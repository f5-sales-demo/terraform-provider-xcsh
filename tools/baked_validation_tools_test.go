// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package main_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestShellValidationUsesBakedToolsWithForkIsolation(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", ".github", "workflows", "ci.yml"))
	if err != nil {
		t.Fatal(err)
	}
	var workflow workflowDocument
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		t.Fatal(err)
	}
	job := workflow.Jobs["validate-shell-scripts"]
	if job["runs-on"] != "managed-socketless" {
		t.Fatal("trusted shell validation must use the canonical socketless runner")
	}
	guard, _ := job["if"].(string)
	if !strings.Contains(guard, "github.event.pull_request.head.repo.full_name == github.repository") {
		t.Fatal("trusted shell validation must exclude fork code")
	}
	forkData, err := os.ReadFile(filepath.Join("..", ".github", "workflows", "_build-test.yml"))
	if err != nil {
		t.Fatal(err)
	}
	var buildWorkflow workflowDocument
	if err := yaml.Unmarshal(forkData, &buildWorkflow); err != nil {
		t.Fatal(err)
	}
	buildJob := buildWorkflow.Jobs["build"]
	const isolatedBuildRunner = "${{ github.event_name == 'pull_request' && github.event.pull_request.head.repo.full_name != github.repository && 'ubuntu-latest' || 'terraform-provider-xcsh-compute' }}"
	if buildJob["runs-on"] != isolatedBuildRunner {
		t.Fatal("fork shell validation must stay in the approved hosted build shard")
	}
	forkSteps, _ := buildJob["steps"].([]any)
	forkFound := false
	for _, value := range forkSteps {
		step, _ := value.(map[string]any)
		if step["name"] != "Validate shell scripts in the existing hosted fork shard" {
			continue
		}
		forkFound = true
		forkGuard, _ := step["if"].(string)
		if !strings.Contains(forkGuard, "github.event.pull_request.head.repo.full_name != github.repository") {
			t.Fatal("downloaded ShellCheck must be used only for isolated forks")
		}
		script, _ := step["run"].(string)
		for _, required := range []string{"sha256sum --check", "version: 0.11.0", "shellcheck\" --severity=warning", "test-github-api-download.sh", "test-classify-spec-release-semantics.sh"} {
			if !strings.Contains(script, required) {
				t.Fatalf("fork shell validation lost %q", required)
			}
		}
	}
	if !forkFound {
		t.Fatal("fork shell validation is missing")
	}
	steps, _ := job["steps"].([]any)
	found := false
	for _, value := range steps {
		step, _ := value.(map[string]any)
		if step["name"] == "Install verified ShellCheck" {
			guard, _ := step["if"].(string)
			if !strings.Contains(guard, "github.event.pull_request.head.repo.full_name != github.repository") {
				t.Fatal("ShellCheck download must be fork-only")
			}
		}
		if step["name"] == "Verify baked ShellCheck" {
			found = true
			if !strings.Contains(step["run"].(string), "0.11.0") {
				t.Fatal("ShellCheck version is not pinned")
			}
		}
	}
	if !found {
		t.Fatal("baked ShellCheck verification is missing")
	}
}

func TestDocumentationIndexFormatterUsesVerifiedBakedBiome(t *testing.T) {
	script := filepath.Join("..", "scripts", "format-provider-docs-index.sh")
	for _, mode := range []string{"trusted", "fork", "wrong version"} {
		t.Run(mode, func(t *testing.T) {
			tmp := t.TempDir()
			bin := filepath.Join(tmp, "bin")
			if err := os.MkdirAll(bin, 0o700); err != nil {
				t.Fatal(err)
			}
			commands := map[string]string{
				"biome": "#!/bin/sh\nif [ \"$1\" = --version ]; then echo \"$TEST_BIOME_VERSION\"; else printf '%s\\n' baked >\"$TEST_TOOL_LOG\"; fi\n",
				"npx":   "#!/bin/sh\n[ \"$1\" = --yes ] && [ \"$2\" = @biomejs/biome@2.5.6 ] || exit 1\nprintf '%s\\n' fork >\"$TEST_TOOL_LOG\"\n",
			}
			for name, text := range commands {
				if err := os.WriteFile(filepath.Join(bin, name), []byte(text), 0o700); err != nil {
					t.Fatal(err)
				}
			}
			version, fork := "2.5.6", "false"
			if mode == "fork" {
				fork = "true"
			}
			if mode == "wrong version" {
				version = "0.0.0"
			}
			log := filepath.Join(tmp, "tool.log")
			cmd := exec.Command("bash", script)
			cmd.Env = append(os.Environ(), "PATH="+bin+string(os.PathListSeparator)+os.Getenv("PATH"), "PROVIDER_FORK_ISOLATION="+fork, "TEST_BIOME_VERSION="+version, "TEST_TOOL_LOG="+log)
			output, err := cmd.CombinedOutput()
			if mode == "wrong version" {
				if err == nil {
					t.Fatal("wrong tool version accepted")
				}
				return
			}
			if err != nil {
				t.Fatalf("formatter failed: %v %s", err, output)
			}
			got, err := os.ReadFile(log)
			if err != nil {
				t.Fatal(err)
			}
			want := "baked\n"
			if mode == "fork" {
				want = "fork\n"
			}
			if string(got) != want {
				t.Fatalf("formatter used %q want %q", got, want)
			}
		})
	}
}
