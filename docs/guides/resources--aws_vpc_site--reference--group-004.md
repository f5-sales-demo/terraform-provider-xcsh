---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-936294423554748d0bf1bc72c8c779409210ed7c9296728868c92a7f0a2fe978"></a>

## ingress_gw.allowed_vip_port.use_https_port — ingress_gw.allowed_vip_port.use_https_port / e12c72480a9a / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8)
- ingress_gw.allowed_vip_port.use_https_port

<a id="canonical-b4117099a0196cd06ebacd71fc3714a03147a56c60ba28ceea70d32227bdcb1c"></a>

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
use_https_port = {}
```

<a id="canonical-a157442c67993ebb2cbb68e2c9262d147568f3581e5735f43af3ba978c8ebb56"></a>

## Direct properties — ingress_gw.allowed_vip_port.use_https_port / e12c72480a9a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7df0754b9214dc2ef9385eee0f89c742fbc731879229def9b81c6c2084bfc6ce"></a>

## Next pages — ingress_gw.allowed_vip_port.use_https_port / e12c72480a9a / 4

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-003.md#canonical-a2dbe9904d7b369a6e2c5c9aaaeeb0b8b15a88c375b28cf16bc0e782212867d8)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-bc6cc2703fabd08d9aa7a7cb7be83e08640f98ae12f7c2610a4c69c813f5f964"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-159ec7ba5b40282e4333be5c6e19aa321710ca7626013a75009a7ee2adadcc1b"></a>

## ingress_gw.az_nodes — ingress_gw.az_nodes / 640df8083f8e / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- ingress_gw.az_nodes

<a id="canonical-88ccc2221b792b3a820c2621b124bfafe8ebfc64901be55fbf60aab179116472"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name")}
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

<a id="canonical-cbebf541b8ea57ca197fc6cace796ccfa78446ee5519a1629c19b07b862e944b"></a>

## Direct properties — ingress_gw.az_nodes / 640df8083f8e / 3

<a id="canonical-f68a641310fbdb4878575aa0028558df4b65b8357db8e12c32eddd883f3ee032"></a>

<a id="canonical-3360bf9fff2890a314fcad5cc0dd69780949a915be67169770a7f8bfd14022d1"></a>

## aws_az_name property — ingress_gw.az_nodes / 640df8083f8e / 4

Type: `"string"`. Optional.

AWS availability zone, must be consistent with the selected AWS region.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-72be73656603fb80a3b40abb6de4b6f9b2c4146627569fa6a038bec3f98d8c14): complete subsection reference.

<a id="canonical-4692b3123dcb72bbfbb70c5603e7e4da55b3534bde93c686d01cfd5749f49ceb"></a>

## Next pages — ingress_gw.az_nodes / 640df8083f8e / 5

- [ingress_gw.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-72be73656603fb80a3b40abb6de4b6f9b2c4146627569fa6a038bec3f98d8c14)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-72be73656603fb80a3b40abb6de4b6f9b2c4146627569fa6a038bec3f98d8c14"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-affa2d6b4e06bdb7311fe8a5536552906de04510403594b646efa9c6bcf7176b"></a>

## ingress_gw.az_nodes.local_subnet — ingress_gw.az_nodes.local_subnet / 559331392a45 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-bc6cc2703fabd08d9aa7a7cb7be83e08640f98ae12f7c2610a4c69c813f5f964)
- ingress_gw.az_nodes.local_subnet

<a id="canonical-b3e4287d40043082c5d3667dc2d00321d1e39a0b65baacb554b52fbb06902c20"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for local subnet.

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
local_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-f511f3f0c02ca3493a06f01e3a5d488e054d0610ebaca79ba339759f250a1b7a"></a>

## Direct properties — ingress_gw.az_nodes.local_subnet / 559331392a45 / 3

<a id="canonical-cd930ea66a8ec459283f9f841150f297f6b164527b4866211ea6ebb23311e6f5"></a>

<a id="canonical-bae78706098108742ff75916afc4252c76c0ca95fa49261ba15ef6d568474a60"></a>

## existing_subnet_id property — ingress_gw.az_nodes.local_subnet / 559331392a45 / 4

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

- [subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-a108833fac9b0ddb8bc0758d9bd83b0bae74617671c6f95ccfac91da827b37d9): complete subsection reference.

<a id="canonical-3e33ffb5ee36250d43936ed9ad26e1acffe4ec454ffe0833af7d42fa1ca3019f"></a>

## Next pages — ingress_gw.az_nodes.local_subnet / 559331392a45 / 5

- [ingress_gw.az_nodes.local_subnet.subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-a108833fac9b0ddb8bc0758d9bd83b0bae74617671c6f95ccfac91da827b37d9)
- [ingress_gw.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-bc6cc2703fabd08d9aa7a7cb7be83e08640f98ae12f7c2610a4c69c813f5f964)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-a108833fac9b0ddb8bc0758d9bd83b0bae74617671c6f95ccfac91da827b37d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08e2181fd83e6811867fe7ea23a247f5aa2ec3324babf4f463e17671fd95052b"></a>

## ingress_gw.az_nodes.local_subnet.subnet_param — ingress_gw.az_nodes.local_subnet.subnet_param / ea5384a49251 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-bc6cc2703fabd08d9aa7a7cb7be83e08640f98ae12f7c2610a4c69c813f5f964)
- [ingress_gw.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-72be73656603fb80a3b40abb6de4b6f9b2c4146627569fa6a038bec3f98d8c14)
- ingress_gw.az_nodes.local_subnet.subnet_param

<a id="canonical-9fb242d2561af114e71edac0e0d3f6a77dbc1b484f2c9ae204b8056740a36da6"></a>

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

<a id="canonical-8ab26feaaac2481bb7f582a9a87cfe976e36c5b74aba56e7f609e08bfb51626c"></a>

## Direct properties — ingress_gw.az_nodes.local_subnet.subnet_param / ea5384a49251 / 3

<a id="canonical-9a60cb3417f1667ee1cd12aefe1936e0895dcd1d14fa2d9e37622bdbea701338"></a>

<a id="canonical-6ed3c11b82b61e4af6024c3d35940ec049e8ef1f84b89d412b43ff17edfb9467"></a>

## ipv4 property — ingress_gw.az_nodes.local_subnet.subnet_param / ea5384a49251 / 4

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

<a id="canonical-22d7b8b37dc8bb2892a81c254b6197b0d72e34a70ca9bad3b08dd710ae901da4"></a>

## Next pages — ingress_gw.az_nodes.local_subnet.subnet_param / ea5384a49251 / 5

- [ingress_gw.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-72be73656603fb80a3b40abb6de4b6f9b2c4146627569fa6a038bec3f98d8c14)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-d755f6778136034ceb42d5cb5880580f5684cf6e38a0f35d217a6caad2f8c7a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ec9af2c82229b938f8739dfef7b4087c87c15d48c692b40ac5ba5f6a1fbc389"></a>

## ingress_gw.performance_enhancement_mode — ingress_gw.performance_enhancement_mode / 80acd19d9912 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- ingress_gw.performance_enhancement_mode

<a id="canonical-026eec207699d0b590341fb8e14158e1fe24dfd10a147b649280cc7a9d3eeea0"></a>

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

<a id="canonical-e752bab125ac300e376f860af2952142ccb39e1c4c092fd7acf90bec5350a913"></a>

## Direct properties — ingress_gw.performance_enhancement_mode / 80acd19d9912 / 3

- [perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-529eeb028ef3b3bda191f758f2324dbc1d5daf940b9357c2d75e28c0bc0e04ff): complete subsection reference.

- [perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-9303b08b2ffe45605f6076a9bbc40f7378ded742e7959a69b36068b3220721c4): complete subsection reference.

<a id="canonical-07fb210e6a6751d7b9eb547a0e329226a3da695384047e319d74b5296c6562de"></a>

## Next pages — ingress_gw.performance_enhancement_mode / 80acd19d9912 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-529eeb028ef3b3bda191f758f2324dbc1d5daf940b9357c2d75e28c0bc0e04ff)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-9303b08b2ffe45605f6076a9bbc40f7378ded742e7959a69b36068b3220721c4)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-529eeb028ef3b3bda191f758f2324dbc1d5daf940b9357c2d75e28c0bc0e04ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef539c5b4540d498f72974dcbf05bc9f839c8f4100b9cbb5713749326042d732"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / a04d57c5fa5f / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-d755f6778136034ceb42d5cb5880580f5684cf6e38a0f35d217a6caad2f8c7a8)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-682ce2c6cc9a7df315850e09dba04d4755fb297dde63965f22229d1de0477084"></a>

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

<a id="canonical-c86f010bf00d85310cf6fe946eff21e49476008aec4be5b98d19c6ccdeb698ba"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / a04d57c5fa5f / 3

- [jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-318da82773f6a4989d5ed3a1873d44de7bb71e1b6fcfa94ab8a2bebb8c708150): complete subsection reference.

- [no_jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-afda27e59626e8e1a9045ba917256a08aaeee0a2b8d8ee6d68b617695edaed88): complete subsection reference.

<a id="canonical-510f71b70d260f49c16c2927c6421a9fe10dbb1453e193e530b71ad3b80773a9"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced / a04d57c5fa5f / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-318da82773f6a4989d5ed3a1873d44de7bb71e1b6fcfa94ab8a2bebb8c708150)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-afda27e59626e8e1a9045ba917256a08aaeee0a2b8d8ee6d68b617695edaed88)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-d755f6778136034ceb42d5cb5880580f5684cf6e38a0f35d217a6caad2f8c7a8)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-318da82773f6a4989d5ed3a1873d44de7bb71e1b6fcfa94ab8a2bebb8c708150"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-497f5d55d21563c3da51fd73c0baaf5a4aca05c31b9c0a2eed2090c00ec61e7e"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / f52dc8650719 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-d755f6778136034ceb42d5cb5880580f5684cf6e38a0f35d217a6caad2f8c7a8)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-529eeb028ef3b3bda191f758f2324dbc1d5daf940b9357c2d75e28c0bc0e04ff)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-34bfa530c2440489a217aeb7f35a1775bc6de23fe5d444831babb18038670686"></a>

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

<a id="canonical-567a31d034969f2800d882befedd705a1d85d4616288b8667645b5a7e43729d2"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / f52dc8650719 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-716e94d9cb82c173f15dad5acf1eb5d3292719391c43cbe4e31b865efcb15067"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo / f52dc8650719 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-529eeb028ef3b3bda191f758f2324dbc1d5daf940b9357c2d75e28c0bc0e04ff)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-afda27e59626e8e1a9045ba917256a08aaeee0a2b8d8ee6d68b617695edaed88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7fcf3a71dbf87c8907e5b82d360359efd70c00782073768078381fa798a9f96f"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 763ff39def08 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-d755f6778136034ceb42d5cb5880580f5684cf6e38a0f35d217a6caad2f8c7a8)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-529eeb028ef3b3bda191f758f2324dbc1d5daf940b9357c2d75e28c0bc0e04ff)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-bd80bfbe2550c6a5123d269a560e376a6c615e412e6579d28ff89a4dba3783a3"></a>

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

<a id="canonical-ac60c10b7e890b7373ee33877f436832b10b08f388cc654b07222796ab09042b"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 763ff39def08 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d03e567615bfd48aef6dbb898d98a80f0c3adf0c4d64461973e6b5192f5f0d4b"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo / 763ff39def08 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-529eeb028ef3b3bda191f758f2324dbc1d5daf940b9357c2d75e28c0bc0e04ff)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-9303b08b2ffe45605f6076a9bbc40f7378ded742e7959a69b36068b3220721c4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-57ffb9ba4a398f0ce33b5c1d8ee8be33ae01a1835f16fe1afb6cf8c19f6e8ce9"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / a41145b123d4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-d755f6778136034ceb42d5cb5880580f5684cf6e38a0f35d217a6caad2f8c7a8)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-9e347e2c67f3aef6c447f630769b61f4fa023686cb6fd9278ac955e8bf9805c5"></a>

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

<a id="canonical-7a9c5e631dc004fcfba911a7fce452158e2c1cd10f80129ce2643d7992afa90e"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / a41145b123d4 / 3

- [jumbo_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-294ece327b8113bdccb3815f805676dcba6d2a4b60a159b0a0b51554f17eb2d9): complete subsection reference.

- [jumbo_enabled](resources--aws_vpc_site--reference--group-004.md#canonical-50f8ad98a640b842b95aa78d54cbba938efae318436385af569dbdab9b908ec6): complete subsection reference.

<a id="canonical-3ee6e390d71a5a65a3b964db25c86073b4bf37b9dbe1709c8d6d559074984216"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced / a41145b123d4 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-294ece327b8113bdccb3815f805676dcba6d2a4b60a159b0a0b51554f17eb2d9)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_vpc_site--reference--group-004.md#canonical-50f8ad98a640b842b95aa78d54cbba938efae318436385af569dbdab9b908ec6)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-d755f6778136034ceb42d5cb5880580f5684cf6e38a0f35d217a6caad2f8c7a8)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-294ece327b8113bdccb3815f805676dcba6d2a4b60a159b0a0b51554f17eb2d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-082df351bedb2d4c12bc80d7a90e6a4c5128a08bc99ddba5a059fd1e0b89e076"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 538c6ff8c88f / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-d755f6778136034ceb42d5cb5880580f5684cf6e38a0f35d217a6caad2f8c7a8)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-9303b08b2ffe45605f6076a9bbc40f7378ded742e7959a69b36068b3220721c4)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-dd2f6f1550004868db1a32696811ab5cd663787de4f33a11811c9b3229c21a73"></a>

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

<a id="canonical-4c1c2dd8eb1e665c11c67c0be5c2bc0515fa164255638e4c497327d5aba7a460"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 538c6ff8c88f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5ffacb75b534c9b3ee043d51f8d364b41e8c7f0b78187df5829c301a9bc4ee5e"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled / 538c6ff8c88f / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-9303b08b2ffe45605f6076a9bbc40f7378ded742e7959a69b36068b3220721c4)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-50f8ad98a640b842b95aa78d54cbba938efae318436385af569dbdab9b908ec6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d255d659754a97bdfda7fb7de095ced2e0d20a5662735b8625b45cd102d26be2"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 5aca6a9aea15 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-3ff8a0f91cfc1a12f0f56ca1ea38289886152f51e71944138a14d403a166acfb)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-d755f6778136034ceb42d5cb5880580f5684cf6e38a0f35d217a6caad2f8c7a8)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-9303b08b2ffe45605f6076a9bbc40f7378ded742e7959a69b36068b3220721c4)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-e270bfa7e387fbdaf0c1023af21b36ac928b601abab6589b5cf920ddae850970"></a>

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

<a id="canonical-2576e46a689b09a4648c15d233234eb95d7e6b42a2de2cf5cc22b107fc1fdd7e"></a>

## Direct properties — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 5aca6a9aea15 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-838a76a27e7c6fcc0f2b892c1d8c4f8703d69008f0a8d940fed5c6a711270e7d"></a>

## Next pages — ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled / 5aca6a9aea15 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-9303b08b2ffe45605f6076a9bbc40f7378ded742e7959a69b36068b3220721c4)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-c7c7ef7fa59a4e4ec7405e53292ad53333a6946b52c70bc83ae9f8f1d704859a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-efc961837e1f14217fc911a6a80c764a7ccb57ab8fa169fa673c204f99f16ad8"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 3ad80b115982 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- kubernetes_upgrade_drain

<a id="canonical-bbbe103d57e01fe88d1deb31dead94903cbc4c2f3f577734e220421bf26f1199"></a>

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

<a id="canonical-b2fc01194af6bb3034bdfc2a5129686949c5fbba3aac45250626248e1b649826"></a>

## Direct properties — kubernetes_upgrade_drain / 3ad80b115982 / 3

- [disable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-5f371783c48cb31a99c3b4948ecd12c775561018e5215e5819fc86ab685e24e4): complete subsection reference.

- [enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-5059daf3169902cf46fb2dffd91efd335943a25c3d03215b6f4f5b8e3c0ce6e7): complete subsection reference.

<a id="canonical-9b9ed019943f8d2bbbdc2badbba76a1e212e0f3ab7961c2eff4df2b4def93e56"></a>

## Next pages — kubernetes_upgrade_drain / 3ad80b115982 / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-5f371783c48cb31a99c3b4948ecd12c775561018e5215e5819fc86ab685e24e4)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-5059daf3169902cf46fb2dffd91efd335943a25c3d03215b6f4f5b8e3c0ce6e7)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-5f371783c48cb31a99c3b4948ecd12c775561018e5215e5819fc86ab685e24e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6a7cd643d94601b3130f7602cc0fa5c344d1eec4ccdb5f370d8557d6817ca53e"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — kubernetes_upgrade_drain.disable_upgrade_drain / f282cef788ce / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-c7c7ef7fa59a4e4ec7405e53292ad53333a6946b52c70bc83ae9f8f1d704859a)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-bb699eedb294da3fa3647539c2af9dd838e27f59ea9e3a4c12beccfdc3440776"></a>

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

<a id="canonical-f39d4859d519e5e369c4d374f247abdb57669b394522a5fe8de746277878336e"></a>

## Direct properties — kubernetes_upgrade_drain.disable_upgrade_drain / f282cef788ce / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-ca06a566acf9ef0a332a391cccb85bb26cc6b4116835d72773d6723b058beda1"></a>

## Next pages — kubernetes_upgrade_drain.disable_upgrade_drain / f282cef788ce / 4

- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-c7c7ef7fa59a4e4ec7405e53292ad53333a6946b52c70bc83ae9f8f1d704859a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-5059daf3169902cf46fb2dffd91efd335943a25c3d03215b6f4f5b8e3c0ce6e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0bbd2ee7eac2d91a7bd195525f074f7284ab53af73d103273e166a400de69330"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — kubernetes_upgrade_drain.enable_upgrade_drain / ce6ecd6c0df4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-c7c7ef7fa59a4e4ec7405e53292ad53333a6946b52c70bc83ae9f8f1d704859a)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-ff1fe051e8206594051ad1484106a0f9ea8d436e2287d594630896752a0d198c"></a>

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

<a id="canonical-e00daad3eb8a61de7e9f52de6653ab6b5538fdfa91c3dd8c7964448ba018aeff"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain / ce6ecd6c0df4 / 3

- [disable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-6edc332af79aa0ae3b26d85ae57e418a332d92c2fa63e2c171399da4c1fa1e68): complete subsection reference.

<a id="canonical-3366a04ccf400626abee7830a93ac63e27487bf583dd1ba6833fd34b32958df2"></a>

<a id="canonical-97c4edb473a08891c3d997bdfc605c757261389caf45ccabf0e92ef45a9db20d"></a>

## drain_max_unavailable_node_count property — kubernetes_upgrade_drain.enable_upgrade_drain / ce6ecd6c0df4 / 4

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

<a id="canonical-f85058ece4cd9bd38675a20ca0113096d1f9447e7a7a1a026803b9d7727e7508"></a>

<a id="canonical-d81a9b5fca4e68dd485def81ee1afb9df49b5b441f5cea9c1c28a7f0dfb31741"></a>

## drain_max_unavailable_node_percentage property — kubernetes_upgrade_drain.enable_upgrade_drain / ce6ecd6c0df4 / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-c861593d3eda4ee3499c5436d02bc775fd107523e0adfc61cd9494d210af083c"></a>

<a id="canonical-20966f3eeb7a4a450c0eaf4f134717f2e705e050d3f75bf6d54aea7f22bc91ee"></a>

## drain_node_timeout property — kubernetes_upgrade_drain.enable_upgrade_drain / ce6ecd6c0df4 / 6

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

- [enable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-9d224bcc6de32d8ce2e6f1f77acc78737343e3840cc481c0871823fa82e57288): complete subsection reference.

<a id="canonical-a83f54d4c52a5a430754a4d1ddd7c415001de8810a807f33927be50be49f619b"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain / ce6ecd6c0df4 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-6edc332af79aa0ae3b26d85ae57e418a332d92c2fa63e2c171399da4c1fa1e68)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-9d224bcc6de32d8ce2e6f1f77acc78737343e3840cc481c0871823fa82e57288)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-c7c7ef7fa59a4e4ec7405e53292ad53333a6946b52c70bc83ae9f8f1d704859a)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-6edc332af79aa0ae3b26d85ae57e418a332d92c2fa63e2c171399da4c1fa1e68"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-80b3261748bb17147d0b9b61948b5ddb79737ba58ee9623770d806f5e708a623"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 4c468a6f3bfc / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-c7c7ef7fa59a4e4ec7405e53292ad53333a6946b52c70bc83ae9f8f1d704859a)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-5059daf3169902cf46fb2dffd91efd335943a25c3d03215b6f4f5b8e3c0ce6e7)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-0e7343ad4575d1041ab6103ad5fe8bfe1d56c6933027c632f1e650036e4f7fa0"></a>

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

<a id="canonical-a47695e6081d0030164dbc814ff3e33eb76d149320dc777adae1c720f16b3afb"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 4c468a6f3bfc / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5152088a9e1cba767f1631fff3f8a3b306a1623787dadba4e2783ef72f73bbd8"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode / 4c468a6f3bfc / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-5059daf3169902cf46fb2dffd91efd335943a25c3d03215b6f4f5b8e3c0ce6e7)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-9d224bcc6de32d8ce2e6f1f77acc78737343e3840cc481c0871823fa82e57288"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1813f6a8a990db3de7f05a78b904dca00ee329a1417cd49d480e6f2a0e241248"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / c40c8d2f6684 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-c7c7ef7fa59a4e4ec7405e53292ad53333a6946b52c70bc83ae9f8f1d704859a)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-5059daf3169902cf46fb2dffd91efd335943a25c3d03215b6f4f5b8e3c0ce6e7)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-d8533ca08156bf9655f300bde1cb5297be84bc5838ba1dfe8734b4cf3679299f"></a>

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

<a id="canonical-afd551eccb1c1867dac94b1262004253f6ccb9385060b2ef6a3c3a919baa6f4c"></a>

## Direct properties — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / c40c8d2f6684 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89164f5b23331a77e7edd49c98d16b6080156a5fb03528c5f75cfc7152b56319"></a>

## Next pages — kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode / c40c8d2f6684 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-5059daf3169902cf46fb2dffd91efd335943a25c3d03215b6f4f5b8e3c0ce6e7)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-e8bd1fa12942769f9086e9fc72f0d6cf4d396489fde894196a56996dfd0a2061"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4d9f48a876698b948fe591efc560d8ab1230509e7c685220935406f541ce6162"></a>

## log_receiver — log_receiver / 5432735b3f8e / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- log_receiver

<a id="canonical-839cc3a1006b8375f716fe9c10910a865c449ba4b01872bdfee2b13f4df2b2ac"></a>

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

- [log_receiver](resources--aws_vpc_site--reference--group-004.md#canonical-839cc3a1006b8375f716fe9c10910a865c449ba4b01872bdfee2b13f4df2b2ac)
- [logs_streaming_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-82a912f582da1406a56e5a0614190eaf3ab538b3ac531c200be7b2efa7d9d193)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-5dcc830e47c14aded9fd2ef897086d32d3d5c650b8b32b64e6dcf0a3e6384f7a"></a>

## Direct properties — log_receiver / 5432735b3f8e / 3

<a id="canonical-3fc544cc704271d19d066e7e223cb63d703e82af5548762e980f5798bb5b2361"></a>

<a id="canonical-46def455b10cfc96c9f9e96fc463a34b81885497963f9fb26f7e925c87873135"></a>

## name property — log_receiver / 5432735b3f8e / 4

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

<a id="canonical-4bb5f6fda9aa625c155b83064f9bd120a8ffd1f0f57f2c5cc8285b76ebe184d9"></a>

<a id="canonical-89075feb651fc53cb52020d660be493fada3c189a987a1a1ac7d6fce17e77bcc"></a>

## namespace property — log_receiver / 5432735b3f8e / 5

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

<a id="canonical-fec36b6f52a30bedf0276045cb785300bbf21c60229518bd688286d4dea113d5"></a>

<a id="canonical-6918b41d346e2c594e34093d482dd0f4cc3971b75548eeaa0fc8fb18fa33cc05"></a>

## tenant property — log_receiver / 5432735b3f8e / 6

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

<a id="canonical-1a2364829a6f080479000253903dc52bc15f106c4672874098667c65c02cc4fe"></a>

## Next pages — log_receiver / 5432735b3f8e / 7

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-8e2259a9b2d220f4369c5133a7a4d2120a94e6dd0544e37a326c6fd85e416ca1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a51fd0696f871ebf5debb089922d0a4696cd97225a3ac6606e9f03e785f1423"></a>

## logs_streaming_disabled — logs_streaming_disabled / 9f95ab55b4db / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- logs_streaming_disabled

<a id="canonical-82a912f582da1406a56e5a0614190eaf3ab538b3ac531c200be7b2efa7d9d193"></a>

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

<a id="canonical-98907fffab8cbfce553658d08083677268225159f12cd8494a840bcff5a904f0"></a>

## Direct properties — logs_streaming_disabled / 9f95ab55b4db / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3a5eb3577ed1c8e779711d3776dc1f07da3d6e8530c4fee7237326210e69d7cf"></a>

## Next pages — logs_streaming_disabled / 9f95ab55b4db / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-befcbfaedf144bd307a72ed3a7799f5b3d479e264e8d1c71aec48e5aa0e3add9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-be676ecba83ea9777039df33471eb9119969094f83f130d879e42a6ff233bd35"></a>

## manual_routing — manual_routing / d8734e5748d4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- manual_routing

<a id="canonical-c864d3ad531ccb13cd7fc450b170b6ea6f1c9458d3bbd77e0a81ece6574449a6"></a>

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
manual_routing = {}
```

