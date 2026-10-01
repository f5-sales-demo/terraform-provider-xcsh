---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-e0efd0372ac60300862d8a2f3dd55a78aa5cc5929ba2bc60466cd66f547bbd9a"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains / 7952e14e38dc / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-50f72f4773ff6a6ce1a50b2b5f5d31ea4ecc61547d908e2172a9a4be76e7b9ad)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2e2cda0caea7bc1ba33c40c0fb6d6849df12cfe6286c0d8a328626b2b932cfee)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-50f72f4773ff6a6ce1a50b2b5f5d31ea4ecc61547d908e2172a9a4be76e7b9ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca919fa481aaf421c58b65ea70ce5ef58e4e271c434e8106f398f2f632444cd0"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / b148ae59157c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2e2cda0caea7bc1ba33c40c0fb6d6849df12cfe6286c0d8a328626b2b932cfee)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a287b8b24ee789e11d0ee6fbc62691da44c371378f04ebe43533c35b74d2f249)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-d4e1af83c000e1d471b5db1804a42a3b5886e740a62b1346851ab20dd25e1830"></a>

Type: `"single"`. Computed.

Configuration parameter for simple login.

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

<a id="canonical-546152d2de78df234ed03783301176bb53e31ad0a327a7862cf31b4b3990d00a"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / b148ae59157c / 3

- [password](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-424abe99eb476e1cc6f670c35c84aec761cba036d408b5f95b3114042821060e): complete subsection reference.

<a id="canonical-8b68c33e7750f8c8d64037843f9a004da09f096bbc14fb867d200b821c0ae774"></a>

<a id="canonical-c98d37c173fe2532c4d85e2b8d72b939c261921fa74f1a437a700bc93987e61b"></a>

## user property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / b148ae59157c / 4

Type: `"string"`. Computed.

Enter the username to assign credentials for the selected domain to crawl.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2fc7fd2867372719776dd5aa15ed4f55a999f15e04bb1a8b721947ef19b7100b"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login / b148ae59157c / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-424abe99eb476e1cc6f670c35c84aec761cba036d408b5f95b3114042821060e)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a287b8b24ee789e11d0ee6fbc62691da44c371378f04ebe43533c35b74d2f249)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-424abe99eb476e1cc6f670c35c84aec761cba036d408b5f95b3114042821060e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86e9fca6ba530a6700637c89b652f78be2a86c993e88956757bc75e248566fe7"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 8949c3fd2cab / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2e2cda0caea7bc1ba33c40c0fb6d6849df12cfe6286c0d8a328626b2b932cfee)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a287b8b24ee789e11d0ee6fbc62691da44c371378f04ebe43533c35b74d2f249)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-50f72f4773ff6a6ce1a50b2b5f5d31ea4ecc61547d908e2172a9a4be76e7b9ad)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-7b5062f91a579148a24aab8363c7a465331f904a678372617051d95702116e27"></a>

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

<a id="canonical-b500dd9f26c99c0d44570a8386cc4d56c6b46a829ccad91974280eed160ab8f0"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 8949c3fd2cab / 3

- [blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-e13ceac5316c7a0a2dcc476109ec7d1c6e5b51a2597aaa2e75b8f1372bf404c9): complete subsection reference.

- [clear_secret_info](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-35d5504a33f21c3a49ab1b0c2d952225def72a0507ebf24a4569782a5eba2a0b): complete subsection reference.

<a id="canonical-eac9ba4031b79e0b6c88b3e080bb9afeff81268d36d9e1d7129852ccdb11f901"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / 8949c3fd2cab / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-e13ceac5316c7a0a2dcc476109ec7d1c6e5b51a2597aaa2e75b8f1372bf404c9)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-35d5504a33f21c3a49ab1b0c2d952225def72a0507ebf24a4569782a5eba2a0b)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-50f72f4773ff6a6ce1a50b2b5f5d31ea4ecc61547d908e2172a9a4be76e7b9ad)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e13ceac5316c7a0a2dcc476109ec7d1c6e5b51a2597aaa2e75b8f1372bf404c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fe3af3684a73b8aa76a042037da32f7d73bf749d761a217a714c51a4ff7732d"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / d404f7ab090e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2e2cda0caea7bc1ba33c40c0fb6d6849df12cfe6286c0d8a328626b2b932cfee)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a287b8b24ee789e11d0ee6fbc62691da44c371378f04ebe43533c35b74d2f249)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-50f72f4773ff6a6ce1a50b2b5f5d31ea4ecc61547d908e2172a9a4be76e7b9ad)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-424abe99eb476e1cc6f670c35c84aec761cba036d408b5f95b3114042821060e)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-a6a56550f521049753beb9a8e80887c04f5fdef1bc9771bc6b097ddaaaec0c9f"></a>

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

<a id="canonical-3fad112305a01bc24eddcd96496c43e2c07812a86897ff62f890c86a4c8b07e0"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / d404f7ab090e / 3

<a id="canonical-9fdeef3733447998bdfd179e015f02a9a3d6d486624bccc2d8c9191d3eeda9b3"></a>

<a id="canonical-c4b676d936b91f61a5a6ac5004f82d7658964a6d8309fdf8a35842eec5f90501"></a>

## decryption_provider property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / d404f7ab090e / 4

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

<a id="canonical-c8813454b5d6f45cfa9c17d06d87e7e416a75b44121ed1d76854b329d6eab1c3"></a>

<a id="canonical-f164d133d304c883d84c3643f37d70338d3d0a60eb266993ac96bf28af3362a7"></a>

## location property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / d404f7ab090e / 5

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

<a id="canonical-de4b87557ccfc18418c971e512eb13f9c216bfbf6b9d33c7c90b37aec3696df4"></a>

<a id="canonical-27892fd9e9fa9decb4e6519945d9387bbf2742e61bf6b8185bce219dc6ea088b"></a>

## store_provider property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / d404f7ab090e / 6

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

<a id="canonical-ff57308b426fbdf1cf167c45610a7d38149fe1fee73c95bca9408ba1900d6560"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / d404f7ab090e / 7

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-424abe99eb476e1cc6f670c35c84aec761cba036d408b5f95b3114042821060e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-35d5504a33f21c3a49ab1b0c2d952225def72a0507ebf24a4569782a5eba2a0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8d610be6f94679f363f92ead013ff5d41c91c129f248eb4b8fecf1b257fd7a1"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / c0842b7eda45 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e)
- [enable_api_discovery.api_crawler.api_crawler_config](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2e2cda0caea7bc1ba33c40c0fb6d6849df12cfe6286c0d8a328626b2b932cfee)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-a287b8b24ee789e11d0ee6fbc62691da44c371378f04ebe43533c35b74d2f249)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-50f72f4773ff6a6ce1a50b2b5f5d31ea4ecc61547d908e2172a9a4be76e7b9ad)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-424abe99eb476e1cc6f670c35c84aec761cba036d408b5f95b3114042821060e)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-56c91f4c977dec9f0c53646593b2ee6d445d5cc5d2a12513cebb346af85ee71c"></a>

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

<a id="canonical-763d5506fc5d32747dec58678e520c07e47faa90d0747313ce68707785ec969e"></a>

## Direct properties — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / c0842b7eda45 / 3

<a id="canonical-fe663e86fca40961b3a7ebd3eb6a1ec30bd66044b9ef4584ef7f770f499ab9d0"></a>

<a id="canonical-c55ae560af92e611404ed3dc99799e7e85920951d4a6326de36127d3779bd01b"></a>

## provider_ref property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / c0842b7eda45 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-50250b88efc2cb081bad1eb0ac6b6e5729d2cbb1bf590704fb44ebf5d0c5e9b5"></a>

<a id="canonical-f2f4cad741a141372af11c4f848e10a89a70244f29a11ea493377350456c606d"></a>

## url property — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / c0842b7eda45 / 5

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

<a id="canonical-bc791edc98e0a75c4df0945a490c7ac3917b42d33aae177159d0c9b5ef751a02"></a>

## Next pages — enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.passwor / c0842b7eda45 / 6

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-424abe99eb476e1cc6f670c35c84aec761cba036d408b5f95b3114042821060e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-1cc475c26c933a448b7721dc3dd8eaffd58e151d50676ea5562d8bdce2a17fe7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1bd37f65f8829291c3ddbafd730bf32829fb09e476992b6a725842f3247bc73"></a>

## enable_api_discovery.api_crawler.disable_api_crawler — enable_api_discovery.api_crawler.disable_api_crawler / 934f112f9baf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e)
- enable_api_discovery.api_crawler.disable_api_crawler

<a id="canonical-5811f84d7bb5018805c409a6f250f383c6259b2ce4eaa3f22c59411e7ccc296f"></a>

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

<a id="canonical-2f927324400be18a632585cc8a9fc5c8cd6c9456d415f71d3130a838bc6acd42"></a>

## Direct properties — enable_api_discovery.api_crawler.disable_api_crawler / 934f112f9baf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5af03c792e176e9beab188de39037f760f5cc86f885b2b983488fcf48a914422"></a>

## Next pages — enable_api_discovery.api_crawler.disable_api_crawler / 934f112f9baf / 4

- [enable_api_discovery.api_crawler](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-30189c8b2ce8252f5fa9e3607ae7d997fa67ff651a7fda7375506a95fa928b1e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-dc1a93946173b466cc20982ac5923dd0373375379651dbc3520e9247977388a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5756148f7f110f09d9047ba65923b8bb3f1424987aab9e180c0b13d7d0061c8"></a>

## enable_api_discovery.api_discovery_from_code_scan — enable_api_discovery.api_discovery_from_code_scan / b9559e56bf9b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- enable_api_discovery.api_discovery_from_code_scan

<a id="canonical-965c81ba9c13ab1d7c915fe96da6e9d2ddb6477837796b5d2b854d827f2ce766"></a>

Type: `"single"`. Computed.

Select Code Base and Repositories.

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

<a id="canonical-9b509cc44d369e7c077979b069e348e8d95f5c89eeb3ace3edd18291a73e6b6f"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan / b9559e56bf9b / 3

- [code_base_integrations](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-451745b331093d2b6e7c4e4db0550fe45265f3ec1b45b05d9b278baa1924b992): complete subsection reference.

<a id="canonical-00477817342dea91de7d5e12019f7def92419d368dddd9a22b758d0f2214456e"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan / b9559e56bf9b / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-451745b331093d2b6e7c4e4db0550fe45265f3ec1b45b05d9b278baa1924b992)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-451745b331093d2b6e7c4e4db0550fe45265f3ec1b45b05d9b278baa1924b992"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05695fe209e4acf4b86b176d158dd56614f7e4b42a71f1d070844067402a73aa"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / b4206932abfc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-dc1a93946173b466cc20982ac5923dd0373375379651dbc3520e9247977388a0)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-faa89fb4fced15dc289a355b0e9030e98b864162046f472bd641cec7ad71c724"></a>

Type: `"list"`. Computed.

Configuration parameter for code base integrations.

Upstream description:

Configuration parameter for code base integrations

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-941f727a39b3f313977403e003dba0a085c6571e72d5a209073717e2231792ca"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / b4206932abfc / 3

- [all_repos](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-4087437360874ebbbf0125a5ca9aa8ca6fa59bea14ade774e0f87dd04cfddc1c): complete subsection reference.

- [code_base_integration](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-7cf51f302651ebafbc326e32cf7420561dde1ed87f82578a5330fdd703b82bd5): complete subsection reference.

- [selected_repos](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3025f5ae9076ab7740cb36ca243d43d685434ccff8ff7e104d82387343117d11): complete subsection reference.

<a id="canonical-ed077160f1ee4fac4b014d0bbbc713083cc6bcbd56391b8e30d3ef3a24f5b09f"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations / b4206932abfc / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-4087437360874ebbbf0125a5ca9aa8ca6fa59bea14ade774e0f87dd04cfddc1c)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-7cf51f302651ebafbc326e32cf7420561dde1ed87f82578a5330fdd703b82bd5)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3025f5ae9076ab7740cb36ca243d43d685434ccff8ff7e104d82387343117d11)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-dc1a93946173b466cc20982ac5923dd0373375379651dbc3520e9247977388a0)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-4087437360874ebbbf0125a5ca9aa8ca6fa59bea14ade774e0f87dd04cfddc1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c78783dfcd419da63e88e0314ff35e10fc3738912d7f55b786fcd2ab46c40717"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 0262d9cd8805 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-dc1a93946173b466cc20982ac5923dd0373375379651dbc3520e9247977388a0)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-451745b331093d2b6e7c4e4db0550fe45265f3ec1b45b05d9b278baa1924b992)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-ff8fde07a55c0e09d7661e308923273956fadac866dca1bcdbfc8981237faa50"></a>

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

<a id="canonical-e1d2a9bd2949c746c39ee34815caffc1f1bee202722570e3c18da56663dff84c"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 0262d9cd8805 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8884ee3cb39070d8bd6e76e7e8504dc6160364b0b3770d1b9548dba0ecb7c494"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_rep / 0262d9cd8805 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-451745b331093d2b6e7c4e4db0550fe45265f3ec1b45b05d9b278baa1924b992)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-7cf51f302651ebafbc326e32cf7420561dde1ed87f82578a5330fdd703b82bd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2cc7091b912e13b3a1bcbaf4e8dde3fbc90a6bde3f9e4582f6030e4b579a91d"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / b95d142b011e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-dc1a93946173b466cc20982ac5923dd0373375379651dbc3520e9247977388a0)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-451745b331093d2b6e7c4e4db0550fe45265f3ec1b45b05d9b278baa1924b992)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-a0cdec1a074747fa205f592ebbfc54e4ad385a22b5ef2606f3d94031e03ff822"></a>

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

<a id="canonical-8d0c7a8004413019119cc68d59b36e824c8c979cfb1e7911de1238508c01de5b"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / b95d142b011e / 3

<a id="canonical-dd6a20f22eb2918cfc5059dd13a363e396ee0070819203aac3aa7f623825ede3"></a>

<a id="canonical-2bb5a4cc15b16695acc0f6748a7da355f690e12d07773497fce6cd766f10a126"></a>

## name property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / b95d142b011e / 4

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

<a id="canonical-e3c1445cda909b7234c080d8acc581a7756716463fe5d2f5887a10fd01850e91"></a>

<a id="canonical-34d1947fc89b1c9c5dd533203036ad8b29b70c90a9f91f996d457d32716f2f35"></a>

## namespace property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / b95d142b011e / 5

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

<a id="canonical-2bd06d8c1d36a5c15ac2713235dcab79aedbe2262573dade1c320a5aeb61769e"></a>

<a id="canonical-fb08cc4df00f79d0d1b211dd45f13ab6abc994156194db4465e8ca8a57bde32f"></a>

## tenant property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / b95d142b011e / 6

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

<a id="canonical-f207809c981ffd9ec88670984eddf8348297f94423487b95fbbf16c16d2a734d"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_ba / b95d142b011e / 7

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-451745b331093d2b6e7c4e4db0550fe45265f3ec1b45b05d9b278baa1924b992)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3025f5ae9076ab7740cb36ca243d43d685434ccff8ff7e104d82387343117d11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16f12e986fcd140b604b7ea618ead1f6eff890c29d33b72a2423bc7bb75b5fb3"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 36c9eb5f5c35 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-dc1a93946173b466cc20982ac5923dd0373375379651dbc3520e9247977388a0)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-451745b331093d2b6e7c4e4db0550fe45265f3ec1b45b05d9b278baa1924b992)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-042e228a07dba8ac88126fa0dc26bf44ec47e7b74ac18ee6b72bb8be4ad8a376"></a>

Type: `"single"`. Computed.

Select which API repositories represent the LB applications.

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

<a id="canonical-4b1687800c6076668c2430e223b62a7ad333d31330da3c6ba0fbf0526949e352"></a>

## Direct properties — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 36c9eb5f5c35 / 3

<a id="canonical-43dcbc90ea42d681b644d4bd28f87c75c06a9bdc8b575c7bca91c2dc53d6168f"></a>

<a id="canonical-19d657798b43579869eb3f23548ca39d4a9462dee6949c94b21d7d1ea73db5b8"></a>

## api_code_repo property — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 36c9eb5f5c35 / 4

Type: `["list", "string"]`. Computed.

Code repository which contain API endpoints.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-38e7872c74aa16836a8e595d2059399089e2490e5c4e775cebb900ce157ffe78"></a>

## Next pages — enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selecte / 36c9eb5f5c35 / 5

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-451745b331093d2b6e7c4e4db0550fe45265f3ec1b45b05d9b278baa1924b992)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d8cea2c986d75f0f461fc48f5299fafbc7f28ef79c4c8fd12b7a5863cda65736"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-709add3368a465714208bbb511fd409ef2cbcf6a256480d8bfbb8269de2b2f97"></a>

## enable_api_discovery.custom_api_auth_discovery — enable_api_discovery.custom_api_auth_discovery / e4ca52ded42a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-40f46cd7be32f654f7d2960a7f4e28b025a57d3dbca5d246831ca0b71829dbef"></a>

Type: `"single"`. Computed.

API Discovery Advanced Settings. API Discovery Advanced settings.

Upstream description:

API Discovery Advanced settings.

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

<a id="canonical-4d61b51b53b3e5de3ccdff660ab4c2bd1569dea6e98d73ff6af19cb4cbe3be9d"></a>

## Direct properties — enable_api_discovery.custom_api_auth_discovery / e4ca52ded42a / 3

- [api_discovery_ref](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-dab962ce0c561482132c2910cdbef39de792dfa2f09e363a57d4b8b56b1f88ea): complete subsection reference.

<a id="canonical-d3c3ccb9e40858f376d159b769057eefcb300b3bad1a5cf2a9f41daf5d6eae2c"></a>

## Next pages — enable_api_discovery.custom_api_auth_discovery / e4ca52ded42a / 4

- [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-dab962ce0c561482132c2910cdbef39de792dfa2f09e363a57d4b8b56b1f88ea)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-dab962ce0c561482132c2910cdbef39de792dfa2f09e363a57d4b8b56b1f88ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b007d874ae0be206817dbccc49c0fb5bc61292fe7cecfae70cb8e14ce5ad3976"></a>

## enable_api_discovery.custom_api_auth_discovery.api_discovery_ref — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 4421eeb6d23d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [enable_api_discovery.custom_api_auth_discovery](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d8cea2c986d75f0f461fc48f5299fafbc7f28ef79c4c8fd12b7a5863cda65736)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-5c9faaf5abf4765d3b552a89661f79db29b6d0c7cf15ab06016929aaaf9116a9"></a>

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

<a id="canonical-b6711f7107eac507083cf8c42e80b6f4c7b761342bd5d48dbde29c1aa89502e0"></a>

## Direct properties — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 4421eeb6d23d / 3

<a id="canonical-58f47ee3bb501d1c0a5d7663456c7f681f9a308647914f99c0d6767b518e8738"></a>

<a id="canonical-af06c6e4d0a75d0c7f63a7c1907021f46c44ace9436527e664e6d8c73785c077"></a>

## name property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 4421eeb6d23d / 4

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

<a id="canonical-82921dc12ba1c1d897e56abcd4e3a6b37d30a736c7847ce6743ad43c0c4d13d2"></a>

<a id="canonical-22119a9457a378bc6c1392e31a512eef95187c60c3c02b6eef630c70745f74db"></a>

## namespace property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 4421eeb6d23d / 5

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

<a id="canonical-97d3d675a403fbf356d434619f0920c13f430fdfc08402b44c776d618847a2eb"></a>

<a id="canonical-cffda2582ab47e5756826b6518ffed2a031d3766e15bec27ff69caf75fb9774b"></a>

## tenant property — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 4421eeb6d23d / 6

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

<a id="canonical-8d8c5de6c6044397ad841d7c5692cb3e1ac403ec5223fb22452fdc949ed98151"></a>

## Next pages — enable_api_discovery.custom_api_auth_discovery.api_discovery_ref / 4421eeb6d23d / 7

- [enable_api_discovery.custom_api_auth_discovery](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d8cea2c986d75f0f461fc48f5299fafbc7f28ef79c4c8fd12b7a5863cda65736)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-924655935a29713b71d9683ada5949479dc9e8a942b2280247cfdf4c3e580b83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c77e5949c77877293752360f55fbc863c3b4268e0f63a39968b85a928c5d967"></a>

## enable_api_discovery.default_api_auth_discovery — enable_api_discovery.default_api_auth_discovery / b278fa14bb2a / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-973673811dbe5e96d4150f89a1e2d509192e0acfc2023f417ad3fd9f035c3489"></a>

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

<a id="canonical-1cdfb559fcbb800d3d33ed3b2646dcfb8c8bea7793497d32b53c86be28f08e36"></a>

## Direct properties — enable_api_discovery.default_api_auth_discovery / b278fa14bb2a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e87afed5e7ada907326758279cdf7d3c1b5af89939f4cdac27725853ebeb98c"></a>

## Next pages — enable_api_discovery.default_api_auth_discovery / b278fa14bb2a / 4

- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d0ed52e1a86221b80749c741711d70dbebdb00aa15bccae4614bd04e4df057a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f356607fa13cfddf7667b68512af49e98f494d23719f196878e378924aea327b"></a>

## enable_api_discovery.disable_learn_from_redirect_traffic — enable_api_discovery.disable_learn_from_redirect_traffic / b5baca85331b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-57603f8c405e1aa87f15fd8fc7da4b829058c44aeb1725f61c86e048883e0ffa"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable learn from redirect traffic.

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

<a id="canonical-9073590f836c42bf056182140a812706e03cf2cb71ab20ccd44fe4ad0dedab3a"></a>

## Direct properties — enable_api_discovery.disable_learn_from_redirect_traffic / b5baca85331b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e3a0a95a005045e877719750258ef04e02f81e3f692fe629ef1a84d0a4563b64"></a>

