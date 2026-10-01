---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-6da3ca0bb79c1aede6a5ca9993c1ea9290aaedac5ae7c93143e7d4b6e1daa5eb"></a>

## primary.default_rr_set_group.ds_record.values.sha256_digest — primary.default_rr_set_group.ds_record.values.sha256_digest / 2719c7b1573a / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-3ef88f0643cae6c90abb5fde17697e969b3f9536dffdfc6919f32a93e600fe90)
- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-5ffdc52aab4896b597f8f36c508520f2da9d29ab4e4d129880080ae08be4b810)
- primary.default_rr_set_group.ds_record.values.sha256_digest

<a id="canonical-654ea7690ee98cc491c81e4eed22052df1995918522198f3e834d26bca9e704d"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 digest.

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

<a id="canonical-5a8c509dd1a6f47348cdb2c68c1ce78cde0a4e730aa08e6a8f006209c63d2e08"></a>

## Direct properties — primary.default_rr_set_group.ds_record.values.sha256_digest / 2719c7b1573a / 3

<a id="canonical-4dafb0b395ca3ef206aed65a460c68cf58e6b70139e720f6a46bf8808010165c"></a>

<a id="canonical-b40bd9824a4c6f4408d51fc555b8f5cddf847d981b112bc0ef215ec0fb2bc972"></a>

## digest property — primary.default_rr_set_group.ds_record.values.sha256_digest / 2719c7b1573a / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 64
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-2729014d25dfee193b59e105d2ed6a828af3585796b00e1946d69b5d0404970a"></a>

## Next pages — primary.default_rr_set_group.ds_record.values.sha256_digest / 2719c7b1573a / 5

- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-5ffdc52aab4896b597f8f36c508520f2da9d29ab4e4d129880080ae08be4b810)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-8abbf950198894d9b203f18a6e8391e221e5971bf4d7c89b8a0285ee150bf7b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-119ba7e24f8448ddb081ab0b04c7aea3c91219e4c42574cde85653eb88a37363"></a>

## primary.default_rr_set_group.ds_record.values.sha384_digest — primary.default_rr_set_group.ds_record.values.sha384_digest / c558bd00cfec / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-3ef88f0643cae6c90abb5fde17697e969b3f9536dffdfc6919f32a93e600fe90)
- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-5ffdc52aab4896b597f8f36c508520f2da9d29ab4e4d129880080ae08be4b810)
- primary.default_rr_set_group.ds_record.values.sha384_digest

<a id="canonical-63aab37fa7dee6df1e2e69f89a90556dddad581d5b06734ed4961b2fc4c7e1e1"></a>

Type: `"single"`. Computed.

Configuration parameter for sha384 digest.

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

<a id="canonical-bb10f19edb46d2e3c8aadbe4a667e9889811af409f9f1020125bd4c8950610d0"></a>

## Direct properties — primary.default_rr_set_group.ds_record.values.sha384_digest / c558bd00cfec / 3

<a id="canonical-9081e150c9ba17b51434b56c2cd5c5b94ad95f1140d7c132059cc2d318debdf6"></a>

<a id="canonical-457f84538faf8988588ff75116c403887c890b2306e85c1bae06e2cdbaaca840"></a>

## digest property — primary.default_rr_set_group.ds_record.values.sha384_digest / c558bd00cfec / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-be7e3459905d3537fc52af9192b3ad9f3833229bc342ff03ad19a9283479fcbc"></a>

## Next pages — primary.default_rr_set_group.ds_record.values.sha384_digest / c558bd00cfec / 5

- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-5ffdc52aab4896b597f8f36c508520f2da9d29ab4e4d129880080ae08be4b810)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-ddc3e22cfd3dd01d590d10bc6f4a41ce27ef3658058ab351d7ea921a4d1b9e33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02da910287d77e3859893ce41aaa29094695ed9a8ba861005617a2ba191cef0e"></a>

## primary.default_rr_set_group.eui48_record — primary.default_rr_set_group.eui48_record / c37b8cc96162 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.eui48_record

<a id="canonical-510e55af08a649aeb199a629b0ba9db14c75629ccaf79c99fc188a72437297e8"></a>

Type: `"single"`. Computed.

Configuration parameter for eui48 record.

Upstream description:

DNS EUI48 Record.

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

<a id="canonical-2321f5b1386f6a60c0946e7f6aae6887731fe6db3d8f0626a3bd1f766f447e4e"></a>

## Direct properties — primary.default_rr_set_group.eui48_record / c37b8cc96162 / 3

<a id="canonical-c74b081322eca48f60f206ae6402f0eded2519426e05ddfb8d80494ca758a82c"></a>

<a id="canonical-39d140610874b575d28c790de49c45d317517ee9a71b6d60f691352e352de58e"></a>

## name property — primary.default_rr_set_group.eui48_record / c37b8cc96162 / 4

Type: `"string"`. Computed.

EUI48 Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-0197220591dbfd0a64529cf640880c999cc2267af43393f69041063dd1fbf081"></a>

<a id="canonical-1dbe3bd149121893fea04b44ff5a9be22e6d341a12229a4799bb2d0f6e925624"></a>

## value property — primary.default_rr_set_group.eui48_record / c37b8cc96162 / 5

Type: `"string"`. Computed.

EUI48 Identifier. A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Upstream description:

A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 17,
  "minLength": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-159a94836c27e7b14c25842aa4e2e3d71e7092b54b3b443d8b6fcc3370b98082"></a>

## Next pages — primary.default_rr_set_group.eui48_record / c37b8cc96162 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-4137c5085f685ae669234993993124f44df8c961c4f967580267db0ada9aff97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb0919966db221951a240efd3ccfce62a89710e447cbe42316ebb23f4d0fb9f6"></a>

## primary.default_rr_set_group.eui64_record — primary.default_rr_set_group.eui64_record / ca256f74bf24 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.eui64_record

<a id="canonical-c0f1cdfefed198669ca4168470c520c647eab55ebd61a1e1b66fffa24659c4d8"></a>

Type: `"single"`. Computed.

Configuration parameter for eui64 record.

Upstream description:

DNS EUI64 Record.

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

<a id="canonical-d1a3bdcc9fada96d0126b83b635b6cf564b9d37f6f769f87378a624c002d8a66"></a>

## Direct properties — primary.default_rr_set_group.eui64_record / ca256f74bf24 / 3

<a id="canonical-54d1d38d99859d69e84f6d4e8bd3f94de95a06a1a6cbc2f97eb5ac5836d1ca02"></a>

<a id="canonical-227492a187bca7e59c6916c5aa9c28a6123c5b59e6376439e729f5365ca8639f"></a>

## name property — primary.default_rr_set_group.eui64_record / ca256f74bf24 / 4

Type: `"string"`. Computed.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-aa93a64babaa6873c6e09c529d1eb3c5dd2c183858bceae849ce1672eb64d971"></a>

<a id="canonical-3e09c34840da656ab8b83cf2965ccc8c09557d77e5cc46863eb707b827a881c8"></a>

## value property — primary.default_rr_set_group.eui64_record / ca256f74bf24 / 5

Type: `"string"`. Computed.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Upstream description:

A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 23,
  "minLength": 23,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 23,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 23,
    "pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-d9192e94a92d0aecea51e21fa5ae42aee60edfdb4db325fad32bf4ff07a5c6f8"></a>

## Next pages — primary.default_rr_set_group.eui64_record / ca256f74bf24 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-7952fecbcf1aa914505a44317d504dfc227553210d64c2d4f31b9a5262998d8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-118b4d2aca0f7d28a10d024a9b9ecaa0a121aec6c84c41af7b51f715c874043f"></a>

## primary.default_rr_set_group.lb_record — primary.default_rr_set_group.lb_record / 3982bce1f66b / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.lb_record

<a id="canonical-f91e65320b4b4381fa10b3eaac701378c1bd08564853254068796e84e3380a22"></a>

Type: `"single"`. Computed.

DNS Load Balancer Record. DNS Load Balancer Record.

Upstream description:

DNS Load Balancer Record.

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

<a id="canonical-5217611c558671fb43edce6695f71699a5fec9bf67622b82f4f45f953b7d674b"></a>

## Direct properties — primary.default_rr_set_group.lb_record / 3982bce1f66b / 3

<a id="canonical-27aabec96ea0ba69970a3773cd48c9f72d7bf3daa0910b8d83eafccb8400addf"></a>

<a id="canonical-3ce27b41b3f32b3c2583fffb7574e0114511f6133f166f8a508824b3bc482022"></a>

## name property — primary.default_rr_set_group.lb_record / 3982bce1f66b / 4

Type: `"string"`. Computed.

Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name
and not a subdomain of a subdomain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
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
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

