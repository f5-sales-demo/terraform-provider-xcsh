---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-71c62bea7bbac98fb590d5eb32b30d9891a464a0da887ce38e7457d8dd203ad9"></a>

## name property — api_specification.api_definition / ab9a6f8bd191 / 4

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

<a id="canonical-74ed2249effb7161130893b8b985ff395bef81d36ab982c55b699c22a4fba5d5"></a>

<a id="canonical-7b83470ab21bffc4375bb5fd54141d8df71122909aea58d2276174462b6aaae5"></a>

## namespace property — api_specification.api_definition / ab9a6f8bd191 / 5

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

<a id="canonical-1f05bdd6d81a8e8e4c1d060430cda5eeea40530b5b108a1159e2749fa278ca9b"></a>

<a id="canonical-fe1975a6a1d33f0b4234c8e766eb083e9bd9c05c0b4c3b4324c037c2633dbfef"></a>

## tenant property — api_specification.api_definition / ab9a6f8bd191 / 6

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

<a id="canonical-9fd607d57a7b519505a195ef5bc34024e47371ac185c6d3dd4d3d9a8fbbb9caa"></a>

## Next pages — api_specification.api_definition / ab9a6f8bd191 / 7

- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0c8b75ead1314a36906010fefd8cac6da585efe7139b0a4957dc06afed994c7"></a>

## api_specification.validation_all_spec_endpoints — api_specification.validation_all_spec_endpoints / 5decff185830 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- api_specification.validation_all_spec_endpoints

<a id="canonical-9f15e391db94e40ae5a58890ec4e02d7dfcae06bd8c4f639b9644f4dabef6040"></a>

Type: `"object"`. single nested block, Optional.

API Inventory. Settings for API Inventory validation.

Upstream description:

Settings for API Inventory validation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

Terraform syntax:

```terraform
validation_all_spec_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-92d27c047db72dbb49b470b2ffeace38135b41ad34460c4cde71043f9ff1be1a"></a>

## Direct properties — api_specification.validation_all_spec_endpoints / 5decff185830 / 3

- [fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e): complete subsection reference.

- [settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1): complete subsection reference.

- [validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729): complete subsection reference.

<a id="canonical-85e9e12ee82e919051ac09fab9d99a1abc6bdc6c80e7f375d6c991be3c3b26be"></a>

## Next pages — api_specification.validation_all_spec_endpoints / 5decff185830 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84db3e3b926c0af3852e9f5f20dafe038430b81cbafc07e30849800d7336a21b"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode — api_specification.validation_all_spec_endpoints.fall_through_mode / 1786b4414e4b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="canonical-3ac48784f9388ed4a4a60cd58013064b434877f4c3364b6520ef6e1cfa81ce24"></a>

Type: `"object"`. single nested block, Optional.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("fall_through_mode_allow",
    "fall_through_mode_custom")}
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
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

Terraform syntax:

```terraform
fall_through_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-09b1f76d48c58b10d0e65e422d20635f3f76f8ca3ab18d7a23b3c742dd319a70"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode / 1786b4414e4b / 3

- [fall_through_mode_allow](resources--http_loadbalancer--reference--group-009.md#canonical-414e0fd0403f4afd5385e238631681f510f68fd52ae0dbc8193140eea4294ca0): complete subsection reference.

- [fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-5a0f1886d20a746d392da33d9bb4406927f1053bf5b0b37e30dfa2c233b1ee73): complete subsection reference.

<a id="canonical-51a763c96a701e10cc0ccfbfdccc5236cc1b111fbb142e7261b43ae2358f9a9e"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode / 1786b4414e4b / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow](resources--http_loadbalancer--reference--group-009.md#canonical-414e0fd0403f4afd5385e238631681f510f68fd52ae0dbc8193140eea4294ca0)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-5a0f1886d20a746d392da33d9bb4406927f1053bf5b0b37e30dfa2c233b1ee73)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-414e0fd0403f4afd5385e238631681f510f68fd52ae0dbc8193140eea4294ca0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1eb41d09cbb70afbb612f4027622ea7aecd89a57ac8d2689d91ffb3015be3ea3"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 932d1f95e1ae / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow

<a id="canonical-8fa70fb506e11cacdd111b63b7b49b17afd41ff473ee120c276a34931878c717"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for fall through mode allow.

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
fall_through_mode_allow = {}
```

<a id="canonical-1347421d037932a9feac931d8caba19b7eb431c04f3d3e73ef092cc1d91b0229"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 932d1f95e1ae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34992579edd2797784f9a6fc3a49c1f8a21f5caaa680f62b4d351412e5c2af09"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 932d1f95e1ae / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5a0f1886d20a746d392da33d9bb4406927f1053bf5b0b37e30dfa2c233b1ee73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a28c43346b0aa6e7b5973c32e33d0c8c1e311bff32b151848fd6cf56591cf432"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 7e99cbab6cf6 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="canonical-915ca64fd02366c9f6a246be4a4bccaa3a60d49f3364bdba4aba02ccbf4bd012"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("open_api_validation_rules")}
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
fall_through_mode_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-63fea02a9d8f9d63eeaa5be0a24e1f4981419e0681ceb990b3f7db65eec1315c"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 7e99cbab6cf6 / 3

- [open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0): complete subsection reference.

<a id="canonical-a588c92503deef7ac4483a2264a365f0b308db936a2e6bc80db77822ceccb24b"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 7e99cbab6cf6 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1086285ed8c6c3cb86b72811cfdd0866a8b31b8f32bf94bd973e16d49f799ba0"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9e236ffc17f3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-5a0f1886d20a746d392da33d9bb4406927f1053bf5b0b37e30dfa2c233b1ee73)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-347b337f1f7d1f0f1fa940194aceb0da746e326dc8d2e116e2b1729da80f4139"></a>

Type: `"object"`. list nested block, Optional.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("action_block",
    "action_report"),
  validators.ConflictingListObjectAttributes("action_block",
    "action_skip"),
  validators.ConflictingListObjectAttributes("action_report",
    "action_skip"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_group"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_group",
    "base_path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-cbd2e62d2bb7074c36c930fea7da31c54bd32345dd612bad6b80249f5beace5e"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9e236ffc17f3 / 3

- [action_block](resources--http_loadbalancer--reference--group-009.md#canonical-edead72fea72e49b586d4583f94fbea7eb4a56b884d697450499f5fa131dc294): complete subsection reference.

- [action_report](resources--http_loadbalancer--reference--group-009.md#canonical-32ff652114c04897ca54d71d62c139756138c384c541547133b601062ec8d4b1): complete subsection reference.

- [action_skip](resources--http_loadbalancer--reference--group-009.md#canonical-1d983ca8a087b53c0ba713d978235e16c4ba8ff32c9023b4ad0682ee72249b0f): complete subsection reference.

- [api_endpoint](resources--http_loadbalancer--reference--group-009.md#canonical-96158fa56f9a1f4c43d95c2c623611996d2fedf9dba94e0beceff487d5a8a955): complete subsection reference.

<a id="canonical-55ee49e0e1cd0b7f860bbeeb5fec955a71a965847d68e7859b82c33d337f3613"></a>

<a id="canonical-8a5dbe4aec5846dbf00c6afe67f3581621e799eba9b04eb9ab9211a8faf7df36"></a>

## api_group property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9e236ffc17f3 / 4

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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

<a id="canonical-144e890466f7c4c2a6d138cd7caea31e221e7d8839c8372cb31660cd668220e4"></a>

<a id="canonical-3b0ce149d677269de5de50d8742175fcb5dd0eb796385fddedcad2396c8dbe56"></a>

## base_path property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9e236ffc17f3 / 5

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-009.md#canonical-16177fc823a949fba1543875690f9bf42d4b453e4c6f0275395ed375f310822c): complete subsection reference.

<a id="canonical-a2b053c76d6f6c58e5e65e1400620a46b0611ddc291539664cd7c889d7ee0360"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9e236ffc17f3 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](resources--http_loadbalancer--reference--group-009.md#canonical-edead72fea72e49b586d4583f94fbea7eb4a56b884d697450499f5fa131dc294)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](resources--http_loadbalancer--reference--group-009.md#canonical-32ff652114c04897ca54d71d62c139756138c384c541547133b601062ec8d4b1)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](resources--http_loadbalancer--reference--group-009.md#canonical-1d983ca8a087b53c0ba713d978235e16c4ba8ff32c9023b4ad0682ee72249b0f)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](resources--http_loadbalancer--reference--group-009.md#canonical-96158fa56f9a1f4c43d95c2c623611996d2fedf9dba94e0beceff487d5a8a955)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](resources--http_loadbalancer--reference--group-009.md#canonical-16177fc823a949fba1543875690f9bf42d4b453e4c6f0275395ed375f310822c)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-5a0f1886d20a746d392da33d9bb4406927f1053bf5b0b37e30dfa2c233b1ee73)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-edead72fea72e49b586d4583f94fbea7eb4a56b884d697450499f5fa131dc294"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fec441941d3dcf14065ac7c0f58404195698c436509e6238281ae5a6ace481c"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 720b573cb893 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-5a0f1886d20a746d392da33d9bb4406927f1053bf5b0b37e30dfa2c233b1ee73)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-03f3349f2464850b340fc7076721bc67511fd04091cce62eb428e203dac196ea"></a>

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
action_block = {}
```

<a id="canonical-26f5783950b9711589aea9df918c88c953a8e2bba92b4a26ac8844a5ef0f6107"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 720b573cb893 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9db4bdcc69c32fdbe8f222192dc8cf3a4b13631b12e0a308c33e7b10bda125fe"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 720b573cb893 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-32ff652114c04897ca54d71d62c139756138c384c541547133b601062ec8d4b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bd9ffb868bd8302a6ca4f0bf266684574f4313a629ba92d2fb5d4f4b8b40d2f8"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / e21b0958bb7e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-5a0f1886d20a746d392da33d9bb4406927f1053bf5b0b37e30dfa2c233b1ee73)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-7c2acae13256e442723fea882969c4c35ded93dc76dc563f3abd542ee0fdc65b"></a>

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
action_report = {}
```

<a id="canonical-49f49f91942b93ba09e30a5d289549ec0488fd5b2c0f79124206d698f7560fcd"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / e21b0958bb7e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4ae62b686a237553084c4d7466de4a52556bb607ab6f68513262d51f4494715"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / e21b0958bb7e / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1d983ca8a087b53c0ba713d978235e16c4ba8ff32c9023b4ad0682ee72249b0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86796ac41d16cbfcdaf197a329108a7bbc3000682b97aed68659e59c94dcd833"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9e995bc7d87c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-5a0f1886d20a746d392da33d9bb4406927f1053bf5b0b37e30dfa2c233b1ee73)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-c2a5c124940df8f81161659aa388fb6c742885d6be015bd5ba6c453a829b2dd5"></a>

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
action_skip = {}
```

