---
page_title: "xcsh_protected_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain reference."
---

# xcsh_protected_domain reference

<a id="canonical-0d8b0bcd910ff49e6d0333184800f1f4d0ef950cb7ec9b09d4025255720e776c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d24df847c1b20b1588f1aa40354b9d16a29e4ba689db73fccb37fd92d67a4aca"></a>

## Property reference — Property reference / b4dea41a37f2 / 2

Breadcrumbs:

- [xcsh_protected_domain](../resources/protected_domain.md#canonical-df1bcdf4ef1e707566d60fcb995be20aae9625e7686c96bfd5a05fe9f6c6ded4)
- Property reference

<a id="canonical-8a9906278c17c29d9eb343dce3064922f05a57fd340c1852db0b4cdf283acf45"></a>

## Direct properties — Property reference / b4dea41a37f2 / 3

<a id="canonical-9ef4fa6a50d263ac8fdc40bfa0879c06803bfacdaf8b425048900e8083d78f94"></a>

<a id="canonical-ed27245def02ca39061068bbeed3a34db12b3bde3b33ae57458f659e43e87a0d"></a>

## annotations property — Property reference / b4dea41a37f2 / 4

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

<a id="canonical-12b70bfbe32f1a1ae8f1ab8946918daf10d9843f9d66cc3867e4d0feaa3f1532"></a>

<a id="canonical-eacaa239bd5414823d4d803558f527306e5f8e9f9561906a6955cd0cdb2b6d24"></a>

## description property — Property reference / b4dea41a37f2 / 5

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

<a id="canonical-914b8b07cc5cd507e8068aa98445785320de3a0514bbf76f57121851cd570013"></a>

<a id="canonical-53568fa66f8e5442fdd418bd7098b6eb12ca5f044b9307b7e6e98e6a660abc37"></a>

## disable property — Property reference / b4dea41a37f2 / 6

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

<a id="canonical-313bae8a02c24740376066afe8177926aaf046bd21a6f1aa97fd532c8efe8979"></a>

<a id="canonical-b27026822a1e0c469269beb2f6c34602764173f248ad81aaabcc19ceae670d0a"></a>

## id property — Property reference / b4dea41a37f2 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-6bc83a720dc4814557dabaa760807715a1c67e641a47eb66cfa54bdbf860b273"></a>

<a id="canonical-539e0c49267b515d6718a0d2a9bee29306bc8d865cdf6a6fdd77b2af6933a88c"></a>

## labels property — Property reference / b4dea41a37f2 / 8

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

<a id="canonical-1f36e96d90b82354d87399855be0ab6fe4d7ae7887f1f64d526b3bcf76204851"></a>

<a id="canonical-5fe0d804644b514c4270b95c702a3c32d0b24a941df815a95858d6bddf642bce"></a>

## name property — Property reference / b4dea41a37f2 / 9

Type: `"string"`. Required.

Name of the Protected Domain. Must be unique within the namespace.

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

<a id="canonical-3c26b40f5e538167d3e2013f7ba8b56e56c09927e5ed30de5f84fab20b46edcb"></a>

<a id="canonical-d99d6c14d0d03e61f913db1e54145ca7fd12ae1e9ebe7884b67f00c71f73f908"></a>

## namespace property — Property reference / b4dea41a37f2 / 10

Type: `"string"`. Required.

Namespace where the Protected Domain is created.

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

<a id="canonical-40be0be075c9bf792c3a98de63e06e8595e5f442931bb51bc5afd40e9d5447b5"></a>

<a id="canonical-4e5cb60e7aa2e12cd05c918250cafb07342cd6dab0db87ba683577db29f48c8e"></a>

## protected_domain property — Property reference / b4dea41a37f2 / 11

Type: `"string"`. Required.

For Client-Side Defense to work on the web pages where you injected the JS, you need to enter the
root domain below.

Upstream description:

For Client-Side Defense to work on the web pages where you injected the JS, you need to enter the
root domain below. Example: if you are adding Client-Side Defense JS on checkout.example.com, you
should enter example.com here.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
  validators.ETLDPlusOneValidator(),
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
    "ves.io.schema.rules.string.etld_plus_one": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.etld_plus_one": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [timeouts](resources--protected_domain--reference--group-001.md#canonical-6890f65ff78716ffc3c9bf34f055c6d88fb48c3d3e74c0b708c9e89b9e3009c7): complete subsection reference.

<a id="canonical-a2ed484f65849fe376d79cb0486bcd8f8c279fe1ae36a8e1f4e43615d5ec7f56"></a>

## All schema paths — Property reference / b4dea41a37f2 / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--protected_domain--reference--group-001.md#canonical-9ef4fa6a50d263ac8fdc40bfa0879c06803bfacdaf8b425048900e8083d78f94) |
| `description` | [description](resources--protected_domain--reference--group-001.md#canonical-12b70bfbe32f1a1ae8f1ab8946918daf10d9843f9d66cc3867e4d0feaa3f1532) |
| `disable` | [disable](resources--protected_domain--reference--group-001.md#canonical-914b8b07cc5cd507e8068aa98445785320de3a0514bbf76f57121851cd570013) |
| `id` | [id](resources--protected_domain--reference--group-001.md#canonical-313bae8a02c24740376066afe8177926aaf046bd21a6f1aa97fd532c8efe8979) |
| `labels` | [labels](resources--protected_domain--reference--group-001.md#canonical-6bc83a720dc4814557dabaa760807715a1c67e641a47eb66cfa54bdbf860b273) |
| `name` | [name](resources--protected_domain--reference--group-001.md#canonical-1f36e96d90b82354d87399855be0ab6fe4d7ae7887f1f64d526b3bcf76204851) |
| `namespace` | [namespace](resources--protected_domain--reference--group-001.md#canonical-3c26b40f5e538167d3e2013f7ba8b56e56c09927e5ed30de5f84fab20b46edcb) |
| `protected_domain` | [protected_domain](resources--protected_domain--reference--group-001.md#canonical-40be0be075c9bf792c3a98de63e06e8595e5f442931bb51bc5afd40e9d5447b5) |
| `timeouts` | [timeouts](resources--protected_domain--reference--group-001.md#canonical-97be014f82173b07cd50207423f09f753c45f3d2cc329b075b85368327056336) |
| `timeouts.create` | [timeouts.create](resources--protected_domain--reference--group-001.md#canonical-70f60096076d36467b07a708ad8ddc57d81db129ac33c762a76f9fe92e51a4ac) |
| `timeouts.delete` | [timeouts.delete](resources--protected_domain--reference--group-001.md#canonical-635960298a2d647ea421259b5787074336a95be57a6c8a633f91136809f68e9b) |
| `timeouts.read` | [timeouts.read](resources--protected_domain--reference--group-001.md#canonical-bc93a0d158950d3d1d52545f0b02c473eff394f6c23b96778784a9e7fdfc1d02) |
| `timeouts.update` | [timeouts.update](resources--protected_domain--reference--group-001.md#canonical-c4d9cee8997bcfbeb21d06317f6b2fbfafd886019022ce80d7798556186f64ec) |

<a id="canonical-04b6516e088697aa081203e02f0b07b98500edc8d26919b775e2bcaaa1a77da2"></a>

## Next pages — Property reference / b4dea41a37f2 / 13

- [timeouts](resources--protected_domain--reference--group-001.md#canonical-6890f65ff78716ffc3c9bf34f055c6d88fb48c3d3e74c0b708c9e89b9e3009c7)
- [xcsh_protected_domain](../resources/protected_domain.md#canonical-df1bcdf4ef1e707566d60fcb995be20aae9625e7686c96bfd5a05fe9f6c6ded4)

<a id="canonical-6890f65ff78716ffc3c9bf34f055c6d88fb48c3d3e74c0b708c9e89b9e3009c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b908c5b9e90d6a136661ea294c08282a39f2d0b98ffb26dbff4365fc69b3387"></a>

## timeouts — timeouts / c40fc262a484 / 2

Breadcrumbs:

- [xcsh_protected_domain](../resources/protected_domain.md#canonical-df1bcdf4ef1e707566d60fcb995be20aae9625e7686c96bfd5a05fe9f6c6ded4)
- [Property reference](resources--protected_domain--reference--group-001.md#canonical-0d8b0bcd910ff49e6d0333184800f1f4d0ef950cb7ec9b09d4025255720e776c)
- timeouts

<a id="canonical-97be014f82173b07cd50207423f09f753c45f3d2cc329b075b85368327056336"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-d5dc03166478e801f8cd676a43c14d13b452e63eee2592e96364175508a63e64"></a>

## Direct properties — timeouts / c40fc262a484 / 3

<a id="canonical-70f60096076d36467b07a708ad8ddc57d81db129ac33c762a76f9fe92e51a4ac"></a>

<a id="canonical-3a762f9f5fce9a8b454c483fc4d924710a3fb2307940020debd65379fde071b3"></a>

## create property — timeouts / c40fc262a484 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-635960298a2d647ea421259b5787074336a95be57a6c8a633f91136809f68e9b"></a>

<a id="canonical-75eb316c10eef58490ac7d31070179cfce437fccae9d124e69486bc4dec859bc"></a>

## delete property — timeouts / c40fc262a484 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-bc93a0d158950d3d1d52545f0b02c473eff394f6c23b96778784a9e7fdfc1d02"></a>

<a id="canonical-999385bfcda80da36c872d112d9a810c0f8d27a1b181be7231e0cf1c0c57f9d7"></a>

## read property — timeouts / c40fc262a484 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-c4d9cee8997bcfbeb21d06317f6b2fbfafd886019022ce80d7798556186f64ec"></a>

<a id="canonical-804219181c3452919f8ccfc8fb4d99ebfe277d22e835907688cd60777119377c"></a>

## update property — timeouts / c40fc262a484 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-85b365eecbb8a2e6efab7415ed490a8f47bedc3d194d0b37a3fb487417608f14"></a>

## Next pages — timeouts / c40fc262a484 / 8

- [Property reference](resources--protected_domain--reference--group-001.md#canonical-0d8b0bcd910ff49e6d0333184800f1f4d0ef950cb7ec9b09d4025255720e776c)
- [xcsh_protected_domain](../resources/protected_domain.md#canonical-df1bcdf4ef1e707566d60fcb995be20aae9625e7686c96bfd5a05fe9f6c6ded4)
