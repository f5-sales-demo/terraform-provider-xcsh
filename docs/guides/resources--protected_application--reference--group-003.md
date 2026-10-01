---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-40bed5a89b1dcca86e70a9350861deeebd132dee690b6bb37839f2a416a30ab8"></a>

## cloudfront.mobile_sdk_config.mobile_identifier.headers — cloudfront.mobile_sdk_config.mobile_identifier.headers / 6ef9d0e50e49 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-002.md#canonical-cf21ec7f03b788f62c0fa1ef9f8f60e2f002fbd4c183bb98c97b975415b39209)
- [cloudfront.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-002.md#canonical-27af8a569a472b2a77c5ea06a6c1ebdbcd31bf63c37fe6714d01b6ef4f90e0db)
- cloudfront.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-04c7d044f238d79f67100a3bb6da6bf1bccbcdd0147d41012babbfcce53865bf"></a>

Type: `"object"`. list nested block, Optional.

List of headers that can be used to identify mobile traffic.

Upstream description:

A list of headers that can be used to identify mobile traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "regex")}
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-963bf2740a2f43e4a7b2441a2d6858c2d8585397506afc1da2dfab2feae174a1"></a>

## Direct properties — cloudfront.mobile_sdk_config.mobile_identifier.headers / 6ef9d0e50e49 / 3

<a id="canonical-54da92f86f18fa1389e2a1f6dda96c02da1f7c4a98e6babedeacb820105de352"></a>

<a id="canonical-d548b6074a831d2143ec3ba21adfd47101fd4fa08d20c4f46e3db3ce922f4501"></a>

## exact property — cloudfront.mobile_sdk_config.mobile_identifier.headers / 6ef9d0e50e49 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\] Header value to match exactly.

Upstream description:

Exclusive with \[regex\] Header value to match exactly.

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
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-5f045cf278bff1e8dcacb7c0bf30d751cebf46f56e6b50623032bef587269576"></a>

<a id="canonical-076555b5ad34b79e4dcfd4fbddfd8a7eb17c4a8e1cf6f07233eb79f4cad977a8"></a>

## name property — cloudfront.mobile_sdk_config.mobile_identifier.headers / 6ef9d0e50e49 / 5

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-22227d571b6ce6051d3a223bf45f1c36754b6a888a351b6bac696466e118cb2c"></a>

<a id="canonical-05684e29b918bb6defa0bf77116fa433ce1b745f352733b51ba147526d1b0cb8"></a>

## regex property — cloudfront.mobile_sdk_config.mobile_identifier.headers / 6ef9d0e50e49 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact\] Regex match of the header value in re2 format.

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
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-db22bf0ddc2fe18db119e6fba7116643e1f5bdb680c198240d87e660c1dc8b40"></a>

## Next pages — cloudfront.mobile_sdk_config.mobile_identifier.headers / 6ef9d0e50e49 / 7

