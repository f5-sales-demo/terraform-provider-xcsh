---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bot_defense_app_infrastructure."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1455, "body_sha256": "sha256:4604cedd854db26d8fb62d1dec246496ee61218b299356817c4dfb6479d1a4cd", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_defense_app_infrastructure:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3d3368cb414ef6890130a5c64346a138d87252424badf8402f264c1532265f55", "source_path": "examples/resources/xcsh_bot_defense_app_infrastructure/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bot_defense_app_infrastructure:example:resource", "parent_id": "xcsh-docs:resources:bot_defense_app_infrastructure:examples", "path": "documentation/resources/bot_defense_app_infrastructure/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3110122032110102-2033130301023300-2113300333023101-0323203101322332-2031233123012213-3130231323023211-1031230112313130-0232233131122120", "registry_path": "docs/guides/resources--bot_defense_app_infrastructure--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_defense_app_infrastructure/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_bot_defense_app_infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/examples/)
- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/)
