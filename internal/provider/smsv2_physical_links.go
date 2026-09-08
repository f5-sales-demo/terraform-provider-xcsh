// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"fmt"
	"sort"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

// validateSMSv2PhysicalLinks joins registered-site operational status to the
// already resolved configuration. Provisioning and object existence cannot
// establish that a physical interface is present and its link is up.
func validateSMSv2PhysicalLinks(configuration, observation client.SMSv2Observation, configured []smsv2ConfiguredInterface, bindings map[string]smsv2BindingModel) error {
	expected, _ := nestedMap(map[string]interface{}(configuration), "metadata")
	actual, _ := nestedMap(map[string]interface{}(observation), "metadata")
	if stringField(actual, "name") != stringField(expected, "name") || stringField(actual, "namespace") != stringField(expected, "namespace") || stringField(actual, "name") == "" {
		return fmt.Errorf("physical link status belongs to a different site or has incomplete identity")
	}
	rawStatuses, ok := observation["status"].([]interface{})
	if !ok {
		return fmt.Errorf("physical link status has no status array")
	}
	statuses := []map[string]interface{}{}
	for _, raw := range rawStatuses {
		status, ok := raw.(map[string]interface{})
		if !ok {
			return fmt.Errorf("physical link status contains a malformed record")
		}
		if status["ver_status"] == nil {
			continue
		}
		_, ok = nestedMap(status, "ver_status")
		if !ok {
			return fmt.Errorf("physical link status has a malformed interface-status document")
		}
		metadata, _ := nestedMap(status, "metadata")
		stale, known := metadata["vtrp_stale"].(bool)
		if stringField(metadata, "creator_class") != "ver" || stringField(metadata, "publish") != "STATUS_PUBLISH" || !known || stale {
			continue
		}
		node := stringField(metadata, "creator_id")
		if node == "" {
			return fmt.Errorf("physical link status has no node publisher identity")
		}
		statuses = append(statuses, status)
	}
	keys := make([]string, 0, len(bindings))
	for key := range bindings {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		binding := bindings[key]
		matches := []smsv2ConfiguredInterface{}
		for _, iface := range configured {
			if iface.Node == binding.Node.ValueString() && iface.MAC == binding.MAC.ValueString() {
				matches = append(matches, iface)
			}
		}
		if len(matches) != 1 {
			return fmt.Errorf("physical link binding %q does not resolve to one configured interface", key)
		}
		iface := matches[0]
		nodes := []map[string]interface{}{}
		for _, status := range statuses {
			publisher, _ := nestedMap(status, "metadata")
			if smsv2NodeMatches(iface.Node, stringField(publisher, "creator_id")) ||
				smsv2SitePublisherMatches(status, observation, iface.Node, stringField(expected, "name")) {
				ver, _ := nestedMap(status, "ver_status")
				nodes = append(nodes, ver)
			}
		}
		if len(nodes) != 1 {
			return fmt.Errorf("physical link binding %q resolved to %d published non-stale node observations", key, len(nodes))
		}
		if err := validateSMSv2PhysicalInterface(nodes[0], iface); err != nil {
			return fmt.Errorf("physical link binding %q: %w", key, err)
		}
	}
	return nil
}

func validateSMSv2PhysicalInterface(status map[string]interface{}, expected smsv2ConfiguredInterface) error {
	rawInterfaces, ok := status["intf_status"].([]interface{})
	if !ok {
		return fmt.Errorf("interface observations are missing")
	}
	matches := []map[string]interface{}{}
	for _, raw := range rawInterfaces {
		iface, ok := raw.(map[string]interface{})
		if !ok {
			return fmt.Errorf("interface observation is malformed")
		}
		if stringField(iface, "name") == expected.Device {
			matches = append(matches, iface)
		}
	}
	if len(matches) != 1 {
		return fmt.Errorf("device resolved to %d interface observations", len(matches))
	}
	got := matches[0]
	mac, err := normalizeSMSv2MAC(stringField(got, "mac"))
	if err != nil || mac != expected.MAC {
		return fmt.Errorf("device MAC disagrees with configured transport identity")
	}
	if stringField(got, "link_type") != "LINK_TYPE_ETHERNET" {
		return fmt.Errorf("device is not reported as an Ethernet link")
	}
	wantNetwork := map[string]string{"slo": "VIRTUAL_NETWORK_SITE_LOCAL", "sli": "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE"}[expected.Role]
	if wantNetwork == "" || stringField(got, "network_type") != wantNetwork {
		return fmt.Errorf("device network role disagrees with configured transport identity")
	}
	up, known := got["link_state"].(bool)
	if !known {
		return fmt.Errorf("link state is missing or malformed")
	}
	if !up {
		return fmt.Errorf("link is down")
	}
	return nil
}

// Current SMSv2 can publish node links under the site identity. Both node
// identifiers and the physical site's exact object reference must agree.
func smsv2SitePublisherMatches(status map[string]interface{}, observation client.SMSv2Observation, node, site string) bool {
	metadata, _ := nestedMap(status, "metadata")
	ver, _ := nestedMap(status, "ver_status")
	physical, _ := nestedMap(map[string]interface{}(observation), "system_metadata")
	uid := stringField(physical, "uid")
	if uid == "" || stringField(metadata, "creator_id") != site ||
		stringField(metadata, "status_id") != node+"_SiteStatusMgr" ||
		stringField(ver, "ver_instance_name") != node+"-"+site {
		return false
	}
	refs, ok := status["object_refs"].([]interface{})
	if !ok {
		return false
	}
	matches := 0
	for _, raw := range refs {
		ref, ok := raw.(map[string]interface{})
		if !ok {
			return false
		}
		if stringField(ref, "kind") == "ves.io.vega.cfg.site.Object" && stringField(ref, "uid") == uid {
			matches++
		}
	}
	return matches == 1
}
