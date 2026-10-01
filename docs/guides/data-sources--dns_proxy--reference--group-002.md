---
page_title: "xcsh_dns_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy reference."
---

# xcsh_dns_proxy reference

<a id="canonical-4827f794125f0e88eb46b774f331157e9d25e7b555d7bea6b93495be6b82a2a7"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 2f956c520f24 / 5

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

<a id="canonical-aadb4c600dd3cb9edef908b78b9d33208abc3e063cb907dc87a565a68cf5069e"></a>

<a id="canonical-edbbea65089dd2d2dbecc8e584bd0578db8b2c3a2a76a1b72217622851751e98"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 2f956c520f24 / 6

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

<a id="canonical-c5dfbdce44fb15b07a8c25cc7d386ba3a02c19ac286d0bd7cff4fadd7983050d"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_publ / 2f956c520f24 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-001.md#canonical-1e6f3c3032c0be1872b90c1fa8a945972be9da7c70a6be7a39a1d0b7f9b6f08d)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-ff29e9aad584e4a7470b72fd27b4d2e4aac34f8b47bc801929171b469fb73ad3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-954de3862dbd0b7ebb5378922becfa7d79146bdfd62f5ede042534c1660dc07a"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_on_public — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public / 4e7988411d42 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public

<a id="canonical-12e07e62837889ed4d39ce59f84682c944c562a642feb7701f1f22147ca29d22"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-426bf8eca1ee9699898de33aa8fffe8eeae6c0e019db57ce5e3483a6cacb4362"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public / 4e7988411d42 / 3

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-afeb43c310dd7a337db0f5ed0593deb5f00f2ad9fb5696cdfb5f9f49eb415399): complete subsection reference.

<a id="canonical-127e396d1e81f875e7eaff9b4430deed545f37bcfd6025d38242e086132c2ce3"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public / 4e7988411d42 / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-afeb43c310dd7a337db0f5ed0593deb5f00f2ad9fb5696cdfb5f9f49eb415399)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-afeb43c310dd7a337db0f5ed0593deb5f00f2ad9fb5696cdfb5f9f49eb415399"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bedab751654eb06bd8ee189a3bcc8225e9d671670f2810f263841e63d41c22bd"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 9b9f40206739 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-ff29e9aad584e4a7470b72fd27b4d2e4aac34f8b47bc801929171b469fb73ad3)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-3d3f07eac8f5b47d5303bf72ea279b5a1b531898745fcd55b03266a3ed2fbef6"></a>

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

<a id="canonical-d06d2d5efbf0cf32ca4329b728741f564195b91ff5c2953df8b01a40bb3f1a15"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 9b9f40206739 / 3

<a id="canonical-18b1e76a95303cad7cc1df145c3fce504474918cbaa87da419f458c356a14e01"></a>

<a id="canonical-6bf9055db287e0c084888bd93557749864c4ea66eb06b6d5dcaebaeca6f7ddb4"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 9b9f40206739 / 4

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

<a id="canonical-885171a9603ed02608e96fd0a96e9d509cf02a36c8fe390ebf912a8e7066883e"></a>

<a id="canonical-8837f36a38e0f8edea167ef7d51fc312c16ceb76c0424f3b5632c37e5e91acf3"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 9b9f40206739 / 5

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

<a id="canonical-8f4174e48a44cb0a028a6740a31dd0bc2f46cfb19a1fc8aabf84c0b07eda6eb3"></a>

<a id="canonical-c4fb42928cf7ae83baac501bb90a2e9346eb003ebd6453d864671dca905020d5"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 9b9f40206739 / 6

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

<a id="canonical-75ed395d92ed858bd7bc44b42196fe01281f6749a17b9a187fd2a180f008b8f4"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ / 9b9f40206739 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-ff29e9aad584e4a7470b72fd27b4d2e4aac34f8b47bc801929171b469fb73ad3)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-c5e01fcffb12a3eab6d30a8dd212340e3a4f649e74a0a0e5a1afe8bdc6e3fe1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-55c1052c9fcde79bd1f3d9bca3172059c3af04d11170c23f003739bf8550ee92"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public / d3794903dc1b / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-b518214d6b63aace97742963c6479668989c17b1674937eb9ae08e812fff9048"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-9e62bf8a672c0c67f77e612ef01c8322325ec6ffd77aa6807150cb678469f973"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public / d3794903dc1b / 3

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-9c5e41178dc3d0dfbe7387d2410ddb66b307663c8fd0740f5a885db5df797eec): complete subsection reference.

<a id="canonical-102481159d4bd7d90e83d142d657291b0ff0ddf6093271be9541a5b8ebd5e9dd"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public / d3794903dc1b / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-9c5e41178dc3d0dfbe7387d2410ddb66b307663c8fd0740f5a885db5df797eec)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-9c5e41178dc3d0dfbe7387d2410ddb66b307663c8fd0740f5a885db5df797eec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0da2e06f5bb3795ce01ca05cb480538230717cc4f09848951b087e1a96b6de72"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / dc8728fd82a9 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-c5e01fcffb12a3eab6d30a8dd212340e3a4f649e74a0a0e5a1afe8bdc6e3fe1f)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-f37e6841bb8722e6cdbba344a31ca49a639e132b0ada680157c4de08cb0d605a"></a>

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

<a id="canonical-d2311696d048ac0c85f1b191ac0d7f8a94083fab2340a7b7cc091f283e4771c7"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / dc8728fd82a9 / 3

<a id="canonical-94f7136c4ace82b6ef88cd7429afa063435a3fb0bf38c99e3ca6d362d66924a0"></a>

<a id="canonical-3ef881504777e0e15c6e57e4b1626284b3e2f956bce18cda42c8ca7de78f07fa"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / dc8728fd82a9 / 4

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

<a id="canonical-6d4ce4d1a41d0d759cca03df1680ee47ad9c3933edf9a092fc956e09a91340fa"></a>

