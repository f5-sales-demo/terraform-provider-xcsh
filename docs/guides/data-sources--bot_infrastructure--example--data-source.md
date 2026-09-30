---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 1052, "body_sha256": "sha256:eb64c846bfbe87b0d9ae4c07ce6ebf3cff38a186ec759632bdff75648cc3d47e", "canonical_id": "xcsh-docs:data-sources:bot_infrastructure:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_infrastructure:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f72b753447890fa3999ee0487a20af9414eac1431bcf20d1ef50048bd9025f06", "source_path": "examples/data-sources/xcsh_bot_infrastructure/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_infrastructure:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_infrastructure:examples", "path": "docs/guides/data-sources--bot_infrastructure--example--data-source.md", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_infrastructure/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md)
- [Examples](data-sources--bot_infrastructure--examples.md)
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

- [Examples](data-sources--bot_infrastructure--examples.md)
- [xcsh_bot_infrastructure](../data-sources/bot_infrastructure.md)
