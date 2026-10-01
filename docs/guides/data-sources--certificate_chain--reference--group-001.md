---
page_title: "xcsh_certificate_chain reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain reference."
---

# xcsh_certificate_chain reference

<a id="canonical-d9436ba22d8301c6a0f02ff16a4d8396e3f87a177d590cb144774d33852ac5b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93c5ccdfcf479c545b06c8818b5e8c2772e0cb7f93267beacac4831760741ca1"></a>

## Property reference — Property reference / 5314490bb355 / 2

Breadcrumbs:

- [xcsh_certificate_chain](../data-sources/certificate_chain.md#canonical-ce3d2afc41637e97fe322a943a0e26b7ffd06d1052a2a27eea5dfed15f9134f9)
- Property reference

<a id="canonical-4a834f3534d12b98d9c74297016f825b611fef3653fc28da7b4ccfce848f4a91"></a>

## Direct properties — Property reference / 5314490bb355 / 3

<a id="canonical-e932ffb8eb56fac8359d711414446b2bd51eb79771240d18f536a912715393a4"></a>

<a id="canonical-70450ab9aa1b3f5790982822d82ad5ff90728dc1b649f5acf9ba2f29a50bf4da"></a>

## annotations property — Property reference / 5314490bb355 / 4

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

<a id="canonical-cb49c05ca0007879b2c9ccd0d9180e3abe08e06b4c1278b1a4dd3354bd8f231c"></a>

<a id="canonical-a8e2cc4bc4110052610e3d1f9c8e7178ef7c1c794ab57b6699789dd9b30ed427"></a>

## certificate_url property — Property reference / 5314490bb355 / 5

Type: `"string"`. Computed.

Certificate chain is the list of intermediate certificates in PEM format including the PEM headers.

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
    "ves.io.schema.rules.string.intermediate_certificate_chain_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.intermediate_certificate_chain_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-ecdca2f79d8473a4e50714cf3d0b01ade2e61efca80a54cd080213e078deb2ba"></a>

<a id="canonical-e6fbc98ba3790b1cdfe212e6c6378ef9be132629b2f6c8e0b7d80db6e2d60d1f"></a>

## description property — Property reference / 5314490bb355 / 6

Type: `"string"`. Computed.

Description of the CertificateChain.

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

<a id="canonical-131de657d3b97833c18e927fb814f3ea49e481b2a5e7953db74e3c72698ab53e"></a>

<a id="canonical-2559093da527597722d925012841aa20e55094807764179d5ceb1520bd7124ba"></a>

## id property — Property reference / 5314490bb355 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-9873e1228d4114ad1f6f3df8248e25d58e5edd3303a19c8092257f65ce545283"></a>

<a id="canonical-6d5f0b8c60e9a3e2e91bef5f8410b2205c2a52e069ab77c5723c775085b4d4b2"></a>

## labels property — Property reference / 5314490bb355 / 8

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

<a id="canonical-526348949a4e4e0dd4a4207d90b298640b2f6d15dd27a0db061da8c67b1dcf25"></a>

<a id="canonical-3e172c635358514490a26a13b5c2541fec6f6c4eee8c64b35e5998ae72cde509"></a>

## name property — Property reference / 5314490bb355 / 9

Type: `"string"`. Required.

Name of the CertificateChain.

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

<a id="canonical-0ee59ab8bbbaa4abd1054a20b3b4bf27928ff722860dd5b91dc6c8635621bf92"></a>

<a id="canonical-cd1113858af4a368deb97cc55a4a484a698d3e0ba67c234c62f7ebacb5865680"></a>

## namespace property — Property reference / 5314490bb355 / 10

Type: `"string"`. Required.

Namespace where the CertificateChain exists.

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

<a id="canonical-27cddb85fa28c6d2bdb3b35fd253776e1289392786833dc019a23dcf97e7e392"></a>

## All schema paths — Property reference / 5314490bb355 / 11

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--certificate_chain--reference--group-001.md#canonical-e932ffb8eb56fac8359d711414446b2bd51eb79771240d18f536a912715393a4) |
| `certificate_url` | [certificate_url](data-sources--certificate_chain--reference--group-001.md#canonical-cb49c05ca0007879b2c9ccd0d9180e3abe08e06b4c1278b1a4dd3354bd8f231c) |
| `description` | [description](data-sources--certificate_chain--reference--group-001.md#canonical-ecdca2f79d8473a4e50714cf3d0b01ade2e61efca80a54cd080213e078deb2ba) |
| `id` | [id](data-sources--certificate_chain--reference--group-001.md#canonical-131de657d3b97833c18e927fb814f3ea49e481b2a5e7953db74e3c72698ab53e) |
| `labels` | [labels](data-sources--certificate_chain--reference--group-001.md#canonical-9873e1228d4114ad1f6f3df8248e25d58e5edd3303a19c8092257f65ce545283) |
| `name` | [name](data-sources--certificate_chain--reference--group-001.md#canonical-526348949a4e4e0dd4a4207d90b298640b2f6d15dd27a0db061da8c67b1dcf25) |
| `namespace` | [namespace](data-sources--certificate_chain--reference--group-001.md#canonical-0ee59ab8bbbaa4abd1054a20b3b4bf27928ff722860dd5b91dc6c8635621bf92) |

<a id="canonical-68bce9c56c67e889a26a71933cb5eba798cc00fb071d84c75ccf959f98e09a56"></a>

## Next pages — Property reference / 5314490bb355 / 12

- [xcsh_certificate_chain](../data-sources/certificate_chain.md#canonical-ce3d2afc41637e97fe322a943a0e26b7ffd06d1052a2a27eea5dfed15f9134f9)
