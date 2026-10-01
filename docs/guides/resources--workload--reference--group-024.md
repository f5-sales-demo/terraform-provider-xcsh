---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-0879dd76c6e0969d16cc037a34e07cfe5e5000d8d08af6f0048cccb9ef4e470c"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 30df9c5b2326 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-023.md#canonical-92034da7aec0493b2329b828ebb181fdc0fe9c674d7aa9cf673f21d43a503eb4)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-3e9090dfa3e4f8ca652e4c4e2d89e201d7df0c3253353f083fd9cd6fd8e9ca38"></a>

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

<a id="canonical-0a61a95dec56852ee22dfa02655138d146f9d7f14e1a9fa1fbb8e2c1bf1666c3"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 30df9c5b2326 / 3

<a id="canonical-2db4212e0b80fe7aaf2d608279927926ec8700d13f3769b06ecbdd2744ed8470"></a>

<a id="canonical-a865c09f2ba656842c6bddd8304d71ec51a4419594800458a2854977368b4a0f"></a>

## path property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 30df9c5b2326 / 4

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

<a id="canonical-b49166eece06906baf59f7dae6ad468c562f17ce0d371b2dd671c15e67ed657d"></a>

<a id="canonical-69395eb1a20f84d440bd2f541744aea0a311b73edf865f4a71e203d2610251da"></a>

## prefix property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 30df9c5b2326 / 5

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

<a id="canonical-65f7b8c979fa026715028cc0b34ec309c140ff4b822051922e7ef2d529c7f07d"></a>

<a id="canonical-b144ac1e9888eb93a995aa8ad99b777300bed9c9268f39410f450cb18970305b"></a>

## regex property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 30df9c5b2326 / 6

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

<a id="canonical-4d3852fff23bd434ebcb50b5749b25faae984b6c6305ce30b26564f9029afda0"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 30df9c5b2326 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-023.md#canonical-92034da7aec0493b2329b828ebb181fdc0fe9c674d7aa9cf673f21d43a503eb4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0ab18006251d7af847343bde7a67a19ab5c7cfc96a4160a143fd713bc5c20f3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9af2b0b5678ae5517db6be8c04a4341b5c4e6ef3d6f21bf834cf2e9b5c70c47"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 822c6e901f23 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-023.md#canonical-92034da7aec0493b2329b828ebb181fdc0fe9c674d7aa9cf673f21d43a503eb4)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-0c8c3ea942882bfd5a7f1021a2cae33f1cf829b0ef9fffc47eeec4e667a4f621"></a>

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

<a id="canonical-19cb23d94d58a2dc6cb4d5691111ce849e271ba9fee71be01aba43defa5432a0"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 822c6e901f23 / 3

<a id="canonical-5d3eb8d7bc78a8b952b40333127c183fbac3916e6b0160e92c776a39901583a8"></a>

<a id="canonical-cbc6ee16038a71ab6b59ec70aa444ddeca136ef2afce8ace82573305c3c48a2e"></a>

## response_body_encoded property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 822c6e901f23 / 4

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

<a id="canonical-c68191b1542907e1548f5ad189b1576373f605e90cea25c7df6c1043d579636c"></a>

<a id="canonical-3da8c744700df5b1949c6dc3e8b99b9c1378b2fc8526d53108cd969b6b1ee140"></a>

## response_code property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 822c6e901f23 / 5

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

<a id="canonical-a341947b9a455112bb4bf3e7f61556a267408226b5916111ac7c5f8cb767da2f"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 822c6e901f23 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-023.md#canonical-92034da7aec0493b2329b828ebb181fdc0fe9c674d7aa9cf673f21d43a503eb4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-647e93f7b8c81fd3024334d860f918fa89c269f0571a2098e7663dd5a918474a"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3e814514f3b0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-99b4ad6120b29ce942d1f079d7f4a5013c84bd862582ee5653c7ff559575d47b"></a>

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

<a id="canonical-58fec32ee96736690ad37c04d16bfb9ec768a4acb52c085d3514eadf23a85c80"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3e814514f3b0 / 3