- [cloudfront.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-002.md#canonical-27af8a569a472b2a77c5ea06a6c1ebdbcd31bf63c37fe6714d01b6ef4f90e0db)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-808bc6b1b372f873cce54fcae939187734656d4aa4770ca9ff5b467f70a385b4"></a>

## cloudfront.protected_endpoints — cloudfront.protected_endpoints / 575ad7c8ee08 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- cloudfront.protected_endpoints

<a id="canonical-1ce96a8cc4d2f9ec2a090b4716d1f65513a8d3a73fb0c4e9e7f8c43e0cce7267"></a>

Type: `"object"`. list nested block, Optional.

List of protected endpoints (max 128 items).

Upstream description:

List of protected endpoints (max 128 items)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("http_methods",
    "path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("flow_label",
    "undefined_flow_label"),
  validators.ConflictingListObjectAttributes("mobile_client",
    "web_client"),
  validators.ConflictingListObjectAttributes("mobile_client",
    "web_mobile_client"),
  validators.ConflictingListObjectAttributes("web_client",
    "web_mobile_client")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protected_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-a70154d4eea0a2a2709e06672151333956f7a73c65bcbfdac063dc8e425b06a5"></a>

## Direct properties — cloudfront.protected_endpoints / 575ad7c8ee08 / 3

- [any_domain](resources--protected_application--reference--group-003.md#canonical-6d45abe321aefd2049e915ba7069d2d76728b1427abd37a8e645b905874d5c7a): complete subsection reference.

- [domain](resources--protected_application--reference--group-003.md#canonical-89c8aea3619dc467198f26955f8a77305957c99d26bb90cdf877b7f13b37442a): complete subsection reference.

- [flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9): complete subsection reference.

<a id="canonical-f61764d96305487347982735124b327c670e50226cb024ee9a9f5df51c2b6e2e"></a>

<a id="canonical-5c11de62461bafb1536f596d483e7626afe1331f825d7ddc2336c485d6008ecd"></a>

## http_methods property — cloudfront.protected_endpoints / 575ad7c8ee08 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](resources--protected_application--reference--group-003.md#canonical-26572fbea0b5a97f998e5cd33dc50ad5cec6eb52e8a1fe2fdedd082e0d49f516): complete subsection reference.

- [mobile_client](resources--protected_application--reference--group-003.md#canonical-27b837e1a76705986c2e896c7defd65dfaf775f5eda460fa07a12bb07fc5b5a4): complete subsection reference.

<a id="canonical-269779a9ada28bbbb11d02b36ea09b58d89ef7317462dd5d78d59183e395dbdd"></a>

<a id="canonical-3258c32e61abc54a96b3fdb59a0c0fb9cfd6abb567e5f1a26f4b25c1ec940b47"></a>

## path property — cloudfront.protected_endpoints / 575ad7c8ee08 / 5

Type: `"string"`. Optional.

Accepts wildcards \* to match multiple characters or ? To match a single character.

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
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  }
}
```

<a id="canonical-5de944d0b4b591bde214466448d5a5d86e37a89b1803bc66193995438d04af12"></a>

<a id="canonical-e4a2fcd34a4823b86775a6d130299818fa282603582d5c4ff4ee7df4515c777c"></a>

## query property — cloudfront.protected_endpoints / 575ad7c8ee08 / 6

Type: `"string"`. Optional.

Enter a regular expression to match your query parameters of interest.

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
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

- [undefined_flow_label](resources--protected_application--reference--group-004.md#canonical-1d2d3979782ee14890ffb72d01eafe38fd61faa7ab83a4a97646f1c687e2c16a): complete subsection reference.

- [web_client](resources--protected_application--reference--group-004.md#canonical-487268b3c34843ee4a74d5aa3f11492d3122613df09482c1e71725f1b44d504f): complete subsection reference.

- [web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16): complete subsection reference.

<a id="canonical-598ee836f83d98fcdf39b55d6349b93e7635995dfa64f41720025e0578ba3eb9"></a>

## Next pages — cloudfront.protected_endpoints / 575ad7c8ee08 / 7

- [cloudfront.protected_endpoints.any_domain](resources--protected_application--reference--group-003.md#canonical-6d45abe321aefd2049e915ba7069d2d76728b1427abd37a8e645b905874d5c7a)
- [cloudfront.protected_endpoints.domain](resources--protected_application--reference--group-003.md#canonical-89c8aea3619dc467198f26955f8a77305957c99d26bb90cdf877b7f13b37442a)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.metadata](resources--protected_application--reference--group-003.md#canonical-26572fbea0b5a97f998e5cd33dc50ad5cec6eb52e8a1fe2fdedd082e0d49f516)
- [cloudfront.protected_endpoints.mobile_client](resources--protected_application--reference--group-003.md#canonical-27b837e1a76705986c2e896c7defd65dfaf775f5eda460fa07a12bb07fc5b5a4)
- [cloudfront.protected_endpoints.undefined_flow_label](resources--protected_application--reference--group-004.md#canonical-1d2d3979782ee14890ffb72d01eafe38fd61faa7ab83a4a97646f1c687e2c16a)
- [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-487268b3c34843ee4a74d5aa3f11492d3122613df09482c1e71725f1b44d504f)
- [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-b18b3283f43ec7feb019c586651cbdeafbe29d71e9d0cba61f6479f1d5e06a16)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-6d45abe321aefd2049e915ba7069d2d76728b1427abd37a8e645b905874d5c7a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-645567b78d0fa84bdb76675d0474e0d55ae71c0735164d495f905036f2b6cd45"></a>

## cloudfront.protected_endpoints.any_domain — cloudfront.protected_endpoints.any_domain / 407d77a68a47 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- cloudfront.protected_endpoints.any_domain

<a id="canonical-12f7886fee20e40a393d9b183b378bdfda4fc514af54f5aa97c8c574c4453917"></a>

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
any_domain = {}
```

<a id="canonical-19af71f1c3b811b6fcaf33e8e6593e9154778fcb6ffc9c02235cb91225e482f2"></a>

## Direct properties — cloudfront.protected_endpoints.any_domain / 407d77a68a47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4b4eec05a35c28f7c0607c2212700cc3c819e83f799dc12d6cc469ea2abf6ce9"></a>

## Next pages — cloudfront.protected_endpoints.any_domain / 407d77a68a47 / 4

- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-89c8aea3619dc467198f26955f8a77305957c99d26bb90cdf877b7f13b37442a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0494235b1abf18ba1a304f0d512cfb8d1cf0d34fba4e1e37c5b872104e50492"></a>

## cloudfront.protected_endpoints.domain — cloudfront.protected_endpoints.domain / b032d48e2dfa / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- cloudfront.protected_endpoints.domain

<a id="canonical-7753624b1b32d0d177d2d1120ebf9d4d78c31bcab5308dcb7fcf1cb4c413e161"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-ab4aaf5411f8eac771c2d7deef8f2a87398003e4b491464e3840d306d8bf1ba7"></a>

## Direct properties — cloudfront.protected_endpoints.domain / b032d48e2dfa / 3

<a id="canonical-d86af2565dd76157e11c899080d0051ede3c400d0ec9a6bd13ffb23bc1987193"></a>

<a id="canonical-7ae0ebc655eb6f5a567c039ea9609914f3d6d0780aff9261209dd57dba770940"></a>

## exact_value property — cloudfront.protected_endpoints.domain / b032d48e2dfa / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-63c643201218fcc4969e3319d4cab743a28c5530eb37aca0fdb7f8bc0ac9eacb"></a>

<a id="canonical-df2a657ae6456b13ebff5994569730da2f1089dbdd6f9197d89dc6529d91f005"></a>

## regex_value property — cloudfront.protected_endpoints.domain / b032d48e2dfa / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-e3bd38293b1fddad1c275c3ff18c2ebef6281b7ec602ef32a54d8df168670eda"></a>

<a id="canonical-dff6ad7ec74059c79b74775eab565bf4ada4a029659982cc7f7984e6211e7469"></a>

## suffix_value property — cloudfront.protected_endpoints.domain / b032d48e2dfa / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-75c3308469bb34c28ab19bf9d038990ac4f42ffcc9aa366e36cc0d65c27f798b"></a>

## Next pages — cloudfront.protected_endpoints.domain / b032d48e2dfa / 7

- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc07f4084c28e13610709db2818f88931de3d9ebb1207474948c4e5a74a28bd0"></a>

## cloudfront.protected_endpoints.flow_label — cloudfront.protected_endpoints.flow_label / 2d26f9507023 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- cloudfront.protected_endpoints.flow_label

<a id="canonical-83a32201ab330d6813e10fcd015cafec407fb4e2b3e4ba52bbd7cca2bf628849"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("account_management",
    "authentication"),
  validators.ConflictingObjectAttributes("account_management",
    "financial_services"),
  validators.ConflictingObjectAttributes("account_management",
    "flight"),
  validators.ConflictingObjectAttributes("account_management",
    "profile_management"),
  validators.ConflictingObjectAttributes("account_management",
    "search"),
  validators.ConflictingObjectAttributes("account_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("authentication",
    "financial_services"),
  validators.ConflictingObjectAttributes("authentication",
    "flight"),
  validators.ConflictingObjectAttributes("authentication",
    "profile_management"),
  validators.ConflictingObjectAttributes("authentication",
    "search"),
  validators.ConflictingObjectAttributes("authentication",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("financial_services",
    "flight"),
  validators.ConflictingObjectAttributes("financial_services",
    "profile_management"),
  validators.ConflictingObjectAttributes("financial_services",
    "search"),
  validators.ConflictingObjectAttributes("financial_services",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("flight",
    "profile_management"),
  validators.ConflictingObjectAttributes("flight",
    "search"),
  validators.ConflictingObjectAttributes("flight",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("profile_management",
    "search"),
  validators.ConflictingObjectAttributes("profile_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("search",
    "shopping_gift_cards")}
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
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

Terraform syntax:

```terraform
flow_label {
  # Configure direct properties listed below.
}
```

<a id="canonical-3db35576b930847408b93d60c7d09dad5a9a59ee740bedde70577c996d775338"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label / 2d26f9507023 / 3

- [account_management](resources--protected_application--reference--group-003.md#canonical-e545d2d85bda83648ae399f5d5bf57fdfdcd9e81d0ed390bac2a66a7af9c1dd2): complete subsection reference.

- [authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214): complete subsection reference.

- [financial_services](resources--protected_application--reference--group-003.md#canonical-692ea01fa276546cc2074827be4f4bf9519afd28943b52af84b6e8f3f53ce986): complete subsection reference.

- [flight](resources--protected_application--reference--group-003.md#canonical-6a1b2bf4c28b2364ffd480cad552a4190aa23617e63f7262544c676ea9e8a548): complete subsection reference.

- [profile_management](resources--protected_application--reference--group-003.md#canonical-e54910f6ce94de74a592ed705d1ddf46731acb511631278f5d75ea1169885bc7): complete subsection reference.

- [search](resources--protected_application--reference--group-003.md#canonical-69603b4df4c676280ea3b16f946a0121f7a8985bc8492753dbf140db99f2b88f): complete subsection reference.

- [shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b): complete subsection reference.

<a id="canonical-8d1e6f2191b711a91d7cefca8726a1c230302d84963b90639c3b9511d11b50eb"></a>

## Next pages — cloudfront.protected_endpoints.flow_label / 2d26f9507023 / 4

- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-e545d2d85bda83648ae399f5d5bf57fdfdcd9e81d0ed390bac2a66a7af9c1dd2)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-692ea01fa276546cc2074827be4f4bf9519afd28943b52af84b6e8f3f53ce986)
- [cloudfront.protected_endpoints.flow_label.flight](resources--protected_application--reference--group-003.md#canonical-6a1b2bf4c28b2364ffd480cad552a4190aa23617e63f7262544c676ea9e8a548)
- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-e54910f6ce94de74a592ed705d1ddf46731acb511631278f5d75ea1169885bc7)
- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-69603b4df4c676280ea3b16f946a0121f7a8985bc8492753dbf140db99f2b88f)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-e545d2d85bda83648ae399f5d5bf57fdfdcd9e81d0ed390bac2a66a7af9c1dd2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-284adf96eb96b972f83d9cf854486444fb477e07de49e527e276f759e6616122"></a>

## cloudfront.protected_endpoints.flow_label.account_management — cloudfront.protected_endpoints.flow_label.account_management / ad9ab6b6cde1 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- cloudfront.protected_endpoints.flow_label.account_management

<a id="canonical-aaafefc4910e854afdc5341410ece28265762c8b45eb839430b885a3e1b7dbd7"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Account Management Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "password_reset")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

Terraform syntax:

```terraform
account_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-3ff789b0ac67273a3f8e114935fdeb8c324e73f1a98bb9eb597ad8f98e8c4487"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.account_management / ad9ab6b6cde1 / 3

- [create](resources--protected_application--reference--group-003.md#canonical-0416d9720358267c0a18816853e81bd6f9c72af8a4191d7a513f8887691ba92d): complete subsection reference.

- [password_reset](resources--protected_application--reference--group-003.md#canonical-f5eef82c6f455c3ce07e2042c5407000577177e2ed2f158ed633d0ab250aaec2): complete subsection reference.

<a id="canonical-bd2082c416c5618f6ab16386e4a21dbbb1555db14f06926e45d9d383f2935799"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.account_management / ad9ab6b6cde1 / 4

- [cloudfront.protected_endpoints.flow_label.account_management.create](resources--protected_application--reference--group-003.md#canonical-0416d9720358267c0a18816853e81bd6f9c72af8a4191d7a513f8887691ba92d)
- [cloudfront.protected_endpoints.flow_label.account_management.password_reset](resources--protected_application--reference--group-003.md#canonical-f5eef82c6f455c3ce07e2042c5407000577177e2ed2f158ed633d0ab250aaec2)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-0416d9720358267c0a18816853e81bd6f9c72af8a4191d7a513f8887691ba92d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-415c3bb263a2de06f15b92dca20b49198a6f6ca769e30e9b7a40a0ba603f01a4"></a>

## cloudfront.protected_endpoints.flow_label.account_management.create — cloudfront.protected_endpoints.flow_label.account_management.create / 41e80c35434c / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-e545d2d85bda83648ae399f5d5bf57fdfdcd9e81d0ed390bac2a66a7af9c1dd2)
- cloudfront.protected_endpoints.flow_label.account_management.create

<a id="canonical-5b2eec6a67db95b939f1df26b638896f266c915709d8d38967130892093f6a76"></a>

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
create = {}
```

<a id="canonical-76aeeb3ce5de45f9144d43155b8f8789f22ccd6f624fb96acef93db7599357d1"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.account_management.create / 41e80c35434c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213d2b7f1e2fc536934c7f18b5e6285755d2b356bef446800365f7d449ea3fb"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.account_management.create / 41e80c35434c / 4

- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-e545d2d85bda83648ae399f5d5bf57fdfdcd9e81d0ed390bac2a66a7af9c1dd2)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-f5eef82c6f455c3ce07e2042c5407000577177e2ed2f158ed633d0ab250aaec2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c27a6c0eb449c73b6b69239a663b64c0b9036e4c6320034e8445b563bace7a7"></a>

## cloudfront.protected_endpoints.flow_label.account_management.password_reset — cloudfront.protected_endpoints.flow_label.account_management.password_reset / 4947acfcf76d / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-e545d2d85bda83648ae399f5d5bf57fdfdcd9e81d0ed390bac2a66a7af9c1dd2)
- cloudfront.protected_endpoints.flow_label.account_management.password_reset

<a id="canonical-d7e490e20403f21d675bfa0c1d9dd59386ee8b4044b44a28eb89d5d520c1085a"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for password reset.

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
password_reset = {}
```

<a id="canonical-6d87159de0d39096c7424fce8690ba6b6a106c8e6cf604b653ad1f903b9d974e"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.account_management.password_reset / 4947acfcf76d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dae55fc676aedb1af5e63008e580cef44fdce6a0bf0efb8dbbddd2e4665ca2d4"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.account_management.password_reset / 4947acfcf76d / 4

- [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-e545d2d85bda83648ae399f5d5bf57fdfdcd9e81d0ed390bac2a66a7af9c1dd2)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a32fcdc83b78045ba5b8310e3070b1cf4c5774891ff589125051286cbe98956"></a>

## cloudfront.protected_endpoints.flow_label.authentication — cloudfront.protected_endpoints.flow_label.authentication / 2183e1bfd7c8 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- cloudfront.protected_endpoints.flow_label.authentication

<a id="canonical-981640247631decf7b8e60c415325a4720cdb8a78510eef204adaa3dc8378ee4"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Authentication Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("login",
    "login_mfa"),
  validators.ConflictingObjectAttributes("login",
    "login_partner"),
  validators.ConflictingObjectAttributes("login",
    "logout"),
  validators.ConflictingObjectAttributes("login",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_mfa",
    "login_partner"),
  validators.ConflictingObjectAttributes("login_mfa",
    "logout"),
  validators.ConflictingObjectAttributes("login_mfa",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_partner",
    "logout"),
  validators.ConflictingObjectAttributes("login_partner",
    "token_refresh"),
  validators.ConflictingObjectAttributes("logout",
    "token_refresh")}
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
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

<a id="canonical-8538b2e971b151119883bbe6eb162f16b0adaccf8eeb1d8d6a47500532317bce"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication / 2183e1bfd7c8 / 3

- [login](resources--protected_application--reference--group-003.md#canonical-94518014655701c746fb757f02e81334103395759ca084c44791df0b3e752c9b): complete subsection reference.

- [login_mfa](resources--protected_application--reference--group-003.md#canonical-4ad995812f702953f79f9f9325695a039e63a136c801b1c5eddcfc3b1892e40f): complete subsection reference.

- [login_partner](resources--protected_application--reference--group-003.md#canonical-c9f71eaf22667a9862c79e450121fcf1ed9b87926fe1c24412fa3abf7db8e84a): complete subsection reference.

- [logout](resources--protected_application--reference--group-003.md#canonical-8bb197cca165166c28fb0c59b22070ed88fb7331c7d3507ab00ba64023555f4b): complete subsection reference.

- [token_refresh](resources--protected_application--reference--group-003.md#canonical-29a01fb1582bb384ebabf02ada9070cf654ffaef8b2eb434635de8a671473fe2): complete subsection reference.

<a id="canonical-fb128925c4b5abdbfbdd5015f486c92dedb23747d42d724c13edcdb6ac214884"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication / 2183e1bfd7c8 / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-94518014655701c746fb757f02e81334103395759ca084c44791df0b3e752c9b)
- [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](resources--protected_application--reference--group-003.md#canonical-4ad995812f702953f79f9f9325695a039e63a136c801b1c5eddcfc3b1892e40f)
- [cloudfront.protected_endpoints.flow_label.authentication.login_partner](resources--protected_application--reference--group-003.md#canonical-c9f71eaf22667a9862c79e450121fcf1ed9b87926fe1c24412fa3abf7db8e84a)
- [cloudfront.protected_endpoints.flow_label.authentication.logout](resources--protected_application--reference--group-003.md#canonical-8bb197cca165166c28fb0c59b22070ed88fb7331c7d3507ab00ba64023555f4b)
- [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](resources--protected_application--reference--group-003.md#canonical-29a01fb1582bb384ebabf02ada9070cf654ffaef8b2eb434635de8a671473fe2)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-94518014655701c746fb757f02e81334103395759ca084c44791df0b3e752c9b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-328b92a653c2286e77888da001a8d4b20d08003a896a978a6fece02d34e4ff0f"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login — cloudfront.protected_endpoints.flow_label.authentication.login / 7ef4dbda239c / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- cloudfront.protected_endpoints.flow_label.authentication.login

<a id="canonical-90a0c7cdaa218792b38dc262258b372c0b748f22bd1c9aa0b75fab7d79833622"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_transaction_result",
    "transaction_result")}
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
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

Terraform syntax:

```terraform
login {
  # Configure direct properties listed below.
}
```

<a id="canonical-35268a286011eda7db0cc9ebce6ac8c00e1f359ab3561bc73ec24aa893e93a9e"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login / 7ef4dbda239c / 3

- [disable_transaction_result](resources--protected_application--reference--group-003.md#canonical-45f0ade184a872e799b6a82093c951a0f3c69f16f0a9e54057d016eab93ae7f6): complete subsection reference.

- [transaction_result](resources--protected_application--reference--group-003.md#canonical-ca998f2510059da71ac31f05cb684e6918bd44fa983312092ee6cd071114a556): complete subsection reference.

<a id="canonical-240cd65f76d133a38f2b8f584ba37e25244a9924decd9646d3b28efa4e2c2b26"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login / 7ef4dbda239c / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result](resources--protected_application--reference--group-003.md#canonical-45f0ade184a872e799b6a82093c951a0f3c69f16f0a9e54057d016eab93ae7f6)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-ca998f2510059da71ac31f05cb684e6918bd44fa983312092ee6cd071114a556)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-45f0ade184a872e799b6a82093c951a0f3c69f16f0a9e54057d016eab93ae7f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b188d0e4caf16c8c520c155b40862115ac2201cd9319fbdd39faa651a07b5550"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result — cloudfront.protected_endpoints.flow_label.authentication.login.disable_transacti / 55912dc6e6d8 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-94518014655701c746fb757f02e81334103395759ca084c44791df0b3e752c9b)
- cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-0bcf11dd58d6f1cda1b51c0f0f6e49e3932bb51aee0ab4eb83f440d66fcc6920"></a>

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
disable_transaction_result = {}
```

<a id="canonical-13624fdf1ef43cb120ca9488652f6aed8fde73d53938b718e14b6372c3d5841a"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login.disable_transacti / 55912dc6e6d8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f6d1ff4c3f4370f039476bbb843c70c0ac0208b1bcc0c8b59eaa19bb22745981"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login.disable_transacti / 55912dc6e6d8 / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-94518014655701c746fb757f02e81334103395759ca084c44791df0b3e752c9b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-ca998f2510059da71ac31f05cb684e6918bd44fa983312092ee6cd071114a556"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-041887668c9d1c4626f342af89d922e814e02bf77d008403f05b72ea15584097"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 2fcd65ad109d / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-94518014655701c746fb757f02e81334103395759ca084c44791df0b3e752c9b)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-76e450f93d4aed1e27042927ea04995f11db83eb979a54cfe2199fd93d96a1fd"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Upstream description:

Bot Defense Transaction ResultType.

Receipt-pinned upstream constraints:

```json
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
transaction_result {
  # Configure direct properties listed below.
}
```

<a id="canonical-96166707330a8166896c6aee0a72a631f3de84ddefdc06b3c3a45259e89292f7"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 2fcd65ad109d / 3

- [failure_conditions](resources--protected_application--reference--group-003.md#canonical-c80055e8550ed1fc8f17819581679623a0a4be3e51cc8a17c002ac1915b1c398): complete subsection reference.

- [success_conditions](resources--protected_application--reference--group-003.md#canonical-6e4a676f7ccca2ca0828a41a3b8d260e879810f6350f0dd8cac2f1d5f4fc1e42): complete subsection reference.

<a id="canonical-fc0fb5298483fad9d53c5bac89d7e1fb300e120408402c11aced33727eab5e3c"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 2fcd65ad109d / 4

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](resources--protected_application--reference--group-003.md#canonical-c80055e8550ed1fc8f17819581679623a0a4be3e51cc8a17c002ac1915b1c398)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions](resources--protected_application--reference--group-003.md#canonical-6e4a676f7ccca2ca0828a41a3b8d260e879810f6350f0dd8cac2f1d5f4fc1e42)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-94518014655701c746fb757f02e81334103395759ca084c44791df0b3e752c9b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-c80055e8550ed1fc8f17819581679623a0a4be3e51cc8a17c002ac1915b1c398"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-837cfa121075e547e53127bb7d37ff2b4cfea5f51693d7f7a7f54c02ccd50aea"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 22735b6ef1bd / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-94518014655701c746fb757f02e81334103395759ca084c44791df0b3e752c9b)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-ca998f2510059da71ac31f05cb684e6918bd44fa983312092ee6cd071114a556)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-ba22f3d76edf5a8ab172867e782c71a67911ce13aee9b967fe2620c4397303b2"></a>

Type: `"object"`. list nested block, Optional.

Failure Conditions. Failure Conditions.

Upstream description:

Failure Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
failure_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-37cd144c203914c39970f7a37e9eb7b2c9b9557461cee43fa9821a72933a6b7e"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 22735b6ef1bd / 3

<a id="canonical-83562465fb7330083906df91ce4f3a36d4d271465876f2771027b4306ba8080f"></a>

<a id="canonical-98ae7e81fe928c4589fe1e15c0c280c315116305ef9b6b6aad56dcd09affa45e"></a>

## name property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 22735b6ef1bd / 4

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-4520995a8bec4ab356e3381c938b712acf7fd80e1070753d9e20869caf4381a0"></a>

<a id="canonical-35a8243e6b4bae9e062438a4cf1012582e72ce30aa1606e255033d894dc9cdec"></a>

## regex_values property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 22735b6ef1bd / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-5fde6390d722bd3660a4c58de7cf4d29ba000f0f5dac78dd12dc1b70d1699162"></a>

<a id="canonical-0c24071fe71030bc8cecc307727ef590de5cf93ced6427005f9dcd4d2e8c928d"></a>

## status property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 22735b6ef1bd / 6

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b424a8cbed879ae24fedb0c9dc887fe6080eebf0ac1de55231eb8d1d189aeaf5"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 22735b6ef1bd / 7

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-ca998f2510059da71ac31f05cb684e6918bd44fa983312092ee6cd071114a556)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-6e4a676f7ccca2ca0828a41a3b8d260e879810f6350f0dd8cac2f1d5f4fc1e42"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2745b25ccbf8b772f81f893176b99b17ce489267fc467a07299f5ee8ee203ba7"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 9a9c259399f5 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-94518014655701c746fb757f02e81334103395759ca084c44791df0b3e752c9b)
- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-ca998f2510059da71ac31f05cb684e6918bd44fa983312092ee6cd071114a556)
- cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-b3bf6d8ffeadb6b039701c31ea2d83531c33c518ccfbf403d12c7f2f371ade04"></a>

Type: `"object"`. list nested block, Optional.

Success Conditions. Success Conditions.

Upstream description:

Success Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
success_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-f89761427465623768029f0b1bfa99f8cca635dc39165aa4ee951ea8508c40ab"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 9a9c259399f5 / 3

<a id="canonical-daaeb317717971fccec44d21f4227b9823b6ead94375c085920375b266e396be"></a>

<a id="canonical-16e9e2dd69c46fec84ec4c733ed4daa62c6eb054881faa4264923a344c6e167c"></a>

## name property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 9a9c259399f5 / 4

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-aaa8013f464a1b5b642be3ca1c5cf9c5f94187bfb762dd0e12e3ef1dd5eea65f"></a>

<a id="canonical-95832afaf97676c00dcb38fcd367dbab90a0dc258aff507dcfdb13caad432efa"></a>

## regex_values property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 9a9c259399f5 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3b304cd79a56bc1bf19f2647c26e763bd2c8ff8af69811fe8384a49720a74ecd"></a>

<a id="canonical-333500bae1cf4566d5d20dab059d4b2d2a84184afc2ba349daa4f491ca5f337f"></a>

## status property — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 9a9c259399f5 / 6

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fdfd9377486d335c5c455ba7a16d1c873ce31be6c3a30fbfaa85b71714bdbfc6"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login.transaction_resul / 9a9c259399f5 / 7

- [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-ca998f2510059da71ac31f05cb684e6918bd44fa983312092ee6cd071114a556)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-4ad995812f702953f79f9f9325695a039e63a136c801b1c5eddcfc3b1892e40f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-467ac4942ec797fa88d1161535ab09086aa25b41f7db73af1e6db4f10d545ad4"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login_mfa — cloudfront.protected_endpoints.flow_label.authentication.login_mfa / 6dc98534be3b / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- cloudfront.protected_endpoints.flow_label.authentication.login_mfa

<a id="canonical-c3c229936db01d04708e709d4bd88d4ad13708cc71ff8bba7f926cb5c5b79200"></a>

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
login_mfa = {}
```

<a id="canonical-6466304a088d44c883442982ae53157c930068d187f9334f0c4aaae100eef9b1"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login_mfa / 6dc98534be3b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c7e258106863f7e182e3ca268f391a3b53e10e256c686757978dc1f4dd3deee1"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login_mfa / 6dc98534be3b / 4

- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-c9f71eaf22667a9862c79e450121fcf1ed9b87926fe1c24412fa3abf7db8e84a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d6bba59f42e4974a284092208b62c61927ebe6861fc1e42021889ee3da8ce8e"></a>

## cloudfront.protected_endpoints.flow_label.authentication.login_partner — cloudfront.protected_endpoints.flow_label.authentication.login_partner / 9df70a3967bd / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- cloudfront.protected_endpoints.flow_label.authentication.login_partner

<a id="canonical-5ecde7bfc6c6e927eee1665462e36ac6ecc0ebc189945b3f5c6e577e7a69afcf"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for login partner.

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
login_partner = {}
```

<a id="canonical-fd2ae2354b800a05d9116294a2a9c4caa7b5a789cc033e0b9478168bd7716312"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.login_partner / 9df70a3967bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-896e4a6efa54375c4f36a835fb4942bcbc733df7849e5005f79e7561adb2e875"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.login_partner / 9df70a3967bd / 4

- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-8bb197cca165166c28fb0c59b22070ed88fb7331c7d3507ab00ba64023555f4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09a97bb7b34a166dacbf8a13b7a1ef5fb172ea0a74cbe42ded6db126b7a76a2c"></a>

## cloudfront.protected_endpoints.flow_label.authentication.logout — cloudfront.protected_endpoints.flow_label.authentication.logout / ae8f85c029ad / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- cloudfront.protected_endpoints.flow_label.authentication.logout

<a id="canonical-e4ebf30528b3cd72350e16c6ad342db14adc908a317c941f8e3303b866827ecd"></a>

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
logout = {}
```

<a id="canonical-16b04c2e693f07b4ae74613f5c90f3c7a943c230e2a4ee7556e0168b2e018072"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.logout / ae8f85c029ad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bfeefac69edf05ea649f0908d4f2d33b4e4ebfb8a807754f2cb3a0dd05ff89e7"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.logout / ae8f85c029ad / 4

- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-29a01fb1582bb384ebabf02ada9070cf654ffaef8b2eb434635de8a671473fe2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3b63db88bc9d886d594a8b657fba8614a177f0ed68acc282c1fa35f1da41be3"></a>

## cloudfront.protected_endpoints.flow_label.authentication.token_refresh — cloudfront.protected_endpoints.flow_label.authentication.token_refresh / 76c663a344d2 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- cloudfront.protected_endpoints.flow_label.authentication.token_refresh

<a id="canonical-1f21eb274eb0d8a92f7cb73696ec56b8697cd96c195cd4cab8519c2cfe9b870d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for token refresh.

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
token_refresh = {}
```

<a id="canonical-25d7cece0f0ca5f96cb24f22e2c3c9ffbb9569fd2cf5bbe05ce2f47ed6ed61ba"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.authentication.token_refresh / 76c663a344d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39ed8bb562c72491a68fa2326276f05036ce10d663fa511435425edb9ac2cf19"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.authentication.token_refresh / 76c663a344d2 / 4

- [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-3b56617dafcbc699762b5856c89c79b6b5f86d4968c2af4e194b83e4b42a4214)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-692ea01fa276546cc2074827be4f4bf9519afd28943b52af84b6e8f3f53ce986"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2899acee21cff02d0a94c0dbf88e6f2674ce3e1f0b4cf016a83f69d8e015add"></a>

## cloudfront.protected_endpoints.flow_label.financial_services — cloudfront.protected_endpoints.flow_label.financial_services / 575e34812a8b / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- cloudfront.protected_endpoints.flow_label.financial_services

<a id="canonical-f52da908064dc74796cc44d32b471d572c9e5f56fa2eed181db323598f14bb5f"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Financial Services Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("apply",
    "money_transfer")}
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
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

Terraform syntax:

```terraform
financial_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-bc3feeeaf2a6b351a33e4ed7f3b3578bf1c6cfd682c227aaa52c55f126b7873c"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.financial_services / 575e34812a8b / 3

- [apply](resources--protected_application--reference--group-003.md#canonical-ef711746acb7b633567812343da78cde1f319fea931ee91e80edde3c1c99ed6d): complete subsection reference.

- [money_transfer](resources--protected_application--reference--group-003.md#canonical-9b7c83b7014a3b8db1fd34f1afaa274131c5cbe56e43ac16fc432e0af58f1c7f): complete subsection reference.

<a id="canonical-65cf5e65abddb7b7282dab349d1273be572ced110d40a9009b84dbf67225e394"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.financial_services / 575e34812a8b / 4

- [cloudfront.protected_endpoints.flow_label.financial_services.apply](resources--protected_application--reference--group-003.md#canonical-ef711746acb7b633567812343da78cde1f319fea931ee91e80edde3c1c99ed6d)
- [cloudfront.protected_endpoints.flow_label.financial_services.money_transfer](resources--protected_application--reference--group-003.md#canonical-9b7c83b7014a3b8db1fd34f1afaa274131c5cbe56e43ac16fc432e0af58f1c7f)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-ef711746acb7b633567812343da78cde1f319fea931ee91e80edde3c1c99ed6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1999504144c56bf01b3577fef2b83c76d31f800b405be426d538c2a9e0b744ef"></a>

## cloudfront.protected_endpoints.flow_label.financial_services.apply — cloudfront.protected_endpoints.flow_label.financial_services.apply / 2fc31627c663 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-692ea01fa276546cc2074827be4f4bf9519afd28943b52af84b6e8f3f53ce986)
- cloudfront.protected_endpoints.flow_label.financial_services.apply

<a id="canonical-548f04d5e0eb1c07d1620d23d357ace6f4fdfc453fc93cf58e62e341cc940cfe"></a>

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
apply = {}
```

<a id="canonical-3043798db3a788fdb5e20b8ee4e5844d7d3dc32e8333f2e959e7701af9557b60"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.financial_services.apply / 2fc31627c663 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-693e17af0957397c45c532539d98f802707dc19b317a25689fa36a8f54743bda"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.financial_services.apply / 2fc31627c663 / 4

- [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-692ea01fa276546cc2074827be4f4bf9519afd28943b52af84b6e8f3f53ce986)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-9b7c83b7014a3b8db1fd34f1afaa274131c5cbe56e43ac16fc432e0af58f1c7f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5541f3157bba8697315a1888f20a0be76b2a8d595f72535ddcf14235dc873266"></a>

## cloudfront.protected_endpoints.flow_label.financial_services.money_transfer — cloudfront.protected_endpoints.flow_label.financial_services.money_transfer / 9d7552b28f11 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-692ea01fa276546cc2074827be4f4bf9519afd28943b52af84b6e8f3f53ce986)
- cloudfront.protected_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-06f3edbc9a417c4daa83744ea0efd9208ecca9a9653f13ea6f9a4c7752acc660"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for money transfer.

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
money_transfer = {}
```

<a id="canonical-81287e6c9cb1f5aa9ccc2c2d9134152730a1959272db4589590dc5914465e383"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.financial_services.money_transfer / 9d7552b28f11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b4d53f911743fe39fb35b0a4b784fe9bdfb9703da9ec2278a5464acd1fad87d"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.financial_services.money_transfer / 9d7552b28f11 / 4

- [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-692ea01fa276546cc2074827be4f4bf9519afd28943b52af84b6e8f3f53ce986)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-6a1b2bf4c28b2364ffd480cad552a4190aa23617e63f7262544c676ea9e8a548"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3265b7f3068a288ba9cd6aa1b22b62dd827f3f619bf3bd656b6663cfc43d9b3"></a>

## cloudfront.protected_endpoints.flow_label.flight — cloudfront.protected_endpoints.flow_label.flight / d46a94a10e37 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- cloudfront.protected_endpoints.flow_label.flight

<a id="canonical-f1e5d2f732f171ccf4b9432accdeed88783aeecb7e8ce55ccc2d29f8623e1fcb"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Upstream description:

Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

Terraform syntax:

```terraform
flight {
  # Configure direct properties listed below.
}
```

<a id="canonical-442d627bcd2a4442df6b586b10aa7caa1d3923e11aef490d4a98c62303172f5e"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.flight / d46a94a10e37 / 3

- [checkin](resources--protected_application--reference--group-003.md#canonical-1760c214053dd77f846614eae652f59d1c12364a08cf8bb59ba5edce4c48df67): complete subsection reference.

<a id="canonical-487e5cefdc4431f53863c2203d9f066aa2176d45857dccc12d48390603e69fd1"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.flight / d46a94a10e37 / 4

- [cloudfront.protected_endpoints.flow_label.flight.checkin](resources--protected_application--reference--group-003.md#canonical-1760c214053dd77f846614eae652f59d1c12364a08cf8bb59ba5edce4c48df67)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-1760c214053dd77f846614eae652f59d1c12364a08cf8bb59ba5edce4c48df67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c725a792d7cf1aed105f8b20c359b7e1f632a4c762c975992f503a18a389ff7"></a>

## cloudfront.protected_endpoints.flow_label.flight.checkin — cloudfront.protected_endpoints.flow_label.flight.checkin / 5c9a6ae305d3 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.flight](resources--protected_application--reference--group-003.md#canonical-6a1b2bf4c28b2364ffd480cad552a4190aa23617e63f7262544c676ea9e8a548)
- cloudfront.protected_endpoints.flow_label.flight.checkin

<a id="canonical-993837945eb512d856b6536a553919c657f448769b4e93eab7737a0b337be032"></a>

Type: `"object"`. single nested block, Optional.

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
checkin {}
```

<a id="canonical-6f5b18035141f2811e48db63f81aa567e7ddfef2a8363f2a5b5b80db99684e00"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.flight.checkin / 5c9a6ae305d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8ae84c45b701a97724963e71e669aa12158fbb9ee7408fbca3d98947bb34f0d9"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.flight.checkin / 5c9a6ae305d3 / 4

- [cloudfront.protected_endpoints.flow_label.flight](resources--protected_application--reference--group-003.md#canonical-6a1b2bf4c28b2364ffd480cad552a4190aa23617e63f7262544c676ea9e8a548)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-e54910f6ce94de74a592ed705d1ddf46731acb511631278f5d75ea1169885bc7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab61256359b1ba565303ecf6c7d3faabbc10fce802944a5383957d2516dd79d6"></a>

## cloudfront.protected_endpoints.flow_label.profile_management — cloudfront.protected_endpoints.flow_label.profile_management / 11253c2bb3a1 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- cloudfront.protected_endpoints.flow_label.profile_management

<a id="canonical-e186dcf62e2fb13c8c6c4f40d55c6ed52ccb66a956f989798eba1d63c5a0c50c"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Profile Management Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "update"),
  validators.ConflictingObjectAttributes("create",
    "view"),
  validators.ConflictingObjectAttributes("update",
    "view")}
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
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

Terraform syntax:

```terraform
profile_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-f1e25cf4d59d69eaf195185ef72e5d514b192e662c41fac5a2b6f3fbb126e52a"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.profile_management / 11253c2bb3a1 / 3

- [create](resources--protected_application--reference--group-003.md#canonical-90ed8cd16dce5d22c60e915ffb451f71c952ee064b036f3478c9fd91e639f4ec): complete subsection reference.

- [update](resources--protected_application--reference--group-003.md#canonical-216d4c642461052d4f196c37ff2d0ab45780b2391220beedb00cf8190ebaa48b): complete subsection reference.

- [view](resources--protected_application--reference--group-003.md#canonical-6fa907c5022bced7d1f112a3120d5a6ef678e73d053068a0bd3310b77e09e6ca): complete subsection reference.

<a id="canonical-80d22cf8b62d1a752eccd14c311e849ba3c7337fd1f8652b2514a4de44099897"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.profile_management / 11253c2bb3a1 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management.create](resources--protected_application--reference--group-003.md#canonical-90ed8cd16dce5d22c60e915ffb451f71c952ee064b036f3478c9fd91e639f4ec)
- [cloudfront.protected_endpoints.flow_label.profile_management.update](resources--protected_application--reference--group-003.md#canonical-216d4c642461052d4f196c37ff2d0ab45780b2391220beedb00cf8190ebaa48b)
- [cloudfront.protected_endpoints.flow_label.profile_management.view](resources--protected_application--reference--group-003.md#canonical-6fa907c5022bced7d1f112a3120d5a6ef678e73d053068a0bd3310b77e09e6ca)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-90ed8cd16dce5d22c60e915ffb451f71c952ee064b036f3478c9fd91e639f4ec"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30ed9f63d3bf04f463371fb056762bbcd041ed7cd6e4d87a8e030b2a5d15ead7"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.create — cloudfront.protected_endpoints.flow_label.profile_management.create / 4e50885da49c / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-e54910f6ce94de74a592ed705d1ddf46731acb511631278f5d75ea1169885bc7)
- cloudfront.protected_endpoints.flow_label.profile_management.create

<a id="canonical-020f739e4d703d2859d1baf7d88833e6d9c2241fa183fd6e46936833dd99f668"></a>

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
create = {}
```

<a id="canonical-d2c4d4bd428702364439d3369b07bc6ea48200ac0c6692e191b41ca119a9ce9c"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.profile_management.create / 4e50885da49c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fcc06528e2a90deb473285bf808debdc0ed714ff8b66b9719a76c511e36ef904"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.profile_management.create / 4e50885da49c / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-e54910f6ce94de74a592ed705d1ddf46731acb511631278f5d75ea1169885bc7)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-216d4c642461052d4f196c37ff2d0ab45780b2391220beedb00cf8190ebaa48b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce6fffffeca58f220cc0c5542c1fb3b8458e16a565f8ee2a9967752b11d20e3b"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.update — cloudfront.protected_endpoints.flow_label.profile_management.update / a201a4e399ee / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-e54910f6ce94de74a592ed705d1ddf46731acb511631278f5d75ea1169885bc7)
- cloudfront.protected_endpoints.flow_label.profile_management.update

