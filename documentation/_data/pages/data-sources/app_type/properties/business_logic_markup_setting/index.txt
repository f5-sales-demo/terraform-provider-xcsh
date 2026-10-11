---
page_title: "business_logic_markup_setting"
subcategory: ""
description: "Settings specifying how API Discovery will be performed."
xcsh_docs: {"aliases": ["business logic markup setting"], "body_bytes": 1337, "body_sha256": "sha256:c4a5e36a213abe1045d96b81c8dc6109c81e7f2bd33834151fa61b10c514265d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:disable_spec", "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:discovered_api_settings", "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:enable"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting", "parent_id": "xcsh-docs:data-sources:app_type:reference", "path": "documentation/data-sources/app_type/properties/business_logic_markup_setting/index.md", "product": "distributed-cloud", "provider_name": "app_type", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-2323122110222310-1232231121122003-3322001112111020-3333201210332002-0003210220103112-1211021101213100-0313132102013202-3311311013201011", "registry_path": "docs/guides/data-sources--app_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["business_logic_markup_setting"], "schema_version": 1, "sections": [{"aliases": ["business logic markup setting disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:disable_spec", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["business_logic_markup_setting", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["business logic markup setting discovered api settings"], "anchor": "section", "description": "Configure Discovered API Settings.", "document_id": "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:discovered_api_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["business_logic_markup_setting", "discovered_api_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["business logic markup setting enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:enable", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["business_logic_markup_setting", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_type/properties/business_logic_markup_setting/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Settings specifying how API Discovery will be performed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["app_typeCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# business_logic_markup_setting

Breadcrumbs:

- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/)
- business_logic_markup_setting

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
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable\",\"enable\"]"
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/business_logic_markup_setting/disable_spec/): complete subsection reference.

- [discovered_api_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/business_logic_markup_setting/discovered_api_settings/): complete subsection reference.

- [enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/business_logic_markup_setting/enable/): complete subsection reference.