<a id="canonical-b0ea0379a601ec82858c68a64e16d21fa1354f55991276c95fb458047f5eb55a"></a>

## Direct properties — manual_routing / d8734e5748d4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-285c8d4d04b8962d16e8cec1a1cf0617e2301bc502bb2a819394c97a7be3e15e"></a>

## Next pages — manual_routing / d8734e5748d4 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-3b567624fc53d7fd5be1fbd116ca3550fbd6c572ba3c41a433a5459854ca051d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e7208029fa7c78848e80f575cf29856c726c59992097892355bf04a6b22ea916"></a>

## no_worker_nodes — no_worker_nodes / a1db0cafa580 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- no_worker_nodes

<a id="canonical-10f20b1b2571ba4f64da21213f026d249076dbe797ad9a63849f43bcad496801"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_worker\_nodes, nodes\_per\_az, total\_nodes; Default: no\_worker\_nodes\] Configuration
parameter for no worker nodes.

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

- [no_worker_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-10f20b1b2571ba4f64da21213f026d249076dbe797ad9a63849f43bcad496801)
- [nodes_per_az](resources--aws_vpc_site--reference--group-001.md#canonical-ee913788c85c8a86da740e494d8f1f0f8873ffbea734dc351e781ade07360004)
- [total_nodes](resources--aws_vpc_site--reference--group-001.md#canonical-7e1d45130ca6ad34af48bed23dd9606cab15fee04d7f9fefacf4afe53898b470)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_worker_nodes = {}
```

<a id="canonical-7eb35b54b562d832f92d3751214e451b6d2254e74eedfbf873f2ffdd4e4aaf99"></a>

## Direct properties — no_worker_nodes / a1db0cafa580 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-881f84be232ee3db0c2186842076f0f7916d1c4294b98cf4c287a4296f126dc7"></a>

## Next pages — no_worker_nodes / a1db0cafa580 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-acf7232f6dd85f603c15af2d430ac186fa48c2a311f068114194e344192d87de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-da1d275822da214f0252d2b962caeb3142b2bc0e597383858e8942143a578e3a"></a>

## offline_survivability_mode — offline_survivability_mode / d25e6a59115d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- offline_survivability_mode

<a id="canonical-7384ccf528ad5878419cf695fde4acb9020a041f00ad21fb507fa7355013db00"></a>

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

<a id="canonical-18aba78022bba05b49b8add8ce846460cb6d80b934afc8824ce12e02c5af34c0"></a>

## Direct properties — offline_survivability_mode / d25e6a59115d / 3

- [enable_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-0f3cc8854ab80ca6d5c13c05dfa80b399ed152117cfbb91d9db430e5fdbf7e4b): complete subsection reference.

- [no_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-75a18bad3b78fb91448deddf140d953fd07f79e6ba0634c9a11c4ad66fd4ff38): complete subsection reference.

<a id="canonical-fae8226652c1ceb8390ff862976d261e23b211f68aa233f19a5c36ff5d85a64b"></a>

## Next pages — offline_survivability_mode / d25e6a59115d / 4

- [offline_survivability_mode.enable_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-0f3cc8854ab80ca6d5c13c05dfa80b399ed152117cfbb91d9db430e5fdbf7e4b)
- [offline_survivability_mode.no_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-75a18bad3b78fb91448deddf140d953fd07f79e6ba0634c9a11c4ad66fd4ff38)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-0f3cc8854ab80ca6d5c13c05dfa80b399ed152117cfbb91d9db430e5fdbf7e4b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5b44fa59f7bb1f8d49df3b9a2dfc22e2f53edb5a479bd2b152f28eeaf72b970"></a>

## offline_survivability_mode.enable_offline_survivability_mode — offline_survivability_mode.enable_offline_survivability_mode / ed2aa6c2daa4 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-acf7232f6dd85f603c15af2d430ac186fa48c2a311f068114194e344192d87de)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-0bcacd0223f947a47a34190e9b4b5993111160abbf60280a2363d6132573ade2"></a>

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

<a id="canonical-6ddde33c65b9b3ff78549b272a2118d532e480f0a46fe3fc310a921c2e56248d"></a>

## Direct properties — offline_survivability_mode.enable_offline_survivability_mode / ed2aa6c2daa4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-85cd6bbe871db081c4f8ecf163432da89b472e4c95a692f8c35bacc6c15673b0"></a>

## Next pages — offline_survivability_mode.enable_offline_survivability_mode / ed2aa6c2daa4 / 4

- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-acf7232f6dd85f603c15af2d430ac186fa48c2a311f068114194e344192d87de)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-75a18bad3b78fb91448deddf140d953fd07f79e6ba0634c9a11c4ad66fd4ff38"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c1cce426b267da96f3c1123b9ee0967255a4614be6811637e58f967d268cdd2"></a>

## offline_survivability_mode.no_offline_survivability_mode — offline_survivability_mode.no_offline_survivability_mode / 19ccb5993c71 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-acf7232f6dd85f603c15af2d430ac186fa48c2a311f068114194e344192d87de)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-0701ab95e79b5872db237799df2b2e4c35fdc34c2d3ccc465af61bc2acf934ad"></a>

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

<a id="canonical-23069a13461fea8eb737cd3debb5c8867158548bdcff8e48f1b02f16f3c5f02b"></a>

## Direct properties — offline_survivability_mode.no_offline_survivability_mode / 19ccb5993c71 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-14d7b5416d1228aa5afe32e43c0a9cead12f8ca0a73d4dee69956fe7c0f30064"></a>

## Next pages — offline_survivability_mode.no_offline_survivability_mode / 19ccb5993c71 / 4

- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-acf7232f6dd85f603c15af2d430ac186fa48c2a311f068114194e344192d87de)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-d9fbbd4f678898f88afff40d99873b7424d6d2c4373892f5fd2ada2490f1581f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fb181c438b2b953a6d3fe123ec738fbe866702ac04178bf04dff5149fccc8d1b"></a>

## os — os / bbe01c4bd9e0 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- os

<a id="canonical-655ec6c3a5c03ca120af2c88da6d11ae517be8eb1e9c3342be4d6473300afa58"></a>

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

<a id="canonical-63c0cda294b1cbb48c5e308c99b8beaf61865361ebb33daf216e6555671dad05"></a>

## Direct properties — os / bbe01c4bd9e0 / 3

- [default_os_version](resources--aws_vpc_site--reference--group-004.md#canonical-3f5a33f13315f676f73f6d4de473d92aee23f45540e257a52ce4c05449a6b746): complete subsection reference.

<a id="canonical-6f49e0d9f0d0817455827ae9b44d8edfd6958cc80e811be5ebcd1f4fa6269c9b"></a>

<a id="canonical-177a20fb374f176d73a001c242b8cf0a5d8aca085c5c8beedcdbfa3fd82a7fb1"></a>

## operating_system_version property — os / bbe01c4bd9e0 / 4

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

<a id="canonical-a7620573340081632e0d2fc3028772e4814758ed327563763d85ee73517c0357"></a>

## Next pages — os / bbe01c4bd9e0 / 5

- [os.default_os_version](resources--aws_vpc_site--reference--group-004.md#canonical-3f5a33f13315f676f73f6d4de473d92aee23f45540e257a52ce4c05449a6b746)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-3f5a33f13315f676f73f6d4de473d92aee23f45540e257a52ce4c05449a6b746"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f7af3e992c705e21af7fddece0e0d6b2ce2beec92b608c21e74f78de9b0877e9"></a>

## os.default_os_version — os.default_os_version / 914b7194e9af / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [os](resources--aws_vpc_site--reference--group-004.md#canonical-d9fbbd4f678898f88afff40d99873b7424d6d2c4373892f5fd2ada2490f1581f)
- os.default_os_version

<a id="canonical-002a82370040155c6062f63027dbefda2bf7554afbe43c59348ef8fb00ff6fcc"></a>

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

<a id="canonical-5a31a81fac1eb2f1085d7ecd8e28f5c6a1179c40b43c899c181b7113279fa759"></a>

## Direct properties — os.default_os_version / 914b7194e9af / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f24a5bbaa72de5669524d4cf61828cfdf3c87ad37819709c541120cc241ac95"></a>

## Next pages — os.default_os_version / 914b7194e9af / 4

- [os](resources--aws_vpc_site--reference--group-004.md#canonical-d9fbbd4f678898f88afff40d99873b7424d6d2c4373892f5fd2ada2490f1581f)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-4fe779f0f64998d18b5a237682b2f74cf8640bded8e9f265be749b40b91d9bc1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-345a9ef49ba54e7246740fed30fbfcc4f603816b3a9b8cd6e226420d2f6f156f"></a>

## private_connectivity — private_connectivity / 6a6770af4065 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- private_connectivity

<a id="canonical-f209eb63cc8a411874e064df721409acd5caf5601b01bf4bc66a44a3062759ea"></a>

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

<a id="canonical-392332239aa1980b98c4b6126c13fe690146e9439d23f64773343223a98d2b05"></a>

## Direct properties — private_connectivity / 6a6770af4065 / 3

- [cloud_link](resources--aws_vpc_site--reference--group-004.md#canonical-689f11168d95882cb4c434ff4865ecaedfd4422f4f8a49fad22b1fa253ed5033): complete subsection reference.

- [inside](resources--aws_vpc_site--reference--group-004.md#canonical-257de297afa48e26e3e03154dd415da1c2f2374c2bb15a331a6ca3c881b3d18d): complete subsection reference.

- [outside](resources--aws_vpc_site--reference--group-004.md#canonical-090abdd95f3d2836bc537e6f8bead026e05b9b37002260697d1bd87a53752b90): complete subsection reference.

<a id="canonical-dc196eca4fbcd43b13ef5636ec86ea7e4dda5142c483c71ef3f50d07242719e0"></a>

## Next pages — private_connectivity / 6a6770af4065 / 4

- [private_connectivity.cloud_link](resources--aws_vpc_site--reference--group-004.md#canonical-689f11168d95882cb4c434ff4865ecaedfd4422f4f8a49fad22b1fa253ed5033)
- [private_connectivity.inside](resources--aws_vpc_site--reference--group-004.md#canonical-257de297afa48e26e3e03154dd415da1c2f2374c2bb15a331a6ca3c881b3d18d)
- [private_connectivity.outside](resources--aws_vpc_site--reference--group-004.md#canonical-090abdd95f3d2836bc537e6f8bead026e05b9b37002260697d1bd87a53752b90)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-689f11168d95882cb4c434ff4865ecaedfd4422f4f8a49fad22b1fa253ed5033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c9c76d26e7ca565d61ae118792b0e5d748cc8715be4093c3089373555516bc75"></a>

## private_connectivity.cloud_link — private_connectivity.cloud_link / f18e7e535edc / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-4fe779f0f64998d18b5a237682b2f74cf8640bded8e9f265be749b40b91d9bc1)
- private_connectivity.cloud_link

<a id="canonical-b663721f1cd9316618e077d17c442bb0a35a40404a907b1db466557691cace73"></a>

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

<a id="canonical-e7f529464502b14aa80d763f3bf3f608c25be40c934b8185abe014556c4ce62c"></a>

## Direct properties — private_connectivity.cloud_link / f18e7e535edc / 3

<a id="canonical-d7326ee3f16add41b9adaf007a6bb4e24913f678837934c416d58b7c6d0f1146"></a>

<a id="canonical-afe5b34b440b0cfe0389b8c3a1c99829556b13f379d16d0a9be2acccb8675e04"></a>

## name property — private_connectivity.cloud_link / f18e7e535edc / 4

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

<a id="canonical-15d477802a1c22a0f5c9041fece65489c53dd3a98f6ef3545e10d67c63a02607"></a>

<a id="canonical-a589702693049ff385e1a6ab8af4eda8d17b0c6e9eb69f2fc0ee298141a0e7a0"></a>

## namespace property — private_connectivity.cloud_link / f18e7e535edc / 5

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

<a id="canonical-691d236ee7dc35894c759c71987790792a85124eb4bf73cdb1797e8c7359288c"></a>

<a id="canonical-a4efe8274351f7aa106cc214c89f4a1b06ca1cef310fd069c4305c8d3e731e23"></a>

## tenant property — private_connectivity.cloud_link / f18e7e535edc / 6

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

<a id="canonical-df89b769f0d518d4470121fa73877b17bf52791222a226b8dad65d2660c1b66b"></a>

## Next pages — private_connectivity.cloud_link / f18e7e535edc / 7

- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-4fe779f0f64998d18b5a237682b2f74cf8640bded8e9f265be749b40b91d9bc1)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-257de297afa48e26e3e03154dd415da1c2f2374c2bb15a331a6ca3c881b3d18d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8d2341ea5b9bf4963b3fee80fcff3b08b004a3d629bf829d639f3659869c3de5"></a>

## private_connectivity.inside — private_connectivity.inside / 8fcdf53099d7 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-4fe779f0f64998d18b5a237682b2f74cf8640bded8e9f265be749b40b91d9bc1)
- private_connectivity.inside

<a id="canonical-9bfdc6ca54536065b47ec6de8108a6396cb9c894d72b341e53c649dc64a31d07"></a>

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

<a id="canonical-f0623e09bb207a13486051d70de9250e210416490cae216bcc72675ecb1b5ef6"></a>

## Direct properties — private_connectivity.inside / 8fcdf53099d7 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d758608893e229865bba4cc8e0d2d460744e48df2ebe931bdb23642d6b711e63"></a>

## Next pages — private_connectivity.inside / 8fcdf53099d7 / 4

- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-4fe779f0f64998d18b5a237682b2f74cf8640bded8e9f265be749b40b91d9bc1)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-090abdd95f3d2836bc537e6f8bead026e05b9b37002260697d1bd87a53752b90"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5719df516b9b157daa9bb6d662fc021a8c40a486da37df5e179de10bb59b55b7"></a>

## private_connectivity.outside — private_connectivity.outside / 23cd2fd06d9c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-4fe779f0f64998d18b5a237682b2f74cf8640bded8e9f265be749b40b91d9bc1)
- private_connectivity.outside

<a id="canonical-46495d930a8858ae3c2515c4952cddce574383b1d622203f6f1845ca377a25d0"></a>

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

<a id="canonical-c141c36a9aef9f9f2a2341119a44c39df39231cbb4b17be036dfd679427d2da1"></a>

## Direct properties — private_connectivity.outside / 23cd2fd06d9c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-46fa5e00e3aa5570a6fe0115a598f57df4e0fe9eb3ca0f08e7bcd2927dd764e0"></a>

## Next pages — private_connectivity.outside / 23cd2fd06d9c / 4

- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-4fe779f0f64998d18b5a237682b2f74cf8640bded8e9f265be749b40b91d9bc1)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-789f7d6e3d560f13cd6566c9347e630b0ad42ebf4f749e81c7ac27bccff98aa0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1dd797b894edad432cba03fc6812d10f35fc493f89e8435c84a300b2717d7e31"></a>

## sw — sw / 89480343328b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- sw

<a id="canonical-f265dad63b8d2fdae9a9fbba4db18760f919f55c4952590194693c507341afa0"></a>

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

<a id="canonical-059d6df76f5994aa805e148c282de3f9d482cd44047c30529b05a384fbb8740f"></a>

## Direct properties — sw / 89480343328b / 3

- [default_sw_version](resources--aws_vpc_site--reference--group-004.md#canonical-0ac767481d2019a605b794413a2733f85cd3c4d9f50437263c3d3cb4fa7dff7e): complete subsection reference.

<a id="canonical-3cf5fdd8fb094823637a106183cc53ab4ca19f12abcc8fb421fac72f4a45d063"></a>

<a id="canonical-0919975cc5d06798b35b25509a875c9fbcf5bf221200c15225eda8f1e9a1d1ab"></a>

## volterra_software_version property — sw / 89480343328b / 4

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

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

<a id="canonical-393fc9d17eb37f24b425b91f0599b03606c8d90a1525ac81bc92088c20de3685"></a>

## Next pages — sw / 89480343328b / 5

- [sw.default_sw_version](resources--aws_vpc_site--reference--group-004.md#canonical-0ac767481d2019a605b794413a2733f85cd3c4d9f50437263c3d3cb4fa7dff7e)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-0ac767481d2019a605b794413a2733f85cd3c4d9f50437263c3d3cb4fa7dff7e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1b450597830cd7ba1ca4c1075a4da2671219641132f0282ff62405e9c795f46b"></a>

## sw.default_sw_version — sw.default_sw_version / e4b273ab9f93 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [sw](resources--aws_vpc_site--reference--group-004.md#canonical-789f7d6e3d560f13cd6566c9347e630b0ad42ebf4f749e81c7ac27bccff98aa0)
- sw.default_sw_version

<a id="canonical-a96a7b1832b2e919ca11e006657c2afbd4a6f04963cb81c935cb4b9bfb52beb7"></a>

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
default_sw_version = {}
```

<a id="canonical-bf58b55d2b0f885985639e5160d6ef5f71ce1bef52f57f8c823f08ef65d295cf"></a>

## Direct properties — sw.default_sw_version / e4b273ab9f93 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-45b0ce006c819efd3dc7c424d66a9df4bad6c7f592e77dd1a8b45a3cdadf9aa6"></a>

## Next pages — sw.default_sw_version / e4b273ab9f93 / 4

- [sw](resources--aws_vpc_site--reference--group-004.md#canonical-789f7d6e3d560f13cd6566c9347e630b0ad42ebf4f749e81c7ac27bccff98aa0)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-e8515c5d0f67a8340f079b06dd5c6221e19fdcba0f0ec762690045b2b82602de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3bdab08bb502c0be4399c9db16c55c45867214d52a730aea905158229c85540e"></a>

## timeouts — timeouts / f5deaef63fc9 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- timeouts

<a id="canonical-ea4b02153fb0ce5ab671520380a1c243fad8654a7b55f64910d86b20c68cc803"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-e4ba94177365b717213585b8e87a801d5f92a81556a447d4c5cd3c47d6dce9f7"></a>

## Direct properties — timeouts / f5deaef63fc9 / 3

<a id="canonical-543e4915d9f1736a25fba5aa471847f04ab0a74057aa5ffc9ce35258ca2a8ce3"></a>

<a id="canonical-d34072eb95e833861fe1d900db1f57ca064e0db0a0af5d084bb487c35ad7fcd9"></a>

## create property — timeouts / f5deaef63fc9 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-092043c5570a6c8990495f1053f3095a4fd602dee01b365fd1f7faf220964635"></a>

<a id="canonical-c7d2c3b66bc83b4448d2f51c3db7cbdf4602c6aee3c9471951f43db506af0340"></a>

## delete property — timeouts / f5deaef63fc9 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-8d4de19d384ec2b914cfd2e47f7f690846fef51bb4e623ca4bb8c978631e11be"></a>

<a id="canonical-07afb2d97eb01c79f5a2b4a9a949271e9a3e09faa94a97230a00a6d3d58eede2"></a>

## read property — timeouts / f5deaef63fc9 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1ef485fb63bd324511a40fb639980e7959378c3d35d111f9e21aad99eeba9a97"></a>

<a id="canonical-7224941a3b0c12f2923f9a03698b749865a69dc9a6e9b111d5f38f5445c218b1"></a>

## update property — timeouts / f5deaef63fc9 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-e3677ae1e896f193a6720d7d832e578a23dc7cb3f3d4a90261b1481276ae635f"></a>

## Next pages — timeouts / f5deaef63fc9 / 8

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-881b53ef82a061e990534deda613d50bc5ae29caab9ded604e7c640ceeaa79a8"></a>

## voltstack_cluster — voltstack_cluster / cb57b62861e3 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- voltstack_cluster

<a id="canonical-4141c9dfecf713cd3dfd532024e078d75bda39e24b2abdbeace7b643dc73f1d6"></a>

Type: `"object"`. single nested block, Optional.

App Stack cluster of single interface AWS nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_certified_hw",
    "az_nodes"),
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
voltstack_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-bc85e8ab42b3b2183092367a560e5e5ce8261c46ab5a50e8aaa878ffa87321b5"></a>

## Direct properties — voltstack_cluster / cb57b62861e3 / 3

- [active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-f039a7d550c8b68cd87b6a81dfd35d8138d111529854edd6ecd8e1e3e4997fe5): complete subsection reference.

- [active_forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-285aef1561a7e1fbdc5feddcfe4dbb66f1bc5d1b1c85e44b009aa63eeede01b5): complete subsection reference.

- [active_network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-b5a9ab6f848a28490c2b6fe3ac83846e8b8f2641b9dc00a9678ad8333f10f8a4): complete subsection reference.

- [allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b): complete subsection reference.

<a id="canonical-0699156be774ac8da6e4ceba8b1ea3ffcc65836ec752ca4f5a4b3d59feb2588a"></a>

<a id="canonical-74998fbbea945e928e0665142ec92a36806a885eda5fc9b35aa0d28791d3714d"></a>

## aws_certified_hw property — voltstack_cluster / cb57b62861e3 / 4

Type: `"string"`. Optional.

\[Enum: aws-byol-voltstack-combo\] AWS Certified Hardware. Name for AWS certified hardware. The only
possible value is \`aws-byol-voltstack-combo\`.

Upstream description:

Name for AWS certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("aws-byol-voltstack-combo"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-voltstack-combo"
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
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-4f5a986fba68a58de02fe9f45a6aa52fc39903f10532528746bd0b1214796051): complete subsection reference.

- [dc_cluster_group](resources--aws_vpc_site--reference--group-004.md#canonical-1bc0b96797cd0ad2e75acf21f3174a7557595dad8052b57278b019c1f2f7a059): complete subsection reference.

- [default_storage](resources--aws_vpc_site--reference--group-004.md#canonical-9a39bc8845556c4b20b61042160e4b89b8dffdafd8fe71833f1263700f74f7aa): complete subsection reference.

- [forward_proxy_allow_all](resources--aws_vpc_site--reference--group-004.md#canonical-5e915d1a5dcba8c8868ce0bf0619335928eb8e28a13fc26fd6359e005f69ebb1): complete subsection reference.

- [global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-a3256201df629eddd4d97416a0a522dc7f5ffdb9feb2aa23e7559a58ce874f55): complete subsection reference.

- [k8s_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-42acdfe9767d853d8032ec5a385b380b28da930f26b46e0e7479964b8d553d3b): complete subsection reference.

- [no_dc_cluster_group](resources--aws_vpc_site--reference--group-004.md#canonical-617137e51ce90d1c4c0adf39c6e17957f864ba1df845dc2b4f661e4052034c9f): complete subsection reference.

- [no_forward_proxy](resources--aws_vpc_site--reference--group-005.md#canonical-b7b42396a1e761eb16b3cd8f83fdab8edfd5705105753a3e0bd41249b57fcdb5): complete subsection reference.

- [no_global_network](resources--aws_vpc_site--reference--group-005.md#canonical-a58d7f27b0b8113205cb12afb55a25ea07c0ef1bc3ef84ba25d6e48b91b12668): complete subsection reference.

- [no_k8s_cluster](resources--aws_vpc_site--reference--group-005.md#canonical-a5ee0767659ffdb3523324ed8d32c9a63f497190a575f168bf8e2c5b20c85f6d): complete subsection reference.

- [no_network_policy](resources--aws_vpc_site--reference--group-005.md#canonical-c3ca312f839a1b1675d8dd56ea64121c40fc467143cee769e110dfc78e9198c9): complete subsection reference.

- [no_outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-744ba0d76e77890668b4eb84789e562adf9ed608af79fdcdf0bd10d49cf337f9): complete subsection reference.

- [outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-e0ef08706068049f92d3b6002a5982c4231e0ba142e9b2592a0c245650411f65): complete subsection reference.

- [sm_connection_public_ip](resources--aws_vpc_site--reference--group-005.md#canonical-fcc2805c3451f2ead066fd48c33fb4f6068a53ac009c8b7db582a9f036e14436): complete subsection reference.

- [sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-005.md#canonical-aec150ff7417f849a375e48d63508ef61711baf607b41f5aa2d4756bb67c6abe): complete subsection reference.

- [storage_class_list](resources--aws_vpc_site--reference--group-005.md#canonical-f0360f04561206ce9769eff21b916b4a7e196174010d8714b4a42766ed88e865): complete subsection reference.

<a id="canonical-c15d4684e62d61453d0b10d0acb72abe93fc1573b9806590fbed81d1a12d2bf6"></a>

## Next pages — voltstack_cluster / cb57b62861e3 / 5

- [voltstack_cluster.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-f039a7d550c8b68cd87b6a81dfd35d8138d111529854edd6ecd8e1e3e4997fe5)
- [voltstack_cluster.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-285aef1561a7e1fbdc5feddcfe4dbb66f1bc5d1b1c85e44b009aa63eeede01b5)
- [voltstack_cluster.active_network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-b5a9ab6f848a28490c2b6fe3ac83846e8b8f2641b9dc00a9678ad8333f10f8a4)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b)
- [voltstack_cluster.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-4f5a986fba68a58de02fe9f45a6aa52fc39903f10532528746bd0b1214796051)
- [voltstack_cluster.dc_cluster_group](resources--aws_vpc_site--reference--group-004.md#canonical-1bc0b96797cd0ad2e75acf21f3174a7557595dad8052b57278b019c1f2f7a059)
- [voltstack_cluster.default_storage](resources--aws_vpc_site--reference--group-004.md#canonical-9a39bc8845556c4b20b61042160e4b89b8dffdafd8fe71833f1263700f74f7aa)
- [voltstack_cluster.forward_proxy_allow_all](resources--aws_vpc_site--reference--group-004.md#canonical-5e915d1a5dcba8c8868ce0bf0619335928eb8e28a13fc26fd6359e005f69ebb1)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-a3256201df629eddd4d97416a0a522dc7f5ffdb9feb2aa23e7559a58ce874f55)
- [voltstack_cluster.k8s_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-42acdfe9767d853d8032ec5a385b380b28da930f26b46e0e7479964b8d553d3b)
- [voltstack_cluster.no_dc_cluster_group](resources--aws_vpc_site--reference--group-004.md#canonical-617137e51ce90d1c4c0adf39c6e17957f864ba1df845dc2b4f661e4052034c9f)
- [voltstack_cluster.no_forward_proxy](resources--aws_vpc_site--reference--group-005.md#canonical-b7b42396a1e761eb16b3cd8f83fdab8edfd5705105753a3e0bd41249b57fcdb5)
- [voltstack_cluster.no_global_network](resources--aws_vpc_site--reference--group-005.md#canonical-a58d7f27b0b8113205cb12afb55a25ea07c0ef1bc3ef84ba25d6e48b91b12668)
- [voltstack_cluster.no_k8s_cluster](resources--aws_vpc_site--reference--group-005.md#canonical-a5ee0767659ffdb3523324ed8d32c9a63f497190a575f168bf8e2c5b20c85f6d)
- [voltstack_cluster.no_network_policy](resources--aws_vpc_site--reference--group-005.md#canonical-c3ca312f839a1b1675d8dd56ea64121c40fc467143cee769e110dfc78e9198c9)
- [voltstack_cluster.no_outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-744ba0d76e77890668b4eb84789e562adf9ed608af79fdcdf0bd10d49cf337f9)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-e0ef08706068049f92d3b6002a5982c4231e0ba142e9b2592a0c245650411f65)
- [voltstack_cluster.sm_connection_public_ip](resources--aws_vpc_site--reference--group-005.md#canonical-fcc2805c3451f2ead066fd48c33fb4f6068a53ac009c8b7db582a9f036e14436)
- [voltstack_cluster.sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-005.md#canonical-aec150ff7417f849a375e48d63508ef61711baf607b41f5aa2d4756bb67c6abe)
- [voltstack_cluster.storage_class_list](resources--aws_vpc_site--reference--group-005.md#canonical-f0360f04561206ce9769eff21b916b4a7e196174010d8714b4a42766ed88e865)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-f039a7d550c8b68cd87b6a81dfd35d8138d111529854edd6ecd8e1e3e4997fe5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4b86c47cddb98f37db18b325594286e74a698d8fdf59713dd19b66513e4c4f37"></a>

## voltstack_cluster.active_enhanced_firewall_policies — voltstack_cluster.active_enhanced_firewall_policies / 7e3d76601e58 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- voltstack_cluster.active_enhanced_firewall_policies

<a id="canonical-0611cf8d680bce3c22018f73ac6b5ee536e36f2669f47b464822c679a441c8a2"></a>

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

<a id="canonical-6e46e5f848f540ff8279555b8aecf9c62e612c05a8d95ed84b54c1b53173956f"></a>

## Direct properties — voltstack_cluster.active_enhanced_firewall_policies / 7e3d76601e58 / 3

- [enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-a32ddea5e057ad5efd0d692492861c926f623afc9948fe52cfc4cb28a1063e40): complete subsection reference.

<a id="canonical-f5f356cb8c6074aadcd968f622fbf7ff5a09227476ac8cf8f2b9d7c602b63dd7"></a>

## Next pages — voltstack_cluster.active_enhanced_firewall_policies / 7e3d76601e58 / 4

- [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-a32ddea5e057ad5efd0d692492861c926f623afc9948fe52cfc4cb28a1063e40)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-a32ddea5e057ad5efd0d692492861c926f623afc9948fe52cfc4cb28a1063e40"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6c8a7821707ca87092483dfc3f93f3332f36bcfc5ad8f6c2ec44209410b69c80"></a>

## voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / e43e747644d3 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-f039a7d550c8b68cd87b6a81dfd35d8138d111529854edd6ecd8e1e3e4997fe5)
- voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-8e202d3883782fb0a9ab00a823287ad400b094cbc5cc67bb126a7259d4287d51"></a>

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

<a id="canonical-4bdc5f760594f3814fd489269960b318aa733d5723b4166dc52f782bb4a1cdce"></a>

## Direct properties — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / e43e747644d3 / 3

<a id="canonical-fdfbcdce328259ff3f1edc81761e743715ae020b4b5298153da67ec1314ecfb8"></a>

<a id="canonical-33e51ebbb078ed705d2bab765e8d9c4c4f4715e7393512fdd67c9994968936bc"></a>

## name property — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / e43e747644d3 / 4

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

<a id="canonical-e34e4c468bfbd5ee913339286fa38b19ac469e46e72a2ce788f479548cf48c2b"></a>

<a id="canonical-35c379b85e37603e88f7227bf02be267646d39057207dc759544513b22dedaf4"></a>

## namespace property — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / e43e747644d3 / 5

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

<a id="canonical-e99589e99d0ae84ddc4e792d14412db0a792fadcab278c30a0c569d9ce33acbc"></a>

<a id="canonical-38ac9467a56a1e5912492c6c979bb1ffa50d6e7eaf0f87f11bbd4d717476b6ff"></a>

## tenant property — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / e43e747644d3 / 6

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

<a id="canonical-17ea5d3d22dd8eb65d5df42b098d3d0bb5bd993b65a42851e17615603bd55849"></a>

## Next pages — voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies / e43e747644d3 / 7

- [voltstack_cluster.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-f039a7d550c8b68cd87b6a81dfd35d8138d111529854edd6ecd8e1e3e4997fe5)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-285aef1561a7e1fbdc5feddcfe4dbb66f1bc5d1b1c85e44b009aa63eeede01b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c89622cfbdb6a43526e6a64d82887005ad3433c006383a4548d654f4321cd2df"></a>

## voltstack_cluster.active_forward_proxy_policies — voltstack_cluster.active_forward_proxy_policies / c9a91aaa8047 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- voltstack_cluster.active_forward_proxy_policies

<a id="canonical-92db8098b32adafeadca027df3310047bdec2750decb80ae5c7495f635be3984"></a>

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

<a id="canonical-e80d21cae490d3a13221c56271506eff5de9a9e2c93b271e4270d8f18f28c63d"></a>

## Direct properties — voltstack_cluster.active_forward_proxy_policies / c9a91aaa8047 / 3

- [forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-77bfe292d2d3732cd19d618227afb974952ef93cb99bd39a2052a731237a80ad): complete subsection reference.

<a id="canonical-774a39b948319041037c5ed76e7575d89de31861ec1e91d0e8bb73a9fa27b70c"></a>

## Next pages — voltstack_cluster.active_forward_proxy_policies / c9a91aaa8047 / 4

- [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-77bfe292d2d3732cd19d618227afb974952ef93cb99bd39a2052a731237a80ad)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-77bfe292d2d3732cd19d618227afb974952ef93cb99bd39a2052a731237a80ad"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d9347b6ffaa33c9d87d0cd1bcf300477ac43721e05351e5f19928a85ea471e33"></a>

## voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 43466814c051 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-285aef1561a7e1fbdc5feddcfe4dbb66f1bc5d1b1c85e44b009aa63eeede01b5)
- voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-d7944c56cd6179fae3aab4a04b1ef5549f6f689147d24685aeb138e04226fb38"></a>

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

<a id="canonical-93c64039b8cf0514abf8dcf8aef6cd455f969cdf2e2ed4879ed682abc47789f9"></a>

## Direct properties — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 43466814c051 / 3

<a id="canonical-a823a85fb83502e1e21cc1ca95528fadf7eea46250ec4f86ec9a2fc665346311"></a>

<a id="canonical-b2056695e575109940b7b6bd749ee9cf3001b05778a0601919098f1ab4a8dbff"></a>

## name property — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 43466814c051 / 4

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

<a id="canonical-7180b76d2f722f6b6b33ee631e84637aa05fc326dc1ada6abf79a8ca9e441427"></a>

<a id="canonical-0246e86d2a030975b7261581a27706f129b498daa0fd419f90b2ac59dd1bfdb4"></a>

## namespace property — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 43466814c051 / 5

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

<a id="canonical-8fc797ba07b560ac4d1e188a748bcadbee6562ebd4bf0e99f83ca64d5b38df98"></a>

<a id="canonical-c5da95fe0d3470faaf2e5775542977653d4a3f2e4e03bbf6f5d53a25e15d476f"></a>

## tenant property — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 43466814c051 / 6

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

<a id="canonical-23dc98fbb4a55752156b030576374df93699fe2c3bd2639374f57b701e1c5eb6"></a>

## Next pages — voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies / 43466814c051 / 7

- [voltstack_cluster.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-285aef1561a7e1fbdc5feddcfe4dbb66f1bc5d1b1c85e44b009aa63eeede01b5)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-b5a9ab6f848a28490c2b6fe3ac83846e8b8f2641b9dc00a9678ad8333f10f8a4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fae5381728d01d33c757688784f16d4897db9a27b53fad6558782fdf2c21447a"></a>

## voltstack_cluster.active_network_policies — voltstack_cluster.active_network_policies / 2eeced4ef4fc / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- voltstack_cluster.active_network_policies

<a id="canonical-59e346f6bb4ab3e23ce74cef0c0212489b55e2605896323b109764b980219085"></a>

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

<a id="canonical-d9feb962f3315120c766d999f6f857a5dea7040b319aa120b931dbfd45fe8b40"></a>

## Direct properties — voltstack_cluster.active_network_policies / 2eeced4ef4fc / 3

- [network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-7b8df37a0ad2024a81f63dabac9643067bbf7642ff1f5b04885e323cfefaa196): complete subsection reference.

<a id="canonical-ed3e2f6ae4ff72483f3efc527f2c2cc91c9fc5a82268f115dbefff65f86c1aba"></a>

## Next pages — voltstack_cluster.active_network_policies / 2eeced4ef4fc / 4

- [voltstack_cluster.active_network_policies.network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-7b8df37a0ad2024a81f63dabac9643067bbf7642ff1f5b04885e323cfefaa196)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-7b8df37a0ad2024a81f63dabac9643067bbf7642ff1f5b04885e323cfefaa196"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2ad730e9273585d7ffee42747e8e1614a1d2e8a83130a11d913bc6b6068aa92"></a>

## voltstack_cluster.active_network_policies.network_policies — voltstack_cluster.active_network_policies.network_policies / 43866a438d09 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.active_network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-b5a9ab6f848a28490c2b6fe3ac83846e8b8f2641b9dc00a9678ad8333f10f8a4)
- voltstack_cluster.active_network_policies.network_policies

<a id="canonical-29a3dabd22a119f5c1f8080de6582a07d7a09c883fac7aeea4b95138e0d63ae1"></a>

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

<a id="canonical-8846c648399b2ad6719bb2ac307a6506819f1a7c04d55120223235fe3c1601c2"></a>

## Direct properties — voltstack_cluster.active_network_policies.network_policies / 43866a438d09 / 3

<a id="canonical-c0d0dc70dfa599b771250759848a5342be7bc8202596b43cd05298d1c6ed3701"></a>

<a id="canonical-0d07cfe73b54bbf9a3b5bfd85b1573dfe62424192a60e14a4c385f57340950a5"></a>

## name property — voltstack_cluster.active_network_policies.network_policies / 43866a438d09 / 4

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

<a id="canonical-1e854af2e2d07dcd695682cb15505423c4b40aa1eae7b4fe9ae44db31421243a"></a>

<a id="canonical-2ffe10363eca1e6e2bd4af5d2c4d30a658f6d726bddb46061a28bc3c27bec4cc"></a>

## namespace property — voltstack_cluster.active_network_policies.network_policies / 43866a438d09 / 5

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

<a id="canonical-efbccace5df3811cd703efee2ecc344e7aaa6669fe90a11bf66c598a76118b01"></a>

<a id="canonical-21e05520a991646bba5e0925e5f5c8bea0744bf1ed9c19e1f5fb322b01c85d5d"></a>

## tenant property — voltstack_cluster.active_network_policies.network_policies / 43866a438d09 / 6

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

<a id="canonical-59a9d360de5ae2e9fb1d3f244cb646132cafa1c26c0c208d3673838afa1d288c"></a>

## Next pages — voltstack_cluster.active_network_policies.network_policies / 43866a438d09 / 7

- [voltstack_cluster.active_network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-b5a9ab6f848a28490c2b6fe3ac83846e8b8f2641b9dc00a9678ad8333f10f8a4)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d599f2736d5d5489e16d981a8b5dc30748614ba5ff2fb64721e5c97a86ece4bc"></a>

## voltstack_cluster.allowed_vip_port — voltstack_cluster.allowed_vip_port / b769f7029446 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- voltstack_cluster.allowed_vip_port

<a id="canonical-8baaa6e298805ac2e9913aa390ca9a66d1a5e05a505a498662e93d09ab3f9a25"></a>

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

<a id="canonical-565ef22f81294b5785e880f64b105155babd523ab34e943972d242192cdd0776"></a>

## Direct properties — voltstack_cluster.allowed_vip_port / b769f7029446 / 3

- [custom_ports](resources--aws_vpc_site--reference--group-004.md#canonical-bbe6526a79ca3a0353cb3e6b08dc3fbab9dd810a679c062abb7c652581495ce4): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-91fc2633f874bf7372cb066c08713c2afb629fc9420687e2c405de49ddb248ff): complete subsection reference.

- [use_http_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-ebd000107392e6a4e87b20ebde6447a72da0cb7dd9d1344b249624bb5fd1b9b3): complete subsection reference.

- [use_http_port](resources--aws_vpc_site--reference--group-004.md#canonical-583d3d8cb971506eeab09f34f808738d75b1fd89e7ad212402f06d874734d7ef): complete subsection reference.

- [use_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-6c762839d61b63744b4fb1eb16e697c9a544819126f600a2e16d7c4481f5a2ab): complete subsection reference.

<a id="canonical-77272765a9f01f291cff777af7c7016d14df11a39fcc966b816b4d674d63cb57"></a>

## Next pages — voltstack_cluster.allowed_vip_port / b769f7029446 / 4

- [voltstack_cluster.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-004.md#canonical-bbe6526a79ca3a0353cb3e6b08dc3fbab9dd810a679c062abb7c652581495ce4)
- [voltstack_cluster.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-91fc2633f874bf7372cb066c08713c2afb629fc9420687e2c405de49ddb248ff)
- [voltstack_cluster.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-ebd000107392e6a4e87b20ebde6447a72da0cb7dd9d1344b249624bb5fd1b9b3)
- [voltstack_cluster.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-004.md#canonical-583d3d8cb971506eeab09f34f808738d75b1fd89e7ad212402f06d874734d7ef)
- [voltstack_cluster.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-6c762839d61b63744b4fb1eb16e697c9a544819126f600a2e16d7c4481f5a2ab)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-bbe6526a79ca3a0353cb3e6b08dc3fbab9dd810a679c062abb7c652581495ce4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e063e947eadd198da1f33ea5e0709c6c07581bd44e405c2955e25c98fef084b"></a>

## voltstack_cluster.allowed_vip_port.custom_ports — voltstack_cluster.allowed_vip_port.custom_ports / b36a9ec31d16 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b)
- voltstack_cluster.allowed_vip_port.custom_ports

<a id="canonical-80b19c12a55f074a30654feac97bb202a0d484efb8b59cb05540b18d3080cb06"></a>

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

<a id="canonical-8b98bf17a424c5021e93bf1eea890e1a034c6921d2ef7bbdcc34bb35580e4781"></a>

## Direct properties — voltstack_cluster.allowed_vip_port.custom_ports / b36a9ec31d16 / 3

<a id="canonical-c867a76d6a1b711150f2d34184a7ec4f0e2dd869a718fcc99847b42dbbbd08fa"></a>

<a id="canonical-5a15247458c794423b92e3d43b18530f86d35b4ed9e38f10d1c06e4c406e5fc8"></a>

## port_ranges property — voltstack_cluster.allowed_vip_port.custom_ports / b36a9ec31d16 / 4

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

<a id="canonical-3b6c89dc531e2ab3d6d89949a2200e2fc3dd0aa370713ec031d3f690bdf32d63"></a>

## Next pages — voltstack_cluster.allowed_vip_port.custom_ports / b36a9ec31d16 / 5

- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-91fc2633f874bf7372cb066c08713c2afb629fc9420687e2c405de49ddb248ff"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bf3b90e555c3b99628903d77b1e59613a1c6217152fcf65a501f16a62b8b6c47"></a>

## voltstack_cluster.allowed_vip_port.disable_allowed_vip_port — voltstack_cluster.allowed_vip_port.disable_allowed_vip_port / 2ab30257b767 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b)
- voltstack_cluster.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-b9831c905909a3e67fac2d00ffdf089325f90b8a63e97497621b48a77bf4e8e4"></a>

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

<a id="canonical-ab56a733fa7030a179d5d582fc7a854cbc9a05cc4a5f07c85e0d3a3990f36932"></a>

## Direct properties — voltstack_cluster.allowed_vip_port.disable_allowed_vip_port / 2ab30257b767 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8bfef8d87c0a61b5a26c37b5b64b637583d1fa2ac11ad958b6baf18636d12475"></a>

## Next pages — voltstack_cluster.allowed_vip_port.disable_allowed_vip_port / 2ab30257b767 / 4

- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-ebd000107392e6a4e87b20ebde6447a72da0cb7dd9d1344b249624bb5fd1b9b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4ac02f1b814ee03b67ebb520cbc974757c2230bc4173b3315a7292649d763292"></a>

## voltstack_cluster.allowed_vip_port.use_http_https_port — voltstack_cluster.allowed_vip_port.use_http_https_port / 8b9d6ec497c6 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b)
- voltstack_cluster.allowed_vip_port.use_http_https_port

<a id="canonical-3129a6cb9a91a78a6e4668715904eed63aef1bb6092f0c578e549a44fc30e4bd"></a>

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

<a id="canonical-a16930fb166e6093efd72e6a18249fcfcc65718616d1a083ebf75ed5f182cb02"></a>

## Direct properties — voltstack_cluster.allowed_vip_port.use_http_https_port / 8b9d6ec497c6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c78f057e4c9722decd636149aebc9f4ba0b4f905355963197b79ec81a836ff6f"></a>

## Next pages — voltstack_cluster.allowed_vip_port.use_http_https_port / 8b9d6ec497c6 / 4

- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-583d3d8cb971506eeab09f34f808738d75b1fd89e7ad212402f06d874734d7ef"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c578771f69fd199c377fc20a7cc8affbfcc10f825312ac9a2452365a99919e72"></a>

## voltstack_cluster.allowed_vip_port.use_http_port — voltstack_cluster.allowed_vip_port.use_http_port / 40f0632a001b / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b)
- voltstack_cluster.allowed_vip_port.use_http_port

<a id="canonical-bca41cf6e80923903a0771b4cd34fbc4dbc8ebc0aecbb53b9007c40ff1c3dfde"></a>

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

<a id="canonical-4216f7c9245bdda359c89ce6bbb4f786e0c057b236ba94244d6eaa47655595aa"></a>

## Direct properties — voltstack_cluster.allowed_vip_port.use_http_port / 40f0632a001b / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0bab2b3681547c9ad00fefcbc60961275f802d4f1089d95aaaea4925a8102113"></a>

## Next pages — voltstack_cluster.allowed_vip_port.use_http_port / 40f0632a001b / 4

- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-6c762839d61b63744b4fb1eb16e697c9a544819126f600a2e16d7c4481f5a2ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-823c31c8c69687f5c7fb1ba85d31532edc2e2fc814c7816d2b32f57ff99be5c0"></a>

## voltstack_cluster.allowed_vip_port.use_https_port — voltstack_cluster.allowed_vip_port.use_https_port / 99cd514ab87c / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b)
- voltstack_cluster.allowed_vip_port.use_https_port

<a id="canonical-bfab9db0e30066753d3ac21482202fb92e91317196e162d4aa2112a2b58e823c"></a>

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
use_https_port = {}
```

<a id="canonical-3dc7712a9ff2129668485ce51e6e398b18ad2d7fd96543f3419c31b6cd1e5d3f"></a>

## Direct properties — voltstack_cluster.allowed_vip_port.use_https_port / 99cd514ab87c / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-833b438a8f27fca638709662e489eb6d3b79fae2645d067a42317c05f253bfe1"></a>

## Next pages — voltstack_cluster.allowed_vip_port.use_https_port / 99cd514ab87c / 4

- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-1fbe740faf3ac262d8259ed53e943686e35a078c0aa867a309d5bc1f358eb20b)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-4f5a986fba68a58de02fe9f45a6aa52fc39903f10532528746bd0b1214796051"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6913192a377958f18ef6549a9152b0407ebfff0c28bd389072945b3f57fd8fd0"></a>

## voltstack_cluster.az_nodes — voltstack_cluster.az_nodes / 2b6f9662a2f7 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- voltstack_cluster.az_nodes

<a id="canonical-98f5e24803464bb6b06272b41e718daf0f212192beaafc31f6ae684fa3f831a6"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name")}
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

<a id="canonical-ad52a28b4b09540bd075df344a3821f57fa282ffb5f62b08e01f26564fdbae47"></a>

## Direct properties — voltstack_cluster.az_nodes / 2b6f9662a2f7 / 3

<a id="canonical-3f3c202c6f858536cac3b0188e06f0d0ac7c748b0f036a160c14d1cdc89317c5"></a>

<a id="canonical-1aea28858157b04e30e2809acbbe5f34525992eca79f3b5ad9d50fe6df34203b"></a>

## aws_az_name property — voltstack_cluster.az_nodes / 2b6f9662a2f7 / 4

Type: `"string"`. Optional.

AWS availability zone, must be consistent with the selected AWS region.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-1cdd2b1aab382de1b4853e70ff361ee0112d759b0f3032aa177dd897e1c07583): complete subsection reference.

