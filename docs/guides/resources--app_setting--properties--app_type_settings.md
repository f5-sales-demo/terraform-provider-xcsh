---
page_title: "app_type_settings"
subcategory: ""
description: "app_type_settings for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 2723, "body_sha256": "sha256:a5be4967db4e49943a4433ae189936f6eea95cfe81c4f3844c0e1b5f2b8d99f5", "canonical_id": "xcsh-docs:resources:app_setting:properties:app_type_settings", "child_ids": ["xcsh-docs:resources:app_setting:properties:app_type_settings:app_type_ref", "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting", "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting"], "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings", "parent_id": "xcsh-docs:resources:app_setting:reference", "path": "docs/guides/resources--app_setting--properties--app_type_settings.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["app_type_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "app_type_settings for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# app_type_settings

Breadcrumbs:

- [xcsh_app_setting](../resources/app_setting.md)
- [Property reference](resources--app_setting--reference.md)
- app_type_settings

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("app_type_ref")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
app_type_settings {
  # Configure direct properties listed below.
}
```

## Direct properties

- [app_type_ref](resources--app_setting--properties--app_type_settings--app_type_ref.md): complete subsection reference.

- [business_logic_markup_setting](resources--app_setting--properties--app_type_settings--business_logic_markup_setting.md): complete subsection reference.

- [timeseries_analyses_setting](resources--app_setting--properties--app_type_settings--timeseries_analyses_setting.md): complete subsection reference.

- [user_behavior_analysis_setting](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting.md): complete subsection reference.

## Next pages

- [app_type_settings.app_type_ref](resources--app_setting--properties--app_type_settings--app_type_ref.md)
- [app_type_settings.business_logic_markup_setting](resources--app_setting--properties--app_type_settings--business_logic_markup_setting.md)
- [app_type_settings.timeseries_analyses_setting](resources--app_setting--properties--app_type_settings--timeseries_analyses_setting.md)
- [app_type_settings.user_behavior_analysis_setting](resources--app_setting--properties--app_type_settings--user_behavior_analysis_setting.md)
- [Property reference](resources--app_setting--reference.md)
- [xcsh_app_setting](../resources/app_setting.md)
