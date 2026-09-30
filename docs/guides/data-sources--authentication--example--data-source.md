---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 1002, "body_sha256": "sha256:3e5b6e4a6eec53b37f3341a4e57fecfdd58eac91c4e8b2f81c867e0b90856041", "canonical_id": "xcsh-docs:data-sources:authentication:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fd8b227ee8d771e5d805c51da258f79cea8ff5031e710c318c14f6217d6efd78", "source_path": "examples/data-sources/xcsh_authentication/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:authentication:example:data-source", "parent_id": "xcsh-docs:data-sources:authentication:examples", "path": "docs/guides/data-sources--authentication--example--data-source.md", "provider_name": "authentication", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_authentication](../data-sources/authentication.md)
- [Examples](data-sources--authentication--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_authentication/data-source.tf`; digest `sha256:fd8b227ee8d771e5d805c51da258f79cea8ff5031e710c318c14f6217d6efd78`.

```terraform
# Authentication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Authentication by name
data "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}

output "authentication_id" {
  value = data.xcsh_authentication.example.id
}
```

## Next pages

- [Examples](data-sources--authentication--examples.md)
- [xcsh_authentication](../data-sources/authentication.md)
