---
page_title: "detection_settings.signature_selection_setting"
subcategory: "Security"
description: "detection_settings.signature_selection_setting for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 4791, "body_sha256": "sha256:d24dd1b581b849e1af01d653ad057a973dd1affb27e8e1279ce8ea9c9a9e0cd0", "canonical_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting", "child_ids": ["xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:attack_type_settings", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:default_attack_type_settings", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:default_signature_setting", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_accuracy_signatures", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_low_accuracy_signatures", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:only_high_accuracy_signatures", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy"], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting", "parent_id": "xcsh-docs:resources:app_firewall:properties:detection_settings", "path": "docs/guides/resources--app_firewall--properties--detection_settings--signature_selection_setting.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["detection_settings", "signature_selection_setting"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/detection_settings/signature_selection_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "detection_settings.signature_selection_setting for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.signature_selection_setting

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Property reference](resources--app_firewall--reference.md)
- [detection_settings](resources--app_firewall--properties--detection_settings.md)
- detection_settings.signature_selection_setting

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Attack Signatures are patterns that identify attacks on a web application and its components.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("attack_type_settings",
    "default_attack_type_settings"),
  validators.ConflictingObjectAttributes("default_signature_setting",
    "signature_settings_by_accuracy"),
  validators.ConflictingObjectAttributes("high_medium_accuracy_signatures",
    "high_medium_low_accuracy_signatures"),
  validators.ConflictingObjectAttributes("high_medium_accuracy_signatures",
    "only_high_accuracy_signatures"),
  validators.ConflictingObjectAttributes("high_medium_low_accuracy_signatures",
    "only_high_accuracy_signatures")}
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
  "x-ves-oneof-field-attack_type_setting": "[\"attack_type_settings\",\"default_attack_type_settings\"]",
  "x-ves-oneof-field-signature_protection_choice": "[\"default_signature_setting\",\"signature_settings_by_accuracy\"]",
  "x-ves-oneof-field-signature_selection_by_accuracy": "[\"high_medium_accuracy_signatures\",\"high_medium_low_accuracy_signatures\",\"only_high_accuracy_signatures\"]"
}
```

Terraform syntax:

```terraform
signature_selection_setting {
  # Configure direct properties listed below.
}
```

## Direct properties

- [attack_type_settings](resources--app_firewall--properties--detection_settings--signature_selection_setting--attack_type_settings.md): complete subsection reference.

- [default_attack_type_settings](resources--app_firewall--properties--detection_settings--signature_selection_setting--default_attack_type_settings.md): complete subsection reference.

- [default_signature_setting](resources--app_firewall--properties--detection_settings--signature_selection_setting--default_signature_setting.md): complete subsection reference.

- [high_medium_accuracy_signatures](resources--app_firewall--properties--detection_settings--signature_selection_setting--high_medium_accuracy_signatures.md): complete subsection reference.

- [high_medium_low_accuracy_signatures](resources--app_firewall--properties--detection_settings--signature_selection_setting--high_medium_low_accuracy_signatures.md): complete subsection reference.

- [only_high_accuracy_signatures](resources--app_firewall--properties--detection_settings--signature_selection_setting--only_high_accuracy_signatures.md): complete subsection reference.

- [signature_settings_by_accuracy](resources--app_firewall--properties--detection_settings--signature_selection_setting--signature_settings_by_accuracy.md): complete subsection reference.

## Next pages

- [detection_settings.signature_selection_setting.attack_type_settings](resources--app_firewall--properties--detection_settings--signature_selection_setting--attack_type_settings.md)
- [detection_settings.signature_selection_setting.default_attack_type_settings](resources--app_firewall--properties--detection_settings--signature_selection_setting--default_attack_type_settings.md)
- [detection_settings.signature_selection_setting.default_signature_setting](resources--app_firewall--properties--detection_settings--signature_selection_setting--default_signature_setting.md)
- [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](resources--app_firewall--properties--detection_settings--signature_selection_setting--high_medium_accuracy_signatures.md)
- [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](resources--app_firewall--properties--detection_settings--signature_selection_setting--high_medium_low_accuracy_signatures.md)
- [detection_settings.signature_selection_setting.only_high_accuracy_signatures](resources--app_firewall--properties--detection_settings--signature_selection_setting--only_high_accuracy_signatures.md)
- [detection_settings.signature_selection_setting.signature_settings_by_accuracy](resources--app_firewall--properties--detection_settings--signature_selection_setting--signature_settings_by_accuracy.md)
- [detection_settings](resources--app_firewall--properties--detection_settings.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
