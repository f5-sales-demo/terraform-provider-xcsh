---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-8a2486a9ab439bb3f4e1278c5dd48984fc9f7a218f3a69151ca51fb440713313"></a>

## presence property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c2936ea10ab4 / 7

Type: `"bool"`. Optional.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-b0a2184e2c728f189307ed35215fa08a6a89ea01ce6e1ccc858c0eb4fabc4aa9"></a>

<a id="canonical-dfc4fbc0ab50960c877126137fbbbb23cdd89f919ec72a7e4daaac72963f615b"></a>

## regex property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c2936ea10ab4 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2397a3287e31a0050d15ed2bd8498cd83b3a69470e428e46d35036041c8660e1"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / c2936ea10ab4 / 9

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-007.md#canonical-edebc88b0590436738dd45e43bca31c34375bfd78d3062c9e14068a30cc7676a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9b3db498bd00276619678589ccdfdc4f0cd5cb11c7a27178c49f1cd4be4dd022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3340773ce47faa327a0827ab6023c8d483a4448be68fba8c9850e62f3bed774"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ed4c63dda34d / 2

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
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="canonical-b37939f540b523d3a6cce05fb2db9fcb4e19dab2949687415db2ca83d1a1aa5e"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
incoming_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-4ed6579081e6134694ebb143225ac29661f86301a879df876d4b0e440048a040"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ed4c63dda34d / 3

