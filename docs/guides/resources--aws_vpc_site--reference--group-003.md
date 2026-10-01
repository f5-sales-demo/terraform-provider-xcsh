---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-f84f3a0c5590124a7cc539e873d270d1f0206f2498e27428b2ab0caf8869ed38"></a>

## ingress_egress_gw.inside_static_routes — ingress_egress_gw.inside_static_routes / ab2cb62ca34d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.inside_static_routes

<a id="canonical-a5606b1c89d944d06ecaabc8f41092a81119f5574f205b165cd6b225c2e23f5f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
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
inside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-01ba29a8fb719c6c21e1b67674a9189986a28ca2130f2e140fdca03530669468"></a>

## Direct properties — ingress_egress_gw.inside_static_routes / ab2cb62ca34d / 3

- [static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd): complete subsection reference.

<a id="canonical-f963d7b4b7970ab65df99b05035637920653395a352c486c8bd27f6724094292"></a>

## Next pages — ingress_egress_gw.inside_static_routes / ab2cb62ca34d / 4

- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c8540172e5bef14ded1b08c80b065b2e347d686a643e29132e1a1229819276d"></a>

## ingress_egress_gw.inside_static_routes.static_route_list — ingress_egress_gw.inside_static_routes.static_route_list / b6900f8c91dd / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- ingress_egress_gw.inside_static_routes.static_route_list

<a id="canonical-c50d100519e47b3228832b654ea4551c1587ac15cb708723e4c481a9a114a90e"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-87e63ddfc228ab3a5d6d09a0c131f685b8c95df0097ebf16886e659a42c37313"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list / b6900f8c91dd / 3

- [custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82): complete subsection reference.

<a id="canonical-bb87f072a634453a021eac7bb787dc033adbc3a0847ce10d031f96cc5e65469f"></a>

<a id="canonical-3b61b07af2b7bc7b79a866c4f8fcce178202a157d7b5bfa2617429c89bd3bd0f"></a>

## simple_static_route property — ingress_egress_gw.inside_static_routes.static_route_list / b6900f8c91dd / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-af54c1fc121fe8ac627a25761fb76d90ef6acbc3cae8730cd4a0fe488f7909c1"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list / b6900f8c91dd / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-275dffd4d26f420ed1671bd34f7099c1de7d890f73a9e5e11dacab9a43c42a9e"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 29c3423ad75c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-6cc64cebefac3bf9ac7eaeff2ae84b123b49472482c4827c4aa3b84e0f97cbed"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
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
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-9093d2c4c001c52eea2aa98e3280881cf7ee1648fb61cdf65cd71542438d158e"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 29c3423ad75c / 3

<a id="canonical-71e910f387a21fa1ce9a2c1bdcef21f6c772f644ffd1938c7b6acb819f46067e"></a>

<a id="canonical-66650e16edb6d0cbdf6f969bfc2c1861cf20fe728f40c07c93bbd91785c02632"></a>

## attrs property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 29c3423ad75c / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--aws_vpc_site--reference--group-003.md#canonical-f62f059e8e7bcfc4da90bfa6f931d283a3d3216143e3297b8184f7db6de96817): complete subsection reference.

- [nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f): complete subsection reference.

- [subnets](resources--aws_vpc_site--reference--group-003.md#canonical-63f22d68a558d400ce64caeb11af03b6615abc7c1cb1015c579b5631f7f4eb5a): complete subsection reference.

<a id="canonical-bc9ede9b03616bc13dfb311039418ebc449859cedd95666a22ebebd79553521f"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route / 29c3423ad75c / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--reference--group-003.md#canonical-f62f059e8e7bcfc4da90bfa6f931d283a3d3216143e3297b8184f7db6de96817)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-63f22d68a558d400ce64caeb11af03b6615abc7c1cb1015c579b5631f7f4eb5a)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-f62f059e8e7bcfc4da90bfa6f931d283a3d3216143e3297b8184f7db6de96817"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dde7d741a1dcda1e1b4a762b20d37a6cdfb350db306602ea279d943cb7cc6f63"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / edb17568b8ab / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-99ad057a07128bf1a92c30779bf02d1ee2676e7ff763569612e330d40ef440fb"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
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
labels {}
```

<a id="canonical-11167bd4266ce4ad032812d8c8c8f0726a9da2ca4275e5006d4e9d9324d08c21"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / edb17568b8ab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-277e9ce7f6932286c4f83963a12774395c4c0be0a360497b7a90b80429b03a7c"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.lab / edb17568b8ab / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e45a5c6ade93122ec75b54917200d0056fcd1c47e8181c93cd6181e2ec0dac0a"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e82cfec317c4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-182bd0bee04b6a9e71c7b729458084e59a6192bf956686d79576fd4cbe151961"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
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
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-bf6a128b5588f5f16b0e4a8f78fac62853def8966e1db67f4f1226677d5a3aca"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e82cfec317c4 / 3

- [interface](resources--aws_vpc_site--reference--group-003.md#canonical-f77584c013fcfdd39d8d289502aa0688e910dbcf95ac7a6a2c8c1e4210f5734b): complete subsection reference.

- [nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-5591cfa51feef509e13d87120de2d22e1e50d7ff31d8ac9c8805c74f76198976): complete subsection reference.

<a id="canonical-7b56a56c928bb2775179dbb8b1aca9eb1a1b731e9252feb3075645fe349b301d"></a>

<a id="canonical-b15aba85472b2939c3ce72adf98d84c161c8248259a14407a8f17b878e6499f4"></a>

## type property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e82cfec317c4 / 4

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-da41084db9934e79fa5e19214a10d71b76650ef6a768305c55360deb20e9b9c0"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e82cfec317c4 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--reference--group-003.md#canonical-f77584c013fcfdd39d8d289502aa0688e910dbcf95ac7a6a2c8c1e4210f5734b)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-5591cfa51feef509e13d87120de2d22e1e50d7ff31d8ac9c8805c74f76198976)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-f77584c013fcfdd39d8d289502aa0688e910dbcf95ac7a6a2c8c1e4210f5734b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e576aec3d69478fe1de030df05d0169eb65ab9c87a93b2fd512897f7565c6ce2"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / f558617f8c23 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-ee4805351a57f430acd51e1d18019899709fcd74843ba8446ddc9d260ded627f"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-68c4ec3f8df81961624fda38a979f6cf51982f18a69a79f32aa3ff013804d64a"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / f558617f8c23 / 3

<a id="canonical-67f8cd4c959a9f8da0883bd05d6509d8101185f64a79d9a91c2e8f20857ef910"></a>

<a id="canonical-95e8a580b5f2ab83e10f9b542deb606bb32dcf0649d64b45a2730b916a862b82"></a>

## kind property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / f558617f8c23 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-bd079428cf3085f25be045f04ef3761b51220c969e0d97d785374525a8f75e07"></a>

<a id="canonical-6ffb5c5b3643cd337c943e309cff18728feb25007406ce9551d780521d40e7bd"></a>

## name property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / f558617f8c23 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-bcdf9c7bcbdf4ac7e1ac9c3c1093ca743b31fc9ee554864186909c0dee29a897"></a>

<a id="canonical-2c2ac4f67239c30db9c23d314a9978e44740c3287c8d5c57e968dcc2f5087a6e"></a>

## namespace property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / f558617f8c23 / 6

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
  }
}
```

<a id="canonical-162708a0a727c1351244d41d16ffa804979c98eea50894236870bcaf395291d5"></a>

<a id="canonical-4db266e43e037a43180bc9a00fe6ceb4e1ebd7344361792fd37b85bf074e886e"></a>

## tenant property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / f558617f8c23 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-313a4ee9deb415f24a011b4d83a318f46302b8e021a644b8ad5473e16f342b80"></a>

<a id="canonical-a6e27af2f498e5c87f0e862b9ab9612b31a802e3c1a05a0fa37f2c10de4b4786"></a>

## uid property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / f558617f8c23 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-107f4ae041589a60bb0b029f6bc4cb046c9aff89afc5947c067e63d7c181708c"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / f558617f8c23 / 9

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-5591cfa51feef509e13d87120de2d22e1e50d7ff31d8ac9c8805c74f76198976"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c43b94f7b542fab6e78725f423056ba5377f5e91f608a0c5eac8d45226b8822c"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 5b16f4b6f218 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-bec0fecdbafac0a8825f3f1b6913196a7459f909e8652f460175362fd05ab060"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-7ae2c1eacf9e24c4ac34d0bca0f66a378638512c64d7de827c092df82b6b3ffb"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 5b16f4b6f218 / 3

- [dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-205f72e54677f6e398fd8378892a9782bdc2ff09b4ad8bced89a1ad21715c762): complete subsection reference.

- [ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-e7298c81e7ef2c8881893527d7dca846aa1a85f8b44042ba63f1349ca0f1c37f): complete subsection reference.

- [ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-ac2fc076fec1d5c7554e5d2f5898c825377a922f5501566beb96994e583e760a): complete subsection reference.

<a id="canonical-0027cd7b508ea42557466840ea81d3cb1ffd3dde90e565381d6900ec04d94337"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 5b16f4b6f218 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-205f72e54677f6e398fd8378892a9782bdc2ff09b4ad8bced89a1ad21715c762)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-e7298c81e7ef2c8881893527d7dca846aa1a85f8b44042ba63f1349ca0f1c37f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-ac2fc076fec1d5c7554e5d2f5898c825377a922f5501566beb96994e583e760a)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-205f72e54677f6e398fd8378892a9782bdc2ff09b4ad8bced89a1ad21715c762"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d7a63ca76b62f28cf9f92d12657f4b91043ef1820c8b970406ae1aa17bce3ab1"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / f34715e7a34b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-5591cfa51feef509e13d87120de2d22e1e50d7ff31d8ac9c8805c74f76198976)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-35d9444279af1d592ce8b96610d2fc64490d56a1b29e1cc2bf6e6057e2d100f7"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
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
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-5799f6d7844092edfdea054cffd4cc14cda8c7394bbd2115f333f3ce70578b4a"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / f34715e7a34b / 3

