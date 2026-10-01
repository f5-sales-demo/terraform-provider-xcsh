---
page_title: "xcsh_gcp_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_gcp_vpc_site reference."
---

# xcsh_gcp_vpc_site reference

<a id="canonical-432769501891b44789983574a248d7ba1347b1e313f0733907ddb52a60cac1aa"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 85c25aa5ec68 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-eab63ac4254ea7eb934be257b7d5ac26b247582615580be17033ef45cb52f4b4"></a>

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

<a id="canonical-d33e6978865aa41cf2df7dad84134ab2edcb32aaf4b3d9a5b021b26c0a7366d9"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 85c25aa5ec68 / 3

<a id="canonical-ab75c7c1a7d6608f7fb9bb100d07e6d2f8d65b4fc039ec45eca79b1597bd5d69"></a>

<a id="canonical-0802a862e6896ed44ad63b3b048e42b2a8ab0ed0f8bdeb18d6f6824e5525f272"></a>

## attrs property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 85c25aa5ec68 / 4

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

- [labels](resources--gcp_vpc_site--reference--group-003.md#canonical-9cf5cbd48bef938d2396c76a34bc5f0ae9e9384b45605ff23f9cc85b5fb4f933): complete subsection reference.

- [nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c): complete subsection reference.

- [subnets](resources--gcp_vpc_site--reference--group-003.md#canonical-5ac13beab669eb3746e95b18cc0f23bd05c5b0d0eb96d0409da7ccfb1829eb01): complete subsection reference.

<a id="canonical-2abb44f57c75c20f54514912cf7e272f1e4bd33e573f073487e15bbabfe4d91c"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 85c25aa5ec68 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](resources--gcp_vpc_site--reference--group-003.md#canonical-9cf5cbd48bef938d2396c76a34bc5f0ae9e9384b45605ff23f9cc85b5fb4f933)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-003.md#canonical-5ac13beab669eb3746e95b18cc0f23bd05c5b0d0eb96d0409da7ccfb1829eb01)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-9cf5cbd48bef938d2396c76a34bc5f0ae9e9384b45605ff23f9cc85b5fb4f933"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6197019d8c609869e3125c5479f9e6d598d6d5a9dc779382942ec2fa8775e6b4"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / c421a276e3a3 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-2abd43c8e800b5027639a6bb346cd530122e2839283cd87a221491bf7c8c9da7"></a>

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

<a id="canonical-741eb88df3ef3412885a4439ecd342792924ec6724c4a4f2988969256830c8c6"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / c421a276e3a3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-adf6c0f2809b2b447aceba799a4ce6366b0acbbfd4e3dbecb6e95bf5fc9de3d5"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / c421a276e3a3 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d637279e19e094aded43246328df8599b29a94632f29289ccfe5af12f4e240a0"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 3bf3e3bd3f31 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-779c38e03356d852b932c31b1a19ae21008fe9bc77e1f3b38ebfae7b5979dc48"></a>

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

<a id="canonical-6b20faf040f4d54bf3816e142195fca14bfaf4ec0fad9b15ccb3b165ca4967db"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 3bf3e3bd3f31 / 3

- [interface](resources--gcp_vpc_site--reference--group-003.md#canonical-5c0549a3e15061d721c5f5b57bd678c25f627b631a28c558f62e1af0f1364389): complete subsection reference.

- [nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-2ac1ad33c8e24db542afc9131dd8a1e24969956f6f6387f85e8c3fbb788eb468): complete subsection reference.

<a id="canonical-649da02f03f021229999887b010dfc174914cf560f61c66b20d8870fa0599a44"></a>

<a id="canonical-d165354884a18b158916185fb34f77ffba1a449f37ff73b1118ee5c561979b1c"></a>

## type property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 3bf3e3bd3f31 / 4

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

<a id="canonical-d2179a83c12a6103dbd1e5ba974fbefa99479ca980bf623542c2d0b6285e1765"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 3bf3e3bd3f31 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--gcp_vpc_site--reference--group-003.md#canonical-5c0549a3e15061d721c5f5b57bd678c25f627b631a28c558f62e1af0f1364389)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-2ac1ad33c8e24db542afc9131dd8a1e24969956f6f6387f85e8c3fbb788eb468)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5c0549a3e15061d721c5f5b57bd678c25f627b631a28c558f62e1af0f1364389"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d8619b4abe7989c3420e25e21c9c72ee78ae2deb80776e726a0896af7e9decec"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / cd8193bc1c8f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-322abc8ae06ec90139186b47c29fe533a3bffbe99669659d64500d3c8e6dc127"></a>

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

<a id="canonical-354281dc6e0ecf5702af3114830f7aa5c9305745eee50050ab9663a5e6696083"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / cd8193bc1c8f / 3

<a id="canonical-a6aeddd84d734ea71fbf47899f3b8c3df6a939ac6a636196f3b29b22a7995c34"></a>

<a id="canonical-8aafe10143dadc14d7950c5ca0c377f797327c69d0a69eaa1fe80956017560b9"></a>

## kind property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / cd8193bc1c8f / 4

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

<a id="canonical-b72a67218def7c45a5d200a69196fcb77b7f803999aa8ff09cdb4bd863bea6ee"></a>

<a id="canonical-44baf83cea86c609aaa94f1f2caa7bdfb3ebd7c19fc01c7db3c9d4c029191796"></a>

## name property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / cd8193bc1c8f / 5

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

<a id="canonical-204efbed3d6e01a4248b5ce93f1d2d0419bb861adc079dbccd5f5ae3bbf6bf88"></a>

<a id="canonical-b97c88b5b52102e528bbd0a1fe8fedafba09e749d5508367226f6a016f3d11cf"></a>

## namespace property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / cd8193bc1c8f / 6

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

<a id="canonical-dc4ece3dbc08c41efdb9e501bf3a38264f562507d21a16d73bd6a130216200ec"></a>

<a id="canonical-e8e7f9be6b48dc253ee719f36ec1bfb78782d037df204932a27d6a30dc7d8c89"></a>

## tenant property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / cd8193bc1c8f / 7

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

<a id="canonical-7dd6ef8dd787053e0ab8d1ce0a07330dbcf190d1558fd159e0ae07b3355e0052"></a>

<a id="canonical-38567f96b5016a33915e63b791941681e04ac95115f48c1e29dcfa1352d17df2"></a>

## uid property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / cd8193bc1c8f / 8

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

<a id="canonical-75f0ba6c558975ae95d0108f161d159ee5f4c9895e4ea8395434387d183c7ba8"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / cd8193bc1c8f / 9

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-2ac1ad33c8e24db542afc9131dd8a1e24969956f6f6387f85e8c3fbb788eb468"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-089e85d7e172bc207541fa74ffbea2c2313598563be13ef66f32ed70c662ccc1"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ed5602e3f39a / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-83642e9293fff071211cceb256c429ec50c4abd6eb3010034b2807db0395c81f"></a>

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

<a id="canonical-9e92ae59572eafa2780e57e3dfc26c99c29ae338473f210b24d2e78f442d1561"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ed5602e3f39a / 3

- [dual_stack](resources--gcp_vpc_site--reference--group-003.md#canonical-d7d94e0c818806e3f621e13d94466d04ce8a84dfb15b27cb259c62245a2e4f72): complete subsection reference.

- [ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-5e9baba149538d0f86aa2cc082c838264d5778f3d46f5c84e517f816a11bd0e3): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-1ca2a6150824bed2693b110331a96bf7105b6e5e4ac473f0222f07ff959304d8): complete subsection reference.

<a id="canonical-70da8d715d9b0defede27b8c4bbae40b0e572253733ef805a7a3eda4f95eddb0"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ed5602e3f39a / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-003.md#canonical-d7d94e0c818806e3f621e13d94466d04ce8a84dfb15b27cb259c62245a2e4f72)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-5e9baba149538d0f86aa2cc082c838264d5778f3d46f5c84e517f816a11bd0e3)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-1ca2a6150824bed2693b110331a96bf7105b6e5e4ac473f0222f07ff959304d8)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-d7d94e0c818806e3f621e13d94466d04ce8a84dfb15b27cb259c62245a2e4f72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60635da3dd68c2115bde57e9b6b6bef8b5b301dc135f3bf28850e270ec49e8ea"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 20a37bf7d203 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-2ac1ad33c8e24db542afc9131dd8a1e24969956f6f6387f85e8c3fbb788eb468)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-9ec8a134aa837255aaf4a145150149804403c7448cd05dfda8810501599aa288"></a>

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

<a id="canonical-5aa2d58077a08c43e98e1237ebfff79273f9ede1316d694d0bbe6dd01974043c"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 20a37bf7d203 / 3

- [ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-907118cb914089b56453b51c7fd0b4615450d4807c9fc3fac372e736c21184a8): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-1345feb9c34bd00a276947cf3f2f4739348c17df7c78a2916d01f25d8abe82b0): complete subsection reference.

<a id="canonical-0d3247b8c21c3e2f73d3db2c42be6af50f69cccd28be78572097679262ab1edc"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 20a37bf7d203 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-907118cb914089b56453b51c7fd0b4615450d4807c9fc3fac372e736c21184a8)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-1345feb9c34bd00a276947cf3f2f4739348c17df7c78a2916d01f25d8abe82b0)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-2ac1ad33c8e24db542afc9131dd8a1e24969956f6f6387f85e8c3fbb788eb468)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-907118cb914089b56453b51c7fd0b4615450d4807c9fc3fac372e736c21184a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-af5e9ecef7dbeea627af1d7da69b6e46e2f3b6b4c0de2de358a89c0864eb08f0"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 066b0651c36d / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-2ac1ad33c8e24db542afc9131dd8a1e24969956f6f6387f85e8c3fbb788eb468)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-003.md#canonical-d7d94e0c818806e3f621e13d94466d04ce8a84dfb15b27cb259c62245a2e4f72)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-c479744450b567fd5c91903d2f517d0621525a0e085684d99d35270d4a85dc5d"></a>

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

<a id="canonical-c06012660d0f5c62d8d42efb260471ac31d1e8dd925637b2ebe87fc2afa9cb91"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 066b0651c36d / 3

<a id="canonical-658f5c09f1e8d67ed4ebaef71a5c16fa439584479e620b6a9afd334f7d944826"></a>

<a id="canonical-90f573a0e10dc0292812d890cbe604de3d77acb9de5d5aee38ca29d12165be96"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 066b0651c36d / 4

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

<a id="canonical-7a7355f56d6e18869cfdac8983d0f9a10e670b8bfa227cf8cb3c44f236c1a50d"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 066b0651c36d / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-003.md#canonical-d7d94e0c818806e3f621e13d94466d04ce8a84dfb15b27cb259c62245a2e4f72)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-1345feb9c34bd00a276947cf3f2f4739348c17df7c78a2916d01f25d8abe82b0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb463bc5ea827a05d1743d2b1285251a97dea8afb3d39acccab89ca261346431"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c9b52b17e5d2 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-2ac1ad33c8e24db542afc9131dd8a1e24969956f6f6387f85e8c3fbb788eb468)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-003.md#canonical-d7d94e0c818806e3f621e13d94466d04ce8a84dfb15b27cb259c62245a2e4f72)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-5862a73c34b43d4eb455da8f2abe438510d3e2e101e0ff135041773a999128fa"></a>

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

