// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Exercise the actual workflow shell against real Git patches. These checks
// cover the data boundary that string-only workflow contracts cannot verify.
func TestGenerationArtifactVerification(t *testing.T) {
	for _, kind := range []string{"provider", "combined"} {
		t.Run(kind, func(t *testing.T) {
			workflow, job, step := "_generate-docs.yml", "generate", "Verify and apply provider generation artifact"
			if kind == "combined" {
				workflow, job, step = "on-merge.yml", "create-regeneration-pr", "Verify and apply combined generation artifact"
			}
			script := extractWorkflowRunStep(t, workflow, job, step)
			cases := []string{"changed", "unchanged", "missing patch", "missing manifest", "corrupt patch", "wrong source", "wrong digest", "wrong kind", "wrong patch file", "malformed manifest", "wrong tool version", "wrong concurrency", "unexpected concurrency", "extra metadata", "wrong terraform version", "wrong docs version", "changed empty patch", "unchanged nonempty patch"}
			if kind == "combined" {
				cases = append(cases, "wrong provider digest", "docs changed empty patch", "docs output mismatch", "combined output mismatch", "stale main")
			}
			for _, name := range cases {
				t.Run(name, func(t *testing.T) {
					work := t.TempDir()
					repo := filepath.Join(work, "repo")
					writeReleaseTestFile(t, repo, "fixture.txt", "before\n", 0o600)
					runReleaseTestCommand(t, repo, nil, "git", "init", "-q")
					runReleaseTestCommand(t, repo, nil, "git", "config", "user.name", "Artifact Test")
					runReleaseTestCommand(t, repo, nil, "git", "config", "user.email", "artifact@example.com")
					runReleaseTestCommand(t, repo, nil, "git", "add", "fixture.txt")
					runReleaseTestCommand(t, repo, nil, "git", "commit", "-qm", "fixture")
					source := strings.TrimSpace(runReleaseTestCommand(t, repo, nil, "git", "rev-parse", "HEAD"))
					remote := filepath.Join(work, "remote.git")
					runReleaseTestCommand(t, work, nil, "git", "init", "--bare", "-q", remote)
					runReleaseTestCommand(t, repo, nil, "git", "remote", "add", "origin", remote)
					runReleaseTestCommand(t, repo, nil, "git", "push", "-q", "origin", "HEAD:refs/heads/main")
					if name == "stale main" {
						writeReleaseTestFile(t, repo, "stale.txt", "changed source\n", 0o600)
						runReleaseTestCommand(t, repo, nil, "git", "add", "stale.txt")
						runReleaseTestCommand(t, repo, nil, "git", "commit", "-qm", "advance main")
						runReleaseTestCommand(t, repo, nil, "git", "push", "-q", "origin", "HEAD:refs/heads/main")
						runReleaseTestCommand(t, repo, nil, "git", "checkout", "-q", source)
					}
					writeReleaseTestFile(t, repo, "fixture.txt", "after\n", 0o600)
					patch := runReleaseTestCommandStdout(t, repo, nil, "git", "diff", "--binary", "--full-index", "--no-ext-diff")
					runReleaseTestCommand(t, repo, nil, "git", "restore", "fixture.txt")
					changed := name != "unchanged" && name != "changed empty patch" && name != "docs changed empty patch"
					if !changed {
						patch = ""
					}
					input := "provider-generation-input"
					patchName := "provider-generation.patch"
					if kind == "combined" {
						input, patchName = "combined-generation-input", "combined-generation.patch"
					}
					artifactDir := filepath.Join(work, input)
					writeReleaseTestFile(t, artifactDir, patchName, patch, 0o600)
					digest := "sha256:" + releaseTestSHA(t, filepath.Join(artifactDir, patchName))
					providerDigest := "sha256:" + strings.Repeat("b", 64)
					tools := map[string]any{"go": "go1.25.13"}
					manifest := map[string]any{"schema_version": 1, "artifact_kind": kind, "source_sha": source, "patch_file": patchName, "patch_sha256": digest, "changed": changed, "go_concurrency": 4, "tool_versions": tools}
					if kind == "combined" {
						tools["terraform"], tools["tfplugindocs"] = "Terraform v1.16.3", "v0.25.0"
						manifest["provider_patch_sha256"], manifest["docs_changed"] = providerDigest, changed
					}
					expectedDocs, expectedProvider := "false", "false"
					if changed {
						expectedDocs = "true"
						if kind == "provider" {
							expectedProvider = "true"
						}
					}
					switch name {
					case "wrong source":
						manifest["source_sha"] = strings.Repeat("a", 40)
					case "wrong digest":
						manifest["patch_sha256"] = "sha256:" + strings.Repeat("c", 64)
					case "wrong kind":
						manifest["artifact_kind"] = "unexpected"
					case "wrong patch file":
						manifest["patch_file"] = "unexpected.patch"
					case "wrong tool version":
						tools["go"] = "go0.0.0"
					case "extra metadata":
						manifest["unexpected"] = true
					case "wrong terraform version":
						tools["terraform"] = "Terraform v0.0.0"
					case "wrong docs version":
						tools["tfplugindocs"] = "v0.0.0"
					case "unexpected concurrency":
						manifest["go_concurrency"] = 1
					case "wrong concurrency":
						manifest["go_concurrency"] = -1
					case "changed empty patch":
						manifest["changed"] = true
					case "unchanged nonempty patch":
						manifest["changed"] = false
					case "wrong provider digest":
						manifest["provider_patch_sha256"] = "sha256:" + strings.Repeat("c", 64)
					case "docs changed empty patch":
						manifest["docs_changed"] = true
					case "docs output mismatch":
						expectedDocs = "false"
					case "combined output mismatch":
						manifest["changed"] = false
					}
					manifestPath := filepath.Join(artifactDir, "generation-manifest.json")
					writeReleaseTestJSON(t, manifestPath, manifest)
					switch name {
					case "missing patch":
						if err := os.Remove(filepath.Join(artifactDir, patchName)); err != nil {
							t.Fatal(err)
						}
					case "missing manifest":
						if err := os.Remove(manifestPath); err != nil {
							t.Fatal(err)
						}
					case "corrupt patch":
						writeReleaseTestFile(t, artifactDir, patchName, "corrupt\n", 0o600)
					case "malformed manifest":
						writeReleaseTestFile(t, artifactDir, "generation-manifest.json", "{\n", 0o600)
					}
					output, err := runWorkflowScript(repo, script, []string{"RUNNER_TEMP=" + work, "EXPECTED_DIGEST=" + digest, "EXPECTED_SOURCE_SHA=" + source, "SOURCE_COMMIT=" + source, "GO_PACKAGE_PARALLELISM=4", "EXPECTED_PROVIDER_DIGEST=" + providerDigest, "EXPECTED_DOCS_CHANGED=" + expectedDocs, "EXPECTED_PROVIDER_CHANGED=" + expectedProvider})
					wantSuccess := name == "changed" || name == "unchanged"
					if (err == nil) != wantSuccess {
						t.Fatalf("success=%v want=%v: %s", err == nil, wantSuccess, output)
					}
					if wantSuccess && changed {
						got, readErr := os.ReadFile(filepath.Join(repo, "fixture.txt"))
						if readErr != nil || string(got) != "after\n" {
							t.Fatalf("verified patch was not applied: %v %s", readErr, got)
						}
					}
				})
			}
		})
	}
}

