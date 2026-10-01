---
page_title: "xcsh_trusted_ca_list reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list reference."
---

# xcsh_trusted_ca_list reference

<a id="canonical-85409a678c309b8757816905fb4db6f52278501f089928c98decb8ccad0ed016"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef202bae29da9fd979266527f2b7969131bb703812d1464872dcea15d99a6202"></a>

## Property reference — Property reference / 1eeafaa38e4d / 2

Breadcrumbs:

- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-2a0175e3a6bbb4a391b22a048ce8c1c99fa0b6d99781050dc30b74fca0cbb5d6)
- Property reference

<a id="canonical-e7273ac8e5c40b03fa7cb05eb67e03c66c028945860c48250366670045c742be"></a>

## Direct properties — Property reference / 1eeafaa38e4d / 3

<a id="canonical-bf4a1d24fbf0cd3f4b7f85ac5600dbb2961f6854d166f4496900d33f957b3d27"></a>

<a id="canonical-85c95ce867480cdc93c35d19890fa37f1cca2a03a75562d83a7363765bc34bef"></a>

## annotations property — Property reference / 1eeafaa38e4d / 4

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

<a id="canonical-11e624d1393c807af5b859dfd80facdb3588acf45da8a930d6297e25fabceb22"></a>

<a id="canonical-a43f9505c4f6f4069a91e1fe2467b967d4f84c25ae0d510aac760ddfe894b8c0"></a>

## description property — Property reference / 1eeafaa38e4d / 5

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

<a id="canonical-5f11c4c82a62a6006142571697100b1f85771f468e53e69fbf97b82188e71e61"></a>

<a id="canonical-7263ff9b9de05fc960c1d5afcf877f7ebd5ea067d8e2e0d8e3514bb34a8e4429"></a>

## disable property — Property reference / 1eeafaa38e4d / 6

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

<a id="canonical-9b7e0ed93b15c1b39bcc1015013333b98153756fd0b96f4f6a94404b240aa104"></a>

<a id="canonical-7a4dc76e082d4d710252711e8eeb660f0289e867260665286abceb49a53e1ecc"></a>

## id property — Property reference / 1eeafaa38e4d / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-063cba405cad9a913bdc2fc64d0d56fdf458ef680ca958ccb75d72e8bdc40f87"></a>

<a id="canonical-cd191d49197556e4c9d8ec557d42a54208e97e7e78d9f21f18b7a4909bad30b4"></a>

## labels property — Property reference / 1eeafaa38e4d / 8

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

<a id="canonical-a6a48d6fdd83bdd531f96544c2248a5d91a3784683284e89c8ec374c1cbc3ea0"></a>

<a id="canonical-e19083bd03136229a07823a8acca9e9a4ba3f266429e6b3124225e8fcdcd92af"></a>

## name property — Property reference / 1eeafaa38e4d / 9

Type: `"string"`. Required.

Name of the Trusted CA List. Must be unique within the namespace.

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

<a id="canonical-cdf6f66ad23600ee2eef848348f061fa67ced79f5b0dca142da82f1791f0af8a"></a>

<a id="canonical-e6ee1aac8bba0f5ce2ca1399ec11255fc80f75ce5cc5a905375ea5ba3180ca83"></a>

## namespace property — Property reference / 1eeafaa38e4d / 10

Type: `"string"`. Required.

Namespace where the Trusted CA List is created.

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

