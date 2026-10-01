---
page_title: "xcsh_ike_phase2_profile reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase2_profile reference."
---

# xcsh_ike_phase2_profile reference

<a id="canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f69a5dbaa586a819045c741fbb6fc1b95eba8cf172397db19e7318508504b420"></a>

## Property reference — Property reference / 3a3ffae02e0a / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)
- Property reference

<a id="canonical-a15d6139969c3b02d98c17a6c83d9943e6b1f8b3157b9ced233b710a4778fed7"></a>

## Direct properties — Property reference / 3a3ffae02e0a / 3

<a id="canonical-983b56dcbf3c75e4b81391d436870e60bfc96935d99695cd844f54f98e0c39fa"></a>

<a id="canonical-a5f5f1609ba90a55d7bf0892937bc9e1ba9580809a46e0844f375381ddc83d7d"></a>

## annotations property — Property reference / 3a3ffae02e0a / 4

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

<a id="canonical-3782675ec08cffbcab3c9648a171dadc07ece11c6d5b86e02d33a4cd2a9155bc"></a>

<a id="canonical-0ed044de64fc2b4abb13e817387e2169bbe49b72f13e19f1fd77e20d1a511267"></a>

## authentication_algos property — Property reference / 3a3ffae02e0a / 5

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

<a id="canonical-c4859c910077aa29d51c23932b2f3a4f39f5d3a208b8331c8ce60bee8c13a88b"></a>

<a id="canonical-3c242d766fc867e8329c5933f216a97f62634cea09fba5130f00618c62b5cfa7"></a>

## description property — Property reference / 3a3ffae02e0a / 6

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