<a id="canonical-e57e3b1621701ac3c653d4a3b53ff2b0a201aaea8318179323132bdf34b8b628"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c9b52b17e5d2 / 3

<a id="canonical-44d229f8691027ca87401be1cd519430a93612afbb928bc6242cb6c20c41b037"></a>

<a id="canonical-89c11ae4a578242b9594c5d316e77ca96d0471a10f7d2c285be72626fdf303f3"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c9b52b17e5d2 / 4

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

<a id="canonical-8d37f04ad0e52d7c12cad1e964bd443d04dbac1fbccab10f29b2227520e40854"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / c9b52b17e5d2 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--gcp_vpc_site--reference--group-003.md#canonical-d7d94e0c818806e3f621e13d94466d04ce8a84dfb15b27cb259c62245a2e4f72)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5e9baba149538d0f86aa2cc082c838264d5778f3d46f5c84e517f816a11bd0e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-478a87feee69db77cc971aa9bf72da9cb6eac429b1bf7b917a4279d91f662ec3"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 8565fdcd9554 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-2ac1ad33c8e24db542afc9131dd8a1e24969956f6f6387f85e8c3fbb788eb468)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-72d955f9e802b86ff2ae156ed8cb0a4394400917090e75ccbdb50f71f2eafa7f"></a>

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

<a id="canonical-3f65c4c856a70f2165ce5846f4f043b57c7d02e4513c43f759150d5bb46db0ac"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 8565fdcd9554 / 3

<a id="canonical-82b4bbf9d3958076a4358924fbacad3623599ff13259abe4204700096e6dfb3a"></a>

<a id="canonical-10ee805392253dd87d8b94a69bbafda2bf7b9ec0cae6731756cd6a7fbae9548b"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 8565fdcd9554 / 4

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

<a id="canonical-8f6abdaa73d7faf90ef7e115efee3187ce574129edde625a1a4935c2c0f68847"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 8565fdcd9554 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-2ac1ad33c8e24db542afc9131dd8a1e24969956f6f6387f85e8c3fbb788eb468)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-1ca2a6150824bed2693b110331a96bf7105b6e5e4ac473f0222f07ff959304d8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c8283c93b3eccb8b2e5ede141ac5cb8dc6772bb471f02e9f65f6d446e273e59e"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ae5ed3b5e68f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--gcp_vpc_site--reference--group-003.md#canonical-5736e59660094bd266ed92c899e642a3d60f7f2ff14f256c4410cc18c988608c)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-2ac1ad33c8e24db542afc9131dd8a1e24969956f6f6387f85e8c3fbb788eb468)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-c7770aed39817279a03b90f7c1a9596890520bf6d5c1b45b92257c3908380326"></a>

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

<a id="canonical-c58ffeb27c797584164ad370267a8e7de7d5ea91f60e0f490fc3fa997e97763e"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ae5ed3b5e68f / 3

<a id="canonical-524882a7e347d31b0124fc614f7d14de39bf0b11d504a7697e4dbe268908668a"></a>

<a id="canonical-8e4747272d408df9bdbff31ace6e3bd9906bd3db5ba58b259111bbcf24134e85"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ae5ed3b5e68f / 4

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

<a id="canonical-09ceacc494533e757fcce90f424f13840929b577fd1b789865a168d24f596d6b"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / ae5ed3b5e68f / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--gcp_vpc_site--reference--group-003.md#canonical-2ac1ad33c8e24db542afc9131dd8a1e24969956f6f6387f85e8c3fbb788eb468)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5ac13beab669eb3746e95b18cc0f23bd05c5b0d0eb96d0409da7ccfb1829eb01"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-edb130fb6550e6a93a0a68abdb111b2368be95737d3eeac231b72adfba832b45"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 8d11bbe83e6c / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-918b6371b0b9b19b3ffd0bb5dadded0b8616c58583c4668fd2e7b6a988530248"></a>

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

<a id="canonical-a888bf9ec02aad809852d096e6464b5857201c2509af4b9143dfb5a49777fab6"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 8d11bbe83e6c / 3

- [ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-6b64bea7608162576438fc51abd21dd6103718f9d15d0b44d159a919a8a3fbed): complete subsection reference.

- [ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-d81d7cc9508ebd71953d2d3cbed1175011a99c8e34c0ba9c24009abdc6ad9f38): complete subsection reference.

<a id="canonical-50080ac19cc7cd5ee969e9b8fc7c067a65ba452c290f5b28cec5e9ec4a737980"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 8d11bbe83e6c / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--gcp_vpc_site--reference--group-003.md#canonical-6b64bea7608162576438fc51abd21dd6103718f9d15d0b44d159a919a8a3fbed)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--gcp_vpc_site--reference--group-003.md#canonical-d81d7cc9508ebd71953d2d3cbed1175011a99c8e34c0ba9c24009abdc6ad9f38)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-6b64bea7608162576438fc51abd21dd6103718f9d15d0b44d159a919a8a3fbed"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e03b3beefbb716c442b0774626396d2aa0d02f699bbd9b2062d751b5a92fef40"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 317f1955a54d / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-003.md#canonical-5ac13beab669eb3746e95b18cc0f23bd05c5b0d0eb96d0409da7ccfb1829eb01)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-93fce2ebee00c7160a7196cacacf97c1e9e42b8edd5146ba93e4db393d84f13b"></a>

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

<a id="canonical-3bd8065c40d0cd1204aaaf42d1d4332382790ef60e45d41ee89504fdfbfd13d2"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 317f1955a54d / 3

<a id="canonical-4c9a6e415e686ced8015543f571b459eb7a880a1b4acc2c1344ad80d7e55be16"></a>

<a id="canonical-0f8dce0f87de46014233bf1700aadc26d126b779559110a7bf914af146d1fb62"></a>

## plen property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 317f1955a54d / 4

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

<a id="canonical-9fb5831e0f3fd694089576e93d1251f900518a71932cce1a9137a68fe67052d5"></a>

<a id="canonical-4a8b81190b325f6716bc242671c1e2ff08d957dbd640fb68c41e2fdcdaac18ea"></a>

## prefix property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 317f1955a54d / 5

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

<a id="canonical-9cc3db2b77d723b12dd7dfbd43f141a1ec2abfc267d67a2f25f72a1c5ea0e6dd"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 317f1955a54d / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-003.md#canonical-5ac13beab669eb3746e95b18cc0f23bd05c5b0d0eb96d0409da7ccfb1829eb01)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-d81d7cc9508ebd71953d2d3cbed1175011a99c8e34c0ba9c24009abdc6ad9f38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c835cacc3b3d2db7b80583adbac1b33a654870a1ff879561e11f454b3c9a880"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 0b3bf194edc1 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_static_routes](resources--gcp_vpc_site--reference--group-002.md#canonical-e3e887e7fb906b41f7e478a0d9854e1f953264dcd3ce9439fd6b32c010f56a78)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--gcp_vpc_site--reference--group-002.md#canonical-3240e7d173b697ec7cd9795af58e465d67139934e84ea6a5e4d56bf35cd44c43)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--gcp_vpc_site--reference--group-002.md#canonical-27cf4bd04cb4f6180d55d8f1d259009714f95a5397ba1dcc1354c6afdae0547d)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-003.md#canonical-5ac13beab669eb3746e95b18cc0f23bd05c5b0d0eb96d0409da7ccfb1829eb01)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-25d295d0eed640b6faa4d4cbbea4cc2fecef6dccded0bd34f3f62ca35135a2fc"></a>

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

<a id="canonical-d81f17afd119a966e57b46ffb6409abe42edd5958b324e66d8dcc6fb72d405de"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 0b3bf194edc1 / 3

<a id="canonical-44f36f8939e4408cd178af5fc72369115574a1913072b23040b241aa8c18e920"></a>

<a id="canonical-197f25935507e3db76bbe392f6e6515654472826d835eca1d86c66c03ec75aa8"></a>

## plen property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 0b3bf194edc1 / 4

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

<a id="canonical-264cad398fdc63af1c4e0faeb749eca461e885df71b779f548aebbd05856698a"></a>

<a id="canonical-1f1d26b0694e13e3976592d91c871e4e594e7ed62aadafd5fceed8dc58b7a1c9"></a>

## prefix property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 0b3bf194edc1 / 5

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

<a id="canonical-2f0f42c836edb7c081df59bef89632edc63608d0a436a6fc18d08d742fb76d24"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 0b3bf194edc1 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--gcp_vpc_site--reference--group-003.md#canonical-5ac13beab669eb3746e95b18cc0f23bd05c5b0d0eb96d0409da7ccfb1829eb01)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-fd6015d0e38216aa4cdf4640bf7e23ca3cdecce3b1cfa6b36fba590869132af1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-76ce52c9ad2520163d12a81d65a0f44dcd9357939aab27ce02f3974c31d5d37a"></a>

## ingress_egress_gw.outside_subnet — ingress_egress_gw.outside_subnet / f9b5456ab476 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.outside_subnet

<a id="canonical-9a578c42bc083926bf10f0050692291d5c62860fec9846fbf277927c21b19e59"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet",
    "new_subnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

Terraform syntax:

