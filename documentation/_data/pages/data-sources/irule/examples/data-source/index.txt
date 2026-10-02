---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_irule."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1190, "body_sha256": "sha256:074eb1c480b1086d8bf23678b2f150ce83061e275e74637928327eb3ab7b3a6d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:irule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:567750397ff4df80dac106ded5aa33f7fcb698c11cea8ada22ebaa5fc6727f22", "source_path": "examples/data-sources/xcsh_irule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:irule:example:data-source", "parent_id": "xcsh-docs:data-sources:irule:examples", "path": "documentation/data-sources/irule/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "irule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2020310331322123-1231313213032103-0220130101121132-0211022211302012-0201022123222333-2021200303030110-1320322300330203-2303102002211031", "registry_path": "docs/guides/data-sources--irule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/irule/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_irule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["iruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/irule/examples/)
- [xcsh_irule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/irule/)
