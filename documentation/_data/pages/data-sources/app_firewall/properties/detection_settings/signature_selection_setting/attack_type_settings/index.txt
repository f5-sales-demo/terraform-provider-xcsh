---
page_title: "detection_settings.signature_selection_setting.attack_type_settings"
subcategory: "Security"
description: "Specifies attack-type settings to be used by WAF."
xcsh_docs: {"aliases": ["detection settings signature selection setting attack type settings"], "body_bytes": 4759, "body_sha256": "sha256:ed45907cfe79cc2e12a3356882089e035d13c1208ae2073aefe6e6bf031bb12d", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:attack_type_settings", "parent_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting", "path": "documentation/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/attack_type_settings/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3221203122131311-3100111230001311-1210213123303233-0310222213330333-1311232002332113-0223211313203110-3113303022120031-3303310001132023", "registry_path": "docs/guides/data-sources--app_firewall--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["detection_settings", "signature_selection_setting", "attack_type_settings"], "schema_version": 1, "sections": [{"aliases": ["disabled attack types"], "anchor": "schema-detection_settings--signature_selection_setting--attack_type_settings--disabled_attack_types", "description": "List of Attack Types that will be ignored and not trigger a detection.", "document_id": "xcsh-docs:data-sources:app_firewall:properties:detection_settings:signature_selection_setting:attack_type_settings", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["detection_settings", "signature_selection_setting", "attack_type_settings", "disabled_attack_types"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/attack_type_settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specifies attack-type settings to be used by WAF.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# detection_settings.signature_selection_setting.attack_type_settings

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/)
- [detection_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/)
- [detection_settings.signature_selection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/)
- detection_settings.signature_selection_setting.attack_type_settings

<a id="section"></a>

Type: `"single"`. Computed.

Specifies attack-type settings to be used by WAF.

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

<a id="schema-detection_settings--signature_selection_setting--attack_type_settings--disabled_attack_types"></a>

### disabled_attack_types property

Type: `["list", "string"]`. Computed.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of Attack Types that will be ignored and not trigger a detection. Possible values are
\`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

List of Attack Types that will be ignored and not trigger a detection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 22,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 22,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "22",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "22",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [detection_settings.signature_selection_setting](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/properties/detection_settings/signature_selection_setting/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_firewall/)