- [ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-31f8b54c7fb80c02a3b864dbc49aad8c873a993071d97e6cb06e2a4c788ecd1c): complete subsection reference.

- [ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-0a0cd18855b4e4e8ab1620a6790006341a78e20f21367c31d6d7472cbbc57145): complete subsection reference.

<a id="canonical-749f7d3a3411a67aac4ac0b78de3fc600136eb83facb04363102d455030e1a78"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / f34715e7a34b / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-31f8b54c7fb80c02a3b864dbc49aad8c873a993071d97e6cb06e2a4c788ecd1c)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-0a0cd18855b4e4e8ab1620a6790006341a78e20f21367c31d6d7472cbbc57145)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-5591cfa51feef509e13d87120de2d22e1e50d7ff31d8ac9c8805c74f76198976)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-31f8b54c7fb80c02a3b864dbc49aad8c873a993071d97e6cb06e2a4c788ecd1c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e1ef79d4b1a43b7b133433bdff13c6d47c1e5685143cb137d0eb9227d954cd2b"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 482e1adc9905 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-5591cfa51feef509e13d87120de2d22e1e50d7ff31d8ac9c8805c74f76198976)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-205f72e54677f6e398fd8378892a9782bdc2ff09b4ad8bced89a1ad21715c762)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-a92468aceaa83423ed319a5e0680122244e8333aeea5d64dc384eb4f96b3c3da"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-b3abdd0295d5a61d6e720587c0a100f01ec98bec67fb9794b1ff434b78e8eb90"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 482e1adc9905 / 3

<a id="canonical-a1bc9647e2eb715cdb46cc3d01d562afd93b2906521b733c3d1c3c807b8aaeb6"></a>

<a id="canonical-772583a2c4e8363a729420e2867ae724996d80f04a8d22741157636d84d53191"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 482e1adc9905 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2b8afa134fe1e5d45ea4aa65dc2354d025630934335a0767c51f6673b85e4d7b"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 482e1adc9905 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-205f72e54677f6e398fd8378892a9782bdc2ff09b4ad8bced89a1ad21715c762)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-0a0cd18855b4e4e8ab1620a6790006341a78e20f21367c31d6d7472cbbc57145"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-942ee420654b10fab84ae1dcb92b8f7429e6e2f709f29ca0ead52a37097c1ffc"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e6683c4dd0e8 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-5591cfa51feef509e13d87120de2d22e1e50d7ff31d8ac9c8805c74f76198976)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-205f72e54677f6e398fd8378892a9782bdc2ff09b4ad8bced89a1ad21715c762)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-381edbdcecfe3151472d8526cdf6a4fe720937176990736dca336a2b285fd9c3"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2751aae882e0508a8f244d337292d108497ea9577d9e65a20a3d908b4dc23453"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e6683c4dd0e8 / 3

<a id="canonical-a1ae9e5e9129f2ddb43b674a538308440e441c4909f5b36653aa81a4640bf648"></a>

<a id="canonical-0dad1a93964e85f37dd88ca75d7cca80242e9dfb3697400ba00cb00cbffec8ae"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e6683c4dd0e8 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-ad620e2c8a6c98d761655bb26ec7f2373ae40870405e0f7f6c960a2dbeaae05e"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / e6683c4dd0e8 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-205f72e54677f6e398fd8378892a9782bdc2ff09b4ad8bced89a1ad21715c762)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-e7298c81e7ef2c8881893527d7dca846aa1a85f8b44042ba63f1349ca0f1c37f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4aa2e16b72b832899f73b7d15d2afbd44a512bd04bcdc4ddd4ad59deb7febc7d"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2aa23bfbf831 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-5591cfa51feef509e13d87120de2d22e1e50d7ff31d8ac9c8805c74f76198976)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-31b1809dc53fd0bffe5be7691b8b410fe84993fbad7ad3ba965bb8c813bd609d"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-e9abda94b91410f422f7ba025c6293bc1761fccd9dad44b4bc2ab48c86e082ce"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2aa23bfbf831 / 3

<a id="canonical-767b8430aeb5df5879074ea43f73cc6f3517fdb92e80638c048b18ee3ad742bf"></a>

<a id="canonical-6f600c7f198c68594f6faefcde9cd277b0e58781b649f9091c17d8f386e76096"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2aa23bfbf831 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-65b730a9bd6b7424ee60db318287f3fec78dcc91696c435a789bc5694d74165d"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2aa23bfbf831 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-5591cfa51feef509e13d87120de2d22e1e50d7ff31d8ac9c8805c74f76198976)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-ac2fc076fec1d5c7554e5d2f5898c825377a922f5501566beb96994e583e760a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78e405a2d90f69c140bb2d9015f29ce5a72d3f3a7a1e1767654a48dbd4ecbd30"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2461b5d6a3e0 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-84d2b38463512fae24cac1b6336f0972d49158b24f79138660dbdd2a2a6e694f)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-5591cfa51feef509e13d87120de2d22e1e50d7ff31d8ac9c8805c74f76198976)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-7e45b8b14b7faaa45a1ffbc7b9fe8ad158b8b257c6461bc34815c56896f79714"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0bf69c9f334dc0cc4580133d67fe4b9aeb616c7e6bd5912edc4e30532c81a4a8"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2461b5d6a3e0 / 3

<a id="canonical-3d4fb5b8d2b870ed33efbd3ac235371023b07357b224b41cd81d2b24f4471f06"></a>

<a id="canonical-785be18e6424feba21e6a55977dbb10746a54481f883a7fcea3a1dbf06195c66"></a>

## addr property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2461b5d6a3e0 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-c4a4f7065bb803f2c6ac8d71b73cab5629d4529cd9fde889090fc2269e346363"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nex / 2461b5d6a3e0 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-5591cfa51feef509e13d87120de2d22e1e50d7ff31d8ac9c8805c74f76198976)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-63f22d68a558d400ce64caeb11af03b6615abc7c1cb1015c579b5631f7f4eb5a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68b3af29bbd82708ae5b5760d9e480966f36f3859d6d61284130908ac0cf6d86"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 92c4e1075c4e / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-943228c37b9e8d5b74cd3a9ece22f578ada06421ee425c92942209ca029405db"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-93405880ce16bcc54032d3f351316512dc534d7d74ec3afa2bef1c5aabedd31a"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 92c4e1075c4e / 3

- [ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-4291b6444469c48bd2527f6f45836c06a23e811f62c76fe9063641b245130985): complete subsection reference.

