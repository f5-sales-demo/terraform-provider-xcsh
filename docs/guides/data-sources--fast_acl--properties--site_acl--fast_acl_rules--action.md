---
page_title: "site_acl.fast_acl_rules.action"
subcategory: ""
description: "site_acl.fast_acl_rules.action for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 2537, "body_sha256": "sha256:3c05b23a2b204156521afd44754181d9334fe2a66faf027880d8ff46469520da", "canonical_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action", "child_ids": ["xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action:policer_action", "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action:protocol_policer_action"], "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules:action", "parent_id": "xcsh-docs:data-sources:fast_acl:properties:site_acl:fast_acl_rules", "path": "docs/guides/data-sources--fast_acl--properties--site_acl--fast_acl_rules--action.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["site_acl", "fast_acl_rules", "action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/properties/site_acl/fast_acl_rules/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "site_acl.fast_acl_rules.action for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_acl.fast_acl_rules.action

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md)
- [Property reference](data-sources--fast_acl--reference.md)
- [site_acl](data-sources--fast_acl--properties--site_acl.md)
- [site_acl.fast_acl_rules](data-sources--fast_acl--properties--site_acl--fast_acl_rules.md)
- site_acl.fast_acl_rules.action

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

- [policer_action](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--policer_action.md): complete subsection reference.

- [protocol_policer_action](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--protocol_policer_action.md): complete subsection reference.

<a id="schema-site_acl--fast_acl_rules--action--simple_action"></a>

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

- [site_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--policer_action.md)
- [site_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--properties--site_acl--fast_acl_rules--action--protocol_policer_action.md)
- [site_acl.fast_acl_rules](data-sources--fast_acl--properties--site_acl--fast_acl_rules.md)
- [xcsh_fast_acl](../data-sources/fast_acl.md)