## Next pages — enable_api_discovery.disable_learn_from_redirect_traffic / b5baca85331b / 4

- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-b22017f4a2fc2d6e6311719ed49308077d507d37fc48b53cd22e86f5c24ba279"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ac5e3281a2768ac2ec8f2362101e7c8d6c7ff7fc8f9d3cf145a90dfc83574b1"></a>

## enable_api_discovery.discovered_api_settings — enable_api_discovery.discovered_api_settings / 0ddc35799405 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- enable_api_discovery.discovered_api_settings

<a id="canonical-7f330e7f34afe326f74104e920608f34ca5e6309f88005a72a5e1b9a4af05c3d"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

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

<a id="canonical-9b12e460bb3b9335e66da253d8a33357c7b14226fda7aeabb5f73bacb71b907d"></a>

## Direct properties — enable_api_discovery.discovered_api_settings / 0ddc35799405 / 3

<a id="canonical-add34a74e9b0731c15cb5add6c8f93a20752d9feeeb0c236fe6b62cf6dea7ec0"></a>

<a id="canonical-045d302a5267fcce0b655e39a1ee2d82b2932c1836ad2c9b212d38ed6545c867"></a>

## purge_duration_for_inactive_discovered_apis property — enable_api_discovery.discovered_api_settings / 0ddc35799405 / 4

Type: `"number"`. Computed.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
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
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-bf9dd869634b21fa56d2ffe1fd3755d2699371e3a0f226f7f5a16ef4681060c3"></a>

## Next pages — enable_api_discovery.discovered_api_settings / 0ddc35799405 / 5

- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-7c48c0b9fc1142de5fa4f1a82551b95191251e7c305bcd8950ef84b0d4709770"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3618bd98c259a42ab9836a6fc629febf8b9cee976c5053ee44c7275167b791a"></a>

## enable_api_discovery.enable_learn_from_redirect_traffic — enable_api_discovery.enable_learn_from_redirect_traffic / 62e685890103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-362cdd7e40aa9ed314b75df965a26d611c84439005f6340264a66b32ea58e4a5"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable learn from redirect traffic.

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

<a id="canonical-70a394638f429a2a5e6c332c266a7f88fd2dff8136dbcf8480f47d9c5df90b13"></a>

## Direct properties — enable_api_discovery.enable_learn_from_redirect_traffic / 62e685890103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03c6af3e9065f780d10929d0e61ca9278d126b1361e3cca7d3c68d5bc3f1b80c"></a>

## Next pages — enable_api_discovery.enable_learn_from_redirect_traffic / 62e685890103 / 4

- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-4d27c391b97ec0bb8400170e9e90720495a006d7b956f6f90243e13bfac2d267)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d544109407ea9e1f60773cffc2af2a175cb0e9bc1a2d48c5ed47d4a622291bca"></a>

## enable_challenge — enable_challenge / 04fcf79b6408 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- enable_challenge

<a id="canonical-263f04a4ba704eaa8c8b05ed55cf0dff4eb594b01121543ffc20f9307a8e293b"></a>

Type: `"single"`. Computed.

Configure auto mitigation i.e risk based challenges for malicious users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

<a id="canonical-d3efe072ec11774a7e55888cf64fd3cd4146569d50a5a247ecc22aaef17c46e5"></a>

## Direct properties — enable_challenge / 04fcf79b6408 / 3

- [captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-5cf48b86f6dd90116298dec0a64e96653fac2ac953a65d6c9f7775b8ab913890): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-268da02d476233bde245864dd95844cdbbbe6973ec72eeaace49917f74897f6e): complete subsection reference.

- [default_js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-458e77cae4047bb630c8ca4f42eba1f7454a525e86ec0aeb456ba54598e10f49): complete subsection reference.

- [default_mitigation_settings](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3ecdc6bdb85ecfdad84bc5c6f2c18ac36dda153bd95cdc132a13b5c7fb955f3a): complete subsection reference.

- [js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-73ee6145db782a97dfa3b43a9d31f48866eef2e5c1c410a25ef01a143ce6c370): complete subsection reference.

- [malicious_user_mitigation](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-74b09a0c8db6d1b443b7c2216e67cc04aa641e61fef298fb11c41ebffe07e3bd): complete subsection reference.

<a id="canonical-5f6f4246e17846c2fe5cb197af25b4f0852c5150904551380d03482e2a18d6ea"></a>

## Next pages — enable_challenge / 04fcf79b6408 / 4

- [enable_challenge.captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-5cf48b86f6dd90116298dec0a64e96653fac2ac953a65d6c9f7775b8ab913890)
- [enable_challenge.default_captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-268da02d476233bde245864dd95844cdbbbe6973ec72eeaace49917f74897f6e)
- [enable_challenge.default_js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-458e77cae4047bb630c8ca4f42eba1f7454a525e86ec0aeb456ba54598e10f49)
- [enable_challenge.default_mitigation_settings](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3ecdc6bdb85ecfdad84bc5c6f2c18ac36dda153bd95cdc132a13b5c7fb955f3a)
- [enable_challenge.js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-73ee6145db782a97dfa3b43a9d31f48866eef2e5c1c410a25ef01a143ce6c370)
- [enable_challenge.malicious_user_mitigation](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-74b09a0c8db6d1b443b7c2216e67cc04aa641e61fef298fb11c41ebffe07e3bd)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-5cf48b86f6dd90116298dec0a64e96653fac2ac953a65d6c9f7775b8ab913890"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc23eacb46137a3d44fe6cf46eb9e2131dfdcd38769e17981cc856b3a178047d"></a>

## enable_challenge.captcha_challenge_parameters — enable_challenge.captcha_challenge_parameters / 6adfe7763765 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-8b1303fb7c4918afaef4378e97e832ae6199e33dbd79cb3f56c60485c82d2b13"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google
Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed
to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will
redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

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

<a id="canonical-1f510e7f671d7739f09d999adeab2d3b2506b999bc2f281f47577ed4b335f8fc"></a>

## Direct properties — enable_challenge.captcha_challenge_parameters / 6adfe7763765 / 3

<a id="canonical-495266c5db37b9fe429e00936a69b36fa23b1f2881915dd82402cf49e5ac83e0"></a>

<a id="canonical-bf3bf977a86332bad510dddae4d686ea486db08383998e5e555c69454ce08946"></a>

## cookie_expiry property — enable_challenge.captcha_challenge_parameters / 6adfe7763765 / 4

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

<a id="canonical-163631089753109fe6be16bda98a0e8bce2bc52cdcf8c8bc3d311c56aa9d0d59"></a>

<a id="canonical-b230441d3f66cbd37c11cf18909da74160893e61d3b07b803584c35bd5768326"></a>

## custom_page property — enable_challenge.captcha_challenge_parameters / 6adfe7763765 / 5

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

<a id="canonical-78ccc9af80d8b120dbbf3788c11afa6bc48c3bac087cfd8d699458d8013c1046"></a>

## Next pages — enable_challenge.captcha_challenge_parameters / 6adfe7763765 / 6

- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-268da02d476233bde245864dd95844cdbbbe6973ec72eeaace49917f74897f6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f5f0e83d080f8402bbda0564dc3dd5b6d5cfcfb27d04c9b846921bc7d87c721"></a>

## enable_challenge.default_captcha_challenge_parameters — enable_challenge.default_captcha_challenge_parameters / b42bda2dad5b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-6bc06583dc3dd86be3b5b973cbdcbd35eda2f2b2d8cfdad50199666b3d3380be"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default captcha challenge parameters.

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

<a id="canonical-06bb642b1181435f99ddaf4855cd8328d4c9b9a9b595900c7c25b391172b0a0f"></a>

## Direct properties — enable_challenge.default_captcha_challenge_parameters / b42bda2dad5b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-674e6c2b207ca1783f479dee447f39ad86cae49e0fc537ce718b9283cb6164a3"></a>

## Next pages — enable_challenge.default_captcha_challenge_parameters / b42bda2dad5b / 4

- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-458e77cae4047bb630c8ca4f42eba1f7454a525e86ec0aeb456ba54598e10f49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f39cfa553f929d89f95ca79945ebed99bfe0e9f48c2ceec0349250c449a48290"></a>

## enable_challenge.default_js_challenge_parameters — enable_challenge.default_js_challenge_parameters / 55ee7995a92f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-f8349088b1513d7de4a0771312233c6e31762c37c66c0d9a8a75a7c234d45ac5"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default js challenge parameters.

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

<a id="canonical-4942de9563f892793d271cc3c0944d6936682e954f48d3fb3fe854844f28b941"></a>

## Direct properties — enable_challenge.default_js_challenge_parameters / 55ee7995a92f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-11c2b427ad640df6db23073fc8eeeee0a651975b711cb8b87793688007b7ac7d"></a>

## Next pages — enable_challenge.default_js_challenge_parameters / 55ee7995a92f / 4

- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3ecdc6bdb85ecfdad84bc5c6f2c18ac36dda153bd95cdc132a13b5c7fb955f3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc946b6f62d71155076b1be785cc064a02739ccf29707a0e12ea8af83e735a9a"></a>

## enable_challenge.default_mitigation_settings — enable_challenge.default_mitigation_settings / 78c3cf03ddd1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- enable_challenge.default_mitigation_settings

<a id="canonical-44290795ad301750aa06eab6ede43857a071c92079df3b2b352a5d25295384cb"></a>

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

<a id="canonical-509c8ea3e673d05f64802ccc7594c79b7ab9f54469d281c9ad453e74c28ce154"></a>

## Direct properties — enable_challenge.default_mitigation_settings / 78c3cf03ddd1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2854cb3a6ba925bf9ea601fb6dcdeb799dea3e050fa2b95c0f5edb082ac70327"></a>

## Next pages — enable_challenge.default_mitigation_settings / 78c3cf03ddd1 / 4

- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-73ee6145db782a97dfa3b43a9d31f48866eef2e5c1c410a25ef01a143ce6c370"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ffbf717c1c8f6e526979c56770b16ebe56b2448e84fa9b080766e68b5f2e735"></a>

## enable_challenge.js_challenge_parameters — enable_challenge.js_challenge_parameters / 19e137dcf355 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- enable_challenge.js_challenge_parameters

<a id="canonical-41c42976bac0f22a6d59bd3a8f587a56ba977a0dc06b17aa81cf695fea4f120d"></a>

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

<a id="canonical-31f4f53b280220308f7a2eaa8b01ac9f79ec878063e9caa5f0792849c4204218"></a>

## Direct properties — enable_challenge.js_challenge_parameters / 19e137dcf355 / 3

<a id="canonical-81752b95c03de25b8fe03ead7e506b26f3cf7dd0b0f37f76c43a08622d05de76"></a>

<a id="canonical-1fcb5ab798e60b2e6cbf6ba9368207c55775ff396a05920cba0b6077037926b0"></a>

## cookie_expiry property — enable_challenge.js_challenge_parameters / 19e137dcf355 / 4

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

<a id="canonical-20c09f97502816416a8045ceddd26a67fcf1fd12c19110586ad3a2ca0446eb82"></a>