- [ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-1683774089788acda688a2802a4ea856fc24d483cf675a99b65d1a1cd3c64680): complete subsection reference.

<a id="canonical-566fb61c1c1e252e3c12c9f42f174584e2d76d39d6a6653ad1fde1dbbea73e6e"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 92c4e1075c4e / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-4291b6444469c48bd2527f6f45836c06a23e811f62c76fe9063641b245130985)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-1683774089788acda688a2802a4ea856fc24d483cf675a99b65d1a1cd3c64680)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-4291b6444469c48bd2527f6f45836c06a23e811f62c76fe9063641b245130985"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4fb1bf0e9774232cf80d34a3c73571081f97798b8b6240a405ea58952898bdd5"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 777a758a5c6d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-63f22d68a558d400ce64caeb11af03b6615abc7c1cb1015c579b5631f7f4eb5a)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-8d5223e129a1aec3639c924d68aa2a7f74545113f2f6fccd4b977d1aff3119ad"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-468118d86cdef221fc63817b9209e29d55aba88cbbdfb8617b914d5d615fe763"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 777a758a5c6d / 3

<a id="canonical-0141159f9caa23578ce3f4f65d98e3d7193327d671567cc415d05fee10e13bd7"></a>

<a id="canonical-d9c93f6a66ae45989082091564395106a8728cc22afd488442032a2c5d722e63"></a>

## plen property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 777a758a5c6d / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-140ff040033ffa09be72f8b57b30b528ce9fef250fb0c10c6267739f395c6f4a"></a>

<a id="canonical-3ad3f26a117e09298f6bb5577c5e29bf65ffdca1a8e240c0111cdd4c1337d998"></a>

## prefix property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 777a758a5c6d / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0e7346d2d66ba97aaf7adc94f48e8c752dff92a7e9d601883e349f1580fdbe2b"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 777a758a5c6d / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-63f22d68a558d400ce64caeb11af03b6615abc7c1cb1015c579b5631f7f4eb5a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-1683774089788acda688a2802a4ea856fc24d483cf675a99b65d1a1cd3c64680"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a07422e22d0b3987bced7ced326b2cb407b5915496009a53c5b36b81d96b5e78"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 644c93bb2b21 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-002.md#canonical-1afb25b9eb991ec8584c6d82821d18e17630c3578c77cc41ce2fb1774c6abf4e)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-9a2bcf5d5c738854af32a44d341bb72b52da6ced2dc24425ef5aa2deeb952abd)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-a4460fb664a0ee9064c3bb332b141bcedda29c884b98a9301cfcdeed56c52f82)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-63f22d68a558d400ce64caeb11af03b6615abc7c1cb1015c579b5631f7f4eb5a)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-8a84201813ee223da47e2dced75059518443e6a95e5e3b6490a59f662025dc75"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-5555a32539d0fa61001682a412fb1e603a074c81a8c837db342d124760ff61db"></a>

## Direct properties — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 644c93bb2b21 / 3

<a id="canonical-86762f21cc00edf46fd6f4a3480782549106a4c1484446addcf336755983d969"></a>

<a id="canonical-915092acb50ccd3ea9ba3b16445815ffba35bb30f03032f27fd609c3c26edaa4"></a>

## plen property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 644c93bb2b21 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-596951496729ecf35c4293b67f66291f0b8daed0884d91aa5a46c96d4a9e85ea"></a>

<a id="canonical-f2a0300f219f9049a54809c3de749967708aeadd8df28d7730ec42fc7c109efd"></a>

## prefix property — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 644c93bb2b21 / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3771694aa684c0924b98f1bb1ccd206db3f9e7aee242b099b91c1669f50b57b7"></a>

## Next pages — ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.sub / 644c93bb2b21 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-63f22d68a558d400ce64caeb11af03b6615abc7c1cb1015c579b5631f7f4eb5a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-925db483a8d11cc7301a046cbd60e34288e2be0b4b2956b9c174f80e93f2cb80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4212bf5fc0213973092bd54a5989b7ce667c7eb8bbc6ee0fb71f68d65d83c61a"></a>

## ingress_egress_gw.no_dc_cluster_group — ingress_egress_gw.no_dc_cluster_group / 2f75cdea7e31 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.no_dc_cluster_group

<a id="canonical-d6773f99d42ee2de2b1e848da04bfbdcc52131ab61cb4fcb6069b38b847299cb"></a>

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
no_dc_cluster_group = {}
```

<a id="canonical-e6db75a30c98902b39d467919eb46ddd97ad2a4b3bc109b0e469a1d29320b98a"></a>

## Direct properties — ingress_egress_gw.no_dc_cluster_group / 2f75cdea7e31 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d2e11cc340fa229b36f441ebf2934ecc03d8f2d58460eeaab5ceabf0661399aa"></a>

## Next pages — ingress_egress_gw.no_dc_cluster_group / 2f75cdea7e31 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-304079dda7dd5d59b838510f594d8679fe4d9b033527672f60df99dfe6b3e501"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86fe8e3ecd4a38bc67dd1f3ef0b72e9d1ade16bb01ab4be0bdc63dd422e41747"></a>

## ingress_egress_gw.no_forward_proxy — ingress_egress_gw.no_forward_proxy / 3beb18ed2e52 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.no_forward_proxy

<a id="canonical-26cf3c496a59097682f36cce9bba072eb0216cda05d2945b5031c2e8d114a5d9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no forward proxy.

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
no_forward_proxy = {}
```

<a id="canonical-4b47b51d2db695067edd1208bf7f4b918b06318afa8767335cc6250f7a47c497"></a>

## Direct properties — ingress_egress_gw.no_forward_proxy / 3beb18ed2e52 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c2d87690ca78113f190ed7f4a58d90bf4570bf359b5c06b923daee697dda76e"></a>

## Next pages — ingress_egress_gw.no_forward_proxy / 3beb18ed2e52 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-b373260e05977dec450fb878a00d0aac7028b0bdcc82fcffee8f6ede03a51058"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c11a70ddfcd482708ba76652985c6aa425491e1be85484ff1d15445d5f3d2783"></a>

## ingress_egress_gw.no_global_network — ingress_egress_gw.no_global_network / 5ebd574b0fe3 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.no_global_network

<a id="canonical-7880605a489b3619ce26517dd1c5092ffa1c25b32a4e04e83b02b56658bd8f1f"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no global network.

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
no_global_network = {}
```

<a id="canonical-7a7d330114de3a5f0d1fe978eda50fa2277c44c03a136dcfb070134dcc460251"></a>

## Direct properties — ingress_egress_gw.no_global_network / 5ebd574b0fe3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2449a2c0abd5bc44edc0ca27d415cd4e7e92ce22aa28f561e2a9a860be16878d"></a>

## Next pages — ingress_egress_gw.no_global_network / 5ebd574b0fe3 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-7f08c3dd89dc58192ae4aabb1edee5c4690eab8217a210fb8632de2eb1cf4aa5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0cc6a74417cca143257c21fbfce16920c9524b48d6beb35f322a71bf47ca4684"></a>

## ingress_egress_gw.no_inside_static_routes — ingress_egress_gw.no_inside_static_routes / b127cd010c92 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.no_inside_static_routes

<a id="canonical-190dab5db21f8b7a151c256c051a3d766ef6353fb68e52ad93a31d79c4fcc6f3"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no inside static routes.

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
no_inside_static_routes = {}
```

<a id="canonical-7704d6b5bc6df7d0889cb07c21e2914246a8be2a6a70ebca2a2c24dceb56130b"></a>

## Direct properties — ingress_egress_gw.no_inside_static_routes / b127cd010c92 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f9ca7dc820b64d72c1f368a34b64172859b0a2768be689e98a32c7319f11e2f7"></a>

## Next pages — ingress_egress_gw.no_inside_static_routes / b127cd010c92 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-3f8d154444742593f7621c73fc280a8494f1eda2ed613ddca74c3f4fa208a26e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddcaa502769e807a578a0d71f89543eb3b7522ff4797fdd154b68be1bd651d24"></a>

## ingress_egress_gw.no_network_policy — ingress_egress_gw.no_network_policy / 5db798ed0778 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.no_network_policy

