---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-64772fe03a3005f3bfd41028f73b6f03d778360cecdb8294d32d25fe8845ad08"></a>

## more_option.request_headers_to_add.secret_value.blindfold_secret_info — more_option.request_headers_to_add.secret_value.blindfold_secret_info / f1812abae078 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-261252cddfc4939b222b92bc32e905eb9810570a7a81260e19ab7d8aa23267c8)
- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-6913d8602665c4129c29c54edd983110ab472ed4798a11d22c1ba9b880eb675e)
- more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-c583a5c2af75a7f0af259d9688d3d0b9af666c7bef9ae757f7e4df6b372898f3"></a>

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

<a id="canonical-535a353f33fb01ed62fcd1f8d251f94ed58e9f4b752d8b568d1d3339b5fe4343"></a>

## Direct properties — more_option.request_headers_to_add.secret_value.blindfold_secret_info / f1812abae078 / 3

<a id="canonical-de7c6166ff851acf7f8051e0a5593a3a76fac92bdb246e2953beca4a006182da"></a>

<a id="canonical-e15c2c22af7d975c8c30e57626ecc0cf7f7bfe8302c6a153536e93b07d958810"></a>

## decryption_provider property — more_option.request_headers_to_add.secret_value.blindfold_secret_info / f1812abae078 / 4

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

<a id="canonical-03e9c9a01678a668b4c1895d0241ce77e32d1375c3e8d1a5d64ce9eaa58af3d7"></a>

<a id="canonical-2e78cdc92330b48613c0a76e53e669cf50d7989198d29ed72d7c602aa67b80ca"></a>

## location property — more_option.request_headers_to_add.secret_value.blindfold_secret_info / f1812abae078 / 5

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

<a id="canonical-809f9ca7c16711379548ba5a51614c2bb61522c04b611c30114538be0c683820"></a>

<a id="canonical-ce358a3e0950e856aba7a85fba64790f49b8fe2119ea21673b3159c3bdc2366d"></a>

## store_provider property — more_option.request_headers_to_add.secret_value.blindfold_secret_info / f1812abae078 / 6

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

<a id="canonical-56317d975bcb3749431bcc50dc82f80ec06fa47f21f8d225c64c7e4794b2364b"></a>

## Next pages — more_option.request_headers_to_add.secret_value.blindfold_secret_info / f1812abae078 / 7

- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-6913d8602665c4129c29c54edd983110ab472ed4798a11d22c1ba9b880eb675e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-dee87c04b906ebacb9bae3622b9fab7344f049f7fa816565b4de3565534a6161"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4f1991e049b0a1f99c9455c6e39e8f689acc6717c60fb23762526d8ec80049f"></a>

## more_option.request_headers_to_add.secret_value.clear_secret_info — more_option.request_headers_to_add.secret_value.clear_secret_info / d2e856b99649 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-020.md#canonical-261252cddfc4939b222b92bc32e905eb9810570a7a81260e19ab7d8aa23267c8)
- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-6913d8602665c4129c29c54edd983110ab472ed4798a11d22c1ba9b880eb675e)
- more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-9f5202d498abc36831416b8be149aab914ebbaae626206be735a964dbcd25b6c"></a>

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

<a id="canonical-de625d062cbaf00c8cce9be2803e79250cdfa4a3267a6967c460dbfd6d81d7d3"></a>

## Direct properties — more_option.request_headers_to_add.secret_value.clear_secret_info / d2e856b99649 / 3

<a id="canonical-85eb2112334884270ee07761ad1a525dd3bd25cfd3a52048f2085b268fafa9d9"></a>

<a id="canonical-daf8db31b887ff65ec71432d2ef2bcba556f5c54d44ba5ac0ff9947b791d174c"></a>

## provider_ref property — more_option.request_headers_to_add.secret_value.clear_secret_info / d2e856b99649 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-63d4fcf0466af35776d064c470f764123da414f711c262341a7a5c57024e5d93"></a>

<a id="canonical-5a403223edae774fb7febce809e724a5b11b76c8856281b925b309ee150e45ef"></a>

## url property — more_option.request_headers_to_add.secret_value.clear_secret_info / d2e856b99649 / 5

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

<a id="canonical-ffff4a8b609a98a57fcc8da71da55b1df3e5b63b89f796b700c3f9328810c528"></a>

## Next pages — more_option.request_headers_to_add.secret_value.clear_secret_info / d2e856b99649 / 6

- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-020.md#canonical-6913d8602665c4129c29c54edd983110ab472ed4798a11d22c1ba9b880eb675e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-14cfba2fad965fc3083c7501a1d8d274dcda649a03615df9f56191389ba2bd4e"></a>

## more_option.response_cookies_to_add — more_option.response_cookies_to_add / 84dd1b11ecd5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- more_option.response_cookies_to_add

<a id="canonical-a4b90105559d76802fe88e86afbb6bdc9d22a99f4bd856b0b025e3322d43f984"></a>

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

<a id="canonical-5d61e8c31d4422e7ba9091d22c64208839f85eed2ec8982a3b718a0d900db9ab"></a>

## Direct properties — more_option.response_cookies_to_add / 84dd1b11ecd5 / 3

<a id="canonical-1cbb263aab1ef942ee34e08578d551a903c6621159bfabb9664ae62087818b05"></a>

<a id="canonical-50bd6f377a71363d68d5bf9f34383cc2ef84bce9543651fddc6a59d6d9e4ccfc"></a>

## add_domain property — more_option.response_cookies_to_add / 84dd1b11ecd5 / 4

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

<a id="canonical-20c20e714935fd1104eec0c9483f8516f6ddab09525e4e97018542aebd720756"></a>

<a id="canonical-1c062c0ac4e8759385c0da258819c5cfd92873fdecb19867600387780cd52f5b"></a>

## add_expiry property — more_option.response_cookies_to_add / 84dd1b11ecd5 / 5

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

