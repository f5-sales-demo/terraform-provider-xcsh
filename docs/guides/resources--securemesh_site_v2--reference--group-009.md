---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-8471b5284905701959fb29c6c9361032eb337e809410e45dd6356a27cffcb956"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / f15c40bf3149 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-008.md#canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-9e646e166817f2e4318cfa5d9ac56a3d97d002ee51eb0d6574c83e9308b3c94a"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0c576e7708140dbc36b64bd80e14b6dc87ef5c286614549b363146b61fec2cc"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / f15c40bf3149 / 3

<a id="canonical-c1fe2531d87e7d956e39cd046407b936fb28502dfc0c1a02274b326c52cdaf1f"></a>

<a id="canonical-50072d734e0e87b6156740da98c4acb0a1aa14824ecdea1be3506f111f340755"></a>

## network_prefix property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / f15c40bf3149 / 4

Type: `"string"`. Optional.

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

<a id="canonical-e31587ecf66ffdc832bea1131a31df04eef3f3b1ce2b2aece8dad2501e5ac019"></a>

<a id="canonical-5477cc79680018e8019fcb12c0010ea556ac4e090c02b79ef44d5e0c25fbd386"></a>

## pool_settings property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / f15c40bf3149 / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

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

- [pools](resources--securemesh_site_v2--reference--group-009.md#canonical-1e05340870a58f2d20c2770b6d84b6ec6178b852d81992f8064fb1979b7c7d42): complete subsection reference.

<a id="canonical-2b9d3eb394c349ec127700c817819da5f49dfb3a324591f375c38b1f0b84c682"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / f15c40bf3149 / 6

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-009.md#canonical-1e05340870a58f2d20c2770b6d84b6ec6178b852d81992f8064fb1979b7c7d42)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-008.md#canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1e05340870a58f2d20c2770b6d84b6ec6178b852d81992f8064fb1979b7c7d42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48b1ffe2599dfdf25011d9d111cb89367cd397b2890c7fddbfa322dd32625dd8"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / e324dbfbbb61 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-008.md#canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-5948b425d72cd4c5cdc03867645116417c4971439c13f1ef39d4f94149e3940e)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-10d0c3f12ed8a48cc3b40af2af6996ce4a6ddb9a90a41db3da7abeff51dbc888"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-1a99090e294611ced64c0f2f6197343b02a280491bb53e83a26ada0f31fb5c16"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / e324dbfbbb61 / 3

<a id="canonical-fd25be7f97802271d1b9b5e283dfabe8a0f27c1a0f56951fa548dc94f4b2fe98"></a>

<a id="canonical-b825e941ebda88b280759202335ac6568347420f02281c3891008f8af04ccec8"></a>

## end_ip property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / e324dbfbbb61 / 4

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

<a id="canonical-adfd4c665634fb8f92dcbe00012c414914abc9bcfbe32471fc606eb60bec31e1"></a>

<a id="canonical-4147937c9f159a9a235a1e5d6b5b47782f1aad1e7ec7bde6c20c7b35439430b3"></a>

## start_ip property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / e324dbfbbb61 / 5

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

<a id="canonical-bbb622936d028954d6ad7105a294d0c55f1c6b90cb2f0799a33d6fa7d7a579e3"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dh / e324dbfbbb61 / 6

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-008.md#canonical-5948b425d72cd4c5cdc03867645116417c4971439c13f1ef39d4f94149e3940e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b00bb0bf541e65537eba77487f37944348350895c53af02d6b163aecd82724e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c54c27ce59d911bd382c96c4b015cd9644659a5405b449ff4b26f1e099e3dd7f"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / 211451efbb6a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-008.md#canonical-1c12c6a4189d310795a700b415763e502e4093ebe556232fde1097fd1c3e89c6)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-008.md#canonical-45cd86413ffc6c1020e5e3260e63638fdeca6ff9b8652fc60a8751cef612cde5)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-008.md#canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-af5060875a4bf2c729e5206c963d1d987a15f8c8b1d3fd3ef7b6324d73ae41e4"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-52da50bb4931dfbd877239bf8787e37a8b32034b3c34a8671fd578f207db6781"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / 211451efbb6a / 3

<a id="canonical-f7df95a790d8ae4208dd0d3a0aba00857f7264e759afe735c99b3afeacf2865b"></a>

<a id="canonical-4ee94db3acce3c61226a8ae049071320837534ecdc83ba18d67f39add7b5c9c6"></a>

## interface_ip_map property — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / 211451efbb6a / 4

Type: `["map", "string"]`. Optional.

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

<a id="canonical-578c4a94d294f9f20f1bc844e2afe719a9fd2656eccef76406738accec032f92"></a>

## Next pages — equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.in / 211451efbb6a / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-008.md#canonical-cbea60ed74d72383917cc47841d496e8d6c9430cd7bb42490e28bcc441ac0081)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-fbf5804b938e70b250d4662bdd3a52ebfa361ea69199b210816c52e72aba3464"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8e559b7e2154bdef98b676e3a6d96c21c6f97cf08abee4eb916cc214cb64f60"></a>

## equinix.not_managed.node_list.interface_list.monitor — equinix.not_managed.node_list.interface_list.monitor / eb5fa5d41d9b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.monitor

<a id="canonical-dba24f4ac4457840d916fed5cbccd0cd7d04ba99553b023ca37951c4924a11f2"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
monitor = {}
```

<a id="canonical-29439187d786bf2a19cab65c0189a44282b07684ffca2b4f865c0846045fba57"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.monitor / eb5fa5d41d9b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-063ffb9f07df22030373ac2e92bd200c4d7ee10a1d20c31ff8baf7810d78df5c"></a>

## Next pages — equinix.not_managed.node_list.interface_list.monitor / eb5fa5d41d9b / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-58f9af6210d386d904b25e71cf30a85287899ea13c6464aee92f1a9a8d7d7f12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6072c7d0bf2908e8129a3624533c1e29b52f4741dee39b24a0700e0b5571aa6c"></a>

## equinix.not_managed.node_list.interface_list.monitor_disabled — equinix.not_managed.node_list.interface_list.monitor_disabled / d7e19b062c25 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-7b01e002e4a75491984bb3b0712faa1b43285c736d990c0e7bf7439052077de5"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
monitor_disabled = {}
```

<a id="canonical-3131510b7ac75c2ce3ebda7c81b9d831a1f871f77844df6f4ee1868ee260209d"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.monitor_disabled / d7e19b062c25 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9aa2565effa72eda2f5dc8901f6a743f813ff9c872365bfcaa41d93b200573e0"></a>

## Next pages — equinix.not_managed.node_list.interface_list.monitor_disabled / d7e19b062c25 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6aec47591d0b4b6ab9b29ac34dd0a58e15c9ace78e3f58baaeb0aa3dee04db18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13fb0435d07027cf01a438b032a4df3d25e919a1a6dee129fc46b9cacf1d7def"></a>

## equinix.not_managed.node_list.interface_list.network_option — equinix.not_managed.node_list.interface_list.network_option / 394d4aa3146b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.network_option

<a id="canonical-44a4289a87cb29397288e74d0079e9c9456879f85b35f9211e3596c8ea05f8e4"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional.

Upstream description:

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
network_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-a287ec2a4c01b233d49a8821e2cef1768e5180b46ae1c8819aa2435f98f54405"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.network_option / 394d4aa3146b / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-009.md#canonical-17fb4e1463d91e08160e6ccbaad03f36392ba3777694d20bc319507d449ca36f): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-009.md#canonical-409a23307c1600e0af488bc2870aabcd94df96b3b5a82af1c5a0bd264ac2d711): complete subsection reference.

<a id="canonical-5022a0dc98218e81024ea3bb758fa8bb36b09d0c035b6cfbc60e6cc0c36baaea"></a>

## Next pages — equinix.not_managed.node_list.interface_list.network_option / 394d4aa3146b / 4

- [equinix.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-009.md#canonical-17fb4e1463d91e08160e6ccbaad03f36392ba3777694d20bc319507d449ca36f)
- [equinix.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-009.md#canonical-409a23307c1600e0af488bc2870aabcd94df96b3b5a82af1c5a0bd264ac2d711)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-17fb4e1463d91e08160e6ccbaad03f36392ba3777694d20bc319507d449ca36f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7a4c76536d08234d4fe8a04be3e1316387f28539487c5bc424bb55c02ff749f"></a>

## equinix.not_managed.node_list.interface_list.network_option.site_local_inside_network — equinix.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 58089091ebdd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-6aec47591d0b4b6ab9b29ac34dd0a58e15c9ace78e3f58baaeb0aa3dee04db18)
- equinix.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-f8505553da5f7fd7969d163ca463bce8cd9181848707512a4e41768a0ee3cdcd"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
site_local_inside_network = {}
```

<a id="canonical-06d54c9e32d97e6ff903c0a4970069540c98a331a4279d4c45b5986bb0c5a1fd"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 58089091ebdd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12b0a95280c8962c2796111f983b916ebdcd59644a4835b8cb70a6c922bf6740"></a>

## Next pages — equinix.not_managed.node_list.interface_list.network_option.site_local_inside_ne / 58089091ebdd / 4

- [equinix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-6aec47591d0b4b6ab9b29ac34dd0a58e15c9ace78e3f58baaeb0aa3dee04db18)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-409a23307c1600e0af488bc2870aabcd94df96b3b5a82af1c5a0bd264ac2d711"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-666fa0ccd07dd062e2387e19612af3b4e0a40f133a2270df3d619f49f15b1ea7"></a>

## equinix.not_managed.node_list.interface_list.network_option.site_local_network — equinix.not_managed.node_list.interface_list.network_option.site_local_network / 0126ed805367 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-6aec47591d0b4b6ab9b29ac34dd0a58e15c9ace78e3f58baaeb0aa3dee04db18)
- equinix.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-29f641ab333a6d0eb4405d48c3f7301855f984c154fc00785af26f723d51ecd7"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
site_local_network = {}
```

<a id="canonical-549a5bc4878071479954ae400e4a1d44e0b920fff15f887387052e1c4c0e1f51"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.network_option.site_local_network / 0126ed805367 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-49ee45b7c033bd844583f3fe81adb8c6ebc21dccca13dd203b2aa4e52caafc97"></a>

## Next pages — equinix.not_managed.node_list.interface_list.network_option.site_local_network / 0126ed805367 / 4

- [equinix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-6aec47591d0b4b6ab9b29ac34dd0a58e15c9ace78e3f58baaeb0aa3dee04db18)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-576e9a39c18b846647c47d4f4518fbe43ab40bfac003f13930a0aebfe06a661f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-451368c3dcc14aab37d93841dae006d66131053d01f6095460f3590fc537b674"></a>

## equinix.not_managed.node_list.interface_list.no_ipv4_address — equinix.not_managed.node_list.interface_list.no_ipv4_address / 20c7dbe56b45 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-94050893703d463e8c194ee7e63d0166d59114744312de022fac097144547183"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_ipv4_address = {}
```

<a id="canonical-2a75294cd8962e9e2bc9fe64045c3b5902d36b9db2e46699f2b3318f7e6ee0b9"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.no_ipv4_address / 20c7dbe56b45 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-87b79b6dd75049e3168ebb8c2bd890b49c1de24d7375040f81ec07b49cb123e3"></a>

## Next pages — equinix.not_managed.node_list.interface_list.no_ipv4_address / 20c7dbe56b45 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-65a381955fa5cd94cc5fa2fbc222152f00bc574e39921048c235ddd028d24854"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05960d6dd4570a273cac9fae8c9307dd066f4972268014c1cc5e4a418273bf2a"></a>

## equinix.not_managed.node_list.interface_list.no_ipv6_address — equinix.not_managed.node_list.interface_list.no_ipv6_address / 538ad8398c35 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-21b6d57d3e1923fc341570b2f1b4a2fd90ad259945af5e892b4ca54f915f918b"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_ipv6_address = {}
```

<a id="canonical-ef6ad896621e4c91e100a254e4da71b7b095953022e6adf345b9e47d34a5523c"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.no_ipv6_address / 538ad8398c35 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb0046bb509d25ca6aaed7d8fe3deee9d516f3628452f1611f68e352339645c8"></a>

## Next pages — equinix.not_managed.node_list.interface_list.no_ipv6_address / 538ad8398c35 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-40cba71be545e59450d86e3f8ff40174e36f79e098e9ded35f2e282a5c77d951"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8407866a5e7e987b805447c57011fb36533aa4a1e1716c2140370ced2ec4af1c"></a>

## equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 3a5bb0833b63 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-fc0ba2c81b2503d74dd56f478da278569fdd9938e450c86a73fc9fd92ab7422f"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
site_to_site_connectivity_interface_disabled = {}
```

<a id="canonical-539f3c26bb4741c5550545ad0d6a914745513b51c427bd67e43f809d86efa2c7"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 3a5bb0833b63 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f0e2d5b9951645bd7a924d1a73e779784bfb8a4bbdb5f815e9b7335e2d6479cc"></a>

## Next pages — equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface / 3a5bb0833b63 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b7349d6306fa8949ba028853754c52c8cfd146d76baed5a41c8267bb705a072a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2a7effb808312393879e277b2863e7fa809f6b75d6526504023adf987eec1a3"></a>

## equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface / cb1a0a5d3539 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-538e6cbc75612f5bdb03464eca1fff96ebbb8ae0207e8611e4e5c6b89560e276"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
site_to_site_connectivity_interface_enabled = {}
```

<a id="canonical-b55307a5161cc0c5223646e77932ab2c765aa897b11b9b891ef3c1af997c4432"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface / cb1a0a5d3539 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a0b8585dfbe95440d44d3c85fa1cae44bdf38491a14c7e08e90ae1ff879b39e"></a>

## Next pages — equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface / cb1a0a5d3539 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-89207d7f8b2c06f093ea659df00de8be4131ba92fa4513d8bf017732eebb960b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a05c1952e1e5d5af49647f409fd637deacc26083dced9bf4bf7ce6c415ce829"></a>

## equinix.not_managed.node_list.interface_list.static_ip — equinix.not_managed.node_list.interface_list.static_ip / 88ddd4268ff4 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.static_ip

<a id="canonical-dc8d5efb704d1030e84f2017af71270e90f50f922a7ec58b6745abe0fb3824f7"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
```

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

Terraform syntax:

```terraform
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1c31184d4bbb734c9b6391147bd1c8fd0198863205bdae58b5187f27e4b7825b"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.static_ip / 88ddd4268ff4 / 3

<a id="canonical-fe76b5eb42115b358faa41bc1ded06bd95b8e03f053c8bd32a4146035d9053f1"></a>

<a id="canonical-a81e51a02815998217ba2aba31d7a69812039d3068aaf4c1d2c3b3c25cea0f9a"></a>

## default_gw property — equinix.not_managed.node_list.interface_list.static_ip / 88ddd4268ff4 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

<a id="canonical-bc55d0c116216d044cf153a6f656954c00d8e99cd5754e4e19e8101c264313ae"></a>

<a id="canonical-a28003f2afb3025c0398f6935aec19fee6b1832ca74123034f56b724e531597b"></a>

## dns_server property — equinix.not_managed.node_list.interface_list.static_ip / 88ddd4268ff4 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-33121affe0277ef37cade40c86109e402091dd8f947496edd203f7cf062ac25a"></a>

<a id="canonical-73ba8402c89a8216f16408952b1c8bd173904bce9304fedb61a4d33dd0e44505"></a>

## ip_address property — equinix.not_managed.node_list.interface_list.static_ip / 88ddd4268ff4 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

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

<a id="canonical-7c3cb1c5017e7d37bfab9c9eac41723feb42541530a42abe6dd8443ace1583f2"></a>

## Next pages — equinix.not_managed.node_list.interface_list.static_ip / 88ddd4268ff4 / 7

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-90a8505774e15c93615dbbcd1864cd8bb1dcd473856f3b8c346fad00ae52edbc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c549f349f6cc37029bd0ad0259aad032aff739a7970534ba96d60e1f35ed5d3"></a>

## equinix.not_managed.node_list.interface_list.static_ipv6_address — equinix.not_managed.node_list.interface_list.static_ipv6_address / 05612da4f2ba / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-449afa59f9a820895b5c97e3cae3c8aa8534d4d9650ebc560e9721300b9ca34d"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
```

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

Terraform syntax:

```terraform
static_ipv6_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-e711d43c9b4163a786942d70ddfafb85c7d7c07d4da86e623f3091d3c9956da0"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.static_ipv6_address / 05612da4f2ba / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-1742b560e3af27e8cf4d8bf0a2c67b42d28c237491c7990b0d47e0f0df7ce400): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-b79edf5d08fc4949cf70beb30b27323ffb9140a736624331f091d264188aeca2): complete subsection reference.

<a id="canonical-7a5fd59b4d3a78bd4191cb4a678ac35ab9eefcf05d9fe70fd7c389843bfec8e5"></a>

## Next pages — equinix.not_managed.node_list.interface_list.static_ipv6_address / 05612da4f2ba / 4

- [equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-1742b560e3af27e8cf4d8bf0a2c67b42d28c237491c7990b0d47e0f0df7ce400)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-b79edf5d08fc4949cf70beb30b27323ffb9140a736624331f091d264188aeca2)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1742b560e3af27e8cf4d8bf0a2c67b42d28c237491c7990b0d47e0f0df7ce400"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-795bf913f4ed962020b532dba5d6ba2edad2ef6a631ef9061b00d5e336ce07c5"></a>

## equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / ce9a1d080c76 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-90a8505774e15c93615dbbcd1864cd8bb1dcd473856f3b8c346fad00ae52edbc)
- equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-a7e1427f7267beac60b7c8d3fd9762d73dd573d94366a6702c2ef8f9ad9b4123"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-d1baaf4571e11692e4be24f9f45497fce5fba5e7ca511b4c76614d564b58b0b5"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / ce9a1d080c76 / 3

<a id="canonical-830dea9e1c48d483aac53178a466ef9bd6efd28560aecf627468666d2cc20c3a"></a>

<a id="canonical-cc994ac5a997435dc700b5057794a2f95a097476e20cef344860f7a393bb8a03"></a>

## interface_ip_map property — equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / ce9a1d080c76 / 4

Type: `["map", "string"]`. Optional.

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

<a id="canonical-5a718a4d9355f3bd043ca2aee0b2d8d389c6b5f9102f735404a111d010180f4a"></a>

## Next pages — equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ / ce9a1d080c76 / 5

- [equinix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-90a8505774e15c93615dbbcd1864cd8bb1dcd473856f3b8c346fad00ae52edbc)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b79edf5d08fc4949cf70beb30b27323ffb9140a736624331f091d264188aeca2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a49c478ca4d3defb21eadcaab660613590373b78db40806602ee6f26bc117175"></a>

## equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / c1ea42fb63da / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-90a8505774e15c93615dbbcd1864cd8bb1dcd473856f3b8c346fad00ae52edbc)
- equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-46b23fe9a7a0fda5ac61a4a9b0bbf486107c4f8a8572b3b90cbe7ce896d92e12"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
```

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

Terraform syntax:

```terraform
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-9878d4e0fcd6146105862c9b407ee797304383d4668c0e234de99c22148e8cf7"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / c1ea42fb63da / 3

<a id="canonical-b853d89e12565cbcf5b3683804ad0d3209c31bedd33e21923a12a9bf435fbbb8"></a>

<a id="canonical-c6fd65b0e0111f9481823c58d84ac30569503bd84f13841dd05e4d52e7da4fe5"></a>

## default_gw property — equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / c1ea42fb63da / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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

<a id="canonical-09f4b5b0746075db6a5f45e368d171c9152ac0939cc509736faf3bea7fdc6fad"></a>

<a id="canonical-7e9ca315e5ba55be3ec2284f8a46d4ecfeaf47af77fc9382ff71e29642c5773b"></a>

## dns_server property — equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / c1ea42fb63da / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-49454c181281467dee3024eb3ee1f634c98eaee4dce257851b2db5e63d684ab2"></a>

<a id="canonical-e2b1e6a4684873278515bf46a2b81afddcff38197093ccb7a5781672732eef30"></a>

## ip_address property — equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / c1ea42fb63da / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

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

<a id="canonical-61a3a52a9be5f7ea3a65da3a3d908b90be32369d8935c711dc0aa555480e43ce"></a>

## Next pages — equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / c1ea42fb63da / 7

- [equinix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-90a8505774e15c93615dbbcd1864cd8bb1dcd473856f3b8c346fad00ae52edbc)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-7884ed7284e85608b088cc98b43a4f6bd2fdb09a059944df8d2c819dc7536ed9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4923405ce10d2deacb6d0059512b35583c08410c5a57d43c507e9e9342be9e1"></a>

## equinix.not_managed.node_list.interface_list.vlan_interface — equinix.not_managed.node_list.interface_list.vlan_interface / fe7d8fb25ddc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-bef28fc316b949e4311589fd003ba9f37420abb7da43fe9107a077b4cfaf68ca)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-8985634e7e0b17670ff971bed673dd8e58c7052d2fb9f9778019b8a9077fad90)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-8293c0c898b3e7c87dea289512cbe8a12de0ce11007babc4bdd7d15f5642f83f)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- equinix.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-e4f4c85757eaea6e951d5b05364113d4a80ef3a9017109b5b2feaf1b0af300e6"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device",
    "vlan_id")}
