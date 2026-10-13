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
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func discoverReadOnlySite(ctx context.Context, c *client.Client) (readOnlySiteReference, error) {
	return discoverReadOnlyInfrastructure(ctx, c, "sites")
}
func discoverReadOnlyInfrastructure(ctx context.Context, c *client.Client, collection string) (readOnlySiteReference, error) {
	var listing struct {
		Items []struct {
			Name     string          `json:"name"`
			Metadata client.Metadata `json:"metadata"`
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
	for _, item := range listing.Items {
		name := item.Name
		if name == "" {
			name = item.Metadata.Name
		}
		if name == "" || strings.ContainsAny(name, "/?# \t\r\n") {
			return readOnlySiteReference{}, errors.New("infrastructure inventory malformed")
		}
		if !strings.HasPrefix(name, TestResourcePrefix) {
			names = append(names, name)
		}
	}
	if len(names) == 0 {
		return readOnlySiteReference{}, errors.New("existing infrastructure absent")
	}
	sort.Strings(names)
	for _, name := range names {
		_, err := readOwnershipProjection(ctx, c, path+"/"+name, "system", name)
		if authoritativeOwnedAbsence(err) {
			continue
		}
		if err != nil {
			return readOnlySiteReference{}, err
		}
		return readOnlySiteReference{Name: name, Namespace: "system"}, nil
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
	ref, err := discoverReadOnlyInfrastructure(ctx, c, collection)
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
func ReadOnlySiteName() string        { return readOnlyReferenceName("site") }
func ReadOnlyVirtualSiteName() string { return readOnlyReferenceName("virtual_site") }
func readOnlyReferenceName(kind string) string {
	ref := readOnlySiteReference{Name: "example-site", Namespace: "system"}
	if !IsMockMode() {
		data, err := os.ReadFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), kind+"-reference.json"))
		if err != nil || json.Unmarshal(data, &ref) != nil || ref.Namespace != "system" || ref.Name == "" {
			panic("verified site reference unavailable")
		}
	}
	return ref.Name
}
