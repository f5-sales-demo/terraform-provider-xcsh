---
page_title: "action.protocol_policer_action"
subcategory: ""
description: "action.protocol_policer_action for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1105, "body_sha256": "sha256:e786785fa3d811dbb2be20302088225f0e85c9f7180b2688cbb38eb2399f1eab", "canonical_id": "xcsh-docs:data-sources:fast_acl_rule:properties:action:protocol_policer_action", "child_ids": ["xcsh-docs:data-sources:fast_acl_rule:properties:action:protocol_policer_action:ref"], "collection_id": "xcsh-docs:data-sources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl_rule:properties:action:protocol_policer_action", "parent_id": "xcsh-docs:data-sources:fast_acl_rule:properties:action", "path": "docs/guides/data-sources--fast_acl_rule--properties--action--protocol_policer_action.md", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["action", "protocol_policer_action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl_rule/properties/action/protocol_policer_action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "action.protocol_policer_action for xcsh_fast_acl_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# action.protocol_policer_action

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md)
- [Property reference](data-sources--fast_acl_rule--reference.md)
- [action](data-sources--fast_acl_rule--properties--action.md)
- action.protocol_policer_action

<a id="section"></a>

Type: `"single"`. Computed.

Protocol Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

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

- [ref](data-sources--fast_acl_rule--properties--action--protocol_policer_action--ref.md): complete subsection reference.

## Next pages

- [action.protocol_policer_action.ref](data-sources--fast_acl_rule--properties--action--protocol_policer_action--ref.md)
- [action](data-sources--fast_acl_rule--properties--action.md)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md)
