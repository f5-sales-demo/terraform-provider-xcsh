// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

func TestMCPArchiveIsByteReproducible(t *testing.T) {
	root := testRepositoryRoot(t)
	repo := t.TempDir()
	for name, body := range map[string]string{
		"docs/resources/a.md":                    "resource\n",
		"docs/data-sources/a.md":                 "data source\n",
		"docs/functions/a.md":                    "function\n",
		"docs/guides/a.md":                       "guide\n",
		"docs/index.md":                          "index\n",
		"docs/specifications/api/index.json":     "{}\n",
		"docs/specifications/api/domains/a.json": "{}\n",
		"tools/minimum-configs.json":             "{}\n",
	} {
		writeReleaseTestFile(t, repo, name, body, 0o600)
	}
	runReleaseTestCommand(t, repo, nil, "git", "init", "-q")
	runReleaseTestCommand(t, repo, nil, "git", "config", "user.name", "Release Test")
	runReleaseTestCommand(t, repo, nil, "git", "config", "user.email", "release@example.com")
	runReleaseTestCommand(t, repo, nil, "git", "add", ".")
	runReleaseTestCommand(t, repo, []string{"GIT_AUTHOR_DATE=2026-01-02T03:04:05Z", "GIT_COMMITTER_DATE=2026-01-02T03:04:05Z"}, "git", "commit", "-qm", "fixture")
	runReleaseTestCommand(t, repo, nil, "git", "tag", "v1.2.3")

	first := filepath.Join(repo, "first.tar.gz")
	second := filepath.Join(repo, "second.tar.gz")
	script := filepath.Join(root, "scripts", "build-mcp-data.sh")
	runReleaseTestCommand(t, repo, nil, script, "v1.2.3", first)
	if err := os.Chtimes(filepath.Join(repo, "docs", "resources", "a.md"), testFutureTime(), testFutureTime()); err != nil {
		t.Fatal(err)
	}
	runReleaseTestCommand(t, repo, nil, script, "v1.2.3", second)
	firstBytes, err := os.ReadFile(first) //nolint:gosec // isolated test fixture
	if err != nil {
		t.Fatal(err)
	}
	secondBytes, err := os.ReadFile(second) //nolint:gosec // isolated test fixture
	if err != nil {
		t.Fatal(err)
	}
	if string(firstBytes) != string(secondBytes) {
		t.Fatal("MCP archive bytes changed when only source mtimes changed")
	}
	listing := runReleaseTestCommand(t, repo, nil, "tar", "-tzf", first)
	var archiveFiles []string
	for _, entry := range strings.Split(strings.TrimSpace(listing), "\n") {
		if entry != "" && !strings.HasSuffix(entry, "/") {
			archiveFiles = append(archiveFiles, entry)
		}
	}
	sort.Strings(archiveFiles)
	expectedFiles := []string{
		"mcp-data/docs/data-sources/a.md",
		"mcp-data/docs/functions/a.md",
		"mcp-data/docs/guides/a.md",
		"mcp-data/docs/index.md",
		"mcp-data/docs/resources/a.md",
		"mcp-data/docs/specifications/api/index.json",
		"mcp-data/docs/specifications/api/domains/a.json",
		"mcp-data/tools/minimum-configs.json",
	}
	sort.Strings(expectedFiles)
	if strings.Join(archiveFiles, "\n") != strings.Join(expectedFiles, "\n") {
		t.Fatalf("MCP archive file set differs:\nwant:\n%s\ngot:\n%s", strings.Join(expectedFiles, "\n"), strings.Join(archiveFiles, "\n"))
	}
}

func TestProviderReleaseVerifierMeasuresAllAssetsAndSignedChecksums(t *testing.T) {
	root := testRepositoryRoot(t)
	work := t.TempDir()
	assets := filepath.Join(work, "assets")
	if err := os.MkdirAll(assets, 0o750); err != nil {
		t.Fatal(err)
	}
	writeReleaseTestFile(t, work, "tools/spec-release.json", "{\"version\":\"fixture\"}\n", 0o600)
	tag := "v9.8.7"
	version := strings.TrimPrefix(tag, "v")
	names := providerReleaseAssetNames(version)
	for _, name := range names {
		if strings.Contains(name, "SHA256SUMS") {
			continue
		}
		writeReleaseTestFile(t, assets, name, "bytes for "+name+"\n", 0o600)
	}
	checksumName := "terraform-provider-xcsh_" + version + "_SHA256SUMS"
	var checksumLines []string
	for _, name := range names {
		if strings.Contains(name, "SHA256SUMS") || strings.HasPrefix(name, "mcp-data-") {
			continue
		}
		checksumLines = append(checksumLines, releaseTestSHA(t, filepath.Join(assets, name))+"  "+name)
	}
	sort.Strings(checksumLines)
	writeReleaseTestFile(t, assets, checksumName, strings.Join(checksumLines, "\n")+"\n", 0o600)
	gpgHome, err := os.MkdirTemp("/tmp", "provider-release-gpg-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(gpgHome) })
	env := []string{"GNUPGHOME=" + gpgHome}
	runReleaseTestCommand(t, work, env, "gpg", "--batch", "--passphrase", "", "--quick-generate-key", "Release Test <release@example.com>", "rsa2048", "sign", "0")
	runReleaseTestCommand(t, work, env, "gpg", "--batch", "--yes", "--pinentry-mode", "loopback", "--passphrase", "", "--detach-sign", "--output", filepath.Join(assets, checksumName+".sig"), filepath.Join(assets, checksumName))

	releasePath := filepath.Join(work, "release.json")
	writeProviderReleaseJSON(t, releasePath, tag, "", assets, names)
	script := filepath.Join(root, "scripts", "verify-provider-release.sh")
	commit := strings.Repeat("a", 40)
	receipt := runReleaseTestCommandStdout(t, work, env, script, tag, commit, releasePath, assets)
	receiptPath := filepath.Join(work, "receipt.json")
	writeReleaseTestFile(t, work, "receipt.json", receipt, 0o600)
	runReleaseTestCommand(t, work, env, script, tag, commit, releasePath, assets, receiptPath)

	writeReleaseTestFile(t, assets, "mcp-data-"+version+".tar.gz", "tampered\n", 0o600)
	cmd := exec.Command(script, tag, commit, releasePath, assets)
	cmd.Dir = work
	cmd.Env = append(os.Environ(), env...)
	output, err := cmd.CombinedOutput()
	if err == nil || !strings.Contains(string(output), "GitHub digest") {
		t.Fatalf("tampered asset was not rejected by measured digest: err=%v\n%s", err, output)
	}
}

func TestPrepareSpecDeliveryReceiptPersistsCommonAndDetailedEvidence(t *testing.T) {
	fixture := newReceiptFixture(t, false, false)
	runReleaseTestCommand(t, fixture.repo, fixture.env, fixture.script)
	result, err := os.ReadFile(fixture.output) //nolint:gosec // isolated test fixture
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(result), "changed=true") || !strings.Contains(string(result), "delivery_id="+fixture.deliveryID) {
		t.Fatalf("receipt outputs are incomplete:\n%s", result)
	}
	for _, removed := range []string{"tools/spec-pending-delivery.json", "tools/spec-regeneration-receipt.json"} {
		if _, statErr := os.Stat(filepath.Join(fixture.repo, removed)); !os.IsNotExist(statErr) {
			t.Fatalf("%s was not removed after durable receipt", removed)
		}
	}
	assertJSONPathExists(t, filepath.Join(fixture.repo, "tools/spec-deliveries.json"), ".deliveries[\""+fixture.deliveryID+"\"]")
	assertJSONPathExists(t, filepath.Join(fixture.repo, "tools/provider-publication-receipts.json"), ".receipts[\""+fixture.deliveryID+"\"].publication.assets")
}

func TestPrepareSpecDeliveryReceiptRejectsForgedStateAndDigest(t *testing.T) {
	for name, fixture := range map[string]*receiptFixture{
		"noncanonical ledger": newReceiptFixture(t, true, false),
		"false API digest":    newReceiptFixture(t, false, true),
	} {
		t.Run(name, func(t *testing.T) {
			cmd := exec.Command(fixture.script)
			cmd.Dir = fixture.repo
			cmd.Env = append(os.Environ(), fixture.env...)
			output, err := cmd.CombinedOutput()
			if err == nil {
				t.Fatalf("invalid receipt state was accepted:\n%s", output)
			}
			if name == "noncanonical ledger" && !strings.Contains(string(output), "noncanonical key") {
				t.Fatalf("forged ledger failed for the wrong reason:\n%s", output)
			}
			if name == "false API digest" && !strings.Contains(string(output), "measured digest") {
				t.Fatalf("false API digest failed for the wrong reason:\n%s", output)
			}
		})
	}
}

func TestPrepareSpecDeliveryReceiptRejectsMutableProviderRelease(t *testing.T) {
	fixture := newReceiptFixture(t, false, false)
	releasePath := filepath.Join(fixture.repo, "release.json")
	data, err := os.ReadFile(releasePath) //nolint:gosec // isolated fixture
	if err != nil {
		t.Fatal(err)
	}
	var release map[string]any
	if err := json.Unmarshal(data, &release); err != nil {
		t.Fatal(err)
	}
	release["immutable"] = false
	writeReleaseTestJSON(t, releasePath, release)

	cmd := exec.Command(fixture.script)
	cmd.Dir = fixture.repo
	cmd.Env = append(os.Environ(), fixture.env...)
	output, runErr := cmd.CombinedOutput()
	if runErr == nil || !strings.Contains(string(output), "absent or incomplete") {
		t.Fatalf("mutable provider release was accepted: err=%v\n%s", runErr, output)
	}
}

func TestPrepareSpecDeliveryReceiptRetryRecognizesDurableMain(t *testing.T) {
	fixture := newReceiptFixture(t, false, false)
	runReleaseTestCommand(t, fixture.repo, fixture.env, fixture.script)
	runReleaseTestCommand(t, fixture.repo, nil, "git", "add", "tools")
	runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "durable receipt")
	runReleaseTestCommand(t, fixture.repo, nil, "git", "push", "-q", "origin", "HEAD:main")
	releasedCommit := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-list", "-n", "1", "v9.8.7"))
	runReleaseTestCommand(t, fixture.repo, nil, "git", "checkout", "-q", "--detach", releasedCommit)
	if err := os.WriteFile(fixture.output, nil, 0o600); err != nil {
		t.Fatal(err)
	}

	runReleaseTestCommand(t, fixture.repo, fixture.env, fixture.script)
	outputs, err := os.ReadFile(fixture.output) //nolint:gosec // isolated fixture
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(outputs)) != "changed=false" {
		t.Fatalf("durable retry did not produce an exact no-op: %s", outputs)
	}
}

func TestPrepareSpecDeliveryReceiptAcceptsExactRecoveryRebind(t *testing.T) {
	fixture := newReceiptFixture(t, false, false)
	rebindReceiptFixture(t, fixture, "", "")
	runReleaseTestCommand(t, fixture.repo, fixture.env, fixture.script)
}

func TestPrepareSpecDeliveryReceiptRejectsInvalidRecoveryRebind(t *testing.T) {
	for _, tc := range []struct {
		name, previousVersion, sourceCommit, want string
	}{
		{
			name:            "identity changed",
			previousVersion: "2.1.207",
			want:            "may change only its source commit",
		},
		{
			name:         "source is not exact parent",
			sourceCommit: strings.Repeat("a", 40),
			want:         "source is not the exact release parent",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fixture := newReceiptFixture(t, false, false)
			rebindReceiptFixture(t, fixture, tc.previousVersion, tc.sourceCommit)
			cmd := exec.Command(fixture.script)
			cmd.Dir = fixture.repo
			cmd.Env = append(os.Environ(), fixture.env...)
			output, runErr := cmd.CombinedOutput()
			if runErr == nil || !strings.Contains(string(output), tc.want) {
				t.Fatalf("invalid recovery rebind was accepted or failed for the wrong reason: err=%v\n%s", runErr, output)
			}
		})
	}
}

