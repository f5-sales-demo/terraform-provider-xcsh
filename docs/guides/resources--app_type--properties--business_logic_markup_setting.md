---
page_title: "business_logic_markup_setting"
subcategory: ""
description: "business_logic_markup_setting for xcsh_app_type."
xcsh_docs: {"aliases": [], "body_bytes": 1787, "body_sha256": "sha256:043572968b81d9c3c0da33b72eae34d2f35a3cf0b1e0babe20cad961c0bff7ed", "canonical_id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting", "child_ids": ["xcsh-docs:resources:app_type:properties:business_logic_markup_setting:disable_spec", "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:discovered_api_settings", "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:enable"], "collection_id": "xcsh-docs:resources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting", "parent_id": "xcsh-docs:resources:app_type:reference", "path": "docs/guides/resources--app_type--properties--business_logic_markup_setting.md", "provider_name": "app_type", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["business_logic_markup_setting"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_type/properties/business_logic_markup_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "business_logic_markup_setting for xcsh_app_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# business_logic_markup_setting

Breadcrumbs:

- [xcsh_app_type](../resources/app_type.md)
- [Property reference](resources--app_type--reference.md)
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

- [disable_spec](resources--app_type--properties--business_logic_markup_setting--disable_spec.md): complete subsection reference.

- [discovered_api_settings](resources--app_type--properties--business_logic_markup_setting--discovered_api_settings.md): complete subsection reference.

- [enable](resources--app_type--properties--business_logic_markup_setting--enable.md): complete subsection reference.

## Next pages

- [business_logic_markup_setting.disable_spec](resources--app_type--properties--business_logic_markup_setting--disable_spec.md)
- [business_logic_markup_setting.discovered_api_settings](resources--app_type--properties--business_logic_markup_setting--discovered_api_settings.md)
- [business_logic_markup_setting.enable](resources--app_type--properties--business_logic_markup_setting--enable.md)
- [Property reference](resources--app_type--reference.md)
- [xcsh_app_type](../resources/app_type.md)
