---
page_title: "xcsh_bot_infrastructure"
subcategory: ""
description: "Manages Bot Infrastructure in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["bot infrastructure"], "body_bytes": 1560, "body_sha256": "sha256:33e06dc193827971854885d0c567612fca898199ffb1baaaeb4a943503ba4039", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:bot_infrastructure:reference", "xcsh-docs:resources:bot_infrastructure:examples", "xcsh-docs:resources:bot_infrastructure:import", "xcsh-docs:resources:bot_infrastructure:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_infrastructure:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/bot_infrastructure/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2222321211302222-3000301022222311-2112020111020303-0211321333200200-0330000200113123-1222031102030300-3031133300033023-1330233231022033", "registry_path": "docs/resources/bot_infrastructure.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_infrastructure/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Manages Bot Infrastructure in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_infrastructure/lifecycle/timeouts/)
