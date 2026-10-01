---
page_title: "xcsh_user_identification reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_user_identification reference."
---

# xcsh_user_identification reference

<a id="canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71bfb96fc97f4ec30127e58a6f8cc03cd1f2494d30d73ecfd4020e42e2fcf07b"></a>

## Property reference — Property reference / 2ffc3e3c069e / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- Property reference

<a id="canonical-b73c0a576d8b9f355a7dfd3cfc13b43e442f1182d4cc85d0b106864a0f935458"></a>

## Direct properties — Property reference / 2ffc3e3c069e / 3

<a id="canonical-7a7d139a9cfa736fda5aeea9bfed59f037b54bf34c6a817b21702998ccfd7161"></a>

<a id="canonical-c635f146278ed247bcf2a22fc8b42493739e28fce486c9f7c8a156be54210fd5"></a>

## annotations property — Property reference / 2ffc3e3c069e / 4

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

<a id="canonical-f778c810ce4c53e052cdf052405134e16ddeb1edbb75a75aade3deab7af3d826"></a>

<a id="canonical-dfbe50806bb739e2bdbbb3c2be6fc64d1bd73bbce09d9c36130e8aa990956596"></a>

## description property — Property reference / 2ffc3e3c069e / 5

Type: `"string"`. Computed.

Description of the UserIdentification.

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

<a id="canonical-a642ea317183cb2b6e8bcbea7db33548d1d53737a2af2111f67bd76632750824"></a>

<a id="canonical-ab2200ef1ccbbf29b6a4e49fe65d09869f23a6005ceae3ef19a83102bae89f2d"></a>

## id property — Property reference / 2ffc3e3c069e / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-649d7d0b17eac1902f6ee3f54e0c8fe5a9742092f074da7db9765320cfd62e6f"></a>

<a id="canonical-2e1be84aca174f11d4ff20df0137ba063ecf7e5638f8276e0fe100cad09daeaf"></a>

## labels property — Property reference / 2ffc3e3c069e / 7

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

<a id="canonical-d226792de1a3f15f1f50d722025d0456786f8f54d2bd22d91d093f26cd8b1b53"></a>

<a id="canonical-afbe7add5917c0f8a3cc0b6e8c6885bb4714586f80198f1114788baecf15b5fd"></a>

## name property — Property reference / 2ffc3e3c069e / 8

Type: `"string"`. Required.

Name of the UserIdentification.

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

<a id="canonical-29534d4fe67391b02428dac96ae3bdf761301f62d972891034e8df331697f163"></a>

<a id="canonical-f7d106fbc41ea416e6b9066a032b4e8bf1ac862f0416f43f31111c26b3d713a4"></a>

## namespace property — Property reference / 2ffc3e3c069e / 9

Type: `"string"`. Required.

Namespace where the UserIdentification exists.

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

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf): complete subsection reference.

<a id="canonical-95c1c77a0a084fa6284e08f95308e49d287db1f08ee8b764078faede4a8002b9"></a>