<a id="canonical-6a6aa4101a346c9d54234a91ae2ae2cafcbcc17a4fa030fd1ed11a910895291c"></a>

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
update = {}
```

<a id="canonical-9d3779453773ed977f991e975eef72cc7ca2cb511581d82db73206cd610ab7b2"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.profile_management.update / a201a4e399ee / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-49a628622a005f6a1d0c391ec4795b1b06b13c5080cfb0f434cd787941173875"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.profile_management.update / a201a4e399ee / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-e54910f6ce94de74a592ed705d1ddf46731acb511631278f5d75ea1169885bc7)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-6fa907c5022bced7d1f112a3120d5a6ef678e73d053068a0bd3310b77e09e6ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bbfd99502ee955613afa46cd5fb5aa076fffa24b0028b24aa285b49db5aca795"></a>

## cloudfront.protected_endpoints.flow_label.profile_management.view — cloudfront.protected_endpoints.flow_label.profile_management.view / c531b98e0715 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-e54910f6ce94de74a592ed705d1ddf46731acb511631278f5d75ea1169885bc7)
- cloudfront.protected_endpoints.flow_label.profile_management.view

<a id="canonical-568e79bd9d85ae1a26c6036211e49a4dbc137e523252874c4de28fab7567a9df"></a>

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
view = {}
```