```

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

Terraform syntax:

```terraform
vlan_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2ae584fe59832cb58f377b7dd7468fe33ea64a7e0e429b2f4f81a07a828f160"></a>

## Direct properties — equinix.not_managed.node_list.interface_list.vlan_interface / fe7d8fb25ddc / 3

<a id="canonical-95a31dd322fb3883eaa7353b9fa381b3530bd1c632948c168513fd3a7aaa7add"></a>

<a id="canonical-fed1647f918d03a5e636689582c2f9f9e9e18b888cc142d78f89eb4a34670261"></a>

## device property — equinix.not_managed.node_list.interface_list.vlan_interface / fe7d8fb25ddc / 4

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-d81e83372229fef08fb0523cf8a84ae6b09dfd66e94d77989c81e7db2ee9fbc1"></a>

<a id="canonical-8c46749cee7278bddd5ce9fa99ebc8f11062be726837de9806f132d5d9d3b178"></a>

## vlan_id property — equinix.not_managed.node_list.interface_list.vlan_interface / fe7d8fb25ddc / 5

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-dae144b48a077820ff38bc0edaa663d463f144f9616efa59c66324dc361f1446"></a>

## Next pages — equinix.not_managed.node_list.interface_list.vlan_interface / fe7d8fb25ddc / 6

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-288e9d7a10024589a37e2347d27eeb2571e78729c2794c4f5b46bb3831c23a1c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c9c61f342c718474ad94bfa872822347d516d70f708bcd7ce7358dfea3955e25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4180ddde77336a7ae7222ae9b04c250034820482f7a634db2793cf38b89dd896"></a>

