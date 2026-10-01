---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-5159cc719fd9d30e6188ce76c4f86d15dbd93c426e384835ac34e1e8558690d8"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 97fe0d14d91d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-d75bbc08a329c9f7cd5c839fc62f228f0108416a52de37131bdf98c115095e77"></a>

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

<a id="canonical-ef7b2880257724c055b7dcb22ab19232dbaeb27c40a5b2e091e5f056caafaa70"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 97fe0d14d91d / 3

<a id="canonical-4ae553415dcb8bdc6e9b2dca724ce7caa072499b9eb6c4d80a895675987b865a"></a>

<a id="canonical-ac874157468d2b398227a5abbc4afca0bcef7f2abef0d368f50f15af33256257"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 97fe0d14d91d / 4

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

<a id="canonical-7b17faae8b2f4c71d459273cbc3618c1cbf3bf18ae41049ac6ceb500de35ba13"></a>

<a id="canonical-cdbffb1d00ad67e99372ca4f73a0c01a1631fa80dc4df4044bb1c3b9d2a46838"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 97fe0d14d91d / 5

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

<a id="canonical-5091643e9575106c8d88184fc1e4efd46020c9653c76ecd77fa014a88e2ff8c7"></a>

<a id="canonical-511924316373cc24fdf794642452427c5508ed679ef21c8d79b7565c33ceb7fe"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 97fe0d14d91d / 6

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

<a id="canonical-fdd1cac3b678f40f9b427178797398a3bcdd21c2c2678196ac11295304d81f1f"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 97fe0d14d91d / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1df4847bf6be57a4006e0010cae62f49cbd97b40062804ae95c7976517a39402"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6550c68e9dfade2d4dd717e9d2aa815eeae92c049dfa7a082a85887aff12e244"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 0485e9aaf722 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-ebb959bca1636e9b21c20ed7c9319b8e0f8be5a66b50cc95f2937faa37bdb7b8"></a>

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

<a id="canonical-20da674b218c49c4d6bc52f916b27a4abb4500d75dc14a753bc2cd743e79fdba"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 0485e9aaf722 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a7065daad1b2279606f6a15c262565be25be291e74a256909a912b9354dc1d88"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 0485e9aaf722 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-34bdf188ecf8ecfaf9609b4979bffb1166a5e4cb989a0e58e8b47c383e2ac65f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d79931badeea541fcfea483d5be1c08a606bed06977e264350ad449b52709a9b"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c642f0327b7 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-4333904fb15ac5c5e9d783d236ecde48392d64cab5afff34c8bfd9ccdaa9a3e7"></a>

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

<a id="canonical-d4c2c91a55359d14dfa1a727c5aa7f8cb08ab9608a7107148a3beec18d325d82"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c642f0327b7 / 3

<a id="canonical-780a894a717c2381db1676d2dcf73e90c36035e1857fc52fc0fdb6c7bd9e6cce"></a>

<a id="canonical-d649a371eb10926bc643e05ef8ef91f6afba46c2a5347f089ceaa751a4559c8e"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c642f0327b7 / 4

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

<a id="canonical-404e529a9e6a4b5422b780b4d7f8b0272d8c3a90befc6629f601f7b2688c8ce3"></a>

<a id="canonical-134a882706ad6b0e823cfa942dea928a5fb3151c355dccc885dcf4c72a652441"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c642f0327b7 / 5

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

<a id="canonical-da4f21cba915d033cfd60965cf5ddbd8d9c168f9fc0ec6b9f737b4a670238d23"></a>

<a id="canonical-6971aed70045f5b4dbcab9912b27088df3b443a00a82cc628500483be385b5a8"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c642f0327b7 / 6

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

<a id="canonical-8cdb0a629ce3429218ab87dd3c320840d82bc987740f04adbc3efea8eb6d8f8a"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c642f0327b7 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1fac8e48a8f05439c39a9dc94fda3a5a57e1d718c809df562b8a69fef74f0913"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1515e222684252c316de124a71fbae67b0e97fa28b6330bd66c3eaba3c3cbd5"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c361cf71cc54 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-50da393414fa2e418768905ee002b485f19193e130fa6b43e3188323fc0988d7"></a>

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

<a id="canonical-14c6832d85f8f2544ee5c911b522b7d422ed13d425c14eb4bbc244f79597aac6"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c361cf71cc54 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b4ed5e4c7db7c9f2674e6b6e5e1f6ef370614fb2a7965ce7a26d6168751e8e56"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / c361cf71cc54 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8dfda5cf4710875fa5ec033f08a74d62d3765c9d97905d178c4568e298951e58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a72f1d1c35d6f1f5042049a5b8eb2e8197c23a54eda3b4ee1f0f9476a727c071"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 13a2f9e543f2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-021.md#canonical-da3edf94cf6aba06d2555eab5614e19c8a42f009f76075a21c3a6ae6723cac08)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-aca8d17f152b58d85f24aa3ebd01b3c329549727b183dcb4bfc311438e7082f8"></a>

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

<a id="canonical-81a2dd3dbb09ac32d4b39a1ba9f9a7ba65a722bd1f1087d8858ce2c4654b5309"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 13a2f9e543f2 / 3

<a id="canonical-eaaec5b2a9fb70bf00a99d6fbfb3bce6d6d7e9e7181405ef4b0342a030d3563f"></a>

<a id="canonical-d23b13d78d07d771f063b87c72d54f3d909912ac729bdcaa066f55179678ac87"></a>

## xfcc_header_elements property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 13a2f9e543f2 / 4

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

