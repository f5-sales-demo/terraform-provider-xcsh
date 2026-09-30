// Copyright (c) 2026 Robin Mordasiewicz. MIT License.

package provider

import (
	"fmt"

	"github.com/f5-sales-demo/terraform-provider-xcsh/internal/client"
)

// preserveRealizedKVMNodes keeps the platform-owned registration graph when
// Terraform deliberately configures an empty bootstrap KVM node list. The
// exact previously reviewed concurrency token is retained; a newer fresh read
// cannot silently contribute unreviewed nodes to the write.
func preserveRealizedKVMNodes(current, desired *client.SecuremeshSiteV2) error {
	configured, ok := nestedMap(desired.Spec, "kvm", "not_managed")
	if !ok {
		return nil
	}
	if nodes, exists := configured["node_list"]; exists {
		values, valid := nodes.([]interface{})
		if !valid {
			return fmt.Errorf("configured KVM node list is malformed")
		}
		if len(values) > 0 {
			return nil
		}
	}
	if current == nil || current.Metadata.Name != desired.Metadata.Name || current.Metadata.Namespace != desired.Metadata.Namespace {
		return fmt.Errorf("fresh KVM site read does not match the reviewed site identity")
	}
	if current.ResourceVersion == "" || current.ResourceVersion != desired.ResourceVersion {
		return fmt.Errorf("fresh KVM site read changed since the reviewed plan; refresh and review before update")
	}
	realized, ok := nestedMap(current.Spec, "kvm", "not_managed")
	if !ok {
		return fmt.Errorf("fresh site read does not contain the reviewed KVM provider")
	}
	nodes, exists := realized["node_list"]
	if !exists {
		return nil
	}
	values, valid := nodes.([]interface{})
	if !valid {
		return fmt.Errorf("realized KVM node list is malformed")
	}
	if len(values) == 0 {
		return nil
	}
	preserved := deepCopySMSv2Map(realized)
	for key, value := range configured {
		if key != "node_list" {
			preserved[key] = value
		}
	}
	desired.Spec["kvm"].(map[string]interface{})["not_managed"] = preserved
	return nil
}
