---
page_title: "detection_settings"
subcategory: "Security"
description: "Specifies detection settings to be used by WAF."
xcsh_docs: {"aliases": ["detection settings"], "body_bytes": 3781, "body_sha256": "sha256:3d069575a25335491c640255f012984116247593e4b947c4637570f32e5fadbe", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:detection_settings:bot_protection_setting", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:default_bot_setting", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:default_violation_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_staging", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_suppression", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_threat_campaigns", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:enable_suppression", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:enable_threat_campaigns", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:stage_new_and_updated_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:stage_new_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violation_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "documentation/data-sources/app_firewall/properties/detection_settings/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings"], "schema_version": 1, "sections": [{"aliases": ["detection settings bot protection setting"], "anchor": "section", "description": "Configuration of WAF Bot Protection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:bot_protection_setting", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings", "bot_protection_setting"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings default bot setting"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:default_bot_setting", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "default_bot_setting"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings default violation settings"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:default_violation_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "default_violation_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings disable staging"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_staging", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "disable_staging"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings disable suppression"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_suppression", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "disable_suppression"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings disable threat campaigns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_threat_campaigns", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "disable_threat_campaigns"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings enable suppression"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:enable_suppression", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "enable_suppression"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings enable threat campaigns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:enable_threat_campaigns", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "enable_threat_campaigns"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings signature selection setting"], "anchor": "section", "description": "Attack Signatures are patterns that identify attacks on a web application and its components.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings stage new and updated signatures"], "anchor": "section", "description": "Attack Signatures staging configuration.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:stage_new_and_updated_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings", "stage_new_and_updated_signatures"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings stage new signatures"], "anchor": "section", "description": "Attack Signatures staging configuration.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:stage_new_signatures", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings", "stage_new_signatures"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings violation settings"], "anchor": "section", "description": "Specifies violation settings to be used by WAF.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violation_settings", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings", "violation_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["detection settings violations view"], "anchor": "section", "description": "List of violation checks that are performed on HTTP request to ensure the requests are properly formatted, detection of evasion techniques and other violations.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["detection_settings", "violations_view"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/detection_settings/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Specifies detection settings to be used by WAF.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["app_firewallCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