<a id="canonical-9bd710f41f51179e9143b1d3fa791c9eb8e502a33af604075fc26e749e457963"></a>

## custom_page property — enable_challenge.js_challenge_parameters / 19e137dcf355 / 5

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

<a id="canonical-6b8a0199e2e63cb20f0c9f569c5d82cbf6724e7c5c8f22a0ac2d99cddd5fd91e"></a>

<a id="canonical-bde9e875259d5010ff98cee47676245fa8f03ac8751b85a0ffe524456522f581"></a>

## js_script_delay property — enable_challenge.js_challenge_parameters / 19e137dcf355 / 6

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

<a id="canonical-e670ccdd4e23bf23bffe42cc8ff376299ccb7d6b8acc14991111616f4489a05a"></a>

## Next pages — enable_challenge.js_challenge_parameters / 19e137dcf355 / 7

- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-74b09a0c8db6d1b443b7c2216e67cc04aa641e61fef298fb11c41ebffe07e3bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ca51a0f2ac9f67f001cb45412f7a27df9166740525f361784673d7d7d0d2cec"></a>

## enable_challenge.malicious_user_mitigation — enable_challenge.malicious_user_mitigation / 6c6bf52642a4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- enable_challenge.malicious_user_mitigation

<a id="canonical-6c965353e22460e02fda2d00d4a103212cce7efb2b8a98a07328aff95b1de4bf"></a>

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

<a id="canonical-d280aee43c6d6478dd9199ec38bc381b0f69cf9f5532a583d5a2bbbf2b276ec3"></a>

## Direct properties — enable_challenge.malicious_user_mitigation / 6c6bf52642a4 / 3

<a id="canonical-4362af6a77d490422e2c0f9f0276102d3a348320d36b1e8d480217568e12ee1b"></a>

<a id="canonical-c10344c36355dd9c0357aefb791c0738388a00d52c9aed2cd09325c5cd0039a9"></a>

## name property — enable_challenge.malicious_user_mitigation / 6c6bf52642a4 / 4

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

<a id="canonical-79ebdb73ef1b48be8dab2ec22d035f1b6b784ba0cc983782796941b92ede58b6"></a>

<a id="canonical-be86042adc3af22695e10a8634b63dcdf24c7768f993b775b9a62bfc37167266"></a>

## namespace property — enable_challenge.malicious_user_mitigation / 6c6bf52642a4 / 5

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

<a id="canonical-4cf30b508cf8156fa67e3f4632a65c1ed176ff6b2818c67adb296aaa51adc3e9"></a>

<a id="canonical-182cd163add5aa58fa945645bfbf54da2d52ac799644ebc296392b61925e1a8d"></a>

## tenant property — enable_challenge.malicious_user_mitigation / 6c6bf52642a4 / 6

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

<a id="canonical-cdbc84681b46c24b14414958cdb866ad7b5b669d27d1f978773b1f97a0f71e57"></a>

## Next pages — enable_challenge.malicious_user_mitigation / 6c6bf52642a4 / 7

- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d08c9244a176c7a71959859b471fc4f32017bb25c0f7babbdfcb3dc9f911841a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f791575cc7afb803eae1edab6bd8161ce0ceef15091e48b7ee7c3fa6b3ba4eb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-738276e05ba546ef400aa677aa8f58dde5768bd9ca7b5c05f0b05fffafba27c7"></a>

## enable_ip_reputation — enable_ip_reputation / a24cc40b4905 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- enable_ip_reputation

<a id="canonical-06bc599ba9a0a22102b4b7d3972f4defd6f1b03c273193df0423bf48581b56a0"></a>

Type: `"single"`. Computed.

IP Threat Category List. List of IP threat categories.

Upstream description:

List of IP threat categories.

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

<a id="canonical-e69cf9545dfc5a878cbb3b5abe6f828ecc91e01d229f9bc1b8879fe1ca7341bb"></a>

## Direct properties — enable_ip_reputation / a24cc40b4905 / 3

<a id="canonical-f9c2d06333d2833a81f2ff67625c6fcf021b8313be3c67c12f439666f5ca7d20"></a>

<a id="canonical-c8a2c9f0840d69dd1fe75f8e1e77a0c13842c88d5cd2b919a76954992da9f052"></a>

## ip_threat_categories property — enable_ip_reputation / a24cc40b4905 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

Upstream description:

If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6d077a9a6be89d5517ce2f2a8f2ed6a889a560d79bff6db5dc2a652832d86395"></a>

## Next pages — enable_ip_reputation / a24cc40b4905 / 5

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-63dc44aae6483b2cb22e58f32ac31fd9657588bb024fedd084e3781a61ba5de3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a5855dbd01cc9c896d73d40c64800ed9cc1cffba57dce9391656e0c292b9375"></a>

## enable_malicious_user_detection — enable_malicious_user_detection / 828b1fa57bf6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- enable_malicious_user_detection

<a id="canonical-3220cb9e442f7fa713187eebb8c8d943e1a28aba7d54521d79cbd7c47a3b7757"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable malicious user detection.

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

<a id="canonical-fd03d98f9c563abaaed8e6becb444d4716fe4f4bfc96aff04fbe9530923cb820"></a>

## Direct properties — enable_malicious_user_detection / 828b1fa57bf6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-73603e7a482a5ca983ca4671d3b7f8ba0655bec52fbe230c28cd17838b231ace"></a>

## Next pages — enable_malicious_user_detection / 828b1fa57bf6 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f09d16a1eeb59b3ba5721dc4983536e2eaf4623b14596e4f6b810709ffc6ae2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45e0aaf4c9ec6e55d58907643074cfbf3dc03b2cba5fe0823393dd59e8a257a0"></a>

## enable_threat_mesh — enable_threat_mesh / d9dc34f5f180 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- enable_threat_mesh

<a id="canonical-417e345f3eb73b352e523c3aa18328b3c0bb6784884972f05b2fa28939cca1c6"></a>

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

<a id="canonical-42e3941e9133e2ed92e8cd7c63bf9720302994d016ace85b9b63b73e08dd1931"></a>

## Direct properties — enable_threat_mesh / d9dc34f5f180 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-969589676fe4cb46b8506a37bea8cf7cbd1a34e3446e88a63f4755b60f767e50"></a>

## Next pages — enable_threat_mesh / d9dc34f5f180 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85090639ec28dc9341c937c20f63d7b851e563cd2e8dbf5b5d7ebbf846464e40"></a>

## graphql_rules — graphql_rules / fd9f7eba17c0 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- graphql_rules

<a id="canonical-c2fd7e162c22449b0662bdf897bb55e66ef3ad48efd7afc32d5e092200a46468"></a>

Type: `"list"`. Computed.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy..

Upstream description:

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-f2be37bbbcb0bd27c09cab8ce9fdf1e4a0d4d25a808d19a570f77d9af1a64b31"></a>

## Direct properties — graphql_rules / fd9f7eba17c0 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-da6fc804e924db228aaef5723d12be6fe577256023408660d11de61436b865e9): complete subsection reference.

<a id="canonical-8ff1af12c1ff70415157b211918d8ae5f76de9c114be87c599ea50a0f7375aba"></a>

<a id="canonical-b97d5390d5a7e5c7b05169cf809872cd05a1a90194f71eb898f021fd144fd192"></a>

## exact_path property — graphql_rules / fd9f7eba17c0 / 4

Type: `"string"`. Computed.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Upstream description:

Specifies the exact path to GraphQL endpoint. Default value is /graphql.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-4f71e55c5d49d909a9b39749cf13a4fd71fa5eb7d5bacd550bf165e5fe2c84a1"></a>

<a id="canonical-1aafbb3bf25246a3653b486933270a2dfb59e6eec87b1600c64585d9120e3365"></a>

## exact_value property — graphql_rules / fd9f7eba17c0 / 5

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

- [graphql_settings](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1f3509297bdca674cf4b43e723e9706eb4edd59c10358678e036485802218005): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-cdf9f20b46addf38e7e525b70b8d2028e41754120c00b8a4abcf78be684af34d): complete subsection reference.

- [method_get](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d9eb7c82005fa53260120d72e26ea410f19dea4038a2383d2a22ae4a69baecc8): complete subsection reference.

- [method_post](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-8f703a3f2906f2bccb4bd9d1fef06429016a5cb680ca451f9d0c2e6d9163cee7): complete subsection reference.

<a id="canonical-f8c06084de20aa1747edb49a40c6625995aed8bf68226e9a960e3b1b109b7f71"></a>

<a id="canonical-8555f5b3f05321f43c33e50a3a302adfd5fca1c1c4ab0102f5bcb5f01b2f4b9e"></a>

## suffix_value property — graphql_rules / fd9f7eba17c0 / 6

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

<a id="canonical-ee190dec9bb13280f5a96d51d80c321b1795c4c00c2fb1fb159f6547614a67fa"></a>

## Next pages — graphql_rules / fd9f7eba17c0 / 7

- [graphql_rules.any_domain](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-da6fc804e924db228aaef5723d12be6fe577256023408660d11de61436b865e9)
- [graphql_rules.graphql_settings](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1f3509297bdca674cf4b43e723e9706eb4edd59c10358678e036485802218005)
- [graphql_rules.metadata](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-cdf9f20b46addf38e7e525b70b8d2028e41754120c00b8a4abcf78be684af34d)
- [graphql_rules.method_get](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-d9eb7c82005fa53260120d72e26ea410f19dea4038a2383d2a22ae4a69baecc8)
- [graphql_rules.method_post](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-8f703a3f2906f2bccb4bd9d1fef06429016a5cb680ca451f9d0c2e6d9163cee7)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-da6fc804e924db228aaef5723d12be6fe577256023408660d11de61436b865e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f57602f26b271bddd9d645b0f89589564a606a4aac5588a4d06741538e9f15c6"></a>

## graphql_rules.any_domain — graphql_rules.any_domain / aa26f3601fda / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- graphql_rules.any_domain

<a id="canonical-778d458f28b3702d37b135c132bec69d8ee5947dee50029c74a2594a03ed92f7"></a>

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

<a id="canonical-c712ac9351c42f2f52690d4c59c4fc2a5506e3059bf56432fdcc78ceda089096"></a>

## Direct properties — graphql_rules.any_domain / aa26f3601fda / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1b7e65a2b62ba9f505871a38bf0a8637d305051a2adf85651b49cf3d64d219e4"></a>

## Next pages — graphql_rules.any_domain / aa26f3601fda / 4

- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-1f3509297bdca674cf4b43e723e9706eb4edd59c10358678e036485802218005"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d9e94744faf0f06983176951cbff35c47a197e9a760c3cc52b2b88564866187"></a>

## graphql_rules.graphql_settings — graphql_rules.graphql_settings / 23c0ee643232 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- graphql_rules.graphql_settings

<a id="canonical-646e38e61cea41ef83059e721fce020603c1ce42362f88f4c35a751fa84267a8"></a>

Type: `"single"`. Computed.

Configuration parameter for graphql settings.

Upstream description:

GraphQL configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allow_introspection_queries_choice": "[\"disable_introspection\",\"enable_introspection\"]"
}
```

<a id="canonical-a9ed43cf042010ac44f04349211f2aad20bb69e7cfe10cc78f842c6741823a83"></a>

## Direct properties — graphql_rules.graphql_settings / 23c0ee643232 / 3

- [disable_introspection](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-524534144f5c72a33fe875817dfb89b198942c01f9f26a0bab99d91c324ba9c3): complete subsection reference.

- [enable_introspection](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-e79ec8e5aab663141f0180c89d4e464a08425e1db1a850e16f47fc923b426f39): complete subsection reference.

<a id="canonical-290f1980133f247b630c0ff5e18496971903bd2debf9742bf3a6ecfa25faa6a9"></a>

<a id="canonical-e708f008b04ba8988285db95355fc691425f4ca67958c9fee2631166ac45909c"></a>

## max_batched_queries property — graphql_rules.graphql_settings / 23c0ee643232 / 4

Type: `"number"`. Computed.

Specify maximum number of queries in a single batched request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-815492afd9a0a66036ac1d4aab17900c02f9ec0e9fad727f821722f8430d74e7"></a>

<a id="canonical-33d4ffa90d3a1c08150473445a08a0fa7150e4476bc69eea5921a911c0154843"></a>

## max_depth property — graphql_rules.graphql_settings / 23c0ee643232 / 5

Type: `"number"`. Computed.

Specify maximum depth for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-f79f7b8fb22f9fe793c6fbc2a77aa2fef91d140cea30428de24ba62236aa080b"></a>

<a id="canonical-82c09ed8001f89333ef3cc09f5f0fd29fbcc9ba9bafcb0a95cb985a57958f11a"></a>

## max_total_length property — graphql_rules.graphql_settings / 23c0ee643232 / 6

Type: `"number"`. Computed.

Specify maximum length in bytes for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16386,
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
    "ves.io.schema.rules.uint32.lte": "16386"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  }
}
```

<a id="canonical-f22524ac87382e29384d26e64e49a2a723c87145d7db3a7ebecba058694ff366"></a>

## Next pages — graphql_rules.graphql_settings / 23c0ee643232 / 7

- [graphql_rules.graphql_settings.disable_introspection](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-524534144f5c72a33fe875817dfb89b198942c01f9f26a0bab99d91c324ba9c3)
- [graphql_rules.graphql_settings.enable_introspection](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-e79ec8e5aab663141f0180c89d4e464a08425e1db1a850e16f47fc923b426f39)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-524534144f5c72a33fe875817dfb89b198942c01f9f26a0bab99d91c324ba9c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53824a59ddde888dd675e512320f1ea5852f02948fd483621283d9dfb48f5e3c"></a>

## graphql_rules.graphql_settings.disable_introspection — graphql_rules.graphql_settings.disable_introspection / b97453a24f89 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- [graphql_rules.graphql_settings](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1f3509297bdca674cf4b43e723e9706eb4edd59c10358678e036485802218005)
- graphql_rules.graphql_settings.disable_introspection

<a id="canonical-f6f93b7372b171fed68f893f0b377027b67c38006e4d41696ebc2324903a9808"></a>

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

<a id="canonical-5e846d0672bb08dd112ba58c7afbd1f89d3d30a644f76f55f532d6178b0a99c2"></a>

## Direct properties — graphql_rules.graphql_settings.disable_introspection / b97453a24f89 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0363eb97ea3b7888e3da95f87dd5210059ffb4edd79dad55d523e8f34396ab95"></a>

## Next pages — graphql_rules.graphql_settings.disable_introspection / b97453a24f89 / 4

- [graphql_rules.graphql_settings](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1f3509297bdca674cf4b43e723e9706eb4edd59c10358678e036485802218005)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e79ec8e5aab663141f0180c89d4e464a08425e1db1a850e16f47fc923b426f39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa94a9b05f7115b3c8dfbd823bf5917110563e0f8d110d4a881789be783265d4"></a>

## graphql_rules.graphql_settings.enable_introspection — graphql_rules.graphql_settings.enable_introspection / 3a6f943ca0c8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- [graphql_rules.graphql_settings](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1f3509297bdca674cf4b43e723e9706eb4edd59c10358678e036485802218005)
- graphql_rules.graphql_settings.enable_introspection

<a id="canonical-70eb14dd0622f8488055a5d9fb4919e80b2842fb0373ffe274d0d3e88bfeb028"></a>

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

<a id="canonical-9dcc5f60421d246e3f838d4a7dcbc6fb64b5ada69ec9827e74b1858dc7af77a6"></a>

## Direct properties — graphql_rules.graphql_settings.enable_introspection / 3a6f943ca0c8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eaa21b50188e87f5782b98d5bf9c4ce3e9d5fff2be851298997ee854fba60af6"></a>

## Next pages — graphql_rules.graphql_settings.enable_introspection / 3a6f943ca0c8 / 4

- [graphql_rules.graphql_settings](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1f3509297bdca674cf4b43e723e9706eb4edd59c10358678e036485802218005)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-cdf9f20b46addf38e7e525b70b8d2028e41754120c00b8a4abcf78be684af34d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b285de0bdb5ad225767901149659030e47e78e90770daa5e467eaebe537aa5b"></a>

## graphql_rules.metadata — graphql_rules.metadata / 8daa1aaee75e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- graphql_rules.metadata

<a id="canonical-28b1d79c41b9656d88ba87dc232d32f584b43632c176b9d02f9f851986c6168e"></a>

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

<a id="canonical-b8fecba1fbf0c448d352e3cdccbe5ec60c3cd864f41a25cb35f15ed7050d034b"></a>

## Direct properties — graphql_rules.metadata / 8daa1aaee75e / 3

<a id="canonical-e9a6ccef25defe65f5f00ab054fa5f6979d105267053612ec67d47ebf66085a6"></a>

<a id="canonical-a08be3a451ea4902133749788efcc8c1eac2e39011c56108bcf83f104ac66d46"></a>

## description_spec property — graphql_rules.metadata / 8daa1aaee75e / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-44f62bda40732c315212ced14b16a3051e98d39bc2d43cac8b301c546e5a8d5f"></a>

<a id="canonical-1aa1e20caeef0f23fcbf3c114862c3535cf1537cbac690258c873dda448f0f62"></a>

## name property — graphql_rules.metadata / 8daa1aaee75e / 5

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

<a id="canonical-0fe7997e87e9205af7f9d04c34b70de602d06a8e967d326ce2aff831661d6717"></a>

## Next pages — graphql_rules.metadata / 8daa1aaee75e / 6

- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d9eb7c82005fa53260120d72e26ea410f19dea4038a2383d2a22ae4a69baecc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f48758522b6de45966c37035c9b6bd8171037c98c24f0fbc1ea75deb436dc49"></a>

## graphql_rules.method_get — graphql_rules.method_get / 4af035beab8f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- graphql_rules.method_get

<a id="canonical-dfa21fcf40c43288f28dc0ac8b24b06549876c8597da03d9ccfeeca1841269fe"></a>

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

<a id="canonical-b74d506fa39ba64bbb379dde64ab37ffedaacd79c7a8f743f630231a44cbb36c"></a>

## Direct properties — graphql_rules.method_get / 4af035beab8f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eb0e8ec45d55ad8624abb46893a52ae23ae82d75f2c0447bc8aae02c04e20947"></a>

## Next pages — graphql_rules.method_get / 4af035beab8f / 4

- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8f703a3f2906f2bccb4bd9d1fef06429016a5cb680ca451f9d0c2e6d9163cee7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a3e08a8c4c158beca82d6e69a007d3d746ff7d615d0c66a3e0d029a71f0f41a"></a>

## graphql_rules.method_post — graphql_rules.method_post / 3a724e833041 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- graphql_rules.method_post

<a id="canonical-07cbedb52149a53240a3f3ea075017a7ac2459e425e7e6ca13b12a8fe9aafb2c"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for method post.

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

<a id="canonical-3d7dc6c30263d9dc43c770b8db4c3d5284290b6b965258a9fb1297775283c463"></a>

## Direct properties — graphql_rules.method_post / 3a724e833041 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9ae2ab0202fea07c1f303c577c3fa8c8776503fa704e277be6d587eae8b43f9"></a>

## Next pages — graphql_rules.method_post / 3a724e833041 / 4

- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f1b85ea1fa6a6e2c26e23ee36f23c36b07e6cbe74339a722ebf7694908ccc9c5)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-008b300c77154c1484b58fff86572204ae1ec82ce5d78ccf747628a43b11a831"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d102ce787d6d2799c68cec6ea37f388992f5dfd8cbf5bfadfcd6d92d9158377c"></a>

## http — http / f57072a9b932 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- http

<a id="canonical-ed150a85893d413d1f14766720f08a3388c0652c2484b1c5faaae17b500c80f1"></a>

Type: `"single"`. Computed.

\[OneOf: http, https, https\_auto\_cert; Default: https\_auto\_cert\] HTTP Choice. Choice for
selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

OneOf alternatives in this subsection:

- [http](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ed150a85893d413d1f14766720f08a3388c0652c2484b1c5faaae17b500c80f1)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-beaafcd7cdb02ab04645248090c26ba38979b5ac1c7982efd2a525b06c0d39a3)
- [https_auto_cert](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-5f2ca249db3615bb879ecfc499fd83a7b2b924f8f0e748722509e5a423395e35)

Select alternatives according to the provider validators above.

<a id="canonical-7a86c8c39549ada9a92de9d21ac3d6fd2a3a92b5ef413d3ea6188f57f92b0e6f"></a>

## Direct properties — http / f57072a9b932 / 3

<a id="canonical-70324f689092e57d4c0c84ed24733a74b0257a74d8d6bc2b8d4abe158dc6b57f"></a>

<a id="canonical-17c129f4053071a7e7fe747d7bc6dfb43417df25adb61de273e58f19db085bd9"></a>

## dns_volterra_managed property — http / f57072a9b932 / 4

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

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

<a id="canonical-bc53e09e33235f754025b005b0635617fd9bc98d9265ced68266c934b3e3fcc9"></a>

<a id="canonical-5b23758fa54c86da9491fdee2fd8707010b8f18d1b4e9e9d90fe69992bfc94bb"></a>

## port property — http / f57072a9b932 / 5

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

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

<a id="canonical-6c1f5ad4c5c5f43a060866c7ecfa582cb53701e56644d7929a3b7559fd389ab2"></a>

<a id="canonical-4ffbb661117937575b63974653ceb750c7609f0b207e47a91077d1e3ab145449"></a>

## port_ranges property — http / f57072a9b932 / 6

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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

<a id="canonical-a230e555ead74ad6e9a2b4063b24feca4144a3f78189a13ce2fe8f919457e320"></a>

## Next pages — http / f57072a9b932 / 7

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79a47dfd85fbd1b48a8b996f71af53141d9f88e7a689665dbf25adaeefa9c2ee"></a>

## https — https / 6c135f7e5b01 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- https

