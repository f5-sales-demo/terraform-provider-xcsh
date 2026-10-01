---
page_title: "xcsh_address_allocator reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator reference."
---

# xcsh_address_allocator reference

<a id="canonical-06d6ef2a4ebffe9fba2c3fad1b3c86d8113e071f5d1178e3a4362816a6103b09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f173eb4f0f33f77fdfd606e92cfcd6022227ac5c75de0f717af55fe951002e9"></a>

## Property reference — Property reference / ce6d0f86c283 / 2

Breadcrumbs:

- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-b37f5189fa71c353983be346c78099ee10d1935bec31c657e73fc56ee3cc9d4b)
- Property reference

<a id="canonical-1c6e2513f58cb738ea40398c2fb224ac956639bbb38d3199b82d96bb31720908"></a>

## Direct properties — Property reference / ce6d0f86c283 / 3

- [address_allocation_scheme](data-sources--address_allocator--reference--group-001.md#canonical-95e50f3bc2e669e33cdc651c7959c4af78684a0a0072835ebc93773e6b5903b9): complete subsection reference.

<a id="canonical-05edcf3d8facc077148ba413c8e02adbc3d3e7372f1bc2209071e15334e63eff"></a>

<a id="canonical-94d7073e87c1c9ffa777f498ae7c721259f69c740d36b6c295fcfb7ceb489e3c"></a>

## address_pool property — Property reference / ce6d0f86c283 / 4

Type: `["list", "string"]`. Computed.

Address pool from which the allocator carves out subnets or addresses to its clients.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-92f962838aecee6021d0b906956007f30d0ffec92b01b62a6f471accd23aaa6f"></a>

<a id="canonical-32c2363012576fe3912a6bfc6a87bd93fcaf5d6f3fff4d2a4e5917e7a37d22ca"></a>

## annotations property — Property reference / ce6d0f86c283 / 5

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-edcf5ea4ff8d93d8c51e3091e7e4052827f0c28d022d01ec6571ac9c370f3e8b"></a>

<a id="canonical-99c3efe6e57673eb6b65c850677853ae5193e9fcabef5c3c29c3f54af1b38a89"></a>

## description property — Property reference / ce6d0f86c283 / 6

Type: `"string"`. Computed.

Description of the AddressAllocator.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-c6573bc47de2edc6b3760ec023475deed8a48839999a3b1cdd054b11939bd0d8"></a>

<a id="canonical-fe4b69a058a202948e2fbd3de1f43621fae1f5db09636f5101b640b6ccecca00"></a>

## id property — Property reference / ce6d0f86c283 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d8cb98f2ee2b5400662e6092460a0afe789365dad4648ddbc4f94b48fdfebdb6"></a>

<a id="canonical-c3bc52195849a62520af332b47cf827c724dbcc5549c89960034a47e81e3665c"></a>

## labels property — Property reference / ce6d0f86c283 / 8

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="canonical-6e97d96ad033b4d9a118afa0de23fb29c460c1bf99a72da18793510e1ce9f648"></a>

<a id="canonical-9ea074ef813edf2b4a9e19a219a059324957940c57dc829e9beba768a31cd5c7"></a>

## mode property — Property reference / ce6d0f86c283 / 9

Type: `"string"`. Computed.

\[Enum: LOCAL|GLOBAL\_PER\_SITE\_NODE\] Mode of the address allocator Address allocator is for VERs
within the local cluster or site Allocation is per site and then per node. Possible values are
\`LOCAL\`, \`GLOBAL\_PER\_SITE\_NODE\`. Defaults to \`LOCAL\`.

Upstream description:

Mode of the address allocator

Address allocator is for VERs within the local cluster or site Allocation is per site and then per
node.

Receipt-pinned upstream constraints:

```json
{
  "default": "LOCAL",
  "enum": [
    "LOCAL",
    "GLOBAL_PER_SITE_NODE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-55bf60cd82ef8ccd79f795adc2c66ba94103951a44d215a0c3bfa8f15d950e13"></a>

<a id="canonical-b886c7366c9ec86ae4b63e398b9463c5b6a79a6bc2b8e123d5b50221d4c91bfc"></a>

## name property — Property reference / ce6d0f86c283 / 10

Type: `"string"`. Required.

Name of the AddressAllocator.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-ef14e414f39d3b110b5ca196af1c32c8d7faf8efa074643bcab621cc26168a8a"></a>

<a id="canonical-92f1cf914d6e9b2c1f2838d5674822dfd29902e433a017b90e512d8beb578f37"></a>

## namespace property — Property reference / ce6d0f86c283 / 11

Type: `"string"`. Required.

Namespace where the AddressAllocator exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

<a id="canonical-436793f57cb64d244d5f20d90287d47562f629e83926bd15577bc88e593e9528"></a>

## All schema paths — Property reference / ce6d0f86c283 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address_allocation_scheme` | [address_allocation_scheme](data-sources--address_allocator--reference--group-001.md#canonical-dfabadf05af8bd98122d2896241dc738bbe4e686d1551f4f6e8ece7a6281886a) |
| `address_allocation_scheme.allocation_unit` | [address_allocation_scheme.allocation_unit](data-sources--address_allocator--reference--group-001.md#canonical-d816ab797958bfe79c1620ef6d511ec724a52827371ff2110dea88db0f121b6e) |
| `address_allocation_scheme.local_interface_address_offset` | [address_allocation_scheme.local_interface_address_offset](data-sources--address_allocator--reference--group-001.md#canonical-183b8df57817c2eda30acf17c6d74206e3d800c1fbc8470309258944dd6297eb) |
| `address_allocation_scheme.local_interface_address_type` | [address_allocation_scheme.local_interface_address_type](data-sources--address_allocator--reference--group-001.md#canonical-2541aff0792e93bfce778d93f29a613403d15bcac2cdeb4897dbe9156d8322fb) |
| `address_pool` | [address_pool](data-sources--address_allocator--reference--group-001.md#canonical-05edcf3d8facc077148ba413c8e02adbc3d3e7372f1bc2209071e15334e63eff) |
| `annotations` | [annotations](data-sources--address_allocator--reference--group-001.md#canonical-92f962838aecee6021d0b906956007f30d0ffec92b01b62a6f471accd23aaa6f) |
| `description` | [description](data-sources--address_allocator--reference--group-001.md#canonical-edcf5ea4ff8d93d8c51e3091e7e4052827f0c28d022d01ec6571ac9c370f3e8b) |
| `id` | [id](data-sources--address_allocator--reference--group-001.md#canonical-c6573bc47de2edc6b3760ec023475deed8a48839999a3b1cdd054b11939bd0d8) |
| `labels` | [labels](data-sources--address_allocator--reference--group-001.md#canonical-d8cb98f2ee2b5400662e6092460a0afe789365dad4648ddbc4f94b48fdfebdb6) |
| `mode` | [mode](data-sources--address_allocator--reference--group-001.md#canonical-6e97d96ad033b4d9a118afa0de23fb29c460c1bf99a72da18793510e1ce9f648) |
| `name` | [name](data-sources--address_allocator--reference--group-001.md#canonical-55bf60cd82ef8ccd79f795adc2c66ba94103951a44d215a0c3bfa8f15d950e13) |
| `namespace` | [namespace](data-sources--address_allocator--reference--group-001.md#canonical-ef14e414f39d3b110b5ca196af1c32c8d7faf8efa074643bcab621cc26168a8a) |

<a id="canonical-20e733acfbaa6e2ac9be0dae04b83c658363c45a6232ba03174913e84b76991c"></a>

## Next pages — Property reference / ce6d0f86c283 / 13

- [address_allocation_scheme](data-sources--address_allocator--reference--group-001.md#canonical-95e50f3bc2e669e33cdc651c7959c4af78684a0a0072835ebc93773e6b5903b9)
- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-b37f5189fa71c353983be346c78099ee10d1935bec31c657e73fc56ee3cc9d4b)

<a id="canonical-95e50f3bc2e669e33cdc651c7959c4af78684a0a0072835ebc93773e6b5903b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93ca886281fb088991afcb9796df884593340c0effbaeb87ad46a351a61b3b69"></a>

## address_allocation_scheme — address_allocation_scheme / 91bd8a7d147d / 2

Breadcrumbs:

- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-b37f5189fa71c353983be346c78099ee10d1935bec31c657e73fc56ee3cc9d4b)
- [Property reference](data-sources--address_allocator--reference--group-001.md#canonical-06d6ef2a4ebffe9fba2c3fad1b3c86d8113e071f5d1178e3a4362816a6103b09)
- address_allocation_scheme

<a id="canonical-dfabadf05af8bd98122d2896241dc738bbe4e686d1551f4f6e8ece7a6281886a"></a>

Type: `"single"`. Computed.

Decides the scheme to be used to allocate addresses from the configured address pool.

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

<a id="canonical-b985fad335b2282f062f9dd0277703c45d921c128b0c7b88def61d3e151efa11"></a>

## Direct properties — address_allocation_scheme / 91bd8a7d147d / 3

<a id="canonical-d816ab797958bfe79c1620ef6d511ec724a52827371ff2110dea88db0f121b6e"></a>

<a id="canonical-998d94e707879ece0bf1a3c1b8e5c70c8fde4241f1df16ec893721caeb344a22"></a>

## allocation_unit property — address_allocation_scheme / 91bd8a7d147d / 4

Type: `"number"`. Computed.

Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30,
subnets of /30 will be allocated from the given address pool.

Upstream description:

Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30,
subnets of /30 will be allocated from the given address pool.

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-183b8df57817c2eda30acf17c6d74206e3d800c1fbc8470309258944dd6297eb"></a>

<a id="canonical-e2c0d68465ea1f6e4ac944287a45db800546054e87fd636e379e3dacdb1bffb0"></a>

## local_interface_address_offset property — address_allocation_scheme / 91bd8a7d147d / 5

Type: `"number"`. Computed.

Used to derive address for the local interface from the allocated subnet. If Local Interface Address
Type is set to 'Offset from beginning of Subnet', this offset value is added to the allocated subnet
and used as the local interface address. For example, if the allocated subnet is 192.0.2.0/24..

Upstream description:

This is used to derive address for the local interface from the allocated subnet.

If Local Interface Address Type is set to "Offset from beginning of Subnet", this offset value is
added to the allocated subnet and used as the local interface address. For example, if the allocated
subnet is 192.0.2.0/24 and offset is set to 2 with Local Interface Address Type set to "Offset from
beginning of Subnet", local interface address of 192.0.2.204 is used.

If Local Interface Address Type is set to "Offset from end of Subnet", this offset value is
subtracted from the end of the allocated subnet and used as the local interface address. For
example, if the allocated subnet is 192.0.2.0/24 and offset is set to 1 with Local Interface Address
Type set to "Offset from end of Subnet", local interface address of 192.0.2.204 is used.

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
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-2541aff0792e93bfce778d93f29a613403d15bcac2cdeb4897dbe9156d8322fb"></a>

<a id="canonical-ca2f94537d7a044a17925f458e397746d6fd792c2473e53fe113ae02c1d61a7d"></a>

## local_interface_address_type property — address_allocation_scheme / 91bd8a7d147d / 6

Type: `"string"`. Computed.

\[Enum:
LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN|LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_END|LOCAL\_INTERFACE\_ADDRESS\_FROM\_PREFIX\]
Dictates how local interface address is derived from the allocated subnet Use Nth address of the
allocated subnet as the local interface address, N being the Local Interface Address Offset. For
example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset is set to 2 and
Local.. Possible values are \`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN\`,
\`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_END\`,
\`LOCAL\_INTERFACE\_ADDRESS\_FROM\_PREFIX\`. Defaults to
\`LOCAL\_INTERFACE\_ADDRESS\_OFFSET\_FROM\_SUBNET\_BEGIN\`.

Upstream description:

Dictates how local interface address is derived from the allocated subnet

Use Nth address of the allocated subnet as the local interface address, N being the Local Interface
Address Offset. For example, if the allocated subnet is 192.0.2.0/24, Local Interface Address Offset
is set to 2 and Local Interface Address Type is set to "Offset from beginning of Subnet", local
address of 192.0.2.204 is used.

Use Nth last address of the allocated subnet as the local interface address, N being the Local
Interface Address Offset. For example, if the allocated subnet is 192.0.2.0/24, Local Interface
Address Offset is set to 1 and Local Interface Address Type is set to "Offset from end of Subnet",
local address of 192.0.2.204 is used.

This case is used for external\_connector.

Receipt-pinned upstream constraints:

```json
{
  "default": "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN",
  "enum": [
    "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN",
    "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_END",
    "LOCAL_INTERFACE_ADDRESS_FROM_PREFIX"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3306f42a2386c748215699c5a1280b9bb837e0832471481026a07432460a5377"></a>

## Next pages — address_allocation_scheme / 91bd8a7d147d / 7

- [Property reference](data-sources--address_allocator--reference--group-001.md#canonical-06d6ef2a4ebffe9fba2c3fad1b3c86d8113e071f5d1178e3a4362816a6103b09)
- [xcsh_address_allocator](../data-sources/address_allocator.md#canonical-b37f5189fa71c353983be346c78099ee10d1935bec31c657e73fc56ee3cc9d4b)
