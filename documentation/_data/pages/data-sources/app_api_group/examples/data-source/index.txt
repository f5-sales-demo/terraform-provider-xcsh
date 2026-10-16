---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_api_group."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1059, "body_sha256": "sha256:4bce887993e8534c75074dad5285f52551b462ce73f71f6437ce59dfb0663834", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_api_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ea62726df10ead3d5992d45e3c012bb70708bc94e4271b50bd5cb8013ce78312", "source_path": "examples/data-sources/xcsh_app_api_group/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_api_group:example:data-source", "parent_id": "xcsh-docs:data-sources:app_api_group:examples", "path": "documentation/data-sources/app_api_group/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3313023333231021-0222020033001332-3103101120320231-3021031120132230-1111122322111110-3301023103232020-0230000020312220-0031200222103313", "registry_path": "docs/guides/data-sources--app_api_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_api_group/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Data source for xcsh_app_api_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_api_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_api_group/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_api_group/data-source.tf`; digest `sha256:ea62726df10ead3d5992d45e3c012bb70708bc94e4271b50bd5cb8013ce78312`.

```terraform
# AppAPIGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppAPIGroup by name
data "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}

output "app_api_group_id" {
  value = data.xcsh_app_api_group.example.id
}
```