## f5_proxy — f5_proxy / 73a3ab35b5e8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- f5_proxy

<a id="canonical-3b88f892a373b4cf98fd02934dffd2d101bec6cb32dfdf6f16a7f94e63af71f2"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
f5_proxy = {}
```

<a id="canonical-d4497090f0174a69ea0bce24e62823d4e132fadb9f5ba042c884a98f87fd29a7"></a>

## Direct properties — f5_proxy / 73a3ab35b5e8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5fed47004424dc5d3a3360d40208f3e5739ccec3d808ae31d9f2f31bc5954b93"></a>

## Next pages — f5_proxy / 73a3ab35b5e8 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4f628c78623d8748137dba5f904f1f4b153cecf3241bd86cc634765a6269f77"></a>

## gcp — gcp / a0ab0bc05668 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- gcp

<a id="canonical-ef56793a72c1e3ea69ab4c977ea4820d56ce89de0243369831359f5239354dfd"></a>

Type: `"object"`. single nested block, Optional.

GCP Provider Type. GCP Provider Type.

Upstream description:

GCP Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

Terraform syntax:

```terraform
gcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-f22a247154e8b908bb2809aea74c0b027af2b4f7a5a04af5ddbc19360fd3bedc"></a>

## Direct properties — gcp / a0ab0bc05668 / 3

