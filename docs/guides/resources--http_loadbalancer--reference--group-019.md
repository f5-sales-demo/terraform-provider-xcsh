---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-cc70dcad06fbfc6f413ed2f7f9293fe95fd84e125d5aed2bf515da20cbc2e1c0"></a>

## Next pages — https.tls_cert_params.use_mtls.no_crl / 30afa9b64747 / 4

- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b02d79626ed0d98f53e005718b6fa3a30bc6c7deea38f903da99d20c3db82d87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19e9adeb670969223b4f6fc6e6739862a591c9cc66820d5762ecc061af53c0cc"></a>

## https.tls_cert_params.use_mtls.trusted_ca — https.tls_cert_params.use_mtls.trusted_ca / 943645433d98 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac)
- https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-20887fb1c5ff31bed4d09b360cd4fb5f1c68d8b96923f31ad3a44e25b2e9191d"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-37b659aaef45d5f09f0080c1e788300e48f6464c3dfc008203274d5b9134e038"></a>

## Direct properties — https.tls_cert_params.use_mtls.trusted_ca / 943645433d98 / 3

<a id="canonical-e0858a10ebf17100761ea56d110c2f93c8147004c154cf93d5b004889c99feca"></a>

<a id="canonical-78a75f3e691d097b401b3259eab31119aad4c1e4527444cc76044283ca649107"></a>

## name property — https.tls_cert_params.use_mtls.trusted_ca / 943645433d98 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-170583a58d590253da5273688f82deba2783818bb936ff4992b77ac1e3a745a4"></a>

<a id="canonical-029f013580b49029685f49bc7038f418a3aa31511fbbfe0f75c4e4e4dc0184ec"></a>

## namespace property — https.tls_cert_params.use_mtls.trusted_ca / 943645433d98 / 5

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
}
```

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

<a id="canonical-59b66e296b71975ff8740d7622f64dd2d3e2329d75e4b3b6997be01fca5f0323"></a>

<a id="canonical-1ae2b0a3f6c850555b992f310c9fbfec9f06df9c8d6f57d73e2c1bed7b1a9581"></a>

## tenant property — https.tls_cert_params.use_mtls.trusted_ca / 943645433d98 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-931e38b8caa2ad03d69521d3869998618d6998ea16cbc7519482bfd6fa149867"></a>

## Next pages — https.tls_cert_params.use_mtls.trusted_ca / 943645433d98 / 7

- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4b63d7f3146ff0f3ebc46e220a87da7eafb5412b0329dc547df8915e04206e7c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2150dbc2c2cc29c6aa1dcd8627dcf0fd2e130d4b413436a25b47ed8d625c11b0"></a>

## https.tls_cert_params.use_mtls.xfcc_disabled — https.tls_cert_params.use_mtls.xfcc_disabled / e0210bf34e4f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac)
- https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-6f9fb3c512b6008655c496aa5a92a0e46aa84b8b74867d1d3d2e3228694d1356"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
xfcc_disabled = {}
```

<a id="canonical-7db4eb43fc5a03a29147e9b8e6492d45c96e0888c35a5f7c3d2897b3b167e0ae"></a>

## Direct properties — https.tls_cert_params.use_mtls.xfcc_disabled / e0210bf34e4f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8164b749415eb9f81bd86db87f7ab502f39f38ce1324f8ce51fef5b79f037b91"></a>

## Next pages — https.tls_cert_params.use_mtls.xfcc_disabled / e0210bf34e4f / 4

- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8c3b0c552caaf6526988ea4b0ae3e98ff612b2909697f15ed31c82ce519e2023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c52fb9798c5309278094c927500cdd30838cfec8f8933d192821db4df39f3c0e"></a>

## https.tls_cert_params.use_mtls.xfcc_options — https.tls_cert_params.use_mtls.xfcc_options / e33e51177f94 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_cert_params](resources--http_loadbalancer--reference--group-018.md#canonical-31ca63fa137c47bc6b3bc3049c6a66381c9558e74ba34ad4e12c7ace5ed5a91f)
- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac)
- https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-36767355aeb78bc1dd7b53ef2ae757d75d06a6846eb4a43e7937f630806f054a"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
```

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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-8056392553ce07b044a2e23d66a66185051cc479702a22ae773afa4ce20289fa"></a>

## Direct properties — https.tls_cert_params.use_mtls.xfcc_options / e33e51177f94 / 3

<a id="canonical-6318b0edfcff2699db3dd30a6847754fd0eaac64f80a42ef568ab769e2b5f9bc"></a>

<a id="canonical-7f3d620a0bf9f9f62cf9728380879ed8594aede2bf459459583e83aef40eef22"></a>

## xfcc_header_elements property — https.tls_cert_params.use_mtls.xfcc_options / e33e51177f94 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-bada57146c67e4b61b75e6a969558445af19ea94dae4cb1b85a7accc8669f2cb"></a>

## Next pages — https.tls_cert_params.use_mtls.xfcc_options / e33e51177f94 / 5

- [https.tls_cert_params.use_mtls](resources--http_loadbalancer--reference--group-018.md#canonical-8c0ad9d7be7da72c3cbec82d199436133f816b111c68e517b73bde71c4aa45ac)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a705ab1aa7d04366ee46a4d59c1e56d0b163b126048c0ebba6992180efbbf08"></a>

## https.tls_parameters — https.tls_parameters / 5b1441cbf6c0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- https.tls_parameters

<a id="canonical-b3593929d6031264c1719ede4b74d07aad57d361a0cf396001e3eb485d9f68b2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
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
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
tls_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-376f9e399c6d25c3a14ff2f4bc7231051376e610083a72c13aa1fd8eaf5fa19a"></a>

## Direct properties — https.tls_parameters / 5b1441cbf6c0 / 3

- [no_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-3439cae70033e63cfd3ed394060d65cfd9949800f99b13cf155992de520c90df): complete subsection reference.

- [tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d): complete subsection reference.

- [tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-7dc6ae9ef4cc527e3c2aab99ee5e4c32d0ec9472bb8e5daeae1b4afa7fabdf00): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac): complete subsection reference.

<a id="canonical-1e4ca03a3d4e39668951f37d321f95449f7f25d4e503173b3fad8a55ebe9e746"></a>

## Next pages — https.tls_parameters / 5b1441cbf6c0 / 4

- [https.tls_parameters.no_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-3439cae70033e63cfd3ed394060d65cfd9949800f99b13cf155992de520c90df)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-7dc6ae9ef4cc527e3c2aab99ee5e4c32d0ec9472bb8e5daeae1b4afa7fabdf00)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-3439cae70033e63cfd3ed394060d65cfd9949800f99b13cf155992de520c90df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb685fda8f8d407eb7ad6e5729444aa7244c0599e698063a06dfaa9f5ac2f6ce"></a>

## https.tls_parameters.no_mtls — https.tls_parameters.no_mtls / de68b0369714 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- https.tls_parameters.no_mtls

<a id="canonical-d6fce74a3c3a330e39c0f06ad6d1cb0bbf6402f8f8a2ab840a851cc2c3ffe74b"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_mtls = {}
```

<a id="canonical-9fa59374319af331a579f90370a04b531a14638018710eca87e6a02fc58ca1e2"></a>

## Direct properties — https.tls_parameters.no_mtls / de68b0369714 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d1b0b1521fa76c44dd9bd112c8d3a88a0078ab39bfb8131a87877d9d11f13858"></a>

## Next pages — https.tls_parameters.no_mtls / de68b0369714 / 4

- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-667d0d11b8b68304dd70ce826b6bc348d6a29823053b2990c4d41384ada0f962"></a>

## https.tls_parameters.tls_certificates — https.tls_parameters.tls_certificates / a673e1f6ae6e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- https.tls_parameters.tls_certificates

<a id="canonical-834f8a36048442966398a28047b00e8ff1cf3b9b059dda2a119beb971ad3c8f4"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-7883b32226e85bf0b4c2b7e802334439dca1d5dc834193b9c4e7c6b48eaec835"></a>

## Direct properties — https.tls_parameters.tls_certificates / a673e1f6ae6e / 3

<a id="canonical-f2be62f19476de6ff6fdfd863551b2cf7d1b36a31547fcf2aafdfaf1f34abbbe"></a>

<a id="canonical-11eebcd7691c1480e24b0d63d9b16ca8e670483f17594b3366d87fce25518223"></a>

## certificate_url property — https.tls_parameters.tls_certificates / a673e1f6ae6e / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [custom_hash_algorithms](resources--http_loadbalancer--reference--group-019.md#canonical-87cff053c5d481295b2f95a3bf34b8dff9217af629bd21a81c48c3a96d16be39): complete subsection reference.

<a id="canonical-5114ce657fd766ced7ef6259448a733a7bf2bde3a6a6707683e08008bd049a9e"></a>

<a id="canonical-8d5d9668b56c3a846e5889b6ad8d3b825b169bb856bcc17908b4a545af13a6e1"></a>

## description_spec property — https.tls_parameters.tls_certificates / a673e1f6ae6e / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--http_loadbalancer--reference--group-019.md#canonical-729b7bbbddf7b8a1230382660ec6622958315c3de37b5a282289f48a4f61b61e): complete subsection reference.

