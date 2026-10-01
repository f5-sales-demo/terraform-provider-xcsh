---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-c54d15c6b9731519610848967cdae87c82fea8fdb1e632354497a0ef296402b0"></a>

## connection_idle_timeout property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0db260ab393c / 6

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

- [default_header](resources--workload--reference--group-009.md#canonical-01a20d81575e883599b8238bf5d93080191118bc959eb902b9f3a568bdfd7067): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-009.md#canonical-0c77b7ceec77d27918f7b57bb58b7d9b94aed98ff8f79bab12dc68f3b54b7a38): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-009.md#canonical-10ec871e51a400cd49707e8c0d7220e6bea395710f2a8c4d4f826a0cf2752759): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-009.md#canonical-6f3da11f37bc38a8dbd127320c120eb172a7efdbf0b9b56c2848e33cd19210b9): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05): complete subsection reference.

<a id="canonical-76ab91be2a3fb3db671b8fb31841a63bb2b62bba614a2fbe4711c2be8057e533"></a>

<a id="canonical-f74bd6637e76cb0c89cd468d71f49bd770a74118e6c2c3f55812a2c8405057c5"></a>

## http_redirect property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0db260ab393c / 7

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

- [non_default_loadbalancer](resources--workload--reference--group-009.md#canonical-cfc975f94039cf9a1f16743bfb61e59c6a42d874ec89badc9339cc1147672467): complete subsection reference.

- [pass_through](resources--workload--reference--group-009.md#canonical-23f59b76141f8ddb6fcc971bfda98c9f7eb589c8e330dd7dbae982dc02323fee): complete subsection reference.

<a id="canonical-4a379c53cccb1d03848da97d0ded1f356630cba277a3982dceb62371198242b1"></a>

<a id="canonical-cf316f857528d9503f17a87ef98ac0b29a06ea18deb88bcee7477c0316ab87d3"></a>

## port property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0db260ab393c / 8

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

<a id="canonical-38e678e1325ac9eb94ca4b59f12dc3dac7f59b63a75bd90b74c17135a492ad11"></a>

<a id="canonical-0b9e508161e8def13e0949c4a92235e07327fdf140d110dd85a69a9fb1cdddab"></a>

## port_ranges property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0db260ab393c / 9

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

<a id="canonical-01257e9ee6d0d6bbd50a00b088d35f4733c0dbd0a37b04734d6711a9587a020b"></a>

<a id="canonical-b883ba1f97ca493800e103417aba4e71b5a1515d617ebf950a49018b229bedfe"></a>

## server_name property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0db260ab393c / 10

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

- [tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-009.md#canonical-ec7bb5761e2146e32ac9162d221b91b367d08a00118e551d73e3cc8fc1de9096): complete subsection reference.

<a id="canonical-59c92da2401a76c072e2083b70aa402a54ea6ddb17bee9e130a7b9b7fe86aed0"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 0db260ab393c / 11

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-009.md#canonical-b0d6167f755948fdf49b0985b7b0671a7e8d7638c8ca714abdaeae3533af70c1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header](resources--workload--reference--group-009.md#canonical-01a20d81575e883599b8238bf5d93080191118bc959eb902b9f3a568bdfd7067)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer](resources--workload--reference--group-009.md#canonical-0c77b7ceec77d27918f7b57bb58b7d9b94aed98ff8f79bab12dc68f3b54b7a38)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize](resources--workload--reference--group-009.md#canonical-10ec871e51a400cd49707e8c0d7220e6bea395710f2a8c4d4f826a0cf2752759)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize](resources--workload--reference--group-009.md#canonical-6f3da11f37bc38a8dbd127320c120eb172a7efdbf0b9b56c2848e33cd19210b9)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer](resources--workload--reference--group-009.md#canonical-cfc975f94039cf9a1f16743bfb61e59c6a42d874ec89badc9339cc1147672467)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through](resources--workload--reference--group-009.md#canonical-23f59b76141f8ddb6fcc971bfda98c9f7eb589c8e330dd7dbae982dc02323fee)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-ec7bb5761e2146e32ac9162d221b91b367d08a00118e551d73e3cc8fc1de9096)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b0d6167f755948fdf49b0985b7b0671a7e8d7638c8ca714abdaeae3533af70c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b139e4911a74f6f3e724ae9f8a1a94ccdc2c3334b5bdf04b2dc64e028a7b75ad"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c6a6fdb767e1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-3131c6e11b51aa535b70fa2eccefe961ac61a28ffcea6ae5ef356e250fe19d6b"></a>

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

<a id="canonical-c2a7e67a904390b30b33cd427fbf3051516c2607d93b032d118ba0bb53e1025e"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c6a6fdb767e1 / 3

- [default_coalescing](resources--workload--reference--group-009.md#canonical-7c7846b79d3bae7f17ae27b4ed3c2958220c3b48aa90793d59e9c0ddf5b4fb9e): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-009.md#canonical-d144151be036f3fad1dc5ea098f0fbfd42bdda8139f7929e5ab69f14344cf98d): complete subsection reference.

<a id="canonical-e9c56b23679f0bcdcd7f142076df79e9f8bc4f47f75a02e5aee192cbd66abefc"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c6a6fdb767e1 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing](resources--workload--reference--group-009.md#canonical-7c7846b79d3bae7f17ae27b4ed3c2958220c3b48aa90793d59e9c0ddf5b4fb9e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](resources--workload--reference--group-009.md#canonical-d144151be036f3fad1dc5ea098f0fbfd42bdda8139f7929e5ab69f14344cf98d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7c7846b79d3bae7f17ae27b4ed3c2958220c3b48aa90793d59e9c0ddf5b4fb9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a2a582b4ec977ee79b3dac4726a9c71d1e3681573377b962b692f8043fa085f"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 81f7bda93736 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-009.md#canonical-b0d6167f755948fdf49b0985b7b0671a7e8d7638c8ca714abdaeae3533af70c1)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-91c40decf1881524db21eedba82dbc5092e36f445e9b6f33a2e48fb2a581a4b1"></a>

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

<a id="canonical-ace060b9eaed0d29f4bf00e9a2d02db7b0c32e0caf69968927fd2bdad5abdd68"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 81f7bda93736 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c5efd419f47cd04caed025f98600c0be1128f2df90115c6db612332aa4c3daca"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 81f7bda93736 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-009.md#canonical-b0d6167f755948fdf49b0985b7b0671a7e8d7638c8ca714abdaeae3533af70c1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d144151be036f3fad1dc5ea098f0fbfd42bdda8139f7929e5ab69f14344cf98d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e5bc8c26fb8a02f3078b566ac47e3635dc38c4414abe9b3a4acb65dc12a763c"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / a7870865cb1e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-009.md#canonical-b0d6167f755948fdf49b0985b7b0671a7e8d7638c8ca714abdaeae3533af70c1)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-be5817d5b208c7e2af12f4aa784b9709bd5c1be1f20dcbc8769c14efae9f3785"></a>

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

<a id="canonical-0da04f36125869f7909994f5eae8e71b29007aa918c45d583a41ca43f2983130"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / a7870865cb1e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f89398556d258548fde3fa709462d5389409e0b6c51d9abda3eb760b50643ba"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / a7870865cb1e / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-009.md#canonical-b0d6167f755948fdf49b0985b7b0671a7e8d7638c8ca714abdaeae3533af70c1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-01a20d81575e883599b8238bf5d93080191118bc959eb902b9f3a568bdfd7067"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ce9657cbed33820b24f611f1689f382e445c3e7c11f675a309db477c63a7f526"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 836c46fca851 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_header

<a id="canonical-d2f13910619cae070ad1d87b310f3f0282f4d07fceb3a761831ccb27d77ab9b7"></a>

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

<a id="canonical-6b0aab8795a1d7aff55fd0fe509e1ac697d889a14479d5c966749d64156e301f"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 836c46fca851 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8722989504b27b5ddae5bcd9b411dd1c73761fd832745cc171b38ede3e7c3d75"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 836c46fca851 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0c77b7ceec77d27918f7b57bb58b7d9b94aed98ff8f79bab12dc68f3b54b7a38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b92857d86619d386ab7bbb4067da00e643f2f190752a605e4ef5bd6705db4673"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 718d6bc0a440 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-44ccbe3c08e1d9995d6518515f6aa4b28c6c132a6e65f240b5cea1080be446c6"></a>

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

<a id="canonical-2eec4806f50b7c56da58231a7a67a43a56426923d3311ad2d062584b1f720bc7"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 718d6bc0a440 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-826735e00ceb9eb8f1fde0103f4c507c2e801eef5808a0cd3def349364475c12"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 718d6bc0a440 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-10ec871e51a400cd49707e8c0d7220e6bea395710f2a8c4d4f826a0cf2752759"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ffa64a457f35e250396c15887213352814653a39f0df0b732db3dfd6914276ce"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / da7dece29763 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-232adeea56a203bdf026b92b39d94edc6f16359ad034fe02a4276aa05c4c7004"></a>

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

<a id="canonical-ef7bcaa349a984e007e93f9a9892a1f910bf94ded6524e888f2325ea3a23190d"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / da7dece29763 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-45e12631b2539b38cef87db15f6562d02701532ae50da6f24ad718a9519deccf"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / da7dece29763 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6f3da11f37bc38a8dbd127320c120eb172a7efdbf0b9b56c2848e33cd19210b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ed23d5210879f6e73c8287ceb97fea767e7cc4b3cbc1e610a4556099e5eaded"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 877dd9e72e10 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-0b2c42274bf4c14a6abfab75d5b3b78ad7291c3737c74679c2d05682214ddd6c"></a>

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

<a id="canonical-ca7e0423671703eb590e47a21a9104d637fb0023456806b6a33482476fada564"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 877dd9e72e10 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b0e81d6effd52090416ff99f6cbf57cfc87e08799780ec615bb1f4e9b8a48d9c"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 877dd9e72e10 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f76702898530e00e8cf079b5adb20424e3481096fa6f4bf3195c0448d37316d5"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4be247fab037 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-812e599ae9559488be7e96f12e0e993c8fce7054942de63d412e91e714cf91d6"></a>

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

<a id="canonical-8975b2770a6d914c287f196a5162d7d70dc6af0128f45380193b661f17160130"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4be247fab037 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0609e2c1d2fa3f358de2efb7df3f86cb359c0e5c668bfbfa79a122a2a0afea1e): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-009.md#canonical-867c841717f1508d92c2625463d1118344e5fbeb0aa77d25da2958bd7563d56e): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-009.md#canonical-becb52e491ef8ff625721b7f1c48e00349517340788f83c9c5c4ab63c51f5311): complete subsection reference.

<a id="canonical-9ac498f0e45c2fc84e97295ae8a4ba4db35df1b963ed5a0c01c2eb475702aef7"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4be247fab037 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0609e2c1d2fa3f358de2efb7df3f86cb359c0e5c668bfbfa79a122a2a0afea1e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-009.md#canonical-867c841717f1508d92c2625463d1118344e5fbeb0aa77d25da2958bd7563d56e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-009.md#canonical-becb52e491ef8ff625721b7f1c48e00349517340788f83c9c5c4ab63c51f5311)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0609e2c1d2fa3f358de2efb7df3f86cb359c0e5c668bfbfa79a122a2a0afea1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1970e3a9bc66ef3c0760f38af873cff4960ff4bccd916dfd064ca628c6061f8"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 345b55bec8a6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-75f4eca5438a861e9f271c5e62e7113c37da64df061670c848b5d9a046284115"></a>

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

<a id="canonical-c16fb409ba3a9b83316553f27c1afb0197ddc869d9f1faeaa4be880171bf78c4"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 345b55bec8a6 / 3

- [header_transformation](resources--workload--reference--group-009.md#canonical-9d123b8544ad38c9a0ebd53040a59741cc1e6046ad96a75b8e1183869aa2cff1): complete subsection reference.

<a id="canonical-51a5462562d7956dc02cc9979149f6413ae01c8e6bf46e11b9d368694f84c472"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 345b55bec8a6 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-9d123b8544ad38c9a0ebd53040a59741cc1e6046ad96a75b8e1183869aa2cff1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9d123b8544ad38c9a0ebd53040a59741cc1e6046ad96a75b8e1183869aa2cff1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc41974973cb94d818e47da380e57702a4ac6c573b427507406f2686b07027ae"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / edfaa3e6026c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0609e2c1d2fa3f358de2efb7df3f86cb359c0e5c668bfbfa79a122a2a0afea1e)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-390b1572ccb4663c93df10d5163b5371e76e8c020ec6d2ae3c3f0b58a68611e9"></a>

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

<a id="canonical-2e6f1b1cfe4dd4262acc47f7dae69f72995c8e3cc6e1b2ddb5ecae1910f8e017"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / edfaa3e6026c / 3

- [default_header_transformation](resources--workload--reference--group-009.md#canonical-8f3db567614944291f5a23976a0273370fbe29e38dba5ba13d1fe245f9e94869): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-009.md#canonical-d4c88f6f44ee516129f43cf7514f69072d0c9f80c4d8f626389d39b3922cd0fe): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-009.md#canonical-46bd612c592e5efc979edd718662e42bafd8af972c3d5008bb787a8710a8499f): complete subsection reference.

<a id="canonical-67a1f2000c36245e51420922d2f664583860bab17d5a3ee503ddf9cc29b0eada"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / edfaa3e6026c / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-009.md#canonical-8f3db567614944291f5a23976a0273370fbe29e38dba5ba13d1fe245f9e94869)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-009.md#canonical-d4c88f6f44ee516129f43cf7514f69072d0c9f80c4d8f626389d39b3922cd0fe)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-009.md#canonical-46bd612c592e5efc979edd718662e42bafd8af972c3d5008bb787a8710a8499f)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0609e2c1d2fa3f358de2efb7df3f86cb359c0e5c668bfbfa79a122a2a0afea1e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8f3db567614944291f5a23976a0273370fbe29e38dba5ba13d1fe245f9e94869"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-de4d9ce65daa5dc253ea6ebfb84ebf57523423ff15b263ca6f201fc7aec22d1e"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 06338d65c6e3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0609e2c1d2fa3f358de2efb7df3f86cb359c0e5c668bfbfa79a122a2a0afea1e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-9d123b8544ad38c9a0ebd53040a59741cc1e6046ad96a75b8e1183869aa2cff1)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-57d3abb17b5139a6cdeeba9611e5805a95e3d728b7d37f226141db65d7d29ea5"></a>

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

<a id="canonical-d627a3d08718ccd9ba388bdaebb6832b0ca94fd3913ff3692d5d164296c8cf47"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 06338d65c6e3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-01129f79c3f0ef06643f572649a82643fb2913f1af5bd749875811b8f610c0eb"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 06338d65c6e3 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-9d123b8544ad38c9a0ebd53040a59741cc1e6046ad96a75b8e1183869aa2cff1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d4c88f6f44ee516129f43cf7514f69072d0c9f80c4d8f626389d39b3922cd0fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d6291c0d8112b73233829aca15cb648fdc6c37b5e7ecde3e41deea4d0ced231a"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 224f12701509 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0609e2c1d2fa3f358de2efb7df3f86cb359c0e5c668bfbfa79a122a2a0afea1e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-9d123b8544ad38c9a0ebd53040a59741cc1e6046ad96a75b8e1183869aa2cff1)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-7bd8a4208c6c20b900b6db9ee7daaf926398bc69867cb9f9a3ff99cb0b9297e0"></a>

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

<a id="canonical-dd3fee6a2a63c6648c25b1db4875700f2915192ea7994fcd63c34a4f3951ca85"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 224f12701509 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cd129492b576d602ee20dc3a2f9000cbe34d82f3bf7774ded493091e00914b1f"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 224f12701509 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-9d123b8544ad38c9a0ebd53040a59741cc1e6046ad96a75b8e1183869aa2cff1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-46bd612c592e5efc979edd718662e42bafd8af972c3d5008bb787a8710a8499f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7b4de1d18de87cd15734ffda8365c2efccc8a688c7be5fd3656408f3ed9c7df9"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 99153bfbeb55 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-009.md#canonical-0609e2c1d2fa3f358de2efb7df3f86cb359c0e5c668bfbfa79a122a2a0afea1e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-9d123b8544ad38c9a0ebd53040a59741cc1e6046ad96a75b8e1183869aa2cff1)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-40154fe74e00839483ad544ddbfba986815064bde175364abcead1eeec16d68b"></a>

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

<a id="canonical-b9d1f88eaadf6a2f60c735be23798e4a33d16f561e0da206fe2c08a630673084"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 99153bfbeb55 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1a2f336fe330919676930d541b29a31be2c474e9d5e0a1e931fa1b01acb4d323"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 99153bfbeb55 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-009.md#canonical-9d123b8544ad38c9a0ebd53040a59741cc1e6046ad96a75b8e1183869aa2cff1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-867c841717f1508d92c2625463d1118344e5fbeb0aa77d25da2958bd7563d56e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af0155572cec6f7d7a156bafac344a5547033fe66a279a3e81eea5b5e5e3a225"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 90d5cace1c19 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-14e0dccded6149f95a72cbd7cf67f3abfa629ef13ec94a1a297d71a77d46f211"></a>

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

<a id="canonical-ddfb77b02a4546e709695b5f6a4cf301df86e4d321a299288028b2ecf009bbab"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 90d5cace1c19 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9f5a61f9f80d6b0afe723564cef88259e62b3d06817701a5fd5bf4df4e3249c4"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 90d5cace1c19 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-becb52e491ef8ff625721b7f1c48e00349517340788f83c9c5c4ab63c51f5311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23b66d88361ebbdc9e51c4e1098f34cea58a38226a711092625943a123822b42"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 87a4a377b6cb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-c15546fdee3b750f01d5c57c3a7fa33a275ba33d1bb12430b26ae639eff9555f"></a>

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

<a id="canonical-727ec829e3b60a9ae1f4c5706d71dfb8f07bd6eeef122fc4876a8181b761c6e9"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 87a4a377b6cb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5af044268fdbec6ecfd44893e86a9a6f3354c7ba591698892aee9140c564966"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 87a4a377b6cb / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-009.md#canonical-af5219ccff7e6408856559981797a33802d183dcfa8dcd05c662174ad8eb1c05)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cfc975f94039cf9a1f16743bfb61e59c6a42d874ec89badc9339cc1147672467"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ad4b426dee8fda3d80a083b0d408f71793d9cd5c70dbe081d23b94983abde67"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 7a1d62d26ee0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-5009c200d041f1caf4316ea13feadd4566b07758fa5fb7c2f4b802dd0aff8a78"></a>

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

<a id="canonical-ab717382064a93328faa30abbc8ba84c3365366d788e0200a73161992ef0ee40"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 7a1d62d26ee0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b25ec1ad5c3d67c7173599eb2b669c568b1d54795881ddca8b166ae981302b40"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 7a1d62d26ee0 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-23f59b76141f8ddb6fcc971bfda98c9f7eb589c8e330dd7dbae982dc02323fee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26acacb9520eb71404cac8ed04659746430451e444ae6557b07ffdb90a7bbb4f"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d2c058c307af / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.pass_through

<a id="canonical-f66598b7c9483c110d0d8eebe09b8a6122b8d2daa903215f9be4680796e33ff3"></a>

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

<a id="canonical-9c48b2e55c5014e0c1998ad3b418d3e52b584bad2d60943d7dda52792b3c65fc"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d2c058c307af / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8959448002a1cf79482fa6c17179e72c7ea8f496bb235b6c7e08e762580db100"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d2c058c307af / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-801129e0b9a22bc512228b27a615689418f6980b2324b97eb656cdb7648dce9f"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / fb05a3cef457 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params

<a id="canonical-23b686cb1d9381d084cb45423e8dc61568a75f23e8a6e51eb437a74484d90fb2"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificates"),
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
tls_cert_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-c33b8fff7fabae9a2a36f0b6461ac7a2ec40c128e7c02aadbc75a684e3a1c2e5"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / fb05a3cef457 / 3

- [certificates](resources--workload--reference--group-009.md#canonical-30ccb84de6c5bf1cec3d15d9a86faa3da35dce69e7f0c0596cb0e930d53aba6f): complete subsection reference.

- [no_mtls](resources--workload--reference--group-009.md#canonical-b94abb9fda4da337c0daa7693c02c42b56db75d94abb19ba54d8e4bfd4036442): complete subsection reference.

- [tls_config](resources--workload--reference--group-009.md#canonical-2fbb5238b5bcd059b9996387b514574dfd8151a4a43c03b3ad9a7986c94df46d): complete subsection reference.

- [use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a): complete subsection reference.

<a id="canonical-1c43ead94410489fea4ccb0ac12c932c0ebe1d5e294e0945a24a8f3b76049fad"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / fb05a3cef457 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates](resources--workload--reference--group-009.md#canonical-30ccb84de6c5bf1cec3d15d9a86faa3da35dce69e7f0c0596cb0e930d53aba6f)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls](resources--workload--reference--group-009.md#canonical-b94abb9fda4da337c0daa7693c02c42b56db75d94abb19ba54d8e4bfd4036442)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-2fbb5238b5bcd059b9996387b514574dfd8151a4a43c03b3ad9a7986c94df46d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-30ccb84de6c5bf1cec3d15d9a86faa3da35dce69e7f0c0596cb0e930d53aba6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f5972d8b4e7d7043e4c5d7737336365494d692b426344dbd4f0f5deb2bf2c0d"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 093655631e93 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-290f65f6c207b7fa5f690b203f733128ce6945a64f426a7373b9b1d6edc58f2e"></a>

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

<a id="canonical-4aa276a08b51df133fb57bb98102abd042940879e20207118acb50ae308c04d1"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 093655631e93 / 3

<a id="canonical-8e8c4a7f0b7ab0ba9b3245ead98a165de5370e42a5295f48771e856ccdf27d1b"></a>

<a id="canonical-9b1c9b3e23c10d9152cc45befa5410c6bb7d7b9f11c71a495e944ac0c7af12c8"></a>

## name property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 093655631e93 / 4

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

<a id="canonical-55bfc98eb1faa3d67bc05fcbdb7a3e00822211530f19fc9f50be2d891b5e33f2"></a>

<a id="canonical-6b387fd97fbe30907fb609c24b513f5791abb19325a4f6be972094f8d8f143c4"></a>

## namespace property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 093655631e93 / 5

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

<a id="canonical-bacf642dcf64f14f247308db346e043a509775e22d75cb09a45b251b35bc3fd2"></a>

<a id="canonical-a8c6f4b14ac910b2ad91c68bbdc2ffcc27426c4d72e6f966d895348d2875b732"></a>

## tenant property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 093655631e93 / 6

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

<a id="canonical-11a8bc12bab2a3d7bb2d6c9a8e25423e667f25a614b3a207984babede407d125"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 093655631e93 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b94abb9fda4da337c0daa7693c02c42b56db75d94abb19ba54d8e4bfd4036442"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e36edab07836db3240c649a3e06d14cf152969c71e56b36f94cef08ce5bfbe1a"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bbec29dd64ed / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-915f5299c69dd3179d353d4f8e78cc94f7385f3aaa11f2c083531d2556a58f24"></a>

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

<a id="canonical-7df6eef1d16978881ea07dd314baeae73a8321e9921d1ce46884c8b90bbc3874"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bbec29dd64ed / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1718feb557c3b65545c4d65a81e44651a7c05ae02f97e025c70c41dcf5ea6242"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / bbec29dd64ed / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2fbb5238b5bcd059b9996387b514574dfd8151a4a43c03b3ad9a7986c94df46d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cf1c17a30adecb026d2b5df38ad865a370939f799f5b88a625c283d8121988ed"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 566e7c00789c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config

<a id="canonical-fc91205a1931e5511319bb615ed490e2b6114a48adefe3641cb0935c961eb9d7"></a>

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

<a id="canonical-a735eb168ed33d7d9384a30e7a866e2dbfb2ecfc1b1b40508544bf8650bd08cb"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 566e7c00789c / 3

- [custom_security](resources--workload--reference--group-009.md#canonical-c81c7ac094607335a67f940f8aa2b17bc4e1323e7a971ce054dd98218117ed94): complete subsection reference.

- [default_security](resources--workload--reference--group-009.md#canonical-cc6738a04723b8b912f94e959a5b78bbd50509828fa67c4db2485af898c35382): complete subsection reference.

- [low_security](resources--workload--reference--group-009.md#canonical-3caa16d1673dbd7dba250153ebee18cd609e102a1f1048ff084e0fec7ccbadf8): complete subsection reference.

- [medium_security](resources--workload--reference--group-009.md#canonical-296841e351d96b535a8471a2976f90eed050de95642ce2fcc3b92d071fb2af3c): complete subsection reference.

<a id="canonical-6f99d7de87427d91084f4db6300f826fa31c0cc924fa3f12b3bc6e668e5c3a31"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 566e7c00789c / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security](resources--workload--reference--group-009.md#canonical-c81c7ac094607335a67f940f8aa2b17bc4e1323e7a971ce054dd98218117ed94)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security](resources--workload--reference--group-009.md#canonical-cc6738a04723b8b912f94e959a5b78bbd50509828fa67c4db2485af898c35382)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security](resources--workload--reference--group-009.md#canonical-3caa16d1673dbd7dba250153ebee18cd609e102a1f1048ff084e0fec7ccbadf8)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security](resources--workload--reference--group-009.md#canonical-296841e351d96b535a8471a2976f90eed050de95642ce2fcc3b92d071fb2af3c)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c81c7ac094607335a67f940f8aa2b17bc4e1323e7a971ce054dd98218117ed94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-054d2554dda2ab9957e2adb05b5a1cfda34e2dc5ba9039cb332d20a88ea10583"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d3748c913b37 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-2fbb5238b5bcd059b9996387b514574dfd8151a4a43c03b3ad9a7986c94df46d)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.custom_security

<a id="canonical-4998829093cca3be6f4f32173f3f1924d8bd71c3b83a994a61460d2fdb138525"></a>

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

<a id="canonical-758b57aaf10287f5c0ce86c67fa9ca0ed9b93f152b00e196b5ba3a141ffcbf5d"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d3748c913b37 / 3

<a id="canonical-b1291ba9b8727cd7bbdc317aa44b746b382231cb4c56d8ab04ba51ba4db7625c"></a>

<a id="canonical-1e5b18ad975a46b5b7920eef72a8a0ebed223a6836e902effbd26b4902308cb0"></a>

## cipher_suites property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d3748c913b37 / 4

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

<a id="canonical-2f279e482074b2b86f3ca0a24a39f72d25344708164377713a9f356b142958ca"></a>

<a id="canonical-6a3d2dc379ddf77bfe8dc387e344accd94420427b696cc834d279765d9a2a2ab"></a>

## max_version property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d3748c913b37 / 5

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

<a id="canonical-56fd6040ea692bb04dadf4d2e45118ce139ab3fc98a8851a9c5bab4abf6ac352"></a>

<a id="canonical-3a79afc2e657318d769a763666b022ad043b16909bb59079cbd365b2872b9630"></a>

## min_version property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d3748c913b37 / 6

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

<a id="canonical-7522be83aab4be3f9a88f7fe4b20c57c60256b0a698db72657bd7e9f5bc657f8"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d3748c913b37 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-2fbb5238b5bcd059b9996387b514574dfd8151a4a43c03b3ad9a7986c94df46d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cc6738a04723b8b912f94e959a5b78bbd50509828fa67c4db2485af898c35382"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1df32e1ab03d33d09b6d59c3d1c1d0edae8cb5da617d5c9f28c5a5b58cf156a"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e0a67edfb28b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-2fbb5238b5bcd059b9996387b514574dfd8151a4a43c03b3ad9a7986c94df46d)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.default_security

<a id="canonical-f2e017f0e0b51bef18a938cf51a78c6cd61ec11861974de3cc368b624206eed0"></a>

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

<a id="canonical-011790f327f525584a6398d2f066b15f9dcfe9df503f65d624ceb75c26c5131a"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e0a67edfb28b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d14ebf504a56e63480bba98ebffc77ebac9907d07e6ed767b8b1e5cdfa58051e"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e0a67edfb28b / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-2fbb5238b5bcd059b9996387b514574dfd8151a4a43c03b3ad9a7986c94df46d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3caa16d1673dbd7dba250153ebee18cd609e102a1f1048ff084e0fec7ccbadf8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-297d6dea634ad8c68fc465a892145ccc1c4d4b28092c5faeda9c9dec4dc8dea5"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d205d59c74e0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-2fbb5238b5bcd059b9996387b514574dfd8151a4a43c03b3ad9a7986c94df46d)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.low_security

<a id="canonical-4fe95d36d8a86b9726970dc0f60ce6db2f436fe3362a2243bea29e36a641c0c0"></a>

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

<a id="canonical-15690c6bb4c045234d965cb533500eeafa938a09fe6155c1aa67ee6fe505e87d"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d205d59c74e0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1ea89983e9552951b9aa6a0a9f3e08dc21b3e54a8207bd370b56f2d234da979a"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / d205d59c74e0 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-2fbb5238b5bcd059b9996387b514574dfd8151a4a43c03b3ad9a7986c94df46d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-296841e351d96b535a8471a2976f90eed050de95642ce2fcc3b92d071fb2af3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f376eec8711dd0575abbcb758a75be624ad5efa6b6bcc7b072ffc0b462651dd7"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b24f56199a1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-2fbb5238b5bcd059b9996387b514574dfd8151a4a43c03b3ad9a7986c94df46d)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config.medium_security

<a id="canonical-e9c39bf9876e36771cc13f7d44d172c30ea617fb2827aca9d89fcff0b6abe2d0"></a>

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

<a id="canonical-cb89065a202b71819325b052ac77875f88cb68d25a844e2d512de19c75272f4b"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b24f56199a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b4e5a07dde8e77f09935d3537f9fef3e5f96f88ee945e5d951ccf45e51ccc0c8"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 2b24f56199a1 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-009.md#canonical-2fbb5238b5bcd059b9996387b514574dfd8151a4a43c03b3ad9a7986c94df46d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a669586e618ba58503702b0b90c24cee441300f9fafb5fb4e40ca5f247e95639"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 7f679cd7cf93 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls

<a id="canonical-c8d2bdde6874f158f84dc0b126a7cabdce31586d337aa8420206dbe624b2672c"></a>

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

<a id="canonical-829d4703225f4a1655b254041a391ebdaef2be20826079195773eb6535d8c6df"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 7f679cd7cf93 / 3

<a id="canonical-dbaed55091b554904c946453521bb03bb632f23e630b9d4560765cbd40a7d5cd"></a>

<a id="canonical-eb5aedbe766e377d4b7bf94e9fd69b784418d96b4419013c6eb367f7e469a286"></a>

## client_certificate_optional property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 7f679cd7cf93 / 4

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

- [crl](resources--workload--reference--group-009.md#canonical-78030271434995d1e5f087b826c21aaea07f5f807c7d485c23ec24cf92dce753): complete subsection reference.

- [no_crl](resources--workload--reference--group-009.md#canonical-b6483438cc82a2832ce8cdd91e9131b8d8e64edc6bd8df8db104b587da7570ab): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-009.md#canonical-df52af02b32b4c8727d543bd6808f23482453c403ef48e2361dcd1772d2318a5): complete subsection reference.

<a id="canonical-28f21b335367edd02e4b9e85dcb87a4ec58aba37c6f4c78b0c6598739da71b83"></a>

<a id="canonical-615ba3e979662d12995d2e98ad905ef83f6d91a1078bb498d56efdc523a8aefe"></a>

## trusted_ca_url property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 7f679cd7cf93 / 5

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

- [xfcc_disabled](resources--workload--reference--group-009.md#canonical-24091da3cfe03182c9abb1ee2bb08884008dd7992dbb674a41054e188fc6fcde): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-009.md#canonical-fd9d39640afc765a046eb58319775758d88e0e6c0e49f2ab8dacff68eb2a58e1): complete subsection reference.

<a id="canonical-05092c74a91d372ea6d83ac79bb9afefb8f0aad78174ac4482e77c5820509da8"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 7f679cd7cf93 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl](resources--workload--reference--group-009.md#canonical-78030271434995d1e5f087b826c21aaea07f5f807c7d485c23ec24cf92dce753)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl](resources--workload--reference--group-009.md#canonical-b6483438cc82a2832ce8cdd91e9131b8d8e64edc6bd8df8db104b587da7570ab)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca](resources--workload--reference--group-009.md#canonical-df52af02b32b4c8727d543bd6808f23482453c403ef48e2361dcd1772d2318a5)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled](resources--workload--reference--group-009.md#canonical-24091da3cfe03182c9abb1ee2bb08884008dd7992dbb674a41054e188fc6fcde)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options](resources--workload--reference--group-009.md#canonical-fd9d39640afc765a046eb58319775758d88e0e6c0e49f2ab8dacff68eb2a58e1)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-78030271434995d1e5f087b826c21aaea07f5f807c7d485c23ec24cf92dce753"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3078d388b53eff489450c171e3cd688c26c7c11408b4653143e67446173729c"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 260d6af1d9fb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.crl

<a id="canonical-47b83d6aeb968882c968df50090c86fe1cab42ab4aea2aa0a77f376b9fcef93a"></a>

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

<a id="canonical-b75e78bc2752a0f3571e24946e07a2a777768f14fb109d6b8a178656eb21e47e"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 260d6af1d9fb / 3

<a id="canonical-9890ae639ba34dbca03cc3a265dbf28e1258b5734869969adcb16ff31a6ae3bd"></a>

<a id="canonical-9aa12c6df6e4e6da241b7211a93315dbaad704d0db72b2bacb191d91ebc8662c"></a>

## name property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 260d6af1d9fb / 4

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

<a id="canonical-414aeb1f687f573c5538923a81baaeae364714f9faf60ec4b0f169d02e045ba1"></a>

<a id="canonical-82204259ef9c15333a2fa2373bdbb9d45626941d8fbf885decd737066cd4e64b"></a>

## namespace property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 260d6af1d9fb / 5

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

<a id="canonical-7a27141ab61440395d7acb419a80db30777d7f5a02151c6e7b7fe925da46e76f"></a>

<a id="canonical-6a2521b2722c21fc48504ea806b42971a5c4f0027b8fd6b3a4c5e34e0921f274"></a>

## tenant property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 260d6af1d9fb / 6

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

<a id="canonical-62508d0b54cb55ffde7c4bac38f1bd832f25ed6c65eb44c237b76b9e543c4e8e"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 260d6af1d9fb / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b6483438cc82a2832ce8cdd91e9131b8d8e64edc6bd8df8db104b587da7570ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a9936c2bc3e6d566e7da4f9127307af9fb2a84add7f26afe7ee3397e9014d47"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e5e753191e2c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-3b0021186249cc369d3fc9bd9f36f17f8af547507809d0ba24b8ad4440881793"></a>

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

<a id="canonical-8054b4f872915db18c708b45bc93864020415949aebfe3ae8c22d9ee933e7a2f"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e5e753191e2c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6f85c0d8f38d163f8aa167b4e9c874c7dee580aa2af7a5f54caa7257a17acdb4"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / e5e753191e2c / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-df52af02b32b4c8727d543bd6808f23482453c403ef48e2361dcd1772d2318a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a04bbeb5b9d2509014d883e54cb63dfa09b2a3290ae3041255063b1e3f0826c4"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4d2c8ba0ec95 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-5e9de88a80e13e78eaaa4ff890876daa846fff816520328e4ce14c02939cf15d"></a>

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

<a id="canonical-8a9a0e4cc196caeb9ee69a00afe3acb767e1b125327986332a8146a64db9c875"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4d2c8ba0ec95 / 3

<a id="canonical-826562fb5a9a65630e6589e1dcd2706ebe82a9524e560effa92ec63fe099081c"></a>

<a id="canonical-8f3dbd1f4c1888fbdc9de6296748709182162b89a83ec94db8f8bc28686755b7"></a>

## name property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4d2c8ba0ec95 / 4

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

<a id="canonical-e64452495cfb7c6f1a74f0141a43b3db1df55fb8e37f2375d1d1bbbcf46d8d9d"></a>

<a id="canonical-d5ed01d20e77f4240b2d6e36facd6826223bfa6b82524d3a1e85105d8d3f7684"></a>

## namespace property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4d2c8ba0ec95 / 5

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

<a id="canonical-ee6444544e85f16e5a144be66897d47d66b712ded497a96b36da7b96b134ad0b"></a>

<a id="canonical-a26d0b5e1618611b9d0908b84ff24deaeed370e5cf6b0cc48b9d6a7f3d326609"></a>

## tenant property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4d2c8ba0ec95 / 6

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

<a id="canonical-81db3461edad3f8f8f597f115fefd534b27a12ac049d5637451f9dc6907a297b"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 4d2c8ba0ec95 / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-24091da3cfe03182c9abb1ee2bb08884008dd7992dbb674a41054e188fc6fcde"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8184e95ac83033d7e8d0a0f288715cafc868e4b65cf3cc9bf9370b38c5158b6b"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c571cae13b0b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-1c294c551224921a07f2561accb09d3107d78a415d8cdbd1d07852d2a83fef09"></a>

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

<a id="canonical-d93c905210600c07a1ce4e80b21bde8de5a5a600185184c3a7b356b622e47f85"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c571cae13b0b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a1290468b5b8268b270f4703c95632dc169a2d26c8f1a201e1272d0826a911ba"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c571cae13b0b / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-fd9d39640afc765a046eb58319775758d88e0e6c0e49f2ab8dacff68eb2a58e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-98db1c1415abe226c960806850663c966bc81ef03dceb4492f7390a92c3741c2"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 726b711740bf / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-009.md#canonical-266977ff404b82e9833a8bf809b55876c13461bef7a730f61617e4015833228e)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-8f86cc668310c4059c69609612a4946446300c2b095509c63266b37c35fd0b5a"></a>

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

<a id="canonical-f5fbd275e224ebc8a42766f3ccf28b7ee30b540abaed0b07f765b37d107aaa6a"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 726b711740bf / 3

<a id="canonical-3d74fb674dd5eb50a7643daaafaa04ca295409faeed9c68a857c3566bfc3e92f"></a>

<a id="canonical-d46f88d7b831006b24e351883d817610fc6ce8e39683ce40b3d998f41932fe1d"></a>

## xfcc_header_elements property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 726b711740bf / 4

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

<a id="canonical-30b1be9157d29118938b7e21c4099768c01ad82d52688217017a50e1cba45a07"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 726b711740bf / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-009.md#canonical-acedf53d91cb76f18fc7f12ef86d090068014f3df114898f542b4e9cde84633a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ec7bb5761e2146e32ac9162d221b91b367d08a00118e551d73e3cc8fc1de9096"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a0680bd2f35117c9b22d0719566f2a9a818951abc6a73e2107db3ffd5f41ca1"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c4a2bf2d31c1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters

<a id="canonical-2cb496c5639e7f332ebf745c48ae3f452a5b172bccbb0fc4f14902276863557e"></a>

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

<a id="canonical-6040375f5b2defbaaff30b6040d0e810fd64e2ae4f648422310fe5ae85ca109f"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c4a2bf2d31c1 / 3

- [no_mtls](resources--workload--reference--group-009.md#canonical-50d17245b0656aecfab560b6441b8bd1dd3d46d604b693d2869d987eb9762f03): complete subsection reference.

- [tls_certificates](resources--workload--reference--group-009.md#canonical-28dc5817a4a1c6b048ff5ced29ab2a06d81d27e89ec818acba7df73777e27da9): complete subsection reference.

- [tls_config](resources--workload--reference--group-010.md#canonical-6f8f2966f060f89120f01db96a7cb36b6fdc26e5e39ba6c65d3560c97db932e9): complete subsection reference.

- [use_mtls](resources--workload--reference--group-010.md#canonical-f4a220d808bb40756c179c8dc28282284248dcef3af8dc560dfcbbd31f4a5d61): complete subsection reference.

<a id="canonical-7f8ef6b57c213ac51a69d3e524a9df4985b93d139bacf9f4f8599d9c0d257278"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / c4a2bf2d31c1 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls](resources--workload--reference--group-009.md#canonical-50d17245b0656aecfab560b6441b8bd1dd3d46d604b693d2869d987eb9762f03)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates](resources--workload--reference--group-009.md#canonical-28dc5817a4a1c6b048ff5ced29ab2a06d81d27e89ec818acba7df73777e27da9)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_config](resources--workload--reference--group-010.md#canonical-6f8f2966f060f89120f01db96a7cb36b6fdc26e5e39ba6c65d3560c97db932e9)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-010.md#canonical-f4a220d808bb40756c179c8dc28282284248dcef3af8dc560dfcbbd31f4a5d61)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-50d17245b0656aecfab560b6441b8bd1dd3d46d604b693d2869d987eb9762f03"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cc811388e9c957a1c4400bba89c927c7ac7673cb1ecb076915f837023b6b1e12"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 6c3c5952bbcb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-ec7bb5761e2146e32ac9162d221b91b367d08a00118e551d73e3cc8fc1de9096)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.no_mtls

<a id="canonical-200aa5cf22f55a115a68adf50653f57c0bc23e5e8737437f3dcba92a8c8568b1"></a>

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

<a id="canonical-7defab50479a740f4b26b5b1f8782bc68f1d0d28434beea2a69916452d40f984"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 6c3c5952bbcb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c8ba83b9b7e3f351a1bb0db347f48a0b1e49535646654abd693866066685cc12"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 6c3c5952bbcb / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-ec7bb5761e2146e32ac9162d221b91b367d08a00118e551d73e3cc8fc1de9096)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-28dc5817a4a1c6b048ff5ced29ab2a06d81d27e89ec818acba7df73777e27da9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b456f5c74638cf0e8136fe31189abca826b6451996b1f6a704819090d9f5b73e"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 1b821c1b5330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https](resources--workload--reference--group-008.md#canonical-e6d1e42780c6860c16eecd28e039ffa4ebd24e38bd1f4a848449f73a956f4e6b)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-009.md#canonical-ec7bb5761e2146e32ac9162d221b91b367d08a00118e551d73e3cc8fc1de9096)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.https.tls_parameters.tls_certificates

<a id="canonical-6ec143865ea777fdc8a5ca816226ca944c13ca63c61f9b8c483222a58ac4d5be"></a>

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

<a id="canonical-2d906ddbf7e64e1eddf2bb0e5a96dee3ba3eb377899e5926119da00ebb5bf0a5"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 1b821c1b5330 / 3

<a id="canonical-34ca50ff30a3c6d3f102da9c71036f8f7efcaea8972237cc25511cace85664ca"></a>

<a id="canonical-c8827753c116554577ab41d28bcd3de7318ee02b431b03f3b2b0dd4344b79851"></a>

## certificate_url property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 1b821c1b5330 / 4

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

- [custom_hash_algorithms](resources--workload--reference--group-010.md#canonical-50fac80a78c5b95ff1bc6840a6d88f1377cf7c7a8528c185c932081c990bff8b): complete subsection reference.

<a id="canonical-e1daf67576abcf5aed9a3eead0d73d6d3cddd143003590b4cba02d3a65222931"></a>
