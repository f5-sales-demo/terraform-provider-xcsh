---
page_title: "xcsh_bot_infrastructure"
subcategory: ""
description: "Manages Bot Infrastructure in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["bot infrastructure"], "body_bytes": 1358, "body_sha256": "sha256:a183a449c1ed6a46a0d4297e1812837749a913559baef3ce4a68946830b068af", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:bot_infrastructure:reference", "xcsh-docs:data-sources:bot_infrastructure:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_infrastructure:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_infrastructure/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210", "registry_path": "docs/data-sources/bot_infrastructure.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_infrastructure/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages Bot Infrastructure in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/examples/)
