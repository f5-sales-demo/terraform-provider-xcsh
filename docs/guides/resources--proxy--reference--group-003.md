---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-609e053e3ba34f92a3ef08680f95bb251adb4d64cb1a53dafd17d829d62a6f3d"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.blindf / 43692e9f241c / 7

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-6142d88b105859ca184c72d96cabf0a447fba55c5a20892c0e5d543a7c98e3c0)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a6d44936c1eeaa514e00676cb4cfba3c369dadfcf076895c35f49e7713164600"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9290d563f79597ae9f31fb74c0491005be43ba2675d0cc5ebdb56d9cd665f1e"></a>

## dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_ / eb936ae603a9 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add](resources--proxy--reference--group-002.md#canonical-fae9c0d07abff912657654993d22c3da6352fdf4cf317a6e8ce52d5d69a406e3)
- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-6142d88b105859ca184c72d96cabf0a447fba55c5a20892c0e5d543a7c98e3c0)
- dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-e8731b958be537a080842874982cc214d1cbbd3431a5d80ae6d2f910974a3dfd"></a>

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

<a id="canonical-439f77547756ea3049fcc923ccc85f762dd391ab97b1ce10d97d75f46a8a6223"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_ / eb936ae603a9 / 3

<a id="canonical-44a87d54c7da78a9044f8a908905ff696656930cf454f4bbdda5ae22189d3345"></a>

<a id="canonical-6884d66a9bacb8a9148a8b4a74c1674f6e8b92f725a94717f77a54c0c51f79a4"></a>

## provider_ref property — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_ / eb936ae603a9 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-adda5788761c2314be2e9676df92eead10fb051d0fb5fe8923eec45fc001296d"></a>

<a id="canonical-d40bf0f51a6ac131e817f011ddefa3826628d0e7213fa75e2a3985aaaa478ba1"></a>

## url property — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_ / eb936ae603a9 / 5

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

<a id="canonical-2021b293e659b45846ef7dbba3c0695e2c8117699702c7e779e1ce7260a788f3"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value.clear_ / eb936ae603a9 / 6

- [dynamic_proxy.https_proxy.more_option.request_cookies_to_add.secret_value](resources--proxy--reference--group-002.md#canonical-6142d88b105859ca184c72d96cabf0a447fba55c5a20892c0e5d543a7c98e3c0)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-be754a741daec268e457f29a08324bc0d402e334ca27b29e3d1d5fb90f4ef0d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-738d44fa6735e71847e417fca42b5fce529c4833e4d66d5e938196c4d5218525"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add — dynamic_proxy.https_proxy.more_option.request_headers_to_add / 57b9ba0758f7 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add

<a id="canonical-ce072dfd8f22b07a7a0584e657cdafd9d28cdf2b78f8adaf9e342ff8577dc012"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-2949ffc265dbddf224327c6d9916a274cbd5b9e08446b32d09ed68bc0d14343d"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.request_headers_to_add / 57b9ba0758f7 / 3

<a id="canonical-efc7994d878284b47658b493c51441d94b1dcfb5abd77f335affcee835c0d6d2"></a>

<a id="canonical-6b5e5435a4ff0e362a9408b62303d80a1b3de0b550b2bdb2962844c663c4459e"></a>

## append property — dynamic_proxy.https_proxy.more_option.request_headers_to_add / 57b9ba0758f7 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d8e65962a54579ec7a124c9aff5c62508eab50bf2f428e04b662121947bb3e49"></a>

<a id="canonical-71d75a8b596df834b9bc4327238f522e1fe0546bf18a171f2da4409048aea78b"></a>

## name property — dynamic_proxy.https_proxy.more_option.request_headers_to_add / 57b9ba0758f7 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--proxy--reference--group-003.md#canonical-572cb769cf0f3be78ef41568e6679e584b4291fba87ea031e56078652809d5df): complete subsection reference.

<a id="canonical-0e268ab06b6c284734e468477205ff53d3851ec725f68001a9f87b4f63aebcc9"></a>

<a id="canonical-9606b9671d02b8d8270910114628e90e5ef5d3d87dde15e90006317380319606"></a>

## value property — dynamic_proxy.https_proxy.more_option.request_headers_to_add / 57b9ba0758f7 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

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

<a id="canonical-7809356e20e1023afb4a42fa430127190e7cb56a6b795074aa145063ed3b12fe"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.request_headers_to_add / 57b9ba0758f7 / 7

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-572cb769cf0f3be78ef41568e6679e584b4291fba87ea031e56078652809d5df)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-572cb769cf0f3be78ef41568e6679e584b4291fba87ea031e56078652809d5df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffbdbc32f64b7f7ba277447c0eba13dd17252b3f2f32e079f8fefd1ea62d5ced"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value / 493e02fcec05 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-003.md#canonical-be754a741daec268e457f29a08324bc0d402e334ca27b29e3d1d5fb90f4ef0d9)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value

<a id="canonical-73126fe1365aa11cd858b53715c3f3b37353bf5f968daebbba22a5c1b42d88bb"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-924c1321b87ceb0882c6868e86809ede0e53d53548cc9118b8d2c3a5eda5831e"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value / 493e02fcec05 / 3

- [blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-c0f87f9ce3188650c774bbc3c0e462a94a9466529077128234e914bc59e987c5): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-003.md#canonical-c7969ff1126d1dfaa60439703fcc936898b7aad20ad257ce6267ce10d8bba01d): complete subsection reference.

<a id="canonical-bf0a0d322bb2e05a2b3c63812f5b4a4407df5c32dc7e103c9395edc087ef06bf"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value / 493e02fcec05 / 4

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-c0f87f9ce3188650c774bbc3c0e462a94a9466529077128234e914bc59e987c5)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-c7969ff1126d1dfaa60439703fcc936898b7aad20ad257ce6267ce10d8bba01d)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-003.md#canonical-be754a741daec268e457f29a08324bc0d402e334ca27b29e3d1d5fb90f4ef0d9)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c0f87f9ce3188650c774bbc3c0e462a94a9466529077128234e914bc59e987c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-720c974994d8b76768e68aa60c9af21deef420c0c5a3dbfc4d9002178839cb8f"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindf / 3ef4ba26847d / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-003.md#canonical-be754a741daec268e457f29a08324bc0d402e334ca27b29e3d1d5fb90f4ef0d9)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-572cb769cf0f3be78ef41568e6679e584b4291fba87ea031e56078652809d5df)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-def5db9c9167aec09fd16d0d41983afcf91e6dd8f1f6ce63eb33b87bec0d9079"></a>

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

<a id="canonical-b5e5a42983bfe0df5195cd67a7bce952ad189c699392e05010b5c1b5c08fc98a"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindf / 3ef4ba26847d / 3

<a id="canonical-dde2e159539d61784b189d2652da5bf878df7f26b09e8c47796c63d0d7f2fcfe"></a>

<a id="canonical-a8e76d6ded88e13423ffdc1ce8858030e9321cb5aee7548aeb1393505af67a4f"></a>

## decryption_provider property — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindf / 3ef4ba26847d / 4

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

<a id="canonical-fc5157defd103dde13ce1960fb9e5925d4d9a23e249ea385f5a9775920a326f6"></a>

<a id="canonical-08f46db747d33f8edaf1f48d24f3986c080683c7dc008f40b3b86818672ca596"></a>

## location property — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindf / 3ef4ba26847d / 5

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

<a id="canonical-b83ce02da65e07f432eaeb8236c58d6b88d1231b72acf58357f7526f70e5c206"></a>

<a id="canonical-663177d53d09dc0105113a0e010aa8daf8974d9227d088cfe7234e7a8db910fe"></a>

## store_provider property — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindf / 3ef4ba26847d / 6

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

<a id="canonical-308307ce2976b6b5c12e1b5491107406c5e912bb79ada23b6c1a939e9a46ea65"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.blindf / 3ef4ba26847d / 7

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-572cb769cf0f3be78ef41568e6679e584b4291fba87ea031e56078652809d5df)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c7969ff1126d1dfaa60439703fcc936898b7aad20ad257ce6267ce10d8bba01d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb714a236df1bab789d7131ce969fcc569ff45544bb93682121c779f2f16ba68"></a>

## dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_ / 0cf21bdc9d19 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add](resources--proxy--reference--group-003.md#canonical-be754a741daec268e457f29a08324bc0d402e334ca27b29e3d1d5fb90f4ef0d9)
- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-572cb769cf0f3be78ef41568e6679e584b4291fba87ea031e56078652809d5df)
- dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-345ad5df88709767ba92e1c629db27f3591c45ea8c2bd33b5df15be8832f72ca"></a>

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

<a id="canonical-30780ff6c20e02e1d22b10b60a2db6fdda3d5c76e33b6fc588e1237c12dcdf90"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_ / 0cf21bdc9d19 / 3

<a id="canonical-b3236a0e8ce5390de9ed502a55c72adad8f914e67dd633f0af33bd567ac32907"></a>

<a id="canonical-ac1718e945d7875c0d214af8e933165e32c9560371a16f29bb1bf9e3e6153af8"></a>

