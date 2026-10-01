---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-5017b918f3ef044ce3166390b28c26b8221afd6253d8752468f8345c0bc2a9c5"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling — https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling / fa44205c1419 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1)
- https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-8442248181211775dcdb08d822003b479570794e960f75ac45b5f7eda7cc0334"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-cb5272ddef8458f2662c10036525a6ec0755c4cf9672a6c9a43585679623ab88"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling / fa44205c1419 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-70a3a885b7a42b613050e8d959e93df9254a699ff08625b598e2a2428d1b705b"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling / fa44205c1419 / 4

- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3d56980d1415ea8c839ec57ed0d9ad43d643f19a2b2af2b7b85280ae717d758b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61e4ff013b5c8886dafe23808ca9594940395fe95f3fcd357ef04b0d50ca2c08"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.private_key — https.tls_cert_options.tls_inline_params.tls_certificates.private_key / 3863c636b56b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key

<a id="canonical-5d5092c27ba887b876e98cd608570006657c0b9f9d47481f31b89cee1f75f54d"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-36cbaba8b8c7a5aa1c70a5516324a92756f6b12becb6041581e85227fa26b8d0"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.private_key / 3863c636b56b / 3

- [blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-19a8c6332f47fee888f1615d71f824ffb3b93861739920d4527c876b3ace1f74): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-c112da46d6d0cd9ace522762d8f695dd45371494d1e9f19725ec5b9b58b944d5): complete subsection reference.

<a id="canonical-d843c840560a0a61d450694c55de2eab939457d21996bee0039c11f6563d2e3f"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.private_key / 3863c636b56b / 4

- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-19a8c6332f47fee888f1615d71f824ffb3b93861739920d4527c876b3ace1f74)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-c112da46d6d0cd9ace522762d8f695dd45371494d1e9f19725ec5b9b58b944d5)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-19a8c6332f47fee888f1615d71f824ffb3b93861739920d4527c876b3ace1f74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e3c944a2955e55faa4a8501fb7fc6751a63ac4f16e5b2ec1a5e9c36393d489c"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / bf1dc7bbb6cf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3d56980d1415ea8c839ec57ed0d9ad43d643f19a2b2af2b7b85280ae717d758b)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-8dd553dde77f8cc2d4209d74d0114acf4690c8dd913fcf3483cc800cf58e3517"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-8478089fe400009fb90bf24bfa249e47d1e485c761dced772d8f225e8444d8b1"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / bf1dc7bbb6cf / 3

<a id="canonical-44982ce4f46193a0b3e160be192d87c8bd92ab9ea97d2ade6bdd9978099c655c"></a>

<a id="canonical-3a04c0f5c4fd14f0b38a4c39e74969acf61739989d2c66d0258b1f0ab0f5e8d4"></a>

## decryption_provider property — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / bf1dc7bbb6cf / 4

Type: `"string"`. Computed.

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

<a id="canonical-b2e34716349b507589899fa98d8eb9afe0d521b7d5e78da7304dbcad6435d21f"></a>

<a id="canonical-96842aced6f92ae222c4e24734c0937f064d55c3302ad37212d9e42db23b421f"></a>

## location property — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / bf1dc7bbb6cf / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

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

<a id="canonical-2cc5543da4074be287cc78afe1dc1ee45592e3987fdac110a95c0861a707ba0f"></a>

<a id="canonical-f5c58daa7544930789e86a89745fd67f29b9dd139d2e2597658530619c916502"></a>

## store_provider property — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / bf1dc7bbb6cf / 6

Type: `"string"`. Computed.

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

<a id="canonical-8b8e0ed4adb47c13a6f55024e57bd127e15ee006f58ef49cb889f226087f1104"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.blindfold_ / bf1dc7bbb6cf / 7

- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3d56980d1415ea8c839ec57ed0d9ad43d643f19a2b2af2b7b85280ae717d758b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c112da46d6d0cd9ace522762d8f695dd45371494d1e9f19725ec5b9b58b944d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2508ea502c31f1a8dee660973fc6657da7d4708845927322f0350945a4c38421"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secr / d844e05e1bbc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3d56980d1415ea8c839ec57ed0d9ad43d643f19a2b2af2b7b85280ae717d758b)
- https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-18baccb3ec3e57e3221e21411709307f2e8d93bcae7f06236663dedf4818cb23"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-7c9a5279bd38d2ef8b8d3c1c46493593fe9be7cf7710f1a44db94b70d9b678b4"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secr / d844e05e1bbc / 3

<a id="canonical-1a7a558a16bc665266d9491abd02153fef43f0a9fb547c9fbcaaf45814b4f843"></a>

<a id="canonical-f706e83839ccddd9ece5a7e942845fca1545eeb6f2c66a4b3a4cabf838cf2c59"></a>

## provider_ref property — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secr / d844e05e1bbc / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-07079cf3af8650173967d9baad506669b0adf5b59bc0ed8d6909b92f938b648f"></a>

<a id="canonical-704f9724016b620019f03b28806aa8b588ef4837fc548331c969e3cfb190e415"></a>

## url property — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secr / d844e05e1bbc / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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

<a id="canonical-6ba40b90416ed4de4712a94f52b03d4eb504488677efc0a0d614aaed2dc426ec"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.private_key.clear_secr / d844e05e1bbc / 6

- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3d56980d1415ea8c839ec57ed0d9ad43d643f19a2b2af2b7b85280ae717d758b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-0a0ff49597b60275fbdfc2986691bcf1f247e11a50e307597042ffd797282635"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32c57373b60670608ca2357caf02fb78f72abd85463865561d9831fe2f2be33e"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults — https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults / 755ca315c3e4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1)
- https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults

<a id="canonical-cb3c8087ea3422be1e09283c995e11305eed025f667a9ff9585c4faca9751f58"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-66221996b21548d1b1065d353b03b5bd273f2f111e41c4c59d6bb2c46481550e"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults / 755ca315c3e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-198d85a55063bfa189f18aefd80ef17f9ba6978aa6fef969c723c2ce2a6bbf04"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults / 755ca315c3e4 / 4

- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-574182211b3e722bb2d0f6d4a1dedc2e117fcca6866a437156769d4326f782d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96b530145ac1625dcf4ecb977598a731089a345cb2791f4f26dc5597ae71a50e"></a>

## https.tls_cert_options.tls_inline_params.tls_config — https.tls_cert_options.tls_inline_params.tls_config / f211494fe23d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- https.tls_cert_options.tls_inline_params.tls_config

<a id="canonical-f10a494f7ebe7a7b0e7bb698c80146894221cca2faae51b5bd635c0eaaad6d51"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

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

<a id="canonical-fcba6e193911594b269f1765e18dd632f2607fb9e17a0fa10145834a1f76a6b7"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_config / f211494fe23d / 3

- [custom_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-072762cec23ef9c394eb6a5f7dc82ccb2b4bfab4c94a2b477887df5775ee9171): complete subsection reference.

- [default_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-5fecc5d862578277a0548f9accdb27e058b1f33d771160e46e412d1b57e1316f): complete subsection reference.

- [low_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-da2be12eccbcb9791c45c5174bfdb65cb90601ed8706d945e17d0720c0a6216e): complete subsection reference.

- [medium_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-570e51d3928542c0288e95f2606cb5cb89001eb79b212e57a81c4f2663919230): complete subsection reference.

<a id="canonical-d9624e54f95b723cbd12438303b3b00e61db0ab8de63e0e889736a3a8b206fcf"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_config / f211494fe23d / 4

- [https.tls_cert_options.tls_inline_params.tls_config.custom_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-072762cec23ef9c394eb6a5f7dc82ccb2b4bfab4c94a2b477887df5775ee9171)
- [https.tls_cert_options.tls_inline_params.tls_config.default_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-5fecc5d862578277a0548f9accdb27e058b1f33d771160e46e412d1b57e1316f)
- [https.tls_cert_options.tls_inline_params.tls_config.low_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-da2be12eccbcb9791c45c5174bfdb65cb90601ed8706d945e17d0720c0a6216e)
- [https.tls_cert_options.tls_inline_params.tls_config.medium_security](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-570e51d3928542c0288e95f2606cb5cb89001eb79b212e57a81c4f2663919230)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-072762cec23ef9c394eb6a5f7dc82ccb2b4bfab4c94a2b477887df5775ee9171"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b64696695bd73a0ea6d2be006990393b81ae3ed43278b4f4f4370bfdcfa2fea"></a>

## https.tls_cert_options.tls_inline_params.tls_config.custom_security — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 0974c0b58c60 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-574182211b3e722bb2d0f6d4a1dedc2e117fcca6866a437156769d4326f782d9)
- https.tls_cert_options.tls_inline_params.tls_config.custom_security

<a id="canonical-db1529c52b8de255a08059e741fb404d71d2b726f73ae726d16c7071ffe3a3fc"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-7d0db89c11d4f04769e9565d5d0fcc5d4ae7e65924268297c8c2cb03793b7c9b"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 0974c0b58c60 / 3

<a id="canonical-e5f10e9d20ac8a62ef054444a42bf4fbac0e8fe642ac04732b3ceeba1c0ae8cd"></a>

<a id="canonical-1171a9a6eede4fe3519b3a051a13653e8ecd93ed929e950c8f397f3d562e3491"></a>

## cipher_suites property — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 0974c0b58c60 / 4

Type: `["list", "string"]`. Computed.

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

<a id="canonical-dac39bb363cc564e36710edcb49b49f60098347d6616da8da104a7b47f37d265"></a>

<a id="canonical-937025d37a68d85ee450dab864bf7f023cb5e61ec867704acad025e33ae6412f"></a>

## max_version property — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 0974c0b58c60 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-622a5af444a6f9701605ba26dcb54013f9f9c49710c7ab35a47cad7e85f37200"></a>

<a id="canonical-aae116f4ef2f35c03e793912bbf5a97a71b5597b30e134aa052fccc24d325f03"></a>

## min_version property — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 0974c0b58c60 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

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

<a id="canonical-519df7101ee068e9ced8fab9c4af24b9ffa8898c39e1848807ef2a2c591dcc5b"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_config.custom_security / 0974c0b58c60 / 7

- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-574182211b3e722bb2d0f6d4a1dedc2e117fcca6866a437156769d4326f782d9)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-5fecc5d862578277a0548f9accdb27e058b1f33d771160e46e412d1b57e1316f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f82d0d684c5d038c092917fc17f79fa5ba2d3ab93dc4805806b80e11238ceaa"></a>

## https.tls_cert_options.tls_inline_params.tls_config.default_security — https.tls_cert_options.tls_inline_params.tls_config.default_security / 225960b2d863 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-574182211b3e722bb2d0f6d4a1dedc2e117fcca6866a437156769d4326f782d9)
- https.tls_cert_options.tls_inline_params.tls_config.default_security

<a id="canonical-1275fca419ee74416a7f4b9d19fa8a169ca13a69a1b1ffb801a8cf0d6dd24c72"></a>

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

<a id="canonical-0b1429e5cb62a284892f3a094fc86025535b270b2078f2f7760cc1e63bff9de4"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_config.default_security / 225960b2d863 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a2e73703f92b0326110c9be55d9bc6e62fbfd3814ddf66a08d445bd217bca5c8"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_config.default_security / 225960b2d863 / 4

- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-574182211b3e722bb2d0f6d4a1dedc2e117fcca6866a437156769d4326f782d9)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-da2be12eccbcb9791c45c5174bfdb65cb90601ed8706d945e17d0720c0a6216e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4610a97088e6cdee01302f62af79a8b047e39cb000b454eadbbf7d496ae65dc"></a>

