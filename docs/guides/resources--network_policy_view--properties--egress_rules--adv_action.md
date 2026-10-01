---
page_title: "egress_rules.adv_action"
subcategory: ""
description: "egress_rules.adv_action for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 2240, "body_sha256": "sha256:86eb8889dbba8b3c938dc4867de0090bedddb9d86b86f0175fb251ca9d7dda3e", "canonical_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:adv_action", "child_ids": [], "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:egress_rules:adv_action", "parent_id": "xcsh-docs:resources:network_policy_view:properties:egress_rules", "path": "docs/guides/resources--network_policy_view--properties--egress_rules--adv_action.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["egress_rules", "adv_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/egress_rules/adv_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "egress_rules.adv_action for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# egress_rules.adv_action

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md)
- [Property reference](resources--network_policy_view--reference.md)
- [egress_rules](resources--network_policy_view--properties--egress_rules.md)
- egress_rules.adv_action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
adv_action {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-egress_rules--adv_action--action"></a>

### action property

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NOLOG",
    "LOG"),
}
```

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

- [egress_rules](resources--network_policy_view--properties--egress_rules.md)
- [xcsh_network_policy_view](../resources/network_policy_view.md)
