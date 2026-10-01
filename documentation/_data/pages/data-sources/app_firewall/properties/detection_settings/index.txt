---
page_title: "detection_settings"
subcategory: "Security"
description: "detection_settings for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 6384, "body_sha256": "sha256:b8ee434599f371e491dcbad9bee3e8220508847d85962d507121b35b13583010", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:detection_settings:bot_protection_setting", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:default_bot_setting", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:default_violation_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_staging", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_suppression", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_threat_campaigns", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:enable_suppression", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:enable_threat_campaigns", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:stage_new_and_updated_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:stage_new_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violation_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view"], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "documentation/data-sources/app_firewall/properties/detection_settings/index.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["detection_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/detection_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "detection_settings for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- detection_settings

<a id="section"></a>

Type: `"single"`. Computed.

Specifies detection settings to be used by WAF.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-bot_protection_choice": "[\"bot_protection_setting\",\"default_bot_setting\"]",
  "x-ves-oneof-field-false_positive_suppression": "[\"disable_suppression\",\"enable_suppression\"]",
  "x-ves-oneof-field-signatures_staging_settings": "[\"disable_staging\",\"stage_new_and_updated_signatures\",\"stage_new_signatures\"]",
  "x-ves-oneof-field-threat_campaign_choice": "[\"disable_threat_campaigns\",\"enable_threat_campaigns\"]",
  "x-ves-oneof-field-violation_detection_setting": "[\"default_violation_settings\",\"violation_settings\"]"
}
```

## Direct properties

- [bot_protection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/bot_protection_setting/): complete subsection reference.

- [default_bot_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/default_bot_setting/): complete subsection reference.

- [default_violation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/default_violation_settings/): complete subsection reference.

- [disable_staging](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/disable_staging/): complete subsection reference.

- [disable_suppression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/disable_suppression/): complete subsection reference.

- [disable_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/disable_threat_campaigns/): complete subsection reference.

- [enable_suppression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/enable_suppression/): complete subsection reference.

- [enable_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/enable_threat_campaigns/): complete subsection reference.

- [signature_selection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/): complete subsection reference.

- [stage_new_and_updated_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/stage_new_and_updated_signatures/): complete subsection reference.

- [stage_new_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/stage_new_signatures/): complete subsection reference.

- [violation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violation_settings/): complete subsection reference.

- [violations_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violations_view/): complete subsection reference.

## Next pages

- [detection_settings.bot_protection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/bot_protection_setting/)
- [detection_settings.default_bot_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/default_bot_setting/)
- [detection_settings.default_violation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/default_violation_settings/)
- [detection_settings.disable_staging](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/disable_staging/)
- [detection_settings.disable_suppression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/disable_suppression/)
- [detection_settings.disable_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/disable_threat_campaigns/)
- [detection_settings.enable_suppression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/enable_suppression/)
- [detection_settings.enable_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/enable_threat_campaigns/)
- [detection_settings.signature_selection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/)
- [detection_settings.stage_new_and_updated_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/stage_new_and_updated_signatures/)
- [detection_settings.stage_new_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/stage_new_signatures/)
- [detection_settings.violation_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violation_settings/)
- [detection_settings.violations_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/violations_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
