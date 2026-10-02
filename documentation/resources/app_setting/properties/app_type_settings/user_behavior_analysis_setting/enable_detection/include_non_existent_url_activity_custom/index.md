---
page_title: "app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom"
subcategory: ""
description: "Non-existent URL Custom Activity Setting."
xcsh_docs: {"aliases": ["app type settings user behavior analysis setting enable detection include non existent url activity custom"], "body_bytes": 3322, "body_sha256": "sha256:d0b3d9882f31771c223442b4a8022fc5977d2378e8d240f6c8f6c35f79c8ba8e", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_custom", "parent_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "path": "documentation/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_custom/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3033220110323013-2323013003303301-1033231123310221-0132003332000112-3223301221302121-3321102103000233-3133131310032101-0123333321232123", "registry_path": "docs/guides/resources--app_setting--reference--group-001.md", "relationships": [{"anchor": "schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom--nonexistent_requests_threshold", "enforcement": "provider-schema", "group": "app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom:RequiredObjectAttributes:nonexistent_requests_threshold", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_custom", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_non_existent_url_activity_custom"], "schema_version": 1, "sections": [{"aliases": ["nonexistent requests threshold"], "anchor": "schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom--nonexistent_requests_threshold", "description": "The percentage of non-existent requests beyond which the system will flag this user as malicious.", "document_id": "xcsh-docs:resources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_custom", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_non_existent_url_activity_custom", "nonexistent_requests_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_custom/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Non-existent URL Custom Activity Setting.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_settingCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/)
- [app_type_settings.user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Non-existent URL Custom Activity Setting.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("nonexistent_requests_threshold")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
include_non_existent_url_activity_custom {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom--nonexistent_requests_threshold"></a>

### nonexistent_requests_threshold property

Type: `"number"`. Optional.

The percentage of non-existent requests beyond which the system will flag this user as malicious.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

## Next pages

- [app_type_settings.user_behavior_analysis_setting.enable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_setting/)
