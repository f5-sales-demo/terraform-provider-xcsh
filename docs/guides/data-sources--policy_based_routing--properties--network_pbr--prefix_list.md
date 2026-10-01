---
page_title: "network_pbr.prefix_list"
subcategory: ""
description: "network_pbr.prefix_list for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 1921, "body_sha256": "sha256:82d8dae91228de051d2337b09bfd5ba5fc43c1f466a48520f132fd4bf7bcb8e0", "canonical_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:prefix_list", "child_ids": [], "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr:prefix_list", "parent_id": "xcsh-docs:data-sources:policy_based_routing:properties:network_pbr", "path": "docs/guides/data-sources--policy_based_routing--properties--network_pbr--prefix_list.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["network_pbr", "prefix_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/properties/network_pbr/prefix_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "network_pbr.prefix_list for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# network_pbr.prefix_list

Breadcrumbs:

- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
- [Property reference](data-sources--policy_based_routing--reference.md)
- [network_pbr](data-sources--policy_based_routing--properties--network_pbr.md)
- network_pbr.prefix_list

<a id="section"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="schema-network_pbr--prefix_list--prefixes"></a>

### prefixes property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Next pages

- [network_pbr](data-sources--policy_based_routing--properties--network_pbr.md)
- [xcsh_policy_based_routing](../data-sources/policy_based_routing.md)
