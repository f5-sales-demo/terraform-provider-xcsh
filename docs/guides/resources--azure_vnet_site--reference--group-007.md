---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-632118241768425b4b3b9fa8b4b013baba5c4146e9fdf6f69e78e1ea7fff9d50"></a>

## Next pages — ingress_egress_gw_ar.node / b2bff600988a / 7

- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-db3c1f7f142e679d6543aac66b1c1bd3346b98047803d3c64fe13a0c566d9a59)
- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-8f1e154b91ff6e6b188def3b996e66ea86a6f4263a2f7eb6ae15101440c8e74e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-db3c1f7f142e679d6543aac66b1c1bd3346b98047803d3c64fe13a0c566d9a59"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-27705609b454df53cd3eff5cf4c74b9d1b59c0ffe05621d15099f0d3517f9dda"></a>

## ingress_egress_gw_ar.node.inside_subnet — ingress_egress_gw_ar.node.inside_subnet / 3f40499c1434 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe)
- ingress_egress_gw_ar.node.inside_subnet

<a id="canonical-01000e8619a63b9431fdca1c8bc2f3abf2c27a6b0dc83bc7f0a31bc8ccbcbb73"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside subnet.

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
inside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-500ef70c6654c2eebd1e847794591186756fe1869412ebf53e7e8c588478b548"></a>

## Direct properties — ingress_egress_gw_ar.node.inside_subnet / 3f40499c1434 / 3

- [subnet](resources--azure_vnet_site--reference--group-007.md#canonical-905b3d453c90f9b64ebf21ef08e0ce526a859f871a0a3dbf477ffff41507e6fd): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-0f7358c209bbdcdd5af95e19b70e47f187fffea1cba9e51f8f9c447c4f093faf): complete subsection reference.

<a id="canonical-48ba8ee272e173334301824140c0f8c5df6282c464fa6615b3a99627972e371b"></a>

## Next pages — ingress_egress_gw_ar.node.inside_subnet / 3f40499c1434 / 4

- [ingress_egress_gw_ar.node.inside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-905b3d453c90f9b64ebf21ef08e0ce526a859f871a0a3dbf477ffff41507e6fd)
- [ingress_egress_gw_ar.node.inside_subnet.subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-0f7358c209bbdcdd5af95e19b70e47f187fffea1cba9e51f8f9c447c4f093faf)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-905b3d453c90f9b64ebf21ef08e0ce526a859f871a0a3dbf477ffff41507e6fd"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53f68c1f539fdd5716f4fdc3ebb6de397e3a89475d0ebf7270cfcac2031d3735"></a>

## ingress_egress_gw_ar.node.inside_subnet.subnet — ingress_egress_gw_ar.node.inside_subnet.subnet / ebc8417422d1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe)
- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-db3c1f7f142e679d6543aac66b1c1bd3346b98047803d3c64fe13a0c566d9a59)
- ingress_egress_gw_ar.node.inside_subnet.subnet

<a id="canonical-0cfda83f262f229c800fb23fe4d4eedecc12a1eae6e0e0f6ebf755e87152aa95"></a>

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

<a id="canonical-4ead4172aef2dedf8a1141c24436ae8f62a1c5f112d1659bf7a57383f7776428"></a>

## Direct properties — ingress_egress_gw_ar.node.inside_subnet.subnet / ebc8417422d1 / 3

<a id="canonical-ab47c4ffb18b34c5af0cff51a8f24f3db06256992791583f30adb3d0eaec29c7"></a>

<a id="canonical-10158f70b6fe56154057aa12c143ce6d537290f08e177235d2d24c123e7c6b76"></a>

## subnet_name property — ingress_egress_gw_ar.node.inside_subnet.subnet / ebc8417422d1 / 4

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

<a id="canonical-f4fbad052ef4c558aef207657f3d756f0eb830d5c702c296cc1c0c2986b13d2c"></a>

<a id="canonical-29122d4b26f919b00ab6183b1e9570d6ee93d4642af15c0e4c50a05eab1b2fbb"></a>

## subnet_resource_grp property — ingress_egress_gw_ar.node.inside_subnet.subnet / ebc8417422d1 / 5

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

- [vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-9d1e1be81ee658a58b1b460acb1fb4e8ce7d0c10d2256e6ffe31afaaa1cfd49e): complete subsection reference.

<a id="canonical-7b289c92c30edcb9c43dbeb4910e4fda37ab399825c1bb5f926bb18775edb959"></a>

## Next pages — ingress_egress_gw_ar.node.inside_subnet.subnet / ebc8417422d1 / 6

- [ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-9d1e1be81ee658a58b1b460acb1fb4e8ce7d0c10d2256e6ffe31afaaa1cfd49e)
- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-db3c1f7f142e679d6543aac66b1c1bd3346b98047803d3c64fe13a0c566d9a59)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9d1e1be81ee658a58b1b460acb1fb4e8ce7d0c10d2256e6ffe31afaaa1cfd49e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d8243ccfd955e2a47e24ffa01019e7918b8035c0134ddca0cabcf2c6158abc0"></a>

## ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group — ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group / 48687a44025f / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe)
- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-db3c1f7f142e679d6543aac66b1c1bd3346b98047803d3c64fe13a0c566d9a59)
- [ingress_egress_gw_ar.node.inside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-905b3d453c90f9b64ebf21ef08e0ce526a859f871a0a3dbf477ffff41507e6fd)
- ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group

<a id="canonical-a4350d82026d55adda505255133d45e14b0f8c857ceeb0b19034b0bfa32ef2b8"></a>

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

<a id="canonical-c6c7fbcc594b923853e1b19cb82bf48c1ffc3c18822832c97ce09e1b5388374e"></a>

## Direct properties — ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group / 48687a44025f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5c69b8dbe6e19459adf68d96cf1300089b6f298aa820a7eda3a4cdbce963e7e7"></a>

## Next pages — ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group / 48687a44025f / 4

- [ingress_egress_gw_ar.node.inside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-905b3d453c90f9b64ebf21ef08e0ce526a859f871a0a3dbf477ffff41507e6fd)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-0f7358c209bbdcdd5af95e19b70e47f187fffea1cba9e51f8f9c447c4f093faf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-428ccb5541078a6cbe5cb2f0ca9da4e408172778f53d4a5274f23e9f63ff530d"></a>

## ingress_egress_gw_ar.node.inside_subnet.subnet_param — ingress_egress_gw_ar.node.inside_subnet.subnet_param / 63715a6e77ea / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe)
- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-db3c1f7f142e679d6543aac66b1c1bd3346b98047803d3c64fe13a0c566d9a59)
- ingress_egress_gw_ar.node.inside_subnet.subnet_param

<a id="canonical-e04588dabf7cfb1e384eda2db38b0206730afbad176d56ec904413cfd7475798"></a>

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

<a id="canonical-7f33dd1032d92ffb0c548a879d53d1f22f9388d3348b47138a0285a5ad5c6e43"></a>

## Direct properties — ingress_egress_gw_ar.node.inside_subnet.subnet_param / 63715a6e77ea / 3

<a id="canonical-eddfcab9a026fcdfd537515943dde86096bb0127280d56281136dea0a15666fa"></a>

<a id="canonical-1e45c7190d86972dcc3e190139d8c07634a8eee2ac0f1ce5f173b7651933dc1d"></a>

## ipv4 property — ingress_egress_gw_ar.node.inside_subnet.subnet_param / 63715a6e77ea / 4

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

<a id="canonical-f2767c3d24e4d50733bf5c5de4c81d6b25fcae2702f4fdaf8363e0b8a2958a53"></a>

## Next pages — ingress_egress_gw_ar.node.inside_subnet.subnet_param / 63715a6e77ea / 5

- [ingress_egress_gw_ar.node.inside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-db3c1f7f142e679d6543aac66b1c1bd3346b98047803d3c64fe13a0c566d9a59)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-8f1e154b91ff6e6b188def3b996e66ea86a6f4263a2f7eb6ae15101440c8e74e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2de372375f2d54592584334cafe94eb0f3c039025d77031e883d9f271a22f310"></a>

## ingress_egress_gw_ar.node.outside_subnet — ingress_egress_gw_ar.node.outside_subnet / 676767a5f272 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe)
- ingress_egress_gw_ar.node.outside_subnet

<a id="canonical-7c3e654e8f51d787eb6ea9e62d920a2b6b745d31d8322bbad98157cd67c1e64c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside subnet.

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
outside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-10af125d9636a18b5639d7cbb1788a073c6f0aefaa7910a310fd76bd56295b99"></a>

## Direct properties — ingress_egress_gw_ar.node.outside_subnet / 676767a5f272 / 3

- [subnet](resources--azure_vnet_site--reference--group-007.md#canonical-ebe40aed26a0b7c974f3df51031a688c824b77003fd2e973092fe6bd5bbb91f4): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-03f1bed4ba4029f74c3bfcd612c43644c04a56e135e797967f03481e63180feb): complete subsection reference.

<a id="canonical-2cd21b97ecd1841745f401b3068099514063e0067bcb52c6bc4aa7edb6b5de54"></a>

## Next pages — ingress_egress_gw_ar.node.outside_subnet / 676767a5f272 / 4

- [ingress_egress_gw_ar.node.outside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-ebe40aed26a0b7c974f3df51031a688c824b77003fd2e973092fe6bd5bbb91f4)
- [ingress_egress_gw_ar.node.outside_subnet.subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-03f1bed4ba4029f74c3bfcd612c43644c04a56e135e797967f03481e63180feb)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ebe40aed26a0b7c974f3df51031a688c824b77003fd2e973092fe6bd5bbb91f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-799f0bf42a3cf196771d7e383154ea9f98de49503dcb2a4663ea41a109404fd0"></a>

## ingress_egress_gw_ar.node.outside_subnet.subnet — ingress_egress_gw_ar.node.outside_subnet.subnet / 1c5ca301d4ea / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe)
- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-8f1e154b91ff6e6b188def3b996e66ea86a6f4263a2f7eb6ae15101440c8e74e)
- ingress_egress_gw_ar.node.outside_subnet.subnet

<a id="canonical-bb3cf80527ebc1ef1db32ffffb5a8d8ac39ebbca295be2d880796ce51b7da453"></a>

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

<a id="canonical-26ca765215b8e3bf2ae59bdd0bad9d626b82bd4344b451a82d3970c872b3d975"></a>

## Direct properties — ingress_egress_gw_ar.node.outside_subnet.subnet / 1c5ca301d4ea / 3

<a id="canonical-8541a2ceb6687a38d4db773acc3ccea53b7f9ea168e32d0b509353b99cd0d0dd"></a>

<a id="canonical-51ce4b91fa87b86e8bd7951912348700c63f10dc658cca63cd3c7d20609b5bee"></a>

