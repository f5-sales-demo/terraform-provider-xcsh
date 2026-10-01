---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-d705b47356e561848fc30193567caaa64d639cd75e640f222a4c10196b10673d"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 842b062259ec / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-c1f0be75e34a82ab353c8348ad97160dd4ab5899956c5eec49f9ea21f8f642ba"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

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
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
stateful {
  # Configure direct properties listed below.
}
```

<a id="canonical-3cc9704b7998143745265d91299fd7d76e42f07deee7333288a9db94697310b0"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 842b062259ec / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-010.md#canonical-e6e3448f315b0d584bd738b8ef4903f8b1399240acaad74b32e0dc5ed4840567): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-010.md#canonical-bffebd1ed4537b6d763db5b6be062096bec308fdb09879eefd683987a7cc0c3c): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-990ab42e5b2524c427e4ab315abcbc4672fc905add04ae66113db8ae693a0d5f): complete subsection reference.

<a id="canonical-09f85346d4c997141baea0f73d819e60052188a3060a1b7a79ffb0b076b726a8"></a>

<a id="canonical-9a95cf67a35c98650c28db345474e1164c83e63ac0972d64f361c1e885c1b087"></a>

## fixed_ip_map property — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 842b062259ec / 4

Type: `["map", "string"]`. Optional.

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-010.md#canonical-9bb3257d2bf4e14cfdea2d22a23de58b137d2ecaeff23cfb60896fe0928ff357): complete subsection reference.

<a id="canonical-d8cf5a94f3ca3100fb19362fc4c8e1d4e5fab7d13cce2b719d7b9c6bb0cc6331"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful / 842b062259ec / 5

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-010.md#canonical-e6e3448f315b0d584bd738b8ef4903f8b1399240acaad74b32e0dc5ed4840567)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-010.md#canonical-bffebd1ed4537b6d763db5b6be062096bec308fdb09879eefd683987a7cc0c3c)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-990ab42e5b2524c427e4ab315abcbc4672fc905add04ae66113db8ae693a0d5f)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-010.md#canonical-9bb3257d2bf4e14cfdea2d22a23de58b137d2ecaeff23cfb60896fe0928ff357)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e6e3448f315b0d584bd738b8ef4903f8b1399240acaad74b32e0dc5ed4840567"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d733386a706475b7841feb4d9c02f5be9cdc52071b72531ef930d6e53f77b743"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / f5d596438c9b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-2673375d5226fdf08bfc726e65157d4e986dcc05e4e2211919efdf96c5c81c09"></a>

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

<a id="canonical-847e95e4364170ad1971fbbd9119d94bf6a90d3b9c4230213dd2dc4dfb2b8925"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / f5d596438c9b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f360833ecaffb8ba322f8063ce26ee3932b4434ae7df438cba495c7711b98cd8"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / f5d596438c9b / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-bffebd1ed4537b6d763db5b6be062096bec308fdb09879eefd683987a7cc0c3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd5ae7d52222a08f1ed28fdf063c96d5e9c33159c56cd087902c4670c8c7f889"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / b6c22b0f7d89 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-b3855c90f1bbf54bf1c7d0a2c717ad32c91d978dfbadfec1cccf7eb3814b8f9c"></a>

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

<a id="canonical-5cc14f2de925cf4a7633c1b9f7688ad3ab50e22c66c14fbd1aa56a19e4fda212"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / b6c22b0f7d89 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4f6bba8aa346545787eaf7ef22b4ee9c94e917740c17eb7524c4278e16f61632"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automa / b6c22b0f7d89 / 4

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-990ab42e5b2524c427e4ab315abcbc4672fc905add04ae66113db8ae693a0d5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47ec0425fa34760e68d793610442d12dbf5b757d7002f7ede46e8a7fc9904899"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 4e7cc00edaa8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-ce773cdbcd25d18573fbf67a8942d15b9e7b1c0b7a214b5d205bd9d127e8b316"></a>

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

<a id="canonical-0652866382353d56d113cc8a4cbf6b40ae44b532c1315849cf486fe627c519f6"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 4e7cc00edaa8 / 3

<a id="canonical-b251feb4726b77b75db87979c1cd9856c4cac77336d1fd477f8d2c958b1b082a"></a>

<a id="canonical-1e42afd90981ab7760b816e1a21a5b26b7b425c1f36feff0d8ab527c89c4b7bc"></a>

## network_prefix property — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 4e7cc00edaa8 / 4

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

<a id="canonical-2ae233cf0f54d7fd2852ee12ee45f120f27279aecf5a54dc70a1023e104eaba2"></a>

<a id="canonical-01d41603a4b3efacb2458286efcbb19e519b2ba8e0f4501e7279f2db023b0734"></a>

## pool_settings property — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 4e7cc00edaa8 / 5

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

- [pools](resources--securemesh_site_v2--reference--group-010.md#canonical-6c0a50cf4cb572c1c1a130ad63c653152632ce38a39cbeb23f06f06c7bf9fa51): complete subsection reference.

<a id="canonical-b4294943e74b34add8bab0c0e4b7c2da8828e3d5b9a5427a4af40587eb4f08ab"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 4e7cc00edaa8 / 6

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-010.md#canonical-6c0a50cf4cb572c1c1a130ad63c653152632ce38a39cbeb23f06f06c7bf9fa51)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6c0a50cf4cb572c1c1a130ad63c653152632ce38a39cbeb23f06f06c7bf9fa51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de126607cecfa50d5db92ec88e5a5a3a9831b012c25800680f584365b43e0427"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 95100087e99e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-990ab42e5b2524c427e4ab315abcbc4672fc905add04ae66113db8ae693a0d5f)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-71f55a72ad79d085e21ac68dca858c7833329096d5bcd10230bf8910268c2b15"></a>

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

<a id="canonical-aca16cbe8f9c54441d9f7c7f63506e0390896189e85b2809f1593eb21629d24e"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 95100087e99e / 3

<a id="canonical-cd19d70abd9512b6b1414775fe617715993071a74f626c58d8fffeddef035edb"></a>

<a id="canonical-d981a4933396e9110383a70f152ce7aa12f82cce2d68db09f03c6296a9de3692"></a>

## end_ip property — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 95100087e99e / 4

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

<a id="canonical-3024bda7772792f0a0bd03e22f890d3ca38086486277440b581656fbcc77bbb4"></a>

<a id="canonical-8bc44a349f629646aceb0fcfc2a21d68405a81ed18f11bde882b6a2ff391b5ff"></a>

## start_ip property — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 95100087e99e / 5

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

<a id="canonical-cd85e109d4aa436bd1e07047562cccaf92e3a082f295b606263adb61a0c61284"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_n / 95100087e99e / 6

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-990ab42e5b2524c427e4ab315abcbc4672fc905add04ae66113db8ae693a0d5f)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9bb3257d2bf4e14cfdea2d22a23de58b137d2ecaeff23cfb60896fe0928ff357"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d660349bf80ca45f06dbc61f33c14b2c445c5c41743f3f47a9d28fc0e0f558c5"></a>

## gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 14ed2a99d302 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0c02fcc33b1b10036e58fbd57a089b7ee30dc31ec6b8346878b954a0fa068871)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-58f001dbdf67c21b4232f189bc5ff9e0d88a4f9c073d8219518fb11a74afe148)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c)
- gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-07106c795bd75be6c1473e754475518cb33880e6228d89e7baf56c71802c8030"></a>

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

<a id="canonical-1216cb5737484e123785635de2dc4b2ff87ec2ee8113c532c06afedae7cd0101"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 14ed2a99d302 / 3

<a id="canonical-7e0636579dce31006c7ff02b27541421faaf3f4bc20b1509cfec913ff6601301"></a>

<a id="canonical-80907a21dfb2b39db84c00473c8e02dfb325eba94d288561666037efc0e9a38e"></a>

## interface_ip_map property — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 14ed2a99d302 / 4

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

<a id="canonical-0edb7268138d49ac9072bad3fb3d43464fdf7b90a867c46ccac6ef0c18cc35c1"></a>

## Next pages — gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interf / 14ed2a99d302 / 5

- [gcp.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-eb98ce95c2c91604c29df18f28e2583dbc4a0fb1f99c8cbf62f33cb4c2592f9c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-65af5831979ef12c631b1c984ac3d4fb2563e5604c340073aaa66f8567da5694"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04ecb70172d4c03216dc6296369438aa58c994e3a5654b528ac2283aa24c51a1"></a>

## gcp.not_managed.node_list.interface_list.monitor — gcp.not_managed.node_list.interface_list.monitor / da115706a9a0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.monitor

<a id="canonical-c57d64557a7f31e76825cf3f87b9f48ea57fb5c673f800f66fdbc8988106cd16"></a>

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

<a id="canonical-516f027454b2dfbb082e5c8bd4c7c386909dafe918f646db97c7ac71f09444a2"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.monitor / da115706a9a0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5084be2d4552ccf8fa27ab95b218202e2f56ded97029feca56acd798b5872001"></a>

## Next pages — gcp.not_managed.node_list.interface_list.monitor / da115706a9a0 / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-941d0f2123dd29a00e7e3d87119b1448b769fbf2a857a44c2af363899501710b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a081c6461d24abf16626e40f4b2ab32f5b4e760b78b0fddf04dd39ff22792045"></a>

## gcp.not_managed.node_list.interface_list.monitor_disabled — gcp.not_managed.node_list.interface_list.monitor_disabled / 5f3a18f369be / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-da7fb5520a8eeaee1194ff04dbf8a55927faaacb8d212cda47f148dc11dbddde"></a>

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

<a id="canonical-ec2f0cefee49bfde499d551c187e92f1954743d46d140f17822513f2ec51fd21"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.monitor_disabled / 5f3a18f369be / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5b229b75828e2d5d34ce8d936c982060742373c8868ecf5c3df9c51891f84b10"></a>

## Next pages — gcp.not_managed.node_list.interface_list.monitor_disabled / 5f3a18f369be / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1cacfff659665dd3d788ff4eafe976e03e12c63d44015cf124750f19beeabbeb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4054b9e8bbdb09d4b613c92c01bd4398e9ca9e2e7c5ae3041b0e983b57b3018"></a>

## gcp.not_managed.node_list.interface_list.network_option — gcp.not_managed.node_list.interface_list.network_option / a55e5ab6c44e / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.network_option

<a id="canonical-7cd3c52c3317f7b1acd64ab17dddccefcc95956f0d35d201ae47ccf3d770fe55"></a>

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

<a id="canonical-0224774e94ba0f7e2312398bf202236018fd0890bd505d9438371d7caf07f4c4"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.network_option / a55e5ab6c44e / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-010.md#canonical-ffe221a20240157413f17016d6dfcdb681535ab1e30989a121c22d8318ee2f14): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-010.md#canonical-42480964a215ae4dd01e797ba038bf050701c83106bf889cb88c3d85c90d4061): complete subsection reference.

<a id="canonical-cd1c6d808d9cadc5b9f2ac606c100da82dd25f8c957bf1715b6eef499392d4a3"></a>

## Next pages — gcp.not_managed.node_list.interface_list.network_option / a55e5ab6c44e / 4

- [gcp.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-010.md#canonical-ffe221a20240157413f17016d6dfcdb681535ab1e30989a121c22d8318ee2f14)
- [gcp.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-010.md#canonical-42480964a215ae4dd01e797ba038bf050701c83106bf889cb88c3d85c90d4061)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ffe221a20240157413f17016d6dfcdb681535ab1e30989a121c22d8318ee2f14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31b7909b673d89fb017f1e9c29a047d5d9d5bff1891fb3c37262b2068e4ec27e"></a>

## gcp.not_managed.node_list.interface_list.network_option.site_local_inside_network — gcp.not_managed.node_list.interface_list.network_option.site_local_inside_networ / 195c816ca105 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-1cacfff659665dd3d788ff4eafe976e03e12c63d44015cf124750f19beeabbeb)
- gcp.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-ee36bc686741ed5856f093b8d6e333394ad7afcc4bb4aa2279ebf1fd83d9eefd"></a>

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

<a id="canonical-53e7a824debf74bccd9cc8f952d9161e12560585b4d72e9a36af6b3e29eb537a"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.network_option.site_local_inside_networ / 195c816ca105 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ae8b73d735737cba7ad88cd5a0f165c6d8cba6e42bdf113957cc2f6564810586"></a>

## Next pages — gcp.not_managed.node_list.interface_list.network_option.site_local_inside_networ / 195c816ca105 / 4

- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-1cacfff659665dd3d788ff4eafe976e03e12c63d44015cf124750f19beeabbeb)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-42480964a215ae4dd01e797ba038bf050701c83106bf889cb88c3d85c90d4061"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1300c4894031e90ff43ce4501108ef798146fb9f8b1d9092d4145900996ff8d"></a>

## gcp.not_managed.node_list.interface_list.network_option.site_local_network — gcp.not_managed.node_list.interface_list.network_option.site_local_network / 17082117029d / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-1cacfff659665dd3d788ff4eafe976e03e12c63d44015cf124750f19beeabbeb)
- gcp.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-5171a07ce0dded5206fde478460db8dc490b72e5d144e121251f6db65ef8a3b0"></a>

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

<a id="canonical-50229f3511d57345f98c6347da15857302b2c8fdd1fd5a20b3209ce1c4657a5f"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.network_option.site_local_network / 17082117029d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e4dd791c34a8234b9f07e12ee587cf4eb692846c0564b9133af211a78a81be0"></a>

## Next pages — gcp.not_managed.node_list.interface_list.network_option.site_local_network / 17082117029d / 4

- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-1cacfff659665dd3d788ff4eafe976e03e12c63d44015cf124750f19beeabbeb)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-63ee8686a1ea8c2c2c0b440cfc37be7530dd315db0775b54c78f7c48b846c7cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-913be02592fe860e04ab203829981fc8b5655e8b56d0ecbf923686ade3b77f6a"></a>

## gcp.not_managed.node_list.interface_list.no_ipv4_address — gcp.not_managed.node_list.interface_list.no_ipv4_address / 20c8f770de6b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-690d7bbad680ac646747dd148225f3510142777b93d7db6133a99ca6b9e8e277"></a>

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

<a id="canonical-54af735ba92fc5dc65c1f001b92b6ae5d694251fe25f440dd7f146e21f5fa9d8"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.no_ipv4_address / 20c8f770de6b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-915cdac6fb05e4634cbe548852889a6143f3eb0a944afa529a904d5f334d6d2a"></a>

## Next pages — gcp.not_managed.node_list.interface_list.no_ipv4_address / 20c8f770de6b / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c04cbb76b113bf3f3f0f20027a03b6ee82a3cc0d7d581852b9ff3ce3812f3296"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4017f7f067244da59383a4eebf2a9e9535c356877a8f084aafc8ee1851ca904"></a>

## gcp.not_managed.node_list.interface_list.no_ipv6_address — gcp.not_managed.node_list.interface_list.no_ipv6_address / 7a2eec2ea139 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-aeb47edfebcab107b332b3f700ad8e6df85dcd34c97163687f2249acef1a1ba2"></a>

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

<a id="canonical-80f483f780a39cc57ba401244e20c828a2fd1f9a47cabdcf215d3d824e26ba13"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.no_ipv6_address / 7a2eec2ea139 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c7e4252a0f253acf6d8e3d11f30d7338a1033d411644387ca9adda74bcf95f46"></a>

## Next pages — gcp.not_managed.node_list.interface_list.no_ipv6_address / 7a2eec2ea139 / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5f86f36cff19c3cd29fb66ddc395fab7476bf1bd10b6875a66aa54c21d1393ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dd48811842849b7dea08ce6821650ea4c85a6a32181fc03138839dc13c7239d"></a>

## gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / aa5aca053e60 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-e1b08a5efa793d2817f0f40d8c430bb9331887bf26a910bc90f1040254dca70a"></a>

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

<a id="canonical-99f4a8f6ae133964cc11ef96e69e59dec36363dadf9b64061f9de1fb05c37041"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / aa5aca053e60 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-580507e0bc9b0174d297f26c07edbfb845c99b122aa9e4948d86053aa433d1a8"></a>

## Next pages — gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_dis / aa5aca053e60 / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-28408f9ee9e1d8ff6f1ce37bea5e41ab59397661411e134d64c8640e3b836856"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cc597058dbf6a848423cda771b136bcb00fd76446521c5d9842875e9c66a0e1"></a>

## gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / 98f4197e2273 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-3f4108a04e6aee02b8f51acec4c6315a7f24fa39f2e4c6681f03bad0b1afe5a8"></a>

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

<a id="canonical-94466d338fb3daf0e74b102f89f204a7796558fcd5b32a397df6bca826d916a7"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / 98f4197e2273 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef5ba10203e75adba9fa7c44f836f574a00000d2fb1664fcb7f25f292be634aa"></a>

## Next pages — gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_ena / 98f4197e2273 / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4c6101f3caa75e81e6a892525da8c929a8ff80591bfdd566d93ac8b7dfdd963f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8980d9c4c6fba97b73c4f357117ab9b00a092d7d31f4d954583029fd23f3ca3"></a>

## gcp.not_managed.node_list.interface_list.static_ip — gcp.not_managed.node_list.interface_list.static_ip / caaea05c8a45 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.static_ip

<a id="canonical-fab8e18bd5ad2d89a8d303655b5a4a66dda52b188ae3bce87d1d0450c5242642"></a>

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

<a id="canonical-f29364dd6247134178ebe96b899e5c7705a8185eba558245eea6e7d45a316efe"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.static_ip / caaea05c8a45 / 3

<a id="canonical-063b9b4cadbad73fc622408762863f0cd2475a7ed039110a093b3550c508f87c"></a>

<a id="canonical-21d13df746ac069b09e1cfadffd88321bbd57f3d4eab213d41317dfee6c2c2a3"></a>

## default_gw property — gcp.not_managed.node_list.interface_list.static_ip / caaea05c8a45 / 4

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

<a id="canonical-c8615ba492513ee204b6b64e639e7a9606634b38fd24863b565b9b596cfd4a14"></a>

<a id="canonical-770a098ec740e78a10296c1e4e6ebb7bd1f15d434e2611fb81a80ce4b9211168"></a>

## dns_server property — gcp.not_managed.node_list.interface_list.static_ip / caaea05c8a45 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-d7e035bb8b824a962d4f8cc6c3fc0f3e988575bb5b718d98553f3af834adaa75"></a>

<a id="canonical-6667cafe517f1da493dbb38021c6570113bdb4edd25912d4f7ac87ea9db60657"></a>

## ip_address property — gcp.not_managed.node_list.interface_list.static_ip / caaea05c8a45 / 6

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

<a id="canonical-5a0b92e59308e7cb250f54ee0d19a02280600c2832d8fc1960ccb8817db2a445"></a>

## Next pages — gcp.not_managed.node_list.interface_list.static_ip / caaea05c8a45 / 7

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-a5035a3970d26e757895a07edf91e186034749f878ed0bc9380f4a89a0eb51b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44cf1029e8c7ad7776241ec4a8eafb2ae31dd221960fa0c6c153befe94ea43b5"></a>

## gcp.not_managed.node_list.interface_list.static_ipv6_address — gcp.not_managed.node_list.interface_list.static_ipv6_address / d18631947002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-e3856f879b9d3c1fd543686a1abbb5475debc4fff512a637cd10ce27449a81ae"></a>

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

<a id="canonical-627afba89aaca19af13edf1f7d635cf8b6f69414520e28d71422e53490a5d645"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.static_ipv6_address / d18631947002 / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-c8238913f22ab4d0e2e9639d755b974d3795293eb09c0cf0603faae86eaaa982): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-e6d9db7753921660aeb145d5ca40b0e38cb16bf1c28ca78f49cda078a6a77271): complete subsection reference.

<a id="canonical-7f424b94e3bec63f74740726cb86f436697f8d3e38bf7e713aef8f4b9fb478b3"></a>

## Next pages — gcp.not_managed.node_list.interface_list.static_ipv6_address / d18631947002 / 4

- [gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-c8238913f22ab4d0e2e9639d755b974d3795293eb09c0cf0603faae86eaaa982)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-e6d9db7753921660aeb145d5ca40b0e38cb16bf1c28ca78f49cda078a6a77271)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-c8238913f22ab4d0e2e9639d755b974d3795293eb09c0cf0603faae86eaaa982"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e02deb8bae26f866f4d000c48e6cd16887ce0d73c1d2d00a9c85d7c90d44ae12"></a>

## gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / f448f4255d12 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-a5035a3970d26e757895a07edf91e186034749f878ed0bc9380f4a89a0eb51b4)
- gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-beaadcdcd47dc078da4a32ee2603ef2335d1fe75b810ac9178e361a469479e53"></a>

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

<a id="canonical-7b7760da68cd204995aa018105a96ad2f841a38702333259ab880cc80679d002"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / f448f4255d12 / 3

<a id="canonical-05fbc100ff5b27b3ff7b55f47dfc223c94cdd5957b5de44e8f081fa3fb32b739"></a>

<a id="canonical-87591c3755dd7f5a58903ef81667dd8ccb39f85f9def57634727a632e9d0b106"></a>

## interface_ip_map property — gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / f448f4255d12 / 4

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

<a id="canonical-da1bda21ff140fc71b06f292bdd17bb03906f95725556e787e044f2bfa932fcf"></a>

## Next pages — gcp.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip / f448f4255d12 / 5

- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-a5035a3970d26e757895a07edf91e186034749f878ed0bc9380f4a89a0eb51b4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-e6d9db7753921660aeb145d5ca40b0e38cb16bf1c28ca78f49cda078a6a77271"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a84342a3d77405ab5a2b1ff320ffb5121b45ebb04688c90cdd71c1d85033ea2"></a>

## gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6b90d0fbf8df / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-a5035a3970d26e757895a07edf91e186034749f878ed0bc9380f4a89a0eb51b4)
- gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-3a7bfaf011888b5a7a10dc315f5a66d68934ba9f479146e622248ee17e3f461c"></a>

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

<a id="canonical-5791d16d356b9604eda6b4b839add01a596fd7234896b01c73a4c13df7221006"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6b90d0fbf8df / 3

<a id="canonical-9f5fb24f89b66d3c8809ac34cde7ba120b3e1c299eb5b6cc7735b07c37cd4bc0"></a>

<a id="canonical-fa4f9cc22c1a9e449912251a0071d5b631f479277db5400b2e434f68a327c76e"></a>

## default_gw property — gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6b90d0fbf8df / 4

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

<a id="canonical-a31182790669e78b3506b3854bca1040bba3639cfc2d9476c465c03e23224b60"></a>

<a id="canonical-1610e2e9c68525e0d3dc29b2de0853bf82be48fbb7739e296688fe4e2161ac83"></a>

## dns_server property — gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6b90d0fbf8df / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-d50594ede9f0320258278f8508ca5097d96d79aeef939d96f21a341954817f8b"></a>

<a id="canonical-c3229a8364463ec7d423114a419d192700f74624fc5424c6c24020fc017b137e"></a>

## ip_address property — gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6b90d0fbf8df / 6

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

<a id="canonical-3c47da0c3a7f7df39e5d9e02c5a753934d73eb0cf3be2384765d135f78411e43"></a>

## Next pages — gcp.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip / 6b90d0fbf8df / 7

- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-a5035a3970d26e757895a07edf91e186034749f878ed0bc9380f4a89a0eb51b4)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-ba3cbf3ac22df7aa55ff82cb6bec071600a413bdf4edee115662459b56451815"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1f4685815ebc5fd1ec29ba4c648114498205b01eb18ac43a3c4a02a74b80116"></a>

## gcp.not_managed.node_list.interface_list.vlan_interface — gcp.not_managed.node_list.interface_list.vlan_interface / bcddea35f783 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-6eb51272a66d083179f0429349942651b103832c0330cea381599982fac4a046)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-d3c7fa797f3e948345d5028478c196b081a25e5e865544a6504c90e57c614011)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-b5217da6e068f0c935089530c517b70b341c4f47af6095f8e2bfe45ac5e454fc)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- gcp.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-7e931e8bb323ea58423bd3e081517bd9ab0b499478a5c6d0b05c36dba59c3f34"></a>

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

<a id="canonical-f86e46fda55ed0cb4aa15e3e2b8923196ab1662dcb26c29a8e71fbbc295eeef7"></a>

## Direct properties — gcp.not_managed.node_list.interface_list.vlan_interface / bcddea35f783 / 3

<a id="canonical-72323d2b8869fcf4007e36a64f586fba7fa90b6c6581834f95c637546aa18edf"></a>

<a id="canonical-0848458d73baf03f719e130981ec26776738fdd6a9d19856ff8d6b0d2b4ced57"></a>

## device property — gcp.not_managed.node_list.interface_list.vlan_interface / bcddea35f783 / 4

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

<a id="canonical-6e47c3e2cb1da88ca564a228709f8ef0ba247347ebdc9bf4c8f04e049d953f4b"></a>

<a id="canonical-b08605229841cb371a1ea4496504ec14731177e58e16c88a3d31c5d22502e7c9"></a>

## vlan_id property — gcp.not_managed.node_list.interface_list.vlan_interface / bcddea35f783 / 5

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

<a id="canonical-c95cdbbf4e8141796e97ac0fd9afc8dd2bab99b9d53eab628121fa1dff21e136"></a>

## Next pages — gcp.not_managed.node_list.interface_list.vlan_interface / bcddea35f783 / 6

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-6e4e4080b770699cc23fd02dc56ebb2300668815b81d1eb36d13132b5d44dd7d)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea79b2aadde005e5744a357f2d157954df94546430aee4d6e1b9c882e1256f04"></a>

## kvm — kvm / e694e6bc246f / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- kvm

<a id="canonical-e016808bc2e6744ef819130102944bf65fe5208507a0e6537c6025615c3624d5"></a>

Type: `"object"`. single nested block, Optional.

KVM Provider Type. KVM Provider Type.

Upstream description:

KVM Provider Type.

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
kvm {
  # Configure direct properties listed below.
}
```

