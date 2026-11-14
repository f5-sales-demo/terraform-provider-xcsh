package provider

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

// OriginalCreationScope is an opt-in, private request prepared by the owning
// executor. ScopeSHA256 binds its reviewed operation; ScopeFileSHA256 in the
// receipt independently binds the exact serialized configuration.
type OriginalCreationScope struct {
	SchemaVersion            int                                  `json:"schema_version"`
	ScopeSHA256              string                               `json:"scope_sha256"`
	Directory                string                               `json:"directory"`
	ProviderExecutableSHA256 string                               `json:"provider_executable_sha256"`
	Resources                []OriginalCreationTarget             `json:"resources"`
	BootstrapAuthorization   *OriginalBootstrapAuthorizationScope `json:"bootstrap_authorization,omitempty"`
}

type OriginalCreationTarget struct {
	TypeName  string `json:"type_name"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// These fields permit independent /proc and executable checks. A receipt alone
// is not authentication: the executor must observe the original child and sink
// descriptors while it runs, then corroborate terminal state separately.
type OriginalCreationProcess struct {
	PID              int    `json:"pid"`
	ParentPID        int    `json:"parent_pid"`
	StartTicks       string `json:"start_ticks"`
	BootIDSHA256     string `json:"boot_id_sha256"`
	ExecutableSHA256 string `json:"executable_sha256"`
	ExecutableDevice uint64 `json:"executable_device"`
	ExecutableInode  uint64 `json:"executable_inode"`
}

type OriginalCreationReceipt struct {
	SchemaVersion          int                     `json:"schema_version"`
	ScopeSHA256            string                  `json:"scope_sha256"`
	ScopeFileSHA256        string                  `json:"scope_file_sha256"`
	Target                 OriginalCreationTarget  `json:"target"`
	CreationIdentitySHA256 string                  `json:"creation_identity_sha256,omitempty"`
	Process                OriginalCreationProcess `json:"process"`
	Nonce                  string                  `json:"nonce"`
}

type OriginalCreationCapture struct {
	mu                 sync.Mutex
	root               *os.Root
	scopeFile          string
	scopeInfo          os.FileInfo
	scopeBytes         []byte
	directory          string
	directoryInfo      os.FileInfo
	receipt            OriginalCreationReceipt
	attempted          bool
	claimName          string
	claimInfo          os.FileInfo
	claimBytes         []byte
	bootstrap          *originalBootstrapLease
	creationReceipt    *OriginalCreationReceipt
	bootstrapAttempted bool
	bootstrapReleased  bool
}

func captureError() error {
	return errors.New("original creation capture is unsafe, changed, ambiguous, or unavailable; retain ownership and reconcile before another submission")
}
func captureHash(s string) bool {
	if len(s) != 64 {
		return false
	}
	b, e := hex.DecodeString(s)
	return e == nil && len(b) == 32 && strings.ToLower(s) == s
}

func privateCapturePath(path string, directory bool) (os.FileInfo, error) {
	if !filepath.IsAbs(path) || filepath.Clean(path) != path {
		return nil, captureError()
	}
	current := string(filepath.Separator)
	for _, part := range strings.Split(strings.TrimPrefix(path, current), string(filepath.Separator)) {
		current = filepath.Join(current, part)
		info, err := os.Lstat(current)
		if err != nil || info.Mode()&os.ModeSymlink != 0 {
			return nil, captureError()
		}
	}
	info, err := os.Lstat(path)
	if err != nil || !captureOwned(info) {
		return nil, captureError()
	}
	if directory {
		if !info.IsDir() || info.Mode().Perm() != 0700 {
			return nil, captureError()
		}
	} else if !info.Mode().IsRegular() || info.Mode().Perm() != 0400 || info.Size() > 65536 {
		return nil, captureError()
	}
	return info, nil
}

func readCaptureScope(file string) ([]byte, os.FileInfo, error) {
	info, err := privateCapturePath(file, false)
	if err != nil {
		return nil, nil, err
	}
	root, err := os.OpenRoot(filepath.Dir(file))
	if err != nil {
		return nil, nil, captureError()
	}
	defer func() { _ = root.Close() }()
	f, err := root.Open(filepath.Base(file))
	if err != nil {
		return nil, nil, captureError()
	}
	defer func() { _ = f.Close() }()
	opened, err := f.Stat()
	if err != nil || !os.SameFile(info, opened) {
		return nil, nil, captureError()
	}
	raw, err := io.ReadAll(io.LimitReader(f, 65537))
	closeErr := f.Close()
	if err != nil || closeErr != nil || len(raw) > 65536 {
		return nil, nil, captureError()
	}
	after, err := privateCapturePath(file, false)
	if err != nil || !os.SameFile(info, after) || !info.ModTime().Equal(after.ModTime()) {
		return nil, nil, captureError()
	}
	return raw, info, nil
}

func captureTargetValid(t OriginalCreationTarget) bool {
	if t.TypeName != "xcsh_token" && t.TypeName != "xcsh_securemesh_site_v2" {
		return false
	}
	for _, s := range []string{t.Name, t.Namespace} {
		if s == "" || len(s) > 253 {
			return false
		}
		for _, r := range s {
			allowed := r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == '-' || r == '.'
			if !allowed {
				return false
			}
		}
	}
	return true
}

// OpenOriginalCreationCapture reserves a target durably before POST. Claims are
// never removed on error or close; interrupted creation requires reconciliation.
func OpenOriginalCreationCapture(typeName, namespace, name string) (*OriginalCreationCapture, error) {
	file := os.Getenv("XCSH_ORIGINAL_CREATION_SCOPE_FILE")
	if file == "" {
		return nil, nil
	}
	raw, info, err := readCaptureScope(file)
	if err != nil {
		return nil, err
	}
	if err := captureUniqueJSON(raw); err != nil {
		return nil, err
	}
	var scope OriginalCreationScope
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err = decoder.Decode(&scope); err != nil {
		return nil, captureError()
	}
	var extra any
	if decoder.Decode(&extra) != io.EOF || (scope.SchemaVersion != 1 && scope.SchemaVersion != 2) || !captureHash(scope.ScopeSHA256) || !captureHash(scope.ProviderExecutableSHA256) || len(scope.Resources) == 0 || len(scope.Resources) > 128 {
		return nil, captureError()
	}
	var scopeFields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &scopeFields); err != nil {
		return nil, captureError()
	}
	if _, declared := scopeFields["bootstrap_authorization"]; scope.SchemaVersion == 1 && declared {
		return nil, captureError()
	}
	if scope.SchemaVersion == 2 {
		if err := bootstrapExactKeys(raw, reflect.TypeOf(OriginalCreationScope{})); err != nil {
			return nil, err
		}
	}
	target := OriginalCreationTarget{TypeName: typeName, Namespace: namespace, Name: name}
	seen := map[OriginalCreationTarget]bool{}
	for _, t := range scope.Resources {
		if !captureTargetValid(t) || seen[t] {
			return nil, captureError()
		}
		seen[t] = true
	}
	if !seen[target] {
		return nil, captureError()
	}
	dirInfo, err := privateCapturePath(scope.Directory, true)
	if err != nil {
		return nil, err
	}
	root, err := os.OpenRoot(scope.Directory)
	if err != nil {
		return nil, captureError()
	}
	keep := false
	defer func() {
		if !keep {
			_ = root.Close()
		}
	}()
	opened, err := root.Stat(".")
	if err != nil || !os.SameFile(dirInfo, opened) {
		return nil, captureError()
	}
	process, err := originalCreationProcess()
	if err != nil || process.ExecutableSHA256 != scope.ProviderExecutableSHA256 {
		return nil, captureError()
	}
	var nonce [16]byte
	if _, err = rand.Read(nonce[:]); err != nil {
		return nil, captureError()
	}
	bootstrap, err := openOriginalBootstrapScope(scope)
	if err != nil {
		return nil, err
	}
	defer func() {
		if !keep {
			bootstrap.close()
		}
	}()
	rawHash := sha256.Sum256(raw)
	capture := &OriginalCreationCapture{root: root, scopeFile: file, scopeInfo: info, scopeBytes: raw, directory: scope.Directory, directoryInfo: dirInfo, bootstrap: bootstrap, receipt: OriginalCreationReceipt{SchemaVersion: 1, ScopeSHA256: scope.ScopeSHA256, ScopeFileSHA256: hex.EncodeToString(rawHash[:]), Target: target, Process: process, Nonce: hex.EncodeToString(nonce[:])}}
	targetBytes, _ := json.Marshal(target)
	targetHash := sha256.Sum256(targetBytes)
	claim := "claim-" + hex.EncodeToString(targetHash[:]) + ".json"
	if err = capture.verify(); err != nil {
		return nil, err
	}
	if err = capture.persist(claim, capture.receipt); err != nil {
		return nil, err
	}
	capture.claimName = claim
	capture.claimInfo, err = root.Lstat(claim)
	if err != nil {
		return nil, captureError()
	}
	capture.claimBytes, err = json.Marshal(capture.receipt)
	if err != nil {
		return nil, captureError()
	}
	if err = capture.verify(); err != nil {
		return nil, err
	}
	keep = true
	return capture, nil
}

func (c *OriginalCreationCapture) verify() error {
	raw, info, err := readCaptureScope(c.scopeFile)
	if err != nil || !os.SameFile(info, c.scopeInfo) || !bytes.Equal(raw, c.scopeBytes) {
		return captureError()
	}
	info, err = privateCapturePath(c.directory, true)
	if err != nil || !os.SameFile(info, c.directoryInfo) {
		return captureError()
	}
	opened, err := c.root.Stat(".")
	if err != nil || !os.SameFile(info, opened) {
		return captureError()
	}
	if c.claimName != "" {
		info, err := c.root.Lstat(c.claimName)
		if err != nil || !os.SameFile(info, c.claimInfo) || !info.Mode().IsRegular() || info.Mode().Perm() != 0400 || !captureOwned(info) {
			return captureError()
		}
		f, err := c.root.Open(c.claimName)
		if err != nil {
			return captureError()
		}
		opened, statErr := f.Stat()
		raw, readErr := io.ReadAll(io.LimitReader(f, 65537))
		closeErr := f.Close()
		if statErr != nil || readErr != nil || closeErr != nil || !os.SameFile(info, opened) || !bytes.Equal(raw, c.claimBytes) {
			return captureError()
		}
	}
	process, err := originalCreationProcess()
	if err != nil || process != c.receipt.Process {
		return captureError()
	}
	return nil
}

func (c *OriginalCreationCapture) persist(name string, receipt OriginalCreationReceipt) error {
	raw, err := json.Marshal(receipt)
	if err != nil {
		return captureError()
	}
	f, err := c.root.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return captureError()
	}
	defer func() { _ = f.Close() }()
	if _, err = f.Write(raw); err != nil {
		return captureError()
	}
	if err = f.Chmod(0400); err != nil {
		return captureError()
	}
	if err = f.Sync(); err != nil {
		return captureError()
	}
	if err = f.Close(); err != nil {
		return captureError()
	}
	dir, err := c.root.Open(".")
	if err != nil {
		return captureError()
	}
	defer func() { _ = dir.Close() }()
	if err = dir.Sync(); err != nil {
		return captureError()
	}
	if err = dir.Close(); err != nil {
		return captureError()
	}
	return c.verify()
}

func (c *OriginalCreationCapture) RecordCreation(hash string) (*OriginalCreationReceipt, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.attempted || c.root == nil {
		return nil, captureError()
	}
	c.attempted = true
	hash = strings.TrimPrefix(hash, "sha256:")
	if !captureHash(hash) {
		return nil, captureError()
	}
	if err := c.verify(); err != nil {
		return nil, err
	}
	receipt := c.receipt
	receipt.CreationIdentitySHA256 = hash
	// Reserve the immutable UID across all targets before publishing evidence.
	// Keep this marker after interruption or failure; it cannot authorize replay.
	if err := c.persist("identity-"+hash+".json", receipt); err != nil {
		return nil, err
	}
	if err := c.persist("receipt-"+receipt.Nonce+".json", receipt); err != nil {
		return nil, err
	}
	c.creationReceipt = &receipt
	return &receipt, nil
}

func (c *OriginalCreationCapture) Close() {
	if c == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.bootstrap.close()
	if c.root != nil {
		_ = c.root.Close()
		c.root = nil
	}
}

// encoding/json accepts duplicate object keys; a provenance scope must not.
func captureUniqueJSON(raw []byte) error {
	d := json.NewDecoder(bytes.NewReader(raw))
	var value func(int) error
	value = func(depth int) error {
		if depth > 8 {
			return captureError()
		}
		token, err := d.Token()
		if err != nil {
			return captureError()
		}
		delim, ok := token.(json.Delim)
		if !ok {
			return nil
		}
		switch delim {
		case '{':
			keys := map[string]bool{}
			for d.More() {
				key, err := d.Token()
				if err != nil {
					return captureError()
				}
				text, ok := key.(string)
				if !ok || keys[text] {
					return captureError()
				}
				keys[text] = true
				if err = value(depth + 1); err != nil {
					return err
				}
			}
			end, err := d.Token()
			if err != nil || end != json.Delim('}') {
				return captureError()
			}
		case '[':
			for d.More() {
				if err = value(depth + 1); err != nil {
					return err
				}
			}
			end, err := d.Token()
			if err != nil || end != json.Delim(']') {
				return captureError()
			}
		default:
			return captureError()
		}
		return nil
	}
	if err := value(0); err != nil {
		return err
	}
	if _, err := d.Token(); err != io.EOF {
		return captureError()
	}
	return nil
}