<a id="canonical-a2768d241ac7eeab3b424b0cf2707abef177504c8733a9032a125c91a3a27cab"></a>

## Next pages — voltstack_cluster.az_nodes / 2b6f9662a2f7 / 5

- [voltstack_cluster.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-1cdd2b1aab382de1b4853e70ff361ee0112d759b0f3032aa177dd897e1c07583)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-1cdd2b1aab382de1b4853e70ff361ee0112d759b0f3032aa177dd897e1c07583"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-50c89071bd1abf957c853a544207ff1dde2fcbd06c3181f4426b2c3177ed752d"></a>

## voltstack_cluster.az_nodes.local_subnet — voltstack_cluster.az_nodes.local_subnet / f7dfdb814d54 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-4f5a986fba68a58de02fe9f45a6aa52fc39903f10532528746bd0b1214796051)
- voltstack_cluster.az_nodes.local_subnet

<a id="canonical-4ce9f218fae841feb3ef8a430a4cdf64f0162bab1f4e90734fe353e336513a0e"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for local subnet.

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
local_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-cc69c3fe2ab6e5fb16a3afe35133569770a464f783554f7a59ad9ea5bad0dfd7"></a>

## Direct properties — voltstack_cluster.az_nodes.local_subnet / f7dfdb814d54 / 3