## https.tls_cert_options.tls_inline_params.tls_config.low_security — https.tls_cert_options.tls_inline_params.tls_config.low_security / 294e1cbf88b6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-574182211b3e722bb2d0f6d4a1dedc2e117fcca6866a437156769d4326f782d9)
- https.tls_cert_options.tls_inline_params.tls_config.low_security

<a id="canonical-4c658e0a7ed0a3a74404b3d6ec1346bb5c444044aedc40e9105b343f7f169e36"></a>

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

<a id="canonical-4fbb4dcc047afacd1cf982d7297907ad325d7f3db1b0e3b4c3cb68c23e6c9817"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_config.low_security / 294e1cbf88b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e02c8c297175efbe514e4780c7a4423fb03d18dabd843cff6bebb4847d19094f"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_config.low_security / 294e1cbf88b6 / 4

- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-574182211b3e722bb2d0f6d4a1dedc2e117fcca6866a437156769d4326f782d9)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-570e51d3928542c0288e95f2606cb5cb89001eb79b212e57a81c4f2663919230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56adf344bf1101697169f14275397662fe60ca84dd8728ff63e787472382a59a"></a>

## https.tls_cert_options.tls_inline_params.tls_config.medium_security — https.tls_cert_options.tls_inline_params.tls_config.medium_security / 610e9db2e228 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-574182211b3e722bb2d0f6d4a1dedc2e117fcca6866a437156769d4326f782d9)
- https.tls_cert_options.tls_inline_params.tls_config.medium_security

<a id="canonical-0c977efdd0919627d90b854bd1217a825fe056fd6f0a539dc688437983402a00"></a>

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

<a id="canonical-3bcf194d65dd10a4432852d418bd9e69fe58b5fe538860f3adeb19ac53513bed"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_config.medium_security / 610e9db2e228 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-352ca5cdea705bd111838ca5b10568f3eca68270a0f0a5d8cc8823a66aadd57f"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_config.medium_security / 610e9db2e228 / 4

- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-574182211b3e722bb2d0f6d4a1dedc2e117fcca6866a437156769d4326f782d9)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8fd995a0c162672111aff3a405547f81d06ad472505b01c6b70de56d3afe1f33"></a>

## https.tls_cert_options.tls_inline_params.use_mtls — https.tls_cert_options.tls_inline_params.use_mtls / e7e622ac1cc4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- https.tls_cert_options.tls_inline_params.use_mtls

<a id="canonical-13a066447a83dce174dc9c6f7f67b3e66c9a0c8b8d72a61724d04acc0a72905e"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

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

<a id="canonical-2ad0db4255446c73383959d5809f61214d0142d53dd58059c54aa06e76aa7960"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls / e7e622ac1cc4 / 3

<a id="canonical-bc660a7c6190038179ba8e6d479082068749fd65334c04f172706c2e4cf8b5ad"></a>

<a id="canonical-b7eac98d6d9b4795aa43f87516817c49d0524c14b2415584ee69db12df81510d"></a>

## client_certificate_optional property — https.tls_cert_options.tls_inline_params.use_mtls / e7e622ac1cc4 / 4

Type: `"bool"`. Computed.

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

- [crl](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-18d21a712b64bad464a0155dd98d01254e55df6ae215abd15ae3e33e29e5103b): complete subsection reference.

- [no_crl](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-34f9e789565c143ddf546849947174ed6966c721612419e92144a69e04f5b27c): complete subsection reference.

- [trusted_ca](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-dee964794359b1ae939791dd696cb2cb51f2483c1b525ba1dedd4e82d4949a94): complete subsection reference.

<a id="canonical-a80aaa7e9c596211e41105ff837e6e69e302e8202d3c7a9eade04de52f8662f0"></a>

<a id="canonical-f5ab13cbef9d9227b0ea6b670a554fecfffedbf73c7ca4c86e5e085ec9d61633"></a>

## trusted_ca_url property — https.tls_cert_options.tls_inline_params.use_mtls / e7e622ac1cc4 / 5

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

- [xfcc_disabled](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-5d63f67a8c1b5afeb0ff121b1605220d0f4f4a7262f8f4e806126b39b21c1fb9): complete subsection reference.

- [xfcc_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-a66308d5c703ac99cf99bee4dd0c1046754d5a7bc5c524e6836e9b39de75bfe8): complete subsection reference.

<a id="canonical-d43edd649c8dd3727f6f7db836daa4951d9f92393dc8d490faa3896cf510e7d1"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls / e7e622ac1cc4 / 6

- [https.tls_cert_options.tls_inline_params.use_mtls.crl](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-18d21a712b64bad464a0155dd98d01254e55df6ae215abd15ae3e33e29e5103b)
- [https.tls_cert_options.tls_inline_params.use_mtls.no_crl](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-34f9e789565c143ddf546849947174ed6966c721612419e92144a69e04f5b27c)
- [https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-dee964794359b1ae939791dd696cb2cb51f2483c1b525ba1dedd4e82d4949a94)
- [https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-5d63f67a8c1b5afeb0ff121b1605220d0f4f4a7262f8f4e806126b39b21c1fb9)
- [https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-a66308d5c703ac99cf99bee4dd0c1046754d5a7bc5c524e6836e9b39de75bfe8)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-18d21a712b64bad464a0155dd98d01254e55df6ae215abd15ae3e33e29e5103b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0482a243e01aef6740b3c746f2942f51a5a030c8dddd75c9ad68a365ef66237a"></a>

## https.tls_cert_options.tls_inline_params.use_mtls.crl — https.tls_cert_options.tls_inline_params.use_mtls.crl / 7a48d4339122 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097)
- https.tls_cert_options.tls_inline_params.use_mtls.crl

<a id="canonical-08f86d0f212e5a25848d4376b9aaf6283a967de9ee092e732786a7c057a3c376"></a>

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

<a id="canonical-81a700f239e357027d64eee4fd416db69cd266fc4488d55a018a63a6c180a33c"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls.crl / 7a48d4339122 / 3

<a id="canonical-4ff4a7b1a198678ac0ca32c41ba7f56bc07a3e029fd081778a471c9ff40532fb"></a>

<a id="canonical-cc56ba27cb172d2ebb7457f48da572400d8b6942267313ece2692ffeda56197d"></a>

## name property — https.tls_cert_options.tls_inline_params.use_mtls.crl / 7a48d4339122 / 4

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

<a id="canonical-92f2b37ad66b5a65034f8447374ece918848f4b9bac7874459b21cdca0a92629"></a>

<a id="canonical-1be9511ab7a407222ca2407a98903eb808f353584f9815a982888fe69ad1b99e"></a>

## namespace property — https.tls_cert_options.tls_inline_params.use_mtls.crl / 7a48d4339122 / 5

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

<a id="canonical-5b0dcd64ba45f166676e49603ea45706dbf3c0fb314d6dc565c2f4e9ee404bdb"></a>

<a id="canonical-d67091341bd82b145a53b9c2dd0b03d52c54c4e86742bd94ea04d88696e2efb7"></a>

## tenant property — https.tls_cert_options.tls_inline_params.use_mtls.crl / 7a48d4339122 / 6

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

<a id="canonical-13ab86d653e0c0e2153c3d1aabdca75201fe98f97df3736fa6b2616d141d1046"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls.crl / 7a48d4339122 / 7

- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-34f9e789565c143ddf546849947174ed6966c721612419e92144a69e04f5b27c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50f3ad7e5e2c36cdddaebf87964aef0e793dafc8cffbbf0e296c8b75244a28c7"></a>

## https.tls_cert_options.tls_inline_params.use_mtls.no_crl — https.tls_cert_options.tls_inline_params.use_mtls.no_crl / f8517575dfd6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097)
- https.tls_cert_options.tls_inline_params.use_mtls.no_crl

<a id="canonical-0b45066fd69cf5e96563738ba0362869bf069b5672689320b914d473424f1aa7"></a>

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

<a id="canonical-ebf103f9b882be0d7161348be3f1a5532a077492783ff7ab4751a44534a345b7"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls.no_crl / f8517575dfd6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3dfcd0cfd6c4678a301c57b881c92fbb47088b3682dc89a63956a2ff2984ef0f"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls.no_crl / f8517575dfd6 / 4

- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-dee964794359b1ae939791dd696cb2cb51f2483c1b525ba1dedd4e82d4949a94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60fc45ed68389bf9a93854ff227716246fde3b0b33843e5897753439b226a9f5"></a>

## https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / 42421eb0551e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097)
- https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca

<a id="canonical-523d388961364276f77882fd95d32c42285b6624ae08b32f4e32360c4d0e6983"></a>

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

<a id="canonical-8ae4c22f69d4d12d5cff409a199346fb9a99e247f3b8fbd32f01fc6ea3f832b3"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / 42421eb0551e / 3

<a id="canonical-e8a88e45c714157ff11ecb1b48be1773a8c92e9eb23196929cfeb50ef99479cf"></a>

<a id="canonical-57d3383ae2a91fe90d4ca815c9c0262e8972e4ee51c3e3dde84549aea02a06f4"></a>

## name property — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / 42421eb0551e / 4

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

<a id="canonical-7e2283046e3b52cb6782df30f2525d1856558aba7ab9bc6f63ae2321a861d274"></a>

<a id="canonical-aba96c4936033f4a63d6bf0f1300b20e4551042c94103d68cc23c8793af034a0"></a>

## namespace property — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / 42421eb0551e / 5

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

<a id="canonical-896d9f398cd96de39e60b24ff56aac1f8df8fdcafe2810bb030e7a8d503cf140"></a>

<a id="canonical-d720b45d005e1608be15efbc7674de0a4cda23452fad22bf47742d87cfeabe5e"></a>

## tenant property — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / 42421eb0551e / 6

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

<a id="canonical-56c8235701ee916cba3180a93ffbd4c4fecc5ddcd31166d816068e014208f6c7"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls.trusted_ca / 42421eb0551e / 7

- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-5d63f67a8c1b5afeb0ff121b1605220d0f4f4a7262f8f4e806126b39b21c1fb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddb90c7c9b3f09f72e9af2a1cbf12cb089cd1b6dde57e107340053412a6e1579"></a>

## https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled / d8a0da5eb761 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097)
- https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled

<a id="canonical-fabd28b5b3d687c61e7679801ac11a65f829a993d52b82f828f8bbee4a21a8db"></a>

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

<a id="canonical-576250a4acc53debaa5d3061fd15543d3b624aff009fc34c99181ef148970c33"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled / d8a0da5eb761 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9f11d42bc6c5efce30fe1afcebe3b5cd8c334d8f578288504780eb24762cba7f"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_disabled / d8a0da5eb761 / 4

- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-a66308d5c703ac99cf99bee4dd0c1046754d5a7bc5c524e6836e9b39de75bfe8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5255326d6642a5a45c66776a3332a90f833a600c2d238ea5ed736d78c52b365"></a>

## https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options / 0cd1886598a2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097)
- https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options

<a id="canonical-05ef00153b4a3aa75459210d76385b0f90808a9e621cc210195231873035c539"></a>

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

<a id="canonical-59ad0211a013b8e81d463cadca2926d950ce3cc1b618cdacc74c93f4f576c617"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options / 0cd1886598a2 / 3

<a id="canonical-5d0131f5fed504388ab00d9c95720e9ee0f58ea16405c91e35b2f01931ed7451"></a>

