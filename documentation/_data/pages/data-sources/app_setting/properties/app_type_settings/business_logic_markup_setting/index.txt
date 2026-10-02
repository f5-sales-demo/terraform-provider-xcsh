---
page_title: "app_type_settings.business_logic_markup_setting"
subcategory: ""
description: "Settings specifying how API Discovery will be performed."
xcsh_docs: {"aliases": ["app type settings business logic markup setting"], "body_bytes": 2027, "body_sha256": "sha256:2c4bef62dbfb13149f71bbc45f6825f82e0bde96616ace967e669130947b7084", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting:disable_spec", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting:enable"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting", "parent_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings", "path": "documentation/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2200113301001221-0032220121230103-3013132020110202-1300003311122230-2331010030321210-3002331313020231-3221120021332233-3310002031123323", "registry_path": "docs/guides/data-sources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "business_logic_markup_setting"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting:disable_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "business_logic_markup_setting", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting:enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "business_logic_markup_setting", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Settings specifying how API Discovery will be performed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_settingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [app_type_settings.business_logic_markup_setting.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/disable_spec/)
- [app_type_settings.business_logic_markup_setting.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/enable/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
