---
page_title: "business_logic_markup_setting"
subcategory: ""
description: "Settings specifying how API Discovery will be performed."
xcsh_docs: {"aliases": ["business logic markup setting"], "body_bytes": 1658, "body_sha256": "sha256:719769cfc9cb83f009dd3a3818ff50f04ea8551ad4b61fd8e649b7e1ce55111a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:app_type:properties:business_logic_markup_setting:disable_spec", "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:discovered_api_settings", "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:enable"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting", "parent_id": "xcsh-docs:resources:app_type:reference", "path": "documentation/resources/app_type/properties/business_logic_markup_setting/index.md", "product": "distributed-cloud", "provider_name": "app_type", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3223322110003203-1122212310332100-2122030232112232-1223332130013002-0311230312202212-3223203333012301-0103213121203013-3021020310331223", "registry_path": "docs/guides/resources--app_type--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "business_logic_markup_setting:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "business_logic_markup_setting:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:enable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["business_logic_markup_setting"], "schema_version": 1, "sections": [{"aliases": ["business logic markup setting disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:disable_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["business_logic_markup_setting", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["business logic markup setting discovered api settings"], "anchor": "section", "description": "Configure Discovered API Settings.", "document_id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:discovered_api_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-business_logic_markup_setting--discovered_api_settings--purge_duration_for_inactive_discovered_apis", "enforcement": "provider-schema", "group": "business_logic_markup_setting.discovered_api_settings:RequiredObjectAttributes:purge_duration_for_inactive_discovered_apis", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:discovered_api_settings", "type": "requires"}], "schema_path": ["business_logic_markup_setting", "discovered_api_settings"], "syntax": "block", "type": "object"}, {"aliases": ["business logic markup setting enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:app_type:properties:business_logic_markup_setting:enable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["business_logic_markup_setting", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_type/properties/business_logic_markup_setting/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Settings specifying how API Discovery will be performed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["app_typeCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
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
EnumExtractionComplete: false
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