func TestRegenerationCommitRequiresExactPendingAttestation(t *testing.T) {
	script := extractWorkflowRunStep(t, "on-merge.yml", "detect-changes", "Identify an exact regeneration completion commit")
	t.Run("delayed pull request association is retried", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		output := filepath.Join(fixture.repo, "attestation-output")
		callCount := filepath.Join(fixture.repo, "association-call-count")
		env := append(append([]string{}, fixture.env...),
			"TARGET_REPOSITORY="+dispatchTarget,
			"GITHUB_OUTPUT="+output,
			"PR_ASSOCIATION_ATTEMPTS=3",
			"REGENERATION_PR_EMPTY_RESPONSES=2",
			"REGENERATION_PR_CALL_COUNT="+callCount,
		)
		result, err := runWorkflowScript(fixture.repo, script, env)
		if err != nil {
			t.Fatalf("delayed PR association was not retried successfully: %v\n%s", err, result)
		}
		outputs, readErr := os.ReadFile(output) //nolint:gosec // isolated test fixture
		if readErr != nil {
			t.Fatal(readErr)
		}
		if !strings.Contains(string(outputs), "is_regeneration=true") {
			t.Fatalf("delayed exact association was not recognized:\n%s", outputs)
		}
		calls, readErr := os.ReadFile(callCount) //nolint:gosec // isolated test fixture
		if readErr != nil {
			t.Fatal(readErr)
		}
		if strings.TrimSpace(string(calls)) != "3" {
			t.Fatalf("association lookup count = %q, want 3", strings.TrimSpace(string(calls)))
		}
	})

	t.Run("exhausted pull request association fails receipt closed", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		output := filepath.Join(fixture.repo, "attestation-output")
		callCount := filepath.Join(fixture.repo, "association-call-count")
		env := append(append([]string{}, fixture.env...),
			"TARGET_REPOSITORY="+dispatchTarget,
			"GITHUB_OUTPUT="+output,
			"PR_ASSOCIATION_ATTEMPTS=3",
			"REGENERATION_PR_EMPTY_RESPONSES=3",
			"REGENERATION_PR_CALL_COUNT="+callCount,
		)
		result, runErr := runWorkflowScript(fixture.repo, script, env)
		if runErr == nil || !strings.Contains(result, "not visible after 3 attempts") {
			t.Fatalf("exhausted association lookup did not fail closed: err=%v\n%s", runErr, result)
		}
		if strings.Contains(result, "processing as a human change") {
			t.Fatalf("receipt-changing commit was downgraded after exhausted lookup:\n%s", result)
		}
		calls, readErr := os.ReadFile(callCount) //nolint:gosec // isolated test fixture
		if readErr != nil {
			t.Fatal(readErr)
		}
		if strings.TrimSpace(string(calls)) != "3" {
			t.Fatalf("association lookup count = %q, want 3", strings.TrimSpace(string(calls)))
		}
	})

	t.Run("squash subject resolves the exact pull request directly", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "--amend", "-qm", "chore: auto-regenerate provider and documentation (#42)")
		mergeCommit := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD"))
		rewriteRegenerationPRMergeCommit(t, fixture, mergeCommit)
		output := filepath.Join(fixture.repo, "attestation-output")
		callCount := filepath.Join(fixture.repo, "association-call-count")
		env := append(append([]string{}, fixture.env...),
			"TARGET_REPOSITORY="+dispatchTarget,
			"GITHUB_OUTPUT="+output,
			"PR_ASSOCIATION_ATTEMPTS=3",
			"REGENERATION_PR_EMPTY_RESPONSES=3",
			"REGENERATION_PR_CALL_COUNT="+callCount,
		)
		result, err := runWorkflowScript(fixture.repo, script, env)
		if err != nil {
			t.Fatalf("exact squash-subject PR lookup failed: %v\n%s", err, result)
		}
		outputs, readErr := os.ReadFile(output) //nolint:gosec // isolated test fixture
		if readErr != nil {
			t.Fatal(readErr)
		}
		if !strings.Contains(string(outputs), "is_regeneration=true") {
			t.Fatalf("exact squash-subject PR was not recognized:\n%s", outputs)
		}
		if _, statErr := os.Stat(callCount); !os.IsNotExist(statErr) {
			t.Fatalf("commit association endpoint was called despite exact PR number: %v", statErr)
		}
	})

	t.Run("zero generated diff still binds pending delivery", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, err := runWorkflowScript(fixture.repo, script, env)
		if err != nil {
			t.Fatalf("valid attestation failed: %v\n%s", err, result)
		}
		outputs, readErr := os.ReadFile(output) //nolint:gosec // isolated test fixture
		if readErr != nil {
			t.Fatal(readErr)
		}
		if !strings.Contains(string(outputs), "is_regeneration=true") {
			t.Fatalf("valid receipt-only regeneration was not recognized:\n%s", outputs)
		}
	})

	t.Run("pending recovery rebinds an existing receipt", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		receiptPath := filepath.Join(fixture.repo, "tools/spec-regeneration-receipt.json")
		data, err := os.ReadFile(receiptPath) //nolint:gosec // isolated test fixture
		if err != nil {
			t.Fatal(err)
		}
		var receipt map[string]any
		if err := json.Unmarshal(data, &receipt); err != nil {
			t.Fatal(err)
		}

		base := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD^"))
		runReleaseTestCommand(t, fixture.repo, nil, "git", "checkout", "-q", "--detach", base)
		receipt["source_commit"] = strings.Repeat("a", 40)
		writeReleaseTestJSON(t, receiptPath, receipt)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", receiptPath)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "rejected regeneration receipt")
		recoverySource := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD"))

		receipt["source_commit"] = recoverySource
		writeReleaseTestJSON(t, receiptPath, receipt)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", receiptPath)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "chore: auto-regenerate provider and documentation")
		recoveryMerge := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD"))
		rewriteRegenerationPRIdentity(t, fixture, recoveryMerge, recoverySource)

		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, runErr := runWorkflowScript(fixture.repo, script, env)
		if runErr != nil {
			t.Fatalf("valid recovery receipt rebind failed: %v\n%s", runErr, result)
		}
		outputs, readErr := os.ReadFile(output) //nolint:gosec // isolated fixture
		if readErr != nil {
			t.Fatal(readErr)
		}
		if !strings.Contains(string(outputs), "is_regeneration=true") {
			t.Fatalf("recovery receipt rebind was not recognized:\n%s", outputs)
		}
	})

	t.Run("recovery receipt may change only its source commit", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		receiptPath := filepath.Join(fixture.repo, "tools/spec-regeneration-receipt.json")
		data, err := os.ReadFile(receiptPath) //nolint:gosec // isolated test fixture
		if err != nil {
			t.Fatal(err)
		}
		var receipt map[string]any
		if err := json.Unmarshal(data, &receipt); err != nil {
			t.Fatal(err)
		}

		base := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD^"))
		runReleaseTestCommand(t, fixture.repo, nil, "git", "checkout", "-q", "--detach", base)
		receipt["source_commit"] = strings.Repeat("a", 40)
		writeReleaseTestJSON(t, receiptPath, receipt)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", receiptPath)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "rejected regeneration receipt")
		recoverySource := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD"))

		receipt["source_commit"] = recoverySource
		receipt["version"] = "2.1.209"
		writeReleaseTestJSON(t, receiptPath, receipt)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", receiptPath)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "chore: auto-regenerate provider and documentation")
		recoveryMerge := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD"))
		rewriteRegenerationPRIdentity(t, fixture, recoveryMerge, recoverySource)

		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, runErr := runWorkflowScript(fixture.repo, script, env)
		if runErr == nil || !strings.Contains(result, "may change only its source commit") {
			t.Fatalf("non-source recovery receipt change was accepted: err=%v\n%s", runErr, result)
		}
	})

	t.Run("unattested receipt change fails closed", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		writeReleaseTestJSON(t, fixture.regenPR, []map[string]any{})
		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, runErr := runWorkflowScript(fixture.repo, script, env)
		if runErr == nil || !strings.Contains(result, "Merged regeneration PR association was not visible after 6 attempts") {
			t.Fatalf("unattested receipt change was accepted: err=%v\n%s", runErr, result)
		}
	})

	t.Run("pending delivery without receipt fails closed", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "rm", "-q", "tools/spec-regeneration-receipt.json")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "--amend", "--no-edit", "--allow-empty", "-q")
		rewriteRegenerationPRMergeCommit(t, fixture, strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD")))
		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, err := runWorkflowScript(fixture.repo, script, env)
		if err == nil || !strings.Contains(result, "pending delivery can be released only") {
			t.Fatalf("unreceipted pending delivery did not fail closed: err=%v\n%s", err, result)
		}
	})

	t.Run("false source receipt fails closed", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		receiptPath := filepath.Join(fixture.repo, "tools/spec-regeneration-receipt.json")
		data, err := os.ReadFile(receiptPath) //nolint:gosec // isolated test fixture
		if err != nil {
			t.Fatal(err)
		}
		var receipt map[string]any
		if err := json.Unmarshal(data, &receipt); err != nil {
			t.Fatal(err)
		}
		receipt["source_commit"] = strings.Repeat("f", 40)
		writeReleaseTestJSON(t, receiptPath, receipt)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", receiptPath)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "--amend", "--no-edit", "-q")
		rewriteRegenerationPRMergeCommit(t, fixture, strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD")))
		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, runErr := runWorkflowScript(fixture.repo, script, env)
		if runErr == nil || !strings.Contains(result, "receipt source differs from the attested PR source") {
			t.Fatalf("false source receipt did not fail closed: err=%v\n%s", runErr, result)
		}
	})

	t.Run("stale but ancestral source fails closed", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		regenerationCommit := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD"))
		sourceCommit := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD^"))
		runReleaseTestCommand(t, fixture.repo, nil, "git", "checkout", "-q", "--detach", sourceCommit)
		writeReleaseTestFile(t, fixture.repo, "intervening.txt", "main advanced\n", 0o600)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", "intervening.txt")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "intervening main change")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "cherry-pick", regenerationCommit)
		rewriteRegenerationPRMergeCommit(t, fixture, strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD")))
		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, runErr := runWorkflowScript(fixture.repo, script, env)
		if runErr == nil || !strings.Contains(result, "source is not the pre-merge main commit") {
			t.Fatalf("stale ancestral regeneration source was accepted: err=%v\n%s", runErr, result)
		}
	})

	t.Run("manual pending recovery bypasses historical merge attestation", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		regenerationCommit := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD"))
		sourceCommit := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD^"))
		runReleaseTestCommand(t, fixture.repo, nil, "git", "checkout", "-q", "--detach", sourceCommit)
		writeReleaseTestFile(t, fixture.repo, "intervening.txt", "main advanced\n", 0o600)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", "intervening.txt")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "intervening main change")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "cherry-pick", regenerationCommit)
		rewriteRegenerationPRMergeCommit(t, fixture, strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD")))
		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "PENDING_RESUME=true", "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, runErr := runWorkflowScript(fixture.repo, script, env)
		if runErr != nil {
			t.Fatalf("manual recovery inspected the historical regeneration merge: %v\n%s", runErr, result)
		}
		outputs, readErr := os.ReadFile(output) //nolint:gosec // isolated fixture
		if readErr != nil {
			t.Fatal(readErr)
		}
		if strings.TrimSpace(string(outputs)) != "is_regeneration=false" {
			t.Fatalf("manual recovery did not enter the human regeneration path: %s", outputs)
		}
	})

	t.Run("exact subject without PR attestation is a human change", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		parent := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD^"))
		runReleaseTestCommand(t, fixture.repo, nil, "git", "checkout", "-q", "--detach", parent)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "--allow-empty", "-qm", "chore: auto-regenerate provider and documentation")
		writeReleaseTestJSON(t, fixture.regenPR, []map[string]any{})
		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, runErr := runWorkflowScript(fixture.repo, script, env)
		if runErr != nil {
			t.Fatalf("human exact-subject commit failed instead of taking the tested path: %v\n%s", runErr, result)
		}
		outputs, readErr := os.ReadFile(output) //nolint:gosec // isolated fixture
		if readErr != nil {
			t.Fatal(readErr)
		}
		if strings.TrimSpace(string(outputs)) != "is_regeneration=false" {
			t.Fatalf("unattested exact subject bypassed build/test: %s", outputs)
		}
	})

	t.Run("durable receipt merge is accepted only with exact pull request attestation", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		runReleaseTestCommand(t, fixture.repo, fixture.env, fixture.script)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", "tools")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "chore: receipt published spec delivery (#43)")
		mergeCommit := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD"))
		writeReleaseTestJSON(t, fixture.regenPR, []map[string]any{{
			"number":           43,
			"state":            "closed",
			"merged_at":        "2026-08-01T12:00:00Z",
			"merge_commit_sha": mergeCommit,
			"title":            "chore: receipt published spec delivery",
			"base": map[string]any{
				"ref": "main", "repo": map[string]any{"full_name": dispatchTarget},
			},
			"head": map[string]any{
				"ref":  "spec-delivery-receipt/" + fixture.deliveryID,
				"repo": map[string]any{"full_name": dispatchTarget},
			},
		}})
		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, runErr := runWorkflowScript(fixture.repo, script, env)
		if runErr != nil {
			t.Fatalf("exact durable receipt merge was rejected: %v\n%s", runErr, result)
		}
		outputs, readErr := os.ReadFile(output) //nolint:gosec // isolated fixture
		if readErr != nil {
			t.Fatal(readErr)
		}
		if strings.TrimSpace(string(outputs)) != "is_regeneration=false" ||
			!strings.Contains(result, "Detected exact durable delivery receipt merge") {
			t.Fatalf("durable receipt merge was not classified as metadata-only:\n%s\n%s", outputs, result)
		}
	})

	t.Run("unattested durable receipt merge fails closed", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		runReleaseTestCommand(t, fixture.repo, fixture.env, fixture.script)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", "tools")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "chore: receipt published spec delivery (#43)")
		writeReleaseTestJSON(t, fixture.regenPR, []map[string]any{})
		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, runErr := runWorkflowScript(fixture.repo, script, env)
		if runErr == nil || !strings.Contains(result, "Receipt merge has no exact pull request attestation") {
			t.Fatalf("unattested durable receipt merge was accepted: err=%v\n%s", runErr, result)
		}
	})

	t.Run("attested receipt merge with an extra change fails closed", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		runReleaseTestCommand(t, fixture.repo, fixture.env, fixture.script)
		writeReleaseTestFile(t, fixture.repo, "unexpected.txt", "not receipt metadata\n", 0o600)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", ".")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "chore: receipt published spec delivery (#43)")
		mergeCommit := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD"))
		writeReleaseTestJSON(t, fixture.regenPR, []map[string]any{{
			"number": 43, "state": "closed", "merged_at": "2026-08-01T12:00:00Z",
			"merge_commit_sha": mergeCommit, "title": "chore: receipt published spec delivery",
			"base": map[string]any{"ref": "main", "repo": map[string]any{"full_name": dispatchTarget}},
			"head": map[string]any{
				"ref":  "spec-delivery-receipt/" + fixture.deliveryID,
				"repo": map[string]any{"full_name": dispatchTarget},
			},
		}})
		output := filepath.Join(fixture.repo, "attestation-output")
		env := append(append([]string{}, fixture.env...), "TARGET_REPOSITORY="+dispatchTarget, "GITHUB_OUTPUT="+output)
		result, runErr := runWorkflowScript(fixture.repo, script, env)
		if runErr == nil || !strings.Contains(result, "outside the exact durable receipt set") {
			t.Fatalf("receipt merge with an extra change was accepted: err=%v\n%s", runErr, result)
		}
	})
}

func TestOnMergeProcessingPrioritizesExactRegenerationIdentity(t *testing.T) {
	script := extractWorkflowRunStep(t, "on-merge.yml", "detect-changes", "Determine if processing should proceed")
	run := func(t *testing.T, isRegeneration string) (string, string, error) {
		t.Helper()
		tmp := t.TempDir()
		output := filepath.Join(tmp, "github-output")
		env := []string{
			"IS_REGENERATION=" + isRegeneration,
			"SPECS=false",
			"CODE=false",
			"TOOLS=false",
			"FUNCTIONS=false",
			"HAS_ANY_CHANGES=false",
			"DELIVERY_METADATA_ONLY=true",
			"PENDING_DELIVERY=false",
			"GITHUB_OUTPUT=" + output,
		}
		result, err := runWorkflowScript(tmp, script, env)
		outputs, readErr := os.ReadFile(output) //nolint:gosec // isolated fixture
		if readErr != nil {
			t.Fatal(readErr)
		}
		return string(outputs), result, err
	}

	t.Run("attested receipt-only regeneration proceeds", func(t *testing.T) {
		outputs, result, err := run(t, "true")
		if err != nil {
			t.Fatalf("exact regeneration classification failed: %v\n%s", err, result)
		}
		if !strings.Contains(outputs, "result=true") || !strings.Contains(result, "exact regeneration completion") {
			t.Fatalf("exact receipt-only regeneration did not proceed:\n%s\n%s", outputs, result)
		}
	})

	t.Run("ordinary receipt-only commit remains skipped", func(t *testing.T) {
		outputs, result, err := run(t, "false")
		if err != nil {
			t.Fatalf("ordinary metadata classification failed: %v\n%s", err, result)
		}
		if !strings.Contains(outputs, "result=false") || !strings.Contains(result, "metadata cannot trigger") {
			t.Fatalf("ordinary receipt-only commit was not skipped:\n%s\n%s", outputs, result)
		}
	})
}

