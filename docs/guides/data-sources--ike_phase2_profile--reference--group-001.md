---
page_title: "xcsh_ike_phase2_profile reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile reference."
---

# xcsh_ike_phase2_profile reference

<a id="canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-961eee97d7456e05e3400aa09c448d8c871d7846fd91a4f73af24c12cf34561d"></a>

## Property reference — Property reference / 5ede625c02f0 / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)
- Property reference

<a id="canonical-14f4c1b76e7128fc4a697ea6f76c773a9e16689e723eb1fa828f62bd17a0b09d"></a>

## Direct properties — Property reference / 5ede625c02f0 / 3

<a id="canonical-c1c50143545f426545c1a767bf7ad40a850cb3f150d13be79748ec408dd6cb28"></a>

<a id="canonical-6bbc50c59f9690d46690fa4f67fdcfbf6dc1614ea825c32412061741f2edb75b"></a>

## annotations property — Property reference / 5ede625c02f0 / 4

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

<a id="canonical-cbd146862f8416371ab48d25d15a04497b6a29cdc09265d828cbd2f81699d503"></a>

<a id="canonical-b03ec0d195c8dac5275963f1821460e2222a159f9ff125195d1473c0ff064e1e"></a>

## authentication_algos property — Property reference / 5ede625c02f0 / 5

Type: `["list", "string"]`. Computed.

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

<a id="canonical-c6eb12f1e81e6e16c60a7ce0d061f373263fdf31729b540a83caed972671bc9e"></a>

<a id="canonical-69737b264f409dc7ff7485333d904510befcda176bd91ce1974efc2a7cda734a"></a>

## description property — Property reference / 5ede625c02f0 / 6

Type: `"string"`. Computed.

Description of the IKEPhase2Profile.

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