<a id="canonical-84b626f40401db77af6111d03bbdf845c3e9f0e5e3ad34f34252686db60f5032"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_network_policy = {}
```

<a id="canonical-a93f958f570def3a9156cb1e2b66bafada8084bba20717dd0975da0c1c5dce1a"></a>

## Direct properties — ingress_egress_gw.no_network_policy / 5db798ed0778 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a247ef7bd8a1ba19b4ce35e559c57661b37471a40a613fb3d3eef7e7842b2a2b"></a>

## Next pages — ingress_egress_gw.no_network_policy / 5db798ed0778 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-e6a7038350d6bdb13a48ed189a70488534b464456f1cd2c9ec10ea58a9a3dd51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a288887990122738e43a2c42ccee85e6555b793f2798ffe5890ca5dde90d8535"></a>

## ingress_egress_gw.no_outside_static_routes — ingress_egress_gw.no_outside_static_routes / ba9bc6fb5c3d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.no_outside_static_routes

<a id="canonical-a2103abfb7a59716f79ff41db6489204399823c85742c4412864bd9e9c8a457b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no outside static routes.

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
no_outside_static_routes = {}
```

<a id="canonical-2c9a8c26ea211f4c58adb62b8e6c13832bc23dafc6ec501790e6fcfe3a58695e"></a>

## Direct properties — ingress_egress_gw.no_outside_static_routes / ba9bc6fb5c3d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-24bc8259faca02592ebf73cb04372ff9894abd385c84c3e14e6a7e47313e3366"></a>

## Next pages — ingress_egress_gw.no_outside_static_routes / ba9bc6fb5c3d / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd099b4046bea4dee43ef1de133b0e815f44b7d2c2d87907cbf32b1685ee148e"></a>

## ingress_egress_gw.outside_static_routes — ingress_egress_gw.outside_static_routes / 4fc36b7be230 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.outside_static_routes

<a id="canonical-ffd7063c32f5664763e7d19290177c80948eda25b25c2caf7b29f5fd553a1ad8"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
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
outside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-dc823c262d891223dd4c461a2d57dac6694c38488fb228145571899d865f66e2"></a>

## Direct properties — ingress_egress_gw.outside_static_routes / 4fc36b7be230 / 3

- [static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f): complete subsection reference.

<a id="canonical-fb39b8c36d1c9de11e12510df31effad77a5564a42c85bd15a3d1af3172345e6"></a>

## Next pages — ingress_egress_gw.outside_static_routes / 4fc36b7be230 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b20bdc0c1f3c70be9b91f1d404d746810303b74ef878e668537b3891e734922"></a>

## ingress_egress_gw.outside_static_routes.static_route_list — ingress_egress_gw.outside_static_routes.static_route_list / 0df4624a228e / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- ingress_egress_gw.outside_static_routes.static_route_list

<a id="canonical-3df3a03596be2ce68ba780d4bdee89340cfc298ca293e24f00ad74732ec79b12"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-30d45675f53dd128da019a6aed1f8ba957356ccc23dbc4b3eea31c9d74e5ea2e"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list / 0df4624a228e / 3

- [custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995): complete subsection reference.

<a id="canonical-f0894809dd82ab5c0360d4db6f012b77c551951d3d2fb8f4b20378d73500944e"></a>

<a id="canonical-63b7b789a7eb296e5a4d95ca0a5b7a2512a2aaf4c3f725297f08561ff8d65808"></a>

## simple_static_route property — ingress_egress_gw.outside_static_routes.static_route_list / 0df4624a228e / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-b755e369a2a1f2bd0936958c48c599d69cde0b10f7e92da890f2fc341ef501b5"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list / 0df4624a228e / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa47806638a9a1a2aa9565642d0e3120f396703fb34505b048c728945b6f13b2"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 993079c665ba / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-d244c24557b2ee4400d26cf0bcf20962677945b17d676f77d51d0bc232e68a0e"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
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
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-a80e1d708273615ee5b1db04c5047c4974a7cb78f7c8c4f5f683a88a06416e2b"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 993079c665ba / 3

<a id="canonical-65259d049942a3dc6a9b772b1afb811404b6dbce21e021dfc80d56995afec8f8"></a>

<a id="canonical-83065508880f376ad682f576668972244c05b162c7e899fa9a60bf2a1c7cc5a0"></a>

## attrs property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 993079c665ba / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--aws_vpc_site--reference--group-003.md#canonical-c34cb4951ac4800b9f5287b9fd7d873fa430b6cdd813f082f10bf0bbbb2d75e2): complete subsection reference.

- [nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda): complete subsection reference.

- [subnets](resources--aws_vpc_site--reference--group-003.md#canonical-8da9bef164795a01b309da45e1d5c9f03460357a80791a6b29ff96dd014d2cbd): complete subsection reference.

<a id="canonical-f894b943600446a76c2f33451ea563c18e36204bdecdd2adf089a00afe18a1a4"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 993079c665ba / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--reference--group-003.md#canonical-c34cb4951ac4800b9f5287b9fd7d873fa430b6cdd813f082f10bf0bbbb2d75e2)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-8da9bef164795a01b309da45e1d5c9f03460357a80791a6b29ff96dd014d2cbd)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-c34cb4951ac4800b9f5287b9fd7d873fa430b6cdd813f082f10bf0bbbb2d75e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7d02d38b7fbbbc1f42354f2b81be24f71f8c6f00e44aace53e83cece22d5630"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / b07b677f5da7 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-dea97b8d643120497762ccd4cc7fb6afbe329833eb2968cb8e383cacdcd88033"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
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
labels {}
```

<a id="canonical-7a59ef7e3f5653adfa670c3737f64bf329bcbc98c58c0a7f11f8841aae91f871"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / b07b677f5da7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d93aac8622c96ff4f0ce35e9a8ebabfd1fbc84b51961b7e77e6af7d5f0bca937"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / b07b677f5da7 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d5208b9b5e7d78ce010ae86c1f7632f092eae439b0da90c6fe9ab875941e046"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / b44237979a7b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-c06dd6dec234901f45d0ec9cf91961922280a5aa07e7504e652490ccd476c896"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
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
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-688261de2732165091e644030b01f08c415de07d093d7f4fadd1dfb8a459c5d5"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / b44237979a7b / 3

- [interface](resources--aws_vpc_site--reference--group-003.md#canonical-8805a6a0d01ad65493ef00391657f509f0c108959736bbbf4a396d7feb1088e4): complete subsection reference.

- [nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-dfffc81cf265b7fcec9894391dfb2d000f8084f9abc15b0338d1e0c183089e83): complete subsection reference.

<a id="canonical-8c6cd3f779d74dbc55c5b8b273d876e2b6c11f0190d08bb27f90e908a242471e"></a>

<a id="canonical-d8b318548ef9040aa5c8478af76f6290846949b5c050f7e28a1984aacb85c08d"></a>

## type property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / b44237979a7b / 4

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-664a9377dd0321ef3cea5ac3bcaa8944d08c44d5b559bb71af8d7e1b8c651d4b"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / b44237979a7b / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--reference--group-003.md#canonical-8805a6a0d01ad65493ef00391657f509f0c108959736bbbf4a396d7feb1088e4)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-dfffc81cf265b7fcec9894391dfb2d000f8084f9abc15b0338d1e0c183089e83)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-8805a6a0d01ad65493ef00391657f509f0c108959736bbbf4a396d7feb1088e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7adfb47b72995623935dd3ffd7089bff4f45b8afd9ab35748588c9255959850"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 27d1c672c117 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-45fb85ba639b337583e617ab21d6b6571b6ad7f915da24e3c25f3783b9a907b4"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-423505780f5a7cee6533570936f98ebd7e56e0f0cca7d2667f9ab5152a81bb75"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 27d1c672c117 / 3

<a id="canonical-965dc1f7372427da2ecba4e3af9944da279332e00c11f4248fa39e230ef82d46"></a>

<a id="canonical-6706028dabb7943a1bdba1fcd9e425c566fd1b672661dcbadc9e510a4a89028b"></a>

## kind property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 27d1c672c117 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-22e00cbe093a9d255a3f0f3f6d212f63f96da3c255a5ca5244e2067768150925"></a>

<a id="canonical-a67adc4299b2a9519a2d5d789314f999376cfcd9de38000e5e095eef952df602"></a>

## name property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 27d1c672c117 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-991be531ebdca6fb47711465e70e986dfeb28b8245d11182e20cc89e37c8f271"></a>

<a id="canonical-05b0b07ce397f33316677b419c47fb47c541ec2701acf7fcddac6b4974d4318d"></a>

## namespace property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 27d1c672c117 / 6

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
  }
}
```

<a id="canonical-4d098eeddf8077b640703b0974197a61dd30fc286860ea16122711aceb0f1191"></a>

<a id="canonical-b7797fffaa5a7485ec03ada97ef7c9ad65dd74ce01da79a421ac0528c1cae6d2"></a>

## tenant property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 27d1c672c117 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-4b808fd6dcf10c03da556d92eabc460679104ccfcc17d1058c531c20aa2fb957"></a>

<a id="canonical-5dad961cfbd1298998a403dff1d2f8b1781489b1e7a3ecf84c1e767c73288390"></a>

## uid property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 27d1c672c117 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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

<a id="canonical-75aff037b58c2d4e3d794020179f5169867c6237d37fcdd9a45a81666814181f"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 27d1c672c117 / 9

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-dfffc81cf265b7fcec9894391dfb2d000f8084f9abc15b0338d1e0c183089e83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4cce79932a0c5602b01cacf7a641c8fa6ea651e365e62dba14e373600365cc94"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ad123ae8719d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-215f73204f16bac87d1b9b5abceba8068a1c42137e34cd4592959d6217f74e11"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-7fd592c8810fdc27ca9dca4691b1ba9d20d2b7f861f9a3a09e15b63b98ae58e2"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ad123ae8719d / 3

- [dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-8858e4fc945f6c0f9ec0611044d91618a95ffc6e544d13182ae0d84d4bc21434): complete subsection reference.

- [ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-eb1e5a91497d9e865169b8200d751aad647aec51fb6de6f2adf2893509869512): complete subsection reference.

- [ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-d322946330d4800a75bd4709ffb368d48586b8efefc1d35f5b9d49000b071eb0): complete subsection reference.

<a id="canonical-9b25e6f0ec2cd95e20e28e8eec5fcdaedefad711b3e984db37e81b7274f557e8"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ad123ae8719d / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-8858e4fc945f6c0f9ec0611044d91618a95ffc6e544d13182ae0d84d4bc21434)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-eb1e5a91497d9e865169b8200d751aad647aec51fb6de6f2adf2893509869512)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-d322946330d4800a75bd4709ffb368d48586b8efefc1d35f5b9d49000b071eb0)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-8858e4fc945f6c0f9ec0611044d91618a95ffc6e544d13182ae0d84d4bc21434"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7230a36147fc995eea681dca80a14a6be9efa7c4e180c681de615fcb01dde442"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / e0255970c06b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-dfffc81cf265b7fcec9894391dfb2d000f8084f9abc15b0338d1e0c183089e83)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-a2f5b97f3e7bc049d9a8c28458802614e465696da2d9ad7f8bd30ac409ce4be8"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
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
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-154ae5f8be14a3f46692b725fc43bff5e5d0df49779988202cc5de57e40d709e"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / e0255970c06b / 3