```terraform
outside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-01eb3ee9fa46896c37a8c2fc01d056b3b02ae04547455f2cf6b0c330222de7d4"></a>

## Direct properties — ingress_egress_gw.outside_subnet / f9b5456ab476 / 3

- [existing_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-52b9db717e18997ccb739e19e05830dd2a5c9be524c6a95b654e4b0b114d86eb): complete subsection reference.

- [new_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-eddbcfed8d3be3a35b12d8a1e7d078ff2a718a476832a1909f96b7de46a2d3f8): complete subsection reference.

<a id="canonical-9929aae3971da30e68907b22bd7e0fdeb88307edfb297b5a9b051aab9eb32ea8"></a>

## Next pages — ingress_egress_gw.outside_subnet / f9b5456ab476 / 4

- [ingress_egress_gw.outside_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-52b9db717e18997ccb739e19e05830dd2a5c9be524c6a95b654e4b0b114d86eb)
- [ingress_egress_gw.outside_subnet.new_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-eddbcfed8d3be3a35b12d8a1e7d078ff2a718a476832a1909f96b7de46a2d3f8)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-52b9db717e18997ccb739e19e05830dd2a5c9be524c6a95b654e4b0b114d86eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-676b900aac85a16f90ff0ffdd16eac1e293c28ea9b851c9bf7ddc6119be9a8e9"></a>

## ingress_egress_gw.outside_subnet.existing_subnet — ingress_egress_gw.outside_subnet.existing_subnet / d43b755db784 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-fd6015d0e38216aa4cdf4640bf7e23ca3cdecce3b1cfa6b36fba590869132af1)
- ingress_egress_gw.outside_subnet.existing_subnet

<a id="canonical-59bdf79f3cfc8595b169625a3b943b553db84927a66c9f83f75b0e2fc824ac1c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name")}
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
existing_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-2fd05033d33bc5a7ed7706321c9a8c89668539dc08ffa4f69155a2a676924bda"></a>

## Direct properties — ingress_egress_gw.outside_subnet.existing_subnet / d43b755db784 / 3

<a id="canonical-2200db1a4075d94eaf41befbe970b4d417d81c333b26a2538c9c53d1e789ae04"></a>

<a id="canonical-43b476aacf5a99c8b00acea565c9f58a2086db3c933a134d737d41f3c263d4b7"></a>

## subnet_name property — ingress_egress_gw.outside_subnet.existing_subnet / d43b755db784 / 4

Type: `"string"`. Optional.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-f383ea7824d84557d35aa8667f55556557f1d99b58d5eff225d678c775da63f7"></a>

## Next pages — ingress_egress_gw.outside_subnet.existing_subnet / d43b755db784 / 5

- [ingress_egress_gw.outside_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-fd6015d0e38216aa4cdf4640bf7e23ca3cdecce3b1cfa6b36fba590869132af1)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-eddbcfed8d3be3a35b12d8a1e7d078ff2a718a476832a1909f96b7de46a2d3f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53b8495621f213145b26dcd097b960f7b9bab36694dce39cf872199ad87148ec"></a>

## ingress_egress_gw.outside_subnet.new_subnet — ingress_egress_gw.outside_subnet.new_subnet / 81b676415685 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.outside_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-fd6015d0e38216aa4cdf4640bf7e23ca3cdecce3b1cfa6b36fba590869132af1)
- ingress_egress_gw.outside_subnet.new_subnet

<a id="canonical-6da8d9343fa11eb8f4cd94d58e72097d3fdffe322d5352d71a8b50a97d2b5f3b"></a>

Type: `"object"`. single nested block, Optional.

GCP subnet parameters Type. Parameters for GCP subnet.

Upstream description:

Parameters for GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4")}
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
new_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-2cd71d28e910b1358feddc21626e67f06f1b554c8c8f197a770e1e8e7692e327"></a>

## Direct properties — ingress_egress_gw.outside_subnet.new_subnet / 81b676415685 / 3

<a id="canonical-50d00fb1f42454eacd39bbc2e70e6138f2a5eb9b96ebdf36bf2f216daa6f5251"></a>

<a id="canonical-cea94dc3eacd33df10ced0153c23d9b083462cc46ffb72431fd9076af0d5e853"></a>

## primary_ipv4 property — ingress_egress_gw.outside_subnet.new_subnet / 81b676415685 / 4

Type: `"string"`. Optional.

IPv4 prefix for this Subnet. It has to be private address space.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  }
}
```

<a id="canonical-dc2b5bfeec8505491435c68a14f4b739a137b3e91c5b857801c0165440349e94"></a>

<a id="canonical-a11087b96b6ed795e91c4e994204a5d362d32184e5a147655cfd0f4fcec4cf24"></a>

## subnet_name property — ingress_egress_gw.outside_subnet.new_subnet / 81b676415685 / 5

Type: `"string"`. Optional.

Name of new VPC Subnet, will be autogenerated if empty.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-88e2513bea66fa8c97fb7e2a2648660c091295e7ab2628567b198b7f0be1fe00"></a>

## Next pages — ingress_egress_gw.outside_subnet.new_subnet / 81b676415685 / 6

- [ingress_egress_gw.outside_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-fd6015d0e38216aa4cdf4640bf7e23ca3cdecce3b1cfa6b36fba590869132af1)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-7bf6c0392d59602a6538d1bc0324b33b8ee64549bf15b6769445b37467b8f67d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ba71f2b401598aae01b946ac99a34f6f8a85d9558cec73a548928e74877acba"></a>

## ingress_egress_gw.performance_enhancement_mode — ingress_egress_gw.performance_enhancement_mode / 8ea8285dfc70 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.performance_enhancement_mode

<a id="canonical-49631701caf3f2a0c8e89a3939443712e3409f13a73a281c090000807839aa4a"></a>

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

<a id="canonical-a4d2b755d702a4b19cba07a685e9485460e94f780b74e590e42f5c29d28496f1"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode / 8ea8285dfc70 / 3

- [perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-8e18efcbc8b3ffbc3209149961b20437c04dbec8a74860ee4ae2cad26237e531): complete subsection reference.

- [perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-fefcfbd02a71d9499e2a38baeb806a81eba607f5ad185939d8585dc0e6c147b1): complete subsection reference.

<a id="canonical-7dcf8c3304ba063c0fc83cbfdc72ce8c005abe177dbd78a9e65bfa5cf92d2e8f"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode / 8ea8285dfc70 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-8e18efcbc8b3ffbc3209149961b20437c04dbec8a74860ee4ae2cad26237e531)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-fefcfbd02a71d9499e2a38baeb806a81eba607f5ad185939d8585dc0e6c147b1)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-8e18efcbc8b3ffbc3209149961b20437c04dbec8a74860ee4ae2cad26237e531"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27277f0dac3b042f9a4709bbab10a79e0d17701680074103fc50c6840e8128da"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / b6090a45bbb1 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-7bf6c0392d59602a6538d1bc0324b33b8ee64549bf15b6769445b37467b8f67d)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-631d4d4fd8afd586b15f074d4aff7ba0342f8377c142c398382ffcbdcb8b4499"></a>

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

<a id="canonical-42297d9f344d1aff10a54bd5e7b7902032d3075eb2b9f7238b6ffd65543bedee"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / b6090a45bbb1 / 3

- [jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-fa4215b3083468b589f4d6eee69ad322e254499e2a50c54e540be12be00721b1): complete subsection reference.

- [no_jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-2880e042870f935581b15a83ae5b0fc3e69be663bed68d6524833d8cf3d0c8d9): complete subsection reference.

<a id="canonical-f452261d6d378f2e21d4ec921d61acf1eff0f74d25dfdeaa1b5f68c0aa834601"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / b6090a45bbb1 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-fa4215b3083468b589f4d6eee69ad322e254499e2a50c54e540be12be00721b1)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-2880e042870f935581b15a83ae5b0fc3e69be663bed68d6524833d8cf3d0c8d9)
- [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-7bf6c0392d59602a6538d1bc0324b33b8ee64549bf15b6769445b37467b8f67d)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-fa4215b3083468b589f4d6eee69ad322e254499e2a50c54e540be12be00721b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-476304bdaa50ce34c8ab478bbeb81257fad27005070fe14935668dd3a14d9097"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 1c082b79c2a9 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-7bf6c0392d59602a6538d1bc0324b33b8ee64549bf15b6769445b37467b8f67d)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-8e18efcbc8b3ffbc3209149961b20437c04dbec8a74860ee4ae2cad26237e531)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-bc1dcd4e3e5f9762c85874eec773dadfd1b3431fae1cfd29147d34f0dbc0706d"></a>

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

<a id="canonical-5140016e7ccf006b2b4abe96e78b19ae94b10d872ae1ce1effa4cb73e0c00137"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 1c082b79c2a9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210b58308c06e7c45aee376a7ff4fb3a191e31551a3c6cc6f3bf17aeb31bd5f"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 1c082b79c2a9 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-8e18efcbc8b3ffbc3209149961b20437c04dbec8a74860ee4ae2cad26237e531)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-2880e042870f935581b15a83ae5b0fc3e69be663bed68d6524833d8cf3d0c8d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0564681b1b83f0212559bf55a0a38531f6077df3b7b5332cf70bbfab502012b"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / e5e084d475cf / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-7bf6c0392d59602a6538d1bc0324b33b8ee64549bf15b6769445b37467b8f67d)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-8e18efcbc8b3ffbc3209149961b20437c04dbec8a74860ee4ae2cad26237e531)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-b17d5838a74705ff59944c2619fade8824e9fe64811f6e18d73dcb4a37e70804"></a>

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

<a id="canonical-195980e6d1d71504aabda2f9fc743a926d531dc00a515a0406eae5bba019da54"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / e5e084d475cf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-581262b793f7b23753533ee6473537b68e466d24a2a272bc6e2834b428081c8e"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / e5e084d475cf / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-8e18efcbc8b3ffbc3209149961b20437c04dbec8a74860ee4ae2cad26237e531)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-fefcfbd02a71d9499e2a38baeb806a81eba607f5ad185939d8585dc0e6c147b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-71e8488570c491c984cf9a3397d1b0a27ec99732bc7b17372520b443a3a14507"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 369f7ee91a66 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-7bf6c0392d59602a6538d1bc0324b33b8ee64549bf15b6769445b37467b8f67d)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-c07a44cd28d6892dec19751848e1bf6d1064f8265c42b850679614a621f70f6b"></a>

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

<a id="canonical-3a96c203ecbcd0fbdaca54664faa09510a1db1b5e764fcf77200679c447ccb57"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 369f7ee91a66 / 3

- [jumbo_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-3c5a58ad47597dd179c2103d51006cb3d32452b47bd7c7ce886f566ca39e0881): complete subsection reference.

- [jumbo_enabled](resources--gcp_vpc_site--reference--group-003.md#canonical-55f2a253bf97437b5c37deba3a080b97d47a10faacbd43850ea748c434eca749): complete subsection reference.

<a id="canonical-5e77f9b3cef45f1a5eca40e97a9791c20d7c9b1660b0cb709c9d2502c584016f"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 369f7ee91a66 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-3c5a58ad47597dd179c2103d51006cb3d32452b47bd7c7ce886f566ca39e0881)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--gcp_vpc_site--reference--group-003.md#canonical-55f2a253bf97437b5c37deba3a080b97d47a10faacbd43850ea748c434eca749)
- [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-7bf6c0392d59602a6538d1bc0324b33b8ee64549bf15b6769445b37467b8f67d)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-3c5a58ad47597dd179c2103d51006cb3d32452b47bd7c7ce886f566ca39e0881"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f06aef31c912e1c47ebdb98f0cb4ad7a7253e52ebb1c0b534421b484aefb4b5"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 2f004421e34e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-7bf6c0392d59602a6538d1bc0324b33b8ee64549bf15b6769445b37467b8f67d)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-fefcfbd02a71d9499e2a38baeb806a81eba607f5ad185939d8585dc0e6c147b1)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-701601133f3569641999edcc0b2b71bfa05ce66f45e50538fb5134ecfdaac584"></a>

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

<a id="canonical-209307ecc7f2405b0809071ae04089972de3f79a06d2cddae74b52fd464332d0"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 2f004421e34e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7e75be6ec1bc6118909897ead2c70dd561b2981e66e2eedd7f72ad377b7f2662"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 2f004421e34e / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-fefcfbd02a71d9499e2a38baeb806a81eba607f5ad185939d8585dc0e6c147b1)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-55f2a253bf97437b5c37deba3a080b97d47a10faacbd43850ea748c434eca749"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a5598253b750c7d86656f4d2465631b3156a428c4082d94a6761453bbbe8409"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / b2ab7c4a08de / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [ingress_egress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-7bf6c0392d59602a6538d1bc0324b33b8ee64549bf15b6769445b37467b8f67d)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-fefcfbd02a71d9499e2a38baeb806a81eba607f5ad185939d8585dc0e6c147b1)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-29c6b8e11ee5b7f724e32b04e51c8ad49ad6a1bdfadb89db1341494e157ef81d"></a>

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

<a id="canonical-a4652be2a9da8ce6a9cdc134fac306acdaafa8b2857402efd218922aef7b34dd"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / b2ab7c4a08de / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-167b1ce068b56b26a1db5deb79cc2e22584fb6c4254ecd72cb86126131932900"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / b2ab7c4a08de / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-fefcfbd02a71d9499e2a38baeb806a81eba607f5ad185939d8585dc0e6c147b1)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-dd54f6695362eb9dc9e39e78abd5a0cf5108081701cee3694a848ce757bb44ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8105a7f1630053a77a9ba5425c08a8bce6cae3cdf812f50ca6698172d7660207"></a>