- [not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011): complete subsection reference.

<a id="canonical-43da289c8d24c65f6ea32409b32de84d03330a0a3e028a6b0434ada61bcf0eaa"></a>

## Next pages — gcp / a0ab0bc05668 / 4

- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-266b1ca64a265e167d2a264237cae2eddb2447cdc82efd45368f07d726ea3a6d"></a>

## gcp.not_managed — gcp.not_managed / 0ed34c2c5a80 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- gcp.not_managed

<a id="canonical-5e25814a58e7fa254bf40ced3bb4b47c989624bf88366f4754dceab9bdc975f3"></a>

Type: `"object"`. single nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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

Terraform syntax:

```terraform
not_managed {
  # Configure direct properties listed below.
}
```

<a id="canonical-2d015fd06eff7ddc8b7aad75708b4e3b68083914078d37cee876cd5014df79a3"></a>

## Direct properties — gcp.not_managed / 0ed34c2c5a80 / 3

- [node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc): complete subsection reference.

<a id="canonical-24de757eddc7b93688cbd517cfd4d2d7ca65c8353f8aeba9287743c92c12bcfd"></a>

## Next pages — gcp.not_managed / 0ed34c2c5a80 / 4

- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bea661b930e5dee29e67c836b48bf74869573267b45f5233d47114d315930ce5"></a>

## gcp.not_managed.node_list — gcp.not_managed.node_list / 6960cbc02e3d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- gcp.not_managed.node_list

<a id="canonical-86a0b7b63bf80676e2bc601a858a6352cc7647d57e3ed9e70e64644c5aca6b86"></a>

Type: `"object"`. list nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-4fdfe4dc7fa8d9202bbcf51748aa021d9a8a98b4e504ec5e3257a49589150fc4"></a>

## Direct properties — gcp.not_managed.node_list / 6960cbc02e3d / 3

<a id="canonical-d2554576e097f47c69819f1506df0a4c2a96559ef4ffe56fa91e627c65dcd822"></a>

<a id="canonical-6124e0ea212dac5e7dbbcc28eebf89bb0bf3fad0d4b7ffa9fcce156cc1873f3c"></a>

## hostname property — gcp.not_managed.node_list / 6960cbc02e3d / 4

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d): complete subsection reference.

<a id="canonical-326fc00435f603c0fa351e6068d53fe71417a34ecebf799a2b2584bff301ff23"></a>

<a id="canonical-4d5807356c8a5841cbf547685238c383967dc08c09bdc8e6a6b199237900cafe"></a>

## public_ip property — gcp.not_managed.node_list / 6960cbc02e3d / 5

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-4c24c3ce92ae2e5d17fd2fe72147c4d24e63b114abfe2dc0173a4fbcb712397a"></a>

<a id="canonical-afc1b536bd0a9fd408641a3ec2f958a7eeb5fbc4a02d1bf3ad47b3d5499c5e94"></a>

## type property — gcp.not_managed.node_list / 6960cbc02e3d / 6

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-a53e7539ca0b25cb913e111c46326ea5422e40ff28f8134363cc89d47e4a6196"></a>

## Next pages — gcp.not_managed.node_list / 6960cbc02e3d / 7

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b460f1bc67a207b92bfeed1092a5aff47cfa93960feaf798b3604c6daa420613"></a>

## gcp.not_managed.node_list.interface_list — gcp.not_managed.node_list.interface_list / 7d269b90bbce / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- gcp.not_managed.node_list.interface_list