## subnet_name property — ingress_egress_gw_ar.node.outside_subnet.subnet / 1c5ca301d4ea / 4

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

<a id="canonical-ebf3df06b2e1f5a1c359ad36497719a41a21c8b06329bdd025dbc3c1b51b8fc0"></a>

<a id="canonical-5ed6d2bde4afdeb43e7dec057e30a53d8118dd3de5a1383ca077a1ae9f9cc886"></a>

## subnet_resource_grp property — ingress_egress_gw_ar.node.outside_subnet.subnet / 1c5ca301d4ea / 5

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

- [vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-351c617a865a4744c43b1add2dd86432d0212dbe371d13b60ceb2a39f57ef715): complete subsection reference.

<a id="canonical-6f21fe97f5f9d07eefe30fc295e7ee7c2a74a7e030a77105de59260f5ac9ec6d"></a>

## Next pages — ingress_egress_gw_ar.node.outside_subnet.subnet / 1c5ca301d4ea / 6

- [ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-351c617a865a4744c43b1add2dd86432d0212dbe371d13b60ceb2a39f57ef715)
- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-8f1e154b91ff6e6b188def3b996e66ea86a6f4263a2f7eb6ae15101440c8e74e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-351c617a865a4744c43b1add2dd86432d0212dbe371d13b60ceb2a39f57ef715"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1a09e7098f7b998ac46e10abfba31ee4e8e4356fd0ddcf87aee76035d2b5e856"></a>

## ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group — ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group / 686e0c1e5ece / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe)
- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-8f1e154b91ff6e6b188def3b996e66ea86a6f4263a2f7eb6ae15101440c8e74e)
- [ingress_egress_gw_ar.node.outside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-ebe40aed26a0b7c974f3df51031a688c824b77003fd2e973092fe6bd5bbb91f4)
- ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group

<a id="canonical-a58cb71cf5f2a69ed37081a841eaa718b7394a214b93362f1864425716eba550"></a>

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

<a id="canonical-58e7b59046b23b1c4f46312d589daf7b7ecfa83675340bd98e1f80dc53b59e29"></a>

## Direct properties — ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group / 686e0c1e5ece / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7cf5f1f5b5952f5161086960ea6b0715ff2d623a760181f75d9ce31a7781378d"></a>

## Next pages — ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group / 686e0c1e5ece / 4

- [ingress_egress_gw_ar.node.outside_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-ebe40aed26a0b7c974f3df51031a688c824b77003fd2e973092fe6bd5bbb91f4)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-03f1bed4ba4029f74c3bfcd612c43644c04a56e135e797967f03481e63180feb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9bd013e441aaf11258b805c07f2d603560ec7671b3cfcde4106ee75a352f565"></a>

## ingress_egress_gw_ar.node.outside_subnet.subnet_param — ingress_egress_gw_ar.node.outside_subnet.subnet_param / 2fd83fd313b2 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.node](resources--azure_vnet_site--reference--group-006.md#canonical-3477bd025743eda61d04d12322ef9c35e50e230c97b2d323226575b1598833fe)
- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-8f1e154b91ff6e6b188def3b996e66ea86a6f4263a2f7eb6ae15101440c8e74e)
- ingress_egress_gw_ar.node.outside_subnet.subnet_param

<a id="canonical-9a05d4cdf40f1c6eaf224e34864ff0bf4de0d0e65523089ffe5244f417bb662b"></a>

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

<a id="canonical-cd2d16cab59a1582e9a30954ba8d4086ac2c484ed7802ddb8b61ec7df4981649"></a>

## Direct properties — ingress_egress_gw_ar.node.outside_subnet.subnet_param / 2fd83fd313b2 / 3

<a id="canonical-e7916f612b3d15bb4e8b5d2c88e41952d70fccc8b7c7fae9808dac9a53b99978"></a>

<a id="canonical-a7c2c46937894f5fddb50845714efadd5769f0dfdbd793cb2a3bd7d58c7a4f96"></a>

## ipv4 property — ingress_egress_gw_ar.node.outside_subnet.subnet_param / 2fd83fd313b2 / 4

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

<a id="canonical-543a8cafea9eeb84462773df7a11c08852abd1e84b8fa94d9635bc1a04ceea04"></a>

## Next pages — ingress_egress_gw_ar.node.outside_subnet.subnet_param / 2fd83fd313b2 / 5

