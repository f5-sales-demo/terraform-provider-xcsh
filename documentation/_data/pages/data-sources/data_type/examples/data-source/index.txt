---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_data_type."
xcsh_docs: {"aliases": [], "body_bytes": 1141, "body_sha256": "sha256:219de3c2d8e728c1cc5f370b29e85678070c02b8ab1786e2e1cf7e96ac8c4d0d", "child_ids": [], "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2ab6eae5bd6da2f1828ca8b3f7f9c368730657551b0232c70f3bd6f7dc6e3501", "source_path": "examples/data-sources/xcsh_data_type/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:data_type:example:data-source", "parent_id": "xcsh-docs:data-sources:data_type:examples", "path": "documentation/data-sources/data_type/examples/data-source/index.md", "provider_name": "data_type", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_data_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/examples/)
- [xcsh_data_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/)