- [add_httponly](resources--http_loadbalancer--reference--group-021.md#canonical-42ac3daafb63ef47941f71a034d4d3a0d8d2ce79a8a18814bd024771a0c72b2a): complete subsection reference.

- [add_partitioned](resources--http_loadbalancer--reference--group-021.md#canonical-ce8966ea63f214faa1c64562939c931c36de1bcc174322ec782023dde830e0ec): complete subsection reference.

<a id="canonical-6e4a4bee86a538aa067e710866c4fa47ead9598b33a01bfe1923f790331094f5"></a>

<a id="canonical-324d1ec34672fbd45449cc60bdb04beaacad49dbc7d65c75129e66b1230c8e1b"></a>

## add_path property — more_option.response_cookies_to_add / 84dd1b11ecd5 / 6

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

- [add_secure](resources--http_loadbalancer--reference--group-021.md#canonical-daef6aaf1b2189c2b867f29446d0ef15d4f3a7dca10c07265762486351c389eb): complete subsection reference.

- [ignore_domain](resources--http_loadbalancer--reference--group-021.md#canonical-22ff5c5d860493e5f9bc67314c89b7e24d6cda3b019a06dc93255f577bbef38a): complete subsection reference.

- [ignore_expiry](resources--http_loadbalancer--reference--group-021.md#canonical-d3907f9c1c4886317aebdbad01323992b38158078d399c2c4ab40b3ad6e4ea39): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-021.md#canonical-f86e59d1560487af3a991e1d0c765fda38ac1adab6b1633535cc7e7a94a4bbc3): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--reference--group-021.md#canonical-6c383273aa4f7124daf6aab38333915bc8c9d20927456b112d8c6e50c5a8e6d9): complete subsection reference.

- [ignore_partitioned](resources--http_loadbalancer--reference--group-021.md#canonical-085e2537e7060f8224a0d45a5527ebd5600696a6a51c9b0b90d96762940e713d): complete subsection reference.

- [ignore_path](resources--http_loadbalancer--reference--group-021.md#canonical-a341f20ff2f14c6c8f2cd3ffe2f4555fdd9e6e280bf1431ff67788fc10241511): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-021.md#canonical-d28b394397609b939cb8a665a6357ef9da8406bc1614c401d03556ce4e90ce44): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-021.md#canonical-0223382fd1c9543be0cb07550749b7e0c44a42349b38bd45fbc353ddfc189b6c): complete subsection reference.

- [ignore_value](resources--http_loadbalancer--reference--group-021.md#canonical-d054e4d7acd1269a374c44672b4f110a125083e2ab70e3cef023577a61136894): complete subsection reference.

<a id="canonical-6921439299b27465433c4c2e451c0cc7beff5d10b88624e9f23b67d287b7d994"></a>

<a id="canonical-9b01c9cb9be0bc73d6f42a5b005d83c927625adcdb65f7952cf30d37618dd07b"></a>

## max_age_value property — more_option.response_cookies_to_add / 84dd1b11ecd5 / 7

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

<a id="canonical-69a4643cf11eec54d53df31865bca696adf964ef431fb33c06b3ea7c757c82cd"></a>

<a id="canonical-d03610b04a745d8909667eef88313f512f44d5508f78cb1c12e5c770b0a50f3d"></a>

## name property — more_option.response_cookies_to_add / 84dd1b11ecd5 / 8

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

<a id="canonical-2ca46a0f4101d4e87af2e352585b91b8db2a70507c2c2433eda83d873a434ce4"></a>

<a id="canonical-2ca2f8122f3c6af5d8dbd6c4c3ba06df43ef996d073efc38c120e16bd048f1b0"></a>

## overwrite property — more_option.response_cookies_to_add / 84dd1b11ecd5 / 9

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

- [samesite_lax](resources--http_loadbalancer--reference--group-021.md#canonical-bccff969a23d7fa8583e2f9938f91d86ba4beaca7a1bd9e401b3994e55e2af07): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-021.md#canonical-7e6021da5fb077a8d4774554e9ff72d8672677b0881226e16e72a183675c51a7): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-021.md#canonical-e0d5bcd9ebcbf71287969e74fd350d224515e905750b8766b85475933573e899): complete subsection reference.

- [secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-7a4c8976ce3a00e6109190166b591890a8ef0156a098d16e71e109562acdefb1): complete subsection reference.

<a id="canonical-20483330dfdbc4262a42cef4ce7c1a0b110e21061a15fc1ab7e86d585ec7ba5e"></a>

<a id="canonical-c8e9ae1c4071a91d5b62d6dd53143ca7d486605bf9b6af6695a7fd1d66dc0bc4"></a>

## value property — more_option.response_cookies_to_add / 84dd1b11ecd5 / 10

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

<a id="canonical-ae67b061b6fa36bd10a23012b2be2f9f5ab7faee994cad0a3c37ad9ea7b9a650"></a>

## Next pages — more_option.response_cookies_to_add / 84dd1b11ecd5 / 11

- [more_option.response_cookies_to_add.add_httponly](resources--http_loadbalancer--reference--group-021.md#canonical-42ac3daafb63ef47941f71a034d4d3a0d8d2ce79a8a18814bd024771a0c72b2a)
- [more_option.response_cookies_to_add.add_partitioned](resources--http_loadbalancer--reference--group-021.md#canonical-ce8966ea63f214faa1c64562939c931c36de1bcc174322ec782023dde830e0ec)
- [more_option.response_cookies_to_add.add_secure](resources--http_loadbalancer--reference--group-021.md#canonical-daef6aaf1b2189c2b867f29446d0ef15d4f3a7dca10c07265762486351c389eb)
- [more_option.response_cookies_to_add.ignore_domain](resources--http_loadbalancer--reference--group-021.md#canonical-22ff5c5d860493e5f9bc67314c89b7e24d6cda3b019a06dc93255f577bbef38a)
- [more_option.response_cookies_to_add.ignore_expiry](resources--http_loadbalancer--reference--group-021.md#canonical-d3907f9c1c4886317aebdbad01323992b38158078d399c2c4ab40b3ad6e4ea39)
- [more_option.response_cookies_to_add.ignore_httponly](resources--http_loadbalancer--reference--group-021.md#canonical-f86e59d1560487af3a991e1d0c765fda38ac1adab6b1633535cc7e7a94a4bbc3)
- [more_option.response_cookies_to_add.ignore_max_age](resources--http_loadbalancer--reference--group-021.md#canonical-6c383273aa4f7124daf6aab38333915bc8c9d20927456b112d8c6e50c5a8e6d9)
- [more_option.response_cookies_to_add.ignore_partitioned](resources--http_loadbalancer--reference--group-021.md#canonical-085e2537e7060f8224a0d45a5527ebd5600696a6a51c9b0b90d96762940e713d)
- [more_option.response_cookies_to_add.ignore_path](resources--http_loadbalancer--reference--group-021.md#canonical-a341f20ff2f14c6c8f2cd3ffe2f4555fdd9e6e280bf1431ff67788fc10241511)
- [more_option.response_cookies_to_add.ignore_samesite](resources--http_loadbalancer--reference--group-021.md#canonical-d28b394397609b939cb8a665a6357ef9da8406bc1614c401d03556ce4e90ce44)
- [more_option.response_cookies_to_add.ignore_secure](resources--http_loadbalancer--reference--group-021.md#canonical-0223382fd1c9543be0cb07550749b7e0c44a42349b38bd45fbc353ddfc189b6c)
- [more_option.response_cookies_to_add.ignore_value](resources--http_loadbalancer--reference--group-021.md#canonical-d054e4d7acd1269a374c44672b4f110a125083e2ab70e3cef023577a61136894)
- [more_option.response_cookies_to_add.samesite_lax](resources--http_loadbalancer--reference--group-021.md#canonical-bccff969a23d7fa8583e2f9938f91d86ba4beaca7a1bd9e401b3994e55e2af07)
- [more_option.response_cookies_to_add.samesite_none](resources--http_loadbalancer--reference--group-021.md#canonical-7e6021da5fb077a8d4774554e9ff72d8672677b0881226e16e72a183675c51a7)
- [more_option.response_cookies_to_add.samesite_strict](resources--http_loadbalancer--reference--group-021.md#canonical-e0d5bcd9ebcbf71287969e74fd350d224515e905750b8766b85475933573e899)
- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-7a4c8976ce3a00e6109190166b591890a8ef0156a098d16e71e109562acdefb1)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-42ac3daafb63ef47941f71a034d4d3a0d8d2ce79a8a18814bd024771a0c72b2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df0479d45b3453b03b4498f68b9559a938548301f6db4b639781e007c694f339"></a>

## more_option.response_cookies_to_add.add_httponly — more_option.response_cookies_to_add.add_httponly / 065cc297ee1b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.add_httponly

<a id="canonical-86b69b0d4fda79f545fb117740ff0a923fda1eefa6587892de13e24f1c3b2a67"></a>

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

<a id="canonical-81dec7c8188551ed8a32d5126d8c53af7f021a79769518f6c67f99cbd029241c"></a>

## Direct properties — more_option.response_cookies_to_add.add_httponly / 065cc297ee1b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64ba5e68b8d01e00b815e557fc97c110a8bb0b43f65ecc21dc9f69a2db758ad1"></a>

## Next pages — more_option.response_cookies_to_add.add_httponly / 065cc297ee1b / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ce8966ea63f214faa1c64562939c931c36de1bcc174322ec782023dde830e0ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96a8096da24ac706e560d64e35f9e9ad2a8e168882dba79a0e606869f66ccd57"></a>

## more_option.response_cookies_to_add.add_partitioned — more_option.response_cookies_to_add.add_partitioned / bc8ce8f769ab / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.add_partitioned

<a id="canonical-fd5f9c306d3dc44e86b1a46f671779f6ac039ef1cc32600f253dee0fe44b6725"></a>

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

<a id="canonical-b26b3dfd45220dc49ef7fef671d6af219d6fe8b4351ec73c6f51dc2cef7875ce"></a>

## Direct properties — more_option.response_cookies_to_add.add_partitioned / bc8ce8f769ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5538247067a1a40b4f33a740e0b9f2153b84c655713088af2236e6f4f4c68461"></a>

## Next pages — more_option.response_cookies_to_add.add_partitioned / bc8ce8f769ab / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-daef6aaf1b2189c2b867f29446d0ef15d4f3a7dca10c07265762486351c389eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff7172812706e67facc6a7d1ba35a4adfa41abeb88d90a0c5f936fe1e23a859c"></a>

## more_option.response_cookies_to_add.add_secure — more_option.response_cookies_to_add.add_secure / d339eba49a8d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.add_secure

<a id="canonical-72102311e92f015b55c40f971a995dabad968d9786a9b5201af1acc809c05c6c"></a>

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

<a id="canonical-9f052382aab174d02343d44950122beacd7771668c33b60531893c1aaee0abbe"></a>

## Direct properties — more_option.response_cookies_to_add.add_secure / d339eba49a8d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1cb03730e06cc7154516194ade4c141a1fb091a1a3d279ff692826db4cea89da"></a>

## Next pages — more_option.response_cookies_to_add.add_secure / d339eba49a8d / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-22ff5c5d860493e5f9bc67314c89b7e24d6cda3b019a06dc93255f577bbef38a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-acbfaea09d9ee8864be7284cf554417062e53a1fa330b7df822f95484763a401"></a>

## more_option.response_cookies_to_add.ignore_domain — more_option.response_cookies_to_add.ignore_domain / abf5bc01f462 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.ignore_domain

<a id="canonical-8d6127e97cc27a4bda7d087cd2af590e31e6308ab0af94b3f7c887e889793ddd"></a>

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

<a id="canonical-6ffe529ea44389fbe27baff209d17c249193461c1e7b69a48928f53f20988dc8"></a>

## Direct properties — more_option.response_cookies_to_add.ignore_domain / abf5bc01f462 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c444fba2eb9224c352da65fe0124e83485378b598b59c88b8595056bbae2bd31"></a>

## Next pages — more_option.response_cookies_to_add.ignore_domain / abf5bc01f462 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d3907f9c1c4886317aebdbad01323992b38158078d399c2c4ab40b3ad6e4ea39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1179832aeea10250c2f9cff73b51bc8adf096bdd6414dc980f7645a2c26811ef"></a>

## more_option.response_cookies_to_add.ignore_expiry — more_option.response_cookies_to_add.ignore_expiry / b46cbbc4bee8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-3749b9142ca9befc2ac4a5ad186cb380d8a4d0501d8615633d280614a4712b1a"></a>

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

<a id="canonical-514cb5b01f8ca2169e1c79c9fc4457f572bd93c1b6fcfc3e980ba2ee571299b1"></a>

## Direct properties — more_option.response_cookies_to_add.ignore_expiry / b46cbbc4bee8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-784eded6da3b19615cf6c1791189393b392b392051ff3c36b955defc9abb3b54"></a>

## Next pages — more_option.response_cookies_to_add.ignore_expiry / b46cbbc4bee8 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f86e59d1560487af3a991e1d0c765fda38ac1adab6b1633535cc7e7a94a4bbc3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8dac4c10935f955a1bd7e60323eb3b3e40f4dac72cbe8113d9a748e16454134"></a>

## more_option.response_cookies_to_add.ignore_httponly — more_option.response_cookies_to_add.ignore_httponly / 7615d3da5527 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-787b66727291fe4d68ff507edf1085ed1aa8b5a0e0703aef77ad12c98d1cc933"></a>

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

<a id="canonical-55ce6a63fd4be06519c560a2b974ae0eaeb71e02cbc012d0c14ea09268afe7b5"></a>

## Direct properties — more_option.response_cookies_to_add.ignore_httponly / 7615d3da5527 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f894d3505ec22ff1fe52175cfae463f72d3822a7f775502019f03f6cadc96d59"></a>

## Next pages — more_option.response_cookies_to_add.ignore_httponly / 7615d3da5527 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-6c383273aa4f7124daf6aab38333915bc8c9d20927456b112d8c6e50c5a8e6d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4806bc410e20303be6e0ef421ebd52dc82d66b43e48661d33111516b04077a83"></a>

## more_option.response_cookies_to_add.ignore_max_age — more_option.response_cookies_to_add.ignore_max_age / b072a11e9d00 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-542c36f72d516293574fe016d908cf6e73a873870180a7f74bdb433a2b2820fb"></a>

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

<a id="canonical-8c325a7c2dfb33092e70479d85640b81e6bb5b389e1a8fcdc081884be61de618"></a>

## Direct properties — more_option.response_cookies_to_add.ignore_max_age / b072a11e9d00 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-363d1d5595d6e697fd687510318966d9cd527c077bfe77da4424bd261af5d6ac"></a>

## Next pages — more_option.response_cookies_to_add.ignore_max_age / b072a11e9d00 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-085e2537e7060f8224a0d45a5527ebd5600696a6a51c9b0b90d96762940e713d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5013c216959e4e6b3b72bfaab1f03621bed96ce63b3ee8a4020855e1c198f20a"></a>

## more_option.response_cookies_to_add.ignore_partitioned — more_option.response_cookies_to_add.ignore_partitioned / 8dbffbae14c9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-4291850719ae8aeb29424db48792fda74ae5f44af6388945ffb6a108f73bbb7f"></a>

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

<a id="canonical-e9f1ea895dab4668991b07b9e27390a69d3be6e54f4267535fad54bce5ebb43c"></a>

## Direct properties — more_option.response_cookies_to_add.ignore_partitioned / 8dbffbae14c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c12c09def34cc34f0d8f1954850cc3d30c0a818e389efb1b3569a087807390ac"></a>

## Next pages — more_option.response_cookies_to_add.ignore_partitioned / 8dbffbae14c9 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a341f20ff2f14c6c8f2cd3ffe2f4555fdd9e6e280bf1431ff67788fc10241511"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ec0d387946885fafae3c3cdac8bf9cf80377d99cba4cd0a53256890c0308c62"></a>

## more_option.response_cookies_to_add.ignore_path — more_option.response_cookies_to_add.ignore_path / 721bd5b39de6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.ignore_path

<a id="canonical-25c226a795602dc1185d2c10be15945a887ed004b5efa2ae95a940382acdf7ec"></a>

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

<a id="canonical-914048e0beffdbfdba7b72845b1831ecc44063b61504252f87492c3fbe47f204"></a>

## Direct properties — more_option.response_cookies_to_add.ignore_path / 721bd5b39de6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d99f3ce4ec4b4729e4cc8d66af8b144f9cbe532ee49c8f6887fa44bae35282e9"></a>

## Next pages — more_option.response_cookies_to_add.ignore_path / 721bd5b39de6 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d28b394397609b939cb8a665a6357ef9da8406bc1614c401d03556ce4e90ce44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d44f56af8312ad09416c72144aaebd2c70863024e8a555292bc509164e2d884f"></a>

## more_option.response_cookies_to_add.ignore_samesite — more_option.response_cookies_to_add.ignore_samesite / 401ef1be14f2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-cc5cf26dbdc65c85a27c91e1d66c1784f65af6ffd253e29de44aa3d582666b48"></a>

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

<a id="canonical-1d626c4f0991ac5aa87b95c7ab804d105abc072a7c2f3c40db17f5fa428b870b"></a>

## Direct properties — more_option.response_cookies_to_add.ignore_samesite / 401ef1be14f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1cc2c0ee9801b96934e5519121f402afa5f1db7e20fd2710817d910bf0db454f"></a>

## Next pages — more_option.response_cookies_to_add.ignore_samesite / 401ef1be14f2 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0223382fd1c9543be0cb07550749b7e0c44a42349b38bd45fbc353ddfc189b6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ead3aa4a5cd850084b38f650f34c97a50e5a4a796bb46d8e2f46a67da52543d7"></a>

## more_option.response_cookies_to_add.ignore_secure — more_option.response_cookies_to_add.ignore_secure / ae65329f53c3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.ignore_secure

<a id="canonical-39ef404bc53ffebc9b7691f802f88fc0bd490ffa988dfb8e628d17dabeec2337"></a>

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

<a id="canonical-b946a36f3a32ffdedb838cd04efe8516dc4d5049c5686e9db60899bc8b6619b2"></a>

## Direct properties — more_option.response_cookies_to_add.ignore_secure / ae65329f53c3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7667bce0a2edb717ae6a36770e37ef7243c0cc75cf6f4ca481f033d06bbb2627"></a>

## Next pages — more_option.response_cookies_to_add.ignore_secure / ae65329f53c3 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d054e4d7acd1269a374c44672b4f110a125083e2ab70e3cef023577a61136894"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9aaedc0f2d701312861b4f2c225204c0cc79eb45313799520446dc8c2d8309a0"></a>

## more_option.response_cookies_to_add.ignore_value — more_option.response_cookies_to_add.ignore_value / c43fac182ea9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.ignore_value

<a id="canonical-f3d3ec0bbc65f621b3ccea2c0fff19893182475335482fb4dcd9a24c344b678d"></a>

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

<a id="canonical-5067fb500a48944f270f198cbdf2351462f0a50a6c425970df48bc5cd4c75a4e"></a>

## Direct properties — more_option.response_cookies_to_add.ignore_value / c43fac182ea9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4719f048223af1a17f9431ecda9685b9e0c1f0b4c1d4a15602fee296a14cfe2a"></a>

## Next pages — more_option.response_cookies_to_add.ignore_value / c43fac182ea9 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bccff969a23d7fa8583e2f9938f91d86ba4beaca7a1bd9e401b3994e55e2af07"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1ac3518436eba0715aaf2f6cdc73beed052f0383847e211b3faf7d69947fbbc3"></a>

## more_option.response_cookies_to_add.samesite_lax — more_option.response_cookies_to_add.samesite_lax / 1ae81e5bde49 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.samesite_lax

<a id="canonical-a9a22b766afbf9d4008defac7b0c4cbb8154b6a1338e8cce32f413c863772d8c"></a>

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

<a id="canonical-92e96fb17946c7f96866738573c2cd9fe8fe40f5e49e1675265464524ce07b8e"></a>

## Direct properties — more_option.response_cookies_to_add.samesite_lax / 1ae81e5bde49 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f60f9906a649baf429f7e10f9cc6beaf8a8adc4cc05cf20ab1839c66483bd814"></a>

## Next pages — more_option.response_cookies_to_add.samesite_lax / 1ae81e5bde49 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7e6021da5fb077a8d4774554e9ff72d8672677b0881226e16e72a183675c51a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef44abf611c959cf8627b94e2912a7960404298fb7eaa3ed4f67e0f63d2bf9fd"></a>

## more_option.response_cookies_to_add.samesite_none — more_option.response_cookies_to_add.samesite_none / 8a6ce19d7ddd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.samesite_none

<a id="canonical-0cad73583b692b1b9accb2d893c943931e856dd4d4179d4b55d849c39264ee58"></a>

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

<a id="canonical-7c63c4f9955195ac92c0c01af89c42bca79e160ab4c45c1c7b5dac55f22bbcef"></a>

## Direct properties — more_option.response_cookies_to_add.samesite_none / 8a6ce19d7ddd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-441c0e3cbfb56fb7ec2b765c2ec9acc403cb41a58b01e2d4b0a249c34417a222"></a>

## Next pages — more_option.response_cookies_to_add.samesite_none / 8a6ce19d7ddd / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e0d5bcd9ebcbf71287969e74fd350d224515e905750b8766b85475933573e899"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74e042e52386b59fcee01e6f337f763ca34d965fcec6af7431b95fddb2bcb58e"></a>

## more_option.response_cookies_to_add.samesite_strict — more_option.response_cookies_to_add.samesite_strict / edf244d91998 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.samesite_strict

<a id="canonical-cd35db1a1f3926e29fa3a9d1d877a4560a72b2608c05142ad59df6639a43cd2a"></a>

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

<a id="canonical-9128bf659c10bacfdb2da3a6bf5e9aa3eb2ae6fecec0e378560f51b731b86833"></a>

## Direct properties — more_option.response_cookies_to_add.samesite_strict / edf244d91998 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-401069cd1ba58b65db9e9328976f2b002ed5eb5868315517cc3d8f7e3ec5e66b"></a>

## Next pages — more_option.response_cookies_to_add.samesite_strict / edf244d91998 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7a4c8976ce3a00e6109190166b591890a8ef0156a098d16e71e109562acdefb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-39a508160097d2058c882849843756d9eddbc8b5d5e54e93b59af050b914c406"></a>

## more_option.response_cookies_to_add.secret_value — more_option.response_cookies_to_add.secret_value / 99b170318c2b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- more_option.response_cookies_to_add.secret_value

<a id="canonical-9891efd0e1cb55d9199b47a7aae2f972e95c784fe2b86dcd5998f5c623838399"></a>

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

<a id="canonical-1ec022324d60e7b93f262158ee1be44caa83ebaca00b8dbb05756698f398cf6e"></a>

## Direct properties — more_option.response_cookies_to_add.secret_value / 99b170318c2b / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-c9b4f14b459ab30b8c8170ddfa705981557eabecd67312a916be3247c02835bf): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-5b860351af3b932e7074596f5fb2f3112e37bf7a08a7b954cbb5f7309685b7c8): complete subsection reference.

<a id="canonical-f97cd2c288a08063bdad9bab0c1900920cfbe74a920617707a2089f7ba66980c"></a>

## Next pages — more_option.response_cookies_to_add.secret_value / 99b170318c2b / 4

- [more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-c9b4f14b459ab30b8c8170ddfa705981557eabecd67312a916be3247c02835bf)
- [more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-5b860351af3b932e7074596f5fb2f3112e37bf7a08a7b954cbb5f7309685b7c8)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c9b4f14b459ab30b8c8170ddfa705981557eabecd67312a916be3247c02835bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dac3285164ea539b52d0a7a683602fad9fd8623924724e05842aa2dffde7e637"></a>

## more_option.response_cookies_to_add.secret_value.blindfold_secret_info — more_option.response_cookies_to_add.secret_value.blindfold_secret_info / d1d826cd0f5c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-7a4c8976ce3a00e6109190166b591890a8ef0156a098d16e71e109562acdefb1)
- more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-d2bd020e35e50af720fd15765772ba7e6d7af4dff0836c9be2d3b6c0bf1ee821"></a>

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

<a id="canonical-67c4bef2b74445ba76336d005c946683aa6d0e85aecba2610de90830f158d722"></a>

## Direct properties — more_option.response_cookies_to_add.secret_value.blindfold_secret_info / d1d826cd0f5c / 3

<a id="canonical-8610531b4001a006c15222bff92973944f5d00045f1c9a1b6c07d5ab6eea23c1"></a>

<a id="canonical-cba78d3e43d3d94d0d780def982c1a9dd78149d2433cc780ed12f1b81cb484b8"></a>

## decryption_provider property — more_option.response_cookies_to_add.secret_value.blindfold_secret_info / d1d826cd0f5c / 4

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

<a id="canonical-218efba818eee5c0290ba21b075b3fadf027efb694bf34621e6ec38cac999165"></a>

<a id="canonical-b87e29dec0eab014086cb7ffa8c44dab2845133991204bf2f8f9c324a8bc7e0e"></a>

## location property — more_option.response_cookies_to_add.secret_value.blindfold_secret_info / d1d826cd0f5c / 5

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

<a id="canonical-ae991b476df06683f1f1a4285264bbca90741b3045cb8ff60ddae812754d1a79"></a>

<a id="canonical-48e659c4f675d0dae994529e51fb64bc5b363db0a9aadfd1a4c67e06b49586a5"></a>

## store_provider property — more_option.response_cookies_to_add.secret_value.blindfold_secret_info / d1d826cd0f5c / 6

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

<a id="canonical-69b2d26c47bc1ae1e1731ebbae1994148263b60a8420cde755480926241540f8"></a>

## Next pages — more_option.response_cookies_to_add.secret_value.blindfold_secret_info / d1d826cd0f5c / 7

- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-7a4c8976ce3a00e6109190166b591890a8ef0156a098d16e71e109562acdefb1)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5b860351af3b932e7074596f5fb2f3112e37bf7a08a7b954cbb5f7309685b7c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-77b2b64a1ad2f3389be96e783bd5c35f0e25bcef84f798f01b1ae14e14163dec"></a>

## more_option.response_cookies_to_add.secret_value.clear_secret_info — more_option.response_cookies_to_add.secret_value.clear_secret_info / ef45eeec612b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-ec6b3c8556cf6ffa11d51687b454365212ab3801fe8820a2ce82413a68c0e4fa)
- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-7a4c8976ce3a00e6109190166b591890a8ef0156a098d16e71e109562acdefb1)
- more_option.response_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-a081ad0814d2eb896afce6579e5f48d9edc7a6b38982f5221e96b47cc96f3686"></a>

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

<a id="canonical-3b46c9239178ff56cb1a0d2d8ace7e0a0b859e6dfa7304e98457d896a170809d"></a>

## Direct properties — more_option.response_cookies_to_add.secret_value.clear_secret_info / ef45eeec612b / 3

<a id="canonical-1344734baba41a7fb8d3eacb344d79f90178b24e06032468549c5a9762ac7a75"></a>

<a id="canonical-df379551b3872d7b0be7da68e1dc23e28adffa82e8854ce805e2ca844faf346f"></a>

## provider_ref property — more_option.response_cookies_to_add.secret_value.clear_secret_info / ef45eeec612b / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-c9d57ef932d34cf4d8febb57f47ec2c91043e6e69a85ec19726f5a65f232be1a"></a>

<a id="canonical-2af15beba004d038667b28cb5dfc2452b033415a0b601f04bf07cc0decc76f72"></a>

## url property — more_option.response_cookies_to_add.secret_value.clear_secret_info / ef45eeec612b / 5

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

<a id="canonical-26fe94f40a05071c9f37616ffca85a0959e1cb517c3dd556cd1716d12dc7099f"></a>

## Next pages — more_option.response_cookies_to_add.secret_value.clear_secret_info / ef45eeec612b / 6

- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-7a4c8976ce3a00e6109190166b591890a8ef0156a098d16e71e109562acdefb1)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7bc015dbe38d29bfa4524071a40731d23f7944255a5d723f15cebdee85b70926"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2aa61668e851c279373099d34f66840525a9cf4dd3df7530d85f65548d4afc2"></a>

## more_option.response_headers_to_add — more_option.response_headers_to_add / 1a84ca76f27e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- more_option.response_headers_to_add

<a id="canonical-714e083900179b61f5ca437c8b718bb7fc2ab25ad861b6bbcd0413767c0be943"></a>

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

<a id="canonical-de86fa4fabcfa2ecdfb26416761f5b40292b274fe1335c851aade4419eabc6df"></a>

## Direct properties — more_option.response_headers_to_add / 1a84ca76f27e / 3

<a id="canonical-23d0adffa3282c2b3bc1a0dbb155fd5aa04faf2cfd79fedb2903b447508a010b"></a>

<a id="canonical-81193634ec2cd72bd17792ab01ac58d4874238e850fdc4cf11608218194feca0"></a>

## append property — more_option.response_headers_to_add / 1a84ca76f27e / 4

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

<a id="canonical-c75e627494f1af4f840ef7078f65b6c633dbd312c9196985b4b544ad74654839"></a>

<a id="canonical-bca8d0fb35211b5a46ccb7b8974c6f254e9d41e8f85f6e3ba2ceda93176ae1e6"></a>

## name property — more_option.response_headers_to_add / 1a84ca76f27e / 5

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

- [secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-f5d90ce95c9884e24647b0f5cda856c1888bb12a05ed905c85115e6e1dc2e72d): complete subsection reference.

<a id="canonical-19bc71844ae85e08e6dd679c9590833246af8a12fa1cb88fb44dcf5154b6fe6d"></a>

<a id="canonical-d102057bbaf1562dc50de747a3f98f8f9c83d7f7c6559bd4ed51ee410f5cdea4"></a>

## value property — more_option.response_headers_to_add / 1a84ca76f27e / 6

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

<a id="canonical-f5204a1c49018a7c339eeb35e1337210d0e6ed9821cec80aef6f4d8e722463d9"></a>

## Next pages — more_option.response_headers_to_add / 1a84ca76f27e / 7

- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-f5d90ce95c9884e24647b0f5cda856c1888bb12a05ed905c85115e6e1dc2e72d)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f5d90ce95c9884e24647b0f5cda856c1888bb12a05ed905c85115e6e1dc2e72d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-776034adea4f0f9151434c326237e1d5bafe9eaf34f2d3c93133ae9ed7c6728a"></a>

## more_option.response_headers_to_add.secret_value — more_option.response_headers_to_add.secret_value / 20189b4bccea / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-7bc015dbe38d29bfa4524071a40731d23f7944255a5d723f15cebdee85b70926)
- more_option.response_headers_to_add.secret_value

<a id="canonical-5dba2669edeeab7293e923bb1fff6de0bb7a630ebef1ccd5282422f590c7eeda"></a>

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

<a id="canonical-1817c199dbf66d9fc0fef5df8c8b614ea872a3e7c89bb7f0ddf615f4464fe3eb"></a>

## Direct properties — more_option.response_headers_to_add.secret_value / 20189b4bccea / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-cc60eb4090b4e63678b27fafa6408799e90368924a617f2d7caff0eeb9451978): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-c986525580e9f4237a37eae13998b419ee94dc58d0a84aa5a4f9238439f77e7a): complete subsection reference.

<a id="canonical-ad6b5710e1d727435836fcbe7ecd1365d26aa31b46f322ae0429e4fdb3a93907"></a>

## Next pages — more_option.response_headers_to_add.secret_value / 20189b4bccea / 4

- [more_option.response_headers_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-cc60eb4090b4e63678b27fafa6408799e90368924a617f2d7caff0eeb9451978)
- [more_option.response_headers_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-c986525580e9f4237a37eae13998b419ee94dc58d0a84aa5a4f9238439f77e7a)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-7bc015dbe38d29bfa4524071a40731d23f7944255a5d723f15cebdee85b70926)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cc60eb4090b4e63678b27fafa6408799e90368924a617f2d7caff0eeb9451978"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57a926030175dbc2ab94f16c45b39fea21cb0dc24a77fdc487add6b62f47fb22"></a>

## more_option.response_headers_to_add.secret_value.blindfold_secret_info — more_option.response_headers_to_add.secret_value.blindfold_secret_info / ef81c5452442 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-7bc015dbe38d29bfa4524071a40731d23f7944255a5d723f15cebdee85b70926)
- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-f5d90ce95c9884e24647b0f5cda856c1888bb12a05ed905c85115e6e1dc2e72d)
- more_option.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-a2c1bb6f2e1d1341db9990050874a559a3cb01e61a2f7796b765b56b79b3a747"></a>

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

<a id="canonical-acae07d79bd334e53f161713b1ac9f80fd43b6fee0c35336089a7948179bc5d4"></a>

## Direct properties — more_option.response_headers_to_add.secret_value.blindfold_secret_info / ef81c5452442 / 3

<a id="canonical-34b42eb343a7f360177b7570163324095f013d11556b917f26085353e7b9cda3"></a>

<a id="canonical-5d5778221ec4e3eda8ba9c863970d762514b0fcf51d288135f0b4fa5eb726aa0"></a>

## decryption_provider property — more_option.response_headers_to_add.secret_value.blindfold_secret_info / ef81c5452442 / 4

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

<a id="canonical-a7d41695bde73322a1f004f919d46e3c9da1b7855605bb8b273cda7cd39d1254"></a>

<a id="canonical-8e31a414727971d2731d46bdf7b9e79a32ccc543c47651a262494dc1d4f79c4f"></a>

## location property — more_option.response_headers_to_add.secret_value.blindfold_secret_info / ef81c5452442 / 5

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

<a id="canonical-ff4c0b9875b715c89f8c1626e516ebf1d7fd4ceeb783b2ae8283666638445f25"></a>

<a id="canonical-d25d2349602224d13d483573be9a1ec6c5eb6335b26bd9e892840a9d0f079efb"></a>

## store_provider property — more_option.response_headers_to_add.secret_value.blindfold_secret_info / ef81c5452442 / 6

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

<a id="canonical-8aea78d99ce7174377278425d7ed3135799f6aa2587255342a6c1c90d90770cd"></a>

## Next pages — more_option.response_headers_to_add.secret_value.blindfold_secret_info / ef81c5452442 / 7

- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-f5d90ce95c9884e24647b0f5cda856c1888bb12a05ed905c85115e6e1dc2e72d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c986525580e9f4237a37eae13998b419ee94dc58d0a84aa5a4f9238439f77e7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95d0ca57c04f02cfe09a2330eefab93c82687b17e820399a89bb1012673e622b"></a>

## more_option.response_headers_to_add.secret_value.clear_secret_info — more_option.response_headers_to_add.secret_value.clear_secret_info / 77d98023819a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-13bbd000e2c68826bf1c0be37491cfc06816379212a16f2fa40ba4bf550b2cee)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-7bc015dbe38d29bfa4524071a40731d23f7944255a5d723f15cebdee85b70926)
- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-f5d90ce95c9884e24647b0f5cda856c1888bb12a05ed905c85115e6e1dc2e72d)
- more_option.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-1a67ee57fe270c9b252a3be32c8b490d6d7671436535f0225d1d03290aef9d50"></a>

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

<a id="canonical-c685068d2231e3057558c1d64e23bb8dc5fc6792050f24e92da49dfce9968c85"></a>

## Direct properties — more_option.response_headers_to_add.secret_value.clear_secret_info / 77d98023819a / 3

<a id="canonical-571366b59a31ca2efde10810f63d5e912297a8065d0cf608c46845406823ed93"></a>

<a id="canonical-7b0017d8d21ceb20f5835e61785043b82dca92029cd2db1f337c1d665f4f7c1f"></a>

## provider_ref property — more_option.response_headers_to_add.secret_value.clear_secret_info / 77d98023819a / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-e61d70e70a2221bca2da1f77822194627e3db4834468381729c6bc67a12d6bd3"></a>

<a id="canonical-68597e47dc1a15e7ed1d9195207d278fe762e4b014ea93e05208bddc02391260"></a>

## url property — more_option.response_headers_to_add.secret_value.clear_secret_info / 77d98023819a / 5

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

<a id="canonical-58773cbb6170819b846381902615bf65f6f2303a36956ff3ee432a6bfbd1dd19"></a>

## Next pages — more_option.response_headers_to_add.secret_value.clear_secret_info / 77d98023819a / 6

- [more_option.response_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-f5d90ce95c9884e24647b0f5cda856c1888bb12a05ed905c85115e6e1dc2e72d)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-0ae28b3b2b8df4f8ce2c8e04c0176280fe9c9128fb14f292dee21e7534b1c785"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5632813dd72c6c4645439b831a93d325b47149fca0612e7cc014fdd32367660f"></a>

## multi_lb_app — multi_lb_app / 3e9ffc5e4d4a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- multi_lb_app

<a id="canonical-5711fccd4ce52343cceaffc81e53a55cf2302e8bed1544f3eaf3ff35482cdf39"></a>

Type: `["object", {}]`. Optional.

\[OneOf: multi\_lb\_app, single\_lb\_app\] Configuration parameter for multi lb app.

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

- [multi_lb_app](resources--http_loadbalancer--reference--group-021.md#canonical-5711fccd4ce52343cceaffc81e53a55cf2302e8bed1544f3eaf3ff35482cdf39)
- [single_lb_app](resources--http_loadbalancer--reference--group-026.md#canonical-61edcd8f7e9848ea5675c000f776f6de16862c029a213391c63a3a015bbcac1f)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
multi_lb_app = {}
```

<a id="canonical-accf68a0834a83aaf84d69223acbd9628633ba9c0e1cc1be238c523459a56803"></a>

## Direct properties — multi_lb_app / 3e9ffc5e4d4a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8d6de20fa47711dcc472594dcd58a7979ee7f9ab1db8009db4c7ea650a29c129"></a>

## Next pages — multi_lb_app / 3e9ffc5e4d4a / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-91668e9b3365ea876e1816b4b40472e96059f8303e3a32daf41ba574857d4efe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50bfa84fa468f736d129ed6d45bf036be980b8388056a7a08ece929b73b9efce"></a>

## no_challenge — no_challenge / 318f991b7f2f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- no_challenge

<a id="canonical-de253bc4120f307cd44e3e5625cf404100e08678cf40944fa82eb5b7f4a67ccb"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no challenge. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_challenge = {}
```

<a id="canonical-e0b78c3385a572d9e02d02d4eebbbbc4c064054742bb4f0cf1086a9245f56703"></a>

## Direct properties — no_challenge / 318f991b7f2f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc8cc40ee431f4f586499e6bc6bf32fcd8476590abdd0e1ecd8a229cae4cd47e"></a>

## Next pages — no_challenge / 318f991b7f2f / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-67a5b67e1c9362160c1235a278e79e0e6a99c2fc8595e33cad1c274c07dc323d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-433bf728413b68b5210a0af4fa17b0984bb68376814855c69670330c8266253f"></a>

## no_service_policies — no_service_policies / a13397f44e5d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- no_service_policies

<a id="canonical-512a95be8c72528d5f4d1fe7510ec1f5d3cabd0ce77b61a10085a3176b9f4dc2"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_service_policies = {}
```

<a id="canonical-72c2397e188ff28fa07da69b67acc0cd05f17c08ac5504dfb2077d453d685acf"></a>

## Direct properties — no_service_policies / a13397f44e5d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-529010b04a49ee9d1253081e229b49a8f0f45c3b06c076a64e6eb51aef02baf8"></a>

## Next pages — no_service_policies / a13397f44e5d / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-787b9a4b463c3f04423a7e1d75fde292ce4d6eb81b52a845fe872fd562409679"></a>

## origin_server_subset_rule_list — origin_server_subset_rule_list / fb9b0a6bbdf0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- origin_server_subset_rule_list

<a id="canonical-c53728f678c2ca9fa4351fc1c663f13849c514a3ee9f892a2fb36ae5c883cf77"></a>

Type: `"object"`. single nested block, Optional.

Origin Server Subset Rule List Type. List of Origin Pools.

Upstream description:

List of Origin Pools.

Receipt-pinned upstream constraints:

```json
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
origin_server_subset_rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330c8432bdc5c1c96ef85166f67e7c2315d7088e9dbe9068a8c9be217d09f52"></a>

## Direct properties — origin_server_subset_rule_list / fb9b0a6bbdf0 / 3

- [origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264): complete subsection reference.

<a id="canonical-c8abc9316c463c63d715a1d5e5beaf10f045a77d373cceed45f2f3a00498d164"></a>

## Next pages — origin_server_subset_rule_list / fb9b0a6bbdf0 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26c8e1417bd3ad7d6a7d726ad2dc6ad781a16268019d9613fe258c0a3f11fa3d"></a>

## origin_server_subset_rule_list.origin_server_subset_rules — origin_server_subset_rule_list.origin_server_subset_rules / 732eb80f5450 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- origin_server_subset_rule_list.origin_server_subset_rules

<a id="canonical-67c37e8c4354f13fdd948d8ce9308286ef249bb2ac13533b964dd7f84453ad67"></a>

Type: `"object"`. list nested block, Optional.

Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN,
Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin
Server Subset is a sequential engine where rules are evaluated one after the other. It's important
to..

Upstream description:

Origin Server Subset Rules allow users to define match condition on Client (IP address, ASN,
Country), IP Reputation, Regional Edge names, Request for subset selection of origin servers. Origin
Server Subset is a sequential engine where rules are evaluated one after the other. It's important
to define the correct order for Origin Server Subset to GET the intended result, rules are evaluated
from top to bottom in the list. When an Origin server subset rule is matched, then this selection
rule takes effect and no more rules are evaluated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("origin_server_subsets_action"),
  validators.ConflictingListObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingListObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingListObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingListObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingListObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingListObjectAttributes("client_selector",
    "none"),
  validators.ConflictingListObjectAttributes("ip_matcher",
    "ip_prefix_list")}
```

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

Terraform syntax:

```terraform
origin_server_subset_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-c83d70edd1b6136c84d97ba9b106b712ff518ae3bcf3b3845cbc276d8b6a86df"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules / 732eb80f5450 / 3

- [any_asn](resources--http_loadbalancer--reference--group-021.md#canonical-72568c8cc1f0902dd65abc87a8d4ebad8847620fb4be535e1dae3d5d6a07d479): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-021.md#canonical-a46c9285e8e48eb76e80a398024201dd62f66a32e4d8d5ed38738a92185bc75e): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-021.md#canonical-eaecbe5fc64a352deefcd5e56317a61c87c94224acbdad91356502c77d542653): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-021.md#canonical-9db9da0aedd8e075416f9269441afbd2ce239e47833594508724d79d8e55aa67): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-021.md#canonical-e679fa829d9ef7ff19eee91f0b4bcb6ad4a2930f79472438a608dd2e896d6615): complete subsection reference.

<a id="canonical-a516fd737ef2959ac97216b81274244cbfe9030ff865057464969c0c884f59f8"></a>

<a id="canonical-9458f8ca0e93f4aa6870c6a33a261fd194693448444a0b80494494e66bc61011"></a>

## country_codes property — origin_server_subset_rule_list.origin_server_subset_rules / 732eb80f5450 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Country Codes List. List of Country Codes. Possible values are \`COUNTRY\_NONE\`, \`COUNTRY\_AD\`,
\`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`, \`COUNTRY\_AI\`, \`COUNTRY\_AL\`,
\`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`, \`COUNTRY\_AQ\`, \`COUNTRY\_AR\`,
\`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`, \`COUNTRY\_AW\`, \`COUNTRY\_AX\`,
\`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`, \`COUNTRY\_BD\`, \`COUNTRY\_BE\`,
\`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`, \`COUNTRY\_BI\`, \`COUNTRY\_BJ\`,
\`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`, \`COUNTRY\_BO\`, \`COUNTRY\_BQ\`,
\`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`, \`COUNTRY\_BV\`, \`COUNTRY\_BW\`,
\`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`, \`COUNTRY\_CC\`, \`COUNTRY\_CD\`,
\`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`, \`COUNTRY\_CI\`, \`COUNTRY\_CK\`,
\`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`, \`COUNTRY\_CO\`, \`COUNTRY\_CR\`,
\`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`, \`COUNTRY\_CW\`, \`COUNTRY\_CX\`,
\`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`, \`COUNTRY\_DJ\`, \`COUNTRY\_DK\`,
\`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`, \`COUNTRY\_EC\`, \`COUNTRY\_EE\`,
\`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`, \`COUNTRY\_ES\`, \`COUNTRY\_ET\`,
\`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`, \`COUNTRY\_FM\`, \`COUNTRY\_FO\`,
\`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`, \`COUNTRY\_GD\`, \`COUNTRY\_GE\`,
\`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`, \`COUNTRY\_GI\`, \`COUNTRY\_GL\`,
\`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`, \`COUNTRY\_GQ\`, \`COUNTRY\_GR\`,
\`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`, \`COUNTRY\_GW\`, \`COUNTRY\_GY\`,
\`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`, \`COUNTRY\_HR\`, \`COUNTRY\_HT\`,
\`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`, \`COUNTRY\_IL\`, \`COUNTRY\_IM\`,
\`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`, \`COUNTRY\_IR\`, \`COUNTRY\_IS\`,
\`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`, \`COUNTRY\_JO\`, \`COUNTRY\_JP\`,
\`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`, \`COUNTRY\_KI\`, \`COUNTRY\_KM\`,
\`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`, \`COUNTRY\_KW\`, \`COUNTRY\_KY\`,
\`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`, \`COUNTRY\_LC\`, \`COUNTRY\_LI\`,
\`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`, \`COUNTRY\_LT\`, \`COUNTRY\_LU\`,
\`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`, \`COUNTRY\_MC\`, \`COUNTRY\_MD\`,
\`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`, \`COUNTRY\_MH\`, \`COUNTRY\_MK\`,
\`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`, \`COUNTRY\_MO\`, \`COUNTRY\_MP\`,
\`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`, \`COUNTRY\_MT\`, \`COUNTRY\_MU\`,
\`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`, \`COUNTRY\_MY\`, \`COUNTRY\_MZ\`,
\`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`, \`COUNTRY\_NF\`, \`COUNTRY\_NG\`,
\`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`, \`COUNTRY\_NP\`, \`COUNTRY\_NR\`,
\`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`, \`COUNTRY\_PA\`, \`COUNTRY\_PE\`,
\`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`, \`COUNTRY\_PK\`, \`COUNTRY\_PL\`,
\`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`, \`COUNTRY\_PS\`, \`COUNTRY\_PT\`,
\`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`, \`COUNTRY\_RE\`, \`COUNTRY\_RO\`,
\`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`, \`COUNTRY\_SA\`, \`COUNTRY\_SB\`,
\`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`, \`COUNTRY\_SG\`, \`COUNTRY\_SH\`,
\`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`, \`COUNTRY\_SL\`, \`COUNTRY\_SM\`,
\`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`, \`COUNTRY\_SS\`, \`COUNTRY\_ST\`,
\`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`, \`COUNTRY\_SZ\`, \`COUNTRY\_TC\`,
\`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`, \`COUNTRY\_TH\`, \`COUNTRY\_TJ\`,
\`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`, \`COUNTRY\_TN\`, \`COUNTRY\_TO\`,
\`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`, \`COUNTRY\_TW\`, \`COUNTRY\_TZ\`,
\`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`, \`COUNTRY\_US\`, \`COUNTRY\_UY\`,
\`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`, \`COUNTRY\_VE\`, \`COUNTRY\_VG\`,
\`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`, \`COUNTRY\_WF\`, \`COUNTRY\_WS\`,
\`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`, \`COUNTRY\_YT\`, \`COUNTRY\_ZA\`,
\`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

List of Country Codes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

- [ip_matcher](resources--http_loadbalancer--reference--group-021.md#canonical-2366270da3a20b5d81ebbbe31f9b33db5be09f979629a5afbde5481ddd37fd1f): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-021.md#canonical-154b80b22444dfb0d70cb0a3da95834f09f977ec18cf944ddf7a66e7c54ebcb1): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-021.md#canonical-08a935119871c441b3473ce4a32f4cc511826867746ff4ecdf676cbeb05e569d): complete subsection reference.

- [none](resources--http_loadbalancer--reference--group-021.md#canonical-a5b3bfb0fc9514908b1ba0435689dbe7b92a540d10e6908d7fd00b57120aedeb): complete subsection reference.

<a id="canonical-7d116fe0383c5e8ca27ea5e1569623c10d34a5c367faef360955dc163eaa6b8d"></a>

<a id="canonical-ba80153f276190098cba1ef22ccb9c210c49d8c810de794b5bc0d7bc527dbca7"></a>

## origin_server_subsets_action property — origin_server_subset_rule_list.origin_server_subset_rules / 732eb80f5450 / 5

Type: `["map", "string"]`. Optional.

Add labels to select one or more origin servers.

Upstream description:

Add labels to select one or more origin servers. Note: The pre-requisite settings to be configured
in the origin pool are: &#8203;1. Add labels to origin servers &#8203;2. Enable subset load
balancing in the Origin Server Subsets section and configure keys in origin server subsets classes.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-f8a0fc7100faa515cb4231d75523d8d9dc22f430caf1c2d745d8490f80afd97d"></a>

<a id="canonical-3270777c55ef5358a4295a70e60e8afa5ad86e8ba4db4e26d2669a1b2dcfc028"></a>

## re_name_list property — origin_server_subset_rule_list.origin_server_subset_rules / 732eb80f5450 / 6

Type: `["list", "string"]`. Optional.

RE Names. List of RE names for match.

Upstream description:

List of RE names for match.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-cc593bd22ff1a69152ffb60662cfe7cdcacd482d3758e9b37b2c1049c2de4ec6"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules / 732eb80f5450 / 7

- [origin_server_subset_rule_list.origin_server_subset_rules.any_asn](resources--http_loadbalancer--reference--group-021.md#canonical-72568c8cc1f0902dd65abc87a8d4ebad8847620fb4be535e1dae3d5d6a07d479)
- [origin_server_subset_rule_list.origin_server_subset_rules.any_ip](resources--http_loadbalancer--reference--group-021.md#canonical-a46c9285e8e48eb76e80a398024201dd62f66a32e4d8d5ed38738a92185bc75e)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_list](resources--http_loadbalancer--reference--group-021.md#canonical-eaecbe5fc64a352deefcd5e56317a61c87c94224acbdad91356502c77d542653)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](resources--http_loadbalancer--reference--group-021.md#canonical-9db9da0aedd8e075416f9269441afbd2ce239e47833594508724d79d8e55aa67)
- [origin_server_subset_rule_list.origin_server_subset_rules.client_selector](resources--http_loadbalancer--reference--group-021.md#canonical-e679fa829d9ef7ff19eee91f0b4bcb6ad4a2930f79472438a608dd2e896d6615)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](resources--http_loadbalancer--reference--group-021.md#canonical-2366270da3a20b5d81ebbbe31f9b33db5be09f979629a5afbde5481ddd37fd1f)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list](resources--http_loadbalancer--reference--group-021.md#canonical-154b80b22444dfb0d70cb0a3da95834f09f977ec18cf944ddf7a66e7c54ebcb1)
- [origin_server_subset_rule_list.origin_server_subset_rules.metadata](resources--http_loadbalancer--reference--group-021.md#canonical-08a935119871c441b3473ce4a32f4cc511826867746ff4ecdf676cbeb05e569d)
- [origin_server_subset_rule_list.origin_server_subset_rules.none](resources--http_loadbalancer--reference--group-021.md#canonical-a5b3bfb0fc9514908b1ba0435689dbe7b92a540d10e6908d7fd00b57120aedeb)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-72568c8cc1f0902dd65abc87a8d4ebad8847620fb4be535e1dae3d5d6a07d479"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee427ce1d96096fc6a7d081312fea4c59776cabb0072c99f2d2140af591ebe72"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.any_asn — origin_server_subset_rule_list.origin_server_subset_rules.any_asn / cc2d97fd4da8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- origin_server_subset_rule_list.origin_server_subset_rules.any_asn

<a id="canonical-e5d3d3c24ab4cfaccb5516240c37a13dc41e6405e511b3c71f2343440787b116"></a>

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
any_asn = {}
```

<a id="canonical-7737f31bf980341d5ea58a5014eb4e9793317e0575337310162ddf6eb08a653b"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.any_asn / cc2d97fd4da8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0efc8b5dde38481774e1bac72b85b8aed836b386c05212544ca09b33427eff26"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.any_asn / cc2d97fd4da8 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a46c9285e8e48eb76e80a398024201dd62f66a32e4d8d5ed38738a92185bc75e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34f5e266ed4a19647f1c4d08f26ce08a0d03e2c7149e358c6f1273de3774528e"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.any_ip — origin_server_subset_rule_list.origin_server_subset_rules.any_ip / 9f40cb24fdc1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- origin_server_subset_rule_list.origin_server_subset_rules.any_ip

<a id="canonical-d795eff266235ec4019d9bd08d8a1599e6f7b13915a9d34f3c43ea224123af7f"></a>

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
any_ip = {}
```

<a id="canonical-f1f7b29a0de9f7d3026f8243e74983981dc554ce64a31cde229adc306f255573"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.any_ip / 9f40cb24fdc1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06cd22a5306cc30df63c6f2fa812df6eb10fd70bff1126b7ce2975a0a1b923ff"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.any_ip / 9f40cb24fdc1 / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-eaecbe5fc64a352deefcd5e56317a61c87c94224acbdad91356502c77d542653"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-626db255b1f873812a472093a3cc22787ad0cf432c18208c2e13ecd9e9ce452b"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.asn_list — origin_server_subset_rule_list.origin_server_subset_rules.asn_list / 2f6fb0c2a361 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_list

<a id="canonical-d51789ff8fb18936b527fe23eba2ae30218bb5e6de16d62b837c9cdbcc5533e5"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-8f005ff7dbe5517847b737ff8dc6ea0659288519714d36153736e29da6f2a22a"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.asn_list / 2f6fb0c2a361 / 3

<a id="canonical-bea461dbeed8c4e145f431d3714c4e38192495b3ebe07e843aed51a6ab7126ea"></a>

<a id="canonical-a17373e4c7092bf8ffd1da499ee86c8d179688bdeeb5de3ae48b9213df4d9313"></a>

## as_numbers property — origin_server_subset_rule_list.origin_server_subset_rules.asn_list / 2f6fb0c2a361 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-35a361fd88cd5e4b4de078ac8f4e834e99e63b933f5282cec44c8522c2ae8edf"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.asn_list / 2f6fb0c2a361 / 5

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9db9da0aedd8e075416f9269441afbd2ce239e47833594508724d79d8e55aa67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c6b24a727e700c746a74190b908f25ff40d6ed1be7ebdbccfee453a4a2a5a7f"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher / 37b37a7f67ca / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher

<a id="canonical-5a1366fb28ecd8888e245e181efc6bb9f6ac06b99611093d0c68091eeeb06aae"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-5983728dd146ffdf1c740f5d03c4dc497d41e9ee80b7b047482678edf8d7871d"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher / 37b37a7f67ca / 3

- [asn_sets](resources--http_loadbalancer--reference--group-021.md#canonical-9ff625726bf20f2d730e5f5432561c94e89eab454b18e17ab5df07dc45e45bcb): complete subsection reference.

<a id="canonical-c0c9f16fa6af20d8cc5192b2f9be00038c032a0b1215a0bf506a81ec3acfeb59"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher / 37b37a7f67ca / 4

- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-021.md#canonical-9ff625726bf20f2d730e5f5432561c94e89eab454b18e17ab5df07dc45e45bcb)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9ff625726bf20f2d730e5f5432561c94e89eab454b18e17ab5df07dc45e45bcb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cda22ccb8f1489fa98d16285e31f254c75d219d45c3ec7924a44d407ba86b5a9"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / bb4c3c854840 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](resources--http_loadbalancer--reference--group-021.md#canonical-9db9da0aedd8e075416f9269441afbd2ce239e47833594508724d79d8e55aa67)
- origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets

<a id="canonical-0df9016164a1ed2e1d27c501d693d0839482065fbf67021d050e43bfddcedbb4"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-e63e645766f43e5f83bc98cda896cc6a852203643850b1e3dda4b5679bbddecc"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / bb4c3c854840 / 3

<a id="canonical-0905bf079f4b41d5f9cb92c4eb18ff9b790792f785a55c9cf746a7d564c90d1e"></a>

<a id="canonical-60f08c4a4cfa82e3825fb64d9c557486d4366d7168ce2f0f4237dd3337c48df6"></a>

## kind property — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / bb4c3c854840 / 4

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

<a id="canonical-8f8647161ab75fa96ced43bc0393d09496774d771faaf0a0d7847efd7d0b0c19"></a>

<a id="canonical-8108d6e5c2591483fb739afc554042fa24c78ade3fdf9eb633212c462e4f6903"></a>

## name property — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / bb4c3c854840 / 5

Type: `"string"`. Optional.

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

<a id="canonical-d97cd8b00ff40f5c209e1107945b048b87e2bffe94193c18a27851765436fa06"></a>

<a id="canonical-9230f1bea27af474e9d67a00393a5b9b73a37b7a806513b0e112b635e154f8f9"></a>

## namespace property — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / bb4c3c854840 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-215d6d659014656a93c40a8b364f60a6be6b14c210382a988a1fd59365ff0e6e"></a>

<a id="canonical-95fa9de9809c7233064d83352bccd9acb25e4e3ebae1cbc0f375eef5c4f450a1"></a>

## tenant property — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / bb4c3c854840 / 7

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

<a id="canonical-193c24480856997d876ee1c30a0af86c262df508bebb3e723e1938a0cc749376"></a>

<a id="canonical-cc4cf6194dfe15af40480a7abb234e6839e916d12975b78d30168cbc33c3a6bd"></a>

## uid property — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / bb4c3c854840 / 8

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

<a id="canonical-12017e36b6946c183b27c6950aa90e3e0b7f0ba8e49a7a34ff99a0dda4645ccd"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher.asn_sets / bb4c3c854840 / 9

- [origin_server_subset_rule_list.origin_server_subset_rules.asn_matcher](resources--http_loadbalancer--reference--group-021.md#canonical-9db9da0aedd8e075416f9269441afbd2ce239e47833594508724d79d8e55aa67)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-e679fa829d9ef7ff19eee91f0b4bcb6ad4a2930f79472438a608dd2e896d6615"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df02193047c8ca8b212beb0588c8991f81c5715b44a147e8c2c3a6c07c435690"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.client_selector — origin_server_subset_rule_list.origin_server_subset_rules.client_selector / 7f493a7dc4a7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- origin_server_subset_rule_list.origin_server_subset_rules.client_selector

<a id="canonical-f6e835adfea5e98950a6b428ad1bb4633b4560c1551961e309c229156ab26029"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-3ddc67295340ef4fb8a99b6bcd149a160781ac7edb9d31c6f1e403fec18f4fae"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.client_selector / 7f493a7dc4a7 / 3

<a id="canonical-034d86476348367d1848d1d97b4d752a3bb5b704f4c982d046edb6fd240b29b6"></a>

<a id="canonical-25c2bd2aaa8ad8f3a944e130bbbe6ac3b34932cea88ea39fb23e629c2d03c552"></a>

## expressions property — origin_server_subset_rule_list.origin_server_subset_rules.client_selector / 7f493a7dc4a7 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-14f4f30bdc42f872a1c09134eb6dabfafd7c952da8287988d27049148d7db3db"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.client_selector / 7f493a7dc4a7 / 5

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-2366270da3a20b5d81ebbbe31f9b33db5be09f979629a5afbde5481ddd37fd1f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e646e258bf69e7ba464e6b0256aa383fa801cb77282c88c5d599238c24e6f99a"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher / db84bfc8526f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher

<a id="canonical-15cbcbda0da45f513d0f05cb834320d2a4722ac1ca723f2e02fb74ca2cdadf73"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-e2f44a08092f4060cab6ed6a0e88a9af10dc00d1db4ced9daa3e7af825247627"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher / db84bfc8526f / 3

<a id="canonical-7dd0d9f6dca4940362800dd9c96f5dcb80fd3d966802566e0d1cb59036a8fa53"></a>

<a id="canonical-1e96efe9c0aa33ab1177cee45babf836c802f55f5e52a38ea5030a02abd4c8da"></a>

## invert_matcher property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher / db84bfc8526f / 4

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [prefix_sets](resources--http_loadbalancer--reference--group-021.md#canonical-429194fc2e55daf1a21a789a66071032f61326357a8feb638ea7067dbedaa7a9): complete subsection reference.

<a id="canonical-bd9489cd5f3a405a8fb83857367adbb6470d604e729104dfad453e08bb090f91"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher / db84bfc8526f / 5

- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-021.md#canonical-429194fc2e55daf1a21a789a66071032f61326357a8feb638ea7067dbedaa7a9)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-429194fc2e55daf1a21a789a66071032f61326357a8feb638ea7067dbedaa7a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8de6ab407199d75ee8bb2157bb469a32d7aef7a8762deb2fe7fb84ecfb01155d"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / 8f8cf6c4f460 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](resources--http_loadbalancer--reference--group-021.md#canonical-2366270da3a20b5d81ebbbe31f9b33db5be09f979629a5afbde5481ddd37fd1f)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets

<a id="canonical-c04c69cd001ce6994c89646b46312e2481bac92b62ba347c42360e31e03d3d32"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-30a3d43c6affe8b21034358bccda4fd67de52becbd53e977761481331eed6145"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / 8f8cf6c4f460 / 3

<a id="canonical-01939e382da26f99aea930ddcff32f558f120b3dd0bcda02706d3dd881a2168d"></a>

<a id="canonical-a0d6ba4371f5d7e60c796aa5a11d1da408fc90d6dd26f6568e16c9c5f3f3a263"></a>

## kind property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / 8f8cf6c4f460 / 4

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

<a id="canonical-5ce5623be6d85bcfb04ae075d0fdddc0be94ad75fe352457b815235c0b2ad2b2"></a>

<a id="canonical-8fe3ba0f63b23bfd0c529666fda3576b5f4277be21c8cb715e277b83387a8012"></a>

## name property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / 8f8cf6c4f460 / 5

Type: `"string"`. Optional.

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

<a id="canonical-67f6be19f27c221e52fd0be50c7400c5b97e66244447fa8063ecb25365b05f4f"></a>

<a id="canonical-737afe807e4b5978614fec03ed00b993de0b101ce4229068c8a92d0c0f76db77"></a>

## namespace property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / 8f8cf6c4f460 / 6

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
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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

<a id="canonical-7a87b1f7a7e30eb0b14226d25bc728281fa9f9f915cb05b2af0781faf3d160e1"></a>

<a id="canonical-fc89b916afe1479a2e085a887d504c2332c51943994e4fd1ba2b596f7a11ef71"></a>

## tenant property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / 8f8cf6c4f460 / 7

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

<a id="canonical-94434aec19a05bc61dcc6b6a82472fe1a8b1b7b3fcb973f2936b3e5526596a21"></a>

<a id="canonical-74b7e3627b47d93afdd76cad483bff0e473c3c3a13f16b4e50e31a480109687a"></a>

## uid property — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / 8f8cf6c4f460 / 8

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

<a id="canonical-4b16b21681edb7ff75aa5d543e32aa58610df5c539e61bace1f9ffc51c52c2ac"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher.prefix_sets / 8f8cf6c4f460 / 9

- [origin_server_subset_rule_list.origin_server_subset_rules.ip_matcher](resources--http_loadbalancer--reference--group-021.md#canonical-2366270da3a20b5d81ebbbe31f9b33db5be09f979629a5afbde5481ddd37fd1f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-154b80b22444dfb0d70cb0a3da95834f09f977ec18cf944ddf7a66e7c54ebcb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3e83b7f83f8437d05f0f63ad154dabeee1c24c5f9fa929869d6c6d2e1b61bb4"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list — origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list / 04d59bee12ca / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list

<a id="canonical-ba35df0c8585d50aafd609b8c18bbb62ba45cf222897e8b1cabc1db133ad49d5"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

Receipt-pinned upstream constraints:

```json
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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-03adab9dcaaef5d7adbbd0a5c55401e60e4743a4e29f41a30d3cd0adea4ac5e3"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list / 04d59bee12ca / 3

<a id="canonical-3513556e46fdc716f16150397b3bada2c233fa1ca2b6e38b86d9e109e4fdfc8d"></a>

<a id="canonical-96764e3b6d14ebd999ffa6cf5a97fcb0b597f6bce05e3ddf21503e7f642f62cf"></a>

## invert_match property — origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list / 04d59bee12ca / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-11b0835130048e6ae66d0d365443fd3216ac0b73f9af8200b85448d17e2b71a2"></a>

<a id="canonical-f750a3e94d905fec883958ab11d63cc540c03c785eb28122aacdce60426f4618"></a>

## ip_prefixes property — origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list / 04d59bee12ca / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5bc5cce7fdc6e452da0fcac380ef48b57e7661c23560a8110c4e4c561fa3c571"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.ip_prefix_list / 04d59bee12ca / 6

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-08a935119871c441b3473ce4a32f4cc511826867746ff4ecdf676cbeb05e569d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64a8fa0ee19c6a851e235cb057d5ea9a6e2d2c20da0baebc3e4dab460d285369"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.metadata — origin_server_subset_rule_list.origin_server_subset_rules.metadata / 2a2deaeb1d67 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- origin_server_subset_rule_list.origin_server_subset_rules.metadata

<a id="canonical-f4ba2b2071ae9a6e0e702ef919826c9162c3f55daa57b867aad0c2c560c1679f"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-9412668fd857035fdf2ea37f680fbec9d0f2747ab5d16905fe1d0b39ada899b4"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.metadata / 2a2deaeb1d67 / 3

<a id="canonical-cf6e4666b8117cec1d4c4ebcddaf77d94d96e7f95fd0ab01fe46f8df33a0aacd"></a>

<a id="canonical-d802aff9d4629e33f2496b900ebef21faefc02a144c8de8f88987d9a6da351e4"></a>

## description_spec property — origin_server_subset_rule_list.origin_server_subset_rules.metadata / 2a2deaeb1d67 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-cd6e09521cc823d4e191a1d337639874bb028522217cce2e40cb2c17974e230f"></a>

<a id="canonical-65f6101f44619bbead4e72a918fcb634e662ad7dc8540f5b5d277e3f526c2d70"></a>

## name property — origin_server_subset_rule_list.origin_server_subset_rules.metadata / 2a2deaeb1d67 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0d924eccfe664d848bbb9ff865fd899930d83e8a8461c8abd06721bcc084bd31"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.metadata / 2a2deaeb1d67 / 6

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a5b3bfb0fc9514908b1ba0435689dbe7b92a540d10e6908d7fd00b57120aedeb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6bd32b63477d1e31c57157d398474e983cbf847dcea20f9ef1adf4da92641326"></a>

## origin_server_subset_rule_list.origin_server_subset_rules.none — origin_server_subset_rule_list.origin_server_subset_rules.none / 070024a8630b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-27f040f8cfe68609b89b4390ae69c1789f9aad0d772f1040ff41a92c5dbc8670)
- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- origin_server_subset_rule_list.origin_server_subset_rules.none

<a id="canonical-2ca055aa3d75bd1355b31a450410354b7f450160f594b18ef3d34e3cc2e8583d"></a>

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
none = {}
```

<a id="canonical-2b96a3ac096b968f997ad37121f98c5c3ac3afae4430a17a3385b3656afbc8c8"></a>

## Direct properties — origin_server_subset_rule_list.origin_server_subset_rules.none / 070024a8630b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7d892a07194189bb18e664d7b97b7a745a0c45cc60ad2c55a7564839af7c4386"></a>

## Next pages — origin_server_subset_rule_list.origin_server_subset_rules.none / 070024a8630b / 4

- [origin_server_subset_rule_list.origin_server_subset_rules](resources--http_loadbalancer--reference--group-021.md#canonical-1121957c1d59150cb3478cb78482df78bc39527cfe598b01db57af0fd971e264)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d740c03a40238dc6db071b29b42b0bf983e2b8a6e5810d4049505af4d5556c6e"></a>

## policy_based_challenge — policy_based_challenge / 0d0734ca2412 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- policy_based_challenge

<a id="canonical-39ad349c197b736b689021789fb7cd999d794a0277cbe70b543cc70730b18cae"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings for policy rule based challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "always_enable_js_challenge"),
  validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("always_enable_js_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("captcha_challenge_parameters",
    "default_captcha_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_js_challenge_parameters",
    "js_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_mitigation_settings",
    "malicious_user_mitigation"),
  validators.ConflictingObjectAttributes("default_temporary_blocking_parameters",
    "temporary_user_blocking")}
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
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-challenge_choice": "[\"always_enable_captcha_challenge\",\"always_enable_js_challenge\",\"no_challenge\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]",
  "x-ves-oneof-field-temporary_blocking_parameters_choice": "[\"default_temporary_blocking_parameters\",\"temporary_user_blocking\"]"
}
```

Terraform syntax:

```terraform
policy_based_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-fa71f5733547a039a8a3cb285a8c473945914dacbfc23546bc26c14e693834c3"></a>

## Direct properties — policy_based_challenge / 0d0734ca2412 / 3

- [always_enable_captcha_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-484e6b24321f1b207c403e86337daa9698bcef128bf9d92e814d31ee2bb3b46a): complete subsection reference.

- [always_enable_js_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-88463651da74c67e01bc4c69a4177e8dac0044d5baeb87e5e547f8fb17f1d41b): complete subsection reference.

- [captcha_challenge_parameters](resources--http_loadbalancer--reference--group-021.md#canonical-33b6b024ed40d43d99f1509bab5aabdb0bef6fa24c5c271563f70b1fb8e25136): complete subsection reference.

- [default_captcha_challenge_parameters](resources--http_loadbalancer--reference--group-021.md#canonical-d08c5745e51dfbdfa701a5f1493826d8b60b62c8c47d1d575abb667b4c419b45): complete subsection reference.

- [default_js_challenge_parameters](resources--http_loadbalancer--reference--group-021.md#canonical-13cfd9c5a2ecdedb12f37df42e23bb47fb04e084f26708ac5b18fc1492de31c2): complete subsection reference.

- [default_mitigation_settings](resources--http_loadbalancer--reference--group-021.md#canonical-8ef6e11cceb81f89e472418800e3af66247fbd50e015c2f3bda4b50d2af98092): complete subsection reference.

- [default_temporary_blocking_parameters](resources--http_loadbalancer--reference--group-021.md#canonical-897e14bb865a57ef8812a0b4478debd3fe7824cb05196dbdc9db9d94f081289f): complete subsection reference.

- [js_challenge_parameters](resources--http_loadbalancer--reference--group-021.md#canonical-ad1f0fa48d596dcca93b94b61eaecdc979cd0e8b8f4dfec8da258b244745b0af): complete subsection reference.

- [malicious_user_mitigation](resources--http_loadbalancer--reference--group-021.md#canonical-7a51c08d2b28f1d93e51d1d80079be669bb8c5812d1a4c2ae2f7c15296ab3dba): complete subsection reference.

- [no_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-ae900055f2f2aeb75dea2bb6d0b0fcefb1bc5a397719826ada883fd09944e969): complete subsection reference.

- [rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf): complete subsection reference.

- [temporary_user_blocking](resources--http_loadbalancer--reference--group-022.md#canonical-3ee94511dbac7c8b39832e21b65490680ec7d621b40e6b37977278368aaaed41): complete subsection reference.

<a id="canonical-32b75f60e9cd7bfeab4f1057bb66a74ce23a17863986c7942634657fe7a16d7f"></a>

## Next pages — policy_based_challenge / 0d0734ca2412 / 4

- [policy_based_challenge.always_enable_captcha_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-484e6b24321f1b207c403e86337daa9698bcef128bf9d92e814d31ee2bb3b46a)
- [policy_based_challenge.always_enable_js_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-88463651da74c67e01bc4c69a4177e8dac0044d5baeb87e5e547f8fb17f1d41b)
- [policy_based_challenge.captcha_challenge_parameters](resources--http_loadbalancer--reference--group-021.md#canonical-33b6b024ed40d43d99f1509bab5aabdb0bef6fa24c5c271563f70b1fb8e25136)
- [policy_based_challenge.default_captcha_challenge_parameters](resources--http_loadbalancer--reference--group-021.md#canonical-d08c5745e51dfbdfa701a5f1493826d8b60b62c8c47d1d575abb667b4c419b45)
- [policy_based_challenge.default_js_challenge_parameters](resources--http_loadbalancer--reference--group-021.md#canonical-13cfd9c5a2ecdedb12f37df42e23bb47fb04e084f26708ac5b18fc1492de31c2)
- [policy_based_challenge.default_mitigation_settings](resources--http_loadbalancer--reference--group-021.md#canonical-8ef6e11cceb81f89e472418800e3af66247fbd50e015c2f3bda4b50d2af98092)
- [policy_based_challenge.default_temporary_blocking_parameters](resources--http_loadbalancer--reference--group-021.md#canonical-897e14bb865a57ef8812a0b4478debd3fe7824cb05196dbdc9db9d94f081289f)
- [policy_based_challenge.js_challenge_parameters](resources--http_loadbalancer--reference--group-021.md#canonical-ad1f0fa48d596dcca93b94b61eaecdc979cd0e8b8f4dfec8da258b244745b0af)
- [policy_based_challenge.malicious_user_mitigation](resources--http_loadbalancer--reference--group-021.md#canonical-7a51c08d2b28f1d93e51d1d80079be669bb8c5812d1a4c2ae2f7c15296ab3dba)
- [policy_based_challenge.no_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-ae900055f2f2aeb75dea2bb6d0b0fcefb1bc5a397719826ada883fd09944e969)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.temporary_user_blocking](resources--http_loadbalancer--reference--group-022.md#canonical-3ee94511dbac7c8b39832e21b65490680ec7d621b40e6b37977278368aaaed41)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-484e6b24321f1b207c403e86337daa9698bcef128bf9d92e814d31ee2bb3b46a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da26b19aaee3b342556faa0336326c6ac751e2bf49874abad16a3b7bbcd09596"></a>

## policy_based_challenge.always_enable_captcha_challenge — policy_based_challenge.always_enable_captcha_challenge / ee0d04024691 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.always_enable_captcha_challenge

<a id="canonical-e2c6223ad2d8b733e5d93de303529ad7c9450b71fb753129d3c061cc0e8e652f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable captcha challenge.

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
always_enable_captcha_challenge = {}
```

<a id="canonical-178b3c6d2d95503b4dcf03410aafe70b3b51da77f5a6d279bc0952fb6877d824"></a>

## Direct properties — policy_based_challenge.always_enable_captcha_challenge / ee0d04024691 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-075b6c56819caeb201066c17a4d78fd1009eb6a5bad56369e55e6f5600412c40"></a>

## Next pages — policy_based_challenge.always_enable_captcha_challenge / ee0d04024691 / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-88463651da74c67e01bc4c69a4177e8dac0044d5baeb87e5e547f8fb17f1d41b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43eccf34d8abb195722004de1fe51d30720b23fe09a3faab9cdeef2ac8c97003"></a>

## policy_based_challenge.always_enable_js_challenge — policy_based_challenge.always_enable_js_challenge / 8b86de6fe27b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.always_enable_js_challenge

<a id="canonical-e2a685c9e7d22e0605b8f8f3a337471cb33b0c41870f6393114e8e380722d69d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable js challenge.

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
always_enable_js_challenge = {}
```

<a id="canonical-d723d65c50288d989b6093075689ed3a9a536b879ab2bfb183663c2bb4d9d50c"></a>

## Direct properties — policy_based_challenge.always_enable_js_challenge / 8b86de6fe27b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d0c553dddffdd240a0d74086c33875cabe38746cb70d1f52148e6916818ac79"></a>

## Next pages — policy_based_challenge.always_enable_js_challenge / 8b86de6fe27b / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-33b6b024ed40d43d99f1509bab5aabdb0bef6fa24c5c271563f70b1fb8e25136"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a42be1066bb7a2afbccb62d4b49ea3790015dcf72eccaa7af91009866d7f8e98"></a>

## policy_based_challenge.captcha_challenge_parameters — policy_based_challenge.captcha_challenge_parameters / 76a3fe97ea16 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-c90fb6d615d84782c262cc5d5a55ff73317e213cd167992eee2119e7bc2b93e3"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
captcha_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-5ab21c785801dd4dd7eb974ef90d72f4a6c352118b8cf324da8654bfdd4d5be2"></a>

## Direct properties — policy_based_challenge.captcha_challenge_parameters / 76a3fe97ea16 / 3

<a id="canonical-ccc195d5195c61db499ce3e0d8166d671caef425b561e710aabcd1693063448d"></a>

<a id="canonical-36849ab2d0081055497e49c8f1bdf2e0063b937e16169ffda530abb9d16137cf"></a>

## cookie_expiry property — policy_based_challenge.captcha_challenge_parameters / 76a3fe97ea16 / 4

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

<a id="canonical-be1c1d4a56ed9a937fac0b0ed5ac9c2321823cb5215b2a18a23ea3d673de6339"></a>

<a id="canonical-cb42efb07e0820e6ff49d05973854631c5ba98fc3d126316bd2fcf1287811660"></a>

## custom_page property — policy_based_challenge.captcha_challenge_parameters / 76a3fe97ea16 / 5

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

<a id="canonical-cb9aeb22de43c9909daf9145ca5eb63e128d4cfa624c23a91ce9528d49a419bd"></a>

## Next pages — policy_based_challenge.captcha_challenge_parameters / 76a3fe97ea16 / 6

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d08c5745e51dfbdfa701a5f1493826d8b60b62c8c47d1d575abb667b4c419b45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7003a2995fb9ad772bf342ea7d71a57d1b589b4d424e6845d2e2d6c2a891432f"></a>

## policy_based_challenge.default_captcha_challenge_parameters — policy_based_challenge.default_captcha_challenge_parameters / e9a7fb2083ee / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-4aaf8fc47bd248d97f96bd459e2d7a975ec667fd0b80887dea96bd09ade4f50c"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_captcha_challenge_parameters = {}
```

<a id="canonical-98189ff3f997a3ef877595d771824f66a2967df41cea17ea9b1dc309a4176a02"></a>

## Direct properties — policy_based_challenge.default_captcha_challenge_parameters / e9a7fb2083ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-de76bdb5bbcb493f3848bb2753dc183a170b8d28a95936e8b96e284cccf31740"></a>

## Next pages — policy_based_challenge.default_captcha_challenge_parameters / e9a7fb2083ee / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-13cfd9c5a2ecdedb12f37df42e23bb47fb04e084f26708ac5b18fc1492de31c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1f4877fc28806a8da68b5ef45a16f001e6e5c466cb39ee443d0597fa88a68ca"></a>

## policy_based_challenge.default_js_challenge_parameters — policy_based_challenge.default_js_challenge_parameters / 2835bdc427a9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-fbacee4a2407d9d3d3fcb437099ffbb14929587755185a47beced7eea98e784f"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_js_challenge_parameters = {}
```

<a id="canonical-5d872534afe87851cd68362165220e473bdec8ba248ca43135e77091257439b9"></a>

## Direct properties — policy_based_challenge.default_js_challenge_parameters / 2835bdc427a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ce8f6e3ccce549f335b17f2f95a7a91988b755736b11f3f08e055d2a0c9277d3"></a>

## Next pages — policy_based_challenge.default_js_challenge_parameters / 2835bdc427a9 / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8ef6e11cceb81f89e472418800e3af66247fbd50e015c2f3bda4b50d2af98092"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0117c8e3435d3eda32d9ad003c5ed0cdb702c4d07ec13f3d5c7ae796d9d52a4"></a>

## policy_based_challenge.default_mitigation_settings — policy_based_challenge.default_mitigation_settings / 553be748714c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-20558de03205c093acbfff7902ad15480711f9b308e37acc96a3d9914fbfc6a3"></a>

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
default_mitigation_settings = {}
```

<a id="canonical-5f464fc27134b7daa14af779c3041ae29d06aa95b91a2f2cb42ac257c43d66a1"></a>

## Direct properties — policy_based_challenge.default_mitigation_settings / 553be748714c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34e724af7c86d2305dea18fc79739e580f15388931159d6c37492ea8fad939c7"></a>

## Next pages — policy_based_challenge.default_mitigation_settings / 553be748714c / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-897e14bb865a57ef8812a0b4478debd3fe7824cb05196dbdc9db9d94f081289f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fabe290456c6bc2bef0397c281739f39efe658a40010ed7417770d4ace6398ec"></a>

## policy_based_challenge.default_temporary_blocking_parameters — policy_based_challenge.default_temporary_blocking_parameters / 4251b6b21fe4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-36db8214bf73c93d3c2f21f630010d6775d30504798334d72c0568daa7fc6da2"></a>

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
default_temporary_blocking_parameters = {}
```

<a id="canonical-8d93c63929d28ad2444d39533b4fe85f2dc3bd36766530b5a228aaf98afe9889"></a>

## Direct properties — policy_based_challenge.default_temporary_blocking_parameters / 4251b6b21fe4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a49a843f2b59e421f5a06e3ff14b34eb1eca7971a8e734568ff3c718e8dda17d"></a>

## Next pages — policy_based_challenge.default_temporary_blocking_parameters / 4251b6b21fe4 / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ad1f0fa48d596dcca93b94b61eaecdc979cd0e8b8f4dfec8da258b244745b0af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0c1ee2a339ce30882ab3c6e730795ed4cf0d70fd5adba37835bb3d62aa984ed"></a>

## policy_based_challenge.js_challenge_parameters — policy_based_challenge.js_challenge_parameters / 80bc98880d59 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-724748e260ba4af3071ee4623ff8ac0faf3392cccb061c92fa16a009c9f6530c"></a>

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
js_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-8824987f19588c31f83535b40357079b68eb00557bdc6b2d461cdbbd959b8b72"></a>

## Direct properties — policy_based_challenge.js_challenge_parameters / 80bc98880d59 / 3

<a id="canonical-95957c34987d161ad99e2909ba7884f1e56aac07d179cf75dca66689cf1898da"></a>

<a id="canonical-eb31440ad63612ec23084a0a698ed129899a9c45ed67fce83cad8886fd7ad111"></a>

## cookie_expiry property — policy_based_challenge.js_challenge_parameters / 80bc98880d59 / 4

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

<a id="canonical-e36b7b5978ab86e262b2032fdc61b504d637de7588ab108653008bf4aee0984b"></a>

<a id="canonical-707b738a4e83bfc63fe8570382f0be4cc981539bbc39d7ff9bb2d4a1d40ab132"></a>

## custom_page property — policy_based_challenge.js_challenge_parameters / 80bc98880d59 / 5

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

<a id="canonical-c82dbcabc3ed58f09e7c830b599aaeb738849bde07300581f89b0f82a79edaa6"></a>

<a id="canonical-80ddd2f21ae98a4f6bbd2d36d3d5fe047faacebec20d600478cc62e92295752d"></a>

## js_script_delay property — policy_based_challenge.js_challenge_parameters / 80bc98880d59 / 6

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

<a id="canonical-99925c7732f800f6fded264306952ffae32031915abd81f00b656f331c64235a"></a>

## Next pages — policy_based_challenge.js_challenge_parameters / 80bc98880d59 / 7

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7a51c08d2b28f1d93e51d1d80079be669bb8c5812d1a4c2ae2f7c15296ab3dba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fba28e5c79a8ab72a52b8be0cc8aba16b7e34b9fec901b9c734459ce899e9606"></a>

## policy_based_challenge.malicious_user_mitigation — policy_based_challenge.malicious_user_mitigation / d6cfe74b139f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-145fc2215ab20665ac9c9e9dbb958f2ac0f73c07a700d24c1a6cb464907bbe76"></a>

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
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-66b15f36c9adf35172c4b7ee34218d277bb78530f1742260321d23ed4657dbc2"></a>

## Direct properties — policy_based_challenge.malicious_user_mitigation / d6cfe74b139f / 3

<a id="canonical-8c6182db6f17865049dcc1f85b9b35c86aedd95e42352c602c0209af950c547b"></a>

<a id="canonical-c6e29f177572822e3b7f7451527f5fa8f352e7a31f2ca65b848cececec3cb6df"></a>

## name property — policy_based_challenge.malicious_user_mitigation / d6cfe74b139f / 4

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

<a id="canonical-a38918d091132967e7cadf767bd621956641ed91269ef36c3388e1c561485582"></a>

<a id="canonical-f0308b71d47dd727f077a2d9c11fadcc3af514b627bdf01b7716c36a7fc397c3"></a>

## namespace property — policy_based_challenge.malicious_user_mitigation / d6cfe74b139f / 5

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

<a id="canonical-6a283f69d878dbef8e887e7b3488e78777d80f06373abb463569a37d5b769b51"></a>

<a id="canonical-81b7a294be664eb32e0886989f84c3b79bf19bb203b24dbe5acc2370dacdee1d"></a>

## tenant property — policy_based_challenge.malicious_user_mitigation / d6cfe74b139f / 6

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

<a id="canonical-ff8c8cb63a64e82fc3cc51759684d3730275477cd6cbc32cd7c0d8ad66ba7751"></a>

## Next pages — policy_based_challenge.malicious_user_mitigation / d6cfe74b139f / 7

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ae900055f2f2aeb75dea2bb6d0b0fcefb1bc5a397719826ada883fd09944e969"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f3ac172d7951020daf56630969e04ad7e08aa4a1cbabc8826a3d80ef7146b35"></a>

## policy_based_challenge.no_challenge — policy_based_challenge.no_challenge / dad44a29d35f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.no_challenge

<a id="canonical-fa2123cab0a36cbc1d6afdc43c00c31a8e6b84c4a18f0077267ec0f9f254da94"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_challenge = {}
```

<a id="canonical-33cf29c0e0b7a896b577b316c9723ef2acefdb8a4c7ac691360a79c47098d60f"></a>

## Direct properties — policy_based_challenge.no_challenge / dad44a29d35f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e85be8611d2be23cf9fbea2a208bd0039762738b36a59463c7abb3aa71af983c"></a>

## Next pages — policy_based_challenge.no_challenge / dad44a29d35f / 4

- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d1c597e2691c866a3105dd3ba4d0a2db4a68b6c6f87cfad6ff8eebf0945e42d"></a>

## policy_based_challenge.rule_list — policy_based_challenge.rule_list / 7d9ea893e946 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- policy_based_challenge.rule_list

<a id="canonical-4981762d4e9d16238ad92cdade8110d4bee1b3c5a20257c338a5f422e0fc1b04"></a>

Type: `"object"`. single nested block, Optional.

List of challenge rules to be used in policy based challenge.

Receipt-pinned upstream constraints:

```json
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
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3d0aeca8be6218fae8bc358706493813a7871fa99fae94bfbb64c02610413abd"></a>

## Direct properties — policy_based_challenge.rule_list / 7d9ea893e946 / 3

- [rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c): complete subsection reference.

<a id="canonical-3fcc27e72a720d2e0b49dc15dad26f74579a1a5e30aea31efa861a8488a4cc24"></a>

## Next pages — policy_based_challenge.rule_list / 7d9ea893e946 / 4

- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92c77d3610966329967c6e70a150d9a1917ca70ffa76512a22bd6df879084d12"></a>

## policy_based_challenge.rule_list.rules — policy_based_challenge.rule_list.rules / 7747c62a24d9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- policy_based_challenge.rule_list.rules

<a id="canonical-498243aef38dbbaed386eea7e071b25db582fb98c483276a301f47eaa23cafae"></a>

Type: `"object"`. list nested block, Optional.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Upstream description:

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1f7ec841ecd81e382e21028ac03f229c1b0ed5dd4e40d7935c4c704ed670a0b8"></a>

## Direct properties — policy_based_challenge.rule_list.rules / 7747c62a24d9 / 3

- [metadata](resources--http_loadbalancer--reference--group-021.md#canonical-32686832245f654c26b40fc5a03f03ffa8d4025bbeebc27407b417cd08e079b2): complete subsection reference.

- [spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6): complete subsection reference.

<a id="canonical-3189eb04299cdbef9e95c93baed6f0a753157a4057b26deb741e510a272b226d"></a>

## Next pages — policy_based_challenge.rule_list.rules / 7747c62a24d9 / 4

- [policy_based_challenge.rule_list.rules.metadata](resources--http_loadbalancer--reference--group-021.md#canonical-32686832245f654c26b40fc5a03f03ffa8d4025bbeebc27407b417cd08e079b2)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-4ba0dd0b848a0cac5f1e5893b11d8a92ba049ca024478259a23668e76c2cbae6)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-32686832245f654c26b40fc5a03f03ffa8d4025bbeebc27407b417cd08e079b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b9241344238810174db2d1514b431753641505cbd339f4fe77be4a6c1d43995"></a>

## policy_based_challenge.rule_list.rules.metadata — policy_based_challenge.rule_list.rules.metadata / 3ea242021dd9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-79948879354188eafa8c0e802af344409a251cdd860384e8fc6ff710a41d9715)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-021.md#canonical-c1ee5fd623a9d08c1d6cc7aa817a0512bb72c111d1ddc441ea44f3c781e3d0cf)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-021.md#canonical-444b3c44b1a1efe37a0bf02e172f10a2ec6d239abc62f6c26a265fb61967f34c)
- policy_based_challenge.rule_list.rules.metadata

<a id="canonical-dd43314fa03ee785cdb6c812ba31f7a7c7c5d352ed258f990e3baa80b302fa17"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-e7ec743b925d25781a323697b9b1a61cc1396d9d979e9cb6a1ddf06a0b1460bc"></a>

## Direct properties — policy_based_challenge.rule_list.rules.metadata / 3ea242021dd9 / 3

<a id="canonical-78bb651c92db0a0d754eae8c6f640b41478bd9a5b65a1a6647ea9abe437e4352"></a>