- [ingress_egress_gw_ar.node.outside_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-8f1e154b91ff6e6b188def3b996e66ea86a6f4263a2f7eb6ae15101440c8e74e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-3e7b6b00c9afdd10dd5ed7a994023188b368fadf3ae535ceaf69484151a058fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8b31ec945a0570f0b861f512d1af1ae39fc8bcfcf60add1e700411efae98c2f0"></a>

## ingress_egress_gw_ar.not_hub — ingress_egress_gw_ar.not_hub / eb71883a07e8 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.not_hub

<a id="canonical-5110084986c5626ab02cebd893c8729ac8f5b67e5fe110d8fe6f4722854027fd"></a>

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
not_hub = {}
```

<a id="canonical-ba81d6e0021b359892eeb22862dd469cb963589aab56dd5822a90101f799a457"></a>

## Direct properties — ingress_egress_gw_ar.not_hub / eb71883a07e8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b073c3223ac789fc9d5f199340dc45f455f734a52918c350358fd547a820c199"></a>

## Next pages — ingress_egress_gw_ar.not_hub / eb71883a07e8 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f031a26de99f374c8cc549b301b5a57c95d1a03741b3b569d50b567aa80fa3f"></a>

## ingress_egress_gw_ar.outside_static_routes — ingress_egress_gw_ar.outside_static_routes / 4b8f4f459afc / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.outside_static_routes

<a id="canonical-0724fd9980a5523fb19c282070380dc8c43009b936f9f7f91d4f40b2b4ff32e5"></a>

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

<a id="canonical-2d627ce4b143804ae727dab7d5097d897756de7cf577f12f8c5dcf312d5b5c69"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes / 4b8f4f459afc / 3

- [static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58): complete subsection reference.

<a id="canonical-163772c8501a12bb5b738d50f7833b5598f16f91837556d358ae9273723dcf70"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes / 4b8f4f459afc / 4

- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-474f67786c9979e384785a85dfa6ba252621c3301e6e8380b6883c1f04e89b55"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list — ingress_egress_gw_ar.outside_static_routes.static_route_list / b2d510c0a18d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- ingress_egress_gw_ar.outside_static_routes.static_route_list

<a id="canonical-44582a36377ce70f9a1a219c11f58e83a9e21c32edf5823f1b0772a39249779c"></a>

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

<a id="canonical-961c4a9b2b14384105921e775137799d52a58c0e6cb48cfc75f707be3ef63a62"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list / b2d510c0a18d / 3

- [custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe): complete subsection reference.

<a id="canonical-d0827a8678b7ad8910c8ead7c18265d464122ffdac979428624c0a1d32d9bc7f"></a>

<a id="canonical-9cdc8d59efabee8b35609421f72b3a22d619cbc186747ee1f0c1ab53beffb230"></a>

## simple_static_route property — ingress_egress_gw_ar.outside_static_routes.static_route_list / b2d510c0a18d / 4

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

<a id="canonical-f2689b08ecf5551e414e95e27a3dae8034e446da8a985325cc46f7c501408233"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list / b2d510c0a18d / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9da5827fb99ccc0e58047c2eb9a12368160fa4d0bc6b1f07969bc481087fe86a"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 49ec911e0893 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-548698a6a0bdfa1aed960d779342fe00b19014ad3821d03f969a55607ad045fe"></a>

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

<a id="canonical-d5ca399f1781456d2629641dc42a34af6939b3b0c0562ea279d78f866812e6c8"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 49ec911e0893 / 3

<a id="canonical-3017d68d26d8a6fa76f9ee8f3d2cae54df943c7521ff6de77751823e6dcd270b"></a>

<a id="canonical-b04372c078ac50ebf44649409e5a065920dd2314db697df422eb95484d8d2fe3"></a>

## attrs property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 49ec911e0893 / 4

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

- [labels](resources--azure_vnet_site--reference--group-007.md#canonical-b9a2107d35836d9560736bad563a4a59c97dfa5566876712ca7aba47c8c682e5): complete subsection reference.

- [nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973): complete subsection reference.

- [subnets](resources--azure_vnet_site--reference--group-007.md#canonical-bbbacacda6edab527e222102f575fd2323276b60e8381926d7c519230cf6f8cb): complete subsection reference.

<a id="canonical-46688a05783f35e82d496725298318d8299c10b725651669da21b11e03f1ca2d"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 49ec911e0893 / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels](resources--azure_vnet_site--reference--group-007.md#canonical-b9a2107d35836d9560736bad563a4a59c97dfa5566876712ca7aba47c8c682e5)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-007.md#canonical-bbbacacda6edab527e222102f575fd2323276b60e8381926d7c519230cf6f8cb)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b9a2107d35836d9560736bad563a4a59c97dfa5566876712ca7aba47c8c682e5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93deba3f13d7566872603522ab669ffd4cec368078077d8017ea6fd76220dbf3"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / dcbb28b4bcd2 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-47c26c5cd2fd7532308b3c3010c12a48c0938e03343cef5e98744c58a0798a9b"></a>

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

<a id="canonical-52d2450dac0feec0496e651d6624b1a2c0cede9f2ba7be4d0def2240b5ed004b"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / dcbb28b4bcd2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-352d3c794f35f8f0a99a2e4b46b05b237cef8476df88dbf71ab3e67bd4feddf1"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / dcbb28b4bcd2 / 4

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09c22eaa52e145681e77d59ffbcbf30d4ab5e4d05afe88e9eeaf0b2f7b35adf1"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / b0592c9ea9e7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-d18be107695c3838dbfe642b2dc236d24263802d42b710090bdaace7238db156"></a>

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

<a id="canonical-b92282b4ffb2e600886619d21a9a12a728fb7e2aec4f5faa426071c9871d5a96"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / b0592c9ea9e7 / 3

- [interface](resources--azure_vnet_site--reference--group-007.md#canonical-df9bd3e66bf5d59221e25e0411730d0c5ee4305e442992b5d6d2fb36417d32a5): complete subsection reference.

- [nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-043c0366474cef1dc77a65d1ad5f31a575f993995c1204159be39b06a0ca7f6e): complete subsection reference.

<a id="canonical-eef6f4596c585ca59ced7b0b6213e5d60354743889cf6f1e23178376eeba340c"></a>

<a id="canonical-18eb0767b337b84722dab39b6ea2618ad34b9728739640e7a60aa80abecb8989"></a>

## type property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / b0592c9ea9e7 / 4

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

<a id="canonical-8efe33053203d422030889d3b1a5241f832bb0cc754c6b307c8dbd864dc4fc7f"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / b0592c9ea9e7 / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--azure_vnet_site--reference--group-007.md#canonical-df9bd3e66bf5d59221e25e0411730d0c5ee4305e442992b5d6d2fb36417d32a5)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-043c0366474cef1dc77a65d1ad5f31a575f993995c1204159be39b06a0ca7f6e)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-df9bd3e66bf5d59221e25e0411730d0c5ee4305e442992b5d6d2fb36417d32a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0368b23699886cc6f741d59c7c306011929272cf252544346cdb125a4dbb7171"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 6cd5ba2f7376 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-b1cc40104d477de97b68ceee09be076fdbff9b599edb639485ccca0dd6442b7f"></a>

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

<a id="canonical-afe1cffe12f551af8e56a9a448d1d1be789163f702a8b949ab8b933a6c9ebb4c"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 6cd5ba2f7376 / 3

<a id="canonical-3e77520e308981b64812f9c775234ffa72670d8ba1814c4870b95c6b93e00619"></a>

<a id="canonical-86f556e10ce3d232e55b4b0d321343b36adca6b99ee81170473e5879da8227b6"></a>

## kind property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 6cd5ba2f7376 / 4

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

<a id="canonical-3896690ee723db2680cfd8f4886ed9a8d3a96c6b40da1a68e09d286c010f47c3"></a>

<a id="canonical-d00ae81b68bcdfce1163f87e05fc74dbfb5b9e05b9ee2403beb8d3131af093cd"></a>

## name property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 6cd5ba2f7376 / 5

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

<a id="canonical-2e8cf5f3c04a17373d840c2ef850129b9c3e00d1a672e8049e10f3c44188e0f1"></a>

<a id="canonical-ef9208222f0a23ec6e8f50e9f333aff908ef230f0daed93d91010e1cd962a533"></a>

## namespace property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 6cd5ba2f7376 / 6

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

<a id="canonical-4f15e2f60c8b9c08498b262c211f29fae27fd89dcdd166d40cfbbdd58f6a3eac"></a>

<a id="canonical-260e63e19b263a80299bcd66d2340da01b9aba28351701a8c934d9358592e3db"></a>

## tenant property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 6cd5ba2f7376 / 7

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

<a id="canonical-6904b8c406666d7e76f905129ff71819edbdaf0a932f261df914a7162163de6e"></a>

<a id="canonical-8ef7cd6bd6d7d1e097ad7e0d6513608bae4b3651964e5f7c1cb9f11bc6aa4190"></a>

## uid property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 6cd5ba2f7376 / 8

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

<a id="canonical-758a2de05bdf71f82aa81af4c90fce49aa8f60c63fe60b58a456d5563e201b4b"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 6cd5ba2f7376 / 9

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-043c0366474cef1dc77a65d1ad5f31a575f993995c1204159be39b06a0ca7f6e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-35d35bfcbaac5cc542823e790fa28b7b123819b8d7b85440b9feaf553c843f14"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 965a72507b2c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-e1ee88f0a587e5e619cf58d7dc6562955d342ecacdf76a3939e23e599e134650"></a>

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

<a id="canonical-a2a188f0209ed4a100a50e1f070e231396255e2c13ad4b0b5635283cef4d5f8f"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 965a72507b2c / 3

- [dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0688cfd28799e65a655441e83cce1b0aa2819716d563d88dfa722b05644f1ae5): complete subsection reference.

- [ipv4](resources--azure_vnet_site--reference--group-007.md#canonical-4eead2df43a0b0ddd1ee2726da95aaa4c666310275a4a49d389ab6a5b2cb5557): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-007.md#canonical-c7a62a9c9b934ea5f42109183058dfb0ee6afaa370d87e6badfe6b34b90b17e3): complete subsection reference.

<a id="canonical-1587fb0413a96ac915f66a40ab0a93aace1caebf0ef461a0c69907270219454f"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 965a72507b2c / 4

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0688cfd28799e65a655441e83cce1b0aa2819716d563d88dfa722b05644f1ae5)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--azure_vnet_site--reference--group-007.md#canonical-4eead2df43a0b0ddd1ee2726da95aaa4c666310275a4a49d389ab6a5b2cb5557)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--azure_vnet_site--reference--group-007.md#canonical-c7a62a9c9b934ea5f42109183058dfb0ee6afaa370d87e6badfe6b34b90b17e3)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-0688cfd28799e65a655441e83cce1b0aa2819716d563d88dfa722b05644f1ae5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e7e1da2c8383bf433103f38ef44ca968b43ac09df91fd63b1dfe74f2f0cc9e2"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 4247c12067a7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-043c0366474cef1dc77a65d1ad5f31a575f993995c1204159be39b06a0ca7f6e)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-4c81c4f7d53f0232a70b8f5aabc0f8f8394fec899bf6ac923a0a87ac8dc477c1"></a>

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

<a id="canonical-cbb1a96cf5617df1b9dde6f8de660672268ed913bd17c481d3dfdcce62352f9e"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 4247c12067a7 / 3

- [ipv4](resources--azure_vnet_site--reference--group-007.md#canonical-271c577c41ced1dced07cd26e24358a6db439421eab03c63578d1147b14f2550): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-007.md#canonical-e8cdf162b73dfd2acbbb0c881dff8a657ed5b5725e1939b57bbc734313a49a0a): complete subsection reference.

<a id="canonical-d9652961f79f16ee2e9476c38e40004b3a862bd771d5f23fcf81673411d297c4"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 4247c12067a7 / 4

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--azure_vnet_site--reference--group-007.md#canonical-271c577c41ced1dced07cd26e24358a6db439421eab03c63578d1147b14f2550)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--azure_vnet_site--reference--group-007.md#canonical-e8cdf162b73dfd2acbbb0c881dff8a657ed5b5725e1939b57bbc734313a49a0a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-043c0366474cef1dc77a65d1ad5f31a575f993995c1204159be39b06a0ca7f6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-271c577c41ced1dced07cd26e24358a6db439421eab03c63578d1147b14f2550"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f51c07937693bbbd5cd7ca4c68de07e29c3e6f67b283f71c4969eae2da922f3f"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4 — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 98dba862908e / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-043c0366474cef1dc77a65d1ad5f31a575f993995c1204159be39b06a0ca7f6e)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0688cfd28799e65a655441e83cce1b0aa2819716d563d88dfa722b05644f1ae5)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4

<a id="canonical-8c4a7a09c45a23896781b1fc6e7a63b0403f8f6befb7915609b8fca9f1152b31"></a>

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

<a id="canonical-de816ff2afb5821a9a3583d7e8732eeb68e53629172480b79346e3bbc906b7b1"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 98dba862908e / 3

<a id="canonical-36edc27e3d022d1bc6f78018b2c83d8e65c413aeb296807965c60ca947f04eaf"></a>

<a id="canonical-597dda37e6763e45c5f505a780262f68fc0f8348f60ce64fbd2a2cff28a73e22"></a>

## addr property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 98dba862908e / 4

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

<a id="canonical-6790914683c77596c8b6da486e49ac47b62b6e6c5fba557ae7a74f0f7f58981c"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 98dba862908e / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0688cfd28799e65a655441e83cce1b0aa2819716d563d88dfa722b05644f1ae5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e8cdf162b73dfd2acbbb0c881dff8a657ed5b5725e1939b57bbc734313a49a0a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-01ced5489559d1941b718c0cc5e9bed2b6426b7bd0c9031e3f5f72271c9cad75"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6 — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / a03772c1b3b5 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-043c0366474cef1dc77a65d1ad5f31a575f993995c1204159be39b06a0ca7f6e)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0688cfd28799e65a655441e83cce1b0aa2819716d563d88dfa722b05644f1ae5)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6

<a id="canonical-ce7bfad249abf0bd52d782396ea472fdf0a20724c3c82281651e8eae8ee51429"></a>

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

<a id="canonical-ea2836d690cc4a593ddb1856244e80f0c2d101bb08746ce690c91919c0c1b800"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / a03772c1b3b5 / 3

<a id="canonical-c8b68115e21e75ec6f0c88836722e80740c2f865380a3c15829ff18a4b76e4d8"></a>

<a id="canonical-1805566945feadaa8ea91b46266a341237e6db3367883d1db40afbd1b761b018"></a>

## addr property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / a03772c1b3b5 / 4

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

<a id="canonical-52af44d0fb83d3ce231c89753c92d50fa9a695209a8d3769b49f52d037729a9a"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / a03772c1b3b5 / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-007.md#canonical-0688cfd28799e65a655441e83cce1b0aa2819716d563d88dfa722b05644f1ae5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-4eead2df43a0b0ddd1ee2726da95aaa4c666310275a4a49d389ab6a5b2cb5557"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-417301e633686f5f8aba411b43cc609ec36c73ebb90aa7fd4f363cd0f3e6570a"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4 — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 21e3fa14329d / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-043c0366474cef1dc77a65d1ad5f31a575f993995c1204159be39b06a0ca7f6e)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4

<a id="canonical-7b7bdb61d7b7fc4d39184b418f5550efd7e63ee152c1bae6bf9fc0939ab09165"></a>

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

<a id="canonical-a9a01d2c1b9180176e052ae7da298f8fb4c1650222eb3014eda771416f8b0c24"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 21e3fa14329d / 3

<a id="canonical-6d381bb1faf732b9092437b4fd19749ca487fd6b09f219275e56ed7dd46ae225"></a>

<a id="canonical-771a16a1e7f2793f980b7958d07eea72570626971f5faae356be2d0fd3f413d6"></a>

## addr property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 21e3fa14329d / 4

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

<a id="canonical-edca5adba551a96f327d3158b34eed12b2e2ad6142e308735aad1e4fc41b7592"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 21e3fa14329d / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-043c0366474cef1dc77a65d1ad5f31a575f993995c1204159be39b06a0ca7f6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c7a62a9c9b934ea5f42109183058dfb0ee6afaa370d87e6badfe6b34b90b17e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b1fe24491bce78e8a76ddd129728443a9cc5b8d91243e878e0dd8fa79c13668b"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6 — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / a777c0a394f7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-007.md#canonical-b58c3c7097e3bdb46cd6e83332ba86b3a0a361a8a8659b1d5e7906fc0235a973)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-043c0366474cef1dc77a65d1ad5f31a575f993995c1204159be39b06a0ca7f6e)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6

<a id="canonical-d6bdf60cce03ac91ac6df701d5075fabbe0f1195095b4786de49833365b4757b"></a>

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

<a id="canonical-cb2beb293b848e919e58a1751ac56218f0a3d145f430811b8d84aa11a5a7d043"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / a777c0a394f7 / 3

<a id="canonical-6655bea155cb743ef260c1918f253b4054f3237495ef9a76cd2dbf887e057ccb"></a>

<a id="canonical-a193e55925dbbb390801b6ee1c377884358e8fd903c5326bf1cbe14839322c0c"></a>

## addr property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / a777c0a394f7 / 4

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

<a id="canonical-2d5982c12c5aa0590efb410e02387a92350dc726af8208b4ce176f9203331444"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / a777c0a394f7 / 5

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-007.md#canonical-043c0366474cef1dc77a65d1ad5f31a575f993995c1204159be39b06a0ca7f6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-bbbacacda6edab527e222102f575fd2323276b60e8381926d7c519230cf6f8cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0e7d48edad6eb283c208349771233c10ba745a28d34273983def7724411a1393"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 0c10c1f5b3ef / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-5df55b06b5a268277134957ac836f52834c5770a13effd31c8e68527d5a586c7"></a>

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

<a id="canonical-711af76a65e9c4b9eab73ecefb8af0d56f91340a80c4613ebb67b784902b2c05"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 0c10c1f5b3ef / 3

- [ipv4](resources--azure_vnet_site--reference--group-007.md#canonical-28395f0cdd3f6eb37428587d77bd2e0a12b645e318d44033ac3385d43cc9b46f): complete subsection reference.

- [ipv6](resources--azure_vnet_site--reference--group-007.md#canonical-44da6bdcd17bdc59f34a1b48cc5bb75bad35246170db3c44758a6227c1c67d30): complete subsection reference.

<a id="canonical-1dcec7bdc35574af64cef4fc3fd9fc475c97c96db018540fe103b1caa48664a5"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 0c10c1f5b3ef / 4

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--azure_vnet_site--reference--group-007.md#canonical-28395f0cdd3f6eb37428587d77bd2e0a12b645e318d44033ac3385d43cc9b46f)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--azure_vnet_site--reference--group-007.md#canonical-44da6bdcd17bdc59f34a1b48cc5bb75bad35246170db3c44758a6227c1c67d30)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-28395f0cdd3f6eb37428587d77bd2e0a12b645e318d44033ac3385d43cc9b46f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7e18059ea7c5073f60fb099a7f54b1f7007b42f58aec99826e156e0447acb42a"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4 — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / cfa0a3856b81 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-007.md#canonical-bbbacacda6edab527e222102f575fd2323276b60e8381926d7c519230cf6f8cb)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4

<a id="canonical-0c17a3275889a54bd0964ae1a981cbb8a1660b1daf2a8ac7767697a76838aa60"></a>

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

<a id="canonical-5e3fdb46cb4dc18537f8129693a055c6f58954c97a337336f3960a5ae4f61082"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / cfa0a3856b81 / 3

<a id="canonical-cea1e12dc164a5b0fbb9336b5f25c2a3af5e4ba3a26604def4ca45dbeb7e87ed"></a>

<a id="canonical-322d3be89460c4c02ab7d1e1990ed83921badbf1d8d0d539dff89a8a74e29595"></a>

## plen property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / cfa0a3856b81 / 4

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

<a id="canonical-52ce0522e1db5fc8673ee59896f40b9f2b6fd55371483bc20d4205dbc028341c"></a>

<a id="canonical-817bb64dad441872f7abe8eb9bb8cf5e499f9e61e4b90d6cd559792bb70407f3"></a>

## prefix property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / cfa0a3856b81 / 5

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

<a id="canonical-0615a582c7192d3d24f44f1528ba830a7e2ec24ec85e1827c18a096dd132d65c"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / cfa0a3856b81 / 6

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-007.md#canonical-bbbacacda6edab527e222102f575fd2323276b60e8381926d7c519230cf6f8cb)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-44da6bdcd17bdc59f34a1b48cc5bb75bad35246170db3c44758a6227c1c67d30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1add18a48e43ae6a1268eb5dced19475d15f43c3bd221ad7e828bd89f60e14ed"></a>

## ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6 — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 398e402f7385 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.outside_static_routes](resources--azure_vnet_site--reference--group-007.md#canonical-d4661b99ff75462b7de04e97cd1a4927e7f64c6a2d17591ea45208dc2deee18a)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-007.md#canonical-52a57dce9ba3ccdd63382078f1535c97e53fae2d570dd34d71ef10bbdce00c58)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-007.md#canonical-2068ad8cf12a58d419d567073ff5bae93024b1293008131ab989cc88681b2afe)
- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-007.md#canonical-bbbacacda6edab527e222102f575fd2323276b60e8381926d7c519230cf6f8cb)
- ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6

<a id="canonical-4bad91a1ae41b09bb48d524fbaae70126edcbb4a09a8b73ecd15e4b17adb1b49"></a>

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

<a id="canonical-2d446d8bca1e5eb9f7ef167cb4b503551977ac788bacaa65c781b3d4d9fa0d93"></a>

## Direct properties — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 398e402f7385 / 3

<a id="canonical-edea51cc57e7fbdba12613c485807ad4115e31559757d2abe750e01a28de79da"></a>

<a id="canonical-15e83976c38d9116ecd4589a5168f0dfaf7d150988af0f50fa45f8fd8ad91d5b"></a>

## plen property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 398e402f7385 / 4

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

<a id="canonical-333172e5b145e8098213fff5b6e0c698f5d9c0df62b86c7f7aa8aee6b0ebdf1a"></a>

<a id="canonical-9fc7426882b85244a13df7aa2ee5c590a37a65fb8308d9a3ba6e865f7d2e70da"></a>

## prefix property — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 398e402f7385 / 5

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

<a id="canonical-f36c40719812b2ca82f2e473e5f2dfdd3fed2b0b5d941969dee59958a29d5898"></a>

## Next pages — ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route / 398e402f7385 / 6

- [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-007.md#canonical-bbbacacda6edab527e222102f575fd2323276b60e8381926d7c519230cf6f8cb)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-8c1fe208b21819c9e2a17ea2396879bed1b8261c750b75f44a9a953ef70031cb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16f425520555d8db0693bfaa51488b6c9baee65987d28235148a7058feace252"></a>

## ingress_egress_gw_ar.performance_enhancement_mode — ingress_egress_gw_ar.performance_enhancement_mode / c41735badd3a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.performance_enhancement_mode

<a id="canonical-21e42d12986bd63f86ebf801dae34bdbd4eec94be4eabaafb5642a1ec403b16f"></a>

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

<a id="canonical-d906dfa11d122981617b4eb6b7c84cb519539dbfaee80939cb373758e95dfc8f"></a>

## Direct properties — ingress_egress_gw_ar.performance_enhancement_mode / c41735badd3a / 3

- [perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9d0edebfe0252a8c65f0b8f1af977821e3ba84afa452e09a01085f2db19068b5): complete subsection reference.

- [perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-e31c8839632bd8b2d313bb19d2069cb7217e9aa95c313039f8e6b7dc62d6b0f6): complete subsection reference.

<a id="canonical-e1c9a1537e6d3ea9ec6b4993262c54f38a1e29ea8eb4cf4e124828e048b783a2"></a>

## Next pages — ingress_egress_gw_ar.performance_enhancement_mode / c41735badd3a / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9d0edebfe0252a8c65f0b8f1af977821e3ba84afa452e09a01085f2db19068b5)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-e31c8839632bd8b2d313bb19d2069cb7217e9aa95c313039f8e6b7dc62d6b0f6)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9d0edebfe0252a8c65f0b8f1af977821e3ba84afa452e09a01085f2db19068b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-396fbb5edb2a9dbeddcf46b17458a06dce0c60da1054d5f439cf4bcfaa7a9378"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced / ceb6ab5bdb24 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-8c1fe208b21819c9e2a17ea2396879bed1b8261c750b75f44a9a953ef70031cb)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-c76bfc97d685647ea4c4a27776d25441b9a64a4cf03255f8b5cd2d27d6bbad69"></a>

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

<a id="canonical-79914e7af758b674d7715b18b36c9646b17c19c1df5a7370e407d8e376fae84c"></a>

## Direct properties — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced / ceb6ab5bdb24 / 3

- [jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-7c723c01d63699cb2aedf863ca5b973b8c286926de2492c646e45dbb03bc14b6): complete subsection reference.

- [no_jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-78b52eeddeee0f05b55144709767c237a6f1061bb6a2f94b12112a43e2ecd978): complete subsection reference.

<a id="canonical-221cbe0e8e416a51f3563c9801000d763579fec8ba1ffea2b7affdbc04ab7246"></a>

## Next pages — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced / ceb6ab5bdb24 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-7c723c01d63699cb2aedf863ca5b973b8c286926de2492c646e45dbb03bc14b6)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-78b52eeddeee0f05b55144709767c237a6f1061bb6a2f94b12112a43e2ecd978)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-8c1fe208b21819c9e2a17ea2396879bed1b8261c750b75f44a9a953ef70031cb)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-7c723c01d63699cb2aedf863ca5b973b8c286926de2492c646e45dbb03bc14b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06b1ad1597378f42b01fbe0e1491e73da6d4323319791131d660fb6f7e405e00"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / b8642b7ea023 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-8c1fe208b21819c9e2a17ea2396879bed1b8261c750b75f44a9a953ef70031cb)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9d0edebfe0252a8c65f0b8f1af977821e3ba84afa452e09a01085f2db19068b5)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-da6fef098185284c2a82666d6a61d58085e4d4929c3551706fb549411c0224b0"></a>

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

<a id="canonical-724fb3a805275d35d412af24e58e247e5854114a56b98f8edeafac15cca55977"></a>

## Direct properties — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / b8642b7ea023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ec265a18d130fc9b846682e65e72d35d19a95c316fb345bd01c6699d613cd8e1"></a>

## Next pages — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / b8642b7ea023 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9d0edebfe0252a8c65f0b8f1af977821e3ba84afa452e09a01085f2db19068b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-78b52eeddeee0f05b55144709767c237a6f1061bb6a2f94b12112a43e2ecd978"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38e4f4f591ec6fc8244ef141c4e54970e1de9c8e4feb15d55ab8ba2cc39f05ca"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / dad0ecc1a4d1 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-8c1fe208b21819c9e2a17ea2396879bed1b8261c750b75f44a9a953ef70031cb)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9d0edebfe0252a8c65f0b8f1af977821e3ba84afa452e09a01085f2db19068b5)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-dcef7cbde7afdb31015eb7d9be03023510536070155b1c9df6a64daf1cfab538"></a>

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

<a id="canonical-4d8ac6c17db714f7e4cf456773632d5337137f868928de3950d7d35deb4e445d"></a>

## Direct properties — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / dad0ecc1a4d1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f7878d822d14e058e8a88f925612a0e5cd6859fb2f97fe428d310b37ca92399a"></a>

## Next pages — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / dad0ecc1a4d1 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9d0edebfe0252a8c65f0b8f1af977821e3ba84afa452e09a01085f2db19068b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e31c8839632bd8b2d313bb19d2069cb7217e9aa95c313039f8e6b7dc62d6b0f6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c3801321dc25b91216f8023c49573fe03a599f6fb235c645ccbec4623406b8c6"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced / 4ac029236d94 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-8c1fe208b21819c9e2a17ea2396879bed1b8261c750b75f44a9a953ef70031cb)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-0241cf589fc8e87f09b8415a6d77c9c708d1c673f8736257251a4e07677897c3"></a>

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

<a id="canonical-c0fe7806b31e633d9bc85ba499298f4aa8efa3c1488a09ec6aa2a10caf22e928"></a>

## Direct properties — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced / 4ac029236d94 / 3

- [jumbo_disabled](resources--azure_vnet_site--reference--group-007.md#canonical-e597328b64adb102a475725eb27c18a5f0d728a9994d20cd14bd42057f34f64f): complete subsection reference.

- [jumbo_enabled](resources--azure_vnet_site--reference--group-007.md#canonical-e978a5fbcb806312eb08923acf71b8de68a93534b0af3453e4f54f8132f9d8a8): complete subsection reference.

<a id="canonical-4fee6307dd3e633f004b118196443c2e8e9a8fa1cdf88f430f1ab54a916b7e36"></a>

## Next pages — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced / 4ac029236d94 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--azure_vnet_site--reference--group-007.md#canonical-e597328b64adb102a475725eb27c18a5f0d728a9994d20cd14bd42057f34f64f)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--azure_vnet_site--reference--group-007.md#canonical-e978a5fbcb806312eb08923acf71b8de68a93534b0af3453e4f54f8132f9d8a8)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-8c1fe208b21819c9e2a17ea2396879bed1b8261c750b75f44a9a953ef70031cb)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e597328b64adb102a475725eb27c18a5f0d728a9994d20cd14bd42057f34f64f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fc2c40322034f86ac793975570fb8b3febee3f78bf4b4928f6ae673040cda49"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_di / 6acc384f4574 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-8c1fe208b21819c9e2a17ea2396879bed1b8261c750b75f44a9a953ef70031cb)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-e31c8839632bd8b2d313bb19d2069cb7217e9aa95c313039f8e6b7dc62d6b0f6)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-144973992151e3c8ea359dbc215192ce60f290c6a3195aaa33f4420673c745f6"></a>

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

<a id="canonical-12fe2351b485d10a971e9fae42819282be7c45aa68b712ebb4efafe1b169f15e"></a>

## Direct properties — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_di / 6acc384f4574 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-606497cd40349820c1981d3a899a3cecd1197224dd525267a92922e526a52b12"></a>

## Next pages — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_di / 6acc384f4574 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-e31c8839632bd8b2d313bb19d2069cb7217e9aa95c313039f8e6b7dc62d6b0f6)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e978a5fbcb806312eb08923acf71b8de68a93534b0af3453e4f54f8132f9d8a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-263a8cf337a48011433eae9708148033075a9a7953d11d703287cd56585d1b82"></a>

## ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_en / 64e62e5c8822 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [ingress_egress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-8c1fe208b21819c9e2a17ea2396879bed1b8261c750b75f44a9a953ef70031cb)
- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-e31c8839632bd8b2d313bb19d2069cb7217e9aa95c313039f8e6b7dc62d6b0f6)
- ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-a5bdaf86951e2bf2e2f37f2990a3675772987841d8adfe22b7da1306af492d3b"></a>

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

<a id="canonical-5e0a2d52517457517cd5e83724330af1c0518ff27b1695fb5c62f1ca7a5ccd6c"></a>

## Direct properties — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_en / 64e62e5c8822 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4ad70ace6fbd6cf9ea3daed5f7e4124a092baff08a301291c685595e28cbe230"></a>

## Next pages — ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_en / 64e62e5c8822 / 4

- [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-e31c8839632bd8b2d313bb19d2069cb7217e9aa95c313039f8e6b7dc62d6b0f6)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-42a38e5c411a6ae53043b919a491382350a63697848e7b49f4113ae2a01abcce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d5ea8326e3a91227c4b4a6860bb042f8196195876c2e09c94a1ce7c507d6d324"></a>

## ingress_egress_gw_ar.sm_connection_public_ip — ingress_egress_gw_ar.sm_connection_public_ip / d03d38482058 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.sm_connection_public_ip

<a id="canonical-bdeb7381104fa3cf1082c873ed7ba70d544f2e94c2ae6dcedddf7aeda263ca76"></a>

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

<a id="canonical-61905c991179d0f5a9854bf1180c429fd9c5c9e2b99991fa5eb5246816cfa39f"></a>

## Direct properties — ingress_egress_gw_ar.sm_connection_public_ip / d03d38482058 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ad5cf9f193a92777d91ec746038799a4d38db627da0f0872d05c57cefb2555bc"></a>

## Next pages — ingress_egress_gw_ar.sm_connection_public_ip / d03d38482058 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-33bb71f61320605de6c5a00f1a99a858c225d67c17dc22d19f49e655eb832e8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f2c9e85842f67d69eb5abad541b3604ee71d919e6b3d45b0335774fce31fa079"></a>

## ingress_egress_gw_ar.sm_connection_pvt_ip — ingress_egress_gw_ar.sm_connection_pvt_ip / 2c733af737e2 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- ingress_egress_gw_ar.sm_connection_pvt_ip

<a id="canonical-d6ebb138f2b59f07db53fa2c3b723acec46c586417df4f7e4b61e2aec1bcb458"></a>

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

<a id="canonical-9e9ef2cc9bad5bae94f80e6874a85ad2040dc2ddc9bb7a556a2e525df24fee44"></a>

## Direct properties — ingress_egress_gw_ar.sm_connection_pvt_ip / 2c733af737e2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-66f367bc02004a79be0b62d0e639b0474fb8c0f1f6f627a1760434f37e65bc92"></a>

## Next pages — ingress_egress_gw_ar.sm_connection_pvt_ip / 2c733af737e2 / 4

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-361833d2e4de66232cf486000c668818ab0f26d3ae5534c58b5a902626ff71b5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bffd21ec8a74b4a9c2224731db5130b8a157b5b71f056b830b6d403662e932bc"></a>

## ingress_gw — ingress_gw / 33bf5be276e9 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- ingress_gw

<a id="canonical-2f3855ab8b16d34ab67bedb123911fa6d3732b53f1fd1ab9c79a82650334425c"></a>

Type: `"object"`. single nested block, Optional.

Single interface Azure ingress site on on Recommended Region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("az_nodes",
    "azure_certified_hw")}
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

<a id="canonical-39fdd78b2f4e1593cd84e3c019da45af61f8cdd753e425ffb0ba0c7dd038044c"></a>

## Direct properties — ingress_gw / 33bf5be276e9 / 3

- [accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-1dd4200d6041bf1abc9c6018da534a7dee663ae4896a5ac1419e1f24651e68f8): complete subsection reference.

- [az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-bc51a02e2a1201e9e35c9f5463865a82d2609892380ed815365fbf7e7b072a49): complete subsection reference.

<a id="canonical-c2e1a9d3745071bf3d74975133e7d6d474d78aba28823fee9f15a1d2bb1ae2a5"></a>

<a id="canonical-2969438acd83f5d7a7bf240b96fc81aa215de4ef2bea4046aec9ab8fa979cf82"></a>

## azure_certified_hw property — ingress_gw / 33bf5be276e9 / 4

Type: `"string"`. Optional.

\[Enum: azure-byol-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware. The only
possible value is \`azure-byol-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("azure-byol-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-voltmesh"
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
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-7f94478a0dcbf76170886f63ba64aa5a5a4fd5a1781885c2b575ef9668641227): complete subsection reference.

<a id="canonical-69bdfe8530563f5789faa1cc798dca2e9143d65d45d6f5329a2c9edb76543951"></a>

## Next pages — ingress_gw / 33bf5be276e9 / 5

- [ingress_gw.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-1dd4200d6041bf1abc9c6018da534a7dee663ae4896a5ac1419e1f24651e68f8)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-bc51a02e2a1201e9e35c9f5463865a82d2609892380ed815365fbf7e7b072a49)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-7f94478a0dcbf76170886f63ba64aa5a5a4fd5a1781885c2b575ef9668641227)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-1dd4200d6041bf1abc9c6018da534a7dee663ae4896a5ac1419e1f24651e68f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4a7870e881a865a46ec0c9e708b9d05807059ac3a4b0706758a7aa1f3af845bb"></a>

## ingress_gw.accelerated_networking — ingress_gw.accelerated_networking / 39873dd7e30a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- ingress_gw.accelerated_networking

<a id="canonical-65f7c5c06332fbc533ffd2f9fad7f71821224aeff812d5f118549b9cfaa1685d"></a>

Type: `"object"`. single nested block, Optional.

Accelerated Networking to reduce Latency, When Mode is toggled, traffic disruption will be seen.
Server applies default when omitted.

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

<a id="canonical-5a49952c1c765837e9a7e04ae704a4d478d6c1bcce2bb6612fccf920e003f77d"></a>

## Direct properties — ingress_gw.accelerated_networking / 39873dd7e30a / 3

- [disable_spec](resources--azure_vnet_site--reference--group-007.md#canonical-6aeef8df39e0d88517afbad2e96a6fa6af44c502d94ecb5fcd97d5d998d625fe): complete subsection reference.

- [enable](resources--azure_vnet_site--reference--group-007.md#canonical-b6a816f22f2d0854de80486cb6e670ea2d5aed524f739fca35c208c637920ef7): complete subsection reference.

<a id="canonical-4768b728048a1c31b2574b5c09c4602dd289dab8f484afb5b08b20e62e4778b4"></a>

## Next pages — ingress_gw.accelerated_networking / 39873dd7e30a / 4

- [ingress_gw.accelerated_networking.disable_spec](resources--azure_vnet_site--reference--group-007.md#canonical-6aeef8df39e0d88517afbad2e96a6fa6af44c502d94ecb5fcd97d5d998d625fe)
- [ingress_gw.accelerated_networking.enable](resources--azure_vnet_site--reference--group-007.md#canonical-b6a816f22f2d0854de80486cb6e670ea2d5aed524f739fca35c208c637920ef7)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6aeef8df39e0d88517afbad2e96a6fa6af44c502d94ecb5fcd97d5d998d625fe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59262da0372cb24949f8f71305ed2b86e38d59adcd71fdc793a88c2001c16beb"></a>

## ingress_gw.accelerated_networking.disable_spec — ingress_gw.accelerated_networking.disable_spec / 86825a564292 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-1dd4200d6041bf1abc9c6018da534a7dee663ae4896a5ac1419e1f24651e68f8)
- ingress_gw.accelerated_networking.disable_spec

<a id="canonical-3c1821c5f47d3a952c56c9a189136ce818fe95548ea8f0aadc9f9b5517e1fdb4"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-b0dc6261f4af1702fcc052e1e92fcc72da1a92f4e630248eeafcb12dddfedd9a"></a>

## Direct properties — ingress_gw.accelerated_networking.disable_spec / 86825a564292 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fe315d65de667e8048699ce27cca5eaeab43355373691a525c09fada916ef5f2"></a>

## Next pages — ingress_gw.accelerated_networking.disable_spec / 86825a564292 / 4

- [ingress_gw.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-1dd4200d6041bf1abc9c6018da534a7dee663ae4896a5ac1419e1f24651e68f8)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b6a816f22f2d0854de80486cb6e670ea2d5aed524f739fca35c208c637920ef7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-02369a81f77c24463e9cfb69548a2d3648e5d6a3b0a852ab78b460df531a1166"></a>

## ingress_gw.accelerated_networking.enable — ingress_gw.accelerated_networking.enable / e316874c095b / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-1dd4200d6041bf1abc9c6018da534a7dee663ae4896a5ac1419e1f24651e68f8)
- ingress_gw.accelerated_networking.enable

<a id="canonical-1d8496a8b7ebc6f8534eb13c5063567b050b9983864433b3b903cb47a8f4386e"></a>

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

<a id="canonical-caa95af411c2e7d45284f2e3fa39b61050454e0a075890b10e3c0ef9ab461a7b"></a>

## Direct properties — ingress_gw.accelerated_networking.enable / e316874c095b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ddc7015478a144d33619f290ea671b48c6667f6899197309fc39c5ace3f35338"></a>

## Next pages — ingress_gw.accelerated_networking.enable / e316874c095b / 4

- [ingress_gw.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-1dd4200d6041bf1abc9c6018da534a7dee663ae4896a5ac1419e1f24651e68f8)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-bc51a02e2a1201e9e35c9f5463865a82d2609892380ed815365fbf7e7b072a49"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7eab02f468d556f5765974302ab2eeb1c9e5b20690f03f0c143fecb92fe5a36"></a>

## ingress_gw.az_nodes — ingress_gw.az_nodes / 83c7b02be15a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- ingress_gw.az_nodes

<a id="canonical-ab0a2fd489b7c252407d7afff81027641dd35a377c588f530f4dcee190b5bafe"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("azure_az")}
```

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
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

Terraform syntax:

```terraform
az_nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-7bb3ac795d257b8d3a01bae40fa9fad49454b0a8ab98dfab1f70f7cf9007af39"></a>

## Direct properties — ingress_gw.az_nodes / 83c7b02be15a / 3

<a id="canonical-ec110d89292e50ed8334d808894a2d239e2108b5df8067860b57f1292555af4e"></a>

<a id="canonical-aa2c6b3c632d16d2364230a04b9f4824d33b4698fe28fb1ac2f087edfe9d25b1"></a>

## azure_az property — ingress_gw.az_nodes / 83c7b02be15a / 4

Type: `"string"`. Optional.

\[Enum: 1|2|3\] Zone depicting a grouping of datacenters within an Azure region. Expecting numeric
input. Possible values are \`1\`, \`2\`, \`3\`.

Upstream description:

A zone depicting a grouping of datacenters within an Azure region. Expecting numeric input.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("1",
    "2",
    "3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "1",
    "2",
    "3"
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"1\\\",\\\"2\\\",\\\"3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"1\\\",\\\"2\\\",\\\"3\\\"]"
  }
}
```

- [local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c5223a3021cd9711367595d0d975c35cc89351f8efb58c93c5dbec8eda8362c5): complete subsection reference.

<a id="canonical-f691361afaf3b6957d077120aa03686d50e995db3a022238a887555fd048903c"></a>

## Next pages — ingress_gw.az_nodes / 83c7b02be15a / 5

- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c5223a3021cd9711367595d0d975c35cc89351f8efb58c93c5dbec8eda8362c5)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c5223a3021cd9711367595d0d975c35cc89351f8efb58c93c5dbec8eda8362c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-409e4839c22470980d267d2419903559d086316270dd017b592ee6a42df8bfbf"></a>

