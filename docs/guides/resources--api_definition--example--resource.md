---
page_title: "Resource"
subcategory: "API Management"
description: "Resource for xcsh_api_definition."
xcsh_docs: {"aliases": [], "body_bytes": 1040, "body_sha256": "sha256:2d4efc98fdbec504c50a23a6bea8029827e8f188ab7a9e6f59f36caa167adce7", "canonical_id": "xcsh-docs:resources:api_definition:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:api_definition:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8eaaa98845c6177dba9c28e17da78fe9cbe8a7d10b85b7c4e55e7857d3d8ba64", "source_path": "examples/resources/xcsh_api_definition/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:api_definition:example:resource", "parent_id": "xcsh-docs:resources:api_definition:examples", "path": "docs/guides/resources--api_definition--example--resource.md", "provider_name": "api_definition", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_definition/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_api_definition.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_definitionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_api_definition](../resources/api_definition.md)
- [Examples](resources--api_definition--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_api_definition/resource.tf`; digest `sha256:8eaaa98845c6177dba9c28e17da78fe9cbe8a7d10b85b7c4e55e7857d3d8ba64`.

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

## Next pages

- [Examples](resources--api_definition--examples.md)
- [xcsh_api_definition](../resources/api_definition.md)
