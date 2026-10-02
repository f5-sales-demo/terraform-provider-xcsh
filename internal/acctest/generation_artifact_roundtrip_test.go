// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerationArtifactRoundTrip(t *testing.T) {
	for _, scenario := range []string{"provider and docs", "docs only", "unchanged"} {
		t.Run(scenario, func(t *testing.T) {
			work := t.TempDir()
			repo := filepath.Join(work, "repo")
			bin := filepath.Join(work, "bin")
			writeReleaseTestFile(t, repo, "provider.txt", "before\n", 0o600)
			writeReleaseTestFile(t, repo, "docs/fixture.md", "before\n", 0o600)
			verifier, err := os.ReadFile(filepath.Join(testRepositoryRoot(t), "scripts", "verify-tfplugindocs.sh"))
			if err != nil {
				t.Fatal(err)
			}
			writeReleaseTestFile(t, repo, "scripts/verify-tfplugindocs.sh", string(verifier), 0o600)
			staging, err := os.ReadFile(filepath.Join(testRepositoryRoot(t), "scripts", "stage-documentation-manifest.py"))
			if err != nil {
				t.Fatal(err)
			}
			writeReleaseTestFile(t, repo, "scripts/stage-documentation-manifest.py", string(staging), 0o600)
			writeDocumentationManifest := func() {
				if err := os.MkdirAll(filepath.Join(repo, "documentation"), 0o700); err != nil {
					t.Fatal(err)
				}
				fixture := filepath.Join(repo, "docs/fixture.md")
				data, err := os.ReadFile(fixture)
				if err != nil {
					t.Fatal(err)
				}
				writeReleaseTestJSON(t, filepath.Join(repo, "documentation/generated-manifest.json"), map[string]any{
					"files": map[string]any{"docs/fixture.md": map[string]any{"bytes": len(data), "sha256": "sha256:" + releaseTestSHA(t, fixture)}},
				})
			}
			writeDocumentationManifest()

			runReleaseTestCommand(t, repo, nil, "git", "init", "-q")
			runReleaseTestCommand(t, repo, nil, "git", "config", "user.name", "Artifact Test")
			runReleaseTestCommand(t, repo, nil, "git", "config", "user.email", "artifact@example.com")
			runReleaseTestCommand(t, repo, nil, "git", "add", ".")
			runReleaseTestCommand(t, repo, nil, "git", "commit", "-qm", "fixture")
			source := strings.TrimSpace(runReleaseTestCommand(t, repo, nil, "git", "rev-parse", "HEAD"))
			writeReleaseTestFile(t, bin, "go", "#!/bin/sh\nif [ \"$1\" = env ]; then echo go1.25.13; else printf '\\tmod\\tgithub.com/hashicorp/terraform-plugin-docs\\tv0.25.0\\th1:fixture\\n'; fi\n", 0o700)
			writeReleaseTestFile(t, bin, "terraform", "#!/bin/sh\necho '{\"terraform_version\":\"1.16.3\"}'\n", 0o700)
			writeReleaseTestFile(t, bin, "tfplugindocs", "#!/bin/sh\nexit 0\n", 0o700)
			env := []string{"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"), "RUNNER_TEMP=" + work, "GITHUB_RUN_ID=1", "GITHUB_RUN_ATTEMPT=1", "GO_PACKAGE_PARALLELISM=4"}
			providerDigest := ""
			if scenario == "provider and docs" {
				writeReleaseTestFile(t, repo, "provider.txt", "after\n", 0o600)
				outputPath := filepath.Join(work, "provider-output")
				script := extractWorkflowRunStep(t, "_generate-provider.yml", "generate", "Create source-bound provider patch")
				output, err := runWorkflowScript(repo, script, append(env, "GITHUB_OUTPUT="+outputPath))
				if err != nil {
					t.Fatalf("provider emission: %v %s", err, output)
				}
				manifest := readArtifactRoundTripManifest(t, filepath.Join(work, "provider-generation/generation-manifest.json"))
				providerDigest = manifest["patch_sha256"].(string)
				runReleaseTestCommand(t, repo, nil, "git", "restore", "provider.txt")
				if err := os.Rename(filepath.Join(work, "provider-generation"), filepath.Join(work, "provider-generation-input")); err != nil {
					t.Fatal(err)
				}
				apply := extractWorkflowRunStep(t, "_generate-docs.yml", "generate", "Verify and apply provider generation artifact")
				output, err = runWorkflowScript(repo, apply, append(env, "EXPECTED_DIGEST="+providerDigest, "EXPECTED_SOURCE_SHA="+source, "EXPECTED_PROVIDER_CHANGED=true"))
				if err != nil {
					t.Fatalf("docs consumption: %v %s", err, output)
				}
				got, err := os.ReadFile(filepath.Join(repo, "provider.txt"))
				if err != nil || string(got) != "after\n" {
					t.Fatalf("docs did not consume provider patch: %v %s", err, got)
				}
			}
			if scenario != "unchanged" {
				writeReleaseTestFile(t, repo, "docs/fixture.md", "after\n", 0o600)
			}
			writeDocumentationManifest()
			expectedTree := strings.TrimSpace(runReleaseTestCommand(t, repo, nil, "git", "add", "."))
			_ = expectedTree
			expectedTree = strings.TrimSpace(runReleaseTestCommand(t, repo, nil, "git", "write-tree"))
			runReleaseTestCommand(t, repo, nil, "git", "reset", "-q")
			combinedOutput := filepath.Join(work, "combined-output")
			emit := extractWorkflowRunStep(t, "_generate-docs.yml", "generate", "Create source-bound combined patch")
			output, err := runWorkflowScript(repo, emit, append(env, "GITHUB_OUTPUT="+combinedOutput, "PROVIDER_DIGEST="+providerDigest))
			if err != nil {
				t.Fatalf("combined emission: %v %s", err, output)
			}
			manifest := readArtifactRoundTripManifest(t, filepath.Join(work, "combined-generation/generation-manifest.json"))
			if manifest["source_sha"] != source || manifest["provider_patch_sha256"] != providerDigest || manifest["changed"] != (scenario != "unchanged") || manifest["docs_changed"] != (scenario != "unchanged") {
				t.Fatalf("inconsistent manifest: %#v", manifest)
			}
			remote := filepath.Join(work, "remote.git")
			runReleaseTestCommand(t, work, nil, "git", "init", "--bare", "-q", remote)
			runReleaseTestCommand(t, repo, nil, "git", "remote", "add", "origin", remote)
			runReleaseTestCommand(t, repo, nil, "git", "push", "-q", "origin", "HEAD:refs/heads/main")
			runReleaseTestCommand(t, repo, nil, "git", "restore", ".")
			if err := os.Rename(filepath.Join(work, "combined-generation"), filepath.Join(work, "combined-generation-input")); err != nil {
				t.Fatal(err)
			}
			providerChanged, docsChanged := "false", "false"
			if scenario == "provider and docs" {
				providerChanged = "true"
			}
			if scenario != "unchanged" {
				docsChanged = "true"
			}
			publish := extractWorkflowRunStep(t, "on-merge.yml", "create-regeneration-pr", "Verify and apply combined generation artifact")
			output, err = runWorkflowScript(repo, publish, append(env, "SOURCE_COMMIT="+source, "EXPECTED_SOURCE_SHA="+source, "EXPECTED_DIGEST="+manifest["patch_sha256"].(string), "EXPECTED_PROVIDER_DIGEST="+providerDigest, "EXPECTED_PROVIDER_CHANGED="+providerChanged, "EXPECTED_DOCS_CHANGED="+docsChanged))
			if err != nil {
				t.Fatalf("publisher consumption: %v %s", err, output)
			}
			runReleaseTestCommand(t, repo, nil, "git", "add", ".")
			actualTree := strings.TrimSpace(runReleaseTestCommand(t, repo, nil, "git", "write-tree"))
			if actualTree != expectedTree {
				t.Fatalf("combined artifact tree=%s want=%s", actualTree, expectedTree)
			}
		})
	}
}

func readArtifactRoundTripManifest(t *testing.T, path string) map[string]any {
	t.Helper()
	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var manifest map[string]any
	if err := json.Unmarshal(body, &manifest); err != nil {
		t.Fatal(err)
	}
	return manifest
}
