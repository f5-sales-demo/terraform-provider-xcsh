---
page_title: "xcsh_trusted_ca_list reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list reference."
---

# xcsh_trusted_ca_list reference

<a id="canonical-dcb8d54ef3f35aabca505c635ade0b5d8ebdf23a0c754eba510299b4cac35f27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbf35106ef6de765477cc34e5fbafeefe09dab13ad0c6e499a8140520fe60787"></a>

## Property reference — Property reference / 542b08955431 / 2

Breadcrumbs:

- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md#canonical-ed5a0ca4002a9acc208559e5d135f6efdf6a9d227170b35fe4e3ac7a7857549a)
- Property reference

<a id="canonical-cff755e621cb4c78c2471585fbd415c92a371841964537bcb3dd9014b91efa97"></a>

## Direct properties — Property reference / 542b08955431 / 3

<a id="canonical-01ff7cf6b3f04d8b7cadbd007b7601235ed8b48bb39a9d432b5e35cbefbbb62d"></a>

<a id="canonical-f21812ebc2da1bbee456910577977a1bf08a23147a1fc843cc286417693e0893"></a>

## annotations property — Property reference / 542b08955431 / 4

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

<a id="canonical-3c60b6b51abb12aaa33e0dc3e96c8a4e74cde481abd0b7aed61e951ac3faf298"></a>

<a id="canonical-9e2819341c86899740fa539638fc3b3201a154965a949d330a6152c00af39528"></a>

## description property — Property reference / 542b08955431 / 5

Type: `"string"`. Computed.

Description of the TrustedCAList.

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

<a id="canonical-52b269c3eff23e4d4fbb5a19bce6ceecadee8e4208a9212d3ad9aa6402464cbc"></a>

<a id="canonical-a8bdf824451a41dc09d50a7a41caccb43693b9c414ae3fe24e731952bc96eed0"></a>

## id property — Property reference / 542b08955431 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1f83d4322d97683d9e75edd66824ad297e426469850edac393359318d07f08f8"></a>

<a id="canonical-9361fa8f9402850fd60c33d13edbfbb0df09c689c15bdb4f5eb293abc5c89cc3"></a>

## labels property — Property reference / 542b08955431 / 7

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

<a id="canonical-f6dddb4bc7c53925e244ab38c9fc9e9a7727bf9f1f9632e0fcfab949b633ccf2"></a>

<a id="canonical-ec99bc0c10f25d524ff39038f4da5f913cec526ba2754f35e3d8e58cf59728f8"></a>

## name property — Property reference / 542b08955431 / 8

Type: `"string"`. Required.

Name of the TrustedCAList.

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

<a id="canonical-8279959deb7a6e264ba72c0f022178be72f496d4ac1dc4fd87527a9a088e0c4c"></a>

<a id="canonical-fe27fa378e8e7d79755684547d65f11c6bc738d87ae5192a7f7452bbba9e1b67"></a>

## namespace property — Property reference / 542b08955431 / 9

Type: `"string"`. Required.

Namespace where the TrustedCAList exists.

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

<a id="canonical-2250513e5468b6e109c3ff287471ab654d4aa02f211ccdb36b739c297bba29b2"></a>

<a id="canonical-941ebc47b9e0e01f85bc08417823dd712b278391122371752a94a2a8aa37b492"></a>

## trusted_ca_url property — Property reference / 542b08955431 / 10

Type: `"string"`. Computed.

Trusted CA certificates for validating certificates.

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

<a id="canonical-48bcb4181f4404477648e5eb93479dce1566fae911d35b6bffbed0ec4d174328"></a>

## All schema paths — Property reference / 542b08955431 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--trusted_ca_list--reference--group-001.md#canonical-01ff7cf6b3f04d8b7cadbd007b7601235ed8b48bb39a9d432b5e35cbefbbb62d) |
| `description` | [description](data-sources--trusted_ca_list--reference--group-001.md#canonical-3c60b6b51abb12aaa33e0dc3e96c8a4e74cde481abd0b7aed61e951ac3faf298) |
| `id` | [id](data-sources--trusted_ca_list--reference--group-001.md#canonical-52b269c3eff23e4d4fbb5a19bce6ceecadee8e4208a9212d3ad9aa6402464cbc) |
| `labels` | [labels](data-sources--trusted_ca_list--reference--group-001.md#canonical-1f83d4322d97683d9e75edd66824ad297e426469850edac393359318d07f08f8) |
| `name` | [name](data-sources--trusted_ca_list--reference--group-001.md#canonical-f6dddb4bc7c53925e244ab38c9fc9e9a7727bf9f1f9632e0fcfab949b633ccf2) |
| `namespace` | [namespace](data-sources--trusted_ca_list--reference--group-001.md#canonical-8279959deb7a6e264ba72c0f022178be72f496d4ac1dc4fd87527a9a088e0c4c) |
| `trusted_ca_url` | [trusted_ca_url](data-sources--trusted_ca_list--reference--group-001.md#canonical-2250513e5468b6e109c3ff287471ab654d4aa02f211ccdb36b739c297bba29b2) |

<a id="canonical-1cf4e40de1b94ac937f588360fb0add9801b5a268ff26361d3821a5be4c3093d"></a>

## Next pages — Property reference / 542b08955431 / 12

- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md#canonical-ed5a0ca4002a9acc208559e5d135f6efdf6a9d227170b35fe4e3ac7a7857549a)
