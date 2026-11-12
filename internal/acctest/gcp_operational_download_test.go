package acctest

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

func TestDownloadSpecsActionGCPOperationalPair(t *testing.T) {
	for _, execution := range []string{"action", "sync"} {
		for _, kind := range []string{"legacy", "complete", "missing-health", "missing-qualification", "missing-release-half", "bad-digest", "tampered-standalone", "tampered-bundled", "duplicate-release-asset"} {
			t.Run(execution+"/"+kind, func(t *testing.T) {
				root := t.TempDir()
				tag := "v2.1.208"
				commit := strings.Repeat("a", 40)
				assets := writeTestSMSv2Assets(t, root, tag, commit)
				content := map[string]string{
					"api-catalog.json":             `{"version":"2.1.208","commands":[]}`,
					"index.json":                   `{"specifications":[{"file":"domains/sites.json"}]}`,
					"minimal-export-defaults.json": `{"version":"2.1.208","defaults":{}}`,
					"openapi.json":                 testSMSv2OpenAPI,
				}
				if kind != "legacy" {
					content["gcp-operational-health-v2.json"] = `{"schema_version":1,"qualification":"unverified","full_ce_acceptance":false}`
					content["gcp-operational-qualification-v2.json"] = `{"schema_version":1,"live_acceptance":false}`
				}
				if kind == "missing-health" {
					delete(content, "gcp-operational-health-v2.json")
				}
				if kind == "missing-qualification" {
					delete(content, "gcp-operational-qualification-v2.json")
				}
				for name, value := range content {
					if err := os.WriteFile(filepath.Join(root, name), []byte(value), 0600); err != nil {
						t.Fatal(err)
					}
					assets[name] = fileSHA256(t, filepath.Join(root, name))
				}
				bundleName := "f5xc-api-specs-" + tag + ".zip"
				bundle, err := os.Create(filepath.Join(root, bundleName))
				if err != nil {
					t.Fatal(err)
				}
				archive := zip.NewWriter(bundle)
				bundled := map[string]string{"index.json": content["index.json"], "domains/sites.json": `{"paths":{}}`}
				for _, name := range []string{"gcp-operational-health-v2.json", "gcp-operational-qualification-v2.json"} {
					if value, ok := content[name]; ok {
						bundled[name] = value
					}
				}
				if kind == "tampered-bundled" {
					bundled["gcp-operational-health-v2.json"] = `{"unreviewed":true}`
				}
				for name, value := range bundled {
					entry, err := archive.Create(name)
					if err != nil {
						t.Fatal(err)
					}
					if _, err = entry.Write([]byte(value)); err != nil {
						t.Fatal(err)
					}
				}
				if err = archive.Close(); err != nil {
					t.Fatal(err)
				}
				if err = bundle.Close(); err != nil {
					t.Fatal(err)
				}
				assets[bundleName] = fileSHA256(t, filepath.Join(root, bundleName))
				if kind == "bad-digest" {
					assets["gcp-operational-health-v2.json"] = "invalid"
				}
				writeSpecReleasePin(t, root, tag, commit, assets)
				names := make([]string, 0, len(assets))
				for name := range assets {
					names = append(names, name)
				}
				sort.Strings(names)
				releaseAssets := []map[string]any{}
				files := map[string]string{}
				for i, name := range names {
					id := i + 1
					files[fmt.Sprint(id)] = name
					if kind == "missing-release-half" && name == "gcp-operational-health-v2.json" {
						continue
					}
					releaseAssets = append(releaseAssets, map[string]any{"id": id, "name": name, "digest": "sha256:" + assets[name]})
				}
				if kind == "duplicate-release-asset" {
					releaseAssets = append(releaseAssets, releaseAssets[0])
				}
				receipt := map[string]any{"assets": qualifiedAssetDigests(assets), "commit": commit, "version": "2.1.208"}
				receiptJSON, err := json.Marshal(receipt)
				if err != nil {
					t.Fatal(err)
				}
				release := map[string]any{"tag_name": tag, "draft": false, "prerelease": false, "immutable": true, "assets": releaseAssets, "body": "<!-- publication-receipt:" + string(receiptJSON) + " -->"}
				for name, value := range map[string]any{"release.json": release, "asset-files.json": files, "commit.json": map[string]string{"sha": commit}} {
					raw, err := json.Marshal(value)
					if err != nil {
						t.Fatal(err)
					}
					if err = os.WriteFile(filepath.Join(root, name), raw, 0600); err != nil {
						t.Fatal(err)
					}
				}
				if kind == "tampered-standalone" {
					if err = os.WriteFile(filepath.Join(root, "gcp-operational-health-v2.json"), []byte(`{"tampered":true}`), 0600); err != nil {
						t.Fatal(err)
					}
				}
				validator, err := os.ReadFile(filepath.Join(testRepositoryRoot(t), "scripts", "validate-smsv2-release.py"))
				if err != nil {
					t.Fatal(err)
				}
				if err = os.WriteFile(filepath.Join(root, "scripts", "validate-smsv2-release.py"), validator, 0600); err != nil {
					t.Fatal(err)
				}
				bin := filepath.Join(root, "bin")
				if err = os.Mkdir(bin, 0700); err != nil {
					t.Fatal(err)
				}
				fixture := `import json,os,pathlib,sys
root=pathlib.Path(os.environ["FIXTURE_ROOT"])
endpoint=next((arg for arg in sys.argv[1:] if arg.startswith("repos/")), "")
with (root/"gh.log").open("a") as log: log.write(endpoint+"\n")
if "/releases/tags/" in endpoint: target="release.json"
elif "/commits/" in endpoint:
 if "--jq" in sys.argv:
  sys.stdout.write(json.loads((root/"commit.json").read_text())["sha"]);raise SystemExit(0)
 target="commit.json"
elif "/releases/assets/" in endpoint: target=json.loads((root/"asset-files.json").read_text())[endpoint.rsplit("/",1)[1]]
else: raise SystemExit(9)
sys.stdout.buffer.write((root/target).read_bytes())
`
				if err = os.WriteFile(filepath.Join(root, "gh-fixture.py"), []byte(fixture), 0600); err != nil {
					t.Fatal(err)
				}
				if err = os.WriteFile(filepath.Join(bin, "gh"), []byte("#!/bin/sh\nexec python3 \"$FIXTURE_ROOT/gh-fixture.py\" \"$@\"\n"), 0700); err != nil {
					t.Fatal(err)
				}
				destination := filepath.Join(root, "specs")
				if err = os.Mkdir(destination, 0700); err != nil {
					t.Fatal(err)
				}
				sentinel := filepath.Join(destination, "original-sentinel")
				if err = os.WriteFile(sentinel, []byte("preserved"), 0600); err != nil {
					t.Fatal(err)
				}
				script := extractActionScript(t)
				if execution == "sync" {
					script = extractSyncOpenAPIStep(t, "Download OpenAPI specs")
				}
				command := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
				command.Dir = root
				command.Env = append(os.Environ(), "PATH="+bin+string(os.PathListSeparator)+os.Getenv("PATH"), "SPEC_DIR="+destination, "SPEC_RELEASE_TAG="+tag, "ENRICHED_REPO=f5-sales-demo/api-specs-enriched", "GH_TOKEN=stub", "GITHUB_WORKSPACE="+root, "FIXTURE_ROOT="+root, "RUNNER_TEMP="+root, "SPEC_VERSION="+tag, "DISPATCH_TARGET_COMMIT="+commit)
				output, err := command.CombinedOutput()
				success := kind == "legacy" || kind == "complete"
				if (err == nil) != success {
					t.Fatalf("action outcome for %s: %v\n%s", kind, err, output)
				}
				if !success {
					if raw, err := os.ReadFile(sentinel); err != nil || string(raw) != "preserved" {
						t.Fatal("rejected download mutated original destination")
					}
					return
				}
				if kind == "complete" {
					for _, name := range []string{"gcp-operational-health-v2.json", "gcp-operational-qualification-v2.json"} {
						raw, err := os.ReadFile(filepath.Join(destination, name))
						if err != nil || string(raw) != content[name] {
							t.Fatalf("qualified companion missing or changed: %s", name)
						}
					}
				}
			})
		}
	}
}
