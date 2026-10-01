---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-b667446d22697ec2602dbdceeb927788d5db6eb3ef6debfcc99b6d7134fba18e"></a>

## namespace property — voltstack_cluster.k8s_cluster / a267cbb72ae9 / 5

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

<a id="canonical-bc876c364d533a18d0bfab9c583ff2f37dd679d6f2815a69796ebb871c64792b"></a>

<a id="canonical-b7b9cd9a1e8a686f4983a75b283e61074469e6cfe5ecad7e23e9dd96aa4f05aa"></a>

## tenant property — voltstack_cluster.k8s_cluster / a267cbb72ae9 / 6

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

<a id="canonical-708cc1541967169408edb8c0aabdd6136dcfc0c3e93f6d0f8e12e704a56c074f"></a>

## Next pages — voltstack_cluster.k8s_cluster / a267cbb72ae9 / 7

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-289e5681a9cf2e53242dcc84d84f97ca791fb0c3383962a8ba1d9a73adaba898"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0ba97a55755611688507e020b2ecd1c68faf0bda0dfa0d341379f2168fe9dc09"></a>

## voltstack_cluster.no_dc_cluster_group — voltstack_cluster.no_dc_cluster_group / 271bfd1a6bcf / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.no_dc_cluster_group

<a id="canonical-01c983227367986d27e94f3093f1492a80191ee629534f67f12f775f035d9fcd"></a>

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

<a id="canonical-7170c37012fa577220b64db3e1864fd347807faae3951e1cdc1ddeb507f3dab9"></a>

## Direct properties — voltstack_cluster.no_dc_cluster_group / 271bfd1a6bcf / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9cf36fc61ea9105b90397c1efce0b8b1049b4a2cb5e4b149bd06df90a79f4fdd"></a>

## Next pages — voltstack_cluster.no_dc_cluster_group / 271bfd1a6bcf / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-da0a9f761b85bb12907956b73c9de07c013d2a221aed8b38eb299e527f86258f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2a38610caa50b6a61386afc3bc5fbb19d7f4f38432495acba73eab5401d4c547"></a>

## voltstack_cluster.no_forward_proxy — voltstack_cluster.no_forward_proxy / 0e2ea829039f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.no_forward_proxy

<a id="canonical-4fd70a9bf0257154f5762bce7eab1bbdb935af4b690e248fc805a4e1c7e04908"></a>

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

<a id="canonical-accc1e85571e0434b9177e549d2ae0eb0a645e25d8a6524197798651a972ead2"></a>

## Direct properties — voltstack_cluster.no_forward_proxy / 0e2ea829039f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-30be9e908b1e6accd106197cc08a9a380407c49db038e4d6b850611fafddcb7e"></a>

## Next pages — voltstack_cluster.no_forward_proxy / 0e2ea829039f / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a832f2f4936449e6da6a29d31b957acde0e04f79c09f32485b52ecabc67ab28c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-85c37709eb56b09220f5d2b8572bd6364e3aedcde595968972f5923085bb6639"></a>

## voltstack_cluster.no_global_network — voltstack_cluster.no_global_network / 4ac4f38b15e4 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.no_global_network

<a id="canonical-17630e6d42a3f013c80569b286acc14e59a8635d7c6856b7bf832c7a6bfd4bab"></a>

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

<a id="canonical-37c8528dff3617d7aa2b42472e309757b8039c3fd9132764ce9221e051b2f5aa"></a>

## Direct properties — voltstack_cluster.no_global_network / 4ac4f38b15e4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-98d2e7a701aab94e6e0a552f66bff8d01677a475ae365487cfd4f2358653eb60"></a>

## Next pages — voltstack_cluster.no_global_network / 4ac4f38b15e4 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-eebc49664723328c5667c67d2c0b2ef51f3daaf7ae64eb1fd7f9bc93b7bdc76d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c0ca7b4cb95029edd16234a1d69131a42da83400eb6b1cb053481c4cc0c2e90"></a>

## voltstack_cluster.no_k8s_cluster — voltstack_cluster.no_k8s_cluster / 59a90501fb09 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.no_k8s_cluster

