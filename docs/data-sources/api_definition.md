---
page_title: "xcsh_api_definition"
subcategory: "API Management"
description: "xcsh_api_definition for xcsh_api_definition."
xcsh_docs: {"aliases": [], "body_bytes": 1251, "body_sha256": "sha256:a9d3aaf0a79b1037733237367fc3171f66e83689c363286b0bc95075dbb8719b", "canonical_id": "xcsh-docs:data-sources:api_definition:fundamentals", "child_ids": ["xcsh-docs:data-sources:api_definition:reference", "xcsh-docs:data-sources:api_definition:examples"], "collection_id": "xcsh-docs:data-sources:api_definition:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_definition:fundamentals", "parent_id": null, "path": "docs/data-sources/api_definition.md", "provider_name": "api_definition", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_definition/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_api_definition for xcsh_api_definition.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_definitionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
# APIDefinition Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDefinition by name
data "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}

output "api_definition_id" {
  value = data.xcsh_api_definition.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--api_definition--reference.md)
- [Examples](../guides/data-sources--api_definition--examples.md)
