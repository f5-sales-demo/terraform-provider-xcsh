---
page_title: "xcsh_bot_infrastructure"
subcategory: ""
description: "xcsh_bot_infrastructure for xcsh_bot_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 1358, "body_sha256": "sha256:157631cef54cac6874b8406e19b87a7f29a1fb42524aaecdb38c339374589155", "canonical_id": "xcsh-docs:resources:bot_infrastructure:fundamentals", "child_ids": ["xcsh-docs:resources:bot_infrastructure:reference", "xcsh-docs:resources:bot_infrastructure:examples", "xcsh-docs:resources:bot_infrastructure:import", "xcsh-docs:resources:bot_infrastructure:timeouts"], "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_infrastructure:fundamentals", "parent_id": null, "path": "docs/resources/bot_infrastructure.md", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_bot_infrastructure for xcsh_bot_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_infrastructure

Breadcrumbs:

- xcsh_bot_infrastructure

Manages Bot Infrastructure in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotInfrastructure Resource Example
# Manages Bot Infrastructure in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BotInfrastructure configuration
resource "xcsh_bot_infrastructure" "example" {
  name      = "example-bot-infrastructure"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--bot_infrastructure--reference.md)
- [Examples](../guides/resources--bot_infrastructure--examples.md)
- [Import](../guides/resources--bot_infrastructure--import.md)
- [Timeouts](../guides/resources--bot_infrastructure--timeouts.md)