<a id="canonical-4c11d0a94aea0864cabb9121cb2ef9afd24e05178de9c4fe3a5e68194e8d80f4"></a>

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
no_k8s_cluster = {}
```

<a id="canonical-565ee54fb83404a0049c368dd121948cb3023a7427007148638fa15bd39d2f51"></a>

## Direct properties — voltstack_cluster.no_k8s_cluster / 59a90501fb09 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8a8beec8eac6f0efed1168874cf1f7304d51b62deed4ed2f0969c1d84576ad57"></a>

## Next pages — voltstack_cluster.no_k8s_cluster / 59a90501fb09 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-07114fe9e1bfc024bf75e993864a9f4f594519b6e0a53400b7b83ec05ea2ce19"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8655edd808311e1860d35d7d323da110193829f71e5d637a29aedbd99487321"></a>

## voltstack_cluster.no_network_policy — voltstack_cluster.no_network_policy / 9a92f55c1d04 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.no_network_policy

<a id="canonical-11ad86581f436ca2b0310ad208751e8665d910c5e552046230cdca56a45844f5"></a>

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

<a id="canonical-1121ff8973771195b41943d3d79b217b8076ee44dd2e40f3b18998a78e622087"></a>

## Direct properties — voltstack_cluster.no_network_policy / 9a92f55c1d04 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1b6a65d9b9224bd1595eac3a0b46c967367a58a8cf67b65bbc904fd3d0dba923"></a>

## Next pages — voltstack_cluster.no_network_policy / 9a92f55c1d04 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-458d512fac829eb4805efac48006a46ceafc83a4a8805838e2191731127078a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23c120a0c4f91851c08bbdede4ca1ffbfcc24beef317b9890267bef223b8b2a7"></a>

## voltstack_cluster.no_outside_static_routes — voltstack_cluster.no_outside_static_routes / 5a0ace4a50b0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.no_outside_static_routes

<a id="canonical-95e01b24560a3b87705254e2d4833e482bd3c9408592632171f675a023dd314e"></a>

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

<a id="canonical-ab86d1da30efdb0fc4a8b30d1e00fc296fb5b555140ddfb4d50475f9a7bb8774"></a>

## Direct properties — voltstack_cluster.no_outside_static_routes / 5a0ace4a50b0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8e52a6bf224e12c4ccc729b3da7f291542cc05c4ada32fb7d18026cb6e85eddf"></a>

## Next pages — voltstack_cluster.no_outside_static_routes / 5a0ace4a50b0 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9b6605bec25e1d5090e057a1371ce3ef73328d1a779d881784d6f3d9c901169e"></a>

## voltstack_cluster.outside_static_routes — voltstack_cluster.outside_static_routes / d399c87d90cc / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.outside_static_routes

<a id="canonical-005a8ac525f4d27590a13a531226caf69cef2edacd675568a2912c96b038d7ff"></a>

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

<a id="canonical-7f48bd01a320881eabb77440e4866064496367030fd19d02f0ceec3c965c1b20"></a>

## Direct properties — voltstack_cluster.outside_static_routes / d399c87d90cc / 3

- [static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26): complete subsection reference.

<a id="canonical-5a6e6180b86c2984badfc6d0916349de50281ac5813b9c68a4372d6a20435091"></a>

## Next pages — voltstack_cluster.outside_static_routes / d399c87d90cc / 4

- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aaf33ca89b0eea40b287f1e5fde849de29945d8b5c3300095d8e6dce0ab1a3f7"></a>

## voltstack_cluster.outside_static_routes.static_route_list — voltstack_cluster.outside_static_routes.static_route_list / a568f88906a3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- voltstack_cluster.outside_static_routes.static_route_list

<a id="canonical-8dbcfee6f76893d09eae53cf6dd10cf5cff13bf1246d4f0503c5b0d4eff7c082"></a>

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

<a id="canonical-6bca68ef6d7f925fa329484a144a37dee19141ef40ab12a2ccf23e5cffde3216"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list / a568f88906a3 / 3

- [custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67): complete subsection reference.

<a id="canonical-7008a25f0b958c6696540c35f862e3838973b209fa31f04928c63d7628b848ea"></a>

<a id="canonical-8fc071453a9f04a3f8095a594f4e9f60a1e3339edc3cc7b1b6deafe9e661a848"></a>

## simple_static_route property — voltstack_cluster.outside_static_routes.static_route_list / a568f88906a3 / 4

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

<a id="canonical-c2b3e16e536190bb7bd1089601c849a5ec457b5b4d25c2c0d9e806fb3ce6ed82"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list / a568f88906a3 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f799d049410f4a7c52465cab82469817398e9098ddfdd7f7894658f313af856"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route / ba50caf89d42 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-9b319f5b2ab86c4e3d5c79385cf42e654dec6881a8e3bf24b2c0612d4dcecab4"></a>

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

<a id="canonical-1f3c42439473df6ad312bbc170dd0647a54e717ffcbfaf44078eaa64c062e81b"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route / ba50caf89d42 / 3

<a id="canonical-afae734288358958ff496942a47214f4f3e176396429d161a9748a6f95f56768"></a>

<a id="canonical-0a818ea95e6fd0ccc91b39592ff2f5da557d71a6484e5ee26d7909a6be6851a0"></a>

## attrs property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route / ba50caf89d42 / 4

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

- [labels](resources--azure_vnet_site--reference--group-009.md#canonical-d003d8196b5f4da5852cd14c26747d06a02573957b2432079283d0d984e7f109): complete subsection reference.

- [nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0): complete subsection reference.

- [subnets](resources--azure_vnet_site--reference--group-009.md#canonical-b6af287a94d9ce178a2ba535801f73450417b815acaa1d07965724c5e0534eea): complete subsection reference.

<a id="canonical-5a1a53eb0f097c191f2726a92879445d346c95c38eaf52b2322afa58a1d5768b"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route / ba50caf89d42 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](resources--azure_vnet_site--reference--group-009.md#canonical-d003d8196b5f4da5852cd14c26747d06a02573957b2432079283d0d984e7f109)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-009.md#canonical-b6af287a94d9ce178a2ba535801f73450417b815acaa1d07965724c5e0534eea)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d003d8196b5f4da5852cd14c26747d06a02573957b2432079283d0d984e7f109"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a63356c2c7201d650d614c737ec3df835ff1c6f913f13e72a1a8e6b64d9d7fb"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.la / 107e7918cf6b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-4793b8751a2f8e46cfed5fa62f54e80dfaf892e69f3cec865bb3127dcf135f98"></a>

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

<a id="canonical-add2d6f04ec2fe2bd6256c2ed5acd0113b25b1fd603a9b9fcf69042b698f7fd6"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.la / 107e7918cf6b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-efe76d0368269fb397cb1481b0cebe00540a5ee7a5f524a27ce7aa242f3cc302"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.la / 107e7918cf6b / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fa76dce88e81559767f9b26a15c99bdd33f8dd3586f89a5e52cc616bf2e9570c"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 8947a449bd60 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-497e3167a82736449de45dd2a1c1e55a7aba3d7be75029c9db95e42c5c25a1a0"></a>

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

<a id="canonical-3872dbe463b7c8e7a53756069e34d3b8ce8b263d2866d7dc97e6ee69442d0c56"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 8947a449bd60 / 3

- [interface](resources--azure_vnet_site--reference--group-009.md#canonical-4d7e452ec0875ae9c61822c710130e3a09d3a6fb8ea36f998062db260e25ace5): complete subsection reference.

- [nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-d400119952bd598fa3e184276ac97ad5922a8da68b438611e9bebeb710b72148): complete subsection reference.

<a id="canonical-13cbeeb33d165733344995289aea5641ede5d364e08f093b712bd7472ea732a6"></a>

<a id="canonical-3583e2b4799fb3286cac77cd80bf137d3717fd5f6f216542ea01503959de123d"></a>

## type property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 8947a449bd60 / 4

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

<a id="canonical-7cd272e71a4744c8f6cdfc41e7e7e4cac8bc4da05c712a1c8142e27636ebe758"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 8947a449bd60 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--azure_vnet_site--reference--group-009.md#canonical-4d7e452ec0875ae9c61822c710130e3a09d3a6fb8ea36f998062db260e25ace5)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-d400119952bd598fa3e184276ac97ad5922a8da68b438611e9bebeb710b72148)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-4d7e452ec0875ae9c61822c710130e3a09d3a6fb8ea36f998062db260e25ace5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c72b1fa4c1ce21e47946cebd60b0547d75c7e812629531f18f9306cee930caf"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9cce9872edc4 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-990269aa9e71352235f50f07ab356d6b131b24f8010e558eeb7e613d296761c1"></a>

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

<a id="canonical-69de2a41a2126d4a126c97dc199e73f60589f3ecf86dedb939427a910b4b21b4"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9cce9872edc4 / 3

<a id="canonical-0d6997b8fb91d3854d18b8b6b418934b4ac2494f7a9752857bd49b728c8bdbf5"></a>

<a id="canonical-06b70e36d3bd29e2f5dcac849f94f157a91329882c1f5a6e82133c4a3f0a8da1"></a>

## kind property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9cce9872edc4 / 4

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

<a id="canonical-6873489b3bf43e618dde321c4c7b88ddde8ad0a09adf38bbfb953339ffca883f"></a>

<a id="canonical-220e4300bea26b0b89d6c046b7370f2ec1d0d093085d0112e2e2d97ed4f5a550"></a>

## name property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9cce9872edc4 / 5

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

<a id="canonical-9ff38918d1050e2b8f73e22229778b3e7fb5ac494503cd4f75d5bd0f0e5dd4ce"></a>

<a id="canonical-02df8f7cc183789e611fa566f04c5d6ac3fa46f879bd062144c3f6e6667a8827"></a>

## namespace property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9cce9872edc4 / 6

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

<a id="canonical-a435260f22159bec3d9b2343304e1123812b2817ff25caf550f0e630fba56295"></a>

<a id="canonical-f74301df852b381382b681816ce80cc007ffdf81833c4adef6938753648d40f0"></a>

## tenant property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9cce9872edc4 / 7

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

<a id="canonical-717fb764cb744a65d2356bb94f1535cea59b978aa6bd7a3197b74849659913b0"></a>

<a id="canonical-36cc42ee048408517a5be44ab8da7b9b17a1fe7946ef5a36da6d76665c0cf3ac"></a>

## uid property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9cce9872edc4 / 8

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

<a id="canonical-8f924b72144badec13e51ff6dbeb0d545b7e3aba268aec8610c4c5f8f89e6874"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 9cce9872edc4 / 9

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d400119952bd598fa3e184276ac97ad5922a8da68b438611e9bebeb710b72148"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f8bf8d2b40af95157e75487e0039f77bd5f8c9180220d2e8b37d61348d5c2b2"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 6852776230de / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-06d122d4bd7f368976e6debf735cfda7dbdf06fb0a7693790f6a000caf57ec4a"></a>

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

<a id="canonical-4b41cfb73ccddcd4f8b5de28dc02819c7dc53345773b29b05a81058259a03b83"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 6852776230de / 3

- [dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-69f92cb64f77919a12532af5501549e68ec410706a786498781f2f3ab1e22336): complete subsection reference.

- [ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-e1139b2cb1fdc323552a5484b1a720d2b19ee86a272eb6361cdf44a556d30233): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-481cfa19da1f83925694a1916457b95db274bae165a59f2874fc79b1cf6fc964): complete subsection reference.

<a id="canonical-4d4795c96b27ee31bb5a5c59bba9c08c2a2bfd7d9254cbe108e6eceabf93c4d1"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 6852776230de / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-69f92cb64f77919a12532af5501549e68ec410706a786498781f2f3ab1e22336)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-e1139b2cb1fdc323552a5484b1a720d2b19ee86a272eb6361cdf44a556d30233)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-481cfa19da1f83925694a1916457b95db274bae165a59f2874fc79b1cf6fc964)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-69f92cb64f77919a12532af5501549e68ec410706a786498781f2f3ab1e22336"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-352ca82b46cdc546e82baf1deb25f961635e87d384da369643d3691217b55e86"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 0998918f7797 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-d400119952bd598fa3e184276ac97ad5922a8da68b438611e9bebeb710b72148)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-f664f3a89e6592a17db8c2d9b9d8f90a07d801c1c17bbbd28957561d2237d1dd"></a>

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

<a id="canonical-b6c5f6e1d88edd3ab527d643a15baa5e7596328497b109a187e5697e426492a5"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 0998918f7797 / 3

- [ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-777280cad1dd66949ccce6f6febe4fc33b68c0f1504c7cf2d2251fc836d76947): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-6c9e9d14cb5fabf4f370f910827832115cfc9f9420231bd2addd5b74709f2e00): complete subsection reference.

<a id="canonical-64fda9858ae910e0c82db4768d36df896730f5af880b7c8c40b4137bf41e2c77"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 0998918f7797 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-777280cad1dd66949ccce6f6febe4fc33b68c0f1504c7cf2d2251fc836d76947)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-6c9e9d14cb5fabf4f370f910827832115cfc9f9420231bd2addd5b74709f2e00)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-d400119952bd598fa3e184276ac97ad5922a8da68b438611e9bebeb710b72148)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-777280cad1dd66949ccce6f6febe4fc33b68c0f1504c7cf2d2251fc836d76947"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9655c63378859ba474cb656410611b6b53eb36900cd5014e5ecd854fe6cd8c82"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / a7bae4625bd8 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-d400119952bd598fa3e184276ac97ad5922a8da68b438611e9bebeb710b72148)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-69f92cb64f77919a12532af5501549e68ec410706a786498781f2f3ab1e22336)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-b90a6a94690a32a0e18168bd4ae738b59900093bde8b29ec07fde4c813b4fbaf"></a>

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

<a id="canonical-117a765eacb561a1e3f36636f19d823860107c3b49277b84b416d01bb01e3112"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / a7bae4625bd8 / 3

<a id="canonical-470c2ff71cce47a1e8820f344d33d65f762d8f03eda8eaa9855c2d1b245ed7f5"></a>

<a id="canonical-3923ee940a4dae9cb2b358917aeb9d0178c5e3d6cea97f1dff937b1facb406af"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / a7bae4625bd8 / 4

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

<a id="canonical-8eb092a87627bd6247bf17089e225751f2aba698a2d59896f7c4aadebcd06a5d"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / a7bae4625bd8 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-69f92cb64f77919a12532af5501549e68ec410706a786498781f2f3ab1e22336)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6c9e9d14cb5fabf4f370f910827832115cfc9f9420231bd2addd5b74709f2e00"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-61406bfced8597bc8d51278e9e4948efb0f1cce394d0204743c83c6bb7c98d81"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / e8457b6f8e14 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-d400119952bd598fa3e184276ac97ad5922a8da68b438611e9bebeb710b72148)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-69f92cb64f77919a12532af5501549e68ec410706a786498781f2f3ab1e22336)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-fe419758de05510826b845e53377095d9a54200f97c8952b3dff72fcbeadcfc4"></a>

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

<a id="canonical-e6b9faa5f4535ae42dadab65813e296be5b6d4b82c3997d185555c019d6cf310"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / e8457b6f8e14 / 3

<a id="canonical-860455ac6595f7986e6289f06c3175e4fe2ac0c287b738d29ffe3b49177be8d1"></a>

<a id="canonical-491c51cd5cc32298f925114a114c5ad89e67f22d61d1f532ea1d9cad197634c1"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / e8457b6f8e14 / 4

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

<a id="canonical-4805d0bf6ea2222ca5f8f6b8878ae97a5682be52c15877c448797b617b72a519"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / e8457b6f8e14 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-009.md#canonical-69f92cb64f77919a12532af5501549e68ec410706a786498781f2f3ab1e22336)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e1139b2cb1fdc323552a5484b1a720d2b19ee86a272eb6361cdf44a556d30233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a49d77112351e2bcadd70a4da4666f93b789594a200e032a06be6e5ab1e6b45"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 986b9b3cd57c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-d400119952bd598fa3e184276ac97ad5922a8da68b438611e9bebeb710b72148)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-677fb4c9cb401ee9e384363598886157d3f2a678afdccb7ae8a46f0f679c72fb"></a>

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

<a id="canonical-717536ed330365108b5d7cce1d4bde42b05cb501305d2b5da078ba8a00b2b72d"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 986b9b3cd57c / 3

<a id="canonical-524d9034a45737db759fbe3dd95300e09b0d9b0afea84d7e305117a50ad7d5c3"></a>

<a id="canonical-a48295deb17ac765bf219e0418e3564f9bcf12928360c3c764e549000262894b"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 986b9b3cd57c / 4

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

<a id="canonical-4926007cf266d6c3a013d7344ce17740ff6b4fcee641e83a118ceba583205b77"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 986b9b3cd57c / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-d400119952bd598fa3e184276ac97ad5922a8da68b438611e9bebeb710b72148)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-481cfa19da1f83925694a1916457b95db274bae165a59f2874fc79b1cf6fc964"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-30e11937c894dd4d5fcc5006c61323c21fc417955ac9dc23fd23bb787a902d5e"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 5ef2d78475be / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-009.md#canonical-cd1c5ed59aef35f24d88b75fb7aaf5bdb259fee30b5bcd7a1b0e1d59d787ddd0)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-d400119952bd598fa3e184276ac97ad5922a8da68b438611e9bebeb710b72148)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-b9d5690be9eaf965ddc0418b3ccd73ea361da9f46a1b043f4d5fe17fdad824fc"></a>

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

<a id="canonical-8b3b13dde27bedb8d1d86c47e44ceeb10311378d6c3e9367d0bf9402ee7d53d2"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 5ef2d78475be / 3

<a id="canonical-58e66c7768cb538e60cd366b93f0d3482f8d9f44ce1d15b8d2dd63fe1e1e58b7"></a>

<a id="canonical-6a690fd0cdb18e868d523dfdd2f75c25dc31041af8fa067d9d7f11e2cffdc1a5"></a>

## addr property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 5ef2d78475be / 4

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

<a id="canonical-80533f829f4100020d31481af7f7385f01c1266e5af1dab053ebf562fd0a1a09"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.ne / 5ef2d78475be / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-009.md#canonical-d400119952bd598fa3e184276ac97ad5922a8da68b438611e9bebeb710b72148)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b6af287a94d9ce178a2ba535801f73450417b815acaa1d07965724c5e0534eea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7facebff6169ec0fb17d65c0f7f0ad43e65a4aa87ce8494451863b2e5a5e80d6"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 2565750e262b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-fe430532bfc8d45628b345e22d5e8c7b56225d959b78dcb942c7d80b010795b1"></a>

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

<a id="canonical-25300b137212ef34405f86162d031103c5c4278d838432d93701fc70fb06dff0"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 2565750e262b / 3

- [ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-2d9392cf812fb15541f281b259742755bb6e02d161b8827233b6fea18527a1f2): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-088bf398eaaf3f528e6018cfe6e89a57acb7d4d69cf3a894f07aef65308232e0): complete subsection reference.

<a id="canonical-bcca08a75206a5b05fdd814f80d32d9bcc5dc65fb500bfa0b6db720a2a4801a9"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / 2565750e262b / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--azure_vnet_site--reference--group-009.md#canonical-2d9392cf812fb15541f281b259742755bb6e02d161b8827233b6fea18527a1f2)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--azure_vnet_site--reference--group-009.md#canonical-088bf398eaaf3f528e6018cfe6e89a57acb7d4d69cf3a894f07aef65308232e0)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-2d9392cf812fb15541f281b259742755bb6e02d161b8827233b6fea18527a1f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-902b16df7d8620223b1a02e558889519424b69cc1c0f2fb6e3490dad874da6a6"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / c4adcac67c89 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-009.md#canonical-b6af287a94d9ce178a2ba535801f73450417b815acaa1d07965724c5e0534eea)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-1d987d88c8d719ab2ae92fa53c7bb8729ccbcf1d45ecce2a9577937717480d06"></a>

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

<a id="canonical-ba0887c657afd1249890033d6edc05fabe1964efe6323c60031a5c666ba8735a"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / c4adcac67c89 / 3

<a id="canonical-77f74d78301b59aeca05b0a84963c99b72ecfd5ca9b488f286daf074e472cfa4"></a>

<a id="canonical-182403120b2d89f4a8478f1a20c678aadb359a3234c92a6e771856aedf06a2ba"></a>

## plen property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / c4adcac67c89 / 4

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

<a id="canonical-2724112f3e3a471f7dab35f6a02cda9903469a6db955a4f4d95f677ced920af8"></a>

<a id="canonical-ad1922f2f7791dbdb0c0df752d56bd573ca6d13cdc91e0aa60944feb2fd4fe41"></a>

## prefix property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / c4adcac67c89 / 5

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

<a id="canonical-75aadced5e8e0b99946477ad19a9b6a73b349ca3070dba98cd231856be56943f"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / c4adcac67c89 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-009.md#canonical-b6af287a94d9ce178a2ba535801f73450417b815acaa1d07965724c5e0534eea)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-088bf398eaaf3f528e6018cfe6e89a57acb7d4d69cf3a894f07aef65308232e0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc4dc998f4d01fc1f6e1ef6f1d0e0d69d75fbaf3a20d64907a44fe94c948d662"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / c198fcbbcf49 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-10b7b5c5ca2525688944c128802937ae34e397f8894fec2dde6dd487195ef2aa)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-009.md#canonical-f1d8e3c9c659ce5861364439eb4efb25d50dff1fbdeb401b432c17f7fc05ce26)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-009.md#canonical-fd10ea1c9b5b71e50aa04ccf6b755da3c3fdf7edd82abb58d3774e6de312bf67)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-009.md#canonical-b6af287a94d9ce178a2ba535801f73450417b815acaa1d07965724c5e0534eea)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-762abe3fd62a33e38c0a504f4a30a991747d9bd98cea4365b027d81aac5ee270"></a>

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

<a id="canonical-bb0e4eb97dbc27f6d87ca1e0e62fd661f5127daa3d578f23062f84e4d3e79faa"></a>

## Direct properties — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / c198fcbbcf49 / 3

<a id="canonical-91743284606e7711c8b1769cf61406c63bca78fc33457b27cd1e027063f3a972"></a>

<a id="canonical-cb912c2911c01193414e554b65b69ee784c69b7ee872477641b6bac2489de574"></a>

## plen property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / c198fcbbcf49 / 4

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

<a id="canonical-3abc7e976f3643fbd5fefa0b890e5f975ca21eec21b24c0f71ecdf466c47e701"></a>

<a id="canonical-25e4ef04a35f853e2db471f762cf20511dae8f01c9aebf782e5c4c8a2a71a40d"></a>

## prefix property — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / c198fcbbcf49 / 5

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

<a id="canonical-d0cc94e5cd63d7aec69f8683de68ab144a70f2e15122319de229b816b13fd603"></a>

## Next pages — voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.su / c198fcbbcf49 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-009.md#canonical-b6af287a94d9ce178a2ba535801f73450417b815acaa1d07965724c5e0534eea)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b97e72482667a76994786d259dab5022563f9130822b995a8826a686419910d5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-febefb3cebe24755608095f96cd109aa21098895502fd320746a493a2747594e"></a>

## voltstack_cluster.sm_connection_public_ip — voltstack_cluster.sm_connection_public_ip / 531fa2276311 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.sm_connection_public_ip

<a id="canonical-0e64710d65dd261cb82363a83a00651ff2bebda65b77187a28e535875f1b1483"></a>

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

<a id="canonical-b64d4e2530c58e48ed839a8c7930a8a40b7ac2de9df2baeff73993351a58ffb5"></a>

## Direct properties — voltstack_cluster.sm_connection_public_ip / 531fa2276311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d1a665c8956ec6d7682cffdbcf6590231f7d8fa32392907dda33f8d65351587e"></a>

## Next pages — voltstack_cluster.sm_connection_public_ip / 531fa2276311 / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c7d7c49bef3b6826e08eb566bee76348916abaaa4a828a15db5482e7b5ef09e9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0efa92e8572e69a30d37c7b13880f99901e78c74d2d45528cd8b77d2a2e535f3"></a>

## voltstack_cluster.sm_connection_pvt_ip — voltstack_cluster.sm_connection_pvt_ip / 5963e7c5990a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.sm_connection_pvt_ip

<a id="canonical-79d3f29e81e4ddb657e1ff55976bd6a272dfef3baa5e9d462c128e46c891eceb"></a>

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

<a id="canonical-9c9ac65d2223b409f2cb3812a7f2fb18e2ed98fa1a695b4dd92191be7d433f2a"></a>

## Direct properties — voltstack_cluster.sm_connection_pvt_ip / 5963e7c5990a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2da0c8f9c9bd3816cec3ee5afc669219e04d62e35d9db9e8d9d8bb1d5da6dd57"></a>

## Next pages — voltstack_cluster.sm_connection_pvt_ip / 5963e7c5990a / 4

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-f5eb60002d9390fbb074b58b83119e443d059edb53f00c5b35b3f02b327b480c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5efaca275b179ffd1aaba8e269160ad1eb42c7d4eb0cd5922143b89632e4c51"></a>

## voltstack_cluster.storage_class_list — voltstack_cluster.storage_class_list / 5859d287a0dd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- voltstack_cluster.storage_class_list

<a id="canonical-ebbbac7297dd69ee98b987cb5e82c19d06f76c3b820550aa1b1eed91fa27375c"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
storage_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-9052e117156d60798c9057ceb3a4d157b173c9b50dd4fe9f1ea620b8bf87c393"></a>

## Direct properties — voltstack_cluster.storage_class_list / 5859d287a0dd / 3

- [storage_classes](resources--azure_vnet_site--reference--group-009.md#canonical-340ef27f6f3c6c66e6f490bab0f529b8745fe3db842804aa75f6da180f5472cc): complete subsection reference.

<a id="canonical-c8945677d70624f34b4d06f29523a6020aafce05b864f2f33a2266a2cbe4c644"></a>

## Next pages — voltstack_cluster.storage_class_list / 5859d287a0dd / 4

- [voltstack_cluster.storage_class_list.storage_classes](resources--azure_vnet_site--reference--group-009.md#canonical-340ef27f6f3c6c66e6f490bab0f529b8745fe3db842804aa75f6da180f5472cc)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-340ef27f6f3c6c66e6f490bab0f529b8745fe3db842804aa75f6da180f5472cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46f38fca8a3eb8aa3ef63c2473859736c5745c7a2edafb9956fe1dcffe1d554b"></a>

## voltstack_cluster.storage_class_list.storage_classes — voltstack_cluster.storage_class_list.storage_classes / a5435b8cebfd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-d13cdf7dc89d470692111f43d9e212d9ccccabf10c25a6043c4d909ce91b8314)
- [voltstack_cluster.storage_class_list](resources--azure_vnet_site--reference--group-009.md#canonical-f5eb60002d9390fbb074b58b83119e443d059edb53f00c5b35b3f02b327b480c)
- voltstack_cluster.storage_class_list.storage_classes

<a id="canonical-173fc68118572451ff62609c04aeefe44de82036ade425259f4f3c3870bf14b6"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name")}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_classes {
  # Configure direct properties listed below.
}
```