<a id="canonical-7a2692f991efbd0ae1f84d73c7ecfcc9c7d53ca768c74797db018acf8b4f18cc"></a>

<a id="canonical-a78854343ad3ef0cf26f547a3d69349285032ed5d8e5ae0bc8de2a5f60d1abd3"></a>

## existing_subnet_id property — voltstack_cluster.az_nodes.local_subnet / f7dfdb814d54 / 4

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

- [subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-ba417654c0bee9bfeee8751c814a64d0dceb350529423f21b35cd1afda9c0719): complete subsection reference.

<a id="canonical-d3dfbb6ae2c4abb3b4215faf30df78b12aef8d0054a56b15c4e6030cde0a5f50"></a>

## Next pages — voltstack_cluster.az_nodes.local_subnet / f7dfdb814d54 / 5

- [voltstack_cluster.az_nodes.local_subnet.subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-ba417654c0bee9bfeee8751c814a64d0dceb350529423f21b35cd1afda9c0719)
- [voltstack_cluster.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-4f5a986fba68a58de02fe9f45a6aa52fc39903f10532528746bd0b1214796051)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-ba417654c0bee9bfeee8751c814a64d0dceb350529423f21b35cd1afda9c0719"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-793448723fcaf70a53def15a20b2be43a2daaab5e7e7f9180f3d0c030a540fbd"></a>