- [headers](resources--workload--reference--group-024.md#canonical-76faae1bb98ecd2fcd66dd940efaec2d14a273ca897fa9f0dc38623e9e13eb5f): complete subsection reference.

<a id="canonical-7b0474c40a308b90b2ca5e0215e16f993e0a4edb27a3c79d791d2fc15d2fb8bf"></a>

<a id="canonical-259a9dc45ce589539a9f5a879f027355620766ec085a856cde7f09a6f15edeeb"></a>

## http_method property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3e814514f3b0 / 4

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

- [incoming_port](resources--workload--reference--group-024.md#canonical-116e075b11ab5568c8069558fdb2fadca1421982c6e9225e21098fe78224a095): complete subsection reference.

- [path](resources--workload--reference--group-024.md#canonical-db24076af932372e142668b49860aa51d033481a6e8568d73934a0457bf11d70): complete subsection reference.

- [route_redirect](resources--workload--reference--group-024.md#canonical-64d5269db315fd944cae60c948b2a6423a008496a9f22462b1d1be57de30933b): complete subsection reference.

<a id="canonical-08af457bbb8154320baca6dd602b84eeffc21825b54890449b4f7c78762dc22d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 3e814514f3b0 / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers](resources--workload--reference--group-024.md#canonical-76faae1bb98ecd2fcd66dd940efaec2d14a273ca897fa9f0dc38623e9e13eb5f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-024.md#canonical-116e075b11ab5568c8069558fdb2fadca1421982c6e9225e21098fe78224a095)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path](resources--workload--reference--group-024.md#canonical-db24076af932372e142668b49860aa51d033481a6e8568d73934a0457bf11d70)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-024.md#canonical-64d5269db315fd944cae60c948b2a6423a008496a9f22462b1d1be57de30933b)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-76faae1bb98ecd2fcd66dd940efaec2d14a273ca897fa9f0dc38623e9e13eb5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0545c20eeb229a3eaf8dda45a553052560d706ae2c3a7adbd2d40b21f89f57b4"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e8f07c71ba5 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-024.md#canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-7c3fbe1505d711e62f925cb39d03a6767fd15c17393547d6268bfd8cffb65f44"></a>

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

<a id="canonical-9e00388e7e591ff7d78be124da6901510875abecbd6bef5bbe310a4973d58b41"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e8f07c71ba5 / 3

<a id="canonical-f72682f189eed13bd74d64f2068bc0785e6736ce7d0d69961ab846a660699ff0"></a>

<a id="canonical-9e169ea58f9b62ff362d2a6f897059e949bae7ed1c3feced305ef4db1042029b"></a>

## exact property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e8f07c71ba5 / 4

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

<a id="canonical-2a90d5e30e87813557c5d8588678f04572bae3a1950ded87cf3994211c83f5cb"></a>

<a id="canonical-4a3196f6ca291cc5484341761f0af7f87e7bf65f61331509e04015166e6070fb"></a>

## invert_match property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e8f07c71ba5 / 5

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

<a id="canonical-c365b7e412ae6f8d07d2f093e2e2b35d77f59f27870ff84832aa8f75094e872d"></a>

<a id="canonical-d304c642693166780f9670514b02b4e5507a1aed64c8a9bb41eacb17324bac1f"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e8f07c71ba5 / 6

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

<a id="canonical-8a884416090abf27c2da0e1b7e3ecd16035bbc942ff2f2cfaf4323d009bd86b2"></a>

<a id="canonical-63d691583f39e53450a66f2b6f9e35d3d03f3bba5af545d80772febb4cda7b20"></a>

## presence property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e8f07c71ba5 / 7

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

<a id="canonical-441572e4b59cbe3df70af773a9a192d8bfc1cb06d431d33cb020b31f924f571c"></a>

<a id="canonical-30f1e3e77e342550ecdeadc09bec790de8e4947fdc3550ace41db04a885502a7"></a>

## regex property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e8f07c71ba5 / 8

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

<a id="canonical-8d2cee20b742e4417d46c94e329e4ddd079ab75ec3b12c067f1ba2cd20a85e86"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e8f07c71ba5 / 9

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-024.md#canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-116e075b11ab5568c8069558fdb2fadca1421982c6e9225e21098fe78224a095"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee2dacf2cdc9939a5acee6d468cef12cf2ef6fdee17dc9ba17f443df1358ce14"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6c89abb93c19 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-024.md#canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-45ea56ea97b4e55a53298eed37bd7ccb93b940dc97102ac060821bcf398aa205"></a>

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

<a id="canonical-42a5bb66260be0dbbaef5df388df8ca6e0862fc5d2737820e2e34409833ccbc8"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6c89abb93c19 / 3

- [no_port_match](resources--workload--reference--group-024.md#canonical-59b2ad93e283aa31795cc7f5956fa83a2904ca2289fd4e22f5f68ee966a5c793): complete subsection reference.

<a id="canonical-8085f6d59f15e595665eb70a8a14bdd32d030572e126ec3c03f7cd3e2d8cc518"></a>

<a id="canonical-2fda54c5e6e4b2ef8d27022de8971f29519d1a6c3ce83c214012d587e51ed062"></a>

## port property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6c89abb93c19 / 4

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

<a id="canonical-f968cc5718036c4ac27c0b390fd782d690d2090250d70af4ba033fb1c5e4f104"></a>

<a id="canonical-8d02c4e8c89cfd03702fd7e30a349b298f46f14533f76e909e0939ef4f57dfe7"></a>

## port_ranges property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6c89abb93c19 / 5

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

<a id="canonical-290f74a2de5adebabbcb65f231aca3096965c7847cab175b4928fa7619abc74c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6c89abb93c19 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](resources--workload--reference--group-024.md#canonical-59b2ad93e283aa31795cc7f5956fa83a2904ca2289fd4e22f5f68ee966a5c793)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-024.md#canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-59b2ad93e283aa31795cc7f5956fa83a2904ca2289fd4e22f5f68ee966a5c793"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25b0b3cd327f2a6a19296c57de411889d9acf25d703e6ffc5e6b18a138b2c035"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1ac2e9f886c0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-024.md#canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-024.md#canonical-116e075b11ab5568c8069558fdb2fadca1421982c6e9225e21098fe78224a095)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-c86541a6a35d6295cf201d86db1f3d25500e893a151d826947f0a8b31f5d3784"></a>

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

<a id="canonical-284003f55abfce94e83f56905bb285ffa59ba14f7a98698488af0a5eb95fcb38"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1ac2e9f886c0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-6ab25c25860c0a4b42a8a011345da71b90c2b37aad673cc905e3b1ae93aa0107"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 1ac2e9f886c0 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-024.md#canonical-116e075b11ab5568c8069558fdb2fadca1421982c6e9225e21098fe78224a095)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-db24076af932372e142668b49860aa51d033481a6e8568d73934a0457bf11d70"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4273ad4a23ff0b4e3b9ebc409962f0b40281b0aeba760caf27909869cbf229a6"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f89fe134ab2c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-024.md#canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-70fb3c3401ca64ce94c85b0c19f54b119a2d7ff53710b8edc8a6e38dc6839389"></a>

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

<a id="canonical-023a3c6e231e1fbb5c8d443344ca04ffc896cb1f3ab67bc83b133eb69668d99a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f89fe134ab2c / 3

<a id="canonical-af7609e93e79ca83c6eea2546d435dce53ccd012a228e235bf0078b8034cff1e"></a>

<a id="canonical-11ff86a93b907ee14ba871317e065d154f9b1b787331903c27437906a14be3fa"></a>

## path property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f89fe134ab2c / 4

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

<a id="canonical-4e2b961bf615ab33c5e754cb36b2592333dd012645d9388b999fd09df86389f7"></a>

<a id="canonical-025d8565b7ef623c724b97e39a3c016cc708e6f221cb71d3a94a36dc24df4562"></a>

## prefix property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f89fe134ab2c / 5

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

<a id="canonical-1f155167227f7b48ad6065526aa9fdfa90b42ccd42025e236f8880d78ed5f632"></a>

<a id="canonical-1e09f0572e5e754fd545b101497176b2e88f8873e2ae0bc16fb10814de46c8ed"></a>

## regex property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f89fe134ab2c / 6

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

<a id="canonical-88e5b095bbe883f3144184fc03187255ecb96aef0fcd9e85c832234cf52626dd"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / f89fe134ab2c / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-024.md#canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-64d5269db315fd944cae60c948b2a6423a008496a9f22462b1d1be57de30933b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6be54727de2d236cc91d33c933d5f6da810f9731c9bc1df83cc00fb2e8f0515"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 230a465fe725 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-024.md#canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-9473cd2e28d4344cce3933138388d5f9b80e8c60e9fc023329c30502ad3bf38c"></a>

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

<a id="canonical-8a5cede19465b79512218824a11c63c8fc3b36bedc1bc167fd079c97bf5d6049"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 230a465fe725 / 3

<a id="canonical-909d8eb4b86dc9ad5370f41e9ce0168f9247e7c8be5016260811e31ef0388f38"></a>

<a id="canonical-9dc781201803e87d8c1934fc60b7aa7d7464da53f64a4e2834ca72e17f6deb56"></a>

## host_redirect property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 230a465fe725 / 4

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

<a id="canonical-194c6e5aad093fefc1b7674e4647d0d3533fc4a3e9158ef43d9fc81120c3d8e3"></a>

<a id="canonical-9c4082fdf6e9d703c8f060405105277bf2ded96bc74e7ff44570278d78292bf1"></a>

## path_redirect property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 230a465fe725 / 5

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

<a id="canonical-5d8879c2490e79c004ba001ee3e3307c9616073c7a92d4e653fe499136e29a56"></a>

<a id="canonical-c9f3f471659eeab084c84e0d73e4a91198aafcff940f5ae80f3e18edd4937da8"></a>

## prefix_rewrite property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 230a465fe725 / 6

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

<a id="canonical-e0f35ca97355c88565c435a86b4a1a84c1440901be05bc1b309647da3b3c64d8"></a>

<a id="canonical-0d23f07f0528bb3490f5aaa95c13d09ef8d2ec8d4b59019668d96d3297369b14"></a>

## proto_redirect property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 230a465fe725 / 7

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

- [remove_all_params](resources--workload--reference--group-024.md#canonical-df5bda3660ea0d14938cd0e2912cd13af84389c79f43aa4273f21c92a6313755): complete subsection reference.

<a id="canonical-6b9b32984990d69895cb5ec0ccd5c7b53b514af07519a96708317d0ec66c9b43"></a>

<a id="canonical-9d70d899e93b49d8193158008e9680c7fb3f68a4dad9879e4413e9f1d12b7922"></a>

## replace_params property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 230a465fe725 / 8

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

<a id="canonical-1ae013e67ca65c1f017c5341e9cb7924470e3cff4f62e1bdc813ba0a56740b1a"></a>

<a id="canonical-b349cb66beb40c5fe558f40d94be44a39f3055d22ac1913529ce7dda3e4cedca"></a>

## response_code property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 230a465fe725 / 9

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

- [retain_all_params](resources--workload--reference--group-024.md#canonical-a852eb9a7eabfe1ea524ca756f48bbeecc8e530f36290704a43ead5ebfb5b55b): complete subsection reference.

<a id="canonical-0b1c9dae98007f931ee28872ea1e2051c685ed83b894ad8088bd33ca39a77ac5"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 230a465fe725 / 10

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](resources--workload--reference--group-024.md#canonical-df5bda3660ea0d14938cd0e2912cd13af84389c79f43aa4273f21c92a6313755)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](resources--workload--reference--group-024.md#canonical-a852eb9a7eabfe1ea524ca756f48bbeecc8e530f36290704a43ead5ebfb5b55b)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-024.md#canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-df5bda3660ea0d14938cd0e2912cd13af84389c79f43aa4273f21c92a6313755"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f522d08e8f5b16c41cc2aca79ed9d7594b83da18482283169a8726e34d14d856"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6ca14e16ed7c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-024.md#canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-024.md#canonical-64d5269db315fd944cae60c948b2a6423a008496a9f22462b1d1be57de30933b)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-00ae8ff91b3dc2126a5e92d5b89b8a0afb71fe7602c76a860064f13a119701a5"></a>

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

<a id="canonical-34039df312516878aef323faff33fca1f0dc25bd022ff66e340b44badf7e40bf"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6ca14e16ed7c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5fedf074bef595530ab665c9bef1cb93d28cf09850a7af60483d7b4133daaebd"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 6ca14e16ed7c / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-024.md#canonical-64d5269db315fd944cae60c948b2a6423a008496a9f22462b1d1be57de30933b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a852eb9a7eabfe1ea524ca756f48bbeecc8e530f36290704a43ead5ebfb5b55b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-534dfb9bf53180e6bb671dc7b6f4f1a248a3b67694bdaa5e6ee633cfdbc15fda"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4072e5249f04 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-024.md#canonical-057d73cd39b234ff81a249908d8e4fa4bfa2bf458391aa010a1edc9f5a344b5f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-024.md#canonical-64d5269db315fd944cae60c948b2a6423a008496a9f22462b1d1be57de30933b)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-a0bbeee5bdee6f96a2a426eb007e05d33f010cbfc33667e834d3455a41ff6b22"></a>

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

<a id="canonical-413073d8f146983d74a2a940eef8870824e6541a8b14b2aa13be7faacddb1a86"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4072e5249f04 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5d1ef41c7673ebd141459bb963e0f98f06c5a7c80851d67693b68c57a21b48e"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 4072e5249f04 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-024.md#canonical-64d5269db315fd944cae60c948b2a6423a008496a9f22462b1d1be57de30933b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8214128d24286124b7c24a2f6af302f3126fd4142aa27b61c6436ba399761ba4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e32bf14b076e96d5e5c419454e9f397d02f3db8a0d03e1bf0d831d10dc96e81"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e3d2dc6e114 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-f6c6d006048139a7835f0fe0b58fd9d041170742a0353ccd85245ed2edc4d4f6"></a>

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

<a id="canonical-814b0adcd67ffb8c31832837c852597e3ac8d247181757215457c9704a7fae52"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e3d2dc6e114 / 3

- [auto_host_rewrite](resources--workload--reference--group-024.md#canonical-075c8d7c6a44829bd02292de3e5892a350f8838021ecbdfc38ef81966d0cd872): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-024.md#canonical-7d194d0a7e0484e5d8989bf29dcef3496c70b23b160ced67df662eee81e3a568): complete subsection reference.

<a id="canonical-1d56ebb5db59874aede414ff63c241fb6129182d8669967dfe159240187c5fb8"></a>

<a id="canonical-3a584176c988ab793f52e1de2d15dfafedb7e202087f69eb7a57bd71535fab74"></a>

## host_rewrite property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e3d2dc6e114 / 4

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

<a id="canonical-6ead28403a21e360e051c56d50c776848de1409e76da077e5fc5efde6400c16d"></a>

<a id="canonical-b6b95156f2f54a5bc31b275d01e5d6ecfefd7d0228fea0acd2ae1526b61ac3f2"></a>

## http_method property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e3d2dc6e114 / 5

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

- [path](resources--workload--reference--group-024.md#canonical-19c7a0299f0a53dbfc72b3270ddb1e683df376b7b0c4cff8ee89361ed5cd8935): complete subsection reference.

<a id="canonical-8ef6245c839b2309896be11ae0baed0d03df4ce0ff5ade9e0750bd4c3a2548df"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 9e3d2dc6e114 / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](resources--workload--reference--group-024.md#canonical-075c8d7c6a44829bd02292de3e5892a350f8838021ecbdfc38ef81966d0cd872)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](resources--workload--reference--group-024.md#canonical-7d194d0a7e0484e5d8989bf29dcef3496c70b23b160ced67df662eee81e3a568)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path](resources--workload--reference--group-024.md#canonical-19c7a0299f0a53dbfc72b3270ddb1e683df376b7b0c4cff8ee89361ed5cd8935)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-075c8d7c6a44829bd02292de3e5892a350f8838021ecbdfc38ef81966d0cd872"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3ff8b114e203a976b9f4e3932b8b93daa40482cfd8d25246a25f393369259ef6"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 506dc853acf2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-024.md#canonical-8214128d24286124b7c24a2f6af302f3126fd4142aa27b61c6436ba399761ba4)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-6ea46b4226df06abbdde2cb8a5991841eaaa0bd6025fc0bd0a4e5c8e33ed7de9"></a>

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

<a id="canonical-374cdfa8962b8f1305ec8a0628c5897821b696b24c75c5f468a8e9afc96b3b89"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 506dc853acf2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fac8ffa9f9883611eded7854d03be376c9f445ef2c08057c94bb38024e4ec751"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 506dc853acf2 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-024.md#canonical-8214128d24286124b7c24a2f6af302f3126fd4142aa27b61c6436ba399761ba4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7d194d0a7e0484e5d8989bf29dcef3496c70b23b160ced67df662eee81e3a568"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b40888a50128b68a20d8497e0360033b5db9f997ec41aed07fdb1c9eb2d08204"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 81d7ee12f6a9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-024.md#canonical-8214128d24286124b7c24a2f6af302f3126fd4142aa27b61c6436ba399761ba4)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-f12544ae161b1dd7afbf2b4b880a07c81a80dab31e810fbc4e7727fab0e0d84e"></a>

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

<a id="canonical-c1be672563daedfc17ddc1c669acfe13b3798d1f8ed5b2a84509bdb2158476e3"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 81d7ee12f6a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2bc991056426d3a2288715dd3c1127447441da879b251e269ecfdcdb2a06f21c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 81d7ee12f6a9 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-024.md#canonical-8214128d24286124b7c24a2f6af302f3126fd4142aa27b61c6436ba399761ba4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-19c7a0299f0a53dbfc72b3270ddb1e683df376b7b0c4cff8ee89361ed5cd8935"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c386d10059b826048641cfafa6f93ad7540a798468b5866b6db04ebe3baae65d"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c7c642c5d25 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-021.md#canonical-f9e4ca1e1b5612586b35235880ec65a64342ddd94e702a5ef1592d520a7ec687)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-023.md#canonical-dddaae7c754cd3d7873d183b5c3ca7fa93637c2535abeed5130c10e33c92ec33)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-023.md#canonical-c85535c9f5ee221117664b0a31a457f2cfd9fc067fe58ecf6b1c415cf827c352)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-024.md#canonical-8214128d24286124b7c24a2f6af302f3126fd4142aa27b61c6436ba399761ba4)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-b00098c79772c870d5911b6f27ad01cc9b4a1f54b88ba7aca116d8b35b71e29f"></a>

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

<a id="canonical-3f222852faf18dd84869f09b0c63331a90cff86fefe2b8e2a984b75d2f23cb56"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c7c642c5d25 / 3

<a id="canonical-2a7509aacc94916396962dcc4944f64c980737fe0f68c50db6438d7595f3950b"></a>

<a id="canonical-3f6438eb8f010a497dfab72db839768237b952fbc57e24a2aff89ca952f6d15c"></a>

## path property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c7c642c5d25 / 4

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

<a id="canonical-b69d23b2871ff17edbc41b186a62d969b6427aa72d621f159ff5931b0ae85010"></a>

<a id="canonical-f48a6d1eb755bbf51112e5e9e40eb347a564325132258e3e0b944f175c1acf8f"></a>

## prefix property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c7c642c5d25 / 5

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

<a id="canonical-2055dce3768c31201498631704bd36d0eeadde0d686bb31213ff094a95c1b26d"></a>

<a id="canonical-cfbb2bf25d0f8df084842427f64b98303c259bd22e7144e644de21129f8c4bad"></a>

## regex property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c7c642c5d25 / 6

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

<a id="canonical-17ba040ba216d5582ec8420dbcf09cda45f12ccab4dca16fdcf7df2a7a80a430"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_lo / 8c7c642c5d25 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-024.md#canonical-8214128d24286124b7c24a2f6af302f3126fd4142aa27b61c6436ba399761ba4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4f57cb95ce001efb49dfd84bf08143f0bdbdf929271f32aff825a054eafdfa5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4520f32143007e7ad7b5f722cca522dad6378e7b4d9105797be79238c3c9ab9"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port / 5dc263ba3b4f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port

<a id="canonical-26ce1268d2c7b1490aa753ea6d4e0d95eaf52eab5c25dc7944f7900146b3af21"></a>

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

<a id="canonical-04f0135520c156941d21b757cbe713b4732e66442d4a8f3d2c9c98d6209da4a9"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port / 5dc263ba3b4f / 3

- [info](resources--workload--reference--group-024.md#canonical-3d05d00abba1632f79de3912007d81b15e869d525848994b0935decd715f7ce4): complete subsection reference.

<a id="canonical-d53c0621371c86e94fa4ac377fa09013ac442e9b2dca64b4c141dcafc32047c4"></a>

<a id="canonical-90c9049ce354c338499595feaa7d2927b5ff2618935377134472344ff7e9934f"></a>

## name property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port / 5dc263ba3b4f / 4

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

<a id="canonical-3c0cca2e116fe89a4e1f95c73e0990a8f8be5910e494b6e639182a976237390f"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port / 5dc263ba3b4f / 5

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-024.md#canonical-3d05d00abba1632f79de3912007d81b15e869d525848994b0935decd715f7ce4)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3d05d00abba1632f79de3912007d81b15e869d525848994b0935decd715f7ce4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66e5bf89b42b886ddf401ca1057c5b16a945a198c67e1c5ca48c56938d82da02"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 567ab73c68a3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-024.md#canonical-4f57cb95ce001efb49dfd84bf08143f0bdbdf929271f32aff825a054eafdfa5d)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info

<a id="canonical-5bedc035d478132c854fc5caa4ea11c608e1a5e5747aed7137f5bb56b2073c98"></a>

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

<a id="canonical-0636da45c196661ffbe61c226624c5859dd31e62015eb58b33e94ec6b090f6ad"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 567ab73c68a3 / 3

<a id="canonical-c85eefd873913384adc75404346b6def5ad99026b24e017554e736c1398253a3"></a>

<a id="canonical-014be1a3738d288410a6b23d711145b26c2bf3acd34df43f550d0ca19faf84ad"></a>

## port property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 567ab73c68a3 / 4

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

<a id="canonical-a4f0edbba91d25c6bf06bbe3ce47898c5d95b9c0ab37d1f6c5b70cb6d67074a1"></a>

<a id="canonical-4a52cee4300500c46ca443069a95b233b403e96b668376df48a8488ab908e746"></a>

## protocol property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 567ab73c68a3 / 5

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

- [same_as_port](resources--workload--reference--group-024.md#canonical-c595014948633bca858843d3337a23aac934bfdac672ad9eed5806ce0a016dfc): complete subsection reference.

<a id="canonical-df985822ce97dccd8ddf6b2adfdc924f787f50a544ee726a0dc0d3b41f32fea4"></a>

<a id="canonical-f4f78afc629792181c49cb84d34792f9a86bfc8ea10b8701706a686a94f0a83b"></a>

## target_port property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 567ab73c68a3 / 6

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

<a id="canonical-b6004ed8b57327409dec858e4b95e1980213f93280aa0ed1814a1ca2a8a43fef"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / 567ab73c68a3 / 7

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port](resources--workload--reference--group-024.md#canonical-c595014948633bca858843d3337a23aac934bfdac672ad9eed5806ce0a016dfc)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-024.md#canonical-4f57cb95ce001efb49dfd84bf08143f0bdbdf929271f32aff825a054eafdfa5d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c595014948633bca858843d3337a23aac934bfdac672ad9eed5806ce0a016dfc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d912556a96b9a11e234d381da62a51c97b3fdc70c8795e9665486891e4f65503"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / c90b59168c37 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-024.md#canonical-4f57cb95ce001efb49dfd84bf08143f0bdbdf929271f32aff825a054eafdfa5d)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-024.md#canonical-3d05d00abba1632f79de3912007d81b15e869d525848994b0935decd715f7ce4)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port

<a id="canonical-2015ce25fd6add78f98f67d3c96e094f4c9f714035ed0e3693fe65ea08c7aa95"></a>

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

<a id="canonical-05257774acc8e6f7aac00f3f997b194d7c781c77eae6ec319b72245ed92dcac4"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / c90b59168c37 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cad777076339836449f24e2e0522efaa90baa19493c13bda68694e52cffa554c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.in / c90b59168c37 / 4

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-024.md#canonical-3d05d00abba1632f79de3912007d81b15e869d525848994b0935decd715f7ce4)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3561ccb026c2d7060c92d2f7b922913f055e6878453c25bf5c512cf37a84e927"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbe1dd47bc4e41e9b6325d790b8b9a170e0fce630bf91bfdb549291f38905ecd"></a>

## stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loa / 11e34abe2b3b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2)
- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer

<a id="canonical-3c528c8203116c0d6d56ab31cc0b30a1c0d38aa8055558e577beb64251d6602e"></a>

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

<a id="canonical-559e192af0ddfff1164abf7f04449af4f18ceeee4da79422b2be42233c825b30"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loa / 11e34abe2b3b / 3

<a id="canonical-674790d7f0d1c9e329ac68b8d874b8ca2c4155aed9a3eae98135fe9c6110127e"></a>

<a id="canonical-f8845e41b4a153d292dee727118fbfea17e22112dc46e8ca9fb55ab300b93f35"></a>

## domains property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loa / 11e34abe2b3b / 4

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

<a id="canonical-baaa7e57424ead61ffec27ce41b839df2b8b677c0aecc186bb7fef004369cd52"></a>

<a id="canonical-444542d6c89ffea0b646aab4ebaeceae34ea9635cb5c143e160360a1e629ddda"></a>

## with_sni property — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loa / 11e34abe2b3b / 5

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

<a id="canonical-0ecc26d28c9ee7837c8607199d23b69ef91aae0f1c3296d532d6152bde679de0"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loa / 11e34abe2b3b / 6

- [stateful_service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-021.md#canonical-e781669fddc2edc89b0b1f6c35cb051a4d49a342608a72820a90a8fb1816a31f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac3821718b29ee9fddacb7c58f7635752a30aa5cfa783dc09681942db2ba2e4e"></a>

## stateful_service.advertise_options.advertise_on_public.port — stateful_service.advertise_options.advertise_on_public.port / 63b40d6b9fb0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- stateful_service.advertise_options.advertise_on_public.port

<a id="canonical-859542a9fdb8f9272354407b73f0b1f55e16cf1716bb1809f7a950cde8df2c6e"></a>

Type: `"object"`. single nested block, Optional.

Advertise Port. Advertise single port.

Upstream description:

Advertise single port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
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
  "x-ves-oneof-field-advertise_choice": "[\"http_loadbalancer\",\"tcp_loadbalancer\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-d51569bdea737c00b68eb0d3a66fd7fe1f85a3989eeff38d47a38ab91c8be185"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port / 63b40d6b9fb0 / 3

- [http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6): complete subsection reference.

- [port](resources--workload--reference--group-027.md#canonical-bb47b02cdccc75dd761ed9ab66ab89e88700f588b8c4e7e2d56f03028776439a): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-027.md#canonical-21344125146a5f70909dbb0ea47658801b199f869e38cd8b15b6524db8ac83dc): complete subsection reference.

<a id="canonical-b51ea555312d1b3b49668e3051fba0bb6537126e9a6c1e76aaca7a0ffa38f43c"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port / 63b40d6b9fb0 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-027.md#canonical-bb47b02cdccc75dd761ed9ab66ab89e88700f588b8c4e7e2d56f03028776439a)
- [stateful_service.advertise_options.advertise_on_public.port.tcp_loadbalancer](resources--workload--reference--group-027.md#canonical-21344125146a5f70909dbb0ea47658801b199f869e38cd8b15b6524db8ac83dc)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d298edb70badab02f61375b05f68469dd3766f4503339aa71a0cfc467820ab2"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer / bbd73ba299b4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer

<a id="canonical-548a02233994518e65787ee955a9f62d9f5ab96924ff576d76588df6d29f9b6a"></a>

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

<a id="canonical-263173bd55606c167fa5c42daacdd6e8127490bd3d10624be74b7601b7d3575a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer / bbd73ba299b4 / 3

- [default_route](resources--workload--reference--group-024.md#canonical-6ca970e472453a680175e7308d0d8ac5382559df9731634e57a688fe5a77c8e2): complete subsection reference.

<a id="canonical-bb0b7ccdca3e60df888d1a6d051644dcea6169c9c35f12d78f2df871dd7519dc"></a>

<a id="canonical-4bdad4e9bf0b30db23e6e7018029cfc86a9f058e7c0369fdf084d7101f418c86"></a>

## domains property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer / bbd73ba299b4 / 4

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

- [http](resources--workload--reference--group-024.md#canonical-9bfd04eaa5c72026d7d9e5bbed797243c7ca7218c784d27cb799ed1605a17e84): complete subsection reference.

- [https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-026.md#canonical-a8e3c43a7c91119e31cbd5fef7dde8ebbc6e56c52af975da2331e4636d1843ca): complete subsection reference.

- [specific_routes](resources--workload--reference--group-027.md#canonical-b0a1888e987fc9d2be4100df113bcbb2b247840272fff2eaeeafbbac38528ae1): complete subsection reference.

<a id="canonical-97c0bcde11063ca5a5ec7f67b3521f31784c0be3e5bd114a570817a359e3e869"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer / bbd73ba299b4 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-024.md#canonical-6ca970e472453a680175e7308d0d8ac5382559df9731634e57a688fe5a77c8e2)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.http](resources--workload--reference--group-024.md#canonical-9bfd04eaa5c72026d7d9e5bbed797243c7ca7218c784d27cb799ed1605a17e84)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-026.md#canonical-a8e3c43a7c91119e31cbd5fef7dde8ebbc6e56c52af975da2331e4636d1843ca)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-027.md#canonical-b0a1888e987fc9d2be4100df113bcbb2b247840272fff2eaeeafbbac38528ae1)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6ca970e472453a680175e7308d0d8ac5382559df9731634e57a688fe5a77c8e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09674a00eefe1ffd458549d9cd6e14b2c7510de1734a816a681a78ada57020e4"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.de / f62e09b84240 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route

<a id="canonical-c8bdc34a1e811c6c01a30b2703767dd23a5bd5223106d18d26d687feb3dabf83"></a>

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

<a id="canonical-70697b16a319c36de7ff925ef1614aa86fcc470e67c3aa7f28f2c8beaf616b26"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.de / f62e09b84240 / 3

- [auto_host_rewrite](resources--workload--reference--group-024.md#canonical-25863e936ac069cba3a0ff401a4bce13b47073be0091da3eaeac1e29c8b98d88): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-024.md#canonical-7d1a28f0c87b4ed8be61f694d08728d047bc16f50458a83aae7aa2a8f17ba2fe): complete subsection reference.

<a id="canonical-f1d04abe613f1965afc9299b8dbb6c03bb05b93a461ff3e301e481fa86c878b2"></a>

<a id="canonical-f941221c952528f926e588f6a062723ee568de5387d325c2db216d05903bca90"></a>

## host_rewrite property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.de / f62e09b84240 / 4

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

<a id="canonical-6ceb13a9684f00e8c36d213aee80a6f691b8713d9b7e211a2165c89aff3d9d53"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.de / f62e09b84240 / 5

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite](resources--workload--reference--group-024.md#canonical-25863e936ac069cba3a0ff401a4bce13b47073be0091da3eaeac1e29c8b98d88)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite](resources--workload--reference--group-024.md#canonical-7d1a28f0c87b4ed8be61f694d08728d047bc16f50458a83aae7aa2a8f17ba2fe)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-25863e936ac069cba3a0ff401a4bce13b47073be0091da3eaeac1e29c8b98d88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9964fce4509bd2b396088b668ca2b1e09610e9ead6bac74a2dd9c929dccba06"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.de / 281d188743cd / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-024.md#canonical-6ca970e472453a680175e7308d0d8ac5382559df9731634e57a688fe5a77c8e2)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-993260273de57458019e186d2ea5ed3c0ffc22d7c10fabb8a3076e9b48aab573"></a>

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

<a id="canonical-db0218edac8b1d17665d4e83ba769c08f01ef7c9d34357511aa8b82e2a9d27b4"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.de / 281d188743cd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-71133a9ab607588ad7e87b3b710d02b7e052b5640d88cc17b37bad3440d666e6"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.de / 281d188743cd / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-024.md#canonical-6ca970e472453a680175e7308d0d8ac5382559df9731634e57a688fe5a77c8e2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7d1a28f0c87b4ed8be61f694d08728d047bc16f50458a83aae7aa2a8f17ba2fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb6be4a3f3f9f035d25693ff53ac0cf0979f8c43ae3faab0f89bdb45432c77e7"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.de / ebb7d6a673ec / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-024.md#canonical-6ca970e472453a680175e7308d0d8ac5382559df9731634e57a688fe5a77c8e2)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-6622b11089d1210eede64441eafb693ab1452aa53a1ffcf4b9de06d9803d590b"></a>

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

<a id="canonical-ace2f707df030802f3a0a5479ae520c85e747c46400847b9ffa36ccb54a20331"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.de / ebb7d6a673ec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03318d732c2c43a880ec6e2b48a9d662df454e3fc486668a65616a7b8f739474"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.de / ebb7d6a673ec / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-024.md#canonical-6ca970e472453a680175e7308d0d8ac5382559df9731634e57a688fe5a77c8e2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9bfd04eaa5c72026d7d9e5bbed797243c7ca7218c784d27cb799ed1605a17e84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1fe8ed4341f6210718d44aafddad1ecc9e428f3d2adf74532126cff1d9cea209"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.http — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 552d54bd1c31 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.http

<a id="canonical-3a18893a9df15928663abfa5be2436c4e8034d7bdd7235f5587c4724ed541502"></a>

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

<a id="canonical-f03ab6206dd54434666059d1a2a548bdec8db35f0f8ad10c518dc4b2a3109a35"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 552d54bd1c31 / 3

<a id="canonical-4e43cfbc9a1f47b60658d1ad992b8ac66a1475db530278c2b242de10eea6f43f"></a>

<a id="canonical-ab826d25d2cad5d1786ea180873d608a4bf785f74b8808123ff6f6021db7e29f"></a>

## dns_volterra_managed property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 552d54bd1c31 / 4

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

<a id="canonical-e92196b6841b7784b67c17835a80cfb1a96e3ff7ff7b970c3aa0d5e271e21b4a"></a>

<a id="canonical-4124bed1d58bd0a908fbad28f7191921c6925ea41c776a752002b3e536f093ba"></a>

## port property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 552d54bd1c31 / 5

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

<a id="canonical-0ec79be3678fb26a48b0d23492f1000f6d3c4d3f7e04fa997c7901f06a993cfe"></a>

<a id="canonical-f887505c3081342aa13c447e9fc021a865b832fc8316c5c9e7425e60c3256eb7"></a>

## port_ranges property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 552d54bd1c31 / 6

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

<a id="canonical-9a68c09b47d6a0b6f8624cafe0863678f34801084f1ee35a6c49a72f7f3076ab"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 552d54bd1c31 / 7

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d857b99a2c0392bcf6e76f66977e185c33f6796dd1c5a11ed7e726a0087324d0"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 44f15baf03bc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https

<a id="canonical-36093cd36ca0193d411dd2f07181298e74a3a6705621d9b1b6d6de1d20f98d68"></a>

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

<a id="canonical-ec298f5d005f6ecefc38394c9a74ef16f35c45599d05582e5b373886aa54fe6a"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 44f15baf03bc / 3

<a id="canonical-4d45f0790ac2599e71a50a76d5ebe742f0550412c95d37c10ad663cef79080d1"></a>

<a id="canonical-a61877814a1750bf3df7bb074d46b65fbc118ed1b625398b0214925aad1b9d96"></a>

## add_hsts property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 44f15baf03bc / 4

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

<a id="canonical-7d47535d58bbd88e40ae152b1ab118f8a7030afb76565f469f46b836d8ecdfbd"></a>

<a id="canonical-4d4d6daba1699e874ee16dba6043939136a2547009c8dfa0979552e8895f3d75"></a>

## append_server_name property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 44f15baf03bc / 5

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

- [coalescing_options](resources--workload--reference--group-024.md#canonical-0da264a78b0645cce2512760b7ea41aab1aaadce2251dca00f653b3ff91bf14d): complete subsection reference.

<a id="canonical-0fed85ceabbc1c1307d4fa155e3598f8c371e0e51f4c62ff6f1a0cb83affbd67"></a>

<a id="canonical-ca715af5dfd2053d8d41b3af1d4f1ebb14b679715b1209fabd393aca3fb99bc3"></a>

## connection_idle_timeout property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 44f15baf03bc / 6

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

- [default_header](resources--workload--reference--group-024.md#canonical-31af0f615650b13076ee78e2ffd93a90e100408851c9043239f32391a17557b9): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-024.md#canonical-e00e1957191e003cf9b218570b3e95a2688490b92cc63d937ac5f8c236a46e84): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-024.md#canonical-464fdb62350db1b2199d4a55fe53ac16e21e7dcbe9756b8f09b8038ec730cfd8): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-024.md#canonical-9203c00238a087128f84f47c9882457244b7306c6ba35785355e60006e2faf5d): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7): complete subsection reference.

<a id="canonical-b082c4a294eb32600abc772dda52e422fe62f88a91bc37ba3cbf61f88dcf167f"></a>

<a id="canonical-e0ffe20492efbd9fd4285989286a310bfc82970f32fc0801dec1ae7d58f577e6"></a>

## http_redirect property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 44f15baf03bc / 7

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

- [non_default_loadbalancer](resources--workload--reference--group-025.md#canonical-852102dc74be9c177a92828aa9e940c8f6f9416323b1a52561778bfb60a87789): complete subsection reference.

- [pass_through](resources--workload--reference--group-025.md#canonical-7b295371e956977ac8a72d8e065faf5d1796e294e82de61df7928f7f6f1e86ba): complete subsection reference.

<a id="canonical-db9cf1efbb67f116c13403bd5ec8e6ce01fc85b1981125886ea09b5aec9ab352"></a>

<a id="canonical-1d40b2506e77416cd8d53fef6e5cd378586d55eb757d9dfe0e2cbc850e33499f"></a>

## port property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 44f15baf03bc / 8

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

<a id="canonical-5093dd0c63592b64c03059ab50dce542eb76dbd3153745b71e89bcd6d66d9a06"></a>

<a id="canonical-510e314f525261d28d00c477c8bfe51a2555791a61776ad64d75141bd583a2ec"></a>

## port_ranges property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 44f15baf03bc / 9

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

<a id="canonical-676767110f4cb3757defd99b72ff528777874137ae4a56391b9a8ba19ad729e8"></a>

<a id="canonical-f89ce3d937dd43be1288552a6abeeecf6847517d89d400123bdcbc971970249f"></a>

## server_name property — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 44f15baf03bc / 10

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

- [tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d): complete subsection reference.

<a id="canonical-8a07279812a956d09838b9f2423ded2b4cbea369a52a0e003f493f74c8eb906d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 44f15baf03bc / 11

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-024.md#canonical-0da264a78b0645cce2512760b7ea41aab1aaadce2251dca00f653b3ff91bf14d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header](resources--workload--reference--group-024.md#canonical-31af0f615650b13076ee78e2ffd93a90e100408851c9043239f32391a17557b9)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer](resources--workload--reference--group-024.md#canonical-e00e1957191e003cf9b218570b3e95a2688490b92cc63d937ac5f8c236a46e84)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize](resources--workload--reference--group-024.md#canonical-464fdb62350db1b2199d4a55fe53ac16e21e7dcbe9756b8f09b8038ec730cfd8)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize](resources--workload--reference--group-024.md#canonical-9203c00238a087128f84f47c9882457244b7306c6ba35785355e60006e2faf5d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-024.md#canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer](resources--workload--reference--group-025.md#canonical-852102dc74be9c177a92828aa9e940c8f6f9416323b1a52561778bfb60a87789)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through](resources--workload--reference--group-025.md#canonical-7b295371e956977ac8a72d8e065faf5d1796e294e82de61df7928f7f6f1e86ba)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-025.md#canonical-3c628a7bc1cdae78492582bbcf546221a1af0a6a9879bb9cc9756246e42ba22d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-025.md#canonical-a8beac99657fd081b6488810eaaffaf6ca1d8f3a81b91771a1ebcb59baf6428d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0da264a78b0645cce2512760b7ea41aab1aaadce2251dca00f653b3ff91bf14d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6f7b52bafde75cd44c1567ad3d7adca024cd85587b9e32759abc6660a379cd8"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 77791ebc48ba / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options

<a id="canonical-a69b8bbf65c1bd23187feda252db55b4653c0c16852d048aa5628cdc9ad35935"></a>

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

<a id="canonical-21a2f8cb3b7f6530f46c0f368240853bab2f1e71504a1209ac890d778bc88371"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 77791ebc48ba / 3

- [default_coalescing](resources--workload--reference--group-024.md#canonical-2d3a73a2cba8f6fea8626bf8c7cb4493725cebdf03e65245c65b99eb9000eb44): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-024.md#canonical-c2d2dc92d745bd2b39b663b512dd82bf96a0ec6acb36a5769be56a5acaa52b4d): complete subsection reference.

<a id="canonical-fc57ee1f02e93b9b57079f2f7ae6b8d5b1b5148b006f2fc22c8c926fba829ea4"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 77791ebc48ba / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing](resources--workload--reference--group-024.md#canonical-2d3a73a2cba8f6fea8626bf8c7cb4493725cebdf03e65245c65b99eb9000eb44)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing](resources--workload--reference--group-024.md#canonical-c2d2dc92d745bd2b39b663b512dd82bf96a0ec6acb36a5769be56a5acaa52b4d)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2d3a73a2cba8f6fea8626bf8c7cb4493725cebdf03e65245c65b99eb9000eb44"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b26f6c2d216f6294349fba338e90674621e8ee1730f5ed991036fcca3028578a"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5316cee7de21 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-024.md#canonical-0da264a78b0645cce2512760b7ea41aab1aaadce2251dca00f653b3ff91bf14d)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-571e7b2011630376d1feeb4e35d3c6ce32611ffd75c64a01b17bab217ba70da9"></a>

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

<a id="canonical-48e29219d43b8826e3141898f3b649ac0f12c31841fb056bcb16a8f24f59ece8"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5316cee7de21 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4f59c9aca4354e04dc290cf96320f6b440c75f232cf858eb368e1ba3451139b"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 5316cee7de21 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-024.md#canonical-0da264a78b0645cce2512760b7ea41aab1aaadce2251dca00f653b3ff91bf14d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c2d2dc92d745bd2b39b663b512dd82bf96a0ec6acb36a5769be56a5acaa52b4d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-268dc04aad518c5c774fbe9c9d97dac57de0eb0893b633d1e6fc570d909402c7"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ca804c65c9c9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-024.md#canonical-0da264a78b0645cce2512760b7ea41aab1aaadce2251dca00f653b3ff91bf14d)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-93d281ac978c47a6887a9d053b87e887413ca8520511c74c2e02164894d27a3f"></a>

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

<a id="canonical-e8d82a8ef583cfa0a9bf5c85beb30955fb642cefc50efcd2bb1c07f43b9d0d2f"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ca804c65c9c9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-eea52428b8a29d820f72f3c70df1f2bc922351afd12d393f4c5d8abc793f7e3e"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ca804c65c9c9 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-024.md#canonical-0da264a78b0645cce2512760b7ea41aab1aaadce2251dca00f653b3ff91bf14d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-31af0f615650b13076ee78e2ffd93a90e100408851c9043239f32391a17557b9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-799517a32767d829a25297187168dba2720bd95e17a9238a2a2026614e78000b"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 49855fab6cde / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header

<a id="canonical-25198c08edb6d39cac7220bb116099fd7c120b30aa23c438fcb425e785aa69f3"></a>

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

<a id="canonical-39aa147368c7566ea8212672a3b9022a4799dc11ee629187f904efc58aff4bcd"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 49855fab6cde / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2b0c1a74d98df6cbfeb3cd2e1126ff432229884f923465118a9889ad15b65d52"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 49855fab6cde / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e00e1957191e003cf9b218570b3e95a2688490b92cc63d937ac5f8c236a46e84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd29b44909f94db02e9eee470ca666950622a5556a5bc2030fbca99b9e4be7fb"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c55f90504eac / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer

<a id="canonical-ee40f5fa54d49fd50bdc1748cfcca8892fc8d697928fd91e9da2b01480e7be6c"></a>

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

<a id="canonical-ddb8e3c395ce2f2c616aff36a3ec740482645fc18a74e242b7dadd2f74be792b"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c55f90504eac / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2511a5cb1e9c3c6695d4081ab4bfaef7cb67ee1b83ef421eb0a01b7e1c855e3d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / c55f90504eac / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-464fdb62350db1b2199d4a55fe53ac16e21e7dcbe9756b8f09b8038ec730cfd8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ba42d38859d9e5fd2e7a46c00dc941a76e2d50d89934d5df4353a6c03fe2acc2"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 254c320bbe3a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize

<a id="canonical-d4b328983684eb5ef9837adaed563d76d3e284338d738b6d44ff74eee26263fb"></a>

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

<a id="canonical-c65e985d1f650081da8e7808e66c84d81461581fa71c7ab61364960584794223"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 254c320bbe3a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ed7bd77131cb23c6754b4cbb1edeec8615f486f6b73c145369a7c4aaca9d328"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 254c320bbe3a / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9203c00238a087128f84f47c9882457244b7306c6ba35785355e60006e2faf5d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c07dd2162f30e084c1f0a51c3f620e60efef453c410295efc6af909ac81696d5"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ca30e2d66bd0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize

<a id="canonical-361aa1651967dc5779a4df18f252fb2e1a3b28e137fae4f43502ab925b966d7b"></a>

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

<a id="canonical-0cf3c658f2d7c10053fb3050162aa740495b7cf8da474e43d3e11880710205a0"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ca30e2d66bd0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f79a1bb01d21ef510c3b1f5e2fbe333025788bdef0afa599a41706fd2d74454d"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / ca30e2d66bd0 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8ec02f648334ffabd08a1a24cb655feb3f94360a06e3130a289d4248a06127c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8af9a84edfd71c4ede3905278f0a3c5ba0ae228837513e59b45afe4bbc47eef3"></a>

## stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 8de0fa712029 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-020.md#canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022)
- [stateful_service.advertise_options.advertise_on_public.port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-024.md#canonical-a309460de31d5a477944352a93673e8a090b1d6c2342355139ca840878a2f2e6)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options

<a id="canonical-05a83f9e91e7297586a172cdd5c158220b75eb9c4141907876f8f95a779e9516"></a>

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

<a id="canonical-147710d14e0a85540edc4687d50c210e6f63088c1398e30e1048cf3356917d87"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 8de0fa712029 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-024.md#canonical-92465093303585154e382d027fedc271e7ad828b7b464e54c366e157a8a8de55): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-025.md#canonical-15441c05ad8def32778d9ef4d5922b2072fdc86634838ba94a24aa885b666301): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-025.md#canonical-ff148f42c7de06c827e69e8040208dcfeb1f0b90105212ab6d7b77792ac87c52): complete subsection reference.

<a id="canonical-35f41b1ef732b231d7a239e926e088093b183f1cdf0f06b849608edffafc0935"></a>

## Next pages — stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.ht / 8de0fa712029 / 4

- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-024.md#canonical-92465093303585154e382d027fedc271e7ad828b7b464e54c366e157a8a8de55)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-025.md#canonical-15441c05ad8def32778d9ef4d5922b2072fdc86634838ba94a24aa885b666301)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-025.md#canonical-ff148f42c7de06c827e69e8040208dcfeb1f0b90105212ab6d7b77792ac87c52)
- [stateful_service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-024.md#canonical-b87765252f8d022952915dc691ece596862e75e46755df054811175f08b62bac)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-92465093303585154e382d027fedc271e7ad828b7b464e54c366e157a8a8de55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
