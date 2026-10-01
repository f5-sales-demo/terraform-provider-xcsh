---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 1256, "body_sha256": "sha256:9dcbc868991e81c5d5d9eafb50d2b643943326d024dd408205d1b7b49106c60b", "child_ids": [], "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d31cbf807fc9dcf209556e05af346ee32ebdbf440fd52fffb2c11e50ae1d1f1a", "source_path": "examples/resources/xcsh_app_setting/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_setting:example:resource", "parent_id": "xcsh-docs:resources:app_setting:examples", "path": "documentation/resources/app_setting/examples/resource/index.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_setting/resource.tf`; digest `sha256:d31cbf807fc9dcf209556e05af346ee32ebdbf440fd52fffb2c11e50ae1d1f1a`.

```terraform
# AppSetting Resource Example
# Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppSetting configuration
resource "xcsh_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/examples/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
