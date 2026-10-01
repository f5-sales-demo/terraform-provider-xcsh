---
page_title: "app_type_settings.business_logic_markup_setting"
subcategory: ""
description: "app_type_settings.business_logic_markup_setting for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 2306, "body_sha256": "sha256:479f71344ac8679200d71a223be42c7cb9205378caa7a61be6df645e472b9767", "child_ids": ["xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting:disable_spec", "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting:enable"], "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings", "path": "documentation/resources/app_setting/properties/app_type_settings/business_logic_markup_setting/index.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["app_type_settings", "business_logic_markup_setting"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/business_logic_markup_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "app_type_settings.business_logic_markup_setting for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.business_logic_markup_setting

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/)
- app_type_settings.business_logic_markup_setting

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Settings specifying how API Discovery will be performed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-learn_from_namespace": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
business_logic_markup_setting {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/business_logic_markup_setting/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/business_logic_markup_setting/enable/): complete subsection reference.

## Next pages

- [app_type_settings.business_logic_markup_setting.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/business_logic_markup_setting/disable_spec/)
- [app_type_settings.business_logic_markup_setting.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/business_logic_markup_setting/enable/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
