// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

const ownershipAnnotation = "f5-sales-demo.com/acceptance-run"

var activeOwnership struct {
	sync.Mutex
	journal *OwnershipJournal
	err     error
	loaded  bool
}

func acceptanceJournal() (*OwnershipJournal, error) {
	activeOwnership.Lock()
	defer activeOwnership.Unlock()
	if !activeOwnership.loaded {
		activeOwnership.loaded = true
		activeOwnership.journal, activeOwnership.err = LoadOwnershipJournal(os.Getenv("XCSH_ACC_JOURNAL"), os.Getenv("XCSH_ACC_RUN_ID"))
	}
	return activeOwnership.journal, activeOwnership.err
}

type ownedProvider struct{ fwprovider.Provider }

func (p *ownedProvider) Actions(ctx context.Context) []func() action.Action {
	return p.Provider.(fwprovider.ProviderWithActions).Actions(ctx)
}
func (p *ownedProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return p.Provider.(fwprovider.ProviderWithEphemeralResources).EphemeralResources(ctx)
}

type ownershipReadContextKey struct{}

func (p *ownedProvider) Configure(ctx context.Context, req fwprovider.ConfigureRequest, resp *fwprovider.ConfigureResponse) {
	p.Provider.Configure(ctx, req, resp)
	if resp.Diagnostics.HasError() || os.Getenv("XCSH_ACC_JOURNAL") == "" {
		return
	}
	c, ok := resp.ResourceData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Ownership tracking unavailable", "Acceptance client missing.")
		return
	}
	if err := attachOwnership(c); err != nil {
		resp.Diagnostics.AddError("Ownership tracking unavailable", "Private ownership journal cannot be loaded.")
	}
}
func attachOwnership(c *client.Client) error {
	if os.Getenv("XCSH_ACC_JOURNAL") == "" {
		return nil
	}
	if _, ok := c.HTTPClient.Transport.(*ownershipTransport); ok {
		return nil
	}
	j, err := acceptanceJournal()
	if err != nil {
		return err
	}
	base := c.HTTPClient.Transport
	if base == nil {
		base = http.DefaultTransport
	}
	c.HTTPClient.Transport = &ownershipTransport{base: base, c: c, j: j}
	return nil
}

type ownershipTransport struct {
	base http.RoundTripper
	c    *client.Client
	j    *OwnershipJournal
}

