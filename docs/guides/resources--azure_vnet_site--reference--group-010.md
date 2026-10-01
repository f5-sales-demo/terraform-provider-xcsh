---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-89709196c8a09ca4309404b6e97f29315f4dd7c5276013bfe765efc09f88ad65"></a>

## Direct properties — voltstack_cluster_ar.node.local_subnet / 731b3af91b0e / 3

- [subnet](resources--azure_vnet_site--reference--group-010.md#canonical-ae36652e2229b303d12b07d552eb5e0f7036410132bacb28d2b479ea4458876d): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-010.md#canonical-1d1fa22656f9654a8f0dbfecdb0cb26bfe7373d55a7257788da5f48a7fff9df2): complete subsection reference.

<a id="canonical-f72047055100d4c7a9723c7296a774636cd0f194eeaa6125773b81b457920464"></a>

## Next pages — voltstack_cluster_ar.node.local_subnet / 731b3af91b0e / 4

- [voltstack_cluster_ar.node.local_subnet.subnet](resources--azure_vnet_site--reference--group-010.md#canonical-ae36652e2229b303d12b07d552eb5e0f7036410132bacb28d2b479ea4458876d)
- [voltstack_cluster_ar.node.local_subnet.subnet_param](resources--azure_vnet_site--reference--group-010.md#canonical-1d1fa22656f9654a8f0dbfecdb0cb26bfe7373d55a7257788da5f48a7fff9df2)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-009.md#canonical-fb7ee52f5f58686836f1724c5f129953df3fd12bcd2f2bcbc459c1d30ea4c086)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ae36652e2229b303d12b07d552eb5e0f7036410132bacb28d2b479ea4458876d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5d9447e0020c540ae8483dfee805aab1f6dd3c5e04fdf39fc41d0a090adfb91"></a>

## voltstack_cluster_ar.node.local_subnet.subnet — voltstack_cluster_ar.node.local_subnet.subnet / 0dabc17abb47 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-009.md#canonical-fb7ee52f5f58686836f1724c5f129953df3fd12bcd2f2bcbc459c1d30ea4c086)
- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-009.md#canonical-bef5d1ec441f891737be5a113eb09bf7ef6dbfab5b5c84741f15bc0cc96227ab)
- voltstack_cluster_ar.node.local_subnet.subnet

<a id="canonical-43b24388389fa38be00c39474a21236281c9fdc639b8aa90c007aa3bddd75ae0"></a>

Type: `"object"`. single nested block, Optional.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name"),
  validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-acd317dfb7976f7eafc9cd34a8273611ca34547f998487279532077a56ebc9c7"></a>

## Direct properties — voltstack_cluster_ar.node.local_subnet.subnet / 0dabc17abb47 / 3

<a id="canonical-fb725d4905559bac2c50b0ab9c4fef29c1de0362172932ec83aa31e88287906b"></a>

<a id="canonical-a1cfa3d4d5e2bc9267b5fe809799ffe1a7efeb8e8ca125bee7e084bac7105977"></a>

## subnet_name property — voltstack_cluster_ar.node.local_subnet.subnet / 0dabc17abb47 / 4

Type: `"string"`. Optional.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-cb4c42739e1a11fa23b7b92dbe8e7c88d97e372a0bab63d03693bd8b76c7a128"></a>

<a id="canonical-ee6b26748746dea263bfda399ad035bd6e9e07ce2dd049f835c7566c265ec47b"></a>

## subnet_resource_grp property — voltstack_cluster_ar.node.local_subnet.subnet / 0dabc17abb47 / 5

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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

- [vnet_resource_group](resources--azure_vnet_site--reference--group-010.md#canonical-4f593e0f8c071bbe8acdebfb7b4e199034c7a11da2e6de50a7bbd9b7e42d2a5e): complete subsection reference.

<a id="canonical-fe7a7cd95dfd2396a249de74bdae21eec6a25f2641222f22c35e2254acb8822f"></a>

## Next pages — voltstack_cluster_ar.node.local_subnet.subnet / 0dabc17abb47 / 6

- [voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-010.md#canonical-4f593e0f8c071bbe8acdebfb7b4e199034c7a11da2e6de50a7bbd9b7e42d2a5e)
- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-009.md#canonical-bef5d1ec441f891737be5a113eb09bf7ef6dbfab5b5c84741f15bc0cc96227ab)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-4f593e0f8c071bbe8acdebfb7b4e199034c7a11da2e6de50a7bbd9b7e42d2a5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5471ce6f698bb05105163aea3f57c12da1b74e5850c61d032d8c6013a947dc5b"></a>

## voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group — voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group / 0513c6d5019d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-009.md#canonical-fb7ee52f5f58686836f1724c5f129953df3fd12bcd2f2bcbc459c1d30ea4c086)
- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-009.md#canonical-bef5d1ec441f891737be5a113eb09bf7ef6dbfab5b5c84741f15bc0cc96227ab)
- [voltstack_cluster_ar.node.local_subnet.subnet](resources--azure_vnet_site--reference--group-010.md#canonical-ae36652e2229b303d12b07d552eb5e0f7036410132bacb28d2b479ea4458876d)
- voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group

<a id="canonical-7d92977f6ccc05b5edbf5bc055a9a40e63e65b928bbb6507131553c9019dc1bd"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

<a id="canonical-42eeb390ac324884dcc190f3521b58ea9f803d1f3f60629136054fe81a425f6f"></a>

## Direct properties — voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group / 0513c6d5019d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b51cc436b0d1e1d8e6c40c584821199afa56ab0d9e8793761a1cf9787bd9efbf"></a>

## Next pages — voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group / 0513c6d5019d / 4

- [voltstack_cluster_ar.node.local_subnet.subnet](resources--azure_vnet_site--reference--group-010.md#canonical-ae36652e2229b303d12b07d552eb5e0f7036410132bacb28d2b479ea4458876d)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-1d1fa22656f9654a8f0dbfecdb0cb26bfe7373d55a7257788da5f48a7fff9df2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f70e6753f5e9e7c82e08e7fdae73383506c7cf7699d4edf6901b4b5f1f9acb4"></a>

## voltstack_cluster_ar.node.local_subnet.subnet_param — voltstack_cluster_ar.node.local_subnet.subnet_param / 58edb0a5e5d0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-009.md#canonical-fb7ee52f5f58686836f1724c5f129953df3fd12bcd2f2bcbc459c1d30ea4c086)
- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-009.md#canonical-bef5d1ec441f891737be5a113eb09bf7ef6dbfab5b5c84741f15bc0cc96227ab)
- voltstack_cluster_ar.node.local_subnet.subnet_param

<a id="canonical-5e8136d81fc39a0cd88907d1760a2f4e9d54fe177290b0484faffcb91bdf55b2"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-0d08e02bfc0bdfd234a625625bc0c87041738f9bd42966e7e628431fd46a337c"></a>

## Direct properties — voltstack_cluster_ar.node.local_subnet.subnet_param / 58edb0a5e5d0 / 3

<a id="canonical-5909500fc884b3442cc04fa92dc46ae414298059620337b8124c78d2c3a788cb"></a>

<a id="canonical-8dac14c586c1d81e3e28be045abc5215c5cd760120464ec0db15609a81d2c3ad"></a>

## ipv4 property — voltstack_cluster_ar.node.local_subnet.subnet_param / 58edb0a5e5d0 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
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
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-78f4c2055eaaa2868a9fe47aa62c4f24436cb760a29d280a7830c66807744022"></a>

## Next pages — voltstack_cluster_ar.node.local_subnet.subnet_param / 58edb0a5e5d0 / 5

- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-009.md#canonical-bef5d1ec441f891737be5a113eb09bf7ef6dbfab5b5c84741f15bc0cc96227ab)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0c7cfca1d2d823e1d834a0c55f66bb98ce3e604c0fc3f4e7628ab50bbfb70ed3"></a>

## voltstack_cluster_ar.outside_static_routes — voltstack_cluster_ar.outside_static_routes / 482e37a53df5 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.outside_static_routes

<a id="canonical-47c9d100a62d8bc72b6de46972982bc828582ea7f294eeb46fcdb62a0c722ee6"></a>

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

<a id="canonical-0700ac3e54f2a0d6ba819b6c2ce309c74f17d1cd326a26e825c0124ca15dd2d7"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes / 482e37a53df5 / 3

- [static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e): complete subsection reference.

<a id="canonical-9454fab27a4ee0a1798d0de7d45a5c02d79dc7e25c5f7aa336c80b8e81982b56"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes / 482e37a53df5 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8f59c79c6c4c05ad817b9e0b8ae67e8a7a3b0a04d1972f8ddc303851ac3c67e1"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list — voltstack_cluster_ar.outside_static_routes.static_route_list / 002f3cf2aa16 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- voltstack_cluster_ar.outside_static_routes.static_route_list

<a id="canonical-da3a244f84f8c4e24ff206a3710445fc60793a2ef8b49602d569e916f20d242c"></a>

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

<a id="canonical-dad3d30ef4122de8e8f526e8ef305f7fb22744684ec70ddada7802049d9a21eb"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list / 002f3cf2aa16 / 3

- [custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590): complete subsection reference.

<a id="canonical-ea6c9f8879907963d2911a2d152bed7cbdc7fcecccde2e5c1c62ef484002ecf5"></a>

<a id="canonical-ff0183a4fe25cb52e21da7b1f9a238b4f13afdfe3f5f58d3003b84cd316abd2d"></a>

## simple_static_route property — voltstack_cluster_ar.outside_static_routes.static_route_list / 002f3cf2aa16 / 4

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

<a id="canonical-f971d700e0bebb8773e6705e432dfb80c85f0c65f4243c349de74408b832361a"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list / 002f3cf2aa16 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dd63a967b7bb473e86aaf16b29a7327b41554a32eddd4621ffe89e19d61ac22"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / a7c8e064f8ac / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-6f43ffb1afdeb513d7f1198e8a2294c72ee01afa6175cfa6745ec66474f75518"></a>

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

<a id="canonical-12593e047be31cc32a4d7a0b371840ba393c37fa429a9826e113af13b20a39a2"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / a7c8e064f8ac / 3

<a id="canonical-f207773167b2bae13e5496f3450cf8ffa967544c1aebcfc110ae47ea56c700a5"></a>

<a id="canonical-14eb6c7063cf9b821e5bc2ac2d86dd9c691b3bf7ff0ae2e5aed375791697da4a"></a>

## attrs property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / a7c8e064f8ac / 4

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

- [labels](resources--azure_vnet_site--reference--group-010.md#canonical-73f272679556b12b31fbce637c1dd37273bcd8e55d85cafce0a016d79cae761e): complete subsection reference.

- [nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e): complete subsection reference.

- [subnets](resources--azure_vnet_site--reference--group-010.md#canonical-843d0f085759619b1e8725ef41476a4a7bc50283fca119b6737e5c162eeab697): complete subsection reference.

<a id="canonical-becf6b887dced77d89404d0909d6b6ffa45a0c4332970f590e479ccdbb465e58"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / a7c8e064f8ac / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels](resources--azure_vnet_site--reference--group-010.md#canonical-73f272679556b12b31fbce637c1dd37273bcd8e55d85cafce0a016d79cae761e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-010.md#canonical-843d0f085759619b1e8725ef41476a4a7bc50283fca119b6737e5c162eeab697)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-73f272679556b12b31fbce637c1dd37273bcd8e55d85cafce0a016d79cae761e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-559ded0f7288efdbb438ed21d5c964223dc5c5f3e462000e17624d2789abd0a8"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 69cd5d5d0f36 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-2ef20272faee1e871f47f60caadfee7324fb57e8eddb76aba776bc30ab830b5e"></a>

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

<a id="canonical-dd6bbafd9421239bf27f5525ec2efd8cef8c17d59c67b1e4fe44c2c86d398ca6"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 69cd5d5d0f36 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-599023e55cffd8e4195ef2d2f91af4a5881d75499d9b73aecfadbe0223abf6dd"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 69cd5d5d0f36 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d1d703f2669575a308b9ff88f1b80ab3732fa1a7bc063b1fa984f22dc61402e"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 6d089e73dadd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-7afe6c90e225840c005860e75156f9f14e44bddbca05bb9b2f6cbd5e7f2ba274"></a>

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

<a id="canonical-b460f0575cb14ff8b982d27b2f82bc1f765f8706b97a94b570f904f76cb22eb3"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 6d089e73dadd / 3

- [interface](resources--azure_vnet_site--reference--group-010.md#canonical-870ed4afdb910b95621bd94c0fdb8cc14ec182342df4462690929e3f066d6bea): complete subsection reference.

- [nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-bede02a6f568a837ddd1f02d7859f3de9c0c5439a55f941edbaaa6f045cf8054): complete subsection reference.

<a id="canonical-87610a039f6b292218c065843712e6ada76bac80016c9bbe19064b3d3d6d0e7a"></a>

<a id="canonical-64314d017580df87ab871dc1aad5af352f158390ca8e273a1073605e535d9bb3"></a>

## type property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 6d089e73dadd / 4

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

<a id="canonical-d84ba00cfcbeba11adc1e5dd93371ed3f8999d18c19c55bec416c562a18195fd"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 6d089e73dadd / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--azure_vnet_site--reference--group-010.md#canonical-870ed4afdb910b95621bd94c0fdb8cc14ec182342df4462690929e3f066d6bea)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-bede02a6f568a837ddd1f02d7859f3de9c0c5439a55f941edbaaa6f045cf8054)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-870ed4afdb910b95621bd94c0fdb8cc14ec182342df4462690929e3f066d6bea"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ddaca520e1d8764e8153140ff00fe5a126c1312cda89b4383c7124732838160"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c708521ac409 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-4d26b2b768b7928744f518c3197686636bcc15b3072627ba3a64f8c5330ed770"></a>

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

<a id="canonical-14970eb0f18c6581386c69986a172dfb649a6c894d3f88cb1241d9010e9cae39"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c708521ac409 / 3

<a id="canonical-aab9224698206516494608e24562418207bb4cded6f3106c557c1b2377076089"></a>

<a id="canonical-2e7183e1cd24b34899b7b3937e56b1a86b347c2138e84370db69efc74c2ce68e"></a>

## kind property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c708521ac409 / 4

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

<a id="canonical-69e81f08a75f5f4ceb1eefae58a1afa65e14663c06ae02a65fa30616d34958fe"></a>

<a id="canonical-ed89a73cb6e00c00fb67d8b4b1f2996c5d0efb832d0171965c1e17dc2e18cd2e"></a>

## name property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c708521ac409 / 5

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

<a id="canonical-47d63ec73663f2edb5614269ff4675e182dd9793730d8b1b265c3322753ccd48"></a>

<a id="canonical-9b5b8e004d421b676fb5ac1b3f8fef1bd177bff444381225bbc736f4fa1d31f6"></a>

## namespace property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c708521ac409 / 6

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

<a id="canonical-2b5c56e52626031ae111d8d165dadb0386c59166a49afa959c7ad6146aaa4e2d"></a>

<a id="canonical-9a719719838aae6ecfbdf01caef00ed9e102f2b81fe5f34e73063c1ab6a0cbb7"></a>

## tenant property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c708521ac409 / 7

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

<a id="canonical-7a4c799c50247b7eea7c885bf8a74e1fdb19a31c34e3ddb504524bb47650ff5f"></a>

<a id="canonical-0d8fa3ce5c93a97ee7a02632cdd2bfa981c209062bbea1ac844d56ca166219e5"></a>

## uid property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c708521ac409 / 8

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

<a id="canonical-259a59f5cb9b4f62b52e55e2c4142e38453147a777a8ab0359135712ce91edff"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c708521ac409 / 9

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-bede02a6f568a837ddd1f02d7859f3de9c0c5439a55f941edbaaa6f045cf8054"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-51138cf44c949eaa4bf08bdcf72e9c23ee71ffcf79cdf723b64753600b6b4aae"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 6961a20611cf / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-9f798f4db9b90802220ff24595d06e3a0763b4c437fda441aaf08ff68e425238"></a>

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

<a id="canonical-28686d251e917e45c076842fbdf38824b1f5c19f405685c235510fe139cba2d6"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 6961a20611cf / 3

- [dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-bf88caf1fea253093024a3d8c5f979b9ecd82e182b06b1410fe16e1a4787b846): complete subsection reference.

- [ipv4](resources--azure_vnet_site--reference--group-010.md#canonical-15d9a4b97e96a000f239b3a3c5b3fc07016809e026458b8e5bdc4eb3800c3806): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-010.md#canonical-657102146c349593437c7478571e1d997e1b838d067b9a66cc411b09e67ef6ba): complete subsection reference.

<a id="canonical-3c0479cfd9aac677fe7ee4690a7eddbac102b8009daa9312ea0ce39eb51fb53e"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 6961a20611cf / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-bf88caf1fea253093024a3d8c5f979b9ecd82e182b06b1410fe16e1a4787b846)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--azure_vnet_site--reference--group-010.md#canonical-15d9a4b97e96a000f239b3a3c5b3fc07016809e026458b8e5bdc4eb3800c3806)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--azure_vnet_site--reference--group-010.md#canonical-657102146c349593437c7478571e1d997e1b838d067b9a66cc411b09e67ef6ba)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-bf88caf1fea253093024a3d8c5f979b9ecd82e182b06b1410fe16e1a4787b846"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d7c3534c2876978ea9262efde7afafb46f3e14459ae5c5a5a72d68bc0e0b7f9"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 1ffa6c3ff480 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-bede02a6f568a837ddd1f02d7859f3de9c0c5439a55f941edbaaa6f045cf8054)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-09b202c126a22b74a2e9c336691654533482673b48249cf6cc460df5eea8ecc3"></a>

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

<a id="canonical-32a697ccdc7eaa9a11b273790ef538b5d31c500c70e2cbec3b4469adcf87f241"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 1ffa6c3ff480 / 3

- [ipv4](resources--azure_vnet_site--reference--group-010.md#canonical-6272557a673c4ea0e0e400d9e05ca12cdab9b810675f26afcabc5ec6b7e81d5c): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-010.md#canonical-8fcde7ccd26ccb33008b452f94dab713c7e5b70d939d16eb5e731bcdc117ddd5): complete subsection reference.

<a id="canonical-2b5c06cb8bd4ea6914f36c445f9cb1d4e2ccd2883542345462ea6a7a6b1f7d60"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 1ffa6c3ff480 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--azure_vnet_site--reference--group-010.md#canonical-6272557a673c4ea0e0e400d9e05ca12cdab9b810675f26afcabc5ec6b7e81d5c)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--azure_vnet_site--reference--group-010.md#canonical-8fcde7ccd26ccb33008b452f94dab713c7e5b70d939d16eb5e731bcdc117ddd5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-bede02a6f568a837ddd1f02d7859f3de9c0c5439a55f941edbaaa6f045cf8054)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6272557a673c4ea0e0e400d9e05ca12cdab9b810675f26afcabc5ec6b7e81d5c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad6a7b4f58084e20fc54a0c4e435bfba2176afc2f31cfbf3bb1ba7fddaf15862"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 87388237f639 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-bede02a6f568a837ddd1f02d7859f3de9c0c5439a55f941edbaaa6f045cf8054)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-bf88caf1fea253093024a3d8c5f979b9ecd82e182b06b1410fe16e1a4787b846)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-fae966df4204b22ae3e5f644afd9050b993091c5c026824f25f96b7fc99bfb6d"></a>

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

<a id="canonical-29049ec2a74a2bdadcf695a86e349dea90442d8dde58f19dd4885bdc5a266286"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 87388237f639 / 3

<a id="canonical-9def7b4bcae1a7fd0779c6eb6e60caa7150debbb317b220159881182b006760d"></a>

<a id="canonical-a25a7533fa205d957ed66582acd55943e79ab22a1cab5336e6bd982b881eb9f2"></a>

## addr property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 87388237f639 / 4

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

<a id="canonical-b74679004885636bdf3d973f4a937c6511b01a9bceb667f082d6de4868f074b1"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 87388237f639 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-bf88caf1fea253093024a3d8c5f979b9ecd82e182b06b1410fe16e1a4787b846)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-8fcde7ccd26ccb33008b452f94dab713c7e5b70d939d16eb5e731bcdc117ddd5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7bf62d0679a330f72bfb8a2f265dcd9cf2d011cfeefec67e355f08d91ae7e37b"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / b0d34eecef42 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-bede02a6f568a837ddd1f02d7859f3de9c0c5439a55f941edbaaa6f045cf8054)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-bf88caf1fea253093024a3d8c5f979b9ecd82e182b06b1410fe16e1a4787b846)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-554b2fa5bb589afbcac9de7d2638323b9629712a4bc7ad153ad2810f3056d93f"></a>

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

<a id="canonical-bcc37834055808587d1bf36e83aeceda16dccd0363e0eabbbf5909c1de9ea17f"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / b0d34eecef42 / 3

<a id="canonical-26eb222d9acbf8c0ef01b808ce084470b7219c7bd923d906f88223ac7a20e40e"></a>

<a id="canonical-5cf61d80c6674e6a643372b73dfccc5ef972ea07f1090a95b414e16d227faf20"></a>

## addr property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / b0d34eecef42 / 4

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

<a id="canonical-2b3435b4a632a8387fe0ff9e2190fdff7e4e597597b7ca37485e9cdcc14b81b6"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / b0d34eecef42 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-bf88caf1fea253093024a3d8c5f979b9ecd82e182b06b1410fe16e1a4787b846)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-15d9a4b97e96a000f239b3a3c5b3fc07016809e026458b8e5bdc4eb3800c3806"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6adf667e20bd5ba6f7462dc3192e90a27b25241a4cdca8adcab43fdfc200ed66"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 602e590b8858 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-bede02a6f568a837ddd1f02d7859f3de9c0c5439a55f941edbaaa6f045cf8054)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-9077c7914e0ce2368d32ade97b8d9147272d7b63aae5047eed5f6a7b3d3c5090"></a>

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

<a id="canonical-0c950f27fbb2c627aa5f2aa5d83d16ebd56c20f970886da680f7f2b8f4e332ec"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 602e590b8858 / 3

<a id="canonical-16531960689a59d4ff15b77fd56e496fafeaa6e8eb10fec70d78a8e2cd1e729b"></a>

<a id="canonical-af57637f87d435fdbc92a4a92295918116958632feba074d0ba4dddec5473ad6"></a>

## addr property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 602e590b8858 / 4

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

<a id="canonical-73dfb32ac54cde117ed43446e944163e0c3c3d2817e75ce5b72ff06e685d4ba2"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 602e590b8858 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-bede02a6f568a837ddd1f02d7859f3de9c0c5439a55f941edbaaa6f045cf8054)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-657102146c349593437c7478571e1d997e1b838d067b9a66cc411b09e67ef6ba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a5e9d88d4118a19a12c956f463268dc239bf9242bce56002c8defc9ddb2fe58"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / b998f7e47143 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-19f2297c48f84e991ae690cf1d00f8b389a48af4dd383e9fbcf80861f0e4881e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-bede02a6f568a837ddd1f02d7859f3de9c0c5439a55f941edbaaa6f045cf8054)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-75756782155bdf0bf1e152810dc63516bc23bc3798154298716642cd15dfebd2"></a>

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

<a id="canonical-8a48ae0de3b9e3e7ca946204f4d8268b5474ae23bb47593f8b3026670ddb782e"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / b998f7e47143 / 3

<a id="canonical-b1f90cef50fcab551d34a7c5ad78b6928af7e50b2a7fbd5bac90df73f6610674"></a>

<a id="canonical-fdc952f2ef2a2adaef0eac786284ca2f6fa6c9ba95c884bdd5c670172f6a2d0a"></a>

## addr property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / b998f7e47143 / 4

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

<a id="canonical-f590141779f321e4612af5b8d1b5feeada8b550cd69155e27f8f2bce12a991d5"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / b998f7e47143 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-bede02a6f568a837ddd1f02d7859f3de9c0c5439a55f941edbaaa6f045cf8054)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-843d0f085759619b1e8725ef41476a4a7bc50283fca119b6737e5c162eeab697"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dbe5cc2d3dd85980cdc44feaeff92072961b50969cb73c4822ebe274dbfff407"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 6aa88d66bcf4 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-e5df575b573d085a90f8a8e51003e5543088fd2560f205dfda22a4376420983e"></a>

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

<a id="canonical-3e21c173546c88021160310bc1d62fed54f29004c44a8b1e2dbb5c3354d3e4b2"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 6aa88d66bcf4 / 3

- [ipv4](resources--azure_vnet_site--reference--group-010.md#canonical-e2882af9be7ecdfe0c624af2e45694da795591db10fca6a8b3d61f038fe83d74): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-010.md#canonical-44338f3afdf6a0d4f7fcbd1beb9ef1b836b228982341175ea1370210b04eff23): complete subsection reference.

<a id="canonical-d947d51073f16c6e5b83bc464d0d2476042e233ef0c8add52887419220224f8f"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 6aa88d66bcf4 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--azure_vnet_site--reference--group-010.md#canonical-e2882af9be7ecdfe0c624af2e45694da795591db10fca6a8b3d61f038fe83d74)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--azure_vnet_site--reference--group-010.md#canonical-44338f3afdf6a0d4f7fcbd1beb9ef1b836b228982341175ea1370210b04eff23)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e2882af9be7ecdfe0c624af2e45694da795591db10fca6a8b3d61f038fe83d74"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fbe806f52ccfc0f53ae867836792ccad0eeea9b1457c9f55a2384d9ff2338dc8"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c77c07c1dca6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-010.md#canonical-843d0f085759619b1e8725ef41476a4a7bc50283fca119b6737e5c162eeab697)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-b6b81ea2c5f56a6634f1fe6113a18d6792f22b3c19782b94404660a5af90fcc0"></a>

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

<a id="canonical-3430ae9b3c21ed5523307aa29818ffef8fdef72a82cb048d03bd58a582495804"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c77c07c1dca6 / 3

<a id="canonical-1ec5924c42631d9e748918fc6e18e58eb0fddbcf11c4440a0502b5cf6d82dce6"></a>

<a id="canonical-390145c0026a7d3128d61199bc1a19e7b1db1bea84f908ac9e95afd42b29e5dd"></a>

## plen property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c77c07c1dca6 / 4

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

<a id="canonical-bdb258dd645c93debd837a62d2cbe7748766bac270511d6e52c2b808475f85ba"></a>

<a id="canonical-cea36317662b079cadcad3f85d3fab163932a1a90572ec41665b6189ec872712"></a>

## prefix property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c77c07c1dca6 / 5

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

<a id="canonical-93cde0931c6595f7c3fa3d3ea2a10597fb932dce726a91b265a54baef1d1213a"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / c77c07c1dca6 / 6

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-010.md#canonical-843d0f085759619b1e8725ef41476a4a7bc50283fca119b6737e5c162eeab697)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-44338f3afdf6a0d4f7fcbd1beb9ef1b836b228982341175ea1370210b04eff23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddd3a22fa92bce346535b7a10c0037594cea74d9e9d2a353ebb10c99213a2317"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 3c4f0586c527 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-54afc8939b459e6c0b165e942cf8556519aa686845ea392710b921fd5ef5fae5)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-66f9656f292e5273eeab63243dc2fec34d06a87a00ffb8661f5f03c0e3b7df9e)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-daf8a44d705b696e79071d6a873e7c943955b8d423d392471200b07c4875e590)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-010.md#canonical-843d0f085759619b1e8725ef41476a4a7bc50283fca119b6737e5c162eeab697)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-b82b5a4338c7272387c26a3cf814bd96cf55c88d20932479cd4f534327664a18"></a>

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

<a id="canonical-ab3daeebeb03e7018ef7ba7f02969e212c5f4dfae0bc662480b47341871a94a4"></a>

## Direct properties — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 3c4f0586c527 / 3

<a id="canonical-d8ff1928e7419749034a038bf59ae65e886a280524049aae0d2a64938c604a6f"></a>

<a id="canonical-9593c8e2593ce5ec01801d3b505a54d5d5608c5a1ea8a47409c9fca7339371cd"></a>

## plen property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 3c4f0586c527 / 4

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

<a id="canonical-dfcbd61d67e6300b5d9c537e7a62200f436f0d6d645a555f58a62b7f2d495aa1"></a>

<a id="canonical-f024a59cd8148d9b4c0af1fbb5783f9b31c24c38f470cd0e8aa309434a401eeb"></a>

## prefix property — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 3c4f0586c527 / 5

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

<a id="canonical-6375a6b44c6a4396e7902dcac34d1fb755034199bceff8d4e82077ff1e663738"></a>

## Next pages — voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route / 3c4f0586c527 / 6

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-010.md#canonical-843d0f085759619b1e8725ef41476a4a7bc50283fca119b6737e5c162eeab697)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-21e611a5dd41974704b93d19fc05194f6d2a79ef96057b3871f374955c652813"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b58ca487115eeb4da8fe89589164fc64e4e13174e3b702c22f126cb4999340ff"></a>

## voltstack_cluster_ar.sm_connection_public_ip — voltstack_cluster_ar.sm_connection_public_ip / fdee452f67cd / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.sm_connection_public_ip

<a id="canonical-6dd1ced6e81e0d63235edf91af074aa977bec9238a935bff52a522fcd0229c7b"></a>

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

<a id="canonical-53f9ba83abde576d43cc6389d88562677136e3afb9b2d7fa8c3f038d3e237e28"></a>

## Direct properties — voltstack_cluster_ar.sm_connection_public_ip / fdee452f67cd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5e6867a1d6c296e7cb15ea24dcaeafcf49a089cf69a6619d850a09dd6beedb5e"></a>

## Next pages — voltstack_cluster_ar.sm_connection_public_ip / fdee452f67cd / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-4040feef611420346b88754f0042d4c28a7a7349ecea3da7822f3c67e787f512"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec1412b207ae512242ebc8e4aad106b0237e9d934579b96f613bbba0e10c2913"></a>

## voltstack_cluster_ar.sm_connection_pvt_ip — voltstack_cluster_ar.sm_connection_pvt_ip / 8ed91e3ef8a6 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.sm_connection_pvt_ip

<a id="canonical-7f0e26df1f8d8a2b79ea66c1917dc65ec75d19b04991e41b9d3ca9875b913d43"></a>

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

<a id="canonical-0f3093f275447006f1f74a1f55485fdce035147d93f0f56a495449f858643466"></a>

## Direct properties — voltstack_cluster_ar.sm_connection_pvt_ip / 8ed91e3ef8a6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cceb94a9786702807b0034f5307e4d01bb3b7bd928dd3728c8a2a108adbdcb93"></a>

## Next pages — voltstack_cluster_ar.sm_connection_pvt_ip / 8ed91e3ef8a6 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ffa39091461c785dd6e267139ed919d5c5313c0a21b41f9fab97d36976d6e5d1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5321004cd7ee7c8d655ac2cb58d22f75e68236f3e8cb474ed34a9a28bac787fb"></a>

## voltstack_cluster_ar.storage_class_list — voltstack_cluster_ar.storage_class_list / 7616498c3643 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- voltstack_cluster_ar.storage_class_list

<a id="canonical-575cd879cc87b95aa326f1bc4a84363191913c5769377df1306c853eae207639"></a>

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

<a id="canonical-988444433762addd2b4be3b41fcb90de77fed85da1515dc0343d9e7d3f1338a9"></a>

## Direct properties — voltstack_cluster_ar.storage_class_list / 7616498c3643 / 3

- [storage_classes](resources--azure_vnet_site--reference--group-010.md#canonical-a2d23cbfaed909eddedb4d1c0baf90ba3ecedd8511226795fb9a11a1074940d7): complete subsection reference.

<a id="canonical-01aa052ffa77656e6fbfe8985759c1b99324d2277a455d0d9e066ce790ee529e"></a>

## Next pages — voltstack_cluster_ar.storage_class_list / 7616498c3643 / 4

- [voltstack_cluster_ar.storage_class_list.storage_classes](resources--azure_vnet_site--reference--group-010.md#canonical-a2d23cbfaed909eddedb4d1c0baf90ba3ecedd8511226795fb9a11a1074940d7)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a2d23cbfaed909eddedb4d1c0baf90ba3ecedd8511226795fb9a11a1074940d7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8dcb8483df1cc731e75f8bd77e7efb2decb917df1fbbebba9fb7775cbf5a5002"></a>

## voltstack_cluster_ar.storage_class_list.storage_classes — voltstack_cluster_ar.storage_class_list.storage_classes / f01559f272ef / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-6f580cdc3a5e6e671545173fe56a2ea390872b3a8bffd392b4590dabc956995f)
- [voltstack_cluster_ar.storage_class_list](resources--azure_vnet_site--reference--group-010.md#canonical-ffa39091461c785dd6e267139ed919d5c5313c0a21b41f9fab97d36976d6e5d1)
- voltstack_cluster_ar.storage_class_list.storage_classes

<a id="canonical-98d0ae5ab126e19e371b33676711b7ba3cd559707949e9c7632b534f4b5a5c57"></a>

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

<a id="canonical-eafbebb3bf3e8b4aa5d03afa0a6ea0c3967a95de258fda634025297fd1500636"></a>

## Direct properties — voltstack_cluster_ar.storage_class_list.storage_classes / f01559f272ef / 3

<a id="canonical-e551d197a52e452be8092f457e47d27343c69ae87d7d9beac032318c0eff62f8"></a>

<a id="canonical-ee252d1ac703828f2d4417232467496c85d4ab5aa42a161125d791071ec73199"></a>

## default_storage_class property — voltstack_cluster_ar.storage_class_list.storage_classes / f01559f272ef / 4

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

<a id="canonical-b73d5777c62095ee3ad9725682d7bbaaf01858c90ae54253d7ba217b1fd44670"></a>

<a id="canonical-6a9b407fae4d664f797db9be283fc651268d95df3a007cc04ab65c3a69fc74bc"></a>

## storage_class_name property — voltstack_cluster_ar.storage_class_list.storage_classes / f01559f272ef / 5

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

<a id="canonical-84b10474de157a6ef47820a2c3019d0ad74a2463656281f1616b6a67536ae0e3"></a>

## Next pages — voltstack_cluster_ar.storage_class_list.storage_classes / f01559f272ef / 6

- [voltstack_cluster_ar.storage_class_list](resources--azure_vnet_site--reference--group-010.md#canonical-ffa39091461c785dd6e267139ed919d5c5313c0a21b41f9fab97d36976d6e5d1)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-804652d0f312cb6f1ffe64d05165eea1d9bb3254ff3dad87b3c1cdd5febb0bdb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47b12e78269c6a552b9e5a3003b6b80390905ca24e285d63f6d61525369d35a3"></a>

## waf_signatures — waf_signatures / 9fd03d3572dc / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- waf_signatures

<a id="canonical-3c7dc0804428602d425da6ea51e7581db67bdbcba25c29e6107ef53235192178"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
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
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-f08f6ed7672713b2aa0f6c380c2c90da32c0a6b6d2234d456268a7e3cdfd37dd"></a>

## Direct properties — waf_signatures / 9fd03d3572dc / 3

- [automatic](resources--azure_vnet_site--reference--group-010.md#canonical-0bb5716f6ef80a3a9e98978f737de6b7d470a5dd9f0e3136fddb3a6d15b860a3): complete subsection reference.

- [manual](resources--azure_vnet_site--reference--group-010.md#canonical-9ab0b1fa9d04a381e76fd800fa8ac4a878fddfef57bea30156112a16d30f7b98): complete subsection reference.

<a id="canonical-1277778f41f21f5f7c7ec7f33a022c0493f73503d015303ec39ae24e32a5298a"></a>

## Next pages — waf_signatures / 9fd03d3572dc / 4

- [waf_signatures.automatic](resources--azure_vnet_site--reference--group-010.md#canonical-0bb5716f6ef80a3a9e98978f737de6b7d470a5dd9f0e3136fddb3a6d15b860a3)
- [waf_signatures.manual](resources--azure_vnet_site--reference--group-010.md#canonical-9ab0b1fa9d04a381e76fd800fa8ac4a878fddfef57bea30156112a16d30f7b98)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-0bb5716f6ef80a3a9e98978f737de6b7d470a5dd9f0e3136fddb3a6d15b860a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2c5f97f9e41d117b25a7e5c2b138dca7c95a2b3d1268483be4b960c9bbe3ef8f"></a>

## waf_signatures.automatic — waf_signatures.automatic / c707b40b5d01 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [waf_signatures](resources--azure_vnet_site--reference--group-010.md#canonical-804652d0f312cb6f1ffe64d05165eea1d9bb3254ff3dad87b3c1cdd5febb0bdb)
- waf_signatures.automatic

<a id="canonical-66b84f9e4b150adf73f1e667091c7291d3b3c81df5594139f03732e2d83ab9a4"></a>

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
automatic = {}
```

<a id="canonical-44fb9227326b5c87a982b8664f7d4a8d0346773cf6c960444db34d927171b1a7"></a>

## Direct properties — waf_signatures.automatic / c707b40b5d01 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2e78a5dbd3abc9b7454a009d31c65246428c76b24cc119244f6cf1e5c0a07d0"></a>

## Next pages — waf_signatures.automatic / c707b40b5d01 / 4

- [waf_signatures](resources--azure_vnet_site--reference--group-010.md#canonical-804652d0f312cb6f1ffe64d05165eea1d9bb3254ff3dad87b3c1cdd5febb0bdb)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9ab0b1fa9d04a381e76fd800fa8ac4a878fddfef57bea30156112a16d30f7b98"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3db2f4b43ed4f1a37139abc6f8e2df1d867ededf6db8f762da0c8609811d46d"></a>

## waf_signatures.manual — waf_signatures.manual / 2d36ca3589d4 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [waf_signatures](resources--azure_vnet_site--reference--group-010.md#canonical-804652d0f312cb6f1ffe64d05165eea1d9bb3254ff3dad87b3c1cdd5febb0bdb)
- waf_signatures.manual

<a id="canonical-241e43bb21a584352c7d9b6b2ffad38b2bc239b97e5ecd564a737acd483c5275"></a>

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
manual = {}
```

<a id="canonical-cef8eb1005ee0e4880fbcfd7dc1d0680703aa7bc3b797a854920e783ebc4b402"></a>

## Direct properties — waf_signatures.manual / 2d36ca3589d4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b3da84bdbd4324f29ba389e9046d0bf502ca74ecf769b72a885e723e26c5c6f4"></a>

## Next pages — waf_signatures.manual / 2d36ca3589d4 / 4

- [waf_signatures](resources--azure_vnet_site--reference--group-010.md#canonical-804652d0f312cb6f1ffe64d05165eea1d9bb3254ff3dad87b3c1cdd5febb0bdb)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