<a id="canonical-67e7fe614b9e99cb2752d68b2479bf720e2075ac4640da76befe58c620596d97"></a>

## xfcc_header_elements property — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options / 0cd1886598a2 / 4

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

<a id="canonical-bcbb53554fe30d982df91ee08d04b11ac444f566eeb44d1e514fc5a5af38da6f"></a>

## Next pages — https.tls_cert_options.tls_inline_params.use_mtls.xfcc_options / 0cd1886598a2 / 5

- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-4a52b5f55a6121f3992d8b42133a29c221108de5d07075f61e7f8570d8cc5970"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b9efba9710ef4b3f263a3d3f474baf73b52d8704dfab1e2399347bfd2bc1fff"></a>

## https_auto_cert — https_auto_cert / 9cc7ad0678a3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- https_auto_cert

<a id="canonical-5f2ca249db3615bb879ecfc499fd83a7b2b924f8f0e748722509e5a423395e35"></a>

Type: `"single"`. Computed.

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

<a id="canonical-9ea34211d320183c2aad5b333c51d6ee7eb4c84bf537828838fb6976da90754e"></a>

## Direct properties — https_auto_cert / 9cc7ad0678a3 / 3

<a id="canonical-89ac5a95605d83b6cc3cb600691ab77d187da0153fa362a63fb9cf9c7588691a"></a>

<a id="canonical-b7854139b63cfd23036fbd0dccaab79832d0a921ed9dbb207b240ecbeeee6f27"></a>

## add_hsts property — https_auto_cert / 9cc7ad0678a3 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-456b24350eebb91d16ca9309b3e55a2be73a5465aee148e62b2a15a8c9f57839"></a>

<a id="canonical-7e7da68f0cb4e2202048c00d57240e5d493e94299ea669064e41ea25fc598e86"></a>

## http_redirect property — https_auto_cert / 9cc7ad0678a3 / 5

Type: `"bool"`. Computed.

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

- [tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-a0c7779e64a78e690cbc63a86ca1f2f1f51dae6425fe9bd070ebd5d24463510c): complete subsection reference.

<a id="canonical-9259126b0e4e2b8790d3423843219836698b5ba86038037ac9ef5555a5503069"></a>

## Next pages — https_auto_cert / 9cc7ad0678a3 / 6

- [https_auto_cert.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-a0c7779e64a78e690cbc63a86ca1f2f1f51dae6425fe9bd070ebd5d24463510c)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-a0c7779e64a78e690cbc63a86ca1f2f1f51dae6425fe9bd070ebd5d24463510c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d3c0badb19fc8be948816728999c85e88b776c3173e1522c69ebb5f0d5e32a6"></a>

## https_auto_cert.tls_config — https_auto_cert.tls_config / 5c31083b20bb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https_auto_cert](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-4a52b5f55a6121f3992d8b42133a29c221108de5d07075f61e7f8570d8cc5970)
- https_auto_cert.tls_config

<a id="canonical-77a10157747ef5c53099b619d43e71b35e83e195a619c0be72c16503e6c20acb"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

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

<a id="canonical-54836c0b6af03086f6f59f47b38ec12dd5c1683d9b43e82ac3f5a7f3cefcda34"></a>

## Direct properties — https_auto_cert.tls_config / 5c31083b20bb / 3

- [tls_11_plus](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0e3d91fc88a39747d6a8179333e8d3e609a6e32afc39aff3375c155767c52bfb): complete subsection reference.

- [tls_12_plus](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-138762d9a1f1043d629f39b82d6d043fd276515f7a03696bebc40366facc4bd2): complete subsection reference.

<a id="canonical-bf34c647c9d896029634411af233a6bef4d8bfdd299c72b1c8166fd3ebd0600b"></a>

## Next pages — https_auto_cert.tls_config / 5c31083b20bb / 4

- [https_auto_cert.tls_config.tls_11_plus](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0e3d91fc88a39747d6a8179333e8d3e609a6e32afc39aff3375c155767c52bfb)
- [https_auto_cert.tls_config.tls_12_plus](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-138762d9a1f1043d629f39b82d6d043fd276515f7a03696bebc40366facc4bd2)
- [https_auto_cert](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-4a52b5f55a6121f3992d8b42133a29c221108de5d07075f61e7f8570d8cc5970)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-0e3d91fc88a39747d6a8179333e8d3e609a6e32afc39aff3375c155767c52bfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c78198279a01843df3a4804537c60b7ee36e0af5cbf0fe1967e9e14c17ee020b"></a>

## https_auto_cert.tls_config.tls_11_plus — https_auto_cert.tls_config.tls_11_plus / acbabc845c6b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https_auto_cert](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-4a52b5f55a6121f3992d8b42133a29c221108de5d07075f61e7f8570d8cc5970)
- [https_auto_cert.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-a0c7779e64a78e690cbc63a86ca1f2f1f51dae6425fe9bd070ebd5d24463510c)
- https_auto_cert.tls_config.tls_11_plus

<a id="canonical-051948f3d2f00cbda7087af870d02889e8bf4f2b34da4d91b9a937c6fc73fd63"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-eb4957fd578a0b08338ca016d3f34e3e37977817a5e5896c3734b7427b97edb6"></a>

## Direct properties — https_auto_cert.tls_config.tls_11_plus / acbabc845c6b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1b1da5ccd296469ee0887a66d9058035893d4abdffcf6559fd25ee80e8ea045"></a>

## Next pages — https_auto_cert.tls_config.tls_11_plus / acbabc845c6b / 4

- [https_auto_cert.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-a0c7779e64a78e690cbc63a86ca1f2f1f51dae6425fe9bd070ebd5d24463510c)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-138762d9a1f1043d629f39b82d6d043fd276515f7a03696bebc40366facc4bd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-749b14fbae98a1e453ac60b82487e0b0837ca769827930cecf8d9d3014d383c7"></a>

## https_auto_cert.tls_config.tls_12_plus — https_auto_cert.tls_config.tls_12_plus / 80367a81d1ed / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https_auto_cert](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-4a52b5f55a6121f3992d8b42133a29c221108de5d07075f61e7f8570d8cc5970)
- [https_auto_cert.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-a0c7779e64a78e690cbc63a86ca1f2f1f51dae6425fe9bd070ebd5d24463510c)
- https_auto_cert.tls_config.tls_12_plus

<a id="canonical-bc468dc548839fd4d49eb92c4f3f836e7fb7a7755cb5f086bb7cd94eb7145f55"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-c6a8180777aee0d265550a4947e962e970eb643054905d90d23174bceb63e962"></a>

## Direct properties — https_auto_cert.tls_config.tls_12_plus / 80367a81d1ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-def650497aa7bdf6c47d54e0f213aa1f8e57c1f2651c12eb219c86f7851b9f25"></a>

## Next pages — https_auto_cert.tls_config.tls_12_plus / 80367a81d1ed / 4

- [https_auto_cert.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-a0c7779e64a78e690cbc63a86ca1f2f1f51dae6425fe9bd070ebd5d24463510c)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-b2d3b8b7553e766c73a39014b67c60a27e520c4fcf4518c9a44961b4ab3fd0d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-145ab183553ef42c90d0d9440c473e1fedd795780c47a9a8d400c63625392b2f"></a>

## js_challenge — js_challenge / 6d42f1313b78 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- js_challenge

<a id="canonical-fca553cc065e71828f2d70eb6f8cbd6340da7828eeba5edc3d1a3075a58f2ec9"></a>

Type: `"single"`. Computed.

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

<a id="canonical-8bfdf5f78cf013bf7386f9243a479c87b0f134d03fea2d6aa65f43ef41047ee8"></a>

## Direct properties — js_challenge / 6d42f1313b78 / 3

<a id="canonical-3b62c45a6181bd27cd6332465aceac455dd9c819f9595edbaf8c23d79ecd8a41"></a>

<a id="canonical-5ce60bd79394846274cb32c8a7db53f0e8bd576d2ad9b1437cfd401b9fecde79"></a>

## cookie_expiry property — js_challenge / 6d42f1313b78 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

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

<a id="canonical-4e675f6dd6d4ad00ddd6fc0f4dcce0067fe56d24056a07b10fb6ce47590e0896"></a>

<a id="canonical-78dd632f51359b0265e1322e5fbf0e3f3736eef3b255ab9fb19d0daea9e46d05"></a>

## custom_page property — js_challenge / 6d42f1313b78 / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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

<a id="canonical-012cbbd2beea7b547559a7407dda28d464ee818186d8628c873a27f31fd2e9d5"></a>

<a id="canonical-0665948e6b47e9e4b47e3c287d1e0557c4d34237dd290a808ac1f6b9724ccfc8"></a>

## js_script_delay property — js_challenge / 6d42f1313b78 / 6

Type: `"number"`. Computed.

Delay introduced by Javascript, in milliseconds.

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

<a id="canonical-535f0a3d2e2ca604e8e367fdc94960cfac6e21352aa6635ac469305c01bb6275"></a>

## Next pages — js_challenge / 6d42f1313b78 / 7

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-642513ef64b18f6204041dc6db89086ab33debfa717cf4e704d162743f72445c"></a>

## jwt_validation — jwt_validation / 1146c9bb4d4b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- jwt_validation

<a id="canonical-9f1a94b81cb3e0930359473ffeed4b4e279ea5fdfa2eabc21024123a6c8b3bfa"></a>

Type: `"single"`. Computed.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Upstream description:

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

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

<a id="canonical-d43a220b45196a1e5a2bbc7b183ae4d6f81c998caff612a88d6e3d2d1cb1a193"></a>

## Direct properties — jwt_validation / 1146c9bb4d4b / 3

- [action](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-9978f1d42316a4ebe5232ab32b566f82655ac8e18095224094cb3e480d5669bc): complete subsection reference.

- [authorization_server](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-344044f66fa3d565f2bb9a885417a0c4932463eccf8f97ecd48abe524d841d69): complete subsection reference.

- [jwks_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-afd7fa41212ffe890166ff6da34b8d2223d21bb50f41fb8fb57a8f1ac72a75d2): complete subsection reference.

- [mandatory_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-07753c15cf5735c2c6d13e18c7efb84baac1ce09406b2e9eff5a536c7efd3d26): complete subsection reference.

- [reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc): complete subsection reference.

- [target](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-833d342629e8039a4fc66294f5a60113d6a38e384d089bc0e46bf1bb5721ffa2): complete subsection reference.

- [token_location](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-44a2e759df6a87b292f41d18536724b72435b870e57aa2574f3ca1802446080e): complete subsection reference.

<a id="canonical-2c86c4c012975e2305247c5ca8a18f044d503d977a4e024dead30a9a23ea42b8"></a>

## Next pages — jwt_validation / 1146c9bb4d4b / 4

- [jwt_validation.action](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-9978f1d42316a4ebe5232ab32b566f82655ac8e18095224094cb3e480d5669bc)
- [jwt_validation.authorization_server](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-344044f66fa3d565f2bb9a885417a0c4932463eccf8f97ecd48abe524d841d69)
- [jwt_validation.jwks_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-afd7fa41212ffe890166ff6da34b8d2223d21bb50f41fb8fb57a8f1ac72a75d2)
- [jwt_validation.mandatory_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-07753c15cf5735c2c6d13e18c7efb84baac1ce09406b2e9eff5a536c7efd3d26)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc)
- [jwt_validation.target](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-833d342629e8039a4fc66294f5a60113d6a38e384d089bc0e46bf1bb5721ffa2)
- [jwt_validation.token_location](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-44a2e759df6a87b292f41d18536724b72435b870e57aa2574f3ca1802446080e)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-9978f1d42316a4ebe5232ab32b566f82655ac8e18095224094cb3e480d5669bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7188e24c82c6cfbeae4c61767844f1f63a534ab1214630d239e1108c38625d66"></a>

