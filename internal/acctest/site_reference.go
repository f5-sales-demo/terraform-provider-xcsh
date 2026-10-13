// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"
)

type readOnlySiteReference struct {
	Name          string `json:"name"`
	Namespace     string `json:"namespace"`
	SelectedSites int    `json:"selected_sites,omitempty"`
}

func discoverReadOnlySite(ctx context.Context, c *client.Client) (readOnlySiteReference, error) {
	return discoverReadOnlyInfrastructure(ctx, c, "sites")
}
func discoverReadOnlyInfrastructure(ctx context.Context, c *client.Client, collection string) (readOnlySiteReference, error) {
	return discoverInfrastructureWithEligibility(ctx, c, collection, false)
}
func discoverMeshVirtualSite(ctx context.Context, c *client.Client) (readOnlySiteReference, error) {
	return discoverInfrastructureWithEligibility(ctx, c, "virtual_sites", true)
}
func discoverInfrastructureWithEligibility(ctx context.Context, c *client.Client, collection string, meshOnly bool) (readOnlySiteReference, error) {
	var listing struct {
		Items []struct {
			Name      string          `json:"name"`
			Namespace string          `json:"namespace"`
			Metadata  client.Metadata `json:"metadata"`
		} `json:"items"`
	}
	path := "/api/config/namespaces/system/" + collection
	if err := c.Get(ctx, path, &listing); err != nil {
		return readOnlySiteReference{}, err
	}
	if listing.Items == nil {
		return readOnlySiteReference{}, errors.New("infrastructure inventory malformed")
	}
	names := []string{}
	refs := map[string]readOnlySiteReference{}
	for _, item := range listing.Items {
		name := item.Name
		if name == "" {
			name = item.Metadata.Name
		}
		if name == "" || strings.ContainsAny(name, "/?# \t\r\n") {
			return readOnlySiteReference{}, errors.New("infrastructure inventory malformed")
		}
		if !strings.HasPrefix(name, TestResourcePrefix) {
			namespace := item.Namespace
			if namespace == "" {
				namespace = item.Metadata.Namespace
			}
			if namespace == "" {
				namespace = "system"
			}
			if strings.ContainsAny(namespace, "/?# \t\r\n") {
				return readOnlySiteReference{}, errors.New("infrastructure inventory malformed")
			}
			key := namespace + "/" + name
			names = append(names, key)
			refs[key] = readOnlySiteReference{Name: name, Namespace: namespace}
		}
	}
	if len(names) == 0 {
		return readOnlySiteReference{}, errors.New("existing infrastructure absent")
	}
	sort.Strings(names)
	for _, key := range names {
		ref := refs[key]
		_, err := readOwnershipProjection(ctx, c, "/api/config/namespaces/"+ref.Namespace+"/"+collection+"/"+ref.Name, ref.Namespace, ref.Name)
		if authoritativeOwnedAbsence(err) {
			continue
		}
		if err != nil {
			return readOnlySiteReference{}, err
		}
		if meshOnly {
			var object client.VirtualSite
			if err := c.Get(ctx, "/api/config/namespaces/"+ref.Namespace+"/virtual_sites/"+ref.Name, &object); err != nil {
				return readOnlySiteReference{}, err
			}
			typ, _ := object.Spec["site_type"].(string)
			if typ == "" {
				return readOnlySiteReference{}, errors.New("infrastructure inventory malformed")
			}
			if typ != "CUSTOMER_EDGE" {
				continue
			}
			selector, _ := object.Spec["site_selector"].(map[string]interface{})
			expressions, _ := selector["expressions"].([]interface{})
			if len(expressions) == 0 {
				continue
			}
			count, err := verifyMeshSelectorExpansion(ctx, c, ref)
			if err != nil {
				return readOnlySiteReference{}, err
			}
			if count == 0 {
				continue
			}
			ref.SelectedSites = count
		}
		return ref, nil
	}
	return readOnlySiteReference{}, errors.New("existing infrastructure absent")
}

