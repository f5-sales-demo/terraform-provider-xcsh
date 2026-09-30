---
page_title: "action"
subcategory: ""
description: "action for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": [], "body_bytes": 2075, "body_sha256": "sha256:3a011c5adf495ac487df2303baa0aa7cb6bec523ca2dcb1aa6626e9062fed783", "canonical_id": "xcsh-docs:data-sources:fast_acl_rule:properties:action", "child_ids": ["xcsh-docs:data-sources:fast_acl_rule:properties:action:policer_action", "xcsh-docs:data-sources:fast_acl_rule:properties:action:protocol_policer_action"], "collection_id": "xcsh-docs:data-sources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl_rule:properties:action", "parent_id": "xcsh-docs:data-sources:fast_acl_rule:reference", "path": "docs/guides/data-sources--fast_acl_rule--properties--action.md", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl_rule/properties/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "action for xcsh_fast_acl_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# action

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md)
- [Property reference](data-sources--fast_acl_rule--reference.md)
- action

<a id="section"></a>

Type: `"single"`. Computed.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

## Direct properties

- [policer_action](data-sources--fast_acl_rule--properties--action--policer_action.md): complete subsection reference.

- [protocol_policer_action](data-sources--fast_acl_rule--properties--action--protocol_policer_action.md): complete subsection reference.

<a id="schema-action--simple_action"></a>

### simple_action property

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
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

- [action.policer_action](data-sources--fast_acl_rule--properties--action--policer_action.md)
- [action.protocol_policer_action](data-sources--fast_acl_rule--properties--action--protocol_policer_action.md)
- [Property reference](data-sources--fast_acl_rule--reference.md)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md)
