---
page_title: "rule_list.rules.advanced_action"
subcategory: ""
description: "rule_list.rules.advanced_action for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2144, "body_sha256": "sha256:a0cca8252e64dee3ac3eb57643a0ba58681b7654bebe10f3dbaccbe13137431a", "canonical_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:advanced_action", "child_ids": [], "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules:advanced_action", "parent_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:rule_list:rules", "path": "docs/guides/data-sources--enhanced_firewall_policy--properties--rule_list--rules--advanced_action.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rule_list", "rules", "advanced_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/properties/rule_list/rules/advanced_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rule_list.rules.advanced_action for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rule_list.rules.advanced_action

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md)
- [Property reference](data-sources--enhanced_firewall_policy--reference.md)
- [rule_list](data-sources--enhanced_firewall_policy--properties--rule_list.md)
- [rule_list.rules](data-sources--enhanced_firewall_policy--properties--rule_list--rules.md)
- rule_list.rules.advanced_action

<a id="section"></a>

Type: `"single"`. Computed.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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

<a id="schema-rule_list--rules--advanced_action--action"></a>

### action property

Type: `"string"`. Computed.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
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

- [rule_list.rules](data-sources--enhanced_firewall_policy--properties--rule_list--rules.md)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md)
