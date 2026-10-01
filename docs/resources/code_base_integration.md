---
page_title: "xcsh_code_base_integration"
subcategory: ""
description: "xcsh_code_base_integration for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 1388, "body_sha256": "sha256:56824819eef96709cb76189ef9a323ba6de40b8c4f4d517e880612d25fddce86", "canonical_id": "xcsh-docs:resources:code_base_integration:fundamentals", "child_ids": ["xcsh-docs:resources:code_base_integration:reference", "xcsh-docs:resources:code_base_integration:examples", "xcsh-docs:resources:code_base_integration:import", "xcsh-docs:resources:code_base_integration:timeouts"], "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "id": "xcsh-docs:resources:code_base_integration:fundamentals", "parent_id": null, "path": "docs/resources/code_base_integration.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_code_base_integration for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/resources--code_base_integration--reference.md)
- [Examples](../guides/resources--code_base_integration--examples.md)
- [Import](../guides/resources--code_base_integration--import.md)
- [Timeouts](../guides/resources--code_base_integration--timeouts.md)
