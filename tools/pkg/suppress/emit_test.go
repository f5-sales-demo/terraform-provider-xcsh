// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package suppress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

const (
	emptyDefaultsDatabase = `{"discovered":0,"failed":0,"generated_at":"2026-01-02T03:04:05Z","resources":[],"skipped":0,"total_resources":0,"version":"test"}`
	validDefaultsDatabase = `{"discovered":1,"failed":0,"generated_at":"2026-01-02T03:04:05Z","resources":[{"category":"simple","defaults":{"spec.flag":{"default_value":false,"path":"spec.flag","type":"bool"}},"resource_name":"probe","status":"discovered"}],"skipped":0,"total_resources":1,"version":"test"}`
)

func TestEmitImportSuppressionsRejectsMalformedExistingFileWithoutChangingIt(t *testing.T) {
	root := t.TempDir()
	databasePath := filepath.Join(root, "api-defaults.json")
	outputPath := filepath.Join(root, "import-default-suppressions.json")
	writeTestFile(t, databasePath, emptyDefaultsDatabase)
	original := []byte(`{"_comment":"measured","Measured":["keep"]`)
	if err := os.WriteFile(outputPath, original, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := EmitImportSuppressions(databasePath, outputPath); err == nil {
		t.Fatal("EmitImportSuppressions() succeeded with malformed existing data")
	}
	after, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(after, original) {
		t.Fatalf("malformed existing output changed: got %q, want %q", after, original)
	}
}

func TestEmitImportSuppressionsPreservesUnobservedMeasuredEntries(t *testing.T) {
	root := t.TempDir()
	databasePath := filepath.Join(root, "api-defaults.json")
	outputPath := filepath.Join(root, "import-default-suppressions.json")
	writeTestFile(t, databasePath, emptyDefaultsDatabase)
	writeTestFile(t, outputPath, `{"_comment":"measured","Measured":["keep"]}`)

	if err := EmitImportSuppressions(databasePath, outputPath); err != nil {
		t.Fatalf("EmitImportSuppressions() error = %v", err)
	}
	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]json.RawMessage
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("generated output is invalid JSON: %v", err)
	}
	var measured []string
	if err := json.Unmarshal(got["Measured"], &measured); err != nil {
		t.Fatalf("decode preserved measurement: %v", err)
	}
	if !reflect.DeepEqual(measured, []string{"keep"}) {
		t.Fatalf("Measured = %v, want [keep]", measured)
	}
}