<a id="canonical-ccc97567ad118baa3d7239dfa013edbaa58400bf29cb6bf40dca54ed423674af"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / dc8728fd82a9 / 5

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

<a id="canonical-2ff842a1b3fac92e3a29a977dd3e757f3798ea214d3e3d7d92d7dfb228fe4379"></a>

<a id="canonical-bcb5938c24f64bacffbe28336a5a90fa03ec9feef362321424ce4904c823526b"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / dc8728fd82a9 / 6

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

<a id="canonical-349c0d59764f7a7883d4fac113a8814052d6adf98039abc6c04305b21d6cde6f"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.publ / dc8728fd82a9 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-c5e01fcffb12a3eab6d30a8dd212340e3a4f649e74a0a0e5a1afe8bdc6e3fe1f)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-621fe6793fe6ca6f30379d7d229bc4ce815d52ecb218219841c5cd35762a06e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64e695c8ae98889d22664c0fa603f99c8983c1a265dcd4df3427b89b94ba552f"></a>

## proxy_advertisement.advertise_custom.advertise_where.site — proxy_advertisement.advertise_custom.advertise_where.site / 63f516f5e37c / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- proxy_advertisement.advertise_custom.advertise_where.site

<a id="canonical-c2b13ce75cbd2e5d8e68cc86bdeb81b34a9c98848ed9cd74a3c677f68e1f2bfb"></a>

Type: `"single"`. Computed.

Defines a reference to a CE site along with network type and an optional IP address where a load
balancer could be advertised.

Upstream description:

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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

<a id="canonical-039c4b26fca9c4c148655a1262ff3a15aad9439d7238b801b0bb9bef70a5099a"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.site / 63f516f5e37c / 3

<a id="canonical-5ea60f4cb67b4dca02af884a2af7281f4b37e2651bb97049c3a264386d5472f1"></a>

<a id="canonical-4bfbbe39a22c17af3182b9d4051ac04db75a3111047accb74b6c9e7f1d17c1db"></a>

## ip property — proxy_advertisement.advertise_custom.advertise_where.site / 63f516f5e37c / 4

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0e0615fca8a23d720f1090b482b72a691ce0dfb0385f001240c34806e898aa37"></a>

<a id="canonical-34f6416b627531b78f10b51ba185eb88dcb4abd2862079bf135af79bc7d50cee"></a>

