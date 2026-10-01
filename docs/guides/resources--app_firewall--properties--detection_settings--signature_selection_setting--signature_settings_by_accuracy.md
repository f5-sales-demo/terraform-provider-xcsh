---
page_title: "detection_settings.signature_selection_setting.signature_settings_by_accuracy"
subcategory: "Security"
description: "detection_settings.signature_selection_setting.signature_settings_by_accuracy for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 4135, "body_sha256": "sha256:4fa13ec5bfcf699e0be431001042ae6acf43936f2daae315c63927b9f1d9f511", "canonical_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting:signature_settings_by_accuracy", "parent_id": "xcsh-docs:resources:app_firewall:properties:detection_settings:signature_selection_setting", "path": "docs/guides/resources--app_firewall--properties--detection_settings--signature_selection_setting--signature_settings_by_accuracy.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["detection_settings", "signature_selection_setting", "signature_settings_by_accuracy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/detection_settings/signature_selection_setting/signature_settings_by_accuracy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "detection_settings.signature_selection_setting.signature_settings_by_accuracy for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.signature_selection_setting.signature_settings_by_accuracy

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Property reference](resources--app_firewall--reference.md)
- [detection_settings](resources--app_firewall--properties--detection_settings.md)
- [detection_settings.signature_selection_setting](resources--app_firewall--properties--detection_settings--signature_selection_setting.md)
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

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
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

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
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

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
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

## Next pages

- [detection_settings.signature_selection_setting](resources--app_firewall--properties--detection_settings--signature_selection_setting.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
