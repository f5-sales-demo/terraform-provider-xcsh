---
page_title: "xcsh_virtual_network reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_virtual_network reference."
---

# xcsh_virtual_network reference

<a id="canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0062404f24ad6904f85d11ab7fdf9b1c8610903013f73c922e32462369cba777"></a>

## Property reference — Property reference / 7eafa07b1cd7 / 2

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
- Property reference

<a id="canonical-0e73a33191fa00d81b74fd518cd1bfc784bf16576a6defbf709d42e6a1cb9fee"></a>

## Direct properties — Property reference / 7eafa07b1cd7 / 3

<a id="canonical-614213ddad9e7e53eabf5e05fdf7970d7d850f8c21d653fda2b1f8d975b508c8"></a>

<a id="canonical-d451ad9b28164043b88625cf87e8f1742765a07c3386c8284a01b658f2433b7c"></a>

## annotations property — Property reference / 7eafa07b1cd7 / 4

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

<a id="canonical-73199cc8cf4ccf16471ab722d1e674e9a20563d9d390bb145f367e4ef14af17c"></a>

<a id="canonical-073b6c9649cea6a9f5aec89f2392f9625ed3d42f12850922c83337afd625404d"></a>

## description property — Property reference / 7eafa07b1cd7 / 5

Type: `"string"`. Computed.

Description of the VirtualNetwork.

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