## provider_ref property — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_ / 0cf21bdc9d19 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-9307256dd812fbe704bb57f17b49adad4e7bd86e7eb78776d17f1ed67c1d5c69"></a>

<a id="canonical-4664de95243451c2c9a69c58397acb5093374e435812280017107291dd393aa8"></a>

## url property — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_ / 0cf21bdc9d19 / 5

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

<a id="canonical-79a9cc79dc647e49b0bc58b11924f5602d63b928751a61a6860f92616d55f71b"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value.clear_ / 0cf21bdc9d19 / 6

- [dynamic_proxy.https_proxy.more_option.request_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-572cb769cf0f3be78ef41568e6679e584b4291fba87ea031e56078652809d5df)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93ef38b08bb5b82b07f0e712d6f220f6d0523542838b2fcf1204abb736a9f52e"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add — dynamic_proxy.https_proxy.more_option.response_cookies_to_add / 155ee173440f / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add

<a id="canonical-1d84e689c48818ddff97d653b62c8d868120ffc466292bf07ba246fb0a1f68b5"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_domain",
    "ignore_domain"),
  validators.ConflictingListObjectAttributes("add_expiry",
    "ignore_expiry"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_partitioned",
    "ignore_partitioned"),
  validators.ConflictingListObjectAttributes("add_path",
    "ignore_path"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "secret_value"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "value"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-78079bae50e7dd1dfcce201dd0260aba0b272212ece6b2cda5db094786eaa7bc"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add / 155ee173440f / 3

<a id="canonical-f6789ce8c613669bd9b2829337acd0ee081f63d6da84ddbde1e63c3ff227bdd3"></a>

<a id="canonical-df317880595c3586daf5012da8ca07f79368477f781820d1756a5386e993d387"></a>

## add_domain property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add / 155ee173440f / 4

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-7ea6ab77a932b9e7c75088ee104f1a817160ba9ca3cd6516ab2e689160f5dcb2"></a>

<a id="canonical-aa038e77427d04d19e392faa5f1e25e9958bc9f610403a69adc4bec8d4a2ac10"></a>

## add_expiry property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add / 155ee173440f / 5

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [add_httponly](resources--proxy--reference--group-003.md#canonical-3a12e45a28ec3467fb740e4c3012c9cb7d2807e1e1832014c9eb92a1626cad69): complete subsection reference.

- [add_partitioned](resources--proxy--reference--group-003.md#canonical-a6834f134ebbd949976a2c95c5fb42323a02d67ca0c6b63a32f4c2b1db596212): complete subsection reference.

<a id="canonical-3964b73c11fd6b8bb2c62edcbf9086b344577d305fb15f170f8e3c812097f48d"></a>

<a id="canonical-32e733c82af7b188e464847392b149a5dee53f443777f4483487d396700d36a7"></a>

## add_path property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add / 155ee173440f / 6

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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

- [add_secure](resources--proxy--reference--group-003.md#canonical-01107677f3e3820491503fdbae79744fb28cfd7a54a9a127ae6558a07a214671): complete subsection reference.

- [ignore_domain](resources--proxy--reference--group-003.md#canonical-9c982705f56e25117f973fb09d43e5084ae26d48fdd3c3ded526243cceeeedb5): complete subsection reference.

- [ignore_expiry](resources--proxy--reference--group-003.md#canonical-524c3fd30d7f05668c2e13bf1ba341ba70763bece6546fecc784f2358e7ae34a): complete subsection reference.

- [ignore_httponly](resources--proxy--reference--group-003.md#canonical-f9d86a52859ceca815ad5125ceb05e17f48ab62e1e7dc7561c2570f073a83d7b): complete subsection reference.

- [ignore_max_age](resources--proxy--reference--group-003.md#canonical-a08909d51e648a8163feba9a2883e2a78390a268f50f8bffaf751506131e37c7): complete subsection reference.

- [ignore_partitioned](resources--proxy--reference--group-003.md#canonical-c44c60ad6cbca6561e4d84e99175cb613e1b6d4d4a289dddb74514d527efc4fd): complete subsection reference.

- [ignore_path](resources--proxy--reference--group-003.md#canonical-8f14d59b4183111a9209da1cb008f5b31ab5fef6270746455aeb68243020cd15): complete subsection reference.

- [ignore_samesite](resources--proxy--reference--group-003.md#canonical-659c1931ba3c6ebd537ab10b39f88f92fcebc47c9f2f7c3f67455a2374bc3819): complete subsection reference.

- [ignore_secure](resources--proxy--reference--group-003.md#canonical-a7e75c90e86fe4f8258cf911bb0dacbef8ec9d00fde604a4b381e7b430382087): complete subsection reference.

- [ignore_value](resources--proxy--reference--group-003.md#canonical-7719c84eaa64186ab06f78f529048a8e0f20f7492d5476db906d26912f0d60a0): complete subsection reference.

<a id="canonical-3c748cbd0e66691206e35dcf6438cea3c91b03e2b5276c850451f24428707a77"></a>

<a id="canonical-6e2318e99efa6d8c791d86c9f7b10fe1bb4deb14e570fe291f5b0db0f5889b65"></a>

## max_age_value property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add / 155ee173440f / 7

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-5fc386f84a7d58de9d577685472fa5d2da541182f6e4f1383b16175745d9fec8"></a>

<a id="canonical-97e736fac4f35924f480d6aeb2ebebc98ddfdcb1df92ad73f85fab1beea88a7d"></a>

## name property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add / 155ee173440f / 8

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-b4e7c27c08e39711eb0b98108dc38bcda293b93ad7ada7990eed58bb5b848c36"></a>

<a id="canonical-690b968c5180812c4273f48ca7f1d239c66a090454a2550446b8eb61e3175321"></a>

## overwrite property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add / 155ee173440f / 9

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](resources--proxy--reference--group-003.md#canonical-1cf64b407579607c907c229c03f89aacce91d8a163f079c487032abb5dc1b787): complete subsection reference.

- [samesite_none](resources--proxy--reference--group-003.md#canonical-25937071e708a4294d08158a2eb50e7f8c9a8d33b88f0d9229c4d9cb12afd3f8): complete subsection reference.

- [samesite_strict](resources--proxy--reference--group-003.md#canonical-e5dfd87293400e16dbac119a439230ce800bb6f56210a17e1140694a021416d4): complete subsection reference.

- [secret_value](resources--proxy--reference--group-003.md#canonical-0d93353ca93264df4d71af2016f54f3c660b2763e3869964a84e0ad00ecdc1c5): complete subsection reference.

<a id="canonical-7b5390589429f465d26dbc3fd56dcdfd8b23d91e2c1248fba4e9daa74a5bc784"></a>

<a id="canonical-69bf6c732e0306791ff32bfac7dcab0f574825a766a52ad3ec124f9cc140e2d6"></a>

## value property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add / 155ee173440f / 10

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

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

<a id="canonical-e4c39db320d86a65c544fdc102877a0ecf8bc11b1332f290607090cf930c1f42"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add / 155ee173440f / 11

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly](resources--proxy--reference--group-003.md#canonical-3a12e45a28ec3467fb740e4c3012c9cb7d2807e1e1832014c9eb92a1626cad69)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned](resources--proxy--reference--group-003.md#canonical-a6834f134ebbd949976a2c95c5fb42323a02d67ca0c6b63a32f4c2b1db596212)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure](resources--proxy--reference--group-003.md#canonical-01107677f3e3820491503fdbae79744fb28cfd7a54a9a127ae6558a07a214671)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain](resources--proxy--reference--group-003.md#canonical-9c982705f56e25117f973fb09d43e5084ae26d48fdd3c3ded526243cceeeedb5)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry](resources--proxy--reference--group-003.md#canonical-524c3fd30d7f05668c2e13bf1ba341ba70763bece6546fecc784f2358e7ae34a)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly](resources--proxy--reference--group-003.md#canonical-f9d86a52859ceca815ad5125ceb05e17f48ab62e1e7dc7561c2570f073a83d7b)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age](resources--proxy--reference--group-003.md#canonical-a08909d51e648a8163feba9a2883e2a78390a268f50f8bffaf751506131e37c7)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned](resources--proxy--reference--group-003.md#canonical-c44c60ad6cbca6561e4d84e99175cb613e1b6d4d4a289dddb74514d527efc4fd)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path](resources--proxy--reference--group-003.md#canonical-8f14d59b4183111a9209da1cb008f5b31ab5fef6270746455aeb68243020cd15)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite](resources--proxy--reference--group-003.md#canonical-659c1931ba3c6ebd537ab10b39f88f92fcebc47c9f2f7c3f67455a2374bc3819)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure](resources--proxy--reference--group-003.md#canonical-a7e75c90e86fe4f8258cf911bb0dacbef8ec9d00fde604a4b381e7b430382087)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value](resources--proxy--reference--group-003.md#canonical-7719c84eaa64186ab06f78f529048a8e0f20f7492d5476db906d26912f0d60a0)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax](resources--proxy--reference--group-003.md#canonical-1cf64b407579607c907c229c03f89aacce91d8a163f079c487032abb5dc1b787)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none](resources--proxy--reference--group-003.md#canonical-25937071e708a4294d08158a2eb50e7f8c9a8d33b88f0d9229c4d9cb12afd3f8)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict](resources--proxy--reference--group-003.md#canonical-e5dfd87293400e16dbac119a439230ce800bb6f56210a17e1140694a021416d4)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-0d93353ca93264df4d71af2016f54f3c660b2763e3869964a84e0ad00ecdc1c5)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-3a12e45a28ec3467fb740e4c3012c9cb7d2807e1e1832014c9eb92a1626cad69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d77648ea23ef18a8b806a751f5d7e26d47f755c8ec82ed468de12de32a907d70"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly / 95cf79de3a43 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly

<a id="canonical-fde5e5547386cc321f7b371871c740feccda33bd6cee8ba5e1857da2f20c5033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

<a id="canonical-9c2e0ec29b62be37d760aff337c0024e0cbe8f9b58ce16248d678ef6e8f229dc"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly / 95cf79de3a43 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6bf4ac6c8d8b3e209113673c41e03f25df3df066a67b1aea180cf5bc1e20e122"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_httponly / 95cf79de3a43 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a6834f134ebbd949976a2c95c5fb42323a02d67ca0c6b63a32f4c2b1db596212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23a9006a751f27d6efb760c8a2b58597ce5d6aae75bca50b1d8fad59a7459cef"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned / ebbd37f0d7f7 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned

<a id="canonical-4ac2d6afb03fb4a4e490ee8f13039ac803f882418fcbea2d6fcffc23fc240bcd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add partitioned.

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
add_partitioned = {}
```

<a id="canonical-c097c5c87d11f6606158e298ebd439a481af9981dafdbb3e083dfbb0581a1481"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned / ebbd37f0d7f7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9ad558a5a03051f6b14df02380691c78b43e7034245630da93f2aba78190852"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_partitioned / ebbd37f0d7f7 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-01107677f3e3820491503fdbae79744fb28cfd7a54a9a127ae6558a07a214671"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52a06c26d2f2cab842ae2100047beb71b96d5b36f46f2f9ee691fcb1d4df3dce"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure / 90ec6c08f28e / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure

<a id="canonical-3846119cf8019aba0dddf71d12de8326a9e63cf1a97b4be71894e93e3448ad46"></a>

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
add_secure = {}
```

<a id="canonical-e155047a26bd451005eeb0b333193e556d80ed3fd4971d3ffd1bec6fd4bb3a93"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure / 90ec6c08f28e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5c782837df810d05d7f1a240dcd95772d712f501acc9b22668541ec124e73617"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.add_secure / 90ec6c08f28e / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-9c982705f56e25117f973fb09d43e5084ae26d48fdd3c3ded526243cceeeedb5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f897f039300171a609489145c6a13be481b8cdb8e2d8c27b5c90207b62ee0207"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain / 2d131bbed406 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain

<a id="canonical-5249f2a8b1f717979a6de522112ce1389114f45c785016a98ffb48d09d58d69d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore domain.

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
ignore_domain = {}
```

<a id="canonical-d9066774b37ad2e2921f5613073c977ee55ed20da8666af8bad8808b4fd120b3"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain / 2d131bbed406 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06bda1817d04d13398b2a2f61310efdddfe5a1cfd3243c5675d36efd39e76fe8"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_domain / 2d131bbed406 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-524c3fd30d7f05668c2e13bf1ba341ba70763bece6546fecc784f2358e7ae34a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e034b16dbfcc4d173172d765df38de0da67ade5e4de35e5b3fc15e179693063b"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry / ec922a95cbc4 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-75b92ea341db100d88df4ea1c2a78ac81c39194084885cf4d97aa05898b36b29"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore expiry.

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
ignore_expiry = {}
```

<a id="canonical-768d30ff716f2c7c0db4e37d07e5695f52662aff0bc7bc9bfb89af80d4d46275"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry / ec922a95cbc4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d60dedcecc3715120623f4e87ce2491bbf045df4f692bfafb9c3d8ca190aeeaf"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_expiry / ec922a95cbc4 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-f9d86a52859ceca815ad5125ceb05e17f48ab62e1e7dc7561c2570f073a83d7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad9d1e4c50ec80124f67d13cb1cabee617c384bc35e1be6224dfd145a7d75076"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly / 766a7f5a4d76 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-3deb605afe342ccd460999e7327c89956952a1ca82389a45a52dd434694a5ee0"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

<a id="canonical-380a2f79b04cfd881410c099627322abfd0f2fb0a9c3fb425a38a3f061ec6699"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly / 766a7f5a4d76 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9db2a58aa6a4bd97b45892d91c7f964d7ab72ffc6c5b22254f24f22456d445bb"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_httponly / 766a7f5a4d76 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a08909d51e648a8163feba9a2883e2a78390a268f50f8bffaf751506131e37c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c21f8a47d78e4dd5bc6b77339796ac3ca8b8bec5c53baa6572fe8ca4ae41e805"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age / 771641e86909 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-093c9895221585b2389d2577c1cdf18dbd97181335eb988c58dacdbaf98b614c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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
ignore_max_age = {}
```

<a id="canonical-787d0400de7ba4887929b639f24d57b3455cb5f9f740b786d9e04a362193e82d"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age / 771641e86909 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f016135e281b1f65d8c9441271cf1b518153f4b5dcad3b1858daec5e0b9abfa2"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_max_age / 771641e86909 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c44c60ad6cbca6561e4d84e99175cb613e1b6d4d4a289dddb74514d527efc4fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bae20d2093ff7028c703b178608a38f4ec52a92ee6bae083a0608663e571c4d1"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned / cba72c1de0cc / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-bcacd9b3a43d9572bc8e074f870497c5a57bd1dc759f926033f9331f686791f9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore partitioned.

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
ignore_partitioned = {}
```

<a id="canonical-9c83442c34e1c0061ab768eae7d88f6dd3f3693dcd53b8319c789d5d6d65169a"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned / cba72c1de0cc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46efd1e128a80217ae3d950a4f39da20ea7fae6377ca2a63ada12c090ab6014d"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_partitioned / cba72c1de0cc / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-8f14d59b4183111a9209da1cb008f5b31ab5fef6270746455aeb68243020cd15"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05ee97ea0b4804b04625f4204def1c0ceea26c5e70170e24320be40a5b39210c"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path / a3305eb9c456 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path

<a id="canonical-a4aa1aa8fa6eb10ab8a5a7c40aa7e2e469b0ad25f7c72dee45afd7679831e689"></a>

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
ignore_path = {}
```

<a id="canonical-e8a3e82176ede2d01df1310264386a13aa063a02b5eb63caf8ade91de309d949"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path / a3305eb9c456 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e59f1ec11cd22eb0094d3873d1a6e7339cc75fe0bf7a3486598efe46259300e"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_path / a3305eb9c456 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-659c1931ba3c6ebd537ab10b39f88f92fcebc47c9f2f7c3f67455a2374bc3819"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29b0b6f450442bd89db379f613e00be9df0a73c2f0f1466f11baf9999f2838c0"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite / c008d699bd5f / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-5e4f1ffd90ea6a6cf38fce7fee2c48736cc2ceaf80c0a16b727c86e4e1cb08b8"></a>

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
ignore_samesite = {}
```

<a id="canonical-e57ef419011ac5e8553135e3c6fedc26343818bce5ea0f83b479563d9d026a75"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite / c008d699bd5f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce6a1a025eeb22962ba8d03bb08193fd9bda27e962a19c2df4497c939746192d"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_samesite / c008d699bd5f / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a7e75c90e86fe4f8258cf911bb0dacbef8ec9d00fde604a4b381e7b430382087"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bef60366c2d584e3e08e670e0a3aac6744867d302f9dcd43377586e0da8aff8"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure / c33fcb0474e1 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure

<a id="canonical-caf4a43ab89fe3f5a53a582b611f2240a01a11b9e893fff116d3a45fd9865170"></a>

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
ignore_secure = {}
```

<a id="canonical-57af42845f2520a287784967c395c2142d60f555c05d5214bc4ea360ef97c6b0"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure / c33fcb0474e1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5271d8f75b56aa7655e6b735fb91d50f663e3310828a7e4cac032067e8f040d3"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_secure / c33fcb0474e1 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-7719c84eaa64186ab06f78f529048a8e0f20f7492d5476db906d26912f0d60a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fb4e35680dece080c1ec9bd87f84f5a435cea033327e3ff5f8a1f60dbdb1406"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value / a0b925e1f1a7 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value

<a id="canonical-216babfb31b6427031a2abbef435c7a7d363a01a89ef46d233ff79a7332b7e56"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore value.

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
ignore_value = {}
```

<a id="canonical-73e3846aec603584c5416f7755d9e51b8eca2a24b4f739ddfc090bc4476471df"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value / a0b925e1f1a7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bce182d49dcb6881fbe1c79037934ce09ea3287708a92de59d00513129e4a8b7"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.ignore_value / a0b925e1f1a7 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-1cf64b407579607c907c229c03f89aacce91d8a163f079c487032abb5dc1b787"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0065116d9700a8a9151076a1920c74232d9b8be87dcd6fc36e0299af4455910c"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax / 78ef9e0bb004 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax

<a id="canonical-211515a0979f56519f75d8887041cc9ef369958ab5946fb10ec69556b13df2cc"></a>

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
samesite_lax = {}
```

<a id="canonical-57e88f99247861e25f96cef081df2e9527e9600367e56c7654e61129cded4f63"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax / 78ef9e0bb004 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd732c2f051b1ab45ac53480a27cc8c8d1f3a97cc2d8f0a6708d50bf246da862"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_lax / 78ef9e0bb004 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-25937071e708a4294d08158a2eb50e7f8c9a8d33b88f0d9229c4d9cb12afd3f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-445731ccb5abcc29ac05df81b5ea3bdae10ca25c7fcd3ee8d3866a9bbde18051"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none / bf9558a3cbc7 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none

<a id="canonical-ea78b164544e10f92ae1b476ba5dbcb22a36a61523c650e4103d387692686c14"></a>

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
samesite_none = {}
```

<a id="canonical-33e777bdcde777aaafbba0d92e8d9947a040eff5c1d943f4a41c065ffdcbf174"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none / bf9558a3cbc7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2861aaa3351c4c48485e7cd660686de12b72ea0459bb4fbde423e97f9312938a"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_none / bf9558a3cbc7 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-e5dfd87293400e16dbac119a439230ce800bb6f56210a17e1140694a021416d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c26697459ab5c88364f2fb168d9db7ea16d404c9d18636b948913186dc2315f"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict / 8c57d5eb2d38 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict

<a id="canonical-1e7514915131ca4206d128173c150ae5b89b901cc40bc47e234bd9256361b348"></a>

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
samesite_strict = {}
```

<a id="canonical-d67f1f184ae74d2251a911730d5d474d777a024cf38c54c04e2a248887244139"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict / 8c57d5eb2d38 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e8016835613ca3416330fca1a1f26c0ffd1b040d519527f70cfe61c0b91bcc91"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.samesite_strict / 8c57d5eb2d38 / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-0d93353ca93264df4d71af2016f54f3c660b2763e3869964a84e0ad00ecdc1c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8706ea70aa1a74823e453fcacf32ffb59996f40d833ff9c59bdad459e8a2ad55"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value / c4a7073784ef / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value

<a id="canonical-163822ebf6b7f36adff9bcdb4c85c3a5672defc29b695a5ac48e19f581cab491"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-e22c7e2ece9420634dd41ad44c84a26a00809a88cd293a8107841b65c0f2dccb"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value / c4a7073784ef / 3

- [blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-01308e8c9937a9e270c9b6445192ee6ce145ecfc18fb24e0c08c9d22a04ca73b): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-003.md#canonical-b36885799a060f9d1ebfa2e1d958a7d9514e8f030e62faf40d39e05a6577ced8): complete subsection reference.

<a id="canonical-340ca9dc039a912052be542ecc89c43cabe114e3ce7b802cd309ede0798d9dfa"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value / c4a7073784ef / 4

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-01308e8c9937a9e270c9b6445192ee6ce145ecfc18fb24e0c08c9d22a04ca73b)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-b36885799a060f9d1ebfa2e1d958a7d9514e8f030e62faf40d39e05a6577ced8)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-01308e8c9937a9e270c9b6445192ee6ce145ecfc18fb24e0c08c9d22a04ca73b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-941c676c929b61915239c29371b0110f38911db3d2efa210ad497bee6675eeae"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blind / a83353313ac4 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-0d93353ca93264df4d71af2016f54f3c660b2763e3869964a84e0ad00ecdc1c5)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-0b3db43cd56782b693caa028ec6c89c89ed3a4015756a6d3b1e8d3051e8e7196"></a>

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

<a id="canonical-d1c27a3fd07b1ec2b57862ec5abf3fbed32c05f4c887df4605d8ecd18066bffb"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blind / a83353313ac4 / 3

<a id="canonical-50b8af6b02e45ca0e6c179249b0f43eb5fcdf7c60e4100bccefb65d7bd238306"></a>

<a id="canonical-80d5c82adcb86a72e4aa83965b3f389e06706514762df5a8b6e5647b42915925"></a>

## decryption_provider property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blind / a83353313ac4 / 4

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

<a id="canonical-ffbe7957824c67e59682f43d924d75cd9a2b4e2431f81d0bff88346fba928202"></a>

<a id="canonical-b86c0e891fcfbf5ecc5f993c77de45cccf7c5faae077574652fabf3b0a7b54cc"></a>

## location property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blind / a83353313ac4 / 5

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

<a id="canonical-2a1c6cc731fbb6b48decb91280e32d175b4d9141474c645a582e3d633477cfaf"></a>

<a id="canonical-e19ff41d8ea3a2805a258bcf7888599144869df704af49a3a60ebb34ba3507c5"></a>

## store_provider property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blind / a83353313ac4 / 6

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

<a id="canonical-7ecece70f0ef112c5dafa783ceeedb9f3e4f63a01b092e2303b5fa054f6134d1"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.blind / a83353313ac4 / 7

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-0d93353ca93264df4d71af2016f54f3c660b2763e3869964a84e0ad00ecdc1c5)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-b36885799a060f9d1ebfa2e1d958a7d9514e8f030e62faf40d39e05a6577ced8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a802434931ce06bcf02aa71c44d49bae75b0d87c3dd12b600ec800c1351993d"></a>

## dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear / 17ea5b640463 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add](resources--proxy--reference--group-003.md#canonical-3ee7e4a7978a294a1d891a19f2b6ae21a0889c1900290f5d6e60a051107b1e74)
- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-0d93353ca93264df4d71af2016f54f3c660b2763e3869964a84e0ad00ecdc1c5)
- dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-3ce960d462604f16a0261ff1f4bdead7aba37d038345e12f9b67ff902442d090"></a>

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

<a id="canonical-0061b5af82d48abb65a4d5c629340bdd69c3feb3dfce003941811edde2da365c"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear / 17ea5b640463 / 3

<a id="canonical-a5394f870e19db11a4e3bb91ec029cb220cffc583edd34510e7a84061630dd9f"></a>

<a id="canonical-7f9719ef31bc37c5ce87d212b927810e2da88079c5e3d3955713430253a8e4e5"></a>

## provider_ref property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear / 17ea5b640463 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-128014b0884c6723fc6964605015f62dc8538327748803cb4229f4a4aa6f7acd"></a>

<a id="canonical-2d4d86aaf8955140190f7ab868d3b79d9fda975c64ac47a14c86c55e41b1cc16"></a>

## url property — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear / 17ea5b640463 / 5

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

<a id="canonical-b107d4676b6c44857c9d72533e25d13dd52e5fd4a1481bef79912c1da537a445"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value.clear / 17ea5b640463 / 6

- [dynamic_proxy.https_proxy.more_option.response_cookies_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-0d93353ca93264df4d71af2016f54f3c660b2763e3869964a84e0ad00ecdc1c5)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-1480e105c5589202b3a5e98369bd4d4a75cefc4cc7735b59ff9fd65040f7d5f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cf01dba0d8a68c198d7d6687e2ff8d60b7596926ccb076e3fea4e03b1af612c"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add — dynamic_proxy.https_proxy.more_option.response_headers_to_add / 29f1e683e6a8 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add

<a id="canonical-5f937c3dc2ef3e629d6d71571408731890587b7cd7df22333e695119148f0ef9"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-789ff3a6b00becede9d4a7511e5a1d4228ced613625d7c129069f91c6b3bb3e4"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_headers_to_add / 29f1e683e6a8 / 3

<a id="canonical-d560f4c8bf2ca77e1578d7e7b36ec679de1b57bbf5f1120a86cbcbbc10c728f4"></a>

<a id="canonical-cf12b6be51c45c52365a87bbec24a174a38971530831be342effa41641678c12"></a>

## append property — dynamic_proxy.https_proxy.more_option.response_headers_to_add / 29f1e683e6a8 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-c1b4013dfd19df143e5f8368bcaaacc10a10b06712db8fab3b76bc63a445d713"></a>

<a id="canonical-eeb8f7bfb10126b799bc64b0caae54e3f4901d1e9c9f207cab0f6615646907df"></a>

## name property — dynamic_proxy.https_proxy.more_option.response_headers_to_add / 29f1e683e6a8 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--proxy--reference--group-003.md#canonical-fdf19ca1b1a3edd59989da615d712a66283918b9feb4f5a00493f7c2c446738e): complete subsection reference.

<a id="canonical-9c46de8e39baff6aa8225697014bb20898c00b57e8205b657f7cc9c5433dcd18"></a>

<a id="canonical-5d8d3a43bacbabe9ccbe173c2befc6f1d8af708cf061fd0428f1e1d1dba19792"></a>

## value property — dynamic_proxy.https_proxy.more_option.response_headers_to_add / 29f1e683e6a8 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

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

<a id="canonical-028d803df249c830bf216f53e353df307adc504c9e19cfeb65fc21157b13b1a8"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_headers_to_add / 29f1e683e6a8 / 7

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-fdf19ca1b1a3edd59989da615d712a66283918b9feb4f5a00493f7c2c446738e)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-fdf19ca1b1a3edd59989da615d712a66283918b9feb4f5a00493f7c2c446738e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e77f4c3a4b63410a516a20c4f91620e2702dcba07f011cb988b773611d971213"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value / 26c42c2c32c7 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-1480e105c5589202b3a5e98369bd4d4a75cefc4cc7735b59ff9fd65040f7d5f4)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value

<a id="canonical-90a95e297fd9b7efb3c44a456b243b49fc24e2490459353b37b646b50b16fd1f"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-adfb6e94c967796b66ff867fb667af8fabb09d2b692a15cbf5f79998a7599c78"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value / 26c42c2c32c7 / 3

- [blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-26501ae5300229143174a083765eb79c72b1b0e83cfec790daee41d3352ecd27): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-003.md#canonical-fbfc1563d60f652139ee5a013ed090a11015bf1569ae435c188cd8a65e60dac7): complete subsection reference.

<a id="canonical-776c7f3bb1261edb470f15373fc7b038deba490bb4beb7c72fecbf122dc86ed6"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value / 26c42c2c32c7 / 4

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-26501ae5300229143174a083765eb79c72b1b0e83cfec790daee41d3352ecd27)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info](resources--proxy--reference--group-003.md#canonical-fbfc1563d60f652139ee5a013ed090a11015bf1569ae435c188cd8a65e60dac7)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-1480e105c5589202b3a5e98369bd4d4a75cefc4cc7735b59ff9fd65040f7d5f4)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-26501ae5300229143174a083765eb79c72b1b0e83cfec790daee41d3352ecd27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72c22ef91737d133ed67a58dc0cfbfda1850bb112beea470adf26a8603b798f3"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blind / 499dc79acff8 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-1480e105c5589202b3a5e98369bd4d4a75cefc4cc7735b59ff9fd65040f7d5f4)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-fdf19ca1b1a3edd59989da615d712a66283918b9feb4f5a00493f7c2c446738e)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-83bf96cdc6c668f5cff0f6fd40335bc037ddff0caaeb98d2dc6c8785a9202cd7"></a>

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

<a id="canonical-681c35b9f7d5a9a0b2af828577118b113a451d1f0bc57d6a5c9a763d4797cb5e"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blind / 499dc79acff8 / 3

<a id="canonical-c75d8edfae2841d22066d20ac322fe75661fae1a8516f023f084a0501f0db3a2"></a>

<a id="canonical-df909c3c511688331ec8add1797a5dfbedb322c15517e8103e5753583ed06657"></a>

## decryption_provider property — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blind / 499dc79acff8 / 4

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

<a id="canonical-6c14227b53bd25801135b2cbe2579f2de0e7245afffd5463238cd037d2925408"></a>

<a id="canonical-40adbf54eb5ceaa804f56cb6af999612c3215b3a74133916d5b494c47157d2db"></a>

## location property — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blind / 499dc79acff8 / 5

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

<a id="canonical-e15beb5ddb5f3560f7bd625fd7e86119c3dfd1a66aa7407e394506f22759d4ee"></a>

<a id="canonical-18fdb3630a275ea0af7b2836211c22036777585c6108921537b3b717f79c2c83"></a>

## store_provider property — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blind / 499dc79acff8 / 6

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

<a id="canonical-c8faa058c106849687338feab6333fcda78fcc98e61251869a851a02f19c075e"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.blind / 499dc79acff8 / 7

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-fdf19ca1b1a3edd59989da615d712a66283918b9feb4f5a00493f7c2c446738e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-fbfc1563d60f652139ee5a013ed090a11015bf1569ae435c188cd8a65e60dac7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fa0bd788c6306d91bf3d94875eaca2579fea4598e11da6099218d73fb4afd2a"></a>

## dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear / 9b48272a54bb / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.more_option](resources--proxy--reference--group-002.md#canonical-50b27dc01c842c8e898a7ced894bb1761bf1277ed7e71b3dc896a9756bd74f22)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add](resources--proxy--reference--group-003.md#canonical-1480e105c5589202b3a5e98369bd4d4a75cefc4cc7735b59ff9fd65040f7d5f4)
- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-fdf19ca1b1a3edd59989da615d712a66283918b9feb4f5a00493f7c2c446738e)
- dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3eaf5caff2dc056b35ad80e1a43faf6aeb735f46df523b67fe50d2789e556cf6"></a>

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

<a id="canonical-56ce2e1064852b3bb43f177c7246553a9bdff2bb0aec3f15e603f0156505967c"></a>

## Direct properties — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear / 9b48272a54bb / 3

<a id="canonical-bf153412e0676c4ab9701fcdc32ce0768434827a00f158f8e47250726f3be66e"></a>

<a id="canonical-66a97a65004f5da7264f8bdd08a73e4b52b91b2c3d4e701f133ae33039c90dce"></a>

## provider_ref property — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear / 9b48272a54bb / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-f9b627d34bbfc4df0ae4650650627e0ebfe00883456d3b917f5af417db36f82b"></a>

<a id="canonical-d65651976b38195cd68251c0e161b8dab6586bb2744c653ce8f1b51c35a64ca4"></a>

## url property — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear / 9b48272a54bb / 5

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

<a id="canonical-f3c6e397bab208712f08acda2066896a53fe7af0255b3dc0fca481d9b971f959"></a>

## Next pages — dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value.clear / 9b48272a54bb / 6

- [dynamic_proxy.https_proxy.more_option.response_headers_to_add.secret_value](resources--proxy--reference--group-003.md#canonical-fdf19ca1b1a3edd59989da615d712a66283918b9feb4f5a00493f7c2c446738e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58cc821a67f83ecc62b6d079bde6a39411c89643e914f83aae37b95c00682698"></a>

## dynamic_proxy.https_proxy.tls_params — dynamic_proxy.https_proxy.tls_params / 72c54868f901 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- dynamic_proxy.https_proxy.tls_params

<a id="canonical-210a2b6bbc6ab27cc4e75722a9f5117d0d44ffe2b7420767729d96a1e30418c7"></a>

Type: `"object"`. single nested block, Optional.

Inline TLS Parameters. Inline TLS parameters.

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
tls_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-da781176ade3470b4e3c6da5a8184ca0577b2ea27536876b00145b830a53751d"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params / 72c54868f901 / 3

- [no_mtls](resources--proxy--reference--group-003.md#canonical-e9d2b01e8b66baf46d9e34fbbb9f8e9594b3ec0b8284c9da09cf0c6ed660e50c): complete subsection reference.

- [tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709): complete subsection reference.

- [tls_config](resources--proxy--reference--group-003.md#canonical-e59934f18e802b6378156696082f874fe63188677b49d7fea450a2f16e7d8edc): complete subsection reference.

- [use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e): complete subsection reference.

<a id="canonical-9f6b9726fa5c8854ccbecfc7212d5227ef514e72529f752d2ab480dd920bd7ad"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params / 72c54868f901 / 4

- [dynamic_proxy.https_proxy.tls_params.no_mtls](resources--proxy--reference--group-003.md#canonical-e9d2b01e8b66baf46d9e34fbbb9f8e9594b3ec0b8284c9da09cf0c6ed660e50c)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709)
- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-e59934f18e802b6378156696082f874fe63188677b49d7fea450a2f16e7d8edc)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-e9d2b01e8b66baf46d9e34fbbb9f8e9594b3ec0b8284c9da09cf0c6ed660e50c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ade2aa57ddc46fa61b4f78a2d00be959574866fd0955268e565555f483e9c81"></a>

## dynamic_proxy.https_proxy.tls_params.no_mtls — dynamic_proxy.https_proxy.tls_params.no_mtls / 586295d11963 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- dynamic_proxy.https_proxy.tls_params.no_mtls

<a id="canonical-d4ee163eb7179b8e39359ce162b32cb47034c9229ce64d7c7f3d3b30a6e128d4"></a>

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

<a id="canonical-0ae5e43c114fe9ef223a50f4f2965518dd20fd96b7580aca3350168390b89554"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.no_mtls / 586295d11963 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e246c1c23707734a429b74199ee8fec823b6f24daa812392680ab23d6ad1d17a"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.no_mtls / 586295d11963 / 4

- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da46d978fc1804480ef15ccf19b610f266b46b6eedcffd3728ef83873ab71ff0"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates — dynamic_proxy.https_proxy.tls_params.tls_certificates / 89f35b69a349 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- dynamic_proxy.https_proxy.tls_params.tls_certificates

<a id="canonical-8d617c6d8034af814ac995ac0c857b181cd0248ca71e1a0a3359df78bbe950f1"></a>

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

<a id="canonical-f0e24f9f6ce4eb0a0db4742e5d5cdae0c4bf6f1b1777f75f42a067c1ae215209"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_certificates / 89f35b69a349 / 3

<a id="canonical-bbf89f9ace7b6a1a09c43d8e18ec2619d1fe15691781061ded5a47e7dbed29ed"></a>

<a id="canonical-53c95a9054f3189522813565d81712671164428edd5a632d8aa1ec78f94a232b"></a>

## certificate_url property — dynamic_proxy.https_proxy.tls_params.tls_certificates / 89f35b69a349 / 4

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

- [custom_hash_algorithms](resources--proxy--reference--group-003.md#canonical-7f57fdecd9ef64b276e248d17953c5fa3b6f9084d54542e916eb900540d450a1): complete subsection reference.

<a id="canonical-6f3a3eec996d9c52528dfae64263200020be520e7a067cccf9746caa005caa24"></a>

<a id="canonical-153861d44de16fb46d7a2617e3a5a27e40368cd319411f9e9269ba663b7faa7b"></a>

## description_spec property — dynamic_proxy.https_proxy.tls_params.tls_certificates / 89f35b69a349 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--proxy--reference--group-003.md#canonical-1967bb1cad00555ed549fc059b2ae8a558a112f8bdd19ce7498b889d69c49919): complete subsection reference.

- [private_key](resources--proxy--reference--group-003.md#canonical-9e9ca17d7e2b30383234dea45fb34f7d8cb14eb8eb3c65623bcbc41953f79dce): complete subsection reference.

- [use_system_defaults](resources--proxy--reference--group-003.md#canonical-6c1f09e8e782c54853aa89aa5b6ff906a33e456b73d102506484794917e1d4ba): complete subsection reference.

<a id="canonical-6e246a2d1a67e2376e5ea0f0c9571b9abb56debc781035e0802de3b668037e12"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_certificates / 89f35b69a349 / 6

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms](resources--proxy--reference--group-003.md#canonical-7f57fdecd9ef64b276e248d17953c5fa3b6f9084d54542e916eb900540d450a1)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling](resources--proxy--reference--group-003.md#canonical-1967bb1cad00555ed549fc059b2ae8a558a112f8bdd19ce7498b889d69c49919)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-9e9ca17d7e2b30383234dea45fb34f7d8cb14eb8eb3c65623bcbc41953f79dce)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults](resources--proxy--reference--group-003.md#canonical-6c1f09e8e782c54853aa89aa5b6ff906a33e456b73d102506484794917e1d4ba)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-7f57fdecd9ef64b276e248d17953c5fa3b6f9084d54542e916eb900540d450a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-73fd7f8e7584fa3c0751395190893031aef9f057b36562351c035b7b2ee7a2c3"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms — dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms / 13a44a21c100 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms

<a id="canonical-df13d7bbd0b2046d83be6150841a67ef7e5d540b828fa9148fe095342293717d"></a>

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

<a id="canonical-b73775ad7370ec8d970d7c9aaec188e9e8401af9e569b25f547b429b67bd8d79"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms / 13a44a21c100 / 3

<a id="canonical-635a2c821ff8710379e89da9331a1d6088449c1d352d52d2cdbab199ff31b72d"></a>

<a id="canonical-9a26330f4d1531dee71769adeecc101200337fdeb974873298562902a2b33f47"></a>

## hash_algorithms property — dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms / 13a44a21c100 / 4

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

<a id="canonical-6487658143df1b4643dc8ef2fc8500d1714e3fab20de7ce2ecf09f7489aef19c"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_certificates.custom_hash_algorithms / 13a44a21c100 / 5

- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-1967bb1cad00555ed549fc059b2ae8a558a112f8bdd19ce7498b889d69c49919"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fc9ead38bcdd9ca06c775ec47cba5874f57b323ade56af947b1e49301e84cf6"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling — dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling / 3cbe57a48875 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-b22f6656a5c5e8c9ddd5ba7633a1e2de2a114f29adcca63d3b3ee7c9f1b059be"></a>

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

<a id="canonical-1dbe1b887300d9f787addfacde0cff7abe8970604fc02eb37344baca99e98411"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling / 3cbe57a48875 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ce40f09519e43df6ec3d12956a5a164b95af4f7120dd25ea0bc9937a6745d82"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_certificates.disable_ocsp_stapling / 3cbe57a48875 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-9e9ca17d7e2b30383234dea45fb34f7d8cb14eb8eb3c65623bcbc41953f79dce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-654b2de77881cb4b375701317620b81e802773377efd0317e6fad9f841740281"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key / c0e659d0e9c4 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key

<a id="canonical-dcccae23044f069608dde72b6783fd21b0c5fdf8eacf8ed76b858c16bd63befb"></a>

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

<a id="canonical-40368a25b1e5ee91744d4e184d35b0699d65c69d03c70187db4fbcda00774ba4"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key / c0e659d0e9c4 / 3

- [blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-6d77df683e655488d934c4a11f822c40cf2fc25b1a3fa9e3042d30a20b2997ee): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-003.md#canonical-03127a925d22ce4130e0d7f673e9ecd39d43b788031e6243cd1738e9822745a7): complete subsection reference.

<a id="canonical-0a31d0b0aecc7d041c997b05bcf462ac69bb0c77742fd5f920279d05f7781672"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key / c0e659d0e9c4 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info](resources--proxy--reference--group-003.md#canonical-6d77df683e655488d934c4a11f822c40cf2fc25b1a3fa9e3042d30a20b2997ee)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info](resources--proxy--reference--group-003.md#canonical-03127a925d22ce4130e0d7f673e9ecd39d43b788031e6243cd1738e9822745a7)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-6d77df683e655488d934c4a11f822c40cf2fc25b1a3fa9e3042d30a20b2997ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf987d4e7b30652a7406ecb31b3cafdd4a1fca56037238e86cbd5fe089c6d1c4"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secr / da3351a0cef0 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-9e9ca17d7e2b30383234dea45fb34f7d8cb14eb8eb3c65623bcbc41953f79dce)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-4a9a58c5d0df485cd935099ee0f31a8756a0e62ae38a021f9ed06a94592136fd"></a>

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

<a id="canonical-a339777adc75267d6836694d5e25dc9cf494367cf6699bd95231eaf456525309"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secr / da3351a0cef0 / 3

<a id="canonical-87389e7338f5dd090043421e82afec5d12636b39b824a7a9605ec159559bd9f4"></a>

<a id="canonical-8899b31999cc74ca701242aff7a4ed3fd14a3369f8a285605693eba0bc5f8429"></a>

## decryption_provider property — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secr / da3351a0cef0 / 4

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

<a id="canonical-a6e7938d35de63e50d4a3d0538b87748d4ab30d08111ca6bb74e886fda5dbbb6"></a>

<a id="canonical-0224c40b1e570541d40e9b9751b6b33aa4ab123cf3466673aaaa9d80ac3392e6"></a>

## location property — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secr / da3351a0cef0 / 5

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

<a id="canonical-86a33e660b6c4ab4f396b4fe8b5ca6b0157b7a291ad9ffda643cfa9fae1649d6"></a>

<a id="canonical-ce1321fc7ee5e53f00c0280c5ff400b4af0b4ec2818187b8a20320fba4d22866"></a>

## store_provider property — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secr / da3351a0cef0 / 6

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

<a id="canonical-1fa5e79cdab67c64a19df35d76b69f6cfcb56f6730076f2730b54586a93468bc"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.blindfold_secr / da3351a0cef0 / 7

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-9e9ca17d7e2b30383234dea45fb34f7d8cb14eb8eb3c65623bcbc41953f79dce)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-03127a925d22ce4130e0d7f673e9ecd39d43b788031e6243cd1738e9822745a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d254e04ea2a9a041ee8ef4fc683dcbc3b910f77dda6dd55b0b0acd6fc8e51c59"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_i / 448bdffc5249 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-9e9ca17d7e2b30383234dea45fb34f7d8cb14eb8eb3c65623bcbc41953f79dce)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-dee9c7546407893d8989b54fdeaf964c977e5332ca3349b138922bfb6f5f71b5"></a>

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

<a id="canonical-52e125f4a3318fc05a05a5fe2faa19e3191f0c62fa587b0e7d1a7d0aaab1ef2d"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_i / 448bdffc5249 / 3

<a id="canonical-6ceb5d6ec3da0d7c825792b12fd823cf9b17458060e4158c67351c99ee44972c"></a>

<a id="canonical-c6ea4bdc3201836ed56f5eee4d975c8348e93242859b2df94801f4d1049a80a8"></a>

## provider_ref property — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_i / 448bdffc5249 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-432fc51992236b894af410f0a7e1ca49132ad220b6e6c13af9da14fcdeffd5e4"></a>

<a id="canonical-8478a09fad3008f10b3470cdb10c72351e51df1abe018498d4c5007a6051d3c9"></a>

## url property — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_i / 448bdffc5249 / 5

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

<a id="canonical-519a70f7fea2d26d902432c60b0fd61720a5a807c974994cdb439163127ff5d1"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key.clear_secret_i / 448bdffc5249 / 6

- [dynamic_proxy.https_proxy.tls_params.tls_certificates.private_key](resources--proxy--reference--group-003.md#canonical-9e9ca17d7e2b30383234dea45fb34f7d8cb14eb8eb3c65623bcbc41953f79dce)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-6c1f09e8e782c54853aa89aa5b6ff906a33e456b73d102506484794917e1d4ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18fbaad86aaa0d6e5b643154d3174211f1e5fc057823b69c21f178441bbc35ff"></a>

## dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults — dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults / bb4ea115728d / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709)
- dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults

<a id="canonical-3e692bf0fd084f1a44772e38b6c08e52b7118770201cc5b01fb3401a3eab0e7a"></a>

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

<a id="canonical-562112b500e5c327b3a6d636f41d07f44498f473fc44f082d162e3bd31567db0"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults / bb4ea115728d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-38568cf06fcba8f94fd05a65b4fbc228deee72b866f2f84a9f26d9c0a69a2fbb"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_certificates.use_system_defaults / bb4ea115728d / 4

- [dynamic_proxy.https_proxy.tls_params.tls_certificates](resources--proxy--reference--group-003.md#canonical-9b0c107cce688739a579ce6e5aa5a2c780e34fb7018cdcd17d8a3a41fccf9709)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-e59934f18e802b6378156696082f874fe63188677b49d7fea450a2f16e7d8edc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-558bac649f2e0f8b3768e3b61530a268c0048d98e337ff66d8ae703c736a0e26"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config — dynamic_proxy.https_proxy.tls_params.tls_config / a85fc58eecb3 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- dynamic_proxy.https_proxy.tls_params.tls_config

<a id="canonical-6d4bebf34276d1a16a9dd763a73dfdfe62fccfa4b12925f3188e875e187b6bcd"></a>

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

<a id="canonical-7cbb81b4b82cc00f89b0c0fdaadf57ea8f57a302ed23127cc71e7db5fcb6e29b"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_config / a85fc58eecb3 / 3

- [custom_security](resources--proxy--reference--group-003.md#canonical-c99930f133b0ac9028f230e4450fbff86f6a2ebe2f455dccec6370557be3d7ce): complete subsection reference.

- [default_security](resources--proxy--reference--group-003.md#canonical-b3548959111fc5edcf2614b3d8337220d5341e69636eb6f851372e7cfc53749f): complete subsection reference.

- [low_security](resources--proxy--reference--group-003.md#canonical-c50f32e7d1191ecaaa2c28b298582940363eb135648b7dd06f627314ca833d25): complete subsection reference.

- [medium_security](resources--proxy--reference--group-003.md#canonical-abe133134d2190c6414ac9d1c6a3164967ca6fc82434c3bef209f06dd73ecb65): complete subsection reference.

<a id="canonical-e69e0a524bbd83a36c41177bd7ea63291508dd93f90f1b07120d8ddb0b3df732"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_config / a85fc58eecb3 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config.custom_security](resources--proxy--reference--group-003.md#canonical-c99930f133b0ac9028f230e4450fbff86f6a2ebe2f455dccec6370557be3d7ce)
- [dynamic_proxy.https_proxy.tls_params.tls_config.default_security](resources--proxy--reference--group-003.md#canonical-b3548959111fc5edcf2614b3d8337220d5341e69636eb6f851372e7cfc53749f)
- [dynamic_proxy.https_proxy.tls_params.tls_config.low_security](resources--proxy--reference--group-003.md#canonical-c50f32e7d1191ecaaa2c28b298582940363eb135648b7dd06f627314ca833d25)
- [dynamic_proxy.https_proxy.tls_params.tls_config.medium_security](resources--proxy--reference--group-003.md#canonical-abe133134d2190c6414ac9d1c6a3164967ca6fc82434c3bef209f06dd73ecb65)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c99930f133b0ac9028f230e4450fbff86f6a2ebe2f455dccec6370557be3d7ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9d580f9d00fe6e1631b8372c1347a6e08ce961156f1eb8cdac293391a813913"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.custom_security — dynamic_proxy.https_proxy.tls_params.tls_config.custom_security / 355e7260bc93 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-e59934f18e802b6378156696082f874fe63188677b49d7fea450a2f16e7d8edc)
- dynamic_proxy.https_proxy.tls_params.tls_config.custom_security

<a id="canonical-1a8e2799d27d0192cffcc11864eae6f38f7ccab2a9bef1d40c32700fe965f176"></a>

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

<a id="canonical-abf18b2fd81f14108bc0e866896629a92560039cf0e4e97990093bc388b26950"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_config.custom_security / 355e7260bc93 / 3

<a id="canonical-e252a7f86d6532538972065467acf7513f93010316b503da7533d63208654385"></a>

<a id="canonical-ae2a35a37c823506160729ff68203026e6a461336d767becef8da39062093733"></a>

## cipher_suites property — dynamic_proxy.https_proxy.tls_params.tls_config.custom_security / 355e7260bc93 / 4

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

<a id="canonical-81cea1944cd8217f7a9aa079212b5cd0489982fe601f72321cf9f853e2ee72db"></a>

<a id="canonical-919d768efa8d9b789362ccb5021515e3aa42d79ff9130ba500a7e6321482bb9b"></a>

## max_version property — dynamic_proxy.https_proxy.tls_params.tls_config.custom_security / 355e7260bc93 / 5

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

<a id="canonical-1d5c797d9cf6f9cdee9f139aef7d9ecd6d73476b0a20e38d083b04ae92e1cbf4"></a>

<a id="canonical-adc2785a80e41e65601f220ef27112f3407e438d361cc1722030ff516fbef625"></a>

## min_version property — dynamic_proxy.https_proxy.tls_params.tls_config.custom_security / 355e7260bc93 / 6

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

<a id="canonical-e0740ad1a89e8e114cc641efca009ea71cf839caf7c2446287b3e84ef737b555"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_config.custom_security / 355e7260bc93 / 7

- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-e59934f18e802b6378156696082f874fe63188677b49d7fea450a2f16e7d8edc)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-b3548959111fc5edcf2614b3d8337220d5341e69636eb6f851372e7cfc53749f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5cc401f42524b91a6a55bb8432e576c36e28e8d4a6d924a127774e5e0215139c"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.default_security — dynamic_proxy.https_proxy.tls_params.tls_config.default_security / 03b3650a9175 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-e59934f18e802b6378156696082f874fe63188677b49d7fea450a2f16e7d8edc)
- dynamic_proxy.https_proxy.tls_params.tls_config.default_security

<a id="canonical-ca51b8f33ca6b5e920a8fdd3f8b6b34e37d173536e26cfea158515a58ca605f3"></a>

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

<a id="canonical-98d6b25c83464c16033d05bef602d2c4f5ca592c1c5f4ff9fbe3e39f2a9a2dd8"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_config.default_security / 03b3650a9175 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c087ad98d0608c2a5a9634ac193545603514d98e799dd6917a516ddab63241c"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_config.default_security / 03b3650a9175 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-e59934f18e802b6378156696082f874fe63188677b49d7fea450a2f16e7d8edc)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-c50f32e7d1191ecaaa2c28b298582940363eb135648b7dd06f627314ca833d25"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0808a34cdd82f9d050536dae1e481b23a1dfb812ba247caf8362e121512be29"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.low_security — dynamic_proxy.https_proxy.tls_params.tls_config.low_security / 295c73a4beed / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-e59934f18e802b6378156696082f874fe63188677b49d7fea450a2f16e7d8edc)
- dynamic_proxy.https_proxy.tls_params.tls_config.low_security

<a id="canonical-c4264a864d4098eb915947defbf9c623bdfa2c498e6a644ebf9828bc69287251"></a>

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

<a id="canonical-25aa907efeed4c608335c12270c71300a78f5d08cb533ba4c0aa55a47a5d5549"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_config.low_security / 295c73a4beed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40bb5457a28cb14e6327b52e43301a6b91a7fda0ff801fc05a1178dd87a4a553"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_config.low_security / 295c73a4beed / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-e59934f18e802b6378156696082f874fe63188677b49d7fea450a2f16e7d8edc)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-abe133134d2190c6414ac9d1c6a3164967ca6fc82434c3bef209f06dd73ecb65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec65b4fa4615600bb44582ed68302fe7cae6033745526284b3be690fcecb0d91"></a>

## dynamic_proxy.https_proxy.tls_params.tls_config.medium_security — dynamic_proxy.https_proxy.tls_params.tls_config.medium_security / e13a15b0a024 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-e59934f18e802b6378156696082f874fe63188677b49d7fea450a2f16e7d8edc)
- dynamic_proxy.https_proxy.tls_params.tls_config.medium_security

<a id="canonical-ac0239c0589bb69addaa83679afde58acf6d7392534a3c505f396da836a9e925"></a>

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

<a id="canonical-71441d795eaa76a066ca3d829726bb8f02565304e486a355d4b37b58c5cbab85"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.tls_config.medium_security / e13a15b0a024 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-120aeb758f6af1d93327692491126e397c69cf0fb6722e0b9ab8e2d8a8521fc9"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.tls_config.medium_security / e13a15b0a024 / 4

- [dynamic_proxy.https_proxy.tls_params.tls_config](resources--proxy--reference--group-003.md#canonical-e59934f18e802b6378156696082f874fe63188677b49d7fea450a2f16e7d8edc)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f38f66c36e2b531563834676199cff707db610149cc8659194e2e0128e601d4"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls — dynamic_proxy.https_proxy.tls_params.use_mtls / bdfc8c0c6ecb / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- dynamic_proxy.https_proxy.tls_params.use_mtls

<a id="canonical-67b03affb463ffe6ea8407ba256cd6efb0d61b17b078ed1e3b4dad05e551528c"></a>

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

<a id="canonical-99a0f490f55763c9be3979638e3be6da12b7001d9c854f11967cb4c2433559c1"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.use_mtls / bdfc8c0c6ecb / 3

<a id="canonical-a7f6f0ee5b05ace772633f74904970559795c2c4d8556a20cf930a1a10192a4b"></a>

<a id="canonical-5c075a7bfdee3ed5b0101579b5a8ded3fb5d7085ddba8dc9ef071afba3f65b4c"></a>

## client_certificate_optional property — dynamic_proxy.https_proxy.tls_params.use_mtls / bdfc8c0c6ecb / 4

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

- [crl](resources--proxy--reference--group-003.md#canonical-01eb0206a298a6045dba19d000c43ed55bf371f8546701f148036afba0961f0c): complete subsection reference.

- [no_crl](resources--proxy--reference--group-003.md#canonical-a66e3c6bb36f2fb14f207e6526d556cf9002023058f1f2b98c2d9edf720ecc98): complete subsection reference.

- [trusted_ca](resources--proxy--reference--group-003.md#canonical-accc141e9005a34d49e9ad19bc2b5cd7ec9cfb1cb236904c98f552c4184e9b8e): complete subsection reference.

<a id="canonical-a5a568571cb10bf1be576b577c6bfc835a5418b265546c2f339cb8a97f03c3e1"></a>

<a id="canonical-09548b896601b6df8d116a1b079491e8b5f497437644cd889530740fbf03af96"></a>

## trusted_ca_url property — dynamic_proxy.https_proxy.tls_params.use_mtls / bdfc8c0c6ecb / 5

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

- [xfcc_disabled](resources--proxy--reference--group-004.md#canonical-f1c470c6447008992c3730d09548f77620e03c2e703777ad78af0a410626b77a): complete subsection reference.

- [xfcc_options](resources--proxy--reference--group-004.md#canonical-063cbc226b3095e04be5ff936f17986f6d6a54844d01ff228964e1712f8e24ca): complete subsection reference.

<a id="canonical-37dca7834053079e8054b601a1b1d009539c49a740468a57fde741879e873a06"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.use_mtls / bdfc8c0c6ecb / 6

- [dynamic_proxy.https_proxy.tls_params.use_mtls.crl](resources--proxy--reference--group-003.md#canonical-01eb0206a298a6045dba19d000c43ed55bf371f8546701f148036afba0961f0c)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl](resources--proxy--reference--group-003.md#canonical-a66e3c6bb36f2fb14f207e6526d556cf9002023058f1f2b98c2d9edf720ecc98)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca](resources--proxy--reference--group-003.md#canonical-accc141e9005a34d49e9ad19bc2b5cd7ec9cfb1cb236904c98f552c4184e9b8e)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_disabled](resources--proxy--reference--group-004.md#canonical-f1c470c6447008992c3730d09548f77620e03c2e703777ad78af0a410626b77a)
- [dynamic_proxy.https_proxy.tls_params.use_mtls.xfcc_options](resources--proxy--reference--group-004.md#canonical-063cbc226b3095e04be5ff936f17986f6d6a54844d01ff228964e1712f8e24ca)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-01eb0206a298a6045dba19d000c43ed55bf371f8546701f148036afba0961f0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a366898f437df3e55ad141c8f11b61602fcd3230f26a6d3942e4ad8161786b6a"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.crl — dynamic_proxy.https_proxy.tls_params.use_mtls.crl / 123a3cac9bd3 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e)
- dynamic_proxy.https_proxy.tls_params.use_mtls.crl

<a id="canonical-c0becc407463f445eab96c70fc13c3101ab758193b6784874355dec25983921c"></a>

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

<a id="canonical-3518673cac89a53ea2fc4d6afc9ef12983d95bbf30130e7e357324ab02725c05"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.use_mtls.crl / 123a3cac9bd3 / 3

<a id="canonical-901a7d42d657db806eb7d4837ed04ce7775ae5b492e9191963532e265bdaa396"></a>

<a id="canonical-75d2bb97ecb998e201631f97b5e2038aef09bdc187986c826e93e852881ca188"></a>

## name property — dynamic_proxy.https_proxy.tls_params.use_mtls.crl / 123a3cac9bd3 / 4

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

<a id="canonical-851cffba861caf64ae73be0f0cdb491dc97c344e031eda4c1a1134cb1cf03567"></a>

<a id="canonical-7e60c4f4a36148b5559ad86a0d2042d2d028343a3397d9f11857c4840b21075a"></a>

## namespace property — dynamic_proxy.https_proxy.tls_params.use_mtls.crl / 123a3cac9bd3 / 5

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

<a id="canonical-b1829e42a784cc052c5a33e2b4e0e29924fb8d062faceac27d7fddb73cf9af0a"></a>

<a id="canonical-9a4025c9bec26f98fd89437d83c6dab92efac3adf7f28c0c3fb8e3996ec1a4e6"></a>

## tenant property — dynamic_proxy.https_proxy.tls_params.use_mtls.crl / 123a3cac9bd3 / 6

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

<a id="canonical-28c7dc6ab6f30f4c6cc1017b4e76a63871b46ce62efff2ac3013e3a32c2e9c89"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.use_mtls.crl / 123a3cac9bd3 / 7

- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-a66e3c6bb36f2fb14f207e6526d556cf9002023058f1f2b98c2d9edf720ecc98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1b94193df5a4cdfbcafff2cc904aac4c5ee32baba6cbb781f57d3f3b5e846ed"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl — dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl / 000e3bdc77c9 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e)
- dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl

<a id="canonical-06f895f01179ae846f56369defe4522aa89d195f8a1fbc5c2340cb06c88afe07"></a>

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

<a id="canonical-8719e6d43bd712b2ab27c6b47fdee57fb9f3789f50d6e509ff9f47dc513cef28"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl / 000e3bdc77c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9ac22d7eca01d945e9d6532cb1b191bf96236ef6a341d2a03658ae59aa4c0a02"></a>

## Next pages — dynamic_proxy.https_proxy.tls_params.use_mtls.no_crl / 000e3bdc77c9 / 4

- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e)
- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)

<a id="canonical-accc141e9005a34d49e9ad19bc2b5cd7ec9cfb1cb236904c98f552c4184e9b8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8d3c05cb88edce456da4c2ae8815e8ebdcd78a877a267868afc5b20506cf30d"></a>

## dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca — dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca / 3fb151a0643d / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-1ec151ccef7a08b4aff97a128da90e78232721fbd549471bff171f64b457edfe)
- [Property reference](resources--proxy--reference--group-001.md#canonical-7edbc5a232bf8f6e801fa159034e2b29d978d133aa2468f12374f845fac70943)
- [dynamic_proxy](resources--proxy--reference--group-001.md#canonical-97896999452925eeecf5ea225ba0d919d1883998e132ff8db5bb1981d088e5fc)
- [dynamic_proxy.https_proxy](resources--proxy--reference--group-002.md#canonical-0f795e2f85ead5e288ed524a7115868e1c2d80026d444d766df22012f10d49a3)
- [dynamic_proxy.https_proxy.tls_params](resources--proxy--reference--group-003.md#canonical-22b339fd009a99d1165b2bf665b38c559d52ba6821e32b7ed8a3cdbc6a380da8)
- [dynamic_proxy.https_proxy.tls_params.use_mtls](resources--proxy--reference--group-003.md#canonical-58c9885c8921697686643f4f627592eb92f0a7ad1b3f22a0d4c06281d7e0c90e)
- dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca

<a id="canonical-917fc74064c3486398f3fe2ae8df819e0f44d7b8c8cb4be23f5e1d65c5d01e23"></a>

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

<a id="canonical-1c5f7d90a3b96a9c0a2350f25c794a433c5f0c2befa6ff6f138b437359a0022b"></a>

## Direct properties — dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca / 3fb151a0643d / 3

<a id="canonical-57ac69d876804f8f73ca823373e0b8bf7f79cd7f34e5ca07e0cda906d2f5c2e6"></a>

<a id="canonical-4fce79b2c17fa316e2e82022c54405e97018fcb397450d5ae527217b66030ba5"></a>

## name property — dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca / 3fb151a0643d / 4

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

<a id="canonical-88e02fbbd954b148188b9aa8ca725c1df5c887a6dad4916af9d6c99bca8c3850"></a>

<a id="canonical-2da11a107a3ec8a5646626ba6aec116b9074ec4b05c8e3a1526df622ebf1948b"></a>

## namespace property — dynamic_proxy.https_proxy.tls_params.use_mtls.trusted_ca / 3fb151a0643d / 5

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

<a id="canonical-9d8de6f07dd5749f0ce6a398049b23f142b0aa8feca65f7ec7add4794ce40566"></a>
