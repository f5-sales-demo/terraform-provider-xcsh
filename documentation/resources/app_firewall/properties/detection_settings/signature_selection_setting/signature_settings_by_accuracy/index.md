---
page_title: "detection_settings.signature_selection_setting.signature_settings_by_accuracy"
subcategory: "Security"
description: "Configuration of WAF Signature Protection."
xcsh_docs: {"aliases": ["detection settings signature selection setting signature settings by accuracy"], "body_bytes": 4639, "body_sha256": "sha256:1464478e5fa4c095b36e2db9ab919f897d49f1cf6849a71c02f2a85a14d66828", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "parent_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting", "path": "documentation/resources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2133003013202031-1023113013011331-3213301103221203-2031333103131130-2012312120103120-3320310112023101-0300032210212201-3001100113000332", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings", "signature_selection_setting", "signature_settings_by_accuracy"], "schema_version": 1, "sections": [{"aliases": ["detection settings signature selection setting signature settings by accuracy high accuracy action"], "anchor": "schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--high_accuracy_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["SIG_BLOCK", "SIG_IGNORE", "SIG_REPORT"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "signature_settings_by_accuracy", "high_accuracy_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["detection settings signature selection setting signature settings by accuracy low accuracy action"], "anchor": "schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--low_accuracy_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["SIG_BLOCK", "SIG_IGNORE", "SIG_REPORT"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "signature_settings_by_accuracy", "low_accuracy_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["detection settings signature selection setting signature settings by accuracy medium accuracy action"], "anchor": "schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--medium_accuracy_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["SIG_BLOCK", "SIG_IGNORE", "SIG_REPORT"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "signature_settings_by_accuracy", "medium_accuracy_action"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Configuration of WAF Signature Protection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["app_firewallCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.signature_selection_setting.signature_settings_by_accuracy

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/)
- [detection_settings.signature_selection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/detection_settings/signature_selection_setting/)
- detection_settings.signature_selection_setting.signature_settings_by_accuracy

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
signature_settings_by_accuracy {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-detection_settings--signature_selection_setting--signature_settings_by_accuracy--high_accuracy_action"></a>

### high_accuracy_action property

Type: `"string"`. Optional.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SIG_BLOCK","SIG_IGNORE","SIG_REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"),
}
```

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

Type: `"string"`. Optional.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SIG_BLOCK","SIG_IGNORE","SIG_REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"),
}
```

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

Type: `"string"`. Optional.

\[Enum: SIG\_BLOCK|SIG\_REPORT|SIG\_IGNORE\] Action to be performed on the request Log and block Log
only Disable detection. Possible values are \`SIG\_BLOCK\`, \`SIG\_REPORT\`, \`SIG\_IGNORE\`.
Defaults to \`SIG\_BLOCK\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SIG_BLOCK","SIG_IGNORE","SIG_REPORT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SIG_BLOCK",
    "SIG_REPORT",
    "SIG_IGNORE"),
}
```

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