<a id="canonical-683da1e6e7b48f7a7f96784a2019446019148ef8d84c9699438d8d0b7f3910a6"></a>

## Direct properties — kvm / e694e6bc246f / 3

- [not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c): complete subsection reference.

<a id="canonical-67869157ac6d1787efc75c459c528886bbe7ec7f1a3b72b43887ec9a1fd9ee5b"></a>

## Next pages — kvm / e694e6bc246f / 4

- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa4cd51c0923848dccac2e693c0ae639b765425458ceb48dd22173b383f48da2"></a>

## kvm.not_managed — kvm.not_managed / b13e3202a223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- kvm.not_managed

<a id="canonical-683728571604308f15ed77491af2c71159d00de27f70f9e39576513aad2b9e67"></a>

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

<a id="canonical-ebc95367fc1d9c692622842d88712bf9b2af994d45e71efa3d9596e5dbf38a0e"></a>

## Direct properties — kvm.not_managed / b13e3202a223 / 3

- [node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596): complete subsection reference.

<a id="canonical-9efe48a9a780a200679b9ab9018c6f41fca7a771d2eeaff10b525c8f8e91a71e"></a>

## Next pages — kvm.not_managed / b13e3202a223 / 4

- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ccd767dbffc5a9ff6de4a6ffbb329966b7066581d28cabf51057c898e1fc1ae2"></a>

