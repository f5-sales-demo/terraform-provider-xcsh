---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-2cb0d30f3f567e7e98e8206de6202c3a40336a7999bf2ab0c5488ad073286c05"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 155b4e229d37 / 3

- [interface](data-sources--aws_vpc_site--reference--group-005.md#canonical-dcb9573375cbcfc21fc5177cfbe2cb40a83d4c7dc8c84a7d50343453a7ae2048): complete subsection reference.

- [nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-f99762732f6eef7d9169b3989bda96e4e0e9b4d8bf04554e0e02bd02520892b3): complete subsection reference.

<a id="canonical-3e0f615565540c7168beee69f2deb0ea6d4ca4c46ce60809e6ffa33fd73641fb"></a>

<a id="canonical-c888613e962435b1767c3a088770b243e751cbc2c004534b8a51e171c49f4fa8"></a>

## type property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 155b4e229d37 / 4

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fd89d280c66e0987dbe80c08f9701e0e737b2047a776d3a5286cbda6efdf0307"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 155b4e229d37 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--aws_vpc_site--reference--group-005.md#canonical-dcb9573375cbcfc21fc5177cfbe2cb40a83d4c7dc8c84a7d50343453a7ae2048)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-f99762732f6eef7d9169b3989bda96e4e0e9b4d8bf04554e0e02bd02520892b3)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-dcb9573375cbcfc21fc5177cfbe2cb40a83d4c7dc8c84a7d50343453a7ae2048"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08f74db98e5c2e3aad759d8c98410d5b5fb4449b136e7b5c9b4a24596a1f8d93"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 53443d90092f / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-fc117b259a0186a106062d2990256669dddd1fe7a759a6558363ac6365334e6d)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-ae52f688a89ba596b353e1dea3f2dd583f69dcc81297a799512f50941f8d0e78)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-004.md#canonical-5d8ad31696ac3f3ca96a45c50fa7746b0b7a044a0ca13167a10b55ae923e1ada)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-a91516dc853dfb59f011abbb2367553a8e4d44cc4be0994c7b1cea88fefab9a0"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-302359c72cb8c0a5f6564c2119a54e2c4e654f077a4cf34c17e181f9b9496571"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 53443d90092f / 3

<a id="canonical-fc918de0222c5145b7cad5978a0ed7014036f808ef780c6aa6100d6ffd9e55b6"></a>

<a id="canonical-8bab48a3f9ae94598a19bb6dd6fabb50609cafd32a2196bc94b5dd02e22a3412"></a>

## kind property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 53443d90092f / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
  }
}
```

<a id="canonical-4f6d01539b122512445f4f9b45925e6541933fcc229e66c5de6475711e47df00"></a>

<a id="canonical-e0a84aba7ec127fbe9412f20947ff0d6c28c765bd8c6edf6e2c47fe1a50c1350"></a>

## name property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 53443d90092f / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
  }
}
```

<a id="canonical-2f7b15c4eb82208326e65d5b1e0c8df50ff069e1006195c146b325335177925c"></a>

<a id="canonical-a8a784b2df628e4a759f118e66b3bb21a4f793c42c1f0e0df08688428f800ec9"></a>

## namespace property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 53443d90092f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
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
  }
}
```

<a id="canonical-8075fd652d3794df0a83d7dbbf588be74b62bb1bfed4ed1865b50bf232cd740a"></a>

<a id="canonical-7fa706235203823405e1b7ea2498e1da58c2b4699cf7cb160c651059aa72e3d9"></a>

## tenant property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 53443d90092f / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
  }
}
```

<a id="canonical-c802f3840b787b89c22f0addb7d3978859fd6f7ce9b51d86d055c8ef874fff39"></a>

<a id="canonical-70d81ca3c6dfe870890493ebf69d77aeec19019359eff62785b3497f172c3352"></a>

