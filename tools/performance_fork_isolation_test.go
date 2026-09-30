// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package main_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestPerformanceWorkflowsPreserveForkIsolation(t *testing.T) {
	const fork = "github.event_name == 'pull_request' && github.event.pull_request.head.repo.full_name != github.repository"
	const trusted = "github.event_name != 'pull_request' || github.event.pull_request.head.repo.full_name == github.repository"
	for filename, jobs := range map[string][]string{
		"_build-test.yml":        {"build", "vet", "race", "lint"},
		"_generate-provider.yml": {"generate"},
		"_generate-docs.yml":     {"generate"},
	} {
		content, err := os.ReadFile(filepath.Join("..", ".github", "workflows", filename))
		if err != nil {
			t.Fatal(err)
		}
		var workflow workflowDocument
		if err := yaml.Unmarshal(content, &workflow); err != nil {
			t.Fatal(err)
		}
		for _, name := range jobs {
			t.Run(filename+"/"+name, func(t *testing.T) {
				job := workflow.Jobs[name]
				if scalarString(job["runs-on"]) != "${{ "+fork+" && 'ubuntu-latest' || 'terraform-provider-xcsh-compute' }}" {
					t.Fatalf("fork route is not isolated: %#v", job["runs-on"])
				}
				env, ok := job["env"].(map[string]any)
				if !ok {
					t.Fatal("missing concurrency environment")
				}
				for _, key := range []string{"GOMAXPROCS", "GO_PACKAGE_PARALLELISM"} {
					if scalarString(env[key]) != "${{ "+fork+" && 1 || inputs.go-concurrency }}" {
						t.Fatalf("%s can use trusted concurrency on a fork: %#v", key, env[key])
					}
				}
				if scalarString(env["GOMEMLIMIT"]) != "${{ "+fork+" && '4GiB' || '16GiB' }}" {
					t.Fatal("fork memory limit is not isolated")
				}
				foundSetup, foundAssertion := false, false
				for _, raw := range job["steps"].([]any) {
					step := raw.(map[string]any)
					title := scalarString(step["name"])
					if title == "Verify governed runner input" {
						foundAssertion = true
						if scalarString(step["if"]) != trusted || !strings.Contains(scalarString(step["run"]), `test "$RUNNER_LABEL" = terraform-provider-xcsh-compute`) {
							t.Fatal("trusted runner input is not asserted")
						}
					}
					if strings.HasPrefix(title, "Set up ") && strings.HasSuffix(title, " for fork isolation") || strings.HasPrefix(title, "Install ") && strings.HasSuffix(title, " for fork isolation") {
						if scalarString(step["if"]) != fork {
							t.Fatalf("%s can install tools on trusted runners", title)
						}
						if title == "Set up Go for fork isolation" {
							foundSetup = true
						}
					}
				}
				if !foundSetup || !foundAssertion {
					t.Fatal("missing fork setup or trusted runner assertion")
				}
				if strings.Contains(string(content), "secrets: inherit") {
					t.Fatal("performance workflow inherits caller secrets")
				}
			})
		}
	}
}