## kvm.not_managed.node_list — kvm.not_managed.node_list / 373ccc665875 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- kvm.not_managed.node_list

<a id="canonical-81a8c884c74a843d8c18a3ba4e4f503e7580824d5cab559a52282c540933805f"></a>

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

<a id="canonical-b5df9ddf8aca3ab6270f4cbfc54d830353456226e71fe6066ab38d3b5c10fbf2"></a>

## Direct properties — kvm.not_managed.node_list / 373ccc665875 / 3

<a id="canonical-38bbf36c42c8b7e6c176621487599eccc3a9d1bb7e42fe2f51e93fa995bde979"></a>

<a id="canonical-8bc1ec0a12cf207d8dd38778693a8182bb9ef021e99dae83096f837b14fadbd4"></a>

## hostname property — kvm.not_managed.node_list / 373ccc665875 / 4

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

- [interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223): complete subsection reference.

<a id="canonical-1badcfa33290d05fae130f4c31cf04fc30aa1dbdc22c2e9341174e9b5083d04f"></a>

<a id="canonical-1fbe6de6a33742ed26361b14e4d4aa40482977719fa1fe2c304768d3b4fe1fb3"></a>

## public_ip property — kvm.not_managed.node_list / 373ccc665875 / 5

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

<a id="canonical-ad79338763872c49da189d32849d4742e0f4064c05689b73094093a1dbf452bb"></a>

