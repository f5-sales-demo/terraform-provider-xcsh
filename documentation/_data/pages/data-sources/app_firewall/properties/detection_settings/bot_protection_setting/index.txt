---
page_title: "detection_settings.bot_protection_setting"
subcategory: "Security"
description: "Configuration of WAF Bot Protection."
xcsh_docs: {"aliases": ["detection settings bot protection setting"], "body_bytes": 3270, "body_sha256": "sha256:7e94b9436560634c8632425e39a8dd92e47598eac2694019292e2293ccb10cd2", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:bot_protection_setting", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings", "path": "documentation/data-sources/app_firewall/properties/detection_settings/bot_protection_setting/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3120233020110133-0002320311033313-1300011030333303-2323113212300203-1223031312021000-3112223031220011-0122021023321331-0131222123323023", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings", "bot_protection_setting"], "schema_version": 1, "sections": [{"aliases": ["detection settings bot protection setting good bot action"], "anchor": "schema-detection_settings--bot_protection_setting--good_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:bot_protection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "bot_protection_setting", "good_bot_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["detection settings bot protection setting malicious bot action"], "anchor": "schema-detection_settings--bot_protection_setting--malicious_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:bot_protection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "bot_protection_setting", "malicious_bot_action"], "syntax": "attribute", "type": "string"}, {"aliases": ["detection settings bot protection setting suspicious bot action"], "anchor": "schema-detection_settings--bot_protection_setting--suspicious_bot_action", "description": "Action to be performed on the request Log and block Log only Disable detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:bot_protection_setting", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "bot_protection_setting", "suspicious_bot_action"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/detection_settings/bot_protection_setting/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration of WAF Bot Protection.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.bot_protection_setting

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/)
- detection_settings.bot_protection_setting

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

<a id="schema-detection_settings--bot_protection_setting--good_bot_action"></a>

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

<a id="schema-detection_settings--bot_protection_setting--malicious_bot_action"></a>

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

<a id="schema-detection_settings--bot_protection_setting--suspicious_bot_action"></a>

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

- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
