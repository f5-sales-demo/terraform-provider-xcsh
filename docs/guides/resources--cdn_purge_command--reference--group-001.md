---
page_title: "xcsh_cdn_purge_command reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command reference."
---

# xcsh_cdn_purge_command reference

<a id="canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89e027a49dd18ce8b39d78f3070dd4bff1417f7a33a9874c6c50aefce45bef11"></a>

## Property reference — Property reference / 296314c137d3 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)
- Property reference

<a id="canonical-18694503b6a001a646756a386c849084f76f3caeafe0fa4510a5e006c77697a3"></a>

## Direct properties — Property reference / 296314c137d3 / 3

<a id="canonical-801d2d9c089d8d62f223f59101d8f852e8175ce8cfb7450a633389708b95e9d1"></a>

<a id="canonical-f9cd724da4067ba18ba0caf28f772ea312f8f8463fc891cf141cc020310f4d7d"></a>

## annotations property — Property reference / 296314c137d3 / 4

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

<a id="canonical-10b7ba4bdd589ccb3ae6e8aad8c8382ee18380f8dc80a55a3453c174b10495fb"></a>

<a id="canonical-939be940b349b5ed7547a64e1b0f5acc0d881010e27332f55bf1b3cb9ee6663e"></a>

## description property — Property reference / 296314c137d3 / 5

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

<a id="canonical-a48052a7b68f17c6acd0d9b4b30e0a9d679178c1347124657908950fe1e39de2"></a>

<a id="canonical-a86f768164bfc65951ecc0ed603228bddbbb66bc98e7930cb1ca75a0ed0c0f3d"></a>

## disable property — Property reference / 296314c137d3 / 6

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