- [ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-450a510f9e0eabb755750da0d9e20d5801585674921c8268bcafc882e9b2401c): complete subsection reference.

- [ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-0514c96a6934735b246751fed014bac120ad414013c92fa0100a66f30690da5c): complete subsection reference.

<a id="canonical-9eadc865a44503eb9748476f8ffbdd9631764e4bcbb0e3a978e5a51b458fdbfd"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / e0255970c06b / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-450a510f9e0eabb755750da0d9e20d5801585674921c8268bcafc882e9b2401c)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-0514c96a6934735b246751fed014bac120ad414013c92fa0100a66f30690da5c)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-dfffc81cf265b7fcec9894391dfb2d000f8084f9abc15b0338d1e0c183089e83)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-450a510f9e0eabb755750da0d9e20d5801585674921c8268bcafc882e9b2401c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3d3a490e7c0aefd8ec3f3075bdead12f19ab2f566896eca91906fb4fd3b184e7"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 524d776b4f02 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-dfffc81cf265b7fcec9894391dfb2d000f8084f9abc15b0338d1e0c183089e83)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-8858e4fc945f6c0f9ec0611044d91618a95ffc6e544d13182ae0d84d4bc21434)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-3bd5d9fe4ca016cf65f1afb1cca07203e031d6b8d9df4bd8f0a20fac5f227b0c"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-7b749b4d0486a641965d33bd665c9ffa0d78131348efeeabcefd717c2e1f9c37"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 524d776b4f02 / 3

<a id="canonical-f1ce4c631268d57a97e2cfdc675a9e7982e377cd9ca2e34f21b465179497e420"></a>

<a id="canonical-b077b7a22b0aae7150d8fc2ab6a3d172fbd140d9a36dcbbc377a463a1ae06238"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 524d776b4f02 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-36b9bce206e088c996a16cb3d804c4d8328dd2bd613f5f072aefc02668bef5a6"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 524d776b4f02 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-8858e4fc945f6c0f9ec0611044d91618a95ffc6e544d13182ae0d84d4bc21434)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-0514c96a6934735b246751fed014bac120ad414013c92fa0100a66f30690da5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6dba2cd0d38c093ea2b213b53bb3f930c58bd425e04ffec105ce9c2cf10d443a"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c91cddecb9d8 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-dfffc81cf265b7fcec9894391dfb2d000f8084f9abc15b0338d1e0c183089e83)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-8858e4fc945f6c0f9ec0611044d91618a95ffc6e544d13182ae0d84d4bc21434)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-597d8929af9fe9ec5df440071a1f0e6b279c618df863f0bf2f40d71d5f03cfaf"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-200bbccab094950d26dba3ec692468df185a00a274622108d47c56bc76d8e259"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c91cddecb9d8 / 3

<a id="canonical-7c337cec75ea276d0862e71ba582270a0bd79b71f99549e248535c38ae32492f"></a>

<a id="canonical-e193dadfdcd5c6b1163d094f2635c423ae156fed76d78472d8c4993fb2b9150e"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c91cddecb9d8 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-7921e1fc3cd5d961f5908cb3d405f1145335cd49b68a107072a3933c19242367"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c91cddecb9d8 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-8858e4fc945f6c0f9ec0611044d91618a95ffc6e544d13182ae0d84d4bc21434)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-eb1e5a91497d9e865169b8200d751aad647aec51fb6de6f2adf2893509869512"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c31ad3a4a3e0be5390512fd87ca1b24aa3ddc0eb45280a9e533b34f83e2b96b"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 699c44a7d3c9 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-dfffc81cf265b7fcec9894391dfb2d000f8084f9abc15b0338d1e0c183089e83)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-1d8a1a2671855cdabe3325ef93a8519aaf2a76017c82e3f7c29dccdfe0cf1a5c"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-bfa890ca815de35cde9d8c361075ce58a48b65b2f0a00fa1097dde3464db0061"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 699c44a7d3c9 / 3

<a id="canonical-f6267bedbec12a85c3cb23cb8cf438811a50cf809ba6367c163635439082e730"></a>

<a id="canonical-7f116d7e1671159a9bb5c243306a40ab40462c89c6de010e7d3cd0de5dd60b3e"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 699c44a7d3c9 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2d547a6bbf694a686cd6e89a1bd1df00d07b1bf44f8c606c71ffe67d4ae4392f"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 699c44a7d3c9 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-dfffc81cf265b7fcec9894391dfb2d000f8084f9abc15b0338d1e0c183089e83)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-d322946330d4800a75bd4709ffb368d48586b8efefc1d35f5b9d49000b071eb0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c358f57ff9f5478901d4e61d2d6a5413044b5d5907930f7e840613d29aaa31b"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / fade94b00348 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-86d5b3d9c00d22a157ba2e6db3dc2449ac54456c3364e4d499771280b5abfcda)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-dfffc81cf265b7fcec9894391dfb2d000f8084f9abc15b0338d1e0c183089e83)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-8c4def8d382958f7ca1496bccca59dbac1f4af6b136f1f151697cd57705f4c2b"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1a501b1f68413774bc75e611b574d1f4383516ad224d3300a26b211dadef0a7b"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / fade94b00348 / 3

<a id="canonical-66e85147055518049801c20bb408ee177af87528a7af8e16f14c69a5eecdb722"></a>

