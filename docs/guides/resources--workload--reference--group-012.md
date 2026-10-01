---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-8d63414a1d25cd119c3dc2d2ab4d793c377d8514f475db0b474ec0076da3c2fa"></a>

## response_code property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 10092f01297b / 9

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

- [retain_all_params](resources--workload--reference--group-012.md#canonical-a2a6923c85bb635f1fb685e08b213edbc6ffd7819d0eb165fd8f6b56a3424296): complete subsection reference.

<a id="canonical-2dd9cc93c471fba0d70de76484b5d7e7dc6904e3d13a1fe8e96a5fee229ef4ae"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 10092f01297b / 10

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](resources--workload--reference--group-012.md#canonical-e711c16faa44fe34d6926520ad159c6d4c3323003fc606f67a01ddf11d76b845)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](resources--workload--reference--group-012.md#canonical-a2a6923c85bb635f1fb685e08b213edbc6ffd7819d0eb165fd8f6b56a3424296)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-ee89c7ffef10a4869121a4a0f4182f23d9694b878b60d29ddab3fb19b3a6f5c2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e711c16faa44fe34d6926520ad159c6d4c3323003fc606f67a01ddf11d76b845"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5440af8bf33fa4b393a2f23fd0c6cc19d4eb35b38adc18049af615b5dd414a2e"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 231d5c677ff5 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-ffbd01fdba16046eb807c00d809a0c34e45fc346864ea2783a5affb6fd62c4a0)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-af469861294cf846760e101cfca19af63a7546ef5c387fc036fc2fc27a75c80a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-ee89c7ffef10a4869121a4a0f4182f23d9694b878b60d29ddab3fb19b3a6f5c2)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-011.md#canonical-8135137038909f79f9ef03c2eb431a534fb1387ad0e08e8c4e5d338bf3a1e029)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-1f49e963f70f812df76142b6cfba92637ac7ed60139c8a9eae8eb8375931bba3"></a>

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

<a id="canonical-3cbb1689291b13a8527a0c6008990a7dc208db860f08986d867d5315d7f2a527"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 231d5c677ff5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3efb07c24fdb6a6e7982fbf3aed016a9c5486ff95d6fef30ceec72304428cdad"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 231d5c677ff5 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-011.md#canonical-8135137038909f79f9ef03c2eb431a534fb1387ad0e08e8c4e5d338bf3a1e029)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a2a6923c85bb635f1fb685e08b213edbc6ffd7819d0eb165fd8f6b56a3424296"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95a1020bedc4d003827f49632df3d8ec40bbb3aa68a5166c2a455f285216816e"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 42f13bde010e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-ffbd01fdba16046eb807c00d809a0c34e45fc346864ea2783a5affb6fd62c4a0)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-af469861294cf846760e101cfca19af63a7546ef5c387fc036fc2fc27a75c80a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-011.md#canonical-ee89c7ffef10a4869121a4a0f4182f23d9694b878b60d29ddab3fb19b3a6f5c2)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-011.md#canonical-8135137038909f79f9ef03c2eb431a534fb1387ad0e08e8c4e5d338bf3a1e029)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-4b695afb0c90ebd55ebb145ca5e936a3f841a12521b48680f84f8e0598578eed"></a>

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

<a id="canonical-bc1306d0dc4a90433c68e5400ab1268142852b375407466c671fe1ac55b53bfc"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 42f13bde010e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1e27c1e01834b1b7118e42cfda41e5791e93a79337de401f5c63fa5615be4672"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 42f13bde010e / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-011.md#canonical-8135137038909f79f9ef03c2eb431a534fb1387ad0e08e8c4e5d338bf3a1e029)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-468f9d8e6407a14364d6dbc8d4c0fc0a77a55bccff3f091f6cdf59696bf88adb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-137e2cdf4786daafe40b1a0d0cc2c390ebbe6421b0bc70863799a662bfa96b78"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 36a5fca8e457 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-ffbd01fdba16046eb807c00d809a0c34e45fc346864ea2783a5affb6fd62c4a0)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-af469861294cf846760e101cfca19af63a7546ef5c387fc036fc2fc27a75c80a)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-907045016ccdb4ddfeeb426fd9a26025c6d0e081856cc043eaa51d8373c6d296"></a>

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

<a id="canonical-84a4619632da8ebafc79b0b1d3808a25a52484d20adffc49848f87e2b9368893"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 36a5fca8e457 / 3

