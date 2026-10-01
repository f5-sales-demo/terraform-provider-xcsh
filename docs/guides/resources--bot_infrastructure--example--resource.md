---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bot_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 1088, "body_sha256": "sha256:6207aa57936f6bb610d4f79fdc8315901e07f8aafd0118a3968dee2cb1aee96a", "canonical_id": "xcsh-docs:resources:bot_infrastructure:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8bbb63f8f4c5acb98b0adb6309584a0770310129d290ad3dfbcb6667452aa9e4", "source_path": "examples/resources/xcsh_bot_infrastructure/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bot_infrastructure:example:resource", "parent_id": "xcsh-docs:resources:bot_infrastructure:examples", "path": "docs/guides/resources--bot_infrastructure--example--resource.md", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_bot_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md)
- [Examples](resources--bot_infrastructure--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_bot_infrastructure/resource.tf`; digest `sha256:8bbb63f8f4c5acb98b0adb6309584a0770310129d290ad3dfbcb6667452aa9e4`.

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

## Next pages

- [Examples](resources--bot_infrastructure--examples.md)
- [xcsh_bot_infrastructure](../resources/bot_infrastructure.md)
