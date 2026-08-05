// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

// The download-api-specs composite action probes for an existing bundle before
// downloading. Its target directory is gitignored, so on a hosted runner it does not
// exist — which is precisely the case the probe has to survive.
//
// It did not. A composite action's `shell: bash` runs as `bash --noprofile --norc -eo
// pipefail`, so `find "$missing" | wc -l` returns 1 through pipefail and `set -e`
// killed the step before mkdir or any download ran. The three hand-rolled copies this
// action replaced were not affected: a workflow `run:` with no `shell:` key uses
// `bash -e {0}` WITHOUT pipefail, so the same line was harmless there. Consolidating
// them is what exposed it.
//
// This test extracts the action's script and runs it against an absent directory with
// a stub `gh` on PATH, asserting it reaches the download rather than dying at the
// probe.
func TestDownloadSpecsActionSurvivesAbsentSpecDir(t *testing.T) {
	script := extractActionScript(t)

	tmp := t.TempDir()
	writeSpecReleasePin(t, tmp, "v2.1.207", strings.Repeat("a", 40), nil)

	// A stub `gh` that records it was called and then fails. Reaching it is the
	// assertion; what it returns afterwards is not this test's concern.
	binDir := filepath.Join(tmp, "bin")
	if err := os.Mkdir(binDir, 0o750); err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(tmp, "gh-was-called")
	stub := "#!/usr/bin/env bash\ntouch " + marker + "\nexit 7\n"
	if err := os.WriteFile(filepath.Join(binDir, "gh"), []byte(stub), 0o700); err != nil { //nolint:gosec // test stub must be executable
		t.Fatal(err)
	}

	absent := filepath.Join(tmp, "no", "such", "spec", "dir")

	cmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = tmp
	cmd.Env = append(os.Environ(),
		"PATH="+binDir+string(os.PathListSeparator)+os.Getenv("PATH"),
		"SPEC_DIR="+absent,
		"SPEC_RELEASE_TAG=v2.1.207",
		"ENRICHED_REPO=f5-sales-demo/api-specs-enriched",
		"GH_TOKEN=stub",
		"GITHUB_WORKSPACE="+tmp,
	)
	out, err := cmd.CombinedOutput()

	if _, statErr := os.Stat(marker); statErr != nil {
		t.Fatalf("action exited before attempting the download against an absent %s.\n"+
			"The probe must not kill the step on a clean checkout.\nerr: %v\noutput:\n%s",
			absent, err, out)
	}

	if _, statErr := os.Stat(absent); !os.IsNotExist(statErr) {
		t.Errorf("failed download mutated the destination instead of staging atomically: %v", statErr)
	}
}

// A plausible local tree is not proof of provenance. The action must revalidate the
// immutable release receipt rather than trusting files that happen to claim the same
// version.
func TestDownloadSpecsActionRevalidatesExistingBundle(t *testing.T) {
	script := extractActionScript(t)

	tmp := t.TempDir()
	writeSpecReleasePin(t, tmp, "v2.1.207", strings.Repeat("a", 40), nil)
	specDir := filepath.Join(tmp, "specs")
	if err := os.MkdirAll(filepath.Join(specDir, "domains"), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(specDir, "index.json"), []byte(`{"specifications":[{"name":"sites"}]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(specDir, "domains", "sites.json"), []byte(`{"paths":{}}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(specDir, "api-catalog.json"), []byte(`{"version":"2.1.207"}`), 0o600); err != nil {
		t.Fatal(err)
	}

	binDir := filepath.Join(tmp, "bin")
	if err := os.Mkdir(binDir, 0o750); err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(tmp, "gh-was-called")
	stub := "#!/usr/bin/env bash\ntouch " + marker + "\nexit 7\n"
	if err := os.WriteFile(filepath.Join(binDir, "gh"), []byte(stub), 0o700); err != nil { //nolint:gosec // test stub must be executable
		t.Fatal(err)
	}

	cmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = tmp
	cmd.Env = append(os.Environ(),
		"PATH="+binDir+string(os.PathListSeparator)+os.Getenv("PATH"),
		"SPEC_DIR="+specDir,
		"SPEC_RELEASE_TAG=v2.1.207",
		"ENRICHED_REPO=f5-sales-demo/api-specs-enriched",
		"GH_TOKEN=stub",
		"GITHUB_WORKSPACE="+tmp,
	)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("unreceipted local bundle was trusted without release validation:\n%s", out)
	}
	if _, statErr := os.Stat(marker); statErr != nil {
		t.Fatalf("action did not query the release for an existing local bundle: %v", statErr)
	}
	if _, statErr := os.Stat(filepath.Join(specDir, "index.json")); statErr != nil {
		t.Fatalf("failed validation destroyed the previous local bundle: %v", statErr)
	}
}