## uid property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 53443d90092f / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
  }
}
```

<a id="canonical-7d5a93b60e3df58fce55d9a5e69a0b73b3a00866e5cb21c84196de8cc932d113"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 53443d90092f / 9

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-004.md#canonical-5d8ad31696ac3f3ca96a45c50fa7746b0b7a044a0ca13167a10b55ae923e1ada)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-f99762732f6eef7d9169b3989bda96e4e0e9b4d8bf04554e0e02bd02520892b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dadf7fb2be8d0a3d68d25843589628e915c4ac0186dd94c6a79bbe86088e51b"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / f6e23cb8cb8c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-fc117b259a0186a106062d2990256669dddd1fe7a759a6558363ac6365334e6d)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-ae52f688a89ba596b353e1dea3f2dd583f69dcc81297a799512f50941f8d0e78)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-004.md#canonical-5d8ad31696ac3f3ca96a45c50fa7746b0b7a044a0ca13167a10b55ae923e1ada)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-cc37e2af2764962d50e2e85a48e6f850eecdd6450a458a999de463e98f9ce40c"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-24f6a082f6219d15234a483e9ef33a91eda7e7fb944f547a5744569a2aff2798"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / f6e23cb8cb8c / 3

- [dual_stack](data-sources--aws_vpc_site--reference--group-005.md#canonical-25ffb17f07d98630bd07a6b217e1eef2907f630f30a6e297148562fc77dd910b): complete subsection reference.

- [ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-db54f8460f8fb6dc11917ab0351c278660b32b8bcd6ade78f4e8f55a58483d57): complete subsection reference.

- [ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-5b1744119da5c4685da12215e07f5b40a7822d516779f1b1e630bec2ea5067e0): complete subsection reference.

<a id="canonical-ffbca121fbbedecb3a46b6b1d1fc5a6a88b18519990daaa7a684cab2b1d70c96"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / f6e23cb8cb8c / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-005.md#canonical-25ffb17f07d98630bd07a6b217e1eef2907f630f30a6e297148562fc77dd910b)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-db54f8460f8fb6dc11917ab0351c278660b32b8bcd6ade78f4e8f55a58483d57)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-5b1744119da5c4685da12215e07f5b40a7822d516779f1b1e630bec2ea5067e0)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-004.md#canonical-5d8ad31696ac3f3ca96a45c50fa7746b0b7a044a0ca13167a10b55ae923e1ada)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-25ffb17f07d98630bd07a6b217e1eef2907f630f30a6e297148562fc77dd910b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd0becac71b1ef1733f44f97026c5c4e821f14da1513461c8ba30a1dca1c7816"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 2d4d1c272d1a / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-fc117b259a0186a106062d2990256669dddd1fe7a759a6558363ac6365334e6d)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-ae52f688a89ba596b353e1dea3f2dd583f69dcc81297a799512f50941f8d0e78)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-004.md#canonical-5d8ad31696ac3f3ca96a45c50fa7746b0b7a044a0ca13167a10b55ae923e1ada)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-f99762732f6eef7d9169b3989bda96e4e0e9b4d8bf04554e0e02bd02520892b3)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-cb49e1d2b69bfdf214b7a47814d28f75c8136b2e6fc3ebe9ceafaa22d2f8b86f"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

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

<a id="canonical-148a726635708d069468ca7217cf7d6cb4a70d322e6a4dcd00f31dcc00f65497"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 2d4d1c272d1a / 3

- [ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-8fc9a28dba31ef55f1e4e1b8171b40158f44dd8199fea4f69bc0d5ff87c49f8b): complete subsection reference.

- [ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-6287e5ad2c1914a30c696ff5cbd184d63c1b37f57147dce504b62489a46ab5d2): complete subsection reference.

<a id="canonical-ffeaf102c0d1270f47ec2484b9c8a8f514c300b47e52d53e2040dea4a07ae21d"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 2d4d1c272d1a / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-8fc9a28dba31ef55f1e4e1b8171b40158f44dd8199fea4f69bc0d5ff87c49f8b)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-6287e5ad2c1914a30c696ff5cbd184d63c1b37f57147dce504b62489a46ab5d2)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-f99762732f6eef7d9169b3989bda96e4e0e9b4d8bf04554e0e02bd02520892b3)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-8fc9a28dba31ef55f1e4e1b8171b40158f44dd8199fea4f69bc0d5ff87c49f8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42eae203f03a86b330ce67111e2bd974d032d6dcd2efee490c5fbf8d7686d111"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 00add8a90ec7 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-fc117b259a0186a106062d2990256669dddd1fe7a759a6558363ac6365334e6d)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-ae52f688a89ba596b353e1dea3f2dd583f69dcc81297a799512f50941f8d0e78)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-004.md#canonical-5d8ad31696ac3f3ca96a45c50fa7746b0b7a044a0ca13167a10b55ae923e1ada)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-f99762732f6eef7d9169b3989bda96e4e0e9b4d8bf04554e0e02bd02520892b3)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-005.md#canonical-25ffb17f07d98630bd07a6b217e1eef2907f630f30a6e297148562fc77dd910b)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-8881e421a8ce6024aae66fab3adf4177007b8e1b0b35df9b6ab26650a9210bf7"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-5a93bd19ac950f1f2fd9e3cde85d47a61ba029b3c50a5333cc1ffedfad7e7a17"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 00add8a90ec7 / 3

<a id="canonical-bdb3a34e171eeaa0b516b5e7a9b2564d751c4074638155b3f3215b31877ca1fb"></a>

<a id="canonical-fc3602fa5d5be7004e1fe2178914c7cbd8205a056b4ead83af24c95c9f8f0e29"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 00add8a90ec7 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-70fa7ed51a60e83134bc09e1543775493c179970ca4b8a4cf31fc5bf6353af61"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 00add8a90ec7 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-005.md#canonical-25ffb17f07d98630bd07a6b217e1eef2907f630f30a6e297148562fc77dd910b)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-6287e5ad2c1914a30c696ff5cbd184d63c1b37f57147dce504b62489a46ab5d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57e4c8c2af5cec23391762ab171f26642f19aa034ebbe471658568a0b9e0b81b"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9272ced1ba6a / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-fc117b259a0186a106062d2990256669dddd1fe7a759a6558363ac6365334e6d)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-ae52f688a89ba596b353e1dea3f2dd583f69dcc81297a799512f50941f8d0e78)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-004.md#canonical-5d8ad31696ac3f3ca96a45c50fa7746b0b7a044a0ca13167a10b55ae923e1ada)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-f99762732f6eef7d9169b3989bda96e4e0e9b4d8bf04554e0e02bd02520892b3)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-005.md#canonical-25ffb17f07d98630bd07a6b217e1eef2907f630f30a6e297148562fc77dd910b)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-57af923f27067f0bef1b643a163d04ee79608fa8a3bc8d4c8509a174f8e2f5a1"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="canonical-c58839c1734ef80c24dc1f2d5a5c66d1212ca6a5a950cb5e84bc6aeace90eeac"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9272ced1ba6a / 3

<a id="canonical-b36ff329768a100c6f9d5b671a74428e2dac6539f556bace573a88de4b7a31b5"></a>

<a id="canonical-b2206b76a091f8e98b60c575aa8d213fa1ade2386a137c87fcba6e9405a60908"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9272ced1ba6a / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-d498c1dfb7859907f721f494293ca363513f3f74187b0b2b6c89f7f15136e32f"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9272ced1ba6a / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_vpc_site--reference--group-005.md#canonical-25ffb17f07d98630bd07a6b217e1eef2907f630f30a6e297148562fc77dd910b)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-db54f8460f8fb6dc11917ab0351c278660b32b8bcd6ade78f4e8f55a58483d57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afe930a5c83b68de4b562cb1965ba21e9964714c9b2a6e7e231c2bff3d37bbba"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 8891f31d8205 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-fc117b259a0186a106062d2990256669dddd1fe7a759a6558363ac6365334e6d)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-ae52f688a89ba596b353e1dea3f2dd583f69dcc81297a799512f50941f8d0e78)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-004.md#canonical-5d8ad31696ac3f3ca96a45c50fa7746b0b7a044a0ca13167a10b55ae923e1ada)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-f99762732f6eef7d9169b3989bda96e4e0e9b4d8bf04554e0e02bd02520892b3)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-77229533f077d8a4a4cafb86f8cb9b1a35b48b7add5f43d9841d55d4f4168702"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-151a66dca6756d8e01f6767fb0aa95f0ae0005db14bd23b3af5ea09eb2f56aa7"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 8891f31d8205 / 3

<a id="canonical-98d9814efe7991b926193901edfdbc731cf7c4e7f61fd954b2ccb343640beffd"></a>

<a id="canonical-721c0d2f65ee18cc2edd23366d2a56779c651aa4eadbee01580343514542f909"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 8891f31d8205 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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

<a id="canonical-e3f6cdfcaaa36fed267a5e21490ed707c2af938bb68643652c951aec9d61587d"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 8891f31d8205 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-f99762732f6eef7d9169b3989bda96e4e0e9b4d8bf04554e0e02bd02520892b3)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-5b1744119da5c4685da12215e07f5b40a7822d516779f1b1e630bec2ea5067e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea7c4c3de35a6f4a5b10228bf302f2691c930afc5348c4e4b31aca527554d2c1"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 12c202715e6d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-fc117b259a0186a106062d2990256669dddd1fe7a759a6558363ac6365334e6d)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-ae52f688a89ba596b353e1dea3f2dd583f69dcc81297a799512f50941f8d0e78)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_vpc_site--reference--group-004.md#canonical-5d8ad31696ac3f3ca96a45c50fa7746b0b7a044a0ca13167a10b55ae923e1ada)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-f99762732f6eef7d9169b3989bda96e4e0e9b4d8bf04554e0e02bd02520892b3)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-b1f65d5acb898a25f831b1f33a63aa9fe2fd7b2dff9b255ea266908760b22cc1"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

<a id="canonical-996d4471e35333b6046bb4e067e742f8434585b759d1adde9fef4e2e6fd0feba"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 12c202715e6d / 3

<a id="canonical-30cf87d4b52c20b594145d1aa157623ef4482ea5f664925dcea5d62c5f3ebee9"></a>

<a id="canonical-2c02ea7da7a6672100b34dda8d4683f3da89ae607fc30d73f57644263039967f"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 12c202715e6d / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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

<a id="canonical-7aea7de3c45b934b917b39aaad5d4f58670cd7ac833278744bd3cc95a6e95774"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 12c202715e6d / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_vpc_site--reference--group-005.md#canonical-f99762732f6eef7d9169b3989bda96e4e0e9b4d8bf04554e0e02bd02520892b3)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-1c2fcf1ec9d797aa24c214c65d86001e3db189de3f6cf941bd6f123d6af37ec0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3eef6eeae6c6c67eebcf641ee64af6ceccb45b716b6dc0980dacb8fe223e1193"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 6de2a9292f23 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-fc117b259a0186a106062d2990256669dddd1fe7a759a6558363ac6365334e6d)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-ae52f688a89ba596b353e1dea3f2dd583f69dcc81297a799512f50941f8d0e78)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-672f39e85c9fdf072317aca3f40a7abef1d9a27e00a02597e5946ef64497cecf"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-5d2804cfd27fc6003ac55646f40418156a045a120b24d951ab1ab1f218a58723"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 6de2a9292f23 / 3

- [ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-d8d4b276cdfaa97003e2f1ebcc429fe9e8b1db6183bc31867aea55e5dd835c13): complete subsection reference.

- [ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-8f2051106fca301fd54fafd97dbb7042dfba5314203958b1f40f69bad45607f5): complete subsection reference.

<a id="canonical-3c7f54da94f75943e42411331f1a846d32ce563acce80cceb1e75d258edbaada"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 6de2a9292f23 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--aws_vpc_site--reference--group-005.md#canonical-d8d4b276cdfaa97003e2f1ebcc429fe9e8b1db6183bc31867aea55e5dd835c13)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--aws_vpc_site--reference--group-005.md#canonical-8f2051106fca301fd54fafd97dbb7042dfba5314203958b1f40f69bad45607f5)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-d8d4b276cdfaa97003e2f1ebcc429fe9e8b1db6183bc31867aea55e5dd835c13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6aa9a95b3eb83b54f6b88c032355f24a9def1e463cdf635664341b26b8285b3e"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / ee33baa6bbd7 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-fc117b259a0186a106062d2990256669dddd1fe7a759a6558363ac6365334e6d)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-ae52f688a89ba596b353e1dea3f2dd583f69dcc81297a799512f50941f8d0e78)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-005.md#canonical-1c2fcf1ec9d797aa24c214c65d86001e3db189de3f6cf941bd6f123d6af37ec0)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-ac12620670a209f5166ce200322c16e9d6af5738c30787951e114480daa626da"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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

<a id="canonical-7d285178a2a0176c70a0f43b56d99a602f09dd7fb00bec9e139a0796d1107848"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / ee33baa6bbd7 / 3

<a id="canonical-6a6b98882adfedc67c0d956ceb460b6ff35b82abcc9a33815193b48169e31f2a"></a>

<a id="canonical-ef806be1f2d817597111689f726feb7dce4c632b6d9cd295b5114864fddab82a"></a>

## plen property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / ee33baa6bbd7 / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-10a741faf449e4bffbc02b7db374b4b90a50df0c400822fe22efbb20fb2400a6"></a>

<a id="canonical-86d32a09f03173f6fe718caeb873370492c425d0f644e8cee6cd56847aedf50a"></a>

## prefix property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / ee33baa6bbd7 / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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

<a id="canonical-87b483541c20d181de41e8ee8a049ec25de38b47408035d893b36d3e1a07294e"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / ee33baa6bbd7 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-005.md#canonical-1c2fcf1ec9d797aa24c214c65d86001e3db189de3f6cf941bd6f123d6af37ec0)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-8f2051106fca301fd54fafd97dbb7042dfba5314203958b1f40f69bad45607f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-755fa3b877e9b5be87b1de22846c1e9cb2c4f9080073f97250e64b9a74efb467"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / d994258fb099 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [voltstack_cluster.outside_static_routes](data-sources--aws_vpc_site--reference--group-004.md#canonical-fc117b259a0186a106062d2990256669dddd1fe7a759a6558363ac6365334e6d)
- [voltstack_cluster.outside_static_routes.static_route_list](data-sources--aws_vpc_site--reference--group-004.md#canonical-ae52f688a89ba596b353e1dea3f2dd583f69dcc81297a799512f50941f8d0e78)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](data-sources--aws_vpc_site--reference--group-004.md#canonical-595702d7115acc2a4f00c89cf8d2f2bca419ac2ce6c76f96baec14f712436607)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-005.md#canonical-1c2fcf1ec9d797aa24c214c65d86001e3db189de3f6cf941bd6f123d6af37ec0)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-0d5580d4e1c1dab90da025b675ea5d352e9ec2636a9fdc42af9a5049d258ee88"></a>

Type: `"single"`. Computed.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

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

<a id="canonical-87d0415ce74a602bc6d52ec2782ca574ff0a4f5029dfd957fa07fc37db982808"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / d994258fb099 / 3

<a id="canonical-ab020fcbae4852c93dae29a8cf749311c64789a31e097ef86b04f38016bcd3d0"></a>

<a id="canonical-78398cc42c0b66ac90f4793a611643a8262abf472938bf2b4cb712f7e8e2de52"></a>

## plen property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / d994258fb099 / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-b01f550f45ff3592bdf84b544e81d42e180512887ee7c26a822f32020b556d5f"></a>

<a id="canonical-ed5c6a7c5c3f58b449c535a640e45d4a524b3d409c7167d4987ff8c5049efb91"></a>

## prefix property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / d994258fb099 / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

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

<a id="canonical-85c0e0fa5bddf3500ea6a306dcc75781d0d55c3ffc32f66408e744c49ef1c6fa"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / d994258fb099 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_vpc_site--reference--group-005.md#canonical-1c2fcf1ec9d797aa24c214c65d86001e3db189de3f6cf941bd6f123d6af37ec0)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-79c6ac5549c7a8592fc75fcf20480505f861d3867d0f61977082398ba8a05a5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f57a3abf864e4fa8e32ef6385df5f699b563c3e4925043b8ed0bf149ecde4e6e"></a>

## voltstack_cluster.sm_connection_public_ip — voltstack_cluster.sm_connection_public_ip / de110ee1ef29 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- voltstack_cluster.sm_connection_public_ip

<a id="canonical-97857809b6e4aaf698e2af018369b2cb71e290a736c573d0c3b92a741ccdeed3"></a>

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

<a id="canonical-6f2aa4b37bcfbaf59c119480883919a47233a8dff7057244e2d56befa10fb4aa"></a>

## Direct properties — voltstack_cluster.sm_connection_public_ip / de110ee1ef29 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f3996b1607d36c53ea4b8474a10ad475f59edc5b67c8acc73510f47c8bddf116"></a>

## Next pages — voltstack_cluster.sm_connection_public_ip / de110ee1ef29 / 4

- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-85752786ad6ab87137e7b05b9f01c15c8c8862cc2c500c3ccc603a4a1c47df2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-933baa8a1019014e2eb162273a293626c7b988039c2f134fdf5235d112a6adfa"></a>

## voltstack_cluster.sm_connection_pvt_ip — voltstack_cluster.sm_connection_pvt_ip / d780b4eaa362 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- voltstack_cluster.sm_connection_pvt_ip

<a id="canonical-8e5ad07413ed37e6218a8390a4a915685ab48b3c5af4c1d81f61f0bdba77c931"></a>

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

<a id="canonical-e2543b5e05555b5f6ea2dd221e9e460fadd381e0a64845ccfcf5109ebc0995bf"></a>

## Direct properties — voltstack_cluster.sm_connection_pvt_ip / d780b4eaa362 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2cfaf8a2c206cdd4656bb8bbc48a2838e455a4e6194c59f1db47cd6f63befdc4"></a>

## Next pages — voltstack_cluster.sm_connection_pvt_ip / d780b4eaa362 / 4

- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-971b039f053e8eebe6ed31d992852f4386d1fc5f48f855ee86755ed5a22a09cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1891c7cab66bbca479801f47d80aae9030ef86f35126f4141a9d58430f285a0c"></a>

## voltstack_cluster.storage_class_list — voltstack_cluster.storage_class_list / 420dcf63322d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- voltstack_cluster.storage_class_list

<a id="canonical-071b825154ef5ff42e2b2eb01aa5669579d44b9b0354af443f14769848cac3cc"></a>

Type: `"single"`. Computed.

Add additional custom storage classes in Kubernetes for this site.

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

<a id="canonical-a150bfeae7e2bc703f3211b4f96b23d0a224d305f1a7f8ec3e0e930bc862d6b3"></a>

## Direct properties — voltstack_cluster.storage_class_list / 420dcf63322d / 3

- [storage_classes](data-sources--aws_vpc_site--reference--group-005.md#canonical-46aa16c4b660d4d9c8e5eee58d0b5efa68483b94bd7595e66d5bf458323201c6): complete subsection reference.

<a id="canonical-483a603e0f7028043c9839322bae32e8c4794729f3385f08c38690e92b342da1"></a>

## Next pages — voltstack_cluster.storage_class_list / 420dcf63322d / 4

- [voltstack_cluster.storage_class_list.storage_classes](data-sources--aws_vpc_site--reference--group-005.md#canonical-46aa16c4b660d4d9c8e5eee58d0b5efa68483b94bd7595e66d5bf458323201c6)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-46aa16c4b660d4d9c8e5eee58d0b5efa68483b94bd7595e66d5bf458323201c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-289f43d2686cdbeaf758a2c32b6e07e268192a36f40f361bb76ae68448667135"></a>

## voltstack_cluster.storage_class_list.storage_classes — voltstack_cluster.storage_class_list.storage_classes / 59de275d224c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-ce03987dc4fa70508f8e432b0fc8b0c65d9dacdc10695d71f9c4f1c61d03dbb9)
- [voltstack_cluster.storage_class_list](data-sources--aws_vpc_site--reference--group-005.md#canonical-971b039f053e8eebe6ed31d992852f4386d1fc5f48f855ee86755ed5a22a09cb)
- voltstack_cluster.storage_class_list.storage_classes

<a id="canonical-9ba2c7dfc1e6ee5bd9723e300341f6d64c75d5c34c6bc205b73f350376d8f886"></a>

Type: `"list"`. Computed.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1c5dca8877a5933366ac3fb3d135fa813bb555de8be603a3f1daa9fdab2a383f"></a>

## Direct properties — voltstack_cluster.storage_class_list.storage_classes / 59de275d224c / 3

<a id="canonical-a613fc40c316ce6e283381a182e20a2f7e9ad37f17c7d8e96651b0eea4bf291a"></a>

<a id="canonical-41cc37750deaa089148d0a48b26b421151def9d01492e6c3d35ddc69944f6bdc"></a>

## default_storage_class property — voltstack_cluster.storage_class_list.storage_classes / 59de275d224c / 4

Type: `"bool"`. Computed.

Make this storage class default storage class for the K8s cluster.

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

<a id="canonical-ddd0b3b50a6a7e3d5f6fa456d3498f38ebe4498b4b0ada470bc78fb76d5b7592"></a>

<a id="canonical-e1f6c30a7ac1037c4dc0a7f209fe811c9bb2fea00cc9bfbc025c4fcb8de108e0"></a>

## storage_class_name property — voltstack_cluster.storage_class_list.storage_classes / 59de275d224c / 5

Type: `"string"`. Computed.

Name of the storage class as it will appear in K8s.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-9e835990fe92a02a46131ee2ef5872dafdb633f295f07ffb5330a395b94afa95"></a>

## Next pages — voltstack_cluster.storage_class_list.storage_classes / 59de275d224c / 6

- [voltstack_cluster.storage_class_list](data-sources--aws_vpc_site--reference--group-005.md#canonical-971b039f053e8eebe6ed31d992852f4386d1fc5f48f855ee86755ed5a22a09cb)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-c482fd7fdef218c974f1bc3d19e9ec98b067330464c5d118dc6b9acff264f21e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-986d2341a12c5b5cbf9182421bcd9bd0fc0fc7de2e0b37a1dcea4ebda086dda2"></a>

## vpc — vpc / 7ddb32bdc351 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- vpc

<a id="canonical-0168a69df836d59bf8bcdf842afa7944a40e90f9cf9b59b64b7063b32141cd71"></a>

Type: `"single"`. Computed.

Defines choice about AWS VPC for a view.

Upstream description:

This defines choice about AWS VPC for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"new_vpc\",\"vpc_id\"]"
}
```