<a id="canonical-7bdd78a3625ff55b2d7eac9ff47fbc8a2d95a801b4d14cc08df709e9d0d0e0cc"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
```

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-32066a7347d77d781c7530898d39cf4241c7719cba5c1c329d3555b4dd5b3009"></a>

## Direct properties — gcp.not_managed.node_list.interface_list / 7d269b90bbce / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-76751a759bd888723b322764cc48f95657f7c1454f183af2fcac56d1b0642102): complete subsection reference.

<a id="canonical-f6f2966ae0c9e5226d071583775935d701b76fb61f1808db1d532557516f277d"></a>

<a id="canonical-0999d86180b635f97f0c7fb56498a5919abecb0fc74b1e6d4070182ccc8edd53"></a>

## description_spec property — gcp.not_managed.node_list.interface_list / 7d269b90bbce / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-009.md#canonical-0f68f199f2f5898bfb2457b75b8388410f6935b7cc7a00c794f0efcf6e1d90e4): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-a5faa0ac1acdae57e16b8dfff475700551e589d1639ef0d98fd2d17746b0f5d5): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871): complete subsection reference.

<a id="canonical-b836ede284280ee57da0e41c761a39cfabbda0372e440bbbd8607a86f3d0d74f"></a>

<a id="canonical-e951c3e32cc08ad56a56d16bee554dc74e8390df6c7331097a1309acbc6e2903"></a>

## is_management property — gcp.not_managed.node_list.interface_list / 7d269b90bbce / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-fb44e79abd05fee5ab3b04aba14a5c99a612498b3842de78b5a4067d23ee1b20"></a>

<a id="canonical-dc35bd57ce67007573a456454b1d5253302476576b2dac3267dff9f537e891af"></a>

## is_primary property — gcp.not_managed.node_list.interface_list / 7d269b90bbce / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-1f068b4bd427f25cabf390d000c09a9a23bf9d7993b826e3b922e7a9637a960d"></a>

<a id="canonical-3fe030535ed66d872771325c9fd5b412f87396b83478a71ac7ba91e44da61418"></a>

## labels property — gcp.not_managed.node_list.interface_list / 7d269b90bbce / 7

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](resources--securemesh_site_v2--reference--group-010.md#canonical-65af5831979ef12c631b1c984ac3d4fb2563e5604c340073aaa66f8567da5694): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-010.md#canonical-941d0f2123dd29a00e7e3d87119b1448b769fbf2a857a44c2af363899501710b): complete subsection reference.

<a id="canonical-f5aab8ea8a24fee1a982747c1600a15492a02eb0b470960304c493fb3524a8bb"></a>

<a id="canonical-359388d108ac21665af42c110179cf9955e9c6d45127100f9fac8ee03626a481"></a>

## mtu property — gcp.not_managed.node_list.interface_list / 7d269b90bbce / 8

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-d7fde9c9062bad6b2fef3c1153942e49ee8a8831b036b21ff6e26644015d6daf"></a>

<a id="canonical-debed863e81fa5ecb69f1a81ff22e07311393403ff50dff5f2a8c54f6fe5cd6e"></a>

## name property — gcp.not_managed.node_list.interface_list / 7d269b90bbce / 9

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-1cacfff659665dd3d788ff4eafe976e03e12c63d44015cf124750f19beeabbeb): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-010.md#canonical-63ee8686a1ea8c2c2c0b440cfc37be7530dd315db0775b54c78f7c48b846c7cf): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-c04cbb76b113bf3f3f0f20027a03b6ee82a3cc0d7d581852b9ff3ce3812f3296): complete subsection reference.

<a id="canonical-19b1b43faf3a6216f952c40b1b1e2eb43c023c74da214825ea84b0e82c08556d"></a>

<a id="canonical-085971e843dc676717be7ebdb19f559fefdc74ee317338534d6e6b2ba73b0ed0"></a>

## priority property — gcp.not_managed.node_list.interface_list / 7d269b90bbce / 10

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-010.md#canonical-5f86f36cff19c3cd29fb66ddc395fab7476bf1bd10b6875a66aa54c21d1393ec): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-010.md#canonical-28408f9ee9e1d8ff6f1ce37bea5e41ab59397661411e134d64c8640e3b836856): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-4c6101f3caa75e81e6a892525da8c929a8ff80591bfdd566d93ac8b7dfdd963f): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-a5035a3970d26e757895a07edf91e186034749f878ed0bc9380f4a89a0eb51b4): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-ba3cbf3ac22df7aa55ff82cb6bec071600a413bdf4edee115662459b56451815): complete subsection reference.

<a id="canonical-8eb080ad452079854ab9cd15efe4a46064b275c8f875183eb15c11d0f3baa9da"></a>

## Next pages — gcp.not_managed.node_list.interface_list / 7d269b90bbce / 11

- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-76751a759bd888723b322764cc48f95657f7c1454f183af2fcac56d1b0642102)
- [gcp.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-009.md#canonical-0f68f199f2f5898bfb2457b75b8388410f6935b7cc7a00c794f0efcf6e1d90e4)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- [gcp.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-a5faa0ac1acdae57e16b8dfff475700551e589d1639ef0d98fd2d17746b0f5d5)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-010.md#canonical-65af5831979ef12c631b1c984ac3d4fb2563e5604c340073aaa66f8567da5694)
- [gcp.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-010.md#canonical-941d0f2123dd29a00e7e3d87119b1448b769fbf2a857a44c2af363899501710b)
- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-1cacfff659665dd3d788ff4eafe976e03e12c63d44015cf124750f19beeabbeb)
- [gcp.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-010.md#canonical-63ee8686a1ea8c2c2c0b440cfc37be7530dd315db0775b54c78f7c48b846c7cf)
- [gcp.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-c04cbb76b113bf3f3f0f20027a03b6ee82a3cc0d7d581852b9ff3ce3812f3296)
- [gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-010.md#canonical-5f86f36cff19c3cd29fb66ddc395fab7476bf1bd10b6875a66aa54c21d1393ec)
- [gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-010.md#canonical-28408f9ee9e1d8ff6f1ce37bea5e41ab59397661411e134d64c8640e3b836856)
- [gcp.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-4c6101f3caa75e81e6a892525da8c929a8ff80591bfdd566d93ac8b7dfdd963f)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-a5035a3970d26e757895a07edf91e186034749f878ed0bc9380f4a89a0eb51b4)
- [gcp.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-ba3cbf3ac22df7aa55ff82cb6bec071600a413bdf4edee115662459b56451815)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-76751a759bd888723b322764cc48f95657f7c1454f183af2fcac56d1b0642102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8483639be9ecf11e50061d250c160c6a971e2b2283f0fa8c6a0d197c68634523"></a>

## gcp.not_managed.node_list.interface_list.bond_interface — gcp.not_managed.node_list.interface_list.bond_interface / 4e20405fd6bd / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.bond_interface

<a id="canonical-81ab931aff109756bcb3d5ac5c1a5393b1a06bc99dff87b09e8f43fa7c3d8269"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-e067111170e7b18c2d4c1abbc38c71e46c807cd68e08d614788d8e64390fe7f4"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.bond_interface / 4e20405fd6bd / 3

- [active_backup](resources--securemesh_site_v2--reference--group-009.md#canonical-afad8b0db84f3b161d654c0e7f5b6111af4f30133eeb08bcf725537f3c166763): complete subsection reference.

<a id="canonical-c286b0d5058d092127ffe5517d93d83956b0268e85587e34eb9702a036b47f1e"></a>

<a id="canonical-379397417aae8e6a13d8c83d168ccae49822d3ee78de85b558d7e3f78402c6c4"></a>

## devices property — gcp.not_managed.node_list.interface_list.bond_interface / 4e20405fd6bd / 4

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](resources--securemesh_site_v2--reference--group-009.md#canonical-e44b8cf5eec46db08438d59ab73ce122d9b783b8aa547bec16b5352f711e52bc): complete subsection reference.

<a id="canonical-fe468c65e1faf376bafae51deae5c772dd4997776b95b53f25d448349a6831f9"></a>

<a id="canonical-18db241df415568f880060475a87dc4743248e407b98bdb63c599307f1f1017c"></a>

## link_polling_interval property — gcp.not_managed.node_list.interface_list.bond_interface / 4e20405fd6bd / 5

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-49e8068b293fa02a3af1b1a05e299f6618cce4dbdcb9ec3144414100c0405406"></a>

<a id="canonical-deee283e9b30347342d3802a69cb9ee4daa16d0153e16447dec5dc4bd0d4abc7"></a>

## link_up_delay property — gcp.not_managed.node_list.interface_list.bond_interface / 4e20405fd6bd / 6

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-081462be5cfe83cba7e9a3a052c0faea4486bc6ae40d1dce09000bea4479c1f3"></a>

<a id="canonical-c19db0721a863dc11044ee9055ec14f4d6f14da9811a7d528e3ddb9bd8b3fafa"></a>

## name property — gcp.not_managed.node_list.interface_list.bond_interface / 4e20405fd6bd / 7

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-c7190f5c9515408ae141a2e00da576fe1ab08b56d1a5adf421925f1a73646196"></a>

## Next pages — gcp.not_managed.node_list.interface_list.bond_interface / 4e20405fd6bd / 8

- [gcp.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-009.md#canonical-afad8b0db84f3b161d654c0e7f5b6111af4f30133eeb08bcf725537f3c166763)
- [gcp.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-009.md#canonical-e44b8cf5eec46db08438d59ab73ce122d9b783b8aa547bec16b5352f711e52bc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-afad8b0db84f3b161d654c0e7f5b6111af4f30133eeb08bcf725537f3c166763"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1dde32dfa7ee4e1c579c912f5abef65df88ba6d410c0beeb6e333d81939c4cf"></a>

## gcp.not_managed.node_list.interface_list.bond_interface.active_backup — gcp.not_managed.node_list.interface_list.bond_interface.active_backup / 6af4661c00fa / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-76751a759bd888723b322764cc48f95657f7c1454f183af2fcac56d1b0642102)
- gcp.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-13d11816d3513aeb2e9652fbd5fdd93f9710d55fff8912b0a57d6d6c953e6309"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

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

Terraform syntax:

```terraform
active_backup = {}
```

<a id="canonical-415b13f5055e6d6443d4bbcd41bb33624a69afe20576c71c76a63465997ec1b2"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.bond_interface.active_backup / 6af4661c00fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1b05696d9aa51df6136af57fe75fb310a1315d719545c467b63fe93ca7d77c52"></a>

## Next pages — gcp.not_managed.node_list.interface_list.bond_interface.active_backup / 6af4661c00fa / 4

- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-76751a759bd888723b322764cc48f95657f7c1454f183af2fcac56d1b0642102)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e44b8cf5eec46db08438d59ab73ce122d9b783b8aa547bec16b5352f711e52bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-695fd2d7538872c7998021def51ce6d7782a712e2f498f0de74230aed2861c24"></a>

## gcp.not_managed.node_list.interface_list.bond_interface.lacp — gcp.not_managed.node_list.interface_list.bond_interface.lacp / 0aa843b08d9b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-76751a759bd888723b322764cc48f95657f7c1454f183af2fcac56d1b0642102)
- gcp.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-4d52074b8d701f7a97de568de7abe96ce4425d2a773d5f5d40fcdeddfa6b88a0"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
```

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

Terraform syntax:

```terraform
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-afb7b48cf34178fe576d99a03d8a0f77aeaafbdc6aec06ac585273089d8775cf"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.bond_interface.lacp / 0aa843b08d9b / 3

<a id="canonical-d80047b9faf53b43f5c1c40dd7d622f9ab7b08d429089b3748812933bf0db643"></a>

<a id="canonical-715da7046db025f26196647d0092fe7005cbec5602c86f0c4874db85e336f597"></a>

## rate property — gcp.not_managed.node_list.interface_list.bond_interface.lacp / 0aa843b08d9b / 4

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-7444b653b7441f37f438911b54a7b9b775f4721fba1ed89f234b99413cf96e15"></a>

## Next pages — gcp.not_managed.node_list.interface_list.bond_interface.lacp / 0aa843b08d9b / 5

- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-76751a759bd888723b322764cc48f95657f7c1454f183af2fcac56d1b0642102)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0f68f199f2f5898bfb2457b75b8388410f6935b7cc7a00c794f0efcf6e1d90e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222ada72c7aca42b7e6bf83227c6b2facefc8f9c82244875c658b67bebdfe94"></a>

## gcp.not_managed.node_list.interface_list.dhcp_client — gcp.not_managed.node_list.interface_list.dhcp_client / e6594b3e404c / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-e77709470d28284f48ec994a42fdb3a113480d945b817c087155c63ac130d66c"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
dhcp_client = {}
```