## ingress_gw.az_nodes.local_subnet — ingress_gw.az_nodes.local_subnet / 3f8c87e66be9 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-bc51a02e2a1201e9e35c9f5463865a82d2609892380ed815365fbf7e7b072a49)
- ingress_gw.az_nodes.local_subnet

<a id="canonical-8b6cf4da67ceb660dfae9ffee88e38556a2bd6d05ea9710ae1d281f9c34ea0e2"></a>

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

<a id="canonical-3f6226149b434e060f2a74ec892fd0347ac2f245faee1f8eda065606e7f9c32e"></a>

## Direct properties — ingress_gw.az_nodes.local_subnet / 3f8c87e66be9 / 3

- [subnet](resources--azure_vnet_site--reference--group-007.md#canonical-a90d7dc570555ee984d2d3ce924158d5bb7a6ed3c35db26a5f544f8b9df3c436): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-6bccd9def60433af31a443bd59e5573dfb61b21bf4051dca164fa1a75d0894aa): complete subsection reference.

<a id="canonical-0b2f0722b1935f59f70711f24d96af64dfa7187bce1e766cab53ea9f8d7dd436"></a>

## Next pages — ingress_gw.az_nodes.local_subnet / 3f8c87e66be9 / 4

- [ingress_gw.az_nodes.local_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-a90d7dc570555ee984d2d3ce924158d5bb7a6ed3c35db26a5f544f8b9df3c436)
- [ingress_gw.az_nodes.local_subnet.subnet_param](resources--azure_vnet_site--reference--group-007.md#canonical-6bccd9def60433af31a443bd59e5573dfb61b21bf4051dca164fa1a75d0894aa)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-bc51a02e2a1201e9e35c9f5463865a82d2609892380ed815365fbf7e7b072a49)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a90d7dc570555ee984d2d3ce924158d5bb7a6ed3c35db26a5f544f8b9df3c436"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-220682d744852e5d2e15d3df37ca59745d9cb62a83d3a72a1f251e3d4213a589"></a>

## ingress_gw.az_nodes.local_subnet.subnet — ingress_gw.az_nodes.local_subnet.subnet / 8ed9f542af88 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-bc51a02e2a1201e9e35c9f5463865a82d2609892380ed815365fbf7e7b072a49)
- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c5223a3021cd9711367595d0d975c35cc89351f8efb58c93c5dbec8eda8362c5)
- ingress_gw.az_nodes.local_subnet.subnet

<a id="canonical-acc0292419ae5bbb8f1fd238aaee85ab0f81f2467c5259e5d29fd5ecbb16253f"></a>

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

<a id="canonical-a8b55aa6f32947c2f3816c2c9e51704e5c10f705271c2413212593c94e6b47fe"></a>

## Direct properties — ingress_gw.az_nodes.local_subnet.subnet / 8ed9f542af88 / 3

<a id="canonical-366d2ce1e6582d7c12384692c39f107951ec3379dd82a544cae645f58d4e083b"></a>

<a id="canonical-561f991c5a2ed3f303fca8d79ddcc4b6f657fc813d7bdb6a384a9c5b9229d971"></a>

## subnet_name property — ingress_gw.az_nodes.local_subnet.subnet / 8ed9f542af88 / 4

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

<a id="canonical-fd8d67db0b39f1650a92444a77ebc0527cbd061f59c989e1c5edf03522ce676e"></a>

<a id="canonical-f4c5875a2c1dfe459313876c79b831b7a558404de7ada26b59e167363769fe80"></a>

## subnet_resource_grp property — ingress_gw.az_nodes.local_subnet.subnet / 8ed9f542af88 / 5

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

- [vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-2f93be76c7cf9596ae7c26c72cc1ab1e1c7dcd78d5fdb5a344fb93dda3636f39): complete subsection reference.

<a id="canonical-554a33e4c2fa835e6455b560f694908775aa2471bb6ef496ac4ee2d10f3c043c"></a>

## Next pages — ingress_gw.az_nodes.local_subnet.subnet / 8ed9f542af88 / 6

- [ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-007.md#canonical-2f93be76c7cf9596ae7c26c72cc1ab1e1c7dcd78d5fdb5a344fb93dda3636f39)
- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c5223a3021cd9711367595d0d975c35cc89351f8efb58c93c5dbec8eda8362c5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-2f93be76c7cf9596ae7c26c72cc1ab1e1c7dcd78d5fdb5a344fb93dda3636f39"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-223f6082b25e58b4b62c8139678304d4c20028d9e3d5aab8be7d837d11acb137"></a>

## ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group — ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group / 16178e09cfdc / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-bc51a02e2a1201e9e35c9f5463865a82d2609892380ed815365fbf7e7b072a49)
- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c5223a3021cd9711367595d0d975c35cc89351f8efb58c93c5dbec8eda8362c5)
- [ingress_gw.az_nodes.local_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-a90d7dc570555ee984d2d3ce924158d5bb7a6ed3c35db26a5f544f8b9df3c436)
- ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group

<a id="canonical-3f02ef238dc2119c58d92f496f5da368f53f14e07ff9b02c2e81f77ea5b6465f"></a>

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

<a id="canonical-c1e750eaa5670d2c2618a2b11a50961325c90db75aba065b5348854d25cc5aca"></a>

## Direct properties — ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group / 16178e09cfdc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0ede5903e14b2f6e0f8254d6c16ad6c130f50380007b597fedb40685b94f36a3"></a>

## Next pages — ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group / 16178e09cfdc / 4

- [ingress_gw.az_nodes.local_subnet.subnet](resources--azure_vnet_site--reference--group-007.md#canonical-a90d7dc570555ee984d2d3ce924158d5bb7a6ed3c35db26a5f544f8b9df3c436)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-6bccd9def60433af31a443bd59e5573dfb61b21bf4051dca164fa1a75d0894aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc069b8e4c8d757fd12ea464d338cb8d0ecfc2b8ca5b3bb963938e4a3efcf6e5"></a>

## ingress_gw.az_nodes.local_subnet.subnet_param — ingress_gw.az_nodes.local_subnet.subnet_param / c684ed45c3f7 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.az_nodes](resources--azure_vnet_site--reference--group-007.md#canonical-bc51a02e2a1201e9e35c9f5463865a82d2609892380ed815365fbf7e7b072a49)
- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c5223a3021cd9711367595d0d975c35cc89351f8efb58c93c5dbec8eda8362c5)
- ingress_gw.az_nodes.local_subnet.subnet_param

<a id="canonical-fbdb0b144d7f2bd595f965f0ecad09e0d945b54f4bf0997b6419ff3db9205302"></a>

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

<a id="canonical-634f9a3b6f5d2986e2f584d3c9e90ddd60df825c55f022a9d7a7102b4a3d0940"></a>

## Direct properties — ingress_gw.az_nodes.local_subnet.subnet_param / c684ed45c3f7 / 3

<a id="canonical-d8dbe9a752065bc1bbc0b2922a21e9aaf941abe0b32cbe3140240aad98eae0ed"></a>

<a id="canonical-284599583ebaafbf816b2e4f5d42c3e5a91a3c82730d6fefab5f9001befb1b2e"></a>

## ipv4 property — ingress_gw.az_nodes.local_subnet.subnet_param / c684ed45c3f7 / 4

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

<a id="canonical-f755e8655415e95df1551e1a5ed2520555e366160abc808c260d70a3b14f0d47"></a>

## Next pages — ingress_gw.az_nodes.local_subnet.subnet_param / c684ed45c3f7 / 5

- [ingress_gw.az_nodes.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c5223a3021cd9711367595d0d975c35cc89351f8efb58c93c5dbec8eda8362c5)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-7f94478a0dcbf76170886f63ba64aa5a5a4fd5a1781885c2b575ef9668641227"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-92e9f49c9726654e9974ccd254b2cfd9a28ef4be16772b5498c41621a41813da"></a>

## ingress_gw.performance_enhancement_mode — ingress_gw.performance_enhancement_mode / 31e44d38bb30 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- ingress_gw.performance_enhancement_mode

<a id="canonical-2b9106fa7fa882455474123b2919107bac85db2e450b841a23315f5fb29bf9d3"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default. Server applies
default when omitted.

Upstream description:

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

<a id="canonical-cb540bd0363af14f8ceaafff35b6f391e57744da472476cbfe2a5210bdcfebd6"></a>

## Direct properties — ingress_gw.performance_enhancement_mode / 31e44d38bb30 / 3

- [perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-d0eeb62ba3621d9974a33b99e92fb988a83001a98d65a2c866833a88e9d55093): complete subsection reference.

- [perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9ebd3586aa53eb97704d1e7f4f41a2e1ac16ae4438c152769a4c4bec39e85c9e): complete subsection reference.

<a id="canonical-c6473336a475accb278aca4d31be9543292baeb10b13e80361974c2c1c0dd905"></a>

## Next pages — ingress_gw.performance_enhancement_mode / 31e44d38bb30 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-d0eeb62ba3621d9974a33b99e92fb988a83001a98d65a2c866833a88e9d55093)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9ebd3586aa53eb97704d1e7f4f41a2e1ac16ae4438c152769a4c4bec39e85c9e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d0eeb62ba3621d9974a33b99e92fb988a83001a98d65a2c866833a88e9d55093"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1be47ac7016fb3c45ae53ebdd2d5a4dbe21054e3a4099f9835261b3c3ce7161f"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 611f343a8551 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-7f94478a0dcbf76170886f63ba64aa5a5a4fd5a1781885c2b575ef9668641227)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-95b6f35f39286c148ce39ba52defb70fd53f8012314b582f7e850a034cc1c731"></a>

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

<a id="canonical-d479e24ccc54b459d6520652e57a5836f94324cfbe8bb9c0fbab342cc5252da3"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 611f343a8551 / 3

- [jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-f22f75f423ed2cbf5294cc5290867b0e30c76d2c7e42a7e2428903b3ade43529): complete subsection reference.

- [no_jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-e772c2373d20f2aa5a9d9e347abc8eb340f0ce81505810939812e1000b469b26): complete subsection reference.

<a id="canonical-a88ed2eca410199f16fb73e0ae24c0d110ea967d5ca04b836eef1ad3af548830"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / 611f343a8551 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-f22f75f423ed2cbf5294cc5290867b0e30c76d2c7e42a7e2428903b3ade43529)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--azure_vnet_site--reference--group-007.md#canonical-e772c2373d20f2aa5a9d9e347abc8eb340f0ce81505810939812e1000b469b26)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-7f94478a0dcbf76170886f63ba64aa5a5a4fd5a1781885c2b575ef9668641227)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-f22f75f423ed2cbf5294cc5290867b0e30c76d2c7e42a7e2428903b3ade43529"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33259fa98938d0fe365b07ad25373151cd64eae52d73e4c5e97d782f84a7028e"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / e2cac894246a / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-7f94478a0dcbf76170886f63ba64aa5a5a4fd5a1781885c2b575ef9668641227)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-d0eeb62ba3621d9974a33b99e92fb988a83001a98d65a2c866833a88e9d55093)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-f1a4a2dd6ece068827d61ab6c56d7634306079ca6ca0bd547a56442b966e0b49"></a>

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

