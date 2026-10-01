---
page_title: "xcsh_workload_flavor reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_workload_flavor reference."
---

# xcsh_workload_flavor reference

<a id="canonical-a1b2df5c2fff8d0b8fbb2c32a43b469b29d4036f48d674235772118f3099a1a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-153d0d9ca80c82624ec898ae5dac520fe8ef7d15a440f6b484d357ca160d4ffa"></a>

## Property reference — Property reference / a7c28b83fd4c / 2

Breadcrumbs:

- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-4dfec73998823e83c8926da172edff562c2f4bfac621f1a18b414eca82495d01)
- Property reference

<a id="canonical-73a993b7649a73f67b31e828bd188d9863d35b48c58d497aca676a5fdfd9be11"></a>

## Direct properties — Property reference / a7c28b83fd4c / 3

<a id="canonical-d15a11c3c5a476c4a5b139eb852a0d532c912e34452787250aed72111d6a1a04"></a>

<a id="canonical-f150e88e2d36b51e1694680b886a80efb430da98baad4861c41f6a0fbdeed901"></a>

## annotations property — Property reference / a7c28b83fd4c / 4

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

<a id="canonical-ac8b25f2b6eebbb986f8b83e87d817b5678a10ea7f07223d3f15d7b46fc505a8"></a>

<a id="canonical-cc047916e941353c2c917078c7b5f58e9a7f0e0e524b3c6ed02e130ef0872c13"></a>

## description property — Property reference / a7c28b83fd4c / 5

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

<a id="canonical-c66f3e751d426573836c75b20546cfe81b60dc7baf4150dd05b5aad59dc5bac3"></a>

<a id="canonical-e220e3d4802861dc2c63712d37b3bf83a53671734780d828ee1d866caf645336"></a>

## disable property — Property reference / a7c28b83fd4c / 6

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

<a id="canonical-c9466dde9fc2ea65df6e8d348d945763830c4e58622e574fb30c014dae9efa8e"></a>

<a id="canonical-b15a6a047fa2148cdbc039beaae870a37850132eec58ade16720d6b7cce2d716"></a>

## ephemeral_storage property — Property reference / a7c28b83fd4c / 7

Type: `"string"`. Optional, Computed.

Ephemeral storage in MiB (mebibyte) allocated for the workload\_flavor.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "1-6000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "1-6000"
  }
}
```

<a id="canonical-e08deed4c3889b8d2777152b943772ca33061bf350e3b0d5a9513bd8931277ec"></a>

<a id="canonical-11f80ae33819ce1b92866ee55adcdba2ca5017db083178ac31ddb47ef7492202"></a>

## id property — Property reference / a7c28b83fd4c / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-91e64bf3bb9966ee7966e79a268442e78dd75884bf7c41f029965d7c3de1d6f8"></a>

<a id="canonical-54099bee9ae4d2ab3d0ba1f32e28758c72a9cd84c0a79b047ba2668929683a65"></a>

## labels property — Property reference / a7c28b83fd4c / 9

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

<a id="canonical-aae43a3457d9b0246ede0e66561dad75bd30de0672decd766707bfc7e47b9583"></a>

<a id="canonical-bf0b6de3a9b9331649da818070348d6067bdba5547f8f29a1f0218ebe693a51d"></a>

## memory property — Property reference / a7c28b83fd4c / 10

Type: `"string"`. Optional, Computed.

Memory in MiB (mebibyte) allocated for the workload\_flavor.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "1-32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "1-32768"
  }
}
```

<a id="canonical-d16d0a74fb089baa44a689f6579cf365a7633a4235632aaa171b46a3069ab044"></a>

<a id="canonical-086f04843bd034238f39e7d8aaf615fa7df67732c4a90ad5826327eb6590c4a1"></a>

## name property — Property reference / a7c28b83fd4c / 11

Type: `"string"`. Required.

Name of the Workload Flavor. Must be unique within the namespace.

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

