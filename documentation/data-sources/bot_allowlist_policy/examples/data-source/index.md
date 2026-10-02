---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_allowlist_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1381, "body_sha256": "sha256:2625e2b78502007c37828fefd39fe822519317508dd49980b41cd0d59b952c74", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e79cb351210b30d9b4d2bf6ff1c7b6c44610a7fd5a55e86ba638364a3bd45c0b", "source_path": "examples/data-sources/xcsh_bot_allowlist_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_allowlist_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_allowlist_policy:examples", "path": "documentation/data-sources/bot_allowlist_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0220001130331330-3110113221021010-3131112022300313-2232121223321222-3121221110311100-1223101022213101-2021023130333132-1102320312332230", "registry_path": "docs/guides/data-sources--bot_allowlist_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_bot_allowlist_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/examples/)
- [xcsh_bot_allowlist_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/)
