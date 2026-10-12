package schema

import "github.com/f5-sales-demo/terraform-provider-xcsh/tools/pkg/openapi"

// measuredNestedDefault fills omissions in pinned enrichment with observed API
// behavior (#2535). Match the parent protobuf identity, rather than a field name
// or Terraform list position, so shared components cover every resource branch.
// The input is copied; upstream defaults and requiredness remain authoritative.
func measuredNestedDefault(parent openapi.Schema, name string, property openapi.Schema) openapi.Schema {
	var value any
	switch parent.XVesProtoMessage {
	case "ves.io.schema.views.cdn_loadbalancer.CdnOriginPoolType":
		if name == "origin_request_timeout" {
			value = "60s"
		}
	case "ves.io.schema.views.workload.ContainerType":
		if name == "init_container" {
			value = false
		}
	case "ves.io.schema.views.workload.ImageType":
		if name == "pull_policy" {
			value = "IMAGE_PULL_POLICY_DEFAULT"
		}
	}
	if value != nil {
		property.XF5XCServerDefault = true
		if property.Default == nil {
			property.Default = value
		}
	}
	return property
}
