---
page_title: "xcsh_bot_allowlist_policy"
subcategory: ""
description: "Reads Bot Allowlist Policy information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["bot allowlist policy"], "body_bytes": 1403, "body_sha256": "sha256:77ffee8dfe6b348c414fdf980a4266f01284d93384b2c19af89db8bedeb2d463", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_allowlist_policy:reference", "xcsh-docs:data-sources:bot_allowlist_policy:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_allowlist_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_allowlist_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_allowlist_policy/index.md", "product": "distributed-cloud", "provider_name": "bot_allowlist_policy", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-0022111012002123-0020202233333300-0020222322013201-0023201320232112-2223012222333122-3222112110213212-1033221302121302-2123032122301213", "registry_path": "docs/data-sources/bot_allowlist_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_allowlist_policy/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Reads Bot Allowlist Policy information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": [], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_allowlist_policy

Breadcrumbs:

- xcsh_bot_allowlist_policy

Reads Bot Allowlist Policy information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_allowlist_policy/examples/)
