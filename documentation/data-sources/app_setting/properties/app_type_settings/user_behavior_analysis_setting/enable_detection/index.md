---
page_title: "app_type_settings.user_behavior_analysis_setting.enable_detection"
subcategory: ""
description: "Various factors about user activity are monitored and analysed to determine malicious users. These settings allow tuning those factors used by the system to detect malicious users."
xcsh_docs: {"aliases": ["app type settings user behavior analysis setting enable detection"], "body_bytes": 13543, "body_sha256": "sha256:41a41644eee5085bef5b4d5358785c1372f4f867fe1fe5f1ba77d8177966664d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:bola_detection_automatic", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_bola_detection", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_bot_defense_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_failed_login_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_forbidden_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_ip_reputation", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_non_existent_url_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_rate_limit", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_waf_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_bot_defense_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_failed_login_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_forbidden_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_ip_reputation", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_automatic", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_custom", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_rate_limit", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_waf_activity"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "parent_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting", "path": "documentation/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/index.md", "product": "distributed-cloud", "provider_name": "app_setting", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0311323113202212-1302021001011002-3222320111030311-0013121322020222-2102202022120010-0203201211213102-3233013000302021-1313022231130301", "registry_path": "docs/guides/data-sources--app_setting--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection"], "schema_version": 1, "sections": [{"aliases": ["app type settings user behavior analysis setting enable detection bola detection automatic"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:bola_detection_automatic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "bola_detection_automatic"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection cooling off period"], "anchor": "schema-app_type_settings--user_behavior_analysis_setting--enable_detection--cooling_off_period", "description": "Exclusive with Malicious user detection assigns a threat level to each user based on their activity. Once a threat level is assigned, the system continues tracking activity from this user and if no further malicious activity is seen, it gradually reduces the threat assessment to lower levels. This field specifies the", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "cooling_off_period"], "syntax": "attribute", "type": "number"}, {"aliases": ["app type settings user behavior analysis setting enable detection exclude bola detection"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_bola_detection", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "exclude_bola_detection"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection exclude bot defense activity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_bot_defense_activity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "exclude_bot_defense_activity"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection exclude failed login activity", "login", "login result", "sign in"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_failed_login_activity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "exclude_failed_login_activity"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection exclude forbidden activity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_forbidden_activity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "exclude_forbidden_activity"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection exclude ip reputation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_ip_reputation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "exclude_ip_reputation"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection exclude non existent url activity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_non_existent_url_activity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "exclude_non_existent_url_activity"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection exclude rate limit"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_rate_limit", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "exclude_rate_limit"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection exclude waf activity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_waf_activity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "exclude_waf_activity"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection include bot defense activity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_bot_defense_activity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_bot_defense_activity"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection include failed login activity", "login", "login result", "sign in"], "anchor": "section", "description": "When enabled, the system monitors persistent failed login attempts from a user. A failed login is detected if a request results in a response code of 401. These settings specify how to use failed login activity to determine suspicious behavior.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_failed_login_activity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_failed_login_activity"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection include forbidden activity"], "anchor": "section", "description": "When L7 policy rules are set up to disallow certain types of requests, the system monitors persistent attempts from a user to send requests which result in policy denies. These settings specify how to use disallowed request activity from a user to determine suspicious behavior.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_forbidden_activity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_forbidden_activity"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection include ip reputation"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_ip_reputation", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_ip_reputation"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection include non existent url activity automatic"], "anchor": "section", "description": "Non-existent URL Automatic Activity Settings.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_automatic", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_non_existent_url_activity_automatic"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection include non existent url activity custom"], "anchor": "section", "description": "Non-existent URL Custom Activity Setting.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_custom", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_non_existent_url_activity_custom"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection include rate limit"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_rate_limit", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_rate_limit"], "syntax": "attribute", "type": "object"}, {"aliases": ["app type settings user behavior analysis setting enable detection include waf activity"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_waf_activity", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection", "include_waf_activity"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Various factors about user activity are monitored and analysed to determine malicious users. These settings allow tuning those factors used by the system to detect malicious users.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# app_type_settings.user_behavior_analysis_setting.enable_detection

Breadcrumbs:

- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/)
- [app_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/)
- [app_type_settings.user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/)
- app_type_settings.user_behavior_analysis_setting.enable_detection

<a id="section"></a>

Type: `"single"`. Computed.

Various factors about user activity are monitored and analysed to determine malicious users. These
settings allow tuning those factors used by the system to detect malicious users.