## ingress_egress_gw.sm_connection_public_ip — ingress_egress_gw.sm_connection_public_ip / 450af282ef0f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.sm_connection_public_ip

<a id="canonical-eca9ec7d3159b6ffbd992edde8871a72d19854b02fbfa7d36032163a616f2431"></a>

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

<a id="canonical-79cf338552796ee725a5ccdb9dec8b8afb2fc06b7acc4cc742bdbdcd121b1f1c"></a>

## Direct properties — ingress_egress_gw.sm_connection_public_ip / 450af282ef0f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-51a4009da25112768bc61461e929d021b9e1bdd08f138794819f0e596e2bc8ff"></a>

## Next pages — ingress_egress_gw.sm_connection_public_ip / 450af282ef0f / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-a23466e2ca253c8f2e844bb917b8d4e2850077c11c5e21634704806eed281cf5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-447c4c1da5e50f7e75c9fa9f82025da4a5501db071b4ae07de5727c9c87dd85c"></a>

## ingress_egress_gw.sm_connection_pvt_ip — ingress_egress_gw.sm_connection_pvt_ip / 34ab56404e6d / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- ingress_egress_gw.sm_connection_pvt_ip

<a id="canonical-64b31924136a5c0b320dfabd7580984c726f0f425c4fb0a115dce902f00a889f"></a>

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

<a id="canonical-07a5f54960910b637b46d4b37f8125c84c9e9fd82da0033e3e6302c8dcbc3bdd"></a>

## Direct properties — ingress_egress_gw.sm_connection_pvt_ip / 34ab56404e6d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-55946d5628000af33a01f425010d4bd8a2d2f7e3aadc1fdea575898e104f0ebf"></a>

## Next pages — ingress_egress_gw.sm_connection_pvt_ip / 34ab56404e6d / 4

- [ingress_egress_gw](resources--gcp_vpc_site--reference--group-001.md#canonical-db41c64ab7587baa2726cc98ecc2c9cdd1eb05bd051e83ba75ab6864a7232110)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7881498e74a05b48c328d756e7a33ccf85d145d9ae0885fb9f8274b3f1ed0afe"></a>

## ingress_gw — ingress_gw / 597fdc0a8259 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- ingress_gw

<a id="canonical-999174b385cdf8f7e4313fe3c8dc8ca1154d6456d47917b6df7b8c1ff499375a"></a>

Type: `"object"`. single nested block, Optional.

GCP Ingress Gateway. Single interface GCP ingress site.

Upstream description:

Single interface GCP ingress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("gcp_certified_hw",
    "gcp_zone_names")}
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

<a id="canonical-a34e80356e45b845af81f06faeaa289a5fcc784ee008d6a7cd393990cc2a56c2"></a>

## Direct properties — ingress_gw / 597fdc0a8259 / 3

<a id="canonical-ea631fe01a07f252748b3866ee9bbe405730eec17b005b0a240ba1003acdf922"></a>

<a id="canonical-e89f088ed4d23b6c4927ff6bdc7433eda902922731003f91061d8575ef3d8272"></a>

## gcp_certified_hw property — ingress_gw / 597fdc0a8259 / 4

Type: `"string"`. Optional.

\[Enum: gcp-byol-voltmesh\] GCP Certified Hardware. Name for GCP certified hardware. The only
possible value is \`gcp-byol-voltmesh\`.

Upstream description:

Name for GCP certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("gcp-byol-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "gcp-byol-voltmesh"
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
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"gcp-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-6d431fc7b5a8a19add0de709ae2280ef78e7875f8beb409f24b35659537e8b9e"></a>

<a id="canonical-3f30f923d07ae5c5d50570a39043e18c5abc6f9814e96ce64c62562e512331e2"></a>

## gcp_zone_names property — ingress_gw / 597fdc0a8259 / 5

Type: `["list", "string"]`. Optional.

X-required List of zones when instances will be created, needs to match with region selected.

Upstream description:

X-required List of zones when instances will be created, needs to match with region selected.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(3),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [local_network](resources--gcp_vpc_site--reference--group-003.md#canonical-1bdf4c2bd27b126fce78d63f221e327afdc813e2b0f2a076f635566a0806902a): complete subsection reference.

- [local_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-325362b4832bfecf2eb7575683ed7ee25e82653d6d4a74cdfb4eb96c6486b0d3): complete subsection reference.

<a id="canonical-abf2353de80c961bf3868e19b484fae2acdc5f99f22c55876ba2e42564e6772f"></a>

<a id="canonical-f087257281f4b50de112c092d51bf1fb5bc918c166be4a20fea95b82fdd97be7"></a>

## node_number property — ingress_gw / 597fdc0a8259 / 6

Type: `"number"`. Optional.

Number of main nodes to create, either 1 or 3.

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
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-6b40ff4816aae52103a699e77462c8e4dc38a9a5d5eec96994ad345d1b74f6a0): complete subsection reference.

<a id="canonical-fe1dd4459b208ef405582d89626aeb1dcd57d1a6c4dab9a38be6e8f504e5133f"></a>

## Next pages — ingress_gw / 597fdc0a8259 / 7

- [ingress_gw.local_network](resources--gcp_vpc_site--reference--group-003.md#canonical-1bdf4c2bd27b126fce78d63f221e327afdc813e2b0f2a076f635566a0806902a)
- [ingress_gw.local_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-325362b4832bfecf2eb7575683ed7ee25e82653d6d4a74cdfb4eb96c6486b0d3)
- [ingress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-6b40ff4816aae52103a699e77462c8e4dc38a9a5d5eec96994ad345d1b74f6a0)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-1bdf4c2bd27b126fce78d63f221e327afdc813e2b0f2a076f635566a0806902a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8628a7f08fb6b50b2c824ac8a558611bbfc2ed4d48fec43d4aff314386cb076e"></a>

## ingress_gw.local_network — ingress_gw.local_network / dd55ff251b2b / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- ingress_gw.local_network

<a id="canonical-565824f57008a1595cbf974e29bb99faba4c94308f7062b13fa4f9580267ca63"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_network",
    "new_network"),
  validators.ConflictingObjectAttributes("existing_network",
    "new_network_autogenerate"),
  validators.ConflictingObjectAttributes("new_network",
    "new_network_autogenerate")}
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
  "x-ves-oneof-field-choice": "[\"existing_network\",\"new_network\",\"new_network_autogenerate\"]"
}
```

Terraform syntax:

```terraform
local_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-2704ea2f3f90617abf7e6470528d7b7e0f7a7400671d446686ab4fab1fc843e1"></a>

## Direct properties — ingress_gw.local_network / dd55ff251b2b / 3

- [existing_network](resources--gcp_vpc_site--reference--group-003.md#canonical-03c1dd3bd28e52b282f1e724e0813b8a3374448a997d6a8d5cbe38c88aae0a61): complete subsection reference.

- [new_network](resources--gcp_vpc_site--reference--group-003.md#canonical-5a10a2fa298512eee01064001d72a70aa0d8dd1b9bd659c662361413fa594373): complete subsection reference.

- [new_network_autogenerate](resources--gcp_vpc_site--reference--group-003.md#canonical-3fbdc14d61a968a123db167fa87c767199bbe0fad734c37d2bc72e4c776af279): complete subsection reference.

<a id="canonical-8c40b98408171948993f44f4b4a1a1b9c4eea66f73b89dd9b99116cdb57c6ad4"></a>

## Next pages — ingress_gw.local_network / dd55ff251b2b / 4

- [ingress_gw.local_network.existing_network](resources--gcp_vpc_site--reference--group-003.md#canonical-03c1dd3bd28e52b282f1e724e0813b8a3374448a997d6a8d5cbe38c88aae0a61)
- [ingress_gw.local_network.new_network](resources--gcp_vpc_site--reference--group-003.md#canonical-5a10a2fa298512eee01064001d72a70aa0d8dd1b9bd659c662361413fa594373)
- [ingress_gw.local_network.new_network_autogenerate](resources--gcp_vpc_site--reference--group-003.md#canonical-3fbdc14d61a968a123db167fa87c767199bbe0fad734c37d2bc72e4c776af279)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-03c1dd3bd28e52b282f1e724e0813b8a3374448a997d6a8d5cbe38c88aae0a61"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ea1482bf7afcd6d624c66c279efb8e268ee9372139adfe28317f6e2835b2dda"></a>

## ingress_gw.local_network.existing_network — ingress_gw.local_network.existing_network / 7a4041f3098f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [ingress_gw.local_network](resources--gcp_vpc_site--reference--group-003.md#canonical-1bdf4c2bd27b126fce78d63f221e327afdc813e2b0f2a076f635566a0806902a)
- ingress_gw.local_network.existing_network

<a id="canonical-56a212bfa3e4738c17a733f08b29b275cd216829ec5f3a578be7ce340cf879eb"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing network.

Upstream description:

Name of existing VPC network.

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
  },
  "x-ves-oneof-field-routing_type": "[]"
}
```

Terraform syntax:

```terraform
existing_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-c297279380e765881b64212748a097dc3bad5d31a3b8f3ecebfbbc04c2bc258c"></a>

## Direct properties — ingress_gw.local_network.existing_network / 7a4041f3098f / 3

<a id="canonical-ee22643c3de58de0d3b9d4c108876329c9fe161e3fad85dd084bd2ff5cd3afc4"></a>

<a id="canonical-37d4956d4e9449c42c4f0a6798b30de665e8de33a03bcf580c86461232b37947"></a>

## name property — ingress_gw.local_network.existing_network / 7a4041f3098f / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-02ad7841eb002609abac2110d84da81aad73101b7679ac1dc8574b415ddfbd0c"></a>

## Next pages — ingress_gw.local_network.existing_network / 7a4041f3098f / 5

- [ingress_gw.local_network](resources--gcp_vpc_site--reference--group-003.md#canonical-1bdf4c2bd27b126fce78d63f221e327afdc813e2b0f2a076f635566a0806902a)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-5a10a2fa298512eee01064001d72a70aa0d8dd1b9bd659c662361413fa594373"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-274a86a9edeb70d3d2f75302ead09c12267bde2484be330d944ac1924c26770f"></a>

## ingress_gw.local_network.new_network — ingress_gw.local_network.new_network / 4760de98ce77 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [ingress_gw.local_network](resources--gcp_vpc_site--reference--group-003.md#canonical-1bdf4c2bd27b126fce78d63f221e327afdc813e2b0f2a076f635566a0806902a)
- ingress_gw.local_network.new_network

<a id="canonical-8ac46e259b17cc2c4f02bc880a9b57915c5736e8b84f77c923f7e5b4f89ed424"></a>

Type: `"object"`. single nested block, Optional.

Parameters to create a new GCP VPC Network.

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
new_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-5c6e18f4f8620b08ac096ab85e555e66ec7a83434f30212687436aae7c404bb7"></a>

## Direct properties — ingress_gw.local_network.new_network / 4760de98ce77 / 3

<a id="canonical-73dd238691969eb15736ea18e18acd3e9c778d2763b386b24f8c01cc2e5ad9b7"></a>

<a id="canonical-46f054b39c053c49335cd5f1d6b4e282b1e0e707a3d11175b6b19e949f999f6d"></a>

## name property — ingress_gw.local_network.new_network / 4760de98ce77 / 4

Type: `"string"`. Optional.

GCP VPC Network Name. Name for your GCP VPC Network.

Upstream description:

Name for your GCP VPC Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3d3ae5bb957e7dbf7a7059d06979a9001bfcfd0745c1a02abf70a39b72f9cad0"></a>

## Next pages — ingress_gw.local_network.new_network / 4760de98ce77 / 5

- [ingress_gw.local_network](resources--gcp_vpc_site--reference--group-003.md#canonical-1bdf4c2bd27b126fce78d63f221e327afdc813e2b0f2a076f635566a0806902a)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-3fbdc14d61a968a123db167fa87c767199bbe0fad734c37d2bc72e4c776af279"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ab558b381425f26a88f01b47dc824da49577e5336d4459e93cdb95f18cfe462"></a>

## ingress_gw.local_network.new_network_autogenerate — ingress_gw.local_network.new_network_autogenerate / 6fbfa3fc3b62 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [ingress_gw.local_network](resources--gcp_vpc_site--reference--group-003.md#canonical-1bdf4c2bd27b126fce78d63f221e327afdc813e2b0f2a076f635566a0806902a)
- ingress_gw.local_network.new_network_autogenerate

<a id="canonical-dc9190c876743380015c38c9e72c06d7c8982193cd1d7c31acf29adf368a2b5f"></a>

Type: `["object", {}]`. Optional.

Create a new GCP VPC Network with autogenerated name.

Receipt-pinned upstream constraints:

```json
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
new_network_autogenerate = {}
```

<a id="canonical-58e23be48bd4391b31aeeaeecc9111af4d82b934cf4cc5d3db158b410e9bd5a3"></a>

## Direct properties — ingress_gw.local_network.new_network_autogenerate / 6fbfa3fc3b62 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d0b228b6959ad1ca4dff15a8e9cb9908455925041a0a7cc629e7528db23b362a"></a>

## Next pages — ingress_gw.local_network.new_network_autogenerate / 6fbfa3fc3b62 / 4

- [ingress_gw.local_network](resources--gcp_vpc_site--reference--group-003.md#canonical-1bdf4c2bd27b126fce78d63f221e327afdc813e2b0f2a076f635566a0806902a)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-325362b4832bfecf2eb7575683ed7ee25e82653d6d4a74cdfb4eb96c6486b0d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58b417e148da33c02f8083baf217f8db359158e1371f0c135a246fd4026021aa"></a>

## ingress_gw.local_subnet — ingress_gw.local_subnet / fc5e572f2315 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- ingress_gw.local_subnet

<a id="canonical-f58fc041d1f2e4976f7ec01d550eca7fb957bdb25c2e9c9cf2aca087106ec6bd"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about GCP VPC network for a view.

Upstream description:

This defines choice about GCP VPC network for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet",
    "new_subnet")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet\",\"new_subnet\"]"
}
```

Terraform syntax:

```terraform
local_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-4e4f21c09fbf06bd56ed1a2fb337a64a04a2de1dfb52185a8fecdcc32d2885ee"></a>

## Direct properties — ingress_gw.local_subnet / fc5e572f2315 / 3

- [existing_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-4f18897f5fc5e9b66a33a9d0060f02631e72f3bffe8a336e5ebcbd28f8887d95): complete subsection reference.

- [new_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-7829d09e58cbf2c612f2fe81c1b3b8727dcc0aecb247bdb9f0e60c10b77cfaa6): complete subsection reference.

<a id="canonical-5b26c530d8d9769685b1a993a952c10c07f84e73a04b36a88c54f0ecd87f8b16"></a>

## Next pages — ingress_gw.local_subnet / fc5e572f2315 / 4

- [ingress_gw.local_subnet.existing_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-4f18897f5fc5e9b66a33a9d0060f02631e72f3bffe8a336e5ebcbd28f8887d95)
- [ingress_gw.local_subnet.new_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-7829d09e58cbf2c612f2fe81c1b3b8727dcc0aecb247bdb9f0e60c10b77cfaa6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-4f18897f5fc5e9b66a33a9d0060f02631e72f3bffe8a336e5ebcbd28f8887d95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-81bac9de709a54a01b7c3cce3e95d02033af224ec6464fa6d87801068316bd11"></a>

## ingress_gw.local_subnet.existing_subnet — ingress_gw.local_subnet.existing_subnet / 62a2dd7d3e6c / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [ingress_gw.local_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-325362b4832bfecf2eb7575683ed7ee25e82653d6d4a74cdfb4eb96c6486b0d3)
- ingress_gw.local_subnet.existing_subnet

<a id="canonical-4e0ccf60e1809038f4a0e2811e88dcd6ecb3a94d1e8d9ed517d633120fbb479f"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name")}
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
existing_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-ee84d20852b9f0b67a3ea4c8ba2055688a8daf2148111a2389c61306d08f40eb"></a>

## Direct properties — ingress_gw.local_subnet.existing_subnet / 62a2dd7d3e6c / 3

<a id="canonical-124e27b442c38d03c220c4f4c4b873ca20419c31d841c48dee8a7df21c5062b0"></a>

<a id="canonical-da4f7cbf5999b2000bd3b130c142d8edf60515ae05fdf628a5b8ed2c2dad8d43"></a>

## subnet_name property — ingress_gw.local_subnet.existing_subnet / 62a2dd7d3e6c / 4

Type: `"string"`. Optional.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-c005074a72b4d27231d38151cf30787289b40d93567b71222cc5d2e1b9e966e3"></a>

## Next pages — ingress_gw.local_subnet.existing_subnet / 62a2dd7d3e6c / 5

- [ingress_gw.local_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-325362b4832bfecf2eb7575683ed7ee25e82653d6d4a74cdfb4eb96c6486b0d3)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-7829d09e58cbf2c612f2fe81c1b3b8727dcc0aecb247bdb9f0e60c10b77cfaa6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f841f6d332daec342bff67bc2ef6450e764a35df02fc03e8aa58e7d4dc3c7926"></a>

## ingress_gw.local_subnet.new_subnet — ingress_gw.local_subnet.new_subnet / 10a434f0859f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [ingress_gw.local_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-325362b4832bfecf2eb7575683ed7ee25e82653d6d4a74cdfb4eb96c6486b0d3)
- ingress_gw.local_subnet.new_subnet

<a id="canonical-837e5a47a9a4677c1f0334683acc38f5f649d5fceec231743dc881ce6611d27b"></a>

Type: `"object"`. single nested block, Optional.

GCP subnet parameters Type. Parameters for GCP subnet.

Upstream description:

Parameters for GCP subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4")}
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
new_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-abd083260e1077f4fa98bf36c21dd9f3c8a316631813d79f4d5ea49926c9c5ce"></a>

## Direct properties — ingress_gw.local_subnet.new_subnet / 10a434f0859f / 3

<a id="canonical-ca61ee10fae123ba71c93bbd02d3fb9b82bea456e5ca33e80c30af6104d63bc2"></a>

<a id="canonical-1b59dbc4aaeccdc7bbfefdadb0c4641099e6f2fb3477a42e2083fb5e56fa35c4"></a>

## primary_ipv4 property — ingress_gw.local_subnet.new_subnet / 10a434f0859f / 4

Type: `"string"`. Optional.

IPv4 prefix for this Subnet. It has to be private address space.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "8"
  }
}
```

<a id="canonical-2939541af982605d8207f82f19a48e7a40a61b105ca9dc1f9b5e855967209fd8"></a>

<a id="canonical-c9c918046cc9172d7f77b795b5c8558a3340d7b9d1dc94ba07d1a3b91f9e38e7"></a>

## subnet_name property — ingress_gw.local_subnet.new_subnet / 10a434f0859f / 5

Type: `"string"`. Optional.

Name of new VPC Subnet, will be autogenerated if empty.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-6f95e9c3c7a46f6d550bd629668042e81a84bdca413ae12ae9879ff8f1c5a616"></a>

## Next pages — ingress_gw.local_subnet.new_subnet / 10a434f0859f / 6

- [ingress_gw.local_subnet](resources--gcp_vpc_site--reference--group-003.md#canonical-325362b4832bfecf2eb7575683ed7ee25e82653d6d4a74cdfb4eb96c6486b0d3)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-6b40ff4816aae52103a699e77462c8e4dc38a9a5d5eec96994ad345d1b74f6a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92fa71fffdc6d037d604dbf245df436f7a57f64030d88c56cb29de2573cd69b6"></a>

## ingress_gw.performance_enhancement_mode — ingress_gw.performance_enhancement_mode / 9821b7630e44 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- ingress_gw.performance_enhancement_mode

<a id="canonical-9df052fe637a6042cce70ccea5ce671ab706ca318d367398ea06bbca947323bd"></a>

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

<a id="canonical-66d8700d707417ed9c4e22063a3796b3a55dd20af8627a712811e365d4407b3f"></a>

## Direct properties — ingress_gw.performance_enhancement_mode / 9821b7630e44 / 3

- [perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-58b8a004db0459507ad334347faf232588c5868d33fbc558dfa94e85e164aaab): complete subsection reference.

- [perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-184393feb1d3072594271ed179eff6416b7ea649c79ea1fba9ab5ad69ad56a83): complete subsection reference.

<a id="canonical-4c383f2cbcc0a9e7d2d8a00095c076005e1b5219facd69114656d61c4c1f8e7d"></a>

## Next pages — ingress_gw.performance_enhancement_mode / 9821b7630e44 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-58b8a004db0459507ad334347faf232588c5868d33fbc558dfa94e85e164aaab)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-184393feb1d3072594271ed179eff6416b7ea649c79ea1fba9ab5ad69ad56a83)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-58b8a004db0459507ad334347faf232588c5868d33fbc558dfa94e85e164aaab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95a084dbff0e63cc916913c7a170f9f7cc7395f017a38464e7ed9a733c35d2c0"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / b8fd186858af / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [ingress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-6b40ff4816aae52103a699e77462c8e4dc38a9a5d5eec96994ad345d1b74f6a0)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-34dd71794ad6565e1e02f055db37d6d1f90964cf2f5f53d859cadbcc71691aaa"></a>

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

<a id="canonical-cf356f626e662c814014574b5c8dcee1c18ee85d940e3afef7453c402ca2fea1"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / b8fd186858af / 3

- [jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-da40ce72da2d2a9ecd661b7b0ad34c8557a616909745d179b8f7d19556a01572): complete subsection reference.

- [no_jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-384d82f598a957e1a04d37a78e89c0c04fd35761f86ed9c30ff4ec08c20aaf19): complete subsection reference.

<a id="canonical-8bd95b8c9492db9263ad883d3fc477f4914f7db2bd37a79b907e805b87f5434c"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / b8fd186858af / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-da40ce72da2d2a9ecd661b7b0ad34c8557a616909745d179b8f7d19556a01572)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--gcp_vpc_site--reference--group-003.md#canonical-384d82f598a957e1a04d37a78e89c0c04fd35761f86ed9c30ff4ec08c20aaf19)
- [ingress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-6b40ff4816aae52103a699e77462c8e4dc38a9a5d5eec96994ad345d1b74f6a0)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-da40ce72da2d2a9ecd661b7b0ad34c8557a616909745d179b8f7d19556a01572"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b47f3e0b1fde65b20c1e8d23a03c5df3a833dc180ba621b7d5c2390daceb5733"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 4922abeda27c / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [ingress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-6b40ff4816aae52103a699e77462c8e4dc38a9a5d5eec96994ad345d1b74f6a0)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-58b8a004db0459507ad334347faf232588c5868d33fbc558dfa94e85e164aaab)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-e062985941c06d8a089bd40059cdd03d8192640f484b813bf0ed18bde53af79e"></a>

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

