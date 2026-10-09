---
page_title: "app_type_settings.business_logic_markup_setting"
subcategory: ""
description: "Settings specifying how API Discovery will be performed."
xcsh_docs: {"aliases": ["app type settings business logic markup setting"], "body_bytes": 1345, "body_sha256": "sha256:0f04165c1c7e2a37a62a20cf80bbe699c963c15d857155e638e275598b0d15e9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting:disable_spec", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting:enable"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting", "parent_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings", "path": "documentation/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323", "registry_path": "docs/guides/data-sources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "business_logic_markup_setting"], "schema_version": 1, "sections": [{"aliases": ["app type settings business logic markup setting disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting:disable_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "business_logic_markup_setting", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings business logic markup setting enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting:enable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "business_logic_markup_setting", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Settings specifying how API Discovery will be performed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["app_settingCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.business_logic_markup_setting

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/)
- app_type_settings.business_logic_markup_setting

<a id="section"></a>

Type: `"single"`. Computed.

Settings specifying how API Discovery will be performed.

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

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/disable_spec/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/enable/): complete subsection reference.
