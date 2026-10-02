---
page_title: "ingress_rules.adv_action"
subcategory: ""
description: "Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and PBRRuleAction."
xcsh_docs: {"aliases": ["ingress rules adv action"], "body_bytes": 2504, "body_sha256": "sha256:0cde34be8c0df7f6b89ae07f6c4dd0745f8904a248fad7adadc193c17d6e64eb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules:adv_action", "parent_id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules", "path": "documentation/resources/network_policy_view/properties/ingress_rules/adv_action/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2221201022221322-1203331112301322-2230223112212012-2231233033133130-3312312301113323-3012311113300322-1200022330023011-3320103102003012", "registry_path": "docs/guides/resources--network_policy_view--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_rules", "adv_action"], "schema_version": 1, "sections": [{"aliases": ["action"], "anchor": "schema-ingress_rules--adv_action--action", "description": "Choice to choose logging or no logging This works together with option selected via NetworkPolicyRuleAction or any other action specified x-example: (No Selection in NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction + AdvancedAction as LOG) = Log and Allow/Deny,", "document_id": "xcsh-docs:resources:network_policy_view:properties:ingress_rules:adv_action", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ingress_rules", "adv_action", "action"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/ingress_rules/adv_action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and PBRRuleAction.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_rules.adv_action

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/)
- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/ingress_rules/)
- ingress_rules.adv_action

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

<a id="schema-ingress_rules--adv_action--action"></a>

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

- [ingress_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/properties/ingress_rules/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_view/)