func (t *ownershipTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method == http.MethodPut && strings.HasPrefix(req.URL.Path, "/api/object_store/") {
		return t.createSwagger(req)
	}
	if req.Method == http.MethodPut {
		route := false
		for _, collection := range ownedCollectionRoutes {
			expected := collection
			if strings.Contains(expected, "%s") {
				parts := strings.Split(req.URL.Path, "/")
				for i, part := range parts {
					if part == "namespaces" && i+1 < len(parts) {
						expected = strings.ReplaceAll(expected, "%s", parts[i+1])
						break
					}
				}
			}
			if strings.HasPrefix(req.URL.Path, expected+"/") {
				route = true
				break
			}
		}
		if route && req.Body != nil {
			data, err := io.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}
			_ = req.Body.Close()
			var payload map[string]interface{}
			if err := json.Unmarshal(data, &payload); err != nil {
				return nil, errors.New("invalid acceptance update request")
			}
			metadata, ok := payload["metadata"].(map[string]interface{})
			if !ok {
				return nil, errors.New("acceptance update metadata missing")
			}
			name, _ := metadata["name"].(string)
			namespace, _ := metadata["namespace"].(string)
			var owned OwnedResource
			found := false
			t.j.mu.Lock()
			for _, existing := range t.j.Resources {
				if existing.Name == name && existing.Namespace == namespace {
					expected := strings.ReplaceAll(ownedCollectionRoutes[existing.Type], "%s", namespace) + "/" + name
					if expected == req.URL.Path {
						owned = existing
						found = true
						break
					}
				}
			}
			t.j.mu.Unlock()
			if !found || owned.UID == "" || owned.Status != "created" {
				return nil, errors.New("update target is not journal-owned")
			}
			observed, err := (apiOwnershipBackend{c: t.c}).Read(req.Context(), owned)
			if err != nil || ownershipUIDDigest(observed.UID) != owned.UID {
				return nil, errors.New("update ownership readback failed")
			}
			annotations, _ := metadata["annotations"].(map[string]interface{})
			if annotations == nil {
				annotations = map[string]interface{}{}
			}
			annotations[ownershipAnnotation] = t.j.RunID
			metadata["annotations"] = annotations
			data, err = json.Marshal(payload)
			if err != nil {
				return nil, err
			}
			clone := req.Clone(req.Context())
			clone.Body = io.NopCloser(bytes.NewReader(data))
			clone.ContentLength = int64(len(data))
			clone.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(data)), nil }
			return t.base.RoundTrip(clone)
		}
	}
	if req.Method != http.MethodPost {
		response, err := t.base.RoundTrip(req)
		if err != nil || req.Method != http.MethodGet || req.Context().Value(ownershipReadContextKey{}) == true || response.StatusCode < 200 || response.StatusCode >= 300 {
			return response, err
		}
		data, err := io.ReadAll(response.Body)
		_ = response.Body.Close()
		if err != nil {
			return nil, err
		}
		var payload interface{}
		if json.Unmarshal(data, &payload) == nil {
			stripOwnershipAnnotation(payload)
			data, err = json.Marshal(payload)
			if err != nil {
				return nil, err
			}
		}
		response.Body = io.NopCloser(bytes.NewReader(data))
		response.ContentLength = int64(len(data))
		response.Header.Set("Content-Length", strconv.Itoa(len(data)))
		return response, nil
	}
	path := req.URL.Path
	if req.Body == nil {
		return nil, errors.New("creation request body missing")
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	_ = req.Body.Close()
	req.Body = io.NopCloser(bytes.NewReader(data))
	var document struct {
		Metadata client.Metadata        `json:"metadata"`
		Spec     map[string]interface{} `json:"spec"`
	}
	if err = json.Unmarshal(data, &document); err != nil {
		return nil, errors.New("invalid acceptance creation request")
	}
	if document.Metadata.Name == "-" {
		return t.base.RoundTrip(req)
	}
	resourceType := ""
	for typ, route := range ownedCollectionRoutes {
		ns := document.Metadata.Namespace
		expected := route
		if strings.Contains(expected, "%s") {
			expected = strings.ReplaceAll(expected, "%s", ns)
		}
		if expected == path {
			resourceType = typ
			break
		}
	}
	if resourceType == "" {
		return nil, errors.New("no ownership handler registered for create operation")
	}
	r := OwnedResource{Type: resourceType, Name: document.Metadata.Name, Namespace: document.Metadata.Namespace}
	if err = t.j.validate(r); err != nil {
		return nil, err
	}
	backend := apiOwnershipBackend{c: t.c}
	if _, err = backend.Read(req.Context(), r); !errors.Is(err, errOwnedAbsent) {
		return nil, errors.New("create target is not authoritatively absent")
	}
	t.j.mu.Lock()
	for _, existing := range t.j.Resources {
		// Exact references and namespace parents define deletion dependencies.
		if existing.Name == r.Namespace || containsOwnedReference(document.Spec, existing.Name) {
			r.Dependencies = append(r.Dependencies, ownedKey(existing))
		}
	}
	t.j.mu.Unlock()
	if err = t.j.ResetAfterVerifiedAbsence(r); err != nil {
		return nil, err
	}
	if err = t.j.Intent(r); err != nil {
		return nil, err
	}
	var payload map[string]interface{}
	if err = json.Unmarshal(data, &payload); err != nil {
		return nil, err
	}
	metadata, ok := payload["metadata"].(map[string]interface{})
	if !ok {
		return nil, errors.New("acceptance metadata missing")
	}
	annotations, ok := metadata["annotations"].(map[string]interface{})
	if !ok {
		annotations = map[string]interface{}{}
	}
	annotations[ownershipAnnotation] = t.j.RunID
	metadata["annotations"] = annotations
	data, err = json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	clone := req.Clone(req.Context())
	clone.Body = io.NopCloser(bytes.NewReader(data))
	clone.ContentLength = int64(len(data))
	clone.GetBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(data)), nil }
	response, err := t.base.RoundTrip(clone)
	logAcceptanceAPIDiagnostic(response)
	if err != nil {
		return response, err
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		metadata, readErr := backend.Read(req.Context(), r)
		if readErr != nil || metadata.Annotations[ownershipAnnotation] != t.j.RunID {
			_ = response.Body.Close()
			return nil, errors.New("creation ownership readback failed")
		}
		if err = t.j.Confirm(r, metadata.UID); err != nil {
			_ = response.Body.Close()
			return nil, errors.New("creation ownership journal confirmation failed")
		}
	}
	return response, nil
}
func containsOwnedReference(value interface{}, name string) bool {
	switch v := value.(type) {
	case string:
		if v == name {
			return true
		}
		for _, component := range strings.Split(v, "/") {
			if component == name {
				return true
			}
		}
		return false
	case map[string]interface{}:
		for _, child := range v {
			if containsOwnedReference(child, name) {
				return true
			}
		}
	case []interface{}:
		for _, child := range v {
			if containsOwnedReference(child, name) {
				return true
			}
		}
	}
	return false
}

func stripOwnershipAnnotation(value interface{}) {
	switch v := value.(type) {
	case map[string]interface{}:
		if annotations, ok := v["annotations"].(map[string]interface{}); ok {
			delete(annotations, ownershipAnnotation)
		}
		for _, child := range v {
			stripOwnershipAnnotation(child)
		}
	case []interface{}:
		for _, child := range v {
			stripOwnershipAnnotation(child)
		}
	}
}
