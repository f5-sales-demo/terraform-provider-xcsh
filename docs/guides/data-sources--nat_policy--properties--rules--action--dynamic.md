---
page_title: "rules.action.dynamic"
subcategory: ""
description: "rules.action.dynamic for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1414, "body_sha256": "sha256:72417ca422de43e7790d1a0a079abffe24f429748900df64cd48e0f61872e540", "canonical_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic:elastic_ips", "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic:pools"], "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action", "path": "docs/guides/data-sources--nat_policy--properties--rules--action--dynamic.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "action", "dynamic"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/action/dynamic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.action.dynamic for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.dynamic

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md)
- [Property reference](data-sources--nat_policy--reference.md)
- [rules](data-sources--nat_policy--properties--rules.md)
- [rules.action](data-sources--nat_policy--properties--rules--action.md)
- rules.action.dynamic

<a id="section"></a>

Type: `"single"`. Computed.

Dynamic Pool. Dynamic Pool Configuration.

Upstream description:

Dynamic Pool Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-pool_choice": "[\"elastic_ips\",\"pools\"]"
}
```

## Direct properties

- [elastic_ips](data-sources--nat_policy--properties--rules--action--dynamic--elastic_ips.md): complete subsection reference.

- [pools](data-sources--nat_policy--properties--rules--action--dynamic--pools.md): complete subsection reference.

## Next pages

- [rules.action.dynamic.elastic_ips](data-sources--nat_policy--properties--rules--action--dynamic--elastic_ips.md)
- [rules.action.dynamic.pools](data-sources--nat_policy--properties--rules--action--dynamic--pools.md)
- [rules.action](data-sources--nat_policy--properties--rules--action.md)
- [xcsh_nat_policy](../data-sources/nat_policy.md)
