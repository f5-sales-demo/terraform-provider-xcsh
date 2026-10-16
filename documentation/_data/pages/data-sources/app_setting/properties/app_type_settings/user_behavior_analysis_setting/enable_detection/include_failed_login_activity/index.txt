---
page_title: "app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity"
subcategory: ""
description: "When enabled, the system monitors persistent failed login attempts from a user. A failed login is detected if a request results in a response code of 401. These settings specify how to use failed login activity to determine suspicious behavior."
xcsh_docs: {"aliases": ["app type settings user behavior analysis setting enable detection include failed login activity", "login", "login result", "sign in"], "body_bytes": 2577, "body_sha256": "sha256:9163d1ff5093321f1859bb395be59387ebacd063aa3dab3f1550d07077283172", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_failed_login_activity", "parent_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "path": "documentation/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_failed_login_activity/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-0003002322222103-1121013210122121-3021222013130321-2210120200003203-2212132111113130-3231021322323121-0032103221220112-3021132011121013", "registry_path": "docs/guides/data-sources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_failed_login_activity"], "schema_version": 1, "sections": [{"aliases": ["app type settings user behavior analysis setting enable detection include failed login activity login failures threshold", "login", "login result", "sign in"], "anchor": "schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_failed_login_activity--login_failures_threshold", "description": "The number of failed logins beyond which the system will flag this user as malicious.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_failed_login_activity", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_failed_login_activity", "login_failures_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_failed_login_activity/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "When enabled, the system monitors persistent failed login attempts from a user. A failed login is detected if a request results in a response code of 401. These settings specify how to use failed login activity to determine suspicious behavior.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["app_settingCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/)
- [app_type_settings.user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/)
- app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity

<a id="section"></a>

Type: `"single"`. Computed.

When enabled, the system monitors persistent failed login attempts from a user. A failed login is
detected if a request results in a response code of 401. These settings specify how to use failed
login activity to determine suspicious behavior.

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

<a id="schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_failed_login_activity--login_failures_threshold"></a>

### login_failures_threshold property

Type: `"number"`. Computed.

The number of failed logins beyond which the system will flag this user as malicious.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-15T10:49:30+00:00"
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
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```
