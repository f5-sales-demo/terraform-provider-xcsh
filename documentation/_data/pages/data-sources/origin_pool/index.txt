---
page_title: "xcsh_origin_pool"
subcategory: "Load Balancing"
description: "Reads an Origin Pool that defines backend servers for load balancer targets."
xcsh_docs: {"aliases": ["backend servers", "origin pool", "origin servers", "upstream servers"], "body_bytes": 1444, "body_sha256": "sha256:a286288f570c48ae090fcfab53ae5b1f6eb23796529d0eaf408ecd935fd7a18d", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:origin_pool:reference", "xcsh-docs:data-sources:origin_pool:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/origin_pool/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203", "registry_path": "docs/data-sources/origin_pool.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:optional", "target_id": "xcsh-docs:resources:healthcheck:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads an Origin Pool that defines backend servers for load balancer targets.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_origin_pool

Breadcrumbs:

- xcsh_origin_pool

Reads an Origin Pool that defines backend servers for load balancer targets.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `healthcheck`.

- healthcheck: Monitor origin server health

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# OriginPool Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing OriginPool by name
data "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}

output "origin_pool_id" {
  value = data.xcsh_origin_pool.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/examples/)
