---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-89bb79ec2d51bd9eebc406c2d03c5e3f8f509486612ad74b0bc763d9a395b83c"></a>

## sub_path property — stateful_service.configuration.parameters.file.mount / b73dd62b1542 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-b6d7699942f5d6bb1905cb29c166831156489a9119135c80609d7c58689755a2"></a>

## Next pages — stateful_service.configuration.parameters.file.mount / b73dd62b1542 / 7

- [stateful_service.configuration.parameters.file](resources--workload--reference--group-027.md#canonical-d57d34733bca8ce08dfef6602adde30afe5bad86783c4f2e32979607cf590be4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-480e2a19e6cba2067fafc92d81ee03af3e8b7d0e05f30f9fd4766359d06c67f7"></a>

## stateful_service.containers — stateful_service.containers / 250138b6710e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- stateful_service.containers

<a id="canonical-5730314004728fd909647cbfaee08cb70feb8e372af10d991c6d208cb061fcaf"></a>

Type: `"object"`. list nested block, Optional.

Containers. Containers to use for service.

Upstream description:

Containers to use for service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("custom_flavor",
    "default_flavor"),
  validators.ConflictingListObjectAttributes("custom_flavor",
    "flavor"),
  validators.ConflictingListObjectAttributes("default_flavor",
    "flavor")}
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
containers {
  # Configure direct properties listed below.
}
```

<a id="canonical-c56fe60b13ffa1f07d87b5d7cafb40f9b02e36d0719b74633d595052e2d61592"></a>

## Direct properties — stateful_service.containers / 250138b6710e / 3

<a id="canonical-3282366df303bd34524a6e2f848495316902aa7698d14c96afacb5025d6bae78"></a>

<a id="canonical-87f1a04296ded77fb55a714a404c95276fa7e6538ccae8325753fc0f5d73463b"></a>

## args property — stateful_service.containers / 250138b6710e / 4

Type: `["list", "string"]`. Optional.

Arguments to the entrypoint. Overrides the docker image's CMD.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-240c3ebe23d4bf31281d0d586b7739f690014b8b834131d4f522c620210444f9"></a>

<a id="canonical-1deac7b8d9a5a9b1450ce28f67401bccbd67d6811401c6364a5539079159f480"></a>

## command property — stateful_service.containers / 250138b6710e / 5

Type: `["list", "string"]`. Optional.

Command to execute. Overrides the docker image's ENTRYPOINT.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](resources--workload--reference--group-028.md#canonical-e7c8724220407cd29fb5712c06e0dc981526b9bdab498d7552f2be4105045f5a): complete subsection reference.

- [default_flavor](resources--workload--reference--group-028.md#canonical-a8efb777e7aeb9924b11109b4492470ae711c10b005ca55e6a7b7f70838b3051): complete subsection reference.

<a id="canonical-c9bf93f9d15de9c2b441bdd5087a90741c4342ec90061a3199eb88555f23a48d"></a>

<a id="canonical-3f1c94b71086e7b9315de387b363f0d12a2cfaaaa2e2f30ae5a056f64d303076"></a>

## flavor property — stateful_service.containers / 250138b6710e / 6

Type: `"string"`. Optional.

\[Enum:
CONTAINER\_FLAVOR\_TYPE\_TINY|CONTAINER\_FLAVOR\_TYPE\_MEDIUM|CONTAINER\_FLAVOR\_TYPE\_LARGE\]
Container Flavor type - CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny Tiny containers have limit of 0.1 vCPU
and 256 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium Medium containers have limit
of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_LARGE: Large Large containers
have.. Possible values are \`CONTAINER\_FLAVOR\_TYPE\_TINY\`, \`CONTAINER\_FLAVOR\_TYPE\_MEDIUM\`,
\`CONTAINER\_FLAVOR\_TYPE\_LARGE\`. Defaults to \`CONTAINER\_FLAVOR\_TYPE\_TINY\`.

Upstream description:

Container Flavor type

&#8203;- CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny

Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium

Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_LARGE: Large

Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTAINER_FLAVOR_TYPE_TINY",
  "enum": [
    "CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [image](resources--workload--reference--group-028.md#canonical-2472577d9d501e0a7f71d3c27e070cf143ceab414de2c94bd70cd0941a8648f2): complete subsection reference.

<a id="canonical-3ce58518b2d512aff23bfb96326ee9e2a08b5817e8b8c75bde714b3ffd158552"></a>

<a id="canonical-26cb27c2a77dab99bcc1539181431bb05a37c2390f18fb385c003f40e0a85827"></a>

## init_container property — stateful_service.containers / 250138b6710e / 7

Type: `"bool"`. Optional.

Specialized container that runs before application container and runs to completion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [liveness_check](resources--workload--reference--group-028.md#canonical-c885fe15ff1aa1344804fe7f8dffbd8c3a259c1a14f1fde14028dd39394bce95): complete subsection reference.

<a id="canonical-efb963cfef0efd57d4dfc9f0ae6e00727cd564dccca428f182cf1fcab9fd97a6"></a>

<a id="canonical-bf92121d9e7971e8a2f7fa66e988d2223b1147f95acbfa848f470775163818c6"></a>

## name property — stateful_service.containers / 250138b6710e / 8

Type: `"string"`. Optional.

Name. Name of the container.

Upstream description:

Name of the container.

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

- [readiness_check](resources--workload--reference--group-028.md#canonical-91bc45320ba6ff7008181894d544fbf7ddcbbf8937e3b3774d41073f5d7a73b9): complete subsection reference.

<a id="canonical-acf81b46f015814420f7800d3f929a6ba3ea9f8dc6427fc434a1c3657787fc7d"></a>

## Next pages — stateful_service.containers / 250138b6710e / 9

- [stateful_service.containers.custom_flavor](resources--workload--reference--group-028.md#canonical-e7c8724220407cd29fb5712c06e0dc981526b9bdab498d7552f2be4105045f5a)
- [stateful_service.containers.default_flavor](resources--workload--reference--group-028.md#canonical-a8efb777e7aeb9924b11109b4492470ae711c10b005ca55e6a7b7f70838b3051)
- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-2472577d9d501e0a7f71d3c27e070cf143ceab414de2c94bd70cd0941a8648f2)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-c885fe15ff1aa1344804fe7f8dffbd8c3a259c1a14f1fde14028dd39394bce95)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-91bc45320ba6ff7008181894d544fbf7ddcbbf8937e3b3774d41073f5d7a73b9)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e7c8724220407cd29fb5712c06e0dc981526b9bdab498d7552f2be4105045f5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95b699141d7790bf9ce6188c8d4996a9f05fc3e4dbd50e23e2d9d3a0de516dc5"></a>

## stateful_service.containers.custom_flavor — stateful_service.containers.custom_flavor / 386543b734e8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- stateful_service.containers.custom_flavor

<a id="canonical-cfb6f7185efa56637eeb92ec3923783c7e40775800ad07635fdf8947fd996001"></a>

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
custom_flavor {
  # Configure direct properties listed below.
}
```

<a id="canonical-b04116a5b5235afb67146afae60e320f3bdbf4167c493db8b46a03da197781e6"></a>

## Direct properties — stateful_service.containers.custom_flavor / 386543b734e8 / 3

<a id="canonical-dcdd6341fe78247e4158882846ba9c5b5536062a41559f2f7276b2b0200bf46e"></a>

<a id="canonical-99a4095161a220d85ccd7fe42402e49e4f3af51e6a6b786fcb0ebcdc0d415ab0"></a>

## name property — stateful_service.containers.custom_flavor / 386543b734e8 / 4

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

<a id="canonical-d4b889d54a5fa64b99196d2bd786f053fb14fa83449ff10fdf7f8dd499c1dec4"></a>

<a id="canonical-430d603a034502337fee39d66dfc5c7d993fa4b8bad17bd005ed542c9e69c591"></a>

## namespace property — stateful_service.containers.custom_flavor / 386543b734e8 / 5

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

<a id="canonical-ba9d1527ebe2fe0db8fe964590d467f5a788ac3576133d4ad33fb90ba27ca35a"></a>

<a id="canonical-2299e52dabef4d36e1aba9449323c5e70827a2105b7b68e29203cbbd2baaff5e"></a>

## tenant property — stateful_service.containers.custom_flavor / 386543b734e8 / 6

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

<a id="canonical-adf1bfff03cc5f7a828d0ed5e44bad5c957fdd75a23e599de27bf4d72308c844"></a>

## Next pages — stateful_service.containers.custom_flavor / 386543b734e8 / 7

- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a8efb777e7aeb9924b11109b4492470ae711c10b005ca55e6a7b7f70838b3051"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eff97758665ec37031ac57411eb1674c0f54c1b1e0e44f9c44361cef882f2a4c"></a>

## stateful_service.containers.default_flavor — stateful_service.containers.default_flavor / b5a9ddecdbba / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- stateful_service.containers.default_flavor

<a id="canonical-ef9712e073f0d124b3cc8e60d83bd1fdfccb02fe6c87ab83c663024c1c53f4f4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default flavor.

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
default_flavor = {}
```

<a id="canonical-75a74fc6ab03a477dc495fb4f35518a1cefb8ab69636ed1327b24006a078c158"></a>

## Direct properties — stateful_service.containers.default_flavor / b5a9ddecdbba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-292c8f290528ab1945dadd51568d90d322e3695e03362cf9672de139da0d49be"></a>

## Next pages — stateful_service.containers.default_flavor / b5a9ddecdbba / 4

- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2472577d9d501e0a7f71d3c27e070cf143ceab414de2c94bd70cd0941a8648f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc9bfc29a45e3cde0dafd637fa741c1a1e7bd2a498e2ed0d7e3d3795fc8013a6"></a>

## stateful_service.containers.image — stateful_service.containers.image / 4c4d028c047f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- stateful_service.containers.image

<a id="canonical-c12ddf26cc7d33efd213470b6f051a90a598fee46ba189adf804c2713cd6a566"></a>

Type: `"object"`. single nested block, Optional.

ImageType configures the image to use, how to pull the image, and the associated secrets to use if
any.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("container_registry",
    "public")}
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
  "x-ves-oneof-field-registry_choice": "[\"container_registry\",\"public\"]"
}
```

Terraform syntax:

```terraform
image {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea3bcd2dc91b7b74b4f370153d719d9dfc29a601e135960d49a801c6137a567a"></a>

## Direct properties — stateful_service.containers.image / 4c4d028c047f / 3

- [container_registry](resources--workload--reference--group-028.md#canonical-745fd6c697ac66e4c83e47bc9c8500b1cbd87cf7d2bf4be37d9563695088d777): complete subsection reference.

<a id="canonical-77781c48dddeb5fc611decc5176d8b34e682aef3bc62a6f46211ef1c26563bae"></a>

<a id="canonical-cf6127e837a5972db7e7a075c6f76ae15c84d4c26aac90f57d24e9eec508df4b"></a>

## name property — stateful_service.containers.image / 4c4d028c047f / 4

Type: `"string"`. Optional.

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed.

Upstream description:

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.

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

- [public](resources--workload--reference--group-028.md#canonical-448916f32fcfe60ccda1bf4da1c44234980967bc2cb19d61e996bb972ddc591f): complete subsection reference.

<a id="canonical-e44a3ceb665de3f0e3ed61a7a0e2ef64702f68ae2879aa4c8833db2a72fd199a"></a>

<a id="canonical-bf4f1323cb81e947ebeafee23a81ef08a1626e5d752798ad34bec666312009a4"></a>

## pull_policy property — stateful_service.containers.image / 4c4d028c047f / 5

Type: `"string"`. Optional.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Upstream description:

Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload

&#8203;- IMAGE\_PULL\_POLICY\_DEFAULT: Default

Default will always pull image if :latest tag is specified in image name. If :latest tag is not
specified in image name, it will pull image only if it does not already exist on the node &#8203;-
IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT: IfNotPresent

Only pull the image if it does not already exist on the node &#8203;- IMAGE\_PULL\_POLICY\_ALWAYS:
Always

Always pull the image &#8203;- IMAGE\_PULL\_POLICY\_NEVER: Never

Never pull the image.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "IMAGE_PULL_POLICY_DEFAULT",
  "enum": [
    "IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-8ae099c2da989359e5f151d26795dc1c08bd5bed87602d27f89d422df3744b69"></a>

## Next pages — stateful_service.containers.image / 4c4d028c047f / 6

- [stateful_service.containers.image.container_registry](resources--workload--reference--group-028.md#canonical-745fd6c697ac66e4c83e47bc9c8500b1cbd87cf7d2bf4be37d9563695088d777)
- [stateful_service.containers.image.public](resources--workload--reference--group-028.md#canonical-448916f32fcfe60ccda1bf4da1c44234980967bc2cb19d61e996bb972ddc591f)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-745fd6c697ac66e4c83e47bc9c8500b1cbd87cf7d2bf4be37d9563695088d777"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a441a4c790728c97d35d349ee7258e7263dede04b13fc99c82d5dc7608e4a36"></a>

## stateful_service.containers.image.container_registry — stateful_service.containers.image.container_registry / 805a0d7de2c1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-2472577d9d501e0a7f71d3c27e070cf143ceab414de2c94bd70cd0941a8648f2)
- stateful_service.containers.image.container_registry

<a id="canonical-d14d7df8cf9cc58a358fefeb248536e56d9c7744fef4522d2c9d27d4d7b9101b"></a>

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
container_registry {
  # Configure direct properties listed below.
}
```

<a id="canonical-b72ac55ebdb902644988858af395da0b93d7efc93cdee8de3a1add427a073982"></a>

## Direct properties — stateful_service.containers.image.container_registry / 805a0d7de2c1 / 3

<a id="canonical-695ef58f1f5278e677b8cd6e67f09fb2095df805f817c0c335e8f8bfb4a616c2"></a>

<a id="canonical-482743a84be590302e6ca843174d788653af5cd0332d0d389ec446975337bd32"></a>

## name property — stateful_service.containers.image.container_registry / 805a0d7de2c1 / 4

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

<a id="canonical-3676f95f7c4791cb2b11c2f7f198242f2f91216d6e8397b89e1271f9371681e8"></a>

<a id="canonical-42f7aff450a05632fe559fd518e0411e9a2aca11a9de85aa756b4d63db2a997e"></a>

## namespace property — stateful_service.containers.image.container_registry / 805a0d7de2c1 / 5

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

<a id="canonical-2919c849ee5e605bf3fbf03a5ce58ba963efea99782c653a78a3b8210b272dc6"></a>

<a id="canonical-30b3f67d5b86492e6f9319e5814956c3d9a3aeab8f883c7a69ffa650f93ea053"></a>

## tenant property — stateful_service.containers.image.container_registry / 805a0d7de2c1 / 6

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

<a id="canonical-7179c79e2be7b4cc660845fe06c83e12cea5de94d621f8470524544554090b64"></a>

## Next pages — stateful_service.containers.image.container_registry / 805a0d7de2c1 / 7

- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-2472577d9d501e0a7f71d3c27e070cf143ceab414de2c94bd70cd0941a8648f2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-448916f32fcfe60ccda1bf4da1c44234980967bc2cb19d61e996bb972ddc591f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-15358c32dd8abb31bb7f004e3588f3a9f068bf9df9b0ceb7a615495561a4ac9d"></a>

## stateful_service.containers.image.public — stateful_service.containers.image.public / 6ef86df3d031 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-2472577d9d501e0a7f71d3c27e070cf143ceab414de2c94bd70cd0941a8648f2)
- stateful_service.containers.image.public

<a id="canonical-240eb6b63f8edc87d00e56ed16e20e5096f4050c1ddd6a25be4665dcbafba59d"></a>

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
public = {}
```

<a id="canonical-2d5f194a46e6889082f756b85ff5fdf27c2f6499753ef3d99e6f05344ba540a6"></a>

## Direct properties — stateful_service.containers.image.public / 6ef86df3d031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-779c5a0b8f46a78aed1a5b07b73edc6fecc6bf07f292ffecafda6154de7953a1"></a>

## Next pages — stateful_service.containers.image.public / 6ef86df3d031 / 4

- [stateful_service.containers.image](resources--workload--reference--group-028.md#canonical-2472577d9d501e0a7f71d3c27e070cf143ceab414de2c94bd70cd0941a8648f2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c885fe15ff1aa1344804fe7f8dffbd8c3a259c1a14f1fde14028dd39394bce95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fc27ac8a656950bacdf16385b9ffd15469f1ce402a6e02efa93da055c173606"></a>

## stateful_service.containers.liveness_check — stateful_service.containers.liveness_check / 3705d3a86a1d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- stateful_service.containers.liveness_check

<a id="canonical-382ff182b5949fecc55299235431e9325125b2d37682e9267eba7f25a1f5b067"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
liveness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-e5042d451ba9a5ed88505be6f225647420742d6f1cbbcda6ff47e930c46ffde0"></a>

## Direct properties — stateful_service.containers.liveness_check / 3705d3a86a1d / 3

- [exec_health_check](resources--workload--reference--group-028.md#canonical-e29499c704ac4a31d0178fa85fd8aa3cd298d99c301fd185f85477d5942d4ef5): complete subsection reference.

<a id="canonical-facd7d882b677c1504bcc28ac215589d72b0c2da7307ad0d350733bb4cf682c0"></a>

<a id="canonical-9106a04d55d0882437e8de4aa41b894181c3ebf66ebe90dfea2d581fe79e2f09"></a>

## healthy_threshold property — stateful_service.containers.liveness_check / 3705d3a86a1d / 4

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](resources--workload--reference--group-028.md#canonical-14df354021e2e480c27873fb009397b21cdda20458b61467bd3f201a8c618880): complete subsection reference.

<a id="canonical-1b55a0ff44bff328646252d28fe0db95ed4b7aa1197c0278a41162fd9a37f70e"></a>

<a id="canonical-7d4b3151677da7b6f5e1e3bc0ca526c32835904c2a8f5782799752885018a69f"></a>

## initial_delay property — stateful_service.containers.liveness_check / 3705d3a86a1d / 5

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-4cdce78315a9fa9845021edcbbb87cccb5d0f872740ffabec9be2b5c384eb81e"></a>

<a id="canonical-5a4c57bd5ae29b7c98492f75cc8cbc3902d1f7151ce09946eb4a66c89e12c153"></a>

## interval property — stateful_service.containers.liveness_check / 3705d3a86a1d / 6

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](resources--workload--reference--group-028.md#canonical-a58bfef463d73f588f7181e2203f6619cf9b5328fd728e0235ee9d150e16847e): complete subsection reference.

<a id="canonical-52b899aa4e3881f58d8f94915c3fc4abd4ee09021dcf238cc75e7945b9a7aeb5"></a>

<a id="canonical-93fe2a54546f214d1291444b3d27a1f562602f82eeeffe31e09c871dd84b556b"></a>

## timeout property — stateful_service.containers.liveness_check / 3705d3a86a1d / 7

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-9c3bd5be33589e2f99cc6beba89f721910ed29ed6a961b96c5be3f6b007e0bee"></a>

<a id="canonical-bceb93718d3578418d737a676183e28fc26bdefbf6400e7d09d33549f5e591b8"></a>

## unhealthy_threshold property — stateful_service.containers.liveness_check / 3705d3a86a1d / 8

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-b9a4e19c3ad955d4dbafafb045e28b15f515cccebe17eccf155f5d656c899708"></a>

## Next pages — stateful_service.containers.liveness_check / 3705d3a86a1d / 9

- [stateful_service.containers.liveness_check.exec_health_check](resources--workload--reference--group-028.md#canonical-e29499c704ac4a31d0178fa85fd8aa3cd298d99c301fd185f85477d5942d4ef5)
- [stateful_service.containers.liveness_check.http_health_check](resources--workload--reference--group-028.md#canonical-14df354021e2e480c27873fb009397b21cdda20458b61467bd3f201a8c618880)
- [stateful_service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-028.md#canonical-a58bfef463d73f588f7181e2203f6619cf9b5328fd728e0235ee9d150e16847e)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e29499c704ac4a31d0178fa85fd8aa3cd298d99c301fd185f85477d5942d4ef5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6ab4e49ae2b825a811a82bec245202cf5bf6a214c98ebb4823ae2dcfb2433e6"></a>

## stateful_service.containers.liveness_check.exec_health_check — stateful_service.containers.liveness_check.exec_health_check / 9bcc6ab42145 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-c885fe15ff1aa1344804fe7f8dffbd8c3a259c1a14f1fde14028dd39394bce95)
- stateful_service.containers.liveness_check.exec_health_check

<a id="canonical-2b6115e2cef701ff2b6ef649b4d6f9f33a16234f354a9716091dcdb0eedfa9db"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("command")}
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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-c76ca8238a0e71e94c3d365d2832c8497e15b282c22f4c5f69168490da4dff8b"></a>

## Direct properties — stateful_service.containers.liveness_check.exec_health_check / 9bcc6ab42145 / 3

<a id="canonical-2e89e29231afef5135aaeeb7a7c744516c1c95035965df05ed79936f32a2890c"></a>

<a id="canonical-4f61c7b0138a8db89184a321f9dc00a446d7e4a0c79f8c98e5491c68a377164b"></a>

## command property — stateful_service.containers.liveness_check.exec_health_check / 9bcc6ab42145 / 4

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-bcfc3f0ee476b40641c354909800b68aac3b6973c87b0afc429842cf9155e979"></a>

## Next pages — stateful_service.containers.liveness_check.exec_health_check / 9bcc6ab42145 / 5

- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-c885fe15ff1aa1344804fe7f8dffbd8c3a259c1a14f1fde14028dd39394bce95)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-14df354021e2e480c27873fb009397b21cdda20458b61467bd3f201a8c618880"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efe9f61c0cbf5c796ce125509bb0a0a3a5bd957d9e363673b851c23009d9afde"></a>

## stateful_service.containers.liveness_check.http_health_check — stateful_service.containers.liveness_check.http_health_check / bdaff6adde0d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-c885fe15ff1aa1344804fe7f8dffbd8c3a259c1a14f1fde14028dd39394bce95)
- stateful_service.containers.liveness_check.http_health_check

<a id="canonical-254491ffb127d487319de0189543aa3bbb675820b696b88bc589e907350f1fd3"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-54226be1825a290b725d6257bbd4508b6564330aecfa20ed27131e08b2a43cd1"></a>

## Direct properties — stateful_service.containers.liveness_check.http_health_check / bdaff6adde0d / 3

<a id="canonical-280c029ae6bc7db208d7b391e64862ee650573ac41ae70a6aebe2e9736fe0ebe"></a>

<a id="canonical-87beb6951fd788a4200649220d77b85f4b9c16fbb546bd74dc2f120879d54656"></a>

## headers property — stateful_service.containers.liveness_check.http_health_check / bdaff6adde0d / 4

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-e10cb7ff2325d2509713699121b5e3e9fd31a133d11999c45f6b84d318ad406e"></a>

<a id="canonical-4472a48aa5ca6dc6f9da98588e52188551fced0d71580ba97792661b5e6fcb33"></a>

## host_header property — stateful_service.containers.liveness_check.http_health_check / bdaff6adde0d / 5

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-570669d223c6d0de780ffd19faac7d200864038c24d85f36fecec872547aab12"></a>

<a id="canonical-6c8e84a15f8e3ebc508843ab1c15be7327afcf66b71d863381bb71686cb807b4"></a>

## path property — stateful_service.containers.liveness_check.http_health_check / bdaff6adde0d / 6

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](resources--workload--reference--group-028.md#canonical-33eb4c1f2fde7ad0460d1f4704d633f3924e465146407e1bf5bb5dee85e88a6f): complete subsection reference.

<a id="canonical-0ed0af89ec5074713f53a47eca51bdc59c97c4a49cb7cc2c646f77b879962936"></a>

## Next pages — stateful_service.containers.liveness_check.http_health_check / bdaff6adde0d / 7

- [stateful_service.containers.liveness_check.http_health_check.port](resources--workload--reference--group-028.md#canonical-33eb4c1f2fde7ad0460d1f4704d633f3924e465146407e1bf5bb5dee85e88a6f)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-c885fe15ff1aa1344804fe7f8dffbd8c3a259c1a14f1fde14028dd39394bce95)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-33eb4c1f2fde7ad0460d1f4704d633f3924e465146407e1bf5bb5dee85e88a6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5465be73d645e559a05fe5396f3da22b2a2920d8a679566c06e4d5d1efafa6c1"></a>

## stateful_service.containers.liveness_check.http_health_check.port — stateful_service.containers.liveness_check.http_health_check.port / 151a43dbf8f0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-c885fe15ff1aa1344804fe7f8dffbd8c3a259c1a14f1fde14028dd39394bce95)
- [stateful_service.containers.liveness_check.http_health_check](resources--workload--reference--group-028.md#canonical-14df354021e2e480c27873fb009397b21cdda20458b61467bd3f201a8c618880)
- stateful_service.containers.liveness_check.http_health_check.port

<a id="canonical-411c9f062a23263bbaf3ef476b7ea6142985540da1a47fbad623725c4590a342"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-d5754722738cca2955d6ef8c9fdfedd416ebf93658cc9d05b703ad3e2499f669"></a>

## Direct properties — stateful_service.containers.liveness_check.http_health_check.port / 151a43dbf8f0 / 3

<a id="canonical-3883a490847bb2bb3c7521e754a48c5677c773cb138482f28741f6d1b0405cfa"></a>

<a id="canonical-dbd4af32b404bbb17d00a0ece8bee8960a4d75f79dd2036606cff1ea0c882353"></a>

## name property — stateful_service.containers.liveness_check.http_health_check.port / 151a43dbf8f0 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-b9e92d970a70a8750bc816d53d2b7e17b4b8fb63f8c693f2d0d2ca57fc9efc8c"></a>

<a id="canonical-fadaed6e7499760e058cd398594c162e49973b8cf533cff7ba159fa62bd0e0b1"></a>

## num property — stateful_service.containers.liveness_check.http_health_check.port / 151a43dbf8f0 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-17496f0b58ebed51b34e96d75e34cea9db99484ba1a428923ed7fba9bddbd444"></a>

## Next pages — stateful_service.containers.liveness_check.http_health_check.port / 151a43dbf8f0 / 6

- [stateful_service.containers.liveness_check.http_health_check](resources--workload--reference--group-028.md#canonical-14df354021e2e480c27873fb009397b21cdda20458b61467bd3f201a8c618880)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a58bfef463d73f588f7181e2203f6619cf9b5328fd728e0235ee9d150e16847e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4684d3f11cb10a1ce102ae4276e2293b375154d893e505e5c3f569c323ec4f3"></a>

## stateful_service.containers.liveness_check.tcp_health_check — stateful_service.containers.liveness_check.tcp_health_check / d29503e63e76 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-c885fe15ff1aa1344804fe7f8dffbd8c3a259c1a14f1fde14028dd39394bce95)
- stateful_service.containers.liveness_check.tcp_health_check

<a id="canonical-0d9a6f42d4f3e2718926f51498b17de73022b82e8f3cefff9aaf6a71c80aa8bc"></a>

Type: `"object"`. single nested block, Optional.

TCPHealthCheckType describes a health check based on opening a TCP connection.

Receipt-pinned upstream constraints:

```json
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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-f5aaf5191b316ee6d21651834c940cecf40ae39ca8c05ecfb0f120d2d14c95da"></a>

