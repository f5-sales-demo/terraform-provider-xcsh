package provider

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

const bootstrapSignatureDomain = "XCSH_BOOTSTRAP_RELEASE_V1\n"
const bootstrapMaximumInteger int64 = 9007199254740991

// OriginalBootstrapAuthorizationScope binds a source-owned authorization key,
// directory and original budget. It is accepted only by scope version 2.
type OriginalBootstrapAuthorizationScope struct {
	SchemaVersion      int    `json:"schema_version"`
	Directory          string `json:"directory"`
	PublicKey          string `json:"public_key"`
	OriginalStartMS    int64  `json:"original_start_ms"`
	OriginalDeadlineMS int64  `json:"original_deadline_ms"`
}
type originalBootstrapLease struct {
	scope         OriginalBootstrapAuthorizationScope
	root          *os.Root
	directoryInfo os.FileInfo
	publicKey     ed25519.PublicKey
}
type originalBootstrapGrant struct {
	SchemaVersion          int                     `json:"schema_version"`
	Kind                   string                  `json:"kind"`
	ScopeSHA256            string                  `json:"scope_sha256"`
	ScopeFileSHA256        string                  `json:"scope_file_sha256"`
	Target                 OriginalCreationTarget  `json:"target"`
	CreationIdentitySHA256 string                  `json:"creation_identity_sha256"`
	Process                OriginalCreationProcess `json:"process"`
	Nonce                  string                  `json:"nonce"`
	ContentSHA256          string                  `json:"content_sha256"`
	IssuedMonotonicMS      int64                   `json:"issued_monotonic_ms"`
	ExpiresAtEpochMS       int64                   `json:"expires_at_epoch_ms"`
}
type originalBootstrapEnvelope struct {
	Payload   json.RawMessage `json:"payload"`
	Signature string          `json:"signature"`
}

func bootstrapError() error {
	return errors.New("original bootstrap authorization is missing, changed, expired, consumed, or unsafe; retain created ownership and reconcile without replay")
}
func bootstrapContentHash(content string) string {
	digest := sha256.Sum256([]byte(content))
	return hex.EncodeToString(digest[:])
}
func bootstrapRawBase64(value string, size int) ([]byte, error) {
	raw, err := base64.RawStdEncoding.DecodeString(value)
	if err != nil || len(raw) != size || base64.RawStdEncoding.EncodeToString(raw) != value {
		return nil, bootstrapError()
	}
	return raw, nil
}
func openOriginalBootstrapScope(scope OriginalCreationScope) (*originalBootstrapLease, error) {
	if scope.SchemaVersion == 1 {
		if scope.BootstrapAuthorization != nil {
			return nil, bootstrapError()
		}
		return nil, nil
	}
	spec := scope.BootstrapAuthorization
	if scope.SchemaVersion != 2 || spec == nil || spec.SchemaVersion != 1 || spec.OriginalStartMS < 0 || spec.OriginalStartMS > bootstrapMaximumInteger || spec.OriginalDeadlineMS < 0 ||
		spec.OriginalDeadlineMS > bootstrapMaximumInteger || spec.OriginalDeadlineMS-spec.OriginalStartMS != 7200000 {
		return nil, bootstrapError()
	}
	pub, err := bootstrapRawBase64(spec.PublicKey, ed25519.PublicKeySize)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(spec.Directory, scope.Directory+string(filepath.Separator)) ||
		strings.HasPrefix(scope.Directory, spec.Directory+string(filepath.Separator)) || spec.Directory == scope.Directory ||
		strings.IndexFunc(spec.Directory, func(r rune) bool { return r < 32 || r == 127 }) >= 0 {
		return nil, bootstrapError()
	}
	info, err := privateCapturePath(spec.Directory, true)
	if err != nil {
		return nil, bootstrapError()
	}
	root, err := os.OpenRoot(spec.Directory)
	if err != nil {
		return nil, bootstrapError()
	}
	lease := &originalBootstrapLease{scope: *spec, root: root, directoryInfo: info, publicKey: ed25519.PublicKey(pub)}
	if err = lease.verify(); err != nil {
		lease.close()
		return nil, err
	}
	if err = lease.withinBudget(); err != nil {
		lease.close()
		return nil, err
	}
	return lease, nil
}
func (b *originalBootstrapLease) close() {
	if b != nil && b.root != nil {
		_ = b.root.Close()
		b.root = nil
	}
}
func (b *originalBootstrapLease) verify() error {
	if b == nil || b.root == nil {
		return bootstrapError()
	}
	info, err := privateCapturePath(b.scope.Directory, true)
	if err != nil || !os.SameFile(info, b.directoryInfo) {
		return bootstrapError()
	}
	opened, err := b.root.Stat(".")
	if err != nil || !os.SameFile(opened, b.directoryInfo) {
		return bootstrapError()
	}
	return nil
}
func (b *originalBootstrapLease) withinBudget() error {
	now, err := bootstrapMonotonicMillis()
	if err != nil || now < b.scope.OriginalStartMS || now >= b.scope.OriginalDeadlineMS {
		return bootstrapError()
	}
	return nil
}
func bootstrapExactKeys(raw []byte, typ reflect.Type) error {
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Slice && typ.Elem().Kind() == reflect.Struct {
		var entries []json.RawMessage
		if err := json.Unmarshal(raw, &entries); err != nil {
			return bootstrapError()
		}
		for _, entry := range entries {
			if err := bootstrapExactKeys(entry, typ.Elem()); err != nil {
				return err
			}
		}
		return nil
	}
	if typ.Kind() != reflect.Struct {
		return nil
	}
	var object map[string]json.RawMessage
	if err := json.Unmarshal(raw, &object); err != nil || object == nil || len(object) != typ.NumField() {
		return bootstrapError()
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		key := strings.Split(field.Tag.Get("json"), ",")[0]
		value, exists := object[key]
		if key == "" || !exists {
			return bootstrapError()
		}
		if err := bootstrapExactKeys(value, field.Type); err != nil {
			return err
		}
	}
	return nil
}

