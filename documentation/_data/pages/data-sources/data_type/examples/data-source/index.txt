---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_data_type."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1240, "body_sha256": "sha256:5431e4da62582ba75f1ee56dc1d90f39a0aeaec1c623ade52095c065cf93b150", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2ab6eae5bd6da2f1828ca8b3f7f9c368730657551b0232c70f3bd6f7dc6e3501", "source_path": "examples/data-sources/xcsh_data_type/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:data_type:example:data-source", "parent_id": "xcsh-docs:data-sources:data_type:examples", "path": "documentation/data-sources/data_type/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3300131002013102-3330112100203301-0303203323233031-3300233202232332-0300221123322012-1000131010203103-2323123300111103-0211020233123100", "registry_path": "docs/guides/data-sources--data_type--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_data_type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["data_typeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
