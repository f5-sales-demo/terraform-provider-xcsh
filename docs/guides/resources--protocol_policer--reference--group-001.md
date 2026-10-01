---
page_title: "xcsh_protocol_policer reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer reference."
---

# xcsh_protocol_policer reference

<a id="canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8fceb7a8b9be7bb879ab09d9a8b3e45ca3dd0f6f92d7ea2f1d71532976c2c63"></a>

## Property reference — Property reference / a634d9b73082 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
- Property reference

<a id="canonical-99e2fa105d3cfaa8ca987c5fc61f5964778e6299cfef4868b344b2a66d8f1852"></a>

## Direct properties — Property reference / a634d9b73082 / 3

<a id="canonical-aac1e85a19b7fe3e1a2fc42b0fab09faf6cbfac6edfd68b3919d302982cd6579"></a>

<a id="canonical-17a2654691b29bdfafc367ca96f3faacbe40ba42dd3f06e1650d6a26a8fe06c6"></a>

## annotations property — Property reference / a634d9b73082 / 4

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

<a id="canonical-7e1fe8654aeb0c27b0e0fda3e5586bd319dc81186a9968f81cc2506dacd972d6"></a>

<a id="canonical-a93e1d0366832df0e27ef6b366bda7a7e82240e651b08ece098a7a3d4bb685f8"></a>

## description property — Property reference / a634d9b73082 / 5

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

<a id="canonical-42f1659821780e92485963227e3cb11cbf4a5368ee2e0607ec3c9caec008703c"></a>

<a id="canonical-30e826b1d0d45461f5c1c770f610d94f05719bfb34b3bc03ab2defafd1525fe1"></a>

## disable property — Property reference / a634d9b73082 / 6

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

<a id="canonical-a0eb272295c0af53240b6ef567f7739145b8573d352882caf52a9c9077fc11fc"></a>

<a id="canonical-d4e0ebc8f6934f13eb226ec3f997c9d3d35ad804f57ee4eac8c55490f3fdf70f"></a>

## id property — Property reference / a634d9b73082 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-6416b37f958c1685cec949fcc67b1544c1d3ce3ef0e7a717a7867493a9660383"></a>

<a id="canonical-92233b55aa424034985d0171721000d10a5364de69bf10156a8a86828527e71f"></a>

## labels property — Property reference / a634d9b73082 / 8

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

<a id="canonical-a99a407a59826eff3544c1fa07ac35ddd20bf2cb93313030e771d8446b1d3451"></a>

<a id="canonical-26cd1a381d8c6406cb8af3a10a9fdc2c7f60a68ef89d1b59cc83d32362db32ce"></a>

## name property — Property reference / a634d9b73082 / 9

Type: `"string"`. Required.

Name of the Protocol Policer. Must be unique within the namespace.

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

<a id="canonical-2f8b8f741a64ed85a94e326b744dfafb17d8594ca184acb277d5626706a88043"></a>

<a id="canonical-b3f1109daf4d74520dc059a507173e9f45c1c60082a78770f4d0724486965e30"></a>

## namespace property — Property reference / a634d9b73082 / 10

Type: `"string"`. Optional, Computed.

Namespace for the Protocol Policer. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-538fc767eec1d120767ed70495f597a3dd5563a39e02bcb92a45129745b5139e): complete subsection reference.

