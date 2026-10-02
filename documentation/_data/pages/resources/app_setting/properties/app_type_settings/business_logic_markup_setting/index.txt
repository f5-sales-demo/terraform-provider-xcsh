---
page_title: "app_type_settings.business_logic_markup_setting"
subcategory: ""
description: "Settings specifying how API Discovery will be performed."
xcsh_docs: {"aliases": ["app type settings business logic markup setting"], "body_bytes": 2306, "body_sha256": "sha256:479f71344ac8679200d71a223be42c7cb9205378caa7a61be6df645e472b9767", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting:disable_spec", "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting:enable"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings", "path": "documentation/resources/app_setting/properties/app_type_settings/business_logic_markup_setting/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0312220223320022-2123030223120221-1320010130121130-2131212110322203-0323311113313203-2231121310211321-3123331110313203-0132120211320032", "registry_path": "docs/guides/resources--app_setting--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "app_type_settings.business_logic_markup_setting:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "app_type_settings.business_logic_markup_setting:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting:enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "business_logic_markup_setting"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting:disable_spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "business_logic_markup_setting", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting:enable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "business_logic_markup_setting", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/business_logic_markup_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Settings specifying how API Discovery will be performed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
