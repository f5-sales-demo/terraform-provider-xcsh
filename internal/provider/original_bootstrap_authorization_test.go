//go:build linux

package provider

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func bootstrapFixture(t *testing.T) (*OriginalCreationCapture, *OriginalCreationReceipt, ed25519.PrivateKey, string) {
	t.Helper()
	root := t.TempDir()
	receipts := filepath.Join(root, "receipts")
	auth := filepath.Join(root, "authorization")
	for _, dir := range []string{receipts, auth} {
		if err := os.Mkdir(dir, 0700); err != nil {
			t.Fatal(err)
		}
	}
	pub, key, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	process, err := originalCreationProcess()
	if err != nil {
		t.Fatal(err)
	}
	now, err := bootstrapMonotonicMillis()
	if err != nil {
		t.Fatal(err)
	}
	scope := map[string]any{
		"schema_version": 2, "scope_sha256": strings.Repeat("a", 64), "directory": receipts,
		"provider_executable_sha256": process.ExecutableSHA256,
		"resources":                  []OriginalCreationTarget{{TypeName: "xcsh_token", Name: "example-token", Namespace: "system"}},
		"bootstrap_authorization": map[string]any{"schema_version": 1, "directory": auth,
			"public_key":        base64.RawStdEncoding.EncodeToString(pub),
			"original_start_ms": now, "original_deadline_ms": now + 7200000},
	}
	raw, err := json.Marshal(scope)
	if err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(root, "scope.json")
	if err := os.WriteFile(file, raw, 0400); err != nil {
		t.Fatal(err)
	}
	t.Setenv("XCSH_ORIGINAL_CREATION_SCOPE_FILE", file)
	c, err := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(c.Close)
	receipt, err := c.RecordCreation(strings.Repeat("b", 64))
	if err != nil {
		t.Fatal(err)
	}
	return c, receipt, key, auth
}
func writeBootstrapGrant(t *testing.T, r *OriginalCreationReceipt, key ed25519.PrivateKey, auth, content string, alter func(map[string]any)) {
	t.Helper()
	now, err := bootstrapMonotonicMillis()
	if err != nil {
		t.Fatal(err)
	}
	payload := map[string]any{
		"schema_version": 1, "kind": "original-bootstrap-release",
		"scope_sha256": r.ScopeSHA256, "scope_file_sha256": r.ScopeFileSHA256,
		"target": r.Target, "creation_identity_sha256": r.CreationIdentitySHA256,
		"process": r.Process, "nonce": r.Nonce, "content_sha256": bootstrapContentHash(content),
		"issued_monotonic_ms": now, "expires_at_epoch_ms": time.Now().Add(time.Hour).UnixMilli(),
	}
	if alter != nil {
		alter(payload)
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	sig := ed25519.Sign(key, append([]byte("XCSH_BOOTSTRAP_RELEASE_V1\n"), raw...))
	envelope, err := json.Marshal(map[string]any{"payload": json.RawMessage(raw), "signature": base64.RawStdEncoding.EncodeToString(sig)})
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(auth, "grant-"+r.Nonce+".json"), envelope, 0400); err != nil {
		t.Fatal(err)
	}
}
func TestBootstrapLegacyCaptureDoesNotWait(t *testing.T) {
	captureFixture(t)
	c, err := OpenOriginalCreationCapture("xcsh_token", "system", "example-token")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	if _, err := c.RecordCreation(strings.Repeat("b", 64)); err != nil {
		t.Fatal(err)
	}
	if err := c.AwaitBootstrapAuthorization(context.Background(), "synthetic.jwt.content"); err != nil {
		t.Fatal(err)
	}
	var disabled *OriginalCreationCapture
	if err := disabled.AwaitBootstrapAuthorization(context.Background(), ""); err != nil {
		t.Fatal(err)
	}
}
func TestBootstrapMissingGrantAndCancellationCannotRelease(t *testing.T) {
	c, r, _, auth := bootstrapFixture(t)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	if err := c.AwaitBootstrapAuthorization(ctx, "synthetic.jwt.content"); err == nil {
		t.Fatal("missing authorization released content")
	}
	if _, err := os.Stat(filepath.Join(auth, "used-"+r.Nonce+".json")); !os.IsNotExist(err) {
		t.Fatal("missing grant consumed")
	}
	if err := c.AwaitBootstrapAuthorization(context.Background(), "synthetic.jwt.content"); err == nil {
		t.Fatal("failed authorization retried")
	}
}
func TestBootstrapValidGrantIsDurablyConsumedOnce(t *testing.T) {
	c, r, key, auth := bootstrapFixture(t)
	writeBootstrapGrant(t, r, key, auth, "synthetic.jwt.content", nil)
	if err := c.AwaitBootstrapAuthorization(context.Background(), "synthetic.jwt.content"); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(auth, "used-"+r.Nonce+".json")
	info, err := os.Stat(path)
	if err != nil || info.Mode().Perm() != 0400 {
		t.Fatal("durable private consumption absent", err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "synthetic.jwt.content") {
		t.Fatal("bootstrap content leaked")
	}
	if err := c.AwaitBootstrapAuthorization(context.Background(), "synthetic.jwt.content"); err == nil {
		t.Fatal("bootstrap release replayed")
	}
}
func TestBootstrapGrantRejectsWrongOriginalBindings(t *testing.T) {
	for _, field := range []string{"scope_sha256", "scope_file_sha256", "creation_identity_sha256", "nonce", "content_sha256", "process", "target", "expires_at_epoch_ms", "issued_monotonic_ms", "signature"} {
		t.Run(field, func(t *testing.T) {
			c, r, key, auth := bootstrapFixture(t)
			signing := key
			if field == "signature" {
				_, signing, _ = ed25519.GenerateKey(rand.Reader)
			}
			writeBootstrapGrant(t, r, signing, auth, "synthetic.jwt.content", func(p map[string]any) {
				switch field {
				case "signature":
				case "process":
					p[field] = OriginalCreationProcess{}
				case "target":
					p[field] = OriginalCreationTarget{TypeName: "xcsh_token", Name: "foreign", Namespace: "system"}
				case "expires_at_epoch_ms":
					p[field] = time.Now().Add(-time.Hour).UnixMilli()
				case "issued_monotonic_ms":
					p[field] = int64(0)
				default:
					p[field] = strings.Repeat("c", 64)
				}
			})
			if err := c.AwaitBootstrapAuthorization(context.Background(), "synthetic.jwt.content"); err == nil {
				t.Fatal("foreign/expired grant released content")
			}
			if _, err := os.Stat(filepath.Join(auth, "used-"+r.Nonce+".json")); !os.IsNotExist(err) {
				t.Fatal("invalid grant consumed")
			}
		})
	}
}
func TestBootstrapGrantSourceReplacementAndUnsafeModesRefuse(t *testing.T) {
	for _, change := range []string{"scope", "directory", "symlink", "mode", "oversize", "malformed"} {
		t.Run(change, func(t *testing.T) {
			c, r, key, auth := bootstrapFixture(t)
			writeBootstrapGrant(t, r, key, auth, "synthetic.jwt.content", nil)
			grant := filepath.Join(auth, "grant-"+r.Nonce+".json")
			switch change {
			case "scope":
				raw, e := os.ReadFile(c.scopeFile)
				if e != nil {
					t.Fatal(e)
				}
				if e = os.Rename(c.scopeFile, c.scopeFile+".old"); e != nil {
					t.Fatal(e)
				}
				if e = os.WriteFile(c.scopeFile, raw, 0400); e != nil {
					t.Fatal(e)
				}
			case "directory":
				if e := os.Rename(auth, auth+".old"); e != nil {
					t.Fatal(e)
				}
				if e := os.Mkdir(auth, 0700); e != nil {
					t.Fatal(e)
				}
			case "symlink":
				if e := os.Rename(grant, grant+".old"); e != nil {
					t.Fatal(e)
				}
				if e := os.Symlink(grant+".old", grant); e != nil {
					t.Fatal(e)
				}
			case "mode":
				if e := os.Chmod(grant, 0600); e != nil {
					t.Fatal(e)
				}
			case "oversize", "malformed":
				if e := os.Remove(grant); e != nil {
					t.Fatal(e)
				}
				raw := []byte("{")
				if change == "oversize" {
					raw = []byte(strings.Repeat("x", 20000))
				}
				if e := os.WriteFile(grant, raw, 0400); e != nil {
					t.Fatal(e)
				}
			}
			if err := c.AwaitBootstrapAuthorization(context.Background(), "synthetic.jwt.content"); err == nil {
				t.Fatal("changed authorization inputs accepted")
			}
		})
	}
}

func TestBootstrapConcurrentReleaseHasOneSuccess(t *testing.T) {
	c, r, key, auth := bootstrapFixture(t)
	writeBootstrapGrant(t, r, key, auth, "synthetic.jwt.content", nil)
	results := make(chan error, 2)
	for i := 0; i < 2; i++ {
		go func() { results <- c.AwaitBootstrapAuthorization(context.Background(), "synthetic.jwt.content") }()
	}
	successes := 0
	for i := 0; i < 2; i++ {
		if <-results == nil {
			successes++
		}
	}
	if successes != 1 {
		t.Fatalf("expected one original release, got %d", successes)
	}
}
func TestBootstrapDecoderRejectsDuplicateMissingAndCaseChangedFields(t *testing.T) {
	for _, raw := range []string{
		`{"payload":{},"signature":"a","signature":"b"}`,
		`{"payload":{},"Signature":"a"}`,
		`{"payload":{}}`,
	} {
		var value originalBootstrapEnvelope
		if bootstrapDecode([]byte(raw), &value) == nil {
			t.Fatal("malformed envelope accepted")
		}
	}
}
func TestBootstrapVersionOneRejectsNewAuthorizationFieldEvenNull(t *testing.T) {
	file, _ := captureFixture(t)
	raw, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	var scope map[string]any
	if err = json.Unmarshal(raw, &scope); err != nil {
		t.Fatal(err)
	}
	scope["bootstrap_authorization"] = nil
	changed, err := json.Marshal(scope)
	if err != nil {
		t.Fatal(err)
	}
	if err = os.Remove(file); err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(file, changed, 0400); err != nil {
		t.Fatal(err)
	}
	if _, err = OpenOriginalCreationCapture("xcsh_token", "system", "example-token"); err == nil {
		t.Fatal("version one accepted a new field")
	}
}