func TestEmitImportSuppressionsRejectsDuplicateMeasuredKeysWithoutChangingFile(t *testing.T) {
	root := t.TempDir()
	databasePath := filepath.Join(root, "api-defaults.json")
	outputPath := filepath.Join(root, "import-default-suppressions.json")
	writeTestFile(t, databasePath, emptyDefaultsDatabase)
	original := []byte(`{"_comment":"measured","Measured":["keep"],"Measured":["lost"]}`)
	if err := os.WriteFile(outputPath, original, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := EmitImportSuppressions(databasePath, outputPath); err == nil {
		t.Fatal("EmitImportSuppressions() accepted duplicate measurement keys")
	}
	after, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(after, original) {
		t.Fatal("duplicate-key failure changed the existing measurement file")
	}
}

func TestEmitImportSuppressionsRejectsMalformedDefaultsDatabase(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(map[string]interface{})
	}{
		{name: "missing resources", mutate: func(db map[string]interface{}) { delete(db, "resources") }},
		{name: "blank version", mutate: func(db map[string]interface{}) { db["version"] = " \t" }},
		{name: "wrong generated_at type", mutate: func(db map[string]interface{}) { db["generated_at"] = 1 }},
		{name: "invalid generated_at", mutate: func(db map[string]interface{}) { db["generated_at"] = "yesterday" }},
		{name: "forbidden api endpoint", mutate: func(db map[string]interface{}) { db["api_endpoint"] = "https://example.invalid" }},
		{name: "null count", mutate: func(db map[string]interface{}) { db["discovered"] = nil }},
		{name: "fractional count", mutate: func(db map[string]interface{}) { db["discovered"] = 1.5 }},
		{name: "count mismatch", mutate: func(db map[string]interface{}) { db["discovered"] = 0 }},
		{name: "resource name whitespace", mutate: func(db map[string]interface{}) { resourceFixture(db)["resource_name"] = "bad name" }},
		{name: "duplicate resource name", mutate: func(db map[string]interface{}) {
			db["resources"] = append(db["resources"].([]interface{}), map[string]interface{}{
				"category": "simple", "resource_name": "probe", "status": "discovered",
			})
		}},
		{name: "blank category", mutate: func(db map[string]interface{}) { resourceFixture(db)["category"] = " " }},
		{name: "wrong category type", mutate: func(db map[string]interface{}) { resourceFixture(db)["category"] = 7 }},
		{name: "wrong status type", mutate: func(db map[string]interface{}) { resourceFixture(db)["status"] = true }},
		{name: "forbidden discovered timestamp", mutate: func(db map[string]interface{}) { resourceFixture(db)["discovered_at"] = "2026-01-02T03:04:05Z" }},
		{name: "forbidden request capture", mutate: func(db map[string]interface{}) {
			resourceFixture(db)["request_sent"] = map[string]interface{}{"metadata": map[string]interface{}{}}
		}},
		{name: "forbidden response capture", mutate: func(db map[string]interface{}) {
			resourceFixture(db)["response_got"] = map[string]interface{}{"metadata": map[string]interface{}{}}
		}},
		{name: "malformed default path", mutate: func(db map[string]interface{}) {
			resource := resourceFixture(db)
			field := resource["defaults"].(map[string]interface{})["spec.flag"]
			delete(resource["defaults"].(map[string]interface{}), "spec.flag")
			field.(map[string]interface{})["path"] = "spec.bad-name"
			resource["defaults"].(map[string]interface{})["spec.bad-name"] = field
		}},
		{name: "default path whitespace", mutate: func(db map[string]interface{}) {
			defaultFixture(db)["path"] = " spec.flag"
		}},
		{name: "default type mismatch", mutate: func(db map[string]interface{}) { defaultFixture(db)["default_value"] = "false" }},
		{name: "null marker", mutate: func(db map[string]interface{}) { defaultFixture(db)["is_marker_block"] = nil }},
		{name: "blank description", mutate: func(db map[string]interface{}) { defaultFixture(db)["description"] = " " }},
		{name: "forbidden failure detail", mutate: func(db map[string]interface{}) { resourceFixture(db)["error"] = "synthetic failure" }},
		{name: "defaults on failed resource", mutate: func(db map[string]interface{}) {
			resourceFixture(db)["status"] = "failed"
			db["discovered"] = float64(0)
			db["failed"] = float64(1)
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := t.TempDir()
			outputPath := filepath.Join(root, "import-default-suppressions.json")
			original := []byte(`{"_comment":"measured","Measured":["keep"]}`)
			writeTestFile(t, outputPath, string(original))
			database := decodeDatabaseFixture(t)
			test.mutate(database)
			databasePath := filepath.Join(root, "api-defaults.json")
			writeJSONTestFile(t, databasePath, database)
			if err := EmitImportSuppressions(databasePath, outputPath); err == nil {
				t.Fatal("EmitImportSuppressions() accepted malformed defaults data")
			}
			after, err := os.ReadFile(outputPath)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(after, original) {
				t.Fatal("malformed defaults data changed the existing output")
			}
		})
	}
}