<a id="canonical-854c578dfa2a7f976a9e8e7a4415380ed5e0235ca73eedef85402ccbf8848523"></a>

## Direct properties — voltstack_cluster.storage_class_list.storage_classes / a5435b8cebfd / 3

<a id="canonical-7ec16662c44df8c4cbda859266a69ad92d8a8b23677b7cd65e7767dd9becdf58"></a>

<a id="canonical-bb446b97055f578e55fbd64b8ef75b4e5e443d05cb544acd50987d9dba56442b"></a>

## default_storage_class property — voltstack_cluster.storage_class_list.storage_classes / a5435b8cebfd / 4

Type: `"bool"`. Optional.

Make this storage class default storage class for the K8s cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2b67c205708341d18cc579c1830af5e52b38b5de7393e407fbefa41cb455679e"></a>

<a id="canonical-1fcf1651ca849a3d9556fc116ed99cf08c25fea9deb09ffc7c8523908831ee1c"></a>

## storage_class_name property — voltstack_cluster.storage_class_list.storage_classes / a5435b8cebfd / 5

Type: `"string"`. Optional.

Name of the storage class as it will appear in K8s.

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
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-6a08614c4a75bbf39f46489cb41dfd00a186c2191e7d3a81c7916e79555502ac"></a>

## Next pages — voltstack_cluster.storage_class_list.storage_classes / a5435b8cebfd / 6

