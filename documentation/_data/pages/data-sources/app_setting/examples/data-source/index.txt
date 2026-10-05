---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_setting."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1266, "body_sha256": "sha256:15a7551581ee1635e11be6eaa186eb760a65521e5ed3b0854e4e8e6837081a15", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4e31d15f083391674e4e309fb8e0ab1858f7a5c2638cd424cc62bee3775dfde7", "source_path": "examples/data-sources/xcsh_app_setting/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_setting:example:data-source", "parent_id": "xcsh-docs:data-sources:app_setting:examples", "path": "documentation/data-sources/app_setting/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3213032032121013-1100310011302030-1002001322301111-2113112133310122-0230113232020223-0130100102302123-3022203102030231-1003020213100130", "registry_path": "docs/guides/data-sources--app_setting--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_app_setting.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
