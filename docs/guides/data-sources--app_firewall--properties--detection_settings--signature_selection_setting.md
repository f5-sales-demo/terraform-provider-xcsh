---
page_title: "detection_settings.signature_selection_setting"
subcategory: "Security"
description: "detection_settings.signature_selection_setting for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 4053, "body_sha256": "sha256:f1127c6ca518bb5fc505848aaff06a2b5971ec7ccf157292193ea785dc2fd410", "canonical_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:attack_type_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:default_attack_type_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:default_signature_setting", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_accuracy_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_low_accuracy_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:only_high_accuracy_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy"], "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings", "path": "docs/guides/data-sources--app_firewall--properties--detection_settings--signature_selection_setting.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["detection_settings", "signature_selection_setting"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "detection_settings.signature_selection_setting for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.signature_selection_setting

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md)
- [Property reference](data-sources--app_firewall--reference.md)
- [detection_settings](data-sources--app_firewall--properties--detection_settings.md)
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

- [attack_type_settings](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--attack_type_settings.md): complete subsection reference.

- [default_attack_type_settings](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--default_attack_type_settings.md): complete subsection reference.

- [default_signature_setting](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--default_signature_setting.md): complete subsection reference.

- [high_medium_accuracy_signatures](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--high_medium_accuracy_signatures.md): complete subsection reference.

- [high_medium_low_accuracy_signatures](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--high_medium_low_accuracy_signatures.md): complete subsection reference.

- [only_high_accuracy_signatures](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--only_high_accuracy_signatures.md): complete subsection reference.

- [signature_settings_by_accuracy](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--signature_settings_by_accuracy.md): complete subsection reference.

## Next pages

- [detection_settings.signature_selection_setting.attack_type_settings](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--attack_type_settings.md)
- [detection_settings.signature_selection_setting.default_attack_type_settings](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--default_attack_type_settings.md)
- [detection_settings.signature_selection_setting.default_signature_setting](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--default_signature_setting.md)
- [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--high_medium_accuracy_signatures.md)
- [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--high_medium_low_accuracy_signatures.md)
- [detection_settings.signature_selection_setting.only_high_accuracy_signatures](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--only_high_accuracy_signatures.md)
- [detection_settings.signature_selection_setting.signature_settings_by_accuracy](data-sources--app_firewall--properties--detection_settings--signature_selection_setting--signature_settings_by_accuracy.md)
- [detection_settings](data-sources--app_firewall--properties--detection_settings.md)
- [xcsh_app_firewall](../data-sources/app_firewall.md)