- [voltstack_cluster.storage_class_list](resources--azure_vnet_site--reference--group-009.md#canonical-f5eb60002d9390fbb074b58b83119e443d059edb53f00c5b35b3f02b327b480c)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-690c9a53b72c6d0cdeee488dd403d9912b2132f985cd54d155ad2d6160fb1d1c"></a>

## voltstack_cluster_ar — voltstack_cluster_ar / ec18bc5827cf / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- voltstack_cluster_ar

<a id="canonical-77d5ee0071ad6b3419b08ffb4c83dd9b8c37e6a49799a8989b298f6baa9195ca"></a>

Type: `"object"`. single nested block, Optional.

App Stack Cluster of single interface Azure nodes.

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
  validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("default_storage",
    "storage_class_list"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("k8s_cluster",
    "no_k8s_cluster"),
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

Terraform syntax:

```terraform
voltstack_cluster_ar {
  # Configure direct properties listed below.
}
```

<a id="canonical-5997689e41fa6f3f5be6ee0724e70e452d46b877eac28a8e01ff5eeb772b6287"></a>

## Direct properties — voltstack_cluster_ar / ec18bc5827cf / 3

- [accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-c2e78ccf67ee4d4433d5e27c99cf0629a9c984be70493b525f8e5e628236e7cd): complete subsection reference.

- [active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-71f2ecfba05031cb997e73c528066702be86b9746673e6492bec30957476576d): complete subsection reference.

- [active_forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-fb13dbc8a5f04e6473307f1a618175d7ab9bbc0d5f4db6349e27153924e78314): complete subsection reference.

- [active_network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-6f7e0b6e15090e23f2a161c319ee5a0198740fd2c07697baed0e5c75405a3fcb): complete subsection reference.

<a id="canonical-c194da1db5ca3609701ea45e5668a9671497b1831fcef52aac4efc609acd152a"></a>

<a id="canonical-36edada1dd84461bd2c2ea6cb39c95f6133cd316046b40a2313600347a87023a"></a>

## azure_certified_hw property — voltstack_cluster_ar / ec18bc5827cf / 4

Type: `"string"`. Optional.

\[Enum: azure-byol-voltstack-combo\] Azure Certified Hardware. Name for Azure certified hardware.
The only possible value is \`azure-byol-voltstack-combo\`.

Upstream description:

Name for Azure certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("azure-byol-voltstack-combo"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-voltstack-combo"
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
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [dc_cluster_group](resources--azure_vnet_site--reference--group-009.md#canonical-23445797a2a89420a7e6dbf95e5906ac49988575995da53770d8132f5ef0b561): complete subsection reference.

- [default_storage](resources--azure_vnet_site--reference--group-009.md#canonical-2e0d49df81683c2f4c17d10ccbf442e93488358fc6e01b4a3ec65466f01fd14f): complete subsection reference.

- [forward_proxy_allow_all](resources--azure_vnet_site--reference--group-009.md#canonical-16a46d3bb246d01a83b0bc3affa4640a70d5dab38984cfd513c6a391831bdb6b): complete subsection reference.

- [global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-327c8911a93a55254e55dc87eb61d258e7d2346754cf78398c2d91006457ff76): complete subsection reference.

- [k8s_cluster](resources--azure_vnet_site--reference--group-009.md#canonical-14c24879afe6f88c2393b198b5ab94421491923c38e53c9f09fa4cb65018497d): complete subsection reference.

- [no_dc_cluster_group](resources--azure_vnet_site--reference--group-009.md#canonical-76547a2a65c1ed5de46908fae8cac50833f6550afd877bda4c9b66ff8f59290e): complete subsection reference.

- [no_forward_proxy](resources--azure_vnet_site--reference--group-009.md#canonical-c380970b6fed57ac45623cf12ccf1176250c57e540f5243de98a55766a12ff37): complete subsection reference.

- [no_global_network](resources--azure_vnet_site--reference--group-009.md#canonical-f378eac304375356f64dfcd32abecf41d3ce1971f72504f40cdc9785185a32fe): complete subsection reference.

- [no_k8s_cluster](resources--azure_vnet_site--reference--group-009.md#canonical-c7daa19b41181c03735c133ea6acdc0a062bb50b52c5b4b311cef41867a053b3): complete subsection reference.

- [no_network_policy](resources--azure_vnet_site--reference--group-009.md#canonical-e4321d9db12919ced8fd473de9cb14446f6d2900ea7e79482acbee783cc6d7d7): complete subsection reference.

- [no_outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-974c731bff5f7cf0f7d3800d682825f63ec3910cbc124947f44c4b73ca9494c6): complete subsection reference.

- [node](resources--azure_vnet_site--reference--group-009.md#canonical-fb7ee52f5f58686836f1724c5f129953df3fd12bcd2f2bcbc459c1d30ea4c086): complete subsection reference.

- [outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5): complete subsection reference.

- [sm_connection_public_ip](resources--azure_vnet_site--reference--group-010.md#canonical-21e611a5dd41974704b93d19fc05194f6d2a79ef96057b3871f374955c652813): complete subsection reference.

- [sm_connection_pvt_ip](resources--azure_vnet_site--reference--group-010.md#canonical-4040feef611420346b88754f0042d4c28a7a7349ecea3da7822f3c67e787f512): complete subsection reference.

- [storage_class_list](resources--azure_vnet_site--reference--group-010.md#canonical-ffa39091461c785dd6e267139ed919d5c5313c0a21b41f9fab97d36976d6e5d1): complete subsection reference.

<a id="canonical-b134bddfe97dd49395f5a03a9b387b17d6834334308630837868d51a7f89694d"></a>

## Next pages — voltstack_cluster_ar / ec18bc5827cf / 5

- [voltstack_cluster_ar.accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-c2e78ccf67ee4d4433d5e27c99cf0629a9c984be70493b525f8e5e628236e7cd)
- [voltstack_cluster_ar.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-71f2ecfba05031cb997e73c528066702be86b9746673e6492bec30957476576d)
- [voltstack_cluster_ar.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-fb13dbc8a5f04e6473307f1a618175d7ab9bbc0d5f4db6349e27153924e78314)
- [voltstack_cluster_ar.active_network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-6f7e0b6e15090e23f2a161c319ee5a0198740fd2c07697baed0e5c75405a3fcb)
- [voltstack_cluster_ar.dc_cluster_group](resources--azure_vnet_site--reference--group-009.md#canonical-23445797a2a89420a7e6dbf95e5906ac49988575995da53770d8132f5ef0b561)
- [voltstack_cluster_ar.default_storage](resources--azure_vnet_site--reference--group-009.md#canonical-2e0d49df81683c2f4c17d10ccbf442e93488358fc6e01b4a3ec65466f01fd14f)
- [voltstack_cluster_ar.forward_proxy_allow_all](resources--azure_vnet_site--reference--group-009.md#canonical-16a46d3bb246d01a83b0bc3affa4640a70d5dab38984cfd513c6a391831bdb6b)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-327c8911a93a55254e55dc87eb61d258e7d2346754cf78398c2d91006457ff76)
- [voltstack_cluster_ar.k8s_cluster](resources--azure_vnet_site--reference--group-009.md#canonical-14c24879afe6f88c2393b198b5ab94421491923c38e53c9f09fa4cb65018497d)
- [voltstack_cluster_ar.no_dc_cluster_group](resources--azure_vnet_site--reference--group-009.md#canonical-76547a2a65c1ed5de46908fae8cac50833f6550afd877bda4c9b66ff8f59290e)
- [voltstack_cluster_ar.no_forward_proxy](resources--azure_vnet_site--reference--group-009.md#canonical-c380970b6fed57ac45623cf12ccf1176250c57e540f5243de98a55766a12ff37)
- [voltstack_cluster_ar.no_global_network](resources--azure_vnet_site--reference--group-009.md#canonical-f378eac304375356f64dfcd32abecf41d3ce1971f72504f40cdc9785185a32fe)
- [voltstack_cluster_ar.no_k8s_cluster](resources--azure_vnet_site--reference--group-009.md#canonical-c7daa19b41181c03735c133ea6acdc0a062bb50b52c5b4b311cef41867a053b3)
- [voltstack_cluster_ar.no_network_policy](resources--azure_vnet_site--reference--group-009.md#canonical-e4321d9db12919ced8fd473de9cb14446f6d2900ea7e79482acbee783cc6d7d7)
- [voltstack_cluster_ar.no_outside_static_routes](resources--azure_vnet_site--reference--group-009.md#canonical-974c731bff5f7cf0f7d3800d682825f63ec3910cbc124947f44c4b73ca9494c6)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-009.md#canonical-fb7ee52f5f58686836f1724c5f129953df3fd12bcd2f2bcbc459c1d30ea4c086)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.sm_connection_public_ip](resources--azure_vnet_site--reference--group-010.md#canonical-21e611a5dd41974704b93d19fc05194f6d2a79ef96057b3871f374955c652813)
- [voltstack_cluster_ar.sm_connection_pvt_ip](resources--azure_vnet_site--reference--group-010.md#canonical-4040feef611420346b88754f0042d4c28a7a7349ecea3da7822f3c67e787f512)
- [voltstack_cluster_ar.storage_class_list](resources--azure_vnet_site--reference--group-010.md#canonical-ffa39091461c785dd6e267139ed919d5c5313c0a21b41f9fab97d36976d6e5d1)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c2e78ccf67ee4d4433d5e27c99cf0629a9c984be70493b525f8e5e628236e7cd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0bf013f204aa25dcc0a0fd0f1ca55e986c29f3d255bdb2cc2770fddb55cf344"></a>

## voltstack_cluster_ar.accelerated_networking — voltstack_cluster_ar.accelerated_networking / 538d71edfa1f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.accelerated_networking

<a id="canonical-3fefbaf14eb0f2127b43fe0b6b45c7cb040debbd7e43b597780926dc6d1afef1"></a>

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

<a id="canonical-967a4c0113bf36dbc03ecf374f9155188ece3c8e9407f0c4120785b601f8f8e9"></a>

## Direct properties — voltstack_cluster_ar.accelerated_networking / 538d71edfa1f / 3

- [disable_spec](resources--azure_vnet_site--reference--group-009.md#canonical-fef736fd6059b641a2e1a1cf47cfd7b423c0e623951f930d17a46b323141170a): complete subsection reference.

- [enable](resources--azure_vnet_site--reference--group-009.md#canonical-51864e348b1ee2dc31b764437e15da57f158a9536dbeb8da62e6d5494c83750b): complete subsection reference.

<a id="canonical-f0c63c61ed7e996322c0c456baf72bca406de9b212c724a8b266d76cb24f1965"></a>

## Next pages — voltstack_cluster_ar.accelerated_networking / 538d71edfa1f / 4

- [voltstack_cluster_ar.accelerated_networking.disable_spec](resources--azure_vnet_site--reference--group-009.md#canonical-fef736fd6059b641a2e1a1cf47cfd7b423c0e623951f930d17a46b323141170a)
- [voltstack_cluster_ar.accelerated_networking.enable](resources--azure_vnet_site--reference--group-009.md#canonical-51864e348b1ee2dc31b764437e15da57f158a9536dbeb8da62e6d5494c83750b)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-fef736fd6059b641a2e1a1cf47cfd7b423c0e623951f930d17a46b323141170a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1a22939a83827cd118eb4518042bdbc739700528a2594a7d3ffcc54b6fd14bc"></a>

## voltstack_cluster_ar.accelerated_networking.disable_spec — voltstack_cluster_ar.accelerated_networking.disable_spec / dd49b0c9ce30 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-c2e78ccf67ee4d4433d5e27c99cf0629a9c984be70493b525f8e5e628236e7cd)
- voltstack_cluster_ar.accelerated_networking.disable_spec

<a id="canonical-ee5ee8d0cf344d6faae0fff554749a5e494579ef033844d71bbbc0d35560230f"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-fa387b34fcc4114f7793953e4bec91c79ac2a686513b118d67240b5b43c8de04"></a>

## Direct properties — voltstack_cluster_ar.accelerated_networking.disable_spec / dd49b0c9ce30 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-646df622d254299290c562cff930b8dfe05eb8abe6fc8d7786b5496cffa54bf9"></a>

## Next pages — voltstack_cluster_ar.accelerated_networking.disable_spec / dd49b0c9ce30 / 4

- [voltstack_cluster_ar.accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-c2e78ccf67ee4d4433d5e27c99cf0629a9c984be70493b525f8e5e628236e7cd)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-51864e348b1ee2dc31b764437e15da57f158a9536dbeb8da62e6d5494c83750b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c28e6ab96c90d6c75b2c89668308ae8b480c2c81518e900b3e87daa9d0dc23e"></a>

## voltstack_cluster_ar.accelerated_networking.enable — voltstack_cluster_ar.accelerated_networking.enable / 6cdbb88b28c8 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-c2e78ccf67ee4d4433d5e27c99cf0629a9c984be70493b525f8e5e628236e7cd)
- voltstack_cluster_ar.accelerated_networking.enable

<a id="canonical-dda1f0067105fcbf9e43858ad23e0b02f81fd21a39ea899167deac86ba6609e7"></a>

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

<a id="canonical-a018af84950a7ca0775e91067007d706e2c0bc1cd5d1092c7e4ed7966db6eeb0"></a>

## Direct properties — voltstack_cluster_ar.accelerated_networking.enable / 6cdbb88b28c8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f2158ac99a09c7383962c245aee1f8e354cbdef26857b89a5ddf33a1c81d9765"></a>

## Next pages — voltstack_cluster_ar.accelerated_networking.enable / 6cdbb88b28c8 / 4

- [voltstack_cluster_ar.accelerated_networking](resources--azure_vnet_site--reference--group-009.md#canonical-c2e78ccf67ee4d4433d5e27c99cf0629a9c984be70493b525f8e5e628236e7cd)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-71f2ecfba05031cb997e73c528066702be86b9746673e6492bec30957476576d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9604b205d848e029c03dae80eaf82b78f8e29f9bd6bce27d22ad0d111170f554"></a>

## voltstack_cluster_ar.active_enhanced_firewall_policies — voltstack_cluster_ar.active_enhanced_firewall_policies / fd0e9951e5ab / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.active_enhanced_firewall_policies

<a id="canonical-0783fec0b1c7d3996814b7012e713036802b8faef185d2be0cd31b81f88db7e6"></a>

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

<a id="canonical-c53031d98bc4f73186bed45a39fb2e04e827e8ad957074b34fcbd0bb2614a4ad"></a>

## Direct properties — voltstack_cluster_ar.active_enhanced_firewall_policies / fd0e9951e5ab / 3

- [enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-7fc1447097140396242b7696806679d48f99c72c4b5375952f4918402330711a): complete subsection reference.

<a id="canonical-7ed0b36618ca84af973bdf6a1a41cb4a036b8cd7dcfb8d02ab944eae937e9fbf"></a>

## Next pages — voltstack_cluster_ar.active_enhanced_firewall_policies / fd0e9951e5ab / 4

- [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-7fc1447097140396242b7696806679d48f99c72c4b5375952f4918402330711a)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-7fc1447097140396242b7696806679d48f99c72c4b5375952f4918402330711a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cec0fae3e55c588aa66240a8fa86dcd0a5c4a8174a8c559a127dfb5be8d6153c"></a>

## voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / 6f88a4f8e169 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-71f2ecfba05031cb997e73c528066702be86b9746673e6492bec30957476576d)
- voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-1e97037614f96905e6440b655c336170adcc6ae409ed73e2b871166dd9bf764d"></a>

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

<a id="canonical-88c50e109723393ebad13a31fb8692975111d6429f40f33a4d49c39885673a2b"></a>

## Direct properties — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / 6f88a4f8e169 / 3

<a id="canonical-555ac37edd2f5c66f87b381ab751bbb6d0ecf828a512e50f851db21db3ff5138"></a>

<a id="canonical-852decf780868603b25b67ff38683f3a61bf46a4b0ce0cbab29556bbc5f031bb"></a>

## name property — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / 6f88a4f8e169 / 4

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

<a id="canonical-6731f8b7fcba7889829c2058bec5a6591dc2c764bed8972621dc61bffd74176f"></a>

<a id="canonical-db124a0e9ac71b39f288ad8818486b213b0498414c224752d38a6025c59c9896"></a>

## namespace property — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / 6f88a4f8e169 / 5

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

<a id="canonical-304d88b529706b3b340c8063ff47a81257bced9e2c842367f5734a91d9a6d093"></a>

<a id="canonical-01ecea35a43debc72584ba4897c2db31fd9ddf4f7b26684b318553a3aabe91a8"></a>

## tenant property — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / 6f88a4f8e169 / 6

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

<a id="canonical-ddc92caa40a1291a124e2cb64332d74d77a8baa85bfa460393e181d976bb37a1"></a>

## Next pages — voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policie / 6f88a4f8e169 / 7

- [voltstack_cluster_ar.active_enhanced_firewall_policies](resources--azure_vnet_site--reference--group-009.md#canonical-71f2ecfba05031cb997e73c528066702be86b9746673e6492bec30957476576d)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-fb13dbc8a5f04e6473307f1a618175d7ab9bbc0d5f4db6349e27153924e78314"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ac25ea7d982bc80d524e7d9712933436ee4c98b3cd33def8b358fee176d12ae"></a>

## voltstack_cluster_ar.active_forward_proxy_policies — voltstack_cluster_ar.active_forward_proxy_policies / 85f92e13f19f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.active_forward_proxy_policies

<a id="canonical-1b2b4f8a838912ebf95cfae7bab468783d955e120555a6b27d16b0bb226078bc"></a>

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

<a id="canonical-157e80529eac254bb61da8dd32cc5ea346cb692873a13d4731449ff517c44673"></a>

## Direct properties — voltstack_cluster_ar.active_forward_proxy_policies / 85f92e13f19f / 3

- [forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-7826963610e69237cf999027c2bb4df3641cf510565b13dcee9392e962cf370b): complete subsection reference.

<a id="canonical-53238fc460d6220f4b420790846cbceffaded6b9b37766698131cd5a1709ad91"></a>

## Next pages — voltstack_cluster_ar.active_forward_proxy_policies / 85f92e13f19f / 4

- [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-7826963610e69237cf999027c2bb4df3641cf510565b13dcee9392e962cf370b)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-7826963610e69237cf999027c2bb4df3641cf510565b13dcee9392e962cf370b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abd26ad6be7679e5d6c4d83c642af66a9f2adf63bd70e56d7b795b959d736f50"></a>

## voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / a36f097e112c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-fb13dbc8a5f04e6473307f1a618175d7ab9bbc0d5f4db6349e27153924e78314)
- voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-9ec5bda6a245d5a3826b8a2e2f7fc06723184a60d3ff34df921362954778fadd"></a>

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

<a id="canonical-95b46fdda71607582e37f84780dcae99585922fb35f742db6c9ceec35793ddf4"></a>

## Direct properties — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / a36f097e112c / 3

<a id="canonical-2163f95d49bea372727d6c842812b79bb2c66364c8bd37598cf357efb66f9aa5"></a>

<a id="canonical-855660f237dda4794175a0d4e580afd04d5915f5fecebefe8cd05c6be970960f"></a>

## name property — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / a36f097e112c / 4

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

<a id="canonical-74f6469329285271c56c87d53aa30387f8efdd16edd3947707c3dc7aebc591a2"></a>

<a id="canonical-10df54f951581ecb2d7469c6b9d4a02d277003ccbbb2e69a9f22653da226337b"></a>

## namespace property — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / a36f097e112c / 5

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

<a id="canonical-6363f87476820a9c8c3c8314ea15626cc0d94cd9b26f04bd68babcdc15762b9a"></a>

<a id="canonical-b13bafaff6691719c6a18a6970bcb87666b6589a39ef3ee3f8232ec6c25b71ab"></a>

## tenant property — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / a36f097e112c / 6

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

<a id="canonical-b006b6b5f8338e4a576ac3102de17d3be84eb14d243424b6ec69845fd1ebcc6d"></a>

## Next pages — voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies / a36f097e112c / 7

- [voltstack_cluster_ar.active_forward_proxy_policies](resources--azure_vnet_site--reference--group-009.md#canonical-fb13dbc8a5f04e6473307f1a618175d7ab9bbc0d5f4db6349e27153924e78314)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6f7e0b6e15090e23f2a161c319ee5a0198740fd2c07697baed0e5c75405a3fcb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2b9247c27a00d30d0ad56c98404458284345a84aea307dbe5bceb4fc3b8090da"></a>

## voltstack_cluster_ar.active_network_policies — voltstack_cluster_ar.active_network_policies / 4018cd44423b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.active_network_policies

<a id="canonical-2cac05aab74ec005cecd00607216922fce88f78f238c7ebe9382ca4fd63028fd"></a>

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

<a id="canonical-5605321ec0963a1915e8c11142aed8973659db1eb823aeb6ebc0c34672fba881"></a>

## Direct properties — voltstack_cluster_ar.active_network_policies / 4018cd44423b / 3

- [network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-87e4c14f064862f2bc144f134c8366f3dc4685d3c51054d0d82d03700331a193): complete subsection reference.

<a id="canonical-894e8248af216ce3bf5a5622a8eca3d73470ef2db4a17866594826f2fd0f004e"></a>

## Next pages — voltstack_cluster_ar.active_network_policies / 4018cd44423b / 4

- [voltstack_cluster_ar.active_network_policies.network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-87e4c14f064862f2bc144f134c8366f3dc4685d3c51054d0d82d03700331a193)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-87e4c14f064862f2bc144f134c8366f3dc4685d3c51054d0d82d03700331a193"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2e94f363f1af1d9ab8b76ce318b14530bda015e497c4d2c556243ab69620d425"></a>

## voltstack_cluster_ar.active_network_policies.network_policies — voltstack_cluster_ar.active_network_policies.network_policies / c215d6494e63 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.active_network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-6f7e0b6e15090e23f2a161c319ee5a0198740fd2c07697baed0e5c75405a3fcb)
- voltstack_cluster_ar.active_network_policies.network_policies

<a id="canonical-60a33093e56d03196c6be6c0a038c4f42da9213e2e38a95c3c2c2194837e463d"></a>

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

<a id="canonical-eb4faf26837ff2afbbfe5aed1d414ac69b424c130612fd1fac04eb15f29636e1"></a>

## Direct properties — voltstack_cluster_ar.active_network_policies.network_policies / c215d6494e63 / 3

<a id="canonical-d238e428110e185d567fcbc67207f3eff3035f018caeb8599cb9faa619c73efe"></a>

<a id="canonical-810899e3f836c4139fba61c2062381ad029bbff324b325537cc7aeb4217fbcc1"></a>

## name property — voltstack_cluster_ar.active_network_policies.network_policies / c215d6494e63 / 4

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

<a id="canonical-34811168311edf1fca10917208c990c6ae3cf34782b5aeb89e5adadc384cd473"></a>

<a id="canonical-965e73ef1fc86dd2036f0f550ef133b95cabb29bed40a3f67a72c3e0f76e5021"></a>

## namespace property — voltstack_cluster_ar.active_network_policies.network_policies / c215d6494e63 / 5

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

<a id="canonical-87c482d041acb2568e885033a9e813634ab7806cf268e8231b3dcc60dd471e72"></a>

<a id="canonical-e6f6268dfecc705a131227d8718814d3393d0ff283b6652501e722871e625c69"></a>

## tenant property — voltstack_cluster_ar.active_network_policies.network_policies / c215d6494e63 / 6

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

<a id="canonical-f839d2ccfd4a6f036d04cdec4211a04d5ebb20ad11d94623cac2eb77149e6f12"></a>

## Next pages — voltstack_cluster_ar.active_network_policies.network_policies / c215d6494e63 / 7

- [voltstack_cluster_ar.active_network_policies](resources--azure_vnet_site--reference--group-009.md#canonical-6f7e0b6e15090e23f2a161c319ee5a0198740fd2c07697baed0e5c75405a3fcb)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-23445797a2a89420a7e6dbf95e5906ac49988575995da53770d8132f5ef0b561"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c43b5b3c550da5fd74756b56d20025a9aca02d4d88de89ea851c1a38a96dc19e"></a>

## voltstack_cluster_ar.dc_cluster_group — voltstack_cluster_ar.dc_cluster_group / 40f00666e860 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.dc_cluster_group

<a id="canonical-71999cc3e0ae4a23424dd911df661d69adec8038d145ac49c0a499dcc774896b"></a>

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
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-3e13151c046fce302becdb999eaaaea80c7ef4bb65e01e54e672318f5db962ab"></a>

## Direct properties — voltstack_cluster_ar.dc_cluster_group / 40f00666e860 / 3

<a id="canonical-de1902c7da977375f6ceb818af69465e3d866895cad615fc02afe13695f0728f"></a>

<a id="canonical-d2592a3e0bb769fd3d8143463953d6e16d0ae9282caeb736f0acc40a61159830"></a>

## name property — voltstack_cluster_ar.dc_cluster_group / 40f00666e860 / 4

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

<a id="canonical-4f66978ed500b8de63617336f156a30ef17bd82e6d564e1f2ad8a9f12898f147"></a>

<a id="canonical-2230bf5927c3fc032ba6e0dbcedf8ba5d1e69c729fc6fc06d96ffbed98c16bcf"></a>

## namespace property — voltstack_cluster_ar.dc_cluster_group / 40f00666e860 / 5

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

<a id="canonical-2c8ac8c9bdf9e3b3c1dd4272db4991642737d17115f839cbca12ff5eb71a96f6"></a>

<a id="canonical-6ef17499508d79e9787ac0dc37757ae8f18ff2cf4ae760469c75ff7715a1284f"></a>

## tenant property — voltstack_cluster_ar.dc_cluster_group / 40f00666e860 / 6

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

<a id="canonical-01f875eeea48f4c37899127d8536135a728d0b26233d680939263a90886204f1"></a>

## Next pages — voltstack_cluster_ar.dc_cluster_group / 40f00666e860 / 7

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-2e0d49df81683c2f4c17d10ccbf442e93488358fc6e01b4a3ec65466f01fd14f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5ece1c1f18120005fa10200cb78f80cb49ad52967e51d4c301e403b1a8c5207"></a>

## voltstack_cluster_ar.default_storage — voltstack_cluster_ar.default_storage / 57c31e3c6254 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.default_storage

<a id="canonical-e7c28727162cec38b1aa77bc03dbbefb436e72e35e062a7b6b7c6725fcb2f6dd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default storage.

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
default_storage = {}
```

<a id="canonical-45dbd7a44d849ed1dc31f3a698b29e427498a8aa8d4ff5456c47cd7bbf63fee4"></a>

## Direct properties — voltstack_cluster_ar.default_storage / 57c31e3c6254 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9a2a9ae9f1a46d066d1291f31a79220f795d19a7bc6e7ab28cc9ff03610700ce"></a>

## Next pages — voltstack_cluster_ar.default_storage / 57c31e3c6254 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-16a46d3bb246d01a83b0bc3affa4640a70d5dab38984cfd513c6a391831bdb6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8c0b22eedcd58d2eda2ffd463575db1e6d88b4a91e17a8234bc54b5ef7236939"></a>

## voltstack_cluster_ar.forward_proxy_allow_all — voltstack_cluster_ar.forward_proxy_allow_all / d9970fe74654 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.forward_proxy_allow_all

<a id="canonical-4f0a077184589fa7ff0c47ecb69be3071dfb48edb4e010f99ea4357cf86cefe1"></a>

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

<a id="canonical-16adf182a29462033b27f662f5997b47592f9bf5ac14eaaa3a0efbc2fee37872"></a>

## Direct properties — voltstack_cluster_ar.forward_proxy_allow_all / d9970fe74654 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fea3bc290061917de5b784d2627f3567602a07de2bff7e3f168af83c775c249b"></a>

## Next pages — voltstack_cluster_ar.forward_proxy_allow_all / d9970fe74654 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-327c8911a93a55254e55dc87eb61d258e7d2346754cf78398c2d91006457ff76"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5f778f0ba8990137d1f863fa02e787c47a2ee34bc5b5586faa90b0921f50394e"></a>

## voltstack_cluster_ar.global_network_list — voltstack_cluster_ar.global_network_list / 21ce4fb73c64 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.global_network_list

<a id="canonical-0a68f701e0fc5e9601277d01d2cf43ee37a145faf6749de6dd21a2df51a060c3"></a>

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

<a id="canonical-1fac3a70372d7edcd2339623904edf6522014e9f74ad484fa8f7ef3a565a3388"></a>

## Direct properties — voltstack_cluster_ar.global_network_list / 21ce4fb73c64 / 3

- [global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-fd03faf3d5158ec579df2f62f3e392fea141b1f3053776557e2daedcdbc770b6): complete subsection reference.

<a id="canonical-c10933b3af769323e2d5ea4b895b305b4742f6dd116ca2598480116711c595f4"></a>

## Next pages — voltstack_cluster_ar.global_network_list / 21ce4fb73c64 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-fd03faf3d5158ec579df2f62f3e392fea141b1f3053776557e2daedcdbc770b6)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-fd03faf3d5158ec579df2f62f3e392fea141b1f3053776557e2daedcdbc770b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9840f4717519ce71d6c3c1050b5ffbf5c78b643da844556fd169a3bfa4e18b1"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections — voltstack_cluster_ar.global_network_list.global_network_connections / 1298bff3ef3e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-327c8911a93a55254e55dc87eb61d258e7d2346754cf78398c2d91006457ff76)
- voltstack_cluster_ar.global_network_list.global_network_connections

<a id="canonical-43b4da316476b2ecb92b3e0381e018b5392be77d391d8392eb1e5c0851ef7e88"></a>

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

<a id="canonical-5f6ccf8ad5e4a068b43ce7a5d6d361a7c243d664ca2c6b92f3e7d6b2ff5eab86"></a>

## Direct properties — voltstack_cluster_ar.global_network_list.global_network_connections / 1298bff3ef3e / 3

- [sli_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-44aab3378318fb18d4fff04c492db69d66074e3abd321bdb29e0722b49fc4039): complete subsection reference.

- [slo_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-9c283b5da85a192ef35de6140729ed1d00f102e53f5340f54dca917cbaa3c2bc): complete subsection reference.

<a id="canonical-839600face50b09e2e95457457e2a05c59fffb8d8a2b625104de501b0d952a05"></a>

## Next pages — voltstack_cluster_ar.global_network_list.global_network_connections / 1298bff3ef3e / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-44aab3378318fb18d4fff04c492db69d66074e3abd321bdb29e0722b49fc4039)
- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-9c283b5da85a192ef35de6140729ed1d00f102e53f5340f54dca917cbaa3c2bc)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-327c8911a93a55254e55dc87eb61d258e7d2346754cf78398c2d91006457ff76)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-44aab3378318fb18d4fff04c492db69d66074e3abd321bdb29e0722b49fc4039"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0b0118133509f6bb92f5201d0612f5eb823d6f9571632f2033dd2fea40c67bd7"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 3f35afd62509 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-327c8911a93a55254e55dc87eb61d258e7d2346754cf78398c2d91006457ff76)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-fd03faf3d5158ec579df2f62f3e392fea141b1f3053776557e2daedcdbc770b6)
- voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-e427bff3f6da567be9b2939267475c16c3fb82468a4e00935e5403ac304dc2c6"></a>

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

