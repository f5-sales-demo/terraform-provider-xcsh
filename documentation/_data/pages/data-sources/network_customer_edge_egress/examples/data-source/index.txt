---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_customer_edge_egress."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1669, "body_sha256": "sha256:ea7d14265fe20dc5da651ece2d6746c37650a086718acbd5452245df423e2701", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_customer_edge_egress:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1a958054768616b5beeebffa91db22fe699af5f80112367c4fcd9c1d753d7368", "source_path": "examples/data-sources/xcsh_network_customer_edge_egress/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_customer_edge_egress:example:data-source", "parent_id": "xcsh-docs:data-sources:network_customer_edge_egress:examples", "path": "documentation/data-sources/network_customer_edge_egress/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_customer_edge_egress", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3010201100132110-0030003012023201-2323030012013011-0022022200121230-2222110121212311-0121121102300132-2310013211133020-2201000201132211", "registry_path": "docs/guides/data-sources--network_customer_edge_egress--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_customer_edge_egress/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_network_customer_edge_egress.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_customer_edge_egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_customer_edge_egress/data-source.tf`; digest `sha256:1a958054768616b5beeebffa91db22fe699af5f80112367c4fcd9c1d753d7368`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/examples/)
- [xcsh_network_customer_edge_egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_egress/)
