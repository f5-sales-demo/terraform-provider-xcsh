---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-5b5f2caad3115982165a5a4a78ca7c352a3465a2d6ff86ef203d1ec1acd4534c"></a>

## cipher_suites property — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 56bc8aa45d97 / 4

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

<a id="canonical-6cc1de2fb5aa92ab89e0310f7dfc39912bb970add5dbede62ae91556a5224f1e"></a>

<a id="canonical-1c378fe7a2d743c91fbaaaa5c55e8a7100fe99f07fc45cccfc7f814511b4a23d"></a>

## max_version property — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 56bc8aa45d97 / 5

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

<a id="canonical-1238577c4782d335bfca9abd40d482380838d77f6497a18855e88e1830254b21"></a>

<a id="canonical-7efbd3e709e9a6fa421cb80168c7beaa72b0bfaccaef79a0f31c589af8cfb575"></a>

## min_version property — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 56bc8aa45d97 / 6

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

<a id="canonical-48207fc464a865bbaa44e574fed8ef405a45ea3fc0f3060660262ed9478fdee3"></a>

## Next pages — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 56bc8aa45d97 / 7

- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-4df5a409dd57c563bffde36f9cf494186e776d781835c6e847c44c94ff489dd8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1805f14acbfc1f2d8a48b2de257b79cfcd83583a0c69ead1e2a7aba9e0609879"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8dff5e8db01b309a958f31e256b3cbe6b77b2f75d030fa1bab476b0352ec66f4"></a>

## https.tls_cert_options.tls_cert_params.tls_config.default_security — https.tls_cert_options.tls_cert_params.tls_config.default_security / 8e0fd61fd742 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-4df5a409dd57c563bffde36f9cf494186e776d781835c6e847c44c94ff489dd8)
- https.tls_cert_options.tls_cert_params.tls_config.default_security

<a id="canonical-3618ad45dea9a318ac81570b962d25834aa1ae08850e46a4c4962a5896214465"></a>

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

<a id="canonical-48cdec594e2c3f1d7a19dd911f9507098c8562679265675415a46a0694d13363"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.tls_config.default_security / 8e0fd61fd742 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b43d5047b95bcf7bdeb4f3257ccf4eb293ffa0724a1034f82b66a11c127135e9"></a>

## Next pages — https.tls_cert_options.tls_cert_params.tls_config.default_security / 8e0fd61fd742 / 4

- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-4df5a409dd57c563bffde36f9cf494186e776d781835c6e847c44c94ff489dd8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e17221937085b12d0e8470a728fa6bf2f65a2e8b42116ab37fefc46c1540d0be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-990556837d9d3d64c88a10d16fc78fef2f64a9a6b3d34c4ccba2448b2d40cab5"></a>

## https.tls_cert_options.tls_cert_params.tls_config.low_security — https.tls_cert_options.tls_cert_params.tls_config.low_security / af9a6de7029e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-4df5a409dd57c563bffde36f9cf494186e776d781835c6e847c44c94ff489dd8)
- https.tls_cert_options.tls_cert_params.tls_config.low_security

<a id="canonical-b402a60f60acd06f4219433c1fea091fd0e633413195aa53cbb21f8e95a5afdb"></a>

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

<a id="canonical-985130c4b7b84479ac5f99a09135f65a09e2661cc146c45dac41875d999000cb"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.tls_config.low_security / af9a6de7029e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0af62c4c5d35e9725dd095cbdf9be2eb3751741276fb73189e450d876a5f4b7"></a>

## Next pages — https.tls_cert_options.tls_cert_params.tls_config.low_security / af9a6de7029e / 4

- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-4df5a409dd57c563bffde36f9cf494186e776d781835c6e847c44c94ff489dd8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-bc1838534f155e4cc22e7f4e2b2f415b409f8ef9e3edcddea0957f88e0f75594"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-836a41d8e8ca9ab0e1bf8c331c5c42de9280f6be569fbe47853fa7c1c01c0a99"></a>

## https.tls_cert_options.tls_cert_params.tls_config.medium_security — https.tls_cert_options.tls_cert_params.tls_config.medium_security / 63cedbad368f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-4df5a409dd57c563bffde36f9cf494186e776d781835c6e847c44c94ff489dd8)
- https.tls_cert_options.tls_cert_params.tls_config.medium_security

<a id="canonical-308fb97e44dcd09984c07d7eeb9e4697cb0b37f113ba9e89ce9bfec7bcd2c53d"></a>

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

<a id="canonical-23e8ffea71621c8038faa56f7d9fbe9b71265eceb6e3272c3550d8923a8e11b4"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.tls_config.medium_security / 63cedbad368f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0b5d489ac8dff9ccb03248ea476c14368a0cee13f55b709fec92918ae2572a62"></a>

## Next pages — https.tls_cert_options.tls_cert_params.tls_config.medium_security / 63cedbad368f / 4