<a id="canonical-ede313d93fd0b4a16fe26ff961daa400ca3038aab34345273f84a375a07289dc"></a>

## type property — kvm.not_managed.node_list / 373ccc665875 / 6

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

<a id="canonical-dcdb7ad4025593bf0aec38d4c9828e687380428e474cfacc47d693f5fab88680"></a>

## Next pages — kvm.not_managed.node_list / 373ccc665875 / 7

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0df7f1c7357d80deef81a32f32d891233adf93c527e3a5173ecde9407c0c4720"></a>

## kvm.not_managed.node_list.interface_list — kvm.not_managed.node_list.interface_list / 9b3130b00fe0 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- kvm.not_managed.node_list.interface_list

<a id="canonical-eda177726202aad338a65bb5d09a010c16b98a0dc791f066f5c9c6c5d964c30f"></a>

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

<a id="canonical-cb6273d30fdc77ce594cd8f4356e937dc41c21aa9845a091e84ab15f79115335"></a>

## Direct properties — kvm.not_managed.node_list.interface_list / 9b3130b00fe0 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-170e0f5da067405e3870bd65926b08d31ccd32cfae2f393614d16db72ec5ea32): complete subsection reference.

<a id="canonical-475c48e5ea1ece848be03a9530595e44724793ddb4a0148d324ccdffd6f3cd0e"></a>

<a id="canonical-a64db78ff7d45585aba902c37e23feda52230fa08e9e727a06f8b71d9e1d979a"></a>

## description_spec property — kvm.not_managed.node_list.interface_list / 9b3130b00fe0 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-010.md#canonical-98c0a52521f3ccdef6737fbfeee8a599593a2a052c1dc84252fd861e3ce48646): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-eccff1a249c0fd91574a7f9e1bb340adf15b6234039269e3f983f41454939eba): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5): complete subsection reference.

<a id="canonical-abd083b21159fa427e0e983209b6998b4e0bb6ebc0cd4bcd5baa89921d708da5"></a>

<a id="canonical-b3b73bdbf41fa76cad920fe9c5f6b4d539e550d41669f943f3ca7b37c2646a6b"></a>

## is_management property — kvm.not_managed.node_list.interface_list / 9b3130b00fe0 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-8d7ffe88a199f2ea235e6676b212b191ace5c8d352bd20f3d8f3ac7f5b2817a1"></a>

<a id="canonical-d85da22c8eb0ca7707078037db7ea754d98380b1a9f7d49760988ab6ffd9d213"></a>

## is_primary property — kvm.not_managed.node_list.interface_list / 9b3130b00fe0 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-fd6311ead706f6fa3d096e3c085b53390a63574efba4addba98ba34f2583db45"></a>

<a id="canonical-7bb0455fd59221b8ecc8aceb9006977ee590aa4965d799f5aafbf1bbdfc1a852"></a>

## labels property — kvm.not_managed.node_list.interface_list / 9b3130b00fe0 / 7

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