<a id="canonical-d4125d993253a8cca3d849148a7659fd547363670d66de88348b2733624384a3"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 13a2f9e543f2 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-021.md#canonical-293c0031367440e573dd9785d10bcea9ef721f077f3d2ed6a72fb8da519171a1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-096fa76ced4e31aea30ef982afc5f894a8f54dd832497e7328007653a6246e7b"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 22b563ed3bc2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-73a5e159a2e15fc7f1b81f142e42c0f5cd6ccb041850e6290018c71f2d98c848"></a>

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

<a id="canonical-555eb6cce67671a3e201057cff5c70d98fee64c0787c34f81990b31d4db846b7"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 22b563ed3bc2 / 3

- [no_mtls](resources--workload--reference--group-022.md#canonical-c8cc5e184370e03197633f3ee32b60598c361f940d0c7fa464dad59b27498583): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9): complete subsection reference.

- [tls_config](resources--workload--reference--group-022.md#canonical-3b01f26e3af33a2e8c89a4d59489ac025852c5f2ae3e88a7a80184fbcfbee608): complete subsection reference.

- [use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21): complete subsection reference.

<a id="canonical-04a8095b1642122b8cb5380f5df9f540ebdfce84806df848c0d68fd175a54d5f"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 22b563ed3bc2 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls](resources--workload--reference--group-022.md#canonical-c8cc5e184370e03197633f3ee32b60598c361f940d0c7fa464dad59b27498583)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-022.md#canonical-3b01f26e3af33a2e8c89a4d59489ac025852c5f2ae3e88a7a80184fbcfbee608)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c8cc5e184370e03197633f3ee32b60598c361f940d0c7fa464dad59b27498583"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba2d9137e2131cd631953fc91d925fffba687c191d4e295aaf3cb9833ee40538"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9bf4fbbdf623 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-1aa2f20183857a7116b74e6a614e41335e9adc1ae4d37d974f3f00834a48214f"></a>

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

<a id="canonical-13941d8d35f7a58555d53d30744618fc008fe8f49d351599698021710f902438"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9bf4fbbdf623 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-35720f1b88af0a575d01869ee526aee2086ad53d378fbba8edb1e3c00f7f2941"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9bf4fbbdf623 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-371993ef3b11d4c93215d4e50cc62847ddca484a0c741f8255559dad36464a3f"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2dbd01a8e8c9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-4820c77e66af267bef188d8d32fdc201ff64458c7c53576e78e5f2bcce6cfcf5"></a>

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

<a id="canonical-f3ca6c0ecf8dda9e158f02a5db878487623ac01424098e0354b95c457b17fe0e"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2dbd01a8e8c9 / 3

<a id="canonical-e463c3babf346a166a5ddf52d34d61b50630c7c2a1e715aab1ab1895870d2bea"></a>

<a id="canonical-3b5ddf60c569ed005157a01ef02d0e3d44dd504b49f395f33fe5cdf805ccd6fb"></a>

## certificate_url property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2dbd01a8e8c9 / 4

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

- [custom_hash_algorithms](resources--workload--reference--group-022.md#canonical-dfc4676d4f2823878ad3ca7dac6a9e4a658fa01f951828d9b3b65f7df0f1b40b): complete subsection reference.

<a id="canonical-3e6b27baac2601bba8ef1b298edbfc7ebadd2de501d09a33ab890ff0ab079c94"></a>

<a id="canonical-8df0e5a877724a351531cbc53cebca3a94543a74b8aa215a119a48b4a9e5bae7"></a>

## description_spec property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2dbd01a8e8c9 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--workload--reference--group-022.md#canonical-c808dcca44c5a24034c7d8b1112cd806645dd52e1cfa4db9def53154828f3228): complete subsection reference.

- [private_key](resources--workload--reference--group-022.md#canonical-d08c49b40773f387e5562d5f37173557a80e7af5f3ab46a42683f6aa6a2b62eb): complete subsection reference.

- [use_system_defaults](resources--workload--reference--group-022.md#canonical-670eade74e1fae4bc35835de6a5378411e334440c7792039f7291041896653f0): complete subsection reference.

<a id="canonical-5f4a7fdcb0462cdfed61d39577a6862dcae6cff2c60db69064f00c6382771e9e"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2dbd01a8e8c9 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms](resources--workload--reference--group-022.md#canonical-dfc4676d4f2823878ad3ca7dac6a9e4a658fa01f951828d9b3b65f7df0f1b40b)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling](resources--workload--reference--group-022.md#canonical-c808dcca44c5a24034c7d8b1112cd806645dd52e1cfa4db9def53154828f3228)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-022.md#canonical-d08c49b40773f387e5562d5f37173557a80e7af5f3ab46a42683f6aa6a2b62eb)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults](resources--workload--reference--group-022.md#canonical-670eade74e1fae4bc35835de6a5378411e334440c7792039f7291041896653f0)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-dfc4676d4f2823878ad3ca7dac6a9e4a658fa01f951828d9b3b65f7df0f1b40b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-29abae8372621a6d40d7e70ae28a4aa1334bdfdaed2c3b7dabbe07181c4581cc"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4dab95a76294 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-31bba327f1b1d9a41e71e25717585064b7451a1585a66c8f9bfd3df444f6db4f"></a>

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

<a id="canonical-8a74871ba4a26415fcaa467d45083274371e2706d37ab818724f116488ded218"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4dab95a76294 / 3

<a id="canonical-b7a9f40ec7e6a10c1602b325c80c367f93e4350888cdab45057b69432316385b"></a>

<a id="canonical-be6de285df55bc80059d4f5e4753b2eeaa6890d871648f5b778909747c2d347b"></a>

## hash_algorithms property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4dab95a76294 / 4

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

<a id="canonical-8775b3d5eaaa04f5948583df0f269ee6b3630b1e080736259ca1eb17f1e040ba"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4dab95a76294 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c808dcca44c5a24034c7d8b1112cd806645dd52e1cfa4db9def53154828f3228"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bfba57d1187581a6367571f81ee2dd626defeb4ea7d7dfc4b490f4af5f5cfc04"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 0e0801fe8b5d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-79c5f8a8e56b1669f2e79ac15ced391b52ee3a01b8b011977b5f8f9817be0854"></a>

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

<a id="canonical-60c2d39bf492011ef200060df37e33c6a69f3d2983ce7fbd6a8c7cb492ae3c87"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 0e0801fe8b5d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5f55e6fc969d6d81daf2a3121b4b5f2c051c8f64a2d51036e79e492791f49d3c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 0e0801fe8b5d / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d08c49b40773f387e5562d5f37173557a80e7af5f3ab46a42683f6aa6a2b62eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f069ba13f984f793bd8bcebaef8f3eb1e33b700f89374672aeb8c040baa1245e"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / bd0d30c64a0f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key

<a id="canonical-b066463d88670e9d0d4434e541077187cb7ac90b455cfb7e0255dcfe399a3884"></a>

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

<a id="canonical-ebe885f7202785ad83cf790c14f095007b37c3376cffaa3ae34263a1427075dc"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / bd0d30c64a0f / 3

- [blindfold_secret_info](resources--workload--reference--group-022.md#canonical-878fe6eea588dd88f058bb5457f6ce768b00224f39bd52ac97cba74c93e1635c): complete subsection reference.

- [clear_secret_info](resources--workload--reference--group-022.md#canonical-f746d3e153040fb4b46877534e503d08348f69bbf68685339ba29a96b35be2db): complete subsection reference.

<a id="canonical-8c1ad8883cb792beac1597348f5f99997d758e2acbdebe280fed554f82baa640"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / bd0d30c64a0f / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](resources--workload--reference--group-022.md#canonical-878fe6eea588dd88f058bb5457f6ce768b00224f39bd52ac97cba74c93e1635c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info](resources--workload--reference--group-022.md#canonical-f746d3e153040fb4b46877534e503d08348f69bbf68685339ba29a96b35be2db)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-878fe6eea588dd88f058bb5457f6ce768b00224f39bd52ac97cba74c93e1635c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e51fc2e9831562710f3e38f831f7bb9dd5449df01af0686b08cb94c5999b25a4"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 49e3727ea265 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-022.md#canonical-d08c49b40773f387e5562d5f37173557a80e7af5f3ab46a42683f6aa6a2b62eb)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3f54ea364fcf5c2b4dba7e588e7022be5987cfcebad7b64e4199ffbb2ebf82b1"></a>

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

<a id="canonical-8e8f634b776e0b9dcf1648d61d3cdca0344b4b1bc4ce5727259fa79db6157d6f"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 49e3727ea265 / 3

<a id="canonical-fc564153921ddecc54e2efa5baefd13dca9a056466af476bc87399c29cf7f6e2"></a>

<a id="canonical-00cc94a8e91b2466f090cf7dd0a0f3baaa6084c0aed931254054cbc15e5ed1f7"></a>

## decryption_provider property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 49e3727ea265 / 4

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

<a id="canonical-235b8620eb3315eaa2805f6dbf3447a15d0dd97edb1d15d9b316a7786924c610"></a>

<a id="canonical-b1fc7e068a4ab59e35bdd440fd0900df84802b3007d09283e84f19749d5fe69d"></a>

## location property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 49e3727ea265 / 5

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

<a id="canonical-ca2e4faffc787de047c94eb41987b9fe11f2dc7032f03d815ba1772e0010d7e1"></a>

<a id="canonical-3971cbde316b0cd911681441b18c4252e2c217ea78d9e5ea330c63a5a89da297"></a>

## store_provider property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 49e3727ea265 / 6

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

<a id="canonical-b388a607651048076016d5e6e57284408728e623dfb16f398b64c4e0c32eae05"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 49e3727ea265 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-022.md#canonical-d08c49b40773f387e5562d5f37173557a80e7af5f3ab46a42683f6aa6a2b62eb)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f746d3e153040fb4b46877534e503d08348f69bbf68685339ba29a96b35be2db"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca81ae2c4a0ff5e96eab2d606bfa61a4df8c2f01ee86c3d5a45913c5a43909ef"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 390bc98d4205 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-022.md#canonical-d08c49b40773f387e5562d5f37173557a80e7af5f3ab46a42683f6aa6a2b62eb)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-5682c4f800bc67cd58e4b83c76b2db8facd454b4561d28f8026de9892205e0e1"></a>

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

<a id="canonical-bae59b51fce81936f041ea1d9fdc7ba08517c23765ecc9931e4b9b444e2255cd"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 390bc98d4205 / 3

<a id="canonical-064f45080e0f8ecff5fd8d14479785e91b32544a3093ed31ffdd10ada1bdd430"></a>

<a id="canonical-444a45371ca75c2c404ccefad850d15ded491d539afb10eff765e78bac901a42"></a>

## provider_ref property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 390bc98d4205 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-e8c572796bf25d7f8d4005114e9593f5ac46825f97e80cc12bdee2dd00b49764"></a>

<a id="canonical-7dbe581c17a4e682b90611be203f83792c800f6b01af9b9c3523d326c2df5fd9"></a>

## url property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 390bc98d4205 / 5

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

<a id="canonical-db6f6039ee9c2ed0828728ebc511970c1d4a967a05776fbfdd4c48be95c07ac2"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 390bc98d4205 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.private_key](resources--workload--reference--group-022.md#canonical-d08c49b40773f387e5562d5f37173557a80e7af5f3ab46a42683f6aa6a2b62eb)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-670eade74e1fae4bc35835de6a5378411e334440c7792039f7291041896653f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c975ff9d69f15e0d0b5c54aa33cea55a0a7f50ecba432beda00af3dd8fc13a1"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1f8cce5ed34e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-8d44b586bf02377f1f5b763febeefba7f735899276bdd510dad810a8458829a6"></a>

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

<a id="canonical-0177c52c00996e66dc88bbd36a714d1bc76e10d7975d6e266ae64953653cc0e4"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1f8cce5ed34e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fb946050f444b93375d47994d4933aa5f7e82502a9001e5b3805081e631be5fd"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1f8cce5ed34e / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-022.md#canonical-3aa8518ec03cb7667ad7c7bab92adcc3e1b8a78fb0704890975d8df351398af9)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3b01f26e3af33a2e8c89a4d59489ac025852c5f2ae3e88a7a80184fbcfbee608"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b98844b9ceea2537fc2572bcf3168b5adb9919b48c0ca673d1d555d6dced3ff"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e375c816a74a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config

<a id="canonical-84fd175e9fe20652550ee426ba53544c66c747cd4ed0fd32f16247461504bd6e"></a>

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

<a id="canonical-d75fdab50a6f6d0dbf70b533df8e344093eacdcc5dde9d4fb7268d166e4d9563"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e375c816a74a / 3

- [custom_security](resources--workload--reference--group-022.md#canonical-c76c2194948c546318601df4f9244de8d1af0acec701dca6511adea9886536cd): complete subsection reference.

- [default_security](resources--workload--reference--group-022.md#canonical-645d4375c3c28c088bef1ed514372fb45beb941281fa11c467500408732663c2): complete subsection reference.

- [low_security](resources--workload--reference--group-022.md#canonical-d9b22d218035769e4fe5ddae3d21294786186923ef2d1fa30701cda5b75a8461): complete subsection reference.

- [medium_security](resources--workload--reference--group-022.md#canonical-29bc6758366ad2e28b6cb74e5c4d67b2cced4d9d295e600007f00ee76385108a): complete subsection reference.

<a id="canonical-8a9f5eebe17223e3652b08557e4a822e4af8a62ab27ddbc1d23b9b589d00f5c0"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e375c816a74a / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security](resources--workload--reference--group-022.md#canonical-c76c2194948c546318601df4f9244de8d1af0acec701dca6511adea9886536cd)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security](resources--workload--reference--group-022.md#canonical-645d4375c3c28c088bef1ed514372fb45beb941281fa11c467500408732663c2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security](resources--workload--reference--group-022.md#canonical-d9b22d218035769e4fe5ddae3d21294786186923ef2d1fa30701cda5b75a8461)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security](resources--workload--reference--group-022.md#canonical-29bc6758366ad2e28b6cb74e5c4d67b2cced4d9d295e600007f00ee76385108a)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c76c2194948c546318601df4f9244de8d1af0acec701dca6511adea9886536cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0724d43402344c025b4c1caf57ee42e54f0994d43989784e26ea0988e2463470"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 46820ccd2c7a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-022.md#canonical-3b01f26e3af33a2e8c89a4d59489ac025852c5f2ae3e88a7a80184fbcfbee608)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.custom_security

<a id="canonical-e8dc1d9eb6bb9429fae50b4ea0bd54c224a818934b232d1aed467c1dc5d7f303"></a>

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

<a id="canonical-8f3059494554175f9da042b21d1f5b12a3ab2404ab1eef592971d6c9cc1abd8d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 46820ccd2c7a / 3

<a id="canonical-03a12bff358cf4cba29b29440b54f350f1beedca487b5bbdd85e50c38bf3dbea"></a>

<a id="canonical-dd22650b0b9a385f3a20ab88321ec3a541f40ff00a56f875852ada5f12824267"></a>

## cipher_suites property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 46820ccd2c7a / 4

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

<a id="canonical-938b2592c431047cf0ff8ec95be24fc580b966a504671cdd31979d3ab730002b"></a>

<a id="canonical-febe17ca0be5b4e515c2f868ae9dc95220e1663951e0024652ee9a67cbf1a6a0"></a>

## max_version property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 46820ccd2c7a / 5

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

<a id="canonical-7a1081bafee706b8ab0d71c3f20a7b310464d4ce755f3d88fd532dd70af53e8a"></a>

<a id="canonical-710c461cf11f7fd6b78ff4f2b9518852362715aca6a5387e58904e9666ce8d85"></a>

## min_version property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 46820ccd2c7a / 6

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

<a id="canonical-5e2fea157e2acc43312475816d3a3a531d3cb6045fa6215c56541e535d8a9897"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 46820ccd2c7a / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-022.md#canonical-3b01f26e3af33a2e8c89a4d59489ac025852c5f2ae3e88a7a80184fbcfbee608)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-645d4375c3c28c088bef1ed514372fb45beb941281fa11c467500408732663c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b621fffaa14ad00182f24e45e2a54e15390a46787ea49987088744812aaa4d0"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 5f26b698dc3e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-022.md#canonical-3b01f26e3af33a2e8c89a4d59489ac025852c5f2ae3e88a7a80184fbcfbee608)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.default_security

<a id="canonical-04bf10a1713a651282893e78678edcaa1c1cadf4b24d028b2e836ee35a2c8408"></a>

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

<a id="canonical-8d5958de1a3999a5b516f005fdf639058abf2796fcfe2cda99e898b88e426688"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 5f26b698dc3e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-277c02eb0c6600c5fd984b7d5bdc7b4a7d33e11b638cc723d7c9b6ccfae85308"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 5f26b698dc3e / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-022.md#canonical-3b01f26e3af33a2e8c89a4d59489ac025852c5f2ae3e88a7a80184fbcfbee608)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d9b22d218035769e4fe5ddae3d21294786186923ef2d1fa30701cda5b75a8461"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f362d4d2d52c67e3419b3ca289f6adc0dbec3f6d0b764b0c57236a95db11949"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 82e62842fa80 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-022.md#canonical-3b01f26e3af33a2e8c89a4d59489ac025852c5f2ae3e88a7a80184fbcfbee608)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.low_security

<a id="canonical-12416c7ddcf5eade47b92367044ec6bef4713e5c618d581fca715dc7f814f7c6"></a>

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

<a id="canonical-1b49e43362ca8b0df9edc18fd27c209ce922b23ffff1a14b54526761e67c7f07"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 82e62842fa80 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1b253e54435d9ed807f9204e1edb3d46b64d2dea563990229f5879430e2ea0c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 82e62842fa80 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-022.md#canonical-3b01f26e3af33a2e8c89a4d59489ac025852c5f2ae3e88a7a80184fbcfbee608)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-29bc6758366ad2e28b6cb74e5c4d67b2cced4d9d295e600007f00ee76385108a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf6d31ced457696f19233701098b330e84f288852966a778e4beaa63e2992ed2"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / db4c2e9f97b6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-022.md#canonical-3b01f26e3af33a2e8c89a4d59489ac025852c5f2ae3e88a7a80184fbcfbee608)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config.medium_security

<a id="canonical-7346e11a466a5eb95fb134ff2ed18bd9a2d51c79c155ba1bc22e329180f362ba"></a>

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

<a id="canonical-6f16e9320658585ef910e45a615415a8942e283ac416153aaef615fae89034ba"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / db4c2e9f97b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2ca8f79d219214cbe3c3f1d58abb1d85d6c6b2c4204387dca2c18004eb019055"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / db4c2e9f97b6 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-022.md#canonical-3b01f26e3af33a2e8c89a4d59489ac025852c5f2ae3e88a7a80184fbcfbee608)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93ce496295cadd8e8f3539d165d730535b0d4141a82e3ec735e7d3b90ca418e3"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2ccdbd76df11 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls

<a id="canonical-2036b8f035e0b4813ff7d400dad747e902a508fdb88a5c0f757aaca0af8b5b1b"></a>

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

<a id="canonical-9d221d89d2a9b0ecb8ca906a1341bc182de9135a64b8ed7850f0afc1d2a2a890"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2ccdbd76df11 / 3

<a id="canonical-d87519f6ae80a451ce4eed7800e4061aba1037af43af7635961ba975dc97d2eb"></a>

<a id="canonical-f964834c952d5692f32e801272ede55da8797b14a866477a4e042d4eb6524dd5"></a>

## client_certificate_optional property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2ccdbd76df11 / 4

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

- [crl](resources--workload--reference--group-022.md#canonical-ff8bac6750243239672dffc5093b22f9a4afa8feab3fb880e569e54c9340e5c7): complete subsection reference.

- [no_crl](resources--workload--reference--group-022.md#canonical-4b7640cbe2b3ac6966915ac40eca0e7dca30408e83769351e85850614c57cd73): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-022.md#canonical-7b83b8d1e9d2c8a3a0263295916517c69ff7d5389b0086850652c97b6991d463): complete subsection reference.

<a id="canonical-dcee7e32bda303c842844d47c2a0ec9977765a6af018fc767282173bbdf36960"></a>

<a id="canonical-a267eaac85c71e7bfb6f41a140dddeb6a14515b83edeb52e1bacd18b121efc6f"></a>

## trusted_ca_url property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2ccdbd76df11 / 5

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

- [xfcc_disabled](resources--workload--reference--group-022.md#canonical-dd90481936aee42d770ba30a7e2696750fc9bc98619027ca1123b4c8bc869f4f): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-022.md#canonical-c60a0317945b17255d41d4ad838d93af32d8ef915136eec6eece9d54becc7013): complete subsection reference.

<a id="canonical-dcd069d175c212ef2a73a505a94f27dc03aeed56166c3e5eab0a8f4138ea31b4"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2ccdbd76df11 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl](resources--workload--reference--group-022.md#canonical-ff8bac6750243239672dffc5093b22f9a4afa8feab3fb880e569e54c9340e5c7)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](resources--workload--reference--group-022.md#canonical-4b7640cbe2b3ac6966915ac40eca0e7dca30408e83769351e85850614c57cd73)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](resources--workload--reference--group-022.md#canonical-7b83b8d1e9d2c8a3a0263295916517c69ff7d5389b0086850652c97b6991d463)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](resources--workload--reference--group-022.md#canonical-dd90481936aee42d770ba30a7e2696750fc9bc98619027ca1123b4c8bc869f4f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](resources--workload--reference--group-022.md#canonical-c60a0317945b17255d41d4ad838d93af32d8ef915136eec6eece9d54becc7013)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ff8bac6750243239672dffc5093b22f9a4afa8feab3fb880e569e54c9340e5c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ad150b3f965618d704cb485343e249fcf987f75e0eb52c216660b7524b589c9"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 38338712fa18 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-0bc5dfc7ac114202510ce0aedb8ea14903b4add2f7164e945863347102f57fbe"></a>

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

<a id="canonical-83a16dc734a3e2f53c7901d825be623dcf2abe2e39759c0528247b461c89cc4d"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 38338712fa18 / 3

<a id="canonical-bc1f14964b1893816dec526dcc14a7863dd8435701c995797416d3e81643d4af"></a>

<a id="canonical-3d477139b862a999146653851d3fe0fd2b085b60a1c9eee8a2b9e84db6317c1d"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 38338712fa18 / 4

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

<a id="canonical-45ea1a331df71c2a89f26eb60b3fe7e9d56be4b8931960d3309fe8685b79784d"></a>

<a id="canonical-5076b51517b4aec49d53631592fe3c87ca20c983c0c56104a9d00032579d7de2"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 38338712fa18 / 5

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

<a id="canonical-4976bf23aeb0c99ba399e883b7bb7b8887065873a4fc65133079ea9340771c86"></a>

<a id="canonical-68fc6fb4a19a2a62313ebf9a3979258cd739c73cac18fc279c68098209bf99a6"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 38338712fa18 / 6

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

<a id="canonical-89bfd4b979da67231866bae8f289ede94b18927ae523a573aea3da8f739b57fa"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 38338712fa18 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4b7640cbe2b3ac6966915ac40eca0e7dca30408e83769351e85850614c57cd73"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33a086426702192576bf53744bd6743de5fea9675bcd1242a3b841861a8d93c0"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / efede98a2e66 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-b644ee5b8652c2e2a3db1f81525da69a8e3a6a6f3dd4b0eea1322cf809c5cdc9"></a>

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

<a id="canonical-90091ff488332563d9bec2f39fcf55bb9b436d7a81c11b6dc35dceb20fcd9b39"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / efede98a2e66 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-603726ed95919562b53823d1ed11218f61c596d99dd6f70894d3707c73e268c6"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / efede98a2e66 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7b83b8d1e9d2c8a3a0263295916517c69ff7d5389b0086850652c97b6991d463"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbfde3ddc5f35ce86911833c18cfcabcf41cfc90959ab8ffba7ac87386be29e2"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 884ff1242360 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-1e708c1c03168618eb2e4ee95e3172e5c2aa6dc5efcdcc439c67d68b7c8bab4a"></a>

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

<a id="canonical-b8ff9f99fecb9d6bd610772fb04e5646e8a88204babf69e7f5df51346536b684"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 884ff1242360 / 3

<a id="canonical-6e123b8853903788aae80bad65380947025adca6637872949583f767f590d155"></a>

<a id="canonical-7ebaf1155c056f53f6e565ff7db0a4c7a5a918381a83c20db990a6d24d105441"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 884ff1242360 / 4

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

<a id="canonical-546c1b9485f4d9d09ddf4546da979d0deae449d0ce69178420dda194a3a0725f"></a>

<a id="canonical-baba826d9f8c16fb3c1b2ba215623f9f0af555f4593ecf0261036e2ce84869aa"></a>

## namespace property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 884ff1242360 / 5

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

<a id="canonical-ecad7e6c010d35b90051e54c8076ba83ecb6e62a9cd9d7ea8f0d17f2aea02bd1"></a>

<a id="canonical-ca0b2805aae106fa5b9d54c2a9ae20370bbac78d0575f2082a3ba7d392d54624"></a>

## tenant property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 884ff1242360 / 6

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

<a id="canonical-f2dca57cd019dfc4c7c91b3fe705c786d43c2debc16f139727ea3de5648cd295"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 884ff1242360 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-dd90481936aee42d770ba30a7e2696750fc9bc98619027ca1123b4c8bc869f4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d380e923334e7c2e761c9efc313f8eccd7a29e7889710da43f4b4770fd0c655e"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 5dac34c733b8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-f20b7d773bbfeb9d5d6fe721d268895e609536333d387024008926552ec16f61"></a>

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

<a id="canonical-2fe2f2f1c1b878587531f1325b15acd54770a29d2166527817fb9412e3b2e2b9"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 5dac34c733b8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40555febfb8d0345d44af829bfeec572e3512b80d4e351c07f7e5ea211e36aeb"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 5dac34c733b8 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c60a0317945b17255d41d4ad838d93af32d8ef915136eec6eece9d54becc7013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9bfc4c02b1552ffa0316be6fc4cd2f770f43a95ad2512fb095a815187cf4a7f0"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e9d859b2e205 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-021.md#canonical-398ea5a255cdab1b82b4ad2576cc7b2f8d7ae7b555a2382f22d11fea8584e7ed)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-022.md#canonical-37dc5503a7937ea8f391e0a2249f289308a180b081263b17445a57a658cf3cf5)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-2c8609c6cd2227ed117ea2a569b7d41f2b378bce6a7c2a90131971d293262800"></a>

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

<a id="canonical-4dea4a0ac7b0de6963d6c737ba3f4d89865cfac6c6446cbae36b5ab7ba59b1d7"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e9d859b2e205 / 3

<a id="canonical-4bd5ffaeee9779806349d8c9b2a92189d47a46b7108d86dbbaf292be27cd8723"></a>

<a id="canonical-7c61b913167fd788182ca9c99f40b496b4ed0af99163c709b27a201266ccf33b"></a>

## xfcc_header_elements property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e9d859b2e205 / 4

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

<a id="canonical-1ff1617e3c1cfe45fc514c07d3b9f25374cceb66c6393eff01a8a5fc1786b440"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e9d859b2e205 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-022.md#canonical-95cd6fc2fbaea2ed7c645ca01c7060bc847d58d40712158ca68599fbf1aa3c21)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8e95b0f301d815d9959bcfd11fe17868b1d410c12fe2585505248cdef4156f40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19e06f966c92491b49d678b6efa35be1a1d896bd139c506be378b331690e90d2"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 20f00e9c53c6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert

<a id="canonical-85a436a7e66473561c08a310fff2a836fb790fdbe8fea65a2edd5becff8f0487"></a>

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

<a id="canonical-467964d47bb7927291c6dc3ae3174815835702170b0b77f8501e0279b9b4196e"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 20f00e9c53c6 / 3

<a id="canonical-45cd1c16d676e2bb4a74edd4af5923258e3e23707ee715f81ed4f16579e5f2b0"></a>

<a id="canonical-1635b5f7d008212a8d405192c2f0b849c2152b952b6f3c16f17d857755e442ed"></a>

## add_hsts property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 20f00e9c53c6 / 4

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

<a id="canonical-50cc768ee5827dcfe97fcfa725d10bc4421443bc1c05b5a479f5131a6a79e017"></a>

<a id="canonical-0895cfc4fd313e77a80b6bf6b359aac8adb78c2de05d03183a878c50a6f23133"></a>

## append_server_name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 20f00e9c53c6 / 5

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

- [coalescing_options](resources--workload--reference--group-022.md#canonical-96dbaa1808de1db06686f0754d55ae200672c755456acffd7c8406805efdae22): complete subsection reference.

<a id="canonical-d2cacb9b55198f40a133f8e84c748f1e50d303294c928c1ca814589ecbf88227"></a>

<a id="canonical-2e649a996c9c034190c92429738d1151349a15adde25a9d150d2b439b2007188"></a>

## connection_idle_timeout property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 20f00e9c53c6 / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
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

- [default_header](resources--workload--reference--group-022.md#canonical-7a5371b61e43eac18e21da792dd6a420a716d3dc98d375be233e486c7d2c44cc): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-022.md#canonical-d1b569908840405a4b72ee1657bde1db81d63ecfa114fc65dfc90b3d45d8a7ee): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-022.md#canonical-5a14edc3d26d54ffebcca3cc64c899c0a6b160c16d7f919926cd5e32171f68cf): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-023.md#canonical-0bf61ff8b991ff311a3aa1beadaa10eb3a467efe99e3d7c19a7f2ee03e71dd10): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-023.md#canonical-b7dba8a6c4aff331156ed26b071c04df64eee255a5b175af6eefcdbbed4a9171): complete subsection reference.

<a id="canonical-407ef82ed36f60cfa835a15689a882e1324f91184fc5b5cef8b43c72aa678ed1"></a>

<a id="canonical-31c9caede0dc5454884e714a39119cabfcf5ff63804127c169033283cab83016"></a>

## http_redirect property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 20f00e9c53c6 / 7

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

- [no_mtls](resources--workload--reference--group-023.md#canonical-9fb1c00f71e144c7f4c6b9f5a32096de943c6c6127c4cb4372cef6a98b639f3c): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-023.md#canonical-dafbe8e472fde81f9f971b879bd9a44c65371026d2255a074683147ae2673a36): complete subsection reference.

- [pass_through](resources--workload--reference--group-023.md#canonical-52f1cb612497bfc1611b5913d8cdf29092435a5ac78b5699e32e42a52c1c05e8): complete subsection reference.

<a id="canonical-f4ea24e5163b34e890010c1d359c598019db2037171029fb97adc86f26beecd9"></a>

<a id="canonical-8756aa001303a9335862cfaf57d268ac4f37657286c3fb123a83103e665dae1f"></a>

## port property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 20f00e9c53c6 / 8

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

<a id="canonical-9ef7722273cad2da060cbb201a78354e3d321ed92eeab3e0bce784f25c3ccd8e"></a>

<a id="canonical-1a30243d75adcbb779232ead206ec07beda25f44684eb36b6d9eed9bcbcf1e55"></a>

## port_ranges property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 20f00e9c53c6 / 9

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

<a id="canonical-efe97c6bcd2ee430f5df295d4187ac2c0db5a61124f38f663e5283e2166a1cbc"></a>

<a id="canonical-fc71f26e5ec2370ac957c777329c23b6241441b9e7841f7508c08684864b0891"></a>

## server_name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 20f00e9c53c6 / 10

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

- [tls_config](resources--workload--reference--group-023.md#canonical-49eb4fb24ccca74f39297f0c9193774c9b30091a2b81123532c3dae218185888): complete subsection reference.

- [use_mtls](resources--workload--reference--group-023.md#canonical-f105d23f1551c555475368bc6d302d784198fcd7417710210e3d0b43075fdfcf): complete subsection reference.

<a id="canonical-27147c326f05871a4f2835af05b5b43bb2ef6be4369f4f898c78722b970ce3ea"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 20f00e9c53c6 / 11

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-022.md#canonical-96dbaa1808de1db06686f0754d55ae200672c755456acffd7c8406805efdae22)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header](resources--workload--reference--group-022.md#canonical-7a5371b61e43eac18e21da792dd6a420a716d3dc98d375be233e486c7d2c44cc)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer](resources--workload--reference--group-022.md#canonical-d1b569908840405a4b72ee1657bde1db81d63ecfa114fc65dfc90b3d45d8a7ee)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.disable_path_normalize](resources--workload--reference--group-022.md#canonical-5a14edc3d26d54ffebcca3cc64c899c0a6b160c16d7f919926cd5e32171f68cf)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.enable_path_normalize](resources--workload--reference--group-023.md#canonical-0bf61ff8b991ff311a3aa1beadaa10eb3a467efe99e3d7c19a7f2ee03e71dd10)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-023.md#canonical-b7dba8a6c4aff331156ed26b071c04df64eee255a5b175af6eefcdbbed4a9171)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.no_mtls](resources--workload--reference--group-023.md#canonical-9fb1c00f71e144c7f4c6b9f5a32096de943c6c6127c4cb4372cef6a98b639f3c)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer](resources--workload--reference--group-023.md#canonical-dafbe8e472fde81f9f971b879bd9a44c65371026d2255a074683147ae2673a36)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.pass_through](resources--workload--reference--group-023.md#canonical-52f1cb612497bfc1611b5913d8cdf29092435a5ac78b5699e32e42a52c1c05e8)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-023.md#canonical-49eb4fb24ccca74f39297f0c9193774c9b30091a2b81123532c3dae218185888)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-023.md#canonical-f105d23f1551c555475368bc6d302d784198fcd7417710210e3d0b43075fdfcf)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-96dbaa1808de1db06686f0754d55ae200672c755456acffd7c8406805efdae22"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d58a7cc176d8ea8a46d5c6220b8148f7ab6dc7b8fe739f65ebd994762dc20565"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 40b52e151ed3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-022.md#canonical-8e95b0f301d815d9959bcfd11fe17868b1d410c12fe2585505248cdef4156f40)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-a42dde182e2f2fd4812fa246f2889126adfab5ca1a37a9d67f0f47d9ea28e369"></a>

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

<a id="canonical-396ae7bfa611a580dad264c3c1e59cd8656afdf1b4242a30e8c4970bbc72a3ea"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 40b52e151ed3 / 3

- [default_coalescing](resources--workload--reference--group-022.md#canonical-a68ab04eb45767b22c9262053ef8f81783b0d343a632a4b6d4172462b2ddd535): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-022.md#canonical-1b1ce6e2f6c69451a5b2cff7844580888950799c442739bd70a863e77e6ad552): complete subsection reference.

<a id="canonical-36f814ef2606fa842853168f3b8530f28a5f2cc14b9b15e1c99979eb95e665ed"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 40b52e151ed3 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](resources--workload--reference--group-022.md#canonical-a68ab04eb45767b22c9262053ef8f81783b0d343a632a4b6d4172462b2ddd535)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](resources--workload--reference--group-022.md#canonical-1b1ce6e2f6c69451a5b2cff7844580888950799c442739bd70a863e77e6ad552)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-022.md#canonical-8e95b0f301d815d9959bcfd11fe17868b1d410c12fe2585505248cdef4156f40)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a68ab04eb45767b22c9262053ef8f81783b0d343a632a4b6d4172462b2ddd535"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afea7b7dd2f3cae6edd7864d173f305200532b999b89d393d121fe1dadea748b"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e17670c6bbd2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-022.md#canonical-8e95b0f301d815d9959bcfd11fe17868b1d410c12fe2585505248cdef4156f40)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-022.md#canonical-96dbaa1808de1db06686f0754d55ae200672c755456acffd7c8406805efdae22)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-f67adf7b50d16440f4b99aea529f4ea5dfe23a0984c6d91b092954ab32b6f83d"></a>

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

<a id="canonical-c1ee4ff2b9507578f802adc2c236297b10f9a56a15178bd1a95b2112e0f5886a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e17670c6bbd2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-39cfb73103222101beafd0ec03dd12817a0451d5208c00fa22909ee74d04cf5f"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / e17670c6bbd2 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-022.md#canonical-96dbaa1808de1db06686f0754d55ae200672c755456acffd7c8406805efdae22)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1b1ce6e2f6c69451a5b2cff7844580888950799c442739bd70a863e77e6ad552"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5cfc665a0c39e33e21033f0bcfdd7abb018f63185897b0cadb5f13905857b44"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2658775057d3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-022.md#canonical-8e95b0f301d815d9959bcfd11fe17868b1d410c12fe2585505248cdef4156f40)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-022.md#canonical-96dbaa1808de1db06686f0754d55ae200672c755456acffd7c8406805efdae22)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-6319b17bd3339e62f77dd26ccdbab4f3c91bdedb48543f8993536cdfc074750f"></a>

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

<a id="canonical-70b5da27f767ca2ca1ab15624daeb3a9c25c8651698ec0453647cfde5a8b5c1e"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2658775057d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0e82a236ac551197e703e4b2a7f07a8f2c5f24674a415e56df694314d3829ee6"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 2658775057d3 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-022.md#canonical-96dbaa1808de1db06686f0754d55ae200672c755456acffd7c8406805efdae22)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7a5371b61e43eac18e21da792dd6a420a716d3dc98d375be233e486c7d2c44cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d561b4f5d4e525d166f06acfa2f56bc64d89ae593bf2c3be27d22f824bf0637"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 812827c37223 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-022.md#canonical-8e95b0f301d815d9959bcfd11fe17868b1d410c12fe2585505248cdef4156f40)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-a49a1829b619c1b6a539bd8f8e0e3950be5a9019825611ac0b8de528ef90279c"></a>

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

<a id="canonical-3d81ef60d4f484b357754da8ba32751a4f1cf5c339b468cc4346d72115091515"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 812827c37223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ba6dce7c03b8ad4d9d1bed4221a2892a8557caa17056788baefe25fce23598e5"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 812827c37223 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-022.md#canonical-8e95b0f301d815d9959bcfd11fe17868b1d410c12fe2585505248cdef4156f40)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d1b569908840405a4b72ee1657bde1db81d63ecfa114fc65dfc90b3d45d8a7ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40d73e410ed21c9012040bdcb8580230608e2a5e003ba777ec7b2284d456e92a"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f54160fc3aac / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-022.md#canonical-8e95b0f301d815d9959bcfd11fe17868b1d410c12fe2585505248cdef4156f40)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-21b74b63d416d0df879d269c3f9a816cf60e80e4e9c4d12ba5f349f95d0d01ae"></a>

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

<a id="canonical-0fa90a10fac9901495057ca826e80da867d6966a28e99cae30cff0e32d941dfa"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f54160fc3aac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8cb645e1598a21f29aa3a7182e687d99bfd4a442c10bd4e55ab521abac4cb36c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f54160fc3aac / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-022.md#canonical-8e95b0f301d815d9959bcfd11fe17868b1d410c12fe2585505248cdef4156f40)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-5a14edc3d26d54ffebcca3cc64c899c0a6b160c16d7f919926cd5e32171f68cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