<a id="canonical-f573ec6022c21374480c2fe6d6aadec924664aa7ac04cb1eba68ea195f260594"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9e995bc7d87c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8408551738c953dc4d338b2946820267b40629807112581b6dfa1a5769c1ffff"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 9e995bc7d87c / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-96158fa56f9a1f4c43d95c2c623611996d2fedf9dba94e0beceff487d5a8a955"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0bee174df16e10778f34a5bcdf344d35c6099faf0b7807cdb72f686dcde89399"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b1db6ea94ae7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-5a0f1886d20a746d392da33d9bb4406927f1053bf5b0b37e30dfa2c233b1ee73)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-62a2ea024ff44eb42b51f6acd56de909291492096e8d9d000b5582a408560732"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea5f82db995af4d9dd578ff5a9a028d4f685c579fcd4915150b4c2246e079681"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b1db6ea94ae7 / 3

<a id="canonical-46d185db77117cb576ed9305134b712f6fa094aacb508c0a7dbdfd9c60f84856"></a>

<a id="canonical-de0566bf23f32aebd3d36f37a715b575d727fd80225fd853a8631a1667e74d67"></a>

## methods property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b1db6ea94ae7 / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-6e79fb60e55f2ffbc91eb71062ab39e9987922e7b6b42d8403682fa1f8195ba3"></a>

<a id="canonical-7c4151bf32d2182e9aea2bd7ce185639a33204188a7971fb44e5e81245195488"></a>

## path property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b1db6ea94ae7 / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-1b975ba9c088524e48686732f713073b7ebc2b562f40087578308f6b6c7da7a8"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / b1db6ea94ae7 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-16177fc823a949fba1543875690f9bf42d4b453e4c6f0275395ed375f310822c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be4639756417cabaf115d574f0ab1aa8f5875cf64f0e1bf00d0cdd3acf6037ed"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 740c40d2ab5e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-dfec026d5d3dc6525775aed1928180fbf8e9487bb3f058faedecc678ebdfcb8e)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-5a0f1886d20a746d392da33d9bb4406927f1053bf5b0b37e30dfa2c233b1ee73)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-9cef5bc076846ae3ae2e7158d5502f58fe1e69359d6cb13b10e7ae2229049841"></a>

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

<a id="canonical-c9b5b65536b2f12f6341c366f951894e0581667eb0fc2d8d272a213005f3d10f"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 740c40d2ab5e / 3

<a id="canonical-bca233dc190d2d8014e10f2db337cb81dd98ae05a6cff714a15b223ba5790c9d"></a>

<a id="canonical-e54eeff48ae44c9a91a51839bc5df8e2114f65cad7f1e56436dc084184b83539"></a>

## description_spec property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 740c40d2ab5e / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-d7862886730d3ffc5884f275165cd56c90d9a388c933e27119371df7e736cfdc"></a>

<a id="canonical-556b10900bff58b6c91ae0d5a32089f8a1d165af743a496733771f004f1468d6"></a>

## name property — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 740c40d2ab5e / 5

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

<a id="canonical-6effe49bf81e59437100027673627eed22000be96b908a53121710c803736d24"></a>

## Next pages — api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_m / 740c40d2ab5e / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-8d45979bad75e084a0248cd28cc3b09398d297fc7c54699397024f83a47304b0)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-160dd5e8b68cca74d882ab5ffef5a570ce11cf11edc3b30ad757d692be4b69a3"></a>

## api_specification.validation_all_spec_endpoints.settings — api_specification.validation_all_spec_endpoints.settings / 650ef8e41f5d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- api_specification.validation_all_spec_endpoints.settings

<a id="canonical-dc6360ca7a7a2fe4d592b81a5db312733418ed20eed1d605050b3448c8182a0a"></a>

Type: `"object"`. single nested block, Optional.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("oversized_body_fail_validation",
    "oversized_body_skip_validation"),
  validators.ConflictingObjectAttributes("property_validation_settings_custom",
    "property_validation_settings_default")}
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
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

Terraform syntax:

```terraform
settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-8cb6f4e5f5b5ea5da45d54be087cef71f1d982396804833e3600a72156df8c6b"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings / 650ef8e41f5d / 3

- [oversized_body_fail_validation](resources--http_loadbalancer--reference--group-009.md#canonical-f6f3f955c066979eb193dbe32bff375ce2e5b5c2ac9c561c64a5e7b17fc96be8): complete subsection reference.

- [oversized_body_skip_validation](resources--http_loadbalancer--reference--group-009.md#canonical-4d9b8dc4db012d3d3e64a2f4f8d09c5126e1e261d381ab687f4d0e3639958442): complete subsection reference.

- [property_validation_settings_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f6baea7a7fe66353ed4fdb306ffcedd597a25c8265d8ee7657a8022ffa2c8b15): complete subsection reference.

- [property_validation_settings_default](resources--http_loadbalancer--reference--group-009.md#canonical-284b257560fb40624ad6e3b2c4f079647af1412e1e125bc2e59c2c197ed05216): complete subsection reference.

<a id="canonical-0dcfcaa853bc1b850426080f5d7b456d6d02506ec95b2b056f1eea6df84267a4"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings / 650ef8e41f5d / 4

- [api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation](resources--http_loadbalancer--reference--group-009.md#canonical-f6f3f955c066979eb193dbe32bff375ce2e5b5c2ac9c561c64a5e7b17fc96be8)
- [api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation](resources--http_loadbalancer--reference--group-009.md#canonical-4d9b8dc4db012d3d3e64a2f4f8d09c5126e1e261d381ab687f4d0e3639958442)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f6baea7a7fe66353ed4fdb306ffcedd597a25c8265d8ee7657a8022ffa2c8b15)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default](resources--http_loadbalancer--reference--group-009.md#canonical-284b257560fb40624ad6e3b2c4f079647af1412e1e125bc2e59c2c197ed05216)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f6f3f955c066979eb193dbe32bff375ce2e5b5c2ac9c561c64a5e7b17fc96be8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e3378bdee0259619f4f81d3bb4e74fc3f6c598f6fdb495f8950e1f893b05658"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / eef55dc5c620 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation

<a id="canonical-2e0f50c6187e467c52122c1b9041651f5259366d0a448137fae17917b4c116ea"></a>

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
oversized_body_fail_validation = {}
```