<a id="canonical-6e9e3897dadddc569bf804de742243b4b0d500cffa31175589d8066c65e7711d"></a>

## Direct properties — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 3f35afd62509 / 3

- [global_vn](resources--azure_vnet_site--reference--group-009.md#canonical-205682214a8edb0c8d3013752d9b9c12c936c61eb8d7239c5f63286488dfa912): complete subsection reference.

<a id="canonical-2e452a338afd21f40bc592524210b76c52ecf3a31528283df6a9d9db693276ea"></a>

## Next pages — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 3f35afd62509 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--azure_vnet_site--reference--group-009.md#canonical-205682214a8edb0c8d3013752d9b9c12c936c61eb8d7239c5f63286488dfa912)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-fd03faf3d5158ec579df2f62f3e392fea141b1f3053776557e2daedcdbc770b6)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-205682214a8edb0c8d3013752d9b9c12c936c61eb8d7239c5f63286488dfa912"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cfe6daccb9ba85785a5368042f68f6b91703ec49e964db777e9fb5a7b3ab0c5b"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 1c79ba7977ce / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-327c8911a93a55254e55dc87eb61d258e7d2346754cf78398c2d91006457ff76)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-fd03faf3d5158ec579df2f62f3e392fea141b1f3053776557e2daedcdbc770b6)
- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-44aab3378318fb18d4fff04c492db69d66074e3abd321bdb29e0722b49fc4039)
- voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-0cf0f1d7934524e4a6078dceb1f0e2932a664dca7431302510b0258e0067d5a2"></a>

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

