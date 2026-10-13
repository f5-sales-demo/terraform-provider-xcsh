// Copyright (c) 2026 Robin Mordasiewicz. MIT License.
package acctest

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
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
			allow := map[string]bool{}
			for _, word := range strings.Fields("message code details failed object for error no in namespace system not dns label must be longer than characters segment configuration wrong kind name reference tenant required missing must specify config network resolver found get store timeout unknown invalid is a valid field value create unavailable interface constraints") {
				allow[word] = true
			}
			words := []string{}
			for _, word := range regexp.MustCompile(`[a-z][a-z0-9_]*`).FindAllString(strings.ToLower(raw), -1) {
				if allow[word] {
					words = append(words, word)
				} else {
					words = append(words, "_")
				}
			}
			record["words"] = words
		}
		records = append(records, record)
	}
	facts["segment_variants"] = records
	var sites struct {
		Items []struct {
			Name, Namespace string
			Metadata        client.Metadata
		}
	}
	if err = c.Get(ctx, "/api/config/namespaces/system/sites", &sites); err != nil {
		return errors.New("site inventory failed")
	}
	readOK, eligible := 0, 0
	for _, item := range sites.Items {
		name := item.Name
		ns := item.Namespace
		if name == "" {
			name = item.Metadata.Name
		}
		if ns == "" {
			ns = item.Metadata.Namespace
		}
		if ns == "" {
			ns = "system"
		}
		if strings.HasPrefix(name, TestResourcePrefix) {
			continue
		}
		var object struct {
			Metadata       client.Metadata
			Spec           map[string]interface{}
			SystemMetadata struct {
				UID string `json:"uid"`
			} `json:"system_metadata"`
		}
		readErr := c.Get(ctx, "/api/config/namespaces/"+ns+"/sites/"+name, &object)
		if readErr != nil {
			continue
		}
		readOK++
		typ, _ := object.Spec["site_type"].(string)
		if typ == "CUSTOMER_EDGE" {
			eligible++
			data, _ := json.Marshal(map[string]interface{}{"metadata": object.Metadata, "spec": object.Spec})
			if writeErr := os.WriteFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), "eligible-site.json"), data, 0600); writeErr != nil {
				return writeErr
			}
		}
	}
	facts["sites_read_ok"] = readOK
	facts["customer_edge_sites"] = eligible
	var inventory struct {
		Items []struct {
			Name, Namespace string
			Metadata        client.Metadata
		}
	}
	if err = c.Get(ctx, "/api/config/namespaces/system/virtual_sites", &inventory); err != nil {
		return errors.New("virtual inventory failed")
	}
	selectorFacts := []map[string]interface{}{}
	for _, item := range inventory.Items {
		name := item.Name
		ns := item.Namespace
		if name == "" {
			name = item.Metadata.Name
		}
		if ns == "" {
			ns = item.Metadata.Namespace
		}
		if ns == "" {
			ns = "system"
		}
		if strings.HasPrefix(name, TestResourcePrefix) {
			continue
		}
		var vs client.VirtualSite
		if readErr := c.Get(ctx, "/api/config/namespaces/"+ns+"/virtual_sites/"+name, &vs); readErr != nil {
			continue
		}
		if vs.Spec["site_type"] != "CUSTOMER_EDGE" {
			continue
		}
		var expansion struct {
			Items []map[string]interface{}
			Error json.RawMessage
		}
		readErr := c.Get(ctx, "/api/config/namespaces/"+ns+"/virtual_sites/"+name+"/selectees", &expansion)
		fact := map[string]interface{}{"selected": len(expansion.Items), "complete": readErr == nil && expansion.Items != nil, "system_namespace": ns == "system"}
		selectorFacts = append(selectorFacts, fact)
	}
	facts["customer_edge_selectors"] = selectorFacts
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
	facts["chosen_selector_system_namespace"] = ref.Namespace == "system"
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
	var eligibleSite struct {
		Metadata client.Metadata
		Spec     map[string]interface{}
	}
	if data, readErr := os.ReadFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), "eligible-site.json")); readErr == nil && json.Unmarshal(data, &eligibleSite) == nil {
		selectorName := CompactRandomName(20)
		expression := "ves.io/siteName in (" + eligibleSite.Metadata.Name + ")"
		_, createErr := c.CreateVirtualSite(ctx, &client.VirtualSite{Metadata: client.Metadata{Name: selectorName, Namespace: "system"}, Spec: map[string]interface{}{
			"site_type": "CUSTOMER_EDGE", "site_selector": map[string]interface{}{"expressions": []string{expression}},
		}})
		facts["disposable_selector_created"] = createErr == nil
		if createErr == nil {
			var expansion struct{ Items []map[string]interface{} }
			readErr := c.Get(ctx, "/api/config/namespaces/system/virtual_sites/"+selectorName+"/selectees", &expansion)
			facts["disposable_selector_read_complete"] = readErr == nil && expansion.Items != nil
			facts["disposable_selector_selected_count"] = len(expansion.Items)
			_, meshErr := c.CreateSiteMeshGroup(ctx, &client.SiteMeshGroup{Metadata: client.Metadata{Name: RandomName("tf-acc-test-selector-mesh"), Namespace: "system"}, Spec: map[string]interface{}{
				"bfd_disabled": map[string]interface{}{}, "disable_re_fallback": map[string]interface{}{},
				"virtual_site": []interface{}{map[string]interface{}{"name": selectorName, "namespace": "system"}}, "full_mesh": map[string]interface{}{"control_and_data_plane_mesh": map[string]interface{}{}},
			}})
			facts["disposable_selector_mesh_created"] = meshErr == nil
			var apiErr *xcsherrors.XCSHError
			if errors.As(meshErr, &apiErr) {
				facts["disposable_selector_mesh_status"] = apiErr.StatusCode
			}
		}
	}
	encoded, err := json.MarshalIndent(facts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(os.Getenv("XCSH_ACC_PRIVATE"), "reference-resolution-facts.json"), encoded, 0600)
}
