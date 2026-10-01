---
page_title: "xcsh_address_allocator reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator reference."
---

# xcsh_address_allocator reference

<a id="canonical-6c53b0269d6a197e4c6beae947a7d28d842ea2701871f0c541ede9d544757f27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-610939a0d933bd1a194ef8b8ada5717297c58030c62d9fbe4d2f87aa6ef1372b"></a>

## Property reference — Property reference / c33d27f562f0 / 2

Breadcrumbs:

- [xcsh_address_allocator](../resources/address_allocator.md#canonical-a90eda9d0678adc17d442e30a4391e1edf799c2a7419e020376896c3e91a3aba)
- Property reference

<a id="canonical-97c1de0feb8c42d2963d32787ea255a954770b52686c6a4c640f79f67da178d0"></a>

## Direct properties — Property reference / c33d27f562f0 / 3

- [address_allocation_scheme](resources--address_allocator--reference--group-001.md#canonical-ecc1bfcbc9d6483751a2b2034da5394781e9987125b3cf138cf4f16fba07704e): complete subsection reference.

<a id="canonical-fe11ee4d0e551b474d8f17a4031c9a46bcc86a206441ca49bf4789c6ce3ff35f"></a>

<a id="canonical-7a79a2aae0bc5fcbf0ced5f831d381a7984f130e13ee56bf32167f3b34f7d386"></a>

## address_pool property — Property reference / c33d27f562f0 / 4

Type: `["list", "string"]`. Required.

Address pool from which the allocator carves out subnets or addresses to its clients.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

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

<a id="canonical-afe69c592251d96f42ac89de0941e136ef22a03caac6c43bfaca0cd069347945"></a>

<a id="canonical-23d527ea93962878aecbe5ed19b04859406439d499444e265872c7b4be42550d"></a>

## annotations property — Property reference / c33d27f562f0 / 5

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-f0ff1de65093dbfdae416db0769d1f9804d6d2656ae8b3701c397400c8347a84"></a>

<a id="canonical-42c71af13312feb6652fbcd3c1a016a0b61181f00fe6b099a65a71a1bc34230f"></a>

## description property — Property reference / c33d27f562f0 / 6

Type: `"string"`. Optional.

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

<a id="canonical-651e0c658cd45510e57a380bd611a5dd4724a22783360bc92f57bbd38d4cc35b"></a>

<a id="canonical-b2ade4609a0d05306af9fa2e0b79854d32bddb8e598d02c35d0985a0fa79bf8b"></a>

## disable property — Property reference / c33d27f562f0 / 7

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="canonical-b145cac39e171403c13c15f580738593ae2a28937f3d3dfc7c224dd6a9b4e2a6"></a>

<a id="canonical-e40bb827a939bbb8eeca50d95089502ccc4d9ff4015847640586454a3c86079c"></a>

## id property — Property reference / c33d27f562f0 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-67880220593fa2e2fce6f7775374e1ed5934ade5d9e662e54cd06f37c503f104"></a>

<a id="canonical-5640afa28303ef6cd23376c16dcbb286517915c38526a9ebc187489d45873a7f"></a>

## labels property — Property reference / c33d27f562f0 / 9

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-750e218c0b82de063f83ba5ed0d2ccca977432dc03c37faf2c9689f348b8c1dc"></a>

<a id="canonical-b466a8bdae09848ad3d5c5127a68ab535e447b05b51f7cad7948db23320a77a3"></a>

## mode property — Property reference / c33d27f562f0 / 10

Type: `"string"`. Optional, Computed.

\[Enum: LOCAL|GLOBAL\_PER\_SITE\_NODE\] Mode of the address allocator Address allocator is for VERs
within the local cluster or site Allocation is per site and then per node. Possible values are
\`LOCAL\`, \`GLOBAL\_PER\_SITE\_NODE\`. Defaults to \`LOCAL\`.

Upstream description:

Mode of the address allocator

Address allocator is for VERs within the local cluster or site Allocation is per site and then per
node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("LOCAL",
    "GLOBAL_PER_SITE_NODE"),
}
```

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

<a id="canonical-ed91648a5c96cfdbcecf913bf8c1e1dfba9f48e205931147a9415fa294e49884"></a>

<a id="canonical-dd17b6a8cbc0f7aea43a4679991ea8de8aa4741804a778be9b7737f35f1bf027"></a>

## name property — Property reference / c33d27f562f0 / 11

Type: `"string"`. Required.

Name of the Address Allocator. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

<a id="canonical-ba872e03ac174d6759c415652139a259b9074bb850fb266e4cdf55a631c4d84a"></a>

<a id="canonical-244fe3e8e6fd72b3cf88577f106d800714a153451e6f391c783f298794dc0796"></a>

## namespace property — Property reference / c33d27f562f0 / 12

Type: `"string"`. Required.

Namespace where the Address Allocator is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [timeouts](resources--address_allocator--reference--group-001.md#canonical-b835e00ca60dc363aca9ca7b4346de500a7fe3c97945b0b64189702003a62d41): complete subsection reference.

<a id="canonical-9dec74e7145e67b1956916e6d93ec2d51cea13eb41ad2a06dec0ddbdd90fa865"></a>

## All schema paths — Property reference / c33d27f562f0 / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address_allocation_scheme` | [address_allocation_scheme](resources--address_allocator--reference--group-001.md#canonical-1bf1de207162f333c5f5f721523482087756344d8ce02b012db3b11484ab2ecc) |
| `address_allocation_scheme.allocation_unit` | [address_allocation_scheme.allocation_unit](resources--address_allocator--reference--group-001.md#canonical-711881a7bfbbad877048c1597cfa2e4c980e80d9972872aff19dd5c08659391b) |
| `address_allocation_scheme.local_interface_address_offset` | [address_allocation_scheme.local_interface_address_offset](resources--address_allocator--reference--group-001.md#canonical-b2f9e1ee979c34330e17546ede2d1db99278e3f9852918cb6344a4c3c59ea809) |
| `address_allocation_scheme.local_interface_address_type` | [address_allocation_scheme.local_interface_address_type](resources--address_allocator--reference--group-001.md#canonical-3ad91f3a95f4522460b4ca78fefd587c91eb386ed52381ff8cc5e73190d39a5e) |
| `address_pool` | [address_pool](resources--address_allocator--reference--group-001.md#canonical-fe11ee4d0e551b474d8f17a4031c9a46bcc86a206441ca49bf4789c6ce3ff35f) |
| `annotations` | [annotations](resources--address_allocator--reference--group-001.md#canonical-afe69c592251d96f42ac89de0941e136ef22a03caac6c43bfaca0cd069347945) |
| `description` | [description](resources--address_allocator--reference--group-001.md#canonical-f0ff1de65093dbfdae416db0769d1f9804d6d2656ae8b3701c397400c8347a84) |
| `disable` | [disable](resources--address_allocator--reference--group-001.md#canonical-651e0c658cd45510e57a380bd611a5dd4724a22783360bc92f57bbd38d4cc35b) |
| `id` | [id](resources--address_allocator--reference--group-001.md#canonical-b145cac39e171403c13c15f580738593ae2a28937f3d3dfc7c224dd6a9b4e2a6) |
| `labels` | [labels](resources--address_allocator--reference--group-001.md#canonical-67880220593fa2e2fce6f7775374e1ed5934ade5d9e662e54cd06f37c503f104) |
| `mode` | [mode](resources--address_allocator--reference--group-001.md#canonical-750e218c0b82de063f83ba5ed0d2ccca977432dc03c37faf2c9689f348b8c1dc) |
| `name` | [name](resources--address_allocator--reference--group-001.md#canonical-ed91648a5c96cfdbcecf913bf8c1e1dfba9f48e205931147a9415fa294e49884) |
| `namespace` | [namespace](resources--address_allocator--reference--group-001.md#canonical-ba872e03ac174d6759c415652139a259b9074bb850fb266e4cdf55a631c4d84a) |
| `timeouts` | [timeouts](resources--address_allocator--reference--group-001.md#canonical-91ee77560923c438c3c494dc00bb6b25f09c93852722d56d634f515c79f42869) |
| `timeouts.create` | [timeouts.create](resources--address_allocator--reference--group-001.md#canonical-209f540eafa12c4bd029238298f53e5b62732259f6782b1a3e859e8fa5fbb372) |
| `timeouts.delete` | [timeouts.delete](resources--address_allocator--reference--group-001.md#canonical-8030a54b9324d2c7d9fb966178336fbf0dbe29511ab9cc1b279b3917cce732fb) |
| `timeouts.read` | [timeouts.read](resources--address_allocator--reference--group-001.md#canonical-0f4d154eb132376ad56b9ce4b6e6d2116140c265664a173e580fd90556ec385d) |
| `timeouts.update` | [timeouts.update](resources--address_allocator--reference--group-001.md#canonical-6c36011f8053b0b7e7ea7987dbd24f87b329ba559e8002ea5b20ba01482e1af8) |

<a id="canonical-5a9bbad2a0b7f598287b58a711181dbeb698bc9ced23da332408af4251ed3396"></a>

## Next pages — Property reference / c33d27f562f0 / 14

- [address_allocation_scheme](resources--address_allocator--reference--group-001.md#canonical-ecc1bfcbc9d6483751a2b2034da5394781e9987125b3cf138cf4f16fba07704e)
- [timeouts](resources--address_allocator--reference--group-001.md#canonical-b835e00ca60dc363aca9ca7b4346de500a7fe3c97945b0b64189702003a62d41)
- [xcsh_address_allocator](../resources/address_allocator.md#canonical-a90eda9d0678adc17d442e30a4391e1edf799c2a7419e020376896c3e91a3aba)

<a id="canonical-ecc1bfcbc9d6483751a2b2034da5394781e9987125b3cf138cf4f16fba07704e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95da453d3e8d3bc19953b67066ae5c11ac848b4b3f5f85df75949fd7d3e0f15f"></a>

## address_allocation_scheme — address_allocation_scheme / e2b7a0663914 / 2

Breadcrumbs:

- [xcsh_address_allocator](../resources/address_allocator.md#canonical-a90eda9d0678adc17d442e30a4391e1edf799c2a7419e020376896c3e91a3aba)
- [Property reference](resources--address_allocator--reference--group-001.md#canonical-6c53b0269d6a197e4c6beae947a7d28d842ea2701871f0c541ede9d544757f27)
- address_allocation_scheme

<a id="canonical-1bf1de207162f333c5f5f721523482087756344d8ce02b012db3b11484ab2ecc"></a>

Type: `"object"`. single nested block, Optional.

Decides the scheme to be used to allocate addresses from the configured address pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("allocation_unit")}
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
address_allocation_scheme {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea9baa1c301ed1d7403e13491e4e41941fafc67c2f6ef8ab3cb4ced6ec4d514f"></a>

## Direct properties — address_allocation_scheme / e2b7a0663914 / 3

<a id="canonical-711881a7bfbbad877048c1597cfa2e4c980e80d9972872aff19dd5c08659391b"></a>

<a id="canonical-a6939446953df6b14906845406ac0c4737362781fe2d3a66d95b50c7aaa8f042"></a>

## allocation_unit property — address_allocation_scheme / e2b7a0663914 / 4

Type: `"number"`. Optional.

Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30,
subnets of /30 will be allocated from the given address pool.

Upstream description:

Prefix length indicating the size of each allocated subnet. For example, if this is specified as 30,
subnets of /30 will be allocated from the given address pool.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

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

<a id="canonical-b2f9e1ee979c34330e17546ede2d1db99278e3f9852918cb6344a4c3c59ea809"></a>

<a id="canonical-d5aed3380bb9da63b58ea1befa4ad7022325521fb6898bfa6ff70d6080c5f8fe"></a>

## local_interface_address_offset property — address_allocation_scheme / e2b7a0663914 / 5

Type: `"number"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

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

<a id="canonical-3ad91f3a95f4522460b4ca78fefd587c91eb386ed52381ff8cc5e73190d39a5e"></a>

<a id="canonical-b502f177781016d263c166a6681d15abbec235eadce7aa24d47ee7e07d3e37be"></a>

## local_interface_address_type property — address_allocation_scheme / e2b7a0663914 / 6

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_BEGIN",
    "LOCAL_INTERFACE_ADDRESS_OFFSET_FROM_SUBNET_END",
    "LOCAL_INTERFACE_ADDRESS_FROM_PREFIX"),
}
```

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

<a id="canonical-803f3c7d824016eb30dbb82d0185f360e0cf1195cb58248337fe6eab6245b11e"></a>

## Next pages — address_allocation_scheme / e2b7a0663914 / 7

- [Property reference](resources--address_allocator--reference--group-001.md#canonical-6c53b0269d6a197e4c6beae947a7d28d842ea2701871f0c541ede9d544757f27)
- [xcsh_address_allocator](../resources/address_allocator.md#canonical-a90eda9d0678adc17d442e30a4391e1edf799c2a7419e020376896c3e91a3aba)

<a id="canonical-b835e00ca60dc363aca9ca7b4346de500a7fe3c97945b0b64189702003a62d41"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-132f3bebec531b4d12a76711b139907b86685c132cb83ecbb4386ebcfe4c2689"></a>

## timeouts — timeouts / 9126692447dd / 2

Breadcrumbs:

- [xcsh_address_allocator](../resources/address_allocator.md#canonical-a90eda9d0678adc17d442e30a4391e1edf799c2a7419e020376896c3e91a3aba)
- [Property reference](resources--address_allocator--reference--group-001.md#canonical-6c53b0269d6a197e4c6beae947a7d28d842ea2701871f0c541ede9d544757f27)
- timeouts

<a id="canonical-91ee77560923c438c3c494dc00bb6b25f09c93852722d56d634f515c79f42869"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-96a394d20dbcdaac7ae308597af944c065c2056c23f36c2a259f29007000dd9e"></a>

## Direct properties — timeouts / 9126692447dd / 3

<a id="canonical-209f540eafa12c4bd029238298f53e5b62732259f6782b1a3e859e8fa5fbb372"></a>

<a id="canonical-e5a3832308065c23b8d8822c60a8f1deef855397b55aaaabcd6da66e24531f39"></a>

## create property — timeouts / 9126692447dd / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-8030a54b9324d2c7d9fb966178336fbf0dbe29511ab9cc1b279b3917cce732fb"></a>

<a id="canonical-5428ec829c274a90698f7b5746cec3b1e6b0dcac2b5405a16f1b9c9a3c70e2a6"></a>

## delete property — timeouts / 9126692447dd / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0f4d154eb132376ad56b9ce4b6e6d2116140c265664a173e580fd90556ec385d"></a>

<a id="canonical-a797430f8460bcbc579180f559b2528a1a50c1db46b0a923e971ed472ef1322f"></a>

## read property — timeouts / 9126692447dd / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-6c36011f8053b0b7e7ea7987dbd24f87b329ba559e8002ea5b20ba01482e1af8"></a>

<a id="canonical-8e8f7358d2a5c6212aaa7f3dccdf28a48593b35cd9b673a71cee30f93100ef65"></a>

## update property — timeouts / 9126692447dd / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-36054863a34adcbc6a4e46903a6f9b7cc46e270a56ca3a3159247d694e0d4efa"></a>

## Next pages — timeouts / 9126692447dd / 8

- [Property reference](resources--address_allocator--reference--group-001.md#canonical-6c53b0269d6a197e4c6beae947a7d28d842ea2701871f0c541ede9d544757f27)
- [xcsh_address_allocator](../resources/address_allocator.md#canonical-a90eda9d0678adc17d442e30a4391e1edf799c2a7419e020376896c3e91a3aba)