// A truncated unzip leaves index.json with an empty domains/. That is JSON in the
// directory, so a "some spec files are present" probe short-circuits on it and the
// download that would repair it never runs — while openapi.RequireSpecs rejects the
// same tree downstream. Measured before this was fixed: such a bundle let
// transform-docs.go exit 0 and rewrite 277 files. The action's completeness test has
// to agree with the guard the generators use.
func TestDownloadSpecsActionRedownloadsTruncatedBundle(t *testing.T) {
	script := extractActionScript(t)

	tmp := t.TempDir()
	writeSpecReleasePin(t, tmp, "v2.1.207", strings.Repeat("a", 40), nil)
	specDir := filepath.Join(tmp, "specs")
	if err := os.MkdirAll(filepath.Join(specDir, "domains"), 0o750); err != nil {
		t.Fatal(err)
	}
	// index.json present, domains/ empty — the truncated shape.
	if err := os.WriteFile(filepath.Join(specDir, "index.json"), []byte(`{}`), 0o600); err != nil {
		t.Fatal(err)
	}

	binDir := filepath.Join(tmp, "bin")
	if err := os.Mkdir(binDir, 0o750); err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(tmp, "gh-was-called")
	stub := "#!/usr/bin/env bash\ntouch " + marker + "\nexit 7\n"
	if err := os.WriteFile(filepath.Join(binDir, "gh"), []byte(stub), 0o700); err != nil { //nolint:gosec // test stub must be executable
		t.Fatal(err)
	}

	cmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = tmp
	cmd.Env = append(os.Environ(),
		"PATH="+binDir+string(os.PathListSeparator)+os.Getenv("PATH"),
		"SPEC_DIR="+specDir,
		"SPEC_RELEASE_TAG=v2.1.207",
		"ENRICHED_REPO=f5-sales-demo/api-specs-enriched",
		"GH_TOKEN=stub",
		"GITHUB_WORKSPACE="+tmp,
	)
	out, _ := cmd.CombinedOutput()

	if _, err := os.Stat(marker); err != nil {
		t.Fatalf("action treated a truncated bundle (index.json, empty domains/) as complete "+
			"and never attempted a download.\noutput:\n%s", out)
	}
}

