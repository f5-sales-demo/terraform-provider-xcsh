---
page_title: "xcsh_ike2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 reference."
---

# xcsh_ike2 reference

<a id="canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bb47f499e06e0ff896025c4af062640afbd9c1ef9057dd0c517b51188301fae"></a>

## Property reference — Property reference / 9e7e22996d03 / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)
- Property reference

<a id="canonical-7c518f834192eb2dd05d95e7dfe0a325e58eb2fcd008a14002f2b9d3812c9fa1"></a>

## Direct properties — Property reference / 9e7e22996d03 / 3

<a id="canonical-c0b03ea3de32344f209fe1b311899add3dc29bfb63b34a3dafdf0a142a16ade1"></a>

<a id="canonical-e38eb6108d0d02a2f4a943cbec726ca873a8b57fc257c571aed740c8b0132b53"></a>

## annotations property — Property reference / 9e7e22996d03 / 4

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

<a id="canonical-58f5dabd71fd2ee03e8392db98d5a40c4e91e434d33c2ca493858c3eae407c4a"></a>

<a id="canonical-9d61b7d958462895dfabd40e8af4508c483617c3f658a7ba82928bab7c200ba6"></a>

## description property — Property reference / 9e7e22996d03 / 5

Type: `"string"`. Computed.

Description of the Ike2.

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

