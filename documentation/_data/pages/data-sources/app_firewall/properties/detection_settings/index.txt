---
page_title: "detection_settings"
subcategory: "Security"
description: "Specifies detection settings to be used by WAF."
xcsh_docs: {"aliases": ["detection settings"], "body_bytes": 6384, "body_sha256": "sha256:b8ee434599f371e491dcbad9bee3e8220508847d85962d507121b35b13583010", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:app_firewall:properties:detection_settings:bot_protection_setting", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:default_bot_setting", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:default_violation_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_staging", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_suppression", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_threat_campaigns", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:enable_suppression", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:enable_threat_campaigns", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:stage_new_and_updated_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:stage_new_signatures", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violation_settings", "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "documentation/data-sources/app_firewall/properties/detection_settings/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2333231122133331-1021202212013303-1200133000300303-2333310030131033-0111221320020202-1121133212011000-2301130211232010-2330121103032210", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings"], "schema_version": 1, "sections": [{"aliases": ["bot protection setting"], "anchor": "section", "description": "Configuration of WAF Bot Protection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:bot_protection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings", "bot_protection_setting"], "syntax": "attribute", "type": "object"}, {"aliases": ["default bot setting"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:default_bot_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "default_bot_setting"], "syntax": "attribute", "type": "object"}, {"aliases": ["default violation settings"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:default_violation_settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "default_violation_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable staging"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_staging", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "disable_staging"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable suppression"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_suppression", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "disable_suppression"], "syntax": "attribute", "type": "object"}, {"aliases": ["disable threat campaigns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:disable_threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "disable_threat_campaigns"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable suppression"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:enable_suppression", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "enable_suppression"], "syntax": "attribute", "type": "object"}, {"aliases": ["enable threat campaigns"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:enable_threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "enable_threat_campaigns"], "syntax": "attribute", "type": "object"}, {"aliases": ["signature selection setting"], "anchor": "section", "description": "Attack Signatures are patterns that identify attacks on a web application and its components.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting"], "syntax": "attribute", "type": "object"}, {"aliases": ["stage new and updated signatures"], "anchor": "section", "description": "Attack Signatures staging configuration.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:stage_new_and_updated_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings", "stage_new_and_updated_signatures"], "syntax": "attribute", "type": "object"}, {"aliases": ["stage new signatures"], "anchor": "section", "description": "Attack Signatures staging configuration.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:stage_new_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings", "stage_new_signatures"], "syntax": "attribute", "type": "object"}, {"aliases": ["violation settings"], "anchor": "section", "description": "Specifies violation settings to be used by WAF.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violation_settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["detection_settings", "violation_settings"], "syntax": "attribute", "type": "object"}, {"aliases": ["violations view"], "anchor": "section", "description": "List of violation checks that are performed on HTTP request to ensure the requests are properly formatted, detection of evasion techniques and other violations.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:violations_view", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["detection_settings", "violations_view"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/detection_settings/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Specifies detection settings to be used by WAF.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
