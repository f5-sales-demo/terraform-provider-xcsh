//go:build linux

package provider

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOriginalCreationReceiptNeverOverwritesExistingFile(t *testing.T) {
	_, dir := captureFixture(t)
	c, err := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	file := filepath.Join(dir, "receipt-"+c.receipt.Nonce+".json")
	original := []byte("preserved-incomplete-capture")
	if err = os.WriteFile(file, original, 0400); err != nil {
		t.Fatal(err)
	}
	if _, err = c.RecordCreation(strings.Repeat("b", 64)); err == nil {
		t.Fatal("existing receipt overwritten")
	}
	raw, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != string(original) {
		t.Fatal("prior evidence changed")
	}
}

func TestOriginalCreationConcurrentUniqueReceipts(t *testing.T) {
	file, dir := captureFixture(t)
	raw, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	var scope OriginalCreationScope
	if err = json.Unmarshal(raw, &scope); err != nil {
		t.Fatal(err)
	}
	scope.Resources = nil
	for _, name := range []string{"example-a", "example-b", "example-c", "example-d"} {
		scope.Resources = append(scope.Resources, OriginalCreationTarget{TypeName: "xcsh_token", Namespace: "system", Name: name})
	}
	raw, err = json.Marshal(scope)
	if err != nil {
		t.Fatal(err)
	}
	if err = os.Chmod(file, 0600); err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(file, raw, 0600); err != nil {
		t.Fatal(err)
	}
	if err = os.Chmod(file, 0400); err != nil {
		t.Fatal(err)
	}
	results := make(chan error, len(scope.Resources))
	for _, target := range scope.Resources {
		go func(target OriginalCreationTarget) {
			c, err := OpenOriginalCreationCapture(target.TypeName, target.Namespace, target.Name)
			if err != nil {
				results <- err
				return
			}
			defer c.Close()
			_, err = c.RecordCreation(fmt.Sprintf("%x", sha256.Sum256([]byte(target.Name))))
			results <- err
		}(target)
	}
	for range scope.Resources {
		if err := <-results; err != nil {
			t.Fatal(err)
		}
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	receipts := 0
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "receipt-") {
			receipts++
		}
	}
	if receipts != len(scope.Resources) {
		t.Fatal("concurrent receipts collided or disappeared")
	}
}

func TestOriginalCreationDuplicateUIDRejected(t *testing.T) {
	file, _ := captureFixture(t)
	raw, e := os.ReadFile(file)
	if e != nil {
		t.Fatal(e)
	}
	var scope OriginalCreationScope
	if e = json.Unmarshal(raw, &scope); e != nil {
		t.Fatal(e)
	}
	scope.Resources = append(scope.Resources, OriginalCreationTarget{TypeName: "xcsh_token", Namespace: "system", Name: "example-other"})
	raw, e = json.Marshal(scope)
	if e != nil {
		t.Fatal(e)
	}
	if e = os.Chmod(file, 0600); e != nil {
		t.Fatal(e)
	}
	if e = os.WriteFile(file, raw, 0600); e != nil {
		t.Fatal(e)
	}
	if e = os.Chmod(file, 0400); e != nil {
		t.Fatal(e)
	}
	first, e := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
	if e != nil {
		t.Fatal(e)
	}
	defer first.Close()
	second, e := OpenOriginalCreationCapture("xcsh_token", "system", "example-other")
	if e != nil {
		t.Fatal(e)
	}
	defer second.Close()
	if _, e = first.RecordCreation(strings.Repeat("b", 64)); e != nil {
		t.Fatal(e)
	}
	if _, e = second.RecordCreation(strings.Repeat("b", 64)); e == nil {
		t.Fatal("same original UID admitted for two targets")
	}
}
