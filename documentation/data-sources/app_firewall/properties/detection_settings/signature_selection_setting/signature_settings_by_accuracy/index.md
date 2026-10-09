---
page_title: "detection_settings.signature_selection_setting.signature_settings_by_accuracy"
subcategory: "Security"
description: "Configuration of WAF Signature Protection."
xcsh_docs: {"aliases": ["detection settings signature selection setting signature settings by accuracy"], "body_bytes": 3165, "body_sha256": "sha256:821dd2ce5640cb537fc4f5ae7212b38409c52d7e771aa8fd5da91c83ab7c5012", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting", "path": "documentation/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-0223310013322210-0201120110122302-3210110002210201-3320130230112032-3011301330100222-2110113001223200-0133333012023221-2021202010331032", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings", "signature_selection_setting", "signature_settings_by_accuracy"], "schema_version": 1, "sections": [{"aliases": ["detection settings signature selection setting signature settings by accuracy high accuracy action"], "anchor": "schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--high_accuracy_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "signature_settings_by_accuracy", "high_accuracy_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["detection settings signature selection setting signature settings by accuracy low accuracy action"], "anchor": "schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--low_accuracy_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "signature_settings_by_accuracy", "low_accuracy_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["detection settings signature selection setting signature settings by accuracy medium accuracy action"], "anchor": "schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--medium_accuracy_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "signature_settings_by_accuracy", "medium_accuracy_action"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Configuration of WAF Signature Protection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["app_firewallCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.signature_selection_setting.signature_settings_by_accuracy

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/)
- [detection_settings.signature_selection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/)
- detection_settings.signature_selection_setting.signature_settings_by_accuracy

<a id="section"></a>

Type: `"single"`. Computed.

Configuration of WAF Signature Protection.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

<a id="schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--high_accuracy_action"></a>

### high_accuracy_action property

Type: `"string"`. Computed.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SIG_BLOCK",
  "enum": [
    "SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--low_accuracy_action"></a>

### low_accuracy_action property

Type: `"string"`. Computed.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SIG_BLOCK",
  "enum": [
    "SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--medium_accuracy_action"></a>

### medium_accuracy_action property

Type: `"string"`. Computed.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SIG_BLOCK",
  "enum": [
    "SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
