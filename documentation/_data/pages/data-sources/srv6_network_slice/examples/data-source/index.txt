---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_srv6_network_slice."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1354, "body_sha256": "sha256:911dea919f59d95b6d7c7573726192bbc7c9e3215f9f1dc5c82777c62393058a", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:srv6_network_slice:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5f3b6cfc9cb152dc2c604228803d1e4acdb80bbc9c71bb42003bba14ed6bbdd1", "source_path": "examples/data-sources/xcsh_srv6_network_slice/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:srv6_network_slice:example:data-source", "parent_id": "xcsh-docs:data-sources:srv6_network_slice:examples", "path": "documentation/data-sources/srv6_network_slice/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "srv6_network_slice", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3202233013313201-3100332002310112-3103033200121200-3031101001110300-1210103233130211-1122300322000211-2132011222121230-1212022312211011", "registry_path": "docs/guides/data-sources--srv6_network_slice--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/srv6_network_slice/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_srv6_network_slice.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["srv6_network_sliceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
