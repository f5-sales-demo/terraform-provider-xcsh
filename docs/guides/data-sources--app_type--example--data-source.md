---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_type."
xcsh_docs: {"aliases": [], "body_bytes": 1021, "body_sha256": "sha256:b0c557ded165a80205bb5698a31b54c28672d3ed40fdfc9b14dfce99ff640644", "canonical_id": "xcsh-docs:data-sources:app_type:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:app_type:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7219c79a69a1e857679b66c93800c9cdd60d2a0ca9b2ed275f17e432b1f04933", "source_path": "examples/data-sources/xcsh_app_type/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_type:example:data-source", "parent_id": "xcsh-docs:data-sources:app_type:examples", "path": "docs/guides/data-sources--app_type--example--data-source.md", "provider_name": "app_type", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_type/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_app_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_type](../data-sources/app_type.md)
- [Examples](data-sources--app_type--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_type/data-source.tf`; digest `sha256:7219c79a69a1e857679b66c93800c9cdd60d2a0ca9b2ed275f17e432b1f04933`.

```terraform
# AppType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppType by name
data "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}

output "app_type_id" {
  value = data.xcsh_app_type.example.id
}
```

## Next pages

- [Examples](data-sources--app_type--examples.md)
- [xcsh_app_type](../data-sources/app_type.md)
