---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_setting."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1266, "body_sha256": "sha256:15a7551581ee1635e11be6eaa186eb760a65521e5ed3b0854e4e8e6837081a15", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4e31d15f083391674e4e309fb8e0ab1858f7a5c2638cd424cc62bee3775dfde7", "source_path": "examples/data-sources/xcsh_app_setting/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_setting:example:data-source", "parent_id": "xcsh-docs:data-sources:app_setting:examples", "path": "documentation/data-sources/app_setting/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3213032032121013-1100310011302030-1002001322301111-2113112133310122-0230113232020223-0130100102302123-3022203102030231-1003020213100130", "registry_path": "docs/guides/data-sources--app_setting--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_app_setting.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_settingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_setting/data-source.tf`; digest `sha256:4e31d15f083391674e4e309fb8e0ab1858f7a5c2638cd424cc62bee3775dfde7`.

```terraform
# AppSetting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppSetting by name
data "xcsh_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "staging"
}

output "app_setting_id" {
  value = data.xcsh_app_setting.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/examples/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
