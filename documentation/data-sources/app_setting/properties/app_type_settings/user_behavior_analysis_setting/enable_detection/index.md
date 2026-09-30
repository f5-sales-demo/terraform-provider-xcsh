---
page_title: "app_type_settings.user_behavior_analysis_setting.enable_detection"
subcategory: ""
description: "app_type_settings.user_behavior_analysis_setting.enable_detection for xcsh_app_setting."
xcsh_docs: {"aliases": [], "body_bytes": 13444, "body_sha256": "sha256:091a515f4324dd1ba89da42f683fec37518a3c95669d67eae2858a6b80a5e564", "child_ids": ["xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:bola_detection_automatic", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_bola_detection", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_bot_defense_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_failed_login_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_forbidden_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_ip_reputation", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_non_existent_url_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_rate_limit", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:exclude_waf_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_bot_defense_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_failed_login_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_forbidden_activity", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_ip_reputation", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_automatic", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_non_existent_url_activity_custom", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_rate_limit", "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection:include_waf_activity"], "collection_id": "xcsh-docs:data-sources:app_setting:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting:enable_detection", "parent_id": "xcsh-docs:data-sources:app_setting:properties:app_type_settings:user_behavior_analysis_setting", "path": "documentation/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/index.md", "provider_name": "app_setting", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["app_type_settings", "user_behavior_analysis_setting", "enable_detection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_setting/properties/app_type_settings/user_behavior_analysis_setting/enable_detection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "app_type_settings.user_behavior_analysis_setting.enable_detection for xcsh_app_setting.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_settingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
