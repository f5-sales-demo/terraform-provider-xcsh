---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1162, "body_sha256": "sha256:17be8f273c059d75de1cb44de030bc25a28cf5cb74913dda5936432f5e963cf1", "canonical_id": "xcsh-docs:data-sources:bot_endpoint_policy:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:742dea18fab3d10fddee899cf0137a64f4bc6d260ffbdc975152403793a768eb", "source_path": "examples/data-sources/xcsh_bot_endpoint_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_endpoint_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:examples", "path": "docs/guides/data-sources--bot_endpoint_policy--example--data-source.md", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_endpoint_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
- [Examples](data-sources--bot_endpoint_policy--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_endpoint_policy/data-source.tf`; digest `sha256:742dea18fab3d10fddee899cf0137a64f4bc6d260ffbdc975152403793a768eb`.

```terraform
# BotEndpointPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotEndpointPolicy by name
data "xcsh_bot_endpoint_policy" "example" {
  name      = "example-bot-endpoint-policy"
  namespace = "staging"
}

output "bot_endpoint_policy_id" {
  value = data.xcsh_bot_endpoint_policy.example.id
}
```

## Next pages

- [Examples](data-sources--bot_endpoint_policy--examples.md)
- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md)