<a id="canonical-bde1c76ac022df61c9879966ec8fe3c86f928af1a3012dfa2f1acbeb332ce9b9"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / eef55dc5c620 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03cbec0a36554e46606339f2b1cbea5204d7f75539250f0246c2c3a542bc6ef4"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_val / eef55dc5c620 / 4

- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-4d9b8dc4db012d3d3e64a2f4f8d09c5126e1e261d381ab687f4d0e3639958442"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a27fa0f02f40a9108192a2dea506312a4cd180dc8d1c848c044944227a2f6606"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / 62bdf866e55a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation

<a id="canonical-f900990a49e735195e129b4b85fa06d3fb5047666e1313268e6b961b51cb1190"></a>

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
oversized_body_skip_validation = {}
```

<a id="canonical-a82df4b8783815b4701a68e937feba5b6c9a9c57e5469cfd54a86f506bcf0eab"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / 62bdf866e55a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e91c3b61d1aa7d9bf37a1767d8e6b6aae1de5fbf50b24a996e8798453bfd3cd8"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_val / 62bdf866e55a / 4

- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f6baea7a7fe66353ed4fdb306ffcedd597a25c8265d8ee7657a8022ffa2c8b15"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-05f465c04556f8d83eac9ebef6c7beee3ebe7ef7fffd910ab1cc28e1fc70d372"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 57c94e720da9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="canonical-ea4c45a5d566bd490bda7f60f078e58dc7d6af5146e009b2dc94ebc07f6ac3c2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for property validation settings custom.

Upstream description:

Custom property validation settings.

Receipt-pinned upstream constraints:

```json
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
property_validation_settings_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-200c758a4f1ff6261d9c842d9425722fa051944c606172377448b870fe7ae4fa"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 57c94e720da9 / 3

- [query_parameters](resources--http_loadbalancer--reference--group-009.md#canonical-7dfa4b5ee5f758edceed713ccb0fc0f10fb9ea3b94cf8ac2f523cd4b8c007b38): complete subsection reference.

<a id="canonical-ccf38d6fd7b2a9b3b8e1ec13f66795b6ee7dfeb5c159aedda2882762f62769fa"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 57c94e720da9 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-009.md#canonical-7dfa4b5ee5f758edceed713ccb0fc0f10fb9ea3b94cf8ac2f523cd4b8c007b38)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7dfa4b5ee5f758edceed713ccb0fc0f10fb9ea3b94cf8ac2f523cd4b8c007b38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59c80251e7de3de52d1476a1ac28df234c1e126f78c8ce0451b48a22aeaf9279"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / d637cd2b5081 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f6baea7a7fe66353ed4fdb306ffcedd597a25c8265d8ee7657a8022ffa2c8b15)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

<a id="canonical-dd03c03ab6b868219b6d4ddce48d30dc3443ed5ea15aaea5a208fe624341a499"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow_additional_parameters",
    "disallow_additional_parameters")}
```

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-f91ba30a0bd05d457017f298aecb414250e5134c7a4c5a2d95f2972a0bb774fc"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / d637cd2b5081 / 3

- [allow_additional_parameters](resources--http_loadbalancer--reference--group-009.md#canonical-bbfbd13c063d2fdbcdb41b264dd42fc58dbdcb11cb856e62b4a3529e5a54e93f): complete subsection reference.

- [disallow_additional_parameters](resources--http_loadbalancer--reference--group-009.md#canonical-a95b93062a6473d3e038a7a0fb73844984c77ad7ed7f9bae0e4083e06a80f70d): complete subsection reference.

<a id="canonical-92518d45d8b67e7d8a340cd767080c8ab6374f2176cb41c060435bc10bcacee9"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / d637cd2b5081 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](resources--http_loadbalancer--reference--group-009.md#canonical-bbfbd13c063d2fdbcdb41b264dd42fc58dbdcb11cb856e62b4a3529e5a54e93f)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](resources--http_loadbalancer--reference--group-009.md#canonical-a95b93062a6473d3e038a7a0fb73844984c77ad7ed7f9bae0e4083e06a80f70d)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f6baea7a7fe66353ed4fdb306ffcedd597a25c8265d8ee7657a8022ffa2c8b15)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bbfbd13c063d2fdbcdb41b264dd42fc58dbdcb11cb856e62b4a3529e5a54e93f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a563e05be77ee42651b71ec96bbcbb066b5833705f860d493a7676af310a76d"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 78c82176e342 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f6baea7a7fe66353ed4fdb306ffcedd597a25c8265d8ee7657a8022ffa2c8b15)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-009.md#canonical-7dfa4b5ee5f758edceed713ccb0fc0f10fb9ea3b94cf8ac2f523cd4b8c007b38)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-08b2d4bd4ad7b48ad3cba09bf1f8f5f6a891602eed30c34b7d5ef6f5d0d7b725"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow additional parameters.

Terraform syntax:

```terraform
allow_additional_parameters = {}
```

<a id="canonical-d5a78dd60a44bad32251131561dbbcbfa7f8b86b4f0c6f9263ee9ff1cdfd6aaa"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 78c82176e342 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-edca5f04f30540e2b078eed0061e8b6c10597c3c6e97dcf710380f4b9f7b24b2"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / 78c82176e342 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-009.md#canonical-7dfa4b5ee5f758edceed713ccb0fc0f10fb9ea3b94cf8ac2f523cd4b8c007b38)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a95b93062a6473d3e038a7a0fb73844984c77ad7ed7f9bae0e4083e06a80f70d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78a3d68da130eab4039cb886e7e058b138339fed667395b24a845ef62692687a"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f458f5678ee4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f6baea7a7fe66353ed4fdb306ffcedd597a25c8265d8ee7657a8022ffa2c8b15)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-009.md#canonical-7dfa4b5ee5f758edceed713ccb0fc0f10fb9ea3b94cf8ac2f523cd4b8c007b38)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-6db59c2b52fa6b9594b53b2ec9efe2f91b082d09e0cb2baffe22424ad1b4f1f6"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disallow additional parameters.

Terraform syntax:

```terraform
disallow_additional_parameters = {}
```

<a id="canonical-fc03628e46fd31a85857e816c3efd89309488e10eb183f96728ac771040ee1e5"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f458f5678ee4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5c9985d4d2f6eea20ae6861ce573a87df7c6f28cae512c4ef41dbdd830822a1f"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / f458f5678ee4 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--http_loadbalancer--reference--group-009.md#canonical-7dfa4b5ee5f758edceed713ccb0fc0f10fb9ea3b94cf8ac2f523cd4b8c007b38)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-284b257560fb40624ad6e3b2c4f079647af1412e1e125bc2e59c2c197ed05216"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6a308e23ef669db5560108bfebb8b2d1fdf9334a96e2aed747e818866e8caf1"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default — api_specification.validation_all_spec_endpoints.settings.property_validation_set / b05e29313d5e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default

<a id="canonical-7a490eaf81fd3cad0191bf7f48bfbb29225b1add9de8ff68eab11c5d51034782"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for property validation settings default.

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
property_validation_settings_default = {}
```

<a id="canonical-9297db34291de3a067631eadc453611f1c895005827636da235d257d778956f1"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.settings.property_validation_set / b05e29313d5e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4cf03f1dac5ef849d6d534be9d89f961b5b9ff31de255679320a633c3321e2e8"></a>

## Next pages — api_specification.validation_all_spec_endpoints.settings.property_validation_set / b05e29313d5e / 4

