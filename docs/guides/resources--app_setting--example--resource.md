---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 1050, "body_sha256": "sha256:6ec3115b800e3f6d3b1f7b8f349e5b5285a032a6313bbdfb00f63232d6f04a2d", "canonical_id": "xcsh-docs:resources:app_setting:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d31cbf807fc9dcf209556e05af346ee32ebdbf440fd52fffb2c11e50ae1d1f1a", "source_path": "examples/resources/xcsh_app_setting/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_setting:example:resource", "parent_id": "xcsh-docs:resources:app_setting:examples", "path": "docs/guides/resources--app_setting--example--resource.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md)
- [Examples](resources--app_setting--examples.md)
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

- [Examples](resources--app_setting--examples.md)
- [xcsh_app_setting](../resources/app_setting.md)
