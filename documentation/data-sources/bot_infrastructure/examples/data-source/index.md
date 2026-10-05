---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_infrastructure."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1357, "body_sha256": "sha256:e7c2ee09e8ae6e33c87f581a615847af0b867e9f28834b69b237ef5e34919576", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_infrastructure:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f72b753447890fa3999ee0487a20af9414eac1431bcf20d1ef50048bd9025f06", "source_path": "examples/data-sources/xcsh_bot_infrastructure/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_infrastructure:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_infrastructure:examples", "path": "documentation/data-sources/bot_infrastructure/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2122303003330312-2112310320230020-0311223203032121-2200300011032333-2223202002211123-1212233100203201-2203011311112220-2303003133112113", "registry_path": "docs/guides/data-sources--bot_infrastructure--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_infrastructure/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_bot_infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_infrastructure/data-source.tf`; digest `sha256:f72b753447890fa3999ee0487a20af9414eac1431bcf20d1ef50048bd9025f06`.

```terraform
# BotInfrastructure Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotInfrastructure by name
data "xcsh_bot_infrastructure" "example" {
  name      = "example-bot-infrastructure"
  namespace = "staging"
}

output "bot_infrastructure_id" {
  value = data.xcsh_bot_infrastructure.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/examples/)
- [xcsh_bot_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/)