- [timeouts](resources--trusted_ca_list--reference--group-001.md#canonical-e3905bfb4ab22205f1c83c5957ee19959c323aabaa2478e5ff41941fdf481b77): complete subsection reference.

<a id="canonical-fef3016c021d522a54e6ac4d068028d7cc3d6e6414450a978100e11b5d2baa41"></a>

<a id="canonical-1e24b189b6982834814235e6ecbfe7b09af821003be1d918cb4d72b4622c643c"></a>

## trusted_ca_url property — Property reference / 1eeafaa38e4d / 11

Type: `"string"`. Optional, Computed.

Trusted CA certificates for validating certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.max_len": "512000",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512000",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-bab344aae6992b0ab0bf9ab55119c235178b496956719f98e3793d700b508559"></a>

## All schema paths — Property reference / 1eeafaa38e4d / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--trusted_ca_list--reference--group-001.md#canonical-bf4a1d24fbf0cd3f4b7f85ac5600dbb2961f6854d166f4496900d33f957b3d27) |
| `description` | [description](resources--trusted_ca_list--reference--group-001.md#canonical-11e624d1393c807af5b859dfd80facdb3588acf45da8a930d6297e25fabceb22) |
| `disable` | [disable](resources--trusted_ca_list--reference--group-001.md#canonical-5f11c4c82a62a6006142571697100b1f85771f468e53e69fbf97b82188e71e61) |
| `id` | [id](resources--trusted_ca_list--reference--group-001.md#canonical-9b7e0ed93b15c1b39bcc1015013333b98153756fd0b96f4f6a94404b240aa104) |
| `labels` | [labels](resources--trusted_ca_list--reference--group-001.md#canonical-063cba405cad9a913bdc2fc64d0d56fdf458ef680ca958ccb75d72e8bdc40f87) |
| `name` | [name](resources--trusted_ca_list--reference--group-001.md#canonical-a6a48d6fdd83bdd531f96544c2248a5d91a3784683284e89c8ec374c1cbc3ea0) |
| `namespace` | [namespace](resources--trusted_ca_list--reference--group-001.md#canonical-cdf6f66ad23600ee2eef848348f061fa67ced79f5b0dca142da82f1791f0af8a) |
| `timeouts` | [timeouts](resources--trusted_ca_list--reference--group-001.md#canonical-0e05dd2b9b5320ab04e0256c71038654921ff77e9efa132a4ad69bedf6d06847) |
| `timeouts.create` | [timeouts.create](resources--trusted_ca_list--reference--group-001.md#canonical-0354a8b3b06450ee4b063559c29ebb5427a0624a9295cb4b6e6362377e20fc23) |
| `timeouts.delete` | [timeouts.delete](resources--trusted_ca_list--reference--group-001.md#canonical-1c3e98784da29f69f33951f06f19345afc39fe0118cbd74672701ef8e6dbeb71) |
| `timeouts.read` | [timeouts.read](resources--trusted_ca_list--reference--group-001.md#canonical-c4955b3034964871212aa327b819cb479aef3f06d6eda41fa6fd57ccb5c1f182) |
| `timeouts.update` | [timeouts.update](resources--trusted_ca_list--reference--group-001.md#canonical-f2b42eb95172d8d4b90141df0f95c4e90e30c7ec76345baf204b3ac6e1f0e34b) |
| `trusted_ca_url` | [trusted_ca_url](resources--trusted_ca_list--reference--group-001.md#canonical-fef3016c021d522a54e6ac4d068028d7cc3d6e6414450a978100e11b5d2baa41) |

<a id="canonical-f9fee68f72701ca4de85b9325eebbf6b3cf56568eac974706d1449504ae99d71"></a>

## Next pages — Property reference / 1eeafaa38e4d / 13

- [timeouts](resources--trusted_ca_list--reference--group-001.md#canonical-e3905bfb4ab22205f1c83c5957ee19959c323aabaa2478e5ff41941fdf481b77)
- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-2a0175e3a6bbb4a391b22a048ce8c1c99fa0b6d99781050dc30b74fca0cbb5d6)

<a id="canonical-e3905bfb4ab22205f1c83c5957ee19959c323aabaa2478e5ff41941fdf481b77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2529e2e96fbe35d207985ba7c2b709a4a0455cb90ba5ffd1fd30326aa13192bd"></a>

## timeouts — timeouts / d26ea723fb38 / 2

Breadcrumbs:

- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-2a0175e3a6bbb4a391b22a048ce8c1c99fa0b6d99781050dc30b74fca0cbb5d6)
- [Property reference](resources--trusted_ca_list--reference--group-001.md#canonical-85409a678c309b8757816905fb4db6f52278501f089928c98decb8ccad0ed016)
- timeouts

<a id="canonical-0e05dd2b9b5320ab04e0256c71038654921ff77e9efa132a4ad69bedf6d06847"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-b123dec89a16f7deef7fbf3f5c15690eba62af49111512615dc74eaa778e803f"></a>

## Direct properties — timeouts / d26ea723fb38 / 3

<a id="canonical-0354a8b3b06450ee4b063559c29ebb5427a0624a9295cb4b6e6362377e20fc23"></a>

<a id="canonical-62a891066d892ae308ba83700185493efcc46d4b686dee8e2bcafb2f0ff727a3"></a>

## create property — timeouts / d26ea723fb38 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1c3e98784da29f69f33951f06f19345afc39fe0118cbd74672701ef8e6dbeb71"></a>

<a id="canonical-6adb6a8991160ff31a53a522619dc331938ce0031a62fc5d971cdb1917f8d3b7"></a>

## delete property — timeouts / d26ea723fb38 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-c4955b3034964871212aa327b819cb479aef3f06d6eda41fa6fd57ccb5c1f182"></a>

<a id="canonical-4c21dc2a66eaa1bd30e454a660ac355411344663ad3e0b53449cfa37fc7ceb50"></a>

## read property — timeouts / d26ea723fb38 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-f2b42eb95172d8d4b90141df0f95c4e90e30c7ec76345baf204b3ac6e1f0e34b"></a>

<a id="canonical-38108da6e7ecb21f1c9f9b9cb375a386f3f81b76cb85b250fe66bad76ea4ec9e"></a>

## update property — timeouts / d26ea723fb38 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-9c4ee344fd02e4db9a88205b48449595d20bef3413c90e5a1007905f9c52b038"></a>

## Next pages — timeouts / d26ea723fb38 / 8

- [Property reference](resources--trusted_ca_list--reference--group-001.md#canonical-85409a678c309b8757816905fb4db6f52278501f089928c98decb8ccad0ed016)
- [xcsh_trusted_ca_list](../resources/trusted_ca_list.md#canonical-2a0175e3a6bbb4a391b22a048ce8c1c99fa0b6d99781050dc30b74fca0cbb5d6)
