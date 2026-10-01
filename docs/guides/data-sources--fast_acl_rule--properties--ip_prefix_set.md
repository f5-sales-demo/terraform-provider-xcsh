---
page_title: "ip_prefix_set"
subcategory: ""
description: "ip_prefix_set for xcsh_fast_acl_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1265, "body_sha256": "sha256:103bdb1423ede1713c54e7f98c1472cc4b141cc28a686ad0d98d5c549d13543c", "canonical_id": "xcsh-docs:data-sources:fast_acl_rule:properties:ip_prefix_set", "child_ids": ["xcsh-docs:data-sources:fast_acl_rule:properties:ip_prefix_set:ref"], "collection_id": "xcsh-docs:data-sources:fast_acl_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl_rule:properties:ip_prefix_set", "parent_id": "xcsh-docs:data-sources:fast_acl_rule:reference", "path": "docs/guides/data-sources--fast_acl_rule--properties--ip_prefix_set.md", "provider_name": "fast_acl_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ip_prefix_set"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl_rule/properties/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ip_prefix_set for xcsh_fast_acl_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_acl_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ip_prefix_set

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md)
- [Property reference](data-sources--fast_acl_rule--reference.md)
- ip_prefix_set

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: ip\_prefix\_set, prefix\] List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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

- [ip_prefix_set](data-sources--fast_acl_rule--properties--ip_prefix_set.md#section)
- [prefix](data-sources--fast_acl_rule--properties--prefix.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [ref](data-sources--fast_acl_rule--properties--ip_prefix_set--ref.md): complete subsection reference.

## Next pages

- [ip_prefix_set.ref](data-sources--fast_acl_rule--properties--ip_prefix_set--ref.md)
- [Property reference](data-sources--fast_acl_rule--reference.md)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md)