func TestOnMergeGeneratorStateTruthTable(t *testing.T) {
	script := extractWorkflowRunStep(t, "on-merge.yml", "generation-state", "Classify generator state")
	type generatorState struct {
		name    string
		result  string
		changed string
	}
	states := []generatorState{
		{name: "failure", result: "failure"},
		{name: "cancelled", result: "cancelled"},
		{name: "skipped", result: "skipped"},
		{name: "success-unchanged", result: "success", changed: "false"},
		{name: "success-changed", result: "success", changed: "true"},
	}
	run := func(t *testing.T, extraEnv ...string) (string, string, error) {
		t.Helper()
		tmp := t.TempDir()
		output := filepath.Join(tmp, "github-output")
		values := map[string]string{
			"DETECT_RESULT":                "success",
			"SHOULD_PROCESS":               "true",
			"IS_REGENERATION":              "false",
			"PENDING_DELIVERY":             "false",
			"RECOVERY_INFRASTRUCTURE_ONLY": "false",
			"BUILD_RESULT":                 "success",
			"PROVIDER_REQUIRED":            "true",
			"PROVIDER_RESULT":              "success",
			"PROVIDER_CHANGED":             "false",
			"DOCS_REQUIRED":                "true",
			"DOCS_RESULT":                  "success",
			"DOCS_CHANGED":                 "false",
			"GITHUB_OUTPUT":                output,
		}
		for _, assignment := range extraEnv {
			key, value, ok := strings.Cut(assignment, "=")
			if !ok {
				t.Fatalf("invalid environment assignment %q", assignment)
			}
			values[key] = value
		}
		env := make([]string, 0, len(values))
		for key, value := range values {
			env = append(env, key+"="+value)
		}
		result, err := runWorkflowScript(tmp, script, env)
		outputs, readErr := os.ReadFile(output) //nolint:gosec // isolated test output
		if readErr != nil && !os.IsNotExist(readErr) {
			t.Fatal(readErr)
		}
		return string(outputs), result, err
	}

	for _, provider := range states {
		for _, docs := range states {
			t.Run("provider-"+provider.name+"-docs-"+docs.name, func(t *testing.T) {
				outputs, result, err := run(t,
					"PROVIDER_RESULT="+provider.result,
					"PROVIDER_CHANGED="+provider.changed,
					"DOCS_RESULT="+docs.result,
					"DOCS_CHANGED="+docs.changed,
				)
				bothSucceeded := provider.result == "success" && docs.result == "success"
				if !bothSucceeded {
					if err == nil {
						t.Fatalf("non-success generator state authorized publication:\n%s", outputs)
					}
					if strings.Contains(outputs, "create_pr=true") || strings.Contains(outputs, "release=true") {
						t.Fatalf("failed generator state emitted authorization: err=%v\n%s\n%s", err, result, outputs)
					}
					return
				}

				if err != nil {
					t.Fatalf("successful generator state was rejected: %v\n%s", err, result)
				}
				wantCreate := provider.changed == "true" || docs.changed == "true"
				if gotCreate := strings.Contains(outputs, "create_pr=true"); gotCreate != wantCreate {
					t.Fatalf("create_pr authorization = %t, want %t:\n%s", gotCreate, wantCreate, outputs)
				}
				if gotRelease := strings.Contains(outputs, "release=true"); gotRelease == wantCreate {
					t.Fatalf("release authorization was not the complement of create_pr:\n%s", outputs)
				}
			})
		}
	}

	t.Run("successful generator missing changed output", func(t *testing.T) {
		outputs, result, err := run(t, "PROVIDER_CHANGED=")
		if err == nil || !strings.Contains(result, "no exact changed result") {
			t.Fatalf("missing changed output was not rejected: err=%v\n%s", err, result)
		}
		if strings.Contains(outputs, "create_pr=true") || strings.Contains(outputs, "release=true") {
			t.Fatalf("missing changed output emitted authorization:\n%s", outputs)
		}
	})

	t.Run("pending delivery requires regeneration PR", func(t *testing.T) {
		outputs, result, err := run(t, "PENDING_DELIVERY=true")
		if err != nil {
			t.Fatalf("pending delivery with successful generators failed: %v\n%s", err, result)
		}
		if !strings.Contains(outputs, "create_pr=true") || strings.Contains(outputs, "release=true") {
			t.Fatalf("pending delivery bypassed regeneration PR:\n%s", outputs)
		}
	})

	t.Run("recovery infrastructure cannot release directly", func(t *testing.T) {
		outputs, result, err := run(t, "RECOVERY_INFRASTRUCTURE_ONLY=true")
		if err != nil {
			t.Fatalf("recovery-only state failed classification: %v\n%s", err, result)
		}
		if strings.Contains(outputs, "create_pr=true") || strings.Contains(outputs, "release=true") {
			t.Fatalf("recovery-only state authorized publication:\n%s", outputs)
		}

		outputs, result, err = run(t,
			"RECOVERY_INFRASTRUCTURE_ONLY=true",
			"PROVIDER_CHANGED=true",
		)
		if err != nil {
			t.Fatalf("recovery generated-change state failed: %v\n%s", err, result)
		}
		if !strings.Contains(outputs, "create_pr=true") || strings.Contains(outputs, "release=true") {
			t.Fatalf("recovery generated change bypassed regeneration PR:\n%s", outputs)
		}
	})

	t.Run("attested regeneration is sole generator bypass", func(t *testing.T) {
		outputs, result, err := run(t,
			"IS_REGENERATION=true",
			"BUILD_RESULT=skipped",
			"PROVIDER_RESULT=skipped",
			"PROVIDER_CHANGED=",
			"DOCS_RESULT=skipped",
			"DOCS_CHANGED=",
		)
		if err != nil {
			t.Fatalf("attested regeneration did not bypass generators: %v\n%s", err, result)
		}
		if !strings.Contains(outputs, "release=true") || strings.Contains(outputs, "create_pr=true") {
			t.Fatalf("attested regeneration did not authorize only release:\n%s", outputs)
		}

		outputs, result, err = run(t,
			"BUILD_RESULT=skipped",
			"PROVIDER_RESULT=skipped",
			"DOCS_RESULT=skipped",
		)
		if err == nil || !strings.Contains(result, "exact successful build") {
			t.Fatalf("unattested state bypassed build and generators: err=%v\n%s", err, result)
		}
		if strings.Contains(outputs, "create_pr=true") || strings.Contains(outputs, "release=true") {
			t.Fatalf("unattested bypass emitted authorization:\n%s", outputs)
		}
	})

	t.Run("non-applicable generators must be skipped", func(t *testing.T) {
		outputs, result, err := run(t,
			"PROVIDER_REQUIRED=false",
			"PROVIDER_RESULT=skipped",
			"PROVIDER_CHANGED=",
			"DOCS_REQUIRED=false",
			"DOCS_RESULT=skipped",
			"DOCS_CHANGED=",
		)
		if err != nil {
			t.Fatalf("properly skipped non-applicable generators failed: %v\n%s", err, result)
		}
		if !strings.Contains(outputs, "release=true") || strings.Contains(outputs, "create_pr=true") {
			t.Fatalf("direct release was not authorized for successful build with no applicable generators:\n%s", outputs)
		}

		outputs, result, err = run(t,
			"PROVIDER_REQUIRED=false",
			"PROVIDER_RESULT=success",
			"DOCS_REQUIRED=false",
			"DOCS_RESULT=skipped",
		)
		if err == nil || !strings.Contains(result, "not skipped") {
			t.Fatalf("non-applicable successful generator was accepted: err=%v\n%s", err, result)
		}
		if strings.Contains(outputs, "create_pr=true") || strings.Contains(outputs, "release=true") {
			t.Fatalf("invalid non-applicable generator emitted authorization:\n%s", outputs)
		}
	})

	t.Run("failed build denies publication", func(t *testing.T) {
		outputs, result, err := run(t, "BUILD_RESULT=failure")
		if err == nil || !strings.Contains(result, "exact successful build") {
			t.Fatalf("failed build was not rejected: err=%v\n%s", err, result)
		}
		if strings.Contains(outputs, "create_pr=true") || strings.Contains(outputs, "release=true") {
			t.Fatalf("failed build emitted authorization:\n%s", outputs)
		}
	})
}

func TestOnMergePublicationJobsConsumeOnlyClassifierAuthorization(t *testing.T) {
	workflowPath := filepath.Join(testRepositoryRoot(t), ".github", "workflows", "on-merge.yml")
	workflowBytes, err := os.ReadFile(workflowPath) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	var workflow struct {
		Jobs map[string]struct {
			If      string            `yaml:"if"`
			Outputs map[string]string `yaml:"outputs"`
		} `yaml:"jobs"`
	}
	if err := yaml.Unmarshal(workflowBytes, &workflow); err != nil {
		t.Fatal(err)
	}
	classifier := workflow.Jobs["generation-state"]
	if classifier.Outputs["create-pr"] != "${{ steps.classify.outputs.create_pr }}" ||
		classifier.Outputs["release"] != "${{ steps.classify.outputs.release }}" {
		t.Fatalf("generation-state does not own both publication authorizations: %+v", classifier.Outputs)
	}

	createCondition := workflow.Jobs["create-regeneration-pr"].If
	if !strings.Contains(createCondition, "needs.generation-state.result == 'success'") ||
		!strings.Contains(createCondition, "needs.generation-state.outputs.create-pr == 'true'") {
		t.Fatalf("regeneration PR is not authorized by generation-state:\n%s", createCondition)
	}
	releaseCondition := workflow.Jobs["tag-release"].If
	if !strings.Contains(releaseCondition, "needs.generation-state.result == 'success'") ||
		!strings.Contains(releaseCondition, "needs.generation-state.outputs.release == 'true'") ||
		!strings.Contains(releaseCondition, "needs.create-regeneration-pr.result == 'skipped'") {
		t.Fatalf("release is not authorized by generation-state and a skipped PR job:\n%s", releaseCondition)
	}
	for job, condition := range map[string]string{
		"create-regeneration-pr": createCondition,
		"tag-release":            releaseCondition,
	} {
		for _, forbidden := range []string{
			"needs.detect-changes.outputs",
			"needs.build-test",
			"needs.regenerate-provider.result",
			"needs.regenerate-provider.outputs",
			"needs.regenerate-docs.result",
			"needs.regenerate-docs.outputs",
		} {
			if strings.Contains(condition, forbidden) {
				t.Fatalf("%s bypasses classifier authorization through %q:\n%s", job, forbidden, condition)
			}
		}
	}
}

func TestProviderPublicationWorkflowsRequireCanonicalToken(t *testing.T) {
	root := testRepositoryRoot(t)
	for _, workflowName := range []string{"discover-defaults.yml", "on-merge.yml", "sync-openapi.yml"} {
		workflowBytes, err := os.ReadFile(filepath.Join(root, ".github", "workflows", workflowName)) //nolint:gosec // fixed repository path
		if err != nil {
			t.Fatal(err)
		}
		workflowText := string(workflowBytes)
		if strings.Contains(workflowText, "AUTO_MERGE_TOKEN") {
			t.Fatalf("%s retains the obsolete downstream-triggering token name", workflowName)
		}
		if !strings.Contains(workflowText, "secrets.REPO_SYNC_TOKEN") {
			t.Fatalf("%s does not use the canonical downstream-triggering token", workflowName)
		}
	}

	for name, fixture := range map[string]struct {
		workflow string
		job      string
		want     string
	}{
		"default discovery":        {workflow: "discover-defaults.yml", job: "discover", want: "default discovery"},
		"openapi sync":             {workflow: "sync-openapi.yml", job: "sync", want: "OpenAPI synchronization"},
		"regeneration publication": {workflow: "on-merge.yml", job: "create-regeneration-pr", want: "regeneration publication"},
		"delivery receipting":      {workflow: "on-merge.yml", job: "receipt-spec-delivery", want: "delivery receipting"},
	} {
		t.Run(name, func(t *testing.T) {
			script := extractWorkflowRunStep(t, fixture.workflow, fixture.job, "Require downstream-triggering token")
			result, runErr := runWorkflowScript(t.TempDir(), script, nil)
			if runErr == nil || !strings.Contains(result, "REPO_SYNC_TOKEN is required for "+fixture.want) {
				t.Fatalf("missing canonical PAT was accepted: err=%v\n%s", runErr, result)
			}
			result, runErr = runWorkflowScript(t.TempDir(), script, []string{"REPO_SYNC_TOKEN=fixture"})
			if runErr != nil {
				t.Fatalf("present canonical PAT was rejected: %v\n%s", runErr, result)
			}
		})
	}
}

func TestDefaultDiscoveryPublishesOnlySanitizedEvidence(t *testing.T) {
	workflowPath := filepath.Join(testRepositoryRoot(t), ".github", "workflows", "discover-defaults.yml")
	workflowBytes, err := os.ReadFile(workflowPath) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	workflowText := string(workflowBytes)
	for _, forbidden := range []string{
		"AUTO_MERGE_TOKEN",
		"XCSH_P12_FILE",
		"XCSH_P12_PASSWORD",
		"continue-on-error",
		"go-version:",
		"discovery.log",
	} {
		if strings.Contains(workflowText, forbidden) {
			t.Fatalf("default discovery retains fail-open or compatibility path %q", forbidden)
		}
	}

	script := extractWorkflowRunStep(t, "discover-defaults.yml", "discover", "Run discovery")
	for _, tc := range []struct {
		name, failed, toolExit, want string
		wantSuccess                  bool
	}{
		{name: "success", failed: "0", toolExit: "0", wantSuccess: true},
		{name: "measured resource failure", failed: "1", toolExit: "0", want: "reported failures"},
		{name: "tool failure", failed: "0", toolExit: "7", want: "reported failures"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tmp := t.TempDir()
			bin := filepath.Join(tmp, "bin")
			writeReleaseTestFile(t, bin, "go", `#!/usr/bin/env bash
printf 'tenant identifier that must stay private\n'
printf 'Discovered:      3\n'
printf 'Failed:          %s\n' "$STUB_FAILED"
exit "$STUB_EXIT"
`, 0o700)
			outputPath := filepath.Join(tmp, "github-output")
			result, runErr := runImplicitWorkflowScript(tmp, script, []string{
				"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"),
				"GITHUB_OUTPUT=" + outputPath,
				"INPUT_MODE=discover",
				"INPUT_RESOURCE=",
				"STUB_FAILED=" + tc.failed,
				"STUB_EXIT=" + tc.toolExit,
			})
			if tc.wantSuccess && runErr != nil {
				t.Fatalf("successful discovery failed: %v\n%s", runErr, result)
			}
			if !tc.wantSuccess && (runErr == nil || !strings.Contains(result, tc.want)) {
				t.Fatalf("failed discovery was reported as success: err=%v\n%s", runErr, result)
			}
			if strings.Contains(result, "tenant identifier") {
				t.Fatalf("raw live-tenant output reached the workflow log:\n%s", result)
			}
			if _, err := os.Stat(filepath.Join(tmp, "discovery-raw.log")); !os.IsNotExist(err) {
				t.Fatalf("raw live-tenant output survived discovery: %v", err)
			}
			evidence, err := os.ReadFile(filepath.Join(tmp, "discovery-evidence.json")) //nolint:gosec // isolated fixture
			if err != nil {
				t.Fatal(err)
			}
			var measured struct {
				Discovered int `json:"discovered"`
				Failed     int `json:"failed"`
				ToolExit   int `json:"tool_exit"`
			}
			if err := json.Unmarshal(evidence, &measured); err != nil {
				t.Fatal(err)
			}
			if measured.Discovered != 3 || strconv.Itoa(measured.Failed) != tc.failed ||
				strconv.Itoa(measured.ToolExit) != tc.toolExit {
				t.Fatalf("sanitized discovery evidence is wrong: %+v", measured)
			}
		})
	}
}

