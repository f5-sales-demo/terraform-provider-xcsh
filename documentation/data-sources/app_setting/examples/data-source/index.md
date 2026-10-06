---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_setting."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1041, "body_sha256": "sha256:f179ffc3dcbb5322e32d2691e6694b53d813cf054df9097c7f6209e908f82123", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4e31d15f083391674e4e309fb8e0ab1858f7a5c2638cd424cc62bee3775dfde7", "source_path": "examples/data-sources/xcsh_app_setting/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_setting:example:data-source", "parent_id": "xcsh-docs:data-sources:app_setting:examples", "path": "documentation/data-sources/app_setting/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3213032032121013-1100310011302030-1002001322301111-2113112133310122-0230113232020223-0130100102302123-3022203102030231-1003020213100130", "registry_path": "docs/guides/data-sources--app_setting--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_app_setting.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["app_settingCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
