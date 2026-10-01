---
page_title: "xcsh_secret_management_access reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access reference."
---

# xcsh_secret_management_access reference

<a id="canonical-c1f03d8ee72adc6b4c08e553cb2ee25cde45f624bb244d1b5cb9617545eeb8f2"></a>

## Next pages — access_info.vault_auth_info.app_role_auth.secret_id / 45f6430b5a95 / 4

- [access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-7787e39d9c44b113fb9cd00b252fe4ac504c2aaca285bca733dabaa3487a773c)
- [access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-e477db4b9c8b0e26d8a2127c50dde697c089b80bc1d046c914aee8f803511e29)
- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-06c78926d465b5c09ae4a9137a0981c0b52426d1d96d985fce7e04c79f60070a)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-7787e39d9c44b113fb9cd00b252fe4ac504c2aaca285bca733dabaa3487a773c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69e3fcd5f6e9f42e0daf60c5a9ce38149fdfaab2d2ad1c8b74d5695a2c78d106"></a>

## access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 3e997698c01b / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a)
- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-06c78926d465b5c09ae4a9137a0981c0b52426d1d96d985fce7e04c79f60070a)
- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-001.md#canonical-8ce6fbd2b6d65d0662abfee3d6e6541f76975430c54554562baf3898f0fc6c88)
- access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info

<a id="canonical-5b371b4ec88addc4bedec4b999e85052b1c37ff0a8543f4a5f171cbdb6900641"></a>

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

<a id="canonical-42b82860a81a64d54336dd2046c0856342c0a7d9b6a5d8bec8cf07e88e3d78bc"></a>

## Direct properties — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 3e997698c01b / 3

<a id="canonical-fea58533dda501a623605e7043e3ac36a35d2217c9290b5d0de944819bb0900a"></a>

<a id="canonical-c2e3959c11fab60735225031d7e52a0b01f4666feae8eeaacfbac5dc1850bee1"></a>

## decryption_provider property — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 3e997698c01b / 4

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

<a id="canonical-861a3856a789e875fff848887cd21c2841f9b044ba7d3df728cabb8b7c70995b"></a>

<a id="canonical-43967004ffdb91a20e7b719e8926c1e1881ab31e46e9a9d288dce199cf914ac4"></a>

## location property — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 3e997698c01b / 5

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

<a id="canonical-666974a7b681c8972cf33e0f4ab7ae9b2553f733795df9d6575501a399049edc"></a>

<a id="canonical-05830e061f4cee7e36188abce2fb6fc7bb2c833a2606414149c175c4bc60f488"></a>

## store_provider property — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 3e997698c01b / 6

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

<a id="canonical-51fb46ad4eb74e066498f46ddf58c2c2ac3922422aaf23366ddfeaac1477a6cb"></a>

## Next pages — access_info.vault_auth_info.app_role_auth.secret_id.blindfold_secret_info / 3e997698c01b / 7

- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-001.md#canonical-8ce6fbd2b6d65d0662abfee3d6e6541f76975430c54554562baf3898f0fc6c88)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-e477db4b9c8b0e26d8a2127c50dde697c089b80bc1d046c914aee8f803511e29"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83931cfbc7bee5ba2e9063e7293ee13128fd4399a536c49aeebdb71c60eaeafd"></a>