func TestRegenerationStaleBranchDeletionUsesPAT(t *testing.T) {
	workflowPath := filepath.Join(testRepositoryRoot(t), ".github", "workflows", "on-merge.yml")
	workflowBytes, err := os.ReadFile(workflowPath) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	var workflow struct {
		Jobs map[string]struct {
			Steps []struct {
				Name string         `yaml:"name"`
				With map[string]any `yaml:"with"`
			} `yaml:"steps"`
		} `yaml:"jobs"`
	}
	if err := yaml.Unmarshal(workflowBytes, &workflow); err != nil {
		t.Fatal(err)
	}
	foundPublicDownload := false
	for _, step := range workflow.Jobs["create-regeneration-pr"].Steps {
		if step.Name == "Download API specs" {
			foundPublicDownload = true
			if step.With["token"] != "${{ github.token }}" {
				t.Fatalf("public spec download receives the write PAT: %+v", step.With)
			}
		}
	}
	if !foundPublicDownload {
		t.Fatal("create-regeneration-pr public spec download step is missing")
	}

	script := extractWorkflowRunStep(t, "on-merge.yml", "create-regeneration-pr", "Close stale auto-regenerate PRs")
	tmp := t.TempDir()
	bin := filepath.Join(tmp, "bin")
	statePath := filepath.Join(tmp, "remote-url")
	branch := "auto-regenerate/" + strings.Repeat("b", 40)
	writeReleaseTestFile(t, bin, "gh", fmt.Sprintf(`#!/usr/bin/env bash
set -euo pipefail
case "$1 $2" in
  "pr list") printf '42\n' ;;
  "pr view") printf '%%s\n' %q ;;
  "pr close") ;;
  *) exit 2 ;;
esac
`, branch), 0o700)
	writeReleaseTestFile(t, bin, "git", `#!/usr/bin/env bash
set -euo pipefail
case "$1 $2" in
  "config --unset-all") exit 0 ;;
  "remote set-url") printf '%s\n' "$4" > "$REMOTE_STATE" ;;
  "ls-remote origin") printf '%040d\t%s\n' 1 "$3" ;;
  "push origin") grep -q '^https://x-access-token:fixture@github.com/f5-sales-demo/terraform-provider-xcsh.git$' "$REMOTE_STATE" ;;
  *) exit 2 ;;
esac
`, 0o700)
	result, runErr := runWorkflowScript(tmp, script, []string{
		"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"),
		"GH_TOKEN=fixture",
		"COMMIT_SHA=" + strings.Repeat("a", 40),
		"REPOSITORY=" + dispatchTarget,
		"REMOTE_STATE=" + statePath,
	})
	if runErr != nil {
		t.Fatalf("stale branch deletion did not authenticate git with the PAT: %v\n%s", runErr, result)
	}
}

func TestReleaseWorkflowSerializesAndClassifiesReleaseState(t *testing.T) {
	root := testRepositoryRoot(t)
	data, err := os.ReadFile(filepath.Join(root, ".github/workflows/_tag-release.yml")) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	var workflow struct {
		Concurrency struct {
			Group            string `yaml:"group"`
			CancelInProgress bool   `yaml:"cancel-in-progress"`
		} `yaml:"concurrency"`
	}
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		t.Fatal(err)
	}
	if workflow.Concurrency.Group != "provider-release-transaction" || workflow.Concurrency.CancelInProgress {
		t.Fatalf("release workflow is not globally serialized: %+v", workflow.Concurrency)
	}

	script := extractWorkflowRunStep(t, "_tag-release.yml", "publish", "Resolve release transaction state")
	for name, fixture := range map[string]struct {
		mode      string
		wantState string
		wantError string
	}{
		"absent":    {mode: "absent", wantState: "absent"},
		"draft":     {mode: "draft", wantState: "draft"},
		"published": {mode: "published", wantState: "published"},
		"duplicate": {mode: "duplicate", wantError: "Release tag resolves to multiple releases"},
		"forbidden": {mode: "forbidden", wantError: "Failed to resolve existing release state"},
	} {
		t.Run(name, func(t *testing.T) {
			tmp := t.TempDir()
			bin := filepath.Join(tmp, "bin")
			stub := `#!/usr/bin/env bash
set -euo pipefail
if [[ "$*" == *immutable-releases* ]]; then
  printf '%s\n' '{"enabled":true}'
  exit 0
fi
if [[ "$*" == *"releases/tags"* ]]; then
  case "$RELEASE_MODE" in
    absent|draft|duplicate) printf '%s\n' 'gh: Not Found (HTTP 404)' >&2; exit 1 ;;
    published) printf '%s\n' '{"tag_name":"v1.2.3","draft":false,"prerelease":false}' ;;
    forbidden) printf '%s\n' 'gh: Forbidden (HTTP 403)' >&2; exit 1 ;;
  esac
  exit 0
fi
if [[ "$*" == *"releases?per_page=100"* ]]; then
  case "$RELEASE_MODE" in
    absent) printf '%s\n' '[]' ;;
    draft) printf '%s\n' '[{"tag_name":"v1.2.3","draft":true,"prerelease":false}]' ;;
    duplicate) printf '%s\n' '[{"tag_name":"v1.2.3","draft":true,"prerelease":false},{"tag_name":"v1.2.3","draft":true,"prerelease":false}]' ;;
    *) exit 88 ;;
  esac
  exit 0
fi
exit 88
`
			writeReleaseTestFile(t, bin, "gh", stub, 0o700)
			output := filepath.Join(tmp, "github-output")
			result, runErr := runWorkflowScript(tmp, script, []string{
				"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"),
				"RELEASE_MODE=" + fixture.mode,
				"RELEASE_TAG=v1.2.3",
				"REPOSITORY=f5-sales-demo/terraform-provider-xcsh",
				"RUNNER_TEMP=" + tmp,
				"GITHUB_OUTPUT=" + output,
				"GH_TOKEN=fixture",
				"REPOSITORY_ADMINISTRATION_TOKEN=administration-fixture",
			})
			if fixture.wantError != "" {
				if runErr == nil || !strings.Contains(result, fixture.wantError) {
					t.Fatalf("release probe failure was misclassified: err=%v\n%s", runErr, result)
				}
				return
			}
			if runErr != nil {
				t.Fatalf("release state %s failed: %v\n%s", fixture.wantState, runErr, result)
			}
			outputs, readErr := os.ReadFile(output) //nolint:gosec // isolated fixture
			if readErr != nil {
				t.Fatal(readErr)
			}
			if !strings.Contains(string(outputs), "state="+fixture.wantState) {
				t.Fatalf("release state was not classified as %s: %s", fixture.wantState, outputs)
			}
		})
	}
}

func TestMutableDraftLookupUsesReleaseCollection(t *testing.T) {
	measure := extractWorkflowRunStep(t, "_tag-release.yml", "publish", "Download and measure all release assets")
	prepublish := extractWorkflowRunStep(t, "_tag-release.yml", "publish", "Publish verified draft")
	published := extractWorkflowRunStep(t, "_tag-release.yml", "publish", "Verify sealed publication")

	for name, script := range map[string]string{
		"measure":    measure,
		"prepublish": prepublish,
	} {
		for _, fragment := range []string{
			`releases?per_page=100`,
			`--paginate`,
			`.tag_name == $tag`,
			`length == 1`,
			`.draft == true`,
		} {
			if !strings.Contains(script, fragment) {
				t.Errorf("%s does not resolve exactly one stable draft from the release collection; missing %q", name, fragment)
			}
		}
	}
	if strings.Contains(prepublish, `releases/tags/${RELEASE_TAG}`) {
		t.Fatal("prepublication draft re-measurement still uses the published-release tag endpoint")
	}
	if !strings.Contains(measure, `if [ "$RELEASE_STATE" = "published" ]; then`) ||
		!strings.Contains(measure, `releases/tags/${RELEASE_TAG}`) {
		t.Fatal("asset measurement does not reserve the exact tag endpoint for an already-published release")
	}
	if !strings.Contains(published, `releases/tags/${RELEASE_TAG}`) {
		t.Fatal("sealed publication verification no longer uses the exact published-release tag endpoint")
	}
}

func TestMCPArchiveDoesNotDirtyGoReleaserCheckout(t *testing.T) {
	buildScript := extractWorkflowRunStep(t, "_tag-release.yml", "publish", "Build MCP archive twice")
	canonicalArchive := `mcp-data-${VERSION}.tar.gz`
	if strings.Contains(buildScript, canonicalArchive) {
		t.Fatalf("MCP reproducibility step writes the canonical archive into the GoReleaser checkout:\n%s", buildScript)
	}

	attachScript := extractWorkflowRunStep(t, "_tag-release.yml", "publish", "Attach deterministic MCP archive to draft")
	copyCommand := `cp "$RUNNER_TEMP/mcp-a.tar.gz" "mcp-data-${VERSION}.tar.gz"`
	uploadCommand := `gh release upload "$RELEASE_TAG" "mcp-data-${VERSION}.tar.gz" --clobber`
	copyIndex := strings.Index(attachScript, copyCommand)
	uploadIndex := strings.Index(attachScript, uploadCommand)
	if copyIndex < 0 || uploadIndex < 0 || copyIndex >= uploadIndex {
		t.Fatalf("MCP archive is not copied to its canonical name immediately before upload:\n%s", attachScript)
	}
}

func TestReleaseImmutabilityChecksUseAdministrationToken(t *testing.T) {
	root := testRepositoryRoot(t)
	releaseData, err := os.ReadFile(filepath.Join(root, ".github/workflows/_tag-release.yml")) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	releaseWorkflow := string(releaseData)
	if strings.Count(releaseWorkflow, "repos/${REPOSITORY}/immutable-releases") != 2 {
		t.Fatal("release workflow must check immutable-release policy before tagging and publishing")
	}
	if !strings.Contains(releaseWorkflow, "GH_TOKEN: ${{ secrets.repository-administration-token }}") ||
		!strings.Contains(releaseWorkflow, "REPOSITORY_ADMINISTRATION_TOKEN: ${{ secrets.repository-administration-token }}") ||
		!strings.Contains(releaseWorkflow, "GH_TOKEN=\"$REPOSITORY_ADMINISTRATION_TOKEN\"") {
		t.Fatal("immutable-release policy probes do not use the dedicated administration credential")
	}
	if !strings.Contains(releaseWorkflow, "GH_TOKEN: ${{ secrets.GITHUB_TOKEN }}") {
		t.Fatal("ordinary release API reads no longer use the job-scoped token")
	}

	for _, caller := range []string{"on-merge.yml", "release-manual.yml"} {
		data, readErr := os.ReadFile(filepath.Join(root, ".github/workflows", caller)) //nolint:gosec // fixed repository path
		if readErr != nil {
			t.Fatal(readErr)
		}
		if strings.Count(string(data), "repository-administration-token: ${{ secrets.REPO_SYNC_TOKEN }}") != 1 {
			t.Fatalf("%s does not pass the established administration credential exactly once", caller)
		}
	}
}

func TestReusableWorkflowsPinExactTriggerSHA(t *testing.T) {
	root := testRepositoryRoot(t)
	for name, count := range map[string]int{
		"_build-test.yml":        2,
		"_generate-provider.yml": 1,
		"_generate-docs.yml":     1,
	} {
		data, err := os.ReadFile(filepath.Join(root, ".github", "workflows", name)) //nolint:gosec // fixed repository workflow
		if err != nil {
			t.Fatal(err)
		}
		workflow := string(data)
		if strings.Contains(workflow, "inputs.ref || github.ref") {
			t.Fatalf("%s still falls back to a moving branch ref", name)
		}
		if got := strings.Count(workflow, "inputs.ref || github.sha"); got != count {
			t.Fatalf("%s exact-SHA checkout count = %d, want %d", name, got, count)
		}
	}

	onMerge, err := os.ReadFile(filepath.Join(root, ".github", "workflows", "on-merge.yml")) //nolint:gosec // fixed repository workflow
	if err != nil {
		t.Fatal(err)
	}
	var caller struct {
		Jobs map[string]struct {
			Uses string         `yaml:"uses"`
			With map[string]any `yaml:"with"`
		} `yaml:"jobs"`
	}
	if err := yaml.Unmarshal(onMerge, &caller); err != nil {
		t.Fatal(err)
	}
	for job, reusable := range map[string]string{
		"build-test":          "./.github/workflows/_build-test.yml",
		"regenerate-provider": "./.github/workflows/_generate-provider.yml",
		"regenerate-docs":     "./.github/workflows/_generate-docs.yml",
	} {
		configuration := caller.Jobs[job]
		if configuration.Uses != reusable || configuration.With["ref"] != "${{ github.sha }}" {
			t.Fatalf("on-merge job %s does not invoke %s at the exact trigger SHA: %+v", job, reusable, configuration)
		}
	}
}