- [value](data-sources--dns_zone--reference--group-002.md#canonical-b570f30983aa46a48869e306a7954f4b2a6b58d48afc31d910be5a17a460ac9f): complete subsection reference.

<a id="canonical-ef411fed716c82b56838a6b1a53b4f7634594e6d9b91b68040fd1275a435ae8f"></a>

## Next pages — primary.default_rr_set_group.lb_record / 3982bce1f66b / 5

- [primary.default_rr_set_group.lb_record.value](data-sources--dns_zone--reference--group-002.md#canonical-b570f30983aa46a48869e306a7954f4b2a6b58d48afc31d910be5a17a460ac9f)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-b570f30983aa46a48869e306a7954f4b2a6b58d48afc31d910be5a17a460ac9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36172ef46810c8a8bf79d36ff23a5eeae8907e30ae406192ec0200e0144658a9"></a>

## primary.default_rr_set_group.lb_record.value — primary.default_rr_set_group.lb_record.value / 599d291e6973 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.lb_record](data-sources--dns_zone--reference--group-002.md#canonical-7952fecbcf1aa914505a44317d504dfc227553210d64c2d4f31b9a5262998d8b)
- primary.default_rr_set_group.lb_record.value

<a id="canonical-de8c7c93762bffe2dce11df4f3731fda9783de0b393675ae3a93c24719fb898d"></a>

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

<a id="canonical-d0647c282e2589fe583cce7ee9daa87143ffb43e7f8cb94d38d8033bc8c6b459"></a>

## Direct properties — primary.default_rr_set_group.lb_record.value / 599d291e6973 / 3

<a id="canonical-5dab3b4e57ea026ddf2d276d60c2b7f8f52ba2ff88b69203052c0106e9e86674"></a>

<a id="canonical-a02d52bc472abe8f5d7d35be178382306fc19ff68be98433a4177cdda32685ca"></a>

## name property — primary.default_rr_set_group.lb_record.value / 599d291e6973 / 4

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

<a id="canonical-662a50796055482ab34ff0ac4737b0121aaad7b17f95c9591d7b9ac1116be626"></a>

<a id="canonical-2f7b95d4cf4960ce5d5d0c5ccd6d7cb9740083cc065105a4688c61fcfc861c40"></a>

## namespace property — primary.default_rr_set_group.lb_record.value / 599d291e6973 / 5

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

<a id="canonical-ec4bfe0f9400cc1230e0aa89c61dba5bb00ae132566a8be7126065a8394d866f"></a>

<a id="canonical-89813685f57a699cae814bf79d4a9cbe6fd65c7781db9ab6c51d82a6e00d320a"></a>

## tenant property — primary.default_rr_set_group.lb_record.value / 599d291e6973 / 6

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

<a id="canonical-87e5e5667a8b8420a05b348a690f58078f25cb368823282c5c1d49ffe7527900"></a>

## Next pages — primary.default_rr_set_group.lb_record.value / 599d291e6973 / 7

- [primary.default_rr_set_group.lb_record](data-sources--dns_zone--reference--group-002.md#canonical-7952fecbcf1aa914505a44317d504dfc227553210d64c2d4f31b9a5262998d8b)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-f46870a2a1d15407f87baf33c3d6a6fd6909e4329a2462f53cc6f5f502719113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6f5ab2e0088a09a35ac0504a35993f1361600fbd2105e15f05b3d5ae7cada2b"></a>

## primary.default_rr_set_group.loc_record — primary.default_rr_set_group.loc_record / a62e61d28555 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.loc_record

<a id="canonical-f26a5d6ad3d522d0e84fb4219b1155bbdb5f12fe2c0a78473cb480b5b9c05ab6"></a>

Type: `"single"`. Computed.

DNS LOC Record. DNS LOC Record.

Upstream description:

DNS LOC Record.

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

<a id="canonical-c8f5fdffc7a2254175e4dd5b667f78159df0615e5097a24783019b9a832d44c1"></a>

## Direct properties — primary.default_rr_set_group.loc_record / a62e61d28555 / 3

<a id="canonical-53b0121cfd2f282b16a030ab404c66732aa1571b88fe7f73189d294e37d5e26f"></a>

<a id="canonical-58c4874d84815af29da1292bf1697010705bb1cc4115a962a576e89edf05529c"></a>

## name property — primary.default_rr_set_group.loc_record / a62e61d28555 / 4

Type: `"string"`. Computed.

LOC Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-3e93c3510c0fbfc01a31cf1c07d910af6fb1db1230b743033d2184cb1fe35263): complete subsection reference.

<a id="canonical-96f914876f4f9dd0247b0bc37f5bd859fed4950326775c056f68fd72438c2a8d"></a>

## Next pages — primary.default_rr_set_group.loc_record / a62e61d28555 / 5

- [primary.default_rr_set_group.loc_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3e93c3510c0fbfc01a31cf1c07d910af6fb1db1230b743033d2184cb1fe35263)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-3e93c3510c0fbfc01a31cf1c07d910af6fb1db1230b743033d2184cb1fe35263"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-575eb4d5b8a8cae24808b00a9e76a0a3bac5eda5cfb90695ab0373a3810b7485"></a>

## primary.default_rr_set_group.loc_record.values — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.loc_record](data-sources--dns_zone--reference--group-002.md#canonical-f46870a2a1d15407f87baf33c3d6a6fd6909e4329a2462f53cc6f5f502719113)
- primary.default_rr_set_group.loc_record.values

<a id="canonical-d7824b5a29aa37b126335cf43c7df72fa3c93c819e6fa1bd9348470647e0a998"></a>

Type: `"list"`. Computed.

LOC Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-910e7c9cdbc920bcb24aaba09c3d599f1c980f4a098e21a765c53691c6e4dbfe"></a>

## Direct properties — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 3

<a id="canonical-f58209df899332759fb4a1dc3d6275fc8a8c4d95713b2e097726cb10591ceb64"></a>

<a id="canonical-7bbc9257235f0db886a39e41dd9361d5876751067859e3610e08b3b73935b30a"></a>

## altitude property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 4

Type: `"number"`. Computed.

Altitude. Altitude in meters.

Upstream description:

Altitude in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-b0621942b4e793169cb300fe2d2448a2da7387adc01c76812a050bbdc1535e27"></a>

<a id="canonical-cffcc561b713ae7e0bd5c933c9075314022d5bf874f56fe9281d91a2b5debe16"></a>

## horizontal_precision property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 5

Type: `"number"`. Computed.

Horizontal Precision. Horizontal Precision in meters.

Upstream description:

Horizontal Precision in meters.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-d4cd70d24dcbd3393b5377839ddd2b08291c6c15d5fdb478040464e62c314674"></a>

<a id="canonical-d4002ae7c8d34a847716e66911e4ebc9aca496be5291042999dfffe0a9e405cd"></a>

## latitude_degree property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 6

Type: `"number"`. Computed.

Latitude degree, an integer between 0 and 90, including 0 and 90.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 90,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0db02b40642b518b105d14fd51cea1aaee59c3c13870a9365387102812fd0414"></a>

<a id="canonical-2a9f83a00d9344e1daee7e321dbceca1d0660d2bbb0bac5f2dc1b0d9e1f80717"></a>

## latitude_hemisphere property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 7

Type: `"string"`. Computed.

\[Enum: N|S\] Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.
Possible values are \`N\`, \`S\`. Defaults to \`N\`.

Upstream description:

Latitude hemisphere can only be N or S

&#8203;- N: North Hemisphere

&#8203;- S: South Hemisphere.

Receipt-pinned upstream constraints:

```json
{
  "default": "N",
  "enum": [
    "N",
    "S"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b251c2ddd9f00f7143fa75d1dcb7a3f1156b1d4b610508cf4bdd0d65ddbf2d47"></a>

<a id="canonical-7d708ca4d756f5e353330d29f630b18413fd670e0659b09af25a017c948605d0"></a>

## latitude_minute property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 8

Type: `"number"`. Computed.

Latitude minute, an integer between 0 and 59, including 0 and 59.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-752a3ccfd6d90fb1f41d136ce2d55a9722f295ddf8d9bd48afd829e9d0ce6131"></a>

<a id="canonical-7108cdd2334cc605b1957db2eea09bb83c90287910c67f9ab8abd5b23fc04386"></a>

## latitude_second property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 9

Type: `"number"`. Computed.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-7e21432814efe3da25a539df9e2e4e3ee0739cacd77e1306219fe8eb63ad78bc"></a>

<a id="canonical-c0ba5cb42c18847120f52a40b92c048b653ad855f99df91e8873d44f9bb0846e"></a>

## location_diameter property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 10

Type: `"number"`. Computed.

Diameter of a sphere enclosing the described entity, in meters.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-d8b9b3fd307d57e7d597b8e7f492b3505d5725eb256340e6cedb96c602c7337e"></a>

<a id="canonical-885286d2c397cee10e3ec87bd3b46c96bfcf49e0e5d184c17e5f8acebfee8d87"></a>

## longitude_degree property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 11

Type: `"number"`. Computed.

Longitude degree, an integer between 0 and 180, including 0 and 180.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-907acf9a160ca7d155adf5fe57fe334c05b702876451bb03f9980a83e847a97b"></a>

<a id="canonical-7510b48f14bede90c35440d590274cb60dbb8c8802cba57f5c6f4d9af2cec696"></a>

## longitude_hemisphere property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 12

Type: `"string"`. Computed.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

Upstream description:

Longitude hemisphere can only be E or W

&#8203;- E: East Hemisphere

&#8203;- W: West Hemisphere.

Receipt-pinned upstream constraints:

```json
{
  "default": "E",
  "enum": [
    "E",
    "W"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-8b6dfe91e887e332abc022a0651aeabfe7b28fcb11707da5dedba2ea5fc33f90"></a>

<a id="canonical-60b693e45a7ad8542be6d8cdeaed7cc1b5bda99cf51937e3e1b2ca9e04a0a23b"></a>

## longitude_minute property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 13

Type: `"number"`. Computed.

Longitude minute, an integer between 0 and 59, including 0 and 59.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-1cf24ef1b2ca3382c7b196eb860abb57c3d954555f19b05b41a8a9e0ef597e9b"></a>

<a id="canonical-3c2dc92f36ac08d05b9b800e5a620ee4311d3a812f2c9e72725f4203697a9623"></a>

## longitude_second property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 14

Type: `"number"`. Computed.

Longitude second, an decimal between 0 and 59.999, including 0 and 59.999.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-953ed9af85851c83a90ec699477753a201b05463d7448473675a8d8cfd1900f4"></a>

<a id="canonical-f293747920963a4299a48515270d021481d29af7f8ce61a2a1320e3adc0cdbdf"></a>

## vertical_precision property — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 15

Type: `"number"`. Computed.

Vertical Precision. Vertical Precision in meters.

Upstream description:

Vertical Precision in meters.

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
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-8cd6fd7c69888b572a533485f3deaad5211dde0190f4e81ae7220b96830cc486"></a>

## Next pages — primary.default_rr_set_group.loc_record.values / ea6081f5dfef / 16

- [primary.default_rr_set_group.loc_record](data-sources--dns_zone--reference--group-002.md#canonical-f46870a2a1d15407f87baf33c3d6a6fd6909e4329a2462f53cc6f5f502719113)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-64661d589d12a3da3ca8de95e02806b23f24e6d6a88e106698338d1eba3a2566"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e42142c84207d145156e92b089de5729fd603b39b642f8c9097c79fc5284a4c7"></a>

## primary.default_rr_set_group.mx_record — primary.default_rr_set_group.mx_record / fdc7e1f7a859 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.mx_record

<a id="canonical-4ed508b2d3ed35877f102ee7827c5cc62ca029d4439cee685ce931edf38d5c59"></a>

Type: `"single"`. Computed.

DNSMXResourceRecord.

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

<a id="canonical-efa11d158a7d270da14bd31e3d945ff787e5941e1c769481ccf8618cacd1e11a"></a>

## Direct properties — primary.default_rr_set_group.mx_record / fdc7e1f7a859 / 3

<a id="canonical-3f5b1ca0de9b6a9c0aefdd04b306d8ba37f9dbd5a2ea6e0106abc63c88108e69"></a>

<a id="canonical-8a0d6ed1e19085ef86148f5ed045640cb6ff1daaae03e94ab7b431b7c3f5e939"></a>

## name property — primary.default_rr_set_group.mx_record / fdc7e1f7a859 / 4

Type: `"string"`. Computed.

MX Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-4573af663147be813883c97d9a7823ba4388b36389368b30663183ed0ffa6538): complete subsection reference.

<a id="canonical-59c184922947c60788bd21aa4913c8d5543a339cf6a71b563456fc4784114949"></a>

## Next pages — primary.default_rr_set_group.mx_record / fdc7e1f7a859 / 5

- [primary.default_rr_set_group.mx_record.values](data-sources--dns_zone--reference--group-002.md#canonical-4573af663147be813883c97d9a7823ba4388b36389368b30663183ed0ffa6538)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-4573af663147be813883c97d9a7823ba4388b36389368b30663183ed0ffa6538"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f63acc0b717f44aefcd6a2fd9df1a6a35982b634921a49d68a9c497722b1ac26"></a>

## primary.default_rr_set_group.mx_record.values — primary.default_rr_set_group.mx_record.values / ec005b4e8866 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.mx_record](data-sources--dns_zone--reference--group-002.md#canonical-64661d589d12a3da3ca8de95e02806b23f24e6d6a88e106698338d1eba3a2566)
- primary.default_rr_set_group.mx_record.values

<a id="canonical-efbbb9c8523e5b4579896e983fe2c45938c932167dc063660a13874ac189c72f"></a>

Type: `"list"`. Computed.

MX Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

<a id="canonical-f7e80813b4e3a7ca8e1f242bc890a61b74bcfe3ffe6edd8d569770776b0b08ee"></a>

## Direct properties — primary.default_rr_set_group.mx_record.values / ec005b4e8866 / 3

<a id="canonical-e6bd310996f3c449f55db5b8e45987ed5bd2861671fdf08a97d43bc8636950bb"></a>

<a id="canonical-a94f7efe0f48a76f83d167ffee4d67dd1cbd4a679fadec164fc4c9e558410cc0"></a>

## domain property — primary.default_rr_set_group.mx_record.values / ec005b4e8866 / 4

Type: `"string"`. Computed.

Mail exchanger domain name, please provide the full hostname, for.

Upstream description:

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-fa98b2319210b46b5847070ac0d532ece87809e6e4fe7fa5d65fb845feaef3a2"></a>

<a id="canonical-842a35571777e599a5327df3be5e3420dfaded26d7c442fe355c5b7c643c05f2"></a>

## priority property — primary.default_rr_set_group.mx_record.values / ec005b4e8866 / 5

Type: `"number"`. Computed.

Priority. Mail exchanger priority code.

Upstream description:

Mail exchanger priority code.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2798dd82b79b70224202f0e4849131d2a9579997bfe225dfb52d62434cdc34ba"></a>

## Next pages — primary.default_rr_set_group.mx_record.values / ec005b4e8866 / 6

- [primary.default_rr_set_group.mx_record](data-sources--dns_zone--reference--group-002.md#canonical-64661d589d12a3da3ca8de95e02806b23f24e6d6a88e106698338d1eba3a2566)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-80a671af6a9b0043a165a908dfc638927e3a2e31434b16a7a3e00c67065f3853"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fa29aa373d34f856646d58d833a6e7236600c9013f01b469bc1cf7ccca64a18"></a>

## primary.default_rr_set_group.naptr_record — primary.default_rr_set_group.naptr_record / e1380006f7ad / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.naptr_record

<a id="canonical-5d0ea0ba3dd17bad18e52ff0808d262d324c9ea8ee2eb94fbc2d9967b20fd28f"></a>

Type: `"single"`. Computed.

Configuration parameter for naptr record.

Upstream description:

DNS NAPTR Record.

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

<a id="canonical-57cfa283ae96a1a6a64b5c1178d6e35fe183a59a200b389ce473ac92e7f3d841"></a>

## Direct properties — primary.default_rr_set_group.naptr_record / e1380006f7ad / 3

<a id="canonical-7420ce7f114f5f72b31bdd14c39ff37f47231e517ada5ed3c7d1cec2c77ae6c2"></a>

<a id="canonical-45ad489b8906dbb76099e2368f259503cb40db172451048f331a77958a242967"></a>

## name property — primary.default_rr_set_group.naptr_record / e1380006f7ad / 4

Type: `"string"`. Computed.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-1085921b4d5fa75f45e4d783f94df9b7cd09987636b61f3a4fa13af14b94d617): complete subsection reference.

<a id="canonical-53f8095eacf3299a452c0b4079504547d46d5f6fa773fbe29d7653a3e879605e"></a>

## Next pages — primary.default_rr_set_group.naptr_record / e1380006f7ad / 5

- [primary.default_rr_set_group.naptr_record.values](data-sources--dns_zone--reference--group-002.md#canonical-1085921b4d5fa75f45e4d783f94df9b7cd09987636b61f3a4fa13af14b94d617)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-1085921b4d5fa75f45e4d783f94df9b7cd09987636b61f3a4fa13af14b94d617"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca8913d4011d464b58cfad4e22f98cd31e8c01a5fdc26c7bf4f5fc18bb6d4b10"></a>

## primary.default_rr_set_group.naptr_record.values — primary.default_rr_set_group.naptr_record.values / e472c26a425b / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-80a671af6a9b0043a165a908dfc638927e3a2e31434b16a7a3e00c67065f3853)
- primary.default_rr_set_group.naptr_record.values

<a id="canonical-b880a18f4f3ce509128b3264cc366709307f04b6a3413f1fb019a3a6502f6d48"></a>

Type: `"list"`. Computed.

NAPTR Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-15f91a42c96c56b8917d8c080e487cfef9da7ca166cb1f90b51bf17b4cbf818d"></a>

## Direct properties — primary.default_rr_set_group.naptr_record.values / e472c26a425b / 3

<a id="canonical-e8e0a0096f3802ee09f0dd51439085b415ebec22bc4688651d50b3f5b2d11fee"></a>

<a id="canonical-095a5136bbf2790b3ea4dc1c3683d2e4d51d8fe8550924eebc0ed3af52998601"></a>

## flags property — primary.default_rr_set_group.naptr_record.values / e472c26a425b / 4

Type: `"string"`. Computed.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  }
}
```

<a id="canonical-30359ee2cb5ae61f1caaccfe320338c4045df18a866131b9bb42943fa22e6251"></a>

<a id="canonical-deae23bd82708ae7b0a644d868204c108ed41f2e00120b4486f4656d6bdcdda5"></a>

## order property — primary.default_rr_set_group.naptr_record.values / e472c26a425b / 5

Type: `"number"`. Computed.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-85263eacad1fec5dde46c2f4cb03693c7f33258da46e3a9f15881e54b2c95917"></a>

<a id="canonical-e7706f87c7b97a9289b6c6157592818525f86061a59fa365f8c9caf9eaf4c050"></a>

## preference property — primary.default_rr_set_group.naptr_record.values / e472c26a425b / 6

Type: `"number"`. Computed.

Preference when records have the same order. A lower number indicates a higher preference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-fc6632232979040382fe4a6b75a397f5fbef5fcf9a34691e6b48f965ce570f22"></a>

<a id="canonical-7968049605019e5a99743e332a48e333d474087a1a9c7ded99d107bf7a9f591a"></a>

## regexp property — primary.default_rr_set_group.naptr_record.values / e472c26a425b / 7

Type: `"string"`. Computed.

Regular expression to construct the next domain name to lookup.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-cf5c3e94a4c2d10003deb2662612846a405caf97a41e076c7605d7a33d4ce29a"></a>

<a id="canonical-ad126e4638599ada1e69ce191490d9f0fe09848ccfc1512b586567ff83f37252"></a>

## replacement property — primary.default_rr_set_group.naptr_record.values / e472c26a425b / 8

Type: `"string"`. Computed.

The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field.

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

<a id="canonical-53074d828f3d8df7f15274ba604d75591586b42f64143590fd1263d4955d3186"></a>

<a id="canonical-061eb01f0d50d434664d8548268a97895a4977cb0b21ab44dc5d9f755f544616"></a>

## service property — primary.default_rr_set_group.naptr_record.values / e472c26a425b / 9

Type: `"string"`. Computed.

Specifies the service(s) available down this rewrite path.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  }
}
```

<a id="canonical-f73e7ec32f77f668892329856000b33138f1467a4c0312955dbcc0d194c5d982"></a>

## Next pages — primary.default_rr_set_group.naptr_record.values / e472c26a425b / 10

- [primary.default_rr_set_group.naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-80a671af6a9b0043a165a908dfc638927e3a2e31434b16a7a3e00c67065f3853)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-b7b5698521f36fec613ee3585b4ac7a5b57373ed28897f7ccc3859a18a84ba13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-490f0326a78e6e1f3b126dccf605a6331a985f9edb8eb501d53e892830fb1b1f"></a>

## primary.default_rr_set_group.ns_record — primary.default_rr_set_group.ns_record / b11c327315d7 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.ns_record

<a id="canonical-d8175dc81584695aa7f4041c6e47c1ec0c8c1fa080d40238f50857367a1441db"></a>

Type: `"single"`. Computed.

DNSNSResourceRecord.

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

<a id="canonical-4d326ebeaba6193fb665a5e6591b0d85261d40e1cedb4ee9a01b0cc25bec7eec"></a>

## Direct properties — primary.default_rr_set_group.ns_record / b11c327315d7 / 3

<a id="canonical-a7e062b7494f3e8f1537b896f6584ce3c5918638ac8d2815a529c895c4e24625"></a>

<a id="canonical-05c29562dc939189bbe25201fce4d9124c8bc935f977b3df61da142f0b59d9a5"></a>

## name property — primary.default_rr_set_group.ns_record / b11c327315d7 / 4

Type: `"string"`. Computed.

NS Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-30623e12f3e7eafe3b405bc7ecd967c7d1feffad4f49293ade8ce40ee836538b"></a>

<a id="canonical-b848fbff1b8cfd29c89304523ab312424f30137f0f96d4c0350829d2dd0a24bf"></a>

## values property — primary.default_rr_set_group.ns_record / b11c327315d7 / 5

Type: `["list", "string"]`. Computed.

Name Servers. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-44f439649ecb977bf6f9d0b6f9ac6ae99ed4d1f289e203cbf7c499e3f005257e"></a>

## Next pages — primary.default_rr_set_group.ns_record / b11c327315d7 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-0eb30da1bc8d2b099a6bcb576ef2dcb4c120ae431000dcb33809e1abfe6f4ed7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1c9f20fe3e354176a7c96436583a073fd8fb95e0e63a6d9eb20536c663642fc"></a>

## primary.default_rr_set_group.ptr_record — primary.default_rr_set_group.ptr_record / ff7c1f93b1d5 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.ptr_record

<a id="canonical-f315e781da4e7643f5968da45534700de6e050fdfb1881c81853089e46ebcbd3"></a>

Type: `"single"`. Computed.

DNSPTRResourceRecord.

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

<a id="canonical-a71a5ce8828056ed23b0022d4f6a520d928bbfbc53ea69fc8c28928d3c46cb81"></a>

## Direct properties — primary.default_rr_set_group.ptr_record / ff7c1f93b1d5 / 3

<a id="canonical-4b0ad3235582088478973c890faa8570458eae895aa6b9f424bc0f37072bfce4"></a>

<a id="canonical-12e4969740fb02778deadc39b8f32f427f62598a3eba412c3f94fa03d241f58e"></a>

## name property — primary.default_rr_set_group.ptr_record / ff7c1f93b1d5 / 4

Type: `"string"`. Computed.

PTR Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-b7465227216c220b57117a69b1de54aee08b57f9fe26936a395178622058f3d4"></a>

<a id="canonical-532c5126932d71bf5526d2a96f9a7fc57162270a737aaf5ba150060ee1d9ab07"></a>

## values property — primary.default_rr_set_group.ptr_record / ff7c1f93b1d5 / 5

Type: `["list", "string"]`. Computed.

Domain Name. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-95832dcd95aeb8ceea7b172f5ea0127196e5652f16bdf87cfbe0bf4b0cc44fc6"></a>

## Next pages — primary.default_rr_set_group.ptr_record / ff7c1f93b1d5 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-ca0b3e919fd28a9269cfb08289e37290cac4f875d0c586ae19393c9685948dce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a14d5ab484d5aa21897a75375811e4c69e239d945c8f4295f0bc002af3d6860b"></a>

## primary.default_rr_set_group.srv_record — primary.default_rr_set_group.srv_record / 286669a8f078 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.srv_record

<a id="canonical-814b76554d181ac9783b10f0691a158f24b5dfc56b7d7510da06f9706f83e201"></a>

Type: `"single"`. Computed.

DNSSRVResourceRecord.

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

<a id="canonical-fb7b4a9fe6c3bf8812341f9f23cff4ed502a38f3423a3b729349cb1fa68ebfa9"></a>

## Direct properties — primary.default_rr_set_group.srv_record / 286669a8f078 / 3

<a id="canonical-3b710e4d3a4cc37c3eb4bfac827dc2d7d59d0fefdf48c537d8e51bfb6bb20870"></a>

<a id="canonical-d8a7fd950073749de4b1bf857cff9789b3f51c72ffcccc67f592a02bcfe8f278"></a>

## name property — primary.default_rr_set_group.srv_record / 286669a8f078 / 4

Type: `"string"`. Computed.

SRV Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-78ac87c74baff3ccdad071e30f1e8b3b63a302a027d130840a16ac39211311e4): complete subsection reference.

<a id="canonical-70deb661eb64b51f49f70150fbdcc72ddbc0308360c831a20b550098f441b982"></a>

## Next pages — primary.default_rr_set_group.srv_record / 286669a8f078 / 5

- [primary.default_rr_set_group.srv_record.values](data-sources--dns_zone--reference--group-002.md#canonical-78ac87c74baff3ccdad071e30f1e8b3b63a302a027d130840a16ac39211311e4)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-78ac87c74baff3ccdad071e30f1e8b3b63a302a027d130840a16ac39211311e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c29db60ef74235a80bb307088d7aaa6f5b7afc6e4bda0d34af1d4848479023e3"></a>

## primary.default_rr_set_group.srv_record.values — primary.default_rr_set_group.srv_record.values / 3f3035a43fb8 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.srv_record](data-sources--dns_zone--reference--group-002.md#canonical-ca0b3e919fd28a9269cfb08289e37290cac4f875d0c586ae19393c9685948dce)
- primary.default_rr_set_group.srv_record.values

<a id="canonical-c5865dff1e67772d6b443955e7a63965b8c2d3564f98600cf545520ad76b160e"></a>

Type: `"list"`. Computed.

SRV Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-ea56715ecb13801abbba65931749a1e9997b5fbbcf19d6e752cd29977dd02c3d"></a>

## Direct properties — primary.default_rr_set_group.srv_record.values / 3f3035a43fb8 / 3

<a id="canonical-01fa21e199199a8a4852f46e75e76a150b28ab891c01541de93797e825e244f2"></a>

<a id="canonical-48a9aa53dc263f50ae269820211bfc764b986bf3bf8875dacf3d5f6a66f5c1b3"></a>

## port property — primary.default_rr_set_group.srv_record.values / 3f3035a43fb8 / 4

Type: `"number"`. Computed.

Port. Port on which the service can be found.

Upstream description:

Port on which the service can be found.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2367d760698b9569bd7730fcee549527c4b524c203f355f87c184788b8996685"></a>

<a id="canonical-491adb16418ffeb9fdac6ad83e22bec031ffe2cb41f03ece46dfa6555983eed1"></a>

## priority property — primary.default_rr_set_group.srv_record.values / 3f3035a43fb8 / 5

Type: `"number"`. Computed.

Priority of the target. A lower number indicates a higher preference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1c9b0e84099a49fbf1524ea59daa5608b39a23ba3a214568b5f97d267388a223"></a>

<a id="canonical-7951a99bcc7fceef0c418a8532f25d95211271fbae5fd225e1c11ae2af8c68c9"></a>

## target property — primary.default_rr_set_group.srv_record.values / 3f3035a43fb8 / 6

Type: `"string"`. Computed.

Hostname of the machine providing the service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-6d6f490ad4e1333c4f9e59ac5260d93d9b231af9a22e07ea8b311d3bf2afe6f7"></a>

<a id="canonical-cb857eb979096745f7afa5f20383dfa8df54a6e1c99f2c823d670501b34a1cbf"></a>

## weight property — primary.default_rr_set_group.srv_record.values / 3f3035a43fb8 / 7

Type: `"number"`. Computed.

Weight of the target. A higher number indicates a higher preference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-33abbf38164aea4e8255cb4abf05f076e9c36a611950656956ebb0a8959857e8"></a>

## Next pages — primary.default_rr_set_group.srv_record.values / 3f3035a43fb8 / 8

- [primary.default_rr_set_group.srv_record](data-sources--dns_zone--reference--group-002.md#canonical-ca0b3e919fd28a9269cfb08289e37290cac4f875d0c586ae19393c9685948dce)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-9d9c34211610604207a24eec3afe20b57ba6a3ce26657a692a52f41a5a67fc49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac7ae8c780314738c91e8233ccead4cf43a91e2bd81fdc8c1ce9cb65f4e27ca7"></a>

## primary.default_rr_set_group.sshfp_record — primary.default_rr_set_group.sshfp_record / 6c932b197edc / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.sshfp_record

<a id="canonical-871bdf8acf9f220f6b0b5f0412202c9c90df26d383ecc774ae9bcec5e63c8b26"></a>

Type: `"single"`. Computed.

Configuration parameter for sshfp record.

Upstream description:

DNS SSHFP Record.

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

<a id="canonical-0e5498c48fc2d201c9821bd108a751344136e8f46ccb51fb264f867aa8f214c6"></a>

## Direct properties — primary.default_rr_set_group.sshfp_record / 6c932b197edc / 3

<a id="canonical-4b69eafc81e713604cdae06f212dc8909ae82c40d2d3d23ff81debabcc27b62a"></a>

<a id="canonical-500f185490a731e81a605ada83915b5b4551e229c5c020c2a15bc1d6b4ef0b84"></a>

## name property — primary.default_rr_set_group.sshfp_record / 6c932b197edc / 4

Type: `"string"`. Computed.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-e20ea637ed05dea872ce3ba03fc65a7c7644d6c63fe3a0f3c506b8ff4366a1d3): complete subsection reference.

<a id="canonical-feaf086003415ae2986f7330774be4284a316a13d32a68a6df0e8db9ece3c7f4"></a>

## Next pages — primary.default_rr_set_group.sshfp_record / 6c932b197edc / 5

- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-e20ea637ed05dea872ce3ba03fc65a7c7644d6c63fe3a0f3c506b8ff4366a1d3)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-e20ea637ed05dea872ce3ba03fc65a7c7644d6c63fe3a0f3c506b8ff4366a1d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5286ded9cd4893e36094f2bd48caebbb0efc27a6bf29bf031e18c9ad66ca185"></a>

## primary.default_rr_set_group.sshfp_record.values — primary.default_rr_set_group.sshfp_record.values / bccb2dcb9797 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-9d9c34211610604207a24eec3afe20b57ba6a3ce26657a692a52f41a5a67fc49)
- primary.default_rr_set_group.sshfp_record.values

<a id="canonical-0237ff5a4e21c521d28af88d42e828238c7e32eac8d099cf8d5274649c1c5aad"></a>

Type: `"list"`. Computed.

SSHFP Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-780b68d73a8767b82bbc5f583a6c74cfea8cf426ff6a3969e7dde76fffe52015"></a>

## Direct properties — primary.default_rr_set_group.sshfp_record.values / bccb2dcb9797 / 3

<a id="canonical-3330a0c5c8b7ac4c3c5166b13c280ceda3fab596e1b38de14b89252841fc35cc"></a>

<a id="canonical-d6276a9bebeba6efa9b5e6576ead40c1a15af620bdb9f3662214da0f25013b67"></a>

## algorithm property — primary.default_rr_set_group.sshfp_record.values / bccb2dcb9797 / 4

Type: `"string"`. Computed.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

Upstream description:

SSHFP algorithm value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM

&#8203;- RSA: RSA

&#8203;- DSA: DSA

&#8203;- ECDSA: ECDSA

&#8203;- Ed25519: Ed25519

&#8203;- Ed448: Ed448.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIEDALGORITHM",
  "enum": [
    "UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [sha1_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-e6bde4b7c830cf3c709f6e74380d26d58c7afb5eeb2ffa75b6f5b62b86fc025d): complete subsection reference.

- [sha256_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-4f0df00b1a10f70041fdd704b0bb817b486cddec0b23e3ba0adc220089181b86): complete subsection reference.

<a id="canonical-f108286fa970197779e2cb1bdd53b3a8fe4123d78252acaeb8a5c6bff5f06bf0"></a>

## Next pages — primary.default_rr_set_group.sshfp_record.values / bccb2dcb9797 / 5

- [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-e6bde4b7c830cf3c709f6e74380d26d58c7afb5eeb2ffa75b6f5b62b86fc025d)
- [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-4f0df00b1a10f70041fdd704b0bb817b486cddec0b23e3ba0adc220089181b86)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-9d9c34211610604207a24eec3afe20b57ba6a3ce26657a692a52f41a5a67fc49)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-e6bde4b7c830cf3c709f6e74380d26d58c7afb5eeb2ffa75b6f5b62b86fc025d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80a2fd0eedb3908f575f7e999bfb873fa4beda28b8e2a3f95e79e62765e5301c"></a>

## primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint — primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint / 4a141c78c8ac / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-9d9c34211610604207a24eec3afe20b57ba6a3ce26657a692a52f41a5a67fc49)
- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-e20ea637ed05dea872ce3ba03fc65a7c7644d6c63fe3a0f3c506b8ff4366a1d3)
- primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint

<a id="canonical-b12a92e147a6cd9037f0fdede420e6a2b98d797489f70f76046c4380a6c93942"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 fingerprint.

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

<a id="canonical-521faef1700688f93941b6f5f4b0850004a832414707a96636e501737832f05e"></a>

## Direct properties — primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint / 4a141c78c8ac / 3

<a id="canonical-a21af051dd1a6e18a631c9bdbdd28a3d01b2e4616969574b29870720c3d5ebaa"></a>

<a id="canonical-657e6f453e0c3628e1286b29e25ba18a8bf2ad757927e83b3157584259c94ee1"></a>

## fingerprint property — primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint / 4a141c78c8ac / 4

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-4d27c58dc5d98d07004499f584a1f9d79ba82a1269383002ccecd49ea3d3154a"></a>

## Next pages — primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint / 4a141c78c8ac / 5

- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-e20ea637ed05dea872ce3ba03fc65a7c7644d6c63fe3a0f3c506b8ff4366a1d3)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-4f0df00b1a10f70041fdd704b0bb817b486cddec0b23e3ba0adc220089181b86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a9f03dc3fb93fb44d414aa9157fd8938e7e3dc8ce748218b9b9bf01b1a04be4"></a>

## primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint — primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint / 8c9a562ea23c / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-9d9c34211610604207a24eec3afe20b57ba6a3ce26657a692a52f41a5a67fc49)
- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-e20ea637ed05dea872ce3ba03fc65a7c7644d6c63fe3a0f3c506b8ff4366a1d3)
- primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint

<a id="canonical-6193f8e7abb75aa521e3804bd429931cb2d9152431fc4f0f2c2ac5a6a4173379"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 fingerprint.

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

<a id="canonical-d8efcb8ae67d1cf286a5fec2d77c949c36096afdb589d838514402c28c96f52d"></a>

## Direct properties — primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint / 8c9a562ea23c / 3

<a id="canonical-57233d7f1efaa1662e990fefd4533a968b9f42a4b0e6cf17aa384cf20c4b4a6d"></a>

<a id="canonical-0d798b7fbe1fe6fb8e5d60743c39d02d092f14458949313ef19387c2c20e3db0"></a>

## fingerprint property — primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint / 8c9a562ea23c / 4

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 64,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-c1ec5c35192643f28712308645dce81d2060053884d0c53854270638fd45e01a"></a>

## Next pages — primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint / 8c9a562ea23c / 5

- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-e20ea637ed05dea872ce3ba03fc65a7c7644d6c63fe3a0f3c506b8ff4366a1d3)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-91640050b110b8aa2107e7cbd04738b4c1a6f9219edda114d055c15eca75582c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33ffdcd7e106f8329932e8df352ad079e40073db4ad7e064af305f31fd80b91c"></a>

## primary.default_rr_set_group.tlsa_record — primary.default_rr_set_group.tlsa_record / f253acb0603d / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.tlsa_record

<a id="canonical-0e894d29ce9c45ca3f67578ce312abf73327e74eea7572290520ce6331260d72"></a>

Type: `"single"`. Computed.

Configuration parameter for tlsa record.

Upstream description:

DNS TLSA Record.

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

<a id="canonical-3a5d58840f85a6b0f67534f1d710cb62c3dc85e5e7fdfc2d0afc4828813bf731"></a>

## Direct properties — primary.default_rr_set_group.tlsa_record / f253acb0603d / 3

<a id="canonical-6c9cff6a3cb901e4e7a653e4e5ae8f88511183e52f32d034269b08eba4955294"></a>

<a id="canonical-7bb1a8fcd17956eea9b81c45dfc069d800e8888b515b74d61855eb37a2e5d160"></a>

## name property — primary.default_rr_set_group.tlsa_record / f253acb0603d / 4

Type: `"string"`. Computed.

TLSA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-1814d95eecff322e4d36c274173ac2f4f52859264f10547dd681ee4a965d0795): complete subsection reference.

<a id="canonical-dfc3e341940eb0f1e3040a03aef1ef9bbeb67a6f1d4049d82ba57ee69ea45f0f"></a>

## Next pages — primary.default_rr_set_group.tlsa_record / f253acb0603d / 5

- [primary.default_rr_set_group.tlsa_record.values](data-sources--dns_zone--reference--group-002.md#canonical-1814d95eecff322e4d36c274173ac2f4f52859264f10547dd681ee4a965d0795)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-1814d95eecff322e4d36c274173ac2f4f52859264f10547dd681ee4a965d0795"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8ca1958712e2a7a62654a6efca3df665684e71b5d8dbfc59561fe4196c18713"></a>

## primary.default_rr_set_group.tlsa_record.values — primary.default_rr_set_group.tlsa_record.values / 816fa4bf70fe / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [primary.default_rr_set_group.tlsa_record](data-sources--dns_zone--reference--group-002.md#canonical-91640050b110b8aa2107e7cbd04738b4c1a6f9219edda114d055c15eca75582c)
- primary.default_rr_set_group.tlsa_record.values

<a id="canonical-c204f36c2b25620d8db07176a6e442f69cbc114effd68f75a8e81c63cd392648"></a>

Type: `"list"`. Computed.

TLSA Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-00bb8efb605f517b33f930701c2725f640ae12163dddd284157c8dce6d7d83bd"></a>

## Direct properties — primary.default_rr_set_group.tlsa_record.values / 816fa4bf70fe / 3

<a id="canonical-114b90e6d79a2ca194cc9475f3180fb535a2dd921e6a92a6251d52f15543a560"></a>

<a id="canonical-9f94cfe643a80fe5c69be870204f63bba5503beb020a23df942a73b305a3a3da"></a>

## certificate_association_data property — primary.default_rr_set_group.tlsa_record.values / 816fa4bf70fe / 4

Type: `"string"`. Computed.

The actual data to be matched given the settings of the other fields.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1b92c812e7267aa382f4e4615b073a6ffb4225fadc997d50cff0ec646d1889ff"></a>

<a id="canonical-f266f16313ad1cf8215bb0b85429170afc9c312295c857f968b29745f96daa81"></a>

## certificate_usage property — primary.default_rr_set_group.tlsa_record.values / 816fa4bf70fe / 5

Type: `"string"`. Computed.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

Upstream description:

&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint

&#8203;- ServiceCertificateConstraint: Service Certificate Constraint

&#8203;- TrustAnchorAssertion: Trust Anchor Assertion

&#8203;- DomainIssuedCertificate: Domain Issued Certificate.

Receipt-pinned upstream constraints:

```json
{
  "default": "CertificateAuthorityConstraint",
  "enum": [
    "CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5b30af807acae286e906fa64efe9b74d9731d5666003ac2ea9bd6bba5d22f038"></a>

<a id="canonical-dea0f3dc01ba414da2a1bb586dcf5da992e77b2166575417ac3f47b610cd0f4e"></a>

## matching_type property — primary.default_rr_set_group.tlsa_record.values / 816fa4bf70fe / 6

Type: `"string"`. Computed.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

Upstream description:

&#8203;- NoHash: No Hash

&#8203;- SHA256: SHA-256

&#8203;- SHA512: SHA-512.

Receipt-pinned upstream constraints:

```json
{
  "default": "NoHash",
  "enum": [
    "NoHash",
    "SHA256",
    "SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-66d7defcd94756c2fa38b7347f004977252e132e95288142f8ddb8f3727f3962"></a>

<a id="canonical-5365c4099fe37d408ee70dd57115334e6ece1a23b89c9bdee75f74b0d49bb986"></a>

## selector property — primary.default_rr_set_group.tlsa_record.values / 816fa4bf70fe / 7

Type: `"string"`. Computed.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

Upstream description:

&#8203;- FullCertificate: Full Certificate

&#8203;- UseSubjectPublicKey: Use Subject Public Key.

Receipt-pinned upstream constraints:

```json
{
  "default": "FullCertificate",
  "enum": [
    "FullCertificate",
    "UseSubjectPublicKey"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-8d9f6eacc199c820e15f25de1f461543b20f4b381aad8d0a63928c8accb9572c"></a>

## Next pages — primary.default_rr_set_group.tlsa_record.values / 816fa4bf70fe / 8

- [primary.default_rr_set_group.tlsa_record](data-sources--dns_zone--reference--group-002.md#canonical-91640050b110b8aa2107e7cbd04738b4c1a6f9219edda114d055c15eca75582c)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-8b0f9fb60b79b92c52f09acdd60e5210cd48a11f397ffba637e2bb8f2e0cd337"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6c0f671387641db27f0dc6ad5c756bbae2e552fe2e7fa60e9c1e6a36e28fc10"></a>

## primary.default_rr_set_group.txt_record — primary.default_rr_set_group.txt_record / bc25e9f95f85 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- primary.default_rr_set_group.txt_record

<a id="canonical-179f09d346e4e66438ee5eb5977c299da0e9087c81faf4eff3ff5faf195aec94"></a>

Type: `"single"`. Computed.

DNSTXTResourceRecord.

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

<a id="canonical-3e1d6a147be0dbd07ba7219dbe44ae85b5f987074db3ffc2fb0aef11d9907b76"></a>

## Direct properties — primary.default_rr_set_group.txt_record / bc25e9f95f85 / 3

<a id="canonical-dee86190a71fbed997427f9bebf3ba94707e62ea06fd08cee554e1bebf9da25b"></a>

<a id="canonical-4d80e607ef735f22496d5aeebed88d100ad816c4c20e8e8c7ff54b8ee706f4c2"></a>

## name property — primary.default_rr_set_group.txt_record / bc25e9f95f85 / 4

Type: `"string"`. Computed.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2ec5ce4e5b151ef5534facca66026e7f270eacdee5adf8e2555749bfcbc2417b"></a>

<a id="canonical-1f20b315dd8a8399e328ef9a791a4db35c10984933f13bacd61f6eced7b36d4d"></a>

## values property — primary.default_rr_set_group.txt_record / bc25e9f95f85 / 5

Type: `["list", "string"]`. Computed.

Text. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ab9e30a2c6c2f0f6fd12e7648a9259ea5d853553b2a08948727d1633718a88e7"></a>

## Next pages — primary.default_rr_set_group.txt_record / bc25e9f95f85 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-bfbe74fc3289491afc1a88e386a7cde7c86812f89651a0d4b32ab4834ca0b515)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-4739966520d2c572b26ed69d5e165bddefce0cd441f1bdc01a53b0559daf1428"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4a5a3f774f354eac3706d183c2304d72fafab520c8737ffa4819f28229848b0"></a>

## primary.default_soa_parameters — primary.default_soa_parameters / 7f673fee0a3f / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- primary.default_soa_parameters

<a id="canonical-604b9eb5eb88a922c4efad10870c5b604ef6fd5e091b8412d7fabd1911f1ff5a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default soa parameters.

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

<a id="canonical-312804aeb8fa2644cfae3ac7f12edbe8e23d0f35969ae73ad80158d03ff7a3fe"></a>

## Direct properties — primary.default_soa_parameters / 7f673fee0a3f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-71653ac8d0e5828728dadc738724b61e1044507321e6b722131d9f62d28a284e"></a>

## Next pages — primary.default_soa_parameters / 7f673fee0a3f / 4

- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-317b4b9b10202b68289285d0affdfd75878c388455a4d278e28ad28df983a152"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2645215f9863049cda58995d38223a04b037954af2c3a93f9bbaaa32d4e25bf6"></a>

## primary.dnssec_mode — primary.dnssec_mode / a20880c194e7 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- primary.dnssec_mode

<a id="canonical-43ec0c7c7a20efc75929dcdeb54460a21c7c3a901f5d7583999d864e89c44656"></a>

Type: `"single"`. Computed.

DNSSEC Mode.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mode": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-74c220d16fcae0d16271e91aaa2059ff077039d3033f2ccb38684c1af17f5f5a"></a>

## Direct properties — primary.dnssec_mode / a20880c194e7 / 3

- [disable_spec](data-sources--dns_zone--reference--group-002.md#canonical-23ce0fba194fcfe3f5505ba991b40e9f92a9522bdecaeca4674e011c30e58996): complete subsection reference.

- [enable](data-sources--dns_zone--reference--group-002.md#canonical-6965c3e99ba9188b21f86741332b18db1a91e7fea3e5b35090c78498632d3810): complete subsection reference.

<a id="canonical-9538f83be03f3c9e16957a04a421b9ee00dac0493dab52c0be12658d20d99392"></a>

## Next pages — primary.dnssec_mode / a20880c194e7 / 4

- [primary.dnssec_mode.disable_spec](data-sources--dns_zone--reference--group-002.md#canonical-23ce0fba194fcfe3f5505ba991b40e9f92a9522bdecaeca4674e011c30e58996)
- [primary.dnssec_mode.enable](data-sources--dns_zone--reference--group-002.md#canonical-6965c3e99ba9188b21f86741332b18db1a91e7fea3e5b35090c78498632d3810)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-23ce0fba194fcfe3f5505ba991b40e9f92a9522bdecaeca4674e011c30e58996"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7df1221279213121ad2ad767d717fa7e7ebc81be30d2060cf81eac84925566a2"></a>

## primary.dnssec_mode.disable_spec — primary.dnssec_mode.disable_spec / b500263f3e6f / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-317b4b9b10202b68289285d0affdfd75878c388455a4d278e28ad28df983a152)
- primary.dnssec_mode.disable_spec

<a id="canonical-1f44a53646b2983e16addf0adabffbbe33a32a436cad308ad69231a2ff507ee5"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-769d8b591bf7bf3025d7a53173f0f42bac5a7bb40590d5a75f8de5d469861b44"></a>

## Direct properties — primary.dnssec_mode.disable_spec / b500263f3e6f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fbe9516115372d09b5523f8e77d8a10c6897ace42f222210cfa531542f9e3ea4"></a>

## Next pages — primary.dnssec_mode.disable_spec / b500263f3e6f / 4

- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-317b4b9b10202b68289285d0affdfd75878c388455a4d278e28ad28df983a152)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-6965c3e99ba9188b21f86741332b18db1a91e7fea3e5b35090c78498632d3810"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1665f2116e909e74b0885f38afdcf74b052920a9f8e2ec0a8bdde2e7424e1f85"></a>

## primary.dnssec_mode.enable — primary.dnssec_mode.enable / 4a7a11e49ed5 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-317b4b9b10202b68289285d0affdfd75878c388455a4d278e28ad28df983a152)
- primary.dnssec_mode.enable

<a id="canonical-498a22ed638be5660333bfbf9088b47bebcbc8c0ed2897b744446e350fe7dae9"></a>

Type: `["object", {}]`. Computed.

Enable. DNSSEC enable.

Upstream description:

DNSSEC enable.

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

<a id="canonical-ecdc8de14b96ec6b5f92e5250fcfc6e2caf46c01f0bebddba2ac2d01ad66c0c8"></a>

## Direct properties — primary.dnssec_mode.enable / 4a7a11e49ed5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ebe9c2a813b5f81a6e85cf188e0ddb10d6fad9711386881516a209dac12f6319"></a>

## Next pages — primary.dnssec_mode.enable / 4a7a11e49ed5 / 4

- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-317b4b9b10202b68289285d0affdfd75878c388455a4d278e28ad28df983a152)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9cac0b01bdc6e992d2fb305b0abe3fb1165f8738f48a7d5ee67e5e07cb12428"></a>

## primary.rr_set_group — primary.rr_set_group / 58ae8946317d / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- primary.rr_set_group

<a id="canonical-ad8b39448944154c872d1aeff99dc227fee3014b9995e147b635e9205d3427cd"></a>

Type: `"list"`. Computed.

Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed
by F5.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-404709bc46f86f339260952a254d70ab81395358c0beee551e915c5ff2fcb6ef"></a>

## Direct properties — primary.rr_set_group / 58ae8946317d / 3

- [metadata](data-sources--dns_zone--reference--group-002.md#canonical-b00be83ef01b864f7bdacf782a9157b54c17dfd5660adb6a795d77fc2acf4ac7): complete subsection reference.

- [rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32): complete subsection reference.

<a id="canonical-6afdd2186a8f5e212709408557c1163a963a503f221ed9e3e5986f3b627142f8"></a>

## Next pages — primary.rr_set_group / 58ae8946317d / 4

- [primary.rr_set_group.metadata](data-sources--dns_zone--reference--group-002.md#canonical-b00be83ef01b864f7bdacf782a9157b54c17dfd5660adb6a795d77fc2acf4ac7)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-b00be83ef01b864f7bdacf782a9157b54c17dfd5660adb6a795d77fc2acf4ac7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-445f745bd7a88e4a938cf6dcd4c75d6ae59df0d599f5e7a44daf3d8095600a4b"></a>

## primary.rr_set_group.metadata — primary.rr_set_group.metadata / 89252850add2 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- primary.rr_set_group.metadata

<a id="canonical-7c0f3b168b0350ee37250b8d9c425c3c7a4a0d1e15974291fd38004143feb106"></a>

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

<a id="canonical-1a296cf0146d678c394d0d800afdc23dc278e3e7d36f1b7dc09be96f592e474e"></a>

## Direct properties — primary.rr_set_group.metadata / 89252850add2 / 3

<a id="canonical-7af1b8fdb8902aa2e9c75206bcceae664d17a43594b802920e017bd93322f6b5"></a>

<a id="canonical-10380d4a479a5731a3360adf1e8ebae2ec23c65afb9c439df5ae3d569924e5ca"></a>

## description_spec property — primary.rr_set_group.metadata / 89252850add2 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-a22dd40a53caea61c516f0de6a7774bd076d78ce717a7592c2bd80bb5ccd09e9"></a>

<a id="canonical-ad6a9e124efbf5f0cfcf3460fa1497b8b0a6a0d9bf4d65323343cb904f72771c"></a>

## name property — primary.rr_set_group.metadata / 89252850add2 / 5

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

<a id="canonical-28fcca2f4673db4ded92c2eeb9e480f50593470253231be279aa213bcb5f78de"></a>

## Next pages — primary.rr_set_group.metadata / 89252850add2 / 6

- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8292027ce1e1ea363ee0a91abba678d800d04e6bf5401640943347299237c0c7"></a>

## primary.rr_set_group.rr_set — primary.rr_set_group.rr_set / 61a866036b32 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- primary.rr_set_group.rr_set

<a id="canonical-f7986ebe5ed53b2253a23fdcf4a471734fa13c9149000f0f64dc99f00df3d425"></a>

Type: `"list"`. Computed.

Resource Record Sets. Collection of DNS resource record sets.

Upstream description:

Collection of DNS resource record sets.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50000,
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
    "ves.io.schema.rules.repeated.max_items": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  }
}
```

<a id="canonical-2b9f922e7e87100c4e6cf8369b8873872e22d50231f1baf232c4af1d3d8aa893"></a>

## Direct properties — primary.rr_set_group.rr_set / 61a866036b32 / 3

- [a_record](data-sources--dns_zone--reference--group-002.md#canonical-9fabd8eb371bcda6dbef406a928ccbc8f18f1a024d3f562900f4541f38acf86e): complete subsection reference.

- [aaaa_record](data-sources--dns_zone--reference--group-002.md#canonical-30fa33e6925003faf9b87cf4c10670272c1377632e5879ede63f16a7ef47fec2): complete subsection reference.

- [afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-705c043313c473b0bf4aa2fd7513102c7674ffbfc1cf8e748b711ca963cfc1bf): complete subsection reference.

- [alias_record](data-sources--dns_zone--reference--group-002.md#canonical-ae98e64bca679df70737639af32f93458f0c24b02a2e3085f4667617d8dc6214): complete subsection reference.

- [caa_record](data-sources--dns_zone--reference--group-002.md#canonical-c94b3ac935f065f3655ef69823cb1c88ab987182907db6619c10925309033ec5): complete subsection reference.

- [cds_record](data-sources--dns_zone--reference--group-002.md#canonical-fbd2eaf9d60049aaa52fc769966c109aee43ed6b58a66219d1d404dff8c47fd9): complete subsection reference.

- [cert_record](data-sources--dns_zone--reference--group-002.md#canonical-8f71d42d744e17ea048e1e9f170cf5acd7f409e73b6bca15712bb308dbde58f4): complete subsection reference.

- [cname_record](data-sources--dns_zone--reference--group-002.md#canonical-c050f8d5b63b59ba93eb329d5b4a402536e9b23baa7a619540485a17db5d950c): complete subsection reference.

<a id="canonical-80e2a92e70d533ededb9b3bcc1ad906dbdfa6c9063e15863197f31d654aafe4f"></a>

<a id="canonical-072fc299a912da7bbebe88172578fd0c333032a68d6409f135efb01609ed9e7d"></a>

## description_spec property — primary.rr_set_group.rr_set / 61a866036b32 / 4

Type: `"string"`. Computed.

Comment. Human-readable description text

- [ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0308432bcdcbcd22fae0117c0311b7e2cb63181ad8619f0fa3d5ef8a8dbe05bd): complete subsection reference.

- [eui48_record](data-sources--dns_zone--reference--group-003.md#canonical-709880ec57d7e4685ef03f71ff548412c3dde7681ab0dceda336b398bef83fd2): complete subsection reference.

- [eui64_record](data-sources--dns_zone--reference--group-003.md#canonical-03ce7d61c3833deb592f05998294842f569a939a7d46517883fb7e7f9482eed1): complete subsection reference.

- [lb_record](data-sources--dns_zone--reference--group-003.md#canonical-6c20f8137cf306dce88f0b4265f888b4d3298580871c69a5a06e886ebc338659): complete subsection reference.

- [loc_record](data-sources--dns_zone--reference--group-003.md#canonical-d309f505f031ab7a4ff2608e5832c457699f46ca702fa82f161885667890bdc4): complete subsection reference.

- [mx_record](data-sources--dns_zone--reference--group-003.md#canonical-ae20712ac50719e301310b652f0b11851224148a83f26796e8d36e5f7cccfe0a): complete subsection reference.

- [naptr_record](data-sources--dns_zone--reference--group-003.md#canonical-9c832aad21ecb00cad1e4c89f8729406860aecc74fecb638e1b5ec6b98a19504): complete subsection reference.

- [ns_record](data-sources--dns_zone--reference--group-003.md#canonical-71210dc397716cff2a83e2fd4597f220ffe3632ca2ffc86e8e930ce2d50568a3): complete subsection reference.

- [ptr_record](data-sources--dns_zone--reference--group-003.md#canonical-a5cc8ab5f4efe1232ec6fc54d96daaf2d33d193a886c5095cede44588163b3fc): complete subsection reference.

- [srv_record](data-sources--dns_zone--reference--group-003.md#canonical-2b4cb5841ad7eb27081233d9bb8cd9f4791d4ce5744a828ddc7f86377f1446b6): complete subsection reference.

- [sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-7886f28784c88ae4eb98416232cc1f6e4ad60fe9cc16f37da1a1d0b6af6ae43e): complete subsection reference.

- [tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-585687c64fe7863dd11f33eab30674fb6cda389e728007a92787a434b499b119): complete subsection reference.

<a id="canonical-17eceb452bfeaca5fac0370c3b5c5754376f183ece5ec61a3b50b99016a4f79d"></a>

<a id="canonical-37d94d4265c9dbde88cbd8f96b95b6131790313a3082ba4e231df0fd8225a491"></a>

## ttl property — primary.rr_set_group.rr_set / 61a866036b32 / 5

Type: `"number"`. Computed.

Time to live. Time-to-live duration in seconds

Upstream description:

Time-to-live duration in seconds

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

- [txt_record](data-sources--dns_zone--reference--group-003.md#canonical-960100e50468b23cb2c3dfe4db63f18677dc3da79863cd6d4fe4365a0c756607): complete subsection reference.

<a id="canonical-ee35d29a3130c4f916f3825f8c4db654a8c84cf9c0c32062eeee639765e2fd40"></a>

## Next pages — primary.rr_set_group.rr_set / 61a866036b32 / 6

- [primary.rr_set_group.rr_set.a_record](data-sources--dns_zone--reference--group-002.md#canonical-9fabd8eb371bcda6dbef406a928ccbc8f18f1a024d3f562900f4541f38acf86e)
- [primary.rr_set_group.rr_set.aaaa_record](data-sources--dns_zone--reference--group-002.md#canonical-30fa33e6925003faf9b87cf4c10670272c1377632e5879ede63f16a7ef47fec2)
- [primary.rr_set_group.rr_set.afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-705c043313c473b0bf4aa2fd7513102c7674ffbfc1cf8e748b711ca963cfc1bf)
- [primary.rr_set_group.rr_set.alias_record](data-sources--dns_zone--reference--group-002.md#canonical-ae98e64bca679df70737639af32f93458f0c24b02a2e3085f4667617d8dc6214)
- [primary.rr_set_group.rr_set.caa_record](data-sources--dns_zone--reference--group-002.md#canonical-c94b3ac935f065f3655ef69823cb1c88ab987182907db6619c10925309033ec5)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-fbd2eaf9d60049aaa52fc769966c109aee43ed6b58a66219d1d404dff8c47fd9)
- [primary.rr_set_group.rr_set.cert_record](data-sources--dns_zone--reference--group-002.md#canonical-8f71d42d744e17ea048e1e9f170cf5acd7f409e73b6bca15712bb308dbde58f4)
- [primary.rr_set_group.rr_set.cname_record](data-sources--dns_zone--reference--group-002.md#canonical-c050f8d5b63b59ba93eb329d5b4a402536e9b23baa7a619540485a17db5d950c)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0308432bcdcbcd22fae0117c0311b7e2cb63181ad8619f0fa3d5ef8a8dbe05bd)
- [primary.rr_set_group.rr_set.eui48_record](data-sources--dns_zone--reference--group-003.md#canonical-709880ec57d7e4685ef03f71ff548412c3dde7681ab0dceda336b398bef83fd2)
- [primary.rr_set_group.rr_set.eui64_record](data-sources--dns_zone--reference--group-003.md#canonical-03ce7d61c3833deb592f05998294842f569a939a7d46517883fb7e7f9482eed1)
- [primary.rr_set_group.rr_set.lb_record](data-sources--dns_zone--reference--group-003.md#canonical-6c20f8137cf306dce88f0b4265f888b4d3298580871c69a5a06e886ebc338659)
- [primary.rr_set_group.rr_set.loc_record](data-sources--dns_zone--reference--group-003.md#canonical-d309f505f031ab7a4ff2608e5832c457699f46ca702fa82f161885667890bdc4)
- [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--reference--group-003.md#canonical-ae20712ac50719e301310b652f0b11851224148a83f26796e8d36e5f7cccfe0a)
- [primary.rr_set_group.rr_set.naptr_record](data-sources--dns_zone--reference--group-003.md#canonical-9c832aad21ecb00cad1e4c89f8729406860aecc74fecb638e1b5ec6b98a19504)
- [primary.rr_set_group.rr_set.ns_record](data-sources--dns_zone--reference--group-003.md#canonical-71210dc397716cff2a83e2fd4597f220ffe3632ca2ffc86e8e930ce2d50568a3)
- [primary.rr_set_group.rr_set.ptr_record](data-sources--dns_zone--reference--group-003.md#canonical-a5cc8ab5f4efe1232ec6fc54d96daaf2d33d193a886c5095cede44588163b3fc)
- [primary.rr_set_group.rr_set.srv_record](data-sources--dns_zone--reference--group-003.md#canonical-2b4cb5841ad7eb27081233d9bb8cd9f4791d4ce5744a828ddc7f86377f1446b6)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-7886f28784c88ae4eb98416232cc1f6e4ad60fe9cc16f37da1a1d0b6af6ae43e)
- [primary.rr_set_group.rr_set.tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-585687c64fe7863dd11f33eab30674fb6cda389e728007a92787a434b499b119)
- [primary.rr_set_group.rr_set.txt_record](data-sources--dns_zone--reference--group-003.md#canonical-960100e50468b23cb2c3dfe4db63f18677dc3da79863cd6d4fe4365a0c756607)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-9fabd8eb371bcda6dbef406a928ccbc8f18f1a024d3f562900f4541f38acf86e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b5d413ff7d30033fc21aee55236f3ff51ae2e9e39dede956ed32c3f3d0eb135"></a>

## primary.rr_set_group.rr_set.a_record — primary.rr_set_group.rr_set.a_record / 4cc05f9b385f / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.a_record

<a id="canonical-c59b8540e2fc762909995f8b8e6ee395215b462bf481bfd626ae8281815f610d"></a>

Type: `"single"`. Computed.

DNSAResourceRecord. A Records

Upstream description:

A Records

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

<a id="canonical-099de1197a5df39674e6bf8d97a3327d78326dbcb5efee82ad354a8485e848c7"></a>

## Direct properties — primary.rr_set_group.rr_set.a_record / 4cc05f9b385f / 3

<a id="canonical-e74addcf33f7a642f0d8f59c215765864338b525abc3fed865eb2d8789eacf2c"></a>

<a id="canonical-b9b878847c05d358546f3dd5cd4ade81d06383a28e50572c86e4070a2234d7ec"></a>

## name property — primary.rr_set_group.rr_set.a_record / 4cc05f9b385f / 4

Type: `"string"`. Computed.

Record name, please provide only the specific subdomain or record name without the base domain.

Upstream description:

A Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-16ed710bcb7475cdd2761e8daf5c2826dcffedcac2a00a7d44db31454b24c5f0"></a>

<a id="canonical-f90b0e18f499a5157ab006d7a860f26c3dd72fc1ac8fdce7fccb1eb7c726deba"></a>

## values property — primary.rr_set_group.rr_set.a_record / 4cc05f9b385f / 5

Type: `["list", "string"]`. Computed.

IPv4 Addresses. A valid IPv4 address, for example: 192.0.2.242.

Upstream description:

A valid IPv4 address, for example: 192.0.2.242.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-32ba2e18581641ee667065015187dabdd390b06c01f9fe1f5cf16f0d182e0ce3"></a>

## Next pages — primary.rr_set_group.rr_set.a_record / 4cc05f9b385f / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-30fa33e6925003faf9b87cf4c10670272c1377632e5879ede63f16a7ef47fec2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4f808ae23114662698628c5beecfb8b15481f5da9984020c968868931fba6b5a"></a>

## primary.rr_set_group.rr_set.aaaa_record — primary.rr_set_group.rr_set.aaaa_record / ee2bf25b15e1 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.aaaa_record

<a id="canonical-8bf46d66a2b3313ca866de0f05be88da6a316c785db81ef4b55557dc1e624639"></a>

Type: `"single"`. Computed.

Configuration parameter for aaaa record.

Upstream description:

RecordSet for AAAA Records.

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

<a id="canonical-bbbff47ec42e0502a022254c7af397ec79f9f729b4f3ea5310efa7a3b74876b9"></a>

## Direct properties — primary.rr_set_group.rr_set.aaaa_record / ee2bf25b15e1 / 3

<a id="canonical-1e700ec01892ce573ec33cd19df4e054138f54dcddb9743ad33c5ab1e3350f34"></a>

<a id="canonical-3e6d9e8853a4eb7019196345305c0875e9a560a9f31df46dee85c2e60ce27e0c"></a>

## name property — primary.rr_set_group.rr_set.aaaa_record / ee2bf25b15e1 / 4

Type: `"string"`. Computed.

AAAA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-6190d72e5925ef8e108017d687a069cf393ff5fe82cf9852d3307b13ce161e89"></a>

<a id="canonical-5e45d5b93bb7fb70cc8275698198ed803f0ddd9266e1c6c8bc3f64b7eb1f9517"></a>

## values property — primary.rr_set_group.rr_set.aaaa_record / ee2bf25b15e1 / 5

Type: `["list", "string"]`. Computed.

IPv6 Addresses. A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Upstream description:

A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e68be4ba76790dc964512837fcaf820c0fc9c0c62f06e00651a29cbaae96268a"></a>

## Next pages — primary.rr_set_group.rr_set.aaaa_record / ee2bf25b15e1 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-705c043313c473b0bf4aa2fd7513102c7674ffbfc1cf8e748b711ca963cfc1bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd0989ec42b1dbc46a1ef482f777f40966f69e623a8f53bb38fc2681f312945b"></a>

## primary.rr_set_group.rr_set.afsdb_record — primary.rr_set_group.rr_set.afsdb_record / 5cc9783e15d7 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.afsdb_record

<a id="canonical-cfcaa49dcf9066a96ba20bd9808f1c1bb54dbf262798671a2c87814e2d792e7c"></a>

Type: `"single"`. Computed.

Configuration parameter for afsdb record.

Upstream description:

DNS AFSDB Record.

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

<a id="canonical-b899ed022f4246b334601791fe0d928ac6f8f56f1ace7dc18ad275e0bdc4c416"></a>

## Direct properties — primary.rr_set_group.rr_set.afsdb_record / 5cc9783e15d7 / 3

<a id="canonical-0ec6f6220fb57416d5f0bc66a011c58a3da429e50fc14824bacaa6be705e296a"></a>

<a id="canonical-5bbfc14fec8b2de2971207b453ff385443103fde400718e8c6f08eb5d75daef4"></a>

## name property — primary.rr_set_group.rr_set.afsdb_record / 5cc9783e15d7 / 4

Type: `"string"`. Computed.

AFSDB Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-b1d7cee9be433b241149084f63bdca378c8549e91d86584f516088afd5ed3f61): complete subsection reference.

<a id="canonical-ad37acd6634f7154c9b2ebf004f18923f236e187e00d4a8c9dd4bc351b5dd306"></a>

## Next pages — primary.rr_set_group.rr_set.afsdb_record / 5cc9783e15d7 / 5

- [primary.rr_set_group.rr_set.afsdb_record.values](data-sources--dns_zone--reference--group-002.md#canonical-b1d7cee9be433b241149084f63bdca378c8549e91d86584f516088afd5ed3f61)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-b1d7cee9be433b241149084f63bdca378c8549e91d86584f516088afd5ed3f61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af1e8d7c65c421acedb9c5efe6b2146bd92a5c84541322b904698bc8a8177f57"></a>

## primary.rr_set_group.rr_set.afsdb_record.values — primary.rr_set_group.rr_set.afsdb_record.values / d44a8430d737 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-705c043313c473b0bf4aa2fd7513102c7674ffbfc1cf8e748b711ca963cfc1bf)
- primary.rr_set_group.rr_set.afsdb_record.values

<a id="canonical-4b060c621406e0c089c13fc11b54200e78848c9f6f974b04aff5c45c4ef36ea2"></a>

Type: `"list"`. Computed.

AFSDB Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-8809904458bd8b99ec3feef22bb85d43f6ef03d0e21029def00afc1f5c7ce57c"></a>

## Direct properties — primary.rr_set_group.rr_set.afsdb_record.values / d44a8430d737 / 3

<a id="canonical-54b52e1de5eebc9831c7c4c8d23358c174fe92afa1601a654ddab7328083b74a"></a>

<a id="canonical-062b415850bec9b3986a8531af7ab3e8dc5f535d2552b2b18ae2885b5976482b"></a>

## hostname property — primary.rr_set_group.rr_set.afsdb_record.values / d44a8430d737 / 4

Type: `"string"`. Computed.

Server name of the AFS cell database server or the DCE name server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-f4ca63f5500cb1fcde93a0b41995b85c88603159641f8507ad8e4845fbab7b6c"></a>

<a id="canonical-44bf1683168baa4a9bea2c796f3494408a6e97b198bb36b3cd4ee92a3725a2bb"></a>

## subtype property — primary.rr_set_group.rr_set.afsdb_record.values / d44a8430d737 / 5

Type: `"string"`. Computed.

\[Enum: NONE|AFSVolumeLocationServer|DCEAuthenticationServer\] AFS Volume Location Server or DCE
Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server -
DCEAuthenticationServer: DCE Authentication Server. Possible values are \`NONE\`,
\`AFSVolumeLocationServer\`, \`DCEAuthenticationServer\`.

Upstream description:

AFS Volume Location Server or DCE Authentication Server.

&#8203;- NONE: NONE

&#8203;- AFSVolumeLocationServer: AFS Volume Location Server

&#8203;- DCEAuthenticationServer: DCE Authentication Server.

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d31582e75fc54f8d88dc6aee50fac1a3cc09a84f401a8ea2cd047fc0008a67d9"></a>

## Next pages — primary.rr_set_group.rr_set.afsdb_record.values / d44a8430d737 / 6

- [primary.rr_set_group.rr_set.afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-705c043313c473b0bf4aa2fd7513102c7674ffbfc1cf8e748b711ca963cfc1bf)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-ae98e64bca679df70737639af32f93458f0c24b02a2e3085f4667617d8dc6214"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c79852c52dce3f7568640974ce63916670c6635ecf7ee12bcc4e694852ae3f4d"></a>

## primary.rr_set_group.rr_set.alias_record — primary.rr_set_group.rr_set.alias_record / 65212491af9b / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.alias_record

<a id="canonical-fc47f5d66110adfac70f5119899f9d7178afed36f72c82123fa9edcb77280f5f"></a>

Type: `"single"`. Computed.

Configuration parameter for alias record.

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

<a id="canonical-3213a2de895d3b164464aa120410f7943310e55231ebe8e1be49a72d45d29078"></a>

## Direct properties — primary.rr_set_group.rr_set.alias_record / 65212491af9b / 3

<a id="canonical-19486adac784aa364f7eac6e2d8b975cefdbdf4afc2d8ccc3e75ec6af1469cfc"></a>

<a id="canonical-628eae6781b252aae26337c59e61026b7b8130cfb91ee6ec6580ec4c0e8fc995"></a>

## value property — primary.rr_set_group.rr_set.alias_record / 65212491af9b / 4

Type: `"string"`. Computed.

Domain. A valid domain name, for example: example.com.

Upstream description:

A valid domain name, for example: example.com.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-f173fb0db028498c22b091b7df26f44e87e22301196e3de6b679aad9d5372863"></a>

## Next pages — primary.rr_set_group.rr_set.alias_record / 65212491af9b / 5

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-c94b3ac935f065f3655ef69823cb1c88ab987182907db6619c10925309033ec5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ce2c65d3f2c6220b72df1ec0d3c165fdcf5fd1f8d2489af57bae4c284016f4e"></a>

## primary.rr_set_group.rr_set.caa_record — primary.rr_set_group.rr_set.caa_record / 611589499d2e / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.caa_record

<a id="canonical-fc5918f983f498ae444bb488cc4a135f8e8ff33c46c4ad23d2204b3399d43df5"></a>

Type: `"single"`. Computed.

DNSCAAResourceRecord.

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

<a id="canonical-0c8af430c85e9303f485515a84ca4360f4943f7ebae7f9939f3b6bf88305fb69"></a>

## Direct properties — primary.rr_set_group.rr_set.caa_record / 611589499d2e / 3

<a id="canonical-e32c4a259cc8d60d06f4bc7ff4cea62e974c47dcb50798481e2d2d0f4fc387f4"></a>

<a id="canonical-986b64b35a947ed438723177c0d139067ce7539454907ed1ed88bbc1091c0c40"></a>

## name property — primary.rr_set_group.rr_set.caa_record / 611589499d2e / 4

Type: `"string"`. Computed.

CAA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-8f0ad42f9cdc99d20c3bf312f171aa2f04663270d17c65c87edaa776b110975d): complete subsection reference.

<a id="canonical-cca33fed4d3381f16f075773e16a1dc54f7a380ac1132bb590e8c2f29722bbdd"></a>

## Next pages — primary.rr_set_group.rr_set.caa_record / 611589499d2e / 5

- [primary.rr_set_group.rr_set.caa_record.values](data-sources--dns_zone--reference--group-002.md#canonical-8f0ad42f9cdc99d20c3bf312f171aa2f04663270d17c65c87edaa776b110975d)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-8f0ad42f9cdc99d20c3bf312f171aa2f04663270d17c65c87edaa776b110975d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb429feae286650ca3f41d103a014b6ec9eae4d039f96c2f789c2f31904e4f8c"></a>

## primary.rr_set_group.rr_set.caa_record.values — primary.rr_set_group.rr_set.caa_record.values / fd7dbc12c121 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.caa_record](data-sources--dns_zone--reference--group-002.md#canonical-c94b3ac935f065f3655ef69823cb1c88ab987182907db6619c10925309033ec5)
- primary.rr_set_group.rr_set.caa_record.values

<a id="canonical-af0bb68d3b5ed4effec1cb08b11954d1229ecd8d0a3ab9d3a8a4df6aaa9cb44f"></a>

Type: `"list"`. Computed.

CAA Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

<a id="canonical-3e136d4f4023edfb2c86a7bee24002e08afdabed6954fe767e5b3a04ac5b7d51"></a>

## Direct properties — primary.rr_set_group.rr_set.caa_record.values / fd7dbc12c121 / 3

<a id="canonical-a5e5f9944ddee3ded3f47a1633918bd93125d1cee1057642c91a66b66e4e84fa"></a>

<a id="canonical-75c0ef939fdc902952f0e0cefde8d11261247f6d3b88690a058474cde87d09fc"></a>

## flags property — primary.rr_set_group.rr_set.caa_record.values / fd7dbc12c121 / 4

Type: `"number"`. Computed.

Flag should be an integer between 0 and 255.

Upstream description:

This flag should be an integer between 0 and 255.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-5193bad77dae77bb7a831114801917b79ae9bc658b04bec4be1c9354e93a0ca8"></a>

<a id="canonical-9401d389a5536ce1c8e8ff8fe15a638be1b4f488012cf77779cfbef861815833"></a>

## tag property — primary.rr_set_group.rr_set.caa_record.values / fd7dbc12c121 / 5

Type: `"string"`. Computed.

\[Enum: issue|issuewild|iodef\] Tag. Tag for categorization and filtering. Possible values are
\`issue\`, \`issuewild\`, \`iodef\`.

Upstream description:

Tag for categorization and filtering

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "issue",
    "issuewild",
    "iodef"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  }
}
```

<a id="canonical-2cde9979fbc4c4ee594962c2e1421295bf7e81e5d82091f1b57616ffe8473cff"></a>

<a id="canonical-0d6b598133e5527d9cc9c64782b0ad3d536d936720b0ac90c22c1e99c6e34396"></a>

## value property — primary.rr_set_group.rr_set.caa_record.values / fd7dbc12c121 / 6

Type: `"string"`. Computed.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-c102041d4faace364fc0fe9483d9be0b6a0c51e6fcf5fcdd14bd0c7999828f4f"></a>

## Next pages — primary.rr_set_group.rr_set.caa_record.values / fd7dbc12c121 / 7

- [primary.rr_set_group.rr_set.caa_record](data-sources--dns_zone--reference--group-002.md#canonical-c94b3ac935f065f3655ef69823cb1c88ab987182907db6619c10925309033ec5)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-fbd2eaf9d60049aaa52fc769966c109aee43ed6b58a66219d1d404dff8c47fd9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7956b8647d8dfc04ff59d5130e98158b73c308cdd90492964835f873f711667"></a>

## primary.rr_set_group.rr_set.cds_record — primary.rr_set_group.rr_set.cds_record / 9befae38dc60 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.cds_record

<a id="canonical-d64aa7eac3a3d580782774a205c26e5d97134d25e480706c04d2fed4c6cc66e0"></a>

Type: `"single"`. Computed.

DNS CDS Record. DNS CDS Record.

Upstream description:

DNS CDS Record.

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

<a id="canonical-3857fd47a8317517276a68c713f3727c59bf57e02a3aebd5eb176f8ddb671b86"></a>

## Direct properties — primary.rr_set_group.rr_set.cds_record / 9befae38dc60 / 3

<a id="canonical-f25c900cce87444866f0249de20c7efbdbeb89832e4d9142a385075b4d030a70"></a>

<a id="canonical-64fe70ecdc07036080fcde68a1efb560afff9d4716a4118a0bda62526351f565"></a>

## name property — primary.rr_set_group.rr_set.cds_record / 9befae38dc60 / 4

Type: `"string"`. Computed.

CDS Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-8d4c1e83de3d0939f708128ed6bd6f6a1d1d6262ce26581661a12381192ebfdb): complete subsection reference.

<a id="canonical-55c3a25957e6780caebe9c3e00ecf70e101f46b2c16ba12217af0aab90129c6b"></a>

## Next pages — primary.rr_set_group.rr_set.cds_record / 9befae38dc60 / 5

- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-8d4c1e83de3d0939f708128ed6bd6f6a1d1d6262ce26581661a12381192ebfdb)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-8d4c1e83de3d0939f708128ed6bd6f6a1d1d6262ce26581661a12381192ebfdb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82b91072e84ea397e04ea70ee17137b59235d2117c4bc7b1b460b321267f3e8f"></a>

## primary.rr_set_group.rr_set.cds_record.values — primary.rr_set_group.rr_set.cds_record.values / 54eaba6cd440 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-fbd2eaf9d60049aaa52fc769966c109aee43ed6b58a66219d1d404dff8c47fd9)
- primary.rr_set_group.rr_set.cds_record.values

<a id="canonical-3b734a95822c51ae245f24a56f83a3fc17571058a321d9687a3023f6de417077"></a>

Type: `"list"`. Computed.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-42e4d3974bc0c33674766475512f2b9fe1785e44295ca07bcffc2559d5362f4f"></a>

## Direct properties — primary.rr_set_group.rr_set.cds_record.values / 54eaba6cd440 / 3

<a id="canonical-7521f90f38d90c4f2033716abf56d5f57533acf551dca88b95a88f538c295e74"></a>

<a id="canonical-af4948b8f4d20071dd5f39165ded07e40278653f151ec1636a0aac1020768780"></a>

## ds_key_algorithm property — primary.rr_set_group.rr_set.cds_record.values / 54eaba6cd440 / 4

Type: `"string"`. Computed.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d74ed67bed30db1f7081a20ffceccdfcc559a56383caa189ba5ba753719ebdce"></a>

<a id="canonical-49de82b14d54ba9005fcb56e84dc87d9ba2b89931f2f167144d30cf9f7ce007a"></a>

## key_tag property — primary.rr_set_group.rr_set.cds_record.values / 54eaba6cd440 / 5

Type: `"number"`. Computed.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-69d3bbc08896960e888062b22784a0b9ca1bbd2ccb782dbd2aa2e44c1fcc1711): complete subsection reference.

- [sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-30b55b80c3ced9e1ed35f11897e45013572a08708e9f1e4de3cbcc37ec74d7e0): complete subsection reference.

- [sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-b314911cebe11d2fbeefaa2d405d7aa9f9ae95ebf10d89ccb38a2231d5970907): complete subsection reference.

<a id="canonical-c40c45a62ef764bdef2bc8f948dc8846708b52c11986b24209c8ae7bf2648de3"></a>

## Next pages — primary.rr_set_group.rr_set.cds_record.values / 54eaba6cd440 / 6

- [primary.rr_set_group.rr_set.cds_record.values.sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-69d3bbc08896960e888062b22784a0b9ca1bbd2ccb782dbd2aa2e44c1fcc1711)
- [primary.rr_set_group.rr_set.cds_record.values.sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-30b55b80c3ced9e1ed35f11897e45013572a08708e9f1e4de3cbcc37ec74d7e0)
- [primary.rr_set_group.rr_set.cds_record.values.sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-b314911cebe11d2fbeefaa2d405d7aa9f9ae95ebf10d89ccb38a2231d5970907)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-fbd2eaf9d60049aaa52fc769966c109aee43ed6b58a66219d1d404dff8c47fd9)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-69d3bbc08896960e888062b22784a0b9ca1bbd2ccb782dbd2aa2e44c1fcc1711"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-294d5089e1cb31f01d81f1aa5eb99028fe746affe01531546b90219fa321be61"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha1_digest — primary.rr_set_group.rr_set.cds_record.values.sha1_digest / ae1f4a537091 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-fbd2eaf9d60049aaa52fc769966c109aee43ed6b58a66219d1d404dff8c47fd9)
- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-8d4c1e83de3d0939f708128ed6bd6f6a1d1d6262ce26581661a12381192ebfdb)
- primary.rr_set_group.rr_set.cds_record.values.sha1_digest

<a id="canonical-b92a26c82a6684dbaa7704a01309f75e2440658e5208bceb46322002690bc4dc"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 digest.

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

<a id="canonical-1e58f19377933e0ef5815668ecd602c207f4c4a92a27d6bbb79440c0d95104a0"></a>

## Direct properties — primary.rr_set_group.rr_set.cds_record.values.sha1_digest / ae1f4a537091 / 3

<a id="canonical-9fe80f480c43a16430fbbbb33f5337496456b3533213fa5d2798e168e1bacf1a"></a>

<a id="canonical-91c2500a7a1b0a5b9a2428bbfa2f370fb2dc32048d8992873c97de80d26c8c59"></a>

## digest property — primary.rr_set_group.rr_set.cds_record.values.sha1_digest / ae1f4a537091 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-6ccad151a0471fda6739541a9a5d28f2f1b76b976db62ee1b66ebe1bfe3ddc32"></a>

## Next pages — primary.rr_set_group.rr_set.cds_record.values.sha1_digest / ae1f4a537091 / 5

- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-8d4c1e83de3d0939f708128ed6bd6f6a1d1d6262ce26581661a12381192ebfdb)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-30b55b80c3ced9e1ed35f11897e45013572a08708e9f1e4de3cbcc37ec74d7e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72195ad837f0d178dbc93a5e34155e760dfff4bf04f11d15b93b28188910e521"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha256_digest — primary.rr_set_group.rr_set.cds_record.values.sha256_digest / 1f44ea7cdb80 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-fbd2eaf9d60049aaa52fc769966c109aee43ed6b58a66219d1d404dff8c47fd9)
- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-8d4c1e83de3d0939f708128ed6bd6f6a1d1d6262ce26581661a12381192ebfdb)
- primary.rr_set_group.rr_set.cds_record.values.sha256_digest

<a id="canonical-808a206664a36f7c1475f3a670d7bc94d19638d1b9073418f6ba46493ea1001f"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 digest.

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

<a id="canonical-5b1597f436b435bc11baae121720641eb056dcc980a26c9742fc7a0bc63168e0"></a>

## Direct properties — primary.rr_set_group.rr_set.cds_record.values.sha256_digest / 1f44ea7cdb80 / 3

<a id="canonical-a23c208940549a0d7d3e46aff7340025c0c2f5d295b3c6feac221cad0b093add"></a>

<a id="canonical-5f1ff3f99eedd0de3e45f92c7c79e0bc8d771516894ad6d44746771c4f2db981"></a>

## digest property — primary.rr_set_group.rr_set.cds_record.values.sha256_digest / 1f44ea7cdb80 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 64
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-0166d29a874fc45b8f90110957dfecf92c8b5bd9456cbb40850b1e0e05ccfe2f"></a>

## Next pages — primary.rr_set_group.rr_set.cds_record.values.sha256_digest / 1f44ea7cdb80 / 5

- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-8d4c1e83de3d0939f708128ed6bd6f6a1d1d6262ce26581661a12381192ebfdb)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-b314911cebe11d2fbeefaa2d405d7aa9f9ae95ebf10d89ccb38a2231d5970907"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e8b774b85d1a7c3b2fddc093f35d79dab667d3eaa6bea7ec74e2dc83e58217fe"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha384_digest — primary.rr_set_group.rr_set.cds_record.values.sha384_digest / 4b822d376d30 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-fbd2eaf9d60049aaa52fc769966c109aee43ed6b58a66219d1d404dff8c47fd9)
- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-8d4c1e83de3d0939f708128ed6bd6f6a1d1d6262ce26581661a12381192ebfdb)
- primary.rr_set_group.rr_set.cds_record.values.sha384_digest