<a id="canonical-c7dcd490da243cd94d317cd3b67434c9e072be0a0774c7f3b9c7238eb34018ab"></a>

## Direct properties — vpc / 7ddb32bdc351 / 3

- [new_vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-2c5ec47f40d08d199c4e319abffbf698d75affa6fe32230bb004015f4fa0e2f0): complete subsection reference.

<a id="canonical-d26f494a21e204a94afae9a3e9787c1b1c5e341d8ce7dace2e3f5375f3922d69"></a>

<a id="canonical-cd81e1dcc0339aa8783c840f395854a1d769768137f8f3ddd4adb9799fa8aac5"></a>

## vpc_id property — vpc / 7ddb32bdc351 / 4

Type: `"string"`. Computed.

Exclusive with \[new\_vpc\] Information about existing VPC ID.

Upstream description:

Exclusive with \[new\_vpc\] Information about existing VPC ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-e888b4a565ebe214cc9e0fe78bb46832ff31d37c51e0705b691823eee98ff02f"></a>

## Next pages — vpc / 7ddb32bdc351 / 5

- [vpc.new_vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-2c5ec47f40d08d199c4e319abffbf698d75affa6fe32230bb004015f4fa0e2f0)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-2c5ec47f40d08d199c4e319abffbf698d75affa6fe32230bb004015f4fa0e2f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7cc50f2f874032b58f94364376088de30c44b4fba72af9a0a864d79e38d9c95"></a>