- [global_network](data-sources--virtual_network--reference--group-001.md#canonical-7e1cae114935b4288f4ae0cce1a77ecbb2d9f5c217e4c55cf70bca80cf248b45): complete subsection reference.

<a id="canonical-0f3baf4f89d8ed439176010d5382a9c4385be6e5836be71d49f92ffa1aaa146e"></a>

<a id="canonical-ad13ac9ffaac84c35fce2852eab7ba17e79c00eb4302cf18a3aaf57db4f6c774"></a>

## id property — Property reference / 7eafa07b1cd7 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-a964f4c53d86e67fd9cf93d8367192ce584aa42f70b736af1c57000a84d0985c"></a>

<a id="canonical-566b8b0d3792076bee34c9ff0b1ab24ab9b84d6c496dfa8355f7f1f8a269de48"></a>

## labels property — Property reference / 7eafa07b1cd7 / 7

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

<a id="canonical-33719296a6522a527f7ee10a8a10ade33f8e4e4fad0ec064224fe64427657bbd"></a>

<a id="canonical-eed72ddd959f9472d592332d40a6aa319dcd97184d1a362d6e979cb10cdb47ea"></a>

## name property — Property reference / 7eafa07b1cd7 / 8

Type: `"string"`. Required.

Name of the VirtualNetwork.

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

<a id="canonical-fa1c678c03cb682491b98f806f882378e33aa6c971434abdf35aee5f18f14b81"></a>

<a id="canonical-809bf827dd67a63fb0defbb80ef5b48e03b19ca702e503e839b1a5c870e7b413"></a>

## namespace property — Property reference / 7eafa07b1cd7 / 9

Type: `"string"`. Optional, Computed.

Namespace where the VirtualNetwork exists.

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

- [site_local_inside_network](data-sources--virtual_network--reference--group-001.md#canonical-88149641890ea53373a3f3c102c1744974876da17d82cec14f53ef32cfd78e0a): complete subsection reference.

- [site_local_network](data-sources--virtual_network--reference--group-001.md#canonical-0975629cb360fcfd561cd3106aa5105ce9726c3b480a137b437355e425fbe3e6): complete subsection reference.

- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-c919e9e97cea9abff9d1810d69b2d9d8508c571abc5dd6fc86ebb1556583ac2d): complete subsection reference.

<a id="canonical-191409c6b9633e5549b469df30e4fae328026a07d0ad61215e6ad4b503ecde80"></a>

## All schema paths — Property reference / 7eafa07b1cd7 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--virtual_network--reference--group-001.md#canonical-614213ddad9e7e53eabf5e05fdf7970d7d850f8c21d653fda2b1f8d975b508c8) |
| `description` | [description](data-sources--virtual_network--reference--group-001.md#canonical-73199cc8cf4ccf16471ab722d1e674e9a20563d9d390bb145f367e4ef14af17c) |
| `global_network` | [global_network](data-sources--virtual_network--reference--group-001.md#canonical-3218230cf7bdd08c9e649ddd24330eb5851395e94abae0364d215725ad00acfc) |
| `id` | [id](data-sources--virtual_network--reference--group-001.md#canonical-0f3baf4f89d8ed439176010d5382a9c4385be6e5836be71d49f92ffa1aaa146e) |
| `labels` | [labels](data-sources--virtual_network--reference--group-001.md#canonical-a964f4c53d86e67fd9cf93d8367192ce584aa42f70b736af1c57000a84d0985c) |
| `name` | [name](data-sources--virtual_network--reference--group-001.md#canonical-33719296a6522a527f7ee10a8a10ade33f8e4e4fad0ec064224fe64427657bbd) |
| `namespace` | [namespace](data-sources--virtual_network--reference--group-001.md#canonical-fa1c678c03cb682491b98f806f882378e33aa6c971434abdf35aee5f18f14b81) |
| `site_local_inside_network` | [site_local_inside_network](data-sources--virtual_network--reference--group-001.md#canonical-a678ea0e12d91f1be1711dfdc642e0d7a1dd95275829a1997263096d78f1850d) |
| `site_local_network` | [site_local_network](data-sources--virtual_network--reference--group-001.md#canonical-a030c6355623dda851d96f3eb8f95bb1491fc73cc8b14119cc827f430b718346) |
| `static_routes` | [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-269bce72797abf0a1bb80fc4f94f95468a400d77d35f59febc8cc5074b793d24) |
| `static_routes.attrs` | [static_routes.attrs](data-sources--virtual_network--reference--group-001.md#canonical-2ad935799339a2c087ceb732c82a5c4f18600c231981ef09080b04f6bbb7b538) |
| `static_routes.default_gateway` | [static_routes.default_gateway](data-sources--virtual_network--reference--group-001.md#canonical-819e638bd0e67d01d41f81da1f05ec568b5c868baae516cf261a3695b7356f74) |
| `static_routes.ip_address` | [static_routes.ip_address](data-sources--virtual_network--reference--group-001.md#canonical-1dab0be898a2bf1ea6e0fe3dca2c4eba740bf8f918cc5296a30eabe85501e072) |
| `static_routes.ip_prefixes` | [static_routes.ip_prefixes](data-sources--virtual_network--reference--group-001.md#canonical-1857c6fe456c9c8ce98bf172593cc06d6a828801f817cea4290db8d5794c0338) |
| `static_routes.node_interface` | [static_routes.node_interface](data-sources--virtual_network--reference--group-001.md#canonical-15d4b253436499b9d2496c6aa437fc36f485bb574e17734254a606d421ea1664) |
| `static_routes.node_interface.list` | [static_routes.node_interface.list](data-sources--virtual_network--reference--group-001.md#canonical-b3f67cde420960bd8e1ae34adf0e29e2ee39d156994961ade8bc828e5445c3ff) |
| `static_routes.node_interface.list.interface` | [static_routes.node_interface.list.interface](data-sources--virtual_network--reference--group-001.md#canonical-cc08dc5fd0bfba2037d5c1569e86c874b64b8363770c7f82910b687a28e741ff) |
| `static_routes.node_interface.list.interface.kind` | [static_routes.node_interface.list.interface.kind](data-sources--virtual_network--reference--group-001.md#canonical-1fc78b86928dab5b95ca72a5fc4adc1c5464bdd808d989f19ac979fe3483305a) |
| `static_routes.node_interface.list.interface.name` | [static_routes.node_interface.list.interface.name](data-sources--virtual_network--reference--group-001.md#canonical-822e12576559c193b8acae0349c401919498641fe7b5389f9f28205547c74c89) |
| `static_routes.node_interface.list.interface.namespace` | [static_routes.node_interface.list.interface.namespace](data-sources--virtual_network--reference--group-001.md#canonical-0d11566d3f740b3d1b2409cf10ace8239a12387a67060436551eda8fce32d02f) |
| `static_routes.node_interface.list.interface.tenant` | [static_routes.node_interface.list.interface.tenant](data-sources--virtual_network--reference--group-001.md#canonical-ae2455be8c4bce520ce0f942d44a3652f852adae0f9bc65f109b43e0b166d417) |
| `static_routes.node_interface.list.interface.uid` | [static_routes.node_interface.list.interface.uid](data-sources--virtual_network--reference--group-001.md#canonical-9839d7e626a7b5d230f820d89261b0a6c4ba00c1166839d9d018ac01fd249d1f) |
| `static_routes.node_interface.list.node` | [static_routes.node_interface.list.node](data-sources--virtual_network--reference--group-001.md#canonical-ab67ebbc9d9b9812ff90c96157a39e59ad1a34b3a36a383f7ee223a2c991341b) |

<a id="canonical-e64107f4655ad5b2d1479f711dc2541f1032a1ad908e5e932e7aaf318ae6e9f2"></a>

## Next pages — Property reference / 7eafa07b1cd7 / 11

- [global_network](data-sources--virtual_network--reference--group-001.md#canonical-7e1cae114935b4288f4ae0cce1a77ecbb2d9f5c217e4c55cf70bca80cf248b45)
- [site_local_inside_network](data-sources--virtual_network--reference--group-001.md#canonical-88149641890ea53373a3f3c102c1744974876da17d82cec14f53ef32cfd78e0a)
- [site_local_network](data-sources--virtual_network--reference--group-001.md#canonical-0975629cb360fcfd561cd3106aa5105ce9726c3b480a137b437355e425fbe3e6)
- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-c919e9e97cea9abff9d1810d69b2d9d8508c571abc5dd6fc86ebb1556583ac2d)
- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)

<a id="canonical-7e1cae114935b4288f4ae0cce1a77ecbb2d9f5c217e4c55cf70bca80cf248b45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-faeb38fa4276b10d1c1fc747f2e66b28f489a7b0d2075d639476597fd5edd1b6"></a>

## global_network — global_network / 1f62a5459a12 / 2

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- global_network

<a id="canonical-3218230cf7bdd08c9e649ddd24330eb5851395e94abae0364d215725ad00acfc"></a>

Type: `["object", {}]`. Computed.

\[OneOf: global\_network, site\_local\_inside\_network, site\_local\_network\] Select the global
virtual-network scope for connectivity across participating sites.

Upstream description:

Select the global virtual-network scope for connectivity across participating sites.

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

OneOf alternatives in this subsection:

- [global_network](data-sources--virtual_network--reference--group-001.md#canonical-3218230cf7bdd08c9e649ddd24330eb5851395e94abae0364d215725ad00acfc)
- [site_local_inside_network](data-sources--virtual_network--reference--group-001.md#canonical-a678ea0e12d91f1be1711dfdc642e0d7a1dd95275829a1997263096d78f1850d)
- [site_local_network](data-sources--virtual_network--reference--group-001.md#canonical-a030c6355623dda851d96f3eb8f95bb1491fc73cc8b14119cc827f430b718346)

Select alternatives according to the provider validators above.

<a id="canonical-75c402aadfe82ff6fc291b5d2e62feb783eb0427c8c4e49d0bdda811fbe355c4"></a>

## Direct properties — global_network / 1f62a5459a12 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d6bbd07f7a31c5b271af80106255e0d09bb37488296b5d1e66203179a80ee1c2"></a>

## Next pages — global_network / 1f62a5459a12 / 4

- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)

<a id="canonical-88149641890ea53373a3f3c102c1744974876da17d82cec14f53ef32cfd78e0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd1da45e7013343892ea7120f0c320510a3c7bca35c2a4082124c2b46dc45670"></a>

## site_local_inside_network — site_local_inside_network / 7e5bec0b2910 / 2

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- site_local_inside_network

<a id="canonical-a678ea0e12d91f1be1711dfdc642e0d7a1dd95275829a1997263096d78f1850d"></a>

Type: `["object", {}]`. Computed.

Select the site-local inside network for site-internal connectivity.

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

<a id="canonical-61a1a5e9beb355c4cf939c2f5aece20125b3f528b8aa77b84a15aa4ac1a92371"></a>

## Direct properties — site_local_inside_network / 7e5bec0b2910 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f228d3a0623744f1d3b00f854d315c367266b3147eb8d75cbf45999ad14b1bda"></a>

## Next pages — site_local_inside_network / 7e5bec0b2910 / 4

- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)

<a id="canonical-0975629cb360fcfd561cd3106aa5105ce9726c3b480a137b437355e425fbe3e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bfc133b7a86d5fba880c7841fb4baaef65ea16951f3767300da3ee158a911050"></a>

## site_local_network — site_local_network / 6ec34b21a6ec / 2

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- site_local_network

<a id="canonical-a030c6355623dda851d96f3eb8f95bb1491fc73cc8b14119cc827f430b718346"></a>

Type: `["object", {}]`. Computed.

Select a site-local virtual network when connectivity must remain within one site.

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

<a id="canonical-300d4f3b30ada248af6bf144e1d88a4208d33febb846c13d97a5e81065f6d7d4"></a>

## Direct properties — site_local_network / 6ec34b21a6ec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9117f918d6c8d9418d7b835309a2ee027d8f1548fc89355e55092fee47fdae2a"></a>

## Next pages — site_local_network / 6ec34b21a6ec / 4

- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)

<a id="canonical-c919e9e97cea9abff9d1810d69b2d9d8508c571abc5dd6fc86ebb1556583ac2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4fe3bc081a85fb2c1cf0c404e5447c3065d734076ffc30030250c6038e7131c"></a>

## static_routes — static_routes / b69705092e03 / 2

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- static_routes

<a id="canonical-269bce72797abf0a1bb80fc4f94f95468a400d77d35f59febc8cc5074b793d24"></a>

Type: `"list"`. Computed.

List of static routes on the virtual network.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 165,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 165,
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
    "ves.io.schema.rules.repeated.max_items": "165",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "165",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-05511ce5984a279fce59d92ff6d59a0d0430e8e54b6cd9e81369a41f299fcb2b"></a>

## Direct properties — static_routes / b69705092e03 / 3

<a id="canonical-2ad935799339a2c087ceb732c82a5c4f18600c231981ef09080b04f6bbb7b538"></a>

<a id="canonical-787a57e91856002d2ae7c1cd3e316a6d8fc223e316aa9de7c1479a08f3665a5b"></a>

## attrs property — static_routes / b69705092e03 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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

- [default_gateway](data-sources--virtual_network--reference--group-001.md#canonical-7830f8fa1afc2582e55636415723fdcbc4d894a8ec4d54c298a423b2fb7f01fe): complete subsection reference.

<a id="canonical-1dab0be898a2bf1ea6e0fe3dca2c4eba740bf8f918cc5296a30eabe85501e072"></a>

<a id="canonical-05289b822c3eeaf25dc254887323b1e36fd91e3126fd402a34b6fb407fdfd3e3"></a>

## ip_address property — static_routes / b69705092e03 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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

<a id="canonical-1857c6fe456c9c8ce98bf172593cc06d6a828801f817cea4290db8d5794c0338"></a>

<a id="canonical-838e1be940faf76fdbfceb94bff84426d6cbb204c84ab14d17d34d5b34214863"></a>

## ip_prefixes property — static_routes / b69705092e03 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--virtual_network--reference--group-001.md#canonical-635b352a9f75fa313734a5cf0969d7c60afaa8c1e0e005d2de8daa2ea180bd18): complete subsection reference.

<a id="canonical-c52072502cbccdc437fe6530b2e5d7d14a8bffcde85916d84b98be031d1d697a"></a>

## Next pages — static_routes / b69705092e03 / 7

- [static_routes.default_gateway](data-sources--virtual_network--reference--group-001.md#canonical-7830f8fa1afc2582e55636415723fdcbc4d894a8ec4d54c298a423b2fb7f01fe)
- [static_routes.node_interface](data-sources--virtual_network--reference--group-001.md#canonical-635b352a9f75fa313734a5cf0969d7c60afaa8c1e0e005d2de8daa2ea180bd18)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)

<a id="canonical-7830f8fa1afc2582e55636415723fdcbc4d894a8ec4d54c298a423b2fb7f01fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1378f91094aeba1edf37fe9e87b48db1558c6e6b4916a49b39abd6970baba78"></a>

## static_routes.default_gateway — static_routes.default_gateway / b88ea3303737 / 2

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-c919e9e97cea9abff9d1810d69b2d9d8508c571abc5dd6fc86ebb1556583ac2d)
- static_routes.default_gateway

<a id="canonical-819e638bd0e67d01d41f81da1f05ec568b5c868baae516cf261a3695b7356f74"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-ccbaecc1cf535ca9e498e47d52aa629386ebf77f49894d92e2d390e48285bd4f"></a>

## Direct properties — static_routes.default_gateway / b88ea3303737 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-15bac170262ad2bc0d487dd56bd8cda71d22cf5804cc6e55b152e41bdb5019b3"></a>

## Next pages — static_routes.default_gateway / b88ea3303737 / 4

- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-c919e9e97cea9abff9d1810d69b2d9d8508c571abc5dd6fc86ebb1556583ac2d)
- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)

<a id="canonical-635b352a9f75fa313734a5cf0969d7c60afaa8c1e0e005d2de8daa2ea180bd18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e28eb29ad99b88e84e63cdc8d973adbd9c0fe39d0564ac07fae9ecc83ce8276a"></a>

## static_routes.node_interface — static_routes.node_interface / fb07da9ba49c / 2

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-c919e9e97cea9abff9d1810d69b2d9d8508c571abc5dd6fc86ebb1556583ac2d)
- static_routes.node_interface

<a id="canonical-15d4b253436499b9d2496c6aa437fc36f485bb574e17734254a606d421ea1664"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-987fab51d6b8b2e395e480d3d26e60d0eb2087d6f7c8fcd5357d2542eefaf41b"></a>

## Direct properties — static_routes.node_interface / fb07da9ba49c / 3

- [list](data-sources--virtual_network--reference--group-001.md#canonical-6561b4cd31d108d2beb93cd81c4284750430f967854ac70a77b9190ab9abf6ca): complete subsection reference.

<a id="canonical-d12fec05a36ce094e365f185fe5add1c2c9cba694e50ed7b0f07aeb41b542958"></a>

## Next pages — static_routes.node_interface / fb07da9ba49c / 4

- [static_routes.node_interface.list](data-sources--virtual_network--reference--group-001.md#canonical-6561b4cd31d108d2beb93cd81c4284750430f967854ac70a77b9190ab9abf6ca)
- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-c919e9e97cea9abff9d1810d69b2d9d8508c571abc5dd6fc86ebb1556583ac2d)
- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)

<a id="canonical-6561b4cd31d108d2beb93cd81c4284750430f967854ac70a77b9190ab9abf6ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a62aba7e30a790e692fd9983f3ff0636e3be4ebca56f1435fe0e992c73be8dc7"></a>

## static_routes.node_interface.list — static_routes.node_interface.list / 5683d1357965 / 2

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-c919e9e97cea9abff9d1810d69b2d9d8508c571abc5dd6fc86ebb1556583ac2d)
- [static_routes.node_interface](data-sources--virtual_network--reference--group-001.md#canonical-635b352a9f75fa313734a5cf0969d7c60afaa8c1e0e005d2de8daa2ea180bd18)
- static_routes.node_interface.list

<a id="canonical-b3f67cde420960bd8e1ae34adf0e29e2ee39d156994961ade8bc828e5445c3ff"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-423e02dce4ed5e349206032517fe79b7b11983eb38b80ab0d74463546b7fe716"></a>

## Direct properties — static_routes.node_interface.list / 5683d1357965 / 3

- [interface](data-sources--virtual_network--reference--group-001.md#canonical-752d0f9465a0f8a8af5e2b801f41e05640728a08883381aec3930bc48ec56245): complete subsection reference.

<a id="canonical-ab67ebbc9d9b9812ff90c96157a39e59ad1a34b3a36a383f7ee223a2c991341b"></a>

<a id="canonical-ba291b518dc13996c57f5394b5a2d6d6d58430168c43d9bec6038df1806213e7"></a>

## node property — static_routes.node_interface.list / 5683d1357965 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-03dec203a2e876a92e7477eb117bc1e4a871b2b9e0fb2f1b74a1471b64842404"></a>

## Next pages — static_routes.node_interface.list / 5683d1357965 / 5

- [static_routes.node_interface.list.interface](data-sources--virtual_network--reference--group-001.md#canonical-752d0f9465a0f8a8af5e2b801f41e05640728a08883381aec3930bc48ec56245)
- [static_routes.node_interface](data-sources--virtual_network--reference--group-001.md#canonical-635b352a9f75fa313734a5cf0969d7c60afaa8c1e0e005d2de8daa2ea180bd18)
- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)

<a id="canonical-752d0f9465a0f8a8af5e2b801f41e05640728a08883381aec3930bc48ec56245"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58ac43880d8c5025ce8f7b41e762a0da3fd6a2a5051ddcd2f10173e7fa1e2f3b"></a>

## static_routes.node_interface.list.interface — static_routes.node_interface.list.interface / 434af4ce5e04 / 2

Breadcrumbs:

- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
- [Property reference](data-sources--virtual_network--reference--group-001.md#canonical-9c449ec5c2f1fe76d44c1c939fe391b4c4e3c20e4e3572d7fc09ce3420ccce57)
- [static_routes](data-sources--virtual_network--reference--group-001.md#canonical-c919e9e97cea9abff9d1810d69b2d9d8508c571abc5dd6fc86ebb1556583ac2d)
- [static_routes.node_interface](data-sources--virtual_network--reference--group-001.md#canonical-635b352a9f75fa313734a5cf0969d7c60afaa8c1e0e005d2de8daa2ea180bd18)
- [static_routes.node_interface.list](data-sources--virtual_network--reference--group-001.md#canonical-6561b4cd31d108d2beb93cd81c4284750430f967854ac70a77b9190ab9abf6ca)
- static_routes.node_interface.list.interface

<a id="canonical-cc08dc5fd0bfba2037d5c1569e86c874b64b8363770c7f82910b687a28e741ff"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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

<a id="canonical-6a951e182a9329d0d5bd220f3bb5a4133dc882df308dacc46b306f3e140d0d21"></a>

## Direct properties — static_routes.node_interface.list.interface / 434af4ce5e04 / 3

<a id="canonical-1fc78b86928dab5b95ca72a5fc4adc1c5464bdd808d989f19ac979fe3483305a"></a>

<a id="canonical-456796921a05be32233ee78291fcb41b2918fc800ea2e5070ec19714e407baaa"></a>

## kind property — static_routes.node_interface.list.interface / 434af4ce5e04 / 4

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

<a id="canonical-822e12576559c193b8acae0349c401919498641fe7b5389f9f28205547c74c89"></a>

<a id="canonical-8b444e9b0815f8eff1b1ed57e7dd2d3be4ec203037f93e7c58c54166c9acde9d"></a>

## name property — static_routes.node_interface.list.interface / 434af4ce5e04 / 5

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

<a id="canonical-0d11566d3f740b3d1b2409cf10ace8239a12387a67060436551eda8fce32d02f"></a>

<a id="canonical-529e62e3bdf0eeb18e089775cdc17ef9751788e48c754ae176ce00b4b51aaa66"></a>

## namespace property — static_routes.node_interface.list.interface / 434af4ce5e04 / 6

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

<a id="canonical-ae2455be8c4bce520ce0f942d44a3652f852adae0f9bc65f109b43e0b166d417"></a>

<a id="canonical-450ae0d45a0a9da489784d64d2d2ac04e1b2ab515b63eaf00c6608242c2c77c3"></a>

## tenant property — static_routes.node_interface.list.interface / 434af4ce5e04 / 7

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

<a id="canonical-9839d7e626a7b5d230f820d89261b0a6c4ba00c1166839d9d018ac01fd249d1f"></a>

<a id="canonical-62de10362d2afe923ba361343ede9797a198dbe57c505254ee14fce3de344e30"></a>

## uid property — static_routes.node_interface.list.interface / 434af4ce5e04 / 8

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

<a id="canonical-b7a4f961a9188fdbed74e1ae6ccabdd012bbd66e61f97f5d2df35e803a6bf397"></a>

## Next pages — static_routes.node_interface.list.interface / 434af4ce5e04 / 9

- [static_routes.node_interface.list](data-sources--virtual_network--reference--group-001.md#canonical-6561b4cd31d108d2beb93cd81c4284750430f967854ac70a77b9190ab9abf6ca)
- [xcsh_virtual_network](../data-sources/virtual_network.md#canonical-962f424f866ca39a0c5bda80d1dc95f430dbcfd7d44fc7a1bcd04ae7c7db361a)
