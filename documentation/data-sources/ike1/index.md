---
page_title: "xcsh_ike1"
subcategory: ""
description: "Reads Ike1 information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["ike1"], "body_bytes": 1231, "body_sha256": "sha256:ed6aaf990e804e182abdc402d7af18b1c8bdfc8cb5fd541e2a4448b3e40f5b44", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:ike1:reference", "xcsh-docs:data-sources:ike1:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike1:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/ike1/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313", "registry_path": "docs/data-sources/ike1.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike1/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Reads Ike1 information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["ike1CreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike1

Breadcrumbs:

- xcsh_ike1

Reads Ike1 information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike1 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike1 by name
data "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}

output "ike1_id" {
  value = data.xcsh_ike1.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/examples/)
