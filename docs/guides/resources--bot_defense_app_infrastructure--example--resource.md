---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bot_defense_app_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 1249, "body_sha256": "sha256:fdb3b1eb0188bb54e66a6d7b1646f5c8874db54b7d1758108361ac20627d5556", "canonical_id": "xcsh-docs:resources:bot_defense_app_infrastructure:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:bot_defense_app_infrastructure:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3d3368cb414ef6890130a5c64346a138d87252424badf8402f264c1532265f55", "source_path": "examples/resources/xcsh_bot_defense_app_infrastructure/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bot_defense_app_infrastructure:example:resource", "parent_id": "xcsh-docs:resources:bot_defense_app_infrastructure:examples", "path": "docs/guides/resources--bot_defense_app_infrastructure--example--resource.md", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_defense_app_infrastructure/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_bot_defense_app_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md)
- [Examples](resources--bot_defense_app_infrastructure--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bot_defense_app_infrastructure/resource.tf`; digest `sha256:3d3368cb414ef6890130a5c64346a138d87252424badf8402f264c1532265f55`.

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

## Next pages

- [Examples](resources--bot_defense_app_infrastructure--examples.md)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md)
