---
page_title: "xcsh_ip_prefix_set reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set reference."
---

# xcsh_ip_prefix_set reference

<a id="canonical-e43df77cb707b03d9ea017a3256270c1386b0044a471db509c179d9d4a2c87e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd68d7d1abaf58c5747ffd8f39446a67d208f9af021d6e38aef0788b531dab0e"></a>

## Property reference — Property reference / cca514a61cc6 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-5a49cde9dc0a3a39ff028cd9a8304b00e480fb0bbedc342df01389fd1e3f30ae)
- Property reference

<a id="canonical-af754f148fcf57ff08423c6d0a1556aa790cdbe94c36d6051da476470e156f40"></a>

## Direct properties — Property reference / cca514a61cc6 / 3

<a id="canonical-d515d7f410ea5445d50169050cc971ef209f4fe048c91f207fd3c58a594b9705"></a>

<a id="canonical-f24a08f87b3c0862b5905cce324f40dc56765ff77ae8f7b52ebec30b5fe405c0"></a>

## annotations property — Property reference / cca514a61cc6 / 4

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

<a id="canonical-533865926a4c2d7c4a07a791c5dd8c05944ef1a202f9a593e673354e9a7bb926"></a>

<a id="canonical-653d3b8fa01978e0bf0f95c7b18870db5c9540bfeb33e8c80e0f627f96d3626b"></a>

## description property — Property reference / cca514a61cc6 / 5

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

<a id="canonical-28287bbacd9291b15e12747cc6646b0b47d458a206b9f4b5ad82bbce01f27bb2"></a>

<a id="canonical-1818b35ea154e4cee10cff23eb4e442d385707dc5de1cd357a16e58ae0c25b1d"></a>

## disable property — Property reference / cca514a61cc6 / 6

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

<a id="canonical-a888c237a1de6964fbe8d8c437343540417d489cd95f4a974bf1aee87789a5ac"></a>

<a id="canonical-e65a73e09186061c3db392720f6fdf410d1ccab94dbfdbd0039b401050a98541"></a>

## id property — Property reference / cca514a61cc6 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipv4_prefixes](resources--ip_prefix_set--reference--group-001.md#canonical-a0ff8a09c9c1618fac3601bd9431ff70f8fa90c94883b7cd21298ab0961f5a2c): complete subsection reference.

<a id="canonical-1011a4954b164082a6040792d27829a01f7e31fec8ae36627d37156025ba34ea"></a>

<a id="canonical-025bce7e3741a81b76c6e51f2a0017f6b243d9b660a2588260f33e1db3145564"></a>

## labels property — Property reference / cca514a61cc6 / 8

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

<a id="canonical-46a9320983f2a1038d31a20904fb23b64f32fe94222d35b17168c8070eab329b"></a>

<a id="canonical-19ba0cf780269b20b64e13129c8aaada9a514c9dcbb1860c2863b300dda5c7e2"></a>

## name property — Property reference / cca514a61cc6 / 9

Type: `"string"`. Required.

Name of the IP Prefix Set. Must be unique within the namespace.

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

<a id="canonical-5a4fcd4fc6309027f4b2bee089d62fe235e2683c7d2693f2ad2c7622f27ee79a"></a>

<a id="canonical-240812b0fb4ddf03cc8547eb8a827a9fddda77a923f74a0fcc8cbc7b1b0d40e7"></a>

## namespace property — Property reference / cca514a61cc6 / 10

Type: `"string"`. Required.

Namespace where the IP Prefix Set is created.

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