- [hard_purge](resources--cdn_purge_command--reference--group-001.md#canonical-5d3c91eb91f1ed335dcd3bbac4422c1eaac871be08b55d492c95de6e69a64db3): complete subsection reference.

<a id="canonical-5d84d8ff502a798d4815328114a143b54f0a225964615cbf144d473427fbf906"></a>

<a id="canonical-95b1a2fa52045bcb2c403107c063bae711d49a69849b6e4790bb2a3f80e1ac37"></a>

## hostname property — Property reference / 296314c137d3 / 7

Type: `"string"`. Optional, Computed.

\[OneOf: hostname, pattern, purge\_all, url\_path\] Exclusive with \[pattern purge\_all url\_path\]
Purge cached content by Hostname.

Upstream description:

Exclusive with \[pattern purge\_all url\_path\] Purge cached content by Hostname.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

OneOf alternatives in this subsection:

- [hostname](resources--cdn_purge_command--reference--group-001.md#canonical-5d84d8ff502a798d4815328114a143b54f0a225964615cbf144d473427fbf906)
- [pattern](resources--cdn_purge_command--reference--group-001.md#canonical-138ce5488ecb46fe887e244104445b5c725c648aa70300a580af6a3fc7be3e71)
- [purge_all](resources--cdn_purge_command--reference--group-001.md#canonical-5204ee45c2a1abc18259abddc74a11a041142221de348627321b9b93ff3d609e)
- [url_path](resources--cdn_purge_command--reference--group-001.md#canonical-68ef86c711231e6c495380c13c6267fd6c16f347af6648b5d03f9fbefe47a94e)

Select alternatives according to the provider validators above.

<a id="canonical-7c55ad5e7fd4a8bf79292da4fdfa7c08253839ad5bdeebd0cdfaf722e46a13cd"></a>

<a id="canonical-b230a70c0710a8bd92dd8659f153242937a5c76e695cb4dc914770943ca2717e"></a>

## id property — Property reference / 296314c137d3 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-c4af412044ee07141328309953e090f07d802ab76f8bfa11862a3a477e2de689"></a>

<a id="canonical-ce9dfdacf6cccd912d3960ea3fd9c6235725729ad855249407cc344bb5c316ce"></a>

## labels property — Property reference / 296314c137d3 / 9

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

<a id="canonical-4c57ee1879ebe0caf7471307b148c50fb0f44c99f6f4da4ee9862261713ad4fb"></a>

<a id="canonical-d0fad5d4e73202a1ec7245d4f325227fa3371a6d9a5b1ca872c6adc3b749b94a"></a>

## name property — Property reference / 296314c137d3 / 10

Type: `"string"`. Required.

Name of the CDN Purge Command. Must be unique within the namespace.

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

<a id="canonical-bbe9ed0bbdad7aca955967b52e2882245dc8789cfea192a8911bbc7ac06410dd"></a>

<a id="canonical-bb9dca4fb8cd39c969eac9e01dfe71a48c92f170534114eaf4840b03fa0dfb08"></a>

## namespace property — Property reference / 296314c137d3 / 11

Type: `"string"`. Required.

Namespace where the CDN Purge Command is created.

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

<a id="canonical-138ce5488ecb46fe887e244104445b5c725c648aa70300a580af6a3fc7be3e71"></a>

<a id="canonical-4c4a4f65516e66edaf510753085e88ccaab405af5ba4faae0ec73d4448ac3016"></a>

## pattern property — Property reference / 296314c137d3 / 12

Type: `"string"`. Optional, Computed.

Exclusive with \[hostname purge\_all url\_path\] Purge cached content using PCRE 1 compliant regular
expression.

Upstream description:

Exclusive with \[hostname purge\_all url\_path\] Purge cached content using PCRE 1 compliant regular
expression.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [purge_all](resources--cdn_purge_command--reference--group-001.md#canonical-5be2c5509e838a2588b7201d762276ea397beae0ae527bb02cd0a19deb7119ab): complete subsection reference.

- [soft_purge](resources--cdn_purge_command--reference--group-001.md#canonical-d3925bc0dfd59dcb541fada2504630cd30530731405b054c46a361088de48746): complete subsection reference.

- [timeouts](resources--cdn_purge_command--reference--group-001.md#canonical-eef4dff2562f1da55e129e79de7875759900b6f01d4d0c312f3576415159f622): complete subsection reference.

<a id="canonical-68ef86c711231e6c495380c13c6267fd6c16f347af6648b5d03f9fbefe47a94e"></a>

<a id="canonical-d0981dd06b395767f0322a64d5273c1c56d07cbd70c46e99d66cea91bfe14610"></a>

## url_path property — Property reference / 296314c137d3 / 13

Type: `"string"`. Optional, Computed.

Exclusive with \[hostname pattern purge\_all\] Purge cache by using a URL path.

Upstream description:

Exclusive with \[hostname pattern purge\_all\] Purge cache by using a URL path.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [virtual_host](resources--cdn_purge_command--reference--group-001.md#canonical-9238c774537ba0bd7d4161904d5d29ff0eb7d8fac8e58f578d8906260e14ced2): complete subsection reference.

<a id="canonical-bdb44ae5939d379ea4f43433849555e05c157598fcbb528901a6b66e523aa9e5"></a>

## All schema paths — Property reference / 296314c137d3 / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cdn_purge_command--reference--group-001.md#canonical-801d2d9c089d8d62f223f59101d8f852e8175ce8cfb7450a633389708b95e9d1) |
| `description` | [description](resources--cdn_purge_command--reference--group-001.md#canonical-10b7ba4bdd589ccb3ae6e8aad8c8382ee18380f8dc80a55a3453c174b10495fb) |
| `disable` | [disable](resources--cdn_purge_command--reference--group-001.md#canonical-a48052a7b68f17c6acd0d9b4b30e0a9d679178c1347124657908950fe1e39de2) |
| `hard_purge` | [hard_purge](resources--cdn_purge_command--reference--group-001.md#canonical-89c875d826e2df5828703163265b726e09a9b7b8f00a025024a5401196f16032) |
| `hostname` | [hostname](resources--cdn_purge_command--reference--group-001.md#canonical-5d84d8ff502a798d4815328114a143b54f0a225964615cbf144d473427fbf906) |
| `id` | [id](resources--cdn_purge_command--reference--group-001.md#canonical-7c55ad5e7fd4a8bf79292da4fdfa7c08253839ad5bdeebd0cdfaf722e46a13cd) |
| `labels` | [labels](resources--cdn_purge_command--reference--group-001.md#canonical-c4af412044ee07141328309953e090f07d802ab76f8bfa11862a3a477e2de689) |
| `name` | [name](resources--cdn_purge_command--reference--group-001.md#canonical-4c57ee1879ebe0caf7471307b148c50fb0f44c99f6f4da4ee9862261713ad4fb) |
| `namespace` | [namespace](resources--cdn_purge_command--reference--group-001.md#canonical-bbe9ed0bbdad7aca955967b52e2882245dc8789cfea192a8911bbc7ac06410dd) |
| `pattern` | [pattern](resources--cdn_purge_command--reference--group-001.md#canonical-138ce5488ecb46fe887e244104445b5c725c648aa70300a580af6a3fc7be3e71) |
| `purge_all` | [purge_all](resources--cdn_purge_command--reference--group-001.md#canonical-5204ee45c2a1abc18259abddc74a11a041142221de348627321b9b93ff3d609e) |
| `soft_purge` | [soft_purge](resources--cdn_purge_command--reference--group-001.md#canonical-e78ace71bc0b6d57304baadb37f48b58c1eab2b7f7d493d761169cf78b4475d4) |
| `timeouts` | [timeouts](resources--cdn_purge_command--reference--group-001.md#canonical-b993d9aa75565991318795497a86c6d97e4631cd0a9e92e3370bce265c97714b) |
| `timeouts.create` | [timeouts.create](resources--cdn_purge_command--reference--group-001.md#canonical-feb5a745ebd3435ae77254760a90309ffb124bb752222644d9a9472bec53f172) |
| `timeouts.delete` | [timeouts.delete](resources--cdn_purge_command--reference--group-001.md#canonical-bbd8819a30b63792e735d2975c97863c2717536a07979d38b3a2b184ebcf2cd5) |
| `timeouts.read` | [timeouts.read](resources--cdn_purge_command--reference--group-001.md#canonical-eb6d62fb33d2bd9a8a48d46014637a0cc2098421a996444e490877e96d1d2c62) |
| `timeouts.update` | [timeouts.update](resources--cdn_purge_command--reference--group-001.md#canonical-7851bd32447b75402005368a27b20bd1da0b9302585762085ac778789ceb99ac) |
| `url_path` | [url_path](resources--cdn_purge_command--reference--group-001.md#canonical-68ef86c711231e6c495380c13c6267fd6c16f347af6648b5d03f9fbefe47a94e) |
| `virtual_host` | [virtual_host](resources--cdn_purge_command--reference--group-001.md#canonical-15e7d1583cec68e96bd9f39e318389e06ccaab44cccd59b2da8516bf8ed097dd) |
| `virtual_host.name` | [virtual_host.name](resources--cdn_purge_command--reference--group-001.md#canonical-6a3262cc5d6e02bacbddbbe272ece4e0313058bfa69f72e7ca7e9dcc087d9a41) |
| `virtual_host.namespace` | [virtual_host.namespace](resources--cdn_purge_command--reference--group-001.md#canonical-9c774c33a282fc52bd02b652003e68633f7cadce36a5a993fb8850f8b1097e03) |
| `virtual_host.tenant` | [virtual_host.tenant](resources--cdn_purge_command--reference--group-001.md#canonical-ea122cf73e7a879a986ec6262f46e83118d720b72f1b8984a80d02c6554c52a6) |

<a id="canonical-b5c75fd7e34bbaefe70d657677294fe0c4ed553d823e42633f3cb14dec36be0b"></a>

## Next pages — Property reference / 296314c137d3 / 15

- [hard_purge](resources--cdn_purge_command--reference--group-001.md#canonical-5d3c91eb91f1ed335dcd3bbac4422c1eaac871be08b55d492c95de6e69a64db3)
- [purge_all](resources--cdn_purge_command--reference--group-001.md#canonical-5be2c5509e838a2588b7201d762276ea397beae0ae527bb02cd0a19deb7119ab)
- [soft_purge](resources--cdn_purge_command--reference--group-001.md#canonical-d3925bc0dfd59dcb541fada2504630cd30530731405b054c46a361088de48746)
- [timeouts](resources--cdn_purge_command--reference--group-001.md#canonical-eef4dff2562f1da55e129e79de7875759900b6f01d4d0c312f3576415159f622)
- [virtual_host](resources--cdn_purge_command--reference--group-001.md#canonical-9238c774537ba0bd7d4161904d5d29ff0eb7d8fac8e58f578d8906260e14ced2)
- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)

<a id="canonical-5d3c91eb91f1ed335dcd3bbac4422c1eaac871be08b55d492c95de6e69a64db3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b2cbba350dbd2907f62f69804254bbf589614c4b2867c08777d2b0172f8fdfc"></a>

## hard_purge — hard_purge / 0ab839e6817c / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)
- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32)
- hard_purge

<a id="canonical-89c875d826e2df5828703163265b726e09a9b7b8f00a025024a5401196f16032"></a>

Type: `["object", {}]`. Optional.

\[OneOf: hard\_purge, soft\_purge\] Enable this option

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

OneOf alternatives in this subsection:

- [hard_purge](resources--cdn_purge_command--reference--group-001.md#canonical-89c875d826e2df5828703163265b726e09a9b7b8f00a025024a5401196f16032)
- [soft_purge](resources--cdn_purge_command--reference--group-001.md#canonical-e78ace71bc0b6d57304baadb37f48b58c1eab2b7f7d493d761169cf78b4475d4)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
hard_purge = {}
```

<a id="canonical-a47aac88f64a503a376c7f5e3d1a9dd48e48cb83033b1b9664bbd7b2e865525f"></a>

## Direct properties — hard_purge / 0ab839e6817c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34709e91e622c024ddfc41f43b9db04c194716761ba505b3163299692df1877a"></a>

## Next pages — hard_purge / 0ab839e6817c / 4

- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32)
- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)

<a id="canonical-5be2c5509e838a2588b7201d762276ea397beae0ae527bb02cd0a19deb7119ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-325f815ddb164abef06a12a4e8519b8e136afd8b88016803f305d7effc0c8cd7"></a>

## purge_all — purge_all / c0aeb3243441 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)
- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32)
- purge_all

<a id="canonical-5204ee45c2a1abc18259abddc74a11a041142221de348627321b9b93ff3d609e"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
purge_all = {}
```

<a id="canonical-60a0d5c352efd1946e620ac78bef14bf2a384ed09ddd13882a9d7a57d93e608a"></a>

## Direct properties — purge_all / c0aeb3243441 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86dcc09ffea5f6b54653d8619944999c46086b0b4c38084996f2be83b87843db"></a>

## Next pages — purge_all / c0aeb3243441 / 4

- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32)
- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)

<a id="canonical-d3925bc0dfd59dcb541fada2504630cd30530731405b054c46a361088de48746"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9f073b2cd02abb4ca3b644a09e54c38f6690091dd8b088aec66e7ff0e936716"></a>

## soft_purge — soft_purge / 9a8093facca0 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)
- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32)
- soft_purge

<a id="canonical-e78ace71bc0b6d57304baadb37f48b58c1eab2b7f7d493d761169cf78b4475d4"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
soft_purge = {}
```

<a id="canonical-2a2e2c52ff66b3b42ecb37ed1a86f65a4543e3b214a72d342cba8b19e305195c"></a>

## Direct properties — soft_purge / 9a8093facca0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7c9a2306f3368c901c8a292cb699e35abdf71faeab1b180f2ac66ce0f60ae1d5"></a>

## Next pages — soft_purge / 9a8093facca0 / 4

- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32)
- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)

<a id="canonical-eef4dff2562f1da55e129e79de7875759900b6f01d4d0c312f3576415159f622"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0e250d3ca159b5866e30a3b0b4d90c9b4af2423b8222d7c397835be23219e74"></a>

## timeouts — timeouts / 237217e227c4 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)
- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32)
- timeouts

<a id="canonical-b993d9aa75565991318795497a86c6d97e4631cd0a9e92e3370bce265c97714b"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1009c75b88cf35c3e14997a8be05b3cfeb2d9fb0b84d774418467002d6cd5c60"></a>

## Direct properties — timeouts / 237217e227c4 / 3

<a id="canonical-feb5a745ebd3435ae77254760a90309ffb124bb752222644d9a9472bec53f172"></a>

<a id="canonical-993170f87c64c668d9c50f672a1d36538d73119e98aee698379eaf8ba39060a2"></a>

## create property — timeouts / 237217e227c4 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-bbd8819a30b63792e735d2975c97863c2717536a07979d38b3a2b184ebcf2cd5"></a>

<a id="canonical-0a2ae01f6d49014a695b8e34e7130129695c7344475aa6527ff86e2163cb862d"></a>

## delete property — timeouts / 237217e227c4 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-eb6d62fb33d2bd9a8a48d46014637a0cc2098421a996444e490877e96d1d2c62"></a>

<a id="canonical-1e5c2848a41398663dea43343048ab63f5f23836bfefd76b45186a3fe9f5aca3"></a>

## read property — timeouts / 237217e227c4 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-7851bd32447b75402005368a27b20bd1da0b9302585762085ac778789ceb99ac"></a>

<a id="canonical-c2c6c9844cbc7c3d2ae7ffabb36ad38fe42d8b71fb0be2bb0b000a8215196bde"></a>

## update property — timeouts / 237217e227c4 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-4916cc3ac0ade416b08462d891970cc9f97d7bcd55547aa30b079bb6311c1203"></a>

## Next pages — timeouts / 237217e227c4 / 8

- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32)
- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)

<a id="canonical-9238c774537ba0bd7d4161904d5d29ff0eb7d8fac8e58f578d8906260e14ced2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c500ed5cb2e8ebffae877ecc98ef8232f1bd217e79a87a17b144d13768011d6"></a>

## virtual_host — virtual_host / 9f16f2175303 / 2

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)
- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32)
- virtual_host

<a id="canonical-15e7d1583cec68e96bd9f39e318389e06ccaab44cccd59b2da8516bf8ed097dd"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
virtual_host {
  # Configure direct properties listed below.
}
```

<a id="canonical-bedc41d647f2cbc9bd15d415c3cb004dbfd364d96832430adc0798361467c38a"></a>

## Direct properties — virtual_host / 9f16f2175303 / 3

<a id="canonical-6a3262cc5d6e02bacbddbbe272ece4e0313058bfa69f72e7ca7e9dcc087d9a41"></a>

<a id="canonical-80030125689cb7dd6780b33d37bc26da897dcb7ce6cfc8c727ee050c6f5d3b67"></a>

## name property — virtual_host / 9f16f2175303 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-9c774c33a282fc52bd02b652003e68633f7cadce36a5a993fb8850f8b1097e03"></a>

<a id="canonical-ce475c165b0aba281f8928aff6c41328a04da6cfd33c3bc255a8ef705e051fa0"></a>

## namespace property — virtual_host / 9f16f2175303 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-ea122cf73e7a879a986ec6262f46e83118d720b72f1b8984a80d02c6554c52a6"></a>

<a id="canonical-20bacd4d167e4f21a60218af2c4d52558b1e86e21d865072105bc824783ea410"></a>

## tenant property — virtual_host / 9f16f2175303 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-cbcc8027f14c0626d57239943cc82096735185f58a250781752a35b62d3acfc9"></a>

## Next pages — virtual_host / 9f16f2175303 / 7

- [Property reference](resources--cdn_purge_command--reference--group-001.md#canonical-0d7ff38d62b870c2da94865b1a7e4f05b5f53d124a92ef6bcc79c52b5a816d32)
- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0b2a857dd2a4345307f28b391793f77c7f11e893b8493fb53ba1965b5e2bf425)
