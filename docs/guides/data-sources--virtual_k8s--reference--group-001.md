---
page_title: "xcsh_virtual_k8s reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s reference."
---

# xcsh_virtual_k8s reference

<a id="canonical-6430b8a0d84baaea3ec220faf51e4b844976a149d4122fc1303f4649e3aeea1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86b54661857bc0d25e925478a4401d9c31babcf122835ed680b2d8e8d659cf64"></a>

## Property reference — Property reference / 6ea73aad3fde / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)
- Property reference

<a id="canonical-9bd62bed743c4c4bed50498ab75fcde34ab93c28132f1c3c879e8ae92365329a"></a>

## Direct properties — Property reference / 6ea73aad3fde / 3

<a id="canonical-6c89bc0960d44b53dd303042fb8d3fd2619a4a2b747f2774fd3b4e39dad8e867"></a>

<a id="canonical-c62fbbcbc6c99702255c26cbb96b3710c569451d01c13e14bff4a368e16bf9ab"></a>

## annotations property — Property reference / 6ea73aad3fde / 4

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

- [default_flavor_ref](data-sources--virtual_k8s--reference--group-001.md#canonical-34f4a6f62e16bd96eb202812b3b0c61dabbef8452116558af907a751164959bd): complete subsection reference.

<a id="canonical-b3407e5961eda522ec9dc6c4f695555311f3c3f6b6f37caaff936e0cf13f60c2"></a>

<a id="canonical-fdbaf15fb2b991d3700bb91adcbc462571d5bac0c70276a6ca42b8d05301334b"></a>

## description property — Property reference / 6ea73aad3fde / 5

Type: `"string"`. Computed.

Description of the VirtualK8S.

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

- [disabled](data-sources--virtual_k8s--reference--group-001.md#canonical-13b6635c55326acc21d6e693c5a91e5f42362c11451d9139dcb61dbf22645257): complete subsection reference.

<a id="canonical-c200c449c717b23b636a61610e6000c3854eedf94055f6809208f73c91c3277d"></a>

<a id="canonical-e5db05ea9da2736b3b4dfd687de9b9cb00dbd28ec8daad33bd79ebfbfcc3ff29"></a>

## id property — Property reference / 6ea73aad3fde / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

- [isolated](data-sources--virtual_k8s--reference--group-001.md#canonical-f12981dac05b1cbef4f3dfc7f62f46d635634e1b0a53505a4324cd40dd81a43c): complete subsection reference.

<a id="canonical-ccaf693c02a654090ace67c1e3bfd83cc20ea9e02fc154f95bdcbb627df964fd"></a>

<a id="canonical-540124a063e4938e7f98e129742f36f5cd759f26030f5b13ff6a15d443c7d00b"></a>

## labels property — Property reference / 6ea73aad3fde / 7

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

<a id="canonical-59f5afea9209a1c810f7161bd145419c6e7068181e19834a63f02ce5e739b776"></a>

<a id="canonical-7ec6e632b0dda97aa65fd2fc2c4a05a144fd6232dc9a0912bf28ec8272b16af4"></a>

## name property — Property reference / 6ea73aad3fde / 8

Type: `"string"`. Required.

Name of the VirtualK8S.

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

<a id="canonical-997ff4dbdb8411600e94edbee614519078a6ec8185afd9970eafa603906d334e"></a>

<a id="canonical-994324ceda64f2aebe957bde7efd15b01bcd46b5cf099e74fb2f7eccd5f41dde"></a>

## namespace property — Property reference / 6ea73aad3fde / 9

Type: `"string"`. Required.

Namespace where the VirtualK8S exists.

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

- [vsite_refs](data-sources--virtual_k8s--reference--group-001.md#canonical-a8dd54f0f954994dceb7465fdc0463114ab3c95e329a69a8c7e4a647c3e434ca): complete subsection reference.

<a id="canonical-35be5d616b40570782dc38a2d776d59590410796fca0206da348ef56506e417c"></a>

## All schema paths — Property reference / 6ea73aad3fde / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--virtual_k8s--reference--group-001.md#canonical-6c89bc0960d44b53dd303042fb8d3fd2619a4a2b747f2774fd3b4e39dad8e867) |
| `default_flavor_ref` | [default_flavor_ref](data-sources--virtual_k8s--reference--group-001.md#canonical-321b4fee3d90a07efbe7cd8aa930683b0a5914a65879ccd30dabf36078f6370d) |
| `default_flavor_ref.name` | [default_flavor_ref.name](data-sources--virtual_k8s--reference--group-001.md#canonical-ccacc53341003ce92ca6200fc5a83bdd4893206035bdf190c7ab72d42acdaf2b) |
| `default_flavor_ref.namespace` | [default_flavor_ref.namespace](data-sources--virtual_k8s--reference--group-001.md#canonical-9933b251824bf54d8d6473d58b04fd09483911a3008f8b4d75dee687c2a48bea) |
| `default_flavor_ref.tenant` | [default_flavor_ref.tenant](data-sources--virtual_k8s--reference--group-001.md#canonical-b3b2e91e2229a3608c35a59071b3037ea8010775f61d24b895c2b24f5547f66c) |
| `description` | [description](data-sources--virtual_k8s--reference--group-001.md#canonical-b3407e5961eda522ec9dc6c4f695555311f3c3f6b6f37caaff936e0cf13f60c2) |
| `disabled` | [disabled](data-sources--virtual_k8s--reference--group-001.md#canonical-2e869b9766575223f2882c31917fa72df0ecb04b5d32da535619b12331aea240) |
| `id` | [id](data-sources--virtual_k8s--reference--group-001.md#canonical-c200c449c717b23b636a61610e6000c3854eedf94055f6809208f73c91c3277d) |
| `isolated` | [isolated](data-sources--virtual_k8s--reference--group-001.md#canonical-f6231dceeb67608556a5cbd299e962941184beea85d80a0dc673ca7186fe4095) |
| `labels` | [labels](data-sources--virtual_k8s--reference--group-001.md#canonical-ccaf693c02a654090ace67c1e3bfd83cc20ea9e02fc154f95bdcbb627df964fd) |
| `name` | [name](data-sources--virtual_k8s--reference--group-001.md#canonical-59f5afea9209a1c810f7161bd145419c6e7068181e19834a63f02ce5e739b776) |
| `namespace` | [namespace](data-sources--virtual_k8s--reference--group-001.md#canonical-997ff4dbdb8411600e94edbee614519078a6ec8185afd9970eafa603906d334e) |
| `vsite_refs` | [vsite_refs](data-sources--virtual_k8s--reference--group-001.md#canonical-a1f83f90fb1851813c6d1f811fa7e12e99cfd2b87dfc53c74c399f9d3488dd75) |
| `vsite_refs.kind` | [vsite_refs.kind](data-sources--virtual_k8s--reference--group-001.md#canonical-db1a8b4f913be8b40a437c1a681c4678e65473a59bd220e2b16678f154f5d300) |
| `vsite_refs.name` | [vsite_refs.name](data-sources--virtual_k8s--reference--group-001.md#canonical-33a1db0c662c34d1992610dec9399e80be09416c59a788a57b8c1ffb96344cf0) |
| `vsite_refs.namespace` | [vsite_refs.namespace](data-sources--virtual_k8s--reference--group-001.md#canonical-f1562ec1e13bdeb9c3f16871e6c7ab1622d2505a367994fd4e0ba1c793222b12) |
| `vsite_refs.tenant` | [vsite_refs.tenant](data-sources--virtual_k8s--reference--group-001.md#canonical-307995d09b3e6828c2fb3a395b86f4d079db80848d300c71ad91e1d76ef5063e) |
| `vsite_refs.uid` | [vsite_refs.uid](data-sources--virtual_k8s--reference--group-001.md#canonical-8c349249544726da30a09471e43e40f54ba74dba6cbf12e9dd0b42f87d1dddb5) |

<a id="canonical-ea6910a50e157c2c66883ec25b6dea1dde1f2724146a42b59945769517e7a620"></a>

## Next pages — Property reference / 6ea73aad3fde / 11

- [default_flavor_ref](data-sources--virtual_k8s--reference--group-001.md#canonical-34f4a6f62e16bd96eb202812b3b0c61dabbef8452116558af907a751164959bd)
- [disabled](data-sources--virtual_k8s--reference--group-001.md#canonical-13b6635c55326acc21d6e693c5a91e5f42362c11451d9139dcb61dbf22645257)
- [isolated](data-sources--virtual_k8s--reference--group-001.md#canonical-f12981dac05b1cbef4f3dfc7f62f46d635634e1b0a53505a4324cd40dd81a43c)
- [vsite_refs](data-sources--virtual_k8s--reference--group-001.md#canonical-a8dd54f0f954994dceb7465fdc0463114ab3c95e329a69a8c7e4a647c3e434ca)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)

<a id="canonical-34f4a6f62e16bd96eb202812b3b0c61dabbef8452116558af907a751164959bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4567dc2dc166a96e251985db4b657de390f8ada59416510cb41619330c57e2a"></a>

## default_flavor_ref — default_flavor_ref / 0fbb5e121208 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)
- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-6430b8a0d84baaea3ec220faf51e4b844976a149d4122fc1303f4649e3aeea1c)
- default_flavor_ref

<a id="canonical-321b4fee3d90a07efbe7cd8aa930683b0a5914a65879ccd30dabf36078f6370d"></a>

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

<a id="canonical-8a41125d48751393175ea9829d77239dc485b5a49254b988221d40fd7b011e08"></a>

## Direct properties — default_flavor_ref / 0fbb5e121208 / 3

<a id="canonical-ccacc53341003ce92ca6200fc5a83bdd4893206035bdf190c7ab72d42acdaf2b"></a>

<a id="canonical-a498b4ce0870638ff2f35259ebbfc1b74ba9a38c6dddad55e7dd5c3c1d6ce6f1"></a>

## name property — default_flavor_ref / 0fbb5e121208 / 4

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

<a id="canonical-9933b251824bf54d8d6473d58b04fd09483911a3008f8b4d75dee687c2a48bea"></a>

<a id="canonical-f48d2329c16dcb353704dcce1a7fd498d34c8e9ff509ccb1722110c417c40664"></a>

## namespace property — default_flavor_ref / 0fbb5e121208 / 5

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

<a id="canonical-b3b2e91e2229a3608c35a59071b3037ea8010775f61d24b895c2b24f5547f66c"></a>

<a id="canonical-b73be597b653ce2619d74b89bfc5e1fdaf2d078658a9840c873eb36bb4bdf03f"></a>

## tenant property — default_flavor_ref / 0fbb5e121208 / 6

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

<a id="canonical-b9d0d44ff898d93f0c67821233ceed0c4a8e5f38434746486b9c5f631dfc5f43"></a>

## Next pages — default_flavor_ref / 0fbb5e121208 / 7

- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-6430b8a0d84baaea3ec220faf51e4b844976a149d4122fc1303f4649e3aeea1c)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)

<a id="canonical-13b6635c55326acc21d6e693c5a91e5f42362c11451d9139dcb61dbf22645257"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d89472fa6fb6b257e90089df434b6cad2f9401e237ac1d69f0be5102904296bc"></a>

## disabled — disabled / 42bc6641da12 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)
- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-6430b8a0d84baaea3ec220faf51e4b844976a149d4122fc1303f4649e3aeea1c)
- disabled

<a id="canonical-2e869b9766575223f2882c31917fa72df0ecb04b5d32da535619b12331aea240"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disabled, isolated\] Enable this option

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

- [disabled](data-sources--virtual_k8s--reference--group-001.md#canonical-2e869b9766575223f2882c31917fa72df0ecb04b5d32da535619b12331aea240)
- [isolated](data-sources--virtual_k8s--reference--group-001.md#canonical-f6231dceeb67608556a5cbd299e962941184beea85d80a0dc673ca7186fe4095)

Select alternatives according to the provider validators above.

<a id="canonical-f3b69d928ce4ebd03ec37fba2716a57bf28ce091f39b061fbfcf42880315e92c"></a>

## Direct properties — disabled / 42bc6641da12 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e81e8dccac75bf8e6294ac5f61f84cc6894d8fe6425ed2f1a93dd04ff190006"></a>

## Next pages — disabled / 42bc6641da12 / 4

- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-6430b8a0d84baaea3ec220faf51e4b844976a149d4122fc1303f4649e3aeea1c)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)

<a id="canonical-f12981dac05b1cbef4f3dfc7f62f46d635634e1b0a53505a4324cd40dd81a43c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8b9f86e53572376e42e250947007b5e28cf69d76cea07da99c0c30e5e7963c0"></a>

## isolated — isolated / c89e2a7a027c / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)
- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-6430b8a0d84baaea3ec220faf51e4b844976a149d4122fc1303f4649e3aeea1c)
- isolated

<a id="canonical-f6231dceeb67608556a5cbd299e962941184beea85d80a0dc673ca7186fe4095"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-83563b0ad197b7cd7533df8658adc8cc746f518d439b6fca8dd873386d7b2a8c"></a>

## Direct properties — isolated / c89e2a7a027c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3eb00fdc6b4fefe6cba5487fb582ab712438a15e6f583076042e493524096e1"></a>

## Next pages — isolated / c89e2a7a027c / 4

- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-6430b8a0d84baaea3ec220faf51e4b844976a149d4122fc1303f4649e3aeea1c)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)

<a id="canonical-a8dd54f0f954994dceb7465fdc0463114ab3c95e329a69a8c7e4a647c3e434ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c7e7ad7fba2f93ff9969c0f281574fd7915e1bdf2f4a21d284544044adb7cde"></a>

## vsite_refs — vsite_refs / 4df828be7297 / 2

Breadcrumbs:

- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)
- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-6430b8a0d84baaea3ec220faf51e4b844976a149d4122fc1303f4649e3aeea1c)
- vsite_refs

<a id="canonical-a1f83f90fb1851813c6d1f811fa7e12e99cfd2b87dfc53c74c399f9d3488dd75"></a>

Type: `"list"`. Computed.

Reference to virtual-sites Default virtual-site of the Virtual K8s object. If no virtual-site is
specified in the Kubernetes API resource object annotations via F5 XC/virtual-sites, then this
virtual-site is used select sites on which to instantiate the Kubernetes API resource object.

Upstream description:

Reference to virtual-sites Default virtual-site of the Virtual K8s object. If no virtual-site is
specified in the Kubernetes API resource object annotations via F5 XC/virtual-sites, then this
virtual-site is used select sites on which to instantiate the Kubernetes API resource object.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-3e726df8185d5c292193ff07895d43f880e9454ce16c1432e29807e80e4c99f6"></a>

## Direct properties — vsite_refs / 4df828be7297 / 3

<a id="canonical-db1a8b4f913be8b40a437c1a681c4678e65473a59bd220e2b16678f154f5d300"></a>

<a id="canonical-89c939577141dbd4282bb785b9bec1919540a78937474cf23b213f40f06e299c"></a>

## kind property — vsite_refs / 4df828be7297 / 4

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

<a id="canonical-33a1db0c662c34d1992610dec9399e80be09416c59a788a57b8c1ffb96344cf0"></a>

<a id="canonical-734fc3f0cb641a4f94c7306e0306b86af8d7f138418111f3409e8d2d9d4ff128"></a>

## name property — vsite_refs / 4df828be7297 / 5

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

<a id="canonical-f1562ec1e13bdeb9c3f16871e6c7ab1622d2505a367994fd4e0ba1c793222b12"></a>

<a id="canonical-7457ebe395cb502b24a3ea8eb5b3241abd64ff6fa934b8c739e8f0d6889b9dae"></a>

## namespace property — vsite_refs / 4df828be7297 / 6

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

<a id="canonical-307995d09b3e6828c2fb3a395b86f4d079db80848d300c71ad91e1d76ef5063e"></a>

<a id="canonical-e53a239551552f049ada9eda7f183d17fe63359715560869e61dbc6e5b233fde"></a>

## tenant property — vsite_refs / 4df828be7297 / 7

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

<a id="canonical-8c349249544726da30a09471e43e40f54ba74dba6cbf12e9dd0b42f87d1dddb5"></a>

<a id="canonical-544c9e63524475a113b2fac41cb47adaab86486bb46d4e9a5034223dcef60bff"></a>

## uid property — vsite_refs / 4df828be7297 / 8

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

<a id="canonical-043481ba103a3170036c09f25e7d43440e503a15a6ea26437c352131ae98de83"></a>

## Next pages — vsite_refs / 4df828be7297 / 9

- [Property reference](data-sources--virtual_k8s--reference--group-001.md#canonical-6430b8a0d84baaea3ec220faf51e4b844976a149d4122fc1303f4649e3aeea1c)
- [xcsh_virtual_k8s](../data-sources/virtual_k8s.md#canonical-0148db5af8d910e157c059c0db0b323f7791d6e9d01c95201d415ea9704f837e)
