---
page_title: "xcsh_ike_phase1_profile reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase1_profile reference."
---

# xcsh_ike_phase1_profile reference

<a id="canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf029b0ec88927fee71ec8bc25514320e1980861d71b5e5b231c695fbd8806e8"></a>

## Property reference — Property reference / 787870ad32c6 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
- Property reference

<a id="canonical-476868a5ebd6462c066004dc1c0595931d8cfd8cd5d6947769e2121060811a25"></a>

## Direct properties — Property reference / 787870ad32c6 / 3

<a id="canonical-1ab9d31702691815e7036c37e66afab21ae8358184c58609f64341ba66dacabb"></a>

<a id="canonical-601edc4c6e945cddb2ab730672c26343897a3298a09f2ea6db76c745f4ab8bbf"></a>

## annotations property — Property reference / 787870ad32c6 / 4

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

<a id="canonical-de2069c52c613383a9b39f79edc04a9a070b9a7edc2b14c81b8c63ed4d9c1ca8"></a>

<a id="canonical-c52406a5fc181ed23a62b3405ba844b9c7420a6415239f87476dc84cd8ff49cd"></a>

## authentication_algos property — Property reference / 787870ad32c6 / 5

Type: `["list", "string"]`. Required.

\[Enum: AUTH\_ALG\_DEFAULT|SHA256\_HMAC|SHA384\_HMAC|SHA512\_HMAC|AUTH\_ALG\_NONE\] Choose one or
more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm encryption
algorithms. Possible values are \`AUTH\_ALG\_DEFAULT\`, \`SHA256\_HMAC\`, \`SHA384\_HMAC\`,
\`SHA512\_HMAC\`, \`AUTH\_ALG\_NONE\`. Defaults to \`AUTH\_ALG\_DEFAULT\`.

Upstream description:

Choose one or more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm
encryption algorithms.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-9e0d85766092af0ec9f8cf82142186d256c72db62bc7c96e901fdd6b40d2856e"></a>

<a id="canonical-f3aa3808b554ee5dfec8b5698d1f470d432873fde281790cfce86e15229dc98e"></a>

## description property — Property reference / 787870ad32c6 / 6

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

<a id="canonical-f9c33a66a4a659e9431356b7c29fc29891345240ab8f519937ae5328d8081850"></a>

<a id="canonical-a98667fff21c0960a5052e6ce481b3f72cfad53535c29c31f884704ce47be7bb"></a>

## dh_group property — Property reference / 787870ad32c6 / 7

Type: `["list", "string"]`. Required.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Choose the acceptable Diffie Hellman (DH) Group or Groups that you are willing to accept as part of
this profile. Possible values are \`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`,
\`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`, \`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`,
\`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`. Defaults to \`DH\_GROUP\_DEFAULT\`.

Upstream description:

Choose the acceptable Diffie Hellman (DH) Group or Groups that you are willing to accept as part of
this profile.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-39fdeeff5d76e631c9c3f8536299be1867b9a44af8ccf62bb4a4248edd1c38d4"></a>

<a id="canonical-f43ee2575951783808d25791b73cd6817b675a89d14539b4c8bafe6b2dc8282c"></a>

## disable property — Property reference / 787870ad32c6 / 8

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

<a id="canonical-e5aec92100dc5782974306df9d5b635e24d8d552bff86a45cf007f5e82dbafab"></a>

<a id="canonical-66658d031ab6ede8e4e8e934fdfe2f8f6ce7d4d49fbcae26bfcd0dbc9b921c3d"></a>

## encryption_algos property — Property reference / 787870ad32c6 / 9

Type: `["list", "string"]`. Required.

\[Enum:
ENC\_ALG\_DEFAULT|AES128\_CBC|AES192\_CBC|AES256\_CBC|TRIPLE\_DES\_CBC|AES128\_GCM|AES192\_GCM|AES256\_GCM\]
Choose one or more encryption algorithms. Possible values are \`ENC\_ALG\_DEFAULT\`,
\`AES128\_CBC\`, \`AES192\_CBC\`, \`AES256\_CBC\`, \`TRIPLE\_DES\_CBC\`, \`AES128\_GCM\`,
\`AES192\_GCM\`, \`AES256\_GCM\`. Defaults to \`ENC\_ALG\_DEFAULT\`.

Upstream description:

Choose one or more encryption algorithms.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-3a8f55f6b2dae65143391ec36745478884e5733078907579bb706cadb3574939"></a>

<a id="canonical-ca02162fbf54e360d851d29ce03f60a24fb90a13fddcab43df6c77ffde807c2a"></a>

## id property — Property reference / 787870ad32c6 / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-2222421004ec11ebd2b2d46d3e23cc31dd351eea9a96f8518ecd3c26bff7b9d1): complete subsection reference.

- [ike_keylifetime_minutes](resources--ike_phase1_profile--reference--group-001.md#canonical-6ffb72544c16ce0507532398154c65ff32bce0c799cfd9bb2515e9049d43e199): complete subsection reference.

<a id="canonical-26d99353d1b5ff8ea3feae00387720c855f04a3c61c23c4991ca5abe2b4e13cc"></a>

<a id="canonical-c3601243e0cd5b1a2ca1819406829ebd533bffe89cced31d6abb24a492eb4b8c"></a>

## labels property — Property reference / 787870ad32c6 / 11

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

<a id="canonical-4d93ed640d53fcc7c16454325ec358138efa1a0eec862b6618ea3d4a6c27480f"></a>

<a id="canonical-5be734a6a9ab3fea8a9de910f97ddbd823768ae1bf45381801eff01ef035c973"></a>

## name property — Property reference / 787870ad32c6 / 12

Type: `"string"`. Required.

Name of the IKE Phase1 Profile. Must be unique within the namespace.

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

<a id="canonical-c6963ec3363b26f8f21ff0a2bd5d5531098755971ea531e1d25c3f847b092f75"></a>

<a id="canonical-6065ac03e48290123d44d7df789cb6df2f3d6b74293096fdd660cd578eaa658a"></a>

## namespace property — Property reference / 787870ad32c6 / 13

Type: `"string"`. Required.

Namespace where the IKE Phase1 Profile is created.

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

<a id="canonical-250d62e370a7e4b0f81a97bf2a239557ff28a99792c0d5157c2bbe04447d23b2"></a>

<a id="canonical-5f42a47ca1f25e0fc98d7b465c09a3ca282476c62e3801cda442924607d01496"></a>

## prf property — Property reference / 787870ad32c6 / 14

Type: `["list", "string"]`. Required.

\[Enum: PRF\_DEFAULT|PRFSHA256|PRFSHA384|PRFSHA512\] PseudoRandomFunction. Select
PseudoRandomFunction for IKE SA. Possible values are \`PRF\_DEFAULT\`, \`PRFSHA256\`, \`PRFSHA384\`,
\`PRFSHA512\`. Defaults to \`PRF\_DEFAULT\`.

Upstream description:

Select PseudoRandomFunction for IKE SA.

Receipt-pinned upstream constraints:

```json
{
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

- [reauth_disabled](resources--ike_phase1_profile--reference--group-001.md#canonical-0b7de1a5a56ca3e7ec7cbaa21bb5a234c8a9008270ff2cfb8a750584aac8843c): complete subsection reference.

- [reauth_timeout_days](resources--ike_phase1_profile--reference--group-001.md#canonical-cab849a5f5d94a96a607b7b1fa6d7476c9e35f90efb63cf88512baba60e9c1aa): complete subsection reference.

- [reauth_timeout_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-b827c49f0fb50055150d58171a841e245c84dd2963b015922532eaaa6c7ff1a1): complete subsection reference.

- [timeouts](resources--ike_phase1_profile--reference--group-001.md#canonical-128ae7f8e31415a7d949724f98c6794f3665e4b887208e1414e35b2abd6932c4): complete subsection reference.

- [use_default_keylifetime](resources--ike_phase1_profile--reference--group-001.md#canonical-ced59d5fc7ac84cff6c37443396541936b62977f50ad5c607a3a0917ff1fc245): complete subsection reference.

<a id="canonical-4956e85adcd519e8ec7dbfd1483d390a59c7d9ac0330741d99c6fa5659303f31"></a>

## All schema paths — Property reference / 787870ad32c6 / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--ike_phase1_profile--reference--group-001.md#canonical-1ab9d31702691815e7036c37e66afab21ae8358184c58609f64341ba66dacabb) |
| `authentication_algos` | [authentication_algos](resources--ike_phase1_profile--reference--group-001.md#canonical-de2069c52c613383a9b39f79edc04a9a070b9a7edc2b14c81b8c63ed4d9c1ca8) |
| `description` | [description](resources--ike_phase1_profile--reference--group-001.md#canonical-9e0d85766092af0ec9f8cf82142186d256c72db62bc7c96e901fdd6b40d2856e) |
| `dh_group` | [dh_group](resources--ike_phase1_profile--reference--group-001.md#canonical-f9c33a66a4a659e9431356b7c29fc29891345240ab8f519937ae5328d8081850) |
| `disable` | [disable](resources--ike_phase1_profile--reference--group-001.md#canonical-39fdeeff5d76e631c9c3f8536299be1867b9a44af8ccf62bb4a4248edd1c38d4) |
| `encryption_algos` | [encryption_algos](resources--ike_phase1_profile--reference--group-001.md#canonical-e5aec92100dc5782974306df9d5b635e24d8d552bff86a45cf007f5e82dbafab) |
| `id` | [id](resources--ike_phase1_profile--reference--group-001.md#canonical-3a8f55f6b2dae65143391ec36745478884e5733078907579bb706cadb3574939) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-fb6b8d637a837ad94600a89ee62f7a5eb29367ab699133188f3332408649624d) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](resources--ike_phase1_profile--reference--group-001.md#canonical-823681fe082c15568fc56f57338fa20e9abc7cae708d1ba1a93cd4e6d83341b0) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](resources--ike_phase1_profile--reference--group-001.md#canonical-9c8417fee8bec268c6a6f5ffd878302cec708a1d35996585b8122543adb2c608) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](resources--ike_phase1_profile--reference--group-001.md#canonical-1706060587b7009fa9e7126ac06fdf571aa5326f56765426e23b180ec819b9dd) |
| `labels` | [labels](resources--ike_phase1_profile--reference--group-001.md#canonical-26d99353d1b5ff8ea3feae00387720c855f04a3c61c23c4991ca5abe2b4e13cc) |
| `name` | [name](resources--ike_phase1_profile--reference--group-001.md#canonical-4d93ed640d53fcc7c16454325ec358138efa1a0eec862b6618ea3d4a6c27480f) |
| `namespace` | [namespace](resources--ike_phase1_profile--reference--group-001.md#canonical-c6963ec3363b26f8f21ff0a2bd5d5531098755971ea531e1d25c3f847b092f75) |
| `prf` | [prf](resources--ike_phase1_profile--reference--group-001.md#canonical-250d62e370a7e4b0f81a97bf2a239557ff28a99792c0d5157c2bbe04447d23b2) |
| `reauth_disabled` | [reauth_disabled](resources--ike_phase1_profile--reference--group-001.md#canonical-2e6edcc044522d80c4fb308b90fa43e252ed3ae7c7cfdfa45c8f0d5a195bcd2d) |
| `reauth_timeout_days` | [reauth_timeout_days](resources--ike_phase1_profile--reference--group-001.md#canonical-35521b39d388e0714048e806eacd4bf385dc5f19d1d8fe6ba36ea22ca41736e4) |
| `reauth_timeout_days.duration` | [reauth_timeout_days.duration](resources--ike_phase1_profile--reference--group-001.md#canonical-bb5526b4b54f986a081e031ad26ac85d366549f84729898d10ffdc43c981790a) |
| `reauth_timeout_hours` | [reauth_timeout_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-e85804da5e104565c8e54dd2d26241234bc24f91409a196c1ecc0d844f5c87b8) |
| `reauth_timeout_hours.duration` | [reauth_timeout_hours.duration](resources--ike_phase1_profile--reference--group-001.md#canonical-546b950c8a97e0b5af7668772065cb7f233ea6400f714d9d232bc5930e02751e) |
| `timeouts` | [timeouts](resources--ike_phase1_profile--reference--group-001.md#canonical-f1497e4dbcce4c0bda73cd8c9b6e2ec38ebd1368480899b93de27cfe99739f91) |
| `timeouts.create` | [timeouts.create](resources--ike_phase1_profile--reference--group-001.md#canonical-17e73d1835f5c442e962987086a583a2a994bce532e2a12529374c296cbc7fdc) |
| `timeouts.delete` | [timeouts.delete](resources--ike_phase1_profile--reference--group-001.md#canonical-0020cf6397b06cee977edb24f999c76491722f4ba0c47604e762e19d67c9e906) |
| `timeouts.read` | [timeouts.read](resources--ike_phase1_profile--reference--group-001.md#canonical-0f6c7da2276f38e17f549667ebfe04314318be50367c05d08ce4b5da91df7fce) |
| `timeouts.update` | [timeouts.update](resources--ike_phase1_profile--reference--group-001.md#canonical-81438ef880141022b49b9a4094e635713dd7df7f9154ce1d86f9e357dc65506b) |
| `use_default_keylifetime` | [use_default_keylifetime](resources--ike_phase1_profile--reference--group-001.md#canonical-bdf5ff75a218f6c604646467a24babcad42c04cd1aa13487bec8d4bf42dab397) |

<a id="canonical-d5a8f10d5e8cc32f78bca8cd344263cf7ca83de37abc81fe706539fff42a589b"></a>

## Next pages — Property reference / 787870ad32c6 / 16

- [ike_keylifetime_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-2222421004ec11ebd2b2d46d3e23cc31dd351eea9a96f8518ecd3c26bff7b9d1)
- [ike_keylifetime_minutes](resources--ike_phase1_profile--reference--group-001.md#canonical-6ffb72544c16ce0507532398154c65ff32bce0c799cfd9bb2515e9049d43e199)
- [reauth_disabled](resources--ike_phase1_profile--reference--group-001.md#canonical-0b7de1a5a56ca3e7ec7cbaa21bb5a234c8a9008270ff2cfb8a750584aac8843c)
- [reauth_timeout_days](resources--ike_phase1_profile--reference--group-001.md#canonical-cab849a5f5d94a96a607b7b1fa6d7476c9e35f90efb63cf88512baba60e9c1aa)
- [reauth_timeout_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-b827c49f0fb50055150d58171a841e245c84dd2963b015922532eaaa6c7ff1a1)
- [timeouts](resources--ike_phase1_profile--reference--group-001.md#canonical-128ae7f8e31415a7d949724f98c6794f3665e4b887208e1414e35b2abd6932c4)
- [use_default_keylifetime](resources--ike_phase1_profile--reference--group-001.md#canonical-ced59d5fc7ac84cff6c37443396541936b62977f50ad5c607a3a0917ff1fc245)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)

<a id="canonical-2222421004ec11ebd2b2d46d3e23cc31dd351eea9a96f8518ecd3c26bff7b9d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07cd63eafd43b1c50f0cf5d9169e8e3943e8fb8042f02ac7f67791992a486a20"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / e50d818aeb2a / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- ike_keylifetime_hours

<a id="canonical-fb6b8d637a837ad94600a89ee62f7a5eb29367ab699133188f3332408649624d"></a>

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

- [ike_keylifetime_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-fb6b8d637a837ad94600a89ee62f7a5eb29367ab699133188f3332408649624d)
- [ike_keylifetime_minutes](resources--ike_phase1_profile--reference--group-001.md#canonical-9c8417fee8bec268c6a6f5ffd878302cec708a1d35996585b8122543adb2c608)
- [use_default_keylifetime](resources--ike_phase1_profile--reference--group-001.md#canonical-bdf5ff75a218f6c604646467a24babcad42c04cd1aa13487bec8d4bf42dab397)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ike_keylifetime_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-b008c64098d7ac4a7f34394095c7c5e7359c3f4cb18c8c7f16c1864c0925ffb7"></a>

## Direct properties — ike_keylifetime_hours / e50d818aeb2a / 3

<a id="canonical-823681fe082c15568fc56f57338fa20e9abc7cae708d1ba1a93cd4e6d83341b0"></a>

<a id="canonical-30e01fd88b44be3e4c2b4275189eb450e31fc36aeb8cbfc2ee3e06f3614be75e"></a>

## duration property — ike_keylifetime_hours / e50d818aeb2a / 4

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

<a id="canonical-def527e59aabcabbb9167180a308f0ee1c7a33e553ba2f5f9a4dd3757a5289f1"></a>

## Next pages — ike_keylifetime_hours / e50d818aeb2a / 5

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)

<a id="canonical-6ffb72544c16ce0507532398154c65ff32bce0c799cfd9bb2515e9049d43e199"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10fa56820cea2744edf6e04ddb84ad81a9d289d75ed673c1cc9a8b1c1da70956"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / 37c41d39da87 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- ike_keylifetime_minutes

<a id="canonical-9c8417fee8bec268c6a6f5ffd878302cec708a1d35996585b8122543adb2c608"></a>

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

<a id="canonical-f25902e92db722d139b83c2213919ef24f1c865e2fd06cce192eaeac4decd008"></a>

## Direct properties — ike_keylifetime_minutes / 37c41d39da87 / 3

<a id="canonical-1706060587b7009fa9e7126ac06fdf571aa5326f56765426e23b180ec819b9dd"></a>

<a id="canonical-27dee3b0854f3f5722cfeff682b87b572eea38b7f175218575c2481d4497f71d"></a>

## duration property — ike_keylifetime_minutes / 37c41d39da87 / 4

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

<a id="canonical-2cdf5bcac17e247f95994d8cb5ca19bdd92df3b1c9b03c1a9d90b42afb1efee5"></a>

## Next pages — ike_keylifetime_minutes / 37c41d39da87 / 5

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)

<a id="canonical-0b7de1a5a56ca3e7ec7cbaa21bb5a234c8a9008270ff2cfb8a750584aac8843c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ae5590df742756545d90bbfed38f45606876d5d5dadfcae3c45338325e68ed0"></a>

## reauth_disabled — reauth_disabled / 4a2e2964fa6c / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- reauth_disabled

<a id="canonical-2e6edcc044522d80c4fb308b90fa43e252ed3ae7c7cfdfa45c8f0d5a195bcd2d"></a>

Type: `["object", {}]`. Optional.

\[OneOf: reauth\_disabled, reauth\_timeout\_days, reauth\_timeout\_hours\] Enable this option

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

OneOf alternatives in this subsection:

- [reauth_disabled](resources--ike_phase1_profile--reference--group-001.md#canonical-2e6edcc044522d80c4fb308b90fa43e252ed3ae7c7cfdfa45c8f0d5a195bcd2d)
- [reauth_timeout_days](resources--ike_phase1_profile--reference--group-001.md#canonical-35521b39d388e0714048e806eacd4bf385dc5f19d1d8fe6ba36ea22ca41736e4)
- [reauth_timeout_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-e85804da5e104565c8e54dd2d26241234bc24f91409a196c1ecc0d844f5c87b8)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
reauth_disabled = {}
```

<a id="canonical-f8f6a520bda2893048802acad2e65d2cb5a0520ba5a8d16e867633b7ad63bddf"></a>

## Direct properties — reauth_disabled / 4a2e2964fa6c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3bc7ac20294e64759b0fae5df44384dbfbc4ebff9fc0dd0a3ddc441660164907"></a>

## Next pages — reauth_disabled / 4a2e2964fa6c / 4

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)

<a id="canonical-cab849a5f5d94a96a607b7b1fa6d7476c9e35f90efb63cf88512baba60e9c1aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29682ba9bc65c1df4fa59baa91baa86e9764782079373983ab1f7d7c72ef0d38"></a>

## reauth_timeout_days — reauth_timeout_days / 8927537b2da1 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- reauth_timeout_days

<a id="canonical-35521b39d388e0714048e806eacd4bf385dc5f19d1d8fe6ba36ea22ca41736e4"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout days.

Upstream description:

Set Duration in days.

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
reauth_timeout_days {
  # Configure direct properties listed below.
}
```

<a id="canonical-e1ac9f0a4e24ccc351331b5a5fe5d6c492ee95296b557236c0a16f1536d64973"></a>

## Direct properties — reauth_timeout_days / 8927537b2da1 / 3

<a id="canonical-bb5526b4b54f986a081e031ad26ac85d366549f84729898d10ffdc43c981790a"></a>

<a id="canonical-14f9e4fda6a859d2ad909b68b129fb44da319f3d705ff9c39d77b1e21b4a5574"></a>

## duration property — reauth_timeout_days / 8927537b2da1 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

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

<a id="canonical-8210e47eebe3581bd518849040ce80be78136cf2efce4cb0865c49fea3d67507"></a>

## Next pages — reauth_timeout_days / 8927537b2da1 / 5

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)

<a id="canonical-b827c49f0fb50055150d58171a841e245c84dd2963b015922532eaaa6c7ff1a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6384812953328327e2b2c354f31763d33fe8ff90d5e7e0d59a84abec06b4745c"></a>

## reauth_timeout_hours — reauth_timeout_hours / 30318928ebed / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- reauth_timeout_hours

<a id="canonical-e85804da5e104565c8e54dd2d26241234bc24f91409a196c1ecc0d844f5c87b8"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout hours.

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

Terraform syntax:

```terraform
reauth_timeout_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-a596993a0f54599b71c07020606a3964a18c776c935234c2f55dcd504f256f29"></a>

## Direct properties — reauth_timeout_hours / 30318928ebed / 3

<a id="canonical-546b950c8a97e0b5af7668772065cb7f233ea6400f714d9d232bc5930e02751e"></a>

<a id="canonical-c07077208b37790d096fe36d7aaae14ac49dbdefc580df2885f30fe3cb7f462d"></a>

## duration property — reauth_timeout_hours / 30318928ebed / 4

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

<a id="canonical-2093902e077b711aebb8a71fa632374b5d476cf652c968b3427bf896541954d1"></a>

## Next pages — reauth_timeout_hours / 30318928ebed / 5

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)

<a id="canonical-128ae7f8e31415a7d949724f98c6794f3665e4b887208e1414e35b2abd6932c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6902820fef5a2ce1acaa72abc43bbd1bc4fbf16e78420032ffa75665cabd49c6"></a>

## timeouts — timeouts / e9462d7a829e / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- timeouts

<a id="canonical-f1497e4dbcce4c0bda73cd8c9b6e2ec38ebd1368480899b93de27cfe99739f91"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-255790a4a560d608971d21b84ecef02422905599347c98ab60441eee0310c821"></a>

## Direct properties — timeouts / e9462d7a829e / 3

<a id="canonical-17e73d1835f5c442e962987086a583a2a994bce532e2a12529374c296cbc7fdc"></a>

<a id="canonical-0ac0ca2a6309447c5e6267790d849ad70cbd5e8434ddb976b85d36752bde2e6a"></a>

## create property — timeouts / e9462d7a829e / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0020cf6397b06cee977edb24f999c76491722f4ba0c47604e762e19d67c9e906"></a>

<a id="canonical-1ec46fb9f63d61a5e41b817fd46b3391605f9dd00efee6863e31ff604dfe1932"></a>

## delete property — timeouts / e9462d7a829e / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0f6c7da2276f38e17f549667ebfe04314318be50367c05d08ce4b5da91df7fce"></a>

<a id="canonical-0227969e452806f34689255d771f385dc6ed120bb467e28b5903f097465664e1"></a>

## read property — timeouts / e9462d7a829e / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-81438ef880141022b49b9a4094e635713dd7df7f9154ce1d86f9e357dc65506b"></a>

<a id="canonical-234f7f6a84de45b50abe59c6ec1d708d2770c74cd427441956e450f590681c43"></a>

## update property — timeouts / e9462d7a829e / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-40d2fe4a214de94d605309dbf98d30813269a6313fc818d8ace62432ae2cfe4e"></a>

## Next pages — timeouts / e9462d7a829e / 8

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)

<a id="canonical-ced59d5fc7ac84cff6c37443396541936b62977f50ad5c607a3a0917ff1fc245"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f01790b6773d01069de57e468b2aeac57cb506e8a3dc074cf505fcd09cc7e88c"></a>

## use_default_keylifetime — use_default_keylifetime / 4490e92c8f3b / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- use_default_keylifetime

<a id="canonical-bdf5ff75a218f6c604646467a24babcad42c04cd1aa13487bec8d4bf42dab397"></a>

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

<a id="canonical-267487d9f95168e4fa40a4692d169bc20ee5e0956ab7cf06a96067a984932e50"></a>

## Direct properties — use_default_keylifetime / 4490e92c8f3b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0724daa277c6b965442b7ede73a3e6daabc9c60b9da6bc733f072f0d5675aabc"></a>

## Next pages — use_default_keylifetime / 4490e92c8f3b / 4

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-f8aa51677734852141980a8fdcb5d03b36b5015e8308abd6ce5e3672da33b979)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-a677d0cdc219c03f7400af85a6c04cc688cf4937e0d720a036517652eef93f68)
