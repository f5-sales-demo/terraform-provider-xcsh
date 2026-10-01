---
page_title: "nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools"
subcategory: ""
description: "nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 5168, "body_sha256": "sha256:173f66e1409452390e5c9710c221a504b3690451381287a1dc3dcb73097ac3cd", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:nutanix:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "path": "docs/guides/data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["nutanix", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/nutanix/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/pools/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [nutanix](data-sources--securemesh_site_v2--properties--nutanix.md)
- [nutanix.not_managed](data-sources--securemesh_site_v2--properties--nutanix--not_managed.md)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list.md)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list.md)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks.md)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="section"></a>

Type: `"list"`. Computed.

List of non overlapping IP address ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip"></a>

### end_ip property

Type: `"string"`. Computed.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="schema-nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip"></a>

### start_ip property

Type: `"string"`. Computed.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

## Next pages

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--properties--nutanix--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