func bootstrapDecode(raw []byte, value any) error {
	if len(raw) == 0 || len(raw) > 16384 || captureUniqueJSON(raw) != nil {
		return bootstrapError()
	}
	if err := bootstrapExactKeys(raw, reflect.TypeOf(value)); err != nil {
		return err
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if decoder.Decode(value) != nil {
		return bootstrapError()
	}
	var extra any
	if decoder.Decode(&extra) != io.EOF {
		return bootstrapError()
	}
	return nil
}
func (b *originalBootstrapLease) readGrant(name string) ([]byte, os.FileInfo, error) {
	if err := b.verify(); err != nil {
		return nil, nil, err
	}
	info, err := b.root.Lstat(name)
	if err != nil {
		return nil, nil, err
	}
	if !info.Mode().IsRegular() || info.Mode().Perm() != 0400 || !captureOwned(info) || info.Size() > 16384 {
		return nil, nil, bootstrapError()
	}
	file, err := b.root.Open(name)
	if err != nil {
		return nil, nil, bootstrapError()
	}
	opened, statErr := file.Stat()
	raw, readErr := io.ReadAll(io.LimitReader(file, 16385))
	closeErr := file.Close()
	if statErr != nil || readErr != nil || closeErr != nil || !os.SameFile(info, opened) || len(raw) > 16384 {
		return nil, nil, bootstrapError()
	}
	after, err := b.root.Lstat(name)
	if err != nil || !os.SameFile(after, info) || !after.ModTime().Equal(info.ModTime()) {
		return nil, nil, bootstrapError()
	}
	return raw, info, nil
}
func (c *OriginalCreationCapture) verifyBootstrapGrant(raw []byte, content string) error {
	var envelope originalBootstrapEnvelope
	if err := bootstrapDecode(raw, &envelope); err != nil {
		return err
	}
	signature, err := bootstrapRawBase64(envelope.Signature, ed25519.SignatureSize)
	if err != nil {
		return err
	}
	message := append([]byte(bootstrapSignatureDomain), envelope.Payload...)
	if !ed25519.Verify(c.bootstrap.publicKey, message, signature) {
		return bootstrapError()
	}
	var grant originalBootstrapGrant
	if err := bootstrapDecode(envelope.Payload, &grant); err != nil {
		return err
	}
	receipt := c.creationReceipt
	now, err := bootstrapMonotonicMillis()
	if err != nil || receipt == nil || grant.SchemaVersion != 1 || grant.Kind != "original-bootstrap-release" ||
		grant.ScopeSHA256 != receipt.ScopeSHA256 || grant.ScopeFileSHA256 != receipt.ScopeFileSHA256 ||
		grant.Target != receipt.Target || grant.CreationIdentitySHA256 != receipt.CreationIdentitySHA256 ||
		grant.Process != receipt.Process || grant.Nonce != receipt.Nonce || grant.ContentSHA256 != bootstrapContentHash(content) ||
		grant.IssuedMonotonicMS < c.bootstrap.scope.OriginalStartMS || grant.IssuedMonotonicMS > now ||
		now-grant.IssuedMonotonicMS > 300000 || grant.ExpiresAtEpochMS > bootstrapMaximumInteger ||
		grant.ExpiresAtEpochMS <= time.Now().UnixMilli() {
		return bootstrapError()
	}
	return c.bootstrap.withinBudget()
}
func (c *OriginalCreationCapture) persistBootstrapUse(raw []byte) error {
	name := "used-" + c.creationReceipt.Nonce + ".json"
	marker, err := json.Marshal(struct {
		SchemaVersion int                     `json:"schema_version"`
		Original      OriginalCreationReceipt `json:"original"`
		GrantSHA256   string                  `json:"grant_sha256"`
	}{1, *c.creationReceipt, bootstrapContentHash(string(raw))})
	if err != nil {
		return bootstrapError()
	}
	file, err := c.bootstrap.root.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return bootstrapError()
	}
	// Failed writes or flushes leave the exclusive marker and original claim.
	writeErr := func() error {
		if _, e := file.Write(marker); e != nil {
			return e
		}
		if e := file.Chmod(0400); e != nil {
			return e
		}
		return file.Sync()
	}()
	closeErr := file.Close()
	if writeErr != nil || closeErr != nil {
		return bootstrapError()
	}
	dir, err := c.bootstrap.root.Open(".")
	if err != nil {
		return bootstrapError()
	}
	syncErr := dir.Sync()
	closeErr = dir.Close()
	if syncErr != nil || closeErr != nil {
		return bootstrapError()
	}
	return nil
}