- [dh_group_set](data-sources--ike2--reference--group-001.md#canonical-b9c8809634ea9006bd4ba29f5e2c93179ba38ccefaade83d943582fdf71a9623): complete subsection reference.

- [disable_pfs](data-sources--ike2--reference--group-001.md#canonical-76a6982ce0ac110bc46a83fb298a1b7609d6132742f76ecc8dadd90eb7f550cb): complete subsection reference.

<a id="canonical-b0a9e79c8eb05e491e75bd8e3bd81f28d42e65e131062f982279f67dbb5ef845"></a>

<a id="canonical-b3f91da1b4812fa4521ebe67ca30e6bada73e07251168aefc35c918645c8abc9"></a>

## id property — Property reference / 9e7e22996d03 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](data-sources--ike2--reference--group-001.md#canonical-18a384b7883dec2f61054ecf1d0284a067cab73db65192eadadc2b3a6e257ad6): complete subsection reference.

- [ike_keylifetime_minutes](data-sources--ike2--reference--group-001.md#canonical-8393b8f320c2cee90a4953ad2c752f585cf0195ad52afd3ca9b6c03acbf2ab93): complete subsection reference.

<a id="canonical-10408f41467a79dc74b1babf6e8f7d685ce6d9afd386334e7e6c276687bb88ab"></a>

<a id="canonical-c9a834ddb71fb0c6057d3baf51dc73f8badc2055e7d5eeb09c739507de5f9e01"></a>

## labels property — Property reference / 9e7e22996d03 / 7

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

<a id="canonical-1e62032208a495ad53a4922e8f2579f96d229454e5615b4e3ef74604f5adc07d"></a>

<a id="canonical-c0785ef7ac74a91ab0148cb6e53e1bc08ea49ae45b64ceb05a78b0609ff8965c"></a>

## name property — Property reference / 9e7e22996d03 / 8

Type: `"string"`. Required.

Name of the Ike2.

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

<a id="canonical-60bb10ebfbedeeb9f1c67597b55a0df7f8d55e98aab799fef9b55cbd1a04eb1f"></a>

<a id="canonical-54aef3aa87f96137e4786b823f0917f1addbcc891eb2b8aea4fbad1cef1e918d"></a>

## namespace property — Property reference / 9e7e22996d03 / 9

Type: `"string"`. Required.

Namespace where the Ike2 exists.

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

- [use_default_keylifetime](data-sources--ike2--reference--group-001.md#canonical-919d2f5e851c79682c6da77495f53c90e4429bf1486f6080be53966cd9d0613e): complete subsection reference.

<a id="canonical-7855288c8ed86131ee70f6b7760ccc0ec45b20cc778af629fb5e4f39e6044252"></a>

## All schema paths — Property reference / 9e7e22996d03 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--ike2--reference--group-001.md#canonical-c0b03ea3de32344f209fe1b311899add3dc29bfb63b34a3dafdf0a142a16ade1) |
| `description` | [description](data-sources--ike2--reference--group-001.md#canonical-58f5dabd71fd2ee03e8392db98d5a40c4e91e434d33c2ca493858c3eae407c4a) |
| `dh_group_set` | [dh_group_set](data-sources--ike2--reference--group-001.md#canonical-f0ab07a66831b9646869e43d315e309da9d11b462f0cb6535106d60ac716035d) |
| `dh_group_set.dh_groups` | [dh_group_set.dh_groups](data-sources--ike2--reference--group-001.md#canonical-246c077079831eddc02f2619a671de62380967194bde3613308489c312c1ee5d) |
| `disable_pfs` | [disable_pfs](data-sources--ike2--reference--group-001.md#canonical-d4550fb3412fbf6ff505074f17b4b6faebe98cfacc780c2e9a51dd98ae68b6b1) |
| `id` | [id](data-sources--ike2--reference--group-001.md#canonical-b0a9e79c8eb05e491e75bd8e3bd81f28d42e65e131062f982279f67dbb5ef845) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](data-sources--ike2--reference--group-001.md#canonical-c602fd7c4d5bf05e991d0d5a33f26f0009814cc7ac562ee9ca82966f2875c3d8) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](data-sources--ike2--reference--group-001.md#canonical-74e3e9536d6c47eb5f87fdb5536242364dace6ca41e5cdbdf42fafc8580ecbcc) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](data-sources--ike2--reference--group-001.md#canonical-749e2a7a5c0ced05915361a6b19eb6b832807bdd4fd6c346d41e21406c79fc4b) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](data-sources--ike2--reference--group-001.md#canonical-5cd5d0c1a289e733afab4e35ab21cc87fae85c44b2e36ba365183ac4958e9df2) |
| `labels` | [labels](data-sources--ike2--reference--group-001.md#canonical-10408f41467a79dc74b1babf6e8f7d685ce6d9afd386334e7e6c276687bb88ab) |
| `name` | [name](data-sources--ike2--reference--group-001.md#canonical-1e62032208a495ad53a4922e8f2579f96d229454e5615b4e3ef74604f5adc07d) |
| `namespace` | [namespace](data-sources--ike2--reference--group-001.md#canonical-60bb10ebfbedeeb9f1c67597b55a0df7f8d55e98aab799fef9b55cbd1a04eb1f) |
| `use_default_keylifetime` | [use_default_keylifetime](data-sources--ike2--reference--group-001.md#canonical-12930694ef3e83ad515ffc8e155e5837b96eef76d29b2fbda6be350a7d76d36d) |

<a id="canonical-57a5819dd038d2b4ab0bd5572414a0d8ac32a278b778bcdcf9385589942f76d4"></a>

## Next pages — Property reference / 9e7e22996d03 / 11

- [dh_group_set](data-sources--ike2--reference--group-001.md#canonical-b9c8809634ea9006bd4ba29f5e2c93179ba38ccefaade83d943582fdf71a9623)
- [disable_pfs](data-sources--ike2--reference--group-001.md#canonical-76a6982ce0ac110bc46a83fb298a1b7609d6132742f76ecc8dadd90eb7f550cb)
- [ike_keylifetime_hours](data-sources--ike2--reference--group-001.md#canonical-18a384b7883dec2f61054ecf1d0284a067cab73db65192eadadc2b3a6e257ad6)
- [ike_keylifetime_minutes](data-sources--ike2--reference--group-001.md#canonical-8393b8f320c2cee90a4953ad2c752f585cf0195ad52afd3ca9b6c03acbf2ab93)
- [use_default_keylifetime](data-sources--ike2--reference--group-001.md#canonical-919d2f5e851c79682c6da77495f53c90e4429bf1486f6080be53966cd9d0613e)
- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)

<a id="canonical-b9c8809634ea9006bd4ba29f5e2c93179ba38ccefaade83d943582fdf71a9623"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-072ff443d0474df13abbae321a2359f060ee9d66d49b44bf6b7870b2103b0a76"></a>

## dh_group_set — dh_group_set / b6eb1d083acc / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)
- [Property reference](data-sources--ike2--reference--group-001.md#canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b)
- dh_group_set

<a id="canonical-f0ab07a66831b9646869e43d315e309da9d11b462f0cb6535106d60ac716035d"></a>

Type: `"single"`. Computed.

\[OneOf: dh\_group\_set, disable\_pfs; Default: disable\_pfs\] Choose the acceptable Diffie
Hellman(DH) Group or Groups that you are willing to accept as part of this profile.

Upstream description:

Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile.

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

- [dh_group_set](data-sources--ike2--reference--group-001.md#canonical-f0ab07a66831b9646869e43d315e309da9d11b462f0cb6535106d60ac716035d)
- [disable_pfs](data-sources--ike2--reference--group-001.md#canonical-d4550fb3412fbf6ff505074f17b4b6faebe98cfacc780c2e9a51dd98ae68b6b1)

Select alternatives according to the provider validators above.

<a id="canonical-8187523fd39dc58f34501685a7325863579b098469048525b09dfaa92cccf2a6"></a>

## Direct properties — dh_group_set / b6eb1d083acc / 3

<a id="canonical-246c077079831eddc02f2619a671de62380967194bde3613308489c312c1ee5d"></a>

<a id="canonical-0ab4fb9dd3293107b4c33f2f3b8b155418b98174c37dad218d9354725813da72"></a>

## dh_groups property — dh_group_set / b6eb1d083acc / 4

Type: `["list", "string"]`. Computed.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Diffie Hellman Groups. Group or collection configuration. Possible values are
\`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`, \`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`,
\`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`, \`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`.
Defaults to \`DH\_GROUP\_DEFAULT\`.

Upstream description:

Group or collection configuration

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

<a id="canonical-3e4de6cd5ca41ef431473386db4552f3aed7fc4dd8dd931eb922dcd58d0cc623"></a>

## Next pages — dh_group_set / b6eb1d083acc / 5

- [Property reference](data-sources--ike2--reference--group-001.md#canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b)
- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)

<a id="canonical-76a6982ce0ac110bc46a83fb298a1b7609d6132742f76ecc8dadd90eb7f550cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-364d0799229648f78f1fec20bfba48887ddd8f2a99248204e8c64ca9e2ca05c4"></a>

## disable_pfs — disable_pfs / e7e14baada1c / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)
- [Property reference](data-sources--ike2--reference--group-001.md#canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b)
- disable_pfs

<a id="canonical-d4550fb3412fbf6ff505074f17b4b6faebe98cfacc780c2e9a51dd98ae68b6b1"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable pfs.

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

<a id="canonical-895be66c54a59f96c0ca2b247443af36a5620a7b54490392f421687edceab49b"></a>

## Direct properties — disable_pfs / e7e14baada1c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-196daacbaf7aea967d5e8226a8f71ba210bccf7681164f787b1453cce32730dc"></a>

## Next pages — disable_pfs / e7e14baada1c / 4

- [Property reference](data-sources--ike2--reference--group-001.md#canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b)
- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)

<a id="canonical-18a384b7883dec2f61054ecf1d0284a067cab73db65192eadadc2b3a6e257ad6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-955f280d05454be5aaef9dce3203f8fca29619197d096fa62f1ffe85961d104d"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / bf3e04eb2648 / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)
- [Property reference](data-sources--ike2--reference--group-001.md#canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b)
- ike_keylifetime_hours

<a id="canonical-c602fd7c4d5bf05e991d0d5a33f26f0009814cc7ac562ee9ca82966f2875c3d8"></a>

Type: `"single"`. Computed.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Upstream description:

Input Hours.

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

- [ike_keylifetime_hours](data-sources--ike2--reference--group-001.md#canonical-c602fd7c4d5bf05e991d0d5a33f26f0009814cc7ac562ee9ca82966f2875c3d8)
- [ike_keylifetime_minutes](data-sources--ike2--reference--group-001.md#canonical-749e2a7a5c0ced05915361a6b19eb6b832807bdd4fd6c346d41e21406c79fc4b)
- [use_default_keylifetime](data-sources--ike2--reference--group-001.md#canonical-12930694ef3e83ad515ffc8e155e5837b96eef76d29b2fbda6be350a7d76d36d)

Select alternatives according to the provider validators above.

<a id="canonical-4282c01be47e10809766484892321021d2c5d8e2374d3491a996985dd4e10ce7"></a>

## Direct properties — ike_keylifetime_hours / bf3e04eb2648 / 3

<a id="canonical-74e3e9536d6c47eb5f87fdb5536242364dace6ca41e5cdbdf42fafc8580ecbcc"></a>

<a id="canonical-e4c78bf61d6944c0067b55d3cbf69f851fed4f1d8c5bd1cc980ff34227a0af5a"></a>

## duration property — ike_keylifetime_hours / bf3e04eb2648 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
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
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-1389af5bce62d3899c4e108e00fc53680278da4babb8250baeae52d4d40cb106"></a>

## Next pages — ike_keylifetime_hours / bf3e04eb2648 / 5

- [Property reference](data-sources--ike2--reference--group-001.md#canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b)
- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)

<a id="canonical-8393b8f320c2cee90a4953ad2c752f585cf0195ad52afd3ca9b6c03acbf2ab93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fb950d67282645a587a52438a1ed6eec3901ac753206beb3aace664d09b6b37"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / e9dcc4e1c08d / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)
- [Property reference](data-sources--ike2--reference--group-001.md#canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b)
- ike_keylifetime_minutes

<a id="canonical-749e2a7a5c0ced05915361a6b19eb6b832807bdd4fd6c346d41e21406c79fc4b"></a>

Type: `"single"`. Computed.

Configuration parameter for ike keylifetime minutes.

Upstream description:

Set IKE Key Lifetime in minutes.

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

<a id="canonical-7298da51c1f6eed6f17bd3e03d4566a4833b2a9ff2d63566d97f401668987edb"></a>

## Direct properties — ike_keylifetime_minutes / e9dcc4e1c08d / 3

<a id="canonical-5cd5d0c1a289e733afab4e35ab21cc87fae85c44b2e36ba365183ac4958e9df2"></a>

<a id="canonical-07e2c54e6f3840746336af578e062b032270afbca8497b3884eca745cb46399d"></a>

## duration property — ike_keylifetime_minutes / e9dcc4e1c08d / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-02a39d7e439a8e6099b8bf7fe2b03d353fe0285498c4e99a5a48839f5e1f7b3c"></a>

## Next pages — ike_keylifetime_minutes / e9dcc4e1c08d / 5

- [Property reference](data-sources--ike2--reference--group-001.md#canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b)
- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)

<a id="canonical-919d2f5e851c79682c6da77495f53c90e4429bf1486f6080be53966cd9d0613e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1e052daa60665dc5f8fdc8ae072aef35b8348294ae1f6dadb088e24712a8564"></a>

## use_default_keylifetime — use_default_keylifetime / 61ec79c9528e / 2

Breadcrumbs:

- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)
- [Property reference](data-sources--ike2--reference--group-001.md#canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b)
- use_default_keylifetime

<a id="canonical-12930694ef3e83ad515ffc8e155e5837b96eef76d29b2fbda6be350a7d76d36d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use default keylifetime.

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

<a id="canonical-2e8dd059e63284f249f4cbff2d7da71efb9002510a81635a2b9b62b6f89100a2"></a>

## Direct properties — use_default_keylifetime / 61ec79c9528e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-10c1e79c6c04a8e8b13b93a2660311c1f55c2788720b3b494ba7e11cd6b36cd7"></a>

## Next pages — use_default_keylifetime / 61ec79c9528e / 4

- [Property reference](data-sources--ike2--reference--group-001.md#canonical-a4a0aa390e5bd1f27ab80f543e69ae2f660dcfa29c8069a70181b729523eae7b)
- [xcsh_ike2](../data-sources/ike2.md#canonical-7e3de9947e7e4e80393326cdc4d084c39aa617a7a1d79b8a3deac10eafa5254e)
