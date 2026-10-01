---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-8e85afd7516bfd431bba5c78526f8b617f0212bdffd13fe88c516a825b515a46"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.a / 5b4aad1e1b81 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--voltstack_site--reference--group-003.md#canonical-c81ed23b266728c5cfed34d6c40f0d14e259d1d3657b5a53bcd23083e2f478c5)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start

<a id="canonical-b65e15ea2d66a5ccc88a9dbf02b3644b624709426b731f768b330e9488da54be"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-c4ecc2236d9dcd7135993d1d727e2ac6fde00d0d520393a54bdc69cc29f4cdff"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.a / 5b4aad1e1b81 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f386c2c3afe20d2519084e07d9a37966f6cc08de5bb51907d5c6b608e7faca43"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.a / 5b4aad1e1b81 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--voltstack_site--reference--group-003.md#canonical-c81ed23b266728c5cfed34d6c40f0d14e259d1d3657b5a53bcd23083e2f478c5)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-99b04f4144eb863908de924dbabf324e58bed0d1bee9e239468bdc5f5b6c24a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6146cf85e204cc9262d934fff17675e2dd296f039e0e3c18281cfbbc354b21b"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 938d069a9578 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--voltstack_site--reference--group-003.md#canonical-c81ed23b266728c5cfed34d6c40f0d14e259d1d3657b5a53bcd23083e2f478c5)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks

<a id="canonical-7470235d253329c38cb75e1ea4bef6b04ec6ae79b7cde214b22ee15f33cd19f6"></a>

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

<a id="canonical-af62b3cdf465a9f3ce4302876024490cbc37636fd6605792fde686dc33e103c0"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 938d069a9578 / 3

<a id="canonical-821a454592b4a1c99e02cc137ccbe68ab7c606a769d1a873cf8d841addd9c697"></a>

<a id="canonical-2766224d5ed0ea837e86a3e9daf01a4d43e726ce7f8f8393b538e0b5517b88cf"></a>

## dgw_address property — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 938d069a9578 / 4

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

<a id="canonical-d3eed82fc560cca9e3c16c4063cb1d84776d8ec24bc6df9dfab82c2f26edbb8b"></a>

<a id="canonical-221b697318af818eb546f7e81661c2400feb33c635008571e2b21f618633f063"></a>

## dns_address property — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 938d069a9578 / 5

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