<a id="canonical-31996bf7e8494936fc2c85cb8f82bb7d3cd230e7b2653afa6bc0c88f6119a8f4"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / fade94b00348 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-fe602c83dbe06a75e9ce9903c1a6a47361a8a79167fc4f8782d98b58b52f6176"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / fade94b00348 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-dfffc81cf265b7fcec9894391dfb2d000f8084f9abc15b0338d1e0c183089e83)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-8da9bef164795a01b309da45e1d5c9f03460357a80791a6b29ff96dd014d2cbd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51841cf5f49086f3875f18d11eed193c4047fa55b09c9f25eb9405b10c6a3312"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 95b3565563bb / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3e6899f1d0812d39deb9b8b952c82fa774e26b1475a9a3536270293dca13fb5f"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-844a7027177ab811885287f47bc7599410088efae2453c8a4bab46721281be48"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 95b3565563bb / 3

- [ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-aa279a663829768deee52061f29f85279a89c4aa37b74792e86ea43b3a56096f): complete subsection reference.

- [ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-40e36de51943010ed36edc594b78d8617b8098939bddeb59ea480e2a5a3a827a): complete subsection reference.

<a id="canonical-3882c8929fadd2754a288f900a8b203ace16b861440ed76114eb855d73dc9dde"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 95b3565563bb / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-aa279a663829768deee52061f29f85279a89c4aa37b74792e86ea43b3a56096f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-40e36de51943010ed36edc594b78d8617b8098939bddeb59ea480e2a5a3a827a)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-aa279a663829768deee52061f29f85279a89c4aa37b74792e86ea43b3a56096f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47d32e7dffe756e29173b14714de4cab268a8fc0a158c232eefdddb84be6689c"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 70d891e70712 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-8da9bef164795a01b309da45e1d5c9f03460357a80791a6b29ff96dd014d2cbd)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-728cde474a74a3ca766bfb85c37780dc5b3006316a024db6f96727cc6a3c7c76"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-050c647e2f058ba727568df2de503db005659821e6d09aba532ee6f95934cf00"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 70d891e70712 / 3

<a id="canonical-93e4a3c0b93465869c57c442210eca3dca88f1b1555a8e1ad86bdff508a7801b"></a>

<a id="canonical-40d1d0ed6f5cae2ded22bb08983bf2675dba943f86bcc9be83292c2160e1903a"></a>

## plen property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 70d891e70712 / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-f65720438783c81ae21e41f61dafd5e457abf5ddc2c9e86ec6311c8d20f4bad7"></a>

<a id="canonical-ee64fdee4f3f695ab28760230e3b946bdf5ff0ef6ee78a1c2c61026556aea65b"></a>

## prefix property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 70d891e70712 / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2e7e9d115ad2a966154d2572df83b52e33548542cd6aff2f9c60c9675fc86623"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 70d891e70712 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-8da9bef164795a01b309da45e1d5c9f03460357a80791a6b29ff96dd014d2cbd)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-40e36de51943010ed36edc594b78d8617b8098939bddeb59ea480e2a5a3a827a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6f297e99b76c73c3817a1cfb4b13f08532d39142062e7f991295438c9db9232c"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / af855d7cf5a7 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1ea0761fe40447fc341a79661e00c0840e83fdb02340b9cf43ceee94e82b7b82)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-c7acefded25f486cc5d966b0f30ce0cd784d41de94bd9721ebaf2ccfb527fb4f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-46b5e9b2ba5be6b9e79f76f876cbb3992aaba7a52cd67a6a1de4d2f492e8e995)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-8da9bef164795a01b309da45e1d5c9f03460357a80791a6b29ff96dd014d2cbd)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-e0eda02837310212f8b8752aa89c87a9faa2b168f99ec58abbbb6d7977c8ba93"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-528f621b11b854a3a0b97e30816c05f51ee139dd7e941ca2b595cc4b61eed8ba"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / af855d7cf5a7 / 3

<a id="canonical-305815cb695398bc8e03096c55fbb7b288c355fece775f9f8cf25c10a1ef0869"></a>

<a id="canonical-a4fa89ea35a70caf3109f8a0266091f3bf387200cdea8216cadd90e6770eeb7d"></a>

## plen property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / af855d7cf5a7 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-589adb9e3b6a96dea4231b9a76361866d893f874beea3e76f0ee5ba25e875bc5"></a>

<a id="canonical-f3f052a46ed2a92ec1e11689aed0ac5414570d89f1ab2a9ed4963b0ba0835f93"></a>

## prefix property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / af855d7cf5a7 / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-52218f35b1d5bc55a31ec9bf977c311e4eea70468923cd6db3556af3f8fabdfd"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / af855d7cf5a7 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-8da9bef164795a01b309da45e1d5c9f03460357a80791a6b29ff96dd014d2cbd)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-9f12e89fc9a44e49873a593cf1db009ec014c5c5b5caffc3b5c8ce1a25c68d61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34d67c8cbd0ab63326e4aa1dc9f28a45fc567f2d7c089d053c8206a6c41c1ef0"></a>

## ingress_egress_gw.performance_enhancement_mode — ingress_egress_gw.performance_enhancement_mode / 70c06722c4a0 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.performance_enhancement_mode

<a id="canonical-b06dd12e5b5cd5e86f9564c8f50a91a9212db50adb646f28a80aaa0fca75147e"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-8ef1d2c3e59b6ca48422f03ca50dc95dec82c9b7616a4fdf902c2f9b97d2bbaa"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode / 70c06722c4a0 / 3

- [perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-d84282dbd041405405881c4a839259a8821f6eca4b78a2543ac03593c94138a5): complete subsection reference.

- [perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-2ee504f1e3db3a4d9a55e77323c34dac294dda69b7c7c26ee0921adca957f567): complete subsection reference.

<a id="canonical-6d5e1968bf57edbb1faae19ce3dfadbb12541dfc991775f0d367eaa6b4f5a95c"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode / 70c06722c4a0 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-d84282dbd041405405881c4a839259a8821f6eca4b78a2543ac03593c94138a5)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-2ee504f1e3db3a4d9a55e77323c34dac294dda69b7c7c26ee0921adca957f567)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-d84282dbd041405405881c4a839259a8821f6eca4b78a2543ac03593c94138a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e635c2ed6b65991f2f37d522d6d001f1bba81273993b14828ba5a3ed642452d5"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 63673f973867 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-9f12e89fc9a44e49873a593cf1db009ec014c5c5b5caffc3b5c8ce1a25c68d61)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-d488a275cfd0d912bda8c1083137ea5d5a60ee8a3e679ba46a68d2e930ee77fb"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-79bf53716ab1b0fdb9e2f2c7c51a92694a0badde9b1cb3c3c6013160656d2c5c"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 63673f973867 / 3

- [jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-b562098e143ab3843671242a154f43e1c0732628ac6be70e92f01ad645f0d8b4): complete subsection reference.

- [no_jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-9829a83c1b525e8a1c3801bffdbf888507e21ac168d162735eb239202730aaba): complete subsection reference.

<a id="canonical-2298abd1c88c20d2cf91a4776aeb10e4966c7c7d32e6ba19308d6ddab6fb08cf"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 63673f973867 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-b562098e143ab3843671242a154f43e1c0732628ac6be70e92f01ad645f0d8b4)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-9829a83c1b525e8a1c3801bffdbf888507e21ac168d162735eb239202730aaba)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-9f12e89fc9a44e49873a593cf1db009ec014c5c5b5caffc3b5c8ce1a25c68d61)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-b562098e143ab3843671242a154f43e1c0732628ac6be70e92f01ad645f0d8b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e2e6b4a16b1effdbc87587b29aaceb0736a82f6da9630837902c7e75aaa42875"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 93236eb28ac5 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-9f12e89fc9a44e49873a593cf1db009ec014c5c5b5caffc3b5c8ce1a25c68d61)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-d84282dbd041405405881c4a839259a8821f6eca4b78a2543ac03593c94138a5)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-668e586a47df6cf8006bcbf478fd3849c4c5e76b47211b1affc1098570c6c198"></a>

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
jumbo = {}
```

<a id="canonical-079f178e204a9db93cb33a184a9fd262e29d2b43333b686e37dc8896c89f40aa"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 93236eb28ac5 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f9a777197e8d649370d9dea62ac767b7fda084c05e7e4a0d0259b9b4833ff4e5"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 93236eb28ac5 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-d84282dbd041405405881c4a839259a8821f6eca4b78a2543ac03593c94138a5)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-9829a83c1b525e8a1c3801bffdbf888507e21ac168d162735eb239202730aaba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1aa1ca01f9fb0b250e5f10e1c67fc432fb65134da2c0762595eed68f74e243e5"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 9b43e8a95d6a / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-9f12e89fc9a44e49873a593cf1db009ec014c5c5b5caffc3b5c8ce1a25c68d61)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-d84282dbd041405405881c4a839259a8821f6eca4b78a2543ac03593c94138a5)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-cbd193e1a954636156b52e0a1495b39d2d5bef036228778b324a99a85c43b62a"></a>

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
no_jumbo = {}
```

