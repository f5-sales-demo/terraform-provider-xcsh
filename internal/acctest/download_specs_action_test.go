// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package acctest

import (
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
		"ENRICHED_REPO=f5-sales-demo/api-specs-enriched",
		"GH_TOKEN=stub",
	)
	out, err := cmd.CombinedOutput()

	if _, statErr := os.Stat(marker); statErr != nil {
		t.Fatalf("action exited before attempting the download against an absent %s.\n"+
			"The probe must not kill the step on a clean checkout.\nerr: %v\noutput:\n%s",
			absent, err, out)
	}

	if _, statErr := os.Stat(absent); statErr != nil {
		t.Errorf("action did not create %s before downloading into it: %v", absent, statErr)
	}
}

// A bundle already in the tree must short-circuit: re-downloading on every job is
// waste, and the local development flow depends on `make download-specs` output being
// honoured rather than silently replaced.
func TestDownloadSpecsActionSkipsWhenBundlePresent(t *testing.T) {
	script := extractActionScript(t)

	tmp := t.TempDir()
	specDir := filepath.Join(tmp, "specs")
	if err := os.MkdirAll(filepath.Join(specDir, "domains"), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(specDir, "index.json"), []byte(`{}`), 0o600); err != nil {
		t.Fatal(err)
	}

	binDir := filepath.Join(tmp, "bin")
	if err := os.Mkdir(binDir, 0o750); err != nil {
		t.Fatal(err)
	}
	// Any invocation of gh is a failure here: the bundle is already present.
	stub := "#!/usr/bin/env bash\necho 'gh must not be called when specs are present' >&2\nexit 1\n"
	if err := os.WriteFile(filepath.Join(binDir, "gh"), []byte(stub), 0o700); err != nil { //nolint:gosec // test stub must be executable
		t.Fatal(err)
	}

	cmd := exec.Command("bash", "--noprofile", "--norc", "-eo", "pipefail", "-c", script)
	cmd.Dir = tmp
	cmd.Env = append(os.Environ(),
		"PATH="+binDir+string(os.PathListSeparator)+os.Getenv("PATH"),
		"SPEC_DIR="+specDir,
		"ENRICHED_REPO=f5-sales-demo/api-specs-enriched",
		"GH_TOKEN=stub",
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("action failed on a spec directory that already holds a bundle: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "existing spec files") {
		t.Errorf("expected the short-circuit message, got:\n%s", out)
	}
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
