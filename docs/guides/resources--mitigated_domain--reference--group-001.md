---
page_title: "xcsh_mitigated_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_mitigated_domain reference."
---

# xcsh_mitigated_domain reference

<a id="canonical-d8a7d36f91ae444348c9b288663ecdfe54e263784cd66b472170a42f06e73224"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e18d95bf9f5660cf73a0a6d5fe8b0122accbf9a40deb5e8d8cfa67116cc542ab"></a>

## Property reference — Property reference / 22e2bef5cfdf / 2

Breadcrumbs:

- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-8dbccf2375d1fbc444e0bd4617f6f345af352b51e200bf181ac86bffcdfce192)
- Property reference

<a id="canonical-f4032f3d103bc4d1cfeb3f86cb13074e90d18146f927b0f928af76394d3f2df5"></a>

## Direct properties — Property reference / 22e2bef5cfdf / 3

<a id="canonical-25b16f4c4b7655b7a954f3b1c3cff6d8709411dce89a4839a0a74fac2f2eb9d0"></a>

<a id="canonical-b4ea7a9a1fec30c8d13a5f5eb8502ea28a04c6e14a74de067a8e77b40f4b65ea"></a>

## annotations property — Property reference / 22e2bef5cfdf / 4

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

<a id="canonical-79ecc2463c18018edba23e3ac0a7d64b6cd36c4d38fc166d639cb21161d21082"></a>

<a id="canonical-1d0aa68ecb8b74520f219a07d701512fe2a17335d665dc303775e6addd0be42c"></a>

## description property — Property reference / 22e2bef5cfdf / 5

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

<a id="canonical-659e016746d62691ebce632fb121f246964360149d12f74e8b37de3b0ac4ac28"></a>

<a id="canonical-dd44665e3f6db9d5c8d480ffc148f00558f549f3b4c510e0f66bd5c2f476d703"></a>

## disable property — Property reference / 22e2bef5cfdf / 6

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

<a id="canonical-487ad32bbf68a0458847f8d0a8ac86d5e27005d037399ffb6f38ab2d92dc6b62"></a>

<a id="canonical-d5f7eb7e3d41eedeb0970b76689bac3374f72dc6be3e5cf25d160ad312b6106f"></a>

## id property — Property reference / 22e2bef5cfdf / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-d0414deb73fb590323fc3ebb8cb7dc4286e13d0fa8d1ffc08bbb19d8af2f7a3c"></a>

<a id="canonical-8f91580415a8af6ca12081a2eb925631fd54882be8ce70252eb82c5429daab1e"></a>

## labels property — Property reference / 22e2bef5cfdf / 8

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

<a id="canonical-31b1fae996db088c7512cf46deb1e23595e6564444bdd29f2b60d68962856c29"></a>

<a id="canonical-f125346570a23a341f603ec500774db43dfbbfc1d84eae6b1fa1d771103a91f3"></a>

## mitigated_domain property — Property reference / 22e2bef5cfdf / 9

Type: `"string"`. Required.

Enter root domain or domain to be entered to mitigated list below. Domains can be entered only one
at a time. In case of conflicting entries, the domain entry takes precedence over the root domain
entry.

Upstream description:

Enter root domain or domain to be entered to mitigated list below. Domains can be entered only one
at a time. In case of conflicting entries, the domain entry takes precedence over the root domain
entry. Example: if you are adding Client-Side Defense JS on checkout.example.com, you should enter
example.com here.

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
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-ad5128b31f7fc6374eb9ea608520f1de33980b2080637738dceb978598ee64c6"></a>

<a id="canonical-7a7a72da5025c8457f1336ce6c9aed542200e8a5f29fdf86672afd4d32f41709"></a>

## name property — Property reference / 22e2bef5cfdf / 10

Type: `"string"`. Required.

Name of the Mitigated Domain. Must be unique within the namespace.

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

<a id="canonical-12edefcad219049087451cedf45afa60631f1d341f64832f7deeae316aa8355a"></a>

<a id="canonical-a1040fe14ea2b4e7ad2bdfed6ede1f1875c42c4f98ea8bc78902b2ddddffde83"></a>

## namespace property — Property reference / 22e2bef5cfdf / 11

Type: `"string"`. Required.

Namespace where the Mitigated Domain is created.

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