<a id="canonical-d2a0e0327c0e983ed0572e60afc291f4aaa2a9e941b373507ac89582113b9243"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.dhcp_client / e6594b3e404c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4f87b5c957d6c8c08413ec8e64b97d31593ffe0b6203b0bae398f62916627f66"></a>

## Next pages — gcp.not_managed.node_list.interface_list.dhcp_client / e6594b3e404c / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-298e91a1a499a505b3182e610c62cd0c6e67ed6dc77611d2e3b07a346ee018d7"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server — gcp.not_managed.node_list.interface_list.dhcp_server / e1bc4208b3cb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-62b6a4dea7afde0a55e2c4303fc8c8fcec1f48c567fd74eae815e01e6583799e"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-fefe4f14d02426e1da28875fe8ffa8fd3b0f7dab4a0937ed85ecba9bab5ac8c9"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.dhcp_server / e1bc4208b3cb / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-009.md#canonical-ec84b8400d8cdc3a1c0d58d9293583e0b842cd15e3898ac3a06828585c4b0cfb): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-009.md#canonical-c95a9952582ee9a81221f03bf60f0f2498b2b99be986078f7a81ba31e6649598): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-80cf78099ab1007691f60bba4cec30f0cf30324fd10393a471978bc4b5fd5e1e): complete subsection reference.

<a id="canonical-09879c7c4a8f17738381831ac3dd533f8d7222778e0fb2c95980c84d0b09d483"></a>

<a id="canonical-d4e0361232d116fc77e49db2c938ce86714e5b1b1d749410ca01f42d8a57a8f7"></a>

## dhcp_option82_tag property — gcp.not_managed.node_list.interface_list.dhcp_server / e1bc4208b3cb / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-09317d427dc32031c4f0c2f33e537bc9aae772d991853317c2b2dd4ce36316c0"></a>

<a id="canonical-ef19177be5d5c37afdc36ef4addf70e3a44faa77e16ee8360a8a5c3b5b0ac906"></a>

## fixed_ip_map property — gcp.not_managed.node_list.interface_list.dhcp_server / e1bc4208b3cb / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

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
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site_v2--reference--group-009.md#canonical-95a059a1a497899b02b2bace13d1cd58cc4a95e9e34a2278bfeb4f9da05cb3f4): complete subsection reference.

<a id="canonical-8052baf730a2c7916fa79b330f027a04426909b49dea277a4b312795d1a5bc0d"></a>

## Next pages — gcp.not_managed.node_list.interface_list.dhcp_server / e1bc4208b3cb / 6