<a id="canonical-71b284a584aa8b51189985a6356e1950b3f3c8bf0b608ae81deba18f957ad807"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 4922abeda27c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bc7a727fb2f3ec5b21c2975d135fcb6708d9d7620df28f00b87abde763b1ab31"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 4922abeda27c / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-58b8a004db0459507ad334347faf232588c5868d33fbc558dfa94e85e164aaab)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-384d82f598a957e1a04d37a78e89c0c04fd35761f86ed9c30ff4ec08c20aaf19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-408ac91af40a5636f2d8cc059c3dd84dfcb665aa5988d597675e796158ddcd09"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 7ce3413574de / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [ingress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-6b40ff4816aae52103a699e77462c8e4dc38a9a5d5eec96994ad345d1b74f6a0)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-58b8a004db0459507ad334347faf232588c5868d33fbc558dfa94e85e164aaab)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-62aa8d0f40e4e868580d32b4a343a97ecb48cb5caabdbffbaaf533f104936d42"></a>

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

<a id="canonical-2ce3fdbe6d1d32672d2a645a53673b88843c58cd41b2a1f3a77802ce016a905a"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 7ce3413574de / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-25d27b51db6e1506d6858ea23c2d430242281601b5dc9f8b3b23f1452ed8c6f9"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 7ce3413574de / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-58b8a004db0459507ad334347faf232588c5868d33fbc558dfa94e85e164aaab)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-184393feb1d3072594271ed179eff6416b7ea649c79ea1fba9ab5ad69ad56a83"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-63d5e104214f80bd055d9a412ff31e8d77452cc22f03894d51e97bc97188b5fc"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 4583ae6c7096 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [ingress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-6b40ff4816aae52103a699e77462c8e4dc38a9a5d5eec96994ad345d1b74f6a0)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-93722bfea2fc8324e764935d6b69cbdb989605236b55d4995abf4de1fd880ece"></a>

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

<a id="canonical-51b73980b5eca025a03b8a495e4d655639612c2a7db069baae989d7c02d963b0"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 4583ae6c7096 / 3

- [jumbo_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-2246a16a57eb68bf2055a6e0b83372676a8651367b29ac135ccb65b3d618ad0b): complete subsection reference.

- [jumbo_enabled](resources--gcp_vpc_site--reference--group-003.md#canonical-3edadeebf49d7cfdaa46aa088e3758684324aba3534dde10e00ee84c512e9eff): complete subsection reference.

<a id="canonical-f0080d374319ecd5b451ecbecb6c67713453b0b226d55ccd0faceb5f5589fdb7"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 4583ae6c7096 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-2246a16a57eb68bf2055a6e0b83372676a8651367b29ac135ccb65b3d618ad0b)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--gcp_vpc_site--reference--group-003.md#canonical-3edadeebf49d7cfdaa46aa088e3758684324aba3534dde10e00ee84c512e9eff)
- [ingress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-6b40ff4816aae52103a699e77462c8e4dc38a9a5d5eec96994ad345d1b74f6a0)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-2246a16a57eb68bf2055a6e0b83372676a8651367b29ac135ccb65b3d618ad0b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d2dc55938191d1df39c52554643b4a9627f00a61ac0704447c82454745cb585"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 7ec9c0f95593 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [ingress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-6b40ff4816aae52103a699e77462c8e4dc38a9a5d5eec96994ad345d1b74f6a0)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-184393feb1d3072594271ed179eff6416b7ea649c79ea1fba9ab5ad69ad56a83)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-eec2dd07ca556a8626e84ded35aecdf7d34599b1deb2554f6bb276a2a5675149"></a>

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

<a id="canonical-0c40282c562fbc152368833de2a9cf3fd36cad8afb0b1508f3dbf6ab8f7cf4bd"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 7ec9c0f95593 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-263d1811503ff5e109f110f8feb4dbd018e01b3aa6e21dbfa8f4a02d2fd2dfbc"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 7ec9c0f95593 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-184393feb1d3072594271ed179eff6416b7ea649c79ea1fba9ab5ad69ad56a83)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-3edadeebf49d7cfdaa46aa088e3758684324aba3534dde10e00ee84c512e9eff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-359c4a44559c71407e37c0cb389f544db79e76f2befe1b2ea3b5f5e37df3136b"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 5aaf30e940ba / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [ingress_gw](resources--gcp_vpc_site--reference--group-003.md#canonical-a9aa8c9aee518ea9d086d1c0672f23c377a919ba9f067fad366a259499a872cc)
- [ingress_gw.performance_enhancement_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-6b40ff4816aae52103a699e77462c8e4dc38a9a5d5eec96994ad345d1b74f6a0)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-184393feb1d3072594271ed179eff6416b7ea649c79ea1fba9ab5ad69ad56a83)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-62a31d96180c6c711fa1e85bd5ee7d4c904b7cfc2af7e8e7d96c3b6532270f88"></a>

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

<a id="canonical-5d3e858936987f1db8be50c11c97baf9e6741653f649032a81478f51a1fee598"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 5aaf30e940ba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a9862032ca081af9ae99c213b22ce7f2a6909e9ab4d3c7443ddf3d1bf8b0c4b3"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 5aaf30e940ba / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--gcp_vpc_site--reference--group-003.md#canonical-184393feb1d3072594271ed179eff6416b7ea649c79ea1fba9ab5ad69ad56a83)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-91992f0d045bc68652deeaf62a5b3849c65f5447dcdd89b44c6c2b97307f6efb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5b4ea45d369be19231c517de17a82b4264cb1cfc7bb6e044d23a9aaa230822db"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 729e26289a1e / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- kubernetes_upgrade_drain

<a id="canonical-a1c706be5022d2e31e11e4567d36e5b95beada9218fb2e60e10dc03c54150083"></a>

Type: `"object"`. single nested block, Optional.

Specify how worker nodes within a site will be upgraded.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_upgrade_drain",
    "enable_upgrade_drain")}
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
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

Terraform syntax:

```terraform
kubernetes_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-8374587c223fcb46fd843cac65176932b592ff757bddbb5610a1bac07275de70"></a>

## Direct properties — kubernetes_upgrade_drain / 729e26289a1e / 3

- [disable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-90f4acd7914d0f28c35b4fb24cc1fedf694120cfb171cde544c5bb6e390e302b): complete subsection reference.

- [enable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-d51765427fe8bee4c2e0557188822d9feb07ef16a2d2f094056478b75d7ffdcd): complete subsection reference.

<a id="canonical-908fd30f2abf857a15ea6a3893f4b76033891f7afb70c0b7a0a6aaa708d8e1e4"></a>

## Next pages — kubernetes_upgrade_drain / 729e26289a1e / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-90f4acd7914d0f28c35b4fb24cc1fedf694120cfb171cde544c5bb6e390e302b)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-d51765427fe8bee4c2e0557188822d9feb07ef16a2d2f094056478b75d7ffdcd)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-90f4acd7914d0f28c35b4fb24cc1fedf694120cfb171cde544c5bb6e390e302b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc12a208cc9a4268085bdca831cdc5da7a7e2a65eef873ee2f2d32f2b3a80fc6"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — kubernetes_upgrade_drain.disable_upgrade_drain / 6351372c0b9f / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-91992f0d045bc68652deeaf62a5b3849c65f5447dcdd89b44c6c2b97307f6efb)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-38422e31f184402be45a057d8b662544bead9089f37ea79391f371dd32c2dcda"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

<a id="canonical-aeb6eefe9ec779ea64d18c42ed3e2f8c903662a158c7f2ff6524212211be2179"></a>

## Direct properties — kubernetes_upgrade_drain.disable_upgrade_drain / 6351372c0b9f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa278571d66dc24cc52d7c6946ff34282b352f598bfc4dd4edea9064b36eaf6c"></a>

## Next pages — kubernetes_upgrade_drain.disable_upgrade_drain / 6351372c0b9f / 4

- [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-91992f0d045bc68652deeaf62a5b3849c65f5447dcdd89b44c6c2b97307f6efb)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-d51765427fe8bee4c2e0557188822d9feb07ef16a2d2f094056478b75d7ffdcd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3b0e7935dcc9cddebf46d49b5f738fcc391b3afbda700521fec111227c4fc22e"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — kubernetes_upgrade_drain.enable_upgrade_drain / 5592820bf173 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-91992f0d045bc68652deeaf62a5b3849c65f5447dcdd89b44c6c2b97307f6efb)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-541dbdb32270271694c2e6a6d17a195867b124117f93e49b1f9f923de0ceeceb"></a>

Type: `"object"`. single nested block, Optional.

Specify batch upgrade settings for worker nodes within a site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("drain_node_timeout"),
  validators.ConflictingObjectAttributes("disable_vega_upgrade_mode",
    "enable_vega_upgrade_mode"),
  validators.ConflictingObjectAttributes("drain_max_unavailable_node_count",
    "drain_max_unavailable_node_percentage")}
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
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

Terraform syntax:

```terraform
enable_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-4b22cd77eb81fc211817db5f40d2f19c88c93e48dcce5e696484df44cd7c4ecc"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain / 5592820bf173 / 3

- [disable_vega_upgrade_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-98e81091ddb9086904431bfb135b6c24f8e8695590794c3bb67a301a1958dea6): complete subsection reference.

<a id="canonical-296a7e4cb6785e58cd2982fdee98574a226ca7a06395ebe06c4373d4e44b7edf"></a>

<a id="canonical-e497950dc889a9132fddf4ee37bd2421acbc10b29156bacfd6e2e9d9d149fe40"></a>

## drain_max_unavailable_node_count property — kubernetes_upgrade_drain.enable_upgrade_drain / 5592820bf173 / 4

Type: `"number"`. Optional.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
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
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-71b23bd989f8931227a0110a2c702a358faca08a6c5a8ce18ea89be67af60920"></a>

<a id="canonical-ef18a367f17c0a94f0e47ee5f73d034f507832f6ab8b0f104d084e781fdc847c"></a>

## drain_max_unavailable_node_percentage property — kubernetes_upgrade_drain.enable_upgrade_drain / 5592820bf173 / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-2a68c3ec4e2c2d7878c17826f3e7fdb6a3c016e174c61eea8c8bc41a3af0a176"></a>

<a id="canonical-6d3018a1d62624b132d7d58c25044abe35c9645ee071292884d11872e4efb1a7"></a>

## drain_node_timeout property — kubernetes_upgrade_drain.enable_upgrade_drain / 5592820bf173 / 6

Type: `"number"`. Optional.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 900),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-98ff247948a8fb743a060ee5545ffb81587224c82694c6117090a4a7b919e0e4): complete subsection reference.