<a id="canonical-6866cec1cdceaa1fe5e069f41785d84de90cf46c49d82824cbac6e590edaa889"></a>

## Direct properties — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 1c79ba7977ce / 3

<a id="canonical-ce23f0a7c62ec76439692c83891ea7baa2737741b6afc589d3be95c7a02b199d"></a>

<a id="canonical-ba6cc2fc60a516bc3f0638c04ff8eaf7dc6d676cc54011efdf77971801950928"></a>

## name property — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 1c79ba7977ce / 4

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

<a id="canonical-b65d64222da9aa09520a95e8287ad74489081e7157e0057cb1f0b191f9fa78e4"></a>

<a id="canonical-fd006c6613965cc8484aa08250cc040c3acf7e18c91e7e283af7cb9364e5cc3f"></a>

## namespace property — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 1c79ba7977ce / 5

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

<a id="canonical-de50dda95c92a9928d58346b7aa241a59b5d4a68c49f6633d8bf6d0d860ddf41"></a>

<a id="canonical-ac3b8bd9848e63aac3bd3dde24e2345392a7619d6cec0d306d0473b1af2eefd4"></a>

## tenant property — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 1c79ba7977ce / 6

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

<a id="canonical-76f7159f4d0206cc89485f560aad4fce8074aed37c3a777adfe4432074916ba7"></a>

