---
page_title: "allow_all"
subcategory: "Security"
description: "allow_all for xcsh_forward_proxy_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1373, "body_sha256": "sha256:a186bc76c22df7e94281b87575819c8e7ce9f1ed6b2f01eae6df7188835d0338", "canonical_id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_all", "child_ids": [], "collection_id": "xcsh-docs:data-sources:forward_proxy_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:forward_proxy_policy:properties:allow_all", "parent_id": "xcsh-docs:data-sources:forward_proxy_policy:reference", "path": "docs/guides/data-sources--forward_proxy_policy--properties--allow_all.md", "provider_name": "forward_proxy_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["allow_all"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/forward_proxy_policy/properties/allow_all/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "allow_all for xcsh_forward_proxy_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forward_proxy_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# allow_all

Breadcrumbs:

- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md)
- [Property reference](data-sources--forward_proxy_policy--reference.md)
- allow_all

<a id="section"></a>

Type: `["object", {}]`. Computed.

\[OneOf: allow\_all, allow\_list, deny\_list, rule\_list\] Enable this option

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

- [allow_all](data-sources--forward_proxy_policy--properties--allow_all.md#section)
- [allow_list](data-sources--forward_proxy_policy--properties--allow_list.md#section)
- [deny_list](data-sources--forward_proxy_policy--properties--deny_list.md#section)
- [rule_list](data-sources--forward_proxy_policy--properties--rule_list.md#section)

Select alternatives according to the provider validators above.

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](data-sources--forward_proxy_policy--reference.md)
- [xcsh_forward_proxy_policy](../data-sources/forward_proxy_policy.md)
