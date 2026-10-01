---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-2ac6ff45e8a223b755d1f3e47d66c564babcb5d115bc891ee7205fa0d75cbf8c"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 86c09ab364af / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-67490586153aadb25a2de4cd6a0489df4e87614263e61ecb8de5a6bda1576040"></a>

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

<a id="canonical-26eff21f9750ca4625b8044d7a9cc915de4f3a970ef161761ef5a6ed001e320f"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 86c09ab364af / 3

<a id="canonical-29f6ae9b59535bfc8cdc156abc30a111be2defba6304e00eced2489bf4eddc97"></a>

<a id="canonical-1b61e8d9d99021a16014d0bb414e9a1b0a23145221300f21b1910cac75625e43"></a>

## attrs property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 86c09ab364af / 4

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

- [labels](resources--azure_vnet_site--reference--group-005.md#canonical-56d04c943b8d8fadf9564142c2d24444e2ece99c224d26a0d7ba6cd3b8a35c10): complete subsection reference.

- [nexthop](resources--azure_vnet_site--reference--group-005.md#canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2): complete subsection reference.

- [subnets](resources--azure_vnet_site--reference--group-005.md#canonical-b55e9e6920ce9d2798e9fc486217d07e23a23c5eda496855ac45ab31fdbd77fe): complete subsection reference.

<a id="canonical-54faf846d1022727df1cff7b45268531bd91445fca78f29634e26a4b9bb99647"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route / 86c09ab364af / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](resources--azure_vnet_site--reference--group-005.md#canonical-56d04c943b8d8fadf9564142c2d24444e2ece99c224d26a0d7ba6cd3b8a35c10)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-005.md#canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-005.md#canonical-b55e9e6920ce9d2798e9fc486217d07e23a23c5eda496855ac45ab31fdbd77fe)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-56d04c943b8d8fadf9564142c2d24444e2ece99c224d26a0d7ba6cd3b8a35c10"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11d4da57d54a65b7204c6a4f0feaf67bc9cc57cc92ad4b9f25b90f3fb44e89a5"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / 81506c87771f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-59c4aaac9f660b64c3ceeb0fd4429f080e1e5de54ca47e2c12271e03d7546902"></a>

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

<a id="canonical-46b0b878f8f8f402d0193921bdf0bd4eea9f01659d7a1c95226d6e09eec8e9ca"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / 81506c87771f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c3ab81a74b6221cb1b137ea8a4d307366dcee63654c1311420b8229a52ba7cb0"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.la / 81506c87771f / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0fce1021c4d63f9f42c0209ab3523826d460236d3e56fb7c771838e760a0d386"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 2f9f9fb2c29c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-4fd3f9a010b20942a147d321d8bdcd9f9ec551b18b83d3acd2ff3828c94d7d2a"></a>

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

<a id="canonical-1b440b587821cc101b604fde2a5c8575f0babb5eef5e47db9a1dcd99226b698a"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 2f9f9fb2c29c / 3

- [interface](resources--azure_vnet_site--reference--group-005.md#canonical-69740257b79071467c779a27dec2a0250b1b3fa9ef7921d6c4b4093897ac965f): complete subsection reference.

- [nexthop_address](resources--azure_vnet_site--reference--group-005.md#canonical-8cf9ebe161d6097f5b63f4cc461a5b37feb992569343aec9aa5d28514f9eb3bb): complete subsection reference.

<a id="canonical-ef4909b0e66a941c31d2a6cd43311b8e08ca898120faf1fe645749c3703ce11f"></a>

<a id="canonical-3a508ad67853254b8a82b03b1d96f19144988db5b40bd22f0522c196c53ef2f6"></a>

## type property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 2f9f9fb2c29c / 4

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

<a id="canonical-632c6848c7912acdd73b16bef82109702d0e9933f863837dae2d363f62450559"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 2f9f9fb2c29c / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--azure_vnet_site--reference--group-005.md#canonical-69740257b79071467c779a27dec2a0250b1b3fa9ef7921d6c4b4093897ac965f)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-005.md#canonical-8cf9ebe161d6097f5b63f4cc461a5b37feb992569343aec9aa5d28514f9eb3bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-69740257b79071467c779a27dec2a0250b1b3fa9ef7921d6c4b4093897ac965f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a8f3113a00f85bc30c0c2c721ce74ef1e43ac3503732dffb8c59bc46fa8b6b2b"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / f4a325a73f60 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-005.md#canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-2625962a0cb5a78e5be46d0a8e93936ce98733f2c760f8692d819079414875e2"></a>

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

<a id="canonical-68d00e4bae0dc577dc27ce28cf88e443a92a1077ae0ebcc5b201bfd00796e4a6"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / f4a325a73f60 / 3

<a id="canonical-f756cf75fd281cda531664822d7020224ce653a26b264211762068abe9925726"></a>

<a id="canonical-be849321393eb849f1534fc5c689bea46c312bcc35eb11f7eae35cb1ea60d3df"></a>

## kind property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / f4a325a73f60 / 4

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

<a id="canonical-4aad97dc22e5c36331f591a6a269512d8a0a31d1e70667e7fa850dff2fb14c85"></a>

<a id="canonical-3e353faf0d9bae2be7edaafbf25044803a9f4bae9639e3a91e8746763443dd9d"></a>

## name property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / f4a325a73f60 / 5

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

<a id="canonical-4bbadfd3f0c78424678622f1c0ffc77413455dc02085ca19845640e62cc5f6de"></a>

<a id="canonical-41636dfe52e8221c6679fc0efcd1cd73254fd2f011d4c7d2253c4a664da7c20d"></a>

## namespace property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / f4a325a73f60 / 6

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

<a id="canonical-6ecc4ffec507ebf4cc77db4d552abf787302ca645d8917c33a1f924581ab54d1"></a>

<a id="canonical-4d0f38d80071b6d26defa6aef11b7a934cabff864f743602f441048b6714713a"></a>

## tenant property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / f4a325a73f60 / 7

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

<a id="canonical-450730d073ec8a31bd66771e88c87e841308d7c545ed01a48bebfe15ef53d17d"></a>

<a id="canonical-b1b39378cf68d05b49b0e89fbde10e17e7930e442a9e67ee0ddf4fb2ec11ca8d"></a>

## uid property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / f4a325a73f60 / 8

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

<a id="canonical-3e992eb9e856c191c9fcdaa7c7fefa1642768f82cf2d73d0d0d25fce350169c1"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / f4a325a73f60 / 9

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-005.md#canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-8cf9ebe161d6097f5b63f4cc461a5b37feb992569343aec9aa5d28514f9eb3bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67ac27f259137f20e9204b325aa8701ea667a6ddda9f35160b356d32a69f25ed"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 8417ad88d582 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-005.md#canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-c6139f619a83c020902e16ccd6dbef0caf24c4bdb44edd4d1160bd406ad74e3c"></a>

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

<a id="canonical-454388c0fe1ba40f3b20a628c2ce26651a6d5d820f5e11635a79d45e19a44591"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 8417ad88d582 / 3

- [dual_stack](resources--azure_vnet_site--reference--group-005.md#canonical-0f236fce495ef23f90a0237665230888127706652143bedf41c542c09c04985a): complete subsection reference.

- [ipv4](resources--azure_vnet_site--reference--group-005.md#canonical-fc66c104b4a616e36a51e180c364855ed5a36f5e61f46ed13b80eb7252930c2b): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-005.md#canonical-0ff1e2b8ba22e33ad58207891d75067944c516b83a889a9f1d9e0fe86a525005): complete subsection reference.

<a id="canonical-b7e533af9c0b862cf98c04f455a3e8685361c937c67132e8249e143054f5946b"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 8417ad88d582 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-005.md#canonical-0f236fce495ef23f90a0237665230888127706652143bedf41c542c09c04985a)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--azure_vnet_site--reference--group-005.md#canonical-fc66c104b4a616e36a51e180c364855ed5a36f5e61f46ed13b80eb7252930c2b)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--azure_vnet_site--reference--group-005.md#canonical-0ff1e2b8ba22e33ad58207891d75067944c516b83a889a9f1d9e0fe86a525005)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-005.md#canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-0f236fce495ef23f90a0237665230888127706652143bedf41c542c09c04985a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2fee3f777704dc7659591fc4ef0b91dce97a3018dc87352a597048e390d0c462"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 0f9ab9b0bbc6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-005.md#canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-005.md#canonical-8cf9ebe161d6097f5b63f4cc461a5b37feb992569343aec9aa5d28514f9eb3bb)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-9f81a6fea3418af5a044b41d8fc89f081f65ec1fa5d5a80bff480b438430b756"></a>

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

<a id="canonical-192b665b892ca5444a4f51f2047fca245c684c9382ad3e87cd78c21c8b1a0b89"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 0f9ab9b0bbc6 / 3

- [ipv4](resources--azure_vnet_site--reference--group-005.md#canonical-5f20dc7c7bc4efc4277563634226ba10d650ff3e6381a182734302f7ad146126): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-005.md#canonical-1bbd75dc587deb9bf529978b34cbaf95b966da26ae890fb092b4487b2362806a): complete subsection reference.

<a id="canonical-7d7e5160854ca178f62e218d944cbcea52e3026876cd10e05fdefbda9ace25d7"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 0f9ab9b0bbc6 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--azure_vnet_site--reference--group-005.md#canonical-5f20dc7c7bc4efc4277563634226ba10d650ff3e6381a182734302f7ad146126)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--azure_vnet_site--reference--group-005.md#canonical-1bbd75dc587deb9bf529978b34cbaf95b966da26ae890fb092b4487b2362806a)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-005.md#canonical-8cf9ebe161d6097f5b63f4cc461a5b37feb992569343aec9aa5d28514f9eb3bb)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-5f20dc7c7bc4efc4277563634226ba10d650ff3e6381a182734302f7ad146126"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ccd02ffa66f6e407d3f0ff922a215c065507400eca54d53957a5e97029af7826"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 8fa738d7de8b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-005.md#canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-005.md#canonical-8cf9ebe161d6097f5b63f4cc461a5b37feb992569343aec9aa5d28514f9eb3bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-005.md#canonical-0f236fce495ef23f90a0237665230888127706652143bedf41c542c09c04985a)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-c52c9517662c1ad0670dfee8bf908b7463e101aa5e66cbae72f1fd8211f383b4"></a>

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

<a id="canonical-f215c57b8ec4f34337e2621c0d16ac1f3aa9f747d7b8dd5935cb0f14779dc509"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 8fa738d7de8b / 3

<a id="canonical-fd0b52d2d68e3e899a76c21fd0c7f10bfa74904e1050f2d0e5ceb4a56f1deaad"></a>

<a id="canonical-fe6b1a4945d9baa3469621854f28b8c24fec84968ab9c7e993c44f178615b291"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 8fa738d7de8b / 4

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

<a id="canonical-6a291037b6068584643cfc77b18ab627eb22eb6ea6b7c56d4b4eb1befbaf6094"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 8fa738d7de8b / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-005.md#canonical-0f236fce495ef23f90a0237665230888127706652143bedf41c542c09c04985a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-1bbd75dc587deb9bf529978b34cbaf95b966da26ae890fb092b4487b2362806a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3fafd4e61bf5744bdca0b11384ea5b23315041463acc3b48fa88239a93143a4c"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 432e76d24fa5 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-005.md#canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-005.md#canonical-8cf9ebe161d6097f5b63f4cc461a5b37feb992569343aec9aa5d28514f9eb3bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-005.md#canonical-0f236fce495ef23f90a0237665230888127706652143bedf41c542c09c04985a)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-4a48ce08c645db341a3174ab16b982e0029e45b1481b40b7870b3a36d9e7e2c0"></a>

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

<a id="canonical-0d0ede8f33e98e6fbf1eabebec284b4a210e5403b7a0f3ff36f21eee3776b9c5"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 432e76d24fa5 / 3

<a id="canonical-2b5c5a39c96219c55f4c42699819c78f70e431143656eaff504815463f678bd2"></a>

<a id="canonical-dfa996fd7a83aed61ec99535e394da5ddedc9208c88b3181d2f8f3262d87799e"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 432e76d24fa5 / 4

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

<a id="canonical-f44970a8dabd195e565d68fb78a0569f0878202c84555f5a641b51ecf3a2d605"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 432e76d24fa5 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-005.md#canonical-0f236fce495ef23f90a0237665230888127706652143bedf41c542c09c04985a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-fc66c104b4a616e36a51e180c364855ed5a36f5e61f46ed13b80eb7252930c2b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df6d1e181ddb0f5376d965dd63d8038c15607eadd9f300d1fdf0691a271838de"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 52a2e7f2d2e6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-005.md#canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-005.md#canonical-8cf9ebe161d6097f5b63f4cc461a5b37feb992569343aec9aa5d28514f9eb3bb)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-97e1fa996300fae92d9dc5894daa8a5cf0ccea9f3ef7880c269143cf865edfc0"></a>

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

<a id="canonical-7003b42905e10e7ba2ade7edac5e32e576e72483a122e25047aa0bb3c56dec1b"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 52a2e7f2d2e6 / 3

<a id="canonical-e316395e5d81bc9ebfb3c0a40370cda51cda220953f7b9d76471bd6ae54ea441"></a>

<a id="canonical-3c2a0f84d73b601918759757f3230dbe90cb988e3d2fc215404e6abe9f7d48fb"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 52a2e7f2d2e6 / 4

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

<a id="canonical-366d0322e3ea02439eff089915412e6870399fa8441409d19facbe33ea1db303"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 52a2e7f2d2e6 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-005.md#canonical-8cf9ebe161d6097f5b63f4cc461a5b37feb992569343aec9aa5d28514f9eb3bb)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-0ff1e2b8ba22e33ad58207891d75067944c516b83a889a9f1d9e0fe86a525005"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6aedcaf4b09651074039470caf5a178cda6d7c206c0ba85b9297ee803600c437"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 59539f7429a6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-005.md#canonical-e2b0b07dc249d1b3f7e85d0e0ac3b144608e8b2b282ba3a5416e3fe7221403d2)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-005.md#canonical-8cf9ebe161d6097f5b63f4cc461a5b37feb992569343aec9aa5d28514f9eb3bb)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-edf9022aec4b32a386fda18ac33de2edbce1fb0c55f26df0c952f032912b9aee"></a>

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

<a id="canonical-b0145fbb7e22a8a441097c5009a006d6188cb869ff4a447235e157435c2527aa"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 59539f7429a6 / 3

<a id="canonical-64c936c4b4e744de4a9fec1bf0e670efc9cc109d36cab6c05a8abf71e44909f0"></a>

<a id="canonical-2203969fb077b596a3d04dc9feaa9130c0a42d74984120b57908cbfcbb499632"></a>

## addr property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 59539f7429a6 / 4

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

<a id="canonical-bc5fd3f94824dfaac0133b9edf5ea8a3e1a653c23bdcf3df0c694afcd137c84a"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.ne / 59539f7429a6 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-005.md#canonical-8cf9ebe161d6097f5b63f4cc461a5b37feb992569343aec9aa5d28514f9eb3bb)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b55e9e6920ce9d2798e9fc486217d07e23a23c5eda496855ac45ab31fdbd77fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d476e162e110ef74d2cc8dabf916473b902cb5d2ccd121575782325c26e38d41"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 9b3a1bce162c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-a7b53d1a2d3ad264da9b61b07fd23cf029155580fa08fe268426b5c936c91a03"></a>

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

<a id="canonical-b64bf9ce7d6bfdd9c1301070f71fa9b027b937f4417da4f30e8b7b7f4fa07390"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 9b3a1bce162c / 3

- [ipv4](resources--azure_vnet_site--reference--group-005.md#canonical-54155cf4611a9c20421b5590ccbd874b1ef250f9a77fe547a049e9b88950ac6a): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-005.md#canonical-41d994e4b1eae1031549c2cfebe6001f3866f1bf98c19711ba43abb5efc24331): complete subsection reference.

<a id="canonical-6e5eee8c9330ad6ebd7492677e9fde8b0496ca51014eb4274a1633598fc9b4a7"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 9b3a1bce162c / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--azure_vnet_site--reference--group-005.md#canonical-54155cf4611a9c20421b5590ccbd874b1ef250f9a77fe547a049e9b88950ac6a)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--azure_vnet_site--reference--group-005.md#canonical-41d994e4b1eae1031549c2cfebe6001f3866f1bf98c19711ba43abb5efc24331)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-54155cf4611a9c20421b5590ccbd874b1ef250f9a77fe547a049e9b88950ac6a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d61b522dd5c2a24f415adfba52eb5e8171573178953d12703e472d1849fe0f7"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 6c3039d6f4d5 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-005.md#canonical-b55e9e6920ce9d2798e9fc486217d07e23a23c5eda496855ac45ab31fdbd77fe)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-702c46d980e3b64c8124e5cc07dfa0e2ac851883131d54cba95aadc1080425ca"></a>

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

<a id="canonical-4ca6c7e5159c856ddb5db05adaba903937fa34691e460231c05ca8336d2ff871"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 6c3039d6f4d5 / 3

<a id="canonical-421356eb127d21254abc06e2aa4b641d9d179777b5930d52816e00b1988a4262"></a>

<a id="canonical-9a3a8d12d86d4b8fe116202cff553c8e3681effef29d3e6efc58c0f7f78ccd6f"></a>

## plen property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 6c3039d6f4d5 / 4

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

<a id="canonical-ec45fc44d86a81465e9ee09a9d93ae61be98f25bd735e78698d16234ff0503ba"></a>

<a id="canonical-e22a0ca2bbdad1686a2704fc57f3072b35bc34291a890dfdc470838c642225b6"></a>

## prefix property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 6c3039d6f4d5 / 5

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

<a id="canonical-b82f6e17bfdfa0f4f77726b88ede79f8d3eea53a34966f44aafdae97bd996fb5"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 6c3039d6f4d5 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-005.md#canonical-b55e9e6920ce9d2798e9fc486217d07e23a23c5eda496855ac45ab31fdbd77fe)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-41d994e4b1eae1031549c2cfebe6001f3866f1bf98c19711ba43abb5efc24331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f32890210632f5e56fa3cf56f3d7c4776b79e2ef1a81ab1a8c12c0dacdd1169c"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 7b513f137903 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.outside_static_routes](resources--azure_vnet_site--reference--group-004.md#canonical-b2cf69d861e0a57ea1c50f0929ee2b476c5e3e8282d0f7a92aa6d14e7b6faa6a)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-004.md#canonical-18c00d1739a185869280757aedf5629b9dbe3f6360e3ae5c1315ab3868f8e1bb)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-004.md#canonical-1c64750fca39996634713907f2ecaa89ef1367766e0469bd3e57de47cc089f81)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-005.md#canonical-b55e9e6920ce9d2798e9fc486217d07e23a23c5eda496855ac45ab31fdbd77fe)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-bfc239715d281f68fea4ee5274d8905b79a78c6899f89310af673b13daaa9868"></a>

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

<a id="canonical-dfb44a837756830014b0b80180f02ad67abb175266e3f34e7e38dc7100882875"></a>

## Direct properties — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 7b513f137903 / 3

<a id="canonical-aea7babe3e1f1e07d6462395fcc015a459567a1d3f97055d314bbe525df404de"></a>

<a id="canonical-6cb3dbfd9cc94615b2e477270695c2c2dd7e9a6da050982132a44f2c158541cd"></a>

## plen property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 7b513f137903 / 4

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

<a id="canonical-402be3fdf2c4d8eb11deb2109aa0eaf0cba4f6d8e9d2daff7fdf26ab45655826"></a>

<a id="canonical-c725fa4a3c22387ad82bb7a11c1008d490ebb1ec555c46c2f95da644d6ca5046"></a>

## prefix property — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 7b513f137903 / 5

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

<a id="canonical-e6209cc70545346eb2e1f77a170165b5b884c26f38304b3bbd4ac8dbf54b11c7"></a>

## Next pages — ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.su / 7b513f137903 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-005.md#canonical-b55e9e6920ce9d2798e9fc486217d07e23a23c5eda496855ac45ab31fdbd77fe)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-73e62657c85cedab976af0728f61832b1d670020dc897977a03fbd5c0e1e2825"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51d2a581b0f0dffdffd676585d95c468392c3037817af8cb3afa77a835cd14b5"></a>

## ingress_egress_gw.performance_enhancement_mode — ingress_egress_gw.performance_enhancement_mode / 228ca652bb7c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.performance_enhancement_mode

<a id="canonical-ab61ef9a75bbde1873151218521eff4d1c0e3f13494f336975056ccd83efc5a2"></a>

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

<a id="canonical-bd6fc4bf211291d0de5c11cff1fdcd455228da9b00cd44600ed4a3c0ea666d62"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode / 228ca652bb7c / 3

- [perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-5e2307b718238b3ec6d848d559557275e9c60b32fd35a0dd1a80fb625053dfd5): complete subsection reference.

- [perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-a8dec9e4b47fdebe5621d023d9cadf522c773d3410c7622581c8f3614344e295): complete subsection reference.

<a id="canonical-56a4a88830931b58093fb355aae49b6a13b58a7939dfe8f7cbcfd6d0bf14dae4"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode / 228ca652bb7c / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-5e2307b718238b3ec6d848d559557275e9c60b32fd35a0dd1a80fb625053dfd5)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-a8dec9e4b47fdebe5621d023d9cadf522c773d3410c7622581c8f3614344e295)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-5e2307b718238b3ec6d848d559557275e9c60b32fd35a0dd1a80fb625053dfd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-40d068fee9f72a0adc1712c6b41fba2427c22c84a8b2a02778da280853695990"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 3dc03f70baba / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-005.md#canonical-73e62657c85cedab976af0728f61832b1d670020dc897977a03fbd5c0e1e2825)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-3207c2532c75dcb0f3753d053cdcc3b1331e52d03018803e3ff6282eaeda07b8"></a>

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

<a id="canonical-e4c435556663ad376b88e6ab7a1af906f36ad151ca9206bbd4a3a50cf09dbb28"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 3dc03f70baba / 3

- [jumbo](resources--azure_vnet_site--reference--group-005.md#canonical-a2f671404abf0d5f84c3fa7e3ce6a74d2120cce60aed6ecdd5c5573dbe913413): complete subsection reference.

- [no_jumbo](resources--azure_vnet_site--reference--group-005.md#canonical-60de3ef2a6d848a68c990f28ebfeb9c4b11c46d0ac9820116d3a5a01dec78072): complete subsection reference.

<a id="canonical-5b3b6cdf7fee777b4699a70dbabcfaf400c072edb74349b6df01a126ba04f200"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 3dc03f70baba / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--azure_vnet_site--reference--group-005.md#canonical-a2f671404abf0d5f84c3fa7e3ce6a74d2120cce60aed6ecdd5c5573dbe913413)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--azure_vnet_site--reference--group-005.md#canonical-60de3ef2a6d848a68c990f28ebfeb9c4b11c46d0ac9820116d3a5a01dec78072)
- [ingress_egress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-005.md#canonical-73e62657c85cedab976af0728f61832b1d670020dc897977a03fbd5c0e1e2825)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a2f671404abf0d5f84c3fa7e3ce6a74d2120cce60aed6ecdd5c5573dbe913413"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8a9a639f8a6ef88f306c7d2665a9986efe6f5fbcd652d34bd1116c510c76882"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 7806897b7285 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-005.md#canonical-73e62657c85cedab976af0728f61832b1d670020dc897977a03fbd5c0e1e2825)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-5e2307b718238b3ec6d848d559557275e9c60b32fd35a0dd1a80fb625053dfd5)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-05b3dba52820c0fb463359585077a1a57dbebb5f6860d12d3b54f2c215d10375"></a>

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

<a id="canonical-69e9141a3c69cf5f92276ac4540590432e5efa8fa3927ceffd9ca21bcdb75ea9"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 7806897b7285 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d07bbc81774bdc305aac769b430c50dfe4235d6f43670efde4bf620c1eab8c3e"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / 7806897b7285 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-5e2307b718238b3ec6d848d559557275e9c60b32fd35a0dd1a80fb625053dfd5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-60de3ef2a6d848a68c990f28ebfeb9c4b11c46d0ac9820116d3a5a01dec78072"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9858c5c619165d8bb7ffce9b7bd85873776006d885553358c4b0dc244616def2"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 2810290aac6c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-005.md#canonical-73e62657c85cedab976af0728f61832b1d670020dc897977a03fbd5c0e1e2825)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-5e2307b718238b3ec6d848d559557275e9c60b32fd35a0dd1a80fb625053dfd5)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-3846eea95f563038f9f05b3b3070c49331178274e2df8d5d2be5aad6505be661"></a>

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

<a id="canonical-e5826b5cefcf15e92e34e2c1925cbc7cbd15c210af52c9c3bebe349f8074dd47"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 2810290aac6c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-bbcd4f3c62c5d23d6d5e93194b8fbe09441a813488602b8dccf09d04a6ccaad9"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 2810290aac6c / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-5e2307b718238b3ec6d848d559557275e9c60b32fd35a0dd1a80fb625053dfd5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a8dec9e4b47fdebe5621d023d9cadf522c773d3410c7622581c8f3614344e295"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d913985770501d6b1db1175fdc139778a6c5a67777aaccaf9cfd8225a66d62c8"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 9007baf319dd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-005.md#canonical-73e62657c85cedab976af0728f61832b1d670020dc897977a03fbd5c0e1e2825)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-4a9f6a14197fa12ba9a6246cecfeb6851c15462018a8d7f57420605d34d5274f"></a>

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

<a id="canonical-34b15473404a57ffab8508826308877b8658c55d39f5410556ff61eb9f9d7cc8"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 9007baf319dd / 3

- [jumbo_disabled](resources--azure_vnet_site--reference--group-005.md#canonical-4310dcb5be8b2d8bd12634d3ef93409b14f80fade68b559fbea8190b8c29a2c0): complete subsection reference.

- [jumbo_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-a2136fb7ef242f4db1210a2046caadfe31687b0c6697871ca19477cb95d88c26): complete subsection reference.

<a id="canonical-c540b00e7bf3c7d840e5acc9206647a6ba2da3d0bc80a4c5bc15d392f0f66609"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 9007baf319dd / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--azure_vnet_site--reference--group-005.md#canonical-4310dcb5be8b2d8bd12634d3ef93409b14f80fade68b559fbea8190b8c29a2c0)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-a2136fb7ef242f4db1210a2046caadfe31687b0c6697871ca19477cb95d88c26)
- [ingress_egress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-005.md#canonical-73e62657c85cedab976af0728f61832b1d670020dc897977a03fbd5c0e1e2825)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-4310dcb5be8b2d8bd12634d3ef93409b14f80fade68b559fbea8190b8c29a2c0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2ff6f719a8445e1c85ca442b814707922973e394c8d73aaa0bf06db271dd2e79"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 25bc6a6bd34c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-005.md#canonical-73e62657c85cedab976af0728f61832b1d670020dc897977a03fbd5c0e1e2825)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-a8dec9e4b47fdebe5621d023d9cadf522c773d3410c7622581c8f3614344e295)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-88039d2619a7f049a87f07e366da0beafc57ea8abb993fac8bff30278b2a4b82"></a>

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

<a id="canonical-05bf341b265873c5f2643713cfa3f05c7e839b7325933fb0bd50f92d14e47711"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 25bc6a6bd34c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06b91fa74af2c43d3d39fe94265eb81aa303a76523a7eabeba8f59c6299caf56"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disab / 25bc6a6bd34c / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-a8dec9e4b47fdebe5621d023d9cadf522c773d3410c7622581c8f3614344e295)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a2136fb7ef242f4db1210a2046caadfe31687b0c6697871ca19477cb95d88c26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97645020e854252e919471b53f1b4c9ac221e3bf444233a94ce6fe829ef464e2"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / eaa9092d8e45 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [ingress_egress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-005.md#canonical-73e62657c85cedab976af0728f61832b1d670020dc897977a03fbd5c0e1e2825)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-a8dec9e4b47fdebe5621d023d9cadf522c773d3410c7622581c8f3614344e295)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-1a5430d54baf575a7112fe2a1e93f91ccf86f44d80eb7627506066bb6ad17d4b"></a>

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

<a id="canonical-fd7fa69f098c53522f2a710131f799e281de779f27bb315aace40c2a64e30004"></a>

## Direct properties — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / eaa9092d8e45 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f0032f8b119b7c0940d77de1fac863c5f8ef8120e18dc51b1ea7a7272d9131f7"></a>

## Next pages — ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabl / eaa9092d8e45 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-005.md#canonical-a8dec9e4b47fdebe5621d023d9cadf522c773d3410c7622581c8f3614344e295)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6b1fc893ac3746b9b3f218800395e2d6fd2743542e41bc0e5cee3adbdb1e97f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19fb3d7e6ce9077d8bf42ba5a44a8d946036eef8527a2f5837c134ebce66874a"></a>

## ingress_egress_gw.sm_connection_public_ip — ingress_egress_gw.sm_connection_public_ip / 52102717aedb / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.sm_connection_public_ip

<a id="canonical-bf86008a4cd4ffdd393741a290db36c7e3ce8dead93cb2070ff91015435378ad"></a>

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

<a id="canonical-d554ede899c2501d93265d82bac7010b322d18cbaf244586019a79ec27a52436"></a>

## Direct properties — ingress_egress_gw.sm_connection_public_ip / 52102717aedb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-af5aaf395b9039252c0aac2be5d53647ef93a72b3b3b2537c7e8966503e579b0"></a>

## Next pages — ingress_egress_gw.sm_connection_public_ip / 52102717aedb / 4

- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c746fabef3426a0824d12667f0394ab0c285816502fb3969836db7edb0112d8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-67c972d385b730aa887c9f15f666921f0e7db370c810e59db49ffa984a471f22"></a>

## ingress_egress_gw.sm_connection_pvt_ip — ingress_egress_gw.sm_connection_pvt_ip / 6de097755396 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- ingress_egress_gw.sm_connection_pvt_ip

<a id="canonical-20e31af37d52ba53e38890e4e09445aea379e4e8c60570b790fd26726a9fe4d7"></a>

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

<a id="canonical-4e2e83319895b55fa9a886e49cf16c85775e8779f96e8c85810ee24eeca83b88"></a>

## Direct properties — ingress_egress_gw.sm_connection_pvt_ip / 6de097755396 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d5551e50fe455a1acb4418515dcdae78efe93a69ffc8c7249fb624f4547346ed"></a>

## Next pages — ingress_egress_gw.sm_connection_pvt_ip / 6de097755396 / 4

- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-84713d0ad72950b74f8b232b5180e4164740ac650342c1ffd50d0fea49bd7218)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3969001f093abf05cbabb08659e97f071dc9882ad242e78e6cee225885555b1a"></a>

## ingress_egress_gw_ar — ingress_egress_gw_ar / 0322983b9e94 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- ingress_egress_gw_ar

<a id="canonical-97a2a2a6fcc0f857b225d133e342bf89cbf0ee35aaabfe9b34059df31a3e4e7f"></a>

Type: `"object"`. single nested block, Optional.

Two interface Azure ingress/egress site on Alternate Region with no support for zones.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("azure_certified_hw"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "dc_cluster_group_outside_vn"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("dc_cluster_group_outside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("hub",
    "not_hub"),
  validators.ConflictingObjectAttributes("inside_static_routes",
    "no_inside_static_routes"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-hub_choice": "[\"hub\",\"not_hub\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

Terraform syntax:

```terraform
ingress_egress_gw_ar {
  # Configure direct properties listed below.
}
```

<a id="canonical-95e7054408dac8286605d6cd449719fd96859be26cfb2ac5bcf3bf3ebffee93d"></a>

## Direct properties — ingress_egress_gw_ar / 0322983b9e94 / 3

- [accelerated_networking](resources--azure_vnet_site--reference--group-005.md#canonical-0cd2d55173887d4ea2dfc00f0ad247f33c85d0c2d9a0a804896699c30ae5f2de): complete subsection reference.

- [active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-005.md#canonical-253ecd420771967e2325f2375fcd04f6475962311b9755dd40c3829c911a50e3): complete subsection reference.

- [active_forward_proxy_policies](resources--azure_vnet_site--reference--group-005.md#canonical-d67a52e48fbf03e9b059dbe2e70a2c37ffac4baecf84542787f78d5101ef6cb3): complete subsection reference.

- [active_network_policies](resources--azure_vnet_site--reference--group-005.md#canonical-14bd033854c92387562900516739819373f8b4362975cf6bf52c07e5b6effbce): complete subsection reference.

<a id="canonical-70d199cdd31ecf855367b2344ad64957716d846d5d4451637273d53c25554ea6"></a>

<a id="canonical-d1e9addd0f4d22395f369235c7509ac2b05e62d495049e7d33dda7ebfdfdcf89"></a>

## azure_certified_hw property — ingress_egress_gw_ar / 0322983b9e94 / 4

Type: `"string"`. Optional.

\[Enum: azure-byol-multi-nic-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware.
The only possible value is \`azure-byol-multi-nic-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("azure-byol-multi-nic-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-multi-nic-voltmesh"
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
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [dc_cluster_group_inside_vn](resources--azure_vnet_site--reference--group-005.md#canonical-e01dde0317a6ecb8d9852c030507d2a582562a879a2a5f850f80abee4c77b56b): complete subsection reference.

- [dc_cluster_group_outside_vn](resources--azure_vnet_site--reference--group-005.md#canonical-947e58b1abafb74a7dd76560b1f9eac886d7fd60144a615f9cd89a5dd5dee0eb): complete subsection reference.

- [forward_proxy_allow_all](resources--azure_vnet_site--reference--group-005.md#canonical-6002ce0d2e665d9b01569e25fb7eed8f9345a3ad0da597685e92c364243839ce): complete subsection reference.

- [global_network_list](resources--azure_vnet_site--reference--group-005.md#canonical-4ed19e66e966b9eda39ebfb2066578c56f60cf14018d888b308272bd5c7593d3): complete subsection reference.

- [hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86): complete subsection reference.

- [inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2): complete subsection reference.

- [no_dc_cluster_group](resources--azure_vnet_site--reference--group-006.md#canonical-818b2f34104f14d32057cca2e534182d9ac2d33a53df4aa9b717a719ada0eb26): complete subsection reference.

- [no_forward_proxy](resources--azure_vnet_site--reference--group-006.md#canonical-9bb715e41a041df39664d0319b1bee1c7e20c1a98800d02496ff67fc172ca6e8): complete subsection reference.

- [no_global_network](resources--azure_vnet_site--reference--group-006.md#canonical-f51b02e48409cf9538906b61dbe13c3a9350e7fd69c4b25e12302657fab47215): complete subsection reference.

- [no_inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-02d31b1e638725cd89df46178569cc3cb222f88646f90faf10df5185e21ecd9e): complete subsection reference.

- [no_network_policy](resources--azure_vnet_site--reference--group-006.md#canonical-86a905690a308b5667688711ac6d7232a89f386b2cb89604be9a4678fdd5b4d4): complete subsection reference.

- [no_outside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-edbfe8c94180d225a7f8eb569f8efc122ae23ded1317587e2dde7c7063d4c534): complete subsection reference.

- [node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe): complete subsection reference.

- [not_hub](resources--azure_vnet_site--reference--group-007.md#canonical-3e7b6b00c9afdd10dd5ed7a994023188b368fadf3ae535ceaf69484151a058fe): complete subsection reference.

- [outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a): complete subsection reference.

- [performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-8c1fe208b21819c9e2a17ea2396879bed1b8261c750b75f44a9a953ef70031cb): complete subsection reference.

- [sm_connection_public_ip](resources--azure_vnet_site--reference--group-007.md#canonical-42a38e5c411a6ae53043b919a491382350a63697848e7b49f4113ae2a01abcce): complete subsection reference.

- [sm_connection_pvt_ip](resources--azure_vnet_site--reference--group-007.md#canonical-33bb71f61320605de6c5a00f1a99a858c225d67c17dc22d19f49e655eb832e8d): complete subsection reference.

<a id="canonical-774d50bee0c09de9b4fbc12f3d556d91b7d6830b0d88304ca0b1996e2cec2a81"></a>

## Next pages — ingress_egress_gw_ar / 0322983b9e94 / 5

- [ingress_egress_gw_ar.accelerated_networking](resources--azure_vnet_site--reference--group-005.md#canonical-0cd2d55173887d4ea2dfc00f0ad247f33c85d0c2d9a0a804896699c30ae5f2de)
- [ingress_egress_gw_ar.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-005.md#canonical-253ecd420771967e2325f2375fcd04f6475962311b9755dd40c3829c911a50e3)
- [ingress_egress_gw_ar.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-005.md#canonical-d67a52e48fbf03e9b059dbe2e70a2c37ffac4baecf84542787f78d5101ef6cb3)
- [ingress_egress_gw_ar.active_network_policies](resources--azure_vnet_site--reference--group-005.md#canonical-14bd033854c92387562900516739819373f8b4362975cf6bf52c07e5b6effbce)
- [ingress_egress_gw_ar.dc_cluster_group_inside_vn](resources--azure_vnet_site--reference--group-005.md#canonical-e01dde0317a6ecb8d9852c030507d2a582562a879a2a5f850f80abee4c77b56b)
- [ingress_egress_gw_ar.dc_cluster_group_outside_vn](resources--azure_vnet_site--reference--group-005.md#canonical-947e58b1abafb74a7dd76560b1f9eac886d7fd60144a615f9cd89a5dd5dee0eb)
- [ingress_egress_gw_ar.forward_proxy_allow_all](resources--azure_vnet_site--reference--group-005.md#canonical-6002ce0d2e665d9b01569e25fb7eed8f9345a3ad0da597685e92c364243839ce)
- [ingress_egress_gw_ar.global_network_list](resources--azure_vnet_site--reference--group-005.md#canonical-4ed19e66e966b9eda39ebfb2066578c56f60cf14018d888b308272bd5c7593d3)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-83daff971568976e1322aa005d7edac6fc071a1d6ad075cc47c167833c9f00a2)
- [ingress_egress_gw_ar.no_dc_cluster_group](resources--azure_vnet_site--reference--group-006.md#canonical-818b2f34104f14d32057cca2e534182d9ac2d33a53df4aa9b717a719ada0eb26)
- [ingress_egress_gw_ar.no_forward_proxy](resources--azure_vnet_site--reference--group-006.md#canonical-9bb715e41a041df39664d0319b1bee1c7e20c1a98800d02496ff67fc172ca6e8)
- [ingress_egress_gw_ar.no_global_network](resources--azure_vnet_site--reference--group-006.md#canonical-f51b02e48409cf9538906b61dbe13c3a9350e7fd69c4b25e12302657fab47215)
- [ingress_egress_gw_ar.no_inside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-02d31b1e638725cd89df46178569cc3cb222f88646f90faf10df5185e21ecd9e)
- [ingress_egress_gw_ar.no_network_policy](resources--azure_vnet_site--reference--group-006.md#canonical-86a905690a308b5667688711ac6d7232a89f386b2cb89604be9a4678fdd5b4d4)
- [ingress_egress_gw_ar.no_outside_static_routes](resources--azure_vnet_site--reference--group-006.md#canonical-edbfe8c94180d225a7f8eb569f8efc122ae23ded1317587e2dde7c7063d4c534)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe)
- [ingress_egress_gw_ar.not_hub](resources--azure_vnet_site--reference--group-007.md#canonical-3e7b6b00c9afdd10dd5ed7a994023188b368fadf3ae535ceaf69484151a058fe)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-8c1fe208b21819c9e2a17ea2396879bed1b8261c750b75f44a9a953ef70031cb)
- [ingress_egress_gw_ar.sm_connection_public_ip](resources--azure_vnet_site--reference--group-007.md#canonical-42a38e5c411a6ae53043b919a491382350a63697848e7b49f4113ae2a01abcce)
- [ingress_egress_gw_ar.sm_connection_pvt_ip](resources--azure_vnet_site--reference--group-007.md#canonical-33bb71f61320605de6c5a00f1a99a858c225d67c17dc22d19f49e655eb832e8d)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-0cd2d55173887d4ea2dfc00f0ad247f33c85d0c2d9a0a804896699c30ae5f2de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e0cd63e0590c4059ffcb29bcbeed6b35d8f94cc1d31c55e6fa8352f875195363"></a>

## ingress_egress_gw_ar.accelerated_networking — ingress_egress_gw_ar.accelerated_networking / d87aee2b0f66 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.accelerated_networking

<a id="canonical-fc16702eab049ec02334ca9894f9bac4790950a1cb9a8ec7f211db5048f9bda4"></a>

Type: `"object"`. single nested block, Optional.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Upstream description:

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-accelerated_networking": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
accelerated_networking {
  # Configure direct properties listed below.
}
```

<a id="canonical-ee057381cb228371451407c064c49245511af940def9c3739c7a9b6e60b5f77d"></a>

## Direct properties — ingress_egress_gw_ar.accelerated_networking / d87aee2b0f66 / 3

- [disable_spec](resources--azure_vnet_site--reference--group-005.md#canonical-b72ae55af43b0e1df0a354f081380a9a22784d4382c562602de8959fb4be2967): complete subsection reference.

- [enable](resources--azure_vnet_site--reference--group-005.md#canonical-ac117eb3447effde8627fcb54ed2b26bcb00128f236b497531f6d5fbc9af5a97): complete subsection reference.

<a id="canonical-379b2f8da012e1c82992994651212f26e330bf6d8f5bfeb6d3bca0cd0d0bae7f"></a>

## Next pages — ingress_egress_gw_ar.accelerated_networking / d87aee2b0f66 / 4

- [ingress_egress_gw_ar.accelerated_networking.disable_spec](resources--azure_vnet_site--reference--group-005.md#canonical-b72ae55af43b0e1df0a354f081380a9a22784d4382c562602de8959fb4be2967)
- [ingress_egress_gw_ar.accelerated_networking.enable](resources--azure_vnet_site--reference--group-005.md#canonical-ac117eb3447effde8627fcb54ed2b26bcb00128f236b497531f6d5fbc9af5a97)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b72ae55af43b0e1df0a354f081380a9a22784d4382c562602de8959fb4be2967"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c30ae9c4bea736391adcf07046cd50af3d3adffe370bba91583d947342f37cd7"></a>

## ingress_egress_gw_ar.accelerated_networking.disable_spec — ingress_egress_gw_ar.accelerated_networking.disable_spec / fc485476ee75 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.accelerated_networking](resources--azure_vnet_site--reference--group-005.md#canonical-0cd2d55173887d4ea2dfc00f0ad247f33c85d0c2d9a0a804896699c30ae5f2de)
- ingress_egress_gw_ar.accelerated_networking.disable_spec

<a id="canonical-4573e7894a2986fcb0a684e3c9ce224003450ae2de52156f4ec2411564e08ccc"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-7a3b045ef0096ff92cdf95b6a48e53c9d0b4fe1f0af2f18e5ebfc04f4c448cdd"></a>

## Direct properties — ingress_egress_gw_ar.accelerated_networking.disable_spec / fc485476ee75 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c195c511fe1cd26713f12157aadd49b9ec3255a3df3192bd269da99e919705e8"></a>

## Next pages — ingress_egress_gw_ar.accelerated_networking.disable_spec / fc485476ee75 / 4

- [ingress_egress_gw_ar.accelerated_networking](resources--azure_vnet_site--reference--group-005.md#canonical-0cd2d55173887d4ea2dfc00f0ad247f33c85d0c2d9a0a804896699c30ae5f2de)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ac117eb3447effde8627fcb54ed2b26bcb00128f236b497531f6d5fbc9af5a97"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1801c698ac758b79dcca7fa71078a4ed4f13f542d640d75301c6df2e5f2066e6"></a>

## ingress_egress_gw_ar.accelerated_networking.enable — ingress_egress_gw_ar.accelerated_networking.enable / c4b835d2075f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.accelerated_networking](resources--azure_vnet_site--reference--group-005.md#canonical-0cd2d55173887d4ea2dfc00f0ad247f33c85d0c2d9a0a804896699c30ae5f2de)
- ingress_egress_gw_ar.accelerated_networking.enable

<a id="canonical-826349f8bf84693dc8ca4d27096cbe5701a60141d6763d88a8556c70638a2d2f"></a>

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
enable = {}
```

<a id="canonical-cbb89a899942037c26c6a103f02367bebfb0f88dd8a6dfd059f013a3d451a218"></a>

## Direct properties — ingress_egress_gw_ar.accelerated_networking.enable / c4b835d2075f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5a7888c3f7669f5fa125cdb5e74da562f87df149cca26bef41abe6fbb103f24a"></a>

## Next pages — ingress_egress_gw_ar.accelerated_networking.enable / c4b835d2075f / 4

- [ingress_egress_gw_ar.accelerated_networking](resources--azure_vnet_site--reference--group-005.md#canonical-0cd2d55173887d4ea2dfc00f0ad247f33c85d0c2d9a0a804896699c30ae5f2de)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-253ecd420771967e2325f2375fcd04f6475962311b9755dd40c3829c911a50e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-69c9bb4d6f4b26f4191c7252f604cf08682ec63c6c977635e5934dd252b5e4fe"></a>

## ingress_egress_gw_ar.active_enhanced_firewall_policies — ingress_egress_gw_ar.active_enhanced_firewall_policies / b619e85e9a1c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.active_enhanced_firewall_policies

<a id="canonical-d1b307736ab0d1773cf9894def75713bcd2fb3a22565e33ab96342f3a5db58fb"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-f98d03f677fa959fb39900f8cdb75cc277050c6ada712e872a6f889f868ebb55"></a>

## Direct properties — ingress_egress_gw_ar.active_enhanced_firewall_policies / b619e85e9a1c / 3

- [enhanced_firewall_policies](resources--azure_vnet_site--reference--group-005.md#canonical-0f78d78b7f144db73ebec3230c4fc348dd01580ce3a47ae3d8140be5c0a0a1f0): complete subsection reference.

<a id="canonical-471c05340fcc304d1e784c5efb8f44036c5fa4f5dd52cbdf89b25d37ea3c2b9f"></a>

## Next pages — ingress_egress_gw_ar.active_enhanced_firewall_policies / b619e85e9a1c / 4

- [ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--azure_vnet_site--reference--group-005.md#canonical-0f78d78b7f144db73ebec3230c4fc348dd01580ce3a47ae3d8140be5c0a0a1f0)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-0f78d78b7f144db73ebec3230c4fc348dd01580ce3a47ae3d8140be5c0a0a1f0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aeb1a952263d2ca863bca856fe5f1eb05dd74898eebc52f01c6f688ba90057cb"></a>

## ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / fc1ccba5ac39 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-005.md#canonical-253ecd420771967e2325f2375fcd04f6475962311b9755dd40c3829c911a50e3)
- ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-e018aa26ad50cee8e7648e09e444926976df8c00fae2127cf084c750c5954ee7"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-f30e38e1298212c8492a336cec4163ea285f17430e7bc961ab96824b848f2a56"></a>

## Direct properties — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / fc1ccba5ac39 / 3

<a id="canonical-9f41cf8321c39c5fc1f28d89a75c66fd957d0d68f613eeb8f19db2ca65da8122"></a>

<a id="canonical-0b58391c75b38060f954fc56a849f13ed0a927261d4d861320edd035b830500f"></a>

## name property — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / fc1ccba5ac39 / 4

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

<a id="canonical-0c6fcafe2123ef9880e9e49363b322d9e4f35da072c0a31cea182c0087a17cf1"></a>

<a id="canonical-09d4a57d7050ed44e647b8f35fd36bff0e52cba9ab5afe7d3c74bb22b3172e9e"></a>

## namespace property — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / fc1ccba5ac39 / 5

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

<a id="canonical-6b175b4b00ccb6e12546f4bcd00f1b8f40e1e4bd3683e620cd7fa07a77a6d14f"></a>

<a id="canonical-500392330828d3a7de0c18f88a74a2eaccf7b51e017c3bc836d1a66709ed9766"></a>

## tenant property — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / fc1ccba5ac39 / 6

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

<a id="canonical-5ab9524f41b2b73f505530eae98bd394c131dd80bc3f57434f71dec9e79d204a"></a>

## Next pages — ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / fc1ccba5ac39 / 7

- [ingress_egress_gw_ar.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-005.md#canonical-253ecd420771967e2325f2375fcd04f6475962311b9755dd40c3829c911a50e3)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d67a52e48fbf03e9b059dbe2e70a2c37ffac4baecf84542787f78d5101ef6cb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53400b806a8eeaa762067d069a903a3a41a216a36c8cbbc829f4d1022302f885"></a>

## ingress_egress_gw_ar.active_forward_proxy_policies — ingress_egress_gw_ar.active_forward_proxy_policies / 3e9d94537df2 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.active_forward_proxy_policies

<a id="canonical-6293564a80caa41960ed29e52b490bed082309f95ac817b0101972ece15e9f79"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0ac1817b184df17a22f115158c6405cacb90fe5805e3316a36a54f8cfd4c5af3"></a>

## Direct properties — ingress_egress_gw_ar.active_forward_proxy_policies / 3e9d94537df2 / 3

- [forward_proxy_policies](resources--azure_vnet_site--reference--group-005.md#canonical-6abd924883cbf2862aefd042e1c430bc5cdccf2ce06194377c404a33d61e7af1): complete subsection reference.

<a id="canonical-14ae0310511c047c5bcf1338c4a9ed37d98a85f9490feea232cc2fa866518307"></a>

## Next pages — ingress_egress_gw_ar.active_forward_proxy_policies / 3e9d94537df2 / 4

- [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies](resources--azure_vnet_site--reference--group-005.md#canonical-6abd924883cbf2862aefd042e1c430bc5cdccf2ce06194377c404a33d61e7af1)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6abd924883cbf2862aefd042e1c430bc5cdccf2ce06194377c404a33d61e7af1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed7c39f0d2beace8c0c97096ce8e81f338384789fb01dd76030550fbbd6ce57b"></a>

## ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / 830c780c2974 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-005.md#canonical-d67a52e48fbf03e9b059dbe2e70a2c37ffac4baecf84542787f78d5101ef6cb3)
- ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-5d83b7d87186b7ad9a404060d2ab95c48136c50cef3966fd92a23f4a13b8f638"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-4a0ece0db3f05b4a432e5fac9b75d518716e85dc971f2c4865fb5d881c8590db"></a>

## Direct properties — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / 830c780c2974 / 3

<a id="canonical-110d1c2b507a57865a0bd058c7d4521cb0ac9db0dd28bda9d76518efcec2b897"></a>

<a id="canonical-57fb9a98d752b15fcbb3c2a9e159799097c1f996896a34c774ab30410d8d3b40"></a>

## name property — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / 830c780c2974 / 4

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

<a id="canonical-5c875da8b841939c28c76270477fc4c5ae50dc77012e9cea730b9a1bd5452e48"></a>

<a id="canonical-62c53ee2f001fec22770aae8a358c0d14c99c0122fb53fddf80867a58967bab7"></a>

## namespace property — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / 830c780c2974 / 5

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

<a id="canonical-98a1bd60f7b996f44cc1017cbc045d9387d57ba76421bcea01845f03a7f8f827"></a>

<a id="canonical-07802dc316f3dfef18c856e632bbb5694894103547cfdeb962cc27af48defce7"></a>

## tenant property — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / 830c780c2974 / 6

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

<a id="canonical-2182ceec959810e20e1c493cbee71f163892099f2b3e46a0701dff5854d50099"></a>

## Next pages — ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies / 830c780c2974 / 7

- [ingress_egress_gw_ar.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-005.md#canonical-d67a52e48fbf03e9b059dbe2e70a2c37ffac4baecf84542787f78d5101ef6cb3)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-14bd033854c92387562900516739819373f8b4362975cf6bf52c07e5b6effbce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dcec029f45e557d678c73bb106cfd025dbe6b722b1a2b958682b767c9cf8408d"></a>

## ingress_egress_gw_ar.active_network_policies — ingress_egress_gw_ar.active_network_policies / 827e3aacdfee / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.active_network_policies

<a id="canonical-ecf62de667e422bf1c5059fbb3b5c21d954bf41da061251789ea9bbf9f8a6adb"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-42665283368a00d653cecfab33404a9b55d8f931e419e44264caaac23f5d9e2e"></a>

## Direct properties — ingress_egress_gw_ar.active_network_policies / 827e3aacdfee / 3

- [network_policies](resources--azure_vnet_site--reference--group-005.md#canonical-cdae1ddc85b0562e117ebb899cd96e6fa4396c75a5b67d2f0c654c433476cd3c): complete subsection reference.

<a id="canonical-93f9184a4066660937a820d6603d1e0de229474b06ff5317ee3ce9a57dc11742"></a>

## Next pages — ingress_egress_gw_ar.active_network_policies / 827e3aacdfee / 4

- [ingress_egress_gw_ar.active_network_policies.network_policies](resources--azure_vnet_site--reference--group-005.md#canonical-cdae1ddc85b0562e117ebb899cd96e6fa4396c75a5b67d2f0c654c433476cd3c)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-cdae1ddc85b0562e117ebb899cd96e6fa4396c75a5b67d2f0c654c433476cd3c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b218117af1ea8158c5dbf02951c99649211ce8f255cd05ebb15e83928b8720f"></a>

## ingress_egress_gw_ar.active_network_policies.network_policies — ingress_egress_gw_ar.active_network_policies.network_policies / 639db47544b9 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.active_network_policies](resources--azure_vnet_site--reference--group-005.md#canonical-14bd033854c92387562900516739819373f8b4362975cf6bf52c07e5b6effbce)
- ingress_egress_gw_ar.active_network_policies.network_policies

<a id="canonical-a7e9cd129334a9870a2f8ab2a69246bbc1c6cf75b364447c4f9da38228452b2e"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-b5f4e6e32fbdd5b955ba449d258d39f9df0e62dc7014ae62aa18ea8f95be046b"></a>

## Direct properties — ingress_egress_gw_ar.active_network_policies.network_policies / 639db47544b9 / 3

<a id="canonical-9d828ef6f699e3e7dda410c8191a4a42ab441f5735d0094631bc7d4eca7c8dee"></a>

<a id="canonical-9567c9e35346283560c03c7fce29c92440bcd748fdbedcbb301797e296cfb30d"></a>

## name property — ingress_egress_gw_ar.active_network_policies.network_policies / 639db47544b9 / 4

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

<a id="canonical-b89b01f5a452bf0171c34369566979b177eeaa5df987118a1c0310bfb96f4ff0"></a>

<a id="canonical-a4505707e9947f34c8a371570ea7c2af8a851afa1e2ac593dc4dc54b6104d091"></a>

## namespace property — ingress_egress_gw_ar.active_network_policies.network_policies / 639db47544b9 / 5

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

<a id="canonical-b1b76f92b08ee7f772d6ddc53454d356d5e50b4ceb5d09242ce912ea94dcb024"></a>

<a id="canonical-08d372431c98fdf5679c2c7c58f5274b39b8425b7a5ad354f7443bdbca45ae6f"></a>

## tenant property — ingress_egress_gw_ar.active_network_policies.network_policies / 639db47544b9 / 6

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

<a id="canonical-367bcca8c82d03d1f3ece6632f3d96ad3ed8c5e51b7e5d80b5fc37516746a68f"></a>

## Next pages — ingress_egress_gw_ar.active_network_policies.network_policies / 639db47544b9 / 7

- [ingress_egress_gw_ar.active_network_policies](resources--azure_vnet_site--reference--group-005.md#canonical-14bd033854c92387562900516739819373f8b4362975cf6bf52c07e5b6effbce)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e01dde0317a6ecb8d9852c030507d2a582562a879a2a5f850f80abee4c77b56b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1c333ee0a2a29f50d1827766ac4eb19202a9fc2ac6740c5ac33667c12da2660c"></a>

## ingress_egress_gw_ar.dc_cluster_group_inside_vn — ingress_egress_gw_ar.dc_cluster_group_inside_vn / ceb9dcc7105e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.dc_cluster_group_inside_vn

<a id="canonical-ccc01b75b35c47b5f2868411cdd7ab0ce4c80ac3f87660fd63ff7a3f2d8dd775"></a>

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
dc_cluster_group_inside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-a28532a295f7fc5700e7fcffc5515548bda057bd3057b5b45379ad1a2b9c8263"></a>

## Direct properties — ingress_egress_gw_ar.dc_cluster_group_inside_vn / ceb9dcc7105e / 3

<a id="canonical-9c8bb2586543eed650719a43ca00fca6f1f03882eba12ea536d8b30aad43c983"></a>

<a id="canonical-a10ca48dc22a0f7336e888bc16e2f7bdb25af0395923e4a239975cffc4900c69"></a>

## name property — ingress_egress_gw_ar.dc_cluster_group_inside_vn / ceb9dcc7105e / 4

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

<a id="canonical-fdc2e93f346650c4c758f20912a39466ab62037a50d33930849bda547e564733"></a>

<a id="canonical-c594749b921acdee36590cba6121b3e3e67654ad1064ad7d5024afa929eb2008"></a>

## namespace property — ingress_egress_gw_ar.dc_cluster_group_inside_vn / ceb9dcc7105e / 5

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

<a id="canonical-422a654641828993efb3ea9a35018fd624ab4712a33a301fc268934cce69950e"></a>

<a id="canonical-730be2458c39b8ccef8815dbaef5db39f5e639f89382ba5ad1882b8876980899"></a>

## tenant property — ingress_egress_gw_ar.dc_cluster_group_inside_vn / ceb9dcc7105e / 6

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

<a id="canonical-58063230d2effbac323c7e377858ae02d153d4725c18c5e0d65b2400bb4ef23a"></a>

## Next pages — ingress_egress_gw_ar.dc_cluster_group_inside_vn / ceb9dcc7105e / 7

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-947e58b1abafb74a7dd76560b1f9eac886d7fd60144a615f9cd89a5dd5dee0eb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7ce7e98c8a8b7f6b78b0bcab54284ba55e43afe8a39452028521cad67b7c3bce"></a>

## ingress_egress_gw_ar.dc_cluster_group_outside_vn — ingress_egress_gw_ar.dc_cluster_group_outside_vn / b19cc0332982 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.dc_cluster_group_outside_vn

<a id="canonical-a920296ee4b658c79c7bf7773a5594835e756cad5729bb15fc5589517f74b6db"></a>

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
dc_cluster_group_outside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-ef49a4c398e4406227b8fca343d14fde2fc80e83f7c97646d9558794fd10beac"></a>

## Direct properties — ingress_egress_gw_ar.dc_cluster_group_outside_vn / b19cc0332982 / 3

<a id="canonical-9c3443f06c6772d5228c4b14b45c1404cc8c144ccc9eefe8e1de7f444c307896"></a>

<a id="canonical-0a4219bf198415de28269365dbb5d551681904c941f204588f7715d6b64ee903"></a>

## name property — ingress_egress_gw_ar.dc_cluster_group_outside_vn / b19cc0332982 / 4

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

<a id="canonical-cc245ee02a467c34856baaf925291f516017aa3e69d6193e6af49c7a2acdd794"></a>

<a id="canonical-6275d3ffd97abe2583330602697b75754d580f1adadc6e7679a040d80e5c9086"></a>

## namespace property — ingress_egress_gw_ar.dc_cluster_group_outside_vn / b19cc0332982 / 5

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

<a id="canonical-a7caecce03a864376af5f42249b5554104013c64cc5a03d2aa2eb3ca6157449c"></a>

<a id="canonical-e978a2f258450dbb6751a8b939b9fc7572cc38544bdc8db1b46a188fb470336a"></a>

## tenant property — ingress_egress_gw_ar.dc_cluster_group_outside_vn / b19cc0332982 / 6

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

<a id="canonical-b26dd1e1970fab578eeef84a3096140ac6354d5147964922999dfa72e5d1ab7a"></a>

## Next pages — ingress_egress_gw_ar.dc_cluster_group_outside_vn / b19cc0332982 / 7

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6002ce0d2e665d9b01569e25fb7eed8f9345a3ad0da597685e92c364243839ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9cc6935434522d5af02aa662037b874dbdfa1ecca1fe8d8d82dfdaa527275dd"></a>

## ingress_egress_gw_ar.forward_proxy_allow_all — ingress_egress_gw_ar.forward_proxy_allow_all / a5816ca2a7f2 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.forward_proxy_allow_all

<a id="canonical-1336f291fee85e6f8958766e93a435c1f31c0b144bb78f7bf86d83f80a44a9eb"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

<a id="canonical-ed99dd6003cddbd7a6a0a9832172b084ca61dd261c6e66469d58ef4446069a62"></a>

## Direct properties — ingress_egress_gw_ar.forward_proxy_allow_all / a5816ca2a7f2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-dfa8f6d8863ba78a691659542ba49d4a2e6d94e591bfab4a36270eb2ddb5fa1a"></a>

## Next pages — ingress_egress_gw_ar.forward_proxy_allow_all / a5816ca2a7f2 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-4ed19e66e966b9eda39ebfb2066578c56f60cf14018d888b308272bd5c7593d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed6dfdb9cf0a2c39ab2c961c9aeac3e151489d4be16652753b49dd92c67f496b"></a>

## ingress_egress_gw_ar.global_network_list — ingress_egress_gw_ar.global_network_list / f663350e03be / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.global_network_list

<a id="canonical-a0f25929ba2a2bef4ed423538ea32e906d54ef4384f513615ff5346c7e4675b3"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-4a347b064571aa1c98285e46f26f27030435650d29106569ecfd736d7e43574d"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list / f663350e03be / 3

- [global_network_connections](resources--azure_vnet_site--reference--group-005.md#canonical-8badaeb09cc425e0bb1045224b8aeeb2b890e5ea5cc125ea47ae0b63004adf4f): complete subsection reference.

<a id="canonical-61ba2b8873b56fef7617c48a94f84034aa80838ddc652270527b9bfa1b2c3a11"></a>

## Next pages — ingress_egress_gw_ar.global_network_list / f663350e03be / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-005.md#canonical-8badaeb09cc425e0bb1045224b8aeeb2b890e5ea5cc125ea47ae0b63004adf4f)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-8badaeb09cc425e0bb1045224b8aeeb2b890e5ea5cc125ea47ae0b63004adf4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-22dbe0228739f41beabb78b4377e0831075e27f1138b800680167f3767683246"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections — ingress_egress_gw_ar.global_network_list.global_network_connections / 47fe013ad3b1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.global_network_list](resources--azure_vnet_site--reference--group-005.md#canonical-4ed19e66e966b9eda39ebfb2066578c56f60cf14018d888b308272bd5c7593d3)
- ingress_egress_gw_ar.global_network_list.global_network_connections

<a id="canonical-c419c4c20680a0df88768171d1594d769ac26b2e1834fe46204aa8eb95e8125c"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-3c093ed064b440d5aef92bd740b79e02005ca55ee33cdc9461c8e2e5351c0b62"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list.global_network_connections / 47fe013ad3b1 / 3

- [sli_to_global_dr](resources--azure_vnet_site--reference--group-005.md#canonical-d9f78026ba8964fdb9d0ec325120d98dec2ede52a20b385d8cc2b9f772496459): complete subsection reference.

- [slo_to_global_dr](resources--azure_vnet_site--reference--group-005.md#canonical-d20ccdd8ac497cf011534fba7bfb1a1e22749db301720c02ab5c8e23dcc04bb2): complete subsection reference.

<a id="canonical-8f6575f963dad75176b21664e2d06eb12ceaab68b24c985fc88034b56455558d"></a>

## Next pages — ingress_egress_gw_ar.global_network_list.global_network_connections / 47fe013ad3b1 / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-005.md#canonical-d9f78026ba8964fdb9d0ec325120d98dec2ede52a20b385d8cc2b9f772496459)
- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-005.md#canonical-d20ccdd8ac497cf011534fba7bfb1a1e22749db301720c02ab5c8e23dcc04bb2)
- [ingress_egress_gw_ar.global_network_list](resources--azure_vnet_site--reference--group-005.md#canonical-4ed19e66e966b9eda39ebfb2066578c56f60cf14018d888b308272bd5c7593d3)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d9f78026ba8964fdb9d0ec325120d98dec2ede52a20b385d8cc2b9f772496459"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07d6ea9b98205e52cbea49d41cca73e57946f833534c075c132a6ec7ea201cf7"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / ab96076a0e76 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.global_network_list](resources--azure_vnet_site--reference--group-005.md#canonical-4ed19e66e966b9eda39ebfb2066578c56f60cf14018d888b308272bd5c7593d3)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-005.md#canonical-8badaeb09cc425e0bb1045224b8aeeb2b890e5ea5cc125ea47ae0b63004adf4f)
- ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-792623dd3edcb9b0b718bb4faa9e42d34cbf2d96ddec66d1b955a661aac45f44"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-3aa53786df5113c34011f71c705e23085a0cb49601a9a2ffd7267e880a08813d"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / ab96076a0e76 / 3

- [global_vn](resources--azure_vnet_site--reference--group-005.md#canonical-bf6a9b25ca30cdf554b6c32126f70f9ae05658a8cee2d1c37bdd996a04e23981): complete subsection reference.

<a id="canonical-90b0fbcb471f4c08b0ef83dab27c6b2e5fef2bca171e8ae98fbf0acbcaa5a547"></a>

## Next pages — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / ab96076a0e76 / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--azure_vnet_site--reference--group-005.md#canonical-bf6a9b25ca30cdf554b6c32126f70f9ae05658a8cee2d1c37bdd996a04e23981)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-005.md#canonical-8badaeb09cc425e0bb1045224b8aeeb2b890e5ea5cc125ea47ae0b63004adf4f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-bf6a9b25ca30cdf554b6c32126f70f9ae05658a8cee2d1c37bdd996a04e23981"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b20343470e661c3526b6a1ace14f2c5e1d4eeef0843135500e3bc267b33dbb4"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / 745f9ab3fd01 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.global_network_list](resources--azure_vnet_site--reference--group-005.md#canonical-4ed19e66e966b9eda39ebfb2066578c56f60cf14018d888b308272bd5c7593d3)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-005.md#canonical-8badaeb09cc425e0bb1045224b8aeeb2b890e5ea5cc125ea47ae0b63004adf4f)
- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-005.md#canonical-d9f78026ba8964fdb9d0ec325120d98dec2ede52a20b385d8cc2b9f772496459)
- ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-35bfe47068c30a3a40d064efa10c23e86d4dfd747d6d75a0a7233dfe7e20878a"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-85790e214a18521b42049a6ff98c2bd7e7fce311dac673e19440ba8f51fa9ecd"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / 745f9ab3fd01 / 3

<a id="canonical-ac4fc97e7c184249fd470170efb075f8c6a67eb00a959c98b17550ce282099ce"></a>

<a id="canonical-f07ab7b5839cdd9d13a3cf02c1054bede601d103c0149b4f96ba0f89f8cca429"></a>

## name property — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / 745f9ab3fd01 / 4

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

<a id="canonical-9fb0181f3313d899229ea99c7c33352746d098bb0531aecd58bc04dfc480c18d"></a>

<a id="canonical-f5caab46a69a268c61ea5dc4711e2fa835b15adec8d666564ec417ec8381d092"></a>

## namespace property — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / 745f9ab3fd01 / 5

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

<a id="canonical-67d0405f1a4189bcb46f1c575093ebec6fc8e4e87633a17ab7e3b006aff8b7ee"></a>

<a id="canonical-c6c9ecb1a294835ab2d2f8c63c72c8ba3bbb5d8ed93336818cb2fb577dd8c392"></a>

## tenant property — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / 745f9ab3fd01 / 6

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

<a id="canonical-634f7b2ca2a98b6c6ef3f0c76d0f9edc8c2b2a0cf11f36351de3f1bd139f82c0"></a>

## Next pages — ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_globa / 745f9ab3fd01 / 7

- [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-005.md#canonical-d9f78026ba8964fdb9d0ec325120d98dec2ede52a20b385d8cc2b9f772496459)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d20ccdd8ac497cf011534fba7bfb1a1e22749db301720c02ab5c8e23dcc04bb2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6461734d47675ab322d43a4a83dd356a474197f439588a88d016b3d417c702d9"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / a8d1f483498e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.global_network_list](resources--azure_vnet_site--reference--group-005.md#canonical-4ed19e66e966b9eda39ebfb2066578c56f60cf14018d888b308272bd5c7593d3)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-005.md#canonical-8badaeb09cc425e0bb1045224b8aeeb2b890e5ea5cc125ea47ae0b63004adf4f)
- ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-354faa548d7872257786a01ff1325c8b06f0f6cc2015ff0bb154df1ea32cc0b8"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

Receipt-pinned upstream constraints:

```json
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
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-689db28cf2f8996cdc56b65261bfc0680b7c11a7b059a289dedfbaf0555e7790"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / a8d1f483498e / 3

- [global_vn](resources--azure_vnet_site--reference--group-005.md#canonical-ad7b4c1705e75c9e834857a82338a90c97b7b98da35717fb53d485bce8e539d5): complete subsection reference.

<a id="canonical-b675fd2ee3003708dfa67cc936b505986df763b5f80d26c7c0ebc596cb84fddb"></a>

## Next pages — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / a8d1f483498e / 4

- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--azure_vnet_site--reference--group-005.md#canonical-ad7b4c1705e75c9e834857a82338a90c97b7b98da35717fb53d485bce8e539d5)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-005.md#canonical-8badaeb09cc425e0bb1045224b8aeeb2b890e5ea5cc125ea47ae0b63004adf4f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ad7b4c1705e75c9e834857a82338a90c97b7b98da35717fb53d485bce8e539d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95ba53b93c01cffa1d1a631be5969cb9935dbf6a9530ea8ad4b42e2fae50ef81"></a>

## ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / f5df30af052a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.global_network_list](resources--azure_vnet_site--reference--group-005.md#canonical-4ed19e66e966b9eda39ebfb2066578c56f60cf14018d888b308272bd5c7593d3)
- [ingress_egress_gw_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-005.md#canonical-8badaeb09cc425e0bb1045224b8aeeb2b890e5ea5cc125ea47ae0b63004adf4f)
- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-005.md#canonical-d20ccdd8ac497cf011534fba7bfb1a1e22749db301720c02ab5c8e23dcc04bb2)
- ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-32515ee1d81eeb6c4d1a217530f2b64ce47f4beca39391d5ea90a4d9f763d175"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-67f0cf8297cc552efc58fbec18c93f488d2dce063d53e8ec2494930e48976159"></a>

## Direct properties — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / f5df30af052a / 3

<a id="canonical-8016a43924fa13f3e8394660ed0adfd4593046aacccf45cd7e39f0bd7fcf6c16"></a>

<a id="canonical-d58c84e48f980927cd069ddf2e12740cc2532eb15ab0657f697927e7348ebdfc"></a>

## name property — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / f5df30af052a / 4

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

<a id="canonical-9a3f2c02b76c1eb73c7fe477bb5c3a8104757387857348afc6602210f4a17aa7"></a>

<a id="canonical-097b50b2694a7aaedcb1e75e672880e7f4c24075c16cbf851b6ce151fb516097"></a>

## namespace property — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / f5df30af052a / 5

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

<a id="canonical-e396f2b78bcebac599896c48552a9207ded8f5b0ed7c2f599fdee01f098be29c"></a>

<a id="canonical-e39f2f146b2375a43576a97eb2b3035c96194694cd81df30eae0ce31dd716df3"></a>

## tenant property — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / f5df30af052a / 6

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

<a id="canonical-ee3b71bbbe6d1de0da6084347365646662858f67f4f9e8ee596dc1e495b63fa9"></a>

## Next pages — ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_globa / f5df30af052a / 7

- [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-005.md#canonical-d20ccdd8ac497cf011534fba7bfb1a1e22749db301720c02ab5c8e23dcc04bb2)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-647a004d298a0873bedd7fbe8943e7bb92dad0000cf2442d6a15f079cca91b83"></a>

## ingress_egress_gw_ar.hub — ingress_egress_gw_ar.hub / 672119df9d93 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.hub

<a id="canonical-788487844c42add5f1922da96839c6979a80efac51bd0c132b57f6ccbf4f42e7"></a>

Type: `"object"`. single nested block, Optional.

Hub VNet type. Hub VNet type.

Upstream description:

Hub VNet type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("express_route_disabled",
    "express_route_enabled")}
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
  "x-ves-oneof-field-express_route_choice": "[\"express_route_disabled\",\"express_route_enabled\"]"
}
```

Terraform syntax:

```terraform
hub {
  # Configure direct properties listed below.
}
```

<a id="canonical-b2c3a0d3be5393e1391047006bf7d1f263272252cf6ee1c749de13dae28e48e8"></a>

## Direct properties — ingress_egress_gw_ar.hub / 672119df9d93 / 3

- [express_route_disabled](resources--azure_vnet_site--reference--group-005.md#canonical-3f0b26200b03929bf64b24146f3284adea2b58427762b195c3b295a2419228b3): complete subsection reference.

- [express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a): complete subsection reference.

- [spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b): complete subsection reference.

<a id="canonical-5251ebd8ec8f1465c03e2b5cdfb56e545a348f7193545f9230f0e002a5324580"></a>

## Next pages — ingress_egress_gw_ar.hub / 672119df9d93 / 4

- [ingress_egress_gw_ar.hub.express_route_disabled](resources--azure_vnet_site--reference--group-005.md#canonical-3f0b26200b03929bf64b24146f3284adea2b58427762b195c3b295a2419228b3)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [ingress_egress_gw_ar.hub.spoke_vnets](resources--azure_vnet_site--reference--group-006.md#canonical-43a5d191b8eee55ba9b749e4293819d3e12e6b2dd09125fa56cbd5421215571b)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3f0b26200b03929bf64b24146f3284adea2b58427762b195c3b295a2419228b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d247833a1088c49f8cef618d24a64b835827d38cc5e3e3b1ea8d08b96b299969"></a>

## ingress_egress_gw_ar.hub.express_route_disabled — ingress_egress_gw_ar.hub.express_route_disabled / b7f63e8e3313 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- ingress_egress_gw_ar.hub.express_route_disabled

<a id="canonical-c771b631bf5470dc071986c01f94878c2da1df20bb38450c9cb0b849dca68d93"></a>

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
express_route_disabled = {}
```

<a id="canonical-d845f4769ad81a295c5102575462e0d07d705ae1a1edee4e812720732d6fd82e"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_disabled / b7f63e8e3313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e856dd252c675d612a34f2e0efd723e64b9e6dc9c497dc9ceffc16696743b7be"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_disabled / b7f63e8e3313 / 4

- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3365fbf5455ddafcbf878a026968c9db93f744e3139fba5beca331fcfb8c2050"></a>

## ingress_egress_gw_ar.hub.express_route_enabled — ingress_egress_gw_ar.hub.express_route_enabled / 87ea369b3d58 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- ingress_egress_gw_ar.hub.express_route_enabled

<a id="canonical-8bdb81378f2bda87f9aa13a2206708b9c31b60dbf577c93b71fc011ed833c119"></a>

Type: `"object"`. single nested block, Optional.

Express Route Configuration. Express Route Configuration.

Upstream description:

Express Route Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("connections"),
  validators.ConflictingObjectAttributes("advertise_to_route_server",
    "do_not_advertise_to_route_server"),
  validators.ConflictingObjectAttributes("auto_asn",
    "custom_asn"),
  validators.ConflictingObjectAttributes("site_registration_over_express_route",
    "site_registration_over_internet"),
  validators.ConflictingObjectAttributes("sku_ergw1az",
    "sku_ergw2az"),
  validators.ConflictingObjectAttributes("sku_ergw1az",
    "sku_high_perf"),
  validators.ConflictingObjectAttributes("sku_ergw1az",
    "sku_standard"),
  validators.ConflictingObjectAttributes("sku_ergw2az",
    "sku_high_perf"),
  validators.ConflictingObjectAttributes("sku_ergw2az",
    "sku_standard"),
  validators.ConflictingObjectAttributes("sku_high_perf",
    "sku_standard")}
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
  "x-ves-oneof-field-asn_choice": "[\"auto_asn\",\"custom_asn\"]",
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_express_route\",\"site_registration_over_internet\"]",
  "x-ves-oneof-field-sku_choice": "[\"sku_ergw1az\",\"sku_ergw2az\",\"sku_high_perf\",\"sku_standard\"]",
  "x-ves-oneof-field-spoke_vnet_routes": "[\"advertise_to_route_server\",\"do_not_advertise_to_route_server\"]"
}
```

Terraform syntax:

```terraform
express_route_enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-f6fe36827d82a85420296cf7d5d895a7d8351531dea61da741284803feacaaca"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled / 87ea369b3d58 / 3

- [advertise_to_route_server](resources--azure_vnet_site--reference--group-005.md#canonical-93bfb28b3ab2d8aacf50c414960a7654436e97c922b33c6601f085e43599c385): complete subsection reference.

- [auto_asn](resources--azure_vnet_site--reference--group-005.md#canonical-d24595ce026d601670bbb5a983e4f5857659dc9d44c05a597c4e8e44222e8618): complete subsection reference.

- [connections](resources--azure_vnet_site--reference--group-005.md#canonical-fcdf88fba2fd9f53af8b5af69def1d59f4eadf9f27bc096dac177d6f5c05bf93): complete subsection reference.

<a id="canonical-fce36dc44e906157a619ff66c92633f3361f53d221edc0aef1068e0db47eeb28"></a>

<a id="canonical-036dcbc703f28cc60f5b7315de45394ff1ab33601107030e7bece77dc79d8a6c"></a>

## custom_asn property — ingress_egress_gw_ar.hub.express_route_enabled / 87ea369b3d58 / 4

Type: `"number"`. Optional.

Exclusive with \[auto\_asn\] Set custom ASN for F5XC Site.

Upstream description:

Exclusive with \[auto\_asn\] Set custom ASN for F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 65535),
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
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "1",
    "ves.io.schema.rules.uint32.lte": "65535",
    "ves.io.schema.rules.uint32.not_in_ranges": "65515,65517,65518,65519,65520,8074,8075,12076,23456"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "1",
    "ves.io.schema.rules.uint32.lte": "65535",
    "ves.io.schema.rules.uint32.not_in_ranges": "65515,65517,65518,65519,65520,8074,8075,12076,23456"
  }
}
```

- [do_not_advertise_to_route_server](resources--azure_vnet_site--reference--group-006.md#canonical-08064fd4d5b38fbe45aeafeab03fb56140e7fc9ef75ecf45d4cbf3acc3c399b4): complete subsection reference.

- [gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-e7406a16de739030a0cf19472d9037dcf845eaf6d76d69ef183a4db0e09a0219): complete subsection reference.

- [route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-c55b5ed2085f9e5b41dcc160c63e37f695add2cb74be66c808f977126f7ea3cd): complete subsection reference.

- [site_registration_over_express_route](resources--azure_vnet_site--reference--group-006.md#canonical-1f85c3c59439548e46dd657bce79e1340e9dd6d7db8ac7940d56f21a43a06f94): complete subsection reference.

- [site_registration_over_internet](resources--azure_vnet_site--reference--group-006.md#canonical-3f42cb467712abca4dfe5ff8a68d1334df921aa3a01af66ab5a00e69fa78a31f): complete subsection reference.

- [sku_ergw1az](resources--azure_vnet_site--reference--group-006.md#canonical-701512fd97c100e810c9dedbcc76e5fa8078c4853732272e77075918cb409a39): complete subsection reference.

- [sku_ergw2az](resources--azure_vnet_site--reference--group-006.md#canonical-803e25e21c265627f6e20f4b3acbe8803ebacdc0945a9ae747f5c051c3502582): complete subsection reference.

- [sku_high_perf](resources--azure_vnet_site--reference--group-006.md#canonical-450890e587b8316cd9f2649d6999aa0268b3ca8eb0ea5b1345f34de01d31edc8): complete subsection reference.

- [sku_standard](resources--azure_vnet_site--reference--group-006.md#canonical-374dea40b689927c70cc506099a6fa9a917e62876e218114c4b97b326a01f29a): complete subsection reference.

<a id="canonical-b85a08e7cfc0fc49bf8fa7a111aa79e3ccb60dcdbf4199168a7dc828962bce72"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled / 87ea369b3d58 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server](resources--azure_vnet_site--reference--group-005.md#canonical-93bfb28b3ab2d8aacf50c414960a7654436e97c922b33c6601f085e43599c385)
- [ingress_egress_gw_ar.hub.express_route_enabled.auto_asn](resources--azure_vnet_site--reference--group-005.md#canonical-d24595ce026d601670bbb5a983e4f5857659dc9d44c05a597c4e8e44222e8618)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](resources--azure_vnet_site--reference--group-005.md#canonical-fcdf88fba2fd9f53af8b5af69def1d59f4eadf9f27bc096dac177d6f5c05bf93)
- [ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server](resources--azure_vnet_site--reference--group-006.md#canonical-08064fd4d5b38fbe45aeafeab03fb56140e7fc9ef75ecf45d4cbf3acc3c399b4)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-e7406a16de739030a0cf19472d9037dcf845eaf6d76d69ef183a4db0e09a0219)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](resources--azure_vnet_site--reference--group-006.md#canonical-c55b5ed2085f9e5b41dcc160c63e37f695add2cb74be66c808f977126f7ea3cd)
- [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route](resources--azure_vnet_site--reference--group-006.md#canonical-1f85c3c59439548e46dd657bce79e1340e9dd6d7db8ac7940d56f21a43a06f94)
- [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet](resources--azure_vnet_site--reference--group-006.md#canonical-3f42cb467712abca4dfe5ff8a68d1334df921aa3a01af66ab5a00e69fa78a31f)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az](resources--azure_vnet_site--reference--group-006.md#canonical-701512fd97c100e810c9dedbcc76e5fa8078c4853732272e77075918cb409a39)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az](resources--azure_vnet_site--reference--group-006.md#canonical-803e25e21c265627f6e20f4b3acbe8803ebacdc0945a9ae747f5c051c3502582)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf](resources--azure_vnet_site--reference--group-006.md#canonical-450890e587b8316cd9f2649d6999aa0268b3ca8eb0ea5b1345f34de01d31edc8)
- [ingress_egress_gw_ar.hub.express_route_enabled.sku_standard](resources--azure_vnet_site--reference--group-006.md#canonical-374dea40b689927c70cc506099a6fa9a917e62876e218114c4b97b326a01f29a)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-93bfb28b3ab2d8aacf50c414960a7654436e97c922b33c6601f085e43599c385"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-006f9371fa3a0afb5ecb2ba71a60f0acd53d4e7e7e1a6676b7cdbc6462e6720d"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server — ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server / ecce802e60d3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server

<a id="canonical-489d54827eab42075a7a0a7ee0533c3784d7294a55cec33af4c37105350cfe98"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for advertise to route server.

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
advertise_to_route_server = {}
```

<a id="canonical-ec52144a735dacad673c05391fdb5cfb5fbc19dc6cdc7907d3359e7285a41b05"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server / ecce802e60d3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f89914e0c0b88fbaa11bce060b70ab38b6c576e28dfa3c7d4874a5e7ffe5c716"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server / ecce802e60d3 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d24595ce026d601670bbb5a983e4f5857659dc9d44c05a597c4e8e44222e8618"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b451deed3da2be982b1c8b427f086422e55ad9b2e3f8e24e8edc139c0cc07ff2"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.auto_asn — ingress_egress_gw_ar.hub.express_route_enabled.auto_asn / 4edcecb248dd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.hub](resources--azure_vnet_site--reference--group-005.md#canonical-bad36d77d9668661adafc6078c3f88ebcbd3f866692376004e816787719f5e86)
- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- ingress_egress_gw_ar.hub.express_route_enabled.auto_asn

<a id="canonical-7993dc0797a8ecc3abaef63b6809ff61aa843dd9cf4052ff333283186051f720"></a>

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
auto_asn = {}
```

<a id="canonical-2b14e41c60387e65c68d1dae4e441d20eba9df3a16a850d154441bcb0ac3f655"></a>

## Direct properties — ingress_egress_gw_ar.hub.express_route_enabled.auto_asn / 4edcecb248dd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-34003a5b410d53c21b127830b2e9bc2c23ebe68c1467af9b22456e87733b661e"></a>

## Next pages — ingress_egress_gw_ar.hub.express_route_enabled.auto_asn / 4edcecb248dd / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](resources--azure_vnet_site--reference--group-005.md#canonical-98c75b1b2c7aff76698e75daad9260827a275cb0a18c5b6e566f03a22d156d8a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-fcdf88fba2fd9f53af8b5af69def1d59f4eadf9f27bc096dac177d6f5c05bf93"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