## network property — proxy_advertisement.advertise_custom.advertise_where.site / 63f516f5e37c / 5

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](data-sources--dns_proxy--reference--group-002.md#canonical-1ee7129a7b51ff9ea3ededa81caa66251f5eb251bf513a91e5ec08a52d64dcef): complete subsection reference.

<a id="canonical-0be7e1b99c08d65ce1aa3a2c0afdefd3ef721271f6b97bc4b4a80ac0e65cf8e3"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.site / 63f516f5e37c / 6

- [proxy_advertisement.advertise_custom.advertise_where.site.site](data-sources--dns_proxy--reference--group-002.md#canonical-1ee7129a7b51ff9ea3ededa81caa66251f5eb251bf513a91e5ec08a52d64dcef)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-1ee7129a7b51ff9ea3ededa81caa66251f5eb251bf513a91e5ec08a52d64dcef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22ba8e2701ef447e4c9a7d99473e3f7f9b8e5d3225f0f31966c5e2092b75e0d3"></a>

## proxy_advertisement.advertise_custom.advertise_where.site.site — proxy_advertisement.advertise_custom.advertise_where.site.site / 2083cecbf194 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--dns_proxy--reference--group-002.md#canonical-621fe6793fe6ca6f30379d7d229bc4ce815d52ecb218219841c5cd35762a06e6)
- proxy_advertisement.advertise_custom.advertise_where.site.site

<a id="canonical-fcca4ceeda1816921bd59423be986322b6ff3ac9c9aa0e605dbe3e1d98392cf1"></a>

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

<a id="canonical-c0c824de80b8a805faeefecc06bcfa5dcb230c69ff31bfc0c0bac8e10060eb72"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.site.site / 2083cecbf194 / 3

<a id="canonical-615fd44cb8ee54d42d4f2c291fcd6080b225fccdabe21027668208f55c687b79"></a>

<a id="canonical-0def724f99deea68b76d2774dbd92da0150828984692285c056e34aedc996826"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.site.site / 2083cecbf194 / 4

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

<a id="canonical-31eea8ea53af5ce0f630e183ebcd73140e2837a49880f10d40258ebb1d6465d3"></a>

<a id="canonical-b74462dac5986d4ffa5a5e8651b09e2dd1a48164247fd4d3fcf2aa7417b6590f"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.site.site / 2083cecbf194 / 5

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

<a id="canonical-021275f7217f86434b881f7885c3e97dc2f4b927a3353a52ab8585d22b7b7f83"></a>

<a id="canonical-ae0a05a55b6676e8a58b65fd303010ebbec8f5528d89a79bd0856be4caf1a983"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.site.site / 2083cecbf194 / 6

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

<a id="canonical-bdce7b370ee901cfa6a8ac735026c7f2ed187c033718539a1a0711a2d0cc2b49"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.site.site / 2083cecbf194 / 7

- [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--dns_proxy--reference--group-002.md#canonical-621fe6793fe6ca6f30379d7d229bc4ce815d52ecb218219841c5cd35762a06e6)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-7a12a16bffa8a3ad33139db767a1f14f3467d4bd47171ea90fc95fb865d3fd07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5273c5b19dc235fce5de15862d89b99edd8a9b14bc73c2b4c41664cf6bf4dace"></a>

## proxy_advertisement.advertise_custom.advertise_where.use_default_port — proxy_advertisement.advertise_custom.advertise_where.use_default_port / baeb9a4e1a84 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- proxy_advertisement.advertise_custom.advertise_where.use_default_port

<a id="canonical-0cad932476e3e3a59ea1593fd8197877611659762cee3fbd90756b65a7c09789"></a>

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

<a id="canonical-2396fe82f03ff91fac1357304ca1d5165a58870d6b5823a01dc57ed4d156acf9"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.use_default_port / baeb9a4e1a84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bd2ea2bddb9894d54c7084ffb8bdd6b86350d96a6ac55629af98bcf898e372fa"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.use_default_port / baeb9a4e1a84 / 4

- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-64e5fd38895b34a6b87c36ef654132398bed54dfe40d6c9b03f8b491cfc979dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7a51efc8d6900bf1794579a725783503bf68e16c16ccc23f50ee2c42fe3e654"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network — proxy_advertisement.advertise_custom.advertise_where.virtual_network / a279bca8e1f8 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network

<a id="canonical-bdea9c85c542a40e96ae779e1cbf4a6c7fef3fd1f88bbe26ee26c4793ff60a32"></a>

Type: `"single"`. Computed.

Parameters to advertise on a given virtual network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

<a id="canonical-90fb53d1ba62076df8b8b59901e40071ee838e47e7cb632ac6e0515cfe9b174b"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network / a279bca8e1f8 / 3

- [default_v6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-541717e2e26b8299fd5040a4615c4dc371a2a081303d7b4536abc33ea6ae92af): complete subsection reference.

- [default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-cf4d2fa11ed7f521e114fd13ff8b75ab54c1c7a1d3374d6af89b52b59b6927e5): complete subsection reference.

<a id="canonical-63046b4ba0617e110ac53cff9ce83d5e4a06f968b4fcb1bca9f5b074d90fb3d9"></a>

<a id="canonical-408b7a4fb18931f02d61c7e022c374f1715355eb7c9cd9b8466ea3179edc3ad9"></a>

## specific_v6_vip property — proxy_advertisement.advertise_custom.advertise_where.virtual_network / a279bca8e1f8 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-048cc0609c288d02cac3e14522b9d6f9a3e65e3906092776b833bb7a0f4dca96"></a>

<a id="canonical-b7b3892e055dbbe059a15e13dedbd32646f6d74d9935ca5ca68afeaa8e2ce8f6"></a>

## specific_vip property — proxy_advertisement.advertise_custom.advertise_where.virtual_network / a279bca8e1f8 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-c62612a387980f165d6af61665a744bcd6e83e4dbf8efbf1b370f54e002465b4): complete subsection reference.

<a id="canonical-b2b6a6acf1e21332f27765f0ac6a7115f59ddddd8791c14ff0d52ba910450092"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network / a279bca8e1f8 / 6

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-541717e2e26b8299fd5040a4615c4dc371a2a081303d7b4536abc33ea6ae92af)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-cf4d2fa11ed7f521e114fd13ff8b75ab54c1c7a1d3374d6af89b52b59b6927e5)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-c62612a387980f165d6af61665a744bcd6e83e4dbf8efbf1b370f54e002465b4)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-541717e2e26b8299fd5040a4615c4dc371a2a081303d7b4536abc33ea6ae92af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ddaf384927726f37037d43a71ed408f099618b263c3790d05d874c724bbc246"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_ / 7eea2024a15b / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-64e5fd38895b34a6b87c36ef654132398bed54dfe40d6c9b03f8b491cfc979dc)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-98c0fafe260d1c3091846a59174d0717658e30efa312af2d7e3a07317e463777"></a>

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

<a id="canonical-61af9a8a5aa9e649a0d6cf8b266aab64beff5eba1f95e05134c53c4a82b52a52"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_ / 7eea2024a15b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b8531d9b725e50ab616ba546e87c557840c207dc3d559f7d53bfc419353621b6"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_ / 7eea2024a15b / 4

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-64e5fd38895b34a6b87c36ef654132398bed54dfe40d6c9b03f8b491cfc979dc)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-cf4d2fa11ed7f521e114fd13ff8b75ab54c1c7a1d3374d6af89b52b59b6927e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-661f7d9a19f763f050b69b84dcadd92e6e73fa41c235b661ebc82f8696745cc8"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip / 03b88dab2663 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-64e5fd38895b34a6b87c36ef654132398bed54dfe40d6c9b03f8b491cfc979dc)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-fdc8083742a7642e1898f73a4d4f4ce478a9a766b96f26d1fc484b030e7377d5"></a>

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

<a id="canonical-212e394a239bdcb87d3242c774362de126ee607086f0f480e605b813a04fae77"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip / 03b88dab2663 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-58ffb332bf7f376d91e918f2be60c0f120d361d2bd413761e82fff16ca9e7983"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip / 03b88dab2663 / 4

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-64e5fd38895b34a6b87c36ef654132398bed54dfe40d6c9b03f8b491cfc979dc)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-c62612a387980f165d6af61665a744bcd6e83e4dbf8efbf1b370f54e002465b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d4053bf734d06db4376f1db9bf86d67db70dfd1173dad872755251a74ac2cb5"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / a763fbdd8c34 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-64e5fd38895b34a6b87c36ef654132398bed54dfe40d6c9b03f8b491cfc979dc)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-2312675bbcc2ecd253dca45e251257199f83f52cdc8ca2d349e7aafe60fe492b"></a>

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

<a id="canonical-ceaa7d6dbb20b128005540f4d0a79c1159e14deb043ea51b682b70c6439ec917"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / a763fbdd8c34 / 3

<a id="canonical-b57f7467897d325bbb08f318c50db31cf756e4ed39d84003a7deb9c9d0c04ef2"></a>

<a id="canonical-8aadb539b80003eae714ac9e2e2088a4051eb9e68f4250899e0d1fb8976c6db0"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / a763fbdd8c34 / 4

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

