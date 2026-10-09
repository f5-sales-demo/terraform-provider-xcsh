---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_customer_edge_defaults."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1386, "body_sha256": "sha256:29afa1effd1c7d435bcdd7665c341b22a7d39b7862451dd2486e3f15b5b80f5e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_customer_edge_defaults:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:02dcb0c87e55fb1eb9417d76ac0888bfe3451a7f9f3e9a6cb8857b1ff2aff366", "source_path": "examples/data-sources/xcsh_network_customer_edge_defaults/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_customer_edge_defaults:example:data-source", "parent_id": "xcsh-docs:data-sources:network_customer_edge_defaults:examples", "path": "documentation/data-sources/network_customer_edge_defaults/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_customer_edge_defaults", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3001112300000133-3223020012320113-1332100113132213-1311103102113012-0213300302030302-0132301020322312-2000212301213233-0231010313312000", "registry_path": "docs/guides/data-sources--network_customer_edge_defaults--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_customer_edge_defaults/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_network_customer_edge_defaults.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_customer_edge_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_customer_edge_defaults/data-source.tf`; digest `sha256:02dcb0c87e55fb1eb9417d76ac0888bfe3451a7f9f3e9a6cb8857b1ff2aff366`.

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
