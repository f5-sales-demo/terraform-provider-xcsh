---
page_title: "xcsh_network_customer_edge_egress"
subcategory: ""
description: "xcsh_network_customer_edge_egress for xcsh_network_customer_edge_egress."
xcsh_docs: {"aliases": [], "body_bytes": 1752, "body_sha256": "sha256:6d5a32f55ae9b20c441b34a81fc44067dce5f75f18c508dc8fd798d490062d63", "child_ids": ["xcsh-docs:data-sources:network_customer_edge_egress:reference", "xcsh-docs:data-sources:network_customer_edge_egress:examples"], "collection_id": "xcsh-docs:data-sources:network_customer_edge_egress:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_customer_edge_egress:fundamentals", "parent_id": null, "path": "documentation/data-sources/network_customer_edge_egress/index.md", "provider_name": "network_customer_edge_egress", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_customer_edge_egress/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_customer_edge_egress for xcsh_network_customer_edge_egress.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_network_customer_edge_egress

Breadcrumbs:

- xcsh_network_customer_edge_egress

Secure Mesh v2 registration IPv4 addresses and egress domains. Legacy Customer Edge values are
intentionally excluded. Values are bundled from the pinned OpenAPI release; this data source
performs no network request. Ports and traffic direction are not encoded in the manifest.

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

data "xcsh_network_customer_edge_egress" "secure_mesh_v2" {}

# Use the domains with an FQDN-aware control. The legacy CE branch is not
# included in this data source.
output "secure_mesh_v2_https_egress" {
  value = {
    direction              = "egress"
    protocol               = "tcp"
    port                   = 443
    registration_addresses = data.xcsh_network_customer_edge_egress.secure_mesh_v2.registration_addresses
    domains                = data.xcsh_network_customer_edge_egress.secure_mesh_v2.domains
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/examples/)
