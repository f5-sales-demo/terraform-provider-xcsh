---
page_title: "xcsh_cminstance reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cminstance reference."
---

# xcsh_cminstance reference

<a id="canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-54917e46cf9c7f2d8c221574e8bbb8ec152befea658239cdd1a418eb6fb77cca"></a>

## Property reference — Property reference / fe44b091156a / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
- Property reference

<a id="canonical-9642f213055bf55ed09af6d81d74e5dc9045e911a60927fc0d0e2b128271ed13"></a>

## Direct properties — Property reference / fe44b091156a / 3

<a id="canonical-e874bc016b54c4639fcb5cd537e668cbd29f733f48357fd6bac3bd0d86b135d7"></a>

<a id="canonical-06e72e7c83595aca67e4f1ad2eecb6fa1bb39abe98b50c4a88ae26ab7d1d4942"></a>

## annotations property — Property reference / fe44b091156a / 4

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

- [api_token](data-sources--cminstance--reference--group-001.md#canonical-9a17ed398623aea6931d85b571015840e4143eed629ba2b501d96fd345f64b4d): complete subsection reference.

<a id="canonical-e7e710409030a8df0f542caf11cd3082efd6b165cfd8baf8ac10a66f74e459e4"></a>

<a id="canonical-0c587e8bfa7529552cc10266dc092f0b38c4033d132d778432fea365b9ccf85a"></a>

## description property — Property reference / fe44b091156a / 5

Type: `"string"`. Computed.

Description of the Cminstance.

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

<a id="canonical-f3298d2708433f8496429fd3780e3fc885c698ad4de4d6c1322a15f495409cce"></a>

<a id="canonical-655f966ec1a4e387eb33166b692dee1c73cb247adf488f19c3583cab8498ee35"></a>

## id property — Property reference / fe44b091156a / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ip](data-sources--cminstance--reference--group-001.md#canonical-3c6339205e271a123677ab965f68ca38a30fdfe92908accda8a304262e0a1455): complete subsection reference.

<a id="canonical-185b7954d27c5dc16558c217c78254222c883df669e60f85ab55f82377285245"></a>

<a id="canonical-c0001409f59df84f981a9da445e9678fdba2dc8d4a8f823dc75e0e5e7c58fe1d"></a>

## labels property — Property reference / fe44b091156a / 7

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

<a id="canonical-2e6e920d1ce6393486cddf58fb4c054ae6444e39e46dbcac4be2743e0faebba2"></a>

<a id="canonical-87564e0df19e725b1f9d94561c773b13f92fd2785ed4279115b6f72f31a1543f"></a>

## name property — Property reference / fe44b091156a / 8

Type: `"string"`. Required.

Name of the Cminstance.

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

<a id="canonical-939a187beb3e26556089434159625b45866d351be1d4bc50f2d6dd128baf589c"></a>

<a id="canonical-f3727b9725b2270448f7b935480c08373195d15fddeebecdcb1a8edfb98498db"></a>

## namespace property — Property reference / fe44b091156a / 9

Type: `"string"`. Required.

Namespace where the Cminstance exists.

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

- [password](data-sources--cminstance--reference--group-001.md#canonical-7a57d86ff7cc215f8bda1c753c142ff62b525be0049e23794b698a0d27e0c4d9): complete subsection reference.

<a id="canonical-580529ad6d5f728e3db0c8e98acc8e44a0f562b7817a5ceb710aed9b36bf8ed1"></a>

<a id="canonical-80e8e735905749f4756f6129a6bdd33358035e3422aa6ac98fa5a1a9774a60fc"></a>

## port property — Property reference / fe44b091156a / 10

Type: `"number"`. Computed.

Port of the Central Manager instance to connect to.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-edd357773ac61c7be50b5d683fcf4e7eaabf1252a708a4f7c1d34d16bee91b1b"></a>

<a id="canonical-0bf10e7a6043342b786720f73b2d39cd5674ae3f7f3f5a89853f90b33aeba1ce"></a>

## username property — Property reference / fe44b091156a / 11

Type: `"string"`. Computed.

Username for the Central Manager instance.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "4"
  }
}
```

<a id="canonical-29613cffc6401490974272fb7583ce0856f4a407603b31e0a3078a55ed4d4ccc"></a>

## All schema paths — Property reference / fe44b091156a / 12

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cminstance--reference--group-001.md#canonical-e874bc016b54c4639fcb5cd537e668cbd29f733f48357fd6bac3bd0d86b135d7) |
| `api_token` | [api_token](data-sources--cminstance--reference--group-001.md#canonical-6f544b3a044dd606c2826f9f21c4b28d9996326ad5633249ae07440c6a2fc704) |
| `api_token.blindfold_secret_info` | [api_token.blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-ca3736c6bd5086dfee3d26c3e2aa51da63e2c3d7ad8e010d8abd56d00e7b0a1a) |
| `api_token.blindfold_secret_info.decryption_provider` | [api_token.blindfold_secret_info.decryption_provider](data-sources--cminstance--reference--group-001.md#canonical-4ade364d2fe9364e88b11418e8f629c630cfb134d2fb0076871f830861643185) |
| `api_token.blindfold_secret_info.location` | [api_token.blindfold_secret_info.location](data-sources--cminstance--reference--group-001.md#canonical-c8e1bcd03b3c4d6dcc9b7e062b69260983941593c5ceebc5f8c5f6a303c69890) |
| `api_token.blindfold_secret_info.store_provider` | [api_token.blindfold_secret_info.store_provider](data-sources--cminstance--reference--group-001.md#canonical-2bac30fc85cbcfae4c64c21a61629e4f83cc3a7b86f1c76d10261c78df1352b2) |
| `api_token.clear_secret_info` | [api_token.clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-7a11c48e475f7f1f4a1a349ec27d757896cab5afeec13aaed0e3e785583843b4) |
| `api_token.clear_secret_info.provider_ref` | [api_token.clear_secret_info.provider_ref](data-sources--cminstance--reference--group-001.md#canonical-4addf483c354bec45d8d5dd41a94068319d93d88cdd8f58efd6b7a854d518026) |
| `api_token.clear_secret_info.url` | [api_token.clear_secret_info.url](data-sources--cminstance--reference--group-001.md#canonical-623093d6f4ac2a68b08bc3ee5d1488e731d23a96f614c1ed7d3498947e00db50) |
| `description` | [description](data-sources--cminstance--reference--group-001.md#canonical-e7e710409030a8df0f542caf11cd3082efd6b165cfd8baf8ac10a66f74e459e4) |
| `id` | [id](data-sources--cminstance--reference--group-001.md#canonical-f3298d2708433f8496429fd3780e3fc885c698ad4de4d6c1322a15f495409cce) |
| `ip` | [ip](data-sources--cminstance--reference--group-001.md#canonical-dabe639ec17c62500a4d7e03fbc18362ec202f49484ef0eaff518f3a66e078e8) |
| `ip.addr` | [ip.addr](data-sources--cminstance--reference--group-001.md#canonical-8d86873aae7619b7789106d38166893f0465e17417c3fd3c7ac652422f80f22d) |
| `labels` | [labels](data-sources--cminstance--reference--group-001.md#canonical-185b7954d27c5dc16558c217c78254222c883df669e60f85ab55f82377285245) |
| `name` | [name](data-sources--cminstance--reference--group-001.md#canonical-2e6e920d1ce6393486cddf58fb4c054ae6444e39e46dbcac4be2743e0faebba2) |
| `namespace` | [namespace](data-sources--cminstance--reference--group-001.md#canonical-939a187beb3e26556089434159625b45866d351be1d4bc50f2d6dd128baf589c) |
| `password` | [password](data-sources--cminstance--reference--group-001.md#canonical-b8ad4060eb3dcc8640eeec174993a559d7911942fd7e01421ba7f3930199e0ff) |
| `password.blindfold_secret_info` | [password.blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-94b435540e17097ac1f44fe06b257f32b7b2c3b331dfe5072aeeaee26aa06054) |
| `password.blindfold_secret_info.decryption_provider` | [password.blindfold_secret_info.decryption_provider](data-sources--cminstance--reference--group-001.md#canonical-008a826e6aab8bd718274d742bd4cf4b0006ffebf6b837b958ee29f12e34d10f) |
| `password.blindfold_secret_info.location` | [password.blindfold_secret_info.location](data-sources--cminstance--reference--group-001.md#canonical-8fc7e35c35934d370a6bf4212db5354df5a08a2a31495425b79f257031b8a8dd) |
| `password.blindfold_secret_info.store_provider` | [password.blindfold_secret_info.store_provider](data-sources--cminstance--reference--group-001.md#canonical-598949b28bea46354cd42b4fbb845030957d4593a3a086a60706b83618c9481f) |
| `password.clear_secret_info` | [password.clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-d3f2eb1ea2dfd1c9675ec79d27e49c2f507140a2c4a8a33114e4e0e52eb7981b) |
| `password.clear_secret_info.provider_ref` | [password.clear_secret_info.provider_ref](data-sources--cminstance--reference--group-001.md#canonical-f22ea42f855dc53f73d004d7150c1291dfe78abaccb613df1dc0accb177abcf5) |
| `password.clear_secret_info.url` | [password.clear_secret_info.url](data-sources--cminstance--reference--group-001.md#canonical-c42973aff656d041fe367c8a9c812e7c08fd808bc7aefe0f026df181972bcc96) |
| `port` | [port](data-sources--cminstance--reference--group-001.md#canonical-580529ad6d5f728e3db0c8e98acc8e44a0f562b7817a5ceb710aed9b36bf8ed1) |
| `username` | [username](data-sources--cminstance--reference--group-001.md#canonical-edd357773ac61c7be50b5d683fcf4e7eaabf1252a708a4f7c1d34d16bee91b1b) |

<a id="canonical-7129229f582405aa17c73ea3817ff7d7dbabc52f4eb97ef1ea1fc3ca268142c4"></a>

## Next pages — Property reference / fe44b091156a / 13

- [api_token](data-sources--cminstance--reference--group-001.md#canonical-9a17ed398623aea6931d85b571015840e4143eed629ba2b501d96fd345f64b4d)
- [ip](data-sources--cminstance--reference--group-001.md#canonical-3c6339205e271a123677ab965f68ca38a30fdfe92908accda8a304262e0a1455)
- [password](data-sources--cminstance--reference--group-001.md#canonical-7a57d86ff7cc215f8bda1c753c142ff62b525be0049e23794b698a0d27e0c4d9)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)

<a id="canonical-9a17ed398623aea6931d85b571015840e4143eed629ba2b501d96fd345f64b4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dae82b6a83c58c1bb69a6affcb55e466d2e60782d58a8c52f3bfed842604d0bd"></a>

## api_token — api_token / 93d19ee28119 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2)
- api_token

<a id="canonical-6f544b3a044dd606c2826f9f21c4b28d9996326ad5633249ae07440c6a2fc704"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-24d644aea0c4919b155d8d89e4a180aac99834ca57218b1db08150d153ccf402"></a>

## Direct properties — api_token / 93d19ee28119 / 3

- [blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-0cda3c7f959ee0e0385e6975b901bc45cf8ff01d33fa5f14a3d417fea2cd4ca9): complete subsection reference.

- [clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-cf29c528bfe2852f37fe27d25fc84f04c6d0412ee6d4840049af826bafcf6438): complete subsection reference.

<a id="canonical-c58c8b6323d0b79017beb31b3220468e926d09e3a587c120f500ad65122395e8"></a>

## Next pages — api_token / 93d19ee28119 / 4

- [api_token.blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-0cda3c7f959ee0e0385e6975b901bc45cf8ff01d33fa5f14a3d417fea2cd4ca9)
- [api_token.clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-cf29c528bfe2852f37fe27d25fc84f04c6d0412ee6d4840049af826bafcf6438)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)

<a id="canonical-0cda3c7f959ee0e0385e6975b901bc45cf8ff01d33fa5f14a3d417fea2cd4ca9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8253db4598145ba1f44a52cbd74347e291d7d9ada951e03bb21497a7fec0583"></a>

## api_token.blindfold_secret_info — api_token.blindfold_secret_info / 80f8bcb30d4b / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2)
- [api_token](data-sources--cminstance--reference--group-001.md#canonical-9a17ed398623aea6931d85b571015840e4143eed629ba2b501d96fd345f64b4d)
- api_token.blindfold_secret_info

<a id="canonical-ca3736c6bd5086dfee3d26c3e2aa51da63e2c3d7ad8e010d8abd56d00e7b0a1a"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-9a82ff1aa95a05300eac515a26bc2550ac6ffe4981308c3877f4257092e07da5"></a>

## Direct properties — api_token.blindfold_secret_info / 80f8bcb30d4b / 3

<a id="canonical-4ade364d2fe9364e88b11418e8f629c630cfb134d2fb0076871f830861643185"></a>

<a id="canonical-2c14d71128bb8395e35ae19ba19d3a214105ab56afb9d6b5d59d6e6edc5488b9"></a>

## decryption_provider property — api_token.blindfold_secret_info / 80f8bcb30d4b / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c8e1bcd03b3c4d6dcc9b7e062b69260983941593c5ceebc5f8c5f6a303c69890"></a>

<a id="canonical-1cbb4fe578f3a5598afca4dc11f3768ee02420fe7e50170f36cc7f8db6f793ff"></a>

## location property — api_token.blindfold_secret_info / 80f8bcb30d4b / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2bac30fc85cbcfae4c64c21a61629e4f83cc3a7b86f1c76d10261c78df1352b2"></a>

<a id="canonical-611be60830e0709806939890915b3fb99ffa11d5ae020ecf836db7e8a6f33b62"></a>

## store_provider property — api_token.blindfold_secret_info / 80f8bcb30d4b / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0fad6724c7b995c59867d3a47dc0288c758587797e6c42814b53a3df2f1dd4b8"></a>

## Next pages — api_token.blindfold_secret_info / 80f8bcb30d4b / 7

- [api_token](data-sources--cminstance--reference--group-001.md#canonical-9a17ed398623aea6931d85b571015840e4143eed629ba2b501d96fd345f64b4d)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)

<a id="canonical-cf29c528bfe2852f37fe27d25fc84f04c6d0412ee6d4840049af826bafcf6438"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99f3b159d4d0756a5d04be659a1761b1ebfa98ba901b8708e008d22fd882a281"></a>

## api_token.clear_secret_info — api_token.clear_secret_info / 638b4ca8c62d / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2)
- [api_token](data-sources--cminstance--reference--group-001.md#canonical-9a17ed398623aea6931d85b571015840e4143eed629ba2b501d96fd345f64b4d)
- api_token.clear_secret_info

<a id="canonical-7a11c48e475f7f1f4a1a349ec27d757896cab5afeec13aaed0e3e785583843b4"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-70b65956d730de4bbd21bfcc7e5b6910b70a25668e19d48eafee19865625c482"></a>

## Direct properties — api_token.clear_secret_info / 638b4ca8c62d / 3

<a id="canonical-4addf483c354bec45d8d5dd41a94068319d93d88cdd8f58efd6b7a854d518026"></a>

<a id="canonical-183400ddca228e68a325eaf32bf7b09746486797193fe55905c323c8fb8ed7ac"></a>

## provider_ref property — api_token.clear_secret_info / 638b4ca8c62d / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-623093d6f4ac2a68b08bc3ee5d1488e731d23a96f614c1ed7d3498947e00db50"></a>

<a id="canonical-3e88d36bd010b8e490d845eb176276256fc4b9b993fc2ee7a458347c51570202"></a>

## url property — api_token.clear_secret_info / 638b4ca8c62d / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-12ed45c908535438b089bce70431c9aeea190d627bf101ba7746ecc84fda59a9"></a>

## Next pages — api_token.clear_secret_info / 638b4ca8c62d / 6

- [api_token](data-sources--cminstance--reference--group-001.md#canonical-9a17ed398623aea6931d85b571015840e4143eed629ba2b501d96fd345f64b4d)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)

<a id="canonical-3c6339205e271a123677ab965f68ca38a30fdfe92908accda8a304262e0a1455"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd95f7e9e3e72349b271e6e6b831ac645e0dc4a4744ab172a2f0ee11c8ca902e"></a>

## ip — ip / 47eb7651f09a / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2)
- ip

<a id="canonical-dabe639ec17c62500a4d7e03fbc18362ec202f49484ef0eaff518f3a66e078e8"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

<a id="canonical-983af264fb2e54a6ed3ab6c68217fa9b9eb9120794ad254a31b2b5f37fc6c422"></a>

## Direct properties — ip / 47eb7651f09a / 3

<a id="canonical-8d86873aae7619b7789106d38166893f0465e17417c3fd3c7ac652422f80f22d"></a>

<a id="canonical-e1e12281ec18d670d89adcdbff868bbd345d1a71b0aaedbba827609dab6bda9b"></a>

## addr property — ip / 47eb7651f09a / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-675427ff35d562d24e0236bfeff7bc8393c39456b604ce32fe05a8b1de5c3315"></a>

## Next pages — ip / 47eb7651f09a / 5

- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)

<a id="canonical-7a57d86ff7cc215f8bda1c753c142ff62b525be0049e23794b698a0d27e0c4d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99e241f9b6366bdce50cda81b8e4d665a8d3f5f2b1b34e6e06f150026ef801c0"></a>

## password — password / 8bce5cf5c098 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2)
- password

<a id="canonical-b8ad4060eb3dcc8640eeec174993a559d7911942fd7e01421ba7f3930199e0ff"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-300179fc43109113166b089da03d0f3be88ed6c24b478ba280e777553c5653d6"></a>

## Direct properties — password / 8bce5cf5c098 / 3

- [blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-4457a30bb379a23a6929fe8f1ed86462a5604f27c79c403ce4380a1a3b9775d3): complete subsection reference.

- [clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-0e5c1d627300c0dcd2e4700614d0c0abddfb563e9de319c8e75dfaf84167c070): complete subsection reference.

<a id="canonical-e2b3918f6de441bd9e8a4b28fff39f538fb5df9d44f9c472db3a534eb867ed7e"></a>

## Next pages — password / 8bce5cf5c098 / 4

- [password.blindfold_secret_info](data-sources--cminstance--reference--group-001.md#canonical-4457a30bb379a23a6929fe8f1ed86462a5604f27c79c403ce4380a1a3b9775d3)
- [password.clear_secret_info](data-sources--cminstance--reference--group-001.md#canonical-0e5c1d627300c0dcd2e4700614d0c0abddfb563e9de319c8e75dfaf84167c070)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)

<a id="canonical-4457a30bb379a23a6929fe8f1ed86462a5604f27c79c403ce4380a1a3b9775d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08b44fc1513082cda7c3c974889342805fe0db1632926b480353a1e8d4322cda"></a>

## password.blindfold_secret_info — password.blindfold_secret_info / e4a73c2eeffd / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2)
- [password](data-sources--cminstance--reference--group-001.md#canonical-7a57d86ff7cc215f8bda1c753c142ff62b525be0049e23794b698a0d27e0c4d9)
- password.blindfold_secret_info

<a id="canonical-94b435540e17097ac1f44fe06b257f32b7b2c3b331dfe5072aeeaee26aa06054"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-f74f539d662498f596944928d4bf5ab4a8973effe3b98567e406d4f5d96385e5"></a>

## Direct properties — password.blindfold_secret_info / e4a73c2eeffd / 3

<a id="canonical-008a826e6aab8bd718274d742bd4cf4b0006ffebf6b837b958ee29f12e34d10f"></a>

<a id="canonical-895f04327017a9e9bba364151e8cb3b5bdd546af8bc8af80440d31171ad18c51"></a>

## decryption_provider property — password.blindfold_secret_info / e4a73c2eeffd / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-8fc7e35c35934d370a6bf4212db5354df5a08a2a31495425b79f257031b8a8dd"></a>

<a id="canonical-d39799bc12b5601eb46d054835eddea1a6a26c3a3fc7b4a6313dacf4378ec763"></a>

## location property — password.blindfold_secret_info / e4a73c2eeffd / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-598949b28bea46354cd42b4fbb845030957d4593a3a086a60706b83618c9481f"></a>

<a id="canonical-0b7fb8432499d695db72a943e332c80da4af7601da917a6ee909e5d14c681e69"></a>

## store_provider property — password.blindfold_secret_info / e4a73c2eeffd / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b72c2e36f9aa764dca73eda78652a1e3d57d8af55786082acf3d20523cc6b2a4"></a>

## Next pages — password.blindfold_secret_info / e4a73c2eeffd / 7

- [password](data-sources--cminstance--reference--group-001.md#canonical-7a57d86ff7cc215f8bda1c753c142ff62b525be0049e23794b698a0d27e0c4d9)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)

<a id="canonical-0e5c1d627300c0dcd2e4700614d0c0abddfb563e9de319c8e75dfaf84167c070"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5aa0442c6a2470c5cefa645aae8f537a337b1048454d81201874e2fe465f6a3d"></a>

## password.clear_secret_info — password.clear_secret_info / fa35fb609e09 / 2

Breadcrumbs:

- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
- [Property reference](data-sources--cminstance--reference--group-001.md#canonical-c9bd3a65c0f26cf5db15191faf5559d6a0edbe97fc6de3a100b3445bac8798e2)
- [password](data-sources--cminstance--reference--group-001.md#canonical-7a57d86ff7cc215f8bda1c753c142ff62b525be0049e23794b698a0d27e0c4d9)
- password.clear_secret_info

<a id="canonical-d3f2eb1ea2dfd1c9675ec79d27e49c2f507140a2c4a8a33114e4e0e52eb7981b"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-c09acc3e050e54561fac4b121445b7b84260f9beb04f0e8cb24b77891d56559c"></a>

## Direct properties — password.clear_secret_info / fa35fb609e09 / 3

<a id="canonical-f22ea42f855dc53f73d004d7150c1291dfe78abaccb613df1dc0accb177abcf5"></a>

<a id="canonical-142639be5352304d6afcadd4ac43b84fb17baae70dd43f95c558bce2caa82af0"></a>

## provider_ref property — password.clear_secret_info / fa35fb609e09 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-c42973aff656d041fe367c8a9c812e7c08fd808bc7aefe0f026df181972bcc96"></a>

<a id="canonical-ca2ad5e668c32bf146dedc0d5051be51db63b77f7e399d4e4731a06c5494f3e5"></a>

## url property — password.clear_secret_info / fa35fb609e09 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-a8f10759d6386fa56511250eae821016e6f94ae02e037eac0590325bd3ce17be"></a>

## Next pages — password.clear_secret_info / fa35fb609e09 / 6

- [password](data-sources--cminstance--reference--group-001.md#canonical-7a57d86ff7cc215f8bda1c753c142ff62b525be0049e23794b698a0d27e0c4d9)
- [xcsh_cminstance](../data-sources/cminstance.md#canonical-02f6c885e68f071b656c6449a3b1dc369f2cba0bd504b6ba060c2f66c5144248)
