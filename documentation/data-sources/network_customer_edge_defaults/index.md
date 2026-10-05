---
page_title: "xcsh_network_customer_edge_defaults"
subcategory: ""
description: "Default DNS and NTP destinations for Customer Edge firewall rules. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest."
xcsh_docs: {"aliases": ["network customer edge defaults"], "body_bytes": 1792, "body_sha256": "sha256:464b0cd3562fd4381bf75fc1efcee4c42ab5b3ca41da8f8cd5eda225e6cf1630", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_customer_edge_defaults:reference", "xcsh-docs:data-sources:network_customer_edge_defaults:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_customer_edge_defaults:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_customer_edge_defaults:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_customer_edge_defaults/index.md", "product": "distributed-cloud", "provider_name": "network_customer_edge_defaults", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2030202230323022-0201110111033011-3112302323000311-1231130201223133-0123022101003012-2310121012022212-0120311221210122-2230332213331100", "registry_path": "docs/data-sources/network_customer_edge_defaults.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_customer_edge_defaults/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Default DNS and NTP destinations for Customer Edge firewall rules. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_customer_edge_defaults

Breadcrumbs:

- xcsh_network_customer_edge_defaults

Default DNS and NTP destinations for Customer Edge firewall rules. Values are bundled from the
pinned OpenAPI release; this data source performs no network request. Ports and traffic direction
are not encoded in the manifest.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_customer_edge_defaults" "system_services" {}

output "customer_edge_default_egress" {
  value = {
    dns = {
      direction    = "egress"
      protocols    = ["udp", "tcp"]
      port         = 53
      destinations = data.xcsh_network_customer_edge_defaults.system_services.dns_servers
    }
    ntp = {
      direction    = "egress"
      protocols    = ["udp"]
      port         = 123
      destinations = data.xcsh_network_customer_edge_defaults.system_services.ntp_servers
    }
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/examples/)
