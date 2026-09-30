---
page_title: "rule_list.rules.applications"
subcategory: ""
description: "rule_list.rules.applications for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1916, "body_sha256": "sha256:6277c5418093c15d597f21ac0261af506e115822826215390883e549b3cb97cd", "child_ids": [], "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:applications", "parent_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules", "path": "documentation/data-sources/enhanced_firewall_policy/properties/rule_list/rules/applications/index.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["rule_list", "rules", "applications"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/properties/rule_list/rules/applications/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.applications for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rule_list.rules.applications

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/)
- [rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/)
- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/)
- rule_list.rules.applications

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="schema-rule_list--rules--applications--applications"></a>

### applications property

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

## Next pages

- [rule_list.rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/properties/rule_list/rules/)
- [xcsh_enhanced_firewall_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/enhanced_firewall_policy/)
