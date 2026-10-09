---
page_title: "xcsh_code_base_integration"
subcategory: ""
description: "Manages integration details in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["code base integration"], "body_bytes": 1590, "body_sha256": "sha256:78b801ab2ec98cb5bf145bd4219bd078072337ef83e3dd8be3c2ced710a6a2bf", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:code_base_integration:reference", "xcsh-docs:resources:code_base_integration:examples", "xcsh-docs:resources:code_base_integration:import", "xcsh-docs:resources:code_base_integration:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/code_base_integration/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002", "registry_path": "docs/resources/code_base_integration.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Manages integration details in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_code_base_integration

Breadcrumbs:

- xcsh_code_base_integration

Manages integration details in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CodeBaseIntegration Resource Example
# Manages integration details in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CodeBaseIntegration configuration
resource "xcsh_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/lifecycle/timeouts/)
