---
page_title: "xcsh_nfv_service reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service reference."
---

# xcsh_nfv_service reference

<a id="canonical-9494a9073c1949c329f151d435e4dfdb15368eba630267a4cb5afec3aeabf201"></a>

## f5_big_ip_aws_service.market_place_image — f5_big_ip_aws_service.market_place_image / 4d2855c181f3 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- f5_big_ip_aws_service.market_place_image

<a id="canonical-1cc63432f816a5d3af8b64ef59998f89108524fa5e92abf4305e4599f05903d1"></a>

Type: `"object"`. single nested block, Optional.

BIG-IP AWS Pay as You Go Image Selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("awafpay_g200_mbps",
    "awafpay_g3_gbps")}
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
  "x-ves-oneof-field-ami_choice": "[\"AWAFPayG200Mbps\",\"AWAFPayG3Gbps\",\"BestPlusPayG200Mbps\",\"best_plus_payg_1gbps\"]"
}
```

Terraform syntax:

```terraform
market_place_image {
  # Configure direct properties listed below.
}
```

<a id="canonical-20d2a7f4a643831facb46acf36cb4c0c75351b5c92e0a3f2b781b7bd1a4db151"></a>

## Direct properties — f5_big_ip_aws_service.market_place_image / 4d2855c181f3 / 3

- [awafpay_g200_mbps](resources--nfv_service--reference--group-002.md#canonical-f5b750c08d67bd503da37e8a6384480c2e57ef5a82ca6d84799a357d02cdc022): complete subsection reference.

- [awafpay_g3_gbps](resources--nfv_service--reference--group-002.md#canonical-f3a5a6d83cb15afa49b7ef0af78b198d21b37372597e6942361fc0e7965fcd45): complete subsection reference.

<a id="canonical-76069b46d16029e7fa7050b13afc4a61e2777257bd9a2f057f44779cc60c9b58"></a>

## Next pages — f5_big_ip_aws_service.market_place_image / 4d2855c181f3 / 4

- [f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps](resources--nfv_service--reference--group-002.md#canonical-f5b750c08d67bd503da37e8a6384480c2e57ef5a82ca6d84799a357d02cdc022)
- [f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps](resources--nfv_service--reference--group-002.md#canonical-f3a5a6d83cb15afa49b7ef0af78b198d21b37372597e6942361fc0e7965fcd45)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-f5b750c08d67bd503da37e8a6384480c2e57ef5a82ca6d84799a357d02cdc022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d5dcece6861088d728e8cde528f89ea17736c90819355e99c1454ebb2a60b57"></a>

## f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps — f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps / 78fa4cf5a73b / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.market_place_image](resources--nfv_service--reference--group-001.md#canonical-d9a50ca9b65a69d0b0323e4f0b7e41ed3ab5bac567fd056d418be98ac15f29e2)
- f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps

<a id="canonical-c7084c8d2b646cb723d4f92d063559267e61c74af78144ae56035468784bf987"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for AWAFPayG200Mbps.

Terraform syntax:

```terraform
awafpay_g200_mbps = {}
```

<a id="canonical-35ee7bcc29157a692586fd2bb32a063eac8c0beb94a9800229b2c2c10d1a9228"></a>

## Direct properties — f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps / 78fa4cf5a73b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d7fd7c67d06fc7daabbb8a5da643fd61f7ed9cc921d4ba81989953152ca5b67b"></a>

## Next pages — f5_big_ip_aws_service.market_place_image.awafpay_g200_mbps / 78fa4cf5a73b / 4

- [f5_big_ip_aws_service.market_place_image](resources--nfv_service--reference--group-001.md#canonical-d9a50ca9b65a69d0b0323e4f0b7e41ed3ab5bac567fd056d418be98ac15f29e2)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-f3a5a6d83cb15afa49b7ef0af78b198d21b37372597e6942361fc0e7965fcd45"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b3e6c596e4e2fd2e6d7c6ae8a54cf326b5827f1618a98036519e76e80fd30e3"></a>

## f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps — f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps / 6d5c3d0b6499 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.market_place_image](resources--nfv_service--reference--group-001.md#canonical-d9a50ca9b65a69d0b0323e4f0b7e41ed3ab5bac567fd056d418be98ac15f29e2)
- f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps

<a id="canonical-374ef557714400b204f2f477f3379f0a1118e2dd1396a8e5544f3201f5673045"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for AWAFPayG3Gbps.

Terraform syntax:

```terraform
awafpay_g3_gbps = {}
```

<a id="canonical-91f0d3c5c477a13b8c06fb28ef593a5138b3d61a18845b927e631678da6a6146"></a>

## Direct properties — f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps / 6d5c3d0b6499 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-fd2f602e888dc28e529be1a644f6e071b48f5c4d977ee9f526d12fd0de77e7ba"></a>

## Next pages — f5_big_ip_aws_service.market_place_image.awafpay_g3_gbps / 6d5c3d0b6499 / 4

- [f5_big_ip_aws_service.market_place_image](resources--nfv_service--reference--group-001.md#canonical-d9a50ca9b65a69d0b0323e4f0b7e41ed3ab5bac567fd056d418be98ac15f29e2)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-339df16e21e9e29d792fcf340e9c363ec54743f68cfd223b9539dd23fe1fa661"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-23d9235ae6f7a92748e5c0356ed7cc83668d2114750302f65425731010f589d5"></a>

## f5_big_ip_aws_service.nodes — f5_big_ip_aws_service.nodes / c95e8f4e2857 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- f5_big_ip_aws_service.nodes

<a id="canonical-fadbd30abbf68d6cfa975fcd4435fe90915961020e3673a5d9ede792c007691f"></a>

Type: `"object"`. list nested block, Optional.

Specify how and where the service nodes are spawned.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name",
    "node_name"),
  validators.ConflictingListObjectAttributes("automatic_prefix",
    "tunnel_prefix"),
  validators.ConflictingListObjectAttributes("mgmt_subnet",
    "reserved_mgmt_subnet")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
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
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-f5b78458022066cfa69adcc188b3c6a83e1608cac08d9db3f6f69c85f25d5670"></a>

## Direct properties — f5_big_ip_aws_service.nodes / c95e8f4e2857 / 3

- [automatic_prefix](resources--nfv_service--reference--group-002.md#canonical-410dae88aac8ee67c950c60db51185f019020833ec6b06bbc398b714a316a8ab): complete subsection reference.

<a id="canonical-9cd8d4056bf96870ae54d1191208c86ccaf68c42eca01d8a3f43123539be34ea"></a>

<a id="canonical-695b79be52ede5d35b9c0770acc51336d9a280937c09d3d30931299fcc7a205f"></a>

## aws_az_name property — f5_big_ip_aws_service.nodes / c95e8f4e2857 / 4

Type: `"string"`. Optional.

The AWS Availability Zone must be consistent with the AWS Region chosen. Please select an AZ in the
same Region as your TGW Site.

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
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([a-z]{2})-([a-z0-9]{4,20})-([a-z0-9]{2})$"
  }
}
```