<a id="canonical-c3344bd38705e44eaaa9f5ea2c951ab5169d602efc5944a028dc6a7c4817b6cb"></a>

<a id="canonical-d7c42734ccf37e6359a0b1c07b2eecccfff091e140f66e1a93870ad91cfea446"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / a763fbdd8c34 / 5

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

<a id="canonical-24948a1ee0db876d4e16fc208ae943ce53401cd6a4788a947f029474ef29a994"></a>

<a id="canonical-9c17f885ea1dc09c351795c78df06030f6073d7a0ab5fd132a9f21df4d548fd8"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / a763fbdd8c34 / 6

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

<a id="canonical-c6325d1666488a7948b14e626c01ec4072905dd8ee76da6d7ca9ca34f79da438"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_net / a763fbdd8c34 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-64e5fd38895b34a6b87c36ef654132398bed54dfe40d6c9b03f8b491cfc979dc)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-0e8f2c87d5a27244330d2a8fa89a3ab266680a51ffadd3b03efc669447be6eb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a56bdddaa438101c6f5bc02dfe7a8c919c9ac0c31e8dcfd3a190fc6dfb6d3ef1"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site — proxy_advertisement.advertise_custom.advertise_where.virtual_site / eba80ca441d9 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site

<a id="canonical-8c4b14b4fa228b4ecf09bb574c2adf62eddb13a7b12ef2b3cd7617a7a83a03db"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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

<a id="canonical-5c010849a0ea49a46c864dcc59308f59ddc7358c98acb03f8e02a9bfd4b5c388"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site / eba80ca441d9 / 3

<a id="canonical-3e8f07772f57d0adcc4d3fecf137db8614372f8a5a045d9d2af798dee77910bb"></a>

<a id="canonical-0168ecd8018c5a20fd8b352f73f8fe833e913206d6e7956e8f8ee184a4257902"></a>

## network property — proxy_advertisement.advertise_custom.advertise_where.virtual_site / eba80ca441d9 / 4

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-86281b93bf149c444aec780e9e83756764b77f7d5167b4e2d5a78ffc5ad0b663): complete subsection reference.

<a id="canonical-260b103d5759d6047366e497158a7adbc2bcff2051dd509fa61c9e25f27367ff"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site / eba80ca441d9 / 5

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-86281b93bf149c444aec780e9e83756764b77f7d5167b4e2d5a78ffc5ad0b663)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-86281b93bf149c444aec780e9e83756764b77f7d5167b4e2d5a78ffc5ad0b663"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e403ea73465439a05c7959f8fb858e1f040efeb86d2480416e3ee0bf4a4ec42"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c4121ae4b905 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-0e8f2c87d5a27244330d2a8fa89a3ab266680a51ffadd3b03efc669447be6eb8)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-6cca8f02619d21f6bf23a340b3b6bc4f07474a6d517e31bbec21446b108dd2ee"></a>

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

<a id="canonical-28995569ae3eb3b4965a950eebe4cb7b0cb49aa2630fb6f36dfa985fe243ed82"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c4121ae4b905 / 3

<a id="canonical-8e31420706552bf805bd98ba9aa35ddbdc67f02d361f4b703ada7a61fcd5798b"></a>

<a id="canonical-0601dfdcf971f8a76fb25e4dc7e031778ee38b6dbd0c1db76fa8909996c564bd"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c4121ae4b905 / 4

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

<a id="canonical-2a23245da545d52e85a4011e84df9b38c27f4f666de368a9ba8233d4c2d24d9a"></a>

<a id="canonical-1550c69bf922dd9814d23863b2cec4040aa98a6f78d4e84e6c9e46c500602365"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c4121ae4b905 / 5

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

<a id="canonical-8effe30ce233bb1bf9a424442a0f78873f36b382dbf89dd44e56c8e0e3cfbea1"></a>

<a id="canonical-ea051f202cbc056ed77c56df6cc6e469ed0b8ffeb1fed6db28d5ab64ec73b51a"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c4121ae4b905 / 6

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

<a id="canonical-887450cf7e4b341d6411944b0bc04fc470ae3cbffc6ae40ea847348544c234d5"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site / c4121ae4b905 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-0e8f2c87d5a27244330d2a8fa89a3ab266680a51ffadd3b03efc669447be6eb8)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-0f754bb182bfae26d5bc167a8b0754deee7eb47ca231f12695a151fb2fa37ddf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-801e781fa846dfa3ce5fcf653a6c00a8286dba68244cd4840e74ba147ec3cc8a"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / fcb0abee6619 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-7e408690a758b100baf4411c57051ddd5ee5c44593114fc7854dbbf37628a4af"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

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

<a id="canonical-66e11cd601efffd993d658616eeeaec98b510a6e3eb8cefc8a9e9a943b6908a7"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / fcb0abee6619 / 3

<a id="canonical-03f412f66798f79b2ea4a6aa0cb7d1c6c504f40f23aedc1fc0726d3e410a77a1"></a>

<a id="canonical-808c2d4abedbe62c1400b09f3086eab0db44358986dc7074d6c1e25759f2101e"></a>

## ip property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / fcb0abee6619 / 4

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-c901df74975eda39a4878ac06ee786ab6f838f9f09269efd4fc95d32431c6ffe"></a>

<a id="canonical-2a10203b6ce398a5060a506c7e604e30b053ac941fd23016a8b4bb75aba0c1a3"></a>

## network property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / fcb0abee6619 / 5

Type: `"string"`. Computed.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Upstream description:

This defines network types to be used on virtual-site with specified VIP

All outside networks. All inside networks.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-972f5871d26cb57511bfdc69860faf46fb1f61e975cc5c919e86ca56798cd6ad): complete subsection reference.