<a id="canonical-d645210a9a92a862b6e48184b4d258594bd8ee09fa5118d79fddbf09114873cd"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.profile_management.view / c531b98e0715 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a82f2b2a5669375a4e63320cb6c3445ac46c474838e9fbe0f994e60d0c93acfd"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.profile_management.view / c531b98e0715 / 4

- [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-e54910f6ce94de74a592ed705d1ddf46731acb511631278f5d75ea1169885bc7)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-69603b4df4c676280ea3b16f946a0121f7a8985bc8492753dbf140db99f2b88f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41dab73fb952876b7ec8116ffb968998f8c889154ece7eb70971df1d671ad61b"></a>

## cloudfront.protected_endpoints.flow_label.search — cloudfront.protected_endpoints.flow_label.search / 12686f4f7b12 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- cloudfront.protected_endpoints.flow_label.search

<a id="canonical-4f6801c3605c073434c3c88ce315256de1e15b32842c675f50e268fb9ae8ad4b"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Upstream description:

Bot Defense Flow Label Search Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("flight_search",
    "product_search"),
  validators.ConflictingObjectAttributes("flight_search",
    "reservation_search"),
  validators.ConflictingObjectAttributes("flight_search",
    "room_search"),
  validators.ConflictingObjectAttributes("product_search",
    "reservation_search"),
  validators.ConflictingObjectAttributes("product_search",
    "room_search"),
  validators.ConflictingObjectAttributes("reservation_search",
    "room_search")}
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
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```

Terraform syntax:

```terraform
search {
  # Configure direct properties listed below.
}
```

<a id="canonical-31d5ef23453402295e77c6271f6e8c7d5990f4cde6e3dc9c110538911685e104"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.search / 12686f4f7b12 / 3

- [flight_search](resources--protected_application--reference--group-003.md#canonical-eb30df84a46114b06726e99fc8e2df3587ca1452e5ceb5ef2d90573dbeb2c9a5): complete subsection reference.

- [product_search](resources--protected_application--reference--group-003.md#canonical-21aefefdbe857f0a950826c9fecaafe650c0e70727a15afdadbdad8649b82148): complete subsection reference.

- [reservation_search](resources--protected_application--reference--group-003.md#canonical-d79446f4583f80a45ccef755dac3c350202f8e2ff218972640faf947aa6ef1a2): complete subsection reference.

- [room_search](resources--protected_application--reference--group-003.md#canonical-9f735aba9a6c6986bf305d8aff5474cc4989bdd4d744165301d3a14bb3313262): complete subsection reference.

<a id="canonical-673fd12b551c563ef2794e494dd818278c446e1cd0a62d29ea82cd3f43fd51a4"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.search / 12686f4f7b12 / 4

- [cloudfront.protected_endpoints.flow_label.search.flight_search](resources--protected_application--reference--group-003.md#canonical-eb30df84a46114b06726e99fc8e2df3587ca1452e5ceb5ef2d90573dbeb2c9a5)
- [cloudfront.protected_endpoints.flow_label.search.product_search](resources--protected_application--reference--group-003.md#canonical-21aefefdbe857f0a950826c9fecaafe650c0e70727a15afdadbdad8649b82148)
- [cloudfront.protected_endpoints.flow_label.search.reservation_search](resources--protected_application--reference--group-003.md#canonical-d79446f4583f80a45ccef755dac3c350202f8e2ff218972640faf947aa6ef1a2)
- [cloudfront.protected_endpoints.flow_label.search.room_search](resources--protected_application--reference--group-003.md#canonical-9f735aba9a6c6986bf305d8aff5474cc4989bdd4d744165301d3a14bb3313262)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-eb30df84a46114b06726e99fc8e2df3587ca1452e5ceb5ef2d90573dbeb2c9a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a86d0bf713cc5478c426df40fe41beb3b283969e98419d004ffdbedbde28307"></a>

## cloudfront.protected_endpoints.flow_label.search.flight_search — cloudfront.protected_endpoints.flow_label.search.flight_search / a7cdfc1dd4dd / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-69603b4df4c676280ea3b16f946a0121f7a8985bc8492753dbf140db99f2b88f)
- cloudfront.protected_endpoints.flow_label.search.flight_search

<a id="canonical-8e6394182be2d3201328935e3a97bd9bd411531b8f7d1afb7b9d1a8f89aa3f12"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for flight search.

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
flight_search = {}
```

