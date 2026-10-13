---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_type."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1011, "body_sha256": "sha256:d1c06d679963e4935416bff986d3478ca5536903a93febc8bdc7869ca7e9a297", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_type:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7219c79a69a1e857679b66c93800c9cdd60d2a0ca9b2ed275f17e432b1f04933", "source_path": "examples/data-sources/xcsh_app_type/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_type:example:data-source", "parent_id": "xcsh-docs:data-sources:app_type:examples", "path": "documentation/data-sources/app_type/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "app_type", "provider_schema_digest": "sha256:500e603af7c92c2c213b5cbdb574fdea5ac4e4825db9ce520b047c223b245685", "provider_type": "data-sources", "registry_anchor": "canonical-2120332221223021-0023301011011233-2130011301010102-0211320120213212-0002010313231021-2230011131203330-0102221100001033-1300033301010133", "registry_path": "docs/guides/data-sources--app_type--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_type/examples/data-source/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Data source for xcsh_app_type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": ["app_typeCreateRequest"], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_type/data-source.tf`; digest `sha256:7219c79a69a1e857679b66c93800c9cdd60d2a0ca9b2ed275f17e432b1f04933`.

```terraform
# AppType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppType by name
data "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}

output "app_type_id" {
  value = data.xcsh_app_type.example.id
}
```
