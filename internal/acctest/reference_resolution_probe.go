// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
	xcsherrors "github.com/f5-sales-demo/terraform-provider-xcsh/internal/errors"
)

func ProbeReferenceResolution(ctx context.Context) error {
	c, err := GetTestClient()
	if err != nil {
		return errors.New("probe client unavailable")
	}
	c.MaxRetries = 0
	facts := map[string]interface{}{"version": 1}
	segmentName := CompactRandomName(20)
	_, err = c.CreateSegment(ctx, &client.Segment{Metadata: client.Metadata{Name: segmentName, Namespace: "system"}, Spec: map[string]interface{}{}})
	if err != nil {
		return errors.New("probe segment creation failed")
	}
	var segment map[string]interface{}
	if err = c.Get(ctx, "/api/config/namespaces/system/segments/"+segmentName, &segment); err != nil {
		return errors.New("probe segment readback failed")
	}
	metadata, _ := segment["metadata"].(map[string]interface{})
	facts["segment_name_matches"] = metadata["name"] == segmentName
	tenant := ""
	for _, parent := range []string{"metadata", "system_metadata"} {
		fields, _ := segment[parent].(map[string]interface{})
		if value, ok := fields["tenant"].(string); ok && value != "" {
			tenant = value
		}
	}
	if tenant == "" {
		u, _ := url.Parse(os.Getenv(EnvXCSHURL))
		tenant = strings.Split(u.Hostname(), ".")[0]
	}
	facts["tenant_available"] = tenant != ""
	records := []map[string]interface{}{}
	for _, variant := range []string{"nested", "nested_tenant", "name_only", "qualified_name", "segment_config"} {
		ref := map[string]interface{}{"name": segmentName, "namespace": "system"}
		if variant == "nested_tenant" {
			ref["tenant"] = tenant
		}
		if variant == "name_only" {
			delete(ref, "namespace")
		}
		if variant == "qualified_name" {
			ref["name"] = tenant + "/system/" + segmentName
			delete(ref, "namespace")
		}
		entry := map[string]interface{}{"segment_network": ref}
		if variant == "segment_config" {
			entry["segment_config"] = map[string]interface{}{}
		}
		_, createErr := c.CreateSecuremeshSiteV2(ctx, &client.SecuremeshSiteV2{Metadata: client.Metadata{Name: RandomName("tf-acc-test-reference"), Namespace: "system"}, Spec: map[string]interface{}{
			"baremetal":  map[string]interface{}{"not_managed": map[string]interface{}{}},
			"disable_ha": map[string]interface{}{}, "no_network_policy": map[string]interface{}{}, "no_forward_proxy": map[string]interface{}{}, "logs_streaming_disabled": map[string]interface{}{}, "block_all_services": map[string]interface{}{},
			"segment_vrf": []interface{}{entry},
		}})
		record := map[string]interface{}{"variant": variant, "accepted": createErr == nil}
		var apiErr *xcsherrors.XCSHError
		if errors.As(createErr, &apiErr) {
			record["status"] = apiErr.StatusCode
			raw, _ := apiErr.Details["raw_response"].(string)
			record["mentions_segment_name"] = strings.Contains(raw, segmentName)
			record["mentions_qualified_reference"] = strings.Contains(raw, tenant+"/system/"+segmentName)
			record["invalid_dns_label"] = strings.Contains(strings.ToLower(raw), "dns label")
		}
		records = append(records, record)
	}
	facts["segment_variants"] = records
	ref, err := discoverMeshVirtualSite(ctx, c)
	if err != nil {
		return errors.New("probe virtual site discovery failed")
	}
	var selectees struct {
		Items []struct{ Name, Namespace, Tenant, Kind string }
		Error json.RawMessage
	}
	if err = c.Get(ctx, "/api/config/namespaces/"+ref.Namespace+"/virtual_sites/"+ref.Name+"/selectees", &selectees); err != nil {
		return errors.New("probe selector expansion failed")
	}
	facts["selectees_present"] = selectees.Items != nil
	facts["selected_count"] = len(selectees.Items)
	facts["expansion_error_present"] = len(selectees.Error) > 0 && string(selectees.Error) != "null" && string(selectees.Error) != "{}"
	siteKinds := map[string]int{}
	for _, site := range selectees.Items {
		siteKinds[site.Kind]++
	}
	// No identities or arbitrary API kinds enter the public facts.
	facts["selected_sites"] = siteKinds["site"]
	records = []map[string]interface{}{}
	for _, variant := range []string{"nested", "nested_tenant", "qualified_name"} {
		r := map[string]interface{}{"name": ref.Name, "namespace": ref.Namespace}
		if variant == "nested_tenant" {
			r["tenant"] = tenant
		}
		if variant == "qualified_name" {
			r["name"] = tenant + "/" + ref.Namespace + "/" + ref.Name
			delete(r, "namespace")
		}
		_, createErr := c.CreateSiteMeshGroup(ctx, &client.SiteMeshGroup{Metadata: client.Metadata{Name: RandomName("tf-acc-test-reference"), Namespace: "system"}, Spec: map[string]interface{}{
			"bfd_disabled": map[string]interface{}{}, "disable_re_fallback": map[string]interface{}{},
			"virtual_site": []interface{}{r}, "full_mesh": map[string]interface{}{"control_and_data_plane_mesh": map[string]interface{}{}},
		}})
		record := map[string]interface{}{"variant": variant, "accepted": createErr == nil}
		var apiErr *xcsherrors.XCSHError
		if errors.As(createErr, &apiErr) {
			record["status"] = apiErr.StatusCode
		}
		records = append(records, record)
	}
	facts["mesh_variants"] = records
	encoded, err := json.MarshalIndent(facts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), "reference-resolution-facts.json"), encoded, 0600)
}