- [private_key](resources--http_loadbalancer--reference--group-019.md#canonical-860b28a2886a0c27c9e7b69c7a1d811512c73eab6f9ef6bfdb84dff1d48851c0): complete subsection reference.

- [use_system_defaults](resources--http_loadbalancer--reference--group-019.md#canonical-6f3cdba686829667bf1e86f64e1438febc72111153ed85a18cdd3c6262c71e93): complete subsection reference.

<a id="canonical-0b9211f24d402ddd9f478ef53049e5a7e491507d2e8a68e87cb97f3ae5ca4f0d"></a>

## Next pages — https.tls_parameters.tls_certificates / a673e1f6ae6e / 6

- [https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--http_loadbalancer--reference--group-019.md#canonical-87cff053c5d481295b2f95a3bf34b8dff9217af629bd21a81c48c3a96d16be39)
- [https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--http_loadbalancer--reference--group-019.md#canonical-729b7bbbddf7b8a1230382660ec6622958315c3de37b5a282289f48a4f61b61e)
- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-860b28a2886a0c27c9e7b69c7a1d811512c73eab6f9ef6bfdb84dff1d48851c0)
- [https.tls_parameters.tls_certificates.use_system_defaults](resources--http_loadbalancer--reference--group-019.md#canonical-6f3cdba686829667bf1e86f64e1438febc72111153ed85a18cdd3c6262c71e93)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-87cff053c5d481295b2f95a3bf34b8dff9217af629bd21a81c48c3a96d16be39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f463e55364dc8fb7429f7ca16e2138057ae950807071dfbe05cf0410c5c01a29"></a>

## https.tls_parameters.tls_certificates.custom_hash_algorithms — https.tls_parameters.tls_certificates.custom_hash_algorithms / 343740ca7f32 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d)
- https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-9e21dce21650c0c637513db59bc4f22ede3e19802b72c1e61a08fd0cf6e4aaed"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
```

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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-a55bc30fd8765c56e497f51e579997efb282dfd66e037e8c4d3090a4fef25294"></a>

## Direct properties — https.tls_parameters.tls_certificates.custom_hash_algorithms / 343740ca7f32 / 3

<a id="canonical-af3049bf80b88581d5a7f6a885456c13dfda3a9154d2e9017fc0ffd8223a0f07"></a>

<a id="canonical-633677d5b99c92a419acde8688b199be1946ae39cd02ea16ddd74fa29a593668"></a>

## hash_algorithms property — https.tls_parameters.tls_certificates.custom_hash_algorithms / 343740ca7f32 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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

<a id="canonical-36f3745042a5dc3e78e3cdb06d0d3afc1cd1a1b5def7a4b651cff00591b498da"></a>

## Next pages — https.tls_parameters.tls_certificates.custom_hash_algorithms / 343740ca7f32 / 5

- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-729b7bbbddf7b8a1230382660ec6622958315c3de37b5a282289f48a4f61b61e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98f885e5f01d90219be90a743bb954d333dd685720dc152c9a8f57717c451b3e"></a>

## https.tls_parameters.tls_certificates.disable_ocsp_stapling — https.tls_parameters.tls_certificates.disable_ocsp_stapling / 70f37ebe28a2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d)
- https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-d9a92e89989d82a07bc27257726edbc118171b0dd161c8917c2aa445dc9d97b3"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

<a id="canonical-75f950f4b8c73d74b5b8606e0b0a41a37242436f3710c2696133831ef7c0846a"></a>

## Direct properties — https.tls_parameters.tls_certificates.disable_ocsp_stapling / 70f37ebe28a2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07e25080e0c231c31df67dc191972fe4b1fee5175001039460c5898c5e76a7fd"></a>

## Next pages — https.tls_parameters.tls_certificates.disable_ocsp_stapling / 70f37ebe28a2 / 4

- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-860b28a2886a0c27c9e7b69c7a1d811512c73eab6f9ef6bfdb84dff1d48851c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d0473aa512d70c37173776589c7c1d68fa7342d7b83a78b34728cab5f945a9d"></a>

## https.tls_parameters.tls_certificates.private_key — https.tls_parameters.tls_certificates.private_key / a865ce356429 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d)
- https.tls_parameters.tls_certificates.private_key

<a id="canonical-7ace765e7a8a7b5505c1e3c0219f1761990448c16ea00d496d50040a0efc531f"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-ebabd1c26a05016ab2d87a4d1f27136cd913f0c8c8be92e6c4e8d48e20ab3b15"></a>

## Direct properties — https.tls_parameters.tls_certificates.private_key / a865ce356429 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-019.md#canonical-9126755fc0ecb23af772662638eeb5bb7d7093870e539e08e0635c600d573cc6): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-019.md#canonical-658892e0fd733768f0a3e6702f110faa86ce3ec6049a6051181e9ea891d7951b): complete subsection reference.

<a id="canonical-06dc776e502019382c8ba28692b30970da04430d24dccebb6c764582443c1034"></a>

## Next pages — https.tls_parameters.tls_certificates.private_key / a865ce356429 / 4

- [https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--http_loadbalancer--reference--group-019.md#canonical-9126755fc0ecb23af772662638eeb5bb7d7093870e539e08e0635c600d573cc6)
- [https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--http_loadbalancer--reference--group-019.md#canonical-658892e0fd733768f0a3e6702f110faa86ce3ec6049a6051181e9ea891d7951b)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9126755fc0ecb23af772662638eeb5bb7d7093870e539e08e0635c600d573cc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c01ea34ebe424152b4de5cf497c0ea58a260860b0b6aee033a5f6d1e7847f16e"></a>

## https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / e1bfbd28b6a0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d)
- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-860b28a2886a0c27c9e7b69c7a1d811512c73eab6f9ef6bfdb84dff1d48851c0)
- https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-aa4a2cf4c75e3f7f3c7fd8880468b72478b5d91e111edb06d5a3a0a2fc5601d0"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-9941f2d0d4e68043ca34d1d55a2c19d4cf0bb809824e47fa2fc70f0870639d37"></a>

## Direct properties — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / e1bfbd28b6a0 / 3

<a id="canonical-7ce43972ef3d39dda0d740e31d3e3701efddb59d704cd91b7df03df5fe49e9dc"></a>

<a id="canonical-a65b3bbf1db6a88d0e144d591d4eb341c2087278cff2dd4193697fc1c1006d5a"></a>

## decryption_provider property — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / e1bfbd28b6a0 / 4

Type: `"string"`. Optional.

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

<a id="canonical-6f3d9a57644cdfdec72f83a0438691f46874b37802dda37de3bcb335566a9970"></a>

<a id="canonical-cbd9e277da9a6511afc9f5155a150d6053407fd848b01f97fa5ed4b16a76fe4c"></a>

## location property — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / e1bfbd28b6a0 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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

<a id="canonical-6f6fd4d38133e463cfb0bc4710b8c790efce190637a56d04d8029026321ed747"></a>

<a id="canonical-22a916f68afcd23bc2240928bb55c8d89184765c782b9abd0ce237f6008510e5"></a>

## store_provider property — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / e1bfbd28b6a0 / 6

Type: `"string"`. Optional.

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

<a id="canonical-b6b661fe6d6f9d94d0591a5588d5e35bd5525f1bfca73fc5ce2a3222435b45ee"></a>

## Next pages — https.tls_parameters.tls_certificates.private_key.blindfold_secret_info / e1bfbd28b6a0 / 7

- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-860b28a2886a0c27c9e7b69c7a1d811512c73eab6f9ef6bfdb84dff1d48851c0)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-658892e0fd733768f0a3e6702f110faa86ce3ec6049a6051181e9ea891d7951b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c0ddd0f8373c209fbc0690e1c49b7424e0a6fa9ad1fd5f61b29f945440a7e2a"></a>

## https.tls_parameters.tls_certificates.private_key.clear_secret_info — https.tls_parameters.tls_certificates.private_key.clear_secret_info / 59214c6d6925 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d)
- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-860b28a2886a0c27c9e7b69c7a1d811512c73eab6f9ef6bfdb84dff1d48851c0)
- https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-db7e6df26cb4cb7043b62fd0060b7cada2cdd76e1fe73b5f646cc30371fc540a"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-c961649d72bfb1aa2f06eb457b6fb2d82d554eb6907cd49839b42a0287761180"></a>

## Direct properties — https.tls_parameters.tls_certificates.private_key.clear_secret_info / 59214c6d6925 / 3

<a id="canonical-aad2e542415c6b286b0f0066a9ce0c8b8ff87fb5c9b1127d801ddedf0025f26f"></a>

<a id="canonical-63a020559feefee6e70bb9840dd41d073894b2e325732e3364b50f499470beb6"></a>

## provider_ref property — https.tls_parameters.tls_certificates.private_key.clear_secret_info / 59214c6d6925 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-98229f0151e6d50b467fc7f113a9b5b0be4e2a0123cffffffc4af4bd1fe243f3"></a>

<a id="canonical-8e75c089eec762611d22d1ef508bf453ff14e46fb82d863752722b48940aada1"></a>

## url property — https.tls_parameters.tls_certificates.private_key.clear_secret_info / 59214c6d6925 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-58958074b019e566791361235790e29b17421535fdb6fc4436c968edf151ecf3"></a>

## Next pages — https.tls_parameters.tls_certificates.private_key.clear_secret_info / 59214c6d6925 / 6

- [https.tls_parameters.tls_certificates.private_key](resources--http_loadbalancer--reference--group-019.md#canonical-860b28a2886a0c27c9e7b69c7a1d811512c73eab6f9ef6bfdb84dff1d48851c0)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6f3cdba686829667bf1e86f64e1438febc72111153ed85a18cdd3c6262c71e93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26af039d9d74430407778a3a00de27a18105e3cbb8a574db94efb13ca07e71af"></a>

## https.tls_parameters.tls_certificates.use_system_defaults — https.tls_parameters.tls_certificates.use_system_defaults / 51c5779e426f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d)
- https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-b3874d16288f1737dcd739a7909e17bfddc485511b377d547cf2816a2aad67a7"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

<a id="canonical-00ec8f064bf0730e998d252001431ad39209341293d3d726d76d2dfac49c79d9"></a>

## Direct properties — https.tls_parameters.tls_certificates.use_system_defaults / 51c5779e426f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-09194948571177624f536636676e66959d843627cdd961487b81bc3a693b8f4a"></a>

## Next pages — https.tls_parameters.tls_certificates.use_system_defaults / 51c5779e426f / 4

- [https.tls_parameters.tls_certificates](resources--http_loadbalancer--reference--group-019.md#canonical-bf8432ea69dd3f1ef6d152cbc130cdf9a8dbb6149f5bf4e371b4b4ef7f4c9a2d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7dc6ae9ef4cc527e3c2aab99ee5e4c32d0ec9472bb8e5daeae1b4afa7fabdf00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64da125ad16daf04b6cc240e611d060453c22db027528b7d125b970937f3b63c"></a>

## https.tls_parameters.tls_config — https.tls_parameters.tls_config / 721469ea7b75 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- https.tls_parameters.tls_config

<a id="canonical-a7c436cd1ce72ec7f5180ba3ac430d7a8f1b104f0f67861a18c7c11c11c93190"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3cca77c9bcd92f956736220d1e53d754d587ff46d33c8fba7eb4343e1525ac67"></a>

## Direct properties — https.tls_parameters.tls_config / 721469ea7b75 / 3

- [custom_security](resources--http_loadbalancer--reference--group-019.md#canonical-adfbe0265db2f2e803c77b790841db16f6646c0e9d252fa53053f52de2b8e0a9): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-019.md#canonical-47d3f1a180e198fab3a99eaec7032ffcc77bf01bccc671aebcc4102145e825e7): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-019.md#canonical-e6174f0c8247aef20c28332a6a562d9e1c89a4dfe557676acf41d435c732f08b): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-019.md#canonical-518631dc271da8f8ecd0d4dcaa716023361a01df60607bd7006a5edd17e7cbe7): complete subsection reference.

<a id="canonical-b0d5d9b3e665f788468e02edaf3f9a5c60108e877e8f1e5eaba44c38ee138f2b"></a>

## Next pages — https.tls_parameters.tls_config / 721469ea7b75 / 4

- [https.tls_parameters.tls_config.custom_security](resources--http_loadbalancer--reference--group-019.md#canonical-adfbe0265db2f2e803c77b790841db16f6646c0e9d252fa53053f52de2b8e0a9)
- [https.tls_parameters.tls_config.default_security](resources--http_loadbalancer--reference--group-019.md#canonical-47d3f1a180e198fab3a99eaec7032ffcc77bf01bccc671aebcc4102145e825e7)
- [https.tls_parameters.tls_config.low_security](resources--http_loadbalancer--reference--group-019.md#canonical-e6174f0c8247aef20c28332a6a562d9e1c89a4dfe557676acf41d435c732f08b)
- [https.tls_parameters.tls_config.medium_security](resources--http_loadbalancer--reference--group-019.md#canonical-518631dc271da8f8ecd0d4dcaa716023361a01df60607bd7006a5edd17e7cbe7)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-adfbe0265db2f2e803c77b790841db16f6646c0e9d252fa53053f52de2b8e0a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ae6ceb99b27722e3e712863712fbaa76dc180ca3be4c29e9929a91a0b6b0730"></a>

## https.tls_parameters.tls_config.custom_security — https.tls_parameters.tls_config.custom_security / 2bbfc66a9d67 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-7dc6ae9ef4cc527e3c2aab99ee5e4c32d0ec9472bb8e5daeae1b4afa7fabdf00)
- https.tls_parameters.tls_config.custom_security

<a id="canonical-24ec858791440407ae1fef61c42004c66ab6ce1c12844c5dc6ccb15bacbbbea1"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
```

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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-4911f9389d18ac4512fc7c5b35b584b3c40cd61e432393f64c44b0914e815757"></a>

## Direct properties — https.tls_parameters.tls_config.custom_security / 2bbfc66a9d67 / 3

<a id="canonical-fcd98422f23b93b79b56e5b3423e9ca0a999111e8c3b4e66b8bc09690eecafb6"></a>

<a id="canonical-7dcf836e25493a887a57a661cef1366a73c26ca5bc7c6190695f5e81aeef4b36"></a>

## cipher_suites property — https.tls_parameters.tls_config.custom_security / 2bbfc66a9d67 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7cd449430f18b0e8ca72107b8d78f92620dff44c9b41f59552a7d3dcbb7fb5b1"></a>

<a id="canonical-40a0b75917357910c79c9543e187a4ae15d8a7701132ddd23974a779a14289fd"></a>

## max_version property — https.tls_parameters.tls_config.custom_security / 2bbfc66a9d67 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-17b98b078b6ac788fc598c49d4ef3fbbe9d47f45ee5f01b9421d15cf3e3589c8"></a>

<a id="canonical-51e800159f0b0d7b09a2147da376f08968fe244f44bbcfdf673014bd30da8038"></a>

## min_version property — https.tls_parameters.tls_config.custom_security / 2bbfc66a9d67 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ce31a2dcdc0ba938761b0228415266d4d918947fc393a32f95f23e6a3c8fa415"></a>

## Next pages — https.tls_parameters.tls_config.custom_security / 2bbfc66a9d67 / 7

- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-7dc6ae9ef4cc527e3c2aab99ee5e4c32d0ec9472bb8e5daeae1b4afa7fabdf00)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-47d3f1a180e198fab3a99eaec7032ffcc77bf01bccc671aebcc4102145e825e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-964474b157fa62a97f7f8543a8e22b7381f297ecb0faf7ed650a374731eb54b4"></a>

## https.tls_parameters.tls_config.default_security — https.tls_parameters.tls_config.default_security / 2ffc09d82054 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-7dc6ae9ef4cc527e3c2aab99ee5e4c32d0ec9472bb8e5daeae1b4afa7fabdf00)
- https.tls_parameters.tls_config.default_security

<a id="canonical-c2038025a399ce535e8c175bd6754306d01620a02e3e627de84fbf574fd7cce4"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_security = {}
```

<a id="canonical-3d8a510b9b7519df039bdbf7fc19cba08c9832795748e7f099867fc1abebc998"></a>

## Direct properties — https.tls_parameters.tls_config.default_security / 2ffc09d82054 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-13225ab20e674a791ab3e8137d0547fadb002157e33ffb53aa98ea5efc54c431"></a>

## Next pages — https.tls_parameters.tls_config.default_security / 2ffc09d82054 / 4

- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-7dc6ae9ef4cc527e3c2aab99ee5e4c32d0ec9472bb8e5daeae1b4afa7fabdf00)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e6174f0c8247aef20c28332a6a562d9e1c89a4dfe557676acf41d435c732f08b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-232642826ad5852e8ca05666cc49347f63dc9d44d69099c671908b8c5b80d1a7"></a>

## https.tls_parameters.tls_config.low_security — https.tls_parameters.tls_config.low_security / 97492eb3b8f2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-7dc6ae9ef4cc527e3c2aab99ee5e4c32d0ec9472bb8e5daeae1b4afa7fabdf00)
- https.tls_parameters.tls_config.low_security

<a id="canonical-51d889e8f4b4bd0be1ef19004b61d84672557b131b9632da9af747a9d03ba9ff"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
low_security = {}
```

<a id="canonical-d984c06106d727f0d0be91567702f0b42ca67999130d767773a99e17c9da92a6"></a>

## Direct properties — https.tls_parameters.tls_config.low_security / 97492eb3b8f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1ed58c723b9e0437b682b357fa999c2af5f43076937fad8bbb36ed540fe4c8e7"></a>

## Next pages — https.tls_parameters.tls_config.low_security / 97492eb3b8f2 / 4

- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-7dc6ae9ef4cc527e3c2aab99ee5e4c32d0ec9472bb8e5daeae1b4afa7fabdf00)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-518631dc271da8f8ecd0d4dcaa716023361a01df60607bd7006a5edd17e7cbe7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b3c132669090f35c3e197facf80e4cf64cd8babc585b28ad9e4616cd13fdfae"></a>

## https.tls_parameters.tls_config.medium_security — https.tls_parameters.tls_config.medium_security / 0ba9665acceb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-7dc6ae9ef4cc527e3c2aab99ee5e4c32d0ec9472bb8e5daeae1b4afa7fabdf00)
- https.tls_parameters.tls_config.medium_security

<a id="canonical-81e2124d36d0a3bdb404513edac3b7562be9cdcbdde612678989546c64c9c062"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
medium_security = {}
```

<a id="canonical-be4701f8943ba5ba8694ad32801e2afffd61428edb291757cbbe161382bbfd76"></a>

## Direct properties — https.tls_parameters.tls_config.medium_security / 0ba9665acceb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-02fc4e3ca2a93b1d05dcdefcbef0e13477189f5a43664555db3c402c6203b25d"></a>

## Next pages — https.tls_parameters.tls_config.medium_security / 0ba9665acceb / 4

- [https.tls_parameters.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-7dc6ae9ef4cc527e3c2aab99ee5e4c32d0ec9472bb8e5daeae1b4afa7fabdf00)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4c88b9397217688badfd0aa809fc2b478fbd85a3cb3dff11ee132b319fb03760"></a>

## https.tls_parameters.use_mtls — https.tls_parameters.use_mtls / a73dc8b0c252 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- https.tls_parameters.use_mtls

<a id="canonical-ca418e2bf1be9bd87c5db67a2fbc729a9578c55f98e7282fc8732d6c04030022"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-d33e0d054f0fdf69c8fca269f76ed0b9d11d78124575b9f44f5172101ec30f2e"></a>

## Direct properties — https.tls_parameters.use_mtls / a73dc8b0c252 / 3

<a id="canonical-0d0b3994c5f246f5eff3d993cbfd80dd1624f55ea1e19743736386a951a348fc"></a>

<a id="canonical-4b1ad8e7dca7e81f6d838c674cec52b6fd9c1ae9a06534b9bb82b15c3602f524"></a>

## client_certificate_optional property — https.tls_parameters.use_mtls / a73dc8b0c252 / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](resources--http_loadbalancer--reference--group-019.md#canonical-2a119724acf784cf43f6748c6d57a61947f528c0c09b7fdbd4213e4800aa0262): complete subsection reference.

- [no_crl](resources--http_loadbalancer--reference--group-019.md#canonical-4c542198ff6a84f5f8b01e066e2356022ecf2289576540140842f02792813495): complete subsection reference.

- [trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-bab693decf160bb85a0a4ae5f59b920e9684b5118be82aa71466fb667ae63f54): complete subsection reference.

<a id="canonical-84e17c567addf16aafbceb92d4e87985f05dff87109f0a7eb941c9cbe659b637"></a>

<a id="canonical-574ead865706208ab0179cbab65e2709c1586b5ff276b67ce561c43a0a383985"></a>

## trusted_ca_url property — https.tls_parameters.use_mtls / a73dc8b0c252 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-ba09c0b87c6094486072ed881fd4ac45e2f6ac14fabae680490d50adae304a08): complete subsection reference.

- [xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-1f7ccec2c6251e842881a3fbea60f51ed9959a011d5e375508713b695dcb397f): complete subsection reference.

<a id="canonical-955ae9d985ab6060f2a891d957a0dc0e79ca638302d8855d0db6b659b134b8ad"></a>

## Next pages — https.tls_parameters.use_mtls / a73dc8b0c252 / 6

- [https.tls_parameters.use_mtls.crl](resources--http_loadbalancer--reference--group-019.md#canonical-2a119724acf784cf43f6748c6d57a61947f528c0c09b7fdbd4213e4800aa0262)
- [https.tls_parameters.use_mtls.no_crl](resources--http_loadbalancer--reference--group-019.md#canonical-4c542198ff6a84f5f8b01e066e2356022ecf2289576540140842f02792813495)
- [https.tls_parameters.use_mtls.trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-bab693decf160bb85a0a4ae5f59b920e9684b5118be82aa71466fb667ae63f54)
- [https.tls_parameters.use_mtls.xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-ba09c0b87c6094486072ed881fd4ac45e2f6ac14fabae680490d50adae304a08)
- [https.tls_parameters.use_mtls.xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-1f7ccec2c6251e842881a3fbea60f51ed9959a011d5e375508713b695dcb397f)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2a119724acf784cf43f6748c6d57a61947f528c0c09b7fdbd4213e4800aa0262"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-326a324b04eb1252f314150793142d9da40d83f29d5f4b0f7344bf1b5608ac54"></a>

## https.tls_parameters.use_mtls.crl — https.tls_parameters.use_mtls.crl / 58690259b5e4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac)
- https.tls_parameters.use_mtls.crl

<a id="canonical-bd308a4d6d6f09e546e9cf7d1168663183b57223e9b2bbf3fe415676d2c23e2b"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-b315869efc5eb14dbefc670ca90952cd29a179efed01f29105ce3175d026b5e6"></a>

## Direct properties — https.tls_parameters.use_mtls.crl / 58690259b5e4 / 3

<a id="canonical-a66a1283a422362f1e2e7bdf774620b8e334de7b5c80726350e902931b6cd58b"></a>

<a id="canonical-6ad85a3d1c9098165c1e7e89586d14a7fe311770f248172fca86e8ac45dcb849"></a>

## name property — https.tls_parameters.use_mtls.crl / 58690259b5e4 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-85e9da96f97a033e76e85533ce2ae56d3b9dc4afdcc9d9dbf7c06f89415abd0a"></a>

<a id="canonical-b3e87d8e04169c03599ecba307cb6ab1013208b799b91a0397e4324be963ddf8"></a>

## namespace property — https.tls_parameters.use_mtls.crl / 58690259b5e4 / 5

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
}
```

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

<a id="canonical-fe112e8b5c58d425df7649198e3e5834ce9b73ad2132be33c32f6f5b02c58d61"></a>

<a id="canonical-1635f7641877fbe8074580ac4df38dfccd639489b6a67516f33b25357fcccd14"></a>

## tenant property — https.tls_parameters.use_mtls.crl / 58690259b5e4 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-34664e38b6004302d57c805196505fda01fbf950df96d3beec33df004309e339"></a>

## Next pages — https.tls_parameters.use_mtls.crl / 58690259b5e4 / 7

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4c542198ff6a84f5f8b01e066e2356022ecf2289576540140842f02792813495"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-401b78e5f52804b7fc5ff80a07a17d3a957bfccaf171fd99e255acb1a4e0e007"></a>

## https.tls_parameters.use_mtls.no_crl — https.tls_parameters.use_mtls.no_crl / 68b7ed0b272b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac)
- https.tls_parameters.use_mtls.no_crl

<a id="canonical-6c0bec5f371e65883b636415341bb88d56793b20798f7b34275ed7726defa976"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_crl = {}
```

<a id="canonical-c03d8666a15f655b042d3f7d37dc24cbb323dc64d519d16e1fb3311d5bda3195"></a>

## Direct properties — https.tls_parameters.use_mtls.no_crl / 68b7ed0b272b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-722e18d835e445b1df23141a19a4cecd9cc3c7e8df49d6b61281a78d965ad358"></a>

## Next pages — https.tls_parameters.use_mtls.no_crl / 68b7ed0b272b / 4

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bab693decf160bb85a0a4ae5f59b920e9684b5118be82aa71466fb667ae63f54"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1999cecb11685563f9b49d3208cd94af490623ca18f7145dbc94285ae8e1c8b7"></a>

## https.tls_parameters.use_mtls.trusted_ca — https.tls_parameters.use_mtls.trusted_ca / 0b74639b8104 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac)
- https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-9afc6882e48806959590289242e3cfe4a05041efb47d1b6f29e23177a9210d3f"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-389254341540e0e33d7c8cb0fb18cacfac59933bae34c31cb40902c771538e5c"></a>

## Direct properties — https.tls_parameters.use_mtls.trusted_ca / 0b74639b8104 / 3

<a id="canonical-6b7c5c5f18bf42371a3033ad5b38d46b1006660653d7f4407420e65b7d75acd2"></a>

<a id="canonical-5aedeebcc75a7284b956f697e462b8eb31d7c64dca07413bb307e4048beacac5"></a>

## name property — https.tls_parameters.use_mtls.trusted_ca / 0b74639b8104 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-65feca61b8fbeeb5bd10f7cada23bace271cc78e24695cf9982e390284605f79"></a>

<a id="canonical-8e0a41898ca86f21c1f8b59a05886995449acdbe0b746d485ea8f5f29a49ebd4"></a>

## namespace property — https.tls_parameters.use_mtls.trusted_ca / 0b74639b8104 / 5

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
}
```

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

<a id="canonical-d88068435699725ce6cdfcd33f35602b82988c2f12ee09af3b9b0c17b4e1ad32"></a>

<a id="canonical-d9b00609a0f6402de917069a5e60aa28049eb18e3fb78d39a902eaef28d09e28"></a>

## tenant property — https.tls_parameters.use_mtls.trusted_ca / 0b74639b8104 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-7051848f3bf33cd6f2e4305c23c4219da39e0592f37d54008dbb865e71b4cd8a"></a>

## Next pages — https.tls_parameters.use_mtls.trusted_ca / 0b74639b8104 / 7

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ba09c0b87c6094486072ed881fd4ac45e2f6ac14fabae680490d50adae304a08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-326c15feec06a2d3d45c0ea76d0077fa5fb6c297560cb385af896b601504efff"></a>

## https.tls_parameters.use_mtls.xfcc_disabled — https.tls_parameters.use_mtls.xfcc_disabled / e94b07601bd0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac)
- https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-81f379d68364f11d6f8bd05e4d301150fe182803346798e659ccf28957b537c3"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
xfcc_disabled = {}
```

<a id="canonical-1834feb0b382a666ddeca6087d5974b62c4d9a7d685b58777a74b7f050f59474"></a>

## Direct properties — https.tls_parameters.use_mtls.xfcc_disabled / e94b07601bd0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ee45b97b49cff5797819a7afc47f7be4cf944b6f7572dce97686128673eb127"></a>

## Next pages — https.tls_parameters.use_mtls.xfcc_disabled / e94b07601bd0 / 4

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1f7ccec2c6251e842881a3fbea60f51ed9959a011d5e375508713b695dcb397f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d16e0f2a4f73eae6d404996a670d9fc56365d6d0c7c2524b88f7f3966f5d478a"></a>

## https.tls_parameters.use_mtls.xfcc_options — https.tls_parameters.use_mtls.xfcc_options / ea3fd2a13d3b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https](resources--http_loadbalancer--reference--group-018.md#canonical-39b9182405c802b182e6b914337b4da12ad4bf0ded160b49869796b301f98012)
- [https.tls_parameters](resources--http_loadbalancer--reference--group-019.md#canonical-6d5c7ce1efa80c26a575c1f2ba94312938bd3c5c92b95992b99bf5d0dc8f22a6)
- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac)
- https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-b08324aab69c0d0cbdb6b4cb87b322f82e047e861910c791aa4288b990b704b7"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
```

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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-6587886c8f9413a9cc7be58d6712b4bed26ad636c0adf126c4314eb48b088217"></a>

## Direct properties — https.tls_parameters.use_mtls.xfcc_options / ea3fd2a13d3b / 3

<a id="canonical-938bb7061d6af328392ad7689a1cd0ac6827c16483557912150d8aabd2638bfd"></a>

<a id="canonical-58ed8080112f3a2f2fc7bce46d02075128b40fff277be761e34c0b18704f5b30"></a>

## xfcc_header_elements property — https.tls_parameters.use_mtls.xfcc_options / ea3fd2a13d3b / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-4f56174d8d69154eaecb047e3e6d3fd429c86fd219b8307ae132e29c29e65e73"></a>

## Next pages — https.tls_parameters.use_mtls.xfcc_options / ea3fd2a13d3b / 5

- [https.tls_parameters.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-79f91f754b5b7324b632351b623ca20c114504948a855bc6e4fcebb6f562adac)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6bdccd22538b0e87c1215498c5545f158f1729648391f6d56322ce7d5cf503b"></a>

## https_auto_cert — https_auto_cert / c443377c5f0d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- https_auto_cert

<a id="canonical-6b6f058e2a7c8a41e1f4634b1f217cf058ab1bd258b9943ea8c4d3640bf10090"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

Terraform syntax:

```terraform
https_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-4142e72d41523ba666470f1a59c7db4f6fd3880ca6a464c4fb8177b54229a36c"></a>

## Direct properties — https_auto_cert / c443377c5f0d / 3

<a id="canonical-ebe7465bb1ac91c595c2572e1a0304b1fffdf86d447a875c73120bc7c4e9f07f"></a>

<a id="canonical-15fec095917c3ee2f836f8b19fdc19e925f96970e869227b802d50a45df627b3"></a>

## add_hsts property — https_auto_cert / c443377c5f0d / 4

Type: `"bool"`. Optional, Computed.

Add HTTP Strict-Transport-Security response header. Defaults to \`false\`. Server applies default
when omitted.

Upstream description:

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-cdd13e2614c5658f8a4129212ad6045f3556988b9ecddff8ef9d871f3be9308a"></a>

<a id="canonical-e3fccfe5837ab23506f23446427c109a2b921c218d3b368d35a2d3af79ed0b33"></a>

## append_server_name property — https_auto_cert / c443377c5f0d / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-b85152d94c15aa4fab6b124e6429b644e116422871fe9fe5e00ab5bf7e2d14d3): complete subsection reference.

<a id="canonical-120a6520792bf455f738753720696074c6ad432816c9457625785cacffde1c5c"></a>

<a id="canonical-5187aef9b0f680a2fadbda62c0171e28d826669a18a6158e3b95a9bc46f799ac"></a>

## connection_idle_timeout property — https_auto_cert / c443377c5f0d / 6

Type: `"number"`. Optional, Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Server
applies default when omitted.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--http_loadbalancer--reference--group-019.md#canonical-97294496451e779d2bdb413f9bcf50d56bf3db135de52f3cd12825cbd27fa298): complete subsection reference.

- [default_loadbalancer](resources--http_loadbalancer--reference--group-019.md#canonical-b409e75dab9c2db15524b177e070fc58b2bd484d1c420da2851dc39fedc3c8e7): complete subsection reference.

- [disable_path_normalize](resources--http_loadbalancer--reference--group-019.md#canonical-14cadd614047bce1353858d719889995a09802a5ea5507cfeaaaac72e433683c): complete subsection reference.

- [enable_path_normalize](resources--http_loadbalancer--reference--group-019.md#canonical-b4a705f0d982bb083cccef6ca6c0c72bc597529638395ce1d0df589176ef2f1d): complete subsection reference.

- [http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a): complete subsection reference.

<a id="canonical-3b11188c6eb0445431f4529390f6943c0e2811a8aa098cfe7ae88f6f567bf523"></a>

<a id="canonical-6e7f3c83874f5dd2f5885438b5168eede5aedee3da25e405aceeebd98fa2c731"></a>

## http_redirect property — https_auto_cert / c443377c5f0d / 7

Type: `"bool"`. Optional, Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS. Defaults to \`false\`. Server applies
default when omitted.

Upstream description:

Redirect HTTP traffic to HTTPS.

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

- [no_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-64611bea0b1362f63ccc686fefc6e145183e245a5ec3c058ec514326e09af27e): complete subsection reference.

- [non_default_loadbalancer](resources--http_loadbalancer--reference--group-019.md#canonical-122a4e407c2b66f22305de1c121ee3e522186c22ce410844f2ffb26044b50743): complete subsection reference.

- [pass_through](resources--http_loadbalancer--reference--group-019.md#canonical-9dcdaecc01f4f9d93f5a952d1f00ae2de5c97efeb1feb3f7ae532478e2c11fb6): complete subsection reference.

<a id="canonical-d26b90532377aab421a4d7cba6f7e25cf83463eafb5baa28ae245d2a6e759d56"></a>

<a id="canonical-17c27bf84cf143b0ed497dae53fccf382e204242cce27395bb78a48e5a982bbf"></a>

## port property — https_auto_cert / c443377c5f0d / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-87640c1ca206348e771a328232b02236e926a3f3a2c7bc4f3854921c5b75a4f9"></a>

<a id="canonical-bb4b8d963fe1ba37f5fe1e83026c86f6e7dcc5e5ea39ea0e719259c6f7625fc3"></a>

## port_ranges property — https_auto_cert / c443377c5f0d / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-34228b9f2ba91a506cb8f53f3fd35e8c0ca7df54a0bd99f91f9b57d1accdb2ab"></a>

<a id="canonical-ae50dc1bbc19d937d557586795d38c9350bd065e83620a8e6a76e7f3345e3326"></a>

## server_name property — https_auto_cert / c443377c5f0d / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-cba85732493542cf0bb17133e5c71383009cdacec91316aa4dbc23819c86fe0a): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d): complete subsection reference.

<a id="canonical-51ffe0f4d46d946e9f01e6c9ca74ac9e570fa2e8eb43761d1f4957f6e10b4eac"></a>

## Next pages — https_auto_cert / c443377c5f0d / 11

- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-b85152d94c15aa4fab6b124e6429b644e116422871fe9fe5e00ab5bf7e2d14d3)
- [https_auto_cert.default_header](resources--http_loadbalancer--reference--group-019.md#canonical-97294496451e779d2bdb413f9bcf50d56bf3db135de52f3cd12825cbd27fa298)
- [https_auto_cert.default_loadbalancer](resources--http_loadbalancer--reference--group-019.md#canonical-b409e75dab9c2db15524b177e070fc58b2bd484d1c420da2851dc39fedc3c8e7)
- [https_auto_cert.disable_path_normalize](resources--http_loadbalancer--reference--group-019.md#canonical-14cadd614047bce1353858d719889995a09802a5ea5507cfeaaaac72e433683c)
- [https_auto_cert.enable_path_normalize](resources--http_loadbalancer--reference--group-019.md#canonical-b4a705f0d982bb083cccef6ca6c0c72bc597529638395ce1d0df589176ef2f1d)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a)
- [https_auto_cert.no_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-64611bea0b1362f63ccc686fefc6e145183e245a5ec3c058ec514326e09af27e)
- [https_auto_cert.non_default_loadbalancer](resources--http_loadbalancer--reference--group-019.md#canonical-122a4e407c2b66f22305de1c121ee3e522186c22ce410844f2ffb26044b50743)
- [https_auto_cert.pass_through](resources--http_loadbalancer--reference--group-019.md#canonical-9dcdaecc01f4f9d93f5a952d1f00ae2de5c97efeb1feb3f7ae532478e2c11fb6)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-cba85732493542cf0bb17133e5c71383009cdacec91316aa4dbc23819c86fe0a)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b85152d94c15aa4fab6b124e6429b644e116422871fe9fe5e00ab5bf7e2d14d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18ea93fc1efab6a79f2974aa2faa0f1734b6780d7205047b603bb363e8e221cd"></a>

## https_auto_cert.coalescing_options — https_auto_cert.coalescing_options / ad60ba7c9319 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- https_auto_cert.coalescing_options

<a id="canonical-c1b66fa1377a819c0c2a6e80d66ef4f10325fd72fb250b384a7efb6a349fba4f"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-44a4a0db73108127af45fd4dcde87a0d837a45a51849ed0bc5a3097e2384ff80"></a>

## Direct properties — https_auto_cert.coalescing_options / ad60ba7c9319 / 3

- [default_coalescing](resources--http_loadbalancer--reference--group-019.md#canonical-f548aa1b7cec256f3602128233c6abea985f5df863be6ab8afd2f4a4144592b0): complete subsection reference.

- [strict_coalescing](resources--http_loadbalancer--reference--group-019.md#canonical-7ac6dc2723f180ba3ffd8a432046543aaa39c30a4cf553bc81efc8a505a2444f): complete subsection reference.

<a id="canonical-79108b3d501458011bc10d7dc7e1135a171f4d5986a686485119d32a09b74ccd"></a>

## Next pages — https_auto_cert.coalescing_options / ad60ba7c9319 / 4

- [https_auto_cert.coalescing_options.default_coalescing](resources--http_loadbalancer--reference--group-019.md#canonical-f548aa1b7cec256f3602128233c6abea985f5df863be6ab8afd2f4a4144592b0)
- [https_auto_cert.coalescing_options.strict_coalescing](resources--http_loadbalancer--reference--group-019.md#canonical-7ac6dc2723f180ba3ffd8a432046543aaa39c30a4cf553bc81efc8a505a2444f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f548aa1b7cec256f3602128233c6abea985f5df863be6ab8afd2f4a4144592b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87297223ddee4d57988e571492a7b7d9d1d052233b674fa749663ddea5e676c5"></a>

## https_auto_cert.coalescing_options.default_coalescing — https_auto_cert.coalescing_options.default_coalescing / 520d08bee76a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-b85152d94c15aa4fab6b124e6429b644e116422871fe9fe5e00ab5bf7e2d14d3)
- https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-5143d742e3ede96577e23ff25f328306bd1782c737eb3c3fed5eac0047a2d0fd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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

Terraform syntax:

```terraform
default_coalescing = {}
```

<a id="canonical-31e2937b0641c2207f0ffe98590d27e7e90fdbf9ca119e89d03eea35a21ebb26"></a>

## Direct properties — https_auto_cert.coalescing_options.default_coalescing / 520d08bee76a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e9878d2921478f0438ee67d3ffbe7f3f9eec4ac4fbfb54f06c1651ecfbc228c6"></a>

## Next pages — https_auto_cert.coalescing_options.default_coalescing / 520d08bee76a / 4

- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-b85152d94c15aa4fab6b124e6429b644e116422871fe9fe5e00ab5bf7e2d14d3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7ac6dc2723f180ba3ffd8a432046543aaa39c30a4cf553bc81efc8a505a2444f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40734bb099ab82b01500a6b5cc44154d05b69219e28a27a86e97bc2fadb3c8f3"></a>

## https_auto_cert.coalescing_options.strict_coalescing — https_auto_cert.coalescing_options.strict_coalescing / bd2e6adecad3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-b85152d94c15aa4fab6b124e6429b644e116422871fe9fe5e00ab5bf7e2d14d3)
- https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-e02fd89eb15bd8ce5648f21a253508f9d170bfaab093752695c88c6562ee4a2a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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

Terraform syntax:

```terraform
strict_coalescing = {}
```

<a id="canonical-8a9e018338e5aeab3553fee902d4b6ff873d6a3ccb093824c89e6e16f6dd3aae"></a>

## Direct properties — https_auto_cert.coalescing_options.strict_coalescing / bd2e6adecad3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-86f280c989310743e26d03ba3b44ee93e2179cbce78e369fca33a94b5e004aef"></a>

## Next pages — https_auto_cert.coalescing_options.strict_coalescing / bd2e6adecad3 / 4

- [https_auto_cert.coalescing_options](resources--http_loadbalancer--reference--group-019.md#canonical-b85152d94c15aa4fab6b124e6429b644e116422871fe9fe5e00ab5bf7e2d14d3)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-97294496451e779d2bdb413f9bcf50d56bf3db135de52f3cd12825cbd27fa298"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38d88c594c5defbf5495e68613b3ea239ae34b7d0d190417b7726f159b989cf9"></a>

## https_auto_cert.default_header — https_auto_cert.default_header / d0e21e7be079 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- https_auto_cert.default_header

<a id="canonical-90d5bb5d1fd54890be2cc66e0bac50738db8f842f273d7b6cc5122b5a0273fef"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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

Terraform syntax:

```terraform
default_header = {}
```

<a id="canonical-761606513e6a6cd5fdb4386145b96a296aa0ee4f61b09aa5fced0614a53b518b"></a>

## Direct properties — https_auto_cert.default_header / d0e21e7be079 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-95bf9cab64ea6831955910a6b98de0ff09e9e56417f0374e2f4d136e3d936b3c"></a>

## Next pages — https_auto_cert.default_header / d0e21e7be079 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b409e75dab9c2db15524b177e070fc58b2bd484d1c420da2851dc39fedc3c8e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-828431adedd0022826d2b8b60584bd6c280ab017171b42beb9504ccf480a4b57"></a>

## https_auto_cert.default_loadbalancer — https_auto_cert.default_loadbalancer / f1d0a4ce53bc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- https_auto_cert.default_loadbalancer

<a id="canonical-c1008befc1e7ad89ba6c6a891f84afd6fbfe58d33ccbd4b16e8cfda132eb75d4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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

Terraform syntax:

```terraform
default_loadbalancer = {}
```

<a id="canonical-a0543c13c5770b5e2cea426bf433090ce008cc193b3040e46ce49ca10fe5c2f1"></a>

## Direct properties — https_auto_cert.default_loadbalancer / f1d0a4ce53bc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b049429e3191d7a174d0af5976b9129d15a8275047982a2cd8a09b4e3d786145"></a>

## Next pages — https_auto_cert.default_loadbalancer / f1d0a4ce53bc / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-14cadd614047bce1353858d719889995a09802a5ea5507cfeaaaac72e433683c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb96b649d50641b6d681b988b30f6a2b7ee69a7b8a6880383c5f2c4d877f1541"></a>

## https_auto_cert.disable_path_normalize — https_auto_cert.disable_path_normalize / 0c71bd324e78 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- https_auto_cert.disable_path_normalize

<a id="canonical-3f9da09a8c441bfe9bffb019546ad9fbccf4d3056bfabdd3ec96638472d79298"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_path_normalize = {}
```

<a id="canonical-feedce2df1479d4721937bed8949462961adb46a1836512d404eb1dfeb608036"></a>

## Direct properties — https_auto_cert.disable_path_normalize / 0c71bd324e78 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-16f02f0ab53e9c64093ccd98e5aacca98e4f49acd2b0ba526df1809d55ea8a36"></a>

## Next pages — https_auto_cert.disable_path_normalize / 0c71bd324e78 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-b4a705f0d982bb083cccef6ca6c0c72bc597529638395ce1d0df589176ef2f1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e85e7e8a153bc23af6deeb0d2e49a9114d8366a2891c09390b0e3af07fcda6fb"></a>

## https_auto_cert.enable_path_normalize — https_auto_cert.enable_path_normalize / d4e749592df0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- https_auto_cert.enable_path_normalize

<a id="canonical-ab1eb844f31fc2c3f0e567622a1f7413788d1a229c23070837ebe2ac91b1d015"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

Terraform syntax:

```terraform
enable_path_normalize = {}
```

<a id="canonical-fca87c48cfe6baec4b2f49d4776e0e27bb2698d855d972d6af0ab2539fd75b9f"></a>

## Direct properties — https_auto_cert.enable_path_normalize / d4e749592df0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dabdc33627c71ecf04824d773d223b331a7ca7a1e446443ce77355eaed8fbff4"></a>

## Next pages — https_auto_cert.enable_path_normalize / d4e749592df0 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6196eda937e6e1bf7944454563b1780f82b9abe5b925ce6844899d24e40f3252"></a>

## https_auto_cert.http_protocol_options — https_auto_cert.http_protocol_options / b3d9963d7ec6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- https_auto_cert.http_protocol_options

<a id="canonical-855de75254ad9f6cf4fe4833a94e959aad43a705a216aef7c81d2f9307ee4fac"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-7d64dd67462e5edd6fcffc3f52963f3e4568be65ecd1173578ae9887fd27b0e8"></a>

## Direct properties — https_auto_cert.http_protocol_options / b3d9963d7ec6 / 3

- [http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-5ca53a905c1283bdf6a5086a05f47a0ff274eca2b5714e6ed6124ea9d74fc908): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--http_loadbalancer--reference--group-019.md#canonical-4e91f0edce7a3188732e404ab3aac1b150bb5185818dd28fb490099a8457b6c6): complete subsection reference.

- [http_protocol_enable_v2_only](resources--http_loadbalancer--reference--group-019.md#canonical-24f58b50749af6949875e169e565435bf0c0d3e290a90b72fd7102ed68129842): complete subsection reference.

<a id="canonical-8e4b7494e7f9b854b4171912adf3479ef7007ece0f41a63fc911b66f38136f03"></a>

## Next pages — https_auto_cert.http_protocol_options / b3d9963d7ec6 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-5ca53a905c1283bdf6a5086a05f47a0ff274eca2b5714e6ed6124ea9d74fc908)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--http_loadbalancer--reference--group-019.md#canonical-4e91f0edce7a3188732e404ab3aac1b150bb5185818dd28fb490099a8457b6c6)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--http_loadbalancer--reference--group-019.md#canonical-24f58b50749af6949875e169e565435bf0c0d3e290a90b72fd7102ed68129842)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5ca53a905c1283bdf6a5086a05f47a0ff274eca2b5714e6ed6124ea9d74fc908"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-721d3131c0f4d8acb378ea8e249869384d6915a6ef383cfb6c779f75e0808dee"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only / 24873836a2ab / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-8ddb5e8abbfa866353d2d66ae1c6973f110ec790864c0f3ea18de99159cca733"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-46afd91a14c72a1b459af296a005813da2c81aebf5083c215610aa5d309945f2"></a>

## Direct properties — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only / 24873836a2ab / 3

- [header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-47e7af06cede0daaeddceebca630fe867d1044d25ca717082a72f0058c5fe42b): complete subsection reference.

<a id="canonical-81460bd59c3ad3703361ec5e125dffbe343462aba39c0763aeb8abb51bc89f01"></a>

## Next pages — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only / 24873836a2ab / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-47e7af06cede0daaeddceebca630fe867d1044d25ca717082a72f0058c5fe42b)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-47e7af06cede0daaeddceebca630fe867d1044d25ca717082a72f0058c5fe42b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8813cb187450022d04b437e50e4ede20819a0dea5ba25b5bf3660a08da9e6ce5"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / 79e9646649a8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-5ca53a905c1283bdf6a5086a05f47a0ff274eca2b5714e6ed6124ea9d74fc908)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-eb19b727ee8dc78fef81803cb141fb385b69fcf599f9b56c746743ab187ad197"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-e8c494117c25d9cf4b3c04b7230de1445130e3bbed6a1d3fa55bb73ed1ea5c96"></a>

## Direct properties — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / 79e9646649a8 / 3

- [default_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-55fd7a547592ec8a49f994b8c9f58c6c57d86bfb23622636588d671b5eaa2183): complete subsection reference.

- [preserve_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-366a62fc2e25db9e487167c7c331b048c5e6ee4291f36eea7fdb0560d1b2f778): complete subsection reference.

- [proper_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-a2c3a7955400cb25e381ee74685174cca7c7bc5eb4a53195e5a432ce77d40887): complete subsection reference.

<a id="canonical-2874019de30fddab86e12f040a75125cb8c5861870fb2aecbc54657d0ba6dfdc"></a>

## Next pages — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / 79e9646649a8 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-55fd7a547592ec8a49f994b8c9f58c6c57d86bfb23622636588d671b5eaa2183)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-366a62fc2e25db9e487167c7c331b048c5e6ee4291f36eea7fdb0560d1b2f778)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-a2c3a7955400cb25e381ee74685174cca7c7bc5eb4a53195e5a432ce77d40887)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-5ca53a905c1283bdf6a5086a05f47a0ff274eca2b5714e6ed6124ea9d74fc908)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-55fd7a547592ec8a49f994b8c9f58c6c57d86bfb23622636588d671b5eaa2183"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2de7d1d403affaa4554616604267c4dc86595aeabaa7e15160dde457c6269522"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / ced718c1f699 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-5ca53a905c1283bdf6a5086a05f47a0ff274eca2b5714e6ed6124ea9d74fc908)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-47e7af06cede0daaeddceebca630fe867d1044d25ca717082a72f0058c5fe42b)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-aabeddd417f35120fb08d4a7dec1d8fba12c0eccbfd395681c6b19fc930c1f7e"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

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
default_header_transformation = {}
```

<a id="canonical-e8921acc4078e95a72ee6f5a90a05a06883a1699edcc71e25a1ab71a5e03d608"></a>

## Direct properties — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / ced718c1f699 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-494c4ce049cac6e17e152f79675fcbf1308a0244fadbaa8396baebbaffb99cf9"></a>

## Next pages — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / ced718c1f699 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-47e7af06cede0daaeddceebca630fe867d1044d25ca717082a72f0058c5fe42b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-366a62fc2e25db9e487167c7c331b048c5e6ee4291f36eea7fdb0560d1b2f778"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00a03ca18bb0d44b6f1f4ddcf2df01167d1726406a9f34143bc41c22c4ee12c4"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / 1691b1dc1941 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-5ca53a905c1283bdf6a5086a05f47a0ff274eca2b5714e6ed6124ea9d74fc908)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-47e7af06cede0daaeddceebca630fe867d1044d25ca717082a72f0058c5fe42b)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-bc14b302d01a0588a5ab943c93e87c33830262ed65d5e65257f521996d59e49b"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

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
preserve_case_header_transformation = {}
```

<a id="canonical-6136e5a3f9f0470b2246957bdb295d7538ffd619fbbd6f7219528f45c66ea86f"></a>

## Direct properties — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / 1691b1dc1941 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-90d85707da7edd6b8e2d514f89784a89b7e82f118b5bf72019bbdb5329276136"></a>

## Next pages — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / 1691b1dc1941 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-47e7af06cede0daaeddceebca630fe867d1044d25ca717082a72f0058c5fe42b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a2c3a7955400cb25e381ee74685174cca7c7bc5eb4a53195e5a432ce77d40887"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b73f868892647b26bf2576762f7fedf47764879134ca89fc0f665611b7012902"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / ebfefca333b6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--http_loadbalancer--reference--group-019.md#canonical-5ca53a905c1283bdf6a5086a05f47a0ff274eca2b5714e6ed6124ea9d74fc908)
- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-47e7af06cede0daaeddceebca630fe867d1044d25ca717082a72f0058c5fe42b)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-7f46217f0efe4b63948b66d84450f339f0f350791c08fa51b8d7f9df3f497701"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

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
proper_case_header_transformation = {}
```

<a id="canonical-1524e7151f996b7b46755bd8c61192f98c5bc168d9d3849fb0f1ecb4a698378f"></a>

## Direct properties — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / ebfefca333b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ec096e0b9136305bfad6108bedd7d3cefad954c9afc87fa06798e6dad9d16a6f"></a>

## Next pages — https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transf / ebfefca333b6 / 4

- [https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--http_loadbalancer--reference--group-019.md#canonical-47e7af06cede0daaeddceebca630fe867d1044d25ca717082a72f0058c5fe42b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4e91f0edce7a3188732e404ab3aac1b150bb5185818dd28fb490099a8457b6c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9edd5c87b67ea239d24745725915f0284ed9c14ac0d248cfec61142b50b8ef4"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 / 496464995eb6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-9bb0435b48cffd0edd586545bf7c5278a36dc0b27618a3ad798cf45d7b7597f2"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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

Terraform syntax:

```terraform
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-e3bcba7547365826b63dd9d22100c5b5cc8e369454f5a80f7d865b7047ebdb0b"></a>

## Direct properties — https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 / 496464995eb6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-22082c538f2ead0dd5f4a17f9082b53507d4aa97ce8a7b221ac8b26eaba6aba7"></a>

## Next pages — https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 / 496464995eb6 / 4

- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-24f58b50749af6949875e169e565435bf0c0d3e290a90b72fd7102ed68129842"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd75fa332d005ae1783ef072427447d91ddeeba28104101c869007027897f410"></a>

## https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — https_auto_cert.http_protocol_options.http_protocol_enable_v2_only / ea627c7faf59 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a)
- https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-da8c2f9e1b2714dad333d092b0f75ed9bc3363a2982249919f7aa34e681c46ca"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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

Terraform syntax:

```terraform
http_protocol_enable_v2_only = {}
```

<a id="canonical-a98c48b901a0f7a9d0784d88191f25be5036fdc36594dc7dc96c0054e82faf18"></a>

## Direct properties — https_auto_cert.http_protocol_options.http_protocol_enable_v2_only / ea627c7faf59 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-83074d7bebaef98bedd89170940254fef84a8e9712277a0d9575552ec7fe570a"></a>

## Next pages — https_auto_cert.http_protocol_options.http_protocol_enable_v2_only / ea627c7faf59 / 4

- [https_auto_cert.http_protocol_options](resources--http_loadbalancer--reference--group-019.md#canonical-a9bd3c4650e1269e6d70e7d97efd3dbebd25b0e0273b91f1d689217d28c2f81a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-64611bea0b1362f63ccc686fefc6e145183e245a5ec3c058ec514326e09af27e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0217deb977412fde8bd9a43a4c6d5ae8ddc1a4a3de05ee0b5174f2a922582c02"></a>

## https_auto_cert.no_mtls — https_auto_cert.no_mtls / 3eb8a518b555 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- https_auto_cert.no_mtls

<a id="canonical-a8f88048b2e42c46b501e4aa5aa6d406f6b45746b3f54a08c2690d4d65e3e4e9"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

Terraform syntax:

```terraform
no_mtls = {}
```

<a id="canonical-7d7630a42504751da77fac049d9418a22eb40f80169f25f70366a4efa23526a5"></a>

## Direct properties — https_auto_cert.no_mtls / 3eb8a518b555 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8755048f66329288b90bb8e973e2b149c84cefc2ef662c3ca537a64b5add43d9"></a>

## Next pages — https_auto_cert.no_mtls / 3eb8a518b555 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-122a4e407c2b66f22305de1c121ee3e522186c22ce410844f2ffb26044b50743"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fea1aaf33e914121b7355b5dfccf1c910193a5d9f875c133b2e208e5c1e65583"></a>

## https_auto_cert.non_default_loadbalancer — https_auto_cert.non_default_loadbalancer / 5108b1eef2f6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- https_auto_cert.non_default_loadbalancer

<a id="canonical-41a52cc5a8250a062c39088571f849088c2b21b3f2b65cd3a5626f53bf7ec574"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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

Terraform syntax:

```terraform
non_default_loadbalancer = {}
```

<a id="canonical-9f999ad80475b8cdd5296641a3d272b4a06ed4c6206e7f6c310e100248ed3746"></a>

## Direct properties — https_auto_cert.non_default_loadbalancer / 5108b1eef2f6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6ddf86ad40474ab777f6c8fe6c6ef12168892fc2f81f843dff7fb13e5390c073"></a>

## Next pages — https_auto_cert.non_default_loadbalancer / 5108b1eef2f6 / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9dcdaecc01f4f9d93f5a952d1f00ae2de5c97efeb1feb3f7ae532478e2c11fb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d1c9081e7f12e9f654b6b905122d8b5314979664b3384d21a1e1bb16e313201"></a>

## https_auto_cert.pass_through — https_auto_cert.pass_through / 26b277fcd38b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- https_auto_cert.pass_through

<a id="canonical-f5a4618ce7b3e269e6f737c83f9a4cd55a264a0d9d8927e5270cf8a52890b911"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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

Terraform syntax:

```terraform
pass_through = {}
```

<a id="canonical-acff450288bf635c00dccbc74e3fcf36810495035de96949cdd1e9cd837beac8"></a>

## Direct properties — https_auto_cert.pass_through / 26b277fcd38b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-705f58e01c226f928fb1fffc1d15b6316f2ceebdcf4f439f94bccb7d4c244c84"></a>

## Next pages — https_auto_cert.pass_through / 26b277fcd38b / 4

- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cba85732493542cf0bb17133e5c71383009cdacec91316aa4dbc23819c86fe0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c7d5f0ac6ff3a978461ceeeab7c784aec0449764d2f7eeb795224b2c1d19ec2a"></a>

## https_auto_cert.tls_config — https_auto_cert.tls_config / 1c5162957935 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- https_auto_cert.tls_config

<a id="canonical-2f620dbe654f0a18bec257209769b6862df8d4335eaa0cfdbcffc6560fd59f49"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-b49617c83eb85f5181b4bfc48d5988c3437094f8d2741e307939c7e836a1440f"></a>

## Direct properties — https_auto_cert.tls_config / 1c5162957935 / 3

- [custom_security](resources--http_loadbalancer--reference--group-019.md#canonical-60208a6a746a6985d17a9f27a2f6aa6689f505af9cc889ddc9fb4a4e977a95e8): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-019.md#canonical-eb883ceb8f584bdc7017f26c0f602187e3918dc9a23bd0aec59b02e1ea6229f3): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-019.md#canonical-d1d2b578f34e1efffea60786027a4b24748120b02a7704f6206551cb96fd1276): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-019.md#canonical-c5b65237e0bdfc1a435ea2fc08e9be7fd7525e8c55834f71316f8a2d88bf98b8): complete subsection reference.

<a id="canonical-4e6422a3c7a46ddc0db7155de97b29be6afe7288ef096a08fc629646b9bf3524"></a>

## Next pages — https_auto_cert.tls_config / 1c5162957935 / 4

- [https_auto_cert.tls_config.custom_security](resources--http_loadbalancer--reference--group-019.md#canonical-60208a6a746a6985d17a9f27a2f6aa6689f505af9cc889ddc9fb4a4e977a95e8)
- [https_auto_cert.tls_config.default_security](resources--http_loadbalancer--reference--group-019.md#canonical-eb883ceb8f584bdc7017f26c0f602187e3918dc9a23bd0aec59b02e1ea6229f3)
- [https_auto_cert.tls_config.low_security](resources--http_loadbalancer--reference--group-019.md#canonical-d1d2b578f34e1efffea60786027a4b24748120b02a7704f6206551cb96fd1276)
- [https_auto_cert.tls_config.medium_security](resources--http_loadbalancer--reference--group-019.md#canonical-c5b65237e0bdfc1a435ea2fc08e9be7fd7525e8c55834f71316f8a2d88bf98b8)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-60208a6a746a6985d17a9f27a2f6aa6689f505af9cc889ddc9fb4a4e977a95e8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4026ef1ec9369ec8597aebe6907e335ae00ed0a265961a83dc68a5f3e6da647"></a>

## https_auto_cert.tls_config.custom_security — https_auto_cert.tls_config.custom_security / 188a52476345 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-cba85732493542cf0bb17133e5c71383009cdacec91316aa4dbc23819c86fe0a)
- https_auto_cert.tls_config.custom_security

<a id="canonical-57dcf9bb494c44ef3abb5456ca69a62b759a319151a06cb37798aad091d3c366"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
```

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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-48d3ce3c4490bb0274d8b86b7d3b1861393c30ab8e10046023c69e9e02d2febe"></a>

## Direct properties — https_auto_cert.tls_config.custom_security / 188a52476345 / 3

<a id="canonical-85ace16eda5b71ccbc08e0ae3e2d54cab04fa615944b954a2e23467652f8ea0c"></a>

<a id="canonical-f39e21a87e1d6a4d832561fbdb7515061b2866e225976a1ea2e78b43809c4227"></a>

## cipher_suites property — https_auto_cert.tls_config.custom_security / 188a52476345 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-00dba4ad4ca5f7681fe61d809e083f96411ef9d0a284d74a75888e096aff6c4c"></a>

<a id="canonical-efec0eb941bdd4a30b7bbbc593eb6fbf70a86414959d4f159ada8ade764db53d"></a>

## max_version property — https_auto_cert.tls_config.custom_security / 188a52476345 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-11d6aab7d06145fd77a61171856cfc41f5f6a14bc6d7be5f6c19aaf9b5a91376"></a>

<a id="canonical-fd7986046310c299af35169ab66330182ad1cf5296b41a4016c876165f1ed8f0"></a>

## min_version property — https_auto_cert.tls_config.custom_security / 188a52476345 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-9e2938b439ac456feecd03255bc334d5b625bb0d72b3e793ebe97fdd38d44b63"></a>

## Next pages — https_auto_cert.tls_config.custom_security / 188a52476345 / 7

- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-cba85732493542cf0bb17133e5c71383009cdacec91316aa4dbc23819c86fe0a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-eb883ceb8f584bdc7017f26c0f602187e3918dc9a23bd0aec59b02e1ea6229f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dee7ebdc7cede493e52955f11f92d3fc4e816401db679a48ee401b2ac3a42905"></a>

## https_auto_cert.tls_config.default_security — https_auto_cert.tls_config.default_security / 29ed7c417039 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-cba85732493542cf0bb17133e5c71383009cdacec91316aa4dbc23819c86fe0a)
- https_auto_cert.tls_config.default_security

<a id="canonical-a2053fc982a14069a28c11bccbe254a43019dd92573b88c025b3694a519ac1d4"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_security = {}
```

<a id="canonical-d34ef5880e20f4e023252043026d174d41b6166ba8017401e688dc1e760a9ccb"></a>

## Direct properties — https_auto_cert.tls_config.default_security / 29ed7c417039 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-700ce5f13867517e51f4a3df4bd88297d77fbf11685cf2075123c99e0e77ecd8"></a>

## Next pages — https_auto_cert.tls_config.default_security / 29ed7c417039 / 4

- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-cba85732493542cf0bb17133e5c71383009cdacec91316aa4dbc23819c86fe0a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d1d2b578f34e1efffea60786027a4b24748120b02a7704f6206551cb96fd1276"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c8915e391ff68328f77286e75f77a3a0676d59f3d052b5a47d0b817d54a5890"></a>

## https_auto_cert.tls_config.low_security — https_auto_cert.tls_config.low_security / dfb0c2e72adc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-cba85732493542cf0bb17133e5c71383009cdacec91316aa4dbc23819c86fe0a)
- https_auto_cert.tls_config.low_security

<a id="canonical-2168130990702ddabe989fcc54a785ca0eba3b73d7d8d884d0ec9208a7fa4853"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
low_security = {}
```

<a id="canonical-05c3282b9089a8f4fd90513d4434096e21565c694a9d7761d666dc7f81e165de"></a>

## Direct properties — https_auto_cert.tls_config.low_security / dfb0c2e72adc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4698bbc81f8dd0da52fd8ceebcc7a10d0adec87b66f5f0d9f232d24e84e34e8f"></a>

## Next pages — https_auto_cert.tls_config.low_security / dfb0c2e72adc / 4

- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-cba85732493542cf0bb17133e5c71383009cdacec91316aa4dbc23819c86fe0a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c5b65237e0bdfc1a435ea2fc08e9be7fd7525e8c55834f71316f8a2d88bf98b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad2c4858c60a7998eb013fc8a2f5114deabcded0731d499fb5bee315ad0c8087"></a>

## https_auto_cert.tls_config.medium_security — https_auto_cert.tls_config.medium_security / 2391d1710345 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-cba85732493542cf0bb17133e5c71383009cdacec91316aa4dbc23819c86fe0a)
- https_auto_cert.tls_config.medium_security

<a id="canonical-21730771078a6a731d7e0160b42b40b5e6fab586bdd636224a537ce397538ec5"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
medium_security = {}
```

<a id="canonical-b586addb94302b3775cefdd4b51837ff28fe9936b47931e313f72cad50828326"></a>

## Direct properties — https_auto_cert.tls_config.medium_security / 2391d1710345 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5bd29cfc54e827644197eb8f3c1f0ef08ae1f6ba149dac70019c2a1c7dc93c18"></a>

## Next pages — https_auto_cert.tls_config.medium_security / 2391d1710345 / 4

- [https_auto_cert.tls_config](resources--http_loadbalancer--reference--group-019.md#canonical-cba85732493542cf0bb17133e5c71383009cdacec91316aa4dbc23819c86fe0a)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba474a129502aaa7d76fc9d90155ac358f6fc70b458e9a1a770aef991f470659"></a>

## https_auto_cert.use_mtls — https_auto_cert.use_mtls / 8d9ff14793d6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- https_auto_cert.use_mtls

<a id="canonical-b5c3667c00cede0fe58f8fc1298eb0d4d35b7a116c07faf7e20e04616bfbfa63"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-f5d33b390f01179782e987dc713c2992705f18def7a84b78b3c3e2013779af00"></a>

## Direct properties — https_auto_cert.use_mtls / 8d9ff14793d6 / 3

<a id="canonical-212e03b232efe61502abdc45daeb0be11fd592bfb34ec34f6dd9e268a8cd3751"></a>

<a id="canonical-259b1bd0ed4111b4071e8bdbebfd18e55331e6ba74fa80b8745c28a3d95a6c49"></a>

## client_certificate_optional property — https_auto_cert.use_mtls / 8d9ff14793d6 / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](resources--http_loadbalancer--reference--group-019.md#canonical-647c21cf53740168b2305f12307a29735aacacfe7e0cfe6f629d20698a15b5b0): complete subsection reference.

- [no_crl](resources--http_loadbalancer--reference--group-019.md#canonical-36820c3387628ad174f1f06469b90a46f0ada2e3aff565d5a74b5d346ea8b20a): complete subsection reference.

- [trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-1afa10ad4d82c994fe007cfe291a0ad06d297c4b2c941eb3607c04d1c1712f33): complete subsection reference.

<a id="canonical-d8f5bf8dffcebd1b7a7d6900cdf643bc18e120120bb64f4c721349640c5f1540"></a>

<a id="canonical-76ebf2aa04fbd8d93afb06e8622a76f6a902bd9f38e193904c7c6b89d294c95c"></a>

## trusted_ca_url property — https_auto_cert.use_mtls / 8d9ff14793d6 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-e3af7e77748ec2390458db79169261c959b803567987405ea28014fb1193cf8e): complete subsection reference.

- [xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-dc5dd1806459e7740f0dec52d5f32382a5f23e90422a515acb1085d250eaa0d4): complete subsection reference.

<a id="canonical-05bacb803d0cac38f914fa26b62fcd71b038a4c26e20001ec3a28609e3213a42"></a>

## Next pages — https_auto_cert.use_mtls / 8d9ff14793d6 / 6

- [https_auto_cert.use_mtls.crl](resources--http_loadbalancer--reference--group-019.md#canonical-647c21cf53740168b2305f12307a29735aacacfe7e0cfe6f629d20698a15b5b0)
- [https_auto_cert.use_mtls.no_crl](resources--http_loadbalancer--reference--group-019.md#canonical-36820c3387628ad174f1f06469b90a46f0ada2e3aff565d5a74b5d346ea8b20a)
- [https_auto_cert.use_mtls.trusted_ca](resources--http_loadbalancer--reference--group-019.md#canonical-1afa10ad4d82c994fe007cfe291a0ad06d297c4b2c941eb3607c04d1c1712f33)
- [https_auto_cert.use_mtls.xfcc_disabled](resources--http_loadbalancer--reference--group-019.md#canonical-e3af7e77748ec2390458db79169261c959b803567987405ea28014fb1193cf8e)
- [https_auto_cert.use_mtls.xfcc_options](resources--http_loadbalancer--reference--group-019.md#canonical-dc5dd1806459e7740f0dec52d5f32382a5f23e90422a515acb1085d250eaa0d4)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-647c21cf53740168b2305f12307a29735aacacfe7e0cfe6f629d20698a15b5b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41373fa99c6a421da6bacae4349079b372d933f081fd0150cc8e782ac369c274"></a>

## https_auto_cert.use_mtls.crl — https_auto_cert.use_mtls.crl / 92edf2882279 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d)
- https_auto_cert.use_mtls.crl

<a id="canonical-9ca1b0fec594f13f83e40cd1073194cc893e1cda36470de2e10ffbb1b8c8cd21"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-2d3a01d90fb2545ec4189b3e8811408297210ba4afca4bb603a9ef046b5c962c"></a>

## Direct properties — https_auto_cert.use_mtls.crl / 92edf2882279 / 3

<a id="canonical-4ac892d58baa9400c14a1765d087de5e3cc9fd2c600a80416727b53acf109caf"></a>

<a id="canonical-2afbf6258e7fc456b6ed43aa691b2a8ee59c12b54d650c176386ccf0bc5a025d"></a>

## name property — https_auto_cert.use_mtls.crl / 92edf2882279 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-91dc1d7c3c4f7399f7ee2df42fb07d8aec7ba7380214f15e8b878098d56b7c36"></a>

<a id="canonical-92ac1b364a1333fd3acbef43c0be4cb4d2ea6423167abf8f4d5c836309f5e9b4"></a>

## namespace property — https_auto_cert.use_mtls.crl / 92edf2882279 / 5

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
}
```

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

<a id="canonical-49fba5e63ddb1d851a7e9c7da7aa1d39aec685425707e4184bdc6ed82c6ef7c5"></a>

<a id="canonical-3f486865d9c138420f3d6c8e827416151a22ffd6da8b5a32a80197bef70ce1be"></a>

## tenant property — https_auto_cert.use_mtls.crl / 92edf2882279 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-48d68cb2f3e7b5a4fe8b1dce488438d4a142acfd9e1f332dc3a123946c2d0827"></a>

## Next pages — https_auto_cert.use_mtls.crl / 92edf2882279 / 7

- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-36820c3387628ad174f1f06469b90a46f0ada2e3aff565d5a74b5d346ea8b20a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76c2cf9de4988c1f5ef21d0cda4bc333f1f4fe5acc74ad947d4c6b3da34ed246"></a>

## https_auto_cert.use_mtls.no_crl — https_auto_cert.use_mtls.no_crl / 451140fb1d77 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d)
- https_auto_cert.use_mtls.no_crl

<a id="canonical-073f51b150591179303e248cd424c441edfce045d6e8bad927f47d496da9442b"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_crl = {}
```

<a id="canonical-c34c323bf70b12f654bd2631ac408beaac89c6ace3306298c654b63e5b7aac2e"></a>

## Direct properties — https_auto_cert.use_mtls.no_crl / 451140fb1d77 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8849df13d60d40de8b2b01b8b7c5efccd53f47ec7ce043bb96844221bc58cafc"></a>

## Next pages — https_auto_cert.use_mtls.no_crl / 451140fb1d77 / 4

- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1afa10ad4d82c994fe007cfe291a0ad06d297c4b2c941eb3607c04d1c1712f33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0455f29832b6c0fb1ff71c4fec4d16346a181764b3ae8b3ba8dc7bc63b4f73ce"></a>

## https_auto_cert.use_mtls.trusted_ca — https_auto_cert.use_mtls.trusted_ca / c03e6b12ac6f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d)
- https_auto_cert.use_mtls.trusted_ca

<a id="canonical-e1dee2498f46677e11018ff7bbd31ee55d1df3438579d2baf9c8c0b1764ac4af"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-81154847a69f7c3d067bb07770316fc1bf380c714580c603f9d8a8c9eda93bbd"></a>

## Direct properties — https_auto_cert.use_mtls.trusted_ca / c03e6b12ac6f / 3

<a id="canonical-ac46599de0e631a86ae89b70cb0212c9e9cbac1d11e12cc8e0275dd526a4cc1f"></a>

<a id="canonical-06e9428d2e637853ef899730610e9c6c2eaefbe32f3dc9c9991caa56ac8642b1"></a>

## name property — https_auto_cert.use_mtls.trusted_ca / c03e6b12ac6f / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-6ed7c7e2a78c5f241e3359d1cf04d2ff02ec7efcd3f839ca20c23fe3663a74f2"></a>

<a id="canonical-80274af7cae9dab095388bc21e3a8fd0064d595398ff137b6008de764b108697"></a>

## namespace property — https_auto_cert.use_mtls.trusted_ca / c03e6b12ac6f / 5

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
}
```

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

<a id="canonical-c125da11f36c71cd94b77a1540f51c93c79e9261cfe6978e61e71f8e53a767f1"></a>

<a id="canonical-5684d192903a65e30f9f5c4cdfd31ba88f3c777c18a4db9f10e414313f496132"></a>

## tenant property — https_auto_cert.use_mtls.trusted_ca / c03e6b12ac6f / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-c44ad262f50ba1a3302328a286052f50f439e7806c5e19d6c461cbece90e39cd"></a>

## Next pages — https_auto_cert.use_mtls.trusted_ca / c03e6b12ac6f / 7

- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e3af7e77748ec2390458db79169261c959b803567987405ea28014fb1193cf8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cd22b871e43638805034dfc45329abc12ada0c7baf5c0dd3929c8b91ac964f27"></a>

## https_auto_cert.use_mtls.xfcc_disabled — https_auto_cert.use_mtls.xfcc_disabled / c2db08208b14 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d)
- https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-ee2343544dadf708c8390474c8b5bed657cd33b50d1f43173dd2c31a8a0d449a"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
xfcc_disabled = {}
```

<a id="canonical-adc175a6065a866200f89d8e870175e8308ad8aa556b37496b0c0cd94d2e1c81"></a>

## Direct properties — https_auto_cert.use_mtls.xfcc_disabled / c2db08208b14 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3d7fbac9a566a755b8f80af51755c9c0366b96fa224b75f586f53d7dc12301ba"></a>

## Next pages — https_auto_cert.use_mtls.xfcc_disabled / c2db08208b14 / 4

- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-dc5dd1806459e7740f0dec52d5f32382a5f23e90422a515acb1085d250eaa0d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87d7d94a1add90296e1a18138200625c9d196aa61d17cd79bb9447d0157f037a"></a>

## https_auto_cert.use_mtls.xfcc_options — https_auto_cert.use_mtls.xfcc_options / 0d78f1aa67c6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-f313e867726d0c24f875066fcaa917f4348e2e777a0878e92631065393f60aa8)
- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d)
- https_auto_cert.use_mtls.xfcc_options

<a id="canonical-f6cd003927a68d07f87b0032822820c72e84ef495a3d8584328281dd27c84f8e"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
```

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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-6c32b8aa9b5720b4de095252f9d3fdf97badd12a5a62040a4e19fce2f365d5d9"></a>

## Direct properties — https_auto_cert.use_mtls.xfcc_options / 0d78f1aa67c6 / 3

<a id="canonical-b319ef23a96429f22b845e1d9618fc191c16b2f0e247fa3e694a28d38448840b"></a>

<a id="canonical-7a052ae33ec39210272def09636ba05b8a07e3f55a655960a5740ef39077fb68"></a>

## xfcc_header_elements property — https_auto_cert.use_mtls.xfcc_options / 0d78f1aa67c6 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-7d1ff012dcdcbb1615adaca5b2522431cd7c0ca3ddb4524cd147c7e65a2a8ba0"></a>

## Next pages — https_auto_cert.use_mtls.xfcc_options / 0d78f1aa67c6 / 5

- [https_auto_cert.use_mtls](resources--http_loadbalancer--reference--group-019.md#canonical-9bba3eb26299916b8419a93da5f300a3fc0aabc62bfa992eb892d5189cbd9a5d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1a9d02dfd2c57fd3a9e31ecff2bb82960607e7bd4a61e0a621dd4cbbf874f729"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6499532c662ad31184ca97e0cd1354ecf342b49c77d4eeed67be5939faa9cd39"></a>

## js_challenge — js_challenge / bd07980e916e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- js_challenge

<a id="canonical-4557dc29e7f082374f2bec8955f6733d0e0c82acc02f9f5f823afb5f2f14e7d0"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript.

With this feature enabled, only clients that are capable of executing Javascript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the Javascript. Javascript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid Javascript challenge for subsequent requests.

Javascript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
```

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
js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-7e6eb292f4ed3ccabdefa48fec9642432b59f666e75fe8e6efe5871efe4be93a"></a>

## Direct properties — js_challenge / bd07980e916e / 3

<a id="canonical-ce7a427ef0d40b812ea064aab50d2c7f92e7e3d401c9e91597c34c9cca652ed6"></a>

<a id="canonical-a98dcc8b3a355844e0db3501a46f9623085c57e82cbc307c7461edb5977231cf"></a>

## cookie_expiry property — js_challenge / bd07980e916e / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-114e07a4186bc19ed0f0958782cee64d5d0ee0d5f2e7f16dde4730bae1b04bc3"></a>

<a id="canonical-733924b2322bf5d62afbec06db760ca872d9537da9b97f246509c534ecb58512"></a>

## custom_page property — js_challenge / bd07980e916e / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-96965c0083ac6f3e8784b7287355ad923c188544529bf817b392744ab9628282"></a>

<a id="canonical-195d40d030cf50a391cbd4da80cb9589cbaeec69f988d08472a53646adc3e3eb"></a>

## js_script_delay property — js_challenge / bd07980e916e / 6

Type: `"number"`. Optional.

Delay introduced by Javascript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-7e1bb20310554b4ddbd48dc0bc00fde05f9f2d9403f09f1a0c4dfd42694f5f70"></a>

## Next pages — js_challenge / bd07980e916e / 7

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4df32e93b4fba0c762fa6eaf59f4319cb5b4db3f51b42644a5d38a9c0965d3ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0dce481584278b466f8ba2d095b0266158845b0aeac4d559580fcbcfaaa5ba9d"></a>

## jwt_validation — jwt_validation / 70f0895252ab / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- jwt_validation

<a id="canonical-eeb4fc1f1b0944ac74318843b98f6d4ebb6d8c5670daa2297b80ca4d142b670a"></a>

Type: `"object"`. single nested block, Optional.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Upstream description:

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("authorization_server",
    "jwks_config")}
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
  "x-ves-oneof-field-jwks_configuration": "[\"authorization_server\",\"jwks_config\"]"
}
```

Terraform syntax:

```terraform
jwt_validation {
  # Configure direct properties listed below.
}
```

<a id="canonical-ffb5946bee0935c51f4dd62cacaac470e669c7a843f22e527ea8a50355177ffc"></a>

## Direct properties — jwt_validation / 70f0895252ab / 3

- [action](resources--http_loadbalancer--reference--group-020.md#canonical-b9201cf244cf27c39eab0c9f115642cc4483bc843faa44dcb385a386bf49c57f): complete subsection reference.

- [authorization_server](resources--http_loadbalancer--reference--group-020.md#canonical-8a0d75570e09c49fa8719c001a88ab35b0956a9973ae355867082dc1752dd674): complete subsection reference.

- [jwks_config](resources--http_loadbalancer--reference--group-020.md#canonical-207912bbcabda8e22d01bf8ca2b79480e8c4733e7149ff9876e43c77f8d44f29): complete subsection reference.

- [mandatory_claims](resources--http_loadbalancer--reference--group-020.md#canonical-c990ea81e10e45c0f4822e07ef17a4b567a240fe6e9ff10ac5dedf426ee2474c): complete subsection reference.

- [reserved_claims](resources--http_loadbalancer--reference--group-020.md#canonical-3a08133a8ca80bbfa9e471c42db51cfd09a8d6eca14bdc98466c89f2ca01016a): complete subsection reference.

- [target](resources--http_loadbalancer--reference--group-020.md#canonical-e3104da7e85673931e1bf16a4ab0a67c12a951ec563354d49d90f8fe5a1f4629): complete subsection reference.

- [token_location](resources--http_loadbalancer--reference--group-020.md#canonical-9073b7d6843748237d154caa9a53d8d56c8367c9123bbaf8de6c42345c346f98): complete subsection reference.
