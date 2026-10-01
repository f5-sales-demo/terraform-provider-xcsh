---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_data_type."
xcsh_docs: {"aliases": [], "body_bytes": 1034, "body_sha256": "sha256:19a945cc4d0762d96fc88094535f653d77616554f646d213f2dd7e3bb26ab302", "canonical_id": "xcsh-docs:data-sources:data_type:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2ab6eae5bd6da2f1828ca8b3f7f9c368730657551b0232c70f3bd6f7dc6e3501", "source_path": "examples/data-sources/xcsh_data_type/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:data_type:example:data-source", "parent_id": "xcsh-docs:data-sources:data_type:examples", "path": "docs/guides/data-sources--data_type--example--data-source.md", "provider_name": "data_type", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_data_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_data_type](../data-sources/data_type.md)
- [Examples](data-sources--data_type--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_data_type/data-source.tf`; digest `sha256:2ab6eae5bd6da2f1828ca8b3f7f9c368730657551b0232c70f3bd6f7dc6e3501`.

```terraform
# DataType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DataType by name
data "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}

output "data_type_id" {
  value = data.xcsh_data_type.example.id
}
```

## Next pages

- [Examples](data-sources--data_type--examples.md)
- [xcsh_data_type](../data-sources/data_type.md)