- [gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-009.md#canonical-ec84b8400d8cdc3a1c0d58d9293583e0b842cd15e3898ac3a06828585c4b0cfb)
- [gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-009.md#canonical-c95a9952582ee9a81221f03bf60f0f2498b2b99be986078f7a81ba31e6649598)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-80cf78099ab1007691f60bba4cec30f0cf30324fd10393a471978bc4b5fd5e1e)
- [gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-009.md#canonical-95a059a1a497899b02b2bace13d1cd58cc4a95e9e34a2278bfeb4f9da05cb3f4)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ec84b8400d8cdc3a1c0d58d9293583e0b842cd15e3898ac3a06828585c4b0cfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf3c9233e594aba69ac55c60e0ff2d5102ed19306df3c22fe6f139d3929e2ebe"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / a58ec0b894d5 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-38e0816034e7ff2b9703da274f71db14faf201fd01f3fe1f18978ac750c0227b"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_end = {}
```

<a id="canonical-f1c033ddc7cb3fe1a4a5bea54f3ae9abbf99d20c22f154e58cf70d388915af66"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / a58ec0b894d5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c4d7a3bd88f7eaffaa0a65fd08629a06b99ab531eb2a4410a92380bc6cb63bd1"></a>

## Next pages — gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / a58ec0b894d5 / 4

- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c95a9952582ee9a81221f03bf60f0f2498b2b99be986078f7a81ba31e6649598"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d763510ba3e119f362122c14f82aa1e698f8d0f533fdae0ca6ec5649b0a5f66f"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / afc884af5310 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-aa680238a72d62fc3e27055d12422a729f58b783a88557f61a0749be206b0379"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic_from_start = {}
```

<a id="canonical-660aa387dbf3cf10e40d2417d087623dbdf36320f042b88e93d5a85ef9cadde6"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / afc884af5310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-050f37589efd4210aa4c74d9f68a0446f8ed3bc1a4532e9ad9282341628a94bd"></a>

## Next pages — gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / afc884af5310 / 4

- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-80cf78099ab1007691f60bba4cec30f0cf30324fd10393a471978bc4b5fd5e1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65993e7afc2f814951211a98cb578d5a4ad6763d165496edafcc6cae8099a33e"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 944e5d8cd85d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-9363fb8e2ad3f29d6e7231ecc9ef0cea861018ecb012f6222a6f6608aace0192"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
```

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

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-86d5b79d7164c4b64ffb1b9f2c61cf3178e98426627d59ce309ff150ad111c5f"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 944e5d8cd85d / 3

<a id="canonical-bb4fc891f7e03f5b7874b0a5323e7d7edf85a1c9895d4d0f1422cd2f70ea385c"></a>

<a id="canonical-1ea26e2ac378891a27439a20b3e035622ee30d47ae3c5c2b9fc8c89374cd2dbd"></a>

## dgw_address property — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 944e5d8cd85d / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-38b817df25018462ec5e0127cdfa33b3c0aaa30a8f050706a76102106f3f7096"></a>

<a id="canonical-b21726286def819f8d1db4302732bd66ab272cfadf16b79e30045049f68072bb"></a>

## dns_address property — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 944e5d8cd85d / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

- [first_address](resources--securemesh_site_v2--reference--group-009.md#canonical-360fdd94d5050281aac362de9a7e54d80839a72316fc7731ac5cbe18695b0656): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-009.md#canonical-75a81216967de9870dbe2db7870c58873860820c8c54f13462785d206adb4b9b): complete subsection reference.

<a id="canonical-64492ea095686e8e1a20e6802448b6a5a419e9f321db5132c1bdb7b9ca839f45"></a>

<a id="canonical-8ca235dc9af36abe56e5b5919e4d7698aba416bb88b25e9d576e31fe98a2c6a7"></a>

## network_prefix property — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 944e5d8cd85d / 6

Type: `"string"`. Optional.

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

<a id="canonical-21cb3c620727074bcac1be179159a14069c0399477b4888ec3802ae3acd10bd0"></a>

<a id="canonical-cc2eac5982a12f4d5090b97f263f0f621bff94e4c49872fc31dbe8fe15367a58"></a>

## pool_settings property — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 944e5d8cd85d / 7

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

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

- [pools](resources--securemesh_site_v2--reference--group-009.md#canonical-6c5e24a65b635a3a4d6f729d162f957df1f34eab0f79cf1a3d1e0e4162176199): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-009.md#canonical-dd161856d7f9efc78fb760d05957b6f8d88ae46e4378fece3e23e796c4c4a0eb): complete subsection reference.

<a id="canonical-28115b00e62360a2de12ff36334a627a6d501b652b2c163e0200d0cd29761e1f"></a>

## Next pages — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 944e5d8cd85d / 8

- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-009.md#canonical-360fdd94d5050281aac362de9a7e54d80839a72316fc7731ac5cbe18695b0656)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-009.md#canonical-75a81216967de9870dbe2db7870c58873860820c8c54f13462785d206adb4b9b)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-009.md#canonical-6c5e24a65b635a3a4d6f729d162f957df1f34eab0f79cf1a3d1e0e4162176199)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-009.md#canonical-dd161856d7f9efc78fb760d05957b6f8d88ae46e4378fece3e23e796c4c4a0eb)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-360fdd94d5050281aac362de9a7e54d80839a72316fc7731ac5cbe18695b0656"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b722d56c6bcc311922a12ff293d12d314e3a399cfd931410d2314bda96908d88"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address / 190d61e297ba / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-80cf78099ab1007691f60bba4cec30f0cf30324fd10393a471978bc4b5fd5e1e)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-b5ab8e8b198fd2fbdcbbdc4543f9da6658ae30f165f1f16080594df57a0b4de6"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
first_address = {}
```

<a id="canonical-533c07013c2436fc560d13e52e6a0a9911f98414fee7b1da7ae25647da8a337c"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address / 190d61e297ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-79148141ed374dede32da648ca10262cbe2a474d7f1bacc1d48d37257a222e51"></a>

## Next pages — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address / 190d61e297ba / 4

- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-80cf78099ab1007691f60bba4cec30f0cf30324fd10393a471978bc4b5fd5e1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-75a81216967de9870dbe2db7870c58873860820c8c54f13462785d206adb4b9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-145b47bfd1b74d0e4432a242506a5c2d20519e6c90bc21f203ce8ca919efd590"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address / a92b111424d9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-80cf78099ab1007691f60bba4cec30f0cf30324fd10393a471978bc4b5fd5e1e)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-b56245f339ce7a9df556a3b1332d1ff0bf9d272b1fd8d2cbd6d8fe57f3509736"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
last_address = {}
```

<a id="canonical-535db047ad1823d84a7cfaf52e963c01283b2e5a65894f43d1ee0615e35cf002"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address / a92b111424d9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53a577ce524e6f922c1319916d4b0b3941f1fb772cfc9aaf48eb053bb5536b10"></a>

## Next pages — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address / a92b111424d9 / 4

- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-80cf78099ab1007691f60bba4cec30f0cf30324fd10393a471978bc4b5fd5e1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6c5e24a65b635a3a4d6f729d162f957df1f34eab0f79cf1a3d1e0e4162176199"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2547a7a5100056c8ea4fe6fd5180dc5b0a7be9a35af7ed4c6ee647d9003a2b2"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / df4e1eaee1a6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-80cf78099ab1007691f60bba4cec30f0cf30324fd10393a471978bc4b5fd5e1e)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-4ee61943a1aa44629a9475ae51309bc9d0f410d7812e8fb2ca4ac4c5876adfcd"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-63cbf965baa514b5f9a68a235f992d8b4c95b2fff84b5f3b62acd8b0ec0f3c81"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / df4e1eaee1a6 / 3

<a id="canonical-7ca3188fe0eaf11b37cd6039f6be6aac8671c6a370e9bcb5db2c7d4580c2c434"></a>

<a id="canonical-7853f1f9291c03a515cfe746f01429ce7eea57cfeb843c5cd9a13b72e081654b"></a>

## end_ip property — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / df4e1eaee1a6 / 4

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-c7cde2a12d2c9b9ad3080edbc651e8745bdf47911905856bdc72d938b72deab5"></a>

<a id="canonical-f9e75911a69f4902f3461a1b23f2f62b4356a3b2355186d0461dd5f7f184f562"></a>

## exclude property — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / df4e1eaee1a6 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-44a1528254cdcd063ce86d04818d408d1444a89ea483fa10cc21be126738c9b9"></a>

<a id="canonical-6a14653ab3e28d6c2f9d108eff1caadf9fd35111da9f49a08f575bfbd659fd1e"></a>

## start_ip property — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / df4e1eaee1a6 / 6

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-660ce7db9df9f9e7a6eb29258c179fee451b3730f101a67df47b1eaccc784834"></a>

## Next pages — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / df4e1eaee1a6 / 7

- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-80cf78099ab1007691f60bba4cec30f0cf30324fd10393a471978bc4b5fd5e1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-dd161856d7f9efc78fb760d05957b6f8d88ae46e4378fece3e23e796c4c4a0eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fe825ec66d728fb05a56bb77bfe5e42707f1d1fb5636a43a4e417f081d3af1b"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / ea44c2a913b0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-80cf78099ab1007691f60bba4cec30f0cf30324fd10393a471978bc4b5fd5e1e)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-b583fb65a9896e3337c291ceb3fb285a3a717b8dc4b9973e5c38f3aff95f6d9a"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
same_as_dgw = {}
```

<a id="canonical-c12284cdeacb83b06c919e5ebe032e77c640eb53c751f213e517cb27fb9932ff"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / ea44c2a913b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6a0ab67bc9bf5176d17b1f5b17da3355d7ff0e00666a13d7e3eaa63d7eaa3966"></a>

## Next pages — gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / ea44c2a913b0 / 4

- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-80cf78099ab1007691f60bba4cec30f0cf30324fd10393a471978bc4b5fd5e1e)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-95a059a1a497899b02b2bace13d1cd58cc4a95e9e34a2278bfeb4f9da05cb3f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eddda1d9b52ec2e2cc8ed487ae124acd9a889850918d6e7ae7a4044286afaf0e"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / f64f45bdb117 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-77cb4965170024266cc889637c77f44f6238d987686844817aee1a83e646b844"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-1897a96717a377a72ee45c888e4e90d4f46a3025e14390534ea66f55a359504e"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / f64f45bdb117 / 3

<a id="canonical-0cc3132358427087ce27f6bc751f31e861ead2c9c5c78e33b1fc7c855d5e14c1"></a>

<a id="canonical-1ba0f5dadada5e943929b4d5e4ffd21b459133e12c58e9704892cf8cc63293e7"></a>

## interface_ip_map property — gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / f64f45bdb117 / 4

Type: `["map", "string"]`. Optional.

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

<a id="canonical-d890a4152b71cc5343eba4156b2fc8dd7eda8a8d01497cae04bc6a8e7b3e6fd7"></a>

## Next pages — gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / f64f45bdb117 / 5

- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-f7f0a35eeb84589b7f9181cc294e767417663b73945c3972cf6a03c86400e537)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a5faa0ac1acdae57e16b8dfff475700551e589d1639ef0d98fd2d17746b0f5d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23254024c73a96ba721cc0468006b62ef62797715224f452e6336dc9d85c7dde"></a>

## gcp.not_managed.node_list.interface_list.ethernet_interface — gcp.not_managed.node_list.interface_list.ethernet_interface / 4b1f963637d0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0b7f6d47484987878d60b48a6526bc3c4b673416ab8da56992eb6a1695046f51"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
```

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

Terraform syntax:

```terraform
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-d30f984f14dbeedd45a819c16f05fab38b57be6d25204d4d1cc647feeb9f58e3"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ethernet_interface / 4b1f963637d0 / 3

<a id="canonical-a275737aa286fff6909a6eb0188148f4247f791e9854fec66e7bbc3bfe9837ad"></a>

<a id="canonical-34879231dc4aafaf65c46747ed9d3044733dee1fdc2d50d0d786e510d1e17da3"></a>

## device property — gcp.not_managed.node_list.interface_list.ethernet_interface / 4b1f963637d0 / 4

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

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
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-8fba613ee71419d8f8d16f054a54f5785635de067006c38dd0a5cb27daee0c26"></a>

<a id="canonical-264be1969d2f764a267c42292977481e7bfa13e4c6858705bc65feae99e7c937"></a>

## mac property — gcp.not_managed.node_list.interface_list.ethernet_interface / 4b1f963637d0 / 5

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-79d4c2c9658447d2d4ece7c044e9e177a934db92387ceb1debe8a15c3c692fea"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ethernet_interface / 4b1f963637d0 / 6

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74f26da748eb9023a97422472355f565bf108944bb82cd690bd18277a97b96df"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config — gcp.not_managed.node_list.interface_list.ipv6_auto_config / 18ad108b0571 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-d9cd38d051ee6b39ec1be7b12e46a700da1b62009f1dc65bc7f54862afbb90ff"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
```

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

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-66cebf886c40c2aa013fa0e9b8f8923493c6c823cb042c36fd7e6475167ab664"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config / 18ad108b0571 / 3

- [host](resources--securemesh_site_v2--reference--group-009.md#canonical-96258308edfc82105d811913025df2d52d9a94d3ae16cc49b8c5faea0c131dda): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148): complete subsection reference.

<a id="canonical-6ffa93f76eb85e21638cb468051a60c144d8b0fcc5ffd617b8c1e392dbc2aa78"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config / 18ad108b0571 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-009.md#canonical-96258308edfc82105d811913025df2d52d9a94d3ae16cc49b8c5faea0c131dda)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-96258308edfc82105d811913025df2d52d9a94d3ae16cc49b8c5faea0c131dda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7c2116a2a8bee35dc7761df7c4fbf7bc92aeecb5d646f830f807e0b603cd2a2"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.host — gcp.not_managed.node_list.interface_list.ipv6_auto_config.host / b995f168b743 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-32f02be131b7c962ecc378fcdd602005dc303012f0b1d5b84bfc3caa36a1f44f"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
host = {}
```

<a id="canonical-ea181de1843e9af32c8dd256b7648fa3734a7befb22a5cecaab104f50420bd7f"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.host / b995f168b743 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-633af8a8ad476b41fbd5186dab88f10a6d24da5dc03aa97e796890d4db4f6676"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.host / b995f168b743 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c6e593f470554804e83e8f1887e517a4cb7468be5b9128aa333a5d7841296be"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router / 950649274dcc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-e4581fc33c26d3efe72d6ab04db5fe95b47c0371567d29f64a710e479d1b6519"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
```

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

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

<a id="canonical-57167309e15c5b2adb5eeb64029be0f114ade2556f71f464662ff84c1755ce0c"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router / 950649274dcc / 3

- [dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-410d30e69d43ec7b4ee98cd68fe829c94988a784c599dfdb63d2dfc500f7e8c5): complete subsection reference.

<a id="canonical-b19bea9972af5c8431e46b53f2536e6058eff939291e557bbf70692906a3051c"></a>

<a id="canonical-1e415baa8d639545d9f131872ad65c751f3a13db500bb4af875b0170763544ed"></a>

## network_prefix property — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router / 950649274dcc / 4

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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

- [stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c): complete subsection reference.

<a id="canonical-666c3fd6d54f908437ec264233dd1d4ef49e1dfc66f7c6ddbab5ba92d81818fd"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router / 950649274dcc / 5

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-410d30e69d43ec7b4ee98cd68fe829c94988a784c599dfdb63d2dfc500f7e8c5)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-410d30e69d43ec7b4ee98cd68fe829c94988a784c599dfdb63d2dfc500f7e8c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-929f5035c00b8c79b6d9ec80cb88ea8360f6af55b4ae12d5e87e07cd87332dc6"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 3e6d9725d070 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-a714785183b09339fb42751ddba4bd3041582204ca3bec54a94f0a67f60c6678"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
```

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

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-8105c46c2b48046bd9496fd448e5c1bba0463939fc02035af6599ffa2a5e2fa4"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 3e6d9725d070 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-009.md#canonical-123b43236a4a6322d250d109455de9146d042c9c54696c85c7c08a0e92d911cf): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-d67c3b0a79b1af60a701cabf0f13a6fd39540d8099bbf0e4bad8625390b3ac04): complete subsection reference.

<a id="canonical-fc4405d3665bae7a87846de8e15f4233e56b3f597e63e6a4ad46f6036ef8d967"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 3e6d9725d070 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-009.md#canonical-123b43236a4a6322d250d109455de9146d042c9c54696c85c7c08a0e92d911cf)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-d67c3b0a79b1af60a701cabf0f13a6fd39540d8099bbf0e4bad8625390b3ac04)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-123b43236a4a6322d250d109455de9146d042c9c54696c85c7c08a0e92d911cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-472d95e74818b12b569a70b3cb8636ed3f1eaaf136b2faa28d1d0a5f084f34a1"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / f156cff0debe / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-410d30e69d43ec7b4ee98cd68fe829c94988a784c599dfdb63d2dfc500f7e8c5)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-db919052cbce75bd3ddff7ecd736256c762716112f2ae02c7ae94a845c2e6e2a"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
```

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

Terraform syntax:

```terraform
configured_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-5e2603e2494ffbd2d5fcca1060ac6dce7b14773c45143a1c54305a36c3cab1c7"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / f156cff0debe / 3

<a id="canonical-56b7375a8e026b5f993f4e2cf4dbeb08bfe74daa976a471be6fed4cf4c21e92b"></a>

<a id="canonical-abac67121c84389a60ec0449e870124c6fbc5bc853bdbdab556602f27e3bbc63"></a>

## dns_list property — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / f156cff0debe / 4

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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

<a id="canonical-cca5b368554ad9e19bf87a63565b00e98a3c47c0ab7c231ca427273dfc891caf"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / f156cff0debe / 5

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-410d30e69d43ec7b4ee98cd68fe829c94988a784c599dfdb63d2dfc500f7e8c5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d67c3b0a79b1af60a701cabf0f13a6fd39540d8099bbf0e4bad8625390b3ac04"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-070ba2764d489159a021836930b1f6c1e53b90194f6c3db7b1f6b82afb9f3da9"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 145d388ca1eb / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-410d30e69d43ec7b4ee98cd68fe829c94988a784c599dfdb63d2dfc500f7e8c5)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-7f895151badf5e0cfaa673ce50e1c9dc0a848b9b37d4e166eced3ad6339ba93a"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
```

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

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-30840557a596ba003670b3659ddc1f5e0b6b8d658eb0627c60bb4e27793d5948"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 145d388ca1eb / 3

<a id="canonical-2c033f49fc0110e75ea50bbace48751369ceed97d8beb2c3e5ce67a20c20701a"></a>

<a id="canonical-f4ea45234960624c4e772296013b4b444f4d2a44ab01fcea71051832d20f09dd"></a>

## configured_address property — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 145d388ca1eb / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

- [first_address](resources--securemesh_site_v2--reference--group-009.md#canonical-6e6ab38b875a9c8d1f5662f546a2760060ddc6fbb7375b8ac67dd5fb68bf32ee): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-009.md#canonical-a41c78e2cd705c760b1beb48ad39fed5289dc0206e17d2589524ff503936b0b3): complete subsection reference.

<a id="canonical-24642e83d850abdf44706cf39879a7422259ea7d2af2b73bb1cd32d2b1d82ecb"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 145d388ca1eb / 5

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-009.md#canonical-6e6ab38b875a9c8d1f5662f546a2760060ddc6fbb7375b8ac67dd5fb68bf32ee)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-009.md#canonical-a41c78e2cd705c760b1beb48ad39fed5289dc0206e17d2589524ff503936b0b3)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-410d30e69d43ec7b4ee98cd68fe829c94988a784c599dfdb63d2dfc500f7e8c5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6e6ab38b875a9c8d1f5662f546a2760060ddc6fbb7375b8ac67dd5fb68bf32ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c261657bc21c950cf4f81b2f04676fdffd6214332e42f66f40760408ef48918"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / fb13dab4a590 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-410d30e69d43ec7b4ee98cd68fe829c94988a784c599dfdb63d2dfc500f7e8c5)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-d67c3b0a79b1af60a701cabf0f13a6fd39540d8099bbf0e4bad8625390b3ac04)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-bb6316e10d454eb5e9912ffdb674c2a281fc26fcd753059259dd485b9156c294"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
first_address = {}
```

