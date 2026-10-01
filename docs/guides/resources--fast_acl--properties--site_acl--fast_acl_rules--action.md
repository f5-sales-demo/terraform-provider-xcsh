---
page_title: "site_acl.fast_acl_rules.action"
subcategory: ""
description: "site_acl.fast_acl_rules.action for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 3127, "body_sha256": "sha256:1c3ab468d8ec76b65fcd674fed32f5ab1d07fcd4fd7e1aa089b22e847f4dd8da", "canonical_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action", "child_ids": ["xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action", "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action:protocol_policer_action"], "collection_id": "xcsh-docs:resources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules:action", "parent_id": "xcsh-docs:resources:fast_acl:properties:site_acl:fast_acl_rules", "path": "docs/guides/resources--fast_acl--properties--site_acl--fast_acl_rules--action.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_acl", "fast_acl_rules", "action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fast_acl/properties/site_acl/fast_acl_rules/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_acl.fast_acl_rules.action for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.fast_acl_rules.action

Breadcrumbs:

- [xcsh_fast_acl](../resources/fast_acl.md)
- [Property reference](resources--fast_acl--reference.md)
- [site_acl](resources--fast_acl--properties--site_acl.md)
- [site_acl.fast_acl_rules](resources--fast_acl--properties--site_acl--fast_acl_rules.md)
- site_acl.fast_acl_rules.action

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

- [policer_action](resources--fast_acl--properties--site_acl--fast_acl_rules--action--policer_action.md): complete subsection reference.

- [protocol_policer_action](resources--fast_acl--properties--site_acl--fast_acl_rules--action--protocol_policer_action.md): complete subsection reference.

<a id="schema-site_acl--fast_acl_rules--action--simple_action"></a>

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

- [site_acl.fast_acl_rules.action.policer_action](resources--fast_acl--properties--site_acl--fast_acl_rules--action--policer_action.md)
- [site_acl.fast_acl_rules.action.protocol_policer_action](resources--fast_acl--properties--site_acl--fast_acl_rules--action--protocol_policer_action.md)
- [site_acl.fast_acl_rules](resources--fast_acl--properties--site_acl--fast_acl_rules.md)
- [xcsh_fast_acl](../resources/fast_acl.md)
