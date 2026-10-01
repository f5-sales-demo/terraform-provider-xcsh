---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_srv6_network_slice."
xcsh_docs: {"aliases": [], "body_bytes": 1354, "body_sha256": "sha256:911dea919f59d95b6d7c7573726192bbc7c9e3215f9f1dc5c82777c62393058a", "child_ids": [], "collection_id": "xcsh-docs:data-sources:srv6_network_slice:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5f3b6cfc9cb152dc2c604228803d1e4acdb80bbc9c71bb42003bba14ed6bbdd1", "source_path": "examples/data-sources/xcsh_srv6_network_slice/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:srv6_network_slice:example:data-source", "parent_id": "xcsh-docs:data-sources:srv6_network_slice:examples", "path": "documentation/data-sources/srv6_network_slice/examples/data-source/index.md", "provider_name": "srv6_network_slice", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/srv6_network_slice/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_srv6_network_slice.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["srv6_network_sliceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_srv6_network_slice](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_srv6_network_slice/data-source.tf`; digest `sha256:5f3b6cfc9cb152dc2c604228803d1e4acdb80bbc9c71bb42003bba14ed6bbdd1`.

```terraform
# Srv6NetworkSlice Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Srv6NetworkSlice by name
data "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"
}

output "srv6_network_slice_id" {
  value = data.xcsh_srv6_network_slice.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/examples/)
- [xcsh_srv6_network_slice](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/srv6_network_slice/)