<a id="canonical-54773d8fbd214cc78ab05b0e85d22b59dd6c1bbf9fd7090b5c2385fbb827ce94"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.search.flight_search / a7cdfc1dd4dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6eb36ab351caf67ba3f578166a291f4e991f4ac3fd7ec9931fa0b605e9619efb"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.search.flight_search / a7cdfc1dd4dd / 4

- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-69603b4df4c676280ea3b16f946a0121f7a8985bc8492753dbf140db99f2b88f)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-21aefefdbe857f0a950826c9fecaafe650c0e70727a15afdadbdad8649b82148"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9361d4f254659231b3dd5ceba9e526cde1e1ca7e6ec1181c2ffe79c8c72ec6b"></a>

## cloudfront.protected_endpoints.flow_label.search.product_search — cloudfront.protected_endpoints.flow_label.search.product_search / ee6abc8ece18 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-69603b4df4c676280ea3b16f946a0121f7a8985bc8492753dbf140db99f2b88f)
- cloudfront.protected_endpoints.flow_label.search.product_search

<a id="canonical-d98943c38434994273e9761c98bc47e789231c6147282989d59cc7d26f60ec19"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for product search.

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
product_search = {}
```

<a id="canonical-c30d867baa0eceaae6b5856f326ba4fcb551869fb76069275022b7eb76fee643"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.search.product_search / ee6abc8ece18 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-64de807dcd4c7abc45b9f194f4494c8ab79484329eef97101fb6a0c8c805bd2c"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.search.product_search / ee6abc8ece18 / 4

- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-69603b4df4c676280ea3b16f946a0121f7a8985bc8492753dbf140db99f2b88f)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-d79446f4583f80a45ccef755dac3c350202f8e2ff218972640faf947aa6ef1a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edfb49855b219324d9bf11e6e5ad46820ff728733f868763dd77d86b5da371cf"></a>

## cloudfront.protected_endpoints.flow_label.search.reservation_search — cloudfront.protected_endpoints.flow_label.search.reservation_search / b355916fcd9d / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-69603b4df4c676280ea3b16f946a0121f7a8985bc8492753dbf140db99f2b88f)
- cloudfront.protected_endpoints.flow_label.search.reservation_search

<a id="canonical-783dcc97359da602704553e2534861ff1b870be1f7589f6d60601906a2217988"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reservation search.

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
reservation_search = {}
```

