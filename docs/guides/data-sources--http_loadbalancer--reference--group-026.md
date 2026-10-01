---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-bcba2a2fb6bc9ab366eacb195dcb9bd5f576f8afbc59b708c078c648ed288e55"></a>

## namespace property — user_identification / dfdcbc05625a / 5

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

<a id="canonical-ed1a2c004eda16ee2ee23295cf95ee2853238a7eb60aa07a16311d76669bfae4"></a>

<a id="canonical-ed460eb38f0eb8826e9c6472ab4ddd934c9cf914db645bc1e6948fbca6f9cee7"></a>

## tenant property — user_identification / dfdcbc05625a / 6

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

<a id="canonical-48c3c5512a9016a19f7a9147251cccc38914c410f3a633f7bc4955aa407092d6"></a>

## Next pages — user_identification / dfdcbc05625a / 7

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9b1db9f26e08fe675b915d5eba0a5c3f450d52d1df79ddfd24b811e274dbcf4"></a>

## waf_exclusion — waf_exclusion / 92dee36b22ad / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- waf_exclusion

<a id="canonical-95bb38c89013ab9379542fadbc24e09f768ce9a6f62b36935e5f3d382832163e"></a>

Type: `"single"`. Computed.

Configuration parameter for waf exclusion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

<a id="canonical-4f5e93b7af6816892957e33fbbe0eca96c950da5d0dcd373975c5493e289749d"></a>

## Direct properties — waf_exclusion / 92dee36b22ad / 3

- [waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332): complete subsection reference.

- [waf_exclusion_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-791555c93951db5cc0b2d5922ce07d72741831d32c2fc875595b2977d3940b60): complete subsection reference.

<a id="canonical-bdc574dfc25f1d5c4eb8fc82df094ffde0cd73002465e61c7182800accbcc60f"></a>

## Next pages — waf_exclusion / 92dee36b22ad / 4

- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- [waf_exclusion.waf_exclusion_policy](data-sources--http_loadbalancer--reference--group-026.md#canonical-791555c93951db5cc0b2d5922ce07d72741831d32c2fc875595b2977d3940b60)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2a98958b3225f7e36c8da3c4ab21c3b49e7d31d3fdac8dabee3c0779eed7d07"></a>

## waf_exclusion.waf_exclusion_inline_rules — waf_exclusion.waf_exclusion_inline_rules / d6f752493b86 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- waf_exclusion.waf_exclusion_inline_rules

<a id="canonical-589e0d03fb96e332a126502ee73bfbc46ae5311b265408b99e466cdb831f9889"></a>

Type: `"single"`. Computed.

List of WAF exclusion rules that will be applied inline.

Upstream description:

A list of WAF exclusion rules that will be applied inline.

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

<a id="canonical-6b3a96594a3fda6261f7232f9e8a1ac7c9d089e12be4b13ee3ae0d0fd90809b3"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules / d6f752493b86 / 3

- [rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa): complete subsection reference.

<a id="canonical-8dba1c6a2b08d4b8539566cc40b557784059ea8620307503588cf92824b364f1"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules / d6f752493b86 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7b70185ff608319b687b5954cec9c9df975bc08d6f34bb97586faf3188a0b57"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules — waf_exclusion.waf_exclusion_inline_rules.rules / 1910f13b6ef5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="canonical-f882dc09ad8145564423998a547f9d09380b2644523cf1fd2eefd237f836787f"></a>

Type: `"list"`. Computed.

Ordered list of WAF Exclusions specific to this Load Balancer.

Upstream description:

An ordered list of WAF Exclusions specific to this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-fa336f470c1c136350d57b7f320623a46e7f907cfbf079d32628368e98575532"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules / 1910f13b6ef5 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-026.md#canonical-44d8f80356fd9b83e2ddfceca8ab0c2c1249d9b728f4daba4fda2cf17a31427a): complete subsection reference.

- [any_path](data-sources--http_loadbalancer--reference--group-026.md#canonical-00105ec4d91813d84c352640e6640ad9db35b92c5388b7ae7413b32326b02b1c): complete subsection reference.

- [app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-026.md#canonical-9863132fbc4c62ac813288fb3c920732e81060a6b7b71e5d8255024298eb845d): complete subsection reference.

<a id="canonical-249d3545a3078b646fa8b3c2d6405c7c348b9f7d6b1be3a1ff665efacd966947"></a>

<a id="canonical-1b266052ad37c4b915260c40b0313f03cfe5e1f96f7413d9656831dfcb5ec167"></a>

## exact_value property — waf_exclusion.waf_exclusion_inline_rules.rules / 1910f13b6ef5 / 4

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-d6083ee5d01d1e1c5b7277031bda7cb9c6144cdcd704fe52fab1ad81b8356a23"></a>

<a id="canonical-8ccf2113151354e51a0a0c052beec5179563a640c37a0b11dc6b1a48cb3bd76d"></a>

## expiration_timestamp property — waf_exclusion.waf_exclusion_inline_rules.rules / 1910f13b6ef5 / 5

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [metadata](data-sources--http_loadbalancer--reference--group-026.md#canonical-902fc8c9ec13a2e6e82cfaa389373de976a0255b84ee656a944e9a6b3ddc2e08): complete subsection reference.

<a id="canonical-af613e75a05c3b421b8c2d408aef59c22cbdfc368f7bbdc01eb5b7555a771267"></a>

<a id="canonical-c6fd85b4a071c15c6aa2ff9a7382fd264c6b89723f28b53d7e7ff5c9ca59c7b9"></a>

## methods property — waf_exclusion.waf_exclusion_inline_rules.rules / 1910f13b6ef5 / 6

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f8bb6debbe4c1f10e458028f3386b3429e7df9a55a15779f3827a7ce8a723623"></a>

<a id="canonical-70c660668f7cc6405406977f86d77828e4cc3cb453ca663d378c5f2938366776"></a>

## path_prefix property — waf_exclusion.waf_exclusion_inline_rules.rules / 1910f13b6ef5 / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-5b97dfa47386bd2d434885a7c39d8b79e5068aacda9ac4364ec0b38e14f0f6aa"></a>

<a id="canonical-714412af11b0e17e4b5561e9d22ed0266dd730ba069ad2bdba62d4ebe1711c96"></a>

## path_regex property — waf_exclusion.waf_exclusion_inline_rules.rules / 1910f13b6ef5 / 8

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-d1cb93f2720327277dfd7f479ffd41618a0e38da7330259b358754b4a59678de"></a>

<a id="canonical-06941c94e8aad858736172819474f6fc019e6d35b3dd75b69af6736fba28d2a8"></a>

## suffix_value property — waf_exclusion.waf_exclusion_inline_rules.rules / 1910f13b6ef5 / 9

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [waf_skip_processing](data-sources--http_loadbalancer--reference--group-026.md#canonical-286856ab948f80fa90fee1ffd437d251f239968392aeb57a983cdba04b7dc491): complete subsection reference.

<a id="canonical-b306244e7d512e1717832bd53cc0004414aaedb4dd89808e97df853d690ddfc3"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules / 1910f13b6ef5 / 10

- [waf_exclusion.waf_exclusion_inline_rules.rules.any_domain](data-sources--http_loadbalancer--reference--group-026.md#canonical-44d8f80356fd9b83e2ddfceca8ab0c2c1249d9b728f4daba4fda2cf17a31427a)
- [waf_exclusion.waf_exclusion_inline_rules.rules.any_path](data-sources--http_loadbalancer--reference--group-026.md#canonical-00105ec4d91813d84c352640e6640ad9db35b92c5388b7ae7413b32326b02b1c)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-026.md#canonical-9863132fbc4c62ac813288fb3c920732e81060a6b7b71e5d8255024298eb845d)
- [waf_exclusion.waf_exclusion_inline_rules.rules.metadata](data-sources--http_loadbalancer--reference--group-026.md#canonical-902fc8c9ec13a2e6e82cfaa389373de976a0255b84ee656a944e9a6b3ddc2e08)
- [waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing](data-sources--http_loadbalancer--reference--group-026.md#canonical-286856ab948f80fa90fee1ffd437d251f239968392aeb57a983cdba04b7dc491)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-44d8f80356fd9b83e2ddfceca8ab0c2c1249d9b728f4daba4fda2cf17a31427a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e40eb7659642672bcc87cc4d5a4cb34b5a8108bd9246850df95e23432c1fa86"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_domain — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / d74b9104a1a5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_domain

<a id="canonical-4409a6278b9ceda9e773b09fa1e500c71e0c9d460a4b8b9e595e328457eca8c4"></a>

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

<a id="canonical-b10040bd0bb57d30c6755b25c41c449da78c41e7c2d78824c1ac32bb702234fb"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / d74b9104a1a5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f15de6d06d3e2aa83d8f9a887d91fa48a5d254ac17c46aba4b84fe52c870ed85"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.any_domain / d74b9104a1a5 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-00105ec4d91813d84c352640e6640ad9db35b92c5388b7ae7413b32326b02b1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aee49fb43ac1f4a1cca65da030af1cc7d43f037df6d1f83033a80bd0922c43f3"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_path — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / 9db3d72b5759 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_path

<a id="canonical-106c2c07d57e6f680824e5c8d8baf758fe2e404c0e7fef9a83fd84afdabb289c"></a>

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

<a id="canonical-c2bdfb89b19bd7c6613cc4f2cd91ce6eafb389bc8b46d1bf745f03131643ffc0"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / 9db3d72b5759 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d62327d1bb3a5a711f05f492796e244c6f7587bcdf998c3f582a97af72bb3b30"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.any_path / 9db3d72b5759 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9863132fbc4c62ac813288fb3c920732e81060a6b7b71e5d8255024298eb845d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0799dfa20151f4b2d1c95a30ba249aaa0e394fe2a35a7dafc5247c7f49a99295"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / 83e554e04e50 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="canonical-2d265be2a09f7faa7242ff3bac6aff101edea724310b3c07e27ec0ce066b01bb"></a>

Type: `"single"`. Computed.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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

<a id="canonical-a510d77435d4352bac77df8dd98ef75dc163b97e44e47a2486d3f46c8bcd010a"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / 83e554e04e50 / 3

- [exclude_attack_type_contexts](data-sources--http_loadbalancer--reference--group-026.md#canonical-c72323ae78f52b55e9403bb907e6c603ed63412a46dab1aaddfe94208c5f0d9d): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--http_loadbalancer--reference--group-026.md#canonical-b8559a1cd46a3dc4f3d0f33ac608ca1156f6e09e2ff97357f7dc765aecbb2171): complete subsection reference.

- [exclude_signature_contexts](data-sources--http_loadbalancer--reference--group-026.md#canonical-9be29cd8aa332541357a4b1e7424fddf14114336d6c244fcefd045cf0577061f): complete subsection reference.

- [exclude_violation_contexts](data-sources--http_loadbalancer--reference--group-026.md#canonical-ac7ca591d07641f93d43ca9b74d141aaf0952ab5b383b7d286e3a6e42fa2008c): complete subsection reference.

<a id="canonical-e07c0c6704c5aa84c042cbabcff751b2ee8e5a4a0ffa67daebfdd5ebfc52d019"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control / 83e554e04e50 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--http_loadbalancer--reference--group-026.md#canonical-c72323ae78f52b55e9403bb907e6c603ed63412a46dab1aaddfe94208c5f0d9d)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--http_loadbalancer--reference--group-026.md#canonical-b8559a1cd46a3dc4f3d0f33ac608ca1156f6e09e2ff97357f7dc765aecbb2171)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts](data-sources--http_loadbalancer--reference--group-026.md#canonical-9be29cd8aa332541357a4b1e7424fddf14114336d6c244fcefd045cf0577061f)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts](data-sources--http_loadbalancer--reference--group-026.md#canonical-ac7ca591d07641f93d43ca9b74d141aaf0952ab5b383b7d286e3a6e42fa2008c)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c72323ae78f52b55e9403bb907e6c603ed63412a46dab1aaddfe94208c5f0d9d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4cc4ae81cd39cbe1d10ed1ad9ed9adc46cac25dbebaa2af1228db4a11521df5"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 0d00057b58fb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-026.md#canonical-9863132fbc4c62ac813288fb3c920732e81060a6b7b71e5d8255024298eb845d)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-da504660c9d8ad93a3b2e641359a195c0b1870232717d610897c880d8d839ab1"></a>

Type: `"list"`. Computed.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-760614bea64492eeced3651d7d492ac85a1b5e809c4ac176171c464040c5be16"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 0d00057b58fb / 3

<a id="canonical-c6e983f95f0fad50dffa2a8315d37cfb5700ac1d5dd817cc11d231ce95e923c5"></a>

<a id="canonical-d9a85c01cbe0446df327682bf7d2034baddec3dd9ef5fd05f7c7da1a2d58c0f7"></a>

## context property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 0d00057b58fb / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d3965187328a0ffbdf8ea4009c071a182df052a5f7aa3fbbd2c6dbc63ae62cf3"></a>

<a id="canonical-f902c4a4e7ca5b293714654dfa501330c055430a16d5f11a2db8d9c3c2d3013c"></a>

## context_name property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 0d00057b58fb / 5

Type: `"string"`. Computed.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3c45b7bad0d4c209914079f2c0002a903bb08145641902a1d5578307e27954f3"></a>

<a id="canonical-b7a4d0e1208bcd1954ba16145f2b955a9927439845c6aed39078e51950d34487"></a>

## exclude_attack_type property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 0d00057b58fb / 6

Type: `"string"`. Computed.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1f441c58142312efb01726e8a2129f2f0f140fa1ceecc96b4bb77e38f45844dc"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 0d00057b58fb / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-026.md#canonical-9863132fbc4c62ac813288fb3c920732e81060a6b7b71e5d8255024298eb845d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b8559a1cd46a3dc4f3d0f33ac608ca1156f6e09e2ff97357f7dc765aecbb2171"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8cc84219329fa0562ee245670c75c0b11bdfebda3721c20aaeb59702317eca2f"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / ef6aa2e083e2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-026.md#canonical-9863132fbc4c62ac813288fb3c920732e81060a6b7b71e5d8255024298eb845d)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-db1d511e363db54f86586cedcc0c3896416f031b4b8de8684fe70db106e19880"></a>

Type: `"list"`. Computed.

Bot Names to be excluded for the defined match criteria.

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

<a id="canonical-01107105c9493ecc0aa9f42883cff1ef1f749486fa6c45dbea480e9548875edb"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / ef6aa2e083e2 / 3

<a id="canonical-ea50727d99837eb4d3a81c5afffdda4f7547821b5739c9fbdb4fa2b085888a67"></a>

<a id="canonical-812490540a4d97bd1a555cfae320cb497ef7c8b544c252d27dd6e9fce8f9d4b0"></a>

## bot_name property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / ef6aa2e083e2 / 4

Type: `"string"`. Computed.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-baf602ff738104a1156dbef7725b184352050658b91e86d21120e541d9edcf82"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / ef6aa2e083e2 / 5

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-026.md#canonical-9863132fbc4c62ac813288fb3c920732e81060a6b7b71e5d8255024298eb845d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9be29cd8aa332541357a4b1e7424fddf14114336d6c244fcefd045cf0577061f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dab419d27d243ba407b3473e5d62f58a8cc86a2f33ab18c4095f190f6fbca9d3"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e8a31ddcfba3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-026.md#canonical-9863132fbc4c62ac813288fb3c920732e81060a6b7b71e5d8255024298eb845d)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-e8d58f15c3006b4f4b73ca712c706d176041e4b7a8a4b6fe9d0e5ba34c8af664"></a>

Type: `"list"`. Computed.

Signature IDs to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-eb93f7ef438f1a4cf8c6ed4af61830a74508e2942d7727cf26c92851d58aca17"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e8a31ddcfba3 / 3

<a id="canonical-7ec69f48dc67d5a318aa9dde1e984632737dff22fca1a6acb7e45faba205d5d1"></a>

<a id="canonical-fa4304162730b12006c8bad0ea4a23f683e62fa189d6e3edbf35a148a752ab35"></a>

## context property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e8a31ddcfba3 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-368b6321da98d8c92b5bf741f325c2872a9c939fe07f06684cd6eecd1f88c368"></a>

<a id="canonical-27be45ef2eb7b36e438d1850efab0860c6658097a1eff27b3f858333dfac5931"></a>

## context_name property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e8a31ddcfba3 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-8a4701d84b88cedb219770940af9a99df9eacde66d88d05b52d5b01f9e8ae977"></a>

<a id="canonical-dc0e86b2fb6ddb361254d802c26220aef2183fbc1c9ad2cf15c84ecd205178e5"></a>

## signature_id property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e8a31ddcfba3 / 6

Type: `"number"`. Computed.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-e85ab320422f0c6a4d722cd50e14c6dbbc96bc6ce4ba6855a118ff8e2d9aa4ed"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / e8a31ddcfba3 / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-026.md#canonical-9863132fbc4c62ac813288fb3c920732e81060a6b7b71e5d8255024298eb845d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ac7ca591d07641f93d43ca9b74d141aaf0952ab5b383b7d286e3a6e42fa2008c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aa07b28594c08d443a3ccde71d5aacd9d4f5ba643b8623321b1a66117bef4091"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 27ba67229170 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-026.md#canonical-9863132fbc4c62ac813288fb3c920732e81060a6b7b71e5d8255024298eb845d)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-9d8dad26fe0fa706f2acc4bfc6a84ce1b11fbc821b541cfdd6d69f86c1495c5e"></a>

Type: `"list"`. Computed.

Violations to be excluded for the defined match criteria.

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

<a id="canonical-cfa7ef07ff8f260943ec905b6cb42ccf1c6c59c228e8b08a5236f02554d43522"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 27ba67229170 / 3

<a id="canonical-629df90ba8acdb6f643e69a084bbe7261fd184f10e7aed144abc0235771d0f87"></a>

<a id="canonical-af6995a4ebf867e28c62f83d5b6280d910f0c7440f3ef44ae44a91eb78e70124"></a>

## context property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 27ba67229170 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c640289f267b584546fb05cc0298aed3a88cc8c09627eebcae453e65375c6441"></a>

<a id="canonical-d1524081c8130636de09f83d53fc82f12496aaac2aae84867e584e40b01b8304"></a>

## context_name property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 27ba67229170 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-fba4aab704ed3844ac19a2996798729ff7d9da6f6f8bc4efbf8412d25821f089"></a>

<a id="canonical-565389b6db248497136a27f57efff2a2deaabe383029728594359d70b760d4af"></a>

## exclude_violation property — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 27ba67229170 / 6

Type: `"string"`. Computed.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Upstream description:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-cdee7a5d9159633fd65ca306bb87c44f8030e0bf69d3845b8738367a81ffb20e"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.ex / 27ba67229170 / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-026.md#canonical-9863132fbc4c62ac813288fb3c920732e81060a6b7b71e5d8255024298eb845d)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-902fc8c9ec13a2e6e82cfaa389373de976a0255b84ee656a944e9a6b3ddc2e08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9f5d7c5b2847d831ce78d7ae6418707e66b3debe7c5f425e70f32f3ba82ba1b"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.metadata — waf_exclusion.waf_exclusion_inline_rules.rules.metadata / 7f995e80531f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- waf_exclusion.waf_exclusion_inline_rules.rules.metadata

<a id="canonical-725096d60c4a8f72397359c11323112993ebd5296cf08ce422f83f03c7b04411"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-b02b2f343b31e3da97777fd5999254bff05e7b6c07656b2b570647eb25354d3a"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.metadata / 7f995e80531f / 3

<a id="canonical-b6d48d0ef635b921d716b07866fa4c9f8b0c471b60f6856f22732058e3031ccf"></a>

<a id="canonical-e128b4658ce87ba592aca7353abd76aa42dff67301d5b0a25dbc75ef90cec5d4"></a>

## description_spec property — waf_exclusion.waf_exclusion_inline_rules.rules.metadata / 7f995e80531f / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-d70faaa0a410cd51d4aaf5f4ffa2190c6e4e901cbcdd244d63f98423806207c1"></a>

<a id="canonical-e19e71bff5d75ce0f9d3dbe3957b1a24de9d8fc3013632584245a2a542397133"></a>

## name property — waf_exclusion.waf_exclusion_inline_rules.rules.metadata / 7f995e80531f / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-dda1807ff1e3c82bb2fd5a2f82f3480bd81747bc5c453d1627a5b1f416ee6924"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.metadata / 7f995e80531f / 6

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-286856ab948f80fa90fee1ffd437d251f239968392aeb57a983cdba04b7dc491"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25a3367b45eb46b6848750c7febdab16addfea745d9ce524ab33b08b8b847b07"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing — waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing / c5d7d5a7a0d3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-2cc2ce4f192de3d9bd8e98e927956aad11bb89d042c6cdd644d49639c31af332)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing

<a id="canonical-faa1a3b5e97769fa712fe63017a7aecbd775a5c2c87155cc89bdaf463cd188ef"></a>

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

<a id="canonical-f799bdcbd33320fc7f0903b386e4f57126b3d0ce8953f1fcee8d9fd09a263c87"></a>

## Direct properties — waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing / c5d7d5a7a0d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-656bc37dff4d3bef6aac12f1c624c3e3d990a675c1d6ed5d6f17e03cbbf9ae93"></a>

## Next pages — waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing / c5d7d5a7a0d3 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-026.md#canonical-5c88fd80bffb59fda16e5352c0fb373c31277972f2b8324a00c6e2d29c8876fa)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-791555c93951db5cc0b2d5922ce07d72741831d32c2fc875595b2977d3940b60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9436c003dc9c8fecccd8a3cf394ea8cc47eac6a481449deaf3854da278ad9c1"></a>

## waf_exclusion.waf_exclusion_policy — waf_exclusion.waf_exclusion_policy / eb9cb1abd812 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- waf_exclusion.waf_exclusion_policy

<a id="canonical-911e2d61c28c1840c20f9fd79f0b65855872e38ee7836f9365851e4dc9b457d8"></a>

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

<a id="canonical-1602fc966ebdb549df7b2740ed8485a85fb83bd60b07bd139f67e1bfd184b3bf"></a>

## Direct properties — waf_exclusion.waf_exclusion_policy / eb9cb1abd812 / 3

<a id="canonical-567745ba8ce089f3c527d8e75d613da0f00f1783432f78f952b1cf14270eec80"></a>

<a id="canonical-166ee04bfb14eb5d297d0651973fefe31df23974deed29a7d1316817fb80b67a"></a>

## name property — waf_exclusion.waf_exclusion_policy / eb9cb1abd812 / 4

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

<a id="canonical-098aa536759da76979f7bcc10e15a4331c0b54b9527093e288b844f7dc0f8c1d"></a>

<a id="canonical-b1c79f4560b469a474e036f51cee65e629d3c40cfc3d9358a7076f99ae2e033b"></a>

## namespace property — waf_exclusion.waf_exclusion_policy / eb9cb1abd812 / 5

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

<a id="canonical-27424489a38e2ac0483f74eaf7b6930d781aff00d8e62d37537d8425e9d932c9"></a>

<a id="canonical-7f7b2acc2e6872ea7e24668495f1701e45da5610e51e129df6a2ce8951341e08"></a>

## tenant property — waf_exclusion.waf_exclusion_policy / eb9cb1abd812 / 6

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

<a id="canonical-8f1156aca792932238fa45ea390ca71c39a5e524873dcd2e09670c1fbd9a2be0"></a>

## Next pages — waf_exclusion.waf_exclusion_policy / eb9cb1abd812 / 7

- [waf_exclusion](data-sources--http_loadbalancer--reference--group-026.md#canonical-a197893b5dbd319a62a2418c76682c7fced4291fdee9027239d8fa6df307bdbe)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
