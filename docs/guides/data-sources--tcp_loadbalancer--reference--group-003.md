---
page_title: "xcsh_tcp_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer reference."
---

# xcsh_tcp_loadbalancer reference

<a id="canonical-f1ec47c8692cea142cc2f571e3fa12f1d4f73b017d877e33337cd7d5ef342290"></a>

## trusted_ca_url property — tls_tcp_auto_cert.use_mtls / d46328e2298e / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-d88ceb2153c36da30b06a1b5ef5ba9628578f1c1289a7ac3688a2306d399cde3): complete subsection reference.

- [xfcc_options](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-77c16ac5e11704f9c0a80510a09537bd22f780080ba636ffd09dfc1e2856dc7f): complete subsection reference.

<a id="canonical-569dfba9f1f453f24dc999d0f1ea8033081c0e7dd48f20f0d5aa4820e665c192"></a>

## Next pages — tls_tcp_auto_cert.use_mtls / d46328e2298e / 6

- [tls_tcp_auto_cert.use_mtls.crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-04fdde51ed3d9f0a1621fe2e940c725b2d1d206fb9f0b600b3cfdb00bee6a7e9)
- [tls_tcp_auto_cert.use_mtls.no_crl](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-dfdac3f21e98272beaa533813199c01ea0a74cef913e719267c32418183737fa)
- [tls_tcp_auto_cert.use_mtls.trusted_ca](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-3cec0f3a70183c7863037189a311aa69680ac922fabf36958b8b30fdeb2ab282)
- [tls_tcp_auto_cert.use_mtls.xfcc_disabled](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-d88ceb2153c36da30b06a1b5ef5ba9628578f1c1289a7ac3688a2306d399cde3)
- [tls_tcp_auto_cert.use_mtls.xfcc_options](data-sources--tcp_loadbalancer--reference--group-003.md#canonical-77c16ac5e11704f9c0a80510a09537bd22f780080ba636ffd09dfc1e2856dc7f)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-04fdde51ed3d9f0a1621fe2e940c725b2d1d206fb9f0b600b3cfdb00bee6a7e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2baf56af8c829644428c103018e7c88fcd292d31484cbeb45a584c580defba3"></a>

## tls_tcp_auto_cert.use_mtls.crl — tls_tcp_auto_cert.use_mtls.crl / fc182839cd6b / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a)
- tls_tcp_auto_cert.use_mtls.crl

<a id="canonical-8b65c7535cb28f815315fb93d2761bf2f62187a800e842ce99d6fe9dadbff094"></a>

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

<a id="canonical-15bdc3eaab9e78b5bf0b342096e0e014e6370ab1e975c4867e647e8d0d586254"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls.crl / fc182839cd6b / 3

<a id="canonical-59467e9c69d3c22791e75134e0ebe759f57a6acc0d692cfaaaf44f0d58fbc0e9"></a>

<a id="canonical-32eb8069098cf28a8ec239970afa881e9d126dd56d5988952e8932b0786bba11"></a>

## name property — tls_tcp_auto_cert.use_mtls.crl / fc182839cd6b / 4

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

<a id="canonical-8f26a75a6ccec616e289f65abc49a2116db94ce40c034d7a514da505a1005e5a"></a>

<a id="canonical-3266c813ee61171c73faf7965c87f041ef4cb33b849ebc1cd4c4272912adbfa5"></a>

## namespace property — tls_tcp_auto_cert.use_mtls.crl / fc182839cd6b / 5

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

<a id="canonical-48e93e69e8c570fe3f39ee42de591849f646d67a54086e7bcc51dee3ef77d39b"></a>

<a id="canonical-ae995f5ac71190af205b288e1f8f9421de5b6156e333ef0a77947d0d2b88fa23"></a>

## tenant property — tls_tcp_auto_cert.use_mtls.crl / fc182839cd6b / 6

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

<a id="canonical-df0190a73d1f8dfd57a85758ef8c56b51f92867c656541264cd2d2e8ee9904de"></a>

## Next pages — tls_tcp_auto_cert.use_mtls.crl / fc182839cd6b / 7

- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-dfdac3f21e98272beaa533813199c01ea0a74cef913e719267c32418183737fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25f8c6ed4ad745423bf39774a715279a12f2dcd9d2f4554fc1d614142df01984"></a>

## tls_tcp_auto_cert.use_mtls.no_crl — tls_tcp_auto_cert.use_mtls.no_crl / 7a0f1d2f1e84 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a)
- tls_tcp_auto_cert.use_mtls.no_crl

<a id="canonical-b4d5e68e2f911d002aa39b00d77ff1a112e6a34c00408f9d3b75138dbd0d6dd3"></a>

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

<a id="canonical-afc4a4d69af3d24c6c4a96d947d404f04fc14c659717f61acabb879e5f843096"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls.no_crl / 7a0f1d2f1e84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2cc7599ba223f9ad70474d784598445f02a8dda4e0c0e0c4fecec7b7dde620d2"></a>

## Next pages — tls_tcp_auto_cert.use_mtls.no_crl / 7a0f1d2f1e84 / 4

- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-3cec0f3a70183c7863037189a311aa69680ac922fabf36958b8b30fdeb2ab282"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d364761fb9447cacc69527be7dba33fbd5d70fffc9add9ed612bfea505e252f8"></a>

## tls_tcp_auto_cert.use_mtls.trusted_ca — tls_tcp_auto_cert.use_mtls.trusted_ca / a0e449adf344 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a)
- tls_tcp_auto_cert.use_mtls.trusted_ca

<a id="canonical-26793d7dacab2547fdc4d25c2cf863b21c8f0acae8b3da17e5ed242282dbd0e5"></a>

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

<a id="canonical-2efb2213cc09019f110c081534d155f20ef285178e52d8fd2b7019a9c9240081"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls.trusted_ca / a0e449adf344 / 3

<a id="canonical-b29ba026fb23d665ac23ae5d9f026be9df4457aad725b1c9be184c391064f7cf"></a>

<a id="canonical-c80bd78e275737d2828fa94f6a5f403f092d723e9270c1182a846f2c4c3f10ad"></a>

## name property — tls_tcp_auto_cert.use_mtls.trusted_ca / a0e449adf344 / 4

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

<a id="canonical-8a0e431c21d596764cd23b4d7ed21969e7741d666d0a750d47f35a44133c4398"></a>

<a id="canonical-8e0a29c694a141c523f1385acbb63dc19343f248b2b77f9645edbd5ad8651c85"></a>

## namespace property — tls_tcp_auto_cert.use_mtls.trusted_ca / a0e449adf344 / 5

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

<a id="canonical-77f13ca43f59f9ca68a631a4f5d49f1cefc5129bdf9650554497bc35c14bb1db"></a>

<a id="canonical-f1e359e52555c8385785a8e584c5c66ee08ace67113a5f54c0e5289303a327e9"></a>

## tenant property — tls_tcp_auto_cert.use_mtls.trusted_ca / a0e449adf344 / 6

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

<a id="canonical-10ecbfb1a21605e0f3db72970b602d64079e7fcbe526a744b3ba5499ea45c676"></a>

## Next pages — tls_tcp_auto_cert.use_mtls.trusted_ca / a0e449adf344 / 7

- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-d88ceb2153c36da30b06a1b5ef5ba9628578f1c1289a7ac3688a2306d399cde3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5b2902ac98f291754d85a5d65dd6255b4b80c276ac15099334e65de0a16feeb"></a>

## tls_tcp_auto_cert.use_mtls.xfcc_disabled — tls_tcp_auto_cert.use_mtls.xfcc_disabled / 4be0238756d7 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a)
- tls_tcp_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-eeb8c4f8dc665473647ade8ac7e5cb09303898d57e1af1098f6ad24661d93ab1"></a>

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

<a id="canonical-d4bf28a007014072520aa6dab512c178fd9757a0e5c5ca1d9a690fff9d31b938"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls.xfcc_disabled / 4be0238756d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1a92134d52cbdbdf1e77124e6b4425f86b4037f6029713c9e18a5a4603667168"></a>

## Next pages — tls_tcp_auto_cert.use_mtls.xfcc_disabled / 4be0238756d7 / 4

- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)

