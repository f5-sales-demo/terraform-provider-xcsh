---
page_title: "detection_settings"
subcategory: "Security"
description: "detection_settings for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 5591, "body_sha256": "sha256:b78d116ab575fef8ced598d9f85b2c5fed97683825cf1bb30a216a34512d28d0", "canonical_id": "xcsh-docs:resources:app_firewall:properties:detection_settings", "child_ids": ["xcsh-docs:resources:app_firewall:properties:detection_settings:bot_protection_setting", "xcsh-docs:resources:app_firewall:properties:detection_settings:default_bot_setting", "xcsh-docs:resources:app_firewall:properties:detection_settings:default_violation_settings", "xcsh-docs:resources:app_firewall:properties:detection_settings:disable_staging", "xcsh-docs:resources:app_firewall:properties:detection_settings:disable_suppression", "xcsh-docs:resources:app_firewall:properties:detection_settings:disable_threat_campaigns", "xcsh-docs:resources:app_firewall:properties:detection_settings:enable_suppression", "xcsh-docs:resources:app_firewall:properties:detection_settings:enable_threat_campaigns", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting", "xcsh-docs:resources:app_firewall:properties:detection_settings:stage_new_and_updated_signatures", "xcsh-docs:resources:app_firewall:properties:detection_settings:stage_new_signatures", "xcsh-docs:resources:app_firewall:properties:detection_settings:violation_settings", "xcsh-docs:resources:app_firewall:properties:detection_settings:violations_view"], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:detection_settings", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "docs/guides/resources--app_firewall--properties--detection_settings.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["detection_settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/detection_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "detection_settings for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# detection_settings

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Property reference](resources--app_firewall--reference.md)
- detection_settings

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies detection settings to be used by WAF.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("bot_protection_setting",
    "default_bot_setting"),
  validators.ConflictingObjectAttributes("default_violation_settings",
    "violation_settings"),
  validators.ConflictingObjectAttributes("disable_staging",
    "stage_new_and_updated_signatures"),
  validators.ConflictingObjectAttributes("disable_staging",
    "stage_new_signatures"),
  validators.ConflictingObjectAttributes("disable_suppression",
    "enable_suppression"),
  validators.ConflictingObjectAttributes("disable_threat_campaigns",
    "enable_threat_campaigns"),
  validators.ConflictingObjectAttributes("stage_new_and_updated_signatures",
    "stage_new_signatures")}
```

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

Terraform syntax:

```terraform
detection_settings {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bot_protection_setting](resources--app_firewall--properties--detection_settings--bot_protection_setting.md): complete subsection reference.

- [default_bot_setting](resources--app_firewall--properties--detection_settings--default_bot_setting.md): complete subsection reference.

- [default_violation_settings](resources--app_firewall--properties--detection_settings--default_violation_settings.md): complete subsection reference.

- [disable_staging](resources--app_firewall--properties--detection_settings--disable_staging.md): complete subsection reference.

- [disable_suppression](resources--app_firewall--properties--detection_settings--disable_suppression.md): complete subsection reference.

- [disable_threat_campaigns](resources--app_firewall--properties--detection_settings--disable_threat_campaigns.md): complete subsection reference.

- [enable_suppression](resources--app_firewall--properties--detection_settings--enable_suppression.md): complete subsection reference.

- [enable_threat_campaigns](resources--app_firewall--properties--detection_settings--enable_threat_campaigns.md): complete subsection reference.

- [signature_selection_setting](resources--app_firewall--properties--detection_settings--signature_selection_setting.md): complete subsection reference.

- [stage_new_and_updated_signatures](resources--app_firewall--properties--detection_settings--stage_new_and_updated_signatures.md): complete subsection reference.

- [stage_new_signatures](resources--app_firewall--properties--detection_settings--stage_new_signatures.md): complete subsection reference.

- [violation_settings](resources--app_firewall--properties--detection_settings--violation_settings.md): complete subsection reference.

- [violations_view](resources--app_firewall--properties--detection_settings--violations_view.md): complete subsection reference.

## Next pages

- [detection_settings.bot_protection_setting](resources--app_firewall--properties--detection_settings--bot_protection_setting.md)
- [detection_settings.default_bot_setting](resources--app_firewall--properties--detection_settings--default_bot_setting.md)
- [detection_settings.default_violation_settings](resources--app_firewall--properties--detection_settings--default_violation_settings.md)
- [detection_settings.disable_staging](resources--app_firewall--properties--detection_settings--disable_staging.md)
- [detection_settings.disable_suppression](resources--app_firewall--properties--detection_settings--disable_suppression.md)
- [detection_settings.disable_threat_campaigns](resources--app_firewall--properties--detection_settings--disable_threat_campaigns.md)
- [detection_settings.enable_suppression](resources--app_firewall--properties--detection_settings--enable_suppression.md)
- [detection_settings.enable_threat_campaigns](resources--app_firewall--properties--detection_settings--enable_threat_campaigns.md)
- [detection_settings.signature_selection_setting](resources--app_firewall--properties--detection_settings--signature_selection_setting.md)
- [detection_settings.stage_new_and_updated_signatures](resources--app_firewall--properties--detection_settings--stage_new_and_updated_signatures.md)
- [detection_settings.stage_new_signatures](resources--app_firewall--properties--detection_settings--stage_new_signatures.md)
- [detection_settings.violation_settings](resources--app_firewall--properties--detection_settings--violation_settings.md)
- [detection_settings.violations_view](resources--app_firewall--properties--detection_settings--violations_view.md)
- [Property reference](resources--app_firewall--reference.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
