---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-1d6346b196255f90e74f495751e8312fba8e2190ac04121d4db2856781b9362e"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b76f43279b34 / 3

- [custom_route_object](resources--workload--reference--group-020.md#canonical-7fbec2e0edd30c5ab4323e57fad55438e4866f4323d472112a492418ad14516e): complete subsection reference.

- [direct_response_route](resources--workload--reference--group-020.md#canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557): complete subsection reference.

- [redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da): complete subsection reference.

- [simple_route](resources--workload--reference--group-020.md#canonical-2e559abeab3e6aeb8f108fd42c1b5ef8bcd936463cdccf09a280711b511a3eae): complete subsection reference.

<a id="canonical-72f43b322e7f23560fcf7e07c073c676c57f4dae54a7e540a8498f04676bbb61"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b76f43279b34 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-020.md#canonical-7fbec2e0edd30c5ab4323e57fad55438e4866f4323d472112a492418ad14516e)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-020.md#canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-020.md#canonical-2e559abeab3e6aeb8f108fd42c1b5ef8bcd936463cdccf09a280711b511a3eae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-7fbec2e0edd30c5ab4323e57fad55438e4866f4323d472112a492418ad14516e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1a60c7c0db20949d905d61a4015f985b14a0ddf15c7bcb78c66159d52b1f850"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 515f45ff9aff / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-10965c4635ac6854a218d2b8eb9c8c1252447dce6a849f54080ab1023647c8a6"></a>

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

<a id="canonical-94d92084be75fdd1b82e1e3c0cca73cadee2abe46ee2393fe761ce0b81406fc6"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 515f45ff9aff / 3

- [caching_disable](resources--workload--reference--group-020.md#canonical-6ac06c1e1743ca82077ba494e04648ecbf32c22d7ff73e67fe885ac37d9db468): complete subsection reference.

- [caching_inherit](resources--workload--reference--group-020.md#canonical-8db05041a33dc376de8aa7074b372af3aad6d02fbef65a2480bd760bc965b0fe): complete subsection reference.

- [route_ref](resources--workload--reference--group-020.md#canonical-3fb8531834047db36d49f4d20a7af25f6013fcc549a329613350b1dcc5249a84): complete subsection reference.

<a id="canonical-5934d96c9cf65ae25b0471360534a100de4c5f1e9716080b72fb613f04a30fba"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 515f45ff9aff / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](resources--workload--reference--group-020.md#canonical-6ac06c1e1743ca82077ba494e04648ecbf32c22d7ff73e67fe885ac37d9db468)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](resources--workload--reference--group-020.md#canonical-8db05041a33dc376de8aa7074b372af3aad6d02fbef65a2480bd760bc965b0fe)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](resources--workload--reference--group-020.md#canonical-3fb8531834047db36d49f4d20a7af25f6013fcc549a329613350b1dcc5249a84)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6ac06c1e1743ca82077ba494e04648ecbf32c22d7ff73e67fe885ac37d9db468"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-193d7778b7c2b658d806470f9b16d55b396ab247062f0dac842a53d5d30082f0"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / bcf1f61fb0eb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-020.md#canonical-7fbec2e0edd30c5ab4323e57fad55438e4866f4323d472112a492418ad14516e)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-0ccf68c9358bd64a37cb41dae79a186481826a710e57f549a98545290a293736"></a>

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

<a id="canonical-a6b9e151964553a5c82316ffbb18979f45e91aa24a9566fe94dce4ab9404f4c7"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / bcf1f61fb0eb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-04efe2f91f58476a41af1a37c09eb18e6dda8b2f603dae0e5ff56c577e8ffdec"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / bcf1f61fb0eb / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-020.md#canonical-7fbec2e0edd30c5ab4323e57fad55438e4866f4323d472112a492418ad14516e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8db05041a33dc376de8aa7074b372af3aad6d02fbef65a2480bd760bc965b0fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e9601fa4b60b243c42c490128425246a9873253e457b14ad1d779f79d7581816"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 83dc988ac674 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-020.md#canonical-7fbec2e0edd30c5ab4323e57fad55438e4866f4323d472112a492418ad14516e)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit

<a id="canonical-473bc5f53b39897350130a7819f58558027bc74097aa7d47af0c3bdbab9ede44"></a>

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

<a id="canonical-56ff4a1b917334685b4db9fcc25ccf717c68ead7f583bd8c921d891651076499"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 83dc988ac674 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9e3b8e01fa0f0c2183714514c0e831053f1b4facfe325c32adb02cddfbef7fbe"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 83dc988ac674 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-020.md#canonical-7fbec2e0edd30c5ab4323e57fad55438e4866f4323d472112a492418ad14516e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-3fb8531834047db36d49f4d20a7af25f6013fcc549a329613350b1dcc5249a84"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cb47d74dce157b4f8d5bb4a8a7e72a5a073771648c7edce3ebfb3f4bd38c103e"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / f5c42980c658 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-020.md#canonical-7fbec2e0edd30c5ab4323e57fad55438e4866f4323d472112a492418ad14516e)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref

<a id="canonical-7c1156850266874da63018b959b853d5f2479cdc9a96b70b724349efcb2feef3"></a>

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

<a id="canonical-ba8c1ee6bb3f6f76ddb5ce7027e10bc5c6e75edef709563f3d7bf2377c324524"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / f5c42980c658 / 3

<a id="canonical-96f5ea073c412f9b32710456461896b3d33b673f9fa8a4f1073bd55a29249524"></a>

<a id="canonical-961cea1ffdd32f59cd1211c0825297c7609ca0f83a608cd16c903092664a1ff8"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / f5c42980c658 / 4

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

<a id="canonical-41c16ec60c4b789b88e23221e9c98ebead1ca18357d2a2bcdacddd7551b4130f"></a>

<a id="canonical-5517f6661e86ebf95b29c50874dcca50a2694c0b625b49280e3e3eb763e89d9f"></a>

## namespace property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / f5c42980c658 / 5

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

<a id="canonical-7dc8d479200ee3c60cf090a892781ad2f0e55349797f55076cf9cdb5016f2973"></a>

<a id="canonical-60e8a8e71a3d822b01fe10da2c34d3d7eadc3a2178554d0900a796fe6527ec07"></a>

## tenant property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / f5c42980c658 / 6

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

<a id="canonical-28e8a35fd1652fcb6547a30ab501045d85f569322b8ce1e0ec47947185353348"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / f5c42980c658 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-020.md#canonical-7fbec2e0edd30c5ab4323e57fad55438e4866f4323d472112a492418ad14516e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a3c52a8a6596d80a45f9f7fef43787dc84b13da9960ea743a18f06533656b4d"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c7631a5d9826 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route

<a id="canonical-7a74da3bea8c50e4fa6e2e48249cad007a81046e6604dbe24028aacf2b7ce05b"></a>

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

<a id="canonical-d266b83d4337b1af56db0c313102398781083d2273edfc0be7891c46b3738f9a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c7631a5d9826 / 3

- [headers](resources--workload--reference--group-020.md#canonical-4c4f6ee118623d3073794398e7ca28b1bbce32131360235a451f83e606826a9a): complete subsection reference.

<a id="canonical-4df179e542c525585046a8d38b826708defc4ff6c170ed908765a377b55f9244"></a>

<a id="canonical-4e638b734bf7dd782f47234cea3e09a52a46831af6f14ff8510ec94ab6a86b2b"></a>

## http_method property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c7631a5d9826 / 4

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

- [incoming_port](resources--workload--reference--group-020.md#canonical-ee3e715581752766722bc7e419991949903fad8fb8bbff28162e901655184a2f): complete subsection reference.

- [path](resources--workload--reference--group-020.md#canonical-8db1f2e156edc0ab3467b1857c5948c7d3e9ba34e11104ab3c5db2a67aafdfd7): complete subsection reference.

- [route_direct_response](resources--workload--reference--group-020.md#canonical-e649a20da45b6b679806f7a9261a24ca20caa104aa20dcb20d7b79bb4fc729bd): complete subsection reference.

<a id="canonical-b4d21bf1e9d888f48042f781c645e443d25d2c67c94df2c8418a5c6151d7472b"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c7631a5d9826 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers](resources--workload--reference--group-020.md#canonical-4c4f6ee118623d3073794398e7ca28b1bbce32131360235a451f83e606826a9a)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-020.md#canonical-ee3e715581752766722bc7e419991949903fad8fb8bbff28162e901655184a2f)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path](resources--workload--reference--group-020.md#canonical-8db1f2e156edc0ab3467b1857c5948c7d3e9ba34e11104ab3c5db2a67aafdfd7)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response](resources--workload--reference--group-020.md#canonical-e649a20da45b6b679806f7a9261a24ca20caa104aa20dcb20d7b79bb4fc729bd)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-4c4f6ee118623d3073794398e7ca28b1bbce32131360235a451f83e606826a9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d4dd38117b071dbf9aa5a1ed4740ee038252ec83fae0883e60ce7b123eec335a"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / e575ffd63f09 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-020.md#canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.headers

<a id="canonical-edff411c3ac2ae7e838beffaaaef6c14587e035084e46e8a474e69162d3c6278"></a>

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

<a id="canonical-d43e501e2f81d9dde3130dbd0afa8c414282a5b864bfbc757cb487236eeabf36"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / e575ffd63f09 / 3

<a id="canonical-5e7e286e852bc57296b3699b996183b2019014f6004f11c9c38f699cd9e168aa"></a>

<a id="canonical-a68b564dcd4c9eeffb56714ca646edcd3da18d4ad686c71dd8e1e445b02319d4"></a>

## exact property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / e575ffd63f09 / 4

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

<a id="canonical-8a138b51b26a21fb73f77addacaf71b2486446d4bd3710065cddf470d44e6f59"></a>

<a id="canonical-6f2eb3c4915a2f1b625d19f18f1f0ed84347ac58fe038cab93d66cf72a9cd84c"></a>

## invert_match property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / e575ffd63f09 / 5

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

<a id="canonical-1a620504d414987cc886e2205e2c858f801ea9e509f71deb277486b9d659049b"></a>

<a id="canonical-fd63d1573ea949e77bac40b0f30f6e4a39aada298d1a10fb0a483dc55dc8298d"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / e575ffd63f09 / 6

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

<a id="canonical-e767c157943a6f636c4f4a5b3f9a24656b2c0be32a1b1949281a4d0ef17c9007"></a>

<a id="canonical-4cb25ec17656a4f72bcbe0e470b44367e5d7f9bfcbcaaca5b86ac3b04c25133c"></a>

## presence property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / e575ffd63f09 / 7

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

<a id="canonical-4a07fd2fdc7f67ed339b18784c38b2dad1343416e19b275b90847f544cee5e2f"></a>

<a id="canonical-dc7f597d74ff23b5a8f2493d34614aaa81675ecf1ed858a9193b1554bc4ec0d9"></a>

## regex property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / e575ffd63f09 / 8

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

<a id="canonical-e33de4f3df1f148529ecedefd2273ccf25cc6e2703b517e5fda01eafa3de88ac"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / e575ffd63f09 / 9

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-020.md#canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-ee3e715581752766722bc7e419991949903fad8fb8bbff28162e901655184a2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5058b9c445df5493447771d7dd96a7c299b544461f5c696d7834e33e9f69304f"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b8df7e6e9aa0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-020.md#canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port

<a id="canonical-46f2ce441d72dc086f298e13b17ce2c95c5349815dbc72a9014442a9017b4371"></a>

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

<a id="canonical-5c2be68daba1804a95b5437d34976e7ebd2ef7417d3d85f80280470d2fb94efe"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b8df7e6e9aa0 / 3

- [no_port_match](resources--workload--reference--group-020.md#canonical-42eb8afe877e9f98fb52f88ff0ff40b2e4e2aa4349bf4b1afb13f99eb4f40676): complete subsection reference.

<a id="canonical-3cd143012bdc6bcaca7d415c5c3104f493dbcd5bc3c5547c2fdb6bc8c72ca71d"></a>

<a id="canonical-588ac2c6cfe6deb3a0a78cc5ac25ddd5f18cb2c59c5b9d22062923b71bc2253f"></a>

## port property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b8df7e6e9aa0 / 4

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

<a id="canonical-ce82e538d3e46faedfd7fe600983a6e617c6db9b078cf2f5a7ab4358cc20439d"></a>

<a id="canonical-ec7441ec7915800ccbd84c42552dc677ee03e14ebef4a5590e31429b01ade3da"></a>

## port_ranges property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b8df7e6e9aa0 / 5

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

<a id="canonical-63b8f8c216e6db065845bb1e9a0721393a593173c96ddd5e93cf8df6eb3468aa"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b8df7e6e9aa0 / 6

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match](resources--workload--reference--group-020.md#canonical-42eb8afe877e9f98fb52f88ff0ff40b2e4e2aa4349bf4b1afb13f99eb4f40676)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-020.md#canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-42eb8afe877e9f98fb52f88ff0ff40b2e4e2aa4349bf4b1afb13f99eb4f40676"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8aff6e5bf0ba5a8c3cfafe1ac1d307ecd387d3d86b5995cf18f4f0c28015d37f"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / dea314a6814c / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-020.md#canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-020.md#canonical-ee3e715581752766722bc7e419991949903fad8fb8bbff28162e901655184a2f)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-e3b52d08d81e7a3154b437af182122a3deb560f5d4810200daa172bc711264d5"></a>

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

<a id="canonical-d7e06eb7dc161ffa28dcb91822c5b805adbc319166026a0122c5cf61c747fdbe"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / dea314a6814c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-65fe8b433d8d903263449e7ba8524e287fe1a02beee7dd7827fb0c7e26d797f6"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / dea314a6814c / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.incoming_port](resources--workload--reference--group-020.md#canonical-ee3e715581752766722bc7e419991949903fad8fb8bbff28162e901655184a2f)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-8db1f2e156edc0ab3467b1857c5948c7d3e9ba34e11104ab3c5db2a67aafdfd7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2824b8d7c23c731579534520698e3fa24ce8597610805e9d8ac1af640abb2cb6"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 83856b8a7fbc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-020.md#canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.path

<a id="canonical-90bfa5f723aafec3f5779415e5c2898dc15ebf07bd0ef1cd284ada186e8f6fca"></a>

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

<a id="canonical-bdb585ac90763311c29d6e09294f2234b96d7ab77af804d9d2f93c426bb14c9b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 83856b8a7fbc / 3

<a id="canonical-1a9be9e81d0592c0229fb25d1b857a50855ed015c478a289ef0cc4f9e00e2123"></a>

<a id="canonical-753716a16af2360d241fc11b702425731178dc60e15be9beec5b6dfeee5713cf"></a>

## path property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 83856b8a7fbc / 4

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

<a id="canonical-f7cb8924afbc2290f4c02f871cb51235c7797f8b84c4b93b15ee59fdc77188bd"></a>

<a id="canonical-3f8037e14102302f6cd8c6fd42ba3df43b97274d3750bde72250b42f1d500f2a"></a>

## prefix property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 83856b8a7fbc / 5

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

<a id="canonical-15f584dec5bb0ed0a3deb1e8fc78e995aad117cf59326647a14d8f9036b87fad"></a>

<a id="canonical-7dd80d440338ca2fa83de0fa3c9765e395017203a753fc5af5309e1c632bd352"></a>

## regex property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 83856b8a7fbc / 6

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

<a id="canonical-800ec51f388b7dbb89992b3633a2ca75c39ff638a08a2bf7f909b9c2141d0766"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 83856b8a7fbc / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-020.md#canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e649a20da45b6b679806f7a9261a24ca20caa104aa20dcb20d7b79bb4fc729bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2949a17430f617d72f65d35837752436aba071264c3d1c53c8f902989d0effad"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / eebe46e1748b / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-020.md#canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route.route_direct_response

<a id="canonical-9909df52b386a37e1816f048e4dcf07b04074f3efb593584f91db142084e215b"></a>

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

<a id="canonical-66f5a0a12f5e8b5a3d6eb7da3be01649cb013a2345775e4f94e72d45327114e4"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / eebe46e1748b / 3

<a id="canonical-9ac0de9690f965e35a468d2ca7a14558ae145b3a89675eff75a190eeaa92f231"></a>

<a id="canonical-0de92b0aeddcfdffd9900f336e4da77afa1dfe79a9b9660cf1c5934fc2a86ce3"></a>

## response_body_encoded property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / eebe46e1748b / 4

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

<a id="canonical-4417f562a21c372697322df5980e3a7b3f78baca1e93bcc35c41d1266ea45a3d"></a>

<a id="canonical-b8a8659b5208b2bd70425b615038169c0d3921d893f46a10069afb60fdf727ef"></a>

## response_code property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / eebe46e1748b / 5

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

<a id="canonical-8897b7ccce116f19350b79358e236e6a658249232c6edb0a2d0c13444df4ad94"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / eebe46e1748b / 6

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-020.md#canonical-cea826a4870b1fff46df62829211ad103112c20fe91b41204c67c7719394c557)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7bfd8141fb3cda923b44b4d3bd36b47aea2685e1cbf77caefaf14acdd0034a4"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 1be4855a16cf / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route

<a id="canonical-b15a0818999761f548a32828b7b50d322936bd47831d193f45ddbc97e60e89bc"></a>

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

<a id="canonical-1cbf5f4ebde80d15cdb864392482b9d714115d74b0ee571a3968102f7f55f187"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 1be4855a16cf / 3

- [headers](resources--workload--reference--group-020.md#canonical-fa5d60ee6e4928dd39a682217f3970327f9bf0757d56bf4cf55dd558c41ee0a3): complete subsection reference.

<a id="canonical-9635e8c63d9b6035c75211feb630f98753b804f78e27d4a251657e4e6c19710f"></a>

<a id="canonical-fd18bf73e2d085d09235d0a880c9e9d9a0c6844e5b2530a5cdf6f0e06030a082"></a>

## http_method property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 1be4855a16cf / 4

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

- [incoming_port](resources--workload--reference--group-020.md#canonical-279830654b9bfadd268625b02c53d8093c47d72c69727958dad211741bc8eb08): complete subsection reference.

- [path](resources--workload--reference--group-020.md#canonical-b686bf6a7dab147a9a30276f89bb6c16f5dfed958711c39bf0e7551995b39dc4): complete subsection reference.

- [route_redirect](resources--workload--reference--group-020.md#canonical-fec12224b7c38ef656236e97e427ce13f4470a666e0deb242d0f5a4c1cf8e6af): complete subsection reference.

<a id="canonical-15e6e4e0ed3170613a534d848fb7beb60f75ccaa0716f8592320a2c52186bc7d"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 1be4855a16cf / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers](resources--workload--reference--group-020.md#canonical-fa5d60ee6e4928dd39a682217f3970327f9bf0757d56bf4cf55dd558c41ee0a3)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-020.md#canonical-279830654b9bfadd268625b02c53d8093c47d72c69727958dad211741bc8eb08)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path](resources--workload--reference--group-020.md#canonical-b686bf6a7dab147a9a30276f89bb6c16f5dfed958711c39bf0e7551995b39dc4)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-020.md#canonical-fec12224b7c38ef656236e97e427ce13f4470a666e0deb242d0f5a4c1cf8e6af)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-fa5d60ee6e4928dd39a682217f3970327f9bf0757d56bf4cf55dd558c41ee0a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eeca31b1acce871555ad4c790e6ab2c1234c141a21317418ce489c2f45d23975"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 87cc84799be0 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.headers

<a id="canonical-6ce949baafca8ef78c6aa882575081d1afcfb380f4613d70fb652383e327e3cc"></a>

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

<a id="canonical-9d1e235cae6b90533f76251f52f597ccd9a9870731d509692415a220d2846da4"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 87cc84799be0 / 3

<a id="canonical-6be03828c08b1d4e76b0ca94a034b382721ddbac696500551a4892ece7759632"></a>

<a id="canonical-f7e5981afbc4d5b02e046dfc75b6873a7e422fb6ef276de1325588909cc0adf5"></a>

## exact property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 87cc84799be0 / 4

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

<a id="canonical-33c7f9244f1bb5ff0623d9aa753ea3796789cf0f4d8b057f44764df1da61c260"></a>

<a id="canonical-35dfdebf65945a8ff0cb8b0799cd08ce185c98747bc0e8a17c7ccc90ff620d95"></a>

## invert_match property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 87cc84799be0 / 5

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

<a id="canonical-c4808204dec115973f9bcbca6d895df72eca0b71ca07666ffef071fe55456e26"></a>

<a id="canonical-2bbfab763e63d2278df30d236f244a4bc740724c98af760797e3a6e79990bc25"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 87cc84799be0 / 6

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

<a id="canonical-8b065119374df4eef8389a78fac25704e6272d1a4fb8bf25fb198fd613a849d4"></a>

<a id="canonical-55944fdd27de8e28e9d436ae5f34de3d5c464cdabefcd6d0c3baa45e0a760f77"></a>

## presence property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 87cc84799be0 / 7

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

<a id="canonical-7b8adc9ca09054aefd767c06f6208bc8c21f402d0454fcc3cda8674eef6eda70"></a>

<a id="canonical-70b110fe2d7bda5dcb386683f066825102b7b3b5c4630b69a5710772fac3759a"></a>

## regex property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 87cc84799be0 / 8

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

<a id="canonical-a7af8831b801969b6dc85743319985d2571af6c341bd5b8e250af5fcc48d3e51"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 87cc84799be0 / 9

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-279830654b9bfadd268625b02c53d8093c47d72c69727958dad211741bc8eb08"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c0d9bbc98d24f3a7bfcfe1635a35be3736186426ce1388197d2a05cbb1b0b39"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 0cffaa925621 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port

<a id="canonical-f80b9e731b735dd28085632fd53e7b3aae7ecf4a1529c2ea61c4ae5e57766955"></a>

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

<a id="canonical-0e0c79af870ae0dd6a0ec9e10d37072f830f21022f27f067f3c1635bb48cf548"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 0cffaa925621 / 3

- [no_port_match](resources--workload--reference--group-020.md#canonical-2cc76c44e918ed51eba902399c3e372f39b7f4bc9a6298f196dc006e95256495): complete subsection reference.

<a id="canonical-c8ee450190f5f70b4f9669a01919a07e0355b6927ac293b7ddee0cf0c23be82d"></a>

<a id="canonical-64d138e182226d15c2c4390b0ea7c2bb3a98877c201cbcda2d90d243ad9634d7"></a>

## port property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 0cffaa925621 / 4

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

<a id="canonical-43aa18f2b45f0cd5de8585443bc2981701d1f00724c566189c8173682ed41583"></a>

<a id="canonical-9d1c28ce54350e532d8cb876977e0f594e960334fef754a3ed21bb57ea6e9efc"></a>

## port_ranges property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 0cffaa925621 / 5

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

<a id="canonical-e7e22f19a3ce35b31ca6c46de5beeebd7cb565d64da65d2fa7a98a31aa4b1590"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 0cffaa925621 / 6

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match](resources--workload--reference--group-020.md#canonical-2cc76c44e918ed51eba902399c3e372f39b7f4bc9a6298f196dc006e95256495)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2cc76c44e918ed51eba902399c3e372f39b7f4bc9a6298f196dc006e95256495"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22fa914fbc61babd4d59b52fcabe99fe72fc514686f6982d7e52488bb523677a"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c6aa81c78a8f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-020.md#canonical-279830654b9bfadd268625b02c53d8093c47d72c69727958dad211741bc8eb08)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port.no_port_match

<a id="canonical-389e93c822829f4c4987a5a904aef00bb6788978dce7879c5932910b693871d9"></a>

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

<a id="canonical-b2bfb595a356ceb502758e2b9e7a1c5f0bbcca243551adfe31fc15ee1a226904"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c6aa81c78a8f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8541efe31b17e298f9673580b07cb9db50c883c8903c6d6cc7ead020d9c73049"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c6aa81c78a8f / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.incoming_port](resources--workload--reference--group-020.md#canonical-279830654b9bfadd268625b02c53d8093c47d72c69727958dad211741bc8eb08)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b686bf6a7dab147a9a30276f89bb6c16f5dfed958711c39bf0e7551995b39dc4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d4ac30e174af683611debe3d1b96453603e386ab3b67bc006b3570d302d5c04"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 6a6b80c121e6 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.path

<a id="canonical-9cb0f11280172b3d7abf75279a21d68ab16e871d042f62fb79e0c5d54bbf2ed1"></a>

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

<a id="canonical-1c9ed791780e053f81f038cda2c7d24836917d1bff05da460d4d81f065127a94"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 6a6b80c121e6 / 3

<a id="canonical-55de6f09aca4001d745f590691d2a961539466b8fc44139eeb53cbdafcb47950"></a>

<a id="canonical-b4fce5e0e91c1190be8262254a069e6a8353d0cd54398e486e1fdf52ea01ca3a"></a>

## path property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 6a6b80c121e6 / 4

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

<a id="canonical-2d6dcfb1027bea5687a18bee7bf4f71acd3692f5719e3531ea78fe011963e194"></a>

<a id="canonical-a2a9eea27ad1b06fca1c64bebadd698644918767bc40076c964f2c7766dcc2b5"></a>

## prefix property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 6a6b80c121e6 / 5

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

<a id="canonical-48c908b588849f8ecbb098a2c4abdfdf7b87b526a9686abfa3d521a54d65bbb5"></a>

<a id="canonical-7df374b63f502d248fa5d77f4056c7da3f2280fe180f95cd93a09a630e594719"></a>

## regex property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 6a6b80c121e6 / 6

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

<a id="canonical-b4287527911b925c881393cc45791f3cb71eb2141b1f87141ad44594f484c24c"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 6a6b80c121e6 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-fec12224b7c38ef656236e97e427ce13f4470a666e0deb242d0f5a4c1cf8e6af"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1caeef105dd3633f3297088a36021c9710b9ffb3655f93fb19056839189f5c83"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b1db10f985f9 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect

<a id="canonical-344f4d00e3dee3c9b075f025e654839779d96da16896f9974e3e686e4a9cc1df"></a>

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

<a id="canonical-3368ea3ecbbc2e8dfd8943c20b0604b3ca2de5718d0604d5ce57ba7d99cdde45"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b1db10f985f9 / 3

<a id="canonical-3c949b0a746933c51c337a6f4b7cf211c974cbbdaf3c6c2f649b7aafc1791d33"></a>

<a id="canonical-a830f2aa9912123a67a40dad2e566a84cb1be48ceda3b956007cb195947b7ac8"></a>

## host_redirect property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b1db10f985f9 / 4

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

<a id="canonical-ed3c90b8268dc72069003b967a2de363d6d58fd2db0f14a6124d3df194717d1e"></a>

<a id="canonical-f1d186b70a2aa7b0334bfb19ff0809a57bd6454cd9ca11bad9400801ab6db00d"></a>

## path_redirect property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b1db10f985f9 / 5

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

<a id="canonical-09590c632c9203b6ac53e40c54e7e5086d90efea1e30b946e4927970b64862fd"></a>

<a id="canonical-bd0dc531ed0fe959a32abfa79e4011a200610ec035634374c4bebca296f8f2b1"></a>

## prefix_rewrite property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b1db10f985f9 / 6

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

<a id="canonical-321d381f8393859c6f3dcc08f96026b8e9246be0fd209efe5f21370e9c13bf04"></a>

<a id="canonical-12a567842fe8a71474c5f5a696b19315d380dfdb27d165a58e8de4954da2bd43"></a>

## proto_redirect property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b1db10f985f9 / 7

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

- [remove_all_params](resources--workload--reference--group-020.md#canonical-481c724ffa4d87c23c6f462337576f432891ef26746210b1273b838951301765): complete subsection reference.

<a id="canonical-342eb9eebc26f71813a1be9b25890d84f2d9112f4f0cc9a436196e2e0408eb23"></a>

<a id="canonical-a2633470f20defae5384fc5d1c6019bd7a5946c2aa7e9031c55cb0ae39b0be20"></a>

## replace_params property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b1db10f985f9 / 8

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

<a id="canonical-f80981411d0925ad3fac5a5f698003c654164c3982dfbc09382c2e68abb260cb"></a>

<a id="canonical-324283c4aed9d74911c9abe0b1fc599247b7e0c91bbb84c823ee9792affc7a06"></a>

## response_code property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b1db10f985f9 / 9

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

- [retain_all_params](resources--workload--reference--group-020.md#canonical-6efe00012691ddd898f0104728a100c0012a036923ab56811f34809b73beb718): complete subsection reference.

<a id="canonical-311623da85fc567b402dcf39f030532a28435eb855af61ba5d6016f27dd82706"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b1db10f985f9 / 10

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params](resources--workload--reference--group-020.md#canonical-481c724ffa4d87c23c6f462337576f432891ef26746210b1273b838951301765)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params](resources--workload--reference--group-020.md#canonical-6efe00012691ddd898f0104728a100c0012a036923ab56811f34809b73beb718)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-481c724ffa4d87c23c6f462337576f432891ef26746210b1273b838951301765"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-049f0212e12da5c50edf7d2b18b7773782d5b0668e890a57b660b6daf0279b82"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b06f0badd902 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-020.md#canonical-fec12224b7c38ef656236e97e427ce13f4470a666e0deb242d0f5a4c1cf8e6af)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-dce1e5d1d85c6e8ed07bc385e7c37eff216bd4e6ca7633d457d1ff057f676639"></a>

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

<a id="canonical-2bf2ba498c08b5115d35616ba2a4b0f1b99578c958b446cc1238e192092576a2"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b06f0badd902 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46ea2eae6cce1d6d1bd0921f7be3f7d49a3b1943a835a8e6d9323fdee0642b20"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / b06f0badd902 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-020.md#canonical-fec12224b7c38ef656236e97e427ce13f4470a666e0deb242d0f5a4c1cf8e6af)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-6efe00012691ddd898f0104728a100c0012a036923ab56811f34809b73beb718"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eadda576520c752a225c37341273dde4e98970d044c7fe2abc5c110758c1d2ee"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 4c069ddfa0ab / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-020.md#canonical-996aaa49b12eb0e6706923143ed2b5905477ae431cfac1f8e44fc7b9b66b54da)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-020.md#canonical-fec12224b7c38ef656236e97e427ce13f4470a666e0deb242d0f5a4c1cf8e6af)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-79a78b361e4799552de173a0e2955499e5c3f22e2b5bde88903738384211a167"></a>

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

<a id="canonical-26fd6631ea2841c0a52ba0f7db711741e4e20d62dda819cdb5c757778a014d0a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 4c069ddfa0ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f8074d4e8003d883ce4e801fa8e6ddbf17479458f1e3ea3f7f305f29abc194ba"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 4c069ddfa0ab / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route.route_redirect](resources--workload--reference--group-020.md#canonical-fec12224b7c38ef656236e97e427ce13f4470a666e0deb242d0f5a4c1cf8e6af)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-2e559abeab3e6aeb8f108fd42c1b5ef8bcd936463cdccf09a280711b511a3eae"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a13350e7ba47fc2887ecf2a145a42c57a626c47d22de4d2a57f5999285736ba"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c4a60535c4cc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route

<a id="canonical-28f6685b218463b167ebf9c6717304b32d1471924a139179179141c05dc5d1f9"></a>

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

<a id="canonical-6dcfe263672f0002be49e26ccd65a0a594e333f334ea6c2d1d82df873dbc3f06"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c4a60535c4cc / 3

- [auto_host_rewrite](resources--workload--reference--group-020.md#canonical-831d3b0713036445ae769f0bd5850ec05902ab26d0f2490e8315e3c2bdb1b76a): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-020.md#canonical-b3cea2d403f07487a1f9215c0d72b99c2f5e8f3ed52e89dedf94be718f524518): complete subsection reference.

<a id="canonical-db42ddd0e07a1c74184b1fa16123868f7c01e1b3f3c28ba1395151fbaacc59e1"></a>

<a id="canonical-81d2da6ce75e258c89e427db97c1189c8576a70d0182b07cf6bbcf6eac18b05c"></a>

## host_rewrite property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c4a60535c4cc / 4

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

<a id="canonical-4c72f52bc088a7fa7a3f4344b6fd5b964c7f62ede866396b3119a2bfbb886e10"></a>

<a id="canonical-2736457c5973a368d8f14f69fa75b0e3d7541d7c91e3982f0eddece8fbaf870a"></a>

## http_method property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c4a60535c4cc / 5

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

- [path](resources--workload--reference--group-020.md#canonical-e5c919a36956552b87a8a16b33e391a8ae7e305b0b1b1ddd55f2636b13ca6b1b): complete subsection reference.

<a id="canonical-2f2a3e7008ed0ebf5fded9ef527e9f796f11652eac9b6de3c260c92002119799"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c4a60535c4cc / 6

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite](resources--workload--reference--group-020.md#canonical-831d3b0713036445ae769f0bd5850ec05902ab26d0f2490e8315e3c2bdb1b76a)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite](resources--workload--reference--group-020.md#canonical-b3cea2d403f07487a1f9215c0d72b99c2f5e8f3ed52e89dedf94be718f524518)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path](resources--workload--reference--group-020.md#canonical-e5c919a36956552b87a8a16b33e391a8ae7e305b0b1b1ddd55f2636b13ca6b1b)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-831d3b0713036445ae769f0bd5850ec05902ab26d0f2490e8315e3c2bdb1b76a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08cb3e9c6a44713497646af2e16b7f18f0e9725641ade6048c0af2aa15b19bed"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 0589e84a1efb / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-020.md#canonical-2e559abeab3e6aeb8f108fd42c1b5ef8bcd936463cdccf09a280711b511a3eae)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.auto_host_rewrite

<a id="canonical-b6c43683b0bd12b73a41ef0d340888d26a71cd96f952a364039c7616d3bb88d2"></a>

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

<a id="canonical-f5f8ea40ebbc21d5020c2e82677c24fc054422a40f67a76275fdbce04be8a8f8"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 0589e84a1efb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b5edf01652c66ef0a35b32ed3318efa47d901fa0013f7627668c42a9a3f5f937"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 0589e84a1efb / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-020.md#canonical-2e559abeab3e6aeb8f108fd42c1b5ef8bcd936463cdccf09a280711b511a3eae)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b3cea2d403f07487a1f9215c0d72b99c2f5e8f3ed52e89dedf94be718f524518"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb3be114bbb6d969ced72d56278511eedf421f3c0b8925499859eac054634f7b"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 9c92d5d98f31 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-020.md#canonical-2e559abeab3e6aeb8f108fd42c1b5ef8bcd936463cdccf09a280711b511a3eae)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.disable_host_rewrite

<a id="canonical-36ff8259667b07d2e802ade6dec8f6ae4e0eddee31da074520bd2321c3af351c"></a>

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

<a id="canonical-b4b13a1ba20e29e1a220ddbca168a0a38b3b660409e672d2c05d123b6dbe506a"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 9c92d5d98f31 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a4a7c73fa925e8af5f08871285ce63bc72e526085b0c75b4f30b22cbaae6abf"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / 9c92d5d98f31 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-020.md#canonical-2e559abeab3e6aeb8f108fd42c1b5ef8bcd936463cdccf09a280711b511a3eae)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e5c919a36956552b87a8a16b33e391a8ae7e305b0b1b1ddd55f2636b13ca6b1b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f97129c3e60e8be6fa60242f60421f20bfd802f725b3e32012712f8cbf8b1c5c"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c48dd13815ab / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-37a8a0e9b14ec8cbd8ea43d099a04fa1dfb4926877075409374899663de89882)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-019.md#canonical-12085629394a59606f007a8c16c0628b4b5c34522811388cb590c912a04fbeae)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-019.md#canonical-13c9a4b24a026095ca642b6bdbbcb8d9392a1b6aac25c2e5b5050f7c95bc1037)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-020.md#canonical-2e559abeab3e6aeb8f108fd42c1b5ef8bcd936463cdccf09a280711b511a3eae)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route.path

<a id="canonical-809e0d6e7150febe53463551c8a27a1884e782a8da0fece7ab36111ddf65384c"></a>

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

<a id="canonical-80392b0434c37847f39b5dbc5fc91172a5726d385483531f557d2ed0053991da"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c48dd13815ab / 3

<a id="canonical-0a538ab69a46a270bce4401486703ddd7e48aadcbce218c92b20e06297e92cd4"></a>

<a id="canonical-359b83c1997e768af5fbe3eb92f6f4452226ff495c7022f82ea1e084c0c2d9d9"></a>

## path property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c48dd13815ab / 4

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

<a id="canonical-516172c6af9120caebe7c78ed735a36ffce2667d68d7fee43c6aa444e6c7429c"></a>

<a id="canonical-bfb696aa4262a01474647e7d92da8b4b3112c3709cf88557e19fcf983d6188ab"></a>

## prefix property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c48dd13815ab / 5

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

<a id="canonical-8bb8f7acf23410796a4c16d8a02c1e1fac1d7e9ab59c00cc34ab0382ceb62b96"></a>

<a id="canonical-f0c1317da7359a558488663b1719ef39fa77cf90290a650b71729cf23b87d803"></a>

## regex property — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c48dd13815ab / 6

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

<a id="canonical-cd899f54b7e346771652c27b435cc6797d8abf4f9a54f85bfb0a959499cf43b4"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.spec / c48dd13815ab / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-020.md#canonical-2e559abeab3e6aeb8f108fd42c1b5ef8bcd936463cdccf09a280711b511a3eae)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-9346adde0cfe4a128d86ed815352cce79b915dbabfc95db14d85f1bcf36002eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac4ddcfcd84f743a41b09b3b07379095301057ef037836ed9fa6f0398ca414ec"></a>

## stateful_service.advertise_options.advertise_custom.ports.port — stateful_service.advertise_options.advertise_custom.ports.port / 4939f1d944c2 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- stateful_service.advertise_options.advertise_custom.ports.port

<a id="canonical-be602fb5215e703175031b36c26f3087c14c8e5a96c7c741bf5135d6b0e1203e"></a>

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

<a id="canonical-32483e55d27af3a7087e2dee01a1079044409734248be8845642a591dedf5621"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.port / 4939f1d944c2 / 3

- [info](resources--workload--reference--group-020.md#canonical-f779590e0a1f556b435cfaeb4c5cd39d77197c3aa053dbad77e7c88b46389485): complete subsection reference.

<a id="canonical-5d1c7978d94f671a80edbb22990ac1a0d9d21694c95a3449ec76c0276912a6b4"></a>

<a id="canonical-95fe0e83119a8844b3653a61b516b76b07aa24cfe9daf2c688d09d101794051f"></a>

## name property — stateful_service.advertise_options.advertise_custom.ports.port / 4939f1d944c2 / 4

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

<a id="canonical-cbc2c026d596f8da807722a343832cced2c1af6c6acb8353972cb00bfa75c960"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.port / 4939f1d944c2 / 5

- [stateful_service.advertise_options.advertise_custom.ports.port.info](resources--workload--reference--group-020.md#canonical-f779590e0a1f556b435cfaeb4c5cd39d77197c3aa053dbad77e7c88b46389485)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f779590e0a1f556b435cfaeb4c5cd39d77197c3aa053dbad77e7c88b46389485"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-068626ee14bb46714fd5b8afde491be65bc07488f56349ec5a61993aaaf0c268"></a>

## stateful_service.advertise_options.advertise_custom.ports.port.info — stateful_service.advertise_options.advertise_custom.ports.port.info / 42c271476862 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-020.md#canonical-9346adde0cfe4a128d86ed815352cce79b915dbabfc95db14d85f1bcf36002eb)
- stateful_service.advertise_options.advertise_custom.ports.port.info

<a id="canonical-c3c9fe9180431ec2f2155b3b082409a1e3e2a79d63c45273ac2ac8b54c854a0d"></a>

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

<a id="canonical-ecd2b23264a62fff7b28851ae567124d056c5d19d88e6d77a37c090e199b8b74"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.port.info / 42c271476862 / 3

<a id="canonical-6ebfabc52b2912af88d5754784011612cc6f3a9f1ac60969c590debc39586dc3"></a>

<a id="canonical-101081375a79486a2ad235262b6a8ade60bb6ecece68f988230bbbc2bd85ba6b"></a>

## port property — stateful_service.advertise_options.advertise_custom.ports.port.info / 42c271476862 / 4

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

<a id="canonical-cb865d41d10631eba5c780666f488d7c6ab40fd012f55005028e6a45440e8243"></a>

<a id="canonical-21e966abebbbd15230f43bb807ba63c178070b73b9388e6e85ce9abc0da4526c"></a>

## protocol property — stateful_service.advertise_options.advertise_custom.ports.port.info / 42c271476862 / 5

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

- [same_as_port](resources--workload--reference--group-020.md#canonical-0d01e088599c30e898cff735ace5cd6fbf45aae82dc47830c0af34fec8103b54): complete subsection reference.

<a id="canonical-34e726f0a173cf94e8a288918fdf2a9dc9bd770fc995220062e70c7639b98866"></a>

<a id="canonical-6f35bc187f55dd58b577ab6d8dcf7cd412f465e3a171ff7dbe877fbba2d8bdd1"></a>

## target_port property — stateful_service.advertise_options.advertise_custom.ports.port.info / 42c271476862 / 6

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

<a id="canonical-c0c1831e7d731940b327263ec63aff35518ec49b310c731877e4328eef1f6d27"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.port.info / 42c271476862 / 7

- [stateful_service.advertise_options.advertise_custom.ports.port.info.same_as_port](resources--workload--reference--group-020.md#canonical-0d01e088599c30e898cff735ace5cd6fbf45aae82dc47830c0af34fec8103b54)
- [stateful_service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-020.md#canonical-9346adde0cfe4a128d86ed815352cce79b915dbabfc95db14d85f1bcf36002eb)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0d01e088599c30e898cff735ace5cd6fbf45aae82dc47830c0af34fec8103b54"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6aa3002058726b803bd317098a85ebf72bab152bc54f7a8663a37a982f126aba"></a>

## stateful_service.advertise_options.advertise_custom.ports.port.info.same_as_port — stateful_service.advertise_options.advertise_custom.ports.port.info.same_as_port / 074ad4373611 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [stateful_service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-020.md#canonical-9346adde0cfe4a128d86ed815352cce79b915dbabfc95db14d85f1bcf36002eb)
- [stateful_service.advertise_options.advertise_custom.ports.port.info](resources--workload--reference--group-020.md#canonical-f779590e0a1f556b435cfaeb4c5cd39d77197c3aa053dbad77e7c88b46389485)
- stateful_service.advertise_options.advertise_custom.ports.port.info.same_as_port

<a id="canonical-99e0b69aa10799a5c9848d96978c6b17e75358174eeb8913d652801a93166220"></a>

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

<a id="canonical-7ac0c6f056c2ea49fa14e17b505532a0a8c3299007a756a1a36e1589996c77ea"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.port.info.same_as_port / 074ad4373611 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4586ee28c554eda13840a7fac54ab9cdfdf1ec486caddffa9c16aa91c7710137"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.port.info.same_as_port / 074ad4373611 / 4

- [stateful_service.advertise_options.advertise_custom.ports.port.info](resources--workload--reference--group-020.md#canonical-f779590e0a1f556b435cfaeb4c5cd39d77197c3aa053dbad77e7c88b46389485)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0846336c6bc25f0f9b5ca92bf4fc57d0fca71df4dd26f175e981080630a50f3d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9daab5449562f8d3b8308be650dda481284cf413b97146cfd9bbc3047dcc5919"></a>

## stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer — stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer / 426e1468b39d / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-23481e1b5870672615468b7fd1f5b8d9540350ee79b618b6e080979620cedf02)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer

<a id="canonical-86cb609496dd9fbdec8f80d043c02e9018bdfebf3ae8d130b94a515408ea2e49"></a>

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

<a id="canonical-0553b65e76f27d7d9e48d20da54373349f1e8695b639959b414d16cdf9b26d0b"></a>

## Direct properties — stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer / 426e1468b39d / 3

<a id="canonical-28878388a3b160bd2942b8b92d52edb320f3c20f04cf7b816cce6ba05753cee0"></a>

<a id="canonical-d35f4e70803ef0d12d7bb6bda96177c4f2b597c31163d89ebb10f2a4ed6f1388"></a>

## domains property — stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer / 426e1468b39d / 4

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

<a id="canonical-cd1fb404b32be17195ac0e76cc13838f3e964bb7481a57ee8096604a9d1a273e"></a>

<a id="canonical-4c85037e6911ea6b5edf0ac54f3ae112f8641dc0dc4a7d64810d07649cb4d34b"></a>

## with_sni property — stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer / 426e1468b39d / 5

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

<a id="canonical-81713a50d7270e2a71bb45f43bcf7c18fcf17fc08e4ae237abfda8d81e2f40a1"></a>

## Next pages — stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer / 426e1468b39d / 6

- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-d5d7afb361308261df26f0378affd762e9d8eecfdf6185eefc258d3dbdcf23ad)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a3b619f037024ba6c32ce619a795e69f40fb1650da727d6bb62996e2468be2a"></a>

## stateful_service.advertise_options.advertise_in_cluster — stateful_service.advertise_options.advertise_in_cluster / cda442e1a85f / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- stateful_service.advertise_options.advertise_in_cluster

<a id="canonical-a44cf0cea05dafc8d48c4919b893ca2648be9b773862e696658282306faa2186"></a>

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

<a id="canonical-dbc7c0c25136526bc6776472dbb05e6d40c92bc8244b1cb2dfec836109b75d56"></a>

## Direct properties — stateful_service.advertise_options.advertise_in_cluster / cda442e1a85f / 3

- [multi_ports](resources--workload--reference--group-020.md#canonical-45d897135e869ec5d9fc84f9ce490d93d685d10b9b644dc0972be10c0d3b957e): complete subsection reference.

- [port](resources--workload--reference--group-020.md#canonical-396880e662d7e5243b3407d5c9e82062f2dfca8f6f9725b8d1db94a0ece1b5f6): complete subsection reference.

<a id="canonical-539a01fbe0c81cea28a5536c571671b998a3d758f95289a6cd47b90aae25d184"></a>

## Next pages — stateful_service.advertise_options.advertise_in_cluster / cda442e1a85f / 4

- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-020.md#canonical-45d897135e869ec5d9fc84f9ce490d93d685d10b9b644dc0972be10c0d3b957e)
- [stateful_service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-020.md#canonical-396880e662d7e5243b3407d5c9e82062f2dfca8f6f9725b8d1db94a0ece1b5f6)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-45d897135e869ec5d9fc84f9ce490d93d685d10b9b644dc0972be10c0d3b957e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c99580f54c7f91efab584b8246a052631824f29f517b108697b9918272f46bfc"></a>

## stateful_service.advertise_options.advertise_in_cluster.multi_ports — stateful_service.advertise_options.advertise_in_cluster.multi_ports / f0f48e61c0fd / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-020.md#canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports

<a id="canonical-9132d15e08073e91378873ac7509552522ebf6b58291d9c761fe8df0af869422"></a>

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

<a id="canonical-58d6b053ab59df39446cb91482bcf50703bfa6668ef67790a87ae5a212a00011"></a>

## Direct properties — stateful_service.advertise_options.advertise_in_cluster.multi_ports / f0f48e61c0fd / 3

- [ports](resources--workload--reference--group-020.md#canonical-0acbd11f0964dba2dd1a73c816c866bc2d5a2edf20fbcd7b1040639b08fd04f2): complete subsection reference.

<a id="canonical-4f952c93d1516635cc13f9de9ff3cb9d9363bd90d1f9deaab4dc9cba2d43a2ea"></a>

## Next pages — stateful_service.advertise_options.advertise_in_cluster.multi_ports / f0f48e61c0fd / 4

- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-020.md#canonical-0acbd11f0964dba2dd1a73c816c866bc2d5a2edf20fbcd7b1040639b08fd04f2)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-020.md#canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-0acbd11f0964dba2dd1a73c816c866bc2d5a2edf20fbcd7b1040639b08fd04f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-45bddffe3e3f7cb0e2a5c5d18e0dad9c92ba7b501498ff76c1335feeda7fad58"></a>

## stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports / 934341ade4db / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-020.md#canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-020.md#canonical-45d897135e869ec5d9fc84f9ce490d93d685d10b9b644dc0972be10c0d3b957e)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports

<a id="canonical-62d7a21921be428a591439d2104fac872e1b7551470e9184a03fae46a0415379"></a>

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

<a id="canonical-29f395af90cedc66bedba1e8a3ba866601e5080a30ae3e6b74610e0904654a55"></a>

## Direct properties — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports / 934341ade4db / 3

- [info](resources--workload--reference--group-020.md#canonical-b9cf6420a29688afd643677a24412f3bd529080959fa95647649e5aeba16de57): complete subsection reference.

<a id="canonical-203906333fb2f05d681f43f89076b8b78f49506f9775314e7f9b16c1daf67745"></a>

<a id="canonical-310418d0ec39692d0aecb3f79920da6c72d0c735a1069f06b2a31497f3001c6e"></a>

## name property — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports / 934341ade4db / 4

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

<a id="canonical-964f059562ec4a330caf557e90120650d4ff11058f40e6d8dafd3d16bd9e8dd0"></a>

## Next pages — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports / 934341ade4db / 5

- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info](resources--workload--reference--group-020.md#canonical-b9cf6420a29688afd643677a24412f3bd529080959fa95647649e5aeba16de57)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-020.md#canonical-45d897135e869ec5d9fc84f9ce490d93d685d10b9b644dc0972be10c0d3b957e)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-b9cf6420a29688afd643677a24412f3bd529080959fa95647649e5aeba16de57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-225f7c66f11030914954cd3d1253eeb2dc43c55176ea359ce41bf6e5c92b6bbf"></a>

## stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info / b02b4fd52f70 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-020.md#canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-020.md#canonical-45d897135e869ec5d9fc84f9ce490d93d685d10b9b644dc0972be10c0d3b957e)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-020.md#canonical-0acbd11f0964dba2dd1a73c816c866bc2d5a2edf20fbcd7b1040639b08fd04f2)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info

<a id="canonical-2073ecbeaac46858b4a0ef3fc2a994486ccb917ad018ef818bd7295bac1e3270"></a>

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

<a id="canonical-fefc8b8d1b9f3ead2a1ba20f6ccd97de23e9499e65d37de6fb0ce1e78b87f5f5"></a>

## Direct properties — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info / b02b4fd52f70 / 3

<a id="canonical-dca63e63c903bb5a7b5f82284073067165c9ba2c8a943c33b65aed6597e03212"></a>

<a id="canonical-6daf1cd5767b1c22f9d0f6911e335ec444087752a6240ad2a848474365b842f1"></a>

## port property — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info / b02b4fd52f70 / 4

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

<a id="canonical-b8146e357ff2d792537018f1c83bfdec81bdba3328c1d8dabb404a06685f1399"></a>

<a id="canonical-27f764159b6d70b2bd47f144513b962c1937417dbd32226b9882a2281aacc83b"></a>

## protocol property — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info / b02b4fd52f70 / 5

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

- [same_as_port](resources--workload--reference--group-020.md#canonical-5354dc1c69d7599e5e84af94da69fd95129c49c1ebfe7513be2cf948adb43092): complete subsection reference.

<a id="canonical-4ce1fac5185137e848d7108a7d43b80834c2b07b906b7a9c8cc48d1afdeb7a30"></a>

<a id="canonical-fd0aeb9be595dcafe75a76e61d2bdf9113028359169deb8bbf5558889b1f97fb"></a>

## target_port property — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info / b02b4fd52f70 / 6

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

<a id="canonical-58b0c67c4856a172f462816467d454386711fa3b6ce7fc4489ad7cfe22a7b9ca"></a>

## Next pages — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info / b02b4fd52f70 / 7

- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port](resources--workload--reference--group-020.md#canonical-5354dc1c69d7599e5e84af94da69fd95129c49c1ebfe7513be2cf948adb43092)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-020.md#canonical-0acbd11f0964dba2dd1a73c816c866bc2d5a2edf20fbcd7b1040639b08fd04f2)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-5354dc1c69d7599e5e84af94da69fd95129c49c1ebfe7513be2cf948adb43092"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c07e89aee2ab0e4e617c22a5947203df6caec10652ee8449df59f959eaf4dcb"></a>

## stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.s / 4a83ce617051 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-020.md#canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports](resources--workload--reference--group-020.md#canonical-45d897135e869ec5d9fc84f9ce490d93d685d10b9b644dc0972be10c0d3b957e)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports](resources--workload--reference--group-020.md#canonical-0acbd11f0964dba2dd1a73c816c866bc2d5a2edf20fbcd7b1040639b08fd04f2)
- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info](resources--workload--reference--group-020.md#canonical-b9cf6420a29688afd643677a24412f3bd529080959fa95647649e5aeba16de57)
- stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.same_as_port

<a id="canonical-911ea1d994a729b32184d9089ae5685bbf906c71645580864480d30dba2d3508"></a>

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

<a id="canonical-980d92c6186032522585cc1ab34bb3854c093661984deb080867934857f2f6c3"></a>

## Direct properties — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.s / 4a83ce617051 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-231619354ad5d9a3db83b43feef0553bb292aa2bcfdd074e380816fd4b456418"></a>

## Next pages — stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info.s / 4a83ce617051 / 4

- [stateful_service.advertise_options.advertise_in_cluster.multi_ports.ports.info](resources--workload--reference--group-020.md#canonical-b9cf6420a29688afd643677a24412f3bd529080959fa95647649e5aeba16de57)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-396880e662d7e5243b3407d5c9e82062f2dfca8f6f9725b8d1db94a0ece1b5f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fbd3d7be391e2eaecc49cdcd0e9c3811b1542f6a2aaddb3e0fc7a7bd4b7bde1"></a>

## stateful_service.advertise_options.advertise_in_cluster.port — stateful_service.advertise_options.advertise_in_cluster.port / a93c01661617 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-020.md#canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77)
- stateful_service.advertise_options.advertise_in_cluster.port

<a id="canonical-d9efa07cd485c687cb91b819d94c83479591f3adf961914779fac9aac6a1c100"></a>

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

<a id="canonical-9507cc281355bd9d4041ace32d255007cdc73ada1e0836c9bb1780eb304bbdd0"></a>

## Direct properties — stateful_service.advertise_options.advertise_in_cluster.port / a93c01661617 / 3

- [info](resources--workload--reference--group-020.md#canonical-f668bc149b2247d849aac1e20cf7d42258c43c73e8955812c53e8c9b61840ace): complete subsection reference.

<a id="canonical-588d231ab0a128d310a902f6fbf94859d3bc71eaca36991765c6592995773883"></a>

## Next pages — stateful_service.advertise_options.advertise_in_cluster.port / a93c01661617 / 4

- [stateful_service.advertise_options.advertise_in_cluster.port.info](resources--workload--reference--group-020.md#canonical-f668bc149b2247d849aac1e20cf7d42258c43c73e8955812c53e8c9b61840ace)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-020.md#canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f668bc149b2247d849aac1e20cf7d42258c43c73e8955812c53e8c9b61840ace"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d29f1a114a2cb9d0feeb132aee0e464204ae303b71cdeff67324238bf7aa6fb"></a>

## stateful_service.advertise_options.advertise_in_cluster.port.info — stateful_service.advertise_options.advertise_in_cluster.port.info / 90de145438b8 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-020.md#canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77)
- [stateful_service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-020.md#canonical-396880e662d7e5243b3407d5c9e82062f2dfca8f6f9725b8d1db94a0ece1b5f6)
- stateful_service.advertise_options.advertise_in_cluster.port.info

<a id="canonical-afa549d1ff1ee209699957e809866f44d8683e91a74df744f4cbd34689d74cb0"></a>

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

<a id="canonical-e62654a9f66cc6603cb80f92651bfcadcd53113bc7630b468753d84fd531e58a"></a>

## Direct properties — stateful_service.advertise_options.advertise_in_cluster.port.info / 90de145438b8 / 3

<a id="canonical-28ee76b964bc508693491d57d6312b4bcfc12343182b8acffb3b1fe0ab2d57f8"></a>

<a id="canonical-c386e617060ec61414baf97d9c78b1348fd0811fe79dbca13ab0b93d4c21e6b1"></a>

## port property — stateful_service.advertise_options.advertise_in_cluster.port.info / 90de145438b8 / 4

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

<a id="canonical-b65dbbdc9896fca65c34f3d26e08a14852046ed345379c4bce607beec6f22dc9"></a>

<a id="canonical-d8e7c68f5c7f065bcff175b4ad10f426896d70a7278cf3a2476e09bedda61e0b"></a>

## protocol property — stateful_service.advertise_options.advertise_in_cluster.port.info / 90de145438b8 / 5

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

- [same_as_port](resources--workload--reference--group-020.md#canonical-e422053cff27de2a2c05c06fb3aa915a9db496376378b55638c8035488c34a57): complete subsection reference.

<a id="canonical-ba51461561389a486b816c4861834c6c0a0dcaf3072ddcc3029d87997113ba0a"></a>

<a id="canonical-6622978a472d611b3fb18320e793d3f547455e96f462cc0ecac01489ff0f53ba"></a>

## target_port property — stateful_service.advertise_options.advertise_in_cluster.port.info / 90de145438b8 / 6

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

<a id="canonical-f7ca4ad3b718329643f7b083e23bc989dbbc24f884a839d457eb95196c0b34f8"></a>

## Next pages — stateful_service.advertise_options.advertise_in_cluster.port.info / 90de145438b8 / 7

- [stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port](resources--workload--reference--group-020.md#canonical-e422053cff27de2a2c05c06fb3aa915a9db496376378b55638c8035488c34a57)
- [stateful_service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-020.md#canonical-396880e662d7e5243b3407d5c9e82062f2dfca8f6f9725b8d1db94a0ece1b5f6)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-e422053cff27de2a2c05c06fb3aa915a9db496376378b55638c8035488c34a57"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30a8432619ca15885cf59b8e4defde4769c35f89f1f8288893b77dfa209b094a"></a>

## stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port — stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port / 8e4671b91f64 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-020.md#canonical-a380f5505dd498de6e06d125c0596bf9c21d808f953ae1192a4909b606d18f77)
- [stateful_service.advertise_options.advertise_in_cluster.port](resources--workload--reference--group-020.md#canonical-396880e662d7e5243b3407d5c9e82062f2dfca8f6f9725b8d1db94a0ece1b5f6)
- [stateful_service.advertise_options.advertise_in_cluster.port.info](resources--workload--reference--group-020.md#canonical-f668bc149b2247d849aac1e20cf7d42258c43c73e8955812c53e8c9b61840ace)
- stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port

<a id="canonical-78137440bc473e867c3d898291fa47cc5b7cb92c3442ba22ed7e20b83bdd7613"></a>

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

<a id="canonical-9d791ae2b6fe15b6114798bb7ef193130c25afc9f75df26d832e15ac9401a282"></a>

## Direct properties — stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port / 8e4671b91f64 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-657238a96be164cb0b7ea94ef67b717732eba0de495ae054da641b430b8f5557"></a>

## Next pages — stateful_service.advertise_options.advertise_in_cluster.port.info.same_as_port / 8e4671b91f64 / 4

- [stateful_service.advertise_options.advertise_in_cluster.port.info](resources--workload--reference--group-020.md#canonical-f668bc149b2247d849aac1e20cf7d42258c43c73e8955812c53e8c9b61840ace)
- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)

<a id="canonical-f047f8b4ebb75db48c54f3e3454edd0b77518d5969376dc864c97c8f94b7d022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ab0938a8e9b9e2ab553b11cde3c841c3e0d44d9ede98234c5b6e1336a901f86"></a>

## stateful_service.advertise_options.advertise_on_public — stateful_service.advertise_options.advertise_on_public / edcc7e053bfc / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-35d9915f43ca5fff8ac85c03830d1af72c049eeb58d2d09b92098817c37219ba)
- [Property reference](resources--workload--reference--group-001.md#canonical-865b40863c3fcc5ef85fd2cac9a0cd4633a6573292d4a6166c678ff46e9a83dc)
- [stateful_service](resources--workload--reference--group-017.md#canonical-a1ae4fd0e881ea29023afda841e3139a5e765d0a27acd6a4f490cedbb1d9e2bf)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-6c0b3eab4eecf719c82b069716fb1369547612efce85679f0a98d0b0f3b5ec7b)
- stateful_service.advertise_options.advertise_on_public

<a id="canonical-ec7302d838dc013ea54ec61971c3c9e488f42b92a3df969b4e712e26cc069e3c"></a>

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

<a id="canonical-4f2d819640fc526cf7bc2175ad84811ad8a58ba856aa833366f18d8cc3c45aac"></a>

## Direct properties — stateful_service.advertise_options.advertise_on_public / edcc7e053bfc / 3

- [multi_ports](resources--workload--reference--group-021.md#canonical-505a640109552eff7be8a81e064f654f0c209e4b9b9a7a659cb44b3121f629b2): complete subsection reference.

- [port](resources--workload--reference--group-024.md#canonical-181a1518b8314bdeef4b5e79a2b8aacc9f75d1d5ed91dcbc4a205f9ad5b37385): complete subsection reference.
