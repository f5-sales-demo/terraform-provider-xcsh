//go:build linux

package provider

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOriginalCreationClaimChanged(t *testing.T) {
	_, dir := captureFixture(t)
	c, e := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
	if e != nil {
		t.Fatal(e)
	}
	defer c.Close()
	entries, e := os.ReadDir(dir)
	if e != nil {
		t.Fatal(e)
	}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "claim-") {
			if e = os.Remove(filepath.Join(dir, entry.Name())); e != nil {
				t.Fatal(e)
			}
		}
	}
	if _, e = c.RecordCreation(strings.Repeat("b", 64)); e == nil {
		t.Fatal("deleted original claim admitted")
	}
}