func TestWorkflowUsesAreImmutable(t *testing.T) {
	root := filepath.Join(testRepositoryRoot(t), ".github")
	err := filepath.Walk(root, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() || (filepath.Ext(path) != ".yml" && filepath.Ext(path) != ".yaml") {
			return nil
		}
		data, err := os.ReadFile(path) //nolint:gosec // fixed repository workflow/action tree
		if err != nil {
			return err
		}
		for lineNumber, line := range strings.Split(string(data), "\n") {
			trimmed := strings.TrimSpace(line)
			if !strings.HasPrefix(trimmed, "uses:") {
				continue
			}
			fields := strings.Fields(strings.TrimSpace(strings.TrimPrefix(trimmed, "uses:")))
			if len(fields) == 0 || strings.HasPrefix(fields[0], "./") {
				continue
			}
			at := strings.LastIndex(fields[0], "@")
			if at < 0 {
				t.Errorf("%s:%d remote use has no ref: %s", path, lineNumber+1, fields[0])
				continue
			}
			ref := fields[0][at+1:]
			if len(ref) != 40 || strings.Trim(ref, "0123456789abcdef") != "" {
				t.Errorf("%s:%d remote use is not pinned to a lowercase commit: %s", path, lineNumber+1, fields[0])
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAutoMergeRequiresCanonicalToken(t *testing.T) {
	root := testRepositoryRoot(t)
	workflowPath := filepath.Join(root, ".github", "workflows", "auto-merge.yml")
	workflowBytes, err := os.ReadFile(workflowPath) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	workflowText := string(workflowBytes)
	for _, forbidden := range []string{"falls back", "GITHUB_TOKEN when"} {
		if strings.Contains(workflowText, forbidden) {
			t.Fatalf("auto-merge retains an alternate authentication path %q", forbidden)
		}
	}
	if !strings.Contains(workflowText, "needs: [require-token]") ||
		strings.Count(workflowText, "github.event.pull_request.head.repo.full_name == github.repository") != 2 {
		t.Fatal("auto-merge call is not gated by same-repository token preflight")
	}

	script := extractWorkflowRunStep(t, "auto-merge.yml", "require-token", "Require downstream-triggering token")
	result, runErr := runWorkflowScript(t.TempDir(), script, nil)
	if runErr == nil || !strings.Contains(result, "REPO_SYNC_TOKEN is required") {
		t.Fatalf("missing auto-merge PAT was accepted: err=%v\n%s", runErr, result)
	}
	result, runErr = runWorkflowScript(t.TempDir(), script, []string{"REPO_SYNC_TOKEN=fixture"})
	if runErr != nil {
		t.Fatalf("present auto-merge PAT was rejected: %v\n%s", runErr, result)
	}
}

func TestEveryReleaseEntrypointPublishes(t *testing.T) {
	root := testRepositoryRoot(t)
	for _, name := range []string{"_tag-release.yml", "release-manual.yml", "on-merge.yml"} {
		data, err := os.ReadFile(filepath.Join(root, ".github", "workflows", name)) //nolint:gosec // fixed repository path
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(string(data), "create-release") {
			t.Fatalf("%s retains a tag-only release path", name)
		}
	}

	data, err := os.ReadFile(filepath.Join(root, ".github", "workflows", "_tag-release.yml")) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	var reusable struct {
		Jobs map[string]struct {
			If string `yaml:"if"`
		} `yaml:"jobs"`
	}
	if err := yaml.Unmarshal(data, &reusable); err != nil {
		t.Fatal(err)
	}
	if reusable.Jobs["publish"].If != "" {
		t.Fatalf("release publication remains conditional: %s", reusable.Jobs["publish"].If)
	}
}

func TestReleaseVersionRecognizesBreakingCommitForms(t *testing.T) {
	script := extractWorkflowRunStep(t, "_tag-release.yml", "tag", "Calculate next version")
	for _, tc := range []struct {
		name    string
		subject string
		body    string
		want    string
	}{
		{name: "scoped breaking subject", subject: "feat(api)!: remove legacy field", want: "v2.0.0"},
		{name: "space breaking trailer", subject: "fix: simplify parser", body: "BREAKING CHANGE: legacy syntax is removed", want: "v2.0.0"},
		{name: "hyphen breaking trailer", subject: "fix: simplify parser", body: "BREAKING-CHANGE: legacy syntax is removed", want: "v2.0.0"},
		{name: "ordinary feature", subject: "feat(api): add exact schema", want: "v1.3.0"},
		{name: "nonbreaking prose", subject: "fix: simplify parser", body: "This is not a BREAKING-CHANGE: trailer", want: "v1.2.4"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			repo := t.TempDir()
			runReleaseTestCommand(t, repo, nil, "git", "init", "-q")
			runReleaseTestCommand(t, repo, nil, "git", "config", "user.name", "Version Test")
			runReleaseTestCommand(t, repo, nil, "git", "config", "user.email", "version@example.com")
			writeReleaseTestFile(t, repo, "version.txt", "base\n", 0o600)
			runReleaseTestCommand(t, repo, nil, "git", "add", "version.txt")
			runReleaseTestCommand(t, repo, nil, "git", "commit", "-qm", "chore: baseline")
			runReleaseTestCommand(t, repo, nil, "git", "tag", "v1.2.3")
			writeReleaseTestFile(t, repo, "version.txt", "changed\n", 0o600)
			runReleaseTestCommand(t, repo, nil, "git", "add", "version.txt")
			args := []string{"commit", "-qm", tc.subject}
			if tc.body != "" {
				args = append(args, "-m", tc.body)
			}
			runReleaseTestCommand(t, repo, nil, "git", args...)
			output := filepath.Join(repo, "github-output")
			result, runErr := runWorkflowScript(repo, script, []string{"GITHUB_OUTPUT=" + output})
			if runErr != nil {
				t.Fatalf("version calculation failed: %v\n%s", runErr, result)
			}
			outputs, err := os.ReadFile(output) //nolint:gosec // isolated test fixture
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(outputs), "new_tag="+tc.want) {
				t.Fatalf("version calculation did not produce %s:\n%s", tc.want, outputs)
			}
		})
	}
}

func TestReleasePublicationRechecksRemoteTagIdentity(t *testing.T) {
	script := extractWorkflowRunStep(t, "_tag-release.yml", "publish", "Publish verified draft")
	tmp := t.TempDir()
	bin := filepath.Join(tmp, "bin")
	logPath := filepath.Join(tmp, "gh.log")
	stub := `#!/usr/bin/env bash
set -euo pipefail
printf '%s\n' "$*" >> "$GH_LOG"
if [[ "$*" == *"/commits/"* ]]; then
  printf '%s\n' "$MOVED_COMMIT"
  exit 0
fi
exit 0
`
	writeReleaseTestFile(t, bin, "gh", stub, 0o700)
	expected := strings.Repeat("a", 40)
	moved := strings.Repeat("b", 40)
	result, runErr := runWorkflowScript(tmp, script, []string{
		"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"),
		"GH_LOG=" + logPath,
		"MOVED_COMMIT=" + moved,
		"GH_TOKEN=fixture",
		"RELEASE_TAG=v1.2.3",
		"RELEASE_COMMIT=" + expected,
		"REPOSITORY=" + dispatchTarget,
	})
	if runErr == nil || !strings.Contains(result, "tag moved before publication") {
		t.Fatalf("moved remote tag did not abort publication: err=%v\n%s", runErr, result)
	}
	logBytes, err := os.ReadFile(logPath) //nolint:gosec // isolated test fixture
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(logBytes), "release edit") {
		t.Fatalf("draft was published after remote identity mismatch:\n%s", logBytes)
	}

	tampered := t.TempDir()
	tamperedBin := filepath.Join(tampered, "bin")
	tamperedLog := filepath.Join(tampered, "gh.log")
	writeReleaseTestFile(t, tampered, "release-notes.md", "notes\n", 0o600)
	writeReleaseTestFile(t, tampered, "provider-receipt.json", "{}\n", 0o600)
	writeReleaseTestJSON(t, filepath.Join(tampered, "release.json"), map[string]any{
		"body": "notes\n", "tag_name": "v1.2.3", "draft": true, "prerelease": false,
	})
	writeReleaseTestFile(t, tamperedBin, "gh", `#!/usr/bin/env bash
set -euo pipefail
printf '%s\n' "$*" >> "$GH_LOG"
if [[ "$*" == *"/commits/"* ]]; then
  printf '%s\n' "$EXPECTED_COMMIT"
elif [ "$1" = api ]; then
  jq -c -s '.' "$RELEASE_JSON"
elif [ "$1 $2" = "release download" ]; then
  while [ "$#" -gt 0 ]; do
    if [ "$1" = --dir ]; then
      printf 'tampered\n' > "$2/asset"
      exit 0
    fi
    shift
  done
  exit 2
fi
`, 0o700)
	writeReleaseTestFile(t, tampered, "scripts/verify-provider-release.sh", `#!/usr/bin/env bash
set -euo pipefail
grep -qx expected "$4/asset"
`, 0o700)
	result, runErr = runWorkflowScript(tampered, script, []string{
		"PATH=" + tamperedBin + string(os.PathListSeparator) + os.Getenv("PATH"),
		"GH_LOG=" + tamperedLog,
		"EXPECTED_COMMIT=" + expected,
		"RELEASE_JSON=" + filepath.Join(tampered, "release.json"),
		"RUNNER_TEMP=" + tampered,
		"GH_TOKEN=fixture",
		"RELEASE_TAG=v1.2.3",
		"RELEASE_COMMIT=" + expected,
		"REPOSITORY=" + dispatchTarget,
	})
	if runErr == nil || !strings.Contains(result, "assets changed before publication") {
		t.Fatalf("mutated draft asset did not abort publication: err=%v\n%s", runErr, result)
	}
	logBytes, err = os.ReadFile(tamperedLog) //nolint:gosec // isolated fixture
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(logBytes), "--draft=false") {
		t.Fatalf("draft was published after asset mutation:\n%s", logBytes)
	}

	for _, step := range []string{"Verify already-published release without mutation", "Verify sealed publication"} {
		body := extractWorkflowRunStep(t, "_tag-release.yml", "publish", step)
		if !strings.Contains(body, `repos/${REPOSITORY}/commits/${RELEASE_TAG}`) {
			t.Fatalf("%s does not re-resolve the remote tag", step)
		}
	}
}

func TestRegenerationBuildUsesBoundedMemory(t *testing.T) {
	workflowPath := filepath.Join(testRepositoryRoot(t), ".github", "workflows", "_generate-provider.yml")
	workflowBytes, err := os.ReadFile(workflowPath) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	var workflow struct {
		Jobs map[string]struct {
			Steps []struct {
				Name string            `yaml:"name"`
				Run  string            `yaml:"run"`
				Env  map[string]string `yaml:"env"`
			} `yaml:"steps"`
		} `yaml:"jobs"`
	}
	if err := yaml.Unmarshal(workflowBytes, &workflow); err != nil {
		t.Fatal(err)
	}
	for _, step := range workflow.Jobs["generate"].Steps {
		if step.Name != "Build to verify" {
			continue
		}
		for key, want := range map[string]string{
			"GOGC":       "10",
			"GOMEMLIMIT": "4GiB",
			"GOMAXPROCS": "1",
		} {
			if got := step.Env[key]; got != want {
				t.Fatalf("regeneration build %s = %q, want %q", key, got, want)
			}
		}
		for _, want := range []string{"go build -p 1", "-gcflags='all=-N -l'", "-v ./..."} {
			if !strings.Contains(step.Run, want) {
				t.Fatalf("regeneration build does not use %q: %s", want, step.Run)
			}
		}
		return
	}
	t.Fatal("regeneration build step not found")
}