- [api_specification.validation_all_spec_endpoints.settings](resources--http_loadbalancer--reference--group-009.md#canonical-bdf74314c5777ee0ee6435a0d1b8ba9b0f492c373b7b306f384c364c8a044be1)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8aaff1d263e297750acbf1bc1c9a72868dc47d8368e52c4693596675520df47"></a>

## api_specification.validation_all_spec_endpoints.validation_mode — api_specification.validation_all_spec_endpoints.validation_mode / 1b8a93d102d9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- api_specification.validation_all_spec_endpoints.validation_mode

<a id="canonical-8b4b9bbdd560c3b0172af87dbd8ecabae1196ece5d4149b104cab7a51034a2b2"></a>

Type: `"object"`. single nested block, Optional.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("response_validation_mode_active",
    "skip_response_validation"),
  validators.ConflictingObjectAttributes("skip_validation",
    "validation_mode_active")}
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
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

Terraform syntax:

```terraform
validation_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-fb66369684122fcc534d4875822e40845e1caeb1b04eb01f4684f73285ad291d"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode / 1b8a93d102d9 / 3

- [response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-cdd57a2bf2a9b94c20f67357072366748ae7895e727ca1caef2ebfcd684975df): complete subsection reference.

- [skip_response_validation](resources--http_loadbalancer--reference--group-009.md#canonical-103e88885d26db3baa1e8879259bd4f6e222872b68758d016467fcf103b7a5cd): complete subsection reference.

- [skip_validation](resources--http_loadbalancer--reference--group-009.md#canonical-bf7bee7d987f182f77ec71fcb010bf139a9c266280bbe8f98ac6650f72d208c0): complete subsection reference.

- [validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-8deec2cd382531339390720d7625f37ddd54da0e6291ed4a80efff629cb765ba): complete subsection reference.

<a id="canonical-20bf56193ab0413619dc146725c98bdea516f644a5fc91b4e4308c43ce9d557b"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode / 1b8a93d102d9 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-cdd57a2bf2a9b94c20f67357072366748ae7895e727ca1caef2ebfcd684975df)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation](resources--http_loadbalancer--reference--group-009.md#canonical-103e88885d26db3baa1e8879259bd4f6e222872b68758d016467fcf103b7a5cd)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_validation](resources--http_loadbalancer--reference--group-009.md#canonical-bf7bee7d987f182f77ec71fcb010bf139a9c266280bbe8f98ac6650f72d208c0)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-8deec2cd382531339390720d7625f37ddd54da0e6291ed4a80efff629cb765ba)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-cdd57a2bf2a9b94c20f67357072366748ae7895e727ca1caef2ebfcd684975df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a60165d15eede3924dcfb10fada18bb0a29b3385eaa9dcdd3225748854fa2e3"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 5c8b854b9714 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active

<a id="canonical-285264877c29685adccf03e88d771a9e2a5c047f8032795b01907df8adf57c34"></a>

Type: `"object"`. single nested block, Optional.

Open API Validation Mode Active. Validation mode properties of response.

Upstream description:

Validation mode properties of response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
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
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
response_validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-b25d45cae89a74d95b7964c98942454abc475954009530df9103308b69a6496c"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 5c8b854b9714 / 3

- [enforcement_block](resources--http_loadbalancer--reference--group-009.md#canonical-5d98dee1a083c6b4c0804b85f2348ec834693db76b4299efe455d57364e34722): complete subsection reference.

- [enforcement_report](resources--http_loadbalancer--reference--group-009.md#canonical-471955e0a72a118cab3d8857b5d46c1ccc4f8060371d0a139482ad662843232a): complete subsection reference.

<a id="canonical-b60d477f9dd89c52c6e53ed6da748947e17d8b6a5bd1553662c10068fb55e76b"></a>

<a id="canonical-f4a77f1072fa876f9372afef7e563278259040595a52490eafd65bbdae0ec7b8"></a>

## response_validation_properties property — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 5c8b854b9714 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-731738d749f50e15d0e53b2def301c958999d19f948423e2fedca71c0bd3da1f"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 5c8b854b9714 / 5

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block](resources--http_loadbalancer--reference--group-009.md#canonical-5d98dee1a083c6b4c0804b85f2348ec834693db76b4299efe455d57364e34722)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report](resources--http_loadbalancer--reference--group-009.md#canonical-471955e0a72a118cab3d8857b5d46c1ccc4f8060371d0a139482ad662843232a)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-5d98dee1a083c6b4c0804b85f2348ec834693db76b4299efe455d57364e34722"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b3203b4973e44f3f35a08695cb9c58dd2aeb75842259e16c1b07f54d4ef18c6"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / d8e10c3fa033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-cdd57a2bf2a9b94c20f67357072366748ae7895e727ca1caef2ebfcd684975df)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-09b02fb227cb317a229787dc7faa8117eeb9b2f11cb6843dbcb72fd8f3ad847b"></a>

Type: `["object", {}]`. Optional.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
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
enforcement_block = {}
```

<a id="canonical-7059f5e0b340b9a27b356f1162bc88ede9889dd158b27c370d695c2d6a258a97"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / d8e10c3fa033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4b3f125e846c4f7f2498b47765f56d892810eb049cf00a3f84d614368e550efc"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / d8e10c3fa033 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-cdd57a2bf2a9b94c20f67357072366748ae7895e727ca1caef2ebfcd684975df)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-471955e0a72a118cab3d8857b5d46c1ccc4f8060371d0a139482ad662843232a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fff5b318bc46fd2b5b0ff16bc92479ac6454596e408cf0d98d81d37eba9bf3b3"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 75956f227c58 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-cdd57a2bf2a9b94c20f67357072366748ae7895e727ca1caef2ebfcd684975df)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-73b01f85eddd1fc487c93e23b75ed00caf9986da8e1593648b88112b2955fe33"></a>

Type: `["object", {}]`. Optional.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
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
enforcement_report = {}
```

<a id="canonical-f8fa66ff280f60e4aa1ba7a1db306a6acdd772d6499f09fdd9f61470dae9354d"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 75956f227c58 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c9d6f0ebe20b539153e570ed714d5708dac7a0761068ed2e6b5bc8c0ea2609b3"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.response_validat / 75956f227c58 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-cdd57a2bf2a9b94c20f67357072366748ae7895e727ca1caef2ebfcd684975df)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-103e88885d26db3baa1e8879259bd4f6e222872b68758d016467fcf103b7a5cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06e3aeadc2696492f1d890d3617c38f124f0eadadf5e4ac9d28e8855b5e0ccec"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation — api_specification.validation_all_spec_endpoints.validation_mode.skip_response_va / bfb56c32bcc7 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

<a id="canonical-ac5ddb26774359bc31fd3291ab03ace6dbe1f77eed422e0453f0771490eb7438"></a>

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
skip_response_validation = {}
```

<a id="canonical-cfc5b40a83c07b60bfa6477f27dcf8a422de56251dc340ab98567afd86c7c9a2"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.skip_response_va / bfb56c32bcc7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-79b95c7aa35f9db58aedf7aa826866defdc0d3c2b3f10aa1abaf21cfc5188d6a"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.skip_response_va / bfb56c32bcc7 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-bf7bee7d987f182f77ec71fcb010bf139a9c266280bbe8f98ac6650f72d208c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e209b378738103e677bde645af3514f098cf57f94393c1234fb71fa8d3cc9ec"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.skip_validation — api_specification.validation_all_spec_endpoints.validation_mode.skip_validation / d2e8df400f52 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_validation

