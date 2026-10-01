---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-e6c99028bea21e0977c3801feee2ef56327ebe92578d8602e23814aa9b8c7d64"></a>

## trusted_ca_url property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 9b4cc392a886 / 5

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

- [xfcc_disabled](resources--workload--reference--group-019.md#canonical-008ff5f55cb0a79841272b6984ca0501e16b75b5038152d81ac5742ae39acdc5): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-019.md#canonical-2e470ed042a6e19fc50908bef9e5baf00473026b87b2bfe9449183de0b0e5da9): complete subsection reference.

<a id="canonical-0de8e617b88dc170edc8219f03e96025ba39e63727c22dc46de532cb7895a9cd"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 9b4cc392a886 / 6

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl](resources--workload--reference--group-019.md#canonical-15601ca81066abfbe49595396789b961e32e4535bcbad6d6ddaf7e3293d20115)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl](resources--workload--reference--group-019.md#canonical-a9e4f366b1c601514b9fc2aba185a577f6837e855669611ce07e1953a3b64eb6)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca](resources--workload--reference--group-019.md#canonical-cc9e8bc9d2af4e22c9c06b8d218107cbce1b9048768f23ca79e20a8b6415d7fe)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled](resources--workload--reference--group-019.md#canonical-008ff5f55cb0a79841272b6984ca0501e16b75b5038152d81ac5742ae39acdc5)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options](resources--workload--reference--group-019.md#canonical-2e470ed042a6e19fc50908bef9e5baf00473026b87b2bfe9449183de0b0e5da9)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-88e031cfad4357b09e79f29ec26027a3ddb61f3d2d19a38b6294452c2c1da6a9)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-15601ca81066abfbe49595396789b961e32e4535bcbad6d6ddaf7e3293d20115"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-724738f4d97d0ea9a889031e53de80cd2890c49da1c3725cdfa8f5196972e9c8"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 42576e812cae / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-88e031cfad4357b09e79f29ec26027a3ddb61f3d2d19a38b6294452c2c1da6a9)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-018.md#canonical-0208f7bdd5e59679b01da63c8e183144c211c76b12b7187c0b67e399e3a2aba4)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.crl

<a id="canonical-3ff178816fffc68d80fedf046a1d375b16ddb15b58d3108ace3ab8343864071d"></a>

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

<a id="canonical-549a1d6127a59c97705fb0d0cfba50d8ebbb0ff5b7fd3514f3584351b5eab598"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 42576e812cae / 3

<a id="canonical-5a47446f901484a8a32d0a1df6ffcfb01fa1e6fa5f2f70c7d99f0d04322f3bf2"></a>

<a id="canonical-ae4f2eb96ac87c64dfe9722e60c1858f38dc81a1b4e9df3106faf04e4b326532"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 42576e812cae / 4

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

<a id="canonical-1f0c33313635053d6b28c67bc62718061f6eebd71056e0e98a98136d943a0eb5"></a>

<a id="canonical-89985ff457ac0b4b10eeb2ff8cac56b8a1fc951988caddca4cbd9970cb5654e9"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 42576e812cae / 5

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

<a id="canonical-a5e4009a554ca4caea72d4e799023f94267424411ee0eb696f57952c7daa528d"></a>

<a id="canonical-fe2dad0c9fca073ba1ec2d2f9701262e4fbff0487d9ea5aa5b765513fdcb2b69"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 42576e812cae / 6

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

<a id="canonical-8dcf4744ef971f24565b2c1cb892a7faf0f438e289f3d362ccbd2724213cc685"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 42576e812cae / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-018.md#canonical-0208f7bdd5e59679b01da63c8e183144c211c76b12b7187c0b67e399e3a2aba4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a9e4f366b1c601514b9fc2aba185a577f6837e855669611ce07e1953a3b64eb6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e250a7692426a636c3636f1b92ddb16a4a7370630f21eb3db453b1dfdb2f012b"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 49e201237965 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-88e031cfad4357b09e79f29ec26027a3ddb61f3d2d19a38b6294452c2c1da6a9)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-018.md#canonical-0208f7bdd5e59679b01da63c8e183144c211c76b12b7187c0b67e399e3a2aba4)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.no_crl

<a id="canonical-6df970d3e55d7fe87b7d21a8b06e9d54c7afe651809e5701024406de0d8949fb"></a>

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

<a id="canonical-41c5bcc406d0b3670ebc826c9f88944f9fed4100232637ece7f8e9ca80d0c3c2"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 49e201237965 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2ff852e5b412d354313e4e2041cadf847205561b15bb9e7775880f32af4ca2de"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 49e201237965 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-018.md#canonical-0208f7bdd5e59679b01da63c8e183144c211c76b12b7187c0b67e399e3a2aba4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cc9e8bc9d2af4e22c9c06b8d218107cbce1b9048768f23ca79e20a8b6415d7fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a23d025857b64de8e814818c1f89dbc006333ba6cbdf246a6ebd396cb17beec"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6f8fdb15c439 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-88e031cfad4357b09e79f29ec26027a3ddb61f3d2d19a38b6294452c2c1da6a9)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-018.md#canonical-0208f7bdd5e59679b01da63c8e183144c211c76b12b7187c0b67e399e3a2aba4)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-3ced81931e33e3db3a4fe577ed7b166413b41b9a004c92739f7668cbf69128b1"></a>

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

<a id="canonical-66e1eeba42d20f05fe1ac1ae0631c88f91f3f4ab1f3730a5d745876004616c3e"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6f8fdb15c439 / 3

<a id="canonical-e32d4fc5df7fd3ba71c86f59c14566ec9a7bfb1cd242a1e52440494e79dd3524"></a>

<a id="canonical-bcb98ef69a1312500b1bd9d4f6f2cbedba8dc0b94e207048c5746c9cc01c948d"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6f8fdb15c439 / 4

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

<a id="canonical-0025df23e5db585540d144a4cdee3ef879ab5f4d91628ba43bfb4a9913310537"></a>

<a id="canonical-427019b2a122c65e94e196c9cf5a5fd71f097c479fb37ddc489743cecfbbdecb"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6f8fdb15c439 / 5

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