## Next pages — voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_globa / 1c79ba7977ce / 7

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-44aab3378318fb18d4fff04c492db69d66074e3abd321bdb29e0722b49fc4039)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9c283b5da85a192ef35de6140729ed1d00f102e53f5340f54dca917cbaa3c2bc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcff6dc94d0e1b32d4b2a4d9e066a1097f7928d941f534dc2685e2a4110e2848"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 4d2c529401d9 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-327c8911a93a55254e55dc87eb61d258e7d2346754cf78398c2d91006457ff76)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-fd03faf3d5158ec579df2f62f3e392fea141b1f3053776557e2daedcdbc770b6)
- voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-d3b6a373392656927b9984a2e3c73f0dda21e5ed0e06afdc7b503bd397e14f90"></a>

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

<a id="canonical-5338434fcdcde509efc74dfb07fc9eb0f0a1a780a901dd9b9d32bc71b0e5a240"></a>

## Direct properties — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 4d2c529401d9 / 3

- [global_vn](resources--azure_vnet_site--reference--group-009.md#canonical-0db0bc3b1a94c1a3820eac75ec1ed58e030c02b88a747edca000f523fe29c862): complete subsection reference.

<a id="canonical-cc6e1cfd8dedb6e7862cf40d47e58d389a42840fd316ce3dc95a5635190e26b5"></a>

## Next pages — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 4d2c529401d9 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--azure_vnet_site--reference--group-009.md#canonical-0db0bc3b1a94c1a3820eac75ec1ed58e030c02b88a747edca000f523fe29c862)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-fd03faf3d5158ec579df2f62f3e392fea141b1f3053776557e2daedcdbc770b6)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-0db0bc3b1a94c1a3820eac75ec1ed58e030c02b88a747edca000f523fe29c862"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbd9d30e196886d688b5c0ef069e477b6d55f3b8aff92f60d057301b40bbc1c9"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 2a40c51718e3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-327c8911a93a55254e55dc87eb61d258e7d2346754cf78398c2d91006457ff76)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-009.md#canonical-fd03faf3d5158ec579df2f62f3e392fea141b1f3053776557e2daedcdbc770b6)
- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-9c283b5da85a192ef35de6140729ed1d00f102e53f5340f54dca917cbaa3c2bc)
- voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-cfe6df1029d6be6afe306b34cb798981c5f946cacce8f95c3595c8d100c8f86e"></a>

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

<a id="canonical-aba79124678a0bca69c91711fc2debdeac83a84f4781dba319ec78ff9fd8ae27"></a>

## Direct properties — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 2a40c51718e3 / 3

<a id="canonical-3a97f21be3a27becb7d9325c851da5a0b762db7cfa5215abcddec3ace2bf881b"></a>

<a id="canonical-aaa6d770068ec682fc820d7d1d9ba62fd33c486ffc27c05c3ca08db43c6cd360"></a>

## name property — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 2a40c51718e3 / 4

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

<a id="canonical-7d9602b2e3e2c2446cba58427caf2d5585d96c27307eabd7899a582e46e23352"></a>

<a id="canonical-f487cb538c3460099bdb0385365ae7dc80e32f7b5604fd515a3558b59413d320"></a>

## namespace property — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 2a40c51718e3 / 5

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

<a id="canonical-c9ccf484fa32322e9f61564567f03e322c40cc06c6639d2b7386948850b599af"></a>

<a id="canonical-97e87c6d69e41ea6cc098f8a8550d5ee7222ce846d0daeb0f046dde1f94c32a9"></a>

## tenant property — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 2a40c51718e3 / 6

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

<a id="canonical-467742e4ab33be55938ced712175cd52499fdedb53f9ea370388f356085aa654"></a>

## Next pages — voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_globa / 2a40c51718e3 / 7

- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-009.md#canonical-9c283b5da85a192ef35de6140729ed1d00f102e53f5340f54dca917cbaa3c2bc)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-14c24879afe6f88c2393b198b5ab94421491923c38e53c9f09fa4cb65018497d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16ff590670f2ed051cf434c4a7d99ca0a8353fabf85d449da3a6dda1b277a40a"></a>

## voltstack_cluster_ar.k8s_cluster — voltstack_cluster_ar.k8s_cluster / e69cbfc21bb3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.k8s_cluster

<a id="canonical-cd372866fbf67013c723f3154d0dae2df39eb2093cc6bea93d306ed98096a2d2"></a>

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
k8s_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-e13c59c99d0d35118c58281255c785edc3ecfe299c286db32580fa16c0da8b9e"></a>

## Direct properties — voltstack_cluster_ar.k8s_cluster / e69cbfc21bb3 / 3

<a id="canonical-c54e803b653b0faf28ab8c9c1422239468a4d4b0f12e196954e73c416b9b06de"></a>