<a id="canonical-fb2720727aa91ee34ae668a440e7ad7504738465f92874079e11e86fea2d0fa7"></a>

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
skip_validation = {}
```

<a id="canonical-b29d56f1b05cdb76098111f0c4bed180bcfe055cf5ec0254f7e0bc954cdd8172"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.skip_validation / d2e8df400f52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-45ffb839c4a051aabd6775c6d6ccc32914a5cf63df4504e22655a86504b9d12d"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.skip_validation / d2e8df400f52 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8deec2cd382531339390720d7625f37ddd54da0e6291ed4a80efff629cb765ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c98036bc55d319e11c39029cd5914b2b2ca076250e1418d3ef5012bf3cd5726"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 539421eacd0d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="canonical-266c8920913e15ee490d198d26089007ae37d7ed29b13a125cbe991edd4067cc"></a>

Type: `"object"`. single nested block, Optional.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
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
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-736560e8ed13ae6aef208fae85456d815dac667bf7d94b16c6d84c4a60206108"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 539421eacd0d / 3

- [enforcement_block](resources--http_loadbalancer--reference--group-009.md#canonical-063444f839d4456479a4b7fe783d02f87a11c218cdcac92dabc1f292222fee4d): complete subsection reference.

- [enforcement_report](resources--http_loadbalancer--reference--group-009.md#canonical-8f89df4c7eb13fef24e14e349f3865e9cbc2b87ad5ef7043d8ad87c7056e61f0): complete subsection reference.

<a id="canonical-86e2d361576cf07a07d86c3b1993c41267e1e4afd507bd70f36d76393e739237"></a>

<a id="canonical-80d0082600ec13fc8814a220aad589407a094d6a027e0584744b44b029c69043"></a>

## request_validation_properties property — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 539421eacd0d / 4

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-aa84a29a611b5e5b23ac4f9bfe1d4638a1f97fad4097594eff4fc761a5af52e6"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 539421eacd0d / 5

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block](resources--http_loadbalancer--reference--group-009.md#canonical-063444f839d4456479a4b7fe783d02f87a11c218cdcac92dabc1f292222fee4d)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report](resources--http_loadbalancer--reference--group-009.md#canonical-8f89df4c7eb13fef24e14e349f3865e9cbc2b87ad5ef7043d8ad87c7056e61f0)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-063444f839d4456479a4b7fe783d02f87a11c218cdcac92dabc1f292222fee4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a96341d13532427a48da938cd855d60aafca8fd88364fd3efdcd76182f4dd190"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 2ffff8ade7ab / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-8deec2cd382531339390720d7625f37ddd54da0e6291ed4a80efff629cb765ba)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-8444895e3867da0c41605ecb48f17f15b1c6325ede669d29abd9130398e019a4"></a>

Type: `["object", {}]`. Optional.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
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
enforcement_block = {}
```

<a id="canonical-4fd67f0bcec70893df2c389d0768ad517aaaf51ccc4641e957f942beacbb7342"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 2ffff8ade7ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-33f77da0afe50769a7e2156150dca4efc5a751c65a4912de636e060b80654f04"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / 2ffff8ade7ab / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-8deec2cd382531339390720d7625f37ddd54da0e6291ed4a80efff629cb765ba)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-8f89df4c7eb13fef24e14e349f3865e9cbc2b87ad5ef7043d8ad87c7056e61f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46b5df41c527ea90f597b325e3aeb7713bcd17316750a801c59ec0821e967476"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / b39b9345c012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_all_spec_endpoints](resources--http_loadbalancer--reference--group-009.md#canonical-28449bf4719d15400a78fc755eadc246aea337db20cd73ec04881764d8ecc15c)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-27a870f5c2cfdc1ca5271a5ed63268c1671eae73b7cd3428c12085f108fa9729)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-8deec2cd382531339390720d7625f37ddd54da0e6291ed4a80efff629cb765ba)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-021813c91bcda1a8db5e2485e9645a0018eda513965b45c636047c89e14fd5e6"></a>

Type: `["object", {}]`. Optional.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
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
enforcement_report = {}
```

<a id="canonical-7b8675ef262c25dfe4a9430ff947ec91dc520d18820225c7b114225803a88bd8"></a>

## Direct properties — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / b39b9345c012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b45e5ada3185d9013a7d7c6c27a7c86f3f7d6beea1f1dc5cf7c68221c914a135"></a>

## Next pages — api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_ / b39b9345c012 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-8deec2cd382531339390720d7625f37ddd54da0e6291ed4a80efff629cb765ba)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1833ac07a0aa9917442bed35b0ea0d8ada9ae025042b0d58454a95cb386b001"></a>

## api_specification.validation_custom_list — api_specification.validation_custom_list / 6bf5fe36d7fb / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- api_specification.validation_custom_list

<a id="canonical-06cce790bbd7379204104194c17c2c5106b46e6642d1f889ccfb8e5f9e530831"></a>

Type: `"object"`. single nested block, Optional.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Upstream description:

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to "Fall Through Mode".

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("open_api_validation_rules")}
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
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

Terraform syntax:

```terraform
validation_custom_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-ad2aadc103489cb7ea1494d0773c9d3ef1b7ba64fa1a7ab98821bb30461bd8d2"></a>

## Direct properties — api_specification.validation_custom_list / 6bf5fe36d7fb / 3

- [fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30): complete subsection reference.

- [open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601): complete subsection reference.

- [settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b): complete subsection reference.

<a id="canonical-83820231f3dbad56d97d260b61e6d9d662c5d4ee629209157a527176967aed39"></a>

## Next pages — api_specification.validation_custom_list / 6bf5fe36d7fb / 4

- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [api_specification.validation_custom_list.settings](resources--http_loadbalancer--reference--group-010.md#canonical-983c0f7b8b65130a581f82668eae61ae86c32f3cb977721f3ee830baad3f8e2b)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da9b102ba46f9460f4614dd8040ac58c2e62a9505f0c1d638cb7660fbd8858de"></a>

## api_specification.validation_custom_list.fall_through_mode — api_specification.validation_custom_list.fall_through_mode / b22ed9c083ca / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- api_specification.validation_custom_list.fall_through_mode

<a id="canonical-35e14c5090e54f9358b9e79a6cbd5810d09f11331a9991b1c2ea4bc52dc0d88d"></a>

Type: `"object"`. single nested block, Optional.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("fall_through_mode_allow",
    "fall_through_mode_custom")}
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
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

Terraform syntax:

```terraform
fall_through_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-e057f5e5cb3db6d2daf6afd3580b33cfccb4e5a00c7072c97f2fc2580524d7a1"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode / b22ed9c083ca / 3

- [fall_through_mode_allow](resources--http_loadbalancer--reference--group-009.md#canonical-998d307205ba09c52f18bdbb2f2a9c6a1711b0ce1e5faea473f524000ee06e48): complete subsection reference.

- [fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f9df52f84288a77aa42540005907df7b7954593de434ed91c38c2924ab52c54f): complete subsection reference.

<a id="canonical-e8b4f4809642fe41a86713ac63edd4737bc075ecd226737649848862f5bc9199"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode / b22ed9c083ca / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow](resources--http_loadbalancer--reference--group-009.md#canonical-998d307205ba09c52f18bdbb2f2a9c6a1711b0ce1e5faea473f524000ee06e48)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f9df52f84288a77aa42540005907df7b7954593de434ed91c38c2924ab52c54f)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-998d307205ba09c52f18bdbb2f2a9c6a1711b0ce1e5faea473f524000ee06e48"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e524bf2e19030700a0d60d6241c98c1e4a0d1af7d3942e941d6ba7691f5daa80"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_all / b8b5ec6db715 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow

<a id="canonical-bbb5850582c666cd813e64fec65ac4304d5dfb0a40f2274cff8c19bf51f11789"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for fall through mode allow.

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
fall_through_mode_allow = {}
```

<a id="canonical-5d820c700a58faf7787a7f5dff4a92b946c473ce606a56a3c807eff2ef6b9397"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_all / b8b5ec6db715 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-916e1867d1258b6a0338590ee86244d66b27acf68bf010a704740c8a7405b981"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_all / b8b5ec6db715 / 4

- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f9df52f84288a77aa42540005907df7b7954593de434ed91c38c2924ab52c54f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b480cbeae41d31ada8d9a61d4b244b731c4bb865ea9f8943b4534229212eae58"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 17b28f00a35a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom

<a id="canonical-8abe57fccaf478b0025e37fd01a2b869d67e3f68775bf24276b84570ff07e940"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("open_api_validation_rules")}
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
fall_through_mode_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-cbe9fa0882bc762fb3a7eda4ed005460410ba882756a8f93d465f4c90a446538"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 17b28f00a35a / 3

- [open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4): complete subsection reference.

<a id="canonical-708510e43d8709c48f6beb16722df97de94165d5c96c403fcc1c455422878795"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 17b28f00a35a / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-552f28a70e03eeb16f17d959ce5e7924ee9e084b3f15b8efa03b7895fe6282d3"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / c3055fecb639 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f9df52f84288a77aa42540005907df7b7954593de434ed91c38c2924ab52c54f)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-4b709c38e940618f0e2c596bb50c7a2007b2131b6e8aa46d008d31f175ae9dd9"></a>

