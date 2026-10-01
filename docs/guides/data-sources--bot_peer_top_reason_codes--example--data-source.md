---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_peer_top_reason_codes."
xcsh_docs: {"aliases": [], "body_bytes": 1141, "body_sha256": "sha256:607dfa504e4722bbf86e180292ed5ff1d6304cd3bf14f06e990a06e9e48bdcea", "canonical_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:088c7908202451a1670b0688576af84d4f3a2cb203efd005dc0803bb8dc3d516", "source_path": "examples/data-sources/xcsh_bot_peer_top_reason_codes/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_peer_top_reason_codes:examples", "path": "docs/guides/data-sources--bot_peer_top_reason_codes--example--data-source.md", "provider_name": "bot_peer_top_reason_codes", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_peer_top_reason_codes/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_peer_top_reason_codes.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md)
- [Examples](data-sources--bot_peer_top_reason_codes--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_peer_top_reason_codes/data-source.tf`; digest `sha256:088c7908202451a1670b0688576af84d4f3a2cb203efd005dc0803bb8dc3d516`.

```terraform
# BotPeerTopReasonCodes DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_bot_peer_top_reason_codes" "example" {
  namespace = "example-value"
}

output "bot_peer_top_reason_codes_result" {
  value = data.xcsh_bot_peer_top_reason_codes.example
}
```

## Next pages

- [Examples](data-sources--bot_peer_top_reason_codes--examples.md)
- [xcsh_bot_peer_top_reason_codes](../data-sources/bot_peer_top_reason_codes.md)
