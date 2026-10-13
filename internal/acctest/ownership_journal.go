// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

var errOwnedAbsent = errors.New("owned resource absent")
var ownershipRunPattern = regexp.MustCompile(`^[a-z0-9]{1,24}$`)

type OwnedResource struct {
	Type          string   `json:"type"`
	Namespace     string   `json:"namespace,omitempty"`
	Name          string   `json:"name"`
	Dependencies  []string `json:"dependencies,omitempty"`
	UID           string   `json:"uid,omitempty"`
	Status        string   `json:"status"`
	Version       string   `json:"version,omitempty"`
	ContentSHA256 string   `json:"content_sha256,omitempty"`
}
type OwnershipJournal struct {
	Version   int             `json:"version"`
	RunID     string          `json:"run_id"`
	Resources []OwnedResource `json:"resources"`
	path      string
	mu        sync.Mutex
}
type CleanupReceipt struct {
	Version             int            `json:"version"`
	Diagnostics         map[string]int `json:"diagnostics,omitempty"`
	JournalSHA256       string         `json:"journal_sha256"`
	Complete            bool           `json:"complete"`
	Owned               int            `json:"owned"`
	ConfirmedAbsent     int            `json:"confirmed_absent"`
	Failures            int            `json:"failures"`
	SecondPassRemaining int            `json:"second_pass_remaining"`
}
type OwnershipBackend interface {
	Read(context.Context, OwnedResource) (client.Metadata, error)
	Delete(context.Context, OwnedResource) error
}

