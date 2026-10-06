---
page_title: "xcsh_bot_defense_app_infrastructure"
subcategory: ""
description: "Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["bot defense app infrastructure"], "body_bytes": 1742, "body_sha256": "sha256:a688fc8d0daca082690e80fd9278d970c69dfd799b426046e9f07c4fe8392323", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:bot_defense_app_infrastructure:reference", "xcsh-docs:resources:bot_defense_app_infrastructure:examples", "xcsh-docs:resources:bot_defense_app_infrastructure:import", "xcsh-docs:resources:bot_defense_app_infrastructure:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_defense_app_infrastructure:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/bot_defense_app_infrastructure/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1203003131331303-1311001012221111-3000012123322322-1122230311323033-0121031312111230-0211100011003100-1131122110310101-3113322201112121", "registry_path": "docs/resources/bot_defense_app_infrastructure.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_defense_app_infrastructure/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages Bot Defense App Infrastructure in a given namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/lifecycle/timeouts/)
