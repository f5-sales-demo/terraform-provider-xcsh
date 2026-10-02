---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_infrastructure."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1357, "body_sha256": "sha256:e7c2ee09e8ae6e33c87f581a615847af0b867e9f28834b69b237ef5e34919576", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_infrastructure:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f72b753447890fa3999ee0487a20af9414eac1431bcf20d1ef50048bd9025f06", "source_path": "examples/data-sources/xcsh_bot_infrastructure/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_infrastructure:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_infrastructure:examples", "path": "documentation/data-sources/bot_infrastructure/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2122303003330312-2112310320230020-0311223203032121-2200300011032333-2223202002211123-1212233100203201-2203011311112220-2303003133112113", "registry_path": "docs/guides/data-sources--bot_infrastructure--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_infrastructure/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_infrastructure.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
