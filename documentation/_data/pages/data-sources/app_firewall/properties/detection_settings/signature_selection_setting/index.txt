---
page_title: "detection_settings.signature_selection_setting"
subcategory: "Security"
description: "detection_settings.signature_selection_setting for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 4996, "body_sha256": "sha256:1bacae231034b1b701585dd37671f1c52dc8ae8fc824420d7b7864d351f405b8", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:attack_type_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:default_attack_type_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:default_signature_setting", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_accuracy_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_low_accuracy_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:only_high_accuracy_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy"], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings", "path": "documentation/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/index.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["detection_settings", "signature_selection_setting"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "detection_settings.signature_selection_setting for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.signature_selection_setting

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/)
- detection_settings.signature_selection_setting

<a id="section"></a>

Type: `"single"`. Computed.

Attack Signatures are patterns that identify attacks on a web application and its components.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-attack_type_setting": "[\"attack_type_settings\",\"default_attack_type_settings\"]",
  "x-ves-oneof-field-signature_protection_choice": "[\"default_signature_setting\",\"signature_settings_by_accuracy\"]",
  "x-ves-oneof-field-signature_selection_by_accuracy": "[\"high_medium_accuracy_signatures\",\"high_medium_low_accuracy_signatures\",\"only_high_accuracy_signatures\"]"
}
```

## Direct properties

- [attack_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/attack_type_settings/): complete subsection reference.

- [default_attack_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/default_attack_type_settings/): complete subsection reference.

- [default_signature_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/default_signature_setting/): complete subsection reference.

- [high_medium_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/high_medium_accuracy_signatures/): complete subsection reference.

- [high_medium_low_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/high_medium_low_accuracy_signatures/): complete subsection reference.

- [only_high_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/only_high_accuracy_signatures/): complete subsection reference.

- [signature_settings_by_accuracy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/): complete subsection reference.

## Next pages

- [detection_settings.signature_selection_setting.attack_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/attack_type_settings/)
- [detection_settings.signature_selection_setting.default_attack_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/default_attack_type_settings/)
- [detection_settings.signature_selection_setting.default_signature_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/default_signature_setting/)
- [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/high_medium_accuracy_signatures/)
- [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/high_medium_low_accuracy_signatures/)
- [detection_settings.signature_selection_setting.only_high_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/only_high_accuracy_signatures/)
- [detection_settings.signature_selection_setting.signature_settings_by_accuracy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
