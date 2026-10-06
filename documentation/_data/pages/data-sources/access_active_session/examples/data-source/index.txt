---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_access_active_session."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1080, "body_sha256": "sha256:212c07030a805997fea950b187aa95b83137f8dec5e70890183a332235e96d7c", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:access_active_session:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d3411f6908febc7c8a9bf8b482b2e026429a95e16a75b366bc2a33e4f32326f7", "source_path": "examples/data-sources/xcsh_access_active_session/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:access_active_session:example:data-source", "parent_id": "xcsh-docs:data-sources:access_active_session:examples", "path": "documentation/data-sources/access_active_session/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "access_active_session", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1122210201123331-0211211312033112-1200212103201123-0130011330021332-0220333011310002-3213212230011310-1131320222212210-1221321100012001", "registry_path": "docs/guides/data-sources--access_active_session--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_session/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_access_active_session.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_access_active_session](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_access_active_session/data-source.tf`; digest `sha256:d3411f6908febc7c8a9bf8b482b2e026429a95e16a75b366bc2a33e4f32326f7`.

```terraform
# AccessActiveSession DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_session" "example" {
  id        = "example-value"
  namespace = "example-value"
}

output "access_active_session_result" {
  value = data.xcsh_access_active_session.example
}
```
