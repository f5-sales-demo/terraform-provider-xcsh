---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-44c09f5e380808be3933a25b4e6f72ceb262b2c6ef223db7ad58fa5535349502"></a>

## connection_idle_timeout property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bd64f81c6992 / 6

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

- [default_header](resources--workload--reference--group-007.md#canonical-00b6423ff22b53ff0db646ee63278d1cd666785304c80453ca4cfe114b97a9a5): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-007.md#canonical-08231de7e23d5c52fc75d469bf6d79224084ef79c48a6ff769120e4c7736a58c): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-007.md#canonical-83307a525155b049f148d0ba9d8960ee02589f69fe9659109d0de0747e05741d): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-007.md#canonical-65e0e67ac628ea20dcd9668fcba666a87bf5b4345fb875d4e9cc6ff3ddac1860): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310): complete subsection reference.

<a id="canonical-e730084844a873a392f0d628d1b634523a1f8e036bb38d3657db70d3137f2e3a"></a>

<a id="canonical-11adbfc4c3c7f781beaeb0e61e3c3c4531f2c98363423bb5a8e1d10b121675b8"></a>

## http_redirect property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bd64f81c6992 / 7

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

- [no_mtls](resources--workload--reference--group-007.md#canonical-57b6b64017d831bba7f2ef998b2cbf633526716760d1a12b657802c085016e40): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-007.md#canonical-69d02bac4fc6589f7b6a653e03c443e53efea82924a453e6d98da1a076822a65): complete subsection reference.

- [pass_through](resources--workload--reference--group-007.md#canonical-8f77b39a096e366c4827f7f2053e6f3ca2eb13308c2f7cfe9dc63cf690bd2563): complete subsection reference.

<a id="canonical-beeb25906d8cac8b67c82e942f110c108fd9ed480456a3843678c8d6d8f0b391"></a>

<a id="canonical-a91fc39581f8a5a1c8218d75b20443f977660ba9dd5dd052632962f3773e0332"></a>

## port property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bd64f81c6992 / 8

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

<a id="canonical-58b536e329588a88f4defbf6f997342c9ac2c61aef0993bd245fc12f20a45f35"></a>

<a id="canonical-4c8ca92784ce5d35ac1a124192722ba17f63f70ab5a8aab44de27a8136ee5134"></a>

## port_ranges property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bd64f81c6992 / 9

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

<a id="canonical-9bd9dc24a4d7144425291f705ac90ea01bef83a01a50a2c4af8f232b9e4e0a35"></a>

<a id="canonical-e66b45e6f4f0dc856672cd8a513e0b4b9e0ffe101207394888dfe84a8d0cf67d"></a>

## server_name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bd64f81c6992 / 10

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

- [tls_config](resources--workload--reference--group-007.md#canonical-ebed6b3277ec9a8bffe634e12a90ed80a8680cec3b888467de4b373f6902e274): complete subsection reference.

- [use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c): complete subsection reference.

<a id="canonical-aff3f43b7951af3edca9b688260febff793108182929258437eba758c7f4d24d"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bd64f81c6992 / 11

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-61816e53c7b194ba062a72aa4e2fa3f7cdb21875cffa5af4027ea24d7ce11c9c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header](resources--workload--reference--group-007.md#canonical-00b6423ff22b53ff0db646ee63278d1cd666785304c80453ca4cfe114b97a9a5)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer](resources--workload--reference--group-007.md#canonical-08231de7e23d5c52fc75d469bf6d79224084ef79c48a6ff769120e4c7736a58c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize](resources--workload--reference--group-007.md#canonical-83307a525155b049f148d0ba9d8960ee02589f69fe9659109d0de0747e05741d)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize](resources--workload--reference--group-007.md#canonical-65e0e67ac628ea20dcd9668fcba666a87bf5b4345fb875d4e9cc6ff3ddac1860)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls](resources--workload--reference--group-007.md#canonical-57b6b64017d831bba7f2ef998b2cbf633526716760d1a12b657802c085016e40)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer](resources--workload--reference--group-007.md#canonical-69d02bac4fc6589f7b6a653e03c443e53efea82924a453e6d98da1a076822a65)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through](resources--workload--reference--group-007.md#canonical-8f77b39a096e366c4827f7f2053e6f3ca2eb13308c2f7cfe9dc63cf690bd2563)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-ebed6b3277ec9a8bffe634e12a90ed80a8680cec3b888467de4b373f6902e274)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-61816e53c7b194ba062a72aa4e2fa3f7cdb21875cffa5af4027ea24d7ce11c9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a6a6fb695cbb99f3cb02773afbe9602158f8c60fd99522512c43c7921f23b638"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 44b0bb9066b8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-1513272e82dc345f57e8c1120c1fcbefa2d8ef44db9e40a4b8fcdad6f7c53440"></a>

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

<a id="canonical-884f30f23d47f139bbfab3477cee98748ba1a193e7ac23e3734d66bdbe662a56"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 44b0bb9066b8 / 3

- [default_coalescing](resources--workload--reference--group-007.md#canonical-8f426e970c5e49f6df8c1444c67d18dc5da1560a76676430c96e7d5441e4a3a6): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-007.md#canonical-adf97bf3b19cf8f638fba4ea0f77dca7c6e6f43f5f66d23432d68134655f6e33): complete subsection reference.

<a id="canonical-78bbae315fe56a7a6abaf9686c4ab71d8b3fd97c5156186506782ec87ef69536"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 44b0bb9066b8 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](resources--workload--reference--group-007.md#canonical-8f426e970c5e49f6df8c1444c67d18dc5da1560a76676430c96e7d5441e4a3a6)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](resources--workload--reference--group-007.md#canonical-adf97bf3b19cf8f638fba4ea0f77dca7c6e6f43f5f66d23432d68134655f6e33)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8f426e970c5e49f6df8c1444c67d18dc5da1560a76676430c96e7d5441e4a3a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25fe877f8d06cbef18f8373abc33322a35c6b6fed01fd27ef51c92df836c862b"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 33e4dc56562a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-61816e53c7b194ba062a72aa4e2fa3f7cdb21875cffa5af4027ea24d7ce11c9c)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-1851b10bd1f60a249da27551892326f2990508bcd828ab8b62ccd5be10d5934e"></a>

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

<a id="canonical-da94f6deda5a5d35bb19fd771042e147f91cb7161b5dc41a4fc563cfc9b7a654"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 33e4dc56562a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6b2846b8e02026aac2b0cf8f0bc2b9b312dd24ba893f42ccae04046f8135c0b4"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 33e4dc56562a / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-61816e53c7b194ba062a72aa4e2fa3f7cdb21875cffa5af4027ea24d7ce11c9c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-adf97bf3b19cf8f638fba4ea0f77dca7c6e6f43f5f66d23432d68134655f6e33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-727f238cf638bff5ae7ed23c6ad3f27ca440b188b58cc6083566301e455971f1"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bb39182278c2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-61816e53c7b194ba062a72aa4e2fa3f7cdb21875cffa5af4027ea24d7ce11c9c)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-ca022129f169ad6c2ab0f198adc056ca37f9d73d1f752b4b83d4fc3dd7f22b32"></a>

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

<a id="canonical-b8c1d20efd7e08a34a95837041311f363fc983863f5875de409be9e693e111eb"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bb39182278c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-32d66f07dd9991c5ad4683e2c835b332f4be23299585160b6b0c507c938c4582"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bb39182278c2 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-61816e53c7b194ba062a72aa4e2fa3f7cdb21875cffa5af4027ea24d7ce11c9c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-00b6423ff22b53ff0db646ee63278d1cd666785304c80453ca4cfe114b97a9a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3a435dbbdfda68ff1f9c0269490150565b322f4dda840da882d0129463d4548"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 28feb1ce393b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-467a32cfcc174ac6c7344bfe6163ed7eae22f197d47191224bb29c056861ad07"></a>

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

<a id="canonical-227f060c769aaf5385467c2fd849b2cc2c21869fe20e171e06112ef2610f802a"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 28feb1ce393b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-19fda581f036148f3fc179dd01220c105c2fc1ac93d6b26fda189ba336fca5c7"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 28feb1ce393b / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-08231de7e23d5c52fc75d469bf6d79224084ef79c48a6ff769120e4c7736a58c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a848c1ee99c0162bc1e83e33eb2e938dfcd7eb38fffd25967bed451757914d76"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 5208c4261a9a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-b4b39ee2de16d872b9670d2a7709423eb1f04717728738cc3fc9b1ab2e531b70"></a>

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

<a id="canonical-3392d351d4385bea9d2f86732b1f7e33975bffe82f3e0bc8e1587aa0e8ecb16e"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 5208c4261a9a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-806a0a07e07fdd2ad9d1c5d7a6059215a493535800f97571801eae1ae8569c7b"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 5208c4261a9a / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-83307a525155b049f148d0ba9d8960ee02589f69fe9659109d0de0747e05741d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51b37760f04575f3aca9a222cf24dd2f16c555afb526368ed35779717d7175ed"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 72bc66351c15 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-05c7ca03b21ea4d4c9c7383745b446c260e1509a718eaf123721535c5eec8095"></a>

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

<a id="canonical-7b01b6d62c2680bb630bbbbbccc790e30288d91440ccb3d8925310b1b6bec58f"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 72bc66351c15 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-714ba781e6a6a8af3f9157916615895358133b5ed289b110bf934380e0e395ca"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 72bc66351c15 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-65e0e67ac628ea20dcd9668fcba666a87bf5b4345fb875d4e9cc6ff3ddac1860"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-472959df9523956ff537d368942120825c4390a8af873dde62790ae3c968cec7"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / ba68047eca8e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-2ffe66c911cee72ef26b8552f861f740c344f1e87185dad148a5958ff4372af3"></a>

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

<a id="canonical-da1c1dd904012bc754e67b83e842e8166c55489c5305d6a93cc0337a0af745cf"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / ba68047eca8e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-97cc50c131f8370afd7c1602ea1243f20067f9d67a1c47bdb90e771602f60ce4"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / ba68047eca8e / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce2b5277a6bdc69906e1d064f0492f11b90aa39ed305f024ce4d4011f8125da6"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / cca257744136 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-9d36afe48402639d89035711ef7b8888c7da5d8f6613ba665049cfd641c49a62"></a>

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

<a id="canonical-a834caaa4ac1ce80ee4f12564f8a7b199553bd45ad6de00301eff28767b0034f"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / cca257744136 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-7eef2cae59366d6966504b548877a9d2d73388dd7d614cccc6852ab97a5339b2): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-007.md#canonical-ce96b73e3e99b22e69c082b266c81c9a1842bd3f17e931545ee612f0bd8e5ee9): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-007.md#canonical-6a4ad0c4befab9ca5cd4553ffcac0552edfe217587e15971ba2a8d02d2ee437d): complete subsection reference.

<a id="canonical-46ecbf46e85ce1c6c494ed1c550826d5b23065392537c38abe16995ed45d007a"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / cca257744136 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-7eef2cae59366d6966504b548877a9d2d73388dd7d614cccc6852ab97a5339b2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-007.md#canonical-ce96b73e3e99b22e69c082b266c81c9a1842bd3f17e931545ee612f0bd8e5ee9)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-007.md#canonical-6a4ad0c4befab9ca5cd4553ffcac0552edfe217587e15971ba2a8d02d2ee437d)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7eef2cae59366d6966504b548877a9d2d73388dd7d614cccc6852ab97a5339b2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c33b0a0418152ea968d224eacb9e0a920b207cc8a7d5e7262ec1f3244da149ff"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 3f1d8fe77b3e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-375326beab71b4aebcb0a88d89a5418fbd0c4c0c9c6f96ebddd8bf32df9f77cc"></a>

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

<a id="canonical-a31bec01a7e451d91ff7aa03d07ebc6b493071f15a13f4ed87110c46e406bb07"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 3f1d8fe77b3e / 3

- [header_transformation](resources--workload--reference--group-007.md#canonical-3a16dd8ed7d7bb0370ee7ab9a10fd568c590901e6fa5faa09effb489643690a6): complete subsection reference.

<a id="canonical-96729ddbdb116b9305a8700ea14fcafc6012465703d9e83fa63a70983e4272eb"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 3f1d8fe77b3e / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-3a16dd8ed7d7bb0370ee7ab9a10fd568c590901e6fa5faa09effb489643690a6)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3a16dd8ed7d7bb0370ee7ab9a10fd568c590901e6fa5faa09effb489643690a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60b12c2581de48f3adf8f1c7ac1ad05ee0e436f9b4b28984395c45471edf487b"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 905fbe7eb395 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-7eef2cae59366d6966504b548877a9d2d73388dd7d614cccc6852ab97a5339b2)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-a768203d71310ba866e59f52eed2c4c0ecb9a821e766a7d21262eb59abea400c"></a>

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

<a id="canonical-84554864544227a09b4b9e0820cd035b996e08514052c00a4052e3dbd2f689b9"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 905fbe7eb395 / 3

- [default_header_transformation](resources--workload--reference--group-007.md#canonical-23d8d2474220045a208128051d93ba674aa17866a81e5b0beb3cc57907fbf3b4): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-007.md#canonical-18346745c32c3379412c86be5d4644d8974333615e85233984b0f31f51476657): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-007.md#canonical-f2654de71f0e05b48ce72af47dd12b846e6ac15fd5af98842d4b0e66c78688a5): complete subsection reference.

<a id="canonical-f6d592cf9721752b32f63a415ce18cc036b623d13718b2f504918d7c8f2fc851"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 905fbe7eb395 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-007.md#canonical-23d8d2474220045a208128051d93ba674aa17866a81e5b0beb3cc57907fbf3b4)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-007.md#canonical-18346745c32c3379412c86be5d4644d8974333615e85233984b0f31f51476657)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-007.md#canonical-f2654de71f0e05b48ce72af47dd12b846e6ac15fd5af98842d4b0e66c78688a5)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-7eef2cae59366d6966504b548877a9d2d73388dd7d614cccc6852ab97a5339b2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-23d8d2474220045a208128051d93ba674aa17866a81e5b0beb3cc57907fbf3b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4e9b057e4207a3d0bebb985b69760f6321f56848de9430af8a900b089623054"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 900ba9419a92 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-7eef2cae59366d6966504b548877a9d2d73388dd7d614cccc6852ab97a5339b2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-3a16dd8ed7d7bb0370ee7ab9a10fd568c590901e6fa5faa09effb489643690a6)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-712e179e8bd570b673c07497ccfe2303ff1e662e87d5d80e970e2339808850a2"></a>

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

<a id="canonical-aeaae2b48339de728dda73c4a25f8775a3837041e3dd5bc09c01752efe5ab072"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 900ba9419a92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b94a30e4bba05c444376cf6db4adb1e3994dd0e6ffab52bb2b5d610f35d90c6b"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 900ba9419a92 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-3a16dd8ed7d7bb0370ee7ab9a10fd568c590901e6fa5faa09effb489643690a6)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-18346745c32c3379412c86be5d4644d8974333615e85233984b0f31f51476657"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aed259e7e6e6e05c04aaf1e0a2d53ce0fcd86500afa5a9687b92ef4ed23476d2"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1372a888b183 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-7eef2cae59366d6966504b548877a9d2d73388dd7d614cccc6852ab97a5339b2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-3a16dd8ed7d7bb0370ee7ab9a10fd568c590901e6fa5faa09effb489643690a6)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-5ea60868eb95eb40ef5018aa928a30cb327e7480fa9ac1b44cf4bc2a790cfd53"></a>

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

<a id="canonical-e34014617d7bcb7c8137b801e709fab4c63817c09462aa9c74df6950f608dc56"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1372a888b183 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-262496e3edca0eb2e45dbeb8d92517872043f94d591a3eb6430c774c56e289db"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1372a888b183 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-3a16dd8ed7d7bb0370ee7ab9a10fd568c590901e6fa5faa09effb489643690a6)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f2654de71f0e05b48ce72af47dd12b846e6ac15fd5af98842d4b0e66c78688a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-521f956de3617e609e261fd590244dfdb8b1d7aa02512ff22d298f47e4e9fef8"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / c256ab593893 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-7eef2cae59366d6966504b548877a9d2d73388dd7d614cccc6852ab97a5339b2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-3a16dd8ed7d7bb0370ee7ab9a10fd568c590901e6fa5faa09effb489643690a6)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-6f58adbb1d77648c4a20bc7039d3caa44b0af7eca84ad66aa8803807ab4211e5"></a>

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

<a id="canonical-ff9d5130e60b9e6d2e9c1a08886af526a5aa0d6894c57f6107bd96925b416605"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / c256ab593893 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34460fb411af6c2310405a0d0dfd1bce5970adf35d2a685197b27b1f751ebe09"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / c256ab593893 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-3a16dd8ed7d7bb0370ee7ab9a10fd568c590901e6fa5faa09effb489643690a6)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ce96b73e3e99b22e69c082b266c81c9a1842bd3f17e931545ee612f0bd8e5ee9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33163dd96e1152e078b4e52f34fde15dfcdccc3367c9381638a4a86d1b64b548"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 40986162ef71 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-abee67ee49939676ed3ba2ab154a1f0e3485c935b64c381145be33064d3a0143"></a>

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

<a id="canonical-547cb5f2822151a5cbe627937835702cbb4e6bd7b32305402c3202a264ad55e7"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 40986162ef71 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e5b074b6c5aee6493a21a71a8d59dcc5995754f581231ebbf7a1c5930d4f21b2"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 40986162ef71 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6a4ad0c4befab9ca5cd4553ffcac0552edfe217587e15971ba2a8d02d2ee437d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a61680f4bac5e2915c62a5c88a95acd77f24c6025ec49984245f468e1037ba0"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / d18b1458511c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-1f31ae1f7206fe04d6a2ffd156f82ee6de653dff655209370ae33437f47dd2a7"></a>

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

<a id="canonical-283605b68008988e59add7157c953e253b9f06ed776c333f323dd5aa8b736809"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / d18b1458511c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b439d9b84f13c20a72e5701e1e33789290df3ea1e8100bd7c38f689b0b1acd54"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / d18b1458511c / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-10c75059268cdac9c68ff3e58d43a7b0365022e85a99a55b9b7e3ee6f91bc310)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-57b6b64017d831bba7f2ef998b2cbf633526716760d1a12b657802c085016e40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ada4103b40c446378ba91b429ea067fb2b040bbdb80410b40c22b23fd67524b2"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 69db511e45cf / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-e93003e7f2b16c66dede66a762877b25581c3f4228116ec008e9425400d060b3"></a>

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

<a id="canonical-16e035697ab22448e89cd4afa6ddd04bbe89c3eec576233ca978be13f833d1f5"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 69db511e45cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd5d4e36548c31e1233689337df2f061e305fa4effb511ac9454592023d89c60"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 69db511e45cf / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-69d02bac4fc6589f7b6a653e03c443e53efea82924a453e6d98da1a076822a65"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d60bedc8d929291f04a3075b898362b21bfa5f6a65a858e5cf10732ea635e71"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 3733d970bce7 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-49c52f4c93cdcc79cedf640547c9fd8347dfbeafb0c286299825a7071bb76c9d"></a>

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

<a id="canonical-8980c508a622185a8cb50461d1dd327a8aa8d6daa5210be11bb7b227ca554396"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 3733d970bce7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1aa7df9c0912b9f1d59215a3437d320e5bf086e4b4c3c57cb71ec04209fa40b2"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 3733d970bce7 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8f77b39a096e366c4827f7f2053e6f3ca2eb13308c2f7cfe9dc63cf690bd2563"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a68f829c9216f2492ce9d3617b4c230e3bc89164e5ae9f58c124b7c5a424a65"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bb294602377c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-0473d84aab164a3265cea5e2c7b2f79658296d3df8632e3dae2cf9748ba84ad3"></a>

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

<a id="canonical-231b81724d8437ba96469502a8150df2e04aa17e8b5d0cedab1425f7ad4d75eb"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bb294602377c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f22a040893d63fba70430ed5b4613d7d91a2436f8f8e24e5fcbc0c46c2169076"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / bb294602377c / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ebed6b3277ec9a8bffe634e12a90ed80a8680cec3b888467de4b373f6902e274"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc4423e1ba08dec069100fadda98cb632e9ffebf5dc608520566c292bcfe5b38"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / ed9d61274ef1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-3a9f9f99946d1208fe4709506c5451d8d4abd4b1b80443c72858ee7eb86e746b"></a>

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

<a id="canonical-b59729868354832cb7c71608a6954b2f9b3e3676fba13c8e0303d1edbd2c8198"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / ed9d61274ef1 / 3

- [custom_security](resources--workload--reference--group-007.md#canonical-a49cd3283db5bd2d624b523b60393ffbf79119e4e0452cbbccedce884f3cfad2): complete subsection reference.

- [default_security](resources--workload--reference--group-007.md#canonical-331877988d74d6271bcbaab83b9b406f1d9163b6f537c6f53dd2eb55e9c8c52e): complete subsection reference.

- [low_security](resources--workload--reference--group-007.md#canonical-369d334c532f3c564c6d6fbd24bfc123e77056146632e18610b85c3f20810e18): complete subsection reference.

- [medium_security](resources--workload--reference--group-007.md#canonical-e72004a5b6d8ebf9dcb4bc72b806155dc1a44b70ed07305fccc7c7ce89486038): complete subsection reference.

<a id="canonical-200dc2494891166ba5ae9524e451a65d0ace3062c61f0361e3bc89fa849242e4"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / ed9d61274ef1 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security](resources--workload--reference--group-007.md#canonical-a49cd3283db5bd2d624b523b60393ffbf79119e4e0452cbbccedce884f3cfad2)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security](resources--workload--reference--group-007.md#canonical-331877988d74d6271bcbaab83b9b406f1d9163b6f537c6f53dd2eb55e9c8c52e)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security](resources--workload--reference--group-007.md#canonical-369d334c532f3c564c6d6fbd24bfc123e77056146632e18610b85c3f20810e18)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security](resources--workload--reference--group-007.md#canonical-e72004a5b6d8ebf9dcb4bc72b806155dc1a44b70ed07305fccc7c7ce89486038)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a49cd3283db5bd2d624b523b60393ffbf79119e4e0452cbbccedce884f3cfad2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d60ecf2c764076bf81f175457fceead18653b07afec3304618690d45f2470b0"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1f616dfc90a8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-ebed6b3277ec9a8bffe634e12a90ed80a8680cec3b888467de4b373f6902e274)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-48c326e3307d1e65727b2c50c5d6a0984bb0b6b7719c04e4dfd5cb50446b4821"></a>

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

<a id="canonical-dd0df00b62780fecbe79c74ba1043f63483e37da44b3d7c70eef46263b3b786e"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1f616dfc90a8 / 3

<a id="canonical-4d77af991bba87782798b46e6656325452cfe18979c5918bdb57a39b7251eb3a"></a>

<a id="canonical-430fb4cc3f1b3742f483f35e556705ef2faefb6eb48fb8028fa9c04dadf3ff39"></a>

## cipher_suites property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1f616dfc90a8 / 4

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

<a id="canonical-e34c5bfc7ad41efb1820299fe65600eb80d4109d8d2c820561e1d7d944f7c66b"></a>

<a id="canonical-3940ae18d1f49675e3558da2a243ef1efa9bbfb78b5c489554bed22d2366a44f"></a>

## max_version property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1f616dfc90a8 / 5

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

<a id="canonical-ddf203842252b4918cd6e3c4c6bcaaeab79a688c15370bfd74185e4eae01b71b"></a>

<a id="canonical-2ba0b664454a3a28b092771578fb406073cecb2fad28d8dfc823fa1414b9fde9"></a>

## min_version property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1f616dfc90a8 / 6

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

<a id="canonical-8d72797744415a2645f1ac4f09822eaaf65be6e711c3b730863ae7b32f85f66e"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1f616dfc90a8 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-ebed6b3277ec9a8bffe634e12a90ed80a8680cec3b888467de4b373f6902e274)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-331877988d74d6271bcbaab83b9b406f1d9163b6f537c6f53dd2eb55e9c8c52e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2a293a8d54999eb13df5551e5eebebaa1e59db9cf2e5f9a24c9c86c34f6d85f"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / a12da14e92b0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-ebed6b3277ec9a8bffe634e12a90ed80a8680cec3b888467de4b373f6902e274)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-c9fe0b8fb411d158603eb0f53a4ab80efb4d36341b50bcbfe1a0756ae5a156e7"></a>

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

<a id="canonical-a68da8ee291e3950df9b80801edd82053281795831c8c59e5f888770e7c08219"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / a12da14e92b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1692a659407500625a6da8d8e36b5d93a10a3203fd2111f474f3617d43e15546"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / a12da14e92b0 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-ebed6b3277ec9a8bffe634e12a90ed80a8680cec3b888467de4b373f6902e274)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-369d334c532f3c564c6d6fbd24bfc123e77056146632e18610b85c3f20810e18"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b48c526b1337d3a5cddfa30b33cba7052bb333bbc6bd2d88cb43aaed709ce223"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / f6c628923917 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-ebed6b3277ec9a8bffe634e12a90ed80a8680cec3b888467de4b373f6902e274)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-aef4e3be08f35ba91a81ad9472e4f424bf462f4fdbc671db8c4ae2479b969dee"></a>

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

<a id="canonical-d5023e902caff003bf5a661749c2e4eb3c59b0faeda61fbd48200a74a616509d"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / f6c628923917 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9491cf834ae0f437484b9e67493b6d72667b3856e7971567c063bcfff280abf"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / f6c628923917 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-ebed6b3277ec9a8bffe634e12a90ed80a8680cec3b888467de4b373f6902e274)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e72004a5b6d8ebf9dcb4bc72b806155dc1a44b70ed07305fccc7c7ce89486038"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-313cd6e43e3e1f5a1977db2aefa83775f884916f06d3ad29696aa6c6a21b4db1"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 05853b28d681 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-ebed6b3277ec9a8bffe634e12a90ed80a8680cec3b888467de4b373f6902e274)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-fcfc2b735c8152291ef9fc48ce2ef6445b750c4c45c161d20ccbb2aafaf9eaec"></a>

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

<a id="canonical-c47b78f06614d2ba45bfecbc3967a4d8cdbf1d6fe379fea391f4c3144793228f"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 05853b28d681 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c3b4bf1d540b08e05275caadf5e4d93c744ba23924a1add3c730ab6efe572493"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 05853b28d681 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-ebed6b3277ec9a8bffe634e12a90ed80a8680cec3b888467de4b373f6902e274)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27d935f58854a9e89c33281bacb3161736c1c2be6b841f357a72d2d05afb48b3"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 77824b9cb622 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-167ab853869012916244a93502c227dfab16a1490cd393313cc8dcf10efc6daf"></a>

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

<a id="canonical-59246711ee42bac53980cde57b82a21e2dec7334a2810dfaed01b401c9c3f577"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 77824b9cb622 / 3

<a id="canonical-af5f370dae3a3ea60a08c657b72e30e80a2e55456e61c691bbdbce481bba759f"></a>

<a id="canonical-e0c50634339097fa11ce52391316ced82a23a86725fd1af7e92dd6427d503545"></a>

## client_certificate_optional property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 77824b9cb622 / 4

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

- [crl](resources--workload--reference--group-007.md#canonical-47c84e0045b3a276c6ee475bc7d863750afbea19e806a3c687afb867f1c08baa): complete subsection reference.

- [no_crl](resources--workload--reference--group-007.md#canonical-f2b4319480cbe6e3f7c8a57aecd23f355ab06617b748e315270233630dc743e0): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-007.md#canonical-f6a42a552a2237af6f801cafc2ac0163f6ba84debbc2cecaec370f0ed9f935ea): complete subsection reference.

<a id="canonical-24395802aced8bdbd46167718af875f2a8558a293455f2e7d065a4cfe86d5956"></a>

<a id="canonical-0a8719c0cf66be69d1f36784b0b39ca23f1445928a4fc7d2279d2528767e915f"></a>

## trusted_ca_url property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 77824b9cb622 / 5

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

- [xfcc_disabled](resources--workload--reference--group-007.md#canonical-6fcf47eaafc701303c627d5ef6ee36edc335f6d4aca2f44516480d0f5c4735ea): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-007.md#canonical-375c94bf013cf621a955e439183c31868aad1a0afcb2b1cbfa1f0e2bc09506c4): complete subsection reference.

<a id="canonical-ebc866924b2bb265edcfde0e2341300cf9079332439c240730be5eb2acd8e8b7"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 77824b9cb622 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl](resources--workload--reference--group-007.md#canonical-47c84e0045b3a276c6ee475bc7d863750afbea19e806a3c687afb867f1c08baa)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl](resources--workload--reference--group-007.md#canonical-f2b4319480cbe6e3f7c8a57aecd23f355ab06617b748e315270233630dc743e0)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca](resources--workload--reference--group-007.md#canonical-f6a42a552a2237af6f801cafc2ac0163f6ba84debbc2cecaec370f0ed9f935ea)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled](resources--workload--reference--group-007.md#canonical-6fcf47eaafc701303c627d5ef6ee36edc335f6d4aca2f44516480d0f5c4735ea)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options](resources--workload--reference--group-007.md#canonical-375c94bf013cf621a955e439183c31868aad1a0afcb2b1cbfa1f0e2bc09506c4)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-47c84e0045b3a276c6ee475bc7d863750afbea19e806a3c687afb867f1c08baa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b7fa1bd57b4b7600bc354b91a063f65a135133b1ae24f01c1ad61bb7b8dd9d36"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 9bdbfb057710 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-f2413b81579d597750a66a221c5397266f107b3d2106e3a965998202c3dfc5ec"></a>

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

<a id="canonical-eb23ea624dc41552687e9d2bd20f02bc9c134c6193af09872fbdfe0669db0d92"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 9bdbfb057710 / 3

<a id="canonical-c47d9d81b5b176b1fa61981a5494f2fcce7105d6b905bf5ab42d56dd88cdba85"></a>

<a id="canonical-b8ad565e9b5c9bcbb5e735d55d2a0cc5c4b7530c1be9856d9d0dae6cbbd7fa9b"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 9bdbfb057710 / 4

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

<a id="canonical-e183d1aa0d6a20718997658f68b8704a5231cd8da29fad016df330165e633d31"></a>

<a id="canonical-16e4c65dfdcc9826a478310d53e472fed3dd02fae6cd6f7979f8c1bf020cb831"></a>

## namespace property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 9bdbfb057710 / 5

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

<a id="canonical-4efa8bc28fcec2b68afdeeb12b3ae9b14da4e5feedd426efedc1d09032c72777"></a>

<a id="canonical-209cb107b6cb25d01e8c75efa09463dc6ae0c3863402120e3e4a326f426bf24e"></a>

## tenant property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 9bdbfb057710 / 6

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

<a id="canonical-b8600c91045b5566b8efce0e8d0af50f4334f74259c6449b5d683910ed176ad5"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 9bdbfb057710 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f2b4319480cbe6e3f7c8a57aecd23f355ab06617b748e315270233630dc743e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3dd5ccccbb8e52b93080173de527020e6eb45458d22d490c9527f2b76cb4c03"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1b62eff053a0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-436068e414c954a130ae81042520400668f9fec6708b388e0364d3eea489b32e"></a>

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

<a id="canonical-f82f4f1a71a54702f1ee77bbb15cec80677935fe4343f4063fad8964b9ae8d11"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1b62eff053a0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5abdf1fefaa35b2c4f3d08b7a0b4d7c013b693da7e0a33705e4237728301a32"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 1b62eff053a0 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f6a42a552a2237af6f801cafc2ac0163f6ba84debbc2cecaec370f0ed9f935ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed434dddf44f13f51596a914f526824e4916b56a688c8180f5f65d138f674a8f"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / e9a06cffe3c8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-052f4df8eeb50da9b758a65d1fb226e16dca590fe4df1e968f0851a1e1018bac"></a>

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

<a id="canonical-bac8316be575954f47790d93034a2061b2723b004ce080bcafb927e0e8d31f40"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / e9a06cffe3c8 / 3

<a id="canonical-5d84832999786017981f376842eb07ee8d904d84143e96b9e4d67b2f65ce7819"></a>

<a id="canonical-5c01afc31cc2ce5ddf2c908c3d62ea795bcfd6968d4b48b6b0f3873735e67175"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / e9a06cffe3c8 / 4

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

<a id="canonical-d8d43b9453def98abf379af50024cbff161a2e78fb2ea6f0a2ec9b3c00f97434"></a>

<a id="canonical-a7d8550076e7a477ed08b3ea742260cef487ab737de77b93c241051e8c6815d8"></a>

## namespace property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / e9a06cffe3c8 / 5

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

<a id="canonical-78a50f418435e240aeca74e29fe3cfb18e8a6a1722ff37c496d57e0706bccf7b"></a>

<a id="canonical-8a1348c313c0f9711687d3cdb997f05af72d82eaef157ccc0f9060c2653f8589"></a>

## tenant property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / e9a06cffe3c8 / 6

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

<a id="canonical-8c4ed6809cce2fbe510803c7873db3e02c0f94f5e6afbafeca873a44ca206acc"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / e9a06cffe3c8 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6fcf47eaafc701303c627d5ef6ee36edc335f6d4aca2f44516480d0f5c4735ea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fe66c55a38b43e555296d6ae690fcbab9e9d1b0570cb6b4b7936f0a91d13ded"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 5daabc474a26 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-e603dea22caa9554d65837fb610ef8c390137422892506b6c9cc8be7011fce26"></a>

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

<a id="canonical-1f62add4fdd382caa98afb388a7450322b1f3931927cc5886082b9e1e3240747"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 5daabc474a26 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8b7edef38607c6b7a04ff33c05664dc41a6e39a285894bdaf34c38898219f943"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / 5daabc474a26 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-375c94bf013cf621a955e439183c31868aad1a0afcb2b1cbfa1f0e2bc09506c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7741648577e2e4ee2871548b4977d919e91c3df7358c4d8a4904dd02ec53beb3"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / aba0ea04931a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-006.md#canonical-d9f795d4a5122d5f07348cd3c91d65f642537fe907205a966d3bd9103ef5a336)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-abaa60afb45848f55ff96dda411f927c4afe9c7e1912575aff94894ce60b3f68"></a>

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

<a id="canonical-1c1286935a31a457554380456798f3e22e12efc97693baa52f5ca913a9c85739"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / aba0ea04931a / 3

<a id="canonical-515547ad3cb4e525562ebaa2f89a645deecc2e2b3e581130446149ea06f790de"></a>

<a id="canonical-db1f749fede9abe2e80228a5a4d6ec8a3a9f541e102993de0427f6beb014c29b"></a>

## xfcc_header_elements property — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / aba0ea04931a / 4

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

<a id="canonical-f9d0c46895c1f037e323725573c0661684358a43c425806e9e83ae3335e22e4e"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_ce / aba0ea04931a / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-092d1e08366c8084b92ca674b7a34cc11dd021b86ec8ccf36e209e77b71c4a2c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c5972512540fc706bc0ed09b7f9a33e566440963ffde68a75abbe8fe2853eeee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d103a90b46ef9d5f41a77f998a1cacebf88fbba3209a856292509feba88a3c26"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / e078dcec1ede / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes

<a id="canonical-68befc0e17ea4e851b9018fbee4d4dfdfc71bfcd4459340385fa19bf980b985b"></a>

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

<a id="canonical-170b247d63f3ece396bb2ed95f4c267845a47bec9e77ea47c7089523af12795d"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / e078dcec1ede / 3

- [routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5): complete subsection reference.

<a id="canonical-2a505aa87c930be0435cc44682c83924329030ec39196a79793ca41b23da920a"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / e078dcec1ede / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74a603f717c46cfb406356fccef196f9350c4f760cea5d53d0df3cd648a23ee7"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0c6f39a4d639 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-c5972512540fc706bc0ed09b7f9a33e566440963ffde68a75abbe8fe2853eeee)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

<a id="canonical-3862f171a037c094e9675cbfc366eb34feedaefa8afb796381672f0d75513cea"></a>

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

<a id="canonical-5af321e51ceb0cfd2577b256964731ffdbe01c5852e564cb033feb58f2221edc"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0c6f39a4d639 / 3

- [custom_route_object](resources--workload--reference--group-007.md#canonical-806ba188e2e33e272d1fde298b342b0f96b0855d3081727bc8bbcda02b1f9196): complete subsection reference.

- [direct_response_route](resources--workload--reference--group-007.md#canonical-edebc88b0590436738dd45e43bca31c34375bfd78d3062c9e14068a30cc7676a): complete subsection reference.

- [redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17): complete subsection reference.

- [simple_route](resources--workload--reference--group-008.md#canonical-beb33c7c18adc7d99d4cf91400f4de949239b18a006453ed6ae00e62fb91f142): complete subsection reference.

<a id="canonical-cdf7b932e868522fd9db35db317911de1f5187589223309e18636e34ac093a03"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0c6f39a4d639 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-007.md#canonical-806ba188e2e33e272d1fde298b342b0f96b0855d3081727bc8bbcda02b1f9196)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-007.md#canonical-edebc88b0590436738dd45e43bca31c34375bfd78d3062c9e14068a30cc7676a)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-beb33c7c18adc7d99d4cf91400f4de949239b18a006453ed6ae00e62fb91f142)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-c5972512540fc706bc0ed09b7f9a33e566440963ffde68a75abbe8fe2853eeee)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-806ba188e2e33e272d1fde298b342b0f96b0855d3081727bc8bbcda02b1f9196"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e03a99b102ace2b1edee8c90cbaaa41fba33235178a2a5071e52668d1b0aa754"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 74d07bff4b81 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-c5972512540fc706bc0ed09b7f9a33e566440963ffde68a75abbe8fe2853eeee)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-67111f195e97363ba65b403a252605e44ec5ac87085b316c0118bcba99f13ac1"></a>

Type: `"object"`. single nested block, Optional.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit")}
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
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

Terraform syntax:

```terraform
custom_route_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-529c58664eb4f7b28e0a7906dc46616f2a3478e26e09d3b4a04050412c6e9df5"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 74d07bff4b81 / 3

- [caching_disable](resources--workload--reference--group-007.md#canonical-871aca09c568fd411145f0f845e43bef544888c5c9a63ed6b7cf4a8b00a4e098): complete subsection reference.

- [caching_inherit](resources--workload--reference--group-007.md#canonical-ec59536d82574f0a2b0ef783a67d0661cead8436816ca68b9959b0e6d289f212): complete subsection reference.

- [route_ref](resources--workload--reference--group-007.md#canonical-d7471b7c61fbfb5c9a4adb198fd2a76409c9ebe80e9dda03a15249a96f6c1fa8): complete subsection reference.

<a id="canonical-5b54967e1c00706ef0666006988a17fa6c28ef75e3dcbc7eeab59ee21ee2079f"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 74d07bff4b81 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](resources--workload--reference--group-007.md#canonical-871aca09c568fd411145f0f845e43bef544888c5c9a63ed6b7cf4a8b00a4e098)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](resources--workload--reference--group-007.md#canonical-ec59536d82574f0a2b0ef783a67d0661cead8436816ca68b9959b0e6d289f212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](resources--workload--reference--group-007.md#canonical-d7471b7c61fbfb5c9a4adb198fd2a76409c9ebe80e9dda03a15249a96f6c1fa8)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-871aca09c568fd411145f0f845e43bef544888c5c9a63ed6b7cf4a8b00a4e098"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c675f7b30eeccd33e54bec7cbc835f1e40a6ec269bc004cef189e0aedaf5c48"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f515762cd942 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-c5972512540fc706bc0ed09b7f9a33e566440963ffde68a75abbe8fe2853eeee)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-007.md#canonical-806ba188e2e33e272d1fde298b342b0f96b0855d3081727bc8bbcda02b1f9196)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-d102e14346d2754db50406558adaf9fb02ebf580bb95137812bf8f7bb6259533"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching disable.

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
caching_disable = {}
```

<a id="canonical-9a13eb0af7e9746eb2e98ddbd71f953bc3fb6bd5b5f4c7b0f8551a4aaac99d6d"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f515762cd942 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-05469640a7a04e3f9c35276851e7b27ca4fb2cf587b2cdf91471c48b8292ba4e"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f515762cd942 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-007.md#canonical-806ba188e2e33e272d1fde298b342b0f96b0855d3081727bc8bbcda02b1f9196)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ec59536d82574f0a2b0ef783a67d0661cead8436816ca68b9959b0e6d289f212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfc6dcad58f9778e4d2877efd88a192dc45d34a5f34bb07ed9875832eaa7f1b5"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 727e94ee5a26 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-c5972512540fc706bc0ed09b7f9a33e566440963ffde68a75abbe8fe2853eeee)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-007.md#canonical-806ba188e2e33e272d1fde298b342b0f96b0855d3081727bc8bbcda02b1f9196)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-4034260bc49338fa85108fc9fa48ffba9984e2e78ea076468a9f92f4062ab8bc"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching inherit.

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
caching_inherit = {}
```

<a id="canonical-dc9fc9ef9c435cb9d136e5272ea58083db64a8fa4b9102810d875c19d1f88c83"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 727e94ee5a26 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f07e436f09f01f48225a27f5bc030b90f896273331a87351b099ac4502691028"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 727e94ee5a26 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-007.md#canonical-806ba188e2e33e272d1fde298b342b0f96b0855d3081727bc8bbcda02b1f9196)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d7471b7c61fbfb5c9a4adb198fd2a76409c9ebe80e9dda03a15249a96f6c1fa8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae71df6c858328618db1547eb58f9994929124d510b220be9c34f1f531e435d4"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c1ae97444dd7 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-c5972512540fc706bc0ed09b7f9a33e566440963ffde68a75abbe8fe2853eeee)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-007.md#canonical-806ba188e2e33e272d1fde298b342b0f96b0855d3081727bc8bbcda02b1f9196)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-d10000ebf7c5ade0e6a9985624feeeafc48f6dca07443162493c8e480d6221bb"></a>

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
route_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-ebf9c0e4e5f2345a8f99bcf2a91c77a2f99eb4d58a7220cf2d2be03f290d5d27"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c1ae97444dd7 / 3

<a id="canonical-70b1e3906f4f6b197cb45b382d75e2ad3dcb9bd5260ba5a827b6439d1308de33"></a>

<a id="canonical-88d9ba885dad20437fb3fb4a3a8e026399a9d08d14a3b73a9e2eb72b83207035"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c1ae97444dd7 / 4

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

<a id="canonical-629d5f4784c022a249bfd0e76161f13df77c5dab7726a42066123b4098295258"></a>

<a id="canonical-d71bd5faf94a5cfeb6804799701cc2e6c32fd4c72c24af46117f4c974c3a4351"></a>

## namespace property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c1ae97444dd7 / 5

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

<a id="canonical-397e94325eb232f2cb0828606dca6b85637adfdf254a2cbf41cc362a2eabc726"></a>

<a id="canonical-407b55915b48a1e87f649e519c48d9b8b40f1ebed7a3334d43f34624b0b14e33"></a>

## tenant property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c1ae97444dd7 / 6

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

<a id="canonical-6b7b137f86deb298ec6a90dea1ad8920f951f6887eb4b3903ddab9fc63d014dc"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c1ae97444dd7 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-007.md#canonical-806ba188e2e33e272d1fde298b342b0f96b0855d3081727bc8bbcda02b1f9196)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-edebc88b0590436738dd45e43bca31c34375bfd78d3062c9e14068a30cc7676a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3eedb4bcaaff430ca6a035afc9fc0743afb44f390ed93e50ddb3db0abae436bc"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0eae35e698b5 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-c5972512540fc706bc0ed09b7f9a33e566440963ffde68a75abbe8fe2853eeee)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-e550940b29a0a75e1aa41d97e2cca1081936af81715a860ecc6cd637658aa040"></a>

Type: `"object"`. single nested block, Optional.

Direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Upstream description:

A direct response route matches on path, incoming header, incoming port and/or HTTP method and
responds directly to the matching traffic.

Receipt-pinned upstream constraints:

```json
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
direct_response_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-a508d1793b181f9f635c00b446251435692fe4072f7be1cd8150cb0b7e07b4f2"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0eae35e698b5 / 3

- [headers](resources--workload--reference--group-007.md#canonical-88b0f7be410794beaa9e3e3921ab71948bfcc22e355da32d825ce92e1bf3d56c): complete subsection reference.

<a id="canonical-52bde652049fe6e846cb922cf2e7551acee555f978c2f99ec7b9bfb951b34508"></a>

<a id="canonical-87468d667bf593dc8fa5d15ab8b8c7a0670e5b71ed0e163269e2437b32f62d5a"></a>

## http_method property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0eae35e698b5 / 4

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](resources--workload--reference--group-008.md#canonical-9b3db498bd00276619678589ccdfdc4f0cd5cb11c7a27178c49f1cd4be4dd022): complete subsection reference.

- [path](resources--workload--reference--group-008.md#canonical-91020060979e06c9cd859e13aeb1e3be095c82ed72da34d968f8511f2ff386a7): complete subsection reference.

- [route_direct_response](resources--workload--reference--group-008.md#canonical-07b301f98a71bbe07b682f826f181b74cb23bad0e2108ee3f2561eae07369187): complete subsection reference.

<a id="canonical-634b47f98e6e9e4330e2870d748ac4f456f4cb7ada558d1e72553ff3037d38ce"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 0eae35e698b5 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers](resources--workload--reference--group-007.md#canonical-88b0f7be410794beaa9e3e3921ab71948bfcc22e355da32d825ce92e1bf3d56c)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-008.md#canonical-9b3db498bd00276619678589ccdfdc4f0cd5cb11c7a27178c49f1cd4be4dd022)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path](resources--workload--reference--group-008.md#canonical-91020060979e06c9cd859e13aeb1e3be095c82ed72da34d968f8511f2ff386a7)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response](resources--workload--reference--group-008.md#canonical-07b301f98a71bbe07b682f826f181b74cb23bad0e2108ee3f2561eae07369187)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-88b0f7be410794beaa9e3e3921ab71948bfcc22e355da32d825ce92e1bf3d56c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e09e66af9103c3314c9500d5069fa95c0dfc5d0851ef83c44eea756ba27095c"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c2936ea10ab4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-b507b8c28bfc88787ce3a46f4f6d231587902fd89c6836d35fc7b6e8b2729192)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-c5972512540fc706bc0ed09b7f9a33e566440963ffde68a75abbe8fe2853eeee)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-007.md#canonical-edebc88b0590436738dd45e43bca31c34375bfd78d3062c9e14068a30cc7676a)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-6179de5f44aab7e857916f6a47902e34a65ca3270becb97960dd68bad3567299"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-efa2bcaa3f028f4744b2854e7740225f0ee862ef44139a47e6ec051ca02b9418"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c2936ea10ab4 / 3

<a id="canonical-32517c2a120cd4924bcd7533ff00f106ac5407b5945fafe7892cca6772434156"></a>

<a id="canonical-c1152a87386693a0bbfdd92809b7639f876ec6a3aa382edb3a32287917686662"></a>

## exact property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c2936ea10ab4 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

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

<a id="canonical-1a5fd97b96269dbc8f25bdc52040402b60fd944da3d16f8b78b6588da96b4dfc"></a>

<a id="canonical-b555cbec4fcf8e1ffa7312ac652728db1c3199d9a27001ea5845ddea028ac0f3"></a>

## invert_match property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c2936ea10ab4 / 5

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-e897998a3a79c9fe0aa98f237392809d7c760fef87d77e99bb1d2dd1f0ad732d"></a>

<a id="canonical-053dfc9c3b1be87138e5ef9d162c8be2e77e7b427b4d205dc739392c9db21a44"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c2936ea10ab4 / 6

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

<a id="canonical-68e69372b6c6830873144079d36d560ae81062674c19aede337bde1600855fc5"></a>
