---
page_title: "xcsh_ike2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike2 reference."
---

# xcsh_ike2 reference

<a id="canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66a46feb888cdff150a90611e7f67739b06dbde2f3a3df1b6938b571f6847bc4"></a>

## Property reference — Property reference / badb06194c81 / 2

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)
- Property reference

<a id="canonical-4ad04068240bde27c86a38d51c9293336ef31ba47348f88d3aa96833271b6fa7"></a>

## Direct properties — Property reference / badb06194c81 / 3

<a id="canonical-c0415e06d9b9d84be619847492607be05c4e88220528a4e353366f4dd43ab45a"></a>

<a id="canonical-a962a89749e0a5bd2c6ff1e28f281dd11bab434b3666ab688037c9905813eb86"></a>

## annotations property — Property reference / badb06194c81 / 4

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

<a id="canonical-9afb8a3a6c1d5b9b5db874175f89b9a352158d69e09a4b936053423dfabb168e"></a>

<a id="canonical-fb626173c2285bc1dbf3becc13ec33260ed2e14025717aad002f9c581d119c0d"></a>

## description property — Property reference / badb06194c81 / 5

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

- [dh_group_set](resources--ike2--reference--group-001.md#canonical-5e339bdd16fe8d2b187014f64f70ed766a4b7981ee870bae0cadd466c5816ebd): complete subsection reference.

<a id="canonical-956d6db3bdc9ad1ac155783e04be212c8cc562a75ce868693052dc4fb9051e8d"></a>

<a id="canonical-3c255c776caa70ebabb56f0827b89d25eff3ea83b8879239bbddc0d3f198290b"></a>

## disable property — Property reference / badb06194c81 / 6

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

- [disable_pfs](resources--ike2--reference--group-001.md#canonical-db452a5a864facc5635ba638f854455baece53f9f289ad563a00110ba3e1a3f6): complete subsection reference.

<a id="canonical-2aa1e131f4a8b4c19d646faf37f07bfd17e608649ceb6430b6cfe5f7d9311904"></a>

<a id="canonical-b7964e0b7adf6e88bacb23aea5f766bfadee347c10ce0a47ca4b8e70536d40fb"></a>

## id property — Property reference / badb06194c81 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](resources--ike2--reference--group-001.md#canonical-6fa131a6da4f2d7ea789417f2a75c4d735b4afc5fdd8c06fcad31a4b1c60214e): complete subsection reference.

- [ike_keylifetime_minutes](resources--ike2--reference--group-001.md#canonical-d0213f8fa0e23bd507d0687440c3b93cf6afbdd69444aea692c4459cd4c96bfc): complete subsection reference.

<a id="canonical-25e0a39df23172a0ca27556029f91ec384738b8ed2e65a66c86f19ddc94b8298"></a>

<a id="canonical-9912280ea93d6023041f00f21d3c14bacf961dd9c5d7827944bc0de219653003"></a>

## labels property — Property reference / badb06194c81 / 8

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

<a id="canonical-0ba2a789212c2432254c689ef022e1f7f5815d34885addb49f072bbce36ae704"></a>

<a id="canonical-d33cc502c1a69d400f530b0f798f7a677f951fbaf339444fef71d2fa4ab900e7"></a>

## name property — Property reference / badb06194c81 / 9

Type: `"string"`. Required.

Name of the Ike2. Must be unique within the namespace.

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

<a id="canonical-22811540d76ffde9cb3ff8ac5375c7537fe83c01bdb8443254b60f146a21eb25"></a>

<a id="canonical-117f0763c9187d744ce952d02a0bc1f7b6f4d8a498da1d7c21d664641ff0d962"></a>

## namespace property — Property reference / badb06194c81 / 10

Type: `"string"`. Required.

Namespace where the Ike2 is created.

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

- [timeouts](resources--ike2--reference--group-001.md#canonical-27bcf79f2091bc55e83a5b3549a143d857069ddd7509f021f5b7d2b746892eb5): complete subsection reference.

- [use_default_keylifetime](resources--ike2--reference--group-001.md#canonical-233c90a6de00a294b022cefaef0386cffe144b38bbefc7b0a614293814a0ae33): complete subsection reference.

<a id="canonical-7d0f6374c6a42342c70a9d5c0d9fa0aa7fbf2d681d06131c98b9ca57d475a5b2"></a>

## All schema paths — Property reference / badb06194c81 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--ike2--reference--group-001.md#canonical-c0415e06d9b9d84be619847492607be05c4e88220528a4e353366f4dd43ab45a) |
| `description` | [description](resources--ike2--reference--group-001.md#canonical-9afb8a3a6c1d5b9b5db874175f89b9a352158d69e09a4b936053423dfabb168e) |
| `dh_group_set` | [dh_group_set](resources--ike2--reference--group-001.md#canonical-65ee580be247036f4d68b9fc5477b6c46bcad95196f2996b11022b34efce4702) |
| `dh_group_set.dh_groups` | [dh_group_set.dh_groups](resources--ike2--reference--group-001.md#canonical-4f8567e8303a2393f3fd2599f4f06d77810e51081e8403c0ae5f99e5df03e9fd) |
| `disable` | [disable](resources--ike2--reference--group-001.md#canonical-956d6db3bdc9ad1ac155783e04be212c8cc562a75ce868693052dc4fb9051e8d) |
| `disable_pfs` | [disable_pfs](resources--ike2--reference--group-001.md#canonical-b02b07e01090dcdc1e2c638c8d91a64d2a811d6febc78cc423ffee78d9afd282) |
| `id` | [id](resources--ike2--reference--group-001.md#canonical-2aa1e131f4a8b4c19d646faf37f07bfd17e608649ceb6430b6cfe5f7d9311904) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](resources--ike2--reference--group-001.md#canonical-9f3a363e4c1abf7ad737ca377ced50d9e339be1fb09f6e6aeddb79173032aa01) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](resources--ike2--reference--group-001.md#canonical-67f8b8de194f6b4cf3fa7cdfecb2ef61cec3a74780097f4528406b04667ae17d) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](resources--ike2--reference--group-001.md#canonical-c9dc1a9f22df2f672b87511686e684fd4979860cd5ab5f2a1675e6f417a895be) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](resources--ike2--reference--group-001.md#canonical-5cf9435ceef278a7b4452a153d2d8562f0d88b7039f23ef45edf8e748d73808f) |
| `labels` | [labels](resources--ike2--reference--group-001.md#canonical-25e0a39df23172a0ca27556029f91ec384738b8ed2e65a66c86f19ddc94b8298) |
| `name` | [name](resources--ike2--reference--group-001.md#canonical-0ba2a789212c2432254c689ef022e1f7f5815d34885addb49f072bbce36ae704) |
| `namespace` | [namespace](resources--ike2--reference--group-001.md#canonical-22811540d76ffde9cb3ff8ac5375c7537fe83c01bdb8443254b60f146a21eb25) |
| `timeouts` | [timeouts](resources--ike2--reference--group-001.md#canonical-7822711fb6b2192c7868099ce53f137aecae98ba4cd81f8a66ffacc00cd92d34) |
| `timeouts.create` | [timeouts.create](resources--ike2--reference--group-001.md#canonical-e0fe5eafe6ffaa6c3e0e9d5a30f216f758311e2a171dddbfe9615ad45d089c9d) |
| `timeouts.delete` | [timeouts.delete](resources--ike2--reference--group-001.md#canonical-fa8fc6499d66194c872deb035fddec49b0fe7496efdd10e35958aea01a9a6f8a) |
| `timeouts.read` | [timeouts.read](resources--ike2--reference--group-001.md#canonical-413d3af5c8b3121866c63d3775658fdaf88d612edf40fd6bf0aee2bcc85345ce) |
| `timeouts.update` | [timeouts.update](resources--ike2--reference--group-001.md#canonical-214ff2f04a7e73385ecd4848e089d38d5d92a061ed87dd3bcc612e8304284e77) |
| `use_default_keylifetime` | [use_default_keylifetime](resources--ike2--reference--group-001.md#canonical-a2c6a3a1d9cc12059e30239431969cedbee5a82a7d2aed473409e61f4cf8e727) |

<a id="canonical-48bc0f5e8e6e7cb4e0accd4a4cf8dbaa13bc3ea0ece02e025af94b01619ddbbb"></a>

## Next pages — Property reference / badb06194c81 / 12

- [dh_group_set](resources--ike2--reference--group-001.md#canonical-5e339bdd16fe8d2b187014f64f70ed766a4b7981ee870bae0cadd466c5816ebd)
- [disable_pfs](resources--ike2--reference--group-001.md#canonical-db452a5a864facc5635ba638f854455baece53f9f289ad563a00110ba3e1a3f6)
- [ike_keylifetime_hours](resources--ike2--reference--group-001.md#canonical-6fa131a6da4f2d7ea789417f2a75c4d735b4afc5fdd8c06fcad31a4b1c60214e)
- [ike_keylifetime_minutes](resources--ike2--reference--group-001.md#canonical-d0213f8fa0e23bd507d0687440c3b93cf6afbdd69444aea692c4459cd4c96bfc)
- [timeouts](resources--ike2--reference--group-001.md#canonical-27bcf79f2091bc55e83a5b3549a143d857069ddd7509f021f5b7d2b746892eb5)
- [use_default_keylifetime](resources--ike2--reference--group-001.md#canonical-233c90a6de00a294b022cefaef0386cffe144b38bbefc7b0a614293814a0ae33)
- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)

<a id="canonical-5e339bdd16fe8d2b187014f64f70ed766a4b7981ee870bae0cadd466c5816ebd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-75f1dca3c7b75951b1b537cecd8fee11936860d5832b5d8250798f9a3d5fa281"></a>

## dh_group_set — dh_group_set / fe89a209ad5c / 2

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)
- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- dh_group_set

<a id="canonical-65ee580be247036f4d68b9fc5477b6c46bcad95196f2996b11022b34efce4702"></a>

Type: `"object"`. single nested block, Optional.

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

- [dh_group_set](resources--ike2--reference--group-001.md#canonical-65ee580be247036f4d68b9fc5477b6c46bcad95196f2996b11022b34efce4702)
- [disable_pfs](resources--ike2--reference--group-001.md#canonical-b02b07e01090dcdc1e2c638c8d91a64d2a811d6febc78cc423ffee78d9afd282)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dh_group_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-9a92fafc81773adc01a51377137ed59c958aaa4d230f13e9dbb361acf3ed8cb2"></a>

## Direct properties — dh_group_set / fe89a209ad5c / 3

<a id="canonical-4f8567e8303a2393f3fd2599f4f06d77810e51081e8403c0ae5f99e5df03e9fd"></a>

<a id="canonical-069f7e1c698232cddb74d624cf6fdd8c4bf443a570bad804e4506ddf2679d0a4"></a>

## dh_groups property — dh_group_set / fe89a209ad5c / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-7959dd302fd401fc501431cb0be0db5f5e484c559cf738c022ff05d651f92bac"></a>

## Next pages — dh_group_set / fe89a209ad5c / 5

- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)

<a id="canonical-db452a5a864facc5635ba638f854455baece53f9f289ad563a00110ba3e1a3f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-249bffdac4fb28b85cd2656d8e1bf735d5026913ea3b4d088acb6419ba645f14"></a>

## disable_pfs — disable_pfs / f2ae563bb653 / 2

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)
- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- disable_pfs

<a id="canonical-b02b07e01090dcdc1e2c638c8d91a64d2a811d6febc78cc423ffee78d9afd282"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_pfs = {}
```

<a id="canonical-63dad8b852c7572df0d379fad32109f89cccd270b1fa8793dc2ea88614aee294"></a>

## Direct properties — disable_pfs / f2ae563bb653 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1a71dac4d22332c7144fba0968960b2e58030934dfa42fe5a3748bf9d47cd9de"></a>

## Next pages — disable_pfs / f2ae563bb653 / 4

- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)

<a id="canonical-6fa131a6da4f2d7ea789417f2a75c4d735b4afc5fdd8c06fcad31a4b1c60214e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ecd422b3c54149eef96379df0482a3ab0595fb367d7b1c3abcd641d7ee87ba8"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / c6fae54e7994 / 2

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)
- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- ike_keylifetime_hours

<a id="canonical-9f3a363e4c1abf7ad737ca377ced50d9e339be1fb09f6e6aeddb79173032aa01"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Upstream description:

Input Hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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

OneOf alternatives in this subsection:

- [ike_keylifetime_hours](resources--ike2--reference--group-001.md#canonical-9f3a363e4c1abf7ad737ca377ced50d9e339be1fb09f6e6aeddb79173032aa01)
- [ike_keylifetime_minutes](resources--ike2--reference--group-001.md#canonical-c9dc1a9f22df2f672b87511686e684fd4979860cd5ab5f2a1675e6f417a895be)
- [use_default_keylifetime](resources--ike2--reference--group-001.md#canonical-a2c6a3a1d9cc12059e30239431969cedbee5a82a7d2aed473409e61f4cf8e727)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ike_keylifetime_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-5043c4408f4627a42b9f4374fd78556c62864e7ba07fad7ed9085f95d6a9910a"></a>

## Direct properties — ike_keylifetime_hours / c6fae54e7994 / 3

<a id="canonical-67f8b8de194f6b4cf3fa7cdfecb2ef61cec3a74780097f4528406b04667ae17d"></a>

<a id="canonical-fda4a5ce311de26fa885a3a7d6929f49126960131dd37077e09e50fdca11f5cd"></a>

## duration property — ike_keylifetime_hours / c6fae54e7994 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

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

<a id="canonical-3ecb6d61a6556c6eb54da971407240b823d70994513b20470d609d418d535425"></a>

## Next pages — ike_keylifetime_hours / c6fae54e7994 / 5

- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)

<a id="canonical-d0213f8fa0e23bd507d0687440c3b93cf6afbdd69444aea692c4459cd4c96bfc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82cebcd5c0510566c6a2fa30d64828334c760c44f68b1fddc773e436eaac0015"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / bf2ea3d1aa5e / 2

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)
- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- ike_keylifetime_minutes

<a id="canonical-c9dc1a9f22df2f672b87511686e684fd4979860cd5ab5f2a1675e6f417a895be"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ike keylifetime minutes.

Upstream description:

Set IKE Key Lifetime in minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
ike_keylifetime_minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-7f220cb76e0ae57cb654d1d3cf28c9e90950155cf3a65221609924bd1b3e31b3"></a>

## Direct properties — ike_keylifetime_minutes / bf2ea3d1aa5e / 3

<a id="canonical-5cf9435ceef278a7b4452a153d2d8562f0d88b7039f23ef45edf8e748d73808f"></a>

<a id="canonical-b631fb244d6ea404ca1274261ff6e24e50718deb123276c3b930f641ba6f3d06"></a>

## duration property — ike_keylifetime_minutes / bf2ea3d1aa5e / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(10, 300),
}
```

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

<a id="canonical-14ee88acc995197c11e091d4ff317d443319857ea12021ddddbf46c46dd55a02"></a>

## Next pages — ike_keylifetime_minutes / bf2ea3d1aa5e / 5

- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)

<a id="canonical-27bcf79f2091bc55e83a5b3549a143d857069ddd7509f021f5b7d2b746892eb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a641411ba32363e677001bec855bea3909a64074e35d88b19e03d255775ab79"></a>

## timeouts — timeouts / 646506bb41f2 / 2

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)
- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- timeouts

<a id="canonical-7822711fb6b2192c7868099ce53f137aecae98ba4cd81f8a66ffacc00cd92d34"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-535c68955c01b37c40a0903db97b53dda6dac8f6cd9723542819a7958a79f939"></a>

## Direct properties — timeouts / 646506bb41f2 / 3

<a id="canonical-e0fe5eafe6ffaa6c3e0e9d5a30f216f758311e2a171dddbfe9615ad45d089c9d"></a>

<a id="canonical-993762880e42489fb2033d4d83b949ef58e8680cca0c43f0311f27d409ea3cab"></a>

## create property — timeouts / 646506bb41f2 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-fa8fc6499d66194c872deb035fddec49b0fe7496efdd10e35958aea01a9a6f8a"></a>

<a id="canonical-a5e4c1def6bc1ac6051d9311e9306448ada5d6bf3a35e2b6ab5bdd17b8ee0509"></a>

## delete property — timeouts / 646506bb41f2 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-413d3af5c8b3121866c63d3775658fdaf88d612edf40fd6bf0aee2bcc85345ce"></a>

<a id="canonical-da0eb726f307d8ba8f6f16534ac843abaf1a51b1da0ce65d835fd2042e3f32fa"></a>

## read property — timeouts / 646506bb41f2 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-214ff2f04a7e73385ecd4848e089d38d5d92a061ed87dd3bcc612e8304284e77"></a>

<a id="canonical-6f8e97905c2ef04d2d31ee43b0dc7fe4d74ee77db0e207531ca291cac87488d0"></a>

## update property — timeouts / 646506bb41f2 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-075d1de5ee0ba765bd919e02e127139798c5fdcdf14acaa782acbfaeca340dee"></a>

## Next pages — timeouts / 646506bb41f2 / 8

- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)

<a id="canonical-233c90a6de00a294b022cefaef0386cffe144b38bbefc7b0a614293814a0ae33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12f1ee60fc49dcd40e410bc59906a7bd33a7ef4b6bfed948c9eb7bd2a0d5bffc"></a>

## use_default_keylifetime — use_default_keylifetime / 9df1f2cdb580 / 2

Breadcrumbs:

- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)
- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- use_default_keylifetime

<a id="canonical-a2c6a3a1d9cc12059e30239431969cedbee5a82a7d2aed473409e61f4cf8e727"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_default_keylifetime = {}
```

<a id="canonical-3394ad623c8bc053799bc455cf584966026b9c1a2974fa9fb22551d209d1b847"></a>

## Direct properties — use_default_keylifetime / 9df1f2cdb580 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b57d0942274d1f79f1d31b07daab76dc0bcafea8657427e3b066edaea03560eb"></a>

## Next pages — use_default_keylifetime / 9df1f2cdb580 / 4

- [Property reference](resources--ike2--reference--group-001.md#canonical-1940acee621836081acda38181ec6599d2a13a430e9017ab699561da79b09b23)
- [xcsh_ike2](../resources/ike2.md#canonical-d3a1c6950be24c85da4b538a2b3d672b6190ed3f7b6b9b23208a094e5821b06e)
