---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_bot_defense_app_infrastructure."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1179, "body_sha256": "sha256:b28a52a6d41253ed05e6e9bf49bb396bdbea02901b9f668a56bb727c268088cb", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_defense_app_infrastructure:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3d3368cb414ef6890130a5c64346a138d87252424badf8402f264c1532265f55", "source_path": "examples/resources/xcsh_bot_defense_app_infrastructure/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:bot_defense_app_infrastructure:example:resource", "parent_id": "xcsh-docs:resources:bot_defense_app_infrastructure:examples", "path": "documentation/resources/bot_defense_app_infrastructure/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "resources", "registry_anchor": "canonical-3110122032110102-2033130301023300-2113300333023101-0323203101322332-2031233123012213-3130231323023211-1031230112313130-0232233131122120", "registry_path": "docs/guides/resources--bot_defense_app_infrastructure--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_defense_app_infrastructure/examples/resource/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Resource for xcsh_bot_defense_app_infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
