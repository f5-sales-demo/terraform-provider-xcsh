---
page_title: "baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks"
subcategory: ""
description: "baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 8210, "body_sha256": "sha256:9362a3c80987c73e45ba8b9024653e88e39ef32d046764f3755bd827e988b21a", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:dhcp_server:dhcp_networks:first_address", "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:dhcp_server:dhcp_networks:last_address", "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:dhcp_server:dhcp_networks:pools", "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:dhcp_server:dhcp_networks:same_as_dgw"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:dhcp_server:dhcp_networks", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:baremetal:not_managed:node_list:interface_list:dhcp_server", "path": "docs/guides/data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["baremetal", "not_managed", "node_list", "interface_list", "dhcp_server", "dhcp_networks"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/baremetal/not_managed/node_list/interface_list/dhcp_server/dhcp_networks/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [baremetal](data-sources--securemesh_site_v2--properties--baremetal.md)
- [baremetal.not_managed](data-sources--securemesh_site_v2--properties--baremetal--not_managed.md)
- [baremetal.not_managed.node_list](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list.md)
- [baremetal.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list.md)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--dhcp_server.md)
- baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="section"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--dgw_address"></a>

### dgw_address property

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="schema-baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--dns_address"></a>

### dns_address property

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--first_address.md): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--last_address.md): complete subsection reference.

<a id="schema-baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--network_prefix"></a>

### network_prefix property

Type: `"string"`. Computed.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="schema-baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--pool_settings"></a>

### pool_settings property

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--pools.md): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--same_as_dgw.md): complete subsection reference.

## Next pages

- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--first_address.md)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--last_address.md)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--pools.md)
- [baremetal.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--dhcp_server--dhcp_networks--same_as_dgw.md)
- [baremetal.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--properties--baremetal--not_managed--node_list--interface_list--dhcp_server.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
