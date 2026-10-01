---
page_title: "xcsh_protocol_policer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer reference."
---

# xcsh_protocol_policer reference

<a id="canonical-1c0f9e479b43f472fa9b158ac9ab0882a2445796723024e7566764effc3aa491"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2e678ee8d9e7f9a32278d9728311a0e2424ae27d0874034962b4e876efac4d8"></a>

## Property reference — Property reference / 7506c35428e6 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
- Property reference

<a id="canonical-ea13e93f61a3af9d784b3795f3a0db240e548405d25d258ab68ec7844e34b9a3"></a>

## Direct properties — Property reference / 7506c35428e6 / 3

<a id="canonical-a5bc89af503d9dd0f561454eb373204d052d0961ec37ac99041fa103a2169a1a"></a>

<a id="canonical-d34d180ab96b1624a09c5deac6124e63e99fbaf7506b0543f52e71406eb060a8"></a>

## annotations property — Property reference / 7506c35428e6 / 4

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

<a id="canonical-4aefab21ff6a0cd8cae41f1fde8b0f39ed60abb31d83c5759e2eb218e956b6e2"></a>

<a id="canonical-47976e110a443b021ad3cf6eb934caf3f050fa539eedd99098468cab501a9c7d"></a>

## description property — Property reference / 7506c35428e6 / 5

Type: `"string"`. Computed.

Description of the ProtocolPolicer.

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

<a id="canonical-0c645a663cad79e88c8bb483f261dfc59d8287e53a4c9ad5cdd9193c30b545db"></a>

<a id="canonical-db705494031e5083b6a9ed53c67de86541e6eb2f67122289bed9b1d1a970f152"></a>

## id property — Property reference / 7506c35428e6 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-947fca13e64688786b763ff3127acaff1d8485cc96526a7ee79d1b1509fd96f7"></a>

<a id="canonical-a76140c04ac952977bc442df667b00981685c539737059ca1d60ae8fccdae5cd"></a>

## labels property — Property reference / 7506c35428e6 / 7

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

<a id="canonical-990ab8ab3313eef8b16f160cf3b74b7a6aac6dbc32506df9d1cd7944dbdf7cda"></a>

<a id="canonical-5dc8c402d39f428fba5c2edefda02b50f15bb76806f162b612350b780d570378"></a>

## name property — Property reference / 7506c35428e6 / 8

Type: `"string"`. Required.

Name of the ProtocolPolicer.

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

<a id="canonical-f0442cf85f2c659ef1f46cad9a260d977804b6de9e478190851725e83f1b92d2"></a>

<a id="canonical-631fab452a8c9bb52c6e509d27a2f4f75c61dc3df865e9d5fd2febd6baca7885"></a>

## namespace property — Property reference / 7506c35428e6 / 9

Type: `"string"`. Optional, Computed.

Namespace where the ProtocolPolicer exists.

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

- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-511cd6275f7e9626e7978434d9e9cc5431c00db328674c4952e5eb754d16f44d): complete subsection reference.

<a id="canonical-46980980ed4f8cdcf2af15cbd48805fb831270c4016cdcc78df28efe950ec2bf"></a>