<a id="canonical-3fcc08455226c4b326eee2570a301e9c26c2649543e16660e5ff07d5ffecc6de"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain / 5592820bf173 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-98e81091ddb9086904431bfb135b6c24f8e8695590794c3bb67a301a1958dea6)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-98ff247948a8fb743a060ee5545ffb81587224c82694c6117090a4a7b919e0e4)
- [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-91992f0d045bc68652deeaf62a5b3849c65f5447dcdd89b44c6c2b97307f6efb)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-98e81091ddb9086904431bfb135b6c24f8e8695590794c3bb67a301a1958dea6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27a23a8fbdf4cc6d4b7d49175afb97a2015099595630a298287c849b888384fc"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 074c2d914003 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-91992f0d045bc68652deeaf62a5b3849c65f5447dcdd89b44c6c2b97307f6efb)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-d51765427fe8bee4c2e0557188822d9feb07ef16a2d2f094056478b75d7ffdcd)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-b3dead804495e953126f6405332e8d5e29d2f1f30f661d28f47a1c3ba6f82dc4"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

<a id="canonical-c9d56dd50965befb81ff0d9dd9fb7fb96a550ca06ba72852e3af2eee91068537"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 074c2d914003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2336d73c24b68d07b8302682b079ac25876e4d31ec081a428a5189599f8fc3c3"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 074c2d914003 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-d51765427fe8bee4c2e0557188822d9feb07ef16a2d2f094056478b75d7ffdcd)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-98ff247948a8fb743a060ee5545ffb81587224c82694c6117090a4a7b919e0e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25f5c3d7bd8db71a2f94b856e5dd79e1a0894690cdcffe6fd4626a37126afac3"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 5668eaec4932 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [kubernetes_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-91992f0d045bc68652deeaf62a5b3849c65f5447dcdd89b44c6c2b97307f6efb)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-d51765427fe8bee4c2e0557188822d9feb07ef16a2d2f094056478b75d7ffdcd)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-c782b3f3c8d9c555bf52dae72211443f85ba398a5f0de682df207602638a4d27"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

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
enable_vega_upgrade_mode = {}
```

<a id="canonical-f009b153639184674db5b7a673e8cb0a7d3b19defe68d0f403a0f266fcf6abcf"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 5668eaec4932 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7fd15ce608a555b062047a0b1f88de085bc9f7726c4eb2cca3f4f7f268cf47ed"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / 5668eaec4932 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--gcp_vpc_site--reference--group-003.md#canonical-d51765427fe8bee4c2e0557188822d9feb07ef16a2d2f094056478b75d7ffdcd)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-4b1d87b55107bf727296e0f74b17d3ab715fa9ab4ce733b2a51208d76bc40e16"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c6724ca9f3199df71bd126c96b566888c30cd054f7ea0d7952fd320d94c82f7"></a>

## log_receiver — log_receiver / 3bbca82f1a9d / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- log_receiver

<a id="canonical-a26d857017679cfda085e0f0a49903fcc032db4921c1de00e087766ff383fd49"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [log_receiver](resources--gcp_vpc_site--reference--group-003.md#canonical-a26d857017679cfda085e0f0a49903fcc032db4921c1de00e087766ff383fd49)
- [logs_streaming_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-f8b5d095a44382632a1481ab7b7f448fa5e9a46803bfd87fe1e8db304f2c49c1)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-b0887aec9c25e3f1622920634c9cd478a5c2a073151bbb103258f87eaa240e8a"></a>

## Direct properties — log_receiver / 3bbca82f1a9d / 3

<a id="canonical-f2f5d2a070cfa515203de8cb8ee80cf3abca174166b5a2e8ef9173d2b7f1fe5f"></a>

<a id="canonical-2056051a971ec1d07b6aada15e5b5819e901e8faaacb32968671f6bf2e7bdea2"></a>

## name property — log_receiver / 3bbca82f1a9d / 4

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

<a id="canonical-7e41f6b87b4b787e7dc7bd9ca47622c05c4bd48bf3d60078bf5387aead3c5c8c"></a>

<a id="canonical-6f4b64eff218e255a1162e0f75d5db2fff6407f793876629c6d667e2269579bd"></a>

## namespace property — log_receiver / 3bbca82f1a9d / 5

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

<a id="canonical-c48869e4cb461325b044f277716159ea150be8d6a41e82aeead6f4f56e481987"></a>

<a id="canonical-e92fdd11a13bc9ce821f1d1b3840de15a923769042e54804b28afc9bf3d16601"></a>

## tenant property — log_receiver / 3bbca82f1a9d / 6

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

<a id="canonical-ac2b12e22fdd2000fc43ee068b047322bf8fae983f093889574aed1c4cd94c6c"></a>

## Next pages — log_receiver / 3bbca82f1a9d / 7

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-c144b8deb7de45a002c3f118644029a062245be86a46e5968311c10907617f46"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7c2712beb47fdc1b4acc932c30fd3a17b36d096dd4363f3dfb6b17ae789cc172"></a>

## logs_streaming_disabled — logs_streaming_disabled / cc90dbfca0f1 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- logs_streaming_disabled

<a id="canonical-f8b5d095a44382632a1481ab7b7f448fa5e9a46803bfd87fe1e8db304f2c49c1"></a>

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
logs_streaming_disabled = {}
```

<a id="canonical-1f1fc4ccb58e40631daa2d741a38a46faa7ba58f2373d0fcfb5506823ae9c654"></a>

## Direct properties — logs_streaming_disabled / cc90dbfca0f1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-03a3b1aed4b297a8ad58f3c434bbca639a3667f9d77148a284c45269856cd513"></a>

## Next pages — logs_streaming_disabled / cc90dbfca0f1 / 4

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-321583ca89a2994ed6ab1089491a0a823abc73ba1e92d28770051589feecf20d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a4dfcea1751091e6d77374da78641c444edb732a88387c6003c4babbdc14678"></a>

## offline_survivability_mode — offline_survivability_mode / 3f3dd40a6bf3 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- offline_survivability_mode

<a id="canonical-22243c16b6b4cce94c82f3fbe7d4e9c68740dc568b60d633ec7c7c8e7befe4d4"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("enable_offline_survivability_mode",
    "no_offline_survivability_mode")}
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
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

Terraform syntax:

```terraform
offline_survivability_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c0550a58e6d03cdecfd221066d5ca0058872752dad4d7ee190ac24876b6f9ae"></a>

## Direct properties — offline_survivability_mode / 3f3dd40a6bf3 / 3

- [enable_offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-cdfdf132e3a4f2834310d9666f3df98682d28d802e50fac3a0992da301d0e731): complete subsection reference.

- [no_offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-f60541bc73a3efdd1b50ee1023b45e1ba790d5b1b1fb408035ec5cd6bd0fd167): complete subsection reference.

<a id="canonical-163771e71eb2c34f5a3bd8fa8732dbc5afe47e22fe41464bdff514f99a3c08a3"></a>

## Next pages — offline_survivability_mode / 3f3dd40a6bf3 / 4

- [offline_survivability_mode.enable_offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-cdfdf132e3a4f2834310d9666f3df98682d28d802e50fac3a0992da301d0e731)
- [offline_survivability_mode.no_offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-f60541bc73a3efdd1b50ee1023b45e1ba790d5b1b1fb408035ec5cd6bd0fd167)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-cdfdf132e3a4f2834310d9666f3df98682d28d802e50fac3a0992da301d0e731"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-105c1199d64076dafe4ced2b33f2a057a7ef3ff221ff86df3776df326cd39932"></a>

## offline_survivability_mode.enable_offline_survivability_mode — offline_survivability_mode.enable_offline_survivability_mode / fd373a65dfa2 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-321583ca89a2994ed6ab1089491a0a823abc73ba1e92d28770051589feecf20d)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-229b89e58d72d2433bb0c5ac8709d6c31d165994c49a623f7f2b6bf0fee71835"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

<a id="canonical-07f106c53aa59cb3c857591215647f3bbbdca3713be7471c53bac6376f5ed467"></a>

## Direct properties — offline_survivability_mode.enable_offline_survivability_mode / fd373a65dfa2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-95847416491c9e3b98d6edb5aff2bc80d93980f412d936fef797a4282937224c"></a>

## Next pages — offline_survivability_mode.enable_offline_survivability_mode / fd373a65dfa2 / 4

- [offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-321583ca89a2994ed6ab1089491a0a823abc73ba1e92d28770051589feecf20d)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-f60541bc73a3efdd1b50ee1023b45e1ba790d5b1b1fb408035ec5cd6bd0fd167"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94f0f0e838194c4966ee3ff2f71e3f92cfff6a50bd01db7b1fa59969a52cf4d6"></a>

## offline_survivability_mode.no_offline_survivability_mode — offline_survivability_mode.no_offline_survivability_mode / 6e46eecdbd24 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-321583ca89a2994ed6ab1089491a0a823abc73ba1e92d28770051589feecf20d)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-499d05184531855d870c44e4d7435c92ad1107f7d67d61eb82147726134b1198"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no offline survivability mode.

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
no_offline_survivability_mode = {}
```

<a id="canonical-a329841e4291052e67b576fe803968bac6149546f18ef226b9f6e79c6d28bfcd"></a>

## Direct properties — offline_survivability_mode.no_offline_survivability_mode / 6e46eecdbd24 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2bb1850ed87d968e6b094579fbe81272656951d0937450f08a103833807cde50"></a>

## Next pages — offline_survivability_mode.no_offline_survivability_mode / 6e46eecdbd24 / 4

