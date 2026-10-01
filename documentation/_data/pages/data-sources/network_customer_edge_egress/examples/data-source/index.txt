---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_customer_edge_egress."
xcsh_docs: {"aliases": [], "body_bytes": 1669, "body_sha256": "sha256:ea7d14265fe20dc5da651ece2d6746c37650a086718acbd5452245df423e2701", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_customer_edge_egress:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1a958054768616b5beeebffa91db22fe699af5f80112367c4fcd9c1d753d7368", "source_path": "examples/data-sources/xcsh_network_customer_edge_egress/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_customer_edge_egress:example:data-source", "parent_id": "xcsh-docs:data-sources:network_customer_edge_egress:examples", "path": "documentation/data-sources/network_customer_edge_egress/examples/data-source/index.md", "provider_name": "network_customer_edge_egress", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_customer_edge_egress/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_customer_edge_egress.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
