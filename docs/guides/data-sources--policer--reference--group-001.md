---
page_title: "xcsh_policer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_policer reference."
---

# xcsh_policer reference

<a id="canonical-f87bd9bfc4cb7f2e29ae21826a8d294799efb0c555b9ca771f74fa96c2470bc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06a365cfec83a1a65fb2fe8deb2678fc953cb714f6b6951a8a185b45e9d01960"></a>

## Property reference — Property reference / b220cc78694f / 2

Breadcrumbs:

- [xcsh_policer](../data-sources/policer.md#canonical-fc38c1c976cffbbb7d1001c95e37ff2eb2f3d9753461e9f18f25729eae4994c0)
- Property reference

<a id="canonical-89873f8cfc9d90fcfc5c03ee97820803bbb7a550c3751429dfbd873016b9253d"></a>

## Direct properties — Property reference / b220cc78694f / 3

<a id="canonical-85d0b7a80ccea94c058e068f199a4835a5e2a58f19799363a274068071ec02cd"></a>

<a id="canonical-49f6830266ce467579b585ace6fee133e1590e059950003c7d54e0a9594c26f6"></a>

## annotations property — Property reference / b220cc78694f / 4

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

<a id="canonical-9d2a55ff68c7d912bdac5a1c75a6e0994bfb21a84af5ce115aef5c7c3d83a598"></a>

<a id="canonical-2a0a2a853185fa48e90f434a7d4618082beb03ae93813cedc8989890f9f67bbb"></a>

## burst_size property — Property reference / b220cc78694f / 5

Type: `"number"`. Computed.

The maximum size permitted for bursts of data. E.g. 10000 pps burst.

Upstream description:

The maximum size permitted for bursts of data. E.g. 10000 pps burst.

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

<a id="canonical-41717a09bdffbbbd9bff14cea4aee4e456153e36b5d3448a09058a3a2cef8996"></a>

<a id="canonical-e87f9f075c26d9032177283b065de11f6c344d8f8f130133af9cbc3a4ced8c40"></a>

## committed_information_rate property — Property reference / b220cc78694f / 6

Type: `"number"`. Computed.

The committed information rate is the guaranteed packets rate for traffic arriving or departing
under normal conditions. E.g. 10000 pps.

Upstream description:

The committed information rate is the guaranteed packets rate for traffic arriving or departing
under normal conditions. E.g. 10000 pps.

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

<a id="canonical-cb8b54ee43f2ad4bcd0888749ebe919a7215ebe85b06ce8d2bc3652958ff8ae4"></a>

<a id="canonical-3e2cb7dfd05ce019af53cc0c25a0abb2471ff83ce1b8b0a59c31a783ca7df065"></a>

## description property — Property reference / b220cc78694f / 7

Type: `"string"`. Computed.

Description of the Policer.

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

<a id="canonical-7cc5d7b637fd4b453ff8de5de60b5c9ad95d190c9d8f0503c1671e5b8b9036c9"></a>

<a id="canonical-864fe35f15eef4ae4df2f0c971d0ca0997327cbc2f30d58a6ad71ad3d07587eb"></a>

## id property — Property reference / b220cc78694f / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-7efb8801a8b348fb23fa0085310b1507af657cb3392c46d9d56dc4a19a755654"></a>

<a id="canonical-d37fc0a241e333682e4500f43547d2be857e1fea9bca5d9bc1a6986b456a0bb4"></a>

## labels property — Property reference / b220cc78694f / 9

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

<a id="canonical-6eda0f2d8cd86ac11d8ce6f93217d945ae1c5d06656cf89c9d58a7e8b757d433"></a>

<a id="canonical-9b4eb8bec8ac2f03e12e5e02c96c8810386fda0c6f6c54b62cfed3d50bf807da"></a>

## name property — Property reference / b220cc78694f / 10

Type: `"string"`. Required.

Name of the Policer.

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

<a id="canonical-78a9406fd5285f4a48e7ace35d0cd51d585810dec3c738d2ec3e9119199d879c"></a>

<a id="canonical-252db8ff4eed327e0ca223e50d064d99579f6a6212167f8c0ff011551a6e5516"></a>

## namespace property — Property reference / b220cc78694f / 11

Type: `"string"`. Required.

Namespace where the Policer exists.

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

<a id="canonical-e4854a6f1799afdbc64b63ef5ee4694549cb2af15e1cc9da71fa39217582542c"></a>

<a id="canonical-48fde3fdea4a9151c89fb99ef8188d732e240b7c09e0ed8fca5024bfdecb98d9"></a>

## policer_mode property — Property reference / b220cc78694f / 12

Type: `"string"`. Computed.

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

<a id="canonical-21c4b4da3336091bf754803498398016a5c3c1ca970c364211febdf29403030d"></a>

<a id="canonical-29998fbb407b36f5b3fa745ad01236ddb8a1379a603b525b09b183fd92583baa"></a>

## policer_type property — Property reference / b220cc78694f / 13

Type: `"string"`. Computed.

\[Enum: POLICER\_SINGLE\_RATE\_TWO\_COLOR\] Specifies the type of Policer Basic Single-Rate
Two-Color Policer. The only possible value is \`POLICER\_SINGLE\_RATE\_TWO\_COLOR\`. Defaults to
\`POLICER\_SINGLE\_RATE\_TWO\_COLOR\`. Server applies default when omitted.

Upstream description:

Specifies the type of Policer

Basic Single-Rate Two-Color Policer.

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

<a id="canonical-f4e0227d5ddf6ef3241fe1293703b7d689476e5ef6e6f8cb4bb593c0f950df9e"></a>

## All schema paths — Property reference / b220cc78694f / 14

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--policer--reference--group-001.md#canonical-85d0b7a80ccea94c058e068f199a4835a5e2a58f19799363a274068071ec02cd) |
| `burst_size` | [burst_size](data-sources--policer--reference--group-001.md#canonical-9d2a55ff68c7d912bdac5a1c75a6e0994bfb21a84af5ce115aef5c7c3d83a598) |
| `committed_information_rate` | [committed_information_rate](data-sources--policer--reference--group-001.md#canonical-41717a09bdffbbbd9bff14cea4aee4e456153e36b5d3448a09058a3a2cef8996) |
| `description` | [description](data-sources--policer--reference--group-001.md#canonical-cb8b54ee43f2ad4bcd0888749ebe919a7215ebe85b06ce8d2bc3652958ff8ae4) |
| `id` | [id](data-sources--policer--reference--group-001.md#canonical-7cc5d7b637fd4b453ff8de5de60b5c9ad95d190c9d8f0503c1671e5b8b9036c9) |
| `labels` | [labels](data-sources--policer--reference--group-001.md#canonical-7efb8801a8b348fb23fa0085310b1507af657cb3392c46d9d56dc4a19a755654) |
| `name` | [name](data-sources--policer--reference--group-001.md#canonical-6eda0f2d8cd86ac11d8ce6f93217d945ae1c5d06656cf89c9d58a7e8b757d433) |
| `namespace` | [namespace](data-sources--policer--reference--group-001.md#canonical-78a9406fd5285f4a48e7ace35d0cd51d585810dec3c738d2ec3e9119199d879c) |
| `policer_mode` | [policer_mode](data-sources--policer--reference--group-001.md#canonical-e4854a6f1799afdbc64b63ef5ee4694549cb2af15e1cc9da71fa39217582542c) |
| `policer_type` | [policer_type](data-sources--policer--reference--group-001.md#canonical-21c4b4da3336091bf754803498398016a5c3c1ca970c364211febdf29403030d) |

<a id="canonical-86f801301bb50006b4086b5bbf8a2f1ff05d1db30625b14b1484ca0804bfdda9"></a>

## Next pages — Property reference / b220cc78694f / 15

- [xcsh_policer](../data-sources/policer.md#canonical-fc38c1c976cffbbb7d1001c95e37ff2eb2f3d9753461e9f18f25729eae4994c0)