- [offline_survivability_mode](resources--gcp_vpc_site--reference--group-003.md#canonical-321583ca89a2994ed6ab1089491a0a823abc73ba1e92d28770051589feecf20d)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-eaae4c5a3696af2cdaa426ef1745cc75f8cc67b832eed73e1b84ffbca525c28b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2241194fe93bb59b878a6744114c5f4bffd927f43e7d093b21bfd011b94218c9"></a>

## os — os / c29aed791991 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- os

<a id="canonical-88ee877961aa11a6145caff7451f332227f90693cced62a81bcf0b7223977d08"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

Terraform syntax:

```terraform
os {
  # Configure direct properties listed below.
}
```

<a id="canonical-9cc1115fd4d3c29544b4dc07b50a578d02bc278971f629e92ab42ac021da0670"></a>

## Direct properties — os / c29aed791991 / 3

- [default_os_version](resources--gcp_vpc_site--reference--group-003.md#canonical-a5d27a6d9a973d0022e400aa5bf4c5842d702ec40523d3d3a9b34b8eb9d173f1): complete subsection reference.

<a id="canonical-95ae728237cbc1ff0eb7a22667f056dfcdd8ad1483c9a8148491db89ab0644c1"></a>

<a id="canonical-cb79552a499b6872b742fe4aad39e52aebfa75bb6c677e613e02d16d355792db"></a>

## operating_system_version property — os / c29aed791991 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-d03b41baefb598e57adae2a3725a3a01063d7dc9a0f8d3eb188ced70f1c37e5b"></a>

## Next pages — os / c29aed791991 / 5

- [os.default_os_version](resources--gcp_vpc_site--reference--group-003.md#canonical-a5d27a6d9a973d0022e400aa5bf4c5842d702ec40523d3d3a9b34b8eb9d173f1)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-a5d27a6d9a973d0022e400aa5bf4c5842d702ec40523d3d3a9b34b8eb9d173f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be24b82428a4bad6f6d213cc7b5f62a366a5da00098ac3acad06d59fc81b4117"></a>

## os.default_os_version — os.default_os_version / d294eee414c2 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [os](resources--gcp_vpc_site--reference--group-003.md#canonical-eaae4c5a3696af2cdaa426ef1745cc75f8cc67b832eed73e1b84ffbca525c28b)
- os.default_os_version

<a id="canonical-cdaf9a488d747065c0e2be29bf8d3ef53a4f115bc2b19f7d5f74a8441e686501"></a>

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
default_os_version = {}
```

<a id="canonical-d07d18fb6d3e7f83edbd1aad515630dfba64a0739b88b63df6acfc6fda4d5eb1"></a>

## Direct properties — os.default_os_version / d294eee414c2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fa3381279f456db0049dd302390d834517aaf30b396def48aac2d39703d438bc"></a>

## Next pages — os.default_os_version / d294eee414c2 / 4

- [os](resources--gcp_vpc_site--reference--group-003.md#canonical-eaae4c5a3696af2cdaa426ef1745cc75f8cc67b832eed73e1b84ffbca525c28b)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-60d557d0083e9ae6c90a25feb164999e2e99fcf4ad1174a90fa437468a9395e1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09b67b143c0752f1f96b92c6ba11e3a8e0e1f21abd23b8dc5d3bddbb3dd9788d"></a>

## private_connect_disabled — private_connect_disabled / 495582ff92da / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- private_connect_disabled

<a id="canonical-afc6f45f786e41688d334f4267b4908fd8fc55ce3574f716dd799807571516cb"></a>

Type: `["object", {}]`. Optional.

\[OneOf: private\_connect\_disabled, private\_connectivity\] Enable this option

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

OneOf alternatives in this subsection:

- [private_connect_disabled](resources--gcp_vpc_site--reference--group-003.md#canonical-afc6f45f786e41688d334f4267b4908fd8fc55ce3574f716dd799807571516cb)
- [private_connectivity](resources--gcp_vpc_site--reference--group-003.md#canonical-06c998e2599f7224ad0891dff32b8d5e831d315c925369a3bc90d2a059c30193)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
private_connect_disabled = {}
```

<a id="canonical-2a6fd13c650e18cff6778ca68d81e1415f18b653dc3d9163a6748e176b946b98"></a>

## Direct properties — private_connect_disabled / 495582ff92da / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7cb7f0b9816c55b176c162707ccb1b768f51033f2b28348f6f8537c1a0e2b8aa"></a>

## Next pages — private_connect_disabled / 495582ff92da / 4

- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-2b156b46727ac0a204195e7e776cb3951197b186926382e9a9299718d824f79f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3e1e49386f3a38e9827788f48285655d7527b08c084f7c3398a2b90c279ca29"></a>

## private_connectivity — private_connectivity / aba1fdb3bc95 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- private_connectivity

<a id="canonical-06c998e2599f7224ad0891dff32b8d5e831d315c925369a3bc90d2a059c30193"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for private connectivity.

Upstream description:

Private Connect Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside",
    "outside")}
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
  "x-ves-oneof-field-network_options": "[\"inside\",\"outside\"]"
}
```

Terraform syntax:

```terraform
private_connectivity {
  # Configure direct properties listed below.
}
```

<a id="canonical-e787a966c9233cc3e77262708e0faa5a27e784c78c099795dac2a707c74b654f"></a>

## Direct properties — private_connectivity / aba1fdb3bc95 / 3

- [cloud_link](resources--gcp_vpc_site--reference--group-003.md#canonical-1fa64a77b6d2b86450cb828d171ddfe9ced979fe656161a2c7e669bbb5f8120f): complete subsection reference.

- [inside](resources--gcp_vpc_site--reference--group-003.md#canonical-410be4eefccb5fabe093603df1022a7d3388ec5479e26bf1ba04a934f52212b1): complete subsection reference.

- [outside](resources--gcp_vpc_site--reference--group-003.md#canonical-2b9356103cd5a30be12f088343c03b418aba6b55cd6f22beaf370186386802ab): complete subsection reference.

<a id="canonical-142fcc5bd8aeadc03d0456472bb9394d685d4ebf6fb387360cde67d6fde6def5"></a>

## Next pages — private_connectivity / aba1fdb3bc95 / 4

- [private_connectivity.cloud_link](resources--gcp_vpc_site--reference--group-003.md#canonical-1fa64a77b6d2b86450cb828d171ddfe9ced979fe656161a2c7e669bbb5f8120f)
- [private_connectivity.inside](resources--gcp_vpc_site--reference--group-003.md#canonical-410be4eefccb5fabe093603df1022a7d3388ec5479e26bf1ba04a934f52212b1)
- [private_connectivity.outside](resources--gcp_vpc_site--reference--group-003.md#canonical-2b9356103cd5a30be12f088343c03b418aba6b55cd6f22beaf370186386802ab)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-1fa64a77b6d2b86450cb828d171ddfe9ced979fe656161a2c7e669bbb5f8120f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0b43e485a87432d028b144a7031ea9bf112651aeed3ef421b4a76bbd6580e11"></a>

## private_connectivity.cloud_link — private_connectivity.cloud_link / 65b9aac58966 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [private_connectivity](resources--gcp_vpc_site--reference--group-003.md#canonical-2b156b46727ac0a204195e7e776cb3951197b186926382e9a9299718d824f79f)
- private_connectivity.cloud_link

<a id="canonical-2400ba220f3ac6c28824d7698e46f773874d7129ed4379c3a45c2a254a2b3fce"></a>

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
cloud_link {
  # Configure direct properties listed below.
}
```

<a id="canonical-e1ea71856f2a0b86716d2a4438a45a6a8883f1153fe08d13ba4993ebfb4142d1"></a>

## Direct properties — private_connectivity.cloud_link / 65b9aac58966 / 3

<a id="canonical-fbb9c6a13582b76ad43013fdce2dd23c210dbe977bb08b3e52b95ba6b2ec5ba2"></a>

<a id="canonical-36e58fe927030d43a16ac06229d50413203322e020f6456e58d29ca04591154a"></a>

## name property — private_connectivity.cloud_link / 65b9aac58966 / 4

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

<a id="canonical-4a37d85be0d4adb91882efa12d87df43af2d88b446e06c20fa5755c739d2f2ec"></a>

<a id="canonical-b664b60d5b8d376fcb045a0c3a4d3c16aba5af31d731c74e855885044bcb9c6c"></a>

## namespace property — private_connectivity.cloud_link / 65b9aac58966 / 5

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

<a id="canonical-24e7992c2003a746297915b8543f782d15e80f095b0a6aed4f21161079305b1f"></a>

<a id="canonical-1c0f0c90bd4e136c9d8383ba56ca5f3ca3c6ff28c57c22e82fcb1a208ae18c38"></a>

## tenant property — private_connectivity.cloud_link / 65b9aac58966 / 6

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

<a id="canonical-156438bbdcb8c02157fb4afb9432ec67f3d1ec0c4037e70e202d4e6534dc54a7"></a>

## Next pages — private_connectivity.cloud_link / 65b9aac58966 / 7

- [private_connectivity](resources--gcp_vpc_site--reference--group-003.md#canonical-2b156b46727ac0a204195e7e776cb3951197b186926382e9a9299718d824f79f)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-410be4eefccb5fabe093603df1022a7d3388ec5479e26bf1ba04a934f52212b1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-64e98a9ba2075c47fa94f325ccc815518d519441048bb147a045fcefe77c9f9b"></a>

## private_connectivity.inside — private_connectivity.inside / 5dc399fd75b9 / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [private_connectivity](resources--gcp_vpc_site--reference--group-003.md#canonical-2b156b46727ac0a204195e7e776cb3951197b186926382e9a9299718d824f79f)
- private_connectivity.inside

<a id="canonical-cb234dcf5f479a2e02f80201badd9438dc9b773dc90dc0f45500a7aa9bd25d0a"></a>

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
inside = {}
```

<a id="canonical-9f6f9cd3544dd82a8872759909f345f4e68bef571c661afb0d7ac6c2b5c53133"></a>

## Direct properties — private_connectivity.inside / 5dc399fd75b9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cbf7381febd5d40a82c96c43bf593e730ce6783963ec910da3ed2323d4d317a3"></a>

## Next pages — private_connectivity.inside / 5dc399fd75b9 / 4

- [private_connectivity](resources--gcp_vpc_site--reference--group-003.md#canonical-2b156b46727ac0a204195e7e776cb3951197b186926382e9a9299718d824f79f)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-2b9356103cd5a30be12f088343c03b418aba6b55cd6f22beaf370186386802ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9d6e23a2820648a6b5a217e84a28ea86aff223dac6b3382b810c0842812362df"></a>

## private_connectivity.outside — private_connectivity.outside / c3e666eb53ea / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- [private_connectivity](resources--gcp_vpc_site--reference--group-003.md#canonical-2b156b46727ac0a204195e7e776cb3951197b186926382e9a9299718d824f79f)
- private_connectivity.outside

<a id="canonical-b172837f4ecae9888073aff680024aca5ea0c9414ce319ac916e9684264b5db2"></a>

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
outside = {}
```

<a id="canonical-ac41b70fa0ae23d63f97f401d2f78a917f06e6c276466c87a5af70e410340d52"></a>

## Direct properties — private_connectivity.outside / c3e666eb53ea / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2e0e64542e9f4bca2de3eb6ab1f4e1bf1c10b2bfaace61cb5571b012a58ffe5c"></a>

## Next pages — private_connectivity.outside / c3e666eb53ea / 4

- [private_connectivity](resources--gcp_vpc_site--reference--group-003.md#canonical-2b156b46727ac0a204195e7e776cb3951197b186926382e9a9299718d824f79f)
- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)

<a id="canonical-267f967c717532fd9b7a7d530db900ad6f3ed5285284fb0cc81876b19418dec2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-782dc2841643900f3132f7ede98d11cd875bcb4c9c8a997e9903412ef8f6c343"></a>

## sw — sw / 761d5743321b / 2

Breadcrumbs:

- [xcsh_gcp_vpc_site](../resources/gcp_vpc_site.md#canonical-1d0101b699d1c052a2b3d4e3aca83ac1574445d59ddcc26b8c6bf8d1a705905a)
- [Property reference](resources--gcp_vpc_site--reference--group-001.md#canonical-915674fe4d0d219acffc2e702613b25031f70cf2f9a438ea6c7f28c7eeb7e6e6)
- sw

<a id="canonical-ac299786eb3acecb781355b7cb3843424b0b2f0aed576816a040d543e47dbb69"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
sw {
  # Configure direct properties listed below.
}
```

<a id="canonical-32053a56edabb75d1dc5d4378d3a56b90f769528006af1743bffc01e32fd2b42"></a>

## Direct properties — sw / 761d5743321b / 3

- [default_sw_version](resources--gcp_vpc_site--reference--group-004.md#canonical-84b59a58878fd650e747bc25e45272cdd659e0b6d53b9d1e4d65fb45118fbf52): complete subsection reference.

<a id="canonical-db56b7c2c6159bb17fc629e032371a6860314a8be27262c65fa4a01eed0ab25a"></a>