- [timeouts](resources--protocol_policer--reference--group-001.md#canonical-d2d3a60fbee3c4b79c03a42a6b0f89da16556073955126d8ab56209e4cc49566): complete subsection reference.

<a id="canonical-bf0a2bdef7f16398a1f73eefa0fada4222adf8ff8f2941e5340329ee3f83d4ef"></a>

## All schema paths — Property reference / a634d9b73082 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--protocol_policer--reference--group-001.md#canonical-aac1e85a19b7fe3e1a2fc42b0fab09faf6cbfac6edfd68b3919d302982cd6579) |
| `description` | [description](resources--protocol_policer--reference--group-001.md#canonical-7e1fe8654aeb0c27b0e0fda3e5586bd319dc81186a9968f81cc2506dacd972d6) |
| `disable` | [disable](resources--protocol_policer--reference--group-001.md#canonical-42f1659821780e92485963227e3cb11cbf4a5368ee2e0607ec3c9caec008703c) |
| `id` | [id](resources--protocol_policer--reference--group-001.md#canonical-a0eb272295c0af53240b6ef567f7739145b8573d352882caf52a9c9077fc11fc) |
| `labels` | [labels](resources--protocol_policer--reference--group-001.md#canonical-6416b37f958c1685cec949fcc67b1544c1d3ce3ef0e7a717a7867493a9660383) |
| `name` | [name](resources--protocol_policer--reference--group-001.md#canonical-a99a407a59826eff3544c1fa07ac35ddd20bf2cb93313030e771d8446b1d3451) |
| `namespace` | [namespace](resources--protocol_policer--reference--group-001.md#canonical-2f8b8f741a64ed85a94e326b744dfafb17d8594ca184acb277d5626706a88043) |
| `protocol_policer` | [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-f8d3b8743f036c31b5768994412ec62aea9ea2ebb374ca0cb0a3e136d1ca361c) |
| `protocol_policer.policer` | [protocol_policer.policer](resources--protocol_policer--reference--group-001.md#canonical-63dc216af102b6416a226212f8d62751930de8b48a07a43fa289fccfce0817dd) |
| `protocol_policer.policer.kind` | [protocol_policer.policer.kind](resources--protocol_policer--reference--group-001.md#canonical-d7f3e52d26074ddb715333b41efeb7d6e68957b69cb43976ecff8bf59d0121e6) |
| `protocol_policer.policer.name` | [protocol_policer.policer.name](resources--protocol_policer--reference--group-001.md#canonical-55cc336301b8fc4eb3d18096ac7323d06c7f4876881eb4a3379cdef84d595f6f) |
| `protocol_policer.policer.namespace` | [protocol_policer.policer.namespace](resources--protocol_policer--reference--group-001.md#canonical-1ac297c5a2507a3bb6ed133280281608eabe61a61c1f6c749d3d804b8163d786) |
| `protocol_policer.policer.tenant` | [protocol_policer.policer.tenant](resources--protocol_policer--reference--group-001.md#canonical-bedbe496b3c6522bfa7fd5e864f03b9b545d0152621c17455d75fbae7e3cfd76) |
| `protocol_policer.policer.uid` | [protocol_policer.policer.uid](resources--protocol_policer--reference--group-001.md#canonical-22f1764008a836b203e2c38aacd400c5e35eb346ca65750c87a4d50077f2c24e) |
| `protocol_policer.protocol` | [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-09db57493089d27b23d707382c571fef599f3c3ea5cf8ecc63a4abddb8d643a7) |
| `protocol_policer.protocol.dns` | [protocol_policer.protocol.dns](resources--protocol_policer--reference--group-001.md#canonical-0822a114c3f4f6a22292471a62d6e0837f3240956bfaf055f06ab7f76569ffcb) |
| `protocol_policer.protocol.icmp` | [protocol_policer.protocol.icmp](resources--protocol_policer--reference--group-001.md#canonical-7d166873939fd61c801fd8ca5fd1d98d5a1a186934a56052e10c3a9f16a0b921) |
| `protocol_policer.protocol.icmp.type` | [protocol_policer.protocol.icmp.type](resources--protocol_policer--reference--group-001.md#canonical-ad75df862838460d6f04909d34cde18e9fc1c1d85cc918218b7b48525529cb2e) |
| `protocol_policer.protocol.tcp` | [protocol_policer.protocol.tcp](resources--protocol_policer--reference--group-001.md#canonical-9e1bbc4c8c8d5cab14dec616a2639461b167f1c7c918895f58bd219fe3769939) |
| `protocol_policer.protocol.tcp.flags` | [protocol_policer.protocol.tcp.flags](resources--protocol_policer--reference--group-001.md#canonical-c937f4ec1ef0410416bd6c752ae1cfeb61427d641d4727eeb539f60d28ae8c11) |
| `protocol_policer.protocol.udp` | [protocol_policer.protocol.udp](resources--protocol_policer--reference--group-001.md#canonical-8b02a11ba7efe249892b0174af9302432f45115ce01bf371d06ee5c2ad900647) |
| `timeouts` | [timeouts](resources--protocol_policer--reference--group-001.md#canonical-47fe590602f9648efc2267aa5d978d0459db63124a2545c0a478c7ae991d1191) |
| `timeouts.create` | [timeouts.create](resources--protocol_policer--reference--group-001.md#canonical-51501f78b04f9d730dd11bfe659ac2baba52a481155f584d12b3be4e1effa203) |
| `timeouts.delete` | [timeouts.delete](resources--protocol_policer--reference--group-001.md#canonical-7c794a1863ced286ce4a4d61add92c9cd8beb5b3d1d3bc2e1d9db58fc717dba5) |
| `timeouts.read` | [timeouts.read](resources--protocol_policer--reference--group-001.md#canonical-d61f0aa32fa01f02228e1b0ab7e6683a4584e35a235a8f5666550554637c1ebf) |
| `timeouts.update` | [timeouts.update](resources--protocol_policer--reference--group-001.md#canonical-68a778ddb0a250d827d20e4471427af42549e9b219566a7a4937cd69174881b2) |

<a id="canonical-b4cdb2cc64a8faf1e70a36f0f90a33ab321ea20e4a5b43336b36b99fa09f169b"></a>

## Next pages — Property reference / a634d9b73082 / 12

- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-538fc767eec1d120767ed70495f597a3dd5563a39e02bcb92a45129745b5139e)
- [timeouts](resources--protocol_policer--reference--group-001.md#canonical-d2d3a60fbee3c4b79c03a42a6b0f89da16556073955126d8ab56209e4cc49566)
- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)

<a id="canonical-538fc767eec1d120767ed70495f597a3dd5563a39e02bcb92a45129745b5139e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a291fa28843841f83e07d497562d84f2a7057cac2d6af9f0a8b88d1885789ef6"></a>

## protocol_policer — protocol_policer / 37e4ae7034ef / 2

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9)
- protocol_policer

<a id="canonical-f8d3b8743f036c31b5768994412ec62aea9ea2ebb374ca0cb0a3e136d1ca361c"></a>

Type: `"object"`. list nested block, Optional.

List of L4 protocol match condition and associated traffic rate limits.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("policer")}
```

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

Terraform syntax:

```terraform
protocol_policer {
  # Configure direct properties listed below.
}
```

<a id="canonical-8ee2420a1b4d234ecf4adb5a55fc63984c55ea021c403a0007afb88403816882"></a>

## Direct properties — protocol_policer / 37e4ae7034ef / 3

- [policer](resources--protocol_policer--reference--group-001.md#canonical-25667e11c66927533cdc27c50e54e8706bc2fd4823c7a34676a5abf95dc9dc4e): complete subsection reference.

- [protocol](resources--protocol_policer--reference--group-001.md#canonical-50a368815e6caffa5d44ea16a71c09618ef5d3ae2c4a19512bcdd5403e979213): complete subsection reference.

<a id="canonical-8c97ccc2acfd29b1253052b8353d6b84fcb30a6a25e491150a69b60511a8b19f"></a>

## Next pages — protocol_policer / 37e4ae7034ef / 4

- [protocol_policer.policer](resources--protocol_policer--reference--group-001.md#canonical-25667e11c66927533cdc27c50e54e8706bc2fd4823c7a34676a5abf95dc9dc4e)
- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-50a368815e6caffa5d44ea16a71c09618ef5d3ae2c4a19512bcdd5403e979213)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9)
- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)

<a id="canonical-25667e11c66927533cdc27c50e54e8706bc2fd4823c7a34676a5abf95dc9dc4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-684882b2154fe3d4f7ecbc4f5df2402a9b5ab5e091fd7421cd42c46c2b7e6391"></a>

## protocol_policer.policer — protocol_policer.policer / b784c841f216 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-538fc767eec1d120767ed70495f597a3dd5563a39e02bcb92a45129745b5139e)
- protocol_policer.policer

<a id="canonical-63dc216af102b6416a226212f8d62751930de8b48a07a43fa289fccfce0817dd"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
policer {
  # Configure direct properties listed below.
}
```

<a id="canonical-c26658ea2f8e41d2978ff27a9b974e61bb63a86cac3a4fa2b811d6665b4eecc4"></a>

## Direct properties — protocol_policer.policer / b784c841f216 / 3

<a id="canonical-d7f3e52d26074ddb715333b41efeb7d6e68957b69cb43976ecff8bf59d0121e6"></a>

<a id="canonical-dac6411f63e20bb868415d46ffb60829bb44bcb2627c9b3196217f29e20f991f"></a>

## kind property — protocol_policer.policer / b784c841f216 / 4

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

<a id="canonical-55cc336301b8fc4eb3d18096ac7323d06c7f4876881eb4a3379cdef84d595f6f"></a>

<a id="canonical-a3036778c9cfc0eecf7abb0fe1ecca9556db3c310b1d05d911086d70bbe3fee2"></a>

## name property — protocol_policer.policer / b784c841f216 / 5

Type: `"string"`. Optional.

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

<a id="canonical-1ac297c5a2507a3bb6ed133280281608eabe61a61c1f6c749d3d804b8163d786"></a>

<a id="canonical-f070a9b47c03611c27d8ec7670901dcbf01cd7044fbe55168c9bb64e086f19fb"></a>

## namespace property — protocol_policer.policer / b784c841f216 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
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

<a id="canonical-bedbe496b3c6522bfa7fd5e864f03b9b545d0152621c17455d75fbae7e3cfd76"></a>

<a id="canonical-fd730ab7813bec0efb66809bf81be7099b1bde158234b0779c2f050073298da5"></a>

## tenant property — protocol_policer.policer / b784c841f216 / 7

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

<a id="canonical-22f1764008a836b203e2c38aacd400c5e35eb346ca65750c87a4d50077f2c24e"></a>

<a id="canonical-dd2de80d4d01af29e5e2b4768bfaf93ae2bcaaacebfdaf60584854e08d318112"></a>

## uid property — protocol_policer.policer / b784c841f216 / 8

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

<a id="canonical-0a3994afdeb98538a001641f819438301731acf6d2db057e0d86b3839f1dfbaa"></a>

## Next pages — protocol_policer.policer / b784c841f216 / 9

- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-538fc767eec1d120767ed70495f597a3dd5563a39e02bcb92a45129745b5139e)
- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)

<a id="canonical-50a368815e6caffa5d44ea16a71c09618ef5d3ae2c4a19512bcdd5403e979213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee3bfe4bbd922e9f59d4fd69af09ad2fd31dd65758bc2c1981e66bc58f41373f"></a>

## protocol_policer.protocol — protocol_policer.protocol / f80b1618cfc7 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-538fc767eec1d120767ed70495f597a3dd5563a39e02bcb92a45129745b5139e)
- protocol_policer.protocol

<a id="canonical-09db57493089d27b23d707382c571fef599f3c3ea5cf8ecc63a4abddb8d643a7"></a>

Type: `"object"`. single nested block, Optional.

Protocol and protocol specific flags to be matched in packet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dns",
    "icmp"),
  validators.ConflictingObjectAttributes("dns",
    "tcp"),
  validators.ConflictingObjectAttributes("dns",
    "udp"),
  validators.ConflictingObjectAttributes("icmp",
    "tcp"),
  validators.ConflictingObjectAttributes("icmp",
    "udp"),
  validators.ConflictingObjectAttributes("tcp",
    "udp")}
```

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

Terraform syntax:

```terraform
protocol {
  # Configure direct properties listed below.
}
```

<a id="canonical-6459684872371081cea4abbecb850739fb5d584a70fc7f8cc0f51d487b12b12c"></a>

## Direct properties — protocol_policer.protocol / f80b1618cfc7 / 3

- [dns](resources--protocol_policer--reference--group-001.md#canonical-4b4487f8115e996fa2c5c9bf795156d86c2cdf26b9b069d2fee16b24f05adeef): complete subsection reference.

- [icmp](resources--protocol_policer--reference--group-001.md#canonical-c6635f5c54a1a9e39f68e5351f90bc32023b89b21a3396297736475811fd55c1): complete subsection reference.

- [tcp](resources--protocol_policer--reference--group-001.md#canonical-1a20b409080df96b5f3bd94cab8435503c9cd19cab2acb5d5ad8c0e55b0ca385): complete subsection reference.

- [udp](resources--protocol_policer--reference--group-001.md#canonical-ee04708dd8d7ef2ac903320a883b4b2e4ccdd255b6585155e63babcaf9d4459d): complete subsection reference.

<a id="canonical-5b05bfd896811f01b48558d0fe8467affcccb07e6b57bc07247890c5d9ec0bfa"></a>

## Next pages — protocol_policer.protocol / f80b1618cfc7 / 4

- [protocol_policer.protocol.dns](resources--protocol_policer--reference--group-001.md#canonical-4b4487f8115e996fa2c5c9bf795156d86c2cdf26b9b069d2fee16b24f05adeef)
- [protocol_policer.protocol.icmp](resources--protocol_policer--reference--group-001.md#canonical-c6635f5c54a1a9e39f68e5351f90bc32023b89b21a3396297736475811fd55c1)
- [protocol_policer.protocol.tcp](resources--protocol_policer--reference--group-001.md#canonical-1a20b409080df96b5f3bd94cab8435503c9cd19cab2acb5d5ad8c0e55b0ca385)
- [protocol_policer.protocol.udp](resources--protocol_policer--reference--group-001.md#canonical-ee04708dd8d7ef2ac903320a883b4b2e4ccdd255b6585155e63babcaf9d4459d)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-538fc767eec1d120767ed70495f597a3dd5563a39e02bcb92a45129745b5139e)
- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)

<a id="canonical-4b4487f8115e996fa2c5c9bf795156d86c2cdf26b9b069d2fee16b24f05adeef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26890259627ced2dd02d4b8c0ac719c2d82cf709356a9d85f054bdfb9c9638af"></a>

## protocol_policer.protocol.dns — protocol_policer.protocol.dns / 9cc25aec2711 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-538fc767eec1d120767ed70495f597a3dd5563a39e02bcb92a45129745b5139e)
- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-50a368815e6caffa5d44ea16a71c09618ef5d3ae2c4a19512bcdd5403e979213)
- protocol_policer.protocol.dns

<a id="canonical-0822a114c3f4f6a22292471a62d6e0837f3240956bfaf055f06ab7f76569ffcb"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
dns = {}
```

<a id="canonical-9d3a01250e9bf0d9297005b340d737fde75e4e29ca42d7ded7cf0dc4532c1b28"></a>

## Direct properties — protocol_policer.protocol.dns / 9cc25aec2711 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b775ba3dc699bf1e988d48c7e34ed322eb1c8c5f3a90f87aeae0985a3dd3d470"></a>

## Next pages — protocol_policer.protocol.dns / 9cc25aec2711 / 4

- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-50a368815e6caffa5d44ea16a71c09618ef5d3ae2c4a19512bcdd5403e979213)
- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)

<a id="canonical-c6635f5c54a1a9e39f68e5351f90bc32023b89b21a3396297736475811fd55c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c071f9ba294737e796481120e2fe28a7fa06de16c8f213f76d77c28116919ae"></a>

## protocol_policer.protocol.icmp — protocol_policer.protocol.icmp / fd4394db27dc / 2

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-538fc767eec1d120767ed70495f597a3dd5563a39e02bcb92a45129745b5139e)
- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-50a368815e6caffa5d44ea16a71c09618ef5d3ae2c4a19512bcdd5403e979213)
- protocol_policer.protocol.icmp

<a id="canonical-7d166873939fd61c801fd8ca5fd1d98d5a1a186934a56052e10c3a9f16a0b921"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
icmp {
  # Configure direct properties listed below.
}
```

<a id="canonical-2b7a8376a8b8f14ace05631b4155f1b4bb5ac04289e13df1b1d5b6f490dcf92e"></a>

## Direct properties — protocol_policer.protocol.icmp / fd4394db27dc / 3

<a id="canonical-ad75df862838460d6f04909d34cde18e9fc1c1d85cc918218b7b48525529cb2e"></a>

<a id="canonical-244c9f6c061ef805cea4abce489dfdade5894cb3c140e720a2dd5b0a8a7f210b"></a>

## type property — protocol_policer.protocol.icmp / fd4394db27dc / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-facc98afdb7978eca71398747bfeedba5fc3d261fe1443334b790473138cf5b6"></a>

## Next pages — protocol_policer.protocol.icmp / fd4394db27dc / 5

- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-50a368815e6caffa5d44ea16a71c09618ef5d3ae2c4a19512bcdd5403e979213)
- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)

<a id="canonical-1a20b409080df96b5f3bd94cab8435503c9cd19cab2acb5d5ad8c0e55b0ca385"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbaf74cb16bb42c508a5d7e0b1dbb8bb61c9545d0eb1e06bca77c7b061f1237a"></a>

## protocol_policer.protocol.tcp — protocol_policer.protocol.tcp / 33c49167e38c / 2

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-538fc767eec1d120767ed70495f597a3dd5563a39e02bcb92a45129745b5139e)
- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-50a368815e6caffa5d44ea16a71c09618ef5d3ae2c4a19512bcdd5403e979213)
- protocol_policer.protocol.tcp

<a id="canonical-9e1bbc4c8c8d5cab14dec616a2639461b167f1c7c918895f58bd219fe3769939"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-39e8129b5622243777932a11eed200068dca09742688a8e0bcbaf6d8f2cca5c2"></a>

## Direct properties — protocol_policer.protocol.tcp / 33c49167e38c / 3

<a id="canonical-c937f4ec1ef0410416bd6c752ae1cfeb61427d641d4727eeb539f60d28ae8c11"></a>

<a id="canonical-b7b4dedbe9127983dbc34341061038d4200bdde13b6e4770302ee586724163b0"></a>

## flags property — protocol_policer.protocol.tcp / 33c49167e38c / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-f847301f86fa987311f2c89a2161b13f2ce2d0329d7d003eaaee62eaa589e8eb"></a>

## Next pages — protocol_policer.protocol.tcp / 33c49167e38c / 5

- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-50a368815e6caffa5d44ea16a71c09618ef5d3ae2c4a19512bcdd5403e979213)
- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)

<a id="canonical-ee04708dd8d7ef2ac903320a883b4b2e4ccdd255b6585155e63babcaf9d4459d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f492f85020de1082c548fac84a0a412ebdef2e232b68eef05515aab88ebd668"></a>

## protocol_policer.protocol.udp — protocol_policer.protocol.udp / 943332b39c59 / 2

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9)
- [protocol_policer](resources--protocol_policer--reference--group-001.md#canonical-538fc767eec1d120767ed70495f597a3dd5563a39e02bcb92a45129745b5139e)
- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-50a368815e6caffa5d44ea16a71c09618ef5d3ae2c4a19512bcdd5403e979213)
- protocol_policer.protocol.udp

<a id="canonical-8b02a11ba7efe249892b0174af9302432f45115ce01bf371d06ee5c2ad900647"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
udp = {}
```

<a id="canonical-91f71a653276ce05d99f060aa1cda241c8b60f9c2e9b6d0f5cf535780e8d9d84"></a>

## Direct properties — protocol_policer.protocol.udp / 943332b39c59 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8f17579b4c9f8628709500ee3a82891bca781cfa2650d5087004a1cc497eb21e"></a>

## Next pages — protocol_policer.protocol.udp / 943332b39c59 / 4

- [protocol_policer.protocol](resources--protocol_policer--reference--group-001.md#canonical-50a368815e6caffa5d44ea16a71c09618ef5d3ae2c4a19512bcdd5403e979213)
- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)

<a id="canonical-d2d3a60fbee3c4b79c03a42a6b0f89da16556073955126d8ab56209e4cc49566"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-341674f6ff80bd6fd17d27d0cb604d83a3781fd933c647160634d95be9d00c4d"></a>

## timeouts — timeouts / f3d6d3d2bc7b / 2

Breadcrumbs:

- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9)
- timeouts

<a id="canonical-47fe590602f9648efc2267aa5d978d0459db63124a2545c0a478c7ae991d1191"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-d5305d72f2644187691ea4a753bbfdef2ee04487d4f14c9f4f6e52181ef318bf"></a>

## Direct properties — timeouts / f3d6d3d2bc7b / 3

<a id="canonical-51501f78b04f9d730dd11bfe659ac2baba52a481155f584d12b3be4e1effa203"></a>

<a id="canonical-9ea158bb32f8cdd186e731c0cfb8d3edeb4f83377476f6925ad813e1f3bd44d7"></a>

## create property — timeouts / f3d6d3d2bc7b / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-7c794a1863ced286ce4a4d61add92c9cd8beb5b3d1d3bc2e1d9db58fc717dba5"></a>

<a id="canonical-df663e2a5f70b69e402720bde6a5b0f835758ce9c062e2d5b8803fde970d4f66"></a>

## delete property — timeouts / f3d6d3d2bc7b / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-d61f0aa32fa01f02228e1b0ab7e6683a4584e35a235a8f5666550554637c1ebf"></a>

<a id="canonical-0f90d829ed6c0f958da1c3a0ee03214597287477a1504f99bc3ea5756118d2d0"></a>

## read property — timeouts / f3d6d3d2bc7b / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-68a778ddb0a250d827d20e4471427af42549e9b219566a7a4937cd69174881b2"></a>

<a id="canonical-12daed775d8564cb4ea445bb170a671d15848340c7a264c26895bb7ceb88dc04"></a>

## update property — timeouts / f3d6d3d2bc7b / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-5d94eb264fa798a5870813d99136eb2a8e5e1e98a8f11db33f5159eb66e57c7f"></a>

## Next pages — timeouts / f3d6d3d2bc7b / 8

- [Property reference](resources--protocol_policer--reference--group-001.md#canonical-9d82204400732627056272da5be92b804b9f4ef52ee632eb743a54b729452ea9)
- [xcsh_protocol_policer](../resources/protocol_policer.md#canonical-fcfc69cbfd0b4f56641ab603a6ae4460d3f308502cba2e0f15ada8529fd6e6aa)