func TestScheduledAcceptanceFailureFailsWorkflow(t *testing.T) {
	workflowPath := filepath.Join(testRepositoryRoot(t), ".github", "workflows", "acc-tests.yml")
	workflowBytes, err := os.ReadFile(workflowPath) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	var workflow struct {
		Jobs map[string]struct {
			ContinueOnError any    `yaml:"continue-on-error"`
			If              string `yaml:"if"`
			Steps           []struct {
				Name            string            `yaml:"name"`
				Env             map[string]string `yaml:"env"`
				Run             string            `yaml:"run"`
				Uses            string            `yaml:"uses"`
				With            map[string]any    `yaml:"with"`
				ContinueOnError any               `yaml:"continue-on-error"`
				If              string            `yaml:"if"`
			} `yaml:"steps"`
		} `yaml:"jobs"`
	}
	if err := yaml.Unmarshal(workflowBytes, &workflow); err != nil {
		t.Fatal(err)
	}
	if workflow.Jobs["real-api-tests"].ContinueOnError != nil {
		t.Fatal("scheduled real API job still suppresses its failure with continue-on-error")
	}
	if !strings.Contains(workflow.Jobs["real-api-tests"].If, "always()") {
		t.Fatal("real-only mode cannot run after its intentionally skipped mock prerequisite")
	}
	for _, step := range workflow.Jobs["summary"].Steps {
		switch step.Name {
		case "Download all reports":
			if step.ContinueOnError != nil {
				t.Fatal("acceptance summary suppresses missing evidence download")
			}
		case "Notify on scheduled failure":
			if strings.Count(step.If, "!= 'success'") != 4 {
				t.Fatalf("scheduled notification does not cover every non-success result: %s", step.If)
			}
		}
	}
	workflowText := string(workflowBytes)
	mockScript := extractWorkflowRunStep(t, "acc-tests.yml", "mock-tests", "Run mock tests")
	if !strings.Contains(mockScript, "go test -json \\\n  -p 1 \\") {
		t.Fatal("mock acceptance tests do not serialize package builds for the constrained runner")
	}
	if !strings.Contains(mockScript, "-gcflags='all=-N -l'") {
		t.Fatal("mock acceptance tests do not disable memory-heavy compiler optimization")
	}
	mockStepFound := false
	for _, step := range workflow.Jobs["mock-tests"].Steps {
		if step.Name != "Run mock tests" {
			continue
		}
		mockStepFound = true
		for key, want := range map[string]string{
			"GOGC":       "10",
			"GOMEMLIMIT": "4GiB",
			"GOMAXPROCS": "1",
		} {
			if got := step.Env[key]; got != want {
				t.Fatalf("mock acceptance %s = %q, want %q", key, got, want)
			}
		}
	}
	if !mockStepFound {
		t.Fatal("mock acceptance test step not found")
	}
	for _, forbidden := range []string{"P12", "p12", "GO_VERSION", "go-version:", "RUNNER_NAME", "batch_delay"} {
		if strings.Contains(workflowText, forbidden) {
			t.Fatalf("acceptance workflow retains pre-production compatibility or moving-toolchain marker %q", forbidden)
		}
	}
	modeExpression := "github.event.inputs.mode || (github.event_name == 'pull_request' && 'pr-subset') || 'full'"
	if strings.Count(workflowText, modeExpression) != 2 {
		t.Fatalf("acceptance mode is not reported and gated with the same event-aware expression")
	}
	setupGoCount := 0
	verifiedImageGo := map[string]bool{}
	for jobName, job := range workflow.Jobs {
		for _, step := range job.Steps {
			if strings.HasPrefix(step.Uses, "actions/setup-go@") {
				setupGoCount++
			}
			if step.Name == "Verify immutable Go toolchain" &&
				strings.Contains(step.Run, `test "$(go env GOVERSION)" = go1.25.12`) {
				verifiedImageGo[jobName] = true
			}
		}
	}
	if setupGoCount != 0 {
		t.Fatalf("acceptance workflow downloads Go in %d managed-socketless jobs", setupGoCount)
	}
	for _, jobName := range []string{"mock-tests", "real-api-tests", "cleanup", "compare-results"} {
		if !verifiedImageGo[jobName] {
			t.Errorf("acceptance job %s does not verify the immutable Go 1.25.12 image toolchain", jobName)
		}
	}
	for _, line := range strings.Split(workflowText, "\n") {
		trimmed := strings.TrimSpace(line)
		if (strings.HasPrefix(trimmed, "echo ") || strings.HasPrefix(trimmed, "printf ")) &&
			(strings.Contains(trimmed, "$XCSH_API_URL") || strings.Contains(trimmed, "$XCSH_API_TOKEN")) {
			t.Fatalf("acceptance workflow prints credential-derived infrastructure data: %s", trimmed)
		}
	}

	for _, tc := range []struct {
		job, step, want string
	}{
		{job: "mock-tests", step: "Run mock tests", want: "Mock tests failed"},
		{job: "real-api-tests", step: "Run real API tests", want: "Real API tests failed"},
	} {
		t.Run(tc.job, func(t *testing.T) {
			tmp := t.TempDir()
			bin := filepath.Join(tmp, "bin")
			writeReleaseTestFile(t, bin, "go", `#!/usr/bin/env bash
set -euo pipefail
if [ "$1" = test ]; then
  printf '%s\n' '{"Action":"fail","Package":"fixture.invalid/failing"}'
  exit 7
fi
format=
output=
for arg in "$@"; do
  case "$arg" in
    -format=*) format=${arg#-format=} ;;
    -output=*) output=${arg#-output=} ;;
  esac
done
[ -n "$output" ] || exit 2
if [ "$format" = json ]; then
  printf '%s\n' '{"total_passed":0,"total_failed":1,"transient_errors":[],"failed_tests":[{"failure_output":"tenant identifier"}]}' > "$output"
else
  printf 'report\n' > "$output"
fi
exit 1
`, 0o700)
			outputPath := filepath.Join(tmp, "github-output")
			script := extractWorkflowRunStep(t, "acc-tests.yml", tc.job, tc.step)
			if !strings.Contains(script, "set -o pipefail") {
				t.Fatal("test step does not enable pipefail before piping go test through tee")
			}
			result, runErr := runImplicitWorkflowScript(tmp, script, []string{
				"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"),
				"GITHUB_OUTPUT=" + outputPath,
				"INPUT_PARALLEL=1",
				"INPUT_TIMEOUT=1",
				"AUTH_METHOD=token",
			})
			if runErr == nil || !strings.Contains(result, tc.want) {
				t.Fatalf("failing go test was reported as success: err=%v\n%s", runErr, result)
			}
			outputs, err := os.ReadFile(outputPath) //nolint:gosec // isolated fixture
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(outputs), "failed=1") {
				t.Fatalf("test reports were not preserved before failure:\n%s", outputs)
			}
			if tc.job == "real-api-tests" {
				report, err := os.ReadFile(filepath.Join(tmp, "test-reports", "real-tests.json")) //nolint:gosec // isolated fixture
				if err != nil {
					t.Fatal(err)
				}
				if strings.Contains(string(report), "failure_output") || strings.Contains(string(report), "tenant identifier") {
					t.Fatalf("real API report retained failure output:\n%s", report)
				}
				if _, err := os.Stat(filepath.Join(tmp, "test-reports", "test-output-real.json")); !os.IsNotExist(err) {
					t.Fatalf("raw live-tenant output survived report generation: %v", err)
				}
			}
		})
	}

	t.Run("sweeper failure", func(t *testing.T) {
		tmp := t.TempDir()
		bin := filepath.Join(tmp, "bin")
		writeReleaseTestFile(t, bin, "go", "#!/usr/bin/env bash\nprintf 'sweeper detail that must stay local\\n'\nexit 7\n", 0o700)
		outputPath := filepath.Join(tmp, "github-output")
		sweepScript := extractWorkflowRunStep(t, "acc-tests.yml", "cleanup", "Run sweepers")
		result, runErr := runImplicitWorkflowScript(tmp, sweepScript, []string{
			"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"),
			"GITHUB_OUTPUT=" + outputPath,
			"XCSH_API_URL=https://example.invalid",
			"XCSH_API_TOKEN=fixture",
		})
		if runErr == nil || !strings.Contains(result, "sweepers failed") {
			t.Fatalf("sweeper failure was reported as success: err=%v\n%s", runErr, result)
		}
		if strings.Contains(result, "sweeper detail") {
			t.Fatalf("sweeper output leaked to the job log:\n%s", result)
		}
	})

	t.Run("mock-real comparison", func(t *testing.T) {
		tmp := t.TempDir()
		for _, report := range []string{"mock-reports/mock-tests.json", "real-reports/real-tests.json"} {
			writeReleaseTestFile(t, tmp, report, "{}\n", 0o600)
		}
		bin := filepath.Join(tmp, "bin")
		writeReleaseTestFile(t, bin, "go", `#!/usr/bin/env bash
set -euo pipefail
output=
format=
for arg in "$@"; do
  case "$arg" in
    -format=*) format=${arg#-format=} ;;
    -output=*) output=${arg#-output=} ;;
  esac
done
if [ "$format" = json ]; then
  printf '%s\n' '{"summary":{"consistent":0,"mock_needs_fix":1}}' > "$output"
else
  printf 'comparison mismatch\n' > "$output"
fi
exit 7
`, 0o700)
		outputPath := filepath.Join(tmp, "github-output")
		comparisonScript := extractWorkflowRunStep(t, "acc-tests.yml", "compare-results", "Run comparison")
		result, runErr := runImplicitWorkflowScript(tmp, comparisonScript, []string{
			"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"),
			"GITHUB_OUTPUT=" + outputPath,
		})
		if runErr == nil || !strings.Contains(result, "Mock server has discrepancies") {
			t.Fatalf("mock/real discrepancy was reported as success: err=%v\n%s", runErr, result)
		}
		outputs, err := os.ReadFile(outputPath) //nolint:gosec // isolated fixture
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(outputs), "available=true") ||
			!strings.Contains(string(outputs), "needs_fix=1") {
			t.Fatalf("comparison evidence was not preserved before failure:\n%s", outputs)
		}
	})

	t.Run("mock-real comparison requires both reports", func(t *testing.T) {
		tmp := t.TempDir()
		outputPath := filepath.Join(tmp, "github-output")
		comparisonScript := extractWorkflowRunStep(t, "acc-tests.yml", "compare-results", "Run comparison")
		result, runErr := runImplicitWorkflowScript(tmp, comparisonScript, []string{
			"GITHUB_OUTPUT=" + outputPath,
		})
		if runErr == nil || !strings.Contains(result, "Mock test JSON evidence is missing") {
			t.Fatalf("missing comparison input was reported as success: err=%v\n%s", runErr, result)
		}
	})

	script := extractWorkflowRunStep(t, "acc-tests.yml", "summary", "Check overall status")
	writeAllAcceptanceEvidence := func(t *testing.T, root string) {
		t.Helper()
		for name, body := range map[string]string{
			"all-reports/mock-test-reports/mock-tests.json":     `{"total_passed":1,"total_failed":0}` + "\n",
			"all-reports/mock-test-reports/mock-tests.xml":      `<testsuite tests="1" failures="0"></testsuite>` + "\n",
			"all-reports/mock-test-reports/mock-tests.md":       "mock evidence\n",
			"all-reports/real-api-test-reports/real-tests.json": `{"total_passed":1,"total_failed":0}` + "\n",
			"all-reports/real-api-test-reports/real-tests.md":   "real evidence\n",
			"all-reports/comparison-reports/comparison.json":    `{"summary":{"consistent":1,"mock_needs_fix":0}}` + "\n",
			"all-reports/comparison-reports/comparison.md":      "comparison evidence\n",
		} {
			writeReleaseTestFile(t, root, name, body, 0o600)
		}
	}

	for _, tc := range []struct {
		name, event, mode, mock, real, comparison, cleanup string
	}{
		{name: "pull request", event: "pull_request", mode: "pr-subset", mock: "success", real: "skipped", comparison: "skipped", cleanup: "skipped"},
		{name: "manual mock only", event: "workflow_dispatch", mode: "mock-only", mock: "success", real: "skipped", comparison: "skipped", cleanup: "skipped"},
		{name: "manual real only", event: "workflow_dispatch", mode: "real-only", mock: "skipped", real: "success", comparison: "skipped", cleanup: "success"},
		{name: "manual full", event: "workflow_dispatch", mode: "full", mock: "success", real: "success", comparison: "success", cleanup: "success"},
		{name: "scheduled full", event: "schedule", mode: "full", mock: "success", real: "success", comparison: "success", cleanup: "success"},
	} {
		t.Run("result matrix accepts "+tc.name, func(t *testing.T) {
			tmp := t.TempDir()
			writeAllAcceptanceEvidence(t, tmp)
			result, runErr := runWorkflowScript(tmp, script, []string{
				"EVENT_NAME=" + tc.event,
				"TEST_MODE=" + tc.mode,
				"MOCK_RESULT=" + tc.mock,
				"REAL_RESULT=" + tc.real,
				"COMPARE_RESULT=" + tc.comparison,
				"CLEANUP_RESULT=" + tc.cleanup,
			})
			if runErr != nil || !strings.Contains(result, "All required tests passed") {
				t.Fatalf("valid acceptance result matrix was rejected: err=%v\n%s", runErr, result)
			}
		})
	}

	for _, tc := range []struct {
		name, event, mode, mock, real, comparison, cleanup, want string
	}{
		{name: "cancelled required mock", event: "pull_request", mode: "pr-subset", mock: "cancelled", real: "skipped", comparison: "skipped", cleanup: "skipped", want: "Mock tests result was cancelled"},
		{name: "all skipped", event: "pull_request", mode: "pr-subset", mock: "skipped", real: "skipped", comparison: "skipped", cleanup: "skipped", want: "Mock tests result was skipped"},
		{name: "scheduled real failure", event: "schedule", mode: "full", mock: "success", real: "failure", comparison: "skipped", cleanup: "success", want: "Real API tests result was failure"},
		{name: "comparison failure", event: "workflow_dispatch", mode: "full", mock: "success", real: "success", comparison: "failure", cleanup: "success", want: "Mock/real comparison result was failure"},
		{name: "cleanup failure", event: "workflow_dispatch", mode: "full", mock: "success", real: "success", comparison: "success", cleanup: "failure", want: "Resource cleanup result was failure"},
		{name: "unknown mode", event: "workflow_dispatch", mode: "compatibility", mock: "success", real: "skipped", comparison: "skipped", cleanup: "skipped", want: "Unsupported acceptance-test event/mode"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tmp := t.TempDir()
			writeAllAcceptanceEvidence(t, tmp)
			result, runErr := runWorkflowScript(tmp, script, []string{
				"EVENT_NAME=" + tc.event,
				"TEST_MODE=" + tc.mode,
				"MOCK_RESULT=" + tc.mock,
				"REAL_RESULT=" + tc.real,
				"COMPARE_RESULT=" + tc.comparison,
				"CLEANUP_RESULT=" + tc.cleanup,
			})
			if runErr == nil || !strings.Contains(result, tc.want) {
				t.Fatalf("acceptance failure was reported as success: err=%v\n%s", runErr, result)
			}
		})
	}

	t.Run("successful jobs require captured evidence", func(t *testing.T) {
		result, runErr := runWorkflowScript(t.TempDir(), script, []string{
			"EVENT_NAME=pull_request",
			"TEST_MODE=pr-subset",
			"MOCK_RESULT=success",
			"REAL_RESULT=skipped",
			"COMPARE_RESULT=skipped",
			"CLEANUP_RESULT=skipped",
		})
		if runErr == nil || !strings.Contains(result, "Required acceptance-test evidence is missing") {
			t.Fatalf("missing acceptance evidence was reported as success: err=%v\n%s", runErr, result)
		}
	})

	t.Run("successful jobs require parseable evidence", func(t *testing.T) {
		tmp := t.TempDir()
		writeAllAcceptanceEvidence(t, tmp)
		writeReleaseTestFile(t, tmp, "all-reports/mock-test-reports/mock-tests.json", "not json\n", 0o600)
		result, runErr := runWorkflowScript(tmp, script, []string{
			"EVENT_NAME=pull_request",
			"TEST_MODE=pr-subset",
			"MOCK_RESULT=success",
			"REAL_RESULT=skipped",
			"COMPARE_RESULT=skipped",
			"CLEANUP_RESULT=skipped",
		})
		if runErr == nil || !strings.Contains(result, "Mock test JSON evidence is malformed") {
			t.Fatalf("malformed acceptance evidence was reported as success: err=%v\n%s", runErr, result)
		}
	})
}

func TestReleaseReadyStateRejectsDescendantOfRegeneration(t *testing.T) {
	root := testRepositoryRoot(t)
	fixture := newReceiptFixture(t, false, false)
	writeReleaseTestFile(t, fixture.repo, "governance.txt", "later commit\n", 0o600)
	runReleaseTestCommand(t, fixture.repo, nil, "git", "add", "governance.txt")
	runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "governance after regeneration")

	validator := filepath.Join(root, "scripts", "validate-provider-delivery-state.sh")
	cmd := exec.Command(validator, "--release-ready")
	cmd.Dir = fixture.repo
	cmd.Env = append(os.Environ(), "TARGET_REPOSITORY="+dispatchTarget)
	output, runErr := cmd.CombinedOutput()
	if runErr == nil || !strings.Contains(string(output), "release commit did not introduce") {
		t.Fatalf("descendant release commit was accepted: err=%v\n%s", runErr, output)
	}
}

func TestReleaseReadyStateAcceptsExactRecoveryRebind(t *testing.T) {
	root := testRepositoryRoot(t)
	fixture := newReceiptFixture(t, false, false)
	rebindReceiptFixture(t, fixture, "", "")

	validator := filepath.Join(root, "scripts", "validate-provider-delivery-state.sh")
	runReleaseTestCommand(t, fixture.repo, []string{"TARGET_REPOSITORY=" + dispatchTarget}, validator, "--release-ready")
}

func TestReleaseReadyStateRejectsInvalidRecoveryRebind(t *testing.T) {
	root := testRepositoryRoot(t)
	validator := filepath.Join(root, "scripts", "validate-provider-delivery-state.sh")
	for _, tc := range []struct {
		name, previousVersion, sourceCommit, want string
	}{
		{
			name:            "identity changed",
			previousVersion: "2.1.207",
			want:            "may change only its source commit",
		},
		{
			name:         "source is not exact parent",
			sourceCommit: strings.Repeat("a", 40),
			want:         "source is not the exact release parent",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fixture := newReceiptFixture(t, false, false)
			rebindReceiptFixture(t, fixture, tc.previousVersion, tc.sourceCommit)
			cmd := exec.Command(validator, "--release-ready")
			cmd.Dir = fixture.repo
			cmd.Env = append(os.Environ(), "TARGET_REPOSITORY="+dispatchTarget)
			output, runErr := cmd.CombinedOutput()
			if runErr == nil || !strings.Contains(string(output), tc.want) {
				t.Fatalf("invalid recovery rebind was accepted or failed for the wrong reason: err=%v\n%s", runErr, output)
			}
		})
	}
}

func TestDeliveryStateValidatorPermitsOnlyInheritedStaleReceiptForPRValidation(t *testing.T) {
	root := testRepositoryRoot(t)
	validator := filepath.Join(root, "scripts", "validate-provider-delivery-state.sh")

	t.Run("unchanged inherited receipt is accepted only with a base reference", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		staleSource := createNonAncestorReceiptSource(t, fixture.repo, "stale-source")
		setRegenerationReceiptSource(t, fixture.repo, staleSource)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", "tools/spec-regeneration-receipt.json")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "persist inherited stale receipt")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "tag", "delivery-base")
		writeReleaseTestFile(t, fixture.repo, "governance.txt", "PR-only change\n", 0o600)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", "governance.txt")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "governance change")

		runReleaseTestCommand(t, fixture.repo, []string{"TARGET_REPOSITORY=" + dispatchTarget}, validator, "--base-ref", "delivery-base")

		cmd := exec.Command(validator)
		cmd.Dir = fixture.repo
		cmd.Env = append(os.Environ(), "TARGET_REPOSITORY="+dispatchTarget)
		output, runErr := cmd.CombinedOutput()
		if runErr == nil || !strings.Contains(string(output), "source is not an ancestor") {
			t.Fatalf("stale receipt without a base reference was accepted: err=%v\n%s", runErr, output)
		}

		cmd = exec.Command(validator, "--release-ready")
		cmd.Dir = fixture.repo
		cmd.Env = append(os.Environ(), "TARGET_REPOSITORY="+dispatchTarget)
		output, runErr = cmd.CombinedOutput()
		if runErr == nil || !strings.Contains(string(output), "release commit did not introduce") {
			t.Fatalf("release-ready validation accepted an inherited stale receipt: err=%v\n%s", runErr, output)
		}

		secondStaleSource := createNonAncestorReceiptSource(t, fixture.repo, "second-stale-source")
		setRegenerationReceiptSource(t, fixture.repo, secondStaleSource)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", "tools/spec-regeneration-receipt.json")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "change stale receipt")
		cmd = exec.Command(validator, "--base-ref", "delivery-base")
		cmd.Dir = fixture.repo
		cmd.Env = append(os.Environ(), "TARGET_REPOSITORY="+dispatchTarget)
		output, runErr = cmd.CombinedOutput()
		if runErr == nil || !strings.Contains(string(output), "source is not an ancestor") {
			t.Fatalf("changed stale receipt was accepted: err=%v\n%s", runErr, output)
		}
	})

	t.Run("new stale receipt is rejected even with a base reference", func(t *testing.T) {
		fixture := newReceiptFixture(t, false, false)
		regenerationPath := filepath.Join(fixture.repo, "tools/spec-regeneration-receipt.json")
		regeneration, err := os.ReadFile(regenerationPath) //nolint:gosec // isolated fixture
		if err != nil {
			t.Fatal(err)
		}
		runReleaseTestCommand(t, fixture.repo, nil, "git", "rm", "-q", "tools/spec-regeneration-receipt.json")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "remove stale receipt")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "tag", "delivery-base")
		staleSource := createNonAncestorReceiptSource(t, fixture.repo, "new-stale-source")
		if err := os.WriteFile(regenerationPath, regeneration, 0o600); err != nil {
			t.Fatal(err)
		}
		setRegenerationReceiptSource(t, fixture.repo, staleSource)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", "tools/spec-regeneration-receipt.json")
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "add stale receipt")

		cmd := exec.Command(validator, "--base-ref", "delivery-base")
		cmd.Dir = fixture.repo
		cmd.Env = append(os.Environ(), "TARGET_REPOSITORY="+dispatchTarget)
		output, runErr := cmd.CombinedOutput()
		if runErr == nil || !strings.Contains(string(output), "source is not an ancestor") {
			t.Fatalf("new stale receipt was accepted: err=%v\n%s", runErr, output)
		}
	})
}