<a id="canonical-66706c535252aea13c3b23249e1da019b37bf07cc1d81d1efbb0bbc77b03d1cb"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.search.reservation_search / b355916fcd9d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2ed54ff6961c43cc29456dd341be4e9f059f83b87fff6295b84638abd89e3902"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.search.reservation_search / b355916fcd9d / 4

- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-69603b4df4c676280ea3b16f946a0121f7a8985bc8492753dbf140db99f2b88f)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-9f735aba9a6c6986bf305d8aff5474cc4989bdd4d744165301d3a14bb3313262"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f046348c9d407762febd3831b6b69b4c3a5d7f87c5215262e1bbea0d58b53bb"></a>

## cloudfront.protected_endpoints.flow_label.search.room_search — cloudfront.protected_endpoints.flow_label.search.room_search / 1b4bfa052c36 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-69603b4df4c676280ea3b16f946a0121f7a8985bc8492753dbf140db99f2b88f)
- cloudfront.protected_endpoints.flow_label.search.room_search

<a id="canonical-ba0d7067ffddfec03099a3977cbb7d777d31db6a4981549974a5cfc717a87f08"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for room search.

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
room_search = {}
```

<a id="canonical-bd746bfb537cb8d68224af5aa46e3886444cc996fe78f659678bdcf567de47e6"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.search.room_search / 1b4bfa052c36 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ef50b9dd7e137c33dab6ce7e197115e9c5c9ac49cfbc2c14ac010811959db6f0"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.search.room_search / 1b4bfa052c36 / 4

- [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-69603b4df4c676280ea3b16f946a0121f7a8985bc8492753dbf140db99f2b88f)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94b584daa9a48a6d87704909337977993cb68e4150d1542df8d3bbd282dc879a"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards — cloudfront.protected_endpoints.flow_label.shopping_gift_cards / c115f09524e2 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards

<a id="canonical-7ec8beac847eb2676a0ef58023d81a67d916def0edf6196c3de59c04ee603a62"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "gift_card_validation"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_add_to_cart"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_order"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_add_to_cart"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_order"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_promo_code_validation",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_promo_code_validation",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_purchase_gift_card",
    "shop_update_quantity")}
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
  "x-ves-oneof-field-label_choice": "[\"gift_card_make_purchase_with_gift_card\",\"gift_card_validation\",\"shop_add_to_cart\",\"shop_checkout\",\"shop_choose_seat\",\"shop_enter_drawing_submission\",\"shop_make_payment\",\"shop_order\",\"shop_price_inquiry\",\"shop_promo_code_validation\",\"shop_purchase_gift_card\",\"shop_update_quantity\"]"
}
```

Terraform syntax:

```terraform
shopping_gift_cards {
  # Configure direct properties listed below.
}
```

<a id="canonical-8990bf287351d73695e75d4e969021231cd888ede09503e3a4fbaf4a342c4cc4"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards / c115f09524e2 / 3

