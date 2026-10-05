---
page_title: "xcsh_code_base_integration"
subcategory: ""
description: "Manages integration details in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["code base integration"], "body_bytes": 1577, "body_sha256": "sha256:e1a180931c9b1073ca26adac51e6be8438adfd697da7da8bff0a7ba3d125ee63", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:code_base_integration:reference", "xcsh-docs:resources:code_base_integration:examples", "xcsh-docs:resources:code_base_integration:import", "xcsh-docs:resources:code_base_integration:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/code_base_integration/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3210121110030031-0102101131101010-3032112220130303-0002122112230220-1300100012102310-1322031221321203-3201301222000010-1330202211303002", "registry_path": "docs/resources/code_base_integration.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages integration details in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/lifecycle/timeouts/)