<a id="canonical-1dcc47065907ca334276467b5d31ca9d67d19f9166a471ac854073408ccf05ee"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 9b43e8a95d6a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-798ca80bf7e1c6e5d6c0dd06357154302331fe1f123ce193091d527db4e90512"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 9b43e8a95d6a / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-d84282dbd041405405881c4a839259a8821f6eca4b78a2543ac03593c94138a5)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-2ee504f1e3db3a4d9a55e77323c34dac294dda69b7c7c26ee0921adca957f567"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94e60190afbc20f836d589f7449bd0c85a4da6590fb22dad048422f37a929077"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / b9e7eba6d991 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-9f12e89fc9a44e49873a593cf1db009ec014c5c5b5caffc3b5c8ce1a25c68d61)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-91199c3fd0ce7467d7b9c2db23e862a5727a55a178826d824626537b50c8cdf9"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-20457d03e65afe4bb2436a8c7f1eb5196d2d6457063cbe7ac8fb738d9a8b10a5"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / b9e7eba6d991 / 3

- [jumbo_disabled](resources--aws_vpc_site--reference--group-003.md#canonical-e4bd78fa6313b891fabf436f63ee25385a7116c937fb0f58f0aad3027f1cd5f3): complete subsection reference.

- [jumbo_enabled](resources--aws_vpc_site--reference--group-003.md#canonical-ac7315802124b184cc372f3569a181f90dd6dc7ffb49c59a8863e9670e56e3f6): complete subsection reference.

<a id="canonical-f817ec44ca946e36271a24a209935a8e2accb49dab5a5f3c7fbffcb440704da5"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / b9e7eba6d991 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_vpc_site--reference--group-003.md#canonical-e4bd78fa6313b891fabf436f63ee25385a7116c937fb0f58f0aad3027f1cd5f3)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_vpc_site--reference--group-003.md#canonical-ac7315802124b184cc372f3569a181f90dd6dc7ffb49c59a8863e9670e56e3f6)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-9f12e89fc9a44e49873a593cf1db009ec014c5c5b5caffc3b5c8ce1a25c68d61)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-e4bd78fa6313b891fabf436f63ee25385a7116c937fb0f58f0aad3027f1cd5f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b488da3da4c45b0b19cd1ec6e4fc73b4a730f990f8d639a27fcf0bdd13d603eb"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 37f0fc421942 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-9f12e89fc9a44e49873a593cf1db009ec014c5c5b5caffc3b5c8ce1a25c68d61)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-2ee504f1e3db3a4d9a55e77323c34dac294dda69b7c7c26ee0921adca957f567)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-867572b25cc62986bba5858f3dbb56fe635993e64face12bd65a437252b63483"></a>

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
jumbo_disabled = {}
```

<a id="canonical-3395b0aaf79a502c3fb20bec77d98a949a7b319b0b0f3e1cfb9c5077daae667b"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 37f0fc421942 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7a8da33cf9cc3d939d87079f1a2a11ff83c383ac6d328a6870732242176607dd"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 37f0fc421942 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-2ee504f1e3db3a4d9a55e77323c34dac294dda69b7c7c26ee0921adca957f567)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-ac7315802124b184cc372f3569a181f90dd6dc7ffb49c59a8863e9670e56e3f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad6a45b59908c91cb07291236bb308eec77e9e2c935918a24d8b961aa9d3c9df"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / 0900c45cae8c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-9f12e89fc9a44e49873a593cf1db009ec014c5c5b5caffc3b5c8ce1a25c68d61)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-2ee504f1e3db3a4d9a55e77323c34dac294dda69b7c7c26ee0921adca957f567)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-33bbc6f224d7aa8abdaecdb13103f06c0168e42696baf2e5316e193a94eff624"></a>

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
jumbo_enabled = {}
```

<a id="canonical-03b9df239052a83a06e27459591be2849da07c26b8b07b835266432c90c37cb4"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / 0900c45cae8c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a66f3d97ca51f327db1f419309beaf0d7e74b9d03abef61fbc067aee994d9619"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / 0900c45cae8c / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-2ee504f1e3db3a4d9a55e77323c34dac294dda69b7c7c26ee0921adca957f567)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-8245ef1b9ec742f19a0f2e280e11739d80e7622dbdf423a8363c438ac16c8992"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-755d7b9aa4e7445116c34aa4fdfe66ade94cad4216590453477043c3c0977d1e"></a>

## ingress_egress_gw.sm_connection_public_ip — ingress_egress_gw.sm_connection_public_ip / 3a9893552741 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.sm_connection_public_ip

<a id="canonical-4f419771b4a9131069cad1b1b5864a8be99ec0dd2aac52e3687e1800ea84de03"></a>

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
sm_connection_public_ip = {}
```

<a id="canonical-36ea27d8d7b14bc4fa12126b979f8454542d3ff9650a1ae9e01540c7ec569fb1"></a>

## Direct properties — ingress_egress_gw.sm_connection_public_ip / 3a9893552741 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b650aed1c0291bcc3edac96bea33dd97484088c7b6d94555e60e3cf171d39b2"></a>

## Next pages — ingress_egress_gw.sm_connection_public_ip / 3a9893552741 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-8a09e1f2df9a35b9e17192f57980319cd9e580ea183832a4d46d0012d6e36149"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f3a664120c4c4790e18fec9cb9b0349a04b7a29baa08adf463241d4a7056255"></a>

## ingress_egress_gw.sm_connection_pvt_ip — ingress_egress_gw.sm_connection_pvt_ip / 55b073fd2151 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- ingress_egress_gw.sm_connection_pvt_ip

<a id="canonical-f8f8c373817ef37f923920090b76b35ac8b8c72c9e990f2ebc8213062ab160d4"></a>

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
sm_connection_pvt_ip = {}
```

<a id="canonical-48d088682f5e95fc08341db3eade7650d11ca334e659d560c8a03031b0d07268"></a>

## Direct properties — ingress_egress_gw.sm_connection_pvt_ip / 55b073fd2151 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3b53a410a211adb21280aa0fc191c690a5d1799c69ffba51e7a9db600947d5ce"></a>

## Next pages — ingress_egress_gw.sm_connection_pvt_ip / 55b073fd2151 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3154c8449a356123e727c1e31d23b173a71b972d47ed39cceb566352b559a08a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6ad9a3e21634fa34a1895f4ec7ab991941dfbe3bc13ca0bc3f9ffcf760ae5fe8"></a>

## ingress_gw — ingress_gw / 77ed93789a0c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- ingress_gw

<a id="canonical-673064a09e6a2668846219b579341e6df28ae8d7a84d2c33fa7479cf571b751e"></a>

Type: `"object"`. single nested block, Optional.

AWS Ingress Gateway. Single interface AWS ingress site.

Upstream description:

Single interface AWS ingress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_certified_hw",
    "az_nodes")}
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
ingress_gw {
  # Configure direct properties listed below.
}
```

<a id="canonical-d3191765d04f6504c5ed2cb96e953eb13ae64cfec87bf30761a790a833213b43"></a>

## Direct properties — ingress_gw / 77ed93789a0c / 3

- [allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8): complete subsection reference.

<a id="canonical-19b85bdb753522a6d4c17aa65abde9e10d52ddc913ad2c3605c16d389c05dd5b"></a>

<a id="canonical-62516fe96e2aba02a2b36fc7a8dca6ade8a0f86635e9de66597ad1befbe55e64"></a>

## aws_certified_hw property — ingress_gw / 77ed93789a0c / 4

Type: `"string"`. Optional.

\[Enum: aws-byol-voltmesh\] AWS Certified Hardware. Name for AWS certified hardware. The only
possible value is \`aws-byol-voltmesh\`.

Upstream description:

Name for AWS certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("aws-byol-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-bc6cc2703fabd08d9aa7a7cb7be83e08640f98ae12f7c2610a4c69c813f5f964): complete subsection reference.

- [performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-d755f6778136034ceb42d5cb5880580f5684cf6e38a0f35d217a6caad2f8c7a8): complete subsection reference.

<a id="canonical-fc28c54825cd4f6e5efb82257d617b673c57ee296b1cc9641ec7e4c6db7351db"></a>

## Next pages — ingress_gw / 77ed93789a0c / 5

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8)
- [ingress_gw.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-bc6cc2703fabd08d9aa7a7cb7be83e08640f98ae12f7c2610a4c69c813f5f964)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-d755f6778136034ceb42d5cb5880580f5684cf6e38a0f35d217a6caad2f8c7a8)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6432943caa325aa87fb7ce46f494c443c6a4706282a294d7c297d7c2be9d74d4"></a>

## ingress_gw.allowed_vip_port — ingress_gw.allowed_vip_port / b55a63c2343b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- ingress_gw.allowed_vip_port

<a id="canonical-c90fa5a4224d92f808c329b82eec9b69a8d6b6017ecf70129f7b1152412e8583"></a>

Type: `"object"`. single nested block, Optional.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-5f0bcc06a603677a75205395a3a432bc524472a18a355a4952e34560b98224bc"></a>

## Direct properties — ingress_gw.allowed_vip_port / b55a63c2343b / 3

- [custom_ports](resources--aws_vpc_site--reference--group-003.md#canonical-cc71034dd4f1349fa5a4ebbed99c4dea5a154cdae165cbc374da0f123e707504): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-09e6324abf690e4bf8e621820d0bea9999207d5f01133746075679e2d8f37e6b): complete subsection reference.

- [use_http_https_port](resources--aws_vpc_site--reference--group-003.md#canonical-d6d876eccc4b754c6d1039bf508dbc758a2516a880409b018b097d12a01149f8): complete subsection reference.

- [use_http_port](resources--aws_vpc_site--reference--group-003.md#canonical-f34ce9e6772a2c6a5a65862f06fbce4cd4de288fc048e55de831d120b83f2824): complete subsection reference.

- [use_https_port](resources--aws_vpc_site--reference--group-003.md#canonical-3a57a193d5cf5d26cd580fc6f232894634230c18117dcbb168227f55b1e427bd): complete subsection reference.

<a id="canonical-5930b8552059c3b0efc4d45e57abd67d1e012ad0f9d92ea0e5f7ec1024cd36fc"></a>

## Next pages — ingress_gw.allowed_vip_port / b55a63c2343b / 4

- [ingress_gw.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-003.md#canonical-cc71034dd4f1349fa5a4ebbed99c4dea5a154cdae165cbc374da0f123e707504)
- [ingress_gw.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-09e6324abf690e4bf8e621820d0bea9999207d5f01133746075679e2d8f37e6b)
- [ingress_gw.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-003.md#canonical-d6d876eccc4b754c6d1039bf508dbc758a2516a880409b018b097d12a01149f8)
- [ingress_gw.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-003.md#canonical-f34ce9e6772a2c6a5a65862f06fbce4cd4de288fc048e55de831d120b83f2824)
- [ingress_gw.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-003.md#canonical-3a57a193d5cf5d26cd580fc6f232894634230c18117dcbb168227f55b1e427bd)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-cc71034dd4f1349fa5a4ebbed99c4dea5a154cdae165cbc374da0f123e707504"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0829475a28d93e1dc50eeb3d55c41c5106d0419d2755e5e56a3f1d2d84d639f8"></a>

## ingress_gw.allowed_vip_port.custom_ports — ingress_gw.allowed_vip_port.custom_ports / b4a9d3d549ad / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8)
- ingress_gw.allowed_vip_port.custom_ports

<a id="canonical-2da8adf8c1e911e99f25f17dbdbf761c1f6b6a346830baf523b33e696db8f664"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
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
custom_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-b433fb74eaa06e9c7089dd68099a3df2d0e88213126400aa831a7979025ac7f0"></a>

## Direct properties — ingress_gw.allowed_vip_port.custom_ports / b4a9d3d549ad / 3

<a id="canonical-0ca6a1abbbe8fadda423c937c017f88a463527493fe0fde4c80b31d78105d182"></a>

<a id="canonical-71556dffd984e6f70b54a28515ddbdaeffa0731cda948e9459deaea197b2e854"></a>

## port_ranges property — ingress_gw.allowed_vip_port.custom_ports / b4a9d3d549ad / 4

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-41b19154fd4c669042bd04d9b353077150b348048a044ce643a63feb6f2c76de"></a>

## Next pages — ingress_gw.allowed_vip_port.custom_ports / b4a9d3d549ad / 5

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-09e6324abf690e4bf8e621820d0bea9999207d5f01133746075679e2d8f37e6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5baeed114410604910cb3f8d375b524568775573dde18ebdf68cea6271bef0ef"></a>

## ingress_gw.allowed_vip_port.disable_allowed_vip_port — ingress_gw.allowed_vip_port.disable_allowed_vip_port / f0eabc291091 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8)
- ingress_gw.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-0fa474b2bb8bfce99b97a7de49e681b4f6942efae1fa670447ff7428102be086"></a>

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
disable_allowed_vip_port = {}
```

