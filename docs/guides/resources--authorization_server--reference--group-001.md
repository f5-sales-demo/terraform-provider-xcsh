---
page_title: "xcsh_authorization_server reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server reference."
---

# xcsh_authorization_server reference

<a id="canonical-2554adc63fdb57eae149d6c2409f7a163791404f3563f7d1e34c7446f838913d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e01d3e94e711315801f29b74d100aae82e31709668bbab17a582b3116f21dfdb"></a>

## Property reference — Property reference / aaf69854c3e7 / 2

Breadcrumbs:

- [xcsh_authorization_server](../resources/authorization_server.md#canonical-db153c7e889fa4d4a8d9c87b182594a4c4c9895ed0e6b0b948eddea4da3b2cbd)
- Property reference

<a id="canonical-3a3186a6bbe981ea88812f12e1eacaa92acdd10a764b5d51e60fd16661df3b59"></a>

## Direct properties — Property reference / aaf69854c3e7 / 3

<a id="canonical-d064222ec9bf98a5e9ef93a79ec8706ab308a708ab6cd930c34bc730cd72e8ea"></a>

<a id="canonical-2866ff57b2ae5a92132b50015602a49e578325b3bf188b45b1bf5bb3357b7abe"></a>

## annotations property — Property reference / aaf69854c3e7 / 4

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

<a id="canonical-9f3af4f64e8493910c9dcad315a392ddaecbf4d5e0b33ceaf8593db1076ec83d"></a>

<a id="canonical-7ca76ed2a81552f9a3502e36a2ecff4e4887222ef660006bf5720ac825141c6c"></a>

## description property — Property reference / aaf69854c3e7 / 5

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

<a id="canonical-4315431453a5545b7dc48fbe50da4ad31b83edc2b2589cca13a4a5d53ecdf4ee"></a>

<a id="canonical-225e1460f07e57d31e79f195570f637f0cce69c8e228f44dfab23e1a666f8242"></a>

## disable property — Property reference / aaf69854c3e7 / 6

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

<a id="canonical-2927a9803e095aaa9d1a56546ff6282c6bd8379a6c2e000b06d345edd2d0b874"></a>

<a id="canonical-2f4d170cf2dd60b2ecb762d3e6289887fd43531708aaa62e1f1f03d54067c7c8"></a>

## id property — Property reference / aaf69854c3e7 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-bfea0f53f170bd14e5afad1b354bafded67c097ed69a570b3bba822931c8e2ff"></a>

<a id="canonical-e724c39f673be40eaa7e551a4597cc2e0840c6ccba3e7f5d606ada78172d1698"></a>

## jwks_uri property — Property reference / aaf69854c3e7 / 8

Type: `"string"`. Required.

X-textBlockContent: Automatic fetching of JWKS will happen once daily. You can also do it manually
from the list of Authorization Servers at any time.

Upstream description:

X-textBlockContent: Automatic fetching of JWKS will happen once daily. You can also do it manually
from the list of Authorization Servers at any time.

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
    "format": "uri",
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.uri": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.uri": "true"
  }
}
```

<a id="canonical-6c21dd91d0451b5b7d3965443fdedf7e289e2028f4e3129ad4d05e7067e37d72"></a>

<a id="canonical-1a4623fce8032d6e15c57b1bcad795527d0c26d2c5cb259239e67105bd030154"></a>

## labels property — Property reference / aaf69854c3e7 / 9

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

<a id="canonical-f63ac82a9584ea1553389a727d6613af623a05eab9e4b78041851dd315682e63"></a>

<a id="canonical-fc00ed51d041e6e1795ad59b4b6dbcab48e2c6b156b62a285681f4991afa4a42"></a>

## name property — Property reference / aaf69854c3e7 / 10

Type: `"string"`. Required.

Name of the Authorization Server. Must be unique within the namespace.

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

<a id="canonical-573a46c203e1d4ec3edc7047a31cf52b5cf60bf6636adc58253660ae2ff63e49"></a>

<a id="canonical-473d20c6f250a0b47c198d72919f6a505a55810a666be8282d2705f3210a6dff"></a>

## namespace property — Property reference / aaf69854c3e7 / 11

Type: `"string"`. Required.

Namespace where the Authorization Server is created.

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

- [timeouts](resources--authorization_server--reference--group-001.md#canonical-f8b90fac92ebe2f574e52212c6cd1b071c869a1bdd0b4fd3e1445b1728586970): complete subsection reference.

<a id="canonical-9c4d60c8ca58ef8abbb417cd462992eb943873e5f5913717728a7fea162608ae"></a>

## All schema paths — Property reference / aaf69854c3e7 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--authorization_server--reference--group-001.md#canonical-d064222ec9bf98a5e9ef93a79ec8706ab308a708ab6cd930c34bc730cd72e8ea) |
| `description` | [description](resources--authorization_server--reference--group-001.md#canonical-9f3af4f64e8493910c9dcad315a392ddaecbf4d5e0b33ceaf8593db1076ec83d) |
| `disable` | [disable](resources--authorization_server--reference--group-001.md#canonical-4315431453a5545b7dc48fbe50da4ad31b83edc2b2589cca13a4a5d53ecdf4ee) |
| `id` | [id](resources--authorization_server--reference--group-001.md#canonical-2927a9803e095aaa9d1a56546ff6282c6bd8379a6c2e000b06d345edd2d0b874) |
| `jwks_uri` | [jwks_uri](resources--authorization_server--reference--group-001.md#canonical-bfea0f53f170bd14e5afad1b354bafded67c097ed69a570b3bba822931c8e2ff) |
| `labels` | [labels](resources--authorization_server--reference--group-001.md#canonical-6c21dd91d0451b5b7d3965443fdedf7e289e2028f4e3129ad4d05e7067e37d72) |
| `name` | [name](resources--authorization_server--reference--group-001.md#canonical-f63ac82a9584ea1553389a727d6613af623a05eab9e4b78041851dd315682e63) |
| `namespace` | [namespace](resources--authorization_server--reference--group-001.md#canonical-573a46c203e1d4ec3edc7047a31cf52b5cf60bf6636adc58253660ae2ff63e49) |
| `timeouts` | [timeouts](resources--authorization_server--reference--group-001.md#canonical-8018f37aa6c819746a9aa8c328c8a1c7ccea03eccbda16bc20c7c1ed668a174a) |
| `timeouts.create` | [timeouts.create](resources--authorization_server--reference--group-001.md#canonical-509e72c6e116ef5ef91757c66d8f607456bde8855942ea6e0ae5aff702b5a0fe) |
| `timeouts.delete` | [timeouts.delete](resources--authorization_server--reference--group-001.md#canonical-ff14dc44239f90ec36611527b18a3e55b598c650e40f00cdf539104c58d944e2) |
| `timeouts.read` | [timeouts.read](resources--authorization_server--reference--group-001.md#canonical-c97d2d68cb28021924de52f97e6d7bfafa70157cf4f941da7f162ef899af8af9) |
| `timeouts.update` | [timeouts.update](resources--authorization_server--reference--group-001.md#canonical-3400ecd7a3b86de7f39b8453ff6d91e0d6aaccad27805f92cad6bb2938b155ec) |

<a id="canonical-18f5ac37c356ff9723af219320264ac9567d836f674a8a7f86450311b3743755"></a>

## Next pages — Property reference / aaf69854c3e7 / 13

- [timeouts](resources--authorization_server--reference--group-001.md#canonical-f8b90fac92ebe2f574e52212c6cd1b071c869a1bdd0b4fd3e1445b1728586970)
- [xcsh_authorization_server](../resources/authorization_server.md#canonical-db153c7e889fa4d4a8d9c87b182594a4c4c9895ed0e6b0b948eddea4da3b2cbd)

<a id="canonical-f8b90fac92ebe2f574e52212c6cd1b071c869a1bdd0b4fd3e1445b1728586970"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f9e58334ed56d78a0f57b8d4a875d34189eef084ffb656c05c3690108b3d750"></a>

## timeouts — timeouts / f94c167e8d55 / 2

Breadcrumbs:

- [xcsh_authorization_server](../resources/authorization_server.md#canonical-db153c7e889fa4d4a8d9c87b182594a4c4c9895ed0e6b0b948eddea4da3b2cbd)
- [Property reference](resources--authorization_server--reference--group-001.md#canonical-2554adc63fdb57eae149d6c2409f7a163791404f3563f7d1e34c7446f838913d)
- timeouts

<a id="canonical-8018f37aa6c819746a9aa8c328c8a1c7ccea03eccbda16bc20c7c1ed668a174a"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0ac65599f009ed8fb24dfcf911e5d21cab1e89eb5d05b5d56107bd0965047860"></a>

## Direct properties — timeouts / f94c167e8d55 / 3

<a id="canonical-509e72c6e116ef5ef91757c66d8f607456bde8855942ea6e0ae5aff702b5a0fe"></a>

<a id="canonical-f6731bad63fc500f77efe8ea98f24cf0aad7c0d33d42a9bf98edf588b8c5f1b2"></a>

## create property — timeouts / f94c167e8d55 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-ff14dc44239f90ec36611527b18a3e55b598c650e40f00cdf539104c58d944e2"></a>

<a id="canonical-feb0d5c2629ec5bf0785f5b86135aad927c1014eee086c18a39b0373532fbc24"></a>

## delete property — timeouts / f94c167e8d55 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-c97d2d68cb28021924de52f97e6d7bfafa70157cf4f941da7f162ef899af8af9"></a>

<a id="canonical-1e579aa2e95ec9695de6c954ec288460285a9da5cacb2e934141adc9b292e6d8"></a>

## read property — timeouts / f94c167e8d55 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3400ecd7a3b86de7f39b8453ff6d91e0d6aaccad27805f92cad6bb2938b155ec"></a>

<a id="canonical-ffd01c7f3087cf7724984adf7a92040d014e088b5eb050e8e92ba5cb1495b6eb"></a>

## update property — timeouts / f94c167e8d55 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-4e23c45c1bfa7c905598df8fb914e0847a235fa564f0f9b3b341bb7077bd191c"></a>

## Next pages — timeouts / f94c167e8d55 / 8

- [Property reference](resources--authorization_server--reference--group-001.md#canonical-2554adc63fdb57eae149d6c2409f7a163791404f3563f7d1e34c7446f838913d)
- [xcsh_authorization_server](../resources/authorization_server.md#canonical-db153c7e889fa4d4a8d9c87b182594a4c4c9895ed0e6b0b948eddea4da3b2cbd)