- [first_address](data-sources--voltstack_site--reference--group-004.md#canonical-559c4dc8a86c79778e3eaac058f580359057a2e40a4c3d2a9f4fadc8d2ddf2bd): complete subsection reference.

- [last_address](data-sources--voltstack_site--reference--group-004.md#canonical-38402db59f4d04dac890a453071f818fa6d3bed6946b6022482d608341462152): complete subsection reference.

<a id="canonical-a4a5794c58788c654a40c1e9d22f012c47a9b73ed5163f20d6e9cf739b0e6e6e"></a>

<a id="canonical-6a51cb83f76d40a39a34e68eb09386a977c2a93cbf826f837e78c05e10fac896"></a>

## network_prefix property — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 938d069a9578 / 6

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

<a id="canonical-2dbbee1c4fcb0a94cc21b141208971d2cd89f0ac7d3b3b3420e2547dbc6988ad"></a>

<a id="canonical-0971f722ab98a2528911e92cc7934c1faf220653af51d1464eb9769ab9df111b"></a>

## pool_settings property — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 938d069a9578 / 7

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

- [pools](data-sources--voltstack_site--reference--group-004.md#canonical-097533dd274555003a6e634550b3f4c95c20075f65219998a7f28ebc32514ff1): complete subsection reference.

- [same_as_dgw](data-sources--voltstack_site--reference--group-004.md#canonical-5bb90e914538034a78102b94ea0256b2caf065683db089b936907b4f4bd18e38): complete subsection reference.

<a id="canonical-2b745f4fec589ca4c4047a6d9b98fa84068f8cd7286a2904b1c31d140e58d81a"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 938d069a9578 / 8

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address](data-sources--voltstack_site--reference--group-004.md#canonical-559c4dc8a86c79778e3eaac058f580359057a2e40a4c3d2a9f4fadc8d2ddf2bd)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address](data-sources--voltstack_site--reference--group-004.md#canonical-38402db59f4d04dac890a453071f818fa6d3bed6946b6022482d608341462152)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools](data-sources--voltstack_site--reference--group-004.md#canonical-097533dd274555003a6e634550b3f4c95c20075f65219998a7f28ebc32514ff1)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](data-sources--voltstack_site--reference--group-004.md#canonical-5bb90e914538034a78102b94ea0256b2caf065683db089b936907b4f4bd18e38)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--voltstack_site--reference--group-003.md#canonical-c81ed23b266728c5cfed34d6c40f0d14e259d1d3657b5a53bcd23083e2f478c5)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-559c4dc8a86c79778e3eaac058f580359057a2e40a4c3d2a9f4fadc8d2ddf2bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b1f15f2f598a78cf10e1a9e10101cdb6f4cc40f51daa95c348a756e2c4b1847"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 6fbeda723b1a / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--voltstack_site--reference--group-003.md#canonical-c81ed23b266728c5cfed34d6c40f0d14e259d1d3657b5a53bcd23083e2f478c5)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-99b04f4144eb863908de924dbabf324e58bed0d1bee9e239468bdc5f5b6c24a6)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address

<a id="canonical-7d415de2a5d530091cc5f79b44312b5c092ba5ad8bd831249323f6f767d707e9"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-38a6d9344df6bf97344d76437c7d9009173be1abe54da2d714a3cce9340e2a35"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 6fbeda723b1a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-58af9288cec0dcc7439c9b7ec455097f5027a03659843de96d2080c723d6c2b0"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 6fbeda723b1a / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-99b04f4144eb863908de924dbabf324e58bed0d1bee9e239468bdc5f5b6c24a6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-38402db59f4d04dac890a453071f818fa6d3bed6946b6022482d608341462152"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-992184e8f6a1b6af5c23cfafd544ba2abc2db3e2480ef2e97d30a5a5621f4aa4"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / aa3ba941a001 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--voltstack_site--reference--group-003.md#canonical-c81ed23b266728c5cfed34d6c40f0d14e259d1d3657b5a53bcd23083e2f478c5)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-99b04f4144eb863908de924dbabf324e58bed0d1bee9e239468bdc5f5b6c24a6)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address

<a id="canonical-ba94246043faf13dfd968557835e5cd00e920216eee909560fdaefede5e57f0f"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-2d0b64d7cc0b1f0b0f07b77f6d1064838988e568d1847faba0c22c09fe470536"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / aa3ba941a001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5dd9842e172ee32f879e4a190e73bc571271a24ff41fe14cef07dbe1c655e710"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / aa3ba941a001 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-99b04f4144eb863908de924dbabf324e58bed0d1bee9e239468bdc5f5b6c24a6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-097533dd274555003a6e634550b3f4c95c20075f65219998a7f28ebc32514ff1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fa6f2bd4fafc661d6832ddb3a3668aa579ba2600c6773acef581f611467c3e3"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 416e6f7fcae7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--voltstack_site--reference--group-003.md#canonical-c81ed23b266728c5cfed34d6c40f0d14e259d1d3657b5a53bcd23083e2f478c5)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-99b04f4144eb863908de924dbabf324e58bed0d1bee9e239468bdc5f5b6c24a6)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools

<a id="canonical-81e6b93b2720c6e27f4a2a9590da43a8d18f7157f5c7fc630cd985c586a72b57"></a>

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

<a id="canonical-dd1cfa184bd57662688ed1908d22f48761a56d861f58c460a6ada370e4892c61"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 416e6f7fcae7 / 3

<a id="canonical-96c7d03abc5f20ee539b607d77974bad81c3a096d6e9384df7fbfadb5a21e385"></a>

<a id="canonical-23e7d96103447fbbf7ee2cc35306fb9387ff61ffc494bea083960ee37cee75f7"></a>

## end_ip property — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 416e6f7fcae7 / 4

Type: `"string"`. Computed.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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

<a id="canonical-b8de3d8bd8bb7a919c63360b1ba00d2816afe8289c70691eb58803884347d36c"></a>

<a id="canonical-38b9933afabcbb2fa8caf2ea1bf1fd49c0e37c6e3998aa07d81181dc313cf85c"></a>

## exclude property — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 416e6f7fcae7 / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-e63ee3195c63e9a2f05427340a1fca95afde8e9c5f9f13e3eb12f6a57b7d67f9"></a>

<a id="canonical-b92ded0982321e2d194f56a9da886ba580687a4b849e552e352a2e106478c109"></a>

## start_ip property — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 416e6f7fcae7 / 6

Type: `"string"`. Computed.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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

<a id="canonical-b1275c4790fab441882c215effd849c6f9d847005588970628f4d6532fbeacfd"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / 416e6f7fcae7 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-99b04f4144eb863908de924dbabf324e58bed0d1bee9e239468bdc5f5b6c24a6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-5bb90e914538034a78102b94ea0256b2caf065683db089b936907b4f4bd18e38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77cccba0b253cc4cb688804797c68a9cd30aae7d6100790a8b058d3a5be52ac2"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / d3d5a3a08df8 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--voltstack_site--reference--group-003.md#canonical-c81ed23b266728c5cfed34d6c40f0d14e259d1d3657b5a53bcd23083e2f478c5)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-99b04f4144eb863908de924dbabf324e58bed0d1bee9e239468bdc5f5b6c24a6)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-69150e8f0699d3e7f7c736ac00a3300f38bf8833c7048b79639219d64052d741"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for same as dgw.

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

<a id="canonical-f8d48fac3ff5b73d2000b713cbe634d3adb3cdf08b92fb03080c54bdd4f302a3"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / d3d5a3a08df8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bd7958d0bf8aab43ad527dc545e48081f965a555768d1079d836ef44bf29dfe3"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.d / d3d5a3a08df8 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-99b04f4144eb863908de924dbabf324e58bed0d1bee9e239468bdc5f5b6c24a6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-26828be93fc0e0ba47ad980e653b454a8e818d7032e5ec0e62526a5b3c9dd4a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24a4f5bd16142ca84d4a3dbda2074e93c7ec16e44635dac8bbe6ca2098b4d21e"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.i / db2c3392510c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--voltstack_site--reference--group-003.md#canonical-c81ed23b266728c5cfed34d6c40f0d14e259d1d3657b5a53bcd23083e2f478c5)
- custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map

<a id="canonical-aad22d0c1bd9b4071dc7425e666570c5a51c0223962d2c87fea1043fc751351b"></a>

Type: `"single"`. Computed.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

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

<a id="canonical-36f3a23e28159b08b9ead743c7c8f4f65979bd6447e4a486b8287e53a271b6b3"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.i / db2c3392510c / 3

<a id="canonical-d37ec492d0ffed0155aea5a5dcb4c07c1b59cfbe611dd7c7590ae13beed0b031"></a>

<a id="canonical-eafb70de6256155bc1e319f10724c3f89844e1a4765e751830b4c6a09a5e5353"></a>

## interface_ip_map property — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.i / db2c3392510c / 4

Type: `["map", "string"]`. Computed.

Specify static IPv4 addresses per site:node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-51af33e524ac8f4a6654a452cf8f98cc06348a0ba0b04e6cdb298f67a3630a1b"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.i / db2c3392510c / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](data-sources--voltstack_site--reference--group-003.md#canonical-c81ed23b266728c5cfed34d6c40f0d14e259d1d3657b5a53bcd23083e2f478c5)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39aa75b9d5600f8f510a8825ccdc5fb3e6365504ef3d715d6504df73baee9b6f"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / affdeb04f09c / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config

<a id="canonical-8c77b5691e1a319c68a9f26e5a898a5fe99b56bbcb1696e1ea6c346415bc5adc"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

<a id="canonical-6ea43ef3df6ecc7282084346a0e4471df33ce1fa898efc3b25cec08082eed6ef"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / affdeb04f09c / 3

- [host](data-sources--voltstack_site--reference--group-004.md#canonical-0ca276d09b40917e0d1041328eec5edd87f6c76091b40406cee39e64ca365a58): complete subsection reference.

- [router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86): complete subsection reference.

<a id="canonical-28aea7083cf09e95bdb3247e41e6d62dee9313043a0d0fb0c222116a41a2fc96"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / affdeb04f09c / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](data-sources--voltstack_site--reference--group-004.md#canonical-0ca276d09b40917e0d1041328eec5edd87f6c76091b40406cee39e64ca365a58)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-0ca276d09b40917e0d1041328eec5edd87f6c76091b40406cee39e64ca365a58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44e9342d264b0a78f963c4c2ec63019668687c6901baf3d1fc111e9af6b96d9e"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 8e9ad0413769 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host

<a id="canonical-75a62835037dba2cb73a7ea620797ecc087758d7edc6f0cac04c66a0d957d529"></a>

Type: `["object", {}]`. Computed.

Hostname or IP address of the target server.

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

<a id="canonical-537804a34d69e96cbe51143e09e70e9f18a046ed704229e6333f338863e09aa5"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 8e9ad0413769 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3d2069310b83325b1494b62dfc4ec3773e438eb2603d5a3a64ccd3ca861aa9f8"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 8e9ad0413769 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f5b5cd4ad001cfcfad7e4050dc67748c8ccb19f242dd396555e66f46ac70858"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 74c4ea13692e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router

<a id="canonical-281e3bcc275677587a9cd05af8b12d221a54a1dcc909e3407bd85960e35157dd"></a>

Type: `"single"`. Computed.

IPV6AutoConfigRouterType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

<a id="canonical-d6cee24cfeb5d52a31c2a1079c284adbb9c3cc12ff2e01881837b3f0e4f78a58"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 74c4ea13692e / 3

- [dns_config](data-sources--voltstack_site--reference--group-004.md#canonical-e5fbc479fbd4e8cf7af196b76a792b3d1b9254643b90cf5279fc982873f2deb6): complete subsection reference.

<a id="canonical-80e89cf0175591eaaa6d967663269093928e73488da1d0367e0c3589f98b0946"></a>

<a id="canonical-64e9d0e08027380f263e86f3633b99cfd76f41e839a5720128023d05812cc7b9"></a>

## network_prefix property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 74c4ea13692e / 4

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](data-sources--voltstack_site--reference--group-004.md#canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6): complete subsection reference.

<a id="canonical-e0f238ce80e6c7a2f89b2a4521b3ddb794cdaa15ed365798a58b6bb11b41efef"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 74c4ea13692e / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-004.md#canonical-e5fbc479fbd4e8cf7af196b76a792b3d1b9254643b90cf5279fc982873f2deb6)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-004.md#canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-e5fbc479fbd4e8cf7af196b76a792b3d1b9254643b90cf5279fc982873f2deb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb125c99cc7b5f6bc2092f0510faff3e137cbd2496e127b322cc85d845cf27b7"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 349c2c889391 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config

<a id="canonical-3beeaa323bbb901d48ce2da2be8cf80759dd9c54b809006839ee88d07183aea4"></a>

Type: `"single"`. Computed.

IPV6DnsConfig.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

<a id="canonical-2891ac4c3616b827d553ca0a8c4ea1061e7644f5bf2f2f4403a3807782037fe7"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 349c2c889391 / 3

- [configured_list](data-sources--voltstack_site--reference--group-004.md#canonical-532a72ff58991061f05f950c32db8cd58273a76018a6c51bc83c2a5e03311707): complete subsection reference.

- [local_dns](data-sources--voltstack_site--reference--group-004.md#canonical-b7e56ca5c228cd36f908fefc50eba0a614fabffdcc1a145d8b26e1df9df4d166): complete subsection reference.

<a id="canonical-a1545d978ba87dcf87bccfc144ed384f08a7b172c652fd768f68c6394c04c258"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 349c2c889391 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](data-sources--voltstack_site--reference--group-004.md#canonical-532a72ff58991061f05f950c32db8cd58273a76018a6c51bc83c2a5e03311707)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--voltstack_site--reference--group-004.md#canonical-b7e56ca5c228cd36f908fefc50eba0a614fabffdcc1a145d8b26e1df9df4d166)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-532a72ff58991061f05f950c32db8cd58273a76018a6c51bc83c2a5e03311707"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d927931bdb9dfe5012ebaa330ae119d0c6c861a25eefd090c2df1995ca98c366"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 235df51e1f23 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-004.md#canonical-e5fbc479fbd4e8cf7af196b76a792b3d1b9254643b90cf5279fc982873f2deb6)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-cd6cc29e71f669b21f1b66489c77ac3d3f7e3c955f14e96388b6de9dfa3fa86f"></a>

Type: `"single"`. Computed.

IPV6DnsList.

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

<a id="canonical-1ce7153e05198705287508a02fd32c580db84c556b9a692f00dce920f1208de6"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 235df51e1f23 / 3

<a id="canonical-f14c78cd2765d287b94c8069a4b43d294bb3b4f55949bd0361c8ed9d8dffc4e4"></a>

<a id="canonical-1b008316a022b629cd8fb9b84a316097f512debd1f9521fa7981f0f104f2808e"></a>

## dns_list property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 235df51e1f23 / 4

Type: `["list", "string"]`. Computed.

List of IPv6 Addresses acting as DNS servers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9e91bc9a429ae3625b1cba5f5b7b80c756e7f8d6ad0f4b9b0f27a032b77eba34"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 235df51e1f23 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-004.md#canonical-e5fbc479fbd4e8cf7af196b76a792b3d1b9254643b90cf5279fc982873f2deb6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-b7e56ca5c228cd36f908fefc50eba0a614fabffdcc1a145d8b26e1df9df4d166"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc6468d5c0a53919f6b102891e56f43b14590f5174f989020a9cb7c557013195"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / d3574863ec07 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-004.md#canonical-e5fbc479fbd4e8cf7af196b76a792b3d1b9254643b90cf5279fc982873f2deb6)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-036c5fb538acd70bef94e16a401d5e013f97eba393f3f0458680aab9c3df64d4"></a>

Type: `"single"`. Computed.

IPV6LocalDnsAddress.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

<a id="canonical-e5df3dafdb2986e74e6f87175b521760cb9f5a43558ac893c5ca527896b28976"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / d3574863ec07 / 3

<a id="canonical-84c4587130da6847221ae7e2400b16bc8fdc9e15489a58a2d74b799084fda149"></a>

<a id="canonical-08df456380366d6aa1de5deb92112ff2781253fce67e72e6afc42cbbe80fa3dd"></a>

## configured_address property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / d3574863ec07 / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

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

- [first_address](data-sources--voltstack_site--reference--group-004.md#canonical-8181be3b91245ee14615ce082b95ac8bb43e3853e1685aae6bbaec7c740c010b): complete subsection reference.

- [last_address](data-sources--voltstack_site--reference--group-004.md#canonical-bcea7906fe85054b080467e2f0ae913af4bca97b97aef1060b68f2ebb5f8b222): complete subsection reference.

<a id="canonical-1ac94480647d3b0cfb953b3b1d6a1327ac90831edae6b3a7647b8a9bdb51f453"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / d3574863ec07 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--voltstack_site--reference--group-004.md#canonical-8181be3b91245ee14615ce082b95ac8bb43e3853e1685aae6bbaec7c740c010b)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--voltstack_site--reference--group-004.md#canonical-bcea7906fe85054b080467e2f0ae913af4bca97b97aef1060b68f2ebb5f8b222)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-004.md#canonical-e5fbc479fbd4e8cf7af196b76a792b3d1b9254643b90cf5279fc982873f2deb6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-8181be3b91245ee14615ce082b95ac8bb43e3853e1685aae6bbaec7c740c010b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e1dce12af71e1e130af44728c758bce4bb1db2ff8d7f4b35626b86709064734"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 1ab64f1575dd / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-004.md#canonical-e5fbc479fbd4e8cf7af196b76a792b3d1b9254643b90cf5279fc982873f2deb6)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--voltstack_site--reference--group-004.md#canonical-b7e56ca5c228cd36f908fefc50eba0a614fabffdcc1a145d8b26e1df9df4d166)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-e5bb56e25f5c2dd0d16529846810549eccc43b1cdb2d95eb4fde036266e12f5f"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-e33258a4595f627bb99211f55fe36661665c99055bd9c889fe06c053820a4677"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 1ab64f1575dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4115230cdffd1d8a84ef812200d71f496b9f66cb165adaaa0dac2c85029c2c6"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 1ab64f1575dd / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--voltstack_site--reference--group-004.md#canonical-b7e56ca5c228cd36f908fefc50eba0a614fabffdcc1a145d8b26e1df9df4d166)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-bcea7906fe85054b080467e2f0ae913af4bca97b97aef1060b68f2ebb5f8b222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2747d2114a99ac18ac9e1aca550866b0454f07a8d31c0ff50dceccadf895e350"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / c49c65f793eb / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-004.md#canonical-e5fbc479fbd4e8cf7af196b76a792b3d1b9254643b90cf5279fc982873f2deb6)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--voltstack_site--reference--group-004.md#canonical-b7e56ca5c228cd36f908fefc50eba0a614fabffdcc1a145d8b26e1df9df4d166)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-3fd66f10aed530b36639809dbf2baedaa126f5341b98aa364556348e74f8e638"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-5dc5f1e451b5666428b71d752aebc0b6fbad531220aa09341a4b3196049cf45b"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / c49c65f793eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4549710d0335197c95db081a2757605f04e0451d47f79e7c19f9f5b802678c87"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / c49c65f793eb / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--voltstack_site--reference--group-004.md#canonical-b7e56ca5c228cd36f908fefc50eba0a614fabffdcc1a145d8b26e1df9df4d166)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f9d8501449e5c8f2a7b2df2251a0064a4a396329b048d17ed5b87103ce174ea"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / fab492c6eab5 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful

<a id="canonical-1e2d3f71b1a52ee7492d967371a63734f337ee982efd0bf5c243de14d5ebd342"></a>

Type: `"single"`. Computed.

DHCPIPV6 Stateful Server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-31373afd9b784ce728d88facc86787863ccec0fcf24e88c5a0719c1315911760"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / fab492c6eab5 / 3

- [automatic_from_end](data-sources--voltstack_site--reference--group-004.md#canonical-4dfc69924179c5fd48bbb9a68b3c46e7df7bbbb169da48dbb0bfdb173003b9b3): complete subsection reference.

- [automatic_from_start](data-sources--voltstack_site--reference--group-004.md#canonical-631b9209534fe10ad3cc2ee7a4b0cce81bcfb1b0ada107fa1b0cba10ce7c68d4): complete subsection reference.

- [dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-23c719f257397dbd5518a99e9ed841c7ccfebef5bfaf4802f9746d90f8c927d8): complete subsection reference.

<a id="canonical-f08561e9a3c8846c84b4857e5ed1e3c57d8dba3a075d7996ef2c885fe40e0991"></a>

<a id="canonical-79ee4d09b7615825280fb5740579c6589f7018c7de4a315d35250ee0ffe37335"></a>

## fixed_ip_map property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / fab492c6eab5 / 4

Type: `["map", "string"]`. Computed.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](data-sources--voltstack_site--reference--group-004.md#canonical-a503d1962ffb72926ba80a5988be58cd0d06310833d478ccc7e3a6f5bf5441c7): complete subsection reference.

<a id="canonical-1a048dd88419b3da105b261ad82fef6edb08de658ddde3e862ce998c0a3d2912"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / fab492c6eab5 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--voltstack_site--reference--group-004.md#canonical-4dfc69924179c5fd48bbb9a68b3c46e7df7bbbb169da48dbb0bfdb173003b9b3)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--voltstack_site--reference--group-004.md#canonical-631b9209534fe10ad3cc2ee7a4b0cce81bcfb1b0ada107fa1b0cba10ce7c68d4)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-23c719f257397dbd5518a99e9ed841c7ccfebef5bfaf4802f9746d90f8c927d8)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--voltstack_site--reference--group-004.md#canonical-a503d1962ffb72926ba80a5988be58cd0d06310833d478ccc7e3a6f5bf5441c7)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-4dfc69924179c5fd48bbb9a68b3c46e7df7bbbb169da48dbb0bfdb173003b9b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211ac3ca0982029134c793fe0dc09ee854124a58f68287e529f2d614c670f5f"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / a1233b4d7aae / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-004.md#canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-70b423e6ef96d12854ded0ff2dd83736c70581edc8ea8b5afb20038b5182b248"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

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

<a id="canonical-92f4cd4864d949c9cfd6750c257448e590a7aaff82c3f7f1806de496b08eafe6"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / a1233b4d7aae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0ccb2a7052eb6a38697b85364f6e10972b742d4abf96212151c181918d8c9c86"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / a1233b4d7aae / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-004.md#canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-631b9209534fe10ad3cc2ee7a4b0cce81bcfb1b0ada107fa1b0cba10ce7c68d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9a386292a09752fdeeb2d84eafd7abcd0af16aa5f4eae7d9731cf2e59fd5a58"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 9c1d41458602 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-004.md#canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-5c5448cda08bde3eedc81b8ecc886ccab862fb2a59fff6de1a1129c1ef55ab74"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

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

<a id="canonical-5ee6012218bf0a6febedb28b30dcfc38e9587a90ac16ffb50161fd1691ed65e9"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 9c1d41458602 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5330850232fb5180a33837f0baf2f66fc24f1bf980ee7b142a55ae66b1987705"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 9c1d41458602 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-004.md#canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-23c719f257397dbd5518a99e9ed841c7ccfebef5bfaf4802f9746d90f8c927d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5fd8de1f7c9540d549bf9ccf8ee93e26252e8026c78283f80845bea7094a5e8"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / ff3a759ec51f / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-004.md#canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-ac589f3824f730b4b73e1b8d28770b5f595010237ee6bc104e892b880ec19de8"></a>

Type: `"list"`. Computed.

List of networks from which DHCP server can allocate IP addresses.

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

<a id="canonical-93f6d5c5ab67c8f620b48520c1e10481f700f4fc3c61d8bb698f565c4e2e7856"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / ff3a759ec51f / 3

<a id="canonical-1283fecfa317effc735806f0137937b36530432c2873ce294aca65d617ea18cb"></a>

<a id="canonical-fb3f2aaea64a17f284085332da0d7678b276b85e73672c1b9734fead9006355f"></a>

## network_prefix property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / ff3a759ec51f / 4

Type: `"string"`. Computed.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Upstream description:

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-e326f303314965d070f7989df6622d2694dcd59d8977ac9314174ee26d4f0026"></a>

<a id="canonical-998b0bd37e8b66db824bd3e71edd83784ebc52ac5984ebae08476ca58b788bfc"></a>

## pool_settings property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / ff3a759ec51f / 5

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

- [pools](data-sources--voltstack_site--reference--group-004.md#canonical-c46d8431c3fcd05b1479a785684535d6ede35aea7bfa7d5daa5804ee0d141d9d): complete subsection reference.

<a id="canonical-4216addde2fafee3618d87ce995440ed30078896cf1e5d4aefb24b4a9fad9001"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / ff3a759ec51f / 6

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--voltstack_site--reference--group-004.md#canonical-c46d8431c3fcd05b1479a785684535d6ede35aea7bfa7d5daa5804ee0d141d9d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-004.md#canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-c46d8431c3fcd05b1479a785684535d6ede35aea7bfa7d5daa5804ee0d141d9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38e127ca38bbe7a8036aed40d35425893b0d20778eae915c26b21d68b4b2560e"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / d2a47c2fbc39 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-004.md#canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-23c719f257397dbd5518a99e9ed841c7ccfebef5bfaf4802f9746d90f8c927d8)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-347749a47dcff0427a2ea31a09cc1727f3e53ce76f4456166a257af4ee7bf4ad"></a>

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

<a id="canonical-4f01eef1ab5712d84d6f13df2565d697d62dc6d8b6a6ca382255bb28a2981621"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / d2a47c2fbc39 / 3

<a id="canonical-67047c35a75185f8518748e00911ab2ea70418b03bc0518fb4e0a279401ad854"></a>

<a id="canonical-52cc8405ccdcd87fe52d1a7c612b9a737039e1f959c814c4f25e76840c70b990"></a>

## end_ip property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / d2a47c2fbc39 / 4

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

<a id="canonical-9b281809c468fc9e074c6163539ce357bb3fc85e3d54239cfa49352a3e6f341e"></a>

<a id="canonical-f866f6c27feeee81a27199bf7c4e121924a644be7c06fa830aa406e1b3802f23"></a>

## start_ip property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / d2a47c2fbc39 / 5

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

<a id="canonical-c905aa63f85b0b375be6ae4c91803bc03e7ae7700926a87d0a2fe252ba41bcd6"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / d2a47c2fbc39 / 6

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--voltstack_site--reference--group-004.md#canonical-23c719f257397dbd5518a99e9ed841c7ccfebef5bfaf4802f9746d90f8c927d8)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-a503d1962ffb72926ba80a5988be58cd0d06310833d478ccc7e3a6f5bf5441c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-319ad32f74a352a2c8d539d221c16b3304f630e38bb66882698f6cdfff97fb40"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 5747418da8a2 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-004.md#canonical-e4e95ffbd93833861b4759fa7f1d4365ff3d73a7493b1664c597530418009d5d)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-004.md#canonical-70fd6140dd7ab7fc660a9b802ee25eba4dedb9c272f01a5a3117bf1b056e8f86)
- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-004.md#canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6)
- custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-d2ce3595148dafa09f1cc72b056b872107b561b0f507f25905e3aef9539a10aa"></a>

Type: `"single"`. Computed.

Map of Interface IPv6 assignments per node.

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

<a id="canonical-2ff7fcd77dcace1c590d8dcdb24e930f8b1199f8d73a20f8a4dd32418afc96db"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 5747418da8a2 / 3

<a id="canonical-72ba1b6972ac44c9f1ffbb2acf8474ea82d993cc7dd474aa76e735599160b814"></a>

<a id="canonical-bd4661d9c315fce0cf5cf7bb8086575132fd36eb7fffc0ab7a881a48be3e74b4"></a>

## interface_ip_map property — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 5747418da8a2 / 4

Type: `["map", "string"]`. Computed.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

<a id="canonical-61cb14075ea959576093c35cdba449a3c238ebca314a39fe53eb65c00309362b"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_con / 5747418da8a2 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-004.md#canonical-a8cda2a3a75e9bf1ebfbd18a9edd18fe6a8bf613fd80aaf2b20fd550ce2e63b6)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-f37a6ce3b0b1dcf8a5ed88f333db9df8534c9a764ac4eada09e27d37670c612d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b5b97188b548bc4106fdb0b9de1f1d4140789bcf73b6d5193b66bbc04b2ce26"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.is_primary — custom_network_config.interface_list.interfaces.ethernet_interface.is_primary / e9449684a43b / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.is_primary

<a id="canonical-62dc1b62a71701a5360af909c79316cc5a3d79c18ec82da54fd2f123fa7d3bea"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-d982af8a1bc67a824d14e4cd6834f04e4a76c4edbc6bf72ed33b93e93fc62d59"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.is_primary / e9449684a43b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-736e5349c82b7c146a7697e2fb87dd34067232517b795598c3aee0238de90ccb"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.is_primary / e9449684a43b / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-722fd5061622bbfa32a2862c46789a3fe5ab8054407c5c06a5594ecc71f59ef9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-03ecfd3725325adcb16fb20ce0c4cd637581d6428de1bcaa365da437a7ee514c"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.monitor — custom_network_config.interface_list.interfaces.ethernet_interface.monitor / 0ee2fb43fcd7 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.monitor

<a id="canonical-410622e0ee647181d289bfc64c39f9f97c15fd32779acfca9568705625cfcdf5"></a>

Type: `["object", {}]`. Computed.

Link Quality Monitoring configuration for a network interface.

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

<a id="canonical-c9fecea9eefefef9f0aa6c73e29a81e00ee046c1eefe237643527a9a10018612"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.monitor / 0ee2fb43fcd7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78000e611995f06cfa1cade94916e07cf676014393bdcc8091e980ce2b4bb066"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.monitor / 0ee2fb43fcd7 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-4c9bb452ce61575f024a64fe9e81588f89936f9ebd18ea9ba158ad058954abd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f627408f1ddd2029b4ef1a6cc3e968eff4b0f3cc6fdc5aff723bd20f3a668a1c"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled — custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disab / d0d2199f8708 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled

<a id="canonical-0cbfd7a3c4870c4202210f47a78222b3d0ddd8598527298cc119fb9f5157dd9c"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-c476f4eb439351b41fb40655ab48915263741accee74d24ab01d2a74e0eae1d7"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disab / d0d2199f8708 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ffce7b18ecf9b93341a36b6f2c4fb81a75a45c13845ac81cf465c80054420e72"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disab / d0d2199f8708 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-0beaeffb632695a915bf472c75c289e334967ee05d193017c43c51aff5d64daf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-166e900e804f5400d0d9040f9afff35928c1bfc2efa5da2af26d0291d4cd33f8"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address — custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_addre / 7f5007b9176e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address

<a id="canonical-87977cda72ec4405af97a60f700107a0828e48301135c30e085f5aa39d8311e4"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-c48ec8f8d1587bd39b77b8b6a92da66248bc37d63b4f148c95f51f795d26409f"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_addre / 7f5007b9176e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b1515dabeccf6f5bbd49f3470f8ea7f7e4057a820b11e177be0d4bac5769c898"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_addre / 7f5007b9176e / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-dbd631947f6a6f6025e279a93a458e07351c7e792c755cc6b5173822c153abda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f56bf709537fb10d82a566fb8ef9dc29cc07e3b2dafd1e6362fcfeb294bab8f2"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.not_primary — custom_network_config.interface_list.interfaces.ethernet_interface.not_primary / 227266f8bb24 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.not_primary

<a id="canonical-a4489b9611d5642ec89d4c445565ed7898d8cec88c0823202226966b830095bd"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for not primary.

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

<a id="canonical-02ba6cb9a48147abc2871ffd70941f08de97c16fb5e77864802f5b0ff0eb4b1d"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.not_primary / 227266f8bb24 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3bae3aef70b19010ec4f8c5a77c8eb6858ab980dcb181065d5fad2045bc4ebe9"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.not_primary / 227266f8bb24 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-e84c3d8440cb228032c8e44f9267edc79d571543ec6840de73f471a05e941053"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b523d71baeb8acb3d12b5757d8de6d6c419d8535986609eda7b2dafddb46187a"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_in / 0afa5d27c03e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network

<a id="canonical-a7f9d1ac5c9e0de00f1cac7a0c1cee75d4bd78c00e75fb31470cec7d56afaaa8"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-ed3500acdec4bc711a7f2cf8bd203d784141643cad008df18db3657e0d12094d"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_in / 0afa5d27c03e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8f15dc129343de25de232f7bd9e20e31c3fb08e6d879e0860230f9f120592541"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_in / 0afa5d27c03e / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-e18052ee905c48af36905429cceee8e2f06154b57592b7eb356326e1ec27b3ed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69f46e867015722de6d078aa9094e78fe221d45056e9e53bafcbea4b11e075f4"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_ne / c8cb32f818ef / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network

<a id="canonical-102b6f29f86f7f41a2262fd596635e943be576e9cc1e066356c4cfe8fdb815e2"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-1897c386ffa5aaaaebfbe7210254bd6e773abf946a7c42df67a0d801969447aa"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_ne / c8cb32f818ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fcdbc9be43aff759154ec7c16beea417cbe5867fbbe9146caa56b65d1978913a"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.site_local_ne / c8cb32f818ef / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-f7d391cedb260d2bc6a620595478beaf697aec05c7ffd6768fd71c720022f365"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9a864b63a93879a0bb958594297ea62bfdf6ed28faa4f8d191ea667f17400654"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip / 2f9716b43973 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip

<a id="canonical-09b4fbdf43279b0cc895969007156ffcd2ddc1ae77d0a9c471047d94eeb6a0e0"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

<a id="canonical-26237d981849c477e4571eab53056ec174210f8e27bb277e0dd9416086271c79"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip / 2f9716b43973 / 3

- [cluster_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-3df795f7b74667a75c9a5f794f4d9f508244927a5bc521913667917e72479737): complete subsection reference.

- [node_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-d847826bd08d81ccc559ac92e0d9dd942c6aa898546214dda7446f3d5b43d6d2): complete subsection reference.

<a id="canonical-0c78a701ccbf908beb40cedb095f480d3a543bd679921b927f70ae76e5c17639"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip / 2f9716b43973 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-3df795f7b74667a75c9a5f794f4d9f508244927a5bc521913667917e72479737)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-d847826bd08d81ccc559ac92e0d9dd942c6aa898546214dda7446f3d5b43d6d2)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-3df795f7b74667a75c9a5f794f4d9f508244927a5bc521913667917e72479737"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba108919730361aa95e0716e9ab28bd3e46e090debc8d482e5f98abfd2f0475e"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.clu / d9c57412c403 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-f7d391cedb260d2bc6a620595478beaf697aec05c7ffd6768fd71c720022f365)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip

<a id="canonical-e9d3b6e118647516bf56b714951e14d78f4656706ae9824f66be684b9fe5c4fc"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for cluster.

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

<a id="canonical-83a1eadc9db33a45d0717501395a7a52038192b4075e2363e3286660a7a0233c"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.clu / d9c57412c403 / 3

<a id="canonical-1b3780a6c35599576c01785bdc402e5658ef69af907c5dc86cd6133dfa17e370"></a>

<a id="canonical-6dce7729ae8457e26646b5db3162b7a6960a680418bf3a8f848d56394ec5fb99"></a>

## interface_ip_map property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.clu / d9c57412c403 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-8e9ecfcb454520cf0271c1757fe76dff0bd8d4021f2fb14e3b18e92ccf01bb9b"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.clu / d9c57412c403 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-f7d391cedb260d2bc6a620595478beaf697aec05c7ffd6768fd71c720022f365)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d847826bd08d81ccc559ac92e0d9dd942c6aa898546214dda7446f3d5b43d6d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c1f24722e362e5a0031bb4f622e8a897997d604697eadcd59579f46e575ebee"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / d2b2cc7eae79 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-f7d391cedb260d2bc6a620595478beaf697aec05c7ffd6768fd71c720022f365)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip

<a id="canonical-87c0e2e2f1dd2f7b63d6f2a53357c48cf15f546d58db7a418d6ff26e6560f37a"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-c0c4f59ac6b628506cffcec362074e0fb86263715a12fbcef3e9b8cc8d87d281"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / d2b2cc7eae79 / 3

<a id="canonical-6ad48ac75c4cc3073493db09f4aa37ec12eb6f4721cde94561914cabce04294d"></a>

<a id="canonical-26219273742f5bb04731b77cc9e07b9a900b4d82abca75d6baed2f9fe88a03b6"></a>

## default_gw property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / d2b2cc7eae79 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-c06e8b439a9813c74ca09d1de795d202d319c59332d511aec3a70b5e037658da"></a>

<a id="canonical-c60ff081e224695dd27d39437cbf1db25fea26e2035463755b1e73c308903c4b"></a>

## dns_server property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / d2b2cc7eae79 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-9e0101c8868c836caf510853086e18e57f68188855b1bb777530770d460d7cbb"></a>

<a id="canonical-0e6bc4bd15ae00cd6e68bf80cc9cb7ea5152795982d9a39c0e72665e803eeec9"></a>

## ip_address property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / d2b2cc7eae79 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-5fe6ff12dfc240afaaec6df6bdd3fe6138d3f3c5a1192742a6a7f61d7fc13ec6"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.nod / d2b2cc7eae79 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-f7d391cedb260d2bc6a620595478beaf697aec05c7ffd6768fd71c720022f365)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-8e500794f968dd3478f8685bd436754bcb8374aac74f8b2a0c571ca73d8126f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b16a0498ebf18864458050462362d249728bd643487283bbcab9026d407c4a8"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 5fc7503b72af / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address

<a id="canonical-46b27d96753d1cc1d276cda00edb2cfae3b016469948c678796ee98e5cf801f9"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

<a id="canonical-0fd41726583b6ea75a0efc7a2caf6ff66b40b45a1ca29d327720cb1dea508649"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 5fc7503b72af / 3

- [cluster_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-48d2ac006ddc753d5e34237aec3564e5195fdc53a6eeee81c5ede9c02d5dd7ec): complete subsection reference.

- [node_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-b4fc89f6c4df6d6e6e5a5318fe3841c7f7238a39ca50f066d71a567094185abe): complete subsection reference.

<a id="canonical-d3b7e4d2f30e36dc303769f180bc286d096c3055e20f466cd9a314d658ffe8b4"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 5fc7503b72af / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-48d2ac006ddc753d5e34237aec3564e5195fdc53a6eeee81c5ede9c02d5dd7ec)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-b4fc89f6c4df6d6e6e5a5318fe3841c7f7238a39ca50f066d71a567094185abe)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-48d2ac006ddc753d5e34237aec3564e5195fdc53a6eeee81c5ede9c02d5dd7ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60be37fdcd4f235b4e57535be098f87971bfb15030be68e7e1c171fc478ea01d"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 478a6479e2d5 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--voltstack_site--reference--group-004.md#canonical-8e500794f968dd3478f8685bd436754bcb8374aac74f8b2a0c571ca73d8126f5)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip

<a id="canonical-e12be3dff3f46bba160fd6b701cb8bc6a98fc4f8904441c224b027c6023e63d0"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for cluster.

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

<a id="canonical-d2699d2b512b1547db3632569692d12ba09dff0af0f75cfac180e532afb84322"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 478a6479e2d5 / 3

<a id="canonical-5a6401b45b371c4ab424cda4d070116523c68c67b49d4f6ab79e962ebd305707"></a>

<a id="canonical-3f1dc446ecd0e270fe580ac63f98348a16f3aae311eff1da1a33af5ad39a69fb"></a>

## interface_ip_map property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 478a6479e2d5 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-9c156936a9073e8d7daa5d821aa631a9a01a5f6694f9e600142966b774fe3cc3"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 478a6479e2d5 / 5

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--voltstack_site--reference--group-004.md#canonical-8e500794f968dd3478f8685bd436754bcb8374aac74f8b2a0c571ca73d8126f5)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-b4fc89f6c4df6d6e6e5a5318fe3841c7f7238a39ca50f066d71a567094185abe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-411fd4f1a66486f0631a5b56a3fba98aa7c40b75035052428ddec1bcf00c0775"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 040f6f1971c3 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--voltstack_site--reference--group-004.md#canonical-8e500794f968dd3478f8685bd436754bcb8374aac74f8b2a0c571ca73d8126f5)
- custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip

<a id="canonical-32eab6bdbc066468c3d731883645223cbcc0807ce3388d7e181775dca301aea9"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-4815ac7dfd992d48bca0af406a5d5485fde2ab36ec6cd83f3376b95df432a07e"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 040f6f1971c3 / 3

<a id="canonical-ca34401af4ca14da92762bab66ee0728b59ba2a7d67275bbdd3a1762acde4b0d"></a>

<a id="canonical-dfb6ae90f2a31e49921c29149f0b8099131776ad94d39e99fe86f6c885388f6e"></a>

## default_gw property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 040f6f1971c3 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-4b5c1a6d927f3d43230fdaa5a69d6a1fb03955cfad161e7541b89f1b5e4de4a1"></a>

<a id="canonical-0d2bf75aa54e072c27bb27796c3445fdf685397fdcd1adc390b3914d1393adc0"></a>

## dns_server property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 040f6f1971c3 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-f25c0c04331dbcf4abc6e63dd32b8f66ae0fa10bf1d938058927a9cf050949cf"></a>

<a id="canonical-00176556f8fbcdcf8bea92aabdb0113f07747cc6089afb2e533aa7c80b42aaa2"></a>

## ip_address property — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 040f6f1971c3 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-4713d22eff191911fa60602a7c5c8803df03a53241e5921ba8d04860fc5c6f65"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_a / 040f6f1971c3 / 7

- [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](data-sources--voltstack_site--reference--group-004.md#canonical-8e500794f968dd3478f8685bd436754bcb8374aac74f8b2a0c571ca73d8126f5)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-24bc6a3dab4c4e3bac978053ceee82adfa2526d2a613843e18c5c500c4f2b1cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-24053208ea14111dfaa569ad356733a56e717e45d78f455d52be38d2b2643e7d"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.storage_network — custom_network_config.interface_list.interfaces.ethernet_interface.storage_netwo / e23bd9880718 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.storage_network

<a id="canonical-8a3b946570684628ebef379827a98d497c706ab343589083145044065a9e2d85"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for storage network.

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

<a id="canonical-d7a78f37f47669b330c47df26028c40851dd6e5f6aa90199ef959f43ad0c15a7"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.storage_netwo / e23bd9880718 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e6e84b0046b248de2f0e943a402277a48e6a58a86944c6ae1be5fa9c126c752a"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.storage_netwo / e23bd9880718 / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-d10b14ef38495609e8af7019fbeb297076524db18ec8006ed616bab97d71dec2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6b3f1a538002f72826b00f7b666025f3b12193d64b7dce67b06ecaad0b0ca54"></a>

## custom_network_config.interface_list.interfaces.ethernet_interface.untagged — custom_network_config.interface_list.interfaces.ethernet_interface.untagged / 2b84780721cc / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- custom_network_config.interface_list.interfaces.ethernet_interface.untagged

<a id="canonical-7951e293ae39e3725e71ed927a03337f1b66092e379501919ade301c72b2a952"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-5c993d876f1793ca17f8afc1c65d3430faf33b092032b0d64b1c4899f9c39924"></a>

## Direct properties — custom_network_config.interface_list.interfaces.ethernet_interface.untagged / 2b84780721cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dc632bfe2df2ae6ca408012a1bef56fbdc224e07d9fa000f036e28da9d0c6b17"></a>

## Next pages — custom_network_config.interface_list.interfaces.ethernet_interface.untagged / 2b84780721cc / 4

- [custom_network_config.interface_list.interfaces.ethernet_interface](data-sources--voltstack_site--reference--group-003.md#canonical-a5953f3011c34394c09390764194fc59b5c55caf6f057a7461fba5b622d8b866)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-f9185e8bfdcf80d92514aa3038cdebe968f5a781c7a323cbc078b329d136e1df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b108fd3d1ce446196269d6d9f2f4bc57e06ab597b32390965174b248343c9c96"></a>

## custom_network_config.interface_list.interfaces.labels — custom_network_config.interface_list.interfaces.labels / 364dde712b1e / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- custom_network_config.interface_list.interfaces.labels

<a id="canonical-9cb87c74aff7e4c2364c342da38e7353b923b6fbe50c4103692b75804e663869"></a>

Type: `"single"`. Computed.

Add Labels for this Interface, these labels can be used in firewall policy.

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

<a id="canonical-d6a50c3954c0a9caaa8f7e7fe13175d988b50faa1277d2083aad8335bbca6ddf"></a>

## Direct properties — custom_network_config.interface_list.interfaces.labels / 364dde712b1e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b7bb84e9b74a91b2792e9a5316c5f08b7b78e046ee721fae6abece3c10a0855"></a>

## Next pages — custom_network_config.interface_list.interfaces.labels / 364dde712b1e / 4

- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-09a63c412db1b530f28538e6c58df1112e91261f41d69db544ff585ed2e3c180"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3ca97cc8d7c54de0a258ac316ee2d3341da08cb84a1e1c5d4b8cf42a6d4c765"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface — custom_network_config.interface_list.interfaces.tunnel_interface / 3b08b229a7a3 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- custom_network_config.interface_list.interfaces.tunnel_interface

<a id="canonical-7c6a70f22659605f11c30f6fd86546a8de3c81092f3e314f7d206e5c75e12b6c"></a>

Type: `"single"`. Computed.

Configuration parameter for tunnel interface.

Upstream description:

Tunnel Interface Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-node_choice": "[\"node\"]"
}
```

<a id="canonical-3a590fefd6088d37590079ec6efdbc1158cb1c0c7af0899ccf2a53435b179f4b"></a>

## Direct properties — custom_network_config.interface_list.interfaces.tunnel_interface / 3b08b229a7a3 / 3

<a id="canonical-b7cfb345d20bdac2c4853bc6449108b06441032b9c3d3709e7eedfbc4bee0e53"></a>

<a id="canonical-4fa718d803ff8c215419d3470600f4faa8813cc6e487769bd5c5be7d52ef3a1d"></a>

## mtu property — custom_network_config.interface_list.interfaces.tunnel_interface / 3b08b229a7a3 / 4

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 9000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

<a id="canonical-bda26aaa430c1d05046b254aece16377ed9be61fb3de9e7bd7b2179b76a308e8"></a>

<a id="canonical-ff8e9ebce94f517801a90f1bf40a6b124f8f195c35d0602395fde58647ac7f1c"></a>

## node property — custom_network_config.interface_list.interfaces.tunnel_interface / 3b08b229a7a3 / 5

Type: `"string"`. Computed.

Exclusive with \[\] Configuration will apply to a given device on the given node.

Upstream description:

Exclusive with \[\] Configuration will apply to a given device on the given node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-369779d8eff4bd8c4aad253efe921a81113decbc2aa979088c21d85243c8bcd3"></a>

<a id="canonical-202b918b173907d4b58fd6d330a1727795de357c6a6f614df4888b1323ee6093"></a>

## priority property — custom_network_config.interface_list.interfaces.tunnel_interface / 3b08b229a7a3 / 6

Type: `"number"`. Computed.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_local_inside_network](data-sources--voltstack_site--reference--group-004.md#canonical-035b8c5007332e92a9d7c7e99e9ab918df522adb16a44d20a1d38a9041e48b77): complete subsection reference.

- [site_local_network](data-sources--voltstack_site--reference--group-004.md#canonical-60a9a152f6a39b0f301e1961c7878e8d60cc8250533c153e5b8c99178bfba852): complete subsection reference.

- [static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-627dcf09b8a6e39c7ce779b26d58e9d406d9e805263da50bac1b09466d0b1804): complete subsection reference.

- [tunnel](data-sources--voltstack_site--reference--group-004.md#canonical-b334ce3dfe11f4c0e84e4114e387dd14c9194c8916699249e275338706d575e6): complete subsection reference.

<a id="canonical-57ac36daafdde106f0902b6b3ee57d1ef16e4dd059bcf5a1905b6439f4579d0b"></a>

## Next pages — custom_network_config.interface_list.interfaces.tunnel_interface / 3b08b229a7a3 / 7

- [custom_network_config.interface_list.interfaces.tunnel_interface.site_local_inside_network](data-sources--voltstack_site--reference--group-004.md#canonical-035b8c5007332e92a9d7c7e99e9ab918df522adb16a44d20a1d38a9041e48b77)
- [custom_network_config.interface_list.interfaces.tunnel_interface.site_local_network](data-sources--voltstack_site--reference--group-004.md#canonical-60a9a152f6a39b0f301e1961c7878e8d60cc8250533c153e5b8c99178bfba852)
- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-627dcf09b8a6e39c7ce779b26d58e9d406d9e805263da50bac1b09466d0b1804)
- [custom_network_config.interface_list.interfaces.tunnel_interface.tunnel](data-sources--voltstack_site--reference--group-004.md#canonical-b334ce3dfe11f4c0e84e4114e387dd14c9194c8916699249e275338706d575e6)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-035b8c5007332e92a9d7c7e99e9ab918df522adb16a44d20a1d38a9041e48b77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8120914978ad34d0dd2a14f9ef618bd9335a78f2b2cc33dde690fcedc1ca48c8"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.site_local_inside_network — custom_network_config.interface_list.interfaces.tunnel_interface.site_local_insi / 032833f1ba96 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--reference--group-004.md#canonical-09a63c412db1b530f28538e6c58df1112e91261f41d69db544ff585ed2e3c180)
- custom_network_config.interface_list.interfaces.tunnel_interface.site_local_inside_network

<a id="canonical-213b60338b276be4e7c275e598bdd6d977c0ef61e9bcbf5ec41080a4853a857b"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-12cd85d13adb4c9a0653a8ec123e31a1c912b4ccbe721daca5ca4f6be0e1b7f1"></a>

## Direct properties — custom_network_config.interface_list.interfaces.tunnel_interface.site_local_insi / 032833f1ba96 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e0918399e0c694ccebd8662dfe7729baac5c4df288e8adef0306082a597e5650"></a>

## Next pages — custom_network_config.interface_list.interfaces.tunnel_interface.site_local_insi / 032833f1ba96 / 4

- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--reference--group-004.md#canonical-09a63c412db1b530f28538e6c58df1112e91261f41d69db544ff585ed2e3c180)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-60a9a152f6a39b0f301e1961c7878e8d60cc8250533c153e5b8c99178bfba852"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3cc5fd7753cf8b6554ebb4126fca438eec499660bfd2f38e2bc11bf401f7e3f"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.site_local_network — custom_network_config.interface_list.interfaces.tunnel_interface.site_local_netw / a94f06420586 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--reference--group-004.md#canonical-09a63c412db1b530f28538e6c58df1112e91261f41d69db544ff585ed2e3c180)
- custom_network_config.interface_list.interfaces.tunnel_interface.site_local_network

<a id="canonical-a14a58d27b2ef5ecafcf65c87664b7f7c513a9737c63600a04b34fac253bd24f"></a>

Type: `["object", {}]`. Computed.

Enable this option

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

<a id="canonical-3500debb239204946eae8324de0250e708ffbc0cdb25a8f77ba69787275c68ab"></a>

## Direct properties — custom_network_config.interface_list.interfaces.tunnel_interface.site_local_netw / a94f06420586 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4d3a67e5b86c92b38702a5a98f81e8e764e0b55a4a1d3107e5aa102b2415c44c"></a>

## Next pages — custom_network_config.interface_list.interfaces.tunnel_interface.site_local_netw / a94f06420586 / 4

- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--reference--group-004.md#canonical-09a63c412db1b530f28538e6c58df1112e91261f41d69db544ff585ed2e3c180)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-627dcf09b8a6e39c7ce779b26d58e9d406d9e805263da50bac1b09466d0b1804"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60ecd31aab0ac8adb92a18404d6c60467c94a780df1a69d3a2028f2910745dbb"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.static_ip — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip / 07b6fa5906a0 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--reference--group-004.md#canonical-09a63c412db1b530f28538e6c58df1112e91261f41d69db544ff585ed2e3c180)
- custom_network_config.interface_list.interfaces.tunnel_interface.static_ip

<a id="canonical-fccd86318a0b16d054c114dcf0edcaac47de14ba3945cbe54ffd490949ff4378"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

<a id="canonical-818eee9f26238c25bd000489e1f185d1f0f936f865ce676478b1131301407fb5"></a>

## Direct properties — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip / 07b6fa5906a0 / 3

- [cluster_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-020768650d13149e00cad273bac98613b20b3b38043170983e6298ca66b0719c): complete subsection reference.

- [node_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-feb4113822b8b70d66653d57c4786b3653ddb4d650b80a44bac1360b86dcc3de): complete subsection reference.

<a id="canonical-9418005c61a3e7fc8e743d90565c8d01fd57a796a442180c5d184bff9bcb0b53"></a>

## Next pages — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip / 07b6fa5906a0 / 4

- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-020768650d13149e00cad273bac98613b20b3b38043170983e6298ca66b0719c)
- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-feb4113822b8b70d66653d57c4786b3653ddb4d650b80a44bac1360b86dcc3de)
- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--reference--group-004.md#canonical-09a63c412db1b530f28538e6c58df1112e91261f41d69db544ff585ed2e3c180)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-020768650d13149e00cad273bac98613b20b3b38043170983e6298ca66b0719c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b22f8dbd6260a4fde1e6d7bdc75ac1aa5b5f0a1984649b13597b88affe6cf950"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.clust / 82e96daa7065 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--reference--group-004.md#canonical-09a63c412db1b530f28538e6c58df1112e91261f41d69db544ff585ed2e3c180)
- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-627dcf09b8a6e39c7ce779b26d58e9d406d9e805263da50bac1b09466d0b1804)
- custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.cluster_static_ip

<a id="canonical-e029b0ad386ec9bf778cd153fad34751ecf4ef03517acbd54fc600b31b257913"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for cluster.

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

<a id="canonical-162b73f10ffd9b8547627c4b50f5ea019dc702c8eca4851fe80b526e6bb7b46b"></a>

## Direct properties — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.clust / 82e96daa7065 / 3

<a id="canonical-bec7807ce4782a8275a824f3883c8d64f96f2acc2b928a0cfb311ad9e8b370ee"></a>

<a id="canonical-c7991ebf7c224e46f7ca441ecbd59f1a553c3be3fe3cbab62e04863d3bce84ad"></a>

## interface_ip_map property — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.clust / 82e96daa7065 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-4b4370fadd3e15e4f26f5a7a308932fe02d138e9f694849d53bbe9ffb2a477a7"></a>

## Next pages — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.clust / 82e96daa7065 / 5

- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-627dcf09b8a6e39c7ce779b26d58e9d406d9e805263da50bac1b09466d0b1804)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-feb4113822b8b70d66653d57c4786b3653ddb4d650b80a44bac1360b86dcc3de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e05ce6b504e37e4aaffa5b80e7c6d366b9c9a962c5240ea7c8639cbec3dcec36"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_static_ip — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_ / 6c7a278aa865 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--reference--group-004.md#canonical-09a63c412db1b530f28538e6c58df1112e91261f41d69db544ff585ed2e3c180)
- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-627dcf09b8a6e39c7ce779b26d58e9d406d9e805263da50bac1b09466d0b1804)
- custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_static_ip

<a id="canonical-d6b6782baf41e5c1cd1ede6c4f29e76692f77c468817f76e0f24c51063d5a5fb"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-480218f1e97fdd37063eee11cb0ea8d587484cd356d84773aab6e905ecb341d4"></a>

## Direct properties — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_ / 6c7a278aa865 / 3

<a id="canonical-79f641ea250c61d77bf534bf7eac741d4df0f6512699eb71f0ff351a4c3dddc2"></a>

<a id="canonical-2586632cf724784dfe62cd6521f94b25403e012dd3460bb2352ece1c5e83be48"></a>

## default_gw property — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_ / 6c7a278aa865 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-78b9f3d9ae03218d068aa5191e7f300d4fee7ed4606ffe743d4ec346ffab0cc2"></a>

<a id="canonical-b258b667a91e62f29eb6850d0f9932dac52d03389a8b03f268e9b6b0a4106231"></a>

## dns_server property — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_ / 6c7a278aa865 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-9c9bbe6a49cf49bca7de9ba409075b2d89b5c7d747cc6ab4ee189080e8c508d2"></a>

<a id="canonical-e91ec65990b447c083cbacafef182d681c090a6ef8b253eddd0f266f5ebc746f"></a>

## ip_address property — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_ / 6c7a278aa865 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-a728d85639764563463734b5e68c3a82ade41324dc7e035a08c16608fecdaed7"></a>

## Next pages — custom_network_config.interface_list.interfaces.tunnel_interface.static_ip.node_ / 6c7a278aa865 / 7

- [custom_network_config.interface_list.interfaces.tunnel_interface.static_ip](data-sources--voltstack_site--reference--group-004.md#canonical-627dcf09b8a6e39c7ce779b26d58e9d406d9e805263da50bac1b09466d0b1804)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-b334ce3dfe11f4c0e84e4114e387dd14c9194c8916699249e275338706d575e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e064852869ed897ff3ce29e624c8c9b0d231d1c8a49e1ba110da20886add0b9"></a>

## custom_network_config.interface_list.interfaces.tunnel_interface.tunnel — custom_network_config.interface_list.interfaces.tunnel_interface.tunnel / 34911a992352 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [custom_network_config.interface_list](data-sources--voltstack_site--reference--group-003.md#canonical-668c2e17eac42d74f26ccb65d2541c362a467413fe81a79a085fc7498a418e8a)
- [custom_network_config.interface_list.interfaces](data-sources--voltstack_site--reference--group-003.md#canonical-84160d92af3a2eb60197656c48dbaf00bd1f4e3c95509f416a79a32fa3414988)
- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--reference--group-004.md#canonical-09a63c412db1b530f28538e6c58df1112e91261f41d69db544ff585ed2e3c180)
- custom_network_config.interface_list.interfaces.tunnel_interface.tunnel

<a id="canonical-da0c4daab4093653a217603b95a5b6e12b44309b15a54cdadfb42e37af908ed1"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-d0fcceb4c769aebbd08f25007901202f36e758edf355fe3fe6cfcbf901cf1f30"></a>

## Direct properties — custom_network_config.interface_list.interfaces.tunnel_interface.tunnel / 34911a992352 / 3

<a id="canonical-30fac4969c9419a0e107f6a29ef8256b5144a61d51d03ce62cd631909e1b5c83"></a>

<a id="canonical-a055a2ecd19549c919dc58cd56e7a682bfd6e65fdc3a4026a17b28fd841b2b0a"></a>

## name property — custom_network_config.interface_list.interfaces.tunnel_interface.tunnel / 34911a992352 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-04ac307101723d732b9cda1fb1664581022054d97ef95cbd790773bbaf03571e"></a>

<a id="canonical-deb20055282486fd88b484fb1b251a1781cc5737c1b280bcd48e2b3139ff92a5"></a>

## namespace property — custom_network_config.interface_list.interfaces.tunnel_interface.tunnel / 34911a992352 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-129c5c53df9e0885bb5661175b81d0f61589168ae244818b8e89fb4c4acf3b28"></a>

<a id="canonical-597ec7278bb6535908ce12ec62b100884bd99b1f3e16d27d3c3ff1430ea41b18"></a>

## tenant property — custom_network_config.interface_list.interfaces.tunnel_interface.tunnel / 34911a992352 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-8b3b0b12a58081e40c1f8193413ff89462b3d04e7a9e1ffe74e9c6f86a830315"></a>

## Next pages — custom_network_config.interface_list.interfaces.tunnel_interface.tunnel / 34911a992352 / 7

- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--reference--group-004.md#canonical-09a63c412db1b530f28538e6c58df1112e91261f41d69db544ff585ed2e3c180)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-8271b8567bb1abc2ae540e8bcd2fcfa06c31b5b46437809aeed4ccdd6755894e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd8c579ad835144a613fad624c53f737724d347fd1b0d759f4f3c93658562d17"></a>

## custom_network_config.no_forward_proxy — custom_network_config.no_forward_proxy / df2eaa811ebf / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- custom_network_config.no_forward_proxy

<a id="canonical-bbc82dc41c54b5dbfbd77c3066c13f54b9b31e5bc78449f66e2e562cf55f4ec6"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no forward proxy.

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

<a id="canonical-020cf535683e77d113ff8cbade7e0d036c05c7147d5fae1eb72731a915a4b46f"></a>

## Direct properties — custom_network_config.no_forward_proxy / df2eaa811ebf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c8e0e7287d29a64d17788178071dbdcdb018813e12e6771777169cf2d8456ff9"></a>

## Next pages — custom_network_config.no_forward_proxy / df2eaa811ebf / 4

- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-df7b86d5786dc9f197de01dfbe4b1dcdc376890720da90509b66eb630a2eb7c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5ffe0321704c803ecf8eb01beb12bc1ec3a7516a6a31d8276a484b70a8389ea"></a>

## custom_network_config.no_global_network — custom_network_config.no_global_network / f508ea67a128 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- custom_network_config.no_global_network

<a id="canonical-8961dbb44815e1809b05f999f3e7d079dea4ea197165a5ea76f2ca37f35014df"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no global network.

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

<a id="canonical-b95d9c3a027385cfad23bca7da6347751d22bf4946853b6fc010fd210301f304"></a>

## Direct properties — custom_network_config.no_global_network / f508ea67a128 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-71216563a447307a36e4f26c7b4fcfe95096ee3546e2e9782ecdcfc6493f58b7"></a>

## Next pages — custom_network_config.no_global_network / f508ea67a128 / 4

- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-fe8c62b2f0bb1f56673102972c2444f482eb56cb0be442d50e31a8752b06e476"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33cac704aa274db37c70187fff224fd21a5cecb7a711c49e6d953ae9e6332564"></a>

## custom_network_config.no_network_policy — custom_network_config.no_network_policy / 1bef653408f2 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- custom_network_config.no_network_policy

<a id="canonical-17b04913a9490898e38943033fda971824ea58b615fa40eb652c6c2bf63ca297"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-391cd99da9f064154ef653d3a50424da0709a165d4542b05396dd7a2dfd6c552"></a>

## Direct properties — custom_network_config.no_network_policy / 1bef653408f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c18a9d7d02c0634ca2c6fbde124f0651350f30b2e60582fe8ee93c17235088d4"></a>

## Next pages — custom_network_config.no_network_policy / 1bef653408f2 / 4

- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)

<a id="canonical-139aada80175683f19dc27480b10078540809b3f566fd29a948f3a3f927970bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1698089dec01b4a13b564eafd65362fa64b04efc3fb8bdf1ad01ffe94555e2a7"></a>

## custom_network_config.sli_config — custom_network_config.sli_config / 3fccf59402db / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-c525865a5846597c8c6439442678872b265504c5c370bcbfc927c71f5fd956c0)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-e470a441d6b912928efcdc07d60d4246ea07c1af02f2c30de1baf077d03bc94d)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-a2bbd12fb6e2064bcd5a324f010ef68909fc4d537defdf39eb8bd3c89b389061)
- custom_network_config.sli_config

<a id="canonical-00c87718bbdb2e2560deed15213e306d6b23dc20d0df193769db175c80140d92"></a>

Type: `"single"`. Computed.

Site local inside network configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

<a id="canonical-0ce97c3c798bfbdec7ebcb5912ee015234e45f5d01291dfac6393e8fdfe6acb5"></a>

## Direct properties — custom_network_config.sli_config / 3fccf59402db / 3

- [no_static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1ce109f60eec6558373fe6450ec154b179bcfa100f0fc7e76ac7cf16f8719d28): complete subsection reference.

- [no_v6_static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-69460bf97c43ab82da9f8774da4e1f3ce36f58fc15f3dbf405460500171b1589): complete subsection reference.

- [static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-358a9d469cb86323ddc241cd78e9cd8a7e6ac058936496e892a30ab6dbd4ad0e): complete subsection reference.

- [static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-f865c88ac10b4305abe16ef3931eb2c5bb0018e4c37f445d3614cb599dddea09): complete subsection reference.