- [https.tls_cert_options.tls_cert_params.tls_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-4df5a409dd57c563bffde36f9cf494186e776d781835c6e847c44c94ff489dd8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-378995a21f0efd94b778acfecf4e11b850493f491c89e0f7759403a4ba2168c9"></a>

## https.tls_cert_options.tls_cert_params.use_mtls — https.tls_cert_options.tls_cert_params.use_mtls / 8d383a9670e2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- https.tls_cert_options.tls_cert_params.use_mtls

<a id="canonical-80ceccd572020da8b65f77db7fd8f483adbe6218fb64af1e13d75b71365a722d"></a>

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

<a id="canonical-964aec7ce41216f4599438f5d5a5b16c7f8e8435e01b17dfc26bc2776ddaebf7"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls / 8d383a9670e2 / 3

<a id="canonical-28d0a2a75a0054d5527cef54bedaac5f82bcd4482bfc48170713bf2457e15382"></a>

<a id="canonical-c357f383b8ff8e8d86a44d2a39c5e57b5961b22fa21e0d4d570b4a0c9eaf917c"></a>

## client_certificate_optional property — https.tls_cert_options.tls_cert_params.use_mtls / 8d383a9670e2 / 4

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

- [crl](resources--cdn_loadbalancer--reference--group-011.md#canonical-d74e0b83d9532b01925e84d8488d6f78c87a7123ef646a84fa29d601cd2b8565): complete subsection reference.

- [no_crl](resources--cdn_loadbalancer--reference--group-011.md#canonical-4f0cd8340b2aa944cd408782b630d17f0134067c5a1ca93727addefb1ffa2372): complete subsection reference.

- [trusted_ca](resources--cdn_loadbalancer--reference--group-011.md#canonical-f41ded3e2e6f8c0955873b0abed9a59d38f5a99605d9810391dfab4383c26bec): complete subsection reference.

<a id="canonical-2f96c3f45b00432053e6d5bfa1dff33d9de25df17d8a47e3c93f4edbe07a422b"></a>

<a id="canonical-b66c7c81183310ab599facbacfaa252ab88c6d838917af7e652eec2b38d1b677"></a>

## trusted_ca_url property — https.tls_cert_options.tls_cert_params.use_mtls / 8d383a9670e2 / 5

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

- [xfcc_disabled](resources--cdn_loadbalancer--reference--group-011.md#canonical-5d8a736c91b21b55971f2afca6b08efbd293811f625af7ff9803b7653b8f71df): complete subsection reference.

- [xfcc_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-4ecf82ae0d6fec1952766bd5e3df2f2a811565fc870ac0111ac492da47b72641): complete subsection reference.

<a id="canonical-b88e21dd3b24eba62cffa61c5a2dda68235f45fd60fa7054866f169aaf2771e8"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls / 8d383a9670e2 / 6

- [https.tls_cert_options.tls_cert_params.use_mtls.crl](resources--cdn_loadbalancer--reference--group-011.md#canonical-d74e0b83d9532b01925e84d8488d6f78c87a7123ef646a84fa29d601cd2b8565)
- [https.tls_cert_options.tls_cert_params.use_mtls.no_crl](resources--cdn_loadbalancer--reference--group-011.md#canonical-4f0cd8340b2aa944cd408782b630d17f0134067c5a1ca93727addefb1ffa2372)
- [https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca](resources--cdn_loadbalancer--reference--group-011.md#canonical-f41ded3e2e6f8c0955873b0abed9a59d38f5a99605d9810391dfab4383c26bec)
- [https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled](resources--cdn_loadbalancer--reference--group-011.md#canonical-5d8a736c91b21b55971f2afca6b08efbd293811f625af7ff9803b7653b8f71df)
- [https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-4ecf82ae0d6fec1952766bd5e3df2f2a811565fc870ac0111ac492da47b72641)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-d74e0b83d9532b01925e84d8488d6f78c87a7123ef646a84fa29d601cd2b8565"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-453d3a112a866286db1995287bdc2bddab3834f88b7b44dc181a7ecceeac0d4b"></a>

## https.tls_cert_options.tls_cert_params.use_mtls.crl — https.tls_cert_options.tls_cert_params.use_mtls.crl / 898dcf3f289c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8)
- https.tls_cert_options.tls_cert_params.use_mtls.crl

<a id="canonical-b3d248fc0ccb8b156abe7671df4d07b28e42deed106fa0c1c14f284d8a38fd5b"></a>

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

<a id="canonical-4f7e9672136d77a39edaacb281fd10ac4c768b3514f5bc8a5ab09043b1df2a90"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls.crl / 898dcf3f289c / 3

<a id="canonical-bb16277985fe980c19e362fcb649ece33f7f67e2f34ed76e4bbb6c8945c71c0e"></a>

<a id="canonical-d3e09e03da096631f4d637e2cd2a6689f6aec525ca3701f3bd88ece18e8821f7"></a>

## name property — https.tls_cert_options.tls_cert_params.use_mtls.crl / 898dcf3f289c / 4

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

<a id="canonical-a9181a8043e1a79dfc67565c525b06805e4a1eab3305db91569e7b4b5c0efbc4"></a>

<a id="canonical-60cbb54357cb4cc93ab3fea26e349c3660fd923cb58f4d0f6e2bf405ff22a378"></a>

## namespace property — https.tls_cert_options.tls_cert_params.use_mtls.crl / 898dcf3f289c / 5

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

<a id="canonical-194b0205091a70903aae76bc1da546d47509399501e4d7a137ca72aef419acc8"></a>

<a id="canonical-596779d819cf25f3f718fac8f0d6a6eed120ad644ddb1e7a71fb220620543fbb"></a>

## tenant property — https.tls_cert_options.tls_cert_params.use_mtls.crl / 898dcf3f289c / 6

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

<a id="canonical-a36b4d7d6d47a91b9d4b03532d2ed5a5dc6cfd9685bcfa072925464c5b6c6d66"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls.crl / 898dcf3f289c / 7

- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4f0cd8340b2aa944cd408782b630d17f0134067c5a1ca93727addefb1ffa2372"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-580b369d84cedf523d8a7623691ac0e196f364ce55adba413ad16a9047d19ba0"></a>

## https.tls_cert_options.tls_cert_params.use_mtls.no_crl — https.tls_cert_options.tls_cert_params.use_mtls.no_crl / dcede474e0d6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8)
- https.tls_cert_options.tls_cert_params.use_mtls.no_crl

<a id="canonical-8a8db3ba550f625ca1f0a70274c7060c3a998f160478238685b3955324c96776"></a>

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

<a id="canonical-838e9ca7e5ed14df16ee37f2a22ddcba28ac38358f0be9683c6029feccdcb569"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls.no_crl / dcede474e0d6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dca8d3f313b5d07ae042a5e63810e2ffc495af9fcf677ca1f737cc32ab063ead"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls.no_crl / dcede474e0d6 / 4

- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f41ded3e2e6f8c0955873b0abed9a59d38f5a99605d9810391dfab4383c26bec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa08288d5db310d3224c32cb840fc54f0b37990da6b8a79f7d86f497ab18c4af"></a>

## https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / 3209abc5cbf3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8)
- https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-0069457185b0db8fa4cbf0a3585c4f9d86bdf9ddcbb43bc639cf1a79e89ad663"></a>

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

<a id="canonical-118de186dc1dee2edc94268e09b32fd4e85070b1160372d919a6451003e42f74"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / 3209abc5cbf3 / 3

<a id="canonical-915390a72841597ac616c234253d4c8f9bcc2480cb0eb59cca75c3d4c6f03719"></a>

<a id="canonical-9f28a4b70d5b586c374aa1a8e9c53d69cd3732f0da3cd94003e1b146497b5aea"></a>

## name property — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / 3209abc5cbf3 / 4

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

<a id="canonical-3cc52096fda74c838390bfc0c3b3721a20b9f5830d235c38bb1e84df614b16ae"></a>

<a id="canonical-753e0b5ea3f5785f748c6d667f631026646d895ff890dfc5226468b0e1e3018a"></a>

## namespace property — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / 3209abc5cbf3 / 5

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

<a id="canonical-858c3c8ad8331a15248bcffbd150b511fb327d05258a7dd1d1cb2005ebb30d60"></a>

<a id="canonical-0dc71b56ee57162f02b6f44493e5edbb79e1c0f26f68fcfd859c675558c093ec"></a>

## tenant property — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / 3209abc5cbf3 / 6

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

<a id="canonical-436f37f070721d33de811a6bb23679d7ce48e411e758f0762a54e520f58da194"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / 3209abc5cbf3 / 7

- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5d8a736c91b21b55971f2afca6b08efbd293811f625af7ff9803b7653b8f71df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94efa878f7af01b6c381778b19b12a779e899855688f88d6aea1066e1b5c97e6"></a>

## https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled / 29de7064992c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8)
- https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-1bc5796c6b3d15e28c395acf4e3fbf5085525760a6b84054fede66a71ee67388"></a>

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

<a id="canonical-30149cfc87acb37d41af993158e296ddc6cbeb1979debc3ece3e005a14286f98"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled / 29de7064992c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a14d3ade2443fe155ae2e38549866fd742b9ff52efbf84a44734a31a08bb4d10"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled / 29de7064992c / 4

- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4ecf82ae0d6fec1952766bd5e3df2f2a811565fc870ac0111ac492da47b72641"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0eed5cb527c7737afbcbca8b3621dae0b78c7abbd5cbdbed6f1ba3dd1a30205"></a>

## https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options / 103a3cd590dd / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_cert_params](resources--cdn_loadbalancer--reference--group-010.md#canonical-fadff03b94e0c5d74340d2ed58c1f2d8a3e84bb577ce7b31f80c2165765711d1)
- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8)
- https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-ef048bf60b5c765cff58c1c5215854a8ac90ce2899f950b85b7a4bc43a689c71"></a>

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

<a id="canonical-7877b071296a1a7f4e0af07a49494ef58e8fb714c71b8fe998d22f4360ac570c"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options / 103a3cd590dd / 3

<a id="canonical-7be0f94869a5eb1bb3371bc9e5f763fc9ce50c19f9861bef71a1107860c20c91"></a>

<a id="canonical-8ebfc1da68dc5bc9e3f30a292dc6608e3171b8c69b95c7d682e32490effe71a3"></a>

## xfcc_header_elements property — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options / 103a3cd590dd / 4

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

<a id="canonical-dcf09f8ce787abeef6270e5b648d0352c192190af2212aa9bb03a13a57525d78"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options / 103a3cd590dd / 5

- [https.tls_cert_options.tls_cert_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-8735be1ccb4030c11c11db4bd0916a0b2db0891a407773421507e28f12bbdeb8)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c96fe22762271601ddca08b502e088d7bf77fa39e9f8e4818bdba823848eab4"></a>

## https.tls_cert_options.tls_inline_params — https.tls_cert_options.tls_inline_params / e8b9297975ea / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- https.tls_cert_options.tls_inline_params

<a id="canonical-cd15c0db17fc764b2a47fbbee7267e2642d208eb418806a4112f0007970249c2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls inline params.

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
tls_inline_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-493be39ba5caf546fe60f786b1d9a6e8ad7fce9c9aa4c3e95c389b098013a083"></a>

## Direct properties — https.tls_cert_options.tls_inline_params / e8b9297975ea / 3

- [no_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-6c3571d33a14405ff358b9df8b9c5b194874eb53c5fb59a17bcb29349b3ac094): complete subsection reference.

- [tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a): complete subsection reference.

- [tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-3eae0557621e97c9aeaf1b8d9c57fcb45f933c63d2b2007d5c061909ec3ede6f): complete subsection reference.

- [use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5): complete subsection reference.

<a id="canonical-907ad9cbcd33dd7bf9cb17cc2a6fbf2d962b7964673e0f33a97d3dcfd2e0bae1"></a>

## Next pages — https.tls_cert_options.tls_inline_params / e8b9297975ea / 4

- [https.tls_cert_options.tls_inline_params.no_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-6c3571d33a14405ff358b9df8b9c5b194874eb53c5fb59a17bcb29349b3ac094)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a)
- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-3eae0557621e97c9aeaf1b8d9c57fcb45f933c63d2b2007d5c061909ec3ede6f)
- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6c3571d33a14405ff358b9df8b9c5b194874eb53c5fb59a17bcb29349b3ac094"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d40490bc1e979b7604b396130f9f5d692c883df4bc32e94c281c24514bd2fb8b"></a>

## https.tls_cert_options.tls_inline_params.no_mtls — https.tls_cert_options.tls_inline_params.no_mtls / a0ed0a650acc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- https.tls_cert_options.tls_inline_params.no_mtls

<a id="canonical-85d9e63141d4b5c1aa554293134aa239e4a5ffdf6b2b6503de1c7694c8acd44b"></a>

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

<a id="canonical-90752f1ec7b16ed332b90cdbc8c8cec493638ae9aa9538002f8a46bc601c0e19"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.no_mtls / a0ed0a650acc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-08f368fb22b211dbad85be458659b1c0fccc68bf9528ace206144f3694ed464d"></a>

## Next pages — https.tls_cert_options.tls_inline_params.no_mtls / a0ed0a650acc / 4

- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b07d1b96f2a5288e2df88115e389d123087f67717fe46ab1190b000434975267"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates — https.tls_cert_options.tls_inline_params.tls_certificates / e976060db9f1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- https.tls_cert_options.tls_inline_params.tls_certificates

<a id="canonical-6fe6a80415e6402f1d318f92cec0a20d0161869bec2973f90b5e75bd03e878d3"></a>

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

<a id="canonical-55e00216c1f13c08a7b617c259f7d4e5e92c0b0b6af0554ed10d056e235cef6f"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates / e976060db9f1 / 3

<a id="canonical-c8c63187e125d9c9c4f8c3c83a5b531fe249142bcc5a5710ccbc9277a4e0a18b"></a>

<a id="canonical-df177e5a39e1ad3bf993f5dd644f73ff2275d9478e04e35fc3c3d15a696eda3c"></a>

## certificate_url property — https.tls_cert_options.tls_inline_params.tls_certificates / e976060db9f1 / 4

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

- [custom_hash_algorithms](resources--cdn_loadbalancer--reference--group-011.md#canonical-dbc351f9f446adab86f39c46a9651a2e2821b7172727acfb79e8a4be997d6130): complete subsection reference.

<a id="canonical-c59b96931d50574ae68462c586b15473e5560d9765ffd9926047ff1962166926"></a>

<a id="canonical-d71239fb8a8df4c6b4debb3716e855c4ffa1d1cade3fe8aab7cf8eaf3995a03f"></a>

## description_spec property — https.tls_cert_options.tls_inline_params.tls_certificates / e976060db9f1 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--cdn_loadbalancer--reference--group-011.md#canonical-66187354de1144925f0fa69f8a3125a14513f5f0535651609dcb39899af58ddf): complete subsection reference.

- [private_key](resources--cdn_loadbalancer--reference--group-011.md#canonical-f66a4ed48a3c1e95242275a9d59347057befb3b7f8f57f4c53efb5e4c0d86273): complete subsection reference.

- [use_system_defaults](resources--cdn_loadbalancer--reference--group-011.md#canonical-080da4863b9e5255d3c2603d9f2d704efb175002701c2b57dfc9163dd9817bab): complete subsection reference.

<a id="canonical-b0b5af9e5b93e03039178efb7cf1a0390e4e4e6c063b34009a9732518d574c79"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates / e976060db9f1 / 6

- [https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms](resources--cdn_loadbalancer--reference--group-011.md#canonical-dbc351f9f446adab86f39c46a9651a2e2821b7172727acfb79e8a4be997d6130)
- [https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling](resources--cdn_loadbalancer--reference--group-011.md#canonical-66187354de1144925f0fa69f8a3125a14513f5f0535651609dcb39899af58ddf)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-011.md#canonical-f66a4ed48a3c1e95242275a9d59347057befb3b7f8f57f4c53efb5e4c0d86273)
- [https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults](resources--cdn_loadbalancer--reference--group-011.md#canonical-080da4863b9e5255d3c2603d9f2d704efb175002701c2b57dfc9163dd9817bab)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-dbc351f9f446adab86f39c46a9651a2e2821b7172727acfb79e8a4be997d6130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b049f60e3116263b0b5c704f14ea47fd0788100c4487130784183041ff29db2"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms — https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms / 685ed391bfb7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a)
- https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms

<a id="canonical-41221c35c279dbdd79461d4a620d85d7775f206a648316cd9bcc49bc51c708c1"></a>

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

<a id="canonical-b23c71ef8fdeccba0e79b16407b14afb9a039df75615d4703b16606b3e52ea12"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms / 685ed391bfb7 / 3

<a id="canonical-cabb9112440055201545e3cda539b7187fa20fe2d3b7915e01285be87eb23a3e"></a>

<a id="canonical-18933c7206922d51e4dafe2dd44f88292d5e6ef569273ffa041a7a0b2d69a2ba"></a>

## hash_algorithms property — https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms / 685ed391bfb7 / 4

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

<a id="canonical-7bc80641b39e492eb9976ebc3b28c55106d82ddc8245332cc33eee5986dfc43e"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms / 685ed391bfb7 / 5

- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-66187354de1144925f0fa69f8a3125a14513f5f0535651609dcb39899af58ddf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-145d711d1a5a4ce6fbc83a58cac9279cbc8780dd56c608619cff0e3471e7ebc6"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling — https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling / 440bed52196d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a)
- https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-1d978e5876ab62122abc2b58c400418d3fbe414248e0eb98f2176fc3c9e54337"></a>

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

<a id="canonical-f45bed1f050ccdc66c7e95f3878f61ef041e4ac263e4acddc595cc14a12a3090"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling / 440bed52196d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1da098a0bafacbc1266ef05c1cb511eb24a0ecb70e69ddfecf87c868a8647634"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling / 440bed52196d / 4

- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f66a4ed48a3c1e95242275a9d59347057befb3b7f8f57f4c53efb5e4c0d86273"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cdc497633406f33f6fb8f0da2ca093a7d2c3d0341751851f612ab55b8fa3737f"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.private_key — https.tls_cert_options.tls_inline_params.tls_certificates.private_key / c59da640c029 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key

<a id="canonical-8bcbcb4da3fb7061c80dc2ad292f41a87dffaacb5e13eecbafde1797971af8fe"></a>

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

<a id="canonical-fe47e59b8fbbc818227f7a542c9dc401a9e045275403e7731f1909e45a68951f"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.private_key / c59da640c029 / 3

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-011.md#canonical-93662784c0dcdfa374cfb10fccd6def2fad0bb207f87fc865193fcf3777e5347): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-011.md#canonical-9e5827ac4107e1932e076f4bc990abfcb7f5e07635107b38743cb8ef0bd06120): complete subsection reference.

<a id="canonical-28168efbcb672d8184b2cde7a1ffb7ea512de121fb7c705fb74e4dc2f0dca8ea"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.private_key / c59da640c029 / 4

- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info](resources--cdn_loadbalancer--reference--group-011.md#canonical-93662784c0dcdfa374cfb10fccd6def2fad0bb207f87fc865193fcf3777e5347)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info](resources--cdn_loadbalancer--reference--group-011.md#canonical-9e5827ac4107e1932e076f4bc990abfcb7f5e07635107b38743cb8ef0bd06120)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-93662784c0dcdfa374cfb10fccd6def2fad0bb207f87fc865193fcf3777e5347"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60749c3429aa17e4b73c0bd7bef21c59ded0a567f44ecb59f51fc43541aeefb1"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / db7127fa483f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-011.md#canonical-f66a4ed48a3c1e95242275a9d59347057befb3b7f8f57f4c53efb5e4c0d86273)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2f786df5ff9693138ef9c5d5b20a12f9757c812079ca5e46b91cbe37e6080918"></a>

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

<a id="canonical-d8b2c0aa4803a44b846539bfaf5aea1d68c0f658c36657574eafefe16eeafa19"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / db7127fa483f / 3

<a id="canonical-22b37dedf32c005058709e1f9975a26378d5fed6316c955e17791d092e41e1b6"></a>

<a id="canonical-910c570912411a83bcd79f9948ac374cf82db3c248b8712d606ac8f764f1bbb1"></a>

## decryption_provider property — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / db7127fa483f / 4

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

<a id="canonical-8994ac490d37236d7264a7ea803bc740dc012d3b51d789b6c485c9c831a3cac8"></a>

<a id="canonical-6d0e9dd3e83bb0379bdcec81ee82e1c167d037a7f689d562e9c86e9fdeb2306a"></a>

## location property — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / db7127fa483f / 5

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

<a id="canonical-f49d654369514a120502e719732d1b745ee325f1ef98d482a307dbe353788bf6"></a>

<a id="canonical-b5564fd3026a1c48e27e81a8cfda6cc23f04d63f8d0ef371d695177f492418c4"></a>

## store_provider property — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / db7127fa483f / 6

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

<a id="canonical-678fc2a57f3aa33d12e7dce2d12331f793e1e74639f3bca1d5deab752b5b0871"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / db7127fa483f / 7

- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-011.md#canonical-f66a4ed48a3c1e95242275a9d59347057befb3b7f8f57f4c53efb5e4c0d86273)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9e5827ac4107e1932e076f4bc990abfcb7f5e07635107b38743cb8ef0bd06120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-428f6fda9007b2daa3d8a05c38264c45bc2f63c4a4ad03d49e64a80866187b2e"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secr / b3fa11816a12 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-011.md#canonical-f66a4ed48a3c1e95242275a9d59347057befb3b7f8f57f4c53efb5e4c0d86273)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-f34cc1d5fbc427554b59457b28c903fe1e9fa5eba480b46ca38da916955df51a"></a>

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

<a id="canonical-a9c7a9e13a728dbb7532b51f090fcf39f298bf77d0b3af636b6559d1513dc5a8"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secr / b3fa11816a12 / 3

<a id="canonical-ec3d5a570c61ca6e2259955568f81300c9876b0417d4ab8e626526b3fc2a4089"></a>

<a id="canonical-9c4e668aa712272d5ddc5efa47967a991a42f2a51afbecacd1ea4b2b2e680789"></a>

## provider_ref property — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secr / b3fa11816a12 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-763cb518d5833881ac83f93284689c104a959b7b099111bb47d72bd84d20c8b5"></a>

<a id="canonical-5427b02b1e0775e4e995f8c2e3d3eed50c09a6713ae5358925773c02d4da44d2"></a>

## url property — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secr / b3fa11816a12 / 5

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

<a id="canonical-d953d99bc74f4015c40ea3de5362d053e9ac39091587cb42867d6f7114ae0c0e"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secr / b3fa11816a12 / 6

- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-011.md#canonical-f66a4ed48a3c1e95242275a9d59347057befb3b7f8f57f4c53efb5e4c0d86273)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-080da4863b9e5255d3c2603d9f2d704efb175002701c2b57dfc9163dd9817bab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e11f00ce14ce7a9f5e5998c9f0e3093b2b2a6dd42eb4c04e6f2dfe65570fbaf"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults — https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults / ff746760f6e1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a)
- https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults

<a id="canonical-2a16a373da9180f3758489752d47026a72d945c995a9e428854cbe537cd04cf6"></a>

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

<a id="canonical-147b2551f1ab906e2278737b5a41e0f85a435f85f755675515df5eb80e66e68a"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults / ff746760f6e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-399ba8f5e0c5e03528237f11754a325111d17323011898d48916793269e34b56"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults / ff746760f6e1 / 4

- [https.tls_cert_options.tls_inline_params.tls_certificates](resources--cdn_loadbalancer--reference--group-011.md#canonical-5c2bfb8a1e9dd38910163723177be6b8064af9c4a25449166919cfb9d0f59a9a)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-3eae0557621e97c9aeaf1b8d9c57fcb45f933c63d2b2007d5c061909ec3ede6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d049610e9cad45e3daf93c8f8cda4a6c8c29867cd0c56737ff67b4a8ce4ef0a7"></a>

## https.tls_cert_options.tls_inline_params.tls_config — https.tls_cert_options.tls_inline_params.tls_config / 0d180425f87b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- https.tls_cert_options.tls_inline_params.tls_config

<a id="canonical-3cc78defbfda5f4016992f0ee26c484047200160a58ba56f64e3e8eb75fcad67"></a>

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

<a id="canonical-9dee759d7e331a0e24039a38204c14f0a6571e3e47882c73c40977015c4d468a"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_config / 0d180425f87b / 3

- [custom_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-33c1c08fac3d1ebb4826a7b6c357241a63941f63bc099be733c156b33819e62b): complete subsection reference.

- [default_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-74994b05a85c14b4662b41c308684e64b8b79957761648526e7ba778bdd0820c): complete subsection reference.

- [low_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-6bcdc09eabc5174feca664e7751565f67c90dda106c45f61666d82c02152361f): complete subsection reference.

- [medium_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-f7446b923bf4f818caf16c15f82c3956413d0ac5a69547bd1414f74b0102f57e): complete subsection reference.

<a id="canonical-4bdb03a4cbb28114615087095c5d48da0a49acaa2cc21d99f6b2a1fb7b249bc3"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_config / 0d180425f87b / 4

- [https.tls_cert_options.tls_inline_params.tls_config.custom_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-33c1c08fac3d1ebb4826a7b6c357241a63941f63bc099be733c156b33819e62b)
- [https.tls_cert_options.tls_inline_params.tls_config.default_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-74994b05a85c14b4662b41c308684e64b8b79957761648526e7ba778bdd0820c)
- [https.tls_cert_options.tls_inline_params.tls_config.low_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-6bcdc09eabc5174feca664e7751565f67c90dda106c45f61666d82c02152361f)
- [https.tls_cert_options.tls_inline_params.tls_config.medium_security](resources--cdn_loadbalancer--reference--group-011.md#canonical-f7446b923bf4f818caf16c15f82c3956413d0ac5a69547bd1414f74b0102f57e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-33c1c08fac3d1ebb4826a7b6c357241a63941f63bc099be733c156b33819e62b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2cba82ae3526fa1693a1cde05dc8da32b9f0104d3978a2650008cc460919cf45"></a>

## https.tls_cert_options.tls_inline_params.tls_config.custom_security — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 099c6767daf5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-3eae0557621e97c9aeaf1b8d9c57fcb45f933c63d2b2007d5c061909ec3ede6f)
- https.tls_cert_options.tls_inline_params.tls_config.custom_security

<a id="canonical-fbe6a9247afc5080bf683ca472c7d36e7dee883353b1b27c33e7ac5bf272c19f"></a>

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

<a id="canonical-4dee4777eda95f2313e64790f0b83a4e6413878fca8444f8fc1f12b98b55df80"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 099c6767daf5 / 3

<a id="canonical-0d85e57b25cea31a255453b1710e1318ee61b724cbdbeaf0a691743d7d242cba"></a>

<a id="canonical-603d035d3448a1b3788017d447fe025f1c519c21e3252cee5ff27a9e93d6895b"></a>

## cipher_suites property — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 099c6767daf5 / 4

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

<a id="canonical-a2fcf54d530b124a5c85696c3adc907b5ed9674791e9a5beaa70df78a5d9e182"></a>

<a id="canonical-527d2b51f753fae36a2daef4192c378012348d007673b6819b6bb4567c3a042b"></a>

## max_version property — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 099c6767daf5 / 5

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

<a id="canonical-37262e962648a6b0d37a700deb050646d3ef469d61eaa683bfb0fa101d8ee587"></a>

<a id="canonical-46ef627fde9fc52c76d624bf799b89cc0aaff64f04115beea0c0602e8c0a668c"></a>

## min_version property — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 099c6767daf5 / 6

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

<a id="canonical-471e4234e7ee7d4f0335281dc67ca4419836dd759531af2c06070a6814e431d8"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 099c6767daf5 / 7

- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-3eae0557621e97c9aeaf1b8d9c57fcb45f933c63d2b2007d5c061909ec3ede6f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-74994b05a85c14b4662b41c308684e64b8b79957761648526e7ba778bdd0820c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04f6923ba7446613015ac9b37363cfad534c0e6138d2d75781d99475a47c4fbd"></a>

## https.tls_cert_options.tls_inline_params.tls_config.default_security — https.tls_cert_options.tls_inline_params.tls_config.default_security / c719ef6aa513 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-3eae0557621e97c9aeaf1b8d9c57fcb45f933c63d2b2007d5c061909ec3ede6f)
- https.tls_cert_options.tls_inline_params.tls_config.default_security

<a id="canonical-4459b4cf2ebb917c5d35a3ca7afff2a079a90faf1e38e9c125363381695e4835"></a>

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

<a id="canonical-2ad96855449b3d23a77daea34a4b8292405361c3139b65c2ad92b4f7c005f2ae"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_config.default_security / c719ef6aa513 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a8def0c0c95b81175a42e1fe2b01a7f461433aec5327ff8037d7b61462534baf"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_config.default_security / c719ef6aa513 / 4

- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-3eae0557621e97c9aeaf1b8d9c57fcb45f933c63d2b2007d5c061909ec3ede6f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-6bcdc09eabc5174feca664e7751565f67c90dda106c45f61666d82c02152361f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fed26ccf6a71c377f257e8c9bdbf7dacd2317696825d322af8a9c74050bea1f0"></a>

## https.tls_cert_options.tls_inline_params.tls_config.low_security — https.tls_cert_options.tls_inline_params.tls_config.low_security / 9c02b4a98847 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-3eae0557621e97c9aeaf1b8d9c57fcb45f933c63d2b2007d5c061909ec3ede6f)
- https.tls_cert_options.tls_inline_params.tls_config.low_security

<a id="canonical-b63b7d892e10ca6c0537b89fa3444376beb5c007f98d9f93e4c6d41b47149bf7"></a>

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

<a id="canonical-633c305295825c270538cc11af59e70fbe34da311f09ab8c6898f3c411918c18"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_config.low_security / 9c02b4a98847 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-926ec821d398af75b1b6f2dc9565ac26ce05148529b6ba48478b8460afb8ec93"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_config.low_security / 9c02b4a98847 / 4

- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-3eae0557621e97c9aeaf1b8d9c57fcb45f933c63d2b2007d5c061909ec3ede6f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-f7446b923bf4f818caf16c15f82c3956413d0ac5a69547bd1414f74b0102f57e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-137d4bb46e403f0477979eaf113af6bba6a46cf551bc60f680f343cf039b51be"></a>

## https.tls_cert_options.tls_inline_params.tls_config.medium_security — https.tls_cert_options.tls_inline_params.tls_config.medium_security / a082d3570866 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-3eae0557621e97c9aeaf1b8d9c57fcb45f933c63d2b2007d5c061909ec3ede6f)
- https.tls_cert_options.tls_inline_params.tls_config.medium_security

<a id="canonical-eb00024d358f2c8cf4b5fc7575cd600db41ab65cff78ffef0f2aa23f80757e90"></a>

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

<a id="canonical-c2f0dc91b2b0c0a6551b417b332b89043a4c4d6a9ba6a68ff3e68d4deef9cbfe"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_config.medium_security / a082d3570866 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f97ffff665b326ad17a7e62af2659d147334dde9b1348d2e6fe2b2061e2e8c6a"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_config.medium_security / a082d3570866 / 4

- [https.tls_cert_options.tls_inline_params.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-3eae0557621e97c9aeaf1b8d9c57fcb45f933c63d2b2007d5c061909ec3ede6f)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05a5ef502b351ea5b635ca661520fee9d453374c2e0c8f0061f24140d41d2656"></a>

## https.tls_cert_options.tls_inline_params.use_mtls — https.tls_cert_options.tls_inline_params.use_mtls / fcb5f5db0807 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- https.tls_cert_options.tls_inline_params.use_mtls

<a id="canonical-81c51dec8239b8bf1d008eb8c810341c42ba5809ba461c99ddd8ff5f58ecc9ec"></a>

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

<a id="canonical-56b5fbbfff3469975c7a6e5521a32163094f14ea0a6832221900f91cbf10b045"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls / fcb5f5db0807 / 3

<a id="canonical-5a9abf9ffc36e27453fb16d7fb5383420afd776c53560e590ceee83813c7db21"></a>

<a id="canonical-b61d691b51dea189e44ab1f7d0eef8fd782245cdc9181b25d2b8d2abe0448e97"></a>

## client_certificate_optional property — https.tls_cert_options.tls_inline_params.use_mtls / fcb5f5db0807 / 4

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

- [crl](resources--cdn_loadbalancer--reference--group-011.md#canonical-45e00ba207def37017391cac9a0623c235fa2d123fee4b45046fcc6d3c420f5b): complete subsection reference.

- [no_crl](resources--cdn_loadbalancer--reference--group-011.md#canonical-c05fba92331ccb5a4fdbccae8379a57aa4261ead8c294e18a9cad0b23c07eea3): complete subsection reference.

- [trusted_ca](resources--cdn_loadbalancer--reference--group-011.md#canonical-4901aa57f5600866f41ac9cfdbe3d5beee808105edb1ea3dd44902549a05ca5e): complete subsection reference.

<a id="canonical-c6efe4ed91b2661466fcbb8890ce43db511229f2494c766533427390b8544c67"></a>

<a id="canonical-3112b28f4df3cce309a5b41e121ffff7260e862e423d02b1759be3c483db5741"></a>

## trusted_ca_url property — https.tls_cert_options.tls_inline_params.use_mtls / fcb5f5db0807 / 5

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

- [xfcc_disabled](resources--cdn_loadbalancer--reference--group-011.md#canonical-b6891065c1ae69c18d3812fb6df3d1655605dc43ec0b19d4987531e913590704): complete subsection reference.

- [xfcc_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-1d282bba2f4940a95245f6249cb212f52f671211241b8c286667dd488b1e5a62): complete subsection reference.

<a id="canonical-8989be0f96cb7e4a72665952cf4197f63c1be8e720fe2c68f44802048c600df9"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls / fcb5f5db0807 / 6

- [https.tls_cert_options.tls_inline_params.use_mtls.crl](resources--cdn_loadbalancer--reference--group-011.md#canonical-45e00ba207def37017391cac9a0623c235fa2d123fee4b45046fcc6d3c420f5b)
- [https.tls_cert_options.tls_inline_params.use_mtls.no_crl](resources--cdn_loadbalancer--reference--group-011.md#canonical-c05fba92331ccb5a4fdbccae8379a57aa4261ead8c294e18a9cad0b23c07eea3)
- [https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca](resources--cdn_loadbalancer--reference--group-011.md#canonical-4901aa57f5600866f41ac9cfdbe3d5beee808105edb1ea3dd44902549a05ca5e)
- [https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled](resources--cdn_loadbalancer--reference--group-011.md#canonical-b6891065c1ae69c18d3812fb6df3d1655605dc43ec0b19d4987531e913590704)
- [https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options](resources--cdn_loadbalancer--reference--group-011.md#canonical-1d282bba2f4940a95245f6249cb212f52f671211241b8c286667dd488b1e5a62)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-45e00ba207def37017391cac9a0623c235fa2d123fee4b45046fcc6d3c420f5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3208bf81930007fe1f901218afae1035c79a8ab76999ad95d400a7a9551876df"></a>

## https.tls_cert_options.tls_inline_params.use_mtls.crl — https.tls_cert_options.tls_inline_params.use_mtls.crl / 3459ecb1cf5e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5)
- https.tls_cert_options.tls_inline_params.use_mtls.crl

<a id="canonical-ebf766e38b5df9ea18b97cb6e2bb561d6c37390ff71fbad327b9f5f69158c2f4"></a>

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

<a id="canonical-691b6c05773951aed9048b8dace01f6aa606768f648ad0634a564df8cb3f4612"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls.crl / 3459ecb1cf5e / 3

<a id="canonical-ca60b1dde77bfaa3cf4d6c48e9f99da3f691205f918ac8703ffe46a26ebdd711"></a>

<a id="canonical-c2ffc324644900f95bc69d3ece91513487d40fce1940a38dc2bddffb0557b55a"></a>

## name property — https.tls_cert_options.tls_inline_params.use_mtls.crl / 3459ecb1cf5e / 4

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

<a id="canonical-5f013ea1c8995965fb3d028c0ce9183cdb8b5a9eb01282ff0fa42548287cd3c3"></a>

<a id="canonical-f9bdecc138b9a78bdb9242e66a007673e47c757427bffb28f39fc71408405895"></a>

## namespace property — https.tls_cert_options.tls_inline_params.use_mtls.crl / 3459ecb1cf5e / 5

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

<a id="canonical-da9f04e6b05b7ff14ba24baee89b245ebe86174dd3517a22fd4fd94ecd59783f"></a>

<a id="canonical-434d9bfe0fd7bb9eaa73ca09d66a2e64c351f983f83bcb050a7ed2393056a0e7"></a>

## tenant property — https.tls_cert_options.tls_inline_params.use_mtls.crl / 3459ecb1cf5e / 6

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

<a id="canonical-d15accdc97161e890baf58a639a66b2c683e98aca2d3cd6aef8dde12c60fc20c"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls.crl / 3459ecb1cf5e / 7

- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c05fba92331ccb5a4fdbccae8379a57aa4261ead8c294e18a9cad0b23c07eea3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b6a7c654f2d0e88c46dd5f7674aae16ec1a6ddb8a21d3dff5701ff19e170631"></a>

## https.tls_cert_options.tls_inline_params.use_mtls.no_crl — https.tls_cert_options.tls_inline_params.use_mtls.no_crl / 1408de81957b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5)
- https.tls_cert_options.tls_inline_params.use_mtls.no_crl

<a id="canonical-053030539c843730f394f2395c531fd90c2a9c5711204f1b4720efb55e29c2a3"></a>

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

<a id="canonical-e4ce942f85ddbf25f87fdff64eeb844375318865695343487cd6c5568d10bcb8"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls.no_crl / 1408de81957b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6d462f4cf3f2aba962cd117ea58dfacaf12c09c2de9b19cf1294e406a3a5f7b2"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls.no_crl / 1408de81957b / 4

- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-4901aa57f5600866f41ac9cfdbe3d5beee808105edb1ea3dd44902549a05ca5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7c4e6fc0000d2e0d5bba25b20c1d05a62aca889ee5bb4762394b5bedaf36966"></a>

## https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / d274da873da1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5)
- https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca

<a id="canonical-d677057299299a620723b87991e9b3ff040f2eb8ae0b74e250d7cd6665e703aa"></a>

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

<a id="canonical-5703b1398e29de5f2a9841ac6c5cdb440a08a34d601a770b0060215e1374c0f2"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / d274da873da1 / 3

<a id="canonical-e4f0288e8e23ccf06237a87f794787d8729f81d7e8553d48b619dc9c96db0761"></a>

<a id="canonical-043d59b73b46c074d742ee4d63ac54320a411d73fa603871b87a7f002c749d68"></a>

## name property — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / d274da873da1 / 4

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

<a id="canonical-88d7da459fd02e628de079dbbd8a3a8c339849ef9588dd86423ed4791c5f81f5"></a>

<a id="canonical-e9e16a45025e89dbd2158afe617a8699239be8b65a7039703e703418e6b29767"></a>

## namespace property — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / d274da873da1 / 5

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

<a id="canonical-c3de1960b0783a83fe050911a9fe23627b8424020f949b1bb1a72302585bc9a4"></a>

<a id="canonical-f36cccc745beef40ecf7a217776281935df8970288b3387726aa26f1dc2506d2"></a>

## tenant property — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / d274da873da1 / 6

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

<a id="canonical-7ce1594f5ca73c093cc264ad58ebd2cf6d7c78fb333105146323a7f94afacb89"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / d274da873da1 / 7

- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b6891065c1ae69c18d3812fb6df3d1655605dc43ec0b19d4987531e913590704"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-42258460c1c3af00372d1eadd271ec863a6d607271ded7b9b3fc9256846ceea6"></a>

## https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled / 1d14770b6e6b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5)
- https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled

<a id="canonical-89ba0b4a6e235d86d83f2a0edf06bfb939a5a357375bec8463d56a5c4e123a01"></a>

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

<a id="canonical-8894d12674c4fab3ed59e307866d164961611cb94f73e5f6263d4fc78890d730"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled / 1d14770b6e6b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9727ab94638d1508bbc3a6712163efa1563a0e216d1a06ae50d2bb5d5b6863e0"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled / 1d14770b6e6b / 4

- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1d282bba2f4940a95245f6249cb212f52f671211241b8c286667dd488b1e5a62"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ec9b14edf6b59a2df8358bedeb544bd9ec9adad8154229c70ce30f3aea0697e"></a>

## https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options / ed7d9a23ed37 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-91bf110360a4a1a30dc035083a949e24838734874a8c984c12a9376c2eecf595)
- [https.tls_cert_options](resources--cdn_loadbalancer--reference--group-010.md#canonical-3f932b92dec7147b65fff04d103540b16d238d7b87e51fd799ef8ad7d33cda9e)
- [https.tls_cert_options.tls_inline_params](resources--cdn_loadbalancer--reference--group-011.md#canonical-057144dabb883c47ae08cfc64b0e66ed7985be130eacf860907c9e83b0725d50)
- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5)
- https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options

<a id="canonical-ee311f62b5a5c25f4cdbc71eb03866df7c4bbf735d2e212b495c9f478c040d95"></a>

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

<a id="canonical-4743982dd04e7d8487fd77fc223b32c75c3fc28bb2d60d1ee7427868089eadec"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options / ed7d9a23ed37 / 3

<a id="canonical-9260024bdf9fc1dd32f071d1241c4ff0282129c7472e431193de5d2cc5d7d757"></a>

<a id="canonical-06003f2a08e5f7d90bb5c1a9c0c6a5072369f6c93569e036d9497f28688a5cf2"></a>

## xfcc_header_elements property — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options / ed7d9a23ed37 / 4

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

<a id="canonical-3dffd71fa9438f30e379d80f4c4c1501ef55db3117a11f99c05adb51e3811466"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options / ed7d9a23ed37 / 5

- [https.tls_cert_options.tls_inline_params.use_mtls](resources--cdn_loadbalancer--reference--group-011.md#canonical-e696626ff7f36d32d7a92a827f08e29d95cf9dfcec758da0b0b21506931ce2a5)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-bfe4be2d539814187ee6209f6c1e2551e3a3a993ebb2b6595877043d81965a7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da1f53801dc26b069cf4fb23fead723c990c50e48eeffbdbfe429c318a9868f2"></a>

## https_auto_cert — https_auto_cert / 2ff7fc8c344e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- https_auto_cert

<a id="canonical-4505e0388e96d2b6532ac846a2fb56512c8f4078a42d43348548aac21f6c9c37"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTPS CDN distribution with bring your own certificates.

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
https_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-78bcac7a35a33e2cbcc05f04788b067556e43cca4e0d8d8c9bef0fd46b269489"></a>

## Direct properties — https_auto_cert / 2ff7fc8c344e / 3

<a id="canonical-db4706d3180c5028a7566671c219760cdbc69a289e6206123fdfe23be73d13e5"></a>

<a id="canonical-bcacd841e2539a7d37a2fda32f25998fbf645a52d503e2cf8a7ff95cee8880ae"></a>

## add_hsts property — https_auto_cert / 2ff7fc8c344e / 4

Type: `"bool"`. Optional.

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

<a id="canonical-f66dacb51f8498327069807b8e383d95bf31dc153159692b70384367f08c7ea4"></a>

<a id="canonical-b3e0834decb16e5c74a7aa16270b787e6eb5ca2146e7b760ad687e4ca0103e8e"></a>

## http_redirect property — https_auto_cert / 2ff7fc8c344e / 5

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

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

- [tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-fc5b40e9a4b63e5f14b98ddd397b5dba0703b62c6d733bf30c54742c0ed68ed4): complete subsection reference.

<a id="canonical-2b7a7f96edbd30ad5c962bd197f3ed389ffd28c6f5c66b9f605cea65a81e3cf5"></a>

## Next pages — https_auto_cert / 2ff7fc8c344e / 6

- [https_auto_cert.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-fc5b40e9a4b63e5f14b98ddd397b5dba0703b62c6d733bf30c54742c0ed68ed4)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-fc5b40e9a4b63e5f14b98ddd397b5dba0703b62c6d733bf30c54742c0ed68ed4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2eacc987fa842732262d2e914b7614486967f04c3381c9b9d20d2b5985a802c3"></a>

## https_auto_cert.tls_config — https_auto_cert.tls_config / 5e3c52d87083 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https_auto_cert](resources--cdn_loadbalancer--reference--group-011.md#canonical-bfe4be2d539814187ee6209f6c1e2551e3a3a993ebb2b6595877043d81965a7b)
- https_auto_cert.tls_config

<a id="canonical-c9cd18afce725e8f098b172dbf3ccf68243cca6db65ab31480eeefe77a56dbfb"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("tls_11_plus",
    "tls_12_plus")}
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
  "x-ves-oneof-field-choice": "[\"tls_11_plus\",\"tls_12_plus\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-31317013ca3868de76052ba851d934c7bbc290caa67d3e62137ff168e39d64ca"></a>

## Direct properties — https_auto_cert.tls_config / 5e3c52d87083 / 3

- [tls_11_plus](resources--cdn_loadbalancer--reference--group-011.md#canonical-ee09216978978c37096208397f880954ae84aa68d26dfa99b3ac7e830b7a0e4d): complete subsection reference.

- [tls_12_plus](resources--cdn_loadbalancer--reference--group-011.md#canonical-53e99e13921408ee2fdea12403942a6cb8cd9e3d671ccb676ecc98a90ae1059e): complete subsection reference.

<a id="canonical-956f5a3ef6bda3a1de3e9eec77ae500289dce4014e9ae4acccc054cd3b8af5fa"></a>

## Next pages — https_auto_cert.tls_config / 5e3c52d87083 / 4

- [https_auto_cert.tls_config.tls_11_plus](resources--cdn_loadbalancer--reference--group-011.md#canonical-ee09216978978c37096208397f880954ae84aa68d26dfa99b3ac7e830b7a0e4d)
- [https_auto_cert.tls_config.tls_12_plus](resources--cdn_loadbalancer--reference--group-011.md#canonical-53e99e13921408ee2fdea12403942a6cb8cd9e3d671ccb676ecc98a90ae1059e)
- [https_auto_cert](resources--cdn_loadbalancer--reference--group-011.md#canonical-bfe4be2d539814187ee6209f6c1e2551e3a3a993ebb2b6595877043d81965a7b)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ee09216978978c37096208397f880954ae84aa68d26dfa99b3ac7e830b7a0e4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97dca58a35064e3ee1c43bd3689c53d7fb67af4dbca2a9ef8da1984a4820cfc4"></a>

## https_auto_cert.tls_config.tls_11_plus — https_auto_cert.tls_config.tls_11_plus / b7f7383220a8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https_auto_cert](resources--cdn_loadbalancer--reference--group-011.md#canonical-bfe4be2d539814187ee6209f6c1e2551e3a3a993ebb2b6595877043d81965a7b)
- [https_auto_cert.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-fc5b40e9a4b63e5f14b98ddd397b5dba0703b62c6d733bf30c54742c0ed68ed4)
- https_auto_cert.tls_config.tls_11_plus

<a id="canonical-4e477e29463872be6f2324138f4618ccda8547ba6b2a6a12f91fca5282efafa9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for tls 11 plus.

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
tls_11_plus = {}
```

<a id="canonical-d51f59e24fa548f22a7ec6bc867eabe85aef51bb1232923221230f6c4d7cd781"></a>

## Direct properties — https_auto_cert.tls_config.tls_11_plus / b7f7383220a8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9efe7334aeb44a1c8102e945a3c96fd0412a765a18a833f720a8265b5bce35d2"></a>

## Next pages — https_auto_cert.tls_config.tls_11_plus / b7f7383220a8 / 4

- [https_auto_cert.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-fc5b40e9a4b63e5f14b98ddd397b5dba0703b62c6d733bf30c54742c0ed68ed4)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-53e99e13921408ee2fdea12403942a6cb8cd9e3d671ccb676ecc98a90ae1059e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0bb09e777767390e120182bc285504d925f50ab4e5164276d5959b5e9143377"></a>

## https_auto_cert.tls_config.tls_12_plus — https_auto_cert.tls_config.tls_12_plus / 348419248fad / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [https_auto_cert](resources--cdn_loadbalancer--reference--group-011.md#canonical-bfe4be2d539814187ee6209f6c1e2551e3a3a993ebb2b6595877043d81965a7b)
- [https_auto_cert.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-fc5b40e9a4b63e5f14b98ddd397b5dba0703b62c6d733bf30c54742c0ed68ed4)
- https_auto_cert.tls_config.tls_12_plus

<a id="canonical-2a0a2d9f9bcda77c405e8b090715d0b44dfa2806c0261b797f920e6ac64c64b4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for tls 12 plus.

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
tls_12_plus = {}
```

<a id="canonical-39789ac15f177eb11ec52141562be728a31b5fd2c572122f76880038a3fa00f5"></a>

## Direct properties — https_auto_cert.tls_config.tls_12_plus / 348419248fad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-10ca81f07232251e95edf6a93a024d25d0832336dbd8a502829dff439c86e3b4"></a>

## Next pages — https_auto_cert.tls_config.tls_12_plus / 348419248fad / 4

- [https_auto_cert.tls_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-fc5b40e9a4b63e5f14b98ddd397b5dba0703b62c6d733bf30c54742c0ed68ed4)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-076789b43998a8bee94a64569f821cf1ab5cde4dac942c574ba859bdb49943b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a3c2268499f52a649babcf17668f207783f9b45843be1387cc73c4efe898bf7"></a>

## js_challenge — js_challenge / f39d9c8e95de / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- js_challenge

<a id="canonical-52c4ddba5ad61b31899f4590c585ce96cd8ae7ff6c559dc9b64266a1842ad358"></a>

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

<a id="canonical-25ed83c8f4bffc93eb589e38dfcaf79a6ecd2eeaf4c2980c49378a171861eb6d"></a>

## Direct properties — js_challenge / f39d9c8e95de / 3

<a id="canonical-b43e0053144df2a79c37f41b7135dafa1f15ec77e401ba5b9fbebf07c2b3d3f5"></a>

<a id="canonical-9a497bd4a0904a34f488608a061d44e6405d2183695da83bc97c352cd8e2cc95"></a>

## cookie_expiry property — js_challenge / f39d9c8e95de / 4

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

<a id="canonical-493f4dc4e4b39f60df43ae98a18fcb9fd636b4514508ce45815486280b430068"></a>

<a id="canonical-7d803f4eda91aec3c2cf8b0d816ba6f4ea110482ba1de41436cf3a506ece300e"></a>

## custom_page property — js_challenge / f39d9c8e95de / 5

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

<a id="canonical-8ae423cf6a96828dbadade41f7423fae66f7956b0b9bbfe7896a44682304efa2"></a>

<a id="canonical-896b2defb66fba805bfe4f8267ff7ed3095af5c5980c1eb2119241666a2dd3ea"></a>

## js_script_delay property — js_challenge / f39d9c8e95de / 6

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

<a id="canonical-3dd096c40d2895c4516d94206f358f410358fd29caf2ff155e981fbba609931f"></a>

## Next pages — js_challenge / f39d9c8e95de / 7

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-99669395019f8ebbf736b8edbc267b6c0af1cf1b135cdd61d645b04e1d883b95"></a>

## jwt_validation — jwt_validation / 08ec768111b6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- jwt_validation

<a id="canonical-c4018fc3d4be03f76944a8d43df65ac53b521254a28de6bbd57e2f28256753d2"></a>

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

<a id="canonical-ecc4157943a84a5a540be759aee8059935b573fd0a5d2fe8d2cbd0d7a2f957ce"></a>

## Direct properties — jwt_validation / 08ec768111b6 / 3

- [action](resources--cdn_loadbalancer--reference--group-011.md#canonical-2bb23a9719b0c485e9256fb2317dcd44bcfacbd0b4e540406f7cb92b492264ca): complete subsection reference.

- [authorization_server](resources--cdn_loadbalancer--reference--group-011.md#canonical-14ad4345127a3d24079ee0e630962176c7bdcea16370a6f74ad055b0d7733ad9): complete subsection reference.

- [jwks_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-c9727f57e69ee48883057a567adff2a5a27d442fa5892649bc8d29496c425b64): complete subsection reference.

- [mandatory_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-b74def61729f28dbf03bc938fe395d07b8e0353898b1e2e9b6737a336bbc4bf0): complete subsection reference.

- [reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d): complete subsection reference.

- [target](resources--cdn_loadbalancer--reference--group-011.md#canonical-989ffff9066adf7df2bdad6479c0382b8c57e0cb606cd07e4fb19b5e07f9fd7b): complete subsection reference.

- [token_location](resources--cdn_loadbalancer--reference--group-011.md#canonical-b203809ffe9d8acf4752853f2fe13f88799f51ab1c589aaa0bc8e392624c98c2): complete subsection reference.

<a id="canonical-cdca0f69e80e784517a9352d91fc160b5e492bf189696116da9740526d21dbfd"></a>

## Next pages — jwt_validation / 08ec768111b6 / 4

- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-011.md#canonical-2bb23a9719b0c485e9256fb2317dcd44bcfacbd0b4e540406f7cb92b492264ca)
- [jwt_validation.authorization_server](resources--cdn_loadbalancer--reference--group-011.md#canonical-14ad4345127a3d24079ee0e630962176c7bdcea16370a6f74ad055b0d7733ad9)
- [jwt_validation.jwks_config](resources--cdn_loadbalancer--reference--group-011.md#canonical-c9727f57e69ee48883057a567adff2a5a27d442fa5892649bc8d29496c425b64)
- [jwt_validation.mandatory_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-b74def61729f28dbf03bc938fe395d07b8e0353898b1e2e9b6737a336bbc4bf0)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d)
- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-011.md#canonical-989ffff9066adf7df2bdad6479c0382b8c57e0cb606cd07e4fb19b5e07f9fd7b)
- [jwt_validation.token_location](resources--cdn_loadbalancer--reference--group-011.md#canonical-b203809ffe9d8acf4752853f2fe13f88799f51ab1c589aaa0bc8e392624c98c2)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2bb23a9719b0c485e9256fb2317dcd44bcfacbd0b4e540406f7cb92b492264ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a952e1fb45c5ed0dd348bead338b6458d276fc14e67887bea9fabe656af8a0b2"></a>

## jwt_validation.action — jwt_validation.action / 8c68e7cd58a2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- jwt_validation.action

<a id="canonical-7b60717fa8853cbfe4ba791719282c78239f68599336ca14ff38c095a1b6f95c"></a>

Type: `"object"`. single nested block, Optional.

Action

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "report")}
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
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-29019b17781f62a9f561d1fc21c31650fbae37c1e252c844056287bfbc0f7f41"></a>

## Direct properties — jwt_validation.action / 8c68e7cd58a2 / 3

- [block](resources--cdn_loadbalancer--reference--group-011.md#canonical-a790f59dfb5642fc8042d4548aa80ef9514b5fd6ad07795525ad7f5451797a0b): complete subsection reference.

- [report](resources--cdn_loadbalancer--reference--group-011.md#canonical-615a62888fac3b28f1b092bfcd9d709777672fefd78d95b618cd64a69bbebadc): complete subsection reference.

<a id="canonical-74238c113be60fe8c384d0751383960df5a50100447a631004dd32808311756d"></a>

## Next pages — jwt_validation.action / 8c68e7cd58a2 / 4

- [jwt_validation.action.block](resources--cdn_loadbalancer--reference--group-011.md#canonical-a790f59dfb5642fc8042d4548aa80ef9514b5fd6ad07795525ad7f5451797a0b)
- [jwt_validation.action.report](resources--cdn_loadbalancer--reference--group-011.md#canonical-615a62888fac3b28f1b092bfcd9d709777672fefd78d95b618cd64a69bbebadc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-a790f59dfb5642fc8042d4548aa80ef9514b5fd6ad07795525ad7f5451797a0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4270d5e98975d4e3fd157f5be0e0208e5c52fea8e6cea398d124291bf667f43e"></a>

## jwt_validation.action.block — jwt_validation.action.block / 648317197ec9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-011.md#canonical-2bb23a9719b0c485e9256fb2317dcd44bcfacbd0b4e540406f7cb92b492264ca)
- jwt_validation.action.block

<a id="canonical-709c6715b20910266da2dbae702eb7bdc58426844a823facebe648755df2765b"></a>

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
block = {}
```

<a id="canonical-23b32e03aa1ede2807bb669eb206b4bfa08820298220d99c2b0b72ec2473e9e7"></a>

## Direct properties — jwt_validation.action.block / 648317197ec9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-77919d077ba6dafba12a0d87b3f63d6dba38e4fcb8232849f73f1f0c7552983e"></a>

## Next pages — jwt_validation.action.block / 648317197ec9 / 4

- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-011.md#canonical-2bb23a9719b0c485e9256fb2317dcd44bcfacbd0b4e540406f7cb92b492264ca)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-615a62888fac3b28f1b092bfcd9d709777672fefd78d95b618cd64a69bbebadc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1453c6ebb768479498a700ad83da7f422867b6eb1dda6b124d3285156377dd28"></a>

## jwt_validation.action.report — jwt_validation.action.report / ea541e2d4bb8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-011.md#canonical-2bb23a9719b0c485e9256fb2317dcd44bcfacbd0b4e540406f7cb92b492264ca)
- jwt_validation.action.report

<a id="canonical-53e278c10dd088193d043b366c1e5b53f70d7f3f031ca178fdb73efaa9f9594e"></a>

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
report = {}
```

<a id="canonical-c45442b56a4d1fa7bb53bcecc130c03bd1c3faf19bda319aca9de17eaab78529"></a>

## Direct properties — jwt_validation.action.report / ea541e2d4bb8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0066afc9540680d9cec51ef355158d170d24572ac967c39ea3abe85683bf6589"></a>

## Next pages — jwt_validation.action.report / ea541e2d4bb8 / 4

- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-011.md#canonical-2bb23a9719b0c485e9256fb2317dcd44bcfacbd0b4e540406f7cb92b492264ca)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-14ad4345127a3d24079ee0e630962176c7bdcea16370a6f74ad055b0d7733ad9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89c173d669f889aa7426f2ca6b0cb4f35dceb49327180fe41efa6026a87aef59"></a>

## jwt_validation.authorization_server — jwt_validation.authorization_server / 0e5939f233fb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- jwt_validation.authorization_server

<a id="canonical-7146b9b1cde6c88dc488f1353e4ed06d0c75a4f48a0b8138671587b6f05eb3c5"></a>

Type: `"object"`. single nested block, Optional.

Reference to Authorization Server object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("authorization_servers")}
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
authorization_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-926cb033705a8488d07a6f375a2a88b8467270d98d56c36dd60e855589dc5ba8"></a>

## Direct properties — jwt_validation.authorization_server / 0e5939f233fb / 3

- [authorization_servers](resources--cdn_loadbalancer--reference--group-011.md#canonical-9463f0c8afb434ab02e90bd032a8adfdabfba348ef5da9aa08b4666522010361): complete subsection reference.

<a id="canonical-9b8cb35df1cc3fd86880401ab413042ed50a2cfdd8e6b76d8acd2ba2d8e17dc1"></a>

## Next pages — jwt_validation.authorization_server / 0e5939f233fb / 4

- [jwt_validation.authorization_server.authorization_servers](resources--cdn_loadbalancer--reference--group-011.md#canonical-9463f0c8afb434ab02e90bd032a8adfdabfba348ef5da9aa08b4666522010361)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-9463f0c8afb434ab02e90bd032a8adfdabfba348ef5da9aa08b4666522010361"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a674f4420ca7168e5e070d687c7d01dcc5bed94597dca21eca87701b75eafde3"></a>

## jwt_validation.authorization_server.authorization_servers — jwt_validation.authorization_server.authorization_servers / 1fff502b17a7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.authorization_server](resources--cdn_loadbalancer--reference--group-011.md#canonical-14ad4345127a3d24079ee0e630962176c7bdcea16370a6f74ad055b0d7733ad9)
- jwt_validation.authorization_server.authorization_servers

<a id="canonical-b371f284a89842bf957717ca818cd447c1a0c1fb9b24880d4ab8822b15b3e476"></a>

Type: `"object"`. list nested block, Optional.

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Upstream description:

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

Terraform syntax:

```terraform
authorization_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-ecfbe00589b04b87e4631365e2497b5859d9f8a745370ca07a280c965772084b"></a>

## Direct properties — jwt_validation.authorization_server.authorization_servers / 1fff502b17a7 / 3

<a id="canonical-8cf865e5852e102a332a0b5184888df6b38f8a31f8593537b809e7c6b4e71aca"></a>

<a id="canonical-b0bf7d023134df170fa8def8db47a02a63e2b198e89abee6f981bcc6f82049e1"></a>

## name property — jwt_validation.authorization_server.authorization_servers / 1fff502b17a7 / 4

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

<a id="canonical-ffb6a7fd10ad664094187f1c52b13c35d0b46dd34bbfed3ccdb721ae811ec5c9"></a>

<a id="canonical-c72256aa40fa11a931b0adb1412f6426ecc53156aff440fdf69f5812d57ad375"></a>

## namespace property — jwt_validation.authorization_server.authorization_servers / 1fff502b17a7 / 5

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

<a id="canonical-f2791338f1dbf541cac2ec563d4fe672456cb25030932f86e23ca680353989e6"></a>

<a id="canonical-134ec201ba678fd3fedba1c9acbf81723cdd123e0bcb0d90f457cc5460b766f9"></a>

## tenant property — jwt_validation.authorization_server.authorization_servers / 1fff502b17a7 / 6

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

<a id="canonical-6d0c4d97e36061495808893dc88b5ba6a36784a5b2efa482bb1f8aa7c448a62a"></a>

## Next pages — jwt_validation.authorization_server.authorization_servers / 1fff502b17a7 / 7

- [jwt_validation.authorization_server](resources--cdn_loadbalancer--reference--group-011.md#canonical-14ad4345127a3d24079ee0e630962176c7bdcea16370a6f74ad055b0d7733ad9)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c9727f57e69ee48883057a567adff2a5a27d442fa5892649bc8d29496c425b64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88f0b12bf8a2b035eb136b919415fffcadcc27f968fe8aa831cee00574e0bba6"></a>

## jwt_validation.jwks_config — jwt_validation.jwks_config / 1c656345c710 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- jwt_validation.jwks_config

<a id="canonical-28cbde7749cfac9ed5ab51f6272d6c0c24b65164787068c354a4c46e27be92a6"></a>

Type: `"object"`. single nested block, Optional.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

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
jwks_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-70ed9234c5c90f75a1b860cfe381eca098fd4de2ed043411129e19118f398af5"></a>

## Direct properties — jwt_validation.jwks_config / 1c656345c710 / 3

<a id="canonical-4025ed6cdac5bfa80aa4ea8643fd2471004d67d251740113729b1f11c4e32d7a"></a>

<a id="canonical-1055367ef9b13e65a6ee3fa55e91119bdeb2c02d4f5a07ba579d5156c2acc2bf"></a>

## cleartext property — jwt_validation.jwks_config / 1c656345c710 / 4

Type: `"string"`. Optional.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

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

<a id="canonical-668c64deb6e116c18cc0ea481bb5a0b99fda537ff0b274fbc2ca56715724d618"></a>

## Next pages — jwt_validation.jwks_config / 1c656345c710 / 5

- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b74def61729f28dbf03bc938fe395d07b8e0353898b1e2e9b6737a336bbc4bf0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffd7a4f4b682866d08bcf06d407b74c3373fce1ec668e74ea950e98d5da4a7ad"></a>

## jwt_validation.mandatory_claims — jwt_validation.mandatory_claims / 4a1fe0ab49c1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- jwt_validation.mandatory_claims

<a id="canonical-bf8fe1cedf1d79068b709d20a61c936a4625209e6e2fe3dd66b9aa6698e1131f"></a>

Type: `"object"`. single nested block, Optional.

Configurable Validation of mandatory Claims.

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
mandatory_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-7db9c16e0b638cba791da4e2922ac06391d6823fca3b4db27d2c65ba9972e522"></a>

## Direct properties — jwt_validation.mandatory_claims / 4a1fe0ab49c1 / 3

<a id="canonical-d2bd13469406a721445aa753525a309c163a0e17a89486ab7c0aa7a353dd1a3c"></a>

<a id="canonical-c0ce2471afc2d973823fdd7c57064b4536106bae56014b760499ee281f3ef11e"></a>

## claim_names property — jwt_validation.mandatory_claims / 4a1fe0ab49c1 / 4

Type: `["list", "string"]`. Optional.

Claim Names. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-e19044a5f8409bfa9313ef531dee9131541a6de4c60285ddff8b6aa64889781c"></a>

## Next pages — jwt_validation.mandatory_claims / 4a1fe0ab49c1 / 5

- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-176b5e1470c3be5abedd8a8763a0d900ac2974490e4b1d512e431cc285fc3926"></a>

## jwt_validation.reserved_claims — jwt_validation.reserved_claims / e8360c200014 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- jwt_validation.reserved_claims

<a id="canonical-e65a0e0dbca717dd5c411ab9fc5d86634c4d488f4ca1a34367d0bbc17763ab24"></a>

Type: `"object"`. single nested block, Optional.

Configurable Validation of reserved Claims.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("audience",
    "audience_disable"),
  validators.ConflictingObjectAttributes("issuer",
    "issuer_disable"),
  validators.ConflictingObjectAttributes("validate_period_disable",
    "validate_period_enable")}
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
  "x-ves-oneof-field-audience_validation": "[\"audience\",\"audience_disable\"]",
  "x-ves-oneof-field-issuer_validation": "[\"issuer\",\"issuer_disable\"]",
  "x-ves-oneof-field-validate_period": "[\"validate_period_disable\",\"validate_period_enable\"]"
}
```

Terraform syntax:

```terraform
reserved_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-fcd68948f800c2c03f59908126aa1c0eb7a966b3e7097aa2872b6fbb94cfabfa"></a>

## Direct properties — jwt_validation.reserved_claims / e8360c200014 / 3

- [audience](resources--cdn_loadbalancer--reference--group-011.md#canonical-c801284174d2c252a6a8c8cff558eba4f9260351d2fa01b4f747e97d50f39cd0): complete subsection reference.

- [audience_disable](resources--cdn_loadbalancer--reference--group-011.md#canonical-677c934841d31bf0218bbba3347100bff5a1bada429777fea319141b9d7c667b): complete subsection reference.

<a id="canonical-a94e60377fd61eb27ea5bb2c5a2b60931cc3b56d92c30f0e90df778df0a611f3"></a>

<a id="canonical-def9fe7ea664caed383fd37f75ea431fa18c020b3d4384e1739975c507d7af7c"></a>

## issuer property — jwt_validation.reserved_claims / e8360c200014 / 4

Type: `"string"`. Optional.

Exact Match. Exclusive with \[issuer\_disable\]

Upstream description:

Exclusive with \[issuer\_disable\]

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

- [issuer_disable](resources--cdn_loadbalancer--reference--group-011.md#canonical-392f57a5d70e6b8a089aca1359e0acff91c2c0c6c5e5da4b05dfdf869e7fd76d): complete subsection reference.

- [validate_period_disable](resources--cdn_loadbalancer--reference--group-011.md#canonical-55bb6cea0e16aeb16dfa12f0f0c625255d9e5727862cfbdb43642f7bb4502f5d): complete subsection reference.

- [validate_period_enable](resources--cdn_loadbalancer--reference--group-011.md#canonical-1614580c420894c90b6bb74ffdcdf8002f9b56ae5c1df39030535d2b5b933992): complete subsection reference.

<a id="canonical-73b99d080f6a28b6cac0da0e47360a5c43be0d5cee83786508b7ccd1482c4762"></a>

## Next pages — jwt_validation.reserved_claims / e8360c200014 / 5

- [jwt_validation.reserved_claims.audience](resources--cdn_loadbalancer--reference--group-011.md#canonical-c801284174d2c252a6a8c8cff558eba4f9260351d2fa01b4f747e97d50f39cd0)
- [jwt_validation.reserved_claims.audience_disable](resources--cdn_loadbalancer--reference--group-011.md#canonical-677c934841d31bf0218bbba3347100bff5a1bada429777fea319141b9d7c667b)
- [jwt_validation.reserved_claims.issuer_disable](resources--cdn_loadbalancer--reference--group-011.md#canonical-392f57a5d70e6b8a089aca1359e0acff91c2c0c6c5e5da4b05dfdf869e7fd76d)
- [jwt_validation.reserved_claims.validate_period_disable](resources--cdn_loadbalancer--reference--group-011.md#canonical-55bb6cea0e16aeb16dfa12f0f0c625255d9e5727862cfbdb43642f7bb4502f5d)
- [jwt_validation.reserved_claims.validate_period_enable](resources--cdn_loadbalancer--reference--group-011.md#canonical-1614580c420894c90b6bb74ffdcdf8002f9b56ae5c1df39030535d2b5b933992)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-c801284174d2c252a6a8c8cff558eba4f9260351d2fa01b4f747e97d50f39cd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56649fcd9ec738f16260f346748c233716f8ff09156c0112a7ec0109f9b619d5"></a>

## jwt_validation.reserved_claims.audience — jwt_validation.reserved_claims.audience / 4ace2dc62480 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d)
- jwt_validation.reserved_claims.audience

<a id="canonical-30c5eeadac8ea76faaec02e155134dfd4cf5220b6dcaacbca4c23e8d77e44152"></a>

Type: `"object"`. single nested block, Optional.

Audiences

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("audiences")}
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
audience {
  # Configure direct properties listed below.
}
```

<a id="canonical-3f55e87e0cc2c8794d6fca7d585327a2f3c4c8bde353c2f79ffa72547bea1c7d"></a>

## Direct properties — jwt_validation.reserved_claims.audience / 4ace2dc62480 / 3

<a id="canonical-60c7b97e3425d23b7e963d926be61ea05e36e8dd243727709dd079282bd52bba"></a>

<a id="canonical-c08a3df576285a574e6f0c2c828bf4f0453a5a82fc12587e8a70b0c4d6118ebc"></a>

## audiences property — jwt_validation.reserved_claims.audience / 4ace2dc62480 / 4

Type: `["list", "string"]`. Optional.

Values. Configuration parameter for audiences

Upstream description:

Configuration parameter for audiences

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-60ad9970d2d60a13844b8e041b4fbaf5c37485fc519cce0c4289b838deb7b8f1"></a>

## Next pages — jwt_validation.reserved_claims.audience / 4ace2dc62480 / 5

- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-677c934841d31bf0218bbba3347100bff5a1bada429777fea319141b9d7c667b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4d54ff507dd425df10d3bf555fb40d62ee2a9a1087ca639bf202f4a5036edc6"></a>

## jwt_validation.reserved_claims.audience_disable — jwt_validation.reserved_claims.audience_disable / 5a94967e5b6b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d)
- jwt_validation.reserved_claims.audience_disable

<a id="canonical-879c984f5c440208e530d8e2efc4e021a5eb7e5f11e90f708cb0e350d26152f0"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for audience disable.

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
audience_disable = {}
```

<a id="canonical-8a9ec46f4f0f3c91d3abf6b8d4f7d9a6d382ab5d6e48afa543627eda3a6fc9a6"></a>

## Direct properties — jwt_validation.reserved_claims.audience_disable / 5a94967e5b6b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f1e32fc714e985ee6aeb8ca5a1748d97dbafd0e98d5d6dfcfddc39b6f60c2363"></a>

## Next pages — jwt_validation.reserved_claims.audience_disable / 5a94967e5b6b / 4

- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-392f57a5d70e6b8a089aca1359e0acff91c2c0c6c5e5da4b05dfdf869e7fd76d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49731ec25d82a20f4dbee715a6250a2dbe3c514030d463854240546fa1e22d08"></a>

## jwt_validation.reserved_claims.issuer_disable — jwt_validation.reserved_claims.issuer_disable / 30ca71367ea5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d)
- jwt_validation.reserved_claims.issuer_disable

<a id="canonical-5ac34e4194a2ab3fa4865f43ea4b940dad95a019743eddbe42ad7982afba2993"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for issuer disable.

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
issuer_disable = {}
```

<a id="canonical-925b686851cd01035de5c1c95b9d18ad1c0cb437feae9dea8cd55e0f9de86a03"></a>

## Direct properties — jwt_validation.reserved_claims.issuer_disable / 30ca71367ea5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b259e9fe4ec81fd32e6486ba5b1b9c4d5be6d6a521f86fa4d81b1d2a1490942b"></a>

## Next pages — jwt_validation.reserved_claims.issuer_disable / 30ca71367ea5 / 4

- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-55bb6cea0e16aeb16dfa12f0f0c625255d9e5727862cfbdb43642f7bb4502f5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9e13bb6863b0f062861f2ada84a6ec823f0d218e337c9f51bc3b4bff3872d6fa"></a>

## jwt_validation.reserved_claims.validate_period_disable — jwt_validation.reserved_claims.validate_period_disable / 28588bdf71ab / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d)
- jwt_validation.reserved_claims.validate_period_disable

<a id="canonical-1bdc94ebf96134bdf19d8894e40cf954d50e9669775fea595d958f1f96f5d765"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period disable.

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
validate_period_disable = {}
```

<a id="canonical-5422fb0d36f0e627a4970e81f458320438f528f7f2eb0b2edd5a8de902b2920b"></a>

## Direct properties — jwt_validation.reserved_claims.validate_period_disable / 28588bdf71ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b1731d958d780210fceee281bf94632976e397ddc65ba5c226d0454911ff419e"></a>

## Next pages — jwt_validation.reserved_claims.validate_period_disable / 28588bdf71ab / 4

- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-1614580c420894c90b6bb74ffdcdf8002f9b56ae5c1df39030535d2b5b933992"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3da6f327c09945dc02d2351a641a94a94cdb2213e953600ff75d0536fefef69"></a>

## jwt_validation.reserved_claims.validate_period_enable — jwt_validation.reserved_claims.validate_period_enable / 959dac8b79a7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d)
- jwt_validation.reserved_claims.validate_period_enable

<a id="canonical-d25948535db4848c9038c9e4710ad81ba462aeff25a5923424f175a0784be7f0"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period enable.

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
validate_period_enable = {}
```

<a id="canonical-fec35abc4297bd4a64018cddbade7b0634aab4d5cac9e99b2c64360eafb9a129"></a>

## Direct properties — jwt_validation.reserved_claims.validate_period_enable / 959dac8b79a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40cb7f884226c0ddacf13a172220d8028b62187e6d08ee0940a04aaa30ff69e9"></a>

## Next pages — jwt_validation.reserved_claims.validate_period_enable / 959dac8b79a7 / 4

- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-011.md#canonical-2972e0d1e152dbefd6d2eee7d49509eeff109ae9f5d0441fe0bc6c213672455d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-989ffff9066adf7df2bdad6479c0382b8c57e0cb606cd07e4fb19b5e07f9fd7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4af437e549fe9bf8630c45d723f98be748de531d311e66b9d88c9fa3ea4ae21e"></a>

## jwt_validation.target — jwt_validation.target / 6499ae71a913 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- jwt_validation.target

<a id="canonical-958ff7ed12a6496741d679f87d7f8946dae411c04185f7e3a4b1ec86bc330ac7"></a>

Type: `"object"`. single nested block, Optional.

Define endpoints for which JWT token validation will be performed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_endpoint",
    "api_groups"),
  validators.ConflictingObjectAttributes("all_endpoint",
    "base_paths"),
  validators.ConflictingObjectAttributes("api_groups",
    "base_paths")}
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
  "x-ves-oneof-field-target": "[\"all_endpoint\",\"api_groups\",\"base_paths\"]"
}
```

Terraform syntax:

```terraform
target {
  # Configure direct properties listed below.
}
```

<a id="canonical-c47ae7933242940720cc77d539ee54ef34351139b3c24db913c30f723a28d9dc"></a>

## Direct properties — jwt_validation.target / 6499ae71a913 / 3

- [all_endpoint](resources--cdn_loadbalancer--reference--group-011.md#canonical-0cf76b19e59acb49bbc5fc0eeb76a68b9c9f1a4124357574b057f0cef09b133c): complete subsection reference.

- [api_groups](resources--cdn_loadbalancer--reference--group-011.md#canonical-ddb14e69c22ce420076a9af39ff79fa57770626ac1cde74c73f2b41019100220): complete subsection reference.

- [base_paths](resources--cdn_loadbalancer--reference--group-011.md#canonical-e25d874938fc013e0e599aa8a332fe2ac9e84e6adbcb2d327467ac308b4e7763): complete subsection reference.

<a id="canonical-615be50897b99a3ba63ce5a67f88b4b2902415ec3660c0c2e6ca74cb2fbb5bc8"></a>

## Next pages — jwt_validation.target / 6499ae71a913 / 4

- [jwt_validation.target.all_endpoint](resources--cdn_loadbalancer--reference--group-011.md#canonical-0cf76b19e59acb49bbc5fc0eeb76a68b9c9f1a4124357574b057f0cef09b133c)
- [jwt_validation.target.api_groups](resources--cdn_loadbalancer--reference--group-011.md#canonical-ddb14e69c22ce420076a9af39ff79fa57770626ac1cde74c73f2b41019100220)
- [jwt_validation.target.base_paths](resources--cdn_loadbalancer--reference--group-011.md#canonical-e25d874938fc013e0e599aa8a332fe2ac9e84e6adbcb2d327467ac308b4e7763)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-0cf76b19e59acb49bbc5fc0eeb76a68b9c9f1a4124357574b057f0cef09b133c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc8f4d0a9152955152313fb917a8a00d7ce4478535d7cb2ee616969a7054f35c"></a>

## jwt_validation.target.all_endpoint — jwt_validation.target.all_endpoint / e9a536ca37f5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-011.md#canonical-989ffff9066adf7df2bdad6479c0382b8c57e0cb606cd07e4fb19b5e07f9fd7b)
- jwt_validation.target.all_endpoint

<a id="canonical-c2f29c744b4905006953be4d688bb3ec0637a1f8cf77db949637ccf227c810fa"></a>

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
all_endpoint = {}
```

<a id="canonical-ba17452ecc842bd7e435f25d4395379b34103dc8a94a711ca73d956c5df9af05"></a>

## Direct properties — jwt_validation.target.all_endpoint / e9a536ca37f5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b7c412fbc3bcdef28cb293a9d6e669936d37694831b85ab8d8e6b2b297141663"></a>

## Next pages — jwt_validation.target.all_endpoint / e9a536ca37f5 / 4

- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-011.md#canonical-989ffff9066adf7df2bdad6479c0382b8c57e0cb606cd07e4fb19b5e07f9fd7b)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-ddb14e69c22ce420076a9af39ff79fa57770626ac1cde74c73f2b41019100220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8ca0281c4e4f51a7ffa93af2db3ca0102190097cefc806919551cdf91076338"></a>

## jwt_validation.target.api_groups — jwt_validation.target.api_groups / 314e4af39ffe / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-011.md#canonical-989ffff9066adf7df2bdad6479c0382b8c57e0cb606cd07e4fb19b5e07f9fd7b)
- jwt_validation.target.api_groups

<a id="canonical-43fc32891be249e02d305cde8943387bb85d3a4050e16cec616c09e10e6d0ad6"></a>

Type: `"object"`. single nested block, Optional.

API Groups.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_groups")}
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
api_groups {
  # Configure direct properties listed below.
}
```

<a id="canonical-5e141267473bbc0722ab0dd42f65e50d3f8317cdcffb915ff73657cbed3f4788"></a>

## Direct properties — jwt_validation.target.api_groups / 314e4af39ffe / 3

<a id="canonical-85117db96ebb0a6ec5436adb5113ef271b96157dd47ee5b485b8e1b926baef7e"></a>

<a id="canonical-83a07c0c2c177442a254393a0e55c7363804672ec1c07d23317ebd5b2430ae2e"></a>

## api_groups property — jwt_validation.target.api_groups / 314e4af39ffe / 4

Type: `["list", "string"]`. Optional.

API Groups. Group or collection configuration

Upstream description:

Group or collection configuration

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-fbc878af9d339fd494727b841b69215722476c586adf92e75300ed3758a73011"></a>

## Next pages — jwt_validation.target.api_groups / 314e4af39ffe / 5

- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-011.md#canonical-989ffff9066adf7df2bdad6479c0382b8c57e0cb606cd07e4fb19b5e07f9fd7b)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-e25d874938fc013e0e599aa8a332fe2ac9e84e6adbcb2d327467ac308b4e7763"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8695a4f10b3c9f9db8e98e72ce11ed047d8a6a613520631b00f5a6dd8cff7e82"></a>

## jwt_validation.target.base_paths — jwt_validation.target.base_paths / 975ba4d9c60c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-011.md#canonical-989ffff9066adf7df2bdad6479c0382b8c57e0cb606cd07e4fb19b5e07f9fd7b)
- jwt_validation.target.base_paths

<a id="canonical-421afda3847b21e8d4a32adb12f11bdbc5d79b4506ecef4b6c47cdd54255c83e"></a>

Type: `"object"`. single nested block, Optional.

Base Paths.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("base_paths")}
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
base_paths {
  # Configure direct properties listed below.
}
```

<a id="canonical-4473df99ef04dda2d92a5ea2b857abfae9d583e7b807ac25fd4511cbca90418d"></a>

## Direct properties — jwt_validation.target.base_paths / 975ba4d9c60c / 3

<a id="canonical-694c331b52aeddab1c6ee245c9b64ac95392247a515f31f3965594fd4b425ac4"></a>

<a id="canonical-9bcb86f1c5a792bae9b8b33561c4284dd362a4d4d66622169296db620da246ca"></a>

## base_paths property — jwt_validation.target.base_paths / 975ba4d9c60c / 4

Type: `["list", "string"]`. Optional.

Prefix Values. File system or URL path

Upstream description:

File system or URL path

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9611b29b5f1e5508188c89cf0bfb3e4e44f897168530dfc763dac9160658ca15"></a>

## Next pages — jwt_validation.target.base_paths / 975ba4d9c60c / 5

- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-011.md#canonical-989ffff9066adf7df2bdad6479c0382b8c57e0cb606cd07e4fb19b5e07f9fd7b)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)

<a id="canonical-b203809ffe9d8acf4752853f2fe13f88799f51ab1c589aaa0bc8e392624c98c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c25f5537ac9cb34708da96bfa03688341ef9cc714f66ca249f03c059a6bd3131"></a>

## jwt_validation.token_location — jwt_validation.token_location / 9809b0becced / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-371adde2be63fed533ab2dce28fba2d800a088907dd21e5448db74ca97a1753b)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-467da3974497605aa4a2c22ffcea869e72329acd5aaa3f4197905ab521e9ebfc)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-ed941b0598c35b0b9845ee273333bbd7ddfe716ab488784720e0406aca406b8d)
- jwt_validation.token_location

<a id="canonical-ea683addf05dcb87491df94ea8a86bb72f79fb1cac5d8cd6a4ff28717ff7ce4b"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for token location.

Upstream description:

Location of JWT in HTTP request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-token_location": "[\"bearer_token\"]"
}
```

Terraform syntax:

```terraform
token_location {
  # Configure direct properties listed below.
}
```

<a id="canonical-6d9c5b2e4708454e268de010252beab202bef528d9d6ba6e34c833f036c8c408"></a>

## Direct properties — jwt_validation.token_location / 9809b0becced / 3

- [bearer_token](resources--cdn_loadbalancer--reference--group-012.md#canonical-b772b8cbe6dcb9788c6cee31a05ff241c10c084208939551659cfd5657f93b7c): complete subsection reference.
