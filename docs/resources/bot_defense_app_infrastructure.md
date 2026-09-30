---
page_title: "xcsh_bot_defense_app_infrastructure"
subcategory: ""
description: "xcsh_bot_defense_app_infrastructure for xcsh_bot_defense_app_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 1441, "body_sha256": "sha256:23cc8144157208de0df20aa61b8d38537c9f73377d8a7a851ead570ed7867d7a", "canonical_id": "xcsh-docs:resources:bot_defense_app_infrastructure:fundamentals", "child_ids": ["xcsh-docs:resources:bot_defense_app_infrastructure:reference", "xcsh-docs:resources:bot_defense_app_infrastructure:examples", "xcsh-docs:resources:bot_defense_app_infrastructure:import", "xcsh-docs:resources:bot_defense_app_infrastructure:timeouts"], "collection_id": "xcsh-docs:resources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_defense_app_infrastructure:fundamentals", "parent_id": null, "path": "docs/resources/bot_defense_app_infrastructure.md", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_defense_app_infrastructure/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bot_defense_app_infrastructure for xcsh_bot_defense_app_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

- [Property reference](../guides/resources--bot_defense_app_infrastructure--reference.md)
- [Examples](../guides/resources--bot_defense_app_infrastructure--examples.md)
- [Import](../guides/resources--bot_defense_app_infrastructure--import.md)
- [Timeouts](../guides/resources--bot_defense_app_infrastructure--timeouts.md)