- [dh_group_set](data-sources--ike_phase2_profile--reference--group-001.md#canonical-156d9a0f217a28938f835e169bc3a0b6270fe27506537f5a0e1bcfeb942bc0c7): complete subsection reference.

- [disable_pfs](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0d0e01b46d38aa6e76991b7190d35a676cace75d45ac86a67748307ba1296d93): complete subsection reference.

<a id="canonical-cd571e2157c7493306a1834f324819d934924485536dacd38520af118f679e39"></a>

<a id="canonical-b3a5ac7d8b8cfa1cd28f88ba6b5da74388a9e77305bda274a88822fd88a7f724"></a>

## encryption_algos property — Property reference / 5ede625c02f0 / 7

Type: `["list", "string"]`. Computed.

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

<a id="canonical-0a5414f0ada7f06bf7a4c451b54c14df56e5fd8becfb169167f5b043b4c62718"></a>

<a id="canonical-8143788259229c839c50b7de70b34f33a4e77664004d62123bc83782bafa7b9e"></a>

## id property — Property reference / 5ede625c02f0 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](data-sources--ike_phase2_profile--reference--group-001.md#canonical-29b6aab12573f1a303b3830d440e12a649eee2ecd2c6972002b28b2a67392249): complete subsection reference.

- [ike_keylifetime_minutes](data-sources--ike_phase2_profile--reference--group-001.md#canonical-d42d0d39e501e98365077d37bb79eb6b9d38dd374ae3637c32e607749ffd390a): complete subsection reference.

<a id="canonical-eb8fc9118e871ac8027bc933c2bdfd5cb28684911b1e65f6319b6c61d38c87d6"></a>

<a id="canonical-a07138cd618495f319f64f617726213a1a6c9a85b1a51154616f9d339f2052fc"></a>

## labels property — Property reference / 5ede625c02f0 / 9

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

<a id="canonical-1d763b0275d454660641ea8e1e8066a00e8e3216f2ae61dba1f22e3cc31a1f84"></a>

<a id="canonical-2a9b6792350674ae6f3931eb3426b7b4dee88d71883a670878d26b2040462861"></a>

## name property — Property reference / 5ede625c02f0 / 10

Type: `"string"`. Required.

Name of the IKEPhase2Profile.

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

<a id="canonical-389d385de8b7a12b82f595c07d0c27044607130d816712c4b6d657409506915b"></a>

<a id="canonical-14dec5e42cd239b607d6945ec8c5fd22984f68edeff4938f9cf49183adb5c61a"></a>

## namespace property — Property reference / 5ede625c02f0 / 11

Type: `"string"`. Required.

Namespace where the IKEPhase2Profile exists.

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

- [use_default_keylifetime](data-sources--ike_phase2_profile--reference--group-001.md#canonical-5d945860411de2759e392e5da68da15687751f7e57674e23ac9a5022825fea93): complete subsection reference.

<a id="canonical-d2693a2aa52d1f13cf929386ba4b3d0c839d87b2b4f03ab22fbc1644b0e7257f"></a>

## All schema paths — Property reference / 5ede625c02f0 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--ike_phase2_profile--reference--group-001.md#canonical-c1c50143545f426545c1a767bf7ad40a850cb3f150d13be79748ec408dd6cb28) |
| `authentication_algos` | [authentication_algos](data-sources--ike_phase2_profile--reference--group-001.md#canonical-cbd146862f8416371ab48d25d15a04497b6a29cdc09265d828cbd2f81699d503) |
| `description` | [description](data-sources--ike_phase2_profile--reference--group-001.md#canonical-c6eb12f1e81e6e16c60a7ce0d061f373263fdf31729b540a83caed972671bc9e) |
| `dh_group_set` | [dh_group_set](data-sources--ike_phase2_profile--reference--group-001.md#canonical-714b8055d514dc40a60d0eb257245761c0fd75b51557bdfc48feaddf8b243b13) |
| `dh_group_set.dh_groups` | [dh_group_set.dh_groups](data-sources--ike_phase2_profile--reference--group-001.md#canonical-ae5c7c44ab9bd52d115582ac67e70a34c08109837ffa8f63f997fef2bd3b5e11) |
| `disable_pfs` | [disable_pfs](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1ff975c5328084705c5ec883ac6fd2f96a499da4a37ad11a7f4ca5f475643d59) |
| `encryption_algos` | [encryption_algos](data-sources--ike_phase2_profile--reference--group-001.md#canonical-cd571e2157c7493306a1834f324819d934924485536dacd38520af118f679e39) |
| `id` | [id](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0a5414f0ada7f06bf7a4c451b54c14df56e5fd8becfb169167f5b043b4c62718) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](data-sources--ike_phase2_profile--reference--group-001.md#canonical-a6042d88be62d5fccff921280dbf74ef3a45e3c88a8cef475a378f99d86c4898) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](data-sources--ike_phase2_profile--reference--group-001.md#canonical-609befcb21b9756e6fa7574835971c5ecd31b1a5c97a9aeb811b0658f7984d95) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](data-sources--ike_phase2_profile--reference--group-001.md#canonical-6d51cc2e95d0dfea81feadf596e3fc79b932dfc07415b942e8df0ac9160745d1) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](data-sources--ike_phase2_profile--reference--group-001.md#canonical-2615bde362805f1a3251914282ca9ae1beed01e258ff9bbeacc685ad88732a33) |
| `labels` | [labels](data-sources--ike_phase2_profile--reference--group-001.md#canonical-eb8fc9118e871ac8027bc933c2bdfd5cb28684911b1e65f6319b6c61d38c87d6) |
| `name` | [name](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1d763b0275d454660641ea8e1e8066a00e8e3216f2ae61dba1f22e3cc31a1f84) |
| `namespace` | [namespace](data-sources--ike_phase2_profile--reference--group-001.md#canonical-389d385de8b7a12b82f595c07d0c27044607130d816712c4b6d657409506915b) |
| `use_default_keylifetime` | [use_default_keylifetime](data-sources--ike_phase2_profile--reference--group-001.md#canonical-7997c7063e5d1df4ba673bbcf7b0f78e80556739e3d4cd86406adc15230b46e9) |

<a id="canonical-3f12622a4923ae3b48cc2ba33bf4ee424e36222b038bd899b72e8c25568c074e"></a>

## Next pages — Property reference / 5ede625c02f0 / 13

- [dh_group_set](data-sources--ike_phase2_profile--reference--group-001.md#canonical-156d9a0f217a28938f835e169bc3a0b6270fe27506537f5a0e1bcfeb942bc0c7)
- [disable_pfs](data-sources--ike_phase2_profile--reference--group-001.md#canonical-0d0e01b46d38aa6e76991b7190d35a676cace75d45ac86a67748307ba1296d93)
- [ike_keylifetime_hours](data-sources--ike_phase2_profile--reference--group-001.md#canonical-29b6aab12573f1a303b3830d440e12a649eee2ecd2c6972002b28b2a67392249)
- [ike_keylifetime_minutes](data-sources--ike_phase2_profile--reference--group-001.md#canonical-d42d0d39e501e98365077d37bb79eb6b9d38dd374ae3637c32e607749ffd390a)
- [use_default_keylifetime](data-sources--ike_phase2_profile--reference--group-001.md#canonical-5d945860411de2759e392e5da68da15687751f7e57674e23ac9a5022825fea93)
- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)

<a id="canonical-156d9a0f217a28938f835e169bc3a0b6270fe27506537f5a0e1bcfeb942bc0c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-920a8332b9acfe9c0399867544795602039f2cdc434ade5400363d6678ce838f"></a>

## dh_group_set — dh_group_set / 6167462ba14d / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)
- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a)
- dh_group_set

<a id="canonical-714b8055d514dc40a60d0eb257245761c0fd75b51557bdfc48feaddf8b243b13"></a>

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

- [dh_group_set](data-sources--ike_phase2_profile--reference--group-001.md#canonical-714b8055d514dc40a60d0eb257245761c0fd75b51557bdfc48feaddf8b243b13)
- [disable_pfs](data-sources--ike_phase2_profile--reference--group-001.md#canonical-1ff975c5328084705c5ec883ac6fd2f96a499da4a37ad11a7f4ca5f475643d59)

Select alternatives according to the provider validators above.

<a id="canonical-01b37a69b43f4e83a41f694627f925276ba3a6f2d0ef678a260c73cf55091cd3"></a>

## Direct properties — dh_group_set / 6167462ba14d / 3

<a id="canonical-ae5c7c44ab9bd52d115582ac67e70a34c08109837ffa8f63f997fef2bd3b5e11"></a>

<a id="canonical-ddb38fd567c4d9178c08060db1af6be40a9f519050238372f8d24f826b4f880f"></a>

## dh_groups property — dh_group_set / 6167462ba14d / 4

Type: `["list", "string"]`. Computed.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile. Possible values are \`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`,
\`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`, \`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`,
\`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`. Defaults to \`DH\_GROUP\_DEFAULT\`.

Upstream description:

Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
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

<a id="canonical-826e05ffc6ff1501244ba17b73c0ab182e0a302fcba78c93a88a518f13cf0870"></a>

## Next pages — dh_group_set / 6167462ba14d / 5

- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a)
- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)

<a id="canonical-0d0e01b46d38aa6e76991b7190d35a676cace75d45ac86a67748307ba1296d93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5452903eb9cf5b42c3521ae1a3004fb737f0a0eb5f054ff7e06d1d3e023ca203"></a>

## disable_pfs — disable_pfs / c95570ec059d / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)
- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a)
- disable_pfs

<a id="canonical-1ff975c5328084705c5ec883ac6fd2f96a499da4a37ad11a7f4ca5f475643d59"></a>

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

<a id="canonical-09067f0272a338dca94e4a1fb3fa2b407bcdf15ac7a4bddedf081be1f060444d"></a>

## Direct properties — disable_pfs / c95570ec059d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d4b2908032755d2190ad4b307ec5831cf915ddf18e188c67f3c719c0386691ff"></a>

## Next pages — disable_pfs / c95570ec059d / 4

- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a)
- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)

<a id="canonical-29b6aab12573f1a303b3830d440e12a649eee2ecd2c6972002b28b2a67392249"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c2b0947d435974e5091175ae48a5d1f40db1af69fd5e603f265863fd151c8ad"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / ecf86584ec0a / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)
- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a)
- ike_keylifetime_hours

<a id="canonical-a6042d88be62d5fccff921280dbf74ef3a45e3c88a8cef475a378f99d86c4898"></a>

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

- [ike_keylifetime_hours](data-sources--ike_phase2_profile--reference--group-001.md#canonical-a6042d88be62d5fccff921280dbf74ef3a45e3c88a8cef475a378f99d86c4898)
- [ike_keylifetime_minutes](data-sources--ike_phase2_profile--reference--group-001.md#canonical-6d51cc2e95d0dfea81feadf596e3fc79b932dfc07415b942e8df0ac9160745d1)
- [use_default_keylifetime](data-sources--ike_phase2_profile--reference--group-001.md#canonical-7997c7063e5d1df4ba673bbcf7b0f78e80556739e3d4cd86406adc15230b46e9)

Select alternatives according to the provider validators above.

<a id="canonical-7775e47fd757e15234fd26c0f5e40d0f474650a06b9a89157038137f88a7a991"></a>

## Direct properties — ike_keylifetime_hours / ecf86584ec0a / 3

<a id="canonical-609befcb21b9756e6fa7574835971c5ecd31b1a5c97a9aeb811b0658f7984d95"></a>

<a id="canonical-c03139c07285e7006de3d6a48bb022dae02e3244132760b9966146ad8b69dddc"></a>

## duration property — ike_keylifetime_hours / ecf86584ec0a / 4

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

<a id="canonical-be715ef5f35903ecd627e0314987f82e9e09c46bcea3cba8387623a4193cb5ea"></a>

## Next pages — ike_keylifetime_hours / ecf86584ec0a / 5

- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a)
- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)

<a id="canonical-d42d0d39e501e98365077d37bb79eb6b9d38dd374ae3637c32e607749ffd390a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d3097e0334fdda75fb2e5e41ff21f671a89dfe27e2a93e3059d3fb4578cc112"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / 2cc1adadf30f / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)
- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a)
- ike_keylifetime_minutes

<a id="canonical-6d51cc2e95d0dfea81feadf596e3fc79b932dfc07415b942e8df0ac9160745d1"></a>

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

<a id="canonical-566ffed9494b8f8e071fe70dd711d643fa828c7575e5e606d92a39a4bf930a7d"></a>

## Direct properties — ike_keylifetime_minutes / 2cc1adadf30f / 3

<a id="canonical-2615bde362805f1a3251914282ca9ae1beed01e258ff9bbeacc685ad88732a33"></a>

<a id="canonical-bc43d2cd4da2c7d8ccf27d63f5219bfc48a0ab339e940060a10386cc804662b2"></a>

## duration property — ike_keylifetime_minutes / 2cc1adadf30f / 4

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

<a id="canonical-ece053a2e528621501e1a30a65759ebf9a265e6cacaa04b00bf0878426a73236"></a>

## Next pages — ike_keylifetime_minutes / 2cc1adadf30f / 5

- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a)
- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)

<a id="canonical-5d945860411de2759e392e5da68da15687751f7e57674e23ac9a5022825fea93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80d9142b1757ed6894aee74f84ede997807a53fa4c8bcde97d4629d3358e5933"></a>

## use_default_keylifetime — use_default_keylifetime / 5045e9fb200f / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)
- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a)
- use_default_keylifetime

<a id="canonical-7997c7063e5d1df4ba673bbcf7b0f78e80556739e3d4cd86406adc15230b46e9"></a>

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

<a id="canonical-40c15bc8ab04a2f429692130a5bc7af6854bb8fd7875d189926cfe8e9e959348"></a>

## Direct properties — use_default_keylifetime / 5045e9fb200f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b6a42ca2edf6bf654d6d7b133f1726261b8967ffd52dd99af2359efd82218542"></a>

## Next pages — use_default_keylifetime / 5045e9fb200f / 4

- [Property reference](data-sources--ike_phase2_profile--reference--group-001.md#canonical-69208af1f987cbde066b5daf25f7fe566428d0259b66669340f64476744b4c9a)
- [xcsh_ike_phase2_profile](../data-sources/ike_phase2_profile.md#canonical-d4f07d0519ebc7b4b0d60c702b18e26168602089d08803fd25362ad1cabf1c5c)