<a id="canonical-4bff9f211533ee5854e33bef54690a2179747da6bebcac5d29f7f36a533a548f"></a>

<a id="canonical-19c3b4183b531ea27aa8f2933060f49b2c38854a0e24ffe3cc906f6c72aa177f"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6f8fdb15c439 / 6

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

<a id="canonical-444d4c871b809fa738c4c4daa0b553990ac57488ad377277781e5f8ef09e266c"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 6f8fdb15c439 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-018.md#canonical-0208f7bdd5e59679b01da63c8e183144c211c76b12b7187c0b67e399e3a2aba4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-008ff5f55cb0a79841272b6984ca0501e16b75b5038152d81ac5742ae39acdc5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d175a50a7280757bbeb4b0bf2980cb58172b646effb9718afe0d0b1f3cbf92fe"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 61c93e59c5ef / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-88e031cfad4357b09e79f29ec26027a3ddb61f3d2d19a38b6294452c2c1da6a9)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-018.md#canonical-0208f7bdd5e59679b01da63c8e183144c211c76b12b7187c0b67e399e3a2aba4)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-ffd14e86579d7d9605b7e4316e7f3e79e3e8afd7dbd5577e99504d1ad3a24744"></a>

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

<a id="canonical-2cc8994946ccea9109a87f44278d4223802b767ab864871db6708a4a729ec0b0"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 61c93e59c5ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-85444b2d954fcd40a12dfa268371c85c7a130edc9e4c641900ef0911d8d5975c"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 61c93e59c5ef / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-018.md#canonical-0208f7bdd5e59679b01da63c8e183144c211c76b12b7187c0b67e399e3a2aba4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2e470ed042a6e19fc50908bef9e5baf00473026b87b2bfe9449183de0b0e5da9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-89813fed524b7994f656d3d1631a18ecb333813da32540ef916b3a1fb858efac"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 87783456bab4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-017.md#canonical-1622054a074c74fd54c41bd170891da7f136459269bb1c252d02e71dde5f37ff)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-018.md#canonical-88e031cfad4357b09e79f29ec26027a3ddb61f3d2d19a38b6294452c2c1da6a9)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-018.md#canonical-0208f7bdd5e59679b01da63c8e183144c211c76b12b7187c0b67e399e3a2aba4)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-f42a4fb6d0e710d8c5c489f28cd12c711fe6e7f5951e035a07477591988d07da"></a>

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

<a id="canonical-f92e5c4021cdb49a8d1ef301db7e3ffe9cbaa8d4baea7daa458204bd93174df4"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 87783456bab4 / 3

<a id="canonical-ce5ffa77d625775a7221ece073776a3954786b3c7e44367a850d44a0c24c0a54"></a>

<a id="canonical-abb960415834561faa2ca1f688315297c11385b5bd2a35d25475f1a3ffea44ad"></a>

## xfcc_header_elements property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 87783456bab4 / 4

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

<a id="canonical-2e4e5dc8897e1a557a6e9abdb6783068492f7ecc9aa3ff5bc6289c16bc0f2adc"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 87783456bab4 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-018.md#canonical-0208f7bdd5e59679b01da63c8e183144c211c76b12b7187c0b67e399e3a2aba4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95a0a99cec7bc648718a26e062c8c759bd2177e53eb5340ec3d6743845aa0f1c"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1887b34b7153 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert

<a id="canonical-78a42549beaa28fc42d6523b9a3ee0f2288fd1f4fe89620d6345710b376fa827"></a>

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

<a id="canonical-6b7d52a6c40e63b4467e96a56964955fe0dccfe00f3b47a280c5cba39b64b01f"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1887b34b7153 / 3

<a id="canonical-d5e77c46aca5164637c02812918dc0aea8b899e3f65c47ee92cb310c705d5e01"></a>

<a id="canonical-99623fe8959882e559d94d82d5a8f14ae652c535403ebc4edf82bfcace4a6fc4"></a>

## add_hsts property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1887b34b7153 / 4

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

<a id="canonical-72aac19e58b870cf6a40fd659dc06f82f6e0a317f886c0bb459be66dde53e9b6"></a>

<a id="canonical-59fd683d554806ccc9a9fcc06bcc4eec4eba5deed4c89d2150e433bfff9cbb05"></a>

## append_server_name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1887b34b7153 / 5

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

