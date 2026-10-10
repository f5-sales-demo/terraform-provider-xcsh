---
page_title: "xcsh_network_customer_edge_egress"
subcategory: ""
description: "Secure Mesh v2 registration IPv4 addresses and egress domains. Legacy Customer Edge values are intentionally excluded. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest."
xcsh_docs: {"aliases": ["network customer edge egress"], "body_bytes": 1864, "body_sha256": "sha256:d2efde6d984b8648a0fb852e2c2beeea5a0ce2ce7129dcba1d2c3b7ba7d2a129", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_customer_edge_egress:reference", "xcsh-docs:data-sources:network_customer_edge_egress:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_customer_edge_egress:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_customer_edge_egress:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_customer_edge_egress/index.md", "product": "distributed-cloud", "provider_name": "network_customer_edge_egress", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3220221103330020-3110221000233302-0211101212002003-1012001301133302-1131222331211321-2333221222011322-0330222303230333-2322220210010320", "registry_path": "docs/data-sources/network_customer_edge_egress.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_customer_edge_egress/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Secure Mesh v2 registration IPv4 addresses and egress domains. Legacy Customer Edge values are intentionally excluded. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/examples/)
