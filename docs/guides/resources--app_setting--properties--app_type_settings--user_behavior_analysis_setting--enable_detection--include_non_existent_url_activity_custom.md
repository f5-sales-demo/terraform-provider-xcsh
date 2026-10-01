---
page_title: "app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom"
subcategory: ""
description: "app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 2968, "body_sha256": "sha256:f8c4b2db1babc4e3f118e4226737129af5bbac263b7611987d7e43d243b79e9c", "canonical_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_custom", "child_ids": [], "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_custom", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "path": "docs/guides/resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_non_existent_url_activity_custom"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_custom/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md)
- [Property reference](resources--app_setting--reference.md)
- [app_type_settings](resources--app_setting--properties--app_type_settings.md)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting.md)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection.md)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Non-existent URL Custom Activity Setting.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("nonexistent_requests_threshold")}
```

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
include_non_existent_url_activity_custom {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom--nonexistent_requests_threshold"></a>

### nonexistent_requests_threshold property

Type: `"number"`. Optional.

The percentage of non-existent requests beyond which the system will flag this user as malicious.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

## Next pages

- [app_type_settings.user_behavior_analysis_setting.enable_detection](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting--enable_detection.md)
- [xcsh_app_setting](../resources/app_setting.md)
