---
page_title: "xcsh_certificate reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate reference."
---

# xcsh_certificate reference

<a id="canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c22fb7eb3017ea77f931080b174d4c457fa8a5eaa756832555e81dd7a746957a"></a>

## Property reference — Property reference / bb3ac3a1e5fb / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
- Property reference

<a id="canonical-0dbe2cfab431c66192c71a7930445c0760d55e31d78eac4902529b9ec122a184"></a>

## Direct properties — Property reference / bb3ac3a1e5fb / 3

<a id="canonical-4d3f84b4034d912d2176c4f76231b294ea56b376edecb7523ded4a39ba431df2"></a>

<a id="canonical-0dad3aa8ebbd389fa07a1353e849c3c411382b35ae7a29e0218f002d6e4ae703"></a>

## annotations property — Property reference / bb3ac3a1e5fb / 4

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

- [certificate_chain](data-sources--certificate--reference--group-001.md#canonical-61631daa5438233b7545b456ade1d463cf3cd15d02690b3477e1efc48abc23eb): complete subsection reference.

<a id="canonical-4b0a523147c960c5a3d5492613da18c4903965fb82887898869b969364600981"></a>

<a id="canonical-8c441d5a489a89c4d51febcc5ff609d79878eeb86777db197827a100e44541a9"></a>

## certificate_url property — Property reference / bb3ac3a1e5fb / 5

Type: `"string"`. Computed.

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

Certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--certificate--reference--group-001.md#canonical-33bec69c73df335038d518566f1606ceef701d272089c026de644653e3b89376): complete subsection reference.

<a id="canonical-6c6cbeedbee72c809c3cc874f889d18b104f114e44020c6ce71e21318846df5e"></a>

<a id="canonical-ccbae0fe00399da6a19e386ea372ad79c279fbe4611fca54ec482555728adca1"></a>

## description property — Property reference / bb3ac3a1e5fb / 6

Type: `"string"`. Computed.

Description of the Certificate.

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

- [disable_ocsp_stapling](data-sources--certificate--reference--group-001.md#canonical-73c01fbf9f0d548c650b6a634e929f9a3ec6079fa88f03af929f87efa00223a4): complete subsection reference.

<a id="canonical-b5a1074c1dd9eea4abc5d20db3de486b0deb1aa31ba5e0fbfd8f7f414cc1dca8"></a>

<a id="canonical-bc50e65cd1c05eed151c096b4a59308517d90435d47138215cfca6fc3320a1e4"></a>

## id property — Property reference / bb3ac3a1e5fb / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-7ae4d278a8c733b44c167ec13170573ec00e43c5e124ec15815c66d935c4286a"></a>

<a id="canonical-6dbab8a2b72dd91d9ed5b984bfd6db3bafa602f388ec072294754082a8e3aec0"></a>

## labels property — Property reference / bb3ac3a1e5fb / 8

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

<a id="canonical-473034aa108f90d2586889f825f6014518a8ac2db0db0f2f10d999db21c283d6"></a>

<a id="canonical-dc2c7b429ea31bf780c32654d43d0e7377d94df6eb0478aaf7613fcce1d257b3"></a>

## name property — Property reference / bb3ac3a1e5fb / 9

Type: `"string"`. Required.

Name of the Certificate.

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

<a id="canonical-3cdac903078c8c1bfe9753067b6c17c0ab056bfb96ff151672eb515fa9bf6ae9"></a>

<a id="canonical-8f857212d9885af52171611456011a70992921381089a425a3d39fd0ae4f26c7"></a>

## namespace property — Property reference / bb3ac3a1e5fb / 10

Type: `"string"`. Required.

Namespace where the Certificate exists.

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

- [private_key](data-sources--certificate--reference--group-001.md#canonical-33fb154e59da212c846c109440810f8d8af9779f410f0a4e074cf1737a099299): complete subsection reference.

- [use_system_defaults](data-sources--certificate--reference--group-001.md#canonical-3a0ac3fa04557cf75b2cedeb6e4320d5f76aa77298737d7c5661230470d62eec): complete subsection reference.

<a id="canonical-105eae4f2802e7c170cbec29462fb952065aa46d060bbe12393ae05e53ba7c56"></a>

## All schema paths — Property reference / bb3ac3a1e5fb / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--certificate--reference--group-001.md#canonical-4d3f84b4034d912d2176c4f76231b294ea56b376edecb7523ded4a39ba431df2) |
| `certificate_chain` | [certificate_chain](data-sources--certificate--reference--group-001.md#canonical-c399e16a702b8f759cea6be1de70433fbe5dd9a884cf450f9f7ee36e0792fc5e) |
| `certificate_chain.name` | [certificate_chain.name](data-sources--certificate--reference--group-001.md#canonical-8fddad31f73da5574309881bbe5992499ac503d1ccff19889bd66b2bec2ea803) |
| `certificate_chain.namespace` | [certificate_chain.namespace](data-sources--certificate--reference--group-001.md#canonical-802f1978f394acd547f47dcfc244b90553a2638586ac1582c6f99fa3a39d7e52) |
| `certificate_chain.tenant` | [certificate_chain.tenant](data-sources--certificate--reference--group-001.md#canonical-1bac82308b1403aa8300f5c014551339d86d00be41ce25e854c45aefd545cd63) |
| `certificate_url` | [certificate_url](data-sources--certificate--reference--group-001.md#canonical-4b0a523147c960c5a3d5492613da18c4903965fb82887898869b969364600981) |
| `custom_hash_algorithms` | [custom_hash_algorithms](data-sources--certificate--reference--group-001.md#canonical-c9e08816f6dd4bcbc438488650dee66c5517feed32673d3cbda16ca4bf05ec6f) |
| `custom_hash_algorithms.hash_algorithms` | [custom_hash_algorithms.hash_algorithms](data-sources--certificate--reference--group-001.md#canonical-8caf79b9b5a5d3eb546a64f2c5ea34c493d48196a9b92b08f304c8182b15106f) |
| `description` | [description](data-sources--certificate--reference--group-001.md#canonical-6c6cbeedbee72c809c3cc874f889d18b104f114e44020c6ce71e21318846df5e) |
| `disable_ocsp_stapling` | [disable_ocsp_stapling](data-sources--certificate--reference--group-001.md#canonical-f8ca53706e93f1dae9fa9f2b26f607fb1cffed3303bea613947f0a3b0b590dc5) |
| `id` | [id](data-sources--certificate--reference--group-001.md#canonical-b5a1074c1dd9eea4abc5d20db3de486b0deb1aa31ba5e0fbfd8f7f414cc1dca8) |
| `labels` | [labels](data-sources--certificate--reference--group-001.md#canonical-7ae4d278a8c733b44c167ec13170573ec00e43c5e124ec15815c66d935c4286a) |
| `name` | [name](data-sources--certificate--reference--group-001.md#canonical-473034aa108f90d2586889f825f6014518a8ac2db0db0f2f10d999db21c283d6) |
| `namespace` | [namespace](data-sources--certificate--reference--group-001.md#canonical-3cdac903078c8c1bfe9753067b6c17c0ab056bfb96ff151672eb515fa9bf6ae9) |
| `private_key` | [private_key](data-sources--certificate--reference--group-001.md#canonical-b6b38d47d82666991f2194e90e27ae9482b7b9081bd2270ac4de4c43718cf6e0) |
| `private_key.blindfold_secret_info` | [private_key.blindfold_secret_info](data-sources--certificate--reference--group-001.md#canonical-463325481cf2643b362b3b0f119d669023578a14261491b9f6c7c1234bb87db3) |
| `private_key.blindfold_secret_info.decryption_provider` | [private_key.blindfold_secret_info.decryption_provider](data-sources--certificate--reference--group-001.md#canonical-e4a16af48362524c0ebbf0474a5cc5128f4c77e0919af1d2816ee48a9dc4559e) |
| `private_key.blindfold_secret_info.location` | [private_key.blindfold_secret_info.location](data-sources--certificate--reference--group-001.md#canonical-48ba06e60ed1f065a25486380917c47e2199a305296fc94aa3fcd50e58f8a0de) |
| `private_key.blindfold_secret_info.store_provider` | [private_key.blindfold_secret_info.store_provider](data-sources--certificate--reference--group-001.md#canonical-210e63d8a0ac39933803290e88f8400b6ebb87c561451ed620b7ce1d21d7544c) |
| `private_key.clear_secret_info` | [private_key.clear_secret_info](data-sources--certificate--reference--group-001.md#canonical-2e8229ae91f472f29f0a5208c6a788fc25eb57ee8f1578171e93e75293df6383) |
| `private_key.clear_secret_info.provider_ref` | [private_key.clear_secret_info.provider_ref](data-sources--certificate--reference--group-001.md#canonical-c2a643eee508ad0f989638574065b6fa622308b70b6509c3ce9619a2a1be8389) |
| `private_key.clear_secret_info.url` | [private_key.clear_secret_info.url](data-sources--certificate--reference--group-001.md#canonical-70e5e16bb450abd6fa2e7fa66cf6e5cadfbd82368e52fa44d31c7c61d1166f33) |
| `use_system_defaults` | [use_system_defaults](data-sources--certificate--reference--group-001.md#canonical-c79f92b904f7f7f06d37786e4a9fae271d0da3bc726d640b59326ec5e31f0cae) |

<a id="canonical-d7657b962392118dcc0ecbbd0d4518a8ca439926310a247273604384ae4c7d51"></a>

## Next pages — Property reference / bb3ac3a1e5fb / 12

- [certificate_chain](data-sources--certificate--reference--group-001.md#canonical-61631daa5438233b7545b456ade1d463cf3cd15d02690b3477e1efc48abc23eb)
- [custom_hash_algorithms](data-sources--certificate--reference--group-001.md#canonical-33bec69c73df335038d518566f1606ceef701d272089c026de644653e3b89376)
- [disable_ocsp_stapling](data-sources--certificate--reference--group-001.md#canonical-73c01fbf9f0d548c650b6a634e929f9a3ec6079fa88f03af929f87efa00223a4)
- [private_key](data-sources--certificate--reference--group-001.md#canonical-33fb154e59da212c846c109440810f8d8af9779f410f0a4e074cf1737a099299)
- [use_system_defaults](data-sources--certificate--reference--group-001.md#canonical-3a0ac3fa04557cf75b2cedeb6e4320d5f76aa77298737d7c5661230470d62eec)
- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)

<a id="canonical-61631daa5438233b7545b456ade1d463cf3cd15d02690b3477e1efc48abc23eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7aeb132116a5d6b738d4bf3d0c765d4c9ae348419477a10c877a0c7e61796411"></a>

## certificate_chain — certificate_chain / 7e7661205f26 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- certificate_chain

<a id="canonical-c399e16a702b8f759cea6be1de70433fbe5dd9a884cf450f9f7ee36e0792fc5e"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-bddb68872f9823e7a20840373f4a229edef7b56b4152aed9e69e18ad444d7494"></a>

## Direct properties — certificate_chain / 7e7661205f26 / 3

<a id="canonical-8fddad31f73da5574309881bbe5992499ac503d1ccff19889bd66b2bec2ea803"></a>

<a id="canonical-3ac0cc6ef468c64c71c1f137a0b205c78a1c11dd47daa0d482b45e16407b678d"></a>

## name property — certificate_chain / 7e7661205f26 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-802f1978f394acd547f47dcfc244b90553a2638586ac1582c6f99fa3a39d7e52"></a>

<a id="canonical-4d32234c5448a439208586e120e568d3541680e76a4636310ecdedf5fd838662"></a>

## namespace property — certificate_chain / 7e7661205f26 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1bac82308b1403aa8300f5c014551339d86d00be41ce25e854c45aefd545cd63"></a>

<a id="canonical-0b6f6fbd912c72ad5fefa5d1e8b2b2b2512ee2be599d7e56cb8d0de84ab641a0"></a>

## tenant property — certificate_chain / 7e7661205f26 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-5e78d39b366e5ad7875f341788697086a75b8bb4ee515e3bab5450aed9cd5659"></a>

## Next pages — certificate_chain / 7e7661205f26 / 7

- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)

<a id="canonical-33bec69c73df335038d518566f1606ceef701d272089c026de644653e3b89376"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fd3cf8edc670baac3f0967eb8583e93addb1e8f9122d0068a8cfd4ec6d17945"></a>

## custom_hash_algorithms — custom_hash_algorithms / 5363e1e0f4ed / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- custom_hash_algorithms

<a id="canonical-c9e08816f6dd4bcbc438488650dee66c5517feed32673d3cbda16ca4bf05ec6f"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_hash\_algorithms, disable\_ocsp\_stapling, use\_system\_defaults; Default:
use\_system\_defaults\] Specifies the hash algorithms to be used.

Upstream description:

Specifies the hash algorithms to be used.

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

- [custom_hash_algorithms](data-sources--certificate--reference--group-001.md#canonical-c9e08816f6dd4bcbc438488650dee66c5517feed32673d3cbda16ca4bf05ec6f)
- [disable_ocsp_stapling](data-sources--certificate--reference--group-001.md#canonical-f8ca53706e93f1dae9fa9f2b26f607fb1cffed3303bea613947f0a3b0b590dc5)
- [use_system_defaults](data-sources--certificate--reference--group-001.md#canonical-c79f92b904f7f7f06d37786e4a9fae271d0da3bc726d640b59326ec5e31f0cae)

Select alternatives according to the provider validators above.

<a id="canonical-a49f21102284f242a597743eeb8bac75d821e17da29da0a0c9e43bd708be615b"></a>

## Direct properties — custom_hash_algorithms / 5363e1e0f4ed / 3

<a id="canonical-8caf79b9b5a5d3eb546a64f2c5ea34c493d48196a9b92b08f304c8182b15106f"></a>

<a id="canonical-7876ab5e1f525222f8e58ece6e45c01816745e8f5f2051ceca295f541c785246"></a>

## hash_algorithms property — custom_hash_algorithms / 5363e1e0f4ed / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0d9f9ea73e81d69c406840623bad59278b6389b25dedd673ad745c82087d7325"></a>

## Next pages — custom_hash_algorithms / 5363e1e0f4ed / 5

- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)

<a id="canonical-73c01fbf9f0d548c650b6a634e929f9a3ec6079fa88f03af929f87efa00223a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65b4fa3fd6d6964299c1d295e895711a043432a22e98df99c6ac1edc71d1a726"></a>

## disable_ocsp_stapling — disable_ocsp_stapling / 3789c1e02d78 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- disable_ocsp_stapling

<a id="canonical-f8ca53706e93f1dae9fa9f2b26f607fb1cffed3303bea613947f0a3b0b590dc5"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-c81a87febf2b1fab06e52177e95e6dac4099a39da2adc8cf5c0a4c5664f72731"></a>

## Direct properties — disable_ocsp_stapling / 3789c1e02d78 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c7567692151bd6e9a207f8468800689bf30bb20252d1e3d0ab128daf8abd36ed"></a>

## Next pages — disable_ocsp_stapling / 3789c1e02d78 / 4

- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)

<a id="canonical-33fb154e59da212c846c109440810f8d8af9779f410f0a4e074cf1737a099299"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3562f51faa3233a639c7c8385c77754763cc3b6a3d2f327cb07bce88be42b319"></a>

## private_key — private_key / 5a9f9f9d4ca0 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- private_key

<a id="canonical-b6b38d47d82666991f2194e90e27ae9482b7b9081bd2270ac4de4c43718cf6e0"></a>

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

<a id="canonical-544ac12f4fa46d02d11a7b3249d859846ea0a4d83ffdda46126173c07f5e00b1"></a>

## Direct properties — private_key / 5a9f9f9d4ca0 / 3

- [blindfold_secret_info](data-sources--certificate--reference--group-001.md#canonical-ceeaa54c64d3a941722f78ae4c67d4eda920237e6fa1df5c9d72df57b358f27b): complete subsection reference.

- [clear_secret_info](data-sources--certificate--reference--group-001.md#canonical-369ae3f4afb3a70a7e30e0dfdecd2fdc5b6626d19ce59abbc5aaef74a789bc2c): complete subsection reference.

<a id="canonical-f918348dd168bac2bb3afd813b2f6743111d53cac2e45a321f595702440dc017"></a>

## Next pages — private_key / 5a9f9f9d4ca0 / 4

- [private_key.blindfold_secret_info](data-sources--certificate--reference--group-001.md#canonical-ceeaa54c64d3a941722f78ae4c67d4eda920237e6fa1df5c9d72df57b358f27b)
- [private_key.clear_secret_info](data-sources--certificate--reference--group-001.md#canonical-369ae3f4afb3a70a7e30e0dfdecd2fdc5b6626d19ce59abbc5aaef74a789bc2c)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)

<a id="canonical-ceeaa54c64d3a941722f78ae4c67d4eda920237e6fa1df5c9d72df57b358f27b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2be63dbd50dc5873a6ef6ba92603967d2d6f82f64e4093ce1918fa912ae26dd"></a>

## private_key.blindfold_secret_info — private_key.blindfold_secret_info / 760479915761 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- [private_key](data-sources--certificate--reference--group-001.md#canonical-33fb154e59da212c846c109440810f8d8af9779f410f0a4e074cf1737a099299)
- private_key.blindfold_secret_info

<a id="canonical-463325481cf2643b362b3b0f119d669023578a14261491b9f6c7c1234bb87db3"></a>

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

<a id="canonical-1c4937f14e1b0ec6c132a76989a8fed0ed3f5864d532cd959f7b2d5a03309f9b"></a>

## Direct properties — private_key.blindfold_secret_info / 760479915761 / 3

<a id="canonical-e4a16af48362524c0ebbf0474a5cc5128f4c77e0919af1d2816ee48a9dc4559e"></a>

<a id="canonical-8d9e124e93ae15ac0b578203ec15eff0b7eb832e02c734c4c13301924c6e4dfd"></a>

## decryption_provider property — private_key.blindfold_secret_info / 760479915761 / 4

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

<a id="canonical-48ba06e60ed1f065a25486380917c47e2199a305296fc94aa3fcd50e58f8a0de"></a>

<a id="canonical-30a9a9a820554c834e29b3bd9248726e626eb0ff121ad585e3d617c081078f85"></a>

## location property — private_key.blindfold_secret_info / 760479915761 / 5

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

<a id="canonical-210e63d8a0ac39933803290e88f8400b6ebb87c561451ed620b7ce1d21d7544c"></a>

<a id="canonical-c58b31c436543ff872eba16387051c0ed6e9a2152c78bea7b6d935a46dff3c4d"></a>

## store_provider property — private_key.blindfold_secret_info / 760479915761 / 6

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

<a id="canonical-60192599591305a6205b0c6cf2da76998ac7dade57215abd83dc485ca8d513b3"></a>

## Next pages — private_key.blindfold_secret_info / 760479915761 / 7

- [private_key](data-sources--certificate--reference--group-001.md#canonical-33fb154e59da212c846c109440810f8d8af9779f410f0a4e074cf1737a099299)
- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)

<a id="canonical-369ae3f4afb3a70a7e30e0dfdecd2fdc5b6626d19ce59abbc5aaef74a789bc2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58a6f8c665fc9e7aa7e59486cd1e4e1e501152acba78b373af3881c5850311d6"></a>

## private_key.clear_secret_info — private_key.clear_secret_info / 41115ff0896e / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- [private_key](data-sources--certificate--reference--group-001.md#canonical-33fb154e59da212c846c109440810f8d8af9779f410f0a4e074cf1737a099299)
- private_key.clear_secret_info

<a id="canonical-2e8229ae91f472f29f0a5208c6a788fc25eb57ee8f1578171e93e75293df6383"></a>

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

<a id="canonical-6d39be43b7d96939175d691c10d1b9f1bca395f9227311c3f3ecb2655dade46e"></a>

## Direct properties — private_key.clear_secret_info / 41115ff0896e / 3

<a id="canonical-c2a643eee508ad0f989638574065b6fa622308b70b6509c3ce9619a2a1be8389"></a>

<a id="canonical-d259042d03587be1a95332dfad79b23684f0c1935e1bfff81d5b97d6849294da"></a>

## provider_ref property — private_key.clear_secret_info / 41115ff0896e / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-70e5e16bb450abd6fa2e7fa66cf6e5cadfbd82368e52fa44d31c7c61d1166f33"></a>

<a id="canonical-a50abf08b13a846c9df0f939059bbefa1220abb8addda44838153b27c959f116"></a>

## url property — private_key.clear_secret_info / 41115ff0896e / 5

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

<a id="canonical-ca353af1d76591600e843fc0f3b10bc44fd94ae83f2c3ade2950ffd89ce74b7a"></a>

## Next pages — private_key.clear_secret_info / 41115ff0896e / 6

- [private_key](data-sources--certificate--reference--group-001.md#canonical-33fb154e59da212c846c109440810f8d8af9779f410f0a4e074cf1737a099299)
- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)

<a id="canonical-3a0ac3fa04557cf75b2cedeb6e4320d5f76aa77298737d7c5661230470d62eec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86365cd02f62494224686ff200adeea393e52516856a8b3a8e25617910e76d78"></a>

## use_system_defaults — use_system_defaults / bf22f8f9cb88 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- use_system_defaults

<a id="canonical-c79f92b904f7f7f06d37786e4a9fae271d0da3bc726d640b59326ec5e31f0cae"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-dd6b70ff5faa095241b88f181c90d96e9309fd2749311da262fb5a6092fe6588"></a>

## Direct properties — use_system_defaults / bf22f8f9cb88 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f50064dbb6815d007ecf0529fc9c02edc08beeed5e64d7bca7b3c099fa292e89"></a>

## Next pages — use_system_defaults / bf22f8f9cb88 / 4

- [Property reference](data-sources--certificate--reference--group-001.md#canonical-2dc6dc0f3cf547c9a867ec8d56d2ac439e02f66e2f161031ecb85fc142b90dc0)
- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