- [monitor](resources--securemesh_site_v2--reference--group-011.md#canonical-86d48cf74cb6ec20a931bbfb5aa21353bea131bd30bbe85fb0b609a188ff8267): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-011.md#canonical-5089a3c8821ddebedbd32922dfc7b09737445d72d94eafa220b5d5fd99af019a): complete subsection reference.

<a id="canonical-434cc3461928e035d511930383fb89c555b4a680be64dce5f5c6872c30abc5e6"></a>

<a id="canonical-274648edfbe19d627e600b3fe4b3c023b1aee3f3827781a38622d8ccd822be00"></a>

## mtu property — kvm.not_managed.node_list.interface_list / 9b3130b00fe0 / 8

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

<a id="canonical-1f60b28bc5e9da9b01c547b61d1f5af8f11203728ed4da9dc381f725601a8362"></a>

<a id="canonical-2c095c3da37f1bc1b7aab3a6a9869f37bc9c903fa1c971a53f46057ebaae3960"></a>

## name property — kvm.not_managed.node_list.interface_list / 9b3130b00fe0 / 9

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

- [network_option](resources--securemesh_site_v2--reference--group-011.md#canonical-95195c28512bf633641d2a35dcb1fc4bac024249c31ef92cebebc0959b739017): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-011.md#canonical-db75d8f8359ee4b155a49ace03af36dbe9bed4e92bc51592bbd0065a32feaee1): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-35113afa738d3ff3faf3314fd5632e2c501dc921050553ef56458426a2fb7c91): complete subsection reference.

<a id="canonical-afb162b1bd6f2e5509fc870871a0497acfa28d1ae5e1ae010306eb2e39095029"></a>

<a id="canonical-c015edf66b9b6936a9712f32c0b859d47ae3c37748bc8223d96436de1e2b4d32"></a>

## priority property — kvm.not_managed.node_list.interface_list / 9b3130b00fe0 / 10

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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-011.md#canonical-fed971e3f1840028a982dba6a3186b22f92e720e2f8464631ef029a5c7557a2f): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-011.md#canonical-c739f48589c4e0507b72d4974b60e5a2462c1bfee830f5dae826c634cd74a755): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-011.md#canonical-1d904574b9d982a55620099dec99e0eb20a899e6103c21e86a73a891796659cf): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-bfe28f2aebf796d26c8e6026a08039c4b9c5bbe882afd60bce2e1bc7b72e47f9): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-6336643535d3cc49ea25eda2a16890f63b134e496b9b9d9c3aff069677cb615b): complete subsection reference.

<a id="canonical-a29c00b04bfca397a47b9595e39fddbd6a10fc718f4f952d3643bc1a56c13a79"></a>

## Next pages — kvm.not_managed.node_list.interface_list / 9b3130b00fe0 / 11

- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-170e0f5da067405e3870bd65926b08d31ccd32cfae2f393614d16db72ec5ea32)
- [kvm.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-010.md#canonical-98c0a52521f3ccdef6737fbfeee8a599593a2a052c1dc84252fd861e3ce48646)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- [kvm.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-eccff1a249c0fd91574a7f9e1bb340adf15b6234039269e3f983f41454939eba)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-011.md#canonical-86d48cf74cb6ec20a931bbfb5aa21353bea131bd30bbe85fb0b609a188ff8267)
- [kvm.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-011.md#canonical-5089a3c8821ddebedbd32922dfc7b09737445d72d94eafa220b5d5fd99af019a)
- [kvm.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-011.md#canonical-95195c28512bf633641d2a35dcb1fc4bac024249c31ef92cebebc0959b739017)
- [kvm.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-011.md#canonical-db75d8f8359ee4b155a49ace03af36dbe9bed4e92bc51592bbd0065a32feaee1)
- [kvm.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-35113afa738d3ff3faf3314fd5632e2c501dc921050553ef56458426a2fb7c91)
- [kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-011.md#canonical-fed971e3f1840028a982dba6a3186b22f92e720e2f8464631ef029a5c7557a2f)
- [kvm.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-011.md#canonical-c739f48589c4e0507b72d4974b60e5a2462c1bfee830f5dae826c634cd74a755)
- [kvm.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-011.md#canonical-1d904574b9d982a55620099dec99e0eb20a899e6103c21e86a73a891796659cf)
- [kvm.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-011.md#canonical-bfe28f2aebf796d26c8e6026a08039c4b9c5bbe882afd60bce2e1bc7b72e47f9)
- [kvm.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-6336643535d3cc49ea25eda2a16890f63b134e496b9b9d9c3aff069677cb615b)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-170e0f5da067405e3870bd65926b08d31ccd32cfae2f393614d16db72ec5ea32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7dc361f6214bc6695f31e3924f9512b867819296ee58105f1a09691f90212de4"></a>

## kvm.not_managed.node_list.interface_list.bond_interface — kvm.not_managed.node_list.interface_list.bond_interface / 18c88814c16a / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.bond_interface

<a id="canonical-e1da462691f1801b3d387c98fba6b79d8589840b474da3af75f47405b4c5fa3f"></a>

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

<a id="canonical-d33e051797562c0ce945d715ba2a80905d5d51f7ddadee2219b6da2f81eb57c6"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.bond_interface / 18c88814c16a / 3

- [active_backup](resources--securemesh_site_v2--reference--group-010.md#canonical-3e446f2e4d49d8bedf671902ed2e393d20219dce069e4e0c43a1d7f855f29050): complete subsection reference.

<a id="canonical-eb6358e1c1ce0b593e0429eed2e7f6ba8602d88fd3419f61cd16420126eee361"></a>

<a id="canonical-866c62cb286ebc5a4722a65cc6178a4116ab6a970e805da38701cbd299396888"></a>

## devices property — kvm.not_managed.node_list.interface_list.bond_interface / 18c88814c16a / 4

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

- [lacp](resources--securemesh_site_v2--reference--group-010.md#canonical-8ce97c8caa1c4327abda80264a5cbf8bb7bd53893b67627e6fa470e88791d710): complete subsection reference.

<a id="canonical-850f68f266cfdd33a97423bded2ddfffb73086dc39deb771c4c00429de23c743"></a>

<a id="canonical-4febf6bb6e2d2cc41062ef8264b7d6a019d422abd95606f25fda6d9851548a99"></a>

## link_polling_interval property — kvm.not_managed.node_list.interface_list.bond_interface / 18c88814c16a / 5

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

<a id="canonical-79a276cbf2e44d519a3f1bcdf98a30303961c1c265c0dcdd912bc3ce45e3b64e"></a>

<a id="canonical-e3d33791d33dfe1d37435faa5d56065f8ff5632a08519992a422c79069d23fe5"></a>

## link_up_delay property — kvm.not_managed.node_list.interface_list.bond_interface / 18c88814c16a / 6

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

<a id="canonical-dadeba80dc9026afc39cf2a723796d41e57ce19d964c32399b0fc28bd582c842"></a>

<a id="canonical-c4d12b9aace9098628213b8befe1cc9ef65393d0c0af6eb5d32722d55adc3444"></a>

## name property — kvm.not_managed.node_list.interface_list.bond_interface / 18c88814c16a / 7

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

<a id="canonical-2b1683ddeab3e5c50e06f5409e1865ec80403c0d9390165f6ad28a4220c36119"></a>

## Next pages — kvm.not_managed.node_list.interface_list.bond_interface / 18c88814c16a / 8

- [kvm.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-010.md#canonical-3e446f2e4d49d8bedf671902ed2e393d20219dce069e4e0c43a1d7f855f29050)
- [kvm.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-010.md#canonical-8ce97c8caa1c4327abda80264a5cbf8bb7bd53893b67627e6fa470e88791d710)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3e446f2e4d49d8bedf671902ed2e393d20219dce069e4e0c43a1d7f855f29050"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2caf0b93696e3b97e0c74f069ccfcb4bdf3ed522bcb4736f03c4a2d7e729240"></a>

## kvm.not_managed.node_list.interface_list.bond_interface.active_backup — kvm.not_managed.node_list.interface_list.bond_interface.active_backup / 4c39ddddb635 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-170e0f5da067405e3870bd65926b08d31ccd32cfae2f393614d16db72ec5ea32)
- kvm.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-4cc52ad59d012786d6e05251498bcfabe8f9b5f297775cc957a7b3712da5f75c"></a>

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

<a id="canonical-9aa19fa38c6c0248138bca5c59a39bef15fbfd37656ee2a4bdaf52b44563370e"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.bond_interface.active_backup / 4c39ddddb635 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-00a9a42f9843ac14a36ac0a0b33a28e67083306a92cacd1012eb50169632f24e"></a>

## Next pages — kvm.not_managed.node_list.interface_list.bond_interface.active_backup / 4c39ddddb635 / 4

- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-170e0f5da067405e3870bd65926b08d31ccd32cfae2f393614d16db72ec5ea32)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-8ce97c8caa1c4327abda80264a5cbf8bb7bd53893b67627e6fa470e88791d710"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e089a5f96e0748dbbcaf556f83519a058e7e4f9556f5b3aa3ca36e37f5af50cc"></a>

## kvm.not_managed.node_list.interface_list.bond_interface.lacp — kvm.not_managed.node_list.interface_list.bond_interface.lacp / 6c6810079226 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-170e0f5da067405e3870bd65926b08d31ccd32cfae2f393614d16db72ec5ea32)
- kvm.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-5b3d5028ff99d3c486e7b3aa82a70288275b781f551ecb648773c7abb426f8c6"></a>

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

<a id="canonical-4491e86f331c53498d3795a93cb9e58ba3991b58e61a1be7ca50d9bd31bce3b1"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.bond_interface.lacp / 6c6810079226 / 3

<a id="canonical-9d171a6c5f89d8f153431945ee437349cbf6f1f2664d3092850a3fa667343640"></a>

<a id="canonical-53bdf9ed063db915032596e67ba55335338e38738b9e3499d964fd4eb449fe80"></a>

## rate property — kvm.not_managed.node_list.interface_list.bond_interface.lacp / 6c6810079226 / 4

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

<a id="canonical-c68ddd737dab5ce367ac647e3f161fc89a86453299c53019290fff68d8de74c8"></a>

## Next pages — kvm.not_managed.node_list.interface_list.bond_interface.lacp / 6c6810079226 / 5

- [kvm.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-170e0f5da067405e3870bd65926b08d31ccd32cfae2f393614d16db72ec5ea32)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-98c0a52521f3ccdef6737fbfeee8a599593a2a052c1dc84252fd861e3ce48646"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e5df15a8ce27fd7d351226d3c97d8100847dc870ff6d2d278b42188d3743ba1"></a>

## kvm.not_managed.node_list.interface_list.dhcp_client — kvm.not_managed.node_list.interface_list.dhcp_client / 24796d00b040 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-3bf178734f2d72c3bf37a01c655a93ca4af49e9153dd78d507d46acef30f3fdf"></a>

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

<a id="canonical-e228343c6e87a8194ae49ba4aa61429aeb0b1ef2a070c9c7d1d057213e44ee66"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.dhcp_client / 24796d00b040 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef1256feb17f1fa870bb37f41a7dd80eb4a7f82561fa0468f190ba359e6a2a11"></a>

## Next pages — kvm.not_managed.node_list.interface_list.dhcp_client / 24796d00b040 / 4

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9e432abe39480bd9481c47632a3703cbfcfbab706d2e1b2d2ca08b4024afbdd"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server — kvm.not_managed.node_list.interface_list.dhcp_server / 868cc5ac6965 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-903e743d31bc266cb1af5eedb02f09b3b6be61c373e088637d0d4c95cad6391d"></a>

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

<a id="canonical-9b162d6f0726ccc104b67f0d81db2ba6d68bf013a5ac6fb5c04896ef548ae056"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.dhcp_server / 868cc5ac6965 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-010.md#canonical-2f89220f66f2fae088b42378e1b0181fdcff06bd890b3a9e717fa0c21de6d146): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-010.md#canonical-2fda93933c82be48c64a961f23a152b26a6e958b4bc64239ac20470ae3524c2e): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-d8c3595a8150f90359686183b9d7f6329c33ca65d365849c37913115928a3f73): complete subsection reference.

<a id="canonical-392753732ecd81997de9214788b3417a92620bdbcf45a7897b1a3a9b51b433dc"></a>

<a id="canonical-3a2778aa7bf6f251d6d0a2b3fe9c73163b907a0f447eb31e0f392a13f6008829"></a>

## dhcp_option82_tag property — kvm.not_managed.node_list.interface_list.dhcp_server / 868cc5ac6965 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-eb89a655f7f68046d41d42d20d75210ea13df7d4e21a9033b08081c0042a9e38"></a>

<a id="canonical-4cab4618b134db035739fbc776c3108203733c676b2b85e3ea396b8ebad276a8"></a>

## fixed_ip_map property — kvm.not_managed.node_list.interface_list.dhcp_server / 868cc5ac6965 / 5

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-010.md#canonical-4284a046dadacbba69773ec15ceb4ae88aa52e18b1eb09b0fe7412bc39420cf3): complete subsection reference.

<a id="canonical-29f4014d0025e97ba2d204af0daa79e7785c231fb658e59596e0961ea6995392"></a>

## Next pages — kvm.not_managed.node_list.interface_list.dhcp_server / 868cc5ac6965 / 6

- [kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-010.md#canonical-2f89220f66f2fae088b42378e1b0181fdcff06bd890b3a9e717fa0c21de6d146)
- [kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-010.md#canonical-2fda93933c82be48c64a961f23a152b26a6e958b4bc64239ac20470ae3524c2e)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-d8c3595a8150f90359686183b9d7f6329c33ca65d365849c37913115928a3f73)
- [kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-010.md#canonical-4284a046dadacbba69773ec15ceb4ae88aa52e18b1eb09b0fe7412bc39420cf3)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-2f89220f66f2fae088b42378e1b0181fdcff06bd890b3a9e717fa0c21de6d146"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94af13d27085f05208fb3aeb2714fa285586b0ad0ac683cd9c312e6c4a8e1eb9"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 7b4393a584dc / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-b3cb61ea2ed4c62fbb6304d3bd260dbe8370bb4d656dd62c784b38484ed4b2d3"></a>

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

<a id="canonical-db4cda4f2f68dbd92df9e683031880e583e7277acb077923afffbb644127ea0e"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 7b4393a584dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03f2f75fe9d973a4ffabff71515a1620a56cc566776184185efc1ba1f6a5502d"></a>

## Next pages — kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_end / 7b4393a584dc / 4

- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-2fda93933c82be48c64a961f23a152b26a6e958b4bc64239ac20470ae3524c2e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8dd2602df142bbea504dafc44abdf5fe65696d22eea2b4fab6631d6a5c0908f3"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / f16de702d048 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-13872c71a29bd9c74ef1793e54b9d6fa0d6ded9b3e2f3d628a64bbd3cb43dab2"></a>

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

<a id="canonical-362e8b91cf7e2b0f0f23c18509bf1291249db383ea7ffb8a57ebd926ce4fd301"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / f16de702d048 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6954ff8676e5f8bd8f9cac6f0b00dcbf4fd34313002e1cbc162ff543f07c7bfb"></a>

## Next pages — kvm.not_managed.node_list.interface_list.dhcp_server.automatic_from_start / f16de702d048 / 4

- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-d8c3595a8150f90359686183b9d7f6329c33ca65d365849c37913115928a3f73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b6735aa05d32422661887637829485ab3fc145f315cf7977ecbe48bf11d70c5"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 98351be7d1b6 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-b94a4842ba3e9a951c9c4f38b22beadf40e68a91a8d4fedecf458d9eb23002dd"></a>

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

<a id="canonical-c383530b57682caa6116d2407fdef68c2886f4df5f7f29cb5fed1fe9877504b0"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 98351be7d1b6 / 3

<a id="canonical-ad3b1f552759763a2059b0ecc1929ae57ab1cb04a2ab26889c11bf22642bbdcb"></a>

<a id="canonical-248f030cc4f9356c6c7067c676943665cc912b05dd7004570de932f42893cd6c"></a>

## dgw_address property — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 98351be7d1b6 / 4

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

<a id="canonical-372a1ca45ddf3cb91aebf710caf84032feefe8ae7b1144bd4bf3f15169ea83d4"></a>

<a id="canonical-b317b11537642a8df9267f8bdbf63d8ac7fabe005470dcb10f57ceec9c2722a9"></a>

## dns_address property — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 98351be7d1b6 / 5

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

- [first_address](resources--securemesh_site_v2--reference--group-010.md#canonical-1c1325261d4ed5b4ba443787bc91381d55fad2441f17bccd3224ee452b950a13): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-010.md#canonical-02f70bbb2571bad8aa3a7aec6c4c25cbc5e518534337c00b13f47b80f85456db): complete subsection reference.

<a id="canonical-6cecb3155ab61ef19e01fd9763fe16963433c83ba2faa40fc956940d148ff719"></a>

<a id="canonical-5e5a69734a21cf28a5db8f302fa8d9ea01ea176cbe0bcca374995ca30eebbe1a"></a>

## network_prefix property — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 98351be7d1b6 / 6

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

<a id="canonical-9c451f580bac73d99ec95ab3a932deda03f37678827f0eab84a00353156180d9"></a>

<a id="canonical-a915efd725b7a92a30c2ebe7165f60532feb0dc16f44cfa50a8507fa60f875f5"></a>

## pool_settings property — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 98351be7d1b6 / 7

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

- [pools](resources--securemesh_site_v2--reference--group-010.md#canonical-93b5f58869fa18ee8d2cfd42d6e0b76e2d0b225440d6d654b30f5fa7e2d43fa6): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-010.md#canonical-3722436a8991484818dfa3041e433baad32338072143085ec063374d5f5f8fb1): complete subsection reference.

<a id="canonical-08707ee2b9a344656d83d0fdefcb2370f1fcbf8a2a05d791ff1ab8f38ac2a1dc"></a>

## Next pages — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks / 98351be7d1b6 / 8

- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-010.md#canonical-1c1325261d4ed5b4ba443787bc91381d55fad2441f17bccd3224ee452b950a13)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-010.md#canonical-02f70bbb2571bad8aa3a7aec6c4c25cbc5e518534337c00b13f47b80f85456db)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-010.md#canonical-93b5f58869fa18ee8d2cfd42d6e0b76e2d0b225440d6d654b30f5fa7e2d43fa6)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-010.md#canonical-3722436a8991484818dfa3041e433baad32338072143085ec063374d5f5f8fb1)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-1c1325261d4ed5b4ba443787bc91381d55fad2441f17bccd3224ee452b950a13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da54f722798df71660e896232cf429a234dfae5d3d48c9d234b126c5078809bd"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address / 631f469f2ec8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-d8c3595a8150f90359686183b9d7f6329c33ca65d365849c37913115928a3f73)
- kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-27e652217aba0fe6411700e5c96ff6e46c42bd6d6890b343d6e9eac0beadce01"></a>

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

<a id="canonical-9df9b3b20a002bc78aea3a7b326d73b9f7d241054981dd3a4d1ca30404033705"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address / 631f469f2ec8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-314de441762c53b6576519f281ca785e7f87bed3094ee93cc0247882bb40a1d2"></a>

## Next pages — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address / 631f469f2ec8 / 4

- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-d8c3595a8150f90359686183b9d7f6329c33ca65d365849c37913115928a3f73)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-02f70bbb2571bad8aa3a7aec6c4c25cbc5e518534337c00b13f47b80f85456db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47e05f0d382e5a0da62d1c0bae812e148d463f17a7d17c170a3f99e1914a2d1a"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address / 959ab9f1d0ab / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-d8c3595a8150f90359686183b9d7f6329c33ca65d365849c37913115928a3f73)
- kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-544f55927c915f81db39702e0a4bdae379e8e6418f1595cb82dd186339b50a65"></a>

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

<a id="canonical-114f76ecad3f60bca254a44570648b02412472661b4fba18c96ec404ee180b86"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address / 959ab9f1d0ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-954da11015c223d23a530b820514aba0647037028ecac0c9a6486202c70d841f"></a>

## Next pages — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address / 959ab9f1d0ab / 4

- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-d8c3595a8150f90359686183b9d7f6329c33ca65d365849c37913115928a3f73)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-93b5f58869fa18ee8d2cfd42d6e0b76e2d0b225440d6d654b30f5fa7e2d43fa6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0380e16aba64faba7a0b95a1407144728c7081fce9238bc5428d91eae878ac8e"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9788a89ee61b / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-d8c3595a8150f90359686183b9d7f6329c33ca65d365849c37913115928a3f73)
- kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-894049d8de8b87b2ca12233793491058c28ba2870e81a4f38690eb2d8ccb540a"></a>

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

<a id="canonical-5896bb8c256fe905104563af000ccf8360921eacce7fbd5e7b975ace088b654b"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9788a89ee61b / 3

<a id="canonical-b7895aecdb7141596b785b470e703c3f4a9fb0138a19f1dbba7e57ce7813d324"></a>

<a id="canonical-71b4e03067018cb2724bb7105e1f5e7e1d440a55a02adad5029875b7785eb682"></a>

## end_ip property — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9788a89ee61b / 4

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

<a id="canonical-549756c08228ff882174923ad330535e2266021334ff746eeb8081bff4d09c0f"></a>

<a id="canonical-358b94501d2f025ecfc31d589408b9d635af6c0905012c37a7e1d57f793edbd1"></a>

## exclude property — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9788a89ee61b / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-eb5be152d5ddabaf8f549ea6c87fee0ca4216859539874b2f7f9b79c075b8116"></a>

<a id="canonical-8b4e86e1ff1dc52c07a8ba80b836bd4e4dccd379c191c39a0b66b27c5c2c060e"></a>

## start_ip property — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9788a89ee61b / 6

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

<a id="canonical-6ae0dfd547b683013782b97e5d4f33a85dda466fe5df54424f42dd3b3b8a0fb4"></a>

## Next pages — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools / 9788a89ee61b / 7

- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-d8c3595a8150f90359686183b9d7f6329c33ca65d365849c37913115928a3f73)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3722436a8991484818dfa3041e433baad32338072143085ec063374d5f5f8fb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da1cc28c07930345b1272030ba9d9ecc51060e4ef866f8fe8447cb77034d84ee"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / 4b0cce785644 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-d8c3595a8150f90359686183b9d7f6329c33ca65d365849c37913115928a3f73)
- kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-9aa709c9b65e84c2ce1b62ecd728700b9a2d65e94575c1948c2b399cb6986bb5"></a>

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

<a id="canonical-dcbed70f0c456887d05cee67d968df30a09bcf7e7a81cfae72962a8c3a0568ca"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / 4b0cce785644 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bb9a1f1dfce5ad96c2adbcdaa378ada2f990086910f20e098d6128ae25c18b0c"></a>

## Next pages — kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw / 4b0cce785644 / 4

- [kvm.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-010.md#canonical-d8c3595a8150f90359686183b9d7f6329c33ca65d365849c37913115928a3f73)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-4284a046dadacbba69773ec15ceb4ae88aa52e18b1eb09b0fe7412bc39420cf3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54e2b5a6427ae2648b3e58c72240a361a1b7a1a7f9c1467c1dfa752b567ac9ce"></a>

## kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / db349b5b6447 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-680b64f13ef2a54d8942568877ee76ccba3961a875d28b61af57954b274ef8eb"></a>

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

<a id="canonical-94412a0d36ac74d1b42a0e431a6b0eed8f80f48192529c6bb83022252239cec5"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / db349b5b6447 / 3

<a id="canonical-7d7cc616c60c12fdaaef70abb28ae9691ae8eec99a48fdfb6a578ccf76c77543"></a>

<a id="canonical-cbcdd1a3dae39ea951a7eeca526ce0185bf5a627694fb674e7b5418b20718044"></a>

## interface_ip_map property — kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / db349b5b6447 / 4

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

<a id="canonical-0800c6eb054e43c5036f6a215c7ab659fd373e341585db9f8f056af497e440fb"></a>

## Next pages — kvm.not_managed.node_list.interface_list.dhcp_server.interface_ip_map / db349b5b6447 / 5

- [kvm.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-010.md#canonical-b8c7b0359696349d80545a6feec588f1eb8805d838241bed218be52d693634bf)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-eccff1a249c0fd91574a7f9e1bb340adf15b6234039269e3f983f41454939eba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3587b78224d4e787b97a4d5222ef3abd3f4542c489e82d183f73208cf663445f"></a>

## kvm.not_managed.node_list.interface_list.ethernet_interface — kvm.not_managed.node_list.interface_list.ethernet_interface / 50bd2d618623 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-6816e7ba2ce89bba459e1fa2705e36d49175dd3aba19f4af4e67291354f9574f"></a>

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

<a id="canonical-d26b7e1829e865fa293c646922db2e0e3f0c32eadfb12776c102bc99b56ccaa9"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ethernet_interface / 50bd2d618623 / 3

<a id="canonical-aae3a4cbbb894ebeff921392baf88435d1c6aa51050bcca03c7eb2bf0661cbea"></a>

<a id="canonical-27ac86b5774850a417d97d95a279475faada9b5c3fe126ef6d5029215ad9c706"></a>

## device property — kvm.not_managed.node_list.interface_list.ethernet_interface / 50bd2d618623 / 4

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

<a id="canonical-6c68f4a76f64122347c30ffae39890946d325c7501ce83cdccf135b652854f20"></a>

<a id="canonical-41fdacce3eae68cb6d3da5d4ce95b8ca0f297ccc802906fc748593d05a257814"></a>

## mac property — kvm.not_managed.node_list.interface_list.ethernet_interface / 50bd2d618623 / 5

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

<a id="canonical-962f4e9293ce2753ce244e05df7f830dbe8de5dceb581e2726ec92f3a5febf53"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ethernet_interface / 50bd2d618623 / 6

- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5572fea7a0c479e086902292282245a776c78ae6d96cdecda53ad3f28aa345de"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config — kvm.not_managed.node_list.interface_list.ipv6_auto_config / b226ce0ac1a8 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-761e41773db43160ebe876a82ce15c65b5dd6d948b5e622407aa8edf4987ec55"></a>

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

<a id="canonical-0542b07fcd0f574fd1c884ab7c856b50585ad6737c0d952ade1b9322e98ab1c5"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config / b226ce0ac1a8 / 3

- [host](resources--securemesh_site_v2--reference--group-010.md#canonical-2993777ccec68725aa4e7752564db6c0265feca3ade91efd68d92be098a8cb1e): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1): complete subsection reference.

<a id="canonical-5d3864132f8eb5e94b50bdff38181249c16cb14fe66c7453480c110225473a58"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config / b226ce0ac1a8 / 4

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-010.md#canonical-2993777ccec68725aa4e7752564db6c0265feca3ade91efd68d92be098a8cb1e)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-2993777ccec68725aa4e7752564db6c0265feca3ade91efd68d92be098a8cb1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb796820c18c6edea0b87d155b253f41d7dcc8e244205841590fc9a3613db51f"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.host — kvm.not_managed.node_list.interface_list.ipv6_auto_config.host / 939080d0d4c9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-d7e859c229418c5ab1a1f1e1239741f0c4e1e5cd9c85a0b98bdcc27995368aad"></a>

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

<a id="canonical-9688f3dc3ecb2e10ce508e2d82f5bbaeccc100f108f080d1ced060daa624c390"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.host / 939080d0d4c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8a253375fc8d99138f21fec304db0b690f672d9a76069aebd1e63f31dac89ba6"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.host / 939080d0d4c9 / 4

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f164eef4fb88aefe781b2f716bee84fc635d2a06417fdcbc95330b05e4ae0d8d"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router / 791422c3c9aa / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-49a29961455afcfda67037e968158e3803c3f60056388c0584a751895768a385"></a>

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

<a id="canonical-51ecf4eca76e1b24447d6c950acb3acd43de26929c89f653d930ee57b996c033"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router / 791422c3c9aa / 3

- [dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-3a1ca3f3867a894bb2fe3df9421f82946150c096fed903bb6c578a0e71683dc8): complete subsection reference.

<a id="canonical-17b5bbfb7c60164ce3d88c405e539b6c9d690c7f87a93d376ac45f9cde3b5cb6"></a>

<a id="canonical-43f824af25aa6eb35f1310b9633fcefce9c3f16800418149a37429fbcca5c3e7"></a>

## network_prefix property — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router / 791422c3c9aa / 4

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

- [stateful](resources--securemesh_site_v2--reference--group-011.md#canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7): complete subsection reference.

<a id="canonical-d8601645ae577b1213788bdf6bf18b74b730c10f09ac274de9791a5de5d781ba"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router / 791422c3c9aa / 5

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-3a1ca3f3867a894bb2fe3df9421f82946150c096fed903bb6c578a0e71683dc8)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-011.md#canonical-1e9c668841f8c962cdd4423e9fb3c42feea0992b55e6de5a2717d4de66ee22b7)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-3a1ca3f3867a894bb2fe3df9421f82946150c096fed903bb6c578a0e71683dc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf8ae3a8e6610682686c09ccbb049bd824d7d6a5fa6084899f05bf98c7edad21"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 3870670d52c7 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-e1beba3d9fbcd9f83a1c626163c6c9667ebc6777eb5c3699c3aef23214873e17"></a>

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

<a id="canonical-9fb583b9bccf2156b6c7be70efc35b22902832a3331cbaba7da349e09f65ea49"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 3870670d52c7 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-010.md#canonical-6943c25d8a07b9d8537eb038295862cf0ff70a177efb25995715487eb5efdf82): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-5915cec528cdf8614eb1a1532ec9a6b289784ec0229a0f0dc88cc20c3831517e): complete subsection reference.

<a id="canonical-5585d5ad1d4f3eafba744c238f78fa1185ee9461bc11127401125a55f9f84649"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config / 3870670d52c7 / 4

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-010.md#canonical-6943c25d8a07b9d8537eb038295862cf0ff70a177efb25995715487eb5efdf82)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-010.md#canonical-5915cec528cdf8614eb1a1532ec9a6b289784ec0229a0f0dc88cc20c3831517e)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-6943c25d8a07b9d8537eb038295862cf0ff70a177efb25995715487eb5efdf82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d588ddca0a3b1f22d15427b220ded542b2bab5ad99bb2bb16f0a649c57591b4"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / 8ef79e8195d9 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-3a1ca3f3867a894bb2fe3df9421f82946150c096fed903bb6c578a0e71683dc8)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-03e00e8a207267649c809f2339817573344f6ccb99ff3ace2e196500f5200d7e"></a>

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

<a id="canonical-7cdaf7576bd731d3bed58ece44d06b125e557ba284b3c2d2243b2cb4fad63817"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / 8ef79e8195d9 / 3

<a id="canonical-f7ac7cb3f9ba733eaf6a08ba179e59717710b602ebdd058271a49ac9951531db"></a>

<a id="canonical-cfc4648849b4e321dacdd09532c3b7275812c45fcfe946cc077c02f0b69198ea"></a>

## dns_list property — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / 8ef79e8195d9 / 4

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

<a id="canonical-e2364b70f2f4f1135fd212a9c09b7b2f43d64f8389177dbd2bdb473ec9618f9b"></a>

## Next pages — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.conf / 8ef79e8195d9 / 5

- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-3a1ca3f3867a894bb2fe3df9421f82946150c096fed903bb6c578a0e71683dc8)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)

<a id="canonical-5915cec528cdf8614eb1a1532ec9a6b289784ec0229a0f0dc88cc20c3831517e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9708187c5325f7f560d9a96ee54d8580a3c0030a5fb1343e1ba2eab431edee90"></a>

## kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 92eb471ee230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-c082aad6996e4faefd74b5f511e52de16c46fc4f52b36a294840082e105babd8)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-b1115461a891b06e7c32d6d3d95ee5ba1d8e0e09c969a6db924997e4e8cfcef2)
- [kvm](resources--securemesh_site_v2--reference--group-010.md#canonical-bca55c8cc2cddcb084b0d55851f2cc7e7c8b6cda0422c7ade50e0d2956b252db)
- [kvm.not_managed](resources--securemesh_site_v2--reference--group-010.md#canonical-367c7006b3b81ddf105ae4b2b887782a293fcaade3264a91e001df48802f246c)
- [kvm.not_managed.node_list](resources--securemesh_site_v2--reference--group-010.md#canonical-9ab37e02d50f7e896f2e34d835174f6372f3da094046875e411160a71f3ec596)
- [kvm.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-010.md#canonical-f7e945b3b18e7734990fb75886344b8d86e093d7b3d455ff6b7394465b87e223)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-341e2d31368215faef8def4004bf01ad8d9333a13026a0728113f4731c44a6b5)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-010.md#canonical-74ec76d1388d50b382b7ec9ef5fef6841a0752a62506d5dd1eb6404570a65bf1)
- [kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-010.md#canonical-3a1ca3f3867a894bb2fe3df9421f82946150c096fed903bb6c578a0e71683dc8)
- kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-bf9b1246ab69a45e3ed3cdf86fa810fa625c2907155b03dc445ecce6ddbb7dde"></a>

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

<a id="canonical-7f32129c39f587d62a3d5e53b80694ea9bab5b0352ae1b85de17f4c4b193bed2"></a>

## Direct properties — kvm.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.loca / 92eb471ee230 / 3

<a id="canonical-36434d6bedf73f806ada94c94040947536a370dde3abcb8e40ba75e404f84ff9"></a>