<a id="canonical-ae108efa70bf94907d4610d57e77c1a703e706a18e0c74c2b0a2175240a5e930"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip / fcb0abee6619 / 6

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-972f5871d26cb57511bfdc69860faf46fb1f61e975cc5c919e86ca56798cd6ad)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-972f5871d26cb57511bfdc69860faf46fb1f61e975cc5c919e86ca56798cd6ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-915f51ca00ae2e2daa85afc7f2790f4408c55f24d29565fbd875ec9650f4a62f"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 49f6a93119f1 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0f754bb182bfae26d5bc167a8b0754deee7eb47ca231f12695a151fb2fa37ddf)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-395fbf136ae9f1182f7a9155796944a456b79d33a34959f157d007b23cb5bc49"></a>

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

<a id="canonical-549a2646de6b696586909cfdfef84cb04a3caf1b53c18ab3a0ecec8d7ac4c768"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 49f6a93119f1 / 3

<a id="canonical-a6495a51d0c36a5b50239cce60b691cc79cea0f3bf5612843fe010040e47198b"></a>

<a id="canonical-0e8a74cee29e9a1558c226ca237930ebc0381e84bce31b5709a6bc9f730b9f87"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 49f6a93119f1 / 4

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

<a id="canonical-8cd1208ad2f7b14c9ea6a8435df058779e1672e289630a116829fd69113f9b47"></a>

<a id="canonical-51075c9d4129a5a909555d1e5dea520d313fe866c24fb05726b13434407b716c"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 49f6a93119f1 / 5

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

<a id="canonical-512e80cfb7d90efce4e6bdb0a88c3c9e7a737ecfb4465d590fdca850e10e61af"></a>

<a id="canonical-abf8380c4e064848a04844c61744110c651e9f078f46549a16ced16ac4220bda"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 49f6a93119f1 / 6

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

<a id="canonical-e03a27b43452a5f734cfbe1f9162080ee90ec1ef7ad14f1626faf27f696e3c7d"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtu / 49f6a93119f1 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0f754bb182bfae26d5bc167a8b0754deee7eb47ca231f12695a151fb2fa37ddf)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-8f16f574f366343c6077bbc105796cde89af60b22ed3ca94986db24ccc10bcdc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b803140ccec030b07e09d65d7c3ba1a8ff6e967af35a3fc0b1cba9418a339d3"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service — proxy_advertisement.advertise_custom.advertise_where.vk8s_service / 59eecf2f6ca7 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service

<a id="canonical-4f5906fe871268ae7f9483e71839524b2f3c93d1630865a594a262f96cc006ad"></a>

Type: `"single"`. Computed.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-61d2adb66532a886e71cbd7760944078140db1a06d97ed8512fecea4ca3064e2"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.vk8s_service / 59eecf2f6ca7 / 3

- [site](data-sources--dns_proxy--reference--group-002.md#canonical-7ffdb722742755b2ef3fcf94825043af4d5f199bcf0ef2cce9e47cee0077ec60): complete subsection reference.

- [virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-f5f93342c8b512ba5f2c3aa1d1128eed4be36d1cc3c87d4c2951f1b380a0b43d): complete subsection reference.

<a id="canonical-ec275723a39fcbd144dbd195dcae65406505d2ec69b4255f0a7557eae2675a08"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.vk8s_service / 59eecf2f6ca7 / 4

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](data-sources--dns_proxy--reference--group-002.md#canonical-7ffdb722742755b2ef3fcf94825043af4d5f199bcf0ef2cce9e47cee0077ec60)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-f5f93342c8b512ba5f2c3aa1d1128eed4be36d1cc3c87d4c2951f1b380a0b43d)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-7ffdb722742755b2ef3fcf94825043af4d5f199bcf0ef2cce9e47cee0077ec60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4c09d29189b9c98fb651dc921344ed2a4cfa3e5427e0c6c01e031c028136677"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / a8ae39e4b5b2 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-8f16f574f366343c6077bbc105796cde89af60b22ed3ca94986db24ccc10bcdc)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-d2942572aa976f5a0de9ea7dc6b406d36bb63f29520e9fc9ffd2eeea80cb699c"></a>

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

<a id="canonical-acf556edc0afed6c95b11f449586fe2655aba78057ee6f4b2683ecf7b699dd32"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / a8ae39e4b5b2 / 3

<a id="canonical-e0ed1363e5494274fbcde0dbfc9050280fcb44f673ec15e6c37aeec3221138db"></a>

<a id="canonical-cdae4a022e8c0fa4b45f3dca69cde5eb383d0eb759a4ba39d8e60a90df47e495"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / a8ae39e4b5b2 / 4

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

<a id="canonical-d0b2e8bc67bd9ce4ab991a8a59e40563fdcea8fb27f89904451260c645363415"></a>

<a id="canonical-a7fad3b6c8f5c821156ebf103042b11c6a654b3cde14cdb4a9ace0e75bad0144"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / a8ae39e4b5b2 / 5

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

<a id="canonical-e5d431a3f01daf5fd1b6a16bc5e66c7ceede05e8b7afbb23df73106961e20aea"></a>

<a id="canonical-610b80fd1e73362cdee2e441dae7550fa7e7f85926a84ba324d32295f763613e"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / a8ae39e4b5b2 / 6

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

<a id="canonical-f099dd640faf1bd78c424fc5ca2dfd90c5f328617f947c1a57af6ad33d923902"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site / a8ae39e4b5b2 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-8f16f574f366343c6077bbc105796cde89af60b22ed3ca94986db24ccc10bcdc)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-f5f93342c8b512ba5f2c3aa1d1128eed4be36d1cc3c87d4c2951f1b380a0b43d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39a9d85a7cbf73fcc7fe3adc1f5079a451249d1038003d7e8a8c04336cdd5d7b"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 48e58d73fe8a / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-19e136c026a592217e452b512c184b6858079e1eeeeaa9bae5c16c7661d6d291)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-3492c07dcf265db64262170171eb16435a5baaf37d91d53eb712c9d0b21b6de8)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-8f16f574f366343c6077bbc105796cde89af60b22ed3ca94986db24ccc10bcdc)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-48f211ab470ef50498ded4604fb9f1e5b6e3bb73a7e1276b95dcfdf774411897"></a>

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

