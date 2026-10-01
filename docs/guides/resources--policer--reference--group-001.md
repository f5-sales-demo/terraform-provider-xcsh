---
page_title: "xcsh_policer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policer reference."
---

# xcsh_policer reference

<a id="canonical-25f1d1c8b4cdf547a8ecbbbcab78c6c88568d856b48014e86944888cf54e0170"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48ba691515003e1db76fd1c3e30c76e20364d583ac38423a8bc637e889cdf728"></a>

## Property reference — Property reference / 0de6145cd01b / 2

Breadcrumbs:

- [xcsh_policer](../resources/policer.md#canonical-9ad5ed8f191e509c3567b5ac04756c7a76ce4526ce5e90ae564852b6550c86d7)
- Property reference

<a id="canonical-5f393788fa57ec620cbff29ffc7db327e9cd00aa5af3ff8f59cca1d00d5057e6"></a>

## Direct properties — Property reference / 0de6145cd01b / 3

<a id="canonical-65341e39cc4d79787f4ef8977156412ee79480287d112a0dae709f238f67d7fb"></a>

<a id="canonical-004b43d95db5a030e19e4908d6f68da0382be83e1954c292e9576c3321db709b"></a>

## annotations property — Property reference / 0de6145cd01b / 4

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

<a id="canonical-383b0a51228b7054b62e515fe85400b3879f6d76b046c8c6d88a68fdd3323c71"></a>

<a id="canonical-420133b14bccff051bb0e02a2454a05c70ce638cb0033f0c33512094e652021e"></a>

## burst_size property — Property reference / 0de6145cd01b / 5

Type: `"number"`. Required.

The maximum size permitted for bursts of data. E.g. 10000 pps burst.

Upstream description:

The maximum size permitted for bursts of data. E.g. 10000 pps burst.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

<a id="canonical-5b1445e42cef54acb4ed6e6c059df5a4508563f6b40609053c40ac5096c9c9ad"></a>

<a id="canonical-a63b161bf3a80e66d4a827d2bf5c224384e3f24698baa84074319ea26e6a9b0e"></a>

## committed_information_rate property — Property reference / 0de6145cd01b / 6

Type: `"number"`. Required.

The committed information rate is the guaranteed packets rate for traffic arriving or departing
under normal conditions. E.g. 10000 pps.

Upstream description:

The committed information rate is the guaranteed packets rate for traffic arriving or departing
under normal conditions. E.g. 10000 pps.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 10000000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10000000,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10000000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "10000000"
  }
}
```

<a id="canonical-a5f16b1912909368ded49d401ce1b2e077a50c280c440228d47764c475cd0222"></a>

<a id="canonical-d5b7f058bc1111083928e859d34d93a355499474a0fff1725805203dc1ec6dcb"></a>

## description property — Property reference / 0de6145cd01b / 7

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

<a id="canonical-a2dc6738cfa5e47ee8074b9ea7ddf1acb70a37dd870f0af74e3b2c07a9684645"></a>

<a id="canonical-1796470c5ea4ba7a650ff7e2267ff351dac583fbe703dcd654a617d9f3796820"></a>

## disable property — Property reference / 0de6145cd01b / 8

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

<a id="canonical-161b0b87b9a1c648b3fde37dcc4f30e0e1a3d8df4c6267794efec66661110e13"></a>

<a id="canonical-6bebf40899ad6f3b4677781692637f71666bf22a44651ee308ee29a16d30c62f"></a>

## id property — Property reference / 0de6145cd01b / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-57e49419538c586d86721abb9b2d7b5707371a6c773769e2a7f9a53454ff64cc"></a>

<a id="canonical-f5f7fb6eb9f4bea9a7d8493fa23ff154cc4ddf40021206c6f41e9eb0dd1becb8"></a>

## labels property — Property reference / 0de6145cd01b / 10

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

<a id="canonical-0c7d83e93b92a8e6d3c90748fa50ea680ef3a8f3ecdf04081ea1423ce2e6d683"></a>

<a id="canonical-52fdafc2d1f96560eba25b7bfbfe10d9a378a0f511b4ce4ebf65aa86ffcc80d9"></a>

## name property — Property reference / 0de6145cd01b / 11

Type: `"string"`. Required.

Name of the Policer. Must be unique within the namespace.

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

<a id="canonical-62e465ab4d4929657347bdf97d18d31bb4d4c2b2d6626980447bc5e3e07b3d27"></a>

<a id="canonical-bfeb6c3d9c9f21b1a6bbe6b3fc8a1b36db450f792d4850834f3befd01c2fef09"></a>

## namespace property — Property reference / 0de6145cd01b / 12

Type: `"string"`. Required.

Namespace where the Policer is created.

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

<a id="canonical-5332b2d52d6f543c38ed865623c1d7df7c5b3900c5b9d8a89007e4208a98fc43"></a>

<a id="canonical-95820830846f33f1caedcf0fdd20288ce802d184e112b78e141178e6b9c8f4c4"></a>

## policer_mode property — Property reference / 0de6145cd01b / 13

Type: `"string"`. Optional, Computed.

\[Enum: POLICER\_MODE\_NOT\_SHARED|POLICER\_MODE\_SHARED\] - POLICER\_MODE\_NOT\_SHARED: Not Shared
A separate policer instance is created for each reference to the policer - POLICER\_MODE\_SHARED:
Shared A common policer instance is used for for all references to the policer. Possible values are
\`POLICER\_MODE\_NOT\_SHARED\`, \`POLICER\_MODE\_SHARED\`. Defaults to
\`POLICER\_MODE\_NOT\_SHARED\`. Server applies default when omitted.

Upstream description:

&#8203;- POLICER\_MODE\_NOT\_SHARED: Not Shared

A separate policer instance is created for each reference to the policer &#8203;-
POLICER\_MODE\_SHARED: Shared

A common policer instance is used for for all references to the policer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("POLICER_MODE_NOT_SHARED",
    "POLICER_MODE_SHARED"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "POLICER_MODE_NOT_SHARED",
  "enum": [
    "POLICER_MODE_NOT_SHARED",
    "POLICER_MODE_SHARED"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ba246b2e3fde6f5abef694828fc1cbfe3f4e3637e7888e6b25f7e8de822c1adb"></a>

<a id="canonical-07f26a0e5d38c3d9cec53c801efde852b52d6f11fd680b1d56efd977ef115565"></a>

## policer_type property — Property reference / 0de6145cd01b / 14

Type: `"string"`. Optional, Computed.

\[Enum: POLICER\_SINGLE\_RATE\_TWO\_COLOR\] Specifies the type of Policer Basic Single-Rate
Two-Color Policer. The only possible value is \`POLICER\_SINGLE\_RATE\_TWO\_COLOR\`. Defaults to
\`POLICER\_SINGLE\_RATE\_TWO\_COLOR\`. Server applies default when omitted.

Upstream description:

Specifies the type of Policer

Basic Single-Rate Two-Color Policer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("POLICER_SINGLE_RATE_TWO_COLOR"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "POLICER_SINGLE_RATE_TWO_COLOR",
  "enum": [
    "POLICER_SINGLE_RATE_TWO_COLOR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [timeouts](resources--policer--reference--group-001.md#canonical-5beee73aafa0a234f457b1ed904581afa4746f4afab552585c1c3dd1b5b76847): complete subsection reference.

<a id="canonical-3314c7cd97418a908bb4815379585925c28e2d5cdb911d34be2527bb974620fc"></a>

## All schema paths — Property reference / 0de6145cd01b / 15

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--policer--reference--group-001.md#canonical-65341e39cc4d79787f4ef8977156412ee79480287d112a0dae709f238f67d7fb) |
| `burst_size` | [burst_size](resources--policer--reference--group-001.md#canonical-383b0a51228b7054b62e515fe85400b3879f6d76b046c8c6d88a68fdd3323c71) |
| `committed_information_rate` | [committed_information_rate](resources--policer--reference--group-001.md#canonical-5b1445e42cef54acb4ed6e6c059df5a4508563f6b40609053c40ac5096c9c9ad) |
| `description` | [description](resources--policer--reference--group-001.md#canonical-a5f16b1912909368ded49d401ce1b2e077a50c280c440228d47764c475cd0222) |
| `disable` | [disable](resources--policer--reference--group-001.md#canonical-a2dc6738cfa5e47ee8074b9ea7ddf1acb70a37dd870f0af74e3b2c07a9684645) |
| `id` | [id](resources--policer--reference--group-001.md#canonical-161b0b87b9a1c648b3fde37dcc4f30e0e1a3d8df4c6267794efec66661110e13) |
| `labels` | [labels](resources--policer--reference--group-001.md#canonical-57e49419538c586d86721abb9b2d7b5707371a6c773769e2a7f9a53454ff64cc) |
| `name` | [name](resources--policer--reference--group-001.md#canonical-0c7d83e93b92a8e6d3c90748fa50ea680ef3a8f3ecdf04081ea1423ce2e6d683) |
| `namespace` | [namespace](resources--policer--reference--group-001.md#canonical-62e465ab4d4929657347bdf97d18d31bb4d4c2b2d6626980447bc5e3e07b3d27) |
| `policer_mode` | [policer_mode](resources--policer--reference--group-001.md#canonical-5332b2d52d6f543c38ed865623c1d7df7c5b3900c5b9d8a89007e4208a98fc43) |
| `policer_type` | [policer_type](resources--policer--reference--group-001.md#canonical-ba246b2e3fde6f5abef694828fc1cbfe3f4e3637e7888e6b25f7e8de822c1adb) |
| `timeouts` | [timeouts](resources--policer--reference--group-001.md#canonical-a2b2bc3a25fd5c1e8b52748a3088227b2d08a71145cb8c2da54c0fdf7141b5d4) |
| `timeouts.create` | [timeouts.create](resources--policer--reference--group-001.md#canonical-15f4f65538e0202f1745a115721e3efe6d747f12e11e753ef02c766d70886186) |
| `timeouts.delete` | [timeouts.delete](resources--policer--reference--group-001.md#canonical-7cb68df97be4ce7b0a23d135267cace88c57b915ec87be32ab69bd54c02c22c9) |
| `timeouts.read` | [timeouts.read](resources--policer--reference--group-001.md#canonical-7d236342048252416731b23ef464e4972d8025ab2f6629a6a155c4b70c9f1eb8) |
| `timeouts.update` | [timeouts.update](resources--policer--reference--group-001.md#canonical-be9a93bd6f6cf3fddb0938c788627749a076d8d65cd4cf3c6609e3202e9f3ebe) |

<a id="canonical-b079dcd7b9722081b152e68f8762dba20c4d0e2bfaf3cda18de072f360946c62"></a>

## Next pages — Property reference / 0de6145cd01b / 16

- [timeouts](resources--policer--reference--group-001.md#canonical-5beee73aafa0a234f457b1ed904581afa4746f4afab552585c1c3dd1b5b76847)
- [xcsh_policer](../resources/policer.md#canonical-9ad5ed8f191e509c3567b5ac04756c7a76ce4526ce5e90ae564852b6550c86d7)

<a id="canonical-5beee73aafa0a234f457b1ed904581afa4746f4afab552585c1c3dd1b5b76847"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-937ee6d45a4560862bd5d0020ff926fbab64b3db91016d21bee5134ffcc4850c"></a>

## timeouts — timeouts / 5d0c237761e2 / 2

Breadcrumbs:

- [xcsh_policer](../resources/policer.md#canonical-9ad5ed8f191e509c3567b5ac04756c7a76ce4526ce5e90ae564852b6550c86d7)
- [Property reference](resources--policer--reference--group-001.md#canonical-25f1d1c8b4cdf547a8ecbbbcab78c6c88568d856b48014e86944888cf54e0170)
- timeouts

<a id="canonical-a2b2bc3a25fd5c1e8b52748a3088227b2d08a71145cb8c2da54c0fdf7141b5d4"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-e486dd24202fb7e7d3e266524e0215a4e534151a61b018818ba323b3df476c4f"></a>

## Direct properties — timeouts / 5d0c237761e2 / 3

<a id="canonical-15f4f65538e0202f1745a115721e3efe6d747f12e11e753ef02c766d70886186"></a>

<a id="canonical-2cd4c2a265134ea14b3b188eec8fd4466bf887e1a404b574473be38f4a7d7b33"></a>

## create property — timeouts / 5d0c237761e2 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-7cb68df97be4ce7b0a23d135267cace88c57b915ec87be32ab69bd54c02c22c9"></a>

<a id="canonical-403511c59f3e2a3e1e130aa2f0767ee603ce6a983dac159e221902a28a1a19bb"></a>

## delete property — timeouts / 5d0c237761e2 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-7d236342048252416731b23ef464e4972d8025ab2f6629a6a155c4b70c9f1eb8"></a>

<a id="canonical-900fb01a49300221c1309e5b7cd0da5acf7d237ac36d49941d7c489bcd0b78f7"></a>

## read property — timeouts / 5d0c237761e2 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-be9a93bd6f6cf3fddb0938c788627749a076d8d65cd4cf3c6609e3202e9f3ebe"></a>

<a id="canonical-f0b2b5cb36d5821203b980c6606b479a3a5a8bfd0efccd91849a8d2310951fb3"></a>

## update property — timeouts / 5d0c237761e2 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-f29887cf8b131c3ebab985d546e3508dcefc9c3b3dfad070cba8ce2a7cb505b1"></a>

## Next pages — timeouts / 5d0c237761e2 / 8

- [Property reference](resources--policer--reference--group-001.md#canonical-25f1d1c8b4cdf547a8ecbbbcab78c6c88568d856b48014e86944888cf54e0170)
- [xcsh_policer](../resources/policer.md#canonical-9ad5ed8f191e509c3567b5ac04756c7a76ce4526ce5e90ae564852b6550c86d7)