## voltstack_cluster.az_nodes.local_subnet.subnet_param — voltstack_cluster.az_nodes.local_subnet.subnet_param / 57bcfeab53ce / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-4f5a986fba68a58de02fe9f45a6aa52fc39903f10532528746bd0b1214796051)
- [voltstack_cluster.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-1cdd2b1aab382de1b4853e70ff361ee0112d759b0f3032aa177dd897e1c07583)
- voltstack_cluster.az_nodes.local_subnet.subnet_param

<a id="canonical-d5311464d37cb019cb48b18e090590765bfe6e41b604f226b0f8740ba0a393bc"></a>

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

<a id="canonical-24b2b55b21eea49b095fbd341385e4e10301aa0ee3ac15b1f9723de3b112ab36"></a>

## Direct properties — voltstack_cluster.az_nodes.local_subnet.subnet_param / 57bcfeab53ce / 3

<a id="canonical-e61dab5b02b7d137254322b8ea823b63befc9337090a87e9ce42a553d58fe550"></a>

<a id="canonical-57313dbc436f330fcb845b4299716e3a8022159699b42e11848a0709f55a63ed"></a>

## ipv4 property — voltstack_cluster.az_nodes.local_subnet.subnet_param / 57bcfeab53ce / 4

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

<a id="canonical-d3179b3f08d1248411a5b8524e929d84ba51a6021a43a09319baf4da3b22252a"></a>

