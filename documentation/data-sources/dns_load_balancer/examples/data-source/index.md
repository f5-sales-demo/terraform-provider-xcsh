---
page_title: "Data source"
subcategory: "DNS"
description: "Data source for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1098, "body_sha256": "sha256:18746a853e94627e4d71c702549d277006f790c8abddf47e685128070882950c", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:94e6c40d8f676ae588608f07a90134d4cad49dd9a1bfd873d996f0534dfb0e30", "source_path": "examples/data-sources/xcsh_dns_load_balancer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_load_balancer:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_load_balancer:examples", "path": "documentation/data-sources/dns_load_balancer/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2103302200311321-3132331213200101-0001232221121130-1203301012232001-3310022231002033-3112211212300321-3022120102202100-3302003033331103", "registry_path": "docs/guides/data-sources--dns_load_balancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/examples/data-source/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Data source for xcsh_dns_load_balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_load_balancer/data-source.tf`; digest `sha256:94e6c40d8f676ae588608f07a90134d4cad49dd9a1bfd873d996f0534dfb0e30`.

```terraform
# DNSLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLoadBalancer by name
data "xcsh_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"
}

output "dns_load_balancer_id" {
  value = data.xcsh_dns_load_balancer.example.id
}
```
