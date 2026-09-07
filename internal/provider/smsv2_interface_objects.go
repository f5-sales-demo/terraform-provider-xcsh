// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"fmt"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

// resolveSMSv2InterfaceObjects joins the configured MAC to a realized object
// through site ownership, node, and ethernet device. Object existence is not
// evidence of link health; configuration objects do not expose that telemetry.
func resolveSMSv2InterfaceObjects(configuration client.SMSv2Observation, configured []smsv2ConfiguredInterface, observation client.SMSv2Observation) ([]smsv2ConfiguredInterface, error) {
	metadata, _ := nestedMap(map[string]interface{}(configuration), "metadata")
	system, _ := nestedMap(map[string]interface{}(configuration), "system_metadata")
	site, namespace, uid := stringField(metadata, "name"), stringField(metadata, "namespace"), stringField(system, "uid")
	if site == "" || namespace == "" || uid == "" {
		return nil, fmt.Errorf("interface discovery requires the site's name, namespace and UID")
	}
	if errors, exists := observation["errors"]; exists && errors != nil {
		entries, ok := errors.([]interface{})
		if !ok || len(entries) != 0 {
			return nil, fmt.Errorf("network_interface discovery returned partial errors")
		}
	}
	items, ok := observation["items"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("network_interface discovery has no items array")
	}
	result := append([]smsv2ConfiguredInterface(nil), configured...)
	for index, expected := range result {
		matches := 0
		for _, raw := range items {
			item, ok := raw.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("network_interface discovery contains a malformed item")
			}
			owner, _ := nestedMap(item, "owner_view")
			if stringField(owner, "kind") != "securemesh_site_v2" || stringField(owner, "name") != site || stringField(owner, "namespace") != namespace || stringField(owner, "uid") != uid {
				continue
			}
			ethernet, _ := nestedMap(item, "get_spec", "ethernet_interface")
			if stringField(ethernet, "node") != expected.Node || stringField(ethernet, "device") != expected.Device {
				continue
			}
			name := stringField(item, "name")
			_, slo := ethernet["site_local_network"]
			_, sli := ethernet["site_local_inside_network"]
			if name == "" || stringField(item, "namespace") != namespace || slo == sli || (expected.Role == "slo") != slo || int64Field(ethernet, "mtu") != expected.MTU {
				return nil, fmt.Errorf("realized network_interface disagrees with configured node %q device %q", expected.Node, expected.Device)
			}
			result[index].Name = name
			matches++
		}
		if matches != 1 {
			return nil, fmt.Errorf("node %q device %q resolved to %d owned network_interface objects", expected.Node, expected.Device, matches)
		}
	}
	return result, nil
}
