---
page_title: "app_type_settings.user_behavior_analysis_setting"
subcategory: ""
description: "app_type_settings.user_behavior_analysis_setting for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 2330, "body_sha256": "sha256:51d15e6b3ede3119c91389d57beeb84220510d045649153ae18d9f7d7cfad00a", "canonical_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting", "child_ids": ["xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:disable_detection", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:disable_learning", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_learning"], "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting", "parent_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings", "path": "docs/guides/data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "app_type_settings.user_behavior_analysis_setting for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# app_type_settings.user_behavior_analysis_setting

Breadcrumbs:

- [xcsh_app_setting](../data-sources/app_setting.md)
- [Property reference](data-sources--app_setting--reference.md)
- [app_type_settings](data-sources--app_setting--properties--app_type_settings.md)
- app_type_settings.user_behavior_analysis_setting

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for user behavior analysis.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-learn_from_namespace": "[\"disable_learning\",\"enable_learning\"]",
  "x-ves-oneof-field-malicious_user_detection": "[\"disable_detection\",\"enable_detection\"]"
}
```

## Direct properties

- [disable_detection](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--disable_detection.md): complete subsection reference.

- [disable_learning](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--disable_learning.md): complete subsection reference.

- [enable_detection](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection.md): complete subsection reference.

- [enable_learning](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_learning.md): complete subsection reference.

## Next pages

- [app_type_settings.user_behavior_analysis_setting.disable_detection](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--disable_detection.md)
- [app_type_settings.user_behavior_analysis_setting.disable_learning](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--disable_learning.md)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection.md)
- [app_type_settings.user_behavior_analysis_setting.enable_learning](data-sources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_learning.md)
- [app_type_settings](data-sources--app_setting--properties--app_type_settings.md)
- [xcsh_app_setting](../data-sources/app_setting.md)