- [no_port_match](resources--workload--reference--group-008.md#canonical-dddb260b0b9e2ea5f8cb10ac90341a82db25c20b0886e3650e63c5924707e93a): complete subsection reference.

<a id="canonical-283eaa807bc01d1ffc5bd5b98d13d57765d8eaf1b052ae3901ee166bf7047a08"></a>

<a id="canonical-4d699ca793e3766210eda26c45f648ba0a0962cb2f0e13cc14890228d5c24906"></a>

## port property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ed4c63dda34d / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-361ede2793c426903b05813f35d96fe405ce08cbc845c363e2cef9afa509c61b"></a>

<a id="canonical-fd8fa2968448343955392ca30f4886e25750f679e16adb61b3cdb14847110ded"></a>

## port_ranges property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ed4c63dda34d / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-f98a263a8fc0c8abe31ab93d192d4bec719f827ffff53406d69ae3fb5fdadbe7"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ed4c63dda34d / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match](resources--workload--reference--group-008.md#canonical-dddb260b0b9e2ea5f8cb10ac90341a82db25c20b0886e3650e63c5924707e93a)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-007.md#canonical-edebc88b0590436738dd45e43bca31c34375bfd78d3062c9e14068a30cc7676a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-dddb260b0b9e2ea5f8cb10ac90341a82db25c20b0886e3650e63c5924707e93a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2b8ae96ebd82985baa516a5f6f394e1e3a91121fc5c851a591395d2e4b3be66"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9ec9f6754f5b / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-008.md#canonical-9b3db498bd00276619678589ccdfdc4f0cd5cb11c7a27178c49f1cd4be4dd022)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-ac7a7c473d6b311ea413c9f1ff622263e70f737953dcc18db01b18fe657015f6"></a>

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
no_port_match = {}
```

<a id="canonical-f70ba20403be257d1f5474adce24bf3e6da71d1e10cc11af66a56b23a1103502"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9ec9f6754f5b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4a18ea46e58478518c67922b0535bf17f9e4845c21dbf8466d256cfca89aa444"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9ec9f6754f5b / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-008.md#canonical-9b3db498bd00276619678589ccdfdc4f0cd5cb11c7a27178c49f1cd4be4dd022)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-91020060979e06c9cd859e13aeb1e3be095c82ed72da34d968f8511f2ff386a7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4092bc53616ec089741cfd4f932345f036cb13a4e890e285b97e97a15e791a20"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f9458e624877 / 2

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
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-9b21da5942c22f08001d7b604f256cf55efce62bfee397eb7af3387a7943e958"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-3f230c0db771d12553aa9d09d72cd6de126541f26acd80094bc661c28a11a869"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f9458e624877 / 3

<a id="canonical-c1dbf417160c668839f59ee053ac1fafd6e65c872b4107ff001f0ef0ace6d757"></a>

<a id="canonical-40a1eba143826eefc6892d7fdc28e7dd4adede07349f3fe816fd39901d4a079f"></a>

## path property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f9458e624877 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-6ba82759da9ee401eb9cc1785f3e82add7869cb7e1d55febe7ccab0a61ae8e7c"></a>

<a id="canonical-8c0f905149a08d788f67b112b59269abe5abb63a7355569f6850806213616c42"></a>

## prefix property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f9458e624877 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-778a32a9bdbcdb0f0638ca4768dde707d44a39ee0fc221f79e393717d3da8f10"></a>

<a id="canonical-3b3f431152bf0d316aeec70a91ac9c13a712e889d42fc06fbadb919ee6c339dd"></a>

## regex property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f9458e624877 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-8b6619136ba34e9d5d5d483a556a03808fcffe942a920ea1b583f9b657ff5c18"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f9458e624877 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-007.md#canonical-edebc88b0590436738dd45e43bca31c34375bfd78d3062c9e14068a30cc7676a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-07b301f98a71bbe07b682f826f181b74cb23bad0e2108ee3f2561eae07369187"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a2dc0270559ca784d08b542832a70a4f5791b4d23a85ef2d6442a6613662ded"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 6d558e90affa / 2

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
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-2f8e8005017130e680b668944437e9de67bfa76f1ce97da92a61156365a22fc1"></a>

Type: `"object"`. single nested block, Optional.

Send this direct response in case of route match action is direct response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_code")}
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
route_direct_response {
  # Configure direct properties listed below.
}
```

<a id="canonical-4fbcefcb1df3739b98ed09bbf5296734b8978ff501777cb18e38b8d92f9fbf14"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 6d558e90affa / 3

<a id="canonical-68cb86dfef44d154ce2fe9659e9aeeeb7a56cc4178713544c3acfe673dd0a781"></a>

<a id="canonical-091b972efaa92e58c0415987f720977f1f3e873fd02d24748b6443d807fa663e"></a>

## response_body_encoded property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 6d558e90affa / 4

Type: `"string"`. Optional.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML.

Upstream description:

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". Base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

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
    "byteLength": {
      "max": 65536
    },
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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0be2a9a76389d657627102d0494bd905aa718df83f141cf076e616a573f278a8"></a>

<a id="canonical-603f9efc7523a0a04f66f68e58b2eff11b4fcbade77f112c201143fe052d0e7a"></a>

## response_code property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 6d558e90affa / 5

Type: `"number"`. Optional.

Response Code. Response code to send.

Upstream description:

Response code to send.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(100, 599),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

<a id="canonical-d5e207fafaf232525e7961abd3797e0c1a64ac221e1a34f55960fb7919f25763"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 6d558e90affa / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-007.md#canonical-edebc88b0590436738dd45e43bca31c34375bfd78d3062c9e14068a30cc7676a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-925dd12acf11fd178ad6b2f0edc41de8b257c5df84c95fc0458c342786a12290"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f3dac9bf176a / 2

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
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-c8a31d99c8010745b4f5b1e00fc628a87f065bbf275c7d53a89db2e542619f49"></a>

Type: `"object"`. single nested block, Optional.

Redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the
matching traffic to a different URL.

Upstream description:

A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects
the matching traffic to a different URL.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
redirect_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-5a90bec93158201507f20db01c04f75f3f5897845b7b1fdbe2472cc4b583cfab"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f3dac9bf176a / 3

- [headers](resources--workload--reference--group-008.md#canonical-6e63b4d3730d5088e3a54438a0d25c6b0066bc4ba917660e26d502fd59f0ae0f): complete subsection reference.

<a id="canonical-6d3e06f62f81b8a04c2138842cd72200eaab6e07fe7aaec2fb96b1d7624eaa36"></a>

<a id="canonical-f794f9b1c9b25b27266d53c5c7e77739d2d3a002d2ef82e0b098f9687269d92e"></a>

## http_method property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f3dac9bf176a / 4

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

- [incoming_port](resources--workload--reference--group-008.md#canonical-48ee1516838cbdc41c510e8890fb1af419b3ddcb4716d59f84ce508fd038f605): complete subsection reference.

- [path](resources--workload--reference--group-008.md#canonical-ef3093f7cd335a4820f4a49777694351a4d6cde988965eb5a10cd3c2f8b5c5b5): complete subsection reference.

- [route_redirect](resources--workload--reference--group-008.md#canonical-47287d5465777ac9483baedb157d1d139e168d987b7151b4f14f1ebfee97503b): complete subsection reference.

<a id="canonical-779cbd1ec04738095f368457a18456f2963b4209ed7c76cefb06ecca0486ccdd"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / f3dac9bf176a / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers](resources--workload--reference--group-008.md#canonical-6e63b4d3730d5088e3a54438a0d25c6b0066bc4ba917660e26d502fd59f0ae0f)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-008.md#canonical-48ee1516838cbdc41c510e8890fb1af419b3ddcb4716d59f84ce508fd038f605)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path](resources--workload--reference--group-008.md#canonical-ef3093f7cd335a4820f4a49777694351a4d6cde988965eb5a10cd3c2f8b5c5b5)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-008.md#canonical-47287d5465777ac9483baedb157d1d139e168d987b7151b4f14f1ebfee97503b)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6e63b4d3730d5088e3a54438a0d25c6b0066bc4ba917660e26d502fd59f0ae0f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46e5f6944b679eddb2ddbf1aca35a305e141e0a458be19ca0b7be472dffded4f"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 46656122f6cc / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-f0b9c8785666af6f48f0deb2808d427ea6161dbf1086331e85cb3c3547ed59fd"></a>

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

<a id="canonical-eb7cef9897777bbd2d1c8ec80906e8ad04f72a53a4097072f2816f3ecbe040da"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 46656122f6cc / 3

<a id="canonical-e6580138c5813fc1154e1a3a2aa2d07e555dac92250f629d7429a23bba6f0c48"></a>

<a id="canonical-142e23bf95164d78d52d7432ad78e1bdd40fbe248d93d6bda07f37bc21f12cc6"></a>

## exact property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 46656122f6cc / 4

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

<a id="canonical-d69d1a694163919e047650e8770a6518ba3edc4fd10d079c7127204df4a632f8"></a>

<a id="canonical-5ad9267a341b0882c8d76f6cf14d535dc244b17e3ab2a82c46bbac55689bc80a"></a>

## invert_match property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 46656122f6cc / 5

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

<a id="canonical-683af23f5b2f38ecff47c121bda9773a0790cde576f29df678a751b631d358ff"></a>

<a id="canonical-18bcc7b5811e5aa33551e258b3cb54459e8698ecb516212e68494a1bb401bd4a"></a>

## name property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 46656122f6cc / 6

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

<a id="canonical-e6a75b0d2f6f47b92beebb8441cca708d666c49050a760a47ea3d2d3130de2ad"></a>

<a id="canonical-349eb26b7dee30b0d0a602aae3f8bb2df1a3f0e588b8b17ed224a7f9b7416fb0"></a>

## presence property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 46656122f6cc / 7

Type: `"bool"`. Optional.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3097db1bb98a4d20066d462884fdf44bbe0848d0f0d17302bb1362f83e9e0af1"></a>

<a id="canonical-5918aca1d4b20faa3af6a314a233e6912ee0ddaa6fc9a18c2dc09c1b07f77d6c"></a>

## regex property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 46656122f6cc / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-e650d500ec5c7a3769fd875f1f444c49c607f38990daef282a257c72da3b51ba"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 46656122f6cc / 9

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-48ee1516838cbdc41c510e8890fb1af419b3ddcb4716d59f84ce508fd038f605"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1898e69a05baa7e519ae96f1be47e61001d77d2f0d26096be29d474e2339e1ae"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9a53200c94c8 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-8b70c3ecb25a27c5134dadfd6bfd4747b3f4531ca26c2d3c88d98092f6ef5c98"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
incoming_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-fe5b85e5e10892f239fa38316729248d01972be5fa167efa402878b5094b2f50"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9a53200c94c8 / 3

- [no_port_match](resources--workload--reference--group-008.md#canonical-506e6bf3512ac5a079441988bc3920162a04f9135e4823e0cb703d2c16e78c5b): complete subsection reference.

<a id="canonical-0751429dc70b8728599279bca1af340679c54611a05ea9e780ffdd63decec5c7"></a>

<a id="canonical-e6869f1baf820a397ae006e831eec476379b0726418fe9f3eb996857ca05ffb0"></a>

## port property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9a53200c94c8 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

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

<a id="canonical-272cb5ddc89304041a01c26657358e6bcf25d501f388371a4ddffc366f96ba00"></a>

<a id="canonical-174ff51757765962318732ecae50e28c1476d1095fe213049ec57d68d8cbe016"></a>

## port_ranges property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9a53200c94c8 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-29fac35cc7e2babc3d1a573ec0702bbc510bd2e7b36419fbfec7d896e034c60e"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 9a53200c94c8 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](resources--workload--reference--group-008.md#canonical-506e6bf3512ac5a079441988bc3920162a04f9135e4823e0cb703d2c16e78c5b)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-506e6bf3512ac5a079441988bc3920162a04f9135e4823e0cb703d2c16e78c5b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43eb43cd4d581af37ac3935837e4da7af35a652b30f7af50b8bcd7cbca8a7c4c"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 41e82f957f72 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-008.md#canonical-48ee1516838cbdc41c510e8890fb1af419b3ddcb4716d59f84ce508fd038f605)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-f5a0f476591cf7478b74096fa162b12f6da2566191beed61bef1e859078e7316"></a>

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
no_port_match = {}
```

<a id="canonical-ac1d7ad7ef1b2704504abba542deece4652d54360aec25116cbff3b48e50753c"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 41e82f957f72 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2caa4d51b45d633a5e61c6a8dcd273bde990f435e91d665475972febd739799"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 41e82f957f72 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-008.md#canonical-48ee1516838cbdc41c510e8890fb1af419b3ddcb4716d59f84ce508fd038f605)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ef3093f7cd335a4820f4a49777694351a4d6cde988965eb5a10cd3c2f8b5c5b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cb244f67a38b32d623d8e32dbe73946e1fecbd74e0f9e3fc0e56fc8814c91f0"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ac7d490d5b99 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-183dbc49df82919539e3f00f47f9b3a957d5842f59a15147b75a26b57e664da2"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-666f0081d90a6e8ccd83fc9fadd11022092302074da42711ae11e0151330dba5"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ac7d490d5b99 / 3

<a id="canonical-a0018ab8747bc1a607d908c3cb74b9d16c91c6db48cd72315ae0bbff27530ca3"></a>

<a id="canonical-ddacaf2069c2a8f8baf0da24a98b4698183f90c4677cac57c71498a52f82a1e8"></a>

## path property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ac7d490d5b99 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-0a7793296e1369b8107eea177956b713ae9f7522c226226b51f732d87b64099d"></a>

<a id="canonical-70c121656e2f67248101b9a353617e3b38932c63084f405fa972493a276f3493"></a>

## prefix property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ac7d490d5b99 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-58a600fe22d21c63da9021a3385e798c9cb6502eac6480c162c2a6c42580feef"></a>

<a id="canonical-6f2dac7a73d4ece5205be61421395de3883370fc6b5f945e047be468c8a670e5"></a>

## regex property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ac7d490d5b99 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-d4bd2cb6bd8022e5a70f9db3b570819f6f338f36e9b437c5b98dfc35f19e5027"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ac7d490d5b99 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-47287d5465777ac9483baedb157d1d139e168d987b7151b4f14f1ebfee97503b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e51b22e0400e9a4a7b92b5327c2777fc91f9de58c78ddb19dda1e2560aca86f9"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / a2b219a2209f / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-97f775a309cde57e23e646d8bf429533b6024c81f67d0725572d3ec351b4bc8b"></a>

Type: `"object"`. single nested block, Optional.

Route redirect parameters when match action is redirect.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path_redirect",
    "prefix_rewrite"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "replace_params"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "retain_all_params"),
  validators.ConflictingObjectAttributes("replace_params",
    "retain_all_params")}
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
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

Terraform syntax:

```terraform
route_redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-124b43f8a0f8befa16c7d60c27b6ccb31f257d61473bb83874255de0e6a402da"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / a2b219a2209f / 3

<a id="canonical-d36ed9e43433d457419b3acbc861a1029770a524e3343603fdadb636e2c9ffe4"></a>

<a id="canonical-84f28828f8e656c3d0858b4dbb018e4cee4e1d3ece7618fffa03b050cf8c924a"></a>

## host_redirect property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / a2b219a2209f / 4

Type: `"string"`. Optional.

Swap host part of incoming URL in redirect URL.

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

<a id="canonical-f8701077834a7915f4dbfa2b5cea8f5e0760b6f54b49ba0c92a8dc3f9ef22b01"></a>

<a id="canonical-ef94830eea6cb5322382222795a129124fcd46ef4178f107a4e8c28f04aec177"></a>

## path_redirect property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / a2b219a2209f / 5

Type: `"string"`. Optional.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Upstream description:

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

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

<a id="canonical-b7fe7a67ddce5f7f12d7a29ea174638d6d4fd457ebd836fe2066c1748d25f958"></a>

<a id="canonical-00395a35b467c9b3176af39bfa15689c14631c90ed7f49a2440ec15f0b3339c4"></a>

## prefix_rewrite property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / a2b219a2209f / 6

Type: `"string"`. Optional.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Upstream description:

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

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

<a id="canonical-90a5654f438e9cb34acf18511b806d2e6c208dee157fe988a92bd2849985babc"></a>

<a id="canonical-09b9ebb4b477ab0fa9b539d84a8062801be66eb25c4b1ed53d0df1bdac5ca97e"></a>

## proto_redirect property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / a2b219a2209f / 7

Type: `"string"`. Optional.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Upstream description:

Swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or
HTTPS When incoming-proto option is specified, swapping of protocol is not done.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("incoming-proto",
    "http",
    "https"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "incoming-proto",
    "http",
    "https"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](resources--workload--reference--group-008.md#canonical-cc06fed167e4dae39cf43dff2e8f2e16c4917261d2f9ce8adaa671514ca1b444): complete subsection reference.

<a id="canonical-5093c5eec4bb98f1bb5d73e85d01d7b50c4a2dc5cbec3ae5b48cf600a14eea2d"></a>

<a id="canonical-3ae52a5b38701b2981e29662f808f7f88bfb7f50bc259348bef6a818ecf92491"></a>

## replace_params property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / a2b219a2209f / 8

Type: `"string"`. Optional.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-eaf8d2fb21356b1eb25b0bf77040e4521308109588217654722af0bdde340cb4"></a>

<a id="canonical-e0072e94921cea729676b2b4e8f45d8ea9067aa6e1e49c0e37248d5b2e129831"></a>

## response_code property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / a2b219a2209f / 9

Type: `"number"`. Optional.

The HTTP status code to use in the redirect response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(599),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
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
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](resources--workload--reference--group-008.md#canonical-fa92edcf2f6850b13a97d5e35abfe9a8a044cc7640db9cc87d16cf691f2d43d4): complete subsection reference.

<a id="canonical-974e0c6fb7bae7520e77279c44dec7d8d5371d25ca71295f1deb804c56812c7c"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / a2b219a2209f / 10

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](resources--workload--reference--group-008.md#canonical-cc06fed167e4dae39cf43dff2e8f2e16c4917261d2f9ce8adaa671514ca1b444)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](resources--workload--reference--group-008.md#canonical-fa92edcf2f6850b13a97d5e35abfe9a8a044cc7640db9cc87d16cf691f2d43d4)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cc06fed167e4dae39cf43dff2e8f2e16c4917261d2f9ce8adaa671514ca1b444"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3d69f7866430e473cd979f31453b4d51403b1523edeae291e43d4b51706edd3"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 1d6b9162b50c / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-008.md#canonical-47287d5465777ac9483baedb157d1d139e168d987b7151b4f14f1ebfee97503b)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-f69d995284cc6697824c6d433ca62c47e9ea003f16026de4285dd18ea1bd549f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for remove all params.

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
remove_all_params = {}
```

<a id="canonical-3d769983a7e53128f1d92256b08c0bb5a83646191ff5576e7b58666d6d93b702"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 1d6b9162b50c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59177e92b530df7177160bc1fffdac44b5cc574b29b5512de3de20ada6c97da2"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 1d6b9162b50c / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-008.md#canonical-47287d5465777ac9483baedb157d1d139e168d987b7151b4f14f1ebfee97503b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-fa92edcf2f6850b13a97d5e35abfe9a8a044cc7640db9cc87d16cf691f2d43d4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-88b9a376819052046ac2dee8ec034abfd250cef0bffcc1a804be594b2fa019ee"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 00c37c9fc803 / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-03212d732fbceebec2b443f677ce9ef6131dd158d254549e772303af55d6fb17)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-008.md#canonical-47287d5465777ac9483baedb157d1d139e168d987b7151b4f14f1ebfee97503b)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-4fc06768a6fef02e7b2b3eeaf29eac04b699ba28043823d2f9723db52c7c29e8"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

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
retain_all_params = {}
```

<a id="canonical-c264566747288fa8bf660b8f2f19f81f4255165387c44c1d84e465ab3b9cac4e"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 00c37c9fc803 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3c02032f8db49debd54dbe279fc4ee4989f19f2fe88be1963bafeeb38efdf222"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 00c37c9fc803 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-008.md#canonical-47287d5465777ac9483baedb157d1d139e168d987b7151b4f14f1ebfee97503b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-beb33c7c18adc7d99d4cf91400f4de949239b18a006453ed6ae00e62fb91f142"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef89fab7380c5d5521733295d69f1d9fd394408c2981b55fea47e324db2fd88e"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 8463fd147a80 / 2

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
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-0bc5056284a65570439457ef288600b6cc3eb0722c2be7e8b3667116ac2827f5"></a>

Type: `"object"`. single nested block, Optional.

Simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Upstream description:

A simple route matches on path and/or HTTP method and forwards the matching traffic to the default
origin pool specified outside.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
simple_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-6a3053597cc1d4d9296c1ebc6ccc32852e9b79f248e6b1c5b317d786949158b7"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 8463fd147a80 / 3

- [auto_host_rewrite](resources--workload--reference--group-008.md#canonical-16b1edfb9a4de25d460818c6404f2d3c0e39b9bd1b1fbd34ac9ab0152847df43): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-008.md#canonical-5558e9aa7deb4d23db70df83c06e944912a0ca25dc49db88c4702a03a69b9eca): complete subsection reference.

<a id="canonical-a74c53d6514f7b8852eb31d50a9edbbb5a479cff917cf2402aae7dfda8f2749a"></a>

<a id="canonical-4a0139af28198e1fc5e6d14ae4ac281ac22e02dfd08b9d361388e43fc2784371"></a>

## host_rewrite property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 8463fd147a80 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-a2a534e7931c5c69a7b257d9d54b7337c90922788b42ffcb9eb6e9441f21a051"></a>

<a id="canonical-9cc68c23c97a4e7219a3ec6fc43774eac5a102433d1a30174e1fc5ce5c063386"></a>

## http_method property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 8463fd147a80 / 5

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

- [path](resources--workload--reference--group-008.md#canonical-e927da709a8999733f6345f28353bd9d789b741951982b9341ef9297e46e6a0d): complete subsection reference.

<a id="canonical-479ea56bc67077fde1408f749022424d55b603300dbce4fa5c5c577d86153bb7"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 8463fd147a80 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](resources--workload--reference--group-008.md#canonical-16b1edfb9a4de25d460818c6404f2d3c0e39b9bd1b1fbd34ac9ab0152847df43)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](resources--workload--reference--group-008.md#canonical-5558e9aa7deb4d23db70df83c06e944912a0ca25dc49db88c4702a03a69b9eca)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path](resources--workload--reference--group-008.md#canonical-e927da709a8999733f6345f28353bd9d789b741951982b9341ef9297e46e6a0d)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-d1ae6606fec2b0503fa341bf4df97f58a49df68a7f3ab6ad8cb64fd053a5d7f5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-16b1edfb9a4de25d460818c6404f2d3c0e39b9bd1b1fbd34ac9ab0152847df43"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-239bc823414161248d8a34bc6292c6a81cca347e2d65a32bf7b0e6435d0fe50a"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 3d267de1e26b / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-beb33c7c18adc7d99d4cf91400f4de949239b18a006453ed6ae00e62fb91f142)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-c80726e8ce240c4436c7aa20fedb6d2ea6c545383fcb90598caf1736eb99b698"></a>

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
auto_host_rewrite = {}
```

<a id="canonical-b14e1df1ca61f2f93a91d74ccd45c77d4e23a381114b1e70e5b4ab332a0aa5b7"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 3d267de1e26b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-75eb100efdad3a748db53c79817d4667d6d0eb2d866dbf00f55107836f1d3154"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 3d267de1e26b / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-beb33c7c18adc7d99d4cf91400f4de949239b18a006453ed6ae00e62fb91f142)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-5558e9aa7deb4d23db70df83c06e944912a0ca25dc49db88c4702a03a69b9eca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-84a814b337a7e81d3d761b2784f461a6e13e342981fd710be028328624f41c96"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 7e142d508a2c / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-beb33c7c18adc7d99d4cf91400f4de949239b18a006453ed6ae00e62fb91f142)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-c7761c95e04fa318e537b98a8bb1aea31e91c4aadad50b9c10ad5fd1c2406346"></a>

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
disable_host_rewrite = {}
```

<a id="canonical-5b29da94757d9f44df01a2c3668e591140de77f0ee81a7b274e9173edfa2cc15"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 7e142d508a2c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9533aef90dc508a538ac1e8d3f3da59917a972099376aae233a00f57362e7490"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / 7e142d508a2c / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-beb33c7c18adc7d99d4cf91400f4de949239b18a006453ed6ae00e62fb91f142)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e927da709a8999733f6345f28353bd9d789b741951982b9341ef9297e46e6a0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aef36aeaa75f26a89355f3bcfd48af83cb8fd25b5f013ecc9bb25f47d72b73eb"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ffbb02f8172c / 2

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
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-beb33c7c18adc7d99d4cf91400f4de949239b18a006453ed6ae00e62fb91f142)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-af07f07eb3572cbe22cba6e85c5f6a477e524669e86391c1b9f5cd73d72e98aa"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-79dee4480dfc7b597f4a26aed1a4674ad6d6b2f4e9ece814200346afb68c8552"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ffbb02f8172c / 3

<a id="canonical-3a7ce82e9e3ae2b1c2f2ba2bac3c35baee5a7f7c39099106c473445ffc3ac518"></a>

<a id="canonical-a431cf06d6b7464e9b31949bb587f20e310dcf9b914d711b42830a3292381093"></a>

## path property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ffbb02f8172c / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-cf7e548d2409d0c64123d0b75a9934df6c1895ca3c6be6c4d7f6c01e2e002c65"></a>

<a id="canonical-3a700834d49b5d91da0235dbc221cd994047e7908c2da93ae420c35436dd6137"></a>

## prefix property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ffbb02f8172c / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-3cede2b9e33b4bc803b5c711c2612914636d75e7e898f1461790f780d9b75306"></a>

<a id="canonical-a6d54831098392c6ecd99aee6b0ebd6fbb759e37a45b85a62fc2fdd8f75ddc07"></a>

## regex property — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ffbb02f8172c / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-bbba697673e3110621fa32c743a28d1bf990368efefeddf50e57b7fafbc8a27b"></a>

## Next pages — service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_rout / ffbb02f8172c / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-beb33c7c18adc7d99d4cf91400f4de949239b18a006453ed6ae00e62fb91f142)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-85b6f4e25c7f1b69d65aa7051f4382aac22e4c6dfbe3dde9ecc33d5532af381b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80d9d2557a09003ac1ab81f9048fea754f849eb321de475a9f0d4f3e77770b54"></a>

## service.advertise_options.advertise_custom.ports.port — service.advertise_options.advertise_custom.ports.port / 255bfd4e9ff3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- service.advertise_options.advertise_custom.ports.port

<a id="canonical-33a34368ad9431199c44eba9c146729a3cd45e16a5e98a8c9fd7be305395c376"></a>

Type: `"object"`. single nested block, Optional.

Port. Port of the workload.

Upstream description:

Port of the workload.

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
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-cfcce39482f3e7de8815a47d78091dc57132a426b114da7bf416c29cf796ff2c"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.port / 255bfd4e9ff3 / 3

- [info](resources--workload--reference--group-008.md#canonical-5abac3580a8e7ed42f15b1d1cf33ab7502c34c66e7f810921301ede4cdf7f979): complete subsection reference.

<a id="canonical-0b009bfae46b4b03df2341953cf786b8e95e8307c9e8406da2cf52843ab64708"></a>

<a id="canonical-bda4b18f60a03dbbd9ac633f792f23664d9159bfd153c6fa1c6f7f05457ade02"></a>

## name property — service.advertise_options.advertise_custom.ports.port / 255bfd4e9ff3 / 4

Type: `"string"`. Optional.

Name. Name of the Port.

Upstream description:

Name of the Port.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-f7d06460ee0b41b42f8916559933813857075a86b0d3852b08dcf201db030693"></a>

## Next pages — service.advertise_options.advertise_custom.ports.port / 255bfd4e9ff3 / 5

- [service.advertise_options.advertise_custom.ports.port.info](resources--workload--reference--group-008.md#canonical-5abac3580a8e7ed42f15b1d1cf33ab7502c34c66e7f810921301ede4cdf7f979)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-5abac3580a8e7ed42f15b1d1cf33ab7502c34c66e7f810921301ede4cdf7f979"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7fa6416a2741b9b5540cdae0c1b93fa0d12253e8f6fa729062fbb8eccebe258"></a>

## service.advertise_options.advertise_custom.ports.port.info — service.advertise_options.advertise_custom.ports.port.info / 5243eba3f822 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-008.md#canonical-85b6f4e25c7f1b69d65aa7051f4382aac22e4c6dfbe3dde9ecc33d5532af381b)
- service.advertise_options.advertise_custom.ports.port.info

<a id="canonical-2722207873e5d26d390f1dc9632ddce8b0e4f7023e53cc59219cf8e65004848d"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Upstream description:

Port information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port"),
  validators.ConflictingObjectAttributes("same_as_port",
    "target_port")}
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
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

Terraform syntax:

```terraform
info {
  # Configure direct properties listed below.
}
```

<a id="canonical-b73fb2f10188c5eadf54acb3850e81e2058ba04a9bf33acebe2b2d463acff3c8"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.port.info / 5243eba3f822 / 3

<a id="canonical-d9ff9a80839832154c08534cd28efd36a5f177ea6c1ea380a117ba180df74eff"></a>

<a id="canonical-359300562941340868b8e1d0a98ffe40f6a0564f0cb8a35bc97e33491c15fb50"></a>

## port property — service.advertise_options.advertise_custom.ports.port.info / 5243eba3f822 / 4

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-9c5fe7c6f95b8717a9787bf16b785954f59ed5c38af4a046e7e04452d454e1f2"></a>

<a id="canonical-cd4de67b3e6ad3897da5720ca6516fff05d865dbdec9766a170207990731504c"></a>

## protocol property — service.advertise_options.advertise_custom.ports.port.info / 5243eba3f822 / 5

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](resources--workload--reference--group-008.md#canonical-54318a718f62424111356ef8277a92d32b2b1efa4a9d3d13853725c0454fdede): complete subsection reference.

<a id="canonical-67ef48ffd17049ae5fd5370eca0e53354702cbfc3765561909cab11b1d382234"></a>

<a id="canonical-af7ff64395693b72c9d7d089713f5f30765827f75e5b15983b3b7616ed79e39e"></a>

## target_port property — service.advertise_options.advertise_custom.ports.port.info / 5243eba3f822 / 6

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-a336b376183d0fb44866df8bc25dedade6f091475675d7c6777080117eb41869"></a>

## Next pages — service.advertise_options.advertise_custom.ports.port.info / 5243eba3f822 / 7

- [service.advertise_options.advertise_custom.ports.port.info.same_as_port](resources--workload--reference--group-008.md#canonical-54318a718f62424111356ef8277a92d32b2b1efa4a9d3d13853725c0454fdede)
- [service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-008.md#canonical-85b6f4e25c7f1b69d65aa7051f4382aac22e4c6dfbe3dde9ecc33d5532af381b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-54318a718f62424111356ef8277a92d32b2b1efa4a9d3d13853725c0454fdede"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf80c220ff8ac59f740192debd99ba17148c9a7144a433dda5e68a7378546df1"></a>

## service.advertise_options.advertise_custom.ports.port.info.same_as_port — service.advertise_options.advertise_custom.ports.port.info.same_as_port / 1b669d7f0c0b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-008.md#canonical-85b6f4e25c7f1b69d65aa7051f4382aac22e4c6dfbe3dde9ecc33d5532af381b)
- [service.advertise_options.advertise_custom.ports.port.info](resources--workload--reference--group-008.md#canonical-5abac3580a8e7ed42f15b1d1cf33ab7502c34c66e7f810921301ede4cdf7f979)
- service.advertise_options.advertise_custom.ports.port.info.same_as_port

<a id="canonical-7c7e4336105910520631877ba7d19f64fe1fb73c6ec717134010f1655d1061df"></a>

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
same_as_port = {}
```

<a id="canonical-9a7b044a0b7446e6a920aceb1863eca78faeb9a86ca8f0e01cfc8f8847f7f282"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.port.info.same_as_port / 1b669d7f0c0b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-54d3227a1e9abd9502ba88953a13f52a4c6b1fa59f8f6042efb83f7d6bbe4913"></a>

## Next pages — service.advertise_options.advertise_custom.ports.port.info.same_as_port / 1b669d7f0c0b / 4

- [service.advertise_options.advertise_custom.ports.port.info](resources--workload--reference--group-008.md#canonical-5abac3580a8e7ed42f15b1d1cf33ab7502c34c66e7f810921301ede4cdf7f979)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bf03b979f8a27555c05ed659a5a2b154f3304086fa331a9795acd884385be280"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18a9cc0beff393963f9a08b302d37c9eebd8acac0a134b7343aa5e07b508330a"></a>

## service.advertise_options.advertise_custom.ports.tcp_loadbalancer — service.advertise_options.advertise_custom.ports.tcp_loadbalancer / df81920aa788 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-99b3633cc45edec16f1cee3a4e8c29358a2425ee5bd8e581064e7b1256d41508)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- service.advertise_options.advertise_custom.ports.tcp_loadbalancer

<a id="canonical-902d18a9595851b944f8c5756e19afbd85adfdbef586d9f700f9b407f1e413d1"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tcp loadbalancer.

Upstream description:

TCP loadbalancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
tcp_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-c93b906581ff7e349a8bcfcd6364209a804121b38502e111bedc90edb6e378d2"></a>

## Direct properties — service.advertise_options.advertise_custom.ports.tcp_loadbalancer / df81920aa788 / 3

<a id="canonical-f04f4466cbaa0a7304aaed367a849b96941b2c75d7198f6e6d689bf4bff1c460"></a>

<a id="canonical-a64969d323d1c37d8432ca90d857856f35bae7157f2543f274c0dc51b3a38bf4"></a>

## domains property — service.advertise_options.advertise_custom.ports.tcp_loadbalancer / df81920aa788 / 4

Type: `["list", "string"]`. Optional.

List of additional domains (host/authority header) that will be matched to this loadbalancer.
Domains are also used for SNI matching if the is true Domains also indicate the list of names for
which DNS resolution will be done by VER.

Upstream description:

A list of additional domains (host/authority header) that will be matched to this loadbalancer.

Domains are also used for SNI matching if the \`with\_sni\` is true Domains also indicate the list
of names for which DNS resolution will be done by VER.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-7075ac883c290df5869a0643414468387f6d5630080fb2de128c14fd5795e1d5"></a>

<a id="canonical-dbfcf780c65d71c64e93963038812b4ed8a3cf6a537c50502d0e9b5cd5b73c98"></a>

## with_sni property — service.advertise_options.advertise_custom.ports.tcp_loadbalancer / df81920aa788 / 5

Type: `"bool"`. Optional.

Set to true to enable TCP loadbalancer with SNI.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2d007b1b375cb88eb21cb3b42c92a6fba9080a292627495cda73c3ee127a318d"></a>

## Next pages — service.advertise_options.advertise_custom.ports.tcp_loadbalancer / df81920aa788 / 6

- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-1fa51fe42a4c906ee610e1d3fd60e1a9468c3b946f4bb06fafe598511adda920)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ca0d5f003333d84ec10c373bcbe87cf86908b3bc8697941cf3671a202dce8f1"></a>

## service.advertise_options.advertise_in_cluster — service.advertise_options.advertise_in_cluster / a005e61a65f5 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- service.advertise_options.advertise_in_cluster

<a id="canonical-f31245e3006df489a8466ff38a92e8b65990ce019a3652b909787b6ead7e2cf7"></a>

Type: `"object"`. single nested block, Optional.

Advertise the workload locally in-cluster.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("multi_ports",
    "port")}
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
  "x-ves-oneof-field-port_choice": "[\"multi_ports\",\"port\"]"
}
```

Terraform syntax:

```terraform
advertise_in_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-e8d8390c3d9d96ba03d6f9a7d43a10ea08146cb0b86c6c021373c4c75b5fd783"></a>

## Direct properties — service.advertise_options.advertise_in_cluster / a005e61a65f5 / 3

- [multi_ports](resources--workload--reference--group-008.md#canonical-ef48cef304055faaa9f30a0ebabc2c3a4c3f9fed60c014f27d132eb2ad5d356e): complete subsection reference.

- [port](resources--workload--reference--group-008.md#canonical-8e459cc553aaf03c23bb5044ba63b480e23fea5bb45eb90e1daa1a4f961f99dc): complete subsection reference.

<a id="canonical-10753f50d805cc510840313294226dce40530990eae60420024eb6d3f2bdba25"></a>

## Next pages — service.advertise_options.advertise_in_cluster / a005e61a65f5 / 4

- [service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-008.md#canonical-ef48cef304055faaa9f30a0ebabc2c3a4c3f9fed60c014f27d132eb2ad5d356e)
- [service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-008.md#canonical-8e459cc553aaf03c23bb5044ba63b480e23fea5bb45eb90e1daa1a4f961f99dc)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ef48cef304055faaa9f30a0ebabc2c3a4c3f9fed60c014f27d132eb2ad5d356e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9989d07bbc567f6c268de3299e97b295f46fe3f704f41b275a2308338e729922"></a>

## service.advertise_options.advertise_in_cluster.multi_ports — service.advertise_options.advertise_in_cluster.multi_ports / 05c132bb02d4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967)
- service.advertise_options.advertise_in_cluster.multi_ports

<a id="canonical-043b5df2905815303e1bfd4771577f34cfb84e2576d372915b53c6e90a904f20"></a>

Type: `"object"`. single nested block, Optional.

Multiple Ports. Multiple ports.

Upstream description:

Multiple ports.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
multi_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-ed5dafd1d86a2fe1cf02db55e61b758c6cdf6688fc930310acd2d84bead7a467"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.multi_ports / 05c132bb02d4 / 3

- [ports](resources--workload--reference--group-008.md#canonical-32b2247fcf507ad424d4baea2333e49fa1ff14d840cb241caacc7c8a4a62b7f6): complete subsection reference.

<a id="canonical-1ca5258aa965f23d504ccb4d0287856140d03c8d4b50eef0e23406734bd07dcf"></a>

## Next pages — service.advertise_options.advertise_in_cluster.multi_ports / 05c132bb02d4 / 4

- [service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-008.md#canonical-32b2247fcf507ad424d4baea2333e49fa1ff14d840cb241caacc7c8a4a62b7f6)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-32b2247fcf507ad424d4baea2333e49fa1ff14d840cb241caacc7c8a4a62b7f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d1557962e63403a8acb27b3276adc306ada8f3bd517f0f92bc43cc64ab6e3013"></a>

## service.advertise_options.advertise_in_cluster.multi_ports.ports — service.advertise_options.advertise_in_cluster.multi_ports.ports / 0a361ad9351f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967)
- [service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-008.md#canonical-ef48cef304055faaa9f30a0ebabc2c3a4c3f9fed60c014f27d132eb2ad5d356e)
- service.advertise_options.advertise_in_cluster.multi_ports.ports

<a id="canonical-f4a2d45d9a8938de8e3b6ed61555b11424454f4ff5a12e84b017fd10c985f3ea"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

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
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-8fdc004cdc54abd2f4fcf54c5c07847be189af3e362b74c2722ca3fa39024f3f"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.multi_ports.ports / 0a361ad9351f / 3

- [info](resources--workload--reference--group-008.md#canonical-4ef34860d148a65122d5905c748ea00557e5a73d17af81cc461b8ee59584a87d): complete subsection reference.

<a id="canonical-253992263845043d8c08e1659f60f921618a3640b3eaff30c5c10a4c720f0f56"></a>

<a id="canonical-f8d4e6c6c6473fb0bf8c244dd65ed80a9a5a2c281883b4585a64d747ce47aa8c"></a>

## name property — service.advertise_options.advertise_in_cluster.multi_ports.ports / 0a361ad9351f / 4

Type: `"string"`. Optional.

Name. Name of the Port.

Upstream description:

Name of the Port.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-d92cfa0e1c3b5ec18ca882f5be8b63d26626383896aac582d984544ec492d221"></a>

## Next pages — service.advertise_options.advertise_in_cluster.multi_ports.ports / 0a361ad9351f / 5

- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](resources--workload--reference--group-008.md#canonical-4ef34860d148a65122d5905c748ea00557e5a73d17af81cc461b8ee59584a87d)
- [service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-008.md#canonical-ef48cef304055faaa9f30a0ebabc2c3a4c3f9fed60c014f27d132eb2ad5d356e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4ef34860d148a65122d5905c748ea00557e5a73d17af81cc461b8ee59584a87d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4314b586e8cb94df4e9a70ed68d6714015658f69a16c370c7dd100b929d6322"></a>

## service.advertise_options.advertise_in_cluster.multi_ports.ports.info — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / ed27bb44aea9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967)
- [service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-008.md#canonical-ef48cef304055faaa9f30a0ebabc2c3a4c3f9fed60c014f27d132eb2ad5d356e)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-008.md#canonical-32b2247fcf507ad424d4baea2333e49fa1ff14d840cb241caacc7c8a4a62b7f6)
- service.advertise_options.advertise_in_cluster.multi_ports.ports.info

<a id="canonical-a6988a68abeaa1fcd75e0bf5de3e980044a0db3c903bdea812f4e9250ef51cd4"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Upstream description:

Port information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port"),
  validators.ConflictingObjectAttributes("same_as_port",
    "target_port")}
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
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

Terraform syntax:

```terraform
info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1483b557979259d3bddc8e2c2ac9edd805316745a1ba3a2032aac0f57b9271d4"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / ed27bb44aea9 / 3

<a id="canonical-d9461525519ca61725f3c5c91037759a7bc2afa97d49a125c119b92df3e761fd"></a>

<a id="canonical-68326688253d7a575094dbfd3eb2992b6178bf3988f0abcef7e791236581ac01"></a>

## port property — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / ed27bb44aea9 / 4

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-e333f09dc75bda1b4f4aad08949ce0b5b6db60fc6af553788425b094408e3ec9"></a>

<a id="canonical-22ace116fa790b26529089caf8f954852f1a9f2539929697f63abac76569264a"></a>

## protocol property — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / ed27bb44aea9 / 5

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](resources--workload--reference--group-008.md#canonical-e32a1082ec858c96a1369fd69d0f88ee37873305912d508dba15c0442d1dde1d): complete subsection reference.

<a id="canonical-4e88c1e522134c817bfcef743b60690d869291c2c57a90fb1457f0d76276300a"></a>

<a id="canonical-a9490672be55ae097ca0fda89df8269b1494d702ae2854e03cb62f8742bb74ba"></a>

## target_port property — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / ed27bb44aea9 / 6

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-7e4e117f7952a1d51bd5abb788a41edb81e69a3df200b2132b72e098f8d28f81"></a>

## Next pages — service.advertise_options.advertise_in_cluster.multi_ports.ports.info / ed27bb44aea9 / 7

- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port](resources--workload--reference--group-008.md#canonical-e32a1082ec858c96a1369fd69d0f88ee37873305912d508dba15c0442d1dde1d)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-008.md#canonical-32b2247fcf507ad424d4baea2333e49fa1ff14d840cb241caacc7c8a4a62b7f6)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e32a1082ec858c96a1369fd69d0f88ee37873305912d508dba15c0442d1dde1d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1db292170eac24902ebc1469048de759f33374ffc456fddfd5472c9fc4d931ad"></a>

## service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port — service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_po / 8c9848847103 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967)
- [service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-008.md#canonical-ef48cef304055faaa9f30a0ebabc2c3a4c3f9fed60c014f27d132eb2ad5d356e)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-008.md#canonical-32b2247fcf507ad424d4baea2333e49fa1ff14d840cb241caacc7c8a4a62b7f6)
- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](resources--workload--reference--group-008.md#canonical-4ef34860d148a65122d5905c748ea00557e5a73d17af81cc461b8ee59584a87d)
- service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port

<a id="canonical-50418fae8dd2dd9ccf746f4bbdb05c89ef3ee8b104dcddde3f192ca4cf55bf6b"></a>

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
same_as_port = {}
```

<a id="canonical-7cc59c1e85386c43a1b6ddde4115c8629a2fa058d81feb0aa3ef84143bbef1ae"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_po / 8c9848847103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-12001600d951787e523b03af8a5bf549ca374e343b809ed44c43abdaabc2e4fb"></a>

## Next pages — service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_po / 8c9848847103 / 4

- [service.advertise_options.advertise_in_cluster.multi_ports.ports.info](resources--workload--reference--group-008.md#canonical-4ef34860d148a65122d5905c748ea00557e5a73d17af81cc461b8ee59584a87d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8e459cc553aaf03c23bb5044ba63b480e23fea5bb45eb90e1daa1a4f961f99dc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3a460c8f33ae2e453d6dbc5a1c8b914f83946b614204178a3cd9d87860b0313"></a>

## service.advertise_options.advertise_in_cluster.port — service.advertise_options.advertise_in_cluster.port / f5bb8f95b199 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967)
- service.advertise_options.advertise_in_cluster.port

<a id="canonical-9147b6a55b5262ef87acf9b3542b0f80e2ff148c5d32607db51d7d76c31b52b9"></a>

Type: `"object"`. single nested block, Optional.

Port. Single port.

Upstream description:

Single port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-94b5bb29351b9dc074a82cd7cec4cba783f00094d97b0de2e2ffda4510867475"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.port / f5bb8f95b199 / 3

- [info](resources--workload--reference--group-008.md#canonical-c02597f0161cf6873e65cf78b8293773e5788e6912247799d9906e774e59d81f): complete subsection reference.

<a id="canonical-1146b37d8631e5b0882f01304dbc6c11b60b2d1fd7a3cecb896a9ec2e0b861ea"></a>

## Next pages — service.advertise_options.advertise_in_cluster.port / f5bb8f95b199 / 4

- [service.advertise_options.advertise_in_cluster.port.info](resources--workload--reference--group-008.md#canonical-c02597f0161cf6873e65cf78b8293773e5788e6912247799d9906e774e59d81f)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c02597f0161cf6873e65cf78b8293773e5788e6912247799d9906e774e59d81f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee184cf366579de05dcdef6ed969d894d1b99b6571a9ca1e3a495d64b9411344"></a>

## service.advertise_options.advertise_in_cluster.port.info — service.advertise_options.advertise_in_cluster.port.info / 3267e3827c70 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967)
- [service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-008.md#canonical-8e459cc553aaf03c23bb5044ba63b480e23fea5bb45eb90e1daa1a4f961f99dc)
- service.advertise_options.advertise_in_cluster.port.info

<a id="canonical-d2e720a277f3122fe856809040454560a46bd125bcbca0d2e9fe9d5edadebac4"></a>

Type: `"object"`. single nested block, Optional.

Port Information. Port information.

Upstream description:

Port information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port"),
  validators.ConflictingObjectAttributes("same_as_port",
    "target_port")}
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
  "x-ves-oneof-field-target_port_choice": "[\"same_as_port\",\"target_port\"]"
}
```

Terraform syntax:

```terraform
info {
  # Configure direct properties listed below.
}
```

<a id="canonical-e545ae9f718cef76218c5a3a3dc2923b8fc892db528d284d35c872e34e04a366"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.port.info / 3267e3827c70 / 3

<a id="canonical-231dc719a59dde6414b123c0e20d5834ed3edd71a1a1844ea53a87919fb0c279"></a>

<a id="canonical-44c274fb2f4ba8fa25de9518c99430f112bc0dee673e8872804de65a92949e6d"></a>

## port property — service.advertise_options.advertise_in_cluster.port.info / 3267e3827c70 / 4

Type: `"number"`. Optional.

Port. Port the workload can be reached on.

Upstream description:

Port the workload can be reached on.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-b3427c3208ee54920aa90b29e7125fdd1a86ccd8834242d340f04df1f3fc3695"></a>

<a id="canonical-e8a0321643fad40ac9ce99375c292464e66c8bbf00c3d16974a0408a64a3978b"></a>

## protocol property — service.advertise_options.advertise_in_cluster.port.info / 3267e3827c70 / 5

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_HTTP|PROTOCOL\_HTTP2|PROTOCOL\_TLS\_WITH\_SNI|PROTOCOL\_UDP\] Type
of protocol - PROTOCOL\_TCP: TCP TCP - PROTOCOL\_HTTP: HTTP HTTP - PROTOCOL\_HTTP2: HTTP2 HTTP2 -
PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI TLS with SNI - PROTOCOL\_UDP: UDP UDP. Possible values are
\`PROTOCOL\_TCP\`, \`PROTOCOL\_HTTP\`, \`PROTOCOL\_HTTP2\`, \`PROTOCOL\_TLS\_WITH\_SNI\`,
\`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

TCP &#8203;- PROTOCOL\_HTTP: HTTP

HTTP &#8203;- PROTOCOL\_HTTP2: HTTP2

HTTP2 &#8203;- PROTOCOL\_TLS\_WITH\_SNI: TLS with SNI

TLS with SNI &#8203;- PROTOCOL\_UDP: UDP

UDP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_HTTP",
    "PROTOCOL_HTTP2",
    "PROTOCOL_TLS_WITH_SNI",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [same_as_port](resources--workload--reference--group-008.md#canonical-3d3028f180896c9359d037a15fca1b1fff49b2b9cba735d4e2f8498046a0152e): complete subsection reference.

<a id="canonical-b779fb1117ff06c6e192247aacb69585224358432fecf7688456aa6558d7b05a"></a>

<a id="canonical-2eb8b463034d2bf311b2e66d5d5c1ecaa7018b8222a4a175ca19ff9b30f62b6f"></a>

## target_port property — service.advertise_options.advertise_in_cluster.port.info / 3267e3827c70 / 6

Type: `"number"`. Optional.

Exclusive with \[same\_as\_port\] Port the workload is listening on.

Upstream description:

Exclusive with \[same\_as\_port\] Port the workload is listening on.

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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-49f37555896f06bd06509f50a475e70dc3daf2e59b5c2e602cd7354094ff084b"></a>

## Next pages — service.advertise_options.advertise_in_cluster.port.info / 3267e3827c70 / 7

- [service.advertise_options.advertise_in_cluster.port.info.same_as_port](resources--workload--reference--group-008.md#canonical-3d3028f180896c9359d037a15fca1b1fff49b2b9cba735d4e2f8498046a0152e)
- [service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-008.md#canonical-8e459cc553aaf03c23bb5044ba63b480e23fea5bb45eb90e1daa1a4f961f99dc)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3d3028f180896c9359d037a15fca1b1fff49b2b9cba735d4e2f8498046a0152e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3786ece481e07b8578d855e6d4ff787e1350bfe8924d7d4299c19ea2e7d0849"></a>

## service.advertise_options.advertise_in_cluster.port.info.same_as_port — service.advertise_options.advertise_in_cluster.port.info.same_as_port / f20cda67b34d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_in_cluster](resources--workload--reference--group-008.md#canonical-3c2f17491c942f70c268ee3196b5f46ff346531af030c2f06723830827076967)
- [service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-008.md#canonical-8e459cc553aaf03c23bb5044ba63b480e23fea5bb45eb90e1daa1a4f961f99dc)
- [service.advertise_options.advertise_in_cluster.port.info](resources--workload--reference--group-008.md#canonical-c02597f0161cf6873e65cf78b8293773e5788e6912247799d9906e774e59d81f)
- service.advertise_options.advertise_in_cluster.port.info.same_as_port

<a id="canonical-b4738e013973621052da91603d7976bb28bfb89fade42b63aeb7a4946ff7184a"></a>

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
same_as_port = {}
```

<a id="canonical-b42ab3849555958a7198cc96e3c55acd41128459af990bac4e06eeef19367063"></a>

## Direct properties — service.advertise_options.advertise_in_cluster.port.info.same_as_port / f20cda67b34d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f4ddbd2923fdb49e8a21a1e870fec2c256fc317fd1155933e72cbb9774468f2"></a>

## Next pages — service.advertise_options.advertise_in_cluster.port.info.same_as_port / f20cda67b34d / 4

- [service.advertise_options.advertise_in_cluster.port.info](resources--workload--reference--group-008.md#canonical-c02597f0161cf6873e65cf78b8293773e5788e6912247799d9906e774e59d81f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee450e50d6a249f5b0dd71d6ec36ee45f62c080db6d030764c9e2b0a6380da94"></a>

## service.advertise_options.advertise_on_public — service.advertise_options.advertise_on_public / c0da7b28e816 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- service.advertise_options.advertise_on_public

<a id="canonical-d1156a0056b32f553413e76830830d0a0cab936658caf5df1f392fc10ee4f9c8"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on Internet with default VIP.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("multi_ports",
    "port")}
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
  "x-ves-oneof-field-advertise_choice": "[\"multi_ports\",\"port\"]"
}
```

Terraform syntax:

```terraform
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-b20a2e133f67d377ab0735550c952b47841e5acde088e5fb7a1c3bce35086a02"></a>

## Direct properties — service.advertise_options.advertise_on_public / c0da7b28e816 / 3

- [multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123): complete subsection reference.

- [port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1): complete subsection reference.

<a id="canonical-4097de59005dffd6a0c5d8cc7a7e6c69d6d0c9f11db844c11e9d80fe87840540"></a>

## Next pages — service.advertise_options.advertise_on_public / c0da7b28e816 / 4

- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ab11f3b688c729bd8dba6ac4c082c17bf3608e3cc3c0baf2851ce44ab4219a46"></a>

## service.advertise_options.advertise_on_public.multi_ports — service.advertise_options.advertise_on_public.multi_ports / db953146ec37 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- service.advertise_options.advertise_on_public.multi_ports

<a id="canonical-4c0916167844536bbea76655ec8aa8747804185826a7f6f1079071d29c59cd34"></a>

Type: `"object"`. single nested block, Optional.

Advertise Multiple Ports. Advertise multiple ports.

Upstream description:

Advertise multiple ports.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ports")}
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
multi_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-cc85c9398f8a76c43a08c8b812d3ab03e14c18343bfb82caaa5d80757e717194"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports / db953146ec37 / 3

- [ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a): complete subsection reference.

<a id="canonical-40013a7fb786f21f7811a937735c7b825b63f2eeba67c12e1beb668d970f0916"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports / db953146ec37 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d3b23871274267a114213e3f6a763a68460536a8a558cbde71c23fa5c07deb1"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports — service.advertise_options.advertise_on_public.multi_ports.ports / 6c69514d77d3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- service.advertise_options.advertise_on_public.multi_ports.ports

<a id="canonical-662875c64830aa66d8a8922db3faeef631f4851580c9c827faa8ee4b723d4426"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
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
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-fb4f3c4344401f913b0aca4e08697e10e7c2b1813c3c2bb95d236997642cf807"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports / 6c69514d77d3 / 3

- [http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d): complete subsection reference.

- [port](resources--workload--reference--group-012.md#canonical-3df6a48ca312dd6a53e3d784178510d1f1840fec69978ddd52219d094c927f4c): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-012.md#canonical-6102c34eb63c02fd55faa6d0818d5cad1f48e01c996b8475abd485db68a82de8): complete subsection reference.

<a id="canonical-b0132a8eee35a3613213fdda60d9d55f3ebab6e0a2c40771b9310bfb64d9bf0d"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports / 6c69514d77d3 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-012.md#canonical-3df6a48ca312dd6a53e3d784178510d1f1840fec69978ddd52219d094c927f4c)
- [service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer](resources--workload--reference--group-012.md#canonical-6102c34eb63c02fd55faa6d0818d5cad1f48e01c996b8475abd485db68a82de8)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0f629952cde3647b7968bcc735fe9c74fd477eed03f3d7043227f3e19fa7100"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 467596e75cc4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer

<a id="canonical-26138cd7780b318cd0f4b093be11f59a0a937cb4682b0b9884ebae76f90b9278"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("default_route",
    "specific_routes"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

Terraform syntax:

```terraform
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-bfccb738f833903b623d6934a1e3313d7cc709b9d81111d64539c77a2f447f65"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 467596e75cc4 / 3

- [default_route](resources--workload--reference--group-008.md#canonical-2e1f554a03e0b95a79e4513c6b619a3f277302f0816d1efbc2790a66b5195f11): complete subsection reference.

<a id="canonical-8ecd0741ec07a103c38806bec50186c8911c74d61e0d2190fd855c1c68ff7f4c"></a>

<a id="canonical-f6f49c21edd20e4e908acf48de98af26478a5e421c3febb76f5aa38cb8424736"></a>

## domains property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 467596e75cc4 / 4

Type: `["list", "string"]`. Optional.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](resources--workload--reference--group-008.md#canonical-8e6f6c0837bd7b7a7693e3629f93c2341434ac4b66e7822d8fb630ecb4921768): complete subsection reference.

- [https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-010.md#canonical-1b8c5834bb7bf53d3894609abb1e85b28031e74b76e1cce7ab713bf0dc822832): complete subsection reference.

- [specific_routes](resources--workload--reference--group-011.md#canonical-ffbd01fdba16046eb807c00d809a0c34e45fc346864ea2783a5affb6fd62c4a0): complete subsection reference.

<a id="canonical-e0d369c7884b609400575de1304a39ad9e910ac0a05bf278f010145321e1ce88"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 467596e75cc4 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-008.md#canonical-2e1f554a03e0b95a79e4513c6b619a3f277302f0816d1efbc2790a66b5195f11)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http](resources--workload--reference--group-008.md#canonical-8e6f6c0837bd7b7a7693e3629f93c2341434ac4b66e7822d8fb630ecb4921768)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-010.md#canonical-1b8c5834bb7bf53d3894609abb1e85b28031e74b76e1cce7ab713bf0dc822832)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-ffbd01fdba16046eb807c00d809a0c34e45fc346864ea2783a5affb6fd62c4a0)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2e1f554a03e0b95a79e4513c6b619a3f277302f0816d1efbc2790a66b5195f11"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47ea642746293b66c6515ef18239f7cd85d32e28d2884ef2cab584554209ef63"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 428c38518320 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route

<a id="canonical-0e9e30127346bf4df12ba76515798321a1f1e7fbc72c892bd24e224772da099c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-e1ed88adbf7f6db10c5d409d0017bdecd091da458b01abf87d9b85cadc78d559"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 428c38518320 / 3

- [auto_host_rewrite](resources--workload--reference--group-008.md#canonical-ba63ea0c30133a9e2dbecd18f6726323fb2b1673701f1fbb74b96eb4f96a2529): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-008.md#canonical-65093dff6d67018376adc92602b3c6a1069cf32508e2c66f97757df8d4f5b37f): complete subsection reference.

<a id="canonical-7fb92bf1a026ee5166884d8e01975e16c49ec59b047052b5731d1ed5ecff8dfe"></a>

<a id="canonical-0236f90f2351b6a7c0dbe7bd03d700d41b3c949649966a31280704177a59b614"></a>

## host_rewrite property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 428c38518320 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-ca6b5cfe582917dc7109300f86618798cf4c8920abdb68aa4b55b1e06a872510"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 428c38518320 / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite](resources--workload--reference--group-008.md#canonical-ba63ea0c30133a9e2dbecd18f6726323fb2b1673701f1fbb74b96eb4f96a2529)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite](resources--workload--reference--group-008.md#canonical-65093dff6d67018376adc92602b3c6a1069cf32508e2c66f97757df8d4f5b37f)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ba63ea0c30133a9e2dbecd18f6726323fb2b1673701f1fbb74b96eb4f96a2529"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2892c86705ac9a64e0715ac16f7a2207262f268689e5e63125932504d3dc70ea"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d071fd1860fa / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-008.md#canonical-2e1f554a03e0b95a79e4513c6b619a3f277302f0816d1efbc2790a66b5195f11)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-89f725c682d98221614a8eb5fe5b9f0f8ed5b01dd3a4d95a95ae7f2f34b235c6"></a>

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
auto_host_rewrite = {}
```

<a id="canonical-80f326248a19b33fffd8410ef16eb87144d2ea4f8da5a383a40dac7c823ca0d1"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d071fd1860fa / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-435977c731c19e69dc90b3fd74052595041c2fec2409accfc78ddc600d371ec8"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d071fd1860fa / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-008.md#canonical-2e1f554a03e0b95a79e4513c6b619a3f277302f0816d1efbc2790a66b5195f11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-65093dff6d67018376adc92602b3c6a1069cf32508e2c66f97757df8d4f5b37f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-972c8b2bb974393a57de68516e251d191e2fc222839eb0136f3d03ab4a81579c"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e5aef3cdc848 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-008.md#canonical-2e1f554a03e0b95a79e4513c6b619a3f277302f0816d1efbc2790a66b5195f11)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-a61ef1c276e0843af2b6ae921455bdacd014d9e436c2080b8194558dba084714"></a>

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
disable_host_rewrite = {}
```

<a id="canonical-f61472bddc44109d91add6913f71ed268e1bd9604bbad63086dd5de75eb42fce"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e5aef3cdc848 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5378d94c958b8b18f953cc6f4ab1d64ced3c52d6735bacce4684880d5c83d22a"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e5aef3cdc848 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.default_route](resources--workload--reference--group-008.md#canonical-2e1f554a03e0b95a79e4513c6b619a3f277302f0816d1efbc2790a66b5195f11)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8e6f6c0837bd7b7a7693e3629f93c2341434ac4b66e7822d8fb630ecb4921768"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f65a2b51a8fc6d3bc5d42f064c039809ec2536a69fe35ced935ffe58874be60a"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e20fa5353520 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.http

<a id="canonical-be16f9b161cdcda392effb63399727570b1e9a5d281cd5f13e0ee39c9237adb1"></a>

Type: `"object"`. single nested block, Optional.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("port",
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
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-fa988ead4e236df1ce87ba06f2afb6efb429453e54207ab66c26aefd21487bfb"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e20fa5353520 / 3

<a id="canonical-6bd98ebe9543246bec5d6cbbe312e8988ce711a35fc9f6ab22df90e6602408c0"></a>

<a id="canonical-3d5258cf24f97d39c911f0e1f0f0fea465c8039f02e402cfe499c9798e606e38"></a>

## dns_volterra_managed property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e20fa5353520 / 4

Type: `"bool"`. Optional.

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

<a id="canonical-5dad3f47cdc21efcbc392c8ffa1917f8b927ca0f101eac1e24dee7a44648f5bb"></a>

<a id="canonical-f1a081718e866bc98ae1465b491e9f7102862eda5cc0c170e4f2e0f7a02e1215"></a>

## port property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e20fa5353520 / 5

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

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

<a id="canonical-83da0d4a4442d4be38288e72082c17143a2c6893a578fe5d9cd3acca54f99bb5"></a>

<a id="canonical-2e6e7aaa12ba195ec2a6c616328aab4e42fc3ab95e5bc568f67b22f36396f362"></a>

## port_ranges property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e20fa5353520 / 6

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

<a id="canonical-214c122356190ba6ecf739aa34485f0d05ce1ddd3b4ae95bef37c657a51b4f3c"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e20fa5353520 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60252c4b60a0bfd5c5f4fedd7aa358c93d56655e41726ee9e52dee8ea008cb0d"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0db260ab393c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https

<a id="canonical-dc0d63cfe89f19b27ad0762fa17bddbe2bc74f582fdec2a0e3842ec19d4a4692"></a>

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
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingObjectAttributes("tls_cert_params",
    "tls_parameters")}
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
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

Terraform syntax:

```terraform
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-cb34f2d2f701d38d8755514fdd09fe428d310fdddf07059cb53e81548ced3f60"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0db260ab393c / 3

<a id="canonical-90a45f154eed6109c087ac6c6a4dbc324cc6322e73274b9efc7173deeeaf4186"></a>

<a id="canonical-487ee9ad417a4936b21f6e4f23a9e63237adff99af1d160c01a379ab03a5a6df"></a>

## add_hsts property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0db260ab393c / 4

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

<a id="canonical-dade76a586822470e065f5a0b58763313a1b377d03bfa906994c41ba80d9595b"></a>

<a id="canonical-f92a32f7b928fd4611d5483acd736fbf5c8de7404f49e49dd39a56f9664eb198"></a>

## append_server_name property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0db260ab393c / 5

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

- [coalescing_options](resources--workload--reference--group-009.md#canonical-b0d6167f755948fdf49b0985b7b0671a7e8d7638c8ca714abdaeae3533af70c1): complete subsection reference.

<a id="canonical-9f6c4d1dd315493eed894fb4e46ecc15c799f59b80362c6c829c765c3d2df27b"></a>