func ownedKey(r OwnedResource) string { return r.Type + "/" + r.Namespace + "/" + r.Name }
func CreateOwnershipJournal(path, run string) (*OwnershipJournal, error) {
	if !ownershipRunPattern.MatchString(run) {
		return nil, errors.New("invalid ownership run")
	}
	j := &OwnershipJournal{Version: 1, RunID: run, Resources: []OwnedResource{}, path: path}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(j)
	if err == nil {
		_, err = f.Write(data)
	}
	if err == nil {
		err = f.Sync()
	}
	closeErr := f.Close()
	if err == nil {
		err = closeErr
	}
	return j, err
}
func LoadOwnershipJournal(path, run string) (*OwnershipJournal, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() || info.Mode().Perm()&0077 != 0 {
		return nil, errors.New("ownership journal is not private")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var j OwnershipJournal
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err = decoder.Decode(&j); err != nil {
		return nil, errors.New("malformed ownership journal")
	}
	var trailing interface{}
	if err = decoder.Decode(&trailing); err != io.EOF {
		return nil, errors.New("trailing ownership journal data")
	}
	if j.Version != 1 || j.RunID != run || !ownershipRunPattern.MatchString(run) || j.Resources == nil {
		return nil, errors.New("ownership journal version or run mismatch")
	}
	j.path = path
	seen := map[string]bool{}
	for _, r := range j.Resources {
		if err = j.validate(r); err != nil {
			return nil, err
		}
		if seen[ownedKey(r)] {
			return nil, errors.New("duplicate ownership resource")
		}
		seen[ownedKey(r)] = true
		if r.Type != "xcsh_swagger_object" {
			if _, ok := ownedReaderRegistry[r.Type]; !ok {
				return nil, errors.New("ownership journal verifier unavailable")
			}
			if _, ok := resourceDeleterRegistry[r.Type]; !ok {
				return nil, errors.New("ownership journal deleter unavailable")
			}
		}
		if r.UID != "" {
			if !regexp.MustCompile(`^[0-9a-f]{64}$`).MatchString(r.UID) {
				return nil, errors.New("ownership journal UID is not a digest")
			}
		}
		if r.Status == "created" && r.UID == "" {
			return nil, errors.New("ownership journal created identity is unconfirmed")
		}
		switch r.Status {
		case "intent", "created", "absent":
		default:
			return nil, errors.New("invalid ownership status")
		}
	}
	return &j, nil
}
func (j *OwnershipJournal) validate(r OwnedResource) error {
	if r.Type == "" || !matchesOwnershipRun(r.Name, j.RunID) || strings.ContainsAny(r.Name+r.Namespace, "/?# \t\r\n") {
		return errors.New("resource outside ownership run")
	}
	return nil
}
func (j *OwnershipJournal) save() error {
	data, err := json.Marshal(j)
	if err != nil {
		return err
	}
	f, err := os.CreateTemp(filepath.Dir(j.path), ".ownership-*")
	if err != nil {
		return err
	}
	temp := f.Name()
	defer func() { _ = os.Remove(temp) }()
	if err = f.Chmod(0600); err == nil {
		_, err = f.Write(data)
	}
	if err == nil {
		err = f.Sync()
	}
	closeErr := f.Close()
	if err == nil {
		err = closeErr
	}
	if err != nil {
		return err
	}
	if err = os.Rename(temp, j.path); err != nil {
		return err
	}
	dir, err := os.Open(filepath.Dir(j.path))
	if err != nil {
		return err
	}
	defer func() { _ = dir.Close() }()
	return dir.Sync()
}
func (j *OwnershipJournal) ResetAfterVerifiedAbsence(r OwnedResource) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	for i, existing := range j.Resources {
		if ownedKey(existing) == ownedKey(r) {
			j.Resources[i].UID = ""
			j.Resources[i].Status = "intent"
			return j.save()
		}
	}
	return nil
}
func (j *OwnershipJournal) Intent(r OwnedResource) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	if err := j.validate(r); err != nil {
		return err
	}
	for _, existing := range j.Resources {
		if ownedKey(existing) == ownedKey(r) {
			if existing.Status == "absent" {
				return errors.New("cannot reuse deleted run identity")
			}
			return nil
		}
	}
	r.Status = "intent"
	r.UID = ""
	j.Resources = append(j.Resources, r)
	return j.save()
}
func ownershipUIDDigest(uid string) string {
	digest := sha256.Sum256([]byte(uid))
	return hex.EncodeToString(digest[:])
}
func (j *OwnershipJournal) Confirm(r OwnedResource, uid string) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	if uid == "" {
		return errors.New("ownership readback lacks UID")
	}
	uid = ownershipUIDDigest(uid)
	for i, existing := range j.Resources {
		if ownedKey(existing) == ownedKey(r) {
			if existing.UID != "" && existing.UID != uid {
				return errors.New("ownership UID changed")
			}
			j.Resources[i].UID = uid
			j.Resources[i].Status = "created"
			return j.save()
		}
	}
	return errors.New("ownership intent missing")
}
func (j *OwnershipJournal) cleanupOrder() ([]int, error) {
	order := []int{}
	visited := map[string]int{}
	indices := map[string]int{}
	for i, r := range j.Resources {
		indices[ownedKey(r)] = i
	}
	var visit func(int) error
	visit = func(i int) error {
		key := ownedKey(j.Resources[i])
		if visited[key] == 1 {
			return errors.New("ownership dependency cycle")
		}
		if visited[key] == 2 {
			return nil
		}
		visited[key] = 1
		for _, dependency := range j.Resources[i].Dependencies {
			if parent, ok := indices[dependency]; ok {
				if err := visit(parent); err != nil {
					return err
				}
			} else {
				return errors.New("ownership dependency missing")
			}
		}
		visited[key] = 2
		order = append(order, i)
		return nil
	}
	for i := range j.Resources {
		if err := visit(i); err != nil {
			return nil, err
		}
	}
	for i, k := 0, len(order)-1; i < k; i, k = i+1, k-1 {
		order[i], order[k] = order[k], order[i]
	}
	return order, nil
}
func (j *OwnershipJournal) Cleanup(ctx context.Context, b OwnershipBackend, attempts int, delay time.Duration) (CleanupReceipt, error) {
	j.mu.Lock()
	defer j.mu.Unlock()
	receipt := CleanupReceipt{Version: 1, Owned: len(j.Resources), Diagnostics: map[string]int{}}
	order, err := j.cleanupOrder()
	if err != nil {
		receipt.Failures++
		receipt.Diagnostics["invalid_dependency_order"]++
		return receipt, err
	}
	failedResources := map[string]bool{}
	for _, i := range order {
		r := j.Resources[i]
		blocked := false
		for _, child := range j.Resources {
			if failedResources[ownedKey(child)] {
				for _, dependency := range child.Dependencies {
					if dependency == ownedKey(r) {
						blocked = true
					}
				}
			}
		}
		if blocked {
			receipt.Diagnostics["dependency_cleanup_failed"]++
			receipt.Failures++
			failedResources[ownedKey(r)] = true
			continue
		}
		absent := false
		fatal := false
		for attempt := 0; attempt < attempts; attempt++ {
			metadata, readErr := b.Read(ctx, r)
			if errors.Is(readErr, errOwnedAbsent) {
				absent = true
				break
			}
			if readErr != nil || metadata.Name != r.Name || metadata.UID == "" || (r.UID != "" && ownershipUIDDigest(metadata.UID) != r.UID) {
				receipt.Diagnostics["ownership_readback_failed"]++
				fatal = true
				break
			}
			if r.Status == "absent" {
				fatal = true
				break
			}
			if r.UID == "" && metadata.Annotations[ownershipAnnotation] != j.RunID {
				receipt.Diagnostics["creation_ownership_unconfirmed"]++
				fatal = true
				break
			}
			if r.UID == "" {
				r.UID = ownershipUIDDigest(metadata.UID)
				r.Status = "created"
				j.Resources[i] = r
				if j.save() != nil {
					fatal = true
					break
				}
			}
			if deleteErr := b.Delete(ctx, r); deleteErr != nil && !errors.Is(deleteErr, errOwnedAbsent) {
				receipt.Diagnostics["delete_failed"]++
				receipt.Diagnostics[cleanupDeleteCategory(deleteErr)]++
				fatal = true
				break
			}
			if waitContext(ctx, delay) != nil {
				receipt.Diagnostics["cleanup_interrupted"]++
				fatal = true
				break
			}
		}
		// Verify after the final delete as well.
		if !absent && !fatal {
			_, readErr := b.Read(ctx, r)
			absent = errors.Is(readErr, errOwnedAbsent)
		}
		if !absent && !fatal {
			receipt.Diagnostics["resource_remains_visible"]++
		}
		if absent {
			j.Resources[i].Status = "absent"
			receipt.ConfirmedAbsent++
		} else {
			receipt.Failures++
			failedResources[ownedKey(r)] = true
		}
	}
	if err = j.save(); err != nil {
		receipt.Failures++
	}
	for _, r := range j.Resources {
		if _, readErr := b.Read(ctx, r); !errors.Is(readErr, errOwnedAbsent) {
			receipt.Diagnostics["second_pass_not_absent"]++
			receipt.SecondPassRemaining++
		}
	}
	data, digestErr := os.ReadFile(j.path)
	if digestErr != nil {
		receipt.Failures++
	} else {
		digest := sha256.Sum256(data)
		receipt.JournalSHA256 = hex.EncodeToString(digest[:])
	}
	receipt.Complete = receipt.Failures == 0 && receipt.ConfirmedAbsent == receipt.Owned && receipt.SecondPassRemaining == 0
	if !receipt.Complete {
		return receipt, errors.New("ownership cleanup verification failed")
	}
	return receipt, nil
}