func createNonAncestorReceiptSource(t *testing.T, repo, branch string) string {
	t.Helper()
	runReleaseTestCommand(t, repo, nil, "git", "checkout", "-q", "-b", branch)
	writeReleaseTestFile(t, repo, branch+".txt", "stale source\n", 0o600)
	runReleaseTestCommand(t, repo, nil, "git", "add", branch+".txt")
	runReleaseTestCommand(t, repo, nil, "git", "commit", "-qm", branch)
	source := strings.TrimSpace(runReleaseTestCommand(t, repo, nil, "git", "rev-parse", "HEAD"))
	runReleaseTestCommand(t, repo, nil, "git", "checkout", "-q", "main")
	return source
}

func setRegenerationReceiptSource(t *testing.T, repo, source string) {
	t.Helper()
	path := filepath.Join(repo, "tools/spec-regeneration-receipt.json")
	data, err := os.ReadFile(path) //nolint:gosec // isolated fixture
	if err != nil {
		t.Fatal(err)
	}
	var receipt map[string]any
	if err := json.Unmarshal(data, &receipt); err != nil {
		t.Fatal(err)
	}
	receipt["source_commit"] = source
	writeReleaseTestJSON(t, path, receipt)
}
func TestDraftCleanupDeletesEveryMeasuredAsset(t *testing.T) {
	script := extractWorkflowRunStep(t, "_tag-release.yml", "publish", "Clear repairable draft artifacts")
	tmp := t.TempDir()
	writeReleaseTestFile(t, tmp, "release-initial.json", `{"assets":[{"id":11},{"id":22}]}`+"\n", 0o600)
	bin := filepath.Join(tmp, "bin")
	logPath := filepath.Join(tmp, "deletions")
	stub := "#!/usr/bin/env bash\nprintf '%s\\n' \"$*\" >> \"$DELETE_LOG\"\n"
	writeReleaseTestFile(t, bin, "gh", stub, 0o700)
	result, runErr := runWorkflowScript(tmp, script, []string{
		"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"),
		"RUNNER_TEMP=" + tmp,
		"REPOSITORY=f5-sales-demo/terraform-provider-xcsh",
		"DELETE_LOG=" + logPath,
		"GH_TOKEN=fixture",
	})
	if runErr != nil {
		t.Fatalf("draft cleanup failed: %v\n%s", runErr, result)
	}
	deletions, err := os.ReadFile(logPath) //nolint:gosec // isolated fixture
	if err != nil {
		t.Fatal(err)
	}
	want := "api --method DELETE repos/f5-sales-demo/terraform-provider-xcsh/releases/assets/11\n" +
		"api --method DELETE repos/f5-sales-demo/terraform-provider-xcsh/releases/assets/22\n"
	if string(deletions) != want {
		t.Fatalf("draft cleanup did not delete the exact measured set:\nwant:\n%s\ngot:\n%s", want, deletions)
	}
}

func TestPrepareReceiptRejectsDuplicateMarkerIncludingEmptyMarker(t *testing.T) {
	fixture := newReceiptFixture(t, false, false)
	data, err := os.ReadFile(filepath.Join(fixture.repo, "release.json")) //nolint:gosec // isolated fixture
	if err != nil {
		t.Fatal(err)
	}
	var release map[string]any
	if err := json.Unmarshal(data, &release); err != nil {
		t.Fatal(err)
	}
	release["body"] = release["body"].(string) + "<!-- provider-publication-receipt: -->\n"
	writeReleaseTestJSON(t, filepath.Join(fixture.repo, "release.json"), release)
	cmd := exec.Command(fixture.script)
	cmd.Dir = fixture.repo
	cmd.Env = append(os.Environ(), fixture.env...)
	output, err := cmd.CombinedOutput()
	if err == nil || !strings.Contains(string(output), "exactly one publication receipt") {
		t.Fatalf("empty duplicate receipt marker was accepted: err=%v\n%s", err, output)
	}
}

func TestDeliveryStateValidatorRejectsCrossBindingAndDeletion(t *testing.T) {
	root := testRepositoryRoot(t)
	fixture := newReceiptFixture(t, false, false)
	runReleaseTestCommand(t, fixture.repo, fixture.env, fixture.script)
	validator := filepath.Join(root, "scripts", "validate-provider-delivery-state.sh")
	runReleaseTestCommand(t, fixture.repo, []string{"TARGET_REPOSITORY=" + dispatchTarget}, validator)

	commonPath := filepath.Join(fixture.repo, "tools/spec-deliveries.json")
	commonBytes, err := os.ReadFile(commonPath) //nolint:gosec // isolated fixture
	if err != nil {
		t.Fatal(err)
	}
	var common map[string]any
	if err := json.Unmarshal(commonBytes, &common); err != nil {
		t.Fatal(err)
	}
	detailedPath := filepath.Join(fixture.repo, "tools/provider-publication-receipts.json")
	detailedBytes, err := os.ReadFile(detailedPath) //nolint:gosec // isolated fixture
	if err != nil {
		t.Fatal(err)
	}
	var detailed map[string]any
	if err := json.Unmarshal(detailedBytes, &detailed); err != nil {
		t.Fatal(err)
	}
	alternateCommit := strings.Repeat("d", 40)
	alternateID := providerDeliveryID(t, "2.1.207", "v2.1.207", alternateCommit)
	alternateDelivery := map[string]any{
		"release_tag": "v2.1.207", "target_commit": alternateCommit, "version": "2.1.207",
	}
	alternateAssets := map[string]string{}
	for i, name := range providerReleaseAssetNames("9.8.6") {
		alternateAssets[name] = fmt.Sprintf("%064x", i+100)
	}
	alternatePublication := map[string]any{
		"assets": qualifiedAssetDigests(alternateAssets), "commit": strings.Repeat("e", 40),
		"spec_release_sha256": strings.Repeat("f", 64), "tag": "v9.8.6", "version": "9.8.6",
	}
	common["deliveries"].(map[string]any)[alternateID] = alternateDelivery
	detailed["receipts"].(map[string]any)[alternateID] = map[string]any{
		"delivery": alternateDelivery, "publication": alternatePublication,
	}
	writeReleaseTestJSON(t, commonPath, common)
	writeReleaseTestJSON(t, detailedPath, detailed)
	runReleaseTestCommand(t, fixture.repo, []string{"TARGET_REPOSITORY=" + dispatchTarget}, validator)
	commonBytes, err = os.ReadFile(commonPath) //nolint:gosec // isolated fixture
	if err != nil {
		t.Fatal(err)
	}
	detailedBytes, err = os.ReadFile(detailedPath) //nolint:gosec // isolated fixture
	if err != nil {
		t.Fatal(err)
	}
	receipt := detailed["receipts"].(map[string]any)[fixture.deliveryID].(map[string]any)
	receipt["delivery"].(map[string]any)["target_commit"] = strings.Repeat("c", 40)
	writeReleaseTestJSON(t, detailedPath, detailed)
	cmd := exec.Command(validator)
	cmd.Dir = fixture.repo
	cmd.Env = append(os.Environ(), "TARGET_REPOSITORY="+dispatchTarget)
	output, runErr := cmd.CombinedOutput()
	if runErr == nil || !strings.Contains(string(output), "cross-bound") {
		t.Fatalf("cross-bound publication evidence was accepted: err=%v\n%s", runErr, output)
	}

	if err := os.WriteFile(detailedPath, detailedBytes, 0o600); err != nil {
		t.Fatal(err)
	}
	// The prepared state is intentionally uncommitted; persist it as the immutable base.
	runReleaseTestCommand(t, fixture.repo, nil, "git", "add", "tools")
	runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "durable receipt")
	runReleaseTestCommand(t, fixture.repo, nil, "git", "tag", "delivery-base")
	if err := json.Unmarshal(commonBytes, &common); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(detailedBytes, &detailed); err != nil {
		t.Fatal(err)
	}
	delete(common["deliveries"].(map[string]any), alternateID)
	delete(detailed["receipts"].(map[string]any), alternateID)
	writeReleaseTestJSON(t, commonPath, common)
	writeReleaseTestJSON(t, detailedPath, detailed)
	cmd = exec.Command(validator, "--base-ref", "delivery-base")
	cmd.Dir = fixture.repo
	cmd.Env = append(os.Environ(), "TARGET_REPOSITORY="+dispatchTarget)
	output, runErr = cmd.CombinedOutput()
	if runErr == nil || !strings.Contains(string(output), "changed or deleted") {
		t.Fatalf("durable receipt deletion was accepted: err=%v\n%s", runErr, output)
	}
}

func TestDeliveryStateValidatorRejectsDuplicateIdentities(t *testing.T) {
	for _, tc := range []struct {
		name                       string
		version, tag, commit, want string
		reuseExistingPublication   bool
	}{
		{
			name: "duplicate spec tag with another commit", version: "2.1.208", tag: "v2.1.208",
			commit: strings.Repeat("d", 40), want: "durable spec deliveries reuse",
		},
		{
			name: "reused provider publication", version: "2.1.207", tag: "v2.1.207",
			commit: strings.Repeat("d", 40), want: "durable receipts reuse",
			reuseExistingPublication: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := testRepositoryRoot(t)
			fixture := newReceiptFixture(t, false, false)
			runReleaseTestCommand(t, fixture.repo, fixture.env, fixture.script)
			commonPath := filepath.Join(fixture.repo, "tools/spec-deliveries.json")
			detailedPath := filepath.Join(fixture.repo, "tools/provider-publication-receipts.json")
			commonBytes, err := os.ReadFile(commonPath) //nolint:gosec // isolated fixture
			if err != nil {
				t.Fatal(err)
			}
			detailedBytes, err := os.ReadFile(detailedPath) //nolint:gosec // isolated fixture
			if err != nil {
				t.Fatal(err)
			}
			var common, detailed map[string]any
			if err := json.Unmarshal(commonBytes, &common); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(detailedBytes, &detailed); err != nil {
				t.Fatal(err)
			}

			deliveryID := providerDeliveryID(t, tc.version, tc.tag, tc.commit)
			delivery := map[string]any{
				"release_tag": tc.tag, "target_commit": tc.commit, "version": tc.version,
			}
			publication := detailed["receipts"].(map[string]any)[fixture.deliveryID].(map[string]any)["publication"]
			if !tc.reuseExistingPublication {
				assets := map[string]string{}
				for i, name := range providerReleaseAssetNames("9.8.6") {
					assets[name] = fmt.Sprintf("%064x", i+100)
				}
				publication = map[string]any{
					"assets": assets, "commit": strings.Repeat("e", 40),
					"spec_release_sha256": strings.Repeat("f", 64), "tag": "v9.8.6", "version": "9.8.6",
				}
			}
			common["deliveries"].(map[string]any)[deliveryID] = delivery
			detailed["receipts"].(map[string]any)[deliveryID] = map[string]any{
				"delivery": delivery, "publication": publication,
			}
			writeReleaseTestJSON(t, commonPath, common)
			writeReleaseTestJSON(t, detailedPath, detailed)

			validator := filepath.Join(root, "scripts", "validate-provider-delivery-state.sh")
			cmd := exec.Command(validator)
			cmd.Dir = fixture.repo
			cmd.Env = append(os.Environ(), "TARGET_REPOSITORY="+dispatchTarget)
			output, runErr := cmd.CombinedOutput()
			if runErr == nil || !strings.Contains(string(output), tc.want) {
				t.Fatalf("duplicate durable identity was accepted: err=%v\n%s", runErr, output)
			}
		})
	}
}

type receiptFixture struct {
	repo       string
	script     string
	output     string
	deliveryID string
	regenPR    string
	env        []string
}

func rewriteRegenerationPRMergeCommit(t *testing.T, fixture *receiptFixture, commit string) {
	t.Helper()
	data, err := os.ReadFile(fixture.regenPR) //nolint:gosec // isolated test fixture
	if err != nil {
		t.Fatal(err)
	}
	var pulls []map[string]any
	if err := json.Unmarshal(data, &pulls); err != nil {
		t.Fatal(err)
	}
	if len(pulls) != 1 {
		t.Fatalf("expected one regeneration PR fixture, found %d", len(pulls))
	}
	pulls[0]["merge_commit_sha"] = commit
	writeReleaseTestJSON(t, fixture.regenPR, pulls)
}

func rewriteRegenerationPRIdentity(t *testing.T, fixture *receiptFixture, mergeCommit, sourceCommit string) {
	t.Helper()
	data, err := os.ReadFile(fixture.regenPR) //nolint:gosec // isolated test fixture
	if err != nil {
		t.Fatal(err)
	}
	var pulls []map[string]any
	if err := json.Unmarshal(data, &pulls); err != nil {
		t.Fatal(err)
	}
	if len(pulls) != 1 {
		t.Fatalf("expected one regeneration PR fixture, found %d", len(pulls))
	}
	pulls[0]["merge_commit_sha"] = mergeCommit
	pulls[0]["head"].(map[string]any)["ref"] = "auto-regenerate/" + sourceCommit
	writeReleaseTestJSON(t, fixture.regenPR, pulls)
}

func rebindReceiptFixture(t *testing.T, fixture *receiptFixture, previousVersion, sourceCommit string) {
	t.Helper()
	receiptPath := filepath.Join(fixture.repo, "tools/spec-regeneration-receipt.json")
	data, err := os.ReadFile(receiptPath) //nolint:gosec // isolated test fixture
	if err != nil {
		t.Fatal(err)
	}
	var canonical map[string]any
	if err := json.Unmarshal(data, &canonical); err != nil {
		t.Fatal(err)
	}
	if previousVersion != "" {
		previous := maps.Clone(canonical)
		previous["version"] = previousVersion
		writeReleaseTestJSON(t, receiptPath, previous)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "add", receiptPath)
		runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "--amend", "--no-edit", "-q")
	}
	parent := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD"))
	if sourceCommit == "" {
		sourceCommit = parent
	}
	canonical["source_commit"] = sourceCommit
	writeReleaseTestJSON(t, receiptPath, canonical)
	runReleaseTestCommand(t, fixture.repo, nil, "git", "add", receiptPath)
	runReleaseTestCommand(t, fixture.repo, nil, "git", "commit", "-qm", "chore: auto-regenerate provider and documentation")
	releasedCommit := strings.TrimSpace(runReleaseTestCommand(t, fixture.repo, nil, "git", "rev-parse", "HEAD"))
	runReleaseTestCommand(t, fixture.repo, nil, "git", "tag", "-f", "v9.8.7")

	releasePath := envValue(t, fixture.env, "RELEASE_JSON")
	releaseData, err := os.ReadFile(releasePath) //nolint:gosec // isolated test fixture
	if err != nil {
		t.Fatal(err)
	}
	var release map[string]any
	if err := json.Unmarshal(releaseData, &release); err != nil {
		t.Fatal(err)
	}
	body := release["body"].(string)
	const markerPrefix = "<!-- provider-publication-receipt:"
	markerStart := strings.Index(body, markerPrefix)
	if markerStart < 0 {
		t.Fatalf("release fixture is missing its publication receipt marker")
	}
	markerEnd := strings.Index(body[markerStart:], " -->")
	if markerEnd < 0 {
		t.Fatalf("release fixture has an unterminated publication receipt marker")
	}
	markerEnd += markerStart
	var publication map[string]any
	if err := json.Unmarshal([]byte(body[markerStart+len(markerPrefix):markerEnd]), &publication); err != nil {
		t.Fatal(err)
	}
	publication["commit"] = releasedCommit
	publicationJSON, err := json.Marshal(publication)
	if err != nil {
		t.Fatal(err)
	}
	release["body"] = body[:markerStart+len(markerPrefix)] + string(publicationJSON) + body[markerEnd:]
	writeReleaseTestJSON(t, releasePath, release)

	rewriteRegenerationPRIdentity(t, fixture, releasedCommit, parent)
	runReleaseTestCommand(t, fixture.repo, nil, "git", "push", "-q", "--force", "origin", "HEAD:main", "v9.8.7")
	for i, assignment := range fixture.env {
		if strings.HasPrefix(assignment, "RELEASED_COMMIT=") {
			fixture.env[i] = "RELEASED_COMMIT=" + releasedCommit
		}
	}
}