<a id="canonical-b740f9c99913c1b04e0f300ec33cd7d586bedc56754fdd41abc0f613a320033f"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / e2cac894246a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-67e4ee7ae650a914f31dfa442f8806af809218c09135f07d173f7d5bc6481627"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / e2cac894246a / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-d0eeb62ba3621d9974a33b99e92fb988a83001a98d65a2c866833a88e9d55093)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-e772c2373d20f2aa5a9d9e347abc8eb340f0ce81505810939812e1000b469b26"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8a4086b81d9b54f8fdb882e40ac46eea10422fbfae7e5eff4b4da6bc1c443ebe"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 00b14ca6d4db / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-7f94478a0dcbf76170886f63ba64aa5a5a4fd5a1781885c2b575ef9668641227)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-d0eeb62ba3621d9974a33b99e92fb988a83001a98d65a2c866833a88e9d55093)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-b594ab468f47075c2e7ba39654f058bf8f235a0fbdb87518a40ade891fb2d69f"></a>

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

<a id="canonical-4be663406703fabd5a88cf2f35f34cccbf988955872e7a348b583f76661a4bd4"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 00b14ca6d4db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f62e161bc72f67fb31e599eee2afcb8bbb340ec49fff612a8f5d8041fe8875f2"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 00b14ca6d4db / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-d0eeb62ba3621d9974a33b99e92fb988a83001a98d65a2c866833a88e9d55093)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-9ebd3586aa53eb97704d1e7f4f41a2e1ac16ae4438c152769a4c4bec39e85c9e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3f9ff62296f5b436a6c3e08696b6c3ba8baae7d1f434e6347f03d6df7ecff74d"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 9caf479262aa / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-7f94478a0dcbf76170886f63ba64aa5a5a4fd5a1781885c2b575ef9668641227)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-2e5373d77a96eeab254ae2a54946f487ab9beadca7907b884e910f2a15d4d516"></a>

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

