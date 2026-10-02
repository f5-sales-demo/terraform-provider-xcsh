---
page_title: "xcsh_bot_defense_app_infrastructure"
subcategory: ""
description: "Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["bot defense app infrastructure"], "body_bytes": 1729, "body_sha256": "sha256:0c0077e930c029f3bdcf1d1ea616836b14bc72c35d96cc15bc107319fd4d9501", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:bot_defense_app_infrastructure:reference", "xcsh-docs:resources:bot_defense_app_infrastructure:examples", "xcsh-docs:resources:bot_defense_app_infrastructure:import", "xcsh-docs:resources:bot_defense_app_infrastructure:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_defense_app_infrastructure:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/bot_defense_app_infrastructure/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121", "registry_path": "docs/resources/bot_defense_app_infrastructure.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_defense_app_infrastructure/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_defense_app_infrastructure

Breadcrumbs:

- xcsh_bot_defense_app_infrastructure

Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotDefenseAppInfrastructure Resource Example
# Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BotDefenseAppInfrastructure configuration
resource "xcsh_bot_defense_app_infrastructure" "example" {
  name      = "example-bot-defense-app-infrastructure"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/lifecycle/timeouts/)