## Next pages — voltstack_cluster.az_nodes.local_subnet.subnet_param / 57bcfeab53ce / 5

- [voltstack_cluster.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-1cdd2b1aab382de1b4853e70ff361ee0112d759b0f3032aa177dd897e1c07583)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-1bc0b96797cd0ad2e75acf21f3174a7557595dad8052b57278b019c1f2f7a059"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5aa533c97aeab42b2d1dd7127d5718a9519fb876be55430f233ae98a34c652e0"></a>

## voltstack_cluster.dc_cluster_group — voltstack_cluster.dc_cluster_group / 1a0219611695 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- voltstack_cluster.dc_cluster_group

<a id="canonical-4993b4df2ccabd0a9775cd25d8c8cc0ae8538a6556798638b0430a0fd3813bc1"></a>

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

<a id="canonical-825bfd8eae2710319176161bd7de492ce6e8564b7ec76ea767cf9e2d9ddbcb6b"></a>

## Direct properties — voltstack_cluster.dc_cluster_group / 1a0219611695 / 3

<a id="canonical-70c5b447464b1d83da51df9912921d7edf048a80634f0dccea2f0045482cde42"></a>

<a id="canonical-7e5a2beed0c7aefd98de7616ae4a57ba2ef36872dc89d6f956d834c9d7f27f8f"></a>