## All schema paths — Property reference / 2ffc3e3c069e / 10

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--user_identification--reference--group-001.md#canonical-7a7d139a9cfa736fda5aeea9bfed59f037b54bf34c6a817b21702998ccfd7161) |
| `description` | [description](data-sources--user_identification--reference--group-001.md#canonical-f778c810ce4c53e052cdf052405134e16ddeb1edbb75a75aade3deab7af3d826) |
| `id` | [id](data-sources--user_identification--reference--group-001.md#canonical-a642ea317183cb2b6e8bcbea7db33548d1d53737a2af2111f67bd76632750824) |
| `labels` | [labels](data-sources--user_identification--reference--group-001.md#canonical-649d7d0b17eac1902f6ee3f54e0c8fe5a9742092f074da7db9765320cfd62e6f) |
| `name` | [name](data-sources--user_identification--reference--group-001.md#canonical-d226792de1a3f15f1f50d722025d0456786f8f54d2bd22d91d093f26cd8b1b53) |
| `namespace` | [namespace](data-sources--user_identification--reference--group-001.md#canonical-29534d4fe67391b02428dac96ae3bdf761301f62d972891034e8df331697f163) |
| `rules` | [rules](data-sources--user_identification--reference--group-001.md#canonical-708f68939d21f259988ee92527551263120490994aaaa6f532c898bffece2a4b) |
| `rules.client_asn` | [rules.client_asn](data-sources--user_identification--reference--group-001.md#canonical-524358079f2b4b47d9593de7cd88c49b2958a448d0d4c09820351a17e8622562) |
| `rules.client_city` | [rules.client_city](data-sources--user_identification--reference--group-001.md#canonical-64daaeb72fc2f3a4c1500732b43ee73bba4807c1d081a59da2f7c59b0ee9186f) |
| `rules.client_country` | [rules.client_country](data-sources--user_identification--reference--group-001.md#canonical-e9b7f3f39685bb1940748453026cfe0555849beb5bce22fdfafbb550d2c00f49) |
| `rules.client_ip` | [rules.client_ip](data-sources--user_identification--reference--group-001.md#canonical-98b07f5b9e166c384f48639781a792fbb53759e78ec80f8734a99c970d0f255d) |
| `rules.client_region` | [rules.client_region](data-sources--user_identification--reference--group-001.md#canonical-13648731656e3d0175a44cf08e4854f9a6c7fa5f087ea63303b060aa1199f582) |
| `rules.cookie_name` | [rules.cookie_name](data-sources--user_identification--reference--group-001.md#canonical-304c2c84c8dd0d9a6cd76037f0dadd26924837b0e55b153907eca55d638a27f4) |
| `rules.http_header_name` | [rules.http_header_name](data-sources--user_identification--reference--group-001.md#canonical-cfd004c16ed6cdf19f0fd35d70784473af304651eb3fa7d2f8555255bbc52572) |
| `rules.ip_and_http_header_name` | [rules.ip_and_http_header_name](data-sources--user_identification--reference--group-001.md#canonical-dd88d91228ab9bfa3cf135defd2a99a2c0ab43597644c72bf75959693f437da6) |
| `rules.ip_and_ja4_tls_fingerprint` | [rules.ip_and_ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-ee0a4e05f0e5402c01b7d7829ca23c0d35972a03d8fb0e06b1e93051e429d435) |
| `rules.ip_and_tls_fingerprint` | [rules.ip_and_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-5e87aab9eb12497b332f9eb56b43727a3da99d99299337af4c1d264f16b3ee25) |
| `rules.ja4_tls_fingerprint` | [rules.ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-7c01dcc174e82aa99acf62de9fecf48aede61ce5fec067e2b09e390499080329) |
| `rules.jwt_claim_name` | [rules.jwt_claim_name](data-sources--user_identification--reference--group-001.md#canonical-e8bb510c737f217b5508158abd5ecaa5574deb6b285abec573972a03089789bf) |
| `rules.none` | [rules.none](data-sources--user_identification--reference--group-001.md#canonical-76dbc06526cd81e02bd14a4da57d682b21c4391ae40de160207eb18e573c2162) |
| `rules.query_param_key` | [rules.query_param_key](data-sources--user_identification--reference--group-001.md#canonical-3ea614a8c203e861103e1fb273bf9f0f4b13eed1dc69a5a6b4fefd7d0d34ef87) |
| `rules.tls_fingerprint` | [rules.tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-53257a8fdbae88d68ee998d8f0bffc2da3d7d4cba19816a9f6639fe7c41fae35) |

<a id="canonical-a9586606c7afa6c05ca7ff22d897dcdd2b5cc35c94785853de7a18ea0d66a7e4"></a>

## Next pages — Property reference / 2ffc3e3c069e / 11

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3ec7e556fa29b1052b7a32a5f9d9a7f352b6ad267342277774e85cc691fcaf1"></a>

## rules — rules / 4fc16e1840b8 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- rules

<a id="canonical-708f68939d21f259988ee92527551263120490994aaaa6f532c898bffece2a4b"></a>

Type: `"list"`. Computed.

Ordered list of rules that are evaluated sequentially against the input fields extracted from an API
request in order to determine a user identifier. Evaluation of the rules is terminated once a user
identifier has been extracted.

Upstream description:

An ordered list of rules that are evaluated sequentially against the input fields extracted from an
API request in order to determine a user identifier. Evaluation of the rules is terminated once a
user identifier has been extracted.

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

<a id="canonical-ab8bc169e03385d88c9411cd69efd02ee2b7b096b713631b9d645385fc830db5"></a>

## Direct properties — rules / 4fc16e1840b8 / 3

- [client_asn](data-sources--user_identification--reference--group-001.md#canonical-a5c47d71a5b4185c2fdf7101e3be1e1e85487cb2b7004a7d46fbe99026dc1d2a): complete subsection reference.

- [client_city](data-sources--user_identification--reference--group-001.md#canonical-0675a0e7cdb287bbda7a52681a6665a14d14c10f66c4431f76a5feff2a0f082b): complete subsection reference.

- [client_country](data-sources--user_identification--reference--group-001.md#canonical-6988cd654f2c796f35b6f7b1d05802ca9f344cc35f9279d818465c2b4520fda9): complete subsection reference.

- [client_ip](data-sources--user_identification--reference--group-001.md#canonical-d5ea901874529ebb354c88224d38392639bb18b20e7ce169f85b3de7b03a075b): complete subsection reference.

- [client_region](data-sources--user_identification--reference--group-001.md#canonical-3f64ec3edd0a470fc16c00d5ff0da35b7eacd256f98e7d73d9540b31d1674b89): complete subsection reference.

<a id="canonical-304c2c84c8dd0d9a6cd76037f0dadd26924837b0e55b153907eca55d638a27f4"></a>

<a id="canonical-7eadeb5a2a367ded327bee5f4f6635c5da4284f06b1fb11423c117a98a070298"></a>

## cookie_name property — rules / 4fc16e1840b8 / 4

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key
tls\_fingerprint\] Use the HTTP cookie value for the given name as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key
tls\_fingerprint\] Use the HTTP cookie value for the given name as user identifier.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-cfd004c16ed6cdf19f0fd35d70784473af304651eb3fa7d2f8555255bbc52572"></a>

<a id="canonical-ce6e8fcb1a35398e082e77db12a35ee5f3add8d6aab5bdf05c735eedd21397a0"></a>

## http_header_name property — rules / 4fc16e1840b8 / 5

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint
ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Use the HTTP header
value for the given name as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint
ja4\_tls\_fingerprint jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Use the HTTP header
value for the given name as user identifier.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-dd88d91228ab9bfa3cf135defd2a99a2c0ab43597644c72bf75959693f437da6"></a>

<a id="canonical-b8334fd58065a7d19e1e0ae7f33dd143ba1cb26076285d2ce15d2e34c5a3ca81"></a>

## ip_and_http_header_name property — rules / 4fc16e1840b8 / 6

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint
jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Name of HTTP header from which the value
should be extracted.

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_ja4\_tls\_fingerprint ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint
jwt\_claim\_name none query\_param\_key tls\_fingerprint\] Name of HTTP header from which the value
should be extracted.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [ip_and_ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-b8cc7abecb8bf718858a664dcc52aba645a89ab745c0a61bdffd69e7fe598ce8): complete subsection reference.

- [ip_and_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-5cf28bcb8adff617dd0d522d415fa46615f744054c4a789b8844845b1bd476c0): complete subsection reference.

- [ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-7f7b2b5f4612ae5ad041a7b8bfe974e8d70c028db052d157aaab5a4ac1b00a83): complete subsection reference.

<a id="canonical-e8bb510c737f217b5508158abd5ecaa5574deb6b285abec573972a03089789bf"></a>

<a id="canonical-da2907d52b05888af9f2f30d03f78cb4ac6ab24af83906382a0a50f5ca2a2d6f"></a>

## jwt_claim_name property — rules / 4fc16e1840b8 / 7

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint none query\_param\_key tls\_fingerprint\] Use the
JWT claim value as user identifier.

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint none query\_param\_key tls\_fingerprint\] Use the
JWT claim value as user identifier.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [none](data-sources--user_identification--reference--group-001.md#canonical-3d45bc7027506e75364a19c17b431b7437f5cd1abc0412a01575ae90f73d9a2a): complete subsection reference.

<a id="canonical-3ea614a8c203e861103e1fb273bf9f0f4b13eed1dc69a5a6b4fefd7d0d34ef87"></a>

<a id="canonical-8d8a822a0aba8ef58b296563aa9c87eac73db3d4179bdb7f8ffc6f0d52073523"></a>

## query_param_key property — rules / 4fc16e1840b8 / 8

Type: `"string"`. Computed.

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none tls\_fingerprint\] Use the
query parameter value for the given key as user..

Upstream description:

Exclusive with \[client\_asn client\_city client\_country client\_ip client\_region cookie\_name
http\_header\_name ip\_and\_http\_header\_name ip\_and\_ja4\_tls\_fingerprint
ip\_and\_tls\_fingerprint ja4\_tls\_fingerprint jwt\_claim\_name none tls\_fingerprint\] Use the
query parameter value for the given key as user identifier.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-eb866a9b789fecaacfc749c12e44453a5c1513e1ad7c4dd1041e3d223e674d83): complete subsection reference.

<a id="canonical-6c1a8830044d62fd08c6bfe17226bff1d752192c7424f367470721986c291250"></a>

## Next pages — rules / 4fc16e1840b8 / 9

- [rules.client_asn](data-sources--user_identification--reference--group-001.md#canonical-a5c47d71a5b4185c2fdf7101e3be1e1e85487cb2b7004a7d46fbe99026dc1d2a)
- [rules.client_city](data-sources--user_identification--reference--group-001.md#canonical-0675a0e7cdb287bbda7a52681a6665a14d14c10f66c4431f76a5feff2a0f082b)
- [rules.client_country](data-sources--user_identification--reference--group-001.md#canonical-6988cd654f2c796f35b6f7b1d05802ca9f344cc35f9279d818465c2b4520fda9)
- [rules.client_ip](data-sources--user_identification--reference--group-001.md#canonical-d5ea901874529ebb354c88224d38392639bb18b20e7ce169f85b3de7b03a075b)
- [rules.client_region](data-sources--user_identification--reference--group-001.md#canonical-3f64ec3edd0a470fc16c00d5ff0da35b7eacd256f98e7d73d9540b31d1674b89)
- [rules.ip_and_ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-b8cc7abecb8bf718858a664dcc52aba645a89ab745c0a61bdffd69e7fe598ce8)
- [rules.ip_and_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-5cf28bcb8adff617dd0d522d415fa46615f744054c4a789b8844845b1bd476c0)
- [rules.ja4_tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-7f7b2b5f4612ae5ad041a7b8bfe974e8d70c028db052d157aaab5a4ac1b00a83)
- [rules.none](data-sources--user_identification--reference--group-001.md#canonical-3d45bc7027506e75364a19c17b431b7437f5cd1abc0412a01575ae90f73d9a2a)
- [rules.tls_fingerprint](data-sources--user_identification--reference--group-001.md#canonical-eb866a9b789fecaacfc749c12e44453a5c1513e1ad7c4dd1041e3d223e674d83)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-a5c47d71a5b4185c2fdf7101e3be1e1e85487cb2b7004a7d46fbe99026dc1d2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-159789dd341d2606c1b3e90a12a37f99bee4087775def0b095d5ef763a2df66d"></a>

## rules.client_asn — rules.client_asn / c7d49215348f / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- rules.client_asn

<a id="canonical-524358079f2b4b47d9593de7cd88c49b2958a448d0d4c09820351a17e8622562"></a>

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

<a id="canonical-0ab0d822f10dd666baa38e0414aef96ecc5ccc86a3146133228c198de6432732"></a>

## Direct properties — rules.client_asn / c7d49215348f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bbb8f2087783ea108ba29993c8c74cb0614f1612d8e6799dff23628bf3f3345e"></a>

## Next pages — rules.client_asn / c7d49215348f / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-0675a0e7cdb287bbda7a52681a6665a14d14c10f66c4431f76a5feff2a0f082b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37e8a0c25a3429ecb118648e0e4256afcfb20ed12664f5b1c0e4ac4e68ae9fd7"></a>

## rules.client_city — rules.client_city / e24165f535aa / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- rules.client_city

<a id="canonical-64daaeb72fc2f3a4c1500732b43ee73bba4807c1d081a59da2f7c59b0ee9186f"></a>

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

<a id="canonical-c8c0e171fad7e5120657dc993ff62a99e466ae0593c0f89ca1761df9c709b743"></a>

## Direct properties — rules.client_city / e24165f535aa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14cdb13e9cd9a775dba1e9bd183e645a08bfadb29ff468b8b6dd4e3b38595eb5"></a>

## Next pages — rules.client_city / e24165f535aa / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-6988cd654f2c796f35b6f7b1d05802ca9f344cc35f9279d818465c2b4520fda9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-912c2ddaf31650091469f24b4a849275e0f05a8d3d641bd0f4eb57c0d824097b"></a>

## rules.client_country — rules.client_country / 9e3c118b699b / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- rules.client_country

<a id="canonical-e9b7f3f39685bb1940748453026cfe0555849beb5bce22fdfafbb550d2c00f49"></a>

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

<a id="canonical-86762bc495efc5115dcc6d0f9ccebbb9a6004fbd5bf2c14458d5a1fee9487e0a"></a>

## Direct properties — rules.client_country / 9e3c118b699b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-343105cd660744de197f9dc12907c4699c5ef60286dfb8e1c680e203f848cb5a"></a>

## Next pages — rules.client_country / 9e3c118b699b / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-d5ea901874529ebb354c88224d38392639bb18b20e7ce169f85b3de7b03a075b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b465f694938a99fe6690b83b78afea15c2a55ee411ceaace8dfc9e71aa830e1c"></a>

## rules.client_ip — rules.client_ip / f4ed72bf38fd / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- rules.client_ip

<a id="canonical-98b07f5b9e166c384f48639781a792fbb53759e78ec80f8734a99c970d0f255d"></a>

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

<a id="canonical-c13ca1ba56ed032146502b4fc6989a37f29162c07fd580950af1a7e977910f3a"></a>

## Direct properties — rules.client_ip / f4ed72bf38fd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-830b9e62cda4789bb855d82fc21a73c9051b7e4b799da9703c6b1349c566236e"></a>

## Next pages — rules.client_ip / f4ed72bf38fd / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-3f64ec3edd0a470fc16c00d5ff0da35b7eacd256f98e7d73d9540b31d1674b89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4061afe190980fdf5e99cfb9d580278fad628eb36f817a1b3b596360131d0dad"></a>

## rules.client_region — rules.client_region / 7d1d64be5dcc / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- rules.client_region

<a id="canonical-13648731656e3d0175a44cf08e4854f9a6c7fa5f087ea63303b060aa1199f582"></a>

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

<a id="canonical-4c5bd2fcf7387c47eeb1e044bf087c47991e6d3be4a77144b0aa83233dce958a"></a>

## Direct properties — rules.client_region / 7d1d64be5dcc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5b4b63c0d541ac3ec068406c5d994e81b2cfa8b9874812dd1233b2d40a20574d"></a>

## Next pages — rules.client_region / 7d1d64be5dcc / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-b8cc7abecb8bf718858a664dcc52aba645a89ab745c0a61bdffd69e7fe598ce8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-423e922f56880e9b2c86b9b4d475955f3d18cbe0bb18339c00f1fccb55d9169e"></a>

## rules.ip_and_ja4_tls_fingerprint — rules.ip_and_ja4_tls_fingerprint / 97601033435b / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- rules.ip_and_ja4_tls_fingerprint

<a id="canonical-ee0a4e05f0e5402c01b7d7829ca23c0d35972a03d8fb0e06b1e93051e429d435"></a>

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

<a id="canonical-412199914d747560ca696dfdaa450f34c92216edfd2b8c2e0c20419dc8388233"></a>

## Direct properties — rules.ip_and_ja4_tls_fingerprint / 97601033435b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7821544403135c9578db320fbac288743cba4917e009c012453260f95a6ea21c"></a>

## Next pages — rules.ip_and_ja4_tls_fingerprint / 97601033435b / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-5cf28bcb8adff617dd0d522d415fa46615f744054c4a789b8844845b1bd476c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9892c03b7a3a6fb78058b736d740f489195c734eb90bbd846afb154d7ac45f75"></a>

## rules.ip_and_tls_fingerprint — rules.ip_and_tls_fingerprint / c8ebfe6719db / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- rules.ip_and_tls_fingerprint

<a id="canonical-5e87aab9eb12497b332f9eb56b43727a3da99d99299337af4c1d264f16b3ee25"></a>

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

<a id="canonical-591e94b8ce588e23f845630c9ce3294c8d61a5936e2da85bc6bd6d55881d2ccf"></a>

## Direct properties — rules.ip_and_tls_fingerprint / c8ebfe6719db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2a4248d08899c78c63a88648e373a29a8ad8f277d84513916d954f509d8ba948"></a>

## Next pages — rules.ip_and_tls_fingerprint / c8ebfe6719db / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-7f7b2b5f4612ae5ad041a7b8bfe974e8d70c028db052d157aaab5a4ac1b00a83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2635c234a0d8dd62bf66f72dbfbe34bdc51c5f852cecfe154bb843656f4a4dd5"></a>

## rules.ja4_tls_fingerprint — rules.ja4_tls_fingerprint / 2a17edec45a4 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- rules.ja4_tls_fingerprint

<a id="canonical-7c01dcc174e82aa99acf62de9fecf48aede61ce5fec067e2b09e390499080329"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ja4 tls fingerprint.

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

<a id="canonical-7f4e8fcf464fb7e956c6533560c7524be90982b824cbd3fa3a35548c6350b45b"></a>

## Direct properties — rules.ja4_tls_fingerprint / 2a17edec45a4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ee82b75ca080d329b3c96ca82dbf99265c5a2c121afb44fbb52f28399f2828c2"></a>

## Next pages — rules.ja4_tls_fingerprint / 2a17edec45a4 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-3d45bc7027506e75364a19c17b431b7437f5cd1abc0412a01575ae90f73d9a2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-378444d1044d8cc8a26950c051d1ba54691684c59bea3958545b2e0d3b9e675d"></a>

## rules.none — rules.none / eb88712da085 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- rules.none

<a id="canonical-76dbc06526cd81e02bd14a4da57d682b21c4391ae40de160207eb18e573c2162"></a>

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

<a id="canonical-5cdcf35ced68e0be16a2740787085fa265f7893507816e89296740a01c65e6cf"></a>

## Direct properties — rules.none / eb88712da085 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3eb7b11cbccddc6bdfa06668bd788bc3f5125d9beb20f22eed04e1c5ed549ef5"></a>

## Next pages — rules.none / eb88712da085 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)

<a id="canonical-eb866a9b789fecaacfc749c12e44453a5c1513e1ad7c4dd1041e3d223e674d83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35a0af789924204845eb67c58a7af8043996252317465d658e8d192de620c00e"></a>

## rules.tls_fingerprint — rules.tls_fingerprint / 71dc09dd8816 / 2

Breadcrumbs:

- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
- [Property reference](data-sources--user_identification--reference--group-001.md#canonical-f464f49eb30c13755f8401e0e930641af6052dbeb168e74f29af74ea3e119f17)
- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- rules.tls_fingerprint

<a id="canonical-53257a8fdbae88d68ee998d8f0bffc2da3d7d4cba19816a9f6639fe7c41fae35"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for tls fingerprint.

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

<a id="canonical-ca54bfebf61fad29db996da1ba06c6f5593bed465c7c440aac996e9dff9e5a50"></a>

## Direct properties — rules.tls_fingerprint / 71dc09dd8816 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-625dfb6166518774c41e44dbc516aa4cafeef2ae97c9463b95229432ff768728"></a>

## Next pages — rules.tls_fingerprint / 71dc09dd8816 / 4

- [rules](data-sources--user_identification--reference--group-001.md#canonical-0be1145374ba0999380c55cf3641e9ed579046cb135285a78eae38e047cd2baf)
- [xcsh_user_identification](../data-sources/user_identification.md#canonical-a181f7e07b2823653ffef09ea2b308d030ae99f9cb03d038485d5730a7902e39)