func TestEmitImportSuppressionsAcceptsEveryDiscoveryStatus(t *testing.T) {
	for _, test := range []struct {
		name   string
		mutate func(map[string]interface{})
	}{
		{name: "discovered", mutate: func(map[string]interface{}) {}},
		{name: "failed", mutate: makeStatusMutation("failed")},
		{name: "skipped", mutate: makeStatusMutation("skipped")},
	} {
		t.Run(test.name, func(t *testing.T) {
			root := t.TempDir()
			databasePath := filepath.Join(root, "api-defaults.json")
			outputPath := filepath.Join(root, "import-default-suppressions.json")
			database := decodeDatabaseFixture(t)
			test.mutate(database)
			writeJSONTestFile(t, databasePath, database)
			writeTestFile(t, outputPath, `{"_comment":"measured","Seed":["keep"]}`)
			if err := EmitImportSuppressions(databasePath, outputPath); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestParseCanonicalSuppressionsRejectsMalformedIdentities(t *testing.T) {
	for name, data := range map[string]string{
		"resource whitespace":  `{"_comment":"measured","Measured Name":["keep"]}`,
		"resource punctuation": `{"_comment":"measured","Measured-Name":["keep"]}`,
		"resource lowercase":   `{"_comment":"measured","measured":["keep"]}`,
		"member whitespace":    `{"_comment":"measured","Measured":["keep member"]}`,
		"member punctuation":   `{"_comment":"measured","Measured":["keep.member"]}`,
		"member uppercase":     `{"_comment":"measured","Measured":["Keep"]}`,
		"member double join":   `{"_comment":"measured","Measured":["keep__member"]}`,
	} {
		t.Run(name, func(t *testing.T) {
			if _, _, err := ParseCanonicalSuppressions([]byte(data), "test.json"); err == nil {
				t.Fatal("ParseCanonicalSuppressions() accepted a malformed identity")
			}
		})
	}
}

func TestReadExistingSuppressionsRejectsCommentOnlyAndNonArrayData(t *testing.T) {
	for name, data := range map[string]string{
		"comment only": `{"_comment":"measured"}`,
		"non-array":    `{"_comment":"measured","Measured":"keep"}`,
	} {
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "import-default-suppressions.json")
			writeTestFile(t, path, data)
			if _, _, err := readExistingSuppressions(path); err == nil {
				t.Fatal("readExistingSuppressions() accepted invalid canonical data")
			}
		})
	}
}

func TestRepositoryDefaultsDatabasePassesStrictValidation(t *testing.T) {
	path := filepath.Join("..", "..", "api-defaults.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := rejectDuplicateJSONKeys(data, path); err != nil {
		t.Fatal(err)
	}
	var database Database
	if err := json.Unmarshal(data, &database); err != nil {
		t.Fatal(err)
	}
	if err := validateDefaultsDatabase(data, database, path); err != nil {
		t.Fatal(err)
	}
	t.Logf("strictly validated %d defaults database resource records", len(database.Resources))
}

func TestRepositoryCanonicalSuppressionsPassStrictValidation(t *testing.T) {
	path := filepath.Join("..", "..", "import-default-suppressions.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	suppressions, _, err := ParseCanonicalSuppressions(data, path)
	if err != nil {
		t.Fatal(err)
	}
	members := 0
	for _, resourceMembers := range suppressions {
		members += len(resourceMembers)
	}
	t.Logf("strictly validated %d canonical resources and %d members", len(suppressions), members)
}

func TestEmitImportSuppressionsCrossProcessMerge(t *testing.T) {
	if os.Getenv("XCSH_SUPPRESS_EMIT_HELPER") == "1" {
		writeTestFile(t, os.Getenv("XCSH_SUPPRESS_READY"), "ready")
		waitForPath(t, os.Getenv("XCSH_SUPPRESS_RELEASE"), 10*time.Second)
		if err := EmitImportSuppressions(os.Getenv("XCSH_SUPPRESS_DB"), os.Getenv("XCSH_SUPPRESS_OUT")); err != nil {
			t.Fatal(err)
		}
		return
	}
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" && runtime.GOOS != "windows" {
		t.Skipf("cross-process file locking is unsupported on %s", runtime.GOOS)
	}

	root := t.TempDir()
	outputPath := filepath.Join(root, "import-default-suppressions.json")
	releasePath := filepath.Join(root, "release")
	writeTestFile(t, outputPath, `{"_comment":"measured","Seed":["keep"]}`)

	const processCount = 8
	type childProcess struct {
		cmd    *exec.Cmd
		output bytes.Buffer
	}
	children := make([]childProcess, processCount)
	for i := range children {
		database := decodeDatabaseFixture(t)
		resource := resourceFixture(database)
		resourceName := fmt.Sprintf("probe_%d", i)
		resource["resource_name"] = resourceName
		databasePath := filepath.Join(root, fmt.Sprintf("database-%d.json", i))
		readyPath := filepath.Join(root, fmt.Sprintf("ready-%d", i))
		writeJSONTestFile(t, databasePath, database)

		cmd := exec.Command(os.Args[0], "-test.run=^TestEmitImportSuppressionsCrossProcessMerge$")
		cmd.Env = append(sanitizedSubprocessEnvironment(),
			"XCSH_SUPPRESS_EMIT_HELPER=1",
			"XCSH_SUPPRESS_DB="+databasePath,
			"XCSH_SUPPRESS_OUT="+outputPath,
			"XCSH_SUPPRESS_READY="+readyPath,
			"XCSH_SUPPRESS_RELEASE="+releasePath,
		)
		cmd.Stdout = &children[i].output
		cmd.Stderr = &children[i].output
		children[i].cmd = cmd
		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}
	}
	for i := range children {
		waitForPath(t, filepath.Join(root, fmt.Sprintf("ready-%d", i)), 10*time.Second)
	}
	writeTestFile(t, releasePath, "release")
	for i := range children {
		if err := children[i].cmd.Wait(); err != nil {
			t.Fatalf("child process %d failed: %v\n%s", i, err, children[i].output.String())
		}
	}

	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	merged, _, err := ParseCanonicalSuppressions(data, outputPath)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(merged), processCount+1; got != want {
		t.Fatalf("merged resource count = %d, want %d", got, want)
	}
	if !reflect.DeepEqual(merged["Seed"], []string{"keep"}) {
		t.Fatal("cross-process merge lost the pre-existing measurement")
	}
	for i := 0; i < processCount; i++ {
		resource := fmt.Sprintf("Probe%d", i)
		if !reflect.DeepEqual(merged[resource], []string{"flag"}) {
			t.Fatalf("cross-process merge lost synthetic resource %d", i)
		}
	}
}