## All schema paths — Property reference / 7506c35428e6 / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--protocol_policer--reference--group-001.md#canonical-a5bc89af503d9dd0f561454eb373204d052d0961ec37ac99041fa103a2169a1a) |
| `description` | [description](data-sources--protocol_policer--reference--group-001.md#canonical-4aefab21ff6a0cd8cae41f1fde8b0f39ed60abb31d83c5759e2eb218e956b6e2) |
| `id` | [id](data-sources--protocol_policer--reference--group-001.md#canonical-0c645a663cad79e88c8bb483f261dfc59d8287e53a4c9ad5cdd9193c30b545db) |
| `labels` | [labels](data-sources--protocol_policer--reference--group-001.md#canonical-947fca13e64688786b763ff3127acaff1d8485cc96526a7ee79d1b1509fd96f7) |
| `name` | [name](data-sources--protocol_policer--reference--group-001.md#canonical-990ab8ab3313eef8b16f160cf3b74b7a6aac6dbc32506df9d1cd7944dbdf7cda) |
| `namespace` | [namespace](data-sources--protocol_policer--reference--group-001.md#canonical-f0442cf85f2c659ef1f46cad9a260d977804b6de9e478190851725e83f1b92d2) |
| `protocol_policer` | [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-6e9e88b4d4d358bc0b69816c2dcee957c1d1633fd8860e809e35c0d00a654597) |
| `protocol_policer.policer` | [protocol_policer.policer](data-sources--protocol_policer--reference--group-001.md#canonical-ba8c623e961967d73f5e8524f19f509d63609a5665f74bf3b62905e307c7d81d) |
| `protocol_policer.policer.kind` | [protocol_policer.policer.kind](data-sources--protocol_policer--reference--group-001.md#canonical-3d67f687071de4ac05f38833bee7f0b73e89055e9c4877b3e5076a7cea4e19b0) |
| `protocol_policer.policer.name` | [protocol_policer.policer.name](data-sources--protocol_policer--reference--group-001.md#canonical-8feef6a94f96ce104d31e73a07842ca9988a406d20cd37a408e96462a38a3ec8) |
| `protocol_policer.policer.namespace` | [protocol_policer.policer.namespace](data-sources--protocol_policer--reference--group-001.md#canonical-db535a29103ecb41ec312a58581795761f8e763b712c2e94bd33db792ab3d60d) |
| `protocol_policer.policer.tenant` | [protocol_policer.policer.tenant](data-sources--protocol_policer--reference--group-001.md#canonical-4b80b9b0799ff7dac9f2543e302d108645d6a44d41484f59371964c8af0c1ebd) |
| `protocol_policer.policer.uid` | [protocol_policer.policer.uid](data-sources--protocol_policer--reference--group-001.md#canonical-243501c2f07205c078dc80415dc3f08c5d70354e3cc6ac7f566215c95b47c94a) |
| `protocol_policer.protocol` | [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-ca86ec9d972f63994be3aa4a76a169489d950a54a920d044cf0c0acf984c5e4f) |
| `protocol_policer.protocol.dns` | [protocol_policer.protocol.dns](data-sources--protocol_policer--reference--group-001.md#canonical-0b78bae95938bfe4bf7dd5c8d4a317bc82e45ab8aba7bbb551ed383c44297720) |
| `protocol_policer.protocol.icmp` | [protocol_policer.protocol.icmp](data-sources--protocol_policer--reference--group-001.md#canonical-06c444443c24dce01d0ea493cb07e4efa5fa04a871c0f5c911355e0d680d5b7f) |
| `protocol_policer.protocol.icmp.type` | [protocol_policer.protocol.icmp.type](data-sources--protocol_policer--reference--group-001.md#canonical-7630b53bac092b884e6b5b19b4bc2c376fcf42781eec9afed5adb2461273510a) |
| `protocol_policer.protocol.tcp` | [protocol_policer.protocol.tcp](data-sources--protocol_policer--reference--group-001.md#canonical-566a5d445ccddae3def2c2786c3263085d16d254fa8dbbfdfec49bfc6c3d629c) |
| `protocol_policer.protocol.tcp.flags` | [protocol_policer.protocol.tcp.flags](data-sources--protocol_policer--reference--group-001.md#canonical-4b452f55ae5221c096eedbd8809a62c0c7a6e634e72c2b14425d1d29c8deb1a4) |
| `protocol_policer.protocol.udp` | [protocol_policer.protocol.udp](data-sources--protocol_policer--reference--group-001.md#canonical-f723b7bf11b0babc47e4bad5802f45d29192c79b9adaef18b0679563ef08ec24) |

<a id="canonical-403b49e1ab4a7e32c6dd95ea84ec3b530144372e372b64ab9e4a8cd708f0289a"></a>

## Next pages — Property reference / 7506c35428e6 / 11

- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-511cd6275f7e9626e7978434d9e9cc5431c00db328674c4952e5eb754d16f44d)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)

<a id="canonical-511cd6275f7e9626e7978434d9e9cc5431c00db328674c4952e5eb754d16f44d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f5cb73150bab5c5443af7a75dc38b5e8c1fe27a223d0261b4295333a9dd66cd"></a>

## protocol_policer — protocol_policer / af7b4c940adc / 2

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-1c0f9e479b43f472fa9b158ac9ab0882a2445796723024e7566764effc3aa491)
- protocol_policer

<a id="canonical-6e9e88b4d4d358bc0b69816c2dcee957c1d1633fd8860e809e35c0d00a654597"></a>

Type: `"list"`. Computed.

List of L4 protocol match condition and associated traffic rate limits.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4fef07454362422fd91264d8d982bececc444a08e4fde51f62d13ad68d4ef074"></a>

## Direct properties — protocol_policer / af7b4c940adc / 3

- [policer](data-sources--protocol_policer--reference--group-001.md#canonical-9eb3f4a33a68822c988cdc9e72c3a5c8eedd580e2aa0319b18b61a18fcd10cd4): complete subsection reference.

- [protocol](data-sources--protocol_policer--reference--group-001.md#canonical-762910771eb373d8085338112563383a5433c06ec873f7aea4624d04da760c9d): complete subsection reference.

<a id="canonical-9b5d2adaf6a8f98076fb4bc9fa6fce8824d9a9509701cdffc9dcafe16c5a408f"></a>

## Next pages — protocol_policer / af7b4c940adc / 4

- [protocol_policer.policer](data-sources--protocol_policer--reference--group-001.md#canonical-9eb3f4a33a68822c988cdc9e72c3a5c8eedd580e2aa0319b18b61a18fcd10cd4)
- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-762910771eb373d8085338112563383a5433c06ec873f7aea4624d04da760c9d)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-1c0f9e479b43f472fa9b158ac9ab0882a2445796723024e7566764effc3aa491)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)

<a id="canonical-9eb3f4a33a68822c988cdc9e72c3a5c8eedd580e2aa0319b18b61a18fcd10cd4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e16b9b35516bf012bade6219b47b53c794a5891e6d1a56c64f7d01e8a37215c1"></a>

## protocol_policer.policer — protocol_policer.policer / 6ef4414ee67e / 2

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-1c0f9e479b43f472fa9b158ac9ab0882a2445796723024e7566764effc3aa491)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-511cd6275f7e9626e7978434d9e9cc5431c00db328674c4952e5eb754d16f44d)
- protocol_policer.policer

<a id="canonical-ba8c623e961967d73f5e8524f19f509d63609a5665f74bf3b62905e307c7d81d"></a>

Type: `"list"`. Computed.

Reference to policer object to apply traffic rate limits.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-bfae37acf1412289e6c3e8dd58ccd6a75080594c96bd4121a2899ba32cb9b9b1"></a>

## Direct properties — protocol_policer.policer / 6ef4414ee67e / 3

<a id="canonical-3d67f687071de4ac05f38833bee7f0b73e89055e9c4877b3e5076a7cea4e19b0"></a>

<a id="canonical-b8e653cbb20532801dfc3561917e42ae0320a2ddd3d2894a5d6f7e24e64eee4b"></a>

## kind property — protocol_policer.policer / 6ef4414ee67e / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-8feef6a94f96ce104d31e73a07842ca9988a406d20cd37a408e96462a38a3ec8"></a>

<a id="canonical-b28274e0c94b9d849265c3e379df4528771c0f37ef8804c4c4a49856f432737b"></a>

## name property — protocol_policer.policer / 6ef4414ee67e / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-db535a29103ecb41ec312a58581795761f8e763b712c2e94bd33db792ab3d60d"></a>

<a id="canonical-7b19ce19dda9931b78ca6aa724f5c22ee55478828c57efd8569d1755028a57ac"></a>

## namespace property — protocol_policer.policer / 6ef4414ee67e / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-4b80b9b0799ff7dac9f2543e302d108645d6a44d41484f59371964c8af0c1ebd"></a>

<a id="canonical-278ff799e1b5803e8b9331ecc317206bdbedac818a8d803602db56699f94d8eb"></a>

## tenant property — protocol_policer.policer / 6ef4414ee67e / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-243501c2f07205c078dc80415dc3f08c5d70354e3cc6ac7f566215c95b47c94a"></a>

<a id="canonical-2c4222242844d4ea3135fa8e1c847a865d7c544bbcb040e5d07fffaa9dbfcecb"></a>

## uid property — protocol_policer.policer / 6ef4414ee67e / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-b7153773964de9fbeaac461f98d74a09d8d74f7c300baa5189447e8d0d672d6b"></a>

## Next pages — protocol_policer.policer / 6ef4414ee67e / 9

- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-511cd6275f7e9626e7978434d9e9cc5431c00db328674c4952e5eb754d16f44d)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)

<a id="canonical-762910771eb373d8085338112563383a5433c06ec873f7aea4624d04da760c9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a2d64941d9c8f486e050e902b781ce33743b6fd79015b3218ff1ac28daa3e85"></a>

## protocol_policer.protocol — protocol_policer.protocol / 21edde5acbe7 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-1c0f9e479b43f472fa9b158ac9ab0882a2445796723024e7566764effc3aa491)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-511cd6275f7e9626e7978434d9e9cc5431c00db328674c4952e5eb754d16f44d)
- protocol_policer.protocol

<a id="canonical-ca86ec9d972f63994be3aa4a76a169489d950a54a920d044cf0c0acf984c5e4f"></a>

Type: `"single"`. Computed.

Protocol and protocol specific flags to be matched in packet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-type": "[\"dns\",\"icmp\",\"tcp\",\"udp\"]"
}
```

<a id="canonical-c0ecbc04e444f6e3904f3d6e53086681c0fec1062b857b4cbd71927ffdff2589"></a>

## Direct properties — protocol_policer.protocol / 21edde5acbe7 / 3

- [dns](data-sources--protocol_policer--reference--group-001.md#canonical-d2c6a4ac7a2cc1fda341d2ed668f9bb510e13a0e389c6186ea61c6b9ee3329e5): complete subsection reference.

- [icmp](data-sources--protocol_policer--reference--group-001.md#canonical-c481ecaef34c9637396589f2a27b67e3abfe7da92edf8f014e092cbbd16cfd61): complete subsection reference.

- [tcp](data-sources--protocol_policer--reference--group-001.md#canonical-c1245a83b905bf86d27d63c4e69ea94c99dfee0e27d79b65f699d7a85b44f6bc): complete subsection reference.

- [udp](data-sources--protocol_policer--reference--group-001.md#canonical-7c54ffd0034ea139913c84483157e946ed00598dbed1008e47974a7fb4e6d41c): complete subsection reference.

<a id="canonical-ed049b90ac5c4132f7d4868449c04d798a4ede1577fe4fa3b73a730e9c711839"></a>

## Next pages — protocol_policer.protocol / 21edde5acbe7 / 4

- [protocol_policer.protocol.dns](data-sources--protocol_policer--reference--group-001.md#canonical-d2c6a4ac7a2cc1fda341d2ed668f9bb510e13a0e389c6186ea61c6b9ee3329e5)
- [protocol_policer.protocol.icmp](data-sources--protocol_policer--reference--group-001.md#canonical-c481ecaef34c9637396589f2a27b67e3abfe7da92edf8f014e092cbbd16cfd61)
- [protocol_policer.protocol.tcp](data-sources--protocol_policer--reference--group-001.md#canonical-c1245a83b905bf86d27d63c4e69ea94c99dfee0e27d79b65f699d7a85b44f6bc)
- [protocol_policer.protocol.udp](data-sources--protocol_policer--reference--group-001.md#canonical-7c54ffd0034ea139913c84483157e946ed00598dbed1008e47974a7fb4e6d41c)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-511cd6275f7e9626e7978434d9e9cc5431c00db328674c4952e5eb754d16f44d)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)

<a id="canonical-d2c6a4ac7a2cc1fda341d2ed668f9bb510e13a0e389c6186ea61c6b9ee3329e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3f4621ac735b8c27f6bf98541568d33d879ef215550e2c119ef7e34914ad1ae"></a>

## protocol_policer.protocol.dns — protocol_policer.protocol.dns / 4a36710df205 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-1c0f9e479b43f472fa9b158ac9ab0882a2445796723024e7566764effc3aa491)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-511cd6275f7e9626e7978434d9e9cc5431c00db328674c4952e5eb754d16f44d)
- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-762910771eb373d8085338112563383a5433c06ec873f7aea4624d04da760c9d)
- protocol_policer.protocol.dns

<a id="canonical-0b78bae95938bfe4bf7dd5c8d4a317bc82e45ab8aba7bbb551ed383c44297720"></a>

Type: `["object", {}]`. Computed.

Match all DNS packets including UDP and TCP.

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

<a id="canonical-2fcc33b3f0891fc4cc2723c062e50761358efa4a1a922e6dfa5fcebd3aae99a9"></a>

## Direct properties — protocol_policer.protocol.dns / 4a36710df205 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12e8f3e2c9a60cdf72bd0a4a02032860bda2cbe4397633c83feeb6b0df497ba7"></a>

## Next pages — protocol_policer.protocol.dns / 4a36710df205 / 4

- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-762910771eb373d8085338112563383a5433c06ec873f7aea4624d04da760c9d)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)

<a id="canonical-c481ecaef34c9637396589f2a27b67e3abfe7da92edf8f014e092cbbd16cfd61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72a626bc80da47c5e2e03a1e1d964d42348e601a723a46902377727068941c5a"></a>

## protocol_policer.protocol.icmp — protocol_policer.protocol.icmp / 3a5cacd8f0da / 2

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-1c0f9e479b43f472fa9b158ac9ab0882a2445796723024e7566764effc3aa491)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-511cd6275f7e9626e7978434d9e9cc5431c00db328674c4952e5eb754d16f44d)
- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-762910771eb373d8085338112563383a5433c06ec873f7aea4624d04da760c9d)
- protocol_policer.protocol.icmp

<a id="canonical-06c444443c24dce01d0ea493cb07e4efa5fa04a871c0f5c911355e0d680d5b7f"></a>

Type: `"single"`. Computed.

ICMP Packet Type. ICMP message type to match in packet.

Upstream description:

ICMP message type to match in packet.

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

<a id="canonical-a2180c5dc1b9da78a55adef2cb30907ce4e70bd9dd470517138936392b76f5b1"></a>

## Direct properties — protocol_policer.protocol.icmp / 3a5cacd8f0da / 3

<a id="canonical-7630b53bac092b884e6b5b19b4bc2c376fcf42781eec9afed5adb2461273510a"></a>

<a id="canonical-9ead4d69feb7c1dc912c744ea3d0c99085ab3b386365068a4e264bf18021d794"></a>

## type property — protocol_policer.protocol.icmp / 3a5cacd8f0da / 4

Type: `["list", "string"]`. Computed.

\[Enum: ECHO\_REPLY|ECHO\_REQUEST|ALL\_ICMP\_MSG\] ICMP message type to be matched in packet.
Possible values are \`ECHO\_REPLY\`, \`ECHO\_REQUEST\`, \`ALL\_ICMP\_MSG\`. Defaults to
\`ECHO\_REPLY\`.

Upstream description:

ICMP message type to be matched in packet.

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

<a id="canonical-8b51677993a0bcff8817268f0f956280350fd8ebd330b4f1e1f5c2a47b4b26c2"></a>

## Next pages — protocol_policer.protocol.icmp / 3a5cacd8f0da / 5

- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-762910771eb373d8085338112563383a5433c06ec873f7aea4624d04da760c9d)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)

<a id="canonical-c1245a83b905bf86d27d63c4e69ea94c99dfee0e27d79b65f699d7a85b44f6bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bafa5ffad489ebbdda60ebcce0da3a252f1c0514963bc803d9ff5a776121d913"></a>

## protocol_policer.protocol.tcp — protocol_policer.protocol.tcp / fa75cab46891 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-1c0f9e479b43f472fa9b158ac9ab0882a2445796723024e7566764effc3aa491)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-511cd6275f7e9626e7978434d9e9cc5431c00db328674c4952e5eb754d16f44d)
- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-762910771eb373d8085338112563383a5433c06ec873f7aea4624d04da760c9d)
- protocol_policer.protocol.tcp

<a id="canonical-566a5d445ccddae3def2c2786c3263085d16d254fa8dbbfdfec49bfc6c3d629c"></a>

Type: `"single"`. Computed.

Specification of TCP flag to be matched in a TCP packet.

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

<a id="canonical-874f92f607492f0e38308451743768aa4a4b0dc5e0fbefa58c747c36843d8b40"></a>

## Direct properties — protocol_policer.protocol.tcp / fa75cab46891 / 3

<a id="canonical-4b452f55ae5221c096eedbd8809a62c0c7a6e634e72c2b14425d1d29c8deb1a4"></a>

<a id="canonical-bd01b598ddadda44cb2b3febd2e24af8f59eb2eb08e94ec020015a94048a735e"></a>

## flags property — protocol_policer.protocol.tcp / fa75cab46891 / 4

Type: `["list", "string"]`. Computed.

\[Enum: FIN|SYN|RST|PSH|ACK|URG|ALL\_TCP\_FLAGS|KEEPALIVE\] TCP flags. TCP flag to be matched in a
TCP packet. Possible values are \`FIN\`, \`SYN\`, \`RST\`, \`PSH\`, \`ACK\`, \`URG\`,
\`ALL\_TCP\_FLAGS\`, \`KEEPALIVE\`. Defaults to \`FIN\`.

Upstream description:

TCP flag to be matched in a TCP packet.

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

<a id="canonical-a0948faa1640de28ce3205eb04ddf1b6ccf4936a680da15a6be9587e9cc92617"></a>

## Next pages — protocol_policer.protocol.tcp / fa75cab46891 / 5

- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-762910771eb373d8085338112563383a5433c06ec873f7aea4624d04da760c9d)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)

<a id="canonical-7c54ffd0034ea139913c84483157e946ed00598dbed1008e47974a7fb4e6d41c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba9287004b4384651742ef418572150dfd587d6e148d703d311180a11df7cd52"></a>

## protocol_policer.protocol.udp — protocol_policer.protocol.udp / ae0f95248c68 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
- [Property reference](data-sources--protocol_policer--reference--group-001.md#canonical-1c0f9e479b43f472fa9b158ac9ab0882a2445796723024e7566764effc3aa491)
- [protocol_policer](data-sources--protocol_policer--reference--group-001.md#canonical-511cd6275f7e9626e7978434d9e9cc5431c00db328674c4952e5eb754d16f44d)
- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-762910771eb373d8085338112563383a5433c06ec873f7aea4624d04da760c9d)
- protocol_policer.protocol.udp

<a id="canonical-f723b7bf11b0babc47e4bad5802f45d29192c79b9adaef18b0679563ef08ec24"></a>

Type: `["object", {}]`. Computed.

UDP Packets. Match all UDP packets.

Upstream description:

Match all UDP packets.

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

<a id="canonical-41b117a6420e5bfc038e5efd9dbe40aa5f1f0dacc38db3d55a9ae91bccbd57ba"></a>

## Direct properties — protocol_policer.protocol.udp / ae0f95248c68 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bedab064be149db252eac37bcc6a99914b66da4ac90f3f7600114010436200fb"></a>

## Next pages — protocol_policer.protocol.udp / ae0f95248c68 / 4

- [protocol_policer.protocol](data-sources--protocol_policer--reference--group-001.md#canonical-762910771eb373d8085338112563383a5433c06ec873f7aea4624d04da760c9d)
- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-3fe015d3b1ba6ff09581d00d9baa87149cb8e458009e6dbcf6d2b038c3b7946e)