## vpc.new_vpc — vpc.new_vpc / e066777e89a9 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-c482fd7fdef218c974f1bc3d19e9ec98b067330464c5d118dc6b9acff264f21e)
- vpc.new_vpc

<a id="canonical-458f41240440241bd926f9145cdba59c9db071450fe3d574027217ded5a730d6"></a>

Type: `"single"`. Computed.

AWS VPC Parameters. Parameters to create new AWS VPC.

Upstream description:

Parameters to create new AWS VPC.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-name_choice": "[\"autogenerate\",\"name_tag\"]"
}
```

<a id="canonical-b41bd113c5db6d7a0751ae2b8b45e6490645f400d6781a5a95838bad4c32fd8e"></a>

## Direct properties — vpc.new_vpc / e066777e89a9 / 3

- [autogenerate](data-sources--aws_vpc_site--reference--group-005.md#canonical-5e04adf742f1fdc8665ad30bdee1b5c8f7f8cd58b3f18100fae705dc0fbfd430): complete subsection reference.

<a id="canonical-e8e2e033a02c013a17cf6ae162f84181e2412faef59152860dc8c7dcb988de8f"></a>

<a id="canonical-6775e136b8d300a1c612dc754728fa6878c8b5c5c502bd3d6fc46631e1a32937"></a>

## name_tag property — vpc.new_vpc / e066777e89a9 / 4

Type: `"string"`. Computed.

Exclusive with \[autogenerate\] Specify the VPC Name.

Upstream description:

Exclusive with \[autogenerate\] Specify the VPC Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-f5456b797aa4a6aa34786e22aaade28ff289bf344f7c97e30ee0ba0d1c6f076c"></a>

<a id="canonical-cb6dc56ad4818f0b7a4b9e5f9c9fb3bb543621131337041af482a203cdd5aecc"></a>

## primary_ipv4 property — vpc.new_vpc / e066777e89a9 / 5

Type: `"string"`. Computed.

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

Upstream description:

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  }
}
```

