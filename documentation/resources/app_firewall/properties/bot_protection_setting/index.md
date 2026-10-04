---
page_title: "bot_protection_setting"
subcategory: "Security"
description: "Configuration of WAF Bot Protection."
xcsh_docs: {"aliases": ["bot protection setting"], "body_bytes": 4115, "body_sha256": "sha256:baea4f6251aff02fdd3250f08b4c73fbde39a481b52021c321efbd21ec88494d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:bot_protection_setting", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "documentation/resources/app_firewall/properties/bot_protection_setting/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3311120131133012-0031000313131223-3122333020233031-0133330330300232-0200300331201303-3223220213230013-1111003221331133-0311210320311111", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_protection_setting"], "schema_version": 1, "sections": [{"aliases": ["bot protection setting good bot action"], "anchor": "schema-bot_protection_setting--good_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:resources:app_firewall:properties:bot_protection_setting", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_protection_setting", "good_bot_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot protection setting malicious bot action"], "anchor": "schema-bot_protection_setting--malicious_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:resources:app_firewall:properties:bot_protection_setting", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_protection_setting", "malicious_bot_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot protection setting suspicious bot action"], "anchor": "schema-bot_protection_setting--suspicious_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:resources:app_firewall:properties:bot_protection_setting", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_protection_setting", "suspicious_bot_action"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/bot_protection_setting/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Configuration of WAF Bot Protection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_protection_setting

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- bot_protection_setting

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bot\_protection\_setting, default\_bot\_setting; Default: default\_bot\_setting\]
Configuration parameter for bot protection setting.

Upstream description:

Configuration of WAF Bot Protection.

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

OneOf alternatives in this subsection:

- [bot_protection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/bot_protection_setting/#section)
- [default_bot_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/default_bot_setting/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bot_protection_setting {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-bot_protection_setting--good_bot_action"></a>

### good_bot_action property

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-bot_protection_setting--malicious_bot_action"></a>

### malicious_bot_action property

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-bot_protection_setting--suspicious_bot_action"></a>

### suspicious_bot_action property

Type: `"string"`. Optional.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("BLOCK",
    "REPORT",
    "IGNORE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "BLOCK",
  "enum": [
    "BLOCK",
    "REPORT",
    "IGNORE"
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