- [dh_group_set](resources--ike_phase2_profile--reference--group-001.md#canonical-1f9976689c848af2b17ea49f0ee3b1cb669ff974e0b04898b320fec8b34675a2): complete subsection reference.

<a id="canonical-872868c1e4c99be1a2bc4fe1d42e78cc670782fcac4c19cfbd7fd024ffff7c04"></a>

<a id="canonical-ddf31bddd314d34c0a53b2962ac0dee5ba17d8906dc12c3fe69c4e010b6771d1"></a>

## disable property — Property reference / 3a3ffae02e0a / 7

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

- [disable_pfs](resources--ike_phase2_profile--reference--group-001.md#canonical-cd94de28078f70f974397f0451c1ce0365a115df703840be866610dc3e770747): complete subsection reference.

<a id="canonical-e39dbe5cae9294a1b226e41dc083168870dcf06c16751863bfa442d6b37c1c91"></a>

<a id="canonical-f463a3b8abe920904f36e0244f517bf5a697c70ab1e87550a14c130b8b94d2f7"></a>

## encryption_algos property — Property reference / 3a3ffae02e0a / 8

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

<a id="canonical-a578d20bf75e44e2cec026b5150dd24c6f4320a826b05455955616ba1cb79e8a"></a>

<a id="canonical-3e41a3b307fec0cbcddcf6bab95f5665c676955923de6b0acc68aae2c8197b34"></a>

## id property — Property reference / 3a3ffae02e0a / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](resources--ike_phase2_profile--reference--group-001.md#canonical-cecfa5bc3a3718fbb5c2accc4245dd089336e9155ddaeb5388f78566022183c7): complete subsection reference.

- [ike_keylifetime_minutes](resources--ike_phase2_profile--reference--group-001.md#canonical-c721e161cdb1d3f56ac708ce0ecad7c7bb82b379c406cab31868ec7b4b4fc988): complete subsection reference.

<a id="canonical-2ff38996fe0bad51bed5927dfdb5726d70e325e1c9877dd73a310735b07f5097"></a>

<a id="canonical-d1c1cb0787c3b70c12dc10dd8af149cbd1b8159491d3b491fd6d9a4a4c09a069"></a>

## labels property — Property reference / 3a3ffae02e0a / 10

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

<a id="canonical-dc6a1388cb7409b2d416a3c91c3673bf9767557e4c6ce49e087dd38b502cd52b"></a>

<a id="canonical-80cb1360b25c7aa66c2b49fdb942a1392698020c9a9997be99b258cc7dee01c8"></a>

## name property — Property reference / 3a3ffae02e0a / 11

Type: `"string"`. Required.

Name of the IKE Phase2 Profile. Must be unique within the namespace.

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

<a id="canonical-404f14fa8d6648010066739ee1bd92dee153b90b39c9ca367e11bb2534afe544"></a>

<a id="canonical-4f84d4aa1b1670be3e34a7c6771d753c901f7c252cdb535d8ff9710003586df0"></a>

## namespace property — Property reference / 3a3ffae02e0a / 12

Type: `"string"`. Required.

Namespace where the IKE Phase2 Profile is created.

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

- [timeouts](resources--ike_phase2_profile--reference--group-001.md#canonical-d1fbddda94ef2438c7e9316654b05d9f63b7ebd7fff18764b5b1ceeb03a5babf): complete subsection reference.

- [use_default_keylifetime](resources--ike_phase2_profile--reference--group-001.md#canonical-f56139b5b0b360f82cc35d43b0dd7ed13c78bb5356c616967b56a15b40f69a14): complete subsection reference.

<a id="canonical-cdc580000575662ab0a332c609572d8eaa02045605e8a3d566065bd7a152d0a0"></a>

## All schema paths — Property reference / 3a3ffae02e0a / 13

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--ike_phase2_profile--reference--group-001.md#canonical-983b56dcbf3c75e4b81391d436870e60bfc96935d99695cd844f54f98e0c39fa) |
| `authentication_algos` | [authentication_algos](resources--ike_phase2_profile--reference--group-001.md#canonical-3782675ec08cffbcab3c9648a171dadc07ece11c6d5b86e02d33a4cd2a9155bc) |
| `description` | [description](resources--ike_phase2_profile--reference--group-001.md#canonical-c4859c910077aa29d51c23932b2f3a4f39f5d3a208b8331c8ce60bee8c13a88b) |
| `dh_group_set` | [dh_group_set](resources--ike_phase2_profile--reference--group-001.md#canonical-82da35aea21fe11d70e682b0e37d48daa6f8713543fdfd83e6f90c42ecd69def) |
| `dh_group_set.dh_groups` | [dh_group_set.dh_groups](resources--ike_phase2_profile--reference--group-001.md#canonical-1d52b44e5d149a92cb1676985bda67d516abd77a6aa5b2360ab9769278b22c32) |
| `disable` | [disable](resources--ike_phase2_profile--reference--group-001.md#canonical-872868c1e4c99be1a2bc4fe1d42e78cc670782fcac4c19cfbd7fd024ffff7c04) |
| `disable_pfs` | [disable_pfs](resources--ike_phase2_profile--reference--group-001.md#canonical-5ac841995d7bf569eaa9749c49398ee557b0a3bd5501174e0d4f88b442c62a62) |
| `encryption_algos` | [encryption_algos](resources--ike_phase2_profile--reference--group-001.md#canonical-e39dbe5cae9294a1b226e41dc083168870dcf06c16751863bfa442d6b37c1c91) |
| `id` | [id](resources--ike_phase2_profile--reference--group-001.md#canonical-a578d20bf75e44e2cec026b5150dd24c6f4320a826b05455955616ba1cb79e8a) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](resources--ike_phase2_profile--reference--group-001.md#canonical-4afbb0123efd6c2048c65c426434b6bf9d9b83ce765c23bd05a367b89a17b5ee) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](resources--ike_phase2_profile--reference--group-001.md#canonical-0589977d3b9761a55089847c9a82827f28eeb7b4c2258237f42de0bdaa3dc54e) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](resources--ike_phase2_profile--reference--group-001.md#canonical-c14cfeb4589f91e2c60af96a3293c9ba60a991e6385bfbe4e1c4a2b7f9ac5330) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](resources--ike_phase2_profile--reference--group-001.md#canonical-f0d057e036b9345a85644cddedd1ae22513a8cf7aa82a726123fc245aca37e26) |
| `labels` | [labels](resources--ike_phase2_profile--reference--group-001.md#canonical-2ff38996fe0bad51bed5927dfdb5726d70e325e1c9877dd73a310735b07f5097) |
| `name` | [name](resources--ike_phase2_profile--reference--group-001.md#canonical-dc6a1388cb7409b2d416a3c91c3673bf9767557e4c6ce49e087dd38b502cd52b) |
| `namespace` | [namespace](resources--ike_phase2_profile--reference--group-001.md#canonical-404f14fa8d6648010066739ee1bd92dee153b90b39c9ca367e11bb2534afe544) |
| `timeouts` | [timeouts](resources--ike_phase2_profile--reference--group-001.md#canonical-6d0650f9d6d9b40970a0e4f943f67b01703209cf50e08dc495335b5504e817da) |
| `timeouts.create` | [timeouts.create](resources--ike_phase2_profile--reference--group-001.md#canonical-d1f407ccbf71e3a8d8106b01b0fda9f2ea6622bbb96904aa867f27a537a9f31c) |
| `timeouts.delete` | [timeouts.delete](resources--ike_phase2_profile--reference--group-001.md#canonical-0bbda21b0a3ddabb0b4874c87916ea535fc1e31e04b4289cd01f4c8aa20f03cd) |
| `timeouts.read` | [timeouts.read](resources--ike_phase2_profile--reference--group-001.md#canonical-905cb36f1a27049584533b5c064e642dfde51c1b82f68ff003207f64b5661f93) |
| `timeouts.update` | [timeouts.update](resources--ike_phase2_profile--reference--group-001.md#canonical-ae4da38766c72cf62e5dac064fa560c1a3ce9047919d036dfafc944ec9ac89d6) |
| `use_default_keylifetime` | [use_default_keylifetime](resources--ike_phase2_profile--reference--group-001.md#canonical-5e85e5ff7c202a7bfee59865cf3ba4c8eff721b70d9edc5cbab2573176a5501f) |

<a id="canonical-37ee93b619cf21fa25b273a51e42a0e69c7df6daf5c045de86d2f40ee141aa6f"></a>

## Next pages — Property reference / 3a3ffae02e0a / 14

- [dh_group_set](resources--ike_phase2_profile--reference--group-001.md#canonical-1f9976689c848af2b17ea49f0ee3b1cb669ff974e0b04898b320fec8b34675a2)
- [disable_pfs](resources--ike_phase2_profile--reference--group-001.md#canonical-cd94de28078f70f974397f0451c1ce0365a115df703840be866610dc3e770747)
- [ike_keylifetime_hours](resources--ike_phase2_profile--reference--group-001.md#canonical-cecfa5bc3a3718fbb5c2accc4245dd089336e9155ddaeb5388f78566022183c7)
- [ike_keylifetime_minutes](resources--ike_phase2_profile--reference--group-001.md#canonical-c721e161cdb1d3f56ac708ce0ecad7c7bb82b379c406cab31868ec7b4b4fc988)
- [timeouts](resources--ike_phase2_profile--reference--group-001.md#canonical-d1fbddda94ef2438c7e9316654b05d9f63b7ebd7fff18764b5b1ceeb03a5babf)
- [use_default_keylifetime](resources--ike_phase2_profile--reference--group-001.md#canonical-f56139b5b0b360f82cc35d43b0dd7ed13c78bb5356c616967b56a15b40f69a14)
- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)

<a id="canonical-1f9976689c848af2b17ea49f0ee3b1cb669ff974e0b04898b320fec8b34675a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81078d199a5dcca670b3090b2774b8a49d6d46aebd5878b860c6d2945087abe5"></a>

## dh_group_set — dh_group_set / 3f2864cb0a03 / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- dh_group_set

<a id="canonical-82da35aea21fe11d70e682b0e37d48daa6f8713543fdfd83e6f90c42ecd69def"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dh\_group\_set, disable\_pfs; Default: disable\_pfs\] Choose the acceptable Diffie
Hellman(DH) Group or Groups that you are willing to accept as part of this profile.

Upstream description:

Choose the acceptable Diffie Hellman(DH) Group or Groups that you are willing to accept as part of
this profile.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dh_groups")}
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

- [dh_group_set](resources--ike_phase2_profile--reference--group-001.md#canonical-82da35aea21fe11d70e682b0e37d48daa6f8713543fdfd83e6f90c42ecd69def)
- [disable_pfs](resources--ike_phase2_profile--reference--group-001.md#canonical-5ac841995d7bf569eaa9749c49398ee557b0a3bd5501174e0d4f88b442c62a62)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dh_group_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-3fed04499e34ea1aa68e2e15862780a95d3f59e5f2e9d49063d80e030f9629d4"></a>

## Direct properties — dh_group_set / 3f2864cb0a03 / 3

<a id="canonical-1d52b44e5d149a92cb1676985bda67d516abd77a6aa5b2360ab9769278b22c32"></a>

<a id="canonical-542f39f9cf0f2f8d1f2c5b0a36dba8911c37093532ff0128dac8eb605bb37e0d"></a>

## dh_groups property — dh_group_set / 3f2864cb0a03 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-76b807db5244c4632866e6ab631e557dc15f52c8a6082f81af0d755519db32a3"></a>

## Next pages — dh_group_set / 3f2864cb0a03 / 5

- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)

<a id="canonical-cd94de28078f70f974397f0451c1ce0365a115df703840be866610dc3e770747"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23b973f0baf811fa39b495f3e44e47dcd029cf32f580db45b396c481fdb880a7"></a>

## disable_pfs — disable_pfs / fea5dafa1e5d / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- disable_pfs

<a id="canonical-5ac841995d7bf569eaa9749c49398ee557b0a3bd5501174e0d4f88b442c62a62"></a>

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

<a id="canonical-65bc8879e0247f2e7d671df12feac56daf485f1e4750e7c886db382a42b64139"></a>

## Direct properties — disable_pfs / fea5dafa1e5d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b7883c7ed2e39a10db0e958e5fcb7b0e9e17c51639ddaaa443ef477abcc7fe9b"></a>

## Next pages — disable_pfs / fea5dafa1e5d / 4

- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)

<a id="canonical-cecfa5bc3a3718fbb5c2accc4245dd089336e9155ddaeb5388f78566022183c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08c43e9b1de24fd8098da6510f6909e66b92c5a56090fa4052d87630852a58c9"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / 419883dbbc1f / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- ike_keylifetime_hours

<a id="canonical-4afbb0123efd6c2048c65c426434b6bf9d9b83ce765c23bd05a367b89a17b5ee"></a>

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

- [ike_keylifetime_hours](resources--ike_phase2_profile--reference--group-001.md#canonical-4afbb0123efd6c2048c65c426434b6bf9d9b83ce765c23bd05a367b89a17b5ee)
- [ike_keylifetime_minutes](resources--ike_phase2_profile--reference--group-001.md#canonical-c14cfeb4589f91e2c60af96a3293c9ba60a991e6385bfbe4e1c4a2b7f9ac5330)
- [use_default_keylifetime](resources--ike_phase2_profile--reference--group-001.md#canonical-5e85e5ff7c202a7bfee59865cf3ba4c8eff721b70d9edc5cbab2573176a5501f)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ike_keylifetime_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-5fc0ef7f46bca3710cd12b10c3fdd08ec1bb24d98d7943ef4a12765502c6728f"></a>

## Direct properties — ike_keylifetime_hours / 419883dbbc1f / 3

<a id="canonical-0589977d3b9761a55089847c9a82827f28eeb7b4c2258237f42de0bdaa3dc54e"></a>

<a id="canonical-6ddfa9ff4ffaf5cdcbff865a1bcaa39313ebca3bfcbbaf74d1a40c6ccce9508d"></a>

## duration property — ike_keylifetime_hours / 419883dbbc1f / 4

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

<a id="canonical-542d08cd162c15893b3cba594bd96c4890d2f4832aad03efda72cc249f8454f3"></a>

## Next pages — ike_keylifetime_hours / 419883dbbc1f / 5

- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)

<a id="canonical-c721e161cdb1d3f56ac708ce0ecad7c7bb82b379c406cab31868ec7b4b4fc988"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1e15f2628eb1bc1ba183343b9d14f50ed582eaf45cd8a856e1f96b2f575f021"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / c06b71258421 / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- ike_keylifetime_minutes

<a id="canonical-c14cfeb4589f91e2c60af96a3293c9ba60a991e6385bfbe4e1c4a2b7f9ac5330"></a>

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

<a id="canonical-27d9416ac6ea42a154cb8d8d17af114acb06b835a18d6560a44f16c845b0533d"></a>

## Direct properties — ike_keylifetime_minutes / c06b71258421 / 3

<a id="canonical-f0d057e036b9345a85644cddedd1ae22513a8cf7aa82a726123fc245aca37e26"></a>

<a id="canonical-d0f294061fc187b46010b1f131a0c1af02f4315298daf289dfdf3ff836219d9d"></a>

## duration property — ike_keylifetime_minutes / c06b71258421 / 4

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

<a id="canonical-173d578f72121ba5abe45b6e9d5ae46985169871195cfde811de528b20498216"></a>

## Next pages — ike_keylifetime_minutes / c06b71258421 / 5

- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)

<a id="canonical-d1fbddda94ef2438c7e9316654b05d9f63b7ebd7fff18764b5b1ceeb03a5babf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa636053b76b4923da95facc52c6ee2ad6ef4474d49d0425ec626a5e7956939a"></a>

## timeouts — timeouts / 92d60b8252f8 / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- timeouts

<a id="canonical-6d0650f9d6d9b40970a0e4f943f67b01703209cf50e08dc495335b5504e817da"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-18b5acd2fee51be0a506a031c51a61de85b78d0c104979bb056cffcf67407179"></a>

## Direct properties — timeouts / 92d60b8252f8 / 3

<a id="canonical-d1f407ccbf71e3a8d8106b01b0fda9f2ea6622bbb96904aa867f27a537a9f31c"></a>

<a id="canonical-e40cd6c45f497b1a219412aa79016b40de80a85c99be24009d70d1d80b8f34bf"></a>

## create property — timeouts / 92d60b8252f8 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0bbda21b0a3ddabb0b4874c87916ea535fc1e31e04b4289cd01f4c8aa20f03cd"></a>

<a id="canonical-bee86adf2b2c7b8955c886271d68e41fd7ea026320d6e568281c05ba801ffd14"></a>

## delete property — timeouts / 92d60b8252f8 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-905cb36f1a27049584533b5c064e642dfde51c1b82f68ff003207f64b5661f93"></a>

<a id="canonical-c0cf7c18302d2db914d3d6d2409cfe4388170c5a1581da888a679a0909155f93"></a>

## read property — timeouts / 92d60b8252f8 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-ae4da38766c72cf62e5dac064fa560c1a3ce9047919d036dfafc944ec9ac89d6"></a>

<a id="canonical-d5b7f26a864371553796fe5cf5e76c7975744497592a37df23f31833d41e3bdc"></a>

## update property — timeouts / 92d60b8252f8 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-d598247d753c9b576199ac7c07aa5d21e9609c9314834ce8028a576b2c534734"></a>

## Next pages — timeouts / 92d60b8252f8 / 8

- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)

<a id="canonical-f56139b5b0b360f82cc35d43b0dd7ed13c78bb5356c616967b56a15b40f69a14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23fc55cb39ddebad418bb60c6781c9b18fab8c5b45daa46cd3ec5ad56f112d40"></a>

## use_default_keylifetime — use_default_keylifetime / b4fbef158ccb / 2

Breadcrumbs:

- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)
- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- use_default_keylifetime

<a id="canonical-5e85e5ff7c202a7bfee59865cf3ba4c8eff721b70d9edc5cbab2573176a5501f"></a>

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

<a id="canonical-849bcc1ea583ff2eda3dae1fe7fee620d03f83382c13e6edd550e5be63eeee08"></a>

## Direct properties — use_default_keylifetime / b4fbef158ccb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4b0cfc0dcb8c349fb774f1a870a46e6b91850c374903d81cb957edc999e41524"></a>

## Next pages — use_default_keylifetime / b4fbef158ccb / 4

- [Property reference](resources--ike_phase2_profile--reference--group-001.md#canonical-880ac98477290434e8fa578e0ae87dda91ef4fd852b03e4243db262ef0076126)
- [xcsh_ike_phase2_profile](../resources/ike_phase2_profile.md#canonical-61d2d3b324a5d2d7c402d1c590f903429b98898ccc9ec864fa8d0a29eca618ad)
