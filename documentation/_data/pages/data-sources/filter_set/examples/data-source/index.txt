---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_filter_set."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1031, "body_sha256": "sha256:2d07b00f6fc5ceead57b22e99c1149f0c9be9e0a1b0e0712fc4eb37b22b86a15", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:filter_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:03a6f567d57f7b6dda2ca579596922f0b9b282f98f909656ba616ef93c767b0d", "source_path": "examples/data-sources/xcsh_filter_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:filter_set:example:data-source", "parent_id": "xcsh-docs:data-sources:filter_set:examples", "path": "documentation/data-sources/filter_set/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2120310123232131-1010111201313002-2222112022000303-1312202013021330-2121202223020110-3313022212330131-1112230210131311-2011010312313322", "registry_path": "docs/guides/data-sources--filter_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/filter_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Data source for xcsh_filter_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["filter_setCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_filter_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_filter_set/data-source.tf`; digest `sha256:03a6f567d57f7b6dda2ca579596922f0b9b282f98f909656ba616ef93c767b0d`.

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