## name property — voltstack_cluster.dc_cluster_group / 1a0219611695 / 4

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

<a id="canonical-37f05279794cc6281003303ceed13fc6b50911e9855deb686306cd30a8f47828"></a>

<a id="canonical-82625d2385ff8a992be2d34f9cbbf72aee7f4702e016d4cfcb0cabe45dd3d579"></a>

## namespace property — voltstack_cluster.dc_cluster_group / 1a0219611695 / 5

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

<a id="canonical-c1bd4fd16a3cab0fce104a16eb6b78a11435a04142c8c8db776fb72945f2ccf2"></a>

<a id="canonical-4213f5e0a3c6a654c6708dd10be480bba37b47072b5d3e2f9b4eca68aab47630"></a>

## tenant property — voltstack_cluster.dc_cluster_group / 1a0219611695 / 6

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

<a id="canonical-3ed31e23dd73d10153b2c4c687560b9886371dfe964d480106e3de0703d286c7"></a>

## Next pages — voltstack_cluster.dc_cluster_group / 1a0219611695 / 7

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-9a39bc8845556c4b20b61042160e4b89b8dffdafd8fe71833f1263700f74f7aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f8c34459ef89e7d518ee27bfdd61183f2caf3080b2b79ec8b268f6b68d08c3a"></a>

## voltstack_cluster.default_storage — voltstack_cluster.default_storage / 704dc925a3e6 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- voltstack_cluster.default_storage

<a id="canonical-8eba3fe8e1f69e8b585e01ddf342c87e4f48e6bf0ba6f16e5caea36232a7ea54"></a>

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

<a id="canonical-6f690f30299fe540382815753776cb369b5b25d943246f6a24285480f6facb22"></a>

## Direct properties — voltstack_cluster.default_storage / 704dc925a3e6 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d4184fd7fe8e29c2a28ea18da53bb04c3fd4d398e5cbd84a5750b4be376b4ece"></a>

## Next pages — voltstack_cluster.default_storage / 704dc925a3e6 / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-5e915d1a5dcba8c8868ce0bf0619335928eb8e28a13fc26fd6359e005f69ebb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dee43ef06a9995babb3050b89fd09b710df10e0fe87253c51c8e6eb3f641e713"></a>

## voltstack_cluster.forward_proxy_allow_all — voltstack_cluster.forward_proxy_allow_all / 6c8cd5e4aaba / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- voltstack_cluster.forward_proxy_allow_all

<a id="canonical-95f860311be77962e580d0fe0c55610dcc97d701306fe3561392cb398d17d6ac"></a>

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

<a id="canonical-c30714aa1f52fbd523f027723ece36bebb34f846c6db234ca12d7c50dfb50b2d"></a>

## Direct properties — voltstack_cluster.forward_proxy_allow_all / 6c8cd5e4aaba / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-89bf413de3120d972cc6634ff4ea02399e6e14e526b137f3ec114b7116c43b98"></a>

## Next pages — voltstack_cluster.forward_proxy_allow_all / 6c8cd5e4aaba / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-a3256201df629eddd4d97416a0a522dc7f5ffdb9feb2aa23e7559a58ce874f55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6e4290ca6ab0637b8bffab2505d111590a055b5b2c13235e7623aaf80424d17c"></a>

## voltstack_cluster.global_network_list — voltstack_cluster.global_network_list / 037841e33dc2 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- voltstack_cluster.global_network_list

<a id="canonical-e1d4954dac5d28ec4cec557541679b4e545f4c2838c6d99511ef3d65e1fa0114"></a>

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

<a id="canonical-f73f443ba3fcf3ea34c65b1bf2d28e8d37d2d8127f0eb472e0b918fcd7848570"></a>

## Direct properties — voltstack_cluster.global_network_list / 037841e33dc2 / 3

- [global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-a842d8a040cf8fbb2b5c126f4706642001b38022146e6049af54dd240bcc5275): complete subsection reference.

<a id="canonical-b642adce3ce097c5357b58c0a9d786ea66c7d49e01edf69b649045ff86bcb858"></a>

## Next pages — voltstack_cluster.global_network_list / 037841e33dc2 / 4

- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-a842d8a040cf8fbb2b5c126f4706642001b38022146e6049af54dd240bcc5275)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-a842d8a040cf8fbb2b5c126f4706642001b38022146e6049af54dd240bcc5275"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8357540c77437c9e1741286990aa0f8a3394383a6782c42d46941f11339286dd"></a>