<a id="canonical-be4c6fc6ada2a8357b5844b2bcaa2d4c2ab851bff271fb4a3f6a846029c47363"></a>

## Direct properties — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 48e58d73fe8a / 3

<a id="canonical-489161520e4ce1a194d14f8b1c8f263a8eb6a508ad7a37b5b3c78cbd9e705c1a"></a>

<a id="canonical-7cb2f04eedb5ad7d1311d55f74059656088d4aeb4cc775be8e32179681fbbd52"></a>

## name property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 48e58d73fe8a / 4

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

<a id="canonical-8a6e5413b1957ee27906b24c21fd05ba77d41b6d8c882036272deb6e8d997c84"></a>

<a id="canonical-6a438f6ef9fe45dec49b58ec3833ddd7c54a2bfb863afd907be6d6985f058132"></a>

## namespace property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 48e58d73fe8a / 5

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

<a id="canonical-7faa77b13198ae0ed186b293dada3c33fb718764ef05a609e98acd59eadcccd4"></a>

<a id="canonical-761b55ce44a0be12f863873907c74cb7ebaa9fedcf143ff0ce3e146202e9d844"></a>

## tenant property — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 48e58d73fe8a / 6

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

<a id="canonical-975b41cb822bff77c919d9545c9ea8eecf9ae50ffe2c76288b401ce96ad5b117"></a>

## Next pages — proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site / 48e58d73fe8a / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-8f16f574f366343c6077bbc105796cde89af60b22ed3ca94986db24ccc10bcdc)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-21fdcf8716231761892b7e61c7bdcd6bcac228499dc6d0cbd9efd8eb82480f99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02f78c8fe2cf7a0653c9d90d842faf2fce3dc05845307a64c717fb47a7a828f3"></a>

## proxy_advertisement.advertise_dualstack_on_public — proxy_advertisement.advertise_dualstack_on_public / c9c05e514adb / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- proxy_advertisement.advertise_dualstack_on_public

<a id="canonical-3963aa67ecea5a48cf21756cfc739b8be6c92876e2ed22f255a1eedabb82310a"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-fc9478c4ab90f5161a9f7e7ae9335d862282e7b3c665343622623b831f1f5b50"></a>

## Direct properties — proxy_advertisement.advertise_dualstack_on_public / c9c05e514adb / 3

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-443ab39daa58856e8c91f6f38b995312f3e9296726f259c791ed1b104788643d): complete subsection reference.

<a id="canonical-2ccecebc58127a6a31013909880847277e1b8a5468434ce4103712f54436e281"></a>

## Next pages — proxy_advertisement.advertise_dualstack_on_public / c9c05e514adb / 4

- [proxy_advertisement.advertise_dualstack_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-443ab39daa58856e8c91f6f38b995312f3e9296726f259c791ed1b104788643d)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-443ab39daa58856e8c91f6f38b995312f3e9296726f259c791ed1b104788643d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ce44375f48894a27aa9bebc7d6f13941a99ca454dba44b6a2c5a8ab9be133db"></a>

## proxy_advertisement.advertise_dualstack_on_public.public_ip — proxy_advertisement.advertise_dualstack_on_public.public_ip / 65f50febb024 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-21fdcf8716231761892b7e61c7bdcd6bcac228499dc6d0cbd9efd8eb82480f99)
- proxy_advertisement.advertise_dualstack_on_public.public_ip

<a id="canonical-de345e6a6a71c0e945ee67bf60d7fba1cf96646eb909eca7e06762831e5f5778"></a>

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

<a id="canonical-b323a00d95b793a6eb75eb83b6f832fb064b8adf972c24247e134636bd91dda4"></a>

## Direct properties — proxy_advertisement.advertise_dualstack_on_public.public_ip / 65f50febb024 / 3

<a id="canonical-3c941c66c5c84ca725f28948ae2b9cb10a84b38f05c61aee927c6487bd0b241b"></a>

<a id="canonical-d978150885a900425a78689535fcaa925bc292b5713998f0a71c0bf062229d80"></a>

## name property — proxy_advertisement.advertise_dualstack_on_public.public_ip / 65f50febb024 / 4

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

<a id="canonical-11dd5f8b7b19f5f3aa7ac388ef028bd8e0b95e4631569edea55bdde3e3cd47bb"></a>

<a id="canonical-3805a52f69850ba733ee3879c1fedb713e19b43097f0b8489643011f5442fffd"></a>

## namespace property — proxy_advertisement.advertise_dualstack_on_public.public_ip / 65f50febb024 / 5

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

<a id="canonical-b085403f66c25aa61678a1a91d0f14c702c2271e3bd5f91a20a5d852a8549eb2"></a>

<a id="canonical-f9a25fcc7e15fe40060d61f1cfc0c01dd68f22ad0c48277be999f5730b0053db"></a>

## tenant property — proxy_advertisement.advertise_dualstack_on_public.public_ip / 65f50febb024 / 6

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

<a id="canonical-ae9668f386a7947a1809111eb50877946cd7579ed44d7ce45562025a55c656df"></a>

## Next pages — proxy_advertisement.advertise_dualstack_on_public.public_ip / 65f50febb024 / 7

- [proxy_advertisement.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-21fdcf8716231761892b7e61c7bdcd6bcac228499dc6d0cbd9efd8eb82480f99)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-f45289809fd7b085fcca1d1df5f4bc10cd0202d1393eb7ba629de991b309111c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef490006585029db96e32b0ad06cc76bb802eaf65221e419e6d5d4a4e07da2ce"></a>