<a id="canonical-0fb267fe2501f1e708689631b00860da048824a1c3587df2ffc29f63e8863ffc"></a>

## Next pages — vpc.new_vpc / e066777e89a9 / 6

- [vpc.new_vpc.autogenerate](data-sources--aws_vpc_site--reference--group-005.md#canonical-5e04adf742f1fdc8665ad30bdee1b5c8f7f8cd58b3f18100fae705dc0fbfd430)
- [vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-c482fd7fdef218c974f1bc3d19e9ec98b067330464c5d118dc6b9acff264f21e)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-5e04adf742f1fdc8665ad30bdee1b5c8f7f8cd58b3f18100fae705dc0fbfd430"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92e6599c26628d1f617d0e2eebdc914040d3581950974a116357326325c3b676"></a>

## vpc.new_vpc.autogenerate — vpc.new_vpc.autogenerate / a54208ce7d82 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-c482fd7fdef218c974f1bc3d19e9ec98b067330464c5d118dc6b9acff264f21e)
- [vpc.new_vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-2c5ec47f40d08d199c4e319abffbf698d75affa6fe32230bb004015f4fa0e2f0)
- vpc.new_vpc.autogenerate

<a id="canonical-4f17d001a7ddfddd0fabf2c4e978587f46e60dd4508e17ac3074562e3b4a1261"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for autogenerate.

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

<a id="canonical-133cb3c3fe909c27caca741b6067bbd4834a4ccbabc454b14d29cedc8522e1ed"></a>

## Direct properties — vpc.new_vpc.autogenerate / a54208ce7d82 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9af795eb5a3e2c49efe17aa750aa006cea97f5f4c69b17a6d6fb8af014573551"></a>

## Next pages — vpc.new_vpc.autogenerate / a54208ce7d82 / 4

- [vpc.new_vpc](data-sources--aws_vpc_site--reference--group-005.md#canonical-2c5ec47f40d08d199c4e319abffbf698d75affa6fe32230bb004015f4fa0e2f0)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-1aaf17a2664f507686e37894dc171e25232558868c23cf23a2ff535374e932d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca4f855050f006b685bb206e13a311216976e70b7809aa0acbdbdf8c031297fd"></a>

## waf_signatures — waf_signatures / 302c65c73dbf / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- waf_signatures

<a id="canonical-8ebffefa7527917c1526c33cdd0d25fdd82a056dfb700c3a99363af0eccbad85"></a>

Type: `"single"`. Computed.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

<a id="canonical-d72c3e3a70164bb4cd6af394e9cb7628467823defdb849f7e502fc583bf0665c"></a>

## Direct properties — waf_signatures / 302c65c73dbf / 3

- [automatic](data-sources--aws_vpc_site--reference--group-005.md#canonical-019fee67523ef23ff78b23369f8d37132d7326034bcfe0f839f429f60578df6d): complete subsection reference.

- [manual](data-sources--aws_vpc_site--reference--group-005.md#canonical-60c403d5309a053e007c8c0bb9136251255785018787413040f0bfde1ea600c5): complete subsection reference.

<a id="canonical-8e804dd4b13e4136b595bb6d5e62ebfa51f76eabc6e7225a033519d96f6901f4"></a>

## Next pages — waf_signatures / 302c65c73dbf / 4

- [waf_signatures.automatic](data-sources--aws_vpc_site--reference--group-005.md#canonical-019fee67523ef23ff78b23369f8d37132d7326034bcfe0f839f429f60578df6d)
- [waf_signatures.manual](data-sources--aws_vpc_site--reference--group-005.md#canonical-60c403d5309a053e007c8c0bb9136251255785018787413040f0bfde1ea600c5)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-019fee67523ef23ff78b23369f8d37132d7326034bcfe0f839f429f60578df6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84b31375d84b0f81f02f870c3b36c7e0545385080654c9f5315d3136989a1aca"></a>

## waf_signatures.automatic — waf_signatures.automatic / d2291412019b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [waf_signatures](data-sources--aws_vpc_site--reference--group-005.md#canonical-1aaf17a2664f507686e37894dc171e25232558868c23cf23a2ff535374e932d3)
- waf_signatures.automatic

<a id="canonical-bc736858814028dbf28a63691e3fc0e4aede256d82f72ca851df9ab0b715f8c6"></a>

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

<a id="canonical-24ad84ded38b164909f6cf0f4d37276554f06b4b9e3f7c552e7fde4edf71ccfa"></a>

## Direct properties — waf_signatures.automatic / d2291412019b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-127a05e9944da59fbe0e8d8d68e7f526d847cfb4208e9d96df7d8ab687e0dc61"></a>

## Next pages — waf_signatures.automatic / d2291412019b / 4

- [waf_signatures](data-sources--aws_vpc_site--reference--group-005.md#canonical-1aaf17a2664f507686e37894dc171e25232558868c23cf23a2ff535374e932d3)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)

<a id="canonical-60c403d5309a053e007c8c0bb9136251255785018787413040f0bfde1ea600c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5644e991f096a801ca943e28875b1a82af97978ddf4d35737584b679591e408"></a>

## waf_signatures.manual — waf_signatures.manual / a5be85ba9f4b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0ae65da546e6296a7be7a38b800cc7b7578f4d214a8a51600c66364338fbe66c)
- [waf_signatures](data-sources--aws_vpc_site--reference--group-005.md#canonical-1aaf17a2664f507686e37894dc171e25232558868c23cf23a2ff535374e932d3)
- waf_signatures.manual

<a id="canonical-923c39f5de1c0e50537b66f5688143106a161e710bad00e4d82b7a3087d0b0d6"></a>

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

<a id="canonical-08fc48be4f5409afa1f3971af91b065bcaa5d7ad335762b3420c92e6b199ac74"></a>

## Direct properties — waf_signatures.manual / a5be85ba9f4b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-55c9e0e774f1f9b619dddfaa2de05ee84a1eb2d406b7b712f7da7cdf41ab337a"></a>

## Next pages — waf_signatures.manual / a5be85ba9f4b / 4

- [waf_signatures](data-sources--aws_vpc_site--reference--group-005.md#canonical-1aaf17a2664f507686e37894dc171e25232558868c23cf23a2ff535374e932d3)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-e044179917c669bafdb839ba10c1a81d9f41961466bb9ae2fc71f3716df4b1ab)
