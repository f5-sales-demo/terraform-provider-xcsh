---
page_title: "app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity"
subcategory: ""
description: "When enabled, the system monitors persistent failed login attempts from a user. A failed login is detected if a request results in a response code of 401. These settings specify how to use failed login activity to determine suspicious behavior."
xcsh_docs: {"aliases": ["app type settings user behavior analysis setting enable detection include failed login activity", "login", "login result", "sign in"], "body_bytes": 3196, "body_sha256": "sha256:3efe70090a59a0c5affaf743fb0c501c4bba84cb964d57ca141a18f7f0049d92", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_failed_login_activity", "parent_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "path": "documentation/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_failed_login_activity/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0003002322222103-1121013210122121-3021222013130321-2210120200003203-2212132111113130-3231021322323121-0032103221220112-3021132011121013", "registry_path": "docs/guides/data-sources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_failed_login_activity"], "schema_version": 1, "sections": [{"aliases": ["app type settings user behavior analysis setting enable detection include failed login activity login failures threshold", "login", "login result", "sign in"], "anchor": "schema-app_type_settings--user_behavior_analysis_setting--enable_detection--include_failed_login_activity--login_failures_threshold", "description": "The number of failed logins beyond which the system will flag this user as malicious.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_failed_login_activity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_failed_login_activity", "login_failures_threshold"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_failed_login_activity/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "When enabled, the system monitors persistent failed login attempts from a user. A failed login is detected if a request results in a response code of 401. These settings specify how to use failed login activity to determine suspicious behavior.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

## Next pages

- [app_type_settings.user_behavior_analysis_setting.enable_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
