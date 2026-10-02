---
page_title: "bot_protection_setting"
subcategory: "Security"
description: "Configuration of WAF Bot Protection."
xcsh_docs: {"aliases": ["bot protection setting"], "body_bytes": 3513, "body_sha256": "sha256:29ce0e83473ada28978b7fc8446c6ddcec6de91e57249646072d60f7f18f616a", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:bot_protection_setting", "parent_id": "xcsh-docs:data-sources:app_firewall:reference", "path": "documentation/data-sources/app_firewall/properties/bot_protection_setting/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2103022221032002-3031232012310300-2201020202303231-0102202313031033-2021133200010103-0330231303130020-0330211202222110-3021203231003220", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_protection_setting"], "schema_version": 1, "sections": [{"aliases": ["good bot action"], "anchor": "schema-bot_protection_setting--good_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:bot_protection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_protection_setting", "good_bot_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["malicious bot action"], "anchor": "schema-bot_protection_setting--malicious_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:bot_protection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_protection_setting", "malicious_bot_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["suspicious bot action"], "anchor": "schema-bot_protection_setting--suspicious_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:bot_protection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_protection_setting", "suspicious_bot_action"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/bot_protection_setting/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration of WAF Bot Protection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [bot_protection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/bot_protection_setting/#section)
- [default_bot_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/default_bot_setting/#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-bot_protection_setting--good_bot_action"></a>

### good_bot_action property

Type: `"string"`. Computed.

\[Enum: BLOCK|REPORT|IGNORE\] Action to be performed on the request Log and block Log only Disable
detection. Possible values are \`BLOCK\`, \`REPORT\`, \`IGNORE\`. Defaults to \`BLOCK\`.

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

Upstream description:

Action to be performed on the request

Log and block Log only Disable detection.

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