## proxy_advertisement.advertise_on_public — proxy_advertisement.advertise_on_public / 5d817d61abbb / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- proxy_advertisement.advertise_on_public

<a id="canonical-a157cacdd0fc6cf07a4d118bd70ca4ab97f66330910a5cc2b64359aee88e18e3"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-55bf648bad52075ac46d4e3e19bffd43b31acf432db07a4f87fbf14be04ed925"></a>

## Direct properties — proxy_advertisement.advertise_on_public / 5d817d61abbb / 3

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-33049f023de256e9b83dc372ecf55f1506dbe6f5d66bc8c36a58f2af53c4a0a7): complete subsection reference.

<a id="canonical-ec00f1aa070f14b1716410124d00cd627838ae40e78623ed0a2d3549dc87723e"></a>

## Next pages — proxy_advertisement.advertise_on_public / 5d817d61abbb / 4

- [proxy_advertisement.advertise_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-33049f023de256e9b83dc372ecf55f1506dbe6f5d66bc8c36a58f2af53c4a0a7)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-33049f023de256e9b83dc372ecf55f1506dbe6f5d66bc8c36a58f2af53c4a0a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d886dd788e22b2e2ea303369b76b806c8c5742aa7e8faf0d313c37b86ad59ef"></a>

## proxy_advertisement.advertise_on_public.public_ip — proxy_advertisement.advertise_on_public.public_ip / 707136b6a644 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-f45289809fd7b085fcca1d1df5f4bc10cd0202d1393eb7ba629de991b309111c)
- proxy_advertisement.advertise_on_public.public_ip

<a id="canonical-355300169a51d13fb2df11b96f46cdaa9dc37e3bebb21eda1ee200f3568c9936"></a>

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

<a id="canonical-79b3d80bccfb6b42511246e7b0924a710ba52917662cabf1393d03b285d2af06"></a>

## Direct properties — proxy_advertisement.advertise_on_public.public_ip / 707136b6a644 / 3

<a id="canonical-37c4cf38e6283d5c1572010f7e0c6b0d48fc382870117fa23083a85e824c079f"></a>

<a id="canonical-6b7cf97ea9375ca5ee16f44e51808d872101d6856b8d276f82be0255685cc8f2"></a>

## name property — proxy_advertisement.advertise_on_public.public_ip / 707136b6a644 / 4

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

<a id="canonical-9011ac00453e3ff14410b2bc8d339638d835b35e6971bb39d45b1a5db1ea8730"></a>

<a id="canonical-2dc0a9c23a3815b61c4d57d9f794375e8433bd768094fd9d22541e85141d27ba"></a>

## namespace property — proxy_advertisement.advertise_on_public.public_ip / 707136b6a644 / 5

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

<a id="canonical-c4d2c4d07d18be7ab0cc59f69767a299d31f0793cfd20271ec2b41db71d795cc"></a>

<a id="canonical-c8dcbe2f9ed77884d7e9c3fa6b279653d0ab20d56bb97d5638c419b3f56e6ea4"></a>

## tenant property — proxy_advertisement.advertise_on_public.public_ip / 707136b6a644 / 6

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

<a id="canonical-5ee21080f553cad7e883417d3c3e22fc1dfd4c00826e70c4a91a8a6874547960"></a>

## Next pages — proxy_advertisement.advertise_on_public.public_ip / 707136b6a644 / 7

- [proxy_advertisement.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-f45289809fd7b085fcca1d1df5f4bc10cd0202d1393eb7ba629de991b309111c)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-5da82d70aa98c07006c1346e0d2ee465f2cb4f5bb1621b7b588ef1deb4a21f24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-232dd1dd081b4613547613cbd43bb77f0d0a8a653e5c92f2c41edb6bd5a932f9"></a>

## proxy_advertisement.advertise_on_public_default_dualstack_vip — proxy_advertisement.advertise_on_public_default_dualstack_vip / 71f2f353f4d9 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- proxy_advertisement.advertise_on_public_default_dualstack_vip

<a id="canonical-02754fc93b9f773c33a8f617d96c6d013429b8e5defe5bea09a9aa6235fe7237"></a>

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

<a id="canonical-650afee58c7e1797e072823795064fa3fecc32e1d10e96d21c4e2cacdc088f83"></a>

## Direct properties — proxy_advertisement.advertise_on_public_default_dualstack_vip / 71f2f353f4d9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4e83cd236a7bd3c07173586714eecf339113968c5b93e55548112b063e126a76"></a>

## Next pages — proxy_advertisement.advertise_on_public_default_dualstack_vip / 71f2f353f4d9 / 4

- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-04548d7062502d76d540d22defd914edb9470bd895a0730923176e8d8ec67d58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3fa44b3dee7889a35b4d43bc7e41e28501bc467e565a6ff2805953993c2a7e2"></a>

## proxy_advertisement.advertise_on_public_default_ipv6_vip — proxy_advertisement.advertise_on_public_default_ipv6_vip / 87167be8f420 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- proxy_advertisement.advertise_on_public_default_ipv6_vip

<a id="canonical-90b526a24342cafc2151115834fff21f24d4fb250fab4289016b4dd022949b7d"></a>

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

<a id="canonical-b4c4da87c19647cace96225bebe9b8d37ba83d075ec60f5e08b70ea7121d17b9"></a>

## Direct properties — proxy_advertisement.advertise_on_public_default_ipv6_vip / 87167be8f420 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-36cdb838871178c05867f7e8e16cece2673600d3e4f327aa7b1483cff67e2af6"></a>

## Next pages — proxy_advertisement.advertise_on_public_default_ipv6_vip / 87167be8f420 / 4

- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-4ae88fb7618e9af821b1d9b75009b587b713589b606525f666c57e6e1bb0a7b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d40011a0bf8ac12aafedbbd1d38047807cf00726fc1c4856f9e27797092c021"></a>

