---
page_title: "xcsh_app_setting"
subcategory: ""
description: "xcsh_app_setting for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 1366, "body_sha256": "sha256:5d8e57f1c01f97b2dd5dc7dd22600f3538c1af29fbb200175c19b41a876a4186", "canonical_id": "xcsh-docs:resources:app_setting:fundamentals", "child_ids": ["xcsh-docs:resources:app_setting:reference", "xcsh-docs:resources:app_setting:examples", "xcsh-docs:resources:app_setting:import", "xcsh-docs:resources:app_setting:timeouts"], "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:fundamentals", "parent_id": null, "path": "docs/resources/app_setting.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_app_setting for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_app_setting

Breadcrumbs:

- xcsh_app_setting

Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--app_setting--reference.md)
- [Examples](../guides/resources--app_setting--examples.md)
- [Import](../guides/resources--app_setting--import.md)
- [Timeouts](../guides/resources--app_setting--timeouts.md)