func envValue(t *testing.T, env []string, name string) string {
	t.Helper()
	prefix := name + "="
	for _, assignment := range env {
		if strings.HasPrefix(assignment, prefix) {
			return strings.TrimPrefix(assignment, prefix)
		}
	}
	t.Fatalf("environment fixture is missing %s", name)
	return ""
}

func newReceiptFixture(t *testing.T, forgedLedger, falseDigest bool) *receiptFixture {
	t.Helper()
	root := testRepositoryRoot(t)
	repo := t.TempDir()
	origin := filepath.Join(t.TempDir(), "origin.git")
	runReleaseTestCommand(t, repo, nil, "git", "init", "-q")
	runReleaseTestCommand(t, repo, nil, "git", "config", "user.name", "Receipt Test")
	runReleaseTestCommand(t, repo, nil, "git", "config", "user.email", "receipt@example.com")
	validatorBytes, err := os.ReadFile(filepath.Join(root, "scripts", "validate-provider-delivery-state.sh")) //nolint:gosec // fixed repository script
	if err != nil {
		t.Fatal(err)
	}
	writeReleaseTestFile(t, repo, "scripts/validate-provider-delivery-state.sh", string(validatorBytes), 0o700)
	commit := strings.Repeat("b", 40)
	deliveryID := providerDeliveryID(t, "2.1.208", "v2.1.208", commit)
	validator, err := os.ReadFile(filepath.Join(root, "scripts", "validate-provider-delivery-state.sh")) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	writeReleaseTestFile(t, repo, "scripts/validate-provider-delivery-state.sh", string(validator), 0o700)
	pending := fmt.Sprintf(`{"delivery_id":%q,"release_tag":"v2.1.208","target_commit":%q,"version":"2.1.208"}`+"\n", deliveryID, commit)
	pin := fmt.Sprintf(`{"assets":{"api-catalog.json":%q,"concurrency_contracts.json":%q,"f5xc-api-specs-v2.1.208.zip":%q,"index.json":%q,"minimal-export-defaults.json":%q,"openapi.json":%q,"smsv2-contract-manifest.json":%q,"smsv2-contract.json":%q,"smsv2-evidence-receipt.json":%q,"smsv2_parity_manifest.json":%q,"upstream-contract-removals.json":%q},"release_tag":"v2.1.208","target_commit":%q,"version":"2.1.208"}`+"\n", "sha256:"+strings.Repeat("1", 64), "sha256:"+strings.Repeat("2", 64), "sha256:"+strings.Repeat("3", 64), "sha256:"+strings.Repeat("4", 64), "sha256:"+strings.Repeat("5", 64), "sha256:"+strings.Repeat("6", 64), "sha256:"+strings.Repeat("7", 64), "sha256:"+strings.Repeat("8", 64), "sha256:"+strings.Repeat("9", 64), "sha256:"+strings.Repeat("a", 64), "sha256:"+strings.Repeat("b", 64), commit)
	writeReleaseTestFile(t, repo, "tools/spec-pending-delivery.json", pending, 0o600)
	writeReleaseTestFile(t, repo, "tools/spec-release.json", pin, 0o600)
	writeReleaseTestFile(t, repo, "tools/spec-version.txt", "v2.1.208\n", 0o600)
	writeReleaseTestFile(t, repo, "tools/spec-deliveries.json", "{\"deliveries\":{},\"version\":1}\n", 0o600)
	writeReleaseTestFile(t, repo, "tools/provider-publication-receipts.json", "{\"receipts\":{},\"version\":1}\n", 0o600)
	runReleaseTestCommand(t, repo, nil, "git", "add", ".")
	runReleaseTestCommand(t, repo, nil, "git", "commit", "-qm", "source")
	sourceCommit := strings.TrimSpace(runReleaseTestCommand(t, repo, nil, "git", "rev-parse", "HEAD"))
	pinSHA := releaseTestSHA(t, filepath.Join(repo, "tools/spec-release.json"))
	regeneration := fmt.Sprintf(`{"delivery_id":%q,"release_tag":"v2.1.208","source_commit":%q,"spec_release_sha256":%q,"target_commit":%q,"version":"2.1.208"}`+"\n", deliveryID, sourceCommit, pinSHA, commit)
	writeReleaseTestFile(t, repo, "tools/spec-regeneration-receipt.json", regeneration, 0o600)
	runReleaseTestCommand(t, repo, nil, "git", "add", ".")
	runReleaseTestCommand(t, repo, nil, "git", "commit", "-qm", "chore: auto-regenerate provider and documentation")
	releasedCommit := strings.TrimSpace(runReleaseTestCommand(t, repo, nil, "git", "rev-parse", "HEAD"))
	providerTag := "v9.8.7"
	runReleaseTestCommand(t, repo, nil, "git", "tag", providerTag)

	providerAssets := map[string]string{}
	for i, name := range providerReleaseAssetNames("9.8.7") {
		providerAssets[name] = fmt.Sprintf("%064x", i+1)
	}
	// providerAssets stays raw hex because the GitHub release fixture below builds
	// its .digest field by prefixing it. The receipt records the qualified
	// "sha256:<hex>" form, so the two compare by equality without prefix arithmetic.
	providerReceipt := map[string]any{
		"assets": qualifiedAssetDigests(providerAssets), "commit": releasedCommit,
		"spec_release_sha256": pinSHA,
		"tag":                 providerTag, "version": "9.8.7",
	}
	if forgedLedger {
		forgedKey := strings.Repeat("f", 64)
		forgedDelivery := map[string]any{"release_tag": "v2.1.100", "target_commit": strings.Repeat("c", 40), "version": "2.1.100"}
		common := map[string]any{"deliveries": map[string]any{forgedKey: forgedDelivery}, "version": 1}
		detailed := map[string]any{"receipts": map[string]any{forgedKey: map[string]any{"delivery": forgedDelivery, "publication": providerReceipt}}, "version": 1}
		writeReleaseTestJSON(t, filepath.Join(repo, "tools/spec-deliveries.json"), common)
		writeReleaseTestJSON(t, filepath.Join(repo, "tools/provider-publication-receipts.json"), detailed)
		runReleaseTestCommand(t, repo, nil, "git", "add", ".")
		runReleaseTestCommand(t, repo, nil, "git", "commit", "--amend", "--no-edit", "-q")
		releasedCommit = strings.TrimSpace(runReleaseTestCommand(t, repo, nil, "git", "rev-parse", "HEAD"))
		runReleaseTestCommand(t, repo, nil, "git", "tag", "-f", providerTag)
		providerReceipt["commit"] = releasedCommit
	}
	receiptBytes, err := json.Marshal(providerReceipt)
	if err != nil {
		t.Fatal(err)
	}
	releaseAssets := make([]map[string]string, 0, len(providerAssets))
	for _, name := range providerReleaseAssetNames("9.8.7") {
		digest := providerAssets[name]
		if falseDigest && name == "mcp-data-9.8.7.tar.gz" {
			digest = strings.Repeat("0", 64)
		}
		releaseAssets = append(releaseAssets, map[string]string{"name": name, "digest": "sha256:" + digest})
	}
	release := map[string]any{"tag_name": providerTag, "draft": false, "prerelease": false, "immutable": true, "assets": releaseAssets, "body": "notes\n<!-- provider-publication-receipt:" + string(receiptBytes) + " -->\n"}
	releasePath := filepath.Join(repo, "release.json")
	writeReleaseTestJSON(t, releasePath, release)
	regenPRPath := filepath.Join(t.TempDir(), "regeneration-pr.json")
	writeReleaseTestJSON(t, regenPRPath, []map[string]any{{
		"number":           42,
		"state":            "closed",
		"merged_at":        "2026-08-01T12:00:00Z",
		"merge_commit_sha": releasedCommit,
		"title":            "chore: auto-regenerate provider and documentation",
		"base": map[string]any{
			"ref": "main", "repo": map[string]any{"full_name": dispatchTarget},
		},
		"head": map[string]any{
			"ref": "auto-regenerate/" + sourceCommit, "repo": map[string]any{"full_name": dispatchTarget},
		},
		"labels": []map[string]string{{"name": "automated"}, {"name": "regeneration"}},
	}})
	binDir := filepath.Join(repo, "bin")
	writeReleaseTestFile(t, binDir, "gh", `#!/usr/bin/env bash
set -euo pipefail
if [[ "$*" == *"/commits/"*"/pulls"* ]]; then
	if [ -n "${REGENERATION_PR_CALL_COUNT:-}" ]; then
		calls=0
		if [ -f "$REGENERATION_PR_CALL_COUNT" ]; then
			calls=$(cat "$REGENERATION_PR_CALL_COUNT")
		fi
		calls=$((calls + 1))
		printf '%s\n' "$calls" > "$REGENERATION_PR_CALL_COUNT"
		if [ "$calls" -le "${REGENERATION_PR_EMPTY_RESPONSES:-0}" ]; then
			printf '[]\n'
			exit 0
		fi
	fi
  cat "$REGENERATION_PR_JSON"
elif [[ "$*" =~ /pulls/[0-9]+$ ]]; then
	jq -c '.[0]' "$REGENERATION_PR_JSON"
else
  cat "$RELEASE_JSON"
fi
`, 0o700)
	runReleaseTestCommand(t, repo, nil, "git", "init", "--bare", "-q", origin)
	runReleaseTestCommand(t, repo, nil, "git", "branch", "-M", "main")
	runReleaseTestCommand(t, repo, nil, "git", "remote", "add", "origin", origin)
	runReleaseTestCommand(t, repo, nil, "git", "push", "-q", "origin", "main", "--tags", "--force")
	output := filepath.Join(repo, "github-output")
	return &receiptFixture{
		repo: repo, script: filepath.Join(root, "scripts", "prepare-spec-delivery-receipt.sh"), output: output,
		deliveryID: deliveryID, regenPR: regenPRPath,
		env: []string{"PATH=" + binDir + string(os.PathListSeparator) + os.Getenv("PATH"), "RELEASE_JSON=" + releasePath,
			"REGENERATION_PR_JSON=" + regenPRPath, "PROVIDER_TAG=" + providerTag, "RELEASED_COMMIT=" + releasedCommit,
			"TARGET_REPOSITORY=" + dispatchTarget, "GITHUB_OUTPUT=" + output, "GH_TOKEN=fixture",
			"PR_ASSOCIATION_RETRY_DELAY_SECONDS=0"},
	}
}

func providerReleaseAssetNames(version string) []string {
	return []string{
		"terraform-provider-xcsh_" + version + "_darwin_amd64.zip",
		"terraform-provider-xcsh_" + version + "_darwin_arm64.zip",
		"terraform-provider-xcsh_" + version + "_freebsd_386.zip",
		"terraform-provider-xcsh_" + version + "_freebsd_amd64.zip",
		"terraform-provider-xcsh_" + version + "_linux_386.zip",
		"terraform-provider-xcsh_" + version + "_linux_amd64.zip",
		"terraform-provider-xcsh_" + version + "_linux_arm.zip",
		"terraform-provider-xcsh_" + version + "_linux_arm64.zip",
		"terraform-provider-xcsh_" + version + "_manifest.json",
		"terraform-provider-xcsh_" + version + "_SHA256SUMS",
		"terraform-provider-xcsh_" + version + "_SHA256SUMS.sig",
		"terraform-provider-xcsh_" + version + "_windows_386.zip",
		"terraform-provider-xcsh_" + version + "_windows_amd64.zip",
		"mcp-data-" + version + ".tar.gz",
	}
}

func writeProviderReleaseJSON(t *testing.T, path, tag, body, assetDir string, names []string) {
	t.Helper()
	assets := make([]map[string]string, 0, len(names))
	for _, name := range names {
		assets = append(assets, map[string]string{"name": name, "digest": "sha256:" + releaseTestSHA(t, filepath.Join(assetDir, name))})
	}
	writeReleaseTestJSON(t, path, map[string]any{"tag_name": tag, "draft": true, "prerelease": false, "body": body, "assets": assets})
}

func writeReleaseTestJSON(t *testing.T, path string, value any) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, append(data, '\n'), 0o600); err != nil {
		t.Fatal(err)
	}
}

func writeReleaseTestFile(t *testing.T, root, name, body string, mode os.FileMode) {
	t.Helper()
	path := filepath.Join(root, name)
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), mode); err != nil {
		t.Fatal(err)
	}
}

func releaseTestSHA(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path) //nolint:gosec // isolated test fixtures
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func runReleaseTestCommand(t *testing.T, dir string, extraEnv []string, name string, args ...string) string {
	t.Helper()
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), extraEnv...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("%s %s failed: %v\n%s", name, strings.Join(args, " "), err, output)
	}
	return string(output)
}

func runReleaseTestCommandStdout(t *testing.T, dir string, extraEnv []string, name string, args ...string) string {
	t.Helper()
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), extraEnv...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("%s %s failed: %v\n%s", name, strings.Join(args, " "), err, stderr.String())
	}
	return string(output)
}

func runWorkflowScript(dir, script string, extraEnv []string) (string, error) {
	cmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), extraEnv...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func runImplicitWorkflowScript(dir, script string, extraEnv []string) (string, error) {
	cmd := exec.Command("bash", "--noprofile", "--norc", "-e", "-c", script)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), extraEnv...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func extractWorkflowRunStep(t *testing.T, workflowName, jobName, stepName string) string {
	t.Helper()
	path := filepath.Join(testRepositoryRoot(t), ".github", "workflows", workflowName)
	data, err := os.ReadFile(path) //nolint:gosec // fixed repository path
	if err != nil {
		t.Fatal(err)
	}
	var workflow struct {
		Jobs map[string]struct {
			Steps []struct {
				Name string `yaml:"name"`
				Run  string `yaml:"run"`
			} `yaml:"steps"`
		} `yaml:"jobs"`
	}
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		t.Fatal(err)
	}
	for _, step := range workflow.Jobs[jobName].Steps {
		if step.Name == stepName {
			return step.Run
		}
	}
	t.Fatalf("workflow step %q/%q not found", jobName, stepName)
	return ""
}

func assertJSONPathExists(t *testing.T, path, expression string) {
	t.Helper()
	cmd := exec.Command("jq", "-e", expression, path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("jq %s %s failed: %v\n%s", expression, path, err, output)
	}
}

func testRepositoryRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	return root
}

func testFutureTime() (value time.Time) {
	return time.Unix(2_000_000_000, 0)
}
