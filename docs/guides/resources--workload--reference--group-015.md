---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-b5ba3c9f2f9759e28542a75552bfa63bbd560a5059c439ffe2fb7dd3e408bc3e"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / c92acbd5ebd9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-014.md#canonical-fb855e824f1e4e4c62bf461c133a7c810a4b4e7d30cf5adfdfacd007106010f1)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-b8d40bf43f0d6a4fbedc8d5246ae83413219ee951634e999de9bee5899a3711c"></a>

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

<a id="canonical-37582f489f2ed0b2400e84cbc8652cfca091a241dd0b16ff5cc88e50b8b31ae1"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / c92acbd5ebd9 / 3

<a id="canonical-7a6fd48e44fcfb13e7ab6551d73ee6105ded16b0a68a85b5d91f3c9d15d60a9c"></a>

<a id="canonical-1738c58db0e932591585a776310eef5f60b81a907bb3dc5c7d14b2e22b92df0a"></a>

## path property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / c92acbd5ebd9 / 4

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

<a id="canonical-ee0dba5ed7a89a8a6f13a44af1e93f7addae3af7c972ced3ad3379afa8371891"></a>

<a id="canonical-135a9df740a40f5cf067a7e464d411d421a871e326530768635b47d9c5513723"></a>

## prefix property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / c92acbd5ebd9 / 5

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

<a id="canonical-c535d9a13b58a745f5291a32ce11d82be2d4d90f4f0144d6f8b895328a45cde5"></a>

<a id="canonical-a85342700dc655dcc1507c5cb8f9fef0ed96908a629ac3e91583ed6d4ba85fd8"></a>

## regex property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / c92acbd5ebd9 / 6

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

<a id="canonical-ad02a23fc815729f2610796356ab3286096258145ed19183110395680cabacc5"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / c92acbd5ebd9 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-014.md#canonical-fb855e824f1e4e4c62bf461c133a7c810a4b4e7d30cf5adfdfacd007106010f1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9341e4c497647c78854443fff55a4b10f409dd3f80e4c2b48dcc74d0751581fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2de35690c66923228a40d1004ef122381965688fa33f802f2ecc12c6f7dea3ad"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 4e2618879f52 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-014.md#canonical-fb855e824f1e4e4c62bf461c133a7c810a4b4e7d30cf5adfdfacd007106010f1)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-1aa6aaa73027fa353be6b43a3e939e7becad2743e7efb5e30224ac516633ccf7"></a>

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

<a id="canonical-3aadd1e26617047f151ff2435f5ef3d24347a3e66125bf873454013d9803465c"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 4e2618879f52 / 3

<a id="canonical-c35aca9e0be31969c6438debeee058eff5d92430cc4d033468d207721aed63be"></a>

<a id="canonical-5d3dc1af06718275233519eb5ce3d7912f311b024afef955edef51cea619503e"></a>

## response_body_encoded property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 4e2618879f52 / 4

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

<a id="canonical-6627556040802025b676e5c79eb9cc57c9af87e230e24a0bbbb0a3009d1b790c"></a>

<a id="canonical-391b95647484eb3ae291e752abc3719ca1f0aaa0688f34d0ad2d22493d45dc88"></a>

## response_code property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 4e2618879f52 / 5

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

<a id="canonical-120b3c83c22b97b1b8c91300cc1443081a8fd2c8bba65e50c444bb367fed2e66"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 4e2618879f52 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-014.md#canonical-fb855e824f1e4e4c62bf461c133a7c810a4b4e7d30cf5adfdfacd007106010f1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37310f8500c0d66689a9c1fe9aba1bf65ae310d6d9eaf5102546288025ad6622"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 43635f8978e8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-63aca01421458e37e4309716859575944ee464e7c4138ba23e6429ba89468a0b"></a>

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

<a id="canonical-7e8c51a9f81f8c80be80641d22c37b5a5b960eccf6f6c8f1e14cfeda56e0c76d"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 43635f8978e8 / 3

