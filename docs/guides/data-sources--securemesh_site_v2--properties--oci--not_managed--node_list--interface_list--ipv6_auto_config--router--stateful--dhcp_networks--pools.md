---
page_title: "oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools"
subcategory: ""
description: "oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 4981, "body_sha256": "sha256:91d84056f0810169cc2cd103147f7f4f59f9943e0044c99e74ed2a08540b20e8", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks:pools", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:oci:not_managed:node_list:interface_list:ipv6_auto_config:router:stateful:dhcp_networks", "path": "docs/guides/data-sources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["oci", "not_managed", "node_list", "interface_list", "ipv6_auto_config", "router", "stateful", "dhcp_networks", "pools"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/oci/not_managed/node_list/interface_list/ipv6_auto_config/router/stateful/dhcp_networks/pools/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [oci](data-sources--securemesh_site_v2--properties--oci.md)
- [oci.not_managed](data-sources--securemesh_site_v2--properties--oci--not_managed.md)
- [oci.not_managed.node_list](data-sources--securemesh_site_v2--properties--oci--not_managed--node_list.md)
- [oci.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list.md)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config.md)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router.md)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful.md)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks.md)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

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

<a id="schema-oci--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip"></a>

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

<a id="schema-oci--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip"></a>

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

- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--properties--oci--not_managed--node_list--interface_list--ipv6_auto_config--router--stateful--dhcp_networks.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
