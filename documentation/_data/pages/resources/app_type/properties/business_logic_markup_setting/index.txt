---
page_title: "business_logic_markup_setting"
subcategory: ""
description: "business_logic_markup_setting for xcsh_app_type."
xcsh_docs: {"aliases": [], "body_bytes": 2394, "body_sha256": "sha256:93fdb990c408eeb9f1be598f589b7443588b6cb21820d031b63a07ba07794095", "child_ids": ["xcsh-docs:resources:app_type:properties:business_logic_markup_setting:disable_spec", "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:discovered_api_settings", "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:enable"], "collection_id": "xcsh-docs:resources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting", "parent_id": "xcsh-docs:resources:app_type:reference", "path": "documentation/resources/app_type/properties/business_logic_markup_setting/index.md", "provider_name": "app_type", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["business_logic_markup_setting"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_type/properties/business_logic_markup_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "business_logic_markup_setting for xcsh_app_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# business_logic_markup_setting

Breadcrumbs:

- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/)
- business_logic_markup_setting

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
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
business_logic_markup_setting {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/business_logic_markup_setting/disable_spec/): complete subsection reference.

- [discovered_api_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/business_logic_markup_setting/discovered_api_settings/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/business_logic_markup_setting/enable/): complete subsection reference.

## Next pages

- [business_logic_markup_setting.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/business_logic_markup_setting/disable_spec/)
- [business_logic_markup_setting.discovered_api_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/business_logic_markup_setting/discovered_api_settings/)
- [business_logic_markup_setting.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/business_logic_markup_setting/enable/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/)
- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/)
