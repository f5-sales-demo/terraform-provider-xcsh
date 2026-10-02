---
page_title: "detection_settings.signature_selection_setting"
subcategory: "Security"
description: "Attack Signatures are patterns that identify attacks on a web application and its components."
xcsh_docs: {"aliases": ["detection settings signature selection setting"], "body_bytes": 5734, "body_sha256": "sha256:6e40ac118a9cfc0863a0c040699b2bc1bf32bb4923f2e472d6553c2b11cdcda0", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:attack_type_settings", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:default_attack_type_settings", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:default_signature_setting", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_accuracy_signatures", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_low_accuracy_signatures", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:only_high_accuracy_signatures", "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting", "parent_id": "xcsh-docs:resources:app_firewall:properties:detection_settings", "path": "documentation/resources/app_firewall/properties/detection_settings/signature_selection_setting/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0301033122212120-1310121301130132-3030213123000103-2201001100122303-3010310331101032-2312032023010303-0333313103012033-0310010211320303", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "detection_settings.signature_selection_setting:ConflictingObjectAttributes:attack_type_settings,default_attack_type_settings", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:attack_type_settings", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "detection_settings.signature_selection_setting:ConflictingObjectAttributes:attack_type_settings,default_attack_type_settings", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:default_attack_type_settings", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "detection_settings.signature_selection_setting:ConflictingObjectAttributes:default_signature_setting,signature_settings_by_accuracy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:default_signature_setting", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "detection_settings.signature_selection_setting:ConflictingObjectAttributes:high_medium_accuracy_signatures,high_medium_low_accuracy_signatures", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_accuracy_signatures", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "detection_settings.signature_selection_setting:ConflictingObjectAttributes:high_medium_accuracy_signatures,only_high_accuracy_signatures", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_accuracy_signatures", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "detection_settings.signature_selection_setting:ConflictingObjectAttributes:high_medium_accuracy_signatures,high_medium_low_accuracy_signatures", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_low_accuracy_signatures", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "detection_settings.signature_selection_setting:ConflictingObjectAttributes:high_medium_low_accuracy_signatures,only_high_accuracy_signatures", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_low_accuracy_signatures", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "detection_settings.signature_selection_setting:ConflictingObjectAttributes:high_medium_accuracy_signatures,only_high_accuracy_signatures", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:only_high_accuracy_signatures", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "detection_settings.signature_selection_setting:ConflictingObjectAttributes:high_medium_low_accuracy_signatures,only_high_accuracy_signatures", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:only_high_accuracy_signatures", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "detection_settings.signature_selection_setting:ConflictingObjectAttributes:default_signature_setting,signature_settings_by_accuracy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings", "signature_selection_setting"], "schema_version": 1, "sections": [{"aliases": ["attack type settings"], "anchor": "section", "description": "Specifies attack-type settings to be used by WAF.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:attack_type_settings", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-detection_settings--signature_selection_setting--attack_type_settings--disabled_attack_types", "enforcement": "provider-schema", "group": "detection_settings.signature_selection_setting.attack_type_settings:RequiredObjectAttributes:disabled_attack_types", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:attack_type_settings", "type": "requires"}], "schema_path": ["detection_settings", "signature_selection_setting", "attack_type_settings"], "syntax": "block", "type": "object"}, {"aliases": ["default attack type settings"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:default_attack_type_settings", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "default_attack_type_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["default signature setting"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:default_signature_setting", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "default_signature_setting"], "syntax": "attribute", "type": "object"}, {"aliases": ["high medium accuracy signatures"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_accuracy_signatures", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "high_medium_accuracy_signatures"], "syntax": "attribute", "type": "object"}, {"aliases": ["high medium low accuracy signatures"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:high_medium_low_accuracy_signatures", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "high_medium_low_accuracy_signatures"], "syntax": "attribute", "type": "object"}, {"aliases": ["only high accuracy signatures"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:only_high_accuracy_signatures", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "only_high_accuracy_signatures"], "syntax": "attribute", "type": "object"}, {"aliases": ["signature settings by accuracy"], "anchor": "section", "description": "Configuration of WAF Signature Protection.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "signature_settings_by_accuracy"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/detection_settings/signature_selection_setting/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Attack Signatures are patterns that identify attacks on a web application and its components.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.signature_selection_setting

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/)
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

- [attack_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/attack_type_settings/): complete subsection reference.

- [default_attack_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/default_attack_type_settings/): complete subsection reference.

- [default_signature_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/default_signature_setting/): complete subsection reference.

- [high_medium_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/high_medium_accuracy_signatures/): complete subsection reference.

- [high_medium_low_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/high_medium_low_accuracy_signatures/): complete subsection reference.

- [only_high_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/only_high_accuracy_signatures/): complete subsection reference.

- [signature_settings_by_accuracy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/): complete subsection reference.

## Next pages

- [detection_settings.signature_selection_setting.attack_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/attack_type_settings/)
- [detection_settings.signature_selection_setting.default_attack_type_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/default_attack_type_settings/)
- [detection_settings.signature_selection_setting.default_signature_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/default_signature_setting/)
- [detection_settings.signature_selection_setting.high_medium_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/high_medium_accuracy_signatures/)
- [detection_settings.signature_selection_setting.high_medium_low_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/high_medium_low_accuracy_signatures/)
- [detection_settings.signature_selection_setting.only_high_accuracy_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/only_high_accuracy_signatures/)
- [detection_settings.signature_selection_setting.signature_settings_by_accuracy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
