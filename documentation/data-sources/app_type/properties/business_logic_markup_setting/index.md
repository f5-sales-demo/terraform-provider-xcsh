---
page_title: "business_logic_markup_setting"
subcategory: ""
description: "Settings specifying how API Discovery will be performed."
xcsh_docs: {"aliases": ["business logic markup setting"], "body_bytes": 2118, "body_sha256": "sha256:52b02ba84c597d94f838899e3dfeb28d427c7fb9257f16a480bce39aed9d892b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:disable_spec", "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:discovered_api_settings", "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:enable"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting", "parent_id": "xcsh-docs:data-sources:app_type:reference", "path": "documentation/data-sources/app_type/properties/business_logic_markup_setting/index.md", "product": "distributed-cloud", "provider_name": "app_type", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2323122110222310-1232231121122003-3322001112111020-3333201210332002-0003210220103112-1211021101213100-0313132102013202-3311311013201011", "registry_path": "docs/guides/data-sources--app_type--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["business_logic_markup_setting"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:disable_spec", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["business_logic_markup_setting", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovered api settings"], "anchor": "section", "description": "Configure Discovered API Settings.", "document_id": "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:discovered_api_settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["business_logic_markup_setting", "discovered_api_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_type:properties:business_logic_markup_setting:enable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["business_logic_markup_setting", "enable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_type/properties/business_logic_markup_setting/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Settings specifying how API Discovery will be performed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_typeCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [business_logic_markup_setting.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/business_logic_markup_setting/disable_spec/)
- [business_logic_markup_setting.discovered_api_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/business_logic_markup_setting/discovered_api_settings/)
- [business_logic_markup_setting.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/business_logic_markup_setting/enable/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/)
- [xcsh_app_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/)
