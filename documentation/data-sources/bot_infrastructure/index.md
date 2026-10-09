---
page_title: "xcsh_bot_infrastructure"
subcategory: ""
description: "Reads Bot Infrastructure information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["bot infrastructure"], "body_bytes": 1383, "body_sha256": "sha256:a632e5430f631d42349893204813d7b20d00433a07ecc0218ea633017073979f", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:bot_infrastructure:reference", "xcsh-docs:data-sources:bot_infrastructure:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_infrastructure:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_infrastructure/index.md", "product": "distributed-cloud", "provider_name": "bot_infrastructure", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210", "registry_path": "docs/data-sources/bot_infrastructure.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_infrastructure/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Reads Bot Infrastructure information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["bot_infrastructureCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_infrastructure

Breadcrumbs:

- xcsh_bot_infrastructure

Reads Bot Infrastructure information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_infrastructure/examples/)
