---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_shape_bot_defense_instance."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1457, "body_sha256": "sha256:0446de403af20f1fc10394ee973fa5eb43424758275ffd04d8b0f9f1d4e9cd65", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:shape_bot_defense_instance:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:559cb2ebbea61576b9b4c66aad8123cbb2b2816c6fc4d4e4939f3b3e386387b6", "source_path": "examples/data-sources/xcsh_shape_bot_defense_instance/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:shape_bot_defense_instance:example:data-source", "parent_id": "xcsh-docs:data-sources:shape_bot_defense_instance:examples", "path": "documentation/data-sources/shape_bot_defense_instance/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "shape_bot_defense_instance", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3322212013131311-1111032230113003-1110202122221132-1033320110201221-1020230113231001-0232332121302032-2023013211302101-3233112231001022", "registry_path": "docs/guides/data-sources--shape_bot_defense_instance--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/shape_bot_defense_instance/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_shape_bot_defense_instance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_shape_bot_defense_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/shape_bot_defense_instance/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/shape_bot_defense_instance/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_shape_bot_defense_instance/data-source.tf`; digest `sha256:559cb2ebbea61576b9b4c66aad8123cbb2b2816c6fc4d4e4939f3b3e386387b6`.

```terraform
# ShapeBotDefenseInstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ShapeBotDefenseInstance by name
data "xcsh_shape_bot_defense_instance" "example" {
  name      = "example-shape-bot-defense-instance"
  namespace = "staging"
}

output "shape_bot_defense_instance_id" {
  value = data.xcsh_shape_bot_defense_instance.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/shape_bot_defense_instance/examples/)
- [xcsh_shape_bot_defense_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/shape_bot_defense_instance/)
