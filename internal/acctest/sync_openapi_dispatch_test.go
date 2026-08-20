// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
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

const (
	dispatchSource = "f5-sales-demo/api-specs-enriched"
	dispatchTarget = "f5-sales-demo/terraform-provider-xcsh"
)

func TestSyncOpenAPIDispatchDeliveryIsIdempotent(t *testing.T) {
	script := extractSyncOpenAPIStep(t, "Validate dispatched release and delivery identity")
	recordScript := extractSyncOpenAPIStep(t, "Record resolved spec version (tracked marker)")
	tmp := t.TempDir()
	ledger := filepath.Join(tmp, "deliveries.json")
	pending := filepath.Join(tmp, "pending.json")
	versionFile := filepath.Join(tmp, "spec-version.txt")
	if err := os.WriteFile(ledger, []byte("{\"deliveries\":{},\"version\":1}\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(versionFile, []byte("v2.1.207\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	commit := strings.Repeat("a", 40)
	deliveryID := providerDeliveryID(t, "2.1.208", "v2.1.208", commit)
	output, runOutput, err := runDispatchContract(
		t, script, tmp, ledger, pending, versionFile, deliveryID, "2.1.208", "v2.1.208", commit,
	)
	if err != nil {
		t.Fatalf("new valid delivery was rejected: %v\n%s", err, runOutput)
	}
	for _, expected := range []string{
		"process=true",
		"delivery_id=" + deliveryID,
		"release_tag=v2.1.208",
		"version=2.1.208",
	} {
		if !strings.Contains(output, expected) {
			t.Fatalf("valid delivery output omitted %q:\n%s", expected, output)
		}
	}

	recordCmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", recordScript)
	recordCmd.Dir = tmp
	recordCmd.Env = append(os.Environ(),
		"RESOLVED_SPEC_VERSION=v2.1.208",
		"DELIVERY_ID="+deliveryID,
		"PENDING_DELIVERY="+pending,
		"RELEASE_VERSION=2.1.208",
		"SPEC_VERSION_FILE="+versionFile,
		"TARGET_COMMIT="+commit,
	)
	if recordOutput, recordErr := recordCmd.CombinedOutput(); recordErr != nil {
		t.Fatalf("recording valid delivery failed: %v\n%s", recordErr, recordOutput)
	}
	ledgerBytes, err := os.ReadFile(ledger) //nolint:gosec // isolated test path
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(ledgerBytes), deliveryID) {
		t.Fatal("delivery was receipted before provider publication completed")
	}
	pendingBytes, err := os.ReadFile(pending) //nolint:gosec // isolated test path
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(pendingBytes), deliveryID) {
		t.Fatal("pending delivery identity was not persisted for post-release receipt")
	}

	output, runOutput, err = runDispatchContract(
		t, script, tmp, ledger, pending, versionFile, deliveryID, "2.1.208", "v2.1.208", commit,
	)
	if err != nil {
		t.Fatalf("duplicate valid delivery failed instead of becoming a no-op: %v\n%s", err, runOutput)
	}
	if !strings.Contains(output, "process=false") || !strings.Contains(output, "resume=true") {
		t.Fatalf("duplicate pending delivery did not request publication reconciliation: %q", output)
	}
}

