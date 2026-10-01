---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_suggest_values."
xcsh_docs: {"aliases": [], "body_bytes": 1066, "body_sha256": "sha256:65012efb777caf840428c88117d66dc0e3636abf08d7230296ab190de6d3db31", "canonical_id": "xcsh-docs:data-sources:bot_suggest_values:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_suggest_values:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c6c1b3639717b7c716bae9d4f7857fddc676935e02dc229f788549aecb0a97ae", "source_path": "examples/data-sources/xcsh_bot_suggest_values/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_suggest_values:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_suggest_values:examples", "path": "docs/guides/data-sources--bot_suggest_values--example--data-source.md", "provider_name": "bot_suggest_values", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_suggest_values/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_suggest_values.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md)
- [Examples](data-sources--bot_suggest_values--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_suggest_values/data-source.tf`; digest `sha256:c6c1b3639717b7c716bae9d4f7857fddc676935e02dc229f788549aecb0a97ae`.

```terraform
# BotSuggestValues DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_suggest_values" "example" {
  namespace = "example-value"
}

output "bot_suggest_values_result" {
  value = data.xcsh_bot_suggest_values.example
}
```

## Next pages

- [Examples](data-sources--bot_suggest_values--examples.md)
- [xcsh_bot_suggest_values](../data-sources/bot_suggest_values.md)