<a id="canonical-863f8fe478da1ff43dcaabb543f724ae4e46d5926d249b11ff85e2552e5c4978"></a>

Type: `"single"`. Computed.

Configuration parameter for sha384 digest.

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

<a id="canonical-669b713f3a50354c0e143dccbd67bb8ae9afd57380330d0f8824a8b1894cf063"></a>

## Direct properties — primary.rr_set_group.rr_set.cds_record.values.sha384_digest / 4b822d376d30 / 3

<a id="canonical-a53332073d83865be6d782544db7e1939d6aadd48adfd44920b8ce286146e596"></a>

<a id="canonical-30d45a5ac340bbe5e4865a81157908c5c96d687bb1fb9953c0e30be9e93d7f29"></a>

## digest property — primary.rr_set_group.rr_set.cds_record.values.sha384_digest / 4b822d376d30 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-0f0d7b495f84a7385d61968a5edbcad9b5edc85094259451f1034a5dfdd02d07"></a>

## Next pages — primary.rr_set_group.rr_set.cds_record.values.sha384_digest / 4b822d376d30 / 5

- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-8d4c1e83de3d0939f708128ed6bd6f6a1d1d6262ce26581661a12381192ebfdb)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-8f71d42d744e17ea048e1e9f170cf5acd7f409e73b6bca15712bb308dbde58f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02067f05ece052e03c333c750ba2f9f2361d15cb3fde059f5121e200d6f261b2"></a>

