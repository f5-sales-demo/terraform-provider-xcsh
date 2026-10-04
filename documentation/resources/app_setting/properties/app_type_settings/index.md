---
page_title: "app_type_settings"
subcategory: ""
description: "List of settings to enable for each AppType, given instance of AppType Exist in this Namespace."
xcsh_docs: {"aliases": ["app type settings"], "body_bytes": 3430, "body_sha256": "sha256:3e5112b92c25bf175253ad4ce91cf007be666a703c086a59fa4e7ea3075b54e2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:app_setting:properties:app_type_settings:app_type_ref", "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting", "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings", "parent_id": "xcsh-docs:resources:app_setting:reference", "path": "documentation/resources/app_setting/properties/app_type_settings/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1232123221311121-0303331012300320-3211001301331213-1232111200120323-0313032203120030-0230001103300323-0213103133030322-3120300003031112", "registry_path": "docs/guides/resources--app_setting--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "app_type_settings:RequiredListObjectAttributes:app_type_ref", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:app_type_ref", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings"], "schema_version": 1, "sections": [{"aliases": ["app type settings app type ref"], "anchor": "section", "description": "The AppType of App instance in current Namespace. Associating an AppType reference, will enable analysis on this instance's generated data.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:app_type_ref", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["app_type_settings", "app_type_ref"], "syntax": "block", "type": "object"}, {"aliases": ["app type settings business logic markup setting"], "anchor": "section", "description": "Settings specifying how API Discovery will be performed.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "app_type_settings.business_logic_markup_setting:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "app_type_settings.business_logic_markup_setting:ConflictingObjectAttributes:disable_spec,enable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:business_logic_markup_setting:enable", "type": "conflicts"}], "schema_path": ["app_type_settings", "business_logic_markup_setting"], "syntax": "block", "type": "object"}, {"aliases": ["app type settings timeseries analyses setting"], "anchor": "section", "description": "Configuration for DDoS Detection.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:timeseries_analyses_setting", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["app_type_settings", "timeseries_analyses_setting"], "syntax": "block", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting"], "anchor": "section", "description": "Configuration for user behavior analysis.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "app_type_settings.user_behavior_analysis_setting:ConflictingObjectAttributes:disable_detection,enable_detection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:disable_detection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "app_type_settings.user_behavior_analysis_setting:ConflictingObjectAttributes:disable_learning,enable_learning", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:disable_learning", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "app_type_settings.user_behavior_analysis_setting:ConflictingObjectAttributes:disable_detection,enable_detection", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "app_type_settings.user_behavior_analysis_setting:ConflictingObjectAttributes:disable_learning,enable_learning", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_learning", "type": "conflicts"}], "schema_path": ["app_type_settings", "user_behavior_analysis_setting"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["app_settingCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/)
- app_type_settings

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of settings to enable for each AppType, given instance of AppType Exist in this Namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("app_type_ref")}
```

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

Terraform syntax:

```terraform
app_type_settings {
  # Configure direct properties listed below.
}
```

## Direct properties

- [app_type_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/app_type_ref/): complete subsection reference.

- [business_logic_markup_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/business_logic_markup_setting/): complete subsection reference.

- [timeseries_analyses_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/): complete subsection reference.

- [user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/): complete subsection reference.

## Next pages

- [app_type_settings.app_type_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/app_type_ref/)
- [app_type_settings.business_logic_markup_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/business_logic_markup_setting/)
- [app_type_settings.timeseries_analyses_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/timeseries_analyses_setting/)
- [app_type_settings.user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