- [gift_card_make_purchase_with_gift_card](resources--protected_application--reference--group-003.md#canonical-d8c3ce2aa064d20b8d567f010ea0716456c34f30e875123625b4a864c49921be): complete subsection reference.

- [gift_card_validation](resources--protected_application--reference--group-003.md#canonical-9ff6f85eb0ce7c09eaab890c371a5c9b32f368516f95715b9faf07355a7aef3c): complete subsection reference.

- [shop_add_to_cart](resources--protected_application--reference--group-003.md#canonical-181d794963f75a0e22aa215c0a8d04c96b06ac457a781bcfbd2c1294c87a1ef5): complete subsection reference.

- [shop_checkout](resources--protected_application--reference--group-003.md#canonical-5b40bd5f5ffce131faa6a23f3f4e014c0e2be6ff62d39db72ac09614352ed795): complete subsection reference.

- [shop_choose_seat](resources--protected_application--reference--group-003.md#canonical-66ffb2d651467aeac32a5ed83a3f9484c60c6e0b7dd31fe4dee0217f10fc56d7): complete subsection reference.

- [shop_enter_drawing_submission](resources--protected_application--reference--group-003.md#canonical-5a86f22b1b6d2c428e0f6f678038cf1a7a409ff3f7300414c8a3f374bb8f16fb): complete subsection reference.

- [shop_make_payment](resources--protected_application--reference--group-003.md#canonical-464b8aad2c00c73e28d52296a31d4eff80e0c7aeef1d9b9598741c809deffd70): complete subsection reference.

- [shop_order](resources--protected_application--reference--group-003.md#canonical-059d5a47901ebd8bc81b3eeca83d20427d4ca915690483a06f5312eed7f3a250): complete subsection reference.

- [shop_price_inquiry](resources--protected_application--reference--group-003.md#canonical-54c0b58f4744fd6075e89f3a8683b06bb0036358342b1ee9f5af5c0a46d55682): complete subsection reference.

- [shop_promo_code_validation](resources--protected_application--reference--group-003.md#canonical-5c41470176f87c10253050a0a24a6b811d4b820f84d6c9cad25df5c44b990cd5): complete subsection reference.

- [shop_purchase_gift_card](resources--protected_application--reference--group-003.md#canonical-0b53f92646212bd675d4b3af889c1fc38c1a7371e30a331f41d44abc32512ca1): complete subsection reference.

- [shop_update_quantity](resources--protected_application--reference--group-003.md#canonical-b65b6c703796164baafa6a6c9a203d279a3d179788e0960ea1d98c0dc5ce75f3): complete subsection reference.

<a id="canonical-164647b1f163a7ed211168fcacfc937c1f402d0a3ddbbd9a133bbf80de5144f2"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards / c115f09524e2 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](resources--protected_application--reference--group-003.md#canonical-d8c3ce2aa064d20b8d567f010ea0716456c34f30e875123625b4a864c49921be)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation](resources--protected_application--reference--group-003.md#canonical-9ff6f85eb0ce7c09eaab890c371a5c9b32f368516f95715b9faf07355a7aef3c)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](resources--protected_application--reference--group-003.md#canonical-181d794963f75a0e22aa215c0a8d04c96b06ac457a781bcfbd2c1294c87a1ef5)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout](resources--protected_application--reference--group-003.md#canonical-5b40bd5f5ffce131faa6a23f3f4e014c0e2be6ff62d39db72ac09614352ed795)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](resources--protected_application--reference--group-003.md#canonical-66ffb2d651467aeac32a5ed83a3f9484c60c6e0b7dd31fe4dee0217f10fc56d7)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](resources--protected_application--reference--group-003.md#canonical-5a86f22b1b6d2c428e0f6f678038cf1a7a409ff3f7300414c8a3f374bb8f16fb)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment](resources--protected_application--reference--group-003.md#canonical-464b8aad2c00c73e28d52296a31d4eff80e0c7aeef1d9b9598741c809deffd70)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order](resources--protected_application--reference--group-003.md#canonical-059d5a47901ebd8bc81b3eeca83d20427d4ca915690483a06f5312eed7f3a250)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](resources--protected_application--reference--group-003.md#canonical-54c0b58f4744fd6075e89f3a8683b06bb0036358342b1ee9f5af5c0a46d55682)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](resources--protected_application--reference--group-003.md#canonical-5c41470176f87c10253050a0a24a6b811d4b820f84d6c9cad25df5c44b990cd5)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](resources--protected_application--reference--group-003.md#canonical-0b53f92646212bd675d4b3af889c1fc38c1a7371e30a331f41d44abc32512ca1)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](resources--protected_application--reference--group-003.md#canonical-b65b6c703796164baafa6a6c9a203d279a3d179788e0960ea1d98c0dc5ce75f3)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-d8c3ce2aa064d20b8d567f010ea0716456c34f30e875123625b4a864c49921be"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bb52b37a723a678123ca2b3688ed35dacc2083eaf3b6ca617684fb9e5f41336"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_pur / e40fdb1f898b / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-a93be5c799e2c476be09cea734be5be424c6f18f0e8985a1446ebaad16c48a6f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for gift card make purchase with gift card.

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
gift_card_make_purchase_with_gift_card = {}
```

<a id="canonical-ee44e4df15ed4fa5659193de70cd753ac18271cec4e2fbf1526b650994dcb615"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_pur / e40fdb1f898b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fdf4bf53899c818ca14a485195691ff6a3b0e29972703a05846dd957af0e4186"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_pur / e40fdb1f898b / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-9ff6f85eb0ce7c09eaab890c371a5c9b32f368516f95715b9faf07355a7aef3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1477d203a389ed1b79c4bddeaec7cf504a58e7bd1fd28469cf9dddefb43bbad6"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validati / 649186e6d734 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-287a7913cbe6a94d19ebdd56d57a3117b725a7bbf47486635d9b1efdc3a7fdcb"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for gift card validation.

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
gift_card_validation = {}
```

<a id="canonical-4c565ae501b3152d876d4e5af35e939e12ef1df5cf3217c343fd8b78aac96c97"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validati / 649186e6d734 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-467ffa42699294b6d35304d75a95c06fcc158f6719279cde521642a2a9cef7a3"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validati / 649186e6d734 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-181d794963f75a0e22aa215c0a8d04c96b06ac457a781bcfbd2c1294c87a1ef5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19c4527136c1e2eb71b02614d357e69f30745bf919a873a6ebf0b2c8758418e1"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart / c51c90e48bec / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-4636ce24ac98300c174284ab1c663d23d05b83fbb32e95ab49f889d36654c2bd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop add to cart.

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
shop_add_to_cart = {}
```

<a id="canonical-9fc909dc092a0f50386ca8733a5c58ab7da0968f47398fd43b32885045351318"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart / c51c90e48bec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-415c04c5bb0104b0cd334ef4b45daaf04b969bfea7c60b74b18e846de1d74651"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart / c51c90e48bec / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-5b40bd5f5ffce131faa6a23f3f4e014c0e2be6ff62d39db72ac09614352ed795"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d8e2bce8e42b2c9f0c775a9323b2c46577dfcbb69cc05e513bd364a8e17033f"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout / 23844f9711f9 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-a2787939120d9f84b59531b7b7651aeabaadd37f0d621c2fb4da184ecd8d7e05"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop checkout.

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
shop_checkout = {}
```

<a id="canonical-2b7df9115c1b256f8a7ebdfea48be8b3b01d4b2483261f76970a49608af16cf5"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout / 23844f9711f9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2fbc48c9c98c157dd68f9f320354af26625db9b255af5a14a520046d261efd08"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout / 23844f9711f9 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-66ffb2d651467aeac32a5ed83a3f9484c60c6e0b7dd31fe4dee0217f10fc56d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-121c8d286e2d9afc2f35e200a7f057b151d3bdedfff19efd7a8325963942fe9d"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat / b6fd8df4732c / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat

<a id="canonical-d3d7bff1125d764be0f06fec4081752f5be017db3b5fbfc6e80785e6744f5f8f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop choose seat.

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
shop_choose_seat = {}
```

<a id="canonical-f0fe6d4594af25d1702e8d8589fa1556ad42973519f2351bf36bc2fdee435cc4"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat / b6fd8df4732c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1b239f6fd81a3dad2cb2ec334385b186ab62f4f7c481a728af499e443f24b30b"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat / b6fd8df4732c / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-5a86f22b1b6d2c428e0f6f678038cf1a7a409ff3f7300414c8a3f374bb8f16fb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36d0257ed8703ba668bc05ecd86bce9a6c57c08b8b2cc96d92f5abaa275f11b3"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing / 6f54b4b66930 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission

<a id="canonical-11413064cb9a7f6da1ebe0839433fe4456cc14b82883c972613d7778f0b4e316"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop enter drawing submission.

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
shop_enter_drawing_submission = {}
```

<a id="canonical-02eeaafe84c02a3ca9d69aaa6033fc1d3efbbda1c2a8c311f05e3cbd0df43019"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing / 6f54b4b66930 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-38c2bc1d06b28be5db790c1b1818d9d196ab38f501c9f2998f876eb28a0ed731"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing / 6f54b4b66930 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-464b8aad2c00c73e28d52296a31d4eff80e0c7aeef1d9b9598741c809deffd70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-703cc1a6426a0fdf948c5e178e9e8064e2ff021619bb4245742f47a8ea5eda1a"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment / 4dbf66f8bf0d / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment

<a id="canonical-0b1295607c98d7176bac09ebb1c8fa98cc4d3b136929f52f7fa5270285212a5b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop make payment.

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
shop_make_payment = {}
```

<a id="canonical-afac6a99b4eb6dcff5c436c484bc0052a48b33a67dbe2449d0ea0adff8fd5fe7"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment / 4dbf66f8bf0d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4d888a9c3587b54bc0a113791504f9e5f04b3cbd96f82ffee817e6dd4132988c"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment / 4dbf66f8bf0d / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-059d5a47901ebd8bc81b3eeca83d20427d4ca915690483a06f5312eed7f3a250"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb1a02cc88230b569bfd26134391624120190f5eb9cd24aa2d0a39457054e3ad"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order / b44c4354f6c3 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order

<a id="canonical-22d8a702fa7af6cac41f2fe59d891a98dbff5d5c5a26ccfa8181a5e2c914d998"></a>

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
shop_order = {}
```

<a id="canonical-2ec92fb0e694dab2e8d816fb09e2ace6a8537448008b6bd6d22dca4f69b0124c"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order / b44c4354f6c3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-10996764e48ba6d98e0409568d5191dfab82f09c69c70fac50ca4b7bf570d3d6"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order / b44c4354f6c3 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-54c0b58f4744fd6075e89f3a8683b06bb0036358342b1ee9f5af5c0a46d55682"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbe18d0dacde9e1dc157e6108caa1f7791a1d3955e15d413ebf3b4b49e925966"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry / a7d4d735c44a / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry

<a id="canonical-4ffdae5f1255b93a2809abb379c73ace557941bb77cd834e369817d652c71140"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop price inquiry.

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
shop_price_inquiry = {}
```

<a id="canonical-9346b9e6e509d4a58f5cf95e287b5e686db113ede8c4228a28496a60c243fa5b"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry / a7d4d735c44a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-24d76632af638a14dcc4c20518f5520854ea156753f3b68015ba89a23b72e05b"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry / a7d4d735c44a / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-5c41470176f87c10253050a0a24a6b811d4b820f84d6c9cad25df5c44b990cd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a272729f3e21fb694d638e205875eb213feba5f2e012d59e308bebac07eed39f"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_va / 35c146259617 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation

<a id="canonical-c3b5459a019322c5efb0ce74746ec65134b2d7eb266c9c13cac779db95920d62"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop promo code validation.

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
shop_promo_code_validation = {}
```

<a id="canonical-6d474a6fd5f0a02ead9de8c4fd25120cfe02ac7b95997b745fff9831b4de9425"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_va / 35c146259617 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-54e77efdecaed63be391ffdf3d736a932bf5350ebb959267063e1074f162d046"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_va / 35c146259617 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-0b53f92646212bd675d4b3af889c1fc38c1a7371e30a331f41d44abc32512ca1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7e3c3b3a6209eb291e87b805c5b42c2512a60e2981be684f08096062eda9005"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift / c440d267f74f / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card

<a id="canonical-df49a6a411dc26eb472b1065af19fed7759596674a61c71a836f9b23fcc48ffa"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop purchase gift card.

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
shop_purchase_gift_card = {}
```

<a id="canonical-3d30a2953ac63b20a58b45b8e1c74700fc64460dd64dee3842961fdbbf515ec8"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift / c440d267f74f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c216bc4c41377d5ddee4d17071fed102d7011f53add82e4f4009f6906855219a"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift / c440d267f74f / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-b65b6c703796164baafa6a6c9a203d279a3d179788e0960ea1d98c0dc5ce75f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1efae8d241e334998a0edb8cb4f225f5b8d185b0bd1032e7423ce41a0f10cd9"></a>

## cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quanti / 0b792bf46027 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-83c56ec943463a73e070d58457ac98a516ecdcb32f3811a4e5cf0e8151e972b9)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity

<a id="canonical-af700b52fba5ecab03f005e141697da1ffad66ebaf867f57eef165d4ea460a17"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop update quantity.

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
shop_update_quantity = {}
```

<a id="canonical-90c77db9e319a0e3a735b1522253004551a6bcca40fc554b522b4dc3061732e2"></a>

## Direct properties — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quanti / 0b792bf46027 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-de98c577ee17f1228c672c536d571e6331ba464a32320444fe989d424a08c97b"></a>

## Next pages — cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quanti / 0b792bf46027 / 4

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-3396cb70835622a8f702fc1f18dd6846698abb591b6d87ce7616deae55d1c02b)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-26572fbea0b5a97f998e5cd33dc50ad5cec6eb52e8a1fe2fdedd082e0d49f516"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2d8ed5123f08852b404a76229aaff8d922c3d299934c83bcd8311245a53e42b"></a>

## cloudfront.protected_endpoints.metadata — cloudfront.protected_endpoints.metadata / 9696f7016fe6 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- cloudfront.protected_endpoints.metadata

<a id="canonical-c6b6e83335af3e580b478be18a158c2bfb7ab8fc68da0367d04656e46f1cf240"></a>

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

<a id="canonical-8fe992b4e1dbb64a118cf0710d2330e81d193f7579d9320d4cc61f7780cdba42"></a>

## Direct properties — cloudfront.protected_endpoints.metadata / 9696f7016fe6 / 3

<a id="canonical-e67d8edf25607cdbf19d8ae18b4cab2558ecdee15605b59262c4a48693b808c9"></a>

<a id="canonical-4e92caf54b32562e6a56a6147a8337b4758de92f4919048fd77a552013b36eb9"></a>

## description_spec property — cloudfront.protected_endpoints.metadata / 9696f7016fe6 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-f66637c2c29e79a015031118243dac74bf362bc1771a6e92388017c387413ba7"></a>

<a id="canonical-1fb1a1f5bb7aea45477c47bab26e26a1c701467faa1ffb4cb2e2d0f5953b1785"></a>

## name property — cloudfront.protected_endpoints.metadata / 9696f7016fe6 / 5

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

<a id="canonical-58d7601c841339cbd0179b339bd5d40f2471d56c2165b4d563965cbccc3133b5"></a>

## Next pages — cloudfront.protected_endpoints.metadata / 9696f7016fe6 / 6

- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-27b837e1a76705986c2e896c7defd65dfaf775f5eda460fa07a12bb07fc5b5a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd20ea408cc391d30c810a1b76e75dc60c27a2fbccb2e3bd49312481b5a62e27"></a>

## cloudfront.protected_endpoints.mobile_client — cloudfront.protected_endpoints.mobile_client / 5bc0689e8eb2 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- cloudfront.protected_endpoints.mobile_client

<a id="canonical-58b0690a51f840faa3198ae2ad73ed65b9a4795b90053764d6a608f1a661c75e"></a>

Type: `"object"`. single nested block, Optional.

Mobile Client. Mobile client configuration OPTIONS.

Upstream description:

Mobile client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "continue")}
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
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\"]"
}
```

Terraform syntax:

```terraform
mobile_client {
  # Configure direct properties listed below.
}
```

<a id="canonical-f16e6e5bf3a9db6c901477947b8a21e0501a814e3ac8927bd40e21eac8e01bef"></a>

## Direct properties — cloudfront.protected_endpoints.mobile_client / 5bc0689e8eb2 / 3

- [block](resources--protected_application--reference--group-003.md#canonical-dbb160616ab5368f95deb492d724a9d974885ebe55c3b91631ca838ff69f6607): complete subsection reference.

- [continue](resources--protected_application--reference--group-004.md#canonical-66ebba0998d14c9bc002aeeda5ae2a25f5af6d78a8be11eda75ee83d50348699): complete subsection reference.

<a id="canonical-8c74e34a7320595a7c6276c701d0f3e7a8c8ab925fa41b1647ddf08230cb030e"></a>

## Next pages — cloudfront.protected_endpoints.mobile_client / 5bc0689e8eb2 / 4

- [cloudfront.protected_endpoints.mobile_client.block](resources--protected_application--reference--group-003.md#canonical-dbb160616ab5368f95deb492d724a9d974885ebe55c3b91631ca838ff69f6607)
- [cloudfront.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-004.md#canonical-66ebba0998d14c9bc002aeeda5ae2a25f5af6d78a8be11eda75ee83d50348699)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-dbb160616ab5368f95deb492d724a9d974885ebe55c3b91631ca838ff69f6607"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abacb428b627be79670ff9780bdaa39f97413929d79acc2a8981bcab9b9f9829"></a>

## cloudfront.protected_endpoints.mobile_client.block — cloudfront.protected_endpoints.mobile_client.block / 827708b05a43 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-56174b845c9baa387b64f950378b62eefe99ad707dcfee4610c49d8d78c3d1f5)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-cbb2cae717d049950b321546a68a8ba507b347b0def5fb02bd1c9cbaa5cad138)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-bd37d941a3cf9514e38843bb752115f37da757683a0f4be8edc324747b79a9ae)
- [cloudfront.protected_endpoints.mobile_client](resources--protected_application--reference--group-003.md#canonical-27b837e1a76705986c2e896c7defd65dfaf775f5eda460fa07a12bb07fc5b5a4)
- cloudfront.protected_endpoints.mobile_client.block

<a id="canonical-dc003019d30a72221594d38b63313e14941e1572e76bc5526726147f816e5c72"></a>

Type: `"object"`. single nested block, Optional.

Block Response for Mobile. Block Response.

Upstream description:

Block Response.

Receipt-pinned upstream constraints:

```json
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
block {
  # Configure direct properties listed below.
}
```

<a id="canonical-a2a1f7779007af8ab7af34ec096017fbedc6d29847c0c568ef60f25088258b13"></a>

## Direct properties — cloudfront.protected_endpoints.mobile_client.block / 827708b05a43 / 3

<a id="canonical-c1f1dbd01e5e13d86586b994d36a4272122dcd2748513e516252545649e50500"></a>

<a id="canonical-0f5dbead18db96eeca40ece96cc176730fa2c7432af19b2da22fb2d66ed9eaf4"></a>

## body property — cloudfront.protected_endpoints.mobile_client.block / 827708b05a43 / 4

Type: `"string"`. Optional.

Body. Custom body message.

Upstream description:

Custom body message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-a15d067ecce83d131919fddf2c544aef8db51d1bc619f61658b5ec191a508aca"></a>

<a id="canonical-2412c510db1ac27425369a79eedb07d7da994c135b7f6f5699f3281981719ec7"></a>

## content_type property — cloudfront.protected_endpoints.mobile_client.block / 827708b05a43 / 5

Type: `"string"`. Optional.

Content type to use in a block response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-890d778f22938af229acc87deb6f75271460c5e6efd3bf7e3fc8cb14095208a4"></a>