func decodeDatabaseFixture(t *testing.T) map[string]interface{} {
	t.Helper()
	var database map[string]interface{}
	if err := json.Unmarshal([]byte(validDefaultsDatabase), &database); err != nil {
		t.Fatal(err)
	}
	return database
}

func resourceFixture(database map[string]interface{}) map[string]interface{} {
	return database["resources"].([]interface{})[0].(map[string]interface{})
}

func defaultFixture(database map[string]interface{}) map[string]interface{} {
	return resourceFixture(database)["defaults"].(map[string]interface{})["spec.flag"].(map[string]interface{})
}

func makeStatusMutation(status string) func(map[string]interface{}) {
	return func(database map[string]interface{}) {
		resource := resourceFixture(database)
		resource["status"] = status
		delete(resource, "defaults")
		database["discovered"] = float64(0)
		database[status] = float64(1)
	}
}

func sanitizedSubprocessEnvironment() []string {
	var environment []string
	for _, entry := range os.Environ() {
		key := strings.SplitN(entry, "=", 2)[0]
		if strings.HasPrefix(key, "XCSH_") || strings.HasPrefix(key, "VES_") || strings.HasPrefix(key, "LITELLM_") {
			continue
		}
		environment = append(environment, entry)
	}
	return environment
}

func writeJSONTestFile(t *testing.T, path string, value interface{}) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatal(err)
	}
}

func waitForPath(t *testing.T, path string, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(path); err == nil {
			return
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("timed out waiting for subprocess coordination file")
}

func writeTestFile(t *testing.T, path, data string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(data), 0o600); err != nil {
		t.Fatal(err)
	}
}
