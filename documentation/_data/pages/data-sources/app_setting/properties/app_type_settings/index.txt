---
page_title: "app_type_settings"
subcategory: ""
description: "List of settings to enable for each AppType, given instance of AppType Exist in this Namespace."
xcsh_docs: {"aliases": ["app type settings"], "body_bytes": 3187, "body_sha256": "sha256:f107cf0d197733814b62b3bfc6f5974198fbc3cb43b8b3a22481ae4da715111c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_setting:properties:app_type_settings:app_type_ref", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings", "parent_id": "xcsh-docs:data-sources:app_setting:reference", "path": "documentation/data-sources/app_setting/properties/app_type_settings/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3001122222232012-1331202012100130-0330330233110010-1002223122303020-2221321010301213-3012310311112012-3331322313022011-3033021023033331", "registry_path": "docs/guides/data-sources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings"], "schema_version": 1, "sections": [{"aliases": ["app type settings app type ref"], "anchor": "section", "description": "The AppType of App instance in current Namespace. Associating an AppType reference, will enable analysis on this instance's generated data.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:app_type_ref", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["app_type_settings", "app_type_ref"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings business logic markup setting"], "anchor": "section", "description": "Settings specifying how API Discovery will be performed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:business_logic_markup_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["app_type_settings", "business_logic_markup_setting"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings timeseries analyses setting"], "anchor": "section", "description": "Configuration for DDoS Detection.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["app_type_settings", "timeseries_analyses_setting"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting"], "anchor": "section", "description": "Configuration for user behavior analysis.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/app_type_settings/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["app_settingCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/)
- app_type_settings

<a id="section"></a>

Type: `"list"`. Computed.

List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Direct properties

- [app_type_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/app_type_ref/): complete subsection reference.

- [business_logic_markup_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/): complete subsection reference.

- [timeseries_analyses_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/): complete subsection reference.

- [user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/): complete subsection reference.

## Next pages

- [app_type_settings.app_type_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/app_type_ref/)
- [app_type_settings.business_logic_markup_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/business_logic_markup_setting/)
- [app_type_settings.timeseries_analyses_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/timeseries_analyses_setting/)
- [app_type_settings.user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