Type: `"object"`. list nested block, Optional.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("action_block",
    "action_report"),
  validators.ConflictingListObjectAttributes("action_block",
    "action_skip"),
  validators.ConflictingListObjectAttributes("action_report",
    "action_skip"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_group"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_group",
    "base_path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-d02ea74d28d9c2ebd59bebebec4b1c8ddd97e6106b57d08457fe8a27afead3b9"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / c3055fecb639 / 3

- [action_block](resources--http_loadbalancer--reference--group-009.md#canonical-1a2f54f40756d9e5f5746da4d220da2eb8abaa16b8a9a0d181c7a4ea020e4e5a): complete subsection reference.

- [action_report](resources--http_loadbalancer--reference--group-009.md#canonical-a7598d7bfcfeee86222ffe9096c6d4be2b65d7f10932277cbab6fdf0a31ff421): complete subsection reference.

- [action_skip](resources--http_loadbalancer--reference--group-009.md#canonical-9d0618c299668bf45bf12cde63a2ab7ae15ad2e71bd8c074dfe3b2e8a9027200): complete subsection reference.

- [api_endpoint](resources--http_loadbalancer--reference--group-009.md#canonical-04e1a835944f0701140151849a303b01645c357343ce87ae99dd08563000f624): complete subsection reference.

<a id="canonical-205e881531a1b6158efd283099cce80c0f46bb77911916c8ae48357c97aad065"></a>

<a id="canonical-0b3077a04ee7befc40b106a6cc3be113f8223515efa9881150a3f08fa190528b"></a>

## api_group property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / c3055fecb639 / 4

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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

<a id="canonical-3fbc7c8679af8a11c28d2a95750b4424ad30d4180a6163f836e478772fe669a6"></a>

<a id="canonical-2fbf86273aa3ad2e4163347e42b8510b5888c6502c733afc9ce338002af08b50"></a>

## base_path property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / c3055fecb639 / 5

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-009.md#canonical-06611c08f5a1e939293eef98f37b8f9e39ff362d548a89a78876748625196ce7): complete subsection reference.

<a id="canonical-38f9e2ca37c25f93a51f59c5b815821873a555eb648e8fcef1cb46766d5bc1fa"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / c3055fecb639 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](resources--http_loadbalancer--reference--group-009.md#canonical-1a2f54f40756d9e5f5746da4d220da2eb8abaa16b8a9a0d181c7a4ea020e4e5a)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](resources--http_loadbalancer--reference--group-009.md#canonical-a7598d7bfcfeee86222ffe9096c6d4be2b65d7f10932277cbab6fdf0a31ff421)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](resources--http_loadbalancer--reference--group-009.md#canonical-9d0618c299668bf45bf12cde63a2ab7ae15ad2e71bd8c074dfe3b2e8a9027200)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](resources--http_loadbalancer--reference--group-009.md#canonical-04e1a835944f0701140151849a303b01645c357343ce87ae99dd08563000f624)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](resources--http_loadbalancer--reference--group-009.md#canonical-06611c08f5a1e939293eef98f37b8f9e39ff362d548a89a78876748625196ce7)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f9df52f84288a77aa42540005907df7b7954593de434ed91c38c2924ab52c54f)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1a2f54f40756d9e5f5746da4d220da2eb8abaa16b8a9a0d181c7a4ea020e4e5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba21d12424686a6ea086ef72102db9620c60861a82823bba3e24f72bce5e88b7"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / b1d26bad2024 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f9df52f84288a77aa42540005907df7b7954593de434ed91c38c2924ab52c54f)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-4b38f9d76d04123368963e3446a615cbc2216a974e02a4b78fa0964a92c397e9"></a>

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
action_block = {}
```

<a id="canonical-360e4ef215370a4caae73de97f2a3362d01de29316c129a22676fce2910901d7"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / b1d26bad2024 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-558b5e78190ff06a11e5389480b200f4b418d7412c9903ea36b6308fecfbdb88"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / b1d26bad2024 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-a7598d7bfcfeee86222ffe9096c6d4be2b65d7f10932277cbab6fdf0a31ff421"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bc133ecaaae28f531536fbdccdbde20bc1dd0b4e9737cff6749f3baebb7155d"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 2a60e2d6608e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f9df52f84288a77aa42540005907df7b7954593de434ed91c38c2924ab52c54f)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-dde8ff330cfef32ed1b6b5dc292353f65ecf2e6fc9000aeb0d116a9f19219caf"></a>

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
action_report = {}
```

<a id="canonical-304063ebcc485885a24c66454e8bb0c264f09edd238cc5db802291915cad202a"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 2a60e2d6608e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cb37acb977529c6828c358d4e4352281c1b769eaa45de34bd751433ade928d6f"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 2a60e2d6608e / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-9d0618c299668bf45bf12cde63a2ab7ae15ad2e71bd8c074dfe3b2e8a9027200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7ec42f9b607f547d6242cee4d6ab4b01bcab3a9b638e00f848c5cf5a2a977cf"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 1593dbb28069 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f9df52f84288a77aa42540005907df7b7954593de434ed91c38c2924ab52c54f)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-e0ab7d5b38445bd3c1e69db3e53187998873ab069537b27e42257277b14e58aa"></a>

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
action_skip = {}
```

<a id="canonical-bc2ca3ef66bbd4571369d30f453dccd327a13eb84a6d21c0be65c27f89a1b6ad"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 1593dbb28069 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f818b4361157ddd1031af48805e6f26ae5f31b3cc42e4421fd4f802226b2f544"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 1593dbb28069 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-04e1a835944f0701140151849a303b01645c357343ce87ae99dd08563000f624"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3a8cbff939b51e29faae236309790d88af86709b2784134f16bb81eaf5be5bb"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 641b9f4bd768 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f9df52f84288a77aa42540005907df7b7954593de434ed91c38c2924ab52c54f)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-35d0377d799fcddabf11adcc47b2eaa1e1b0253cc4b593c82d719a9640a51e0a"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-df80efae171ae117beccac1c0671f94b7f40c3b0a55dbe931e8dd51f0c2a8220"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 641b9f4bd768 / 3

<a id="canonical-a32deebcbd34599a87bb2bf8448348f660b82f204ec81ec43a6d1787788be88e"></a>

<a id="canonical-c540ece15c0877eb12b1a0bbdf6ffbfa3a5e4c9196fdef0e7a157a2fc6d1403e"></a>

## methods property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 641b9f4bd768 / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4e170156113409ca4226b08b81c230ac839f13486dec4d6d39f67128a9ecc638"></a>

<a id="canonical-dabd6bb34912b2e167a851e9ea291c10bd0ef4d57da5b28dcc91ff032f078537"></a>

## path property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 641b9f4bd768 / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-b83d47f57b1d2a40cddf0ce60f0a1c314ed45c4133c29cc75b8f19be1da5361a"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / 641b9f4bd768 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-06611c08f5a1e939293eef98f37b8f9e39ff362d548a89a78876748625196ce7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45969087c92ae5146bc1916d3dd03d0a1e82b3838dc74e34d6724b91d91a3788"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / d1789bd5bd62 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.fall_through_mode](resources--http_loadbalancer--reference--group-009.md#canonical-ed2e4be85d1d7d6dd4ac1565c3eb17110a2f461efa0d03773b9e2d6b683a2a30)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--http_loadbalancer--reference--group-009.md#canonical-f9df52f84288a77aa42540005907df7b7954593de434ed91c38c2924ab52c54f)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-9be02ed9f1b37ebca7ed6c9dc8a2ed27b6579bb2e7d792949a892931d4973594"></a>

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

<a id="canonical-af19cb75cc968d7f769cfa462ca0a00cb8b0cfce3bffaa64c7530d58dce3f915"></a>

## Direct properties — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / d1789bd5bd62 / 3

<a id="canonical-5f631679f85fe254ffecb4850a881e84d07ddf3887c6dc5daf66c706775037c7"></a>

<a id="canonical-0bd5245f73c8cc2cf41e2845933f4fdadeb947ebf04732e566e7183e6eb9142c"></a>

## description_spec property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / d1789bd5bd62 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-5581485000d7b2969ed98a7e39bb82277e3880c28b74fb0d238f8c06c9fabf34"></a>

<a id="canonical-ceec142985f4236d67c36ff256fc14cd0c75f3175bf72c7572598592fc98b1ce"></a>

## name property — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / d1789bd5bd62 / 5

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

<a id="canonical-25362cdfc1b4f19b8a83b724f3f10561262e16b135b374e52105177ec260becc"></a>

## Next pages — api_specification.validation_custom_list.fall_through_mode.fall_through_mode_cus / d1789bd5bd62 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-72e877b019535bc254929026d75ec97b100072160d605c5d74e7d293890f57d4)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abfc02d18ea1a14886896c84087d4dc2601076e3bfb3ed824022c63bd0564119"></a>

## api_specification.validation_custom_list.open_api_validation_rules — api_specification.validation_custom_list.open_api_validation_rules / 0bbeb5234252 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="canonical-a81f9aee4c20ae4f20fdefec83b622679099696abfa9ae876ab4559a33654847"></a>

Type: `"object"`. list nested block, Optional.

Validation List. Rule or policy definition

Upstream description:

Rule or policy definition

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_group"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_group",
    "base_path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-8a782eb54e49158813452bc4a6f7a49029ca55d8e0e8596f9ccca272f247bcbd"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules / 0bbeb5234252 / 3

- [any_domain](resources--http_loadbalancer--reference--group-009.md#canonical-22e6f34ada48f9e0ca612543438be0c780be2fa4b66fc9fbcf55725af4faf50f): complete subsection reference.

- [api_endpoint](resources--http_loadbalancer--reference--group-009.md#canonical-33c96b1a2245fe1adb814968da0a74f0ecb81e72a9c285eef8bfad8e22b6f9f0): complete subsection reference.

<a id="canonical-6d98cbc2c29b668fa9b67cd8969a7f347258643a411453eef088cb23546afd39"></a>

<a id="canonical-b29d4a2677c6f738cdd9cf0a04a9b991a737a1487699f4c30a103db230653285"></a>

## api_group property — api_specification.validation_custom_list.open_api_validation_rules / 0bbeb5234252 / 4

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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

<a id="canonical-8a38710cfa0b3feb5fca657e61657b5b966119b6d38044e338b071743e7c36f1"></a>

<a id="canonical-ef54cdfc73e1a932b974c39f4ff869ecd1ecbabfec4034044af8fec6d22c4938"></a>

## base_path property — api_specification.validation_custom_list.open_api_validation_rules / 0bbeb5234252 / 5

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-009.md#canonical-d909cf4697742fedfea4bb0454328cfa4cf8ac5e1f4c09801bf45fdfb6e4fced): complete subsection reference.

<a id="canonical-48b17995572fce3f7ee39bda6f19a7b9a058b52bddd909fae7e33a12b21b7a5b"></a>

<a id="canonical-3ebbcce625294edad388c598ea2d2ae8a33e9ff3f84462c0f7443ee52529aebd"></a>

## specific_domain property — api_specification.validation_custom_list.open_api_validation_rules / 0bbeb5234252 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

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
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf): complete subsection reference.

<a id="canonical-8e63f97f4cee2a0795edbac49660bfdca00f70b528309115ee7eb3ef2029ab12"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules / 0bbeb5234252 / 7

- [api_specification.validation_custom_list.open_api_validation_rules.any_domain](resources--http_loadbalancer--reference--group-009.md#canonical-22e6f34ada48f9e0ca612543438be0c780be2fa4b66fc9fbcf55725af4faf50f)
- [api_specification.validation_custom_list.open_api_validation_rules.api_endpoint](resources--http_loadbalancer--reference--group-009.md#canonical-33c96b1a2245fe1adb814968da0a74f0ecb81e72a9c285eef8bfad8e22b6f9f0)
- [api_specification.validation_custom_list.open_api_validation_rules.metadata](resources--http_loadbalancer--reference--group-009.md#canonical-d909cf4697742fedfea4bb0454328cfa4cf8ac5e1f4c09801bf45fdfb6e4fced)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-22e6f34ada48f9e0ca612543438be0c780be2fa4b66fc9fbcf55725af4faf50f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f24b9d13850bf04e7e751c169ff844430d00e274edecc543d620a6935fa29819"></a>

## api_specification.validation_custom_list.open_api_validation_rules.any_domain — api_specification.validation_custom_list.open_api_validation_rules.any_domain / f1aba5df9ce5 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- api_specification.validation_custom_list.open_api_validation_rules.any_domain

<a id="canonical-8ea84ed05d475dacbda9788dbddd53990398985a3d44bd442226898c941d12dd"></a>

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

<a id="canonical-b203c6ad583cc6facd0f87407f66bc9f9de6206010a97c1e42966bf7647e1a60"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.any_domain / f1aba5df9ce5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa7b58dee5ac96228fdb35de5494da1b0d007d5e1c81a1d87ce38dc35e1581a4"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.any_domain / f1aba5df9ce5 / 4

- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-33c96b1a2245fe1adb814968da0a74f0ecb81e72a9c285eef8bfad8e22b6f9f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf0498d16cb606911faeaaad8f4901fa24af3e1869ae1232c54537e38f49bfdb"></a>

## api_specification.validation_custom_list.open_api_validation_rules.api_endpoint — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 0596fd827e3e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

<a id="canonical-e644b12e0dd2eaf528d5a0d1536900647df7b0fc797a9aa456f2ffc9df566482"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-ef45d7255c80f338de744c80d168512ebb1cf9dc415d75a9cd23ccf725016355"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 0596fd827e3e / 3

<a id="canonical-c4e1202ac8573952be06906de51b1a5c600a1278cf09ef905be871ebe4a55fe7"></a>

<a id="canonical-8fcf9b487747afc8c417280e646cc0e243ca58ae0134f718f82ae3ff5a849764"></a>

## methods property — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 0596fd827e3e / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b8d94bbb65b577607ec358abea79563a820932ae6bb3b03ee2bc2118567aa172"></a>

<a id="canonical-61b095c3a9444d8003bc1a3ece7a2ada0dceff193cdf3293ffab5b4d7bb687a5"></a>

## path property — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 0596fd827e3e / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-9f9ee3078e1b25776c01c1756260ba980e7474ec152f08685c5abb0b0f63b572"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.api_endpoint / 0596fd827e3e / 6

- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-d909cf4697742fedfea4bb0454328cfa4cf8ac5e1f4c09801bf45fdfb6e4fced"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87cd66940f1331e571b4a9a6bf3e996b92b9d14ce0aad78530f6ef2ce98f4bd3"></a>

## api_specification.validation_custom_list.open_api_validation_rules.metadata — api_specification.validation_custom_list.open_api_validation_rules.metadata / f974c3b5ec65 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- api_specification.validation_custom_list.open_api_validation_rules.metadata

<a id="canonical-e9134f848c17054244f1c45010739495cc70d0b0ef56664e70c4db3d22fac1db"></a>

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

<a id="canonical-ce1ec058f3962ded810c346c92f6c90ec0bbe9ec2c87562f7254ecc7d9073c15"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.metadata / f974c3b5ec65 / 3

<a id="canonical-612e8c081dc62f4ba0afc4d1c680e166aed3e0e616a96644b7c71f231e8ba484"></a>

<a id="canonical-01880c723c50c202a4d26a7a08abd5546988f24ec184b67c14461a47bbd66b09"></a>

## description_spec property — api_specification.validation_custom_list.open_api_validation_rules.metadata / f974c3b5ec65 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-c7e2d5d1375b2a32646231339432dcf27e434a3e4aaba0a68219483356ecc0cc"></a>

<a id="canonical-07644634dc850e7997aa1d4353b070e9a44287ff5bc4bec4f0f73bb25da0cfb8"></a>

## name property — api_specification.validation_custom_list.open_api_validation_rules.metadata / f974c3b5ec65 / 5

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

<a id="canonical-53434aee5d603f21c2836924e927c05b35651c0fb0f5be3872886d22687159ec"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.metadata / f974c3b5ec65 / 6

- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c0d3235c0564013917377d11532d02e57fcf059f39c64419422a6f989d62049"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 477ac36b8232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode

<a id="canonical-24ee70283902caa41b5cdd633ebd041a4c23bbf15a206b6e6e7d4f17a179803b"></a>

Type: `"object"`. single nested block, Optional.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("response_validation_mode_active",
    "skip_response_validation"),
  validators.ConflictingObjectAttributes("skip_validation",
    "validation_mode_active")}
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
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

Terraform syntax:

```terraform
validation_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-7ba983825c76b3f6d82120cafcfd881221386711481ca7c697a13cacfbfb07f8"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 477ac36b8232 / 3

- [response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-75c242e8792a34997fdb1a2e4027abcf1fa562173886993eaad90294fe85033b): complete subsection reference.

- [skip_response_validation](resources--http_loadbalancer--reference--group-009.md#canonical-f41d0db68beb2ef7b9257597b1d3bc1274f58c4f7a11b3208da2f27f4cb758e0): complete subsection reference.

- [skip_validation](resources--http_loadbalancer--reference--group-009.md#canonical-1f7e92cc3815f4620a59ae8c24905f741d9cf5537134669954c049d812788dc0): complete subsection reference.

- [validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-dc6a3e18b1c100d758cfe54d98983bac9119a64fd12b057253c5e470994f9dd8): complete subsection reference.

<a id="canonical-6d98267d00413662a0f724799d55241ff4d8faa0cc48b57faea5801d4cea78ea"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 477ac36b8232 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-75c242e8792a34997fdb1a2e4027abcf1fa562173886993eaad90294fe85033b)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation](resources--http_loadbalancer--reference--group-009.md#canonical-f41d0db68beb2ef7b9257597b1d3bc1274f58c4f7a11b3208da2f27f4cb758e0)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation](resources--http_loadbalancer--reference--group-009.md#canonical-1f7e92cc3815f4620a59ae8c24905f741d9cf5537134669954c049d812788dc0)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-dc6a3e18b1c100d758cfe54d98983bac9119a64fd12b057253c5e470994f9dd8)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-75c242e8792a34997fdb1a2e4027abcf1fa562173886993eaad90294fe85033b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-188074b891a8fcb2c18504a0117329ac93a9fadda36e48034f138cbe616d40f0"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / db54ef65694b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="canonical-33fb0b7921422ac42d9ec640e136006448856b8d52299ab5d1156753005237fc"></a>

Type: `"object"`. single nested block, Optional.

Open API Validation Mode Active. Validation mode properties of response.

Upstream description:

Validation mode properties of response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
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
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
response_validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-03e5fdd1e2ed13e9fab140729d448a76fd218797cb3127d4e928fc2062acbcd0"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / db54ef65694b / 3

- [enforcement_block](resources--http_loadbalancer--reference--group-009.md#canonical-efe4e32a3129003285ec03c8d16d2c07166067757f96710fdf6189ca5839ba01): complete subsection reference.

- [enforcement_report](resources--http_loadbalancer--reference--group-009.md#canonical-767fa1fbeb609e9164e70e9b20fff3e95dcda94fb3fb0ad7bd4ef5be2da81138): complete subsection reference.

<a id="canonical-5dbd76c4b805fd38ef8c217aafea2781e540395c81babcbfd3a0c99b68fdeb8a"></a>

<a id="canonical-08b8765373ef9fc7f4cdeb31f466bb49482cebe0f5907671a19f421c8287a668"></a>

## response_validation_properties property — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / db54ef65694b / 4

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9552ebdd0bc2f778d0422d485d9cb9b770d645cb675fc988557f3cc34006021a"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / db54ef65694b / 5

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block](resources--http_loadbalancer--reference--group-009.md#canonical-efe4e32a3129003285ec03c8d16d2c07166067757f96710fdf6189ca5839ba01)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report](resources--http_loadbalancer--reference--group-009.md#canonical-767fa1fbeb609e9164e70e9b20fff3e95dcda94fb3fb0ad7bd4ef5be2da81138)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-efe4e32a3129003285ec03c8d16d2c07166067757f96710fdf6189ca5839ba01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a0993fbfe03cd41e34aaecabb485e5cd1b6550e13e153ae6d8a2f395e1cab69"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 2c549f37526f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-75c242e8792a34997fdb1a2e4027abcf1fa562173886993eaad90294fe85033b)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-ff0fa805046af67547d80f0c2778f08f0a385fdd374a74893478af4ae07ba22b"></a>

Type: `["object", {}]`. Optional.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
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
enforcement_block = {}
```

<a id="canonical-cecaddeb110b7829471176e7ad9e092b08064ac23d7922fadceff980410b5fad"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 2c549f37526f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bf9c4224aa188e589f2ed6444613fba7435ab6b64a3f8b035359a32fd084d0f0"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 2c549f37526f / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-75c242e8792a34997fdb1a2e4027abcf1fa562173886993eaad90294fe85033b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-767fa1fbeb609e9164e70e9b20fff3e95dcda94fb3fb0ad7bd4ef5be2da81138"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ced62fbbd0b0e304d0fb5debdaf9b6a1471aea666d1457f96aa96538372fe408"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / fe33dfd07879 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-75c242e8792a34997fdb1a2e4027abcf1fa562173886993eaad90294fe85033b)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-1896aaf488af709588e3d65e5c71b7687e3915e227a771ce59bc58360b3b7bae"></a>

Type: `["object", {}]`. Optional.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
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
enforcement_report = {}
```

<a id="canonical-98029e3b4c722441873dd9172f273a0824563ef8d7a73083d634c1d23fa488a8"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / fe33dfd07879 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5726fdc84687d9745adb443d2a18fa5130e5be120e16ae205958efcfb8a34e4f"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / fe33dfd07879 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--http_loadbalancer--reference--group-009.md#canonical-75c242e8792a34997fdb1a2e4027abcf1fa562173886993eaad90294fe85033b)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-f41d0db68beb2ef7b9257597b1d3bc1274f58c4f7a11b3208da2f27f4cb758e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-833e33c2d75d8aa76841009d28806cc695b5f9a878bce65af1f417923898d724"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / eba12222364b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation

<a id="canonical-8a2962eca1149475bc926a83ce079dabe012e580533b3e7033102767d4fc11a5"></a>

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
skip_response_validation = {}
```

<a id="canonical-282fa2e603cd29c5b996fdfb37c798a7353f1202be46a86bc0d8972127ce21fa"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / eba12222364b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dfe9d3bbfd0dfcfeaa69815bedd99b0da695db9fa8010a50e5018f408985d9d7"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / eba12222364b / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-1f7e92cc3815f4620a59ae8c24905f741d9cf5537134669954c049d812788dc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f733d4376b5c200015a383e12ec982d4200775505a81e0cbcd7d002db42f061d"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 731e708d4ab3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-94b4d5b45c140f447678a0a06e4e71718643f40b80e3f8675c3128b7dac4eb2f)
- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-e39f46ee17d4d8075f3d741ef024d557ae4090d0ab94a0b4063a8cce9d955c10)
- [api_specification.validation_custom_list](resources--http_loadbalancer--reference--group-009.md#canonical-d0b19f04e97ec7b3a15988db3291da92c5fec8ada37dc4e08e29ad4a8768bce4)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--http_loadbalancer--reference--group-009.md#canonical-871283c10740c5f020e9ec61c8e18d845ce5a2124b0cf332c843dc80bf985601)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation

<a id="canonical-c971246c74fd415875b4e771bc4360dbb70722b02cee9bdb7d16546931339a92"></a>

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
skip_validation = {}
```

<a id="canonical-35fc15c0154fdadc4b5023080113d0360659d310c1238cc75e739b0292ec3e65"></a>

## Direct properties — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 731e708d4ab3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ddf061cde7f206871928d38232eb82650b35b5c6c722506b2e95873a31e65746"></a>

## Next pages — api_specification.validation_custom_list.open_api_validation_rules.validation_mo / 731e708d4ab3 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--http_loadbalancer--reference--group-009.md#canonical-7092af96e2bbe1a0001f3798a1a4c0196c698be541608d84309a038923d194bf)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-7b45dee760877c1f305714c7dd9c6975c40a205aed3ea2fb9502895dc70ebd63)

<a id="canonical-dc6a3e18b1c100d758cfe54d98983bac9119a64fd12b057253c5e470994f9dd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
