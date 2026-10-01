---
page_title: "xcsh_api_definition"
subcategory: "API Management"
description: "xcsh_api_definition for xcsh_api_definition."
xcsh_docs: {"aliases": [], "body_bytes": 1620, "body_sha256": "sha256:61bb8092ffe0f064a4ef48c3f791f7a66599ae5bbd414533401e793fde867e73", "child_ids": ["xcsh-docs:resources:api_definition:reference", "xcsh-docs:resources:api_definition:examples", "xcsh-docs:resources:api_definition:import", "xcsh-docs:resources:api_definition:timeouts"], "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_definition:fundamentals", "parent_id": null, "path": "documentation/resources/api_definition/index.md", "provider_name": "api_definition", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_api_definition for xcsh_api_definition.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_definitionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_api_definition

Breadcrumbs:

- xcsh_api_definition

Manages API Definition in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Optional integrations: `api_endpoint`.

- api_endpoint: Endpoints defined by this API

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APIDefinition Resource Example
# Manages API Definition in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDefinition configuration
resource "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/api_definition/lifecycle/timeouts/)