- [auto_host_rewrite](resources--workload--reference--group-012.md#canonical-0e9122c68b6da2bab7769395f2ab7e3e3febd5a139456a5f82d1d5aab9272a19): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-012.md#canonical-50aaf50c2d436f00f3d0763c07ef04f9932dbe5afcbfb2f2644ac6c2c1abf382): complete subsection reference.

<a id="canonical-16cbda8299cd69183bc10f6246986214a9fe91b31d7f0abb4cbaa38385c1df42"></a>

<a id="canonical-c56a270f18bde114984d58ac0db521b4919fb3c9ed4916d985f5c9a5f2d21ab0"></a>

## host_rewrite property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 36a5fca8e457 / 4

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

<a id="canonical-b4d1ce80510c3f10f5b2f796a4b981e73448ea6340dabe5efdee96db2fc6cc0d"></a>

<a id="canonical-a4c47ed55e798eaa3b26441fd32768ac3709cee5de4a6363644d716e7eb32b97"></a>

## http_method property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 36a5fca8e457 / 5

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

- [path](resources--workload--reference--group-012.md#canonical-bdaf2e3b2b4632e80abef8df053772271a7b8e45126a83c1aec8739d97f6edfd): complete subsection reference.

<a id="canonical-a89cb3b980d1b36446ecad11b7b18f13dc4ee8c3979c83c08c6579f36f6f1d97"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 36a5fca8e457 / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](resources--workload--reference--group-012.md#canonical-0e9122c68b6da2bab7769395f2ab7e3e3febd5a139456a5f82d1d5aab9272a19)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](resources--workload--reference--group-012.md#canonical-50aaf50c2d436f00f3d0763c07ef04f9932dbe5afcbfb2f2644ac6c2c1abf382)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path](resources--workload--reference--group-012.md#canonical-bdaf2e3b2b4632e80abef8df053772271a7b8e45126a83c1aec8739d97f6edfd)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-af469861294cf846760e101cfca19af63a7546ef5c387fc036fc2fc27a75c80a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0e9122c68b6da2bab7769395f2ab7e3e3febd5a139456a5f82d1d5aab9272a19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b52f17f982d41945179db35ade2566ff86cc13715e5a4cbf4378047e788ba573"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 162f7ec55cb8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-ffbd01fdba16046eb807c00d809a0c34e45fc346864ea2783a5affb6fd62c4a0)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-af469861294cf846760e101cfca19af63a7546ef5c387fc036fc2fc27a75c80a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-468f9d8e6407a14364d6dbc8d4c0fc0a77a55bccff3f091f6cdf59696bf88adb)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-8f4a97b11febed34b72e4a5a1ab1f6a7a6b15712561ff4ea18c64b6745f21a6f"></a>

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

<a id="canonical-d26b00a21f18aea7492c46dd8e1f08f0afa1f5f6505dac4f77a333f3f8c2f1dd"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 162f7ec55cb8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-53a514ec84dd7d1d603b524080a15f48a26bf286002456e43b29d7f4acf5e010"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 162f7ec55cb8 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-468f9d8e6407a14364d6dbc8d4c0fc0a77a55bccff3f091f6cdf59696bf88adb)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-50aaf50c2d436f00f3d0763c07ef04f9932dbe5afcbfb2f2644ac6c2c1abf382"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b0f7f4b7b6566470ed2d549e7c505628c223fb23afa7b0afe780141822d14328"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 46ab280a1b10 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-ffbd01fdba16046eb807c00d809a0c34e45fc346864ea2783a5affb6fd62c4a0)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-af469861294cf846760e101cfca19af63a7546ef5c387fc036fc2fc27a75c80a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-468f9d8e6407a14364d6dbc8d4c0fc0a77a55bccff3f091f6cdf59696bf88adb)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-72b763060c92ff5e008f87118c3b459fbfac9957587673fe35c6b163cdb112cc"></a>

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

<a id="canonical-ac9772047ef95e87e9cc779ff199110358530e1043002df8459c2d4a4d9d26a8"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 46ab280a1b10 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-839a5d10762e15b85fe6161a2013fc76aa08919d4a829f6a40d3bac67236fc13"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / 46ab280a1b10 / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-468f9d8e6407a14364d6dbc8d4c0fc0a77a55bccff3f091f6cdf59696bf88adb)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-bdaf2e3b2b4632e80abef8df053772271a7b8e45126a83c1aec8739d97f6edfd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6baab06216daa0bb9c9232863f66d0b24c586f008048429b4d6be0bd0d052bf3"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / cb173964d8db / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer](resources--workload--reference--group-008.md#canonical-6221b23fbec385fc79597f6296a01651ca047d39084019900ebb81bb756a6d0d)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-011.md#canonical-ffbd01fdba16046eb807c00d809a0c34e45fc346864ea2783a5affb6fd62c4a0)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-011.md#canonical-af469861294cf846760e101cfca19af63a7546ef5c387fc036fc2fc27a75c80a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-468f9d8e6407a14364d6dbc8d4c0fc0a77a55bccff3f091f6cdf59696bf88adb)
- service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-d507d0bfa715cc9aa0054a423d9efaaf2f6a26dd2766998978ac0e9d99944962"></a>

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

<a id="canonical-0d079ff7126dfd5ef9d37abe3c6c25e969b709c481720674073acb6faed81e25"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / cb173964d8db / 3

<a id="canonical-d5caea4bc70181e5d1ee6f8191a779747f54d93639c5e5a99adf94c0ff8ea948"></a>

<a id="canonical-4149dce9bf0c6afb3263542a50ef3d93682e62f7b9ed919b17e2c63f47ccec71"></a>

## path property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / cb173964d8db / 4

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

<a id="canonical-f89edd64f838086e48031808d645123e1108a1e2c42a83d615fd5d0d6dda95a6"></a>

<a id="canonical-c1d717fe261157865e9e77cce55acd9a587683d641d65d8126d3bd7e42fb370f"></a>

## prefix property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / cb173964d8db / 5

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

<a id="canonical-dd4d25828738ce378630a48f73c56a1b40e6a9d29633f103b5331d83a0a6c1f8"></a>

<a id="canonical-40324cf94889e6447eba7ac45392183c93f3aef7fb047eb2e7cb67b6040914e0"></a>

## regex property — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / cb173964d8db / 6

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

<a id="canonical-35621b6fe371b1810b09b9e7d841c3e9bd141507714bf71c839d103b3ee42031"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalance / cb173964d8db / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-012.md#canonical-468f9d8e6407a14364d6dbc8d4c0fc0a77a55bccff3f091f6cdf59696bf88adb)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3df6a48ca312dd6a53e3d784178510d1f1840fec69978ddd52219d094c927f4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-160ec11e4972722159544e3a6550b782541f1306b15306711b22ff9a734244b1"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.port — service.advertise_options.advertise_on_public.multi_ports.ports.port / 1af4bf7b4acf / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- service.advertise_options.advertise_on_public.multi_ports.ports.port

<a id="canonical-c133f55e272e8e791ed2634661dc66bd358169634628dc4c712efc1c70102b28"></a>

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

<a id="canonical-e4ecd5944cc4f6ec7e7409d789ce3f0041bc05d171321b24c3b11bd3840c3f61"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.port / 1af4bf7b4acf / 3

- [info](resources--workload--reference--group-012.md#canonical-0ab0f93a4c18c17d85e13339b732901b121a9fba5a566b6747379337090925bb): complete subsection reference.

<a id="canonical-80db78ffdd7d017a28c669ce71ecc3640c56475bc52c3d8c04ccd66af046ca4b"></a>

<a id="canonical-a91394e516075524b9c2bdf6ff14442379c294bfd8583eed008457fb97e7ab09"></a>

## name property — service.advertise_options.advertise_on_public.multi_ports.ports.port / 1af4bf7b4acf / 4

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

<a id="canonical-4f1ece0075880cf28447254071ed59f75dbf69f94c5cd4d8a4180c44bdec4956"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.port / 1af4bf7b4acf / 5

- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-012.md#canonical-0ab0f93a4c18c17d85e13339b732901b121a9fba5a566b6747379337090925bb)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0ab0f93a4c18c17d85e13339b732901b121a9fba5a566b6747379337090925bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f3a91a84d1958dbed02371909a2e9f230838004266b9683c8c0e2b6d364cff62"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.port.info — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / 90bb682505dc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-012.md#canonical-3df6a48ca312dd6a53e3d784178510d1f1840fec69978ddd52219d094c927f4c)
- service.advertise_options.advertise_on_public.multi_ports.ports.port.info

<a id="canonical-2acf76e0ce25e58044869cfb1f78bbbd76487672719f1796849cbb738eb21473"></a>

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

<a id="canonical-20fa7b0f0658dabfa77ae5d058bc4c09faec13a4066f0127869e922e202534d2"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / 90bb682505dc / 3

<a id="canonical-88d08056e932a3a23308bc6c971f7a4d51b8d8968784951edfcccdba6cc31a90"></a>

<a id="canonical-cfd392ae7d7dd1d7eed549128156b55194d8c8d326e7239422077a23a57dd499"></a>

## port property — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / 90bb682505dc / 4

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

<a id="canonical-367884b44fd0b89f89f905ac3657a5b2ba899a88a911e66d81243fa768c5db64"></a>

<a id="canonical-386a4990cc9170b5148066747f0578df3fc59c6592da01cc041801467be54386"></a>

## protocol property — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / 90bb682505dc / 5

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

- [same_as_port](resources--workload--reference--group-012.md#canonical-32c687b176573052c340cc859b100c0218edfe5be9586e14f2b6450b1bb66705): complete subsection reference.

<a id="canonical-818e28d8279e151bc5544cb4a0d7a03245f794076e8a859bdd0133e7dfff673e"></a>

<a id="canonical-99fd5cd88d9fc98d7db8ffe59862db821a3d4af6e4e80007ec8b513295323a9a"></a>

## target_port property — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / 90bb682505dc / 6

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

<a id="canonical-206de04ae70b2827af9aad00ba8b145c7959872547e969bcbc302942874ea5a6"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.port.info / 90bb682505dc / 7

- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port](resources--workload--reference--group-012.md#canonical-32c687b176573052c340cc859b100c0218edfe5be9586e14f2b6450b1bb66705)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-012.md#canonical-3df6a48ca312dd6a53e3d784178510d1f1840fec69978ddd52219d094c927f4c)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-32c687b176573052c340cc859b100c0218edfe5be9586e14f2b6450b1bb66705"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83fa018f1787787f96f2441add0a2cb655f054ebcfad42fe2ebd479c4dc21f29"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port — service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_a / b01d1dbd2c2c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port](resources--workload--reference--group-012.md#canonical-3df6a48ca312dd6a53e3d784178510d1f1840fec69978ddd52219d094c927f4c)
- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-012.md#canonical-0ab0f93a4c18c17d85e13339b732901b121a9fba5a566b6747379337090925bb)
- service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_as_port

<a id="canonical-5ae967ab1fbd9db79ad96b9edd3ca41bdc18a6b3fd534093a47286f45310dbbc"></a>

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

<a id="canonical-95b80e411d8911ba869472e0e5eef4ed2e97a3041314702fd24ad51468a2f117"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_a / b01d1dbd2c2c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7e4a57282210c9b70b27d87d7cf6e926f91619afaebc7eaea1a845ac4a58ddce"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.port.info.same_a / b01d1dbd2c2c / 4

- [service.advertise_options.advertise_on_public.multi_ports.ports.port.info](resources--workload--reference--group-012.md#canonical-0ab0f93a4c18c17d85e13339b732901b121a9fba5a566b6747379337090925bb)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6102c34eb63c02fd55faa6d0818d5cad1f48e01c996b8475abd485db68a82de8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-271e6c0775112eda8372eb2cf45801bf378a39d677a283204274615b43e6445d"></a>

## service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer — service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer / d571c1f58c4f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.multi_ports](resources--workload--reference--group-008.md#canonical-e4497353be851a10536a90a6c24460407e35c08509b5fcda99aa69a75be28123)
- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer

<a id="canonical-f17ed529826661d7b1b12e216dbf27b0a14633e44ac16890aa052548f4f2fa9a"></a>

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

<a id="canonical-08ad85546bcbcb21ad543bbb9af963116fbbbf3981ee590a366c44874d294f58"></a>

## Direct properties — service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer / d571c1f58c4f / 3

<a id="canonical-eb4d8f949079c240c476476069058628af7a0a1733622c2730612718e8bcdf73"></a>

<a id="canonical-ffe394ed31b5f8f63b80c8c0665ac7d2a4f40dd7641aabc5b7e113f2a1da76c7"></a>

## domains property — service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer / d571c1f58c4f / 4

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

<a id="canonical-324494d3f9851eede6631ce9669d480f3d78485d57fbffe77386f302a4a59d45"></a>

<a id="canonical-42c08a00b3c8fd1eb43616a83cb49a78918b857679555c994d8b26f32c575651"></a>

## with_sni property — service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer / d571c1f58c4f / 5

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

<a id="canonical-2f6c76773663937931b376f8bbc80d18dae79702a26d37bda2d2e91300806466"></a>

## Next pages — service.advertise_options.advertise_on_public.multi_ports.ports.tcp_loadbalancer / d571c1f58c4f / 6

- [service.advertise_options.advertise_on_public.multi_ports.ports](resources--workload--reference--group-008.md#canonical-b57af1182bb2c6a43e823da4c079543f5f471bc4e5d0f6bfbb0949007018251a)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52fa465f9b07a5c9a7aa520b71d4de8b40aa6e1e340ed5a54571ed55f19fefba"></a>

## service.advertise_options.advertise_on_public.port — service.advertise_options.advertise_on_public.port / 5d6e7183a57d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- service.advertise_options.advertise_on_public.port

<a id="canonical-58a1e8f56fb1477a465e70473146029afbc9a9a8d6204d5978369e7c9e2dbf32"></a>

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

<a id="canonical-21b84b577a1d4c2d01b54d44c91780a48939e0909af66a3ef126609bbbe68532"></a>

## Direct properties — service.advertise_options.advertise_on_public.port / 5d6e7183a57d / 3

- [http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc): complete subsection reference.

- [port](resources--workload--reference--group-015.md#canonical-11e66f627d5d7a33e10e8d9045af1ca879560472658dea83732f50de97228951): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-015.md#canonical-427355fe9cb5b0c0971a8e939299dc09ef6e767ec08e74cbbab5065ec1526dca): complete subsection reference.

<a id="canonical-2f58bfd69da0b1fc5565f938669c7aee89151733539a1a757dbbb0f8d1a65713"></a>

## Next pages — service.advertise_options.advertise_on_public.port / 5d6e7183a57d / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.port](resources--workload--reference--group-015.md#canonical-11e66f627d5d7a33e10e8d9045af1ca879560472658dea83732f50de97228951)
- [service.advertise_options.advertise_on_public.port.tcp_loadbalancer](resources--workload--reference--group-015.md#canonical-427355fe9cb5b0c0971a8e939299dc09ef6e767ec08e74cbbab5065ec1526dca)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c73d35652f123cb2d1e25716f87ff130ee4eb4acf6f0510ca3aabf9eaea36906"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer — service.advertise_options.advertise_on_public.port.http_loadbalancer / b9baeec8bced / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- service.advertise_options.advertise_on_public.port.http_loadbalancer

<a id="canonical-6a208d4dac4c097c4743a28d32ff3ffc9c0c0ef8edfe753f2a7dbf0f51b45750"></a>

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

<a id="canonical-6b1312067ea912267e39f72d3973ec136f58498777d72317531df367364d8aa2"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer / b9baeec8bced / 3

- [default_route](resources--workload--reference--group-012.md#canonical-0ad19415be35a8134f35266813fe7d4915c39fa28347314e8dfa3bb33f6264c7): complete subsection reference.

<a id="canonical-006144e8105dfd2c419dbe0325b990c81ecee398bc52697f56b4903a8f1798ac"></a>

<a id="canonical-fb845840a6fb91d84c55aa4ac825149b72dc09ce069abc4f28e8c8f4d08c53ee"></a>

## domains property — service.advertise_options.advertise_on_public.port.http_loadbalancer / b9baeec8bced / 4

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

- [http](resources--workload--reference--group-012.md#canonical-783171af4a3b694641a5b2ecaa8fe51508eb975b23daa155ad0ad886d21f6140): complete subsection reference.

- [https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-013.md#canonical-af097b6f6116d0cbb4580bc318179bf7da7a4923b0776b78bbdd457f4624a1cf): complete subsection reference.

- [specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435): complete subsection reference.

<a id="canonical-12b364a6015cb7e19c188e921e7b5d27299bc1de84bfaeb4413ac5d630a8f774"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer / b9baeec8bced / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-012.md#canonical-0ad19415be35a8134f35266813fe7d4915c39fa28347314e8dfa3bb33f6264c7)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.http](resources--workload--reference--group-012.md#canonical-783171af4a3b694641a5b2ecaa8fe51508eb975b23daa155ad0ad886d21f6140)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https_auto_cert](resources--workload--reference--group-013.md#canonical-af097b6f6116d0cbb4580bc318179bf7da7a4923b0776b78bbdd457f4624a1cf)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.specific_routes](resources--workload--reference--group-014.md#canonical-6b77f771518a0983b3004d06c376980821adc8fab32c1a876d40594c8ef19435)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0ad19415be35a8134f35266813fe7d4915c39fa28347314e8dfa3bb33f6264c7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c96adfaabf9d384f838d80e95d62a3805f62e2f1a230a8af6b64e6733bc43c5d"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 8cd4ca1e5e11 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route

<a id="canonical-172c88d71f8b5b36e72b31a7d08540a910cafa5295a6e231d008893dec949719"></a>

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

<a id="canonical-ae10c2bab1627311d76385195f187ffb42a0b76e2b396f33b63fe5c4418c8e91"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 8cd4ca1e5e11 / 3

- [auto_host_rewrite](resources--workload--reference--group-012.md#canonical-4f987cffbef99bceafbf34a51ec947e22d50a1cbd69c6ced7b88b80b52c5d156): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-012.md#canonical-d428ed80bad405d7784f7d7118755b07476c0a664cf08576a9755a82219391e7): complete subsection reference.

<a id="canonical-c6407a458bf61d6c90a5571bd1dc60c40fc2261c9c8a77562415c73c8532380c"></a>

<a id="canonical-002aae4e3b32f4ddd757d22fb17c6d3f0888b2b03e8ca7a1333689c136a30008"></a>

## host_rewrite property — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 8cd4ca1e5e11 / 4

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

<a id="canonical-9135acd753e2ef4b779e583c9a52250d1a702ae179435d3156665025a6cc0613"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 8cd4ca1e5e11 / 5

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite](resources--workload--reference--group-012.md#canonical-4f987cffbef99bceafbf34a51ec947e22d50a1cbd69c6ced7b88b80b52c5d156)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite](resources--workload--reference--group-012.md#canonical-d428ed80bad405d7784f7d7118755b07476c0a664cf08576a9755a82219391e7)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4f987cffbef99bceafbf34a51ec947e22d50a1cbd69c6ced7b88b80b52c5d156"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5593611771ed767a8488cfaad17e17ed63a136a6ac84ca42397f5abe18c1d81e"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 4e491896e44a / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-012.md#canonical-0ad19415be35a8134f35266813fe7d4915c39fa28347314e8dfa3bb33f6264c7)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-a787f56b72a03acc1ae8744c356c6012e6f12e9c212735e48246a068bc4af258"></a>

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

<a id="canonical-af3e89930ca55a9a2416693d53950f86af5b16fc1e4202d6a25c4697d0cceb1a"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 4e491896e44a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bd48bc02c798c81b38a67bc993ee4c867474a2dd0cf5143d49c0e6ec568c0454"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / 4e491896e44a / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-012.md#canonical-0ad19415be35a8134f35266813fe7d4915c39fa28347314e8dfa3bb33f6264c7)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d428ed80bad405d7784f7d7118755b07476c0a664cf08576a9755a82219391e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c0650bb25cddb77c2f1dd980c900996ea0d173736a62e1af10cc42de0ac2702"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / f2b3408a7cc5 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-012.md#canonical-0ad19415be35a8134f35266813fe7d4915c39fa28347314e8dfa3bb33f6264c7)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-f21894b085f0ecbb9007faf377c5641ed9eee64876514acd936f44b220222fa0"></a>

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

<a id="canonical-e817e6c3f6416a0de8ded770010443b55b8e62e0f4b2f3173fe7d4d6ddc995a2"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / f2b3408a7cc5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-56738b1c8d248e08a23f35d40411c30fa19586dbe5dcaf99c42987b35e23a0df"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.default_rou / f2b3408a7cc5 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.default_route](resources--workload--reference--group-012.md#canonical-0ad19415be35a8134f35266813fe7d4915c39fa28347314e8dfa3bb33f6264c7)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-783171af4a3b694641a5b2ecaa8fe51508eb975b23daa155ad0ad886d21f6140"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2930517d4371a2ada9c7fcbeb26960c7bf9c5703e8890577f5145dc24d3702f"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.http — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 904d18296b05 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.http

<a id="canonical-0deeb0a2bda4ba34ec45a10287b7ba8249ba7b3aa0b87844df77b6af0eb63c7c"></a>

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

<a id="canonical-326ae75ce27a818d42f1a3123694258e51bbbd2efbf13f434ac58260fa946eb1"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 904d18296b05 / 3

<a id="canonical-8a53d38bfa8f4da285a3b27d48861a97593b8013f9bbddfba47ef8ed896288fb"></a>

<a id="canonical-be8147cc2e7be8d90e9e908e562cf0e912025e4cb33180c319601def41109bbd"></a>

## dns_volterra_managed property — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 904d18296b05 / 4

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

<a id="canonical-e02e79ce37faaa95718eb8fcae4c5cbb10ffcbd7c125080da3beeb9757a2fa0b"></a>

<a id="canonical-43a37fc29f79be9a604ec92e0eab76e6f16a5485af61718475420e29fe737e2e"></a>

## port property — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 904d18296b05 / 5

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

<a id="canonical-3bc18d59737a84e34cb78b42e3c37ec22fe778555c112818039450ef06163627"></a>

<a id="canonical-6ec8947d2b1794f0a6b4dafd68388cce8cce32acf8daeea47f4007855d72fb72"></a>

## port_ranges property — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 904d18296b05 / 6

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

<a id="canonical-a2ccfd3224623570ae567a4d205c457b4525ac607b0ea4460bcc144031525e27"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.http / 904d18296b05 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-70dbf42078f9f026f5dcb7867698003e4f9fa0920a719d32ec2b3459de9265db"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 6d7c939bfc7f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https

<a id="canonical-bb302bef22cfc6e67631837da6734a11741ff3cee7fc212c3c6896e97b5ffc9d"></a>

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

<a id="canonical-e86c03cf78bf5e831589ef6bb669be90c93de70dde1cae1e66b0a33b73a31aea"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 6d7c939bfc7f / 3

<a id="canonical-2618bfc074e54ee4504fc77cfa4c4306ab975d998889f70a4682ae9a93daf3a2"></a>

<a id="canonical-652bc5f77a1c2c22f265ba3845a8245d12a5c53756cfa28de5238f635ca50995"></a>

## add_hsts property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 6d7c939bfc7f / 4

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

<a id="canonical-13bc4f93dc6b5db525090dad176aebd4064f6ebd1d059c1c96b1c3ae6476553f"></a>

<a id="canonical-e4d1762fcb4abf3687580eac8882f9782c120cd6ac49f18bd9a1ba70f2bcea10"></a>

## append_server_name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 6d7c939bfc7f / 5

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

- [coalescing_options](resources--workload--reference--group-012.md#canonical-dd2634313bcd37a60c5d729394a167fbc298dd8fb2ed7ae98a105cf3758e62a8): complete subsection reference.

<a id="canonical-bdff170b224a94be5659c805af5dfb27687e4609fdbcc2874294aba4a1418c77"></a>

<a id="canonical-a3c508d39057d21a6b1f9f432bb94d9ab4624314cbcdf9ada53b0a59cfb7c4f7"></a>

## connection_idle_timeout property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 6d7c939bfc7f / 6

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

- [default_header](resources--workload--reference--group-012.md#canonical-f9d197072a6b7874a7b7aae5e04289447428a87e114d85e429e98862b0628d31): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-012.md#canonical-787b7d10f9ac85e11c39b2535dbbeb34ec9e217de1d570d9e50c71574f9dd2c1): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-012.md#canonical-a9d760b070c192aa2ee6578b35f8857e0a797c40cea7309a5bdeda6d605d817d): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-012.md#canonical-ef8f3579ede963d1a983452b62dda9926bf6d8a99969d47d44de218aa348ea4b): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b): complete subsection reference.

<a id="canonical-c1272ed559526363bfa5ff841b6aeff8364f9ed7fbe0683ee601bfc97bc2a79e"></a>

<a id="canonical-4311b740eecb96dc5c092151e5219fb8795cea4e153a3a3ee8fb1dc8d2d54fb2"></a>

## http_redirect property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 6d7c939bfc7f / 7

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

- [non_default_loadbalancer](resources--workload--reference--group-012.md#canonical-21192f9ac919ff00667d635dc054bfc757937509f37f2dbdf3987c6a0ec5c0ee): complete subsection reference.

- [pass_through](resources--workload--reference--group-012.md#canonical-7db59adac8ffdb78aa082ec701a37ab2eb64663e47676515049945c720c48ade): complete subsection reference.

<a id="canonical-69d862579719563b156b8ed5fa731fae8706a3df5052cbf7090f64ef6de95be4"></a>

<a id="canonical-138b9cfbe9990bcf01cadafa633434492d39b36181bbf14e1cb0c0655f2cd7a5"></a>

## port property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 6d7c939bfc7f / 8

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

<a id="canonical-0813837bb849f79fc559034c7d9fb026082d1969fd9c4549f0026a5809c0d670"></a>

<a id="canonical-7658a4e8a2a140e391fb984f7fd875159775900753730f009f6f0024bb3c5e8d"></a>

## port_ranges property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 6d7c939bfc7f / 9

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

<a id="canonical-64ef03e6a264cd99e8735dc16a407bcc335bc331fc723aebfe6451c2ecd52ae7"></a>

<a id="canonical-0a731519e51798b527715ebbcd1ed47cbebd8e231949a98c3b98ba2ce741c7bd"></a>

## server_name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 6d7c939bfc7f / 10

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

- [tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef): complete subsection reference.

- [tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334): complete subsection reference.

<a id="canonical-faa22cd892e8480d0518c936e467969a6f87f9150e12c5aab18e7e2397d3d9c2"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https / 6d7c939bfc7f / 11

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-012.md#canonical-dd2634313bcd37a60c5d729394a167fbc298dd8fb2ed7ae98a105cf3758e62a8)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header](resources--workload--reference--group-012.md#canonical-f9d197072a6b7874a7b7aae5e04289447428a87e114d85e429e98862b0628d31)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer](resources--workload--reference--group-012.md#canonical-787b7d10f9ac85e11c39b2535dbbeb34ec9e217de1d570d9e50c71574f9dd2c1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize](resources--workload--reference--group-012.md#canonical-a9d760b070c192aa2ee6578b35f8857e0a797c40cea7309a5bdeda6d605d817d)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize](resources--workload--reference--group-012.md#canonical-ef8f3579ede963d1a983452b62dda9926bf6d8a99969d47d44de218aa348ea4b)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer](resources--workload--reference--group-012.md#canonical-21192f9ac919ff00667d635dc054bfc757937509f37f2dbdf3987c6a0ec5c0ee)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through](resources--workload--reference--group-012.md#canonical-7db59adac8ffdb78aa082ec701a37ab2eb64663e47676515049945c720c48ade)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-013.md#canonical-6d0d87d68751c425af6449419937a1eeaea337ed3042f834031b1605d2915334)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-dd2634313bcd37a60c5d729394a167fbc298dd8fb2ed7ae98a105cf3758e62a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed28d9299eac71787e40b090deeadf3e488477724c5f9900fcfff927224923d1"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 0c2aeb6f1e70 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options

<a id="canonical-9a65c66339809a6f2bf5e74c48e85d39cf0354b5f752572202e22b1515994c5a"></a>

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

<a id="canonical-4cab3f0f4d6bc5d5028a03bc2b35e96024b4fef49d26994fb8eb34a81effd770"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 0c2aeb6f1e70 / 3

- [default_coalescing](resources--workload--reference--group-012.md#canonical-d2c6c25d2f6fc6edcffb95d990ecc6cb5130d25df5a6fd253ee2634fce419659): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-012.md#canonical-a6f75225c5575c06c69299cb133e0967923a98791b2e658bfb035cf810b053c5): complete subsection reference.

<a id="canonical-127e9a23984fb4d36a1ff52bd69a7cf48484c1854880c06b09cc71e07bc3850c"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 0c2aeb6f1e70 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing](resources--workload--reference--group-012.md#canonical-d2c6c25d2f6fc6edcffb95d990ecc6cb5130d25df5a6fd253ee2634fce419659)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing](resources--workload--reference--group-012.md#canonical-a6f75225c5575c06c69299cb133e0967923a98791b2e658bfb035cf810b053c5)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-d2c6c25d2f6fc6edcffb95d990ecc6cb5130d25df5a6fd253ee2634fce419659"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d1e45f54105edc06ba4e6dfa907b2cf44c2152769513e9b0cd51f4ad61b9184"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 298c38a4ab11 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-012.md#canonical-dd2634313bcd37a60c5d729394a167fbc298dd8fb2ed7ae98a105cf3758e62a8)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-91c1bd729905df85bba03c6ab163e0e12f6ace741ac2ffc0ded710ada77193d6"></a>

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

<a id="canonical-768c59abbf96c665c0851f673663b57dc3f3410fe7e3e81fea1716b5596df976"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 298c38a4ab11 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a7e35b74810eb2f051f42d454eab712b4d12a74ea73717b3179b1e1608399773"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 298c38a4ab11 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-012.md#canonical-dd2634313bcd37a60c5d729394a167fbc298dd8fb2ed7ae98a105cf3758e62a8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a6f75225c5575c06c69299cb133e0967923a98791b2e658bfb035cf810b053c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be861d60f26fd441d6d9fac1df77e6a8c75c24c1a2e249e09a2269134edc03c1"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 86d70378f39b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-012.md#canonical-dd2634313bcd37a60c5d729394a167fbc298dd8fb2ed7ae98a105cf3758e62a8)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-8c5d63a0956b84ab27bb07726495f8efdee9a61f25f22ee14e89e8e6a6290284"></a>

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

<a id="canonical-0659510997b94b76f779b3e971a7020ed4e8bfe1a6c8ab24c96f56bba33d1b94"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 86d70378f39b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5b6436fa65f01539466faa3d8d200e9d6b2680df444d675c420c347edce7f722"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coale / 86d70378f39b / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.coalescing_options](resources--workload--reference--group-012.md#canonical-dd2634313bcd37a60c5d729394a167fbc298dd8fb2ed7ae98a105cf3758e62a8)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f9d197072a6b7874a7b7aae5e04289447428a87e114d85e429e98862b0628d31"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fcc32f36ed7c05d967a9b22cbf6d6d964883403d2e30d9139b22c5aff6c86fa5"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / 9427ca5660b6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_header

<a id="canonical-07b7fcee8d433f7832b7ac0e3d4154fb91387467a42932fb460fe0375779bf44"></a>

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

<a id="canonical-18f1c0fa2f46b30b7d03696c4e671646b8543a6a8d586883fd76d01b1da81d56"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / 9427ca5660b6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b54e6c7a17ee46f378e41a9a0e4456a17b5b29eee17e8396c3f9adb1bad520cb"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / 9427ca5660b6 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-787b7d10f9ac85e11c39b2535dbbeb34ec9e217de1d570d9e50c71574f9dd2c1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34d170815484522fe3cbcec2d65ea95b2bda6242e2bc33b1505be0a7a9305832"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / 3d0271570ea5 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.default_loadbalancer

<a id="canonical-01898c604e56cc244faa332cbddf6383f443bbb88df5d29e8233ccb2715af0ec"></a>

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

<a id="canonical-588ff11d6efb04a6ee26285253b2a218edfac1bd138f8ff10ce86a253939f41d"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / 3d0271570ea5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1c52ecfc8f192307a83eb2a5d8f694c3d84688688313a76f7226c8bc0c03b240"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.defau / 3d0271570ea5 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a9d760b070c192aa2ee6578b35f8857e0a797c40cea7309a5bdeda6d605d817d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-993cd1f39b7e16ab564ef979af80447f43764595a4c2b479078a3f1c5de07bba"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disab / a65d8077cb3c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disable_path_normalize

<a id="canonical-316862323e5eae14f1a3aa42dc63412a48e53e22ea616be285a9bb1541835659"></a>

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

<a id="canonical-a89999222a6fc3043c54ad0c753e7c6e323af87fd36609620cbfbeb119c04eab"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disab / a65d8077cb3c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-171a32480d6975a7a14a7e4ac2d1ca73e9edb594c6944a651910321dafba0976"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.disab / a65d8077cb3c / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ef8f3579ede963d1a983452b62dda9926bf6d8a99969d47d44de218aa348ea4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2620de459eeebcfa726362111334d7897277258f65ab55a0011a97d9b8d2e4e7"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enabl / 4aa4bb1ef834 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enable_path_normalize

<a id="canonical-f7f2acd51c65d680dfda6bbdf83c1482d3e0ffafc60df0191128a432d0cc20af"></a>

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

<a id="canonical-c13d2ad8f670292bff22e93c4fddc0a2e9989d3a1074927b3a9ad0db52347298"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enabl / 4aa4bb1ef834 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ddcf074dc792e8e1641abd7edb56412978dac9e2473c9fe12021a83aebc67814"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.enabl / 4aa4bb1ef834 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d443611c9b98cc2e4bc380e785f26e2fb028d5e26cecaac97fb6c3fa31448dfc"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / dc1f6fda0e90 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options

<a id="canonical-50f72bc779860c28d64053429963e53e5ad0b05250f13f7ea8317b7e7b26c69a"></a>

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

<a id="canonical-42e48976f4a19e97588918dedc8f8379f779272d9aa63e0a1939a41774568745"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / dc1f6fda0e90 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-012.md#canonical-b38bc9345af2dbf6ae72562dbd2795593a454abcce55a58b582f131c34a9d179): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-012.md#canonical-c1c7b4ea99b55862cf261f62878c7b5a4a7373a3d2cb0f7dba72c21465ee2e2f): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-012.md#canonical-af3f69f6e53e3a9d3cbec66675f3f23f2c9ace3c6d1065ebaca272651950da5e): complete subsection reference.

<a id="canonical-8314c086490c61b7d5929eda023793524ce54cf800ed41323af0c1011da5313a"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / dc1f6fda0e90 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-012.md#canonical-b38bc9345af2dbf6ae72562dbd2795593a454abcce55a58b582f131c34a9d179)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-012.md#canonical-c1c7b4ea99b55862cf261f62878c7b5a4a7373a3d2cb0f7dba72c21465ee2e2f)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-012.md#canonical-af3f69f6e53e3a9d3cbec66675f3f23f2c9ace3c6d1065ebaca272651950da5e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b38bc9345af2dbf6ae72562dbd2795593a454abcce55a58b582f131c34a9d179"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8da756dce114a0d743d9be90347a95ae5a9de69440a1dae6d1faabf10d1241f7"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c5bc71ad4576 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-1696ef311e0d5f1619211376b17686b463ade3c7f16b415f8c6d74f38c9603cc"></a>

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

<a id="canonical-7d5f7df002a3ef5df38f1a476dcc04f7cf1a1dbcf66f7d7ebacb123083d01e07"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c5bc71ad4576 / 3

- [header_transformation](resources--workload--reference--group-012.md#canonical-dd670faa4d2738047e1b04ae34ff99b083a94d521e786e08b6630f307e36f7e3): complete subsection reference.

<a id="canonical-0518acf6b81592e90576a42f0f6b42e963a148460704377afe889ef779912c42"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c5bc71ad4576 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-012.md#canonical-dd670faa4d2738047e1b04ae34ff99b083a94d521e786e08b6630f307e36f7e3)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-dd670faa4d2738047e1b04ae34ff99b083a94d521e786e08b6630f307e36f7e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b39d0abefc538716ddda65aa06eeec5e21aeb5573baee328454e7d921e399bf"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / fe202b40e728 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-012.md#canonical-b38bc9345af2dbf6ae72562dbd2795593a454abcce55a58b582f131c34a9d179)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-670683eb6abc947c844d6cbaefc81030e642c7da122e75cf53fb5e3f4f482b34"></a>

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

<a id="canonical-4f2041c1902fa92fa415c8662aa94da2c0be3d5b5eee67f8ba7075ca60509141"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / fe202b40e728 / 3

- [default_header_transformation](resources--workload--reference--group-012.md#canonical-6dba142cee35e8058c528f6f0aae2d4ebeb549790a0e86ce8648adef4f40c56b): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-012.md#canonical-8714d4cc6433db78da7f53e4b4f7a99140641e4881f84a4faed2af2b78d4478c): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-012.md#canonical-dee8a0435093b1e88a70585cfcf644cda2f9b5f2cc73f7e71deb8619b2153761): complete subsection reference.

<a id="canonical-6cd90fc4761ba7b504e685d2f634ec00f9a5f4b2822518232d33b7a343b0a3df"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / fe202b40e728 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-012.md#canonical-6dba142cee35e8058c528f6f0aae2d4ebeb549790a0e86ce8648adef4f40c56b)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-012.md#canonical-8714d4cc6433db78da7f53e4b4f7a99140641e4881f84a4faed2af2b78d4478c)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-012.md#canonical-dee8a0435093b1e88a70585cfcf644cda2f9b5f2cc73f7e71deb8619b2153761)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-012.md#canonical-b38bc9345af2dbf6ae72562dbd2795593a454abcce55a58b582f131c34a9d179)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6dba142cee35e8058c528f6f0aae2d4ebeb549790a0e86ce8648adef4f40c56b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d9dd65f98c6af730305808e4cee5fef053bccadbdbc0dc5eb1686d9e91ab687"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c35ee54342bc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-012.md#canonical-b38bc9345af2dbf6ae72562dbd2795593a454abcce55a58b582f131c34a9d179)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-012.md#canonical-dd670faa4d2738047e1b04ae34ff99b083a94d521e786e08b6630f307e36f7e3)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-d0dd04e0c8ba52fadca0ffc1d66b116a1b7041538e9b3bfba2873ef013adf47c"></a>

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

<a id="canonical-c94ffbea9962f97fc306a41d8e81ffdc373b7d08052bb5832c9f6147c46c3868"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c35ee54342bc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-831be6e8366e247397a2f65162620dfd60bf332c58d4baa81a8c396008678c7a"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / c35ee54342bc / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-012.md#canonical-dd670faa4d2738047e1b04ae34ff99b083a94d521e786e08b6630f307e36f7e3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8714d4cc6433db78da7f53e4b4f7a99140641e4881f84a4faed2af2b78d4478c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1db664d100842f7dfa4b45f53a3b060ce2d434e9e52d593a0e8f7876a25b0e22"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / d8c80a218edd / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-012.md#canonical-b38bc9345af2dbf6ae72562dbd2795593a454abcce55a58b582f131c34a9d179)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-012.md#canonical-dd670faa4d2738047e1b04ae34ff99b083a94d521e786e08b6630f307e36f7e3)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-dbb9d9c740c24ce9ad83adadf18d192054dace01a3655acc992dea42a283a808"></a>

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

<a id="canonical-847bb82fdaf91e64eb611aff7a67b2e8eaff4f1182104d621f8ed17f07d536b8"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / d8c80a218edd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-aa9cb00d77dbde9b6b8f8af149238538b8f64c96e9397bcd7fb758c080135407"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / d8c80a218edd / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-012.md#canonical-dd670faa4d2738047e1b04ae34ff99b083a94d521e786e08b6630f307e36f7e3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-dee8a0435093b1e88a70585cfcf644cda2f9b5f2cc73f7e71deb8619b2153761"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fd74e48d3b951122a07eac7a48041dea448f162f0b2f6be37af2f6b177f11d6"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 789cd9d82562 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-012.md#canonical-b38bc9345af2dbf6ae72562dbd2795593a454abcce55a58b582f131c34a9d179)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-012.md#canonical-dd670faa4d2738047e1b04ae34ff99b083a94d521e786e08b6630f307e36f7e3)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-e80b2b8599811dbeaa53a08c8ba2bbf3aa571615cecc94652bebc129ebe68b96"></a>

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

<a id="canonical-aa0a5d20b6573be35cda3dfeca64125365cb8e822f7e43b6350ad5540386b277"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 789cd9d82562 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b63b89599192f54c984c367b717aee8cfffe9d37168132cd2520fbcf3b5b74eb"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 789cd9d82562 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-012.md#canonical-dd670faa4d2738047e1b04ae34ff99b083a94d521e786e08b6630f307e36f7e3)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-c1c7b4ea99b55862cf261f62878c7b5a4a7373a3d2cb0f7dba72c21465ee2e2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-712208a0997f7bacc08a1788adf85f4852c2568a2faa34785d292235b05884da"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2 — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 64b4a438ee5e / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-e11fb4883dac23d750c83be9bbaa3a16cb5a215949d997adc65160c4847cbbde"></a>

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

<a id="canonical-ae2e2c05308e101ce5e089a3b49bf01501e5862bcea9854feb253287be95c7f9"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 64b4a438ee5e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4bb780c7194405369afbcb60f1d8448f3a6aef34621314a2bd84d067063765c9"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 64b4a438ee5e / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-af3f69f6e53e3a9d3cbec66675f3f23f2c9ace3c6d1065ebaca272651950da5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ee5708efe9df2681a265ec0c9fc0a072413cd06b35e223aafd6ff413fcf67c56"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 62ca553095fb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-8bd01ff90de7d118af38140d22a9bb4ec47df106a36478f7f4f4e963ec364580"></a>

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

<a id="canonical-e689542e75b5c77b8aa36a895663cc8de36d3e1affdfe1e18fecd5f74fa055ee"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 62ca553095fb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e4d4a19d8911571c7460fd759bba0fc2f2283f2b2810993686fa205fc53dbbfd"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_ / 62ca553095fb / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.http_protocol_options](resources--workload--reference--group-012.md#canonical-1726dcef95e1456972b40505b8d6de6fab85ec41903966d9c01cfe365e0c191b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-21192f9ac919ff00667d635dc054bfc757937509f37f2dbdf3987c6a0ec5c0ee"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1e0017ea2ce0bf99ab48fcef40a4285ea6637bc5f823695558bc13b41b277315"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_d / b97a8a94c4cb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_default_loadbalancer

<a id="canonical-5cc7a9b530abe12a3f75c50d1b788aabf92f33cc65733482654f0e01a20ae71a"></a>

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

<a id="canonical-01a0c1e37f91aa3fcdd817d74fc35c8c97878dd322ee0e00f4dc5eb6ba1b34ae"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_d / b97a8a94c4cb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0404a9e36b43607fb8bfefd08f9d076fe3ca3c733d9693b7b3ab3b7e96acfdd1"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.non_d / b97a8a94c4cb / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7db59adac8ffdb78aa082ec701a37ab2eb64663e47676515049945c720c48ade"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8192ecde576edf9e69c3bd3591604af77fc4c5fa1187c6e8b778c267643fe694"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_ / 3c7ea84409f2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_through

<a id="canonical-ab674a2df86e1f05716c0032f94c9bbe3ae2969b27c56c07cc12716b47a4bd99"></a>

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

<a id="canonical-38c0d21aa29eb8aafb6c9d5e45a748c2556638bd1a20847a3a921d83871196c7"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_ / 3c7ea84409f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-225084b37275f0133c07d0f67e112fac753fc641ce357a9b390191dc9ed06350"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.pass_ / 3c7ea84409f2 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9ccc121e87dba2ce19411aba07a76cd4c5322a2d42099eed60a7e67cefbad2fa"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / f21c821041be / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params

<a id="canonical-aecadb484620a0edb11999bb12d1018bec0c8a0ac1f8f7ed098b7d442aa7412c"></a>

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

<a id="canonical-2d443c7cdc3d4b568f7feaec806c6150c19e5b2b7d8eeec9a69c9a9ba3cedd56"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / f21c821041be / 3

- [certificates](resources--workload--reference--group-012.md#canonical-7eb6ff2ea6f0c00d4d6c4dadf179c00e16002bfc975e0047df1dd9eda95c91d2): complete subsection reference.

- [no_mtls](resources--workload--reference--group-012.md#canonical-0518bbab00b32d9e803a2838e72f4195e1b2fa10fa294ad12dc564b9657f0e1e): complete subsection reference.

- [tls_config](resources--workload--reference--group-012.md#canonical-fed915352cd59c93c7a841b7a9cd8d5968444b638685afec110ad49b12bc67c4): complete subsection reference.

- [use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474): complete subsection reference.

<a id="canonical-31869bad000efc76afc1cdd4cd759a73220a25627bdb28a9a7a8d3cfc55b2db0"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / f21c821041be / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates](resources--workload--reference--group-012.md#canonical-7eb6ff2ea6f0c00d4d6c4dadf179c00e16002bfc975e0047df1dd9eda95c91d2)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls](resources--workload--reference--group-012.md#canonical-0518bbab00b32d9e803a2838e72f4195e1b2fa10fa294ad12dc564b9657f0e1e)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.tls_config](resources--workload--reference--group-012.md#canonical-fed915352cd59c93c7a841b7a9cd8d5968444b638685afec110ad49b12bc67c4)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.use_mtls](resources--workload--reference--group-013.md#canonical-5f52f209e7695c8b268068e20e4d3c53abad589ec4b704ca36f01d0c1d267474)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7eb6ff2ea6f0c00d4d6c4dadf179c00e16002bfc975e0047df1dd9eda95c91d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-10cca15ca68383581f1221bc04db56abfd30afaee8d52f1c913577260b2f6183"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 28916210ef65 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.certificates

<a id="canonical-dca3b7066745fe060def26c4b899f3ac44fb782c28785a6ac4654f1a52f7aaf2"></a>

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

<a id="canonical-ac3b362d9aa113d3bfb08b89dd72a9113b73e66c7fa19d2b411a208243af71bb"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 28916210ef65 / 3

<a id="canonical-b4eee723a0231881ead251fd65746a3c1699aaba1a8d1fd74d6143f8cec56d2c"></a>

<a id="canonical-4faebc23ddba9bfbe9b93e6782136ccee2575b510201aea34bdda00109202ccc"></a>

## name property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 28916210ef65 / 4

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

<a id="canonical-d9cbff6a617d47bb9ae6e773f18e8cdd02c012e07843f667358ba7671501a23f"></a>

<a id="canonical-1742d98a4b18c4d55b84bf999d8dfb8738ad2b771e33c3ea59d40070a0862cdf"></a>

## namespace property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 28916210ef65 / 5

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

<a id="canonical-9e72b7ca5172a56faeac11b127dc23ffb0515be768d9db9e59c469bae2c702e5"></a>

<a id="canonical-73a98ab6ba10e40832edfbbbee28700ba15d70a098ce2966c95d33585d681210"></a>

## tenant property — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 28916210ef65 / 6

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

<a id="canonical-94e9fdf980189488f2a8fc217bd70adf48f30ca5dfb8588f663d8132f4f724bc"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 28916210ef65 / 7

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0518bbab00b32d9e803a2838e72f4195e1b2fa10fa294ad12dc564b9657f0e1e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbc31260c9ee242e2783dfcde6b6498f78f5bfd95799edec82d5d16667cd2957"></a>

## service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 63f388cce5d1 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [service](resources--workload--reference--group-005.md#canonical-3878aa4259028f4891bf5a01e0e221cdb7c4d36f7eff864876c0158b26b73805)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-b9479d010afb8d2944cc4748edbd4e68599cf55d895c8be384000ac39782231c)
- [service.advertise_options.advertise_on_public](resources--workload--reference--group-008.md#canonical-6a412acfbb0120491654d23529c408d7962f92fd8786267fc418feb534cf15d1)
- [service.advertise_options.advertise_on_public.port](resources--workload--reference--group-012.md#canonical-c736a28c89aaf6a7a82aec43d2a00fb2359f5d1edfe6725b1fc2a9bc5a7432f1)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer](resources--workload--reference--group-012.md#canonical-110270702a2024774878e1f9ad1558bf9ab53f64e26de08837f836c2d1d457bc)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https](resources--workload--reference--group-012.md#canonical-a90389dd8194c28e8b75d5d4407c55f7435be600926ad039259a22ea4a1d8b55)
- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params.no_mtls

<a id="canonical-e61817bd9a3aa253424e2e003defcd7bea9ce074dcaf7e76b885ed0a98dc6cce"></a>

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

<a id="canonical-390e075c0bba89d1abd17d65623619e75841cdb0ba42726859887d78ef9099d2"></a>

## Direct properties — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 63f388cce5d1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-57518e87cd5cabd2e7db20d15e9d077e3ba22b94edc9014d63b7d7f253fe1175"></a>

## Next pages — service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_c / 63f388cce5d1 / 4

- [service.advertise_options.advertise_on_public.port.http_loadbalancer.https.tls_cert_params](resources--workload--reference--group-012.md#canonical-89bf1b3dd501e9cb38b4cfb8a7b8dce67b9a1a203705ac38e60dc82b8f88faef)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-fed915352cd59c93c7a841b7a9cd8d5968444b638685afec110ad49b12bc67c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
