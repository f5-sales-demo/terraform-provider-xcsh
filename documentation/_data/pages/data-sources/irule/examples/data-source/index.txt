---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_irule."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 983, "body_sha256": "sha256:ce72cdfb2d8eaf103090c5edaa6e8020ae3035431b021ec681de8da4f932a761", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:irule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:567750397ff4df80dac106ded5aa33f7fcb698c11cea8ada22ebaa5fc6727f22", "source_path": "examples/data-sources/xcsh_irule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:irule:example:data-source", "parent_id": "xcsh-docs:data-sources:irule:examples", "path": "documentation/data-sources/irule/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "irule", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2020310331322123-1231313213032103-0220130101121132-0211022211302012-0201022123222333-2021200303030110-1320322300330203-2303102002211031", "registry_path": "docs/guides/data-sources--irule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/irule/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_irule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["iruleCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_irule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/irule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/irule/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_irule/data-source.tf`; digest `sha256:567750397ff4df80dac106ded5aa33f7fcb698c11cea8ada22ebaa5fc6727f22`.

```terraform
# Irule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Irule by name
data "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"
}

output "irule_id" {
  value = data.xcsh_irule.example.id
}
```
