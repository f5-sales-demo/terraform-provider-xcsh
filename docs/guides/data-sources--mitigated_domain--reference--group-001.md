---
page_title: "xcsh_mitigated_domain reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_mitigated_domain reference."
---

# xcsh_mitigated_domain reference

<a id="canonical-219ebdfe9e83d72a62a01c844ad18ac7c83ab7629b5a58cda760c069bc6d72e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ada03621cad2145c02022d4ca1a70f03580cf8b1b1b366d71a2ed587c0a2e91"></a>

## Property reference — Property reference / 74ce534668ac / 2

Breadcrumbs:

- [xcsh_mitigated_domain](../data-sources/mitigated_domain.md#canonical-efcfbe08e10eae8cd34e645079b3b2ddf4a333bb259bfc7f7c6325c8868a5fb4)
- Property reference

<a id="canonical-2abc8be4b89b680c3e5a493f2e5581f555b74743c866b847fbd84bcbcccb2001"></a>

## Direct properties — Property reference / 74ce534668ac / 3

<a id="canonical-4a1c6c3c62ba457c5e4b743140d276413ee17cb4f17b495db21599e381660928"></a>

<a id="canonical-e091cf84c0a4a6672c207f841788c0b61e33a79d501dc71bb2788d1973dc18cf"></a>

## annotations property — Property reference / 74ce534668ac / 4

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

<a id="canonical-1dbad280082146527016d46c74038b01019347b0b7216c3bd0abff60afa89957"></a>

<a id="canonical-c5e8b3e1b6a2eddb0996b68c9db00eadbbd8b698c8c48cd3136ebd1a3fa8d10e"></a>

## description property — Property reference / 74ce534668ac / 5

Type: `"string"`. Computed.

Description of the MitigatedDomain.

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

<a id="canonical-375d015d3abaf23c798c80fc68887d2bb9cc32c1040a5fa84d0e75d0ad89f09a"></a>

<a id="canonical-ff20a12a08b54687e398f8d8813ac108d851024d30d89ace3905a58bfc2ffbc8"></a>

## id property — Property reference / 74ce534668ac / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-7325fd0cca7924e6e0aa0cd37f4fcd5669eea2596ca320c20b97fd77a69eec18"></a>

<a id="canonical-62fd108fba9c8edc38bdfa8cc7d2dd197e8da7e07aa9a9faeb966c91895e1fcc"></a>

## labels property — Property reference / 74ce534668ac / 7

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

<a id="canonical-21a6d18d43356c9b009b1b1a79296a66a477ae0ca66f2706c825aa7a191df6e9"></a>

<a id="canonical-7b6ae52e6e6a4203ce66111693bc552472056f6902a1dc962834cbaa7084e921"></a>

## mitigated_domain property — Property reference / 74ce534668ac / 8

Type: `"string"`. Computed.

Enter root domain or domain to be entered to mitigated list below. Domains can be entered only one
at a time. In case of conflicting entries, the domain entry takes precedence over the root domain
entry.

Upstream description:

Enter root domain or domain to be entered to mitigated list below. Domains can be entered only one
at a time. In case of conflicting entries, the domain entry takes precedence over the root domain
entry. Example: if you are adding Client-Side Defense JS on checkout.example.com, you should enter
example.com here.

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

<a id="canonical-0babdfc5a4d99b9a455a0e24ce49d34b9955055f04c3a1a0b43b89b95162d1c1"></a>

<a id="canonical-bbfaaa8c96446771c2797dde161f0a38ef0ded075a0fdc07f634dda355594104"></a>

## name property — Property reference / 74ce534668ac / 9

Type: `"string"`. Required.

Name of the MitigatedDomain.

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

<a id="canonical-145cb663da282c148501503a09a030bb786411e3c08ce2d499396b65b873b7b6"></a>

<a id="canonical-aa9ac780c731f881700469467b7b3abdc13c6a7355ab4ae62e00933b312da7d1"></a>

## namespace property — Property reference / 74ce534668ac / 10

Type: `"string"`. Required.

Namespace where the MitigatedDomain exists.

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

<a id="canonical-70f946be494775938457e7e4ab76e76f36842e161c389e124b17d0e7fc346c7c"></a>

## All schema paths — Property reference / 74ce534668ac / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--mitigated_domain--reference--group-001.md#canonical-4a1c6c3c62ba457c5e4b743140d276413ee17cb4f17b495db21599e381660928) |
| `description` | [description](data-sources--mitigated_domain--reference--group-001.md#canonical-1dbad280082146527016d46c74038b01019347b0b7216c3bd0abff60afa89957) |
| `id` | [id](data-sources--mitigated_domain--reference--group-001.md#canonical-375d015d3abaf23c798c80fc68887d2bb9cc32c1040a5fa84d0e75d0ad89f09a) |
| `labels` | [labels](data-sources--mitigated_domain--reference--group-001.md#canonical-7325fd0cca7924e6e0aa0cd37f4fcd5669eea2596ca320c20b97fd77a69eec18) |
| `mitigated_domain` | [mitigated_domain](data-sources--mitigated_domain--reference--group-001.md#canonical-21a6d18d43356c9b009b1b1a79296a66a477ae0ca66f2706c825aa7a191df6e9) |
| `name` | [name](data-sources--mitigated_domain--reference--group-001.md#canonical-0babdfc5a4d99b9a455a0e24ce49d34b9955055f04c3a1a0b43b89b95162d1c1) |
| `namespace` | [namespace](data-sources--mitigated_domain--reference--group-001.md#canonical-145cb663da282c148501503a09a030bb786411e3c08ce2d499396b65b873b7b6) |

<a id="canonical-b8f2f80e51fc89d092ea36602d4dc7d921c7374fc4c9710a48358b185b760a06"></a>

## Next pages — Property reference / 74ce534668ac / 12

- [xcsh_mitigated_domain](../data-sources/mitigated_domain.md#canonical-efcfbe08e10eae8cd34e645079b3b2ddf4a333bb259bfc7f7c6325c8868a5fb4)
