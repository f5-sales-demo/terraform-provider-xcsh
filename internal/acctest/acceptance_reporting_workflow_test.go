// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAcceptanceMockRetainsEvidenceBeforeIncompleteStreamGate(t *testing.T) {
	script := extractWorkflowRunStep(t, "acc-tests.yml", "mock-tests", "Run mock tests")
	dir := t.TempDir()
	bin := filepath.Join(dir, "bin")
	if err := os.Mkdir(bin, 0700); err != nil {
		t.Fatal(err)
	}
	stub := `#!/usr/bin/env bash
if [[ "$1" == "test" ]]; then
  printf '%s\n' "$@" > arguments
  echo '{"Action":"start","Package":"example.invalid/provider"}'
echo COMPILER_SENTINEL >&2
  exit 1
fi
for arg in "$@"; do
  case "$arg" in
    -output=*) target="${arg#-output=}" ;;
    -format=*) format="${arg#-format=}" ;;
  esac
done
cat >/dev/null
case "$format" in
  json) echo '{"stream_completeness":{"complete":false},"total_tests":0,"total_passed":0,"total_failed":0}' > "$target" ;;
  markdown) echo 'Stream complete: false' > "$target" ;;
  junit) echo '<testsuites/>' > "$target" ;;
esac
exit 1
`
	if err := os.WriteFile(filepath.Join(bin, "go"), []byte(stub), 0700); err != nil {
		t.Fatal(err)
	}
	_, err := runWorkflowScript(dir, script, []string{
		"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH"),
		"GITHUB_OUTPUT=" + filepath.Join(dir, "outputs"),
		"INPUT_TEST_PATTERN=^TestAccFixture$",
	})
	if err == nil {
		t.Fatal("incomplete stream passed")
	}
	for _, name := range []string{"mock-tests.json", "mock-tests.md"} {
		if data, err := os.ReadFile(filepath.Join(dir, "test-reports", name)); err != nil || len(data) == 0 {
			t.Errorf("missing failed-run evidence %s", name)
		}
	}
	stream, err := os.ReadFile(filepath.Join(dir, "test-reports", "test-output-mock.json"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(stream), "COMPILER_SENTINEL") {
		t.Error("compiler stderr contaminated JSON stream")
	}
	args, err := os.ReadFile(filepath.Join(dir, "arguments"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(args), "^TestAccFixture$") {
		t.Error("focused input was not transported to go test")
	}
}

func TestAcceptanceLiveConcurrencyDoesNotDependOnBranch(t *testing.T) {
	data, err := os.ReadFile("../../.github/workflows/acc-tests.yml")
	if err != nil {
		t.Fatal(err)
	}
	source := string(data)
	start := strings.Index(source, "  real-api-tests:")
	end := strings.Index(source[start:], "    steps:")
	job := source[start : start+end]
	if !strings.Contains(job, "group: acceptance-live-tenant") || !strings.Contains(job, "cancel-in-progress: false") {
		t.Fatal("tenant live mutation jobs can overlap across branches")
	}
}