func TestSyncOpenAPIDispatchRejectsStaleAndForgedDelivery(t *testing.T) {
	script := extractSyncOpenAPIStep(t, "Validate dispatched release and delivery identity")
	tmp := t.TempDir()
	ledger := filepath.Join(tmp, "deliveries.json")
	pending := filepath.Join(tmp, "pending.json")
	versionFile := filepath.Join(tmp, "spec-version.txt")
	if err := os.WriteFile(ledger, []byte("{\"deliveries\":{},\"version\":1}\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(versionFile, []byte("v2.1.208\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	commit := strings.Repeat("b", 40)

	_, output, err := runDispatchContract(
		t, script, tmp, ledger, pending, versionFile, strings.Repeat("0", 64), "2.1.209", "v2.1.209", commit,
	)
	if err == nil || !strings.Contains(output, "delivery_id does not match") {
		t.Fatalf("forged delivery_id was not rejected explicitly: err=%v\n%s", err, output)
	}

	staleID := providerDeliveryID(t, "2.1.207", "v2.1.207", commit)
	_, output, err = runDispatchContract(
		t, script, tmp, ledger, pending, versionFile, staleID, "2.1.207", "v2.1.207", commit,
	)
	if err == nil || !strings.Contains(output, "stale dispatch") {
		t.Fatalf("stale release was not rejected explicitly: err=%v\n%s", err, output)
	}
}

func TestSyncOpenAPIDispatchDurableRetryIsExactNoOp(t *testing.T) {
	script := extractSyncOpenAPIStep(t, "Validate dispatched release and delivery identity")
	tmp := t.TempDir()
	ledger := filepath.Join(tmp, "deliveries.json")
	pending := filepath.Join(tmp, "pending.json")
	versionFile := filepath.Join(tmp, "spec-version.txt")
	commit := strings.Repeat("b", 40)
	version := "2.1.209"
	tag := "v" + version
	deliveryID := providerDeliveryID(t, version, tag, commit)
	writeReleaseTestJSON(t, ledger, map[string]any{
		"deliveries": map[string]any{
			deliveryID: map[string]string{
				"release_tag": tag, "target_commit": commit, "version": version,
			},
		},
		"version": 1,
	})
	if err := os.WriteFile(versionFile, []byte("v2.1.209\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	outputs, runOutput, err := runDispatchContract(
		t, script, tmp, ledger, pending, versionFile, deliveryID, version, tag, commit,
	)
	if err != nil {
		t.Fatalf("exact durable retry failed instead of becoming a no-op: %v\n%s", err, runOutput)
	}
	if !strings.Contains(outputs, "process=false") || strings.Contains(outputs, "process=true") || strings.Contains(outputs, "resume=true") {
		t.Fatalf("exact durable retry did not produce a terminal no-op: %q", outputs)
	}
}

func TestSyncOpenAPIDispatchRejectsReusedDurableIdentityDimensions(t *testing.T) {
	script := extractSyncOpenAPIStep(t, "Validate dispatched release and delivery identity")
	incomingCommit := strings.Repeat("b", 40)
	incomingVersion := "2.1.209"
	incomingTag := "v" + incomingVersion
	incomingID := providerDeliveryID(t, incomingVersion, incomingTag, incomingCommit)
	type mutation struct {
		key      string
		delivery map[string]string
		want     string
	}
	mutations := map[string]mutation{
		"conflicting delivery ID": {
			key: incomingID,
			delivery: map[string]string{
				"release_tag": "v2.1.208", "target_commit": strings.Repeat("a", 40), "version": "2.1.208",
			},
			want: "delivery ledger entry disagrees with the dispatched identity",
		},
		"reused release tag": {
			key: providerDeliveryID(t, incomingVersion, incomingTag, strings.Repeat("a", 40)),
			delivery: map[string]string{
				"release_tag": incomingTag, "target_commit": strings.Repeat("a", 40), "version": incomingVersion,
			},
			want: "release_tag " + incomingTag + " was previously delivered under a different identity",
		},
		"reused version": {
			key: providerDeliveryID(t, incomingVersion, "v2.1.208", strings.Repeat("a", 40)),
			delivery: map[string]string{
				"release_tag": "v2.1.208", "target_commit": strings.Repeat("a", 40), "version": incomingVersion,
			},
			want: "version " + incomingVersion + " was previously delivered under a different identity",
		},
		"reused target commit": {
			key: providerDeliveryID(t, "2.1.208", "v2.1.208", incomingCommit),
			delivery: map[string]string{
				"release_tag": "v2.1.208", "target_commit": incomingCommit, "version": "2.1.208",
			},
			want: "target_commit " + incomingCommit + " was previously delivered under a different identity",
		},
	}

	for name, mutation := range mutations {
		t.Run(name, func(t *testing.T) {
			tmp := t.TempDir()
			ledger := filepath.Join(tmp, "deliveries.json")
			pending := filepath.Join(tmp, "pending.json")
			versionFile := filepath.Join(tmp, "spec-version.txt")
			writeReleaseTestJSON(t, ledger, map[string]any{
				"deliveries": map[string]any{mutation.key: mutation.delivery},
				"version":    1,
			})
			if err := os.WriteFile(versionFile, []byte("v2.1.208\n"), 0o600); err != nil {
				t.Fatal(err)
			}

			outputs, runOutput, err := runDispatchContract(
				t, script, tmp, ledger, pending, versionFile,
				incomingID, incomingVersion, incomingTag, incomingCommit,
			)
			if err == nil || !strings.Contains(runOutput, mutation.want) {
				t.Fatalf("durable identity mutation was not rejected explicitly: err=%v\n%s", err, runOutput)
			}
			if strings.Contains(outputs, "process=true") {
				t.Fatalf("rejected durable identity emitted process authorization:\n%s", outputs)
			}
		})
	}
}

func TestSyncOpenAPIUsesDispatchedTagAndRetainsCatalog(t *testing.T) {
	determine := extractSyncOpenAPIStep(t, "Determine spec version")
	download := extractSyncOpenAPIStep(t, "Download OpenAPI specs")
	record := extractSyncOpenAPIStep(t, "Record resolved spec version (tracked marker)")

	if !strings.Contains(determine, `SPEC_VERSION="$DISPATCH_RELEASE_TAG"`) {
		t.Fatal("repository_dispatch does not replace the manual tag with release_tag")
	}
	if strings.Contains(download, ".[0].assets") ||
		!strings.Contains(download, `releases/tags/${SPEC_VERSION}`) {
		t.Fatal("download is not scoped to the resolved release tag")
	}
	if strings.Contains(determine, "releases/latest") || strings.Contains(determine, `"latest"`) {
		t.Fatal("version resolution still contains a mutable latest fallback")
	}
	catalogDownload := strings.Index(download, `for asset in api-catalog.json index.json minimal-export-defaults.json openapi.json`)
	promotion := strings.Index(download, `cp -a "${STAGED_SPECS}/." "${SPEC_DIR}/"`)
	if catalogDownload < 0 || promotion < 0 || catalogDownload > promotion {
		t.Fatal("api-catalog.json is not retained in the staged bundle before promotion")
	}
	if strings.Contains(record, ".deliveries[$id]") ||
		!strings.Contains(record, "{delivery_id:$id,release_tag:$tag,target_commit:$commit,version:$version}") {
		t.Fatal("sync must stage pending identity without writing the completed-delivery ledger")
	}
}

func TestSyncOpenAPIRetryPushUsesPATAndTriggersChecks(t *testing.T) {
	path := filepath.Join("..", "..", ".github", "workflows", "sync-openapi.yml")
	data, err := os.ReadFile(path) //nolint:gosec // fixed repository workflow
	if err != nil {
		t.Fatal(err)
	}
	var workflow struct {
		Jobs map[string]struct {
			Permissions map[string]string `yaml:"permissions"`
			Steps       []struct {
				Name string            `yaml:"name"`
				Uses string            `yaml:"uses"`
				With map[string]any    `yaml:"with"`
				Env  map[string]string `yaml:"env"`
				Run  string            `yaml:"run"`
			} `yaml:"steps"`
		} `yaml:"jobs"`
	}
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		t.Fatal(err)
	}
	syncJob := workflow.Jobs["sync"]
	if len(syncJob.Permissions) != 0 {
		t.Fatalf("sync GITHUB_TOKEN retains unused permissions: %+v", syncJob.Permissions)
	}

	var checkout, push *struct {
		Name string            `yaml:"name"`
		Uses string            `yaml:"uses"`
		With map[string]any    `yaml:"with"`
		Env  map[string]string `yaml:"env"`
		Run  string            `yaml:"run"`
	}
	for i := range syncJob.Steps {
		step := &syncJob.Steps[i]
		switch step.Name {
		case "Checkout":
			checkout = step
		case "Create branch and commit":
			push = step
		}
	}
	if checkout == nil || push == nil {
		t.Fatal("sync checkout or branch-push step is missing")
	}
	if checkout.With["persist-credentials"] != false ||
		checkout.With["token"] != "${{ secrets.REPO_SYNC_TOKEN }}" {
		t.Fatalf("checkout can persist or use the automatic token: %+v", checkout.With)
	}
	if push.Env["GH_TOKEN"] != "${{ secrets.REPO_SYNC_TOKEN }}" ||
		push.Env["REPOSITORY"] != "${{ github.repository }}" {
		t.Fatalf("branch push is not bound to the intended PAT/repository: %+v", push.Env)
	}
	for _, required := range []string{
		`git config --unset-all http.https://github.com/.extraheader`,
		`git remote set-url origin "https://x-access-token:${GH_TOKEN}@github.com/${REPOSITORY}.git"`,
		`git push --force-with-lease`,
	} {
		if !strings.Contains(push.Run, required) {
			t.Fatalf("branch push omitted required PAT wiring %q", required)
		}
	}
}

func TestSyncOpenAPIDownloadPromotesExactReleaseCatalog(t *testing.T) {
	script := extractSyncOpenAPIStep(t, "Download OpenAPI specs")
	tmp := t.TempDir()
	specDir := filepath.Join(tmp, "specs")
	if err := os.Mkdir(specDir, 0o750); err != nil {
		t.Fatal(err)
	}
	stale := filepath.Join(specDir, "stale.json")
	if err := os.WriteFile(stale, []byte(`{"stale":true}`), 0o600); err != nil {
		t.Fatal(err)
	}
	bundle := filepath.Join(tmp, "bundle.zip")
	writeTestSpecBundle(t, bundle)
	catalog := filepath.Join(tmp, "catalog.json")
	catalogBody := `{"version":"2.1.208","commands":[{"name":"vesctl"}]}`
	if err := os.WriteFile(catalog, []byte(catalogBody), 0o600); err != nil {
		t.Fatal(err)
	}
	index := filepath.Join(tmp, "index.json")
	indexBody := `{"specifications":[{"file":"domains/sites.json"}]}`
	if err := os.WriteFile(index, []byte(indexBody), 0o600); err != nil {
		t.Fatal(err)
	}
	openapi := filepath.Join(tmp, "openapi.json")
	openapiBody := `{"openapi":"3.0.0","paths":{}}`
	if err := os.WriteFile(openapi, []byte(openapiBody), 0o600); err != nil {
		t.Fatal(err)
	}
	minimal := filepath.Join(tmp, "minimal-export-defaults.json")
	writeTestSMSv2Assets(t, tmp, "v2.1.208", strings.Repeat("a", 40))
	validator, err := os.ReadFile(filepath.Join(testRepositoryRoot(t), "scripts", "validate-smsv2-release.py"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(tmp, "scripts"), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmp, "scripts", "validate-smsv2-release.py"), validator, 0o600); err != nil {
		t.Fatal(err)
	}
	minimalBody := `{"version":"2.1.208","defaults":{}}`
	if err := os.WriteFile(minimal, []byte(minimalBody), 0o600); err != nil {
		t.Fatal(err)
	}

	binDir := filepath.Join(tmp, "bin")
	if err := os.Mkdir(binDir, 0o750); err != nil {
		t.Fatal(err)
	}
	logPath := filepath.Join(tmp, "gh.log")
	stub := `#!/usr/bin/env bash
printf '%s\n' "$*" >> "$GH_LOG"
case "$*" in
  *releases/assets/301*) cat "$BUNDLE_ZIP" ;;
  *releases/assets/302*) cat "$CATALOG_FILE" ;;
  *releases/assets/303*) cat "$INDEX_FILE" ;;
  *releases/assets/304*) cat "$MINIMAL_FILE" ;;
  *releases/assets/305*) cat "$OPENAPI_FILE" ;;
  *releases/assets/306*) cat "$SMSV2_MANIFEST_FILE" ;;
  *releases/assets/307*) cat "$SMSV2_CONTRACT_FILE" ;;
  *releases/assets/308*) cat "$SMSV2_EVIDENCE_FILE" ;;
  *commits/v2.1.208*) printf '%s\n' "$TAG_COMMIT" ;;
  *releases/tags/v2.1.208*)
    bundle_sha=$(shasum -a 256 "$BUNDLE_ZIP" | awk '{print $1}')
    catalog_sha=$(shasum -a 256 "$CATALOG_FILE" | awk '{print $1}')
    index_sha=$(shasum -a 256 "$INDEX_FILE" | awk '{print $1}')
    minimal_sha=$(shasum -a 256 "$MINIMAL_FILE" | awk '{print $1}')
    openapi_sha=$(shasum -a 256 "$OPENAPI_FILE" | awk '{print $1}')
    manifest_sha=$(shasum -a 256 "$SMSV2_MANIFEST_FILE" | cut -d " " -f1)
    contract_sha=$(shasum -a 256 "$SMSV2_CONTRACT_FILE" | cut -d " " -f1)
    evidence_sha=$(shasum -a 256 "$SMSV2_EVIDENCE_FILE" | cut -d " " -f1)
    jq -cn \
      --arg bundle_sha "$bundle_sha" \
      --arg catalog_sha "$catalog_sha" \
      --arg commit "$TAG_COMMIT" \
      --arg index_sha "$index_sha" \
      --arg minimal_sha "$minimal_sha" \
      --arg openapi_sha "$openapi_sha" \
      --arg manifest_sha "$manifest_sha" \
      --arg contract_sha "$contract_sha" \
      --arg evidence_sha "$evidence_sha" '
      {tag_name: "v2.1.208", draft: false, prerelease: false, immutable: true, assets: [
        {id: 302, name: "api-catalog.json", digest: ("sha256:" + $catalog_sha)},
        {id: 301, name: "f5xc-api-specs-v2.1.208.zip", digest: ("sha256:" + $bundle_sha)},
        {id: 303, name: "index.json", digest: ("sha256:" + $index_sha)},
        {id: 304, name: "minimal-export-defaults.json", digest: ("sha256:" + $minimal_sha)},
        {id: 305, name: "openapi.json", digest: ("sha256:" + $openapi_sha)},
        {id: 306, name: "smsv2-contract-manifest.json", digest: ("sha256:" + $manifest_sha)},
        {id: 307, name: "smsv2-contract.json", digest: ("sha256:" + $contract_sha)},
        {id: 308, name: "smsv2-evidence-receipt.json", digest: ("sha256:" + $evidence_sha)}
      ], body: ("<!-- publication-receipt:" + ({
        assets: {
          "api-catalog.json": ("sha256:" + $catalog_sha),
          "f5xc-api-specs-v2.1.208.zip": ("sha256:" + $bundle_sha),
          "index.json": ("sha256:" + $index_sha),
          "minimal-export-defaults.json": ("sha256:" + $minimal_sha),
          "openapi.json": ("sha256:" + $openapi_sha),
          "smsv2-contract-manifest.json": ("sha256:" + $manifest_sha),
          "smsv2-contract.json": ("sha256:" + $contract_sha),
          "smsv2-evidence-receipt.json": ("sha256:" + $evidence_sha)
        }, commit: $commit, version: "2.1.208"
      } | tojson) + " -->")}' ;;
  *) echo "unexpected gh invocation: $*" >&2; exit 9 ;;
esac
`
	if err := os.WriteFile(filepath.Join(binDir, "gh"), []byte(stub), 0o700); err != nil { //nolint:gosec // test stub must be executable
		t.Fatal(err)
	}

	downloadEnv := []string{
		"PATH=" + binDir + string(os.PathListSeparator) + os.Getenv("PATH"),
		"SPEC_DIR=" + specDir,
		"SPEC_VERSION=v2.1.208",
		"ENRICHED_REPO=" + dispatchSource,
		"GH_TOKEN=stub",
		"GH_LOG=" + logPath,
		"BUNDLE_ZIP=" + bundle,
		"CATALOG_FILE=" + catalog,
		"INDEX_FILE=" + index,
		"MINIMAL_FILE=" + minimal,
		"OPENAPI_FILE=" + openapi,
		"SMSV2_MANIFEST_FILE=" + filepath.Join(tmp, "smsv2-contract-manifest.json"),
		"SMSV2_CONTRACT_FILE=" + filepath.Join(tmp, "smsv2-contract.json"),
		"SMSV2_EVIDENCE_FILE=" + filepath.Join(tmp, "smsv2-evidence-receipt.json"),
		"DISPATCH_TARGET_COMMIT=" + strings.Repeat("a", 40),
		"RUNNER_TEMP=" + filepath.Join(tmp, "runner"),
		"TAG_COMMIT=" + strings.Repeat("a", 40),
	}
	cmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = tmp
	cmd.Env = append(os.Environ(), downloadEnv...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("workflow download step failed: %v\n%s", err, output)
	}
	gotCatalog, err := os.ReadFile(filepath.Join(specDir, "api-catalog.json")) //nolint:gosec // isolated test path
	if err != nil {
		t.Fatalf("workflow promotion lost api-catalog.json: %v", err)
	}
	if string(gotCatalog) != catalogBody {
		t.Fatalf("workflow promoted the wrong catalog bytes: %s", gotCatalog)
	}
	pinBytes, err := os.ReadFile(filepath.Join(tmp, "tools", "spec-release.json")) //nolint:gosec // isolated test path
	if err != nil {
		t.Fatalf("workflow did not persist the verified release pin: %v", err)
	}
	var pin struct {
		Assets       map[string]string `json:"assets"`
		ReleaseTag   string            `json:"release_tag"`
		TargetCommit string            `json:"target_commit"`
		Version      string            `json:"version"`
	}
	if err := json.Unmarshal(pinBytes, &pin); err != nil {
		t.Fatalf("release pin is not JSON: %v", err)
	}
	// The pin records digests as "sha256:<hex>", the same form GitHub reports and the
	// receipt carries, so the three compare by equality with no prefix arithmetic.
	wantCatalogDigest := fmt.Sprintf("sha256:%x", sha256.Sum256([]byte(catalogBody)))
	if pin.ReleaseTag != "v2.1.208" || pin.Version != "2.1.208" ||
		pin.TargetCommit != strings.Repeat("a", 40) ||
		pin.Assets["api-catalog.json"] != wantCatalogDigest {
		t.Fatalf("release pin does not preserve the verified receipt identity: %+v", pin)
	}
	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Fatalf("stale pre-release content survived exact bundle promotion: %v", err)
	}
	logBytes, err := os.ReadFile(logPath) //nolint:gosec // isolated test path
	if err != nil {
		t.Fatal(err)
	}
	log := string(logBytes)
	if !strings.Contains(log, "releases/tags/v2.1.208") ||
		strings.Contains(log, "api repos/f5-sales-demo/api-specs-enriched/releases --jq") {
		t.Fatalf("workflow did not use only the exact release endpoint:\n%s", log)
	}

	writeSymlinkSpecBundle(t, bundle)
	sentinel := filepath.Join(specDir, "existing-bundle-must-survive")
	if err := os.WriteFile(sentinel, []byte("preserve\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	cmd = exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = tmp
	cmd.Env = append(os.Environ(), downloadEnv...)
	output, err = cmd.CombinedOutput()
	if err == nil || !strings.Contains(string(output), "unsupported archive entry type") {
		t.Fatalf("sync accepted a symlink-bearing spec archive: err=%v\n%s", err, output)
	}
	if _, err := os.Stat(sentinel); err != nil {
		t.Fatalf("rejected sync archive mutated the prior spec tree: %v", err)
	}
}

func TestSyncOpenAPIDownloadRejectsMutableRelease(t *testing.T) {
	script := extractSyncOpenAPIStep(t, "Download OpenAPI specs")
	tmp := t.TempDir()
	binDir := filepath.Join(tmp, "bin")
	if err := os.Mkdir(binDir, 0o750); err != nil {
		t.Fatal(err)
	}
	stub := `#!/usr/bin/env bash
printf '%s\n' '{"tag_name":"v2.1.208","draft":false,"prerelease":false,"immutable":false,"assets":[]}'
`
	if err := os.WriteFile(filepath.Join(binDir, "gh"), []byte(stub), 0o700); err != nil { //nolint:gosec // executable test stub
		t.Fatal(err)
	}
	cmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = tmp
	cmd.Env = append(os.Environ(),
		"PATH="+binDir+string(os.PathListSeparator)+os.Getenv("PATH"),
		"SPEC_DIR="+filepath.Join(tmp, "specs"),
		"SPEC_VERSION=v2.1.208",
		"ENRICHED_REPO="+dispatchSource,
		"GH_TOKEN=stub",
		"DISPATCH_TARGET_COMMIT="+strings.Repeat("a", 40),
		"RUNNER_TEMP="+filepath.Join(tmp, "runner"),
	)
	output, err := cmd.CombinedOutput()
	if err == nil || !strings.Contains(string(output), "final and immutable") {
		t.Fatalf("mutable spec release was accepted: err=%v\n%s", err, output)
	}
}

func providerDeliveryID(t *testing.T, version, tag, commit string) string {
	t.Helper()
	identity := map[string]string{
		"commit":     commit,
		"event_type": "enriched-specs-updated",
		"source":     dispatchSource,
		"tag":        tag,
		"target":     dispatchTarget,
		"version":    version,
	}
	canonical, err := json.Marshal(identity)
	if err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf("%x", sha256.Sum256(canonical))
}

func runDispatchContract(
	t *testing.T,
	script, dir, ledger, pending, versionFile, deliveryID, version, tag, commit string,
) (string, string, error) {
	t.Helper()
	outputFile := filepath.Join(dir, "github-output")
	if err := os.WriteFile(outputFile, nil, 0o600); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(),
		"EVENT_NAME=repository_dispatch",
		"DELIVERY_ID="+deliveryID,
		"RELEASE_TAG="+tag,
		"RELEASE_VERSION="+version,
		"TARGET_COMMIT="+commit,
		"TRIGGER_SOURCE="+dispatchSource,
		"DELIVERY_LEDGER="+ledger,
		"PENDING_DELIVERY="+pending,
		"CURRENT_SPEC_VERSION_FILE="+versionFile,
		"TARGET_REPOSITORY="+dispatchTarget,
		"ENRICHED_REPO="+dispatchSource,
		"GITHUB_OUTPUT="+outputFile,
	)
	runOutput, err := cmd.CombinedOutput()
	outputs, readErr := os.ReadFile(outputFile) //nolint:gosec // isolated test path
	if readErr != nil {
		t.Fatal(readErr)
	}
	return string(outputs), string(runOutput), err
}

func extractSyncOpenAPIStep(t *testing.T, name string) string {
	t.Helper()
	path := filepath.Join("..", "..", ".github", "workflows", "sync-openapi.yml")
	data, err := os.ReadFile(path) //nolint:gosec // fixed repo-relative path
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
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
		t.Fatalf("parsing %s: %v", path, err)
	}
	for _, step := range workflow.Jobs["sync"].Steps {
		if step.Name == name {
			if step.Run == "" {
				t.Fatalf("workflow step %q has no run body", name)
			}
			return step.Run
		}
	}
	t.Fatalf("workflow step %q not found", name)
	return ""
}