## proxy_advertisement.advertise_on_public_default_vip — proxy_advertisement.advertise_on_public_default_vip / 803e6f0423cb / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- proxy_advertisement.advertise_on_public_default_vip

<a id="canonical-32600ac7754133bb5c5a9fce8a9d05daef8ef970807c5035742fdbfb9972e173"></a>

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

<a id="canonical-5f30de882ffb804cf974d695aa72f8b65f7eadf3c07aaa2493614e632c21bec7"></a>

## Direct properties — proxy_advertisement.advertise_on_public_default_vip / 803e6f0423cb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ec978a04c37e479244a509bf2a478569ad1a8739ff6da5c86411a65ad0875533"></a>

## Next pages — proxy_advertisement.advertise_on_public_default_vip / 803e6f0423cb / 4

- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-cb10f0b6654a042f07f268fa0f73ecaa482009ef09967a66690c1fa4dad64ce7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73f823dd16d131eac86fc8fd859b56259e437279f6fdfbfe69f7f5ced528fa67"></a>

## proxy_advertisement.advertise_v6_on_public — proxy_advertisement.advertise_v6_on_public / 279d834d358a / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- proxy_advertisement.advertise_v6_on_public

<a id="canonical-51c3c45931d9e05c50604070c4e664ee44024ebcbb67efa90e05f666aff1c924"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

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

<a id="canonical-a6ea0459648ae95f8061f1d479fdc95218f324f36c8ec95a3579f09ec222d909"></a>

## Direct properties — proxy_advertisement.advertise_v6_on_public / 279d834d358a / 3

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-00e833f549ee35bc206b7f83c8fbb66bf523ba4ae3678127bb73b65c2c177f2d): complete subsection reference.

<a id="canonical-38e8f8b753b9b20881a8245c5cdaebdee4b44701ce3a5c89073c349411e37c66"></a>

## Next pages — proxy_advertisement.advertise_v6_on_public / 279d834d358a / 4

- [proxy_advertisement.advertise_v6_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-00e833f549ee35bc206b7f83c8fbb66bf523ba4ae3678127bb73b65c2c177f2d)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-00e833f549ee35bc206b7f83c8fbb66bf523ba4ae3678127bb73b65c2c177f2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29bfd92a8825788fbfb038e3941ac341a9b8a3e97a90b904f274c65cfff5bfd4"></a>

## proxy_advertisement.advertise_v6_on_public.public_ip — proxy_advertisement.advertise_v6_on_public.public_ip / 0210df26546e / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [proxy_advertisement.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-cb10f0b6654a042f07f268fa0f73ecaa482009ef09967a66690c1fa4dad64ce7)
- proxy_advertisement.advertise_v6_on_public.public_ip

<a id="canonical-d0a3b2c79b45bc6af09c709ea5dc47b11084076b3e12e93bcd575e471417a170"></a>

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

<a id="canonical-3596474be4881ddc405c3ed681001372418c5df9710ad4a75918803b5745bbd1"></a>

## Direct properties — proxy_advertisement.advertise_v6_on_public.public_ip / 0210df26546e / 3

<a id="canonical-0d82b4e6b2bbb311beec40b43b98b07979b625015fdd9d8dfd282e72813e462f"></a>

<a id="canonical-2b00435dc74a030b506a6854ee75c10af6abe35906f5ffcddcea60311971215b"></a>

## name property — proxy_advertisement.advertise_v6_on_public.public_ip / 0210df26546e / 4

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

<a id="canonical-fb4e6a1d7e18dfaa81dde820309d7ad2fb853e7f785883c2489a3fc8877f427d"></a>

<a id="canonical-0aee015794f98ae54f7b2760c36270260cc2533282abd5e466ba207573be572a"></a>

## namespace property — proxy_advertisement.advertise_v6_on_public.public_ip / 0210df26546e / 5

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

<a id="canonical-a12a0e5f259d49576c93b2deb2adf2052f599a6ee860e2a800947886b7b94aaa"></a>

<a id="canonical-f8074e18850c12f93acd6c04db206e37521938682a468a93a048d08ad68abfc4"></a>

## tenant property — proxy_advertisement.advertise_v6_on_public.public_ip / 0210df26546e / 6

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

<a id="canonical-fdf1b9386b47df943f91cb64460b73b9a84ac63300558abc8f671919da607295"></a>

## Next pages — proxy_advertisement.advertise_v6_on_public.public_ip / 0210df26546e / 7

- [proxy_advertisement.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-cb10f0b6654a042f07f268fa0f73ecaa482009ef09967a66690c1fa4dad64ce7)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)

<a id="canonical-467fcbabf8498a0bc252658d3c5ba41ee23f6e92f141c3f3e050d16febb486fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7c416bd6a2ce30cc27b919ea25d0695137e7fce05a512fe0e7f3fefb94d6aae"></a>

## proxy_advertisement.do_not_advertise — proxy_advertisement.do_not_advertise / c9703c0b8a90 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-5ff17290d4fe5259d5b90400f5bb92f8e2a7d71f076a1cbc5303680963556408)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- proxy_advertisement.do_not_advertise

<a id="canonical-6049129f23630cf5b93004df2c0a123064ba72308bc0b0b6e9fe725b55754d62"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

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

<a id="canonical-f8c2eac73c9a1dfa481fa202573ec7a5c06550916feb27ed2130a11d3705b861"></a>

## Direct properties — proxy_advertisement.do_not_advertise / c9703c0b8a90 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-90c952de9a4cd0d950d3e7964a28ef8be2be5653c4df19e7a2d42c0bd427ade1"></a>

## Next pages — proxy_advertisement.do_not_advertise / c9703c0b8a90 / 4

- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-6fe78e334e533bf69d5ed47fceee3a41149e6613fa66bd1521656defb4bd9a1f)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-5a1e27cfe185ad04f3b462669b6039f0dc05fd23851daabf1cdff6983bd6f1d6)