- [timeouts](resources--ip_prefix_set--reference--group-001.md#canonical-b9858b0cc1ba803c57763065b0503566d521f5cc21ce18822348708d535f0a27): complete subsection reference.

<a id="canonical-617e08f142e5726685464ec50e14389468b22cdacd8519dfcce7a9435051a4b9"></a>

## All schema paths — Property reference / cca514a61cc6 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--ip_prefix_set--reference--group-001.md#canonical-d515d7f410ea5445d50169050cc971ef209f4fe048c91f207fd3c58a594b9705) |
| `description` | [description](resources--ip_prefix_set--reference--group-001.md#canonical-533865926a4c2d7c4a07a791c5dd8c05944ef1a202f9a593e673354e9a7bb926) |
| `disable` | [disable](resources--ip_prefix_set--reference--group-001.md#canonical-28287bbacd9291b15e12747cc6646b0b47d458a206b9f4b5ad82bbce01f27bb2) |
| `id` | [id](resources--ip_prefix_set--reference--group-001.md#canonical-a888c237a1de6964fbe8d8c437343540417d489cd95f4a974bf1aee87789a5ac) |
| `ipv4_prefixes` | [ipv4_prefixes](resources--ip_prefix_set--reference--group-001.md#canonical-1e6890760f2a0a239cb4fc96301cfc5693bb63fd4de348bc592e2a0f73485434) |
| `ipv4_prefixes.description_spec` | [ipv4_prefixes.description_spec](resources--ip_prefix_set--reference--group-001.md#canonical-d663e27cdbb420f33dddc06555289316b1cc89b9467dc9359f291d81a5a314dd) |
| `ipv4_prefixes.ipv4_prefix` | [ipv4_prefixes.ipv4_prefix](resources--ip_prefix_set--reference--group-001.md#canonical-615c17912cc78d7d838d0387b96e0d449503119acf3fc84c7fb13f10e548bfe6) |
| `labels` | [labels](resources--ip_prefix_set--reference--group-001.md#canonical-1011a4954b164082a6040792d27829a01f7e31fec8ae36627d37156025ba34ea) |
| `name` | [name](resources--ip_prefix_set--reference--group-001.md#canonical-46a9320983f2a1038d31a20904fb23b64f32fe94222d35b17168c8070eab329b) |
| `namespace` | [namespace](resources--ip_prefix_set--reference--group-001.md#canonical-5a4fcd4fc6309027f4b2bee089d62fe235e2683c7d2693f2ad2c7622f27ee79a) |
| `timeouts` | [timeouts](resources--ip_prefix_set--reference--group-001.md#canonical-33c14b1ed33a4a2ce601c1d28fe1dd5604477ded121dec11b70a22533a4cef08) |
| `timeouts.create` | [timeouts.create](resources--ip_prefix_set--reference--group-001.md#canonical-0555e66ace5ad2174ddd84492e988e4d021dee054d3c777d553f2895afdd56c4) |
| `timeouts.delete` | [timeouts.delete](resources--ip_prefix_set--reference--group-001.md#canonical-14b784918a19683e6c39d5b7698fc65e521266dce9b11223869a6e206fa902b0) |
| `timeouts.read` | [timeouts.read](resources--ip_prefix_set--reference--group-001.md#canonical-3da9455b984e6021db02150c6662c6a49e2115517c6a81462751ad1095cc2605) |
| `timeouts.update` | [timeouts.update](resources--ip_prefix_set--reference--group-001.md#canonical-6c61a92a777c9d12d6029378116f0ae47c814aa04dafe9cc7ce42f2136282d13) |

<a id="canonical-3e50e65edf1eef9ab248c0b6931443da11c9e1aa20a8022b88a51ebd2b5c3900"></a>

## Next pages — Property reference / cca514a61cc6 / 12

- [ipv4_prefixes](resources--ip_prefix_set--reference--group-001.md#canonical-a0ff8a09c9c1618fac3601bd9431ff70f8fa90c94883b7cd21298ab0961f5a2c)
- [timeouts](resources--ip_prefix_set--reference--group-001.md#canonical-b9858b0cc1ba803c57763065b0503566d521f5cc21ce18822348708d535f0a27)
- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-5a49cde9dc0a3a39ff028cd9a8304b00e480fb0bbedc342df01389fd1e3f30ae)

<a id="canonical-a0ff8a09c9c1618fac3601bd9431ff70f8fa90c94883b7cd21298ab0961f5a2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88476914eb592dcef79090e884590e0c1fcb1ac026c0d3c5df392f2cadc9cda0"></a>

## ipv4_prefixes — ipv4_prefixes / e0fddc043128 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-5a49cde9dc0a3a39ff028cd9a8304b00e480fb0bbedc342df01389fd1e3f30ae)
- [Property reference](resources--ip_prefix_set--reference--group-001.md#canonical-e43df77cb707b03d9ea017a3256270c1386b0044a471db509c179d9d4a2c87e9)
- ipv4_prefixes

<a id="canonical-1e6890760f2a0a239cb4fc96301cfc5693bb63fd4de348bc592e2a0f73485434"></a>

Type: `"object"`. list nested block, Optional.

IPv4 Prefixes. List of IPv4 prefixes with description.

Upstream description:

List of IPv4 prefixes with description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ipv4_prefix")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ipv4_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-755defe9090a694e530e7df11fe0a2a3b8e39d72af564b9d8140e27fe462f2c0"></a>

## Direct properties — ipv4_prefixes / e0fddc043128 / 3

<a id="canonical-d663e27cdbb420f33dddc06555289316b1cc89b9467dc9359f291d81a5a314dd"></a>

<a id="canonical-19fe67ac1e5d3ed3b448f10d20f6b4acc3e5babd2a55928f3812870df1fb0c1d"></a>

## description_spec property — ipv4_prefixes / e0fddc043128 / 4

Type: `"string"`. Optional.

Description. Human-readable description text

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="canonical-615c17912cc78d7d838d0387b96e0d449503119acf3fc84c7fb13f10e548bfe6"></a>

<a id="canonical-77e390ec4338e5a4c4e5ec73b0a2595fb026a2f7909a93244e673b9089b708c4"></a>

## ipv4_prefix property — ipv4_prefixes / e0fddc043128 / 5

Type: `"string"`. Optional.

IPv4 Prefix. IP address configuration

Upstream description:

IP address configuration

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-8e2520f5209472b061790f6fa303bb22ee812ed9971f06b771d04f679e15c8eb"></a>

## Next pages — ipv4_prefixes / e0fddc043128 / 6

- [Property reference](resources--ip_prefix_set--reference--group-001.md#canonical-e43df77cb707b03d9ea017a3256270c1386b0044a471db509c179d9d4a2c87e9)
- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-5a49cde9dc0a3a39ff028cd9a8304b00e480fb0bbedc342df01389fd1e3f30ae)

<a id="canonical-b9858b0cc1ba803c57763065b0503566d521f5cc21ce18822348708d535f0a27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c02cdd502495351fa1d1f9fe9fc2c4293183ce44d083d194ae7de1edc8c3b075"></a>

## timeouts — timeouts / 6689c1e5b091 / 2

Breadcrumbs:

- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-5a49cde9dc0a3a39ff028cd9a8304b00e480fb0bbedc342df01389fd1e3f30ae)
- [Property reference](resources--ip_prefix_set--reference--group-001.md#canonical-e43df77cb707b03d9ea017a3256270c1386b0044a471db509c179d9d4a2c87e9)
- timeouts

<a id="canonical-33c14b1ed33a4a2ce601c1d28fe1dd5604477ded121dec11b70a22533a4cef08"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-53bf6c74158a713c760dcbcb54414ffacfb6f37148b3ee32d32eed49246126ea"></a>

## Direct properties — timeouts / 6689c1e5b091 / 3

<a id="canonical-0555e66ace5ad2174ddd84492e988e4d021dee054d3c777d553f2895afdd56c4"></a>

<a id="canonical-d3f200b643abe4ee575b4aedafdd5aa5bf91b4060402bff7f6b618b9e9fc636a"></a>

## create property — timeouts / 6689c1e5b091 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-14b784918a19683e6c39d5b7698fc65e521266dce9b11223869a6e206fa902b0"></a>

<a id="canonical-11dcb35424dd7a79c6ec8eff6f0417760b869b596eb88e66592d0cd3266f076b"></a>

## delete property — timeouts / 6689c1e5b091 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3da9455b984e6021db02150c6662c6a49e2115517c6a81462751ad1095cc2605"></a>

<a id="canonical-6f3cc586e9980b0ef21b13f219a2f031a290519295e8cc34b528779ea8d42ad0"></a>

## read property — timeouts / 6689c1e5b091 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-6c61a92a777c9d12d6029378116f0ae47c814aa04dafe9cc7ce42f2136282d13"></a>

<a id="canonical-c6705f9be6e9d7ce07544f5515c9bbb950feb54a2c2ab387d0d174181a0d0140"></a>

## update property — timeouts / 6689c1e5b091 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f60c35d10147df243f4b846d0e15e50c27e221b022426fe0b1a9dfb47219db2d"></a>

## Next pages — timeouts / 6689c1e5b091 / 8

- [Property reference](resources--ip_prefix_set--reference--group-001.md#canonical-e43df77cb707b03d9ea017a3256270c1386b0044a471db509c179d9d4a2c87e9)
- [xcsh_ip_prefix_set](../resources/ip_prefix_set.md#canonical-5a49cde9dc0a3a39ff028cd9a8304b00e480fb0bbedc342df01389fd1e3f30ae)
