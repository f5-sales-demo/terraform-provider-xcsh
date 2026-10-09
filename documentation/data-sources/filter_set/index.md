---
page_title: "xcsh_filter_set"
subcategory: ""
description: "Reads Filter Set information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["filter set"], "body_bytes": 1295, "body_sha256": "sha256:cba6ba4224320f8d7c000cddbe08141d90170250dbf4c261493beb39f5b09e9c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:filter_set:reference", "xcsh-docs:data-sources:filter_set:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:filter_set:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/filter_set/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3220320303320131-0131310102101111-3102330020220020-3222001011311133-0032232323232000-2210333331011111-1110011031200321-2111213033230001", "registry_path": "docs/data-sources/filter_set.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/filter_set/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Reads Filter Set information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["filter_setCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_filter_set

Breadcrumbs:

- xcsh_filter_set

Reads Filter Set information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FilterSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FilterSet by name
data "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"
}

output "filter_set_id" {
  value = data.xcsh_filter_set.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/examples/)
