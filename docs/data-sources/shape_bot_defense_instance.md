---
page_title: "xcsh_shape_bot_defense_instance"
subcategory: ""
description: "xcsh_shape_bot_defense_instance for xcsh_shape_bot_defense_instance."
xcsh_docs: {"aliases": [], "body_bytes": 1451, "body_sha256": "sha256:df9362a61bae460a1a9ff127bb8d1c807ae8dd68e8cb65714e420b1c4e4d7eb3", "canonical_id": "xcsh-docs:data-sources:shape_bot_defense_instance:fundamentals", "child_ids": ["xcsh-docs:data-sources:shape_bot_defense_instance:reference", "xcsh-docs:data-sources:shape_bot_defense_instance:examples"], "collection_id": "xcsh-docs:data-sources:shape_bot_defense_instance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:shape_bot_defense_instance:fundamentals", "parent_id": null, "path": "docs/data-sources/shape_bot_defense_instance.md", "provider_name": "shape_bot_defense_instance", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/shape_bot_defense_instance/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_shape_bot_defense_instance for xcsh_shape_bot_defense_instance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_shape_bot_defense_instance

Breadcrumbs:

- xcsh_shape_bot_defense_instance

Manages a Shape Bot Defense Instance resource in F5 Distributed Cloud for get virtual host from a
given namespace. configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--shape_bot_defense_instance--reference.md)
- [Examples](../guides/data-sources--shape_bot_defense_instance--examples.md)
