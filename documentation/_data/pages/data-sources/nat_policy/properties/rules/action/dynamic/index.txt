---
page_title: "rules.action.dynamic"
subcategory: ""
description: "rules.action.dynamic for xcsh_nat_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1912, "body_sha256": "sha256:e14f1b5695c9f68ba74f6ff2c042bfe21e5de00260968a81bdbb543442904cc2", "child_ids": ["xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic:elastic_ips", "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic:pools"], "collection_id": "xcsh-docs:data-sources:nat_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nat_policy:properties:rules:action:dynamic", "parent_id": "xcsh-docs:data-sources:nat_policy:properties:rules:action", "path": "documentation/data-sources/nat_policy/properties/rules/action/dynamic/index.md", "provider_name": "nat_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["rules", "action", "dynamic"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nat_policy/properties/rules/action/dynamic/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.action.dynamic for xcsh_nat_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nat_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.action.dynamic

Breadcrumbs:

- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/)
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

- [elastic_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic/elastic_ips/): complete subsection reference.

- [pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic/pools/): complete subsection reference.

## Next pages

- [rules.action.dynamic.elastic_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic/elastic_ips/)
- [rules.action.dynamic.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/dynamic/pools/)
- [rules.action](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/properties/rules/action/)
- [xcsh_nat_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nat_policy/)