## primary.rr_set_group.rr_set.cert_record — primary.rr_set_group.rr_set.cert_record / b13eea83b851 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.cert_record

<a id="canonical-7feeb9be87c907fd69d8d2fd398f225bffd3a4c367d015f325d8cd71cb4c6ecf"></a>

Type: `"single"`. Computed.

Configuration parameter for cert record.

Upstream description:

DNS CERT Record.

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

<a id="canonical-d3f66e9a3ccf8b1433c339f2c1f75f73c9a48e5be5e10d100b4f015b4f3ea865"></a>

## Direct properties — primary.rr_set_group.rr_set.cert_record / b13eea83b851 / 3

<a id="canonical-062a2d0ee952165251f63149889fe14a7c4b40bd000d0f599062916f33ae4bb4"></a>

<a id="canonical-74050ffeb3400b02c4b59603c1b69e6ed4ad1bad534a6d44617fa3efd4c2c408"></a>

## name property — primary.rr_set_group.rr_set.cert_record / b13eea83b851 / 4

Type: `"string"`. Computed.

CERT Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-8e21e69baa9c8db97e721756e7010768cf43d4dfec1c1e38723c87628aa187fc): complete subsection reference.

<a id="canonical-f0de71b2b42a9302c3c4d95649ea4acc3d59910f5db1d3aa407bc73f9a514134"></a>