## Direct properties — stateful_service.containers.liveness_check.tcp_health_check / d29503e63e76 / 3

- [port](resources--workload--reference--group-028.md#canonical-4bcf3fa36e96803fa6ca7a861bd7d7d3ccb1ef9f7efd57be15fdb3481c46800c): complete subsection reference.

<a id="canonical-95b77057de13c015d7b960e0b88045ae81fcd923672ef1981347820509e92dc7"></a>

## Next pages — stateful_service.containers.liveness_check.tcp_health_check / d29503e63e76 / 4

- [stateful_service.containers.liveness_check.tcp_health_check.port](resources--workload--reference--group-028.md#canonical-4bcf3fa36e96803fa6ca7a861bd7d7d3ccb1ef9f7efd57be15fdb3481c46800c)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-c885fe15ff1aa1344804fe7f8dffbd8c3a259c1a14f1fde14028dd39394bce95)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4bcf3fa36e96803fa6ca7a861bd7d7d3ccb1ef9f7efd57be15fdb3481c46800c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa17a9b1dd3eee7b1876f87fd0f4652691dbdcee109f7ecdf20b0a67789dc6e7"></a>

## stateful_service.containers.liveness_check.tcp_health_check.port — stateful_service.containers.liveness_check.tcp_health_check.port / 09c508887777 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.liveness_check](resources--workload--reference--group-028.md#canonical-c885fe15ff1aa1344804fe7f8dffbd8c3a259c1a14f1fde14028dd39394bce95)
- [stateful_service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-028.md#canonical-a58bfef463d73f588f7181e2203f6619cf9b5328fd728e0235ee9d150e16847e)
- stateful_service.containers.liveness_check.tcp_health_check.port

<a id="canonical-a285215b8bdf7129123c00db4fc3a1b71d4c0fa031f8bcd93cbed59a692c9271"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-36235ebc82eac8426066bbc9ce8a4dec0cfad26ef77348f02993f603ad11aec1"></a>

## Direct properties — stateful_service.containers.liveness_check.tcp_health_check.port / 09c508887777 / 3

<a id="canonical-f3605be54b18d6ed018faa0052b336d45d7fec66aabc606e4ebe4d8ffa7d455f"></a>

<a id="canonical-7c89aa7c894721331754ed48d430c586dd6d2c2484702eba9408543902abf5ac"></a>

## name property — stateful_service.containers.liveness_check.tcp_health_check.port / 09c508887777 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-898609a5f6a59979d753c5010d8e656be80a7b2f58cd49e95d24371dcb9b50bb"></a>

<a id="canonical-c2f8e0fa278f81ca62b161185f3aa8b6f2752a55fc34db7e16e8503b96d41a08"></a>

## num property — stateful_service.containers.liveness_check.tcp_health_check.port / 09c508887777 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-8d7cf3c409d29e05c81571159a2b814db0eb4ac660b47d971e08690b9246614f"></a>

## Next pages — stateful_service.containers.liveness_check.tcp_health_check.port / 09c508887777 / 6

- [stateful_service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-028.md#canonical-a58bfef463d73f588f7181e2203f6619cf9b5328fd728e0235ee9d150e16847e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-91bc45320ba6ff7008181894d544fbf7ddcbbf8937e3b3774d41073f5d7a73b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-460c5500cce9a685c3c1b1501b8987aa8a26488bd16c239a14e1f7e64d361872"></a>

## stateful_service.containers.readiness_check — stateful_service.containers.readiness_check / 86857ceda411 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- stateful_service.containers.readiness_check

<a id="canonical-3ba34c4cdf99f94ffee807d798cb6412b68309e198a0ac05e28a2f707882cae9"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
readiness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-ed6b383f429c4b4b446271dcc47fb025c83a221869957b200c19eb0ab6a36787"></a>

## Direct properties — stateful_service.containers.readiness_check / 86857ceda411 / 3

- [exec_health_check](resources--workload--reference--group-028.md#canonical-98d22817f1a2657d0a062babd6e1d210ee24d5d2a09ac8d3a08487727bd100f8): complete subsection reference.

<a id="canonical-a08949a404730396701ad5c287a32ba2e3d754ce1b3e218afc1792f9200a0201"></a>

<a id="canonical-b496b94d21932ab850cd10ae902ead5fb2ea5a8fa12cdc6d04db47f334d09257"></a>

## healthy_threshold property — stateful_service.containers.readiness_check / 86857ceda411 / 4

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](resources--workload--reference--group-028.md#canonical-d3dfe5ab9c3c8e563ac6e0a5abea26e4bdf1929a75e607cd312742a7fad60b01): complete subsection reference.

<a id="canonical-8eeef28095c76be6e88958a3b93d8740f7d0ff8cf5b2e82b156ceb9b55953b3d"></a>

<a id="canonical-4710047f53636b9aa4d456032c717f7412e673b6bcbce6569c5cbb2f3145b337"></a>

## initial_delay property — stateful_service.containers.readiness_check / 86857ceda411 / 5

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-d554f6b6e35f1607cd42a00c7510bbcfd52dde68892dc2c3255c6fa42e4b8fb9"></a>

<a id="canonical-25e6125ce32da1f70f5558bc428edeec3ce9681e61937634d2294b594b70b639"></a>

## interval property — stateful_service.containers.readiness_check / 86857ceda411 / 6

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](resources--workload--reference--group-028.md#canonical-b0a641f4835d535e4cd49e6afa10102e6c13513d954aeff642b64cf13aeebf0e): complete subsection reference.

<a id="canonical-6f10f87a5a331bee2217dd7425072d7007c50f12ec2a7fe836fc965b3abffcf7"></a>

<a id="canonical-d9cdacab020a7b43063007bb76d3ab3ca64569b6c7c485ced8694d68c9099d82"></a>

## timeout property — stateful_service.containers.readiness_check / 86857ceda411 / 7

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-ea8910b729e3f5bc882afdf820ecfe8b23bc04d6b77dfa4f0da361cfee1e0b1e"></a>

<a id="canonical-f5acbca8554640f091462e977326ec4140a1c88e29a84b5eb9829cf27a9b4047"></a>

## unhealthy_threshold property — stateful_service.containers.readiness_check / 86857ceda411 / 8

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-bb5770e9796e79a761603969b4e6e5bcc4cb7f0fdf67ee7378ad4d5e277db393"></a>

## Next pages — stateful_service.containers.readiness_check / 86857ceda411 / 9

- [stateful_service.containers.readiness_check.exec_health_check](resources--workload--reference--group-028.md#canonical-98d22817f1a2657d0a062babd6e1d210ee24d5d2a09ac8d3a08487727bd100f8)
- [stateful_service.containers.readiness_check.http_health_check](resources--workload--reference--group-028.md#canonical-d3dfe5ab9c3c8e563ac6e0a5abea26e4bdf1929a75e607cd312742a7fad60b01)
- [stateful_service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-028.md#canonical-b0a641f4835d535e4cd49e6afa10102e6c13513d954aeff642b64cf13aeebf0e)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-98d22817f1a2657d0a062babd6e1d210ee24d5d2a09ac8d3a08487727bd100f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-524a228e99b111f956fd4787b36146fe380e503206a8fdb1080b3205e467b370"></a>

## stateful_service.containers.readiness_check.exec_health_check — stateful_service.containers.readiness_check.exec_health_check / 23fb07d992ad / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-91bc45320ba6ff7008181894d544fbf7ddcbbf8937e3b3774d41073f5d7a73b9)
- stateful_service.containers.readiness_check.exec_health_check

<a id="canonical-49f18acc9ae4aa8da88971d49e6e75be6ac65d19e54808ebd0699e0d82232730"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("command")}
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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-361e193082aa7927376a3d63ceca250fac9c84c66d93272e969b272163e4f64a"></a>

## Direct properties — stateful_service.containers.readiness_check.exec_health_check / 23fb07d992ad / 3

<a id="canonical-7545424a0c9aeb6e70e0eec705336b075b3780500a491f3c44ab2992de649da5"></a>

<a id="canonical-e4a710815be9100bd042ce685383fc53e4665a47305235d348770771742f6205"></a>

## command property — stateful_service.containers.readiness_check.exec_health_check / 23fb07d992ad / 4

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-d68ed3374513b5abf686b6cb6fefba83c892aeb153732a3f78ad09cb233fedb5"></a>

## Next pages — stateful_service.containers.readiness_check.exec_health_check / 23fb07d992ad / 5

- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-91bc45320ba6ff7008181894d544fbf7ddcbbf8937e3b3774d41073f5d7a73b9)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d3dfe5ab9c3c8e563ac6e0a5abea26e4bdf1929a75e607cd312742a7fad60b01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d64af4bb0bf8162fc7c7d1cffc038f66e7ac5ac34203d54ecc6f30ea7495ba0"></a>

## stateful_service.containers.readiness_check.http_health_check — stateful_service.containers.readiness_check.http_health_check / d83414155289 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-91bc45320ba6ff7008181894d544fbf7ddcbbf8937e3b3774d41073f5d7a73b9)
- stateful_service.containers.readiness_check.http_health_check

<a id="canonical-8339c0d41ffdde07f4d78ab7ff77c9580bb5cd93322731c39e49c8df27ce72dd"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-96f0413e2a7d962e73956d8409cbe73f64a7bdf5b8fc149635048c437639eb2b"></a>

## Direct properties — stateful_service.containers.readiness_check.http_health_check / d83414155289 / 3

<a id="canonical-3dc4e12967bda802231bbde148dabbe502f78da1af99d32887c373e6439a0a20"></a>

<a id="canonical-bcfafa4e4a5dacd4ad6e5e6be73e0e54228921edd5c997f36e29f93567f21e9d"></a>

## headers property — stateful_service.containers.readiness_check.http_health_check / d83414155289 / 4

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-304184df358dcf08d60f11b06353d651380f175bead0e60f245afbbbaf193538"></a>

<a id="canonical-99763b877b0e5162b574137c1b81e52eae2eac57f23959cbc1e19cf744b84fe2"></a>

## host_header property — stateful_service.containers.readiness_check.http_health_check / d83414155289 / 5

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-96f4dadfdb092094f905d7045d0825114cb17c95f72c55b02cc3677c0e74b1d6"></a>

<a id="canonical-6c4f0901cc324c3cb38ef43e3517d205990ee5cfa6d8041895f2bfbec0024872"></a>

## path property — stateful_service.containers.readiness_check.http_health_check / d83414155289 / 6

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](resources--workload--reference--group-028.md#canonical-ed704e21383d68f0d6c8b61e8039c73211ee0fec59a117c5be23c5942a4e041c): complete subsection reference.

<a id="canonical-21c3ae0c340b6e0867e26ac2282293bdee99fff349e5a59db8913255e5ae3368"></a>

## Next pages — stateful_service.containers.readiness_check.http_health_check / d83414155289 / 7

- [stateful_service.containers.readiness_check.http_health_check.port](resources--workload--reference--group-028.md#canonical-ed704e21383d68f0d6c8b61e8039c73211ee0fec59a117c5be23c5942a4e041c)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-91bc45320ba6ff7008181894d544fbf7ddcbbf8937e3b3774d41073f5d7a73b9)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ed704e21383d68f0d6c8b61e8039c73211ee0fec59a117c5be23c5942a4e041c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ce8fa649119e12e0d58d488031edd9bada705fd8235c7d9ccff191c158b1cf3"></a>

## stateful_service.containers.readiness_check.http_health_check.port — stateful_service.containers.readiness_check.http_health_check.port / b0668394990b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-91bc45320ba6ff7008181894d544fbf7ddcbbf8937e3b3774d41073f5d7a73b9)
- [stateful_service.containers.readiness_check.http_health_check](resources--workload--reference--group-028.md#canonical-d3dfe5ab9c3c8e563ac6e0a5abea26e4bdf1929a75e607cd312742a7fad60b01)
- stateful_service.containers.readiness_check.http_health_check.port

<a id="canonical-295bde65a88550927c114e3504dd5501ebfbd41c25b3f2f0d1d6a351a9fde29b"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-e8491dc4ccf0de9f46d3523caf5587dc9128f3bb79054a29189223008eb85cd7"></a>

## Direct properties — stateful_service.containers.readiness_check.http_health_check.port / b0668394990b / 3

<a id="canonical-72611a517a16436412a272aae808d3f4b7bfdbd98b00517822e1f9e52afea884"></a>

<a id="canonical-8b32c57ae8d5653541001c0311f95f86b87f48899d2f97f296e3bf00dcf9ecfd"></a>

## name property — stateful_service.containers.readiness_check.http_health_check.port / b0668394990b / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-66f9310d2c0893ab23184b0ec49b192ff0286e57e398530df8eccd903ce8ab90"></a>

<a id="canonical-f22a121d8b3093b35cdaa4ae53c7570c5efcf35b587ec2b903db087a3c5de9b1"></a>

## num property — stateful_service.containers.readiness_check.http_health_check.port / b0668394990b / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-f1accde72df835e23a5d7c466a9bb24825805ba227544d8f9c3c13b058bb0d93"></a>

## Next pages — stateful_service.containers.readiness_check.http_health_check.port / b0668394990b / 6

- [stateful_service.containers.readiness_check.http_health_check](resources--workload--reference--group-028.md#canonical-d3dfe5ab9c3c8e563ac6e0a5abea26e4bdf1929a75e607cd312742a7fad60b01)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b0a641f4835d535e4cd49e6afa10102e6c13513d954aeff642b64cf13aeebf0e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-763e109c59e1ac774e89388fc15ec866c835b4c20350432346e0f082499c4bf3"></a>

## stateful_service.containers.readiness_check.tcp_health_check — stateful_service.containers.readiness_check.tcp_health_check / 80e6ad351ec3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-91bc45320ba6ff7008181894d544fbf7ddcbbf8937e3b3774d41073f5d7a73b9)
- stateful_service.containers.readiness_check.tcp_health_check

<a id="canonical-041562476efd18ba52ad9ab0a5c2d2ff7c9bca3f1242a96b0f0e608a2afd4faa"></a>

Type: `"object"`. single nested block, Optional.

TCPHealthCheckType describes a health check based on opening a TCP connection.

Receipt-pinned upstream constraints:

```json
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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-bedf596fad927acf79e5c8e1876ffa175ccac556decbe3767a6995b8decaa7fa"></a>

## Direct properties — stateful_service.containers.readiness_check.tcp_health_check / 80e6ad351ec3 / 3

- [port](resources--workload--reference--group-028.md#canonical-d801fa183c6cc0aff4e211191e5dbeba454624e44c84a79488c41ff85c35165e): complete subsection reference.

<a id="canonical-1434cb38445657fa748dd725d1d039fccc4f217d1cb9a31b5f838c8d354526e1"></a>

## Next pages — stateful_service.containers.readiness_check.tcp_health_check / 80e6ad351ec3 / 4

- [stateful_service.containers.readiness_check.tcp_health_check.port](resources--workload--reference--group-028.md#canonical-d801fa183c6cc0aff4e211191e5dbeba454624e44c84a79488c41ff85c35165e)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-91bc45320ba6ff7008181894d544fbf7ddcbbf8937e3b3774d41073f5d7a73b9)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d801fa183c6cc0aff4e211191e5dbeba454624e44c84a79488c41ff85c35165e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-20ae9d4a9ce265c9fb2bd89ce307c4ef318dfdfd33ec33cfb1eadf1a0b7883bc"></a>

## stateful_service.containers.readiness_check.tcp_health_check.port — stateful_service.containers.readiness_check.tcp_health_check.port / deb414cd0c88 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-e99bf3bb456823083e4129ffac6f5e01d1780828fbcb8d3560f97df96d0905f1)
- [stateful_service.containers.readiness_check](resources--workload--reference--group-028.md#canonical-91bc45320ba6ff7008181894d544fbf7ddcbbf8937e3b3774d41073f5d7a73b9)
- [stateful_service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-028.md#canonical-b0a641f4835d535e4cd49e6afa10102e6c13513d954aeff642b64cf13aeebf0e)
- stateful_service.containers.readiness_check.tcp_health_check.port

<a id="canonical-5f2371dc14a3403ddcb42bd5311b998a513fea1d81c0d5dfbd80f3db873347e5"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-848ade8e93cfd6fc65c60ad2a75612f1de74d497e87fd212012a563634437d40"></a>

## Direct properties — stateful_service.containers.readiness_check.tcp_health_check.port / deb414cd0c88 / 3

<a id="canonical-3f9b27c48b7ec0274215af11099673071b1905917f0a0b6a4fabab3c4d904f98"></a>

<a id="canonical-c8edb169526c35e215fa928cf3a82b9c173b988ce50c41dbce0fb1483cc81c73"></a>

## name property — stateful_service.containers.readiness_check.tcp_health_check.port / deb414cd0c88 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-98798e728aac38e4b2f7881af18d7a8c87704808c70c0f8d311e18226179dd22"></a>

<a id="canonical-47213bb3d39faad9f952f4cbed3de4124082402b4ef1e1dbecd00074a2b286a7"></a>

## num property — stateful_service.containers.readiness_check.tcp_health_check.port / deb414cd0c88 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-b90c7a73a50a355660f571656481844fde4020e56f502f0344dee247bfebae2f"></a>

## Next pages — stateful_service.containers.readiness_check.tcp_health_check.port / deb414cd0c88 / 6

- [stateful_service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-028.md#canonical-b0a641f4835d535e4cd49e6afa10102e6c13513d954aeff642b64cf13aeebf0e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d87da1f8ae782fee16cdfd107b92a0e65c681bf767530c6cc3c1c089b8f544a2"></a>

## stateful_service.deploy_options — stateful_service.deploy_options / 1fe66cb10fbf / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- stateful_service.deploy_options

<a id="canonical-daa63b1d2bd61c9cdee9b2b5e16d75419d7ecb6c602f4d1660f4b87099a1f8d2"></a>

Type: `"object"`. single nested block, Optional.

Deploy OPTIONS are used to configure the workload deployment OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_res",
    "default_virtual_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_ce_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("all_res",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_ce_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("default_virtual_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_ce_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_virtual_sites",
    "deploy_re_sites"),
  validators.ConflictingObjectAttributes("deploy_ce_virtual_sites",
    "deploy_re_virtual_sites"),
  validators.ConflictingObjectAttributes("deploy_re_sites",
    "deploy_re_virtual_sites")}
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
  "x-ves-oneof-field-deploy_choice": "[\"all_res\",\"default_virtual_sites\",\"deploy_ce_sites\",\"deploy_ce_virtual_sites\",\"deploy_re_sites\",\"deploy_re_virtual_sites\"]"
}
```

Terraform syntax:

```terraform
deploy_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-75ade84c729c8392ec75247ca9cb54288c26aa7f64fe1229a65e411fc9b7da25"></a>

## Direct properties — stateful_service.deploy_options / 1fe66cb10fbf / 3

- [all_res](resources--workload--reference--group-028.md#canonical-3f21a307ae498f1001b965635a64f1e75c91728e9d9fcca5749e4aa71b04343f): complete subsection reference.

- [default_virtual_sites](resources--workload--reference--group-028.md#canonical-8e6f8a8261ed3675e0caec5a6880dd278f131351b0e7efec42c53de8dbbd5873): complete subsection reference.

- [deploy_ce_sites](resources--workload--reference--group-028.md#canonical-f317a5e39414499eed96bd542f8fd93136464247af731ebb3d5a0d3da5dc7be0): complete subsection reference.

- [deploy_ce_virtual_sites](resources--workload--reference--group-028.md#canonical-cbca6554517323b778587fbf1bc335362b53efdbd0f9da12c38c6a0e0b9d8edc): complete subsection reference.

- [deploy_re_sites](resources--workload--reference--group-028.md#canonical-ca6381aecf772eea7ba4062737bd53347ec90fd52deabe3cc99165a3e00f2daa): complete subsection reference.

- [deploy_re_virtual_sites](resources--workload--reference--group-028.md#canonical-1f037266f52514c7a0edc54aa6e0e2a9a1973a3ffe824acd03c2d932408852cf): complete subsection reference.

<a id="canonical-edfef2e1bd78af8900bd2dee4028b093551d143f3e3a5e027f57269072ed5367"></a>

## Next pages — stateful_service.deploy_options / 1fe66cb10fbf / 4

- [stateful_service.deploy_options.all_res](resources--workload--reference--group-028.md#canonical-3f21a307ae498f1001b965635a64f1e75c91728e9d9fcca5749e4aa71b04343f)
- [stateful_service.deploy_options.default_virtual_sites](resources--workload--reference--group-028.md#canonical-8e6f8a8261ed3675e0caec5a6880dd278f131351b0e7efec42c53de8dbbd5873)
- [stateful_service.deploy_options.deploy_ce_sites](resources--workload--reference--group-028.md#canonical-f317a5e39414499eed96bd542f8fd93136464247af731ebb3d5a0d3da5dc7be0)
- [stateful_service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-028.md#canonical-cbca6554517323b778587fbf1bc335362b53efdbd0f9da12c38c6a0e0b9d8edc)
- [stateful_service.deploy_options.deploy_re_sites](resources--workload--reference--group-028.md#canonical-ca6381aecf772eea7ba4062737bd53347ec90fd52deabe3cc99165a3e00f2daa)
- [stateful_service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-028.md#canonical-1f037266f52514c7a0edc54aa6e0e2a9a1973a3ffe824acd03c2d932408852cf)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3f21a307ae498f1001b965635a64f1e75c91728e9d9fcca5749e4aa71b04343f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5568ceb4a580856ddd1d36973a12ce165a42c5b7acab751ff7d29d27cbfaac6c"></a>

## stateful_service.deploy_options.all_res — stateful_service.deploy_options.all_res / da429f7cf7ae / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- stateful_service.deploy_options.all_res

<a id="canonical-6d3159275ec2f3b56e93b4b77088f7444767478528c8745992239b55e817fa2a"></a>

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
all_res = {}
```

<a id="canonical-416544760f72be75f1324295540108ae3b2dfaf2dd340a324307bf5098129ca5"></a>

## Direct properties — stateful_service.deploy_options.all_res / da429f7cf7ae / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74ec50dc2435c988aa6ac6c3419bff6d50901476eed3f14fe499a89917b27804"></a>

## Next pages — stateful_service.deploy_options.all_res / da429f7cf7ae / 4

- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8e6f8a8261ed3675e0caec5a6880dd278f131351b0e7efec42c53de8dbbd5873"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0be1c8aa03556c61ed645319a9b83d7c0d04ba6cdab4539316df897da5c5c1b1"></a>

## stateful_service.deploy_options.default_virtual_sites — stateful_service.deploy_options.default_virtual_sites / 06f0833176f4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- stateful_service.deploy_options.default_virtual_sites

<a id="canonical-3da82370cc0762b2fa308677fbd51f392534b52377ece5402da7a48ba7af1c83"></a>

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
default_virtual_sites = {}
```

<a id="canonical-daa5d34ae8f14b5f64ee40a36d5b78739cbe96080ba9536592a0ac199cf4421f"></a>

## Direct properties — stateful_service.deploy_options.default_virtual_sites / 06f0833176f4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c83fd1767a49452c64ab7b6e1d7906190c7034f3f283791327c29299f3919e46"></a>

## Next pages — stateful_service.deploy_options.default_virtual_sites / 06f0833176f4 / 4

- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f317a5e39414499eed96bd542f8fd93136464247af731ebb3d5a0d3da5dc7be0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0a7adb1107adbdf30964cb9f95f0f37d689272a56b824c92255e93ee0708d56"></a>

## stateful_service.deploy_options.deploy_ce_sites — stateful_service.deploy_options.deploy_ce_sites / 1cc019ea7016 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- stateful_service.deploy_options.deploy_ce_sites

<a id="canonical-07d84c85eb3dcb9375c06d70cd6065e3c8d18e95c347efa10644b14dd03330bd"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Customer sites.

Upstream description:

This defines a way to deploy a workload on specific Customer sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
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
deploy_ce_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-0229c8f1e52f86f8df19a70fbb1ecda647ded9170914e5a580e8a36f174ab2ba"></a>

## Direct properties — stateful_service.deploy_options.deploy_ce_sites / 1cc019ea7016 / 3

- [site](resources--workload--reference--group-028.md#canonical-89e06417745f819a86aac35f2bd25d44119dc02a7aa955dfc5a09e2866e3acb8): complete subsection reference.

<a id="canonical-90cd835178694f8d33697a133d62e4665ce8575b2d8a9deaf4e8adda90448422"></a>

## Next pages — stateful_service.deploy_options.deploy_ce_sites / 1cc019ea7016 / 4

- [stateful_service.deploy_options.deploy_ce_sites.site](resources--workload--reference--group-028.md#canonical-89e06417745f819a86aac35f2bd25d44119dc02a7aa955dfc5a09e2866e3acb8)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-89e06417745f819a86aac35f2bd25d44119dc02a7aa955dfc5a09e2866e3acb8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b47df53da694c7d14d84cd1bdecd4f1e2a1f13255acef2376067b27ed6139be"></a>

## stateful_service.deploy_options.deploy_ce_sites.site — stateful_service.deploy_options.deploy_ce_sites.site / 5268714e5f8b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- [stateful_service.deploy_options.deploy_ce_sites](resources--workload--reference--group-028.md#canonical-f317a5e39414499eed96bd542f8fd93136464247af731ebb3d5a0d3da5dc7be0)
- stateful_service.deploy_options.deploy_ce_sites.site

<a id="canonical-b6944f01d29a9da8e378ce55fe86c6fd9673380c8310d82c52479e8636d46da2"></a>

Type: `"object"`. list nested block, Optional.

Which customer sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-279fb6caee1666b016ed53cf22ccbf61449c56f7d71f9045a32d0c7947201b5c"></a>

## Direct properties — stateful_service.deploy_options.deploy_ce_sites.site / 5268714e5f8b / 3

<a id="canonical-423392fc2b345a55776be0e397dbe66bf606d1f33792cbe95d208d38c28bc6eb"></a>

<a id="canonical-279b1ed01e19d3ae899bb9c7a80fe1c7e3df5074a1b801637f12e56cafb8bfbf"></a>

## name property — stateful_service.deploy_options.deploy_ce_sites.site / 5268714e5f8b / 4

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

<a id="canonical-fbe5ef12817a6560a21fcf7be5b2d2b19e21486e1ee0b0626652417e8283206c"></a>

<a id="canonical-5c18bd147b6e97ce6081e4a155e83fa7c5e94cf97e5a254d2c56cbeeee373252"></a>

## namespace property — stateful_service.deploy_options.deploy_ce_sites.site / 5268714e5f8b / 5

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

<a id="canonical-f6bc62e181f8d2b202654263e760f19494adcd55b19c87fddfe19f3cd6d0639d"></a>

<a id="canonical-e4e782b36f3b7ad2a8259522703c222f6cd32102469423ef1e7dc9fa3c736f64"></a>

## tenant property — stateful_service.deploy_options.deploy_ce_sites.site / 5268714e5f8b / 6

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

<a id="canonical-95f892fc6a0b869ef09fe402b9557be9c3e5ba020627f7ee7084218fc7465312"></a>

## Next pages — stateful_service.deploy_options.deploy_ce_sites.site / 5268714e5f8b / 7

- [stateful_service.deploy_options.deploy_ce_sites](resources--workload--reference--group-028.md#canonical-f317a5e39414499eed96bd542f8fd93136464247af731ebb3d5a0d3da5dc7be0)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cbca6554517323b778587fbf1bc335362b53efdbd0f9da12c38c6a0e0b9d8edc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-287a18a7c01f1aba09f192a9d3780d4128c350353db5234f0eb515a08fb31720"></a>

## stateful_service.deploy_options.deploy_ce_virtual_sites — stateful_service.deploy_options.deploy_ce_virtual_sites / 4cb7c43363b2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- stateful_service.deploy_options.deploy_ce_virtual_sites

<a id="canonical-85f67d990bdc8c39eb9c540d9a3dbd2082de7cf3da7e49be0a7743326946d4a1"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Customer virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Customer virtual sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("virtual_site")}
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
deploy_ce_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-6183c8e2341b420260bbcecca7586a215d2a775b8e5a3584e6f76d917f8188c2"></a>

## Direct properties — stateful_service.deploy_options.deploy_ce_virtual_sites / 4cb7c43363b2 / 3

- [virtual_site](resources--workload--reference--group-028.md#canonical-719ded1619757576086f4d53f044c0a447c1d681141bb58821872fdbfbcad8cb): complete subsection reference.

<a id="canonical-a782febd9411bae80d8a7d1dbd138637510f9679917a5c783520be4136104bb3"></a>

## Next pages — stateful_service.deploy_options.deploy_ce_virtual_sites / 4cb7c43363b2 / 4

- [stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site](resources--workload--reference--group-028.md#canonical-719ded1619757576086f4d53f044c0a447c1d681141bb58821872fdbfbcad8cb)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-719ded1619757576086f4d53f044c0a447c1d681141bb58821872fdbfbcad8cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6960db4aee07977d3e574a15529211aedf49cee995cb6bd3de9b24389e60af4a"></a>

## stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site — stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site / 17cc816d8b37 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- [stateful_service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-028.md#canonical-cbca6554517323b778587fbf1bc335362b53efdbd0f9da12c38c6a0e0b9d8edc)
- stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site

<a id="canonical-0a1398f2b7aa1d9024dd3976bff6de9818ac9d693f86e3afe5e912c6dc2d6fee"></a>

Type: `"object"`. list nested block, Optional.

Which customer virtual sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-63b3eb2b36210b090fd5ad0d28910a1073fb1b2dd27a3862b8f7f15ce29ccc7a"></a>

## Direct properties — stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site / 17cc816d8b37 / 3

<a id="canonical-e6aef0a347ad320c1d61ee89908354b0a2fbcc8e8d06cf004066283451fa0703"></a>

<a id="canonical-1835ea2234b7a5585f09b1531cecf4cd261d919ea24c7d7b5826ac7d73677f96"></a>

## name property — stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site / 17cc816d8b37 / 4

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

<a id="canonical-ca9417114b8ccefc0e4021703961f7bbfb36c075f968f8c2082515080a16518b"></a>

<a id="canonical-51e692a088d94619cc6260088bd62fc721aa90748e10ff7596231f25fee988bd"></a>

## namespace property — stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site / 17cc816d8b37 / 5

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

<a id="canonical-ccf0ebb1ff3ce41d56c99bc0f244b35e1b69b007f20c107cb1c202cb19a70059"></a>

<a id="canonical-b1c352b2b433615bf5ad8f6543363c5cbba65eb2371a3cb54f680d5980070e56"></a>

## tenant property — stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site / 17cc816d8b37 / 6

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

<a id="canonical-f3ad638cb5275cfde96d576acdd4778a1792de796dacae4558907b2a9c42f192"></a>

## Next pages — stateful_service.deploy_options.deploy_ce_virtual_sites.virtual_site / 17cc816d8b37 / 7

- [stateful_service.deploy_options.deploy_ce_virtual_sites](resources--workload--reference--group-028.md#canonical-cbca6554517323b778587fbf1bc335362b53efdbd0f9da12c38c6a0e0b9d8edc)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ca6381aecf772eea7ba4062737bd53347ec90fd52deabe3cc99165a3e00f2daa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9bec4ca22e8d87b6701b76e4e83ae57e73a356cca46c69d9f5fc7a3330e29b9"></a>

## stateful_service.deploy_options.deploy_re_sites — stateful_service.deploy_options.deploy_re_sites / 77604a88bac9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- stateful_service.deploy_options.deploy_re_sites

<a id="canonical-59ecb5b177cfd7ea521a337e43b4073cfde57b9ea7493fc3e1fbb7f1dc2d4738"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("site")}
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
deploy_re_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-82e559ebc48fc317a819c1c1f86d49c5f1dc8f8c24ac9f653c83ff406850e2fa"></a>

## Direct properties — stateful_service.deploy_options.deploy_re_sites / 77604a88bac9 / 3

- [site](resources--workload--reference--group-028.md#canonical-a1241f9a0368b278640e17b770e46cd35fc7e1e2eec317e90114d706827b93b9): complete subsection reference.

<a id="canonical-f231247f7419fb0d332a5263e6ce655554194a6ef2624d452c06d197755c7a83"></a>

## Next pages — stateful_service.deploy_options.deploy_re_sites / 77604a88bac9 / 4

- [stateful_service.deploy_options.deploy_re_sites.site](resources--workload--reference--group-028.md#canonical-a1241f9a0368b278640e17b770e46cd35fc7e1e2eec317e90114d706827b93b9)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a1241f9a0368b278640e17b770e46cd35fc7e1e2eec317e90114d706827b93b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a43632d2a2c9d68c705b227821181e468d05aeb197926f6ea329c4625331d094"></a>

## stateful_service.deploy_options.deploy_re_sites.site — stateful_service.deploy_options.deploy_re_sites.site / dcb3fabfd1a5 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- [stateful_service.deploy_options.deploy_re_sites](resources--workload--reference--group-028.md#canonical-ca6381aecf772eea7ba4062737bd53347ec90fd52deabe3cc99165a3e00f2daa)
- stateful_service.deploy_options.deploy_re_sites.site

<a id="canonical-7cb1a7ec93a052c1fcd753457f4c72373edc9a93f8370cf4c012b044b3e96e90"></a>

Type: `"object"`. list nested block, Optional.

Which regional edge sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-974cde67411e983f023b926b49debf346d10916e4ede3c2b62d932f4628c62dd"></a>

## Direct properties — stateful_service.deploy_options.deploy_re_sites.site / dcb3fabfd1a5 / 3

<a id="canonical-df53ceb7c3e70edbc8cf110ef7a32560657eded9ff40cfcecf3c66958bfc33b1"></a>

<a id="canonical-9fdb072a56d9ab4128c8e2b01b5ec8a833dd05ebd2dc77bc514f927bcb2a9e68"></a>

## name property — stateful_service.deploy_options.deploy_re_sites.site / dcb3fabfd1a5 / 4

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

<a id="canonical-5d2a7424f9b82d796274079472e5a631cef410110c8243fb3e0725b2369c09c8"></a>

<a id="canonical-339e3ddc5c6f1bb536e8f454efab2b7adeaf257067737bfb2da672bae2431843"></a>

## namespace property — stateful_service.deploy_options.deploy_re_sites.site / dcb3fabfd1a5 / 5

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

<a id="canonical-7c9ac5dda9d786ecb6df1a1581ca165c5abe263cc38ae07c15c7d237327b779a"></a>

<a id="canonical-5d3926c754ae997c2c69e8858fab950af6b98fb725f3c3fd7470f65b305c33d8"></a>

## tenant property — stateful_service.deploy_options.deploy_re_sites.site / dcb3fabfd1a5 / 6

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

<a id="canonical-74252cfc587424f10531a96e55bd46ed5a62ae67de3c746398fff8c65840ca8d"></a>

## Next pages — stateful_service.deploy_options.deploy_re_sites.site / dcb3fabfd1a5 / 7

- [stateful_service.deploy_options.deploy_re_sites](resources--workload--reference--group-028.md#canonical-ca6381aecf772eea7ba4062737bd53347ec90fd52deabe3cc99165a3e00f2daa)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1f037266f52514c7a0edc54aa6e0e2a9a1973a3ffe824acd03c2d932408852cf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1b058bb3a1226a2e0ac9bcdeedbf8d21636c1f261bf3ec9370a506dd1b2e890"></a>

## stateful_service.deploy_options.deploy_re_virtual_sites — stateful_service.deploy_options.deploy_re_virtual_sites / df258b57ae53 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- stateful_service.deploy_options.deploy_re_virtual_sites

<a id="canonical-be08f8b156148a1233d426fafd6b98f2ecd002fe59729d5178397196ef5c8d53"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to deploy a workload on specific Regional Edge virtual sites.

Upstream description:

This defines a way to deploy a workload on specific Regional Edge virtual sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("virtual_site")}
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
deploy_re_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-ea34b8c422bcbaf55fbef19f9e0df83b68cfb9c6a7b88408edd0a7b6f8f042c5"></a>

## Direct properties — stateful_service.deploy_options.deploy_re_virtual_sites / df258b57ae53 / 3

- [virtual_site](resources--workload--reference--group-028.md#canonical-ec0b60d14d58852805f18b660b67b76368679cb4cea315fdb028ce270b76287a): complete subsection reference.

<a id="canonical-c8d7dcc35b0a0616262391797c70de3c4be783b107d45b05b8eac0d1ec2e0258"></a>

## Next pages — stateful_service.deploy_options.deploy_re_virtual_sites / df258b57ae53 / 4

- [stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site](resources--workload--reference--group-028.md#canonical-ec0b60d14d58852805f18b660b67b76368679cb4cea315fdb028ce270b76287a)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ec0b60d14d58852805f18b660b67b76368679cb4cea315fdb028ce270b76287a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40883ca4060f5467db6d483d938f5c7c4c9653049595600720c8c9c9eb6bda04"></a>

## stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site — stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site / fd24c695c624 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.deploy_options](resources--workload--reference--group-028.md#canonical-2513a93c02711ec77866b205feb6413e7dd2117ba4c06380bf5e0b29878808c3)
- [stateful_service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-028.md#canonical-1f037266f52514c7a0edc54aa6e0e2a9a1973a3ffe824acd03c2d932408852cf)
- stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site

<a id="canonical-b67f0fb8bad8df520cb4096ff35fe327ecf77e93ae739f67de7f24bef47d8028"></a>

Type: `"object"`. list nested block, Optional.

Which regional edge virtual sites should this workload be deployed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-9324bde900ae251e3d3b0fca1651312ddb1abad16f5a1878d1886686ad521b4d"></a>

## Direct properties — stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site / fd24c695c624 / 3

<a id="canonical-8f7bc4f33a83c850a21f1cb361882a8949a320cbafdb1acf12fd8662a8525b48"></a>

<a id="canonical-948bb1cd97101ac91cb4e06f0c8be02611ff700201fff8856833fcb496125a9e"></a>

## name property — stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site / fd24c695c624 / 4

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

<a id="canonical-1dd876d7309d862780c87e34e94eeb420b8ae46cff0cf21209a2eb61d298268d"></a>

<a id="canonical-dd9a5911a681cdbd523e1bf66196843c6928f3f22d7757bb8bac77e15e065c14"></a>

## namespace property — stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site / fd24c695c624 / 5

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

<a id="canonical-c303beac995dfd50225d97c3ca2576abf76a887f2cf2f24abd8ad1ea31be3c4b"></a>

<a id="canonical-ea1e3d529de62850243d571b5b0c13b8bffbbbc774c0b56b96d8f51895e9ed80"></a>

## tenant property — stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site / fd24c695c624 / 6

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

<a id="canonical-a3a61d2c15dd34c86b54716df2f7bd103e74120849afdd257a81371caf5ccbae"></a>

## Next pages — stateful_service.deploy_options.deploy_re_virtual_sites.virtual_site / fd24c695c624 / 7

- [stateful_service.deploy_options.deploy_re_virtual_sites](resources--workload--reference--group-028.md#canonical-1f037266f52514c7a0edc54aa6e0e2a9a1973a3ffe824acd03c2d932408852cf)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a203e04acd7f749603be1522b30eccf5c29f4b281dddb74e04ca8905b0310f7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bebd6a0c94463a2197c55511f9fa02764e88cf57fd26b6169e324cbdeab1a06"></a>

## stateful_service.persistent_volumes — stateful_service.persistent_volumes / 736c7d782874 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- stateful_service.persistent_volumes

<a id="canonical-b07f162e7beae570e2149f641962dbc2a3fd27068ff77d10d6afe359a8910def"></a>

Type: `"object"`. list nested block, Optional.

Persistent storage configuration for the service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
persistent_volumes {
  # Configure direct properties listed below.
}
```

<a id="canonical-f08e0dccebc33a48f3415c51f0217433ad0564ddd9a210f961b8c64198f14c35"></a>

## Direct properties — stateful_service.persistent_volumes / 736c7d782874 / 3

<a id="canonical-7f2e05f92ad9a16635992713678f545d278fc5cbeb4fe4abad477be6b8067ee6"></a>

<a id="canonical-e1074c672850c3375c1b93ffa746aef12d797f651544b825d2a468a52cb9adb2"></a>

## name property — stateful_service.persistent_volumes / 736c7d782874 / 4

Type: `"string"`. Optional.

Name. Name of the volume.

Upstream description:

Name of the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

- [persistent_volume](resources--workload--reference--group-028.md#canonical-96d80b9ab6aeae048f8373415d53f468f984bb0a9d04c5115486f4970769bade): complete subsection reference.

<a id="canonical-0fabe7bc582ba7e067fb293070bec25823182249f09f5ee9ee2faa0526ee253e"></a>

## Next pages — stateful_service.persistent_volumes / 736c7d782874 / 5

- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-028.md#canonical-96d80b9ab6aeae048f8373415d53f468f984bb0a9d04c5115486f4970769bade)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-96d80b9ab6aeae048f8373415d53f468f984bb0a9d04c5115486f4970769bade"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-582a513a40e36cdfd89e869016583fdd8d4ce03acabc56699224716b3e3c3789"></a>

## stateful_service.persistent_volumes.persistent_volume — stateful_service.persistent_volumes.persistent_volume / 7d5049a5940d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.persistent_volumes](resources--workload--reference--group-028.md#canonical-a203e04acd7f749603be1522b30eccf5c29f4b281dddb74e04ca8905b0310f7e)
- stateful_service.persistent_volumes.persistent_volume

<a id="canonical-26b4c930bfeb283271bd7fc115be59d630ffad23477f32a2f1281e9344780d27"></a>

Type: `"object"`. single nested block, Optional.

Volume containing the Persistent Storage for the workload.

Receipt-pinned upstream constraints:

```json
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
persistent_volume {
  # Configure direct properties listed below.
}
```

<a id="canonical-8001e9197d6e911e37a67b8b28d228f49c4ea0d2e4aaaa3793bf07e6ab1ce5c0"></a>

## Direct properties — stateful_service.persistent_volumes.persistent_volume / 7d5049a5940d / 3

- [mount](resources--workload--reference--group-028.md#canonical-f15e2de5868f5b79e1a06d0550b234f00f0644cdaf222718d7b6a9e3dd0a29ea): complete subsection reference.

- [storage](resources--workload--reference--group-028.md#canonical-b35950953195ce91738cef5e736178f473d4cceeb032ff08e17ce59baaec8f1c): complete subsection reference.

<a id="canonical-d37c5b208d871a07baabc79134120671c16c8ea1d97003a9e3a43a5e0fc8e0a1"></a>

## Next pages — stateful_service.persistent_volumes.persistent_volume / 7d5049a5940d / 4

- [stateful_service.persistent_volumes.persistent_volume.mount](resources--workload--reference--group-028.md#canonical-f15e2de5868f5b79e1a06d0550b234f00f0644cdaf222718d7b6a9e3dd0a29ea)
- [stateful_service.persistent_volumes.persistent_volume.storage](resources--workload--reference--group-028.md#canonical-b35950953195ce91738cef5e736178f473d4cceeb032ff08e17ce59baaec8f1c)
- [stateful_service.persistent_volumes](resources--workload--reference--group-028.md#canonical-a203e04acd7f749603be1522b30eccf5c29f4b281dddb74e04ca8905b0310f7e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f15e2de5868f5b79e1a06d0550b234f00f0644cdaf222718d7b6a9e3dd0a29ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c6b6b01341ebebeccf16316cf97b3fe61b110fa02279c786ee489f1ca492e75"></a>

## stateful_service.persistent_volumes.persistent_volume.mount — stateful_service.persistent_volumes.persistent_volume.mount / 4de7b6bc91dd / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.persistent_volumes](resources--workload--reference--group-028.md#canonical-a203e04acd7f749603be1522b30eccf5c29f4b281dddb74e04ca8905b0310f7e)
- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-028.md#canonical-96d80b9ab6aeae048f8373415d53f468f984bb0a9d04c5115486f4970769bade)
- stateful_service.persistent_volumes.persistent_volume.mount

<a id="canonical-57e771cc2c80c963a648cd14273bf20910befdbe99aac11a6e283961300a2c6a"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-12a410d7bb0a900a581e43a776282ff60deff671a7174353f7f9d11530f4f898"></a>

## Direct properties — stateful_service.persistent_volumes.persistent_volume.mount / 4de7b6bc91dd / 3

<a id="canonical-271ecd66e7a5aa5439e4bc071f4c122446995026f703d76775c6a159d21f7028"></a>

<a id="canonical-8cff7e86aea65a73fae6c0499009f0bb98611c8647abdcc2cd999e7574146fa4"></a>

## mode property — stateful_service.persistent_volumes.persistent_volume.mount / 4de7b6bc91dd / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-64ae3db9d68a4e5dcee8518817b845d9621b7ef79715c1ccfd88179bfc30cae0"></a>

<a id="canonical-dc341077f7086edd179c6e91abb3485707f63da6b6bfe2ca94e83dce45ed5374"></a>

## mount_path property — stateful_service.persistent_volumes.persistent_volume.mount / 4de7b6bc91dd / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    },
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-77dd8f9b2a64e8e748752b44f37f6c93eadf256843d98178178dc02d3b0be7a8"></a>

<a id="canonical-3bf838ae193ee872270d642de5a983eea55e1104d50b859cbb8e3d6cd6b31396"></a>

## sub_path property — stateful_service.persistent_volumes.persistent_volume.mount / 4de7b6bc91dd / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-72cd6fe0a4a60d476a69a1367222e2915710e35b54804bf84bafecaef61905d8"></a>

## Next pages — stateful_service.persistent_volumes.persistent_volume.mount / 4de7b6bc91dd / 7

- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-028.md#canonical-96d80b9ab6aeae048f8373415d53f468f984bb0a9d04c5115486f4970769bade)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b35950953195ce91738cef5e736178f473d4cceeb032ff08e17ce59baaec8f1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3def00dae87779d0968cb7cb8385286ca6ff4f7d7c8508a4af4c92a20a2f7f35"></a>

## stateful_service.persistent_volumes.persistent_volume.storage — stateful_service.persistent_volumes.persistent_volume.storage / 7b8f30309a3b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.persistent_volumes](resources--workload--reference--group-028.md#canonical-a203e04acd7f749603be1522b30eccf5c29f4b281dddb74e04ca8905b0310f7e)
- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-028.md#canonical-96d80b9ab6aeae048f8373415d53f468f984bb0a9d04c5115486f4970769bade)
- stateful_service.persistent_volumes.persistent_volume.storage

<a id="canonical-870b00d85d280e92b8590cbbf1e05f59400e2dfee0987b58d33dc22cea9e9999"></a>

Type: `"object"`. single nested block, Optional.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Upstream description:

Persistent storage configuration is used to configure Persistent Volume Claim (PVC)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_size"),
  validators.ConflictingObjectAttributes("class_name",
    "default")}
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
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-169081e36f28fbc2f89a8f53bb2efb49d209c920e7fb03bb83d1655e980e04c4"></a>

## Direct properties — stateful_service.persistent_volumes.persistent_volume.storage / 7b8f30309a3b / 3

<a id="canonical-96878230224bd8493839924c91631e2cfe2e5add6b4ffa10d22b0eb20ca7c0d7"></a>

<a id="canonical-85517a6b20270b46a1e28fadad77eb38bee68d52017eb0e04e9d2d206ff47421"></a>

## access_mode property — stateful_service.persistent_volumes.persistent_volume.storage / 7b8f30309a3b / 4

Type: `"string"`. Optional.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Upstream description:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ACCESS_MODE_READ_WRITE_ONCE",
  "enum": [
    "ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0c524680f2e047fc5ca185709406ce2acf65cffab18e9c95e677eb6d901e1cc0"></a>

<a id="canonical-be4d73293f945368cf67b581b3691c6c4918a85155a65535ee9171c098c80455"></a>

## class_name property — stateful_service.persistent_volumes.persistent_volume.storage / 7b8f30309a3b / 5

Type: `"string"`. Optional.

Exclusive with \[default\] Use the specified class name.

Upstream description:

Exclusive with \[default\] Use the specified class name.

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

- [default](resources--workload--reference--group-028.md#canonical-c85802f7210630200a1cd899e202ef541e4bdb9a5cbc38bedb0614832c32892a): complete subsection reference.

<a id="canonical-1ac36ce20fba743f4e6c19fbce49e13a4bffb8298238eb1b4926025b9324b0d2"></a>

<a id="canonical-ce06ffd0273f701f12dd861fb8a6e09a705e107e037f0b8d8e1b5eaf408f5cd4"></a>

## storage_size property — stateful_service.persistent_volumes.persistent_volume.storage / 7b8f30309a3b / 6

Type: `"number"`. Optional.

Size (in GiB). Size in GiB of the persistent storage.

Upstream description:

Size in GiB of the persistent storage.

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
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-a618cc43cb67ea037abdd88fcfea8abee88768e89d53b01c3c4cfb6a2e70fba2"></a>

## Next pages — stateful_service.persistent_volumes.persistent_volume.storage / 7b8f30309a3b / 7

- [stateful_service.persistent_volumes.persistent_volume.storage.default](resources--workload--reference--group-028.md#canonical-c85802f7210630200a1cd899e202ef541e4bdb9a5cbc38bedb0614832c32892a)
- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-028.md#canonical-96d80b9ab6aeae048f8373415d53f468f984bb0a9d04c5115486f4970769bade)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c85802f7210630200a1cd899e202ef541e4bdb9a5cbc38bedb0614832c32892a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2870a71347ef6d5e195ea2eeb8daceb61daaceb8dd96e788f3e97d6d86b5c865"></a>

## stateful_service.persistent_volumes.persistent_volume.storage.default — stateful_service.persistent_volumes.persistent_volume.storage.default / 000d249c8084 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.persistent_volumes](resources--workload--reference--group-028.md#canonical-a203e04acd7f749603be1522b30eccf5c29f4b281dddb74e04ca8905b0310f7e)
- [stateful_service.persistent_volumes.persistent_volume](resources--workload--reference--group-028.md#canonical-96d80b9ab6aeae048f8373415d53f468f984bb0a9d04c5115486f4970769bade)
- [stateful_service.persistent_volumes.persistent_volume.storage](resources--workload--reference--group-028.md#canonical-b35950953195ce91738cef5e736178f473d4cceeb032ff08e17ce59baaec8f1c)
- stateful_service.persistent_volumes.persistent_volume.storage.default

<a id="canonical-6bae153e8c98d3360096dc856650cdc052995739b008c568e8d5f04d2313c8ce"></a>

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
default = {}
```

<a id="canonical-c89eaa2fc56d16b6321379ccb773709bc68bf87d3bd9611df3d1b2bf7cdbcd96"></a>

## Direct properties — stateful_service.persistent_volumes.persistent_volume.storage.default / 000d249c8084 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2e2919b938c372d75fcf2dcf7b6d4118cf2c39ed36d4e0705b81dc8d5d96cca"></a>

## Next pages — stateful_service.persistent_volumes.persistent_volume.storage.default / 000d249c8084 / 4

- [stateful_service.persistent_volumes.persistent_volume.storage](resources--workload--reference--group-028.md#canonical-b35950953195ce91738cef5e736178f473d4cceeb032ff08e17ce59baaec8f1c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-29a60e695fa2ddf2da5fe0b273aff907f4dfd3d5b85fe204b155210cd4f6636c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-440041dd09058ee2da8934c451bc149062bf8141ef5243c20cba1fe432a9e355"></a>

## stateful_service.scale_to_zero — stateful_service.scale_to_zero / cc7d94cf3e84 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- stateful_service.scale_to_zero

<a id="canonical-96e6e383b70e1dc85c8c15ddded85d282a3fc113f74ce8ebc6d79670ad0f604c"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for scale to zero.

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
scale_to_zero = {}
```

<a id="canonical-01df859aca6bce43de49df49c34f228378e1b64ef5559d5676a33051e4817b52"></a>

## Direct properties — stateful_service.scale_to_zero / cc7d94cf3e84 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-822ed880016696c9887e082bb23ee08afa9e6a30f7d3a2aed0373eb806485d87"></a>

## Next pages — stateful_service.scale_to_zero / cc7d94cf3e84 / 4

- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-38e7ef0ae08e8226d79f01ebfd032f90bd70887814b91e9ee1610bebdab59227"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-611471eb1d5ac6e378b4326f148815a45d03054ea181c69908e8f209ec4ab584"></a>

## stateful_service.volumes — stateful_service.volumes / 214b96a038b7 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- stateful_service.volumes

<a id="canonical-02ba93ca269249c6f2cccb0523aa79087e05d7ba304bd5df0476631519eb59c7"></a>

Type: `"object"`. list nested block, Optional.

Ephemeral Volumes. Ephemeral volumes for the service.

Upstream description:

Ephemeral volumes for the service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("empty_dir",
    "host_path")}
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
    }
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
volumes {
  # Configure direct properties listed below.
}
```

<a id="canonical-37f24f286922b0934dfde0043e9ef00b38475689012e09aef5c5cd35612115a3"></a>

## Direct properties — stateful_service.volumes / 214b96a038b7 / 3

- [empty_dir](resources--workload--reference--group-028.md#canonical-6f54be1db77c1206c706b220311f1837344b1235660eedcd5ab1390923ac3830): complete subsection reference.

- [host_path](resources--workload--reference--group-028.md#canonical-205c706e9066dbd7c4ca5324dc84fd01b0f54d123f4ae858548198dfa9881299): complete subsection reference.

<a id="canonical-e742e9effc5bbe15161eb6c288b0c4e5bdd11d2f3408ba17c8a6d16de57aef2e"></a>

<a id="canonical-ecc1cf88cb1123e4eb2d492b7043682697372cf8adac492444cfaf6e35db0d5a"></a>

## name property — stateful_service.volumes / 214b96a038b7 / 4

Type: `"string"`. Optional.

Name. Name of the volume.

Upstream description:

Name of the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

<a id="canonical-3661f3ccc965b83817a8640d96ecd9c0b51554dc82b8d7842363e305b295bfb1"></a>

## Next pages — stateful_service.volumes / 214b96a038b7 / 5

- [stateful_service.volumes.empty_dir](resources--workload--reference--group-028.md#canonical-6f54be1db77c1206c706b220311f1837344b1235660eedcd5ab1390923ac3830)
- [stateful_service.volumes.host_path](resources--workload--reference--group-028.md#canonical-205c706e9066dbd7c4ca5324dc84fd01b0f54d123f4ae858548198dfa9881299)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6f54be1db77c1206c706b220311f1837344b1235660eedcd5ab1390923ac3830"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a791a215f11ece88fb3e762802706de23d8d72f96153e28df1f0bfd22dd53cd0"></a>

## stateful_service.volumes.empty_dir — stateful_service.volumes.empty_dir / 2c8c2be8317b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.volumes](resources--workload--reference--group-028.md#canonical-38e7ef0ae08e8226d79f01ebfd032f90bd70887814b91e9ee1610bebdab59227)
- stateful_service.volumes.empty_dir

<a id="canonical-8df048265b714e9f21ed931a70ad93262de27278aaed458b82f10dc99fa59c03"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("size_limit")}
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
empty_dir {
  # Configure direct properties listed below.
}
```

<a id="canonical-c2b1ef27c9b178812d413f51f0a0ebff67a7e1dece0c6baf625101c28294641c"></a>

## Direct properties — stateful_service.volumes.empty_dir / 2c8c2be8317b / 3

- [mount](resources--workload--reference--group-028.md#canonical-55b223c2e5b7b48bb092a210dae8c6807b4c020b96c7c5bd15aaca8c9d69f559): complete subsection reference.

<a id="canonical-bd7f20df2cfdaa80ae4450010fa2d4df4368e1077171ca769f4ddd25c47626ad"></a>

<a id="canonical-36d06129b9c3e7148d2b8244a16d626005fa685ccd083834a33533f34ef37b2a"></a>

## size_limit property — stateful_service.volumes.empty_dir / 2c8c2be8317b / 4

Type: `"number"`. Optional.

Size Limit (in GiB). Configuration parameter for size limit

Upstream description:

Configuration parameter for size limit

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
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-ca7863095f0a4e73d34bf4cc186976e16a2af7eab8e981770a043b948c722e97"></a>

## Next pages — stateful_service.volumes.empty_dir / 2c8c2be8317b / 5

- [stateful_service.volumes.empty_dir.mount](resources--workload--reference--group-028.md#canonical-55b223c2e5b7b48bb092a210dae8c6807b4c020b96c7c5bd15aaca8c9d69f559)
- [stateful_service.volumes](resources--workload--reference--group-028.md#canonical-38e7ef0ae08e8226d79f01ebfd032f90bd70887814b91e9ee1610bebdab59227)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-55b223c2e5b7b48bb092a210dae8c6807b4c020b96c7c5bd15aaca8c9d69f559"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07b401ffa1e7cbeae4e354ce39b2d794325a9cb2d43b027865b75069add2942b"></a>

## stateful_service.volumes.empty_dir.mount — stateful_service.volumes.empty_dir.mount / 05f4a939937d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.volumes](resources--workload--reference--group-028.md#canonical-38e7ef0ae08e8226d79f01ebfd032f90bd70887814b91e9ee1610bebdab59227)
- [stateful_service.volumes.empty_dir](resources--workload--reference--group-028.md#canonical-6f54be1db77c1206c706b220311f1837344b1235660eedcd5ab1390923ac3830)
- stateful_service.volumes.empty_dir.mount

<a id="canonical-bd73c48fd4c7413230c147ce3e6f24787dd9e89b6f49ed86d3dd1ed04810f9af"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-319b9b2cd15bfb593946de6adeec7ec4f52631c284c1b863e43e0f0376e1c71d"></a>

## Direct properties — stateful_service.volumes.empty_dir.mount / 05f4a939937d / 3

<a id="canonical-a7762b83a1f39fc1a6f08ccabe70f8b7ba22ae961241950f446a0f3809b6df87"></a>

<a id="canonical-d096ad6ccaf0da4be61107ee7400437cdb6ab837a54e522c05565f10ee9c07dc"></a>

## mode property — stateful_service.volumes.empty_dir.mount / 05f4a939937d / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-d43f7614d782b11fefbc3645de77bcf4c775ef40eac0eacedad69be849b20cd7"></a>

<a id="canonical-5678b6d6f16971ff7089b20ee387e61426116414a307d4a734cf69e7c50c9c12"></a>

## mount_path property — stateful_service.volumes.empty_dir.mount / 05f4a939937d / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    },
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-de6d031e5fabce16ca6849ad774b42ed15ea7112cdd5eb5a5971dbfebe4893e8"></a>

<a id="canonical-9689eb3282d846fc4bd266e7189522f9b5736138d9dcafc457cbc28b7ae6a2e6"></a>

## sub_path property — stateful_service.volumes.empty_dir.mount / 05f4a939937d / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-39aab56d5db5a45875983037342edc80c94ab01ff36e5a022483e7b0c51a8967"></a>

## Next pages — stateful_service.volumes.empty_dir.mount / 05f4a939937d / 7

- [stateful_service.volumes.empty_dir](resources--workload--reference--group-028.md#canonical-6f54be1db77c1206c706b220311f1837344b1235660eedcd5ab1390923ac3830)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-205c706e9066dbd7c4ca5324dc84fd01b0f54d123f4ae858548198dfa9881299"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd881c98de2a2a86bf010d788ca1a997ae76bf5b4869ffa4f6905e0526965223"></a>

## stateful_service.volumes.host_path — stateful_service.volumes.host_path / 498d67bd8e39 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.volumes](resources--workload--reference--group-028.md#canonical-38e7ef0ae08e8226d79f01ebfd032f90bd70887814b91e9ee1610bebdab59227)
- stateful_service.volumes.host_path

<a id="canonical-dbe7b255a173a33b3fc0661b9af4c08766c4b79a6ab3e30efed8bbf264983f97"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a host mapped path into the workload.

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
host_path {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0706ab7e3c425504d73dc43d36b5ffd783f4d19c5d7af00c362206d7a073015"></a>

## Direct properties — stateful_service.volumes.host_path / 498d67bd8e39 / 3

- [mount](resources--workload--reference--group-028.md#canonical-c4dfb3de724f8ad62bde6f17d93ca1935507133c501d76e358d5cbcddcd231ce): complete subsection reference.

<a id="canonical-6df109488dc6c63d8e8b8573816f5aca917eeb97942a01167aa85ff0a3a01099"></a>

<a id="canonical-8badc4c924b69688ab32c383b21a0bf9e94e76d038f84a1ba9ed882c9f885908"></a>

## path property — stateful_service.volumes.host_path / 498d67bd8e39 / 4

Type: `"string"`. Optional.

Path. Path of the directory on the host.

Upstream description:

Path of the directory on the host.

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "[^\\\\0]+"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "[^\\\\0]+"
  }
}
```

<a id="canonical-29b1d7ffd1710f95cda35ae53799dcefad9355160abbd9fcc37aa3c88766e85b"></a>

## Next pages — stateful_service.volumes.host_path / 498d67bd8e39 / 5

- [stateful_service.volumes.host_path.mount](resources--workload--reference--group-028.md#canonical-c4dfb3de724f8ad62bde6f17d93ca1935507133c501d76e358d5cbcddcd231ce)
- [stateful_service.volumes](resources--workload--reference--group-028.md#canonical-38e7ef0ae08e8226d79f01ebfd032f90bd70887814b91e9ee1610bebdab59227)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c4dfb3de724f8ad62bde6f17d93ca1935507133c501d76e358d5cbcddcd231ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-36aa349e156a695885757a07e7d691b1de00344b0682fdf00c720abbaf11f011"></a>

## stateful_service.volumes.host_path.mount — stateful_service.volumes.host_path.mount / 68911b59a07a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.volumes](resources--workload--reference--group-028.md#canonical-38e7ef0ae08e8226d79f01ebfd032f90bd70887814b91e9ee1610bebdab59227)
- [stateful_service.volumes.host_path](resources--workload--reference--group-028.md#canonical-205c706e9066dbd7c4ca5324dc84fd01b0f54d123f4ae858548198dfa9881299)
- stateful_service.volumes.host_path.mount

<a id="canonical-4860c871dbf67cf345c52842a80c89f36cc7fe9e783f2e52a4fd36a72c41857e"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-d0a4b5d4c4a065d970d2dedf24b32da46f646bba3d35382eefcb596fe0f1fc23"></a>

## Direct properties — stateful_service.volumes.host_path.mount / 68911b59a07a / 3

<a id="canonical-c6a9644024ebe0e0304a4c7d94240532d756449322e0f943cb7cb821f15ffb0a"></a>

<a id="canonical-305bfd7d9b3396c6168ee0b1da6253c76571bce56907108d1093426b4bea82bc"></a>

## mode property — stateful_service.volumes.host_path.mount / 68911b59a07a / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1f744b0010205142b84ed0828f2a3bcc4275b1b0754bc86401386234a41c41f8"></a>

<a id="canonical-cb532038a813249fdaa93befd8cbc61be08c885e36f9476b8cd9b4ebf84c6b35"></a>

## mount_path property — stateful_service.volumes.host_path.mount / 68911b59a07a / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    },
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-e46d8f6a38fcc812a5630ce2b0b787685d933869879c22e49752f3de6b93f3f1"></a>

<a id="canonical-d3177a0f498eb89872390af9c99183d460620e185080aed66595406446f10e4c"></a>

## sub_path property — stateful_service.volumes.host_path.mount / 68911b59a07a / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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

<a id="canonical-1171637b5d51ac456bde54a5e5461e40ed644a80065f55b2b7d7ca95ecf51494"></a>

## Next pages — stateful_service.volumes.host_path.mount / 68911b59a07a / 7

- [stateful_service.volumes.host_path](resources--workload--reference--group-028.md#canonical-205c706e9066dbd7c4ca5324dc84fd01b0f54d123f4ae858548198dfa9881299)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-923a5edb0b626cfe6c45cb6f5d6d5e0ff2ccab8133ffed2c2cdcdabe8d3e7698"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cea738db9be2817688f3568c6425423cab585e061026287ddd242b2ebf9603f"></a>

## timeouts — timeouts / 9348da332b07 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- timeouts

<a id="canonical-ed902ba0352c11c9f01a16010daf5f74c4454d62d37253221cd683527735e3fa"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-f6693fc08a34c57464b6d75c5fc10a1d599fe5100e1b4271578a2613bdd8f93a"></a>

## Direct properties — timeouts / 9348da332b07 / 3

<a id="canonical-f3db28dd22f8afc794baede88adff9cf089d148a3033270c8bc539829799e6e9"></a>

<a id="canonical-5e1e07129f3bf65963c2300e7e8f1ab525e69199dc408c208817bcb2011f1ded"></a>

## create property — timeouts / 9348da332b07 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0163693782f5d81626662d481776b5b52d44f5fca5d90e12a80b09ee0c329a5a"></a>

<a id="canonical-3cb066c8d480b1d7176a4660a7f664e602136205c2fe389b807ebd19e56be857"></a>

## delete property — timeouts / 9348da332b07 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-b62c6c67583ec09fccf5a051a947cf853901907b7d51b5f75f069327f2a9db6e"></a>

<a id="canonical-a435107d0dcf1fd09291acfcd99e1f3cb0146b3df6cec4d20575cfbdb52aadb2"></a>

## read property — timeouts / 9348da332b07 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0386be23cb74ffe05ddfc8232463d857257481c006063dc6e550a27256668d7d"></a>

<a id="canonical-d9d2b2e29604479a52dce71817a13eced7f1d204ee6795449052a1535830742f"></a>

## update property — timeouts / 9348da332b07 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-a6a7e260d76933f6292830a3ef6d3f5e44709f72f6c0a583ad1c9e3d820b387e"></a>

## Next pages — timeouts / 9348da332b07 / 8

- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
