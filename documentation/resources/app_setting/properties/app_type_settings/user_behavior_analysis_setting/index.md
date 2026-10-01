---
page_title: "app_type_settings.user_behavior_analysis_setting"
subcategory: ""
description: "app_type_settings.user_behavior_analysis_setting for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 3446, "body_sha256": "sha256:64537beb5fe682246f9ec8ca182027a8892d65af06469992fd7d4767008f7989", "child_ids": ["xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:disable_detection", "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:disable_learning", "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_learning"], "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings", "path": "documentation/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/index.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "app_type_settings.user_behavior_analysis_setting for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.user_behavior_analysis_setting

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/)
- app_type_settings.user_behavior_analysis_setting

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration for user behavior analysis.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_detection",
    "enable_detection"),
  validators.ConflictingObjectAttributes("disable_learning",
    "enable_learning")}
```

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

Terraform syntax:

```terraform
user_behavior_analysis_setting {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/disable_detection/): complete subsection reference.

- [disable_learning](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/disable_learning/): complete subsection reference.

- [enable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/): complete subsection reference.

- [enable_learning](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_learning/): complete subsection reference.

## Next pages

- [app_type_settings.user_behavior_analysis_setting.disable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/disable_detection/)
- [app_type_settings.user_behavior_analysis_setting.disable_learning](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/disable_learning/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/)
- [app_type_settings.user_behavior_analysis_setting.enable_learning](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_learning/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