## access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info — access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info / 81d9d0156c3f / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a)
- [access_info.vault_auth_info.app_role_auth](data-sources--secret_management_access--reference--group-001.md#canonical-06c78926d465b5c09ae4a9137a0981c0b52426d1d96d985fce7e04c79f60070a)
- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-001.md#canonical-8ce6fbd2b6d65d0662abfee3d6e6541f76975430c54554562baf3898f0fc6c88)
- access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info

<a id="canonical-d0bd1ce0c97981266a5b8df6395ce6ed04206864709699e0c904c7a8e4c6a969"></a>

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

<a id="canonical-227c335c538a0191c33e3acdf131e0004de5731d5c3d5216d641c2d319c9712e"></a>

## Direct properties — access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info / 81d9d0156c3f / 3

<a id="canonical-1553b328697d64b2c8d9502135680f243c592fbed829196333df63f80365ae1a"></a>

<a id="canonical-f4c5d53a47024329e3b0bebe69c8dbfb25fb01b5438606c17d7f7811426c692d"></a>

## provider_ref property — access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info / 81d9d0156c3f / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-316b62ab3833f139e1f5b63e152db5d01c8fb1ee5e37fc952de9b0a4d7e5803f"></a>

<a id="canonical-7fe741ba2a449034fd76e88366104973e79175b53cbc67a308ca66185bc7d7d8"></a>

## url property — access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info / 81d9d0156c3f / 5

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

<a id="canonical-a606b101c839b0847bb1fe1fe377684036ec27c0469498bb3080e22d648bdd06"></a>

## Next pages — access_info.vault_auth_info.app_role_auth.secret_id.clear_secret_info / 81d9d0156c3f / 6

- [access_info.vault_auth_info.app_role_auth.secret_id](data-sources--secret_management_access--reference--group-001.md#canonical-8ce6fbd2b6d65d0662abfee3d6e6541f76975430c54554562baf3898f0fc6c88)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-a0a97aa2e4da9e47ecc9104e222afad13599ad2ee0c29b3d97ac4085a25161a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78c97ceff0132090940002c8be969c911e33f1bbf6a32305a54fbfaf10948fba"></a>

## access_info.vault_auth_info.token — access_info.vault_auth_info.token / 995abcc33ba2 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a)
- access_info.vault_auth_info.token

<a id="canonical-0a6b89a8e23e911e2ad9195db9e354d6a24496cdac56ac756acd9ffcb0909bf4"></a>

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

<a id="canonical-f99c95321715fbb03e092b4b4b883681dbd62324c02db24195c7cc556aae3a19"></a>

## Direct properties — access_info.vault_auth_info.token / 995abcc33ba2 / 3

- [blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-fa184a49e89f12a4d84c687b24e84f1ec177e8b8e68675a05d9647f5f34ea6c4): complete subsection reference.

- [clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-4e1bebe7682790779e2de88040318edfd6418d7c502e71f040b1c8f14847af17): complete subsection reference.

<a id="canonical-a04733d4538d253cd263a4f5e32630a6939159c1829c76f94f9b7f945b8f42f1"></a>

## Next pages — access_info.vault_auth_info.token / 995abcc33ba2 / 4

- [access_info.vault_auth_info.token.blindfold_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-fa184a49e89f12a4d84c687b24e84f1ec177e8b8e68675a05d9647f5f34ea6c4)
- [access_info.vault_auth_info.token.clear_secret_info](data-sources--secret_management_access--reference--group-002.md#canonical-4e1bebe7682790779e2de88040318edfd6418d7c502e71f040b1c8f14847af17)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-fa184a49e89f12a4d84c687b24e84f1ec177e8b8e68675a05d9647f5f34ea6c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c64d53003a8dfc0be95a97d0c0ea8dc34c8283372dda23ff1352d4cbaf86879e"></a>

## access_info.vault_auth_info.token.blindfold_secret_info — access_info.vault_auth_info.token.blindfold_secret_info / 367b70f6570a / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a)
- [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-002.md#canonical-a0a97aa2e4da9e47ecc9104e222afad13599ad2ee0c29b3d97ac4085a25161a3)
- access_info.vault_auth_info.token.blindfold_secret_info

<a id="canonical-80458f73ec1a42d7b25b10b0763f1a938f3a5dcc52cd20b52f26e0b2de33123a"></a>

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

<a id="canonical-50a2030bff1be589586b2f241203af933a27100d2abef5e4c6bba8dc8d796fa5"></a>

## Direct properties — access_info.vault_auth_info.token.blindfold_secret_info / 367b70f6570a / 3

<a id="canonical-9b4e236ecbb27fdc953fdd6b0934784f06dc1c67339d5c052f1be0f83e266cd1"></a>

<a id="canonical-340eaf62280da894e2c5ec0194379dea818aadee517a56ec46586c0653fdf31a"></a>

## decryption_provider property — access_info.vault_auth_info.token.blindfold_secret_info / 367b70f6570a / 4

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

<a id="canonical-bc8ff4876123da87d37be029aac6db68502ca3fd4d8c84adf8ddd33f1693c77d"></a>

<a id="canonical-9e4358d4360de39f93ea624573ea215f275d19b5648071fbfc371b0c0a45d8a6"></a>

## location property — access_info.vault_auth_info.token.blindfold_secret_info / 367b70f6570a / 5

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

<a id="canonical-548626d85b7161805bb7836cae7d0c6ae5cc5fc1a456c28b61cfa5e6e74f9c0c"></a>

<a id="canonical-ec38c806a1b60664d62d53a7b183500ebf82173ba0adc331399cfa5572e75452"></a>

## store_provider property — access_info.vault_auth_info.token.blindfold_secret_info / 367b70f6570a / 6

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

<a id="canonical-bc9c33446b3c60b8a8ca1d5b9b1c5ab54812642eb3d5662090be0f47fd9a7c06"></a>

## Next pages — access_info.vault_auth_info.token.blindfold_secret_info / 367b70f6570a / 7

- [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-002.md#canonical-a0a97aa2e4da9e47ecc9104e222afad13599ad2ee0c29b3d97ac4085a25161a3)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-4e1bebe7682790779e2de88040318edfd6418d7c502e71f040b1c8f14847af17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9668b97579484fdf016147d7bc1ccbd41579e013a415574a138ad6f1417cd16f"></a>

## access_info.vault_auth_info.token.clear_secret_info — access_info.vault_auth_info.token.clear_secret_info / 72a66e0841ba / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [access_info](data-sources--secret_management_access--reference--group-001.md#canonical-7a900c69f90fcf2acc5a93f0d9cbf231aa3324cb8040d7a180f241926abba21f)
- [access_info.vault_auth_info](data-sources--secret_management_access--reference--group-001.md#canonical-cece745041c07f913191557897066c124d26fc8e4d232ac2c7d60084a364c44a)
- [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-002.md#canonical-a0a97aa2e4da9e47ecc9104e222afad13599ad2ee0c29b3d97ac4085a25161a3)
- access_info.vault_auth_info.token.clear_secret_info

<a id="canonical-7a42ccbba549450011eb623f709762e0f9c4b7cc7c8f8814013859d4e720f6f7"></a>

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

<a id="canonical-5fc1f212272438cb34b1ba7d3c78db2535e8d5c51916a9bd4e7cf82e8fcee297"></a>

## Direct properties — access_info.vault_auth_info.token.clear_secret_info / 72a66e0841ba / 3

<a id="canonical-04cee01e2f7b809fe837761f754ecf4a84711bb30335faaf30dff7b2485d0f55"></a>

<a id="canonical-f0b8028014cab23a0d21da583184aecf34f9eac3d4957607e3dee82b4b5a8381"></a>

## provider_ref property — access_info.vault_auth_info.token.clear_secret_info / 72a66e0841ba / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-9b8e627a9c6d09435fd341172a15968797e430b2f4cf639c845071b34432ae47"></a>

<a id="canonical-c47d5639f2ed9eaab6c45269eeecad39f837d63f8cc2055754e611f6f9d3adfd"></a>

## url property — access_info.vault_auth_info.token.clear_secret_info / 72a66e0841ba / 5

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

<a id="canonical-2f78171133f1c9cd30b1dc862c80bd19a125cd9a661fa3ea92da1d8be8ef0ea2"></a>

## Next pages — access_info.vault_auth_info.token.clear_secret_info / 72a66e0841ba / 6

- [access_info.vault_auth_info.token](data-sources--secret_management_access--reference--group-002.md#canonical-a0a97aa2e4da9e47ecc9104e222afad13599ad2ee0c29b3d97ac4085a25161a3)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dae6a6b49863a708c72728ffccad4042b3ee52350166d1b304da7e4601b56709"></a>

## where — where / 1198562c9790 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- where

<a id="canonical-c20b6c23fffc096f6c342afeafd03fed39b0fb3c9e05d3ce64b3a0aa1ba51533"></a>

Type: `"single"`. Computed.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local..

Upstream description:

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

<a id="canonical-0019593fdfea2980bc637cf524b947912650cf61226f6aa28376f7177a5987c1"></a>

## Direct properties — where / 1198562c9790 / 3

- [site](data-sources--secret_management_access--reference--group-002.md#canonical-a72691d12ea34974f8130503c5c77d5481c502766f41f106548f70c17d11daf9): complete subsection reference.

- [virtual_network](data-sources--secret_management_access--reference--group-002.md#canonical-44436a94516e99c47818a77299bb58211f0809ecde991909399bca507ad54807): complete subsection reference.

- [virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-77da69a85d76ee29e23c71afba443f2021b34cf90b0945e5696061b568013aa6): complete subsection reference.

<a id="canonical-1e8f9a278622bdae387e2d37f016f91d35f72e799aa44443681851517e787493"></a>

## Next pages — where / 1198562c9790 / 4

- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-a72691d12ea34974f8130503c5c77d5481c502766f41f106548f70c17d11daf9)
- [where.virtual_network](data-sources--secret_management_access--reference--group-002.md#canonical-44436a94516e99c47818a77299bb58211f0809ecde991909399bca507ad54807)
- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-77da69a85d76ee29e23c71afba443f2021b34cf90b0945e5696061b568013aa6)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-a72691d12ea34974f8130503c5c77d5481c502766f41f106548f70c17d11daf9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a8fe4a273e7336347c8e0b63e69919f645746980391307ab028f041b54a05ec"></a>

## where.site — where.site / b245d35bc761 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- where.site

<a id="canonical-fec7f24ae73cbd19402e1fafddc17e6a93e4edb40cc4602def484926589d6a3c"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a site configuration object.

Upstream description:

This specifies a direct reference to a site configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-7af3901f8913f1d572d85669794653870ea52993ad091069787332e3d5de67b1"></a>

## Direct properties — where.site / b245d35bc761 / 3

- [disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-085a6b6b1ed7b73958a9a4c9189d9c60d701ff8bbe16e5b3fc7c05fac1a1590d): complete subsection reference.

- [enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-9571f7ed7d46ff8b17380acf464e318ec537efcc5ef32cdc3c5bcf74c9ac3b49): complete subsection reference.

<a id="canonical-a425bc40fb365dfe718b426af5a24e386af61b20d6407e6c09a57455987da4ca"></a>

<a id="canonical-4a90100e9e1769d78c27e3fe0b70370b359002529e5197af71a7f874384d7e56"></a>

## network_type property — where.site / b245d35bc761 / 4

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--secret_management_access--reference--group-002.md#canonical-faafef12bcb814a3a71044bbe04f5f451071de622a9fce6110671b267d673111): complete subsection reference.

<a id="canonical-0e15817aee6fb7f77c620fd9c0f64d52e872e9e83456b0da52e56d0707a8e24f"></a>

## Next pages — where.site / b245d35bc761 / 5

- [where.site.disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-085a6b6b1ed7b73958a9a4c9189d9c60d701ff8bbe16e5b3fc7c05fac1a1590d)
- [where.site.enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-9571f7ed7d46ff8b17380acf464e318ec537efcc5ef32cdc3c5bcf74c9ac3b49)
- [where.site.ref](data-sources--secret_management_access--reference--group-002.md#canonical-faafef12bcb814a3a71044bbe04f5f451071de622a9fce6110671b267d673111)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-085a6b6b1ed7b73958a9a4c9189d9c60d701ff8bbe16e5b3fc7c05fac1a1590d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f68e68ce7d8f7bb3e3016d8c565162468c366c90161984110e9879e13a8e8351"></a>

## where.site.disable_internet_vip — where.site.disable_internet_vip / 2caa0bae0afc / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-a72691d12ea34974f8130503c5c77d5481c502766f41f106548f70c17d11daf9)
- where.site.disable_internet_vip

<a id="canonical-86644d4b38ebba5fad9fd82b33cfe96730283206979dcd801d02526128537cf2"></a>

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

<a id="canonical-08df9ac883c987bd85aa7b43e2ed32125b6558911f0e48f4fa0baa6fef5da0cd"></a>

## Direct properties — where.site.disable_internet_vip / 2caa0bae0afc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c12dcadd34810f213b80c23ef83c314b8517e375aa83b0dd49ee846876ce4ac6"></a>

## Next pages — where.site.disable_internet_vip / 2caa0bae0afc / 4

- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-a72691d12ea34974f8130503c5c77d5481c502766f41f106548f70c17d11daf9)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-9571f7ed7d46ff8b17380acf464e318ec537efcc5ef32cdc3c5bcf74c9ac3b49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eea3ecf31935725d0eadcf33502012d35975fd90156e38cfb637bad36a2f194a"></a>

## where.site.enable_internet_vip — where.site.enable_internet_vip / c0b348d934ec / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-a72691d12ea34974f8130503c5c77d5481c502766f41f106548f70c17d11daf9)
- where.site.enable_internet_vip

<a id="canonical-5569b54acfa7ec67102cda2969b02a0932426d0f816cf080bd1175aa98a69d1b"></a>

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

<a id="canonical-4b151ff81617119b4407b0d34c48054e26c93e50700a8bbbc0874bafeb65a8ea"></a>

## Direct properties — where.site.enable_internet_vip / c0b348d934ec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-584b31272c612728268245aee3c14806e153c0f627f3d424a24f94517c2b0010"></a>

## Next pages — where.site.enable_internet_vip / c0b348d934ec / 4

- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-a72691d12ea34974f8130503c5c77d5481c502766f41f106548f70c17d11daf9)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-faafef12bcb814a3a71044bbe04f5f451071de622a9fce6110671b267d673111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7de743b2f77a7c1a9e650f7a107e0db16422588a82b8f076649a3845f765753"></a>

## where.site.ref — where.site.ref / cd2b243290fd / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-a72691d12ea34974f8130503c5c77d5481c502766f41f106548f70c17d11daf9)
- where.site.ref

<a id="canonical-b82dc783eac399bf1520012a8fec6d35a11fd913cf68b6564c61a93c0d655789"></a>

Type: `"list"`. Computed.

Reference. A site direct reference.

Upstream description:

A site direct reference.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-7b5d7fb55fce7e118e8950a5a62cb133ae6c7aeeb524ad0e79cf490f9264513a"></a>

## Direct properties — where.site.ref / cd2b243290fd / 3

<a id="canonical-eb7a5273efaf3a19514a3744c04e849edf55786a3fc7822c50971d162d172012"></a>

<a id="canonical-b8574304d4a2cfa1cda89f68f5886336b13bbc7dfb91906e36571e319bc9f351"></a>

## kind property — where.site.ref / cd2b243290fd / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-0c405da0ce1c2bd7f70f719b40d16dd0eb74b918a2d566cb4cf0daaff59a20ab"></a>

<a id="canonical-b3893cad1958a22e2eec14156e8a6d16cdadd951259e497cda8762fc649fe019"></a>

## name property — where.site.ref / cd2b243290fd / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-1aafb86024dc8cd81578fc8e305ca917cf3fb65a0f5b4379a73d1cfc75c72c4b"></a>

<a id="canonical-f458e5f8a95b2ac5a5eedf08a3abb4b11d97a5cfab51ee4a3305339f61b21ff2"></a>

## namespace property — where.site.ref / cd2b243290fd / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-5fd7b960d74b6ccea4a83681dcbfd8c8930322e38119b7fea21fd8b3e4c79f32"></a>

<a id="canonical-97e6245e3fbf7957b72173abe638676c06f4f501f939fa30bd3b86892f3d1df4"></a>

## tenant property — where.site.ref / cd2b243290fd / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-e581269cc591de2fcd7bcdd9cea5c3b054a66744d47d0deb7b49a876feb48626"></a>

<a id="canonical-833d8b482bb821a4a2212fc942f1071cd828bb09c251a1f080d280c0b361abc8"></a>

## uid property — where.site.ref / cd2b243290fd / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-ca4c0679911eafcebbe54636cdae0118cd7e28b33972a8cdedbf45945fc386a7"></a>

## Next pages — where.site.ref / cd2b243290fd / 9

- [where.site](data-sources--secret_management_access--reference--group-002.md#canonical-a72691d12ea34974f8130503c5c77d5481c502766f41f106548f70c17d11daf9)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-44436a94516e99c47818a77299bb58211f0809ecde991909399bca507ad54807"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f92c70a99af3d924604bfa13dc98035187387d1092eaf1bb7c7576f2b8e86a7"></a>

## where.virtual_network — where.virtual_network / 6545192eeaaa / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- where.virtual_network

<a id="canonical-3c5ad0e9ebb1682491b37c0c03b1d9848e619dc246ef8ac374037ea1444b0bbc"></a>

Type: `"single"`. Computed.

Specifies a direct reference to a network configuration object.

Upstream description:

This specifies a direct reference to a network configuration object.

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

<a id="canonical-e53830168f98c074088029600310090f08281c4c0173e44b0e46b43f3d96a8b0"></a>

## Direct properties — where.virtual_network / 6545192eeaaa / 3

- [ref](data-sources--secret_management_access--reference--group-002.md#canonical-b9683534602fdee9d79c42ce1b58ca475eca96cc97d8547bdcf3a7386d04f7f0): complete subsection reference.

<a id="canonical-c0fcbda33e0f078addbc6b53780f75fb6ea6b9ab0bbbf77eec5cfaf811bc6275"></a>

## Next pages — where.virtual_network / 6545192eeaaa / 4

- [where.virtual_network.ref](data-sources--secret_management_access--reference--group-002.md#canonical-b9683534602fdee9d79c42ce1b58ca475eca96cc97d8547bdcf3a7386d04f7f0)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-b9683534602fdee9d79c42ce1b58ca475eca96cc97d8547bdcf3a7386d04f7f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a479e210124c106b51aaa5fa3388582a4341dfd765fd8fab5b205179af8fb731"></a>

## where.virtual_network.ref — where.virtual_network.ref / 9d4de4e3bbcd / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- [where.virtual_network](data-sources--secret_management_access--reference--group-002.md#canonical-44436a94516e99c47818a77299bb58211f0809ecde991909399bca507ad54807)
- where.virtual_network.ref

<a id="canonical-5cbe56e11873703a1592ad788a7ef4f1e674affedddbbdd28278bb0fd0921220"></a>

Type: `"list"`. Computed.

Reference. A virtual network direct reference.

Upstream description:

A virtual network direct reference.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-f1ae7743e745b05eb834b8c0f404ba8e78c50d0f414abdbe830163b377ccc0f9"></a>

## Direct properties — where.virtual_network.ref / 9d4de4e3bbcd / 3

<a id="canonical-67006bc5f99913e934e1b8b012b4e039af3023c2c0852575db921c191cbfd6bc"></a>

<a id="canonical-84dff2f437a0d962a5cf43638710a479d402b2977d988d11547598464f042504"></a>

## kind property — where.virtual_network.ref / 9d4de4e3bbcd / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-b47bdbb8fff05ec0db9cefb4d9e69888180c6c29a27d284c863bb03d38bf9a87"></a>

<a id="canonical-2f44f4e0aa3d28554e9fa8cfc2a7e6fad932714ed17f29948c1681043e05b8e2"></a>

## name property — where.virtual_network.ref / 9d4de4e3bbcd / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-4c74e9e157da85f7e8e54c71810ebe0487fc87d79367803713114ec3cfb31f40"></a>

<a id="canonical-6fbc4c326ecf582beefb4a31cd2527f28252798a938a020815f0f6b373dbb95b"></a>

## namespace property — where.virtual_network.ref / 9d4de4e3bbcd / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-04b30ab73411a55a433db3a1a5cced269299e2b92fe5899da71a3251d624ef55"></a>

<a id="canonical-d936e4ca3ec6ca052e44a73e1380ed972aaf909c943e2132ad7760915073cba1"></a>

## tenant property — where.virtual_network.ref / 9d4de4e3bbcd / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-7131ea14e2863efa9eafc94b8c20265ac7a27b5caa0dda3a8a4912f4630529ce"></a>

<a id="canonical-32b59a3fa72adb368cb65b76290e8916c36e49b22ca6b96f6100fab57b2f6917"></a>

## uid property — where.virtual_network.ref / 9d4de4e3bbcd / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-95a10f950e1500db58e7731c64345943e56864f11f11fad5abcaf91d3615ba34"></a>

## Next pages — where.virtual_network.ref / 9d4de4e3bbcd / 9

- [where.virtual_network](data-sources--secret_management_access--reference--group-002.md#canonical-44436a94516e99c47818a77299bb58211f0809ecde991909399bca507ad54807)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-77da69a85d76ee29e23c71afba443f2021b34cf90b0945e5696061b568013aa6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-790577cf564aa32f4e887c6cf81e8981c0d9cda7644064c261a78a107e0e17aa"></a>

## where.virtual_site — where.virtual_site / 26e6766ec8af / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- where.virtual_site

<a id="canonical-813f538ac0ebb50de9172bcb70378f2116784ef4a9c64126b11cf6338ee68562"></a>

Type: `"single"`. Computed.

Virtual Site. A reference to virtual\_site object.

Upstream description:

A reference to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-afed96feda5658c35670f4f2e73396357524fd01cab07d5adf01561dedfa1f0f"></a>

## Direct properties — where.virtual_site / 26e6766ec8af / 3

- [disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-a22e21649e3d48b50168d9cefb00c122e767a969009f2f8009de4af68ef4ed70): complete subsection reference.

- [enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-2d23e11c26a59f47977a27fccebce716fa76c9573c4be0a06c368c24ef4e6fb0): complete subsection reference.

<a id="canonical-7f38b6b5d4098b00202675c4d9337933820a1057b112e1e591ea98d47bcf3809"></a>

<a id="canonical-4fe4488b3f60e4233c51bfa958aae4a53285e2b3311d745be9654c259a072ec8"></a>

## network_type property — where.virtual_site / 26e6766ec8af / 4

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--secret_management_access--reference--group-002.md#canonical-5741bbdaf940b14ae300e032785212569c1ec4042767610c128223eeff880299): complete subsection reference.

<a id="canonical-b56113b51820edeb91d7cb30779fdfe0c801fc7f37dccd79f3936cb86dcaaa0a"></a>

## Next pages — where.virtual_site / 26e6766ec8af / 5

- [where.virtual_site.disable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-a22e21649e3d48b50168d9cefb00c122e767a969009f2f8009de4af68ef4ed70)
- [where.virtual_site.enable_internet_vip](data-sources--secret_management_access--reference--group-002.md#canonical-2d23e11c26a59f47977a27fccebce716fa76c9573c4be0a06c368c24ef4e6fb0)
- [where.virtual_site.ref](data-sources--secret_management_access--reference--group-002.md#canonical-5741bbdaf940b14ae300e032785212569c1ec4042767610c128223eeff880299)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-a22e21649e3d48b50168d9cefb00c122e767a969009f2f8009de4af68ef4ed70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-157acf89660e3805c472dadc06fc1ef269a3e698115f9e0dbaaefa6ab55fe89d"></a>

## where.virtual_site.disable_internet_vip — where.virtual_site.disable_internet_vip / b1835a5d41b2 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-77da69a85d76ee29e23c71afba443f2021b34cf90b0945e5696061b568013aa6)
- where.virtual_site.disable_internet_vip

<a id="canonical-5edeaa0cf231d4cb00eb70ae2960303a85dda2b63fcc54bc1e6710822be66463"></a>

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

<a id="canonical-5b8ebb041b22ba291e0ab263c2114f8d69ebfb34675fd0a94b9dc47eda809c23"></a>

## Direct properties — where.virtual_site.disable_internet_vip / b1835a5d41b2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-423d86502f71bab6b42bb0942a9061fb75ec4b02a74c56e7dc421022e0fcd780"></a>

## Next pages — where.virtual_site.disable_internet_vip / b1835a5d41b2 / 4

- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-77da69a85d76ee29e23c71afba443f2021b34cf90b0945e5696061b568013aa6)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-2d23e11c26a59f47977a27fccebce716fa76c9573c4be0a06c368c24ef4e6fb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89a3b767a83c4649c945c8c082ccdcae12e0e77bdb489c6e6b06e0c1bb20aed9"></a>

## where.virtual_site.enable_internet_vip — where.virtual_site.enable_internet_vip / a95715fe8b41 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-77da69a85d76ee29e23c71afba443f2021b34cf90b0945e5696061b568013aa6)
- where.virtual_site.enable_internet_vip

<a id="canonical-580b2edb06d778c720397c358cb83079f610990b21088ec217cdc5bb31a5fc81"></a>

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

<a id="canonical-28f013eccd1788542e9e10c9b3bc0f2f41192449b226ebc2d7951b80811cdf9f"></a>

## Direct properties — where.virtual_site.enable_internet_vip / a95715fe8b41 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-481bab52648643eaa015cf39505e3b6999e01005778cb33d9934ac56570cafd0"></a>

## Next pages — where.virtual_site.enable_internet_vip / a95715fe8b41 / 4

- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-77da69a85d76ee29e23c71afba443f2021b34cf90b0945e5696061b568013aa6)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)

<a id="canonical-5741bbdaf940b14ae300e032785212569c1ec4042767610c128223eeff880299"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a35543ad8bfd70e9564db22ab5c4513207d1436dfc12be2a05610ff83c72c740"></a>

## where.virtual_site.ref — where.virtual_site.ref / a15074f54972 / 2

Breadcrumbs:

- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
- [Property reference](data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [where](data-sources--secret_management_access--reference--group-002.md#canonical-087b3328a7dd9f4c2fa093cef4c65b95201bded492d218a604e245f78d36f091)
- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-77da69a85d76ee29e23c71afba443f2021b34cf90b0945e5696061b568013aa6)
- where.virtual_site.ref

<a id="canonical-6e340a1d3890b7a12f9485fbce02d6f9089da239a90a5e4d0d9e3a07c02a1b67"></a>

Type: `"list"`. Computed.

Reference. A virtual\_site direct reference.

Upstream description:

A virtual\_site direct reference.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-a621f826cd81a41a375359c0e26e8d375b1cf625859c74b9c2a0a89c1e3d09e2"></a>

## Direct properties — where.virtual_site.ref / a15074f54972 / 3

<a id="canonical-a2b4dc75d694b397b23ace282e6ee7ccf772aa10f7c592d097696ba925311f38"></a>

<a id="canonical-c819f23a3492724aa7ddf01bf82a0b5fc4d04fbae8fbbfcd79ca5211e7763901"></a>

## kind property — where.virtual_site.ref / a15074f54972 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-eeb8c1fb2365d8b6ba9b3836dd10340ce72de04043c66b6a65b5af0886197bfc"></a>

<a id="canonical-4a4f4a1e4f24b52248aa8b961bfed93c8306dbe008a6595ab801ebc1714a9e70"></a>

## name property — where.virtual_site.ref / a15074f54972 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-49b07b3ea46cc8b82f96007af15e1ea11cb616c16561e4aaad0f154db7d1d008"></a>

<a id="canonical-b93ecf1ee9b2797dab9e61001845ad7ad64b1405543b74b5c90714b9c4c2d268"></a>

## namespace property — where.virtual_site.ref / a15074f54972 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2858000a57bb431f3f1aae3ac8f76d441b4f8f4e2bb7b6f0fab8dd624cf01053"></a>

<a id="canonical-38a2419ed33c1202024f63a59671d6c4790eae12ed404b957ed062290af31d92"></a>

## tenant property — where.virtual_site.ref / a15074f54972 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-f52875049ff87f0c88b7d714afba0f1bddcdf7d50f1d7fae32e490074bf3c1d3"></a>

<a id="canonical-a0fe47b73452ea65a0f44da79e037e7e55736e8a88efeacf572a19d953e84134"></a>

## uid property — where.virtual_site.ref / a15074f54972 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-81c9090b745c256e67bed452c1de261162d91be2ecc74ece32d2efe7c17f645d"></a>

## Next pages — where.virtual_site.ref / a15074f54972 / 9

- [where.virtual_site](data-sources--secret_management_access--reference--group-002.md#canonical-77da69a85d76ee29e23c71afba443f2021b34cf90b0945e5696061b568013aa6)
- [xcsh_secret_management_access](../data-sources/secret_management_access.md#canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b)