type apiOwnershipBackend struct {
	c   *client.Client
	run string
}

func (b apiOwnershipBackend) Read(ctx context.Context, r OwnedResource) (client.Metadata, error) {
	if r.Type == "xcsh_swagger_object" {
		return b.readSwagger(ctx, r)
	}
	read, ok := ownedReaderRegistry[r.Type]
	if !ok {
		return client.Metadata{}, fmt.Errorf("no API verifier registered for resource type")
	}
	metadata, err := read(context.WithValue(ctx, ownershipReadContextKey{}, true), b.c, r.Namespace, r.Name)
	if authoritativeOwnedAbsence(err) {
		return client.Metadata{}, errOwnedAbsent
	}
	return metadata, err
}
func (b apiOwnershipBackend) Delete(ctx context.Context, r OwnedResource) error {
	if r.Type == "xcsh_namespace" {
		err := deleteOwnedNamespace(ctx, b.c, r)
		if authoritativeOwnedAbsence(err) {
			return errOwnedAbsent
		}
		return err
	}
	if r.Type == "xcsh_swagger_object" {
		return b.deleteSwagger(ctx, r)
	}
	delete, ok := resourceDeleterRegistry[r.Type]
	if !ok {
		return errors.New("no deleter registered for resource type")
	}
	err := delete(ctx, b.c, r.Namespace, r.Name)
	if authoritativeOwnedAbsence(err) {
		return errOwnedAbsent
	}
	return err
}

func authoritativeOwnedAbsence(err error) bool {
	var apiErr *xcsherrors.XCSHError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != 404 {
		return false
	}
	if code, ok := apiErr.Details["api_code"].(string); ok {
		return code == "NOT_FOUND" || code == "5"
	}
	if raw, ok := apiErr.Details["raw_response"].(string); ok {
		var response struct {
			Code interface{} `json:"code"`
		}
		if json.Unmarshal([]byte(raw), &response) != nil {
			return false
		}
		return response.Code == float64(5) || response.Code == "NOT_FOUND" || response.Code == "5"
	}
	return false
}

// Ownership metadata is separate from mutable resource metadata in the API.
func readOwnershipProjection(ctx context.Context, c *client.Client, path, namespace, name string) (client.Metadata, error) {
	var projection struct {
		Metadata       client.Metadata `json:"metadata"`
		SystemMetadata struct {
			UID string `json:"uid"`
		} `json:"system_metadata"`
	}
	if err := c.Get(ctx, path, &projection); err != nil {
		return client.Metadata{}, err
	}
	metadata := projection.Metadata
	if metadata.UID == "" {
		metadata.UID = projection.SystemMetadata.UID
	}
	if metadata.Name != name || metadata.UID == "" || (metadata.Namespace != "" && namespace != "" && metadata.Namespace != namespace) {
		return client.Metadata{}, errors.New("invalid ownership readback")
	}
	return metadata, nil
}

func cleanupDeleteCategory(err error) string {
	switch err.Error() {
	case "namespace cascade ownership or empty inventory unverified":
		return "namespace_empty_inventory_unverified"
	case "namespace cascade response malformed":
		return "namespace_command_response_malformed"
	case "namespace cascade item failed":
		return "namespace_command_item_failed"
	}
	var apiErr *xcsherrors.XCSHError
	if errors.As(err, &apiErr) {
		switch apiErr.StatusCode {
		case 400:
			return "delete_http_400"
		case 403:
			return "delete_http_403"
		case 404:
			return "delete_http_404"
		case 429:
			return "delete_http_429"
		default:
			return "delete_api_error"
		}
	}
	return "delete_error_unclassified"
}
