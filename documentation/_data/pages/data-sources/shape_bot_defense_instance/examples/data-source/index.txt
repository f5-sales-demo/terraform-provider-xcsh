---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_shape_bot_defense_instance."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1187, "body_sha256": "sha256:27d11a4b251de5201f85022fa5bea438ad1eb6f6b0afde027758073c3f49c4d8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:shape_bot_defense_instance:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:559cb2ebbea61576b9b4c66aad8123cbb2b2816c6fc4d4e4939f3b3e386387b6", "source_path": "examples/data-sources/xcsh_shape_bot_defense_instance/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:shape_bot_defense_instance:example:data-source", "parent_id": "xcsh-docs:data-sources:shape_bot_defense_instance:examples", "path": "documentation/data-sources/shape_bot_defense_instance/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "shape_bot_defense_instance", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3322212013131311-1111032230113003-1110202122221132-1033320110201221-1020230113231001-0232332121302032-2023013211302101-3233112231001022", "registry_path": "docs/guides/data-sources--shape_bot_defense_instance--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/shape_bot_defense_instance/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_shape_bot_defense_instance.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
