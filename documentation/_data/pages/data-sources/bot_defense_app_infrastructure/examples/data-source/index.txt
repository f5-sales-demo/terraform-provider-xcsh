---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_defense_app_infrastructure."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1227, "body_sha256": "sha256:8dc8559df3fe9d4015ede312edc368eb098e2a56b5595f0af2707b7c650c0285", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9d1dcbef65f7cbdf8cc89c39b723d01c2f54dd36e7a2d02f47f90c88068c73e0", "source_path": "examples/data-sources/xcsh_bot_defense_app_infrastructure/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:examples", "path": "documentation/data-sources/bot_defense_app_infrastructure/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3310222001301223-3222221111211002-1103321222223201-0132033202120020-3012311220222302-0210012033102201-0233023132131000-2102300310203212", "registry_path": "docs/guides/data-sources--bot_defense_app_infrastructure--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_defense_app_infrastructure/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_bot_defense_app_infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_defense_app_infrastructure/data-source.tf`; digest `sha256:9d1dcbef65f7cbdf8cc89c39b723d01c2f54dd36e7a2d02f47f90c88068c73e0`.

```terraform
# BotDefenseAppInfrastructure Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotDefenseAppInfrastructure by name
data "xcsh_bot_defense_app_infrastructure" "example" {
  name      = "example-bot-defense-app-infrastructure"
  namespace = "staging"
}

output "bot_defense_app_infrastructure_id" {
  value = data.xcsh_bot_defense_app_infrastructure.example.id
}
```
