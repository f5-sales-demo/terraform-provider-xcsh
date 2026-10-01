---
page_title: "allow_all"
subcategory: ""
description: "allow_all for xcsh_enhanced_firewall_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1866, "body_sha256": "sha256:9908a2bbc063959e29c71aa79b82527ced8720e0d52a957d26741d5adf95c90d", "canonical_id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:allow_all", "child_ids": [], "collection_id": "xcsh-docs:data-sources:enhanced_firewall_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:enhanced_firewall_policy:properties:allow_all", "parent_id": "xcsh-docs:data-sources:enhanced_firewall_policy:reference", "path": "docs/guides/data-sources--enhanced_firewall_policy--properties--allow_all.md", "provider_name": "enhanced_firewall_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["allow_all"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/enhanced_firewall_policy/properties/allow_all/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_all for xcsh_enhanced_firewall_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["enhanced_firewall_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_all

Breadcrumbs:

- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md)
- [Property reference](data-sources--enhanced_firewall_policy--reference.md)
- allow_all

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all, allowed\_destinations, allowed\_sources, denied\_destinations, denied\_sources,
deny\_all, rule\_list\] Enable this option. Defaults to \`map\[\]\`. Server applies default when
omitted.

Upstream description:

This can be used for messages where no values are needed.

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

OneOf alternatives in this subsection:

- [allow_all](data-sources--enhanced_firewall_policy--properties--allow_all.md#section)
- [allowed_destinations](data-sources--enhanced_firewall_policy--properties--allowed_destinations.md#section)
- [allowed_sources](data-sources--enhanced_firewall_policy--properties--allowed_sources.md#section)
- [denied_destinations](data-sources--enhanced_firewall_policy--properties--denied_destinations.md#section)
- [denied_sources](data-sources--enhanced_firewall_policy--properties--denied_sources.md#section)
- [deny_all](data-sources--enhanced_firewall_policy--properties--deny_all.md#section)
- [rule_list](data-sources--enhanced_firewall_policy--properties--rule_list.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--enhanced_firewall_policy--reference.md)
- [xcsh_enhanced_firewall_policy](../data-sources/enhanced_firewall_policy.md)
