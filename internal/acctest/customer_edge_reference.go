// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

var errEligibleInfrastructureAbsent = errors.New("eligible existing infrastructure absent")

func discoverCustomerEdgeSites(ctx context.Context, c *client.Client) ([]readOnlySiteReference, error) {
	var list struct {
		Items []struct {
			Name      string          `json:"name"`
			Namespace string          `json:"namespace"`
			Metadata  client.Metadata `json:"metadata"`
		} `json:"items"`
	}
	if err := c.Get(ctx, "/api/config/namespaces/system/sites", &list); err != nil {
		return nil, err
	}
	if list.Items == nil {
		return nil, errors.New("site inventory malformed")
	}
	refs := []readOnlySiteReference{}
	for _, item := range list.Items {
		name := item.Name
		if name == "" {
			name = item.Metadata.Name
		}
		ns := item.Namespace
		if ns == "" {
			ns = item.Metadata.Namespace
		}
		if ns == "" {
			ns = "system"
		}
		if name == "" || strings.ContainsAny(name+ns, "/ ?#\t\r\n") {
			return nil, errors.New("site identity malformed")
		}
		if !strings.HasPrefix(name, TestResourcePrefix) {
			refs = append(refs, readOnlySiteReference{Name: name, Namespace: ns})
		}
	}
	sort.Slice(refs, func(i, j int) bool { return refs[i].Namespace+"/"+refs[i].Name < refs[j].Namespace+"/"+refs[j].Name })
	eligible := []readOnlySiteReference{}
	for _, ref := range refs {
		path := "/api/config/namespaces/" + ref.Namespace + "/sites/" + ref.Name
		var object struct {
			Metadata       client.Metadata `json:"metadata"`
			SystemMetadata struct {
				UID string `json:"uid"`
			} `json:"system_metadata"`
			Spec struct {
				Type string `json:"site_type"`
			} `json:"spec"`
		}
		if err := c.Get(ctx, path, &object); err != nil {
			if authoritativeOwnedAbsence(err) {
				continue
			}
			return nil, err
		}
		uid := object.Metadata.UID
		if uid == "" {
			uid = object.SystemMetadata.UID
		}
		if object.Metadata.Name != ref.Name || uid == "" || (object.Metadata.Namespace != "" && object.Metadata.Namespace != ref.Namespace) || object.Spec.Type == "" {
			return nil, errors.New("site type or identity malformed")
		}
		if object.Spec.Type == "CUSTOMER_EDGE" {
			eligible = append(eligible, ref)
		}
	}
	if len(eligible) == 0 {
		return nil, errEligibleInfrastructureAbsent
	}
	return eligible, nil
}
func RequireCustomerEdgeSite(t *testing.T) {
	t.Helper()
	SkipIfNotAccTest(t)
	if IsMockMode() {
		return
	}
	c, err := GetTestClient()
	if err != nil {
		t.Fatal("site client unavailable")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	refs, err := discoverCustomerEdgeSites(ctx, c)
	if err != nil {
		if errors.Is(err, errEligibleInfrastructureAbsent) {
			t.Log("[PREREQUISITE:infrastructure_absent method=GET type=customer_edge_site]")
			t.Skip("typed site inventory contains no customer edge")
		}
		t.Fatal("customer edge inventory unverified")
	}
	if err := saveSiteReference(refs[0]); err != nil {
		t.Fatal("site reference receipt unavailable")
	}
	t.Log("[PREREQUISITE:infrastructure_verified method=GET type=customer_edge_site]")
}
func saveSiteReference(ref readOnlySiteReference) error {
	data, err := json.Marshal(ref)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), "site-reference.json"), data, 0600)
}
