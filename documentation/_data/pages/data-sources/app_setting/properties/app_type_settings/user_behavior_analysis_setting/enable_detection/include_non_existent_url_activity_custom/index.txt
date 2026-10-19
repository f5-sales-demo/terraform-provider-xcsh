---
page_title: "app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom"
subcategory: ""
description: "Non-existent URL Custom Activity Setting."
xcsh_docs: {"aliases": ["app type settings user behavior analysis setting enable detection include non existent url activity custom"], "body_bytes": 2541, "body_sha256": "sha256:e23c3980e632f94c758fbc08554a01d78bc18c95fc30da3cf93b0ed64315bdf8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_custom", "parent_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "path": "documentation/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_custom/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1113332013020121-3220203212302100-2032013132131110-2312000021021212-1112023022133003-2021332321011002-1321101132010211-3101302110213130", "registry_path": "docs/guides/data-sources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_non_existent_url_activity_custom"], "schema_version": 1, "sections": [{"aliases": ["app type settings user behavior analysis setting enable detection include non existent url activity custom nonexistent requests threshold"], "anchor": "schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom--nonexistent_requests_threshold", "description": "The percentage of non-existent requests beyond which the system will flag this user as malicious.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_custom", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_non_existent_url_activity_custom", "nonexistent_requests_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_custom/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Non-existent URL Custom Activity Setting.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["app_settingCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/)
- [app_type_settings.user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom

<a id="section"></a>

Type: `"single"`. Computed.

Non-existent URL Custom Activity Setting.

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

## Direct properties

<a id="schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_non_existent_url_activity_custom--nonexistent_requests_threshold"></a>

### nonexistent_requests_threshold property

Type: `"number"`. Computed.

The percentage of non-existent requests beyond which the system will flag this user as malicious.

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
      "validatedAt": "2026-10-17T10:10:42+00:00"
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
