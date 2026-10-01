---
page_title: "action"
subcategory: ""
description: "action for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": [], "body_bytes": 2770, "body_sha256": "sha256:c6bd4b77ff2d8b3fb03747e39e7cdbf9d1ae7e719af62889dcb0e5943ec6a281", "canonical_id": "xcsh-docs:resources:fast_acl_rule:properties:action", "child_ids": ["xcsh-docs:resources:fast_acl_rule:properties:action:policer_action", "xcsh-docs:resources:fast_acl_rule:properties:action:protocol_policer_action"], "collection_id": "xcsh-docs:resources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl_rule:properties:action", "parent_id": "xcsh-docs:resources:fast_acl_rule:reference", "path": "docs/guides/resources--fast_acl_rule--properties--action.md", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl_rule/properties/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "action for xcsh_fast_acl_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# action

Breadcrumbs:

- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md)
- [Property reference](resources--fast_acl_rule--reference.md)
- action

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("policer_action",
    "protocol_policer_action"),
  validators.ConflictingObjectAttributes("policer_action",
    "simple_action"),
  validators.ConflictingObjectAttributes("protocol_policer_action",
    "simple_action")}
```

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

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

## Direct properties

- [policer_action](resources--fast_acl_rule--properties--action--policer_action.md): complete subsection reference.

- [protocol_policer_action](resources--fast_acl_rule--properties--action--protocol_policer_action.md): complete subsection reference.

<a id="schema-action--simple_action"></a>

### simple_action property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

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

- [action.policer_action](resources--fast_acl_rule--properties--action--policer_action.md)
- [action.protocol_policer_action](resources--fast_acl_rule--properties--action--protocol_policer_action.md)
- [Property reference](resources--fast_acl_rule--reference.md)
- [xcsh_fast_acl_rule](../resources/fast_acl_rule.md)