Upstream description:

Various factors about user activity are monitored and analysed to determine malicious users. These
settings allow tuning those factors used by the system to detect malicious users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-bola_activity_choice": "[\"bola_detection_automatic\",\"exclude_bola_detection\"]",
  "x-ves-oneof-field-bot_defense_activity_choice": "[\"exclude_bot_defense_activity\",\"include_bot_defense_activity\"]",
  "x-ves-oneof-field-cooling_off_period_setting": "[\"cooling_off_period\"]",
  "x-ves-oneof-field-failed_login_activity_choice": "[\"exclude_failed_login_activity\",\"include_failed_login_activity\"]",
  "x-ves-oneof-field-forbidden_activity_choice": "[\"exclude_forbidden_activity\",\"include_forbidden_activity\"]",
  "x-ves-oneof-field-ip_reputation_choice": "[\"exclude_ip_reputation\",\"include_ip_reputation\"]",
  "x-ves-oneof-field-non_existent_url_activity_choice": "[\"exclude_non_existent_url_activity\",\"include_non_existent_url_activity_automatic\",\"include_non_existent_url_activity_custom\"]",
  "x-ves-oneof-field-rate_limit_choice": "[\"exclude_rate_limit\",\"include_rate_limit\"]",
  "x-ves-oneof-field-waf_activity_choice": "[\"exclude_waf_activity\",\"include_waf_activity\"]"
}
```

## Direct properties

- [bola_detection_automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/bola_detection_automatic/): complete subsection reference.

<a id="schema-app_type_settings--user_behavior_analysis_setting--enable_detection--cooling_off_period"></a>

### cooling_off_period property

Type: `"number"`. Computed.

Exclusive with \[\] Malicious user detection assigns a threat level to each user based on their
activity. Once a threat level is assigned, the system continues tracking activity from this user and
if no further malicious activity is seen, it gradually reduces the threat assessment to lower
levels..

Upstream description:

Exclusive with \[\] Malicious user detection assigns a threat level to each user based on their
activity. Once a threat level is assigned, the system continues tracking activity from this user and
if no further malicious activity is seen, it gradually reduces the threat assessment to lower
levels. This field specifies the time period, in minutes, used by the system to decay a user's
threat level from a high to medium or medium to low or low to none.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 120,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 5
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "120"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "5",
    "ves.io.schema.rules.uint32.lte": "120"
  }
}
```

- [exclude_bola_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_bola_detection/): complete subsection reference.

- [exclude_bot_defense_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_bot_defense_activity/): complete subsection reference.

- [exclude_failed_login_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_failed_login_activity/): complete subsection reference.

- [exclude_forbidden_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_forbidden_activity/): complete subsection reference.

- [exclude_ip_reputation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_ip_reputation/): complete subsection reference.

- [exclude_non_existent_url_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_non_existent_url_activity/): complete subsection reference.

- [exclude_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_rate_limit/): complete subsection reference.

- [exclude_waf_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_waf_activity/): complete subsection reference.

- [include_bot_defense_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_bot_defense_activity/): complete subsection reference.

- [include_failed_login_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_failed_login_activity/): complete subsection reference.

- [include_forbidden_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_forbidden_activity/): complete subsection reference.

- [include_ip_reputation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_ip_reputation/): complete subsection reference.

- [include_non_existent_url_activity_automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_automatic/): complete subsection reference.

- [include_non_existent_url_activity_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_custom/): complete subsection reference.

- [include_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_rate_limit/): complete subsection reference.

- [include_waf_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_waf_activity/): complete subsection reference.

## Next pages

- [app_type_settings.user_behavior_analysis_setting.enable_detection.bola_detection_automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/bola_detection_automatic/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bola_detection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_bola_detection/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_bot_defense_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_bot_defense_activity/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_failed_login_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_failed_login_activity/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_forbidden_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_forbidden_activity/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_ip_reputation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_ip_reputation/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_non_existent_url_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_non_existent_url_activity/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_rate_limit/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.exclude_waf_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/exclude_waf_activity/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_bot_defense_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_bot_defense_activity/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_failed_login_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_failed_login_activity/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_forbidden_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_forbidden_activity/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_ip_reputation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_ip_reputation/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_automatic/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_non_existent_url_activity_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_non_existent_url_activity_custom/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_rate_limit/)
- [app_type_settings.user_behavior_analysis_setting.enable_detection.include_waf_activity](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/include_waf_activity/)
- [app_type_settings.user_behavior_analysis_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/)
- [xcsh_app_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_setting/)