// Every lookup must be scoped to the tracked tag, and the standalone catalog must
// survive promotion into SPEC_DIR. The old implementation selected .[0] from the
// release list and downloaded the catalog before deleting SPEC_DIR/*, so a newer
// release could be mixed into an older config and the catalog was always lost.
func TestDownloadSpecsActionUsesPinnedReleaseAndRetainsCatalog(t *testing.T) {
	script := extractActionScript(t)
	tmp := t.TempDir()
	specDir := filepath.Join(tmp, "specs")
	binDir := filepath.Join(tmp, "bin")
	if err := os.Mkdir(binDir, 0o750); err != nil {
		t.Fatal(err)
	}

	bundle := filepath.Join(tmp, "bundle.zip")
	writeTestSpecBundle(t, bundle)

	catalog := filepath.Join(tmp, "catalog.json")
	if err := os.WriteFile(catalog, []byte(`{"version":"2.1.208","commands":[]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	minimal := filepath.Join(tmp, "minimal-export-defaults.json")
	if err := os.WriteFile(minimal, []byte(`{"version":"2.1.208","defaults":{}}`), 0o600); err != nil {
		t.Fatal(err)
	}
	index := filepath.Join(tmp, "index.json")
	if err := os.WriteFile(index, []byte(`{"specifications":[{"file":"domains/sites.json"}]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	openapi := filepath.Join(tmp, "openapi.json")
	if err := os.WriteFile(openapi, []byte(`{"openapi":"3.0.0","paths":{}}`), 0o600); err != nil {
		t.Fatal(err)
	}
	tagCommit := strings.Repeat("a", 40)
	writeSpecReleasePin(t, tmp, "v2.1.208", tagCommit, map[string]string{
		"api-catalog.json":             fileSHA256(t, catalog),
		"f5xc-api-specs-v2.1.208.zip":  fileSHA256(t, bundle),
		"index.json":                   fileSHA256(t, index),
		"minimal-export-defaults.json": fileSHA256(t, minimal),
		"openapi.json":                 fileSHA256(t, openapi),
	})
	logPath := filepath.Join(tmp, "gh.log")
	stub := `#!/usr/bin/env bash
printf '%s\n' "$*" >> "$GH_LOG"
case "$*" in
  *releases/assets/101*) cat "$BUNDLE_ZIP" ;;
  *releases/assets/202*) cat "$CATALOG_FILE" ;;
  *releases/assets/203*) cat "$MINIMAL_FILE" ;;
  *releases/assets/204*) cat "$INDEX_FILE" ;;
  *releases/assets/205*) cat "$OPENAPI_FILE" ;;
  *commits/v2.1.208*) printf '%s\n' "$TAG_COMMIT" ;;
  *releases/tags/v2.1.208*)
    bundle_sha=$(shasum -a 256 "$BUNDLE_ZIP" | awk '{print $1}')
    catalog_sha=$(shasum -a 256 "$CATALOG_FILE" | awk '{print $1}')
    index_sha=$(shasum -a 256 "$INDEX_FILE" | awk '{print $1}')
    minimal_sha=$(shasum -a 256 "$MINIMAL_FILE" | awk '{print $1}')
    openapi_sha=$(shasum -a 256 "$OPENAPI_FILE" | awk '{print $1}')
    jq -cn \
      --arg bundle_sha "$bundle_sha" \
      --arg catalog_sha "$catalog_sha" \
      --arg commit "$TAG_COMMIT" \
      --arg index_sha "$index_sha" \
      --arg minimal_sha "$minimal_sha" \
      --arg openapi_sha "$openapi_sha" '
      {tag_name: "v2.1.208", draft: false, prerelease: false, immutable: true, assets: [
        {id: 202, name: "api-catalog.json", digest: ("sha256:" + $catalog_sha)},
        {id: 101, name: "f5xc-api-specs-v2.1.208.zip", digest: ("sha256:" + $bundle_sha)},
        {id: 204, name: "index.json", digest: ("sha256:" + $index_sha)},
        {id: 203, name: "minimal-export-defaults.json", digest: ("sha256:" + $minimal_sha)},
        {id: 205, name: "openapi.json", digest: ("sha256:" + $openapi_sha)}
      ], body: ("<!-- publication-receipt:" + ({
        assets: {
          "api-catalog.json": ("sha256:" + $catalog_sha),
          "f5xc-api-specs-v2.1.208.zip": ("sha256:" + $bundle_sha),
          "index.json": ("sha256:" + $index_sha),
          "minimal-export-defaults.json": ("sha256:" + $minimal_sha),
          "openapi.json": ("sha256:" + $openapi_sha)
        }, commit: $commit, version: "2.1.208"
      } | tojson) + " -->")}' ;;
  *) echo "unexpected gh invocation: $*" >&2; exit 9 ;;
esac
`
	if err := os.WriteFile(filepath.Join(binDir, "gh"), []byte(stub), 0o700); err != nil { //nolint:gosec // test stub must be executable
		t.Fatal(err)
	}

	actionEnv := []string{
		"PATH=" + binDir + string(os.PathListSeparator) + os.Getenv("PATH"),
		"SPEC_DIR=" + specDir,
		"SPEC_RELEASE_TAG=v2.1.208",
		"ENRICHED_REPO=f5-sales-demo/api-specs-enriched",
		"GH_TOKEN=stub",
		"GITHUB_WORKSPACE=" + tmp,
		"GH_LOG=" + logPath,
		"BUNDLE_ZIP=" + bundle,
		"CATALOG_FILE=" + catalog,
		"MINIMAL_FILE=" + minimal,
		"INDEX_FILE=" + index,
		"OPENAPI_FILE=" + openapi,
		"TAG_COMMIT=" + tagCommit,
	}
	cmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = tmp
	cmd.Env = append(os.Environ(), actionEnv...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("exact-release download failed: %v\n%s", err, out)
	}

	gotCatalog, err := os.ReadFile(filepath.Join(specDir, "api-catalog.json")) //nolint:gosec // isolated test path
	if err != nil {
		t.Fatalf("promoted bundle lost api-catalog.json: %v", err)
	}
	if string(gotCatalog) != `{"version":"2.1.208","commands":[]}` {
		t.Fatalf("promoted catalog differs from dispatched release asset: %s", gotCatalog)
	}
	logBytes, err := os.ReadFile(logPath) //nolint:gosec // isolated test path
	if err != nil {
		t.Fatal(err)
	}
	log := string(logBytes)
	if !strings.Contains(log, "releases/tags/v2.1.208") {
		t.Fatalf("release lookup was not pinned to v2.1.208:\n%s", log)
	}
	if strings.Contains(log, "api repos/f5-sales-demo/api-specs-enriched/releases --jq") {
		t.Fatalf("release lookup used the mutable release list:\n%s", log)
	}

	writeSymlinkSpecBundle(t, bundle)
	writeSpecReleasePin(t, tmp, "v2.1.208", tagCommit, map[string]string{
		"api-catalog.json":             fileSHA256(t, catalog),
		"f5xc-api-specs-v2.1.208.zip":  fileSHA256(t, bundle),
		"index.json":                   fileSHA256(t, index),
		"minimal-export-defaults.json": fileSHA256(t, minimal),
		"openapi.json":                 fileSHA256(t, openapi),
	})
	sentinel := filepath.Join(specDir, "existing-bundle-must-survive")
	if err := os.WriteFile(sentinel, []byte("preserve\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	cmd = exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = tmp
	cmd.Env = append(os.Environ(), actionEnv...)
	out, err = cmd.CombinedOutput()
	if err == nil || !strings.Contains(string(out), "unsupported archive entry type") {
		t.Fatalf("symlink-bearing spec archive was accepted: err=%v\n%s", err, out)
	}
	if _, err := os.Stat(sentinel); err != nil {
		t.Fatalf("rejected archive mutated the prior spec tree: %v", err)
	}
}

func TestDownloadSpecsActionRejectsDestinationOutsideWorkspace(t *testing.T) {
	script := extractActionScript(t)
	workspace := t.TempDir()
	outside := filepath.Join(t.TempDir(), "specs")
	writeSpecReleasePin(t, workspace, "v2.1.208", strings.Repeat("a", 40), nil)
	cmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = workspace
	cmd.Env = append(os.Environ(),
		"SPEC_DIR="+outside,
		"SPEC_RELEASE_TAG=v2.1.208",
		"ENRICHED_REPO=f5-sales-demo/api-specs-enriched",
		"GH_TOKEN=stub",
		"GITHUB_WORKSPACE="+workspace,
	)
	output, err := cmd.CombinedOutput()
	if err == nil || !strings.Contains(string(output), "non-root descendant of GITHUB_WORKSPACE") {
		t.Fatalf("destination outside the workspace was accepted: err=%v\n%s", err, output)
	}
	if _, statErr := os.Stat(outside); !os.IsNotExist(statErr) {
		t.Fatalf("rejected outside destination was mutated: %v", statErr)
	}
}

func TestDownloadSpecsActionRejectsMutableRelease(t *testing.T) {
	tag := "v2.1.208"
	commit := strings.Repeat("a", 40)
	receipt := testSpecPublicationReceipt(tag, commit)
	metadata := testSpecReleaseMetadata(t, tag, false, []map[string]any{receipt}, nil)

	output, err := runDownloadActionWithMetadata(t, metadata, commit)
	if err == nil || !strings.Contains(output, "exact tag and be final and immutable") {
		t.Fatalf("mutable release was accepted: err=%v\n%s", err, output)
	}
}

func TestDownloadSpecsActionRejectsContradictoryPublicationReceipt(t *testing.T) {
	tag := "v2.1.208"
	commit := strings.Repeat("a", 40)
	tests := []struct {
		name   string
		mutate func(map[string]any)
	}{
		{
			name: "commit differs from tag and pin",
			mutate: func(receipt map[string]any) {
				receipt["commit"] = strings.Repeat("b", 40)
			},
		},
		{
			name: "asset digest differs from pin",
			mutate: func(receipt map[string]any) {
				receipt["assets"].(map[string]string)["index.json"] = strings.Repeat("f", 64)
			},
		},
		{
			name: "asset set is incomplete",
			mutate: func(receipt map[string]any) {
				delete(receipt["assets"].(map[string]string), "openapi.json")
			},
		},
		{
			name: "version differs from pin",
			mutate: func(receipt map[string]any) {
				receipt["version"] = "2.1.209"
			},
		},
		{
			name: "receipt has an uncontracted key",
			mutate: func(receipt map[string]any) {
				receipt["mutable_alias"] = true
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			receipt := testSpecPublicationReceipt(tag, commit)
			test.mutate(receipt)
			metadata := testSpecReleaseMetadata(t, tag, true, []map[string]any{receipt}, nil)

			output, err := runDownloadActionWithMetadata(t, metadata, commit)
			if err == nil || !strings.Contains(output, "does not exactly match the tracked release pin") {
				t.Fatalf("contradictory publication receipt was accepted: err=%v\n%s", err, output)
			}
		})
	}
}

func TestDownloadSpecsActionRequiresExactlyOnePublicationReceipt(t *testing.T) {
	tag := "v2.1.208"
	commit := strings.Repeat("a", 40)
	receipt := testSpecPublicationReceipt(tag, commit)
	tests := []struct {
		name     string
		receipts []map[string]any
	}{
		{name: "missing", receipts: nil},
		{name: "duplicate", receipts: []map[string]any{receipt, receipt}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			metadata := testSpecReleaseMetadata(t, tag, true, test.receipts, nil)
			output, err := runDownloadActionWithMetadata(t, metadata, commit)
			if err == nil || !strings.Contains(output, "exactly one publication receipt") {
				t.Fatalf("%s publication receipt set was accepted: err=%v\n%s", test.name, err, output)
			}
		})
	}
}

func TestDownloadSpecsActionRejectsWrongTagIdentity(t *testing.T) {
	tag := "v2.1.208"
	commit := strings.Repeat("a", 40)
	receipt := testSpecPublicationReceipt(tag, commit)
	tests := []struct {
		name               string
		releaseTag         string
		resolvedCommit     string
		expectedDiagnostic string
	}{
		{
			name:               "release metadata names another tag",
			releaseTag:         "v2.1.209",
			resolvedCommit:     commit,
			expectedDiagnostic: "exact tag and be final and immutable",
		},
		{
			name:               "tag resolves to another commit",
			releaseTag:         tag,
			resolvedCommit:     strings.Repeat("b", 40),
			expectedDiagnostic: "tag commit differs from the expected upstream commit",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			metadata := testSpecReleaseMetadata(t, test.releaseTag, true, []map[string]any{receipt}, nil)
			output, err := runDownloadActionWithMetadata(t, metadata, test.resolvedCommit)
			if err == nil || !strings.Contains(output, test.expectedDiagnostic) {
				t.Fatalf("wrong tag identity was accepted: err=%v\n%s", err, output)
			}
		})
	}
}

func TestDownloadSpecsActionRejectsGitHubDigestMismatch(t *testing.T) {
	tag := "v2.1.208"
	commit := strings.Repeat("a", 40)
	receipt := testSpecPublicationReceipt(tag, commit)
	metadata := testSpecReleaseMetadata(t, tag, true, []map[string]any{receipt}, map[string]string{
		"index.json": strings.Repeat("f", 64),
	})

	output, err := runDownloadActionWithMetadata(t, metadata, commit)
	if err == nil || !strings.Contains(output, "GitHub asset digests do not exactly match") {
		t.Fatalf("GitHub digest contradicting the receipt was accepted: err=%v\n%s", err, output)
	}
}

func runDownloadActionWithMetadata(t *testing.T, metadata []byte, resolvedCommit string) (string, error) {
	t.Helper()
	tmp := t.TempDir()
	tag := "v2.1.208"
	writeSpecReleasePin(t, tmp, tag, strings.Repeat("a", 40), nil)

	metadataPath := filepath.Join(tmp, "release.json")
	if err := os.WriteFile(metadataPath, metadata, 0o600); err != nil {
		t.Fatal(err)
	}
	binDir := filepath.Join(tmp, "bin")
	if err := os.Mkdir(binDir, 0o750); err != nil {
		t.Fatal(err)
	}
	stub := `#!/usr/bin/env bash
case "$*" in
  *releases/tags/v2.1.208*) cat "$STUB_RELEASE_METADATA" ;;
  *commits/v2.1.208*) printf '%s\n' "$RESOLVED_COMMIT" ;;
  *) echo "unexpected gh invocation: $*" >&2; exit 88 ;;
esac
`
	if err := os.WriteFile(filepath.Join(binDir, "gh"), []byte(stub), 0o700); err != nil { //nolint:gosec // executable test stub
		t.Fatal(err)
	}

	cmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", extractActionScript(t))
	cmd.Dir = tmp
	cmd.Env = append(os.Environ(),
		"PATH="+binDir+string(os.PathListSeparator)+os.Getenv("PATH"),
		"SPEC_DIR="+filepath.Join(tmp, "specs"),
		"SPEC_RELEASE_TAG="+tag,
		"ENRICHED_REPO=f5-sales-demo/api-specs-enriched",
		"GH_TOKEN=stub",
		"GITHUB_WORKSPACE="+tmp,
		"STUB_RELEASE_METADATA="+metadataPath,
		"RESOLVED_COMMIT="+resolvedCommit,
	)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func testSpecPublicationReceipt(tag, commit string) map[string]any {
	return map[string]any{
		"assets":  qualifiedAssetDigests(testSpecAssetDigests(tag, nil)),
		"commit":  commit,
		"version": strings.TrimPrefix(tag, "v"),
	}
}

func testSpecReleaseMetadata(
	t *testing.T,
	tag string,
	immutable bool,
	receipts []map[string]any,
	digestOverrides map[string]string,
) []byte {
	t.Helper()
	digests := testSpecAssetDigests("v2.1.208", digestOverrides)
	names := []string{
		"api-catalog.json",
		"f5xc-api-specs-v2.1.208.zip",
		"index.json",
		"minimal-export-defaults.json",
		"openapi.json",
	}
	assets := make([]map[string]any, 0, len(names))
	for id, name := range names {
		assets = append(assets, map[string]any{
			"digest": "sha256:" + digests[name],
			"id":     id + 1,
			"name":   name,
		})
	}
	body := "release notes\n"
	for _, receipt := range receipts {
		receiptJSON, err := json.Marshal(receipt)
		if err != nil {
			t.Fatal(err)
		}
		body += "<!-- publication-receipt:" + string(receiptJSON) + " -->\n"
	}
	metadata, err := json.Marshal(map[string]any{
		"assets":     assets,
		"body":       body,
		"draft":      false,
		"immutable":  immutable,
		"prerelease": false,
		"tag_name":   tag,
	})
	if err != nil {
		t.Fatal(err)
	}
	return metadata
}

func testSpecAssetDigests(tag string, overrides map[string]string) map[string]string {
	bundle := "f5xc-api-specs-" + tag + ".zip"
	digests := map[string]string{
		"api-catalog.json":             strings.Repeat("0", 64),
		bundle:                         strings.Repeat("0", 64),
		"index.json":                   strings.Repeat("0", 64),
		"minimal-export-defaults.json": strings.Repeat("0", 64),
		"openapi.json":                 strings.Repeat("0", 64),
	}
	for name, digest := range overrides {
		digests[name] = digest
	}
	return digests
}

// qualifiedAssetDigests converts raw hex digests into the "sha256:<hex>" form the
// pin and the publication receipt both record — the same form GitHub reports in
// .assets[].digest, so the three compare by equality with no prefix arithmetic.
//
// testSpecAssetDigests deliberately keeps returning raw hex: testSpecReleaseMetadata
// builds GitHub's field by prefixing it itself, and the mutation tests override
// individual digests with raw values. Qualifying at the point of use rather than at
// the source keeps both of those call sites honest.
func qualifiedAssetDigests(digests map[string]string) map[string]string {
	qualified := make(map[string]string, len(digests))
	for name, digest := range digests {
		if strings.HasPrefix(digest, "sha256:") {
			qualified[name] = digest
			continue
		}
		qualified[name] = "sha256:" + digest
	}
	return qualified
}

func writeSpecReleasePin(
	t *testing.T,
	root, tag, commit string,
	overrides map[string]string,
) {
	t.Helper()
	version := strings.TrimPrefix(tag, "v")
	digests := qualifiedAssetDigests(testSpecAssetDigests(tag, overrides))
	pin := map[string]any{
		"assets":        digests,
		"release_tag":   tag,
		"target_commit": commit,
		"version":       version,
	}
	body, err := json.MarshalIndent(pin, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	toolsDir := filepath.Join(root, "tools")
	if err := os.MkdirAll(toolsDir, 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(toolsDir, "spec-release.json"), append(body, '\n'), 0o600); err != nil {
		t.Fatal(err)
	}
}

func fileSHA256(t *testing.T, path string) string {
	t.Helper()
	body, err := os.ReadFile(path) //nolint:gosec // isolated test path
	if err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf("%x", sha256.Sum256(body))
}

// extractActionScript returns the shell body of the composite action's single step,
// with GitHub's ${{ }} input expansions resolved to the environment variables the
// tests set. Reading the real file means the test cannot drift from what CI runs.
func extractActionScript(t *testing.T) string {
	t.Helper()

	path := filepath.Join("..", "..", ".github", "actions", "download-api-specs", "action.yml")
	data, err := os.ReadFile(path) //nolint:gosec // fixed repo-relative path
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}

	var action struct {
		Runs struct {
			Steps []struct {
				Run string `yaml:"run"`
			} `yaml:"steps"`
		} `yaml:"runs"`
	}
	if err := yaml.Unmarshal(data, &action); err != nil {
		t.Fatalf("parsing %s: %v", path, err)
	}
	if len(action.Runs.Steps) != 1 {
		t.Fatalf("expected exactly one step in the composite action, found %d", len(action.Runs.Steps))
	}

	script := action.Runs.Steps[0].Run
	if script == "" {
		t.Fatal("composite action step has an empty run body")
	}
	return script
}

func writeTestSpecBundle(t *testing.T, path string) {
	t.Helper()
	zipFile, err := os.Create(path) //nolint:gosec // isolated test temporary
	if err != nil {
		t.Fatal(err)
	}
	writer := zip.NewWriter(zipFile)
	entries := []struct {
		name string
		body string
	}{
		{"index.json", `{"specifications":[{"file":"domains/sites.json"}]}`},
		{"openapi.json", `{"openapi":"3.0.0","paths":{}}`},
		{"domains/sites.json", `{"openapi":"3.0.0","paths":{}}`},
	}
	for _, item := range entries {
		entry, createErr := writer.Create(item.name)
		if createErr != nil {
			t.Fatal(createErr)
		}
		if _, writeErr := entry.Write([]byte(item.body)); writeErr != nil {
			t.Fatal(writeErr)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if err := zipFile.Close(); err != nil {
		t.Fatal(err)
	}
}

func writeSymlinkSpecBundle(t *testing.T, path string) {
	t.Helper()
	zipFile, err := os.Create(path) //nolint:gosec // isolated test temporary
	if err != nil {
		t.Fatal(err)
	}
	writer := zip.NewWriter(zipFile)
	header := &zip.FileHeader{Name: "domains/leak.json", Method: zip.Store}
	header.SetMode(os.ModeSymlink | 0o777)
	entry, err := writer.CreateHeader(header)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := entry.Write([]byte("../../outside.json")); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if err := zipFile.Close(); err != nil {
		t.Fatal(err)
	}
}