## voltstack_cluster.global_network_list.global_network_connections — voltstack_cluster.global_network_list.global_network_connections / 574e5fd04f6d / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-a3256201df629eddd4d97416a0a522dc7f5ffdb9feb2aa23e7559a58ce874f55)
- voltstack_cluster.global_network_list.global_network_connections

<a id="canonical-765cccf3b3e176562e20b796cb20d633b16dc0527fd6da32cdbcb2e4c01a8d4e"></a>

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

<a id="canonical-a77e08348e77add8ad68d02573a579838f9caa63a1ae60566633eba4ec880b1a"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections / 574e5fd04f6d / 3

- [sli_to_global_dr](resources--aws_vpc_site--reference--group-004.md#canonical-3413ffa239937bc116655ed5067596b2260713ad48d9f781dc38bafef821df30): complete subsection reference.

- [slo_to_global_dr](resources--aws_vpc_site--reference--group-004.md#canonical-0f53dd34818b19d60b9a119e01c1a06fd3d4c92a1fb02b7caa10ec84131d3af0): complete subsection reference.

<a id="canonical-ba03a7b13a991c49fc9b83427711429efdc3b08765f152a8a57fedcb7f63b451"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections / 574e5fd04f6d / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-004.md#canonical-3413ffa239937bc116655ed5067596b2260713ad48d9f781dc38bafef821df30)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-004.md#canonical-0f53dd34818b19d60b9a119e01c1a06fd3d4c92a1fb02b7caa10ec84131d3af0)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-a3256201df629eddd4d97416a0a522dc7f5ffdb9feb2aa23e7559a58ce874f55)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-3413ffa239937bc116655ed5067596b2260713ad48d9f781dc38bafef821df30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a0024b72b9eadb68e13f363598decc083ac2395c285845cefb005248ebebc04"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 5f5cb67ea118 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-a3256201df629eddd4d97416a0a522dc7f5ffdb9feb2aa23e7559a58ce874f55)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-a842d8a040cf8fbb2b5c126f4706642001b38022146e6049af54dd240bcc5275)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-d3bc23e6e873f20ab21d23a9dfe2264d1399c9e599f6e3b6b8a9b6d79c878022"></a>

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

<a id="canonical-9b9c33907ce3a84a2fda0f39153679cd118214727385459cc2fedab734a666e9"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 5f5cb67ea118 / 3

- [global_vn](resources--aws_vpc_site--reference--group-004.md#canonical-1a6de99832042c33f77e33c76b872527f14617c2221edbc21484693692ee3b85): complete subsection reference.

<a id="canonical-c0038fcb3232786b0501060199b7fd8d367c7750df4ee3fc4c5384b79f190a9b"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 5f5cb67ea118 / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-004.md#canonical-1a6de99832042c33f77e33c76b872527f14617c2221edbc21484693692ee3b85)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-a842d8a040cf8fbb2b5c126f4706642001b38022146e6049af54dd240bcc5275)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-1a6de99832042c33f77e33c76b872527f14617c2221edbc21484693692ee3b85"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3c5c9075ff056a5a357e45f0907d9ea6a1b78f52cd5ac6730ba818c5f84fc6bb"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 8953d0e24c68 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-a3256201df629eddd4d97416a0a522dc7f5ffdb9feb2aa23e7559a58ce874f55)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-a842d8a040cf8fbb2b5c126f4706642001b38022146e6049af54dd240bcc5275)
- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-004.md#canonical-3413ffa239937bc116655ed5067596b2260713ad48d9f781dc38bafef821df30)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-f44b85c33882d89c530d15f1cb07e6d343948dd113d69cf8f02ce640311e4a3d"></a>

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

<a id="canonical-6a906194287e42c05f0dcf1579bdf7804ce40563aeb728b62d28feacf068a0cd"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 8953d0e24c68 / 3

<a id="canonical-a1e0dfd13bc4015a55507630b29898248784fb14786c5af37d35f9990a8149f7"></a>

<a id="canonical-5ef5cb23999389288e2997cf85903dc112be87900861a658caa2c5495b81117c"></a>

## name property — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 8953d0e24c68 / 4

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

<a id="canonical-c442d3e50f58ae2eda211d50746acdbaf3c3fe5a56db520bc9d803f9f2997276"></a>

<a id="canonical-ef83fea9d86c4b7c0f1279d93f3e86b9aececa140209463fbcd67b47dbe8854f"></a>

## namespace property — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 8953d0e24c68 / 5

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

<a id="canonical-70ee482d803c4e8f23b28e0c49dfcb0bb1681ae9b9c0206aa9674bd8b5cabded"></a>

<a id="canonical-97fdd483e82814c972f671c4fcab6231a34db41d120b5e143aa8a4ef4ba4f994"></a>

## tenant property — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 8953d0e24c68 / 6

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

<a id="canonical-f5a8f74b8ab97e5c20e00cb3d26d0b2df7646e44f3785ddb1db20d775d456d6a"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.sli_to_global_d / 8953d0e24c68 / 7

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-004.md#canonical-3413ffa239937bc116655ed5067596b2260713ad48d9f781dc38bafef821df30)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-0f53dd34818b19d60b9a119e01c1a06fd3d4c92a1fb02b7caa10ec84131d3af0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-33f86dba822878282a05af99e5eeffaf25c47333930cb8961e5ce18f6f1beb33"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / e21712eb9618 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-a3256201df629eddd4d97416a0a522dc7f5ffdb9feb2aa23e7559a58ce874f55)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-a842d8a040cf8fbb2b5c126f4706642001b38022146e6049af54dd240bcc5275)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-49901eec3ff4f801d14a0985b83784970df1fe075eb078d388220387fc6a62d5"></a>

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

<a id="canonical-7e0a11d2cd6ef406fc2b1af7297d1e249d314e0688ce5f1c94f29b731de51098"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / e21712eb9618 / 3

- [global_vn](resources--aws_vpc_site--reference--group-004.md#canonical-e483d40a0d4e9483149e649f7bc29bcb40d2c37a181bc6951e02a59e0d77d20c): complete subsection reference.

<a id="canonical-45732907ffdcf087d06c70f08cd741ea373cae23e32c57b6f10143c259cfadd2"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / e21712eb9618 / 4

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-004.md#canonical-e483d40a0d4e9483149e649f7bc29bcb40d2c37a181bc6951e02a59e0d77d20c)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-a842d8a040cf8fbb2b5c126f4706642001b38022146e6049af54dd240bcc5275)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-e483d40a0d4e9483149e649f7bc29bcb40d2c37a181bc6951e02a59e0d77d20c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-204e2cd0f89211fc6519e67536690b1e1fb7197ea19e58db67fa1560b22a56a2"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / fdce0951c5dc / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-a3256201df629eddd4d97416a0a522dc7f5ffdb9feb2aa23e7559a58ce874f55)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-a842d8a040cf8fbb2b5c126f4706642001b38022146e6049af54dd240bcc5275)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-004.md#canonical-0f53dd34818b19d60b9a119e01c1a06fd3d4c92a1fb02b7caa10ec84131d3af0)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-10607dbdf3c7916e6e34ecec605cfca0b6bd37da6add237b73c9d2b5aaf4e0ea"></a>

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

<a id="canonical-c1006b5f08ce917ae169a8965799521d3f558cd57610e2ce9a6619437e88d83c"></a>

## Direct properties — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / fdce0951c5dc / 3

<a id="canonical-c12fca647de0420e92e7e3f64f55c987a35aff638cda4ed93f57bd89d440a0c6"></a>

<a id="canonical-729f5dd3bc93256412387d7a6d699b36d5a535c35dd270a31ce7f9d330d0dc21"></a>

## name property — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / fdce0951c5dc / 4

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

<a id="canonical-0447d4c1e12fc0f6106bf860682b657a14bd1767a97d2fa7c80e97228828e404"></a>

<a id="canonical-a3f0d91b7aeaf774538406919774d2ad769814945eb32fb9f7a0f9be34a3b755"></a>

## namespace property — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / fdce0951c5dc / 5

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

<a id="canonical-26c4db3943e9ba1ff0ccc2f29b0184c08b412c69c21e7c110afaef8518a4f1d8"></a>

<a id="canonical-1a4e9c478db62bcea77bba6dc32b0b22dff25f6788349eedc38914901978fcfa"></a>

## tenant property — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / fdce0951c5dc / 6

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

<a id="canonical-8dfd0816fbf258d4776de0ba60f0990168299a1eb5132c7f519b53364de28e12"></a>

## Next pages — voltstack_cluster.global_network_list.global_network_connections.slo_to_global_d / fdce0951c5dc / 7

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-004.md#canonical-0f53dd34818b19d60b9a119e01c1a06fd3d4c92a1fb02b7caa10ec84131d3af0)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-42acdfe9767d853d8032ec5a385b380b28da930f26b46e0e7479964b8d553d3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-830140d7be79c3d378c0be6de682e93241764fa14d38980b04693933cb7447a8"></a>

## voltstack_cluster.k8s_cluster — voltstack_cluster.k8s_cluster / c2863ee4ee25 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- voltstack_cluster.k8s_cluster

<a id="canonical-9b8ac3341b1a0ff84de6cfe1c550cdc3725dbca9f07b34bfc73db737381c28db"></a>

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

<a id="canonical-d16db4362ff10a1e8758a0133ab94eebb896d53d6688daafc531cd652a155544"></a>

## Direct properties — voltstack_cluster.k8s_cluster / c2863ee4ee25 / 3

<a id="canonical-2108e0b5ad37ae8e44fa15cd7f623f4022b8940531df7759e45ea117186d42e0"></a>

<a id="canonical-48d3ffbb2333425bb5e635897dfbd7df839ee2155677a2d1f03e4cc0d512afc1"></a>

## name property — voltstack_cluster.k8s_cluster / c2863ee4ee25 / 4

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

<a id="canonical-04c21825666451762d37f51929b9f63b39d6857d94c0ac2af05870c8c1d3aa2c"></a>

<a id="canonical-c736a32a334767cfab3ed4a86605a4cee26d8bf0a205536bc41dd3f66ef18cd7"></a>

## namespace property — voltstack_cluster.k8s_cluster / c2863ee4ee25 / 5

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

<a id="canonical-d5072b7fe7def1509a0e13db6095613c4fd29c92bb1ad86c280256dc7457bae7"></a>

<a id="canonical-d2e010419822c915d67923c1230c1c501b2b4d12891c28ed8a358f50e63057c0"></a>

## tenant property — voltstack_cluster.k8s_cluster / c2863ee4ee25 / 6

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

<a id="canonical-a414bdc786af86fdac5dfe9ab006fa82f76f229cdfdb9a1be981fa064a73b1f9"></a>

## Next pages — voltstack_cluster.k8s_cluster / c2863ee4ee25 / 7

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)

<a id="canonical-617137e51ce90d1c4c0adf39c6e17957f864ba1df845dc2b4f661e4052034c9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-74467a9353d3fb6c54830b18e96d12803b57087c4c76d4b6b2afa294bb856b76"></a>

## voltstack_cluster.no_dc_cluster_group — voltstack_cluster.no_dc_cluster_group / bcac3073a7db / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-596185b78dcd7c81fad398b9cb95469116ab1857588f8fec08d4b5deb27d1d75)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-31bf91e3704c52909574bea812a91182baf60517dbc465c961fba35a136a753c)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-b9c0dc061e9044db2f839eab725e620c26d2611088a4e89f7ffa98e53ff49843)
- voltstack_cluster.no_dc_cluster_group

<a id="canonical-06013a72e64ecd35bca2a9c69923d8aec11e636fd527fce9795f9792aaf7290f"></a>

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

<a id="canonical-8b37254304605f12a3603835f0bbc2b0d0d6486d4e1c487f3a9e89d3ce22c58c"></a>

## Direct properties — voltstack_cluster.no_dc_cluster_group / bcac3073a7db / 3

This is an empty object or choice marker. It has no direct properties.
