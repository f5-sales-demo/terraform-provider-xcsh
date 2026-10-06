---
page_title: "bot_protection_setting"
subcategory: "Security"
description: "Configuration of WAF Bot Protection."
xcsh_docs: {"aliases": ["bot protection setting"], "body_bytes": 2965, "body_sha256": "sha256:136743600b011c3b302030944247ca78403d6eb4d333836a1f97b2d09ad4f1bd", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:bot_protection_setting", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "documentation/data-sources/app_firewall/properties/bot_protection_setting/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2103022221032002-3031232012310300-2201020202303231-0102202313031033-2021133200010103-0330231303130020-0330211202222110-3021203231003220", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_protection_setting"], "schema_version": 1, "sections": [{"aliases": ["bot protection setting good bot action"], "anchor": "schema-bot_protection_setting--good_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:bot_protection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_protection_setting", "good_bot_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot protection setting malicious bot action"], "anchor": "schema-bot_protection_setting--malicious_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:bot_protection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_protection_setting", "malicious_bot_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot protection setting suspicious bot action"], "anchor": "schema-bot_protection_setting--suspicious_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:bot_protection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_protection_setting", "suspicious_bot_action"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/bot_protection_setting/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration of WAF Bot Protection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_protection_setting

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- bot_protection_setting

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: bot\_protection\_setting, default\_bot\_setting; Default: default\_bot\_setting\]
Configuration parameter for bot protection setting.

Additional upstream details:

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

- [bot_protection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/bot_protection_setting/#section)
- [default_bot_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_bot_setting/#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-bot_protection_setting--good_bot_action"></a>

### good_bot_action property

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

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

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

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

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

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