- [timeouts](resources--mitigated_domain--reference--group-001.md#canonical-a05e4c39ab6cb47fb6022115dbbde646de1d7829e2a1c943cbb0854981c8df91): complete subsection reference.

<a id="canonical-d972e5e9f217ce8f9d65b0e52e4e9df895e20269c9c435865679c9fc6d1cda72"></a>

## All schema paths — Property reference / 22e2bef5cfdf / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--mitigated_domain--reference--group-001.md#canonical-25b16f4c4b7655b7a954f3b1c3cff6d8709411dce89a4839a0a74fac2f2eb9d0) |
| `description` | [description](resources--mitigated_domain--reference--group-001.md#canonical-79ecc2463c18018edba23e3ac0a7d64b6cd36c4d38fc166d639cb21161d21082) |
| `disable` | [disable](resources--mitigated_domain--reference--group-001.md#canonical-659e016746d62691ebce632fb121f246964360149d12f74e8b37de3b0ac4ac28) |
| `id` | [id](resources--mitigated_domain--reference--group-001.md#canonical-487ad32bbf68a0458847f8d0a8ac86d5e27005d037399ffb6f38ab2d92dc6b62) |
| `labels` | [labels](resources--mitigated_domain--reference--group-001.md#canonical-d0414deb73fb590323fc3ebb8cb7dc4286e13d0fa8d1ffc08bbb19d8af2f7a3c) |
| `mitigated_domain` | [mitigated_domain](resources--mitigated_domain--reference--group-001.md#canonical-31b1fae996db088c7512cf46deb1e23595e6564444bdd29f2b60d68962856c29) |
| `name` | [name](resources--mitigated_domain--reference--group-001.md#canonical-ad5128b31f7fc6374eb9ea608520f1de33980b2080637738dceb978598ee64c6) |
| `namespace` | [namespace](resources--mitigated_domain--reference--group-001.md#canonical-12edefcad219049087451cedf45afa60631f1d341f64832f7deeae316aa8355a) |
| `timeouts` | [timeouts](resources--mitigated_domain--reference--group-001.md#canonical-7c566af09a173a752d0843c51a0436085dba60d60d7b07adc1dfb585ad82acef) |
| `timeouts.create` | [timeouts.create](resources--mitigated_domain--reference--group-001.md#canonical-c7ba91daae8bb58a42d78c33106ad17278f115b6704a7f5a4af3eda6cee1d5c1) |
| `timeouts.delete` | [timeouts.delete](resources--mitigated_domain--reference--group-001.md#canonical-308139bf9e9ceb2b911220b8b0a3334edf803c1ab4dd078697f955861193fe62) |
| `timeouts.read` | [timeouts.read](resources--mitigated_domain--reference--group-001.md#canonical-b1a2e98ddc43ad8c841c7692b595aa7cdbc3dbf212d84396ea5f4da17158af54) |
| `timeouts.update` | [timeouts.update](resources--mitigated_domain--reference--group-001.md#canonical-3da8c47d4e021090c960ffd677308c44802789791b0c61998d9722739b264741) |

<a id="canonical-be29889e183caff90f842f45ebd070017dc271473dca0075c6b965e18d7960a7"></a>

## Next pages — Property reference / 22e2bef5cfdf / 13

- [timeouts](resources--mitigated_domain--reference--group-001.md#canonical-a05e4c39ab6cb47fb6022115dbbde646de1d7829e2a1c943cbb0854981c8df91)
- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-8dbccf2375d1fbc444e0bd4617f6f345af352b51e200bf181ac86bffcdfce192)

<a id="canonical-a05e4c39ab6cb47fb6022115dbbde646de1d7829e2a1c943cbb0854981c8df91"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2996693fa5bd95f2eb9fb0193c1207bbeb35edd1359adffab6bd8a7ad46d22ca"></a>

## timeouts — timeouts / bbf45f7a81fc / 2

Breadcrumbs:

- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-8dbccf2375d1fbc444e0bd4617f6f345af352b51e200bf181ac86bffcdfce192)
- [Property reference](resources--mitigated_domain--reference--group-001.md#canonical-d8a7d36f91ae444348c9b288663ecdfe54e263784cd66b472170a42f06e73224)
- timeouts

<a id="canonical-7c566af09a173a752d0843c51a0436085dba60d60d7b07adc1dfb585ad82acef"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-cbc0af2b28eab5ca0cc231af288fb8d3eb01075e543657b4251886f64f9f74c9"></a>

## Direct properties — timeouts / bbf45f7a81fc / 3

<a id="canonical-c7ba91daae8bb58a42d78c33106ad17278f115b6704a7f5a4af3eda6cee1d5c1"></a>

<a id="canonical-9b0f0730b81ee52683acedd6faa34c2c8ad7e9cfda0ada0be551ec58e5479936"></a>

## create property — timeouts / bbf45f7a81fc / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-308139bf9e9ceb2b911220b8b0a3334edf803c1ab4dd078697f955861193fe62"></a>

<a id="canonical-7a5dda0b7c6715e0bf6a4474f3aabf5226b0e7df641cb510383b02fd982490e1"></a>

## delete property — timeouts / bbf45f7a81fc / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-b1a2e98ddc43ad8c841c7692b595aa7cdbc3dbf212d84396ea5f4da17158af54"></a>

<a id="canonical-d403b1794267fc2aa7ce1c8b78ff90543f5663331889a5279e41d434fdfb8293"></a>

## read property — timeouts / bbf45f7a81fc / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3da8c47d4e021090c960ffd677308c44802789791b0c61998d9722739b264741"></a>

<a id="canonical-08ae2381e5148d2d559e5b687ca913aa4b194ca4c8b97677ccc668d78e350db8"></a>

## update property — timeouts / bbf45f7a81fc / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1b9d5584d70a3b6f514e6d9814ca4a07d26a6d6f70ca48f5c647b46a24bd6873"></a>

## Next pages — timeouts / bbf45f7a81fc / 8

- [Property reference](resources--mitigated_domain--reference--group-001.md#canonical-d8a7d36f91ae444348c9b288663ecdfe54e263784cd66b472170a42f06e73224)
- [xcsh_mitigated_domain](../resources/mitigated_domain.md#canonical-8dbccf2375d1fbc444e0bd4617f6f345af352b51e200bf181ac86bffcdfce192)