## Next pages — primary.rr_set_group.rr_set.cert_record / b13eea83b851 / 5

- [primary.rr_set_group.rr_set.cert_record.values](data-sources--dns_zone--reference--group-002.md#canonical-8e21e69baa9c8db97e721756e7010768cf43d4dfec1c1e38723c87628aa187fc)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-8e21e69baa9c8db97e721756e7010768cf43d4dfec1c1e38723c87628aa187fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77fa647036f8f73147863d28ae2fb5004ac3d1bfc07bfe8413c382292d336537"></a>

## primary.rr_set_group.rr_set.cert_record.values — primary.rr_set_group.rr_set.cert_record.values / fc630534b62b / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.cert_record](data-sources--dns_zone--reference--group-002.md#canonical-8f71d42d744e17ea048e1e9f170cf5acd7f409e73b6bca15712bb308dbde58f4)
- primary.rr_set_group.rr_set.cert_record.values

<a id="canonical-9c6f945acfbe2f50551c6371776d9397cbad8f5980fb6c33c02a4eb1904a265a"></a>

Type: `"list"`. Computed.

CERT Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-95b2ca4f9f5706e716a8695fcc0da6b015b8eba4f2f891d5afddc80fc188ed86"></a>

## Direct properties — primary.rr_set_group.rr_set.cert_record.values / fc630534b62b / 3

<a id="canonical-ff3b6466a275be6a487788a0fba01b48fff064df8d181c788326529efa033359"></a>

<a id="canonical-52e4a4f00ab7682054be36a3ecb8ed014278df98bf227aaabfce06f82f84df93"></a>

## algorithm property — primary.rr_set_group.rr_set.cert_record.values / fc630534b62b / 4

Type: `"string"`. Computed.

\[Enum: RESERVEDALGORITHM|RSAMD5|DH|DSASHA1|ECC|RSASHA1ALGORITHM|INDIRECT|PRIVATEDNS|PRIVATEOID\]
CERT algorithm value must be compatible with the specified algorithm. - RESERVEDALGORITHM:
RESERVEDALGORITHM - RSAMD5: RSAMD5 - DH: DH - DSASHA1: DSASHA1 - ECC: ECC - RSASHA1ALGORITHM:
RSA-SHA1 - INDIRECT: INDIRECT - PRIVATEDNS: PRIVATEDNS - PRIVATEOID: PRIVATEOID. Possible values are
\`RESERVEDALGORITHM\`, \`RSAMD5\`, \`DH\`, \`DSASHA1\`, \`ECC\`, \`RSASHA1ALGORITHM\`, \`INDIRECT\`,
\`PRIVATEDNS\`, \`PRIVATEOID\`. Defaults to \`RESERVEDALGORITHM\`.

Upstream description:

CERT algorithm value must be compatible with the specified algorithm.

&#8203;- RESERVEDALGORITHM: RESERVEDALGORITHM

&#8203;- RSAMD5: RSAMD5

&#8203;- DH: DH

&#8203;- DSASHA1: DSASHA1

&#8203;- ECC: ECC

&#8203;- RSASHA1ALGORITHM: RSA-SHA1

&#8203;- INDIRECT: INDIRECT

&#8203;- PRIVATEDNS: PRIVATEDNS

&#8203;- PRIVATEOID: PRIVATEOID.

Receipt-pinned upstream constraints:

```json
{
  "default": "RESERVEDALGORITHM",
  "enum": [
    "RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d0dd8ec9d4cda3dd5dda55225f37dab6efc0d64017f8fb6943ba40b4815521c9"></a>

<a id="canonical-34bb65536a84a9b65d5e54bc167fd96e4211e870399e927b148b131549a17566"></a>

## cert_key_tag property — primary.rr_set_group.rr_set.cert_record.values / fc630534b62b / 5

Type: `"number"`. Computed.

Key Tag. Tag for categorization and filtering

Upstream description:

Tag for categorization and filtering

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-417e838d67c19b4a418b0ec96e15a84faf59c776f08e2d0499d2966e8f04b183"></a>

<a id="canonical-884be8b7c9c621c9d896eb3628b8f2a8efad5913d2650f8874412ed4828b7585"></a>

## cert_type property — primary.rr_set_group.rr_set.cert_record.values / fc630534b62b / 6

Type: `"string"`. Computed.

\[Enum: INVALIDCERTTYPE|PKIX|SPKI|PGP|IPKIX|ISPKI|IPGP|ACPKIX|IACPKIX|URI\_|OID\] CERT type value
must be compatible with the specified types. - INVALIDCERTTYPE: INVALIDCERTTYPE - PKIX: PKIX - SPKI:
SPKI - PGP: PGP - IPKIX: IPKIX - ISPKI: ISPKI - IPGP: IPGP - ACPKIX: ACPKIX - IACPKIX: IACPKIX -
URI\_: URI - OID: OID. Possible values are \`INVALIDCERTTYPE\`, \`PKIX\`, \`SPKI\`, \`PGP\`,
\`IPKIX\`, \`ISPKI\`, \`IPGP\`, \`ACPKIX\`, \`IACPKIX\`, \`URI\_\`, \`OID\`. Defaults to
\`INVALIDCERTTYPE\`.

Upstream description:

CERT type value must be compatible with the specified types.

&#8203;- INVALIDCERTTYPE: INVALIDCERTTYPE

&#8203;- PKIX: PKIX

&#8203;- SPKI: SPKI

&#8203;- PGP: PGP

&#8203;- IPKIX: IPKIX

&#8203;- ISPKI: ISPKI

&#8203;- IPGP: IPGP

&#8203;- ACPKIX: ACPKIX

&#8203;- IACPKIX: IACPKIX

&#8203;- URI\_: URI

&#8203;- OID: OID.

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALIDCERTTYPE",
  "enum": [
    "INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4acd0d999fe0fd86e8fa7f162231c09fd675a2d9913940e7aa2eae6ec323425f"></a>

<a id="canonical-ddfe84e11e5cd706c85a376ba0533f93be4e620b33fef61b7469e09d236b09ed"></a>

## certificate property — primary.rr_set_group.rr_set.cert_record.values / fc630534b62b / 7

Type: `"string"`. Computed.

Certificate. Certificate in base 64 format.

Upstream description:

Certificate in base 64 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-7bd45891a283dd349aabb2007656efd3b085f67734f258ccf40956ba1d26ad80"></a>

## Next pages — primary.rr_set_group.rr_set.cert_record.values / fc630534b62b / 8

- [primary.rr_set_group.rr_set.cert_record](data-sources--dns_zone--reference--group-002.md#canonical-8f71d42d744e17ea048e1e9f170cf5acd7f409e73b6bca15712bb308dbde58f4)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-c050f8d5b63b59ba93eb329d5b4a402536e9b23baa7a619540485a17db5d950c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0134f2f563c9a6b20c0eae6b1e392eb687343c4876477f9546f7ec5c54d16ac"></a>

## primary.rr_set_group.rr_set.cname_record — primary.rr_set_group.rr_set.cname_record / 9194a25e2d27 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.cname_record

<a id="canonical-4c1a7ccf02cf6149d5735bb50814f4093e891ef1df0b8b1daafa0fd212c960c2"></a>

Type: `"single"`. Computed.

DNSCNAMEResourceRecord.

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

<a id="canonical-2e38685bb79923017209df5df70828d618b2e3d4d6926e138bec59cd1f055115"></a>

## Direct properties — primary.rr_set_group.rr_set.cname_record / 9194a25e2d27 / 3

<a id="canonical-83432a0fdad6f981172b4f83e093664e00dc0d1f0181e39c718d5a4813ef88f9"></a>

<a id="canonical-4a23fc4774e7d8565e64823efb5873afdb9fed3b2f5883c7391084827ef9dbd0"></a>

## name property — primary.rr_set_group.rr_set.cname_record / 9194a25e2d27 / 4

Type: `"string"`. Computed.

CName Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-fc737e60e617bff0cb26bd8de49b5d2eae9c676bf131b6a58e5995f2c95bb9ac"></a>

<a id="canonical-1b8f912aec5146c30f919551ead92fcfd317c461415d651efdc590e9ecb68da5"></a>

## value property — primary.rr_set_group.rr_set.cname_record / 9194a25e2d27 / 5

Type: `"string"`. Computed.

Domain. Configuration parameter for value

Upstream description:

Configuration parameter for value

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-aaa80a933d5679043961844a72c1b7df86af43eee5d1fdb43c9f8115af12d931"></a>

## Next pages — primary.rr_set_group.rr_set.cname_record / 9194a25e2d27 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-0308432bcdcbcd22fae0117c0311b7e2cb63181ad8619f0fa3d5ef8a8dbe05bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c46bbd3097d719583986bcc97de18724128e05f9f76d786747dafa2d3546a8d9"></a>

## primary.rr_set_group.rr_set.ds_record — primary.rr_set_group.rr_set.ds_record / b5f9911be0e8 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- primary.rr_set_group.rr_set.ds_record

<a id="canonical-fb6edbaed969e5cf8438eb882bad11d7780a30e30965a9ce145c0e34dde02839"></a>

Type: `"single"`. Computed.

DNS DS Record. DNS DS Record.

Upstream description:

DNS DS Record.

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

<a id="canonical-43fc89154d82ce1afb738c1a17b2c1784c0b88cf679f6386742840c569ccbd3a"></a>

## Direct properties — primary.rr_set_group.rr_set.ds_record / b5f9911be0e8 / 3

<a id="canonical-9810127fbcdbcb69c7b72d8046f2ffd196f04499ecc46a35a5dc1c5ee109af80"></a>

<a id="canonical-270b7dba1d70fc6993fe3bde98e6aa045d9a5bf775bb51676875b8784713e659"></a>

## name property — primary.rr_set_group.rr_set.ds_record / b5f9911be0e8 / 4

Type: `"string"`. Computed.

DS Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-39f06e65dbea4e404da6b97eccbd64303ce2aceb60f9d161e5e931fcce4c2c9c): complete subsection reference.

<a id="canonical-dfa6ee78f616618f5ab529bad6df9419979e1c77c893b622b337ae4769ed5886"></a>

## Next pages — primary.rr_set_group.rr_set.ds_record / b5f9911be0e8 / 5

- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-39f06e65dbea4e404da6b97eccbd64303ce2aceb60f9d161e5e931fcce4c2c9c)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-39f06e65dbea4e404da6b97eccbd64303ce2aceb60f9d161e5e931fcce4c2c9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c268612a9aa8fa699fa7cc21af203eb872ed3de73c43078762d31489cf5d6f8d"></a>

## primary.rr_set_group.rr_set.ds_record.values — primary.rr_set_group.rr_set.ds_record.values / 263990a0526f / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-78018855a314933db8c9e2b1edd52e65b80fdeff100fdff7b9f615ac179d2205)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-e2c087a24d742a83be947011866125a9d148d22585a1d5114ce07f607fdf9efd)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-5ba8f5e6fad1eb0a038e7cfaced675e3c3b3b98c890130b6b343e5d91acc71e2)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-3971871b39262a8955a50d77f48cffe290b0c387b692d15ca2c72fd07c550e32)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0308432bcdcbcd22fae0117c0311b7e2cb63181ad8619f0fa3d5ef8a8dbe05bd)
- primary.rr_set_group.rr_set.ds_record.values

<a id="canonical-82926de713e09cf0d3b41ebf9519e0b8ed216aff7841301b2c22e5fb4c7381cf"></a>

Type: `"list"`. Computed.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-8f38b4f1257ffd53a014c7bd3052356604f0b7393e0285e4c07c1aa2fa145b40"></a>

## Direct properties — primary.rr_set_group.rr_set.ds_record.values / 263990a0526f / 3

<a id="canonical-805b5db828215dd92054223e66c7d06777b9fc6e7e5fb1985a81f5979a85ae27"></a>

<a id="canonical-577e792e055c791fb5edc690d53f099041da2d271ed89d9ecd5db00138a6d469"></a>

## ds_key_algorithm property — primary.rr_set_group.rr_set.ds_record.values / 263990a0526f / 4

Type: `"string"`. Computed.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-acaec0306ebb65b08ebe55d00a49855c58a97850bc0fe1bd7ba2c95531fcdf49"></a>

<a id="canonical-f4f41c5c10a420c7a57505b14aaacadc62887391073b064e264e7ea74b0df50d"></a>

## key_tag property — primary.rr_set_group.rr_set.ds_record.values / 263990a0526f / 5

Type: `"number"`. Computed.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-e3a768678d36930bfffe27717d246f418532bf322bf2926a7de6127f12a9587d): complete subsection reference.

- [sha256_digest](data-sources--dns_zone--reference--group-003.md#canonical-b58fd918f65b9c174a2110f19d1f177a293cd5825a765eb914707e3999f657cb): complete subsection reference.

- [sha384_digest](data-sources--dns_zone--reference--group-003.md#canonical-c41c505aa680f3668929f4b32ca3e4936276d775404a4916b2132889c0aa9aee): complete subsection reference.

<a id="canonical-bbbaac07930643a15c932a6b83c1523c34bc4882793e3e3d701c7ce443f3b678"></a>

## Next pages — primary.rr_set_group.rr_set.ds_record.values / 263990a0526f / 6

- [primary.rr_set_group.rr_set.ds_record.values.sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-e3a768678d36930bfffe27717d246f418532bf322bf2926a7de6127f12a9587d)
- [primary.rr_set_group.rr_set.ds_record.values.sha256_digest](data-sources--dns_zone--reference--group-003.md#canonical-b58fd918f65b9c174a2110f19d1f177a293cd5825a765eb914707e3999f657cb)
- [primary.rr_set_group.rr_set.ds_record.values.sha384_digest](data-sources--dns_zone--reference--group-003.md#canonical-c41c505aa680f3668929f4b32ca3e4936276d775404a4916b2132889c0aa9aee)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-002.md#canonical-0308432bcdcbcd22fae0117c0311b7e2cb63181ad8619f0fa3d5ef8a8dbe05bd)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-02bc847e36c00371c54bc379ac1dbb58c1c0a02c7c78750a8c5190f02a02af5d)

<a id="canonical-e3a768678d36930bfffe27717d246f418532bf322bf2926a7de6127f12a9587d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
