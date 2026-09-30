---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 961, "body_sha256": "sha256:127c0bb9010aa36cc2e825d9e8d265a84c1c5b1f71da59dff29bd7d5e68c7a34", "canonical_id": "xcsh-docs:data-sources:app_setting:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4e31d15f083391674e4e309fb8e0ab1858f7a5c2638cd424cc62bee3775dfde7", "source_path": "examples/data-sources/xcsh_app_setting/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_setting:example:data-source", "parent_id": "xcsh-docs:data-sources:app_setting:examples", "path": "docs/guides/data-sources--app_setting--example--data-source.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md)
- [Examples](data-sources--app_setting--examples.md)
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

- [Examples](data-sources--app_setting--examples.md)
- [xcsh_app_setting](../data-sources/app_setting.md)