<a id="canonical-355f7a32c89683fa397c8f5da4bbd93dd7874be74e491b032b90685baa780d99"></a>

## Direct properties — ingress_gw.allowed_vip_port.disable_allowed_vip_port / f0eabc291091 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f71583f1219230849569211373d9950bd1b39e6977d824c716990ae9b0c33480"></a>

## Next pages — ingress_gw.allowed_vip_port.disable_allowed_vip_port / f0eabc291091 / 4

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-d6d876eccc4b754c6d1039bf508dbc758a2516a880409b018b097d12a01149f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52320531d1cc10205d09c734699e8f116e800e4a11dcc8617944d3275abd3737"></a>

## ingress_gw.allowed_vip_port.use_http_https_port — ingress_gw.allowed_vip_port.use_http_https_port / 9e47c048089c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8)
- ingress_gw.allowed_vip_port.use_http_https_port

<a id="canonical-f77a87fbb6d87fe365edf63021f384dc34a2818ea25ba67518bff2df4950d3af"></a>

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
use_http_https_port = {}
```

<a id="canonical-ce5c23919acd1f9e9cd036370520baa8a68c583f48846e97394f0edd08ffe626"></a>

## Direct properties — ingress_gw.allowed_vip_port.use_http_https_port / 9e47c048089c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-63d5714697bb4f9bf6c1b7c86cb77fa243c4e9d2fcea3d5567ac3ab37bfe7cc3"></a>

## Next pages — ingress_gw.allowed_vip_port.use_http_https_port / 9e47c048089c / 4

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-f34ce9e6772a2c6a5a65862f06fbce4cd4de288fc048e55de831d120b83f2824"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26ce1bd780739d539d8f35811ace92a5050c4bcc575b8d3e3a9b8605de576f2a"></a>

## ingress_gw.allowed_vip_port.use_http_port — ingress_gw.allowed_vip_port.use_http_port / 5ef587038e50 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8)
- ingress_gw.allowed_vip_port.use_http_port

<a id="canonical-8bf2432399d25b296b89ea498eb8e67b31c4a291770e51f4be6cb28f8fd884e0"></a>

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
use_http_port = {}
```

<a id="canonical-7d32354c095a517c8fc56fce268ae58b753665f91d816949bea67634185a5e0e"></a>

## Direct properties — ingress_gw.allowed_vip_port.use_http_port / 5ef587038e50 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b9e20a73c2538cc6e45215d64a44d29eb336272937b330685812d502717951e7"></a>

## Next pages — ingress_gw.allowed_vip_port.use_http_port / 5ef587038e50 / 4

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-3a57a193d5cf5d26cd580fc6f232894634230c18117dcbb168227f55b1e427bd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