- [mgmt_subnet](resources--nfv_service--reference--group-002.md#canonical-5fa40a80a761b7131c8f477f9f5a8ffd030c17382aef4e68791f15aaf4425e89): complete subsection reference.

<a id="canonical-cb308a3a26a0952be331e62c11e5219e47d8fab9c4e12c582fde65a67bbdb5fb"></a>

<a id="canonical-bc948a30ecaf806c585b113ed69b95af3b25a1d286b23abdc967fc3f6b21b636"></a>

## node_name property — f5_big_ip_aws_service.nodes / c95e8f4e2857 / 5

Type: `"string"`. Optional.

Node Name will be used to assign as hostname to the service.

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
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [reserved_mgmt_subnet](resources--nfv_service--reference--group-002.md#canonical-8452d2d2b84c3d797dc416031f8fbc09bd4b560975f4411369c1f69197f5755f): complete subsection reference.

<a id="canonical-4395c9ede84d44a752ecccf5a89c6c69ac57304dcc6172799c8fc2ff7f9b5c98"></a>

<a id="canonical-884673678cfe14feb8d886139ee6d7f4fc6d2e29de4108bd11dfdcd5b9b48364"></a>

## tunnel_prefix property — f5_big_ip_aws_service.nodes / c95e8f4e2857 / 6

Type: `"string"`. Optional.

Exclusive with \[automatic\_prefix\] Enter IP prefix for the tunnel, it has to be /30.

Upstream description:

Exclusive with \[automatic\_prefix\] Enter IP prefix for the tunnel, it has to be /30.

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

<a id="canonical-04f38d04300a50ed0e4437e703c94c3c31443f7eef9dd360387bfa4578c76a61"></a>

## Next pages — f5_big_ip_aws_service.nodes / c95e8f4e2857 / 7

- [f5_big_ip_aws_service.nodes.automatic_prefix](resources--nfv_service--reference--group-002.md#canonical-410dae88aac8ee67c950c60db51185f019020833ec6b06bbc398b714a316a8ab)
- [f5_big_ip_aws_service.nodes.mgmt_subnet](resources--nfv_service--reference--group-002.md#canonical-5fa40a80a761b7131c8f477f9f5a8ffd030c17382aef4e68791f15aaf4425e89)
- [f5_big_ip_aws_service.nodes.reserved_mgmt_subnet](resources--nfv_service--reference--group-002.md#canonical-8452d2d2b84c3d797dc416031f8fbc09bd4b560975f4411369c1f69197f5755f)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-410dae88aac8ee67c950c60db51185f019020833ec6b06bbc398b714a316a8ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-28fe5c82b7368c9626c2fab29a97281873f8a617970569ed12f594cd9ab82f27"></a>

## f5_big_ip_aws_service.nodes.automatic_prefix — f5_big_ip_aws_service.nodes.automatic_prefix / 74dc7ae84c1f / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.nodes](resources--nfv_service--reference--group-002.md#canonical-339df16e21e9e29d792fcf340e9c363ec54743f68cfd223b9539dd23fe1fa661)
- f5_big_ip_aws_service.nodes.automatic_prefix

<a id="canonical-39aa4c2677e729f4e674b22177da306e201450bc4bd4b2c1082cb2b2cc37ef7b"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic prefix.

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
automatic_prefix = {}
```

<a id="canonical-c6b19d35c132a5eb2846e002025ed189259514d8ad7d4c74d2163a9496cdfa7c"></a>

## Direct properties — f5_big_ip_aws_service.nodes.automatic_prefix / 74dc7ae84c1f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a99fb757755264ecb0c482339c6fa146ebcdad12c5e4725263925e3e83f314b9"></a>

## Next pages — f5_big_ip_aws_service.nodes.automatic_prefix / 74dc7ae84c1f / 4

- [f5_big_ip_aws_service.nodes](resources--nfv_service--reference--group-002.md#canonical-339df16e21e9e29d792fcf340e9c363ec54743f68cfd223b9539dd23fe1fa661)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-5fa40a80a761b7131c8f477f9f5a8ffd030c17382aef4e68791f15aaf4425e89"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7834ac5693860f09c7202a92116bbb937e949bf32895baee815fe2041a7c7632"></a>

## f5_big_ip_aws_service.nodes.mgmt_subnet — f5_big_ip_aws_service.nodes.mgmt_subnet / 88b7e43d1b63 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.nodes](resources--nfv_service--reference--group-002.md#canonical-339df16e21e9e29d792fcf340e9c363ec54743f68cfd223b9539dd23fe1fa661)
- f5_big_ip_aws_service.nodes.mgmt_subnet

<a id="canonical-1a08c5015a18a055dda2b28cedf19408fbeb3a12c0f860ec490c26a6a99c5b93"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for mgmt subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
mgmt_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-ad76ad8e6c260d566ea38b1e342ece2d84995e7a7c15af52452ce940c7a2f3f3"></a>

## Direct properties — f5_big_ip_aws_service.nodes.mgmt_subnet / 88b7e43d1b63 / 3

<a id="canonical-e545c95d05f4deb9522d42fdb8079cc47ac407d2eaa14207a230e757b8d4916c"></a>

<a id="canonical-c44e21a9960a742bbd0d93c862728a86d66bd5619930be571ce218520791f802"></a>

## existing_subnet_id property — f5_big_ip_aws_service.nodes.mgmt_subnet / 88b7e43d1b63 / 4

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](resources--nfv_service--reference--group-002.md#canonical-b077f913e15ebb40b55a3d5c4560f8db6f6df697d1092f1002bdc68236ba93c5): complete subsection reference.

<a id="canonical-cb12fe4a2f00ffc2f499791f3e0bf9ee278f615b6abab9e88016a2eaf54938d3"></a>

## Next pages — f5_big_ip_aws_service.nodes.mgmt_subnet / 88b7e43d1b63 / 5

- [f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param](resources--nfv_service--reference--group-002.md#canonical-b077f913e15ebb40b55a3d5c4560f8db6f6df697d1092f1002bdc68236ba93c5)
- [f5_big_ip_aws_service.nodes](resources--nfv_service--reference--group-002.md#canonical-339df16e21e9e29d792fcf340e9c363ec54743f68cfd223b9539dd23fe1fa661)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-b077f913e15ebb40b55a3d5c4560f8db6f6df697d1092f1002bdc68236ba93c5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6eaa68e3cc9f7898ec70e5224380925776b296aacbe4b6670521386dec3217b2"></a>

## f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param — f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param / 8de33afc233c / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.nodes](resources--nfv_service--reference--group-002.md#canonical-339df16e21e9e29d792fcf340e9c363ec54743f68cfd223b9539dd23fe1fa661)
- [f5_big_ip_aws_service.nodes.mgmt_subnet](resources--nfv_service--reference--group-002.md#canonical-5fa40a80a761b7131c8f477f9f5a8ffd030c17382aef4e68791f15aaf4425e89)
- f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param

<a id="canonical-17d97287c39533e97427b2f5a0d0b0eb8642638796ad32a82e3630aa63f4a311"></a>

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

<a id="canonical-0b26d4a32f94d91baa88a698107664e26048922ddb1de71efaaf0edb0a15239f"></a>

## Direct properties — f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param / 8de33afc233c / 3

<a id="canonical-3d92be3a4f0c0a1254ff29326cba629f7285a67616f3fd31bb63a09d53a97592"></a>

<a id="canonical-5b4875a18ba6c143e25db6217a72edeccdb582a9582763b3ac5be494b8afb7c5"></a>

## ipv4 property — f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param / 8de33afc233c / 4

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

<a id="canonical-8dd1638a85a77c175bc6132d1648154f0f10e72089057b93901f2e4457a87b2d"></a>

## Next pages — f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param / 8de33afc233c / 5

- [f5_big_ip_aws_service.nodes.mgmt_subnet](resources--nfv_service--reference--group-002.md#canonical-5fa40a80a761b7131c8f477f9f5a8ffd030c17382aef4e68791f15aaf4425e89)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-8452d2d2b84c3d797dc416031f8fbc09bd4b560975f4411369c1f69197f5755f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4adfde101595d5f54908759e618a787a277f5c362064c69725f847b567bb9679"></a>

## f5_big_ip_aws_service.nodes.reserved_mgmt_subnet — f5_big_ip_aws_service.nodes.reserved_mgmt_subnet / 1c69b4e81f17 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [f5_big_ip_aws_service](resources--nfv_service--reference--group-001.md#canonical-a398a983a6c5e9e184b4c525d893a490dcb13f6c3e442c43b2e0d505d9641ca4)
- [f5_big_ip_aws_service.nodes](resources--nfv_service--reference--group-002.md#canonical-339df16e21e9e29d792fcf340e9c363ec54743f68cfd223b9539dd23fe1fa661)
- f5_big_ip_aws_service.nodes.reserved_mgmt_subnet

<a id="canonical-35ede095a92e65d39dd73627d9cbd5cbff403eac877d5e5f1e742f7f44d8d4f9"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved mgmt subnet.

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
reserved_mgmt_subnet = {}
```

<a id="canonical-1dbd2183300cbd8355b13687ce40e80082b19f66839fdd8dd8f912c1cea2e93e"></a>

## Direct properties — f5_big_ip_aws_service.nodes.reserved_mgmt_subnet / 1c69b4e81f17 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0b3a2b5e64b4a12881363f94a3f785a1257a269dfcbd9fb57ce54c18ba7d117b"></a>

## Next pages — f5_big_ip_aws_service.nodes.reserved_mgmt_subnet / 1c69b4e81f17 / 4

- [f5_big_ip_aws_service.nodes](resources--nfv_service--reference--group-002.md#canonical-339df16e21e9e29d792fcf340e9c363ec54743f68cfd223b9539dd23fe1fa661)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-09f1372720f1e78ab3d6fc0c7331ea8ba85521f7679847283663b3d69222c208"></a>

## https_management — https_management / 3a3a01043efc / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- https_management

<a id="canonical-b4518939056897104846d93f414610b8e7ef75ac36012d24571976aa6b8b425c"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for https management.

Upstream description:

HTTPS based configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domain_suffix"),
  validators.ConflictingObjectAttributes("advertise_on_internet",
    "advertise_on_internet_default_vip"),
  validators.ConflictingObjectAttributes("advertise_on_internet",
    "advertise_on_sli_vip"),
  validators.ConflictingObjectAttributes("advertise_on_internet",
    "advertise_on_slo_internet_vip"),
  validators.ConflictingObjectAttributes("advertise_on_internet",
    "advertise_on_slo_sli"),
  validators.ConflictingObjectAttributes("advertise_on_internet",
    "advertise_on_slo_vip"),
  validators.ConflictingObjectAttributes("advertise_on_internet_default_vip",
    "advertise_on_sli_vip"),
  validators.ConflictingObjectAttributes("advertise_on_internet_default_vip",
    "advertise_on_slo_internet_vip"),
  validators.ConflictingObjectAttributes("advertise_on_internet_default_vip",
    "advertise_on_slo_sli"),
  validators.ConflictingObjectAttributes("advertise_on_internet_default_vip",
    "advertise_on_slo_vip"),
  validators.ConflictingObjectAttributes("advertise_on_sli_vip",
    "advertise_on_slo_internet_vip"),
  validators.ConflictingObjectAttributes("advertise_on_sli_vip",
    "advertise_on_slo_sli"),
  validators.ConflictingObjectAttributes("advertise_on_sli_vip",
    "advertise_on_slo_vip"),
  validators.ConflictingObjectAttributes("advertise_on_slo_internet_vip",
    "advertise_on_slo_sli"),
  validators.ConflictingObjectAttributes("advertise_on_slo_internet_vip",
    "advertise_on_slo_vip"),
  validators.ConflictingObjectAttributes("advertise_on_slo_sli",
    "advertise_on_slo_vip"),
  validators.ConflictingObjectAttributes("default_https_port",
    "https_port")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_on_internet\",\"advertise_on_internet_default_vip\",\"advertise_on_sli_vip\",\"advertise_on_slo_internet_vip\",\"advertise_on_slo_sli\",\"advertise_on_slo_vip\"]",
  "x-ves-oneof-field-internet_choice": "[]",
  "x-ves-oneof-field-port_choice": "[\"default_https_port\",\"https_port\"]"
}
```

Terraform syntax:

```terraform
https_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-411f39d094d0ddf2db34330b6bbe0184f9f5f0d5f8fb09875b9b7f21bbe5dc46"></a>

## Direct properties — https_management / 3a3a01043efc / 3

- [advertise_on_internet](resources--nfv_service--reference--group-002.md#canonical-18816543fde40e39c1c48526afda700cd90632a6b683f2a598a2009cfd76594b): complete subsection reference.

- [advertise_on_internet_default_vip](resources--nfv_service--reference--group-002.md#canonical-952e94d129a50e5a5f7bc820895afafedddeb175e11cd6c69c1cf597351c441a): complete subsection reference.

- [advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2): complete subsection reference.

- [advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c): complete subsection reference.

- [advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46): complete subsection reference.

- [advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580): complete subsection reference.

- [default_https_port](resources--nfv_service--reference--group-003.md#canonical-4d3e72934bfc108341684f9a988111dde4fcdd5edc9a8f34d8eb22a612cddd3e): complete subsection reference.

<a id="canonical-dde6155ccfaa0e6d618b72ed7f5b47333a61d150d8b6aaa0c69c16ac8bf1fa4b"></a>

<a id="canonical-98c015e55ee00060f43b1642010d08d39008b329b5469558cacfb631f176dedc"></a>

## domain_suffix property — https_management / 3a3a01043efc / 4

Type: `"string"`. Optional.

Domain suffix will be used along with node name to form URL to access node management.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-1b2cebb33b94189f870fa07a4eda9b6e96ab7bc92c8b8b5fdf46aa0080f80501"></a>

<a id="canonical-02b75beba7e16cb77949c8c1ffd8362ca99659da19fb813844567c52519b3c83"></a>

## https_port property — https_management / 3a3a01043efc / 5

Type: `"number"`. Optional.

Exclusive with \[default\_https\_port\] Enter TCP port number.

Upstream description:

Exclusive with \[default\_https\_port\] Enter TCP port number.

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

<a id="canonical-08d8b31e270dd89415f906174341c8ccb06684e69f472315b5601b89145295a6"></a>

## Next pages — https_management / 3a3a01043efc / 6

- [https_management.advertise_on_internet](resources--nfv_service--reference--group-002.md#canonical-18816543fde40e39c1c48526afda700cd90632a6b683f2a598a2009cfd76594b)
- [https_management.advertise_on_internet_default_vip](resources--nfv_service--reference--group-002.md#canonical-952e94d129a50e5a5f7bc820895afafedddeb175e11cd6c69c1cf597351c441a)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_sli](resources--nfv_service--reference--group-003.md#canonical-30022b387cd576f0c782dfcba45fea9f6211f83d534f947e7512e992c2608a46)
- [https_management.advertise_on_slo_vip](resources--nfv_service--reference--group-003.md#canonical-dbbbdded43c0c30607bcb8e329159cdf82634f24371abc298c763310ecccf580)
- [https_management.default_https_port](resources--nfv_service--reference--group-003.md#canonical-4d3e72934bfc108341684f9a988111dde4fcdd5edc9a8f34d8eb22a612cddd3e)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-18816543fde40e39c1c48526afda700cd90632a6b683f2a598a2009cfd76594b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bcc35035cd2e4e405e0b75ad617730ebb0d0dc81906482a2f0c7b9b55c5379ad"></a>

## https_management.advertise_on_internet — https_management.advertise_on_internet / 54137af60568 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- https_management.advertise_on_internet

<a id="canonical-3b10091445e68467f2dc70d0d1e36723c851e08499c661d3d62687bc350f0b60"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
advertise_on_internet {
  # Configure direct properties listed below.
}
```

<a id="canonical-1dcc60356b2961a22948955d7bca57feaca48a6bac48a26a66e99b746a07fbf0"></a>

## Direct properties — https_management.advertise_on_internet / 54137af60568 / 3

- [public_ip](resources--nfv_service--reference--group-002.md#canonical-10797d51edf30028eb9475505f1ac3ffdd78a3760b1a3e6a44dfcaf9e8b8805a): complete subsection reference.

<a id="canonical-ac57c83345aa736020e4135a2f884c39c9b6514e4e8ea1d52c233b74dd6a3044"></a>

## Next pages — https_management.advertise_on_internet / 54137af60568 / 4

- [https_management.advertise_on_internet.public_ip](resources--nfv_service--reference--group-002.md#canonical-10797d51edf30028eb9475505f1ac3ffdd78a3760b1a3e6a44dfcaf9e8b8805a)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-10797d51edf30028eb9475505f1ac3ffdd78a3760b1a3e6a44dfcaf9e8b8805a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-49d48ebf5e49120b6149c950594a5ac09d6b9485b22c80425e766ff35b8166d2"></a>

## https_management.advertise_on_internet.public_ip — https_management.advertise_on_internet.public_ip / 175ca036db64 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_internet](resources--nfv_service--reference--group-002.md#canonical-18816543fde40e39c1c48526afda700cd90632a6b683f2a598a2009cfd76594b)
- https_management.advertise_on_internet.public_ip

<a id="canonical-782b2e3131ed654806a4e98f59de9d0dd592a2d2656a083326ebfc69a0d16308"></a>

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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-a83775bb6c39564edecc2fc7b12bef75fd3692f82812ecb3b27b9d45baa44266"></a>

## Direct properties — https_management.advertise_on_internet.public_ip / 175ca036db64 / 3

<a id="canonical-9295477d0db70f40eb79fca519d6f150bf2d25f30b8899112393a660c834547b"></a>

<a id="canonical-9b69c071b67666398e09e2c522c421d6081933642cd0bd0f8e15332e96e2a8e7"></a>

## name property — https_management.advertise_on_internet.public_ip / 175ca036db64 / 4

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

<a id="canonical-d31aa133f06800231f7f56705f79df54f4d0298bbc715f06a909fdeb81168a02"></a>

<a id="canonical-6ad59757d56fd5cd7f9d43e22b64994894553ad2e18598c3597110c4bb4af290"></a>

## namespace property — https_management.advertise_on_internet.public_ip / 175ca036db64 / 5

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

<a id="canonical-c568dfdb0491536c6043d483fb49105362515c562828666e4d4d2a78feb0af29"></a>

<a id="canonical-8579b2d3fa44498bdaf32bb08d30a00c422082a693681119951700bf60dd11c7"></a>

## tenant property — https_management.advertise_on_internet.public_ip / 175ca036db64 / 6

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

<a id="canonical-cb9330107cfeec356692aaca31d5d815286cfe6750135f648538b93def19cfcb"></a>

## Next pages — https_management.advertise_on_internet.public_ip / 175ca036db64 / 7

- [https_management.advertise_on_internet](resources--nfv_service--reference--group-002.md#canonical-18816543fde40e39c1c48526afda700cd90632a6b683f2a598a2009cfd76594b)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-952e94d129a50e5a5f7bc820895afafedddeb175e11cd6c69c1cf597351c441a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f43714e94ac33a15433cd596391fec30bd6d2e1ceb1e2869701ee298dc66f96"></a>

## https_management.advertise_on_internet_default_vip — https_management.advertise_on_internet_default_vip / 20feab55c982 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- https_management.advertise_on_internet_default_vip

<a id="canonical-ceca158e74c183b3d2d2204bde0a31fe379637faac00de2a0c03cd4ffc696ae5"></a>

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
advertise_on_internet_default_vip = {}
```

<a id="canonical-51c7f1ba5fcf27f38a7976ee60a664d83bac816b74cc72476d8456751843e4d9"></a>

## Direct properties — https_management.advertise_on_internet_default_vip / 20feab55c982 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c3a56cd6deeb9cbe4818f492356c6e3cb3bfad0b8e35ec7f617a77914f653069"></a>

## Next pages — https_management.advertise_on_internet_default_vip / 20feab55c982 / 4

- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-601d1bfafc97caa1579776822657d23df866af6aed696348f1be74395054bfb8"></a>

## https_management.advertise_on_sli_vip — https_management.advertise_on_sli_vip / 02a15c406e5c / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- https_management.advertise_on_sli_vip

<a id="canonical-3f882978171e511f48b1ba30574f7529eec5107bc5af8911be3b8dcdb064387f"></a>

Type: `"object"`. single nested block, Optional.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
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
advertise_on_sli_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-26fb8686052a2c2e35650f2e7fd07c4f0ab56175d6d926dd1b2da2ef696c3bee"></a>

## Direct properties — https_management.advertise_on_sli_vip / 02a15c406e5c / 3

- [no_mtls](resources--nfv_service--reference--group-002.md#canonical-4853d8aae4a60772025557044b9af5dbfaf244af1e069861ec2377f509e478bb): complete subsection reference.

- [tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701): complete subsection reference.

- [tls_config](resources--nfv_service--reference--group-002.md#canonical-d1a648731a53dfe450f2260c554666be91a89d63d41248e33bc92e52d4158406): complete subsection reference.

- [use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe): complete subsection reference.

<a id="canonical-133d57d062431ef7c0de39d7049e2b6b3d1313eb42e381c9b420ae4340bb73fd"></a>

## Next pages — https_management.advertise_on_sli_vip / 02a15c406e5c / 4

- [https_management.advertise_on_sli_vip.no_mtls](resources--nfv_service--reference--group-002.md#canonical-4853d8aae4a60772025557044b9af5dbfaf244af1e069861ec2377f509e478bb)
- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701)
- [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-d1a648731a53dfe450f2260c554666be91a89d63d41248e33bc92e52d4158406)
- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-4853d8aae4a60772025557044b9af5dbfaf244af1e069861ec2377f509e478bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-44710675aa0c7bed547046e389ee644ca01dec882bb5dd16c2f79cc78f6cb003"></a>

## https_management.advertise_on_sli_vip.no_mtls — https_management.advertise_on_sli_vip.no_mtls / 2d4a6a832115 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- https_management.advertise_on_sli_vip.no_mtls

<a id="canonical-6ce6f81f6b020ffd3327a917e9f83d3bc2dc5061ae8466930bb3afa05fc9f415"></a>

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

<a id="canonical-92aff8f2a3d3b6eea18ebd29307a1b5d9111b251edec4604f20efeac729195fa"></a>

## Direct properties — https_management.advertise_on_sli_vip.no_mtls / 2d4a6a832115 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a75b136aae87577a5182f46d198f5ac2eaf86fd8301583047d6a298ec56646ac"></a>

## Next pages — https_management.advertise_on_sli_vip.no_mtls / 2d4a6a832115 / 4

- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b44078a08230a8253f89741fdd1fecb3cbc312ca1df24a3a61021476e8309ea"></a>

## https_management.advertise_on_sli_vip.tls_certificates — https_management.advertise_on_sli_vip.tls_certificates / 9166a5b4a497 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- https_management.advertise_on_sli_vip.tls_certificates

<a id="canonical-3d9f35642287f71fe43f45f66446a74e75d105530fcc7d6bea80594f54ae23ff"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-cb53eb7e844eba9b99abc9368385eb43764f8d63576e323ad0376a8e19d90fe4"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates / 9166a5b4a497 / 3

<a id="canonical-b150106cdf741889867f6845e2dc440db97e1fcd4ec3a4d8ade2f748533ecd72"></a>

<a id="canonical-b53dc227292a058868168180a836659bc8b66249f7708a45e55b72329105135f"></a>

## certificate_url property — https_management.advertise_on_sli_vip.tls_certificates / 9166a5b4a497 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--nfv_service--reference--group-002.md#canonical-a92913c7c7ce3157c1a2dd9fd926bd3174237f141818ada8643b12e5ffcf4256): complete subsection reference.

<a id="canonical-7336ccbb4a1779c5a2a775f85d1d8316590b89e9239a776740b5702914e861fc"></a>

<a id="canonical-ac4c175cf307add4c95434e6b130f58454a4c48a4532fdbc24e22e943ef0047f"></a>

## description_spec property — https_management.advertise_on_sli_vip.tls_certificates / 9166a5b4a497 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--nfv_service--reference--group-002.md#canonical-051ec0d4c7962ac509c60a8f6ba8873f43733cf3d6d11df7c317b45c57f6472c): complete subsection reference.

- [private_key](resources--nfv_service--reference--group-002.md#canonical-fe8daac3a4c58b0738c7622a9ef11ddcff4f12fb46750313c6c9f04850a43692): complete subsection reference.

- [use_system_defaults](resources--nfv_service--reference--group-002.md#canonical-47d5d764b6ce5822ad677e84e9042bb2862deebfc1801840b871cbcac0952f3e): complete subsection reference.

<a id="canonical-9d5acc08478e79f49ab1ec0b3ff073e4451e01eeb7145424889ead048b919bfd"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates / 9166a5b4a497 / 6

- [https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-002.md#canonical-a92913c7c7ce3157c1a2dd9fd926bd3174237f141818ada8643b12e5ffcf4256)
- [https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-002.md#canonical-051ec0d4c7962ac509c60a8f6ba8873f43733cf3d6d11df7c317b45c57f6472c)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-fe8daac3a4c58b0738c7622a9ef11ddcff4f12fb46750313c6c9f04850a43692)
- [https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-002.md#canonical-47d5d764b6ce5822ad677e84e9042bb2862deebfc1801840b871cbcac0952f3e)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-a92913c7c7ce3157c1a2dd9fd926bd3174237f141818ada8643b12e5ffcf4256"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef3f00fcdd8018b146c7e6faab2919f3cd8f12539063611653bff412f98152ce"></a>

## https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms — https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms / bc56a4e5029c / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701)
- https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-2ea395e648a2473491d7a16060a837c93bf975d8915fd2bf6ffc6d865d77403d"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-a528be131e35a9a6e0de5bd22e9d858d9091db98c7ac53fbcf3e8a1a28f1ebec"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms / bc56a4e5029c / 3

<a id="canonical-12357436615cb4dd496eb0494a567f399e2f238cd4d50784da0450b8111331e0"></a>

<a id="canonical-cc298c1dd49b47ac247a74c91d1b3aaa94332603425fa9c4146c32bcdf668910"></a>

## hash_algorithms property — https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms / bc56a4e5029c / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9f283f5d6e0b01d7af2ba465729fd0d2fb1ba8fe343d56d69f565c827ad3e08e"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.custom_hash_algorithms / bc56a4e5029c / 5

- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-051ec0d4c7962ac509c60a8f6ba8873f43733cf3d6d11df7c317b45c57f6472c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-342a75ab69cc76887a3c959f1cc118679a7be686dc698836601e27c75624ea54"></a>

## https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling — https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling / b3d7871f439d / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701)
- https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-6b30ffcba6c173e586c221fe002c997a40af5be1a2d42c58f653b62317a3b193"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-3aa04e9292a96fd45512d98164012ce3829e7140105fef217a650bec1ecad9c9"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling / b3d7871f439d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5df1b3ac583669aab936080a6171237b0c923e3a2a73c5118b7893263713f4d5"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.disable_ocsp_stapling / b3d7871f439d / 4

- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-fe8daac3a4c58b0738c7622a9ef11ddcff4f12fb46750313c6c9f04850a43692"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a429bd3df6d0f00165efc318c3973eca477a213d0661eb593c7f17780261e808"></a>

## https_management.advertise_on_sli_vip.tls_certificates.private_key — https_management.advertise_on_sli_vip.tls_certificates.private_key / 32dc0bea28bc / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701)
- https_management.advertise_on_sli_vip.tls_certificates.private_key

<a id="canonical-f2e2876452be96b1fb2573cbb77e4831ac9ca1a66538b25a22b207925f3171a8"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1b497d4281685d31bac5273056781f161b1512a92bec342cc319baa70d4f8106"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.private_key / 32dc0bea28bc / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-002.md#canonical-7821fb3ee862e18865371a0a2e98b718170fb25bdb0ac1f3a695391e66382424): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-002.md#canonical-6bd488d10955e68fd27be5da205fe27c00b1478b56af8418cb1751f85082e4ce): complete subsection reference.

<a id="canonical-a704813f2a7f7ce704dca5b2ef1461baf9a8d4d98284d59e26d6fb798699517a"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.private_key / 32dc0bea28bc / 4

- [https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-002.md#canonical-7821fb3ee862e18865371a0a2e98b718170fb25bdb0ac1f3a695391e66382424)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-002.md#canonical-6bd488d10955e68fd27be5da205fe27c00b1478b56af8418cb1751f85082e4ce)
- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-7821fb3ee862e18865371a0a2e98b718170fb25bdb0ac1f3a695391e66382424"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5e877985d220014db39d7d8c7a755981e0e240c5e95dc038586e95377d788e9"></a>

## https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / abeab4a0853a / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-fe8daac3a4c58b0738c7622a9ef11ddcff4f12fb46750313c6c9f04850a43692)
- https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-7e8c2d36965687a7d7eb3a256515d13bbdeae78648362b537ff692ea6f9b1066"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3f33e67a206424993d9bc1e78abb410aa2680cc7d963b89d9a7c872099d8f280"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / abeab4a0853a / 3

<a id="canonical-f390615e467cc8fce1bac5b84f1d17b8b67067d9731a1bb4be9fc1a3b6076f16"></a>

<a id="canonical-adb1d4da6278e113de632cfbde2f2309a58396442ea021903c5106ea124d6ca8"></a>

## decryption_provider property — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / abeab4a0853a / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-c12d63df38050a701274610d00170d53360a15eb37c57c10778ae9c2d38a3afa"></a>

<a id="canonical-ad8c99eece3a39c98665d3539f4b73922a16b31c5b2bcc7c660756af95fa0c8f"></a>

## location property — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / abeab4a0853a / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-9f953212973c1959001b3ef5a900a74dda57527c47588e5f3a82bd120738f754"></a>

<a id="canonical-bd1b4195a7f1f941fb0f0c95cfb91cee76dfad889ce39eab47b5c46d5da9099a"></a>

## store_provider property — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / abeab4a0853a / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-fa4a646b06342ba56d2624c685f37f4d6c923d31987a4b44ca288dda4c9a9062"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.private_key.blindfold_sec / abeab4a0853a / 7

- [https_management.advertise_on_sli_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-fe8daac3a4c58b0738c7622a9ef11ddcff4f12fb46750313c6c9f04850a43692)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-6bd488d10955e68fd27be5da205fe27c00b1478b56af8418cb1751f85082e4ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b881d19670d6398a555dc2397c37290b11acad1d3db29dd3fa97c54c551619d8"></a>

## https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info — https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_ / 8484eb697e96 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701)
- [https_management.advertise_on_sli_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-fe8daac3a4c58b0738c7622a9ef11ddcff4f12fb46750313c6c9f04850a43692)
- https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_info

<a id="canonical-5d039e5829dc1747f7c1eafcaaba4ee03e8974a3a0b68bfe6fef416021f82cfb"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-f3a3eb332f8fa16d1b7513f5c31f2e36f449c0aa682e7068923a0098c9f079d7"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_ / 8484eb697e96 / 3

<a id="canonical-5b657f86dc67feb541d822ea48622d8ab3243dfccc4c041e9fe1db4c746b85a6"></a>

<a id="canonical-82af6f1f13815eec58aa7e27bc04c692709bd49f30a72617b3359740087b2437"></a>

## provider_ref property — https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_ / 8484eb697e96 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-f2e185e9d3da5701b2f0ac0d5ad531697a410901aec02590614a54cf2c43c94a"></a>

<a id="canonical-9f4041659ecfa225c9bb1d1e40eb65b014c6455a64b4a128f125c1951109041d"></a>

## url property — https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_ / 8484eb697e96 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-74a713a9b24a0bee4042393571c167ee0b226308747c4f478fb3ecf6c0132591"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.private_key.clear_secret_ / 8484eb697e96 / 6

- [https_management.advertise_on_sli_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-fe8daac3a4c58b0738c7622a9ef11ddcff4f12fb46750313c6c9f04850a43692)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-47d5d764b6ce5822ad677e84e9042bb2862deebfc1801840b871cbcac0952f3e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ec092928310e4ea82fc624ac04f9d30c67b8358748f6b8cad2f1ffd3dfe4f62d"></a>

## https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults — https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults / 300e68c0bf4f / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701)
- https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults

<a id="canonical-21d71019f6509ab7713e77c8c4c61f9e0fa7b103675d0667c1fccf0355a4d322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-cf294027b69b020475fb5ac18191876311cdc42c161ed24b8cb1eec88208b144"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults / 300e68c0bf4f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-59ad8f1aa92c0c881b838571ae6233ad8c8092195bcc129822e97d7e668c5da3"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_certificates.use_system_defaults / 300e68c0bf4f / 4

- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-a270e5bea74cc234884fe3098034df1f0e82a6e9e05b4b60fc5179886fd03701)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-d1a648731a53dfe450f2260c554666be91a89d63d41248e33bc92e52d4158406"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4e77d7f5328a924ea46634596a528e5400e34b1c44990e0a735bd5a41f46463"></a>

## https_management.advertise_on_sli_vip.tls_config — https_management.advertise_on_sli_vip.tls_config / 0a1750bf8ce0 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- https_management.advertise_on_sli_vip.tls_config

<a id="canonical-d1007795e2ed398441bfd1f3ec0ba65ff1d216e59200c07cd0077b6003edcebe"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-4ac448fef9833c193274753c6f4744bcdc99be29c6a9a45ffa69c80194d13a1e"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_config / 0a1750bf8ce0 / 3

- [custom_security](resources--nfv_service--reference--group-002.md#canonical-2d15bebc531a71d6a8441630842a354fe9d824b8ca72a43b1ae897254e2a6820): complete subsection reference.

- [default_security](resources--nfv_service--reference--group-002.md#canonical-c339c6e12188bedc465ce7295e9116d493e397234e2398814aba212df9a9a3f1): complete subsection reference.

- [low_security](resources--nfv_service--reference--group-002.md#canonical-3e6387de757e028acd864e4ba0718c5c44e7d9d9944e95a84d716fc29bfed76d): complete subsection reference.

- [medium_security](resources--nfv_service--reference--group-002.md#canonical-50a8ccd45d40e7b917a226bd46573c174096e6dcb79fd0b9cb0376caa9100243): complete subsection reference.

<a id="canonical-b95f3a3acfc53d2f2c21d2a698ce48ef4ff117f7cfb880a74803a9184cf4d837"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_config / 0a1750bf8ce0 / 4

- [https_management.advertise_on_sli_vip.tls_config.custom_security](resources--nfv_service--reference--group-002.md#canonical-2d15bebc531a71d6a8441630842a354fe9d824b8ca72a43b1ae897254e2a6820)
- [https_management.advertise_on_sli_vip.tls_config.default_security](resources--nfv_service--reference--group-002.md#canonical-c339c6e12188bedc465ce7295e9116d493e397234e2398814aba212df9a9a3f1)
- [https_management.advertise_on_sli_vip.tls_config.low_security](resources--nfv_service--reference--group-002.md#canonical-3e6387de757e028acd864e4ba0718c5c44e7d9d9944e95a84d716fc29bfed76d)
- [https_management.advertise_on_sli_vip.tls_config.medium_security](resources--nfv_service--reference--group-002.md#canonical-50a8ccd45d40e7b917a226bd46573c174096e6dcb79fd0b9cb0376caa9100243)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-2d15bebc531a71d6a8441630842a354fe9d824b8ca72a43b1ae897254e2a6820"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a1348d5ba409ac0b29f4900aed6c7b0a71aa4757a37788369631a99ded430dd"></a>

## https_management.advertise_on_sli_vip.tls_config.custom_security — https_management.advertise_on_sli_vip.tls_config.custom_security / dfcdfad2a063 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-d1a648731a53dfe450f2260c554666be91a89d63d41248e33bc92e52d4158406)
- https_management.advertise_on_sli_vip.tls_config.custom_security

<a id="canonical-396d63c6998b4786856bf7f9a498f6218258897995d9544f58edc2962f8318a7"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-f18ee76e84ec9bebf9793fe504d59aef47cf3416bacd406d528ad83cc3fb3312"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_config.custom_security / dfcdfad2a063 / 3

<a id="canonical-cad1fdd4574b8f4da20fd4aafac5027bf0c39a6a43d1a3a3ca7c5a8bc9094544"></a>

<a id="canonical-6facc6c78d25efa6938b93dd0ffbc7fc348ff643444d51190186c5de15f41914"></a>

## cipher_suites property — https_management.advertise_on_sli_vip.tls_config.custom_security / dfcdfad2a063 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-4db9a5d704a682aaaa7d9c23b8fa2a433c5e06451d95d1806b04e4f675556474"></a>

<a id="canonical-786471cef7395b738c69fda81d4fd64bfc91c4d8ce11195b1f6c5d0de81d4d31"></a>

## max_version property — https_management.advertise_on_sli_vip.tls_config.custom_security / dfcdfad2a063 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fd011b44902f2d9b6ea9695728dc8c77ed56945fba86e9723eaed89efffad02e"></a>

<a id="canonical-c620a277cd36a52b5b8a0177e7515f87f7958c90b12301ae5ec6434ffbfb06aa"></a>

## min_version property — https_management.advertise_on_sli_vip.tls_config.custom_security / dfcdfad2a063 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f5d4c146a09a6aba07176c654c58dad11e3e9b425fbc7a5e0b647238de936d80"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_config.custom_security / dfcdfad2a063 / 7

- [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-d1a648731a53dfe450f2260c554666be91a89d63d41248e33bc92e52d4158406)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-c339c6e12188bedc465ce7295e9116d493e397234e2398814aba212df9a9a3f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16eedefeafbb83cdfe95afb9c42afad48fdcadeb5fc1a8c81c3e8aec36cbd42f"></a>

## https_management.advertise_on_sli_vip.tls_config.default_security — https_management.advertise_on_sli_vip.tls_config.default_security / e472d9d09f70 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-d1a648731a53dfe450f2260c554666be91a89d63d41248e33bc92e52d4158406)
- https_management.advertise_on_sli_vip.tls_config.default_security

<a id="canonical-c2498e069c5e5f1b5e6a1e0a9936bd3c76c7a133287c1fa94ec10ca2de6c1892"></a>

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
default_security = {}
```

<a id="canonical-b6c507bd8c01079b3b8f52d4330e174c05cd0058287303c571df0a12e70e1ed5"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_config.default_security / e472d9d09f70 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-da4e7240d1ed012d6551f0669313a062314c6aae202ecf03d812a779583bc691"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_config.default_security / e472d9d09f70 / 4

- [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-d1a648731a53dfe450f2260c554666be91a89d63d41248e33bc92e52d4158406)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-3e6387de757e028acd864e4ba0718c5c44e7d9d9944e95a84d716fc29bfed76d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5d91d9b10e0315a07b0a6f1245159b226bb90c5d74ff4fd7fd9d7d345ab4acf9"></a>

## https_management.advertise_on_sli_vip.tls_config.low_security — https_management.advertise_on_sli_vip.tls_config.low_security / ac85e870c827 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-d1a648731a53dfe450f2260c554666be91a89d63d41248e33bc92e52d4158406)
- https_management.advertise_on_sli_vip.tls_config.low_security

<a id="canonical-b0713211203b5df2b384800d6ea3b65b4cbd69cb6393b76e24416f0a8eec9c3c"></a>

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
low_security = {}
```

<a id="canonical-b32fcd1969034fea11f7d7edcee98b2085418fc027f9f23247f3318789ada028"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_config.low_security / ac85e870c827 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1dd69e9d5e99eca6db3c72ceb079f62de79431cbddb9ab1cfe568a3fe7253875"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_config.low_security / ac85e870c827 / 4

- [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-d1a648731a53dfe450f2260c554666be91a89d63d41248e33bc92e52d4158406)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-50a8ccd45d40e7b917a226bd46573c174096e6dcb79fd0b9cb0376caa9100243"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d51dc3ff3e39187f800e2f477372a74886c4ad662c421091ec5eed14ab8e4485"></a>

## https_management.advertise_on_sli_vip.tls_config.medium_security — https_management.advertise_on_sli_vip.tls_config.medium_security / 1c9bbc47464e / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-d1a648731a53dfe450f2260c554666be91a89d63d41248e33bc92e52d4158406)
- https_management.advertise_on_sli_vip.tls_config.medium_security

<a id="canonical-469ace6133b155baf459cd23a6f7672bccc574852c240974e32a9e2e4272296d"></a>

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
medium_security = {}
```

<a id="canonical-58327d2936887cd6bc70c15cdcd67620b69ffda53418613023056f1af03af8c4"></a>

## Direct properties — https_management.advertise_on_sli_vip.tls_config.medium_security / 1c9bbc47464e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c0415876b208147f8dc773f14c45bbe978dd08e07cc3c97a5caaa16c3a9da1b"></a>

## Next pages — https_management.advertise_on_sli_vip.tls_config.medium_security / 1c9bbc47464e / 4

- [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-d1a648731a53dfe450f2260c554666be91a89d63d41248e33bc92e52d4158406)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-16dcf258eb9d3162e1c1334cc2071354f4a1d2c4cfac0e190ad103f231eeab96"></a>

## https_management.advertise_on_sli_vip.use_mtls — https_management.advertise_on_sli_vip.use_mtls / cae18441badf / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- https_management.advertise_on_sli_vip.use_mtls

<a id="canonical-b0514d0ec1db1c0428aeb098a8d4395a219dbe44861f5a29dc90ab96e828a08c"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-1dd58b2651e8647e76d789177b968cceabdcfcb4cd974bc2c8de4f1c41e0b958"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls / cae18441badf / 3

<a id="canonical-7a2fb951508c67e93bf4d2717cf6d140539d8173f25dc9fe1fcbefc0a4a0fd43"></a>

<a id="canonical-427b44cd20a79c051ae81c2419bd01ca1e45ceb08beae011b3b7758898b65f87"></a>

## client_certificate_optional property — https_management.advertise_on_sli_vip.use_mtls / cae18441badf / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--nfv_service--reference--group-002.md#canonical-c57a7d7c0c1491b573717235f4baffc2d9c2386e7d086552830270c515d7ab13): complete subsection reference.

- [no_crl](resources--nfv_service--reference--group-002.md#canonical-09d2f82dc2fccecf9e16cb736f996ca5e29634d3262a448c067d676145214a13): complete subsection reference.

- [trusted_ca](resources--nfv_service--reference--group-002.md#canonical-eaf6ec402668374a64dd8655e1e48449cf65ffae9619a43bd8d38470353c5b4e): complete subsection reference.

<a id="canonical-242a86d2392c33e6242d93798c73263483cceb08aba34216a5a51788099c1b82"></a>

<a id="canonical-7ae42f9ea7cb685fd31c77232e84a6c38185c4ebeb519a90550201c64a4204b0"></a>

## trusted_ca_url property — https_management.advertise_on_sli_vip.use_mtls / cae18441badf / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--nfv_service--reference--group-002.md#canonical-bf681d9023b3c18ab521bc8447c0b8918d67785543200ccc9f304378d5805c8b): complete subsection reference.

- [xfcc_options](resources--nfv_service--reference--group-002.md#canonical-f39f61c37bfc321d1929fe4b8fc5c11d4d8f01ac8e9eed06cf8dc88125ccb99d): complete subsection reference.

<a id="canonical-ade168c1099d48e60eae5731f2e7a3cc8e9f40bb1afe84757d443fa79bfffbd7"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls / cae18441badf / 6

- [https_management.advertise_on_sli_vip.use_mtls.crl](resources--nfv_service--reference--group-002.md#canonical-c57a7d7c0c1491b573717235f4baffc2d9c2386e7d086552830270c515d7ab13)
- [https_management.advertise_on_sli_vip.use_mtls.no_crl](resources--nfv_service--reference--group-002.md#canonical-09d2f82dc2fccecf9e16cb736f996ca5e29634d3262a448c067d676145214a13)
- [https_management.advertise_on_sli_vip.use_mtls.trusted_ca](resources--nfv_service--reference--group-002.md#canonical-eaf6ec402668374a64dd8655e1e48449cf65ffae9619a43bd8d38470353c5b4e)
- [https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-002.md#canonical-bf681d9023b3c18ab521bc8447c0b8918d67785543200ccc9f304378d5805c8b)
- [https_management.advertise_on_sli_vip.use_mtls.xfcc_options](resources--nfv_service--reference--group-002.md#canonical-f39f61c37bfc321d1929fe4b8fc5c11d4d8f01ac8e9eed06cf8dc88125ccb99d)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-c57a7d7c0c1491b573717235f4baffc2d9c2386e7d086552830270c515d7ab13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4991f5050e59d8edf265f2c2d6853511979b295970a1b33f75d200340e65a31a"></a>

## https_management.advertise_on_sli_vip.use_mtls.crl — https_management.advertise_on_sli_vip.use_mtls.crl / 90faf245ce38 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe)
- https_management.advertise_on_sli_vip.use_mtls.crl

<a id="canonical-48f539fd2f28a8cc222b6e4c5fdf8542601d3e191cb733da4c67b523a2e62df0"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-722321cbe46786f6144c035350e0cda9caec86af92635228966738dc01d5c9e1"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls.crl / 90faf245ce38 / 3

<a id="canonical-e7113fedeee1fd743eb1126a52111bb84b1d4052311b2e3cec3250be0bb36bac"></a>

<a id="canonical-7d23093f0ffb0a44d79f567a0fe968e20f8a64e0b12fa97f12fbc39dab30e6e4"></a>

## name property — https_management.advertise_on_sli_vip.use_mtls.crl / 90faf245ce38 / 4

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

<a id="canonical-a7c8c2755651e6eb088860076630c1c15084c6f3c7327bfb6ef76beb663e77c1"></a>

<a id="canonical-aa11ce6fa43315710f7ff47fbfd6b6c7c6172830b711a48cba62679d8087aac8"></a>

## namespace property — https_management.advertise_on_sli_vip.use_mtls.crl / 90faf245ce38 / 5

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

<a id="canonical-785085bc37a9bd46be707e2cfea9f4de00b67658c3396865585109adb366132a"></a>

<a id="canonical-d205929e46bbc02573ebd8636d8a0c0943274eb1e5528c91cea5b7af94e920fd"></a>

## tenant property — https_management.advertise_on_sli_vip.use_mtls.crl / 90faf245ce38 / 6

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

<a id="canonical-49e64b5c15c472cd38fa951709f0980697934fcbec311cc19d8635f7175a8a76"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls.crl / 90faf245ce38 / 7

- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-09d2f82dc2fccecf9e16cb736f996ca5e29634d3262a448c067d676145214a13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-07b6741a2fcc61ed9b0cc33a35f5250d73e2190f7ff4efe23198b4dc0ef79d2f"></a>

## https_management.advertise_on_sli_vip.use_mtls.no_crl — https_management.advertise_on_sli_vip.use_mtls.no_crl / 0fda636d5573 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe)
- https_management.advertise_on_sli_vip.use_mtls.no_crl

<a id="canonical-97bd1ab411dca56becfac6631a7783de0226668dacb72c805da23706c729fe0c"></a>

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
no_crl = {}
```

<a id="canonical-6ec8a98a3eb4594df7763e9feabf188c35b11124a6d3506e6dec948ba4b3bed8"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls.no_crl / 0fda636d5573 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-f91d82b0851b320beec1c135c0610be7fe7126f81a9ac2c0e701c7b9a5009d9c"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls.no_crl / 0fda636d5573 / 4

- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-eaf6ec402668374a64dd8655e1e48449cf65ffae9619a43bd8d38470353c5b4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-073c1f7479bf17d7a5459e1aa44fbf8367930e465fdb12c859a767bda989d59a"></a>

## https_management.advertise_on_sli_vip.use_mtls.trusted_ca — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / e5823b295406 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe)
- https_management.advertise_on_sli_vip.use_mtls.trusted_ca

<a id="canonical-559678e572c7aa08a07c920a647fa9e60555651a07da394a4c9aedb3af899647"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-3f217126e089b584bb8e5201dfdce427553cecfc2165a3ddbc34a4587d3b22ec"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / e5823b295406 / 3

<a id="canonical-e8f2a004bab40706989237bbc2d694f633bf215b45c9c811faa357d5e9ab78cf"></a>

<a id="canonical-0cdc0df99f87331b82387599d4565a1f774c9e2e172aa0f499c93bbd03d0079b"></a>

## name property — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / e5823b295406 / 4

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

<a id="canonical-2da27d96403d003cf1639908959cda77e9a83ca6060c4808b7026a5817de1414"></a>

<a id="canonical-6bb6e56f9d230179883889c1c7d0d04cfd0a796f2c2b6641c3b9ebf85653d1aa"></a>

## namespace property — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / e5823b295406 / 5

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

<a id="canonical-49a3f39e0a20732157f1ada4848581795610ef67e771947da6c51b8cc3950f0e"></a>

<a id="canonical-80603e47952f17d3dd2415b9481bfd92d5342532ff3bbe6b8b4f7303dd88ca92"></a>

## tenant property — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / e5823b295406 / 6

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

<a id="canonical-766d6553dc7cc186cd3a7d283dacc4e516c5996ad0d573d0edad552176e61833"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls.trusted_ca / e5823b295406 / 7

- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-bf681d9023b3c18ab521bc8447c0b8918d67785543200ccc9f304378d5805c8b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0f97386bf8bb5682c2c0947ba02b47f1e8fb94787fad6c27465667a785a31d2d"></a>

## https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled — https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled / b2dee5498f68 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe)
- https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled

<a id="canonical-3c5815595dff603936215bc5f8318fb842f148a17f7e99648ea2b7ffcca1ad4b"></a>

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
xfcc_disabled = {}
```

<a id="canonical-fd27353e27250387b93918f4da1edf7a926a1de5e206525e142d43a99501396f"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled / b2dee5498f68 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9acfd857de22201216a27ba0feae59094e176378c47cf06c94709c4f2c370624"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls.xfcc_disabled / b2dee5498f68 / 4

- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-f39f61c37bfc321d1929fe4b8fc5c11d4d8f01ac8e9eed06cf8dc88125ccb99d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ad827ce6354de96f902c528f5091d7bd300712b6d60c581027aef53d3043c310"></a>

## https_management.advertise_on_sli_vip.use_mtls.xfcc_options — https_management.advertise_on_sli_vip.use_mtls.xfcc_options / 745fba32cb85 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_sli_vip](resources--nfv_service--reference--group-002.md#canonical-3d004d84a43cb893970b0b31ca50bd31ae8d74c3488818a4f99ac8a4127fb3f2)
- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe)
- https_management.advertise_on_sli_vip.use_mtls.xfcc_options

<a id="canonical-fb512985260795fd1c0cc4799711030e224a425c405ffe727011d6cf06f5542b"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-6e25e5c6c2c6fa56304d63577676c04abf2a90f0d1921faaa473ce2af77e3a77"></a>

## Direct properties — https_management.advertise_on_sli_vip.use_mtls.xfcc_options / 745fba32cb85 / 3

<a id="canonical-f695ad14a2c01b14887dda2c55eba1786331cf8d46522d545e89f36f270686c0"></a>

<a id="canonical-14d2ab5b8895217ec7fb38a33020de880da4fd8bdea83e454aa8d53c02a7d271"></a>

## xfcc_header_elements property — https_management.advertise_on_sli_vip.use_mtls.xfcc_options / 745fba32cb85 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-5a9a70d6531d3d31dbceb1303a1c091aa33bcf7ba4b4c524a8626f199da9cb1e"></a>

## Next pages — https_management.advertise_on_sli_vip.use_mtls.xfcc_options / 745fba32cb85 / 5

- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-fb9914069882c4cc2452fe211af3391d1cddb6227642ecb391994a3d82856dfe)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6563bb9194e106442b61c056d0b04a355b737edc3f99069bcfce5715ccd8f165"></a>

## https_management.advertise_on_slo_internet_vip — https_management.advertise_on_slo_internet_vip / 70032fce3748 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- https_management.advertise_on_slo_internet_vip

<a id="canonical-000f5f6ab55ffd53b27bf332ae67a8bb00cf69f053e1fecd012a61caf6dab875"></a>

Type: `"object"`. single nested block, Optional.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
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
advertise_on_slo_internet_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-ac91a71ea13174655b5b14294173b5d2bcd98865d13436ec8d1668c3953f32bd"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip / 70032fce3748 / 3

- [no_mtls](resources--nfv_service--reference--group-002.md#canonical-541daa691bfc0c415495a9f8b931ec3130f375f4157a4b3b1e96d9059bce1671): complete subsection reference.

- [tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d): complete subsection reference.

- [tls_config](resources--nfv_service--reference--group-002.md#canonical-677b7827dbbe8f90d4bf9632f2f0580748009c49b5b52bd1de17d2d308d5b06d): complete subsection reference.

- [use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6): complete subsection reference.

<a id="canonical-901e09ccbb4799e0c94d0acecd0b868235134e0c063ea4e07053c3cb7afd03ea"></a>

## Next pages — https_management.advertise_on_slo_internet_vip / 70032fce3748 / 4

- [https_management.advertise_on_slo_internet_vip.no_mtls](resources--nfv_service--reference--group-002.md#canonical-541daa691bfc0c415495a9f8b931ec3130f375f4157a4b3b1e96d9059bce1671)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d)
- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-677b7827dbbe8f90d4bf9632f2f0580748009c49b5b52bd1de17d2d308d5b06d)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-541daa691bfc0c415495a9f8b931ec3130f375f4157a4b3b1e96d9059bce1671"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e946086543f62da7822e654470be68522bfadca3a43ddd42cc279b328fd3db47"></a>

## https_management.advertise_on_slo_internet_vip.no_mtls — https_management.advertise_on_slo_internet_vip.no_mtls / 68e2c667eeab / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- https_management.advertise_on_slo_internet_vip.no_mtls

<a id="canonical-e5f3b0f001919704b049448b2e7a0b7067d066358ed7a514e97280687328baa9"></a>

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

<a id="canonical-a94c9e4bf233ce0dd411b10d83d1ff86c829f5e2d2beca3af2d913cb87a500cf"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.no_mtls / 68e2c667eeab / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-71122179514371e15be9ee85fc3133fb9377a2d54fc2cf4ebab9443168c71511"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.no_mtls / 68e2c667eeab / 4

- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b04f488eac48c951c65aba1c1c32b42572904860c23cd3bfc7a8c43b22a16ed"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates — https_management.advertise_on_slo_internet_vip.tls_certificates / 2e3ce86f855e / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- https_management.advertise_on_slo_internet_vip.tls_certificates

<a id="canonical-3f978cad766147c3a88d222418b99b63af2c330fe53ca6a634411a53db67870f"></a>

Type: `"object"`. list nested block, Optional.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-4c384db51f3069914da49f982e573bb78df5421c2269fb44a7856b9a142ff463"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates / 2e3ce86f855e / 3

<a id="canonical-2f4ece0bc719c5541a0650e1264cbb54beb0462038cf08ac29f801c6bfb77b34"></a>

<a id="canonical-cf3dceb059b8d3453d57243444c9c227f84e78d63f8f1e47294d5961feadbd54"></a>

## certificate_url property — https_management.advertise_on_slo_internet_vip.tls_certificates / 2e3ce86f855e / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--nfv_service--reference--group-002.md#canonical-ef7bd8ec8505b278ecb4061c15b0047ac1f13915515ecbee660b5c66f061abde): complete subsection reference.

<a id="canonical-3af7b41d8d10309b278922da334bf7ee0e1c103f9a04b5c221e02dae1c3474cf"></a>

<a id="canonical-a589e42d6b58abb7801c90e815b42ba65554e55d293bd868c3343f1cb50199a9"></a>

## description_spec property — https_management.advertise_on_slo_internet_vip.tls_certificates / 2e3ce86f855e / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--nfv_service--reference--group-002.md#canonical-bddb3474690c622f3726f77d0fe9d6bc55fecfcbf332cf62cf76d8ebb3506c33): complete subsection reference.

- [private_key](resources--nfv_service--reference--group-002.md#canonical-353ec4349becfb16fe48f11c6fa6a37b6124390807a2582015eb4a04fc240746): complete subsection reference.

- [use_system_defaults](resources--nfv_service--reference--group-002.md#canonical-3bf379d049f3da3042e6de1c5e0b70be1387bc8ff9ea20811fb763781b468652): complete subsection reference.

<a id="canonical-573350bf1d6eaaf81f007fe9134f9a7d3524311018ee5246232d7d632c16a415"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates / 2e3ce86f855e / 6

- [https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms](resources--nfv_service--reference--group-002.md#canonical-ef7bd8ec8505b278ecb4061c15b0047ac1f13915515ecbee660b5c66f061abde)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling](resources--nfv_service--reference--group-002.md#canonical-bddb3474690c622f3726f77d0fe9d6bc55fecfcbf332cf62cf76d8ebb3506c33)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-353ec4349becfb16fe48f11c6fa6a37b6124390807a2582015eb4a04fc240746)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults](resources--nfv_service--reference--group-002.md#canonical-3bf379d049f3da3042e6de1c5e0b70be1387bc8ff9ea20811fb763781b468652)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-ef7bd8ec8505b278ecb4061c15b0047ac1f13915515ecbee660b5c66f061abde"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18ea424b0783af88b09f557ad775a479f2fbefe4a3b21cb8bda8f22b95503dcd"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms — https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algo / faa5b028940d / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d)
- https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algorithms

<a id="canonical-71a7270ad0146bc368d3039d5dd4883a4f46f2339ffdae36f67cd4309d0cd399"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-19fe4aa6a3e16e5b2b82b8f8338d773b4e44847cf48185e78b02a8e5da63be06"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algo / faa5b028940d / 3

<a id="canonical-2263253804d8a554fe7f8b324fa6b6c58ed7014c3b97e1ccb6eb1a5ad9c5c0f0"></a>

<a id="canonical-43783b2360a86351ea614ed347d758c01b289ec6ea8955fa92700854ef0e6713"></a>

## hash_algorithms property — https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algo / faa5b028940d / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-a25ec84c0fb412d3d0d88fccafad244e2576153150e6f0155cf249ff42aad4ed"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.custom_hash_algo / faa5b028940d / 5

- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-bddb3474690c622f3726f77d0fe9d6bc55fecfcbf332cf62cf76d8ebb3506c33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4aedced9c3b416cfa8eabf9a375fcf9fcb27a55543dcdad2fa9dd040855637a6"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling — https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_sta / 07e0ded5bd44 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d)
- https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_stapling

<a id="canonical-030fd2cb1f48a310b8db555074259487df3a64b9d24b235ef1f3f3a9934adb0d"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-addbb480ab795997c7346b93c748a569ab190f6ae85ce140628f23e56b9103b1"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_sta / 07e0ded5bd44 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c7697cd43e0e13bbf9cdbcf2176db26b265acc1bd6887e2a111d8ecf869ca033"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.disable_ocsp_sta / 07e0ded5bd44 / 4

- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-353ec4349becfb16fe48f11c6fa6a37b6124390807a2582015eb4a04fc240746"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e19ab9c770e07819b001f823b516e50c2c8bdd1677faece8dfbac2e49efda35b"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.private_key — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key / b355910f8d38 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d)
- https_management.advertise_on_slo_internet_vip.tls_certificates.private_key

<a id="canonical-cd90a3ec39661810f0ed3d28226e25d1bf9164d7a3c62da8280526d9779dd30d"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-15a35b0d4f093203565d6dd0f60ea8d43279ab40378f6aa6d21628ed2a66e0f1"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key / b355910f8d38 / 3

- [blindfold_secret_info](resources--nfv_service--reference--group-002.md#canonical-36bee8c58e476379064a51b79c9d0c6dd620743ad286678754664c5d6a7b4970): complete subsection reference.

- [clear_secret_info](resources--nfv_service--reference--group-002.md#canonical-8a48a77618ab7d3e722839ae937c4fbcd6a79dc4a240a0b661959f00dbc4c6a1): complete subsection reference.

<a id="canonical-9447b7342d6c4390278ced8808ede20555a1a8f8de03a4e059cffa9163bb0140"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key / b355910f8d38 / 4

- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info](resources--nfv_service--reference--group-002.md#canonical-36bee8c58e476379064a51b79c9d0c6dd620743ad286678754664c5d6a7b4970)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info](resources--nfv_service--reference--group-002.md#canonical-8a48a77618ab7d3e722839ae937c4fbcd6a79dc4a240a0b661959f00dbc4c6a1)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-36bee8c58e476379064a51b79c9d0c6dd620743ad286678754664c5d6a7b4970"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-843e88ffabb662faf3acd11336c647c0cdff18509c9ac665a5644b2c7ddd3e10"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 6fbae446d6d9 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-353ec4349becfb16fe48f11c6fa6a37b6124390807a2582015eb4a04fc240746)
- https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-bb7286b22e7c70a35e84661ff4c521a464c4dd6ae52947632370805cbfe78543"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-f598cb7f09a57a81e8eca160c903c7c329600652fd977a8faaf01065772516b8"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 6fbae446d6d9 / 3

<a id="canonical-bf869e8f447ae0abe94cbf1cbe7984b5e82e220972133283a3d0c5bd8d9c3506"></a>

<a id="canonical-778e280a1d3362a11a30b2654944feaad2ba15a149c836969e9ca36ec0f6689a"></a>

## decryption_provider property — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 6fbae446d6d9 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-b68a1cb45844d2ca3e19a4f994a7ed30f58d4cbdf904118cb1f805bd3fd5e379"></a>

<a id="canonical-2860faaa89d394b9ae902563b58e13f781a3d59d7efc0bd1b8624b9a0315643f"></a>

## location property — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 6fbae446d6d9 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-9bf1ec2c62b32873868cb869290822bc786e4701e5b2f6a23c1fe3b005dcf8bd"></a>

<a id="canonical-c88a65d9caaf75e836f71e53a7793e5d1431c0473ab62890782c2cec3f5aa7b7"></a>

## store_provider property — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 6fbae446d6d9 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-e53961ee7bf41244f0212b1b97c382ab79f0da2c6c2fd89327db33240581d6c1"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.blin / 6fbae446d6d9 / 7

- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-353ec4349becfb16fe48f11c6fa6a37b6124390807a2582015eb4a04fc240746)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-8a48a77618ab7d3e722839ae937c4fbcd6a79dc4a240a0b661959f00dbc4c6a1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a79b0491b48951793c1a7fa726ee967c42443776889e08042932257759fe6917"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clea / 83ba6fc872a3 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d)
- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-353ec4349becfb16fe48f11c6fa6a37b6124390807a2582015eb4a04fc240746)
- https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clear_secret_info

<a id="canonical-c8a8cc4a713602948b7777447730082095a2f6948fbacd994c6c3dd162f06119"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-d21fd7c4cd4556de7dfd56fe098ce45231c892e3576c035ffcc6283268466cfe"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clea / 83ba6fc872a3 / 3

<a id="canonical-39fb675559a22b2f2a9975e19997ae4caeb85fb5de9a954ca06edf21ac397450"></a>

<a id="canonical-25a46243360c292144c12fe95b1701cbef5b0687815329d68e662b103457c245"></a>

## provider_ref property — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clea / 83ba6fc872a3 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-90754b90a9c50a83d588234887364a069217346ce4b37a935c7ad2a2ebbb9295"></a>

<a id="canonical-8468a46dd1241abb88a967bd75ce8f9b9a1f97032abc7b68937a96fe315f4295"></a>

## url property — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clea / 83ba6fc872a3 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-4888671226e928ada07b9d20af941c178d4360083886bfc25319981e89d7e7df"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.private_key.clea / 83ba6fc872a3 / 6

- [https_management.advertise_on_slo_internet_vip.tls_certificates.private_key](resources--nfv_service--reference--group-002.md#canonical-353ec4349becfb16fe48f11c6fa6a37b6124390807a2582015eb4a04fc240746)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-3bf379d049f3da3042e6de1c5e0b70be1387bc8ff9ea20811fb763781b468652"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2389aacdd3efc0e98220b8c18af468cb0b46c822f8e3b74fa9f43dd055aad3b9"></a>

## https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults — https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defau / 2d96b5776a47 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d)
- https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defaults

<a id="canonical-285b5caf9a9977872fb61c9156aedc628e17504cadfd571794da69e8b6c41530"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-309fbad66238be3d91bf8e95a91230b608cb581d3150a1e08d57acd3b5a54678"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defau / 2d96b5776a47 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-cc0af78af921c5900a26da023ff470f17efd81dd60997429fe92caa824ab92d9"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_certificates.use_system_defau / 2d96b5776a47 / 4

- [https_management.advertise_on_slo_internet_vip.tls_certificates](resources--nfv_service--reference--group-002.md#canonical-c02aa8cfed2138a7cb92ea25cdeddfc2bcd8e756a4afcc4e2afcbaf0cd54957d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-677b7827dbbe8f90d4bf9632f2f0580748009c49b5b52bd1de17d2d308d5b06d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80b29f502837c2abe995077441bc489f260bcc50a6a0f74bd9c761dd1c025542"></a>

## https_management.advertise_on_slo_internet_vip.tls_config — https_management.advertise_on_slo_internet_vip.tls_config / 08c96f73283c / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- https_management.advertise_on_slo_internet_vip.tls_config

<a id="canonical-7a4264429185ac5cbee11a32acf90af12bd5ee04265683a8755f14397de8dab8"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-f93294aede2e2ccf6f7970490eb988e33e894b13d9a33f683674631af4cb22f8"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_config / 08c96f73283c / 3

- [custom_security](resources--nfv_service--reference--group-002.md#canonical-5d47fb7ef012272c1ec3bbce8049918cb72894174593189a54589d5cf4f1c703): complete subsection reference.

- [default_security](resources--nfv_service--reference--group-002.md#canonical-54e07ec5692b12867add438d179968abae6c8f1393be2a204f7b7fa50851bb80): complete subsection reference.

- [low_security](resources--nfv_service--reference--group-002.md#canonical-470c03c857b80934b98c1e0169d7bc0998e3bb6f0a8f764547aae7f95a8dbb30): complete subsection reference.

- [medium_security](resources--nfv_service--reference--group-002.md#canonical-d8cafba8444f38019342e3c6e7b7a30b79168b0385a29717119564f98e3075c8): complete subsection reference.

<a id="canonical-0da5cee2f15e591e42d1bc9f0e2d44d5d96c1f704e982ce79ae5425ec7c6fe06"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_config / 08c96f73283c / 4

- [https_management.advertise_on_slo_internet_vip.tls_config.custom_security](resources--nfv_service--reference--group-002.md#canonical-5d47fb7ef012272c1ec3bbce8049918cb72894174593189a54589d5cf4f1c703)
- [https_management.advertise_on_slo_internet_vip.tls_config.default_security](resources--nfv_service--reference--group-002.md#canonical-54e07ec5692b12867add438d179968abae6c8f1393be2a204f7b7fa50851bb80)
- [https_management.advertise_on_slo_internet_vip.tls_config.low_security](resources--nfv_service--reference--group-002.md#canonical-470c03c857b80934b98c1e0169d7bc0998e3bb6f0a8f764547aae7f95a8dbb30)
- [https_management.advertise_on_slo_internet_vip.tls_config.medium_security](resources--nfv_service--reference--group-002.md#canonical-d8cafba8444f38019342e3c6e7b7a30b79168b0385a29717119564f98e3075c8)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-5d47fb7ef012272c1ec3bbce8049918cb72894174593189a54589d5cf4f1c703"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1853bf6ef9ce9c284c1cb3f421796e96a5c03216efc8137fd025637afa84caa4"></a>

## https_management.advertise_on_slo_internet_vip.tls_config.custom_security — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / 66214e377964 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-677b7827dbbe8f90d4bf9632f2f0580748009c49b5b52bd1de17d2d308d5b06d)
- https_management.advertise_on_slo_internet_vip.tls_config.custom_security

<a id="canonical-54993e8047f3a1b7fb616cd7ea12583d98b70d5d75d616ce5aa66768591db9af"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-b432d5f38f452f877d52ffba34ec0b45f399dde1835c535a5c0b2fdd5ef9a156"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / 66214e377964 / 3

<a id="canonical-a178516ed5942be1ebc3b0199d083ffbd2aa552d092e42ce7c3b0be39da18509"></a>

<a id="canonical-31763e86ab998af84983c112c13526809a54c60030097ecb39f5c404112a0fd3"></a>

## cipher_suites property — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / 66214e377964 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-112f0db57388611891b59cb23bbd7a43990d9938031bf0bab2ba40f2399aa253"></a>

<a id="canonical-32db31727675176bb0b7a9271c3888bc7e7b3946680bc04afba568cf3d33545f"></a>

## max_version property — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / 66214e377964 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-4cccec009779c07c949028c91278d994aa7118d8cee0cba6f1f7c13ca0f9df03"></a>

<a id="canonical-7277a62c62fbfbfbdde84c25e30fd3021fbbca515505933e1e16f5c8ec006035"></a>

## min_version property — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / 66214e377964 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-054c4893429dc4e736a7f37c469a3d7210e41acf72fddef6170cdf4b724fece9"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_config.custom_security / 66214e377964 / 7

- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-677b7827dbbe8f90d4bf9632f2f0580748009c49b5b52bd1de17d2d308d5b06d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-54e07ec5692b12867add438d179968abae6c8f1393be2a204f7b7fa50851bb80"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4baea7592a266481050034c0a889d3169cc67578271eb121c939066581a42758"></a>

## https_management.advertise_on_slo_internet_vip.tls_config.default_security — https_management.advertise_on_slo_internet_vip.tls_config.default_security / c9d18f4fce12 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-677b7827dbbe8f90d4bf9632f2f0580748009c49b5b52bd1de17d2d308d5b06d)
- https_management.advertise_on_slo_internet_vip.tls_config.default_security

<a id="canonical-9527220466bde32fcf88ca2b5bd2fba29f2f010b2f54e208e670a492284d31d9"></a>

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
default_security = {}
```

<a id="canonical-471637bc5caf03d3a00a10cca80ae29df66eb2c422ee965ee610afca675ae38f"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_config.default_security / c9d18f4fce12 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-29b5cf7478fb7890ec4172ad69aa957befa9c6ed23a156c28b38fd5eb0b6de82"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_config.default_security / c9d18f4fce12 / 4

- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-677b7827dbbe8f90d4bf9632f2f0580748009c49b5b52bd1de17d2d308d5b06d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-470c03c857b80934b98c1e0169d7bc0998e3bb6f0a8f764547aae7f95a8dbb30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86d0ede2e9ddccfe8113785295c0dc06db05090a9ab00f805e43adc65cb8b589"></a>

## https_management.advertise_on_slo_internet_vip.tls_config.low_security — https_management.advertise_on_slo_internet_vip.tls_config.low_security / 5e5d9678b81a / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-677b7827dbbe8f90d4bf9632f2f0580748009c49b5b52bd1de17d2d308d5b06d)
- https_management.advertise_on_slo_internet_vip.tls_config.low_security

<a id="canonical-6f4c0e32310fe40980b9d6f8ff788202e98b5d3fc2c69128334dd6de35e5202b"></a>

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
low_security = {}
```

<a id="canonical-b26eb77a2d9f1e2a4d1e860fbed6317fb44d30dc96dce7c9aa83a67de73ed378"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_config.low_security / 5e5d9678b81a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3feb1598de95f0ea0e9d8abd4bf702d1488bc2bdb505784e1aebecbd2112ea23"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_config.low_security / 5e5d9678b81a / 4

- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-677b7827dbbe8f90d4bf9632f2f0580748009c49b5b52bd1de17d2d308d5b06d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-d8cafba8444f38019342e3c6e7b7a30b79168b0385a29717119564f98e3075c8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-00b688c5ad99caaa6cf2a4a5f27068f304cd60b1c42dc21d618ff1d17df2df34"></a>

## https_management.advertise_on_slo_internet_vip.tls_config.medium_security — https_management.advertise_on_slo_internet_vip.tls_config.medium_security / f231fa578b2e / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-677b7827dbbe8f90d4bf9632f2f0580748009c49b5b52bd1de17d2d308d5b06d)
- https_management.advertise_on_slo_internet_vip.tls_config.medium_security

<a id="canonical-2e2bc386781befbde2919ff6ec17f84173905b45fdfcf5c1c93f6d838dad3a62"></a>

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
medium_security = {}
```

<a id="canonical-849b547af7e748de781f9e2a7143fd421101b50ad17c119b239c0fb07d06e3e0"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.tls_config.medium_security / f231fa578b2e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-610637b65a496feae4961100d98d9f51a91abbe19390e5ee17d7aaaaa9e113ba"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.tls_config.medium_security / f231fa578b2e / 4

- [https_management.advertise_on_slo_internet_vip.tls_config](resources--nfv_service--reference--group-002.md#canonical-677b7827dbbe8f90d4bf9632f2f0580748009c49b5b52bd1de17d2d308d5b06d)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1cab41b35128ea989bcfd4eba1f8aba4c6dc01dac154dfbae56d3f4237284a7e"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls — https_management.advertise_on_slo_internet_vip.use_mtls / fa16bda4e9bf / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- https_management.advertise_on_slo_internet_vip.use_mtls

<a id="canonical-238239b99e0b186337027f89b497ff97e1e443be95bde40e0c1a4a2fb79031c6"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-ff0bb5caf2ca7b77b4b06bdc94c2a4e26a2c836fe2e96fa05ec3f9da9bc40b4c"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls / fa16bda4e9bf / 3

<a id="canonical-a5087ee407ed7bdb70fb6b529eefd9c3b3ad143513e57b72696d1570d6f7918e"></a>

<a id="canonical-01a565ac3ab7a6abea197b7675a54607d4ff8ce972f07ca4e92374f886d3d0de"></a>

## client_certificate_optional property — https_management.advertise_on_slo_internet_vip.use_mtls / fa16bda4e9bf / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](resources--nfv_service--reference--group-002.md#canonical-b6d43adeaa600223b9c59613ad611e74adbc0fd0c53565ec882aad5dd9c62d6c): complete subsection reference.

- [no_crl](resources--nfv_service--reference--group-002.md#canonical-c3378e3e39fc1ed44a8e75357b46ddcdc7a17307957943631258dd9e2cf9fdb3): complete subsection reference.

- [trusted_ca](resources--nfv_service--reference--group-002.md#canonical-3012dc52701ee049dc0d07b4383d37f1935e11b510dc3baf300dec71ca8b3c8c): complete subsection reference.

<a id="canonical-64a4bcc97fd004a2060e92e33287e1e449af093750d4b4e4085b293c596a8872"></a>

<a id="canonical-a374cb1858fc37668e40dbd06725cdabe79a8f63f14cf65cd6d3e7a6615442f5"></a>

## trusted_ca_url property — https_management.advertise_on_slo_internet_vip.use_mtls / fa16bda4e9bf / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--nfv_service--reference--group-002.md#canonical-d44a55bc5b2d81024fdab8a41b991bd5b0334d8f8c7bfd0d9041012a271f7e2d): complete subsection reference.

- [xfcc_options](resources--nfv_service--reference--group-003.md#canonical-a69d6df5da51dc13c9779e3e9d6ad788951c1a687499a422603c3abc63c8cd73): complete subsection reference.

<a id="canonical-eb56df00e28fa2a79691a67f4419b4a691b21b74b210b33612599ebd60eefdd7"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls / fa16bda4e9bf / 6

- [https_management.advertise_on_slo_internet_vip.use_mtls.crl](resources--nfv_service--reference--group-002.md#canonical-b6d43adeaa600223b9c59613ad611e74adbc0fd0c53565ec882aad5dd9c62d6c)
- [https_management.advertise_on_slo_internet_vip.use_mtls.no_crl](resources--nfv_service--reference--group-002.md#canonical-c3378e3e39fc1ed44a8e75357b46ddcdc7a17307957943631258dd9e2cf9fdb3)
- [https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca](resources--nfv_service--reference--group-002.md#canonical-3012dc52701ee049dc0d07b4383d37f1935e11b510dc3baf300dec71ca8b3c8c)
- [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled](resources--nfv_service--reference--group-002.md#canonical-d44a55bc5b2d81024fdab8a41b991bd5b0334d8f8c7bfd0d9041012a271f7e2d)
- [https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_options](resources--nfv_service--reference--group-003.md#canonical-a69d6df5da51dc13c9779e3e9d6ad788951c1a687499a422603c3abc63c8cd73)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-b6d43adeaa600223b9c59613ad611e74adbc0fd0c53565ec882aad5dd9c62d6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ac9921f96d8baeaca8b4bd6a2bd0f39f64ffaa94e7137aa65f14e1cb77af56b7"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.crl — https_management.advertise_on_slo_internet_vip.use_mtls.crl / 56ad0da1a545 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6)
- https_management.advertise_on_slo_internet_vip.use_mtls.crl

<a id="canonical-cc457d9583dbe52a37abe10fe2fc40db837b42d94664ecb811053a56fa1ccf8a"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-f854a1f943270cf21b183f2d20c949bf1a5459989aed2944be7890072cb9a6c2"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls.crl / 56ad0da1a545 / 3

<a id="canonical-0c7e64eea5fd33b579482d80f7d8f7283188bae34812f55c283a15451d465abd"></a>

<a id="canonical-660369be88f5619a803107b3a53ddfad4e5b3e7e0b7800b28d49cad8dca3179c"></a>

## name property — https_management.advertise_on_slo_internet_vip.use_mtls.crl / 56ad0da1a545 / 4

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

<a id="canonical-4f5063317c9b71b0226c530a9b9b882d45cfa6437deefc73460384f92f366f62"></a>

<a id="canonical-51cf74f3fe4540f3b38d85bce68617f70160a0d6f21e228932258296a999d49d"></a>

## namespace property — https_management.advertise_on_slo_internet_vip.use_mtls.crl / 56ad0da1a545 / 5

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

<a id="canonical-520882b7c29eadcc9a81e3eb88eb901db2b61df769bce1f8d3eade9e94d9e771"></a>

<a id="canonical-d24500fb81a0f6f65c690aba8ab13513f6dfdb91813d85fbc35b6e3af5577a6d"></a>

## tenant property — https_management.advertise_on_slo_internet_vip.use_mtls.crl / 56ad0da1a545 / 6

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

<a id="canonical-94ba86519dbe7b32f905d87c8b044b9db713b2082a76772c10d7a564e6c4578d"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls.crl / 56ad0da1a545 / 7

- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-c3378e3e39fc1ed44a8e75357b46ddcdc7a17307957943631258dd9e2cf9fdb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f5bd9bba5680d257485b0de3d9a7afde0b46177ee0467e96a273dbdcad5dcd14"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.no_crl — https_management.advertise_on_slo_internet_vip.use_mtls.no_crl / fbb1804515c7 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6)
- https_management.advertise_on_slo_internet_vip.use_mtls.no_crl

<a id="canonical-1d3f5a0415a5f101c780089a936cfedfe03b37bdd2a0c8c6c8b6796d95698909"></a>

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
no_crl = {}
```

<a id="canonical-50127d4daf863cff2a399be813858c595fab5046acc72ab3a0de09348d5dc969"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls.no_crl / fbb1804515c7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-389c75cf768bbf493edab90e7fa13776c07d9c5f76f567cfd6cfe73d3e659356"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls.no_crl / fbb1804515c7 / 4

- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-3012dc52701ee049dc0d07b4383d37f1935e11b510dc3baf300dec71ca8b3c8c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3916b4799007718c063dd5cd7e666fc153aa82a41c9d7e0d8368e364e1b52721"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 8f052bd9d569 / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6)
- https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca

<a id="canonical-e6814ab6b9764ab73914070e1ec90991be3ce4f57514d5d6683923fb9c8756e6"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-60aea2c28344ef0cc81f1cb0e4dd9ceb1e395b4631910e206eaa3a53782fea0f"></a>

## Direct properties — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 8f052bd9d569 / 3

<a id="canonical-e660aaeafb36729fbca234023e6741acda7d9403c3abddf5bde70d3e5e35ae20"></a>

<a id="canonical-b7a1f531977ee3e360534c9607d273f4b6f4ad4cfc191efc157b20728ec3e02b"></a>

## name property — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 8f052bd9d569 / 4

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

<a id="canonical-8c8bf9e443dd5125fa3fde937ce5fc88515a7541170599f6654269599d320679"></a>

<a id="canonical-71cc937b846dc45a2772e34b921eca858a44472321881fafadae93a866fd0b62"></a>

## namespace property — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 8f052bd9d569 / 5

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

<a id="canonical-cc5ac10d92a7cebe7d4e2d1e62920560d6aa28d8f440da037da664a660f4f6b1"></a>

<a id="canonical-a85610e8492367491915657f2b69824ccd10094b2962782d12c55d8c9be815b8"></a>

## tenant property — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 8f052bd9d569 / 6

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

<a id="canonical-902e91da60a3d0d497ff36bcfeba84f76d5f3f3efa62c1ca59482e45a5b95dfa"></a>

## Next pages — https_management.advertise_on_slo_internet_vip.use_mtls.trusted_ca / 8f052bd9d569 / 7

- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6)
- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)

<a id="canonical-d44a55bc5b2d81024fdab8a41b991bd5b0334d8f8c7bfd0d9041012a271f7e2d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5bae528683b5821823b03fba606db0f3052cdc0519325725f619e34c3f58195e"></a>

## https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled — https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled / 63d7369f7dea / 2

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md#canonical-2be57939079325a93785e24152f0bc745de9e02cf389bf33da3baca2a2ce1b2a)
- [Property reference](resources--nfv_service--reference--group-001.md#canonical-4f83719cb3e08028ce8f20e6ccecde153a38aedb2349dc1f6ee54b514aa84878)
- [https_management](resources--nfv_service--reference--group-002.md#canonical-86d8730512d9865d9a594d5452f337a54d17d7f2f539c7ddca065b32f4328a51)
- [https_management.advertise_on_slo_internet_vip](resources--nfv_service--reference--group-002.md#canonical-cd7a014a77a7464ee01bdc9918de6c755195ebcf97fd000b19523ae733884b9c)
- [https_management.advertise_on_slo_internet_vip.use_mtls](resources--nfv_service--reference--group-002.md#canonical-a8cd513a1608b86e1f196b42d35787d058f4210ab856187e96cf80249c2cffd6)
- https_management.advertise_on_slo_internet_vip.use_mtls.xfcc_disabled

<a id="canonical-decda601ba3a53169fe3362b578fd0988f98372ba453065c3c2b95a49a1abf57"></a>

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
xfcc_disabled = {}
```