## jwt_validation.action — jwt_validation.action / 9b5aff818f38 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- jwt_validation.action

<a id="canonical-4b01426e29c42a268e554a2328b01e4925cec9fecd0880e4322c99f2a55c40f0"></a>

Type: `"single"`. Computed.

Action

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

<a id="canonical-3ac4123f6be3917fa954f4af949d43ab8b70289cea8d5209b200d6c8a5380102"></a>

## Direct properties — jwt_validation.action / 9b5aff818f38 / 3

- [block](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-b178a4d40bdeb55cf3a712aa8ac3cd84e36ff3b75b095432bc09481dffdc1db7): complete subsection reference.

- [report](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8649df3035be97446e0d87ae48a9fe505ccd2aa92eb175b5899713d3c829bcce): complete subsection reference.

<a id="canonical-db3e90673e6de52ae147f6b8b9cdcb374ed6485768aa1c812935ec3c45991c4f"></a>

## Next pages — jwt_validation.action / 9b5aff818f38 / 4

- [jwt_validation.action.block](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-b178a4d40bdeb55cf3a712aa8ac3cd84e36ff3b75b095432bc09481dffdc1db7)
- [jwt_validation.action.report](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8649df3035be97446e0d87ae48a9fe505ccd2aa92eb175b5899713d3c829bcce)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-b178a4d40bdeb55cf3a712aa8ac3cd84e36ff3b75b095432bc09481dffdc1db7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df0b48772a78a64a426d70baf69703fcd904b5bddedb76d7efb4e4423d4f954a"></a>

## jwt_validation.action.block — jwt_validation.action.block / 2c1fd3ba1828 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.action](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-9978f1d42316a4ebe5232ab32b566f82655ac8e18095224094cb3e480d5669bc)
- jwt_validation.action.block

<a id="canonical-276e5e5c1c211e1b8ca7b837e9515030b4da4b7e7b4541d97008d03798442670"></a>

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

<a id="canonical-6df43c82d8caa23c46fd1ba64633d7b74c81529b528112a8f1afeb89b5493f7f"></a>

## Direct properties — jwt_validation.action.block / 2c1fd3ba1828 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-684a87e3d0b0fe1546771e69c1d5a0128e546864e5ed46155ff7f0ea819f4fc6"></a>

## Next pages — jwt_validation.action.block / 2c1fd3ba1828 / 4

