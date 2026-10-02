---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_suggest_values."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1272, "body_sha256": "sha256:c36e89540e0ff79e31ab93de790e9aeccc37b09ac4baac0d2d9bab0b0832ebd0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_suggest_values:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c6c1b3639717b7c716bae9d4f7857fddc676935e02dc229f788549aecb0a97ae", "source_path": "examples/data-sources/xcsh_bot_suggest_values/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_suggest_values:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_suggest_values:examples", "path": "documentation/data-sources/bot_suggest_values/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_suggest_values", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0313312323011010-1031111310122203-3012110312110231-3120311032102232-2001012313110231-1121323110102121-3010110120222122-2232101100310000", "registry_path": "docs/guides/data-sources--bot_suggest_values--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_suggest_values/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_suggest_values.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_suggest_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/examples/)
- [xcsh_bot_suggest_values](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_suggest_values/)
