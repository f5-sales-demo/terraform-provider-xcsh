---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_customer_edge_defaults."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1668, "body_sha256": "sha256:50a0c26c55b1ae39e4bff0dcdd8c93e728d29158e49b2e92d708cfe6a8bfa95b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_customer_edge_defaults:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:02dcb0c87e55fb1eb9417d76ac0888bfe3451a7f9f3e9a6cb8857b1ff2aff366", "source_path": "examples/data-sources/xcsh_network_customer_edge_defaults/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_customer_edge_defaults:example:data-source", "parent_id": "xcsh-docs:data-sources:network_customer_edge_defaults:examples", "path": "documentation/data-sources/network_customer_edge_defaults/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_customer_edge_defaults", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3001112300000133-3223020012320113-1332100113132213-1311103102113012-0213300302030302-0132301020322312-2000212301213233-0231010313312000", "registry_path": "docs/guides/data-sources--network_customer_edge_defaults--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_customer_edge_defaults/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_network_customer_edge_defaults.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/examples/)
- [xcsh_network_customer_edge_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_customer_edge_defaults/)
