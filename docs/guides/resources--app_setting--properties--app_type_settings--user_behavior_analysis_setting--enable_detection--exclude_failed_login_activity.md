---
page_title: "app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity"
subcategory: ""
description: "app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 1606, "body_sha256": "sha256:7a464e69c117e2a8b40d4fe72d8303e90861fdbc283fa1137193b7e706f6b600", "canonical_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_failed_login_activity", "child_ids": [], "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_failed_login_activity", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "path": "docs/guides/resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--exclude_failed_login_activity.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "exclude_failed_login_activity"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_failed_login_activity/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md)
- [Property reference](resources--app_setting--reference.md)
- [app_type_settings](resources--app_setting--properties--app_type_settings.md)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting.md)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection.md)
- app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for exclude failed login activity.

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
exclude_failed_login_activity = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection.md)
- [xcsh_app_setting](../resources/app_setting.md)