func TestDocsGenerationPropagatesPackageConcurrency(t *testing.T) {
	script := extractWorkflowRunStep(t, "_generate-docs.yml", "generate", "Generate documentation with runner profiling")
	for _, profile := range []bool{false, true} {
		t.Run(map[bool]string{false: "direct", true: "profiled"}[profile], func(t *testing.T) {
			tmp := t.TempDir()
			bin := filepath.Join(tmp, "bin")
			writeReleaseTestFile(t, tmp, "scripts/generate-provider-docs.sh", "#!/bin/sh\n[ \"$GOFLAGS\" = '-p=4' ]\n", 0o700)
			if profile {
				writeReleaseTestFile(t, bin, "runner-profile", "#!/bin/sh\nwhile [ \"$1\" != -- ]; do shift; done\nshift\nexec \"$@\"\n", 0o700)
			}
			// Keep an empty local PATH directory first; runner-profile is not part of
			// the test host contract, so a fixture exists only for the profiled case.
			env := []string{"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"), "RUNNER_TEMP=" + tmp, "GO_PACKAGE_PARALLELISM=4", "GOFLAGS=-p=1"}
			output, err := runWorkflowScript(tmp, script, env)
			if err != nil {
				t.Fatalf("documentation subprocess concurrency failed: %v %s", err, output)
			}
		})
	}
}

func TestRegenerationPublisherPreservesExistingAttestedBranch(t *testing.T) {
	script := extractWorkflowRunStep(t, "on-merge.yml", "create-regeneration-pr", "Create branch and PR")
	tmp := t.TempDir()
	bin := filepath.Join(tmp, "bin")
	writeReleaseTestFile(t, bin, "git", `#!/usr/bin/env bash
set -euo pipefail
if [ "$1" = ls-remote ]; then
  printf '%040d\trefs/heads/auto-regenerate/%040d\n' 2 1
  exit 0
fi
printf 'unexpected mutation: %s\n' "$*" > "$MUTATION_LOG"
exit 2
`, 0o700)
	log := filepath.Join(tmp, "mutation")
	output, err := runWorkflowScript(tmp, script, []string{"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"), "COMMIT_SHA=" + strings.Repeat("a", 40), "MUTATION_LOG=" + log, "PROVIDER_CHANGED=false", "DOCS_CHANGED=false"})
	if err == nil {
		t.Fatalf("existing branch was accepted: %s", output)
	}
	if _, err := os.Stat(log); !os.IsNotExist(err) {
		t.Fatalf("publisher attempted to mutate an existing branch: %v %s", err, output)
	}
}