<a id="canonical-db219c3d1975272e842e009a35b6c63c5980e64ee4e705641818cc6ac93aad00"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 9caf479262aa / 3

- [jumbo_disabled](resources--azure_vnet_site--reference--group-007.md#canonical-093b6c806f92b35847d008f76d37ad592a9fa89662214e36ec9a938058089a79): complete subsection reference.

- [jumbo_enabled](resources--azure_vnet_site--reference--group-007.md#canonical-a871521c884358f31066f8084450291267598a4915bed840d49c08324746263d): complete subsection reference.

<a id="canonical-0700b00b26696addaa035c563cab527e2ae571a51e2b6f41594c3e9ca6f666f1"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / 9caf479262aa / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--azure_vnet_site--reference--group-007.md#canonical-093b6c806f92b35847d008f76d37ad592a9fa89662214e36ec9a938058089a79)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--azure_vnet_site--reference--group-007.md#canonical-a871521c884358f31066f8084450291267598a4915bed840d49c08324746263d)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-7f94478a0dcbf76170886f63ba64aa5a5a4fd5a1781885c2b575ef9668641227)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-093b6c806f92b35847d008f76d37ad592a9fa89662214e36ec9a938058089a79"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d53b7eae46bb59221c7f4cff54396e616abd8986db86adbbba0057ffd07efb35"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 33d21e7206ec / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-7f94478a0dcbf76170886f63ba64aa5a5a4fd5a1781885c2b575ef9668641227)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9ebd3586aa53eb97704d1e7f4f41a2e1ac16ae4438c152769a4c4bec39e85c9e)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-e5e79005e4c0b874a9c23932bd6355f77943ef730d4a9002ae94687bcae52407"></a>

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

<a id="canonical-9a22c52be98a8fc601dabcc61bc4cccb62b327725463adba0cf518f6420fb220"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 33d21e7206ec / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b45e79cca6b895fe4e1968ee874d8aaa6d16a3c4f7fec3bcb6a07ef5d3903d86"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 33d21e7206ec / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9ebd3586aa53eb97704d1e7f4f41a2e1ac16ae4438c152769a4c4bec39e85c9e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a871521c884358f31066f8084450291267598a4915bed840d49c08324746263d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-567a9bbbcb61c13d85991fc2f398cc437e31d91ffeb81ebae9053f908baa9dd4"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 7c96f4e8686c / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-606a1b6abf72175f974cbd00a7a6307f97e16173f014344502dfc9ac1ca14228)
- [ingress_gw.performance_enhancement_mode](resources--azure_vnet_site--reference--group-007.md#canonical-7f94478a0dcbf76170886f63ba64aa5a5a4fd5a1781885c2b575ef9668641227)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9ebd3586aa53eb97704d1e7f4f41a2e1ac16ae4438c152769a4c4bec39e85c9e)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-4e6decf7bf242b869f631455b4acc2009abdc2b9d1509c305b79b045149c7209"></a>

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

<a id="canonical-5bf48dc3dba0e70637e046b9090e30025344a3429c9e0d2423f8cc3c196ef02c"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 7c96f4e8686c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-06e0a4eda9502709b6dd1d5be4b228dfc91ea34830fdfdfc3fc7b0563d73d8d6"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 7c96f4e8686c / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--reference--group-007.md#canonical-9ebd3586aa53eb97704d1e7f4f41a2e1ac16ae4438c152769a4c4bec39e85c9e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f6a9c63fc4d90df091e12cf7968dc677c189081ce350548705630efd55b49568"></a>

## ingress_gw_ar — ingress_gw_ar / 558cad6a0a75 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- ingress_gw_ar

<a id="canonical-b41f84df7f5035ca2597d66f790fb253e09d49063cbd453fee824186373ffb2b"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ingress gw ar.

Upstream description:

Single interface Azure ingress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("azure_certified_hw")}
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
ingress_gw_ar {
  # Configure direct properties listed below.
}
```

<a id="canonical-d0ac8435dcf79788ca77b0c26e34e4bb76427b79ca2c36b6a7ce0ab3a1a380e2"></a>

## Direct properties — ingress_gw_ar / 558cad6a0a75 / 3

- [accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-a00ca808f947d1ddeda9960d2c7134e0e72ebeda1ab93bb3cb4d365c9a8e45de): complete subsection reference.

<a id="canonical-a0bca04dbfc0d1945231818df9b118d4c43b8f6fab5b7e0f12e3bd2cb342cb80"></a>

<a id="canonical-ec3f488ab7667ccbf66deb946deff492a95a73b589e7578e2121da4e5171dd06"></a>

## azure_certified_hw property — ingress_gw_ar / 558cad6a0a75 / 4

Type: `"string"`. Optional.

\[Enum: azure-byol-voltmesh\] Azure Certified Hardware. Name for Azure certified hardware. The only
possible value is \`azure-byol-voltmesh\`.

Upstream description:

Name for Azure certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("azure-byol-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "azure-byol-voltmesh"
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
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"azure-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [node](resources--azure_vnet_site--reference--group-007.md#canonical-87b70f7d188d23ff41722d94e8837588c63ec02ab72058e0fcb45f7d81e73a69): complete subsection reference.

- [performance_enhancement_mode](resources--azure_vnet_site--reference--group-008.md#canonical-3b7cf54e88a203c51b49b91964fd227cf387be5f726e62633846067d9131bbd2): complete subsection reference.

<a id="canonical-ece83ac09c261f63926c52ca64cb36ec828d5c5da0f499722f807d4e7e22c653"></a>

## Next pages — ingress_gw_ar / 558cad6a0a75 / 5

- [ingress_gw_ar.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-a00ca808f947d1ddeda9960d2c7134e0e72ebeda1ab93bb3cb4d365c9a8e45de)
- [ingress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-87b70f7d188d23ff41722d94e8837588c63ec02ab72058e0fcb45f7d81e73a69)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--reference--group-008.md#canonical-3b7cf54e88a203c51b49b91964fd227cf387be5f726e62633846067d9131bbd2)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-a00ca808f947d1ddeda9960d2c7134e0e72ebeda1ab93bb3cb4d365c9a8e45de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca44ba6b39905a2f65fbeea74e78aa13f2be804e160bcb24dfd366c68b175e7e"></a>

## ingress_gw_ar.accelerated_networking — ingress_gw_ar.accelerated_networking / 90e516ece781 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- ingress_gw_ar.accelerated_networking

<a id="canonical-6691a96ae0b0dede067fb16e8a1acf4f5ada626a5d4d9ebd675efef3979f5684"></a>

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

<a id="canonical-485a6fdbd9b3dbe16de3ba54db3618fb6bf0aec74357db9ef96031ec310bb19b"></a>

## Direct properties — ingress_gw_ar.accelerated_networking / 90e516ece781 / 3

- [disable_spec](resources--azure_vnet_site--reference--group-007.md#canonical-b8898560b044945d3ab0d50fe78e51c98cb26420e3e20048f7938f596f2430ca): complete subsection reference.

- [enable](resources--azure_vnet_site--reference--group-007.md#canonical-ff63910373b8299ff769e8e712c001c06539a937670c6d7f4735209233e9ba0c): complete subsection reference.

<a id="canonical-ae867c906d45126f71b4304b659756f3ab6ea6aa88b50509e68a3ca79c1ec218"></a>

## Next pages — ingress_gw_ar.accelerated_networking / 90e516ece781 / 4

- [ingress_gw_ar.accelerated_networking.disable_spec](resources--azure_vnet_site--reference--group-007.md#canonical-b8898560b044945d3ab0d50fe78e51c98cb26420e3e20048f7938f596f2430ca)
- [ingress_gw_ar.accelerated_networking.enable](resources--azure_vnet_site--reference--group-007.md#canonical-ff63910373b8299ff769e8e712c001c06539a937670c6d7f4735209233e9ba0c)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-b8898560b044945d3ab0d50fe78e51c98cb26420e3e20048f7938f596f2430ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f4724d869187d466d75a002ceeab1061960ad732d09223037c31ac46c6e47c3b"></a>

## ingress_gw_ar.accelerated_networking.disable_spec — ingress_gw_ar.accelerated_networking.disable_spec / 9ca46e3a8ae0 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-a00ca808f947d1ddeda9960d2c7134e0e72ebeda1ab93bb3cb4d365c9a8e45de)
- ingress_gw_ar.accelerated_networking.disable_spec

<a id="canonical-c60a73178e948a8ec99413df8b2bfbf4e3b058266592a8f676f7f4c1d6ec2c57"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-ca17952a1b8b1b4e5f0d23c267876223e7e5c3ae51bb408441d8f4775d0ee530"></a>

## Direct properties — ingress_gw_ar.accelerated_networking.disable_spec / 9ca46e3a8ae0 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-00463903afd6c0349c398bf9a4285368398ebf0482d86841eecefe2f93966047"></a>

## Next pages — ingress_gw_ar.accelerated_networking.disable_spec / 9ca46e3a8ae0 / 4

- [ingress_gw_ar.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-a00ca808f947d1ddeda9960d2c7134e0e72ebeda1ab93bb3cb4d365c9a8e45de)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-ff63910373b8299ff769e8e712c001c06539a937670c6d7f4735209233e9ba0c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c0d3287050185447ac585093420d1d654a862d5d719ac6c781c785876f943ad9"></a>

## ingress_gw_ar.accelerated_networking.enable — ingress_gw_ar.accelerated_networking.enable / 5bd4d06e9314 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-a00ca808f947d1ddeda9960d2c7134e0e72ebeda1ab93bb3cb4d365c9a8e45de)
- ingress_gw_ar.accelerated_networking.enable

<a id="canonical-df717926aef5a8b774a53dadff6dbe08af1f47cff8427e2daba8e1e04dd57ac1"></a>

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

<a id="canonical-90442a4280040bdda0477ad75c929b9ee74fe0da5588ac82444f30cc29a79927"></a>

## Direct properties — ingress_gw_ar.accelerated_networking.enable / 5bd4d06e9314 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7ae296eb925f33b201eb6a131d6917e39674a605211e3b3dad49cb82edc3e137"></a>

## Next pages — ingress_gw_ar.accelerated_networking.enable / 5bd4d06e9314 / 4

- [ingress_gw_ar.accelerated_networking](resources--azure_vnet_site--reference--group-007.md#canonical-a00ca808f947d1ddeda9960d2c7134e0e72ebeda1ab93bb3cb4d365c9a8e45de)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-87b70f7d188d23ff41722d94e8837588c63ec02ab72058e0fcb45f7d81e73a69"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-aab13699800a3e5e23e93b44cc9f9e1fbdfe99be4402ec4476df2468a9690a6d"></a>

## ingress_gw_ar.node — ingress_gw_ar.node / 5736e176b5c3 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- ingress_gw_ar.node

<a id="canonical-7746fec8c731f66d8359586aed692ab68f5a8e6d632af03935545933c59ac355"></a>

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

<a id="canonical-89cefb9b4d7be55004ea7a8e6ae443626846de3462e5252c64b57cae4e5e625e"></a>

## Direct properties — ingress_gw_ar.node / 5736e176b5c3 / 3

<a id="canonical-2918255c9aa013c4eb3fbe378692868c92aa2bdc7162197cae46779c88b482d7"></a>

<a id="canonical-99377ecf06d50b6daa4d9df9faee607378adb7b915a40cf4250fafde51ea6df9"></a>

## fault_domain property — ingress_gw_ar.node / 5736e176b5c3 / 4

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

- [local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c40a230559356b8bdb345774131fdd36f2125524a601a66842dee698de5a0873): complete subsection reference.

<a id="canonical-2bb9e8f20f699c649bc7fbb6082f98f95600f1d0eb7130d7ccc930a5e08ca4ae"></a>

<a id="canonical-fbd027c41a9e1e2e9eca074b840f9b72a1caa74419f14f647595b2c73f2919c9"></a>

## node_number property — ingress_gw_ar.node / 5736e176b5c3 / 5

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

<a id="canonical-7e5a9716024660ca52041cb6fde9b26c4e5f317f870ac57bf0357c7352c99071"></a>

<a id="canonical-6773f0a54844aacd441cd55ecac81d927c0d5958c75a3f602d8d86dee5ca170e"></a>

## update_domain property — ingress_gw_ar.node / 5736e176b5c3 / 6

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

<a id="canonical-1015e33594eb970a7391fa0d1c93816e3891831c346e9aa78ecaa8ade9e5aa32"></a>

## Next pages — ingress_gw_ar.node / 5736e176b5c3 / 7

- [ingress_gw_ar.node.local_subnet](resources--azure_vnet_site--reference--group-007.md#canonical-c40a230559356b8bdb345774131fdd36f2125524a601a66842dee698de5a0873)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)

<a id="canonical-c40a230559356b8bdb345774131fdd36f2125524a601a66842dee698de5a0873"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e30ab80c194e967946992da9d51124947a14fe623ff620b0e7bec0bd975a1eee"></a>

## ingress_gw_ar.node.local_subnet — ingress_gw_ar.node.local_subnet / 6f38c925dfdc / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-b087f208aea9bf78194523dff06d0d9350f1017d40cb35999cbaea1a437ac4d7)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-c663d8e15601cf9de2f980c11b1d3fc2561d340189ef651c56155a79bbbd7b6e)
- [ingress_gw_ar](resources--azure_vnet_site--reference--group-007.md#canonical-d0e215b74519793a40bfc7211040a95320d415ace2a874843eeb2ee665a518c4)
- [ingress_gw_ar.node](resources--azure_vnet_site--reference--group-007.md#canonical-87b70f7d188d23ff41722d94e8837588c63ec02ab72058e0fcb45f7d81e73a69)
- ingress_gw_ar.node.local_subnet

<a id="canonical-e50ace24d553e8f2c015f37882c65e35f9b81afecdc6ac0704e30fdb7f851c18"></a>

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

<a id="canonical-0f004153f30d56a54c065d4337f9aa5eb1aac9f351c8eb527aae280a4bc3f1a2"></a>

## Direct properties — ingress_gw_ar.node.local_subnet / 6f38c925dfdc / 3

- [subnet](resources--azure_vnet_site--reference--group-008.md#canonical-589e6304a3b881158b62acb9f1ffbda5e070f464ca55ad8236f02c18a58964a8): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-008.md#canonical-5278ceccd573239486fadf316b56502fac6d8825a9f4cea180295e9fc41d4c75): complete subsection reference.