func requireReadOnlySite(t *testing.T) { requireReadOnlyReference(t, "site", "sites") }
func RequireReadOnlyVirtualSite(t *testing.T) {
	t.Helper()
	if IsMockMode() {
		return
	}
	requireReadOnlyReference(t, "virtual_site", "virtual_sites")
}
func requireReadOnlyReference(t *testing.T, kind, collection string) {
	t.Helper()
	c, err := GetTestClient()
	if err != nil {
		t.Fatal("infrastructure client unavailable")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var ref readOnlySiteReference
	var errDiscover error
	if kind == "virtual_site" {
		ref, errDiscover = discoverMeshVirtualSite(ctx, c)
	} else {
		ref, errDiscover = discoverReadOnlyInfrastructure(ctx, c, collection)
	}
	err = errDiscover
	if err != nil {
		if classifyCapabilityProbe(err) == capabilityDenied {
			t.Logf("[PREREQUISITE:operation_denied method=GET type=%s status=403]", kind)
			t.Skip("existing site read denied")
		}
		if err.Error() == "existing infrastructure absent" {
			t.Logf("[PREREQUISITE:infrastructure_absent method=GET type=%s]", kind)
			t.Skip("no independently owned site available")
		}
		t.Fatal("existing site reference readback failed")
	}
	private := os.Getenv("XCSH_ACC_PRIVATE")
	if private == "" {
		t.Fatal("private site reference workspace missing")
	}
	data, err := json.Marshal(ref)
	if err != nil {
		t.Fatal("site reference encoding failed")
	}
	if err = os.WriteFile(filepath.Join(private, kind+"-reference.json"), data, 0600); err != nil {
		t.Fatal("private site reference persistence failed")
	}
	t.Logf("[PREREQUISITE:infrastructure_verified method=GET type=%s]", kind)
}

// ReadOnlySiteName returns a verified private reference. It never imports or
// journals the appliance. Fixture audits supply their own synthetic argument.
func ReadOnlySiteName() string                 { return readOnlyReferenceName("site") }
func ReadOnlyVirtualSiteName() string          { return readOnlyReference("virtual_site").Name }
func ReadOnlyVirtualSiteNamespace() string     { return readOnlyReference("virtual_site").Namespace }
func readOnlyReferenceName(kind string) string { return readOnlyReference(kind).Name }
func readOnlyReference(kind string) readOnlySiteReference {
	ref := readOnlySiteReference{Name: "example-site", Namespace: "system"}
	if !IsMockMode() {
		data, err := os.ReadFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), kind+"-reference.json"))
		if err != nil || json.Unmarshal(data, &ref) != nil || ref.Namespace == "" || ref.Name == "" {
			panic("verified site reference unavailable")
		}
	}
	return ref
}

// verifyMeshSelectorExpansion uses the API's resolved selection and exact site
// readback. Empty successful selections are absent; uncertain expansion fails.
func verifyMeshSelectorExpansion(ctx context.Context, c *client.Client, ref readOnlySiteReference) (int, error) {
	var expansion struct {
		Items []struct {
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
			Kind      string `json:"kind"`
		} `json:"items"`
		Error json.RawMessage `json:"error"`
	}
	path := "/api/config/namespaces/" + ref.Namespace + "/virtual_sites/" + ref.Name + "/selectees"
	if err := c.Get(ctx, path, &expansion); err != nil {
		return 0, err
	}
	if expansion.Items == nil || (len(expansion.Error) > 0 && string(expansion.Error) != "null" && string(expansion.Error) != "{}") {
		return 0, errors.New("mesh selector expansion malformed")
	}
	seen := map[string]bool{}
	for _, item := range expansion.Items {
		if item.Name == "" || item.Namespace == "" || strings.ContainsAny(item.Name+item.Namespace, "/?# \t\r\n") || (item.Kind != "site" && !strings.HasSuffix(item.Kind, ".site")) {
			return 0, errors.New("mesh selector site identity malformed")
		}
		key := item.Namespace + "/" + item.Name
		if seen[key] {
			return 0, errors.New("mesh selector duplicate site")
		}
		seen[key] = true
		if _, err := readOwnershipProjection(ctx, c, "/api/config/namespaces/"+item.Namespace+"/sites/"+item.Name, item.Namespace, item.Name); err != nil {
			// A selected stale site is an inconsistent expansion, never an absent selector.
			return 0, errors.New("mesh selector selected-site readback failed")
		}
	}
	return len(expansion.Items), nil
}
