---
page_title: "xcsh_bot_suggest_values"
subcategory: ""
description: "xcsh_bot_suggest_values for xcsh_bot_suggest_values."
xcsh_docs: {"aliases": [], "body_bytes": 1058, "body_sha256": "sha256:e6ba199701546eceacedc355271b852450488c7b36802a22669806de3f8ecb6c", "canonical_id": "xcsh-docs:data-sources:bot_suggest_values:fundamentals", "child_ids": ["xcsh-docs:data-sources:bot_suggest_values:reference", "xcsh-docs:data-sources:bot_suggest_values:examples"], "collection_id": "xcsh-docs:data-sources:bot_suggest_values:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_suggest_values:fundamentals", "parent_id": null, "path": "docs/data-sources/bot_suggest_values.md", "provider_name": "bot_suggest_values", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_suggest_values/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bot_suggest_values for xcsh_bot_suggest_values.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_bot_suggest_values

Breadcrumbs:

- xcsh_bot_suggest_values

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--bot_suggest_values--reference.md)
- [Examples](../guides/data-sources--bot_suggest_values--examples.md)
