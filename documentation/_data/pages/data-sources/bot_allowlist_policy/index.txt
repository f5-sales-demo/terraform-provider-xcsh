---
page_title: "xcsh_bot_allowlist_policy"
subcategory: ""
description: "Manages a Bot Allowlist Policy resource in F5 Distributed Cloud for get bot allowlist policy. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["bot allowlist policy"], "body_bytes": 1457, "body_sha256": "sha256:374025568fe9ea94ea0c5fba718560020c1503588d504a2b123a7e6fdaaeb19c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_allowlist_policy:reference", "xcsh-docs:data-sources:bot_allowlist_policy:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_allowlist_policy/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0022111012002123-0020202233333300-0020222322013201-0023201320232112-2223012222333122-3222112110213212-1033221302121302-2123032122301213", "registry_path": "docs/data-sources/bot_allowlist_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Bot Allowlist Policy resource in F5 Distributed Cloud for get bot allowlist policy. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_allowlist_policy

Breadcrumbs:

- xcsh_bot_allowlist_policy

Manages a Bot Allowlist Policy resource in F5 Distributed Cloud for get bot allowlist policy.
configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/examples/)