- [jwt_validation.action](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-9978f1d42316a4ebe5232ab32b566f82655ac8e18095224094cb3e480d5669bc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8649df3035be97446e0d87ae48a9fe505ccd2aa92eb175b5899713d3c829bcce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2f63f39233fa688f0d8ada3ecc466175b5bf86493830b3ea8c67b22bd7da60c"></a>

## jwt_validation.action.report — jwt_validation.action.report / 32ec698f9f12 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.action](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-9978f1d42316a4ebe5232ab32b566f82655ac8e18095224094cb3e480d5669bc)
- jwt_validation.action.report

<a id="canonical-bd5f09fe9eb0b003d156c37b4061c1e1fdc1f41e0535abb48e44df31a3764742"></a>

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

<a id="canonical-9da0deb10f0036c479923917706d43bfac5dfdfc445c375f995a8658236e4e87"></a>

## Direct properties — jwt_validation.action.report / 32ec698f9f12 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c6bc1168039d8e0db6eda2d90204ffd559c910077cf86d05903d0aa89b5ef5ae"></a>

## Next pages — jwt_validation.action.report / 32ec698f9f12 / 4

- [jwt_validation.action](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-9978f1d42316a4ebe5232ab32b566f82655ac8e18095224094cb3e480d5669bc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-344044f66fa3d565f2bb9a885417a0c4932463eccf8f97ecd48abe524d841d69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c246a4e63ecccc3fe044641769a2d6d28aea67364068b08a4e7ca86af891ec3a"></a>

## jwt_validation.authorization_server — jwt_validation.authorization_server / 9e253ca42d4a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- jwt_validation.authorization_server

<a id="canonical-bec526dd49e09b846fe05466cf6af643e5f55b91a0a43e90284dd486871d6a44"></a>

Type: `"single"`. Computed.

Reference to Authorization Server object.

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

<a id="canonical-3b37820cd8201fc428ed5eeaba507b1cec1434eb47681fc27b0cb6d590e37527"></a>

## Direct properties — jwt_validation.authorization_server / 9e253ca42d4a / 3

- [authorization_servers](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-9a96d9e041f3b7137cf00c3686345ddef260bae17791f76cf37aa3d6d39d0bb0): complete subsection reference.

<a id="canonical-858af0c63894dde161a13da9d7f369f37e51332b67ae0630c4e6cdca31f724bf"></a>

## Next pages — jwt_validation.authorization_server / 9e253ca42d4a / 4

- [jwt_validation.authorization_server.authorization_servers](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-9a96d9e041f3b7137cf00c3686345ddef260bae17791f76cf37aa3d6d39d0bb0)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-9a96d9e041f3b7137cf00c3686345ddef260bae17791f76cf37aa3d6d39d0bb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9ad2e953edd5be55144178ffc439cdb29a3f4ecfad59e0720911180ebb2f612"></a>

## jwt_validation.authorization_server.authorization_servers — jwt_validation.authorization_server.authorization_servers / aa4ecd5525b8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.authorization_server](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-344044f66fa3d565f2bb9a885417a0c4932463eccf8f97ecd48abe524d841d69)
- jwt_validation.authorization_server.authorization_servers

<a id="canonical-52c6d67ee9db069a5c517f9f3d8a179087ec31879d4e3b08c159457f03f86ca6"></a>

Type: `"list"`. Computed.

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Upstream description:

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

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

<a id="canonical-07cf5c2b2bf37ca0816b8513424ffbcd6cb774ac6f515f5c6b24ff837ec5e16c"></a>

## Direct properties — jwt_validation.authorization_server.authorization_servers / aa4ecd5525b8 / 3

<a id="canonical-aa0bd3648614d1483392da84d3ea87c158c582e91db01e42a6393e85fef6ac3d"></a>

<a id="canonical-e77e04cc12e11275e4d35a7aff1eafa2970da547ea95a38a0bf1eae12e983de9"></a>

## name property — jwt_validation.authorization_server.authorization_servers / aa4ecd5525b8 / 4

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

<a id="canonical-fb36bc15b52c8d83e0a1351df3f2ea5e9195664d2390413dd07cffbf95967eda"></a>

<a id="canonical-eaac1619011d911d98e7fc1e4f8d361e58f16ac250531aeafe3e1b7d4eae0b97"></a>

## namespace property — jwt_validation.authorization_server.authorization_servers / aa4ecd5525b8 / 5

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

<a id="canonical-029ddc0afa9513f20f401a6c04dbe8548e519523203ec0cbcde8923cb4b51b82"></a>

<a id="canonical-c0d5bb5eee5a3fc91a2dae60cb2d316b80f889a6877b25eabcbc3ebb24c83490"></a>

## tenant property — jwt_validation.authorization_server.authorization_servers / aa4ecd5525b8 / 6

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

<a id="canonical-d1d79bac1fe8b8d67601c489150514549f6a8dfcea7d50519b4ff171ac5e8768"></a>

## Next pages — jwt_validation.authorization_server.authorization_servers / aa4ecd5525b8 / 7

- [jwt_validation.authorization_server](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-344044f66fa3d565f2bb9a885417a0c4932463eccf8f97ecd48abe524d841d69)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-afd7fa41212ffe890166ff6da34b8d2223d21bb50f41fb8fb57a8f1ac72a75d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4994c8e21c53a56bef1d803cd2be796d6356d10b5d93b8f2ac1d663fc2d0ba6"></a>

## jwt_validation.jwks_config — jwt_validation.jwks_config / b53d23b7b89e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- jwt_validation.jwks_config

<a id="canonical-31a8e3f59c2b30ce5f82fc2775e01bf75520a5313c5654b0f2bdaa4e21b2ce81"></a>

Type: `"single"`. Computed.

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

<a id="canonical-734cd40cad9f8c81e03f686e3447c680ba3adf0f7d174c316afd85a0e5aab15c"></a>

## Direct properties — jwt_validation.jwks_config / b53d23b7b89e / 3

<a id="canonical-cad8ce61e0a98cbbe1e9eca9bd2b558a0872d80910db10ec2fb9181a71d9f75a"></a>

<a id="canonical-20c5f3aefd14d0163033b4d87ee9eaf2450f44110d15b7d6b98b02aa98af683c"></a>

## cleartext property — jwt_validation.jwks_config / b53d23b7b89e / 4

Type: `"string"`. Computed.

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

<a id="canonical-e3d4e9d7a490bd08e76eaec6108619545da75433af00da0fb0cd95debf3b4657"></a>

## Next pages — jwt_validation.jwks_config / b53d23b7b89e / 5

- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-07753c15cf5735c2c6d13e18c7efb84baac1ce09406b2e9eff5a536c7efd3d26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92f183852a46c8eb4ef7b16445bb144f61886667b34da7c016051cc6ad1e19cd"></a>

## jwt_validation.mandatory_claims — jwt_validation.mandatory_claims / 1a3e0c78b9e6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- jwt_validation.mandatory_claims

<a id="canonical-0a417201d90b482be4aaf619eaf64f69625d91957ceeb994bbd3510808497e13"></a>

Type: `"single"`. Computed.

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

<a id="canonical-ec0feab05ef92e80b5afb76d9bbb39595a5e233b66200b96de0a014ad0939670"></a>

## Direct properties — jwt_validation.mandatory_claims / 1a3e0c78b9e6 / 3

<a id="canonical-d71adffa82fc7b415ba0c29ace369998a9be00ea1f244fc619b84db492124e95"></a>

<a id="canonical-0dc6ea7abba8155e69261daa421253092cdf33d334c9279d84de94de8d4e7a89"></a>

## claim_names property — jwt_validation.mandatory_claims / 1a3e0c78b9e6 / 4

Type: `["list", "string"]`. Computed.

Claim Names. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-9d362855a16192548eabfb884544a755ea982cd8f5ca94d3d7f3869399afc768"></a>

## Next pages — jwt_validation.mandatory_claims / 1a3e0c78b9e6 / 5

- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-614176b2adefd49e5426393bdb75cb134a4a3fdf653a07b8016c9be620bd7b76"></a>

## jwt_validation.reserved_claims — jwt_validation.reserved_claims / 6bdd1954a09a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- jwt_validation.reserved_claims

<a id="canonical-5bf28f1c9ba57893850b71cc4d5f3a8b676c842df77298cf116b342af2029277"></a>

Type: `"single"`. Computed.

Configurable Validation of reserved Claims.

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

<a id="canonical-b8e010b3c0e3151fa5c63491f0832a273e2ebfb35648c3fe22d6ba5e9dada63b"></a>

## Direct properties — jwt_validation.reserved_claims / 6bdd1954a09a / 3

- [audience](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e28382c3111227ae6fe8b2471c0b7a92ce023b14b1b9f09acdca12e022aa2f61): complete subsection reference.

- [audience_disable](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-c8d8f640ac074eadbc321f8bf7432e12b3a429fb59bf59ad41cf8019a3351be8): complete subsection reference.

<a id="canonical-eb22025be498d58aa300f18a0b5113a703b55799ca1d91b1b12d83618219eb8d"></a>

<a id="canonical-f67a2ca1af3073c7d827677a18d34d5ead431b8b188854de5cb6a8ec630ab2aa"></a>

## issuer property — jwt_validation.reserved_claims / 6bdd1954a09a / 4

Type: `"string"`. Computed.

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

- [issuer_disable](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-03228c2b146b37c4a20935795dadbf9d023cdba18937f1abcf663b78b32cfcc9): complete subsection reference.

- [validate_period_disable](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-68d5d29f005bfd8c2777738826500d3d0760cc6f45269262c53b3dea499be62c): complete subsection reference.

- [validate_period_enable](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-f7cce97654b154339440f6ee3ee6ad28d4667077e416150061c65d522a53c4e6): complete subsection reference.

<a id="canonical-488f180d8a3ca7eb15e831e46562d892d5bf583a41058ea04bbaf050352b418e"></a>

## Next pages — jwt_validation.reserved_claims / 6bdd1954a09a / 5

- [jwt_validation.reserved_claims.audience](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e28382c3111227ae6fe8b2471c0b7a92ce023b14b1b9f09acdca12e022aa2f61)
- [jwt_validation.reserved_claims.audience_disable](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-c8d8f640ac074eadbc321f8bf7432e12b3a429fb59bf59ad41cf8019a3351be8)
- [jwt_validation.reserved_claims.issuer_disable](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-03228c2b146b37c4a20935795dadbf9d023cdba18937f1abcf663b78b32cfcc9)
- [jwt_validation.reserved_claims.validate_period_disable](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-68d5d29f005bfd8c2777738826500d3d0760cc6f45269262c53b3dea499be62c)
- [jwt_validation.reserved_claims.validate_period_enable](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-f7cce97654b154339440f6ee3ee6ad28d4667077e416150061c65d522a53c4e6)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e28382c3111227ae6fe8b2471c0b7a92ce023b14b1b9f09acdca12e022aa2f61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-990443d9b88ab5f753b2d641c3c564269775d920becdbf9b852ee53c0b356f0f"></a>

## jwt_validation.reserved_claims.audience — jwt_validation.reserved_claims.audience / 11054df01f56 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc)
- jwt_validation.reserved_claims.audience

<a id="canonical-6099a61610dcd251be9242ae6898f7d2a8f23aed2f2a625a4b2b1b0f0576e8d1"></a>

Type: `"single"`. Computed.

Audiences

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

<a id="canonical-1ba2046714bae07e849a1f7f2f246fae9c03a706053c9428a2e67460b821c54d"></a>

## Direct properties — jwt_validation.reserved_claims.audience / 11054df01f56 / 3

<a id="canonical-1120d6d37c07eb69868d1c711eca48fb1545e628a0af3d5005d93c7c3dc596d5"></a>

<a id="canonical-6bffba29160d1cc4270e62dcd15647492239e209ba37f072617a34f33981bfc3"></a>

## audiences property — jwt_validation.reserved_claims.audience / 11054df01f56 / 4

Type: `["list", "string"]`. Computed.

Values. Configuration parameter for audiences

Upstream description:

Configuration parameter for audiences

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

<a id="canonical-1108023611113d3ef4ccd0251f2fa9c99d62c0f70a982e5827b0dc78480bc51b"></a>

## Next pages — jwt_validation.reserved_claims.audience / 11054df01f56 / 5

- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c8d8f640ac074eadbc321f8bf7432e12b3a429fb59bf59ad41cf8019a3351be8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edd8fe991e363800b1c377ca5f7876e9d6d52afa6666869143262127708d5934"></a>

## jwt_validation.reserved_claims.audience_disable — jwt_validation.reserved_claims.audience_disable / 2ebee6b88d02 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc)
- jwt_validation.reserved_claims.audience_disable

<a id="canonical-f3eaeb1b8e638e688e1d563885b1719c55c3bee28c778694905207197bdb6d33"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1a828b3313f3f6e7563c82f2e9c719701a0321d72064fd9ba4f57f5807468d91"></a>

## Direct properties — jwt_validation.reserved_claims.audience_disable / 2ebee6b88d02 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-691315816b24d978e5f91ae1784266f0e3422ea74dd2851375cc4ca7e01850ad"></a>

## Next pages — jwt_validation.reserved_claims.audience_disable / 2ebee6b88d02 / 4

- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-03228c2b146b37c4a20935795dadbf9d023cdba18937f1abcf663b78b32cfcc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27c2bb5a579091fb6ff303bfd568225f4ae6a25ab13941025b03f1d902ce4cf5"></a>

## jwt_validation.reserved_claims.issuer_disable — jwt_validation.reserved_claims.issuer_disable / 8defb2422eb9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc)
- jwt_validation.reserved_claims.issuer_disable

<a id="canonical-d8ba4e3741ce21bfd2a4140f62be8c7c234f4891d63af39fdd308ca3db1fd3e1"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-135a201b5d10aa70a529c094e4f98a9479feb3e91be6fe7af0e44b5c7cd108f2"></a>

## Direct properties — jwt_validation.reserved_claims.issuer_disable / 8defb2422eb9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-44b727ee549c6cbe5e090c34d57d3118f3dc64de19dda9c4aaf993be6dec6a81"></a>

## Next pages — jwt_validation.reserved_claims.issuer_disable / 8defb2422eb9 / 4

- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-68d5d29f005bfd8c2777738826500d3d0760cc6f45269262c53b3dea499be62c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-109f4ba10f684c468281f0c43f9656500a42cf92a2f704e544ed425040d02fe0"></a>

## jwt_validation.reserved_claims.validate_period_disable — jwt_validation.reserved_claims.validate_period_disable / 998c4dc59a63 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc)
- jwt_validation.reserved_claims.validate_period_disable

<a id="canonical-99808749e244ec539443b81c0dbc446ce698211bc756bec6364776e410f04be7"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-08c735601d7859e0f7193d1514e3bd47f5aeb7bea0b2efa61c4f74e4e2e4eeff"></a>

## Direct properties — jwt_validation.reserved_claims.validate_period_disable / 998c4dc59a63 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a299f69539df6233c9106515bc99a3413b4d3a55c9ecb26787d7570b22c9ec0e"></a>

## Next pages — jwt_validation.reserved_claims.validate_period_disable / 998c4dc59a63 / 4

- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f7cce97654b154339440f6ee3ee6ad28d4667077e416150061c65d522a53c4e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85e2cf5b8d3c9f6a0e9816bb6e4fdec0237e0edda74952d0e94557dcc91159f0"></a>

## jwt_validation.reserved_claims.validate_period_enable — jwt_validation.reserved_claims.validate_period_enable / 3d9d42f32cf6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc)
- jwt_validation.reserved_claims.validate_period_enable

<a id="canonical-188c86125570f10cab95ca8b1d7da73f7f2d1b37fe92ea2e44074c018ac2668d"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-6796ec0d8150160b196b1084bb13f13c1fdd8c7317742b1edbad453abca74498"></a>

## Direct properties — jwt_validation.reserved_claims.validate_period_enable / 3d9d42f32cf6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-576ae14c0b5ecbf64580d8c987a3d3672c4917c5e312af41a7a16b29576421c7"></a>

## Next pages — jwt_validation.reserved_claims.validate_period_enable / 3d9d42f32cf6 / 4

- [jwt_validation.reserved_claims](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8073b2a163ffafe3abff34cb9e8117b2cf2f704da1021a67520660ba9c4929bc)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-833d342629e8039a4fc66294f5a60113d6a38e384d089bc0e46bf1bb5721ffa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8babc8ea8f98b326d12d3c8322dec718b6ef91b3f9ef4ddc96599f7a1016f539"></a>

## jwt_validation.target — jwt_validation.target / 9a420a16a659 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- jwt_validation.target

<a id="canonical-d7d2e6d2beab86211da23e8577818d4a9a9673efafea3aa36c612d13388ffaee"></a>

Type: `"single"`. Computed.

Define endpoints for which JWT token validation will be performed.

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

<a id="canonical-3e541c8a9ec2396adf71c8105d7df340df54884699f81846a3ab9e802d523aa7"></a>

## Direct properties — jwt_validation.target / 9a420a16a659 / 3

- [all_endpoint](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-36252f63c88ca3e97d3e120ba81850902422a5a605247028c0ea1550b3b5da95): complete subsection reference.

- [api_groups](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-af6636b72b54e984e893dd683e385148411bcd6d8ae88396c690225a7e83fc08): complete subsection reference.

- [base_paths](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-ad6115b89cb15c68561b001447f07fd15bebd00668a628204e356f50d5053005): complete subsection reference.

<a id="canonical-70d8969d816af75ca9cc2b87eb6e10a58833accd4e8c6159b72a6205f0c9bcf0"></a>

## Next pages — jwt_validation.target / 9a420a16a659 / 4

- [jwt_validation.target.all_endpoint](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-36252f63c88ca3e97d3e120ba81850902422a5a605247028c0ea1550b3b5da95)
- [jwt_validation.target.api_groups](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-af6636b72b54e984e893dd683e385148411bcd6d8ae88396c690225a7e83fc08)
- [jwt_validation.target.base_paths](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-ad6115b89cb15c68561b001447f07fd15bebd00668a628204e356f50d5053005)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-36252f63c88ca3e97d3e120ba81850902422a5a605247028c0ea1550b3b5da95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0899db44330cbdc5d03cebceb81a5a5a9435f1381d85c3ca4b4f35080bdf501"></a>

## jwt_validation.target.all_endpoint — jwt_validation.target.all_endpoint / 87108868fe38 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.target](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-833d342629e8039a4fc66294f5a60113d6a38e384d089bc0e46bf1bb5721ffa2)
- jwt_validation.target.all_endpoint

<a id="canonical-0744512440ae89e509d33ae7368a0009f4fecbf2dcfa8e6c8e79760737782f09"></a>

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

<a id="canonical-162ffc793d78d0e53d4f69826531d6511fa9334587e4795acf74c59338d7fa83"></a>

## Direct properties — jwt_validation.target.all_endpoint / 87108868fe38 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-933a355e4122ce7263ac79ee432379112abdb5407bcd74cbf86befd04f91376f"></a>

## Next pages — jwt_validation.target.all_endpoint / 87108868fe38 / 4

- [jwt_validation.target](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-833d342629e8039a4fc66294f5a60113d6a38e384d089bc0e46bf1bb5721ffa2)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-af6636b72b54e984e893dd683e385148411bcd6d8ae88396c690225a7e83fc08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edf422cd279ec883b264ab2d359487c9386589e0f1c1217b2ff7c2a8583f4030"></a>

## jwt_validation.target.api_groups — jwt_validation.target.api_groups / c6adef6074f9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.target](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-833d342629e8039a4fc66294f5a60113d6a38e384d089bc0e46bf1bb5721ffa2)
- jwt_validation.target.api_groups

<a id="canonical-13bf8e7a6302d2642acfa6dbb3c881214422ba4ac12a835c2c363845af2022b0"></a>

Type: `"single"`. Computed.

API Groups.

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

<a id="canonical-ebccd16c91c94684346ce40330cd489882bcc5dc62e422b5ef1bc60e53f7282e"></a>

## Direct properties — jwt_validation.target.api_groups / c6adef6074f9 / 3

<a id="canonical-ed4af74e62c75aa3b555faaffc62ecca8ad89d7dd3896e010f5d0c893993acb2"></a>

<a id="canonical-825ac9e689633262bcb866247a2a0fe5cb9404527267a07ebd4f1329d519b192"></a>

## api_groups property — jwt_validation.target.api_groups / c6adef6074f9 / 4

Type: `["list", "string"]`. Computed.

API Groups. Group or collection configuration

Upstream description:

Group or collection configuration

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

<a id="canonical-cf21d9609c1ef291ccdf59ca15233c586eac06d0d46f22cf9aaa0c207fd7ba7a"></a>

## Next pages — jwt_validation.target.api_groups / c6adef6074f9 / 5

- [jwt_validation.target](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-833d342629e8039a4fc66294f5a60113d6a38e384d089bc0e46bf1bb5721ffa2)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ad6115b89cb15c68561b001447f07fd15bebd00668a628204e356f50d5053005"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ece74c8d8c1e88b6151a21213109e288ca89f4590ed2dde1ea09adbe5923f373"></a>

## jwt_validation.target.base_paths — jwt_validation.target.base_paths / 946bf0244fc5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.target](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-833d342629e8039a4fc66294f5a60113d6a38e384d089bc0e46bf1bb5721ffa2)
- jwt_validation.target.base_paths

<a id="canonical-e1a5c45e6a457975bb5f87e5cf8773055e48e2242a3f055b75cc40391a32fa98"></a>

Type: `"single"`. Computed.

Base Paths.

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

<a id="canonical-c5f23b2c6abe323b53c5d4ab85e32bc098f4adfa3161cd962491f7a19165aa25"></a>

## Direct properties — jwt_validation.target.base_paths / 946bf0244fc5 / 3

<a id="canonical-dc420c45e742fb3f9fb0a4d6ab4a2b1b49f06105e062e68a8cf06981c80b6eb7"></a>

<a id="canonical-213d607182888f2a0bbb869596b95dfab4f8724040c9eb4de5764756a3dc59be"></a>

## base_paths property — jwt_validation.target.base_paths / 946bf0244fc5 / 4

Type: `["list", "string"]`. Computed.

Prefix Values. File system or URL path

Upstream description:

File system or URL path

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

<a id="canonical-4a1d83249b8fdba6a46af2c2bb682c17ae7868cf7a8faac7fb45e3f57bb9c048"></a>

## Next pages — jwt_validation.target.base_paths / 946bf0244fc5 / 5

- [jwt_validation.target](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-833d342629e8039a4fc66294f5a60113d6a38e384d089bc0e46bf1bb5721ffa2)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-44a2e759df6a87b292f41d18536724b72435b870e57aa2574f3ca1802446080e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ff9492f6da5f698e195bd7bc96069ec83983c07e60c6cbfb7403caf9856a324"></a>

## jwt_validation.token_location — jwt_validation.token_location / 095238143562 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- jwt_validation.token_location

<a id="canonical-455999228fbbd3097dd7bce4dc681c743a6c0b611e36cfda3db519b93fcd379c"></a>

Type: `"single"`. Computed.

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

<a id="canonical-df338c140ceeb6b72da79fcbabc9e4dee2f816ca40efa689629cdbcfcfaba66a"></a>

## Direct properties — jwt_validation.token_location / 095238143562 / 3

- [bearer_token](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-bb6b47cf05e9b3c5eef6a240c1f640ac288dd98593af727e26d915aae1a0c17e): complete subsection reference.

<a id="canonical-c7e9362da30837bf38d4c08041e9f647d392a26bc85d186b93537a68a99240d8"></a>

## Next pages — jwt_validation.token_location / 095238143562 / 4

- [jwt_validation.token_location.bearer_token](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-bb6b47cf05e9b3c5eef6a240c1f640ac288dd98593af727e26d915aae1a0c17e)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-bb6b47cf05e9b3c5eef6a240c1f640ac288dd98593af727e26d915aae1a0c17e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbc9ea9d964bcf7e5861f5caf6e67768f29acdc425acff90895273ecaeb1dcc0"></a>

## jwt_validation.token_location.bearer_token — jwt_validation.token_location.bearer_token / 1c71908799ab / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-e9c0248b9dccacb6877ca2f8bb94b11c27ba55338a07f3ecd602982f201cd5b4)
- [jwt_validation.token_location](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-44a2e759df6a87b292f41d18536724b72435b870e57aa2574f3ca1802446080e)
- jwt_validation.token_location.bearer_token

<a id="canonical-a6a73aa769459ad7d52b0eee2dca3b53eb5407ac006cc2041dbb3c9c989a18ee"></a>

Type: `"single"`. Computed.

Configuration parameter for bearer token.

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

<a id="canonical-103e81dcb4adad65fff1607ccc03c80017ec032de8e5e3a6f98733318f05e850"></a>

## Direct properties — jwt_validation.token_location.bearer_token / 1c71908799ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e0974210d1d6d81cc6fd4246de46b422c5dded749a6cb74733cce85521756895"></a>

## Next pages — jwt_validation.token_location.bearer_token / 1c71908799ab / 4

- [jwt_validation.token_location](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-44a2e759df6a87b292f41d18536724b72435b870e57aa2574f3ca1802446080e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-7ffe063e878f3a3744dbd23f60e50c97276352019aab855a5dbf85a4d8b05f75"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f9617868c251b5da83a21b6fd8117305b05dbde52f4b16a4428766b1c393326b"></a>

## l7_ddos_action_block — l7_ddos_action_block / 940e1777cd03 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- l7_ddos_action_block

<a id="canonical-19fa633827bb9040f88fa462b0abe6c8487ff3570e7a5976b7df3a7ab4b13b93"></a>

Type: `["object", {}]`. Computed.

\[OneOf: l7\_ddos\_action\_block, l7\_ddos\_action\_default, l7\_ddos\_action\_js\_challenge;
Default: l7\_ddos\_action\_default\] Enable this option

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

- [l7_ddos_action_block](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-19fa633827bb9040f88fa462b0abe6c8487ff3570e7a5976b7df3a7ab4b13b93)
- [l7_ddos_action_default](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-b86f39b8666f2e8e3b1830d8296d9b5b4cc5343aae1b51e5ebe137ed20cc058a)
- [l7_ddos_action_js_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-423490beab0ffbffc54d38b96fb8f1bae639229c1f9b5e286815c9f7822ceea5)

Select alternatives according to the provider validators above.

<a id="canonical-e5e51135eb131e1e0c92481a996f40b6c81021b681e05ea0b832db23125b407e"></a>

## Direct properties — l7_ddos_action_block / 940e1777cd03 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c7f131292d7812f2442f7f6a3a9f4b29f8c35c76f91910be758a2a58c7cef979"></a>

## Next pages — l7_ddos_action_block / 940e1777cd03 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-60c1d2cae5f8d7ef4588ef22dc10e409f573ad59f7fa0749e60a4cc7bb905003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11828f8f9345160ae4c170141bda5f203dba347dc960cdf1a0258ffe58e43246"></a>

## l7_ddos_action_default — l7_ddos_action_default / 67f60d182cc8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- l7_ddos_action_default

<a id="canonical-b86f39b8666f2e8e3b1830d8296d9b5b4cc5343aae1b51e5ebe137ed20cc058a"></a>

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

<a id="canonical-70908946d1c6ece54c7b2398640f27acf8f25354f3a39253c889121d59e96276"></a>

## Direct properties — l7_ddos_action_default / 67f60d182cc8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb80a074fbcf287c061cbe1c90b01c479c7d206f1c55006484109288cd413728"></a>

## Next pages — l7_ddos_action_default / 67f60d182cc8 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-4064035bd1d3552bb9ccb82ffbae60db2624cc9c79d7f8b305061c5af59d3f48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b0a76f977704d6a69e69e525da51e146f59b0b6f9ca61e8ffab7e9e9dfc52e5"></a>

## l7_ddos_action_js_challenge — l7_ddos_action_js_challenge / 8a05b372efd2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- l7_ddos_action_js_challenge

<a id="canonical-423490beab0ffbffc54d38b96fb8f1bae639229c1f9b5e286815c9f7822ceea5"></a>

Type: `"single"`. Computed.

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

<a id="canonical-ee95a4107be5f702f46f1108dfff681c90f13c54ca54928230a8cef6c5162c3b"></a>

## Direct properties — l7_ddos_action_js_challenge / 8a05b372efd2 / 3

<a id="canonical-92c94ecd0e26d404c5fc67ed0aee87106c01a957e73448f95a57ce82f65e7678"></a>

<a id="canonical-3f149b949f14a82d17c7c89c7e6a168a0f5975a689b38547295b4c704918bcd4"></a>

## cookie_expiry property — l7_ddos_action_js_challenge / 8a05b372efd2 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

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

<a id="canonical-008cab72450350d9d4656f337ef514568c2680d2b7754ecfdb7df57354fcc7e4"></a>

<a id="canonical-486a3038565c47b7b8fbe92eaf627ff67eb2d176379c9b7bd8463a3d2d530a61"></a>

## custom_page property — l7_ddos_action_js_challenge / 8a05b372efd2 / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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

<a id="canonical-54765d856246dff8a894776ca8145c02ef5fc3821408535846c2598725de792f"></a>

<a id="canonical-248e0261cb05531642e3af72277fc1d1671236a330f7d93a1305c1ebedea7431"></a>

## js_script_delay property — l7_ddos_action_js_challenge / 8a05b372efd2 / 6

Type: `"number"`. Computed.

Delay introduced by Javascript, in milliseconds.

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

<a id="canonical-960907abc39b199d1b6baa21522b856162365a20ac0750b1c999d72894aa7c70"></a>

## Next pages — l7_ddos_action_js_challenge / 8a05b372efd2 / 7

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-6e529dfa5d955f8b4335cc866733193c137c5e9d565daf1b27a3fa9022f7ec64"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a627c3db4d0fd456219f2a9346a73eb441b880ea01d3525c03976d72b3d312cb"></a>

## no_challenge — no_challenge / 92a2865f9b72 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- no_challenge

<a id="canonical-f4c54e092755798ade79101de0a3985b7ea12009c3514145f0c763f5ebdfe3d9"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge.

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

<a id="canonical-20f91bbe84c22a87885170571cc665491e9a2e1a39bc08749829fa0559199db6"></a>

## Direct properties — no_challenge / 92a2865f9b72 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-814c42b25c697bec70a79c345cfa4a9bb6e53a881171cbe1609a6e80cc0629c5"></a>

## Next pages — no_challenge / 92a2865f9b72 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-107e8e679b873906e8381ab7c4ccbea59f6bd95936693ad2aff2271ba7a5e3a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f7be3de0120e80e357d8ec82d0eac7d0d872e42968edd0711b86e7c1c9ba6e4"></a>

## no_service_policies — no_service_policies / db71521e7375 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- no_service_policies

<a id="canonical-9b7e94307e87a6cd2b20fc90068e31b2e4c48f4b4e2945e954a233679bba53b0"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no service policies.

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

<a id="canonical-bd9c246f64ea7a3e3ad5f1c6e857842e051c267869718a351a94491025b198a2"></a>

## Direct properties — no_service_policies / db71521e7375 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2762e043407fffe4ffe0decebd91caf9fda4328bb183114aa861b3f3b61e7f0c"></a>

## Next pages — no_service_policies / db71521e7375 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-649984c83db4fec5c9ae9862b44a002a5397c26d9c9256cb6c6e050bea7439b6"></a>

## origin_pool — origin_pool / 547a7746a250 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- origin_pool

<a id="canonical-4290960a7ae8d208d862fe1f0b29bdda9742acb025ec998a0ed3e7ac9e6de8cd"></a>

Type: `"single"`. Computed.

Configuration parameter for origin pool.

Upstream description:

Origin Pool for the CDN distribution.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-3c94c5b2cfbd15b1c6d7005136ec8ba6b3196a68ed3911bc984ea511415bb2fe"></a>

## Direct properties — origin_pool / 547a7746a250 / 3

- [more_origin_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-90416fdf2d7e1f84fccdb21dbdabb29227e0ee775a17e98d3b4728652253848d): complete subsection reference.

- [no_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8db3610fb4681543f01e32b471da3bbb996ddb95069052587036b2214879880c): complete subsection reference.

<a id="canonical-65d150d712be408dcb6204ed46827ff61c2c78288393aef0d535120b107e98e0"></a>

<a id="canonical-efe73ed08c5972279eee295d46f44fc10e30d18a74ce8f6d5ec694c7cd42ed1b"></a>

## origin_request_timeout property — origin_pool / 547a7746a250 / 4

Type: `"string"`. Computed.

Configures the time after which a request to the origin will time out waiting for a response.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [origin_servers](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-450049616cdcf87323e1a3fb2252e52303cf2c9dcb6a5eebad4b851cc1788109): complete subsection reference.

- [public_name](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-b12db07aa36ac570baefe9974dd026759f3b85a397686f26b7e15d1547f00053): complete subsection reference.

- [use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50): complete subsection reference.

<a id="canonical-abb2298d0a081b26e921247a27b636354953732a7077848b467ed086694e8e0a"></a>

## Next pages — origin_pool / 547a7746a250 / 5

- [origin_pool.more_origin_options](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-90416fdf2d7e1f84fccdb21dbdabb29227e0ee775a17e98d3b4728652253848d)
- [origin_pool.no_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8db3610fb4681543f01e32b471da3bbb996ddb95069052587036b2214879880c)
- [origin_pool.origin_servers](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-450049616cdcf87323e1a3fb2252e52303cf2c9dcb6a5eebad4b851cc1788109)
- [origin_pool.public_name](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-b12db07aa36ac570baefe9974dd026759f3b85a397686f26b7e15d1547f00053)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-90416fdf2d7e1f84fccdb21dbdabb29227e0ee775a17e98d3b4728652253848d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc575eba34d30e8034515968373b3668c950d6a399359d7926973f7fd7ffa210"></a>

## origin_pool.more_origin_options — origin_pool.more_origin_options / 979924d3bb16 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- origin_pool.more_origin_options

<a id="canonical-61373de2cd473a22dc32717e70b4147cde97fc99922bacba119759c346478387"></a>

Type: `"single"`. Computed.

Configuration parameter for more origin options.

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

<a id="canonical-86ee22a606cb926ce40ecfdb2cb8648a2e2f7842dfbcdda7b20348d178ca5c28"></a>

## Direct properties — origin_pool.more_origin_options / 979924d3bb16 / 3

<a id="canonical-f24d924b3bacb38c6134c9fe81772aba14bbbbd16c67d46d0ccf8d08d456e677"></a>

<a id="canonical-64c0fed271dfe11a28038c87fee569d1888a4197072aa409bf867c484a5d1eb4"></a>

## enable_byte_range_request property — origin_pool.more_origin_options / 979924d3bb16 / 4

Type: `"bool"`. Computed.

Choice to enable/disable byte range requests towards origin.

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

<a id="canonical-2c8e5b081acdd417a97df0e6cc1a52d82418eb408952c6e287bf9d6d0d187222"></a>

<a id="canonical-0419386447e29851911b7dd727338022dc39d71ca87d1c096625e1c643432c93"></a>

## websocket_proxy property — origin_pool.more_origin_options / 979924d3bb16 / 5

Type: `"bool"`. Computed.

Option to enable proxying of websocket connections to the origin server.

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

<a id="canonical-2c8304523c5b11b5595f6a352e94b385b446726526dd6fc83df1d70a04890cdd"></a>

## Next pages — origin_pool.more_origin_options / 979924d3bb16 / 6

- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8db3610fb4681543f01e32b471da3bbb996ddb95069052587036b2214879880c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e4681c22c5e364dc98b1a2a4b73729a5bac350b70dcc0cbc18ce1baea3e870cf"></a>

## origin_pool.no_tls — origin_pool.no_tls / e688581fb125 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- origin_pool.no_tls

<a id="canonical-13b3592ca782ce1bbd2fa95d85d992bab13248b846b80a90b44a426584571c0d"></a>

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

<a id="canonical-dfa3a256b688755ef851603a79e9b4b3a854ddedf92d37022f14cebf62b27658"></a>

## Direct properties — origin_pool.no_tls / e688581fb125 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e6bd8d92319a73130b5a7392148f677cf0def56c827149a1fb607fafd5cac86"></a>

## Next pages — origin_pool.no_tls / e688581fb125 / 4

- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-450049616cdcf87323e1a3fb2252e52303cf2c9dcb6a5eebad4b851cc1788109"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19b2027aef7882685c937b28f395c13244738001507175b378fe1cb7296ecd71"></a>

## origin_pool.origin_servers — origin_pool.origin_servers / 3689f237df70 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- origin_pool.origin_servers

<a id="canonical-f08b56e46b2dbaaa79013146774f83853b0d3e7eb493e662a06a66e63f257f6b"></a>

Type: `"list"`. Computed.

List Of Origin Servers. List of original servers.

Upstream description:

List of original servers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-397902f1b7fe3a651925effb7f3ffb4344294a6b3283e30e48383f28ac83055f"></a>

## Direct properties — origin_pool.origin_servers / 3689f237df70 / 3

<a id="canonical-27ddda56d228a0b56d9c34a6a51fc454f9889cbe8894e05f5184a3e0507bf8f7"></a>

<a id="canonical-2c582c9c67908eb2f20189f84fc9056abe85ad40639549e20df86a9c812972e9"></a>

## port property — origin_pool.origin_servers / 3689f237df70 / 4

Type: `"number"`. Computed.

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

Upstream description:

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

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

- [public_ip](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1e507944ae256217f7d25ecc75cd46cef0050a98cb5952819bc5a5a894a36bc6): complete subsection reference.

- [public_name](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-22f395fe8109dd25606fe9b39ba8c1508cab993d0edd1461c4fd185ea245748f): complete subsection reference.

<a id="canonical-c20ef9dcfba8c39811fdefcc6862387c0a8f1ed8ddd19809c27817055902c9a7"></a>

## Next pages — origin_pool.origin_servers / 3689f237df70 / 5

- [origin_pool.origin_servers.public_ip](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1e507944ae256217f7d25ecc75cd46cef0050a98cb5952819bc5a5a894a36bc6)
- [origin_pool.origin_servers.public_name](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-22f395fe8109dd25606fe9b39ba8c1508cab993d0edd1461c4fd185ea245748f)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-1e507944ae256217f7d25ecc75cd46cef0050a98cb5952819bc5a5a894a36bc6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea988417236ec66748a5ff83049eb01366e540b9a8de1b54b5426856994501ff"></a>

## origin_pool.origin_servers.public_ip — origin_pool.origin_servers.public_ip / 7ee0bd828328 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.origin_servers](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-450049616cdcf87323e1a3fb2252e52303cf2c9dcb6a5eebad4b851cc1788109)
- origin_pool.origin_servers.public_ip

<a id="canonical-60c09d2d85d7ae51adcaed65b1a2ea32129fb9a6f8e61bc9863ad96a273f8687"></a>

Type: `"single"`. Computed.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-8319f97f193a7fb26fa6a2bc128bbc44fee2f4d96ad60176d5e761538b99c59e"></a>

## Direct properties — origin_pool.origin_servers.public_ip / 7ee0bd828328 / 3

<a id="canonical-c831f136f7b4beed71ce9bb0975d94c111e7f0be08cc84652012f9365703ffcc"></a>

<a id="canonical-eecf9e0a1b0ad0687d1a1a0b1762a3757c3c97337eaabd6a4cf7c693041db475"></a>

## ip property — origin_pool.origin_servers.public_ip / 7ee0bd828328 / 4

Type: `"string"`. Computed.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

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

<a id="canonical-342e1b96eb6c587b2e4c5900b9639b9d7f161d2755192a57fa904ad03cd0fb48"></a>

## Next pages — origin_pool.origin_servers.public_ip / 7ee0bd828328 / 5

- [origin_pool.origin_servers](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-450049616cdcf87323e1a3fb2252e52303cf2c9dcb6a5eebad4b851cc1788109)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-22f395fe8109dd25606fe9b39ba8c1508cab993d0edd1461c4fd185ea245748f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e6d2202ae32302d3fc3feba59d4e0049f6c85da265b12ca65663299f3ba2475"></a>

## origin_pool.origin_servers.public_name — origin_pool.origin_servers.public_name / 9c8ef081a92b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.origin_servers](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-450049616cdcf87323e1a3fb2252e52303cf2c9dcb6a5eebad4b851cc1788109)
- origin_pool.origin_servers.public_name

<a id="canonical-c6e0df3185329567281e2aeaf4d150fd9232db721f2ab57e8568bbf13a35eafa"></a>

Type: `"single"`. Computed.

Specify origin server with public DNS name.

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

<a id="canonical-883d51e3ab69529650994de010067f7dc4b3ae4f700a3cb857d30e4784f023f9"></a>

## Direct properties — origin_pool.origin_servers.public_name / 9c8ef081a92b / 3

<a id="canonical-18ed6f52be1649e1c90e68c0ce75d9fa9faa8ecbd6ccc05c7596f5d1ff69bc36"></a>

<a id="canonical-59eb20940ad0f4b6f9a566931e625b9a8d5de427bb8f98d229e06dc78c5b2e3a"></a>

## dns_name property — origin_pool.origin_servers.public_name / 9c8ef081a92b / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-aa28e3baa7f50efab506c597b469f9d7042cf7241b0381a1a5047fe87a5109ef"></a>

<a id="canonical-c1765fda4d850effe0f43de08d26d7ce0809029ae040285ab363de6ef271f4e8"></a>

## refresh_interval property — origin_pool.origin_servers.public_name / 9c8ef081a92b / 5

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-8b43dd69e27bf250bfb51fa2d693985b3eaac4b58eb9c322579534a52e3c5b83"></a>

## Next pages — origin_pool.origin_servers.public_name / 9c8ef081a92b / 6

- [origin_pool.origin_servers](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-450049616cdcf87323e1a3fb2252e52303cf2c9dcb6a5eebad4b851cc1788109)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-b12db07aa36ac570baefe9974dd026759f3b85a397686f26b7e15d1547f00053"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-248840a5880c74ea6a433bb9d0c3ad9980c41bbb223413b57b6e0aa32f0d256e"></a>

## origin_pool.public_name — origin_pool.public_name / 4bfcd05f1aa5 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- origin_pool.public_name

<a id="canonical-9983b8fcd92c1b8b644507c857c7f50d7c6d84925442206429727f997b2a071d"></a>

Type: `"single"`. Computed.

Specify origin server with public DNS name.

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

<a id="canonical-f85402038b7fc51e500503d48494d9e65f9bc513b1c0838328800608bde25bfa"></a>

## Direct properties — origin_pool.public_name / 4bfcd05f1aa5 / 3

<a id="canonical-b6209f2dc5e0e4b93defd6762e29532649a334062c11c2549a5dcc27ced287eb"></a>

<a id="canonical-8382e83bae6603ef00888240ab95a194b5b0b0f72d0a5fd430f61aacb4785de9"></a>

## dns_name property — origin_pool.public_name / 4bfcd05f1aa5 / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-52483d2fbab011a41c8b2eb869b693277dceddc65038265b38274256cf73b72e"></a>

<a id="canonical-90a33918a639f0795185291d405c875f1f5a747d501be438f61a1133ca188af3"></a>

## refresh_interval property — origin_pool.public_name / 4bfcd05f1aa5 / 5

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-074c7220f10a7b3352ea8ad0d0535d8d9ff43338900d99582488b942829570c7"></a>

## Next pages — origin_pool.public_name / 4bfcd05f1aa5 / 6

- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e02765c408d19243ba3628a617f8a0371028fb565eaefa3f0a2bb93aa6c7fbd6"></a>

## origin_pool.use_tls — origin_pool.use_tls / 292af0171eaf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- origin_pool.use_tls

<a id="canonical-0381dc876022549b7cd8497fd155cd85d46f8477c1e1d036dd5f6efaa3bf6124"></a>

Type: `"single"`. Computed.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

<a id="canonical-5ee660417253fce5a14bc6acb9eaf7aa53829ac94d6275bf0457fc4debd379e3"></a>

## Direct properties — origin_pool.use_tls / 292af0171eaf / 3

- [default_session_key_caching](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-34631111ddb1972f7ac8b4b2826b42f8bac98057e3fa162196be071f0b536ec8): complete subsection reference.

- [disable_session_key_caching](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-a1f510fecdd3934f6e7ddd2fc80460163aafa3778ac3d3dac88d36b5e7fb5882): complete subsection reference.

- [disable_sni](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-56bb114871a1dd39d2e7fbadf87b9d2be9a92f9b3fc6a61b2f80d35db3e476d7): complete subsection reference.

<a id="canonical-e8a7bceb8dfdc2ffa7eae34e7c42eb33c6d27838916c4bcd124a407220500fa3"></a>

<a id="canonical-04c3942f3867bcbb4379b78d8643fc87b3faa074f6efcc753b9edd910a0b4308"></a>

## max_session_keys property — origin_pool.use_tls / 292af0171eaf / 4

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-883c2f8749ee274dd4ab23d86bfd2652fb4392f2bb59da5a4f7fe606e8109842): complete subsection reference.

- [skip_server_verification](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-9758e61f7b10658022aedc47c6a6779fbe44f90d423056c89ea84b088fd48762): complete subsection reference.

<a id="canonical-78db75f38510d1eb2fa60ff035bd4069811a3282505e9494d3df089d6c5c3261"></a>

<a id="canonical-b00a5bdd114754af19704f37b2777d6f5376cc1c9f6b156f8b6c7d6a246c344f"></a>

## sni property — origin_pool.use_tls / 292af0171eaf / 5

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8afdf7f8cf145375fd99cca3d9e4220566392bcd21755c507eab2df2a267d3f4): complete subsection reference.

- [use_host_header_as_sni](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-dd65c566b56ea97b3d0a4838833b3a6f009c6b9c5ef34d0ae1e7226350f7322f): complete subsection reference.

- [use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c71ccb872d61a1214a94a25a65841a6c04bdfe46193e54e7746c81b828588213): complete subsection reference.

- [use_mtls_obj](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ee98a678519e4f1a9cda4c6aa3ba19b6f2dd8396c25a322c9fbec0cf854492f7): complete subsection reference.

- [use_server_verification](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d6d3e059c3665f643b999fdf3812795ba2895e20bd5b1ae4ea731267fdd56184): complete subsection reference.

- [volterra_trusted_ca](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-79df3deef533facec1df2592d14b280dd2b8986e8cc122c85987c485e8800cfe): complete subsection reference.

<a id="canonical-dc929f2eea1bdf00ebcb9827c33c562065424394f565f2b526bfc3eb20dcbee0"></a>

## Next pages — origin_pool.use_tls / 292af0171eaf / 6

- [origin_pool.use_tls.default_session_key_caching](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-34631111ddb1972f7ac8b4b2826b42f8bac98057e3fa162196be071f0b536ec8)
- [origin_pool.use_tls.disable_session_key_caching](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-a1f510fecdd3934f6e7ddd2fc80460163aafa3778ac3d3dac88d36b5e7fb5882)
- [origin_pool.use_tls.disable_sni](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-56bb114871a1dd39d2e7fbadf87b9d2be9a92f9b3fc6a61b2f80d35db3e476d7)
- [origin_pool.use_tls.no_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-883c2f8749ee274dd4ab23d86bfd2652fb4392f2bb59da5a4f7fe606e8109842)
- [origin_pool.use_tls.skip_server_verification](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-9758e61f7b10658022aedc47c6a6779fbe44f90d423056c89ea84b088fd48762)
- [origin_pool.use_tls.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-8afdf7f8cf145375fd99cca3d9e4220566392bcd21755c507eab2df2a267d3f4)
- [origin_pool.use_tls.use_host_header_as_sni](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-dd65c566b56ea97b3d0a4838833b3a6f009c6b9c5ef34d0ae1e7226350f7322f)
- [origin_pool.use_tls.use_mtls](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-c71ccb872d61a1214a94a25a65841a6c04bdfe46193e54e7746c81b828588213)
- [origin_pool.use_tls.use_mtls_obj](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-ee98a678519e4f1a9cda4c6aa3ba19b6f2dd8396c25a322c9fbec0cf854492f7)
- [origin_pool.use_tls.use_server_verification](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-d6d3e059c3665f643b999fdf3812795ba2895e20bd5b1ae4ea731267fdd56184)
- [origin_pool.use_tls.volterra_trusted_ca](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-79df3deef533facec1df2592d14b280dd2b8986e8cc122c85987c485e8800cfe)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-34631111ddb1972f7ac8b4b2826b42f8bac98057e3fa162196be071f0b536ec8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-600a87129865795a5f8c4e0bc7f1ea4d13ba8baf3c02514e918f3e106a786105"></a>

## origin_pool.use_tls.default_session_key_caching — origin_pool.use_tls.default_session_key_caching / 4c9875faead2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- origin_pool.use_tls.default_session_key_caching

<a id="canonical-68eb98198674fda2dba9cd66e5437ddc9e51c9b002e53c594574bf95dc0f4c9a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

<a id="canonical-a0c6f6984e31b9fb8469e6a1028e4391f9f9d2fe8ba3d742af54c83a91576187"></a>

## Direct properties — origin_pool.use_tls.default_session_key_caching / 4c9875faead2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-af24576fdd62687a57ec4dee0fec610793b84d1bfbe3ee3cc15da6132a519e2e"></a>

## Next pages — origin_pool.use_tls.default_session_key_caching / 4c9875faead2 / 4

- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-a1f510fecdd3934f6e7ddd2fc80460163aafa3778ac3d3dac88d36b5e7fb5882"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-177de2362f470d244c45ea2f6a3c5988615952bfe27e5c7040070580670df3f4"></a>

## origin_pool.use_tls.disable_session_key_caching — origin_pool.use_tls.disable_session_key_caching / 9ffe518eccbb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- origin_pool.use_tls.disable_session_key_caching

<a id="canonical-213c2cfe23cd54a13d938804373acd5f8a908155b37afd2eb6a745eee994598d"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable session key caching.

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

<a id="canonical-92937adbf0c5ea32af655bf6d9021ccead3479fd6be0eb8f399900c3f62861a6"></a>

## Direct properties — origin_pool.use_tls.disable_session_key_caching / 9ffe518eccbb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ab0d918a9e7133c3307bfbfbca3f7b014e58cf0c87b80e58f0e4032e3cd03553"></a>

## Next pages — origin_pool.use_tls.disable_session_key_caching / 9ffe518eccbb / 4

- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-56bb114871a1dd39d2e7fbadf87b9d2be9a92f9b3fc6a61b2f80d35db3e476d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35a5a34dd7451982fb2094102462c595394a0ee96af9fab6340814a73752164d"></a>

## origin_pool.use_tls.disable_sni — origin_pool.use_tls.disable_sni / 8866ef972a79 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- origin_pool.use_tls.disable_sni

<a id="canonical-2febc76fa6288a23ebbe985b4fcf4f606e3da84d6e2e3556c10a148a3360dbce"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable sni.

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

<a id="canonical-6c95d4793990945cd229f4c255eecc881547a00528840d1ae94737929f2cdffa"></a>

## Direct properties — origin_pool.use_tls.disable_sni / 8866ef972a79 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3fa9a498683bb14512be27d151c140c4cbcd93287c77d7127a8746bfbf21b4b0"></a>

## Next pages — origin_pool.use_tls.disable_sni / 8866ef972a79 / 4

- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-883c2f8749ee274dd4ab23d86bfd2652fb4392f2bb59da5a4f7fe606e8109842"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b48b82706f8a96ebb510086dfeda4fc58d69c263461dde572180d34dde01f21f"></a>

## origin_pool.use_tls.no_mtls — origin_pool.use_tls.no_mtls / 96583c8ccd82 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- origin_pool.use_tls.no_mtls

<a id="canonical-75b5cee1b37ec745cc28de470cdac13d39d512054c798e9870f30c6d05e7d20b"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-bbbfaba7ad711e4e893cafc0dc59552f4007402c4063d8711d3259d2b4c4064f"></a>

## Direct properties — origin_pool.use_tls.no_mtls / 96583c8ccd82 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1652232df9a59a12466d62813bb64ed1b1fcbc9a6c4961dce1378d354bc8d63d"></a>

## Next pages — origin_pool.use_tls.no_mtls / 96583c8ccd82 / 4

- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-9758e61f7b10658022aedc47c6a6779fbe44f90d423056c89ea84b088fd48762"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b6baf4c4cef3410a4471abb60cada94622d21e71352efc453318365c20862a0"></a>

## origin_pool.use_tls.skip_server_verification — origin_pool.use_tls.skip_server_verification / babf50981c52 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- origin_pool.use_tls.skip_server_verification

<a id="canonical-e0d16a767df7ed0b0f15d6d402fdcb05db90134ca12c1ee42398e4812b4dc871"></a>

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

<a id="canonical-afdda079e0f9a48c12b19d1aeb04ed5fe12adb48ef52018c94fcfe3b3423b4d2"></a>

## Direct properties — origin_pool.use_tls.skip_server_verification / babf50981c52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-37fb97568ec83212fecb792a11eb7a9258533e1f43ffd9694fc65d24caff5638"></a>

## Next pages — origin_pool.use_tls.skip_server_verification / babf50981c52 / 4

- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8afdf7f8cf145375fd99cca3d9e4220566392bcd21755c507eab2df2a267d3f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4dfcbaf24dfcf0dc56ae5e4c2e37fbbe2869e15f1b20787b49294466d14aa59"></a>

## origin_pool.use_tls.tls_config — origin_pool.use_tls.tls_config / 67321f3fc13c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [origin_pool](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2bb6e1e94d7f9314d889119c7c85fbf1bf8b508164ddd71ba5a583ccf8b25654)
- [origin_pool.use_tls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-21958b2ec893b946e5983d5551766301213202c56530a745016192a51eceef50)
- origin_pool.use_tls.tls_config

<a id="canonical-faa79a929d6f50b7e55f5f3c17a2220531ac55820685f5de27f8b84829a70944"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```