<a id="canonical-f645d8e5825c9a692a3b00e98af181ba0461f2d2f516e3c57a898e5de5ad1ef5"></a>

## name property — voltstack_cluster_ar.k8s_cluster / e69cbfc21bb3 / 4

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

<a id="canonical-b9d82fe944b20edfb89566beae832193805d36dfee9b0070f639bfe65edfbd00"></a>

<a id="canonical-39688e123e42aa63c1110fced01ba027d22a0417fe20fc2151c2fd87039852e5"></a>

## namespace property — voltstack_cluster_ar.k8s_cluster / e69cbfc21bb3 / 5

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

<a id="canonical-28f83ffe8aa8fd18cec3da11b058e5812cf5ba36b61b4dfa2eb6dea275df19aa"></a>

<a id="canonical-b5e27723313ea5af25af6a289ed57c9cfc0a366aa2fe6fc05ecb0da858c463c0"></a>

## tenant property — voltstack_cluster_ar.k8s_cluster / e69cbfc21bb3 / 6

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

<a id="canonical-bfecd537394878c31ef7248804269409ef47b74d3efb78a504da9861a3907208"></a>

## Next pages — voltstack_cluster_ar.k8s_cluster / e69cbfc21bb3 / 7

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-76547a2a65c1ed5de46908fae8cac50833f6550afd877bda4c9b66ff8f59290e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0d96dcd076448e50fd2a3cd6cf4b5e50233b2f72ecbaffa3016d704a423791a6"></a>

## voltstack_cluster_ar.no_dc_cluster_group — voltstack_cluster_ar.no_dc_cluster_group / 63aac84391bb / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.no_dc_cluster_group

<a id="canonical-0da07f89d3dd42a0cd445645767b0ea314087513a8af3452a8f6955c90ed8b1d"></a>

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

<a id="canonical-0bd0fa31543eb42cfdc03e7d283227881c88c8a6929374d79094a140957ec779"></a>

## Direct properties — voltstack_cluster_ar.no_dc_cluster_group / 63aac84391bb / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ca6ebd6efd77e475e139b6ac494c09ecb9d558e18808650157587822dc53319c"></a>

## Next pages — voltstack_cluster_ar.no_dc_cluster_group / 63aac84391bb / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c380970b6fed57ac45623cf12ccf1176250c57e540f5243de98a55766a12ff37"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8ffc139675ce35baf56503835b4ba7439a42bd8626db4b88fa26429753b5657e"></a>

## voltstack_cluster_ar.no_forward_proxy — voltstack_cluster_ar.no_forward_proxy / e041b23d09e0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.no_forward_proxy

<a id="canonical-6ad2f3ca5a1a6221d13d4066024dc949e84e8385f82de527e8f4caff4fd6a0da"></a>

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

<a id="canonical-643820ded1a563c67bac36e7455bc55e3d2e5419cd8d5bd75c94afcb5261b766"></a>

## Direct properties — voltstack_cluster_ar.no_forward_proxy / e041b23d09e0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a69897fd520b29073590fceee9b383b09d6185b4e5611c7de2f12924872a188"></a>

## Next pages — voltstack_cluster_ar.no_forward_proxy / e041b23d09e0 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-f378eac304375356f64dfcd32abecf41d3ce1971f72504f40cdc9785185a32fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c742b6f6a65ecdb608024dd078c342aa2a853302cded112e41a0708e194f0c8e"></a>

## voltstack_cluster_ar.no_global_network — voltstack_cluster_ar.no_global_network / 7ab1f119e594 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.no_global_network

<a id="canonical-990ee95ba201ec43bd4a84c08b9ea69a8d7543cadc562abb45d409097dd1943d"></a>

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

<a id="canonical-7aec33f7b3600153d1cb95ea10c4754f1e365a904fba43e14c264c33eba130af"></a>

## Direct properties — voltstack_cluster_ar.no_global_network / 7ab1f119e594 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ff12d9591f9bfc89791b0a341d0ee0c8f153d4f6e6c1c294267947f241866577"></a>

## Next pages — voltstack_cluster_ar.no_global_network / 7ab1f119e594 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c7daa19b41181c03735c133ea6acdc0a062bb50b52c5b4b311cef41867a053b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-68478f2ab0dd583bf8c8582eb2283505f73ccac3fd0085441ea5edc13b6cd021"></a>

## voltstack_cluster_ar.no_k8s_cluster — voltstack_cluster_ar.no_k8s_cluster / 3452c6c125b7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.no_k8s_cluster

<a id="canonical-26ede9fe577314e7f729776b623a7fd98c0c1fd9d982f5573d2762002c7a4870"></a>

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
no_k8s_cluster = {}
```

<a id="canonical-25ae3ab3718f6f7b6ffd4baf4fa13c756e1f337175ae83dd0b8aaf066072597b"></a>

## Direct properties — voltstack_cluster_ar.no_k8s_cluster / 3452c6c125b7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-40dc4af16693ff50b325986b966e264d2d9efc79c43c5adc6a49153247f8a71d"></a>

## Next pages — voltstack_cluster_ar.no_k8s_cluster / 3452c6c125b7 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e4321d9db12919ced8fd473de9cb14446f6d2900ea7e79482acbee783cc6d7d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-226d407715b4ca419af99e8b5dc00fc9c2e8741aac5d5bb0f3276be3f79db15c"></a>

## voltstack_cluster_ar.no_network_policy — voltstack_cluster_ar.no_network_policy / f97c33dd0532 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.no_network_policy

<a id="canonical-3aec6272abaa1827f626c30a73911be469fccf7f2eeca11364de26c44f382806"></a>

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

<a id="canonical-060dfd835565665e39bef975bd69de45dc5004561f4609b2c480cec4edd6e6dd"></a>

## Direct properties — voltstack_cluster_ar.no_network_policy / f97c33dd0532 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2bf1e2685f1e9c259ae366a889b1eb945e49397170b2ce97fb24d5945eae86cc"></a>

## Next pages — voltstack_cluster_ar.no_network_policy / f97c33dd0532 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-974c731bff5f7cf0f7d3800d682825f63ec3910cbc124947f44c4b73ca9494c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d811d77fc6144149d081928694177aedeb1a374d30eff3aabccb6d4cf610663"></a>

## voltstack_cluster_ar.no_outside_static_routes — voltstack_cluster_ar.no_outside_static_routes / 8055ba8cc48d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.no_outside_static_routes

<a id="canonical-2b8f194b7445c5ea98b930cd0c95505a10ee9c34d387e1fbe811e83d6871cded"></a>

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

<a id="canonical-4897e81d007b3d2187be2c306c0f6d0e42c726ea9891ad8139033cb012a84dd7"></a>

## Direct properties — voltstack_cluster_ar.no_outside_static_routes / 8055ba8cc48d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1618860c55c9e4e14711cca11e01dd7141ecf41a61a4fe11084e93731ef20049"></a>

## Next pages — voltstack_cluster_ar.no_outside_static_routes / 8055ba8cc48d / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-fb7ee52f5f58686836f1724c5f129953df3fd12bcd2f2bcbc459c1d30ea4c086"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-96cc5922a6d869e5eda3cdf380636ad8b3fe02daba57ead2507624904f79eb85"></a>

## voltstack_cluster_ar.node — voltstack_cluster_ar.node / c88c27790607 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.node

<a id="canonical-568c633908f77f4f23ed91c365a7b8b8daf07f10069fbd947b2ee69df5e1d648"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating Single interface Node for Alternate Region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fault_domain",
    "node_number",
    "update_domain")}
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
node {
  # Configure direct properties listed below.
}
```

<a id="canonical-f2a73341807e69ee8754e1d6ce164d793f23d437a144184e9288507dd5f33969"></a>

## Direct properties — voltstack_cluster_ar.node / c88c27790607 / 3

<a id="canonical-29c6bab058b2bac3e0aad54eae5a8ed34cca8799e8cbf432f0cd9de5d06203d1"></a>

<a id="canonical-4c42f004842966916f0a385192530211935c5ef7f565b7bf762aceff6a8aa198"></a>

## fault_domain property — voltstack_cluster_ar.node / c88c27790607 / 4

Type: `"number"`. Optional.

Namuber of fault domains to be used while creating the availability set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 3),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  }
}
```

- [local_subnet](resources--azure_vnet_site--reference--group-009.md#canonical-bef5d1ec441f891737be5a113eb09bf7ef6dbfab5b5c84741f15bc0cc96227ab): complete subsection reference.

<a id="canonical-3d50066966e391b3040da5a6de1839f21ecd4b34554f577cee81fedc93c8ef8f"></a>

<a id="canonical-82e5dbfa2b0ce07340dca0c983f2f2aeb66a5fbc589d8e1473e39ca9b8b675fa"></a>

## node_number property — voltstack_cluster_ar.node / c88c27790607 / 5

Type: `"number"`. Optional.

Number of main nodes to create, either 1 or 3.

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
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

<a id="canonical-1af663558a25e4d5699900d5057f30f35c78afde5f21d319e5201533e318a867"></a>

<a id="canonical-6189829155c55188e5e54a735180b309ca7f4ca67271f6137810644f5b42f6d6"></a>

## update_domain property — voltstack_cluster_ar.node / c88c27790607 / 6

Type: `"number"`. Optional.

Namuber of update domains to be used while creating the availability set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-d66c2d07f9493ef60d030b9788e94e22ada860bfb12ff4dd610bedd7666a5940"></a>

## Next pages — voltstack_cluster_ar.node / c88c27790607 / 7

- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-009.md#canonical-bef5d1ec441f891737be5a113eb09bf7ef6dbfab5b5c84741f15bc0cc96227ab)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-bef5d1ec441f891737be5a113eb09bf7ef6dbfab5b5c84741f15bc0cc96227ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fdfb440c8f3f42126c2414d6cec4886557fbf719444e8139978d13222599fe57"></a>

## voltstack_cluster_ar.node.local_subnet — voltstack_cluster_ar.node.local_subnet / 731b3af91b0e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-009.md#canonical-fb7ee52f5f58686836f1724c5f129953df3fd12bcd2f2bcbc459c1d30ea4c086)
- voltstack_cluster_ar.node.local_subnet

<a id="canonical-f9f657fa0ed0a19451ff145b78cd906ddca66fa00d405bcbf239d90935dc8a25"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for local subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
local_subnet {
  # Configure direct properties listed below.
}
```
