//go:build linux

package provider

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func captureFixture(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	dir := filepath.Join(root, "receipts")
	if err := os.Mkdir(dir, 0700); err != nil {
		t.Fatal(err)
	}
	process, err := originalCreationProcess()
	if err != nil {
		t.Fatal(err)
	}
	scope := OriginalCreationScope{ProviderExecutableSHA256: process.ExecutableSHA256, SchemaVersion: 1, ScopeSHA256: strings.Repeat("a", 64), Directory: dir, Resources: []OriginalCreationTarget{{TypeName: "xcsh_token", Name: "example-token", Namespace: "system"}}}
	raw, err := json.Marshal(scope)
	if err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(root, "scope.json")
	if err = os.WriteFile(file, raw, 0400); err != nil {
		t.Fatal(err)
	}
	t.Setenv("XCSH_ORIGINAL_CREATION_SCOPE_FILE", file)
	return file, dir
}

func TestOriginalCreationCaptureDisabled(t *testing.T) {
	t.Setenv("XCSH_ORIGINAL_CREATION_SCOPE_FILE", "")
	capture, err := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
	if err != nil || capture != nil {
		t.Fatalf("disabled capture changed legacy behavior: %v", err)
	}
}

func TestOriginalCreationCaptureSafeReceipt(t *testing.T) {
	_, dir := captureFixture(t)
	capture, err := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
	if err != nil {
		t.Fatal(err)
	}
	if capture == nil {
		t.Fatal("original capture not admitted")
	}
	receipt, err := capture.RecordCreation(strings.Repeat("b", 64))
	if err != nil {
		t.Fatal(err)
	}
	if receipt.SchemaVersion != 1 || receipt.ScopeSHA256 != strings.Repeat("a", 64) || receipt.CreationIdentitySHA256 != strings.Repeat("b", 64) || receipt.Process.PID != os.Getpid() {
		t.Fatal("original source projection differs")
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	found := 0
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), "receipt-") {
			continue
		}
		found++
		file := filepath.Join(dir, entry.Name())
		info, err := os.Lstat(file)
		if err != nil {
			t.Fatal(err)
		}
		if info.Mode().Perm() != 0400 {
			t.Fatal("receipt is not private immutable output")
		}
		raw, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		for _, secret := range []string{"jwt", "api_token", "cloud_init", "raw_state"} {
			if strings.Contains(string(raw), secret) {
				t.Fatal("private material exported")
			}
		}
	}
	if found != 1 {
		t.Fatal("one original receipt required")
	}
	if _, err = capture.RecordCreation(strings.Repeat("b", 64)); err == nil {
		t.Fatal("original creation receipt replay admitted")
	}
}

func TestOriginalCreationCaptureWrongTargetAndChangedScope(t *testing.T) {
	file, _ := captureFixture(t)
	if _, err := OpenOriginalCreationCapture("xcsh_token", "system", "foreign-token"); err == nil {
		t.Fatal("foreign target admitted")
	}
	capture, err := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(file, 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(file, []byte("{}"), 0400); err != nil {
		t.Fatal(err)
	}
	if _, err = capture.RecordCreation(strings.Repeat("b", 64)); err == nil {
		t.Fatal("changed original scope admitted")
	}
}

func TestOriginalCreationCaptureUnsafeInputs(t *testing.T) {
	for _, kind := range []string{"public-scope", "public-sink", "symlink-scope", "symlink-sink", "unknown-field", "trailing-json", "duplicate-json-key", "duplicate-target", "wrong-executable", "relative-sink"} {
		t.Run(kind, func(t *testing.T) {
			file, dir := captureFixture(t)
			switch kind {
			case "public-scope":
				os.Chmod(file, 0644)
			case "public-sink":
				os.Chmod(dir, 0755)
			case "symlink-scope":
				os.Rename(file, file+".original")
				os.Symlink(file+".original", file)
			case "symlink-sink":
				os.Rename(dir, dir+".original")
				os.Symlink(dir+".original", dir)
			default:
				raw, _ := os.ReadFile(file)
				var scope OriginalCreationScope
				json.Unmarshal(raw, &scope)
				switch kind {
				case "duplicate-target":
					scope.Resources = append(scope.Resources, scope.Resources[0])
				case "wrong-executable":
					scope.ProviderExecutableSHA256 = strings.Repeat("c", 64)
				case "relative-sink":
					scope.Directory = "receipts"
				}
				switch kind {
				case "duplicate-json-key":
					raw = append(raw[:len(raw)-1], []byte(",\"schema_version\":1}")...)
				case "unknown-field":
					raw = append(raw[:len(raw)-1], []byte(",\"unexpected\":true}")...)
				case "trailing-json":
					raw = append(raw, []byte("{}")...)
				default:
					raw, _ = json.Marshal(scope)
				}
				os.Chmod(file, 0600)
				os.WriteFile(file, raw, 0600)
				os.Chmod(file, 0400)
			}
			if _, err := OpenOriginalCreationCapture("xcsh_token", "system", "example-token"); err == nil {
				t.Fatal("unsafe capture admitted")
			}
		})
	}
}

func TestOriginalCreationCaptureInterruptedClaimCannotReplay(t *testing.T) {
	captureFixture(t)
	first, err := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
	if err != nil {
		t.Fatal(err)
	}
	first.Close()
	if _, err := OpenOriginalCreationCapture("xcsh_token", "system", "example-token"); err == nil {
		t.Fatal("interrupted create replay admitted")
	}
}

func TestOriginalCreationCaptureMissingUIDAndSinkReplacement(t *testing.T) {
	for _, kind := range []string{"missing-uid", "malformed-uid", "sink-replaced", "sink-permissions"} {
		t.Run(kind, func(t *testing.T) {
			_, dir := captureFixture(t)
			capture, err := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
			if err != nil {
				t.Fatal(err)
			}
			defer capture.Close()
			hash := strings.Repeat("b", 64)
			switch kind {
			case "missing-uid":
				hash = ""
			case "malformed-uid":
				hash = "credential-value"
			case "sink-replaced":
				os.Rename(dir, dir+".original")
				os.Mkdir(dir, 0700)
			case "sink-permissions":
				os.Chmod(dir, 0755)
			}
			if _, err := capture.RecordCreation(hash); err == nil {
				t.Fatal("invalid capture completed")
			}
		})
	}
}

func TestOriginalCreationCaptureConcurrentSingleClaim(t *testing.T) {
	captureFixture(t)
	results := make(chan bool, 16)
	for i := 0; i < 16; i++ {
		go func() {
			c, e := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
			if c != nil {
				defer c.Close()
			}
			results <- e == nil
		}()
	}
	admitted := 0
	for i := 0; i < 16; i++ {
		if <-results {
			admitted++
		}
	}
	if admitted != 1 {
		t.Fatalf("got %d admitted concurrent creates", admitted)
	}
}
