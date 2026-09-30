---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_shape_bot_defense_instance."
xcsh_docs: {"aliases": [], "body_bytes": 1358, "body_sha256": "sha256:33dfb0179b899d90b2e11f021d0608ac21611f4cec55a1527de6c01fe2489bbd", "child_ids": [], "collection_id": "xcsh-docs:data-sources:shape_bot_defense_instance:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:559cb2ebbea61576b9b4c66aad8123cbb2b2816c6fc4d4e4939f3b3e386387b6", "source_path": "examples/data-sources/xcsh_shape_bot_defense_instance/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:shape_bot_defense_instance:example:data-source", "parent_id": "xcsh-docs:data-sources:shape_bot_defense_instance:examples", "path": "documentation/data-sources/shape_bot_defense_instance/examples/data-source/index.md", "provider_name": "shape_bot_defense_instance", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/shape_bot_defense_instance/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_shape_bot_defense_instance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
