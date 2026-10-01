---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-98fa9210460dd66423749f53f55eea96ea8dcaebc538f6aeb5df6149e8250ccf"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 498b804e2a84 / 3

- [certificates](resources--workload--reference--group-006.md#canonical-ebb089030cb8a06d067ea2f043a2476a1174f27325168a2b8089842aad8936bb): complete subsection reference.

- [no_mtls](resources--workload--reference--group-006.md#canonical-f28c47c08233e9ad30f2fe37d12f17ba5ffc0e7337d8d3d76a5729cc189c3fa2): complete subsection reference.

- [tls_config](resources--workload--reference--group-006.md#canonical-74863938e66c27a1f562a3e8650c94dd2fea60f9add0a05d8b6bbb9b78ba782b): complete subsection reference.

- [use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5): complete subsection reference.

<a id="canonical-6e484f3311256f30850dc18028aa27af64e4c145c4a9100a6f3db70f750e9fc2"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 498b804e2a84 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates](resources--workload--reference--group-006.md#canonical-ebb089030cb8a06d067ea2f043a2476a1174f27325168a2b8089842aad8936bb)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls](resources--workload--reference--group-006.md#canonical-f28c47c08233e9ad30f2fe37d12f17ba5ffc0e7337d8d3d76a5729cc189c3fa2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-74863938e66c27a1f562a3e8650c94dd2fea60f9add0a05d8b6bbb9b78ba782b)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ebb089030cb8a06d067ea2f043a2476a1174f27325168a2b8089842aad8936bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68e59a8b9cb8866d2a4238ef01bf0dd1743e731626c25ffe64b3cc54e44263ab"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 4c073faf9319 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-11829254e39cb41559c0dafd943368cce656d333e0dfcd273cd45d26a3803758"></a>

Type: `"object"`. list nested block, Optional.

Select one or more certificates with any domain names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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

Terraform syntax:

```terraform
certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-c7790909fd57e2ffb3e4c432e2d8197dcdb05aa2558f3498ce8207827ebce711"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 4c073faf9319 / 3

<a id="canonical-004610651a7d25c6bf5a0b4d80421650ef19db7cbeac9ebd5acadbf32412e76f"></a>

<a id="canonical-46f8790a6e2f17625432a1229080068aab741ab7b8106bc60e91a19b54cbb67b"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 4c073faf9319 / 4

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

<a id="canonical-45398564e0eb85e8349dcc743fc19169241aaca293701b22251ec4ec75ac2581"></a>

<a id="canonical-a039b2c64cf31771b2db2ad617b704f809bf15b461074a13bc5a0b039d15310c"></a>

## namespace property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 4c073faf9319 / 5

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

<a id="canonical-73114c38283afdecfa7bf89de8985a04de1c2b93f4e0c5c7928c678e40978909"></a>

<a id="canonical-de648979f7a06a3a66799a12807eed0be7de2fba0893ab22f766424a4ff07e74"></a>

## tenant property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 4c073faf9319 / 6

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

<a id="canonical-ab202d222e57d63cd300a670964cec773799a29179153ae8909476efc2a3e70b"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 4c073faf9319 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f28c47c08233e9ad30f2fe37d12f17ba5ffc0e7337d8d3d76a5729cc189c3fa2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4a852e6c0e7cb2315e21d3e1dd8a4ceeef65f198ebedca39c3a36a6df2521b5"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / f308f61921bf / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-fef6eb38792d3472135433a6242a00d2fbc6449e15e1ce5e2fefacad4a92e024"></a>

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

<a id="canonical-dc5b635c27a30b9ba408023e45e2b0f354c7bd0cf039d549e04a0113fd4dd427"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / f308f61921bf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9925dc30c69d72ef761177b2ca32ba85060c447564390925a2bd431565d5ab4"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / f308f61921bf / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-74863938e66c27a1f562a3e8650c94dd2fea60f9add0a05d8b6bbb9b78ba782b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89c7fd9cf6dffb2ba24e6a8b36414f0d8e4cc38eb13493701f5160efa57ed5c7"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 0ae46fb43319 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-2223394c67dac153e4dd6e6ae704077875004acf311e221904b633250c617156"></a>

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

<a id="canonical-3d96c80507791c8685ad29b50d058e8a19b8138db1e61d5aa97f8dbf632f018d"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 0ae46fb43319 / 3

- [custom_security](resources--workload--reference--group-006.md#canonical-c24f835bb3a46b5fa771a11f6f22b2dd7f7c39f38326e3d72784029abe4c9544): complete subsection reference.

- [default_security](resources--workload--reference--group-006.md#canonical-974ca828a984c4cc11cf880c742661792b1f6c222d27abaedd89a80d914e8425): complete subsection reference.

- [low_security](resources--workload--reference--group-006.md#canonical-77e36c6c7913b62002f98afa9f132d9ff693e210a5675129bb0c36297af32189): complete subsection reference.

- [medium_security](resources--workload--reference--group-006.md#canonical-6387873b72e1e7940e5d87a53b416ad602ee1a005a4c12f8a0be2a6e81e56516): complete subsection reference.

<a id="canonical-39cfbf61ba397d1694348786dff131a965cc726632d34222a43abae92610cd0c"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 0ae46fb43319 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](resources--workload--reference--group-006.md#canonical-c24f835bb3a46b5fa771a11f6f22b2dd7f7c39f38326e3d72784029abe4c9544)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security](resources--workload--reference--group-006.md#canonical-974ca828a984c4cc11cf880c742661792b1f6c222d27abaedd89a80d914e8425)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security](resources--workload--reference--group-006.md#canonical-77e36c6c7913b62002f98afa9f132d9ff693e210a5675129bb0c36297af32189)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](resources--workload--reference--group-006.md#canonical-6387873b72e1e7940e5d87a53b416ad602ee1a005a4c12f8a0be2a6e81e56516)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c24f835bb3a46b5fa771a11f6f22b2dd7f7c39f38326e3d72784029abe4c9544"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c9e5f58a513e9eeccb9735ab8c27f318b26692b840432eed30ade9c5604fb76"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 8331daa2489a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-74863938e66c27a1f562a3e8650c94dd2fea60f9add0a05d8b6bbb9b78ba782b)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-9bfc6a9e62d92c0fce2c3cc23bb3b72c863263efb56fcdd4eaf865b97ef50533"></a>

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

<a id="canonical-58c4703611e92b21937bab230f20fd6d187537a478a20c8a4b5090fa5bd24391"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 8331daa2489a / 3

<a id="canonical-6cf192bf6fa85831937b312372bc9b114c301ebcfa68d6117cd868ad3acc4f9f"></a>

<a id="canonical-1cecefb543413304ba49c2a0a297a1ce88ce0de11f08918880c1c88b215d364f"></a>

## cipher_suites property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 8331daa2489a / 4

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

<a id="canonical-00fc174d1f98af470bee2dd20f4343bcd088d0fcf649b854484200198871cc3a"></a>

<a id="canonical-be20733ce225069a92320595459c0c5ad5b14a17884853e4a035a8c1c0947c06"></a>

## max_version property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 8331daa2489a / 5

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

<a id="canonical-6fe73f64118f84066d671f8ac6b4b1974204eded3d59132d8dfd2c839a353967"></a>

<a id="canonical-6971a19d6807993d2fcefb3449a874573eb51d4cc0298ec3fd8da1ab88e95724"></a>

## min_version property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 8331daa2489a / 6

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

<a id="canonical-afd057d5baafe393a9a73893cdec9a13fefbab92a55a37cf6907ac1faad24cdf"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 8331daa2489a / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-74863938e66c27a1f562a3e8650c94dd2fea60f9add0a05d8b6bbb9b78ba782b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-974ca828a984c4cc11cf880c742661792b1f6c222d27abaedd89a80d914e8425"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f458a9683faee738d2f0ab9f5d4533ca4aeaca80a47d4bca19727f5d95a1cdb"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 280048d2a2f6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-74863938e66c27a1f562a3e8650c94dd2fea60f9add0a05d8b6bbb9b78ba782b)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-bc6088de51eeec17a426f1dbbf03955667c7e5e53f562e717fc203eccc7ecd85"></a>

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

<a id="canonical-9b537878114ebb5491ba38fff83abfece1be0f5541cbac36159a0599c793844d"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 280048d2a2f6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-518afa5d35ea66fa2370599d2daa96d846fc2f0f357a6b1e90c675bac046d9ba"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 280048d2a2f6 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-74863938e66c27a1f562a3e8650c94dd2fea60f9add0a05d8b6bbb9b78ba782b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-77e36c6c7913b62002f98afa9f132d9ff693e210a5675129bb0c36297af32189"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e701e51fa528a912fbb8380749866f44981087a737130e865b852488d991542e"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / c9c8b3ea3039 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-74863938e66c27a1f562a3e8650c94dd2fea60f9add0a05d8b6bbb9b78ba782b)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-dabc8fb045ee8499e79bb77ebf994973e3790651bb970d4be06cca9d09d5df09"></a>

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

<a id="canonical-24106478891b32b334eebc250ab29693d738b42a8d25d37209263becddd75d99"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / c9c8b3ea3039 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f327928a3597c28634de97d9401d850f9edd869375e2d8cdf1be6abc1da16d7"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / c9c8b3ea3039 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-74863938e66c27a1f562a3e8650c94dd2fea60f9add0a05d8b6bbb9b78ba782b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6387873b72e1e7940e5d87a53b416ad602ee1a005a4c12f8a0be2a6e81e56516"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e88524753dd33b2a4149f100458c2c8b2b539b5da1f5cb7f1e298a8fdf33a62"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 2477624abf2d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-74863938e66c27a1f562a3e8650c94dd2fea60f9add0a05d8b6bbb9b78ba782b)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-e07b653b01c1b9d8229c1ea41b3df409eedde15f6c7762581cc3452251721665"></a>

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

<a id="canonical-75f1158db8c03976d4b387fcbd33b9ddf30de34be53bd06b4b2f24ab162a5ac3"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 2477624abf2d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2787582247d191b6d71f2b8e3e893429aef40b21032ae661f35af1f141c381d5"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 2477624abf2d / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-006.md#canonical-74863938e66c27a1f562a3e8650c94dd2fea60f9add0a05d8b6bbb9b78ba782b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7846a80e5fd59b9aff8b952c78d127b4b7be78855e5e403a466fd29389bcae76"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 6c09dc778e7f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-f155e7955deaa41f10b040a1135babe01dc2265b8bb9a192bfa28ff5a82d816b"></a>

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

<a id="canonical-5bda2d90b9fc3244bccf9cac00bdd588ca07ef9367e9517c5ef3b1e4cc4a8f4e"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 6c09dc778e7f / 3

<a id="canonical-ebef93ddd4f51897b6bb541dc30fdb7872678886c90a20d055535e06662f665f"></a>

<a id="canonical-b2e6de27bd40855638dc60d6540c9275d88f2842ba7909f3396f2064b495ef75"></a>

## client_certificate_optional property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 6c09dc778e7f / 4

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

- [crl](resources--workload--reference--group-006.md#canonical-fce0599f1e9baef9d53c8fbcd826b9e35ceab17d18a8e0d8095681c294790061): complete subsection reference.

- [no_crl](resources--workload--reference--group-006.md#canonical-deb9f2870944a10f4b5ea7024bb19cdf4ce29a79cc53e73b940da762ad5e49af): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-006.md#canonical-92efce193a0fff6576387fc3e2b0c2efdd5747d400e5f6b6f3c5402e87081f77): complete subsection reference.

<a id="canonical-0e385aa590b8f32f816b987f64b4f602143d5e8790e94286c6bc4f9bfe109886"></a>

<a id="canonical-dcc869d00eed040db57471b4fd19171bb6bd6ca9a00a56f0bf513846c4fb3bdb"></a>

## trusted_ca_url property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 6c09dc778e7f / 5

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

- [xfcc_disabled](resources--workload--reference--group-006.md#canonical-bfb98eb99ce261d20e55dbdd8949c4de49e9905de5c467b46f09211e60220ba9): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-006.md#canonical-027a923e545c03e8ddf1e5f33d290e6069c89146f091ca82c974912e4d636a32): complete subsection reference.

<a id="canonical-629a313ac64e3f154c48d16c27b9e37d9146e7088eb8acef50c1862410f0296f"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 6c09dc778e7f / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl](resources--workload--reference--group-006.md#canonical-fce0599f1e9baef9d53c8fbcd826b9e35ceab17d18a8e0d8095681c294790061)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](resources--workload--reference--group-006.md#canonical-deb9f2870944a10f4b5ea7024bb19cdf4ce29a79cc53e73b940da762ad5e49af)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](resources--workload--reference--group-006.md#canonical-92efce193a0fff6576387fc3e2b0c2efdd5747d400e5f6b6f3c5402e87081f77)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](resources--workload--reference--group-006.md#canonical-bfb98eb99ce261d20e55dbdd8949c4de49e9905de5c467b46f09211e60220ba9)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](resources--workload--reference--group-006.md#canonical-027a923e545c03e8ddf1e5f33d290e6069c89146f091ca82c974912e4d636a32)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-fce0599f1e9baef9d53c8fbcd826b9e35ceab17d18a8e0d8095681c294790061"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86f103147c1cabef2981b73be159552a49d8beca377e927cae818b32657070eb"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 171d22cec86e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-49e0435555c9a6839b84aea572303f1943c75ccec3229d848c67c1c15c8ed859"></a>

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

<a id="canonical-2f8cf0ab0c7be14295b400d5e92ea517f44d3daf14eea62908aff4f0b151eae5"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 171d22cec86e / 3

<a id="canonical-7f3ad7710affd2cc689b405466b8f164d51c4b035f545add4652e7444793d576"></a>

<a id="canonical-70f6fb7a05354109af4e95bd76e5c7286c93b9c3471139a24959ab0962f79734"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 171d22cec86e / 4

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

<a id="canonical-20a055cb91501d643069037f616a05c270810ac46b24a69cbc8409623cefead2"></a>

<a id="canonical-8c7da3f4955694fc17d0d17247447a93d848f732ff890b163ad8f4ba8517c071"></a>

## namespace property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 171d22cec86e / 5

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

<a id="canonical-2bf358dd17e868003d40f82f9bd4b8cc5e08e3663e562d6994de1de6a4f1fef2"></a>

<a id="canonical-94882e9033730e9a3dd168e28d426e51505b0fc63f45e0dd37f80658a5be621e"></a>

## tenant property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 171d22cec86e / 6

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

<a id="canonical-5d6428d9898968bc0383945a571c62e0c9884305bd77a391ab61315d6b1ef393"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 171d22cec86e / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-deb9f2870944a10f4b5ea7024bb19cdf4ce29a79cc53e73b940da762ad5e49af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79271014795d54a6788f98c79c6df2bd89d5a10bde857011b73bce3ab6be409f"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 8b85458abb17 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-40558c06d0baa2426295fc65b0dc9034598fc8ec1a4a30155ea2f084e38f7794"></a>

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

<a id="canonical-296e96b60c5dafa5a12c3a6350e194f778a81aa6bc58be6f148dbdb97f2169f7"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 8b85458abb17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-db405230c205b8308c94672f779414e586a842271a5522839f696b1c2b338b26"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 8b85458abb17 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-92efce193a0fff6576387fc3e2b0c2efdd5747d400e5f6b6f3c5402e87081f77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43230e5e6c204235e92c34eefb30d4b9fbdbb5a9c6b6e20afa247f068bf5198f"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 719ccce1abc9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-6e329c331ffb408572e14d6161a3bb2c37c10abcb28b509fdd315f813f20d0d5"></a>

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

<a id="canonical-2d40889fedd1f094bcf87642c33327e294c20c6a18255f8bbf252dead60dc6a1"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 719ccce1abc9 / 3

<a id="canonical-c6d9c0012286e2162b5cb3308167854cd22834b1b59082329f0d8206937b30b0"></a>

<a id="canonical-4f49d8c85e75594b4dc9e35ca561d5927a8218559d4b221b474e207cea4a16b8"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 719ccce1abc9 / 4

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

<a id="canonical-d3c3f226200e1a2633a11f42b664187fb568df5d28d48663146052ddda32514c"></a>

<a id="canonical-3f3bea498af7cc76fcfeaefcb32bde384709c362bb7baaa52a13a261aa9ebd9c"></a>

## namespace property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 719ccce1abc9 / 5

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

<a id="canonical-1f70145afb407216e3a1a37f4e98a8571bbec2a94c2431c41969d956fe5ae136"></a>

<a id="canonical-f2dcae3cff8e81af7b5c59445ed3af29033746a2655b34d6e4f72c8f7dc60a26"></a>

## tenant property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 719ccce1abc9 / 6

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

<a id="canonical-bfc65c874625fc0191bacfda210440954ee8283aef42e6dbf6cd425ab6cc9fbc"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 719ccce1abc9 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bfb98eb99ce261d20e55dbdd8949c4de49e9905de5c467b46f09211e60220ba9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afb19b5ef13a9723451c84ba5f6bb1725e368e3c2f5cf454eed6bd8a7e32e6c0"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 173f4a45a61d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-2a7bff5904184b0d6c8467db00d8df8e9355d7056dab7304b8c073d85cbb30a8"></a>

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

<a id="canonical-38825591c52d1014256466ce30609b4efb4003e1f8eeacf341e4d91b4c78d636"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 173f4a45a61d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89d169ba6535b7bb3590dfab4b2ad9ed059cfd182819fb3317f843ab13d8d966"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / 173f4a45a61d / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-027a923e545c03e8ddf1e5f33d290e6069c89146f091ca82c974912e4d636a32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a5bae722126fb6c0e6bc3fe088c5feb751ebe2479f748ec64295c5e229122ab"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / d2e11645866b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-005.md#canonical-137ff8a95d257e7b9a269ba147dbae38b7fe843cc507f89319608cfafe000444)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-9c7c857533209d16b6fc32d3af02ae362ad1eed97f9941612dc5bbd7874a9eb8"></a>

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

<a id="canonical-a4d08ad4587d560e0b3b72fdd4fe990e6dff11074a51b7c4dd1b6cc686433a85"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / d2e11645866b / 3

<a id="canonical-9c72dd76aa1aec18aa153dc0a4c526d7bac72c0a7fff9df78f966ce5d254fa9a"></a>

<a id="canonical-42fca7e812165682d7e33c02b5231c6517d827792d93ea4bb635ab829a6dc99b"></a>

## xfcc_header_elements property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / d2e11645866b / 4

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

<a id="canonical-70de712f0550f9b9560f10c9012d56b83d85f8159b1129aa1224301449d5ec3d"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cer / d2e11645866b / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-006.md#canonical-51ff0618d86abbad6443d3fdc442ccb63626dcd2d65643e0f413facad8d6b6a5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13f6d808afdbf67273e131ba95700ec68101cde0fef5e14911212a2102a6d032"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / dabf1a59c01c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-3311cdacc9adc53909a6e435b238bf9c72f2edc254c40fcc173010948ec89250"></a>

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

<a id="canonical-76ecea8db228e1d9fa6090ae90e3382bbb5f2f1b7cafe0acbea4553c6f841126"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / dabf1a59c01c / 3

- [no_mtls](resources--workload--reference--group-006.md#canonical-6a4cb98643019624c4c39733887cc4c8067fd021cf4194f7c9882c818d8092e6): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa): complete subsection reference.

- [tls_config](resources--workload--reference--group-006.md#canonical-1d9d9284b2f5b6b53fb88aa2943153c1441beb8eac3f22e7120b8fae42718c2d): complete subsection reference.

- [use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b): complete subsection reference.

<a id="canonical-6413f970cf821b3adbe3c870119f0fb7771a4776c5aae50c95dc419b6094015e"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / dabf1a59c01c / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls](resources--workload--reference--group-006.md#canonical-6a4cb98643019624c4c39733887cc4c8067fd021cf4194f7c9882c818d8092e6)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-1d9d9284b2f5b6b53fb88aa2943153c1441beb8eac3f22e7120b8fae42718c2d)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6a4cb98643019624c4c39733887cc4c8067fd021cf4194f7c9882c818d8092e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2c648e5ef262acfc157c2f81198459b6d7a6cc10051adfe9972837b3a840b3e"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / f635c2bcbe9b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-76e41d6ab3fc0cb3b66135362a805a6e9f6d112d44ae99e57cb82e610478e369"></a>

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

<a id="canonical-9e7196013368aad131dfac0c7fb646a246c93569c4ed5c3bb93feb7297d5da29"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / f635c2bcbe9b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4874795f3d72f6309f66648156ad3fcd847223ceecde7c5910b888d938782a1"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / f635c2bcbe9b / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a57e8c35f2a305e8974cc7c034ac61e5fe55ef9ebadc8085a19565ea3f4547a"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 008337747809 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-73ab4dbd67542af7c6e18195c23691e2581900c4b7c6e5c4c77b40e00ab252bb"></a>

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

<a id="canonical-f780afa54adcc3980c07f6f45faf84a5e4a7949b15c4dd372eb929dcd39c83c4"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 008337747809 / 3

<a id="canonical-484d27dcf0ac1e9b3b6cd2949a022978f0cfcfeade35b77eaefb8735202646da"></a>

<a id="canonical-e60811bebbf297a858e34318ac901ec4cc1245a9d20ef613d62f076d1568eac8"></a>

## certificate_url property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 008337747809 / 4

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

- [custom_hash_algorithms](resources--workload--reference--group-006.md#canonical-ad539aa62155e59c070ce65086edaf486b39fa404db062bf7d6fc3608e6c3d21): complete subsection reference.

<a id="canonical-cb927407c53a101d7d8e352332a9048e2cf5b2d792a617a2b67269ab68d6cb70"></a>

<a id="canonical-541c613b92af1279348b5a8d85efe919cd0d9292fdaa0cbae9855d20729736d4"></a>

## description_spec property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 008337747809 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-006.md#canonical-fcde080b4163c57750880c55d14400156da6b242247b0e43163cae52c5ae5779): complete subsection reference.

- [private_key](resources--workload--reference--group-006.md#canonical-cd5cf0ea328d652fc7e0c5cfb8f06c4c3b5a47da7b7e9e80712ae6884bfb3ac0): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-006.md#canonical-dc812dca469672cf6e2bf5d8383d8c2306f0c5fd7d7b4b258aeaf438fb1ae1b4): complete subsection reference.

<a id="canonical-aed73929cc1e0c41ac0a6df104fe35d59625f73444787cb052e230c8ee288339"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 008337747809 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--workload--reference--group-006.md#canonical-ad539aa62155e59c070ce65086edaf486b39fa404db062bf7d6fc3608e6c3d21)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--workload--reference--group-006.md#canonical-fcde080b4163c57750880c55d14400156da6b242247b0e43163cae52c5ae5779)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-006.md#canonical-cd5cf0ea328d652fc7e0c5cfb8f06c4c3b5a47da7b7e9e80712ae6884bfb3ac0)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults](resources--workload--reference--group-006.md#canonical-dc812dca469672cf6e2bf5d8383d8c2306f0c5fd7d7b4b258aeaf438fb1ae1b4)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ad539aa62155e59c070ce65086edaf486b39fa404db062bf7d6fc3608e6c3d21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98ce36c8146407de66534af78e204318091a90ae33ef721bbbff202e2a444846"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / a07993f47031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-d4bda26c4bb23197a70c5c8ea71b17da1939602fbb1f2b6e07f07f4a9e762a0b"></a>

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

<a id="canonical-0c32b5e824952405003379264019558652f8e355966110413204244279222c1c"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / a07993f47031 / 3

<a id="canonical-ed82c1763346c63c1f2b69ef1f24eb4d3e048be1fa126b622f0599c039221eba"></a>

<a id="canonical-4a7404cd5f73b695412291d5bc479a24e3ea80d5ac41fba8daf6c7d5a03a1695"></a>

## hash_algorithms property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / a07993f47031 / 4

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

<a id="canonical-57a39a94ce07d27702850cbb77e547ba1386d54f5dfeb94b0c348034055b4ca6"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / a07993f47031 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-fcde080b4163c57750880c55d14400156da6b242247b0e43163cae52c5ae5779"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae0b7a55ce10d7fe506afd2451495bf6fc6b51fd74a34307d8247e1419fb66ef"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / a21823364a37 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-6e48228e89bb6effe70064a8adf82a38d8084dc158ed56103345ddb9f14144a9"></a>

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

<a id="canonical-3684f519ec455a4a996ac1cf643077743e9e772433833c113d7ad7f2ceacb2ce"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / a21823364a37 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8b6d16ea0daa3c671dd3d01abfb1f971b85b6a5589651de8633fc6cf1f250d05"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / a21823364a37 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cd5cf0ea328d652fc7e0c5cfb8f06c4c3b5a47da7b7e9e80712ae6884bfb3ac0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70db8584d7d0d9114e5544a9ee55d775889d0804cfa4c39ef1924ec146c4b588"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 2066cb07e77e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-cef5e57ed46f1c73bc27bd134fe5a582e83b954b0e4344b3af5a8751bd008398"></a>

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

<a id="canonical-d6d9e2f1136b37868afc535f022c891ff60206f0f9f622a5a2f103dc89a66027"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 2066cb07e77e / 3

- [blindfold_secret_info](resources--workload--reference--group-006.md#canonical-de4ad30efcd61c0200b62dc581b9cd11cc1332aaa2eeb5e3d3534db96c20fc7e): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-006.md#canonical-dc7307de996974b1344682ac103e62828b211d582d79c3e347d5ffb283593782): complete subsection reference.

<a id="canonical-ee286f502d33f0921cdbe1512e3dcd8444a44e52da94c22ebc2788fb63853a2d"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 2066cb07e77e / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--workload--reference--group-006.md#canonical-de4ad30efcd61c0200b62dc581b9cd11cc1332aaa2eeb5e3d3534db96c20fc7e)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--workload--reference--group-006.md#canonical-dc7307de996974b1344682ac103e62828b211d582d79c3e347d5ffb283593782)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-de4ad30efcd61c0200b62dc581b9cd11cc1332aaa2eeb5e3d3534db96c20fc7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6b9382120572c9e46bffada2e72f4291796a6cd07581dacf68521e712ddf3d2"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 78414bfa19fb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-006.md#canonical-cd5cf0ea328d652fc7e0c5cfb8f06c4c3b5a47da7b7e9e80712ae6884bfb3ac0)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-e1c56be8be0e415c05256565c2544ac35c1d472e8dfdf5ea8fe80dc2ab32f77d"></a>

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

<a id="canonical-32e12939a5d635fcb421325de049fc2f36f0178b741e1a7acaeb3ded0215bbcb"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 78414bfa19fb / 3

<a id="canonical-c4e0575030d68c863c1ccc660d1d7fadadac8c68958fbe1ca6df0b993359885a"></a>

<a id="canonical-9e852753fe7a88d58ea69f7fc1bb5d653e25b50ca0979e06be36c8b9d9a25795"></a>

## decryption_provider property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 78414bfa19fb / 4

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

<a id="canonical-41ac6dcc727bb7d0b2e341e8cfd9702f7ffeb0ba020053dabf912be0ff1b6467"></a>

<a id="canonical-27f3c5bd53e81ddfc8a3d624447f3798b8a1ea47187f574819b1667356d897a6"></a>

## location property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 78414bfa19fb / 5

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

<a id="canonical-9e5623e9e4af9cf345d5e29678c6b6d05c56fd8795846d887bfa711de38d60b7"></a>

<a id="canonical-3c321ce07c1cd2f76f2207b4c73022067398f6e7cc688efa1e32548dffc84a7f"></a>

## store_provider property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 78414bfa19fb / 6

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

<a id="canonical-b09ff8c154c073e8b2df1807e779dd8847a0eb3124a9ed39567b45ddee0692ea"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 78414bfa19fb / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-006.md#canonical-cd5cf0ea328d652fc7e0c5cfb8f06c4c3b5a47da7b7e9e80712ae6884bfb3ac0)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-dc7307de996974b1344682ac103e62828b211d582d79c3e347d5ffb283593782"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-730621a9eeca634567f4f30c5d2b8c284425a1149784deba3e96f19d07d1a3f6"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / e35600aa99e6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-006.md#canonical-cd5cf0ea328d652fc7e0c5cfb8f06c4c3b5a47da7b7e9e80712ae6884bfb3ac0)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-f506527a3cc472536b6954517b38e5250b279c4e8dca0e0fa2f1c721f06e4318"></a>

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

<a id="canonical-e72c3a2c5dfc377c5e01574aabc8d2e8c9baa890ae25f5dc530b0b8e2f618183"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / e35600aa99e6 / 3

<a id="canonical-023155961ec31d6a136257de08da9b5c1caa2c60aebe7aba957d75814a530cf8"></a>

<a id="canonical-784f8805ff8407fc5f9c98366b1e55681cf064898466733ce4a96d7f4bfcbab0"></a>

## provider_ref property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / e35600aa99e6 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-49978900add092daa814562358b66b0254946678b2d18d1ad28d9801150a2f88"></a>

<a id="canonical-dbcb0fc0209a5f4768c8d91c7edacaefe9327afe5f985f786d79ca8917218fa9"></a>

## url property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / e35600aa99e6 / 5

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

<a id="canonical-41e003ada58f3d5ab2c469f1ac1f2728a775d661ce556b0bbfca25bad9fff96f"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / e35600aa99e6 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-006.md#canonical-cd5cf0ea328d652fc7e0c5cfb8f06c4c3b5a47da7b7e9e80712ae6884bfb3ac0)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-dc812dca469672cf6e2bf5d8383d8c2306f0c5fd7d7b4b258aeaf438fb1ae1b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38b00e9fbe32bfc7dc609f2176456d5da1e2f609327e0e893a7662273556f663"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / e45ddb1ee47a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-26d897eeda4dfa572056205d21ae7b7dab799bee6cab8ee5333d46514b99e394"></a>

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

<a id="canonical-99fd6d7b356c9a9d319975b8d9d9fb41d4d7e8462c92c9257eef286dc35df5fd"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / e45ddb1ee47a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a7b58e4c603fef95574c40f02158b5699769dae00d5622f32b891ede95c45c23"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / e45ddb1ee47a / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-006.md#canonical-8b322df898af62eb85dc9d4b4e23465de3212e05d2b7bb84b8f85f3692e002fa)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1d9d9284b2f5b6b53fb88aa2943153c1441beb8eac3f22e7120b8fae42718c2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4e93e4a55ed435d1d50a3e18ac9902ff764dc6e76f04aa0d58549070ab7f71d"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 4cc54cf9da97 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-151a51a3b99dec04a9db8a453565ae3d7c42c351a472ead3ba3862bb7baecfb6"></a>

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

<a id="canonical-eeb2936c1f3bf152b2a032dc506dfd1a72359388c5bdbf5f4d4879cc713ef230"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 4cc54cf9da97 / 3

- [custom_security](resources--workload--reference--group-006.md#canonical-f09fc590348faa2d0f8032b750a62a007d58a5b608a24c3db55628ca40738a47): complete subsection reference.

- [default_security](resources--workload--reference--group-006.md#canonical-85bf9e5ce1e9ab0e023848f737ba93e580b2fdf87dcc668e84c3d047e8a90285): complete subsection reference.

- [low_security](resources--workload--reference--group-006.md#canonical-86b993ee5c168165beceec5920ed9dc3fef93d3c7e2109313a460fd64e09b954): complete subsection reference.

- [medium_security](resources--workload--reference--group-006.md#canonical-03e2e23070b5b44403e17251abe4e987533904ac808fa55d71723b8ddab22cf2): complete subsection reference.

<a id="canonical-2f7e60aa770c0044e6ac404a2dc1dd7c273c0152a63ddf62da4d45a332d4ed15"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 4cc54cf9da97 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security](resources--workload--reference--group-006.md#canonical-f09fc590348faa2d0f8032b750a62a007d58a5b608a24c3db55628ca40738a47)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security](resources--workload--reference--group-006.md#canonical-85bf9e5ce1e9ab0e023848f737ba93e580b2fdf87dcc668e84c3d047e8a90285)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security](resources--workload--reference--group-006.md#canonical-86b993ee5c168165beceec5920ed9dc3fef93d3c7e2109313a460fd64e09b954)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security](resources--workload--reference--group-006.md#canonical-03e2e23070b5b44403e17251abe4e987533904ac808fa55d71723b8ddab22cf2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f09fc590348faa2d0f8032b750a62a007d58a5b608a24c3db55628ca40738a47"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3922743a6f9325724fe8738523da7cbf983d7b3c7c6501530f1fd4cf002fda1"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 233088653af5 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-1d9d9284b2f5b6b53fb88aa2943153c1441beb8eac3f22e7120b8fae42718c2d)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-544c0a6b99910694c17e8019de2c1871f3fac83f46c08b80038fda87bccc5239"></a>

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

<a id="canonical-4059865790d5da43ece86c29d73a457cfde8c4899924a1dd190ad2cf7b8f7181"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 233088653af5 / 3

<a id="canonical-5616b3505503071fd9242154e6b6c7865602eea889ca0951051aec32742ae807"></a>

<a id="canonical-b6a8f089fcfb2ed6fc8421101a4b20f6f13523eaab2f6b7f4d4aa84d4b037a5d"></a>

## cipher_suites property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 233088653af5 / 4

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

<a id="canonical-2abe54f18b25d5295d7a5053ef98fe00119a5f297bbef2b320af4e4f3815adb5"></a>

<a id="canonical-0a8a43dc41edfaa4649f0916d0532b77e98b594c337ddace28f9b59ad94c7cc5"></a>

## max_version property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 233088653af5 / 5

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

<a id="canonical-6459a8934c696d2f68f17c6e62a97916a747d82e342b9e2cb39a11a7a1b78e3a"></a>

<a id="canonical-91218ebdc5a03ac821dcf799c0894ddd884ccb6b304e91d8287dbd07c06bf496"></a>

## min_version property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 233088653af5 / 6

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

<a id="canonical-9c249e658a29daa46c27166e972360112262a44e7196955fd4dd5cb6c7e93281"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 233088653af5 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-1d9d9284b2f5b6b53fb88aa2943153c1441beb8eac3f22e7120b8fae42718c2d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-85bf9e5ce1e9ab0e023848f737ba93e580b2fdf87dcc668e84c3d047e8a90285"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13bb3c30494c04f27edb0ea79112cb128f6fec12f2903b00f6f080ddcb2f7697"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 384a56ba045f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-1d9d9284b2f5b6b53fb88aa2943153c1441beb8eac3f22e7120b8fae42718c2d)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-e56502d2e34d2bb160a7c6009fd4da0d043a6df530697f4ac5e88dc146706276"></a>

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

<a id="canonical-cf045ba573a7aa5042601f4b98f6855cc4b9aa03cf3df26a3e380fb4991bcafb"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 384a56ba045f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c0479c0de2cbfcb2fa82d98893975f7ad7c9a1db2dff0a332238b1e085c2fb24"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 384a56ba045f / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-1d9d9284b2f5b6b53fb88aa2943153c1441beb8eac3f22e7120b8fae42718c2d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-86b993ee5c168165beceec5920ed9dc3fef93d3c7e2109313a460fd64e09b954"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c2c272d9ea60d6d2c8c016fbc92bf8fe7ff399a6d9928ee88c9a33bb3ea83a9"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / d03db1e98520 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-1d9d9284b2f5b6b53fb88aa2943153c1441beb8eac3f22e7120b8fae42718c2d)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-b2ac7ead5a80403a7748df5b39cdb4a3ebc693ed8621034798813372bcb79e44"></a>

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

<a id="canonical-3ec880b9de1138e5a43f0c307b2561455cddb32ea84d3db0a9c907d41ad3b10b"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / d03db1e98520 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7f691e5a281519ec31f15b9659e2d19c77971c606407ae32c75355cc8814fa0f"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / d03db1e98520 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-1d9d9284b2f5b6b53fb88aa2943153c1441beb8eac3f22e7120b8fae42718c2d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-03e2e23070b5b44403e17251abe4e987533904ac808fa55d71723b8ddab22cf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-661c5ba01a474f9235e32823c50b1a4356d67659655bd9013f6ea887b3fe2068"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 7046092042d2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-1d9d9284b2f5b6b53fb88aa2943153c1441beb8eac3f22e7120b8fae42718c2d)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-84ad163e98f4b7f596e33911dfba9fd9e1cbf5556532743d9000af048560251a"></a>

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

<a id="canonical-bbd580d3909344264be87db27ad3a987b56c275ed28fe113ae11ba348d117791"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 7046092042d2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f24efa10683f72aab9b2af8dee103492e3c575b96e80fd10f9737ee9c2c00b24"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 7046092042d2 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-006.md#canonical-1d9d9284b2f5b6b53fb88aa2943153c1441beb8eac3f22e7120b8fae42718c2d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf30989f9fa148e3dca1d8da5f8180bf07d7b519a15037939e0430d6367543a3"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 09611e41ab40 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-591a3164ff8b05e1063504f3a1c41a9b470932a54d54bd0c6bd578c8db8fdf27"></a>

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

<a id="canonical-9240d2fb30c3f134ee4d29a67cb96208ec6cf19a3d9d29a07f34a16a35bcddb4"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 09611e41ab40 / 3

<a id="canonical-4d28561a177deb78690fb517db6e63ae869cad8ebdc71e9fb48f9a4eaeb53051"></a>

<a id="canonical-a9d6223f9b1ec04239687e81c652ea7a1f67981724b04744b07d8a33f8248d1f"></a>

## client_certificate_optional property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 09611e41ab40 / 4

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

- [crl](resources--workload--reference--group-006.md#canonical-f0d6a23f4c338b4cf13770fc6099ad1424868557faa3410da501db71856c108b): complete subsection reference.

- [no_crl](resources--workload--reference--group-006.md#canonical-11347524f196175ebe1e569b631208ac326b5b5d00cb31634e7d7076525487dd): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-006.md#canonical-c625f81e19bfd7bb811ded4373247ab12a1f66d1057ae49de991f6ca30500619): complete subsection reference.

<a id="canonical-e61455dfa81b9e6b832134828c89e8ea9c3e8dc7e0e3334049c00f3eddcd6c06"></a>

<a id="canonical-8528c7e9f4b81e3ed204223fbbb75774a806c095b3cd33abf5c5c7563d3e8a35"></a>

## trusted_ca_url property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 09611e41ab40 / 5

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

- [xfcc_disabled](resources--workload--reference--group-006.md#canonical-d15651433a5dafdd78f80c7b89041383b23371bb0fe8c05b20199e046faa8283): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-006.md#canonical-82495ac75c713fb50702ad4ffb5a0961de816b58290eda42fd4d5c2bea2a5333): complete subsection reference.

<a id="canonical-663bde6bf68c90b736114ca1f6fe4c9b91ddd8b65d923e8450c216cfd1bf02f0"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 09611e41ab40 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl](resources--workload--reference--group-006.md#canonical-f0d6a23f4c338b4cf13770fc6099ad1424868557faa3410da501db71856c108b)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](resources--workload--reference--group-006.md#canonical-11347524f196175ebe1e569b631208ac326b5b5d00cb31634e7d7076525487dd)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](resources--workload--reference--group-006.md#canonical-c625f81e19bfd7bb811ded4373247ab12a1f66d1057ae49de991f6ca30500619)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](resources--workload--reference--group-006.md#canonical-d15651433a5dafdd78f80c7b89041383b23371bb0fe8c05b20199e046faa8283)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](resources--workload--reference--group-006.md#canonical-82495ac75c713fb50702ad4ffb5a0961de816b58290eda42fd4d5c2bea2a5333)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f0d6a23f4c338b4cf13770fc6099ad1424868557faa3410da501db71856c108b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e87c7895d0b0692fd7626b21ca280da4b37c8143c8ae40fb30e5ae41df7ee8b"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0824ea155f06 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-f7822eb294ff0e1cc515ad8cad59c23bdd126498d7fbd66da157f2510dadb86a"></a>

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

<a id="canonical-ff18eea84505083ac7228930eece0f5acb4e2a78bed2708e3fc14cc083b4a36c"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0824ea155f06 / 3

<a id="canonical-8811a12a40d7918be90680c713c3fae3114805a0299c2c98ec51255e8655ec73"></a>

<a id="canonical-9a48363ee229b9612a7ff7262fff4c9f27a2fd9a554396f3d507fb83dd4c6965"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0824ea155f06 / 4

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

<a id="canonical-55a1c7ba62b55c1265a1c8c0c7b4fd4d2e0352cac4c6565231df04eec29747f3"></a>

<a id="canonical-1dceb97d40d1ffa3d13c834d78a11758e468e9fcecd7f9239a74df0d4cbf340f"></a>

## namespace property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0824ea155f06 / 5

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

<a id="canonical-a78de28259fc7860ad79360147064f9a33a95ff4c1110eae95b217ca90941cb1"></a>

<a id="canonical-84a1145d1ab7b36553076e7990c45a08e9bc12960c955bbd7510dae230e31ee9"></a>

## tenant property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0824ea155f06 / 6

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

<a id="canonical-916e01f5a21c18333aaa3cd17efa3899b69548e7de42a65b04d31e2bcfed6956"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0824ea155f06 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-11347524f196175ebe1e569b631208ac326b5b5d00cb31634e7d7076525487dd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-458af2871e679615c085ea3de2ef6798029d6fa061afd80fe10a9f2bbdc988b7"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / b5635055dede / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-029cbb006ef23eb44677220fb3c443c0a174035b23ce4cc77031f63548fab801"></a>

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

<a id="canonical-602f4eaceb41013f15193d45b8b09f46ef6eab358f2fb6604f9b7ef61f21fd6d"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / b5635055dede / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0dc7007db361cf7d9aa0529870a5416576fdc92db2e5d64935e98a6fd97f6706"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / b5635055dede / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c625f81e19bfd7bb811ded4373247ab12a1f66d1057ae49de991f6ca30500619"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9339cfd9063206a68771373a56891d432b71b14b912bfc62dfb7746e3f40626"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0d8bd7145a08 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-2caf51aba8ae69ad04f24c5fc5ab55aa6f512710d238623e790484fcdbe78b4e"></a>

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

<a id="canonical-af2fb6109193b99631570dc40f23c699f3377ba15ecc51edca5c55a4311d5988"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0d8bd7145a08 / 3

<a id="canonical-85a4186aafb4d8a5d64c7aa2ba364ddd67e53ef05236024716e22f4578689d7b"></a>

<a id="canonical-acafd6860b649a9230565b82e3b44833111afc1ed7aa62eb7b81cdb5c0d1a145"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0d8bd7145a08 / 4

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

<a id="canonical-52dbf8368f4ac1daa2013ede739198b0e19d04cf73a131696ea181bd17a0b4e1"></a>

<a id="canonical-0a18420b1d4b10654847f1cfecd7d07ed5f8aed5d684d5062cef344bec33d1ce"></a>

## namespace property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0d8bd7145a08 / 5

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

<a id="canonical-6e152931248ade095eea969f4897273d751eaabb64e8b7c8e1bc833889f9732a"></a>

<a id="canonical-799e122105bdaa2a7b8dd904bfb8695c3a61c0e9c144a1c7638c23c184306983"></a>

## tenant property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0d8bd7145a08 / 6

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

<a id="canonical-1b477694cd89bd9c76ff867e6444277148b2e68bb6d99bd5232b1f91c60b1b8d"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 0d8bd7145a08 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d15651433a5dafdd78f80c7b89041383b23371bb0fe8c05b20199e046faa8283"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2c037af904c51c2680ddc8d6a42e697e562c9c74470311695fe2b298e4a6c4d"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / e0d2c32f25ad / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-a2592c9e3e5f71bf18f0874c9dfbbc500efbe85015b8530240b671103e9bd970"></a>

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

<a id="canonical-d487ab817d7997e978cee1542ea78a14de2afc81285e1dab91fe0d20f296c6fa"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / e0d2c32f25ad / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9520302c8290bbcb5885f4d70483c81f6ff9b759de06b16260816b851e13ace2"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / e0d2c32f25ad / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-82495ac75c713fb50702ad4ffb5a0961de816b58290eda42fd4d5c2bea2a5333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e6d4f31a0a2b5136a07469b214223d35b426c6bf729159d53e1fb0f1484ebb8"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 8bf1b7feb8cf / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-bc34c468e413b25ffab5c447517524bed929c0b96f8107be27a2c70f14bcef45)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-cbc48f3d69bdae68858971c3bd0f7ef81a2402af198b106a45dd7171932f55b8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-cf65971131f6d5819c60e2bd2e5a074c41a817ee0280de61792ce641e22b9e5d"></a>

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

<a id="canonical-c8a65f3198e5f2bbf8dd47877dcc115b4635b1180f8309a2c6c56f7a8e0360bc"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 8bf1b7feb8cf / 3

<a id="canonical-c64fef3d7b7e1ca66d139b9e1d5542fa90a038b91cb983e595f475934d6db5de"></a>

<a id="canonical-ecdef7d0811f42e0f78c4a65b9cfcc182487ae2e0d7fd1c1b4c74fbdca6b3840"></a>

## xfcc_header_elements property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 8bf1b7feb8cf / 4

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

<a id="canonical-b0241fd9c28f8330b7153fe49531dc1a198959f135757e0e28ab0dfdcb9f2d3c"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_par / 8bf1b7feb8cf / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-616ca0801a2cf1b6241e859f68715f67392ac1e2e761bcc142c10d7c4b0b6a1b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b026d0c2b0553e3b908370475ec2e980188b8725459a88ede20121ea5a0e7095"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bd64f81c6992 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert

<a id="canonical-5bfe73533baab5eb7015adea38cbd920aa9fba7644aafef473254707d03970db"></a>

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

<a id="canonical-ce5b3e2990b548229aa5170e1499e5da9114fe236649dbe47887fd66cb7c4fc7"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bd64f81c6992 / 3

<a id="canonical-54a1889a5b0bb2392451a6240a5811fa62992e0935aee72f5597badf33fcbf7c"></a>

<a id="canonical-69169a54cdec403226f6db4f3c344ba7a923b1f1cdf4b31f960181632d69cd62"></a>

## add_hsts property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bd64f81c6992 / 4

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

<a id="canonical-5fb1ac707aa08a6a23fc5e23e95bba8c7fd70e486a4024ff4633cd6bbd327154"></a>

<a id="canonical-60bb314c019458caaf7ba7321a97bf0e304c1b299293e83ab0810da1d460b263"></a>

## append_server_name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bd64f81c6992 / 5

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

- [coalescing_options](resources--workload--reference--group-007.md#canonical-61816e53c7b194ba062a72aa4e2fa3f7cdb21875cffa5af4027ea24d7ce11c9c): complete subsection reference.

<a id="canonical-60bfe7372aa70c014d68eccfdf7a8bb2b98f0815c4e43e00f0e99f78b915f12c"></a>
