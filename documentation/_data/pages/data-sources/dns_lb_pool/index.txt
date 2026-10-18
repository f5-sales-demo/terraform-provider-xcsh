---
page_title: "xcsh_dns_lb_pool"
subcategory: ""
description: "Reads an existing DNS Load Balancer pool in its namespace."
xcsh_docs: {"aliases": ["dns lb pool"], "body_bytes": 1292, "body_sha256": "sha256:e27b4f7d005aa2f63ab2c8b56f63ee38d84bf3d5a0b7710b7ac0d58586e9de4a", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_lb_pool:reference", "xcsh-docs:data-sources:dns_lb_pool:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_lb_pool:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/dns_lb_pool/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010", "registry_path": "docs/data-sources/dns_lb_pool.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_pool/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Reads an existing DNS Load Balancer pool in its namespace.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_lb_pool

Breadcrumbs:

- xcsh_dns_lb_pool

Reads an existing DNS Load Balancer pool in its namespace.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSLBPool Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLBPool by name
data "xcsh_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}

output "dns_lb_pool_id" {
  value = data.xcsh_dns_lb_pool.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_lb_pool/examples/)
