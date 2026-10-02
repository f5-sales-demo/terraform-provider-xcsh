---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_type."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1227, "body_sha256": "sha256:5f266ee60e1edb80174125d85b31d31577c7d4913f28f813e58cfa0897372f26", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_type:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:7219c79a69a1e857679b66c93800c9cdd60d2a0ca9b2ed275f17e432b1f04933", "source_path": "examples/data-sources/xcsh_app_type/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_type:example:data-source", "parent_id": "xcsh-docs:data-sources:app_type:examples", "path": "documentation/data-sources/app_type/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "app_type", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2120332221223021-0023301011011233-2130011301010102-0211320120213212-0002010313231021-2230011131203330-0102221100001033-1300033301010133", "registry_path": "docs/guides/data-sources--app_type--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_type/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_app_type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_typeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/examples/)
- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/)