- [coalescing_options](resources--workload--reference--group-019.md#canonical-d8d3969647b351569282b91f493fac00cb06e07b735743a61d17f7d86a445ebe): complete subsection reference.

<a id="canonical-4f2923aa08fbbd74671d2ed366f171718a9315e0e641c59a65fd325b8cb31c79"></a>

<a id="canonical-c0c5d5a1b8f3db439ddc3f6418a51cd876f29e6871801677c1b69051ed313ffa"></a>

## connection_idle_timeout property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1887b34b7153 / 6

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

- [default_header](resources--workload--reference--group-019.md#canonical-9b807feb415255284261b09b587b918faf47da9ef2c3fe7a6f7d98ba92d6a728): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-019.md#canonical-81be64ac6b4d3c47b1ef722e5f411c998bc685e461077ff28f46ddcbebfb96e6): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-019.md#canonical-2e072e559a8e9e18bb6d6d89ce0c09f7d55c0844291ecba96d0900e3a1f0a534): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-019.md#canonical-3380d97ea372eda74b7bbc71dc2203431571c2583a325f4e6305c5430c59b79d): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe): complete subsection reference.

<a id="canonical-a32261ab98180d9a9ab95bebf4f6659733caa2b28a8b42839b5a7d0a2ee6c084"></a>

<a id="canonical-a6b1b424adb9c94205329225cdc370548ec006674d9486c78ad19e5dc4f5f5d5"></a>

## http_redirect property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1887b34b7153 / 7

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

- [no_mtls](resources--workload--reference--group-019.md#canonical-53303c5fb3c4f861dd3167c2417f23beac9c602e00c753b327175b18ea910e71): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-019.md#canonical-3d0477a1d6976e1f9e8401dee1d41170755558b22c9c6025b4306887ef3f9ba5): complete subsection reference.

- [pass_through](resources--workload--reference--group-019.md#canonical-c3f6ce13d0fb0a9a0c627c7cf998643718eeef204cf8557b85d4bfef6fda281b): complete subsection reference.

<a id="canonical-bdfa861d441960e89bccb8b8fec33fb2ff6073ab50a46ccf256f83d17b4cbed2"></a>

<a id="canonical-35875170554a353b8a0550fa4a7c4aee21c1239db8b4051c214c7b4b7d5078e2"></a>

## port property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1887b34b7153 / 8

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

<a id="canonical-e4fd932d291a691bfe7fd502cf445f16a10f69d09c1a6065d7025552ff4d01bb"></a>

<a id="canonical-ae26c0efaa215b43765ebf1cd1f01ef20ae2edb2a2f158b6218ec6ca8554c490"></a>

## port_ranges property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1887b34b7153 / 9

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

<a id="canonical-385f537bd7bf9e8bca00e95dbad6dabb2f00c089d1c045f3052843002419eb62"></a>

<a id="canonical-38091f35b88fb0f067249d0da7735bcd6e9901fb4d276c26394ebd9e82167882"></a>

## server_name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1887b34b7153 / 10

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

- [tls_config](resources--workload--reference--group-019.md#canonical-d036da666bf50d78482211dd9755254574a3a7061ae84af4319a77859ad15c37): complete subsection reference.

- [use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8): complete subsection reference.

<a id="canonical-1edff883502a5da3870a921eedfff112c8320849d62665f7d9044bf51a7a193f"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1887b34b7153 / 11

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-019.md#canonical-d8d3969647b351569282b91f493fac00cb06e07b735743a61d17f7d86a445ebe)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header](resources--workload--reference--group-019.md#canonical-9b807feb415255284261b09b587b918faf47da9ef2c3fe7a6f7d98ba92d6a728)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer](resources--workload--reference--group-019.md#canonical-81be64ac6b4d3c47b1ef722e5f411c998bc685e461077ff28f46ddcbebfb96e6)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize](resources--workload--reference--group-019.md#canonical-2e072e559a8e9e18bb6d6d89ce0c09f7d55c0844291ecba96d0900e3a1f0a534)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize](resources--workload--reference--group-019.md#canonical-3380d97ea372eda74b7bbc71dc2203431571c2583a325f4e6305c5430c59b79d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls](resources--workload--reference--group-019.md#canonical-53303c5fb3c4f861dd3167c2417f23beac9c602e00c753b327175b18ea910e71)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer](resources--workload--reference--group-019.md#canonical-3d0477a1d6976e1f9e8401dee1d41170755558b22c9c6025b4306887ef3f9ba5)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through](resources--workload--reference--group-019.md#canonical-c3f6ce13d0fb0a9a0c627c7cf998643718eeef204cf8557b85d4bfef6fda281b)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-019.md#canonical-d036da666bf50d78482211dd9755254574a3a7061ae84af4319a77859ad15c37)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d8d3969647b351569282b91f493fac00cb06e07b735743a61d17f7d86a445ebe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d431b10c211cf8aa3c96db8dfd3eb7ee060ffb54bdf85369e5da02f2bc0c0bd4"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a4c12236e8ca / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-c3954dc1caabbccb9eac404fb21ad8b1f4c88d4a49a51c864afaa4816c0fb72e"></a>

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

<a id="canonical-e1909510118753aa1c3eafe58858a7ae9a4e4096dfc5db796374e72cdb39b63f"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a4c12236e8ca / 3

- [default_coalescing](resources--workload--reference--group-019.md#canonical-5203866ebe0a72cc1de0a98d2abb6649970ffb7612ccfb027ddd9e1ee40e6e43): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-019.md#canonical-7fe065332af680fef3aa268467420275491dcdd787970453503ba8a6b04f8b7d): complete subsection reference.

<a id="canonical-2365473906c6a2f9b99b17ee912e0dd81902e3b498065c9953a0abd9f4a3e054"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a4c12236e8ca / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](resources--workload--reference--group-019.md#canonical-5203866ebe0a72cc1de0a98d2abb6649970ffb7612ccfb027ddd9e1ee40e6e43)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](resources--workload--reference--group-019.md#canonical-7fe065332af680fef3aa268467420275491dcdd787970453503ba8a6b04f8b7d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-5203866ebe0a72cc1de0a98d2abb6649970ffb7612ccfb027ddd9e1ee40e6e43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cbda079a2b876a00d9dd74516ffb5083e129a660fd8a8168efe5db07f664033"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bf139b8cf864 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-019.md#canonical-d8d3969647b351569282b91f493fac00cb06e07b735743a61d17f7d86a445ebe)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-9196f2674bee7d345d60f3a0f6ca73a04402a08ee7237e7e66157cfe95f5adf7"></a>

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

<a id="canonical-207f3224160fd127a545bdf3569a4c532be775901eedb39a767e781882034373"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bf139b8cf864 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c56b745ca3ac8390b9ff916d6464b982d21fd83f2a23ee2793e452224decf44b"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bf139b8cf864 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-019.md#canonical-d8d3969647b351569282b91f493fac00cb06e07b735743a61d17f7d86a445ebe)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7fe065332af680fef3aa268467420275491dcdd787970453503ba8a6b04f8b7d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f34b0c1486da2b346d81009d41e2a363b063ef99a91f0583650e7d29aa518dfe"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7de471e38c04 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-019.md#canonical-d8d3969647b351569282b91f493fac00cb06e07b735743a61d17f7d86a445ebe)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-2c56bc9049c90b961814bfc3e4f3b0f7cd9e0d950fd5770ee9942cea8ce49956"></a>

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

<a id="canonical-faef36e66748ca555f90107c921e68cb31992c5044260c64aba9bbee42f66eb0"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7de471e38c04 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8da96c875cefba3a1d8d05c946033c24d668965b11477491fc01bd3bdd84b71d"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 7de471e38c04 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-019.md#canonical-d8d3969647b351569282b91f493fac00cb06e07b735743a61d17f7d86a445ebe)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9b807feb415255284261b09b587b918faf47da9ef2c3fe7a6f7d98ba92d6a728"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ca0e66a20a31d41ba7342898e48434e24e499e7d816db5a1b243ab4f2e24965"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 033bbeb754ea / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-5da03d6c22162896c4649811f6915da6b178449936fb09d047229af81de43c2f"></a>

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

<a id="canonical-fd670012cfde796fab39d089e2ae9fadb28c6e007a807fc8a91f4c276f49cd2a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 033bbeb754ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d124241d95b21abe9538f98004f05427660aeac81f0a99394960f851d84d2a5f"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 033bbeb754ea / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-81be64ac6b4d3c47b1ef722e5f411c998bc685e461077ff28f46ddcbebfb96e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-12eb674fab2044b388a1054503726a7c0ff92e5ec20f6968505fd4e85d8ef084"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b13ed3cc94dc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-039bc215608a4dd33104b8917a3530cf7c86d2692393ebd6af5a26e3e9e43122"></a>

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

<a id="canonical-c55c6b0ea563597dc32422d70b17dfd3c8be70b4ba72ce41cdc7f95260a9ddbb"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b13ed3cc94dc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6aaf5217019489ae150d09706b4e11a0f91d1975b536adce4449afb3df05fd85"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b13ed3cc94dc / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2e072e559a8e9e18bb6d6d89ce0c09f7d55c0844291ecba96d0900e3a1f0a534"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a6ef4392efcd92156997d0ce84f4e68ea7f4cf97da34ce7eb97356d87a35586"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4f04d8b6ce76 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-c737d1851cdbd6f71076e2442b5928487fb41e95f9276a0edb433141ec21314c"></a>

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
disable_path_normalize = {}
```

<a id="canonical-18d1904c5dc3ea0d0d31153b90cf06301e2c1d67e03e397ea86ac1ec280ec8c6"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4f04d8b6ce76 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-560f7fb6b0a6e1fc305f181abc34cbc16cacd9b18a3a50c35412aea6a4984d0e"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4f04d8b6ce76 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3380d97ea372eda74b7bbc71dc2203431571c2583a325f4e6305c5430c59b79d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-925e813b7c8175175d1f5d9650bee1b171cfdd5a638ab668facd6d6fc15ee564"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / afb3b8725b5b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-c886f8c1d3e8b8d65fcb36d13f73487a32ac5621465ee20f268f87e7e1b11526"></a>

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
enable_path_normalize = {}
```

<a id="canonical-72b99502f7a3d253312d27a65a0dd6ca7965edcd5d658e0527b33ef6cc2c000e"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / afb3b8725b5b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5a7d4f83fbdd66f92bbe7548fcb6a87b151d5233e84254832463b2f51d7faf4"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / afb3b8725b5b / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70bd32503ed3b74e5dc217fe295fdb9ecf5e070c0b9eb2a030d4d347a43f4ea4"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 13bf37d218fb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-9c64261903899bd0c1dd0dd178ae7f5819b7b47b335b319e1350f7c6d2fae0a8"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-f5e56869959c6f07e32e06325280ba8a0b6ddbd7677c55b7c28d705586beb469"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 13bf37d218fb / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-b1e69cf1f2938816ef417b5627d2b1993f4fd4f6df0aaa1d6a18623e162a264d): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-019.md#canonical-b3b010507fffb43f08f0a7219d0435b0e7cb08b6586f6ec8957ab2efe8b8030b): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-019.md#canonical-3214d305d90d1d3b93dd5c96c7aaafdbd4b0bc6e4a19e12e99c2c2ab992c95a0): complete subsection reference.

<a id="canonical-3d64d55ea46e3622debad607bb9d797d632c1d6409c09c66949357c652de8312"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 13bf37d218fb / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-b1e69cf1f2938816ef417b5627d2b1993f4fd4f6df0aaa1d6a18623e162a264d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-019.md#canonical-b3b010507fffb43f08f0a7219d0435b0e7cb08b6586f6ec8957ab2efe8b8030b)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-019.md#canonical-3214d305d90d1d3b93dd5c96c7aaafdbd4b0bc6e4a19e12e99c2c2ab992c95a0)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b1e69cf1f2938816ef417b5627d2b1993f4fd4f6df0aaa1d6a18623e162a264d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-596487199707b243747e290bb70554f7c0232915e1b91af5b364849588c17177"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f578e8f95df8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-f3e45b92980c6e3a561312ab312f589996c581174aa27c82a4ca086d0a20af40"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
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
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-7bc48161c5fb9a84592f21328ded292c12c679094d83eb201740296943ba0baa"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f578e8f95df8 / 3

- [header_transformation](resources--workload--reference--group-019.md#canonical-f0e94b853496cfda9176d762c67bd8218d00dc3ac468b5bec806ad02bef721f2): complete subsection reference.

<a id="canonical-550c98a53c9969193a52e67b08ebdf96d92f413c747b84b7c3c221dfe35bda45"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f578e8f95df8 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-019.md#canonical-f0e94b853496cfda9176d762c67bd8218d00dc3ac468b5bec806ad02bef721f2)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f0e94b853496cfda9176d762c67bd8218d00dc3ac468b5bec806ad02bef721f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7711ea82523fd7b442385d0ce5dd10178ab1851642d926f834b4efa719145956"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / aa7d325c186a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-b1e69cf1f2938816ef417b5627d2b1993f4fd4f6df0aaa1d6a18623e162a264d)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-1cd9d96d52d0727b2a9ead5a6c69f12e2662ba1e0271291ee580d0d9d108ce05"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-0529035d99925a8b3b6f4e6a08d3d9842550db50e6d709e7a79eaaa14349edef"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / aa7d325c186a / 3

- [default_header_transformation](resources--workload--reference--group-019.md#canonical-61c3b42f0cecac89ad513144f0ad76f27ebfc1438752c743a9d36ceb67c331fe): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-019.md#canonical-f4fc16ca73e621263787113bd06f406d28958bae2e6a5dad7965e94b193006c5): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-019.md#canonical-cf6e83a22631ba32f26ba279a9945e1bc9c3f5168134b253dd6342a2a4344b3a): complete subsection reference.

<a id="canonical-7dea77c677dd099f795c1c850aee77af41ff6779a85c285ac7e896c56d4a9d25"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / aa7d325c186a / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-019.md#canonical-61c3b42f0cecac89ad513144f0ad76f27ebfc1438752c743a9d36ceb67c331fe)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-019.md#canonical-f4fc16ca73e621263787113bd06f406d28958bae2e6a5dad7965e94b193006c5)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-019.md#canonical-cf6e83a22631ba32f26ba279a9945e1bc9c3f5168134b253dd6342a2a4344b3a)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-b1e69cf1f2938816ef417b5627d2b1993f4fd4f6df0aaa1d6a18623e162a264d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-61c3b42f0cecac89ad513144f0ad76f27ebfc1438752c743a9d36ceb67c331fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6930c76a8bd03da452bb5b49eec1f20ccf7b72c8e4a7490ed3e32802e77bedb"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 02d8b573bbaa / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-b1e69cf1f2938816ef417b5627d2b1993f4fd4f6df0aaa1d6a18623e162a264d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-019.md#canonical-f0e94b853496cfda9176d762c67bd8218d00dc3ac468b5bec806ad02bef721f2)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-ad3d3988e93d9bed327b5597a68f80a3b05715d194525cf57aaecd82aec828e2"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
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
default_header_transformation = {}
```

<a id="canonical-83391865db95064528b236d17fe1c0e9b85170a950b934b1415390d87f8127d8"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 02d8b573bbaa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a0801ab0b0f394dc13e4cdd7b7bbd8c7b097224b3e5f5951d4b168436cf9aa9e"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 02d8b573bbaa / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-019.md#canonical-f0e94b853496cfda9176d762c67bd8218d00dc3ac468b5bec806ad02bef721f2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f4fc16ca73e621263787113bd06f406d28958bae2e6a5dad7965e94b193006c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef6de39808330a8812c5ebc2cde77e93a360709bd1f53a605007c3fbdacecbc6"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f0d3c86f98e6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-b1e69cf1f2938816ef417b5627d2b1993f4fd4f6df0aaa1d6a18623e162a264d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-019.md#canonical-f0e94b853496cfda9176d762c67bd8218d00dc3ac468b5bec806ad02bef721f2)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-9756433b4b53e36dcb3d69d2bf08ed2d568d6efbb0d73c0a26a0ab23a04e4267"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
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
preserve_case_header_transformation = {}
```

<a id="canonical-49e1d3d3ceafc2634f7f58803eeb5ebc926d7150b569296890f8a6fa69941843"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f0d3c86f98e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-07b0fb133a4c8dd7ebd9ec5db0844b811142a814501d46284cdcb7be4e50b35f"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / f0d3c86f98e6 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-019.md#canonical-f0e94b853496cfda9176d762c67bd8218d00dc3ac468b5bec806ad02bef721f2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cf6e83a22631ba32f26ba279a9945e1bc9c3f5168134b253dd6342a2a4344b3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a4a67dbbbcd224afa8f04e3f1c658ecf5ffbe524a2c013071742f1770c4d1c2"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b8272c05fdaa / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-019.md#canonical-b1e69cf1f2938816ef417b5627d2b1993f4fd4f6df0aaa1d6a18623e162a264d)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-019.md#canonical-f0e94b853496cfda9176d762c67bd8218d00dc3ac468b5bec806ad02bef721f2)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-e816ba20f0e9b0cf273e0d465fa0d135f0d39d6e62131b9e92018f5d4193df41"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
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
proper_case_header_transformation = {}
```

<a id="canonical-d91879040ab0eb6d94926799358872736b9f627ebaa402924a4741f8366fb98b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b8272c05fdaa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4f7d59976779369777025c91cf4506ef82db874cdbee7fa50f18fbebd637e51a"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b8272c05fdaa / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-019.md#canonical-f0e94b853496cfda9176d762c67bd8218d00dc3ac468b5bec806ad02bef721f2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b3b010507fffb43f08f0a7219d0435b0e7cb08b6586f6ec8957ab2efe8b8030b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0a8891694582c47fa9bbe29b17e562f84bfb72156e9d8f4deb7305fffabf419"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a4c5e783201b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2c97418bff3dc925b8c11c36121764790c726810e62326c8bfb18f25f4218005"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-ee7363c54c2cfde19e479ca8eff42ae92e974502130e1094cf386a7ad7bd0c1a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a4c5e783201b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12a55118ca6d1b9c6288e53ee30eb098caf81e97bbb40afc34618e1d509e68db"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / a4c5e783201b / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3214d305d90d1d3b93dd5c96c7aaafdbd4b0bc6e4a19e12e99c2c2ab992c95a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9fe3b4dc8f1f68e4bd9b4bfdd28b84e1e9cf35a1bf65a264f6ab1bce9c1c25e7"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1acc4d474d3d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-cef0b9d79edf3863c5682726065c65a7c3d8e221f591591f3bf35022ae440528"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

<a id="canonical-82530c0ad4a436adebe84ccabc9ea0046994dfd23db9e3fb8b9af092103c775b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1acc4d474d3d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa7ac52393411552603041aa50cc05c92021082cb489bb03079e4cc4bf5275ff"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 1acc4d474d3d / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-019.md#canonical-a82d1eb959a595476a929efb701828f4a00d99554e80247332379348c7818bbe)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-53303c5fb3c4f861dd3167c2417f23beac9c602e00c753b327175b18ea910e71"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71916aad8a7c1c379411a8e935dec72f4c8e7d30be70500fd8aaacc81f619895"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c9dd876ee7ce / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-8b194bfc0c64d305a60fc1de311b155d22e2364a416d0adee4bf49d7f4b6b086"></a>

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

<a id="canonical-4fd0f9c8504812201a52202a3a9c3b7ff31e73782f9557c8ca51e340321fae28"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c9dd876ee7ce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-286f500a12344ff411bdea257b694ecb3e156d71048a97eaeaa594253659f1bf"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / c9dd876ee7ce / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3d0477a1d6976e1f9e8401dee1d41170755558b22c9c6025b4306887ef3f9ba5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb684b55aba921bbaa62602ff3a9a178487ad7ea73369d85ea5482b7c4094277"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / dba44fd4d561 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-ef264d65de3dfcb37614fb81f711e787f9ddae08519ab764dd74346457d3b841"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

<a id="canonical-92a22edfc4e1d39956e09246c0ff18de76aa193e9ab79c9c89fe63c7e7ba8cf4"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / dba44fd4d561 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dffb779b6a27427cb35365b0fb0ea23a861f52b6a3386b9b4ce6fd71442a7fe3"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / dba44fd4d561 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c3f6ce13d0fb0a9a0c627c7cf998643718eeef204cf8557b85d4bfef6fda281b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103e8f904ef84efae3f61365640f6499cc1d84bc17bf07a3334233ebbfa4a75"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 2b1c4da12221 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-7b2e2a6565c39c1564e1b2201dd5559bd268bad44f7aa7e8db94f3b0033ea166"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

<a id="canonical-76e89c4c8e529e53efc0e2cb97f95f4cfdd09e6e06c2dd7b77f276c643cd48d8"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 2b1c4da12221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bbda4657ce53f36b1e7611e55b1f41c3e54a7510f37829b3346aac4caa929850"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 2b1c4da12221 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d036da666bf50d78482211dd9755254574a3a7061ae84af4319a77859ad15c37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f998d7c45fcb28711f790ca0c22167f93243fdba9150dbf2fc5a1b6b5958e1c"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 827026a529ae / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-59b876e9674411120f4632589e541b44d5a7fbc7b7282bae537b331bc7422f62"></a>

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

<a id="canonical-d68a61aaaba6e175ab361d2ba358a1f4ad0ef1fd74cec9b52b6bb2f5310e75a2"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 827026a529ae / 3

- [custom_security](resources--workload--reference--group-019.md#canonical-63965b14c9a4a25791f4cfbc1885df3e5b3a661e711dd054c21061fe8aaecaca): complete subsection reference.

- [default_security](resources--workload--reference--group-019.md#canonical-1f58b76751ee30950edc9f675cbe935b8492bb922230c16dfd37a294404babce): complete subsection reference.

- [low_security](resources--workload--reference--group-019.md#canonical-dacb2ad49278af7313145792cb5589df90b81cc4715050b829fec9d553caa56c): complete subsection reference.

- [medium_security](resources--workload--reference--group-019.md#canonical-047025bc2cfeb23ecace28e37dcdadf8f32c43735654b970144b8418aa96a3ae): complete subsection reference.

<a id="canonical-544afb76a3f73aac9866088c74960896c1c25f1a5e120e9c3fcfdd411a253ed0"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 827026a529ae / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security](resources--workload--reference--group-019.md#canonical-63965b14c9a4a25791f4cfbc1885df3e5b3a661e711dd054c21061fe8aaecaca)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security](resources--workload--reference--group-019.md#canonical-1f58b76751ee30950edc9f675cbe935b8492bb922230c16dfd37a294404babce)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security](resources--workload--reference--group-019.md#canonical-dacb2ad49278af7313145792cb5589df90b81cc4715050b829fec9d553caa56c)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security](resources--workload--reference--group-019.md#canonical-047025bc2cfeb23ecace28e37dcdadf8f32c43735654b970144b8418aa96a3ae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-63965b14c9a4a25791f4cfbc1885df3e5b3a661e711dd054c21061fe8aaecaca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9453c6847b83845a48fb7377f9a12ce09f8a50d47bc347db074ff23984575866"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5106cfbca91e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-019.md#canonical-d036da666bf50d78482211dd9755254574a3a7061ae84af4319a77859ad15c37)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-d30ac58f19576226bc848ba7d74fb66c1a1c36872114234808fa92428142fc63"></a>

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

<a id="canonical-7d5704423c37a2355e208a9462f7d95a963759cb9ab282fdb4f9f9b70d5d6ed1"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5106cfbca91e / 3

<a id="canonical-a51c2f18f8d7a64600e8534686a58445e72c148302d6cc4d1a64e76edd4b7b49"></a>

<a id="canonical-7342a8a284e322f3f8294b2c6ca0c912b37cac7660fc8f7cc43494cd4d974103"></a>

## cipher_suites property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5106cfbca91e / 4

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

<a id="canonical-a5b92ed48116fa9b91a9a5270923e343e4272000c3914f5d3e539ab76a25ba4e"></a>

<a id="canonical-a7f376c3fe350950671103ccc3646ca3ccc4c8c47f99369f60f0ef74a3ccbc18"></a>

## max_version property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5106cfbca91e / 5

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

<a id="canonical-1ccfc009742538ffc08d8d84434589b9f62fcc4e480ef01774534de90fdeb1c9"></a>

<a id="canonical-abf18e64262214146978d27dc8c6bb6cf92721602b8f88f7ba55e5c517812fb4"></a>

## min_version property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5106cfbca91e / 6

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

<a id="canonical-ac470bc21c122e04779a38ac4995f1e9e3a7d74135baeff1066cdaad703d00a2"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 5106cfbca91e / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-019.md#canonical-d036da666bf50d78482211dd9755254574a3a7061ae84af4319a77859ad15c37)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1f58b76751ee30950edc9f675cbe935b8492bb922230c16dfd37a294404babce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1290bf151fe7ea2eab056204dd3123232c36b7575f27d25e39d1881147eac996"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 60eba0ed254d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-019.md#canonical-d036da666bf50d78482211dd9755254574a3a7061ae84af4319a77859ad15c37)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-53e03d11f932a7db83222475ed25940222710a3205b0568c53b98fd4b14974de"></a>

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

<a id="canonical-7e9c1d52d4a21dfb047f75d5b65ed343669da72c97651f97f9112007fa7300d3"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 60eba0ed254d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c8ba6736dfa0dacb664f6ccca1386a9b3e5cfd7ca4d1e1b27d03281ad705bb1"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 60eba0ed254d / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-019.md#canonical-d036da666bf50d78482211dd9755254574a3a7061ae84af4319a77859ad15c37)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-dacb2ad49278af7313145792cb5589df90b81cc4715050b829fec9d553caa56c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea9936d8521b7f54b8bf3563b18a2febcb7950bf7eb6be5bca6938fe90b5ac7d"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / d54ce9ce3ca8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-019.md#canonical-d036da666bf50d78482211dd9755254574a3a7061ae84af4319a77859ad15c37)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-ee2bb4983772b4b979552a85c0eefecb1412aed06c75b3b5a0aa569f26b8fbb3"></a>

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

<a id="canonical-7a1118899ca05893ff131252c1485a2589eb1d9dc99ddc06f9e49466ca79a66c"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / d54ce9ce3ca8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0cbfdbbcb19d6c1aef6501758b83947cad7768fdd50ab70983a88c8b2bf53f2c"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / d54ce9ce3ca8 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-019.md#canonical-d036da666bf50d78482211dd9755254574a3a7061ae84af4319a77859ad15c37)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-047025bc2cfeb23ecace28e37dcdadf8f32c43735654b970144b8418aa96a3ae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76716db49adc07cffced802005aae08992c2b3b12d86fa4bd2f6493ae5c6f83b"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bc6daae9ac4c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-019.md#canonical-d036da666bf50d78482211dd9755254574a3a7061ae84af4319a77859ad15c37)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-7d59b186ab03143f2c1d148eba904d9fe598b619d7c3307112a99850ed472c75"></a>

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

<a id="canonical-365a41214464f405188993459a7ef3a50bed0d6d518a1d9f96344fa521e8f048"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bc6daae9ac4c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-23a7cdc1b5f3cc4ac831def2b498c690bfd97ffd8edf74144d3e76b14952ccf4"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / bc6daae9ac4c / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-019.md#canonical-d036da666bf50d78482211dd9755254574a3a7061ae84af4319a77859ad15c37)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a23be27cb0392f77fbe444227dc47b42cb0d5464dec6ec2bb0bf40db059efed"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b9e23a8eba32 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-e8c6e44ac9c5dfff3a09e620f14d23109a350d586b0350dce0ccab03731fcf74"></a>

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

<a id="canonical-50bce1d2a2da20d9a0b059fcfed31ab50107cbeab10b4458988689b789ef4db9"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b9e23a8eba32 / 3

<a id="canonical-16fab4ce8389d205de98e775c425ae89f2a61834ecf4bc10591d00cf42ef3ea1"></a>

<a id="canonical-ea2d5389edfd7adbc53dbc050bd0c9ccbb3fbd00ffe6e0a41ab0237597fad988"></a>

## client_certificate_optional property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b9e23a8eba32 / 4

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

- [crl](resources--workload--reference--group-019.md#canonical-d71e790fc3f21ef8b91d46a93f1e31ad81496e3506fc951a95e61d10f81c82e3): complete subsection reference.

- [no_crl](resources--workload--reference--group-019.md#canonical-fa207f17fcd32a38ef5e8e047b8fc88adbd4a918c97e3a61d90e34235ecc38f0): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-019.md#canonical-8359959b62320b564a8501539d6bb0f5614d32014b8640c91bc4220d3b55913b): complete subsection reference.

<a id="canonical-0300539d8ab57a23f8d4eca387bb792b60f77628c3f719310106cc2cd0b0f746"></a>

<a id="canonical-fc076c5a4de289d6ca61c4fc5929ca95e72e33007d63650feb51a44627fc4446"></a>

## trusted_ca_url property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b9e23a8eba32 / 5

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

- [xfcc_disabled](resources--workload--reference--group-019.md#canonical-93fa8573efe9d45ea29a4f7802e2fdfbee90e80a3bcbf63f86a927a405a3f2e0): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-019.md#canonical-3519dc39c96d2dc6adb3f47f8245accfaa90235a0cca0d0c8b3f99ecb2a78c11): complete subsection reference.

<a id="canonical-40abf9589bfab1cd2f43caeac32e53db1ec8267c8a067c5a0739672bdcc631a8"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / b9e23a8eba32 / 6

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl](resources--workload--reference--group-019.md#canonical-d71e790fc3f21ef8b91d46a93f1e31ad81496e3506fc951a95e61d10f81c82e3)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl](resources--workload--reference--group-019.md#canonical-fa207f17fcd32a38ef5e8e047b8fc88adbd4a918c97e3a61d90e34235ecc38f0)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca](resources--workload--reference--group-019.md#canonical-8359959b62320b564a8501539d6bb0f5614d32014b8640c91bc4220d3b55913b)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled](resources--workload--reference--group-019.md#canonical-93fa8573efe9d45ea29a4f7802e2fdfbee90e80a3bcbf63f86a927a405a3f2e0)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options](resources--workload--reference--group-019.md#canonical-3519dc39c96d2dc6adb3f47f8245accfaa90235a0cca0d0c8b3f99ecb2a78c11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d71e790fc3f21ef8b91d46a93f1e31ad81496e3506fc951a95e61d10f81c82e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0261945697201e259122360e083e5cf4d5fb96876889e2025c8d12319abd9100"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 75d60feab3c4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-b25350ab98109002a5ccbaa2ccde23c569b55e37750d7e74e07ec774ef4b952e"></a>

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

<a id="canonical-83a15111c8a200c4509b6070e0c84115b748a82d9033f0d78030801a33f878a8"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 75d60feab3c4 / 3

<a id="canonical-d5c000224db55dbe66282ef1f121295bca955551956bbed0b75ea36ace708890"></a>

<a id="canonical-4db2525b9aa5743de3451c285ea8e4973b857f84d4813f6588efbaa6f88bfd98"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 75d60feab3c4 / 4

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

<a id="canonical-a0821e86dbc93781a964ba8ce5a8ccb07a59c5b84707da6b71493f0e130a4261"></a>

<a id="canonical-4621f55d583e979d5e9f71d4f24779ee698a8edf34198caf526dfc912780eac7"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 75d60feab3c4 / 5

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

<a id="canonical-27ab0fda87ea69b774d6cbaa6e42b8b259f428936e689afadcc5efc37c69dc69"></a>

<a id="canonical-356eff4dc336582f84f65b4d2cc034487287ccfdac944b90edbf5d54cefab7cc"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 75d60feab3c4 / 6

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

<a id="canonical-ddb12d36c6b059f8a09f5d670a6863d4fbd551a11c0b430a5947590a24eae8c3"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 75d60feab3c4 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-fa207f17fcd32a38ef5e8e047b8fc88adbd4a918c97e3a61d90e34235ecc38f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0c05c27e7996868fbe9d1fbbeb2b6caf56860def2b66762f2525695fb19afda"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4a17b0d14501 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-a070a4510b0f6f4d99d84e39f8865fac613637eb3447e233f863df6c69d9fb32"></a>

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

<a id="canonical-ce3e5e99d402cc90d5038251aa5737e0bb68ccd7ca52ad6a3689efcb5b8af1ea"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4a17b0d14501 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-366b166e76e2f2dc7fa18420583e970c472db57badba5f8e531bbcfe6e1fdb40"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 4a17b0d14501 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8359959b62320b564a8501539d6bb0f5614d32014b8640c91bc4220d3b55913b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35e35f381e9ae8f5a1a81a391668bff17e945723b586b488420412fa0c4476fb"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 58a938752209 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-9c6d79a7435cdff03bfc44cc48d68ebc4aa8562457b0fe10dec79da9020ec56c"></a>

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

<a id="canonical-060922f913829f1dbaa35d30eb459ffae87f9f911626bd9feac5b359fcb00375"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 58a938752209 / 3

<a id="canonical-0ca5488ef0f04d2ba8475d4ffc8002875e6f19789da34be3d3e378dabac580c6"></a>

<a id="canonical-128eb45e2c3c5ed995644369e34906b31d01e5e677cf22174d35f8c35353cc6c"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 58a938752209 / 4

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

<a id="canonical-b4bcd799ca7ec75c018b387d46a9e08be04b56476ebc3596d050d508ee5c3613"></a>

<a id="canonical-885539ec05b16749f477727282968913ec744b29b6b47f9fa20713e3fd1c8cb9"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 58a938752209 / 5

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

<a id="canonical-03cf6e2de8f70e67feaf7bf320fa67272a3758699979199cab2117257feacb13"></a>

<a id="canonical-c828c5d8240539ba8082c6fda6845bd6bb65a1ab909be9dc2241dbc6b0b5cfde"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 58a938752209 / 6

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

<a id="canonical-4ba320e1ecdd3736d995f53f842af8036d7d6b02d866acfe0c4eeeed51e6d9bd"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 58a938752209 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-93fa8573efe9d45ea29a4f7802e2fdfbee90e80a3bcbf63f86a927a405a3f2e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7efb421d6dda19249e314d7622e955396984d41a66bc8ceab8350dc6344fe3c5"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 77a6e1d37d84 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-0065efa3a17898eeb8300c063420adb43e0ee525a55c22b22d0f82e41b7ee061"></a>

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

<a id="canonical-819601e18c3c407bc776d4763adc2745bb56c8c73e96447dcb38e3e993e23527"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 77a6e1d37d84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ae58f8a87c93d1fe9b7ddc51b6ec79ced55ccb1dec29387dcc586f46913e3bf"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 77a6e1d37d84 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3519dc39c96d2dc6adb3f47f8245accfaa90235a0cca0d0c8b3f99ecb2a78c11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20c638454f3012d063f7b668d7754867fb02ee910a30e36e88d9dea414ffec81"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 3a53ccb17136 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-ce66ed1d7b48dc7c09733c25f662bdffbe9a3d5d17d8942dcd5ecabb4e3fbe11)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-60157fc12996830cfb2b640b32e93ad35b17390a2d9eb2d405b541a76182d54c"></a>

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

<a id="canonical-c79dca94fd6b25420ed8d88cfff4873ecf14f2f690528167c771e9c221634d5b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 3a53ccb17136 / 3

<a id="canonical-0883f4e3723e864826f2a616728991335f16f78140a2f7c83aeadef0e59b99e2"></a>

<a id="canonical-74a875d70b1c00be9d7d0f05cb69b154c0903823f9b7269c3647d1250c0c4e2f"></a>

## xfcc_header_elements property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 3a53ccb17136 / 4

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

<a id="canonical-1f574f016eb8f023830b1576ea03e78a09d9e600f6953b8a6d74490f77e43211"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http / 3a53ccb17136 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-019.md#canonical-71c9401151295115396c7974d0a39944cbc3d0404f597dd7aef7a0ab825126a8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9228d4262aca723705909584c97e7ed029c79c079822a17016401f1654b6381b"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 997543bcccca / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes

<a id="canonical-cd0194705280fb7d2e14bb35463b473dd39878d7ed4be1e7e53e4dfa73439bfa"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

Receipt-pinned upstream constraints:

```json
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
specific_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-9990590f44fd4915d2f63b32bfb82a04412fcdeb271d7bfe9ce6869345f13e92"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 997543bcccca / 3

- [routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037): complete subsection reference.

<a id="canonical-4d92edf985671fe39b4bdac976db2c7d466a1be846b9978172f5bf8f97120f08"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 997543bcccca / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-554c09b11ef703aad34359d39e3f10cd2cf6a6bf2d06dc2dfe21f05467fb1b9a"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b76f43279b34 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

<a id="canonical-b1b876b64e010438ee74b1726404f2e4cb5dcbcd66b77f817df681ddd0d370cf"></a>

Type: `"object"`. list nested block, Optional.

Routes. Routes for this loadbalancer.

Upstream description:

Routes for this loadbalancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_route_object",
    "direct_response_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "simple_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "simple_route"),
  validators.ConflictingListObjectAttributes("redirect_route",
    "simple_route")}
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```