<a id="canonical-beaafcd7cdb02ab04645248090c26ba38979b5ac1c7982efd2a525b06c0d39a3"></a>

Type: `"single"`. Computed.

Choice for selecting CDN Distribution with bring your own certificates.

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

<a id="canonical-c687a232a489aab984c913f28a50e79a2b4ad70a81b2a2720c00039ea04dea3b"></a>

## Direct properties — https / 6c135f7e5b01 / 3

<a id="canonical-b7bb99a340f1ab9524716af3410531ba27f3722a9aa401c6602b0a4df1add08c"></a>

<a id="canonical-dc310dd21bcccfdf2ec0388f819240017623db26b086a2224a32f0c56d4ad452"></a>

## add_hsts property — https / 6c135f7e5b01 / 4

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

<a id="canonical-5e78718fc8bf1518efc984980f3a092bc53a20287b47abd8b8d2d9aaceadfa68"></a>

<a id="canonical-eb1855498e344a57fb192210b029d294ec1d6bd9b4afab7c85bb61a4bd4fd6a4"></a>

## http_redirect property — https / 6c135f7e5b01 / 5

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

- [tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65): complete subsection reference.

<a id="canonical-d3d39c2fc2cc3c62ea605553d8129414089b304195dfdc83b9c909b541dca30d"></a>

## Next pages — https / 6c135f7e5b01 / 6

- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e02f021f150478a81e35d563fe10b83cc681f3b93dab31a7c07d02a23a2c59e"></a>

## https.tls_cert_options — https.tls_cert_options / 4318e0c9fa0d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- https.tls_cert_options

<a id="canonical-75751357b03c4e141c6b22e9a7e4c29618e0160131ed70dce517af9cb6aa962a"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert options.

Upstream description:

TLS Certificate OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_inline_params\"]"
}
```

<a id="canonical-452dfea0372aa25900fe8bc68365783ca2bc7f751b2fdeb46e930486cf8de587"></a>

## Direct properties — https.tls_cert_options / 4318e0c9fa0d / 3

- [tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e): complete subsection reference.

- [tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984): complete subsection reference.

<a id="canonical-5cafdcf6dcdfbfabc5a67dad313ba6d36db155f8ee9b683cbf74fc2e65e1c5fd"></a>

## Next pages — https.tls_cert_options / 4318e0c9fa0d / 4

- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c9e0c89b574b1236865587b2b69123f2474449ab777b77dc79b499694322126"></a>

## https.tls_cert_options.tls_cert_params — https.tls_cert_options.tls_cert_params / 939841bf1d3f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- https.tls_cert_options.tls_cert_params

<a id="canonical-e0a8fa82b525145098c6f7aa2e5b487db6df3c7e78540f7637a458f2bf48e309"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

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

<a id="canonical-7298818aa128e7d30077e70a28d868e2f55d0a0bdc921aae3f09f64c2afc7360"></a>

## Direct properties — https.tls_cert_options.tls_cert_params / 939841bf1d3f / 3

- [certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-996b02a920735b7c277c667b49f705f661fe7dd00b447f3944d05dc23fe5d2bc): complete subsection reference.

- [no_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-af9fa2f221b5654b8a49ce1df30b6f871cea0cb64cc912c9ba2a94fb401dc3ac): complete subsection reference.

- [tls_config](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b8f93ac4d57ce7e70e4b3abdff5097c700abc331074cb84664449091c9e169f0): complete subsection reference.

- [use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217): complete subsection reference.

<a id="canonical-3562142e2710b65999a6ae478cd2c19ed6a326ad5c4b7442d4be5a93f54ba2f3"></a>

## Next pages — https.tls_cert_options.tls_cert_params / 939841bf1d3f / 4

- [https.tls_cert_options.tls_cert_params.certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-996b02a920735b7c277c667b49f705f661fe7dd00b447f3944d05dc23fe5d2bc)
- [https.tls_cert_options.tls_cert_params.no_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-af9fa2f221b5654b8a49ce1df30b6f871cea0cb64cc912c9ba2a94fb401dc3ac)
- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b8f93ac4d57ce7e70e4b3abdff5097c700abc331074cb84664449091c9e169f0)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-996b02a920735b7c277c667b49f705f661fe7dd00b447f3944d05dc23fe5d2bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74655b192344d85861c166534973575d8f859c629f0baa079685da0ea5e22b38"></a>

## https.tls_cert_options.tls_cert_params.certificates — https.tls_cert_options.tls_cert_params.certificates / 8956cf85a293 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- https.tls_cert_options.tls_cert_params.certificates

<a id="canonical-0e4eb0fddfe6932c0676f8d690abccb61a27fcc328d6d202a6b8e8fbc46c1f60"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ca0e3b6443037bcad20f829270b882dcebe312770feb192b9c3beb58f8283701"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.certificates / 8956cf85a293 / 3

<a id="canonical-bd7c2c096745c8cf49285af1cf2190ad45e2fbab1426204fd563b41853f136ce"></a>

<a id="canonical-319cd9bba81a0a0bd2f736485c7332ea7560c81934dd9bd37f04b93c7756317c"></a>

## name property — https.tls_cert_options.tls_cert_params.certificates / 8956cf85a293 / 4

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

<a id="canonical-61bab438b54b16803d7476b7ef8b00711c6467869c252761945eff80613ea954"></a>

<a id="canonical-48693f9b3f2f3a0dba96e14dd5131a2bda423155204a3f9eeb8f513bdb328043"></a>

## namespace property — https.tls_cert_options.tls_cert_params.certificates / 8956cf85a293 / 5

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

<a id="canonical-d9dacabdab5f2da4fd8b2b24df54a057c2083d03387d6d0e9978aceb17a90172"></a>

<a id="canonical-90705e617364d08bacfba6ea198d715bc0383580c76daafe0b80c7ddaba30fce"></a>

## tenant property — https.tls_cert_options.tls_cert_params.certificates / 8956cf85a293 / 6

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

<a id="canonical-e4eb7a8366a7fff5b3660caafb1a6e81e03a355b54fb5a1bbfb0047089248f7d"></a>

## Next pages — https.tls_cert_options.tls_cert_params.certificates / 8956cf85a293 / 7

- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-af9fa2f221b5654b8a49ce1df30b6f871cea0cb64cc912c9ba2a94fb401dc3ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-785089a0df65570a9a4b26ac11176e83cb8952fe7b29adaa60d6d7260c65901a"></a>

## https.tls_cert_options.tls_cert_params.no_mtls — https.tls_cert_options.tls_cert_params.no_mtls / 3b1558e8c7f4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- https.tls_cert_options.tls_cert_params.no_mtls

<a id="canonical-3cdf6e6164c95118cfda3fe0da7c4cf52b17fdc098a2203feded0e29de0280d1"></a>

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

<a id="canonical-914512fbfb601de88e139416629987250c54506967e6d1833af37e5e5a5c819a"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.no_mtls / 3b1558e8c7f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-390d1bb9018efbdfc88381687782206464db1fb4dd47ada399e964fdafbe0dc1"></a>

## Next pages — https.tls_cert_options.tls_cert_params.no_mtls / 3b1558e8c7f4 / 4

- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-b8f93ac4d57ce7e70e4b3abdff5097c700abc331074cb84664449091c9e169f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36ac77e39e0755a9ec521744cc36498300ce520064f2c5761410f661074104a7"></a>

## https.tls_cert_options.tls_cert_params.tls_config — https.tls_cert_options.tls_cert_params.tls_config / fe1e1cd0f233 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- https.tls_cert_options.tls_cert_params.tls_config

<a id="canonical-6125fff04dcfefa16e425fbc7bc74c7e03531be5e8f02bd89d0ae8f15860851f"></a>

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

<a id="canonical-f3b1c0c3b40d7aefe1bbe213ea6e09dbd0278f24610bd89e02df794a512f295e"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.tls_config / fe1e1cd0f233 / 3

- [custom_security](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2b0b266b12ee0b89895ef7036c97dd86052314127a6690019bea255fbc9d6f94): complete subsection reference.

- [default_security](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b34e691af2a1647d7d58a1352e24c50325ecadb4c7e09d6d58b7a5ef2b694949): complete subsection reference.

- [low_security](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0f4d043f9c9d345cd3c5a07f86af1595469997086e9e67a9d40394d6924be462): complete subsection reference.

- [medium_security](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-6eeb532c0cd11923330988f4807c7d812d3e5122b59c12d0ee0d70e6803f69b1): complete subsection reference.

<a id="canonical-48a7905e7da29904d54bbdd97423bf51db9b4293b250640a93c22a1b44882ad6"></a>

## Next pages — https.tls_cert_options.tls_cert_params.tls_config / fe1e1cd0f233 / 4

- [https.tls_cert_options.tls_cert_params.tls_config.custom_security](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2b0b266b12ee0b89895ef7036c97dd86052314127a6690019bea255fbc9d6f94)
- [https.tls_cert_options.tls_cert_params.tls_config.default_security](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b34e691af2a1647d7d58a1352e24c50325ecadb4c7e09d6d58b7a5ef2b694949)
- [https.tls_cert_options.tls_cert_params.tls_config.low_security](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0f4d043f9c9d345cd3c5a07f86af1595469997086e9e67a9d40394d6924be462)
- [https.tls_cert_options.tls_cert_params.tls_config.medium_security](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-6eeb532c0cd11923330988f4807c7d812d3e5122b59c12d0ee0d70e6803f69b1)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2b0b266b12ee0b89895ef7036c97dd86052314127a6690019bea255fbc9d6f94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8bcc12c22f8efd4c1c6d7a8a190d45ed64707187b0586a4e4fba65959ae261b"></a>

## https.tls_cert_options.tls_cert_params.tls_config.custom_security — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 1226ec304f00 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b8f93ac4d57ce7e70e4b3abdff5097c700abc331074cb84664449091c9e169f0)
- https.tls_cert_options.tls_cert_params.tls_config.custom_security

<a id="canonical-b8ea2e78084546f4bc715d9c39b0f4fd0fa8c33be93bdb10d85b971d21606680"></a>

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

<a id="canonical-25037f15f5c96f5d185387a79e27be5c495d908a38c62579dbb7f7d068ced567"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 1226ec304f00 / 3

<a id="canonical-0a988642e68326e1c14bf5807c017feb3212ada180ce99da4cd7bc1a97ef3394"></a>

<a id="canonical-69d50f183205f348e684941cb07853b83679b7c72620e5a3969075c867bf1670"></a>

## cipher_suites property — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 1226ec304f00 / 4

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

<a id="canonical-31c490e39ef6dab0683c40d5cd17a8f38c6caa91ef208a5f3cf0e3e9572979d8"></a>

<a id="canonical-f6173d8e74f6b7c83ef1ffb9167bc3fd5ebd7553a401e5c84bde56d433b7c5fe"></a>

## max_version property — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 1226ec304f00 / 5

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

<a id="canonical-bf17ded4ce6b64448b160640d1d0d1a6b275faed6d9fa3147562a8a4c9689adc"></a>

<a id="canonical-68c383b875d4ede66b41f2fd32af09d574623a3ee30cba2f412afecdcc6db334"></a>

## min_version property — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 1226ec304f00 / 6

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

<a id="canonical-f1bf3be21d6b409c23a9c416bcf60b7e53c796eee370bf043d05c57d13320286"></a>

## Next pages — https.tls_cert_options.tls_cert_params.tls_config.custom_security / 1226ec304f00 / 7

- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b8f93ac4d57ce7e70e4b3abdff5097c700abc331074cb84664449091c9e169f0)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-b34e691af2a1647d7d58a1352e24c50325ecadb4c7e09d6d58b7a5ef2b694949"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b8444421a385ca25d329df923c3a95964c3698689adb0aa3099684b41c94726"></a>

## https.tls_cert_options.tls_cert_params.tls_config.default_security — https.tls_cert_options.tls_cert_params.tls_config.default_security / 702e52710d85 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b8f93ac4d57ce7e70e4b3abdff5097c700abc331074cb84664449091c9e169f0)
- https.tls_cert_options.tls_cert_params.tls_config.default_security

<a id="canonical-fcb8763e647051fb51c35436aa31b9ff3518aad09c51fa56402af6c6d0400a20"></a>

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

<a id="canonical-a0d9eb4ff5744de8a3ff0b0ba70739b4d3e15f68ee16372f4796c788f6e4dfb2"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.tls_config.default_security / 702e52710d85 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-920bd8903724cd901c431b40724a16d1d77a10f781ff54e5646a02a6bd4fc95b"></a>

## Next pages — https.tls_cert_options.tls_cert_params.tls_config.default_security / 702e52710d85 / 4

- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b8f93ac4d57ce7e70e4b3abdff5097c700abc331074cb84664449091c9e169f0)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-0f4d043f9c9d345cd3c5a07f86af1595469997086e9e67a9d40394d6924be462"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cccf4cf5268695b4decf0c52b06934954425ce4448b023a846d1a94e7f8caddd"></a>

## https.tls_cert_options.tls_cert_params.tls_config.low_security — https.tls_cert_options.tls_cert_params.tls_config.low_security / 9c3a216df49c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b8f93ac4d57ce7e70e4b3abdff5097c700abc331074cb84664449091c9e169f0)
- https.tls_cert_options.tls_cert_params.tls_config.low_security

<a id="canonical-9fd26b9b85aba4e5ac32cad54000d944bb45baa832e3638f2300e4ebdb60adb4"></a>

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

<a id="canonical-21f069efa01992d009ef32a0aca91583cb366cdef1cb8d68a9aee71689bb27c4"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.tls_config.low_security / 9c3a216df49c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d18b209d66ac9c446b98f7a7ce3ae9b97bdd4ff8b98bbab7db1e9ba6a6c3d27e"></a>

## Next pages — https.tls_cert_options.tls_cert_params.tls_config.low_security / 9c3a216df49c / 4

- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b8f93ac4d57ce7e70e4b3abdff5097c700abc331074cb84664449091c9e169f0)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-6eeb532c0cd11923330988f4807c7d812d3e5122b59c12d0ee0d70e6803f69b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41691634a949c32138dfb39618bd38e89817d817c033bd3fa871d45e96499c05"></a>

## https.tls_cert_options.tls_cert_params.tls_config.medium_security — https.tls_cert_options.tls_cert_params.tls_config.medium_security / c9282bfdcb7c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b8f93ac4d57ce7e70e4b3abdff5097c700abc331074cb84664449091c9e169f0)
- https.tls_cert_options.tls_cert_params.tls_config.medium_security

<a id="canonical-e01cf938e41e6640e6ba67b0e8aaedb6c52ba5f3494469da005f7e4f6a1878bd"></a>

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

<a id="canonical-5359c47408d8cb5c3a195f1ea5a84b611f8c47edb0deb49e29cc016ddea6dce0"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.tls_config.medium_security / c9282bfdcb7c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2fb351e24a34b746d85a0c95cbb401181633e84179c98e16f610c6f669c0d67"></a>

## Next pages — https.tls_cert_options.tls_cert_params.tls_config.medium_security / c9282bfdcb7c / 4

- [https.tls_cert_options.tls_cert_params.tls_config](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b8f93ac4d57ce7e70e4b3abdff5097c700abc331074cb84664449091c9e169f0)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c585e1a74d49d4cd3df762c29f46db281bfaddc2c9aad0b5bf9934bdc81d0831"></a>

## https.tls_cert_options.tls_cert_params.use_mtls — https.tls_cert_options.tls_cert_params.use_mtls / 3ffb442ebc9b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- https.tls_cert_options.tls_cert_params.use_mtls

<a id="canonical-b407ac2d6f0f4463d3a684ff8b745b47e727aa8193d3b21b3f09b6ff784bec64"></a>

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

<a id="canonical-f1dc20d0855e98b1c17dc7f342bbf08bca0b7d57e22a7da3bd5a1811d17193f7"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls / 3ffb442ebc9b / 3

<a id="canonical-a9f85a7289a66309ec65c66cbbbdaca98652357c5c4bddf595f0ae0e1b3a67f9"></a>

<a id="canonical-0177b8702bc33366cff83c1ee6948f4d4755f2c157e4dd09a3e9a52eff6394cb"></a>

## client_certificate_optional property — https.tls_cert_options.tls_cert_params.use_mtls / 3ffb442ebc9b / 4

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

- [crl](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ba90d0d1f1f3655eb1feb2951d046de6eac7be99992cbeb43e09460d37198a87): complete subsection reference.

- [no_crl](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-e2477f9350d10f64411a71c7cf08731bc33299c17fa29050a95c778ba5f8cd3b): complete subsection reference.

- [trusted_ca](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-37360ed01a55f065c6375de2d753fe47f093269e767af2509547668a8b2d8047): complete subsection reference.

<a id="canonical-8f22a7fd5d9e0ec7a74515b6290e95ebcacbf8181a6343a35fbceec328a2f6fc"></a>

<a id="canonical-40f6fb28c52c4975061bc571e8830aba5880e73e1d085e208a4b7468f1e0dc03"></a>

## trusted_ca_url property — https.tls_cert_options.tls_cert_params.use_mtls / 3ffb442ebc9b / 5

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

- [xfcc_disabled](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f12574c44f9eb7da3b778a5d06b84a34c498c3227060b6cee09529a5a45cdf2f): complete subsection reference.

- [xfcc_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b3acceeceff2ec7fca4ba3ee4c1f984f2ff9dafb7df0091e1a2d3a4c84c6de6b): complete subsection reference.

<a id="canonical-595f32fcbca14d299c8d653cd8da05dd7b2539a4fe89a537e99a78164c0c3b86"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls / 3ffb442ebc9b / 6

- [https.tls_cert_options.tls_cert_params.use_mtls.crl](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ba90d0d1f1f3655eb1feb2951d046de6eac7be99992cbeb43e09460d37198a87)
- [https.tls_cert_options.tls_cert_params.use_mtls.no_crl](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-e2477f9350d10f64411a71c7cf08731bc33299c17fa29050a95c778ba5f8cd3b)
- [https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-37360ed01a55f065c6375de2d753fe47f093269e767af2509547668a8b2d8047)
- [https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f12574c44f9eb7da3b778a5d06b84a34c498c3227060b6cee09529a5a45cdf2f)
- [https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-b3acceeceff2ec7fca4ba3ee4c1f984f2ff9dafb7df0091e1a2d3a4c84c6de6b)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ba90d0d1f1f3655eb1feb2951d046de6eac7be99992cbeb43e09460d37198a87"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-91f14e15514c5aa20525b5358e71e9d29ce2599e02cfc957cdeff8ea9eaec9dd"></a>

## https.tls_cert_options.tls_cert_params.use_mtls.crl — https.tls_cert_options.tls_cert_params.use_mtls.crl / 499754b990c2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217)
- https.tls_cert_options.tls_cert_params.use_mtls.crl

<a id="canonical-347b6be48c024b9e13a1a4f29f0fac0d475ce649ef9acb04cc240deed1267b21"></a>

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

<a id="canonical-e4ff2d7270bc800336999d861c8caf15fb8e6b3f622997b97e2f09fc92cf85eb"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls.crl / 499754b990c2 / 3

<a id="canonical-61ca95dd5cf00ded7624a4205f4ed573a35d724e36b0364a380eaac16c4a769c"></a>

<a id="canonical-7904ba21f2b10b3a31ddf2de4ca4b5663e98b2197cc536adf28f4e4fa2fbe322"></a>

## name property — https.tls_cert_options.tls_cert_params.use_mtls.crl / 499754b990c2 / 4

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

<a id="canonical-f9a121e20a728c4666eb531d0c367ae04ad72a94c38914931f77a4f1de99b722"></a>

<a id="canonical-54d2c543e5b735c4db8f3ecf6fb64d9b527b4a468b50f09002c8ccddfa739b59"></a>

## namespace property — https.tls_cert_options.tls_cert_params.use_mtls.crl / 499754b990c2 / 5

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

<a id="canonical-43b626a6c4e563e3d7ba7f25e933989b69b9f966ed90c9b987e2ec60d3e91937"></a>

<a id="canonical-d2dd30ce1ce7e2f8ca7c5bcfbd4b9763730d5be92f6e9247c76acaededcf27cb"></a>

## tenant property — https.tls_cert_options.tls_cert_params.use_mtls.crl / 499754b990c2 / 6

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

<a id="canonical-2565fdb11c42586f6b068b3750a643a3e9d90bca8dcac16dca8ede5bd3bad828"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls.crl / 499754b990c2 / 7

- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e2477f9350d10f64411a71c7cf08731bc33299c17fa29050a95c778ba5f8cd3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf3557ad224f07c5edc3b0dedc1cdfcde1275e162c55686fbcb334f1a8a5c83c"></a>

## https.tls_cert_options.tls_cert_params.use_mtls.no_crl — https.tls_cert_options.tls_cert_params.use_mtls.no_crl / ba19a279ab0d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217)
- https.tls_cert_options.tls_cert_params.use_mtls.no_crl

<a id="canonical-9923cb0f545a02f80c3a89bec7946fa7554a6f6bbc2dc97bf5daecf09369fc1a"></a>

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

<a id="canonical-e8e4ae8b5d03a6ae591a77d17a81b61bc2f2281a6709f351eacfd2e92e83bb72"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls.no_crl / ba19a279ab0d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-376a557d2bfbfbdb1ae4db34bc7f7b3ccd1220b5721047f0a188846192695288"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls.no_crl / ba19a279ab0d / 4

- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-37360ed01a55f065c6375de2d753fe47f093269e767af2509547668a8b2d8047"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28e42604160b38c3a14c5b0b63eac645048a6c58fa24085b29de89e0efe7f3bf"></a>

## https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / d803772985b1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217)
- https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-d4d511f951d4c2fabf97bd9e0e9bc23c4efeaf7625de58323e3ecd01b0e5a5c2"></a>

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

<a id="canonical-4aded74a44863e3fcc47f14a56673b41b9f3069262889031c1793a0ca3854a77"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / d803772985b1 / 3

<a id="canonical-57c7674507f6b54891354ef29f11d2bc0d646718a90201162a21e2d523dd1b5d"></a>

<a id="canonical-da16cc674f9ee26d0e1e4ae3688a855e1b6e2efb0faff229d1c89750492b4a88"></a>

## name property — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / d803772985b1 / 4

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

<a id="canonical-c25957b6cf2ba7e7a6d097d8b8368cb5712f222d7c9b6a8d987526540196518c"></a>

<a id="canonical-5fed1a5ee57e20797cdbd679612f10bfc537c87af37fb020d723788f60315660"></a>

## namespace property — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / d803772985b1 / 5

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

<a id="canonical-39ab5243870bd8c9b9ed19aa4a80e87e670019ed45aca857618daa5f60c28601"></a>

<a id="canonical-9933ecaf0eaddada4305bece00c84fb290a5275db14e76ac4b422c846c11da15"></a>

## tenant property — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / d803772985b1 / 6

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

<a id="canonical-c4f715dd60e732e49b00e71b2deffe44583029048978fd7f1efc89582199d488"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls.trusted_ca / d803772985b1 / 7

- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f12574c44f9eb7da3b778a5d06b84a34c498c3227060b6cee09529a5a45cdf2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-235225a92cf2f3bada6d8dc396a907a4a086c69a88f1b86bc8a33444a9f46345"></a>

## https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled / 7d2505c3aa81 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217)
- https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-bc82a58ff096365311de909d7001aa9a59bd71a3466ec80a712255e52b040aba"></a>

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

<a id="canonical-614b5f1d1ff85d655f94ab76aeeaf2db0b4bdf5770e1acf0e3e6a0a7ca104c67"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled / 7d2505c3aa81 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e784b65117bee8c35d8d53f76da6e63c7ccdfed602d1fb22cb089a9a8c372af2"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_disabled / 7d2505c3aa81 / 4

- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-b3acceeceff2ec7fca4ba3ee4c1f984f2ff9dafb7df0091e1a2d3a4c84c6de6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fda5064816599527a3ce951f4fa54b9ad0ffdaff11af709590f8e81fe37faae0"></a>

## https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options / c56b5f02a767 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_cert_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-929ccee005b1482d26ddc1c02b9a84f09f402d579bd9b6b09c5a007b4ad0578e)
- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217)
- https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-7c9e3c80e4215c40f29bfa0f2fdb5652a8669999b26319027b6594eb9a582340"></a>

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

<a id="canonical-1765b514a4585a0ca973909286d4b5558e5a8cef3bb930b0085f36ae5bfdd653"></a>

## Direct properties — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options / c56b5f02a767 / 3

<a id="canonical-764095c50bf3b90e53a2a0217548bf77f6a970474c5980b42e21bb7f4ee034d8"></a>

<a id="canonical-fae069008f781fa723975b87846281040ccbda6577e223bd7890d886a9c43d83"></a>

## xfcc_header_elements property — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options / c56b5f02a767 / 4

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

<a id="canonical-5d74dec6fce2f5d8777f99ad36408f2f298387377eb96435106b5caf011c9a56"></a>

## Next pages — https.tls_cert_options.tls_cert_params.use_mtls.xfcc_options / c56b5f02a767 / 5

- [https.tls_cert_options.tls_cert_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2c2af69488ff9ab601e5e15fc3ac9ef4d8b6a978ee6e0162260d986145068217)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0fe046c8d74a3962339650c6d140beca4dfc577f48a10a8f90319625b415ffc"></a>

## https.tls_cert_options.tls_inline_params — https.tls_cert_options.tls_inline_params / fbfbfedde766 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- https.tls_cert_options.tls_inline_params

<a id="canonical-4aca1cea1cec1d55f61c00ee888450983429116cc204f434e62efd50177e1bb1"></a>

Type: `"single"`. Computed.

Configuration parameter for tls inline params.

Upstream description:

Inline TLS parameters.

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

<a id="canonical-88c107e45929459126d6da61b95ab2215771e8ffeccd02219b0000560c306f83"></a>

## Direct properties — https.tls_cert_options.tls_inline_params / fbfbfedde766 / 3

- [no_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-fa82af9501f6aeee0ecfed5fe25b5daab9cc1f85cad53c355e09587d4b8a6ea5): complete subsection reference.

- [tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1): complete subsection reference.

- [tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-574182211b3e722bb2d0f6d4a1dedc2e117fcca6866a437156769d4326f782d9): complete subsection reference.

- [use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097): complete subsection reference.

<a id="canonical-46af3a37db0f888e4120efbfb183c005bfebec6f9ed44441eaa8bc6c7fc1acec"></a>

## Next pages — https.tls_cert_options.tls_inline_params / fbfbfedde766 / 4

- [https.tls_cert_options.tls_inline_params.no_mtls](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-fa82af9501f6aeee0ecfed5fe25b5daab9cc1f85cad53c355e09587d4b8a6ea5)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1)
- [https.tls_cert_options.tls_inline_params.tls_config](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-574182211b3e722bb2d0f6d4a1dedc2e117fcca6866a437156769d4326f782d9)
- [https.tls_cert_options.tls_inline_params.use_mtls](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-48661bc3667c9841bd24b9b415a174ccea952e93bd5c464a121755cfd5b92097)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-fa82af9501f6aeee0ecfed5fe25b5daab9cc1f85cad53c355e09587d4b8a6ea5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1737a6edda27bdb0dd1af0ed8ce7c3a5e23e6a554df7f9735d2467159420bcc"></a>

## https.tls_cert_options.tls_inline_params.no_mtls — https.tls_cert_options.tls_inline_params.no_mtls / c568ab48a9cf / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- https.tls_cert_options.tls_inline_params.no_mtls

<a id="canonical-d26bd19a3f0b13dd10ba3482ea491f40223331d5f76161316b6e5c05247ecb11"></a>

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

<a id="canonical-3d5ab6087d7cbfe551852d009c136aabeaf8bcba7fce4302d813a250d474c6c9"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.no_mtls / c568ab48a9cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-788b507c39aa6e5d3d1a47e548193d24ea38fb36a8f9b89f0dbd3878f5edbf6f"></a>

## Next pages — https.tls_cert_options.tls_inline_params.no_mtls / c568ab48a9cf / 4

- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a09ef70512d017b4e1781b88ff45b97bc9cab9e120168339965979826980e604"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates — https.tls_cert_options.tls_inline_params.tls_certificates / 36fef117dc70 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- https.tls_cert_options.tls_inline_params.tls_certificates

<a id="canonical-bb8abedfec246116c580481b66eac13802daa1c6fc2033404e9b9648029d53ba"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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

<a id="canonical-bed4f2b1a741e7ad3448395240c2013764ee6902d51aabff2de3cbbb41de032e"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates / 36fef117dc70 / 3

<a id="canonical-dd1a20ebea0177e3737a20856d7cbfbe8552da0aef0e08167eeebfae4f8b3b7d"></a>

<a id="canonical-00b92da8aeb907103965191bf25b00d2dc7a21b3d6065ac69983a9693a059f02"></a>

## certificate_url property — https.tls_cert_options.tls_inline_params.tls_certificates / 36fef117dc70 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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

- [custom_hash_algorithms](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-dc9e00d6f556f421e02ae8efc944bdb35db5452fc127faaf6fa9b0b8966609af): complete subsection reference.

<a id="canonical-f0df85eb516bf67ac9353de9da7c9a089ce3ae485c54dabd2dc24890b960ca3a"></a>

<a id="canonical-06d4195b1e0c884d2d5f4b15ac024f9dcce80772377dbfd4179f5d7f35d44180"></a>

## description_spec property — https.tls_cert_options.tls_inline_params.tls_certificates / 36fef117dc70 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-67575e7a21691aaf0836a98b0eaec12758e148de648d9fe89c7518fa87f40822): complete subsection reference.

- [private_key](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3d56980d1415ea8c839ec57ed0d9ad43d643f19a2b2af2b7b85280ae717d758b): complete subsection reference.

- [use_system_defaults](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0a0ff49597b60275fbdfc2986691bcf1f247e11a50e307597042ffd797282635): complete subsection reference.

<a id="canonical-43b85dd1bb3484ec4171c7feac4f5b23e813bface7df9359b39d934a9fc9561e"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates / 36fef117dc70 / 6

- [https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-dc9e00d6f556f421e02ae8efc944bdb35db5452fc127faaf6fa9b0b8966609af)
- [https.tls_cert_options.tls_inline_params.tls_certificates.disable_ocsp_stapling](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-67575e7a21691aaf0836a98b0eaec12758e148de648d9fe89c7518fa87f40822)
- [https.tls_cert_options.tls_inline_params.tls_certificates.private_key](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3d56980d1415ea8c839ec57ed0d9ad43d643f19a2b2af2b7b85280ae717d758b)
- [https.tls_cert_options.tls_inline_params.tls_certificates.use_system_defaults](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0a0ff49597b60275fbdfc2986691bcf1f247e11a50e307597042ffd797282635)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-dc9e00d6f556f421e02ae8efc944bdb35db5452fc127faaf6fa9b0b8966609af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a72d0d24ff9b4b40d646403af01e6c777d507c42726525876d56d3afe1a42960"></a>

## https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms — https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms / f95f94e13dd2 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [https](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-ecdbae149b8ba9ab361be31238c648e67d23a59ca92b2d8f18e6ba12d97a21af)
- [https.tls_cert_options](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f8355af7ebbf3d07f5de71e93d8112caf76f4216ee7f12a5bc0c671546f0da65)
- [https.tls_cert_options.tls_inline_params](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-f73899440519b3f0f15e8d6cfb63f16ab46dd733d72edd3ca0e555a7650b9984)
- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1)
- https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms

<a id="canonical-cc52b1acb0bc5383729ef14ce467af29793755336d0cb2ed0dc899f1c54865fe"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-4c283d0d8813ec916c348edbeabd475ee4644141e31874c0acf47cece56bd48b"></a>

## Direct properties — https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms / f95f94e13dd2 / 3

<a id="canonical-f942153aa1e126e8b4d6915a1e2b84c257627fb7852de3aab95acd5f078a1957"></a>

<a id="canonical-7b2c52bf391414cf9fe570b9d589adac03d4641a26b4c01397e20a57e1f835d7"></a>

## hash_algorithms property — https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms / f95f94e13dd2 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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

<a id="canonical-98fcd823ab55cdb343ae6c43f2593a72bda766211941c354be2bfe7545a701ce"></a>

## Next pages — https.tls_cert_options.tls_inline_params.tls_certificates.custom_hash_algorithms / f95f94e13dd2 / 5

- [https.tls_cert_options.tls_inline_params.tls_certificates](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0d63fafab76e3e11cb7283dcf7d98f28495489a8cf658156f8c32daa65a466f1)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-67575e7a21691aaf0836a98b0eaec12758e148de648d9fe89c7518fa87f40822"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
