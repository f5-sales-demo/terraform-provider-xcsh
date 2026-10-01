---
page_title: "app_type_settings.user_behavior_analysis_setting.disable_detection"
subcategory: ""
description: "app_type_settings.user_behavior_analysis_setting.disable_detection for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 1311, "body_sha256": "sha256:a23bd92aa241c55e6630caf786976c19b1b15d29dba405f413f8507e9a0a394e", "canonical_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:disable_detection", "child_ids": [], "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:disable_detection", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting", "path": "docs/guides/resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--disable_detection.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "disable_detection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/disable_detection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "app_type_settings.user_behavior_analysis_setting.disable_detection for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.user_behavior_analysis_setting.disable_detection

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md)
- [Property reference](resources--app_setting--reference.md)
- [app_type_settings](resources--app_setting--properties--app_type_settings.md)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting.md)
- app_type_settings.user_behavior_analysis_setting.disable_detection

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable detection.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
disable_detection = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting.md)
- [xcsh_app_setting](../resources/app_setting.md)