<a id="canonical-77c16ac5e11704f9c0a80510a09537bd22f780080ba636ffd09dfc1e2856dc7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5aab4123b2cede0bd671ca3314836f4b36546967f0d75b4a1cadb8b37dde8a64"></a>

## tls_tcp_auto_cert.use_mtls.xfcc_options — tls_tcp_auto_cert.use_mtls.xfcc_options / 505d9eff965c / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
- [Property reference](data-sources--tcp_loadbalancer--reference--group-001.md#canonical-13ef50def1ca3ae923b44ea13751f5a714a48d50a09cddc18621352d8522588b)
- [tls_tcp_auto_cert](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-17eb82266e303a3d9bed39b3a1288b032620aaf6227dde0f4948ac347bb7b929)
- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a)
- tls_tcp_auto_cert.use_mtls.xfcc_options

<a id="canonical-a7d148f76edcf35d7e6673cd275a6125d11953863696c8b0a3be6802e6b109e8"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

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

<a id="canonical-a911704f79ac397d67c9dc3efbc786104dba56ba01c3d469eebcf75286ff3315"></a>

## Direct properties — tls_tcp_auto_cert.use_mtls.xfcc_options / 505d9eff965c / 3

<a id="canonical-4fbd7a9ebbb27e4690d7e94033bb1aa0b3d65a8b59feba065ad4be8e6eea48e1"></a>

<a id="canonical-9d6f0ed53ce77cd75d018765309c91aed90bf3091415197a34dc8006f1503187"></a>

## xfcc_header_elements property — tls_tcp_auto_cert.use_mtls.xfcc_options / 505d9eff965c / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-646312d4561e2bb6af63e3ae64dbcf345cd279ded97e06edd17e5d6dfcdbf396"></a>

## Next pages — tls_tcp_auto_cert.use_mtls.xfcc_options / 505d9eff965c / 5

- [tls_tcp_auto_cert.use_mtls](data-sources--tcp_loadbalancer--reference--group-002.md#canonical-4135ba6765ed7ad6d4b0b68635fe24b2d2e47b6faf6626aa736e1905618b7f5a)
- [xcsh_tcp_loadbalancer](../data-sources/tcp_loadbalancer.md#canonical-ff619d4c6a92fa41f6695468ec3a5d1538b009b416a57e2136569462b99a7d6f)