- [headers](resources--workload--reference--group-015.md#canonical-18d3c99b1d0943809702654e9bd53e9aeeee27088bf3daccaa357632158b2f23): complete subsection reference.

<a id="canonical-ad63268ff6219642f23e4e1f9d9b33ef9dffd4d6db21e16682bcad18418086da"></a>

<a id="canonical-7d4c7025c9e41ed1d7b2e8ab3b56f8d8b62d886a9efae3a2f22a1b4a9eea712c"></a>

## http_method property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 43635f8978e8 / 4

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

- [incoming_port](resources--workload--reference--group-015.md#canonical-228e7b489a342eb7e6d1b1de7a0d43119650202540820ef1e0e542643447aaa5): complete subsection reference.

- [path](resources--workload--reference--group-015.md#canonical-f8e5f40ea06c252d8992722cbeb9c758b542c4560e21e4776dec312c0ef9b385): complete subsection reference.

- [route_redirect](resources--workload--reference--group-015.md#canonical-1a64a8cbb460ad959c2ceeadf3d9f4e35274e75502d7d8361f952664e6a09460): complete subsection reference.

<a id="canonical-6311852c26d21ac5b56d2be45e753eec536452712338815660fa7e6787fdcc8d"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 43635f8978e8 / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers](resources--workload--reference--group-015.md#canonical-18d3c99b1d0943809702654e9bd53e9aeeee27088bf3daccaa357632158b2f23)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-015.md#canonical-228e7b489a342eb7e6d1b1de7a0d43119650202540820ef1e0e542643447aaa5)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path](resources--workload--reference--group-015.md#canonical-f8e5f40ea06c252d8992722cbeb9c758b542c4560e21e4776dec312c0ef9b385)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-015.md#canonical-1a64a8cbb460ad959c2ceeadf3d9f4e35274e75502d7d8361f952664e6a09460)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-18d3c99b1d0943809702654e9bd53e9aeeee27088bf3daccaa357632158b2f23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3a185a2ffc8b6cd02ba56f192a13f256e09959a153d30eb4533e07bd1460d066"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 944c4da512db / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-ff1c73e43ca04b99ec82f722cb37740850ab9b7851ebf2ec8eb3c08ebc047a47"></a>

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

<a id="canonical-83cbec09385a47e42dcd00c3321d6c969c9128ca6341303e6118a17c822be049"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 944c4da512db / 3

<a id="canonical-95bf78961a8387e4d04817c94dc4c7ce07953f6d63e5994e70954ce29dd77d1a"></a>

<a id="canonical-61aea15a4394949458baed47ef1ad0a5ade3bbb4804147768d313557dd997fb3"></a>

## exact property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 944c4da512db / 4

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

<a id="canonical-82120854f9e2f4a9187dd2de0d5660af30c9beec8fe688e109863c4f8b8057ce"></a>

<a id="canonical-6ca7bab8947ec70c974dc38305a143f4d73ebf0ed55193c86e3d6a1158e8f330"></a>

## invert_match property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 944c4da512db / 5

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

<a id="canonical-8acb2cbf1fbd6182782c870c7c2c26a9c011ab8bbf52f78a207027c989905ed2"></a>

<a id="canonical-b9b2579eb1d00ab1c581876ded80684439df76c4ffe33f14edcf2092f55618f4"></a>

## name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 944c4da512db / 6

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

<a id="canonical-e0637926a59b8d07312fa49a5dc12ebaca22580f3542fafc5e4205cd9b059612"></a>

<a id="canonical-cd6e1c6796552182cd208511d2c7a6e68543674413363de1e8c88c021aa39cfc"></a>

## presence property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 944c4da512db / 7

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

<a id="canonical-73b86d4424bb35ed763045449e698bf43393abfa2b3724ca1b4d9d30eff233bf"></a>

<a id="canonical-e3aa5344fc19734829e04ad86b926ca850259e42fecae13868f59ec7b1e22668"></a>

## regex property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 944c4da512db / 8

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

<a id="canonical-054c5e26b36311ecc713446e6063a43ce0b610c5c33e1000de970d8a69c79244"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 944c4da512db / 9

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-228e7b489a342eb7e6d1b1de7a0d43119650202540820ef1e0e542643447aaa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f61c223f7a9965d70215c4290d31561e76053d986db0c30a22d7474fc9fe8957"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / c875b4ac0966 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-2b6afbae007ad31e6523858f3c0613a2bd21939bbb0fe0e61134b2feb650e47d"></a>

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

<a id="canonical-369ccebfd9fb436adee2c8dbc61fae53c0244705068f49dfd74256ba1ced68e4"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / c875b4ac0966 / 3

- [no_port_match](resources--workload--reference--group-015.md#canonical-b376cea506aed7d4cd75ce6b19eae2c06f22ea3e7ec72f9c6a83787069e19103): complete subsection reference.

<a id="canonical-278f41a4cd4568f10369f5d97a93c82664aa20fd29a2ba841ae46490f9c0bd0b"></a>

<a id="canonical-1f9d0954b663704d188c7e22b91558e947f296471a109ca0b9f307f478a05f66"></a>

## port property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / c875b4ac0966 / 4

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

<a id="canonical-da2b8a42249352f1f2e52314d6ce5cb0694f942524ebb58e42b891679e423fad"></a>

<a id="canonical-9dd90a7c5f142c78d54cfd6ab7dbc53d5fbbe9caa2b277dc10c3af40c5d118e2"></a>

## port_ranges property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / c875b4ac0966 / 5

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

<a id="canonical-9880447e4e349a6c3bcf18103808af8d1091c56d41df44a077a930ab539d9cc2"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / c875b4ac0966 / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](resources--workload--reference--group-015.md#canonical-b376cea506aed7d4cd75ce6b19eae2c06f22ea3e7ec72f9c6a83787069e19103)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b376cea506aed7d4cd75ce6b19eae2c06f22ea3e7ec72f9c6a83787069e19103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bef6d7e79e7d7920e030302a2daddb3dc602b3d360718fb22fb65f255806332"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / cef512e8066b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-015.md#canonical-228e7b489a342eb7e6d1b1de7a0d43119650202540820ef1e0e542643447aaa5)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-38053833faa9bbe6482dd5f375f5328d748a0253c4d7f08472d29fd42cac2a53"></a>

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

<a id="canonical-a3418479171164ac777e68eddc225f38d13feeaaf26df4825f57e5745f36511e"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / cef512e8066b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-347bd07a7a44f02384bb077ca4a13c409dd176822235508c617d33691aa1e6f9"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / cef512e8066b / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-015.md#canonical-228e7b489a342eb7e6d1b1de7a0d43119650202540820ef1e0e542643447aaa5)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f8e5f40ea06c252d8992722cbeb9c758b542c4560e21e4776dec312c0ef9b385"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8d194e7058b1b9fb924c719f39b9ccf60ac8aa02f3fc47f7d0471b162db8c38"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 69773c4eef58 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-cf29af95b4ec2b379d1614758bababf7ecfc62272eb7271b9d30626773eb491a"></a>

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

<a id="canonical-f7cd4ee3c3951deb2e65aafe6685084bd988a83d6ea1aa3fe703d70caf277d69"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 69773c4eef58 / 3

<a id="canonical-780f72f2dfc6840fdf23eb48d2a741d99eb5aae4a3ce525e4797bf647a98a88d"></a>

<a id="canonical-c0f236a5cb7eed8857d3d685feebca18227cbdf749e68115c8eb4ac77af5983d"></a>

## path property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 69773c4eef58 / 4

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

<a id="canonical-fb70cdb0278f23f0751de6ae941bbd656b5e72c496d6fdd066f3e23551858784"></a>

<a id="canonical-b565b360c3c4f1beeb90ec987323873ef578ad0ca04b10467bc1e30c1d2d5f4d"></a>

## prefix property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 69773c4eef58 / 5

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

<a id="canonical-33b0223a8b6ee4ef8f6bb6410aa3a19f0680a2762ad03825750c1474a9471a85"></a>

<a id="canonical-6f827a086e5b31161b9758552cde28e3b8520653f4b515bcf6378c9c52a743c4"></a>

## regex property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 69773c4eef58 / 6

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

<a id="canonical-51fbcbeabb06ad81b9958ce294a74346696a189783072fcb61d0a434744e0344"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 69773c4eef58 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1a64a8cbb460ad959c2ceeadf3d9f4e35274e75502d7d8361f952664e6a09460"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dee261b40b298b1da0b82b60f0eab0cd01e6baf4d475aef3388b779edcd364b3"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 273c57ed4140 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-106c1ee589cefcd41f2e02eaf27b988504aee93e2de212e01c39831b2aff2df1"></a>

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

<a id="canonical-9c6546e65df0cc227152c60f3ca334c275042f5f3da86c7c6922d3d1346d8102"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 273c57ed4140 / 3

<a id="canonical-98bfada2fc279d6a837919445d712efa84eb3c71ee4785638ba0f73ee80ffcc7"></a>

<a id="canonical-f2c8972b9657fc0c95e7e5ac8b9d69e2bf590d14c199f474a6d9cbc20f307d54"></a>

## host_redirect property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 273c57ed4140 / 4

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

<a id="canonical-dd957ddb26e068a73bc4f3a731641e667dcc3db54f7ddc3a2bb62dd116ce4c74"></a>

<a id="canonical-60affb9b185a63c3dea4bc0decf3c1cf9cc9c42e588e707ad9e13e97959391bc"></a>

## path_redirect property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 273c57ed4140 / 5

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

<a id="canonical-1690743019ec875a2c3540514c3f689bfd43fba8607181499b772e11d7c04cd8"></a>

<a id="canonical-f5a5442d4a086191e89691e54935e3032185290707392ab7fa5c640867360961"></a>

## prefix_rewrite property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 273c57ed4140 / 6

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

<a id="canonical-b35cc95dd28f21eb9cf453ebd0719d2413f1acbddbfded670137b3369ba7dc69"></a>

<a id="canonical-db21cd8df9b4e9867a72f9d01c006e97d3419e1ec95d5a055ffdeb9ba0fb4698"></a>

## proto_redirect property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 273c57ed4140 / 7

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

- [remove_all_params](resources--workload--reference--group-015.md#canonical-54205acb966973222126fa07282cbeca050e44d48c72c89978a4733adf64ec08): complete subsection reference.

<a id="canonical-0adf166aa953602d89d0e7fc02e7a0db4910bb8b5ee5f5d7c60d43fe868a73ed"></a>

<a id="canonical-5937778d489a30ea9ab920290befeecba3aed528e3a0e66a1935adf1ab219496"></a>

## replace_params property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 273c57ed4140 / 8

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

<a id="canonical-65d8c9e7b6b12f91e6b1ce149a9039a5690f0a4b72f25857553ce4ab940b00fd"></a>

<a id="canonical-6a9c907c3c2483e28a30cf066fe0888d21f1b51bc27eab30ecf5ba55a58eb7fe"></a>

## response_code property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 273c57ed4140 / 9

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

- [retain_all_params](resources--workload--reference--group-015.md#canonical-e99f8de0c96913553dea2597e62606ac5c6a1f53a96ed11820e68b786c614d69): complete subsection reference.

<a id="canonical-2233115ddd92961f888716c01eb3407a554140474b57016321ce08025c2205c7"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 273c57ed4140 / 10

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](resources--workload--reference--group-015.md#canonical-54205acb966973222126fa07282cbeca050e44d48c72c89978a4733adf64ec08)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](resources--workload--reference--group-015.md#canonical-e99f8de0c96913553dea2597e62606ac5c6a1f53a96ed11820e68b786c614d69)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-54205acb966973222126fa07282cbeca050e44d48c72c89978a4733adf64ec08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f673a84dad8ff192673329289e777d0e39b7186e82377f9d9f408ad4bd6b14f9"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 8a27b588fa21 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-015.md#canonical-1a64a8cbb460ad959c2ceeadf3d9f4e35274e75502d7d8361f952664e6a09460)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-4e45185aa3502f8de362c32befd946271f2de5a9c17990c03a442eeffd7dff57"></a>

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

<a id="canonical-45b3a74f9cf4a0734f0d3a4ea5c42579fd52db054bf1fe95d14279e49c1f1978"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 8a27b588fa21 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cac7d044dcf48e292b81be7d682de36ac5fa5cf7313205ba941557d61fc3e95f"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 8a27b588fa21 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-015.md#canonical-1a64a8cbb460ad959c2ceeadf3d9f4e35274e75502d7d8361f952664e6a09460)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e99f8de0c96913553dea2597e62606ac5c6a1f53a96ed11820e68b786c614d69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a067c432ad6fe8e6f3241e549dc0f11c62dcc3b59e377ff2ea95a144f67ec56e"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / cfcd61ea6598 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-015.md#canonical-eaad75abe51d826b75461e377e8d784c20bc93bfb7d21bebf977f5b59293fcbe)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-015.md#canonical-1a64a8cbb460ad959c2ceeadf3d9f4e35274e75502d7d8361f952664e6a09460)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-4740daca4f7b80d2ffd9aa3e38c211cb3358fbe773d1322a608634db6ac305f2"></a>

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

<a id="canonical-bc511be0bce0653736c4cc1e43d850c743f0bc10fbf38f0399f95191f3743a09"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / cfcd61ea6598 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b5a306b239153f5d009b4be87158f4dc859017156c99dc3bc218ea70f91cf91"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / cfcd61ea6598 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-015.md#canonical-1a64a8cbb460ad959c2ceeadf3d9f4e35274e75502d7d8361f952664e6a09460)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-dccbcfe0ed23a00dcaa47a904f13229d34ec751001e9559f53ca8bfd2bbb4e88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8315a87504976058c5c0047d76955290c839f3c944f69c5a0cb5c06ebf4d738"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 1afc3328fd5b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-8a02266a0b5c2c073ebbc24fec9f06d9749b50b05717fbce1ca7c58bb6162abd"></a>

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

<a id="canonical-a7d116ed870ccfd77035eda429dcebee5c911d32b44248e4f2707eb66f302698"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 1afc3328fd5b / 3

- [auto_host_rewrite](resources--workload--reference--group-015.md#canonical-3e702cfe623d79a413152eb75002cb789d8d0d277dc81a1cecf53ae95c8bb4d2): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-015.md#canonical-8643e3ab09ec3d047e7e88f333c8d18ed85d131cc493c500f33ada3e6f5e5bf2): complete subsection reference.

<a id="canonical-586d253b28989b1d80f1e438320fbcee1f626769f163786f79038764c7ee0f3f"></a>

<a id="canonical-c6f9cbf41d4f10eb6b7b1a45c183549582e242066317a81c4a70d0468f83316f"></a>

## host_rewrite property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 1afc3328fd5b / 4

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

<a id="canonical-1da1463a86a548006f38acc17f6fe1f44d69f50c18f301a9bc2d2f7b1d46a4de"></a>

<a id="canonical-30e95fb35f69ca251f0ac2c8912902602822a3ddac91d14d278815b6ee8020ad"></a>

## http_method property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 1afc3328fd5b / 5

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

- [path](resources--workload--reference--group-015.md#canonical-e1c0cb34c4647cfe5b0d556cf2295a3e43ec92c0d2107691c01588343f3b4d6f): complete subsection reference.

<a id="canonical-49fe2204350d4cc2952d5bafbcaf0d6041735a7b1b6f70a2a83f9c666c789947"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 1afc3328fd5b / 6

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](resources--workload--reference--group-015.md#canonical-3e702cfe623d79a413152eb75002cb789d8d0d277dc81a1cecf53ae95c8bb4d2)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](resources--workload--reference--group-015.md#canonical-8643e3ab09ec3d047e7e88f333c8d18ed85d131cc493c500f33ada3e6f5e5bf2)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path](resources--workload--reference--group-015.md#canonical-e1c0cb34c4647cfe5b0d556cf2295a3e43ec92c0d2107691c01588343f3b4d6f)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3e702cfe623d79a413152eb75002cb789d8d0d277dc81a1cecf53ae95c8bb4d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af8cb0921f6e2717c8fa3e58859a823608dcbe629bfb8c5b5c6c427ff051f94e"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 850b57fefd5b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-dccbcfe0ed23a00dcaa47a904f13229d34ec751001e9559f53ca8bfd2bbb4e88)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-fde3eb46041af13b4907463df114eb456942fe7da89907263f9867e67a3e0e1a"></a>

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

<a id="canonical-fc4f68e8d4ca175ac0198e673d497d6e554e4e55f5cbf5030ceee2f1038e8102"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 850b57fefd5b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-82f0a5a0a420683c60ef4c80fda5901e577993ad995bbecebd18a95820a026de"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 850b57fefd5b / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-dccbcfe0ed23a00dcaa47a904f13229d34ec751001e9559f53ca8bfd2bbb4e88)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8643e3ab09ec3d047e7e88f333c8d18ed85d131cc493c500f33ada3e6f5e5bf2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-588c3ae56a9dff7be167fa4ba87961eb3314717543600b03263c19739279496f"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 36d271074b98 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-dccbcfe0ed23a00dcaa47a904f13229d34ec751001e9559f53ca8bfd2bbb4e88)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-22e672f2196a6b0373efa5b98eddd01dd39c1ee4c90bfbd690f2037954880feb"></a>

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

<a id="canonical-20802ca02dfefe0c451208f4b5544f85a7052900757afdedda8977f638bb9ebc"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 36d271074b98 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f4ddd26c8b1a4d7315cdc4207f88644f48160a1740960bf1fa9648aaaf15f6d1"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 36d271074b98 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-dccbcfe0ed23a00dcaa47a904f13229d34ec751001e9559f53ca8bfd2bbb4e88)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e1c0cb34c4647cfe5b0d556cf2295a3e43ec92c0d2107691c01588343f3b4d6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1388e31f6552fc56a0c7335d7b10154282dc9207f5c01ae584f57c0157944a59"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 96bd5e30db58 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-014.md#canonical-78c808604766f67e60ccddff9d8264348dab301f07630938ce853888134b6f2e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-dccbcfe0ed23a00dcaa47a904f13229d34ec751001e9559f53ca8bfd2bbb4e88)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-74871f8dda97f1c75f95580929f681ce2d063cee0c7c4fd1d6dccf620837a16c"></a>

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

<a id="canonical-bbf05d7301084c4bd74615d23f9767359fc0a5cba6e8da77cf441b73a84a5706"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 96bd5e30db58 / 3

<a id="canonical-911fc8b67007d71c4d447e42aaa5c549f2fc6b843984dc433d9bec50601dbc38"></a>

<a id="canonical-fb950a323805a9a3b4f8dd3ecf063c6bcf18cde954202d7d2deb78f5831c60c1"></a>

## path property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 96bd5e30db58 / 4

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

<a id="canonical-71e51795a2cbcb08278e4fbde5eea3a39d34236f8895c400aa3ae7fa9da64072"></a>

<a id="canonical-b872f54fa8853b3a9bdf98fa433ac9628c8d44146077cbe91f24e45219382060"></a>

## prefix property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 96bd5e30db58 / 5

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

<a id="canonical-8063d75f5f023dd2b5944419640879c55f22a49d3273239bb6d1159d9b01a90d"></a>

<a id="canonical-14069bf8b4782a5dbd1662614c7fb36b47fdb0154406810438b646bd58514f4b"></a>

## regex property — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 96bd5e30db58 / 6

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

<a id="canonical-306830ffb886732468ce38b1181df5ae09c0b3cda41c8951fe381cdc7b6b7247"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_ro / 96bd5e30db58 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-015.md#canonical-dccbcfe0ed23a00dcaa47a904f13229d34ec751001e9559f53ca8bfd2bbb4e88)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-11e66f627d5d7a33e10e8d9045af1ca879560472658dea83732f50de97228951"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbe978c5f2bbadb1bd45585825202d1c2f64797b67f77df3492849179b203130"></a>

## service.advertise_options.advertise_on_public.port.port — service.advertise_options.advertise_on_public.port.port / 6326ad445c9c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- service.advertise_options.advertise_on_public.port.port

<a id="canonical-cf1a735ffaf9330e53ba1214bb489ac9f70c9bebd76a84d6796263ffb6c40e9c"></a>

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

<a id="canonical-3a82fdfa9f26aecd2b20d2bdf5bdb38213aa47aa82864965ce8da7c1f59042ec"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.port / 6326ad445c9c / 3

- [info](resources--workload--reference--group-015.md#canonical-afe528bfe9b71ae3136a1b37027d6d9ef2c8a3b31d4d031f83d4885cb018007d): complete subsection reference.

<a id="canonical-33adb4e7f33ee5557ed5fbbd15353fdd0b8c45cd25de12c3f765fecae850c766"></a>

## Next pages — service.advertise_options.advertise_on_public.port.port / 6326ad445c9c / 4

- [service.advertise_options.advertise_on_public.port.port.info](resources--workload--reference--group-015.md#canonical-afe528bfe9b71ae3136a1b37027d6d9ef2c8a3b31d4d031f83d4885cb018007d)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-afe528bfe9b71ae3136a1b37027d6d9ef2c8a3b31d4d031f83d4885cb018007d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4faedeb49159530c47e406a55ebcc83981ceb06c6caab24c2d0208aa7ac6ba46"></a>

## service.advertise_options.advertise_on_public.port.port.info — service.advertise_options.advertise_on_public.port.port.info / bf9de205554d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-015.md#canonical-11e66f627d5d7a33e10e8d9045af1ca879560472658dea83732f50de97228951)
- service.advertise_options.advertise_on_public.port.port.info

<a id="canonical-40c2c6a3d548e40928c0fc564d9f9217a7eb07fee6515da9c8ceffb7be4e79f0"></a>

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

<a id="canonical-72e36835bdd309838e1e473a2e56cdd1b71e8cc64bbad9fb3c846eaeae2d531d"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.port.info / bf9de205554d / 3

<a id="canonical-9a0c38712e244d206fe1715d2734115b5e06210342b29adefca1790ee913727e"></a>

<a id="canonical-fa360523d59dee9e5a897c4805a74ebb32687bf32d252223fefbafc84c7b93c1"></a>

## port property — service.advertise_options.advertise_on_public.port.port.info / bf9de205554d / 4

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

<a id="canonical-1776000894a835662b8555f13526b2831ebe06975ca0c06e58e26bfae14cbea2"></a>

<a id="canonical-e1c8123985c652a251f004f49b531ef2492d6ede848b314969caad7e515a12c8"></a>

## protocol property — service.advertise_options.advertise_on_public.port.port.info / bf9de205554d / 5

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

- [same_as_port](resources--workload--reference--group-015.md#canonical-4cd0c54b91a9a0b69c8f7ae68ae3b03ddfe95a297625438b5bbd30f7d06e5d00): complete subsection reference.

<a id="canonical-de256bc88311aa74521ec5454130a741203a2601af7051e9fccf90a59da551a2"></a>

<a id="canonical-41086c3c90e77dc50e8a57e10d78ab6b13a5928690ee2846bf52cf9e0448e126"></a>

## target_port property — service.advertise_options.advertise_on_public.port.port.info / bf9de205554d / 6

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

<a id="canonical-823b1768227094ded2c018d5e1841fc18f221c55c7bee12bf0ceea1715d76eeb"></a>

## Next pages — service.advertise_options.advertise_on_public.port.port.info / bf9de205554d / 7

- [service.advertise_options.advertise_on_public.port.port.info.same_as_port](resources--workload--reference--group-015.md#canonical-4cd0c54b91a9a0b69c8f7ae68ae3b03ddfe95a297625438b5bbd30f7d06e5d00)
- [service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-015.md#canonical-11e66f627d5d7a33e10e8d9045af1ca879560472658dea83732f50de97228951)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4cd0c54b91a9a0b69c8f7ae68ae3b03ddfe95a297625438b5bbd30f7d06e5d00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca6582e782dbfa700bdf55a3209035bdf76c31e9815674b0e5fa41c99fe41308"></a>

## service.advertise_options.advertise_on_public.port.port.info.same_as_port — service.advertise_options.advertise_on_public.port.port.info.same_as_port / 58c18a04d6f9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-015.md#canonical-11e66f627d5d7a33e10e8d9045af1ca879560472658dea83732f50de97228951)
- [service.advertise_options.advertise_on_public.port.port.info](resources--workload--reference--group-015.md#canonical-afe528bfe9b71ae3136a1b37027d6d9ef2c8a3b31d4d031f83d4885cb018007d)
- service.advertise_options.advertise_on_public.port.port.info.same_as_port

<a id="canonical-63bb57e07352edc3ce29e28517e0341503164c64d1bd718c7c948398a57f8d58"></a>

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

<a id="canonical-4533cdc55d46b013f8c0b9bbad0a2dd75e8ca26e3158c0a805bf00d5ce800d69"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.port.info.same_as_port / 58c18a04d6f9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d1ba81f2c6ca209f824cba6a714e0ba5fef051c6d9820dd8281de5c522115a1"></a>

## Next pages — service.advertise_options.advertise_on_public.port.port.info.same_as_port / 58c18a04d6f9 / 4

- [service.advertise_options.advertise_on_public.port.port.info](resources--workload--reference--group-015.md#canonical-afe528bfe9b71ae3136a1b37027d6d9ef2c8a3b31d4d031f83d4885cb018007d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-427355fe9cb5b0c0971a8e939299dc09ef6e767ec08e74cbbab5065ec1526dca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d857a962e0adbca0ba1c8285594f21d3fd53d1e0be4f8e0234aec2aa6bb91341"></a>

## service.advertise_options.advertise_on_public.port.tcp_loadbalancer — service.advertise_options.advertise_on_public.port.tcp_loadbalancer / de6a1b41981d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- service.advertise_options.advertise_on_public.port.tcp_loadbalancer

<a id="canonical-2448e9c59bacc2948548c8006ddbb90114c17d87d51035ffb77eaed2976945a7"></a>

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

<a id="canonical-cd0385681f43789ce2c8ea01be004a7a3752d2730a4b81736681b65a1db42f67"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.tcp_loadbalancer / de6a1b41981d / 3

<a id="canonical-06bb7f5a420f4c7b8e5fbad3d6103db690aa5058b2f70b5957c2a8fa548dd719"></a>

<a id="canonical-13c17254932ad0d3cd0f92ae7bb268a7211b4a4444ee875bc8347baaa09f0a25"></a>

## domains property — service.advertise_options.advertise_on_public.port.tcp_loadbalancer / de6a1b41981d / 4

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

<a id="canonical-e9d8c56c02a3d3c1a6b2d0b55660b4ff27237b0b3bf7843b2dfe444f04ca9c72"></a>

<a id="canonical-bffaf03027b925e6bc0740f24912514794f43c623af710c446baa8b895bb7ba0"></a>

## with_sni property — service.advertise_options.advertise_on_public.port.tcp_loadbalancer / de6a1b41981d / 5

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

<a id="canonical-a8b4db4a0fbe17e077cf559237dc68d503e6963a3d6a5186d7af71917f07afda"></a>

## Next pages — service.advertise_options.advertise_on_public.port.tcp_loadbalancer / de6a1b41981d / 6

- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-34c8bee64727675e6ab55d819968b7a37e89ce671f1a8c0fd53905e1cd97c2bf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5830f5619180a48924991f97c3e27271f6a8d2225565cdb792b95af1f39cf28"></a>

## service.advertise_options.do_not_advertise — service.advertise_options.do_not_advertise / d4d88fd5d7ba / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- service.advertise_options.do_not_advertise

<a id="canonical-d7e1a5bdf95d7a37f99b0694cfde3ee11bb694b76feae8aa2da825238b7ad695"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

<a id="canonical-c96e3d41c2f4c74cf69665a0c1059a03832dad3b1575384f1d2e1a90ad02a666"></a>

## Direct properties — service.advertise_options.do_not_advertise / d4d88fd5d7ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1f9e5a71c46f84ef6f56f3d10ceaf518b518b853efe2e46ba4b597d73f9009c3"></a>

## Next pages — service.advertise_options.do_not_advertise / d4d88fd5d7ba / 4

- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-18c1d9d98cecc2c731f787c41b62d9d968a8d90541ec19c8272ab7e594301dbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4daf856b56f98dfa0a0c413bfd45344c8ffd8dfdabaa29a89fce13b4167332e8"></a>

## service.configuration — service.configuration / 9bf187fc994f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- service.configuration

<a id="canonical-df8e07d8afed369c95b3e2002979d1264f67c66b459d1cafbf3a738713a8a4b1"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameters of the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
configuration {
  # Configure direct properties listed below.
}
```

<a id="canonical-14fb5da76160553f1637dfa042b700b569b46345ebcf100d3e1798f1c8139dcf"></a>

## Direct properties — service.configuration / 9bf187fc994f / 3

- [parameters](resources--workload--reference--group-015.md#canonical-bafb9261a9c9a5775ba4907d29f5b49983f752c2232c3ec5804724941159361d): complete subsection reference.

<a id="canonical-90e44967e7de14376e117479ceaadf266c5a0fc54618eccea728339d34f69013"></a>

## Next pages — service.configuration / 9bf187fc994f / 4

- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-bafb9261a9c9a5775ba4907d29f5b49983f752c2232c3ec5804724941159361d)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bafb9261a9c9a5775ba4907d29f5b49983f752c2232c3ec5804724941159361d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ebc28e356958f77166e8b96e6cc08fdffe79dfaa134473b230e86f4b41af526"></a>

## service.configuration.parameters — service.configuration.parameters / 905c80b5f678 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.configuration](resources--workload--reference--group-015.md#canonical-18c1d9d98cecc2c731f787c41b62d9d968a8d90541ec19c8272ab7e594301dbf)
- service.configuration.parameters

<a id="canonical-c3dc91ab9927575776e50775e4837f5d0b6903a5594fb577c7939d38d3344b05"></a>

Type: `"object"`. list nested block, Optional.

Parameters. Parameters for the workload.

Upstream description:

Parameters for the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("env_var",
    "file")}
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-f64f49a6e75162cbc5937cf4e19cb797d5ab06350e788d3141fb5625f25f521b"></a>

## Direct properties — service.configuration.parameters / 905c80b5f678 / 3

- [env_var](resources--workload--reference--group-015.md#canonical-02f15cd7b34576155d88e51645d0278e2bc78a202e7ca286ef7d6e37cc31bd98): complete subsection reference.

- [file](resources--workload--reference--group-015.md#canonical-09f327e710b38643897b2d9b4c491565c3902d11a18f8c66a777b07f2d432684): complete subsection reference.

<a id="canonical-b9a797979175171ad420fffda06f3d220cadf668704cf278de235bfce8450020"></a>

## Next pages — service.configuration.parameters / 905c80b5f678 / 4

- [service.configuration.parameters.env_var](resources--workload--reference--group-015.md#canonical-02f15cd7b34576155d88e51645d0278e2bc78a202e7ca286ef7d6e37cc31bd98)
- [service.configuration.parameters.file](resources--workload--reference--group-015.md#canonical-09f327e710b38643897b2d9b4c491565c3902d11a18f8c66a777b07f2d432684)
- [service.configuration](resources--workload--reference--group-015.md#canonical-18c1d9d98cecc2c731f787c41b62d9d968a8d90541ec19c8272ab7e594301dbf)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-02f15cd7b34576155d88e51645d0278e2bc78a202e7ca286ef7d6e37cc31bd98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60786563a351e4aa2769ba89c1e08136c27d66e926d5efdbf1ae9b73e46984df"></a>

## service.configuration.parameters.env_var — service.configuration.parameters.env_var / 0f51c5f13522 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.configuration](resources--workload--reference--group-015.md#canonical-18c1d9d98cecc2c731f787c41b62d9d968a8d90541ec19c8272ab7e594301dbf)
- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-bafb9261a9c9a5775ba4907d29f5b49983f752c2232c3ec5804724941159361d)
- service.configuration.parameters.env_var

<a id="canonical-fa031ae93c3387083ec8db43e9d946d0e71d625e253ad3b630d8a1a866a790a8"></a>

Type: `"object"`. single nested block, Optional.

Environment Variable. Environment Variable.

Upstream description:

Environment Variable.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
env_var {
  # Configure direct properties listed below.
}
```

<a id="canonical-1fa4833e3eb333fde8977d3c366ac9ca162a753f0f139dcd3ab1148912512bf0"></a>

## Direct properties — service.configuration.parameters.env_var / 0f51c5f13522 / 3

<a id="canonical-7095d2e65c6d6a1d1b505b0d8f62ddca0dcaec2f8d67b953532d4848cdbed203"></a>

<a id="canonical-cb9818b181816512430696eb683e3e495f09c6692500e2ef732b616d5c7b23b4"></a>

## name property — service.configuration.parameters.env_var / 0f51c5f13522 / 4

Type: `"string"`. Optional.

Name. Name of Environment Variable.

Upstream description:

Name of Environment Variable.

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

<a id="canonical-5a0dc65a67886f49b8bd718f263cdd79251e3fb532fb98405dc2f3a53286bee2"></a>

<a id="canonical-6c52e1530efb622a63404d5327ffa930845a9c0c8e4cec5bb538d8c3b7692550"></a>

## value property — service.configuration.parameters.env_var / 0f51c5f13522 / 5

Type: `"string"`. Optional.

Value. Value of Environment Variable.

Upstream description:

Value of Environment Variable.

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

<a id="canonical-51aceca93283d5dd95ceb0d245e4ec8cf38162369c65a78062a151f407f47583"></a>

## Next pages — service.configuration.parameters.env_var / 0f51c5f13522 / 6

- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-bafb9261a9c9a5775ba4907d29f5b49983f752c2232c3ec5804724941159361d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-09f327e710b38643897b2d9b4c491565c3902d11a18f8c66a777b07f2d432684"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc634204fd6295521237e1395e8d35d5500992f3940088a1332157fa11d5450c"></a>

## service.configuration.parameters.file — service.configuration.parameters.file / 65697580a0a7 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.configuration](resources--workload--reference--group-015.md#canonical-18c1d9d98cecc2c731f787c41b62d9d968a8d90541ec19c8272ab7e594301dbf)
- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-bafb9261a9c9a5775ba4907d29f5b49983f752c2232c3ec5804724941159361d)
- service.configuration.parameters.file

<a id="canonical-1f889af803131acb022bbec0a49cdbae936dfaaf6a6a00f5c1edace7a05efb38"></a>

Type: `"object"`. single nested block, Optional.

Configuration File. Configuration File for the workload.

Upstream description:

Configuration File for the workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name",
    "volume_name")}
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
file {
  # Configure direct properties listed below.
}
```

<a id="canonical-e0ab6df499167c4cf4acf7daa216c6d98308e9bc780ef730307b49b2a42871f3"></a>

## Direct properties — service.configuration.parameters.file / 65697580a0a7 / 3

<a id="canonical-21d4b77c14904252980908f8ac9e7ad0c2732db4b51ba9de9f6f68952a9bcd3e"></a>

<a id="canonical-4db070246293f321a05d18ab10bdc8e5a9c29c6ea8c920da852b21499c602eca"></a>

## data property — service.configuration.parameters.file / 65697580a0a7 / 4

Type: `"string"`. Optional.

Data. File data

Upstream description:

File data

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(16384),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16384,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 16384,
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
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16384",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [mount](resources--workload--reference--group-015.md#canonical-8dde41e6a795bd9e51367c394978c686ba468f75bc8be1a6a16674ef97620ec8): complete subsection reference.

<a id="canonical-87f60ca3342f6e159bd8867f9c9b708444f40b212d8b0d800013c73c382824a3"></a>

<a id="canonical-9d824665ad908f30ea66f24b69ec4b3d8c65f4edecb2ebca9bf9656332fdaf90"></a>

## name property — service.configuration.parameters.file / 65697580a0a7 / 5

Type: `"string"`. Optional.

Name. Name of the file.

Upstream description:

Name of the file.

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

<a id="canonical-61826f29a4a95ff271e9167cd6dfbe8e1ce6aa08842c035bb229dd97fe28284f"></a>

<a id="canonical-eff8d976e754cd4e2e05ade9f45a3702cddc11d0ae75eaebef2a1e1bb583466a"></a>

## volume_name property — service.configuration.parameters.file / 65697580a0a7 / 6

Type: `"string"`. Optional.

Volume Name. Name of the Volume.

Upstream description:

Name of the Volume.

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

<a id="canonical-6f7debc299d3f496c51960176cafe822b04f4f7aa18a12858106cdc94fc6bc81"></a>

## Next pages — service.configuration.parameters.file / 65697580a0a7 / 7

- [service.configuration.parameters.file.mount](resources--workload--reference--group-015.md#canonical-8dde41e6a795bd9e51367c394978c686ba468f75bc8be1a6a16674ef97620ec8)
- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-bafb9261a9c9a5775ba4907d29f5b49983f752c2232c3ec5804724941159361d)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8dde41e6a795bd9e51367c394978c686ba468f75bc8be1a6a16674ef97620ec8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-72e28a8af467455212b8d61a977552e9fc7aa72f49752df324b551d3f136cbb1"></a>

## service.configuration.parameters.file.mount — service.configuration.parameters.file.mount / 6e5740cfac8d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.configuration](resources--workload--reference--group-015.md#canonical-18c1d9d98cecc2c731f787c41b62d9d968a8d90541ec19c8272ab7e594301dbf)
- [service.configuration.parameters](resources--workload--reference--group-015.md#canonical-bafb9261a9c9a5775ba4907d29f5b49983f752c2232c3ec5804724941159361d)
- [service.configuration.parameters.file](resources--workload--reference--group-015.md#canonical-09f327e710b38643897b2d9b4c491565c3902d11a18f8c66a777b07f2d432684)
- service.configuration.parameters.file.mount

<a id="canonical-15895f2c1466640c07976ef8e6c277fdbaa55812ae6ff142276cfa797c7923ae"></a>

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

<a id="canonical-67fda35d3b6dab1c7f93828d41804c11f61e15907086dbca3cea9c3b450469b8"></a>

## Direct properties — service.configuration.parameters.file.mount / 6e5740cfac8d / 3

<a id="canonical-8a272ab3df13dbe36c62da758b29ba5ce1ff77ab695cca948aa0bc6234d188ca"></a>

<a id="canonical-13154bb70677fad883ccb42fef4319d0804a6a369e0649ad41837774557b6d92"></a>

## mode property — service.configuration.parameters.file.mount / 6e5740cfac8d / 4

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

<a id="canonical-1fc78507a26b1fc7e42994bef8585e0c5e5b3ff154734111f9af7b54a0c6335c"></a>

<a id="canonical-e5ea76ff27388a4b9cb73dd11f917763c497caa7a1f31070ed44426330bb582b"></a>

## mount_path property — service.configuration.parameters.file.mount / 6e5740cfac8d / 5

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

<a id="canonical-a64f970dd1a364640f101fbfcf0218ddaf947782c181ee68f79a001a60e27601"></a>

<a id="canonical-5fc7c1367bc224ffa5daa6a6f8bae5305567411bb9c6694940de395aa1ec7948"></a>

## sub_path property — service.configuration.parameters.file.mount / 6e5740cfac8d / 6

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

<a id="canonical-2209f9a49afb8634c56e89317edf8b8f83e09b75ca9a2337a69ca27c2012f4e3"></a>

## Next pages — service.configuration.parameters.file.mount / 6e5740cfac8d / 7

- [service.configuration.parameters.file](resources--workload--reference--group-015.md#canonical-09f327e710b38643897b2d9b4c491565c3902d11a18f8c66a777b07f2d432684)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5703a82ae308f03fc9221ebe8fd240527652b05975f06f27fc439be3af43adb9"></a>

## service.containers — service.containers / 6595f97a5422 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- service.containers

<a id="canonical-cd3cd3ff7b31a97ea1656d66955d8c3d2f234b8dc2faac4675376a0150f215b9"></a>

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

<a id="canonical-c33e0bfe763db0fc3f53cb595905fc2157362b05a237d30c3e4772af464de60a"></a>

## Direct properties — service.containers / 6595f97a5422 / 3

<a id="canonical-a932884c353abc753f5a37dd00b91e298ff63bedd65dbd32617da504e95e40cf"></a>

<a id="canonical-96f2b9e51dffc8a3bf7982bd0208078402d4edf01c0da4bc6051cef428c98a72"></a>

## args property — service.containers / 6595f97a5422 / 4

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

<a id="canonical-4faeaa794e59a5fbd7d451597be228b8a7a8fed563b7ccbeb2925d141eddff58"></a>

<a id="canonical-1a19bce9728932910323d11f88fe4dccba23a46541431abb2280bfe0761f5ff0"></a>

## command property — service.containers / 6595f97a5422 / 5

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

- [custom_flavor](resources--workload--reference--group-015.md#canonical-aa110fa03ca0a4f70dcb7a0b03c490015e07c7fc2877b3a92f09ada17cf685a6): complete subsection reference.

- [default_flavor](resources--workload--reference--group-015.md#canonical-c1bce050dcf762e071c3461a716dfd6b9b9462887d599a3ba29509e3080187f6): complete subsection reference.

<a id="canonical-cd5c506a493492f31c780bcdcb2e1163dbaa1ecba70744a6c279e73f3dcfc535"></a>

<a id="canonical-a8837062906544088fe9a3ed4761dd9ce935330246d21050170d69cedb4c8be8"></a>

## flavor property — service.containers / 6595f97a5422 / 6

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

- [image](resources--workload--reference--group-015.md#canonical-55938d0de3a6ab9691c628685e2c0e6df146b8cdf089968d4a87850dce03d926): complete subsection reference.

<a id="canonical-fb3cd817cc3bd2f06cf49858fc345e4f1dd2cfee54b19e267024b6990759ccc8"></a>

<a id="canonical-012d7cf04b72486f9730a5d3198519507ed6c825f052afed71fdad43ecfaa6ae"></a>

## init_container property — service.containers / 6595f97a5422 / 7

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

- [liveness_check](resources--workload--reference--group-015.md#canonical-9c5d1446ca9c13b82f49b3937282cdf56a75392faefe1cb909e2d3913af0a08f): complete subsection reference.

<a id="canonical-4d56456088efe10f2b75c4babeb33cac62e97f893cfa53d2bd3457f5c3b4be55"></a>

<a id="canonical-d79b6c803da87ada9f4ceb8fd2def646aa7d4f93b1400b70ca1b9389126e3909"></a>

## name property — service.containers / 6595f97a5422 / 8

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

- [readiness_check](resources--workload--reference--group-015.md#canonical-8c84db60fef3270c2e5c9acf436cd213cc7de6870d6207184ca297b292e72b6f): complete subsection reference.

<a id="canonical-d357e64675569e22bb1ae455c69ace099da090bb108a6bbaab5593d0e0e48aad"></a>

## Next pages — service.containers / 6595f97a5422 / 9

- [service.containers.custom_flavor](resources--workload--reference--group-015.md#canonical-aa110fa03ca0a4f70dcb7a0b03c490015e07c7fc2877b3a92f09ada17cf685a6)
- [service.containers.default_flavor](resources--workload--reference--group-015.md#canonical-c1bce050dcf762e071c3461a716dfd6b9b9462887d599a3ba29509e3080187f6)
- [service.containers.image](resources--workload--reference--group-015.md#canonical-55938d0de3a6ab9691c628685e2c0e6df146b8cdf089968d4a87850dce03d926)
- [service.containers.liveness_check](resources--workload--reference--group-015.md#canonical-9c5d1446ca9c13b82f49b3937282cdf56a75392faefe1cb909e2d3913af0a08f)
- [service.containers.readiness_check](resources--workload--reference--group-015.md#canonical-8c84db60fef3270c2e5c9acf436cd213cc7de6870d6207184ca297b292e72b6f)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-aa110fa03ca0a4f70dcb7a0b03c490015e07c7fc2877b3a92f09ada17cf685a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-919f1a5b1a62dc1ce9b3027b99ad0931ad7e81389f670c04e4d7a5cb9ce27203"></a>

## service.containers.custom_flavor — service.containers.custom_flavor / af60bbf26af5 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- service.containers.custom_flavor

<a id="canonical-bc9850148d3baf7d823f29e9fecdbafa13d8e2ee3d4f4d61043822e7772cae4f"></a>

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

<a id="canonical-a71b46d2ce866aeab87ddfd34d36896d7b20545b41cfef9e63dee013caaf3472"></a>

## Direct properties — service.containers.custom_flavor / af60bbf26af5 / 3

<a id="canonical-a49c9f98dd1566a722377c354bac331b227a217ccef22f37c98454004f84b076"></a>

<a id="canonical-1b8d4abc983fedb1075c00a7f9b81624e089566487203ec292b6a4bc8dbac3cb"></a>

## name property — service.containers.custom_flavor / af60bbf26af5 / 4

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

<a id="canonical-b61aa4f88e87a4b0317d2971c9c67b1e9c6e5e34c25b5f24902bd58765cdd973"></a>

<a id="canonical-a317d3ca3fb5c2ee222fc75d2ad7a9b04a3f8324abe1fb8c874048e1f3148724"></a>

## namespace property — service.containers.custom_flavor / af60bbf26af5 / 5

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

<a id="canonical-d74e0f6011ab11154cb12bb88be48f3df0dea6dc2ac6b6a0b0adbcba50b03ba1"></a>

<a id="canonical-82091f9302521149a893a5a1e281730c278030efb3ffb18290775101c6f08717"></a>

## tenant property — service.containers.custom_flavor / af60bbf26af5 / 6

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

<a id="canonical-4229690b39ffdb89bf0e46876c3569da7799756ba7de2bd5e023c3507ba58be1"></a>

## Next pages — service.containers.custom_flavor / af60bbf26af5 / 7

- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c1bce050dcf762e071c3461a716dfd6b9b9462887d599a3ba29509e3080187f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e62b8d059db1dbfdacc051b25817b898d0f6602048ca6d102145f202b2e0a3c"></a>

## service.containers.default_flavor — service.containers.default_flavor / 67711f076754 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- service.containers.default_flavor

<a id="canonical-e7aa3db002ea87b7c17478464b601e4af311cfd0e566b8e2bf6b2b75461d7589"></a>

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

<a id="canonical-11a66b628d17a98f910ae621093d85e0c99178f99d64c3ce1bf2080f83f5b479"></a>

## Direct properties — service.containers.default_flavor / 67711f076754 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0c007e048e399779688dab274469e9442ff32e82abcbce71f50290a1fb0f6a54"></a>

## Next pages — service.containers.default_flavor / 67711f076754 / 4

- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-55938d0de3a6ab9691c628685e2c0e6df146b8cdf089968d4a87850dce03d926"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-facb301082b920866515b4a03bce6692d5759ce915424bca6f71364e6e4d9b80"></a>

## service.containers.image — service.containers.image / 64bdc669ac4c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- service.containers.image

<a id="canonical-b47b550a41580cf52b763193cb8f8fb90c3c7b6a5ccec70f19f7c8a6a0e59395"></a>

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

<a id="canonical-d669cb5002f10c436baaef2bf5ceee34b91ecf0efa84455bcfca899ed15f979e"></a>

## Direct properties — service.containers.image / 64bdc669ac4c / 3

- [container_registry](resources--workload--reference--group-015.md#canonical-09356cab2abf7b050c5b5d5e93af472efbb9bd056b98a1b9ebfa489093e81933): complete subsection reference.

<a id="canonical-59a3194c897b892100d5f6268d238570c5fc18a15d027ee8b9a611e3d67e8e61"></a>

<a id="canonical-8992a837d56935467b4cd83ea6cb57bdc3c811507e591377364088aeda0a4043"></a>

## name property — service.containers.image / 64bdc669ac4c / 4

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

- [public](resources--workload--reference--group-015.md#canonical-41c2379b1e8777249955cd87a19975aa78be9d791eb8b7b655845ee40bf09829): complete subsection reference.

<a id="canonical-fd305901e7a7a94ab2277751f31233aa35173658c72f61fc8b1f1f3a380b017c"></a>

<a id="canonical-48d7b84827d020b99043e770220c4dcb473beb6efe58a7e622020faef21a9690"></a>

## pull_policy property — service.containers.image / 64bdc669ac4c / 5

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

<a id="canonical-3ae3dfa94d5aa364b7093f8b7fb0781244ab36185ccde0011a4cb4d02ae281e8"></a>

## Next pages — service.containers.image / 64bdc669ac4c / 6

- [service.containers.image.container_registry](resources--workload--reference--group-015.md#canonical-09356cab2abf7b050c5b5d5e93af472efbb9bd056b98a1b9ebfa489093e81933)
- [service.containers.image.public](resources--workload--reference--group-015.md#canonical-41c2379b1e8777249955cd87a19975aa78be9d791eb8b7b655845ee40bf09829)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-09356cab2abf7b050c5b5d5e93af472efbb9bd056b98a1b9ebfa489093e81933"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c07199ce192f590393afe9ac58234a4e9b5649c3819de3f3d73e93edec729b43"></a>

## service.containers.image.container_registry — service.containers.image.container_registry / 72d950b38229 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.image](resources--workload--reference--group-015.md#canonical-55938d0de3a6ab9691c628685e2c0e6df146b8cdf089968d4a87850dce03d926)
- service.containers.image.container_registry

<a id="canonical-d7076a52a73603d9818a88a55cf034b4d70640d14348ec7b5c9afbd13bd93f0f"></a>

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

<a id="canonical-71010e2c0f283cdeaf03bbe98b65f7f2bdc93cf07a422f7d53342bfbc69048eb"></a>

## Direct properties — service.containers.image.container_registry / 72d950b38229 / 3

<a id="canonical-fcded7e0b7a5cecb810bff3085cd9619121a065435a6189cc6a866e8f68e7d52"></a>

<a id="canonical-a8a3172baf1a0d186d023a65e3c66242e8db4f8149a004dd998337755c5983b6"></a>

## name property — service.containers.image.container_registry / 72d950b38229 / 4

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

<a id="canonical-9106dec9e6d857d35c9dcd78a99d73b653a927971e1dda7ec9b118cc5a6be239"></a>

<a id="canonical-af839e1ea8b439bbec008ad3a71b67e851dbaa9a405de2245526f2f97ef083e5"></a>

## namespace property — service.containers.image.container_registry / 72d950b38229 / 5

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

<a id="canonical-eaa62eec6cb1bf16d323e21bc20d39f9bd0e48db768415b56cd0c41261cd3f63"></a>

<a id="canonical-dbeafcccc820234ceb8bf7fa70b1ed6db0a5175abf7386fc5bcd844e1f2c306a"></a>

## tenant property — service.containers.image.container_registry / 72d950b38229 / 6

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

<a id="canonical-0e4f4f78dbc32dbbeb748c24b11d198cfe823d3a728958ba2974e81922592df2"></a>

## Next pages — service.containers.image.container_registry / 72d950b38229 / 7

- [service.containers.image](resources--workload--reference--group-015.md#canonical-55938d0de3a6ab9691c628685e2c0e6df146b8cdf089968d4a87850dce03d926)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-41c2379b1e8777249955cd87a19975aa78be9d791eb8b7b655845ee40bf09829"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5e8b75349b0f6b5251c75e7bbaace362cd7ff1bd815a40aa285dedfef25f91b1"></a>

## service.containers.image.public — service.containers.image.public / c98dc89dc76e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.image](resources--workload--reference--group-015.md#canonical-55938d0de3a6ab9691c628685e2c0e6df146b8cdf089968d4a87850dce03d926)
- service.containers.image.public

<a id="canonical-4aa7ae6c0f148d5b90b12bcf9ecf082e0ea67cb8dad97b23f0a4c787790ea688"></a>

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

<a id="canonical-b2467a305bea6934a900ada1744831515fd534bba1b26df56911e45eee94d5a5"></a>

## Direct properties — service.containers.image.public / c98dc89dc76e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a532900a813646b65b7b92cbc9d1555e47081ccb81bee2901f4dada71fe898a3"></a>

## Next pages — service.containers.image.public / c98dc89dc76e / 4

- [service.containers.image](resources--workload--reference--group-015.md#canonical-55938d0de3a6ab9691c628685e2c0e6df146b8cdf089968d4a87850dce03d926)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9c5d1446ca9c13b82f49b3937282cdf56a75392faefe1cb909e2d3913af0a08f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf16e8d077b82c94ff5cb665f4f04c9331cd28e013c94cd026c9a33542ba0193"></a>

## service.containers.liveness_check — service.containers.liveness_check / 7316b2ef9bbc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- service.containers.liveness_check

<a id="canonical-e0db1e874129aa9551f534aa5ae0669d860c7d376954014cfc9418df710c2b6f"></a>

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

<a id="canonical-b00b79579337788f7ac7d9c10a12574367daa1d74a4cb287a80cfce6edf5a02b"></a>

## Direct properties — service.containers.liveness_check / 7316b2ef9bbc / 3

- [exec_health_check](resources--workload--reference--group-015.md#canonical-c6b259c972cf03933ff549172623d9c3abad18fc1de89bf03c00f1461037c7ff): complete subsection reference.

<a id="canonical-e6989ae4c5e1c73ce779927738242ede3f590c0118f463ea779e8765adefbfa5"></a>

<a id="canonical-afa7e1ce62111f1eca18e704ce6cd6cbd826bfb343ef548177ebeecfb6fad878"></a>

## healthy_threshold property — service.containers.liveness_check / 7316b2ef9bbc / 4

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

- [http_health_check](resources--workload--reference--group-015.md#canonical-575c5a1ac90a581e4ce652bab95bf49375d9c14a7aa12728697635e8347ab4f0): complete subsection reference.

<a id="canonical-93c0a450318102ff1f271ccf2407742b79748909573b1352e62dc2e7ddf0d6cf"></a>

<a id="canonical-e98ebfa5a538f213cd09474cf00d820330ce4aaa361868a3d604ad50010863b1"></a>

## initial_delay property — service.containers.liveness_check / 7316b2ef9bbc / 5

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

<a id="canonical-fe9cd627cf025879c104840c4083ce9aaca1a2c0881cf6b132a041f14566531e"></a>

<a id="canonical-f1f263e550c263ef363bd096661f9da882e99c0dfedc5074da2c8c09c23a2761"></a>

## interval property — service.containers.liveness_check / 7316b2ef9bbc / 6

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

- [tcp_health_check](resources--workload--reference--group-015.md#canonical-934f46745896a0bb292e97f47aafff29b13a98a8f215b386b9c05fb5a58da8a2): complete subsection reference.

<a id="canonical-c45a98faa430a8a546ba41478fee28c65a0f8989018dae6aa3e73ee7df62fa11"></a>

<a id="canonical-cd954bd6c25ae2c089dffec7b2bc664e0ada4b4ae59f62dc2655f80030c22517"></a>

## timeout property — service.containers.liveness_check / 7316b2ef9bbc / 7

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

<a id="canonical-01995b2f118e1a951c80464495359ea9369c4105a5ca2ab0962304a649decf5a"></a>

<a id="canonical-fc8c69b7ef6aa3ed3ef0961e677444253ea63214808438f02501e5b713f91007"></a>

## unhealthy_threshold property — service.containers.liveness_check / 7316b2ef9bbc / 8

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

<a id="canonical-775aadec3e789423729c87175d3754e878547b859fdaf21f2cfb7b0c01cb9eaf"></a>

## Next pages — service.containers.liveness_check / 7316b2ef9bbc / 9

- [service.containers.liveness_check.exec_health_check](resources--workload--reference--group-015.md#canonical-c6b259c972cf03933ff549172623d9c3abad18fc1de89bf03c00f1461037c7ff)
- [service.containers.liveness_check.http_health_check](resources--workload--reference--group-015.md#canonical-575c5a1ac90a581e4ce652bab95bf49375d9c14a7aa12728697635e8347ab4f0)
- [service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-015.md#canonical-934f46745896a0bb292e97f47aafff29b13a98a8f215b386b9c05fb5a58da8a2)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c6b259c972cf03933ff549172623d9c3abad18fc1de89bf03c00f1461037c7ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9f3b68ce4e7d15af1b7384788dc19ada07c6870bceddc4aa0e157d068c8e65b"></a>

## service.containers.liveness_check.exec_health_check — service.containers.liveness_check.exec_health_check / 0299de07f4df / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.liveness_check](resources--workload--reference--group-015.md#canonical-9c5d1446ca9c13b82f49b3937282cdf56a75392faefe1cb909e2d3913af0a08f)
- service.containers.liveness_check.exec_health_check

<a id="canonical-e705f7900a007f6b122a6869746d0c90222e966234205dc233b53daacc3d57e1"></a>

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

<a id="canonical-fbf88d50e4895784c728f1098813f856bd0c4f9451d13787e246f49940d69886"></a>

## Direct properties — service.containers.liveness_check.exec_health_check / 0299de07f4df / 3

<a id="canonical-7cceb67717c68ec51b2f9b655c52a634e65be61fc5a55244203c911f0a99a9b7"></a>

<a id="canonical-67c9d75213b710e7062ada6640220318b23bf5896ce82eac748e1f98c7994efe"></a>

## command property — service.containers.liveness_check.exec_health_check / 0299de07f4df / 4

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

<a id="canonical-9a4a6aaca84434b0b056c8d1ca6b74d89bca61fad3bda314c9680f6239251b5f"></a>

## Next pages — service.containers.liveness_check.exec_health_check / 0299de07f4df / 5

- [service.containers.liveness_check](resources--workload--reference--group-015.md#canonical-9c5d1446ca9c13b82f49b3937282cdf56a75392faefe1cb909e2d3913af0a08f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-575c5a1ac90a581e4ce652bab95bf49375d9c14a7aa12728697635e8347ab4f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3cfc7607d1e35b5b81aa5ed503c48f84e0296a505152320a179c661869be0324"></a>

## service.containers.liveness_check.http_health_check — service.containers.liveness_check.http_health_check / 8bf10a4bac65 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.liveness_check](resources--workload--reference--group-015.md#canonical-9c5d1446ca9c13b82f49b3937282cdf56a75392faefe1cb909e2d3913af0a08f)
- service.containers.liveness_check.http_health_check

<a id="canonical-29a00d331e25b8e4acdfde00d67867d1b683802dc5bf6207cb84376ef5b13675"></a>

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

<a id="canonical-3b8161f1a1422fc88ad2bb1f7dcb629b3be83855e29cfefc4ffe6b92de3415b8"></a>

## Direct properties — service.containers.liveness_check.http_health_check / 8bf10a4bac65 / 3

<a id="canonical-5b66d6b0109830e7ad49a0918cb75c3014c9bd7728ff0a0acf5d8087931ad059"></a>

<a id="canonical-4bef0c9ab9001cabdeb819fba37ca88dda01062e471b2aa971e9f0b3c7cf21ec"></a>

## headers property — service.containers.liveness_check.http_health_check / 8bf10a4bac65 / 4

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

<a id="canonical-d80000e9be3dcd2ef7bb35c8ddadace6d1d05158d0fa44b82844776423ac04bc"></a>

<a id="canonical-12505e20a2905682d118e62417d1a442c4e35b1c688ff549fce6813042e6496f"></a>

## host_header property — service.containers.liveness_check.http_health_check / 8bf10a4bac65 / 5

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

<a id="canonical-3576baef3807c9ba5567d54e8146882c7bf369a00be791b9fe249cdb3fd73fae"></a>

<a id="canonical-5104db5570f9298e42787013b2312b0bd743ed37d9b8c070ceeabcb87629f71b"></a>

## path property — service.containers.liveness_check.http_health_check / 8bf10a4bac65 / 6

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

- [port](resources--workload--reference--group-015.md#canonical-0ec65514d869d6947a8835db1fc6a2a86616edfba1cca8173c694723403bfacb): complete subsection reference.

<a id="canonical-e443eda439c90f1e2a0d497b9e36403bd794df753a2069fed4c7bc2879c52275"></a>

## Next pages — service.containers.liveness_check.http_health_check / 8bf10a4bac65 / 7

- [service.containers.liveness_check.http_health_check.port](resources--workload--reference--group-015.md#canonical-0ec65514d869d6947a8835db1fc6a2a86616edfba1cca8173c694723403bfacb)
- [service.containers.liveness_check](resources--workload--reference--group-015.md#canonical-9c5d1446ca9c13b82f49b3937282cdf56a75392faefe1cb909e2d3913af0a08f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0ec65514d869d6947a8835db1fc6a2a86616edfba1cca8173c694723403bfacb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec58b588a0465fe8216abe0ae1e8ee6a6ba8edc243285efc99e9108e22225586"></a>

## service.containers.liveness_check.http_health_check.port — service.containers.liveness_check.http_health_check.port / 60e7302c7db4 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.liveness_check](resources--workload--reference--group-015.md#canonical-9c5d1446ca9c13b82f49b3937282cdf56a75392faefe1cb909e2d3913af0a08f)
- [service.containers.liveness_check.http_health_check](resources--workload--reference--group-015.md#canonical-575c5a1ac90a581e4ce652bab95bf49375d9c14a7aa12728697635e8347ab4f0)
- service.containers.liveness_check.http_health_check.port

<a id="canonical-179b30c02dca7ac36b631e198c58a234044d029a5ce2432ded085eeb90f600c6"></a>

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

<a id="canonical-9878c36aa9ad83ba60c507e8cc41838d71304b3f64ed0f86f3895de3edbab912"></a>

## Direct properties — service.containers.liveness_check.http_health_check.port / 60e7302c7db4 / 3

<a id="canonical-0a969fe3cd220313291f71d5cb4c4abb5ba765a285f6e481bc18e19c43f86e28"></a>

<a id="canonical-9dbea5aeb65cc144e99910bede04bf60d13e6c91951f0bd39b5204d06e4999aa"></a>

## name property — service.containers.liveness_check.http_health_check.port / 60e7302c7db4 / 4

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

<a id="canonical-f93ee365f325c351a9b5023e2e944832ed2a9942e655bf9837d73ab3fa388046"></a>

<a id="canonical-271ccfa304c9e05e35dd6a00f83feaf711d883141a47ed31aa454d35b327e549"></a>

## num property — service.containers.liveness_check.http_health_check.port / 60e7302c7db4 / 5

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

<a id="canonical-c896ce1dd440e8ac8daa383e1ab6221b47b3411e2f43fb81082da8723372bf05"></a>

## Next pages — service.containers.liveness_check.http_health_check.port / 60e7302c7db4 / 6

- [service.containers.liveness_check.http_health_check](resources--workload--reference--group-015.md#canonical-575c5a1ac90a581e4ce652bab95bf49375d9c14a7aa12728697635e8347ab4f0)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-934f46745896a0bb292e97f47aafff29b13a98a8f215b386b9c05fb5a58da8a2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13831f6c0260e4d4ad1846bac3fd2375f5b22ea7979b17cf7050c7d930814208"></a>

## service.containers.liveness_check.tcp_health_check — service.containers.liveness_check.tcp_health_check / 6a17be64dbd7 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.liveness_check](resources--workload--reference--group-015.md#canonical-9c5d1446ca9c13b82f49b3937282cdf56a75392faefe1cb909e2d3913af0a08f)
- service.containers.liveness_check.tcp_health_check

<a id="canonical-0aa849240b2eb896468d37ddaf8ccbb89e3dc7ba36d96d1d88576543e916c842"></a>

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

<a id="canonical-d15b6192ba2f1bbf825cb9f987f4b91e8a7c46f4f459ee9c0de9517e19a05702"></a>

## Direct properties — service.containers.liveness_check.tcp_health_check / 6a17be64dbd7 / 3

- [port](resources--workload--reference--group-015.md#canonical-0f1dcb4ab49708acf1d31bd51816b47f0ecd9857fff0ba744f662db6bbd086fd): complete subsection reference.

<a id="canonical-b45b44d738073e0e7758cb2adf6174cbb56f15ac62d32fc09b291ac28f74ef63"></a>

## Next pages — service.containers.liveness_check.tcp_health_check / 6a17be64dbd7 / 4

- [service.containers.liveness_check.tcp_health_check.port](resources--workload--reference--group-015.md#canonical-0f1dcb4ab49708acf1d31bd51816b47f0ecd9857fff0ba744f662db6bbd086fd)
- [service.containers.liveness_check](resources--workload--reference--group-015.md#canonical-9c5d1446ca9c13b82f49b3937282cdf56a75392faefe1cb909e2d3913af0a08f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0f1dcb4ab49708acf1d31bd51816b47f0ecd9857fff0ba744f662db6bbd086fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2d847d1d91d3e60bc3a36eacda8ee5daa4fc285df834e88dbaa49a58730e87bb"></a>

## service.containers.liveness_check.tcp_health_check.port — service.containers.liveness_check.tcp_health_check.port / 62b39d5d99d3 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.liveness_check](resources--workload--reference--group-015.md#canonical-9c5d1446ca9c13b82f49b3937282cdf56a75392faefe1cb909e2d3913af0a08f)
- [service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-015.md#canonical-934f46745896a0bb292e97f47aafff29b13a98a8f215b386b9c05fb5a58da8a2)
- service.containers.liveness_check.tcp_health_check.port

<a id="canonical-b661ec7a874c04e7401c043678fa3edc3d4a11e201e75819fd6eb00bbddfd277"></a>

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

<a id="canonical-a5415869553a9a3e7fe7b081aae8b87f34594d8af7d7f6669068b620ff6000a4"></a>

## Direct properties — service.containers.liveness_check.tcp_health_check.port / 62b39d5d99d3 / 3

<a id="canonical-7fd15ae9ccf17950b753e6ad685c2bbe009a0160f6e77c898ec4f55f72a50763"></a>

<a id="canonical-526fd20f933a164e077a1a09765bb7af6196b02212cfa61a70e1bd33c56a4333"></a>

## name property — service.containers.liveness_check.tcp_health_check.port / 62b39d5d99d3 / 4

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

<a id="canonical-9011be58d7d396d0efa260de024c4cf9a1a43f3abdb59d8284c1a2c937580e1a"></a>

<a id="canonical-1b7e81ac5830e71754843bdd9dcc515df328ac8c0bff7ce440f56383939edbb2"></a>

## num property — service.containers.liveness_check.tcp_health_check.port / 62b39d5d99d3 / 5

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

<a id="canonical-790bfd799ae9be186246c5d90713cc5b20f6bcc1a4d0a31c033291e2ed51ce50"></a>

## Next pages — service.containers.liveness_check.tcp_health_check.port / 62b39d5d99d3 / 6

- [service.containers.liveness_check.tcp_health_check](resources--workload--reference--group-015.md#canonical-934f46745896a0bb292e97f47aafff29b13a98a8f215b386b9c05fb5a58da8a2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8c84db60fef3270c2e5c9acf436cd213cc7de6870d6207184ca297b292e72b6f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c9ac96d17bc7d8207c9e48ef1860682a0907e37a5dbe82eef192b6b9411a498"></a>

## service.containers.readiness_check — service.containers.readiness_check / 782a8553c46b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- service.containers.readiness_check

<a id="canonical-c2abae6672fbbf1529ee31ecee96a3ae442ec99aa8a241e9169f3c5c4e30f946"></a>

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

<a id="canonical-e8cca8db1319349f1453202056ecb7aa254757ad899545526ea325187e036575"></a>

## Direct properties — service.containers.readiness_check / 782a8553c46b / 3

- [exec_health_check](resources--workload--reference--group-015.md#canonical-9998862e8fc4930913796443e9091bd6ae5fde205aa0cd6927553d942885a311): complete subsection reference.

<a id="canonical-f3528fe1b7da8d67915508c457f4ed6bb57016c1dfd26fc9304e9ade5c27732f"></a>

<a id="canonical-c4b0514be5a461c7bc583fbced10ffab79d93630d9c19460ac8731c4a3b99f81"></a>

## healthy_threshold property — service.containers.readiness_check / 782a8553c46b / 4

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

- [http_health_check](resources--workload--reference--group-015.md#canonical-82b719c7307729eec6e569e822b334ccfeabf207f05b35e5f10acdfb5d4d5896): complete subsection reference.

<a id="canonical-6b8793c0071a0992e28e68a7efdbdadba40d9530e42140615c267bad8f6a259c"></a>

<a id="canonical-0034c33de1e143312b9cdef0d6705dedfb9ca9db4858bd5b4ec7597ddb21588e"></a>

## initial_delay property — service.containers.readiness_check / 782a8553c46b / 5

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

<a id="canonical-1a01b5c26a6e06ea235cddb14d231d7ac06d50f21f60b98fea70775bbfb7055a"></a>

<a id="canonical-0e1ea01fa8cfe66a21d750d3e08d0bcba8b74e2f33c859dc6793664b4a9c1a56"></a>

## interval property — service.containers.readiness_check / 782a8553c46b / 6

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

- [tcp_health_check](resources--workload--reference--group-016.md#canonical-bae7e7156dfc3765696e12988318f708df2763faadc27866633aa63112a9ba7b): complete subsection reference.

<a id="canonical-b8224f719f07918b565eeba64a5022f868df512edb33aad53e49d42b1cb1281d"></a>

<a id="canonical-072a75b5d3f372a5d728c0ebf1c45a6e93f758255848c2612bcfc9f29c18e526"></a>

## timeout property — service.containers.readiness_check / 782a8553c46b / 7

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

<a id="canonical-4dd117bf0dd7a0e9da35eae0b332d201f89d2571071b37d7a507e66472c5bfe0"></a>

<a id="canonical-4545950567e35cede70aa9741b0ed90454f912c96357de8c51cf871bbf6bd678"></a>

## unhealthy_threshold property — service.containers.readiness_check / 782a8553c46b / 8

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

<a id="canonical-5351290aa7a5de14851231543085ec71e8bb7d40b3780c5cbd37e962830250a7"></a>

## Next pages — service.containers.readiness_check / 782a8553c46b / 9

- [service.containers.readiness_check.exec_health_check](resources--workload--reference--group-015.md#canonical-9998862e8fc4930913796443e9091bd6ae5fde205aa0cd6927553d942885a311)
- [service.containers.readiness_check.http_health_check](resources--workload--reference--group-015.md#canonical-82b719c7307729eec6e569e822b334ccfeabf207f05b35e5f10acdfb5d4d5896)
- [service.containers.readiness_check.tcp_health_check](resources--workload--reference--group-016.md#canonical-bae7e7156dfc3765696e12988318f708df2763faadc27866633aa63112a9ba7b)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9998862e8fc4930913796443e9091bd6ae5fde205aa0cd6927553d942885a311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-90ce2e3dc11705bd30ebc1c83627c1be689f859f3341673396f57895232e0714"></a>

## service.containers.readiness_check.exec_health_check — service.containers.readiness_check.exec_health_check / f27db93f2b19 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.containers](resources--workload--reference--group-015.md#canonical-3f88422b353a65363a1ce1b0d566b8c6aa769a87c87c23cb23b8978323a9c01a)
- [service.containers.readiness_check](resources--workload--reference--group-015.md#canonical-8c84db60fef3270c2e5c9acf436cd213cc7de6870d6207184ca297b292e72b6f)
- service.containers.readiness_check.exec_health_check

<a id="canonical-79459b9953d98d23a6c64516d432feba9076cdd84d3d8e57f7df1a510c070a8d"></a>

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

<a id="canonical-3a527af72aca8e0049532056a3b2317b43af85470d4b330b3401ba8a23f7438b"></a>

## Direct properties — service.containers.readiness_check.exec_health_check / f27db93f2b19 / 3

<a id="canonical-9edc4d3c88634c609a6b11055ecbdd6e29c26988934d5c4610c64478b829d063"></a>

<a id="canonical-df442bc6f8931a781d0345c98d5b5a0d2d7f9da6c6cd625466d6eb7c991a4c2f"></a>

## command property — service.containers.readiness_check.exec_health_check / f27db93f2b19 / 4

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

<a id="canonical-7998d21e7531b6ee6bb617d72c563c50d96cd2f6f2eff85dd09fa38331b70741"></a>

## Next pages — service.containers.readiness_check.exec_health_check / f27db93f2b19 / 5

- [service.containers.readiness_check](resources--workload--reference--group-015.md#canonical-8c84db60fef3270c2e5c9acf436cd213cc7de6870d6207184ca297b292e72b6f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-82b719c7307729eec6e569e822b334ccfeabf207f05b35e5f10acdfb5d4d5896"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