<a id="canonical-8d88243d05f44b26a367acb011b05fb0240426b6806aa4161434c5b43ffb7e3d"></a>

<a id="canonical-6fac3cfbadab0d9553f42a4a4b2151937839e2233c9201ed941b4c1ced906cdd"></a>

## namespace property — Property reference / a7c28b83fd4c / 12

Type: `"string"`. Optional, Computed.

Namespace for the Workload Flavor. The F5 XC API restricts this resource to the shared namespace; it
defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("shared")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("shared"),
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

- [timeouts](resources--workload_flavor--reference--group-001.md#canonical-ee9efcd2d4239a14ebfa0eb13f721e5d7db0428bf5c7247669283ef549818122): complete subsection reference.

<a id="canonical-b07bdab44d22c32cbcb4b2f9758c1e4762d4eed9a62ba54ab4796f66a71fdb59"></a>

<a id="canonical-532c64d7b0735f2fea7838dbd02e943c7f11d99495d274111ed6810918854490"></a>

## vcpus property — Property reference / a7c28b83fd4c / 13

Type: `"number"`. Optional, Computed.

Number of vCPUs allocated for the workload\_flavor. Each vCPU is a thread on a CPU core.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "8.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "8.0"
  }
}
```

<a id="canonical-e116829b4e3b2feadaefa48254294d6b1674d65d3a14bc4ec0f2b1c024c34665"></a>

## All schema paths — Property reference / a7c28b83fd4c / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--workload_flavor--reference--group-001.md#canonical-d15a11c3c5a476c4a5b139eb852a0d532c912e34452787250aed72111d6a1a04) |
| `description` | [description](resources--workload_flavor--reference--group-001.md#canonical-ac8b25f2b6eebbb986f8b83e87d817b5678a10ea7f07223d3f15d7b46fc505a8) |
| `disable` | [disable](resources--workload_flavor--reference--group-001.md#canonical-c66f3e751d426573836c75b20546cfe81b60dc7baf4150dd05b5aad59dc5bac3) |
| `ephemeral_storage` | [ephemeral_storage](resources--workload_flavor--reference--group-001.md#canonical-c9466dde9fc2ea65df6e8d348d945763830c4e58622e574fb30c014dae9efa8e) |
| `id` | [id](resources--workload_flavor--reference--group-001.md#canonical-e08deed4c3889b8d2777152b943772ca33061bf350e3b0d5a9513bd8931277ec) |
| `labels` | [labels](resources--workload_flavor--reference--group-001.md#canonical-91e64bf3bb9966ee7966e79a268442e78dd75884bf7c41f029965d7c3de1d6f8) |
| `memory` | [memory](resources--workload_flavor--reference--group-001.md#canonical-aae43a3457d9b0246ede0e66561dad75bd30de0672decd766707bfc7e47b9583) |
| `name` | [name](resources--workload_flavor--reference--group-001.md#canonical-d16d0a74fb089baa44a689f6579cf365a7633a4235632aaa171b46a3069ab044) |
| `namespace` | [namespace](resources--workload_flavor--reference--group-001.md#canonical-8d88243d05f44b26a367acb011b05fb0240426b6806aa4161434c5b43ffb7e3d) |
| `timeouts` | [timeouts](resources--workload_flavor--reference--group-001.md#canonical-e7740bfee14b7f3c0700e82e8ae4fa3a2a4dcdf5856d5dc418906bc74864853f) |
| `timeouts.create` | [timeouts.create](resources--workload_flavor--reference--group-001.md#canonical-f08d66f8dee33cb85dd94a6e06ca62826460933ce86fd9b6b2682839b32e0e75) |
| `timeouts.delete` | [timeouts.delete](resources--workload_flavor--reference--group-001.md#canonical-cfedfc5f97cb1f7485f6e085dd5dfd7b41e26052f369783e1f04e135b587c102) |
| `timeouts.read` | [timeouts.read](resources--workload_flavor--reference--group-001.md#canonical-8bd48d0fcb903e686dd0ac20609c4f39fdb38c3d17f95d60051660dbb25b9a4f) |
| `timeouts.update` | [timeouts.update](resources--workload_flavor--reference--group-001.md#canonical-92936231d2e3cf17e8d83bff89338d74df2571bafe37fb5823e56e529cabe697) |
| `vcpus` | [vcpus](resources--workload_flavor--reference--group-001.md#canonical-b07bdab44d22c32cbcb4b2f9758c1e4762d4eed9a62ba54ab4796f66a71fdb59) |

<a id="canonical-57314d78e33aab1a62f1334a1ad1fafd52d70d82e100f34fc643d50fd5689668"></a>

## Next pages — Property reference / a7c28b83fd4c / 15

- [timeouts](resources--workload_flavor--reference--group-001.md#canonical-ee9efcd2d4239a14ebfa0eb13f721e5d7db0428bf5c7247669283ef549818122)
- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-4dfec73998823e83c8926da172edff562c2f4bfac621f1a18b414eca82495d01)

<a id="canonical-ee9efcd2d4239a14ebfa0eb13f721e5d7db0428bf5c7247669283ef549818122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e009133c3bdec77e7874d3abad2fa606db5437715d2e857d0be86967eece27bf"></a>

## timeouts — timeouts / a7f38276c495 / 2

Breadcrumbs:

- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-4dfec73998823e83c8926da172edff562c2f4bfac621f1a18b414eca82495d01)
- [Property reference](resources--workload_flavor--reference--group-001.md#canonical-a1b2df5c2fff8d0b8fbb2c32a43b469b29d4036f48d674235772118f3099a1a7)
- timeouts

<a id="canonical-e7740bfee14b7f3c0700e82e8ae4fa3a2a4dcdf5856d5dc418906bc74864853f"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-38efb74786bb38997a33bfaa0c20e1bc65eb60b94f36c9670ff09c4b4a64c8fa"></a>

## Direct properties — timeouts / a7f38276c495 / 3

<a id="canonical-f08d66f8dee33cb85dd94a6e06ca62826460933ce86fd9b6b2682839b32e0e75"></a>

<a id="canonical-04e73f9406633502df28ab537c02dcb78ccb5ccd90408a2a80335a4e9bd7adca"></a>

## create property — timeouts / a7f38276c495 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-cfedfc5f97cb1f7485f6e085dd5dfd7b41e26052f369783e1f04e135b587c102"></a>

<a id="canonical-08027e973392a5c2aaf15aa9a71ede1e788a7ad2d1836114d87071cb9c173cac"></a>

## delete property — timeouts / a7f38276c495 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-8bd48d0fcb903e686dd0ac20609c4f39fdb38c3d17f95d60051660dbb25b9a4f"></a>

<a id="canonical-ea4a949812274c7b35fd1b1238a7968ad623c19118580192772b0dda04125b2e"></a>

## read property — timeouts / a7f38276c495 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-92936231d2e3cf17e8d83bff89338d74df2571bafe37fb5823e56e529cabe697"></a>

<a id="canonical-4e8ea03625e371feba789cb2b3b918d6e463054eb532054701da3815ac8195a5"></a>

## update property — timeouts / a7f38276c495 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-82923496f81b0ea10e95aa50b87785d2b663878c8caee7fad75e60f981c4e27a"></a>

## Next pages — timeouts / a7f38276c495 / 8

- [Property reference](resources--workload_flavor--reference--group-001.md#canonical-a1b2df5c2fff8d0b8fbb2c32a43b469b29d4036f48d674235772118f3099a1a7)
- [xcsh_workload_flavor](../resources/workload_flavor.md#canonical-4dfec73998823e83c8926da172edff562c2f4bfac621f1a18b414eca82495d01)
