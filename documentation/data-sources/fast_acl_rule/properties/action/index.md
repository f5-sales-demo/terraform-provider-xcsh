---
page_title: "action"
subcategory: ""
description: "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic."
xcsh_docs: {"aliases": ["action"], "body_bytes": 2582, "body_sha256": "sha256:cd2eda7ebf108abf868049867c53908a4991404b7969a2f165758d9177266ac6", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:fast_acl_rule:properties:action:policer_action", "xcsh-docs:data-sources:fast_acl_rule:properties:action:protocol_policer_action"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl_rule:properties:action", "parent_id": "xcsh-docs:data-sources:fast_acl_rule:reference", "path": "documentation/data-sources/fast_acl_rule/properties/action/index.md", "product": "distributed-cloud", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2323322103100113-0330221003132312-0022232121231222-2001112100311031-3033102332230020-2200221013103021-0101203130002020-0333000311113233", "registry_path": "docs/guides/data-sources--fast_acl_rule--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["action"], "schema_version": 1, "sections": [{"aliases": ["policer action"], "anchor": "section", "description": "Reference to policer object.", "document_id": "xcsh-docs:data-sources:fast_acl_rule:properties:action:policer_action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["action", "policer_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["protocol policer action"], "anchor": "section", "description": "Reference to policer object.", "document_id": "xcsh-docs:data-sources:fast_acl_rule:properties:action:protocol_policer_action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["action", "protocol_policer_action"], "syntax": "attribute", "type": "object"}, {"aliases": ["simple action"], "anchor": "schema-action--simple_action", "description": "FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the traffic Forward the traffic.", "document_id": "xcsh-docs:data-sources:fast_acl_rule:properties:action", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["action", "simple_action"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl_rule/properties/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "FastAclRuleAction specifies possible action to be applied on traffic, possible action include dropping, forwarding or ratelimiting the traffic.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# action

Breadcrumbs:

- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/)
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

- [policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/action/policer_action/): complete subsection reference.

- [protocol_policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/action/protocol_policer_action/): complete subsection reference.

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

- [action.policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/action/policer_action/)
- [action.protocol_policer_action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/action/protocol_policer_action/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/properties/)
- [xcsh_fast_acl_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl_rule/)
