---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_allowlist_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1129, "body_sha256": "sha256:b83af0f1c6fc6db42f343269d58287b1724974b0ee2084c3ec0748be3ad4bf44", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e79cb351210b30d9b4d2bf6ff1c7b6c44610a7fd5a55e86ba638364a3bd45c0b", "source_path": "examples/data-sources/xcsh_bot_allowlist_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_allowlist_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:examples", "path": "documentation/data-sources/bot_allowlist_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-0220001130331330-3110113221021010-3131112022300313-2232121223321222-3121221110311100-1223101022213101-2021023130333132-1102320312332230", "registry_path": "docs/guides/data-sources--bot_allowlist_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_bot_allowlist_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_allowlist_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_allowlist_policy/data-source.tf`; digest `sha256:e79cb351210b30d9b4d2bf6ff1c7b6c44610a7fd5a55e86ba638364a3bd45c0b`.

```terraform
# BotAllowlistPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotAllowlistPolicy by name
data "xcsh_bot_allowlist_policy" "example" {
  name      = "example-bot-allowlist-policy"
  namespace = "staging"
}

output "bot_allowlist_policy_id" {
  value = data.xcsh_bot_allowlist_policy.example.id
}
```
