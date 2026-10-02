---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_defense_app_infrastructure."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1509, "body_sha256": "sha256:52cebd344e8b072bd2b299de52d7db8db81f3efd68d61e763f7612b2576b11ae", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9d1dcbef65f7cbdf8cc89c39b723d01c2f54dd36e7a2d02f47f90c88068c73e0", "source_path": "examples/data-sources/xcsh_bot_defense_app_infrastructure/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:examples", "path": "documentation/data-sources/bot_defense_app_infrastructure/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3310222001301223-3222221111211002-1103321222223201-0132033202120020-3012311220222302-0210012033102201-0233023132131000-2102300310203212", "registry_path": "docs/guides/data-sources--bot_defense_app_infrastructure--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_defense_app_infrastructure/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_bot_defense_app_infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/examples/)
- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/)