// AwaitBootstrapAuthorization withholds Token content until the exact original
// signed grant is verified and durably consumed. Failure never permits a replay.
// Capture-disabled and version 1 calls preserve their prior behavior.
func (c *OriginalCreationCapture) AwaitBootstrapAuthorization(ctx context.Context, content string) error {
	if c == nil {
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.bootstrap == nil {
		return nil
	}
	if c.bootstrapAttempted || c.root == nil || c.creationReceipt == nil ||
		ctx == nil || c.creationReceipt.Target.TypeName != "xcsh_token" || content == "" || len(content) > 262144 {
		return bootstrapError()
	}
	c.bootstrapAttempted = true
	if err := c.verify(); err != nil {
		return err
	}
	name := "grant-" + c.creationReceipt.Nonce + ".json"
	for {
		if ctx.Err() != nil {
			return bootstrapError()
		}
		if err := c.bootstrap.withinBudget(); err != nil {
			return err
		}
		if err := c.bootstrap.verify(); err != nil {
			return err
		}
		scope, info, err := readCaptureScope(c.scopeFile)
		if err != nil || !os.SameFile(info, c.scopeInfo) || !bytes.Equal(scope, c.scopeBytes) {
			return bootstrapError()
		}
		raw, original, err := c.bootstrap.readGrant(name)
		if errors.Is(err, os.ErrNotExist) {
			timer := time.NewTimer(50 * time.Millisecond)
			select {
			case <-ctx.Done():
				timer.Stop()
				return bootstrapError()
			case <-timer.C:
			}
			continue
		}
		if err != nil {
			return bootstrapError()
		}
		if err = c.verify(); err != nil {
			return err
		}
		if err = c.verifyBootstrapGrant(raw, content); err != nil {
			return err
		}
		again, after, err := c.bootstrap.readGrant(name)
		if err != nil || !os.SameFile(original, after) || !bytes.Equal(raw, again) {
			return bootstrapError()
		}
		if err = c.persistBootstrapUse(raw); err != nil {
			return err
		}
		if ctx.Err() != nil {
			return bootstrapError()
		}
		if err = c.verify(); err != nil {
			return err
		}
		if err = c.bootstrap.verify(); err != nil {
			return err
		}
		if err = c.bootstrap.withinBudget(); err != nil {
			return err
		}
		final, last, err := c.bootstrap.readGrant(name)
		if err != nil || !os.SameFile(original, last) || !bytes.Equal(raw, final) {
			return bootstrapError()
		}
		if err := c.verifyBootstrapGrant(final, content); err != nil {
			return err
		}
		c.bootstrapReleased = true
		return nil
	}
}