<a id="canonical-dd5401804addeb4a05c67c9354c33d4d5859a000937f4565f48ef2bfb7303369"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / fb13dab4a590 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cec183a79361efcee06643ec44d81f16086ab12a90b5664015334095c146c565"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / fb13dab4a590 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-d67c3b0a79b1af60a701cabf0f13a6fd39540d8099bbf0e4bad8625390b3ac04)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a41c78e2cd705c760b1beb48ad39fed5289dc0206e17d2589524ff503936b0b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-897f9bd9759b243242291cccc6cef55cc23c271dfe7a542961dcbb7511889018"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / ad36b5a55816 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-410d30e69d43ec7b4ee98cd68fe829c94988a784c599dfdb63d2dfc500f7e8c5)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-d67c3b0a79b1af60a701cabf0f13a6fd39540d8099bbf0e4bad8625390b3ac04)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-a915265332f80062d1eadbc058351745ff377e2ef402fb71f9acbefa6a89c912"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
last_address = {}
```

<a id="canonical-d53f84024861b1b2786973d1ffb7a24b8d15e0697d9f009899174a93a611fe82"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / ad36b5a55816 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe2ba8fa5b49699331173d29a055a35e8a94a54f5b8816316332f0f4e1e2f912"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / ad36b5a55816 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-d67c3b0a79b1af60a701cabf0f13a6fd39540d8099bbf0e4bad8625390b3ac04)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
